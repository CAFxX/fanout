package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
func genFn0(p0 int16, p1 uint) uint {
    v4 := []int{-2147483648}
    switch "" {
    case "foo bar":
        switch "hello" {
        case "hello":
            fmt.Println(p0)
        case "foo bar":
            if false {
                p0 = 32767
            }
        default:
            _ = len(v4)
        }
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = 32767
        }
    }
    if ((p0 / ((774) | 1)) < 42) {
        v5 := GenS1{F0: -38, F1: -101, F2: 42}
        if false {
            v6 := GenS1{F0: -112, F1: 42, F2: 1}
            fmt.Println(p1, p1)
            _ = v6
        }
        var v7 uint32 = 2147483648
        _ = v5
        _ = v7
    } else {
        p0 = (p0 + p0)
        var v8 uint64 = 42
        _ = v8
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v4
    return 684
}
func genFn1() uint {
    var v9 uint = 4294967295
    fmt.Println(v9)
    _ = v9
    return genFn0(619, (v9 + v9))
}

func main() {
    var v1 int16 = 511
    var v2 int8 = 1
    var v3 int64 = 9223372036854775807
    if true {
        v10 := GenS1{F0: -630, F1: -2147483648, F2: 762}
        _ = v10
    }
    var v11 uint32 = 632
    v12 := []int{1}
    var v13 bool = false
    v14 := GenS1{F0: v3, F1: 304, F2: 18446744073709551615}
    v15 := GenS1{F0: (v3 + (v3 + v3)), F1: 1, F2: 0}
    if v13 {
        var v16 int16 = v1
        var v17 string = ("" + "\n\t")
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 42, F1: -2147483648, F2: 0}
    fmt.Println(v11, v11, v13)
    v19 := GenS1{F0: (v3 >> ((710) & 63)), F1: 0, F2: 622}
    var v20 bool = v13
    var v21 int = 2147483647
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
