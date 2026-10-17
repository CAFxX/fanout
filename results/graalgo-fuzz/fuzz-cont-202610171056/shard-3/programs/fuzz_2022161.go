package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int32 }
type GenS3 struct { F0 int; F1 uint; F2 bool }

func main() {
    var v1 uint16 = 255
    var v2 int32 = 0
    var v3 float64 = (0.1 * 0.0)
    var v4 int64 = 0
    v5 := make(map[uint32]uint32)
    v6 := []int{}
    var v7 float64 = v3
    v7 = (v7 * (v7 + v3))
    fmt.Println(v4, v4)
    for fi8 := 0; fi8 < 19; fi8++ {
        v4 += (v4 + v4)
        _ = fi8
    }
    v2 = 1
    v9 := GenS1{F0: ("a" + "hello"), F1: uint64(v1)}
    switch "hello" {
    case "a":
        v5[716] = 813
    case "hello":
        fmt.Println(v1)
    }
    var v10 uint64 = 1
    var v11 float64 = (v3 - v7)
    switch "hello" {
    case "hello":
        v10 = v10
    default:
        fmt.Println(v2, v10)
    }
    delete(v5, 65535)
    v10 -= 0
    v12 := GenS2{F0: 2, F1: false, F2: (v2 - int32(v10))}
    v10 += (v10 | 18446744073709551615)
    var v13 int = 2147483647
    fmt.Println(v3, v2)
    if (0.0 != (v11 - (v7 * v3))) {
        v1 = 0
        delete(v5, 0)
        v1 += 255
    } else {
        if true {
            switch "" {
            case "foo bar":
                var v14 int16 = 1
                _ = v14
            default:
                if (-1 < (v13 / ((v13) | 1))) {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        fmt.Println(v13)
                        var v16 uint64 = v10
                        delete(v5, 42)
                        _ = fi15
                        _ = v16
                    }
                    v13 = (v13 & v13)
                    var v17 float32 = 0.0
                    _ = v17
                }
            }
            var v18 uint16 = v1
            _ = v18
        }
    }
    v19 := GenS3{F0: v13, F1: 119, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
