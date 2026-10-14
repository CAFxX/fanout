package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint }
func genFn0(p0 bool) uint16 {
    p0 = false
    p0 = (p0 || p0)
    var v5 int32 = 2147483647
    _ = p0
    _ = v5
    return 65535
}
func genFn1() int16 {
    switch "a" {
    case "":
        var v6 float32 = ((0.0 + float32(float64(3.14159))) * 0.1)
        _ = v6
    case "a":
        var v7 int = 155
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: 42, F1: ("" + "foo bar"), F2: (356 > 474)}
        _ = v8
    }
    var v9 uint32 = 2
    _ = v9
    return -1
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]bool)
    var v3 float64 = 1.5
    var v4 bool = false
    v10 := GenS2{F0: 0, F1: (65535 > 661), F2: 1}
    if (-566 == -1) {
        fmt.Println(v3, v4, v3)
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int8 = 127
            _ = fi11
            _ = v12
        }
        var v13 int32 = -1
        _ = v13
    }
    _ = len(v2)
    var v14 float32 = 0.0
    v15 := []int{-2147483648, (v1 - v1), v1}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v1, fi16)
        _ = fi16
    }
    v4 = v4
    var v17 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
