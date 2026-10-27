package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }
type GenS2 struct { F0 bool }
func genFn0() uint8 {
    var v7 uint8 = 127
    v7 = (^(v7))
    v8 := []int{42, 469, 9223372036854775807, 0}
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 float32 = float32(1.5)
    var v2 float32 = 1e10
    var v3 int64 = 9223372036854775807
    var v4 float64 = (3.14159 + (-2.25 - (0.1 - 1.5)))
    v5 := []int{}
    var v6 uint32 = 699
    v9 := GenS1{F0: -477, F1: 42, F2: 126}
    var v10 int64 = (^(v3))
    v5 = append(v5, -2147483648)
    v11 := GenS1{F0: -292, F1: -266, F2: -73}
    var v12 uint8 = 42
    var v13 int32 = -897
    if true {
        var v14 uint = 4294967295
        v13 += v13
        _ = v14
    }
    if len(v5) > 0 {
        v5[0] = 299
    }
    var v15 float64 = v4
    if (false && ((float64(0.1) - ((v15 - 1e10) + v4)) <= float64((v1 / v1)))) {
        var v16 uint = 0
        if true {
            var v17 uint8 = (^(v12))
            v18 := GenS1{F0: (v13 % (((v13 | (v13 & v13))) | 1)), F1: v13, F2: 42}
            fmt.Println(v10, v16, v12)
            _ = v17
            _ = v18
        }
        fmt.Println(v4, v16, v6)
        _ = v16
    } else {
        v5 = append(v5, 528)
        var v19 uint = 444
        _ = v19
    }
    _ = len(v5)
    v20 := GenS1{F0: v13, F1: (v13 << ((1) & 63)), F2: 100}
    v21 := GenS2{F0: false}
    if true {
        if false {
            if true {
                fmt.Println(v2, v10, v2)
                v2 = 1.5
            }
            for fi22 := 0; fi22 < 8; fi22++ {
                v12 += v12
                fmt.Println(v10)
                _ = fi22
            }
        }
        v5 = append(v5, 2147483647)
    } else {
        v13 = v13
    }
    v12 += (v12 + v12)
    v23 := GenS1{F0: (v13 % ((679) | 1)), F1: v13, F2: 42}
    var v24 int64 = (v10 / (((v10 - v3)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
