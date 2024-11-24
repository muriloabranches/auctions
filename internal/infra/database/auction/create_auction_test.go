package auction

import (
	"context"
	"testing"
	"time"

	"github.com/muriloabranches/auctions/internal/entity/auction_entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateAuction(t *testing.T) {
	t.Setenv("AUCTION_INTERVAL", "1s")

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("Should create and complete auction after interval", func(t *mtest.T) {
		t.AddMockResponses(mtest.CreateSuccessResponse())

		repo := &AuctionRepository{Collection: t.Coll}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := repo.CreateAuction(ctx, &auction_entity.Auction{
			Id:          uuid.New().String(),
			ProductName: "Product 1",
			Category:    "Category 1",
			Description: "Description 1",
			Condition:   auction_entity.New,
			Status:      auction_entity.Active,
			Timestamp:   time.Now(),
		})
		assert.Nil(t, err, "Auction creation expected to succeed")

		ev := t.GetSucceededEvent()
		assert.NotNil(t, ev, "Expected an event after auction creation")
		assert.Equal(t, "insert", ev.CommandName, "Expected 'insert' event after auction creation")
		assert.Nil(t, t.GetSucceededEvent(), "No more events expected immediately after creation")

		t.AddMockResponses(mtest.CreateSuccessResponse())
		time.Sleep(time.Second)

		ev = t.GetSucceededEvent()
		assert.NotNil(t, ev, "Expected an update event after auction completion")
		assert.Equal(t, "update", ev.CommandName, "Expected 'update' event after auction completion")
	})
}
