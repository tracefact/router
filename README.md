# router

对默认路由进行了一些扩展，使用“+”代表有且一个，“\*”位于中间时代表多个，当“\*”位于末尾时，代表没有或者多个。

使用方法：

```go
// 匹配 /a/b/c、/a/b/c/d、/a/b/ ...
router.HandleFunc("/a/b/*", handler)

// 匹配 /a/b/c、/a/d/c ...
router.HandleFunc("/a/+/c", handler)

// 匹配 /a/b/c、/a/b/d/e/c
router.HandleFunc("/a/*/c", handler)
```

更多用法，可以参见 router_test.go
