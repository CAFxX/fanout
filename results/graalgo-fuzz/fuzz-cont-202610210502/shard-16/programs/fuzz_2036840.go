package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int32 }
func genFn0(p0 uint64, p1 int16) int {
    p1 = p1
    fmt.Println(p1)
    if (false && true) {
        p1 = (p1 % ((p1) | 1))
    }
    if (-429 < -9223372036854775808) {
        v4 := GenS1{F0: 128, F1: p1}
        fmt.Println(p1)
        var v5 uint = 0
        _ = v4
        _ = v5
    }
    _ = p0
    _ = p1
    return -301
}
func genFn1(p0 int, p1 uint16, p2 uint32) bool {
    var v6 int64 = 1
    switch genFn0(18446744073709551615, 1000) {
    case 220:
        v7 := GenS1{F0: 24, F1: 216}
        _ = v7
    default:
        switch p0 {
        case 4294967295:
            v6 = (v6 / ((0) | 1))
        case -979:
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 uint16 = p1
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 int32 = -2147483648
                    _ = fi10
                    _ = v11
                }
                v6 += (v6 / ((650) | 1))
                _ = fi8
                _ = v9
            }
        case 0:
            p2 = p2
        default:
            fmt.Println(p0, p2)
        }
    }
    v12 := GenS1{F0: 53, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return (p2 == p2)
}

