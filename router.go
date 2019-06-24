package router

import (
	"fmt"
	"net/http"
	"strings"
)

// 保存用户的 pattern 和 对应的handlerFunc
var handlers = map[string]http.HandlerFunc{}

// 保存用户pattern，主要是为了按顺序，因为map没有顺序
var patterns = []string{}

// HandleFunc 处理请求
func HandleFunc(pattern string, handler http.HandlerFunc) {

	if !checkPattern(pattern) {
		panic("pattern格式有误，检查是否存在下面模式：/+/*、/*/+、/*/*，可使用/*代替")
	}

	var filter = func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		for _, k := range patterns {
			if match(k, path) {
				handlers[k](w, r)
				return
			}
		}

		// 全都不匹配则404
		http.NotFound(w, r)
	}

	if _, ok := handlers[pattern]; !ok {
		// 如果不存在，在list中加一个
		patterns = append(patterns, pattern)
	} else {
		panic("URL模式 " + pattern + " 重复注册了！")
	}

	handlers[pattern] = handler

	p2 := transform(pattern)
	http.HandleFunc(p2, filter)
}

// 用户传进来的模式
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
		fmt.Printf("pattern:[%v]%v, \tpath:[%v]", i, x, pathIndex)

		if pathIndex >= pathLen {
			// pattern还没结束，但是path已经结束了
			fmt.Printf("nil \n")
			return false
		}
		fmt.Printf("%v \n", pathArr[pathIndex])

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
			// *出现在最后一个，则肯定满足
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

	// arr2 没有匹配完，则肯定不满足
	if pathIndex != (pathLen - 1) {
		return false
	}

	return true
}

// 将形如 /a/b/+/c 的，转换为 /a/b/
func transform(pattern string) string {
	if pattern == "/" {
		return pattern
	}

	plusIndex := strings.Index(pattern, "+")
	starIndex := strings.Index(pattern, "*")

	// 不存在模式，则直接退出
	if plusIndex == -1 && starIndex == -1 {
		return pattern
	}

	index := 0

	// 取靠前面的作为起点
	if plusIndex != -1 && starIndex == -1 {
		index = plusIndex
	}
	if plusIndex == -1 && starIndex != -1 {
		index = starIndex
	}
	if plusIndex != -1 && starIndex != -1 {
		if plusIndex < starIndex {
			index = plusIndex
		} else {
			index = starIndex
		}
	}

	if index >= 0 {
		slashIndex := strings.LastIndex(pattern[:index+1], "/")
		pattern = pattern[:slashIndex+1]
	}

	return pattern
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
