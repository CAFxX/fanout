package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
func genFn0(p0 uint64, p1 bool) uint16 {
    var v7 bool = p1
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 930
}

func main() {
    var v1 uint8 = 245
    v2 := make(map[int]uint32)
    var v3 int64 = 0
    var v4 int16 = 1000
    var v5 uint16 = 32767
    v6 := make(map[int]uint32)
    var v9 uint32 = 135
    if true {
        v10 := []int{}
        _ = v10
    }
    v1 = uint8(v5)
    var v11 uint64 = 2
    switch ("" + "hello") {
    case "hello":
        fmt.Println(v3)
    default:
        var v12 int16 = -950
        _ = v12
    }
    v13 := GenS1{F0: (uint(v5) <= 336), F1: "foo bar", F2: (v4 * 32767)}
    switch "a" {
    case "hello":
        v14 := v2[128]
        _ = v14
    }
    fmt.Println(v1, v4)
    switch "a" {
    case "\n\t":
        var v15 int64 = -9223372036854775808
        _ = v15
    case "foo bar":
        var v16 int16 = (v4 / (((-(v4))) | 1))
        _ = v16
    case "":
        for fi17 := 0; fi17 < 6; fi17++ {
            v5 = v5
            _ = fi17
        }
    default:
        v18 := GenS2{F0: "\n\t", F1: (v5 & 0), F2: uint(v11)}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
