package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0() int32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint16 = 32767
        _ = fi4
        _ = v5
    }
    var v6 int32 = 42
    var v7 int16 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        v7 -= v7
        for fi9 := 0; fi9 < 20; fi9++ {
            v6 = v6
            v10 := GenS1{F0: -1, F1: 571, F2: "\n\t"}
            fmt.Println(v6, fi9)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v7 = 1000
    _ = v6
    _ = v7
    return (v6 + -951)
}
func genFn1(p0 bool) uint8 {
    var v11 int8 = 100
    p0 = p0
    fmt.Println(v11)
    fmt.Println(v11, v11)
    _ = p0
    _ = v11
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 float64 = (0.1 / float64(0.1))
    switch (("\n\t" + "hello") + (("\n\t" + "hello") + "a")) {
    case "foo bar":
        fmt.Println(v3)
    case "a":
        var v12 int16 = 42
        _ = v12
    }
    var v13 float64 = -2.25
    v14 := []int{-215, -9223372036854775808, -522, 585}
    v15 := GenS2{F0: 45, F1: (-1 > 0)}
    var v16 float32 = ((1.5 / 1.0) / (float32((v3 + (3.14159 * v13))) / 0.1))
    var v17 int = -560
    var v18 int64 = -146
    if false {
        var v19 bool = true
        var v20 bool = v19
        _ = v19
        _ = v20
    } else {
        v21 := GenS1{F0: -2147483648, F1: 260, F2: "a"}
        v16 -= float32((0.1 - (1e10 + v3)))
        var v22 uint16 = 0
        _ = v21
        _ = v22
    }
    var v23 uint = 42
    if true {
        var v24 int8 = 100
        fmt.Println(v18, v23, v3)
        _ = v24
    }
    var v25 bool = ((false && true) || (86 >= 2147483647))
    _ = len(v2)
    v18 -= 42
    var v26 float32 = -2.25
    switch ("" + "hello") {
    case "\n\t":
        var v27 uint = v23
        _ = v27
    default:
        v16 += v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
}
