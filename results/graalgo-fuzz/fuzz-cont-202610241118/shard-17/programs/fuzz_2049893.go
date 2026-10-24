package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int64, p1 int32) uint16 {
    p0 -= (p0 ^ p0)
    p0 -= p0
    _ = p0
    _ = p1
    return 425
}

func main() {
    var v1 int8 = 0
    var v2 int8 = (v1 | 100)
    v3 := make(map[string]int)
    v4 := GenS1{F0: genFn0(1, -1), F1: 0}
    v5 := []int{-9223372036854775808}
    if (!(!true)) {
        v3["hello"] = 42
    }
    v2 -= v1
    var v6 uint8 = 18
    v3[("" + ((("a" + "") + "a") + "foo bar"))] = 4294967295
    v7 := GenS1{F0: 1, F1: 1}
    var v8 uint32 = 922
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
