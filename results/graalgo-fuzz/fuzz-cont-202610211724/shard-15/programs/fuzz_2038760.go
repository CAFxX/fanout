package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int32; F2 bool }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 int8; F1 int32; F2 int16 }
type GenS8 struct { F0 int8; F1 int; F2 string }

func main() {
    var v1 int64 = -1
    var v2 int16 = -1
    var v3 float64 = 1.5
    v4 := make(map[string]bool)
    var v5 float64 = 1.5
    v6 := make(map[int]string)
    v7 := GenS1{F0: 812, F1: 990}
    v8 := []int{}
    switch "a" {
    case "foo bar":
        v2 -= 0
    default:
        v9 := GenS2{F0: -481}
        _ = v9
    }
    v3 += v3
    v10 := GenS3{F0: ((false && false) && false)}
    if false {
        v2 = (v2 << ((4294967295) & 63))
        v11 := GenS2{F0: 0}
        v12 := GenS4{F0: 2147483647}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS1{F0: fi13, F1: 1}
            switch ("" + "hello") {
            case "hello":
                var v15 uint64 = 2
                _ = v15
            default:
                v5 -= -2.25
            }
            _ = fi13
            _ = v14
        }
        if (2 != 0) {
            v16 := GenS1{F0: 797, F1: 164}
            if (!false) {
                v17 := GenS5{F0: -1, F1: 1, F2: false}
                fmt.Println(v2, v2, v2)
                if (!(!false)) {
                    if true {
                        fmt.Println(v1)
                        v18 := GenS6{F0: 127, F1: true}
                        var v19 int16 = (v2 + 990)
                        _ = v18
                        _ = v19
                    } else {
                        v5 = (((-2.25 * v3) * (0.1 + (0.0 - v5))) - v3)
                        var v20 uint64 = 302
                        v21 := v6[526]
                        _ = v20
                        _ = v21
                    }
                    var v22 uint8 = 1
                    v23 := GenS7{F0: 0, F1: -2147483648, F2: v2}
                    _ = v22
                    _ = v23
                }
                _ = v17
            }
            _ = v16
        } else {
            v1 = (v1 + (v1 + (v1 | (-(v1)))))
            v1 = (v1 & v1)
            fmt.Println(v2, v3)
        }
        var v24 uint32 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v1)
        _ = fi25
    }
    var v26 int = 1
    v27 := GenS8{F0: 83, F1: (v26 ^ v26), F2: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
