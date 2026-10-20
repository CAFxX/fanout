package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int; F1 int32; F2 int64 }
type GenS7 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 int16, p1 int, p2 uint8) int32 {
    v6 := []int{}
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: uint(p2)}
        _ = fi7
        _ = v8
    }
    var v9 bool = true
    var v10 int8 = 127
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            v10 -= v10
            v13 := GenS1{F0: 1}
            v9 = false
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return 42
}
func genFn1(p0 int16, p1 uint, p2 int16) uint64 {
    p1 += p1
    for fi14 := 0; fi14 < 6; fi14++ {
        p2 -= (p0 >> ((65535) & 63))
        fmt.Println(p2)
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    v1 := []int{}
    var v2 uint16 = 2
    var v3 int8 = 0
    v4 := make(map[string]bool)
    v5 := []string{("" + ("" + ("a" + "foo bar"))), (("a" + "") + ""), "\n\t"}
    fmt.Println(v2, v3)
    v3 += v3
    for fi15 := 0; fi15 < 9; fi15++ {
        switch "hello" {
        case "\n\t":
            v2 = 1
        case "a":
            fmt.Println(fi15, v2, fi15)
        }
        v16 := GenS1{F0: 984}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 406}
    if (1e10 < float32(-2.25)) {
        var v18 int16 = 1
        _ = v18
    } else {
        var v19 int32 = 2147483647
        switch (("\n\t" + "a") + "hello") {
        case "":
            var v20 uint = 711
            _ = v20
        default:
            var v21 int32 = v19
            _ = v21
        }
        _ = v19
    }
    var v22 int32 = 0
    if (!(0 >= genFn0(-382, -1, 223))) {
        v2 -= 0
    } else {
        v3 -= (v3 | (v3 | (^(v3))))
        var v23 uint32 = 2
        if len(v1) > 0 {
            v1[0] = 128
        }
        _ = v23
    }
    v22 -= (^(v22))
    switch ((("" + "foo bar") + "foo bar") + ("foo bar" + ("\n\t" + "a"))) {
    case "":
        v2 -= v2
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS2{F0: fi24}
        _ = fi24
        _ = v25
    }
    if false {
        if (18446744073709551615 == 18446744073709551615) {
            if false {
                v22 = 42
                var v26 uint32 = 2147483647
                _ = v26
            }
            v27 := GenS2{F0: 32768}
            if true {
                v22 += v22
                fmt.Println(v2)
                delete(v4, "\n\t")
            } else {
                if false {
                    _ = len(v4)
                    for fi28 := 0; fi28 < 8; fi28++ {
                        var v29 int = (^(fi28))
                        v29 = (fi28 % (((v29 >> ((2) & 63))) | 1))
                        switch v29 {
                        case -681:
                            if ((!false) || true) {
                                for fi30 := 0; fi30 < 6; fi30++ {
                                    v4["hello"] = (false && (v3 > v3))
                                    _ = fi30
                                }
                                fmt.Println(v22, v2, v3)
                                var v31 int32 = 42
                                _ = v31
                            } else {
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    switch (fi32 & (fi28 << ((446) & 63))) {
                                    case -317:
                                        if false {
                                            if false {
                                                var v33 int = 4294967295
                                                _ = v33
                                            }
                                            fmt.Println(v3, fi28, fi28)
                                            delete(v4, "\n\t")
                                        } else {
                                            v3 -= (v3 << ((255) & 63))
                                            v3 -= v3
                                            switch "a" {
                                            case "hello":
                                                var v34 float32 = 1e10
                                                _ = v34
                                            case "foo bar":
                                                var v35 string = "\n\t"
                                                _ = v35
                                            default:
                                                fmt.Println(v22)
                                            }
                                        }
                                    }
                                    _ = fi32
                                }
                                for fi36 := 0; fi36 < 11; fi36++ {
                                    fmt.Println(fi36, v22, v2)
                                    v37 := GenS2{F0: 2147483647}
                                    switch "" {
                                    case "":
                                        var v38 uint64 = 2
                                        _ = v38
                                    default:
                                        v22 = v22
                                    }
                                    _ = fi36
                                    _ = v37
                                }
                            }
                        default:
                            switch "hello" {
                            case "hello":
                                for fi39 := 0; fi39 < 9; fi39++ {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v22)
                                    case "hello":
                                        var v40 float32 = (1e10 + 1.5)
                                        _ = v40
                                    case "":
                                        if (18446744073709551615 == 42) {
                                            var v41 uint = 766
                                            v42 := GenS1{F0: v41}
                                            v22 -= -126
                                            var v43 int16 = 32767
                                            _ = v41
                                            _ = v42
                                            _ = v43
                                        } else {
                                            if ((696 != -9223372036854775808) && ((-(v29)) >= (v29 & (fi39 / ((v29) | 1))))) {
                                                if len(v1) > 0 {
                                                    v1[0] = (fi39 ^ 128)
                                                }
                                                if (-1 <= 481) {
                                                    v44 := v4[("foo bar" + ("a" + "\n\t"))]
                                                    v45 := GenS2{F0: v29}
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    _ = len(v5)
                                                }
                                                switch ("hello" + "\n\t") {
                                                case "":
                                                    var v46 int32 = v22
                                                    var v47 int32 = genFn0(-1, (fi28 * v29), 2)
                                                    fmt.Println(v22, v29, v2)
                                                    _ = v46
                                                    _ = v47
                                                default:
                                                    var v48 uint = 0
                                                    _ = v48
                                                }
                                            } else {
                                                delete(v4, (("foo bar" + ("\n\t" + ((("\n\t" + "\n\t") + "foo bar") + "\n\t"))) + ("a" + (("foo bar" + "\n\t") + "foo bar"))))
                                                var v49 int64 = 813
                                                _ = v49
                                            }
                                            var v50 uint32 = 65535
                                            fmt.Println(fi39)
                                            var v51 int32 = 2147483647
                                            v4[("hello" + "hello")] = false
                                            _ = v50
                                            _ = v51
                                        }
                                    }
                                    _ = fi39
                                }
                            case "a":
                                for fi52 := 0; fi52 < 18; fi52++ {
                                    if (false || ((!false) || (false || (v3 < v3)))) {
                                        var v53 string = "foo bar"
                                        if (!((^(v2)) < v2)) {
                                            fmt.Println(v53, v3, v29)
                                            fmt.Println(v29, fi28)
                                        } else {
                                            v54 := GenS3{F0: v29, F1: 644}
                                            switch (fi28 ^ -2147483648) {
                                            case 127:
                                                if (246 < (^(v2))) {
                                                    var v55 int8 = int8(v22)
                                                    _ = v55
                                                } else {
                                                    if ((920 >= 0) && (false && ((223 > 18446744073709551615) && ((v3 ^ v3) < int8(v2))))) {
                                                        switch fi28 {
                                                        case 0:
                                                            v56 := v4[v53]
                                                            _ = v56
                                                        }
                                                        switch int(v22) {
                                                        case 9223372036854775807:
                                                            var v57 uint16 = v2
                                                            _ = v57
                                                        case -9223372036854775808:
                                                            fmt.Println(v22, v2)
                                                            fmt.Println(v53)
                                                            if (0 != 255) {
                                                                if true {
                                                                    var v58 uint32 = 2
                                                                    _ = v58
                                                                } else {
                                                                    v4[((v53 + (("" + "") + "")) + v53)] = false
                                                                    if false {
                                                                        var v59 uint16 = v2
                                                                        if (0 > 0) {
                                                                            v60 := GenS1{F0: 0}
                                                                            if (v53 > (v53 + "hello")) {
                                                                                var v61 bool = (!true)
                                                                                v59 += 951
                                                                                _ = v61
                                                                            } else {
                                                                                v2 = v59
                                                                            }
                                                                            var v62 int8 = 127
                                                                            _ = v60
                                                                            _ = v62
                                                                        }
                                                                        v63 := GenS4{F0: -9223372036854775808}
                                                                        _ = v59
                                                                        _ = v63
                                                                    }
                                                                    switch 431 {
                                                                    case -499:
                                                                        switch (v29 / (((fi52 & fi28)) | 1)) {
                                                                        case 127:
                                                                            if true {
                                                                                if false {
                                                                                    v2 += (v2 & 0)
                                                                                    v64 := GenS3{F0: 0, F1: 2147483647}
                                                                                    _ = v64
                                                                                }
                                                                            }
                                                                        case 0:
                                                                            fmt.Println(v29, v22, fi28)
                                                                        default:
                                                                            if true {
                                                                                _ = len(v1)
                                                                                v65 := GenS4{F0: -1}
                                                                                _ = v65
                                                                            }
                                                                        }
                                                                    case -1:
                                                                        fmt.Println(fi52, v2)
                                                                    case 127:
                                                                        if false {
                                                                            var v66 uint64 = 18446744073709551615
                                                                            _ = v66
                                                                        } else {
                                                                            v3 += v3
                                                                            v3 += (v3 * (^(v3)))
                                                                            var v67 string = v53
                                                                            var v68 int32 = (v22 << ((65535) & 63))
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                v2 = 389
                                                            }
                                                        default:
                                                            v69 := GenS4{F0: -779}
                                                            _ = v69
                                                        }
                                                        v22 -= (v22 >> ((42) & 63))
                                                    }
                                                    _ = len(v4)
                                                }
                                            case 44:
                                                fmt.Println(fi52)
                                            default:
                                                if false {
                                                    v3 += (v3 | 1)
                                                    var v70 bool = (v53 < (v53 + (v53 + ("a" + v53))))
                                                    _ = v70
                                                } else {
                                                    switch ("foo bar" + v53) {
                                                    case "foo bar":
                                                        _ = len(v1)
                                                    default:
                                                        if false {
                                                            switch 42 {
                                                            case 274:
                                                                delete(v4, v53)
                                                            case 440:
                                                                var v71 int = fi52
                                                                v72 := GenS1{F0: 18446744073709551615}
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                            if (1.5 == (float64(-2.25) - (1e10 - 0.0))) {
                                                                v4["a"] = true
                                                                v3 += (^(v3))
                                                                var v73 uint64 = 18446744073709551615
                                                                _ = v73
                                                            } else {
                                                                var v74 int8 = 0
                                                                v22 -= (v22 * (v22 % (((v22 << ((1) & 63))) | 1)))
                                                                if false {
                                                                    switch 91 {
                                                                    case 501:
                                                                        v2 = (^(v2))
                                                                    case 998:
                                                                        if true {
                                                                            if (0 >= 42) {
                                                                                if false {
                                                                                    v3 -= 0
                                                                                    var v75 uint16 = 438
                                                                                    _ = v75
                                                                                }
                                                                            } else {
                                                                                var v76 int = (fi52 << ((4294967295) & 63))
                                                                                _ = v76
                                                                            }
                                                                            var v77 uint16 = (v2 >> ((42) & 63))
                                                                            _ = v77
                                                                        } else {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = ((v53 + (v53 + "a")) + v53)
                                                                            }
                                                                            var v78 uint64 = 18446744073709551615
                                                                            _ = v78
                                                                        }
                                                                    case 127:
                                                                        var v79 bool = true
                                                                        _ = len(v4)
                                                                        _ = v79
                                                                    }
                                                                }
                                                                v3 -= v74
                                                                _ = v74
                                                            }
                                                            fmt.Println(v3, v2, v3)
                                                            var v80 uint32 = 1
                                                            switch (v29 >> ((859) & 63)) {
                                                            case 128:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v81 := GenS5{F0: 142, F1: 861}
                                                                    _ = v81
                                                                case "\n\t":
                                                                    v4["\n\t"] = false
                                                                default:
                                                                    fmt.Println(fi52, v22, v3)
                                                                }
                                                            case 146:
                                                                fmt.Println(v29, v22, fi52)
                                                                v82 := v4[(v53 + v53)]
                                                                v83 := GenS3{F0: (fi28 ^ 671), F1: v22}
                                                                v1 = append(v1, fi52)
                                                                switch v53 {
                                                                case "":
                                                                    if v82 {
                                                                        v84 := GenS4{F0: 489}
                                                                        v2 += v2
                                                                        _ = len(v4)
                                                                        _ = v84
                                                                    } else {
                                                                        if (!v82) {
                                                                            var v85 uint8 = 255
                                                                            fmt.Println(v80, v80)
                                                                            var v86 float32 = 0.0
                                                                            _ = v85
                                                                            _ = v86
                                                                        } else {
                                                                            v29 = (v29 ^ v29)
                                                                            v87 := GenS2{F0: 70}
                                                                            fmt.Println(v80)
                                                                            _ = v87
                                                                        }
                                                                        fmt.Println(v29, v22)
                                                                        v1 = append(v1, 603)
                                                                    }
                                                                default:
                                                                    v53 = "foo bar"
                                                                }
                                                                var v88 uint32 = (^(v80))
                                                                _ = v82
                                                                _ = v83
                                                                _ = v88
                                                            case -64:
                                                                v3 = -1
                                                            }
                                                            _ = v80
                                                        } else {
                                                            v89 := GenS5{F0: 2147483647, F1: 0}
                                                            v22 = (-(v22))
                                                            if ((18446744073709551615 > 1) && true) {
                                                                var v90 uint8 = 255
                                                                var v91 uint32 = 1
                                                                if (!true) {
                                                                    var v92 uint64 = 0
                                                                    fmt.Println(v2)
                                                                    v22 -= v22
                                                                    _ = v92
                                                                }
                                                                _ = v90
                                                                _ = v91
                                                            } else {
                                                                if (1 == 127) {
                                                                    v93 := GenS1{F0: 754}
                                                                    _ = v93
                                                                }
                                                                var v94 uint64 = 4294967295
                                                                delete(v4, v53)
                                                                var v95 int32 = v22
                                                                v29 = fi52
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                            _ = v89
                                                        }
                                                    }
                                                    fmt.Println(v29, fi28)
                                                    fmt.Println(v29, v3, v53)
                                                }
                                            }
                                            if false {
                                                var v96 int64 = int64(v3)
                                                v97 := GenS3{F0: (fi28 & -973), F1: genFn0(-1, fi28, 22)}
                                                _ = v96
                                                _ = v97
                                            } else {
                                                var v98 string = v53
                                                v3 = 1
                                                switch (fi28 / ((2147483647) | 1)) {
                                                case -2147483648:
                                                    v22 += v22
                                                default:
                                                    if ((v22 * v22) < genFn0(-32768, fi28, 42)) {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v99 := GenS3{F0: (fi52 | 2147483647), F1: (v22 & (v22 * v22))}
                                                            var v100 uint = 42
                                                            if false {
                                                                _ = len(v1)
                                                                if (0 > (-(v22))) {
                                                                    v29 = (fi28 & (-(fi28)))
                                                                    fmt.Println(fi28, v3)
                                                                    v4["foo bar"] = (true || false)
                                                                    v101 := GenS1{F0: v100}
                                                                    fmt.Println(v53, fi52)
                                                                    _ = v101
                                                                } else {
                                                                    if false {
                                                                        v2 = 255
                                                                        v22 = (^(v22))
                                                                        v100 = v100
                                                                        v2 = (v2 / (((^(v2))) | 1))
                                                                    }
                                                                    _ = len(v4)
                                                                }
                                                            } else {
                                                                if (42 <= -533) {
                                                                    var v102 uint32 = 65535
                                                                    var v103 int32 = genFn0(676, (v29 / (((fi52 | fi28)) | 1)), 0)
                                                                    _ = v102
                                                                    _ = v103
                                                                } else {
                                                                    if (!(0.0 >= 0.1)) {
                                                                        v104 := GenS5{F0: 2147483648, F1: 93}
                                                                        var v105 int = fi28
                                                                        _ = v104
                                                                        _ = v105
                                                                    }
                                                                }
                                                            }
                                                            _ = v100
                                                            _ = v99
                                                        case "":
                                                            var v106 int8 = (^(v3))
                                                            _ = v106
                                                        }
                                                    } else {
                                                        var v107 float64 = 1.5
                                                        _ = v107
                                                    }
                                                }
                                                _ = v98
                                            }
                                            _ = v54
                                        }
                                        if true {
                                            var v108 int8 = (v3 & -1)
                                            _ = v108
                                        }
                                        _ = v53
                                    } else {
                                        switch (fi52 | (fi28 % ((v29) | 1))) {
                                        case -2147483648:
                                            v2 = v2
                                        }
                                        var v109 string = ("a" + "a")
                                        v2 = v2
                                        v3 -= -17
                                        v110 := v4[(v109 + (v109 + v109))]
                                        _ = v109
                                        _ = v110
                                    }
                                    v111 := GenS5{F0: 42, F1: -9223372036854775808}
                                    v112 := GenS4{F0: -9223372036854775808}
                                    _ = fi52
                                    _ = v111
                                    _ = v112
                                }
                            default:
                                for fi113 := 0; fi113 < 11; fi113++ {
                                    v3 -= 42
                                    v114 := GenS6{F0: (fi113 + (fi28 >> ((255) & 63))), F1: (v22 ^ (v22 + (v22 / ((genFn0(-32768, fi28, 0)) | 1)))), F2: 1}
                                    _ = fi113
                                    _ = v114
                                }
                            }
                        }
                        _ = fi28
                        _ = v29
                    }
                } else {
                    var v115 float32 = 0.0
                    _ = v115
                }
            }
            _ = v27
        } else {
            var v116 int16 = -1
            for fi117 := 0; fi117 < 3; fi117++ {
                v1 = append(v1, (-(fi117)))
                _ = fi117
            }
            for fi118 := 0; fi118 < 4; fi118++ {
                switch "" {
                case "\n\t":
                    switch fi118 {
                    case -9223372036854775808:
                        var v119 string = "foo bar"
                        _ = v119
                    case 339:
                        var v120 uint16 = v2
                        _ = v120
                    }
                }
                if false {
                    v116 += (v116 / ((v116) | 1))
                }
                var v121 uint16 = 295
                _ = fi118
                _ = v121
            }
            _ = v116
        }
        for fi122 := 0; fi122 < 10; fi122++ {
            var v123 int = (fi122 >> ((669) & 63))
            v124 := GenS7{F0: true, F1: v22, F2: (v3 + v3)}
            var v125 string = "a"
            _ = fi122
            _ = v123
            _ = v124
            _ = v125
        }
    } else {
        var v126 bool = (42 > 500)
        _ = v126
    }
    fmt.Println(v2, v2)
    fmt.Println(v2, v2)
    var v127 int = 32768
    v3 += v3
    for fi128 := 0; fi128 < 7; fi128++ {
        if false {
            var v129 float64 = -2.25
            _ = len(v1)
            v3 -= v3
            _ = v129
        } else {
            for fi130 := 0; fi130 < 6; fi130++ {
                if false {
                    var v131 int = (v127 - v127)
                    _ = v131
                }
                v132 := GenS1{F0: 4294967295}
                _ = fi130
                _ = v132
            }
            v127 += fi128
        }
        v133 := GenS7{F0: false, F1: (-(v22)), F2: v3}
        _ = fi128
        _ = v133
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v127
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
