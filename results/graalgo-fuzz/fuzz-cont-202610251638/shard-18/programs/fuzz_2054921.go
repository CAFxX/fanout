package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint; F2 int }
func genFn0() uint64 {
    var v4 float32 = 0.0
    if (-128 < -86) {
        v5 := []int{-9223372036854775808, -2147483648}
        v4 -= -2.25
        v5 = append(v5, 32767)
        _ = v5
    }
    var v6 uint8 = 128
    _ = v4
    _ = v6
    return 357
}

func main() {
    v1 := []uint32{0, 2147483647}
    var v2 uint16 = 42
    var v3 string = "\n\t"
    if len(v1) > 0 {
        v1[0] = 147
    }
    v3 = v3
    _ = len(v1)
    v7 := GenS1{F0: -32, F1: -9223372036854775808, F2: 918}
    var v8 int = 2147483647
    var v9 float32 = (1e10 * (((3.14159 + 3.14159) * 0.1) / -2.25))
    fmt.Println(v3, v9, v8)
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            var v11 bool = (v2 != (v2 >> ((42) & 63)))
            v8 -= 939
            fmt.Println(v8)
            _ = v11
        }
        var v12 int16 = -480
        _ = fi10
        _ = v12
    }
    v3 = v3
    var v13 string = "hello"
    if (v9 != v9) {
        if false {
            v3 = "\n\t"
            v14 := GenS2{F0: uint(v2), F1: 134, F2: 18446744073709551615}
            var v15 int8 = -79
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                if ((-1 < 128) && true) {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 int = 32767
                        _ = fi17
                        _ = v18
                    }
                    v2 -= (^(v2))
                }
                _ = fi16
            }
            var v19 int32 = -2147483648
            _ = v19
        }
        var v20 uint8 = 0
        fmt.Println(v9)
        _ = v20
    } else {
        v21 := GenS3{F0: 153, F1: 1, F2: 0}
        v2 = 1
        if false {
            var v22 string = (((v13 + "a") + "a") + v3)
            _ = v22
        } else {
            fmt.Println(v8, v2, v2)
            v23 := GenS2{F0: 1, F1: -605, F2: genFn0()}
            _ = v23
        }
        _ = v21
    }
    if (!true) {
        var v24 uint16 = 87
        _ = v24
    }
    var v25 float64 = (-2.25 * (0.0 / 1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
