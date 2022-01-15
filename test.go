package main

type st struct {
	bool
	id   int
	name string
	add  string
	v    *val
	v2   val
	arr  []int
	ma   map[string]*val
}

type val struct {
	v1 int
	v2 int
}
type sf int
type ssf = val2
type val2 struct {
	v1 int
	v2 int
}
