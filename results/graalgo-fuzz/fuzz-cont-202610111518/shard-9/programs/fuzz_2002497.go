package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }

func main() {
    v1 := []int{-636, -9223372036854775808, -2147483648, 32768}
    v2 := make(map[string]int)
    var v3 uint = 198
    var v4 float64 = 1e10
    var v5 int16 = 32767
    v6 := GenS1{F0: 2}
    var v7 int16 = (v5 & v5)
    v8 := GenS2{F0: "foo bar"}
    if (!false) {
        v9 := GenS1{F0: v3}
        var v10 int = -2147483648
        _ = v10
        _ = v9
    }
    v5 = v5
    v5 += (^(v5))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
