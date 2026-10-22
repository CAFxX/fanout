package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int8; F2 int32 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0() bool {
    var v5 uint32 = 1
    v6 := []int{-959, -2147483648}
    if (32767 >= -1) {
        var v7 int16 = 442
        _ = v7
    }
    _ = v5
    _ = v6
    return ("\n\t" > "foo bar")
}
func genFn1(p0 uint) int {
    if ((313 <= 32767) && (!(false && false))) {
        v8 := GenS1{F0: "foo bar"}
        var v9 uint64 = 25
        v10 := GenS1{F0: (("" + "hello") + ("hello" + "\n\t"))}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if false {
            fmt.Println(p0, fi11)
            p0 += (p0 * 347)
            p0 = (p0 | (p0 * (^(p0))))
        } else {
            v12 := []int{(fi11 << (((p0 % ((1) | 1))) & 63)), fi11, fi11, (fi11 * fi11)}
            _ = v12
        }
        var v13 int8 = 1
        switch "a" {
        case "":
            v14 := GenS1{F0: ""}
            _ = v14
        case "hello":
            v15 := GenS1{F0: ""}
            _ = v15
        }
        _ = fi11
        _ = v13
    }
    v16 := GenS2{F0: false, F1: -1, F2: -242}
    _ = p0
    _ = v16
    return 2147483647
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 42
    var v3 int32 = 2147483647
    var v4 int32 = (v3 >> ((98) & 63))
    v17 := GenS3{F0: -2147483648, F1: 0}
    v18 := GenS1{F0: "\n\t"}
    v19 := []int{-96, -9223372036854775808}
    v20 := GenS1{F0: (("hello" + "\n\t") + ("" + (("hello" + "foo bar") + "")))}
    v3 = v3
    v2 = -128
    v1 += ((v1 * ((1.5 + (((v1 + v1) + v1) * v1)) - v1)) * 0.0)
    var v21 uint = 375
    fmt.Println(v2, v1, v1)
    _ = len(v19)
    v3 = v3
    switch "" {
    case "foo bar":
        var v22 int8 = (v2 / ((v2) | 1))
        _ = v22
    case "a":
        v19 = append(v19, 128)
    case "\n\t":
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS3{F0: v3, F1: 2}
            _ = fi23
            _ = v24
        }
    default:
        var v25 uint = 42
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
