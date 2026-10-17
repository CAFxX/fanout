package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 uint }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 int16, p1 int16, p2 bool) bool {
    v6 := []int{-689, -1, -1, 236}
    switch "foo bar" {
    case "a":
        var v7 float32 = 0.1
        _ = v7
    case "hello":
        switch (("" + "hello") + "a") {
        case "foo bar":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "hello":
            var v8 uint64 = 18446744073709551615
            _ = v8
        }
    case "\n\t":
        var v9 float64 = float64(float32(0.1))
        _ = v9
    default:
        fmt.Println(p0, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (((1.5 - 0.1) * 1.5) < float64(float32(1e10)))
}
func genFn1() int32 {
    var v10 uint32 = 905
    if (!false) {
        if true {
            if (1 >= -128) {
                v11 := GenS1{F0: 32767}
                v10 -= (v10 & v10)
                _ = v11
            } else {
                v12 := GenS2{F0: -2147483648, F1: 4294967295, F2: 0}
                _ = v12
            }
        }
        v10 -= 2147483647
    } else {
        fmt.Println(v10)
        var v13 uint = 42
        _ = v13
    }
    v14 := GenS2{F0: 42, F1: 961, F2: 65535}
    v15 := []int{}
    var v16 uint = 989
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 605
    var v3 int64 = -9223372036854775808
    var v4 string = ""
    var v5 uint = 2
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 int32 = int32(v3)
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: 0, F1: 66, F2: (v4 + v4)}
    switch "a" {
    case "\n\t":
        v20 := GenS3{F0: -459, F1: -124, F2: "foo bar"}
        _ = v20
    case "hello":
        v1 = v1
    default:
        if false {
            fmt.Println(v2)
            fmt.Println(v2)
        }
    }
    if (false && (0.1 >= (3.14159 + -2.25))) {
        v1 = v1
        v21 := []int{}
        _ = v21
    } else {
        v4 = "hello"
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v1 = 42
        var v23 uint64 = 4294967295
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 int = fi24
            var v26 string = v4
            var v27 uint8 = 128
            v28 := GenS4{F0: fi24, F1: 1}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi22
        _ = v23
    }
    v4 = v4
    v1 = (v1 + (^(v1)))
    v1 = v1
    v29 := []int{}
    v30 := GenS3{F0: -32768, F1: 1, F2: (("" + (v4 + v4)) + "foo bar")}
    switch -1 {
    case 185:
        var v31 float32 = 1e10
        _ = v31
    case 9223372036854775807:
        var v32 int8 = 0
        _ = v32
    case 2147483647:
        var v33 int = 267
        _ = v33
    default:
        fmt.Println(v2, v3)
    }
    v4 = "hello"
    v3 = (-(v3))
    var v34 int32 = -83
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
