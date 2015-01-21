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

func Query(v interface{}) *simplejson.Json {
	j := simplejson.New()
	j.Set("query", v)
	return j
}

func MatchAll(v interface{}) *simplejson.Json {
	j := simplejson.New()
	if v == nil {
		j.Set("match_all", struct{}{})
	} else {
		j.Set("match_all", v)
	}
	return j
}
func Boost(v interface{}) interface{} {
	j := simplejson.New()
	j.Set("boost", v)
	return j
}

func Array(v ...interface{}) []interface{} {
	return v
}

type DocDsl struct {
	*simplejson.Json
}

func Doc() *DocDsl {
	return &DocDsl{simplejson.New()}
}
func (d *DocDsl) Index(v string) *DocDsl {
	d.Set("_index", v)
	return d
}
func (d *DocDsl) Type(v string) *DocDsl {
	d.Set("_type", v)
	return d
}
func (d *DocDsl) Id(v string) *DocDsl {
	d.Set("_id", v)
	return d
}
func (d *DocDsl) Fileds(v ...string) *DocDsl {
	d.Set("fields", v)
	return d
}
func (d *DocDsl) Source(v interface{}) *DocDsl {
	d.Set("_source", v)
	return d
}
func (d *DocDsl) Routing(v string) *DocDsl {
	d.Set("_routing", v)
	return d
}
func (d *DocDsl) Encode() []byte {
	b, err := d.Bytes()
	if err != nil {
		panic(err)
	}
	return b
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
type RangeQuery struct {
	field string
	Json  *simplejson.Json `json:"range"`
}

func Range(field string) *RangeQuery {
	return &RangeQuery{field, simplejson.New()}
}
func (j *RangeQuery) Gte(v interface{}) *RangeQuery {
	j.Json.SetPath([]string{j.field, "gte"}, v)
	return j
}
func (j *RangeQuery) Gt(v interface{}) *RangeQuery {
	j.Json.SetPath([]string{j.field, "gt"}, v)
	return j
}

func (j *RangeQuery) Lte(v interface{}) *RangeQuery {
	j.Json.SetPath([]string{j.field, "lte"}, v)
	return j
}

func (j *RangeQuery) Lt(v interface{}) *RangeQuery {
	j.Json.SetPath([]string{j.field, "lt"}, v)
	return j
}
func (j *RangeQuery) From(v interface{}) *RangeQuery {
	j.Json.SetPath([]string{j.field, "from"}, v)
	return j
}

func (j *RangeQuery) To(v interface{}) *RangeQuery {
	j.Json.SetPath([]string{j.field, "to"}, v)
	return j
}
func (j *RangeQuery) Boost(v float64) *RangeQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *RangeQuery) TimeZone(v interface{}) *RangeQuery {
	j.Json.SetPath([]string{j.field, "time_zone"}, v)
	return j
}
func (j *RangeQuery) Cache(v bool) *RangeQuery {
	j.Json.Set("_cache", v)
	return j
}
func (j *RangeQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type TermQuery struct {
	field string
	Json  *simplejson.Json `json:"term"`
}

func Term(field string) *TermQuery {
	return &TermQuery{field, simplejson.New()}
}
func (j *TermQuery) Val(v interface{}) *TermQuery {
	j.Json.Set(j.field, v)
	return j
}
func (j *TermQuery) Value(v interface{}) *TermQuery {
	j.Json.SetPath([]string{j.field, "value"}, v)
	return j
}
func (j *TermQuery) Term(v interface{}) *TermQuery {
	j.Json.SetPath([]string{j.field, "term"}, v)
	return j
}
func (j *TermQuery) Cache(v bool) *TermQuery {
	j.Json.Set("_cache", v)
	return j
}
func (j *TermQuery) Boost(v interface{}) *TermQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *TermQuery) Encode() []byte {
	return encode(j)
}

// v is map[string]string
// map["query"]="this is a test"
// map["operator"]="and"
//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type MatchQuery struct {
	field string
	Json  *simplejson.Json `json:"match"`
}

func Match(field string) *MatchQuery {
	return &MatchQuery{field, simplejson.New()}
}

func (j *MatchQuery) Value(v interface{}) *MatchQuery {
	j.Json.SetPath([]string{j.field}, v)
	return j
}
func (j *MatchQuery) Query(v interface{}) *MatchQuery {
	j.Json.SetPath([]string{j.field, "query"}, v)
	return j
}
func (j *MatchQuery) Operator(operator string) *MatchQuery {
	j.Json.SetPath([]string{j.field, "operator"}, operator)
	return j
}

//none or all
func (j *MatchQuery) ZeroTermsQuery(p string) *MatchQuery {
	j.Json.SetPath([]string{j.field, "zero_terms_query"}, p)
	return j
}

//range [0..1)
func (j *MatchQuery) CutOffFrequency(p float64) *MatchQuery {
	j.Json.SetPath([]string{j.field, "cutoff_frequency"}, p)
	return j
}
func (j *MatchQuery) Slop(p int) *MatchQuery {
	j.Json.SetPath([]string{j.field, "slop"}, p)
	return j
}

