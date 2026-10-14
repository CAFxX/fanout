package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 int; F1 uint16; F2 string }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0(p0 int32, p1 string) int8 {
    var v5 string = p1
    var v6 float32 = 0.0
    v5 = (v5 + "\n\t")
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -15
}
func genFn1(p0 int32, p1 int32) uint16 {
    p1 = p1
    p0 -= (p0 >> ((4294967295) & 63))
    if ("a" > "foo bar") {
        v7 := GenS1{F0: 65535, F1: 187, F2: 0}
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{-1}
            var v10 uint = 4294967295
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        switch -553 {
        case 36:
            p0 -= p0
        case 4294967295:
            var v11 float32 = float32(((0.0 / -2.25) + ((1e10 * 1.5) / ((1.5 / ((1e10 * 1.5) - 1e10)) + -2.25))))
            _ = v11
        }
        switch "hello" {
        case "a":
            p1 = p1
        case "foo bar":
            switch "a" {
            case "hello":
                p0 += (p1 & p0)
            case "foo bar":
                v12 := GenS2{F0: -2147483648, F1: 32767, F2: "\n\t"}
                _ = v12
            default:
                var v13 int32 = (p0 / ((p1) | 1))
                _ = v13
            }
        }
        var v14 bool = false
        _ = v14
    }
    v15 := GenS1{F0: 255, F1: 1, F2: 32768}
    var v16 uint16 = 584
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return v16
}

func main() {
    var v1 float64 = 1e10
    var v2 int16 = -1
    v3 := []int{9223372036854775807}
    var v4 int = 0
    v17 := GenS3{F0: 376, F1: -2147483648, F2: 42}
    for fi18 := 0; fi18 < 19; fi18++ {
        switch v4 {
        case 0:
            switch fi18 {
            case -9223372036854775808:
                v19 := GenS2{F0: fi18, F1: 890, F2: "foo bar"}
                _ = v19
            case 898:
                if true {
                    var v20 uint = 220
                    var v21 int64 = -9223372036854775808
                    var v22 uint = v20
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    v4 = (v4 >> ((18446744073709551615) & 63))
                    v2 = v2
                }
            }
        case 4294967295:
            v2 -= v2
        }
        _ = fi18
    }
    var v23 uint32 = 42
    if (true || (3.14159 > 1.5)) {
        fmt.Println(v4, v1, v23)
        v2 = v2
        var v24 int = 152
        _ = v24
    }
    v4 = (v4 << ((65535) & 63))
    v1 += v1
    var v25 uint = 0
    var v26 uint64 = 0
    if true {
        var v27 uint = 42
        switch 2147483647 {
        case 9223372036854775807:
            v28 := GenS3{F0: 354, F1: -1, F2: 119}
            _ = v28
        case -900:
            var v29 uint16 = 32767
            _ = v29
        case 128:
            if (65535 >= 404) {
                if (0 > v2) {
                    if (!true) {
                        if (v2 == -32768) {
                            v30 := GenS4{F0: genFn0(42, "hello"), F1: 41, F2: v26}
                            v1 -= float64(-2.25)
                            fmt.Println(v1, v23)
                            _ = v30
                        }
                    }
                    v27 += 182
                    if (!(255 < 42)) {
                        switch "foo bar" {
                        case "foo bar":
                            v25 += v25
                        default:
                            fmt.Println(v23)
                        }
                        var v31 uint32 = 4294967295
                        _ = v31
                    }
                }
                fmt.Println(v26, v4, v4)
                var v32 int32 = 792
                _ = v32
            } else {
                switch "foo bar" {
                case "\n\t":
                    for fi33 := 0; fi33 < 6; fi33++ {
                        fmt.Println(v23, v2)
                        fmt.Println(v23, v23, v2)
                        _ = fi33
                    }
                case "a":
                    switch -2147483648 {
                    case 9223372036854775807:
                        var v34 uint64 = v26
                        _ = v34
                    case 2147483647:
                        v35 := GenS4{F0: 0, F1: 56, F2: v26}
                        _ = v35
                    }
                case "":
                    v36 := GenS3{F0: 42, F1: 2147483647, F2: 0}
                    _ = v36
                }
                var v37 int32 = 122
                v4 = (v4 ^ (-(v4)))
                _ = v37
            }
        default:
            var v38 int = (^(v4))
            _ = v38
        }
        v1 = 1e10
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
