package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS8 struct { F0 uint16; F1 int64; F2 uint }
type GenS9 struct { F0 int32; F1 int32; F2 uint32 }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 int64; F1 int8; F2 string }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint; F1 bool }
type GenS17 struct { F0 int16; F1 int; F2 string }
type GenS18 struct { F0 uint64; F1 uint16 }
type GenS19 struct { F0 uint64; F1 uint; F2 int }
func genFn0() int8 {
    var v4 float32 = 1e10
    for fi5 := 0; fi5 < 14; fi5++ {
        switch "\n\t" {
        case "":
            fmt.Println(fi5)
        default:
            fmt.Println(fi5, fi5)
        }
        var v6 string = ("hello" + (("\n\t" + "a") + "a"))
        _ = fi5
        _ = v6
    }
    fmt.Println(v4)
    _ = v4
    return 126
}
func genFn1(p0 int8, p1 int, p2 bool) string {
    fmt.Println(p2, p1, p1)
    switch p1 {
    case 128:
        var v7 uint8 = 0
        _ = v7
    case -472:
        if (0.1 < 1e10) {
            fmt.Println(p2, p1, p1)
        } else {
            if p2 {
                p2 = false
                v8 := []int{42}
                for fi9 := 0; fi9 < 4; fi9++ {
                    p1 = 850
                    p1 -= 783
                    v8 = append(v8, 128)
                    _ = fi9
                }
                _ = v8
            } else {
                if p2 {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v11 := []int{1, int(p0)}
                        v12 := GenS1{F0: 550, F1: p1, F2: 18446744073709551615}
                        p0 -= (p0 ^ (p0 & p0))
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    fmt.Println(p2)
                    p2 = (42 < 993)
                } else {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 int = fi13
                        for fi15 := 0; fi15 < 6; fi15++ {
                            switch "hello" {
                            case "a":
                                var v16 uint8 = 115
                                _ = v16
                            case "foo bar":
                                var v17 int32 = 332
                                _ = v17
                            }
                            if (p2 && (!(false || false))) {
                                p2 = ((p2 && (!(9223372036854775807 <= (p1 * 1)))) || (1 != -2147483648))
                                var v18 uint8 = 42
                                _ = v18
                            }
                            _ = fi15
                        }
                        _ = fi13
                        _ = v14
                    }
                    if p2 {
                        var v19 int16 = -32768
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            p1 += p1
                            fmt.Println(p0, p0, p1)
                            p1 += p1
                            _ = fi20
                        }
                        var v21 bool = true
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS2{F0: 32767, F1: fi22}
                    for fi24 := 0; fi24 < 4; fi24++ {
                        fmt.Println(p0, fi24, p2)
                        var v25 float32 = 1e10
                        v26 := GenS3{F0: (p0 % ((p0) | 1))}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                }
            }
            fmt.Println(p0, p0)
        }
    case 0:
        p0 = genFn0()
    default:
        p1 = p1
    }
    var v27 bool = p2
    var v28 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return "a"
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []int{}
    var v3 float32 = -2.25
    v29 := GenS4{F0: 405, F1: -9223372036854775808}
    var v30 bool = ((!false) || false)
    switch v1 {
    case -2147483648:
        switch 0 {
        case 127:
            if v30 {
                if true {
                    fmt.Println(v30)
                    var v31 bool = v30
                    _ = v31
                }
                v30 = true
            }
        case -466:
            v2 = append(v2, v1)
        }
    case 2147483647:
        for fi32 := 0; fi32 < 20; fi32++ {
            if v30 {
                v33 := GenS1{F0: 42, F1: (v1 - (fi32 | fi32)), F2: 255}
                _ = len(v2)
                var v34 int64 = 0
                _ = v33
                _ = v34
            }
            v30 = (-2.25 == float32(1e10))
            v30 = ((2 > 1) || v30)
            _ = fi32
        }
    case 128:
        if v30 {
            var v35 int64 = -657
            for fi36 := 0; fi36 < 10; fi36++ {
                switch v1 {
                case 1:
                    switch fi36 {
                    case 4294967295:
                        if (!(-1 < genFn0())) {
                            v37 := GenS5{F0: 255, F1: 1}
                            v38 := GenS4{F0: 32767, F1: (-(v35))}
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS6{F0: (v35 << ((255) & 63)), F1: (false && (42 > 1))}
                            for fi40 := 0; fi40 < 17; fi40++ {
                                fmt.Println(fi40, v1)
                                v30 = (4294967295 <= 4294967295)
                                _ = fi40
                            }
                            fmt.Println(fi36, fi36)
                            _ = v39
                        }
                    case 679:
                        v30 = false
                    default:
                        fmt.Println(v1, v30)
                    }
                case -9223372036854775808:
                    var v41 int8 = 1
                    _ = v41
                case -1:
                    v42 := GenS7{F0: 490, F1: 2, F2: 0}
                    _ = v42
                default:
                    switch "hello" {
                    case "hello":
                        if (!v30) {
                            v43 := GenS8{F0: 32767, F1: v35, F2: 539}
                            v30 = (!((168 <= 32767) && true))
                            _ = v43
                        } else {
                            v3 += 0.0
                        }
                    case "foo bar":
                        switch (("hello" + "") + "hello") {
                        case "":
                            var v44 uint16 = 2
                            _ = v44
                        default:
                            if v30 {
                                v1 -= 374
                                switch ("a" + "") {
                                case "\n\t":
                                    v30 = v30
                                default:
                                    if v30 {
                                        fmt.Println(v30)
                                    } else {
                                        var v45 int16 = 1000
                                        v46 := GenS2{F0: 0, F1: (fi36 / ((285) | 1))}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v45
                                        _ = v46
                                    }
                                }
                            }
                        }
                    default:
                        switch v1 {
                        case -389:
                            var v47 float32 = (1e10 + v3)
                            _ = v47
                        case -2147483648:
                            if len(v2) > 0 {
                                v2[0] = (fi36 | fi36)
                            }
                        case -9223372036854775808:
                            v48 := GenS8{F0: 26, F1: 0, F2: 18446744073709551615}
                            _ = v48
                        default:
                            var v49 uint16 = 2
                            _ = v49
                        }
                    }
                }
                _ = fi36
            }
            v3 = v3
            _ = v35
        } else {
            switch "a" {
            case "hello":
                v3 -= (v3 / (v3 + float32((0.0 / 1e10))))
            case "foo bar":
                v3 -= 1.5
            }
            v2 = append(v2, -258)
        }
    default:
        if len(v2) > 0 {
            v2[0] = (v1 + v1)
        }
    }
    var v50 uint32 = 0
    switch 907 {
    case 9223372036854775807:
        v51 := GenS4{F0: 2, F1: 0}
        _ = v51
    case 1:
        var v52 uint8 = 51
        _ = v52
    case 42:
        if false {
            _ = len(v2)
        }
    default:
        v50 -= (v50 / ((v50) | 1))
    }
    if (!v30) {
        v53 := GenS1{F0: v50, F1: v1, F2: 65535}
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        var v54 uint64 = 0
        _ = v53
        _ = v54
    }
    var v55 uint = 42
    switch v1 {
    case 9223372036854775807:
        if (0.1 <= -2.25) {
            v50 = v50
        } else {
            switch "foo bar" {
            case "foo bar":
                v3 = v3
            }
            v56 := GenS9{F0: 1, F1: -1, F2: v50}
            fmt.Println(v1, v55, v3)
            _ = v56
        }
    case 1:
        var v57 float32 = 0.1
        _ = v57
    }
    v58 := GenS10{F0: v1, F1: ("a" <= ("" + (("hello" + "hello") + "a")))}
    if v30 {
        for fi59 := 0; fi59 < 12; fi59++ {
            switch ("\n\t" + "hello") {
            case "\n\t":
                if (0.1 <= float64(v3)) {
                    v1 += 2147483647
                    var v60 uint8 = 214
                    var v61 int8 = -1
                    _ = v60
                    _ = v61
                } else {
                    v62 := GenS2{F0: 255, F1: 32767}
                    for fi63 := 0; fi63 < 13; fi63++ {
                        if (42 != 193) {
                            var v64 bool = v30
                            _ = v64
                        } else {
                            v65 := GenS11{F0: -9223372036854775808, F1: genFn0(), F2: "a"}
                            _ = v65
                        }
                        if true {
                            var v66 int32 = 813
                            _ = v66
                        }
                        _ = fi63
                    }
                    _ = v62
                }
            case "":
                fmt.Println(v30)
            }
            switch "a" {
            case "\n\t":
                var v67 uint64 = 814
                _ = v67
            case "a":
                for fi68 := 0; fi68 < 5; fi68++ {
                    v2 = append(v2, fi68)
                    _ = fi68
                }
            case "":
                for fi69 := 0; fi69 < 18; fi69++ {
                    switch (fi69 << (((v55 >> ((v55) & 63))) & 63)) {
                    case 128:
                        if v30 {
                            if ((^(v50)) >= (v50 + (v50 - (v50 % (((v50 * (v50 + (v50 >> ((42) & 63))))) | 1))))) {
                                var v70 uint64 = 42
                                if false {
                                    if false {
                                        v70 += 4294967295
                                        _ = len(v2)
                                    } else {
                                        _ = len(v2)
                                        v70 = 4294967295
                                        fmt.Println(v1, v3, v70)
                                    }
                                    if (v30 && v30) {
                                        if (("hello" == "\n\t") || ((float64(v3) >= float64((v3 - (1e10 / 1e10)))) || v30)) {
                                            fmt.Println(v1, v1)
                                        }
                                        var v71 bool = (!true)
                                        _ = len(v2)
                                        v72 := GenS9{F0: -1, F1: -52, F2: v50}
                                        fmt.Println(v30)
                                        switch "foo bar" {
                                        case "":
                                            if (!v71) {
                                                var v73 uint16 = 65535
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 127
                                                }
                                                _ = v73
                                            } else {
                                                var v74 int16 = 0
                                                v50 = 197
                                                if true {
                                                    v70 += uint64(v50)
                                                    switch ("a" + "hello") {
                                                    case "a":
                                                        v50 = (v50 / (((^(v50))) | 1))
                                                    default:
                                                        if v71 {
                                                            v75 := GenS12{F0: (v74 ^ v74)}
                                                            if v71 {
                                                                v55 = v55
                                                                if (-2 >= v1) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v1
                                                                    }
                                                                    var v76 float64 = (1e10 * 1.5)
                                                                    _ = v76
                                                                } else {
                                                                    v55 += (^(v55))
                                                                    v2 = append(v2, (fi69 * (fi69 - (v1 << ((2) & 63)))))
                                                                }
                                                            }
                                                            _ = v75
                                                        } else {
                                                            if false {
                                                                v77 := GenS1{F0: v50, F1: (fi69 << ((v55) & 63)), F2: (v55 >> (((v55 | v55)) & 63))}
                                                                v78 := GenS1{F0: 0, F1: fi69, F2: v55}
                                                                switch ("foo bar" + "foo bar") {
                                                                case "a":
                                                                    if v71 {
                                                                        fmt.Println(v70, v1, v71)
                                                                        fmt.Println(fi59, fi69, v3)
                                                                        v2 = append(v2, fi59)
                                                                    } else {
                                                                        var v79 bool = (!(!true))
                                                                        _ = v79
                                                                    }
                                                                case "hello":
                                                                    _ = len(v2)
                                                                default:
                                                                    v71 = (618 == 42)
                                                                }
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            v1 = fi59
                                                            if v30 {
                                                                v3 = v3
                                                                var v80 uint32 = (v50 + v50)
                                                                _ = v80
                                                            }
                                                        }
                                                    }
                                                    v2 = append(v2, (fi59 % ((fi69) | 1)))
                                                } else {
                                                    v3 -= v3
                                                }
                                                _ = v74
                                            }
                                        case "hello":
                                            var v81 int64 = 0
                                            _ = v81
                                        case "a":
                                            fmt.Println(v1)
                                        default:
                                            fmt.Println(v50, v30)
                                        }
                                        _ = v71
                                        _ = v72
                                    } else {
                                        v82 := GenS3{F0: 100}
                                        switch (fi69 * -2147483648) {
                                        case -872:
                                            v70 += (v70 >> (((v55 * v55)) & 63))
                                            v3 -= (0.0 - float32((-2.25 * 1e10)))
                                            if (1 == -2147483648) {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if (3.14159 == 3.14159) {
                                                        if (v30 || v30) {
                                                            _ = len(v2)
                                                            v1 += (fi69 + fi69)
                                                            v55 = (v55 ^ v55)
                                                            var v83 float64 = -2.25
                                                            _ = v83
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi69 * (fi69 / (((-(v1))) | 1)))
                                                        }
                                                    } else {
                                                        v55 = v55
                                                        if v30 {
                                                            v2 = append(v2, (-(fi59)))
                                                            v84 := GenS13{F0: 64}
                                                            var v85 uint32 = 42
                                                            var v86 int16 = int16(v70)
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            if true {
                                                                v1 = 9223372036854775807
                                                                switch fi59 {
                                                                case 42:
                                                                    var v87 uint64 = uint64(v50)
                                                                    _ = v87
                                                                case 146:
                                                                    if ((v70 * v70) < 4294967295) {
                                                                        v88 := GenS6{F0: 421, F1: (fi69 == v1)}
                                                                        if (v55 > (v55 | v55)) {
                                                                            var v89 int = (v1 * fi69)
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi59
                                                                            }
                                                                            v90 := GenS3{F0: 0}
                                                                            _ = v89
                                                                            _ = v90
                                                                        } else {
                                                                            var v91 int32 = 2147483647
                                                                            if (0 > -1) {
                                                                                v92 := GenS13{F0: 0}
                                                                                v1 = (fi59 % (((v1 - (v1 ^ fi59))) | 1))
                                                                                _ = v92
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = int(v70)
                                                                            }
                                                                            v93 := GenS10{F0: (^(fi59)), F1: (v91 == v91)}
                                                                            var v94 uint32 = (v50 >> (((v55 << ((4294967295) & 63))) & 63))
                                                                            _ = v91
                                                                            _ = v93
                                                                            _ = v94
                                                                        }
                                                                        _ = v88
                                                                    } else {
                                                                        v55 -= v55
                                                                    }
                                                                case 2147483647:
                                                                    var v95 uint = uint(v50)
                                                                    _ = v95
                                                                default:
                                                                    switch v1 {
                                                                    case -9223372036854775808:
                                                                        if ((1 >= 1) || true) {
                                                                            v55 = v55
                                                                            v96 := GenS4{F0: 32767, F1: -9223372036854775808}
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                    v97 := GenS7{F0: v70, F1: 65535, F2: 0}
                                                                    v1 += v1
                                                                    _ = v97
                                                                }
                                                                fmt.Println(v3, fi69, v3)
                                                                if true {
                                                                    fmt.Println(fi69)
                                                                }
                                                            } else {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    if v30 {
                                                                        v55 -= v55
                                                                    } else {
                                                                        v70 = v70
                                                                        if v30 {
                                                                            v1 -= v1
                                                                            var v98 uint64 = (^(v70))
                                                                            v1 += 2147483647
                                                                            _ = v98
                                                                        } else {
                                                                            var v99 int = v1
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v50 = (^(v50))
                                                                            case "":
                                                                                if false {
                                                                                    var v100 float64 = 3.14159
                                                                                    _ = v100
                                                                                }
                                                                            case "\n\t":
                                                                                v101 := GenS9{F0: 476, F1: int32(v70), F2: (v50 * (v50 / ((v50) | 1)))}
                                                                                _ = v101
                                                                            default:
                                                                                if (1 < -1) {
                                                                                    var v102 uint32 = v50
                                                                                    var v103 uint8 = 255
                                                                                    fmt.Println(v50, v30, v3)
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                }
                                                                            }
                                                                            _ = v99
                                                                        }
                                                                    }
                                                                case "a":
                                                                    if (-5 > 100) {
                                                                        v104 := GenS4{F0: 32767, F1: int64(v50)}
                                                                        fmt.Println(v55, fi69, v30)
                                                                        _ = v104
                                                                    } else {
                                                                        v3 += 0.0
                                                                        v105 := GenS8{F0: 42, F1: 0, F2: v55}
                                                                        _ = v105
                                                                    }
                                                                default:
                                                                    if (!v30) {
                                                                        fmt.Println(fi69, v1)
                                                                    }
                                                                    switch (fi69 % ((1) | 1)) {
                                                                    case 2147483647:
                                                                        var v106 uint16 = 32767
                                                                        _ = v106
                                                                    case -244:
                                                                        if v30 {
                                                                            var v107 uint8 = 26
                                                                            var v108 int16 = 699
                                                                            var v109 float32 = 1e10
                                                                            fmt.Println(v109)
                                                                            v110 := GenS14{F0: "hello"}
                                                                            _ = v107
                                                                            _ = v108
                                                                            _ = v109
                                                                            _ = v110
                                                                        } else {
                                                                            v55 -= v55
                                                                            v3 -= (float32(-2.25) * ((v3 * v3) / 1e10))
                                                                            var v111 float64 = 1e10
                                                                            _ = v111
                                                                        }
                                                                    default:
                                                                        v112 := GenS15{F0: v70}
                                                                        _ = v112
                                                                    }
                                                                    var v113 int32 = 1
                                                                    _ = v113
                                                                }
                                                                v1 += fi59
                                                                v114 := GenS3{F0: -69}
                                                                _ = v114
                                                            }
                                                            v115 := GenS16{F0: v55, F1: v30}
                                                            _ = v115
                                                        }
                                                        if false {
                                                            if v30 {
                                                                if (!(v30 || v30)) {
                                                                    v3 += 1e10
                                                                    if (-2.25 != (v3 * 1e10)) {
                                                                        v1 += (fi69 | (^(fi59)))
                                                                        v116 := GenS12{F0: 247}
                                                                        _ = v116
                                                                    }
                                                                } else {
                                                                    fmt.Println(v3)
                                                                    if false {
                                                                        if (127 > genFn0()) {
                                                                            var v117 string = ("foo bar" + "a")
                                                                            v118 := GenS15{F0: v70}
                                                                            v1 = (-(fi59))
                                                                            _ = v117
                                                                            _ = v118
                                                                        } else {
                                                                            var v119 bool = v30
                                                                            v120 := GenS3{F0: 0}
                                                                            _ = v119
                                                                            _ = v120
                                                                        }
                                                                    }
                                                                    fmt.Println(v30, fi69, v1)
                                                                }
                                                                v121 := GenS8{F0: 315, F1: -1, F2: v55}
                                                                fmt.Println(v30, v1, v55)
                                                                v122 := GenS8{F0: 1, F1: 9223372036854775807, F2: (^(v55))}
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v121
                                                                _ = v122
                                                            }
                                                        } else {
                                                            var v123 uint64 = v70
                                                            if ("" >= ("" + "foo bar")) {
                                                                var v124 int = 42
                                                                _ = v124
                                                            }
                                                            v125 := GenS13{F0: 1}
                                                            _ = v123
                                                            _ = v125
                                                        }
                                                    }
                                                }
                                                v70 = uint64(v55)
                                            } else {
                                                var v126 int32 = 451
                                                _ = v126
                                            }
                                            if (v30 || v30) {
                                                v70 += 96
                                            } else {
                                                var v127 uint64 = (^(v70))
                                                if (v30 || ((-2.25 * 1e10) <= 1e10)) {
                                                    if (!(v3 >= float32(3.14159))) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v128 := GenS17{F0: 1000, F1: fi59, F2: ("hello" + "foo bar")}
                                                        _ = v128
                                                    } else {
                                                        if v30 {
                                                            v129 := GenS7{F0: v127, F1: (v55 / (((v55 | v55)) | 1)), F2: (v50 ^ v50)}
                                                            if (false && v30) {
                                                                var v130 string = "hello"
                                                                fmt.Println(v50, v55, v55)
                                                                _ = v130
                                                            }
                                                            v131 := GenS2{F0: 32767, F1: fi59}
                                                            v132 := GenS16{F0: uint(v127), F1: v30}
                                                            v1 += -403
                                                            _ = v129
                                                            _ = v131
                                                            _ = v132
                                                        }
                                                    }
                                                    v133 := GenS18{F0: (v127 / (((v70 - v127)) | 1)), F1: 1}
                                                    _ = len(v2)
                                                    fmt.Println(v70)
                                                    v30 = (!((!((v30 || true) && v30)) && true))
                                                    var v134 string = ("" + "foo bar")
                                                    _ = v133
                                                    _ = v134
                                                }
                                                v135 := GenS9{F0: 42, F1: -2147483648, F2: (v50 - 42)}
                                                _ = v127
                                                _ = v135
                                            }
                                        case -1:
                                            if (!(false || ((("hello" + "") <= "foo bar") || (2147483647 <= 1)))) {
                                                v1 = fi69
                                            } else {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi59 << (((v55 & 65535)) & 63))
                                                }
                                            }
                                        case -6:
                                            switch (fi69 ^ (^(fi69))) {
                                            case 157:
                                                if (9223372036854775807 == 9223372036854775807) {
                                                    v55 += v55
                                                } else {
                                                    fmt.Println(v3, v1)
                                                    v55 = uint(v50)
                                                }
                                            case -223:
                                                if (((false && (true && ((!(!true)) || (((100 > genFn0()) || v30) && (v30 && (v50 > v50)))))) || (2 == 129)) && true) {
                                                    if (!(65535 <= 65535)) {
                                                        if v30 {
                                                            _ = len(v2)
                                                            v136 := GenS7{F0: v70, F1: v55, F2: (v50 * v50)}
                                                            var v137 int8 = 100
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                    }
                                                    if (1 >= (fi59 / ((fi69) | 1))) {
                                                        fmt.Println(v1, fi69, v3)
                                                        var v138 int32 = -838
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if v30 {
                                                            v30 = true
                                                        } else {
                                                            var v139 uint16 = 652
                                                            if len(v2) > 0 {
                                                                v2[0] = fi69
                                                            }
                                                            v55 = uint(v139)
                                                            _ = v139
                                                        }
                                                        _ = v138
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = fi59
                                                        }
                                                        if false {
                                                            v55 = v55
                                                        } else {
                                                            var v140 int = (v1 ^ (-(fi69)))
                                                            v141 := GenS6{F0: 42, F1: v30}
                                                            v2 = append(v2, fi59)
                                                            _ = v140
                                                            _ = v141
                                                        }
                                                        if v30 {
                                                            var v142 uint64 = v70
                                                            _ = v142
                                                        } else {
                                                            v2 = append(v2, (fi59 / ((0) | 1)))
                                                            v55 = v55
                                                            var v143 uint16 = 0
                                                            _ = v143
                                                        }
                                                    }
                                                    _ = len(v2)
                                                } else {
                                                    if v30 {
                                                        var v144 string = "\n\t"
                                                        _ = v144
                                                    } else {
                                                        v70 += v70
                                                        v55 = 255
                                                        v145 := GenS13{F0: 2}
                                                        _ = v145
                                                    }
                                                }
                                            default:
                                                var v146 int16 = 721
                                                _ = v146
                                            }
                                        }
                                        _ = v82
                                    }
                                    v30 = true
                                } else {
                                    var v147 float64 = ((-2.25 * ((((float64(((v3 / 1e10) * v3)) + float64((float32(0.1) + (v3 - (v3 - (v3 * (float32(float64(float32(0.0))) * (1.5 * float32((float64((((v3 - v3) + v3) + v3)) - (float64(1e10) + 3.14159))))))))))) / 3.14159) / -2.25) + 1e10)) / 1.0)
                                    v3 = v3
                                    var v148 int32 = -831
                                    _ = v147
                                    _ = v148
                                }
                                v50 = 2147483648
                                v149 := GenS19{F0: v70, F1: 110, F2: (fi69 + (fi59 << (((v55 - v55)) & 63)))}
                                v150 := GenS2{F0: 1, F1: (v1 + 32767)}
                                _ = v149
                                _ = v150
                                _ = v70
                            }
                        } else {
                            v151 := GenS8{F0: 813, F1: 42, F2: 255}
                            _ = len(v2)
                            _ = v151
                        }
                    case -9223372036854775808:
                        switch ("a" + "hello") {
                        case "":
                            switch "hello" {
                            case "\n\t":
                                v55 = v55
                            case "":
                                _ = len(v2)
                            case "foo bar":
                                var v152 bool = v30
                                _ = v152
                            }
                        }
                    }
                    _ = len(v2)
                    v50 -= v50
                    _ = fi69
                }
            }
            _ = fi59
        }
        _ = len(v2)
        for fi153 := 0; fi153 < 4; fi153++ {
            if true {
                var v154 uint32 = (v50 + v50)
                v154 += (v50 * v154)
                _ = v154
            } else {
                for fi155 := 0; fi155 < 14; fi155++ {
                    v3 = v3
                    var v156 uint32 = v50
                    _ = fi155
                    _ = v156
                }
                if (766 < 88) {
                    if (-1 > -63) {
                        var v157 uint = (v55 - 1)
                        if true {
                            var v158 int64 = -9223372036854775808
                            _ = v158
                        } else {
                            fmt.Println(v50)
                            var v159 float32 = ((1e10 / -2.25) * (3.14159 - v3))
                            _ = v159
                        }
                        _ = v157
                    } else {
                        for fi160 := 0; fi160 < 6; fi160++ {
                            fmt.Println(v1, v30, v50)
                            switch (^(v1)) {
                            case 32767:
                                v55 = (v55 - v55)
                                var v161 int8 = -110
                                switch (((("a" + "a") + "hello") + "") + "\n\t") {
                                case "hello":
                                    if true {
                                        v30 = false
                                        var v162 float32 = float32(3.14159)
                                        v162 = (v3 * ((v162 / v162) + (float32(3.14159) / 0.1)))
                                        _ = v162
                                    } else {
                                        var v163 uint64 = 18446744073709551615
                                        var v164 int64 = 9223372036854775807
                                        _ = v163
                                        _ = v164
                                    }
                                case "\n\t":
                                    v50 -= 97
                                case "a":
                                    var v165 bool = v30
                                    _ = v165
                                }
                                fmt.Println(v50, v55)
                                v30 = ((v161 >> ((v55) & 63)) < v161)
                                _ = v161
                            case 598:
                                if len(v2) > 0 {
                                    v2[0] = (v1 | (fi153 * (fi153 - 128)))
                                }
                            case 32768:
                                v2 = append(v2, 9223372036854775807)
                            }
                            v50 += v50
                            _ = fi160
                        }
                    }
                    v3 -= -2.25
                } else {
                    if true {
                        var v166 uint16 = 255
                        fmt.Println(v166, v166, v30)
                        var v167 int32 = -405
                        _ = v166
                        _ = v167
                    } else {
                        for fi168 := 0; fi168 < 20; fi168++ {
                            var v169 float64 = 0.1
                            v50 -= v50
                            switch 128 {
                            case 1:
                                switch 146 {
                                case -171:
                                    var v170 uint = (v55 | (v55 % (((^(v55))) | 1)))
                                    _ = v170
                                default:
                                    _ = len(v2)
                                }
                            default:
                                if (370 > 42) {
                                    switch ((("" + "hello") + ("hello" + "foo bar")) + ("foo bar" + "\n\t")) {
                                    case "":
                                        fmt.Println(v3, v30)
                                    case "a":
                                        fmt.Println(v3, fi153)
                                    }
                                    v171 := GenS11{F0: -542, F1: 0, F2: ""}
                                    _ = v171
                                }
                                _ = len(v2)
                                var v172 uint8 = 137
                                _ = v172
                            }
                            _ = fi168
                            _ = v169
                        }
                        v55 += v55
                    }
                }
                v1 = (v1 * 361)
            }
            for fi173 := 0; fi173 < 2; fi173++ {
                v2 = append(v2, 32767)
                v1 += 1
                if v30 {
                    v30 = true
                }
                _ = fi173
            }
            _ = fi153
        }
    } else {
        var v174 uint8 = 1
        for fi175 := 0; fi175 < 4; fi175++ {
            if true {
                for fi176 := 0; fi176 < 20; fi176++ {
                    v55 -= (v55 * v55)
                    fmt.Println(fi175)
                    v174 = (v174 | 255)
                    fmt.Println(v174, v30)
                    _ = fi176
                }
                switch ("foo bar" + "a") {
                case "\n\t":
                    if v30 {
                        _ = len(v2)
                        _ = len(v2)
                    } else {
                        v2 = append(v2, -1)
                        v174 = (v174 - (v174 & 0))
                        var v177 uint64 = 42
                        _ = v177
                    }
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    for fi178 := 0; fi178 < 15; fi178++ {
                        switch -366 {
                        case 128:
                            var v179 float64 = (((-2.25 + (0.1 + 1.5)) * 1.5) + ((1.5 + 0.0) - 1.5))
                            _ = v179
                        case 4294967295:
                            if len(v2) > 0 {
                                v2[0] = (fi175 >> ((v55) & 63))
                            }
                        }
                        if (!(uint(v50) == (v55 << ((v55) & 63)))) {
                            var v180 uint64 = uint64(v50)
                            _ = v180
                        }
                        _ = fi178
                    }
                }
                if (v30 && (true || v30)) {
                    var v181 bool = false
                    v1 = (fi175 / (((v1 >> ((2) & 63))) | 1))
                    var v182 uint8 = v174
                    _ = v181
                    _ = v182
                }
            } else {
                v183 := GenS14{F0: (("" + ("foo bar" + ("a" + "a"))) + "foo bar")}
                switch "" {
                case "foo bar":
                    fmt.Println(v1)
                case "":
                    v184 := GenS1{F0: v50, F1: (v1 ^ 128), F2: (v55 / (((v55 >> ((v55) & 63))) | 1))}
                    _ = v184
                default:
                    var v185 uint32 = 112
                    _ = v185
                }
                _ = v183
            }
            var v186 uint64 = 545
            v186 += v186
            _ = fi175
            _ = v186
        }
        _ = v174
    }
    v50 -= 0
    v2 = append(v2, v1)
    for fi187 := 0; fi187 < 14; fi187++ {
        switch v1 {
        case 1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 4294967295:
            for fi188 := 0; fi188 < 11; fi188++ {
                v189 := GenS17{F0: -1, F1: fi188, F2: "hello"}
                _ = fi188
                _ = v189
            }
        case -1:
            if len(v2) > 0 {
                v2[0] = 42
            }
        default:
            if (-203 < 1) {
                v3 = (v3 + ((v3 + ((v3 / (float32(((-2.25 + 0.1) - (float64(3.14159) + ((((3.14159 - (1e10 / 3.14159)) / 3.14159) / 0.1) + float64(v3))))) / v3)) - v3)) * v3))
                if true {
                    var v190 float32 = 3.14159
                    _ = v190
                }
                var v191 int = (-(fi187))
                _ = v191
            }
        }
        v192 := GenS3{F0: 1}
        _ = fi187
        _ = v192
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v50
    _ = v55
    _ = v58
}
