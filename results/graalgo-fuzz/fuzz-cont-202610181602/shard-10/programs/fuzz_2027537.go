package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int8; F1 uint8 }
func genFn0(p0 int8, p1 uint, p2 uint32) uint64 {
    switch 127 {
    case -55:
        v5 := []int{32768, 300, -9223372036854775808}
        _ = v5
    case -714:
        p0 -= p0
    }
    p1 += 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    return 261
}
func genFn1(p0 uint16, p1 bool) uint64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 float32 = ((1.5 - 1e10) * 0.0)
        _ = fi6
        _ = v7
    }
    v8 := []int{128, 42, 127}
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint32 = 913
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 194
    v12 := GenS1{F0: 42, F1: 42}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v8
    return uint64(p0)
}

func main() {
    var v1 int16 = 0
    v2 := []int{}
    var v3 int8 = 1
    var v4 float32 = float32(1.5)
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: 4294967295, F1: fi13}
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v16 int8 = 0
        v3 -= v16
        v16 += v16
        _ = v16
    } else {
        v2 = append(v2, 4294967295)
        v4 = v4
        if true {
            v17 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
            _ = v17
        } else {
            var v18 uint64 = 1
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        fmt.Println(v4, v15, v15)
        switch fi19 {
        case -9223372036854775808:
            v4 = v4
        case 128:
            switch ("foo bar" + "foo bar") {
            case "":
                var v20 float64 = 3.14159
                _ = v20
            case "a":
                v21 := GenS1{F0: 4294967295, F1: fi19}
                _ = v21
            case "foo bar":
                var v22 float64 = float64(float32(-2.25))
                _ = v22
            default:
                v23 := GenS2{F0: v3, F1: 2}
                _ = v23
            }
        }
        _ = fi19
    }
    v3 += (v3 - (-(v3)))
    v2 = append(v2, 1)
    fmt.Println(v1)
    v24 := GenS1{F0: 0, F1: -9223372036854775808}
    var v25 uint64 = 4294967295
    v26 := GenS1{F0: (v25 % ((genFn0(-34, 566, 114)) | 1)), F1: 2147483647}
    switch "hello" {
    case "foo bar":
        switch "foo bar" {
        case "a":
            v4 = (0.0 * float32(0.0))
        case "\n\t":
            var v27 uint8 = 79
            _ = v27
        default:
            var v28 string = "hello"
            _ = v28
        }
    case "a":
        if ((!true) && true) {
            for fi29 := 0; fi29 < 10; fi29++ {
                fmt.Println(v1, v25, v4)
                _ = len(v2)
                v30 := GenS2{F0: (v3 % (((-(v3))) | 1)), F1: 0}
                _ = fi29
                _ = v30
            }
            var v31 int64 = -1
            _ = v31
        }
    case "":
        var v32 uint = 82
        _ = v32
    }
    var v33 float64 = (-2.25 * float64(3.14159))
    v25 += (v25 % ((uint64(v15)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
}
