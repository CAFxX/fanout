package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint }
type GenS6 struct { F0 uint32; F1 string; F2 bool }
type GenS7 struct { F0 int64; F1 int32; F2 uint64 }
type GenS8 struct { F0 bool }
func genFn0() int8 {
    v6 := GenS1{F0: 65535, F1: "foo bar"}
    v7 := []int{9223372036854775807, -1, 0}
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint) uint16 {
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int64 = 9223372036854775807
        for fi10 := 0; fi10 < 10; fi10++ {
            p0 = p0
            fmt.Println(v9, p0, fi8)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{fi11, (fi11 & 600), fi11, fi11}
        var v13 int32 = 2147483647
        fmt.Println(fi11, v13, fi11)
        _ = fi11
        _ = v12
        _ = v13
    }
    p0 = p0
    p0 += 995
    v14 := GenS2{F0: 42, F1: 0, F2: (p0 & p0)}
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 int16 = -839
    v2 := []int{4294967295}
    v3 := []string{"hello", "foo bar", "\n\t"}
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]int)
    switch ("" + "") {
    case "hello":
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
    case "foo bar":
        for fi15 := 0; fi15 < 8; fi15++ {
            v2 = append(v2, (fi15 | -111))
            var v16 uint16 = 1
            v16 += (v16 % (((v16 % (((v16 * (v16 * v16))) | 1))) | 1))
            _ = fi15
            _ = v16
        }
    default:
        v17 := GenS1{F0: 1, F1: ("a" + (("\n\t" + "a") + "\n\t"))}
        _ = v17
    }
    var v18 uint = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -1 {
    case 128:
        var v19 int64 = 42
        _ = v19
    case 32768:
        v1 += v1
    case 807:
        for fi20 := 0; fi20 < 3; fi20++ {
            _ = len(v5)
            var v21 uint64 = 947
            var v22 float32 = ((float32(1e10) + (0.1 / -2.25)) / float32(1e10))
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        switch "hello" {
        case "foo bar":
            var v24 uint = 42
            _ = v24
        case "\n\t":
            fmt.Println(v18, fi23)
        case "a":
            var v25 uint = v18
            _ = v25
        default:
            switch -82 {
            case 1:
                v4["\n\t"] = 2147483648
            case 32768:
                for fi26 := 0; fi26 < 8; fi26++ {
                    fmt.Println(v18)
                    fmt.Println(fi26, fi26)
                    var v27 int64 = -9223372036854775808
                    _ = fi26
                    _ = v27
                }
            }
        }
        _ = fi23
    }
    var v28 uint32 = 42
    v29 := GenS3{F0: 1, F1: -1}
    v18 -= v18
    v30 := GenS3{F0: 181, F1: 46}
    fmt.Println(v18, v18, v1)
    var v31 int32 = 0
    v31 += v31
    var v32 string = "hello"
    for fi33 := 0; fi33 < 13; fi33++ {
        fmt.Println(v31, v28, v18)
        var v34 float64 = (((1e10 / -2.25) + 3.14159) * (1.5 - 0.0))
        _ = fi33
        _ = v34
    }
    if true {
        if false {
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 int8 = 100
                _ = fi35
                _ = v36
            }
            if false {
                if ((true || ((true && true) || false)) || true) {
                    v37 := GenS4{F0: -125, F1: v28, F2: 234}
                    _ = v37
                }
                v31 += (^(v31))
            }
            v38 := GenS5{F0: v31, F1: 2147483648, F2: 255}
            _ = v38
        } else {
            v28 -= v28
            fmt.Println(v18, v28)
            v31 = (v31 % ((0) | 1))
        }
        var v39 int = -1
        var v40 bool = ((true || ((-(v39)) > v39)) || ((v18 ^ 628) > v18))
        _ = v39
        _ = v40
    }
    v41 := GenS6{F0: 2, F1: "\n\t", F2: false}
    v1 = v1
    switch 42 {
    case 128:
        switch "a" {
        case "hello":
            if (804 > 18446744073709551615) {
                switch 9223372036854775807 {
                case 32767:
                    if (false && false) {
                        var v42 float32 = (float32(-2.25) / 1.0)
                        var v43 int8 = 15
                        v32 = ""
                        _ = v42
                        _ = v43
                    }
                case 32768:
                    v44 := GenS2{F0: v1, F1: v31, F2: v18}
                    _ = v44
                case -2147483648:
                    v18 = 255
                default:
                    fmt.Println(v28)
                }
                v45 := GenS4{F0: -20, F1: 42, F2: 165}
                _ = v45
            }
        case "foo bar":
            switch -667 {
            case 543:
                fmt.Println(v32, v28, v1)
            default:
                v46 := GenS6{F0: v28, F1: (v32 + v32), F2: (false && (42 < 18446744073709551615))}
                _ = v46
            }
        default:
            switch ((("hello" + ("hello" + v32)) + ("a" + "\n\t")) + (("foo bar" + ((("a" + v32) + "foo bar") + "hello")) + ("a" + v32))) {
            case "hello":
                var v47 int64 = 42
                _ = v47
            case "\n\t":
                switch -1 {
                case 32768:
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                case 740:
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                default:
                    if false {
                        if (521 > 4294967295) {
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = -1
                                }
                            } else {
                                switch "hello" {
                                case "hello":
                                    v48 := GenS7{F0: 321, F1: (v31 - (v31 / ((v31) | 1))), F2: uint64(v18)}
                                    _ = v48
                                case "\n\t":
                                    var v49 float32 = -2.25
                                    _ = v49
                                case "":
                                    v1 = (v1 & int16(v28))
                                }
                            }
                        } else {
                            if (!(214 > 1)) {
                                v50 := GenS8{F0: false}
                                v3 = append(v3, (("foo bar" + (v32 + ("hello" + (("foo bar" + "\n\t") + v32)))) + "a"))
                                _ = v50
                            } else {
                                if false {
                                    var v51 string = v32
                                    v32 = "\n\t"
                                    _ = v51
                                }
                                var v52 uint = 196
                                switch "hello" {
                                case "a":
                                    if (42 > 128) {
                                        for fi53 := 0; fi53 < 11; fi53++ {
                                            if ((true && (!((genFn0() == -128) && (!false)))) && (!(false && true))) {
                                                var v54 int8 = -128
                                                v55 := GenS2{F0: v1, F1: (-(v31)), F2: v18}
                                                switch -2147483648 {
                                                case 42:
                                                    switch (fi53 & fi53) {
                                                    case 4294967295:
                                                        fmt.Println(v54, v32, fi53)
                                                    case -1:
                                                        v18 = 255
                                                    }
                                                case -194:
                                                    v56 := GenS5{F0: v31, F1: (v28 >> (((v18 >> ((v52) & 63))) & 63)), F2: (v52 % ((66) | 1))}
                                                    _ = v56
                                                }
                                                _ = v54
                                                _ = v55
                                            }
                                            _ = fi53
                                        }
                                    }
                                case "foo bar":
                                    var v57 float64 = 0.0
                                    _ = v57
                                }
                                _ = v52
                            }
                            switch v32 {
                            case "foo bar":
                                var v58 float64 = 1e10
                                _ = v58
                            case "\n\t":
                                var v59 int16 = v1
                                _ = v59
                            case "":
                                var v60 float32 = float32(0.1)
                                _ = v60
                            }
                        }
                    } else {
                        if (1 <= 127) {
                            for fi61 := 0; fi61 < 14; fi61++ {
                                var v62 uint64 = 2
                                delete(v4, "")
                                _ = fi61
                                _ = v62
                            }
                            fmt.Println(v18)
                            var v63 uint16 = 32767
                            _ = v63
                        } else {
                            v18 = 1
                            var v64 float64 = (3.14159 * (3.14159 / float64(1e10)))
                            _ = v64
                        }
                    }
                }
            case "a":
                switch v32 {
                case "foo bar":
                    _ = len(v2)
                case "a":
                    for fi65 := 0; fi65 < 8; fi65++ {
                        v32 = ""
                        _ = fi65
                    }
                case "hello":
                    if (!(true && (42 <= (v1 * (-(v1)))))) {
                        for fi66 := 0; fi66 < 16; fi66++ {
                            var v67 string = "hello"
                            _ = fi66
                            _ = v67
                        }
                    }
                }
            }
        }
    case 67:
        delete(v5, (v28 >> ((v18) & 63)))
    default:
        for fi68 := 0; fi68 < 8; fi68++ {
            fmt.Println(v32)
            v32 = "a"
            var v69 int8 = -128
            _ = fi68
            _ = v69
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v5
}
