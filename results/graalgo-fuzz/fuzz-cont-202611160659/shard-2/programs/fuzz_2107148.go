package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
type GenS4 struct { F0 int }
func genFn0() uint16 {
    if false {
        var v6 uint32 = 562
        var v7 uint8 = 255
        v6 += 4294967295
        _ = v6
        _ = v7
    }
    var v8 float32 = 3.14159
    var v9 bool = (100 >= 100)
    v9 = (65535 == 2147483647)
    var v10 uint16 = 1
    _ = v10
    _ = v8
    _ = v9
    return (v10 / (((v10 << ((99) & 63))) | 1))
}

func main() {
    var v1 float32 = 0.0
    var v2 string = ("foo bar" + "")
    var v3 int32 = -284
    var v4 int16 = 42
    v5 := []int{127}
    switch "a" {
    case "\n\t":
        v3 = v3
    case "foo bar":
        switch 32767 {
        case -778:
            if (!((false || true) && true)) {
                v11 := GenS1{F0: 985, F1: ("\n\t" + ("" + (v2 + (v2 + v2)))), F2: 0}
                var v12 uint8 = 1
                _ = v11
                _ = v12
            }
        }
    case "a":
        fmt.Println(v1, v3)
    }
    v2 = v2
    var v13 int = 0
    v1 += v1
    switch "a" {
    case "":
        v14 := GenS1{F0: 2147483648, F1: "hello", F2: 1}
        _ = v14
    default:
        _ = len(v5)
    }
    if (!true) {
        if true {
            v15 := GenS2{F0: 1}
            v1 -= 0.0
            _ = v15
        }
    } else {
        v13 -= 42
    }
    var v16 uint32 = 2
    if ((false && (1 < v13)) && true) {
        v17 := GenS1{F0: v16, F1: v2, F2: -52}
        _ = v17
    }
    var v18 int64 = -749
    v19 := GenS3{F0: v3, F1: v16, F2: (v18 / (((v18 / (((v18 % ((v18) | 1))) | 1))) | 1))}
    v16 += 879
    fmt.Println(v3)
    if false {
        var v20 uint32 = 813
        _ = v20
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!(!(v2 != v2))) {
        fmt.Println(v16)
    } else {
        var v21 uint64 = 915
        if (true || true) {
            var v22 int = 1
            v4 = (v4 * (-(v4)))
            _ = v22
        } else {
            if true {
                v21 = 473
                fmt.Println(v13, v21)
            } else {
                var v23 int = 128
                var v24 int32 = (v3 * -165)
                _ = v23
                _ = v24
            }
        }
        v2 = v2
        _ = v21
    }
    v25 := GenS4{F0: v13}
    for fi26 := 0; fi26 < 16; fi26++ {
        v13 += fi26
        v13 += (fi26 % (((fi26 * v13)) | 1))
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
