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

type SearchFacet struct {
	Json *simplejson.Json `json:"facets"`
}

func Facets() *SearchFacet {
	return &SearchFacet{simplejson.New()}
}

func (j *SearchFacet) Facets(key string, val ...*SearchFacet) *SearchFacet {
	if js, ok := j.Json.CheckGet(key); ok {
		for _, v := range val {
			js.Set("facets", v.Json)
		}
	} else {
		for _, v := range val {
			j.Json.Set(key, v)
		}
	}
	return j
}

// func Facets(v interface{}) *simplejson.Json {
// 	j := simplejson.New()
// 	j.Set("facets", v)
// 	return j
// }
func From(v int) *simplejson.Json {
	j := simplejson.New()
	j.Set("from", v)
	return j
}
func To(v int) *simplejson.Json {
	j := simplejson.New()
	j.Set("to", v)
	return j
}
func FromTo(from, to int) *simplejson.Json {
	j := simplejson.New()
	j.Set("from", from)
	j.Set("to", to)
	return j
}

type TermsFacets struct {
	Json   *simplejson.Json `json:"terms"`
	parent *SearchFacet
}

func (j *SearchFacet) Terms(field string) *TermsFacets {
	js := &TermsFacets{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *TermsFacets) Field(v ...string) *TermsFacets {
	if len(v) > 1 {
		j.Json.Set("field", v)
	} else {
		j.Json.Set("field", v[0])
	}
	return j
}
func (j *TermsFacets) Order(v string) *TermsFacets {
	j.Json.Set("order", v)
	return j
}
func (j *TermsFacets) AllTerms(v bool) *TermsFacets {
	j.Json.Set("all_terms", v)
	return j
}
func (j *TermsFacets) Size(v int) *TermsFacets {
	j.Json.Set("size", v)
	return j
}
func (j *TermsFacets) Exclude(v ...string) *TermsFacets {
	j.Json.Set("exclude", v)
	return j
}
func (j *TermsFacets) Regex(v string) *TermsFacets {
	j.Json.Set("regex", v)
	return j
}
func (j *TermsFacets) RegexFlags(v string) *TermsFacets {
	j.Json.Set("regex_flags", v)
	return j
}
func (j *TermsFacets) Script(v string) *TermsFacets {
	j.Json.Set("script", v)
	return j
}
func (j *TermsFacets) ScriptField(v string) *TermsFacets {
	j.Json.Set("script_field", v)
	return j
}
func (j *TermsFacets) Index(v string) *TermsFacets {
	j.Json.Set("_index", v)
	return j
}
func (j *TermsFacets) Facet() *SearchFacet {
	return j.parent
}

type RangeFacets struct {
	Json   *simplejson.Json `json:"range"`
	parent *SearchFacet
}

func (j *SearchFacet) Range(field string) *RangeFacets {
	js := &RangeFacets{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *RangeFacets) Field(v string) *RangeFacets {
	j.Json.Set("field", v)
	return j
}
func (j *RangeFacets) KeyField(v string) *RangeFacets {
	j.Json.Set("key_field", v)
	return j
}
func (j *RangeFacets) ValueField(v string) *RangeFacets {
	j.Json.Set("value_field", v)
	return j
}
func (j *RangeFacets) KeyScript(v string) *RangeFacets {
	j.Json.Set("key_script", v)
	return j
}
func (j *RangeFacets) ValueScript(v string) *RangeFacets {
	j.Json.Set("value_script", v)
	return j
}
func (j *RangeFacets) Ranges(v ...interface{}) *RangeFacets {
	j.Json.Set("ranges", v)
	return j
}
func (j *RangeFacets) Facet() *SearchFacet {
	return j.parent
}

type HistogramFacets struct {
	Json   *simplejson.Json `json:"histogram"`
	parent *SearchFacet
}

func (j *SearchFacet) Histogram(field string) *HistogramFacets {
	js := &HistogramFacets{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *HistogramFacets) Field(v string) *HistogramFacets {
	j.Json.Set("field", v)
	return j
}
func (j *HistogramFacets) TimeInterval(v string) *HistogramFacets {
	j.Json.Set("time_interval", v)
	return j
}
func (j *HistogramFacets) Interval(v int) *HistogramFacets {
	j.Json.Set("interval", v)
	return j
}
func (j *HistogramFacets) KeyField(v string) *HistogramFacets {
	j.Json.Set("key_field", v)
	return j
}
func (j *HistogramFacets) ValueField(v string) *HistogramFacets {
	j.Json.Set("value_field", v)
	return j
}
func (j *HistogramFacets) KeyScript(v string) *HistogramFacets {
	j.Json.Set("key_script", v)
	return j
}
func (j *HistogramFacets) ValueScript(v string) *HistogramFacets {
	j.Json.Set("value_script", v)
	return j
}
func (j *HistogramFacets) Params(v map[string]interface{}) *HistogramFacets {
	j.Json.Set("params", v)
	return j
}
func (j *HistogramFacets) Facet() *SearchFacet {
	return j.parent
}

type DateHistogramFacets struct {
	Json   *simplejson.Json `json:"date_histogram"`
	parent *SearchFacet
}

func (j *SearchFacet) DateHistogram(field string) *DateHistogramFacets {
	js := &DateHistogramFacets{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *DateHistogramFacets) Field(v string) *DateHistogramFacets {
	j.Json.Set("field", v)
	return j
}
func (j *DateHistogramFacets) Interval(v string) *DateHistogramFacets {
	j.Json.Set("interval", v)
	return j
}
func (j *DateHistogramFacets) KeyField(v string) *DateHistogramFacets {
	j.Json.Set("key_field", v)
	return j
}
func (j *DateHistogramFacets) ValueField(v string) *DateHistogramFacets {
	j.Json.Set("value_field", v)
	return j
}
func (j *DateHistogramFacets) KeyScript(v string) *DateHistogramFacets {
	j.Json.Set("key_script", v)
	return j
}
func (j *DateHistogramFacets) ValueScript(v string) *DateHistogramFacets {
	j.Json.Set("value_script", v)
	return j
}
func (j *DateHistogramFacets) TimeZone(v string) *DateHistogramFacets {
	j.Json.Set("time_zone", v)
	return j
}
func (j *DateHistogramFacets) PreZone(v string) *DateHistogramFacets {
	j.Json.Set("pre_zone", v)
	return j
}
func (j *DateHistogramFacets) PostZone(v string) *DateHistogramFacets {
	j.Json.Set("post_zone", v)
	return j
}
func (j *DateHistogramFacets) PreOffset(v string) *DateHistogramFacets {
	j.Json.Set("pre_offset", v)
	return j
}
func (j *DateHistogramFacets) PostOffset(v string) *DateHistogramFacets {
	j.Json.Set("post_offset", v)
	return j
}
func (j *DateHistogramFacets) Factor(v float64) *DateHistogramFacets {
	j.Json.Set("factor", v)
	return j
}
func (j *DateHistogramFacets) PreZoneAdjustLargeInterval(v bool) *DateHistogramFacets {
	j.Json.Set("pre_zone_adjust_large_interval", v)
	return j
}
func (j *DateHistogramFacets) Facet() *SearchFacet {
	return j.parent
}

func (j *SearchFacet) Filter(field string, v interface{}) *SearchFacet {
	j.Json.SetPath([]string{field, "filter"}, v)
	return j
}

func (j *SearchFacet) Query(field string, v interface{}) *SearchFacet {
	j.Json.SetPath([]string{field, "query"}, v)
	return j
}

type StatisticalFacets struct {
	Json   *simplejson.Json `json:"statistical"`
	parent *SearchFacet
}

func (j *SearchFacet) Statistical(field string) *StatisticalFacets {
	js := &StatisticalFacets{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *StatisticalFacets) Field(v ...string) *StatisticalFacets {
	if len(v) > 1 {
		j.Json.Set("field", v)
	} else {
		j.Json.Set("field", v[0])
	}
	return j
}
func (j *StatisticalFacets) script(v string) *StatisticalFacets {
	j.Json.Set("script", v)
	return j
}
func (j *StatisticalFacets) Params(v map[string]interface{}) *StatisticalFacets {
	j.Json.Set("params", v)
	return j
}
func (j *StatisticalFacets) Facet() *SearchFacet {
	return j.parent
}

type TermsStatsFacets struct {
	Json   *simplejson.Json `json:"terms_stats"`
	parent *SearchFacet
}

func (j *SearchFacet) TermsStats(field string) *TermsStatsFacets {
	js := &TermsStatsFacets{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}

func (j *TermsStatsFacets) KeyField(v string) *TermsStatsFacets {
	j.Json.Set("key_field", v)
	return j
}
func (j *TermsStatsFacets) ValueField(v string) *TermsStatsFacets {
	j.Json.Set("value_field", v)
	return j
}

func (j *TermsStatsFacets) Order(v string) *TermsStatsFacets {
	j.Json.Set("order", v)
	return j
}
func (j *TermsStatsFacets) Lang(v string) *TermsStatsFacets {
	j.Json.Set("lang", v)
	return j
}
func (j *TermsStatsFacets) ShardSize(v int) *TermsStatsFacets {
	j.Json.Set("shard_size", v)
	return j
}
func (j *TermsStatsFacets) Size(v int) *TermsStatsFacets {
	j.Json.Set("size", v)
	return j
}
func (j *TermsStatsFacets) Script(v string) *TermsStatsFacets {
	j.Json.Set("script", v)
	return j
}
func (j *TermsStatsFacets) Params(v map[string]interface{}) *TermsStatsFacets {
	j.Json.Set("params", v)
	return j
}
func (j *TermsStatsFacets) Facet() *SearchFacet {
	return j.parent
}

type GeoDistanceFacets struct {
	Json   *simplejson.Json `json:"terms_stats"`
	parent *SearchFacet
}

func (j *SearchFacet) GeoDistance(field string) *GeoDistanceFacets {
	js := &GeoDistanceFacets{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}

func (j *GeoDistanceFacets) PinLocation(v interface{}) *GeoDistanceFacets {
	j.Json.Set("pin.location", v)
	return j
}
func (j *GeoDistanceFacets) Ranges(v ...interface{}) *GeoDistanceFacets {
	j.Json.Set("ranges", v)
	return j
}

func (j *GeoDistanceFacets) ValueField(v string) *GeoDistanceFacets {
	j.Json.Set("value_field", v)
	return j
}

func (j *GeoDistanceFacets) ValueScript(v string) *GeoDistanceFacets {
	j.Json.Set("value_script", v)
	return j
}
func (j *GeoDistanceFacets) Params(v map[string]interface{}) *GeoDistanceFacets {
	j.Json.Set("params", v)
	return j
}
func (j *GeoDistanceFacets) Facet() *SearchFacet {
	return j.parent
}
