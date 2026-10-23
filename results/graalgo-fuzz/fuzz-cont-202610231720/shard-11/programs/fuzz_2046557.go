package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int16; F2 int64 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint8, p1 bool, p2 uint8) int {
    if false {
        v6 := []int{-1000, 42, 32768, 174}
        v6 = append(v6, 32767)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v6
    }
    v7 := []int{}
    for fi8 := 0; fi8 < 6; fi8++ {
        v7 = append(v7, (fi8 << ((42) & 63)))
        if (127 > -128) {
            for fi9 := 0; fi9 < 15; fi9++ {
                p1 = false
                var v10 int = (fi9 - (fi8 & (fi8 ^ (fi9 & (fi9 + fi9)))))
                switch fi9 {
                case -756:
                    fmt.Println(p2, p1)
                case -2147483648:
                    var v11 uint8 = (p0 & (p2 / ((p2) | 1)))
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
            var v12 int8 = 127
            _ = v12
        } else {
            if p1 {
                v13 := GenS1{F0: ((("hello" + "\n\t") + "") + "a")}
                _ = v13
            } else {
                var v14 int = 42
                _ = v14
            }
            fmt.Println(fi8)
            p2 -= p0
        }
        if len(v7) > 0 {
            v7[0] = fi8
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32768
}
func genFn1(p0 uint8) uint16 {
    if false {
        fmt.Println(p0)
        if false {
            if false {
                p0 -= (p0 << ((718) & 63))
                var v15 int = 2147483647
                _ = v15
            }
            var v16 uint = 2
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := []int{(fi17 << (((^(v16))) & 63)), fi17, fi17}
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    } else {
        v19 := []int{genFn0(60, true, (p0 % (((p0 + (p0 | p0))) | 1))), 4294967295, 42}
        var v20 string = "hello"
        _ = v19
        _ = v20
    }
    fmt.Println(p0)
    fmt.Println(p0)
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS1{F0: (("hello" + "") + "")}
        var v23 uint32 = 2147483648
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: ("\n\t" + "foo bar")}
    _ = p0
    _ = v24
    return 411
}

func main() {
    var v1 string = ((((("\n\t" + ("hello" + "foo bar")) + ("" + "hello")) + "\n\t") + "\n\t") + (("foo bar" + "\n\t") + "hello"))
    var v2 float32 = 3.14159
    var v3 uint64 = 0
    v4 := []string{("foo bar" + ("\n\t" + ("hello" + v1))), ("hello" + v1)}
    v5 := []uint32{}
    switch 32767 {
    case -234:
        var v25 int8 = 1
        _ = v25
    case 0:
        if (1.5 >= 0.0) {
            v1 = ""
        }
    case 32767:
        var v26 string = "\n\t"
        _ = v26
    }
    if (false && false) {
        if (45 < 255) {
            v4 = append(v4, v1)
            v27 := GenS1{F0: v1}
            fmt.Println(v2)
            _ = v27
        }
        var v28 int64 = 448
        var v29 int16 = 0
        _ = v28
        _ = v29
    } else {
        v30 := GenS1{F0: v1}
        var v31 int16 = -655
        _ = v30
        _ = v31
    }
    v2 += (v2 * v2)
    for fi32 := 0; fi32 < 20; fi32++ {
        v33 := GenS1{F0: (v1 + "hello")}
        var v34 uint = 2
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        for fi36 := 0; fi36 < 5; fi36++ {
            switch "foo bar" {
            case "hello":
                if true {
                    fmt.Println(v3)
                    v2 = (-2.25 / v2)
                }
            case "":
                v3 += v3
            }
            var v37 int64 = 9223372036854775807
            _ = fi36
            _ = v37
        }
        v38 := GenS2{F0: 772, F1: 606, F2: 9223372036854775807}
        switch 128 {
        case 1:
            v39 := GenS3{F0: 1, F1: 2}
            _ = v39
        case -2147483648:
            v2 += v2
        default:
            for fi40 := 0; fi40 < 7; fi40++ {
                if false {
                    if len(v4) > 0 {
                        v4[0] = (((v1 + ((((v1 + ((("\n\t" + v1) + "a") + ((v1 + (("hello" + (v1 + (("a" + "a") + v1))) + "a")) + v1))) + v1) + v1) + (v1 + "a"))) + v1) + v1)
                    }
                }
                fmt.Println(v2, v3, v1)
                v1 = v1
                _ = fi40
            }
        }
        _ = fi35
        _ = v38
    }
    var v41 uint = 927
    if (float64(0.0) == float64(0.0)) {
        fmt.Println(v1, v1)
    }
    v1 = ("a" + v1)
    for fi42 := 0; fi42 < 20; fi42++ {
        _ = len(v5)
        v43 := GenS1{F0: ("a" + ("a" + v1))}
        _ = fi42
        _ = v43
    }
    v3 = 0
    if false {
        var v44 float64 = 3.14159
        var v45 uint = 440
        _ = v44
        _ = v45
    } else {
        v46 := GenS4{F0: 42, F1: 856}
        switch v1 {
        case "a":
            fmt.Println(v41, v3)
        case "hello":
            for fi47 := 0; fi47 < 12; fi47++ {
                if false {
                    fmt.Println(v3, fi47, fi47)
                } else {
                    if (!false) {
                        v3 = v3
                        v41 = v41
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                v48 := GenS2{F0: 2147483648, F1: 42, F2: -408}
                _ = fi47
                _ = v48
            }
        case "":
            var v49 int32 = 42
            _ = v49
        default:
            v50 := GenS4{F0: 9223372036854775807, F1: 2}
            _ = v50
        }
        if (1 == 42) {
            fmt.Println(v2, v3)
            switch genFn0(128, true, 107) {
            case 128:
                for fi51 := 0; fi51 < 10; fi51++ {
                    if ((fi51 | fi51) > (fi51 / ((-874) | 1))) {
                        if true {
                            if (32767 >= int16(v41)) {
                                switch (fi51 & (-(fi51))) {
                                case -604:
                                    v2 -= ((v2 - float32((1e10 + -2.25))) - 1e10)
                                case 128:
                                    v52 := GenS3{F0: -9223372036854775808, F1: 196}
                                    _ = v52
                                case 2147483647:
                                    if ((16 != 1) && (!(false && (!false)))) {
                                        var v53 int = (fi51 + fi51)
                                        v3 = v3
                                        _ = v53
                                    }
                                }
                            } else {
                                var v54 uint8 = 30
                                if true {
                                    if (true || false) {
                                        v55 := GenS5{F0: v3}
                                        var v56 uint64 = (v3 ^ v3)
                                        v57 := GenS2{F0: 982, F1: genFn1(v54), F2: 1}
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                    } else {
                                        if (false && (612 >= 255)) {
                                            for fi58 := 0; fi58 < 10; fi58++ {
                                                fmt.Println(v41)
                                                fmt.Println(v41)
                                                v3 = v3
                                                _ = fi58
                                            }
                                            for fi59 := 0; fi59 < 20; fi59++ {
                                                var v60 int8 = 0
                                                _ = fi59
                                                _ = v60
                                            }
                                            fmt.Println(v41)
                                        }
                                    }
                                    if (0 > genFn1(127)) {
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                        if (2147483647 < 2147483648) {
                                            var v61 int8 = 1
                                            var v62 float32 = float32(0.1)
                                            v4 = append(v4, v1)
                                            _ = v61
                                            _ = v62
                                        } else {
                                            fmt.Println(v1, v41)
                                            if (!(v41 != v41)) {
                                                for fi63 := 0; fi63 < 18; fi63++ {
                                                    _ = len(v5)
                                                    _ = len(v5)
                                                    if (((v54 % ((v54) | 1)) < (v54 / ((v54) | 1))) && (((v1 + "\n\t") + "\n\t") < (v1 + v1))) {
                                                        var v64 int16 = 32767
                                                        _ = v64
                                                    }
                                                    v54 = 42
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    switch v1 {
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "":
                                                            var v65 int64 = -554
                                                            _ = v65
                                                        default:
                                                            v54 += 255
                                                        }
                                                    }
                                                    _ = fi63
                                                }
                                                v1 = v1
                                                v1 = v1
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = v1
                                            }
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            v2 = v2
                                        default:
                                            for fi66 := 0; fi66 < 1; fi66++ {
                                                v67 := GenS6{F0: -9223372036854775808, F1: 32767, F2: -433}
                                                v68 := GenS7{F0: v3, F1: 1}
                                                _ = fi66
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                    }
                                    v54 += (^(v54))
                                }
                                var v69 uint8 = (v54 ^ v54)
                                _ = v54
                                _ = v69
                            }
                            var v70 uint32 = 65535
                            v71 := GenS8{F0: -678, F1: -32768}
                            _ = v70
                            _ = v71
                        } else {
                            v3 += v3
                        }
                        if len(v5) > 0 {
                            v5[0] = 814
                        }
                    } else {
                        switch 2147483647 {
                        case 9223372036854775807:
                            v72 := GenS9{F0: 178}
                            _ = v72
                        default:
                            v1 = (v1 + "hello")
                        }
                        v5 = append(v5, 710)
                    }
                    _ = fi51
                }
            case 2147483647:
                fmt.Println(v1, v3, v3)
            default:
                fmt.Println(v3)
            }
            v3 += v3
        } else {
            v2 = 1.5
            switch ("" + ("hello" + "")) {
            case "a":
                switch "\n\t" {
                case "":
                    v73 := GenS2{F0: 106, F1: uint16(v3), F2: 9223372036854775807}
                    _ = v73
                default:
                    fmt.Println(v3, v3, v41)
                }
            }
            v74 := GenS4{F0: -635, F1: v3}
            _ = v74
        }
        _ = v46
    }
    for fi75 := 0; fi75 < 4; fi75++ {
        v76 := GenS9{F0: 128}
        _ = fi75
        _ = v76
    }
    if false {
        var v77 int32 = 42
        if false {
            var v78 string = v1
            _ = v78
        } else {
            for fi79 := 0; fi79 < 19; fi79++ {
                var v80 uint16 = 65535
                if (2 == uint8(v80)) {
                    v81 := GenS10{F0: -323, F1: 493}
                    _ = v81
                }
                _ = fi79
                _ = v80
            }
        }
        var v82 int = 128
        _ = v77
        _ = v82
    }
    var v83 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v83
}
