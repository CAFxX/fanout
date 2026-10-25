package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
type GenS3 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int16 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]uint32)
    var v3 int16 = 574
    v4 := []string{"foo bar", "hello"}
    v5 := []int{-1, -2147483648, 128}
    var v6 uint16 = 0
    if true {
        var v7 uint8 = 1
        var v8 uint64 = 18446744073709551615
        _ = v7
        _ = v8
    }
    v3 -= -1
    v4 = append(v4, "hello")
    var v9 uint64 = 0
    if (true || true) {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int64 = 855
            v12 := GenS1{F0: (^(fi10)), F1: v11, F2: false}
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(v9, v9)
        fmt.Println(v3, v3, v6)
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        switch "" {
        case "\n\t":
            var v14 int64 = 9223372036854775807
            _ = v14
        case "":
            v15 := GenS2{F0: 1, F1: (v3 % ((v3) | 1)), F2: 1}
            _ = v15
        }
        _ = fi13
    }
    fmt.Println(v1, v1)
    var v16 uint16 = v6
    v17 := GenS3{F0: 100, F1: v6, F2: 59}
    if (2147483647 == 0) {
        v18 := GenS4{F0: 15, F1: (-(v3))}
        _ = v18
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    var v19 float64 = float64(float32(0.1))
    v6 -= uint16(v9)
    fmt.Println(v6, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
