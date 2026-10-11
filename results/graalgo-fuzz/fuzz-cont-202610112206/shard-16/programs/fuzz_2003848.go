package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
func genFn0() int32 {
    v4 := GenS1{F0: 0, F1: -68}
    v5 := GenS2{F0: 0, F1: 18446744073709551615, F2: -9223372036854775808}
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 float32 = 3.14159
        _ = fi6
        _ = v7
    }
    switch 1 {
    case -2147483648:
        var v8 string = "a"
        _ = v8
    default:
        var v9 uint32 = 2147483648
        _ = v9
    }
    var v10 int64 = 1
    _ = v10
    _ = v4
    _ = v5
    return -2147483648
}
func genFn1() bool {
    var v11 int32 = 2147483647
    var v12 uint = 42
    switch -681 {
    case 32767:
        fmt.Println(v12)
    case 1:
        v11 = (v11 - genFn0())
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: 2, F1: 872, F2: -445}
        v11 += v11
        for fi15 := 0; fi15 < 15; fi15++ {
            v11 = v11
            var v16 int = fi15
            var v17 int16 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 int16 = -112
    var v2 float32 = (1e10 * (0.1 - 1e10))
    v3 := []uint32{2147483647, 728}
    v1 = (v1 & v1)
    v1 = (v1 % (((v1 + v1)) | 1))
    var v18 int32 = 0
    var v19 bool = true
    var v20 int32 = v18
    fmt.Println(v1, v20)
    v19 = v19
    var v21 float32 = (v2 / 1e10)
    var v22 float32 = v21
    v21 = ((v2 / v22) + v22)
    v23 := GenS2{F0: 1, F1: 18446744073709551615, F2: 752}
    v2 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
