package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
func genFn0(p0 uint, p1 int16, p2 int8) int32 {
    if (false || false) {
        if (859 == p0) {
            v6 := GenS1{F0: 1, F1: 0, F2: (p2 & p2)}
            var v7 int16 = (p1 / ((p1) | 1))
            _ = v6
            _ = v7
        }
        v8 := GenS2{F0: p0, F1: ((1e10 >= (1.5 / -2.25)) && ((p1 | 1000) < (p1 + (p1 * (p1 ^ (^(p1))))))), F2: p2}
        _ = v8
    }
    p2 = -128
    v9 := GenS2{F0: p0, F1: true, F2: 42}
    p1 = 0
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 int64 = -521
    var v2 uint = 2
    var v3 float32 = 0.0
    var v4 uint32 = 2147483647
    v5 := make(map[int]uint32)
    switch "foo bar" {
    case "\n\t":
        if false {
            if true {
                v10 := GenS1{F0: uint16(v2), F1: (v4 | (v4 - (v4 & 2))), F2: 127}
                var v11 int16 = 32767
                _ = v10
                _ = v11
            }
            switch "a" {
            case "\n\t":
                v12 := []int{-9223372036854775808, 0}
                _ = v12
            case "a":
                v1 = v1
            }
            v3 = (v3 - 0.1)
        }
    }
    switch 128 {
    case -647:
        var v13 uint8 = 54
        _ = v13
    case 4294967295:
        v1 = -1
    case 1:
        v14 := []int{}
        _ = v14
    default:
        v15 := GenS2{F0: v2, F1: true, F2: 127}
        _ = v15
    }
    _ = len(v5)
    var v16 float64 = 3.14159
    if ((true && true) && true) {
        var v17 bool = true
        _ = v17
    }
    var v18 int64 = v1
    if false {
        fmt.Println(v1, v3)
        v19 := v5[4294967295]
        fmt.Println(v3, v18)
        _ = v19
    } else {
        switch ("\n\t" + "hello") {
        case "foo bar":
            var v20 float64 = v16
            _ = v20
        case "":
            var v21 int64 = -263
            _ = v21
        case "\n\t":
            v3 -= v3
        default:
            v18 += 1
        }
        v18 -= (v1 % ((v18) | 1))
        switch -236 {
        case 0:
            v22 := GenS1{F0: 172, F1: v4, F2: 42}
            _ = v22
        case 127:
            fmt.Println(v2, v2, v1)
        case 9223372036854775807:
            v23 := []int{9223372036854775807, -2147483648}
            _ = v23
        default:
            v24 := []int{42, 4294967295}
            _ = v24
        }
    }
    var v25 float32 = -2.25
    var v26 string = (((("\n\t" + (("\n\t" + "") + ("a" + "foo bar"))) + "\n\t") + ("\n\t" + (("foo bar" + (("\n\t" + "foo bar") + "\n\t")) + ""))) + "")
    v27 := []int{4294967295}
    var v28 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
