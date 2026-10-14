package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
func genFn0(p0 int16, p1 int32) int16 {
    var v6 float32 = (-2.25 + 0.1)
    var v7 int8 = int8(p1)
    v6 += v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p0 - (p0 >> ((4294967295) & 63)))
}

func main() {
    var v1 float64 = 0.0
    var v2 int64 = -1
    var v3 int32 = 0
    var v4 bool = true
    var v5 float64 = v1
    v4 = (!(43 <= 49))
    v8 := []int{127, -229}
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: "foo bar", F1: (^(v2))}
        _ = fi9
        _ = v10
    }
    var v11 bool = v4
    if v11 {
        v5 = v1
    }
    var v12 bool = (42 <= genFn0(900, v3))
    v4 = (0.1 <= 0.1)
    v2 = (v2 * v2)
    v11 = (((9 > 100) || v11) && v4)
    v13 := GenS1{F0: "hello", F1: -740}
    v1 -= (float64(0.1) / 1.0)
    v14 := GenS1{F0: "foo bar", F1: v2}
    v1 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