//phrase, phrase_prefix
func (j *MatchQuery) Type(p string) *MatchQuery {
	j.Json.SetPath([]string{j.field, "type"}, p)
	return j
}
func (j *MatchQuery) Encode() []byte {
	return encode(j)
}

type MatchPhraseQuery struct {
	field string
	Json  *simplejson.Json `json:"match_phrase"`
}

func MatchPhrase(field string) *MatchPhraseQuery {
	return &MatchPhraseQuery{field, simplejson.New()}
}

func (j *MatchPhraseQuery) Value(v interface{}) *MatchPhraseQuery {
	j.Json.SetPath([]string{j.field}, v)
	return j
}
func (j *MatchPhraseQuery) Query(v interface{}) *MatchPhraseQuery {
	j.Json.SetPath([]string{j.field, "query"}, v)
	return j
}
func (j *MatchPhraseQuery) Analyzer(v interface{}) *MatchPhraseQuery {
	j.Json.SetPath([]string{j.field, "analyzer"}, v)
	return j
}
func (j *MatchPhraseQuery) PhraseSlop(v int) *MatchPhraseQuery {
	j.Json.SetPath([]string{j.field, "phrase_slop"}, v)
	return j
}
func (j *MatchPhraseQuery) Boost(v float64) *MatchPhraseQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *MatchPhraseQuery) Encode() []byte {
	return encode(j)
}

type MatchPhrasePrefixQuery struct {
	field string
	Json  *simplejson.Json `json:"match_phrase_prefix"`
}

