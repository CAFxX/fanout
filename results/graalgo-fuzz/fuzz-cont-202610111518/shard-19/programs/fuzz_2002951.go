package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 uint, p1 int8, p2 string) uint32 {
    v5 := GenS1{F0: 9223372036854775807, F1: -967}
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int = (fi6 % (((fi6 % ((0) | 1))) | 1))
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 int = 128
    var v2 int32 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    v2 += 287
    var v8 uint32 = 2
    if (true && false) {
        var v9 float32 = ((float32((1.5 + (-2.25 / 1e10))) - 1.5) - (1.5 / (3.14159 - float32(1e10))))
        v10 := []int{885, v1, 127, v1}
        _ = v10
        _ = v9
    }
    var v11 uint = 18446744073709551615
    var v12 uint = 42
    v13 := []int{-308}
    var v14 float32 = 1.5
    if (false || (255 > 745)) {
        v15 := v3[(-(v1))]
        v8 = (^(v15))
        if true {
            var v16 int8 = -128
            _ = v16
        } else {
            if (!true) {
                if (1 != 156) {
                    var v17 float64 = ((0.0 / 3.14159) + float64(-2.25))
                    if true {
                        delete(v4, v15)
                        var v18 int64 = 42
                        _ = v18
                    }
                    v1 += (v1 ^ 548)
                    _ = v17
                } else {
                    var v19 uint8 = 127
                    for fi20 := 0; fi20 < 7; fi20++ {
                        fmt.Println(v19, v12, v15)
                        _ = fi20
                    }
                    _ = v19
                }
            }
        }
        _ = v15
    } else {
        fmt.Println(v11)
        var v21 uint8 = 166
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
