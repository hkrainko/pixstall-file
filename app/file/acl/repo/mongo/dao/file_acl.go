package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-file/domain/file/acl/model"
)

type FileACL struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	ID       string             `bson:"id"`
	Owner    string             `bson:"owner"`
	ACL      map[string]bool    `bson:"acl"`
	IsPublic bool               `bson:"isPublic"`
	State    model.FileState    `bson:"state"`
}

func (f FileACL) ToDomainFileAcl() model.FileACL {
	return model.FileACL{
		ID:       f.ID,
		Owner:    f.Owner,
		ACL:      f.ACL,
		IsPublic: f.IsPublic,
		State:    f.State,
	}
}
