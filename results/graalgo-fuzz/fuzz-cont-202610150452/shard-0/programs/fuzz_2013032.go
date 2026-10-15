package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -128
    v3 := []int{9223372036854775807, 32767, 32767}
    v4 := []string{"a", "hello"}
    v2 = (-(v2))
    v2 = -128
    v2 -= (-(v2))
    v5 := GenS1{F0: 313, F1: "hello"}
    v2 = v2
    v1[2] = 42
    v2 -= v2
    v6 := GenS1{F0: 516, F1: ("foo bar" + ("foo bar" + ((("" + "") + "foo bar") + ("\n\t" + "hello"))))}
    v7 := v1[46]
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v8)
        var v10 uint8 = 2
        fmt.Println(v10, v8, v2)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int8 = (v2 - (v2 + v2))
        v7 -= v7
        var v13 uint32 = v8
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
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
