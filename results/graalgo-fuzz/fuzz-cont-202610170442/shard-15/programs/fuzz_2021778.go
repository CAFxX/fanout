package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 bool }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 int64; F1 bool; F2 string }
type GenS4 struct { F0 int16 }

func main() {
    v1 := []string{}
    var v2 int64 = 0
    var v3 int8 = 100
    var v4 int16 = -272
    var v5 uint = 112
    v6 := GenS1{F0: 0, F1: false, F2: true}
    v4 = -1
    if false {
        var v7 bool = false
        _ = v7
    } else {
        var v8 uint16 = 394
        v3 -= (v3 ^ v3)
        fmt.Println(v3)
        _ = v8
    }
    var v9 uint32 = 1
    var v10 uint64 = 4294967295
    v11 := GenS2{F0: false, F1: (v4 ^ (v4 << (((v5 << ((0) & 63))) & 63))), F2: v5}
    for fi12 := 0; fi12 < 7; fi12++ {
        switch "" {
        case "foo bar":
            v13 := GenS1{F0: v5, F1: true, F2: (("" + ("" + "")) < "")}
            _ = v13
        case "":
            var v14 int8 = 127
            _ = v14
        default:
            v15 := GenS3{F0: v2, F1: true, F2: "a"}
            _ = v15
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS3{F0: v2, F1: true, F2: "foo bar"}
            _ = fi16
            _ = v17
        }
        _ = fi12
    }
    v5 += uint(v10)
    if false {
        v18 := GenS4{F0: (-(v4))}
        var v19 float64 = 3.14159
        _ = v18
        _ = v19
    } else {
        v1 = append(v1, "\n\t")
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v3 -= (-(v3))
        v2 = int64(fi20)
        _ = fi20
    }
    switch ("hello" + "") {
    case "a":
        _ = len(v1)
    default:
        if false {
            var v21 int64 = (v2 + (v2 >> (((v5 ^ 303)) & 63)))
            v1 = append(v1, "")
            var v22 int32 = 42
            _ = v21
            _ = v22
        }
    }
    var v23 uint16 = 1
    var v24 int32 = -375
    v1 = append(v1, ("\n\t" + ("foo bar" + "foo bar")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
