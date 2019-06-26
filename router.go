package router

import (
	"net/http"
	"strings"
)

// MyRouter 路由
type myRouter struct {
	// 保存用户的 pattern 和 对应的handlerFunc
	handlers map[string]http.HandlerFunc

	// 保存用户pattern，主要是为了按顺序，因为map没有顺序
	patterns []string
}

// DefaultRouter 路由
var DefaultRouter = &myRouter{handlers: map[string]http.HandlerFunc{}, patterns: []string{}}

// ServeHTTP
func (x *myRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	for _, k := range x.patterns {
		if match(k, path) {
			x.handlers[k](w, r)
			return
		}
	}

	// 全都不匹配则404
	http.NotFound(w, r)
}

func (x *myRouter) HandleFunc(pattern string, handler http.HandlerFunc) {
	if !checkPattern(pattern) {
		panic("pattern格式有误，检查是否存在下面模式：/+/*、/*/+、/*/*，可使用/*代替")
	}

	if _, ok := x.handlers[pattern]; !ok {
		// 如果不存在，在list中加一个
		x.patterns = append(x.patterns, pattern)
	} else {
		panic("URL模式 " + pattern + " 重复注册了！")
	}

	x.handlers[pattern] = handler
}

// 验证模式和路径是否匹配
func match(pattern string, path string) bool {
	if pattern == path {
		return true
	}

	if path+"/" == pattern {
		return true
	}

	// patternArr 由pattern生成的array
	// pathArr 由path生成的array
	patternArr := strings.Split(pattern, "/")
	pathArr := strings.Split(path, "/")

	patternLen := len(patternArr)
	pathLen := len(pathArr)
	pathIndex := 0

	// 路径比模式还短，肯定有问题
	if pathLen < patternLen {
		return false
	}

	for i, x := range patternArr {
		//fmt.Printf("pattern:[%v]%v, \tpath:[%v]", i, x, pathIndex)

		if pathIndex >= pathLen {
			// pattern还没结束，但是path已经结束了
			//fmt.Printf("nil \n")
			return false
		}
		//fmt.Printf("%v \n", pathArr[pathIndex])

		if x == pathArr[pathIndex] {
			// 进入下一次循环，pathIndex才可以加
			if i != (patternLen - 1) {
				pathIndex++
				// fmt.Printf("pathIndex++, %v \n", pathIndex)
			}
			continue
		}

		if x == "+" {
			// fmt.Printf("%v, %v\n", arr1Len, arr2Len)
			// 如果是最后一个，则不需要循环了，否则直接退出了
			if i == (patternLen-1) && patternLen == pathLen {
				// + 号不可以匹配空
				if pathArr[pathIndex] == "" {
					return false
				}
				return true
			}

			// 进入下一次循环，pathIndex才可以加
			if i != (patternLen - 1) {
				pathIndex++
			}
			continue
		}

		if x == "*" {
			// *出现在最后一个，则认为其满足(即使 path最后一项为空 也认为满足)
			if i == (patternLen - 1) {
				return true
			}

			// 下一个要匹配的项
			next := patternArr[i+1]
			for c := pathIndex; c < pathLen; c++ {
				if pathArr[c] == next {
					pathIndex = c
				}
			}
			continue
		}

		return false
	}

	// path 没有匹配完，则肯定不满足
	if pathIndex != (pathLen - 1) {
		return false
	}

	return true
}

// 验证模式，不能出现 /+/* 或者 /*/+ ，或者 /*/*
// 因为它们都等于 /*/
func checkPattern(pattern string) bool {
	arr := strings.Split(pattern, "/")
	pre := "this_is_a_place_holder"

	for _, x := range arr {
		if x == "*" && pre == "+" {
			return false
		}
		if x == "+" && pre == "*" {
			return false
		}
		if x == "*" && pre == "*" {
			return false
		}
		pre = x
	}

	return true
}
