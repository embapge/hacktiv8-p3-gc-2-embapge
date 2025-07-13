package repository

import (
	"context"
	"time"

	"p3-graded-challenge-1-embapge/shopping-service/internal/domain"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) repository.ProductRepository {
	return &productRepository{
		collection: collection,
	}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	product.CreatedAt = time.Now()
	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}

func (r *productRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
	var product domain.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, id primitive.ObjectID, product *domain.Product) (*domain.Product, error) {
	product.UpdatedAt = time.Now()
	update := bson.M{
		"$set": product,
	}
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}
	return r.FindByID(ctx, id)
}

func (r *productRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
