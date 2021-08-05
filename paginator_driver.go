package database

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

type parsableRecord struct {
	Page   uint
	Limit  uint8
	Sort   string
	Order  string
	Search string
	Tags   map[string]interface{}
}

type pagination struct {
	page   uint
	limit  uint8
	sort   string
	order  string
	search string
	tags   map[string]interface{}

	limits []uint8
	sorts  []string
	meta   map[string]interface{}

	count uint64
	from  uint64
	to    uint64
	total uint
}

func (p *pagination) validateLimit(limit uint8) bool {
	if p.limits != nil && len(p.limits) > 0 {
		for _, l := range p.limits {
			if limit == l {
				return true
			}
		}
	}
	return false
}

func (p *pagination) validateSort(sort string) bool {
	if p.sorts != nil && len(p.sorts) > 0 {
		for _, s := range p.sorts {
			if sort == s {
				return true
			}
		}
	}
	return false
}

func (p *pagination) init(limits []uint8, defaultLimit uint8, sorts []string, defaultSort string, queryString string) error {
	// initialize
	p.page = 1
	p.limit = 25
	p.sort = "id"
	p.order = "asc"
	p.tags = make(map[string]interface{}, 0)
	p.limits = make([]uint8, 0)
	p.sorts = make([]string, 0)
	p.meta = make(map[string]interface{}, 0)
	// proccess input
	if limits != nil {
		p.limits = limits
	}
	if p.validateLimit(defaultLimit) {
		p.limit = defaultLimit
	}

	if sorts != nil {
		p.sorts = sorts
	}
	if p.validateSort(defaultSort) {
		p.sort = defaultSort
	}
	// parse query string
	var base64decoded []byte
	if _, err := base64.StdEncoding.Decode(base64decoded, []byte(queryString)); err != nil {
		return err
	}

	data := parsableRecord{}
	if err := json.Unmarshal(base64decoded, &data); err != nil {
		return nil
	}

	if data.Page > 0 {
		p.page = data.Page
	}

	if p.validateLimit(data.Limit) {
		p.limit = data.Limit
	}

	if p.validateSort(data.Sort) {
		p.sort = data.Sort
	}

	if data.Order == "asc" || data.Order == "desc" {
		p.order = strings.ToUpper(data.Order)
	}

	p.search = data.Search

	if data.Tags != nil {
		p.tags = data.Tags
	}

	return nil
}

// SetPage set current page
func (p *pagination) SetPage(page uint) {
	p.page = page
}

// GetPage get current page
func (p *pagination) GetPage() uint {
	return p.page
}

// SetLimit set limit
func (p *pagination) SetLimit(limit uint8) {
	if p.validateLimit(limit) {
		p.limit = limit
	}
}

// GetLimit get limit
func (p *pagination) GetLimit() uint8 {
	return p.limit
}

// SetSort set sort
func (p *pagination) SetSort(sort string) {
	if p.validateSort(sort) {
		p.sort = sort
	}
}

// GetSort get sort
func (p *pagination) GetSort() string {
	return p.sort
}

// SetOrder set order
func (p *pagination) SetOrder(order string) {
	if order == "asc" || order == "desc" {
		p.order = order
	}
}

// GetOrder get order
func (p *pagination) GetOrder() string {
	return p.order
}

// SetSearch set search key
func (p *pagination) SetSearch(search string) {
	p.search = search
}

// GetSearch get search key
func (p *pagination) GetSearch() string {
	return p.search
}

// Settags set tags list
func (p *pagination) SetTags(tags map[string]interface{}) {
	if tags != nil {
		p.tags = tags
	}
}

// Gettags get tags list
func (p *pagination) GetTags() map[string]interface{} {
	return p.tags
}

// SetTag set tag
func (p *pagination) SetTag(key string, value interface{}) {
	p.tags[key] = value
}

// GetTag get tag
func (p *pagination) GetTag(key string) interface{} {
	return p.tags[key]
}

// HasTag check if tag exists
func (p *pagination) HasTag(key string) bool {
	_, ok := p.tags[key]
	return ok
}

// SliceTag get slice tag or return fallback if tag not exists
func (p *pagination) SliceTag(key string, fallback []interface{}) ([]interface{}, bool) {
	if val, ok := p.tags[key]; ok {
		if sliceVal, ok := val.([]interface{}); ok {
			return sliceVal, true
		}
	}
	return fallback, false
}

