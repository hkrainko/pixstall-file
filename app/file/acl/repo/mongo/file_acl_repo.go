package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-file/app/file/acl/repo/mongo/dao"
	error2 "pixstall-file/domain/error"
	"pixstall-file/domain/file/acl"
	"pixstall-file/domain/file/acl/model"
)

type mongoFileAclRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	FileAclCollection = "FileAcl"
)

func NewMongoFileAclRepo(db *mongo.Database) acl.Repo {
	return &mongoFileAclRepo{
		db:         db,
		collection: db.Collection(FileAclCollection),
	}
}

func (m mongoFileAclRepo) AddFileACL(ctx context.Context, fileACL model.FileACL) (*string, error) {
	daoFileACL := dao.FileACL{
		ObjectID: primitive.ObjectID{},
		ID:       fileACL.ID,
		Owner:    fileACL.Owner,
		ACL:      fileACL.ACL,
		State:    fileACL.State,
	}
	result, err := m.collection.InsertOne(ctx, daoFileACL)
	if err != nil {
		fmt.Printf("AddFileACL error %v\n", err)
		return nil, err
	}
	fmt.Printf("AddFileACL %v", result.InsertedID)
	dFileACL := daoFileACL.ToDomainFileAcl()
	return &dFileACL.ID, nil
}

func (m mongoFileAclRepo) GetFileACL(ctx context.Context, id string) (*model.FileACL, error) {
	daoFileACL := dao.FileACL{}
	err := m.collection.FindOne(ctx, bson.M{"id": id}, nil).Decode(&daoFileACL)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, error2.NotFoundError
		default:
			return nil, error2.UnknownError
		}
	}
	dFileACL := daoFileACL.ToDomainFileAcl()
	return &dFileACL, nil
}

func (m mongoFileAclRepo) UpdateFileACL(ctx context.Context, id string, updater model.FileACLUpdater) error {
	filter := bson.M{
		"id": id,
	}
	u := bson.D{}
	setter := bson.D{}
	if updater.State != nil {
		setter = append(setter, bson.E{Key: "state", Value: updater.State})
	}
	u = append(u, bson.E{Key: "$set", Value: setter})
	_, err := m.collection.UpdateOne(ctx, filter, u)
	if err != nil {
		return error2.UnknownError
	}
	return nil
}
