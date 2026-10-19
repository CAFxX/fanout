package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 int16; F1 int32; F2 uint16 }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS15 struct { F0 int; F1 string; F2 uint32 }
type GenS16 struct { F0 int32; F1 int16; F2 int8 }
type GenS17 struct { F0 bool; F1 int; F2 int32 }
type GenS18 struct { F0 uint64; F1 uint16; F2 string }
type GenS19 struct { F0 int16; F1 bool; F2 int32 }
type GenS20 struct { F0 string; F1 uint64 }
type GenS21 struct { F0 uint; F1 uint32; F2 int32 }
type GenS22 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS23 struct { F0 uint }
func genFn0() int {
    v5 := GenS1{F0: 2}
    switch -9223372036854775808 {
    case 32767:
        v6 := []int{-429}
        _ = v6
    case 556:
        var v7 string = "a"
        _ = v7
    case -840:
        var v8 float64 = (0.0 - (-2.25 * (((3.14159 - (1.5 + 3.14159)) + (((1.5 + (0.1 * ((-2.25 - 0.1) / 1.0))) + 1e10) - (1e10 + (-2.25 / 1.0)))) * 0.1)))
        _ = v8
    }
    var v9 int8 = 1
    v10 := GenS1{F0: 54}
    v11 := GenS1{F0: 128}
    _ = v10
    _ = v11
    _ = v5
    _ = v9
    return 32768
}
func genFn1() uint32 {
    if true {
        v12 := []int{619}
        var v13 int16 = 42
        _ = v12
        _ = v13
    } else {
        var v14 float32 = ((1e10 + 1e10) - 3.14159)
        var v15 int = -9223372036854775808
        _ = v14
        _ = v15
    }
    var v16 uint16 = 42
    if false {
        v17 := GenS2{F0: (false || (293 >= 603)), F1: 255}
        fmt.Println(v16)
        _ = v17
    }
    _ = v16
    return 519
}

