package datastore

import (
	"embed"
	"encoding/json"
	"fmt"
)

// staticDatastore is the current internal implementation
type staticDatastore struct {
	Records Records `json:"records"`
}

var (
	staticDatastoreFile = ".datastore.json"

	//go:embed .datastore.json
	staticDatastoreContent embed.FS
)

func NewStaticDS() (Datastore, error) {

	// static: read embedded datastore file
	rawContent, err := staticDatastoreContent.ReadFile(staticDatastoreFile)
	if err != nil {
		return nil, err
	}

	// static: deserialize into datastore structure
	records := new(staticDatastore)
	if err = json.Unmarshal(rawContent, records); err != nil {
		return nil, err
	}

	return records, nil
}

func (ds staticDatastore) DescribeRecord(name string, recordType RecordType) (*Record, error) {

	//  describe: iterate datastore looking for record
	for _, record := range ds.Records {
		if record.Name == name && record.Type == recordType {
			return record, nil
		}
	}

	// describe: return complete record set
	return nil, fmt.Errorf("%s not found", name)

}
