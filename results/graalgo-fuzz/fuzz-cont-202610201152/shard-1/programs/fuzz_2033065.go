package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint16; F1 string; F2 uint16 }
type GenS7 struct { F0 string; F1 int; F2 uint64 }
type GenS8 struct { F0 uint32; F1 int64; F2 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 uint; F1 string }
type GenS13 struct { F0 bool; F1 uint8; F2 uint }
type GenS14 struct { F0 bool; F1 int8 }
type GenS15 struct { F0 int64; F1 bool; F2 int8 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int8; F1 bool }
type GenS20 struct { F0 uint8; F1 uint8 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 int; F1 uint16 }
type GenS23 struct { F0 int64 }
func genFn0() uint8 {
    var v5 bool = false
    var v6 int16 = -364
    _ = v5
    _ = v6
    return 239
}
func genFn1(p0 int, p1 uint64, p2 bool) int8 {
    if p2 {
        v7 := []int{p0}
        var v8 uint8 = 103
        _ = v7
        _ = v8
    } else {
        switch p0 {
        case 2147483647:
            var v9 uint16 = 484
            _ = v9
        default:
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS1{F0: 2, F1: true, F2: 1}
                if p2 {
                    fmt.Println(fi10, fi10)
                    if p2 {
                        for fi12 := 0; fi12 < 14; fi12++ {
                            switch (fi12 >> ((758) & 63)) {
                            case -2147483648:
                                var v13 float32 = -2.25
                                var v14 uint64 = (^(p1))
                                v15 := []int{fi12, 42, (fi12 >> ((2) & 63)), int(p1)}
                                _ = v13
                                _ = v14
                                _ = v15
                            default:
                                var v16 bool = (!(((!(!(!(false || false)))) || ((!p2) && (fi10 == -1))) && (!((1 < genFn0()) && (-548 >= 32767)))))
                                _ = v16
                            }
                            var v17 string = "foo bar"
                            v18 := []int{fi12, 32767, p0, -418}
                            _ = fi12
                            _ = v17
                            _ = v18
                        }
                        if false {
                            v19 := GenS1{F0: genFn0(), F1: (0 < -816), F2: 89}
                            _ = v19
                        }
                    } else {
                        v20 := []int{(fi10 | (p0 % ((p0) | 1))), fi10, (p0 + (fi10 >> ((0) & 63)))}
                        _ = v20
                    }
                    p2 = ((p0 - (fi10 >> ((65535) & 63))) <= fi10)
                }
                v21 := GenS1{F0: genFn0(), F1: p2, F2: -128}
                _ = fi10
                _ = v11
                _ = v21
            }
        }
        if (!false) {
            if p2 {
                p0 = -9223372036854775808
            }
            v22 := []int{-1}
            _ = v22
        } else {
            var v23 bool = p2
            fmt.Println(v23, v23, p2)
            _ = v23
        }
    }
    var v24 int8 = 1
    if p2 {
        p1 = p1
        p0 -= 765
        var v25 bool = (!(p2 && (2 > 0)))
        _ = v25
    }
    if false {
        v26 := GenS1{F0: 2, F1: (1e10 >= (1.5 + (1e10 * 1e10))), F2: (^(v24))}
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint64 = p1
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 int32 = -1
                v32 := []int{}
                switch "a" {
                case "\n\t":
                    v24 = v24
                }
                _ = fi30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 11; fi33++ {
                v34 := []int{}
                _ = fi33
                _ = v34
            }
            _ = fi29
        }
        _ = v26
    }
    v24 -= (v24 & 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 127
}

