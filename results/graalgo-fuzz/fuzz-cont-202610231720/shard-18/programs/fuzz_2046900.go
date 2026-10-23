package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
func genFn0(p0 bool) int64 {
    v6 := GenS1{F0: p0, F1: 2}
    if ((!((!(true && p0)) && ((((1e10 + 1.5) * 0.0) + (1e10 - 3.14159)) < (3.14159 + -2.25)))) || p0) {
        if (-1 <= 406) {
            var v7 uint64 = 42
            v7 = v7
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int8 = 42
                var v10 uint32 = 2147483648
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            fmt.Println(p0)
            p0 = ((-2.25 + 3.14159) > 3.14159)
            var v11 int64 = 9223372036854775807
            _ = v11
        }
    } else {
        switch 404 {
        case 128:
            var v12 int16 = 676
            _ = v12
        default:
            p0 = p0
        }
        v13 := GenS1{F0: (-949 != 0), F1: 286}
        p0 = (!p0)
        _ = v13
    }
    v14 := GenS1{F0: false, F1: 2}
    if p0 {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    var v15 int8 = 1
    _ = p0
    _ = v14
    _ = v15
    _ = v6
    return 984
}
func genFn1(p0 int16, p1 int16) uint {
    if true {
        switch 1 {
        case -223:
            var v16 uint = 0
            _ = v16
        case 0:
            var v17 int8 = -56
            _ = v17
        case 9223372036854775807:
            var v18 int = 0
            _ = v18
        default:
            if ((p0 ^ (-(p0))) == 617) {
                p0 -= p0
            }
        }
        v19 := []int{-233}
        _ = v19
    }
    p0 += p1
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint16 = 255
    var v2 int = -1
    v3 := make(map[int]uint32)
    var v4 int8 = 105
    v5 := []string{("" + (("foo bar" + ("a" + ("a" + "a"))) + "")), "a", "", ""}
    if len(v5) > 0 {
        v5[0] = ((("\n\t" + "") + "a") + "a")
    }
    delete(v3, (v2 & 32768))
    delete(v3, (v2 / ((v2) | 1)))
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 float32 = (3.14159 + 0.0)
            var v23 uint16 = 0
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    v24 := GenS1{F0: false, F1: 4294967295}
    switch (v2 + v2) {
    case 9223372036854775807:
        v4 += int8(v1)
    case 128:
        v4 += -128
    case 4294967295:
        fmt.Println(v4)
    default:
        v1 = 2
    }
    v2 = 83
    var v25 float32 = 0.1
    v5 = append(v5, ("a" + ("a" + "hello")))
    var v26 uint16 = (v1 ^ (v1 / (((v1 % ((v1) | 1))) | 1)))
    if false {
        var v27 uint64 = 0
        v28 := GenS1{F0: true, F1: 893}
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: true, F1: 42}
    if (v1 < v26) {
        fmt.Println(v4, v26, v1)
        v26 = (v1 ^ 1)
        var v30 int8 = v4
        _ = v30
    } else {
        for fi31 := 0; fi31 < 7; fi31++ {
            _ = len(v5)
            _ = fi31
        }
    }
    v2 = 4294967295
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
