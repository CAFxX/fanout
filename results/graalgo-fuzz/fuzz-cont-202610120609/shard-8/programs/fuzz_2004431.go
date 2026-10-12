package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint8 }

func main() {
    var v1 uint16 = 42
    var v2 int64 = 42
    var v3 float32 = (-2.25 + (1.5 / 0.1))
    var v4 uint64 = 4294967295
    var v5 uint64 = v4
    v6 := make(map[string]uint32)
    v4 = (v4 - v4)
    v1 -= 2
    _ = len(v6)
    var v7 float64 = (3.14159 / 1.5)
    switch -937 {
    case -849:
        fmt.Println(v1, v4)
    case 13:
        var v8 int8 = 1
        _ = v8
    case 816:
        v1 -= (^(v1))
    }
    switch "" {
    case "":
        _ = len(v6)
    case "foo bar":
        v9 := GenS1{F0: 32767}
        _ = v9
    case "hello":
        var v10 float64 = (float64(v3) - 3.14159)
        _ = v10
    default:
        if true {
            v4 = v5
            v1 -= 65535
        } else {
            var v11 uint64 = (v4 % ((v5) | 1))
            _ = v11
        }
    }
    fmt.Println(v1)
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: 32767}
        _ = fi12
        _ = v13
    }
    v1 = (v1 ^ uint16(v4))
    v14 := GenS2{F0: 1, F1: -448, F2: 42}
    if false {
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v3)
            v16 := GenS2{F0: 609, F1: 42, F2: 128}
            v17 := []int{(fi15 + fi15), (fi15 + fi15), fi15, int(v2)}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v18 := []int{922}
    var v19 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
