package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0() int16 {
    v5 := []int{-1, 32768, -9223372036854775808, 4294967295}
    var v6 float32 = 1.5
    _ = v5
    _ = v6
    return 700
}

func main() {
    var v1 uint64 = 2
    var v2 int32 = 1
    var v3 int8 = -1
    var v4 int8 = 0
    var v7 uint8 = 0
    var v8 uint = 42
    if true {
        var v9 uint64 = 42
        var v10 int32 = (v2 % ((v2) | 1))
        var v11 uint32 = 2
        _ = v10
        _ = v11
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v8, v1)
    case "":
        v12 := GenS1{F0: (v7 + (^(v7))), F1: 32767, F2: (v8 * 18446744073709551615)}
        _ = v12
    case "\n\t":
        var v13 uint32 = 1
        _ = v13
    default:
        var v14 uint = (v8 ^ (v8 + (v8 | (v8 / (((^(v8))) | 1)))))
        _ = v14
    }
    switch 32768 {
    case 32767:
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := []int{fi15, fi15, fi15}
            v17 := GenS1{F0: uint8(v8), F1: 2, F2: v8}
            if true {
                v4 = (v4 << ((545) & 63))
            } else {
                v7 += (v7 << ((v8) & 63))
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    case -434:
        if ("\n\t" == "hello") {
            v18 := []int{-874}
            v18 = append(v18, 944)
            v8 -= v8
            _ = v18
        }
    case 9223372036854775807:
        for fi19 := 0; fi19 < 9; fi19++ {
            if (42 <= genFn0()) {
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 uint32 = 2
                    v3 = v4
                    _ = fi20
                    _ = v21
                }
            }
            v3 -= -128
            _ = fi19
        }
    }
    v22 := GenS2{F0: 2147483647, F1: v7}
    v23 := GenS1{F0: (v7 << ((195) & 63)), F1: 2, F2: (v8 << ((748) & 63))}
    v24 := []int{128, -231, -309, 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
