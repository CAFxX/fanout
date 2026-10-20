package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 0
    var v3 uint32 = 2
    var v4 float64 = ((((-2.25 + 3.14159) / -2.25) + 1.5) - float64(float32((0.0 * 0.1))))
    v3 -= 1
    fmt.Println(v3)
    var v5 int = 4294967295
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(v4, v5)
        if (3.14159 >= -2.25) {
            v2 = 1
            v3 += (v3 << ((255) & 63))
            v7 := GenS1{F0: -1}
            _ = v7
        } else {
            var v8 float64 = (v4 + (v4 / 1.0))
            _ = v8
        }
        if (-1 < -9223372036854775808) {
            v9 := GenS2{F0: 100, F1: -1}
            var v10 uint = 42
            fmt.Println(v3)
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    var v11 int8 = -1
    v11 -= (v11 ^ (-(v11)))
    v5 += (v5 + v5)
    v4 = ((v4 + v4) / ((v4 + v4) + 0.1))
    v12 := GenS3{F0: 65535, F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
