package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 int16) int16 {
    var v5 int16 = p0
    switch 0 {
    case -262:
        switch "a" {
        case "":
            fmt.Println(p0, v5)
        case "foo bar":
            if (!true) {
                for fi6 := 0; fi6 < 7; fi6++ {
                    p0 = int16(fi6)
                    var v7 int8 = 1
                    for fi8 := 0; fi8 < 10; fi8++ {
                        fmt.Println(v5)
                        _ = fi8
                    }
                    _ = fi6
                    _ = v7
                }
                v9 := GenS1{F0: ("\n\t" + (("" + ("\n\t" + ("a" + "foo bar"))) + ("\n\t" + "\n\t"))), F1: "foo bar"}
                var v10 uint64 = 687
                _ = v10
                _ = v9
            }
        case "a":
            v11 := []int{}
            _ = v11
        default:
            var v12 int16 = v5
            _ = v12
        }
    case -176:
        var v13 int16 = -336
        _ = v13
    }
    switch "hello" {
    case "\n\t":
        if ((v5 >> ((0) & 63)) <= v5) {
            v5 = (-(v5))
            v14 := GenS1{F0: ("a" + ""), F1: ("\n\t" + ("a" + ""))}
            _ = v14
        } else {
            v15 := []int{}
            v16 := GenS2{F0: 1}
            _ = len(v15)
            _ = v15
            _ = v16
        }
    case "foo bar":
        switch "hello" {
        case "\n\t":
            v17 := []int{-9223372036854775808, 128, -778, 127}
            _ = v17
        }
    case "hello":
        var v18 int32 = -767
        _ = v18
    }
    _ = p0
    _ = v5
    return (v5 * p0)
}
func genFn1(p0 uint8) uint8 {
    v19 := []int{-211, 32768}
    if len(v19) > 0 {
        v19[0] = 127
    }
    p0 = (^(p0))
    var v20 int8 = 100
    _ = p0
    _ = v19
    _ = v20
    return (p0 | 42)
}

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = (v1 + -657)
    v3 := make(map[uint32]int)
    var v4 uint32 = 4294967295
    v4 -= (v4 | v4)
    v21 := GenS2{F0: 1}
    v22 := GenS2{F0: -1}
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS1{F0: "foo bar", F1: "\n\t"}
        switch int(v4) {
        case -420:
            v2 -= -2147483648
        default:
            v25 := GenS3{F0: ("hello" + "\n\t"), F1: "\n\t"}
            _ = v25
        }
        v26 := []int{}
        _ = fi23
        _ = v24
        _ = v26
    }
    v1 -= v2
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 int8 = -128
        var v29 int8 = int8(v2)
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(v1)
        _ = fi30
    }
    if true {
        v4 = 65535
        v31 := []int{}
        _ = v31
    }
    v32 := GenS2{F0: -820}
    var v33 int32 = (v2 + 0)
    _ = len(v3)
    fmt.Println(v1, v2, v2)
    v34 := []int{0, 4294967295, -784}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
