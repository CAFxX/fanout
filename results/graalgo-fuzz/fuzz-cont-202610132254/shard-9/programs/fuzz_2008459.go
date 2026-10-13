package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint; F1 uint }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 string; F1 string; F2 bool }
type GenS11 struct { F0 int64; F1 bool }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    var v4 int8 = -1
    var v5 bool = true
    var v6 float64 = 1.5
    for fi7 := 0; fi7 < 19; fi7++ {
        switch fi7 {
        case 405:
            switch (fi7 / ((fi7) | 1)) {
            case 1:
                if v5 {
                    if (18446744073709551615 > 42) {
                        v1 = 0
                        v8 := GenS1{F0: int16(v1)}
                        _ = v8
                    }
                }
            case 2147483647:
                switch (fi7 - fi7) {
                case -283:
                    v9 := GenS1{F0: 299}
                    _ = v9
                case 32768:
                    v2[("" + ("foo bar" + ((("hello" + "foo bar") + ("\n\t" + "foo bar")) + ("" + ("" + "hello")))))] = (("a" + "\n\t") + (("foo bar" + "a") + "a"))
                default:
                    v6 += ((v6 - 0.0) * 0.1)
                }
            default:
                switch "\n\t" {
                case "":
                    var v10 uint64 = 18446744073709551615
                    _ = v10
                case "foo bar":
                    if ((2 == (v1 / ((v1) | 1))) || (1 != -32768)) {
                        v5 = v5
                        for fi11 := 0; fi11 < 15; fi11++ {
                            if ((float32(v6) > ((1.5 - 0.1) * 1.5)) || true) {
                                switch ("\n\t" + "hello") {
                                case "":
                                    if (("hello" > "hello") && v5) {
                                        var v12 float32 = -2.25
                                        switch 797 {
                                        case 2147483647:
                                            var v13 bool = true
                                            _ = v13
                                        case 4294967295:
                                            switch (("hello" + ("\n\t" + "a")) + ((("\n\t" + "\n\t") + "a") + "\n\t")) {
                                            case "foo bar":
                                                v5 = (((!v5) || v5) && (!v5))
                                            case "hello":
                                                v5 = v5
                                                var v14 uint16 = v1
                                                _ = v14
                                            default:
                                                v15 := GenS2{F0: "foo bar", F1: -2147483648, F2: 128}
                                                _ = v15
                                            }
                                        case -799:
                                            var v16 float64 = float64(((1e10 + -2.25) + (0.0 - 3.14159)))
                                            _ = v16
                                        }
                                        v6 = v6
                                        var v17 float64 = float64(v12)
                                        _ = v12
                                        _ = v17
                                    }
                                case "hello":
                                    v18 := []int{(-(fi11)), 35, 2147483647, (fi7 % ((-1) | 1))}
                                    _ = v18
                                }
                            } else {
                                fmt.Println(v1)
                            }
                            if ("" <= "foo bar") {
                                var v19 bool = (0 > 1)
                                switch fi7 {
                                case 128:
                                    fmt.Println(v4)
                                    var v20 uint8 = 0
                                    var v21 uint32 = 2147483648
                                    v6 -= v6
                                    _ = v20
                                    _ = v21
                                case 524:
                                    if v5 {
                                        var v22 uint32 = 2
                                        v1 += (v1 ^ 1)
                                        _ = v22
                                    }
                                case -2147483648:
                                    v19 = true
                                }
                                var v23 string = ""
                                _ = v19
                                _ = v23
                            }
                            fmt.Println(fi7)
                            _ = fi11
                        }
                        for fi24 := 0; fi24 < 8; fi24++ {
                            if (!v5) {
                                var v25 float32 = 0.0
                                v6 -= (v6 - float64(float32(((((((v6 / v6) - v6) + 0.1) - v6) - ((v6 / 1.5) + v6)) + v6))))
                                _ = v25
                            }
                            var v26 float64 = 1e10
                            _ = fi24
                            _ = v26
                        }
                    } else {
                        var v27 uint = 481
                        if false {
                            var v28 string = ("a" + "\n\t")
                            var v29 int = fi7
                            for fi30 := 0; fi30 < 14; fi30++ {
                                fmt.Println(v29, v5, fi7)
                                v5 = v5
                                v1 -= 32767
                                _ = fi30
                            }
                            _ = v28
                            _ = v29
                        } else {
                            var v31 int16 = -32768
                            v4 = v4
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 8; fi32++ {
                            fmt.Println(v5)
                            v33 := []int{fi32, 953, fi7, (fi7 ^ fi7)}
                            _ = fi32
                            _ = v33
                        }
                        _ = v27
                    }
                default:
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v6 = v6
                        v35 := []int{(fi34 >> ((4294967295) & 63)), (fi7 | -1), (fi34 & 79), (fi7 & fi7)}
                        v1 += (^(v1))
                        fmt.Println(v6, fi34)
                        if (!(v1 <= v1)) {
                            switch "foo bar" {
                            case "":
                                var v36 int32 = 604
                                _ = v36
                            }
                            fmt.Println(v6, fi34)
                        }
                        v4 = v4
                        fmt.Println(fi7)
                        _ = fi34
                        _ = v35
                    }
                }
            }
        case 32767:
            for fi37 := 0; fi37 < 9; fi37++ {
                v38 := []int{(fi7 >> ((1) & 63))}
                var v39 float64 = float64(-2.25)
                _ = fi37
                _ = v38
                _ = v39
            }
        case 127:
            for fi40 := 0; fi40 < 2; fi40++ {
                v5 = v5
                _ = fi40
            }
        default:
            switch fi7 {
            case 2147483647:
                var v41 uint16 = (v1 / ((715) | 1))
                _ = v41
            case 576:
                v6 += v6
            }
        }
        var v42 int16 = -32768
        _ = fi7
        _ = v42
    }
    v6 = (((3.14159 + (v6 - 1e10)) - v6) - v6)
    var v43 string = ("\n\t" + ("\n\t" + ("hello" + "hello")))
    fmt.Println(v43)
    switch -818 {
    case -462:
        var v44 uint = 65535
        _ = v44
    case 2147483647:
        if (((4294967295 == 18446744073709551615) && (!false)) && v5) {
            switch ((v43 + v43) + v43) {
            case "foo bar":
                if ((true || ((false && false) && (((3.14159 * 0.1) > 1e10) || (!false)))) || (true && v5)) {
                    v43 = (v43 + "")
                    v5 = ((v5 || true) || (v5 && (!v5)))
                } else {
                    for fi45 := 0; fi45 < 5; fi45++ {
                        for fi46 := 0; fi46 < 17; fi46++ {
                            v4 = v4
                            var v47 float64 = v6
                            _ = fi46
                            _ = v47
                        }
                        v4 = (v4 % ((v4) | 1))
                        _ = fi45
                    }
                    var v48 float32 = -2.25
                    fmt.Println(v6)
                    _ = v48
                }
            }
        }
    case 0:
        switch 32767 {
        case 2147483647:
            var v49 int64 = 42
            _ = v49
        }
    default:
        if (v6 == 1.5) {
            if (!false) {
                if (0.0 >= (3.14159 - (((0.1 - 1.5) * (0.1 / (1e10 + (0.1 + 0.0)))) * (0.0 + (-2.25 * (0.0 + 3.14159)))))) {
                    if v5 {
                        v5 = v5
                        var v50 int16 = 606
                        _ = v50
                    }
                    _ = len(v2)
                    if v5 {
                        if (-1 > 1) {
                            delete(v2, ((v43 + ((("\n\t" + ((v43 + ("hello" + "")) + v43)) + ("hello" + v43)) + v43)) + (v43 + v43)))
                            if v5 {
                                for fi51 := 0; fi51 < 11; fi51++ {
                                    var v52 int32 = 1
                                    var v53 float32 = 3.14159
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                            }
                        }
                        if (108 == 1) {
                            var v54 uint8 = 127
                            v55 := []int{}
                            v4 += -2
                            _ = v54
                            _ = v55
                        } else {
                            switch 0 {
                            case 767:
                                var v56 int8 = 42
                                _ = v56
                            default:
                                v5 = v5
                            }
                            var v57 uint64 = 669
                            _ = v57
                        }
                        fmt.Println(v5, v5)
                    } else {
                        for fi58 := 0; fi58 < 19; fi58++ {
                            var v59 bool = (!(42 <= 759))
                            v5 = v5
                            _ = fi58
                            _ = v59
                        }
                    }
                }
                if v5 {
                    fmt.Println(v43)
                    v60 := []int{42, 899}
                    delete(v3, 2)
                    _ = v60
                }
                for fi61 := 0; fi61 < 2; fi61++ {
                    for fi62 := 0; fi62 < 9; fi62++ {
                        var v63 int16 = -32768
                        _ = fi62
                        _ = v63
                    }
                    if (true || true) {
                        var v64 bool = (2147483647 >= 614)
                        _ = v64
                    } else {
                        for fi65 := 0; fi65 < 13; fi65++ {
                            if (127 == 2) {
                                var v66 bool = false
                                v67 := v3[1]
                                if v66 {
                                    if v66 {
                                        var v68 int8 = (v4 % ((99) | 1))
                                        fmt.Println(fi61)
                                        _ = v68
                                    } else {
                                        v69 := v3[4294967295]
                                        switch (fi61 - (fi65 - fi61)) {
                                        case 32767:
                                            fmt.Println(v67)
                                        case 0:
                                            var v70 float32 = float32(v6)
                                            _ = v70
                                        }
                                        _ = v69
                                    }
                                    switch fi61 {
                                    case 32767:
                                        v6 += (v6 / float64(float32(float64(((1.5 - -2.25) / 0.1)))))
                                    }
                                    var v71 int64 = 653
                                    var v72 int32 = -25
                                    var v73 uint16 = v1
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                                _ = v66
                                _ = v67
                            } else {
                                var v74 bool = v5
                                var v75 uint16 = (v1 << ((255) & 63))
                                delete(v3, 65535)
                                var v76 uint32 = 4294967295
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                            v77 := []int{}
                            _ = fi65
                            _ = v77
                        }
                        v78 := GenS2{F0: v43, F1: -2147483648, F2: 127}
                        _ = v78
                    }
                    _ = fi61
                }
            } else {
                for fi79 := 0; fi79 < 3; fi79++ {
                    v80 := GenS3{F0: v5, F1: 18446744073709551615}
                    fmt.Println(fi79, fi79, v1)
                    _ = fi79
                    _ = v80
                }
                for fi81 := 0; fi81 < 10; fi81++ {
                    switch "" {
                    case "a":
                        v5 = true
                    }
                    var v82 float64 = v6
                    var v83 int64 = -9223372036854775808
                    _ = fi81
                    _ = v82
                    _ = v83
                }
                for fi84 := 0; fi84 < 15; fi84++ {
                    var v85 bool = (v5 || v5)
                    fmt.Println(v6)
                    v85 = v5
                    _ = fi84
                    _ = v85
                }
            }
            v86 := GenS1{F0: -619}
            var v87 uint16 = v1
            _ = v86
            _ = v87
        }
    }
    v88 := GenS4{F0: -128}
    if (v6 >= v6) {
        v89 := GenS5{F0: -302, F1: 0}
        _ = v89
    } else {
        v6 -= (3.14159 - ((3.14159 - (v6 / v6)) / (1e10 * v6)))
        v90 := GenS2{F0: v43, F1: -2147483648, F2: 149}
        v91 := GenS6{F0: 1000, F1: 325, F2: -801}
        _ = v90
        _ = v91
    }
    _ = len(v2)
    v6 -= 3.14159
    switch ("\n\t" + ((v43 + v43) + v43)) {
    case "":
        v92 := []int{42}
        _ = v92
    }
    var v93 uint = 326
    if v5 {
        v94 := GenS6{F0: -32768, F1: 65535, F2: 904}
        v95 := GenS7{F0: -737, F1: v4}
        switch -1 {
        case 1:
            switch -9223372036854775808 {
            case -829:
                v5 = (1 == -751)
            case 1:
                var v96 int32 = 42
                _ = v96
            case 2147483647:
                var v97 uint16 = v1
                _ = v97
            }
        case -1:
            v1 = (v1 + v1)
        case 128:
            for fi98 := 0; fi98 < 15; fi98++ {
                switch fi98 {
                case 4294967295:
                    var v99 string = "foo bar"
                    _ = v99
                case 128:
                    var v100 float64 = float64(3.14159)
                    _ = v100
                case -743:
                    v101 := []int{fi98, fi98, 424, (^(fi98))}
                    _ = v101
                default:
                    v3[759] = (0.0 >= -2.25)
                }
                _ = fi98
            }
        default:
            var v102 string = v43
            _ = v102
        }
        _ = v94
        _ = v95
    } else {
        var v103 int32 = -1
        _ = v103
    }
    for fi104 := 0; fi104 < 5; fi104++ {
        if v5 {
            for fi105 := 0; fi105 < 5; fi105++ {
                delete(v2, "foo bar")
                if true {
                    v106 := GenS8{F0: v93, F1: 700}
                    var v107 int8 = 42
                    _ = v106
                    _ = v107
                } else {
                    var v108 float64 = 0.0
                    fmt.Println(v6, v1)
                    _ = v108
                }
                if v5 {
                    var v109 uint8 = 51
                    v1 = v1
                    _ = v109
                }
                _ = fi105
            }
            var v110 float64 = float64(1e10)
            v6 += (v110 * v6)
            _ = v110
        }
        if v5 {
            var v111 int16 = -32768
            _ = v111
        }
        switch fi104 {
        case 687:
            for fi112 := 0; fi112 < 6; fi112++ {
                v1 += (v1 << ((779) & 63))
                v113 := []int{}
                v114 := GenS6{F0: 101, F1: (v93 * 2), F2: 916}
                if ((!v5) || (!v5)) {
                    v115 := GenS9{F0: fi104, F1: (^(v4))}
                    if (!(336 <= 1)) {
                        v1 = (v1 ^ (v1 + (^(v1))))
                        v1 = 1
                    }
                    _ = v115
                }
                v4 = 100
                _ = fi112
                _ = v113
                _ = v114
            }
        case 112:
            if v5 {
                switch ((((((("hello" + v43) + v43) + v43) + (("\n\t" + v43) + v43)) + v43) + (("\n\t" + v43) + v43)) + v43) {
                case "foo bar":
                    for fi116 := 0; fi116 < 6; fi116++ {
                        v117 := GenS6{F0: 1, F1: 0, F2: -1}
                        _ = fi116
                        _ = v117
                    }
                default:
                    for fi118 := 0; fi118 < 3; fi118++ {
                        var v119 int8 = (v4 ^ v4)
                        _ = fi118
                        _ = v119
                    }
                }
            } else {
                var v120 int8 = (v4 % ((int8(v1)) | 1))
                v120 = (v120 / (((-(v120))) | 1))
                fmt.Println(v43, v93)
                _ = v120
            }
        default:
            switch "hello" {
            case "foo bar":
                if v5 {
                    var v121 float64 = (v6 / 1.5)
                    v122 := v3[0]
                    for fi123 := 0; fi123 < 13; fi123++ {
                        fmt.Println(v122, v4)
                        v124 := []int{(-(fi123)), fi104, -1}
                        _ = fi123
                        _ = v124
                    }
                    _ = v121
                    _ = v122
                } else {
                    v93 -= v93
                    switch 32767 {
                    case 0:
                        fmt.Println(v43)
                    case 32767:
                        switch fi104 {
                        case -952:
                            v125 := []int{}
                            _ = v125
                        }
                    case -1:
                        switch ("" + v43) {
                        case "a":
                            fmt.Println(v93, v4)
                        case "":
                            v126 := GenS4{F0: v4}
                            _ = v126
                        case "foo bar":
                            v127 := GenS10{F0: v43, F1: v43, F2: v5}
                            _ = v127
                        default:
                            var v128 float64 = (float64(0.0) / ((v6 / v6) - v6))
                            _ = v128
                        }
                    default:
                        if (-9223372036854775808 > -775) {
                            v43 = "foo bar"
                            v129 := GenS2{F0: "foo bar", F1: -3, F2: 42}
                            v130 := GenS11{F0: 42, F1: (335 > 0)}
                            _ = v129
                            _ = v130
                        } else {
                            var v131 int64 = -9223372036854775808
                            v132 := GenS7{F0: 1, F1: -89}
                            if v5 {
                                v133 := GenS3{F0: (!(575 != 4294967295)), F1: (v93 | 255)}
                                switch fi104 {
                                case 2147483647:
                                    v93 = (v93 + (v93 % ((0) | 1)))
                                case 127:
                                    var v134 uint16 = v1
                                    _ = v134
                                }
                                v135 := GenS5{F0: 0, F1: v131}
                                _ = v133
                                _ = v135
                            }
                            _ = v131
                            _ = v132
                        }
                    }
                    var v136 int64 = 9223372036854775807
                    _ = v136
                }
            case "a":
                v4 -= v4
            }
        }
        _ = fi104
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v88
    _ = v93
}
