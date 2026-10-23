package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 int64; F2 bool }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 uint; F1 uint; F2 string }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 int8; F1 int32; F2 int }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS14 struct { F0 int32; F1 uint16; F2 bool }
type GenS15 struct { F0 int16; F1 int64 }
type GenS16 struct { F0 uint16; F1 int64 }
type GenS17 struct { F0 uint32; F1 int64 }
type GenS18 struct { F0 int; F1 uint32 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 uint16; F1 uint8 }
type GenS21 struct { F0 int64; F1 int8; F2 uint64 }
type GenS22 struct { F0 uint16; F1 bool; F2 int }
type GenS23 struct { F0 int64; F1 int }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 int32; F1 uint64; F2 int }
func genFn0(p0 bool, p1 bool, p2 uint8) uint64 {
    v7 := []int{128, 32768}
    if (p1 && (0 < 32767)) {
        v7 = append(v7, 42)
        switch 128 {
        case 713:
            for fi8 := 0; fi8 < 12; fi8++ {
                switch "hello" {
                case "\n\t":
                    var v9 uint64 = 2
                    _ = v9
                }
                var v10 uint8 = p2
                for fi11 := 0; fi11 < 8; fi11++ {
                    if len(v7) > 0 {
                        v7[0] = (fi11 / (((-(fi8))) | 1))
                    }
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        p0 = p0
                    }
                    if false {
                        if (p0 || p0) {
                            v12 := GenS1{F0: false, F1: 2, F2: 42}
                            _ = v12
                        } else {
                            var v13 float64 = float64(1e10)
                            _ = v13
                        }
                        if (!true) {
                            if p0 {
                                p1 = (!(-1 < 1))
                                p1 = p1
                            }
                            v14 := GenS2{F0: 42, F1: -1, F2: 2}
                            v15 := GenS2{F0: 65535, F1: 0, F2: p2}
                            var v16 uint16 = 42
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            p2 -= (p2 * v10)
                            var v17 string = ""
                            _ = v17
                        }
                        switch "" {
                        case "hello":
                            p0 = p0
                        case "foo bar":
                            var v18 int64 = -1
                            _ = v18
                        case "\n\t":
                            switch ("hello" + "") {
                            case "hello":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            case "foo bar":
                                p2 = 254
                            case "a":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                        }
                    } else {
                        if p1 {
                            p0 = p0
                        }
                        var v19 float32 = (0.0 * (-2.25 * 3.14159))
                        v19 -= -2.25
                        _ = v19
                    }
                    var v20 uint8 = p2
                    _ = fi11
                    _ = v20
                }
                _ = fi8
                _ = v10
            }
        default:
            if p1 {
                if (498 != 42) {
                    v21 := GenS3{F0: 1, F1: 2}
                    _ = len(v7)
                    p0 = (!(!(!p0)))
                    _ = v21
                }
                for fi22 := 0; fi22 < 8; fi22++ {
                    p1 = false
                    _ = fi22
                }
            } else {
                if (1 <= 1) {
                    fmt.Println(p0, p1)
                    p2 = (p2 << ((464) & 63))
                } else {
                    if false {
                        p2 = p2
                        for fi23 := 0; fi23 < 2; fi23++ {
                            if (!p1) {
                                var v24 int8 = -1
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    v26 := GenS4{F0: -1}
                                    _ = len(v7)
                                    switch "" {
                                    case "a":
                                        v7 = append(v7, fi25)
                                    case "\n\t":
                                        v7 = append(v7, 9223372036854775807)
                                    case "foo bar":
                                        v24 = v24
                                    default:
                                        var v27 float32 = (float32(((1e10 / float64(1e10)) + -2.25)) - -2.25)
                                        _ = v27
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v24
                            }
                            _ = fi23
                        }
                        var v28 uint64 = 0
                        _ = v28
                    } else {
                        v29 := GenS1{F0: (110 <= 617), F1: 1, F2: -1}
                        var v30 int = -678
                        _ = v29
                        _ = v30
                    }
                    p1 = (-4 <= 1)
                }
                var v31 int = 410
                var v32 int16 = 32767
                _ = v31
                _ = v32
            }
        }
    } else {
        var v33 int64 = 42
        _ = v33
    }
    if ((1e10 + 0.0) >= (-2.25 + 0.1)) {
        var v34 uint8 = (p2 | (p2 / ((p2) | 1)))
        _ = v34
    } else {
        v35 := GenS1{F0: true, F1: 42, F2: 0}
        _ = v35
    }
    v36 := GenS5{F0: 32767, F1: 936}
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v7
    return 310
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = 0.0
    var v3 uint = 4294967295
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    var v6 uint = v3
    v1 += v1
    v37 := GenS6{F0: 168}
    v38 := GenS6{F0: -540}
    v39 := GenS3{F0: 127, F1: 0}
    var v40 int32 = 2147483647
    v41 := GenS7{F0: -1, F1: 506, F2: (65535 > 1)}
    v42 := GenS8{F0: -128, F1: 32767}
    v43 := []int{}
    if true {
        v3 = 0
        v44 := GenS2{F0: 318, F1: (^(v40)), F2: v1}
        _ = v44
    } else {
        v45 := GenS5{F0: -32768, F1: 311}
        _ = v45
    }
    if true {
        var v46 int64 = 9223372036854775807
        v1 -= (^(v1))
        _ = v46
    } else {
        switch "foo bar" {
        case "":
            if (v2 < v2) {
                var v47 int8 = 127
                v40 = v40
                _ = v47
            } else {
                var v48 uint8 = v1
                var v49 float32 = v2
                _ = v48
                _ = v49
            }
        default:
            v2 = 0.1
        }
        var v50 int8 = 0
        v51 := GenS3{F0: 1, F1: 918}
        _ = v50
        _ = v51
    }
    var v52 int = -2147483648
    if false {
        v53 := GenS9{F0: (v6 << ((v3) & 63)), F1: (v6 | v3), F2: "hello"}
        fmt.Println(v2, v3)
        v54 := GenS3{F0: 42, F1: 42}
        _ = v53
        _ = v54
    } else {
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
        var v55 string = "\n\t"
        _ = v55
    }
    v1 = v1
    if true {
        _ = len(v43)
        switch ((("a" + ((("hello" + "\n\t") + ("\n\t" + "hello")) + ("foo bar" + "a"))) + "foo bar") + ("a" + (("" + "") + (("" + ("\n\t" + ("a" + (("" + "hello") + "\n\t")))) + "")))) {
        case "":
            v6 = (v3 | v6)
        case "\n\t":
            v40 = (v40 & v40)
        }
        if (100 > 114) {
            switch "a" {
            case "":
                if false {
                    v43 = append(v43, v52)
                    for fi56 := 0; fi56 < 16; fi56++ {
                        var v57 string = "hello"
                        switch (v57 + v57) {
                        case "foo bar":
                            for fi58 := 0; fi58 < 7; fi58++ {
                                v59 := GenS2{F0: 128, F1: -2147483648, F2: (v1 + 2)}
                                var v60 int = (fi58 & fi56)
                                v61 := v5[(v60 / ((v52) | 1))]
                                if (!(1 == 1)) {
                                    var v62 float64 = (0.1 - (-2.25 * -2.25))
                                    switch v57 {
                                    case "foo bar":
                                        v63 := GenS4{F0: 393}
                                        _ = v63
                                    case "":
                                        if v61 {
                                            var v64 bool = (!(!v61))
                                            delete(v5, (-(fi56)))
                                            var v65 float64 = 1e10
                                            _ = v64
                                            _ = v65
                                        }
                                    case "a":
                                        var v66 int16 = -1
                                        _ = v66
                                    }
                                    _ = v62
                                } else {
                                    v67 := v4[1]
                                    if ((v52 - fi56) < fi58) {
                                        if (false || true) {
                                            switch (fi58 - (fi56 * 881)) {
                                            case -2147483648:
                                                var v68 uint16 = 0
                                                _ = v68
                                            default:
                                                var v69 string = ((v57 + ("\n\t" + "")) + "a")
                                                _ = v69
                                            }
                                            fmt.Println(v3, v67, fi58)
                                        } else {
                                            if (!(v67 && v67)) {
                                                if len(v43) >= 2 {
                                                    v43 = v43[1:len(v43)]
                                                }
                                                switch v60 {
                                                case 2147483647:
                                                    var v70 int64 = int64(v40)
                                                    _ = v70
                                                case 9223372036854775807:
                                                    v6 -= v3
                                                case 259:
                                                    if (!false) {
                                                        var v71 int16 = 0
                                                        v72 := GenS2{F0: 2, F1: (v40 & v40), F2: 128}
                                                        v73 := GenS8{F0: 1, F1: uint16(v6)}
                                                        v74 := GenS9{F0: v6, F1: v3, F2: (v57 + v57)}
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        v6 = v6
                                                        v75 := GenS8{F0: 1, F1: 491}
                                                        v3 = (^(v3))
                                                        if (!v61) {
                                                            if true {
                                                                v52 -= v60
                                                            }
                                                        } else {
                                                            v76 := GenS6{F0: v52}
                                                            v52 += (fi58 % (((fi58 & fi58)) | 1))
                                                            v6 += (v3 / ((v6) | 1))
                                                            _ = v76
                                                        }
                                                        if (2147483647 > 710) {
                                                            if false {
                                                                switch v57 {
                                                                case "a":
                                                                    delete(v5, v52)
                                                                }
                                                                if false {
                                                                    var v77 int32 = (v40 ^ (v40 ^ v40))
                                                                    var v78 int32 = v40
                                                                    if (-257 != -1) {
                                                                        if (-128 > 0) {
                                                                            var v79 int = (fi58 % ((v52) | 1))
                                                                            v80 := GenS10{F0: (fi58 | -2147483648), F1: 42}
                                                                            v77 = 0
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        if (v57 <= (v57 + v57)) {
                                                                            v81 := GenS1{F0: false, F1: 849, F2: 42}
                                                                            v60 -= (fi58 & fi56)
                                                                            v82 := GenS11{F0: 127, F1: v40, F2: (-(v60))}
                                                                            if v67 {
                                                                                v40 = v78
                                                                                v77 -= (v78 & (v78 - (v40 - 1)))
                                                                                if false {
                                                                                    switch fi58 {
                                                                                    case 1:
                                                                                        var v83 int16 = -32768
                                                                                        _ = v83
                                                                                    }
                                                                                    v40 = v77
                                                                                    switch v57 {
                                                                                    case "a":
                                                                                        switch v57 {
                                                                                        case "foo bar":
                                                                                            v84 := GenS12{F0: v78, F1: v1}
                                                                                            _ = v84
                                                                                        case "":
                                                                                            v85 := v5[fi56]
                                                                                            _ = v85
                                                                                        case "hello":
                                                                                            v6 = v6
                                                                                        default:
                                                                                            delete(v5, fi56)
                                                                                        }
                                                                                    default:
                                                                                        if v67 {
                                                                                            var v86 int16 = 32767
                                                                                            var v87 uint = 255
                                                                                            v88 := GenS13{F0: (^(v40)), F1: 2147483647, F2: 0}
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                            _ = v88
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v89 := GenS14{F0: (v77 + v40), F1: 954, F2: v61}
                                                                                    v90 := GenS15{F0: int16(fi58), F1: 1}
                                                                                    v91 := GenS16{F0: 1, F1: 937}
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                }
                                                                            }
                                                                            v40 = v78
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            v92 := GenS14{F0: v77, F1: 255, F2: (42 > 1)}
                                                                            fmt.Println(v77)
                                                                            v77 += 42
                                                                            _ = v92
                                                                        }
                                                                        var v93 uint16 = 255
                                                                        var v94 float64 = (((0.1 + ((1.5 * 1.5) / 1.5)) + float64(0.0)) / (0.1 / 1e10))
                                                                        var v95 int64 = 675
                                                                        _ = v93
                                                                        _ = v94
                                                                        _ = v95
                                                                    } else {
                                                                        var v96 int = (fi56 - -214)
                                                                        v5[(fi58 | (v60 + v52))] = true
                                                                        fmt.Println(v67, v77)
                                                                        _ = v96
                                                                    }
                                                                    _ = v77
                                                                    _ = v78
                                                                } else {
                                                                    v97 := v5[v52]
                                                                    var v98 bool = v61
                                                                    v3 = (^(v3))
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                            }
                                                            fmt.Println(v1, v67, v60)
                                                            if true {
                                                                switch ((((v57 + v57) + "foo bar") + "foo bar") + "\n\t") {
                                                                case "":
                                                                    v3 = (^(v3))
                                                                case "foo bar":
                                                                    v99 := GenS6{F0: (^(v52))}
                                                                    _ = v99
                                                                case "hello":
                                                                    switch v52 {
                                                                    case 42:
                                                                        switch ("\n\t" + v57) {
                                                                        case "hello":
                                                                            if true {
                                                                                v61 = (v61 || (0.1 > 0.1))
                                                                                var v100 float32 = 1.5
                                                                                fmt.Println(v100, v2, fi56)
                                                                                _ = v100
                                                                            } else {
                                                                                var v101 int64 = -601
                                                                                v57 = v57
                                                                                _ = v101
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (2 > 991) {
                                                                            switch v57 {
                                                                            case "":
                                                                                v67 = v61
                                                                            case "a":
                                                                                v6 -= v6
                                                                            case "hello":
                                                                                var v102 int16 = 42
                                                                                _ = v102
                                                                            }
                                                                            delete(v5, fi56)
                                                                        }
                                                                    }
                                                                default:
                                                                    switch 42 {
                                                                    case 1:
                                                                        if true {
                                                                            if v67 {
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v52 = v52
                                                                                case "hello":
                                                                                    if (!((!true) && v61)) {
                                                                                        v2 = (1.5 / float32(float64(3.14159)))
                                                                                        var v103 int = (fi58 & fi56)
                                                                                        _ = v103
                                                                                    } else {
                                                                                        if (!v61) {
                                                                                            var v104 string = (("" + ("hello" + (v57 + v57))) + v57)
                                                                                            _ = v104
                                                                                        }
                                                                                        v105 := GenS17{F0: 1, F1: -9223372036854775808}
                                                                                        v40 = (-(v40))
                                                                                        _ = v105
                                                                                    }
                                                                                case "":
                                                                                    v106 := GenS18{F0: (v60 * fi58), F1: 2147483648}
                                                                                    _ = v106
                                                                                default:
                                                                                    switch (fi58 - (v60 + int(v3))) {
                                                                                    case -1:
                                                                                        var v107 int64 = 864
                                                                                        _ = v107
                                                                                    case 506:
                                                                                        var v108 float64 = (0.0 + float64(v2))
                                                                                        _ = v108
                                                                                    case 32767:
                                                                                        switch v57 {
                                                                                        case "":
                                                                                            fmt.Println(v3)
                                                                                        case "foo bar":
                                                                                            v3 += (v3 * v6)
                                                                                        case "a":
                                                                                            v67 = v67
                                                                                            if v67 {
                                                                                                if (1 > (v40 & v40)) {
                                                                                                    delete(v4, 4294967295)
                                                                                                    var v109 int32 = -119
                                                                                                    _ = v109
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v52)
                                                                                                v110 := v5[184]
                                                                                                var v111 int = (v52 << ((v6) & 63))
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                        default:
                                                                                            var v112 int = v60
                                                                                            _ = v112
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    switch fi56 {
                                                                                    case -1:
                                                                                        if false {
                                                                                            var v113 float64 = -2.25
                                                                                            _ = v113
                                                                                        }
                                                                                    case 32767:
                                                                                        var v114 int32 = (^(v40))
                                                                                        _ = v114
                                                                                    case 1:
                                                                                        v5[(fi58 << ((v6) & 63))] = ((v3 & v3) == 844)
                                                                                    default:
                                                                                        v40 -= 1
                                                                                    }
                                                                                    v3 += 2
                                                                                case "hello":
                                                                                    v67 = true
                                                                                case "\n\t":
                                                                                    if ((^(v1)) <= (v1 >> ((v3) & 63))) {
                                                                                        if (v3 <= (v6 + v6)) {
                                                                                            if len(v43) > 0 {
                                                                                                v43[0] = (^(v60))
                                                                                            }
                                                                                            _ = len(v43)
                                                                                        }
                                                                                    } else {
                                                                                        var v115 int8 = 52
                                                                                        if (v61 || false) {
                                                                                            var v116 float32 = ((-2.25 + ((v2 / 1e10) * (0.1 - v2))) / float32(3.14159))
                                                                                            if false {
                                                                                                if true {
                                                                                                    switch ("foo bar" + v57) {
                                                                                                    case "hello":
                                                                                                        switch v57 {
                                                                                                        case "\n\t":
                                                                                                            switch v60 {
                                                                                                            case 32768:
                                                                                                                fmt.Println(v52, fi58)
                                                                                                            case 67:
                                                                                                                v67 = ((v3 >> ((v6) & 63)) > v6)
                                                                                                            default:
                                                                                                                if (v67 && true) {
                                                                                                                    if false {
                                                                                                                        v67 = v61
                                                                                                                        v117 := GenS19{F0: ((^(v3)) >= v6)}
                                                                                                                        _ = v117
                                                                                                                    } else {
                                                                                                                        v2 += v2
                                                                                                                        if true {
                                                                                                                            switch v57 {
                                                                                                                            case "a":
                                                                                                                                v3 = (v3 >> (((v3 | 300)) & 63))
                                                                                                                            }
                                                                                                                            v61 = v61
                                                                                                                            var v118 int64 = 9223372036854775807
                                                                                                                            _ = v118
                                                                                                                        } else {
                                                                                                                            var v119 bool = false
                                                                                                                            _ = v119
                                                                                                                        }
                                                                                                                        v120 := GenS20{F0: 42, F1: (v1 ^ v1)}
                                                                                                                        var v121 bool = v67
                                                                                                                        _ = v120
                                                                                                                        _ = v121
                                                                                                                    }
                                                                                                                    if v61 {
                                                                                                                        if ((v115 & -128) >= v115) {
                                                                                                                            fmt.Println(v52)
                                                                                                                            switch 128 {
                                                                                                                            case 770:
                                                                                                                                var v122 uint32 = uint32(v3)
                                                                                                                                _ = v122
                                                                                                                            case 376:
                                                                                                                                switch (v52 & 0) {
                                                                                                                                case 128:
                                                                                                                                    if v67 {
                                                                                                                                        var v123 int32 = v40
                                                                                                                                        _ = v123
                                                                                                                                    } else {
                                                                                                                                        v124 := GenS21{F0: 1, F1: (^(v115)), F2: genFn0((4294967295 < uint64(v6)), v67, v1)}
                                                                                                                                        v125 := GenS12{F0: -2147483648, F1: v1}
                                                                                                                                        _ = v124
                                                                                                                                        _ = v125
                                                                                                                                    }
                                                                                                                                case 32767:
                                                                                                                                    if len(v43) >= 2 {
                                                                                                                                        v43 = v43[1:len(v43)]
                                                                                                                                    }
                                                                                                                                    delete(v5, v52)
                                                                                                                                    v126 := GenS22{F0: 65535, F1: true, F2: 42}
                                                                                                                                    _ = v126
                                                                                                                                case 0:
                                                                                                                                    var v127 uint = v6
                                                                                                                                    _ = v127
                                                                                                                                }
                                                                                                                            case 94:
                                                                                                                                fmt.Println(v60)
                                                                                                                            default:
                                                                                                                                v1 -= (^(v1))
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if false {
                                                                                                                            v57 = v57
                                                                                                                            v128 := v5[int(v1)]
                                                                                                                            _ = v128
                                                                                                                        } else {
                                                                                                                            if v67 {
                                                                                                                                v2 += 1.5
                                                                                                                                v60 = v52
                                                                                                                            }
                                                                                                                            if (v67 && false) {
                                                                                                                                var v129 uint8 = (v1 & 128)
                                                                                                                                _ = v129
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                var v130 float32 = (v116 + v2)
                                                                                                                                if (false || (v40 <= (v40 - (v40 / (((v40 + 42)) | 1))))) {
                                                                                                                                    if v67 {
                                                                                                                                        var v131 int32 = v40
                                                                                                                                        _ = len(v5)
                                                                                                                                        v132 := GenS23{F0: 0, F1: v60}
                                                                                                                                        _ = v131
                                                                                                                                        _ = v132
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v133 int8 = (v115 | v115)
                                                                                                                                _ = v130
                                                                                                                                _ = v133
                                                                                                                            }
                                                                                                                            v1 += (v1 - (v1 / ((1) | 1)))
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if v61 {
                                                                                                                            _ = len(v5)
                                                                                                                            v43 = append(v43, 115)
                                                                                                                            v1 += 1
                                                                                                                        }
                                                                                                                        v6 += (v6 & v6)
                                                                                                                        if v67 {
                                                                                                                            v67 = (72 <= v1)
                                                                                                                            var v134 uint16 = 5
                                                                                                                            v40 += (v40 >> (((v3 % (((v3 | v3)) | 1))) & 63))
                                                                                                                            _ = v134
                                                                                                                        }
                                                                                                                        fmt.Println(v3, v3)
                                                                                                                    }
                                                                                                                    v135 := GenS24{F0: 297}
                                                                                                                    _ = v135
                                                                                                                }
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            switch fi58 {
                                                                                                            case 1:
                                                                                                                var v136 string = "a"
                                                                                                                _ = v136
                                                                                                            case 127:
                                                                                                                if ((true && true) || v67) {
                                                                                                                    var v137 float64 = 0.1
                                                                                                                    _ = v137
                                                                                                                } else {
                                                                                                                    v138 := GenS7{F0: 42, F1: 1, F2: true}
                                                                                                                    var v139 uint32 = 65535
                                                                                                                    _ = v138
                                                                                                                    _ = v139
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(v116, v115)
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        v140 := v5[(^(v60))]
                                                                                                        _ = v140
                                                                                                    default:
                                                                                                        v67 = (-1 != 0)
                                                                                                    }
                                                                                                    v60 += v60
                                                                                                    if v67 {
                                                                                                        var v141 int32 = 0
                                                                                                        _ = v141
                                                                                                    } else {
                                                                                                        _ = len(v43)
                                                                                                        v115 -= v115
                                                                                                        if (!false) {
                                                                                                            v1 = 134
                                                                                                            v3 = uint(v1)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v142 float32 = 0.1
                                                                                                if (-32768 < 42) {
                                                                                                    var v143 int8 = 100
                                                                                                    _ = v143
                                                                                                }
                                                                                                var v144 int16 = -32768
                                                                                                _ = v142
                                                                                                _ = v144
                                                                                            }
                                                                                            _ = v116
                                                                                        }
                                                                                        if (1 == -1) {
                                                                                            v67 = v61
                                                                                        } else {
                                                                                            v61 = false
                                                                                        }
                                                                                        _ = v115
                                                                                    }
                                                                                default:
                                                                                    v3 = v6
                                                                                }
                                                                                v40 += 550
                                                                            }
                                                                            v6 = 2
                                                                            if ((!v61) && true) {
                                                                                if len(v43) > 0 {
                                                                                    v43[0] = 128
                                                                                }
                                                                                var v145 float32 = 1.5
                                                                                if (v3 > (^(v3))) {
                                                                                    fmt.Println(v61, v145, v1)
                                                                                    v146 := GenS22{F0: 2, F1: v67, F2: (fi58 << (((v6 ^ v6)) & 63))}
                                                                                    v147 := GenS4{F0: 146}
                                                                                    var v148 float32 = v145
                                                                                    _ = v146
                                                                                    _ = v147
                                                                                    _ = v148
                                                                                }
                                                                                _ = v145
                                                                            }
                                                                        } else {
                                                                            var v149 int = 42
                                                                            v6 += v3
                                                                            _ = v149
                                                                        }
                                                                    case 32768:
                                                                        v150 := GenS14{F0: v40, F1: 255, F2: (!false)}
                                                                        _ = v150
                                                                    case 2147483647:
                                                                        if false {
                                                                            v57 = ((("\n\t" + (v57 + v57)) + ("foo bar" + "")) + v57)
                                                                            _ = len(v43)
                                                                            switch v52 {
                                                                            case -615:
                                                                                fmt.Println(v57, v52, v2)
                                                                            case 0:
                                                                                var v151 float64 = 1.5
                                                                                _ = v151
                                                                            case -215:
                                                                                var v152 float32 = 0.0
                                                                                _ = v152
                                                                            }
                                                                            v153 := GenS25{F0: (v1 & (v1 - v1))}
                                                                            _ = v153
                                                                        }
                                                                    default:
                                                                        var v154 float32 = v2
                                                                        _ = v154
                                                                    }
                                                                }
                                                                if (!(!v67)) {
                                                                    v155 := GenS26{F0: v40, F1: 0, F2: fi56}
                                                                    if ((1e10 * -2.25) < 3.14159) {
                                                                        if v67 {
                                                                            fmt.Println(v60, v2, fi58)
                                                                            fmt.Println(v61, v60, v3)
                                                                            v4[4294967295] = v61
                                                                        }
                                                                        v156 := GenS1{F0: v61, F1: 32767, F2: -89}
                                                                        _ = v156
                                                                    } else {
                                                                        var v157 int8 = 0
                                                                        _ = v157
                                                                    }
                                                                    _ = v155
                                                                }
                                                            }
                                                            v158 := v5[32767]
                                                            _ = v158
                                                        } else {
                                                            if len(v43) > 0 {
                                                                v43[0] = (fi58 ^ (v60 & v60))
                                                            }
                                                            v2 += v2
                                                            fmt.Println(v1, v3, v6)
                                                            var v159 uint32 = uint32(v1)
                                                            if false {
                                                                v52 -= fi56
                                                            }
                                                            _ = v159
                                                        }
                                                        _ = v75
                                                    }
                                                default:
                                                    delete(v5, v60)
                                                }
                                                fmt.Println(v61)
                                            }
                                        }
                                        v40 = (v40 * (v40 >> (((v3 + v6)) & 63)))
                                    } else {
                                        var v160 uint = uint(v1)
                                        _ = v160
                                    }
                                    _ = v67
                                }
                                v161 := GenS19{F0: (v52 > 27)}
                                _ = fi58
                                _ = v161
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                        case "a":
                            var v162 int = (fi56 % ((fi56) | 1))
                            _ = v162
                        default:
                            var v163 string = v57
                            _ = v163
                        }
                        var v164 int = int(v6)
                        _ = fi56
                        _ = v164
                        _ = v57
                    }
                }
            case "hello":
                if len(v43) > 0 {
                    v43[0] = -1
                }
            default:
                var v165 int16 = 0
                _ = v165
            }
            v166 := v4[42]
            if ("hello" >= "\n\t") {
                var v167 float32 = 1.5
                for fi168 := 0; fi168 < 11; fi168++ {
                    v169 := GenS10{F0: v52, F1: 65535}
                    switch (fi168 & v52) {
                    case 1:
                        switch "" {
                        case "\n\t":
                            var v170 int16 = -246
                            _ = v170
                        case "hello":
                            if (!(100 == -53)) {
                                switch "a" {
                                case "a":
                                    for fi171 := 0; fi171 < 1; fi171++ {
                                        if v166 {
                                            v4[4294967295] = ((0.0 / 0.1) > -2.25)
                                            v3 += 1
                                            var v172 uint8 = 2
                                            _ = v172
                                        } else {
                                            fmt.Println(v167, fi171, fi171)
                                        }
                                        _ = fi171
                                    }
                                case "\n\t":
                                    var v173 uint = 660
                                    _ = v173
                                case "":
                                    if v166 {
                                        for fi174 := 0; fi174 < 7; fi174++ {
                                            if ((1e10 < v2) || v166) {
                                                if (710 != genFn0(false, false, 42)) {
                                                    fmt.Println(v52)
                                                    _ = len(v5)
                                                    v3 -= v6
                                                    var v175 uint32 = 2
                                                    var v176 int16 = -44
                                                    if false {
                                                        if (!false) {
                                                            var v177 uint32 = (^(v175))
                                                            v178 := GenS12{F0: -1, F1: (v1 | uint8(v6))}
                                                            if v166 {
                                                                var v179 float32 = v167
                                                                _ = v179
                                                            }
                                                            _ = v177
                                                            _ = v178
                                                        } else {
                                                            var v180 int = int(v175)
                                                            var v181 int8 = 100
                                                            v6 = (v6 >> ((v3) & 63))
                                                            v182 := GenS3{F0: 94, F1: 1}
                                                            if v166 {
                                                                fmt.Println(v181)
                                                            } else {
                                                                fmt.Println(fi174)
                                                                if ((v1 <= 63) || (v176 > v176)) {
                                                                    var v183 int = (v180 ^ v180)
                                                                    _ = v183
                                                                } else {
                                                                    var v184 int8 = (v181 / (((-(v181))) | 1))
                                                                    _ = v184
                                                                }
                                                                v176 += v176
                                                            }
                                                            _ = v180
                                                            _ = v181
                                                            _ = v182
                                                        }
                                                        var v185 uint64 = 799
                                                        var v186 bool = (9223372036854775807 <= 0)
                                                        _ = v185
                                                        _ = v186
                                                    }
                                                    _ = v175
                                                    _ = v176
                                                } else {
                                                    switch fi174 {
                                                    case -773:
                                                        v6 = v6
                                                    case 9223372036854775807:
                                                        v187 := GenS2{F0: 2147483648, F1: v40, F2: v1}
                                                        _ = v187
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    delete(v5, (v52 - 526))
                                                    if false {
                                                        var v188 int8 = -128
                                                        _ = v188
                                                    }
                                                }
                                            }
                                            fmt.Println(v2)
                                            _ = len(v43)
                                            _ = fi174
                                        }
                                        v52 = (v52 | (fi168 - v52))
                                        for fi189 := 0; fi189 < 20; fi189++ {
                                            v52 = -1
                                            if v166 {
                                                v190 := GenS14{F0: 1, F1: 255, F2: v166}
                                                if v166 {
                                                    _ = len(v4)
                                                } else {
                                                    fmt.Println(v167, v1, v167)
                                                    switch "a" {
                                                    case "":
                                                        _ = len(v5)
                                                    case "hello":
                                                        v2 = float32(1e10)
                                                    default:
                                                        var v191 uint32 = 1
                                                        _ = v191
                                                    }
                                                    fmt.Println(fi168)
                                                }
                                                _ = v190
                                            } else {
                                                if v166 {
                                                    switch "foo bar" {
                                                    case "a":
                                                        if v166 {
                                                            v167 += (1e10 * (float32(1e10) - float32(1.5)))
                                                            v2 = ((3.14159 / v2) / v167)
                                                            var v192 bool = v166
                                                            _ = v192
                                                        } else {
                                                            v193 := v4[144]
                                                            _ = v193
                                                        }
                                                    case "hello":
                                                        if false {
                                                            v194 := GenS22{F0: 255, F1: true, F2: (fi168 % ((fi189) | 1))}
                                                            _ = v194
                                                        }
                                                    case "foo bar":
                                                        v195 := GenS5{F0: 1, F1: 32767}
                                                        _ = v195
                                                    default:
                                                        if len(v43) >= 2 {
                                                            v43 = v43[1:len(v43)]
                                                        }
                                                    }
                                                    var v196 uint8 = (v1 * (^(v1)))
                                                    v40 = (v40 - 2147483647)
                                                    _ = v196
                                                }
                                            }
                                            _ = fi189
                                        }
                                    }
                                }
                            }
                        }
                    case 9223372036854775807:
                        if (("\n\t" + ("" + "foo bar")) < ("hello" + "\n\t")) {
                            v1 = (v1 >> (((^(v3))) & 63))
                            switch (v52 * int(v6)) {
                            case 32768:
                                if false {
                                    var v197 int16 = 983
                                    _ = v197
                                }
                            case 127:
                                fmt.Println(v52, v2, v52)
                            default:
                                for fi198 := 0; fi198 < 2; fi198++ {
                                    var v199 uint = v3
                                    _ = fi198
                                    _ = v199
                                }
                            }
                            var v200 float32 = v2
                            _ = v200
                        } else {
                            var v201 uint64 = 0
                            _ = v201
                        }
                    case 0:
                        v3 = (v3 ^ v6)
                    }
                    switch 222 {
                    case 32767:
                        fmt.Println(v3)
                    case 173:
                        v40 = -2147483648
                    }
                    _ = fi168
                    _ = v169
                }
                _ = v167
            }
            _ = v166
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            var v202 bool = false
            _ = v202
        case "":
            fmt.Println(v6, v52)
        default:
            v6 -= (v6 - 65535)
        }
        var v203 string = ""
        _ = v203
    }
    for fi204 := 0; fi204 < 7; fi204++ {
        for fi205 := 0; fi205 < 2; fi205++ {
            v5[v52] = (1 > 0)
            _ = fi205
        }
        if false {
            delete(v4, 682)
            if false {
                v2 = ((v2 + v2) - (-2.25 * v2))
                var v206 int8 = 0
                _ = v206
            } else {
                switch int(v1) {
                case -743:
                    v1 += v1
                case 795:
                    v40 = 2147483647
                case 899:
                    var v207 uint16 = 255
                    _ = v207
                }
                v6 -= v3
            }
            if true {
                var v208 uint16 = 993
                var v209 uint8 = (v1 * (v1 >> (((v3 / ((v6) | 1))) & 63)))
                switch fi204 {
                case 566:
                    var v210 bool = (!(false && (true && false)))
                    _ = v210
                case 32767:
                    v3 = v6
                default:
                    if ((v3 % ((uint(v208)) | 1)) >= (v3 + v6)) {
                        v1 -= v1
                    } else {
                        if len(v43) >= 2 {
                            v43 = v43[1:len(v43)]
                        }
                    }
                }
                _ = v208
                _ = v209
            }
        }
        _ = fi204
    }
    v211 := GenS16{F0: 993, F1: 311}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v211
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v52
    _ = v6
}
