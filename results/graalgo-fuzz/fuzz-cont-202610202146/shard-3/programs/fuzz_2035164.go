package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0() uint8 {
    if false {
        var v7 uint8 = 127
        _ = v7
    }
    var v8 uint64 = 42
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: v8}
            v8 += v8
            _ = fi10
            _ = v11
        }
        v12 := []int{(fi9 / ((-482) | 1)), 127}
        _ = fi9
        _ = v12
    }
    v13 := GenS1{F0: (v8 | v8)}
    _ = v13
    _ = v8
    return 137
}
func genFn1() uint32 {
    if true {
        v14 := GenS1{F0: 18446744073709551615}
        var v15 uint16 = 296
        _ = v14
        _ = v15
    }
    var v16 uint64 = 510
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := []int{(fi17 % ((32767) | 1)), 1, fi17}
        v19 := GenS2{F0: 4294967295}
        var v20 uint8 = 0
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 uint = 18446744073709551615
    v22 := []int{-2147483648, 427}
    _ = v16
    _ = v21
    _ = v22
    return 2
}

func main() {
    var v1 float32 = 1.5
    var v2 uint32 = 4294967295
    var v3 int64 = 9223372036854775807
    var v4 int32 = 0
    var v5 float64 = (1.5 - 3.14159)
    v6 := []string{"a", "hello", ""}
    switch ("hello" + (("" + "hello") + "hello")) {
    case "hello":
        if len(v6) > 0 {
            v6[0] = ("\n\t" + ((("" + "foo bar") + (("hello" + "a") + (("\n\t" + "hello") + "hello"))) + "hello"))
        }
    case "\n\t":
        v2 = v2
    }
    if (true && (0.0 > ((v1 - v1) - -2.25))) {
        fmt.Println(v4, v4, v4)
        var v23 int8 = 42
        v5 -= v5
        _ = v23
    }
    v1 += (v1 + (v1 * (v1 - v1)))
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    fmt.Println(v5, v4, v3)
    v24 := GenS1{F0: 166}
    fmt.Println(v3)
    if false {
        fmt.Println(v1, v5, v1)
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v5, v3)
        case "\n\t":
            if false {
                _ = len(v6)
            } else {
                switch "a" {
                case "a":
                    switch 128 {
                    case -2147483648:
                        v6 = append(v6, "")
                    default:
                        var v25 int = 42
                        _ = v25
                    }
                case "hello":
                    var v26 int8 = int8(v3)
                    _ = v26
                }
                var v27 int32 = v4
                v2 = (v2 ^ (v2 >> ((0) & 63)))
                _ = v27
            }
        case "a":
            v2 = 1
        default:
            v5 -= ((((((float64(v1) * 1e10) - v5) + v5) / 1.5) + v5) / v5)
        }
    } else {
        if len(v6) > 0 {
            v6[0] = ""
        }
        v3 = 1
    }
    v28 := GenS3{F0: 65535, F1: (v4 + (v4 * (-(v4))))}
    fmt.Println(v5)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
