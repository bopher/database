# Database

A set of database types, driver, query builder and paginator for sql based databases.

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

## Paginator

Advanced pagination manager with tag and meta support.

### Create Paginator

Constructor functions parse paginator inputs (page, limit, etc.) from query string.

```go
// Create a new paginator instance with initial values
NewPaginatorWithDefaults(limits []uint8, defaultLimit uint8, sorts []string, defaultSort string, queryString string) Paginator

// Create a new paginator instance with default params
// This function is equal NewPaginatorWithDefaults([]uint8{10, 25, 50, 100}, 25, []string{"id"}, "id", queryString)
NewPaginator(queryString string) Paginator
```

#### Query String Structure

Paginator get incoming parameters as Base64 encoded string (queryString). Query string contains following structure:

**Note:** If your query string not came as encoded string you can use manual method for setting paginator driver.

- Page (`uint`): page number
- Limit (`uint8`): limit
- Sort (`string`): sort
- Order (`string`): order asc or desc
- Search (`string`): search phrase
- Tags (`map[string]interface{}`): tags and filters.

#### Usage

Paginator interface contains following methods:

##### SetPage

Set current page.

```go
// Signature:
SetPage(page uint)
```

##### GetPage

Get current page.

```go
// Signature:
GetPage() uint
```

##### SetLimit

Set limit.

```go
SetLimit(limit uint8)
```

##### GetLimit

Get limit.

```go
// Signature:
GetLimit() uint8
```

##### SetSort

Set sort.

```go
SetSort(sort string)
```

##### GetSort

Get sort.

```go
GetSort() string
```

##### SetOrder

Set order.

```go
SetOrder(order string)
```

##### GetOrder

Get order.

```go
// Signature:
GetOrder() string
```

##### SetSearch

Set search phrase.

```go
// Signature:
SetSearch(search string)
```

##### GetSearch

Get search key

```go
// Signature:
GetSearch() string
```

##### Tags

Tags is a list of filters passed from client side for filtering query result.

###### SetTags

Set tags list

```go
// Signature:
SetTags(tags map[string]interface{})
```

###### GetTags

Get tags list

```go
// Signature:
GetTags() map[string]interface{}
```

###### SetTag

Set tag.

```go
// Signature:
SetTag(key string, value interface{})
```

###### GetTag

Get tag.

```go
// Signature:
GetTag(key string) interface{}
```

###### HasTag

Check if tag exists.

```go
// Signature:
HasTag(key string) bool
```

###### Get Tag With Type

For getting tags with type you can use helper getter methods. getter methods accept a fallback value and returns fallback if value not exists or not in type. getter methods follow ok pattern. Getter methods list:

```go
// SliceTag get slice tag or return fallback if tag not exists.
(key string, fallback []interface{}) ([]interface{}, bool)
// StringTag get string tag or return fallback if tag not exists
StringTag(key string, fallback string) (string, bool)
// StringSliceTag get string slice tag or return fallback if tag not exists
StringSliceTag(key string, fallback []string) ([]string, bool)
// BoolTag get bool tag or return fallback if tag not exists
BoolTag(key string, fallback bool) (bool, bool)
// BoolSliceTag get bool slice tag or return fallback if tag not exists
BoolSliceTag(key string, fallback []bool) ([]bool, bool)
// Float64Tag get float64 tag or return fallback if tag not exists
Float64Tag(key string, fallback float64) (float64, bool)
// Float64SliceTag get float64 slice tag or return fallback if tag not exists
Float64SliceTag(key string, fallback []float64) ([]float64, bool)
// Int64Tag get int64 tag or return fallback if tag not exists
Int64Tag(key string, fallback int64) (int64, bool)
// Int64SliceTag get int64 slice tag or return fallback if tag not exists
Int64SliceTag(key string, fallback []int64) ([]int64, bool)
```

##### Meta

Meta data are extra information attached to response.

###### SetMeta

Set meta data.

```go
// Signature:
SetMeta(key string, value interface{})
```

###### GetMeta

Get meta.

```go
// Signature:
GetMeta(key string) interface{}
```

###### HasMeta

Check if meta exists.

```go
// Signature:
HasMeta(key string) bool
```

###### MetaData

Get meta data list.

```go
// Signature:
MetaData() map[string]interface{}
```

###### Get Meta With Type

For getting tags with type you can use helper getter methods. getter methods accept a fallback value and returns fallback if value not exists or not in type. getter methods follow ok pattern. Getter methods list:

```go
// SliceMeta get slice meta or return fallback if meta not exists
SliceMeta(key string, fallback []interface{}) ([]interface{}, bool)
// StringMeta get string meta or return fallback if meta not exists
StringMeta(key string, fallback string) (string, bool)
// StringSliceMeta get string slice slice meta or return fallback if meta not exists
StringSliceMeta(key string, fallback []string) ([]string, bool)
// BoolMeta get bool meta or return fallback if meta not exists
BoolMeta(key string, fallback bool) (bool, bool)
// BoolSliceMeta get bool slice meta or return fallback if meta not exists
BoolSliceMeta(key string, fallback []bool) ([]bool, bool)
// Float64Meta get float64 meta or return fallback if meta not exists
Float64Meta(key string, fallback float64) (float64, bool)
// Float64SliceMeta get float64 slice meta or return fallback if meta not exists
Float64SliceMeta(key string, fallback []float64) ([]float64, bool)
// Int64Meta get int64 meta or return fallback if meta not exists
Int64Meta(key string, fallback int64) (int64, bool)
// Int64SliceMeta get int64 slice meta or return fallback if meta not exists
Int64SliceMeta(key string, fallback []int64) ([]int64, bool)
```

##### SetCount

Set total records count.

```go
SetCount(count uint64)
```

##### GetCount

Get records count.

```go
// Signature:
GetCount() uint64
```

##### From

Get from record position.

```go
// Signature:
From() uint64
```

##### To

Get to record position.

```go
// Signature:
To() uint64
```

##### Total

Get total pages.

```go
// Signature:
Total() uint
```

##### SQL

Get sql order and limit command as string.

```go
// Signature:
SQL() string
```

##### Response

Get response for json.

Response returns following data:

**Note:** Meta data added to this response.

```json
{
  "page": 0,
  "limit": 0,
  "sort": "",
  "order": "",
  "search": "",
  "count": 0,
  "from": 0,
  "to": 0,
  "total": 0
}
```

```go
// Signature:
Response() map[string]interface{}
```
