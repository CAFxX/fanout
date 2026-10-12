package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 bool; F2 int64 }
type GenS5 struct { F0 bool; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int64; F2 int16 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32; F1 int16; F2 int16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string; F1 int; F2 int8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int8; F1 int }
type GenS17 struct { F0 uint; F1 uint16 }
type GenS18 struct { F0 int32; F1 uint16 }
type GenS19 struct { F0 bool; F1 uint16; F2 uint64 }
func genFn0(p0 uint16, p1 int, p2 int16) uint32 {
    v4 := []int{}
    fmt.Println(p0, p1)
    p1 = p1
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 += p0
        var v6 bool = (0.0 == ((3.14159 / 3.14159) / 1e10))
        _ = fi5
        _ = v6
    }
    var v7 int = (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 2147483648
}
func genFn1() bool {
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = 79
        v9 = -128
        var v10 int16 = -1
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{18, 42, 1}
    _ = v11
    return false
}

func main() {
    v1 := []uint32{42, 65535, 2147483647, 945}
    var v2 uint = 216
    v3 := []string{("hello" + "hello"), "\n\t"}
    v2 -= v2
    if true {
        switch 850 {
        case -9223372036854775808:
            var v12 bool = true
            _ = v12
        case -817:
            var v13 float64 = -2.25
            _ = v13
        case 127:
            for fi14 := 0; fi14 < 16; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    switch (fi15 % ((fi14) | 1)) {
                    case 14:
                        switch (fi15 % ((fi15) | 1)) {
                        case 9223372036854775807:
                            v2 -= v2
                        case -2147483648:
                            var v16 uint32 = 989
                            var v17 int16 = 32767
                            if len(v3) > 0 {
                                v3[0] = ("\n\t" + "foo bar")
                            }
                            v16 -= v16
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v18 := GenS1{F0: -1, F1: false}
                            v19 := GenS1{F0: 42, F1: false}
                            var v20 bool = false
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                        case -292:
                            fmt.Println(fi14)
                        default:
                            if (false && false) {
                                fmt.Println(fi15)
                                _ = len(v1)
                                fmt.Println(fi15, fi15)
                            }
                        }
                    case -259:
                        v1 = append(v1, genFn0(510, fi14, -713))
                    default:
                        v2 = v2
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(fi14, fi14, v2)
                        if (1e10 == (0.1 + (-2.25 / 1e10))) {
                            v1 = append(v1, 893)
                        }
                    }
                    _ = fi15
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    v2 += v2
                    _ = fi21
                }
                _ = fi14
            }
        }
        if ((-2.25 * 3.14159) != 3.14159) {
            fmt.Println(v2)
            v2 = v2
            v2 -= (v2 + v2)
        }
        fmt.Println(v2)
    }
    var v22 float32 = 0.0
    switch "a" {
    case "":
        if true {
            if false {
                var v23 uint64 = 18446744073709551615
                _ = v23
            } else {
                var v24 int = -254
                _ = v24
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                if len(v1) > 0 {
                    v1[0] = 318
                }
                if true {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        if (!((64 > 128) || false)) {
                            v27 := GenS1{F0: 42, F1: (32767 <= 1000)}
                            v28 := GenS2{F0: -9223372036854775808, F1: 42, F2: 0}
                            var v29 uint16 = 0
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            fmt.Println(v22, fi26)
                        }
                        _ = fi26
                    }
                    if true {
                        switch (("\n\t" + "a") + "foo bar") {
                        case "foo bar":
                            v30 := GenS3{F0: -128}
                            _ = v30
                        }
                        if false {
                            var v31 uint8 = 127
                            switch (("\n\t" + ("a" + ("\n\t" + "a"))) + ("\n\t" + "")) {
                            case "foo bar":
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    var v33 uint8 = v31
                                    switch (-(fi25)) {
                                    case 445:
                                        var v34 int32 = -2147483648
                                        _ = v34
                                    case 128:
                                        var v35 uint8 = (v31 - (^(v33)))
                                        _ = v35
                                    case 2147483647:
                                        var v36 uint16 = 255
                                        _ = v36
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                            case "a":
                                fmt.Println(v22, v31, fi25)
                            case "hello":
                                fmt.Println(fi25)
                            }
                            _ = v31
                        }
                        var v37 int16 = 1000
                        _ = v37
                    }
                }
                _ = fi25
            }
        } else {
            var v38 int8 = 42
            _ = v38
        }
    default:
        v39 := GenS2{F0: 0, F1: genFn0(0, 32767, -500), F2: 4294967295}
        _ = v39
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        fmt.Println(v22, fi40)
        _ = fi40
    }
    v2 -= 1
    v2 -= 1
    if (246 == 4294967295) {
        v41 := GenS4{F0: false, F1: ((((0.1 * 1.5) <= (1.5 * 1e10)) && false) || (!true)), F2: -401}
        _ = v41
    } else {
        if (-646 <= -353) {
            var v42 float32 = v22
            _ = v42
        }
        for fi43 := 0; fi43 < 12; fi43++ {
            v22 = v22
            if true {
                var v44 bool = false
                if (!v44) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 = (v2 | v2)
                    for fi45 := 0; fi45 < 11; fi45++ {
                        if true {
                            v22 -= -2.25
                            v44 = ("\n\t" == "\n\t")
                            var v46 string = "foo bar"
                            _ = v46
                        } else {
                            var v47 uint64 = 173
                            _ = v47
                        }
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v48 string = ("hello" + "")
                            fmt.Println(v48, fi45)
                            var v49 uint8 = 255
                            _ = v48
                            _ = v49
                        }
                        _ = fi45
                    }
                } else {
                    v50 := GenS3{F0: -128}
                    fmt.Println(v2, v22)
                    _ = v50
                }
                fmt.Println(v44, fi43, v22)
                _ = v44
            }
            v51 := GenS5{F0: true, F1: -242, F2: 255}
            _ = fi43
            _ = v51
        }
        if (182 >= 122) {
            if (false && true) {
                v52 := GenS6{F0: 18446744073709551615, F1: -9223372036854775808, F2: 1}
                if false {
                    v53 := GenS7{F0: -32768, F1: 1}
                    fmt.Println(v2, v22)
                    _ = v53
                } else {
                    v2 += v2
                    _ = len(v3)
                    switch 42 {
                    case -906:
                        if false {
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                for fi54 := 0; fi54 < 3; fi54++ {
                                    v2 = (v2 & (v2 << ((v2) & 63)))
                                    v22 += (3.14159 / (((v22 + (v22 / ((((1e10 * 1.5) + float32(float64(v22))) - (v22 - 0.1)) / v22))) / (v22 * v22)) / 0.1))
                                    for fi55 := 0; fi55 < 16; fi55++ {
                                        var v56 int64 = 594
                                        if false {
                                            if (true || (!(true || (!(255 > 981))))) {
                                                var v57 uint64 = 4294967295
                                                var v58 uint8 = 1
                                                _ = v57
                                                _ = v58
                                            }
                                            v59 := GenS8{F0: (((1e10 / 1.5) / ((float64(1e10) + 1e10) + -2.25)) < (0.0 * -2.25))}
                                            _ = v59
                                        } else {
                                            v22 -= (v22 + v22)
                                        }
                                        _ = fi55
                                        _ = v56
                                    }
                                    _ = fi54
                                }
                            }
                            var v60 int8 = 0
                            var v61 int64 = 1
                            _ = v60
                            _ = v61
                        } else {
                            v62 := GenS9{F0: 699}
                            _ = v62
                        }
                    case 4294967295:
                        switch -2147483648 {
                        case 0:
                            if (!(false && (!false))) {
                                v63 := GenS1{F0: 850, F1: true}
                                _ = v63
                            } else {
                                for fi64 := 0; fi64 < 20; fi64++ {
                                    if false {
                                        if true {
                                            var v65 uint64 = 1
                                            v66 := GenS5{F0: (!true), F1: 42, F2: 127}
                                            var v67 bool = false
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        } else {
                                            v2 -= 473
                                        }
                                        for fi68 := 0; fi68 < 5; fi68++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v22 = ((float32(0.0) + v22) - float32(1e10))
                                            var v69 uint32 = uint32(v2)
                                            _ = fi68
                                            _ = v69
                                        }
                                    } else {
                                        var v70 int64 = -9223372036854775808
                                        _ = v70
                                    }
                                    v2 = v2
                                    _ = fi64
                                }
                            }
                        case 9223372036854775807:
                            var v71 int8 = -128
                            _ = v71
                        default:
                            v22 -= v22
                        }
                    }
                }
                if true {
                    if true {
                        v72 := GenS4{F0: true, F1: (true && true), F2: 9223372036854775807}
                        _ = v72
                    }
                }
                _ = v52
            } else {
                var v73 uint = (^(v2))
                _ = v73
            }
            switch 2147483647 {
            case -2147483648:
                v74 := GenS5{F0: (true && ((v2 >= 18446744073709551615) && (float32(3.14159) < float32((0.1 / ((float64((1.5 * v22)) / 0.1) / 1.5)))))), F1: 412, F2: 2}
                _ = v74
            default:
                if ((!true) && false) {
                    v2 = 532
                    v3 = append(v3, ("hello" + (("foo bar" + ("\n\t" + "foo bar")) + "a")))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            if (32768 <= 847) {
                v75 := GenS10{F0: 0}
                fmt.Println(v22)
                _ = v75
            } else {
                fmt.Println(v22, v22)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "foo bar" {
                case "":
                    var v76 uint8 = 2
                    _ = v76
                case "foo bar":
                    var v77 int16 = -543
                    _ = v77
                }
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 625
            }
            v22 += 0.1
            v22 = v22
        }
    }
    v22 += 1e10
    var v78 float32 = 1.5
    if (float32((3.14159 / 0.1)) < 3.14159) {
        var v79 int = -9223372036854775808
        _ = v79
    } else {
        v80 := GenS11{F0: 2147483648, F1: -1, F2: -32768}
        _ = v80
    }
    var v81 uint16 = 255
    var v82 int64 = 1
    fmt.Println(v81)
    for fi83 := 0; fi83 < 17; fi83++ {
        if false {
            v84 := GenS12{F0: 588}
            var v85 int8 = 82
            _ = v84
            _ = v85
        }
        _ = fi83
    }
    switch (("" + "") + "a") {
    case "":
        if false {
            if false {
                if (1e10 == 1.5) {
                    v78 -= v78
                    var v86 uint = v2
                    _ = v86
                }
                var v87 uint32 = 2147483648
                v88 := GenS11{F0: (v87 >> (((v2 * 1)) & 63)), F1: 42, F2: 42}
                _ = v87
                _ = v88
            } else {
                if len(v1) > 0 {
                    v1[0] = 15
                }
            }
            if (("a" <= "hello") || true) {
                v89 := GenS13{F0: "\n\t", F1: 2147483647, F2: 0}
                if (!((("" + "hello") >= "\n\t") || true)) {
                    var v90 uint = v2
                    for fi91 := 0; fi91 < 17; fi91++ {
                        v81 += 674
                        for fi92 := 0; fi92 < 18; fi92++ {
                            var v93 uint8 = 0
                            var v94 float64 = -2.25
                            v94 = v94
                            _ = fi92
                            _ = v93
                            _ = v94
                        }
                        var v95 float64 = ((1e10 - 0.0) / 1.5)
                        _ = fi91
                        _ = v95
                    }
                    v2 -= v90
                    _ = v90
                }
                v96 := GenS1{F0: (v82 / ((v82) | 1)), F1: false}
                _ = v89
                _ = v96
            } else {
                fmt.Println(v82)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            fmt.Println(v22)
        }
    case "a":
        if true {
            v22 -= float32(((0.1 / ((float64((v78 / 1.0)) - (1.5 + 1.5)) + 1.5)) - (((-2.25 + 0.0) * (float64(1.5) * 1.5)) / 3.14159)))
            switch (("hello" + "") + "a") {
            case "hello":
                var v97 bool = false
                _ = v97
            case "":
                if true {
                    v98 := GenS6{F0: 140, F1: (v82 / (((v82 >> ((v2) & 63))) | 1)), F2: 1000}
                    _ = v98
                } else {
                    v1 = append(v1, 65535)
                }
            }
        } else {
            v78 += (v78 * v78)
            if false {
                var v99 uint16 = 32767
                if (((v22 * v78) / 0.1) >= float32(1e10)) {
                    v3 = append(v3, (("foo bar" + "hello") + "a"))
                    fmt.Println(v99, v22, v82)
                    var v100 int8 = 42
                    _ = v100
                }
                _ = v99
            } else {
                if ((v2 - v2) > 1) {
                    fmt.Println(v2)
                    var v101 int16 = -299
                    _ = v101
                } else {
                    if ((-(v82)) == -9223372036854775808) {
                        v102 := GenS14{F0: false}
                        v82 = v82
                        _ = v102
                    }
                    switch ("foo bar" + "") {
                    case "foo bar":
                        var v103 int16 = -1
                        _ = v103
                    case "\n\t":
                        if (-103 > 46) {
                            var v104 string = "hello"
                            v22 = ((v78 * v78) - v78)
                            v104 = v104
                            _ = v104
                        } else {
                            var v105 uint64 = 0
                            v106 := GenS15{F0: 42}
                            if (127 <= 42) {
                                var v107 int16 = 32767
                                for fi108 := 0; fi108 < 9; fi108++ {
                                    var v109 int = fi108
                                    var v110 int16 = v107
                                    _ = fi108
                                    _ = v109
                                    _ = v110
                                }
                                for fi111 := 0; fi111 < 10; fi111++ {
                                    if ((!(("" + "foo bar") < "a")) || (v82 < v82)) {
                                        var v112 uint = v2
                                        _ = v112
                                    } else {
                                        if false {
                                            switch "a" {
                                            case "foo bar":
                                                for fi113 := 0; fi113 < 6; fi113++ {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    fmt.Println(v2)
                                                    _ = fi113
                                                }
                                            case "hello":
                                                v114 := GenS16{F0: 42, F1: (fi111 ^ fi111)}
                                                _ = v114
                                            case "":
                                                var v115 uint32 = uint32(v2)
                                                _ = v115
                                            }
                                            for fi116 := 0; fi116 < 9; fi116++ {
                                                switch "hello" {
                                                case "\n\t":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                if false {
                                                    var v117 int = int(v105)
                                                    _ = v117
                                                } else {
                                                    var v118 uint8 = 2
                                                    v105 -= (v105 & 0)
                                                    var v119 float64 = float64(v78)
                                                    _ = v118
                                                    _ = v119
                                                }
                                                var v120 int32 = 1
                                                _ = fi116
                                                _ = v120
                                            }
                                        } else {
                                            for fi121 := 0; fi121 < 4; fi121++ {
                                                if (!false) {
                                                    switch fi121 {
                                                    case 32768:
                                                        if (fi111 == (fi121 << ((2) & 63))) {
                                                            v2 = 65535
                                                            if len(v3) > 0 {
                                                                v3[0] = ("a" + "foo bar")
                                                            }
                                                            var v122 int32 = 42
                                                            _ = v122
                                                        }
                                                    case 4294967295:
                                                        v123 := GenS17{F0: (^(v2)), F1: (v81 / (((v81 / (((v81 + v81)) | 1))) | 1))}
                                                        _ = v123
                                                    }
                                                } else {
                                                    var v124 float64 = float64(v78)
                                                    fmt.Println(v107)
                                                    _ = v124
                                                }
                                                var v125 bool = false
                                                _ = fi121
                                                _ = v125
                                            }
                                        }
                                        var v126 uint32 = 1
                                        _ = v126
                                    }
                                    v127 := GenS18{F0: 27, F1: (^(v81))}
                                    for fi128 := 0; fi128 < 15; fi128++ {
                                        v2 = 42
                                        _ = fi128
                                    }
                                    _ = fi111
                                    _ = v127
                                }
                                _ = v107
                            } else {
                                _ = len(v3)
                            }
                            _ = v105
                            _ = v106
                        }
                    default:
                        fmt.Println(v81, v22, v82)
                    }
                }
            }
            fmt.Println(v81)
        }
    case "\n\t":
        v2 -= 1
    default:
        var v129 bool = (!(42 >= 255))
        _ = v129
    }
    if true {
        v130 := GenS13{F0: ("foo bar" + (("hello" + "hello") + ("foo bar" + "hello"))), F1: 2147483647, F2: 100}
        switch 32767 {
        case -136:
            v3 = append(v3, "a")
        case 998:
            var v131 float32 = 1.5
            _ = v131
        default:
            v132 := GenS13{F0: "hello", F1: 127, F2: -128}
            _ = v132
        }
        if (((("a" + "foo bar") > ("\n\t" + "\n\t")) || false) || (32768 > 0)) {
            v2 += (v2 % ((v2) | 1))
            for fi133 := 0; fi133 < 16; fi133++ {
                for fi134 := 0; fi134 < 16; fi134++ {
                    v135 := GenS19{F0: true, F1: v81, F2: 4294967295}
                    _ = fi134
                    _ = v135
                }
                var v136 int8 = int8(v82)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi133
                _ = v136
            }
        }
        _ = v130
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v137 int32 = 42
        if (!false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v82 = (v82 + 158)
        }
        _ = v137
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v78
    _ = v81
    _ = v82
}