func main() {
    v1 := []string{"", ("foo bar" + (("\n\t" + "\n\t") + "hello")), ("" + "hello"), "foo bar"}
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v13 := GenS1{F0: 1, F1: -40}
    var v14 int64 = 1
    v15 := GenS1{F0: 60, F1: 449}
    fmt.Println(v14)
    switch "\n\t" {
    case "foo bar":
        switch ("\n\t" + (("hello" + ("" + "hello")) + "a")) {
        case "foo bar":
            v16 := GenS2{F0: 0, F1: (("hello" + "a") + (("a" + "foo bar") + "hello"))}
            _ = v16
        case "hello":
            var v17 float32 = (3.14159 + 0.1)
            _ = v17
        default:
            switch genFn0(1, 100) {
            case 42:
                for fi18 := 0; fi18 < 19; fi18++ {
                    var v19 int32 = -1
                    _ = fi18
                    _ = v19
                }
            }
        }
    case "a":
        if ((((("hello" + "a") + "") + "hello") <= "") && (true || (42 > 1))) {
            v14 = (-(v14))
            switch ((("\n\t" + (("" + ("foo bar" + ("" + "foo bar"))) + "a")) + "\n\t") + "\n\t") {
            case "\n\t":
                for fi20 := 0; fi20 < 13; fi20++ {
                    if ((-(fi20)) > genFn0(4294967295, 628)) {
                        v21 := GenS3{F0: 1000, F1: -415, F2: 1}
                        var v22 int64 = v14
                        _ = v21
                        _ = v22
                    }
                    v14 = -1
                    var v23 uint = 2
                    _ = fi20
                    _ = v23
                }
            case "":
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 uint64 = 4294967295
                    var v26 int16 = -32768
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case "foo bar":
                var v27 float64 = float64(float32((-2.25 - 0.0)))
                _ = v27
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v14, fi28)
                }
                _ = fi28
            }
        } else {
            if (414 == 156) {
                v14 = (v14 * v14)
                v29 := GenS4{F0: (false && false), F1: 708}
                if true {
                    switch -163 {
                    case -516:
                        v30 := GenS5{F0: ("foo bar" + ("foo bar" + ("foo bar" + ""))), F1: 2147483648, F2: 2147483648}
                        _ = v30
                    case 127:
                        if false {
                            v14 = v14
                            for fi31 := 0; fi31 < 10; fi31++ {
                                switch (fi31 | (fi31 | (-(fi31)))) {
                                case 355:
                                    var v32 int64 = (-(v14))
                                    _ = v32
                                }
                                v14 -= v14
                                fmt.Println(v14, fi31)
                                _ = fi31
                            }
                        }
                    case -2147483648:
                        if false {
                            v14 += (v14 + -1)
                            v14 += 9223372036854775807
                        }
                    }
                    for fi33 := 0; fi33 < 3; fi33++ {
                        var v34 int64 = (v14 * v14)
                        _ = fi33
                        _ = v34
                    }
                }
                _ = v29
            }
            v14 = 9223372036854775807
        }
    }
    var v35 int32 = -1
    fmt.Println(v14, v14)
    var v36 int32 = (v35 % (((v35 << ((255) & 63))) | 1))
    v37 := GenS4{F0: false, F1: 0}
    _ = len(v1)
    if true {
        v35 = v36
    } else {
        fmt.Println(v36, v35)
    }
    v36 += (v36 * (-(v35)))
    v38 := GenS6{F0: 42, F1: v14}
    _ = len(v1)
    _ = len(v1)
    switch 4294967295 {
    case -647:
        switch (("hello" + "a") + "a") {
        case "":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        case "\n\t":
            if true {
                fmt.Println(v14)
                if ((v14 + 1) < (v14 >> ((42) & 63))) {
                    fmt.Println(v35, v14)
                    switch "a" {
                    case "\n\t":
                        if (true && (v14 != v14)) {
                            v35 = (^(v36))
                        }
                    case "hello":
                        if (1e10 <= 3.14159) {
                            v39 := GenS6{F0: 0, F1: (v14 >> ((0) & 63))}
                            var v40 bool = false
                            v35 -= (v36 * v35)
                            _ = v39
                            _ = v40
                        } else {
                            if true {
                                switch "hello" {
                                case "hello":
                                    var v41 string = "hello"
                                    _ = v41
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                default:
                                    for fi42 := 0; fi42 < 11; fi42++ {
                                        if ((("hello" + ("a" + "\n\t")) + ("" + "a")) >= ("hello" + ("\n\t" + ("" + "hello")))) {
                                            fmt.Println(v36, v35)
                                            v3[fi42] = (!(42 >= 2147483647))
                                        }
                                        delete(v3, fi42)
                                        _ = fi42
                                    }
                                }
                                v43 := GenS7{F0: 1, F1: 18446744073709551615}
                                _ = v43
                            } else {
                                var v44 int64 = 9223372036854775807
                                v44 -= (-(v44))
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v1)
                                case "\n\t":
                                    if ("a" != "\n\t") {
                                        _ = len(v2)
                                        if false {
                                            v45 := GenS8{F0: v35}
                                            v14 = (-(v44))
                                            for fi46 := 0; fi46 < 9; fi46++ {
                                                switch genFn0(0, -740) {
                                                case -9223372036854775808:
                                                    switch (fi46 * fi46) {
                                                    case 841:
                                                        v47 := GenS6{F0: 1, F1: v14}
                                                        _ = v47
                                                    case 0:
                                                        v14 = (v44 - (-(v44)))
                                                    case -1:
                                                        fmt.Println(v35, fi46, v35)
                                                    default:
                                                        v44 += (v44 >> ((255) & 63))
                                                    }
                                                }
                                                var v48 float64 = 1e10
                                                var v49 int16 = 0
                                                _ = fi46
                                                _ = v48
                                                _ = v49
                                            }
                                            _ = v45
                                        }
                                    }
                                }
                                _ = v44
                            }
                        }
                    default:
                        fmt.Println(v35)
                    }
                    v1 = append(v1, "foo bar")
                } else {
                    var v50 int64 = (-(v14))
                    var v51 uint16 = 281
                    _ = v50
                    _ = v51
                }
                if true {
                    fmt.Println(v14, v35, v36)
                } else {
                    var v52 string = "foo bar"
                    for fi53 := 0; fi53 < 15; fi53++ {
                        if false {
                            v35 += 1
                            var v54 int16 = -1
                            _ = v54
                        }
                        switch "" {
                        case "":
                            v55 := v2[544]
                            _ = v55
                        }
                        _ = fi53
                    }
                    v36 = v36
                    _ = v52
                }
            }
        default:
            for fi56 := 0; fi56 < 8; fi56++ {
                var v57 int = genFn0(4294967295, -466)
                _ = fi56
                _ = v57
            }
        }
    case 9223372036854775807:
        var v58 string = (("a" + ("hello" + "foo bar")) + "\n\t")
        _ = v58
    default:
        for fi59 := 0; fi59 < 11; fi59++ {
            var v60 uint = 4294967295
            fmt.Println(v14)
            _ = fi59
            _ = v60
        }
    }
    for fi61 := 0; fi61 < 4; fi61++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3[(^(fi61))] = (!true)
        v35 = (v35 * (v35 / ((v35) | 1)))
        _ = fi61
    }
    v14 = v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
