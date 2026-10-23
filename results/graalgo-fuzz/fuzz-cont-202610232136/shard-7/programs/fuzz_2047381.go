package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint }
type GenS5 struct { F0 string; F1 int; F2 int32 }
type GenS6 struct { F0 string; F1 bool; F2 int64 }
type GenS7 struct { F0 uint; F1 int32; F2 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 string; F1 uint16; F2 uint16 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 uint16; F1 uint64 }
type GenS14 struct { F0 int32; F1 uint64 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint32; F1 int }
type GenS17 struct { F0 string; F1 int32; F2 int32 }
type GenS18 struct { F0 uint32 }
func genFn0(p0 int, p1 int32, p2 uint16) string {
    v7 := []int{p0}
    var v8 uint64 = 18446744073709551615
    p0 = (p0 & (p0 - p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := []uint32{4294967295, 536, 65535, 2147483648}
    v2 := []int{32768, 32768}
    v3 := []int{317, -1, 2147483647}
    var v4 float64 = float64(1.5)
    v5 := make(map[int]string)
    var v6 bool = false
    switch 1 {
    case 32767:
        var v9 string = "\n\t"
        _ = v9
    case 4294967295:
        v4 += float64(1e10)
    case 0:
        v4 += (((1.5 * v4) - 0.1) / -2.25)
    default:
        var v10 uint = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int32 = -2147483648
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        if v6 {
            switch (("" + "") + "") {
            case "\n\t":
                switch (fi13 - fi13) {
                case 1:
                    _ = len(v2)
                case 0:
                    var v14 int32 = -47
                    _ = v14
                }
            case "a":
                if len(v3) > 0 {
                    v3[0] = fi13
                }
            }
            v4 = (v4 * 0.0)
            switch fi13 {
            case 32768:
                for fi15 := 0; fi15 < 7; fi15++ {
                    v4 = 1.5
                    _ = fi15
                }
            default:
                if (!(0 <= 42)) {
                    var v16 int32 = -944
                    _ = v16
                } else {
                    switch "foo bar" {
                    case "hello":
                        v2 = append(v2, 4294967295)
                    case "foo bar":
                        v17 := GenS1{F0: 42}
                        _ = v17
                    case "\n\t":
                        switch ("a" + "") {
                        case "foo bar":
                            v18 := GenS1{F0: -436}
                            _ = v18
                        default:
                            var v19 int32 = -404
                            _ = v19
                        }
                    default:
                        var v20 uint16 = 670
                        _ = v20
                    }
                    if ((v6 || (!(!(1 == -1)))) && false) {
                        if len(v1) > 0 {
                            v1[0] = 905
                        }
                        if (9223372036854775807 <= (-(fi13))) {
                            v21 := GenS1{F0: -9223372036854775808}
                            v6 = v6
                            if (-128 == -128) {
                                fmt.Println(v4)
                                var v22 int8 = 0
                                var v23 int = (fi13 & (fi13 + fi13))
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        }
                    } else {
                        v24 := GenS1{F0: 810}
                        if (896 != 1) {
                            var v25 uint = 1
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 5; fi26++ {
                                fmt.Println(v4, fi26)
                                v6 = v6
                                switch "foo bar" {
                                case "":
                                    v27 := GenS2{F0: 128, F1: 76}
                                    _ = v27
                                case "a":
                                    delete(v5, fi26)
                                default:
                                    switch "hello" {
                                    case "hello":
                                        var v28 uint8 = 0
                                        _ = v28
                                    case "a":
                                        switch "" {
                                        case "foo bar":
                                            switch "hello" {
                                            case "hello":
                                                if (65535 < 550) {
                                                    _ = len(v2)
                                                    v1 = append(v1, 0)
                                                }
                                            case "a":
                                                v29 := GenS2{F0: 38, F1: 86}
                                                var v30 string = ""
                                                if v6 {
                                                    var v31 int8 = 1
                                                    _ = v31
                                                } else {
                                                    var v32 int8 = -7
                                                    if (2 > 0) {
                                                        v33 := GenS3{F0: 94}
                                                        if (v32 < 127) {
                                                            if (639 >= 2) {
                                                                var v34 int64 = 9223372036854775807
                                                                v6 = ((v34 != -1) && ("foo bar" > "foo bar"))
                                                                v35 := GenS2{F0: 0, F1: 0}
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                            switch ((v30 + "") + (v30 + ("foo bar" + v30))) {
                                                            case "":
                                                                var v36 bool = (!v6)
                                                                _ = v36
                                                            default:
                                                                switch 4294967295 {
                                                                case 823:
                                                                    var v37 uint8 = 127
                                                                    _ = v37
                                                                case 9223372036854775807:
                                                                    v6 = v6
                                                                case 2147483647:
                                                                    var v38 float64 = ((v4 + (1.5 / 1.0)) - 3.14159)
                                                                    _ = v38
                                                                default:
                                                                    if v6 {
                                                                        v6 = ((v6 && v6) && (-863 == -1))
                                                                        v39 := GenS1{F0: -9223372036854775808}
                                                                        if (257 == 2147483647) {
                                                                            fmt.Println(v30, v32, v4)
                                                                            v6 = true
                                                                        } else {
                                                                            v4 -= (float64(1.5) - (-2.25 * float64(1e10)))
                                                                        }
                                                                        _ = v39
                                                                    }
                                                                }
                                                            }
                                                            var v40 int32 = -2147483648
                                                            if len(v1) > 0 {
                                                                v1[0] = 2147483647
                                                            }
                                                            _ = v40
                                                        }
                                                        var v41 int32 = 1
                                                        if false {
                                                            var v42 string = v30
                                                            v6 = (9223372036854775807 == -918)
                                                            _ = v42
                                                        }
                                                        fmt.Println(fi13, v30)
                                                        _ = v33
                                                        _ = v41
                                                    }
                                                    fmt.Println(v4, v30, v32)
                                                    v5[fi26] = v30
                                                    v43 := GenS1{F0: 461}
                                                    _ = v32
                                                    _ = v43
                                                }
                                                _ = v29
                                                _ = v30
                                            case "":
                                                v6 = ("foo bar" <= (((("foo bar" + (((("a" + ("hello" + ("" + ("foo bar" + "")))) + ("hello" + "")) + ("foo bar" + ("hello" + "a"))) + "a")) + "") + "a") + ""))
                                            default:
                                                fmt.Println(v6, fi13)
                                            }
                                        case "\n\t":
                                            v44 := GenS4{F0: 0, F1: 1, F2: 42}
                                            _ = v44
                                        case "":
                                            v4 -= ((v4 - v4) / ((v4 * v4) / v4))
                                        }
                                    case "\n\t":
                                        if true {
                                            var v45 uint = 255
                                            _ = v45
                                        } else {
                                            v4 += ((1.5 / 1.5) - 0.1)
                                            fmt.Println(fi26, fi26, fi26)
                                            fmt.Println(fi13)
                                        }
                                    default:
                                        if v6 {
                                            v5[(fi13 & fi13)] = "\n\t"
                                        }
                                    }
                                }
                                if (v6 || v6) {
                                    var v46 uint8 = 68
                                    _ = v46
                                } else {
                                    v6 = v6
                                }
                                _ = fi26
                            }
                        }
                        v47 := GenS5{F0: ("" + "\n\t"), F1: fi13, F2: -876}
                        _ = v24
                        _ = v47
                    }
                    var v48 uint32 = 519
                    _ = v48
                }
            }
        }
        var v49 string = "\n\t"
        _ = fi13
        _ = v49
    }
    v6 = v6
    switch 1 {
    case 128:
        if (42 == 642) {
            var v50 uint = 643
            _ = v50
        } else {
            if true {
                fmt.Println(v4)
                if (74 != -1) {
                    v51 := GenS4{F0: 65535, F1: 191, F2: 0}
                    var v52 int64 = 0
                    _ = v51
                    _ = v52
                }
                var v53 uint8 = 83
                _ = v53
            } else {
                var v54 uint16 = 255
                if (42 < -1) {
                    var v55 uint16 = (v54 - v54)
                    v54 = (v54 ^ 951)
                    _ = v55
                } else {
                    for fi56 := 0; fi56 < 15; fi56++ {
                        fmt.Println(v4, fi56, fi56)
                        v4 -= ((v4 - v4) / v4)
                        var v57 string = (("" + "\n\t") + "a")
                        _ = fi56
                        _ = v57
                    }
                    v5[128] = "hello"
                    delete(v5, 32767)
                }
                for fi58 := 0; fi58 < 18; fi58++ {
                    v54 = v54
                    if v6 {
                        fmt.Println(v54, fi58)
                        if (4294967295 < 42) {
                            for fi59 := 0; fi59 < 20; fi59++ {
                                var v60 int8 = 10
                                _ = fi59
                                _ = v60
                            }
                            fmt.Println(v6, fi58)
                            v61 := GenS4{F0: 32767, F1: 2, F2: 837}
                            _ = v61
                        } else {
                            v62 := GenS5{F0: "hello", F1: fi58, F2: 42}
                            if ((v54 | (v54 / ((v54) | 1))) >= (v54 / (((v54 / ((v54) | 1))) | 1))) {
                                fmt.Println(v6, v54, v6)
                                if (429 >= 4294967295) {
                                    var v63 int = (-(fi58))
                                    _ = v63
                                }
                                switch ("hello" + ("\n\t" + ((("a" + (((("" + (("\n\t" + ("a" + "foo bar")) + "\n\t")) + "a") + "") + "foo bar")) + ("a" + "a")) + ("hello" + ("" + "foo bar"))))) {
                                case "foo bar":
                                    var v64 uint32 = 42
                                    _ = v64
                                case "\n\t":
                                    var v65 float64 = ((v4 / ((0.1 * (v4 + v4)) + 3.14159)) / float64((1.5 + (0.1 / (0.1 * ((0.0 * (0.0 / 0.1)) * (((float32(((v4 - (1e10 * -2.25)) * v4)) + 1.5) + (0.0 - 0.0)) * float32(((0.1 + v4) - 1.5)))))))))
                                    _ = v65
                                case "a":
                                    var v66 int64 = 42
                                    _ = v66
                                default:
                                    var v67 string = ("foo bar" + "\n\t")
                                    _ = v67
                                }
                            } else {
                                fmt.Println(v4, v6)
                                v6 = (((float64((0.0 - (1.5 * -2.25))) == v4) || true) && (false || (-1 < -1)))
                            }
                            _ = v62
                        }
                        v68 := GenS6{F0: "a", F1: v6, F2: -1}
                        _ = v68
                    } else {
                        var v69 int8 = 3
                        var v70 uint8 = 255
                        fmt.Println(v4, v4, fi58)
                        _ = v69
                        _ = v70
                    }
                    switch -1 {
                    case 2147483647:
                        for fi71 := 0; fi71 < 2; fi71++ {
                            switch "" {
                            case "foo bar":
                                switch "" {
                                case "a":
                                    var v72 string = ("" + ("a" + "\n\t"))
                                    _ = v72
                                case "hello":
                                    if v6 {
                                        v4 = v4
                                        v5[(fi71 % ((890) | 1))] = ""
                                        fmt.Println(fi58, fi71)
                                        if (!(116 <= v54)) {
                                            var v73 int16 = -1
                                            var v74 float64 = (v4 - v4)
                                            _ = v73
                                            _ = v74
                                        }
                                    } else {
                                        var v75 uint = uint(v54)
                                        var v76 string = ""
                                        fmt.Println(v54)
                                        _ = v75
                                        _ = v76
                                    }
                                }
                                if v6 {
                                    v5[9223372036854775807] = "hello"
                                    if (!(v6 && v6)) {
                                        v77 := GenS6{F0: "foo bar", F1: (4294967295 == 2), F2: 42}
                                        if (v54 > 989) {
                                            var v78 int = (fi58 + fi71)
                                            var v79 float32 = (0.1 / 1.0)
                                            _ = v78
                                            _ = v79
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 32767
                                            }
                                        }
                                        v4 -= v4
                                        _ = v77
                                    } else {
                                        var v80 int8 = -86
                                        v81 := GenS7{F0: 4294967295, F1: 868, F2: v6}
                                        var v82 uint = 240
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    }
                                }
                            default:
                                v54 -= v54
                            }
                            v4 = (v4 * ((float64(0.1) / v4) / v4))
                            _ = fi71
                        }
                    case 42:
                        v83 := GenS5{F0: ("foo bar" + "\n\t"), F1: fi58, F2: 42}
                        _ = v83
                    default:
                        if (!false) {
                            for fi84 := 0; fi84 < 5; fi84++ {
                                switch (fi58 >> ((42) & 63)) {
                                case 42:
                                    if (0 >= -442) {
                                        var v85 bool = (!false)
                                        _ = v85
                                    } else {
                                        _ = len(v3)
                                        fmt.Println(v6, fi84, fi84)
                                        var v86 bool = v6
                                        _ = v86
                                    }
                                case 150:
                                    delete(v5, int(v54))
                                default:
                                    var v87 uint32 = 0
                                    _ = v87
                                }
                                delete(v5, (fi58 << ((370) & 63)))
                                var v88 float32 = (((1e10 + (float32(1.5) + 3.14159)) - (3.14159 + (3.14159 / 1e10))) * ((1.5 + -2.25) + 0.0))
                                _ = fi84
                                _ = v88
                            }
                        } else {
                            var v89 uint64 = 42
                            var v90 bool = v6
                            fmt.Println(fi58, v4, v89)
                            _ = v89
                            _ = v90
                        }
                    }
                    _ = fi58
                }
                _ = v54
            }
        }
    case 634:
        v6 = v6
    }
    v91 := GenS7{F0: 246, F1: 42, F2: true}
    if ((18446744073709551615 == 49) || (0 >= 1)) {
        fmt.Println(v4, v6)
        switch "\n\t" {
        case "":
            var v92 float32 = 1.5
            _ = v92
        case "\n\t":
            var v93 bool = v6
            _ = v93
        case "a":
            if v6 {
                var v94 uint16 = 0
                fmt.Println(v94)
                fmt.Println(v6, v94)
                _ = v94
            }
        default:
            switch 32768 {
            case 42:
                switch ("a" + "\n\t") {
                case "a":
                    switch -64 {
                    case 32767:
                        var v95 uint32 = 2147483647
                        _ = v95
                    case -1:
                        switch "hello" {
                        case "hello":
                            v96 := GenS6{F0: "foo bar", F1: v6, F2: 0}
                            _ = v96
                        case "a":
                            _ = len(v3)
                        default:
                            v97 := v5[-9223372036854775808]
                            _ = v97
                        }
                    case 128:
                        var v98 int8 = 42
                        _ = v98
                    }
                case "":
                    v4 -= (v4 * v4)
                case "\n\t":
                    v6 = false
                }
            case 32768:
                v4 -= v4
            }
        }
        v4 += v4
    } else {
        var v99 int32 = 0
        v6 = v6
        if v6 {
            fmt.Println(v99, v99, v4)
            for fi100 := 0; fi100 < 2; fi100++ {
                v101 := GenS8{F0: 30}
                _ = fi100
                _ = v101
            }
        } else {
            if (((0.1 / float32(3.14159)) * float32(v4)) == 3.14159) {
                var v102 int8 = 1
                _ = v102
            }
        }
        _ = v99
    }
    fmt.Println(v6)
    var v103 bool = v6
    if v6 {
        var v104 uint = 2
        switch "a" {
        case "hello":
            var v105 string = "hello"
            _ = v105
        case "\n\t":
            switch "\n\t" {
            case "a":
                switch 9223372036854775807 {
                case -410:
                    var v106 uint32 = 589
                    _ = v106
                case 1:
                    var v107 float32 = 3.14159
                    _ = v107
                case -9223372036854775808:
                    switch -74 {
                    case 2147483647:
                        fmt.Println(v6)
                    default:
                        var v108 string = "hello"
                        _ = v108
                    }
                default:
                    var v109 uint32 = 0
                    _ = v109
                }
            default:
                switch 683 {
                case -2147483648:
                    for fi110 := 0; fi110 < 1; fi110++ {
                        for fi111 := 0; fi111 < 10; fi111++ {
                            var v112 int8 = 1
                            _ = fi111
                            _ = v112
                        }
                        _ = fi110
                    }
                case -537:
                    for fi113 := 0; fi113 < 12; fi113++ {
                        v104 -= (^(v104))
                        v103 = (-2.25 == (-2.25 * (-2.25 - (1e10 * 1.5))))
                        _ = fi113
                    }
                }
            }
        case "foo bar":
            switch ("\n\t" + "\n\t") {
            case "":
                var v114 uint = v104
                _ = v114
            case "foo bar":
                v115 := GenS5{F0: "foo bar", F1: 1, F2: -659}
                _ = v115
            default:
                var v116 uint = 1
                _ = v116
            }
        default:
            if v6 {
                if ((!v6) && v103) {
                    for fi117 := 0; fi117 < 14; fi117++ {
                        switch (("hello" + ("hello" + "foo bar")) + "\n\t") {
                        case "hello":
                            v6 = false
                        default:
                            if (-843 >= 1) {
                                if (!false) {
                                    v4 = ((-2.25 + 0.1) * (v4 * v4))
                                } else {
                                    if (!v6) {
                                        for fi118 := 0; fi118 < 14; fi118++ {
                                            if v6 {
                                                delete(v5, fi118)
                                                if v6 {
                                                    v119 := GenS9{F0: 815, F1: (!true)}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v120 int64 = 1
                                                    _ = v119
                                                    _ = v120
                                                } else {
                                                    var v121 float64 = v4
                                                    _ = v121
                                                }
                                                fmt.Println(v6, fi117)
                                            } else {
                                                switch fi118 {
                                                case -9223372036854775808:
                                                    if (18446744073709551615 < 0) {
                                                        switch fi118 {
                                                        case -673:
                                                            if (v103 || (v104 >= v104)) {
                                                                if (((65535 <= 2) && false) && (v103 || v103)) {
                                                                    var v122 uint32 = 42
                                                                    v4 -= v4
                                                                    var v123 float64 = 1e10
                                                                    _ = v122
                                                                    _ = v123
                                                                }
                                                                var v124 int = (fi117 << ((255) & 63))
                                                                v104 = v104
                                                                _ = v124
                                                            } else {
                                                                v4 -= ((float64(3.14159) * (-2.25 / ((v4 + float64(1.5)) * v4))) * v4)
                                                            }
                                                        default:
                                                            if (0 > 163) {
                                                                v104 += 4294967295
                                                            }
                                                        }
                                                        v4 -= 1e10
                                                        if (0.0 < (1e10 / 0.1)) {
                                                            fmt.Println(v6)
                                                        } else {
                                                            fmt.Println(fi117, fi117)
                                                            switch fi118 {
                                                            case -9223372036854775808:
                                                                var v125 int8 = int8(fi117)
                                                                _ = v125
                                                            case 0:
                                                                var v126 bool = v6
                                                                _ = v126
                                                            default:
                                                                fmt.Println(v6, v6)
                                                            }
                                                            v4 += v4
                                                            v127 := GenS9{F0: 42, F1: (!true)}
                                                            switch (-(fi118)) {
                                                            case -979:
                                                                switch 127 {
                                                                case -590:
                                                                    v103 = (!v103)
                                                                default:
                                                                    fmt.Println(v103, v103)
                                                                }
                                                            case 4294967295:
                                                                switch "" {
                                                                case "a":
                                                                    var v128 uint32 = 65535
                                                                    _ = v128
                                                                case "foo bar":
                                                                    if v6 {
                                                                        v129 := GenS5{F0: "", F1: (fi118 * 32767), F2: 0}
                                                                        _ = v129
                                                                    }
                                                                }
                                                            case 32768:
                                                                var v130 string = "\n\t"
                                                                _ = v130
                                                            default:
                                                                v131 := GenS10{F0: ("hello" + (("hello" + (("a" + "\n\t") + "foo bar")) + ("foo bar" + "foo bar"))), F1: 32767, F2: 503}
                                                                _ = v131
                                                            }
                                                            _ = v127
                                                        }
                                                    } else {
                                                        v132 := GenS7{F0: v104, F1: 42, F2: true}
                                                        _ = len(v5)
                                                        v133 := GenS11{F0: ("a" + (("hello" + "") + "foo bar"))}
                                                        v6 = true
                                                        v134 := GenS4{F0: 65535, F1: 1, F2: v104}
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    }
                                                    v104 = v104
                                                case -706:
                                                    var v135 int16 = 1000
                                                    _ = v135
                                                case 4294967295:
                                                    if (!v103) {
                                                        v6 = v103
                                                        if v103 {
                                                            if (255 < 239) {
                                                                if ((!true) || (565 > 42)) {
                                                                    var v136 string = ("\n\t" + (("\n\t" + "foo bar") + ("foo bar" + "\n\t")))
                                                                    _ = v136
                                                                } else {
                                                                    var v137 uint = v104
                                                                    _ = v137
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = 2147483647
                                                            }
                                                        }
                                                    } else {
                                                        v5[(fi118 / ((fi118) | 1))] = "\n\t"
                                                        v104 = 1
                                                    }
                                                default:
                                                    if (v6 && (v103 && true)) {
                                                        v138 := GenS9{F0: 42, F1: (false || (!(2147483647 >= 2147483648)))}
                                                        fmt.Println(v4, v103)
                                                        _ = v138
                                                    } else {
                                                        if (!v6) {
                                                            if (!(!(v104 == 4294967295))) {
                                                                if v103 {
                                                                    fmt.Println(fi117)
                                                                } else {
                                                                    v103 = (false && (!v6))
                                                                    var v139 uint8 = 42
                                                                    var v140 float64 = (-2.25 * v4)
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                            } else {
                                                                var v141 uint = (^(v104))
                                                                v142 := GenS12{F0: 1, F1: 4294967295}
                                                                _ = len(v3)
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                            var v143 float64 = float64(0.1)
                                                            _ = v143
                                                        }
                                                        var v144 string = ("foo bar" + ("a" + "a"))
                                                        if (v6 && false) {
                                                            if (!v6) {
                                                                delete(v5, (^(fi117)))
                                                            }
                                                        }
                                                        if v6 {
                                                            v145 := GenS6{F0: v144, F1: false, F2: int64(fi117)}
                                                            v144 = v144
                                                            switch ((((v144 + (((("" + "hello") + "") + (v144 + "")) + (v144 + v144))) + v144) + "\n\t") + (v144 + "")) {
                                                            case "hello":
                                                                var v146 bool = v6
                                                                _ = v146
                                                            case "a":
                                                                var v147 int32 = -1
                                                                _ = v147
                                                            case "foo bar":
                                                                v144 = ""
                                                            }
                                                            _ = v145
                                                        }
                                                        var v148 uint8 = 141
                                                        v149 := GenS5{F0: v144, F1: (fi118 % ((fi118) | 1)), F2: -2147483648}
                                                        _ = v144
                                                        _ = v148
                                                        _ = v149
                                                    }
                                                }
                                            }
                                            if (!true) {
                                                v104 = (v104 | (v104 << (((^(v104))) & 63)))
                                                v4 = (v4 / (1.5 - (float64(0.0) + 3.14159)))
                                                if v103 {
                                                    if (v6 && true) {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if (-83 > 52) {
                                                                var v150 string = ("\n\t" + "hello")
                                                                _ = v150
                                                            } else {
                                                                switch (fi118 >> ((v104) & 63)) {
                                                                case 127:
                                                                    v4 += (0.1 * 0.1)
                                                                }
                                                            }
                                                            v4 = v4
                                                            v151 := GenS13{F0: 0, F1: 1}
                                                            _ = v151
                                                        case "":
                                                            var v152 int16 = 32767
                                                            _ = v152
                                                        case "foo bar":
                                                            v103 = (-1 >= 128)
                                                        default:
                                                            if v6 {
                                                                if true {
                                                                    var v153 float32 = -2.25
                                                                    fmt.Println(v104)
                                                                    v103 = true
                                                                    _ = v153
                                                                } else {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v154 int8 = -1
                                                                        _ = v154
                                                                    default:
                                                                        v155 := GenS7{F0: v104, F1: -2147483648, F2: (255 < 1)}
                                                                        _ = v155
                                                                    }
                                                                    _ = len(v2)
                                                                    v104 += 84
                                                                }
                                                            } else {
                                                                v103 = (226 < 130)
                                                                v103 = ("" <= "\n\t")
                                                            }
                                                        }
                                                        delete(v5, 128)
                                                        _ = len(v3)
                                                        var v156 string = ("\n\t" + "")
                                                        _ = v156
                                                    }
                                                } else {
                                                    fmt.Println(v6, v104, v6)
                                                    var v157 float32 = (1e10 - -2.25)
                                                    if (v103 || false) {
                                                        var v158 string = "\n\t"
                                                        v4 = 3.14159
                                                        var v159 uint64 = 2
                                                        _ = v158
                                                        _ = v159
                                                    }
                                                    var v160 uint64 = 4294967295
                                                    var v161 uint64 = (v160 << ((v104) & 63))
                                                    _ = v157
                                                    _ = v160
                                                    _ = v161
                                                }
                                            } else {
                                                var v162 uint16 = 65535
                                                v6 = (9223372036854775807 >= -591)
                                                var v163 int64 = 793
                                                _ = v162
                                                _ = v163
                                            }
                                            var v164 float64 = v4
                                            _ = fi118
                                            _ = v164
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                    }
                                    for fi165 := 0; fi165 < 2; fi165++ {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v103 = v6
                                        _ = fi165
                                    }
                                    var v166 uint16 = 65535
                                    _ = v166
                                }
                                switch (("a" + ("\n\t" + "a")) + (("a" + ("\n\t" + "a")) + "")) {
                                case "foo bar":
                                    switch ("a" + (("" + "hello") + "a")) {
                                    case "\n\t":
                                        for fi167 := 0; fi167 < 12; fi167++ {
                                            var v168 int8 = 99
                                            if v103 {
                                                var v169 int8 = 42
                                                if true {
                                                    var v170 uint64 = 18446744073709551615
                                                    _ = v170
                                                } else {
                                                    fmt.Println(v104)
                                                    var v171 int64 = 42
                                                    var v172 float32 = (float32(v4) / 3.14159)
                                                    _ = v171
                                                    _ = v172
                                                }
                                                var v173 uint16 = 566
                                                _ = v169
                                                _ = v173
                                            } else {
                                                if true {
                                                    if v103 {
                                                        switch ((("hello" + "\n\t") + "") + "a") {
                                                        case "\n\t":
                                                            v174 := v5[fi117]
                                                            var v175 int = 404
                                                            _ = v174
                                                            _ = v175
                                                        case "a":
                                                            v176 := GenS14{F0: 2147483647, F1: 2}
                                                            _ = v176
                                                        }
                                                        if ((2147483647 != 445) || v103) {
                                                            v103 = v6
                                                            fmt.Println(fi117, v104, v168)
                                                        } else {
                                                            var v177 uint8 = 8
                                                            _ = v177
                                                        }
                                                        v103 = ((false && false) && (0 > 121))
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi117 * 1)
                                                        }
                                                        v103 = (32767 < 65535)
                                                    }
                                                    var v178 int32 = -2147483648
                                                    v104 = v104
                                                    _ = v178
                                                } else {
                                                    fmt.Println(v168, v103)
                                                    fmt.Println(v4)
                                                }
                                            }
                                            _ = fi167
                                            _ = v168
                                        }
                                    default:
                                        var v179 int32 = 164
                                        _ = v179
                                    }
                                case "hello":
                                    switch (("foo bar" + ("a" + ((("" + (("\n\t" + "hello") + "\n\t")) + "a") + ("\n\t" + ("\n\t" + "a"))))) + ("foo bar" + "")) {
                                    case "":
                                        var v180 int64 = 1
                                        _ = v180
                                    }
                                default:
                                    switch (fi117 >> ((v104) & 63)) {
                                    case 0:
                                        if false {
                                            var v181 uint32 = 1
                                            _ = v181
                                        }
                                    case 1:
                                        if (!(false && (1 != 0))) {
                                            if false {
                                                var v182 int16 = 332
                                                fmt.Println(v6, v182, fi117)
                                                _ = v182
                                            }
                                            if (302 >= 2) {
                                                if v103 {
                                                    var v183 int8 = 19
                                                    _ = v183
                                                } else {
                                                    for fi184 := 0; fi184 < 18; fi184++ {
                                                        v6 = true
                                                        if v103 {
                                                            v6 = false
                                                            fmt.Println(v6, v104)
                                                            switch (("a" + ("foo bar" + "foo bar")) + "foo bar") {
                                                            case "a":
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi184
                                                                }
                                                            case "\n\t":
                                                                v185 := GenS15{F0: 1}
                                                                _ = v185
                                                            default:
                                                                var v186 int = (-(fi117))
                                                                _ = v186
                                                            }
                                                            var v187 float64 = 0.1
                                                            _ = v187
                                                        }
                                                        var v188 int16 = 1000
                                                        var v189 uint32 = 2
                                                        delete(v5, (fi184 * fi117))
                                                        _ = fi184
                                                        _ = v188
                                                        _ = v189
                                                    }
                                                    switch 32768 {
                                                    case 128:
                                                        var v190 uint8 = 95
                                                        _ = v190
                                                    case -2147483648:
                                                        v4 = v4
                                                    }
                                                }
                                                v191 := GenS5{F0: (("hello" + "\n\t") + "hello"), F1: fi117, F2: 1}
                                                _ = v191
                                            }
                                            for fi192 := 0; fi192 < 5; fi192++ {
                                                v6 = ((3.14159 == (v4 - 0.0)) || (-9223372036854775808 <= 1))
                                                var v193 uint32 = 2
                                                _ = fi192
                                                _ = v193
                                            }
                                        }
                                    default:
                                        fmt.Println(v6)
                                    }
                                }
                                v194 := GenS5{F0: "foo bar", F1: 32768, F2: -159}
                                _ = v194
                            } else {
                                switch 32767 {
                                case 2147483647:
                                    _ = len(v5)
                                case -1:
                                    var v195 bool = false
                                    _ = v195
                                }
                                switch ("foo bar" + "a") {
                                case "hello":
                                    var v196 int16 = 1000
                                    _ = v196
                                case "a":
                                    _ = len(v5)
                                case "":
                                    switch "foo bar" {
                                    case "a":
                                        v197 := v5[(fi117 * (fi117 - fi117))]
                                        _ = v197
                                    case "foo bar":
                                        var v198 int16 = 1000
                                        _ = v198
                                    case "hello":
                                        v6 = ((0.1 + (-2.25 + ((1.5 / float32((1.5 * v4))) * -2.25))) != float32(1.5))
                                    default:
                                        fmt.Println(v103)
                                    }
                                default:
                                    var v199 int = (fi117 >> ((4294967295) & 63))
                                    _ = v199
                                }
                            }
                        }
                        v200 := GenS11{F0: ("" + "foo bar")}
                        _ = fi117
                        _ = v200
                    }
                    var v201 uint = (v104 * (v104 - (v104 >> ((2) & 63))))
                    for fi202 := 0; fi202 < 20; fi202++ {
                        fmt.Println(v6, v103)
                        var v203 float32 = (-2.25 / 3.14159)
                        _ = fi202
                        _ = v203
                    }
                    _ = v201
                }
                switch -95 {
                case -1:
                    if (!(-864 != 0)) {
                        var v204 float32 = float32(v4)
                        var v205 int16 = 1
                        fmt.Println(v103, v103)
                        _ = v204
                        _ = v205
                    } else {
                        var v206 int = 9223372036854775807
                        _ = v206
                    }
                default:
                    switch 789 {
                    case 469:
                        v103 = true
                    case 32767:
                        var v207 int32 = 0
                        _ = v207
                    case -9223372036854775808:
                        for fi208 := 0; fi208 < 6; fi208++ {
                            v103 = false
                            v209 := GenS10{F0: "hello", F1: 42, F2: 32767}
                            _ = fi208
                            _ = v209
                        }
                    default:
                        switch 455 {
                        case 32768:
                            switch "" {
                            case "":
                                v210 := GenS5{F0: ("\n\t" + "hello"), F1: 128, F2: 1}
                                _ = v210
                            }
                        case 32767:
                            var v211 int32 = 2147483647
                            _ = v211
                        }
                    }
                }
                fmt.Println(v4)
            }
        }
        _ = v104
    } else {
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        var v212 float64 = (1.5 + ((float64((0.1 * 1.5)) / v4) - v4))
        _ = v212
    }
    if ("foo bar" != ("a" + ("" + "foo bar"))) {
        v6 = (540 != 809)
        var v213 int64 = 397
        if (0 <= 460) {
            v214 := GenS2{F0: 42, F1: 42}
            if true {
                switch -1 {
                case -2147483648:
                    switch -160 {
                    case 0:
                        v213 -= (v213 / ((v213) | 1))
                    }
                case 0:
                    v213 -= v213
                default:
                    fmt.Println(v4, v213, v4)
                }
                var v215 int32 = 410
                for fi216 := 0; fi216 < 7; fi216++ {
                    v217 := GenS16{F0: 65535, F1: 398}
                    fmt.Println(v4, v103)
                    _ = fi216
                    _ = v217
                }
                _ = v215
            } else {
                switch -591 {
                case -103:
                    v213 = -73
                case 4294967295:
                    if v6 {
                        for fi218 := 0; fi218 < 18; fi218++ {
                            var v219 int = 127
                            _ = fi218
                            _ = v219
                        }
                        var v220 uint = 781
                        _ = v220
                    } else {
                        for fi221 := 0; fi221 < 17; fi221++ {
                            if len(v3) > 0 {
                                v3[0] = (fi221 / (((fi221 << ((18446744073709551615) & 63))) | 1))
                            }
                            _ = fi221
                        }
                        var v222 int16 = 0
                        if v6 {
                            v222 = (v222 - v222)
                            v4 = (v4 - v4)
                            v223 := GenS17{F0: ((("foo bar" + "\n\t") + "hello") + ("" + ("\n\t" + ("foo bar" + ("foo bar" + ("" + "a")))))), F1: 1, F2: 853}
                            _ = v223
                        }
                        _ = v222
                    }
                case 127:
                    var v224 float32 = float32((1e10 + (v4 / v4)))
                    _ = v224
                }
                var v225 uint = 603
                for fi226 := 0; fi226 < 12; fi226++ {
                    if true {
                        for fi227 := 0; fi227 < 19; fi227++ {
                            if (65535 != 2147483648) {
                                v4 += float64(3.14159)
                                fmt.Println(v225, v225)
                                fmt.Println(v213)
                            } else {
                                var v228 float64 = v4
                                v103 = v103
                                v4 -= (1e10 - (((v228 + ((1.5 + 0.1) + 1.5)) / v4) / (float64((3.14159 * (3.14159 * 3.14159))) / v228)))
                                _ = v228
                            }
                            delete(v5, 127)
                            _ = fi227
                        }
                        switch "a" {
                        case "a":
                            fmt.Println(v225, v4)
                        case "":
                            var v229 float32 = 1e10
                            _ = v229
                        default:
                            for fi230 := 0; fi230 < 20; fi230++ {
                                v225 -= v225
                                _ = fi230
                            }
                        }
                    }
                    switch fi226 {
                    case 1:
                        v4 = (v4 / float64(-2.25))
                    case 4294967295:
                        if (("foo bar" + "a") < "a") {
                            fmt.Println(v6, v103, v225)
                            if v6 {
                                switch "\n\t" {
                                case "\n\t":
                                    var v231 uint8 = 0
                                    _ = v231
                                default:
                                    v232 := GenS18{F0: 2147483647}
                                    _ = v232
                                }
                                if ((("foo bar" + "\n\t") + "\n\t") < "a") {
                                    var v233 uint16 = 21
                                    if v6 {
                                        var v234 uint16 = v233
                                        _ = v234
                                    }
                                    var v235 uint8 = 42
                                    _ = v233
                                    _ = v235
                                }
                            } else {
                                var v236 bool = ((-1 >= -716) || (!(1.5 >= (3.14159 + 0.1))))
                                var v237 int8 = 1
                                _ = v236
                                _ = v237
                            }
                        } else {
                            v6 = (!true)
                        }
                    }
                    switch (((((("\n\t" + "hello") + "foo bar") + "foo bar") + (("foo bar" + "\n\t") + "")) + "hello") + (("" + "hello") + ("" + ("" + "")))) {
                    case "hello":
                        _ = len(v3)
                    case "\n\t":
                        v213 += 1
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi226
                }
                _ = v225
            }
            var v238 string = "foo bar"
            _ = v214
            _ = v238
        } else {
            for fi239 := 0; fi239 < 9; fi239++ {
                for fi240 := 0; fi240 < 11; fi240++ {
                    switch (fi239 + fi240) {
                    case 0:
                        var v241 bool = false
                        _ = v241
                    case 4294967295:
                        var v242 int64 = (v213 % ((v213) | 1))
                        _ = v242
                    default:
                        v103 = (3.14159 >= (((float32((0.0 - (((v4 - v4) * (v4 + v4)) * float64((1.5 - float32(v4)))))) - 0.1) * (0.1 + (-2.25 / (0.1 + (float32((float64((0.0 * -2.25)) * (float64(0.1) * v4))) / 3.14159))))) * 3.14159))
                    }
                    _ = fi240
                }
                v103 = v6
                _ = fi239
            }
            v243 := GenS4{F0: 255, F1: 111, F2: 893}
            _ = v243
        }
        _ = v213
    }
    v244 := GenS13{F0: 2, F1: 4294967295}
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v2
    _ = v244
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
}
