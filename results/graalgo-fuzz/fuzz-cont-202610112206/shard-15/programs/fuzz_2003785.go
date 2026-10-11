package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
func genFn0() uint8 {
    v5 := []int{-780, -2147483648}
    switch -1 {
    case 42:
        var v6 uint16 = 868
        _ = v6
    default:
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    var v8 int8 = -128
    switch 9223372036854775807 {
    case -1:
        v9 := GenS1{F0: -689}
        _ = v9
    default:
        v8 = v8
    }
    _ = v5
    _ = v8
    return 128
}
func genFn1(p0 int16, p1 int) uint16 {
    if false {
        v10 := []int{p1, p1, -9223372036854775808, (p1 >> ((1) & 63))}
        _ = v10
    } else {
        v11 := []int{(p1 & (^(p1))), p1, p1}
        fmt.Println(p0, p1)
        var v12 float32 = 3.14159
        _ = v11
        _ = v12
    }
    var v13 uint32 = 65535
    _ = p0
    _ = p1
    _ = v13
    return 255
}

func main() {
    var v1 uint16 = 0
    var v2 int8 = 42
    v3 := make(map[int]string)
    var v4 int32 = 2147483647
    var v14 int16 = 453
    var v15 uint8 = 255
    for fi16 := 0; fi16 < 11; fi16++ {
        v1 += genFn1(42, fi16)
        switch fi16 {
        case 2147483647:
            if ((!("hello" <= "")) && true) {
                v1 += genFn1(1, (fi16 - 42))
                fmt.Println(v15, v15)
            } else {
                v1 = v1
            }
        case 1:
            v17 := GenS2{F0: (v15 >> ((255) & 63))}
            _ = v17
        default:
            v18 := GenS2{F0: v15}
            _ = v18
        }
        _ = fi16
    }
    var v19 uint64 = 4294967295
    v20 := GenS2{F0: v15}
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 bool = true
        v23 := GenS1{F0: (v14 * (v14 ^ (-(v14))))}
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 uint32 = 2147483647
    v25 := GenS3{F0: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
