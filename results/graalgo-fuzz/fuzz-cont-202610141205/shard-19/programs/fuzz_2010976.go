package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 uint }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint16; F1 int64; F2 int32 }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 int32; F1 int64; F2 uint32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string; F1 int64; F2 uint8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int32; F1 int16; F2 int32 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int64; F1 int; F2 uint64 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint; F1 uint32; F2 uint }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int64; F1 int }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16, p1 int32) int64 {
    var v6 uint64 = 18446744073709551615
    fmt.Println(p0, p1)
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v6 = (v6 + (^(v6)))
            if true {
                v6 += v6
            }
            p1 -= -461
            var v9 float64 = 1.5
            _ = fi8
            _ = v9
        }
        var v10 uint16 = 228
        _ = fi7
        _ = v10
    }
    if false {
        v11 := []int{4294967295, 32768, -9223372036854775808, -1}
        if false {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 uint32 = 421
                v13 = (v13 | 1)
                _ = fi12
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 uint = 255
                fmt.Println(v15, fi14, fi14)
                _ = fi14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint32 = 42
            _ = fi16
            _ = v17
        }
        _ = v11
    } else {
        for fi18 := 0; fi18 < 16; fi18++ {
            p1 = p1
            _ = fi18
        }
        fmt.Println(p0)
    }
    fmt.Println(v6, p1, v6)
    _ = p0
    _ = p1
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = 100
    var v3 float32 = 0.0
    v4 := []int{4294967295, 492}
    v5 := make(map[uint32]string)
    if (!false) {
        var v19 float32 = v3
        _ = v19
    } else {
        fmt.Println(v2)
        v20 := GenS1{F0: 432, F1: (v1 * (v1 ^ v1))}
        _ = v20
    }
    var v21 uint = 1
    var v22 int8 = (v2 % ((1) | 1))
    var v23 float32 = v3
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS2{F0: ("\n\t" + "\n\t")}
        if false {
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v26 bool = false
                _ = v26
            }
            for fi27 := 0; fi27 < 5; fi27++ {
                switch "\n\t" {
                case "foo bar":
                    if false {
                        var v28 uint8 = 255
                        _ = v28
                    } else {
                        if (2 != 961) {
                            var v29 uint32 = 4294967295
                            if true {
                                v29 -= (v29 >> ((v21) & 63))
                                v30 := GenS3{F0: 255}
                                v21 = (^(v21))
                                if false {
                                    _ = len(v5)
                                    var v31 int32 = 42
                                    fmt.Println(v22, v31)
                                    _ = v31
                                }
                                _ = len(v4)
                                fmt.Println(fi24, v3)
                                _ = v30
                            }
                            var v32 float64 = 3.14159
                            _ = v29
                            _ = v32
                        } else {
                            var v33 float64 = 0.1
                            switch (fi24 ^ -356) {
                            case 2147483647:
                                var v34 string = ""
                                _ = v34
                            case 127:
                                if (0 == -377) {
                                    var v35 bool = false
                                    var v36 string = "foo bar"
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v1 = v1
                                    v37 := GenS4{F0: "hello", F1: 1, F2: v21}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    switch fi27 {
                                    case 42:
                                        var v38 int = (fi24 ^ (fi27 * (fi24 + (-(fi24)))))
                                        _ = v38
                                    case 4294967295:
                                        var v39 bool = true
                                        _ = v39
                                    default:
                                        fmt.Println(v23, v21)
                                    }
                                    fmt.Println(v23, v21, fi24)
                                }
                            }
                            fmt.Println(v3, fi24, v33)
                            v40 := GenS3{F0: 827}
                            if true {
                                _ = len(v5)
                            } else {
                                switch -2147483648 {
                                case -2147483648:
                                    v41 := GenS5{F0: ("foo bar" + ("hello" + (("\n\t" + "a") + "a"))), F1: uint32(v1)}
                                    _ = v41
                                case 42:
                                    v42 := v5[0]
                                    _ = v42
                                case 32767:
                                    var v43 uint32 = 1
                                    _ = v43
                                default:
                                    var v44 uint64 = 1
                                    _ = v44
                                }
                                if (0 > -1) {
                                    if ((^(v21)) >= v21) {
                                        if true {
                                            v45 := GenS2{F0: ("hello" + "a")}
                                            _ = v45
                                        } else {
                                            v46 := GenS6{F0: false, F1: 1}
                                            v47 := GenS7{F0: -9223372036854775808, F1: "\n\t"}
                                            _ = v46
                                            _ = v47
                                        }
                                        v48 := v5[475]
                                        _ = v48
                                    }
                                    if (true || (42 == 300)) {
                                        if (!(42 >= 0)) {
                                            v5[2147483647] = "hello"
                                            if false {
                                                switch fi24 {
                                                case 42:
                                                    var v49 bool = (fi24 > -2147483648)
                                                    if true {
                                                        switch fi24 {
                                                        case -1:
                                                            var v50 string = ""
                                                            _ = v50
                                                        case -2147483648:
                                                            v23 = v23
                                                        default:
                                                            v3 += v23
                                                        }
                                                        v1 = (v1 - (^(v1)))
                                                        var v51 int = (fi24 >> ((937) & 63))
                                                        _ = v51
                                                    }
                                                    fmt.Println(v49)
                                                    _ = v49
                                                case -9223372036854775808:
                                                    if len(v4) > 0 {
                                                        v4[0] = fi27
                                                    }
                                                case 675:
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v52 int16 = int16(v21)
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v53 float32 = v23
                                                            _ = v53
                                                        case "":
                                                            v54 := GenS7{F0: genFn0((v52 >> ((646) & 63)), -2147483648), F1: (("a" + "foo bar") + "a")}
                                                            _ = v54
                                                        case "foo bar":
                                                            _ = len(v5)
                                                        default:
                                                            v23 = 0.1
                                                        }
                                                        v1 += 409
                                                        _ = v52
                                                    case "hello":
                                                        v55 := GenS8{F0: 42, F1: genFn0(704, 42), F2: 0}
                                                        _ = v55
                                                    default:
                                                        if true {
                                                            fmt.Println(fi24)
                                                            if true {
                                                                if len(v4) > 0 {
                                                                    v4[0] = fi27
                                                                }
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                if true {
                                                                    if (!false) {
                                                                        if true {
                                                                            delete(v5, 328)
                                                                        }
                                                                        fmt.Println(v21, fi27)
                                                                        var v56 int8 = (v2 / ((100) | 1))
                                                                        v3 += (1e10 / v23)
                                                                        var v57 uint32 = 501
                                                                        _ = v56
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 float32 = v3
                                                                        _ = v58
                                                                    }
                                                                } else {
                                                                    v59 := GenS7{F0: -711, F1: "foo bar"}
                                                                    if (!(!("hello" < ("foo bar" + (("\n\t" + "a") + ""))))) {
                                                                        var v60 uint32 = 319
                                                                        if false {
                                                                            switch ((("hello" + ((("a" + "") + "foo bar") + "foo bar")) + "") + (("foo bar" + "a") + "hello")) {
                                                                            case "":
                                                                                v61 := GenS4{F0: ("hello" + "foo bar"), F1: v60, F2: 513}
                                                                                _ = v61
                                                                            case "\n\t":
                                                                                v62 := GenS1{F0: 65535, F1: (v1 % ((v1) | 1))}
                                                                                _ = v62
                                                                            case "a":
                                                                                if false {
                                                                                    var v63 uint = (^(v21))
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                            if ((true || (!(!((!((!(!true)) || true)) && true)))) && ((!(787 > v21)) && true)) {
                                                                                fmt.Println(v1, v3)
                                                                            }
                                                                            fmt.Println(v3, v60)
                                                                        }
                                                                        _ = v60
                                                                    }
                                                                    v64 := GenS5{F0: "foo bar", F1: 821}
                                                                    _ = v59
                                                                    _ = v64
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v2, v22)
                                                        }
                                                    }
                                                default:
                                                    var v65 int32 = -153
                                                    _ = v65
                                                }
                                                switch fi27 {
                                                case 128:
                                                    if (652 < (fi27 | (fi24 + fi27))) {
                                                        if (235 <= 2147483647) {
                                                            var v66 int = fi27
                                                            v67 := GenS8{F0: 123, F1: 1, F2: 188}
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        switch 42 {
                                                        case 9223372036854775807:
                                                            v5[4294967295] = "foo bar"
                                                        case -9223372036854775808:
                                                            var v68 int64 = -1
                                                            _ = v68
                                                        case 43:
                                                            v5[1] = ("hello" + ("" + "\n\t"))
                                                        }
                                                    } else {
                                                        v69 := GenS9{F0: 157, F1: "a"}
                                                        var v70 int = (-(fi24))
                                                        switch "\n\t" {
                                                        case "a":
                                                            v70 -= fi27
                                                            v71 := GenS10{F0: -500, F1: genFn0(-541, 42), F2: 28}
                                                            v2 += (v2 + (v22 >> (((v21 % ((18446744073709551615) | 1))) & 63)))
                                                            _ = v71
                                                        }
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                case 42:
                                                    fmt.Println(v22, v1, v22)
                                                }
                                            } else {
                                                fmt.Println(v1)
                                                if true {
                                                    v2 -= v2
                                                    if true {
                                                        _ = len(v5)
                                                    }
                                                }
                                                var v72 uint = (v21 * 255)
                                                v4 = append(v4, fi27)
                                                _ = len(v5)
                                                var v73 int16 = 902
                                                v74 := GenS11{F0: (v21 << ((v21) & 63))}
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            }
                                            var v75 int64 = 0
                                            v5[283] = "\n\t"
                                            _ = v75
                                        } else {
                                            v2 = -128
                                            if (0 != 1) {
                                                if ((0.1 + (-2.25 * (3.14159 / ((-2.25 / 3.14159) / (v33 - 0.0))))) != (((v33 - (v33 / 1.0)) - v33) / 1e10)) {
                                                    if (!false) {
                                                        fmt.Println(v23, v23, v33)
                                                        v22 = (^(v22))
                                                        v2 -= (v2 + 1)
                                                    } else {
                                                        var v76 int16 = 32767
                                                        if (1 >= 171) {
                                                            if false {
                                                                v77 := GenS3{F0: 2}
                                                                v78 := GenS3{F0: 65535}
                                                                _ = v77
                                                                _ = v78
                                                            } else {
                                                                v79 := v5[1]
                                                                switch 0 {
                                                                case -2147483648:
                                                                    var v80 int16 = v76
                                                                    _ = v80
                                                                default:
                                                                    _ = len(v4)
                                                                }
                                                                delete(v5, 944)
                                                                delete(v5, 65535)
                                                                var v81 int = 127
                                                                _ = v79
                                                                _ = v81
                                                            }
                                                        }
                                                        var v82 int32 = 2147483647
                                                        _ = v76
                                                        _ = v82
                                                    }
                                                    v21 = (v21 >> ((18446744073709551615) & 63))
                                                    var v83 uint64 = 0
                                                    _ = v83
                                                }
                                            }
                                        }
                                        v84 := GenS12{F0: ("" + (("" + (((((("a" + "foo bar") + ((("hello" + "") + ("foo bar" + "\n\t")) + ("hello" + "hello"))) + "foo bar") + ("a" + "foo bar")) + "a") + "a")) + (((("" + "a") + "foo bar") + "a") + (("a" + "foo bar") + "a")))), F1: 1, F2: 0}
                                        _ = v84
                                    } else {
                                        if (!true) {
                                            v23 = ((v3 * (3.14159 + float32((v33 / v33)))) * (v23 + v3))
                                        }
                                        _ = len(v4)
                                    }
                                    if ((v1 == 59) && (v22 >= v2)) {
                                        v33 = 0.1
                                    }
                                    v23 = (v3 - (v23 / 0.1))
                                    fmt.Println(v33, v22)
                                } else {
                                    v85 := GenS13{F0: v1}
                                    fmt.Println(v23)
                                    _ = v85
                                }
                            }
                            v86 := v5[2147483647]
                            _ = v33
                            _ = v40
                            _ = v86
                        }
                        fmt.Println(v3, v2, v1)
                        v87 := GenS14{F0: 1}
                        _ = v87
                    }
                }
                if (0 >= 42) {
                    v88 := GenS1{F0: 255, F1: (v1 % ((v1) | 1))}
                    v22 += -84
                    _ = v88
                } else {
                    var v89 int64 = 101
                    fmt.Println(v23, v2, v21)
                    _ = v89
                }
                _ = fi27
            }
        } else {
            if (true && false) {
                v2 = (v2 + 0)
            }
        }
        if true {
            var v90 string = (("a" + ((("a" + "\n\t") + "hello") + "hello")) + ((("\n\t" + ("hello" + "")) + "hello") + "hello"))
            v91 := GenS15{F0: -2147483648, F1: 54, F2: -8}
            _ = v90
            _ = v91
        }
        _ = fi24
        _ = v25
    }
    v1 += (v1 >> (((v21 & 735)) & 63))
    switch 0 {
    case 2147483647:
        if ((^(v1)) <= 389) {
            v1 = 42
        } else {
            switch 704 {
            case -1:
                v92 := GenS16{F0: 65535}
                _ = v92
            }
        }
    default:
        var v93 int16 = -32768
        _ = v93
    }
    v3 += (((1e10 + (float32(0.0) / ((v3 / v3) - float32(0.1)))) * (v23 * v3)) * (((v23 / -2.25) - 3.14159) / (v3 - (v23 - (v3 - -2.25)))))
    for fi94 := 0; fi94 < 19; fi94++ {
        v22 = int8(v21)
        v95 := GenS8{F0: 0, F1: 9223372036854775807, F2: -2147483648}
        v96 := v5[914]
        _ = fi94
        _ = v95
        _ = v96
    }
    switch "a" {
    case "\n\t":
        if true {
            v21 = v21
            v22 = 1
        } else {
            var v97 float64 = 0.0
            var v98 string = ""
            fmt.Println(v2)
            _ = v97
            _ = v98
        }
    case "a":
        var v99 bool = ("a" < "")
        _ = v99
    case "hello":
        v1 = (^(v1))
    }
    for fi100 := 0; fi100 < 10; fi100++ {
        switch "a" {
        case "a":
            if ((v21 ^ v21) != v21) {
                for fi101 := 0; fi101 < 2; fi101++ {
                    var v102 uint64 = v1
                    v21 = 4294967295
                    _ = fi101
                    _ = v102
                }
                v23 += 1.5
                for fi103 := 0; fi103 < 10; fi103++ {
                    var v104 int = (fi100 + (-(fi100)))
                    _ = len(v5)
                    if (-1 > 1) {
                        v23 += v23
                        var v105 int32 = -102
                        _ = v105
                    }
                    _ = fi103
                    _ = v104
                }
            } else {
                switch "\n\t" {
                case "":
                    for fi106 := 0; fi106 < 18; fi106++ {
                        if (false || true) {
                            var v107 bool = (((((v3 + v3) / float32(float64(v3))) * (float32((0.0 * -2.25)) * float32(3.14159))) * float32(float64(1.5))) < (-2.25 - 3.14159))
                            v108 := GenS17{F0: (("hello" + "") + "foo bar")}
                            _ = v107
                            _ = v108
                        } else {
                            switch (fi106 * fi100) {
                            case 899:
                                delete(v5, 4294967295)
                            case -285:
                                v3 -= (0.1 / 3.14159)
                            case 4294967295:
                                if (v21 == (v21 ^ v21)) {
                                    v4 = append(v4, fi106)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                        }
                        _ = fi106
                    }
                case "foo bar":
                    if (!(850 >= 9223372036854775807)) {
                        switch fi100 {
                        case -9223372036854775808:
                            switch ("foo bar" + "foo bar") {
                            case "":
                                var v109 int64 = 9223372036854775807
                                _ = v109
                            case "a":
                                if false {
                                    v22 = v2
                                } else {
                                    var v110 int = fi100
                                    for fi111 := 0; fi111 < 8; fi111++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            if false {
                                                var v112 uint64 = (v1 & 115)
                                                fmt.Println(v110)
                                                _ = v112
                                            }
                                        case "hello":
                                            switch "" {
                                            case "a":
                                                fmt.Println(v2, fi100, v3)
                                            case "":
                                                switch "a" {
                                                case "":
                                                    v21 += v21
                                                case "foo bar":
                                                    var v113 int32 = -1
                                                    _ = v113
                                                case "hello":
                                                    if false {
                                                        fmt.Println(fi111)
                                                        v21 = v21
                                                    }
                                                default:
                                                    v23 = (v23 + (v23 * ((0.1 / ((((v23 / v3) + 1e10) + (v23 * (v3 * (((float32(1e10) / 1e10) * v23) / float32(float64(1e10)))))) * 0.1)) + v23)))
                                                }
                                            default:
                                                var v114 bool = true
                                                _ = v114
                                            }
                                        case "\n\t":
                                            var v115 bool = (v23 != (v3 / (3.14159 / v23)))
                                            _ = v115
                                        default:
                                            v23 = v3
                                        }
                                        fmt.Println(v3)
                                        _ = fi111
                                    }
                                    v116 := GenS7{F0: 9223372036854775807, F1: ((("foo bar" + "hello") + "\n\t") + "hello")}
                                    _ = v110
                                    _ = v116
                                }
                            default:
                                for fi117 := 0; fi117 < 14; fi117++ {
                                    v5[2147483647] = ""
                                    v2 -= 42
                                    _ = fi117
                                }
                            }
                        case 32768:
                            var v118 uint = v21
                            _ = v118
                        case -170:
                            fmt.Println(v21, v2)
                        }
                        v21 -= (^(v21))
                    }
                }
                v2 += v2
                switch (-(fi100)) {
                case 42:
                    v21 = 449
                case -1:
                    if true {
                        v5[42] = "\n\t"
                        fmt.Println(v2)
                    } else {
                        switch "a" {
                        case "a":
                            if ((^(v21)) < (v21 % ((v21) | 1))) {
                                for fi119 := 0; fi119 < 14; fi119++ {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v120 float32 = v23
                                    var v121 uint16 = 65535
                                    switch (fi100 - 128) {
                                    case -1:
                                        v122 := GenS5{F0: "foo bar", F1: 0}
                                        _ = v122
                                    }
                                    _ = fi119
                                    _ = v120
                                    _ = v121
                                }
                                var v123 string = ("hello" + "\n\t")
                                _ = v123
                            }
                        case "hello":
                            for fi124 := 0; fi124 < 3; fi124++ {
                                _ = len(v5)
                                _ = fi124
                            }
                        case "foo bar":
                            for fi125 := 0; fi125 < 1; fi125++ {
                                if true {
                                    if (32767 < -290) {
                                        v21 = 611
                                        v126 := GenS8{F0: 1, F1: 298, F2: 975}
                                        var v127 bool = false
                                        _ = v126
                                        _ = v127
                                    } else {
                                        v21 = (^(v21))
                                        v3 += (0.1 * v23)
                                        v23 += v23
                                    }
                                    v3 -= 0.0
                                } else {
                                    v128 := GenS18{F0: genFn0(-1, 2147483647), F1: fi100, F2: (v1 + 583)}
                                    if (false || false) {
                                        v2 -= (v22 - v2)
                                        v129 := GenS14{F0: 1000}
                                        _ = v129
                                    }
                                    if (v23 >= (3.14159 - 1.5)) {
                                        if (0.0 < (0.0 * 0.1)) {
                                            fmt.Println(fi100, v1)
                                            fmt.Println(v23)
                                            v130 := GenS19{F0: 255}
                                            _ = v130
                                        }
                                        fmt.Println(v21)
                                        if false {
                                            v21 += v21
                                        }
                                        var v131 int16 = -32768
                                        _ = v131
                                    }
                                    _ = v128
                                }
                                _ = fi125
                            }
                        }
                        v22 = (v2 | 34)
                    }
                case 2147483647:
                    for fi132 := 0; fi132 < 15; fi132++ {
                        var v133 float64 = ((3.14159 * ((0.1 / float64(-2.25)) * -2.25)) / float64(v3))
                        fmt.Println(fi132)
                        _ = fi132
                        _ = v133
                    }
                }
            }
        }
        switch ("a" + "foo bar") {
        case "":
            for fi134 := 0; fi134 < 15; fi134++ {
                switch -532 {
                case 1:
                    var v135 uint8 = 2
                    _ = v135
                case -1:
                    if true {
                        v5[806] = (("\n\t" + "hello") + "")
                        switch ("a" + ("hello" + "a")) {
                        case "a":
                            var v136 float64 = ((3.14159 / (3.14159 * -2.25)) + float64((0.0 - v23)))
                            _ = v136
                        default:
                            _ = len(v5)
                        }
                    } else {
                        var v137 int16 = 42
                        _ = v137
                    }
                default:
                    if true {
                        v138 := v5[1]
                        switch fi134 {
                        case -760:
                            switch 2147483647 {
                            case 4294967295:
                                var v139 uint64 = v1
                                _ = v139
                            case 57:
                                if (uint16(v21) != 144) {
                                    var v140 int64 = -125
                                    if false {
                                        v4 = append(v4, fi134)
                                    }
                                    _ = len(v5)
                                    _ = v140
                                }
                            case 32768:
                                var v141 float32 = float32(float64((1e10 + v3)))
                                _ = v141
                            default:
                                var v142 bool = true
                                _ = v142
                            }
                        }
                        _ = v138
                    } else {
                        var v143 uint16 = 32767
                        v144 := v5[2]
                        fmt.Println(v144, fi134)
                        _ = v143
                        _ = v144
                    }
                }
                var v145 int8 = 67
                switch "a" {
                case "hello":
                    v23 = (v23 * ((v3 / v3) + ((((v3 - float32(((-2.25 + 0.0) * (0.0 / 3.14159)))) / v3) + v3) * v3)))
                case "":
                    if false {
                        var v146 int16 = 1000
                        switch 0 {
                        case 913:
                            v146 = 0
                        case 42:
                            var v147 int32 = -2147483648
                            _ = v147
                        case 833:
                            v148 := GenS14{F0: v146}
                            _ = v148
                        default:
                            switch ("a" + "foo bar") {
                            case "a":
                                var v149 int16 = v146
                                _ = v149
                            case "foo bar":
                                _ = len(v4)
                            }
                        }
                        switch "" {
                        case "foo bar":
                            var v150 uint = (v21 & 2)
                            _ = v150
                        case "\n\t":
                            _ = len(v5)
                        case "":
                            switch ("hello" + "foo bar") {
                            case "a":
                                v21 -= 255
                            case "":
                                var v151 uint16 = 1
                                _ = v151
                            default:
                                v146 = 1000
                            }
                        }
                        _ = v146
                    }
                default:
                    fmt.Println(fi134, v145, v3)
                }
                v152 := GenS20{F0: (^(v21)), F1: 867, F2: (v21 & (^(v21)))}
                if len(v4) > 0 {
                    v4[0] = 32768
                }
                if (v145 != (v145 / ((v22) | 1))) {
                    if true {
                        var v153 string = ("a" + "foo bar")
                        switch fi134 {
                        case -649:
                            switch v153 {
                            case "\n\t":
                                var v154 uint32 = 2
                                _ = v154
                            case "hello":
                                var v155 bool = (v21 == 255)
                                _ = v155
                            }
                        case 895:
                            if (-1 < -2147483648) {
                                var v156 uint16 = 2
                                _ = v156
                            }
                        case 4294967295:
                            switch 72 {
                            case 751:
                                var v157 uint32 = 627
                                _ = v157
                            case 2147483647:
                                if ((v145 / ((127) | 1)) != v145) {
                                    v2 = (-(v145))
                                } else {
                                    var v158 float64 = 0.0
                                    _ = v158
                                }
                            default:
                                _ = len(v5)
                            }
                            var v159 uint64 = v1
                            var v160 float64 = 1e10
                            _ = v159
                            _ = v160
                        }
                        _ = v153
                    } else {
                        delete(v5, 2147483648)
                    }
                    v161 := GenS21{F0: (v2 & 67)}
                    fmt.Println(fi134, v1, v2)
                    _ = v161
                }
                switch "" {
                case "foo bar":
                    _ = len(v5)
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "hello":
                    fmt.Println(v2, v2, v3)
                }
                _ = fi134
                _ = v145
                _ = v152
            }
        default:
            v162 := GenS19{F0: 203}
            _ = v162
        }
        v1 -= v1
        _ = fi100
    }
    if true {
        for fi163 := 0; fi163 < 18; fi163++ {
            v164 := v5[2]
            _ = fi163
            _ = v164
        }
        for fi165 := 0; fi165 < 10; fi165++ {
            var v166 bool = false
            if (((849 < 2) || ((-(v2)) < (v22 ^ 68))) && v166) {
                switch "" {
                case "foo bar":
                    if (1e10 <= 0.0) {
                        for fi167 := 0; fi167 < 13; fi167++ {
                            var v168 uint16 = 32767
                            v169 := GenS7{F0: genFn0(32767, 505), F1: "a"}
                            v170 := GenS22{F0: 386, F1: 42}
                            _ = fi167
                            _ = v168
                            _ = v169
                            _ = v170
                        }
                    }
                case "":
                    if v166 {
                        if true {
                            var v171 uint8 = 109
                            var v172 bool = ((0.0 <= float64(((v3 * v23) - 0.0))) && v166)
                            _ = v171
                            _ = v172
                        } else {
                            delete(v5, 0)
                            v173 := GenS23{F0: (v21 >> ((764) & 63))}
                            for fi174 := 0; fi174 < 6; fi174++ {
                                switch fi165 {
                                case 32767:
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v23)
                                    case "hello":
                                        v1 = v1
                                    case "\n\t":
                                        switch fi174 {
                                        case 312:
                                            switch "" {
                                            case "":
                                                var v175 uint32 = 2147483648
                                                _ = v175
                                            case "\n\t":
                                                v176 := GenS6{F0: v166, F1: genFn0(1000, 1)}
                                                _ = v176
                                            case "foo bar":
                                                delete(v5, 1)
                                            default:
                                                var v177 string = "foo bar"
                                                _ = v177
                                            }
                                        case 1:
                                            v21 -= (v21 | v21)
                                        case 191:
                                            fmt.Println(v22, fi165, v22)
                                        }
                                    default:
                                        switch "a" {
                                        case "":
                                            var v178 int32 = 127
                                            _ = v178
                                        case "hello":
                                            v179 := GenS14{F0: 1000}
                                            _ = v179
                                        }
                                        v180 := GenS2{F0: "foo bar"}
                                        var v181 uint8 = 1
                                        _ = v180
                                        _ = v181
                                    }
                                }
                                v182 := v5[2147483647]
                                _ = fi174
                                _ = v182
                            }
                            _ = v173
                        }
                    }
                }
                for fi183 := 0; fi183 < 7; fi183++ {
                    fmt.Println(v2, v22, v2)
                    if (!v166) {
                        v21 = (v21 ^ v21)
                        v21 = 18446744073709551615
                    }
                    _ = fi183
                }
            } else {
                var v184 int8 = 1
                v3 = float32(-2.25)
                var v185 int8 = 1
                _ = v184
                _ = v185
            }
            _ = fi165
            _ = v166
        }
    } else {
        fmt.Println(v21)
        fmt.Println(v21, v23, v21)
        v186 := GenS24{F0: 4294967295, F1: 380}
        _ = v186
    }
    switch "a" {
    case "":
        v5[42] = ("hello" + "foo bar")
    case "a":
        switch "\n\t" {
        case "hello":
            v187 := GenS4{F0: "foo bar", F1: 1, F2: 18446744073709551615}
            _ = v187
        }
    }
    v1 -= 0
    v4 = append(v4, 511)
    var v188 int16 = 633
    var v189 int32 = 441
    fmt.Println("v1:", v1)
    fmt.Println("v188:", v188)
    fmt.Println("v189:", v189)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v188
    _ = v189
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
