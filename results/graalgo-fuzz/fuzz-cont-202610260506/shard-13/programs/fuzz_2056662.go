package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 bool, p1 int8, p2 string) int64 {
    var v5 uint64 = 109
    p0 = false
    fmt.Println(p2)
    v5 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}

func main() {
    v1 := []string{"foo bar", "", "", ((((("foo bar" + (("a" + ("" + ("" + "foo bar"))) + "hello")) + "a") + "\n\t") + "foo bar") + "hello")}
    var v2 uint32 = 0
    var v3 uint16 = 1
    var v4 uint32 = v2
    fmt.Println(v4, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 -= (v2 - v4)
    switch -170 {
    case 127:
        var v6 int64 = -446
        _ = v6
    }
    v4 -= (^(v4))
    var v7 uint32 = (v2 & 2)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    if (42 < 100) {
        fmt.Println(v2)
        fmt.Println(v2)
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: "", F1: uint32(v3)}
        fmt.Println(v7)
        _ = fi8
        _ = v9
    }
    v7 += v4
    v3 = (v3 | v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
