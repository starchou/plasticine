package main

import (
	"strconv"

	es "github.com/starchou/plasticine"
)

type User struct {
	Id   int
	Name string
}

func main() {
	s := es.NewServer("localhost", "9200")
	//s.Debug(true)
	user := &User{1, "star"}
	//Insert
	s.Index("plasticine", "user", strconv.Itoa(user.Id), nil, user)
	//Update
	user.Name = "chou"
	s.Index("plasticine", "user", strconv.Itoa(user.Id), nil, user)

	//Delete
	s.Delete("plasticine", "user", "1", nil)

	//Search DSL
	// es.Search().Index("plasticine").Type("News").From(1).Size(10).Query(
	// 	es.QueryString().Query("what is plasticine").Fields("Title", "Content"),
	// ).Filter(
	// 	es.And().And(
	// 		es.Range("Created").From("2012-12-10T15:00:00-08:00").To("2013-5-10T15:00:00-08:00"),
	// 		es.Terms().Terms("IsDelete", "0"),
	// 	),
	// ).Result(s)

	//Search DSL (facet)
	// es.Search().Query(es.MatchAll(nil)).Facets(
	// 	es.Facets().Terms("my_facet").Field("tag").Size(10).Order("term").Facet(),
	// ).Result(s)
}
