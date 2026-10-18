package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0() uint {
    var v7 int64 = 798
    v7 = -1
    _ = v7
    return 65535
}
func genFn1(p0 int8) int32 {
    switch 128 {
    case -2147483648:
        p0 += p0
    case 944:
        fmt.Println(p0)
    case 818:
        switch -925 {
        case -111:
            v8 := []int{0, -91, 32768}
            _ = v8
        default:
            p0 += (p0 & (p0 & p0))
        }
    default:
        var v9 int = -2147483648
        _ = v9
    }
    var v10 uint16 = 32767
    fmt.Println(p0, p0)
    _ = p0
    _ = v10
    return 836
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]int)
    var v3 float64 = (float64((0.1 + 1.5)) + float64(-2.25))
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 int = 2147483647
    var v11 uint8 = 5
    fmt.Println(v1)
    v2[0] = (v6 / ((v6) | 1))
    var v12 string = "hello"
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int64 = -273
        fmt.Println(v3, v6, v6)
        fmt.Println(v12)
        _ = fi13
        _ = v14
    }
    var v15 float64 = 0.1
    var v16 uint16 = 476
    v17 := GenS1{F0: (v11 << ((2) & 63)), F1: 42, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
