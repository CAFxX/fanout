package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS4 struct { F0 int16; F1 bool; F2 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int16, p1 int16) int {
    var v7 int64 = 9223372036854775807
    v7 += 1
    _ = p0
    _ = p1
    _ = v7
    return 769
}
func genFn1(p0 int8, p1 uint64) int32 {
    p1 = p1
    p1 += p1
    p1 += p1
    var v8 float32 = 0.1
    p1 += (p1 | p1)
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 uint16 = 2
    var v4 uint16 = (v3 - v3)
    var v5 int64 = -69
    var v6 bool = true
    if v6 {
        var v9 float64 = -2.25
        if v6 {
            v6 = v6
        } else {
            v10 := GenS1{F0: 128}
            _ = v10
        }
        v11 := GenS1{F0: 227}
        _ = v11
        _ = v9
    } else {
        v12 := GenS2{F0: 1, F1: "foo bar"}
        _ = v12
    }
    if v6 {
        var v13 float32 = ((-2.25 + ((0.0 / 1.5) + -2.25)) / 1.5)
        _ = v13
    } else {
        if v6 {
            v5 += v5
            v14 := GenS3{F0: 1, F1: (!false), F2: 569}
            var v15 bool = v6
            _ = v14
            _ = v15
        }
        v3 -= 0
        v16 := GenS2{F0: 42, F1: "a"}
        _ = v16
    }
    v1["a"] = false
    if v6 {
        v17 := GenS2{F0: 2, F1: "foo bar"}
        _ = v17
    } else {
        delete(v1, (("foo bar" + "foo bar") + "foo bar"))
        switch 2147483647 {
        case -569:
            v2 = v5
        default:
            for fi18 := 0; fi18 < 12; fi18++ {
                if v6 {
                    v4 = (^(v3))
                    fmt.Println(v4, v4, v6)
                }
                _ = fi18
            }
        }
        fmt.Println(v2)
    }
    v5 = (^(v5))
    v6 = v6
    v2 = v5
    if v6 {
        v4 += (v3 ^ v4)
        var v19 uint8 = 128
        _ = v19
    } else {
        v4 = (v3 << ((367) & 63))
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS2{F0: 167, F1: "hello"}
        _ = len(v1)
        var v22 float32 = 0.0
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 float64 = 1e10
    switch genFn0(-1, -910) {
    case 4294967295:
        v24 := GenS1{F0: 0}
        _ = v24
    default:
        var v25 uint64 = 18446744073709551615
        _ = v25
    }
    fmt.Println(v23)
    if v6 {
        v26 := GenS4{F0: 1000, F1: (!(!true)), F2: v2}
        _ = v26
    }
    v1[(("\n\t" + ("\n\t" + "hello")) + "")] = ("foo bar" != "foo bar")
    var v27 uint32 = 830
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := GenS5{F0: genFn1(42, 386)}
        fmt.Println(v27, v4, v23)
        var v30 int32 = 234
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
