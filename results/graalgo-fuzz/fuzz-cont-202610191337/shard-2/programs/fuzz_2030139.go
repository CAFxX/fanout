package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 uint64; F1 uint8; F2 string }
type GenS15 struct { F0 uint8; F1 uint32 }
type GenS16 struct { F0 int32; F1 uint16 }
type GenS17 struct { F0 uint8 }
func genFn0() uint {
    var v7 float64 = (0.1 + float64(1e10))
    var v8 int64 = 42
    var v9 int64 = (^(v8))
    var v10 int = -618
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 string, p1 bool, p2 int32) int8 {
    p2 -= (p2 * p2)
    if (!p1) {
        fmt.Println(p1, p2)
    } else {
        v11 := GenS1{F0: 4294967295}
        v12 := GenS1{F0: 42}
        v13 := GenS1{F0: 2147483647}
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(p0)
    var v14 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := []string{"", "\n\t", ("" + (("hello" + "hello") + "foo bar")), (("a" + "") + ("a" + ((("\n\t" + ("" + (("hello" + "\n\t") + ("hello" + "foo bar")))) + ("a" + ("" + "\n\t"))) + "foo bar")))}
    var v3 int16 = -295
    v4 := []string{"hello"}
    v5 := []string{"\n\t", ""}
    v6 := make(map[string]string)
    v3 = (v3 >> ((1) & 63))
    if false {
        var v15 uint8 = 128
        switch ("hello" + "a") {
        case "a":
            for fi16 := 0; fi16 < 14; fi16++ {
                if (!((0.0 * 0.0) > (0.1 + 0.0))) {
                    fmt.Println(fi16, v3, v3)
                }
                _ = fi16
            }
        case "":
            fmt.Println(v15)
        case "\n\t":
            v3 = 0
        default:
            var v17 int = 2147483647
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            switch "foo bar" {
            case "foo bar":
                v1 += (^(v1))
            case "hello":
                if true {
                    if true {
                        if false {
                            v1 = (v1 & (v1 % ((v1) | 1)))
                            var v19 int8 = 127
                            if true {
                                var v20 int = fi18
                                fmt.Println(v1)
                                var v21 uint32 = 42
                                _ = v20
                                _ = v21
                            } else {
                                var v22 int64 = -1
                                _ = v22
                            }
                            _ = v19
                        }
                        v23 := GenS2{F0: ("\n\t" + "\n\t"), F1: 2}
                        v15 += (v15 ^ (v15 % ((128) | 1)))
                        _ = v23
                    } else {
                        v24 := GenS2{F0: "\n\t", F1: (v15 ^ (v15 + 127))}
                        var v25 string = ("a" + ((("foo bar" + "foo bar") + "foo bar") + "foo bar"))
                        _ = v24
                        _ = v25
                    }
                    if false {
                        if true {
                            v26 := GenS1{F0: 42}
                            _ = v26
                        } else {
                            if len(v4) > 0 {
                                v4[0] = (("\n\t" + "hello") + "")
                            }
                            fmt.Println(v3)
                            fmt.Println(v3)
                        }
                        v27 := GenS3{F0: 42, F1: -425}
                        v28 := GenS2{F0: ("\n\t" + ""), F1: (v15 >> ((288) & 63))}
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(v1, v3)
                        for fi29 := 0; fi29 < 5; fi29++ {
                            v15 = 2
                            if (!false) {
                                if false {
                                    _ = len(v6)
                                    v1 += v1
                                    if false {
                                        switch fi29 {
                                        case -9223372036854775808:
                                            if (!true) {
                                                var v30 string = "\n\t"
                                                v15 = (^(v15))
                                                _ = v30
                                            } else {
                                                var v31 int64 = 389
                                                var v32 int = fi29
                                                if len(v2) > 0 {
                                                    v2[0] = "hello"
                                                }
                                                _ = v31
                                                _ = v32
                                            }
                                        case 4294967295:
                                            if true {
                                                v3 += int16(fi18)
                                            }
                                        case -1:
                                            v33 := GenS1{F0: 664}
                                            v3 += v3
                                            _ = v33
                                        }
                                    }
                                    _ = len(v5)
                                    if true {
                                        _ = len(v4)
                                        if false {
                                            fmt.Println(fi18, fi29)
                                            var v34 uint8 = v15
                                            if (!((((4294967295 >= 2147483647) || ((fi18 >= fi29) || (4294967295 >= 4294967295))) && (true || false)) && true)) {
                                                var v35 uint = 4294967295
                                                _ = v35
                                            } else {
                                                var v36 uint32 = 42
                                                fmt.Println(fi29, v15, v36)
                                                _ = v36
                                            }
                                            v37 := GenS4{F0: 1}
                                            v34 -= v34
                                            _ = v34
                                            _ = v37
                                        }
                                        v38 := GenS5{F0: -621}
                                        v15 -= v15
                                        _ = v38
                                    } else {
                                        var v39 int32 = 1
                                        var v40 int16 = v3
                                        _ = v39
                                        _ = v40
                                    }
                                    var v41 int8 = 38
                                    var v42 int = (fi29 % ((fi18) | 1))
                                    v43 := GenS3{F0: 42, F1: -2147483648}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v2 = append(v2, (("\n\t" + (("foo bar" + "a") + "a")) + "foo bar"))
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            } else {
                                var v44 string = ("" + "hello")
                                var v45 int64 = 42
                                _ = v44
                                _ = v45
                            }
                            _ = fi29
                        }
                    }
                }
            }
            v3 = -1
            for fi46 := 0; fi46 < 4; fi46++ {
                fmt.Println(fi46, v1, v15)
                _ = fi46
            }
            _ = fi18
        }
        _ = v15
    }
    if (709 <= 2) {
        for fi47 := 0; fi47 < 15; fi47++ {
            for fi48 := 0; fi48 < 7; fi48++ {
                switch ("\n\t" + "hello") {
                case "a":
                    v3 = v3
                    if (!(32767 > 65535)) {
                        var v49 int32 = 500
                        v50 := GenS6{F0: (((("" + "foo bar") + "\n\t") + ("hello" + ("hello" + ("a" + "\n\t")))) + "hello"), F1: 127}
                        _ = v49
                        _ = v50
                    } else {
                        switch fi48 {
                        case -95:
                            switch "" {
                            case "\n\t":
                                var v51 float64 = 3.14159
                                if true {
                                    v51 += (float64(1.5) / (1.5 * 1e10))
                                    delete(v6, "\n\t")
                                }
                                if (0 <= (v1 * v1)) {
                                    _ = len(v2)
                                    fmt.Println(fi47)
                                    v1 += v1
                                }
                                _ = v51
                            }
                        case 1:
                            if len(v4) > 0 {
                                v4[0] = ("a" + ("hello" + ""))
                            }
                        case 4294967295:
                            v4 = append(v4, "foo bar")
                        default:
                            fmt.Println(v1, fi48, v1)
                            v1 -= v1
                            if false {
                                v52 := GenS1{F0: 2}
                                if false {
                                    if true {
                                        if true {
                                            v53 := GenS7{F0: (true || (-1 != v3))}
                                            _ = v53
                                        }
                                        v3 = (^(v3))
                                    }
                                    v1 = 4294967295
                                }
                                _ = v52
                            }
                        }
                        var v54 uint8 = 255
                        _ = v54
                    }
                    switch -9223372036854775808 {
                    case -416:
                        v55 := GenS1{F0: 985}
                        _ = v55
                    case -998:
                        var v56 int16 = v3
                        _ = v56
                    case 2147483647:
                        var v57 float64 = 3.14159
                        _ = v57
                    }
                case "\n\t":
                    v58 := GenS5{F0: -9223372036854775808}
                    _ = v58
                case "foo bar":
                    var v59 bool = false
                    _ = v59
                default:
                    var v60 int32 = 0
                    _ = v60
                }
                _ = len(v5)
                v61 := GenS5{F0: -9223372036854775808}
                _ = fi48
                _ = v61
            }
            _ = fi47
        }
        v62 := GenS8{F0: 492}
        delete(v6, "")
        _ = v62
    } else {
        var v63 uint32 = 2
        v5 = append(v5, "hello")
        _ = v63
    }
    var v64 int32 = 904
    for fi65 := 0; fi65 < 20; fi65++ {
        if (164 < 0) {
            var v66 uint = genFn0()
            if true {
                if (!(2 > uint8(v1))) {
                    v67 := GenS7{F0: ((v66 ^ (v66 / (((v66 & v66)) | 1))) != genFn0())}
                    switch "" {
                    case "foo bar":
                        v68 := v6[("hello" + (("\n\t" + (("a" + ("hello" + "a")) + ("foo bar" + ("foo bar" + "foo bar")))) + ("hello" + ((("\n\t" + "hello") + "a") + ("hello" + (("\n\t" + "\n\t") + "hello"))))))]
                        _ = v68
                    }
                    fmt.Println(fi65, v3, v3)
                    _ = v67
                }
            }
            _ = v66
        } else {
            var v69 uint16 = 255
            _ = v69
        }
        if (!true) {
            for fi70 := 0; fi70 < 6; fi70++ {
                var v71 int64 = 9223372036854775807
                v72 := GenS5{F0: -9223372036854775808}
                _ = fi70
                _ = v71
                _ = v72
            }
            fmt.Println(fi65)
        } else {
            if (0 > 32767) {
                var v73 uint = 1
                for fi74 := 0; fi74 < 18; fi74++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi74
                }
                for fi75 := 0; fi75 < 7; fi75++ {
                    v76 := GenS5{F0: 42}
                    _ = fi75
                    _ = v76
                }
                _ = v73
            }
            fmt.Println(v64, v3)
            v64 -= (v64 << ((genFn0()) & 63))
        }
        for fi77 := 0; fi77 < 3; fi77++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                v3 = (v3 * (-(v3)))
                v78 := GenS9{F0: 493, F1: 65535}
                _ = v78
            } else {
                v79 := GenS10{F0: v1, F1: int8(v1), F2: 255}
                fmt.Println(fi65)
                _ = v79
            }
            var v80 uint64 = 4294967295
            var v81 int16 = (v3 * (v3 & (v3 + v3)))
            if true {
                fmt.Println(v80, v3, v3)
                switch ("" + ("foo bar" + "\n\t")) {
                case "a":
                    var v82 int32 = 2147483647
                    _ = v82
                default:
                    fmt.Println(fi65, v81)
                }
                v1 = v80
            }
            _ = fi77
            _ = v80
            _ = v81
        }
        _ = fi65
    }
    if true {
        for fi83 := 0; fi83 < 8; fi83++ {
            var v84 string = ("a" + "hello")
            _ = fi83
            _ = v84
        }
        switch -130 {
        case 42:
            for fi85 := 0; fi85 < 16; fi85++ {
                fmt.Println(v64, v1)
                _ = fi85
            }
        case 127:
            v86 := GenS11{F0: 257}
            _ = v86
        }
    } else {
        _ = len(v4)
    }
    var v87 uint16 = 65535
    v64 += 2147483647
    if ((-2.25 <= ((float64((0.0 * (float32((-2.25 + (1.5 - (3.14159 + -2.25)))) + 0.1))) * 0.0) / 1.0)) || (true && (!(1 <= 18446744073709551615)))) {
        fmt.Println(v1, v64)
        v87 = (^(v87))
        switch 9223372036854775807 {
        case 127:
            v3 = (-(v3))
        case 120:
            fmt.Println(v1, v64, v87)
        case 2147483647:
            for fi88 := 0; fi88 < 11; fi88++ {
                if (false && false) {
                    var v89 int32 = v64
                    if false {
                        v87 = v87
                        v90 := GenS9{F0: (^(v1)), F1: (v87 - 42)}
                        var v91 float32 = float32(3.14159)
                        _ = v90
                        _ = v91
                    }
                    _ = v89
                }
                var v92 int64 = 9223372036854775807
                v93 := GenS4{F0: (v92 | (v92 >> ((18446744073709551615) & 63)))}
                _ = fi88
                _ = v92
                _ = v93
            }
        }
    } else {
        var v94 uint64 = v1
        v95 := GenS12{F0: 578, F1: 42}
        for fi96 := 0; fi96 < 16; fi96++ {
            v87 = v87
            if (162 >= 42) {
                for fi97 := 0; fi97 < 18; fi97++ {
                    if ((v94 ^ 2) >= 4294967295) {
                        if ((float32((float64(0.0) / 3.14159)) - 1.5) <= 0.1) {
                            if ((true && (!false)) && true) {
                                switch "hello" {
                                case "foo bar":
                                    v98 := v6[("a" + "\n\t")]
                                    _ = v98
                                case "":
                                    var v99 int32 = (v64 + 983)
                                    _ = v99
                                }
                                fmt.Println(v94)
                            } else {
                                _ = len(v6)
                            }
                            var v100 uint16 = uint16(v1)
                            _ = v100
                        }
                        v94 = v94
                        switch "foo bar" {
                        case "":
                            v87 = (^(v87))
                        case "hello":
                            switch (^(fi97)) {
                            case -1:
                                v1 = 0
                            case 2147483647:
                                if (("hello" + "a") == (("\n\t" + "") + ("hello" + "a"))) {
                                    v101 := GenS6{F0: "", F1: uint8(v94)}
                                    if false {
                                        v94 -= 18446744073709551615
                                        v87 -= (v87 / ((255) | 1))
                                    }
                                    v94 = v94
                                    var v102 int8 = 0
                                    v103 := GenS1{F0: 42}
                                    _ = v101
                                    _ = v102
                                    _ = v103
                                } else {
                                    var v104 int16 = v3
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(v64)
                                    _ = v104
                                }
                            default:
                                var v105 int8 = 104
                                _ = v105
                            }
                        }
                        var v106 string = ("foo bar" + ("\n\t" + ""))
                        var v107 uint32 = 118
                        var v108 string = v106
                        _ = v106
                        _ = v107
                        _ = v108
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v3 = (v3 & (-(v3)))
                        default:
                            if (!true) {
                                var v109 uint32 = 2147483647
                                v109 += (v109 << ((959) & 63))
                                var v110 uint16 = (v87 | (v87 / (((^(v87))) | 1)))
                                if (v109 < v109) {
                                    switch (fi96 % ((int(v1)) | 1)) {
                                    case -925:
                                        fmt.Println(v109, v87, v110)
                                    case 100:
                                        v94 -= 42
                                    case -9223372036854775808:
                                        var v111 float64 = (((1.5 * (3.14159 - 3.14159)) + 3.14159) - float64(1e10))
                                        _ = v111
                                    }
                                } else {
                                    if (uint64(v110) < uint64(v87)) {
                                        v87 -= v110
                                        switch ("" + "hello") {
                                        case "foo bar":
                                            v64 += (^(v64))
                                        case "":
                                            var v112 int64 = -123
                                            _ = v112
                                        case "hello":
                                            fmt.Println(v109)
                                            fmt.Println(v1, v109)
                                        }
                                    } else {
                                        v113 := GenS12{F0: 1, F1: 221}
                                        var v114 int = -553
                                        if (!true) {
                                            var v115 float64 = 3.14159
                                            _ = v115
                                        } else {
                                            v64 += 0
                                            v110 += v110
                                            fmt.Println(v87)
                                        }
                                        _ = v113
                                        _ = v114
                                    }
                                    if true {
                                        if ((3.14159 + (3.14159 * 3.14159)) > -2.25) {
                                            var v116 string = "\n\t"
                                            v117 := GenS13{F0: (fi96 ^ (fi97 ^ fi96)), F1: 255}
                                            _ = v116
                                            _ = v117
                                        } else {
                                            if true {
                                                fmt.Println(fi97, v64, v109)
                                            } else {
                                                var v118 uint32 = v109
                                                var v119 uint8 = 0
                                                _ = v118
                                                _ = v119
                                            }
                                        }
                                        v110 = (v110 & v87)
                                    } else {
                                        switch "a" {
                                        case "a":
                                            switch "foo bar" {
                                            case "\n\t":
                                                v120 := GenS5{F0: -1}
                                                _ = v120
                                            case "hello":
                                                if false {
                                                    v109 = (v109 + v109)
                                                    var v121 uint32 = 2147483648
                                                    v110 -= 65535
                                                    _ = v121
                                                }
                                                if (1 < int8(v87)) {
                                                    switch (fi97 / ((-2147483648) | 1)) {
                                                    case 32768:
                                                        fmt.Println(v64)
                                                    case 1:
                                                        fmt.Println(v110, v3, v94)
                                                    case 128:
                                                        var v122 int64 = -741
                                                        _ = v122
                                                    default:
                                                        var v123 float64 = (0.1 / 1.0)
                                                        _ = v123
                                                    }
                                                    var v124 string = "\n\t"
                                                    v125 := GenS14{F0: 2, F1: 216, F2: (v124 + "\n\t")}
                                                    _ = v124
                                                    _ = v125
                                                } else {
                                                    fmt.Println(v1, v109, fi97)
                                                    v64 -= (-(v64))
                                                }
                                            case "a":
                                                v6[("a" + "a")] = "foo bar"
                                            default:
                                                fmt.Println(fi97, v87)
                                                var v126 int8 = -125
                                                v127 := GenS15{F0: 127, F1: v109}
                                                _ = v126
                                                _ = v127
                                            }
                                        default:
                                            if false {
                                                var v128 uint32 = v109
                                                v109 = (^(v109))
                                                fmt.Println(fi96)
                                                _ = v128
                                            } else {
                                                var v129 int64 = -1
                                                v1 = v94
                                                fmt.Println(v87)
                                                _ = v129
                                            }
                                        }
                                        switch int(v3) {
                                        case -2147483648:
                                            _ = len(v2)
                                        case 127:
                                            var v130 float64 = (0.1 - ((0.1 + 0.1) * -2.25))
                                            _ = v130
                                        }
                                    }
                                    v87 = (v87 % (((v110 << ((genFn0()) & 63))) | 1))
                                    if (18446744073709551615 != 75) {
                                        v110 += (v87 % ((2) | 1))
                                        var v131 uint32 = 42
                                        v132 := GenS16{F0: (v64 | 0), F1: (v87 | (^(v110)))}
                                        _ = v131
                                        _ = v132
                                    }
                                    switch 32768 {
                                    case 32767:
                                        var v133 uint = 0
                                        _ = v133
                                    }
                                }
                                switch "" {
                                case "foo bar":
                                    if (65535 <= (v110 >> ((1) & 63))) {
                                        v1 = (v1 | 223)
                                        v109 = (v109 >> ((4294967295) & 63))
                                    }
                                case "":
                                    v134 := GenS17{F0: 255}
                                    _ = v134
                                case "\n\t":
                                    v110 += v87
                                }
                                var v135 float32 = (((0.1 - (1.5 * 1.5)) + 1.5) - 1.5)
                                _ = v109
                                _ = v110
                                _ = v135
                            }
                        }
                        fmt.Println(v3)
                        v136 := GenS7{F0: false}
                        _ = v136
                    }
                    _ = fi97
                }
                var v137 float32 = 1e10
                fmt.Println(fi96, v137, v137)
                _ = v137
            }
            fmt.Println(v3)
            _ = fi96
        }
        _ = v94
        _ = v95
    }
    for fi138 := 0; fi138 < 8; fi138++ {
        if (18446744073709551615 > 18446744073709551615) {
            var v139 int8 = genFn1(("\n\t" + "foo bar"), true, (-(v64)))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v140 int = (-(fi138))
            _ = v139
            _ = v140
        } else {
            v141 := GenS14{F0: v1, F1: 255, F2: (("a" + ("" + "foo bar")) + "\n\t")}
            _ = v141
        }
        _ = fi138
    }
    v3 -= v3
    fmt.Println(v1, v3, v3)
    v87 -= 255
    _ = len(v6)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v87
}
