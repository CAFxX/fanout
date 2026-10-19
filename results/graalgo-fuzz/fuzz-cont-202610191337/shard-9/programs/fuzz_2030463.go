package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 uint64, p1 int, p2 uint) uint16 {
    if (!false) {
        var v4 uint64 = p0
        v5 := GenS1{F0: 1, F1: p0, F2: 655}
        _ = v4
        _ = v5
    } else {
        v6 := GenS2{F0: 2147483647, F1: 128}
        p0 -= (p0 | (p0 / (((p0 * (p0 % ((277) | 1)))) | 1)))
        _ = v6
    }
    if (!(!true)) {
        var v7 int64 = 81
        var v8 string = "hello"
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{(p1 << ((65535) & 63)), (p1 * p1)}
        p0 = 0
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        p2 = p2
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 string) string {
    switch 4294967295 {
    case 0:
        v12 := []int{32767, 0, -655}
        _ = v12
    case 32767:
        for fi13 := 0; fi13 < 3; fi13++ {
            switch p0 {
            case "a":
                v14 := GenS1{F0: -1, F1: 908, F2: 65535}
                _ = v14
            case "\n\t":
                p0 = p0
            }
            v15 := []int{fi13, (fi13 << ((473) & 63))}
            v16 := GenS1{F0: 718, F1: 2, F2: 32767}
            _ = fi13
            _ = v15
            _ = v16
        }
    }
    v17 := GenS1{F0: -811, F1: 42, F2: 2}
    switch 2147483647 {
    case 42:
        p0 = ""
    case 113:
        v18 := GenS1{F0: -773, F1: 18446744073709551615, F2: 596}
        _ = v18
    case 9223372036854775807:
        var v19 uint16 = 909
        _ = v19
    default:
        for fi20 := 0; fi20 < 18; fi20++ {
            p0 = p0
            _ = fi20
        }
    }
    _ = p0
    _ = v17
    return (p0 + (p0 + "a"))
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v21 int32 = 0
    fmt.Println(v1, v1)
    if false {
        v22 := GenS2{F0: v1, F1: 333}
        switch 385 {
        case -64:
            var v23 bool = true
            _ = v23
        case 9223372036854775807:
            v1 = (v1 | v1)
        }
        _ = v22
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v2[("foo bar" + "foo bar")] = fi24
        _ = fi24
    }
    if false {
        v21 = -2147483648
        fmt.Println(v21)
        if (1 > 1) {
            var v25 uint8 = 166
            fmt.Println(v25, v25, v25)
            v26 := []int{-340, 2147483647}
            _ = v25
            _ = v26
        }
    }
    var v27 float64 = ((1.5 - 0.1) - 0.1)
    var v28 int16 = 1
    fmt.Println(v1, v27, v28)
    v1 = v1
    var v29 int32 = 0
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 bool = true
        v31 = v31
        v32 := v2["a"]
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v29)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v28, v21)
    default:
        v27 = float64(float32(v27))
    }
    v33 := []int{1, -9223372036854775808, 9223372036854775807, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
}
