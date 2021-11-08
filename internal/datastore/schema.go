package datastore

// RecordType represents the type of item is being requested.
type RecordType string

var (
	// Magic is an RecordType ENUM value of "Magic". Indicates the card
	// conversation is for magic.
	Magic RecordType = "Magic"

	// LimitBreak is an RecordType ENUM value of "Limit Break". Indicates the
	// card conversation is for one or more Limit Breaks.
	LimitBreak RecordType = "Limit Break"
)

// Records is a collection of datastore records.
type Records []*Record

// Record represents a single datastore record.
type Record struct {
	Name           string     `json:"name"`
	Type           RecordType `json:"type"`
	CardEquivalent string     `json:"card_equivalent"`
	CardMagicRatio float64    `json:"card_magic_ratio"`
}
