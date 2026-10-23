package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int8 = 0
    var v2 float64 = -2.25
    v3 := make(map[int]uint32)
    var v4 int16 = 338
    var v5 int16 = 32767
    v6 := make(map[int]string)
    switch "\n\t" {
    case "hello":
        fmt.Println(v4, v5, v4)
    case "":
        fmt.Println(v1, v1)
    case "foo bar":
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 15; fi8++ {
                if false {
                    var v9 float32 = 0.0
                    _ = v9
                }
                v10 := GenS1{F0: 65535, F1: 2}
                var v11 uint64 = 18446744073709551615
                var v12 uint32 = 2147483647
                _ = fi8
                _ = v10
                _ = v11
                _ = v12
            }
            v5 = int16(v1)
            _ = fi7
        }
    default:
        fmt.Println(v2, v1)
    }
    var v13 uint32 = 0
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: 65535, F1: 1}
        fmt.Println(v5)
        _ = fi14
        _ = v15
    }
    fmt.Println(v2, v2, v1)
    var v16 int32 = -2147483648
    var v17 int8 = v1
    v4 += (v4 & v5)
    v18 := GenS2{F0: (v16 / ((2147483647) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
