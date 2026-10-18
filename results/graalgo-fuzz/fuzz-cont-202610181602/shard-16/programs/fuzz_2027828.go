package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 string, p1 int32) uint {
    p1 -= p1
    if (18446744073709551615 < 707) {
        var v4 uint16 = 2
        v5 := GenS1{F0: 639, F1: (^(v4))}
        _ = v4
        _ = v5
    }
    switch -2147483648 {
    case 4294967295:
        var v6 bool = (-2.25 > 0.1)
        _ = v6
    }
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 367, F1: 2}
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 int32) int16 {
    p1 = (p1 & p1)
    var v8 int32 = 739
    v9 := []int{1, 42}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -146
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]uint32)
    var v3 uint = 286
    v10 := GenS1{F0: 719, F1: 442}
    v1 = v1
    var v11 int = 4294967295
    var v12 uint8 = 161
    fmt.Println(v12, v3)
    var v13 string = (v1 + v1)
    v14 := []int{}
    v13 = v1
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: 2, F1: 42}
        var v17 uint32 = 42
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 42, F1: v13}
    switch (("\n\t" + v13) + v1) {
    case "\n\t":
        v1 = v13
    case "hello":
        v11 -= v11
    case "foo bar":
        switch "foo bar" {
        case "":
            fmt.Println(v1, v1, v11)
        default:
            v19 := GenS3{F0: 9223372036854775807, F1: (("" == (v13 + "")) || (false && true))}
            _ = v19
        }
    }
    if (!(2147483647 != -667)) {
        for fi20 := 0; fi20 < 8; fi20++ {
            v2["foo bar"] = 42
            _ = fi20
        }
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            fmt.Println(fi21, v1)
            _ = fi21
        }
        v3 -= v3
        var v22 string = (((v1 + "foo bar") + ("" + ("" + v1))) + "")
        _ = v22
    }
    fmt.Println(v3, v3, v13)
    var v23 uint32 = 65535
    for fi24 := 0; fi24 < 16; fi24++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        var v25 uint32 = 42
        _ = fi24
        _ = v25
    }
    v23 = v23
    v14 = append(v14, -925)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
