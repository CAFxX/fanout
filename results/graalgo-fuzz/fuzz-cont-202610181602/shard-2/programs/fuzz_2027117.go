package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    var v1 float64 = ((((1e10 * float64(3.14159)) + ((float64(1e10) + (-2.25 + 0.1)) * (float64(1e10) / (0.1 / (float64(((0.1 + 1.5) * 1e10)) - ((float64(float32(1.5)) * 0.0) * 0.0)))))) - ((((3.14159 * (1.5 / ((1.5 * 0.1) + ((0.1 / 3.14159) - 1.5)))) * (0.0 * ((1e10 * -2.25) / 3.14159))) - 1.5) - 1.5)) - 3.14159)
    var v2 int8 = 0
    v3 := make(map[string]bool)
    var v4 uint8 = 128
    var v5 uint64 = 42
    var v6 float32 = (1e10 + -2.25)
    var v7 int64 = 1
    v8 := GenS1{F0: (v2 >> ((439) & 63)), F1: (^(v5))}
    var v9 uint64 = (v5 >> ((65535) & 63))
    if (!(!false)) {
        v10 := []int{-708, 2147483647, -92, 127}
        _ = v10
    } else {
        if true {
            var v11 int8 = v2
            v12 := []int{32768, 9223372036854775807, 9223372036854775807, -965}
            _ = v11
            _ = v12
        } else {
            v13 := []int{int(v2), -1, 42}
            if true {
                v14 := GenS1{F0: (v2 + (v2 << ((4294967295) & 63))), F1: (^(v9))}
                v9 -= 0
                _ = v14
            } else {
                v6 += (v6 / (v6 + v6))
            }
            _ = v13
        }
    }
    switch 1 {
    case 272:
        for fi15 := 0; fi15 < 1; fi15++ {
            v4 += v4
            v2 -= v2
            _ = fi15
        }
    case 2147483647:
        if (2 == 0) {
            v1 += v1
        }
    case -279:
        v16 := GenS1{F0: (v2 & v2), F1: (^(v5))}
        _ = v16
    }
    v5 -= (v9 - (v5 & v9))
    var v17 int64 = v7
    var v18 bool = (1 == 120)
    v1 = (v1 + ((v1 * v1) - float64(v6)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
