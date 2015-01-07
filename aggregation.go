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

type SearchAggregation struct {
	Json *simplejson.Json `json:"aggs"`
}

func Aggregation() *SearchAggregation {
	return &SearchAggregation{simplejson.New()}
}

func (j *SearchAggregation) Aggregation(key string, val ...*SearchAggregation) *SearchAggregation {
	if js, ok := j.Json.CheckGet(key); ok {
		for _, v := range val {
			js.Set("aggs", v.Json)
		}
	} else {
		for _, v := range val {
			j.Json.Set(key, v)
		}
	}
	return j
}
func (j *SearchAggregation) Encode() []byte {
	return encode(j)
}

type MinAggregation struct {
	Json   *simplejson.Json `json:"min"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Min(field string) *MinAggregation {
	minj := &MinAggregation{simplejson.New(), j}
	j.Json.Set(field, minj)
	return minj
}
func (j *MinAggregation) Field(field string) *MinAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *MinAggregation) Script(v string) *MinAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *MinAggregation) Params(k string, v interface{}) *MinAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}

func (j *MinAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type MaxAggregation struct {
	Json   *simplejson.Json `json:"max"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Max(field string) *MaxAggregation {
	js := &MaxAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *MaxAggregation) Field(field string) *MaxAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *MaxAggregation) Script(v string) *MaxAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *MaxAggregation) Params(k string, v interface{}) *MaxAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *MaxAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type SumAggregation struct {
	Json   *simplejson.Json `json:"sum"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Sum(field string) *SumAggregation {
	js := &SumAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *SumAggregation) Field(field string) *SumAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *SumAggregation) Script(v string) *SumAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *SumAggregation) Params(k string, v interface{}) *SumAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *SumAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type AvgAggregation struct {
	Json   *simplejson.Json `json:"avg"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Avg(field string) *AvgAggregation {
	js := &AvgAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *AvgAggregation) Field(field string) *AvgAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *AvgAggregation) Script(v string) *AvgAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *AvgAggregation) Params(k string, v interface{}) *AvgAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}

