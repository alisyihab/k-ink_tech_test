package repo

import (
	"context"
	"registration-smart-reward/internal/config"
	"registration-smart-reward/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Insert(ctx context.Context, user *model.User) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() UserRepository {
	return &userRepository{
		collection: config.GetMongoDB().Collection("users"),
	}
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Insert(ctx context.Context, user *model.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
