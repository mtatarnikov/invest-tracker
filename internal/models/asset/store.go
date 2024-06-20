package asset

import (
	"invest-tracker/pkg/storage"
)

type AssetStore interface {
	Save(db storage.Database, a Asset) error
}

type RealAssetStore struct{}

func (s *RealAssetStore) Save(db storage.Database, a Asset) error {
	Save(db, a)
	return nil
}
