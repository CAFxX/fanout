package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint64 }
type GenS10 struct { F0 int16; F1 int; F2 uint16 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 int; F1 uint8; F2 int32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 bool; F1 string }
type GenS17 struct { F0 bool; F1 uint32 }
type GenS18 struct { F0 int32; F1 int16 }
type GenS19 struct { F0 int32; F1 string; F2 int }
type GenS20 struct { F0 string; F1 int16 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int }
type GenS23 struct { F0 uint64; F1 int64 }
type GenS24 struct { F0 int64; F1 int8 }
type GenS25 struct { F0 int64; F1 int64; F2 int }
type GenS26 struct { F0 uint }
type GenS27 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 uint16, p1 int16, p2 uint8) int {
    p2 += (p2 - (p2 + 93))
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int32 = 42
        v7 := []int{42, (fi5 << ((42) & 63))}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint32 = 65535
    var v9 float64 = (0.1 - 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1() int {
    switch ("" + ("\n\t" + "")) {
    case "a":
        var v10 uint8 = 42
        _ = v10
    case "":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint64 = 179
            v13 := GenS1{F0: (fi11 ^ fi11), F1: 2}
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(fi14, fi11, fi11)
                v12 = v12
                _ = fi14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    switch ("" + "a") {
    case "":
        v15 := []int{9223372036854775807, genFn0(2, 42, 238), -763, 398}
        _ = v15
    case "foo bar":
        v16 := []int{9223372036854775807, 128, genFn0(255, 580, 60)}
        _ = v16
    case "\n\t":
        switch ("a" + "\n\t") {
        case "":
            var v17 string = "a"
            _ = v17
        case "hello":
            var v18 float64 = 1.5
            _ = v18
        default:
            if ((0.1 == -2.25) && true) {
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 uint32 = 727
                    var v21 string = "a"
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        }
    default:
        for fi22 := 0; fi22 < 11; fi22++ {
            if (!((!(2 >= 2147483648)) || (false || true))) {
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 uint32 = 0
                    _ = fi23
                    _ = v24
                }
                if true {
                    v25 := GenS1{F0: 334, F1: 891}
                    fmt.Println(fi22)
                    var v26 uint32 = 658
                    _ = v25
                    _ = v26
                } else {
                    var v27 uint = 459
                    if ((!(734 <= 9223372036854775807)) && (((0.0 / 1.0) + float32(0.0)) != 0.1)) {
                        fmt.Println(v27)
                    } else {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            v29 := []int{}
                            v30 := GenS1{F0: 210, F1: 4294967295}
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = v27
                }
                for fi31 := 0; fi31 < 5; fi31++ {
                    var v32 float32 = ((float32(0.1) + 0.0) * (1.5 - ((1e10 / (1e10 + ((float32((3.14159 * 0.1)) * (1.5 - ((0.1 - 1.5) / float32(1e10)))) / ((1.5 - 1e10) * (0.1 * (1e10 + 1e10)))))) / (0.0 + float32(3.14159)))))
                    _ = fi31
                    _ = v32
                }
            }
            v33 := GenS1{F0: genFn0(854, -32768, 203), F1: 723}
            _ = fi22
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 int = fi34
        _ = fi34
        _ = v35
    }
    if false {
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 int32 = -2147483648
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 bool = (42 > 1000)
                if (!v39) {
                    var v40 float32 = 1.5
                    fmt.Println(v40, fi38, v37)
                    var v41 int32 = -1
                    v42 := []int{fi36}
                    switch ("\n\t" + "a") {
                    case "a":
                        _ = len(v42)
                    case "":
                        v37 -= (^(v41))
                    default:
                        switch "foo bar" {
                        case "hello":
                            if len(v42) > 0 {
                                v42[0] = fi38
                            }
                        case "foo bar":
                            if (!(4294967295 < 4294967295)) {
                                if (v39 && false) {
                                    if (-32768 > -901) {
                                        var v43 int32 = v41
                                        _ = len(v42)
                                        v41 += (v37 / (((-(v43))) | 1))
                                        v42 = append(v42, (fi38 / ((964) | 1)))
                                        switch fi38 {
                                        case 42:
                                            var v44 uint = 42
                                            _ = v44
                                        case 4294967295:
                                            var v45 int8 = 127
                                            _ = v45
                                        }
                                        v46 := GenS1{F0: 660, F1: 255}
                                        v43 -= (v41 & -552)
                                        _ = v43
                                        _ = v46
                                    } else {
                                        v47 := GenS2{F0: 403, F1: true, F2: 100}
                                        _ = v47
                                    }
                                    switch (("\n\t" + "foo bar") + "hello") {
                                    case "\n\t":
                                        var v48 uint8 = 0
                                        _ = v48
                                    }
                                } else {
                                    if false {
                                        v42 = append(v42, fi36)
                                        v49 := GenS1{F0: (fi36 * (fi38 | 42)), F1: 851}
                                        switch -9223372036854775808 {
                                        case 836:
                                            v50 := GenS3{F0: 255, F1: 169}
                                            _ = v50
                                        case -27:
                                            v51 := GenS1{F0: (fi38 % ((-207) | 1)), F1: 65535}
                                            _ = v51
                                        default:
                                            if v39 {
                                                var v52 uint32 = 65535
                                                _ = v52
                                            }
                                        }
                                        v53 := GenS1{F0: fi36, F1: 4294967295}
                                        if v39 {
                                            v39 = v39
                                            if len(v42) >= 2 {
                                                v42 = v42[1:len(v42)]
                                            }
                                            if (v39 || (v39 || true)) {
                                                if len(v42) > 0 {
                                                    v42[0] = 869
                                                }
                                                var v54 int16 = 0
                                                var v55 int64 = 263
                                                _ = v54
                                                _ = v55
                                            } else {
                                                if v39 {
                                                    switch -9223372036854775808 {
                                                    case 0:
                                                        v40 -= -2.25
                                                        v40 += (3.14159 / 1.0)
                                                        var v56 string = "\n\t"
                                                        _ = v56
                                                    default:
                                                        switch ("hello" + ("a" + ("" + "a"))) {
                                                        case "foo bar":
                                                            if v39 {
                                                                fmt.Println(v40, v40)
                                                                var v57 float32 = ((v40 - 0.1) * v40)
                                                                v58 := GenS4{F0: 2, F1: v39, F2: ""}
                                                                _ = v57
                                                                _ = v58
                                                            } else {
                                                                var v59 bool = (255 == 74)
                                                                _ = v59
                                                            }
                                                        case "":
                                                            v42 = append(v42, (fi38 + genFn0(65535, 242, 127)))
                                                        case "a":
                                                            var v60 int = (fi38 % (((^(fi36))) | 1))
                                                            _ = v60
                                                        }
                                                    }
                                                    v61 := GenS4{F0: 42, F1: v39, F2: ""}
                                                    var v62 int32 = (v37 | (v41 + (v41 | v41)))
                                                    _ = v61
                                                    _ = v62
                                                }
                                            }
                                            switch genFn0(255, 124, 42) {
                                            case 2147483647:
                                                var v63 uint = 2
                                                _ = v63
                                            case 0:
                                                var v64 int64 = -9223372036854775808
                                                _ = v64
                                            case 4294967295:
                                                if false {
                                                    var v65 int16 = -1
                                                    v66 := GenS5{F0: v65, F1: (v65 / (((v65 >> ((205) & 63))) | 1))}
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    if v39 {
                                                        v40 = v40
                                                    }
                                                }
                                            }
                                        } else {
                                            if true {
                                                switch (fi36 >> ((306) & 63)) {
                                                case 42:
                                                    if (-546 < 367) {
                                                        var v67 int8 = -1
                                                        _ = v67
                                                    }
                                                case 0:
                                                    fmt.Println(fi36, fi36, v37)
                                                }
                                            } else {
                                                var v68 uint8 = 255
                                                if v39 {
                                                    _ = len(v42)
                                                } else {
                                                    switch ("a" + "foo bar") {
                                                    case "\n\t":
                                                        v37 = 0
                                                    }
                                                }
                                                _ = v68
                                            }
                                            if (((!(!true)) || true) || true) {
                                                v37 += 238
                                            } else {
                                                var v69 uint = 18446744073709551615
                                                var v70 int64 = 125
                                                _ = v69
                                                _ = v70
                                            }
                                        }
                                        _ = v49
                                        _ = v53
                                    }
                                    v37 -= v37
                                }
                                if v39 {
                                    var v71 float64 = -2.25
                                    if v39 {
                                        if v39 {
                                            switch (fi38 ^ fi38) {
                                            case 4294967295:
                                                v71 += (1.5 / ((v71 / (v71 * 1e10)) + v71))
                                            case -2147483648:
                                                v71 += (float64((3.14159 + 1e10)) / 3.14159)
                                            }
                                            var v72 bool = (((v41 * v37) != 42) || (!(((v39 || ((!(-991 >= 42)) || (fi36 > fi36))) || (0 == 0)) && (!(!(v41 < (^(v41))))))))
                                            v73 := GenS2{F0: 65535, F1: ((true || (!(-1 == -122))) || false), F2: 100}
                                            v74 := GenS6{F0: 1000, F1: 42}
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                    } else {
                                        fmt.Println(v39)
                                        var v75 bool = v39
                                        _ = v75
                                    }
                                    var v76 uint = 255
                                    _ = v71
                                    _ = v76
                                } else {
                                    v40 += (1e10 + 1.5)
                                    switch fi36 {
                                    case 32767:
                                        v77 := GenS2{F0: 42, F1: v39, F2: 42}
                                        _ = v77
                                    default:
                                        v41 = -2147483648
                                    }
                                    switch 0 {
                                    case 1:
                                        var v78 int8 = 0
                                        _ = v78
                                    case -108:
                                        var v79 uint16 = 458
                                        _ = v79
                                    case 0:
                                        v39 = (!(65535 == 255))
                                    default:
                                        v41 -= (v41 << ((698) & 63))
                                    }
                                    var v80 int32 = (v41 >> ((18446744073709551615) & 63))
                                    v80 = v41
                                    _ = v80
                                }
                                if false {
                                    switch (("foo bar" + (("hello" + "hello") + (("a" + "\n\t") + "a"))) + ((("" + ("hello" + "\n\t")) + ("\n\t" + "foo bar")) + "a")) {
                                    case "foo bar":
                                        if len(v42) > 0 {
                                            v42[0] = (-(fi36))
                                        }
                                    case "":
                                        var v81 uint64 = 18446744073709551615
                                        _ = v81
                                    case "hello":
                                        switch "foo bar" {
                                        case "foo bar":
                                            if len(v42) >= 2 {
                                                v42 = v42[1:len(v42)]
                                            }
                                            v82 := GenS7{F0: 635}
                                            _ = v82
                                        case "a":
                                            v42 = append(v42, genFn0(65535, -32768, 0))
                                        }
                                    default:
                                        v41 -= (v41 / (((-(v37))) | 1))
                                    }
                                    v39 = v39
                                    if len(v42) > 0 {
                                        v42[0] = 1
                                    }
                                    v39 = v39
                                    fmt.Println(v41, fi36)
                                } else {
                                    _ = len(v42)
                                    var v83 float32 = ((((1.5 * v40) * v40) + (float32((-2.25 - 0.0)) - v40)) * (v40 * float32(0.1)))
                                    var v84 int8 = 0
                                    _ = v83
                                    _ = v84
                                }
                            } else {
                                v41 -= v37
                                var v85 int = fi38
                                v41 -= (v37 % ((v37) | 1))
                                _ = v85
                            }
                        }
                    }
                    v86 := GenS6{F0: 545, F1: 42}
                    _ = v40
                    _ = v41
                    _ = v42
                    _ = v86
                } else {
                    switch fi38 {
                    case -1:
                        v87 := GenS8{F0: 42, F1: 2}
                        _ = v87
                    case 162:
                        switch (fi38 % ((fi36) | 1)) {
                        case -1:
                            v88 := GenS9{F0: 7, F1: 0}
                            _ = v88
                        case -189:
                            v89 := []int{(fi36 ^ 0), fi36, 0, genFn0(2, -388, 42)}
                            _ = v89
                        case 42:
                            if v39 {
                                v37 -= -152
                            } else {
                                v37 = v37
                                v90 := GenS5{F0: -32768, F1: -780}
                                _ = v90
                            }
                        }
                    }
                }
                _ = fi38
                _ = v39
            }
            for fi91 := 0; fi91 < 2; fi91++ {
                v92 := []int{(fi91 * genFn0(808, 804, 255)), fi91, (fi91 & (fi36 | fi36)), (fi91 & (fi36 / ((fi36) | 1)))}
                if (!false) {
                    switch "hello" {
                    case "foo bar":
                        v93 := GenS9{F0: 42, F1: 836}
                        _ = v93
                    case "\n\t":
                        var v94 uint32 = 2147483648
                        fmt.Println(fi36, fi91)
                        var v95 uint16 = 42
                        _ = v94
                        _ = v95
                    case "hello":
                        var v96 float32 = -2.25
                        _ = v96
                    default:
                        if ((("\n\t" + ("hello" + "foo bar")) > "") && true) {
                            v97 := GenS2{F0: 165, F1: (-1 != 111), F2: -128}
                            v37 += (v37 & 42)
                            var v98 int = (fi36 << ((255) & 63))
                            v37 -= (-(v37))
                            _ = v97
                            _ = v98
                        }
                    }
                    v37 = (v37 * 2147483647)
                } else {
                    _ = len(v92)
                    if false {
                        v37 = v37
                    } else {
                        fmt.Println(fi91, fi36)
                    }
                }
                fmt.Println(v37)
                _ = fi91
                _ = v92
            }
            _ = fi36
            _ = v37
        }
        var v99 bool = (42 <= 1)
        var v100 bool = v99
        _ = v100
        _ = v99
    } else {
        v101 := GenS10{F0: -366, F1: -175, F2: 32767}
        for fi102 := 0; fi102 < 3; fi102++ {
            if false {
                var v103 uint8 = 127
                v103 += v103
                _ = v103
            }
            var v104 int32 = 881
            _ = fi102
            _ = v104
        }
        if true {
            if ("foo bar" != "\n\t") {
                var v105 uint8 = 0
                _ = v105
            }
        } else {
            var v106 int32 = 631
            var v107 int = 451
            if (!true) {
                for fi108 := 0; fi108 < 8; fi108++ {
                    for fi109 := 0; fi109 < 20; fi109++ {
                        v107 += genFn0(922, -1, 0)
                        v107 += 89
                        _ = fi109
                    }
                    var v110 uint64 = 4294967295
                    for fi111 := 0; fi111 < 11; fi111++ {
                        fmt.Println(fi108, v107, fi111)
                        _ = fi111
                    }
                    _ = fi108
                    _ = v110
                }
                fmt.Println(v107)
                if false {
                    switch ("a" + "\n\t") {
                    case "a":
                        v106 += (v106 << ((42) & 63))
                    case "hello":
                        for fi112 := 0; fi112 < 16; fi112++ {
                            v113 := GenS11{F0: "\n\t", F1: 779}
                            _ = fi112
                            _ = v113
                        }
                    case "\n\t":
                        v107 = genFn0(1, 0, 42)
                    }
                    switch "" {
                    case "\n\t":
                        fmt.Println(v107)
                    case "":
                        for fi114 := 0; fi114 < 16; fi114++ {
                            v106 -= v106
                            _ = fi114
                        }
                    }
                    var v115 uint64 = 0
                    _ = v115
                }
            }
            _ = v106
            _ = v107
        }
        _ = v101
    }
    return 1
}

func main() {
    v1 := []string{"", "hello", "foo bar", ""}
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]uint32)
    var v4 uint = 18446744073709551615
    delete(v3, "a")
    if (4294967295 != 2) {
        if ((v2 * 18446744073709551615) == v2) {
            v116 := GenS11{F0: "", F1: 32767}
            for fi117 := 0; fi117 < 7; fi117++ {
                v4 += (v4 / (((v4 >> (((^(v4))) & 63))) | 1))
                var v118 int = 0
                for fi119 := 0; fi119 < 18; fi119++ {
                    switch (fi117 - v118) {
                    case 0:
                        fmt.Println(fi117)
                    case 2147483647:
                        _ = len(v3)
                        v120 := v3["foo bar"]
                        switch "" {
                        case "":
                            fmt.Println(v118)
                        case "\n\t":
                            switch "\n\t" {
                            case "\n\t":
                                delete(v3, "a")
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                            case "a":
                                switch 485 {
                                case 32768:
                                    var v121 int64 = 81
                                    _ = v121
                                case 990:
                                    var v122 string = ((("a" + "\n\t") + "a") + ("\n\t" + "a"))
                                    _ = v122
                                default:
                                    if false {
                                        v4 = (v4 / (((v4 % ((2) | 1))) | 1))
                                    } else {
                                        fmt.Println(v118)
                                    }
                                }
                            }
                            var v123 float32 = 3.14159
                            _ = v123
                        default:
                            switch "foo bar" {
                            case "a":
                                var v124 uint32 = v120
                                _ = v124
                            case "\n\t":
                                var v125 uint8 = 50
                                _ = v125
                            case "foo bar":
                                v120 -= v120
                                v1 = append(v1, ("\n\t" + (("" + "foo bar") + "foo bar")))
                                if false {
                                    var v126 uint = v4
                                    v127 := GenS12{F0: (fi119 - 434), F1: 1, F2: 1}
                                    v120 += (v120 >> ((v4) & 63))
                                    v2 -= v2
                                    _ = v126
                                    _ = v127
                                } else {
                                    v4 += v4
                                    if (42 != 127) {
                                        switch 128 {
                                        case 287:
                                            if false {
                                                if (1 < 151) {
                                                    v4 += v4
                                                    v128 := v3["foo bar"]
                                                    if (("foo bar" + "foo bar") < "hello") {
                                                        var v129 uint8 = 42
                                                        _ = v129
                                                    } else {
                                                        v4 = v4
                                                        var v130 int32 = -2147483648
                                                        _ = v130
                                                    }
                                                    _ = v128
                                                }
                                            }
                                        default:
                                            delete(v3, "a")
                                        }
                                        var v131 int64 = 778
                                        _ = v131
                                    }
                                    if true {
                                        switch "a" {
                                        case "":
                                            v4 -= (^(v4))
                                        case "\n\t":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "a":
                                            v132 := GenS13{F0: (^(v4))}
                                            _ = v132
                                        default:
                                            var v133 float32 = -2.25
                                            _ = v133
                                        }
                                        delete(v3, (("" + "foo bar") + "a"))
                                        v4 = (v4 * (v4 - v4))
                                    }
                                }
                            }
                        }
                        _ = v120
                    case -2147483648:
                        switch ("a" + "") {
                        case "hello":
                            switch -2147483648 {
                            case 69:
                                v4 -= 65535
                            case -9223372036854775808:
                                if false {
                                    v1 = append(v1, "\n\t")
                                }
                            case 0:
                                var v134 float64 = (1.5 - (((0.1 / -2.25) - (((float64((-2.25 - 0.0)) - 3.14159) - ((0.0 * -2.25) * ((float64(-2.25) / 1e10) * 0.0))) * 1.5)) / 3.14159))
                                _ = v134
                            }
                        }
                    default:
                        v118 = (fi119 + v118)
                    }
                    var v135 int16 = -515
                    if false {
                        if (430 <= 32767) {
                            v2 = v2
                        }
                        var v136 float32 = 0.1
                        _ = v136
                    }
                    _ = fi119
                    _ = v135
                }
                _ = fi117
                _ = v118
            }
            var v137 uint16 = 65535
            _ = v116
            _ = v137
        } else {
            var v138 float64 = (3.14159 - 0.1)
            _ = v138
        }
    } else {
        switch "" {
        case "a":
            switch -902 {
            case -9223372036854775808:
                v4 += 65535
            default:
                if true {
                    fmt.Println(v2, v4)
                    var v139 int64 = -140
                    _ = v139
                } else {
                    switch -369 {
                    case -1:
                        for fi140 := 0; fi140 < 8; fi140++ {
                            v2 = 457
                            for fi141 := 0; fi141 < 16; fi141++ {
                                v4 -= v4
                                switch "foo bar" {
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                case "":
                                    if true {
                                        var v142 uint8 = 160
                                        fmt.Println(v2, fi140, v4)
                                        v2 = v2
                                        _ = v142
                                    } else {
                                        switch ("a" + ("" + "foo bar")) {
                                        case "\n\t":
                                            var v143 int = -1
                                            _ = v143
                                        case "":
                                            v144 := GenS10{F0: 1000, F1: fi140, F2: 0}
                                            _ = v144
                                        case "foo bar":
                                            var v145 int32 = 42
                                            _ = v145
                                        default:
                                            var v146 int8 = 42
                                            _ = v146
                                        }
                                        v147 := GenS14{F0: "foo bar"}
                                        _ = v147
                                    }
                                case "\n\t":
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                }
                                var v148 int64 = -988
                                if (!(2147483648 >= 2147483648)) {
                                    if true {
                                        v149 := v3[((("" + "a") + ((("hello" + ("foo bar" + "foo bar")) + "a") + "\n\t")) + "\n\t")]
                                        _ = v149
                                    } else {
                                        v150 := GenS2{F0: 449, F1: (int32(fi140) != -2147483648), F2: -1}
                                        switch ("foo bar" + "a") {
                                        case "":
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v151 float32 = (1e10 / 0.1)
                                                _ = v151
                                            case "":
                                                if true {
                                                    v152 := GenS2{F0: 2147483647, F1: true, F2: 127}
                                                    if (!true) {
                                                        fmt.Println(v4, fi141, v148)
                                                        v153 := GenS5{F0: -32768, F1: 1}
                                                        _ = v153
                                                    } else {
                                                        var v154 int64 = (v148 << ((v4) & 63))
                                                        fmt.Println(fi141)
                                                        _ = v154
                                                    }
                                                    var v155 int8 = -112
                                                    _ = v152
                                                    _ = v155
                                                } else {
                                                    var v156 uint32 = 66
                                                    var v157 uint16 = 321
                                                    _ = v156
                                                    _ = v157
                                                }
                                            case "a":
                                                v148 -= (v148 & -1)
                                            }
                                        case "a":
                                            fmt.Println(v148)
                                        default:
                                            var v158 int = genFn0(32767, -32768, 128)
                                            _ = v158
                                        }
                                        _ = v150
                                    }
                                    v2 = (^(v2))
                                    var v159 uint16 = 1
                                    var v160 bool = (v2 < (v2 >> ((65535) & 63)))
                                    switch (fi140 ^ fi141) {
                                    case 604:
                                        v159 += v159
                                        v3[""] = 893
                                        fmt.Println(v148, v2, v160)
                                        var v161 int64 = (v148 - 215)
                                        _ = v161
                                    case 126:
                                        var v162 uint32 = 65535
                                        _ = v162
                                    default:
                                        switch fi141 {
                                        case 0:
                                            var v163 int16 = -606
                                            _ = v163
                                        }
                                    }
                                    var v164 int16 = -32768
                                    v165 := GenS15{F0: (v164 - (v164 ^ v164))}
                                    _ = v159
                                    _ = v160
                                    _ = v164
                                    _ = v165
                                }
                                _ = fi141
                                _ = v148
                            }
                            _ = fi140
                        }
                    case 32767:
                        var v166 uint16 = 255
                        _ = v166
                    default:
                        if (9223372036854775807 >= 1) {
                            switch genFn0(42, -32768, 1) {
                            case 42:
                                v167 := GenS15{F0: -1}
                                _ = v167
                            case 74:
                                v2 -= (v2 | 18446744073709551615)
                            case 282:
                                v3["hello"] = 42
                            default:
                                for fi168 := 0; fi168 < 3; fi168++ {
                                    if (v4 >= (v4 - uint(v2))) {
                                        var v169 uint8 = 1
                                        v170 := GenS16{F0: false, F1: "\n\t"}
                                        v4 = v4
                                        _ = v169
                                        _ = v170
                                    }
                                    _ = fi168
                                }
                            }
                        } else {
                            for fi171 := 0; fi171 < 17; fi171++ {
                                if (972 != 1) {
                                    for fi172 := 0; fi172 < 4; fi172++ {
                                        v2 += (v2 ^ (v2 << ((462) & 63)))
                                        _ = fi172
                                    }
                                } else {
                                    for fi173 := 0; fi173 < 4; fi173++ {
                                        if true {
                                            switch ("hello" + "foo bar") {
                                            case "foo bar":
                                                if false {
                                                    v2 = 0
                                                    if ((false || false) || (((!(false || (v2 <= v2))) || true) && true)) {
                                                        if true {
                                                            v2 += v2
                                                            var v174 int32 = -2147483648
                                                            _ = v174
                                                        } else {
                                                            v175 := GenS6{F0: 1, F1: 42}
                                                            _ = v175
                                                        }
                                                        switch (fi173 + int(v4)) {
                                                        case 127:
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                switch -9223372036854775808 {
                                                                case 9223372036854775807:
                                                                    v176 := GenS17{F0: (-128 != -83), F1: 1}
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v176
                                                                default:
                                                                    if (2 <= 65535) {
                                                                        v177 := v3[""]
                                                                        var v178 uint32 = (v177 + 234)
                                                                        if false {
                                                                            v179 := GenS7{F0: (v2 / ((v2) | 1))}
                                                                            v4 += (v4 / ((0) | 1))
                                                                            v180 := GenS18{F0: 2147483647, F1: -1}
                                                                            _ = v179
                                                                            _ = v180
                                                                        } else {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = ("foo bar" + "hello")
                                                                            }
                                                                            fmt.Println(fi171, v4)
                                                                            if (1 >= 71) {
                                                                                var v181 uint8 = 255
                                                                                if (false || false) {
                                                                                    if true {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            fmt.Println(v181)
                                                                                        case "":
                                                                                            var v182 uint64 = (v2 + 1)
                                                                                            _ = v182
                                                                                        case "hello":
                                                                                            if true {
                                                                                                fmt.Println(fi171)
                                                                                                switch fi171 {
                                                                                                case 128:
                                                                                                    v183 := GenS2{F0: (v178 - (v178 & 824)), F1: true, F2: 42}
                                                                                                    _ = v183
                                                                                                case -321:
                                                                                                    var v184 uint8 = 2
                                                                                                    _ = v184
                                                                                                }
                                                                                                v178 = 2147483648
                                                                                                var v185 int8 = 0
                                                                                                _ = v185
                                                                                            } else {
                                                                                                v178 += (v178 ^ (v177 << ((1) & 63)))
                                                                                            }
                                                                                        }
                                                                                        switch 827 {
                                                                                        case -2147483648:
                                                                                            if false {
                                                                                                if true {
                                                                                                    var v186 uint32 = (v178 / ((0) | 1))
                                                                                                    _ = v186
                                                                                                }
                                                                                            }
                                                                                        case 128:
                                                                                            v177 = (^(v178))
                                                                                        }
                                                                                        if false {
                                                                                            if (!(!false)) {
                                                                                                var v187 int32 = 1
                                                                                                _ = v187
                                                                                            }
                                                                                            if (((!((-781 < 1) && (("" + "hello") <= ""))) || (fi173 == genFn1())) || true) {
                                                                                                v3["foo bar"] = v177
                                                                                                v188 := GenS4{F0: 398, F1: (2147483647 > 42), F2: ""}
                                                                                                if (fi171 < genFn0(588, 520, (v181 - v181))) {
                                                                                                    if false {
                                                                                                        v4 = 0
                                                                                                        if ((-(fi171)) > fi173) {
                                                                                                            if (false || ((0.1 + 1e10) > 3.14159)) {
                                                                                                                var v189 float64 = (1e10 / 3.14159)
                                                                                                                v189 -= 0.0
                                                                                                                _ = v189
                                                                                                            } else {
                                                                                                                switch "hello" {
                                                                                                                case "foo bar":
                                                                                                                    v181 += uint8(v2)
                                                                                                                case "a":
                                                                                                                    switch "" {
                                                                                                                    case "":
                                                                                                                        v181 += 127
                                                                                                                    }
                                                                                                                }
                                                                                                                var v190 uint64 = (v2 + (v2 % ((v2) | 1)))
                                                                                                                if ((fi173 + (fi171 + genFn1())) < (fi171 + fi173)) {
                                                                                                                    if true {
                                                                                                                        var v191 int = (fi171 | 42)
                                                                                                                        v177 -= v177
                                                                                                                        v192 := GenS7{F0: uint64(v178)}
                                                                                                                        _ = v191
                                                                                                                        _ = v192
                                                                                                                    } else {
                                                                                                                        fmt.Println(v181)
                                                                                                                        fmt.Println(v2)
                                                                                                                    }
                                                                                                                    switch (fi171 + genFn0(0, 1000, (v181 << (((v4 + v4)) & 63)))) {
                                                                                                                    case -9223372036854775808:
                                                                                                                        switch ("hello" + "") {
                                                                                                                        case "":
                                                                                                                            v178 = (^(v178))
                                                                                                                        default:
                                                                                                                            delete(v3, "\n\t")
                                                                                                                        }
                                                                                                                    case 9:
                                                                                                                        var v193 float64 = ((float64(((1.5 / 0.1) / 1e10)) * -2.25) * 1.5)
                                                                                                                        v190 = (v2 ^ (v2 | (v190 / ((v2) | 1))))
                                                                                                                        v177 = (v178 ^ uint32(v181))
                                                                                                                        _ = v193
                                                                                                                    case -914:
                                                                                                                        var v194 uint8 = (^(v181))
                                                                                                                        _ = v194
                                                                                                                    default:
                                                                                                                        fmt.Println(v190, v190, v2)
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v190
                                                                                                            }
                                                                                                        } else {
                                                                                                            fmt.Println(v181)
                                                                                                            v177 = v177
                                                                                                            fmt.Println(v181, v181)
                                                                                                        }
                                                                                                    }
                                                                                                    v195 := GenS13{F0: v4}
                                                                                                    _ = v195
                                                                                                } else {
                                                                                                    var v196 bool = (!(!(772 >= v178)))
                                                                                                    var v197 string = ""
                                                                                                    _ = v196
                                                                                                    _ = v197
                                                                                                }
                                                                                                var v198 bool = false
                                                                                                _ = v188
                                                                                                _ = v198
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v199 float64 = float64(3.14159)
                                                                                        _ = v199
                                                                                    }
                                                                                    v178 += (v177 << ((137) & 63))
                                                                                    var v200 int8 = -16
                                                                                    _ = v200
                                                                                }
                                                                                switch genFn0(255, 0, v181) {
                                                                                case -1:
                                                                                    var v201 uint = (v4 >> ((18446744073709551615) & 63))
                                                                                    v2 = (v2 / ((v2) | 1))
                                                                                    _ = v201
                                                                                case -866:
                                                                                    switch (fi173 % (((fi171 << ((18446744073709551615) & 63))) | 1)) {
                                                                                    case 9223372036854775807:
                                                                                        var v202 float64 = 1.5
                                                                                        _ = v202
                                                                                    case 128:
                                                                                        v203 := GenS13{F0: v4}
                                                                                        _ = v203
                                                                                    default:
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        case "\n\t":
                                                                                            if ((v4 >> ((v4) & 63)) > 255) {
                                                                                                v178 = v177
                                                                                                if false {
                                                                                                    v204 := GenS7{F0: 4294967295}
                                                                                                    _ = v204
                                                                                                } else {
                                                                                                    var v205 int = fi173
                                                                                                    _ = v205
                                                                                                }
                                                                                            } else {
                                                                                                v2 -= v2
                                                                                                var v206 int64 = 0
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    v178 -= (v177 ^ v178)
                                                                                                }
                                                                                                _ = v206
                                                                                            }
                                                                                        default:
                                                                                            v4 = v4
                                                                                        }
                                                                                    }
                                                                                case 2147483647:
                                                                                    switch genFn1() {
                                                                                    case 32768:
                                                                                        switch fi173 {
                                                                                        case 2147483647:
                                                                                            v207 := v3[("" + "\n\t")]
                                                                                            _ = v207
                                                                                        }
                                                                                        v177 = v178
                                                                                        fmt.Println(fi173, v4, v2)
                                                                                    default:
                                                                                        if true {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            if (!((((("a" + (("a" + "foo bar") + "\n\t")) + ("\n\t" + "hello")) + "\n\t") + "") <= "")) {
                                                                                                switch "foo bar" {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v178, v4)
                                                                                                    var v208 float32 = (1e10 * -2.25)
                                                                                                    v1 = append(v1, (("\n\t" + "a") + ""))
                                                                                                    _ = v208
                                                                                                default:
                                                                                                    var v209 int16 = -1
                                                                                                    _ = v209
                                                                                                }
                                                                                                v178 = (v178 + (v177 * v177))
                                                                                            }
                                                                                            if (false && (!(349 != 401))) {
                                                                                                if (!true) {
                                                                                                    if (true || false) {
                                                                                                        v177 = 4294967295
                                                                                                        v181 -= v181
                                                                                                    }
                                                                                                    var v210 uint64 = 15
                                                                                                    var v211 int = -9223372036854775808
                                                                                                    _ = v210
                                                                                                    _ = v211
                                                                                                } else {
                                                                                                    delete(v3, ("a" + "\n\t"))
                                                                                                    v212 := GenS15{F0: 42}
                                                                                                    _ = v212
                                                                                                }
                                                                                                if true {
                                                                                                    var v213 int16 = 657
                                                                                                    if false {
                                                                                                        var v214 uint32 = 2147483648
                                                                                                        _ = v214
                                                                                                    }
                                                                                                    _ = v213
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    if (!(true || (!((-1 >= -766) || true)))) {
                                                                                        _ = len(v3)
                                                                                    } else {
                                                                                        var v215 string = ("foo bar" + ("a" + ("" + ("hello" + "hello"))))
                                                                                        fmt.Println(v178)
                                                                                        _ = v215
                                                                                    }
                                                                                }
                                                                                _ = v181
                                                                            } else {
                                                                                v216 := GenS16{F0: ((fi171 % ((fi171) | 1)) >= (fi171 << ((v4) & 63))), F1: ("foo bar" + (("a" + "\n\t") + "foo bar"))}
                                                                                v178 += (v178 - v178)
                                                                                if (-9223372036854775808 >= 1) {
                                                                                    v177 -= 61
                                                                                } else {
                                                                                    v217 := GenS19{F0: -2147483648, F1: "hello", F2: fi173}
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = ""
                                                                                    }
                                                                                    v178 += v178
                                                                                    _ = v217
                                                                                }
                                                                                _ = v216
                                                                            }
                                                                            if false {
                                                                                if (!((v4 / ((18446744073709551615) | 1)) < (^(v4)))) {
                                                                                    v178 += 1
                                                                                    v218 := GenS7{F0: (v2 & v2)}
                                                                                    switch (fi173 / ((-78) | 1)) {
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(fi173, fi173, v177)
                                                                                        v177 = v178
                                                                                    case 42:
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = "\n\t"
                                                                                        }
                                                                                    case 32768:
                                                                                        v178 = (^(v177))
                                                                                    default:
                                                                                        var v219 string = ""
                                                                                        _ = v219
                                                                                    }
                                                                                    _ = v218
                                                                                } else {
                                                                                    switch (fi173 | (fi173 + (fi171 % ((fi173) | 1)))) {
                                                                                    case 4294967295:
                                                                                        var v220 float32 = 1.5
                                                                                        _ = v220
                                                                                    }
                                                                                    v1 = append(v1, (((("\n\t" + "foo bar") + ("hello" + ("hello" + (((("a" + "a") + ("a" + "")) + ("a" + "a")) + "")))) + "a") + "foo bar"))
                                                                                    v2 += v2
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v177
                                                                        _ = v178
                                                                    } else {
                                                                        v2 = (v2 % ((v2) | 1))
                                                                        var v221 uint8 = 42
                                                                        fmt.Println(v4, fi171, v2)
                                                                        _ = v221
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                var v222 int64 = 1
                                                                _ = v222
                                                            }
                                                        case 128:
                                                            switch genFn0(255, -405, 255) {
                                                            case 128:
                                                                v223 := GenS20{F0: "", F1: -534}
                                                                _ = v223
                                                            case -397:
                                                                var v224 uint16 = 512
                                                                _ = v224
                                                            }
                                                            v2 = (^(v2))
                                                            var v225 int32 = 56
                                                            _ = v225
                                                        }
                                                    } else {
                                                        if (!((v2 | 18446744073709551615) >= 267)) {
                                                            switch (fi173 * fi173) {
                                                            case 1:
                                                                v226 := GenS7{F0: v2}
                                                                _ = v226
                                                            case -216:
                                                                var v227 float32 = -2.25
                                                                _ = v227
                                                            case 731:
                                                                v1 = append(v1, "foo bar")
                                                            }
                                                        } else {
                                                            if (true || (v4 < (v4 >> ((v4) & 63)))) {
                                                                v228 := GenS21{F0: (v4 >> ((4294967295) & 63))}
                                                                v2 += v2
                                                                v229 := GenS22{F0: fi171}
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v228
                                                                _ = v229
                                                            }
                                                            v2 = v2
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    v230 := GenS13{F0: (^(v4))}
                                                    var v231 string = "\n\t"
                                                    v231 = v231
                                                    _ = v230
                                                    _ = v231
                                                } else {
                                                    if true {
                                                        var v232 uint16 = 1
                                                        _ = v232
                                                    }
                                                    var v233 uint16 = 65535
                                                    _ = v233
                                                }
                                            default:
                                                v234 := GenS2{F0: 42, F1: (2147483647 >= 2), F2: -128}
                                                _ = v234
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                switch ("\n\t" + "") {
                                                case "hello":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            case "foo bar":
                                                var v235 float32 = (0.0 / float32(-2.25))
                                                _ = v235
                                            case "a":
                                                if true {
                                                    v236 := GenS11{F0: ("a" + ""), F1: 255}
                                                    _ = v236
                                                } else {
                                                    v4 = 65535
                                                    fmt.Println(fi173, v4)
                                                    if (127 >= -16) {
                                                        var v237 int8 = -95
                                                        _ = v237
                                                    }
                                                }
                                            }
                                            v1 = append(v1, ("hello" + "a"))
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if true {
                                            var v238 uint8 = 42
                                            v239 := v3[(("\n\t" + "hello") + "\n\t")]
                                            _ = v238
                                            _ = v239
                                        } else {
                                            var v240 int32 = 686
                                            _ = v240
                                        }
                                        v2 -= 1
                                        _ = fi173
                                    }
                                    fmt.Println(v4, fi171, v4)
                                    var v241 string = "hello"
                                    _ = v241
                                }
                                _ = fi171
                            }
                            for fi242 := 0; fi242 < 14; fi242++ {
                                for fi243 := 0; fi243 < 14; fi243++ {
                                    v4 = v4
                                    v244 := GenS22{F0: (fi242 + fi243)}
                                    _ = fi243
                                    _ = v244
                                }
                                var v245 uint16 = 2
                                _ = fi242
                                _ = v245
                            }
                            var v246 uint = v4
                            _ = v246
                        }
                    }
                    for fi247 := 0; fi247 < 5; fi247++ {
                        v2 = (^(v2))
                        var v248 int64 = -9223372036854775808
                        v248 = (v248 >> (((v4 | (v4 << (((^(v4))) & 63)))) & 63))
                        _ = fi247
                        _ = v248
                    }
                }
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    if true {
        v4 -= (v4 & v4)
        v2 = v2
    } else {
        if (65535 > 0) {
            v4 += v4
            v4 = (^(v4))
        }
        for fi249 := 0; fi249 < 9; fi249++ {
            var v250 uint64 = v2
            fmt.Println(fi249)
            if true {
                v1 = append(v1, "a")
                if ("\n\t" > (("foo bar" + "\n\t") + "a")) {
                    v2 -= uint64(v4)
                    if true {
                        v4 = v4
                        fmt.Println(v2, v2, v4)
                    } else {
                        v251 := GenS17{F0: true, F1: 42}
                        for fi252 := 0; fi252 < 15; fi252++ {
                            v4 = 2
                            if false {
                                if true {
                                    v253 := GenS23{F0: (v2 - (^(v250))), F1: 334}
                                    _ = v253
                                }
                            } else {
                                v254 := GenS24{F0: 0, F1: -14}
                                _ = v254
                            }
                            v2 = v2
                            switch (((("hello" + "hello") + "a") + ("\n\t" + "")) + "foo bar") {
                            case "hello":
                                v3["foo bar"] = 42
                            }
                            _ = fi252
                        }
                        _ = v251
                    }
                } else {
                    var v255 int16 = 1
                    v2 += v2
                    if false {
                        for fi256 := 0; fi256 < 9; fi256++ {
                            var v257 bool = false
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(fi249, v4)
                            case "hello":
                                switch (-(fi249)) {
                                case 765:
                                    var v258 uint8 = 2
                                    _ = v258
                                default:
                                    v257 = ((fi249 - (fi256 - (^(fi256)))) != (fi256 | (fi249 & genFn0(32767, v255, 42))))
                                }
                            case "":
                                if v257 {
                                    v1 = append(v1, "hello")
                                    var v259 int16 = 1
                                    switch genFn1() {
                                    case 1:
                                        switch (fi256 >> (((v4 * v4)) & 63)) {
                                        case -1:
                                            var v260 int16 = (v255 | v259)
                                            _ = v260
                                        case 419:
                                            fmt.Println(v259, fi249, v257)
                                        case -812:
                                            delete(v3, ("" + (((("hello" + "foo bar") + "") + "") + "foo bar")))
                                        default:
                                            v250 = uint64(v4)
                                        }
                                    default:
                                        v255 += v259
                                    }
                                    _ = v259
                                }
                                fmt.Println(fi256, v250, v257)
                                v2 = v250
                                v261 := GenS25{F0: 1, F1: 9223372036854775807, F2: fi256}
                                v4 = (v4 + (v4 - 65535))
                                _ = len(v3)
                                v250 -= v250
                                _ = v261
                            }
                            _ = fi256
                            _ = v257
                        }
                        for fi262 := 0; fi262 < 20; fi262++ {
                            v255 -= 841
                            v1 = append(v1, ("foo bar" + ("\n\t" + "hello")))
                            var v263 int64 = -9223372036854775808
                            _ = fi262
                            _ = v263
                        }
                    } else {
                        fmt.Println(v4)
                        v2 = v2
                        fmt.Println(v250)
                    }
                    _ = v255
                }
            }
            _ = fi249
            _ = v250
        }
        var v264 int16 = 0
        _ = v264
    }
    for fi265 := 0; fi265 < 1; fi265++ {
        fmt.Println(v4, v2, v2)
        var v266 int16 = 32767
        _ = fi265
        _ = v266
    }
    var v267 int8 = 42
    switch "" {
    case "a":
        if (true && ((v4 ^ v4) == (v4 ^ v4))) {
            v268 := GenS22{F0: 2147483647}
            v269 := GenS26{F0: (v4 + v4)}
            fmt.Println(v4, v2)
            _ = v268
            _ = v269
        } else {
            for fi270 := 0; fi270 < 14; fi270++ {
                v2 = (v2 + (v2 | (^(v2))))
                _ = fi270
            }
            var v271 bool = true
            _ = v271
        }
    case "":
        v272 := GenS24{F0: -1, F1: v267}
        _ = v272
    case "\n\t":
        switch 128 {
        case -321:
            for fi273 := 0; fi273 < 7; fi273++ {
                v274 := v3["\n\t"]
                v267 = (-(v267))
                _ = fi273
                _ = v274
            }
        case 0:
            v275 := GenS27{F0: 0, F1: -294, F2: 599}
            _ = v275
        case -442:
            v276 := GenS16{F0: (826 != -9223372036854775808), F1: "hello"}
            _ = v276
        default:
            v277 := GenS27{F0: 61, F1: int16(v267), F2: 700}
            _ = v277
        }
    default:
        var v278 int64 = -1
        _ = v278
    }
    if (42 == 100) {
        v279 := GenS12{F0: -723, F1: 42, F2: 390}
        for fi280 := 0; fi280 < 19; fi280++ {
            v267 = v267
            _ = fi280
        }
        v4 -= (v4 ^ 65535)
        _ = v279
    }
    v267 = (v267 ^ -128)
    _ = len(v3)
    v281 := GenS19{F0: 0, F1: (("" + ("\n\t" + "")) + "\n\t"), F2: genFn0(255, 42, 97)}
    for fi282 := 0; fi282 < 1; fi282++ {
        v283 := GenS16{F0: ((true || true) && false), F1: "hello"}
        v267 += (v267 - v267)
        _ = fi282
        _ = v283
    }
    delete(v3, ("" + "\n\t"))
    delete(v3, "foo bar")
    switch ("\n\t" + "") {
    case "":
        switch 32768 {
        case -863:
            if (!((v2 ^ 42) != (v2 / (((v2 & 2)) | 1)))) {
                var v284 int8 = -1
                _ = v284
            }
        case -9223372036854775808:
            var v285 float64 = float64(1e10)
            _ = v285
        case 127:
            v4 = v4
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "foo bar":
        v286 := GenS6{F0: -32768, F1: 1}
        _ = v286
    }
    for fi287 := 0; fi287 < 3; fi287++ {
        var v288 int16 = 42
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        fmt.Println(v2, fi287, v2)
        _ = fi287
        _ = v288
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        switch 42 {
        case 32768:
            var v289 int = 856
            _ = v289
        case -9223372036854775808:
            for fi290 := 0; fi290 < 2; fi290++ {
                v267 -= 14
                _ = fi290
            }
        case 42:
            fmt.Println(v2, v4, v4)
        default:
            v2 += (v2 * 42)
        }
    case "":
        var v291 uint = v4
        _ = v291
    case "\n\t":
        v2 -= 18446744073709551615
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v267:", v267)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v267
    _ = v281
    _ = v3
    _ = v4
}
