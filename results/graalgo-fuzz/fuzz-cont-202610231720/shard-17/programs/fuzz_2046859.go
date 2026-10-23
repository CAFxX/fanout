package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint = 186
    var v2 bool = (128 < 0)
    v3 := []uint32{1, 42, 2147483647}
    var v4 float64 = 0.0
    var v5 float64 = ((1e10 * ((0.1 + (v4 * (((0.0 + float64((0.1 * (3.14159 / 3.14159)))) * ((v4 - v4) - (v4 / ((3.14159 * v4) + (0.0 + (v4 - (0.1 * v4))))))) * v4))) - 1e10)) + (v4 + ((0.0 * (float64(0.1) - float64(-2.25))) - (0.0 - -2.25))))
    var v6 int32 = 0
    v7 := GenS1{F0: ((-784 <= -9223372036854775808) && (((float64(3.14159) * 3.14159) / v4) <= -2.25))}
    var v8 int16 = int16(v6)
    var v9 uint32 = 4294967295
    fmt.Println(v5, v4)
    fmt.Println(v5)
    var v10 uint16 = 278
    fmt.Println(v6)
    if len(v3) > 0 {
        v3[0] = (v9 << (((^(v1))) & 63))
    }
    var v11 uint32 = (^(v9))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
