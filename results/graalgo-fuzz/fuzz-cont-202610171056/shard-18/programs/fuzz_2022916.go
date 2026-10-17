package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16, p1 string) int64 {
    if false {
        for fi4 := 0; fi4 < 7; fi4++ {
            p0 = p0
            var v5 uint32 = 2147483648
            v6 := GenS1{F0: 42, F1: p0}
            _ = fi4
            _ = v5
            _ = v6
        }
        p1 = ((p1 + p1) + "hello")
    }
    v7 := GenS1{F0: 96, F1: (p0 | p0)}
    switch "" {
    case "a":
        v8 := GenS2{F0: 4294967295, F1: 0}
        _ = v8
    case "":
        p0 -= (p0 ^ (p0 % ((p0) | 1)))
    case "hello":
        if false {
            if (65535 < 189) {
                fmt.Println(p0, p1)
                var v9 float64 = 1.5
                for fi10 := 0; fi10 < 18; fi10++ {
                    v9 -= ((v9 - v9) + ((v9 * float64(1e10)) * (v9 / v9)))
                    var v11 int8 = -1
                    v12 := []int{(fi10 - -1), fi10, fi10}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            }
            v13 := []int{26}
            _ = v13
        }
    default:
        var v14 int16 = 1000
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := GenS2{F0: 10, F1: 42}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 uint16) string {
    var v17 uint = 18446744073709551615
    var v18 bool = true
    _ = p0
    _ = v17
    _ = v18
    return ""
}

func main() {
    var v1 bool = true
    v2 := []string{"foo bar", ("" + "")}
    v3 := make(map[string]bool)
    v19 := GenS2{F0: 2, F1: 255}
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint64 = 2
        var v22 int = fi20
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        fmt.Println(v1, fi23)
        v1 = true
        v3[(((("hello" + (("a" + ("\n\t" + "hello")) + "\n\t")) + "\n\t") + (((("" + "foo bar") + "a") + "\n\t") + ("\n\t" + ""))) + "hello")] = (722 <= 325)
        _ = fi23
    }
    if ((0 <= -2147483648) && (!v1)) {
        var v24 int32 = 42
        fmt.Println(v24)
        v1 = (18446744073709551615 > 42)
        _ = v24
    }
    var v25 int16 = 0
    if (v25 != v25) {
        if (v25 >= v25) {
            v1 = true
        }
        v26 := GenS2{F0: 2147483648, F1: 128}
        _ = v26
    } else {
        fmt.Println(v25)
        var v27 uint32 = 552
        _ = v27
    }
    v25 -= (-(v25))
    var v28 uint8 = 42
    for fi29 := 0; fi29 < 2; fi29++ {
        if (!v1) {
            v28 = (v28 * v28)
            fmt.Println(fi29)
            v25 = v25
        }
        if len(v2) > 0 {
            v2[0] = ("hello" + "\n\t")
        }
        _ = fi29
    }
    var v30 float64 = 0.0
    if (!false) {
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 uint64 = 4294967295
            var v33 uint8 = v28
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v30
}
