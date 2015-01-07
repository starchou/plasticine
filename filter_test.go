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

func Test_And(t *testing.T) {
	s := Filtered().Query(
		Term("name.first").Val("shay"),
	).Filter(And().And(
		Range("postDate").From("2010-03-01").To("2010-04-01"),
		Prefix("name.second").Val("ba"),
	))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"and, ":[{"range":{"postDate":{"from":"2010-03-01","to":"2010-04-01"}}},{"prefix":{"name.second":"ba"}}]},"query":{"term":{"name.first":"shay"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Query(
		Term("name.first").Val("shay"),
	).Filter(And().Filters(
		Range("postDate").From("2010-03-01").To("2010-04-01"),
		Prefix("name.second").Val("ba"),
	).Cache(true))
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"and":{"_cache":true,"filters":[{"range":{"postDate":{"from":"2010-03-01","to":"2010-04-01"}}},{"prefix":{"name.second":"ba"}}]}},"query":{"term":{"name.first":"shay"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_BoolFilter(t *testing.T) {
	s := Filtered().Query(
		QueryString().DefaultField("message").Query("elasticsearch"),
	).Filter(
		Bool().Must(
			Term("tag").Val("now"),
		).MustNot(
			Range("age").From(10).To(20),
		).Should(
			Term("tag").Val("sometag"),
			Term("tag").Val("sometagtag"),
		))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"bool":{"must":{"term":{"tag":"now"}},"must_not":{"range":{"age":{"from":10,"to":20}}},"should":[{"term":{"tag":"sometag"}},{"term":{"tag":"sometagtag"}}]}},"query":{"query_string":{"default_field":"message","query":"elasticsearch"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Exists(t *testing.T) {
	s := Filtered().Filter(
		Exists().Field("user"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"exists":{"field":"user"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_GeoBoundingBox(t *testing.T) {
	s := Filtered().Filter(
		GeoBoundingBox().TopLeft(map[string]float64{"lat": 40.01, "lon": -71.12}).BottomRight(
			map[string]float64{"lat": 40.01, "lon": -71.12}),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"geo_bounding_box":{"pin.location":{"bottom_right":{"lat":40.01,"lon":-71.12},"top_left":{"lat":40.01,"lon":-71.12}}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Filter(
		GeoBoundingBox().TopLeft([]float64{40.01, -71.12}).BottomRight([]float64{40.01, -71.12}),
	)
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"geo_bounding_box":{"pin.location":{"bottom_right":[40.01,-71.12],"top_left":[40.01,-71.12]}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Filter(
		GeoBoundingBox().TopLeft("40.73, -74.1").BottomRight("40.01, -71.12"),
	)
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"geo_bounding_box":{"pin.location":{"bottom_right":"40.01, -71.12","top_left":"40.73, -74.1"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_GeoDistance(t *testing.T) {
	s := Filtered().Filter(
		GeoDistance().PinLocation(map[string]float64{"lat": 40.01, "lon": -71.12}).Distance("200km"),
	).Query(MatchAll(nil))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"geo_distance":{"distance":"200km","pin.location":{"lat":40.01,"lon":-71.12}}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Filter(
		GeoDistance().PinLocation([]float64{40.01, -71.12}).Distance("200km"),
	).Query(MatchAll(nil))
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"geo_distance":{"distance":"200km","pin.location":[40.01,-71.12]}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Filter(
		GeoDistance().PinLocation("40.73, -74.1").Distance("200km"),
	).Query(MatchAll(nil))
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"geo_distance":{"distance":"200km","pin.location":"40.73, -74.1"}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_GeoDistanceRange(t *testing.T) {
	s := Filtered().Filter(
		GeoDistanceRange().PinLocation(map[string]float64{"lat": 40.01, "lon": -71.12}).From("200km").To("400km"),
	).Query(MatchAll(nil))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"geo_distance_range":{"from":"200km","pin.location":{"lat":40.01,"lon":-71.12},"to":"400km"}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_GeoPolygon(t *testing.T) {
	s := Filtered().Filter(
		GeoPolygon().Points(
			map[string]float64{"lat": 40.01, "lon": -71.12},
			map[string]float64{"lat": 40.01, "lon": -71.12},
		))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"geo_polygon":{"person.location":{"points":[{"lat":40.01,"lon":-71.12},{"lat":40.01,"lon":-71.12}]}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Filter(
		GeoPolygon().Points(
			[]float64{40.01, -71.12},
			[]float64{40.01, -71.12},
		))
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"geo_polygon":{"person.location":{"points":[[40.01,-71.12],[40.01,-71.12]]}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Filter(
		GeoPolygon().Points("drn5x1g8cu2y", "40.73, -74.1", "40.01, -71.12"),
	)
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"geo_polygon":{"person.location":{"points":["drn5x1g8cu2y","40.73, -74.1","40.01, -71.12"]}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_GeoShapeFilter(t *testing.T) {
	s := Filtered().Query(MatchAll(nil)).Filter(
		GeoShape().Type("envelope").Coordinates([]float64{13.400544, 52.530286}))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"geo_shape":{"location":{"shape":{"coordinates":[13.400544,52.530286],"type":"envelope"}}}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_GeohashCell(t *testing.T) {
	s := Filtered().Filter(
		GeohashCell().Lat(12.22).Lon(32.122).Neighbors(true).Precision(3),
	).Query(MatchAll(nil))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"geohash_cell":{"neighbors":true,"pin":{"lat":12.22,"lon":32.122},"precision":3}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_HasChildFilter(t *testing.T) {
	s := HasChild().Type("comment").MaxChildren(10).MinChildren(2).Filter(
		Term("user").Val("john"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"has_child":{"filter":{"term":{"user":"john"}},"max_children":10,"min_children":2,"type":"comment"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_HasParentFilter(t *testing.T) {
	s := HasParent().Type("blog").Filter(
		Term("tag").Val("something"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"has_parent":{"filter":{"term":{"tag":"something"}},"parent_type":"blog"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_IndicesFilter(t *testing.T) {
	s := Indices().Indices("index1", "index2").Filter(
		Term("tag").Val("wow"),
	).NoMatchFilter(
		Term("tag").Val("kow"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"indices":{"filter":{"term":{"tag":"wow"}},"indices":["index1","index2"],"no_match_filter":{"term":{"tag":"kow"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Limit(t *testing.T) {
	s := Filtered().Filter(
		Limit(100),
	).Query(Term("tag").Val("wow"))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"limit":{"value":100}},"query":{"term":{"tag":"wow"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_MatchAllFilter(t *testing.T) {
	s := Constant().Filter(
		MatchAll(nil),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Missing(t *testing.T) {
	s := Constant().Filter(
		Missing().Field("user").Existence(true).NullValue(false),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"missing":{"existence":true,"field":"user","null_value":false}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_NestedFilter(t *testing.T) {
	s := Filtered().Filter(
		Nested().Path("obj1").Filter(
			Bool().Must(
				Term("obj1.name").Val("blue"),
				Range("obj1.count").Gt(5),
			)).Cache(true)).Query(MatchAll(nil))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"nested":{"_cache":true,"filter":{"bool":{"must":[{"term":{"obj1.name":"blue"}},{"range":{"obj1.count":{"gt":5}}}]}},"path":"obj1"}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Not(t *testing.T) {
	s := Filtered().Query(
		Term("name.first").Val("shay"),
	).Filter(Not().Not(
		Range("postDate").From("2010-03-01").To("2010-04-01"),
	))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"not":{"range":{"postDate":{"from":"2010-03-01","to":"2010-04-01"}}}},"query":{"term":{"name.first":"shay"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Query(
		Term("name.first").Val("shay"),
	).Filter(Not().Filters(
		Range("postDate").From("2010-03-01").To("2010-04-01"),
	).Cache(true))
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"not":{"_cache":true,"filters":{"range":{"postDate":{"from":"2010-03-01","to":"2010-04-01"}}}}},"query":{"term":{"name.first":"shay"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Or(t *testing.T) {
	s := Filtered().Query(
		Term("name.first").Val("shay"),
	).Filter(Or().Or(
		Term("name.second").Val("banon"),
		Term("name.nick").Val("kimchy"),
	))
	b := s.Encode()
	got := string(b)
	expected := `{"filtered":{"filter":{"or":[{"term":{"name.second":"banon"}},{"term":{"name.nick":"kimchy"}}]},"query":{"term":{"name.first":"shay"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
	s = Filtered().Query(
		Term("name.first").Val("shay"),
	).Filter(Or().Filters(
		Term("name.second").Val("banon"),
		Term("name.nick").Val("kimchy"),
	).Cache(true))
	b = s.Encode()
	got = string(b)
	expected = `{"filtered":{"filter":{"or":{"_cache":true,"filters":[{"term":{"name.second":"banon"}},{"term":{"name.nick":"kimchy"}}]}},"query":{"term":{"name.first":"shay"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_PrefixFilter(t *testing.T) {
	s := Constant().Filter(
		Prefix("user").Val("ki").Cache(true),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"prefix":{"_cache":true,"user":"ki"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_QueryStringFilter(t *testing.T) {
	s := Constant().Filter(
		Query(
			QueryString().Query("this AND that OR thus"),
		))
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"query":{"query_string":{"query":"this AND that OR thus"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_RangeFilter(t *testing.T) {
	s := Constant().Filter(
		Range("age").Lte(10).Gte(20))
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"range":{"age":{"gte":20,"lte":10}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_RegexpFilter(t *testing.T) {
	s := Constant().Query(MatchAll(nil)).Filter(
		Regexp("name.first").Value("s.*y").Flags(
			"INTERSECTION|COMPLEMENT|EMPTY").Name("test").Cache(true).CacheKey("key"),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"regexp":{"_cache":true,"_cache_key":"key","_name":"test","name.first":{"flags":"INTERSECTION|COMPLEMENT|EMPTY","value":"s.*y"}}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_ScriptFilter(t *testing.T) {
	s := Constant().Query(MatchAll(nil)).Filter(
		Script().Script("doc['num1'].value > param1").Params("param1", 5),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"script":{"params":{"param1":5},"script":"doc['num1'].value \u003e param1"}},"query":{"match_all":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_TermFilter(t *testing.T) {
	s := Constant().Filter(
		Term("user").Val("kimchy").Cache(true))
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"term":{"_cache":true,"user":"kimchy"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_TermsFilter(t *testing.T) {
	s := Constant().Filter(
		Terms().Terms("user", "kimchy", "elasticsearch").Execution("bool").Cache(true))
	b := s.Encode()
	got := string(b)
	expected := `{"constant_score":{"filter":{"terms":{"_cache":true,"execution":"bool","user":["kimchy","elasticsearch"]}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_Type(t *testing.T) {
	s := Type("my_type")
	b, err := s.Encode()
	if err != nil {
		t.Error(err)
	}
	got := string(b)
	expected := `{"type":{"value":"my_type"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
