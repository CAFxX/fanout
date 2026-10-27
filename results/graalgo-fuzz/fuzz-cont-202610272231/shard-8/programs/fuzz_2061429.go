package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1e10
    var v3 float64 = 0.0
    v4 := make(map[uint32]uint32)
    var v5 bool = (1 != 0)
    var v6 int64 = -218
    v3 = v3
    var v7 uint = 2
    delete(v4, 702)
    var v8 float64 = 0.0
    var v9 uint8 = 127
    v10 := v4[uint32(v7)]
    switch "\n\t" {
    case "\n\t":
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: v6, F1: 1}
            fmt.Println(v9)
            _ = fi11
            _ = v12
        }
    case "":
        for fi13 := 0; fi13 < 17; fi13++ {
            v3 -= v8
            _ = fi13
        }
    default:
        switch 4294967295 {
        case 32768:
            if v5 {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v5 = (42 == 0)
                    _ = fi14
                }
                v9 += (v9 / (((v9 ^ v9)) | 1))
            }
        case 2147483647:
            var v15 uint8 = (v9 | (v9 | (v9 * v9)))
            _ = v15
        }
    }
    v16 := v4[66]
    fmt.Println(v8, v7)
    v17 := GenS1{F0: -9223372036854775808, F1: 127}
    v16 = v10
    fmt.Println(v3, v7, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
