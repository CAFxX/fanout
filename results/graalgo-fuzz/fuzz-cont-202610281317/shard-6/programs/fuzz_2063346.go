package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (-901 != -2147483648)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483647
    var v5 uint16 = 65535
    var v6 int64 = 0
    var v7 int = 701
    var v8 float32 = float32(0.0)
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint64 = 42
            v2 = (false || true)
            _ = fi9
            _ = v10
        }
        fmt.Println(v5)
        v8 = float32((((-2.25 / 1e10) / (0.1 / 0.1)) + (1.5 - (3.14159 / 3.14159))))
    } else {
        var v11 uint = 4294967295
        v4 -= (v4 * 478)
        _ = v11
    }
    fmt.Println(v4, v6)
    v12 := GenS1{F0: v5, F1: 32767}
    v13 := v1[(^(v4))]
    v14 := GenS2{F0: 837, F1: (v5 | 982), F2: 1}
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 bool = (((2 >= 1) && false) && (true && v2))
        _ = fi15
        _ = v16
    }
    var v17 int32 = 2147483647
    v8 = (float32(0.1) * 0.1)
    switch (v7 & (-(v7))) {
    case 304:
        for fi18 := 0; fi18 < 17; fi18++ {
            if true {
                v19 := GenS1{F0: 42, F1: -915}
                _ = v19
            } else {
                v1[(v4 & 639)] = ("" + "hello")
            }
            _ = fi18
        }
    case 169:
        _ = len(v1)
    default:
        var v20 uint = 934
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
