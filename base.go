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
	"strconv"
)

type BaseResponse struct {
	Ok      bool             `json:"ok"`
	Index   string           `json:"_index,omitempty"`
	Type    string           `json:"_type,omitempty"`
	Id      string           `json:"_id,omitempty"`
	Source  *json.RawMessage `json:"_source,omitempty"` // depends on the schema you've defined
	Version int              `json:"_version,omitempty"`
	Found   bool             `json:"found,omitempty"`
	Exists  bool             `json:"exists,omitempty"`
	Created bool             `json:"created,omitempty"`
	Matches []string         `json:"matches,omitempty"` // percolate matches
}

type SuggestionOption struct {
	Payload json.RawMessage `json:"payload"`
	//Score   Float64Nullable `json:"score,omitempty"`
	Score float64 `json:"score,omitempty"`
	Text  string  `json:"text"`
}

type Suggestion struct {
	Length  int                `json:"length"`
	Offset  int                `json:"offset"`
	Options []SuggestionOption `json:"options"`
	Text    string             `json:"text"`
}

type Suggestions map[string][]Suggestion

type SearchResult struct {
	RawJSON      []byte
	Took         int             `json:"took"`
	TimedOut     bool            `json:"timed_out"`
	ShardStatus  Status          `json:"_shards"`
	Hits         Hits            `json:"hits"`
	Facets       json.RawMessage `json:"facets,omitempty"` // structure varies on query
	ScrollId     string          `json:"_scroll_id,omitempty"`
	Aggregations json.RawMessage `json:"aggregations,omitempty"` // structure varies on query
	Suggestions  Suggestions     `json:"suggest,omitempty"`
}

type Hits struct {
	Total int `json:"total"`
	//	MaxScore float32 `json:"max_score"`
	Hits []Hit `json:"hits"`
}

func (h *Hits) Len() int {
	return len(h.Hits)
}

type HighlightClass map[string][]string

type Hit struct {
	Index string  `json:"_index"`
	Type  string  `json:"_type,omitempty"`
	Id    string  `json:"_id"`
	Score float64 `json:"_score,omitempty"` // Filters (no query) dont have score, so is null
	//Score       Float64Nullable  `json:"_score,omitempty"` // Filters (no query) dont have score, so is null
	Source      *json.RawMessage `json:"_source"` // marshalling left to consumer
	Fields      *json.RawMessage `json:"fields"`  // when a field arg is passed to ES, instead of _source it returns fields
	Explanation *Explanation     `json:"_explanation,omitempty"`
	Highlight   *HighlightClass  `json:"highlight,omitempty"`
}
type Explanation struct {
	Value       float64        `json:"value"`
	Description string         `json:"description"`
	Details     []*Explanation `json:"details,omitempty"`
}

// Elasticsearch returns some invalid (according to go) json, with floats having...
//
// json: cannot unmarshal null into Go value of type float64 (see last field.)
//
// "hits":{"total":6808,"max_score":null,
//    "hits":[{"_index":"10user","_type":"user","_id":"751820","_score":null,
type Float64Nullable float64

func (i *Float64Nullable) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	if in, err := strconv.ParseFloat(string(data), 64); err != nil {
		return err
	} else {
		*i = Float64Nullable(in)
	}
	return nil
}

// StatusInt is required because /_optimize, at least, returns its status as
// strings instead of integers.
type StatusInt int

func (self *StatusInt) UnmarshalJSON(b []byte) error {
	s := ""
	if json.Unmarshal(b, &s) == nil {
		if i, err := strconv.Atoi(s); err == nil {
			*self = StatusInt(i)
			return nil
		}
	}
	i := 0
	err := json.Unmarshal(b, &i)
	if err == nil {
		*self = StatusInt(i)
	}
	return err
}

func (self *StatusInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(*self)
}

// StatusBool is required because /_optimize, at least, returns its status as
// strings instead of booleans.
type StatusBool bool

func (self *StatusBool) UnmarshalJSON(b []byte) error {
	s := ""
	if json.Unmarshal(b, &s) == nil {
		switch s {
		case "true":
			*self = StatusBool(true)
			return nil
		case "false":
			*self = StatusBool(false)
			return nil
		default:
		}
	}
	b2 := false
	err := json.Unmarshal(b, &b2)
	if err == nil {
		*self = StatusBool(b2)
	}
	return err
}

func (self *StatusBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(*self)
}

type Status struct {
	Total      StatusInt `json:"total"`
	Successful StatusInt `json:"successful"`
	Failed     StatusInt `json:"failed"`
	Failures   []Failure `json:"failures,omitempty"`
}

type Failure struct {
	Index  string    `json:"index"`
	Shard  StatusInt `json:"shard"`
	Reason string    `json:"reason"`
}
