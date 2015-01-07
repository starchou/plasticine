// Copyright 2015 star Chou, All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package plasticine

import "github.com/bitly/go-simplejson"

type SearchFilter struct {
	Json *simplejson.Json
}

func Filter() *SearchFilter {
	return &SearchFilter{simplejson.New()}
}
func (j *SearchFilter) Filter(v interface{}) *SearchFilter {
	j.Json.Set("filter", v)
	return j
}

func (j *SearchFilter) Encode() []byte {
	var b []byte
	var err error
	if Debug {
		b, err = j.Json.EncodePretty()
	} else {
		b, err = j.Json.Encode()
	}
	if err != nil {
		return nil
	}
	return b
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-filtered-query.html
type FilteredQuery struct {
	Json *simplejson.Json `json:"filtered"`
}

func Filtered() *FilteredQuery {
	return &FilteredQuery{simplejson.New()}
}
func (j *FilteredQuery) Query(v ...interface{}) *FilteredQuery {
	if len(v) > 1 {
		j.Json.Set("query", v)
	} else {
		j.Json.Set("query", v[0])
	}
	return j
}

func (j *FilteredQuery) Filter(v ...interface{}) *FilteredQuery {
	if len(v) > 1 {
		j.Json.Set("filter", v)
	} else {
		j.Json.Set("filter", v[0])
	}
	return j
}

// strategy parameter (leap_frog_query_first,leap_frog_filter_first,leap_frog,query_first,random_access_${threshold},random_access_always)
//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-filtered-query.html#_filter_strategy
func (j *FilteredQuery) Strategy(v string) *FilteredQuery {
	j.Json.Set("strategy", v)
	return j
}

func (j *FilteredQuery) Encode() []byte {
	return encode(j)
}

func Limit(v int) *simplejson.Json {
	j := simplejson.New()
	j.SetPath([]string{"limit", "value"}, v)
	return j
}
func Type(v string) *simplejson.Json {
	j := simplejson.New()
	j.SetPath([]string{"type", "value"}, v)
	return j
}

type ExistsFilter struct {
	Json *simplejson.Json `json:"exists"`
}

func Exists() *ExistsFilter {
	return &ExistsFilter{simplejson.New()}
}

func (j *ExistsFilter) Field(v string) *ExistsFilter {
	j.Json.Set("field", v)
	return j
}

type MissingFilter struct {
	Json *simplejson.Json `json:"missing"`
}

func Missing() *MissingFilter {
	return &MissingFilter{simplejson.New()}
}

func (j *MissingFilter) Field(v string) *MissingFilter {
	j.Json.Set("field", v)
	return j
}
func (j *MissingFilter) Existence(v bool) *MissingFilter {
	j.Json.Set("existence", v)
	return j
}
func (j *MissingFilter) NullValue(v bool) *MissingFilter {
	j.Json.Set("null_value", v)
	return j
}

type AndFilter struct {
	Json *simplejson.Json `json:"and"`
}

func And() *AndFilter {
	return &AndFilter{simplejson.New()}
}

func (j *AndFilter) And(v ...interface{}) interface{} {
	js := simplejson.New()
	js.Set("and, ", v)
	return js
}
func (j *AndFilter) Filters(v ...interface{}) *AndFilter {
	if len(v) > 1 {
		j.Json.Set("filters", v)
	} else {
		j.Json.Set("filters", v[0])
	}
	return j
}
func (j *AndFilter) Cache(v bool) *AndFilter {
	j.Json.Set("_cache", v)
	return j
}

type NotFilter struct {
	Json *simplejson.Json `json:"not"`
}

func Not() *NotFilter {
	return &NotFilter{simplejson.New()}
}

func (j *NotFilter) Not(v ...interface{}) interface{} {
	js := simplejson.New()
	if len(v) > 1 {
		js.Set("not", v)
	} else {
		js.Set("not", v[0])
	}
	return js
}
func (j *NotFilter) Filters(v ...interface{}) *NotFilter {
	if len(v) > 1 {
		j.Json.Set("filters", v)
	} else {
		j.Json.Set("filters", v[0])
	}
	return j
}
func (j *NotFilter) Cache(v bool) *NotFilter {
	j.Json.Set("_cache", v)
	return j
}

type OrFilter struct {
	Json *simplejson.Json `json:"or"`
}

func Or() *OrFilter {
	return &OrFilter{simplejson.New()}
}

func (j *OrFilter) Or(v ...interface{}) interface{} {
	js := simplejson.New()
	if len(v) > 1 {
		js.Set("or", v)
	} else {
		js.Set("or", v[0])
	}
	return js
}
func (j *OrFilter) Filters(v ...interface{}) *OrFilter {
	if len(v) > 1 {
		j.Json.Set("filters", v)
	} else {
		j.Json.Set("filters", v[0])
	}
	return j
}
func (j *OrFilter) Cache(v bool) *OrFilter {
	j.Json.Set("_cache", v)
	return j
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-geo-bounding-box-filter.html
type GeoBoundingBoxFilter struct {
	Json *simplejson.Json `json:"geo_bounding_box"`
}

func GeoBoundingBox() *GeoBoundingBoxFilter {
	return &GeoBoundingBoxFilter{simplejson.New()}
}

func (j *GeoBoundingBoxFilter) TopLeft(v interface{}) *GeoBoundingBoxFilter {
	j.Json.SetPath([]string{"pin.location", "top_left"}, v)
	return j
}
func (j *GeoBoundingBoxFilter) BottomRight(v interface{}) *GeoBoundingBoxFilter {
	j.Json.SetPath([]string{"pin.location", "bottom_right"}, v)
	return j
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-geo-distance-filter.html
type GeoDistanceFilter struct {
	Json *simplejson.Json `json:"geo_distance"`
}

func GeoDistance() *GeoDistanceFilter {
	return &GeoDistanceFilter{simplejson.New()}
}
func (j *GeoDistanceFilter) DistanceType(v interface{}) *GeoDistanceFilter {
	j.Json.Set("distance_type", v)
	return j
}
func (j *GeoDistanceFilter) OptimizeBbox(v interface{}) *GeoDistanceFilter {
	j.Json.Set("optimize_bbox", v)
	return j
}
func (j *GeoDistanceFilter) Distance(v interface{}) *GeoDistanceFilter {
	j.Json.Set("distance", v)
	return j
}
func (j *GeoDistanceFilter) PinLocation(v interface{}) *GeoDistanceFilter {
	j.Json.Set("pin.location", v)
	return j
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-geo-distance-range-filter.html
type GeoDistanceRangeFilter struct {
	Json *simplejson.Json `json:"geo_distance_range"`
}

func GeoDistanceRange() *GeoDistanceRangeFilter {
	return &GeoDistanceRangeFilter{simplejson.New()}
}

func (j *GeoDistanceRangeFilter) Gte(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("gte", v)
	return j
}
func (j *GeoDistanceRangeFilter) Gt(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("gt", v)
	return j
}

func (j *GeoDistanceRangeFilter) Lte(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("lte", v)
	return j
}

func (j *GeoDistanceRangeFilter) Lt(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("lt", v)
	return j
}
func (j *GeoDistanceRangeFilter) From(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("from", v)
	return j
}

func (j *GeoDistanceRangeFilter) To(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("to", v)
	return j
}
func (j *GeoDistanceRangeFilter) IncludeUpper(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("include_upper", v)
	return j
}

func (j *GeoDistanceRangeFilter) IncludeLower(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("include_lower", v)
	return j
}
func (j *GeoDistanceRangeFilter) PinLocation(v interface{}) *GeoDistanceRangeFilter {
	j.Json.Set("pin.location", v)
	return j
}
func (j *GeoDistanceRangeFilter) Cache(v bool) *GeoDistanceRangeFilter {
	j.Json.Set("_cache", v)
	return j
}

type GeoPolygonFilter struct {
	Json *simplejson.Json `json:"geo_polygon"`
}

func GeoPolygon() *GeoPolygonFilter {
	return &GeoPolygonFilter{simplejson.New()}
}

func (j *GeoPolygonFilter) Points(v ...interface{}) *GeoPolygonFilter {
	j.Json.SetPath([]string{"person.location", "points"}, v)
	return j
}
func (j *GeoPolygonFilter) Cache(v bool) *GeoPolygonFilter {
	j.Json.Set("_cache", v)
	return j
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-geohash-cell-filter.html
type GeohashCellFilter struct {
	Json *simplejson.Json `json:"geohash_cell"`
}

func GeohashCell() *GeohashCellFilter {
	return &GeohashCellFilter{simplejson.New()}
}

func (j *GeohashCellFilter) Lat(v interface{}) *GeohashCellFilter {
	j.Json.SetPath([]string{"pin", "lat"}, v)
	return j
}
func (j *GeohashCellFilter) Lon(v interface{}) *GeohashCellFilter {
	j.Json.SetPath([]string{"pin", "lon"}, v)
	return j
}
func (j *GeohashCellFilter) Precision(v interface{}) *GeohashCellFilter {
	j.Json.Set("precision", v)
	return j
}
func (j *GeohashCellFilter) Neighbors(v bool) *GeohashCellFilter {
	j.Json.Set("neighbors", v)
	return j
}
func (j *GeohashCellFilter) Cache(v bool) *GeohashCellFilter {
	j.Json.Set("_cache", v)
	return j
}

type ScriptFilter struct {
	Json *simplejson.Json `json:"script"`
}

func Script() *ScriptFilter {
	return &ScriptFilter{simplejson.New()}
}
func (j *ScriptFilter) Script(v interface{}) *ScriptFilter {
	j.Json.Set("script", v)
	return j
}
func (j *ScriptFilter) Params(k string, v interface{}) *ScriptFilter {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *ScriptFilter) Cache(v bool) *ScriptFilter {
	j.Json.Set("_cache", v)
	return j
}
func (j *ScriptFilter) Map() interface{} {
	m, err := j.Json.Map()
	if err != nil {
		return nil
	}
	return m
}
