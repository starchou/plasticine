
## API Status

Here's the current API status.

### APIs

- [x] Search (most queries, filters, facets, aggregations etc. are implemented: see below)
- [x] Index
- [x] Get
- [x] Delete
- [x] Delete By Query
- [x] Update
- [ ] Multi Get
- [ ] Bulk
- [ ] Bulk UDP
- [ ] Term vectors
- [ ] Multi term vectors
- [ ] Count
- [ ] Validate
- [ ] Explain
- [x] Search
- [ ] Search shards
- [ ] Search template
- [x] Facets (most are implemented, see below)
- [x] Aggregates (most are implemented, see below)
- [x] Multi Search
- [ ] Percolate
- [ ] More like this
- [ ] Benchmark

### Indices

- [x] Create index
- [x] Delete index
- [ ] Indices exists
- [ ] Open/close index
- [ ] Put mapping
- [ ] Get mapping
- [ ] Get field mapping
- [x] Types exist
- [ ] Delete mapping
- [ ] Index aliases
- [ ] Update indices settings
- [ ] Get settings
- [ ] Analyze
- [ ] Index templates
- [ ] Warmers
- [x] Status
- [ ] Indices stats
- [ ] Indices segments
- [ ] Indices recovery
- [x] Clear cache
- [x] Flush
- [x] Refresh
- [x] Optimize

### Snapshot and Restore

- [ ] Snapshot
- [ ] Restore
- [ ] Snapshot status
- [ ] Monitoring snapshot/restore progress
- [ ] Partial restore

### Cat APIs

Not implemented. Those are better suited for operating with Elasticsearch
on the command line.

### Cluster

- [x] Health
- [x] State
- [x] Stats
- [x] Pending cluster tasks
- [x] Cluster reroute
- [x] Cluster update settings
- [x] Nodes stats
- [x] Nodes info
- [ ] Nodes hot_threads
- [x] Nodes shutdown

### Query DSL

#### Queries

- [x] `match`
- [x] `multi_match`
- [x] `bool`
- [x] `boosting`
- [x] `common_terms`
- [x] `constant_score`
- [x] `dis_max`
- [x] `filtered`
- [x] `fuzzy_like_this_query` (`flt`)
- [x] `fuzzy_like_this_field_query` (`flt_field`)
- [x] `function_score`
- [x] `fuzzy`
- [x] `geo_shape`
- [x] `has_child`
- [x] `has_parent`
- [x] `ids`
- [x] `indices`
- [x] `match_all`
- [x] `mlt`
- [x] `mlt_field`
- [x] `nested`
- [x] `prefix`
- [x] `query_string`
- [x] `simple_query_string`
- [x] `range`
- [x] `regexp`
- [x] `span_first`
- [x] `span_multi_term`
- [x] `span_near`
- [x] `span_not`
- [x] `span_or`
- [x] `span_term`
- [x] `term`
- [x] `terms`
- [x] `top_children`
- [x] `wildcard`
- [x] `minimum_should_match`
- [x] `multi_term_query_rewrite`
- [ ] `template_query`

#### Filters

- [x] `and`
- [x] `bool`
- [x] `exists`
- [x] `geo_bounding_box`
- [x] `geo_distance`
- [x] `geo_distance_range`
- [x] `geo_polygon`
- [x] `geoshape`
- [x] `geohash`
- [x] `has_child`
- [x] `has_parent`
- [x] `ids`
- [x] `indices`
- [x] `limit`
- [x] `match_all`
- [x] `missing`
- [x] `nested`
- [x] `not`
- [x] `or`
- [x] `prefix`
- [x] `query`
- [x] `range`
- [x] `regexp`
- [x] `script`
- [x] `term`
- [x] `terms`
- [x] `type`

### Facets

- [x] Terms
- [x] Range
- [x] Histogram
- [x] Date Histogram
- [x] Filter
- [x] Query
- [x] Statistical
- [x] Terms Stats
- [x] Geo Distance

### Aggregations

- [x] min
- [x] max
- [x] sum
- [x] avg
- [x] stats
- [x] extended stats
- [x] value count
- [x] percentiles
- [x] percentile ranks
- [x] cardinality
- [x] geo bounds
- [ ] top hits
- [x] scripted metric
- [x] global
- [x] filter
- [ ] filters
- [x] missing
- [x] nested
- [ ] reverse nested
- [x] children
- [x] terms
- [x] significant terms
- [x] range
- [x] date range
- [x] ipv4 range
- [x] histogram
- [x] date histogram
- [x] geo distance
- [ ] geohash grid

### Scan

Scrolling through documents (e.g. `search_type=scan`) are implemented via
the `Scroll` and `Scan` services.