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

import "testing"

func Test_Match(t *testing.T) {
	Debug = false
	s := Match("name").Query("star").Operator("and")
	b := s.Encode()
	got := string(b)
	expected := `{"match":{"name":{"operator":"and","query":"star"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_MatchPhrase(t *testing.T) {
	s := MatchPhrase("message").Value("star")
	b := s.Encode()
	got := string(b)
	expected := `{"match_phrase":{"message":"star"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_MatchPhrasePrefix(t *testing.T) {
	s := MatchPhrasePrefix("message").Value("star")
	b := s.Encode()
	got := string(b)
	expected := `{"match_phrase_prefix":{"message":"star"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Bool(t *testing.T) {
	s := Bool().Must(
		Term("user").Val("star"),
	).MustNot(
		Range("age").Lt(50).Gte(30).Boost(2.0),
	).Should(
		Term("user").Val("star"),
		Term("user").Val("chou"),
	).Must(Bool().Should(
		Term("h").Val("v"),
	).MinimumShouldMatch(1))
	b := s.Encode()
	got := string(b)
	expected := `{"bool":{"must":{"bool":{"minimum_should_match":1,"should":{"term":{"h":"v"}}}},"must_not":{"range":{"age":{"boost":2,"gte":30,"lt":50}}},"should":[{"term":{"user":"star"}},{"term":{"user":"chou"}}]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}

}
func Test_Booting(t *testing.T) {
	//Debug = true
	s := Boosting().Boost(0.1).Negative(
		Term("user").Val("star"),
	).Positive(
		Term("user").Val("star"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"boosting":{"negative":{"term":{"user":"star"}},"negative_boost":0.1,"positive":{"term":{"user":"star"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_MultiMatch(t *testing.T) {
	s := MultiMatch().Field("1", "2").Query("star").Type("cross_fields").Operator("and")
	b := s.Encode()
	got := string(b)
	expected := `{"multi_match":{"fields":["1","2"],"operator":"and","query":"star","type":"cross_fields"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}

}
func Test_Common(t *testing.T) {
	s := Common().Query("nelly the elephant as a cartoon").Frequency(0.001).Operator(
		"and").MinimumShouldMatch(map[string]int{"low_freq": 3, "high_freq": 7})
	b := s.Encode()
	got := string(b)
	expected := `{"common":{"body":{"cutoff_frequency":0.001,"low_freq_operator":"and","minimum_should_match":{"high_freq":7,"low_freq":3},"query":"nelly the elephant as a cartoon"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}

}
func Test_Constant(t *testing.T) {
	s := Constant().Query(
		Term("user").Val("star"),
	).Filter(
		Term("user").Val("chou"),
	).Boost(1.2)
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"boost":1.2,"filter":{"term":{"user":"chou"}},"query":{"term":{"user":"star"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_DisjunctionMax(t *testing.T) {
	s := DisjunctionMax().Query(
		Term("user").Val("star"),
		Term("user").Val("chou"),
	).TieBreaker(0.1).Boost(1.2)
	b := s.Encode()
	got := string(b)
	expected := `{"dis_max":{"boost":1.2,"queries":[{"term":{"user":"star"}},{"term":{"user":"chou"}}],"tie_breaker":0.1}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = DisjunctionMax().Query(
		Match("user").Value("star"),
		Match("user").Value("chou"),
	).TieBreaker(0.1).Boost(1.2)
	b = s.Encode()
	got = string(b)
	expected = `{"dis_max":{"boost":1.2,"queries":[{"match":{"user":"star"}},{"match":{"user":"chou"}}],"tie_breaker":0.1}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Filtered(t *testing.T) {
	s := Filtered().Query(
		Match("user").Value("star"),
		Match("user").Value("chou"),
	).Filter(
		Range("created").Gte("now - 1d / d"),
	).Strategy("leap_frog")
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"range":{"created":{"gte":"now - 1d / d"}}},"query":[{"match":{"user":"star"}},{"match":{"user":"chou"}}],"strategy":"leap_frog"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Flt(t *testing.T) {
	s := Flt().Fields("name.first", "name.last").LikeText("text like this one").MaxQueryTerms(12)
	b := s.Encode()
	got := string(b)
	expected := `{"fuzzy_like_this":{"fields":["name.first","name.last"],"like_text":"text like this one","max_query_terms":12}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_FltField(t *testing.T) {
	s := FltField("name.first").LikeText("text like this one").MaxQueryTerms(12)
	b := s.Encode()
	got := string(b)
	expected := `{"fuzzy_like_this_field":{"name.first":{"like_text":"text like this one","max_query_terms":12}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Fuzzy(t *testing.T) {
	s := Fuzzy("name").Fuzzy("star")
	b := s.Encode()
	got := string(b)
	expected := `{"fuzzy":{"name":"star"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Fuzzy("name").Value("star").Fuzziness(2).PrefixLength(1).Boost(1.0).MaxExpansions(50)
	b = s.Encode()
	got = string(b)
	expected = `{"fuzzy":{"name":{"boost":1,"fuzziness":2,"max_expansions":50,"prefix_length":1,"value":"star"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_GeoShape(t *testing.T) {
	s := GeoShape().Type("envelope").Coordinates([2][2]int{{13, 53}, {14, 52}})
	b := s.Encode()
	got := string(b)
	expected := `{"geo_shape":{"location":{"shape":{"coordinates":[[13,53],[14,52]],"type":"envelope"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_HasChild(t *testing.T) {
	s := HasChild().Type("blog_tag").ScoreMode("sum").MaxChildren(10).MinChildren(1).Query(
		Term("tag").Val("something"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"has_child":{"max_children":10,"min_children":1,"query":{"term":{"tag":"something"}},"score_mode":"sum","type":"blog_tag"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_HasParent(t *testing.T) {
	s := HasParent().Type("blog").ScoreMode("score").Query(
		Term("tag").Val("something"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"has_parent":{"parent_type":"blog","query":{"term":{"tag":"something"}},"score_mode":"score"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Ids(t *testing.T) {
	s := Ids().Type("blog", "my_type").Values("1", "2", "3")
	b := s.Encode()
	got := string(b)
	expected := `{"ids":{"type":["blog","my_type"],"values":["1","2","3"]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Indices(t *testing.T) {
	s := Indices().Indices("index1", "index2").Query(
		Term("tag").Val("wow"),
	).NoMatchQuery(
		Term("tag").Val("wow"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"indices":{"indices":["index1","index2"],"no_match_query":{"term":{"tag":"wow"}},"query":{"term":{"tag":"wow"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_MatchAll(t *testing.T) {
	s := MatchAll(Boost(1.2))
	b, err := s.Encode()
	if err != nil {
		t.Error(err)
	}
	got := string(b)
	expected := `{"match_all":{"boost":1.2}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = MatchAll(nil)
	b, err = s.Encode()
	if err != nil {
		t.Error(err)
	}
	got = string(b)
	expected = `{"match_all":{}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_MoreLikeThis(t *testing.T) {
	s := MoreLikeThis().Fields("name.first", "name.last").Docs(
		Doc("yn", "Info", "2"),
		Doc("sc", "Lemma", "1"),
	).Ids("1", "2").MinTermFreq(2).MaxQueryTerms(12)
	b := s.Encode()
	got := string(b)
	expected := `{"more_like_this":{"docs":[{"_id":"2","_index":"yn","_type":"Info"},{"_id":"1","_index":"sc","_type":"Lemma"}],"fields":["name.first","name.last"],"ids":["1","2"],"max_query_terms":12,"min_term_freq":2}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_MoreLikeThisField(t *testing.T) {
	s := MoreLikeThisField("name.first").LikeText("text like this one").MinTermFreq(2).MaxQueryTerms(12)
	b := s.Encode()
	got := string(b)
	expected := `{"more_like_this_field":{"name.first":{"like_text":"text like this one","max_query_terms":12,"min_term_freq":2}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Nested(t *testing.T) {
	s := Nested().Path("obj1").ScoreMode("none").Query(
		Bool().Must(
			Match("obj1.name").Value("blue"),
			Range("obj1.count").Gt(5),
		))
	b := s.Encode()
	got := string(b)
	expected := `{"nested":{"path":"obj1","query":{"bool":{"must":[{"match":{"obj1.name":"blue"}},{"range":{"obj1.count":{"gt":5}}}]}},"score_mode":"none"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Prefix(t *testing.T) {
	s := Prefix("user").Val("v")
	b := s.Encode()
	got := string(b)
	expected := `{"prefix":{"user":"v"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Prefix("user").Value("v").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"prefix":{"user":{"boost":2,"value":"v"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Prefix("user").Prefix("v").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"prefix":{"user":{"boost":2,"prefix":"v"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_QueryString(t *testing.T) {
	s := QueryString().DefaultField("content").Query("this AND that OR thus")
	b := s.Encode()
	got := string(b)
	expected := `{"query_string":{"default_field":"content","query":"this AND that OR thus"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Regexp(t *testing.T) {
	s := Regexp("name.first").Val("s.*y")
	b := s.Encode()
	got := string(b)
	expected := `{"regexp":{"name.first":"s.*y"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Regexp("name.first").Value("s.*y").Flags("INTERSECTION|COMPLEMENT|EMPTY").Boost(1.2)
	b = s.Encode()
	got = string(b)
	expected = `{"regexp":{"name.first":{"boost":1.2,"flags":"INTERSECTION|COMPLEMENT|EMPTY","value":"s.*y"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_SpanFirst(t *testing.T) {
	s := SpanFirst().Match("user", "star").End(3)
	b := s.Encode()
	got := string(b)
	expected := `{"span_first":{"end":3,"match":{"span_term":{"user":"star"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_SpanMulti(t *testing.T) {
	s := SpanMulti().Match(
		Prefix("user").Value("ki"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"span_multi":{"match":{"prefix":{"user":{"value":"ki"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = SpanMulti().Match(
		Prefix("user").Value("ki").Boost(1.2),
	)
	b = s.Encode()
	got = string(b)
	expected = `{"span_multi":{"match":{"prefix":{"user":{"boost":1.2,"value":"ki"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = SpanMulti().Match(
		Regexp("user").Value("ki").Boost(1.2),
	)
	b = s.Encode()
	got = string(b)
	expected = `{"span_multi":{"match":{"regexp":{"user":{"boost":1.2,"value":"ki"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_SpanNear(t *testing.T) {
	s := SpanNear().Clauses(
		SpanTerm("user").Val("value1"),
		SpanTerm("user").Val("value2"),
		SpanTerm("user").Val("value3"),
	).Slop(12).CollectPayloads(false).InOrder(false)
	b := s.Encode()
	got := string(b)
	expected := `{"span_near":{"clauses":[{"span_term":{"user":"value1"}},{"span_term":{"user":"value2"}},{"span_term":{"user":"value3"}}],"collect_payloads":false,"in_order":false,"slop":12}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_SpanNot(t *testing.T) {
	s := SpanNot().Include(
		SpanTerm("field1").Val("hoya"),
	).Exclude(
		SpanNear().Clauses(
			SpanTerm("user").Val("hoya"),
			SpanTerm("user").Val("la"),
		).Slop(12).CollectPayloads(false).InOrder(false),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"span_not":{"exclude":{"span_near":{"clauses":[{"span_term":{"user":"hoya"}},{"span_term":{"user":"la"}}],"collect_payloads":false,"in_order":false,"slop":12}},"include":{"span_term":{"field1":"hoya"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_SpanOr(t *testing.T) {
	s := SpanOr().Clauses(
		SpanTerm("user").Val("value1"),
		SpanTerm("user").Val("value2"),
		SpanTerm("user").Val("value3"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"span_or":{"clauses":[{"span_term":{"user":"value1"}},{"span_term":{"user":"value2"}},{"span_term":{"user":"value3"}}]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_SpanTerm(t *testing.T) {
	s := SpanTerm("user").Val("v")
	b := s.Encode()
	got := string(b)
	expected := `{"span_term":{"user":"v"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = SpanTerm("user").Value("v").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"span_term":{"user":{"boost":2,"value":"v"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = SpanTerm("user").Term("v").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"span_term":{"user":{"boost":2,"term":"v"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Term(t *testing.T) {
	s := Term("user").Val("v")
	b := s.Encode()
	got := string(b)
	expected := `{"term":{"user":"v"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Term("user").Value("v").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"term":{"user":{"boost":2,"value":"v"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Term("user").Term("v").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"term":{"user":{"boost":2,"term":"v"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Terms(t *testing.T) {
	s := Terms().Tags("blue", "pill").MinimumShouldMatch(2)
	b := s.Encode()
	got := string(b)
	expected := `{"terms":{"minimum_should_match":2,"tags":["blue","pill"]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_TopChildren(t *testing.T) {
	s := TopChildren().Type("blog_tag").Query(
		Term("tag").Val("something")).Factor(5).IncrementalFactor(3).Scope("my_scope").Score("max")
	b := s.Encode()
	got := string(b)
	expected := `{"top_children":{"_scope":"my_scope","factor":5,"incremental_factor":3,"query":{"term":{"tag":"something"}},"score":"max","type":"blog_tag"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Wildcard(t *testing.T) {
	s := Wildcard("user").Val("ki*y")
	b := s.Encode()
	got := string(b)
	expected := `{"wildcard":{"user":"ki*y"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Wildcard("user").Value("ki*y").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"wildcard":{"user":{"boost":2,"value":"ki*y"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Wildcard("user").Wildcard("ki*y").Boost(2.0)
	b = s.Encode()
	got = string(b)
	expected = `{"wildcard":{"user":{"boost":2,"wildcard":"ki*y"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Template(t *testing.T) {
	s := Template().QueryMap("match_{{template}}", struct{}{}).Params("template", "all")
	b := s.Encode()
	got := string(b)
	expected := `{"template":{"params":{"template":"all"},"query":{"match_{{template}}":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
