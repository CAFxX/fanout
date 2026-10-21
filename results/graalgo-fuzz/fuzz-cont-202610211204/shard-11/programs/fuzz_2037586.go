package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int8) int16 {
    var v6 uint64 = 18446744073709551615
    if (-542 >= 42) {
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 uint32 = 0
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    }
    v10 := GenS1{F0: 1, F1: 128}
    _ = p0
    _ = v10
    _ = v6
    return 0
}

func main() {
    var v1 uint8 = 244
    var v2 uint16 = 255
    var v3 int32 = 0
    var v4 float64 = float64((0.1 + 1e10))
    v5 := make(map[string]string)
    var v11 int32 = 2147483647
    switch -273 {
    case 105:
        fmt.Println(v3, v1)
    default:
        if (!(1 < 338)) {
            v12 := GenS2{F0: 0, F1: 2, F2: v11}
            _ = v12
        } else {
            _ = len(v5)
        }
    }
    v11 += v3
    v3 = (v11 / ((-3) | 1))
    fmt.Println(v4)
    v2 += 1
    var v13 uint = 357
    v2 = v2
    if (128 > uint8(v13)) {
        if true {
            v14 := GenS3{F0: 99, F1: -107}
            v15 := []int{32768}
            _ = v14
            _ = v15
        } else {
            v16 := v5["hello"]
            v17 := []int{int(v11), 128, -182, -2147483648}
            fmt.Println(v13)
            _ = v16
            _ = v17
        }
    }
    v18 := []int{1, 240}
    v1 -= uint8(v13)
    if (!(491 < 4294967295)) {
        v11 += (-(v3))
    }
    v11 = (v11 >> ((v13) & 63))
    var v19 int64 = 42
    switch "hello" {
    case "":
        fmt.Println(v2, v2, v4)
    }
    var v20 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
