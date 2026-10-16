package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 498
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 float32 = 1e10
    var v5 uint64 = 18446744073709551615
    v6 := GenS1{F0: 0, F1: (v1 * v1), F2: false}
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 9; fi8++ {
        v1 = (^(v1))
        switch "" {
        case "hello":
            v1 = (^(v1))
        }
        v1 += (v1 & (v1 % (((v1 ^ v1)) | 1)))
        _ = fi8
    }
    _ = len(v3)
    var v9 uint32 = 2147483648
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint64 = 881
        v4 -= v4
        _ = fi10
        _ = v11
    }
    if false {
        fmt.Println(v5)
    }
    v12 := v3[9223372036854775807]
    fmt.Println(v4)
    for fi13 := 0; fi13 < 4; fi13++ {
        switch fi13 {
        case 127:
            fmt.Println(v5, v4, v4)
        case 358:
            v14 := []int{(fi13 >> (((v1 >> ((65535) & 63))) & 63)), fi13, (-(fi13))}
            _ = v14
        case 2147483647:
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(v12)
                _ = fi15
            }
        }
        var v16 int64 = -1
        _ = fi13
        _ = v16
    }
    v17 := []int{1, 128, 166, 1}
    v2[(v9 / (((v9 * v12)) | 1))] = 4294967295
    v18 := GenS2{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
