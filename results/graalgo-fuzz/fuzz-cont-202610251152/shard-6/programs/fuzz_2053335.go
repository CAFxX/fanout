package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 uint16 = 734
    fmt.Println(v4)
    var v5 float64 = 0.1
    fmt.Println(v5)
    if (0.0 == -2.25) {
        var v6 uint16 = v4
        _ = v6
    }
    v5 = (((3.14159 + float64(float32(v5))) - ((v5 * v5) * (1.5 * float64((float32(1.5) + 0.1))))) - float64(3.14159))
    v7 := GenS1{F0: 42, F1: -2147483648}
    var v8 uint16 = (v4 / ((0) | 1))
    v9 := GenS1{F0: 32767, F1: 1}
    var v10 uint = 42
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int8 = 0
        _ = fi11
        _ = v12
    }
    switch 608 {
    case 42:
        var v13 uint = (v10 | (v10 / (((v10 * 1)) | 1)))
        _ = v13
    case -2147483648:
        if true {
            v8 -= v8
        }
    default:
        v14 := []int{9223372036854775807, 128, 127, 2147483647}
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
