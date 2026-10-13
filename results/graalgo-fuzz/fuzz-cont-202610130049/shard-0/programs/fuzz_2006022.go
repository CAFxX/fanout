package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int }
type GenS2 struct { F0 string; F1 uint8; F2 int8 }
func genFn0(p0 uint64, p1 uint16) int64 {
    p1 = (p1 ^ (p1 & p1))
    p0 += (p0 >> ((2) & 63))
    v4 := []int{42}
    p0 = 844
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 float64 = 0.1
    var v2 int = 0
    var v3 uint32 = 123
    if true {
        switch (v2 ^ v2) {
        case 32767:
            var v5 int16 = -1
            _ = v5
        case -933:
            fmt.Println(v2, v3, v3)
        case 2147483647:
            var v6 uint32 = v3
            _ = v6
        }
        v3 = v3
    }
    var v7 float32 = float32(0.0)
    v3 += (v3 >> ((500) & 63))
    if true {
        var v8 uint32 = (v3 * (^(v3)))
        v1 = v1
        fmt.Println(v3, v2)
        _ = v8
    } else {
        v2 += (^(v2))
        v9 := GenS1{F0: v3, F1: 0, F2: v2}
        _ = v9
    }
    v3 += 2147483647
    v7 -= 3.14159
    var v10 uint16 = 2
    var v11 uint64 = 42
    v12 := GenS2{F0: ("\n\t" + ""), F1: 203, F2: 1}
    v3 -= v3
    if (127 <= 41) {
        v7 -= float32(1e10)
        var v13 uint64 = v11
        v7 = ((float32(float64(0.0)) * v7) * 1e10)
        _ = v13
    }
    var v14 uint16 = (v10 >> ((255) & 63))
    fmt.Println(v1, v7, v14)
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := []int{fi15, v2}
        _ = fi15
        _ = v16
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v7
}
