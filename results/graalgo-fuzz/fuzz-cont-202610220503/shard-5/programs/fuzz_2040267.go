package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 bool; F2 bool }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 int16; F1 bool; F2 uint }
type GenS13 struct { F0 uint; F1 bool; F2 int32 }
func genFn0(p0 uint64, p1 int8, p2 uint32) int64 {
    v6 := []int{-500, int(p1)}
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            p1 = 42
            p2 = (p2 % ((p2) | 1))
            _ = fi8
        }
        var v9 uint8 = 42
        v9 = 42
        _ = fi7
        _ = v9
    }
    var v10 float32 = 0.0
    if ((("a" + "\n\t") + "hello") <= "foo bar") {
        var v11 float64 = 0.0
        _ = v11
    }
    var v12 uint32 = (p2 | 65535)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v12
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 223
    var v3 int = -9223372036854775808
    var v4 int = v3
    v5 := make(map[uint32]bool)
    var v13 uint16 = 250
    v2 -= 79
    v14 := GenS1{F0: v2, F1: (1 != 1), F2: (v3 - (v4 % ((v4) | 1)))}
    v2 = (v2 ^ 42)
    for fi15 := 0; fi15 < 9; fi15++ {
        if (1.5 <= 3.14159) {
            switch "a" {
            case "hello":
                fmt.Println(v4, v13)
            case "foo bar":
                v13 = (v13 * (v13 % ((65535) | 1)))
            default:
                for fi16 := 0; fi16 < 18; fi16++ {
                    fmt.Println(v2, fi15)
                    if false {
                        v17 := v1[65535]
                        fmt.Println(v13, v3)
                        var v18 string = "hello"
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(v2, v3, fi15)
                    _ = fi16
                }
            }
        }
        _ = fi15
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        fmt.Println(v13, v4)
        for fi20 := 0; fi20 < 19; fi20++ {
            if (2147483648 >= 2) {
                fmt.Println(v13)
                v21 := GenS1{F0: (v2 | 29), F1: (-128 != 83), F2: fi19}
                v22 := GenS2{F0: 1, F1: 18446744073709551615}
                _ = v21
                _ = v22
            }
            var v23 uint16 = 1
            _ = fi20
            _ = v23
        }
        fmt.Println(fi19, fi19)
        _ = fi19
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        switch "hello" {
        case "\n\t":
            v13 = v13
        case "":
            v25 := GenS3{F0: 42, F1: v13, F2: ("foo bar" + "foo bar")}
            _ = v25
        case "foo bar":
            switch -1 {
            case -2147483648:
                v2 = 219
            default:
                if true {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v13 -= 1
                        _ = fi26
                    }
                }
            }
        default:
            v27 := []int{9223372036854775807, (v3 >> ((42) & 63)), 1}
            _ = v27
        }
        v2 += v2
        var v28 int32 = 2147483647
        _ = fi24
        _ = v28
    }
    var v29 float64 = -2.25
    if (((-2.25 / 3.14159) + 1e10) == 0.1) {
        switch "" {
        case "\n\t":
            v29 -= v29
        case "foo bar":
            v30 := []int{v3, (v4 % (((v4 ^ v3)) | 1))}
            _ = v30
        }
        switch "hello" {
        case "a":
            if (true || (42 == -725)) {
                var v31 string = "foo bar"
                _ = v31
            } else {
                var v32 uint64 = 2
                v33 := GenS4{F0: int64(v13), F1: (!(!(9223372036854775807 <= -315)))}
                var v34 uint64 = (v32 | 759)
                _ = v32
                _ = v33
                _ = v34
            }
        case "hello":
            if false {
                if false {
                    v35 := GenS3{F0: -128, F1: v13, F2: "\n\t"}
                    var v36 uint32 = 2147483648
                    if (float32((1e10 + v29)) > -2.25) {
                        if ("a" != (((((("\n\t" + "foo bar") + "") + "") + "a") + (("\n\t" + "a") + "a")) + "")) {
                            if (!("a" == "")) {
                                if false {
                                    v4 = 399
                                    var v37 uint32 = v36
                                    _ = v37
                                } else {
                                    var v38 float64 = 0.0
                                    _ = v38
                                }
                            } else {
                                var v39 uint16 = (v13 + v13)
                                _ = v39
                            }
                            v40 := []int{}
                            if (((1e10 / v29) * ((float64(3.14159) / v29) - (v29 - (0.1 - 0.1)))) != float64((3.14159 / 1.0))) {
                                var v41 float32 = 3.14159
                                for fi42 := 0; fi42 < 8; fi42++ {
                                    if (44 != 127) {
                                        if false {
                                            v43 := GenS5{F0: v2}
                                            _ = v43
                                        }
                                        if (!(false && false)) {
                                            v44 := GenS3{F0: -1, F1: v13, F2: (("hello" + "a") + "hello")}
                                            for fi45 := 0; fi45 < 19; fi45++ {
                                                v2 = v2
                                                _ = len(v1)
                                                var v46 int32 = 0
                                                _ = fi45
                                                _ = v46
                                            }
                                            v47 := GenS3{F0: -1, F1: (^(v13)), F2: "\n\t"}
                                            _ = v44
                                            _ = v47
                                        } else {
                                            v13 = v13
                                            if (32767 < int16(v4)) {
                                                fmt.Println(v2)
                                                v48 := GenS6{F0: v13, F1: (v3 + v3)}
                                                v2 += v2
                                                _ = v48
                                            } else {
                                                switch fi42 {
                                                case 32768:
                                                    v41 = ((-2.25 / (v41 * (0.0 + ((1.5 * (((((v41 / ((3.14159 / 1.0) + v41)) - v41) + 1e10) - v41) / 1.0)) * v41)))) - (v41 / v41))
                                                case 2147483647:
                                                    var v49 int16 = int16(v2)
                                                    _ = v49
                                                case 0:
                                                    if (42 == 42) {
                                                        v29 += (0.0 - float64(v41))
                                                        switch ("foo bar" + "hello") {
                                                        case "foo bar":
                                                            var v50 uint64 = 0
                                                            _ = v50
                                                        case "hello":
                                                            for fi51 := 0; fi51 < 1; fi51++ {
                                                                v3 = (-(fi51))
                                                                v52 := GenS5{F0: v2}
                                                                _ = fi51
                                                                _ = v52
                                                            }
                                                        }
                                                        var v53 int = int(v36)
                                                        _ = v53
                                                    }
                                                }
                                            }
                                            for fi54 := 0; fi54 < 18; fi54++ {
                                                if (1e10 <= 1.5) {
                                                    v55 := GenS6{F0: v13, F1: 9223372036854775807}
                                                    if true {
                                                        if ((true || (-126 != 103)) && false) {
                                                            if true {
                                                                if len(v40) >= 2 {
                                                                    v40 = v40[1:len(v40)]
                                                                }
                                                                v56 := GenS7{F0: -123}
                                                                fmt.Println(v13)
                                                                _ = v56
                                                            } else {
                                                                if ((v2 >> ((255) & 63)) != 255) {
                                                                    v3 -= -563
                                                                }
                                                            }
                                                            if true {
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    if (((("" + "hello") + "hello") + ("foo bar" + "\n\t")) > ("" + ((("a" + ("" + "a")) + ("hello" + ("" + "foo bar"))) + "\n\t"))) {
                                                                        if (((1e10 + v41) - v41) <= 3.14159) {
                                                                            var v57 uint32 = v36
                                                                            var v58 uint32 = (v57 ^ v57)
                                                                            switch (fi54 % ((fi54) | 1)) {
                                                                            case 0:
                                                                                v41 = (v41 + v41)
                                                                            default:
                                                                                delete(v5, v57)
                                                                                v59 := GenS8{F0: -1, F1: true, F2: (v13 < 42)}
                                                                                v1[42] = "foo bar"
                                                                                v60 := GenS9{F0: 947}
                                                                                var v61 string = "foo bar"
                                                                                _ = v59
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                            _ = v57
                                                                            _ = v58
                                                                        } else {
                                                                            v29 = ((-2.25 - float64(v41)) / -2.25)
                                                                            fmt.Println(v36)
                                                                        }
                                                                        v4 -= fi42
                                                                        fmt.Println(v41, fi42, v36)
                                                                    }
                                                                case "a":
                                                                    var v62 int32 = 0
                                                                    _ = v62
                                                                case "hello":
                                                                    if false {
                                                                        v63 := GenS6{F0: 72, F1: (-(v4))}
                                                                        var v64 uint32 = 2147483647
                                                                        _ = v63
                                                                        _ = v64
                                                                    } else {
                                                                        if len(v40) >= 2 {
                                                                            v40 = v40[1:len(v40)]
                                                                        }
                                                                        if false {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                var v65 int32 = 42
                                                                                _ = v65
                                                                            default:
                                                                                v4 -= v4
                                                                            }
                                                                            v36 = (v36 + 673)
                                                                        } else {
                                                                            var v66 uint32 = (v36 % (((^(v36))) | 1))
                                                                            switch (v4 / (((-(v3))) | 1)) {
                                                                            case 32767:
                                                                                if ((v3 == (v4 - fi42)) || (false || false)) {
                                                                                    v41 -= v41
                                                                                } else {
                                                                                    v1[(^(v66))] = ("foo bar" + "hello")
                                                                                }
                                                                            case -434:
                                                                                if true {
                                                                                    var v67 uint32 = uint32(v2)
                                                                                    _ = v67
                                                                                } else {
                                                                                    var v68 uint16 = uint16(v2)
                                                                                    v3 -= (fi54 << ((1) & 63))
                                                                                    var v69 int16 = int16(v66)
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                }
                                                                            }
                                                                            _ = v66
                                                                        }
                                                                        v70 := GenS2{F0: 42, F1: 1}
                                                                        _ = v70
                                                                    }
                                                                default:
                                                                    var v71 float32 = 1.5
                                                                    _ = v71
                                                                }
                                                            } else {
                                                                v72 := GenS10{F0: (v36 - 2)}
                                                                v73 := GenS3{F0: 127, F1: (^(v13)), F2: "a"}
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            v29 -= v29
                                                        } else {
                                                            v29 = ((v29 * v29) + v29)
                                                            switch fi42 {
                                                            case 158:
                                                                v13 = (v13 / ((v13) | 1))
                                                            }
                                                            var v74 bool = true
                                                            if true {
                                                                delete(v1, (^(v36)))
                                                            } else {
                                                                if ((v36 - 42) > (v36 + v36)) {
                                                                    if (v29 == 3.14159) {
                                                                        v74 = (!((v29 * v29) >= (v29 / 3.14159)))
                                                                        v29 = -2.25
                                                                        var v75 int64 = 42
                                                                        v36 = v36
                                                                        v76 := GenS2{F0: -128, F1: 347}
                                                                        if false {
                                                                            var v77 int16 = 0
                                                                            _ = v77
                                                                        } else {
                                                                            if (889 > 1) {
                                                                                switch 32767 {
                                                                                case -369:
                                                                                    v2 += uint8(v36)
                                                                                case 4294967295:
                                                                                    v1[(v36 / (((v36 * v36)) | 1))] = ""
                                                                                case -9223372036854775808:
                                                                                    delete(v1, (v36 % ((v36) | 1)))
                                                                                default:
                                                                                    if ((922 > 42) && (469 == 295)) {
                                                                                        v4 -= 127
                                                                                        fmt.Println(v36, v36)
                                                                                    } else {
                                                                                        v78 := GenS8{F0: -128, F1: (!(-32768 == 42)), F2: v74}
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v79 uint32 = (v36 & (v36 * 2147483647))
                                                                                fmt.Println(v75)
                                                                                _ = v79
                                                                            }
                                                                            v40 = append(v40, (fi42 * int(v13)))
                                                                        }
                                                                        var v80 uint64 = 0
                                                                        _ = v75
                                                                        _ = v76
                                                                        _ = v80
                                                                    } else {
                                                                        var v81 uint = 942
                                                                        v3 = -712
                                                                        _ = len(v1)
                                                                        v41 += (float32(v29) - 1.5)
                                                                        switch ("a" + "\n\t") {
                                                                        case "hello":
                                                                            switch (-(fi42)) {
                                                                            case -780:
                                                                                var v82 uint = v81
                                                                                _ = v82
                                                                            case 32767:
                                                                                delete(v1, (v36 % ((v36) | 1)))
                                                                            case -494:
                                                                                fmt.Println(fi42)
                                                                            default:
                                                                                fmt.Println(fi42)
                                                                            }
                                                                        case "\n\t":
                                                                            if true {
                                                                                v83 := GenS9{F0: 525}
                                                                                _ = v83
                                                                            }
                                                                        case "":
                                                                            v74 = v74
                                                                            fmt.Println(fi42)
                                                                        default:
                                                                            var v84 bool = ("" >= "")
                                                                            _ = v84
                                                                        }
                                                                        v85 := GenS5{F0: v2}
                                                                        _ = v81
                                                                        _ = v85
                                                                    }
                                                                    var v86 int64 = -1
                                                                    v2 -= v2
                                                                    _ = v86
                                                                } else {
                                                                    fmt.Println(v3)
                                                                }
                                                                if v74 {
                                                                    v5[(v36 / (((v36 | 134)) | 1))] = false
                                                                }
                                                                var v87 int64 = 9223372036854775807
                                                                _ = v87
                                                            }
                                                            var v88 uint = 255
                                                            _ = v74
                                                            _ = v88
                                                        }
                                                    } else {
                                                        if (("\n\t" <= "") && (false && ((-(v4)) > 1))) {
                                                            v89 := GenS3{F0: 0, F1: v13, F2: "hello"}
                                                            var v90 int8 = 55
                                                            switch 9223372036854775807 {
                                                            case 2147483647:
                                                                v13 -= (v13 / ((v13) | 1))
                                                            case 128:
                                                                var v91 float64 = 1e10
                                                                _ = v91
                                                            case 32767:
                                                                if false {
                                                                    v4 -= fi54
                                                                    fmt.Println(v13, v13)
                                                                    switch ("hello" + "\n\t") {
                                                                    case "\n\t":
                                                                        delete(v5, v36)
                                                                    case "":
                                                                        v36 += (v36 * 2147483647)
                                                                    default:
                                                                        fmt.Println(fi42, fi42, v4)
                                                                    }
                                                                }
                                                            }
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        v13 -= uint16(v36)
                                                        var v92 float64 = v29
                                                        _ = v92
                                                    }
                                                    v13 = 32767
                                                    _ = v55
                                                } else {
                                                    v29 -= v29
                                                }
                                                if (1.5 > v29) {
                                                    var v93 int64 = 1
                                                    fmt.Println(fi54, v4, v29)
                                                    _ = v93
                                                } else {
                                                    if (42 < 1) {
                                                        v36 += 471
                                                    } else {
                                                        var v94 int16 = -32768
                                                        v40 = append(v40, v3)
                                                        _ = v94
                                                    }
                                                    switch (-(fi42)) {
                                                    case 250:
                                                        fmt.Println(fi42, v13)
                                                    }
                                                    var v95 uint16 = (^(v13))
                                                    if len(v40) > 0 {
                                                        v40[0] = int(v2)
                                                    }
                                                    _ = v95
                                                }
                                                _ = fi54
                                            }
                                        }
                                    }
                                    _ = fi42
                                }
                                _ = v41
                            } else {
                                switch 32767 {
                                case -667:
                                    var v96 string = ("hello" + "hello")
                                    _ = v96
                                case -178:
                                    if ((int32(v4) < -892) || ((0 < 18446744073709551615) || true)) {
                                        v29 -= v29
                                        v97 := GenS3{F0: -52, F1: (v13 & v13), F2: ""}
                                        v98 := GenS1{F0: (^(v2)), F1: true, F2: v4}
                                        _ = v97
                                        _ = v98
                                    } else {
                                        var v99 float32 = 1.5
                                        if true {
                                            v3 += v3
                                            var v100 float32 = 1e10
                                            _ = v100
                                        }
                                        _ = v99
                                    }
                                case 9223372036854775807:
                                    for fi101 := 0; fi101 < 10; fi101++ {
                                        for fi102 := 0; fi102 < 4; fi102++ {
                                            v103 := v5[1]
                                            switch ("foo bar" + "a") {
                                            case "\n\t":
                                                var v104 uint = 65535
                                                v105 := GenS11{F0: (v3 + (v4 + fi101)), F1: 32767}
                                                var v106 int32 = 0
                                                _ = v104
                                                _ = v105
                                                _ = v106
                                            default:
                                                v107 := GenS9{F0: 1}
                                                _ = v107
                                            }
                                            var v108 uint32 = v36
                                            _ = fi102
                                            _ = v103
                                            _ = v108
                                        }
                                        _ = fi101
                                    }
                                default:
                                    if len(v40) >= 2 {
                                        v40 = v40[1:len(v40)]
                                    }
                                }
                                switch 32768 {
                                case 62:
                                    fmt.Println(v2, v3)
                                case 1:
                                    if true {
                                        if true {
                                            if (true || (0 == 65535)) {
                                                fmt.Println(v3, v13, v2)
                                            } else {
                                                fmt.Println(v2, v13)
                                                v109 := GenS9{F0: 652}
                                                var v110 uint = 1
                                                _ = v109
                                                _ = v110
                                            }
                                            var v111 bool = false
                                            _ = v111
                                        }
                                        var v112 uint64 = 581
                                        _ = v112
                                    }
                                default:
                                    v29 = (v29 - v29)
                                }
                                v29 = 3.14159
                            }
                            _ = v40
                        }
                        var v113 uint32 = v36
                        v114 := GenS10{F0: v36}
                        _ = v113
                        _ = v114
                    } else {
                        _ = len(v5)
                        v115 := GenS1{F0: v2, F1: (!(!(32767 < v13))), F2: 618}
                        _ = v115
                    }
                    _ = v35
                    _ = v36
                }
                if (!(197 > 65535)) {
                    var v116 uint16 = v13
                    v29 = (1e10 * 1e10)
                    _ = v116
                } else {
                    fmt.Println(v4, v2, v4)
                }
                v3 -= 9223372036854775807
            }
        default:
            switch (("hello" + "\n\t") + "foo bar") {
            case "":
                v1[65535] = "a"
            case "foo bar":
                v13 = v13
            default:
                if (3.14159 == v29) {
                    for fi117 := 0; fi117 < 15; fi117++ {
                        v118 := GenS12{F0: 584, F1: (!false), F2: 12}
                        v29 -= v29
                        _ = fi117
                        _ = v118
                    }
                    v119 := GenS4{F0: genFn0(42, 82, 1), F1: true}
                    _ = v119
                }
            }
        }
        switch (v3 % ((1) | 1)) {
        case 0:
            var v120 int = (v3 >> ((2) & 63))
            _ = v120
        case 4294967295:
            if (false && (926 != 18446744073709551615)) {
                for fi121 := 0; fi121 < 18; fi121++ {
                    switch (fi121 << ((65535) & 63)) {
                    case -513:
                        v13 -= (v13 / (((v13 ^ v13)) | 1))
                    case 32768:
                        v3 += (v3 + v4)
                    case 127:
                        var v122 int8 = -1
                        _ = v122
                    }
                    var v123 int64 = genFn0(829, 42, 2)
                    _ = fi121
                    _ = v123
                }
                switch v4 {
                case 9223372036854775807:
                    var v124 int64 = -9223372036854775808
                    _ = v124
                case 0:
                    v125 := GenS13{F0: 1, F1: true, F2: -2147483648}
                    _ = v125
                case 127:
                    v13 = (v13 / ((v13) | 1))
                default:
                    v3 = v3
                }
            }
        default:
            var v126 float32 = float32(v29)
            _ = v126
        }
    }
    fmt.Println(v2, v4)
    v3 = 32768
    v29 += v29
    var v127 int16 = -888
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v127:", v127)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v127
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
