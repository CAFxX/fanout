package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    v1 := []uint32{815}
    var v2 uint32 = 42
    var v3 float64 = ((float64((1e10 + 1.5)) / ((float64(0.1) * (3.14159 + (((3.14159 + (((float64((0.0 + (0.1 + float32(-2.25)))) / 3.14159) + 0.0) / 1.0)) / (float64(-2.25) + (3.14159 / 3.14159))) + 3.14159))) + ((0.0 * -2.25) - 0.1))) - ((float64(0.0) - ((0.0 * float64(1.5)) / 1e10)) * (1.5 + ((3.14159 - 1e10) - ((0.1 / 0.1) - (0.1 / 0.1))))))
    v4 := make(map[uint32]string)
    fmt.Println(v2, v3)
    v5 := GenS1{F0: 652, F1: -760}
    fmt.Println(v3)
    var v6 int8 = -1
    for fi7 := 0; fi7 < 3; fi7++ {
        v4[(v2 + 65535)] = "a"
        _ = fi7
    }
    switch -833 {
    case -734:
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint32 = v2
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint16 = 65535
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    case 32768:
        fmt.Println(v3, v3, v2)
    default:
        var v12 bool = (uint(v2) <= 596)
        _ = v12
    }
    v13 := GenS2{F0: v2, F1: 4294967295}
    v3 = (((0.0 / v3) - v3) / (0.1 / float64((1.5 - float32(((v3 * v3) + (v3 - v3)))))))
    v14 := GenS1{F0: 255, F1: 757}
    var v15 uint = 509
    fmt.Println(v2, v2, v6)
    fmt.Println(v15, v15, v2)
    var v16 uint32 = (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
