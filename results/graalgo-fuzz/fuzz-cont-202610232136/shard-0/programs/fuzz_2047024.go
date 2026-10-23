package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint64; F2 string }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 bool) uint {
    var v5 float64 = 0.1
    switch 568 {
    case 2147483647:
        var v6 uint64 = 1
        _ = v6
    case 1:
        p0 = p0
    default:
        switch 42 {
        case 8:
            fmt.Println(p1, v5, v5)
        }
    }
    p0 = (p0 ^ 909)
    v5 -= (float64(-2.25) * float64((3.14159 + 3.14159)))
    if (((4294967295 < 2) || p1) && (!true)) {
        p1 = ((v5 - float64((-2.25 / -2.25))) == v5)
        p1 = false
    } else {
        v7 := []int{}
        var v8 float64 = (0.0 * (v5 / float64((1e10 + float32(v5)))))
        fmt.Println(v8, p0, p0)
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1() uint64 {
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: -1}
        _ = fi9
        _ = v10
    }
    var v11 int = 700
    var v12 string = "hello"
    v13 := GenS2{F0: 255, F1: 17}
    v14 := GenS1{F0: 42}
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 379
}

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 2
    var v3 bool = false
    var v4 int64 = 9223372036854775807
    var v15 int = int(v1)
    var v16 bool = (!((!v3) || v3))
    v2 -= (v2 % (((v2 >> ((652) & 63))) | 1))
    v17 := GenS3{F0: 2147483647, F1: 1, F2: "a"}
    if (float64(float32(((1e10 - 1.5) - 0.0))) == 0.0) {
        var v18 bool = v16
        _ = v18
    } else {
        var v19 float64 = 1.5
        var v20 bool = v16
        _ = v19
        _ = v20
    }
    var v21 uint16 = 222
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := []int{}
        _ = fi22
        _ = v23
    }
    if false {
        v2 -= (v2 * v2)
        v24 := GenS4{F0: -1}
        _ = v24
    }
    fmt.Println(v16, v16, v2)
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS1{F0: -1}
        v2 = v2
        _ = fi25
        _ = v26
    }
    v16 = v3
    var v27 uint16 = v21
    v27 -= (v27 + 2)
    v2 -= v2
    fmt.Println(v27, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
