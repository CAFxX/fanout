package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]string)
    var v3 string = "a"
    v3 = "\n\t"
    v4 := []int{32767, 2147483647, -9223372036854775808, 4294967295}
    fmt.Println(v3, v3)
    v5 := GenS1{F0: ((0.1 == 0.0) || true)}
    fmt.Println(v3)
    v3 = v3
    v1 = ((((v1 * (float32(float64(0.1)) - (0.1 - 1.5))) / v1) + v1) * v1)
    v6 := GenS2{F0: 0, F1: 1}
    switch 986 {
    case 2147483647:
        var v7 float64 = (float64(((3.14159 * v1) + float32((0.1 * 0.0)))) / (3.14159 - 0.1))
        _ = v7
    case 32768:
        fmt.Println(v1)
    case -1:
        fmt.Println(v1, v1)
    }
    v1 = v1
    var v8 float32 = v1
    for fi9 := 0; fi9 < 8; fi9++ {
        v1 -= v8
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v1)
        _ = fi9
    }
    var v10 int8 = 42
    if len(v4) > 0 {
        v4[0] = 1
    }
    v10 += 100
    v10 += (v10 << ((2) & 63))
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
