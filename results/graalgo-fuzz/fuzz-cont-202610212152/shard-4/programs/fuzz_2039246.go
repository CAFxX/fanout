package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint, p1 int64) int8 {
    var v5 uint8 = 0
    v5 = v5
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 int = 2147483647
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    var v6 int = v2
    v6 -= v2
    if true {
        fmt.Println(v6, v6)
        for fi7 := 0; fi7 < 12; fi7++ {
            v6 += v6
            for fi8 := 0; fi8 < 10; fi8++ {
                v2 = (-(v2))
                _ = fi8
            }
            _ = fi7
        }
        if true {
            var v9 uint16 = 42
            var v10 uint32 = 471
            fmt.Println(v2)
            _ = v10
            _ = v9
        }
    } else {
        _ = len(v1)
        v6 -= 232
    }
    var v11 bool = (float64(0.0) < ((1.5 / (1e10 / float64(float32((0.1 - (3.14159 - (1.5 - float64(1e10)))))))) - 0.1))
    v6 += -2147483648
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v2)
        v11 = (2 >= 2)
        _ = fi12
    }
    v2 += 9223372036854775807
    v2 -= v6
    fmt.Println(v6, v11)
    fmt.Println(v2, v2)
    v6 -= (v2 ^ (^(v2)))
    var v13 uint = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4[635] = 42
    v11 = v11
    v14 := GenS1{F0: 65535, F1: -630}
    v13 += 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
