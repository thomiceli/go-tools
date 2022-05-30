package tparamsource

// https://staticcheck.io/issues/1282

import (
	tparamsink "local-type-param-sink"
	"testing"
)

func TestFoo(t *testing.T) { //@ used("TestFoo", true), used("t", true)
	type EmptyStruct struct{} //@ used("EmptyStruct", true)
	_ = tparamsink.TypeOfType[EmptyStruct]()
}
