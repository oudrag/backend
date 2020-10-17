package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/oudrag/server/internal/domain/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const usersCollection = "users"

var ctx = context.Background()

type userSchema struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	DisplayName  string             `bson:"display_name"`
	Avatar       string             `bson:"avatar"`
	RegisteredAt time.Time          `bson:"registered_at"`
	RegisterType string             `bson:"register_type"`
}

type UserRepository struct {
	db *mongo.Collection
}

func (r *UserRepository) Index() error {
	_, err := r.db.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	})

	return err
}

func (r *UserRepository) Load(id string) (*auth.User, error) {
	res := r.db.FindOne(ctx, bson.D{{"_id", id}})

	if err := res.Err(); err != nil {
		return nil, err
	}

	var doc *userSchema
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return convertDoc(doc), nil
}

func (r *UserRepository) LoadWithEmail(email string) (*auth.User, error) {
	res := r.db.FindOne(ctx, bson.D{{"email", email}})

	if err := res.Err(); err != nil {
		return nil, err
	}

	var doc *userSchema
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return convertDoc(doc), nil
}

func (r *UserRepository) Save(user *auth.User) error {
	res, err := r.db.InsertOne(ctx, convertToDoc(user))
	if err != nil {
		return err
	}

	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.SetID(id.String())
	}

	return nil
}

func convertToDoc(user *auth.User) userSchema {
	doc := userSchema{
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		DisplayName:  user.DisplayName,
		Avatar:       user.Avatar,
		RegisteredAt: time.Now(),
		RegisterType: user.RegisterType.Value(),
	}

	return doc
}

func convertDoc(doc *userSchema) *auth.User {
	return &auth.User{
		ID:           doc.ID.String(),
		Name:         doc.Name,
		Email:        doc.Email,
		Password:     doc.Password,
		DisplayName:  doc.DisplayName,
		Avatar:       doc.Avatar,
		RegisteredAt: doc.RegisteredAt,
		RegisterType: auth.RegisterTypeFromValue(doc.RegisterType),
	}
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db.Collection(usersCollection)}
}