func main() {
    var v1 uint8 = 128
    var v2 bool = (0 < -1)
    var v3 float64 = (0.0 * (-2.25 / -2.25))
    var v4 uint = 70
    fmt.Println(v1, v1)
    v35 := []int{4294967295, 502, 408}
    v36 := GenS1{F0: 0, F1: false, F2: int8(v1)}
    if v2 {
        v4 = (v4 & (v4 >> ((v4) & 63)))
    }
    if true {
        if v2 {
            switch -9223372036854775808 {
            case 74:
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
            case 42:
                v4 -= v4
            case 128:
                var v37 int64 = 42
                _ = v37
            default:
                var v38 bool = (false && (2 >= 42))
                _ = v38
            }
            v3 -= float64(1e10)
        } else {
            v35 = append(v35, 1)
        }
        for fi39 := 0; fi39 < 14; fi39++ {
            switch ("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) {
            case "foo bar":
                var v40 int32 = 0
                _ = v40
            case "\n\t":
                if (!v2) {
                    fmt.Println(fi39, v4)
                    var v41 uint = (v4 + (^(v4)))
                    _ = v41
                } else {
                    v1 = 42
                    if (-82 <= 9223372036854775807) {
                        var v42 float64 = float64(1e10)
                        _ = v42
                    } else {
                        if false {
                            for fi43 := 0; fi43 < 15; fi43++ {
                                v44 := GenS1{F0: v1, F1: v2, F2: 37}
                                _ = fi43
                                _ = v44
                            }
                            var v45 bool = false
                            _ = v45
                        }
                    }
                    var v46 uint = 981
                    _ = v46
                }
            default:
                for fi47 := 0; fi47 < 8; fi47++ {
                    v2 = v2
                    v3 = v3
                    fmt.Println(v4)
                    var v48 int64 = -303
                    var v49 uint32 = 2147483647
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            }
            _ = fi39
        }
        var v50 uint = v4
        _ = v50
    } else {
        fmt.Println(v3, v2, v1)
        for fi51 := 0; fi51 < 1; fi51++ {
            for fi52 := 0; fi52 < 18; fi52++ {
                if true {
                    if (42 != (fi51 | (-(fi52)))) {
                        fmt.Println(v1, v1, v2)
                    } else {
                        if (fi51 < 42) {
                            v35 = append(v35, fi52)
                            _ = len(v35)
                            v4 -= v4
                        } else {
                            v4 += (v4 - v4)
                        }
                        if (-748 != 42) {
                            v53 := GenS1{F0: 255, F1: false, F2: -1}
                            _ = v53
                        }
                        v1 += (v1 ^ v1)
                        if v2 {
                            v3 -= 1.5
                            v1 = (^(v1))
                            var v54 float64 = ((3.14159 - 0.0) / ((((v3 / 1.5) + float64(-2.25)) * (1.5 * (v3 + v3))) - float64((3.14159 / 1.5))))
                            if v2 {
                                v55 := GenS1{F0: 42, F1: v2, F2: -118}
                                if ((3.14159 * float32(((((v54 / v54) + 0.0) / (v3 * -2.25)) - float64(1.5)))) == 1e10) {
                                    v4 -= (v4 / ((v4) | 1))
                                    v56 := GenS2{F0: 1, F1: (v4 << (((v4 + v4)) & 63)), F2: ((false || (2 >= 42)) && (!v2))}
                                    v57 := GenS3{F0: 42}
                                    fmt.Println(fi51, v4, fi51)
                                    if (true && (741 == 42)) {
                                        v58 := GenS4{F0: ("" + ""), F1: 255}
                                        v2 = v2
                                        _ = len(v35)
                                        _ = v58
                                    }
                                    switch ((("foo bar" + ("\n\t" + "hello")) + "a") + "hello") {
                                    case "":
                                        v59 := GenS4{F0: "\n\t", F1: (v1 * v1)}
                                        _ = v59
                                    default:
                                        var v60 uint64 = 0
                                        _ = v60
                                    }
                                    _ = v56
                                    _ = v57
                                }
                                _ = v55
                            } else {
                                switch 127 {
                                case 2147483647:
                                    if v2 {
                                        v4 = (v4 % ((v4) | 1))
                                    }
                                case -838:
                                    v1 -= (^(v1))
                                }
                                v61 := GenS5{F0: 1, F1: (v1 + v1), F2: int64(fi52)}
                                v2 = (!(("a" + "hello") == "hello"))
                                v62 := GenS5{F0: 655, F1: (v1 | uint8(v4)), F2: -1}
                                switch (fi52 ^ fi52) {
                                case 2147483647:
                                    switch "a" {
                                    case "":
                                        var v63 string = (("a" + (("a" + "\n\t") + "")) + ("a" + "a"))
                                        _ = v63
                                    case "a":
                                        fmt.Println(v54, fi52)
                                    default:
                                        v64 := GenS1{F0: genFn0(), F1: (0 > 916), F2: 100}
                                        v65 := GenS1{F0: (v1 & 107), F1: ("a" <= "\n\t"), F2: -104}
                                        v66 := GenS2{F0: 255, F1: v4, F2: false}
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                case 960:
                                    if false {
                                        fmt.Println(v1, fi52)
                                        if len(v35) > 0 {
                                            v35[0] = fi52
                                        }
                                    }
                                default:
                                    v67 := GenS3{F0: 4294967295}
                                    _ = v67
                                }
                                var v68 uint8 = (v1 * v1)
                                switch "a" {
                                case "":
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                case "\n\t":
                                    v69 := GenS3{F0: 0}
                                    _ = v69
                                case "a":
                                    v35 = append(v35, 4294967295)
                                default:
                                    if v2 {
                                        var v70 uint64 = 42
                                        switch "a" {
                                        case "":
                                            var v71 int32 = -1
                                            _ = v71
                                        case "\n\t":
                                            fmt.Println(v4)
                                        case "foo bar":
                                            v35 = append(v35, -1)
                                        default:
                                            v72 := GenS4{F0: "", F1: (v68 >> (((v4 - 0)) & 63))}
                                            _ = v72
                                        }
                                        _ = v70
                                    }
                                }
                                if ((!v2) || v2) {
                                    if v2 {
                                        if (1 != 281) {
                                            var v73 int32 = int32(fi52)
                                            var v74 int64 = -679
                                            var v75 int = (fi52 & (fi52 | fi51))
                                            var v76 bool = false
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                        var v77 uint8 = (v68 - v1)
                                        _ = v77
                                    } else {
                                        v68 = (v1 + v68)
                                        v78 := GenS3{F0: 319}
                                        _ = v78
                                    }
                                    var v79 bool = true
                                    if v2 {
                                        var v80 float64 = v3
                                        var v81 uint32 = 905
                                        var v82 string = "foo bar"
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    } else {
                                        if v2 {
                                            if true {
                                                var v83 uint = 42
                                                var v84 float32 = (-2.25 / 1.5)
                                                _ = v83
                                                _ = v84
                                            }
                                            v68 += (^(v68))
                                            _ = len(v35)
                                        }
                                        v85 := GenS5{F0: 0, F1: 24, F2: -9223372036854775808}
                                        v1 += (v68 / (((v1 & (v68 ^ v1))) | 1))
                                        _ = v85
                                    }
                                    _ = v79
                                } else {
                                    if true {
                                        v4 = v4
                                        fmt.Println(fi52, v54, fi52)
                                    }
                                    if len(v35) > 0 {
                                        v35[0] = (fi52 - (fi51 & -176))
                                    }
                                    var v86 int = 2147483647
                                    var v87 int8 = 42
                                    var v88 float32 = 0.0
                                    _ = len(v35)
                                    _ = v86
                                    _ = v87
                                    _ = v88
                                }
                                _ = v61
                                _ = v62
                                _ = v68
                            }
                            _ = v54
                        } else {
                            var v89 uint8 = v1
                            _ = v89
                        }
                    }
                    if (v2 || (int64(v1) <= 42)) {
                        switch -195 {
                        case -2147483648:
                            v3 = ((v3 * (0.1 / 0.1)) + (0.0 / (v3 - 0.0)))
                        }
                    } else {
                        if (v2 || (1 <= v4)) {
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                            v35 = append(v35, 9223372036854775807)
                            if (((42 > 114) || v2) && (1 > 2)) {
                                var v90 int64 = -403
                                v4 -= 607
                                var v91 bool = v2
                                _ = v90
                                _ = v91
                            }
                            switch "foo bar" {
                            case "hello":
                                v35 = append(v35, (^(fi51)))
                            case "foo bar":
                                fmt.Println(fi52, v3)
                                switch fi51 {
                                case -945:
                                    var v92 float32 = ((0.1 * float32(((3.14159 - v3) - v3))) - float32((v3 + v3)))
                                    _ = v92
                                case 497:
                                    var v93 uint32 = 65535
                                    _ = v93
                                default:
                                    if (float64(-2.25) < (1.5 + 0.0)) {
                                        var v94 uint8 = (^(v1))
                                        _ = v94
                                    }
                                }
                            default:
                                v95 := GenS6{F0: 1, F1: ("" + "foo bar"), F2: 255}
                                _ = v95
                            }
                            v4 -= 180
                        } else {
                            var v96 int16 = 0
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                            _ = v96
                        }
                        v3 += -2.25
                        v97 := GenS7{F0: ("a" + ""), F1: fi51, F2: 0}
                        if (v2 && (1000 != 1000)) {
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        } else {
                            if len(v35) > 0 {
                                v35[0] = (fi52 * 127)
                            }
                            var v98 int32 = -1
                            var v99 uint64 = 97
                            _ = v98
                            _ = v99
                        }
                        _ = v97
                    }
                } else {
                    if v2 {
                        var v100 uint64 = 583
                        switch (fi52 << (((^(v4))) & 63)) {
                        case 2147483647:
                            switch (("foo bar" + "foo bar") + "foo bar") {
                            case "\n\t":
                                if (2147483647 > 1) {
                                    fmt.Println(v1, v4, v1)
                                } else {
                                    fmt.Println(v1)
                                    v1 = (v1 << ((1) & 63))
                                }
                                v101 := GenS4{F0: ("" + (("\n\t" + "\n\t") + "a")), F1: 143}
                                switch (fi51 / ((fi52) | 1)) {
                                case -2147483648:
                                    v35 = append(v35, 96)
                                default:
                                    var v102 uint = (^(v4))
                                    _ = v102
                                }
                                _ = v101
                            case "a":
                                if (v2 && v2) {
                                    v103 := GenS3{F0: v100}
                                    if false {
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                        if ((1.5 > 0.1) || v2) {
                                            if (839 <= 255) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v104 float32 = 0.0
                                                    var v105 uint16 = 42
                                                    _ = v104
                                                    _ = v105
                                                default:
                                                    if ("" <= ("a" + "foo bar")) {
                                                        if v2 {
                                                            v100 += v100
                                                            switch ("" + "hello") {
                                                            case "a":
                                                                v4 = 2
                                                            default:
                                                                var v106 string = ("" + "\n\t")
                                                                _ = v106
                                                            }
                                                        }
                                                    } else {
                                                        v107 := GenS8{F0: uint32(v1), F1: -9223372036854775808, F2: 764}
                                                        _ = v107
                                                    }
                                                }
                                                switch "" {
                                                case "":
                                                    var v108 uint32 = 2
                                                    _ = v108
                                                case "hello":
                                                    if ((^(v4)) < v4) {
                                                        var v109 string = "foo bar"
                                                        v4 = 0
                                                        _ = v109
                                                    }
                                                case "foo bar":
                                                    v110 := GenS2{F0: 2, F1: 639, F2: true}
                                                    _ = v110
                                                }
                                            } else {
                                                v111 := GenS9{F0: int8(v1)}
                                                v112 := GenS10{F0: (v1 << (((v4 - (^(v4)))) & 63))}
                                                if v2 {
                                                    switch fi52 {
                                                    case 2147483647:
                                                        v100 -= 18446744073709551615
                                                    }
                                                    v113 := GenS11{F0: (v4 / ((v4) | 1)), F1: -64}
                                                    _ = v113
                                                } else {
                                                    v114 := GenS12{F0: v4, F1: ""}
                                                    _ = len(v35)
                                                    _ = v114
                                                }
                                                _ = v111
                                                _ = v112
                                            }
                                        } else {
                                            var v115 int64 = -1
                                            fmt.Println(v2)
                                            switch (fi51 << (((v4 * (v4 + v4))) & 63)) {
                                            case -183:
                                                switch "foo bar" {
                                                case "a":
                                                    if ((!v2) && (!(18446744073709551615 == (v4 - v4)))) {
                                                        fmt.Println(v115, v2, v100)
                                                        fmt.Println(v4)
                                                        var v116 int32 = -1
                                                        var v117 int32 = v116
                                                        var v118 uint32 = 1
                                                        switch ((("" + (("\n\t" + ("a" + "foo bar")) + "\n\t")) + "hello") + "foo bar") {
                                                        case "foo bar":
                                                            v4 += (v4 - 4294967295)
                                                        case "\n\t":
                                                            fmt.Println(v118, v3)
                                                        case "a":
                                                            fmt.Println(v118)
                                                        default:
                                                            var v119 uint = v4
                                                            _ = v119
                                                        }
                                                        _ = v116
                                                        _ = v117
                                                        _ = v118
                                                    } else {
                                                        var v120 int16 = 0
                                                        var v121 int32 = 2147483647
                                                        if (0 != 24) {
                                                            v122 := GenS9{F0: 21}
                                                            v123 := GenS13{F0: (2 > 65535), F1: (v1 + (v1 * v1)), F2: (v4 / (((v4 | (v4 / ((42) | 1)))) | 1))}
                                                            if (true && (false || true)) {
                                                                var v124 uint64 = 4294967295
                                                                v1 -= (v1 ^ 128)
                                                                if (v124 >= (v124 & 2)) {
                                                                    if v2 {
                                                                        var v125 uint = 18446744073709551615
                                                                        var v126 uint = v4
                                                                        _ = v125
                                                                        _ = v126
                                                                    }
                                                                }
                                                                _ = v124
                                                            }
                                                            _ = v122
                                                            _ = v123
                                                        }
                                                        _ = v120
                                                        _ = v121
                                                    }
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if ((fi51 * (fi51 - fi51)) <= fi51) {
                                                            if (!(-1 == 100)) {
                                                                v127 := GenS14{F0: (((1e10 + 3.14159) / 3.14159) < -2.25), F1: -93}
                                                                _ = v127
                                                            }
                                                            v100 -= (v100 & (v100 >> ((v4) & 63)))
                                                        }
                                                    case "\n\t":
                                                        var v128 uint16 = 65535
                                                        var v129 int16 = 42
                                                        v130 := GenS15{F0: v115, F1: v2, F2: 68}
                                                        if len(v35) > 0 {
                                                            v35[0] = 365
                                                        }
                                                        _ = v128
                                                        _ = v129
                                                        _ = v130
                                                    case "hello":
                                                        var v131 int64 = (v115 + 9223372036854775807)
                                                        _ = v131
                                                    default:
                                                        switch (fi52 | 32768) {
                                                        case -197:
                                                            v1 += v1
                                                        case 127:
                                                            var v132 string = ("" + (("foo bar" + "hello") + ""))
                                                            _ = v132
                                                        case -933:
                                                            if len(v35) >= 2 {
                                                                v35 = v35[1:len(v35)]
                                                            }
                                                        default:
                                                            v133 := GenS10{F0: (v1 + 255)}
                                                            _ = v133
                                                        }
                                                    }
                                                case "\n\t":
                                                    _ = len(v35)
                                                }
                                            case -408:
                                                fmt.Println(v4, v3)
                                            case -2147483648:
                                                if true {
                                                    var v134 uint32 = 32
                                                    if (!(v4 > (v4 << ((1) & 63)))) {
                                                        v135 := GenS16{F0: int8(v1)}
                                                        v35 = append(v35, (fi51 | (fi51 + -29)))
                                                        v2 = (3.14159 <= 3.14159)
                                                        _ = v135
                                                    }
                                                    _ = v134
                                                } else {
                                                    switch fi52 {
                                                    case 1:
                                                        fmt.Println(v115)
                                                    case 42:
                                                        v1 = (v1 - 1)
                                                    }
                                                    var v136 float64 = ((v3 * ((v3 - 1e10) + v3)) + v3)
                                                    switch (fi52 % (((fi52 | (-(fi52)))) | 1)) {
                                                    case 520:
                                                        var v137 int = fi51
                                                        _ = v137
                                                    case -459:
                                                        v138 := GenS6{F0: 42, F1: ("a" + "hello"), F2: 682}
                                                        _ = v138
                                                    case 848:
                                                        if len(v35) >= 2 {
                                                            v35 = v35[1:len(v35)]
                                                        }
                                                    default:
                                                        v4 = v4
                                                    }
                                                    _ = v136
                                                }
                                                if v2 {
                                                    v3 -= ((float64(float32(float64((1e10 + 1.5)))) / 1.0) / v3)
                                                    v3 -= 1.5
                                                    var v139 float64 = ((v3 / 0.1) * v3)
                                                    _ = v139
                                                } else {
                                                    var v140 string = ""
                                                    _ = v140
                                                }
                                                v141 := GenS17{F0: 32767}
                                                _ = v141
                                            }
                                            _ = v115
                                        }
                                        var v142 float32 = 1.5
                                        _ = v142
                                    } else {
                                        _ = len(v35)
                                        fmt.Println(fi52, v1)
                                        v143 := GenS18{F0: 1}
                                        _ = v143
                                    }
                                    _ = v103
                                }
                            }
                        case 128:
                            var v144 uint = (v4 << ((v4) & 63))
                            _ = v144
                        default:
                            var v145 uint16 = 255
                            _ = v145
                        }
                        _ = v100
                    }
                }
                _ = fi52
            }
            _ = fi51
        }
    }
    var v146 int = -2147483648
    v147 := GenS14{F0: (9223372036854775807 != 35), F1: 42}
    if (0.1 != 3.14159) {
        v3 += 1e10
        for fi148 := 0; fi148 < 16; fi148++ {
            if ((-9223372036854775808 != -971) && false) {
                var v149 float64 = ((v3 / 1e10) - (v3 - v3))
                switch "\n\t" {
                case "foo bar":
                    if (!(!(-321 > 62))) {
                        v4 -= v4
                        v150 := GenS19{F0: 21, F1: true}
                        v151 := GenS17{F0: 169}
                        _ = v150
                        _ = v151
                    }
                case "a":
                    var v152 int32 = -2147483648
                    _ = v152
                case "":
                    if v2 {
                        var v153 uint16 = 387
                        fmt.Println(v1, v4, v149)
                        v154 := GenS5{F0: 1000, F1: (v1 % (((^(v1))) | 1)), F2: -80}
                        _ = v153
                        _ = v154
                    } else {
                        var v155 uint64 = 4294967295
                        if v2 {
                            for fi156 := 0; fi156 < 4; fi156++ {
                                switch "" {
                                case "\n\t":
                                    v157 := GenS20{F0: (v1 * genFn0()), F1: v1}
                                    _ = v157
                                case "a":
                                    v155 = v155
                                }
                                _ = fi156
                            }
                            for fi158 := 0; fi158 < 3; fi158++ {
                                switch "foo bar" {
                                case "hello":
                                    if len(v35) > 0 {
                                        v35[0] = fi148
                                    }
                                case "":
                                    v1 = v1
                                case "\n\t":
                                    if v2 {
                                        v3 -= ((float64(0.1) - ((v149 * float64(1e10)) * v3)) - -2.25)
                                    }
                                default:
                                    if (v2 && (0 < -1)) {
                                        v35 = append(v35, 968)
                                        v149 -= float64((float32(float64(float32(v149))) + (-2.25 + float32(1e10))))
                                        v159 := GenS13{F0: false, F1: genFn0(), F2: (v4 % ((v4) | 1))}
                                        _ = v159
                                    }
                                }
                                if (-64 != 127) {
                                    if false {
                                        v3 = (-2.25 - (v3 / (float64(float32(v3)) + (v3 / 0.1))))
                                    } else {
                                        var v160 uint32 = 910
                                        _ = v160
                                    }
                                    v149 = v3
                                    var v161 uint8 = v1
                                    var v162 int = fi158
                                    v146 -= fi158
                                    _ = v161
                                    _ = v162
                                } else {
                                    v149 += v149
                                }
                                switch (fi148 + -710) {
                                case -9223372036854775808:
                                    switch ("foo bar" + ("\n\t" + "\n\t")) {
                                    case "\n\t":
                                        var v163 int16 = 32767
                                        _ = v163
                                    case "":
                                        v4 -= (v4 * (v4 % ((4294967295) | 1)))
                                    default:
                                        switch (v146 & (v146 ^ (fi158 % ((fi148) | 1)))) {
                                        case -76:
                                            v1 += (v1 << ((27) & 63))
                                            var v164 uint8 = (v1 << ((v4) & 63))
                                            fmt.Println(v149)
                                            _ = v164
                                        default:
                                            if len(v35) >= 2 {
                                                v35 = v35[1:len(v35)]
                                            }
                                        }
                                    }
                                case 32767:
                                    fmt.Println(v2, v149)
                                case -455:
                                    if len(v35) > 0 {
                                        v35[0] = -36
                                    }
                                }
                                v165 := GenS11{F0: v4, F1: -32768}
                                if (!v2) {
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(fi158, v149)
                                    case "":
                                        if len(v35) > 0 {
                                            v35[0] = (fi148 % ((fi148) | 1))
                                        }
                                    }
                                    v166 := GenS17{F0: 255}
                                    if (!(v155 != (v155 << ((42) & 63)))) {
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                    }
                                    _ = v166
                                } else {
                                    var v167 float32 = (1.5 - (3.14159 * 1.5))
                                    v2 = (2147483647 <= 0)
                                    v155 -= 42
                                    v168 := GenS4{F0: "a", F1: (v1 * (^(v1)))}
                                    _ = v167
                                    _ = v168
                                }
                                _ = fi158
                                _ = v165
                            }
                        } else {
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        }
                        _ = v155
                    }
                default:
                    if v2 {
                        v169 := GenS21{F0: 1000}
                        _ = v169
                    } else {
                        var v170 float32 = 0.1
                        var v171 uint = 1
                        var v172 float32 = v170
                        _ = v170
                        _ = v171
                        _ = v172
                    }
                }
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
                _ = v149
            } else {
                fmt.Println(fi148)
                if len(v35) > 0 {
                    v35[0] = fi148
                }
                v173 := GenS22{F0: fi148, F1: 32767}
                _ = v173
            }
            for fi174 := 0; fi174 < 2; fi174++ {
                var v175 int16 = -32768
                v175 = (v175 ^ (^(v175)))
                v176 := GenS23{F0: 9223372036854775807}
                _ = fi174
                _ = v175
                _ = v176
            }
            for fi177 := 0; fi177 < 11; fi177++ {
                v4 = (v4 + (v4 >> ((737) & 63)))
                fmt.Println(fi177, v2, fi177)
                v4 -= v4
                _ = fi177
            }
            _ = fi148
        }
    }
    if ((false || v2) && true) {
        for fi178 := 0; fi178 < 13; fi178++ {
            v1 -= (v1 % (((v1 ^ v1)) | 1))
            if ((0 <= -1) && false) {
                v146 += fi178
            }
            _ = fi178
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v146
    _ = v147
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
