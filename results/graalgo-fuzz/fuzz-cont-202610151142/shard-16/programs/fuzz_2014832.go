package main

import (
    "fmt"
)

func genFn0() int64 {
    var v4 float32 = -2.25
    var v5 int32 = 42
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 255
    var v3 bool = false
    var v6 uint64 = 810
    v6 = 1000
    v7 := []int{-744, -1, 322, 32767}
    var v8 uint32 = 2147483648
    v2 = v2
    if (!v3) {
        v8 -= (^(v8))
    }
    v6 += v6
    v1[-9223372036854775808] = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
