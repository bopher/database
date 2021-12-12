# Database

A set of database types, driver and query builder for sql based databases.

## Nullable Types

database package contains nullable datatype for working with nullable data. nullable types implements **Scanners**, **Valuers**, **Marshaler** and **Unmarshaler** interfaces.

**Note:** You can use `Val` method to get variable nullable value.

**Note:** Slice types is a comma separated list of variable that stored as string in database. e.g.: "1,2,3,4"

### Available Nullable Types

```go
import "github.com/bopher/database/types"
var a types.NullBool
var a types.NullFloat32
var a types.Float32Slice
var a types.NullFloat64
var a types.Float64Slice
var a types.NullInt
var a types.IntSlice
var a types.NullInt8
var a types.Int8Slice
var a types.NullInt16
var a types.Int16Slice
var a types.NullInt32
var a types.Int32Slice
var a types.NullInt64
var a types.Int64Slice
var a types.NullString
var a types.StringSlice
var a types.NullTime
var a types.NullUInt
var a types.UIntSlice
var a types.NullUInt8
var a types.UInt8Slice
var a types.NullUInt16
var a types.UInt16Slice
var a types.NullUInt32
var a types.UInt32Slice
var a types.NullUInt64
var a types.UInt64Slice
```

## MySQL Driver

Create new MySQL connection. this function return a `"github.com/jmoiron/sqlx"` instance.

```go
// Signature:
NewMySQLConnector(host string, username string, password string, database string) (*sqlx.DB, error)

// Example:
import "github.com/bopher/database"
db, err := database.NewMySQLConnector("", "root", "root", "myDB")
```

## Query Builder

Make complex query use `Query` structure.

**Note:** You can use special `@in` keyword in your query and query builder make a `IN(params)` query for you.

### Query Builder Methods

#### Add

Add new query.

```go
// Signature:
Add(q Query)
```

#### Query

Get query string.

```go
// Signature:
Query() string
```

#### Params

Get query builder parameters.

```go
// Signature:
Params() []interface{}
```

### Query Structure Fields

**Type** _(String)_: Determine query type `AND`, `OR`, etc.

**Query** _(String)_: Query string.

**Params** _[]interface{}_: Query parameters.

**Closure** _bool_: Determine query is sub query or not.

```go
import "github.com/bopher/database"
import "fmt"
var qBuilder database.QueryBuilder
qBuilder.Add(database.Query{
    Query:  "firstname LIKE '%?%'",
    Params: []interface{}{"john"},
})
qBuilder.Add(database.Query{
    Type: "AND",
    Query:  "role @in",
    Params: []interface{}{"admin", "support", "user"},
})
qBuilder.Add(database.Query{
    Type: "AND",
    Query:  "age > ? AND age < ?",
    Params: []interface{}{15, 30},
    Closure: true,
})
fmt.Print(qBuilder.Query()) // firstname LIKE '%?%' AND role IN(?, ?, ?) AND (age > ? AND age < ?)
fmt.Print(qBuilder.Params()) // ["john", "admin", "support", "user", 15, 30]
```
