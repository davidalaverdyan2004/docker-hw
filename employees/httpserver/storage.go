package httpserver

import "sync"

var employees = []Employee{}
var nextID = 1
var mtx sync.Mutex
