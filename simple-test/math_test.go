package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 9 // given or arrange

	actual := Add(4, 5) // when or act

	if actual != expected { //then or assert
		t.Errorf("expected %d is not equal to actual %d", expected, actual)
	}
}
