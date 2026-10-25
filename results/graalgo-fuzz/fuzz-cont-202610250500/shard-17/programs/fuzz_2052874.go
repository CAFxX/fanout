package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 int16) int8 {
    var v5 int = 32768
    var v6 int8 = 100
    _ = p0
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -128
    var v3 bool = false
    var v4 int16 = 858
    v4 -= v4
    var v7 int64 = 42
    v3 = v3
    if (368 > 220) {
        var v8 int16 = v4
        v9 := []int{}
        _ = v8
        _ = v9
    }
    if true {
        v3 = (-1 >= -659)
        v4 = v4
        v10 := GenS1{F0: (-(v2)), F1: (v4 | v4), F2: ""}
        _ = v10
    }
    fmt.Println(v7)
    v11 := GenS1{F0: (-(v2)), F1: v4, F2: ("" + "a")}
    v12 := GenS1{F0: v2, F1: v4, F2: "foo bar"}
    fmt.Println(v4, v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
