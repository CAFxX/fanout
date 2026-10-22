package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 string }
type GenS3 struct { F0 uint32; F1 int; F2 uint16 }
func genFn0(p0 int8, p1 int, p2 bool) uint64 {
    v6 := GenS1{F0: 255, F1: 94}
    if (p1 != (p1 >> ((1000) & 63))) {
        v7 := []int{(p1 / ((p1) | 1)), p1}
        var v8 int16 = 32767
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS2{F0: p0, F1: p2, F2: "foo bar"}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := []int{(p1 >> ((255) & 63)), (fi11 + 42), fi11}
            var v13 uint64 = 4294967295
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    if true {
        v14 := GenS3{F0: 267, F1: (p1 >> ((2) & 63)), F2: 65535}
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS2{F0: 100, F1: true, F2: "foo bar"}
                _ = fi16
                _ = v17
            }
            v18 := []int{(-(p1)), (^(fi15)), p1, (p1 * 0)}
            var v19 int64 = -9223372036854775808
            _ = fi15
            _ = v18
            _ = v19
        }
        _ = v14
    } else {
        if p2 {
            var v20 uint32 = 0
            _ = v20
        } else {
            p0 = p0
            p0 -= p0
        }
        switch (p1 + -985) {
        case 32767:
            var v21 int = p1
            _ = v21
        case 128:
            v22 := []int{(-(p1)), 42, 9223372036854775807}
            _ = v22
        default:
            v23 := []int{42}
            _ = v23
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 uint = 42
            _ = fi24
            _ = v25
        }
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}

func main() {
    v1 := []uint32{1}
    var v2 float64 = ((((1e10 + 1.5) * 1e10) - ((1.5 * 0.0) + float64(-2.25))) / (-2.25 / 1.0))
    var v3 int16 = 1
    var v4 int = 1
    var v5 int16 = (-(v3))
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS3{F0: 820, F1: v4, F2: 32767}
        _ = fi26
        _ = v27
    }
    v5 = -553
    switch v4 {
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 2147483647:
        if true {
            var v28 uint = 255
            var v29 int8 = 100
            _ = v28
            _ = v29
        }
    case -819:
        var v30 int64 = 0
        _ = v30
    default:
        if false {
            fmt.Println(v4)
            var v31 float64 = v2
            _ = v31
        }
    }
    _ = len(v1)
    var v32 uint8 = 2
    var v33 int = 0
    for fi34 := 0; fi34 < 9; fi34++ {
        v32 = v32
        _ = fi34
    }
    if false {
        var v35 int8 = 1
        fmt.Println(v35)
        _ = v35
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v32, v2)
    case "hello":
        fmt.Println(v4, v32)
    default:
        v2 = v2
    }
    fmt.Println(v3, v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
