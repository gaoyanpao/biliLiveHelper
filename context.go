package biliLiveHelper

import (
	"github.com/bitly/go-simplejson"
	"math"
	"sync"
)

const (
	abortIndex = math.MaxInt8 / 2
)

type Context struct {
	Cmd       CmdType
	Msg       *simplejson.Json
	keys      map[string]interface{}
	keysMutex *sync.RWMutex
	handlers  HandleChain
	index     int8
}

func NewContext(cmdType CmdType, msg *simplejson.Json) *Context {
	return &Context{
		Cmd:       cmdType,
		Msg:       msg,
		index:     -1,
	}
}

func (c *Context) Set(key string, value interface{}) {
	if c.keysMutex == nil {
		c.keysMutex = &sync.RWMutex{}
	}

	c.keysMutex.Lock()
	if c.keys == nil {
		c.keys = make(map[string]interface{})
	}

	c.keys[key] = value
	c.keysMutex.Unlock()
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	if c.keysMutex == nil {
		c.keysMutex = &sync.RWMutex{}
	}

	c.keysMutex.RLock()
	value, exists = c.keys[key]
	c.keysMutex.RUnlock()
	return
}

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index].Handle(c)
		c.index++
	}
}


func (c *Context) Abort() {
	c.index = abortIndex
}


func (c *Context) IsAbort() bool {
	return c.index >= abortIndex
}