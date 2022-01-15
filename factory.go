// Code generated by "spawner "; DO NOT EDIT.

package main

type factory func() interface{}

func Spawner(name string) (interface{}, bool) {
	f, ok := spawner[name]
	if !ok {
		return nil, ok
	}
	return f(), true
}

var spawner = map[string]factory{
	"st":   func() interface{} { return &st{} },
	"val":  func() interface{} { return &val{} },
	"Dal2": func() interface{} { return &Dal2{} },
}
