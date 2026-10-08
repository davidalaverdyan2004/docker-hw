package httpserver

import (
	"employees/db"
	"sync"
)

var employees = []db.Employee{}
var nextID = 1
var mtx sync.Mutex
