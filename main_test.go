package main

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	t.Error(errors.New("test"))
}
