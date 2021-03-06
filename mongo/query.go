package mongo

import (
	"launchpad.net/mgo/bson"
	"github.com/ungerik/go-start/model"
	"launchpad.net/mgo"
)

///////////////////////////////////////////////////////////////////////////////
// Query

/*
Query is the interface with all query methods for mongo.
*/
type Query interface {
	subDocumentSelector() string
	bsonSelector() bson.M
	mongoQuery() (q *mgo.Query, err error)

	Selector() string

	ParentQuery() Query
	Collection() *Collection

	SubDocument(selector string) Query
	Skip(int) Query
	Limit(int) Query
	Sort(selector string) Query                               // Chain Sort() and SortReverse() for multi value sorting
	SortReverse(selector string) Query                        // Chain Sort() and SortReverse() for multi value sorting
	SortFunc(less func(a, b interface{}) bool) model.Iterator // Last query of chain

	// FilterX must be the first query on a Collection
	IsFilter() bool
	Filter(selector string, value interface{}) Query
	FilterWhere(javascript string) Query

	// Filter via a Go function. Note that all documents have to be loaded
	// in memory in order for Go code to be able to filter it.
	FilterFunc(passFilter model.FilterFunc) model.Iterator
	FilterRef(selector string, ref ...Ref) Query
	FilterEqualCaseInsensitive(selector string, str string) Query
	FilterNotEqual(selector string, value interface{}) Query
	FilterLess(selector string, value interface{}) Query
	FilterGreater(selector string, value interface{}) Query
	FilterLessEqual(selector string, value interface{}) Query
	FilterGreaterEqual(selector string, value interface{}) Query
	FilterModulo(selector string, divisor, result interface{}) Query
	FilterIn(selector string, values ...interface{}) Query
	FilterNotIn(selector string, values ...interface{}) Query
	FilterAllIn(selector string, values ...interface{}) Query
	FilterArraySize(selector string, size int) Query
	FilterStartsWith(selector string, str string) Query
	FilterStartsWithCaseInsensitive(selector string, str string) Query
	FilterEndsWith(selector string, str string) Query
	FilterEndsWithCaseInsensitive(selector string, str string) Query
	FilterContains(selector string, str string) Query
	FilterContainsCaseInsensitive(selector string, str string) Query
	FilterExists(selector string, exists bool) Query

	Or() Query

	// Statistics
	Count() (n int, err error)
	// Distinct() int
	Explain() string

	// Read
	One() (document interface{}, err error)
	TryOne() (document interface{}, found bool, err error)
	GetOrCreateOne() (document interface{}, found bool, err error)

	Iterator() model.Iterator
	OneID() (id bson.ObjectId, err error)
	TryOneID() (id bson.ObjectId, found bool, err error)
	IDs() (ids []bson.ObjectId, err error)
	Refs() (refs []Ref, err error)

	// RemoveAll ignores Skip() and Limit()
	RemoveAll() error
}
