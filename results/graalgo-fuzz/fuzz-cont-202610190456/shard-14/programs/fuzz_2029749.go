package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 uint; F2 uint64 }
type GenS9 struct { F0 string; F1 string; F2 bool }
type GenS10 struct { F0 string; F1 uint32; F2 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64; F1 uint32; F2 string }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int8; F1 uint; F2 string }
type GenS17 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 uint) string {
    v4 := GenS1{F0: -85}
    for fi5 := 0; fi5 < 19; fi5++ {
        p0 -= p0
        for fi6 := 0; fi6 < 2; fi6++ {
            switch (fi6 << (((p0 + (^(p0)))) & 63)) {
            case 32768:
                var v7 uint16 = 0
                _ = v7
            case 128:
                p0 = 740
                p0 = (p0 >> ((421) & 63))
                switch -1 {
                case 0:
                    v8 := GenS2{F0: 4294967295, F1: "foo bar"}
                    _ = v8
                case 9223372036854775807:
                    var v9 int = (-(fi5))
                    _ = v9
                case 42:
                    var v10 uint32 = uint32(p0)
                    _ = v10
                }
            }
            fmt.Println(p0)
            p0 = p0
            _ = fi6
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            p0 = p0
            v12 := []int{fi11}
            _ = fi11
            _ = v12
        }
        _ = fi5
    }
    switch "" {
    case "a":
        var v13 uint64 = 1
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS3{F0: false, F1: (^(p0))}
            p0 = 255
            _ = fi14
            _ = v15
        }
    case "":
        var v16 int = 2147483647
        _ = v16
    default:
        if (false || (812 == 42)) {
            for fi17 := 0; fi17 < 4; fi17++ {
                switch "a" {
                case "a":
                    var v18 bool = true
                    _ = v18
                case "":
                    fmt.Println(p0, p0)
                }
                _ = fi17
            }
        } else {
            if true {
                fmt.Println(p0)
            }
            switch 32767 {
            case 4294967295:
                if false {
                    switch -893 {
                    case 1:
                        switch -907 {
                        case 2147483647:
                            if ("hello" > "a") {
                                p0 = 674
                            }
                        case -981:
                            if true {
                                p0 += p0
                                if (uint8(p0) <= 1) {
                                    v19 := GenS4{F0: (p0 / ((1) | 1)), F1: 319, F2: 136}
                                    _ = v19
                                } else {
                                    for fi20 := 0; fi20 < 4; fi20++ {
                                        p0 -= p0
                                        var v21 uint = (p0 % ((p0) | 1))
                                        _ = fi20
                                        _ = v21
                                    }
                                }
                            }
                        }
                    }
                    p0 = 836
                } else {
                    var v22 int = 32768
                    var v23 int8 = -128
                    v24 := GenS3{F0: (!(32767 == 0)), F1: p0}
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
        }
    }
    p0 = 364
    for fi25 := 0; fi25 < 5; fi25++ {
        p0 = (p0 - p0)
        _ = fi25
    }
    _ = p0
    _ = v4
    return "\n\t"
}
func genFn1(p0 uint8, p1 string) uint8 {
    for fi26 := 0; fi26 < 7; fi26++ {
        v27 := []int{fi26, fi26}
        _ = fi26
        _ = v27
    }
    v28 := []int{1, 4294967295}
    fmt.Println(p1)
    switch ((((((p1 + p1) + p1) + p1) + (p1 + "")) + p1) + p1) {
    case "foo bar":
        fmt.Println(p1, p1)
    case "":
        var v29 bool = false
        _ = v29
    default:
        if len(v28) > 0 {
            v28[0] = 4294967295
        }
    }
    v30 := GenS1{F0: -128}
    _ = p0
    _ = p1
    _ = v28
    _ = v30
    return 2
}

func main() {
    var v1 int = 2147483647
    v2 := []string{}
    v3 := []string{(("hello" + (("" + "hello") + (("a" + "hello") + "foo bar"))) + "\n\t"), (("a" + ("\n\t" + "")) + "foo bar"), ""}
    for fi31 := 0; fi31 < 4; fi31++ {
        if (true || (v1 >= fi31)) {
            if (((1e10 * 0.0) <= 0.1) && (42 <= 127)) {
                v32 := GenS4{F0: 42, F1: 2, F2: 543}
                v33 := GenS5{F0: 1}
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 14; fi34++ {
                v1 = (fi34 * (fi34 >> ((165) & 63)))
                var v35 uint = 740
                v36 := GenS6{F0: 100, F1: (v35 << ((2) & 63)), F2: ""}
                _ = fi34
                _ = v35
                _ = v36
            }
        } else {
            if true {
                for fi37 := 0; fi37 < 14; fi37++ {
                    if true {
                        if (1 < 2) {
                            v1 += fi37
                            v1 += fi37
                        }
                        v1 += (-(v1))
                    }
                    v38 := GenS7{F0: 2}
                    if ((fi37 - 0) != (fi31 << ((2) & 63))) {
                        var v39 float64 = 1.5
                        fmt.Println(v39, fi37)
                        v40 := GenS8{F0: 42, F1: 255, F2: 42}
                        _ = v39
                        _ = v40
                    }
                    v2 = append(v2, "")
                    var v41 uint8 = 2
                    _ = fi37
                    _ = v38
                    _ = v41
                }
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                if ("a" <= "foo bar") {
                    _ = len(v3)
                    if false {
                        v3 = append(v3, ("a" + "a"))
                        if (4294967295 < 4294967295) {
                            v1 = v1
                            var v42 int32 = 2147483647
                            _ = v42
                        } else {
                            fmt.Println(v1, v1)
                        }
                    }
                } else {
                    var v43 uint = 2
                    var v44 string = "hello"
                    switch v44 {
                    case "":
                        switch v1 {
                        case -625:
                            switch (fi31 % (((fi31 | fi31)) | 1)) {
                            case 4294967295:
                                v3 = append(v3, (v44 + v44))
                            case -241:
                                v2 = append(v2, "\n\t")
                            case 42:
                                switch -994 {
                                case -9223372036854775808:
                                    switch "" {
                                    case "\n\t":
                                        switch (-(fi31)) {
                                        case -896:
                                            for fi45 := 0; fi45 < 14; fi45++ {
                                                v43 -= (v43 | v43)
                                                v43 = (v43 - 65535)
                                                _ = fi45
                                            }
                                        case 9223372036854775807:
                                            for fi46 := 0; fi46 < 20; fi46++ {
                                                v47 := GenS7{F0: (^(v43))}
                                                v2 = append(v2, v44)
                                                if true {
                                                    if true {
                                                        v43 += v43
                                                        fmt.Println(v1, v43)
                                                        v44 = v44
                                                        switch (v44 + v44) {
                                                        case "foo bar":
                                                            if false {
                                                                v43 = 42
                                                                v44 = v44
                                                                fmt.Println(v44, fi46)
                                                            }
                                                        case "a":
                                                            v48 := GenS8{F0: -88, F1: v43, F2: 4294967295}
                                                            _ = v48
                                                        case "":
                                                            switch v44 {
                                                            case "":
                                                                var v49 int32 = -1
                                                                _ = v49
                                                            case "a":
                                                                fmt.Println(fi46)
                                                            default:
                                                                v50 := GenS8{F0: 1000, F1: (v43 + (v43 << (((v43 + (v43 | v43))) & 63))), F2: 368}
                                                                _ = v50
                                                            }
                                                        }
                                                        _ = len(v3)
                                                        v44 = v44
                                                    } else {
                                                        v51 := GenS1{F0: 0}
                                                        var v52 float32 = 0.0
                                                        if len(v2) > 0 {
                                                            v2[0] = ""
                                                        }
                                                        if true {
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                            var v53 int16 = 264
                                                            if false {
                                                                fmt.Println(v44, v1)
                                                                var v54 int64 = 42
                                                                _ = v54
                                                            } else {
                                                                v43 -= v43
                                                                switch ("\n\t" + (v44 + v44)) {
                                                                case "a":
                                                                    var v55 int = fi46
                                                                    _ = v55
                                                                case "":
                                                                    v2 = append(v2, (v44 + ("a" + v44)))
                                                                case "foo bar":
                                                                    var v56 float32 = 0.1
                                                                    _ = v56
                                                                default:
                                                                    var v57 float32 = (1e10 - (v52 * v52))
                                                                    _ = v57
                                                                }
                                                            }
                                                            _ = v53
                                                        } else {
                                                            if (v43 < 2) {
                                                                var v58 int = 9223372036854775807
                                                                var v59 uint64 = 0
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            var v60 bool = true
                                                            v60 = (((1 == 42) && true) || false)
                                                            _ = v60
                                                        }
                                                        var v61 int = 32767
                                                        _ = v51
                                                        _ = v52
                                                        _ = v61
                                                    }
                                                    var v62 uint32 = 518
                                                    _ = v62
                                                } else {
                                                    v44 = v44
                                                    v63 := GenS9{F0: ((v44 + "\n\t") + v44), F1: v44, F2: (1 >= 42)}
                                                    _ = v63
                                                }
                                                _ = fi46
                                                _ = v47
                                            }
                                        default:
                                            v64 := GenS7{F0: (v43 + (v43 * (v43 + v43)))}
                                            _ = v64
                                        }
                                    case "foo bar":
                                        var v65 string = ("hello" + "")
                                        _ = v65
                                    default:
                                        v43 += v43
                                    }
                                case 1:
                                    v43 -= (v43 << ((v43) & 63))
                                default:
                                    var v66 uint32 = 65535
                                    _ = v66
                                }
                            }
                        }
                    default:
                        for fi67 := 0; fi67 < 9; fi67++ {
                            fmt.Println(v1, fi31)
                            if (537 == 0) {
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                var v68 uint16 = 1
                                var v69 uint32 = 2147483648
                                v1 -= (v1 >> ((v43) & 63))
                                v70 := GenS10{F0: (("\n\t" + (v44 + "foo bar")) + (v44 + v44)), F1: (v69 ^ 2), F2: 100}
                                var v71 uint8 = 42
                                switch (((v44 + v44) + (v44 + (v44 + "a"))) + v44) {
                                case "\n\t":
                                    switch (v44 + ("a" + v44)) {
                                    case "hello":
                                        v2 = append(v2, ("foo bar" + "a"))
                                    default:
                                        fmt.Println(fi31, v43, v43)
                                    }
                                case "a":
                                    v68 += (v68 * 37)
                                case "foo bar":
                                    v72 := GenS2{F0: 999, F1: v44}
                                    _ = v72
                                }
                                _ = v68
                                _ = v69
                                _ = v70
                                _ = v71
                            } else {
                                _ = len(v2)
                                var v73 int32 = 0
                                _ = v73
                            }
                            _ = fi67
                        }
                    }
                    _ = v43
                    _ = v44
                }
            } else {
                _ = len(v3)
            }
            var v74 uint8 = 127
            for fi75 := 0; fi75 < 7; fi75++ {
                var v76 int8 = 42
                var v77 int32 = 1
                v1 += (v1 << ((594) & 63))
                _ = fi75
                _ = v76
                _ = v77
            }
            _ = v74
        }
        _ = fi31
    }
    v1 = (v1 >> ((1) & 63))
    var v78 string = "hello"
    var v79 int32 = -775
    if ((42 >= 18446744073709551615) && (!("foo bar" != "hello"))) {
        switch (v78 + v78) {
        case "a":
            switch v1 {
            case -2147483648:
                v80 := GenS11{F0: 42}
                _ = v80
            case 4294967295:
                v1 -= v1
            }
        case "foo bar":
            switch -442 {
            case 2147483647:
                v1 = (v1 + v1)
            case 1:
                for fi81 := 0; fi81 < 16; fi81++ {
                    if len(v3) > 0 {
                        v3[0] = v78
                    }
                    var v82 uint16 = 944
                    if (false || (-128 > -65)) {
                        v79 += v79
                        v1 += (v1 - (v1 >> ((2) & 63)))
                        for fi83 := 0; fi83 < 9; fi83++ {
                            var v84 uint64 = 42
                            var v85 int = (fi83 << ((2) & 63))
                            if true {
                                var v86 uint8 = 128
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v86
                            }
                            _ = fi83
                            _ = v84
                            _ = v85
                        }
                    }
                    _ = fi81
                    _ = v82
                }
            default:
                for fi87 := 0; fi87 < 5; fi87++ {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        for fi88 := 0; fi88 < 7; fi88++ {
                            v89 := GenS4{F0: 65535, F1: 1, F2: 42}
                            fmt.Println(v78, fi88, v78)
                            if false {
                                fmt.Println(fi88, v79)
                                fmt.Println(fi88)
                                if false {
                                    v1 += 32768
                                    v78 = v78
                                    switch (v1 & 947) {
                                    case 42:
                                        v90 := GenS12{F0: 0, F1: 2147483648, F2: ("" + (v78 + v78))}
                                        _ = v90
                                    default:
                                        v91 := GenS13{F0: -1}
                                        _ = v91
                                    }
                                } else {
                                    v92 := GenS14{F0: 393, F1: (1 == 0), F2: 1}
                                    _ = v92
                                }
                            } else {
                                v93 := GenS4{F0: 65535, F1: 42, F2: 4294967295}
                                fmt.Println(v78, v1)
                                fmt.Println(v78, v79)
                                _ = v93
                            }
                            _ = fi88
                            _ = v89
                        }
                        var v94 float32 = 1e10
                        _ = v94
                    } else {
                        v1 = fi87
                    }
                    var v95 float32 = (3.14159 - ((3.14159 + -2.25) * 1.5))
                    _ = fi87
                    _ = v95
                }
            }
        default:
            var v96 int32 = v79
            _ = v96
        }
        if ((v79 * (v79 ^ (v79 | (v79 | (^(v79)))))) >= -201) {
            _ = len(v2)
            switch (v78 + v78) {
            case "foo bar":
                v97 := GenS15{F0: -128}
                _ = v97
            case "hello":
                if true {
                    fmt.Println(v79)
                    v98 := GenS6{F0: -87, F1: 941, F2: "a"}
                    _ = v98
                } else {
                    var v99 bool = (!(("a" + v78) > ("" + v78)))
                    _ = v99
                }
            case "\n\t":
                var v100 uint = 1
                _ = v100
            }
            var v101 int8 = -1
            _ = v101
        } else {
            if (470 <= -433) {
                if (false && true) {
                    v79 -= (v79 % ((v79) | 1))
                    switch (v1 & v1) {
                    case -9223372036854775808:
                        switch v78 {
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = v78
                            }
                        }
                    case -1:
                        switch 32767 {
                        case 42:
                            var v102 uint64 = 4294967295
                            _ = v102
                        case 9223372036854775807:
                            v78 = "foo bar"
                        case 2147483647:
                            for fi103 := 0; fi103 < 20; fi103++ {
                                var v104 int8 = 127
                                _ = fi103
                                _ = v104
                            }
                        default:
                            var v105 int16 = -32768
                            _ = v105
                        }
                    case -2147483648:
                        v106 := GenS16{F0: 3, F1: 18446744073709551615, F2: v78}
                        _ = v106
                    }
                }
                v79 += (v79 * (v79 ^ (v79 / (((v79 * (^(v79)))) | 1))))
                v3 = append(v3, (v78 + v78))
            }
        }
        var v107 uint8 = 76
        _ = v107
    } else {
        v78 = v78
        switch v1 {
        case 127:
            var v108 uint64 = 4294967295
            _ = v108
        case -926:
            v79 += (v79 << ((4294967295) & 63))
        case -806:
            for fi109 := 0; fi109 < 17; fi109++ {
                for fi110 := 0; fi110 < 17; fi110++ {
                    v111 := GenS3{F0: true, F1: 65535}
                    fmt.Println(v79, v79)
                    _ = fi110
                    _ = v111
                }
                v78 = ((v78 + "foo bar") + "foo bar")
                _ = fi109
            }
        default:
            var v112 uint32 = 505
            _ = v112
        }
    }
    var v113 uint16 = 1
    v78 = v78
    for fi114 := 0; fi114 < 2; fi114++ {
        v113 = v113
        switch (v1 << ((235) & 63)) {
        case -2147483648:
            v79 += v79
        case 750:
            fmt.Println(fi114)
        case 0:
            switch ("foo bar" + "") {
            case "hello":
                fmt.Println(v79, v79)
            default:
                var v115 int = -748
                _ = v115
            }
        }
        if false {
            for fi116 := 0; fi116 < 14; fi116++ {
                var v117 uint32 = 2147483647
                var v118 float64 = 1.5
                if (-32768 > 32767) {
                    if len(v2) > 0 {
                        v2[0] = v78
                    }
                    var v119 uint8 = 5
                    _ = v119
                }
                v120 := GenS10{F0: "a", F1: 147, F2: 8}
                _ = fi116
                _ = v117
                _ = v118
                _ = v120
            }
            for fi121 := 0; fi121 < 18; fi121++ {
                v79 += v79
                _ = fi121
            }
        } else {
            if (-1 < 42) {
                var v122 uint64 = 0
                var v123 uint = uint(v113)
                for fi124 := 0; fi124 < 10; fi124++ {
                    v125 := GenS17{F0: false, F1: -1, F2: (v123 & v123)}
                    _ = fi124
                    _ = v125
                }
                _ = v122
                _ = v123
            } else {
                var v126 int64 = -29
                _ = v126
            }
            v1 -= (fi114 - (^(v1)))
        }
        _ = fi114
    }
    fmt.Println(v1, v79)
    if len(v2) > 0 {
        v2[0] = ""
    }
    var v127 int16 = -381
    var v128 int8 = 42
    v127 += 1
    var v129 bool = ((-2.25 <= (0.1 + (-2.25 - ((3.14159 + 0.1) / ((0.1 / ((1.5 / 0.1) / -2.25)) - -2.25))))) || (127 > 180))
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v113
    _ = v127
    _ = v128
    _ = v129
    _ = v2
    _ = v3
    _ = v78
    _ = v79
}
