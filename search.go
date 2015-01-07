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

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bitly/go-simplejson"
)

type SearchDsl struct {
	Json   *simplejson.Json
	_index []string
	_type  []string
	para   map[string]string
}

//the DSL search of usage:
//	s := NewServer("localhost", "9200")
//	Search().Index("plasticine").Type("user").From(1).Size(10).Query(
// 		QueryString().Query("star").Fields("name", "nickname"),
// 	).Filter(
// 		Terms().Terms("age", "20"),
// 	).Result(s)
func Search() *SearchDsl {
	return &SearchDsl{
		simplejson.New(),
		nil, nil, nil,
	}
}

type Jsoner interface {
	Encode() []byte
}

func (j *SearchDsl) Encode() []byte {
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

func (j *SearchDsl) Index(v ...string) *SearchDsl {
	j._index = v
	return j
}
func (j *SearchDsl) Type(v ...string) *SearchDsl {
	j._type = v
	return j
}
func (j *SearchDsl) SearchType(v string) *SearchDsl {
	j.para["search_type"] = v
	return j
}
func (j *SearchDsl) Scroll(v string) *SearchDsl {
	j.para["scroll"] = v
	return j
}
func (j *SearchDsl) Preference(v string) *SearchDsl {
	j.para["preference"] = v
	return j
}
func (j *SearchDsl) From(v interface{}) *SearchDsl {
	j.Json.Set("from", v)
	return j
}
func (j *SearchDsl) Size(v interface{}) *SearchDsl {
	j.Json.Set("size", v)
	return j
}
func (j *SearchDsl) Explain(v bool) *SearchDsl {
	j.Json.Set("explain", v)
	return j
}
func (j *SearchDsl) Version(v bool) *SearchDsl {
	j.Json.Set("version", v)
	return j
}

func (j *SearchDsl) MinScore(v float64) *SearchDsl {
	j.Json.Set("min_score", v)
	return j
}

//v is map[string]float64{"index1":1.2}
func (j *SearchDsl) IndicesBoost(v ...interface{}) *SearchDsl {
	if len(v) > 1 {
		for _, val := range v {
			j.Json.Set("indices_boost", val)
		}
	} else {
		j.Json.Set("indices_boost", v[0])
	}
	return j
}
func (j *SearchDsl) Aggs(v ...interface{}) *SearchDsl {
	if len(v) > 1 {
		for i := 0; i < len(v); i++ {
			js := v[i].(*SearchAggregation)
			if js != nil {
				m, err := js.Json.Map()
				if err == nil {
					for key, val := range m {
						j.Json.SetPath([]string{"aggs", key}, val)
					}
				}
			}
		}
	} else {
		j.Json.Set("aggs", v[0])
	}
	return j
}

func (j *SearchDsl) Facets(v ...interface{}) *SearchDsl {
	for i := 0; i < len(v); i++ {
		js := v[i].(*SearchFacet)
		if js != nil {
			m, err := js.Json.Map()
			if err == nil {
				for key, val := range m {
					j.Json.SetPath([]string{"facets", key}, val)
				}
			}
		}
	}
	return j
}
func (j *SearchDsl) Query(v interface{}) *SearchDsl {
	j.Json.Set("query", v)
	return j
}
func (j *SearchDsl) Filter(v interface{}) *SearchDsl {
	j.Json.Set("filter", v)
	return j
}
func (j *SearchDsl) Sort(v ...interface{}) *SearchDsl {
	j.Json.Set("sort", v)
	return j
}
func (j *SearchDsl) Source(v ...interface{}) *SearchDsl {
	if len(v) > 1 {
		for i := 0; i < len(v); i++ {
			js := v[i].(*simplejson.Json)
			if js != nil {
				m, err := js.Map()
				if err == nil {
					for key, val := range m {
						j.Json.SetPath([]string{"_source", key}, val)
					}
				}
			}
		}
	} else {
		j.Json.Set("_source", v[0])
	}
	return j
}
func (j *SearchDsl) Fields(v ...interface{}) *SearchDsl {
	if len(v) > 0 {
		j.Json.Set("fields", v)
	} else {
		j.Json.Set("fields", []string{})
	}
	return j
}
func (j *SearchDsl) PartialFields(k string, v ...*simplejson.Json) *SearchDsl {
	if len(v) > 1 {
		for i := 0; i < len(v); i++ {
			m, err := v[i].Map()
			if err == nil {
				for key, val := range m {
					j.Json.SetPath([]string{"partial_fields", k, key}, val)
				}
			}
		}
	} else {
		j.Json.SetPath([]string{"partial_fields", k}, v[0])
	}
	return j
}
func (j *SearchDsl) ScriptFields(k string, v interface{}) *SearchDsl {
	j.Json.SetPath([]string{"script_fields", k}, v)
	return j
}
func (j *SearchDsl) FielddataFields(field ...string) *SearchDsl {
	if len(field) > 0 {
		j.Json.Set("fielddata_fields", field)
	} else {
		j.Json.Set("fielddata_fields", []string{})
	}
	return j
}
func (j *SearchDsl) PostFilter(v interface{}) *SearchDsl {
	j.Json.Set("post_filter", v)
	return j
}

func (j *SearchDsl) Highlight(v interface{}) *SearchDsl {
	if js := v.(*SearchHighlight); js != nil {
		j.Json.Set("highlight", js.Json)
	}
	return j
}
func (j *SearchDsl) Rescoring(v ...interface{}) *SearchDsl {
	if len(v) < 2 {
		if js := v[0].(*SearchRescore); js != nil {
			j.Json.Set("rescore", js.Json)
		}
	} else {
		jsons := make([]*simplejson.Json, len(v))
		for i := 0; i < len(v); i++ {
			if js := v[i].(*SearchRescore); js != nil {
				jsons[i] = js.Json
			}
		}
		j.Json.Set("rescore", jsons)
	}
	return j
}

type SearchHighlight struct {
	field string
	Json  *simplejson.Json `json:"highlight"`
}

func Highlight(field string) *SearchHighlight {
	return &SearchHighlight{field, simplejson.New()}
}
func (j *SearchHighlight) Order(v string) *SearchHighlight {
	j.Json.Set("order", v)
	return j
}
func (j *SearchHighlight) Tags(v ...string) *SearchHighlight {
	pre_tags := []string{}
	post_tags := []string{}
	for _, val := range v {
		pre_tags = append(pre_tags, fmt.Sprintf("<%v>", val))
		post_tags = append(post_tags, fmt.Sprintf("</%v>", val))
	}
	j.Json.Set("pre_tags", pre_tags)
	j.Json.Set("post_tags", post_tags)
	return j
}
func (j *SearchHighlight) FragmentSize(v int) *SearchHighlight {
	j.Json.SetPath([]string{"fields", j.field, "fragment_size"}, v)
	return j
}
func (j *SearchHighlight) NumberOfFragments(v int) *SearchHighlight {
	j.Json.SetPath([]string{"fields", j.field, "number_of_fragments"}, v)
	return j
}
func (j *SearchHighlight) NoMatchSize(v int) *SearchHighlight {
	j.Json.SetPath([]string{"fields", j.field, "no_match_size"}, v)
	return j
}

func (j *SearchHighlight) ForceSource(v bool) *SearchHighlight {
	j.Json.SetPath([]string{"fields", j.field, "force_source"}, v)
	return j
}

//values are:plain, postings and fvh
func (j *SearchHighlight) Type(v string) *SearchHighlight {
	j.Json.SetPath([]string{"fields", j.field, "type"}, v)
	return j
}
func (j *SearchHighlight) HighlightQuery(v interface{}) *SearchHighlight {
	j.Json.SetPath([]string{"fields", j.field, "highlight_query"}, v)
	return j
}
func (j *SearchHighlight) MatchedFields(v interface{}) *SearchHighlight {
	j.Json.SetPath([]string{"fields", j.field, "matched_fields"}, v)
	return j
}

type SearchRescore struct {
	Json *simplejson.Json `json:"rescore"`
}

func Rescoring() *SearchRescore {
	return &SearchRescore{simplejson.New()}
}
func (j *SearchRescore) WindowSize(v interface{}) *SearchRescore {
	j.Json.Set("window_size", v)
	return j
}
func (j *SearchRescore) ScoreMode(v interface{}) *SearchRescore {
	j.Json.SetPath([]string{"query", "score_mode"}, v)
	return j
}
func (j *SearchRescore) QueryWeight(v float64) *SearchRescore {
	j.Json.SetPath([]string{"query", "query_weight"}, v)
	return j
}
func (j *SearchRescore) RescoreQueryWeight(v float64) *SearchRescore {
	j.Json.SetPath([]string{"query", "rescore_query_weight"}, v)
	return j
}
func (j *SearchRescore) RescoreQuery(v interface{}) *SearchRescore {
	j.Json.SetPath([]string{"query", "rescore_query"}, v)
	return j
}

func Sort(v ...interface{}) *simplejson.Json {
	j := simplejson.New()
	if len(v) > 1 {
		j.Set("sort", v)
	} else {
		j.Set("sort", v[0])
	}
	return j
}
func Desc(v string) *simplejson.Json {
	j := simplejson.New()
	j.Set(v, "desc")
	return j
}
func Asc(v string) *simplejson.Json {
	j := simplejson.New()
	j.Set(v, "asc")
	return j
}
func Order(field, order string) *simplejson.Json {
	j := simplejson.New()
	j.SetPath([]string{field, "order"}, order)
	return j
}

func encode(v interface{}) []byte {
	var b []byte
	var err error
	if Debug {
		b, err = json.MarshalIndent(v, "", "  ")
	} else {
		b, err = json.Marshal(v)
	}
	if err != nil {
		return nil
	}
	return b
}

func Include(v ...string) *simplejson.Json {
	j := simplejson.New()
	if len(v) > 1 {
		j.Set("include", v)
	} else {
		j.Set("include", v[0])
	}
	return j
}
func Exclude(v ...string) *simplejson.Json {
	j := simplejson.New()
	if len(v) > 1 {
		j.Set("exclude", v)
	} else {
		j.Set("exclude", v[0])
	}
	return j
}
func (j *SearchDsl) Result(s *Server) (*SearchResult, error) {
	indices := strings.Join(j._index, ",")
	types := strings.Join(j._type, ",")
	ctx := Post(s.getUrl(getPath(indices, types, "_search"), getQuery(j.para)), s.hr)
	ctx.setData(j.Encode())
	if Debug {
		println("dsl:" + string(j.Encode()))
	}
	return ctx.GetSearch()
}
