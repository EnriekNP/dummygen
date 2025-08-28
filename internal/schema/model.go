package schema

type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Faker         string `json:"faker,omitempty"`
	Primary       bool   `json:"primary,omitempty"`
	Unique        bool   `json:"unique"`
	Nullable      bool   `json:"nullable"`
	AutoIncrement bool   `json:"auto_increment,omitempty"`
}

type Table struct {
	Name    string   `json:"table"`
	Count   int      `json:"count"`
	Columns []Column `json:"columns"`
}
