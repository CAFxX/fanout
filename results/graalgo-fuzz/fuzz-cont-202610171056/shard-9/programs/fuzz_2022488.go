package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 bool; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 uint8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 uint; F2 string }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int64 }
func genFn0() bool {
    var v7 int16 = 705
    if (42 >= (^(v7))) {
        v8 := GenS1{F0: -376, F1: 506}
        _ = v8
    }
    _ = v7
    return (128 >= 1)
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]int)
    var v3 int32 = -1
    var v4 uint = 0
    v5 := make(map[uint32]int)
    var v6 uint = (^(v4))
    fmt.Println(v4, v6)
    v9 := GenS1{F0: 9223372036854775807, F1: uint32(v4)}
    v3 = v3
    if false {
        if true {
            switch 32767 {
            case -2147483648:
                v4 = v4
            case 2147483647:
                var v10 uint = v4
                _ = v10
            case 757:
                _ = len(v2)
            default:
                var v11 int64 = 42
                _ = v11
            }
            var v12 int32 = (v3 + int32(v4))
            _ = v12
        } else {
            v13 := []int{-2147483648, 1, -9223372036854775808}
            _ = v13
        }
        var v14 uint32 = 2147483648
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 float64 = 1e10
            if false {
                var v17 uint32 = (v14 >> ((2) & 63))
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := []int{fi18, (fi18 % (((fi15 / (((^(fi15))) | 1))) | 1))}
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                if true {
                    v14 -= (v14 % ((v14) | 1))
                    fmt.Println(v16, v4, v3)
                } else {
                    fmt.Println(v6, v16)
                    if ((v14 & v14) <= v14) {
                        v14 = uint32(v4)
                    } else {
                        fmt.Println(v6, fi15, fi15)
                    }
                    var v20 bool = true
                    _ = v20
                }
            }
            switch (v1 + (("foo bar" + v1) + "foo bar")) {
            case "a":
                switch -9223372036854775808 {
                case 9223372036854775807:
                    for fi21 := 0; fi21 < 3; fi21++ {
                        if ((42 >= 42) && false) {
                            if (true && (v6 <= v4)) {
                                v22 := []int{(fi15 ^ fi21), fi21}
                                v23 := GenS2{F0: 32767, F1: 100, F2: (v14 & 2)}
                                _ = v22
                                _ = v23
                            } else {
                                var v24 uint32 = (^(v14))
                                _ = v24
                            }
                            v6 = (^(v6))
                            v6 -= v4
                        }
                        var v25 bool = true
                        fmt.Println(v3, v16)
                        _ = fi21
                        _ = v25
                    }
                case -9223372036854775808:
                    var v26 uint8 = 53
                    _ = v26
                default:
                    v27 := GenS3{F0: 65535, F1: -231, F2: 65535}
                    _ = v27
                }
            case "\n\t":
                v3 -= (v3 ^ v3)
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 int8 = 100
            if false {
                var v30 bool = false
                var v31 uint = v4
                _ = v30
                _ = v31
            } else {
                switch "" {
                case "\n\t":
                    v1 = (v1 + v1)
                default:
                    switch (fi28 | (fi28 >> (((v6 << (((v4 + v4)) & 63))) & 63))) {
                    case 0:
                        for fi32 := 0; fi32 < 19; fi32++ {
                            if (0 != 4294967295) {
                                if (1e10 > 1e10) {
                                    var v33 uint = 4294967295
                                    v3 = (-(v3))
                                    switch int(v6) {
                                    case 0:
                                        var v34 float64 = 3.14159
                                        _ = v34
                                    default:
                                        _ = len(v2)
                                    }
                                    _ = v33
                                }
                            } else {
                                var v35 int32 = v3
                                var v36 float64 = ((1.5 - 0.1) * 0.0)
                                switch v1 {
                                case "a":
                                    if ((true && true) || true) {
                                        v37 := GenS4{F0: 0, F1: true}
                                        switch (v1 + "a") {
                                        case "hello":
                                            if false {
                                                v36 -= (v36 * v36)
                                                var v38 string = (((v1 + (v1 + "foo bar")) + v1) + "foo bar")
                                                _ = v38
                                            } else {
                                                fmt.Println(v3, v4)
                                                v6 -= v4
                                            }
                                        case "":
                                            var v39 bool = false
                                            v40 := []int{(fi28 & (-(fi32)))}
                                            v6 = (v4 / ((480) | 1))
                                            _ = v39
                                            _ = v40
                                        case "a":
                                            fmt.Println(v1, v4)
                                        default:
                                            v3 = (v3 | v35)
                                        }
                                        v4 -= v6
                                        _ = v37
                                    } else {
                                        var v41 int16 = -307
                                        _ = v41
                                    }
                                }
                                if false {
                                    v6 = v6
                                    if (!false) {
                                        var v42 uint = 913
                                        fmt.Println(v35)
                                        _ = v42
                                    }
                                    var v43 string = "foo bar"
                                    _ = v43
                                } else {
                                    fmt.Println(v6, v35, v29)
                                    v44 := []int{(fi28 * fi32), (fi32 | fi32)}
                                    _ = v44
                                }
                                _ = v35
                                _ = v36
                            }
                            if (!((1 == 0) && true)) {
                                switch 9223372036854775807 {
                                case -53:
                                    if (!false) {
                                        v4 = (v4 << ((65535) & 63))
                                        if true {
                                            switch -7 {
                                            case 32768:
                                                if ((0 > int64(v29)) || (356 == 114)) {
                                                    v45 := v2[v1]
                                                    v46 := []int{(fi28 ^ v45)}
                                                    _ = v45
                                                    _ = v46
                                                }
                                            default:
                                                if false {
                                                    var v47 uint8 = 123
                                                    _ = v47
                                                }
                                            }
                                            fmt.Println(fi28)
                                            var v48 uint8 = 99
                                            var v49 bool = true
                                            _ = v48
                                            _ = v49
                                        }
                                        v1 = v1
                                        if (fi28 <= fi32) {
                                            var v50 int32 = v3
                                            var v51 uint64 = 452
                                            _ = v50
                                            _ = v51
                                        } else {
                                            v52 := GenS2{F0: 879, F1: 23, F2: 2147483648}
                                            if (!(65535 < 171)) {
                                                if ((147 == 639) && ((-32768 < int16(v29)) && true)) {
                                                    switch v1 {
                                                    case "\n\t":
                                                        var v53 uint64 = 388
                                                        v4 = (^(v6))
                                                        v54 := GenS1{F0: 53, F1: 2147483647}
                                                        switch v1 {
                                                        case "\n\t":
                                                            v55 := GenS5{F0: v6}
                                                            var v56 int = fi32
                                                            _ = v55
                                                            _ = v56
                                                        case "":
                                                            var v57 uint16 = 32767
                                                            _ = v57
                                                        case "a":
                                                            fmt.Println(v4, v53, v6)
                                                        default:
                                                            var v58 uint16 = 231
                                                            _ = v58
                                                        }
                                                        _ = v53
                                                        _ = v54
                                                    case "a":
                                                        fmt.Println(v3, fi28)
                                                        v59 := GenS1{F0: -988, F1: 0}
                                                        var v60 string = (v1 + "")
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    var v61 int = fi28
                                                    v3 = (v3 * (^(v3)))
                                                    _ = v61
                                                } else {
                                                    var v62 uint8 = 255
                                                    switch v1 {
                                                    case "":
                                                        var v63 float32 = 0.0
                                                        _ = v63
                                                    }
                                                    if true {
                                                        var v64 uint32 = 2147483647
                                                        _ = v64
                                                    } else {
                                                        if ((int(v3) != fi32) && (false && false)) {
                                                            switch "" {
                                                            case "a":
                                                                v65 := GenS4{F0: 4294967295, F1: true}
                                                                _ = v65
                                                            default:
                                                                v66 := GenS4{F0: 0, F1: (1e10 == 1e10)}
                                                                _ = v66
                                                            }
                                                            var v67 uint16 = 440
                                                            var v68 int32 = -106
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            var v69 uint32 = 0
                                                            v5[(v69 % (((v69 >> (((v6 / (((v6 & 0)) | 1))) & 63))) | 1))] = 42
                                                            if ((!((65535 == 1) && true)) || (((v1 + "") + "a") < "hello")) {
                                                                if true {
                                                                    v4 = (v4 << ((728) & 63))
                                                                } else {
                                                                    switch fi32 {
                                                                    case -1:
                                                                        v3 = (v3 - v3)
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                    v70 := GenS5{F0: v6}
                                                                    switch v1 {
                                                                    case "a":
                                                                        var v71 int64 = -9223372036854775808
                                                                        _ = v71
                                                                    case "":
                                                                        var v72 string = ("a" + v1)
                                                                        _ = v72
                                                                    case "hello":
                                                                        v1 = v1
                                                                    }
                                                                    _ = v70
                                                                }
                                                                fmt.Println(v6, v62)
                                                            } else {
                                                                if (((((3.14159 + 3.14159) + ((0.1 / 1.0) - 0.0)) * (-2.25 / 0.1)) - float64(3.14159)) == 0.1) {
                                                                    delete(v5, (v69 | v69))
                                                                    v73 := GenS5{F0: v4}
                                                                    var v74 bool = (437 < 1000)
                                                                    v75 := []int{(-(fi28)), int(v4), (^(fi32)), -9223372036854775808}
                                                                    _ = v73
                                                                    _ = v74
                                                                    _ = v75
                                                                } else {
                                                                    var v76 string = v1
                                                                    fmt.Println(v62, v62, fi32)
                                                                    if (((!(!(!true))) || (-160 > (fi28 << ((0) & 63)))) && true) {
                                                                        if (392 == 42) {
                                                                            v29 = (v29 + v29)
                                                                            var v77 int32 = (v3 << ((v4) & 63))
                                                                            _ = v77
                                                                        } else {
                                                                            if (false || true) {
                                                                                switch -9223372036854775808 {
                                                                                case 395:
                                                                                    v78 := GenS3{F0: 65535, F1: 524, F2: 65535}
                                                                                    _ = v78
                                                                                case 127:
                                                                                    var v79 float64 = -2.25
                                                                                    _ = v79
                                                                                }
                                                                                var v80 int = (fi32 / ((fi28) | 1))
                                                                                var v81 string = "a"
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                var v82 bool = ((!false) && true)
                                                                                _ = v82
                                                                            }
                                                                            v69 = v69
                                                                            fmt.Println(v4, v29)
                                                                        }
                                                                    }
                                                                    fmt.Println(fi28, v1, v76)
                                                                    var v83 int8 = v29
                                                                    _ = v76
                                                                    _ = v83
                                                                }
                                                                v84 := GenS6{F0: 42, F1: (-1 != -9223372036854775808), F2: -308}
                                                                _ = v84
                                                            }
                                                            _ = v69
                                                        }
                                                        if false {
                                                            var v85 uint16 = 32767
                                                            _ = v85
                                                        }
                                                        v86 := []int{-2147483648, fi32, (fi32 >> ((v6) & 63)), 42}
                                                        _ = v86
                                                    }
                                                    _ = v62
                                                }
                                                var v87 int8 = (^(v29))
                                                v88 := GenS5{F0: v6}
                                                _ = v87
                                                _ = v88
                                            } else {
                                                switch ("foo bar" + (v1 + ("foo bar" + ("foo bar" + (v1 + v1))))) {
                                                case "":
                                                    v29 -= (v29 >> ((v6) & 63))
                                                case "a":
                                                    var v89 float32 = 3.14159
                                                    _ = v89
                                                case "foo bar":
                                                    if true {
                                                        v90 := []int{fi32, fi28, fi28}
                                                        v91 := GenS2{F0: 588, F1: (v29 & (v29 ^ (^(v29)))), F2: 65535}
                                                        switch "\n\t" {
                                                        case "":
                                                            var v92 int = fi32
                                                            _ = v92
                                                        case "\n\t":
                                                            var v93 int32 = (-(v3))
                                                            var v94 uint32 = 42
                                                            v6 += v4
                                                            v29 = (-(v29))
                                                            if false {
                                                                var v95 uint8 = 2
                                                                _ = v95
                                                            } else {
                                                                v96 := GenS2{F0: 65535, F1: (v29 & 42), F2: uint32(v4)}
                                                                v97 := GenS2{F0: 1, F1: 1, F2: (v94 % ((4294967295) | 1))}
                                                                if false {
                                                                    if (-32768 < 606) {
                                                                        if true {
                                                                            if len(v90) >= 2 {
                                                                                v90 = v90[1:len(v90)]
                                                                            }
                                                                            v93 += -2147483648
                                                                            v98 := GenS4{F0: 65535, F1: ((true && (127 >= 128)) && ((int(v6) == fi28) && true))}
                                                                            _ = v98
                                                                        } else {
                                                                            if ((fi32 << (((v4 ^ v4)) & 63)) == 4) {
                                                                                v5[(v94 - v94)] = fi28
                                                                                v1 = ("a" + v1)
                                                                                v2["\n\t"] = (fi32 - fi32)
                                                                            } else {
                                                                                v99 := GenS7{F0: -493}
                                                                                v90 = append(v90, (fi32 >> ((v6) & 63)))
                                                                                switch v1 {
                                                                                case "hello":
                                                                                    fmt.Println(v93)
                                                                                case "\n\t":
                                                                                    v90 = append(v90, fi32)
                                                                                case "":
                                                                                    var v100 uint16 = 0
                                                                                    _ = v100
                                                                                }
                                                                                _ = v99
                                                                            }
                                                                        }
                                                                        if len(v90) > 0 {
                                                                            v90[0] = fi32
                                                                        }
                                                                    } else {
                                                                        var v101 int32 = (v93 >> ((v6) & 63))
                                                                        switch (fi28 * (fi32 >> ((v4) & 63))) {
                                                                        case 9223372036854775807:
                                                                            var v102 int = int(v29)
                                                                            _ = v102
                                                                        case -2147483648:
                                                                            v4 = v6
                                                                        case -9223372036854775808:
                                                                            var v103 uint64 = 42
                                                                            _ = v103
                                                                        }
                                                                        v104 := GenS2{F0: 32767, F1: v29, F2: (^(v94))}
                                                                        _ = v101
                                                                        _ = v104
                                                                    }
                                                                    if (true && true) {
                                                                        v93 += v93
                                                                    } else {
                                                                        v105 := GenS1{F0: (fi32 ^ fi32), F1: uint32(v6)}
                                                                        _ = v105
                                                                    }
                                                                } else {
                                                                    var v106 int16 = -1
                                                                    switch v1 {
                                                                    case "hello":
                                                                        var v107 int64 = 103
                                                                        v90 = append(v90, fi28)
                                                                        switch (fi28 & (fi32 * 128)) {
                                                                        case -1:
                                                                            v90 = append(v90, (^(fi32)))
                                                                        case 128:
                                                                            switch (v1 + ("foo bar" + v1)) {
                                                                            case "foo bar":
                                                                                var v108 uint64 = 614
                                                                                _ = v108
                                                                            case "a":
                                                                                v109 := GenS8{F0: v1}
                                                                                _ = v109
                                                                            case "\n\t":
                                                                                var v110 int = fi32
                                                                                _ = v110
                                                                            }
                                                                        case -9223372036854775808:
                                                                            v6 = (v4 ^ (^(v4)))
                                                                        }
                                                                        _ = v107
                                                                    case "foo bar":
                                                                        switch (fi28 + int(v106)) {
                                                                        case -791:
                                                                            var v111 uint16 = 2
                                                                            _ = v111
                                                                        case 4294967295:
                                                                            var v112 string = v1
                                                                            _ = v112
                                                                        default:
                                                                            v93 -= v93
                                                                        }
                                                                    case "\n\t":
                                                                        switch ("a" + "a") {
                                                                        case "\n\t":
                                                                            v93 += (v93 * (v3 * 827))
                                                                        case "foo bar":
                                                                            v106 = (v106 >> (((v4 << ((708) & 63))) & 63))
                                                                        case "a":
                                                                            v4 += (v4 ^ (v6 & (v6 % ((v6) | 1))))
                                                                        default:
                                                                            v3 = (v3 % ((v93) | 1))
                                                                            if (v6 == (v4 & (v6 << ((282) & 63)))) {
                                                                                switch fi28 {
                                                                                case -148:
                                                                                    if true {
                                                                                        switch (fi28 / ((fi32) | 1)) {
                                                                                        case 9223372036854775807:
                                                                                            if true {
                                                                                                if len(v90) > 0 {
                                                                                                    v90[0] = (fi32 % (((fi32 + fi28)) | 1))
                                                                                                }
                                                                                                var v113 uint8 = 42
                                                                                                _ = v113
                                                                                            } else {
                                                                                                _ = len(v2)
                                                                                                if true {
                                                                                                    var v114 float64 = (0.0 / float64((1.5 + (3.14159 + (float32((3.14159 * 3.14159)) + ((1.5 - 3.14159) - -2.25))))))
                                                                                                    if len(v90) >= 2 {
                                                                                                        v90 = v90[1:len(v90)]
                                                                                                    }
                                                                                                    _ = v114
                                                                                                } else {
                                                                                                    if false {
                                                                                                        _ = len(v90)
                                                                                                        switch fi32 {
                                                                                                        case -103:
                                                                                                            var v115 uint16 = 0
                                                                                                            _ = v115
                                                                                                        case -2147483648:
                                                                                                            switch ("hello" + (((v1 + v1) + v1) + "\n\t")) {
                                                                                                            case "foo bar":
                                                                                                                if (v94 != v94) {
                                                                                                                    v106 = -1
                                                                                                                    var v116 int16 = v106
                                                                                                                    _ = v116
                                                                                                                } else {
                                                                                                                    if (518 >= 32767) {
                                                                                                                        v117 := v5[(v94 / (((v94 / ((v94) | 1))) | 1))]
                                                                                                                        v118 := GenS3{F0: 0, F1: int64(v29), F2: 1}
                                                                                                                        var v119 uint = (v6 / ((v6) | 1))
                                                                                                                        _ = v117
                                                                                                                        _ = v118
                                                                                                                        _ = v119
                                                                                                                    } else {
                                                                                                                        var v120 int32 = v93
                                                                                                                        switch int(v106) {
                                                                                                                        case 4294967295:
                                                                                                                            v4 += v6
                                                                                                                        case -9223372036854775808:
                                                                                                                            _ = len(v2)
                                                                                                                        }
                                                                                                                        _ = v120
                                                                                                                    }
                                                                                                                }
                                                                                                            case "":
                                                                                                                v29 -= 0
                                                                                                            default:
                                                                                                                v121 := GenS7{F0: (v106 + (v106 % (((v106 * 1)) | 1)))}
                                                                                                                _ = v121
                                                                                                            }
                                                                                                        default:
                                                                                                            var v122 float32 = (0.0 + 3.14159)
                                                                                                            _ = v122
                                                                                                        }
                                                                                                        switch (fi28 / (((fi32 & (fi28 | 0))) | 1)) {
                                                                                                        case -881:
                                                                                                            v123 := GenS5{F0: 255}
                                                                                                            _ = v123
                                                                                                        case 4294967295:
                                                                                                            delete(v2, (v1 + "hello"))
                                                                                                        case 0:
                                                                                                            if (false && true) {
                                                                                                                delete(v2, v1)
                                                                                                            }
                                                                                                        default:
                                                                                                            v124 := GenS7{F0: -887}
                                                                                                            _ = v124
                                                                                                        }
                                                                                                    } else {
                                                                                                        v125 := GenS1{F0: int(v29), F1: v94}
                                                                                                        v90 = append(v90, (fi32 + 511))
                                                                                                        _ = v125
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case -1:
                                                                                            v93 = v3
                                                                                        default:
                                                                                            var v126 uint8 = 2
                                                                                            _ = v126
                                                                                        }
                                                                                        v127 := GenS1{F0: 127, F1: v94}
                                                                                        _ = v127
                                                                                    }
                                                                                }
                                                                                var v128 uint = 18446744073709551615
                                                                                _ = v128
                                                                            } else {
                                                                                v129 := GenS9{F0: (v106 | v106)}
                                                                                switch "a" {
                                                                                case "":
                                                                                    v29 += (v29 >> ((v6) & 63))
                                                                                case "hello":
                                                                                    if (!false) {
                                                                                        var v130 int32 = (v93 % ((-2147483648) | 1))
                                                                                        var v131 uint8 = 142
                                                                                        v130 = v3
                                                                                        _ = v130
                                                                                        _ = v131
                                                                                    } else {
                                                                                        v132 := GenS9{F0: v106}
                                                                                        if true {
                                                                                            v1 = "hello"
                                                                                            v133 := GenS10{F0: fi32, F1: 196}
                                                                                            switch (((v1 + ("hello" + v1)) + "\n\t") + "\n\t") {
                                                                                            case "hello":
                                                                                                if false {
                                                                                                    v134 := GenS6{F0: v106, F1: (!true), F2: 217}
                                                                                                    _ = v134
                                                                                                }
                                                                                            case "\n\t":
                                                                                                fmt.Println(v94, v106)
                                                                                            case "foo bar":
                                                                                                v135 := GenS11{F0: 128}
                                                                                                _ = v135
                                                                                            default:
                                                                                                v6 -= uint(v94)
                                                                                            }
                                                                                            v3 -= 32
                                                                                            _ = v133
                                                                                        } else {
                                                                                            v1 = ((v1 + ("a" + (v1 + "hello"))) + v1)
                                                                                            var v136 int64 = 0
                                                                                            _ = v136
                                                                                        }
                                                                                        _ = v132
                                                                                    }
                                                                                default:
                                                                                    if true {
                                                                                        var v137 string = (((("\n\t" + "foo bar") + ("hello" + "a")) + v1) + "hello")
                                                                                        _ = v137
                                                                                    } else {
                                                                                        fmt.Println(v1, v4)
                                                                                        var v138 int64 = 1
                                                                                        _ = v138
                                                                                    }
                                                                                }
                                                                                _ = v129
                                                                            }
                                                                            var v139 int64 = 1
                                                                            _ = v139
                                                                        }
                                                                    }
                                                                    var v140 uint32 = v94
                                                                    _ = v106
                                                                    _ = v140
                                                                }
                                                                fmt.Println(v6, v4, v94)
                                                                if (42 > 42) {
                                                                    var v141 int32 = (v3 % ((v3) | 1))
                                                                    var v142 uint16 = 1
                                                                    _ = v141
                                                                    _ = v142
                                                                } else {
                                                                    switch (fi32 * 128) {
                                                                    case 0:
                                                                        var v143 int8 = v29
                                                                        _ = v143
                                                                    case -9223372036854775808:
                                                                        switch (fi28 << (((^(v6))) & 63)) {
                                                                        case -114:
                                                                            var v144 int16 = -152
                                                                            _ = v144
                                                                        case 42:
                                                                            var v145 int16 = int16(v94)
                                                                            _ = v145
                                                                        case 128:
                                                                            v93 -= v93
                                                                        default:
                                                                            if (!true) {
                                                                                var v146 int64 = 9223372036854775807
                                                                                if len(v90) >= 2 {
                                                                                    v90 = v90[1:len(v90)]
                                                                                }
                                                                                _ = v146
                                                                            } else {
                                                                                v29 += 100
                                                                            }
                                                                        }
                                                                    case 916:
                                                                        var v147 string = v1
                                                                        _ = v147
                                                                    default:
                                                                        fmt.Println(v93, v1)
                                                                    }
                                                                    fmt.Println(v3, v6)
                                                                    switch fi28 {
                                                                    case -547:
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                            switch v1 {
                                                            case "hello":
                                                                if len(v90) > 0 {
                                                                    v90[0] = (-(fi32))
                                                                }
                                                            case "":
                                                                var v148 int16 = 42
                                                                _ = v148
                                                            }
                                                            switch (-(fi28)) {
                                                            case 4294967295:
                                                                v149 := GenS7{F0: int16(v4)}
                                                                _ = v149
                                                            case 931:
                                                                v4 += v6
                                                            default:
                                                                delete(v5, v94)
                                                            }
                                                            if true {
                                                                if (!true) {
                                                                    if false {
                                                                        v3 += v3
                                                                        v150 := GenS11{F0: 0}
                                                                        _ = len(v5)
                                                                        _ = v150
                                                                    } else {
                                                                        fmt.Println(v4, v93, v29)
                                                                        v6 += (v6 / (((v6 & (v4 >> ((v4) & 63)))) | 1))
                                                                        v94 += (^(v94))
                                                                    }
                                                                }
                                                                v151 := GenS10{F0: (fi32 >> ((42) & 63)), F1: uint8(v4)}
                                                                _ = v151
                                                            } else {
                                                                fmt.Println(v29, v29)
                                                                v90 = append(v90, (fi28 / (((fi32 / (((fi32 + int(v94))) | 1))) | 1)))
                                                            }
                                                            _ = len(v2)
                                                            fmt.Println(fi32)
                                                            v94 += (v94 | (v94 >> ((v4) & 63)))
                                                            _ = v93
                                                            _ = v94
                                                        default:
                                                            if (586 < 255) {
                                                                v152 := GenS11{F0: 128}
                                                                v153 := GenS1{F0: fi32, F1: 2}
                                                                switch 127 {
                                                                case 4294967295:
                                                                    var v154 int8 = v29
                                                                    _ = v154
                                                                case 128:
                                                                    var v155 uint32 = 2147483648
                                                                    _ = v155
                                                                }
                                                                _ = v152
                                                                _ = v153
                                                            } else {
                                                                if true {
                                                                    switch (fi32 % (((fi28 - fi28)) | 1)) {
                                                                    case 248:
                                                                        if true {
                                                                            v6 -= 390
                                                                            if false {
                                                                                switch fi32 {
                                                                                case 32767:
                                                                                    var v156 uint16 = 501
                                                                                    _ = v156
                                                                                case -2147483648:
                                                                                    switch v1 {
                                                                                    case "hello":
                                                                                        if true {
                                                                                            var v157 int32 = v3
                                                                                            var v158 int = (fi32 - fi28)
                                                                                            var v159 string = v1
                                                                                            _ = v157
                                                                                            _ = v158
                                                                                            _ = v159
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v160 := GenS12{F0: 0, F1: v6}
                                                                                        _ = v160
                                                                                    }
                                                                                case -1:
                                                                                    switch fi32 {
                                                                                    case 305:
                                                                                        v6 = v4
                                                                                    case 32768:
                                                                                        var v161 uint16 = 22
                                                                                        switch v1 {
                                                                                        case "a":
                                                                                            switch fi28 {
                                                                                            case -9223372036854775808:
                                                                                                _ = len(v90)
                                                                                            case 128:
                                                                                                fmt.Println(fi28, v161)
                                                                                                v3 += 1
                                                                                                if (v3 > (^(v3))) {
                                                                                                    v90 = append(v90, fi28)
                                                                                                    var v162 bool = true
                                                                                                    switch fi32 {
                                                                                                    case 127:
                                                                                                        var v163 string = (v1 + (v1 + v1))
                                                                                                        _ = v163
                                                                                                    case 9223372036854775807:
                                                                                                        if v162 {
                                                                                                            if v162 {
                                                                                                                var v164 int32 = (v3 >> (((v6 << ((v4) & 63))) & 63))
                                                                                                                v5[2147483648] = fi32
                                                                                                                fmt.Println(fi28, fi28, v161)
                                                                                                                _ = v164
                                                                                                            } else {
                                                                                                                delete(v5, 1)
                                                                                                            }
                                                                                                            if v162 {
                                                                                                                if (998 == 42) {
                                                                                                                    var v165 int64 = 0
                                                                                                                    var v166 uint8 = 2
                                                                                                                    _ = v165
                                                                                                                    _ = v166
                                                                                                                }
                                                                                                                var v167 uint = v6
                                                                                                                var v168 float64 = (1.5 * (((1.5 + 1e10) - 0.1) / 1.0))
                                                                                                                _ = v167
                                                                                                                _ = v168
                                                                                                            }
                                                                                                        } else {
                                                                                                            v169 := GenS13{F0: v161}
                                                                                                            _ = v169
                                                                                                        }
                                                                                                    case -1:
                                                                                                        v4 -= v6
                                                                                                    }
                                                                                                    v162 = v162
                                                                                                    _ = v162
                                                                                                } else {
                                                                                                    var v170 int64 = -1
                                                                                                    _ = v170
                                                                                                }
                                                                                                if ((v29 ^ v29) == (^(v29))) {
                                                                                                    var v171 int = (fi28 & (fi32 & (fi32 ^ (fi28 - 127))))
                                                                                                    fmt.Println(v171)
                                                                                                    v172 := GenS14{F0: (v3 & (v3 - 1)), F1: 548, F2: "a"}
                                                                                                    _ = v171
                                                                                                    _ = v172
                                                                                                }
                                                                                                if true {
                                                                                                    var v173 float32 = (0.1 + 1e10)
                                                                                                    _ = v173
                                                                                                } else {
                                                                                                    v174 := v5[2147483648]
                                                                                                    _ = v174
                                                                                                }
                                                                                                v6 -= v4
                                                                                                switch (fi28 % ((fi28) | 1)) {
                                                                                                case -9223372036854775808:
                                                                                                    var v175 int64 = -9223372036854775808
                                                                                                    _ = v175
                                                                                                case 9223372036854775807:
                                                                                                    if true {
                                                                                                        switch "a" {
                                                                                                        case "foo bar":
                                                                                                            v161 += (v161 >> ((1) & 63))
                                                                                                        case "a":
                                                                                                            var v176 uint8 = 1
                                                                                                            _ = v176
                                                                                                        case "":
                                                                                                            if ((1.5 != (-2.25 * 0.0)) || false) {
                                                                                                                v177 := GenS7{F0: 42}
                                                                                                                var v178 int = 32767
                                                                                                                _ = v177
                                                                                                                _ = v178
                                                                                                            }
                                                                                                        default:
                                                                                                            var v179 int = (fi32 >> ((v4) & 63))
                                                                                                            _ = v179
                                                                                                        }
                                                                                                        switch v1 {
                                                                                                        case "a":
                                                                                                            _ = len(v90)
                                                                                                        default:
                                                                                                            if len(v90) > 0 {
                                                                                                                v90[0] = 9223372036854775807
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v180 float32 = 1e10
                                                                                                        var v181 int16 = -32768
                                                                                                        _ = v180
                                                                                                        _ = v181
                                                                                                    }
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                _ = len(v2)
                                                                                                var v182 uint32 = 2
                                                                                                _ = v182
                                                                                            default:
                                                                                                switch v1 {
                                                                                                case "a":
                                                                                                    v161 = v161
                                                                                                case "hello":
                                                                                                    var v183 int16 = 42
                                                                                                    _ = v183
                                                                                                case "":
                                                                                                    var v184 uint32 = 65535
                                                                                                    v4 = v6
                                                                                                    _ = v184
                                                                                                default:
                                                                                                    _ = len(v5)
                                                                                                }
                                                                                            }
                                                                                        case "foo bar":
                                                                                            var v185 string = (v1 + v1)
                                                                                            _ = v185
                                                                                        default:
                                                                                            v186 := GenS10{F0: (-(fi32)), F1: 255}
                                                                                            _ = v186
                                                                                        }
                                                                                        if (303 < 423) {
                                                                                            var v187 uint64 = 4294967295
                                                                                            _ = v187
                                                                                        } else {
                                                                                            fmt.Println(v3, v161)
                                                                                        }
                                                                                        _ = v161
                                                                                    }
                                                                                default:
                                                                                    switch fi28 {
                                                                                    case 127:
                                                                                        var v188 int64 = 0
                                                                                        _ = v188
                                                                                    default:
                                                                                        if ("a" < v1) {
                                                                                            v6 += 65535
                                                                                            var v189 int16 = 119
                                                                                            if false {
                                                                                                v190 := v5[2147483648]
                                                                                                v191 := GenS7{F0: 1000}
                                                                                                if len(v90) > 0 {
                                                                                                    v90[0] = fi28
                                                                                                }
                                                                                                _ = len(v90)
                                                                                                _ = len(v90)
                                                                                                _ = v190
                                                                                                _ = v191
                                                                                            } else {
                                                                                                switch (fi32 >> (((v6 % (((v6 << ((v6) & 63))) | 1))) & 63)) {
                                                                                                case -9223372036854775808:
                                                                                                    if (487 <= 9223372036854775807) {
                                                                                                        fmt.Println(v189, fi28)
                                                                                                        var v192 float64 = 1.5
                                                                                                        var v193 uint8 = 42
                                                                                                        _ = v192
                                                                                                        _ = v193
                                                                                                    }
                                                                                                case -388:
                                                                                                    if false {
                                                                                                        v194 := GenS11{F0: 128}
                                                                                                        _ = v194
                                                                                                    } else {
                                                                                                        if (false || (!(v4 >= v4))) {
                                                                                                            var v195 int64 = 9223372036854775807
                                                                                                            v6 = (^(v6))
                                                                                                            delete(v2, "a")
                                                                                                            _ = v195
                                                                                                        }
                                                                                                        var v196 uint16 = 2
                                                                                                        _ = v196
                                                                                                    }
                                                                                                case 1:
                                                                                                    var v197 uint16 = 548
                                                                                                    _ = v197
                                                                                                }
                                                                                            }
                                                                                            v29 -= v29
                                                                                            _ = v189
                                                                                        }
                                                                                    }
                                                                                    var v198 int32 = (^(v3))
                                                                                    _ = v198
                                                                                }
                                                                                v90 = append(v90, -1)
                                                                                if (0.1 <= 0.0) {
                                                                                    v3 = v3
                                                                                    if len(v90) > 0 {
                                                                                        v90[0] = (fi32 / (((fi32 * 150)) | 1))
                                                                                    }
                                                                                } else {
                                                                                    _ = len(v5)
                                                                                    fmt.Println(v29)
                                                                                }
                                                                            } else {
                                                                                var v199 float32 = -2.25
                                                                                v2[(v1 + v1)] = (fi32 % (((^(fi32))) | 1))
                                                                                _ = v199
                                                                            }
                                                                        } else {
                                                                            v200 := GenS11{F0: 59}
                                                                            var v201 int32 = (v3 >> ((0) & 63))
                                                                            _ = v200
                                                                            _ = v201
                                                                        }
                                                                    case 32768:
                                                                        if (!((-(fi32)) >= fi32)) {
                                                                            var v202 uint = 42
                                                                            delete(v5, 2147483648)
                                                                            if (!false) {
                                                                                if false {
                                                                                    var v203 uint16 = 73
                                                                                    _ = v203
                                                                                }
                                                                                v202 -= 0
                                                                                if (!(v3 < v3)) {
                                                                                    var v204 string = "foo bar"
                                                                                    v204 = ("" + "")
                                                                                    var v205 uint8 = 0
                                                                                    _ = v204
                                                                                    _ = v205
                                                                                }
                                                                            } else {
                                                                                var v206 uint16 = 42
                                                                                _ = v206
                                                                            }
                                                                            _ = v202
                                                                        } else {
                                                                            var v207 int = fi32
                                                                            _ = v207
                                                                        }
                                                                    case 1:
                                                                        switch 517 {
                                                                        case 127:
                                                                            switch ("hello" + "") {
                                                                            case "":
                                                                                if true {
                                                                                    _ = len(v90)
                                                                                }
                                                                            default:
                                                                                v208 := GenS12{F0: (v3 >> ((0) & 63)), F1: (v4 % ((v4) | 1))}
                                                                                v3 = (^(v3))
                                                                                _ = v208
                                                                            }
                                                                        case 0:
                                                                            v209 := GenS15{F0: (v3 - (v3 ^ v3))}
                                                                            _ = v209
                                                                        }
                                                                    default:
                                                                        fmt.Println(v4, v4, fi32)
                                                                    }
                                                                    v210 := GenS5{F0: (v6 ^ 65535)}
                                                                    fmt.Println(v3)
                                                                    _ = v210
                                                                }
                                                                v6 = 64
                                                                v211 := GenS4{F0: 9, F1: true}
                                                                v1 = v1
                                                                var v212 int64 = 42
                                                                v212 += -9223372036854775808
                                                                _ = v211
                                                                _ = v212
                                                            }
                                                        }
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                }
                                            }
                                            _ = v52
                                        }
                                        v213 := GenS16{F0: 42}
                                        v6 -= v4
                                        v214 := v5[529]
                                        v3 -= 2147483647
                                        _ = v213
                                        _ = v214
                                    }
                                case 128:
                                    if (42 < 0) {
                                        v215 := []int{}
                                        v6 -= v4
                                        v6 += v4
                                        _ = v215
                                    }
                                default:
                                    switch "a" {
                                    case "":
                                        fmt.Println(v3, v1)
                                    case "foo bar":
                                        var v216 int16 = 13
                                        _ = v216
                                    }
                                    var v217 float64 = 1e10
                                    if true {
                                        v1 = v1
                                        var v218 int = fi28
                                        _ = v218
                                    }
                                    _ = v217
                                }
                                if false {
                                    if false {
                                        v5[56] = fi32
                                        v6 += v6
                                    }
                                    if ((!true) && true) {
                                        if false {
                                            fmt.Println(v29, v3, v4)
                                        } else {
                                            fmt.Println(v29, v6)
                                        }
                                        v219 := GenS2{F0: 1, F1: v29, F2: 0}
                                        var v220 float64 = 3.14159
                                        _ = v219
                                        _ = v220
                                    } else {
                                        var v221 uint8 = 128
                                        v2[v1] = fi28
                                        _ = v221
                                    }
                                }
                            }
                            _ = fi32
                        }
                    case 2147483647:
                        var v222 float32 = 1e10
                        _ = v222
                    case 4294967295:
                        var v223 uint16 = 411
                        _ = v223
                    }
                }
                v3 = (v3 << (((v6 + (v4 ^ (v4 * (v6 << ((18446744073709551615) & 63)))))) & 63))
            }
            _ = fi28
            _ = v29
        }
    }
    v224 := GenS12{F0: (-(v3)), F1: 4294967295}
    v6 = (^(v6))
    if (0.0 != 0.1) {
        if (-525 > -1) {
            v225 := GenS3{F0: 0, F1: 42, F2: 42}
            var v226 uint16 = 32767
            _ = v225
            _ = v226
        }
        for fi227 := 0; fi227 < 1; fi227++ {
            for fi228 := 0; fi228 < 17; fi228++ {
                switch (fi227 * 871) {
                case -77:
                    var v229 int64 = 42
                    _ = v229
                case 9223372036854775807:
                    v1 = v1
                case 32767:
                    fmt.Println(v4, v1, v3)
                }
                _ = fi228
            }
            _ = fi227
        }
        fmt.Println(v6, v3, v6)
    } else {
        fmt.Println(v4, v6)
        v230 := []int{127}
        var v231 uint64 = 369
        _ = v230
        _ = v231
    }
    for fi232 := 0; fi232 < 17; fi232++ {
        fmt.Println(v3, fi232, v3)
        v1 = v1
        var v233 int32 = 2147483647
        _ = fi232
        _ = v233
    }
    if true {
        v234 := []int{}
        _ = v234
    } else {
        switch v1 {
        case "":
            v6 -= (v6 ^ 255)
        }
        v235 := GenS16{F0: 1}
        v6 = v6
        _ = v235
    }
    v1 = (v1 + (("a" + ((v1 + v1) + v1)) + v1))
    v1 = ""
    var v236 int8 = -73
    v4 = 255
    if false {
        if (v6 != (^(v6))) {
            for fi237 := 0; fi237 < 4; fi237++ {
                var v238 int = (fi237 % (((fi237 << ((v4) & 63))) | 1))
                _ = fi237
                _ = v238
            }
            var v239 float64 = 1e10
            var v240 uint8 = 255
            _ = v239
            _ = v240
        }
        delete(v2, v1)
        if (0 < 2) {
            v236 += v236
            if (!true) {
                v236 -= (v236 - (v236 - 42))
            }
        }
    }
    var v241 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v236:", v236)
    fmt.Println("v241:", v241)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v224
    _ = v236
    _ = v241
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
