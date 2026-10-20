package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int; F1 string; F2 int64 }

func main() {
    var v1 int32 = 219
    var v2 float32 = (-2.25 - 1e10)
    var v3 float64 = (((0.0 - 0.0) / (float64(((1e10 * 1e10) * float32(1.5))) / 3.14159)) - ((float64(v2) / (float64(v2) + (1.5 * (0.1 * 0.0)))) - -2.25))
    var v4 int8 = 100
    v5 := make(map[uint32]uint32)
    if (!(v1 < v1)) {
        if (826 < 277) {
            switch 4294967295 {
            case 4294967295:
                fmt.Println(v2, v4)
            }
            if false {
                if false {
                    v2 -= v2
                    v6 := GenS1{F0: "a", F1: 2}
                    switch 42 {
                    case -727:
                        fmt.Println(v1, v1, v3)
                    default:
                        v4 += (v4 % ((v4) | 1))
                    }
                    _ = v6
                }
            } else {
                v7 := []int{471, 9223372036854775807, 2147483647, 127}
                fmt.Println(v1)
                _ = v7
            }
        }
        var v8 int64 = 600
        _ = v8
    }
    var v9 int8 = 42
    var v10 float32 = 0.1
    v1 += (-(v1))
    var v11 string = (("hello" + "") + "")
    var v12 uint8 = 2
    v13 := []int{-2147483648}
    if len(v13) > 0 {
        v13[0] = 0
    }
    _ = len(v13)
    v14 := GenS2{F0: 1, F1: (("" + (v11 + v11)) + (v11 + v11)), F2: 0}
    fmt.Println(v1)
    v12 = (v12 * v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
