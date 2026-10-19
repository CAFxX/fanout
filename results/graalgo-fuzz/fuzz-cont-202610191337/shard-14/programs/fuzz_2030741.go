package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 bool }
type GenS8 struct { F0 uint16; F1 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint; F1 int64 }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint64; F1 int64; F2 int }
type GenS15 struct { F0 uint32; F1 int }
type GenS16 struct { F0 uint32; F1 uint64 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int64; F1 int }
type GenS19 struct { F0 int16; F1 string }
type GenS20 struct { F0 uint32; F1 uint8 }
type GenS21 struct { F0 int32; F1 int64 }
type GenS22 struct { F0 bool; F1 bool; F2 uint32 }
type GenS23 struct { F0 int32; F1 uint64; F2 int8 }
type GenS24 struct { F0 int }
type GenS25 struct { F0 int32; F1 uint16 }
type GenS26 struct { F0 uint16; F1 int8; F2 bool }
type GenS27 struct { F0 uint32 }
type GenS28 struct { F0 int16; F1 int64; F2 uint }
type GenS29 struct { F0 uint64; F1 int32; F2 int16 }
type GenS30 struct { F0 int8 }
type GenS31 struct { F0 int32; F1 int64 }
func genFn0(p0 int, p1 uint8, p2 int16) int32 {
    if (!((!(false && true)) || (-30 >= -695))) {
        v5 := GenS1{F0: (p0 >> ((18446744073709551615) & 63)), F1: 1}
        if false {
            switch "" {
            case "":
                p1 -= p1
            default:
                p0 += -2147483648
            }
            if false {
                p0 += p0
                if true {
                    var v6 uint = 255
                    var v7 bool = true
                    _ = v6
                    _ = v7
                }
                v8 := GenS1{F0: 1, F1: int8(p0)}
                _ = v8
            } else {
                var v9 float64 = (0.1 / 1.0)
                _ = v9
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                p2 = p2
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := GenS2{F0: (p1 * 127)}
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
        _ = v5
    }
    var v13 uint16 = 0
    v14 := GenS1{F0: (p0 + p0), F1: -47}
    fmt.Println(p2, v13)
    for fi15 := 0; fi15 < 4; fi15++ {
        v13 -= (v13 | v13)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int = -720
    var v16 int32 = genFn0(v4, 255, 42)
    var v17 float64 = 0.1
    v4 -= (v4 % (((v4 >> ((65535) & 63))) | 1))
    var v18 string = "a"
    for fi19 := 0; fi19 < 10; fi19++ {
        if false {
            var v20 bool = (true || false)
            _ = v20
        } else {
            var v21 int16 = -190
            v21 += (v21 - v21)
            _ = v21
        }
        if false {
            v22 := GenS3{F0: 666, F1: 127}
            _ = v22
        }
        _ = fi19
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(fi23, v16, v18)
        v17 = 3.14159
        _ = len(v1)
        _ = fi23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        switch v18 {
        case "hello":
            var v25 uint32 = 65535
            _ = v25
        case "\n\t":
            var v26 int64 = 42
            _ = v26
        case "foo bar":
            if (((0 >= 18446744073709551615) || false) || false) {
                _ = len(v2)
                v16 = (v16 / ((42) | 1))
                if (("" + ("a" + v18)) == v18) {
                    if true {
                        if true {
                            switch "foo bar" {
                            case "a":
                                for fi27 := 0; fi27 < 6; fi27++ {
                                    if true {
                                        switch (v18 + "") {
                                        case "hello":
                                            if true {
                                                _ = len(v3)
                                                var v28 uint32 = 2
                                                var v29 bool = (("hello" + v18) < v18)
                                                _ = v28
                                                _ = v29
                                            }
                                        case "":
                                            if (true && true) {
                                                v30 := GenS2{F0: 127}
                                                v17 = float64(-2.25)
                                                var v31 uint8 = 235
                                                _ = v30
                                                _ = v31
                                            }
                                        case "foo bar":
                                            var v32 int64 = -1
                                            _ = v32
                                        default:
                                            fmt.Println(v18, fi27)
                                        }
                                        if true {
                                            if true {
                                                var v33 int32 = 0
                                                if (!(false || true)) {
                                                    v34 := []int{(fi24 * v4), (v4 << ((278) & 63))}
                                                    _ = v34
                                                }
                                                _ = len(v2)
                                                _ = v33
                                            } else {
                                                v16 -= v16
                                                switch fi27 {
                                                case 127:
                                                    if (v16 <= genFn0((fi24 - v4), 0, 1)) {
                                                        if ((v17 > float64((1.5 * 0.1))) && true) {
                                                            v3[42] = (32767 > 65535)
                                                            v35 := GenS3{F0: 980, F1: 0}
                                                            _ = v35
                                                        } else {
                                                            if (v4 <= fi24) {
                                                                if ((true || true) || (3.14159 <= v17)) {
                                                                    v36 := GenS4{F0: (2 == 0)}
                                                                    var v37 int16 = 1
                                                                    fmt.Println(v16)
                                                                    _ = v36
                                                                    _ = v37
                                                                } else {
                                                                    var v38 bool = true
                                                                    v39 := []int{}
                                                                    var v40 float64 = v17
                                                                    _ = v38
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                v41 := GenS4{F0: (fi24 < (^(fi24)))}
                                                                _ = v41
                                                            }
                                                            var v42 float64 = v17
                                                            fmt.Println(fi27, v17, v42)
                                                            v43 := GenS5{F0: 42}
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                    } else {
                                                        v44 := []int{32768, (-(fi27)), (v4 | 1)}
                                                        v45 := v3[2147483647]
                                                        if (!(18446744073709551615 > 18446744073709551615)) {
                                                            v4 -= -667
                                                            if (129 >= 2) {
                                                                fmt.Println(fi27, v17, v45)
                                                            }
                                                        }
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                case 4294967295:
                                                    var v46 uint64 = 117
                                                    v47 := GenS1{F0: (fi27 | (v4 >> ((18446744073709551615) & 63))), F1: 43}
                                                    v17 -= v17
                                                    _ = v46
                                                    _ = v47
                                                }
                                                if true {
                                                    v48 := GenS6{F0: 0}
                                                    v49 := GenS3{F0: 2147483647, F1: 49}
                                                    var v50 int16 = 359
                                                    var v51 float32 = -2.25
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                            v18 = (v18 + (v18 + ("hello" + "")))
                                            if (!true) {
                                                fmt.Println(v17, v16, v17)
                                            } else {
                                                v52 := []int{v4, fi27, (fi27 + v4), (fi24 * 1)}
                                                var v53 bool = false
                                                if (624 != 999) {
                                                    v54 := GenS7{F0: 2, F1: 234, F2: (764 != -1)}
                                                    v18 = v18
                                                    switch v4 {
                                                    case -2147483648:
                                                        _ = len(v1)
                                                    case 42:
                                                        if len(v52) >= 2 {
                                                            v52 = v52[1:len(v52)]
                                                        }
                                                        var v55 uint64 = 4294967295
                                                        _ = v55
                                                    default:
                                                        v18 = ""
                                                    }
                                                    _ = v54
                                                }
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                    }
                                    v56 := GenS8{F0: 751, F1: v16}
                                    _ = fi27
                                    _ = v56
                                }
                            case "hello":
                                for fi57 := 0; fi57 < 3; fi57++ {
                                    var v58 uint8 = 1
                                    switch v4 {
                                    case 127:
                                        var v59 float32 = ((3.14159 * float32((((v17 * v17) + v17) - v17))) - 0.0)
                                        _ = v59
                                    case -382:
                                        v2[9223372036854775807] = 1
                                    case -1:
                                        v17 = v17
                                    default:
                                        var v60 float64 = v17
                                        _ = v60
                                    }
                                    v61 := GenS9{F0: -113}
                                    _ = fi57
                                    _ = v58
                                    _ = v61
                                }
                            case "\n\t":
                                fmt.Println(v17)
                            default:
                                if false {
                                    var v62 uint16 = 2
                                    v63 := []int{v4, 42}
                                    v64 := GenS5{F0: 65535}
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                } else {
                                    switch -694 {
                                    case -9223372036854775808:
                                        if true {
                                            for fi65 := 0; fi65 < 18; fi65++ {
                                                v66 := []int{v4, (fi65 >> ((255) & 63))}
                                                delete(v2, 32767)
                                                var v67 int32 = 940
                                                fmt.Println(v18, v17, v67)
                                                v68 := GenS1{F0: (fi24 / (((fi65 | fi65)) | 1)), F1: 72}
                                                _ = fi65
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            }
                                            for fi69 := 0; fi69 < 5; fi69++ {
                                                fmt.Println(v17, v16)
                                                _ = fi69
                                            }
                                            var v70 uint64 = 0
                                            _ = v70
                                        }
                                    default:
                                        v71 := GenS3{F0: 42, F1: 1}
                                        _ = v71
                                    }
                                    fmt.Println(v16)
                                    if true {
                                        switch v18 {
                                        case "\n\t":
                                            v72 := GenS8{F0: 815, F1: v16}
                                            _ = v72
                                        case "a":
                                            var v73 float64 = -2.25
                                            _ = v73
                                        case "":
                                            v74 := GenS2{F0: 226}
                                            _ = v74
                                        default:
                                            for fi75 := 0; fi75 < 16; fi75++ {
                                                if ((-32768 >= 32767) && ((0.1 - 3.14159) != 3.14159)) {
                                                    fmt.Println(v18, v18)
                                                    switch (fi24 >> ((255) & 63)) {
                                                    case 0:
                                                        if (v4 >= (fi24 ^ (v4 % ((-1) | 1)))) {
                                                            if false {
                                                                switch v18 {
                                                                case "\n\t":
                                                                    v16 = (v16 >> ((255) & 63))
                                                                }
                                                            } else {
                                                                v17 += (float64(3.14159) * ((v17 * -2.25) / (-2.25 - float64(((1e10 - float32(((3.14159 - float64(3.14159)) + 0.1))) + (float32(-2.25) - 0.1))))))
                                                            }
                                                            v76 := GenS2{F0: 127}
                                                            v77 := GenS10{F0: 710, F1: 795}
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    case -840:
                                                        delete(v1, v18)
                                                    case 1:
                                                        if true {
                                                            var v78 float32 = 3.14159
                                                            v79 := []int{}
                                                            _ = len(v79)
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            fmt.Println(v16, v18, v18)
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    var v80 int16 = 1000
                                                    v4 += (fi75 % (((v4 & fi24)) | 1))
                                                    _ = v80
                                                } else {
                                                    if (false || (!true)) {
                                                        v17 = v17
                                                        v81 := []int{v4}
                                                        if true {
                                                            v82 := GenS1{F0: 2147483647, F1: 45}
                                                            _ = v82
                                                        } else {
                                                            v83 := v3[42]
                                                            v17 = ((v17 - float64(3.14159)) + v17)
                                                            _ = v83
                                                        }
                                                        _ = v81
                                                    }
                                                }
                                                _ = fi75
                                            }
                                        }
                                    } else {
                                        v4 = (-(v4))
                                        v84 := []int{}
                                        var v85 uint32 = 2147483648
                                        _ = v84
                                        _ = v85
                                    }
                                }
                            }
                        }
                        switch fi24 {
                        case 32768:
                            v86 := []int{(v4 | 0), fi24}
                            _ = v86
                        case 27:
                            fmt.Println(v18)
                        case 1:
                            if true {
                                v16 -= genFn0(605, 127, -9)
                            } else {
                                v87 := []int{(v4 | 2147483647), 0, fi24}
                                switch (fi24 | fi24) {
                                case 42:
                                    for fi88 := 0; fi88 < 17; fi88++ {
                                        if true {
                                            fmt.Println(v17, fi24)
                                            if true {
                                                var v89 uint16 = 1
                                                var v90 string = v18
                                                if (1 > 68) {
                                                    v17 += 0.0
                                                } else {
                                                    var v91 int64 = 9223372036854775807
                                                    v92 := GenS11{F0: 85, F1: 717}
                                                    switch v90 {
                                                    case "":
                                                        var v93 int16 = 1000
                                                        _ = v93
                                                    case "\n\t":
                                                        if (!false) {
                                                            switch 9223372036854775807 {
                                                            case 0:
                                                                v94 := GenS7{F0: 2147483647, F1: 127, F2: (42 == 2)}
                                                                _ = v94
                                                            default:
                                                                switch v18 {
                                                                case "hello":
                                                                    v17 += (((1.5 * float64(1.5)) * (float64(((float32(0.1) - 0.1) + (-2.25 - (3.14159 + 0.1)))) / (0.0 + (1e10 * 0.1)))) + v17)
                                                                case "a":
                                                                    if true {
                                                                        if ((v91 / ((v91) | 1)) <= v91) {
                                                                            if (true && (!(4294967295 < 2))) {
                                                                                var v95 uint64 = 338
                                                                                _ = v95
                                                                            }
                                                                            var v96 string = v18
                                                                            v4 += (fi24 ^ (v4 + (v4 ^ (^(fi24)))))
                                                                            _ = v96
                                                                        }
                                                                        v97 := GenS12{F0: ""}
                                                                        _ = v97
                                                                    }
                                                                default:
                                                                    if len(v87) > 0 {
                                                                        v87[0] = v4
                                                                    }
                                                                }
                                                            }
                                                            switch (fi24 & -588) {
                                                            case 9223372036854775807:
                                                                switch fi88 {
                                                                case 32767:
                                                                    var v98 uint8 = 128
                                                                    _ = v98
                                                                case -1:
                                                                    v90 = ""
                                                                default:
                                                                    v18 = (v18 + (v18 + v90))
                                                                }
                                                            case -9223372036854775808:
                                                                v16 = v16
                                                            }
                                                            if (!false) {
                                                                v99 := GenS3{F0: 0, F1: 1}
                                                                switch (v4 >> ((0) & 63)) {
                                                                case 157:
                                                                    fmt.Println(v4, v89)
                                                                    switch (fi24 << ((531) & 63)) {
                                                                    case -1:
                                                                        v90 = v90
                                                                    default:
                                                                        v3[42] = (!(497 <= 295))
                                                                        v4 = (-(fi24))
                                                                        var v100 uint16 = (v89 >> ((4294967295) & 63))
                                                                        _ = v100
                                                                    }
                                                                }
                                                                fmt.Println(v17)
                                                                _ = v99
                                                            } else {
                                                                v91 += v91
                                                                var v101 int = (v4 & fi24)
                                                                v87 = append(v87, (fi88 >> ((0) & 63)))
                                                                _ = v101
                                                            }
                                                        } else {
                                                            if false {
                                                                v102 := GenS10{F0: 697, F1: -1}
                                                                var v103 uint16 = v89
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                            v4 -= (fi88 & v4)
                                                        }
                                                    default:
                                                        v104 := v2[fi88]
                                                        _ = v104
                                                    }
                                                    v18 = (v18 + v18)
                                                    _ = v91
                                                    _ = v92
                                                }
                                                _ = v89
                                                _ = v90
                                            }
                                            var v105 bool = (!true)
                                            fmt.Println(v18)
                                            var v106 string = ("foo bar" + v18)
                                            _ = v105
                                            _ = v106
                                        } else {
                                            switch v18 {
                                            case "\n\t":
                                                v107 := GenS7{F0: 2147483648, F1: 0, F2: true}
                                                _ = v107
                                            case "foo bar":
                                                if (317 <= 562) {
                                                    v4 -= v4
                                                    var v108 int16 = 874
                                                    _ = v108
                                                }
                                            case "":
                                                v16 += 0
                                            default:
                                                switch v18 {
                                                case "a":
                                                    var v109 uint64 = 2
                                                    _ = v109
                                                case "":
                                                    switch 32767 {
                                                    case -1:
                                                        if (280 > 1) {
                                                            v16 -= -1
                                                            fmt.Println(v17, v18)
                                                            var v110 string = (v18 + v18)
                                                            _ = v110
                                                        }
                                                    case 32767:
                                                        if ((0.1 + 3.14159) < 1e10) {
                                                            if true {
                                                                delete(v2, v4)
                                                                if len(v87) >= 2 {
                                                                    v87 = v87[1:len(v87)]
                                                                }
                                                                v111 := GenS13{F0: -478}
                                                                v87 = append(v87, (fi24 & (-(fi24))))
                                                                v4 = (^(fi88))
                                                                _ = v111
                                                            } else {
                                                                v112 := GenS1{F0: (-(fi88)), F1: -103}
                                                                var v113 int16 = 326
                                                                if false {
                                                                    v113 -= 32767
                                                                }
                                                                var v114 uint64 = 892
                                                                if len(v87) >= 2 {
                                                                    v87 = v87[1:len(v87)]
                                                                }
                                                                var v115 bool = (!true)
                                                                _ = v112
                                                                _ = v113
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                        }
                                                    default:
                                                        var v116 int8 = int8(fi88)
                                                        _ = v116
                                                    }
                                                    v117 := v2[v4]
                                                    _ = v117
                                                default:
                                                    if (!false) {
                                                        switch (fi24 % ((32767) | 1)) {
                                                        case 2147483647:
                                                            if ((119 > 232) && (!(genFn0(fi24, 255, 32767) != v16))) {
                                                                var v118 int64 = -399
                                                                var v119 uint16 = 1
                                                                _ = v118
                                                                _ = v119
                                                            }
                                                        case 32768:
                                                            if len(v87) >= 2 {
                                                                v87 = v87[1:len(v87)]
                                                            }
                                                        case 305:
                                                            v120 := GenS9{F0: 100}
                                                            _ = v120
                                                        }
                                                    } else {
                                                        var v121 int8 = -83
                                                        if (v16 <= (v16 >> ((1) & 63))) {
                                                            v4 = (v4 * v4)
                                                            if true {
                                                                v122 := GenS8{F0: 123, F1: genFn0((v4 + (fi88 ^ fi24)), 255, 141)}
                                                                switch fi24 {
                                                                case 702:
                                                                    var v123 uint16 = 1
                                                                    _ = v123
                                                                }
                                                                _ = v122
                                                            }
                                                            var v124 string = "foo bar"
                                                            _ = v124
                                                        } else {
                                                            if (!false) {
                                                                var v125 int32 = v16
                                                                if false {
                                                                    var v126 int64 = 0
                                                                    if false {
                                                                        _ = len(v2)
                                                                    }
                                                                    v4 = fi88
                                                                    v18 = v18
                                                                    var v127 float64 = (v17 - v17)
                                                                    if (false && (!false)) {
                                                                        if ((v16 + (v125 & v16)) >= -667) {
                                                                            v121 -= (v121 ^ (v121 + 127))
                                                                            switch (fi24 << ((65535) & 63)) {
                                                                            case 0:
                                                                                if true {
                                                                                    switch (("\n\t" + v18) + v18) {
                                                                                    case "foo bar":
                                                                                        v128 := v1[""]
                                                                                        _ = v128
                                                                                    case "hello":
                                                                                        var v129 int8 = 0
                                                                                        _ = v129
                                                                                    case "":
                                                                                        var v130 int8 = v121
                                                                                        _ = v130
                                                                                    }
                                                                                    switch v18 {
                                                                                    case "hello":
                                                                                        fmt.Println(fi24, v125)
                                                                                    case "a":
                                                                                        v126 = v126
                                                                                    default:
                                                                                        v4 = fi24
                                                                                        v131 := GenS9{F0: (v121 / ((42) | 1))}
                                                                                        _ = v131
                                                                                    }
                                                                                } else {
                                                                                    if (!false) {
                                                                                        v132 := v1[("foo bar" + ("" + (("hello" + "hello") + v18)))]
                                                                                        var v133 uint64 = 382
                                                                                        v127 = v127
                                                                                        _ = v132
                                                                                        _ = v133
                                                                                    } else {
                                                                                        if true {
                                                                                            v18 = v18
                                                                                            v125 = int32(v126)
                                                                                            if (!(v18 != "foo bar")) {
                                                                                                var v134 uint32 = 1
                                                                                                if true {
                                                                                                    var v135 int8 = v121
                                                                                                    var v136 float64 = ((0.0 / 0.1) / float64(1e10))
                                                                                                    v135 -= v135
                                                                                                    _ = v135
                                                                                                    _ = v136
                                                                                                }
                                                                                                _ = v134
                                                                                            }
                                                                                            v137 := GenS6{F0: v126}
                                                                                            switch 0 {
                                                                                            case -9223372036854775808:
                                                                                                var v138 int16 = 925
                                                                                                _ = v138
                                                                                            case 4294967295:
                                                                                                v4 -= (^(v4))
                                                                                            default:
                                                                                                v127 = ((0.1 * v17) + 0.1)
                                                                                            }
                                                                                            _ = v137
                                                                                        } else {
                                                                                            v16 += -1
                                                                                            v121 = v121
                                                                                        }
                                                                                        switch v4 {
                                                                                        case -9223372036854775808:
                                                                                            v139 := GenS14{F0: 750, F1: 1, F2: -1}
                                                                                            _ = v139
                                                                                        case 2147483647:
                                                                                            fmt.Println(fi88, v121, fi24)
                                                                                        default:
                                                                                            if (false || true) {
                                                                                                fmt.Println(v4, v17, fi88)
                                                                                                var v140 int64 = (-(v126))
                                                                                                _ = v140
                                                                                            }
                                                                                        }
                                                                                        v141 := GenS9{F0: v121}
                                                                                        var v142 uint64 = 30
                                                                                        if false {
                                                                                            v127 -= ((float64((3.14159 + 3.14159)) * 3.14159) + 1.5)
                                                                                        } else {
                                                                                            if true {
                                                                                                var v143 int16 = 32767
                                                                                                _ = v143
                                                                                            }
                                                                                            var v144 uint = 198
                                                                                            _ = v144
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        v17 += float64(float32(v127))
                                                                                        if len(v87) >= 2 {
                                                                                            v87 = v87[1:len(v87)]
                                                                                        }
                                                                                        if (128 == 255) {
                                                                                            v145 := GenS8{F0: 32767, F1: 42}
                                                                                            if (((((float32(float64(-2.25)) / (1.5 + (-2.25 - -2.25))) * 3.14159) * 0.0) * 0.1) == 3.14159) {
                                                                                                v146 := GenS15{F0: 0, F1: (^(v4))}
                                                                                                _ = v146
                                                                                            }
                                                                                            v147 := GenS10{F0: 255, F1: v126}
                                                                                            _ = v145
                                                                                            _ = v147
                                                                                        } else {
                                                                                            var v148 uint32 = 2147483648
                                                                                            _ = v148
                                                                                        }
                                                                                        v121 += v121
                                                                                        if false {
                                                                                            switch v18 {
                                                                                            case "foo bar":
                                                                                                switch (v18 + "a") {
                                                                                                case "":
                                                                                                    v121 -= (^(v121))
                                                                                                    _ = len(v3)
                                                                                                    v149 := GenS16{F0: 721, F1: (v142 + v142)}
                                                                                                    switch 4294967295 {
                                                                                                    case 2147483647:
                                                                                                        var v150 int64 = (v126 & (v126 / ((v126) | 1)))
                                                                                                        _ = v150
                                                                                                    }
                                                                                                    var v151 int16 = int16(v142)
                                                                                                    _ = v149
                                                                                                    _ = v151
                                                                                                default:
                                                                                                    var v152 int64 = 395
                                                                                                    _ = v152
                                                                                                }
                                                                                                if (true || (!false)) {
                                                                                                    fmt.Println(fi88, v18, v125)
                                                                                                }
                                                                                                if (1 <= 32767) {
                                                                                                    v121 = (v121 << ((255) & 63))
                                                                                                    if ((((v18 + "a") == (("foo bar" + v18) + "a")) || false) && (2147483647 == 350)) {
                                                                                                        if false {
                                                                                                            v153 := GenS4{F0: true}
                                                                                                            v87 = append(v87, 32768)
                                                                                                            v154 := GenS9{F0: int8(v125)}
                                                                                                            _ = v153
                                                                                                            _ = v154
                                                                                                        }
                                                                                                    }
                                                                                                    var v155 int = fi24
                                                                                                    _ = v155
                                                                                                } else {
                                                                                                    v126 = v126
                                                                                                }
                                                                                            case "hello":
                                                                                                var v156 uint32 = 449
                                                                                                _ = v156
                                                                                            case "":
                                                                                                if false {
                                                                                                    _ = len(v87)
                                                                                                    v157 := GenS9{F0: (v121 >> ((42) & 63))}
                                                                                                    var v158 uint16 = 1
                                                                                                    _ = v157
                                                                                                    _ = v158
                                                                                                } else {
                                                                                                    fmt.Println(v16)
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    if len(v87) >= 2 {
                                                                                                        v87 = v87[1:len(v87)]
                                                                                                    }
                                                                                                    _ = len(v87)
                                                                                                    var v159 int32 = 2147483647
                                                                                                    v4 = (v4 ^ (-(v4)))
                                                                                                    switch (fi88 & fi24) {
                                                                                                    case -173:
                                                                                                        if len(v87) >= 2 {
                                                                                                            v87 = v87[1:len(v87)]
                                                                                                        }
                                                                                                    case 310:
                                                                                                        v160 := GenS9{F0: v121}
                                                                                                        _ = v160
                                                                                                    default:
                                                                                                        var v161 int32 = 42
                                                                                                        _ = v161
                                                                                                    }
                                                                                                    v162 := GenS1{F0: (fi88 ^ 2147483647), F1: 42}
                                                                                                    v163 := v1[v18]
                                                                                                    if (!true) {
                                                                                                        v164 := GenS2{F0: 2}
                                                                                                        _ = v164
                                                                                                    }
                                                                                                    _ = v159
                                                                                                    _ = v162
                                                                                                    _ = v163
                                                                                                } else {
                                                                                                    v17 = 0.1
                                                                                                    var v165 float64 = 0.1
                                                                                                    _ = v165
                                                                                                }
                                                                                            }
                                                                                            v126 = (v126 << ((42) & 63))
                                                                                        } else {
                                                                                            v166 := GenS8{F0: 2, F1: (v16 << ((452) & 63))}
                                                                                            _ = v166
                                                                                        }
                                                                                        _ = v141
                                                                                        _ = v142
                                                                                    }
                                                                                }
                                                                            case 32768:
                                                                                switch (v4 & (^(v4))) {
                                                                                case -9223372036854775808:
                                                                                    var v167 int64 = v126
                                                                                    _ = v167
                                                                                }
                                                                            case 4294967295:
                                                                                var v168 float32 = float32(1e10)
                                                                                _ = v168
                                                                            }
                                                                            if false {
                                                                                switch (fi24 >> ((179) & 63)) {
                                                                                case 32768:
                                                                                    if len(v87) > 0 {
                                                                                        v87[0] = (fi88 / (((fi24 / (((fi88 % (((fi24 - (fi24 / ((v4) | 1)))) | 1))) | 1))) | 1))
                                                                                    }
                                                                                case 128:
                                                                                    if true {
                                                                                        if (v125 > (v125 * v16)) {
                                                                                            var v169 float32 = ((-2.25 + (1e10 / (1.5 - 0.0))) - 3.14159)
                                                                                            _ = v169
                                                                                        } else {
                                                                                            var v170 uint = 255
                                                                                            _ = v170
                                                                                        }
                                                                                        if (128 >= 4294967295) {
                                                                                            _ = len(v87)
                                                                                            if false {
                                                                                                var v171 bool = true
                                                                                                v125 = 2147483647
                                                                                                _ = v171
                                                                                            } else {
                                                                                                var v172 int64 = v126
                                                                                                if ((((!true) && (true || (0 > 18446744073709551615))) && true) && true) {
                                                                                                    var v173 uint8 = 127
                                                                                                    _ = v173
                                                                                                } else {
                                                                                                    v87 = append(v87, (fi88 * fi24))
                                                                                                }
                                                                                                var v174 uint8 = 42
                                                                                                _ = v172
                                                                                                _ = v174
                                                                                            }
                                                                                        } else {
                                                                                            if true {
                                                                                                switch v18 {
                                                                                                case "\n\t":
                                                                                                    delete(v2, v4)
                                                                                                case "foo bar":
                                                                                                    var v175 int16 = -32768
                                                                                                    _ = v175
                                                                                                default:
                                                                                                    v125 = v125
                                                                                                }
                                                                                            } else {
                                                                                                v176 := GenS7{F0: 877, F1: 0, F2: true}
                                                                                                var v177 uint = 18446744073709551615
                                                                                                _ = v176
                                                                                                _ = v177
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if (!true) {
                                                                                            if false {
                                                                                                _ = len(v87)
                                                                                                v1[v18] = v18
                                                                                            } else {
                                                                                                v178 := GenS17{F0: 2}
                                                                                                delete(v3, 4294967295)
                                                                                                if (259 != 2) {
                                                                                                    v179 := GenS2{F0: 0}
                                                                                                    _ = v179
                                                                                                }
                                                                                                _ = v178
                                                                                            }
                                                                                            var v180 int16 = 835
                                                                                            v181 := GenS9{F0: v121}
                                                                                            _ = v180
                                                                                            _ = v181
                                                                                        }
                                                                                        switch (fi88 + -665) {
                                                                                        case 1:
                                                                                            switch (v4 & 42) {
                                                                                            case -879:
                                                                                                fmt.Println(fi24)
                                                                                            case 4294967295:
                                                                                                v4 = fi88
                                                                                                v182 := GenS18{F0: (v126 % ((802) | 1)), F1: (fi24 + (^(fi88)))}
                                                                                                var v183 uint64 = 804
                                                                                                _ = v182
                                                                                                _ = v183
                                                                                            case 1:
                                                                                                switch v18 {
                                                                                                case "a":
                                                                                                    v87 = append(v87, (fi24 + v4))
                                                                                                case "hello":
                                                                                                    var v184 string = (v18 + v18)
                                                                                                    delete(v1, v184)
                                                                                                    v185 := GenS1{F0: v4, F1: v121}
                                                                                                    _ = v184
                                                                                                    _ = v185
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v4)
                                                                                                }
                                                                                            default:
                                                                                                delete(v2, fi24)
                                                                                            }
                                                                                        case -384:
                                                                                            switch (v18 + v18) {
                                                                                            case "":
                                                                                                var v186 int32 = (v125 / (((^(v16))) | 1))
                                                                                                _ = v186
                                                                                            case "a":
                                                                                                _ = len(v87)
                                                                                                v125 = genFn0((fi88 / ((fi88) | 1)), 214, -32768)
                                                                                            default:
                                                                                                var v187 float32 = (0.1 + (1.5 / float32(v127)))
                                                                                                _ = v187
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v17, v125)
                                                                                        var v188 float32 = 1.5
                                                                                        _ = v188
                                                                                    }
                                                                                case 42:
                                                                                    if true {
                                                                                        fmt.Println(v16, fi24, v125)
                                                                                        v4 += fi88
                                                                                        if ((v121 & v121) == (-(v121))) {
                                                                                            var v189 int64 = (-(v126))
                                                                                            var v190 int = fi24
                                                                                            _ = v189
                                                                                            _ = v190
                                                                                        } else {
                                                                                            var v191 uint32 = 4294967295
                                                                                            v126 = (^(v126))
                                                                                            v192 := v1["\n\t"]
                                                                                            _ = v191
                                                                                            _ = v192
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v18 = "hello"
                                                                                }
                                                                                var v193 uint64 = 0
                                                                                fmt.Println(v126)
                                                                                _ = v193
                                                                            } else {
                                                                                fmt.Println(v121, fi88, v127)
                                                                                switch v4 {
                                                                                case -2147483648:
                                                                                    v126 = v126
                                                                                case 2147483647:
                                                                                    delete(v2, fi24)
                                                                                case 4294967295:
                                                                                    switch (((v18 + "a") + v18) + v18) {
                                                                                    case "a":
                                                                                        if len(v87) > 0 {
                                                                                            v87[0] = (fi24 | v4)
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v16 = -2147483648
                                                                                    default:
                                                                                        var v194 uint8 = 2
                                                                                        _ = v194
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    var v195 uint8 = 232
                                                                                    v196 := GenS4{F0: false}
                                                                                    v125 = (v16 + (v16 % ((genFn0((fi24 & v4), 252, 32767)) | 1)))
                                                                                    _ = v195
                                                                                    _ = v196
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                var v197 uint8 = 2
                                                                                if (1 < (fi24 % (((-(fi24))) | 1))) {
                                                                                    v125 -= v125
                                                                                    fmt.Println(v197, v125)
                                                                                }
                                                                                switch v4 {
                                                                                case 128:
                                                                                    if ((-(v125)) <= 42) {
                                                                                        _ = len(v87)
                                                                                        if true {
                                                                                            var v198 uint8 = v197
                                                                                            var v199 bool = (!(!true))
                                                                                            v87 = append(v87, 4294967295)
                                                                                            v18 = "a"
                                                                                            if len(v87) > 0 {
                                                                                                v87[0] = (-(fi24))
                                                                                            }
                                                                                            _ = v198
                                                                                            _ = v199
                                                                                        } else {
                                                                                            v125 = (v125 | genFn0((v4 << ((65535) & 63)), 127, 1000))
                                                                                            if ((3.14159 / 1.0) == float32(v17)) {
                                                                                                var v200 int16 = 1000
                                                                                                if (false || (0.0 != (1.5 / ((float32(v127) * 3.14159) * 0.0)))) {
                                                                                                    var v201 int8 = (v121 / ((-128) | 1))
                                                                                                    v202 := GenS19{F0: v200, F1: (("a" + v18) + v18)}
                                                                                                    switch "" {
                                                                                                    case "hello":
                                                                                                        switch 117 {
                                                                                                        case -1:
                                                                                                            v203 := v1[v18]
                                                                                                            _ = v203
                                                                                                        case 32768:
                                                                                                            var v204 bool = (v127 >= v17)
                                                                                                            _ = v204
                                                                                                        case 32767:
                                                                                                            delete(v3, 0)
                                                                                                        }
                                                                                                    case "a":
                                                                                                        v205 := GenS20{F0: 645, F1: (v197 / ((0) | 1))}
                                                                                                        _ = v205
                                                                                                    case "\n\t":
                                                                                                        var v206 int = 0
                                                                                                        _ = v206
                                                                                                    default:
                                                                                                        var v207 uint16 = 0
                                                                                                        _ = v207
                                                                                                    }
                                                                                                    v16 -= genFn0((^(v4)), 221, v200)
                                                                                                    var v208 uint = 65535
                                                                                                    _ = v201
                                                                                                    _ = v202
                                                                                                    _ = v208
                                                                                                } else {
                                                                                                    switch 2147483647 {
                                                                                                    case 9223372036854775807:
                                                                                                        v200 = (v200 * v200)
                                                                                                    case 1:
                                                                                                        v4 = fi24
                                                                                                    default:
                                                                                                        var v209 uint32 = 4294967295
                                                                                                        v210 := GenS17{F0: 412}
                                                                                                        _ = v209
                                                                                                        _ = v210
                                                                                                    }
                                                                                                    v211 := v3[4294967295]
                                                                                                    v4 -= v4
                                                                                                    _ = v211
                                                                                                }
                                                                                                _ = v200
                                                                                            } else {
                                                                                                v212 := GenS21{F0: (v125 / ((v125) | 1)), F1: v126}
                                                                                                if (!true) {
                                                                                                    var v213 bool = ((42 == -32768) && false)
                                                                                                    v214 := v3[632]
                                                                                                    if v213 {
                                                                                                        fmt.Println(v125, v126, v127)
                                                                                                        v127 = float64(1e10)
                                                                                                    } else {
                                                                                                        _ = len(v1)
                                                                                                        switch int(v126) {
                                                                                                        case 1:
                                                                                                            if (241 != 4294967295) {
                                                                                                                v125 -= 1
                                                                                                                v215 := GenS21{F0: v16, F1: v126}
                                                                                                                fmt.Println(v125, v16)
                                                                                                                _ = v215
                                                                                                            }
                                                                                                        default:
                                                                                                            v216 := GenS5{F0: 42}
                                                                                                            _ = v216
                                                                                                        }
                                                                                                        fmt.Println(v18, v197)
                                                                                                    }
                                                                                                    _ = v213
                                                                                                    _ = v214
                                                                                                } else {
                                                                                                    _ = len(v2)
                                                                                                    var v217 float64 = ((v127 - v127) + v127)
                                                                                                    _ = v217
                                                                                                }
                                                                                                var v218 int16 = -91
                                                                                                _ = v212
                                                                                                _ = v218
                                                                                            }
                                                                                            var v219 int32 = (v125 | (v16 | 1))
                                                                                            v17 += ((v17 * v127) / (v17 - v17))
                                                                                            _ = v219
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v220 uint8 = 1
                                                                                    _ = v220
                                                                                }
                                                                                _ = v197
                                                                            } else {
                                                                                var v221 bool = false
                                                                                _ = v221
                                                                            }
                                                                            var v222 int32 = (-(v125))
                                                                            var v223 int8 = int8(v126)
                                                                            _ = v222
                                                                            _ = v223
                                                                        }
                                                                        var v224 float64 = -2.25
                                                                        if false {
                                                                            switch v18 {
                                                                            case "":
                                                                                if ((fi24 & v4) < -2147483648) {
                                                                                    var v225 uint = 0
                                                                                    _ = v225
                                                                                }
                                                                                if (true || (v18 >= v18)) {
                                                                                    v226 := GenS22{F0: false, F1: (!false), F2: 481}
                                                                                    switch 0 {
                                                                                    case -2147483648:
                                                                                        var v227 int16 = -378
                                                                                        _ = v227
                                                                                    default:
                                                                                        var v228 int8 = 0
                                                                                        _ = v228
                                                                                    }
                                                                                    var v229 uint64 = 4294967295
                                                                                    v224 -= v127
                                                                                    _ = v226
                                                                                    _ = v229
                                                                                } else {
                                                                                    v230 := GenS10{F0: 42, F1: (-(v126))}
                                                                                    var v231 uint16 = 2
                                                                                    fmt.Println(v16, v126, fi24)
                                                                                    _ = v230
                                                                                    _ = v231
                                                                                }
                                                                                v232 := GenS16{F0: 42, F1: 360}
                                                                                _ = v232
                                                                            case "hello":
                                                                                v126 = v126
                                                                            case "\n\t":
                                                                                if len(v87) >= 2 {
                                                                                    v87 = v87[1:len(v87)]
                                                                                }
                                                                            default:
                                                                                v233 := GenS9{F0: (v121 * v121)}
                                                                                _ = v233
                                                                            }
                                                                            if (true && (747 < 42)) {
                                                                                v4 += v4
                                                                                var v234 uint8 = 255
                                                                                _ = v234
                                                                            } else {
                                                                                if (false && true) {
                                                                                    var v235 int32 = (v125 / ((v16) | 1))
                                                                                    delete(v1, (v18 + v18))
                                                                                    _ = v235
                                                                                }
                                                                                v126 += (v126 << ((255) & 63))
                                                                            }
                                                                            v224 -= v17
                                                                        }
                                                                        var v236 uint32 = 602
                                                                        _ = v224
                                                                        _ = v236
                                                                    } else {
                                                                        v237 := v3[607]
                                                                        if true {
                                                                            fmt.Println(v4)
                                                                        }
                                                                        _ = v237
                                                                    }
                                                                    _ = v126
                                                                    _ = v127
                                                                } else {
                                                                    var v238 uint16 = 398
                                                                    v3[593] = false
                                                                    v121 = 23
                                                                    if false {
                                                                        if (-9223372036854775808 >= 9223372036854775807) {
                                                                            var v239 uint = 42
                                                                            v18 = ("hello" + v18)
                                                                            _ = v239
                                                                        }
                                                                    } else {
                                                                        v240 := v3[0]
                                                                        _ = v240
                                                                    }
                                                                    _ = v238
                                                                }
                                                                _ = v125
                                                            } else {
                                                                if (((((v17 - float64(1.5)) * v17) / (v17 * (3.14159 - 3.14159))) + v17) < (v17 * 1e10)) {
                                                                    v241 := GenS23{F0: v16, F1: 18446744073709551615, F2: v121}
                                                                    fmt.Println(fi88, v4, fi88)
                                                                    _ = v241
                                                                } else {
                                                                    if false {
                                                                        v242 := GenS24{F0: (-(v4))}
                                                                        v243 := GenS9{F0: 118}
                                                                        var v244 float32 = float32(v17)
                                                                        var v245 int32 = (v16 & (v16 ^ v16))
                                                                        _ = v242
                                                                        _ = v243
                                                                        _ = v244
                                                                        _ = v245
                                                                    }
                                                                    fmt.Println(v16, v4, fi88)
                                                                    var v246 int = fi24
                                                                    _ = v246
                                                                }
                                                                v247 := GenS25{F0: v16, F1: 850}
                                                                v18 = "foo bar"
                                                                v248 := GenS26{F0: 45, F1: (v121 >> ((18446744073709551615) & 63)), F2: true}
                                                                v249 := GenS27{F0: 375}
                                                                _ = v247
                                                                _ = v248
                                                                _ = v249
                                                            }
                                                        }
                                                        var v250 float64 = (float64((0.1 * ((0.0 / (-2.25 + (((float32((v17 - float64(1e10))) - 1.5) + ((0.0 / 1.5) + ((3.14159 - (0.1 - 1.5)) + 1e10))) + 1.5))) + 0.0))) + (float64(((1e10 + 1e10) - (-2.25 / (1.5 * float32(v17))))) / v17))
                                                        fmt.Println(v16, v121)
                                                        _ = v121
                                                        _ = v250
                                                    }
                                                }
                                            }
                                            v16 = v16
                                            var v251 int32 = v16
                                            _ = v251
                                        }
                                        _ = fi88
                                    }
                                }
                                if false {
                                    v17 += (v17 * (v17 + (v17 - v17)))
                                    var v252 uint = 18446744073709551615
                                    fmt.Println(v18, v252, v17)
                                    _ = v252
                                }
                                _ = v87
                            }
                        default:
                            fmt.Println(v17, fi24, v16)
                        }
                    } else {
                        switch v18 {
                        case "hello":
                            var v253 uint32 = 494
                            _ = v253
                        }
                        switch 42 {
                        case -2147483648:
                            switch (fi24 >> ((1) & 63)) {
                            case 32768:
                                var v254 int32 = genFn0((v4 << ((398) & 63)), 255, -573)
                                _ = v254
                            case 128:
                                var v255 int32 = 42
                                _ = v255
                            default:
                                v256 := GenS28{F0: 1, F1: -1, F2: 231}
                                _ = v256
                            }
                        default:
                            if ((true && false) && (true || (2147483647 > 4294967295))) {
                                var v257 int16 = -32768
                                _ = v257
                            } else {
                                v3[42] = false
                                if false {
                                    for fi258 := 0; fi258 < 20; fi258++ {
                                        var v259 uint64 = 4294967295
                                        fmt.Println(v4, v4)
                                        _ = fi258
                                        _ = v259
                                    }
                                    fmt.Println(v16, fi24, v16)
                                    v2[fi24] = fi24
                                } else {
                                    var v260 uint16 = 140
                                    v4 -= v4
                                    var v261 int = -139
                                    _ = v260
                                    _ = v261
                                }
                                fmt.Println(v17)
                            }
                        }
                        var v262 uint8 = 127
                        _ = v262
                    }
                    fmt.Println(v16, v18)
                    var v263 uint = 2
                    _ = v263
                }
            } else {
                fmt.Println(v17)
                v264 := []int{(v4 / ((1) | 1)), (v4 >> ((1) & 63))}
                switch fi24 {
                case 165:
                    delete(v1, v18)
                case 2147483647:
                    v265 := GenS29{F0: 0, F1: (^(v16)), F2: -1}
                    _ = v265
                default:
                    var v266 uint8 = 2
                    _ = v266
                }
                _ = v264
            }
        default:
            v17 += float64((1.5 * (-2.25 + (((1.5 / -2.25) - -2.25) * (((float32(0.1) - 3.14159) * 1.5) * -2.25)))))
        }
        _ = fi24
    }
    for fi267 := 0; fi267 < 4; fi267++ {
        _ = len(v3)
        _ = fi267
    }
    v268 := GenS29{F0: 561, F1: v16, F2: -323}
    v4 -= (v4 - v4)
    var v269 int64 = 1
    v4 = 42
    v269 = v269
    switch v18 {
    case "a":
        v18 = v18
    }
    if (65535 > 374) {
        var v270 float32 = (3.14159 / (0.1 / 1.0))
        var v271 uint = 255
        for fi272 := 0; fi272 < 19; fi272++ {
            if (v17 >= 1.5) {
                fmt.Println(v18, v4)
                for fi273 := 0; fi273 < 5; fi273++ {
                    v274 := v2[(-(fi273))]
                    if false {
                        v275 := GenS23{F0: 1, F1: 504, F2: 1}
                        var v276 int8 = -1
                        _ = v275
                        _ = v276
                    }
                    fmt.Println(v4)
                    _ = fi273
                    _ = v274
                }
                fmt.Println(v18, v4)
            }
            _ = fi272
        }
        _ = v270
        _ = v271
    } else {
        switch (v18 + (("\n\t" + (v18 + v18)) + v18)) {
        case "a":
            var v277 string = "a"
            _ = v277
        case "foo bar":
            v278 := GenS30{F0: 127}
            _ = v278
        }
        if (!false) {
            if true {
                for fi279 := 0; fi279 < 5; fi279++ {
                    switch (v18 + (v18 + v18)) {
                    case "\n\t":
                        if false {
                            var v280 int8 = 100
                            _ = v280
                        }
                    case "foo bar":
                        v281 := GenS22{F0: false, F1: true, F2: 0}
                        _ = v281
                    case "a":
                        v269 = (v269 ^ (v269 / (((v269 % ((v269) | 1))) | 1)))
                    default:
                        v282 := GenS9{F0: 1}
                        _ = v282
                    }
                    v283 := GenS27{F0: 2147483647}
                    v284 := []int{fi279, -405, (fi279 + v4)}
                    _ = fi279
                    _ = v283
                    _ = v284
                }
                v16 = 1
            } else {
                v285 := []int{(-(v4))}
                _ = v285
            }
        }
    }
    fmt.Println(v269, v4, v17)
    if (!(648 == 194)) {
        if false {
            delete(v3, 65535)
            var v286 int32 = genFn0((v4 ^ v4), 128, 1)
            v287 := GenS31{F0: 1, F1: (^(v269))}
            _ = v286
            _ = v287
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v269:", v269)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v268
    _ = v269
    _ = v3
    _ = v4
}
