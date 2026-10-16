package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
type GenS3 struct { F0 int32; F1 int8; F2 int8 }
type GenS4 struct { F0 int16; F1 int16; F2 string }

func main() {
    var v1 uint = 65535
    var v2 float64 = float64(-2.25)
    v3 := []uint32{780}
    v4 := []int{32767}
    var v5 float64 = 3.14159
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 42}
        switch "" {
        case "\n\t":
            v2 = (0.1 - v2)
        case "":
            v8 := GenS2{F0: 0, F1: true, F2: 293}
            _ = v8
        default:
            var v9 float64 = v2
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    switch "\n\t" {
    case "\n\t":
        var v10 uint8 = 124
        _ = v10
    }
    if true {
        if (3.14159 != -2.25) {
            var v11 uint16 = 1
            _ = len(v3)
            v12 := GenS2{F0: 42, F1: (!((("" + ((("a" + "") + (("a" + "\n\t") + "hello")) + "foo bar")) + "") < ("foo bar" + "a"))), F2: 1}
            _ = v11
            _ = v12
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v13 float64 = 1e10
    if true {
        v5 = v2
        v14 := GenS3{F0: -795, F1: -1, F2: 100}
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 int16 = 1000
            var v17 uint64 = 214
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    switch 32768 {
    case -685:
        v4 = append(v4, -404)
    case 2147483647:
        var v18 float64 = (v2 / float64((3.14159 * 1e10)))
        _ = v18
    case 1:
        if (-1 != 100) {
            v19 := GenS1{F0: 352}
            fmt.Println(v2, v1)
            var v20 bool = false
            _ = v19
            _ = v20
        } else {
            _ = len(v3)
        }
    }
    v4 = append(v4, -9223372036854775808)
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint64 = 64
        v2 -= (v13 + v13)
        _ = fi21
        _ = v22
    }
    switch "a" {
    case "":
        var v23 uint = v1
        _ = v23
    default:
        v2 -= ((v5 / v2) + (3.14159 - ((v2 * 1e10) * (v13 * v2))))
    }
    if len(v3) > 0 {
        v3[0] = 923
    }
    v24 := GenS4{F0: -823, F1: 365, F2: "\n\t"}
    fmt.Println(v5, v1, v2)
    var v25 int8 = 127
    v13 -= v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
