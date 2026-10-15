package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0() uint32 {
    switch "\n\t" {
    case "\n\t":
        var v6 uint = 2
        _ = v6
    case "":
        if false {
            if (842 <= -1) {
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := []int{32767, fi7}
                    _ = fi7
                    _ = v8
                }
            } else {
                v9 := GenS1{F0: 1, F1: 65535}
                var v10 int8 = -53
                _ = v10
                _ = v9
            }
            var v11 int = 128
            v12 := []int{(^(v11)), v11}
            _ = v11
            _ = v12
        }
    default:
        v13 := GenS1{F0: -1, F1: 42}
        _ = v13
    }
    v14 := []int{-9223372036854775808, 83, 862}
    var v15 uint = 61
    _ = v14
    _ = v15
    return 417
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]int)
    var v3 int64 = 1
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    var v16 uint16 = v1
    var v17 int = -9223372036854775808
    var v18 int16 = 0
    for fi19 := 0; fi19 < 12; fi19++ {
        _ = len(v4)
        var v20 string = ""
        _ = fi19
        _ = v20
    }
    fmt.Println(v1)
    v21 := GenS1{F0: -39, F1: v1}
    var v22 uint = 771
    v23 := []int{v17}
    var v24 uint64 = 42
    var v25 int32 = 1
    v4[1] = (v17 != (v17 / (((v17 % (((v17 / ((v17) | 1))) | 1))) | 1)))
    fmt.Println(v16, v3, v24)
    _ = len(v23)
    var v26 uint8 = 1
    v17 += v17
    if len(v23) > 0 {
        v23[0] = (v17 & (-(v17)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
