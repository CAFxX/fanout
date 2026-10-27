package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (1.5 * 1e10)
    var v3 float32 = (float32(-2.25) + -2.25)
    var v4 uint8 = 1
    var v5 uint32 = 525
    switch 1 {
    case 9223372036854775807:
        v6 := []int{929, 266, 2147483647}
        _ = v6
    default:
        v2 = v2
    }
    var v7 uint16 = 206
    var v8 int64 = 171
    switch 290 {
    case -254:
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int32 = -1
            v7 = (v7 >> ((1) & 63))
            _ = fi9
            _ = v10
        }
    case 9223372036854775807:
        if false {
            var v11 uint32 = 65535
            var v12 float64 = v2
            _ = v11
            _ = v12
        }
    default:
        _ = len(v1)
    }
    var v13 float64 = 0.0
    if ((v7 * 0) != (v7 >> ((4294967295) & 63))) {
        v14 := []int{-1, 374}
        var v15 uint16 = (v7 * v7)
        var v16 bool = false
        _ = v14
        _ = v15
        _ = v16
    } else {
        if true {
            v8 = (v8 / (((v8 % ((v8) | 1))) | 1))
            for fi17 := 0; fi17 < 12; fi17++ {
                v18 := GenS1{F0: 255, F1: v8}
                v4 -= 255
                if (0 != 2147483647) {
                    fmt.Println(v2, v3)
                    v3 = (1.5 / v3)
                } else {
                    v13 = (v13 - float64(v3))
                }
                _ = fi17
                _ = v18
            }
        } else {
            var v19 int = 9223372036854775807
            fmt.Println(v8)
            _ = v19
        }
        v4 -= (v4 * v4)
    }
    var v20 uint16 = v7
    var v21 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
