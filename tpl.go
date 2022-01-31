package main

const tplFactoryOutline = `{comment}
package {package}

type factory func() interface{}

{spawnerFunc}

{poolPut}

var spawner = map[string]factory{
	{mapContext}
}

`

const tplPoolOutline = `{comment}
package {package}


import "sync"

var spawnerPools = map[string]*sync.Pool{
	{poolObject}
}
`

const tplSpawner = `
func Spawner(name string, newPool ...bool) (interface{},bool) {
	f ,ok := spawner[name]
	if !ok{
		return nil,ok
	}
	return f(),true
}`

const tplSpawnerPool = `
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

const tplPoolPut = `
func Put(name string, x interface{}){
	pool, ok := spawnerPools[name]
	if !ok {
		return
	}
	pool.Put(x)
}`
const tplWithoutPoolPut = `
func Put(name string, x interface{}){}`

const tplMapField = `"{package}.{name}":func() interface{} { return &{name}{} },
`

const tplMapPoolField = `"{package}.{name}":{New: func() interface{}{ return spawner["{package}.{name}"]()}},
`
