package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("a" + "\n\t")
    var v3 uint8 = 127
    var v4 uint32 = 2147483647
    v3 = (v3 >> ((0) & 63))
    var v5 uint16 = 42
    var v6 float64 = (((3.14159 + (float64(-2.25) + (1e10 / 1e10))) - 1.5) * (3.14159 / (0.1 + -2.25)))
    var v7 int16 = -160
    delete(v1, 531)
    switch 32767 {
    case 993:
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 bool = (619 >= 255)
            _ = len(v1)
            fmt.Println(v3, fi8)
            _ = fi8
            _ = v9
        }
    case -157:
        _ = len(v1)
    case -2147483648:
        var v10 string = v2
        _ = v10
    default:
        var v11 float64 = 3.14159
        _ = v11
    }
    fmt.Println(v7, v2, v2)
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: 100}
        v7 -= 42
        _ = fi12
        _ = v13
    }
    v1[-232] = (v4 % ((876) | 1))
    v6 = (-2.25 * ((v6 * 0.1) + 1.5))
    v14 := []int{42, -9223372036854775808, 2147483647, 752}
    v6 = v6
    v3 -= (v3 / ((1) | 1))
    fmt.Println(v7)
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS2{F0: -128, F1: (v4 + uint32(v5))}
        var v17 bool = ((v5 * (v5 - (v5 % ((v5) | 1)))) != v5)
        var v18 uint64 = 42
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint16 = (v5 * v5)
    v3 += uint8(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
