package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{2147483647}
    var v3 int16 = -32768
    v4 := make(map[int]string)
    var v5 int32 = 2147483647
    var v6 uint = 255
    v6 -= 2
    var v7 bool = (!("" < "a"))
    v8 := GenS1{F0: 127, F1: v3, F2: 42}
    for fi9 := 0; fi9 < 15; fi9++ {
        if ((1.5 <= -2.25) && v7) {
            v6 = (v6 % ((v6) | 1))
        }
        _ = fi9
    }
    v4[122] = "hello"
    v2 = append(v2, -9223372036854775808)
    fmt.Println(v3, v6, v6)
    v10 := GenS2{F0: false}
    switch "" {
    case "a":
        v7 = true
    case "\n\t":
        v3 = v3
    case "hello":
        _ = len(v4)
    default:
        switch "" {
        case "\n\t":
            v11 := GenS1{F0: 128, F1: v3, F2: 42}
            _ = v11
        default:
            var v12 int = -9223372036854775808
            _ = v12
        }
    }
    if (v5 == v5) {
        var v13 float32 = (3.14159 * 0.0)
        v14 := GenS3{F0: (v5 | v5)}
        fmt.Println(v3, v6, v3)
        _ = v13
        _ = v14
    } else {
        v3 += v3
        var v15 int64 = 9223372036854775807
        var v16 uint8 = 2
        _ = v15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
