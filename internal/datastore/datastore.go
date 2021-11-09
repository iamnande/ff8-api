package datastore

import (
	"embed"
	"encoding/json"
)

//go:generate mockgen -destination=./mocks/datastore.go -package mocks github.com/iamnande/ff8-api/internal/datastore Datastore

// Datastore interface describes the methods that all Datastore providers
// must satisfy.
type Datastore interface {
	DescribeRecord(name string, recordType RecordType) (*Record, error)
}

// datastore is the currently used implementation of the Datastore.
type datastore struct {
	Records Records `json:"records"`
}

// compile time validation that the current implementation satisfies the
// defined interface.
var _ Datastore = (*datastore)(nil)

var (
	staticDatastoreFile = ".datastore.json"

	//go:embed .datastore.json
	staticDatastoreContent embed.FS
)

// NewDatastore creates a fresh in stance of a static, file based, Datastore
// implementation.
// TODO: test this package.
func NewDatastore() (Datastore, error) {

	// new: read embedded datastore file
	rawContent, err := staticDatastoreContent.ReadFile(staticDatastoreFile)
	if err != nil {
		return nil, err
	}

	// new: deserialize into datastore structure
	records := new(datastore)
	if err = json.Unmarshal(rawContent, records); err != nil {
		return nil, err
	}

	// new: return new instance of the datastore
	return records, nil

}

func (ds datastore) DescribeRecord(name string, recordType RecordType) (*Record, error) {

	//  describe: iterate datastore looking for record
	for _, record := range ds.Records {
		if record.Name == name && record.Type == recordType {
			return record, nil
		}
	}

	// describe: return complete record set
	return nil, ErrItemNotFound

}
