# DummyGen

Generate dummy data from JSON/XML schema models for MySQL (Golang).  

> ⚠️ Project is in active development (v1).


---

## Features (v1)

- Define schema with tables and columns
- Supports common shortcuts: `email`, `name`, `uuid`, `date`
- Primary key, unique, nullable, and auto_increment support
- Normalization ensures consistent types and flags

---

## Project Structure
```
dummygen/
│── cmd/ # CLI commands (planned)
│── internal/
│ ├── schema/ # Schema structs + normalization
│ └── generator/ # SQL / dummy data generator
│── go.mod
```

---

## Getting Started

```bash
git clone https://github.com/EnriekNP/dummygen.git
cd dummygen
go mod tidy
```
- Use the internal packages (schema, generator) to load schema, normalize, and generate SQL/dummy data.
- CLI and full generator are in development.

---

## Example Schema (JSON)
```json
{
  "table": "users",
  "count": 5,
  "columns": [
    { "name": "id", "type": "int", "primary": true, "auto_increment": true },
    { "name": "email", "type": "email" },
    { "name": "name", "type": "name" },
    { "name": "created_at", "type": "date" }
  ]
}
```