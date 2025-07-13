package repository

import (
	"context"
	"p3-graded-challenge-1-embapge/shopping-service/internal/domain"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(collection *mongo.Collection) repository.TransactionRepository {
	return &transactionRepository{collection: collection}
}

func (r *transactionRepository) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = transaction.CreatedAt
	result, err := r.collection.InsertOne(ctx, transaction)
	if err != nil {
		return nil, err
	}
	transaction.ID = result.InsertedID.(primitive.ObjectID)
	return transaction, nil
}

func (r *transactionRepository) FindAll(ctx context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var transaction domain.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *transactionRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) Update(ctx context.Context, id primitive.ObjectID, transaction *domain.Transaction) (*domain.Transaction, error) {
	transaction.UpdatedAt = time.Now()
	update := bson.M{"$set": transaction}
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}
	return r.FindByID(ctx, id)
}

func (r *transactionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
