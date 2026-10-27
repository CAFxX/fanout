package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 700
    var v3 uint32 = 1
    var v4 uint64 = 44
    v5 := GenS1{F0: -1, F1: 2147483647}
    fmt.Println(v3, v4)
    for fi6 := 0; fi6 < 12; fi6++ {
        v3 -= v3
        _ = fi6
    }
    _ = len(v1)
    var v7 uint32 = 2147483648
    var v8 int8 = -128
    if (-1 > 431) {
        var v9 uint64 = v4
        if true {
            fmt.Println(v4, v9, v7)
        }
        _ = v9
    } else {
        if true {
            v10 := v1[292]
            switch 0 {
            case 0:
                for fi11 := 0; fi11 < 2; fi11++ {
                    switch fi11 {
                    case 9223372036854775807:
                        var v12 float64 = (-2.25 - 1e10)
                        _ = v12
                    case 42:
                        var v13 string = "a"
                        _ = v13
                    default:
                        var v14 uint8 = 42
                        _ = v14
                    }
                    var v15 int = fi11
                    v16 := GenS2{F0: uint64(v2)}
                    _ = fi11
                    _ = v15
                    _ = v16
                }
            }
            v17 := GenS1{F0: (v10 * v10), F1: v10}
            _ = v10
            _ = v17
        } else {
            v3 += v7
            var v18 uint32 = (v7 % ((v7) | 1))
            _ = v18
        }
        v19 := []int{32767, 32}
        _ = v19
    }
    v20 := GenS3{F0: 9223372036854775807, F1: (v4 ^ 639), F2: (^(v4))}
    v21 := GenS4{F0: -230, F1: 32768}
    var v22 uint16 = 1
    if true {
        v23 := GenS5{F0: 42}
        _ = v23
    }
    delete(v1, 127)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
