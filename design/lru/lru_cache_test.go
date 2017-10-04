package lru_test

import (
	"testing"
	"github.com/h4ckm03d/coding-practice/design/lru"
)

func TestLru2(t *testing.T) {
	c := lru.Constructor(2)
	c.Put(1,1)
	c.Put(2,2)
	if c.Get(1) == 1{
		t.Log("Success")
	}else{
		t.Error("should return 1")
	}

	c.Put(3,3)

	if c.Get(2) == -1{
		t.Log("Success")
	}else{
		t.Error("should return -1")
	}

	c.Put(4,4)

	if c.Get(1) == -1{
		t.Log("Success")
	}else{
		t.Error("should return -1")
	}

	if c.Get(3) == 3{
		t.Log("Success")
	}else{
		t.Error("should return 3")
	}

	if c.Get(4) == 4{
		t.Log("Success")
	}else{
		t.Error("should return 4")
	}
}