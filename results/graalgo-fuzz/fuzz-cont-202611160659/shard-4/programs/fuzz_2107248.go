package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int32 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint8 }
func genFn0() uint16 {
    v4 := []int{0, 128}
    var v5 uint64 = 42
    v6 := GenS1{F0: 42, F1: ("foo bar" + "hello"), F2: 0}
    var v7 int16 = -32768
    fmt.Println(v7)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 703
}
func genFn1(p0 string, p1 int32) int {
    switch p0 {
    case "\n\t":
        v8 := GenS1{F0: -1, F1: (p0 + p0), F2: -2147483648}
        _ = v8
    case "foo bar":
        var v9 uint = 711
        _ = v9
    }
    var v10 uint8 = 42
    var v11 int32 = int32(v10)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -2147483648
}

func main() {
    var v1 bool = true
    var v2 int64 = 752
    var v3 int = 42
    v12 := GenS2{F0: 2, F1: 2147483647, F2: 128}
    v2 = v2
    var v13 float64 = ((1e10 - float64(1e10)) / 1.0)
    var v14 int32 = -863
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int32 = (v14 / ((v14) | 1))
        v13 += v13
        fmt.Println(fi15, v14)
        _ = fi15
        _ = v16
    }
    var v17 float64 = 0.1
    v17 = (v13 * v17)
    v18 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
