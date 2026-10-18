package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 float64 = (1.5 - (((-2.25 - ((0.0 / 1.5) / 3.14159)) / 3.14159) + 1e10))
    var v4 string = "a"
    v4 = "hello"
    var v5 int16 = -32768
    v3 += v3
    v6 := []int{-2147483648, 9223372036854775807, 32768}
    var v7 uint = 1
    v3 = ((v3 / v3) * (v3 - 3.14159))
    var v8 uint64 = 641
    v3 += (v3 / v3)
    var v9 float32 = float32(v3)
    fmt.Println(v8)
    var v10 uint = (v7 ^ 1)
    var v11 uint = (v7 >> ((v7) & 63))
    v9 += (0.1 + v9)
    v11 = 725
    v12 := GenS1{F0: -2147483648}
    if (v8 > (v8 >> (((v7 - (v10 & 255))) & 63))) {
        if len(v6) > 0 {
            v6[0] = -225
        }
        v13 := GenS1{F0: -359}
        v10 = v10
        _ = v13
    } else {
        fmt.Println(v10, v4)
        v14 := GenS1{F0: 288}
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint16 = 0
            for fi17 := 0; fi17 < 9; fi17++ {
                v9 = 0.0
                v18 := GenS2{F0: -9223372036854775808, F1: 325}
                v7 = v7
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    if (true && false) {
        var v19 int8 = 0
        var v20 bool = false
        var v21 uint64 = (v8 + v8)
        _ = v19
        _ = v20
        _ = v21
    } else {
        v7 += v7
        delete(v2, 500)
        v22 := GenS1{F0: 322}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
