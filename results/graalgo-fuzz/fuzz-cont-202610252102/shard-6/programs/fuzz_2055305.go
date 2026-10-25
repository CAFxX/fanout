package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string; F1 uint8; F2 bool }
func genFn0() uint8 {
    v7 := GenS1{F0: 100}
    for fi8 := 0; fi8 < 1; fi8++ {
        switch "hello" {
        case "foo bar":
            if false {
                var v9 int = fi8
                v10 := GenS1{F0: -25}
                v11 := GenS1{F0: 1}
                _ = v10
                _ = v11
                _ = v9
            }
        case "":
            v12 := []int{}
            _ = v12
        }
        var v13 uint32 = 4294967295
        _ = fi8
        _ = v13
    }
    var v14 int = -9223372036854775808
    fmt.Println(v14)
    _ = v14
    _ = v7
    return 255
}
func genFn1(p0 int8, p1 uint32, p2 uint16) uint8 {
    p1 += (^(p1))
    v15 := GenS2{F0: 42, F1: 0, F2: -1}
    p1 += (^(p1))
    p2 += p2
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            var v17 float64 = (0.1 + (float64(float32(1.5)) * 0.0))
            switch "\n\t" {
            case "\n\t":
                v18 := GenS2{F0: 0, F1: -856, F2: 1}
                _ = v18
            case "hello":
                switch fi16 {
                case 32767:
                    p2 -= p2
                default:
                    var v19 float32 = (3.14159 / (0.1 - ((((((((0.1 - float32(v17)) + 3.14159) * 0.1) + 1.5) + 1e10) * 0.0) + 1.5) + -2.25)))
                    _ = v19
                }
            case "a":
                p0 = (p0 / ((p0) | 1))
            default:
                for fi20 := 0; fi20 < 18; fi20++ {
                    fmt.Println(p2, fi20, v17)
                    v21 := []int{}
                    v22 := GenS3{F0: 42, F1: (fi16 & (fi16 | -1))}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            fmt.Println(p2)
            _ = v17
        }
        var v23 uint = 471
        _ = fi16
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 1
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 bool = (2 != 0)
    var v5 int8 = v1
    var v6 int32 = 1
    fmt.Println(v5, v4, v4)
    if false {
        var v24 uint8 = 42
        _ = v24
    }
    var v25 uint8 = 127
    v26 := GenS4{F0: -2147483648}
    var v27 uint = 18446744073709551615
    var v28 uint32 = 127
    switch 4294967295 {
    case 0:
        v29 := GenS1{F0: 89}
        _ = v29
    default:
        v4 = v4
    }
    fmt.Println(v1, v4)
    var v30 int32 = v6
    switch ("hello" + "") {
    case "":
        v31 := GenS5{F0: 739, F1: (v28 - (v28 + v28)), F2: 191}
        _ = v31
    case "foo bar":
        switch -9223372036854775808 {
        case -594:
            v32 := []int{127, -2147483648, 0, 32768}
            _ = v32
        case -391:
            if (-1 > 9223372036854775807) {
                switch 4294967295 {
                case 0:
                    fmt.Println(v30)
                case 128:
                    if (!v4) {
                        v28 -= uint32(v25)
                        v33 := []int{-9223372036854775808, 46, 1}
                        _ = v33
                    }
                case 415:
                    v34 := GenS2{F0: -549, F1: -537, F2: 922}
                    _ = v34
                }
                v6 += -1
            }
        case -511:
            v35 := []int{4294967295, -1, 32768}
            _ = v35
        }
    case "\n\t":
        var v36 int32 = v6
        _ = v36
    default:
        v6 += v30
    }
    fmt.Println(v5, v27)
    v37 := GenS6{F0: v27}
    switch "\n\t" {
    case "\n\t":
        v4 = v4
    case "hello":
        if (!true) {
            var v38 bool = ((!v4) && v4)
            delete(v3, "\n\t")
            _ = v38
        }
    case "a":
        switch -1 {
        case 32767:
            fmt.Println(v5, v5)
        case -9223372036854775808:
            switch (("hello" + "foo bar") + (("a" + ("foo bar" + ("\n\t" + ("\n\t" + "\n\t")))) + ("hello" + "foo bar"))) {
            case "a":
                v39 := []int{787}
                _ = v39
            case "":
                v40 := GenS2{F0: 89, F1: 2147483647, F2: 42}
                _ = v40
            default:
                fmt.Println(v6)
            }
        }
    default:
        var v41 uint64 = uint64(v27)
        _ = v41
    }
    if (v6 >= v30) {
        v42 := []int{}
        v43 := GenS7{F0: "hello", F1: (^(v25)), F2: v4}
        _ = v42
        _ = v43
    }
    v5 = v1
    switch "hello" {
    case "":
        v44 := []int{2147483647, int(v28), -1, 4294967295}
        _ = v44
    default:
        switch "\n\t" {
        case "":
            if ((1000 < int16(v28)) && true) {
                v28 = v28
                v5 -= (v1 ^ v5)
            }
        case "a":
            for fi45 := 0; fi45 < 3; fi45++ {
                fmt.Println(v30)
                fmt.Println(v6, v27, v5)
                var v46 int8 = -37
                _ = fi45
                _ = v46
            }
        case "hello":
            v4 = (!false)
        default:
            var v47 float64 = (-2.25 + (-2.25 / (-2.25 * 0.1)))
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
