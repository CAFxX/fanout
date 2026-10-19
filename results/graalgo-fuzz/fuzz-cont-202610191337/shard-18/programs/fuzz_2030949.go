package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 uint32 }
type GenS5 struct { F0 int32 }
func genFn0(p0 string, p1 int16, p2 int64) uint64 {
    switch 4294967295 {
    case -632:
        p0 = ""
    case 2147483647:
        switch 9223372036854775807 {
        case 2147483647:
            p1 = 1
        case -286:
            v4 := []int{2147483647, 127, 231}
            _ = v4
        }
    default:
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 int8 = -128
            var v7 uint64 = 0
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    p0 = p0
    v8 := GenS1{F0: 2, F1: p0, F2: 850}
    v9 := []int{830}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 int32 = -401
    var v2 uint16 = 65535
    var v3 uint16 = v2
    v10 := GenS2{F0: false, F1: genFn0(("a" + (("a" + "foo bar") + "\n\t")), 738, 1)}
    v3 = v2
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: false, F1: 42}
        for fi13 := 0; fi13 < 8; fi13++ {
            if true {
                fmt.Println(fi11, v2, fi11)
                var v14 uint64 = genFn0("foo bar", -113, -254)
                _ = v14
            } else {
                v2 += 42
            }
            _ = fi13
        }
        var v15 string = ("foo bar" + "")
        _ = fi11
        _ = v12
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS3{F0: (v1 | 0)}
        v18 := GenS2{F0: true, F1: 1}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v1, v3)
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    var v21 uint64 = 24
    v22 := GenS4{F0: (v1 + v1), F1: 9223372036854775807, F2: 414}
    var v23 uint = uint(v21)
    v1 += v1
    v24 := GenS5{F0: (-(v1))}
    v25 := []int{}
    fmt.Println(v21)
    v25 = append(v25, 764)
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint16 = (v2 | uint16(v21))
        var v28 int8 = 1
        fmt.Println(v23, v27, v2)
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = len(v25)
    v2 = 42
    for fi29 := 0; fi29 < 15; fi29++ {
        v30 := GenS5{F0: 0}
        fmt.Println(v3)
        var v31 int64 = -9223372036854775808
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
