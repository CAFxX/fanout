package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS9 struct { F0 int64; F1 int; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 int32; F1 uint16; F2 int8 }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint; F1 uint }
type GenS18 struct { F0 int16; F1 uint8 }
type GenS19 struct { F0 string; F1 int; F2 int64 }
func genFn0() uint {
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := []int{32768, (fi8 % ((127) | 1)), (fi8 | 302)}
                var v10 int64 = -550
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            var v11 uint32 = 529
            v12 := GenS1{F0: 127, F1: 897}
            _ = v11
            _ = v12
        }
        var v13 uint = 18446744073709551615
        var v14 float32 = 1.5
        _ = fi7
        _ = v13
        _ = v14
    }
    if (false && ("" != ("hello" + "foo bar"))) {
        v15 := GenS1{F0: 1, F1: 556}
        _ = v15
    }
    var v16 int = -2147483648
    _ = v16
    return 2
}
func genFn1() int8 {
    v17 := GenS1{F0: 127, F1: 0}
    switch -919 {
    case 0:
        var v18 uint = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := []int{(fi19 / ((fi19) | 1))}
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int = fi19
            v22 -= (fi21 % ((v22) | 1))
            var v23 int8 = 42
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    var v24 uint64 = 42
    _ = v17
    _ = v24
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int8 = 100
    var v3 int8 = (v2 & (v2 | 1))
    var v4 bool = false
    var v5 uint32 = 2147483647
    var v6 float64 = 3.14159
    if false {
        if (!v4) {
            v6 = 1.5
        }
        var v25 int8 = 42
        _ = v25
    } else {
        var v26 int32 = -1
        if (true || ((-(v26)) == -216)) {
            switch 0 {
            case 32768:
                var v27 int8 = (v3 & v2)
                _ = v27
            default:
                v28 := GenS2{F0: 1}
                _ = v28
            }
            v26 = v26
            for fi29 := 0; fi29 < 3; fi29++ {
                for fi30 := 0; fi30 < 13; fi30++ {
                    v31 := []int{(fi30 - -634), fi30}
                    var v32 uint64 = 18446744073709551615
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                var v33 int8 = genFn1()
                _ = fi29
                _ = v33
            }
        }
        v6 -= float64(-2.25)
        _ = v26
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 uint16 = 805
            var v37 uint64 = 0
            var v38 uint16 = v36
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    var v39 int32 = 369
    v1 = 2147483647
    v5 = (v5 / ((730) | 1))
    for fi40 := 0; fi40 < 7; fi40++ {
        if true {
            var v41 string = "\n\t"
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 float64 = (float64(-2.25) + -2.25)
                if v4 {
                    v39 = v39
                    fmt.Println(v1)
                    v44 := []int{128, fi42}
                    fmt.Println(v2, v6, fi42)
                    var v45 string = v41
                    _ = v44
                    _ = v45
                } else {
                    if (!(v5 >= (^(v5)))) {
                        v46 := GenS3{F0: fi42, F1: 255, F2: 0}
                        _ = v46
                    } else {
                        var v47 uint64 = 363
                        _ = v47
                    }
                    v3 = (v2 / ((v2) | 1))
                }
                _ = fi42
                _ = v43
            }
            var v48 bool = v4
            _ = v41
            _ = v48
        }
        for fi49 := 0; fi49 < 8; fi49++ {
            var v50 uint8 = 127
            if (!(v4 && v4)) {
                if v4 {
                    v6 = ((((float64(-2.25) * 3.14159) * 0.0) / 3.14159) - (-2.25 + -2.25))
                    v51 := []int{fi49, fi40}
                    if v4 {
                        var v52 uint32 = (v1 | (v1 >> ((42) & 63)))
                        v53 := GenS4{F0: v5}
                        if len(v51) >= 2 {
                            v51 = v51[1:len(v51)]
                        }
                        v4 = ((("hello" + "\n\t") + "\n\t") == "a")
                        _ = v52
                        _ = v53
                    } else {
                        switch "" {
                        case "foo bar":
                            switch fi40 {
                            case 32768:
                                v51 = append(v51, fi40)
                            case 127:
                                var v54 uint8 = (v50 >> ((genFn0()) & 63))
                                _ = v54
                            case 4294967295:
                                fmt.Println(fi40)
                            default:
                                switch "a" {
                                case "a":
                                    if ((v50 % ((v50) | 1)) <= v50) {
                                        v6 += v6
                                        if (39 != 812) {
                                            if true {
                                                if (((((v6 / (v6 * (float64(float32(-2.25)) * float64(((-2.25 - ((float32((0.1 / v6)) * float32(1.5)) / 3.14159)) + ((-2.25 / 1e10) - 1e10)))))) - ((v6 + (((1.5 + 3.14159) + 0.0) * -2.25)) - v6)) - (float64(3.14159) / v6)) <= v6) && v4) {
                                                    fmt.Println(v1, v1, v4)
                                                } else {
                                                    v39 += 2147483647
                                                    var v55 int16 = 1000
                                                    v6 = (v6 / v6)
                                                    _ = v55
                                                }
                                                v56 := GenS5{F0: "\n\t"}
                                                _ = v56
                                            } else {
                                                var v57 uint64 = 1
                                                if (!((-(fi40)) != 0)) {
                                                    v58 := GenS6{F0: v57, F1: 255, F2: (v50 - v50)}
                                                    _ = v58
                                                } else {
                                                    v3 -= (v2 - 1)
                                                    if v4 {
                                                        var v59 float32 = (1e10 + float32((float64(1.5) / 3.14159)))
                                                        v60 := GenS7{F0: (v5 - 4294967295), F1: v5}
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        if (345 >= 1) {
                                                            switch fi40 {
                                                            case 127:
                                                                var v61 int16 = 32767
                                                                _ = v61
                                                            case 32768:
                                                                _ = len(v51)
                                                            case -9223372036854775808:
                                                                if false {
                                                                    if (true && v4) {
                                                                        v6 = (v6 - float64(((float32(v6) / 1.5) - (3.14159 * 1.5))))
                                                                    } else {
                                                                        v50 += v50
                                                                        fmt.Println(v39, fi49)
                                                                        var v62 float64 = v6
                                                                        _ = v62
                                                                    }
                                                                    _ = len(v51)
                                                                }
                                                            default:
                                                                if ((-(v39)) <= v39) {
                                                                    if len(v51) >= 2 {
                                                                        v51 = v51[1:len(v51)]
                                                                    }
                                                                    v63 := GenS7{F0: (^(v5)), F1: v1}
                                                                    _ = v63
                                                                } else {
                                                                    var v64 float32 = 1e10
                                                                    v50 += (v50 << ((2) & 63))
                                                                    var v65 uint = genFn0()
                                                                    v51 = append(v51, (fi49 * (fi49 << (((v65 % ((genFn0()) | 1))) & 63))))
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            }
                                                        }
                                                        var v66 int64 = -1
                                                        var v67 uint16 = 1
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                    if v4 {
                                                        v39 -= v39
                                                    } else {
                                                        fmt.Println(v6, v5, v3)
                                                        switch "" {
                                                        case "foo bar":
                                                            switch "" {
                                                            case "\n\t":
                                                                var v68 int8 = genFn1()
                                                                _ = v68
                                                            }
                                                        case "":
                                                            var v69 uint16 = 0
                                                            _ = v69
                                                        case "a":
                                                            var v70 uint32 = v5
                                                            _ = v70
                                                        default:
                                                            switch (fi49 | (-(fi40))) {
                                                            case 979:
                                                                if false {
                                                                    fmt.Println(v1, v57)
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v57
                                            }
                                            if (int16(fi49) < -32768) {
                                                if (4294967295 > v5) {
                                                    switch "a" {
                                                    case "foo bar":
                                                        if len(v51) > 0 {
                                                            v51[0] = 42
                                                        }
                                                    case "hello":
                                                        fmt.Println(fi40)
                                                    }
                                                    var v71 uint32 = v1
                                                    _ = v71
                                                } else {
                                                    v50 += v50
                                                }
                                                var v72 float64 = 0.0
                                                _ = v72
                                            } else {
                                                if (((v50 + v50) <= (^(v50))) && ((v50 / ((42) | 1)) > v50)) {
                                                    var v73 uint8 = (^(v50))
                                                    if len(v51) >= 2 {
                                                        v51 = v51[1:len(v51)]
                                                    }
                                                    _ = v73
                                                }
                                                var v74 int = fi49
                                                v74 -= fi40
                                                _ = v74
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi40, v3, v50)
                                    }
                                default:
                                    v75 := GenS7{F0: 42, F1: (v5 & v5)}
                                    _ = v75
                                }
                            }
                        case "a":
                            v51 = append(v51, fi49)
                        default:
                            if len(v51) > 0 {
                                v51[0] = fi40
                            }
                        }
                        if (v4 || false) {
                            var v76 int64 = -35
                            if len(v51) > 0 {
                                v51[0] = 4294967295
                            }
                            v2 = genFn1()
                            _ = v76
                        } else {
                            if (-910 <= -801) {
                                v39 += (v39 | (v39 & (^(v39))))
                                if len(v51) >= 2 {
                                    v51 = v51[1:len(v51)]
                                }
                                if (-9223372036854775808 >= fi49) {
                                    if (((v3 << ((42) & 63)) < genFn1()) && true) {
                                        var v77 int32 = (^(v39))
                                        _ = v77
                                    } else {
                                        _ = len(v51)
                                        if len(v51) > 0 {
                                            v51[0] = int(v2)
                                        }
                                    }
                                    switch (fi49 - -507) {
                                    case -2147483648:
                                        var v78 int64 = 1
                                        _ = v78
                                    case -433:
                                        if true {
                                            v79 := GenS8{F0: v5, F1: 650, F2: (v3 & v2)}
                                            if false {
                                                fmt.Println(fi40, v2, fi40)
                                                var v80 bool = false
                                                _ = v80
                                            }
                                            v81 := GenS9{F0: 1, F1: fi40, F2: v3}
                                            _ = v79
                                            _ = v81
                                        } else {
                                            var v82 float32 = 1.5
                                            var v83 float32 = (-2.25 + v82)
                                            var v84 uint32 = v5
                                            _ = v82
                                            _ = v83
                                            _ = v84
                                        }
                                    case 195:
                                        switch fi49 {
                                        case -110:
                                            if true {
                                                v4 = v4
                                                if len(v51) > 0 {
                                                    v51[0] = (fi40 | fi49)
                                                }
                                                if v4 {
                                                    switch 127 {
                                                    case 32767:
                                                        _ = len(v51)
                                                    case 304:
                                                        v5 = v1
                                                    case 2147483647:
                                                        if false {
                                                            fmt.Println(v39, fi40)
                                                            fmt.Println(v2, fi49, v39)
                                                        }
                                                    }
                                                }
                                            } else {
                                                if len(v51) >= 2 {
                                                    v51 = v51[1:len(v51)]
                                                }
                                                if ("" < ("hello" + "a")) {
                                                    v39 = (v39 & (v39 / (((-(v39))) | 1)))
                                                    if true {
                                                        v85 := GenS10{F0: 127}
                                                        v86 := GenS9{F0: 0, F1: 0, F2: v2}
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                } else {
                                                    if len(v51) >= 2 {
                                                        v51 = v51[1:len(v51)]
                                                    }
                                                    v87 := GenS11{F0: (v39 >> ((109) & 63)), F1: 255}
                                                    _ = v87
                                                }
                                            }
                                        case 42:
                                            switch (^(fi49)) {
                                            case 386:
                                                switch fi40 {
                                                case 807:
                                                    if v4 {
                                                        if len(v51) >= 2 {
                                                            v51 = v51[1:len(v51)]
                                                        }
                                                    } else {
                                                        switch (fi49 / ((633) | 1)) {
                                                        case -609:
                                                            var v88 bool = ((((v6 / ((0.1 / (v6 / (v6 * v6))) * (v6 / v6))) / (float64(1.5) + -2.25)) - ((v6 - v6) * (1.5 / 0.1))) != v6)
                                                            _ = v88
                                                        }
                                                        if len(v51) >= 2 {
                                                            v51 = v51[1:len(v51)]
                                                        }
                                                        var v89 int32 = (v39 ^ v39)
                                                        _ = v89
                                                    }
                                                default:
                                                    v51 = append(v51, int(v3))
                                                }
                                            case -9223372036854775808:
                                                var v90 uint8 = v50
                                                _ = v90
                                            case 127:
                                                v50 = (^(v50))
                                            }
                                            fmt.Println(v50, fi40)
                                            v1 = (v1 & (v1 + (v5 << ((genFn0()) & 63))))
                                        case 32768:
                                            var v91 int16 = 0
                                            _ = v91
                                        default:
                                            v50 = v50
                                        }
                                    }
                                    fmt.Println(v2, v6, v39)
                                    v50 += (v50 / ((v50) | 1))
                                }
                                var v92 bool = true
                                v93 := GenS7{F0: v1, F1: v5}
                                switch ("\n\t" + "") {
                                case "a":
                                    if v4 {
                                        var v94 int16 = 1
                                        v95 := GenS5{F0: "a"}
                                        _ = v94
                                        _ = v95
                                    }
                                default:
                                    if ((^(v1)) <= 1) {
                                        if false {
                                            var v96 string = "hello"
                                            fmt.Println(fi40, v96, v50)
                                            var v97 uint16 = 623
                                            _ = v96
                                            _ = v97
                                        }
                                        v39 = (v39 ^ (^(v39)))
                                        v50 = (v50 ^ (v50 - v50))
                                    }
                                }
                                v3 -= (v3 + (v3 << ((18446744073709551615) & 63)))
                                v5 -= 959
                                if (65535 < 1) {
                                    if true {
                                        v51 = append(v51, fi49)
                                        if v92 {
                                            if true {
                                                fmt.Println(v39)
                                            }
                                            if (!v4) {
                                                switch (fi49 % ((32768) | 1)) {
                                                case 838:
                                                    switch 32767 {
                                                    case 2147483647:
                                                        _ = len(v51)
                                                    default:
                                                        if (false && (747 < 2)) {
                                                            switch 127 {
                                                            case 42:
                                                                v2 -= v2
                                                            case 0:
                                                                if len(v51) > 0 {
                                                                    v51[0] = (fi40 & (fi49 | 127))
                                                                }
                                                            case 9223372036854775807:
                                                                var v98 string = (("foo bar" + ((("" + (("" + "\n\t") + ("" + "foo bar"))) + "hello") + "")) + (("hello" + "\n\t") + "\n\t"))
                                                                _ = v98
                                                            default:
                                                                v3 -= int8(fi40)
                                                            }
                                                        }
                                                    }
                                                case 42:
                                                    v99 := GenS12{F0: (v39 >> ((genFn0()) & 63)), F1: 874, F2: (v3 & (v3 | genFn1()))}
                                                    _ = v99
                                                }
                                                var v100 uint64 = 2
                                                _ = v100
                                            }
                                            var v101 uint8 = v50
                                            _ = v101
                                        } else {
                                            fmt.Println(v92, v50, v6)
                                        }
                                        var v102 uint8 = (^(v50))
                                        _ = v102
                                    } else {
                                        var v103 float32 = 3.14159
                                        var v104 int16 = -759
                                        _ = v103
                                        _ = v104
                                    }
                                    var v105 float64 = v6
                                    _ = v105
                                }
                                _ = v92
                                _ = v93
                            } else {
                                v51 = append(v51, fi49)
                            }
                        }
                        var v106 uint = 971
                        _ = v106
                    }
                    v2 = (v3 % ((v3) | 1))
                    v4 = (3.14159 != 1.5)
                    _ = v51
                }
                if ((("a" + ("hello" + "hello")) >= "\n\t") || v4) {
                    var v107 uint16 = 566
                    _ = v107
                } else {
                    v108 := []int{fi40, fi40}
                    _ = v108
                }
            } else {
                var v109 string = (((("foo bar" + "hello") + "foo bar") + "foo bar") + ("a" + "a"))
                v110 := []int{(-(fi49)), fi40, (fi40 ^ fi49)}
                v1 += v1
                _ = v109
                _ = v110
            }
            fmt.Println(v6, v3, v50)
            _ = fi49
            _ = v50
        }
        for fi111 := 0; fi111 < 10; fi111++ {
            v1 -= 2147483648
            v112 := []int{fi40}
            _ = fi111
            _ = v112
        }
        _ = fi40
    }
    switch 128 {
    case 442:
        var v113 bool = v4
        _ = v113
    case -2147483648:
        switch -1 {
        case 355:
            v4 = v4
        case 127:
            fmt.Println(v39)
        case 883:
            var v114 uint16 = 0
            _ = v114
        default:
            v115 := GenS13{F0: (v1 << ((4294967295) & 63)), F1: 733}
            _ = v115
        }
    case 2147483647:
        var v116 int = -127
        _ = v116
    }
    switch ("" + ("hello" + "hello")) {
    case "\n\t":
        for fi117 := 0; fi117 < 4; fi117++ {
            switch fi117 {
            case -376:
                var v118 int32 = v39
                _ = v118
            case 1:
                v2 += -1
            case 745:
                var v119 int = 627
                _ = v119
            default:
                if (fi117 < 32767) {
                    if ((v3 & v3) > -115) {
                        v120 := []int{fi117, fi117, fi117}
                        if len(v120) > 0 {
                            v120[0] = (fi117 % ((-1) | 1))
                        }
                        _ = v120
                    } else {
                        v121 := []int{fi117, fi117, (fi117 ^ (fi117 << ((0) & 63)))}
                        fmt.Println(v3, v1)
                        _ = v121
                    }
                    v1 += v5
                    v122 := GenS14{F0: 0}
                    _ = v122
                } else {
                    v1 = v1
                    if false {
                        fmt.Println(v39, fi117)
                        var v123 int32 = 1
                        for fi124 := 0; fi124 < 9; fi124++ {
                            var v125 int32 = v123
                            _ = fi124
                            _ = v125
                        }
                        _ = v123
                    } else {
                        if ((1 > (v3 >> ((18446744073709551615) & 63))) && v4) {
                            v126 := GenS15{F0: int16(fi117)}
                            for fi127 := 0; fi127 < 19; fi127++ {
                                v128 := []int{}
                                _ = fi127
                                _ = v128
                            }
                            var v129 bool = (v1 < (^(v1)))
                            _ = v126
                            _ = v129
                        } else {
                            if v4 {
                                v6 += v6
                            }
                            for fi130 := 0; fi130 < 18; fi130++ {
                                v131 := []int{fi130, (fi117 ^ (^(fi130))), fi117}
                                _ = fi130
                                _ = v131
                            }
                            if (2 < 1) {
                                v132 := GenS3{F0: (fi117 & fi117), F1: 749, F2: 128}
                                v4 = ((v2 & v2) >= v3)
                                _ = v132
                            } else {
                                switch -337 {
                                case 4294967295:
                                    v133 := GenS16{F0: 0}
                                    _ = v133
                                case 32768:
                                    v134 := []int{(fi117 * fi117), (fi117 & fi117), (fi117 * 32767), 2147483647}
                                    _ = v134
                                default:
                                    for fi135 := 0; fi135 < 17; fi135++ {
                                        fmt.Println(v3, v2)
                                        v136 := []int{fi117}
                                        _ = fi135
                                        _ = v136
                                    }
                                }
                                switch fi117 {
                                case -835:
                                    var v137 string = "foo bar"
                                    _ = v137
                                case 9223372036854775807:
                                    for fi138 := 0; fi138 < 19; fi138++ {
                                        v139 := GenS1{F0: -15, F1: 18446744073709551615}
                                        _ = fi138
                                        _ = v139
                                    }
                                default:
                                    if true {
                                        v140 := GenS1{F0: -20, F1: 42}
                                        if false {
                                            var v141 float64 = v6
                                            v142 := []int{fi117, fi117, 127, fi117}
                                            if (1 >= 862) {
                                                var v143 uint8 = 2
                                                var v144 bool = (35 == 32767)
                                                v145 := GenS17{F0: 18446744073709551615, F1: 42}
                                                _ = v143
                                                _ = v144
                                                _ = v145
                                            } else {
                                                v39 += (v39 & (v39 + (v39 << ((genFn0()) & 63))))
                                                v6 = v6
                                                fmt.Println(v1, fi117, v3)
                                            }
                                            _ = v141
                                            _ = v142
                                        } else {
                                            for fi146 := 0; fi146 < 14; fi146++ {
                                                if (("foo bar" <= ("" + ("hello" + "\n\t"))) || (4294967295 > 18446744073709551615)) {
                                                    fmt.Println(v3, v1)
                                                    if v4 {
                                                        v6 -= v6
                                                    } else {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v6 = 1e10
                                                        }
                                                        if false {
                                                            if (v4 || false) {
                                                                fmt.Println(v3)
                                                                var v147 uint16 = 92
                                                                var v148 int16 = 0
                                                                _ = v147
                                                                _ = v148
                                                            }
                                                            v39 -= int32(v2)
                                                            v149 := []int{fi117, -79, 958, fi117}
                                                            _ = v149
                                                        }
                                                        fmt.Println(fi117)
                                                    }
                                                    var v150 uint = 527
                                                    _ = v150
                                                } else {
                                                    var v151 int32 = v39
                                                    _ = v151
                                                }
                                                _ = fi146
                                            }
                                        }
                                        _ = v140
                                    } else {
                                        fmt.Println(fi117, v5)
                                    }
                                }
                            }
                        }
                        v3 = 127
                        for fi152 := 0; fi152 < 4; fi152++ {
                            var v153 int64 = 1
                            v154 := []int{fi152, fi152, -9223372036854775808, (fi152 & (^(fi117)))}
                            _ = fi152
                            _ = v153
                            _ = v154
                        }
                    }
                }
            }
            _ = fi117
        }
    case "hello":
        for fi155 := 0; fi155 < 8; fi155++ {
            if v4 {
                var v156 int = fi155
                v157 := GenS12{F0: v39, F1: 0, F2: (v3 & v2)}
                _ = v156
                _ = v157
            } else {
                var v158 int = -675
                v159 := []int{v158, v158, v158}
                v160 := GenS6{F0: 18446744073709551615, F1: 0, F2: 42}
                _ = v158
                _ = v159
                _ = v160
            }
            v2 -= v3
            for fi161 := 0; fi161 < 15; fi161++ {
                var v162 float32 = ((1.5 + 3.14159) / -2.25)
                v163 := []int{-1, (^(fi161)), 1}
                if (false && v4) {
                    var v164 uint32 = (v5 & (v5 & (v1 ^ (v1 - v5))))
                    _ = v164
                } else {
                    switch "foo bar" {
                    case "a":
                        var v165 int = fi161
                        _ = v165
                    case "hello":
                        var v166 int64 = int64(v2)
                        _ = v166
                    case "foo bar":
                        var v167 int = (fi161 & fi155)
                        _ = v167
                    }
                }
                _ = fi161
                _ = v162
                _ = v163
            }
            _ = fi155
        }
    }
    v168 := GenS18{F0: -530, F1: 127}
    for fi169 := 0; fi169 < 18; fi169++ {
        for fi170 := 0; fi170 < 15; fi170++ {
            v39 = 0
            v4 = false
            if (v4 || true) {
                switch fi170 {
                case 426:
                    if v4 {
                        v1 = (v5 ^ (v5 + v1))
                    } else {
                        switch fi170 {
                        case 109:
                            var v171 float32 = (((1e10 / (-2.25 + 1.5)) - float32(-2.25)) / 0.1)
                            fmt.Println(fi170, v5, fi169)
                            if (-889 <= 955) {
                                if v4 {
                                    fmt.Println(v5)
                                } else {
                                    if v4 {
                                        var v172 int = 4294967295
                                        var v173 bool = v4
                                        _ = v172
                                        _ = v173
                                    } else {
                                        v4 = v4
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v174 int32 = v39
                                            _ = v174
                                        case "hello":
                                            v2 = (v2 * 1)
                                        case "\n\t":
                                            fmt.Println(v1)
                                        default:
                                            v175 := []int{(fi170 << ((910) & 63))}
                                            _ = v175
                                        }
                                        fmt.Println(v6, v2, v4)
                                    }
                                }
                                v176 := []int{(fi170 & (fi170 - fi170)), fi169, fi169, (-(fi170))}
                                v6 += (float64(v171) / float64(float32(1e10)))
                                _ = v176
                            } else {
                                v3 = (v3 << ((1) & 63))
                                v177 := GenS4{F0: 852}
                                _ = v177
                            }
                            _ = v171
                        default:
                            fmt.Println(v2, v2)
                        }
                        v178 := GenS19{F0: "hello", F1: int(v3), F2: 42}
                        var v179 int = fi169
                        if (-9223372036854775808 < 0) {
                            v39 = v39
                        }
                        _ = v178
                        _ = v179
                    }
                }
                var v180 uint32 = (v5 % (((v1 & (v1 ^ (^(v1))))) | 1))
                _ = v180
            } else {
                v181 := []int{}
                _ = v181
            }
            var v182 int16 = 32767
            _ = fi170
            _ = v182
        }
        fmt.Println(fi169)
        v2 = v2
        _ = fi169
    }
    var v183 bool = (18446744073709551615 == 4294967295)
    var v184 float32 = 0.1
    var v185 bool = v4
    v186 := []int{-2147483648, -1, 398, 42}
    var v187 int64 = 596
    for fi188 := 0; fi188 < 8; fi188++ {
        v187 = (^(v187))
        var v189 string = "a"
        _ = fi188
        _ = v189
    }
    for fi190 := 0; fi190 < 4; fi190++ {
        switch "foo bar" {
        case "\n\t":
            var v191 string = ("\n\t" + "hello")
            _ = v191
        case "hello":
            if true {
                v186 = append(v186, fi190)
                for fi192 := 0; fi192 < 4; fi192++ {
                    var v193 uint64 = 0
                    _ = fi192
                    _ = v193
                }
            }
        default:
            v194 := GenS18{F0: -32768, F1: 42}
            _ = v194
        }
        _ = fi190
    }
    fmt.Println("v1:", v1)
    fmt.Println("v183:", v183)
    fmt.Println("v184:", v184)
    fmt.Println("v185:", v185)
    fmt.Println("slicelen v186:", len(v186), v186)
    fmt.Println("v187:", v187)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v168
    _ = v183
    _ = v184
    _ = v185
    _ = v186
    _ = v187
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
