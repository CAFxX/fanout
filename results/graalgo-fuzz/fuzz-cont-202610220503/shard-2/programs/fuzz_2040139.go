package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 string; F1 uint; F2 int32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint8; F1 int8; F2 string }
type GenS8 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 uint16 }

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 65535
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v5 float64 = float64(v1)
    v6 := make(map[uint32]bool)
    v2 = 32767
    fmt.Println(v5)
    if ((4294967295 < 4294967295) && true) {
        if true {
            if (!(false && false)) {
                fmt.Println(v5, v1, v2)
                fmt.Println(v5)
            }
        } else {
            var v7 string = ("\n\t" + (("" + "hello") + ("a" + "foo bar")))
            var v8 bool = (282 >= 2147483647)
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 42, F1: v2, F2: 1}
        v5 = v5
        _ = v9
    }
    switch ("" + "") {
    case "":
        var v10 uint32 = uint32(v2)
        _ = v10
    case "\n\t":
        if false {
            switch -1 {
            case 127:
                switch "\n\t" {
                case "foo bar":
                    switch 686 {
                    case 32768:
                        fmt.Println(v1, v5)
                    }
                default:
                    v1 -= v1
                }
            case 0:
                if ((255 != 0) || (!((-1 == 1) || true))) {
                    _ = len(v4)
                }
            case -64:
                fmt.Println(v1)
            default:
                v11 := []int{-191}
                _ = v11
            }
        } else {
            fmt.Println(v5, v5, v5)
        }
    case "a":
        if (167 != 42) {
            var v12 int8 = 91
            _ = v12
        }
    default:
        fmt.Println(v1, v2, v2)
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            v2 -= v2
            v5 += v5
            fmt.Println(fi14, v1)
            _ = fi14
        }
        var v15 uint16 = (v2 | (v2 % (((v2 & v2)) | 1)))
        fmt.Println(v1)
        _ = fi13
        _ = v15
    }
    v16 := GenS1{F0: 440, F1: v2, F2: 42}
    fmt.Println(v5, v2)
    v17 := []int{281, -2147483648, 699}
    _ = len(v4)
    fmt.Println(v2)
    switch "" {
    case "":
        var v18 float32 = float32(v5)
        _ = v18
    case "hello":
        var v19 bool = true
        _ = v19
    case "a":
        if true {
            v20 := GenS2{F0: "\n\t", F1: 4294967295, F2: 0}
            _ = v20
        }
    }
    v21 := GenS3{F0: false, F1: 0, F2: v2}
    if false {
        fmt.Println(v2)
        v5 = 3.14159
    } else {
        if ((v5 - v5) >= v5) {
            if false {
                var v22 int64 = -105
                if ((0 >= 2) || false) {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        for fi24 := 0; fi24 < 17; fi24++ {
                            var v25 float64 = v5
                            _ = fi24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 8; fi26++ {
                            fmt.Println(v2, v1)
                            _ = fi26
                        }
                        var v27 uint64 = 1
                        _ = fi23
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 14; fi28++ {
                        for fi29 := 0; fi29 < 16; fi29++ {
                            var v30 uint8 = 42
                            switch fi29 {
                            case 4294967295:
                                v2 += (v2 + (v2 >> ((18446744073709551615) & 63)))
                            case 127:
                                v31 := GenS4{F0: (^(v2)), F1: 4294967295}
                                _ = v31
                            }
                            if false {
                                v32 := GenS4{F0: v2, F1: 4294967295}
                                if len(v17) > 0 {
                                    v17[0] = (fi29 / (((fi29 | 127)) | 1))
                                }
                                _ = v32
                            } else {
                                var v33 uint = 65535
                                v30 -= 2
                                _ = v33
                            }
                            _ = fi29
                            _ = v30
                        }
                        switch "a" {
                        case "hello":
                            for fi34 := 0; fi34 < 2; fi34++ {
                                if (-1 < 2147483647) {
                                    v22 = v22
                                } else {
                                    delete(v3, -585)
                                }
                                v1 += v1
                                var v35 int16 = -860
                                _ = fi34
                                _ = v35
                            }
                        case "a":
                            var v36 string = "hello"
                            _ = v36
                        }
                        _ = fi28
                    }
                    var v37 float32 = v1
                    _ = v37
                }
                _ = v22
            } else {
                var v38 uint16 = v2
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 uint = 255
                    for fi41 := 0; fi41 < 8; fi41++ {
                        v38 = 0
                        switch 711 {
                        case 2147483647:
                            switch -849 {
                            case 808:
                                if true {
                                    if ((("" + "") != "\n\t") || false) {
                                        v42 := GenS5{F0: ("foo bar" + "")}
                                        if false {
                                            var v43 int = fi39
                                            var v44 int64 = -868
                                            _ = len(v17)
                                            _ = v43
                                            _ = v44
                                        } else {
                                            fmt.Println(v40)
                                            fmt.Println(v5)
                                        }
                                        _ = v42
                                    } else {
                                        var v45 int64 = -1
                                        _ = v45
                                    }
                                    v17 = append(v17, (fi41 / ((fi39) | 1)))
                                } else {
                                    if ((1 != 113) && false) {
                                        var v46 float32 = (((v1 / 1.0) / ((-2.25 - (float32((-2.25 + v5)) * ((v1 * v1) * ((v1 * -2.25) - float32(v5))))) * (v1 * (v1 + -2.25)))) * (1e10 + float32(v5)))
                                        _ = v46
                                    } else {
                                        v5 = -2.25
                                        v2 -= v38
                                        v40 -= 0
                                    }
                                }
                            }
                        }
                        _ = fi41
                    }
                    for fi47 := 0; fi47 < 9; fi47++ {
                        if (4 == 128) {
                            v40 = (v40 >> (((v40 ^ (v40 % (((v40 - 160)) | 1)))) & 63))
                            var v48 int16 = -949
                            _ = v48
                        } else {
                            v49 := GenS1{F0: (fi39 % (((fi47 % (((fi47 & fi39)) | 1))) | 1)), F1: v2, F2: -851}
                            if ((v40 >= (v40 * (v40 - (v40 & (v40 % (((v40 + uint(v2))) | 1)))))) && true) {
                                var v50 uint8 = 203
                                _ = v50
                            }
                            _ = v49
                        }
                        v2 += v38
                        _ = fi47
                    }
                    _ = fi39
                    _ = v40
                }
                switch (("a" + "hello") + "\n\t") {
                case "\n\t":
                    var v51 bool = (((0.1 + v5) * (v5 + (v5 - v5))) >= v5)
                    _ = v51
                }
                _ = v38
            }
            if (!(1e10 == 0.0)) {
                if (843 <= 1) {
                    for fi52 := 0; fi52 < 19; fi52++ {
                        _ = len(v17)
                        v53 := v6[uint32(v2)]
                        _ = fi52
                        _ = v53
                    }
                    var v54 uint64 = 1
                    v17 = append(v17, 1)
                    _ = v54
                } else {
                    if ((-9223372036854775808 <= 42) && (2 <= 2)) {
                        if false {
                            fmt.Println(v1, v5)
                        }
                    } else {
                        v2 -= 589
                        var v55 int = 1
                        _ = v55
                    }
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    if false {
                        if true {
                            if false {
                                for fi56 := 0; fi56 < 4; fi56++ {
                                    var v57 uint8 = 2
                                    for fi58 := 0; fi58 < 17; fi58++ {
                                        var v59 int = fi56
                                        switch "foo bar" {
                                        case "hello":
                                            var v60 int8 = -1
                                            _ = v60
                                        case "foo bar":
                                            var v61 uint64 = 735
                                            var v62 uint64 = v61
                                            var v63 float32 = v1
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        case "a":
                                            var v64 bool = false
                                            v2 += 934
                                            if len(v17) > 0 {
                                                v17[0] = (^(fi56))
                                            }
                                            _ = v64
                                        }
                                        v65 := GenS2{F0: (("hello" + (("a" + "hello") + ("a" + ("foo bar" + ("a" + "a"))))) + "foo bar"), F1: 4294967295, F2: 0}
                                        _ = fi58
                                        _ = v59
                                        _ = v65
                                    }
                                    _ = fi56
                                    _ = v57
                                }
                            } else {
                                v66 := GenS4{F0: (v2 / (((v2 >> ((2) & 63))) | 1)), F1: 1}
                                var v67 uint64 = 843
                                if (true && (-683 <= -340)) {
                                    fmt.Println(v5, v2)
                                    if false {
                                        for fi68 := 0; fi68 < 4; fi68++ {
                                            var v69 float32 = float32(v5)
                                            if ((("a" + ("hello" + ("hello" + "hello"))) + "") > "\n\t") {
                                                for fi70 := 0; fi70 < 12; fi70++ {
                                                    var v71 uint32 = 854
                                                    v67 += 1
                                                    v2 = (v2 >> ((1) & 63))
                                                    _ = fi70
                                                    _ = v71
                                                }
                                                if true {
                                                    v72 := GenS1{F0: 900, F1: 2, F2: -599}
                                                    if false {
                                                        for fi73 := 0; fi73 < 14; fi73++ {
                                                            var v74 int64 = 0
                                                            if (-2147483648 == 0) {
                                                                if (!(-2147483648 > 921)) {
                                                                    switch ("" + "a") {
                                                                    case "":
                                                                        v69 -= ((float32((float64(1.5) + (float64((v69 + float32(1.5))) / float64((v1 / ((v1 * (v1 + 0.1)) - v1)))))) + v69) / float32((v5 / v5)))
                                                                    case "a":
                                                                        fmt.Println(v69, v69, v69)
                                                                    default:
                                                                        if true {
                                                                            switch 127 {
                                                                            case 32768:
                                                                                v69 -= 0.1
                                                                            }
                                                                            var v75 int32 = -2147483648
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                    fmt.Println(v74)
                                                                    if len(v17) > 0 {
                                                                        v17[0] = int(v74)
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v67, v1, fi73)
                                                            v76 := v4[2147483647]
                                                            if (v67 > (v67 * (v67 / (((v67 * v67)) | 1)))) {
                                                                var v77 int32 = -817
                                                                v78 := GenS5{F0: (("hello" + ("hello" + (("foo bar" + (("" + "hello") + ("foo bar" + ("hello" + ("\n\t" + ""))))) + "a"))) + (("\n\t" + ("foo bar" + "a")) + ("hello" + ("hello" + "hello"))))}
                                                                _ = v77
                                                                _ = v78
                                                            } else {
                                                                if true {
                                                                    v69 = 1e10
                                                                    v2 -= (v2 + (v2 + v2))
                                                                } else {
                                                                    if (!(v74 >= (v74 - v74))) {
                                                                        if (!false) {
                                                                            v79 := v6[666]
                                                                            fmt.Println(v1, fi68, v5)
                                                                            var v80 uint = 18446744073709551615
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                    } else {
                                                                        v76 = v76
                                                                    }
                                                                    v76 = v76
                                                                    v2 -= v2
                                                                }
                                                            }
                                                            _ = fi73
                                                            _ = v74
                                                            _ = v76
                                                        }
                                                        v81 := GenS3{F0: (float64(v1) < (v5 / v5)), F1: -1, F2: 29}
                                                        _ = v81
                                                    } else {
                                                        v82 := v4[0]
                                                        v83 := GenS6{F0: uint32(v67), F1: 1}
                                                        v84 := GenS7{F0: 169, F1: 42, F2: (("hello" + (("foo bar" + "\n\t") + "")) + ("" + ""))}
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    for fi85 := 0; fi85 < 13; fi85++ {
                                                        v86 := GenS2{F0: ("hello" + ("a" + "a")), F1: 255, F2: -971}
                                                        _ = len(v17)
                                                        _ = fi85
                                                        _ = v86
                                                    }
                                                    _ = v72
                                                }
                                                if (!false) {
                                                    var v87 bool = false
                                                    if (1 != 0) {
                                                        v69 += (float32(((-2.25 / v5) / (v5 - v5))) + v1)
                                                        v17 = append(v17, -1)
                                                    }
                                                    _ = v87
                                                } else {
                                                    var v88 uint = 8
                                                    v88 -= uint(v2)
                                                    _ = v88
                                                }
                                            }
                                            var v89 uint64 = (v67 | (v67 | 139))
                                            _ = fi68
                                            _ = v69
                                            _ = v89
                                        }
                                        if true {
                                            if true {
                                                var v90 uint64 = 4294967295
                                                _ = v90
                                            }
                                        }
                                        for fi91 := 0; fi91 < 12; fi91++ {
                                            v2 = (v2 | (^(v2)))
                                            var v92 uint = 47
                                            fmt.Println(v2, v92, v67)
                                            _ = fi91
                                            _ = v92
                                        }
                                    }
                                    _ = len(v4)
                                }
                                _ = v66
                                _ = v67
                            }
                        }
                        v3[42] = 2147483647
                        if false {
                            var v93 float64 = v5
                            v93 += v5
                            _ = v93
                        } else {
                            var v94 uint = 796
                            for fi95 := 0; fi95 < 19; fi95++ {
                                if len(v17) > 0 {
                                    v17[0] = (fi95 / ((688) | 1))
                                }
                                switch fi95 {
                                case 432:
                                    for fi96 := 0; fi96 < 17; fi96++ {
                                        fmt.Println(v2, fi96, v1)
                                        _ = fi96
                                    }
                                case -86:
                                    v2 = v2
                                }
                                var v97 int = fi95
                                _ = fi95
                                _ = v97
                            }
                            _ = v94
                        }
                    }
                }
                _ = len(v17)
                if (-1 <= 42) {
                    fmt.Println(v5, v2)
                    v98 := GenS8{F0: 42, F1: -296, F2: 2}
                    _ = v98
                } else {
                    v99 := GenS1{F0: -942, F1: v2, F2: 850}
                    _ = v99
                }
            } else {
                v1 += v1
                for fi100 := 0; fi100 < 12; fi100++ {
                    if ((0.0 / (float32((0.1 - 3.14159)) + ((v1 - 1e10) / (float32(1.5) - ((v1 - v1) / v1))))) == float32(-2.25)) {
                        var v101 bool = true
                        _ = v101
                    }
                    for fi102 := 0; fi102 < 20; fi102++ {
                        switch fi102 {
                        case -461:
                            if false {
                                if false {
                                    v1 = (0.0 - v1)
                                    _ = len(v17)
                                    fmt.Println(v1)
                                } else {
                                    fmt.Println(v5)
                                    v1 += 3.14159
                                }
                                v2 = (^(v2))
                            }
                        case 2147483647:
                            v2 -= (^(v2))
                        default:
                            if false {
                                if len(v17) > 0 {
                                    v17[0] = fi102
                                }
                                v1 = ((0.0 - v1) * v1)
                                if false {
                                    v1 = v1
                                } else {
                                    var v103 int64 = 1
                                    _ = v103
                                }
                            }
                        }
                        _ = len(v3)
                        v104 := GenS9{F0: true, F1: "a"}
                        _ = fi102
                        _ = v104
                    }
                    _ = fi100
                }
                v2 = v2
            }
            for fi105 := 0; fi105 < 7; fi105++ {
                if len(v17) > 0 {
                    v17[0] = (fi105 * fi105)
                }
                _ = fi105
            }
        } else {
            if true {
                if true {
                    v17 = append(v17, 1)
                }
                if (("hello" > "foo bar") || true) {
                    var v106 uint = 1
                    for fi107 := 0; fi107 < 12; fi107++ {
                        v1 -= float32(v5)
                        v106 = (v106 & (v106 + uint(v2)))
                        _ = fi107
                    }
                    _ = v106
                } else {
                    var v108 int = 32767
                    fmt.Println(v2, v108, v108)
                    _ = v108
                }
                var v109 string = "foo bar"
                _ = v109
            } else {
                switch "a" {
                case "\n\t":
                    fmt.Println(v5, v1)
                }
            }
            if false {
                v110 := GenS6{F0: 65535, F1: 301}
                v17 = append(v17, 127)
                if false {
                    if true {
                        if (!(937 <= 0)) {
                            if ("hello" < ((("" + ("hello" + "\n\t")) + ("" + "a")) + ("foo bar" + "a"))) {
                                switch "a" {
                                case "hello":
                                    var v111 uint = 65535
                                    _ = v111
                                case "":
                                    v112 := GenS3{F0: (0 != -1), F1: 1000, F2: v2}
                                    _ = v112
                                default:
                                    fmt.Println(v2, v1, v5)
                                }
                            }
                            fmt.Println(v5, v2, v2)
                        } else {
                            fmt.Println(v5)
                            var v113 string = "foo bar"
                            _ = v113
                        }
                        var v114 uint64 = 2
                        _ = v114
                    }
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                } else {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    var v115 int64 = 42
                    _ = v115
                }
                _ = v110
            } else {
                for fi116 := 0; fi116 < 6; fi116++ {
                    if true {
                        fmt.Println(v1, fi116)
                        v117 := GenS10{F0: v2}
                        _ = v117
                    }
                    if (!(-1 < 0)) {
                        v118 := GenS2{F0: "\n\t", F1: 4294967295, F2: 426}
                        _ = len(v17)
                        _ = v118
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v119 := GenS9{F0: true, F1: "foo bar"}
                        _ = v119
                    default:
                        for fi120 := 0; fi120 < 2; fi120++ {
                            if true {
                                var v121 string = ((((("" + "") + ("\n\t" + (("" + (("\n\t" + "") + ("\n\t" + "hello"))) + (((("hello" + (("foo bar" + "foo bar") + "foo bar")) + "foo bar") + "a") + ("foo bar" + ("\n\t" + "")))))) + (((("" + "hello") + "foo bar") + "foo bar") + "")) + "hello") + "hello")
                                v122 := GenS3{F0: true, F1: 173, F2: (v2 + v2)}
                                _ = v121
                                _ = v122
                            } else {
                                var v123 uint8 = 255
                                _ = v123
                            }
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            v5 = (v5 - (((float64((float32(-2.25) / v1)) / 1e10) - v5) * -2.25))
                            _ = fi120
                        }
                    }
                    _ = fi116
                }
                var v124 int8 = -1
                if true {
                    var v125 string = ("a" + "\n\t")
                    _ = v125
                }
                _ = v124
            }
            v1 -= v1
        }
    }
    v126 := GenS8{F0: 59, F1: -1, F2: 65535}
    if (156 < 627) {
        v17 = append(v17, -830)
        fmt.Println(v1)
        v4[65535] = uint32(v2)
    }
    if true {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        var v127 uint8 = 127
        _ = v127
    } else {
        var v128 float64 = v5
        _ = v128
    }
    var v129 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v126
    _ = v129
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
