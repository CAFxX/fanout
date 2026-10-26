package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int }

func main() {
    var v1 uint8 = 127
    var v2 float64 = 3.14159
    var v3 int16 = 130
    var v4 int16 = v3
    var v5 int8 = 42
    if true {
        var v6 int = -940
        v7 := []int{v6, 9223372036854775807}
        _ = len(v7)
        _ = v6
        _ = v7
    }
    fmt.Println(v3, v5)
    v1 -= (v1 * 163)
    var v8 uint32 = 745
    if (((v1 / (((v1 / (((v1 << ((2) & 63))) | 1))) | 1)) >= v1) && false) {
        fmt.Println(v1, v5)
    } else {
        if (42 > -774) {
            var v9 float64 = ((v2 * v2) + v2)
            var v10 uint32 = v8
            v11 := GenS1{F0: 996, F1: 1}
            _ = v10
            _ = v11
            _ = v9
        } else {
            if (v5 == (v5 | -128)) {
                if (-9223372036854775808 > 13) {
                    var v12 float64 = ((-2.25 * 1e10) - 0.0)
                    _ = v12
                } else {
                    var v13 string = "foo bar"
                    _ = v13
                }
            } else {
                var v14 int = 32768
                v15 := []int{v14, (v14 % (((v14 & (^(v14)))) | 1))}
                _ = v14
                _ = v15
            }
            v5 = (v5 / ((v5) | 1))
            var v16 uint64 = 18446744073709551615
            _ = v16
        }
        var v17 uint32 = (v8 / ((v8) | 1))
        _ = v17
    }
    var v18 int8 = (v5 - 127)
    var v19 int8 = v18
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int = (fi20 * fi20)
        var v22 int64 = 919
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: true}
    fmt.Println(v19)
    fmt.Println(v8, v4, v19)
    switch 1 {
    case -1:
        v24 := GenS3{F0: 128, F1: 32768}
        _ = v24
    case 128:
        fmt.Println(v5)
    }
    fmt.Println(v4, v1)
    var v25 int64 = -9223372036854775808
    if ((!false) || true) {
        v26 := GenS2{F0: (!true)}
        _ = v26
    } else {
        v3 = 848
        v8 = 2147483647
    }
    switch "" {
    case "a":
        var v27 uint = 65535
        _ = v27
    case "\n\t":
        switch ("a" + ("hello" + "hello")) {
        case "hello":
            v28 := []int{}
            _ = v28
        }
    default:
        v29 := []int{int(v8), 4294967295}
        _ = v29
    }
    v4 = (v4 % (((^(v3))) | 1))
    var v30 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
