package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0() int32 {
    v7 := []int{548, 128, 127}
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = 65
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint64 = 1
        _ = fi10
        _ = v11
    }
    _ = v7
    return -139
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int32 = 93
    v4 := []int{}
    v5 := make(map[int]uint32)
    var v6 uint32 = 1
    v3 += (^(v3))
    var v12 uint32 = (v6 - v6)
    v6 = 65535
    v6 -= (v6 >> ((255) & 63))
    var v13 bool = (!((^(v6)) > v12))
    var v14 uint64 = 4294967295
    fmt.Println(v12)
    v15 := GenS1{F0: 0}
    if v13 {
        var v16 uint8 = 130
        _ = v16
    } else {
        if false {
            _ = len(v2)
            if ("a" == "") {
                var v17 int8 = 100
                v6 = v6
                _ = v17
            } else {
                v14 = 264
            }
            switch "" {
            case "\n\t":
                var v18 float64 = 0.0
                _ = v18
            case "foo bar":
                fmt.Println(v13, v12, v12)
            case "":
                var v19 string = "a"
                _ = v19
            }
        } else {
            fmt.Println(v12, v14)
            v20 := GenS2{F0: (false || false), F1: 274}
            v21 := GenS1{F0: 4294967295}
            _ = v20
            _ = v21
        }
    }
    var v22 bool = (-1 >= -86)
    for fi23 := 0; fi23 < 17; fi23++ {
        v14 -= 1
        v3 = (-(v3))
        _ = fi23
    }
    v24 := GenS3{F0: (v12 - uint32(v14)), F1: -889, F2: (9223372036854775807 > 1)}
    fmt.Println(v6)
    v12 -= 42
    var v25 uint64 = 42
    fmt.Println(v14)
    v22 = (false || v22)
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS4{F0: -47, F1: 132}
        var v28 uint64 = (v14 ^ v14)
        fmt.Println(v25, v14)
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
