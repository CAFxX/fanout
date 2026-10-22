package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 bool; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint) uint32 {
    for fi4 := 0; fi4 < 15; fi4++ {
        switch -473 {
        case 127:
            for fi5 := 0; fi5 < 14; fi5++ {
                var v6 int64 = 9223372036854775807
                _ = fi5
                _ = v6
            }
        }
        v7 := GenS1{F0: 140, F1: 42, F2: 1}
        v8 := []int{371}
        _ = fi4
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: false}
    _ = p0
    _ = v9
    return 0
}
func genFn1() string {
    var v10 int = 32768
    var v11 int64 = -9223372036854775808
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := []int{4294967295, (-(v10))}
        v14 := GenS2{F0: (!(0 <= 1))}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 4294967295, F1: 42, F2: 69}
    fmt.Println(v10, v11)
    _ = v10
    _ = v11
    _ = v15
    return "hello"
}

func main() {
    var v1 uint = 92
    v2 := make(map[string]bool)
    var v3 uint8 = 115
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            var v17 uint8 = v3
            var v18 int8 = -128
            v19 := GenS3{F0: 255, F1: true, F2: (fi16 >> ((uint(v17)) & 63))}
            _ = v17
            _ = v18
            _ = v19
        } else {
            v20 := GenS4{F0: -128}
            _ = v20
        }
        _ = fi16
    }
    v3 = v3
    v21 := []int{-2147483648}
    var v22 float64 = (((1.5 * float64(1e10)) / 3.14159) - (float64(((1.5 + 0.0) + 1.5)) + 0.0))
    v23 := v2[""]
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int64 = -9223372036854775808
        switch 1 {
        case 32767:
            var v26 int16 = 32767
            _ = v26
        case -24:
            var v27 uint64 = 42
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    switch (("a" + "") + "\n\t") {
    case "hello":
        var v28 string = "foo bar"
        _ = v28
    case "a":
        _ = len(v21)
    case "\n\t":
        switch 127 {
        case -128:
            fmt.Println(v22, v23)
        case 128:
            v2["\n\t"] = v23
        case 127:
            fmt.Println(v3, v3)
        }
    default:
        v29 := GenS1{F0: 4294967295, F1: 42, F2: 70}
        _ = v29
    }
    v1 = (^(v1))
    if (0 == 1) {
        v30 := GenS5{F0: -108}
        var v31 bool = (-2.25 < (3.14159 / 1.5))
        fmt.Println(v22, v23, v23)
        _ = v30
        _ = v31
    }
    v32 := v2["a"]
    var v33 uint16 = 712
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
}
