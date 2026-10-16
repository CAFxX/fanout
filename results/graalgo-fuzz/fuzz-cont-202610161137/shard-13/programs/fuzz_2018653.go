package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16; F1 string; F2 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 2
    v3 := make(map[int]string)
    v4 := GenS1{F0: 26, F1: v2}
    fmt.Println(v1, v1)
    if (-1 <= 2147483647) {
        v5 := []int{128, 32768}
        fmt.Println(v2, v2)
        _ = len(v3)
        _ = v5
    }
    if true {
        var v6 int8 = 100
        v6 = v6
        _ = v6
    } else {
        var v7 int16 = -32768
        v8 := GenS2{F0: 255, F1: "\n\t", F2: (^(v7))}
        v7 = (v7 & -32768)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v2 = v2
        _ = fi9
    }
    if (84 <= 0) {
        fmt.Println(v2, v1)
    } else {
        var v10 int64 = (v1 / ((v1) | 1))
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int64 = v1
        _ = fi11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint16 = 65535
            v15 := []int{(fi13 | (fi13 >> ((v2) & 63))), 0, fi13, (fi13 ^ fi13)}
            var v16 float64 = 0.0
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        v1 += -419
        v2 = (v2 * 42)
    }
    var v17 int = 2147483647
    var v18 uint16 = 65535
    v17 += v17
    v17 -= (v17 * v17)
    v18 -= (^(v18))
    switch "" {
    case "foo bar":
        for fi19 := 0; fi19 < 1; fi19++ {
            v1 -= 98
            v18 -= (v18 * (^(v18)))
            delete(v3, (v17 & (-(fi19))))
            _ = fi19
        }
    case "a":
        v18 = 0
    case "hello":
        var v20 bool = (!(1 != 2147483648))
        _ = v20
    }
    v21 := []int{1}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v17)
        _ = fi22
    }
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
