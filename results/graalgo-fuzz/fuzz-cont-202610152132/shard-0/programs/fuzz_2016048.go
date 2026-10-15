package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 string }
func genFn0() int64 {
    var v4 int8 = -109
    var v5 uint64 = 959
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 1
    var v3 float32 = (1e10 / (1.5 * 1.5))
    v6 := GenS1{F0: (v2 * 1), F1: -1, F2: -13}
    var v7 uint = 1
    fmt.Println(v1)
    v8 := GenS1{F0: v2, F1: -1, F2: 1}
    fmt.Println(v1, v3)
    var v9 uint8 = (^(v2))
    var v10 int = -126
    if (((3.14159 / 1e10) - -2.25) >= ((1.5 / 0.1) + 0.1)) {
        v3 -= float32(1.5)
    } else {
        v7 = (^(v7))
        var v11 float64 = 0.0
        _ = v11
    }
    v3 = -2.25
    v12 := []int{}
    v3 -= 1.5
    var v13 float32 = v3
    v7 = v7
    for fi14 := 0; fi14 < 2; fi14++ {
        v2 = (v9 * 0)
        v15 := GenS2{F0: "a"}
        if (((-2.25 - float32((-2.25 + (0.0 / 1e10)))) * v13) != ((1e10 - v3) / (((3.14159 + ((v3 * (v13 - float32(float64(float32((1e10 / (0.1 / 0.1))))))) - (v13 / v3))) * float32(0.0)) + 0.0))) {
            fmt.Println(v13)
            if true {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
            v9 += (v2 | uint8(v7))
        }
        _ = fi14
        _ = v15
    }
    var v16 uint8 = (v9 % ((v9) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
