package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint }

func main() {
    v1 := []string{"\n\t", "hello", "\n\t"}
    var v2 uint8 = 183
    v3 := []string{(("" + ("foo bar" + "foo bar")) + ""), "foo bar", "a"}
    v4 := []int{-557, 4294967295, -244, 9223372036854775807}
    var v5 float64 = (float64(1e10) - 1.5)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint = 42
        var v8 int32 = 42
        v9 := GenS1{F0: 42, F1: v8, F2: -1}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    switch ("a" + "hello") {
    case "hello":
        fmt.Println(v5, v2)
    default:
        var v10 uint = 42
        _ = v10
    }
    _ = len(v4)
    v2 -= (v2 ^ (v2 + v2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += 0
    if true {
        v5 += float64(0.0)
        if true {
            switch "hello" {
            case "a":
                v11 := GenS2{F0: 0, F1: -405}
                _ = v11
            case "\n\t":
                v12 := GenS3{F0: 9223372036854775807, F1: 0, F2: 1}
                _ = v12
            default:
                switch (("hello" + "") + "a") {
                case "\n\t":
                    switch "foo bar" {
                    case "hello":
                        if false {
                            if (("" + "hello") > ("" + ("foo bar" + "foo bar"))) {
                                _ = len(v1)
                            }
                        }
                    case "a":
                        switch "foo bar" {
                        case "\n\t":
                            v13 := GenS2{F0: -259, F1: -1}
                            _ = v13
                        case "hello":
                            v14 := GenS4{F0: 2, F1: 42}
                            _ = v14
                        }
                    case "foo bar":
                        var v15 bool = false
                        _ = v15
                    }
                case "a":
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v5 = v5
                        _ = fi16
                    }
                case "hello":
                    var v17 float64 = v5
                    _ = v17
                }
            }
            v18 := GenS5{F0: 2}
            _ = v18
        } else {
            v5 = -2.25
            v5 += 3.14159
            var v19 float32 = 1e10
            _ = v19
        }
        fmt.Println(v5, v5)
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v5 -= v5
        _ = fi20
    }
    if ((((1.5 / 3.14159) / (1.5 / -2.25)) / 3.14159) >= ((3.14159 / 0.1) / 3.14159)) {
        v5 -= v5
    } else {
        v21 := GenS3{F0: 2147483647, F1: 402, F2: 2}
        var v22 int16 = -1
        _ = v21
        _ = v22
    }
    switch "" {
    case "a":
        v2 = v2
    default:
        v23 := GenS3{F0: 837, F1: 407, F2: 622}
        _ = v23
    }
    v24 := GenS4{F0: 816, F1: 128}
    var v25 uint = uint(v2)
    v5 += (1e10 * ((v5 * (1.5 / v5)) * v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