func main() {
    var v1 float64 = 0.1
    var v2 int8 = -111
    var v3 int32 = -451
    var v4 uint64 = 42
    switch 585 {
    case -2147483648:
        v3 = (^(v3))
    default:
        if true {
            var v18 uint16 = 32767
            var v19 int16 = 1000
            if true {
                if true {
                    v4 = v4
                    for fi20 := 0; fi20 < 8; fi20++ {
                        fmt.Println(fi20, v3, v3)
                        _ = fi20
                    }
                    var v21 int32 = -1
                    _ = v21
                }
            } else {
                var v22 uint32 = 2147483648
                v1 = 1e10
                var v23 uint8 = 42
                _ = v22
                _ = v23
            }
            _ = v18
            _ = v19
        } else {
            switch -2147483648 {
            case 128:
                var v24 bool = (127 <= 0)
                _ = v24
            }
            if (-276 != 721) {
                for fi25 := 0; fi25 < 15; fi25++ {
                    if ((!false) || ("\n\t" == ((("hello" + "a") + "\n\t") + ("\n\t" + (("a" + "") + ("foo bar" + "foo bar")))))) {
                        fmt.Println(v4)
                        var v26 uint32 = 2147483648
                        if (3.14159 > v1) {
                            v27 := GenS2{F0: false, F1: 255}
                            v28 := []int{}
                            _ = v27
                            _ = v28
                        }
                        _ = v26
                    }
                    v29 := []int{(fi25 ^ (fi25 / (((fi25 >> ((275) & 63))) | 1))), (fi25 >> ((1) & 63))}
                    if (65535 > 824) {
                        switch "\n\t" {
                        case "a":
                            v30 := GenS3{F0: 216, F1: v4}
                            _ = v30
                        case "\n\t":
                            switch "hello" {
                            case "foo bar":
                                if (false && ((v2 / ((v2) | 1)) == (v2 - (v2 / ((v2) | 1))))) {
                                    if (0 > 0) {
                                        v3 = -854
                                        if ((true || true) && true) {
                                            for fi31 := 0; fi31 < 20; fi31++ {
                                                switch (("" + "") + ("" + ("hello" + "foo bar"))) {
                                                case "foo bar":
                                                    var v32 uint8 = 42
                                                    _ = v32
                                                }
                                                var v33 int = genFn0()
                                                _ = fi31
                                                _ = v33
                                            }
                                        } else {
                                            v1 = ((float64(1.5) + v1) + (v1 + ((v1 / (v1 / 1.5)) - (float64(-2.25) + float64(1.5)))))
                                        }
                                    } else {
                                        if (0.1 == 1e10) {
                                            for fi34 := 0; fi34 < 18; fi34++ {
                                                var v35 int32 = v3
                                                var v36 uint8 = 1
                                                _ = fi34
                                                _ = v35
                                                _ = v36
                                            }
                                            var v37 uint = 1
                                            _ = v37
                                        }
                                        if len(v29) > 0 {
                                            v29[0] = fi25
                                        }
                                        switch "hello" {
                                        case "a":
                                            switch 32768 {
                                            case 2147483647:
                                                v1 = 0.1
                                            case 563:
                                                if len(v29) > 0 {
                                                    v29[0] = (fi25 >> ((4294967295) & 63))
                                                }
                                            default:
                                                v38 := GenS3{F0: 281, F1: v4}
                                                _ = v38
                                            }
                                        case "hello":
                                            v4 += v4
                                        case "\n\t":
                                            var v39 float32 = float32(float64(0.0))
                                            _ = v39
                                        }
                                    }
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    v3 = -2147483648
                                } else {
                                    for fi40 := 0; fi40 < 7; fi40++ {
                                        if true {
                                            if (!(int32(v4) <= v3)) {
                                                v29 = append(v29, fi25)
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(fi40, v1)
                                                    var v41 int = fi40
                                                    _ = v41
                                                case "foo bar":
                                                    if (255 >= 0) {
                                                        v2 -= -1
                                                        v42 := GenS1{F0: 1}
                                                        v43 := GenS2{F0: true, F1: 65535}
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                                fmt.Println(v1, v2)
                                                if ((fi25 + (fi25 * fi40)) == genFn0()) {
                                                    var v44 uint32 = 2147483647
                                                    switch ("hello" + ("\n\t" + "\n\t")) {
                                                    case "":
                                                        v44 -= v44
                                                    case "hello":
                                                        if false {
                                                            if (false && false) {
                                                                _ = len(v29)
                                                                var v45 uint8 = 0
                                                                _ = v45
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v4 = (v4 >> ((uint(v44)) & 63))
                                                    default:
                                                        if (!false) {
                                                            v44 = (v44 + v44)
                                                            if len(v29) >= 2 {
                                                                v29 = v29[1:len(v29)]
                                                            }
                                                            v46 := GenS3{F0: 297, F1: (v4 - 4294967295)}
                                                            _ = v46
                                                        } else {
                                                            fmt.Println(v1, fi40, v44)
                                                            if ((fi25 * genFn0()) <= int(v2)) {
                                                                v3 = 2147483647
                                                                v47 := GenS1{F0: 42}
                                                                v4 = v4
                                                                fmt.Println(fi40)
                                                                v48 := GenS3{F0: 65535, F1: (^(v4))}
                                                                switch 1 {
                                                                case -206:
                                                                    var v49 float32 = 1e10
                                                                    _ = v49
                                                                case 42:
                                                                    if true {
                                                                        if (!(0.1 < 0.0)) {
                                                                            if false {
                                                                                switch -1 {
                                                                                case 127:
                                                                                    _ = len(v29)
                                                                                case 32767:
                                                                                    if false {
                                                                                        fmt.Println(v4, v44)
                                                                                    } else {
                                                                                        switch "foo bar" {
                                                                                        case "a":
                                                                                            var v50 uint16 = 1
                                                                                            _ = v50
                                                                                        case "\n\t":
                                                                                            var v51 int64 = 0
                                                                                            _ = v51
                                                                                        case "foo bar":
                                                                                            var v52 uint = uint(v44)
                                                                                            _ = v52
                                                                                        default:
                                                                                            var v53 bool = (("\n\t" + "hello") == ("foo bar" + "\n\t"))
                                                                                            _ = v53
                                                                                        }
                                                                                        fmt.Println(v3, v3, v2)
                                                                                    }
                                                                                case 0:
                                                                                    var v54 float64 = (0.1 * 3.14159)
                                                                                    _ = v54
                                                                                default:
                                                                                    v1 += (float64(0.1) / 1.0)
                                                                                }
                                                                                v55 := GenS3{F0: 395, F1: (v4 + v4)}
                                                                                _ = v55
                                                                            } else {
                                                                                v4 = v4
                                                                            }
                                                                            v56 := GenS4{F0: ("" + "\n\t"), F1: -474}
                                                                            _ = v56
                                                                        } else {
                                                                            fmt.Println(v2, v4, v3)
                                                                        }
                                                                        v1 += v1
                                                                        var v57 uint32 = v44
                                                                        _ = v57
                                                                    }
                                                                case 53:
                                                                    var v58 bool = (v4 < (v4 & (^(v4))))
                                                                    _ = v58
                                                                default:
                                                                    var v59 float32 = float32(-2.25)
                                                                    _ = v59
                                                                }
                                                                v4 += v4
                                                                if false {
                                                                    var v60 int16 = 32767
                                                                    v44 = v44
                                                                    _ = v60
                                                                } else {
                                                                    if len(v29) >= 2 {
                                                                        v29 = v29[1:len(v29)]
                                                                    }
                                                                    v44 = (v44 ^ 407)
                                                                }
                                                                v29 = append(v29, int(v2))
                                                                var v61 int = (fi25 & (fi40 & -1))
                                                                if (((v3 ^ v3) == 754) || false) {
                                                                    if (true && ((v4 <= 0) && true)) {
                                                                        v2 = 100
                                                                        if true {
                                                                            if true {
                                                                                v1 = ((v1 / v1) / (-2.25 - v1))
                                                                                switch genFn0() {
                                                                                case 1:
                                                                                    fmt.Println(fi40)
                                                                                case 32767:
                                                                                    v4 -= (v4 + (v4 * v4))
                                                                                case 32768:
                                                                                    v2 = -1
                                                                                }
                                                                            }
                                                                            fmt.Println(v44, v4, v2)
                                                                            if (0.1 == -2.25) {
                                                                                v62 := GenS2{F0: (v1 == -2.25), F1: 645}
                                                                                _ = v62
                                                                            }
                                                                            var v63 int = genFn0()
                                                                            if false {
                                                                                switch ("hello" + "") {
                                                                                case "":
                                                                                    v64 := GenS1{F0: 1}
                                                                                    _ = v64
                                                                                case "a":
                                                                                    v1 = (v1 + ((3.14159 - (((float64(((0.1 + -2.25) / 0.1)) * (1.5 - -2.25)) - 1.5) - 1e10)) + v1))
                                                                                case "\n\t":
                                                                                    if (1 <= (-(v3))) {
                                                                                        fmt.Println(v63, v3)
                                                                                        var v65 uint = 4294967295
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            var v66 bool = false
                                                                                            _ = v66
                                                                                        case "":
                                                                                            v44 += (^(v44))
                                                                                        default:
                                                                                            v67 := GenS5{F0: 32767}
                                                                                            _ = v67
                                                                                        }
                                                                                        _ = v65
                                                                                    } else {
                                                                                        v2 -= v2
                                                                                        switch (-(fi25)) {
                                                                                        case 23:
                                                                                            v2 += (v2 / ((127) | 1))
                                                                                        case 1:
                                                                                            v3 -= (-(v3))
                                                                                        case 0:
                                                                                            if true {
                                                                                                var v68 int64 = 0
                                                                                                v1 += v1
                                                                                                v68 = v68
                                                                                                _ = v68
                                                                                            }
                                                                                        }
                                                                                        v4 += v4
                                                                                        if (389 == -885) {
                                                                                            switch (fi40 >> ((18446744073709551615) & 63)) {
                                                                                            case 42:
                                                                                                if false {
                                                                                                    var v69 int8 = (v2 >> ((4294967295) & 63))
                                                                                                    var v70 uint8 = 0
                                                                                                    _ = v69
                                                                                                    _ = v70
                                                                                                } else {
                                                                                                    fmt.Println(v1, v1)
                                                                                                    var v71 uint32 = 924
                                                                                                    _ = v71
                                                                                                }
                                                                                                v72 := GenS5{F0: -32768}
                                                                                                _ = v72
                                                                                            case 2147483647:
                                                                                                v73 := GenS3{F0: 255, F1: v4}
                                                                                                _ = v73
                                                                                            default:
                                                                                                if true {
                                                                                                    fmt.Println(v61)
                                                                                                    fmt.Println(v1, v3)
                                                                                                    v2 = (v2 / (((v2 ^ (^(v2)))) | 1))
                                                                                                } else {
                                                                                                    if false {
                                                                                                        var v74 float32 = (0.0 - (float32(v1) + float32((float64((1.5 + ((-2.25 * -2.25) - 1.5))) / 3.14159))))
                                                                                                        v63 = (fi40 ^ v63)
                                                                                                        var v75 uint64 = (v4 % ((0) | 1))
                                                                                                        _ = v74
                                                                                                        _ = v75
                                                                                                    }
                                                                                                    if true {
                                                                                                        v44 = (v44 << ((305) & 63))
                                                                                                    } else {
                                                                                                        v76 := GenS6{F0: (v44 << ((255) & 63)), F1: 946}
                                                                                                        v29 = append(v29, (v61 ^ fi25))
                                                                                                        fmt.Println(v63, v44)
                                                                                                        _ = v76
                                                                                                    }
                                                                                                    v77 := GenS4{F0: ("foo bar" + ("a" + "")), F1: -2147483648}
                                                                                                    _ = v77
                                                                                                }
                                                                                            }
                                                                                            v4 = v4
                                                                                            var v78 int8 = (v2 / ((-1) | 1))
                                                                                            _ = v78
                                                                                        } else {
                                                                                            fmt.Println(v44)
                                                                                            v61 += v61
                                                                                            var v79 uint8 = 101
                                                                                            v80 := GenS5{F0: 32767}
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v2, v4, v4)
                                                                                }
                                                                                fmt.Println(v63, v61)
                                                                            }
                                                                            v63 += fi25
                                                                            _ = v63
                                                                        } else {
                                                                            var v81 uint16 = 42
                                                                            var v82 uint64 = 42
                                                                            var v83 uint16 = uint16(v4)
                                                                            _ = v81
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                        switch (("foo bar" + "\n\t") + ("\n\t" + "")) {
                                                                        case "foo bar":
                                                                            v84 := GenS4{F0: ("foo bar" + "\n\t"), F1: (-(v3))}
                                                                            _ = v84
                                                                        case "":
                                                                            var v85 int8 = (v2 * v2)
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(fi25)
                                                                            var v86 uint = uint(v44)
                                                                            v4 -= (v4 + (v4 << ((255) & 63)))
                                                                            _ = v86
                                                                        }
                                                                        switch 515 {
                                                                        case 4294967295:
                                                                            v87 := GenS4{F0: "", F1: 891}
                                                                            _ = v87
                                                                        default:
                                                                            v88 := GenS7{F0: v44, F1: 150}
                                                                            _ = v88
                                                                        }
                                                                        v89 := GenS4{F0: "foo bar", F1: (v3 ^ v3)}
                                                                        _ = len(v29)
                                                                        _ = v89
                                                                    }
                                                                    if true {
                                                                        if len(v29) >= 2 {
                                                                            v29 = v29[1:len(v29)]
                                                                        }
                                                                        var v90 int64 = -1
                                                                        var v91 uint8 = 255
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        fmt.Println(v1, fi25)
                                                                    case "":
                                                                        v1 -= float64(0.0)
                                                                        fmt.Println(v1)
                                                                        var v92 int8 = 1
                                                                        _ = v92
                                                                    default:
                                                                        var v93 uint32 = v44
                                                                        _ = v93
                                                                    }
                                                                }
                                                                if (v4 < (v4 << ((65535) & 63))) {
                                                                    if ((-(fi25)) < genFn0()) {
                                                                        v4 += (^(v4))
                                                                    }
                                                                    v44 = (v44 & v44)
                                                                }
                                                                var v94 string = "\n\t"
                                                                var v95 int16 = 32767
                                                                var v96 string = v94
                                                                fmt.Println(v44, v94)
                                                                _ = v47
                                                                _ = v48
                                                                _ = v61
                                                                _ = v94
                                                                _ = v95
                                                                _ = v96
                                                            } else {
                                                                switch 235 {
                                                                case -728:
                                                                    var v97 string = "a"
                                                                    _ = v97
                                                                case 504:
                                                                    if (!false) {
                                                                        if (!(!true)) {
                                                                            v44 += genFn1()
                                                                            v98 := GenS4{F0: "\n\t", F1: 1}
                                                                            _ = v98
                                                                        } else {
                                                                            v99 := GenS4{F0: "foo bar", F1: v3}
                                                                            v29 = append(v29, 560)
                                                                            v29 = append(v29, genFn0())
                                                                            _ = v99
                                                                        }
                                                                        v44 += (v44 | v44)
                                                                        v100 := GenS8{F0: (fi40 % ((fi40) | 1)), F1: -1}
                                                                        _ = v100
                                                                    }
                                                                    v3 = v3
                                                                default:
                                                                    if (((-(v2)) < (v2 * v2)) || false) {
                                                                        v3 = (v3 + v3)
                                                                        if ("hello" <= "") {
                                                                            switch "" {
                                                                            case "a":
                                                                                v3 -= (v3 * (v3 ^ (v3 >> ((970) & 63))))
                                                                            }
                                                                            v101 := GenS9{F0: 1000, F1: (-(v3)), F2: 533}
                                                                            v102 := GenS1{F0: 42}
                                                                            switch (-(fi25)) {
                                                                            case 983:
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    var v103 bool = ((fi25 > fi40) || true)
                                                                                    _ = v103
                                                                                case "":
                                                                                    if (!false) {
                                                                                        fmt.Println(fi25, v1, v2)
                                                                                        var v104 uint16 = 32767
                                                                                        if (!true) {
                                                                                            var v105 float64 = v1
                                                                                            if len(v29) >= 2 {
                                                                                                v29 = v29[1:len(v29)]
                                                                                            }
                                                                                            _ = v105
                                                                                        }
                                                                                        _ = v104
                                                                                    } else {
                                                                                        if (true || true) {
                                                                                            v4 = 2
                                                                                        } else {
                                                                                            var v106 int64 = -75
                                                                                            v107 := GenS10{F0: 255, F1: 332}
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        }
                                                                                        v108 := GenS11{F0: v3}
                                                                                        fmt.Println(v3, fi40)
                                                                                        v44 += (v44 + genFn1())
                                                                                        _ = v108
                                                                                    }
                                                                                default:
                                                                                    if ((fi40 | fi25) <= (-(fi40))) {
                                                                                        v109 := GenS12{F0: 18446744073709551615}
                                                                                        v2 = v2
                                                                                        fmt.Println(v1, v4)
                                                                                        _ = len(v29)
                                                                                        _ = v109
                                                                                    }
                                                                                }
                                                                            case 32768:
                                                                                var v110 int64 = 0
                                                                                _ = v110
                                                                            default:
                                                                                v3 = (^(v3))
                                                                            }
                                                                            if len(v29) > 0 {
                                                                                v29[0] = (^(fi25))
                                                                            }
                                                                            v2 += (v2 + (v2 & v2))
                                                                            _ = v101
                                                                            _ = v102
                                                                        } else {
                                                                            var v111 uint8 = 2
                                                                            _ = v111
                                                                        }
                                                                        if len(v29) >= 2 {
                                                                            v29 = v29[1:len(v29)]
                                                                        }
                                                                    }
                                                                }
                                                                v112 := GenS13{F0: 12}
                                                                _ = v112
                                                            }
                                                            switch (fi40 << ((563) & 63)) {
                                                            case -1:
                                                                if len(v29) > 0 {
                                                                    v29[0] = (fi40 - fi25)
                                                                }
                                                            default:
                                                                fmt.Println(fi40, v4)
                                                            }
                                                            v3 = 42
                                                            v2 -= 71
                                                        }
                                                    }
                                                    if true {
                                                        v2 -= (v2 * v2)
                                                    }
                                                    var v113 int = fi40
                                                    _ = v113
                                                    _ = v44
                                                } else {
                                                    _ = len(v29)
                                                }
                                                if false {
                                                    if ("hello" != (("a" + "hello") + "\n\t")) {
                                                        v29 = append(v29, fi25)
                                                        if (30 == 0) {
                                                            var v114 int8 = v2
                                                            _ = v114
                                                        } else {
                                                            var v115 float64 = -2.25
                                                            _ = v115
                                                        }
                                                        var v116 float32 = (1.5 / 1.5)
                                                        _ = v116
                                                    } else {
                                                        if true {
                                                            fmt.Println(v1)
                                                        } else {
                                                            var v117 uint32 = 0
                                                            v118 := GenS14{F0: -74, F1: (v4 >> ((65535) & 63)), F2: 42}
                                                            var v119 int32 = v3
                                                            _ = v117
                                                            _ = v118
                                                            _ = v119
                                                        }
                                                    }
                                                    v2 += -128
                                                } else {
                                                    v3 = v3
                                                    v1 += v1
                                                    v1 = v1
                                                }
                                            } else {
                                                v29 = append(v29, (fi40 / ((4294967295) | 1)))
                                            }
                                            v120 := GenS15{F0: fi40, F1: "", F2: 1}
                                            _ = v120
                                        } else {
                                            v3 = v3
                                            v2 = (-(v2))
                                        }
                                        fmt.Println(fi40, fi25, v1)
                                        _ = fi40
                                    }
                                    var v121 uint16 = 187
                                    for fi122 := 0; fi122 < 2; fi122++ {
                                        switch (fi25 | fi25) {
                                        case 339:
                                            v123 := GenS1{F0: 2}
                                            _ = v123
                                        case 641:
                                            v2 = v2
                                        default:
                                            var v124 int16 = 1
                                            _ = v124
                                        }
                                        switch ("\n\t" + ((("hello" + "") + (("hello" + ((("foo bar" + "") + "foo bar") + "foo bar")) + ("\n\t" + "hello"))) + "a")) {
                                        case "a":
                                            v29 = append(v29, int(v2))
                                        }
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        _ = fi122
                                    }
                                    _ = v121
                                }
                            case "":
                                v4 = (^(v4))
                            }
                        case "hello":
                            v4 += v4
                        default:
                            fmt.Println(fi25, v3)
                        }
                    } else {
                        for fi125 := 0; fi125 < 10; fi125++ {
                            fmt.Println(v3)
                            v2 += (v2 + (v2 << ((735) & 63)))
                            _ = fi125
                        }
                        _ = len(v29)
                        var v126 bool = true
                        _ = v126
                    }
                    _ = fi25
                    _ = v29
                }
                var v127 uint32 = 4294967295
                v1 = ((v1 / 1.0) + float64(1.5))
                _ = v127
            } else {
                v2 += (v2 << ((838) & 63))
                v128 := GenS9{F0: -32768, F1: 42, F2: 32767}
                _ = v128
            }
            switch 127 {
            case -453:
                for fi129 := 0; fi129 < 19; fi129++ {
                    if false {
                        v130 := []int{456, (^(fi129))}
                        v1 = 3.14159
                        _ = v130
                    } else {
                        if false {
                            switch -870 {
                            case -2147483648:
                                v2 -= v2
                            case 127:
                                for fi131 := 0; fi131 < 17; fi131++ {
                                    var v132 int = (fi131 + (fi129 % ((fi129) | 1)))
                                    switch "foo bar" {
                                    case "\n\t":
                                        v133 := []int{(fi131 ^ 1), v132}
                                        _ = v133
                                    case "foo bar":
                                        var v134 int64 = -1
                                        _ = v134
                                    }
                                    v135 := []int{}
                                    var v136 int = (v132 / (((fi131 * (fi129 >> ((255) & 63)))) | 1))
                                    _ = fi131
                                    _ = v132
                                    _ = v135
                                    _ = v136
                                }
                            case -9223372036854775808:
                                if (68 >= (v2 % (((v2 + v2)) | 1))) {
                                    v137 := GenS16{F0: int32(fi129), F1: -32768, F2: v2}
                                    if (517 == 545) {
                                        if true {
                                            var v138 int8 = (v2 & (v2 + (^(v2))))
                                            var v139 float64 = 1e10
                                            v140 := GenS13{F0: 255}
                                            _ = v138
                                            _ = v139
                                            _ = v140
                                        } else {
                                            for fi141 := 0; fi141 < 20; fi141++ {
                                                if (("foo bar" + "hello") <= "a") {
                                                    var v142 uint64 = 2
                                                    if false {
                                                        v4 = (v4 >> ((2) & 63))
                                                        v3 += v3
                                                    } else {
                                                        fmt.Println(v3)
                                                    }
                                                    _ = v142
                                                } else {
                                                    v143 := GenS17{F0: false, F1: (fi129 << ((65535) & 63)), F2: (v3 >> ((930) & 63))}
                                                    _ = v143
                                                }
                                                var v144 int16 = 32767
                                                var v145 float32 = float32(v1)
                                                _ = fi141
                                                _ = v144
                                                _ = v145
                                            }
                                            v3 = v3
                                        }
                                    }
                                    var v146 uint32 = genFn1()
                                    _ = v137
                                    _ = v146
                                } else {
                                    v147 := GenS4{F0: "hello", F1: v3}
                                    _ = v147
                                }
                            }
                            var v148 uint64 = v4
                            v149 := GenS18{F0: v148, F1: 2, F2: "a"}
                            _ = v148
                            _ = v149
                        } else {
                            v150 := []int{int(v2), 0, fi129, (fi129 - (-(fi129)))}
                            v3 = v3
                            _ = v150
                        }
                        switch ("hello" + "a") {
                        case "foo bar":
                            v151 := GenS19{F0: -1, F1: false, F2: 185}
                            _ = v151
                        case "":
                            var v152 int64 = -656
                            _ = v152
                        case "hello":
                            for fi153 := 0; fi153 < 13; fi153++ {
                                v154 := []int{-1, genFn0(), fi129}
                                var v155 uint8 = 1
                                _ = fi153
                                _ = v154
                                _ = v155
                            }
                        default:
                            if true {
                                var v156 uint8 = 255
                                v157 := GenS10{F0: uint(v4), F1: (v4 >> ((828) & 63))}
                                _ = v156
                                _ = v157
                            }
                        }
                        if false {
                            switch fi129 {
                            case -284:
                                v158 := []int{fi129}
                                _ = v158
                            case -727:
                                var v159 float32 = 0.1
                                _ = v159
                            case 4294967295:
                                if (!(((3.14159 * (v1 * float64(float32(1e10)))) / (1.5 + (v1 * float64(0.0)))) < v1)) {
                                    var v160 float64 = 1e10
                                    v2 = (v2 * (v2 % ((v2) | 1)))
                                    var v161 uint = 1
                                    _ = v160
                                    _ = v161
                                }
                            default:
                                fmt.Println(v3, v3)
                            }
                            fmt.Println(v3, v3, v3)
                            for fi162 := 0; fi162 < 19; fi162++ {
                                v163 := GenS15{F0: genFn0(), F1: (("\n\t" + ("a" + "a")) + ("" + "\n\t")), F2: 65535}
                                _ = fi162
                                _ = v163
                            }
                        } else {
                            if true {
                                switch genFn0() {
                                case 9223372036854775807:
                                    if false {
                                        var v164 int16 = 32767
                                        v165 := GenS5{F0: -1}
                                        _ = v164
                                        _ = v165
                                    } else {
                                        var v166 int16 = -738
                                        v167 := []int{224, (^(fi129))}
                                        v168 := GenS3{F0: 412, F1: 853}
                                        _ = v166
                                        _ = v167
                                        _ = v168
                                    }
                                }
                                var v169 string = "\n\t"
                                _ = v169
                            }
                            fmt.Println(v4)
                        }
                    }
                    var v170 int8 = (v2 << ((1) & 63))
                    _ = fi129
                    _ = v170
                }
            case -9223372036854775808:
                v171 := []int{2147483647}
                _ = v171
            case 9223372036854775807:
                v172 := GenS12{F0: 65535}
                _ = v172
            default:
                switch -359 {
                case 32767:
                    if false {
                        switch -512 {
                        case 32767:
                            v1 = v1
                        case -2147483648:
                            var v173 float32 = 1e10
                            _ = v173
                        default:
                            var v174 int32 = (v3 + 0)
                            _ = v174
                        }
                    } else {
                        v175 := []int{4294967295, 128}
                        v3 -= (v3 << ((939) & 63))
                        v2 = v2
                        _ = v175
                    }
                case 0:
                    fmt.Println(v1)
                case -2147483648:
                    if false {
                        v3 = -2147483648
                        var v176 float32 = 0.1
                        _ = v176
                    }
                default:
                    for fi177 := 0; fi177 < 10; fi177++ {
                        if true {
                            v2 = v2
                            v178 := GenS20{F0: "hello", F1: v4}
                            v179 := GenS21{F0: 4294967295, F1: 2147483648, F2: -2147483648}
                            _ = v178
                            _ = v179
                        } else {
                            fmt.Println(v3, v2, v1)
                            if true {
                                var v180 uint32 = 4294967295
                                v1 -= v1
                                _ = v180
                            }
                            var v181 uint64 = 84
                            _ = v181
                        }
                        _ = fi177
                    }
                }
            }
        }
    }
    var v182 float32 = 1.5
    v183 := GenS9{F0: -1, F1: v3, F2: 657}
    v184 := []int{-165}
    v3 = 1
    if ((4294967295 <= 425) || true) {
        var v185 float64 = ((((v1 - v1) * v1) * v1) / 0.1)
        _ = v185
    }
    if len(v184) > 0 {
        v184[0] = 32768
    }
    v182 -= float32((((v1 - (1.5 - v1)) + (((v1 / v1) / (((0.1 * 3.14159) - ((0.1 / 1.0) + v1)) - v1)) + v1)) * (3.14159 * (float64(v182) + 3.14159))))
    switch "hello" {
    case "hello":
        var v186 int8 = (v2 * -38)
        _ = v186
    }
    var v187 string = ("\n\t" + "")
    v188 := GenS5{F0: -32768}
    v189 := GenS5{F0: 1}
    v190 := GenS22{F0: 988, F1: 255, F2: (v4 / ((0) | 1))}
    if true {
        v187 = v187
    } else {
        _ = len(v184)
        for fi191 := 0; fi191 < 6; fi191++ {
            if ((!true) || ((-32768 >= -32768) || (v1 < v1))) {
                var v192 uint32 = genFn1()
                if false {
                    v1 -= float64(v182)
                    v2 += (v2 & (v2 + v2))
                } else {
                    v4 = (v4 | 18446744073709551615)
                }
                _ = v192
            }
            _ = fi191
        }
        if false {
            v193 := GenS23{F0: 2}
            v187 = ((v187 + v187) + ((v187 + v187) + ("" + ((v187 + "") + "\n\t"))))
            v1 = ((1.5 + float64(v182)) / (v1 + v1))
            _ = v193
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v182:", v182)
    fmt.Println("slicelen v184:", len(v184), v184)
    fmt.Println("v187:", v187)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v182
    _ = v183
    _ = v184
    _ = v187
    _ = v188
    _ = v189
    _ = v190
    _ = v2
    _ = v3
    _ = v4
}
