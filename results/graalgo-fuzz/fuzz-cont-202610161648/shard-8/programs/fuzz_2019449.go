package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint; F1 bool; F2 int8 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int16; F1 int32 }
type GenS12 struct { F0 uint16; F1 string; F2 int32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int; F1 int16 }
type GenS16 struct { F0 int64; F1 int16; F2 uint32 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int64 }
func genFn0() int64 {
    if ("" <= (("foo bar" + ("hello" + "hello")) + "hello")) {
        var v4 int8 = -128
        fmt.Println(v4)
        _ = v4
    }
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v5, fi6)
        if true {
            v7 := []int{}
            _ = v7
        }
        if true {
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS1{F0: 2, F1: -1, F2: ("hello" + ((("a" + ("a" + ("" + "a"))) + "\n\t") + "foo bar"))}
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: 2, F1: 1000, F2: "hello"}
            v11 := []int{fi6, (fi6 | -2147483648), -1}
            _ = v10
            _ = v11
        } else {
            if false {
                for fi12 := 0; fi12 < 11; fi12++ {
                    if (!((((!(255 >= (v5 << ((65535) & 63)))) || false) && ((!false) || true)) || (42 == 2))) {
                        if true {
                            if (1 <= 690) {
                                var v13 float64 = float64(0.0)
                                switch "foo bar" {
                                case "a":
                                    v13 -= (3.14159 - v13)
                                case "\n\t":
                                    v14 := []int{fi12, fi12, fi6, fi12}
                                    _ = v14
                                case "foo bar":
                                    v15 := GenS2{F0: 0, F1: "\n\t"}
                                    _ = v15
                                default:
                                    var v16 int32 = 42
                                    _ = v16
                                }
                                v5 = v5
                                _ = v13
                            }
                        } else {
                            var v17 string = "hello"
                            v5 = (v5 << ((0) & 63))
                            var v18 float32 = 0.1
                            _ = v17
                            _ = v18
                        }
                        if (104 == -12) {
                            fmt.Println(v5, fi6)
                            var v19 float64 = float64(1.5)
                            fmt.Println(v19, fi6)
                            _ = v19
                        }
                        v5 += (v5 ^ (v5 >> ((1) & 63)))
                    } else {
                        var v20 string = "foo bar"
                        v20 = v20
                        _ = v20
                    }
                    var v21 uint8 = 195
                    v22 := []int{-9223372036854775808, fi12, (fi12 - (-(fi12))), fi12}
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = fi12
                    _ = v21
                    _ = v22
                }
                fmt.Println(fi6, fi6)
                v23 := GenS3{F0: -108}
                _ = v23
            } else {
                fmt.Println(v5, fi6)
            }
            v24 := GenS3{F0: 1}
            var v25 int = (fi6 & (fi6 * (fi6 ^ (fi6 | fi6))))
            _ = v24
            _ = v25
        }
        _ = fi6
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 3; fi27++ {
            if (false && (255 != 42)) {
                fmt.Println(v5)
                switch ("\n\t" + "foo bar") {
                case "a":
                    fmt.Println(fi26, fi26, fi26)
                case "hello":
                    if (164 > 0) {
                        switch 128 {
                        case 401:
                            var v28 string = ("" + "a")
                            _ = v28
                        case -2147483648:
                            v5 = (^(v5))
                        case 1:
                            v5 = 2
                        }
                        v29 := []int{fi27, fi26}
                        _ = v29
                    }
                }
                var v30 float32 = 0.0
                _ = v30
            } else {
                v31 := GenS1{F0: 298, F1: 319, F2: "\n\t"}
                _ = v31
            }
            _ = fi27
        }
        _ = fi26
    }
    v5 += v5
    _ = v5
    return 729
}
func genFn1(p0 uint, p1 uint8, p2 int8) int8 {
    p0 -= p0
    v32 := GenS4{F0: -545, F1: 572}
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return p2
}

