package datastore

type RecordType string

var (
	Magic      RecordType = "Magic"
	LimitBreak RecordType = "Limit Break"
)

type Records []*Record

type Record struct {
	Name           string     `json:"name"`
	Type           RecordType `json:"type"`
	CardEquivalent string     `json:"card_equivalent"`
	CardMagicRatio float64    `json:"card_magic_ratio"`
}

// Datastore is a generic datastore definition for implementations to satisfy.
type Datastore interface {
	DescribeRecord(name string, recordType RecordType) (*Record, error)
}
