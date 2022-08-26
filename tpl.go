package main

const (
	tplFactoryOutline = `{comment}
package {package}

type factory func() interface{}

{spawnerFunc}

{poolPut}

var spawner = map[string]factory{
	{mapContext}
}

`

	tplPoolOutline = `{comment}
package {package}


import "sync"

var spawnerPools = map[string]*sync.Pool{
	{poolObject}
}
`

	// get build func without sync.pool
	tplSpawner = `
func Spawner(name string, newPool ...bool) (interface{},bool) {
	f ,ok := spawner[name]
	if !ok{
		return nil,ok
	}
	return f(),true
}`

	tplSpawnerPool = `
func Spawner(name string, newPool ...bool) (interface{}, bool) {
	if len(newPool) > 0 && newPool[0] {
		p, ok := spawnerPools[name]
		if !ok {
			return nil,false
		}
		return p.Get(),true
	}
	f, ok := spawner[name]
	if !ok {
		return nil, ok
	}
	return f(), true
}`

	tplPoolPut = `
func Put(name string, x interface{}){
	pool, ok := spawnerPools[name]
	if !ok {
		return
	}
	pool.Put(x)
}
`

	tplWithoutPoolPut = `
func Put(name string, x interface{}){}`

	tplMapField = `"{package}.{name}":func() interface{} { return &{name}{} },
`

	tplMapPoolField = `"{package}.{name}":{New: func() interface{}{ return spawner["{package}.{name}"]()}},
`
)
