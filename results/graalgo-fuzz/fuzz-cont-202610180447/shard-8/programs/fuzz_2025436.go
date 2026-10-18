package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    var v1 uint = 451
    var v2 float32 = 3.14159
    v3 := make(map[int]string)
    var v4 int16 = 1000
    var v5 uint64 = 42
    var v6 uint = 2
    v7 := []int{127, 982, -774, 32767}
    for fi8 := 0; fi8 < 15; fi8++ {
        v5 = (^(v5))
        _ = fi8
    }
    v4 += (v4 / ((v4) | 1))
    v4 = (-(v4))
    v7 = append(v7, 9223372036854775807)
    var v9 int64 = 42
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    v3[1] = "foo bar"
    var v10 uint8 = 128
    var v11 int16 = 974
    v5 = 4294967295
    if false {
        var v12 bool = (-372 <= 1)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            v4 += (v4 ^ v4)
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: (v4 - 1), F1: 255, F2: v5}
        v16 := GenS2{F0: 127, F1: (^(v6))}
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
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
    _ = v9
}
