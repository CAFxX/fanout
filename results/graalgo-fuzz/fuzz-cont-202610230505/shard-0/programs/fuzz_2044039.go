package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 int; F2 uint8 }
func genFn0(p0 int) string {
    if false {
        p0 = (p0 ^ (p0 / (((p0 << ((4294967295) & 63))) | 1)))
    } else {
        fmt.Println(p0)
    }
    if true {
        p0 -= (p0 >> ((1) & 63))
        var v5 string = "a"
        v6 := GenS1{F0: 4294967295}
        _ = v5
        _ = v6
    } else {
        var v7 float64 = 1.5
        var v8 float32 = (((-2.25 * 1.5) - (0.1 * 1.5)) * 1e10)
        _ = v7
        _ = v8
    }
    _ = p0
    return "hello"
}
func genFn1(p0 uint32, p1 int16) uint32 {
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint64 = 249
        if false {
            fmt.Println(p1, fi9, p1)
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int32 = -2147483648
                v13 := GenS1{F0: 42}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v14 := []int{}
        _ = fi9
        _ = v10
        _ = v14
    }
    var v15 int16 = 978
    p0 = p0
    v16 := []int{306, 1}
    if ((!false) && ((("" + ("" + "hello")) + "hello") < "foo bar")) {
        v16 = append(v16, 1)
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return (p0 | p0)
}

func main() {
    var v1 string = ("" + "")
    v2 := make(map[int]bool)
    var v3 int16 = 42
    v4 := []string{((v1 + v1) + ""), "foo bar"}
    for fi17 := 0; fi17 < 16; fi17++ {
        switch (v1 + v1) {
        case "\n\t":
            switch 42 {
            case 42:
                fmt.Println(v3, fi17, fi17)
            case 376:
                if (genFn1(0, v3) >= genFn1(genFn1(707, (v3 & v3)), v3)) {
                    delete(v2, (-(fi17)))
                    v18 := v2[42]
                    var v19 int64 = 9223372036854775807
                    _ = v18
                    _ = v19
                }
            case 812:
                if false {
                    v1 = "foo bar"
                    v3 = (^(v3))
                }
            }
        default:
            var v20 uint8 = 176
            _ = v20
        }
        v3 = v3
        _ = fi17
    }
    switch 42 {
    case 127:
        v3 = 0
    default:
        v3 = (v3 / (((^(v3))) | 1))
    }
    v1 = v1
    v3 -= v3
    fmt.Println(v1, v1)
    v3 -= -774
    var v21 float32 = 1e10
    var v22 int = -49
    fmt.Println(v1, v1)
    if true {
        var v23 int32 = 282
        _ = v23
    } else {
        var v24 uint32 = 248
        v25 := GenS1{F0: 18446744073709551615}
        _ = v24
        _ = v25
    }
    _ = len(v4)
    var v26 string = "hello"
    v27 := GenS2{F0: 128, F1: 65535}
    if len(v4) > 0 {
        v4[0] = (v26 + v26)
    }
    fmt.Println(v1, v22)
    v28 := GenS3{F0: v1, F1: 128, F2: 42}
    if false {
        var v29 uint8 = 128
        v2[v22] = true
        v1 = v26
        _ = v29
    } else {
        var v30 bool = (2 < 699)
        var v31 string = (v26 + v26)
        _ = v30
        _ = v31
    }
    var v32 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
}
