package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    var v1 int32 = -564
    var v2 uint16 = 775
    var v3 int16 = -398
    v4 := make(map[int]int)
    v5 := make(map[uint32]string)
    var v6 float64 = 3.14159
    var v7 uint = 2
    var v8 uint = v7
    switch "" {
    case "":
        v4[128] = 4294967295
    case "hello":
        v9 := v5[2147483647]
        _ = v9
    }
    fmt.Println(v6)
    fmt.Println(v7, v8)
    v3 = (^(v3))
    v2 = (v2 ^ 1)
    var v10 int16 = 211
    var v11 uint32 = 293
    switch 0 {
    case 128:
        fmt.Println(v11)
    case 32767:
        for fi12 := 0; fi12 < 2; fi12++ {
            if false {
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS1{F0: (v1 | v1), F1: (v3 & v3)}
                    _ = fi13
                    _ = v14
                }
                v7 = 375
            } else {
                v8 -= (v8 / ((v7) | 1))
                var v15 int64 = -9223372036854775808
                _ = v15
            }
            var v16 int32 = v1
            v5[uint32(v2)] = "a"
            _ = fi12
            _ = v16
        }
    case 4294967295:
        v17 := []int{-2147483648, -2147483648, 0, 0}
        _ = v17
    default:
        if (!((v1 / ((0) | 1)) < v1)) {
            for fi18 := 0; fi18 < 10; fi18++ {
                v1 -= 206
                _ = fi18
            }
        }
    }
    v6 = float64(1.5)
    v3 = v10
    var v19 int8 = -81
    v20 := v5[v11]
    for fi21 := 0; fi21 < 20; fi21++ {
        if ((189 != 1) || ((false && false) && ((v2 >= (v2 ^ v2)) || false))) {
            v1 = 1
            fmt.Println(v11, v20, v7)
            var v22 uint32 = (^(v11))
            _ = v22
        }
        v11 = (v11 - 42)
        _ = fi21
    }
    v3 -= (v10 & v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
