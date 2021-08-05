package database

// Query object
type Query struct {
	Type    string
	Query   string
	Params  []interface{}
	Closure bool
}
