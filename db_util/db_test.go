package dbutil

import "testing"

func TestDbInfo(t *testing.T) {
	Init("root:121511yml@tcp(127.0.0.1:3306)/test_db")
	GetTableStruct("hello_world")
}
