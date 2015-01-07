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

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-terms-facet.html
func Test_FacetsTerms(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().Terms("my_facet").Field("tag").Size(10).Order("term").Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"my_facet":{"terms":{"field":"tag","order":"term","size":10}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_FacetsTerms1(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().Terms("my_facet").Field("tag").Exclude("term1", "term2").Regex("_regex expression here_").RegexFlags(
			"DOTALL").Script("term == 'aaa' ? true : false").Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"my_facet":{"terms":{"exclude":["term1","term2"],"field":"tag","regex":"_regex expression here_","regex_flags":"DOTALL","script":"term == 'aaa' ? true : false"}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-range-facet.html
func Test_FacetsRange(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().Range("range1").Field("field_name").Ranges(
			To(15), FromTo(20, 70), FromTo(70, 120), From(120),
		).Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"range1":{"range":{"field":"field_name","ranges":[{"to":15},{"from":20,"to":70},{"from":70,"to":120},{"from":120}]}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-histogram-facet.html
func Test_FacetsHistogram(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().Histogram("histo1").KeyField("key_field_name").ValueField("value_field_name").Interval(100).KeyScript(
			"doc['date'].date.minuteOfHour * factor1").ValueScript("doc['num1'].value + factor2").Params(
			map[string]interface{}{"factor1": 2, "factor2": 3},
		).Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"histo1":{"histogram":{"interval":100,"key_field":"key_field_name","key_script":"doc['date'].date.minuteOfHour * factor1","params":{"factor1":2,"factor2":3},"value_field":"value_field_name","value_script":"doc['num1'].value + factor2"}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-date-histogram-facet.html
func Test_FacetsDateHistogram(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().DateHistogram("histo1").KeyField("timestamp").ValueScript("doc['price'].value * 2").Interval("day").Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"histo1":{"date_histogram":{"interval":"day","key_field":"timestamp","value_script":"doc['price'].value * 2"}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-filter-facet.html
func Test_FacetsFilter(t *testing.T) {
	s := Search().Facets(
		Facets().Filter("wow_facet", Term("tag").Val("wow")),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"wow_facet":{"filter":{"term":{"tag":"wow"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-query-facet.html
func Test_FacetsQuery(t *testing.T) {
	s := Search().Facets(
		Facets().Query("wow_facet", Term("tag").Val("wow")),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"wow_facet":{"query":{"term":{"tag":"wow"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-statistical-facet.html
func Test_FacetsStatistical(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().Statistical("stat1").Field("num1", "num2").script("(doc['num1'].value + doc['num2'].value) * factor").Params(
			map[string]interface{}{"factor": 5}).Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"stat1":{"statistical":{"field":["num1","num2"],"params":{"factor":5},"script":"(doc['num1'].value + doc['num2'].value) * factor"}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-terms-stats-facet.html
func Test_FacetsTermsStats(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().TermsStats("tag_price_stats").KeyField("tag").ValueField("price").Size(10).Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"tag_price_stats":{"terms_stats":{"key_field":"tag","size":10,"value_field":"price"}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-facets-geo-distance-facet.html
func Test_FacetsGeoDistance(t *testing.T) {
	s := Search().Query(MatchAll(nil)).Facets(
		Facets().GeoDistance("geo1").PinLocation(map[string]float64{"lat": 40.01, "lon": -71.12}).Ranges(
			To(10), FromTo(10, 20), FromTo(20, 100), From(100),
		).Facet(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"facets":{"geo1":{"terms_stats":{"pin.location":{"lat":40.01,"lon":-71.12},"ranges":[{"to":10},{"from":10,"to":20},{"from":20,"to":100},{"from":100}]}}},"query":{"match_all":{}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
