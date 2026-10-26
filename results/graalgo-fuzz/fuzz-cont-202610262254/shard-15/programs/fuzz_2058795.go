package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0() int {
    if (134 > 42) {
        for fi4 := 0; fi4 < 19; fi4++ {
            var v5 uint = 65535
            var v6 uint8 = 128
            _ = fi4
            _ = v5
            _ = v6
        }
        var v7 int64 = -9223372036854775808
        v8 := GenS1{F0: 42}
        _ = v7
        _ = v8
    } else {
        switch 2147483647 {
        case 9223372036854775807:
            var v9 int = 9223372036854775807
            _ = v9
        default:
            v10 := []int{128, -9223372036854775808, 0, 127}
            _ = v10
        }
        var v11 uint32 = 374
        _ = v11
    }
    var v12 int32 = 0
    v12 = (v12 % (((v12 * v12)) | 1))
    v12 += v12
    v13 := GenS1{F0: 255}
    _ = v12
    _ = v13
    return 32767
}

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 42
    var v3 int64 = 0
    var v14 float32 = 0.0
    if true {
        v2 += (v2 ^ v2)
        var v15 string = "\n\t"
        _ = v15
    }
    var v16 uint = 2
    v17 := []int{-9223372036854775808, -2147483648, 128, 9223372036854775807}
    if false {
        var v18 uint32 = 2147483647
        v1 += v1
        v19 := GenS2{F0: 41, F1: false}
        _ = v18
        _ = v19
    }
    if (3.14159 != -2.25) {
        fmt.Println(v14, v1, v3)
        v17 = append(v17, 338)
    } else {
        switch 42 {
        case 42:
            v16 += 873
        }
        v1 = (v1 % ((uint8(v2)) | 1))
        var v20 int64 = int64(v2)
        _ = v20
    }
    if (!(!(v1 < (v1 + (^(v1)))))) {
        var v21 float32 = v14
        _ = v21
    }
    v17 = append(v17, 2147483647)
    var v22 int32 = -907
    v16 += v16
    var v23 float32 = ((v14 + (-2.25 - float32(float64(v14)))) / -2.25)
    v22 = (v22 + 0)
    if len(v17) > 0 {
        v17[0] = 2147483647
    }
    v24 := GenS2{F0: 0, F1: false}
    v22 = 1
    var v25 bool = ((v14 - (v23 + v14)) < v23)
    var v26 int64 = (v3 ^ 42)
    var v27 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
