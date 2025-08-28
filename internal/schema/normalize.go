package schema

func NormalizeTable(t *Table) {
	for i, col := range t.Columns {
		if shortcut, ok := shortcuts[col.Type]; ok {
			if col.Faker == "" {
				col.Faker = shortcut.Faker
			}
			col.Type = shortcut.BaseType
		}

		// Default: nullable = false unless explicitly set true
		if !col.Nullable {
			col.Nullable = false // default not null
		}

		if col.Primary {
			col.Nullable = false // PK must be NOT NULL
			col.Unique = true    // PK must be UNIQUE
			// auto_increment remains opt-in
		}

		t.Columns[i] = col
	}
}