func main() {
    var v1 uint = 632
    v2 := make(map[uint32]string)
    var v3 int8 = 100
    var v33 float64 = (((1e10 * 3.14159) + ((0.0 * 1.5) - -2.25)) + ((0.0 * 0.1) + (1e10 + float64(0.1))))
    fmt.Println(v3, v1)
    for fi34 := 0; fi34 < 7; fi34++ {
        if (0 == -1) {
            fmt.Println(fi34, v33)
            v35 := GenS5{F0: (true && (((406 != 384) && true) || ((!false) || false))), F1: v1}
            _ = v35
        }
        if false {
            v36 := GenS3{F0: v3}
            var v37 int16 = 669
            var v38 int64 = -9223372036854775808
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    var v39 bool = false
    var v40 int64 = -1
    v41 := GenS4{F0: 4294967295, F1: 1}
    for fi42 := 0; fi42 < 15; fi42++ {
        v43 := []int{}
        if ((true || true) || v39) {
            v44 := GenS1{F0: (v1 * (v1 << ((v1) & 63))), F1: 0, F2: ""}
            delete(v2, 42)
            _ = v44
        } else {
            var v45 int64 = genFn0()
            if len(v43) > 0 {
                v43[0] = (fi42 % ((int(v45)) | 1))
            }
            _ = v45
        }
        for fi46 := 0; fi46 < 2; fi46++ {
            if ((v1 & 977) < (v1 % (((v1 % (((v1 & 0)) | 1))) | 1))) {
                if ((^(v40)) == v40) {
                    v3 += 0
                    var v47 uint8 = 161
                    _ = v47
                } else {
                    v48 := GenS1{F0: (v1 << ((v1) & 63)), F1: 42, F2: ("\n\t" + (((("" + "") + "a") + ("foo bar" + "")) + "\n\t"))}
                    v33 = v33
                    var v49 string = "foo bar"
                    _ = v48
                    _ = v49
                }
                v40 -= v40
                switch (("\n\t" + "foo bar") + "a") {
                case "":
                    if false {
                        v50 := GenS6{F0: (v1 | 176), F1: (!false), F2: v3}
                        var v51 uint = (v1 << ((2) & 63))
                        _ = v50
                        _ = v51
                    } else {
                        var v52 int32 = -2147483648
                        fmt.Println(fi42, fi46)
                        v39 = v39
                        _ = v52
                    }
                case "hello":
                    switch 9223372036854775807 {
                    case -2147483648:
                        v40 += -9223372036854775808
                    default:
                        switch (fi42 >> ((302) & 63)) {
                        case 32768:
                            var v53 uint16 = 65535
                            _ = v53
                        default:
                            var v54 int16 = 42
                            _ = v54
                        }
                    }
                    var v55 uint64 = 754
                    _ = v55
                case "foo bar":
                    fmt.Println(v33)
                default:
                    v56 := GenS4{F0: (fi46 >> (((^(v1))) & 63)), F1: -2147483648}
                    _ = v56
                }
            }
            _ = fi46
        }
        _ = fi42
        _ = v43
    }
    v3 = (v3 % (((v3 - v3)) | 1))
    v57 := GenS4{F0: 42, F1: int32(v1)}
    var v58 bool = v39
    if v58 {
        switch "\n\t" {
        case "hello":
            var v59 float64 = v33
            _ = v59
        case "":
            switch ("a" + "a") {
            case "a":
                switch ("foo bar" + "\n\t") {
                case "hello":
                    v33 = (((float64((-2.25 - (1.5 - 1e10))) + (float64(float32((1e10 / v33))) + v33)) + v33) * v33)
                case "":
                    v33 = 0.0
                }
            }
        case "\n\t":
            fmt.Println(v33, v58, v58)
        default:
            v60 := []int{9223372036854775807, -968, -1, 32768}
            _ = v60
        }
        _ = len(v2)
        fmt.Println(v3, v39, v40)
    } else {
        v3 -= (v3 / (((v3 + (v3 + (v3 % ((1) | 1))))) | 1))
        for fi61 := 0; fi61 < 17; fi61++ {
            var v62 bool = (v39 || (584 != 1))
            delete(v2, 0)
            switch (fi61 | 967) {
            case -2147483648:
                if v62 {
                    v1 = 4294967295
                    var v63 int = 32767
                    _ = v63
                }
            case 32768:
                if (0 < 2147483647) {
                    if v39 {
                        var v64 float32 = 0.0
                        _ = v64
                    } else {
                        if v58 {
                            v3 -= v3
                        } else {
                            for fi65 := 0; fi65 < 4; fi65++ {
                                if v39 {
                                    if v58 {
                                        switch fi61 {
                                        case 9223372036854775807:
                                            if (0 != 119) {
                                                v39 = (genFn0() == -765)
                                            } else {
                                                v2[425] = "foo bar"
                                            }
                                        case 128:
                                            fmt.Println(v58, v33, v58)
                                        default:
                                            if v58 {
                                                if v39 {
                                                    switch fi65 {
                                                    case -9223372036854775808:
                                                        v39 = false
                                                    default:
                                                        switch (fi65 / ((fi61) | 1)) {
                                                        case 0:
                                                            v39 = v58
                                                            v66 := GenS3{F0: (^(v3))}
                                                            _ = v66
                                                        }
                                                    }
                                                    switch (fi65 ^ 32767) {
                                                    case 2147483647:
                                                        var v67 float32 = 1e10
                                                        _ = v67
                                                    case 32767:
                                                        var v68 int = fi65
                                                        _ = v68
                                                    case -9223372036854775808:
                                                        if (1000 > 0) {
                                                            fmt.Println(fi61)
                                                            var v69 int8 = (v3 % ((v3) | 1))
                                                            var v70 uint = 653
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                    }
                                                } else {
                                                    if (110 > -518) {
                                                        v71 := GenS7{F0: v1, F1: 1}
                                                        var v72 int = 1
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    fmt.Println(v62)
                                                    v58 = (42 != 18446744073709551615)
                                                }
                                                if v39 {
                                                    var v73 bool = false
                                                    fmt.Println(v3)
                                                    _ = v73
                                                }
                                                v1 += v1
                                            } else {
                                                if ((18446744073709551615 == 65535) && v62) {
                                                    if v58 {
                                                        if (((!true) || v62) && true) {
                                                            var v74 uint16 = 32767
                                                            v3 = v3
                                                            switch 42 {
                                                            case 127:
                                                                if v39 {
                                                                    var v75 bool = (float32((1e10 - v33)) != (0.0 * 0.0))
                                                                    if (552 >= 18446744073709551615) {
                                                                        var v76 bool = (!false)
                                                                        var v77 uint16 = 32767
                                                                        _ = v76
                                                                        _ = v77
                                                                    } else {
                                                                        var v78 float32 = float32(v33)
                                                                        v79 := []int{754, 42}
                                                                        switch (fi65 / (((fi65 >> ((v1) & 63))) | 1)) {
                                                                        case -312:
                                                                            fmt.Println(v58)
                                                                        case 32767:
                                                                            var v80 uint = (v1 >> ((v1) & 63))
                                                                            var v81 float64 = float64((v78 - float32((v33 - -2.25))))
                                                                            var v82 bool = v58
                                                                            _ = v80
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                    var v83 uint = 800
                                                                    _ = v75
                                                                    _ = v83
                                                                }
                                                                fmt.Println(fi61)
                                                                fmt.Println(fi61, v33)
                                                            case 9223372036854775807:
                                                                switch "a" {
                                                                case "foo bar":
                                                                    fmt.Println(fi65, v58, v62)
                                                                case "hello":
                                                                    v62 = ((1 > 67) || v62)
                                                                    var v84 uint8 = 42
                                                                    _ = v84
                                                                case "a":
                                                                    if v39 {
                                                                        v2[uint32(v74)] = ("\n\t" + (("foo bar" + ((("foo bar" + "a") + ("\n\t" + "")) + ("foo bar" + "a"))) + "foo bar"))
                                                                        v40 = (v40 % ((-906) | 1))
                                                                        v62 = v39
                                                                    }
                                                                default:
                                                                    v33 += 0.0
                                                                }
                                                            case -9223372036854775808:
                                                                v1 -= (v1 + v1)
                                                            default:
                                                                var v85 string = "a"
                                                                _ = v85
                                                            }
                                                            _ = v74
                                                        }
                                                        var v86 float64 = 0.1
                                                        v58 = true
                                                        _ = v86
                                                    } else {
                                                        switch fi65 {
                                                        case 127:
                                                            v62 = ("\n\t" < ("\n\t" + "\n\t"))
                                                        case 753:
                                                            var v87 int = -379
                                                            _ = v87
                                                        case 2147483647:
                                                            v1 = 255
                                                        default:
                                                            v88 := v2[42]
                                                            _ = v88
                                                        }
                                                    }
                                                    if (2147483647 <= 9223372036854775807) {
                                                        fmt.Println(fi65)
                                                    } else {
                                                        switch 32768 {
                                                        case 47:
                                                            v1 -= (v1 % ((255) | 1))
                                                        case 128:
                                                            switch 42 {
                                                            case 127:
                                                                v62 = ((!v58) || v58)
                                                            case 42:
                                                                var v89 float64 = (0.1 / (v33 / (float64(-2.25) + v33)))
                                                                _ = v89
                                                            case 2147483647:
                                                                var v90 uint = v1
                                                                _ = v90
                                                            default:
                                                                var v91 int8 = v3
                                                                _ = v91
                                                            }
                                                        case 32768:
                                                            if (!(2 == 2)) {
                                                                v40 -= 9223372036854775807
                                                                v39 = v39
                                                                var v92 uint8 = 0
                                                                v58 = true
                                                                v1 = v1
                                                                delete(v2, 2147483647)
                                                                fmt.Println(fi61)
                                                                _ = v92
                                                            }
                                                        default:
                                                            if v39 {
                                                                v3 -= (v3 & 100)
                                                                v1 = 165
                                                            }
                                                            var v93 float64 = v33
                                                            _ = v93
                                                        }
                                                        v33 = (3.14159 - v33)
                                                        switch -826 {
                                                        case -795:
                                                            if (!("hello" < "hello")) {
                                                                var v94 int16 = 12
                                                                v1 = (v1 * v1)
                                                                _ = v94
                                                            } else {
                                                                if true {
                                                                    v95 := []int{-1}
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v96 int32 = 42
                                                                        _ = v96
                                                                    case "a":
                                                                        var v97 int16 = -532
                                                                        _ = v97
                                                                    default:
                                                                        switch (fi65 | fi61) {
                                                                        case 4294967295:
                                                                            v62 = false
                                                                        case 127:
                                                                            if len(v95) >= 2 {
                                                                                v95 = v95[1:len(v95)]
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v95
                                                                }
                                                                if ((v40 == v40) && false) {
                                                                    var v98 bool = (1 <= 141)
                                                                    _ = v98
                                                                } else {
                                                                    if v39 {
                                                                        if (297 != 2) {
                                                                            fmt.Println(v1)
                                                                        } else {
                                                                            if ((v62 && (float32((v33 * (v33 + (v33 - (v33 / 0.1))))) <= float32(0.0))) && false) {
                                                                                if ((^(v40)) <= v40) {
                                                                                    v99 := GenS5{F0: false, F1: (^(v1))}
                                                                                    v33 += (v33 * v33)
                                                                                    _ = v99
                                                                                } else {
                                                                                    v39 = (true && ((!(42 != 964)) && true))
                                                                                    v58 = (v39 && v58)
                                                                                    v100 := GenS3{F0: v3}
                                                                                    _ = v100
                                                                                }
                                                                                var v101 uint = (^(v1))
                                                                                _ = v101
                                                                            } else {
                                                                                if (318 >= (v40 * (v40 >> (((v1 - v1)) & 63)))) {
                                                                                    var v102 uint64 = 42
                                                                                    v103 := GenS8{F0: int64(v3), F1: 2}
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                }
                                                                                var v104 string = "\n\t"
                                                                                var v105 int64 = (v40 % (((v40 / ((v40) | 1))) | 1))
                                                                                if v39 {
                                                                                    _ = len(v2)
                                                                                    switch 4294967295 {
                                                                                    case -9223372036854775808:
                                                                                        if v39 {
                                                                                            switch fi61 {
                                                                                            case 42:
                                                                                                var v106 uint32 = 130
                                                                                                var v107 uint32 = (v106 ^ (v106 * (v106 + 692)))
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            case 127:
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    v33 = v33
                                                                                                }
                                                                                            case 1:
                                                                                                if v58 {
                                                                                                    fmt.Println(v40, v105, v3)
                                                                                                    fmt.Println(v58, v105, v1)
                                                                                                    v108 := GenS9{F0: 32767, F1: genFn1(v1, 42, (v3 & v3))}
                                                                                                    _ = v108
                                                                                                } else {
                                                                                                    fmt.Println(v39)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case -278:
                                                                                        var v109 uint64 = 792
                                                                                        _ = v109
                                                                                    case -1:
                                                                                        if ((float32((v33 + 1.5)) <= (((1.5 * (0.1 + (0.0 - 1.5))) + float32(-2.25)) + (1e10 * (float32(v33) / ((float32(float64(3.14159)) - 1e10) * ((0.1 + (-2.25 + float32(v33))) / (float32((v33 - 1e10)) * 1.5))))))) && (183 >= 17)) {
                                                                                            var v110 uint16 = 776
                                                                                            _ = v110
                                                                                        } else {
                                                                                            var v111 int16 = -1
                                                                                            var v112 int64 = v40
                                                                                            if false {
                                                                                                var v113 uint64 = 514
                                                                                                _ = v113
                                                                                            }
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v1)
                                                                                        v3 += v3
                                                                                        fmt.Println(fi61)
                                                                                    }
                                                                                    v105 -= v105
                                                                                } else {
                                                                                    v114 := GenS5{F0: (!true), F1: 2}
                                                                                    v104 = "\n\t"
                                                                                    var v115 uint16 = 42
                                                                                    _ = v114
                                                                                    _ = v115
                                                                                }
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                            v1 += 18446744073709551615
                                                                            var v116 uint8 = 42
                                                                            var v117 uint32 = 830
                                                                            v118 := GenS8{F0: v40, F1: (v116 << (((v1 / ((v1) | 1))) & 63))}
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                        }
                                                                        switch fi61 {
                                                                        case -33:
                                                                            if (2 == v1) {
                                                                                v62 = (!(-2.25 != 1.5))
                                                                                switch 127 {
                                                                                case -109:
                                                                                    switch fi61 {
                                                                                    case 32767:
                                                                                        v62 = v58
                                                                                    }
                                                                                case -44:
                                                                                    _ = len(v2)
                                                                                case 1:
                                                                                    switch fi65 {
                                                                                    case 42:
                                                                                        fmt.Println(fi61)
                                                                                    default:
                                                                                        v39 = v58
                                                                                    }
                                                                                default:
                                                                                    _ = len(v2)
                                                                                }
                                                                            }
                                                                        case -9223372036854775808:
                                                                            if false {
                                                                                var v119 int32 = 0
                                                                                v58 = v39
                                                                                var v120 uint16 = uint16(v1)
                                                                                switch (fi65 + (^(fi65))) {
                                                                                case -37:
                                                                                    if v39 {
                                                                                        v1 -= (v1 << ((18446744073709551615) & 63))
                                                                                        fmt.Println(fi61, v33, v58)
                                                                                    }
                                                                                case 2147483647:
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        if v39 {
                                                                                            v58 = ((fi61 ^ fi65) > -552)
                                                                                        }
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            fmt.Println(v119)
                                                                                            switch (^(fi65)) {
                                                                                            case 2147483647:
                                                                                                var v121 uint = (v1 - v1)
                                                                                                var v122 uint64 = 4294967295
                                                                                                v33 -= 0.1
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            case -231:
                                                                                                v123 := []int{127, fi65, int(v120), (fi65 + fi65)}
                                                                                                _ = v123
                                                                                            case 642:
                                                                                                var v124 bool = v58
                                                                                                _ = v124
                                                                                            }
                                                                                        } else {
                                                                                            var v125 uint8 = 95
                                                                                            var v126 float32 = 0.0
                                                                                            v58 = false
                                                                                            _ = v125
                                                                                            _ = v126
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v119 = v119
                                                                                        v2[2147483647] = "foo bar"
                                                                                        var v127 float32 = -2.25
                                                                                        _ = v127
                                                                                    default:
                                                                                        fmt.Println(v119, v3, v62)
                                                                                        v128 := GenS4{F0: (fi61 << ((v1) & 63)), F1: v119}
                                                                                        _ = v128
                                                                                    }
                                                                                case 32768:
                                                                                    v129 := []int{}
                                                                                    _ = v129
                                                                                }
                                                                                _ = v119
                                                                                _ = v120
                                                                            } else {
                                                                                if v62 {
                                                                                    if (v58 && (float32((0.0 - ((1.5 * v33) * (0.0 + 1e10)))) != 3.14159)) {
                                                                                        fmt.Println(v39)
                                                                                        v130 := GenS1{F0: 1, F1: 0, F2: ("" + "hello")}
                                                                                        _ = v130
                                                                                    } else {
                                                                                        fmt.Println(v40, v33, v39)
                                                                                    }
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        var v131 uint32 = 0
                                                                                        _ = v131
                                                                                    }
                                                                                }
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v132 := GenS2{F0: 21, F1: ("a" + "")}
                                                                                    _ = v132
                                                                                }
                                                                                v133 := []int{(fi61 * (fi65 | fi65)), fi65}
                                                                                _ = v133
                                                                            }
                                                                        }
                                                                        fmt.Println(v39, fi61)
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(v33, v3, fi65)
                                                                            v134 := []int{(fi65 + int(v3))}
                                                                            _ = v134
                                                                        }
                                                                    }
                                                                }
                                                                var v135 int8 = (v3 - (v3 ^ (-(v3))))
                                                                _ = v135
                                                            }
                                                        case 9223372036854775807:
                                                            v3 = (v3 / ((genFn1((^(v1)), 42, -13)) | 1))
                                                        }
                                                        v62 = (true && (true || v58))
                                                    }
                                                    if true {
                                                        fmt.Println(v62, fi61, v62)
                                                        if v62 {
                                                            v3 = (-(v3))
                                                            v62 = (254 != 1)
                                                            v40 += (v40 - -1)
                                                        }
                                                        var v136 uint16 = 255
                                                        _ = v136
                                                    }
                                                } else {
                                                    switch fi65 {
                                                    case 9223372036854775807:
                                                        v3 -= (v3 & (v3 / ((100) | 1)))
                                                    default:
                                                        v3 -= v3
                                                    }
                                                    v137 := GenS10{F0: -643}
                                                    v1 = 65535
                                                    var v138 int8 = genFn1(686, 167, (v3 * (v3 - 100)))
                                                    if false {
                                                        v139 := GenS2{F0: (v138 | 127), F1: "\n\t"}
                                                        v140 := []int{(fi65 - fi65), -9223372036854775808}
                                                        if v62 {
                                                            v140 = append(v140, (fi61 | fi61))
                                                        }
                                                        _ = v139
                                                        _ = v140
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                    v62 = v62
                                                    var v141 float32 = 0.0
                                                    v142 := GenS11{F0: -32768, F1: int32(v138)}
                                                    v143 := v2[0]
                                                    _ = v137
                                                    _ = v138
                                                    _ = v141
                                                    _ = v142
                                                    _ = v143
                                                }
                                            }
                                        }
                                        v144 := GenS5{F0: false, F1: (v1 - (v1 % ((v1) | 1)))}
                                        if v62 {
                                            if ((float64(-2.25) / 1.5) >= (-2.25 / (v33 - (0.1 - 0.0)))) {
                                                v145 := GenS12{F0: 32767, F1: "\n\t", F2: 702}
                                                _ = v145
                                            } else {
                                                if (v40 < v40) {
                                                    v40 -= v40
                                                    var v146 int64 = (v40 | (v40 / ((v40) | 1)))
                                                    v62 = (fi61 == fi65)
                                                    v40 -= (-(v40))
                                                    _ = v146
                                                } else {
                                                    var v147 float64 = v33
                                                    fmt.Println(fi61, v33)
                                                    var v148 float32 = 0.1
                                                    _ = v147
                                                    _ = v148
                                                }
                                                v149 := []int{-958, 32768}
                                                if (!(!v58)) {
                                                    v2[65535] = "a"
                                                    v39 = v58
                                                    switch "\n\t" {
                                                    case "":
                                                        v2[347] = (("hello" + "\n\t") + "hello")
                                                    case "\n\t":
                                                        var v150 int16 = 1000
                                                        _ = v150
                                                    }
                                                } else {
                                                    var v151 uint32 = 0
                                                    v3 = (-(v3))
                                                    if v58 {
                                                        if true {
                                                            fmt.Println(v40)
                                                            fmt.Println(v151)
                                                        } else {
                                                            if ((0.1 < (1.5 * 0.1)) || v39) {
                                                                fmt.Println(v1, v1)
                                                                v152 := GenS13{F0: 417}
                                                                _ = v152
                                                            } else {
                                                                var v153 int16 = int16(fi65)
                                                                var v154 uint = v1
                                                                _ = v153
                                                                _ = v154
                                                            }
                                                            var v155 bool = (v33 != v33)
                                                            v62 = v155
                                                            _ = v155
                                                        }
                                                        v39 = (0.1 == v33)
                                                    }
                                                    _ = v151
                                                }
                                                _ = v149
                                            }
                                        } else {
                                            if ((((float32(v33) + -2.25) + -2.25) / 1e10) >= -2.25) {
                                                v1 = 4294967295
                                                var v156 bool = v39
                                                var v157 uint16 = 1
                                                _ = v156
                                                _ = v157
                                            } else {
                                                fmt.Println(v1, fi61)
                                            }
                                            v2[573] = "a"
                                            var v158 uint = v1
                                            _ = v158
                                        }
                                        _ = v144
                                    } else {
                                        switch (fi61 << (((^(v1))) & 63)) {
                                        case -316:
                                            if v62 {
                                                var v159 float32 = 0.0
                                                _ = v159
                                            } else {
                                                if v58 {
                                                    fmt.Println(v58)
                                                    v62 = v39
                                                    v160 := GenS14{F0: 607}
                                                    _ = v160
                                                } else {
                                                    v39 = (!false)
                                                    if true {
                                                        v62 = (("\n\t" != "a") || ((v40 + genFn0()) != v40))
                                                        v3 = (-(v3))
                                                        var v161 uint8 = 249
                                                        var v162 int16 = 0
                                                        _ = v161
                                                        _ = v162
                                                    }
                                                }
                                                var v163 uint64 = 42
                                                v164 := []int{(fi65 % ((2147483647) | 1)), (fi61 >> (((v1 + 1)) & 63))}
                                                _ = v163
                                                _ = v164
                                            }
                                        case 129:
                                            var v165 float32 = float32((v33 - (v33 * -2.25)))
                                            _ = v165
                                        case 0:
                                            v3 += 1
                                        }
                                    }
                                    v166 := GenS14{F0: 2}
                                    switch "\n\t" {
                                    case "hello":
                                        if v58 {
                                            v167 := v2[2]
                                            if v58 {
                                                fmt.Println(v167)
                                                if false {
                                                    switch (v167 + v167) {
                                                    case "hello":
                                                        var v168 uint16 = 0
                                                        _ = v168
                                                    case "a":
                                                        v40 = genFn0()
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(fi65, v39, fi65)
                                                    if v62 {
                                                        fmt.Println(v62, fi65, v40)
                                                        if (!v39) {
                                                            fmt.Println(fi65, v39, v62)
                                                        } else {
                                                            fmt.Println(v58, fi65)
                                                            if true {
                                                                v169 := []int{fi61, (^(fi65))}
                                                                v1 += v1
                                                                _ = v169
                                                            } else {
                                                                v3 += (v3 * 59)
                                                                v170 := []int{fi61, fi61, fi65, 465}
                                                                _ = v170
                                                            }
                                                        }
                                                    }
                                                    v58 = (!v58)
                                                }
                                            } else {
                                                var v171 uint = v1
                                                if false {
                                                    if (false && (255 != 65535)) {
                                                        _ = len(v2)
                                                    }
                                                    var v172 uint = 255
                                                    v173 := GenS15{F0: fi61, F1: -1}
                                                    v174 := GenS2{F0: (v3 | genFn1(v1, 127, genFn1((v172 ^ (v171 | (v1 ^ (v172 % ((v172) | 1))))), 127, (v3 + v3)))), F1: (v167 + v167)}
                                                    _ = len(v2)
                                                    switch (-(fi61)) {
                                                    case 128:
                                                        if ((!(v39 || (false && v39))) && v62) {
                                                            v175 := GenS3{F0: v3}
                                                            var v176 int = (fi61 << ((18446744073709551615) & 63))
                                                            if v62 {
                                                                v40 += (^(v40))
                                                                v39 = v58
                                                            }
                                                            _ = v175
                                                            _ = v176
                                                        } else {
                                                            v1 -= 4294967295
                                                        }
                                                    case -9223372036854775808:
                                                        v3 += (v3 | v3)
                                                    default:
                                                        v33 = float64(float32(v33))
                                                    }
                                                    v1 -= v1
                                                    _ = v172
                                                    _ = v173
                                                    _ = v174
                                                }
                                                fmt.Println(v1)
                                                _ = v171
                                            }
                                            _ = v167
                                        } else {
                                            switch ("" + "hello") {
                                            case "a":
                                                if v58 {
                                                    v177 := GenS16{F0: v40, F1: 32767, F2: 1}
                                                    fmt.Println(v3, v33, v40)
                                                    _ = v177
                                                } else {
                                                    v178 := GenS17{F0: ("foo bar" + "foo bar")}
                                                    v58 = v62
                                                    switch fi61 {
                                                    case -2147483648:
                                                        var v179 float32 = 0.1
                                                        _ = v179
                                                    }
                                                    _ = v178
                                                }
                                            case "\n\t":
                                                v39 = (!v62)
                                            case "hello":
                                                v180 := GenS15{F0: -669, F1: -918}
                                                _ = v180
                                            default:
                                                v181 := GenS18{F0: (v40 / ((v40) | 1))}
                                                _ = v181
                                            }
                                        }
                                    case "a":
                                        switch (fi61 / ((2147483647) | 1)) {
                                        case 32768:
                                            delete(v2, 2)
                                        case 128:
                                            switch "a" {
                                            case "\n\t":
                                                fmt.Println(v58, v39, fi65)
                                            case "":
                                                var v182 string = ""
                                                _ = v182
                                            default:
                                                v2[2] = "foo bar"
                                            }
                                        case 9223372036854775807:
                                            v33 -= -2.25
                                        default:
                                            _ = len(v2)
                                        }
                                    default:
                                        if (!(65535 != 2)) {
                                            switch fi61 {
                                            case -9223372036854775808:
                                                var v183 uint = v1
                                                _ = v183
                                            case 136:
                                                var v184 uint = v1
                                                _ = v184
                                            case -1000:
                                                var v185 uint = v1
                                                _ = v185
                                            default:
                                                var v186 uint16 = 255
                                                _ = v186
                                            }
                                            var v187 float32 = 3.14159
                                            if false {
                                                if v62 {
                                                    v188 := []int{127}
                                                    _ = v188
                                                } else {
                                                    _ = len(v2)
                                                    v40 -= -1
                                                    var v189 float64 = 0.1
                                                    _ = v189
                                                }
                                                switch int(v40) {
                                                case -1:
                                                    var v190 int8 = v3
                                                    _ = v190
                                                default:
                                                    var v191 int64 = 1
                                                    _ = v191
                                                }
                                                v192 := GenS9{F0: 0, F1: 89}
                                                _ = v192
                                            }
                                            _ = v187
                                        }
                                        v193 := GenS9{F0: 255, F1: (v3 * (-(v3)))}
                                        _ = v193
                                    }
                                    _ = v166
                                } else {
                                    v194 := []int{}
                                    _ = v194
                                }
                                delete(v2, 2)
                                var v195 uint = 4294967295
                                fmt.Println(fi61)
                                var v196 int64 = genFn0()
                                if (("" <= "foo bar") && (v3 != v3)) {
                                    v3 -= genFn1(2, 128, v3)
                                }
                                var v197 int8 = (v3 & genFn1(v195, 255, v3))
                                fmt.Println(v1, v197)
                                _ = fi65
                                _ = v195
                                _ = v196
                                _ = v197
                            }
                        }
                    }
                } else {
                    var v198 uint8 = uint8(v1)
                    _ = v198
                }
            case 2147483647:
                v62 = v58
            default:
                v199 := GenS10{F0: -1}
                _ = v199
            }
            _ = fi61
            _ = v62
        }
        for fi200 := 0; fi200 < 6; fi200++ {
            var v201 int = 127
            v33 -= ((v33 / 1e10) + 1e10)
            _ = fi200
            _ = v201
        }
    }
    v202 := GenS3{F0: 42}
    var v203 int = 432
    v3 += (v3 >> (((v1 ^ 4294967295)) & 63))
    var v204 uint16 = 678
    fmt.Println(v203, v203)
    v204 = (v204 % ((v204) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v203:", v203)
    fmt.Println("v204:", v204)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v202
    _ = v203
    _ = v204
    _ = v3
    _ = v33
    _ = v39
    _ = v40
    _ = v41
    _ = v57
    _ = v58
}
