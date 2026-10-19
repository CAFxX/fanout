package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint32; F1 int; F2 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 bool; F1 int32; F2 uint32 }
type GenS13 struct { F0 string; F1 uint64; F2 uint16 }
type GenS14 struct { F0 int16; F1 int32; F2 uint8 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint16; F1 int8; F2 uint }
func genFn0(p0 string) uint {
    switch "" {
    case "":
        for fi5 := 0; fi5 < 16; fi5++ {
            fmt.Println(p0, fi5)
            v6 := GenS1{F0: 127, F1: (fi5 & 32768), F2: 152}
            _ = fi5
            _ = v6
        }
    case "\n\t":
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(p0, fi7)
            _ = fi7
        }
    case "hello":
        var v8 int = 1
        _ = v8
    default:
        v9 := GenS1{F0: -128, F1: 9223372036854775807, F2: 0}
        _ = v9
    }
    if false {
        if true {
            var v10 int16 = -1
            fmt.Println(v10)
            _ = v10
        }
        var v11 uint16 = 2
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(fi12, p0, p0)
            _ = fi12
        }
        _ = v11
    } else {
        switch 42 {
        case 2147483647:
            var v13 uint64 = 2
            _ = v13
        case -2147483648:
            var v14 int32 = -2147483648
            _ = v14
        case 127:
            v15 := GenS1{F0: -78, F1: -1, F2: 1000}
            _ = v15
        default:
            if false {
                var v16 uint32 = 659
                for fi17 := 0; fi17 < 6; fi17++ {
                    if (-556 > int64(fi17)) {
                        v18 := GenS1{F0: 1, F1: 571, F2: 32767}
                        v19 := GenS1{F0: 11, F1: (fi17 % ((127) | 1)), F2: 1}
                        _ = v18
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            var v21 int16 = 1
                            var v22 int = fi17
                            v16 = (v16 | v16)
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        switch -1 {
                        case 4294967295:
                            var v23 int8 = 0
                            _ = v23
                        default:
                            var v24 int16 = -145
                            _ = v24
                        }
                        var v25 int16 = 0
                        _ = v25
                    }
                    switch -2147483648 {
                    case -9223372036854775808:
                        for fi26 := 0; fi26 < 6; fi26++ {
                            var v27 int64 = 1
                            var v28 int64 = 42
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    case 1:
                        p0 = ((p0 + (("\n\t" + (p0 + ("a" + "\n\t"))) + p0)) + "hello")
                    default:
                        v29 := GenS2{F0: 0}
                        _ = v29
                    }
                    if (!((p0 + p0) != ((("foo bar" + (p0 + ("hello" + p0))) + p0) + p0))) {
                        if true {
                            if true {
                                var v30 uint16 = 177
                                _ = v30
                            } else {
                                fmt.Println(v16, p0, p0)
                                var v31 int = (fi17 + 9223372036854775807)
                                _ = v31
                            }
                        }
                        for fi32 := 0; fi32 < 1; fi32++ {
                            v33 := GenS3{F0: -1, F1: 42, F2: 0}
                            var v34 int32 = 0
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        v16 -= v16
                    } else {
                        v35 := GenS3{F0: 9223372036854775807, F1: 1, F2: 9223372036854775807}
                        _ = v35
                    }
                    _ = fi17
                }
                _ = v16
            } else {
                var v36 bool = true
                for fi37 := 0; fi37 < 1; fi37++ {
                    if (float32(3.14159) == 1.5) {
                        if false {
                            p0 = p0
                            var v38 float64 = 0.0
                            _ = v38
                        } else {
                            var v39 int8 = -128
                            p0 = p0
                            v40 := GenS4{F0: int64(fi37), F1: 0, F2: 4294967295}
                            _ = v39
                            _ = v40
                        }
                        var v41 bool = true
                        p0 = p0
                        _ = v41
                    }
                    v42 := []int{fi37, 973, fi37, fi37}
                    _ = fi37
                    _ = v42
                }
                _ = v36
            }
        }
        for fi43 := 0; fi43 < 19; fi43++ {
            for fi44 := 0; fi44 < 20; fi44++ {
                v45 := GenS1{F0: 1, F1: (fi43 << ((65535) & 63)), F2: 859}
                if (3.14159 == 1.5) {
                    switch (fi44 >> ((0) & 63)) {
                    case 9223372036854775807:
                        p0 = ""
                    }
                    p0 = "hello"
                    var v46 int8 = 127
                    p0 = (p0 + "\n\t")
                    v46 = (v46 % (((^(v46))) | 1))
                    fmt.Println(fi44, p0, v46)
                    var v47 float64 = 3.14159
                    var v48 bool = true
                    v46 = 0
                    _ = v46
                    _ = v47
                    _ = v48
                } else {
                    var v49 int16 = -32768
                    if (-1 == 127) {
                        fmt.Println(fi44)
                        var v50 int = (fi43 & (-(fi44)))
                        v51 := GenS5{F0: 2, F1: -94, F2: 42}
                        var v52 uint8 = 0
                        if false {
                            switch fi43 {
                            case 128:
                                if (!true) {
                                    switch 1 {
                                    case -9223372036854775808:
                                        if (100 != 100) {
                                            var v53 int16 = v49
                                            v49 = 42
                                            _ = v53
                                        } else {
                                            if false {
                                                v54 := []int{(fi44 & (fi43 % (((v50 >> ((0) & 63))) | 1))), -635}
                                                _ = len(v54)
                                                _ = v54
                                            }
                                        }
                                    }
                                    fmt.Println(v50, v52, fi43)
                                } else {
                                    v55 := []int{v50, fi43}
                                    v55 = append(v55, 2147483647)
                                    v55 = append(v55, (-(v50)))
                                    switch v50 {
                                    case 766:
                                        if false {
                                            var v56 uint64 = 4294967295
                                            fmt.Println(v52)
                                            var v57 int32 = 1
                                            _ = v56
                                            _ = v57
                                        }
                                    default:
                                        var v58 uint8 = 42
                                        _ = v58
                                    }
                                    _ = v55
                                }
                                switch -281 {
                                case 59:
                                    var v59 int8 = 127
                                    _ = v59
                                default:
                                    switch fi44 {
                                    case 127:
                                        p0 = "\n\t"
                                    case 76:
                                        v60 := []int{(fi44 - (fi44 + fi43)), fi44}
                                        if (true || (-1 >= 100)) {
                                            v61 := GenS6{F0: (-(v49))}
                                            _ = v61
                                        }
                                        if (-1 > -1) {
                                            v62 := GenS7{F0: 511, F1: 234, F2: 255}
                                            if (0.1 == 3.14159) {
                                                v49 = (v49 << ((18446744073709551615) & 63))
                                            }
                                            switch p0 {
                                            case "":
                                                if false {
                                                    if true {
                                                        var v63 int64 = 207
                                                        v64 := GenS5{F0: 2, F1: 42, F2: 42}
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                }
                                            default:
                                                if (0 >= -1) {
                                                    var v65 int16 = v49
                                                    switch ((((p0 + ("" + "a")) + "foo bar") + "a") + p0) {
                                                    case "a":
                                                        v52 -= 237
                                                    case "hello":
                                                        var v66 uint64 = 854
                                                        _ = v66
                                                    case "":
                                                        p0 = (("" + p0) + p0)
                                                        var v67 int = v50
                                                        if true {
                                                            p0 = (p0 + p0)
                                                            if true {
                                                                _ = len(v60)
                                                            } else {
                                                                var v68 uint = 331
                                                                v67 -= (-(v67))
                                                                v69 := GenS2{F0: -380}
                                                                var v70 float64 = 0.1
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            p0 = p0
                                                        } else {
                                                            var v71 uint16 = 32767
                                                            _ = len(v60)
                                                            v60 = append(v60, v50)
                                                            if (!(float32((1e10 * 0.0)) >= ((1.5 * (0.1 - 1.5)) * (0.0 - (0.0 - 0.0))))) {
                                                                if len(v60) >= 2 {
                                                                    v60 = v60[1:len(v60)]
                                                                }
                                                                v71 -= (v71 + v71)
                                                            } else {
                                                                v72 := GenS7{F0: 2, F1: 263, F2: (v71 | v71)}
                                                                _ = v72
                                                            }
                                                            _ = v71
                                                        }
                                                        v60 = append(v60, 2147483647)
                                                        _ = v67
                                                    }
                                                    _ = v65
                                                }
                                                v49 = (v49 % (((v49 / (((v49 - v49)) | 1))) | 1))
                                            }
                                            _ = v62
                                        } else {
                                            v50 -= 795
                                            if true {
                                                var v73 float64 = ((0.1 * -2.25) / ((0.1 * 0.1) + -2.25))
                                                if (!(546 <= 0)) {
                                                    v50 = v50
                                                } else {
                                                    var v74 uint = 18446744073709551615
                                                    _ = v74
                                                }
                                                var v75 int16 = (-(v49))
                                                _ = v73
                                                _ = v75
                                            } else {
                                                var v76 int64 = -1
                                                v52 = (v52 & v52)
                                                switch (p0 + ("foo bar" + p0)) {
                                                case "":
                                                    switch (((p0 + (p0 + "hello")) + (p0 + p0)) + p0) {
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "foo bar":
                                                            p0 = ((p0 + p0) + p0)
                                                        }
                                                    case "hello":
                                                        fmt.Println(fi43)
                                                    case "a":
                                                        v50 = v50
                                                    default:
                                                        var v77 uint = 2
                                                        _ = v77
                                                    }
                                                }
                                                switch -2147483648 {
                                                case 1:
                                                    v49 = -598
                                                default:
                                                    if false {
                                                        p0 = (p0 + p0)
                                                        v52 = (v52 | 42)
                                                        v76 += 9223372036854775807
                                                    } else {
                                                        if true {
                                                            fmt.Println(p0, p0)
                                                            v50 = (fi43 / ((v50) | 1))
                                                            switch v50 {
                                                            case 32767:
                                                                if (v76 >= (-(v76))) {
                                                                    switch p0 {
                                                                    case "foo bar":
                                                                        v78 := GenS2{F0: -263}
                                                                        _ = v78
                                                                    case "a":
                                                                        var v79 bool = true
                                                                        _ = v79
                                                                    case "":
                                                                        if false {
                                                                            var v80 uint = 252
                                                                            var v81 uint32 = 2
                                                                            v82 := GenS8{F0: v81, F1: 0, F2: 42}
                                                                            switch v50 {
                                                                            case 2147483647:
                                                                                var v83 int = v50
                                                                                if true {
                                                                                    v50 = v50
                                                                                } else {
                                                                                    switch p0 {
                                                                                    case "":
                                                                                        v52 = (v52 & (v52 >> ((255) & 63)))
                                                                                    }
                                                                                    v80 = 87
                                                                                }
                                                                                v84 := GenS5{F0: 846, F1: 127, F2: 255}
                                                                                _ = v83
                                                                                _ = v84
                                                                            }
                                                                            _ = v80
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            if true {
                                                                                var v85 float64 = -2.25
                                                                                var v86 int64 = (v76 ^ (v76 - (v76 + -1)))
                                                                                var v87 int8 = 110
                                                                                _ = v85
                                                                                _ = v86
                                                                                _ = v87
                                                                            }
                                                                        }
                                                                    }
                                                                    if ((!(false || true)) || false) {
                                                                        switch (-(fi44)) {
                                                                        case 2147483647:
                                                                            if (!((v49 & v49) >= v49)) {
                                                                                switch p0 {
                                                                                case "a":
                                                                                    var v88 int8 = int8(v52)
                                                                                    _ = v88
                                                                                default:
                                                                                    var v89 bool = (!((0.1 * ((((((0.0 - ((((((-2.25 - (1.5 - (0.1 / ((((1e10 / 0.1) / -2.25) - float32(-2.25)) - float32(-2.25))))) + (0.1 - 0.1)) - 0.1) - 1.5) / (1e10 / 3.14159)) / 1.0)) * -2.25) * 3.14159) - 3.14159) / float32(3.14159)) - 0.1)) >= (-2.25 + 1e10)))
                                                                                    _ = v89
                                                                                }
                                                                                fmt.Println(fi44, v52)
                                                                                var v90 int16 = (-(v49))
                                                                                _ = v90
                                                                            }
                                                                        case 32768:
                                                                            var v91 bool = true
                                                                            _ = v91
                                                                        default:
                                                                            if false {
                                                                                v49 = v49
                                                                                if ((v49 % ((-32768) | 1)) < v49) {
                                                                                    _ = len(v60)
                                                                                }
                                                                                v92 := GenS5{F0: 65535, F1: 42, F2: 32767}
                                                                                var v93 int16 = v49
                                                                                _ = v92
                                                                                _ = v93
                                                                            } else {
                                                                                _ = len(v60)
                                                                                if len(v60) > 0 {
                                                                                    v60[0] = 2147483647
                                                                                }
                                                                                if true {
                                                                                    if ((255 == 0) || true) {
                                                                                        switch p0 {
                                                                                        case "foo bar":
                                                                                            v76 += 9223372036854775807
                                                                                        case "hello":
                                                                                            var v94 uint64 = 2
                                                                                            _ = v94
                                                                                        case "\n\t":
                                                                                            _ = len(v60)
                                                                                        default:
                                                                                            v49 += 1
                                                                                        }
                                                                                    } else {
                                                                                        var v95 uint32 = 2
                                                                                        switch (fi43 << ((65535) & 63)) {
                                                                                        case 128:
                                                                                            if len(v60) > 0 {
                                                                                                v60[0] = 9223372036854775807
                                                                                            }
                                                                                        case -571:
                                                                                            if (42 == 42) {
                                                                                                var v96 int16 = v49
                                                                                                v60 = append(v60, fi43)
                                                                                                _ = v96
                                                                                            } else {
                                                                                                fmt.Println(fi44, v52)
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            v97 := GenS8{F0: 1, F1: (^(fi44)), F2: (v76 & v76)}
                                                                                            _ = v97
                                                                                        default:
                                                                                            switch (fi44 % (((v50 + (fi44 + 499))) | 1)) {
                                                                                            case 127:
                                                                                                switch ("\n\t" + (p0 + ("a" + p0))) {
                                                                                                case "a":
                                                                                                    var v98 int16 = (v49 >> ((4294967295) & 63))
                                                                                                    _ = v98
                                                                                                default:
                                                                                                    if len(v60) >= 2 {
                                                                                                        v60 = v60[1:len(v60)]
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v76, v52)
                                                                                            }
                                                                                        }
                                                                                        var v99 uint16 = 0
                                                                                        _ = v95
                                                                                        _ = v99
                                                                                    }
                                                                                    var v100 string = ("\n\t" + p0)
                                                                                    var v101 uint16 = 42
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                } else {
                                                                                    if len(v60) >= 2 {
                                                                                        v60 = v60[1:len(v60)]
                                                                                    }
                                                                                    if (-2.25 >= 1e10) {
                                                                                        if false {
                                                                                            if true {
                                                                                                switch p0 {
                                                                                                case "":
                                                                                                    fmt.Println(p0, v52, v50)
                                                                                                case "a":
                                                                                                    switch -351 {
                                                                                                    case 67:
                                                                                                        switch ((((p0 + (p0 + ("" + (p0 + p0)))) + p0) + ((p0 + ("\n\t" + ((p0 + "a") + (p0 + ("a" + ((p0 + "\n\t") + (p0 + ("hello" + p0)))))))) + p0)) + (p0 + (p0 + p0))) {
                                                                                                        case "hello":
                                                                                                            v76 = (v76 >> ((584) & 63))
                                                                                                        case "foo bar":
                                                                                                            v76 = 42
                                                                                                        default:
                                                                                                            v49 = (v49 - v49)
                                                                                                        }
                                                                                                    case 42:
                                                                                                        var v102 bool = true
                                                                                                        _ = v102
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    if true {
                                                                                                        var v103 string = p0
                                                                                                        switch v50 {
                                                                                                        case -2147483648:
                                                                                                            p0 = ((v103 + "hello") + p0)
                                                                                                        default:
                                                                                                            var v104 int8 = 0
                                                                                                            _ = v104
                                                                                                        }
                                                                                                        _ = v103
                                                                                                    }
                                                                                                default:
                                                                                                    var v105 uint16 = 1
                                                                                                    _ = v105
                                                                                                }
                                                                                            } else {
                                                                                                switch fi44 {
                                                                                                case 178:
                                                                                                    v76 = int64(v49)
                                                                                                case -305:
                                                                                                    var v106 float64 = 0.0
                                                                                                    _ = v106
                                                                                                case 127:
                                                                                                    switch (fi44 >> ((232) & 63)) {
                                                                                                    case 0:
                                                                                                        v107 := GenS9{F0: (p0 < (p0 + p0))}
                                                                                                        _ = v107
                                                                                                    default:
                                                                                                        var v108 int64 = (-(v76))
                                                                                                        _ = v108
                                                                                                    }
                                                                                                }
                                                                                                v109 := GenS2{F0: 1}
                                                                                                fmt.Println(v49, v76, v76)
                                                                                                fmt.Println(p0)
                                                                                                fmt.Println(v50, v50)
                                                                                                if (true || (893 <= v76)) {
                                                                                                    fmt.Println(v49, v49, p0)
                                                                                                }
                                                                                                _ = v109
                                                                                            }
                                                                                            v110 := GenS8{F0: 1, F1: (fi44 ^ (fi43 % ((fi44) | 1))), F2: (v76 % ((v76) | 1))}
                                                                                            _ = len(v60)
                                                                                            _ = v110
                                                                                        } else {
                                                                                            v76 = v76
                                                                                            v111 := GenS8{F0: 2147483647, F1: 128, F2: (v76 + (v76 | 9223372036854775807))}
                                                                                            var v112 int32 = 1
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        }
                                                                                        var v113 float32 = ((0.0 + 3.14159) + (3.14159 - 3.14159))
                                                                                        if false {
                                                                                            var v114 string = p0
                                                                                            var v115 int16 = v49
                                                                                            v50 = -2147483648
                                                                                            switch 127 {
                                                                                            case 1:
                                                                                                var v116 uint = 1
                                                                                                _ = v116
                                                                                            case 253:
                                                                                                if false {
                                                                                                    if false {
                                                                                                        fmt.Println(fi44)
                                                                                                        fmt.Println(v76, p0)
                                                                                                        if (v76 == (v76 | 9223372036854775807)) {
                                                                                                            fmt.Println(v76, v50)
                                                                                                            var v117 int64 = v76
                                                                                                            _ = v117
                                                                                                        } else {
                                                                                                            switch ("foo bar" + "foo bar") {
                                                                                                            case "\n\t":
                                                                                                                var v118 bool = false
                                                                                                                _ = v118
                                                                                                            case "hello":
                                                                                                                var v119 uint32 = 1
                                                                                                                _ = v119
                                                                                                            default:
                                                                                                                var v120 int64 = v76
                                                                                                                fmt.Println(v120)
                                                                                                                _ = v120
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v121 := GenS1{F0: 0, F1: (fi44 * v50), F2: (v49 / (((v49 & int16(v52))) | 1))}
                                                                                                    switch p0 {
                                                                                                    case "":
                                                                                                        var v122 int16 = v49
                                                                                                        _ = v122
                                                                                                    }
                                                                                                    _ = v121
                                                                                                } else {
                                                                                                    var v123 uint8 = 221
                                                                                                    v123 = v123
                                                                                                    var v124 int8 = -1
                                                                                                    _ = v123
                                                                                                    _ = v124
                                                                                                }
                                                                                            default:
                                                                                                v52 += (v52 >> ((875) & 63))
                                                                                            }
                                                                                            _ = v114
                                                                                            _ = v115
                                                                                        } else {
                                                                                            v125 := GenS10{F0: 27, F1: 185}
                                                                                            _ = v125
                                                                                        }
                                                                                        _ = v113
                                                                                    }
                                                                                }
                                                                                if (1 >= v76) {
                                                                                    var v126 uint64 = 4294967295
                                                                                    _ = v126
                                                                                }
                                                                            }
                                                                        }
                                                                        if false {
                                                                            v52 -= 255
                                                                            fmt.Println(v50, v52)
                                                                            if len(v60) > 0 {
                                                                                v60[0] = (v50 | -520)
                                                                            }
                                                                        }
                                                                        if (-10 < 45) {
                                                                            v127 := GenS11{F0: 255, F1: (v52 & v52)}
                                                                            _ = v127
                                                                        } else {
                                                                            v50 += -2147483648
                                                                            if len(v60) >= 2 {
                                                                                v60 = v60[1:len(v60)]
                                                                            }
                                                                            if (2 >= 42) {
                                                                                var v128 uint = 309
                                                                                fmt.Println(v52, v49, v52)
                                                                                v50 = int(v52)
                                                                                _ = v128
                                                                            } else {
                                                                                v50 = (v50 & fi44)
                                                                            }
                                                                            v50 += (fi44 | (v50 >> ((90) & 63)))
                                                                            var v129 int16 = v49
                                                                            _ = v129
                                                                        }
                                                                    }
                                                                    if (65535 <= uint(v52)) {
                                                                        var v130 int = (^(fi44))
                                                                        if len(v60) > 0 {
                                                                            v60[0] = (v50 ^ (fi44 ^ fi44))
                                                                        }
                                                                        _ = v130
                                                                    } else {
                                                                        if false {
                                                                            var v131 int = (fi43 | (fi43 - (fi43 / (((v50 - 32767)) | 1))))
                                                                            var v132 string = ((("a" + "foo bar") + (p0 + "foo bar")) + p0)
                                                                            _ = v131
                                                                            _ = v132
                                                                        }
                                                                    }
                                                                    v76 = v76
                                                                    if ((2 <= 326) || false) {
                                                                        v49 += v49
                                                                    } else {
                                                                        v49 = v49
                                                                    }
                                                                    if true {
                                                                        v60 = append(v60, fi43)
                                                                    } else {
                                                                        p0 = (p0 + ((p0 + "hello") + "hello"))
                                                                        v50 -= fi43
                                                                        var v133 int64 = 42
                                                                        _ = v133
                                                                    }
                                                                }
                                                            case 628:
                                                                if (true && true) {
                                                                    if (2147483647 == 0) {
                                                                        v134 := GenS11{F0: 0, F1: 2}
                                                                        _ = v134
                                                                    } else {
                                                                        v135 := GenS10{F0: 18, F1: 0}
                                                                        _ = v135
                                                                    }
                                                                }
                                                            default:
                                                                _ = len(v60)
                                                            }
                                                        }
                                                        if ((false && (!((v52 | 170) > 42))) && false) {
                                                            if (false && false) {
                                                                v60 = append(v60, -1)
                                                                v76 = (^(v76))
                                                                if len(v60) > 0 {
                                                                    v60[0] = 4294967295
                                                                }
                                                            } else {
                                                                v60 = append(v60, (v50 >> ((18446744073709551615) & 63)))
                                                                if false {
                                                                    var v136 uint32 = 4294967295
                                                                    v136 = 2147483648
                                                                    v60 = append(v60, -780)
                                                                    _ = v136
                                                                } else {
                                                                    var v137 float64 = 1e10
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        v138 := GenS12{F0: (((-(v76)) > (v76 >> ((255) & 63))) || (!(v76 > v76))), F1: 17, F2: 763}
                                                                        _ = v138
                                                                    case "hello":
                                                                        if len(v60) >= 2 {
                                                                            v60 = v60[1:len(v60)]
                                                                        }
                                                                    case "foo bar":
                                                                        p0 = (p0 + ("a" + ("hello" + "")))
                                                                    }
                                                                    v50 += v50
                                                                    _ = v137
                                                                }
                                                                var v139 int32 = 2147483647
                                                                fmt.Println(fi43, fi43, v50)
                                                                if (-107 == 1) {
                                                                    fmt.Println(fi44, v76)
                                                                    v140 := GenS13{F0: "a", F1: 0, F2: 42}
                                                                    v52 += (^(v52))
                                                                    _ = v140
                                                                } else {
                                                                    v76 = v76
                                                                }
                                                                _ = v139
                                                            }
                                                            v50 += 0
                                                            v60 = append(v60, 127)
                                                        }
                                                        v76 += (v76 * 9223372036854775807)
                                                    }
                                                }
                                                _ = v76
                                            }
                                            var v141 float32 = -2.25
                                            _ = v141
                                        }
                                        switch fi43 {
                                        case 0:
                                            if len(v60) > 0 {
                                                v60[0] = 4294967295
                                            }
                                        default:
                                            var v142 int32 = -212
                                            _ = v142
                                        }
                                        var v143 int64 = -9223372036854775808
                                        _ = v143
                                        _ = v60
                                    case 42:
                                        switch 32768 {
                                        case -2147483648:
                                            var v144 int8 = 0
                                            _ = v144
                                        case -9223372036854775808:
                                            var v145 int64 = 9223372036854775807
                                            _ = v145
                                        case -185:
                                            v146 := []int{fi44, -9223372036854775808, 127}
                                            _ = v146
                                        }
                                        v147 := []int{v50}
                                        _ = v147
                                    }
                                }
                            }
                            var v148 float32 = -2.25
                            v50 -= fi43
                            _ = v148
                        }
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    v149 := GenS3{F0: 1, F1: 127, F2: 808}
                    _ = v149
                    _ = v49
                }
                var v150 int8 = 42
                if (!(!(!true))) {
                    switch p0 {
                    case "hello":
                        v151 := []int{0}
                        _ = v151
                    default:
                        v150 += v150
                        v152 := []int{fi43, 4294967295, 2147483647, (fi43 | (fi44 | (-(fi43))))}
                        _ = v152
                    }
                    var v153 string = p0
                    v154 := GenS1{F0: v150, F1: (fi44 + fi43), F2: 0}
                    _ = v153
                    _ = v154
                }
                if (4294967295 > 2) {
                    var v155 uint = 2
                    if false {
                        v150 = v150
                        switch fi43 {
                        case 42:
                            var v156 int16 = -1
                            _ = v156
                        default:
                            switch (p0 + (p0 + "")) {
                            case "a":
                                var v157 int8 = v150
                                _ = v157
                            case "hello":
                                v158 := []int{}
                                _ = v158
                            case "foo bar":
                                v159 := []int{fi44, 2147483647}
                                _ = v159
                            }
                        }
                        p0 = ("" + p0)
                    } else {
                        switch p0 {
                        case "a":
                            v155 = (v155 % ((v155) | 1))
                        default:
                            var v160 float64 = 1.5
                            _ = v160
                        }
                        if true {
                            v161 := GenS9{F0: ((!false) && true)}
                            var v162 uint64 = 0
                            _ = v161
                            _ = v162
                        }
                        var v163 int8 = v150
                        _ = v163
                    }
                    _ = v155
                }
                _ = fi44
                _ = v150
                _ = v45
            }
            v164 := GenS11{F0: 0, F1: 135}
            _ = fi43
            _ = v164
        }
        v165 := []int{1, 127, 120, 554}
        _ = v165
    }
    v166 := []int{32768, -2147483648, 0, -9223372036854775808}
    _ = len(v166)
    _ = p0
    _ = v166
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 330
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    v2 = (v2 | (v2 ^ (^(v2))))
    for fi167 := 0; fi167 < 20; fi167++ {
        v2 = 875
        _ = fi167
    }
    v168 := GenS14{F0: 0, F1: -1, F2: 255}
    switch "hello" {
    case "\n\t":
        fmt.Println(v2)
    default:
        v169 := GenS15{F0: -1}
        _ = v169
    }
    fmt.Println(v2)
    v2 = (v2 % (((v2 % ((1) | 1))) | 1))
    v2 = (v2 % ((v2) | 1))
    var v170 uint8 = 2
    fmt.Println(v170)
    _ = len(v4)
    if true {
        for fi171 := 0; fi171 < 2; fi171++ {
            var v172 float32 = float32((0.0 * 1.5))
            v3[42] = (!(!(735 <= 0)))
            _ = fi171
            _ = v172
        }
        _ = len(v4)
    }
    v2 += -2147483648
    v173 := GenS16{F0: 818, F1: -1, F2: 2}
    v174 := GenS10{F0: -128, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v170:", v170)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v168
    _ = v170
    _ = v173
    _ = v174
    _ = v2
    _ = v3
    _ = v4
}
