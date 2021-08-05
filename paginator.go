package database

// Paginator interface
type Paginator interface {
	// SetPage set current page
	SetPage(page uint)
	// GetPage get current page
	GetPage() uint
	// SetLimit set limit
	SetLimit(limit uint8)
	// GetLimit get limit
	GetLimit() uint8
	// SetSort set sort
	SetSort(sort string)
	// GetSort get sort
	GetSort() string
	// SetOrder set order
	SetOrder(order string)
	// GetOrder get order
	GetOrder() string
	// SetSearch set search phrase
	SetSearch(search string)
	// GetSearch get search phrase
	GetSearch() string
	// SetTags set tags list
	SetTags(tags map[string]interface{})
	// GetTags get tags list
	GetTags() map[string]interface{}
	// SetTag set tag
	SetTag(key string, value interface{})
	// GetTag get tag
	GetTag(key string) interface{}
	// HasTag check if tag exists
	HasTag(key string) bool
	// SliceTag get slice tag or return fallback if tag not exists
	SliceTag(key string, fallback []interface{}) ([]interface{}, bool)
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
	// SetMeta set meta data
	SetMeta(key string, value interface{})
	// GetMeta get meta
	GetMeta(key string) interface{}
	// HasMeta check if meta exists
	HasMeta(key string) bool
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
	// MetaData get meta data list
	MetaData() map[string]interface{}
	// SetCount Set total records count
	SetCount(count uint64)
	// GetCount get records count
	GetCount() uint64
	// From get from record position
	From() uint64
	// To get to record position
	To() uint64
	// Total get total pages
	Total() uint
	// SQL get sql order and limit command as string
	SQL() string
	// Response get response for json
	Response() map[string]interface{}
}
