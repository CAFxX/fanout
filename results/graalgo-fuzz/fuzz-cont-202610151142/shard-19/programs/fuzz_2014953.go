package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
func genFn0(p0 int) int32 {
    p0 = p0
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint32 = 429
        _ = fi6
        _ = v7
    }
    _ = p0
    return 307
}

func main() {
    var v1 uint8 = 103
    var v2 float64 = 1e10
    var v3 bool = false
    var v4 bool = v3
    var v5 bool = (1 >= -2147483648)
    var v8 int32 = int32(v1)
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(v4)
            var v11 uint = 1
            v8 -= v8
            var v12 string = "\n\t"
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := []int{}
        _ = fi9
        _ = v13
    }
    if (255 >= 18446744073709551615) {
        v14 := []int{-2147483648, 4294967295, 9223372036854775807, 2147483647}
        fmt.Println(v4)
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := GenS1{F0: (fi15 | (fi15 >> ((0) & 63))), F1: 42, F2: 1}
            var v17 uint32 = 0
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    v4 = (false || (295 != 4294967295))
    var v18 int16 = -1
    var v19 uint8 = 100
    v8 -= 1
    for fi20 := 0; fi20 < 8; fi20++ {
        v19 += (v1 / ((128) | 1))
        _ = fi20
    }
    v8 += (^(v8))
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 float64 = 1.5
        v19 += v19
        _ = fi21
        _ = v22
    }
    if false {
        v23 := GenS2{F0: 65535, F1: uint16(v1)}
        _ = v23
    }
    v19 -= v1
    fmt.Println(v8)
    v1 += (v1 & (v1 / (((v1 | v19)) | 1)))
    fmt.Println(v3)
    v3 = v5
    var v24 uint64 = 39
    v24 -= v24
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
