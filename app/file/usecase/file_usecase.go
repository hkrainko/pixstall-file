package usecase

import (
	"context"
	"log"
	error2 "pixstall-file/domain/error"
	"pixstall-file/domain/file"
	"pixstall-file/domain/file/acl"
	"pixstall-file/domain/file/acl/model"
	model2 "pixstall-file/domain/file/model"
	image_processing "pixstall-file/domain/image/image-processing"
	"strings"
)

type fileUseCase struct {
	fileRepo file.Repo
	fileAclRepo acl.Repo
	imageProcessingRepo image_processing.Repo
}

func NewFileUseCase(fileRepo file.Repo, fileAclRepo acl.Repo, imageProcessingRepo image_processing.Repo) file.UseCase {
	return fileUseCase{
		fileRepo: fileRepo,
		fileAclRepo: fileAclRepo,
		imageProcessingRepo: imageProcessingRepo,
	}
}

func (f fileUseCase) SaveFile(ctx context.Context, fileData *[]byte, fileType model2.FileType, name string, owner string, acl []string) (*string, error) {

	files, err := NewFileFactory(fileType, f.imageProcessingRepo).getFiles(fileData, fileType, name)
	if err != nil || len(*files) <= 0 {
		return nil, error2.UnknownError
	}
	dFiles := *files
	paths, err := f.fileRepo.SaveFiles(ctx, dFiles)
	if err != nil {
		return nil, error2.UnknownError
	}
	log.Printf("SaveFile into paths:%v\n", paths)
	aclMap := make(map[string]bool)
	for _, v := range acl {
		aclMap[v] = true
	}
	fileACL := model.FileACL{
		ID:    dFiles[0].ID,
		Owner: owner,
		ACL:   aclMap,
		State: model.FileStateActive,
	}
	_, err = f.fileAclRepo.AddFileACL(ctx, fileACL)
	if err != nil {
		return nil, err
	}

	return &dFiles[0].RawPath, nil
}

func (f fileUseCase) IsAccessible(ctx context.Context, accessUserID *string, fileType model2.FileType, prefix string) (*bool, error) {
	if !fileType.IsValid() {
		return nil, error2.NotFoundError
	}
	if f.isPermanentPublic(fileType) {
		result := true
		return &result, nil
	}
	fileID := f.removeSizeSuffix(prefix)
	fileAcl, err := f.fileAclRepo.GetFileACL(ctx, fileID)
	if err != nil {
		return nil, err
	}
	if *accessUserID == fileAcl.Owner {
		result := true
		return &result, nil
	}
	if _, ok := fileAcl.ACL["*"]; ok {
		result := true
		return &result, nil
	}
	if _, ok := fileAcl.ACL[*accessUserID]; ok {
		result := true
		return &result, nil
	}
	result := false
	return &result, nil
}

func (f fileUseCase) isInitialPublic(fileType model2.FileType) bool {
	switch fileType {
	case model2.FileTypeArtwork,
	model2.FileTypeRoof,
	model2.FileTypeOpenCommission,
	model2.FileTypeProfile:
		return true
	default:
		return false
	}
}

func (f fileUseCase) isPermanentPublic(fileType model2.FileType) bool {
	switch fileType {
	case model2.FileTypeRoof,
		model2.FileTypeProfile:
		return true
	default:
		return false
	}
}

func (f fileUseCase) removeSizeSuffix(prefix string) string {
	ss := strings.Split(prefix, "_")
	if len(ss) <= 1 {
		return prefix
	}
	return ss[0]
}