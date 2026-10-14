package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 uint64; F1 int8; F2 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int8; F1 string; F2 int8 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint16; F1 string }
type GenS18 struct { F0 string; F1 int32; F2 uint16 }
func genFn0(p0 uint8, p1 uint64, p2 int8) int {
    p1 = 42
    var v6 uint32 = 7
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 int, p1 int, p2 string) bool {
    p0 += p1
    v7 := []int{(p1 & p1)}
    if (true && (-780 < -295)) {
        p1 = p0
        v8 := GenS1{F0: 744, F1: 0}
        p0 = (p0 >> ((551) & 63))
        _ = v8
    } else {
        var v9 bool = true
        switch p0 {
        case -2147483648:
            v10 := GenS2{F0: -1, F1: 42}
            _ = v10
        case 4294967295:
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int64 = 702
                _ = fi11
                _ = v12
            }
        default:
            v9 = (0.1 == 0.1)
        }
        var v13 uint64 = 606
        _ = v13
        _ = v9
    }
    var v14 float32 = ((0.1 / 1.5) + (1.5 + float32((3.14159 + (3.14159 - (-2.25 - 3.14159))))))
    var v15 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v7
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 65535
    var v3 int16 = 42
    var v4 uint8 = 1
    v5 := make(map[int]uint32)
    v16 := []int{32768}
    v3 = (v3 / ((v3) | 1))
    if ((v3 >> ((uint(v4)) & 63)) > -904) {
        if false {
            if false {
                if (!(false || false)) {
                    v1[(v2 << ((255) & 63))] = false
                }
                var v17 int16 = 0
                fmt.Println(v2, v3)
                _ = v17
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                v3 = v3
                _ = fi18
            }
            switch (("a" + ("foo bar" + "")) + ("hello" + "foo bar")) {
            case "a":
                delete(v5, 249)
            }
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            switch genFn0(128, 2, int8(v2)) {
            case 2147483647:
                switch (fi19 + (fi19 % ((32768) | 1))) {
                case 128:
                    var v20 uint16 = 65535
                    _ = v20
                case 591:
                    v21 := GenS3{F0: ""}
                    _ = v21
                case 127:
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 uint = 42
                        var v24 uint8 = (^(v4))
                        var v25 string = ("\n\t" + "a")
                        switch (-(fi22)) {
                        case 9223372036854775807:
                            if ((!(false || true)) || (!true)) {
                                if false {
                                    _ = len(v16)
                                } else {
                                    v26 := GenS3{F0: v25}
                                    if (true || true) {
                                        if false {
                                            if ((false || (!true)) && (!(0 >= 1))) {
                                                if true {
                                                    switch 0 {
                                                    case 2147483647:
                                                        fmt.Println(fi22)
                                                    case 4294967295:
                                                        v2 = 1
                                                    case 238:
                                                        v27 := GenS4{F0: (v23 >> ((v23) & 63)), F1: 0}
                                                        _ = v27
                                                    }
                                                } else {
                                                    v28 := GenS3{F0: ("\n\t" + v25)}
                                                    var v29 string = "hello"
                                                    v4 -= v4
                                                    _ = v28
                                                    _ = v29
                                                }
                                                v2 = (v2 >> ((v23) & 63))
                                            } else {
                                                var v30 float64 = -2.25
                                                _ = v30
                                            }
                                            var v31 int8 = 126
                                            _ = v31
                                        } else {
                                            var v32 uint8 = 78
                                            var v33 float32 = ((float32(1.5) / 3.14159) * ((1.5 / (-2.25 / 1.0)) / ((1e10 - 3.14159) / (0.0 + (0.0 + 3.14159)))))
                                            _ = v32
                                            _ = v33
                                        }
                                        v4 -= (v24 | 1)
                                        v16 = append(v16, 32768)
                                    }
                                    v24 += 233
                                    _ = v26
                                }
                                var v34 int8 = -128
                                v35 := GenS5{F0: false, F1: false}
                                _ = v34
                                _ = v35
                            } else {
                                var v36 int64 = -1
                                _ = v36
                            }
                        default:
                            var v37 int32 = 0
                            _ = v37
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                default:
                    var v38 float32 = ((0.1 * ((1.5 - 0.0) * 1.5)) * (0.1 - 0.0))
                    _ = v38
                }
            case -1:
                fmt.Println(fi19)
            case -2147483648:
                if (3.14159 < (float64((3.14159 / float32((((3.14159 - (1.5 - 0.0)) + 0.0) + 1.5)))) * 0.0)) {
                    v39 := GenS5{F0: false, F1: (-2.25 <= (0.0 / 0.1))}
                    var v40 uint64 = uint64(v4)
                    for fi41 := 0; fi41 < 9; fi41++ {
                        var v42 uint32 = 355
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        v4 = uint8(v2)
                        var v43 int = fi41
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    _ = v39
                    _ = v40
                } else {
                    if true {
                        var v44 uint = 65535
                        _ = v44
                    } else {
                        if false {
                            fmt.Println(v4)
                            fmt.Println(v4, v2)
                        }
                        for fi45 := 0; fi45 < 13; fi45++ {
                            if (!(0 < 2)) {
                                if (!(true && (42 != 18446744073709551615))) {
                                    v16 = append(v16, fi45)
                                    var v46 string = "hello"
                                    _ = v46
                                } else {
                                    fmt.Println(v2, v3, v3)
                                }
                                v47 := GenS1{F0: 42, F1: -1}
                                _ = v47
                            } else {
                                if false {
                                    switch "foo bar" {
                                    case "a":
                                        v48 := GenS2{F0: 1, F1: v4}
                                        _ = v48
                                    case "":
                                        v2 = v2
                                    case "foo bar":
                                        var v49 float32 = (3.14159 * (3.14159 * ((-2.25 + (-2.25 + (((((3.14159 / (float32((((((1e10 + (-2.25 - ((float64(0.0) * (1.5 / 1.0)) * 1e10))) * 0.1) / 1.0) * 0.0) + 1.5)) + 1e10)) + 0.1) / 3.14159) + ((((0.1 - 1e10) + 1.5) - (1e10 + 3.14159)) + 1.5)) * 0.1))) / 1e10)))
                                        _ = v49
                                    }
                                    fmt.Println(fi45, v4, fi45)
                                    fmt.Println(fi19, v4, v2)
                                } else {
                                    if (0 == 4294967295) {
                                        _ = len(v5)
                                        var v50 int8 = 42
                                        _ = v50
                                    } else {
                                        if (42 != 1) {
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            switch fi19 {
                                            case 32767:
                                                v3 = (^(v3))
                                            case 2147483647:
                                                v51 := GenS6{F0: 4294967295, F1: 100, F2: 1}
                                                _ = v51
                                            case -3:
                                                _ = len(v16)
                                            default:
                                                v52 := GenS7{F0: 1}
                                                var v53 int16 = 645
                                                _ = v52
                                                _ = v53
                                            }
                                            fmt.Println(fi45, fi19)
                                        }
                                        if false {
                                            if (!true) {
                                                v3 = 0
                                                var v54 uint32 = (v2 * v2)
                                                var v55 uint16 = 130
                                                var v56 float64 = (float64(1e10) * 0.0)
                                                var v57 uint32 = v2
                                                v16 = append(v16, (fi19 / (((fi45 | 127)) | 1)))
                                                if (-949 == 2147483647) {
                                                    if len(v16) >= 2 {
                                                        v16 = v16[1:len(v16)]
                                                    }
                                                }
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            } else {
                                                v3 = v3
                                                v58 := GenS2{F0: 100, F1: v4}
                                                _ = v58
                                            }
                                            v59 := GenS5{F0: (0 < 18446744073709551615), F1: (!((v4 ^ (v4 >> ((1) & 63))) > v4))}
                                            v3 += (-(v3))
                                            _ = v59
                                        }
                                        v60 := GenS6{F0: 42, F1: 0, F2: 609}
                                        _ = v60
                                    }
                                    v4 += (v4 / ((v4) | 1))
                                    switch fi45 {
                                    case 375:
                                        _ = len(v5)
                                    default:
                                        v2 += (v2 % ((v2) | 1))
                                    }
                                    v16 = append(v16, fi19)
                                }
                                if true {
                                    switch (fi19 * fi45) {
                                    case 938:
                                        v61 := GenS8{F0: 128}
                                        _ = v61
                                    case -9223372036854775808:
                                        if (v2 < v2) {
                                            fmt.Println(v2, fi19)
                                        } else {
                                            if false {
                                                v62 := GenS3{F0: "\n\t"}
                                                _ = v62
                                            } else {
                                                v3 -= (v3 ^ 3)
                                                var v63 uint8 = 255
                                                v2 = v2
                                                _ = v63
                                            }
                                            if true {
                                                fmt.Println(v4, fi45)
                                                v64 := GenS8{F0: 12}
                                                var v65 uint64 = 18446744073709551615
                                                _ = v64
                                                _ = v65
                                            } else {
                                                var v66 int32 = 2147483647
                                                v67 := GenS6{F0: 1, F1: 42, F2: 42}
                                                _ = v66
                                                _ = v67
                                            }
                                        }
                                    case 42:
                                        v4 = 42
                                    default:
                                        if (!((^(v2)) >= (v2 % (((^(v2))) | 1)))) {
                                            var v68 int = fi19
                                            _ = v68
                                        }
                                    }
                                    switch ("foo bar" + "foo bar") {
                                    case "\n\t":
                                        _ = len(v16)
                                    }
                                } else {
                                    v2 -= (^(v2))
                                    fmt.Println(v4, fi19, fi45)
                                }
                                switch "\n\t" {
                                case "a":
                                    var v69 uint64 = 1
                                    _ = v69
                                case "":
                                    if (42 <= 292) {
                                        v16 = append(v16, (-(fi45)))
                                    } else {
                                        switch "a" {
                                        case "":
                                            v16 = append(v16, -9223372036854775808)
                                        case "hello":
                                            v2 = v2
                                        }
                                    }
                                }
                            }
                            v1[0] = ("hello" != (("foo bar" + "a") + "hello"))
                            _ = fi45
                        }
                        if false {
                            var v70 int = (fi19 >> ((42) & 63))
                            var v71 int16 = v3
                            fmt.Println(v71, v4, v4)
                            _ = v70
                            _ = v71
                        } else {
                            v4 += v4
                            if (v2 == v2) {
                                v16 = append(v16, fi19)
                                v4 = 42
                            } else {
                                var v72 float32 = 0.1
                                _ = v72
                            }
                        }
                    }
                }
            }
            switch ("foo bar" + ("hello" + ("a" + "hello"))) {
            case "":
                var v73 int = fi19
                _ = v73
            default:
                fmt.Println(v3, v3)
            }
            _ = fi19
        }
        v4 -= (^(v4))
    } else {
        v2 += 378
        for fi74 := 0; fi74 < 14; fi74++ {
            v2 = (v2 - 0)
            _ = len(v16)
            var v75 uint8 = (v4 | v4)
            _ = fi74
            _ = v75
        }
        for fi76 := 0; fi76 < 9; fi76++ {
            v3 -= -32768
            _ = fi76
        }
    }
    for fi77 := 0; fi77 < 20; fi77++ {
        var v78 uint = 4294967295
        _ = fi77
        _ = v78
    }
    var v79 int16 = (v3 * -32768)
    if (true && (0 < 963)) {
        var v80 uint = 367
        for fi81 := 0; fi81 < 15; fi81++ {
            v79 = v3
            var v82 bool = (!true)
            for fi83 := 0; fi83 < 14; fi83++ {
                delete(v5, 32767)
                v82 = false
                _ = fi83
            }
            _ = fi81
            _ = v82
        }
        v5[128] = (v2 / ((v2) | 1))
        _ = v80
    } else {
        if true {
            if (("hello" + (("\n\t" + "a") + ("foo bar" + "foo bar"))) >= "\n\t") {
                v2 = (^(v2))
                for fi84 := 0; fi84 < 15; fi84++ {
                    _ = len(v16)
                    switch ("\n\t" + "a") {
                    case "":
                        for fi85 := 0; fi85 < 18; fi85++ {
                            if true {
                                if false {
                                    v4 = (v4 / (((v4 ^ (v4 >> ((712) & 63)))) | 1))
                                    if false {
                                        switch fi85 {
                                        case 445:
                                            fmt.Println(fi84, v2, v4)
                                        case -114:
                                            var v86 bool = (!false)
                                            _ = v86
                                        default:
                                            var v87 int = fi84
                                            _ = v87
                                        }
                                        v5[(fi84 * (fi84 + genFn0(v4, 42, 18)))] = (v2 >> ((42) & 63))
                                    }
                                }
                                v88 := GenS9{F0: 2}
                                _ = v88
                            } else {
                                v89 := GenS9{F0: 348}
                                fmt.Println(v4)
                                v79 = (v79 / ((v79) | 1))
                                _ = v89
                            }
                            _ = fi85
                        }
                    default:
                        for fi90 := 0; fi90 < 14; fi90++ {
                            if (0.0 < 1.5) {
                                v2 += v2
                                v79 = v79
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            }
                            var v91 float64 = -2.25
                            v79 -= v79
                            v91 += (v91 * (v91 + ((float64(0.0) + v91) + v91)))
                            _ = fi90
                            _ = v91
                        }
                    }
                    v79 += -50
                    _ = fi84
                }
                v3 = (v79 ^ 0)
            } else {
                v3 -= (v3 * -1)
                if (!true) {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    v4 -= 255
                    switch "foo bar" {
                    case "hello":
                        for fi92 := 0; fi92 < 7; fi92++ {
                            fmt.Println(fi92)
                            for fi93 := 0; fi93 < 8; fi93++ {
                                var v94 int8 = 1
                                v2 += v2
                                v3 = (v3 - int16(fi93))
                                _ = fi93
                                _ = v94
                            }
                            _ = fi92
                        }
                    case "foo bar":
                        if false {
                            fmt.Println(v4)
                            var v95 float32 = 0.0
                            _ = v95
                        }
                    }
                } else {
                    var v96 int16 = v79
                    _ = v96
                }
                v97 := GenS10{F0: 154}
                _ = v97
            }
            switch ("\n\t" + (("\n\t" + "foo bar") + ("hello" + "foo bar"))) {
            case "a":
                if ((true || false) && false) {
                    for fi98 := 0; fi98 < 13; fi98++ {
                        switch "" {
                        case "\n\t":
                            var v99 bool = (!(-1 > genFn0(v4, 18446744073709551615, 0)))
                            _ = v99
                        case "a":
                            if true {
                                var v100 float64 = ((0.1 * 1.5) / 1.0)
                                if len(v16) > 0 {
                                    v16[0] = fi98
                                }
                                _ = v100
                            }
                        default:
                            var v101 bool = (v4 == 2)
                            _ = v101
                        }
                        if ((false || (1 == 1)) && false) {
                            for fi102 := 0; fi102 < 9; fi102++ {
                                if (("" + "foo bar") >= ((((("a" + ("" + ("hello" + "hello"))) + ("\n\t" + "")) + "a") + ((("hello" + ("\n\t" + ("foo bar" + ("foo bar" + ("foo bar" + "hello"))))) + "hello") + ("foo bar" + "foo bar"))) + (("a" + "hello") + ("" + "foo bar")))) {
                                    switch ("hello" + ("" + "foo bar")) {
                                    case "foo bar":
                                        fmt.Println(v2, v2, fi102)
                                    }
                                } else {
                                    v103 := GenS11{F0: 42, F1: "foo bar"}
                                    _ = v103
                                }
                                var v104 float32 = (0.1 - 1.5)
                                v5[32768] = 65535
                                _ = fi102
                                _ = v104
                            }
                            fmt.Println(v79, v2)
                            for fi105 := 0; fi105 < 19; fi105++ {
                                v4 += 164
                                _ = fi105
                            }
                        } else {
                            for fi106 := 0; fi106 < 19; fi106++ {
                                fmt.Println(v2, v4, v4)
                                var v107 float32 = 3.14159
                                _ = fi106
                                _ = v107
                            }
                        }
                        var v108 uint32 = v2
                        _ = fi98
                        _ = v108
                    }
                    for fi109 := 0; fi109 < 19; fi109++ {
                        v4 -= v4
                        _ = fi109
                    }
                    v110 := GenS3{F0: "foo bar"}
                    _ = v110
                } else {
                    switch ((("hello" + ("foo bar" + ("\n\t" + "hello"))) + ((("\n\t" + ((("\n\t" + "a") + "") + "")) + "\n\t") + "a")) + "foo bar") {
                    case "":
                        v2 -= (^(v2))
                    case "\n\t":
                        switch 2147483647 {
                        case 128:
                            v111 := GenS4{F0: 893, F1: uint64(v2)}
                            _ = v111
                        case -9223372036854775808:
                            for fi112 := 0; fi112 < 16; fi112++ {
                                for fi113 := 0; fi113 < 18; fi113++ {
                                    fmt.Println(fi112, fi113)
                                    _ = fi113
                                }
                                _ = fi112
                            }
                        case -1000:
                            fmt.Println(v79, v79)
                        }
                    case "hello":
                        for fi114 := 0; fi114 < 14; fi114++ {
                            v4 -= 128
                            _ = fi114
                        }
                    default:
                        var v115 string = "a"
                        _ = v115
                    }
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        var v116 string = "\n\t"
                        _ = v116
                    case -1:
                        v2 += v2
                    case 1:
                        v117 := GenS12{F0: 640}
                        _ = v117
                    default:
                        for fi118 := 0; fi118 < 3; fi118++ {
                            var v119 bool = false
                            v2 -= 2147483648
                            _ = fi118
                            _ = v119
                        }
                    }
                }
            default:
                var v120 int64 = -9223372036854775808
                _ = v120
            }
        }
        fmt.Println(v79, v4, v4)
        var v121 int8 = 100
        _ = v121
    }
    var v122 uint8 = (^(v4))
    if (0 >= 0) {
        var v123 float32 = ((1.5 + 1e10) / 0.1)
        if (!false) {
            var v124 int16 = (-(v3))
            if ((v4 + v122) <= (v4 - 255)) {
                var v125 int32 = 0
                if (v2 == v2) {
                    v126 := GenS13{F0: 42}
                    _ = v126
                }
                var v127 int16 = 1000
                _ = v125
                _ = v127
            }
            _ = v124
        } else {
            if true {
                v3 = (v3 * (^(v79)))
            }
        }
        v128 := GenS13{F0: 0}
        _ = v123
        _ = v128
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        var v129 int32 = 42
        _ = v129
    case "hello":
        v130 := GenS13{F0: 42}
        _ = v130
    default:
        fmt.Println(v122, v3)
    }
    v79 = v3
    if (!(-2.25 < (float64(1e10) + -2.25))) {
        for fi131 := 0; fi131 < 20; fi131++ {
            v16 = append(v16, fi131)
            for fi132 := 0; fi132 < 11; fi132++ {
                if true {
                    if (((3.14159 - 1.5) > 3.14159) || false) {
                        var v133 float64 = (3.14159 / 1.5)
                        var v134 uint16 = 32767
                        _ = v133
                        _ = v134
                    } else {
                        var v135 int16 = v3
                        v3 -= (v79 * 0)
                        _ = v135
                    }
                    if (!true) {
                        v136 := GenS14{F0: v3}
                        _ = v136
                    } else {
                        if (42 == 255) {
                            var v137 int16 = (v3 << ((90) & 63))
                            fmt.Println(fi131, v79, v122)
                            fmt.Println(v137)
                            var v138 uint32 = (^(v2))
                            _ = v137
                            _ = v138
                        } else {
                            var v139 int16 = int16(fi131)
                            var v140 int8 = 42
                            if false {
                                v141 := GenS3{F0: "hello"}
                                if false {
                                    _ = len(v5)
                                    var v142 int16 = v3
                                    var v143 int64 = 1
                                    _ = v142
                                    _ = v143
                                } else {
                                    v144 := GenS2{F0: v140, F1: (v122 * (v122 * (v4 * 255)))}
                                    v2 -= 2147483648
                                    v3 -= v139
                                    _ = v144
                                }
                                _ = v141
                            }
                            if (42 >= -2147483648) {
                                var v145 uint32 = v2
                                _ = v145
                            } else {
                                v146 := GenS15{F0: v140, F1: "\n\t", F2: (^(v140))}
                                if (32767 != 65535) {
                                    var v147 uint64 = 18446744073709551615
                                    switch "" {
                                    case "hello":
                                        var v148 int8 = (v140 >> ((uint(v2)) & 63))
                                        _ = v148
                                    case "\n\t":
                                        v16 = append(v16, genFn0(v122, 42, (v140 >> ((2) & 63))))
                                    default:
                                        if false {
                                            var v149 bool = (255 > 79)
                                            _ = v149
                                        } else {
                                            v150 := GenS13{F0: 75}
                                            _ = len(v16)
                                            var v151 int = int(v140)
                                            _ = v150
                                            _ = v151
                                        }
                                    }
                                    var v152 int16 = v79
                                    v16 = append(v16, -9223372036854775808)
                                    _ = v147
                                    _ = v152
                                } else {
                                    var v153 uint8 = (v122 << ((65535) & 63))
                                    var v154 uint = 65535
                                    v79 -= (v79 ^ (v3 ^ 32767))
                                    v79 = 42
                                    var v155 int = 32768
                                    _ = v153
                                    _ = v154
                                    _ = v155
                                }
                                var v156 int32 = 2147483647
                                _ = v146
                                _ = v156
                            }
                            v5[(fi132 * fi131)] = 4294967295
                            _ = v139
                            _ = v140
                        }
                        v2 += uint32(v122)
                        fmt.Println(v2, v122)
                    }
                }
                _ = fi132
            }
            v157 := GenS9{F0: 255}
            _ = fi131
            _ = v157
        }
        v158 := GenS16{F0: ""}
        if (false && false) {
            for fi159 := 0; fi159 < 19; fi159++ {
                var v160 int8 = int8(fi159)
                _ = len(v16)
                var v161 int8 = int8(v3)
                _ = fi159
                _ = v160
                _ = v161
            }
            v79 -= v3
            for fi162 := 0; fi162 < 6; fi162++ {
                var v163 uint16 = 255
                for fi164 := 0; fi164 < 13; fi164++ {
                    v165 := GenS2{F0: -111, F1: 3}
                    fmt.Println(fi164)
                    if true {
                        v122 = v4
                        v166 := GenS3{F0: ("hello" + "")}
                        _ = len(v1)
                        switch "\n\t" {
                        case "\n\t":
                            v3 += 1
                        default:
                            fmt.Println(fi162)
                        }
                        _ = v166
                    }
                    _ = fi164
                    _ = v165
                }
                if ((v163 >= uint16(v2)) && true) {
                    if (1e10 < 1.5) {
                        _ = len(v1)
                        v1[(v2 - 434)] = ("" < ("hello" + ""))
                    } else {
                        switch fi162 {
                        case 0:
                            v167 := GenS5{F0: ((v2 << ((4294967295) & 63)) >= (v2 % ((uint32(v122)) | 1))), F1: (2147483647 == -1)}
                            _ = v167
                        }
                    }
                    fmt.Println(v2)
                    var v168 uint8 = (v122 / ((1) | 1))
                    _ = v168
                } else {
                    fmt.Println(v2)
                    v79 = v79
                }
                _ = fi162
                _ = v163
            }
        }
        _ = v158
    }
    for fi169 := 0; fi169 < 15; fi169++ {
        v2 += (v2 + v2)
        fmt.Println(v4)
        _ = fi169
    }
    v170 := GenS17{F0: 1, F1: "a"}
    v171 := GenS18{F0: "hello", F1: 590, F2: 481}
    if (false || true) {
        v79 += v79
        v2 = v2
    } else {
        for fi172 := 0; fi172 < 14; fi172++ {
            for fi173 := 0; fi173 < 9; fi173++ {
                v122 = 42
                var v174 bool = true
                delete(v5, (fi172 & (fi173 + genFn0(36, 42, 127))))
                _ = fi173
                _ = v174
            }
            for fi175 := 0; fi175 < 20; fi175++ {
                var v176 int32 = -506
                _ = fi175
                _ = v176
            }
            fmt.Println(fi172, v79, v79)
            _ = fi172
        }
        v177 := GenS4{F0: 26, F1: 0}
        var v178 uint = 65535
        _ = v177
        _ = v178
    }
    if len(v16) > 0 {
        v16[0] = 9223372036854775807
    }
    v179 := GenS16{F0: ("" + "")}
    if false {
        fmt.Println(v3, v79)
        fmt.Println(v4, v4)
    } else {
        v180 := GenS5{F0: ((v79 != (v3 / ((273) | 1))) && false), F1: true}
        _ = v180
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v122:", v122)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v79:", v79)
    _ = v1
    _ = v122
    _ = v16
    _ = v170
    _ = v171
    _ = v179
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
}
