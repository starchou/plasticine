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

func Test_Include(t *testing.T) {
	s := Include("obj1.*", "obj2.*")
	b, _ := s.Encode()
	got := string(b)
	expected := `{"include":["obj1.*","obj2.*"]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Include("obj1.*")
	b, _ = s.Encode()
	got = string(b)
	expected = `{"include":"obj1.*"}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Exclude(t *testing.T) {
	s := Exclude("obj1.*", "obj2.*")
	b, _ := s.Encode()
	got := string(b)
	expected := `{"exclude":["obj1.*","obj2.*"]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Exclude("obj1.*")
	b, _ = s.Encode()
	got = string(b)
	expected = `{"exclude":"obj1.*"}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Source1(t *testing.T) {
	s := Search().Source(true)
	b := s.Encode()
	got := string(b)
	expected := `{"_source":true}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Source2(t *testing.T) {
	s := Search().Source([]string{"obj1.*", "obj2.*"})
	b := s.Encode()
	got := string(b)
	expected := `{"_source":["obj1.*","obj2.*"]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Source3(t *testing.T) {
	s := Search().Source(
		Include("obj1.*", "obj2.*"),
		Exclude("*.description"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"_source":{"exclude":"*.description","include":["obj1.*","obj2.*"]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Field(t *testing.T) {
	s := Search().Fields()
	b := s.Encode()
	got := string(b)
	expected := `{"fields":[]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Search().Fields("user", "postDate")
	b = s.Encode()
	got = string(b)
	expected = `{"fields":["user","postDate"]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_PartialFields(t *testing.T) {
	s := Search().PartialFields("partial1", Include("obj1.*"))
	b := s.Encode()
	got := string(b)
	expected := `{"partial_fields":{"partial1":{"include":"obj1.*"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Search().PartialFields("partial1", Include("obj1.*"), Exclude("obj1.*", "obj2.*"))
	b = s.Encode()
	got = string(b)
	expected = `{"partial_fields":{"partial1":{"exclude":["obj1.*","obj2.*"],"include":"obj1.*"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_ScriptFields(t *testing.T) {
	s := Search().ScriptFields("test1",
		Script().Script("doc['my_field_name'].value * 2").Map(),
	).ScriptFields("test2",
		Script().Script("doc['my_field_name'].value  * factor").Params("factor", 2.0).Map(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"script_fields":{"test1":{"script":"doc['my_field_name'].value * 2"},"test2":{"params":{"factor":2},"script":"doc['my_field_name'].value  * factor"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Search().ScriptFields("test1",
		Script().Script("_source.obj1.obj2").Map(),
	)
	b = s.Encode()
	got = string(b)
	expected = `{"script_fields":{"test1":{"script":"_source.obj1.obj2"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_FielddataFields(t *testing.T) {
	s := Search().FielddataFields("test1", "test2")
	b := s.Encode()
	got := string(b)
	expected := `{"fielddata_fields":["test1","test2"]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_PostFilter(t *testing.T) {
	s := Search().Query(
		Filtered().Filter(
			Bool().Must(
				Term("color").Val("red"),
				Term("brand").Val("gucci"),
			),
		),
	).Aggs(
		Aggregation().Term("colors", "color"),
		Aggregation().Filter("color_red", Term("color").Val("red")),
		Aggregation().Term("models", "model"),
	).PostFilter(
		Term("color").Val("red"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"color_red":{"filter":{"term":{"color":"red"}}},"colors":{"term":{"field":"color"}},"models":{"term":{"field":"model"}}},"post_filter":{"term":{"color":"red"}},"query":{"filtered":{"filter":{"bool":{"must":[{"term":{"color":"red"}},{"term":{"brand":"gucci"}}]}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Highlight(t *testing.T) {
	s := Search().Highlight(
		Highlight("content").Type("plain").ForceSource(true).Tags("span").HighlightQuery(
			Bool().Must(Match("content").Query("foo bar")).Should(
				MatchPhrase("content").Query("foo bar").PhraseSlop(1).Boost(1.2),
			).MinimumShouldMatch(0),
		))
	b := s.Encode()
	got := string(b)
	expected := `{"highlight":{"fields":{"content":{"force_source":true,"highlight_query":{"bool":{"minimum_should_match":0,"must":{"match":{"content":{"query":"foo bar"}}},"should":{"match_phrase":{"content":{"boost":1.2,"phrase_slop":1,"query":"foo bar"}}}}},"type":"plain"}},"post_tags":["\u003c/span\u003e"],"pre_tags":["\u003cspan\u003e"]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_Rescoring(t *testing.T) {
	s := Search().Rescoring(
		Rescoring().WindowSize(50).QueryWeight(1.2).RescoreQueryWeight(3.2).RescoreQuery(
			Match("field1").Query("the quick brown").Type("phrase").Slop(2),
		),
		Rescoring().WindowSize(50).QueryWeight(1.2).RescoreQueryWeight(3.2).RescoreQuery(
			Match("field1").Query("the quick brown").Type("phrase").Slop(2),
		))
	b := s.Encode()
	got := string(b)
	expected := `{"rescore":[{"query":{"query_weight":1.2,"rescore_query":{"match":{"field1":{"query":"the quick brown","slop":2,"type":"phrase"}}},"rescore_query_weight":3.2},"window_size":50},{"query":{"query_weight":1.2,"rescore_query":{"match":{"field1":{"query":"the quick brown","slop":2,"type":"phrase"}}},"rescore_query_weight":3.2},"window_size":50}]}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_MultiSort(t *testing.T) {
	s := Sort(
		Order("date", "desc"),
		Desc("time"),
		"_score",
	)
	b, err := s.Encode()
	if err != nil {
		t.Error(err)
	} else {
		got := string(b)
		expected := `{"sort":[{"date":{"order":"desc"}},{"time":"desc"},"_score"]}`
		if got != expected {
			t.Errorf("expected\n%s\n,got:\n%s", expected, got)
		}
	}
}

func Test_Sort(t *testing.T) {
	s := Sort(
		Desc("date"),
		Asc("time"),
	)
	b, err := s.Encode()
	if err != nil {
		t.Error(err)
	} else {
		got := string(b)
		expected := `{"sort":[{"date":"desc"},{"time":"asc"}]}`
		if got != expected {
			t.Errorf("expected\n%s\n,got:\n%s", expected, got)
		}
	}
}
func Test_Desc(t *testing.T) {
	s := Sort(Desc("date"))
	b, err := s.Encode()
	if err != nil {
		t.Error(err)
	} else {
		got := string(b)
		expected := `{"sort":{"date":"desc"}}`
		if got != expected {
			t.Errorf("expected\n%s\n,got:\n%s", expected, got)
		}
	}
}
func Test_Asc(t *testing.T) {
	s := Sort(Asc("date"))
	b, err := s.Encode()
	if err != nil {
		t.Error(err)
	} else {
		got := string(b)
		expected := `{"sort":{"date":"asc"}}`
		if got != expected {
			t.Errorf("expected\n%s\n,got:\n%s", expected, got)
		}
	}
}

func Test_Search(t *testing.T) {
	s := Search().From(10).Size(20).Version(true).Explain(false).MinScore(0.45).Query(
		Filtered().Query(Match("content").Query("aaa")),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"explain":false,"from":10,"min_score":0.45,"query":{"filtered":{"query":{"match":{"content":{"query":"aaa"}}}}},"size":20,"version":true}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
