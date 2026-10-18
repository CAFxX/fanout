package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 uint }

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{42, 2}
    var v3 float32 = -2.25
    v4 := []uint32{588, 2147483648}
    v3 -= 1.5
    v5 := GenS1{F0: -521, F1: -1, F2: 771}
    v3 = 3.14159
    fmt.Println(v3, v1)
    var v6 int = 2147483647
    v3 -= float32(v1)
    switch v6 {
    case 9223372036854775807:
        v7 := GenS2{F0: (v6 >> ((385) & 63)), F1: -800, F2: 1}
        _ = v7
    case 42:
        v4 = append(v4, 65535)
    default:
        if false {
            fmt.Println(v6, v6)
            var v8 float32 = v3
            _ = len(v2)
            _ = v8
        } else {
            _ = len(v2)
        }
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(v3)
        v6 = 9223372036854775807
        _ = fi9
    }
    v6 = 35
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
