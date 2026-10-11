package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 uint64; F1 string; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 75
    var v3 uint = 42
    v4 := []int{-1, 128, 32767}
    v5 := make(map[string]uint32)
    var v6 uint16 = 32767
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 15; fi8++ {
        v3 = v3
        _ = fi8
    }
    var v9 uint8 = 253
    if false {
        v10 := GenS1{F0: (^(v3)), F1: -758, F2: 42}
        v11 := GenS2{F0: 42, F1: (("" + "\n\t") + ""), F2: 2}
        v12 := v1[1]
        _ = v10
        _ = v11
        _ = v12
    } else {
        v7 -= float64(3.14159)
        var v13 bool = true
        v7 -= (v7 * v7)
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v3, v6, v2)
        var v15 int32 = -2147483648
        v6 = (v6 >> (((v3 - v3)) & 63))
        _ = fi14
        _ = v15
    }
    var v16 int32 = -904
    v7 = float64(0.0)
    v6 -= v6
    v1[4294967295] = ""
    var v17 int64 = 9223372036854775807
    var v18 int = 127
    var v19 int64 = (v17 + v17)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (v7 < v7) {
        fmt.Println(v16, v16)
        if (4294967295 > 2147483648) {
            v2 = (^(v2))
        }
        if (-1 == (v19 * (-(v19)))) {
            v17 -= 9223372036854775807
            fmt.Println(v18, v7, v18)
        } else {
            fmt.Println(v2)
        }
    }
    var v20 uint = (v3 | v3)
    v21 := GenS2{F0: 452, F1: "", F2: (v20 * v20)}
    var v22 float64 = float64((((3.14159 / (0.0 + -2.25)) / (3.14159 / 0.1)) / 1e10))
    v23 := v1[2]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
