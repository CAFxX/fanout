package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 int32; F2 uint }
func genFn0(p0 bool, p1 uint64) int64 {
    p1 += p1
    var v5 uint64 = (p1 & p1)
    _ = p0
    _ = p1
    _ = v5
    return 246
}
func genFn1(p0 string) int16 {
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p0)
    case "foo bar":
        var v6 uint32 = 2
        _ = v6
    }
    switch 128 {
    case 128:
        fmt.Println(p0)
    default:
        p0 = p0
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint = 791
    v2 := []int{4294967295, 127, 128, 4294967295}
    var v3 int8 = 100
    var v4 float64 = (0.1 / (1.5 - 0.0))
    var v7 uint32 = 405
    v8 := GenS1{F0: "a", F1: -2147483648}
    var v9 uint32 = v7
    v10 := GenS1{F0: "\n\t", F1: -366}
    if false {
        v4 = (1.5 + (v4 * (1e10 + 3.14159)))
        if (!("" >= "a")) {
            v11 := GenS2{F0: 18446744073709551615, F1: 781, F2: -424}
            var v12 int16 = 1
            _ = v11
            _ = v12
        }
        v3 -= int8(v1)
    }
    v9 = (v7 | v7)
    if true {
        switch ((("a" + "foo bar") + "a") + "") {
        case "\n\t":
            _ = len(v2)
        case "":
            var v13 int32 = 604
            _ = v13
        case "a":
            var v14 bool = (((false && true) || true) || (42 == 2))
            _ = v14
        default:
            v15 := GenS3{F0: 42, F1: 2147483647, F2: (^(v9))}
            _ = v15
        }
        v16 := GenS1{F0: ("hello" + ""), F1: 32767}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            switch 977 {
            case 42:
                switch ("" + "hello") {
                case "\n\t":
                    v18 := GenS3{F0: 0, F1: -160, F2: v9}
                    _ = v18
                case "foo bar":
                    v4 -= (v4 + v4)
                }
            case 1:
                v9 = v9
            case 4294967295:
                v1 += (v1 / ((1) | 1))
            }
            _ = fi17
        }
        var v19 uint = v1
        if len(v2) > 0 {
            v2[0] = 0
        }
        _ = v19
    }
    if (!false) {
        switch "hello" {
        case "":
            v3 = int8(v7)
        }
        switch ((("\n\t" + "\n\t") + "a") + "\n\t") {
        case "\n\t":
            v20 := GenS4{F0: (v1 ^ (v1 * (v1 * (v1 | (^(v1))))))}
            _ = v20
        case "hello":
            v4 = float64((3.14159 - 1.5))
        case "foo bar":
            fmt.Println(v9, v3)
        default:
            if true {
                fmt.Println(v7, v7)
                var v21 bool = (125 > 0)
                v21 = v21
                _ = v21
            }
        }
        fmt.Println(v9, v4, v7)
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        switch (fi22 % ((-1) | 1)) {
        case 127:
            v23 := GenS5{F0: v1, F1: -388, F2: (v1 % ((v1) | 1))}
            _ = v23
        case 128:
            v24 := GenS4{F0: 18446744073709551615}
            _ = v24
        }
        var v25 uint16 = 354
        _ = fi22
        _ = v25
    }
    v4 = ((1.5 / 1.0) + v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
