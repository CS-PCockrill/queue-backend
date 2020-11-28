package mongodb

import (
	"github.com/CS-PCockrill/queue/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type StoreFunctions struct{
	CLIENT *mongo.Client
}
func (s *StoreFunctions) RegisterStore(store *models.Store) (int, error) {
	// Register a store after registering a user...
	return 0, nil
}

func (s *StoreFunctions) InsertOneProduct(product *models.Item) (int, error) {
	// Insert just one product if there is only 1
	return 0, nil
}

func (s *StoreFunctions) InsertManyProducts(products ...*models.Item) (int, error) {
	// Insert many products if there is an array of products
	return 0, nil
}

