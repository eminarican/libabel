package session

import (
	"encoding/json"
	"github.com/df-mc/goleveldb/leveldb"
	"github.com/df-mc/goleveldb/leveldb/opt"
	"github.com/google/uuid"
	"os"
)

type Provider struct {
	db *leveldb.DB
}

func NewProvider(path string) (*Provider, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}
	db, err := leveldb.OpenFile(path, &opt.Options{Compression: opt.SnappyCompression})
	if err != nil {
		return nil, err
	}
	return &Provider{db: db}, nil
}

func (p *Provider) Save(id uuid.UUID, d Data) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return p.db.Put(id[:], b, nil)
}

func (p *Provider) Load(id uuid.UUID) (Data, error) {
	b, err := p.db.Get(id[:], nil)
	if err != nil {
		return Data{}, err
	}
	var d Data
	err = json.Unmarshal(b, &d)
	if err != nil {
		return Data{}, err
	}
	return d, nil
}

func (p *Provider) Close() error {
	return p.db.Close()
}
