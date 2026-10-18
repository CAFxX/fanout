package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 int32; F1 int16; F2 int }
type GenS12 struct { F0 int64; F1 int32; F2 uint64 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0(p0 uint16) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{fi4, (fi4 << ((766) & 63)), (fi4 | (fi4 ^ fi4))}
        var v6 bool = true
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{}
        if true {
            p0 = (p0 ^ 173)
            v9 := GenS1{F0: -817, F1: p0, F2: 4294967295}
            switch "a" {
            case "hello":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case "foo bar":
                var v10 string = "foo bar"
                _ = v10
            default:
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 int = -9223372036854775808
                    _ = fi11
                    _ = v12
                }
            }
            _ = v9
        }
        var v13 bool = (42 > 1)
        _ = fi7
        _ = v13
        _ = v8
    }
    _ = p0
    return false
}

func main() {
    var v1 uint32 = 42
    var v2 uint = 2
    v3 := []int{4294967295, 4294967295, -592}
    for fi14 := 0; fi14 < 5; fi14++ {
        v2 = uint(v1)
        v3 = append(v3, (fi14 | fi14))
        _ = fi14
    }
    if false {
        v2 = 1
        if false {
            var v15 string = ""
            var v16 float64 = 3.14159
            fmt.Println(v16, v15)
            _ = v15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 1; fi17++ {
            if (65535 <= 65535) {
                switch -110 {
                case -285:
                    switch (("\n\t" + "a") + "foo bar") {
                    case "a":
                        var v18 uint8 = 127
                        _ = v18
                    default:
                        if true {
                            var v19 uint32 = (v1 << ((4294967295) & 63))
                            _ = v19
                        } else {
                            switch (fi17 ^ (fi17 | int(v1))) {
                            case 4294967295:
                                var v20 string = "hello"
                                _ = v20
                            case 405:
                                if true {
                                    v2 -= (^(v2))
                                } else {
                                    v3 = append(v3, (-(fi17)))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v21 int = (^(fi17))
                                    _ = v21
                                }
                            case -2147483648:
                                if len(v3) > 0 {
                                    v3[0] = fi17
                                }
                            }
                            v1 -= v1
                            switch "\n\t" {
                            case "a":
                                var v22 int64 = 1
                                _ = v22
                            case "":
                                v2 += 1
                            case "\n\t":
                                var v23 int8 = -2
                                _ = v23
                            }
                        }
                    }
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 11; fi24++ {
                        var v25 uint = (v2 + (^(v2)))
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi24
                        _ = v25
                    }
                case 0:
                    for fi26 := 0; fi26 < 10; fi26++ {
                        switch (-(fi26)) {
                        case 481:
                            fmt.Println(fi17, fi17)
                        default:
                            v1 += (^(v1))
                        }
                        var v27 uint16 = 2
                        _ = fi26
                        _ = v27
                    }
                }
                v1 = (v1 | (v1 | (^(v1))))
                switch "\n\t" {
                case "\n\t":
                    switch "hello" {
                    case "\n\t":
                        var v28 int16 = 1000
                        _ = v28
                    }
                default:
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v30 := GenS2{F0: (^(fi17)), F1: 115}
                        _ = fi29
                        _ = v30
                    }
                }
            } else {
                var v31 bool = true
                for fi32 := 0; fi32 < 19; fi32++ {
                    v33 := GenS2{F0: -2147483648, F1: 1}
                    v1 = 4294967295
                    _ = fi32
                    _ = v33
                }
                v31 = v31
                _ = v31
            }
            switch fi17 {
            case 9223372036854775807:
                v34 := GenS3{F0: 0, F1: fi17}
                _ = v34
            default:
                v2 = v2
            }
            _ = fi17
        }
        if false {
            if true {
                v35 := GenS1{F0: 499, F1: 255, F2: v1}
                _ = v35
            } else {
                v2 = (v2 >> (((v2 << ((v2) & 63))) & 63))
                var v36 uint32 = (v1 | v1)
                _ = v36
            }
        } else {
            v37 := GenS4{F0: 1, F1: -818}
            if false {
                v1 += v1
                switch ("a" + "") {
                case "":
                    v38 := GenS5{F0: ""}
                    _ = v38
                default:
                    v1 = v1
                }
            }
            v39 := GenS4{F0: 0, F1: 32767}
            _ = v37
            _ = v39
        }
    }
    v40 := GenS5{F0: ""}
    v1 += v1
    var v41 float64 = -2.25
    for fi42 := 0; fi42 < 13; fi42++ {
        v43 := GenS6{F0: -2147483648, F1: false}
        v1 = 2147483647
        if (fi42 >= fi42) {
            switch ("a" + "") {
            case "\n\t":
                if false {
                    v3 = append(v3, fi42)
                }
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = (fi42 | fi42)
                }
            case "hello":
                v2 -= 255
            default:
                for fi44 := 0; fi44 < 20; fi44++ {
                    v1 -= 471
                    if (!(!true)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (2147483647 >= -1) {
                            var v45 int = 4294967295
                            v46 := GenS4{F0: 2, F1: 1000}
                            var v47 int16 = -51
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        fmt.Println(v1, v2, fi44)
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v48 uint64 = 42
                        var v49 uint16 = 168
                        _ = v48
                        _ = v49
                    }
                    _ = fi44
                }
            }
            var v50 uint8 = 127
            v51 := GenS7{F0: (!false)}
            _ = v50
            _ = v51
        } else {
            fmt.Println(v41, v41, v2)
        }
        _ = fi42
        _ = v43
    }
    v41 -= (1e10 + (v41 / v41))
    for fi52 := 0; fi52 < 15; fi52++ {
        var v53 float64 = 3.14159
        v1 = 726
        _ = fi52
        _ = v53
    }
    var v54 float64 = float64(3.14159)
    v55 := GenS6{F0: -2147483648, F1: (1 == -1)}
    switch (("foo bar" + "foo bar") + "") {
    case "hello":
        for fi56 := 0; fi56 < 18; fi56++ {
            switch "" {
            case "":
                if false {
                    if (!(!false)) {
                        if true {
                            v41 = v54
                            v57 := GenS2{F0: (fi56 / (((fi56 - (fi56 >> ((18446744073709551615) & 63)))) | 1)), F1: 0}
                            _ = v57
                        }
                    } else {
                        var v58 bool = (!false)
                        v59 := GenS8{F0: (v1 << ((v2) & 63)), F1: 32768}
                        for fi60 := 0; fi60 < 16; fi60++ {
                            if len(v3) > 0 {
                                v3[0] = fi56
                            }
                            _ = fi60
                        }
                        _ = v58
                        _ = v59
                    }
                } else {
                    v41 += float64(0.1)
                    switch fi56 {
                    case 137:
                        v54 -= v54
                    case -132:
                        v3 = append(v3, -1)
                    default:
                        v61 := GenS8{F0: v1, F1: fi56}
                        _ = v61
                    }
                    if false {
                        if (65535 != 0) {
                            v62 := GenS2{F0: fi56, F1: uint16(v1)}
                            v63 := GenS9{F0: "foo bar", F1: 42}
                            v1 -= 1
                            _ = v62
                            _ = v63
                        } else {
                            if (!(fi56 == fi56)) {
                                v54 = 1.5
                            } else {
                                v2 += v2
                                for fi64 := 0; fi64 < 19; fi64++ {
                                    var v65 uint32 = v1
                                    _ = fi64
                                    _ = v65
                                }
                            }
                            v2 = 447
                        }
                        v2 -= (^(v2))
                        if false {
                            v1 = v1
                        } else {
                            if false {
                                v54 += ((float64((-2.25 - -2.25)) - v41) * (v54 - (v41 + v54)))
                                v1 += v1
                            }
                            switch fi56 {
                            case 0:
                                for fi66 := 0; fi66 < 7; fi66++ {
                                    v54 = ((0.1 * 0.1) / 0.1)
                                    fmt.Println(v2, fi56, v1)
                                    var v67 uint32 = 6
                                    var v68 float64 = v54
                                    if (0 == 0) {
                                        if (!true) {
                                            if true {
                                                var v69 uint8 = 1
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v69
                                            } else {
                                                v70 := GenS1{F0: 1, F1: 103, F2: (v67 * v1)}
                                                v3 = append(v3, 127)
                                                _ = v70
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                v71 := GenS10{F0: -1, F1: 255}
                                                var v72 uint8 = 42
                                                _ = v71
                                                _ = v72
                                            case "":
                                                switch (fi56 & (fi66 >> ((0) & 63))) {
                                                case 4294967295:
                                                    v73 := GenS5{F0: ("hello" + "a")}
                                                    _ = v73
                                                case 32767:
                                                    var v74 uint8 = 243
                                                    _ = v74
                                                }
                                            default:
                                                if true {
                                                    if false {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v75 uint64 = 0
                                                            _ = v75
                                                        case "foo bar":
                                                            if (1.5 <= 1.5) {
                                                                switch fi56 {
                                                                case 715:
                                                                    v1 -= (v67 & v1)
                                                                }
                                                                _ = len(v3)
                                                                switch (fi66 - (fi56 % (((^(fi56))) | 1))) {
                                                                case -2147483648:
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v1 = (v67 & 1)
                                                                    case "":
                                                                        v76 := GenS1{F0: 0, F1: 821, F2: v1}
                                                                        _ = v76
                                                                    }
                                                                case 522:
                                                                    var v77 int64 = 1
                                                                    _ = v77
                                                                case 124:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                                v78 := GenS10{F0: 0, F1: 0}
                                                                var v79 uint32 = 632
                                                                v67 = (v67 ^ (v79 / ((v79) | 1)))
                                                                fmt.Println(v41, v1, v79)
                                                                v79 = 749
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                        }
                                                    } else {
                                                        v68 += (3.14159 - 1e10)
                                                        v1 = 2147483648
                                                        switch (fi66 + (fi66 ^ fi66)) {
                                                        case 0:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v67 -= v1
                                                            }
                                                        case 516:
                                                            v68 += v54
                                                        }
                                                    }
                                                }
                                            }
                                            v80 := GenS3{F0: 255, F1: fi66}
                                            _ = v80
                                        } else {
                                            if ((1 < 0) && (!false)) {
                                                var v81 bool = false
                                                v2 -= v2
                                                _ = v81
                                            } else {
                                                if ((true || ((v2 & 42) == v2)) || false) {
                                                    fmt.Println(fi56, fi56)
                                                } else {
                                                    v82 := GenS6{F0: 1, F1: (9223372036854775807 != -1)}
                                                    switch ("\n\t" + "\n\t") {
                                                    case "\n\t":
                                                        fmt.Println(v54, v68)
                                                    case "a":
                                                        var v83 int8 = 1
                                                        _ = v83
                                                    }
                                                    v84 := GenS11{F0: 2147483647, F1: 1, F2: -561}
                                                    _ = v82
                                                    _ = v84
                                                }
                                                var v85 int32 = 2147483647
                                                _ = v85
                                            }
                                            v86 := GenS12{F0: -819, F1: -163, F2: 377}
                                            v2 -= (v2 / ((v2) | 1))
                                            _ = v86
                                        }
                                    }
                                    _ = fi66
                                    _ = v67
                                    _ = v68
                                }
                            case 1:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v87 := GenS13{F0: 32767}
                            _ = v87
                        }
                    }
                }
            default:
                v88 := GenS14{F0: ("hello" + ("a" + ("foo bar" + "\n\t"))), F1: 0, F2: 4294967295}
                _ = v88
            }
            v54 -= (v41 * (v54 * ((v54 / (v41 / v41)) - v54)))
            _ = fi56
        }
    case "":
        v3 = append(v3, 2147483647)
    default:
        v54 -= v54
    }
    if false {
        if (false && (!(0 == 1000))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            if true {
                var v89 uint8 = 167
                v1 = v1
                _ = v89
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v54
    _ = v55
}
