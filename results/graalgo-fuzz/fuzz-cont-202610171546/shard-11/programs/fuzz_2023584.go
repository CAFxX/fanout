package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint8 = 2
    var v2 float64 = float64(float32(1e10))
    var v3 float32 = 3.14159
    v4 := make(map[int]uint32)
    var v5 int = 1
    v6 := GenS1{F0: 2, F1: 42, F2: v1}
    var v7 bool = (true && (false || (!(42 <= 18446744073709551615))))
    if v7 {
        v3 = (v3 + v3)
        fmt.Println(v3)
        v7 = v7
    }
    var v8 uint64 = 18446744073709551615
    _ = len(v4)
    if (0 > -50) {
        fmt.Println(v3)
    }
    v9 := GenS2{F0: (v7 && false)}
    v5 = (-(v5))
    if (v7 || false) {
        var v10 uint = uint(v1)
        _ = v10
    } else {
        switch ("hello" + "hello") {
        case "":
            v11 := v4[(v5 - v5)]
            _ = v11
        case "a":
            v12 := GenS2{F0: (false || v7)}
            _ = v12
        default:
            for fi13 := 0; fi13 < 5; fi13++ {
                v3 = v3
                var v14 int16 = -1
                _ = fi13
                _ = v14
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
