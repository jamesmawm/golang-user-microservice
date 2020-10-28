package data

import (
	"fmt"
	"testing"
	"reflect"
)

/**
Test to ensure that a default database is generated
*/
func TestNewDatabase(t *testing.T) {
	db := newDatabase()

	fmt.Print("type of db: ", reflect.TypeOf(db))
	if reflect.TypeOf(db).String() != "*gorm.DB" {
		t.Errorf("newDatabase didn't return a default database, instead got %s", reflect.TypeOf(db).String())
	}
}
