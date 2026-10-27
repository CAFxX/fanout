package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }

func main() {
    v1 := []string{"hello", "hello", "", ("foo bar" + "hello")}
    var v2 string = "hello"
    v3 := make(map[int]int)
    var v4 uint32 = 2147483648
    var v5 int8 = -128
    v6 := []int{-355}
    v5 = v5
    v5 += v5
    v2 = "a"
    fmt.Println(v2)
    v7 := GenS1{F0: 65535, F1: (("foo bar" + ("" + "hello")) + v2)}
    for fi8 := 0; fi8 < 8; fi8++ {
        v3[fi8] = fi8
        var v9 int64 = -400
        _ = fi8
        _ = v9
    }
    v10 := v3[128]
    v10 = (v10 - (v10 >> ((1) & 63)))
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int8 = v5
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 718, F1: v2}
    var v14 int = v10
    if (((520 >= 0) || false) && true) {
        fmt.Println(v5, v14)
        var v15 float64 = ((((-2.25 / float64(-2.25)) - 0.0) / 3.14159) / 1.5)
        _ = v15
    }
    var v16 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
