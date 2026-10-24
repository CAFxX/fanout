package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }

func main() {
    var v1 uint32 = 89
    var v2 float64 = (-2.25 * 1e10)
    var v3 uint32 = v1
    var v4 uint16 = 254
    v5 := make(map[int]string)
    var v6 float32 = (float32(v2) * ((0.1 * float32(v2)) + 1e10))
    v7 := v5[-424]
    v8 := []int{}
    var v9 uint16 = (v4 | v4)
    v1 -= 65535
    var v10 uint8 = 175
    _ = len(v8)
    var v11 bool = (v6 != v6)
    v12 := GenS1{F0: (v1 + v3), F1: (("hello" + (v7 + ("a" + v7))) + (v7 + "hello")), F2: v11}
    v13 := GenS1{F0: (v3 ^ v1), F1: "hello", F2: v11}
    if len(v8) > 0 {
        v8[0] = int(v3)
    }
    fmt.Println(v7, v1)
    v14 := GenS1{F0: uint32(v4), F1: (("\n\t" + v7) + ((v7 + v7) + v7)), F2: false}
    if false {
        var v15 float64 = (((1e10 * 0.1) + v2) + float64(v6))
        _ = v15
    }
    if v11 {
        fmt.Println(v9, v9)
        switch 937 {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 7; fi16++ {
                v2 = (v2 / v2)
                _ = fi16
            }
        case 9223372036854775807:
            v9 = v4
        case 1:
            var v17 int = 1000
            _ = v17
        default:
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        delete(v5, -1)
    }
    var v18 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
