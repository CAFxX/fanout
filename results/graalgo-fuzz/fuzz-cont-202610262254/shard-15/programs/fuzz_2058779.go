package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0(p0 int64, p1 uint, p2 int16) int32 {
    var v4 uint8 = 42
    v4 = v4
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 942
}
func genFn1(p0 uint8) uint64 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 9; fi5++ {
        if true {
            fmt.Println(p0)
        }
        fmt.Println(fi5, p0)
        fmt.Println(p0, fi5)
        _ = fi5
    }
    v6 := GenS1{F0: 53}
    _ = p0
    _ = v6
    return 497
}

func main() {
    var v1 int64 = 42
    var v2 float64 = 0.1
    var v3 uint16 = 1
    switch "" {
    case "hello":
        v7 := []int{930, 9223372036854775807, -1}
        _ = v7
    default:
        var v8 uint32 = uint32(v3)
        _ = v8
    }
    v9 := GenS1{F0: -128}
    v2 = (3.14159 + float64(0.1))
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int16 = -513
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 42
    v13 := []int{-565, 251, 32768, 42}
    v12 += (v12 >> ((42) & 63))
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v14 int8 = 42
    fmt.Println(v12)
    if true {
        v15 := GenS2{F0: genFn1(2), F1: 44}
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 float32 = (1e10 * (3.14159 / 1.0))
            fmt.Println(fi16, v1)
            _ = fi16
            _ = v17
        }
        var v18 uint8 = 128
        _ = v15
        _ = v18
    }
    switch 42 {
    case 1:
        var v19 int8 = (v14 >> ((65535) & 63))
        _ = v19
    default:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
