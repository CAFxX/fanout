package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 int; F2 int8 }
type GenS3 struct { F0 string; F1 int32; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 int, p1 uint8) bool {
    var v4 int8 = 100
    p1 -= 2
    _ = p0
    _ = p1
    _ = v4
    return false
}
func genFn1(p0 int16, p1 int32) int {
    var v5 uint64 = 42
    v5 = 2
    switch 4294967295 {
    case 127:
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(fi6, fi6, p1)
            if (("a" <= "") && false) {
                if ("" <= ("" + ("" + "\n\t"))) {
                    fmt.Println(fi6, p0, p1)
                    var v7 float64 = 1.5
                    v7 = 1e10
                    _ = v7
                } else {
                    p1 = p1
                }
                if false {
                    fmt.Println(p0, v5, p0)
                    v8 := []int{(fi6 % ((498) | 1)), 1}
                    var v9 float64 = -2.25
                    _ = v8
                    _ = v9
                } else {
                    if true {
                        var v10 float64 = 3.14159
                        v10 = v10
                        p0 -= p0
                        _ = v10
                    }
                }
            }
            var v11 int8 = -128
            _ = fi6
            _ = v11
        }
    default:
        p1 = -1
    }
    _ = p0
    _ = p1
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    switch "a" {
    case "foo bar":
        v12 := GenS1{F0: 9223372036854775807, F1: (-763 != (-(v1)))}
        _ = v12
    }
    fmt.Println(v1)
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 961, F1: (!(-1 <= -1))}
        var v15 int8 = 127
        _ = fi13
        _ = v14
        _ = v15
    }
    switch (v1 + v1) {
    case 32768:
        v16 := GenS2{F0: 42, F1: (v1 >> ((18446744073709551615) & 63)), F2: 0}
        _ = v16
    default:
        v1 -= v1
    }
    v17 := []int{2147483647, 42}
    v18 := GenS3{F0: ("hello" + "a"), F1: -1, F2: 2}
    if len(v17) > 0 {
        v17[0] = v1
    }
    v19 := GenS4{F0: 9223372036854775807, F1: 127, F2: 1000}
    v1 += genFn1(-1, -571)
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(v1, fi20)
        v21 := GenS3{F0: "hello", F1: -576, F2: 1}
        var v22 uint8 = 193
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int32 = -2
            _ = fi24
            _ = v25
        }
        if true {
            v26 := GenS5{F0: 819, F1: -921, F2: ("hello" + (("hello" + "") + ("" + "foo bar")))}
            if len(v17) > 0 {
                v17[0] = v1
            }
            _ = v26
        }
        switch v1 {
        case 1:
            v1 = v1
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
