package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 bool; F1 uint8; F2 int64 }
func genFn0(p0 bool, p1 string, p2 int16) uint32 {
    switch "hello" {
    case "\n\t":
        var v4 uint64 = 4294967295
        _ = v4
    case "":
        fmt.Println(p1, p0)
    }
    fmt.Println(p2, p2, p0)
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 1}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 65535
}

func main() {
    var v1 bool = (1 > 128)
    var v2 int8 = 1
    var v3 uint = 2
    if true {
        if v1 {
            var v8 uint = v3
            v1 = (!(42 == 221))
            var v9 int8 = (v2 * v2)
            _ = v8
            _ = v9
        } else {
            v2 += 1
            v1 = v1
            switch "foo bar" {
            case "\n\t":
                v10 := GenS1{F0: (v3 ^ v3)}
                _ = v10
            case "hello":
                fmt.Println(v2, v2)
            }
        }
        v11 := []int{}
        _ = v11
    } else {
        if (!true) {
            v12 := []int{166, 128, 629}
            var v13 string = "hello"
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint32 = 4294967295
                var v16 uint64 = 18446744073709551615
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v12
            _ = v13
        } else {
            v17 := []int{1, 2147483647}
            v1 = ((!v1) || v1)
            _ = v17
        }
    }
    switch "\n\t" {
    case "hello":
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(v1)
            _ = fi18
        }
    case "foo bar":
        v19 := []int{695, 2147483647, 32767, 2147483647}
        _ = v19
    case "a":
        v3 += (^(v3))
    }
    v20 := GenS2{F0: v3}
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            switch (^(fi22)) {
            case 128:
                v2 -= 109
            case 2147483647:
                var v23 int64 = 9223372036854775807
                _ = v23
            default:
                if v1 {
                    switch "" {
                    case "hello":
                        switch fi21 {
                        case 127:
                            switch "\n\t" {
                            case "":
                                v24 := GenS1{F0: (^(v3))}
                                _ = v24
                            case "\n\t":
                                var v25 uint16 = 42
                                _ = v25
                            case "foo bar":
                                var v26 int64 = 500
                                _ = v26
                            default:
                                var v27 int8 = (-(v2))
                                v28 := GenS1{F0: v3}
                                _ = v27
                                _ = v28
                            }
                        case 32767:
                            v1 = v1
                        case 9223372036854775807:
                            var v29 int8 = 100
                            _ = v29
                        default:
                            switch (fi22 & (-(fi21))) {
                            case 2147483647:
                                v2 -= (v2 % ((v2) | 1))
                            case -476:
                                v30 := GenS2{F0: v3}
                                _ = v30
                            }
                        }
                    case "a":
                        fmt.Println(v1, v2, v3)
                    case "\n\t":
                        var v31 int64 = -9223372036854775808
                        _ = v31
                    default:
                        if (!(!true)) {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v3, v1)
                            case "foo bar":
                                v32 := []int{}
                                _ = v32
                            default:
                                var v33 uint64 = 0
                                _ = v33
                            }
                            if (42 == 42) {
                                switch "\n\t" {
                                case "\n\t":
                                    if (551 != 454) {
                                        var v34 bool = v1
                                        var v35 int32 = 42
                                        _ = v34
                                        _ = v35
                                    }
                                case "":
                                    v1 = ((v3 + v3) >= 0)
                                }
                            }
                        }
                    }
                } else {
                    v2 = (v2 * v2)
                    v36 := []int{(fi22 / (((-(fi22))) | 1)), fi21, (fi21 / ((1) | 1))}
                    v37 := GenS3{F0: 1, F1: (-310 <= 1)}
                    _ = v36
                    _ = v37
                }
            }
            _ = fi22
        }
        fmt.Println(v3, v1)
        var v38 string = ""
        _ = fi21
        _ = v38
    }
    v39 := []int{}
    if (1e10 <= (float64((-2.25 - 0.1)) * 0.1)) {
        if v1 {
            for fi40 := 0; fi40 < 8; fi40++ {
                switch (fi40 % (((fi40 | (fi40 * (^(fi40))))) | 1)) {
                case 9223372036854775807:
                    v41 := GenS4{F0: ((128 != 128) && v1), F1: 58, F2: -9223372036854775808}
                    _ = v41
                case -2147483648:
                    if (0 < 4294967295) {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            v1 = v1
                            v2 = 100
                            if len(v39) > 0 {
                                v39[0] = 128
                            }
                            _ = fi42
                        }
                        var v43 string = "hello"
                        _ = v43
                    }
                case 25:
                    var v44 bool = true
                    _ = v44
                }
                _ = fi40
            }
        } else {
            v45 := GenS2{F0: v3}
            _ = v45
        }
        var v46 int = 535
        switch (v46 + (v46 + (^(v46)))) {
        case 4294967295:
            var v47 int64 = -9223372036854775808
            _ = v47
        case 611:
            var v48 uint64 = 18446744073709551615
            _ = v48
        case 9223372036854775807:
            switch (v46 - (v46 * (v46 / ((v46) | 1)))) {
            case 1:
                var v49 uint8 = 255
                _ = v49
            }
        default:
            v3 += v3
        }
        _ = v46
    } else {
        for fi50 := 0; fi50 < 5; fi50++ {
            switch (fi50 - 4294967295) {
            case 127:
                _ = len(v39)
            case 2147483647:
                v51 := GenS1{F0: 4294967295}
                _ = v51
            case -55:
                v52 := GenS3{F0: 9223372036854775807, F1: (85 >= 65)}
                _ = v52
            }
            _ = fi50
        }
    }
    _ = len(v39)
    fmt.Println(v2, v1, v2)
    var v53 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v53
}
