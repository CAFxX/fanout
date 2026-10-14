package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int, p1 int, p2 uint16) uint32 {
    var v4 int16 = 42
    var v5 uint16 = p2
    p2 = (p2 ^ (p2 >> ((4294967295) & 63)))
    fmt.Println(v4, p0, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 bool, p2 uint16) uint16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint64 = 977
            v9 := []int{127, fi6, (-(fi6))}
            var v10 int32 = 1
            v11 := GenS1{F0: 1000, F1: (p0 | -398)}
            v12 := GenS2{F0: 127}
            _ = fi7
            _ = v10
            _ = v11
            _ = v12
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    switch 0 {
    case 537:
        var v13 uint64 = 523
        _ = v13
    case 32767:
        var v14 int64 = 1
        _ = v14
    case -918:
        p0 = p0
    default:
        var v15 string = (("hello" + "\n\t") + ("" + ("\n\t" + "")))
        _ = v15
    }
    v16 := []int{942, 400}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return (p2 + p2)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int16 = -96
    v3 = 1
    var v17 int16 = 1
    v18 := GenS1{F0: v17, F1: 9223372036854775807}
    if false {
        fmt.Println(v17, v3)
        v19 := GenS1{F0: v17, F1: 894}
        fmt.Println(v3)
        _ = v19
    }
    v20 := v2[""]
    v3 = 103
    fmt.Println(v3, v3, v20)
    fmt.Println(v17)
    v3 = 604
    fmt.Println(v20, v3, v17)
    var v21 uint32 = 1
    v22 := []int{0, 128}
    if (((^(v17)) > v17) && v20) {
        v17 = v17
        fmt.Println(v21)
        switch "foo bar" {
        case "\n\t":
            v20 = v20
        case "foo bar":
            fmt.Println(v21, v20, v21)
        case "hello":
            if ((v17 % ((v17) | 1)) < v17) {
                if v20 {
                    _ = len(v1)
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    var v23 bool = (!v20)
                    _ = v23
                } else {
                    v24 := GenS2{F0: -128}
                    v21 = genFn0(32767, 4294967295, genFn1(1, (true || (!true)), 2))
                    v25 := GenS2{F0: -60}
                    _ = v24
                    _ = v25
                }
                v20 = false
            }
        }
    }
    fmt.Println(v3, v3)
    if v20 {
        var v26 uint64 = 4294967295
        v21 = v21
        v27 := v2[("" + "a")]
        _ = v26
        _ = v27
    }
    v17 += (v3 | v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