func MatchPhrasePrefix(field string) *MatchPhrasePrefixQuery {
	return &MatchPhrasePrefixQuery{field, simplejson.New()}
}
func (j *MatchPhrasePrefixQuery) Value(v interface{}) *MatchPhrasePrefixQuery {
	j.Json.SetPath([]string{j.field}, v)
	return j
}
func (j *MatchPhrasePrefixQuery) Query(v interface{}) *MatchPhrasePrefixQuery {
	j.Json.SetPath([]string{j.field, "query"}, v)
	return j
}
func (j *MatchPhrasePrefixQuery) MaxExpansions(v interface{}) *MatchPhrasePrefixQuery {
	j.Json.SetPath([]string{j.field, "max_expansions"}, v)
	return j
}
func (j *MatchPhrasePrefixQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html
type MultiMatchQuery struct {
	Json *simplejson.Json `json:"multi_match"`
}

func MultiMatch() *MultiMatchQuery {
	j := simplejson.New()
	return &MultiMatchQuery{j}
}
func (j *MultiMatchQuery) Query(query string) *MultiMatchQuery {
	j.Json.Set("query", query)
	return j
}

//type parameter best_fields,most_fields,cross_fields,phrase,phrase_prefix
//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html#multi-match-types
func (j *MultiMatchQuery) Type(_type string) *MultiMatchQuery {
	j.Json.Set("type", _type)
	return j
}
func (j *MultiMatchQuery) Operator(operator string) *MultiMatchQuery {
	j.Json.Set("operator", operator)
	return j
}
func (j *MultiMatchQuery) Field(field ...string) *MultiMatchQuery {
	if len(field) > 1 {
		j.Json.Set("fields", field)
	} else {
		j.Json.Set("fields", field[0])
	}
	return j
}

func (j *MultiMatchQuery) Encode() []byte {
	return encode(j)
}

type BooleanQuery struct {
	Json *simplejson.Json `json:"bool"`
}

func Bool() *BooleanQuery {
	return &BooleanQuery{simplejson.New()}
}
func (j *BooleanQuery) Must(v ...interface{}) *BooleanQuery {
	if len(v) > 1 {
		j.Json.Set("must", v)
	} else {
		j.Json.Set("must", v[0])
	}
	return j
}
func (j *BooleanQuery) MustNot(v ...interface{}) *BooleanQuery {
	if len(v) > 1 {
		j.Json.Set("must_not", v)
	} else {
		j.Json.Set("must_not", v[0])
	}
	return j
}
func (j *BooleanQuery) Should(v ...interface{}) *BooleanQuery {
	if len(v) > 1 {
		j.Json.Set("should", v)
	} else {
		j.Json.Set("should", v[0])
	}
	return j
}
func (j *BooleanQuery) MinimumShouldMatch(v interface{}) *BooleanQuery {
	j.Json.Set("minimum_should_match", v)
	return j
}

//float64
func (j *BooleanQuery) Boost(v float64) *BooleanQuery {
	j.Json.Set("boost", v)
	return j
}

func (j *BooleanQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-boosting-query.html
type BoostingQuery struct {
	Json *simplejson.Json `json:"boosting"`
}

func Boosting() *BoostingQuery {
	return &BoostingQuery{simplejson.New()}
}
func (j *BoostingQuery) Positive(v interface{}) *BoostingQuery {
	j.Json.Set("positive", v)
	return j
}
func (j *BoostingQuery) Negative(v interface{}) *BoostingQuery {
	j.Json.Set("negative", v)
	return j
}
func (j *BoostingQuery) Boost(v float64) *BoostingQuery {
	j.Json.Set("negative_boost", v)
	return j
}

func (j *BoostingQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-common-terms-query.html
type CommonQuery struct {
	Json *simplejson.Json `json:"common"`
}

func Common() *CommonQuery {
	return &CommonQuery{simplejson.New()}
}
func (j *CommonQuery) Query(v interface{}) *CommonQuery {
	j.Json.SetPath([]string{"body", "query"}, v)
	return j
}
func (j *CommonQuery) Frequency(v interface{}) *CommonQuery {
	j.Json.SetPath([]string{"body", "cutoff_frequency"}, v)
	return j
}
func (j *CommonQuery) Operator(v interface{}) *CommonQuery {
	j.Json.SetPath([]string{"body", "low_freq_operator"}, v)
	return j
}

func (j *CommonQuery) MinimumShouldMatch(v interface{}) *CommonQuery {
	j.Json.SetPath([]string{"body", "minimum_should_match"}, v)
	return j
}

func (j *CommonQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-constant-score-query.html
type ConstantQuery struct {
	Json *simplejson.Json `json:"constant_score"`
}

func Constant() *ConstantQuery {
	return &ConstantQuery{simplejson.New()}
}

func (j *ConstantQuery) Boost(v float64) *ConstantQuery {
	j.Json.Set("boost", v)
	return j
}
func (j *ConstantQuery) Filter(v interface{}) *ConstantQuery {
	j.Json.Set("filter", v)
	return j
}
func (j *ConstantQuery) Query(v interface{}) *ConstantQuery {
	j.Json.Set("query", v)
	return j
}

func (j *ConstantQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-dis-max-query.html
type DisjunctionMaxQuery struct {
	Json *simplejson.Json `json:"dis_max"`
}

func DisjunctionMax() *DisjunctionMaxQuery {
	return &DisjunctionMaxQuery{simplejson.New()}
}
func (j *DisjunctionMaxQuery) Query(v ...interface{}) *DisjunctionMaxQuery {
	j.Json.Set("queries", v)
	return j
}

func (j *DisjunctionMaxQuery) Boost(v float64) *DisjunctionMaxQuery {
	j.Json.Set("boost", v)
	return j
}
func (j *DisjunctionMaxQuery) TieBreaker(v float64) *DisjunctionMaxQuery {
	j.Json.Set("tie_breaker", v)
	return j
}

func (j *DisjunctionMaxQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-flt-query.html
type FuzzyLikeThisQuery struct {
	Json *simplejson.Json `json:"fuzzy_like_this"`
}

func Flt() *FuzzyLikeThisQuery {
	return &FuzzyLikeThisQuery{simplejson.New()}
}

func (j *FuzzyLikeThisQuery) Fields(v ...interface{}) *FuzzyLikeThisQuery {
	if len(v) > 1 {
		j.Json.Set("fields", v)
	} else {
		j.Json.Set("fields", v[0])
	}
	return j
}
func (j *FuzzyLikeThisQuery) LikeText(v interface{}) *FuzzyLikeThisQuery {
	j.Json.Set("like_text", v)
	return j
}
func (j *FuzzyLikeThisQuery) MaxQueryTerms(v int) *FuzzyLikeThisQuery {
	j.Json.Set("max_query_terms", v)
	return j
}
func (j *FuzzyLikeThisQuery) IgnoreTf(v bool) *FuzzyLikeThisQuery {
	j.Json.Set("ignore_tf", v)
	return j
}
func (j *FuzzyLikeThisQuery) Boost(v float64) *FuzzyLikeThisQuery {
	j.Json.Set("boost", v)
	return j
}
func (j *FuzzyLikeThisQuery) Fuzziness(v float64) *FuzzyLikeThisQuery {
	j.Json.Set("fuzziness", v)
	return j
}
func (j *FuzzyLikeThisQuery) PrefixLength(v int) *FuzzyLikeThisQuery {
	j.Json.Set("prefix_length", v)
	return j
}
func (j *FuzzyLikeThisQuery) Analyzer(v interface{}) *FuzzyLikeThisQuery {
	j.Json.Set("analyzer", v)
	return j
}

func (j *FuzzyLikeThisQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-flt-field-query.html
type FuzzyLikeThisFiedQuery struct {
	field string
	Json  *simplejson.Json `json:"fuzzy_like_this_field"`
}

func FltField(field string) *FuzzyLikeThisFiedQuery {
	return &FuzzyLikeThisFiedQuery{field, simplejson.New()}
}

func (j *FuzzyLikeThisFiedQuery) LikeText(v interface{}) *FuzzyLikeThisFiedQuery {
	j.Json.SetPath([]string{j.field, "like_text"}, v)
	return j
}
func (j *FuzzyLikeThisFiedQuery) MaxQueryTerms(v int) *FuzzyLikeThisFiedQuery {
	j.Json.SetPath([]string{j.field, "max_query_terms"}, v)
	return j
}
func (j *FuzzyLikeThisFiedQuery) IgnoreTf(v bool) *FuzzyLikeThisFiedQuery {
	j.Json.SetPath([]string{j.field, "ignore_tf"}, v)
	return j
}
func (j *FuzzyLikeThisFiedQuery) Boost(v float64) *FuzzyLikeThisFiedQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *FuzzyLikeThisFiedQuery) Fuzziness(v float64) *FuzzyLikeThisFiedQuery {
	j.Json.SetPath([]string{j.field, "fuzziness"}, v)
	return j
}
func (j *FuzzyLikeThisFiedQuery) PrefixLength(v int) *FuzzyLikeThisFiedQuery {
	j.Json.SetPath([]string{j.field, "prefix_length"}, v)
	return j
}
func (j *FuzzyLikeThisFiedQuery) Analyzer(v interface{}) *FuzzyLikeThisFiedQuery {
	j.Json.SetPath([]string{j.field, "analyzer"}, v)
	return j
}

func (j *FuzzyLikeThisFiedQuery) Encode() []byte {
	return encode(j)
}

type FuzzyQuery struct {
	field string
	Json  *simplejson.Json `json:"fuzzy"`
}

func Fuzzy(field string) *FuzzyQuery {
	return &FuzzyQuery{field, simplejson.New()}
}

func (j *FuzzyQuery) Fuzzy(v interface{}) *FuzzyQuery {
	j.Json.SetPath([]string{j.field}, v)
	return j
}
func (j *FuzzyQuery) Value(v interface{}) *FuzzyQuery {
	j.Json.SetPath([]string{j.field, "value"}, v)
	return j
}
func (j *FuzzyQuery) Boost(v float64) *FuzzyQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *FuzzyQuery) Fuzziness(v interface{}) *FuzzyQuery {
	j.Json.SetPath([]string{j.field, "fuzziness"}, v)
	return j
}
func (j *FuzzyQuery) PrefixLength(v int) *FuzzyQuery {
	j.Json.SetPath([]string{j.field, "prefix_length"}, v)
	return j
}
func (j *FuzzyQuery) MaxExpansions(v int) *FuzzyQuery {
	j.Json.SetPath([]string{j.field, "max_expansions"}, v)
	return j
}

func (j *FuzzyQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-geo-shape-query.html
type GeoShapeQuery struct {
	Json *simplejson.Json `json:"geo_shape"`
}

func GeoShape() *GeoShapeQuery {
	return &GeoShapeQuery{simplejson.New()}
}

func (j *GeoShapeQuery) Type(v string) *GeoShapeQuery {
	j.Json.SetPath([]string{"location", "shape", "type"}, v)
	return j
}
func (j *GeoShapeQuery) Coordinates(v interface{}) *GeoShapeQuery {
	j.Json.SetPath([]string{"location", "shape", "coordinates"}, v)
	return j
}
func (j *GeoShapeQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-has-child-query.html
type HasChildQuery struct {
	Json *simplejson.Json `json:"has_child"`
}

func HasChild() *HasChildQuery {
	return &HasChildQuery{simplejson.New()}
}

func (j *HasChildQuery) Type(v string) *HasChildQuery {
	j.Json.Set("type", v)
	return j
}
func (j *HasChildQuery) ScoreMode(v string) *HasChildQuery {
	j.Json.Set("score_mode", v)
	return j
}
func (j *HasChildQuery) MinChildren(v int) *HasChildQuery {
	j.Json.Set("min_children", v)
	return j
}
func (j *HasChildQuery) MaxChildren(v int) *HasChildQuery {
	j.Json.Set("max_children", v)
	return j
}
func (j *HasChildQuery) Query(v interface{}) *HasChildQuery {
	j.Json.Set("query", v)
	return j
}
func (j *HasChildQuery) Filter(v interface{}) *HasChildQuery {
	j.Json.Set("filter", v)
	return j
}
func (j *HasChildQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-has-parent-query.html
type HasParentQuery struct {
	Json *simplejson.Json `json:"has_parent"`
}

func HasParent() *HasParentQuery {
	return &HasParentQuery{simplejson.New()}
}
func (j *HasParentQuery) Type(v string) *HasParentQuery {
	j.Json.Set("parent_type", v)
	return j
}
func (j *HasParentQuery) ScoreMode(v string) *HasParentQuery {
	j.Json.Set("score_mode", v)
	return j
}
func (j *HasParentQuery) Query(v interface{}) *HasParentQuery {
	j.Json.Set("query", v)
	return j
}
func (j *HasParentQuery) Filter(v interface{}) *HasParentQuery {
	j.Json.Set("filter", v)
	return j
}
func (j *HasParentQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html
type IdsQuery struct {
	Json *simplejson.Json `json:"ids"`
}

func Ids() *IdsQuery {
	return &IdsQuery{simplejson.New()}
}
func (j *IdsQuery) Type(v ...string) *IdsQuery {
	if len(v) > 1 {
		j.Json.Set("type", v)
	} else {
		j.Json.Set("type", v[0])
	}
	return j
}
func (j *IdsQuery) Values(v ...interface{}) *IdsQuery {
	if len(v) > 1 {
		j.Json.Set("values", v)
	} else {
		j.Json.Set("values", v[0])
	}
	return j
}
func (j *IdsQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-indices-query.html
type IndicesQuery struct {
	Json *simplejson.Json `json:"indices"`
}

func Indices() *IndicesQuery {
	return &IndicesQuery{simplejson.New()}
}
func (j *IndicesQuery) Indices(v ...string) *IndicesQuery {
	if len(v) > 1 {
		j.Json.Set("indices", v)
	} else {
		j.Json.Set("index", v[0])
	}
	return j
}
func (j *IndicesQuery) Query(v interface{}) *IndicesQuery {
	j.Json.Set("query", v)
	return j
}
func (j *IndicesQuery) NoMatchQuery(v interface{}) *IndicesQuery {
	j.Json.Set("no_match_query", v)
	return j
}
func (j *IndicesQuery) Filter(v interface{}) *IndicesQuery {
	j.Json.Set("filter", v)
	return j
}
func (j *IndicesQuery) NoMatchFilter(v interface{}) *IndicesQuery {
	j.Json.Set("no_match_filter", v)
	return j
}
func (j *IndicesQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-mlt-query.html
type MoreLikeThisQuery struct {
	Json *simplejson.Json `json:"more_like_this"`
}

func MoreLikeThis() *MoreLikeThisQuery {
	return &MoreLikeThisQuery{simplejson.New()}
}
func (j *MoreLikeThisQuery) Fields(v ...string) *MoreLikeThisQuery {
	if len(v) > 1 {
		j.Json.Set("fields", v)
	} else {
		j.Json.Set("fields", v[0])
	}
	return j
}
func (j *MoreLikeThisQuery) LikeText(v string) *MoreLikeThisQuery {
	j.Json.Set("like_text", v)
	return j
}
func (j *MoreLikeThisQuery) Docs(v ...interface{}) *MoreLikeThisQuery {
	if len(v) > 1 {
		j.Json.Set("docs", v)
	} else {
		j.Json.Set("docs", v[0])
	}
	return j
}
func (j *MoreLikeThisQuery) Ids(v ...string) *MoreLikeThisQuery {
	if len(v) > 1 {
		j.Json.Set("ids", v)
	} else {
		j.Json.Set("ids", v[0])
	}
	return j
}
func (j *MoreLikeThisQuery) Include(v bool) *MoreLikeThisQuery {
	j.Json.Set("include", v)
	return j
}
func (j *MoreLikeThisQuery) Exclude(v bool) *MoreLikeThisQuery {
	j.Json.Set("exclude", v)
	return j
}
func (j *MoreLikeThisQuery) PercentTermsToMatch(v float64) *MoreLikeThisQuery {
	j.Json.Set("percent_terms_to_match", v)
	return j
}
func (j *MoreLikeThisQuery) MinTermFreq(v int) *MoreLikeThisQuery {
	j.Json.Set("min_term_freq", v)
	return j
}
func (j *MoreLikeThisQuery) MaxQueryTerms(v int) *MoreLikeThisQuery {
	j.Json.Set("max_query_terms", v)
	return j
}
func (j *MoreLikeThisQuery) StopWords(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("stop_words", v)
	return j
}
func (j *MoreLikeThisQuery) MinDocFreq(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("min_doc_freq", v)
	return j
}
func (j *MoreLikeThisQuery) MaxDocFreq(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("max_doc_freq", v)
	return j
}
func (j *MoreLikeThisQuery) MinWordLength(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("min_word_length", v)
	return j
}
func (j *MoreLikeThisQuery) MaxWordLength(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("max_word_length", v)
	return j
}
func (j *MoreLikeThisQuery) BoostTerms(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("boost_terms", v)
	return j
}
func (j *MoreLikeThisQuery) Boost(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("boost", v)
	return j
}
func (j *MoreLikeThisQuery) Analyzer(v interface{}) *MoreLikeThisQuery {
	j.Json.Set("analyzer", v)
	return j
}
func (j *MoreLikeThisQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-mlt-field-query.html
type MoreLikeThisFieldQuery struct {
	field string
	Json  *simplejson.Json `json:"more_like_this_field"`
}

func MoreLikeThisField(field string) *MoreLikeThisFieldQuery {
	return &MoreLikeThisFieldQuery{field, simplejson.New()}
}
func (j *MoreLikeThisFieldQuery) LikeText(v string) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "like_text"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) PercentTermsToMatch(v float64) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "percent_terms_to_match"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) MinTermFreq(v int) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "min_term_freq"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) MaxQueryTerms(v int) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "max_query_terms"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) StopWords(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "stop_words"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) MinDocFreq(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "min_doc_freq"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) MaxDocFreq(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "max_doc_freq"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) MinWordLength(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "min_word_length"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) MaxWordLength(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "max_word_length"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) BoostTerms(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "boost_terms"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) Boost(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) Analyzer(v interface{}) *MoreLikeThisFieldQuery {
	j.Json.SetPath([]string{j.field, "analyzer"}, v)
	return j
}
func (j *MoreLikeThisFieldQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html#query-dsl-nested-query
type NestedQuery struct {
	Json *simplejson.Json `json:"nested"`
}

func Nested() *NestedQuery {
	return &NestedQuery{simplejson.New()}
}
func (j *NestedQuery) Path(v string) *NestedQuery {
	j.Json.Set("path", v)
	return j
}
func (j *NestedQuery) Query(v interface{}) *NestedQuery {
	j.Json.Set("query", v)
	return j
}
func (j *NestedQuery) Filter(v interface{}) *NestedQuery {
	j.Json.Set("filter", v)
	return j
}
func (j *NestedQuery) ScoreMode(v string) *NestedQuery {
	j.Json.Set("score_mode", v)
	return j
}
func (j *NestedQuery) Cache(v bool) *NestedQuery {
	j.Json.Set("_cache", v)
	return j
}
func (j *NestedQuery) Join(v bool) *NestedQuery {
	j.Json.Set("join", v)
	return j
}
func (j *NestedQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixQuery struct {
	field string
	Json  *simplejson.Json `json:"prefix"`
}

func Prefix(field string) *PrefixQuery {
	return &PrefixQuery{field, simplejson.New()}
}
func (j *PrefixQuery) Val(v interface{}) *PrefixQuery {
	j.Json.Set(j.field, v)
	return j
}
func (j *PrefixQuery) Value(v interface{}) *PrefixQuery {
	j.Json.SetPath([]string{j.field, "value"}, v)
	return j
}
func (j *PrefixQuery) Prefix(v interface{}) *PrefixQuery {
	j.Json.SetPath([]string{j.field, "prefix"}, v)
	return j
}
func (j *PrefixQuery) Boost(v interface{}) *PrefixQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *PrefixQuery) Cache(v bool) *PrefixQuery {
	j.Json.Set("_cache", v)
	return j
}
func (j *PrefixQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html
type QueryStringQuery struct {
	Json *simplejson.Json `json:"query_string"`
}

func QueryString() *QueryStringQuery {
	return &QueryStringQuery{simplejson.New()}
}

func (j *QueryStringQuery) Query(v interface{}) *QueryStringQuery {
	j.Json.Set("query", v)
	return j
}
func (j *QueryStringQuery) DefaultField(v interface{}) *QueryStringQuery {
	j.Json.Set("default_field", v)
	return j
}
func (j *QueryStringQuery) Fields(v ...interface{}) *QueryStringQuery {
	j.Json.Set("fields", v)
	return j
}

func (j *QueryStringQuery) DefaultOperator(v interface{}) *QueryStringQuery {
	j.Json.Set("default_operator", v)
	return j
}
func (j *QueryStringQuery) Analyzer(v interface{}) *QueryStringQuery {
	j.Json.Set("analyzer", v)
	return j
}
func (j *QueryStringQuery) AllowLeadingWildcard(v interface{}) *QueryStringQuery {
	j.Json.Set("allow_leading_wildcard", v)
	return j
}
func (j *QueryStringQuery) LowercaseExpandedTerms(v interface{}) *QueryStringQuery {
	j.Json.Set("lowercase_expanded_terms", v)
	return j
}
func (j *QueryStringQuery) EnablePositionIncrements(v interface{}) *QueryStringQuery {
	j.Json.Set("enable_position_increments", v)
	return j
}
func (j *QueryStringQuery) FuzzyMaxExpansions(v interface{}) *QueryStringQuery {
	j.Json.Set("fuzzy_max_expansions", v)
	return j
}
func (j *QueryStringQuery) Fuzziness(v interface{}) *QueryStringQuery {
	j.Json.Set("fuzziness", v)
	return j
}
func (j *QueryStringQuery) FuzzyPrefixLength(v interface{}) *QueryStringQuery {
	j.Json.Set("fuzzy_prefix_length", v)
	return j
}
func (j *QueryStringQuery) Boost(v interface{}) *QueryStringQuery {
	j.Json.Set("boost", v)
	return j
}
func (j *QueryStringQuery) AnalyzeWildcard(v interface{}) *QueryStringQuery {
	j.Json.Set("analyze_wildcard", v)
	return j
}
func (j *QueryStringQuery) AutoGeneratePhraseQueries(v interface{}) *QueryStringQuery {
	j.Json.Set("auto_generate_phrase_queries", v)
	return j
}
func (j *QueryStringQuery) MinimumShouldMatch(v interface{}) *QueryStringQuery {
	j.Json.Set("minimum_should_match", v)
	return j
}
func (j *QueryStringQuery) Lenient(v interface{}) *QueryStringQuery {
	j.Json.Set("lenient", v)
	return j
}
func (j *QueryStringQuery) Locale(v interface{}) *QueryStringQuery {
	j.Json.Set("locale", v)
	return j
}
func (j *QueryStringQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-simple-query-string-query.html
type SimpleQueryStringQuery struct {
	Json *simplejson.Json `json:"simple_query_string"`
}

func SimpleQueryString() *SimpleQueryStringQuery {
	return &SimpleQueryStringQuery{simplejson.New()}
}

func (j *SimpleQueryStringQuery) Query(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("query", v)
	return j
}
func (j *SimpleQueryStringQuery) Fields(v ...interface{}) *SimpleQueryStringQuery {
	if len(v) > 1 {
		j.Json.Set("fields", v)
	} else {
		j.Json.Set("fields", v[0])
	}
	return j
}
func (j *SimpleQueryStringQuery) DefaultOperator(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("default_operator", v)
	return j
}
func (j *SimpleQueryStringQuery) Analyzer(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("analyzer", v)
	return j
}

//The available flags are: ALL, NONE, AND, OR, NOT, PREFIX, PHRASE, PRECEDENCE, ESCAPE, WHITESPACE, FUZZY, NEAR,
// and SLOP
func (j *SimpleQueryStringQuery) Flags(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("flags", v)
	return j
}
func (j *SimpleQueryStringQuery) LowercaseExpandedTerms(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("lowercase_expanded_terms", v)
	return j
}

func (j *SimpleQueryStringQuery) Lenient(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("lenient", v)
	return j
}
func (j *SimpleQueryStringQuery) Locale(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("locale", v)
	return j
}
func (j *SimpleQueryStringQuery) Boost(v interface{}) *SimpleQueryStringQuery {
	j.Json.Set("boost", v)
	return j
}
func (j *SimpleQueryStringQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-regexp-query.html
type RegexpQuery struct {
	field string
	Json  *simplejson.Json `json:"regexp"`
}

func Regexp(field string) *RegexpQuery {
	return &RegexpQuery{field, simplejson.New()}
}

func (j *RegexpQuery) Val(v string) *RegexpQuery {
	j.Json.Set(j.field, v)
	return j
}
func (j *RegexpQuery) Value(v string) *RegexpQuery {
	j.Json.SetPath([]string{j.field, "value"}, v)
	return j
}
func (j *RegexpQuery) Flags(v string) *RegexpQuery {
	j.Json.SetPath([]string{j.field, "flags"}, v)
	return j
}
func (j *RegexpQuery) Boost(v float64) *RegexpQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *RegexpQuery) Name(v string) *RegexpQuery {
	j.Json.Set("_name", v)
	return j
}
func (j *RegexpQuery) Cache(v bool) *RegexpQuery {
	j.Json.Set("_cache", v)
	return j
}
func (j *RegexpQuery) CacheKey(v string) *RegexpQuery {
	j.Json.Set("_cache_key", v)
	return j
}
func (j *RegexpQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-span-first-query.html
type SpanFirstQuery struct {
	Json *simplejson.Json `json:"span_first"`
}

func SpanFirst() *SpanFirstQuery {
	return &SpanFirstQuery{simplejson.New()}
}
func (j *SpanFirstQuery) Match(k, v string) *SpanFirstQuery {
	j.Json.SetPath([]string{"match", "span_term", k}, v)
	return j
}
func (j *SpanFirstQuery) End(v int) *SpanFirstQuery {
	j.Json.Set("end", v)
	return j
}
func (j *SpanFirstQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-span-multi-term-query.html
type SpanMultiQuery struct {
	Json *simplejson.Json `json:"span_multi"`
}

func SpanMulti() *SpanMultiQuery {
	return &SpanMultiQuery{simplejson.New()}
}
func (j *SpanMultiQuery) Match(v interface{}) *SpanMultiQuery {
	j.Json.Set("match", v)
	return j
}
func (j *SpanMultiQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-span-near-query.html
type SpanNearQuery struct {
	Json *simplejson.Json `json:"span_near"`
}

func SpanNear() *SpanNearQuery {
	return &SpanNearQuery{simplejson.New()}
}
func (j *SpanNearQuery) Clauses(v ...interface{}) *SpanNearQuery {
	if len(v) > 1 {
		j.Json.Set("clauses", v)
	} else {
		j.Json.Set("clauses", v[0])
	}
	return j
}
func (j *SpanNearQuery) Slop(v interface{}) *SpanNearQuery {
	j.Json.Set("slop", v)
	return j
}
func (j *SpanNearQuery) InOrder(v bool) *SpanNearQuery {
	j.Json.Set("in_order", v)
	return j
}
func (j *SpanNearQuery) CollectPayloads(v bool) *SpanNearQuery {
	j.Json.Set("collect_payloads", v)
	return j
}
func (j *SpanNearQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-span-not-query.html
type SpanNotQuery struct {
	Json *simplejson.Json `json:"span_not"`
}

func SpanNot() *SpanNotQuery {
	return &SpanNotQuery{simplejson.New()}
}
func (j *SpanNotQuery) Include(v interface{}) *SpanNotQuery {
	j.Json.Set("include", v)
	return j
}
func (j *SpanNotQuery) Exclude(v interface{}) *SpanNotQuery {
	j.Json.Set("exclude", v)
	return j
}
func (j *SpanNotQuery) Pre(v interface{}) *SpanNotQuery {
	j.Json.Set("pre", v)
	return j
}
func (j *SpanNotQuery) Post(v interface{}) *SpanNotQuery {
	j.Json.Set("post", v)
	return j
}
func (j *SpanNotQuery) Dist(v interface{}) *SpanNotQuery {
	j.Json.Set("dist", v)
	return j
}
func (j *SpanNotQuery) Encode() []byte {
	return encode(j)
}

//
type SpanOrQuery struct {
	Json *simplejson.Json `json:"span_or"`
}

func SpanOr() *SpanOrQuery {
	return &SpanOrQuery{simplejson.New()}
}
func (j *SpanOrQuery) Clauses(v ...interface{}) *SpanOrQuery {
	if len(v) > 1 {
		j.Json.Set("clauses", v)
	} else {
		j.Json.Set("clauses", v[0])
	}
	return j
}
func (j *SpanOrQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-span-term-query.html
type SpanTermQuery struct {
	field string
	Json  *simplejson.Json `json:"span_term"`
}

func SpanTerm(field string) *SpanTermQuery {
	return &SpanTermQuery{field, simplejson.New()}
}
func (j *SpanTermQuery) Val(v interface{}) *SpanTermQuery {
	j.Json.Set(j.field, v)
	return j
}
func (j *SpanTermQuery) Value(v interface{}) *SpanTermQuery {
	j.Json.SetPath([]string{j.field, "value"}, v)
	return j
}
func (j *SpanTermQuery) Term(v interface{}) *SpanTermQuery {
	j.Json.SetPath([]string{j.field, "term"}, v)
	return j
}
func (j *SpanTermQuery) Boost(v interface{}) *SpanTermQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *SpanTermQuery) Encode() []byte {
	return encode(j)
}

type TermsQuery struct {
	Json *simplejson.Json `json:"terms"`
}

func Terms() *TermsQuery {
	return &TermsQuery{simplejson.New()}
}
func (j *TermsQuery) Tags(v ...interface{}) *TermsQuery {
	if len(v) > 1 {
		j.Json.Set("tags", v)
	} else {
		j.Json.Set("tags", v[0])
	}
	return j
}
func (j *TermsQuery) MinimumShouldMatch(v interface{}) *TermsQuery {
	j.Json.Set("minimum_should_match", v)
	return j
}
func (j *TermsQuery) Terms(k string, v ...interface{}) *TermsQuery {
	j.Json.Set(k, v)
	return j
}
func (j *TermsQuery) Execution(v string) *TermsQuery {
	j.Json.Set("execution", v)
	return j
}
func (j *TermsQuery) Cache(v bool) *TermsQuery {
	j.Json.Set("_cache", v)
	return j
}
func (j *TermsQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-top-children-query.html
type TopChildrenQuery struct {
	Json *simplejson.Json `json:"top_children"`
}

func TopChildren() *TopChildrenQuery {
	return &TopChildrenQuery{simplejson.New()}
}
func (j *TopChildrenQuery) Type(v interface{}) *TopChildrenQuery {
	j.Json.Set("type", v)
	return j
}
func (j *TopChildrenQuery) Query(v interface{}) *TopChildrenQuery {
	j.Json.Set("query", v)
	return j
}
func (j *TopChildrenQuery) Score(v interface{}) *TopChildrenQuery {
	j.Json.Set("score", v)
	return j
}
func (j *TopChildrenQuery) Factor(v interface{}) *TopChildrenQuery {
	j.Json.Set("factor", v)
	return j
}
func (j *TopChildrenQuery) Scope(v interface{}) *TopChildrenQuery {
	j.Json.Set("_scope", v)
	return j
}
func (j *TopChildrenQuery) IncrementalFactor(v interface{}) *TopChildrenQuery {
	j.Json.Set("incremental_factor", v)
	return j
}
func (j *TopChildrenQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html
type WildcardQuery struct {
	field string
	Json  *simplejson.Json `json:"wildcard"`
}

func Wildcard(field string) *WildcardQuery {
	return &WildcardQuery{field, simplejson.New()}
}
func (j *WildcardQuery) Val(v interface{}) *WildcardQuery {
	j.Json.Set(j.field, v)
	return j
}
func (j *WildcardQuery) Value(v interface{}) *WildcardQuery {
	j.Json.SetPath([]string{j.field, "value"}, v)
	return j
}
func (j *WildcardQuery) Wildcard(v interface{}) *WildcardQuery {
	j.Json.SetPath([]string{j.field, "wildcard"}, v)
	return j
}
func (j *WildcardQuery) Boost(v interface{}) *WildcardQuery {
	j.Json.SetPath([]string{j.field, "boost"}, v)
	return j
}
func (j *WildcardQuery) Encode() []byte {
	return encode(j)
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/query-dsl-template-query.html
type TemplateQuery struct {
	Json *simplejson.Json `json:"template"`
}

func Template() *TemplateQuery {
	return &TemplateQuery{simplejson.New()}
}
func (j *TemplateQuery) Query(v interface{}) *TemplateQuery {
	j.Json.Set("query", v)
	return j
}

//struct{}{}
func (j *TemplateQuery) QueryMap(k string, v interface{}) *TemplateQuery {
	j.Json.SetPath([]string{"query", k}, v)
	return j
}
func (j *TemplateQuery) Params(k string, v interface{}) *TemplateQuery {
	j.Json.SetPath([]string{"params", k}, v)
	return j
}
func (j *TemplateQuery) Encode() []byte {
	return encode(j)
}
