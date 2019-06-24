package router

import (
	"testing"
)

func TestMatch(t *testing.T) {
	var pattern string
	var path string
	var v bool

	// pattern = "/+"
	// path = "/a"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/b/*"
	// path = "/a/b/c/d"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/b/*/d/+/f"
	// path = "/a/b/c/x/z/d/e"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	pattern = "/a/+/c"
	path = "/a/b/c"
	v = match(pattern, path)
	if !v {
		t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	}

	// pattern = "/"
	// path = "/"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/b"
	// path = "/a/b"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/+"
	// path = "/a/b"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/+"
	// path = "/a/b/b"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/+/c"
	// path = "/a/b/c"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/+/c"
	// path = "/a/b/c/d"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/+/+/c"
	// path = "/a/b/c"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/+/+/c"
	// path = "/a"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/*"
	// path = "/a"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/*"
	// path = "/a/b"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/*"
	// path = "/a/b/c/d/e"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/*/c/d"
	// path = "/a/b/e/c"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/*/c/+"
	// path = "/a/b/e/c"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/*/c/+"
	// path = "/a/b/e/c/d"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/a/*/c/+"
	// path = "/a/b/e/c/d/e"
	// v = match(pattern, path)
	// if v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

	// pattern = "/+/+/+"
	// path = "/a/b/e"
	// v = match(pattern, path)
	// if !v {
	// 	t.Errorf("ERROR! pattern: %v, path: %v\n", pattern, path)
	// }

}
