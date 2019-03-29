package bolt

import (
	"encoding/json"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
)

// Store Store
type Store interface {
	Create(key []byte, v interface{}, bucketName string) error
	Get(key []byte, v interface{}, bucket string) error
}

type store struct {
	Volume string
	Name   string
}

// New creates a new bolt store
func New(dumbPath, dbName string) Store {
	return &store{Volume: dumbPath, Name: dbName}
}

func connect(p, dbName string) (*bolt.DB, error) {
	h, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	rc := path.Join(h, p)

	if _, err := os.Stat(rc); os.IsNotExist(err) {
		os.Mkdir(rc, os.ModePerm)
	}

	return bolt.Open(path.Join(rc, dbName), 0600, &bolt.Options{Timeout: 1 * time.Second})
}

func (s *store) Create(key []byte, v interface{}, bucketName string) error {
	db, err := connect(s.Volume, s.Name)
	if err != nil {
		return err
	}
	defer db.Close()
	bArr, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return b.Put(key, bArr)
	})
}

func (s *store) Get(key []byte, v interface{}, bucket string) error {
	db, err := connect(s.Volume, s.Name)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.View(func(tx *bolt.Tx) error {
		bck := tx.Bucket([]byte(bucket))
		if bck == nil {
			return ErrBucketNotFound
		}

		out := bck.Get(key)
		if out == nil {
			return ErrValueDoesntExist
		}
		return json.Unmarshal(out, &v)
	})

}
