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

func Test_AggregationMin(t *testing.T) {
	s := Aggregation()
	s.Min("min_price").Field("price").Script("_value * correction").Params("correction", 1.2)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"min_price":{"min":{"field":"price","params":{"correction":1.2},"script":"_value * correction"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationMax(t *testing.T) {
	s := Aggregation()
	s.Max("max_price").Field("price").Script(
		"_value * correction").Params("correction", 1.3)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"max_price":{"max":{"field":"price","params":{"correction":1.3},"script":"_value * correction"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationSum(t *testing.T) {
	s := Aggregation()
	s.Sum("daytime_return").Field("change").Script("_value * correction").Params("correction", 1.3)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"daytime_return":{"sum":{"field":"change","params":{"correction":1.3},"script":"_value * correction"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationAvg(t *testing.T) {
	s := Aggregation()
	s.Avg("avg_corrected_grade").Field("grade").Script("_value * correction").Params("correction", 3.2)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"avg_corrected_grade":{"avg":{"field":"grade","params":{"correction":3.2},"script":"_value * correction"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationStats(t *testing.T) {
	s := Aggregation()
	s.Stats("stats_corrected_grade").Field("grade").Script("_value * correction").Params("correction", 3.2)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"stats_corrected_grade":{"stats":{"field":"grade","params":{"correction":3.2},"script":"_value * correction"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationExtendedStats(t *testing.T) {
	s := Aggregation()
	s.ExtendedStats("grades_stats").Field("grade").Script("_value * correction").Params("correction", 3.2)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"grades_stats":{"extended_stats":{"field":"grade","params":{"correction":3.2},"script":"_value * correction"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_AggregationValueCount(t *testing.T) {
	s := Aggregation()
	s.ValueCount("grades_count").Field("grade").Script("_value * correction")
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"grades_count":{"value_count":{"field":"grade","script":"_value * correction"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationPercentiles(t *testing.T) {
	s := Aggregation()
	s.Percentiles("load_time_outlier").Field("load_time").Params("timeUnit", 2000).Percents(
		95, 99, 99.9,
	).Compression(200).Script("doc['load_time'].value / timeUnit")
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"load_time_outlier":{"percentiles":{"compression":200,"field":"load_time","params":{"timeUnit":2000},"percents":[95,99,99.9],"script":"doc['load_time'].value / timeUnit"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationPercentileRanks(t *testing.T) {
	s := Aggregation()
	s.PercentileRanks("load_time_outlier").Field("load_time").Value(
		3, 5,
	).Script("doc['load_time'].value / timeUnit")
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"load_time_outlier":{"percentile_ranks":{"field":"load_time","script":"doc['load_time'].value / timeUnit","value":[3,5]}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationCardinality(t *testing.T) {
	s := Aggregation()
	s.Cardinality("author_count").Field("author").PrecisionThreshold(100).Rehash(false).Script("doc['author.first_name'].value + ' ' + doc['author.last_name'].value")
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"author_count":{"cardinality":{"field":"author","precision_threshold":100,"rehash":false,"script":"doc['author.first_name'].value + ' ' + doc['author.last_name'].value"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
func Test_AggregationGeoBounds(t *testing.T) {
	s := Aggregation()
	s.GeoBounds("viewport").Field("location").WrapLongitude("true")
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"viewport":{"geo_bounds":{"field":"location","wrap_longitude":"true"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationScriptedMetric(t *testing.T) {
	s := Aggregation()
	s.ScriptedMetric("profit").InitScript("_agg['transactions'] = []").MapScript(
		"if (doc['type'].value == \"sale\") { _agg.transactions.add(doc['amount'].value) } else { _agg.transactions.add(-1 * doc['amount'].value) }",
	).CombineScript("profit = 0; for (t in _agg.transactions) { profit += t }; return profit").ReduceScript("profit = 0; for (a in _aggs) { profit += a }; return profit")
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"profit":{"scripted_metric":{"combine_script":"profit = 0; for (t in _agg.transactions) { profit += t }; return profit","init_script":"_agg['transactions'] = []","map_script":"if (doc['type'].value == \"sale\") { _agg.transactions.add(doc['amount'].value) } else { _agg.transactions.add(-1 * doc['amount'].value) }","reduce_script":"profit = 0; for (a in _aggs) { profit += a }; return profit"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationGlobal(t *testing.T) {
	s := Aggregation().Global("all_products")
	s.Aggregation("all_products",
		Aggregation().Avg("avg_price").Field("price").Aggs(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"all_products":{"aggs":{"avg_price":{"avg":{"field":"price"}}},"global":{}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func Test_AggregationFilter(t *testing.T) {
	s := Aggregation().Filter("in_stock_products",
		Range("stock").Lt(0),
	)
	s.Aggregation("in_stock_products",
		Aggregation().Avg("avg_price").Field("price").Aggs(),
	)
	b := s.Encode()
	got := string(b)
	expected := `{"aggs":{"in_stock_products":{"aggs":{"avg_price":{"avg":{"field":"price"}}},"filter":{"range":{"stock":{"lt":0}}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

//http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-aggregations-bucket-filters-aggregation.html
func Test_AggregationFilters(t *testing.T) {
	// s := Aggregation().Filters("messages", v)
	// b := s.Encode()
	// got := string(b)
	// expected := `{"aggs":{"in_stock_products":{"aggs":{"avg_price":{"avg":{"field":"price"}}},"filter":{"range":{"stock":{"lt":0}}}}}}`
	// if got != expected {
	// 	t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	// }
}
