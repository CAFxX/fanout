package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int32 = 0
    var v2 float32 = 0.0
    var v3 uint16 = 42
    v4 := make(map[int]uint32)
    var v5 int64 = -9223372036854775808
    if ((false && ("" <= "")) || (!(true && ((^(v3)) == (^(v3)))))) {
        v5 = v5
        _ = len(v4)
    }
    v6 := GenS1{F0: ((32768 > 438) && (255 <= 0)), F1: 2}
    v5 -= 7
    fmt.Println(v3)
    v7 := GenS2{F0: (v3 + v3)}
    var v8 float64 = (0.0 * 1.5)
    v9 := GenS3{F0: 974}
    var v10 float64 = v8
    v8 += float64((v2 - ((v2 * v2) + ((v2 / (((v2 - 1.5) - 1e10) + (((((float32(float64(v2)) - v2) + v2) / v2) + 3.14159) / (v2 / v2)))) * (v2 * 3.14159)))))
    var v11 uint32 = 2147483647
    v1 += (v1 | (v1 * 699))
    delete(v4, -455)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
