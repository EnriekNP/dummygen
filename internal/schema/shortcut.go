package schema

var shortcuts = map[string]struct {
	BaseType string
	Faker    string
}{
	"email": {"string", "email"},
	"name":  {"string", "name"},
	"uuid":  {"string", "uuid"},
	"date":  {"date", "past_date"},
}
