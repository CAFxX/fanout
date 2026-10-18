package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
func genFn0(p0 uint32, p1 bool, p2 uint16) int8 {
    var v4 uint16 = (p2 << ((255) & 63))
    if (0 > 775) {
        fmt.Println(p1, p0, p0)
        var v5 string = ""
        fmt.Println(v5)
        _ = v5
    }
    if p1 {
        v6 := []int{209, -918, 713}
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            var v8 uint = 4294967295
            var v9 int = fi7
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 0, F1: 9223372036854775807, F2: 127}
            _ = v10
        }
        var v11 int = -9223372036854775808
        _ = fi7
        _ = v11
    }
    if (!p1) {
        v12 := []int{}
        v4 -= p2
        v13 := GenS1{F0: -877, F1: int64(p2), F2: -128}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 uint64, p1 string) int8 {
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: -9223372036854775808, F1: 1, F2: -11}
        _ = fi14
        _ = v15
    }
    v16 := []int{}
    _ = p0
    _ = p1
    _ = v16
    return genFn0(219, false, 32767)
}

func main() {
    v1 := []int{}
    var v2 uint16 = 2
    v3 := []uint32{}
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    var v17 int16 = -32768
    v17 = (v17 & (v17 | v17))
    v17 = v17
    if false {
        _ = len(v3)
        v2 = 806
        var v18 string = (("hello" + "a") + "a")
        _ = v18
    } else {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        v2 = v2
        switch "a" {
        case "foo bar":
            v17 -= v17
        default:
            var v19 uint = 1
            _ = v19
        }
    }
    if false {
        var v20 uint16 = 42
        _ = v20
    } else {
        switch "" {
        case "a":
            if (v17 >= v17) {
                var v21 float32 = (1.5 / (-2.25 + 0.1))
                _ = v21
            }
        case "":
            v17 += v17
        default:
            var v22 uint32 = 2
            _ = v22
        }
    }
    _ = len(v1)
    fmt.Println(v2)
    switch 128 {
    case 9223372036854775807:
        switch "" {
        case "foo bar":
            switch "hello" {
            case "hello":
                v2 -= (v2 << ((477) & 63))
            case "foo bar":
                var v23 float64 = 3.14159
                _ = v23
            case "\n\t":
                v2 = 0
            }
        case "\n\t":
            fmt.Println(v2, v2)
        case "":
            var v24 int64 = 1
            _ = v24
        }
    default:
        var v25 int8 = 44
        _ = v25
    }
    var v26 uint64 = 42
    if ((42 < 1) && true) {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
}