// StringTag get string tag or return fallback if tag not exists
func (p *pagination) StringTag(key string, fallback string) (string, bool) {
	if val, ok := p.tags[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// StringSliceTag get string slice tag or return fallback if tag not exists
func (p *pagination) StringSliceTag(key string, fallback []string) ([]string, bool) {
	if val, ok := p.tags[key]; ok {
		if strVal, ok := val.([]string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// BoolTag get bool tag or return fallback if tag not exists
func (p *pagination) BoolTag(key string, fallback bool) (bool, bool) {
	if val, ok := p.tags[key]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// BoolSliceTag get bool slice tag or return fallback if tag not exists
func (p *pagination) BoolSliceTag(key string, fallback []bool) ([]bool, bool) {
	if val, ok := p.tags[key]; ok {
		if boolVal, ok := val.([]bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// Float64Tag get float64 tag or return fallback if tag not exists
func (p *pagination) Float64Tag(key string, fallback float64) (float64, bool) {
	if val, ok := p.tags[key]; ok {
		if floatVal, ok := val.(float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Float64SliceTag get float64 slice tag or return fallback if tag not exists
func (p *pagination) Float64SliceTag(key string, fallback []float64) ([]float64, bool) {
	if val, ok := p.tags[key]; ok {
		if floatVal, ok := val.([]float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Int64Tag get int64 tag or return fallback if tag not exists
func (p *pagination) Int64Tag(key string, fallback int64) (int64, bool) {
	if val, ok := p.tags[key]; ok {
		if intVal, ok := val.(int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// Int64SliceTag get int64 slice tag or return fallback if tag not exists
func (p *pagination) Int64SliceTag(key string, fallback []int64) ([]int64, bool) {
	if val, ok := p.tags[key]; ok {
		if intVal, ok := val.([]int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// SetMeta set meta data
func (p *pagination) SetMeta(key string, value interface{}) {
	p.meta[key] = value
}

// GetMeta get meta
func (p *pagination) GetMeta(key string) interface{} {
	return p.meta[key]
}

// HasMeta check if meta exists
func (p *pagination) HasMeta(key string) bool {
	_, ok := p.meta[key]
	return ok
}

// SliceMeta get slice meta or return fallback if meta not exists
func (p *pagination) SliceMeta(key string, fallback []interface{}) ([]interface{}, bool) {
	if val, ok := p.meta[key]; ok {
		if sliceVal, ok := val.([]interface{}); ok {
			return sliceVal, true
		}
	}
	return fallback, false
}

// StringMeta get string meta or return fallback if meta not exists
func (p *pagination) StringMeta(key string, fallback string) (string, bool) {
	if val, ok := p.meta[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// StringSliceMeta get string slice slice meta or return fallback if meta not exists
func (p *pagination) StringSliceMeta(key string, fallback []string) ([]string, bool) {
	if val, ok := p.meta[key]; ok {
		if strVal, ok := val.([]string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// BoolMeta get bool meta or return fallback if meta not exists
func (p *pagination) BoolMeta(key string, fallback bool) (bool, bool) {
	if val, ok := p.meta[key]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// BoolSliceMeta get bool slice meta or return fallback if meta not exists
func (p *pagination) BoolSliceMeta(key string, fallback []bool) ([]bool, bool) {
	if val, ok := p.meta[key]; ok {
		if boolVal, ok := val.([]bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// Float64Meta get float64 meta or return fallback if meta not exists
func (p *pagination) Float64Meta(key string, fallback float64) (float64, bool) {
	if val, ok := p.meta[key]; ok {
		if floatVal, ok := val.(float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Float64SliceMeta get float64 slice meta or return fallback if meta not exists
func (p *pagination) Float64SliceMeta(key string, fallback []float64) ([]float64, bool) {
	if val, ok := p.meta[key]; ok {
		if floatVal, ok := val.([]float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Int64Meta get int64 meta or return fallback if meta not exists
func (p *pagination) Int64Meta(key string, fallback int64) (int64, bool) {
	if val, ok := p.meta[key]; ok {
		if intVal, ok := val.(int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// Int64SliceMeta get int64 slice meta or return fallback if meta not exists
func (p *pagination) Int64SliceMeta(key string, fallback []int64) ([]int64, bool) {
	if val, ok := p.meta[key]; ok {
		if intVal, ok := val.([]int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// MetaData get meta data list
func (p *pagination) MetaData() map[string]interface{} {
	return p.meta
}

// SetCount Set total records count
func (p *pagination) SetCount(count uint64) {
	p.count = count
	p.total = uint(math.Ceil(float64(p.count) / float64(p.limit)))
	if p.page > p.total {
		p.page = p.total
	}
	if p.page < 1 {
		p.page = 1
	}

	p.from = (uint64(p.page-1) * uint64(p.limit)) + 1
	if p.from < 1 {
		p.from = 1
	}

	p.to = p.from + uint64(p.limit)
	if p.to > count {
		p.to = count
	}
}

// GetCount get records count
func (p *pagination) GetCount() uint64 {
	return p.count
}

// From get from record position
func (p *pagination) From() uint64 {
	return p.from
}

// To get to record position
func (p *pagination) To() uint64 {
	return p.to
}

// Total get total pages
func (p *pagination) Total() uint {
	return p.total
}

// SQL get sql order and limit command as string
func (p *pagination) SQL() string {
	return fmt.Sprintf(" ORDER BY %s %s LIMIT %d, %d", p.sort, p.order, p.from-1, p.limit)
}

// Response get response for json
func (p *pagination) Response() map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range p.meta {
		res[k] = v
	}
	res["page"] = p.page
	res["limit"] = p.limit
	res["sort"] = p.sort
	res["order"] = p.order
	res["search"] = p.search
	res["count"] = p.count
	res["from"] = p.from
	res["to"] = p.to
	res["total"] = p.total
	return res
}
