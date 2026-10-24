package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 uint) int32 {
    var v5 uint = (p0 ^ p0)
    p0 = 2
    fmt.Println(v5, p0)
    _ = p0
    _ = v5
    return 42
}
func genFn1() uint64 {
    var v6 uint16 = 32767
    if true {
        if (!((-100 < -32768) && true)) {
            switch "" {
            case "\n\t":
                var v7 int = 1
                _ = v7
            case "hello":
                if ((!false) && (127 == 255)) {
                    var v8 int64 = -624
                    if true {
                        var v9 uint = 407
                        v9 += 18446744073709551615
                        _ = v9
                    }
                    _ = v8
                }
            }
            v10 := GenS1{F0: 2147483648, F1: genFn0(216), F2: 4294967295}
            _ = v10
        }
    } else {
        v11 := GenS2{F0: -2147483648}
        _ = v11
    }
    var v12 int32 = -609
    _ = v12
    _ = v6
    return 2
}

func main() {
    var v1 int32 = 0
    var v2 uint = 2
    var v3 uint8 = 251
    var v4 int16 = 4
    if false {
        fmt.Println(v4)
    }
    v13 := []int{-9223372036854775808, 486}
    v4 -= (v4 ^ (v4 % ((v4) | 1)))
    if (!false) {
        v1 = v1
        _ = len(v13)
        v2 += 403
    }
    v14 := GenS1{F0: 442, F1: v1, F2: 4294967295}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    switch "a" {
    case "a":
        v1 = v1
    case "\n\t":
        var v15 int32 = (v1 & (v1 << (((v2 << (((v2 << ((v2) & 63))) & 63))) & 63)))
        _ = v15
    default:
        var v16 int8 = 87
        _ = v16
    }
    if false {
        var v17 string = ("" + "")
        v18 := GenS3{F0: (v17 + "\n\t")}
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
