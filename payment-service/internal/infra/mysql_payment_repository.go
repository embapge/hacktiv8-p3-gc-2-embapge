package infra

import (
	"context"
	"p3-graded-challenge-2-embapge/payment-service/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type paymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(collection *mongo.Collection) domain.PaymentRepository {
	return &paymentRepository{collection: collection}
}

func (r *paymentRepository) Create(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = payment.CreatedAt
	payment.Status = "pending"
	result, err := r.collection.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}
	payment.ID = result.InsertedID.(primitive.ObjectID)
	return payment, nil
}

func (r *paymentRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&payment)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}
func (r *paymentRepository) GetAll(ctx context.Context) (*[]domain.Payment, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var payments []domain.Payment
	for cursor.Next(ctx) {
		var payment domain.Payment
		if err := cursor.Decode(&payment); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &payments, nil
}

func (r *paymentRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
