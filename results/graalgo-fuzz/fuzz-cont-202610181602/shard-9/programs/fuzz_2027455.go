package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 int64, p1 string) int8 {
    fmt.Println(p1)
    fmt.Println(p1)
    var v7 uint = 18446744073709551615
    fmt.Println(v7, p0, p0)
    switch 881 {
    case 128:
        for fi8 := 0; fi8 < 3; fi8++ {
            v7 = v7
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS1{F0: true, F1: -21}
                _ = fi9
                _ = v10
            }
            p0 = 1
            _ = fi8
        }
    case -9223372036854775808:
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                var v12 uint64 = 479
                _ = v12
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return -128
}

func main() {
    var v1 int16 = 1000
    var v2 uint = 170
    var v3 float64 = 1e10
    v4 := []int{884}
    v5 := []string{}
    var v6 bool = true
    switch ("a" + "\n\t") {
    case "":
        fmt.Println(v3, v1, v3)
    case "hello":
        for fi13 := 0; fi13 < 15; fi13++ {
            v6 = ((((-2.25 - v3) + (v3 - 1.5)) * 0.0) == v3)
            _ = fi13
        }
    case "a":
        if true {
            if len(v5) > 0 {
                v5[0] = "hello"
            }
        }
    default:
        var v14 int16 = (v1 - v1)
        _ = v14
    }
    var v15 int = 32767
    fmt.Println(v6)
    switch ("\n\t" + "hello") {
    case "\n\t":
        v1 -= (-(v1))
    }
    if (499 >= 42) {
        v6 = (v6 || (("\n\t" != "a") && (0 < 42)))
    }
    v16 := GenS2{F0: v15, F1: (v2 >> (((v2 % ((30) | 1))) & 63))}
    if true {
        v3 += 0.1
        v6 = (v15 >= (v15 ^ v15))
        v17 := GenS2{F0: v15, F1: 42}
        _ = v17
    } else {
        var v18 float32 = ((0.0 + -2.25) + ((1.5 * (0.0 / 1.0)) + 1e10))
        if v6 {
            v19 := GenS2{F0: v15, F1: v2}
            _ = v19
        }
        _ = v18
    }
    var v20 int64 = 0
    for fi21 := 0; fi21 < 5; fi21++ {
        v3 -= (0.0 + (float64(0.0) / v3))
        _ = fi21
    }
    if v6 {
        for fi22 := 0; fi22 < 12; fi22++ {
            v2 = (^(v2))
            v23 := GenS3{F0: 1}
            _ = fi22
            _ = v23
        }
        if (!v6) {
            v24 := GenS1{F0: v6, F1: genFn0(v20, ("\n\t" + ("foo bar" + "")))}
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            _ = v24
        }
    }
    fmt.Println(v6)
    switch "foo bar" {
    case "\n\t":
        var v25 int = (v15 / ((v15) | 1))
        _ = v25
    case "":
        fmt.Println(v2, v20)
    }
    v26 := GenS2{F0: v15, F1: 65535}
    v27 := GenS1{F0: v6, F1: 127}
    var v28 int32 = 403
    var v29 uint64 = 2
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 uint8 = 254
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