func (j *AvgAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type StatsAggregation struct {
	Json   *simplejson.Json `json:"stats"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Stats(field string) *StatsAggregation {
	js := &StatsAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *StatsAggregation) Field(field string) *StatsAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *StatsAggregation) Script(v string) *StatsAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *StatsAggregation) Params(k string, v interface{}) *StatsAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *StatsAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type ExtendedStatsAggregation struct {
	Json   *simplejson.Json `json:"extended_stats"`
	parent *SearchAggregation
}

func (j *SearchAggregation) ExtendedStats(field string) *ExtendedStatsAggregation {
	js := &ExtendedStatsAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *ExtendedStatsAggregation) Field(field string) *ExtendedStatsAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *ExtendedStatsAggregation) Script(v string) *ExtendedStatsAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *ExtendedStatsAggregation) Params(k string, v interface{}) *ExtendedStatsAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *ExtendedStatsAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type ValueCountAggregation struct {
	Json   *simplejson.Json `json:"value_count"`
	parent *SearchAggregation
}

func (j *SearchAggregation) ValueCount(field string) *ValueCountAggregation {
	js := &ValueCountAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *ValueCountAggregation) Field(field string) *ValueCountAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *ValueCountAggregation) Script(v string) *ValueCountAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *ValueCountAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type PercentilesAggregation struct {
	Json   *simplejson.Json `json:"percentiles"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Percentiles(field string) *PercentilesAggregation {
	js := &PercentilesAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *PercentilesAggregation) Field(field string) *PercentilesAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *PercentilesAggregation) Script(v string) *PercentilesAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *PercentilesAggregation) Percents(v ...float64) *PercentilesAggregation {
	j.Json.Set("percents", v)
	return j
}
func (j *PercentilesAggregation) Params(k string, v interface{}) *PercentilesAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *PercentilesAggregation) Compression(v int) *PercentilesAggregation {
	j.Json.Set("compression", v)
	return j
}
func (j *PercentilesAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type PercentileRanksAggregation struct {
	Json   *simplejson.Json `json:"percentile_ranks"`
	parent *SearchAggregation
}

func (j *SearchAggregation) PercentileRanks(field string) *PercentileRanksAggregation {
	js := &PercentileRanksAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *PercentileRanksAggregation) Field(field string) *PercentileRanksAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *PercentileRanksAggregation) Script(v string) *PercentileRanksAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *PercentileRanksAggregation) Value(v ...int) *PercentileRanksAggregation {
	j.Json.Set("value", v)
	return j
}
func (j *PercentileRanksAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type CardinalityAggregation struct {
	Json   *simplejson.Json `json:"cardinality"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Cardinality(field string) *CardinalityAggregation {
	js := &CardinalityAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *CardinalityAggregation) Field(field string) *CardinalityAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *CardinalityAggregation) Script(v string) *CardinalityAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *CardinalityAggregation) Rehash(v bool) *CardinalityAggregation {
	j.Json.Set("rehash", v)
	return j
}
func (j *CardinalityAggregation) PrecisionThreshold(v int) *CardinalityAggregation {
	j.Json.Set("precision_threshold", v)
	return j
}
func (j *CardinalityAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type GeoBoundsAggregation struct {
	Json   *simplejson.Json `json:"geo_bounds"`
	parent *SearchAggregation
}

func (j *SearchAggregation) GeoBounds(field string) *GeoBoundsAggregation {
	js := &GeoBoundsAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *GeoBoundsAggregation) Field(field string) *GeoBoundsAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *GeoBoundsAggregation) WrapLongitude(v string) *GeoBoundsAggregation {
	j.Json.Set("wrap_longitude", v)
	return j
}
func (j *GeoBoundsAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type ScriptedMetricAggregation struct {
	Json   *simplejson.Json `json:"scripted_metric"`
	parent *SearchAggregation
}

func (j *SearchAggregation) ScriptedMetric(field string) *ScriptedMetricAggregation {
	js := &ScriptedMetricAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *ScriptedMetricAggregation) InitScript(v string) *ScriptedMetricAggregation {
	j.Json.Set("init_script", v)
	return j
}
func (j *ScriptedMetricAggregation) MapScript(v string) *ScriptedMetricAggregation {
	j.Json.Set("map_script", v)
	return j
}

func (j *ScriptedMetricAggregation) CombineScript(v string) *ScriptedMetricAggregation {
	j.Json.Set("combine_script", v)
	return j
}

func (j *ScriptedMetricAggregation) ReduceScript(v string) *ScriptedMetricAggregation {
	j.Json.Set("reduce_script", v)
	return j
}
func (j *ScriptedMetricAggregation) Aggs() *SearchAggregation {
	return j.parent
}

func (j *SearchAggregation) Global(field string) *SearchAggregation {
	j.Json.SetPath([]string{field, "global"}, struct{}{})
	return j
}

func (j *SearchAggregation) Filter(field string, v interface{}) *SearchAggregation {
	j.Json.SetPath([]string{field, "filter"}, v)
	return j
}

func (j *SearchAggregation) Filters(field string, v interface{}) *SearchAggregation {
	j.Json.SetPath([]string{field, "filters"}, v)
	return j
}

type MissingAggregation struct {
	Json   *simplejson.Json `json:"missing"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Missing(field string) *MissingAggregation {
	js := &MissingAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *MissingAggregation) Field(field string) *MissingAggregation {
	j.Json.Set("field", field)
	return j
}
func (j *MissingAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type NestedAggregation struct {
	Json   *simplejson.Json `json:"nested"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Nested(field string) *NestedAggregation {
	js := &NestedAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *NestedAggregation) Path(v string) *NestedAggregation {
	j.Json.Set("path", v)
	return j
}
func (j *NestedAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type ChildrenAggregation struct {
	Json   *simplejson.Json `json:"children"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Children(field string) *ChildrenAggregation {
	js := &ChildrenAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *ChildrenAggregation) Type(v string) *ChildrenAggregation {
	j.Json.Set("type", v)
	return j
}
func (j *ChildrenAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type TermsAggregation struct {
	Json   *simplejson.Json `json:"terms"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Terms(field string) *TermsAggregation {
	js := &TermsAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *TermsAggregation) Field(v string) *TermsAggregation {
	j.Json.Set("field", v)
	return j
}
func (j *TermsAggregation) MinDocCount(v int) *TermsAggregation {
	j.Json.Set("min_doc_count", v)
	return j
}
func (j *TermsAggregation) Size(v int) *TermsAggregation {
	j.Json.Set("size", v)
	return j
}
func (j *TermsAggregation) Script(v string) *TermsAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *TermsAggregation) Order(v map[string]string) *TermsAggregation {
	j.Json.Set("order", v)
	return j
}
func (j *TermsAggregation) Include(v interface{}) *TermsAggregation {
	j.Json.Set("include", v)
	return j
}
func (j *TermsAggregation) Exclude(v interface{}) *TermsAggregation {
	j.Json.Set("exclude", v)
	return j
}
func (j *TermsAggregation) CollectMode(v string) *TermsAggregation {
	j.Json.Set("collect_mode", v)
	return j
}
func (j *TermsAggregation) ExecutionHint(v string) *TermsAggregation {
	j.Json.Set("execution_hint", v)
	return j
}
func (j *TermsAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type RangeAggregation struct {
	Json   *simplejson.Json `json:"range"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Range(field string) *RangeAggregation {
	js := &RangeAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *RangeAggregation) Field(v string) *RangeAggregation {
	j.Json.Set("field", v)
	return j
}
func (j *RangeAggregation) Keyed(v bool) *RangeAggregation {
	j.Json.Set("keyed", v)
	return j
}
func (j *RangeAggregation) Params(k string, v interface{}) *RangeAggregation {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *RangeAggregation) Script(v string) *RangeAggregation {
	j.Json.Set("script", v)
	return j
}
func (j *RangeAggregation) Ranges(v ...map[string]interface{}) *RangeAggregation {
	j.Json.Set("ranges", v)
	return j
}
func (j *RangeAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type DateRangeAggregation struct {
	Json   *simplejson.Json `json:"date_range"`
	parent *SearchAggregation
}

func (j *SearchAggregation) DateRange(field string) *DateRangeAggregation {
	js := &DateRangeAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *DateRangeAggregation) Field(v string) *DateRangeAggregation {
	j.Json.Set("field", v)
	return j
}
func (j *DateRangeAggregation) Format(v string) *DateRangeAggregation {
	j.Json.Set("format", v)
	return j
}
func (j *DateRangeAggregation) Ranges(v ...map[string]interface{}) *DateRangeAggregation {
	j.Json.Set("ranges", v)
	return j
}
func (j *DateRangeAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type IpRangeAggregation struct {
	Json   *simplejson.Json `json:"ip_range"`
	parent *SearchAggregation
}

func (j *SearchAggregation) IpRange(field string) *IpRangeAggregation {
	js := &IpRangeAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *IpRangeAggregation) Field(v string) *IpRangeAggregation {
	j.Json.Set("field", v)
	return j
}
func (j *IpRangeAggregation) Ranges(v ...map[string]interface{}) *IpRangeAggregation {
	j.Json.Set("ranges", v)
	return j
}
func (j *IpRangeAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type HistogramAggregation struct {
	Json   *simplejson.Json `json:"histogram"`
	parent *SearchAggregation
}

func (j *SearchAggregation) Histogram(field string) *HistogramAggregation {
	js := &HistogramAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *HistogramAggregation) Field(v string) *HistogramAggregation {
	j.Json.Set("field", v)
	return j
}
func (j *HistogramAggregation) MinDocCount(v int) *HistogramAggregation {
	j.Json.Set("min_doc_count", v)
	return j
}
func (j *HistogramAggregation) Interval(v int) *HistogramAggregation {
	j.Json.Set("interval", v)
	return j
}
func (j *HistogramAggregation) Keyed(v bool) *HistogramAggregation {
	j.Json.Set("keyed", v)
	return j
}
func (j *HistogramAggregation) Order(v map[string]string) *HistogramAggregation {
	j.Json.Set("order", v)
	return j
}
func (j *HistogramAggregation) Aggs() *SearchAggregation {
	return j.parent
}

type DateHistogramAggregation struct {
	Json   *simplejson.Json `json:"date_histogram"`
	parent *SearchAggregation
}

func (j *SearchAggregation) DateHistogram(field string) *DateHistogramAggregation {
	js := &DateHistogramAggregation{simplejson.New(), j}
	j.Json.Set(field, js)
	return js
}
func (j *DateHistogramAggregation) Field(v string) *DateHistogramAggregation {
	j.Json.Set("field", v)
	return j
}
func (j *DateHistogramAggregation) Interval(v string) *DateHistogramAggregation {
	j.Json.Set("interval", v)
	return j
}
func (j *DateHistogramAggregation) Format(v string) *DateHistogramAggregation {
	j.Json.Set("format", v)
	return j
}
func (j *DateHistogramAggregation) Aggs() *SearchAggregation {
	return j.parent
}
func (j *SearchAggregation) Term(field string, v interface{}) *SearchAggregation {
	j.Json.SetPath([]string{field, "term", "field"}, v)
	return j
}
