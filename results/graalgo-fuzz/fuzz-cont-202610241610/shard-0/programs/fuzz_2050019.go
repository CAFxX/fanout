package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int; F1 int64; F2 uint }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint; F2 string }
type GenS9 struct { F0 uint16; F1 int32; F2 int64 }
type GenS10 struct { F0 uint32; F1 string; F2 int32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint32; F1 int; F2 uint32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int; F1 int; F2 bool }
type GenS17 struct { F0 uint8; F1 uint64; F2 int }
type GenS18 struct { F0 int16; F1 string; F2 bool }
type GenS19 struct { F0 int; F1 bool; F2 int64 }
type GenS20 struct { F0 int16; F1 uint8 }
func genFn0() int64 {
    switch 863 {
    case 97:
        var v7 int16 = 1
        _ = v7
    case 9223372036854775807:
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(fi9, fi8)
                var v10 bool = (72 <= -75)
                v10 = (!true)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case -2147483648:
        var v11 bool = false
        _ = v11
    default:
        var v12 float32 = -2.25
        _ = v12
    }
    switch -9223372036854775808 {
    case 42:
        var v13 uint = 495
        _ = v13
    case -9223372036854775808:
        var v14 uint = 269
        _ = v14
    }
    return 42
}

func main() {
    v1 := []uint32{104, 1, 25, 4294967295}
    var v2 int = -808
    var v3 int64 = -383
    var v4 uint32 = 42
    v5 := make(map[int]bool)
    var v6 float32 = (0.0 * 0.0)
    if ((32767 > 245) || (false && true)) {
        if len(v1) > 0 {
            v1[0] = 480
        }
    }
    delete(v5, -9223372036854775808)
    var v15 int64 = (v3 + v3)
    v3 += (v3 << ((919) & 63))
    fmt.Println(v2, v6, v6)
    fmt.Println(v15, v6)
    delete(v5, (v2 & v2))
    v2 = (v2 * 2147483647)
    v16 := GenS1{F0: 1, F1: (v3 + genFn0())}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = (v2 * v2)
        fmt.Println(v2, v18)
        v19 := GenS1{F0: 119, F1: genFn0()}
        _ = fi17
        _ = v18
        _ = v19
    }
    if (65535 != 0) {
        v20 := v5[2147483647]
        var v21 int16 = 0
        _ = v20
        _ = v21
    } else {
        v22 := GenS1{F0: 107, F1: 1}
        var v23 int = (v2 - (v2 / ((32767) | 1)))
        if true {
            var v24 uint8 = 254
            fmt.Println(v4, v6)
            if ((42 != (v3 % ((v15) | 1))) || true) {
                fmt.Println(v23, v2)
                for fi25 := 0; fi25 < 18; fi25++ {
                    v3 -= int64(v23)
                    switch v23 {
                    case 435:
                        v26 := v5[218]
                        _ = v26
                    case 2147483647:
                        if (true && (!(1 <= -122))) {
                            var v27 float64 = 0.1
                            v2 += (v23 / (((fi25 << ((1) & 63))) | 1))
                            if (v27 < 0.0) {
                                fmt.Println(v23)
                                v27 = (v27 - 0.1)
                            } else {
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    switch "foo bar" {
                                    case "":
                                        var v29 uint32 = (v4 - 2147483648)
                                        _ = v29
                                    default:
                                        switch 128 {
                                        case -1:
                                            if (103 < 1) {
                                                if false {
                                                    v3 = int64(v4)
                                                }
                                                var v30 int16 = -149
                                                _ = v30
                                            } else {
                                                if true {
                                                    fmt.Println(v2)
                                                    v6 -= v6
                                                    v31 := GenS1{F0: 4294967295, F1: (-(v3))}
                                                    _ = v31
                                                } else {
                                                    switch ("" + "foo bar") {
                                                    case "hello":
                                                        if (true || false) {
                                                            switch fi28 {
                                                            case 9223372036854775807:
                                                                var v32 int16 = -32768
                                                                _ = len(v1)
                                                                _ = v32
                                                            case 0:
                                                                if false {
                                                                    v33 := GenS1{F0: 1, F1: genFn0()}
                                                                    v4 = 2
                                                                    _ = v33
                                                                }
                                                            }
                                                            v34 := GenS2{F0: 0}
                                                            v15 += genFn0()
                                                            _ = v34
                                                        }
                                                    case "":
                                                        if true {
                                                            var v35 uint32 = v4
                                                            var v36 float32 = 0.1
                                                            v27 += (v27 + (((v27 + float64(-2.25)) * (1e10 - ((3.14159 / v27) / (-2.25 * (1.5 + v27))))) + v27))
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        if (((!false) || (true || false)) || true) {
                                                            if false {
                                                                v5[(v23 - 2147483647)] = ("\n\t" > ("a" + "a"))
                                                            }
                                                            fmt.Println(v23, v15)
                                                            if true {
                                                                if (!(9223372036854775807 <= 1)) {
                                                                    var v37 int32 = -4
                                                                    var v38 int8 = 88
                                                                    var v39 int8 = (v38 % ((v38) | 1))
                                                                    _ = v37
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    var v40 int = fi25
                                                                    var v41 uint8 = v24
                                                                    v42 := GenS2{F0: 252}
                                                                    _ = v40
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v43 float32 = (v6 * v6)
                                                            _ = v43
                                                        } else {
                                                            var v44 int64 = v3
                                                            if (-1 != 44) {
                                                                if false {
                                                                    v24 = (v24 | (v24 >> ((363) & 63)))
                                                                    if (!false) {
                                                                        v45 := GenS1{F0: uint64(v24), F1: v44}
                                                                        var v46 int8 = 127
                                                                        _ = v45
                                                                        _ = v46
                                                                    } else {
                                                                        var v47 uint16 = 542
                                                                        _ = v47
                                                                    }
                                                                    var v48 float32 = 1.5
                                                                    var v49 uint = 65535
                                                                    _ = v48
                                                                    _ = v49
                                                                } else {
                                                                    v4 = (^(v4))
                                                                    var v50 string = ""
                                                                    v50 = v50
                                                                    switch 0 {
                                                                    case 42:
                                                                        fmt.Println(fi25, fi28)
                                                                    }
                                                                    switch (v2 - fi28) {
                                                                    case 673:
                                                                        var v51 int = -9223372036854775808
                                                                        _ = v51
                                                                    default:
                                                                        fmt.Println(v3)
                                                                    }
                                                                    v52 := GenS3{F0: 32767}
                                                                    if ((false && (0 == (v15 & genFn0()))) && (-2.25 <= v6)) {
                                                                        var v53 int16 = 943
                                                                        _ = v53
                                                                    }
                                                                    v27 += (((((v27 / 1.0) + v27) - 1e10) * 1.5) / v27)
                                                                    var v54 uint16 = 42
                                                                    switch ((v50 + "hello") + ((v50 + (v50 + "foo bar")) + "a")) {
                                                                    case "\n\t":
                                                                        if false {
                                                                            if (0 != 4294967295) {
                                                                                v1 = append(v1, (v4 ^ (v4 % (((v4 >> ((596) & 63))) | 1))))
                                                                                var v55 float64 = v27
                                                                                v56 := GenS3{F0: uint16(v4)}
                                                                                var v57 uint64 = 18446744073709551615
                                                                                if (v57 > v57) {
                                                                                    if true {
                                                                                        delete(v5, 32768)
                                                                                    }
                                                                                }
                                                                                _ = v55
                                                                                _ = v56
                                                                                _ = v57
                                                                            }
                                                                            if ((("\n\t" + v50) <= "a") || ((!false) && false)) {
                                                                                if false {
                                                                                    fmt.Println(v27, v6, v3)
                                                                                    fmt.Println(v27, v24)
                                                                                } else {
                                                                                    if true {
                                                                                        switch (v2 & (fi28 + fi28)) {
                                                                                        case -9223372036854775808:
                                                                                            var v58 int64 = v15
                                                                                            _ = v58
                                                                                        case -569:
                                                                                            v59 := GenS4{F0: 255, F1: (501 <= 0), F2: v24}
                                                                                            _ = v59
                                                                                        default:
                                                                                            v6 -= (v6 / v6)
                                                                                        }
                                                                                    }
                                                                                    v60 := GenS1{F0: 42, F1: genFn0()}
                                                                                    switch ((((((v50 + "\n\t") + v50) + (v50 + "")) + v50) + v50) + v50) {
                                                                                    case "":
                                                                                        v6 = v6
                                                                                    case "hello":
                                                                                        switch v50 {
                                                                                        case "hello":
                                                                                            var v61 string = (v50 + v50)
                                                                                            _ = v61
                                                                                        case "\n\t":
                                                                                            fmt.Println(v15)
                                                                                        }
                                                                                    default:
                                                                                        var v62 int32 = 0
                                                                                        _ = v62
                                                                                    }
                                                                                    _ = v60
                                                                                }
                                                                                var v63 uint64 = 0
                                                                                _ = v63
                                                                            }
                                                                            var v64 float64 = 1e10
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    _ = v50
                                                                    _ = v52
                                                                    _ = v54
                                                                }
                                                                v65 := GenS1{F0: 4294967295, F1: v3}
                                                                var v66 uint16 = 255
                                                                if false {
                                                                    var v67 string = "foo bar"
                                                                    v4 = v4
                                                                    switch -1 {
                                                                    case 127:
                                                                        var v68 int = (fi25 | -1)
                                                                        _ = v68
                                                                    case 1:
                                                                        v2 += (fi25 | (-(fi28)))
                                                                    case -664:
                                                                        v66 -= 729
                                                                    }
                                                                    var v69 int = int(v15)
                                                                    v23 += v2
                                                                    v15 += int64(v24)
                                                                    var v70 uint64 = 18446744073709551615
                                                                    _ = v67
                                                                    _ = v69
                                                                    _ = v70
                                                                } else {
                                                                    if false {
                                                                        var v71 uint = uint(v4)
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            var v72 float32 = ((v6 / v6) + float32(v27))
                                                                            _ = v72
                                                                        case "a":
                                                                            v73 := GenS5{F0: (^(v66)), F1: 100}
                                                                            _ = v73
                                                                        default:
                                                                            _ = len(v5)
                                                                        }
                                                                        var v74 uint16 = v66
                                                                        v75 := GenS4{F0: (v71 << ((v71) & 63)), F1: true, F2: v24}
                                                                        v76 := v5[(fi28 * 436)]
                                                                        _ = v71
                                                                        _ = v74
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                    if (true || ((!false) || false)) {
                                                                        if (0 >= 924) {
                                                                            var v77 uint64 = 4294967295
                                                                            var v78 int32 = -2147483648
                                                                            var v79 int32 = v78
                                                                            _ = v77
                                                                            _ = v78
                                                                            _ = v79
                                                                        } else {
                                                                            v23 += (fi28 / ((308) | 1))
                                                                        }
                                                                        if true {
                                                                            v3 -= v15
                                                                        }
                                                                        v2 -= (v23 + (v2 | (fi28 >> ((65535) & 63))))
                                                                        delete(v5, -9223372036854775808)
                                                                        var v80 float64 = v27
                                                                        var v81 uint64 = 42
                                                                        var v82 int8 = 42
                                                                        v80 = ((0.1 - 1.5) * v80)
                                                                        _ = v80
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                    v83 := GenS4{F0: 0, F1: true, F2: 120}
                                                                    _ = v83
                                                                }
                                                                fmt.Println(v44)
                                                                if ((v3 >> ((915) & 63)) == v44) {
                                                                    var v84 uint32 = (v4 & (v4 + 1))
                                                                    v4 -= (v84 >> ((2) & 63))
                                                                    _ = v84
                                                                }
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                v4 = v4
                                                                var v85 int = v2
                                                                _ = v85
                                                            }
                                                            fmt.Println(v3)
                                                            _ = v44
                                                        }
                                                    case "a":
                                                        if (!false) {
                                                            var v86 int8 = 103
                                                            if (true || (!(v4 == v4))) {
                                                                var v87 string = ("" + (("" + "hello") + "foo bar"))
                                                                _ = v87
                                                            } else {
                                                                var v88 int32 = 457
                                                                if true {
                                                                    v1 = append(v1, 2147483648)
                                                                    var v89 uint = 42
                                                                    if false {
                                                                        v4 = 0
                                                                        v90 := GenS6{F0: fi25, F1: -1, F2: 255}
                                                                        _ = v90
                                                                    }
                                                                    _ = v89
                                                                } else {
                                                                    _ = len(v1)
                                                                }
                                                                switch v23 {
                                                                case 1:
                                                                    switch ("hello" + "\n\t") {
                                                                    case "\n\t":
                                                                        if true {
                                                                            if (((1000 <= 1) || ("" < "")) || false) {
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                case "":
                                                                                    var v91 string = "foo bar"
                                                                                    _ = v91
                                                                                }
                                                                                v92 := GenS7{F0: (v86 % ((v86) | 1)), F1: 200}
                                                                                _ = v92
                                                                            }
                                                                        } else {
                                                                            var v93 uint16 = 546
                                                                            var v94 float32 = -2.25
                                                                            _ = v93
                                                                            _ = v94
                                                                        }
                                                                        v95 := GenS3{F0: 32767}
                                                                        v96 := GenS8{F0: int64(v4), F1: 65535, F2: ("hello" + "\n\t")}
                                                                        v86 += -124
                                                                        fmt.Println(v88, v2, v15)
                                                                        _ = v95
                                                                        _ = v96
                                                                    case "a":
                                                                        switch v23 {
                                                                        case 877:
                                                                            if false {
                                                                                switch ("" + "hello") {
                                                                                case "hello":
                                                                                    switch ("a" + ("hello" + "hello")) {
                                                                                    case "foo bar":
                                                                                        var v97 uint64 = 659
                                                                                        _ = v97
                                                                                    case "hello":
                                                                                        v98 := GenS7{F0: -12, F1: 823}
                                                                                        _ = v98
                                                                                    default:
                                                                                        var v99 int16 = int16(v88)
                                                                                        _ = v99
                                                                                    }
                                                                                default:
                                                                                    v4 = 2147483648
                                                                                }
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    var v100 uint = 65535
                                                                                    _ = v100
                                                                                case "\n\t":
                                                                                    var v101 int64 = v15
                                                                                    _ = v101
                                                                                default:
                                                                                    if true {
                                                                                        v102 := GenS4{F0: 42, F1: (v3 >= (v15 * (^(v15)))), F2: 232}
                                                                                        _ = v102
                                                                                    } else {
                                                                                        var v103 float32 = (1e10 + (v6 + v6))
                                                                                        v27 = 1e10
                                                                                        _ = v103
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v15 += (v3 ^ -9223372036854775808)
                                                                                delete(v5, (fi28 + -418))
                                                                                if false {
                                                                                    v104 := GenS9{F0: 65535, F1: v88, F2: v15}
                                                                                    v105 := GenS9{F0: 0, F1: (v88 % ((v88) | 1)), F2: (v15 / ((v15) | 1))}
                                                                                    var v106 int32 = v88
                                                                                    _ = v104
                                                                                    _ = v105
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(v3, v88, v2)
                                                                        }
                                                                    case "foo bar":
                                                                        if false {
                                                                            switch "a" {
                                                                            case "hello":
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        v23 = (v23 - (fi25 & v2))
                                                                                    }
                                                                                case "hello":
                                                                                    v107 := GenS10{F0: (^(v4)), F1: "\n\t", F2: (-(v88))}
                                                                                    v24 = (v24 << ((0) & 63))
                                                                                    _ = len(v5)
                                                                                    _ = v107
                                                                                case "":
                                                                                    switch ("foo bar" + "foo bar") {
                                                                                    case "hello":
                                                                                        if false {
                                                                                            v6 -= v6
                                                                                            fmt.Println(v23)
                                                                                            v23 = 127
                                                                                            v108 := GenS7{F0: int8(v24), F1: 2}
                                                                                            v109 := GenS4{F0: 255, F1: (((v88 & v88) < (v88 * (-(v88)))) || (1e10 < v6)), F2: v24}
                                                                                            v27 += (v27 / v27)
                                                                                            _ = v108
                                                                                            _ = v109
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v110 uint64 = 1
                                                                                        v5[v23] = (2 <= 65535)
                                                                                        _ = v110
                                                                                    case "\n\t":
                                                                                        var v111 int32 = v88
                                                                                        _ = v111
                                                                                    default:
                                                                                        var v112 int64 = -1
                                                                                        _ = v112
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                case 575:
                                                                    if true {
                                                                        if (4294967295 > 204) {
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                if true {
                                                                                    if (false && ((((-2.25 * -2.25) / 3.14159) == v27) || false)) {
                                                                                        v113 := GenS3{F0: 670}
                                                                                        if true {
                                                                                            v114 := GenS11{F0: ("hello" + "a")}
                                                                                            _ = v114
                                                                                        }
                                                                                        v115 := GenS12{F0: (v4 + v4), F1: fi28, F2: (^(v4))}
                                                                                        var v116 int32 = int32(v24)
                                                                                        v117 := GenS13{F0: 65535}
                                                                                        _ = v113
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                        _ = v117
                                                                                    } else {
                                                                                        var v118 uint8 = (v24 ^ (v24 & v24))
                                                                                        _ = v118
                                                                                    }
                                                                                } else {
                                                                                    var v119 uint16 = 65535
                                                                                    var v120 uint32 = v4
                                                                                    v121 := GenS7{F0: (v86 >> ((18446744073709551615) & 63)), F1: 18446744073709551615}
                                                                                    switch (fi28 + fi28) {
                                                                                    case 42:
                                                                                        v119 += 255
                                                                                    case 9223372036854775807:
                                                                                        var v122 bool = (0.0 > (v6 * float32(-2.25)))
                                                                                        _ = v122
                                                                                    }
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                    _ = v121
                                                                                }
                                                                            case "":
                                                                                v123 := GenS7{F0: (v86 | 78), F1: 837}
                                                                                _ = v123
                                                                            }
                                                                            _ = len(v1)
                                                                            v124 := GenS1{F0: 2, F1: genFn0()}
                                                                            _ = v124
                                                                        } else {
                                                                            var v125 int16 = 100
                                                                            _ = v125
                                                                        }
                                                                        v3 = v15
                                                                    } else {
                                                                        var v126 int8 = 45
                                                                        var v127 string = "\n\t"
                                                                        _ = v126
                                                                        _ = v127
                                                                    }
                                                                    if false {
                                                                        v128 := GenS14{F0: 721, F1: 2, F2: (v4 * 188)}
                                                                        fmt.Println(v6)
                                                                        if false {
                                                                            var v129 int = (-(fi25))
                                                                            _ = len(v1)
                                                                            v86 += v86
                                                                            _ = v129
                                                                        }
                                                                        _ = v128
                                                                    } else {
                                                                        v130 := GenS15{F0: 7}
                                                                        _ = v130
                                                                    }
                                                                }
                                                                if ((-(v86)) > (v86 - -128)) {
                                                                    v6 = 3.14159
                                                                    _ = len(v1)
                                                                }
                                                                var v131 float32 = (((3.14159 / (float32(((0.0 * v27) + 1e10)) - 1e10)) * v6) / v6)
                                                                v3 += int64(v4)
                                                                fmt.Println(fi25)
                                                                _ = v131
                                                                _ = v88
                                                            }
                                                            v132 := GenS12{F0: v4, F1: v2, F2: 454}
                                                            _ = v132
                                                            _ = v86
                                                        } else {
                                                            v3 = (^(v3))
                                                        }
                                                    default:
                                                        fmt.Println(v27, v2)
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = v4
                                                    }
                                                    _ = len(v5)
                                                }
                                                v27 = (float64(((((-2.25 / (3.14159 + 1e10)) / v6) * v6) + v6)) + v27)
                                                v5[(v23 + 670)] = (!true)
                                            }
                                        }
                                        v133 := GenS16{F0: v23, F1: (v23 << ((1) & 63)), F2: true}
                                        if true {
                                            switch "hello" {
                                            case "hello":
                                                v4 -= (^(v4))
                                            case "\n\t":
                                                v4 -= (v4 / ((v4) | 1))
                                            case "a":
                                                var v134 uint8 = v24
                                                switch "a" {
                                                case "":
                                                    v23 -= 346
                                                case "foo bar":
                                                    if false {
                                                        delete(v5, fi25)
                                                        var v135 int32 = 747
                                                        _ = v135
                                                    }
                                                }
                                                if false {
                                                    switch (fi25 * 839) {
                                                    case 4294967295:
                                                        if (28 == 2) {
                                                            fmt.Println(fi28, v2, v6)
                                                            fmt.Println(v23, v23)
                                                            var v136 int16 = 42
                                                            _ = v136
                                                        } else {
                                                            v2 = 799
                                                        }
                                                    }
                                                    v15 += (v3 << ((608) & 63))
                                                    _ = len(v5)
                                                }
                                                _ = v134
                                            }
                                            v1 = append(v1, 42)
                                            v137 := GenS17{F0: (v24 >> ((284) & 63)), F1: 2, F2: (-(v2))}
                                            _ = v137
                                        }
                                        v23 -= v23
                                        var v138 uint16 = 2
                                        _ = v133
                                        _ = v138
                                    }
                                    switch "a" {
                                    case "hello":
                                        if false {
                                            var v139 float64 = ((v27 - (v27 + (v27 / ((v27 - v27) / v27)))) + float64(3.14159))
                                            _ = v139
                                        }
                                    default:
                                        v23 -= (v23 << ((234) & 63))
                                    }
                                    if false {
                                        v6 = 0.1
                                        v140 := GenS18{F0: 1, F1: "", F2: true}
                                        v27 += v27
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v140
                                    } else {
                                        var v141 uint16 = 845
                                        _ = v141
                                    }
                                    _ = fi28
                                }
                                v142 := GenS2{F0: 42}
                                fmt.Println(v15)
                                _ = v142
                            }
                            _ = v27
                        } else {
                            switch (("" + "") + (("\n\t" + (("a" + (("" + "foo bar") + "hello")) + ("foo bar" + "foo bar"))) + "foo bar")) {
                            case "a":
                                if (false && ((v24 * 42) >= v24)) {
                                    v5[32767] = (53 >= 65535)
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        fmt.Println(v2, v24, v2)
                                    case "\n\t":
                                        v23 -= v23
                                    default:
                                        var v143 uint8 = (^(v24))
                                        _ = v143
                                    }
                                    if true {
                                        if (fi25 > (v23 & -759)) {
                                            var v144 float32 = v6
                                            if (v3 > (v3 | (^(v15)))) {
                                                for fi145 := 0; fi145 < 6; fi145++ {
                                                    if (false || true) {
                                                        if (false || false) {
                                                            var v146 int16 = 399
                                                            v6 -= v6
                                                            v1 = append(v1, (v4 | (v4 << ((42) & 63))))
                                                            _ = v146
                                                        } else {
                                                            if (32767 == -491) {
                                                                var v147 int8 = -128
                                                                fmt.Println(v24, v4, v23)
                                                                _ = v147
                                                            } else {
                                                                if true {
                                                                    var v148 uint32 = v4
                                                                    v149 := GenS4{F0: 1, F1: true, F2: v24}
                                                                    var v150 int64 = genFn0()
                                                                    var v151 string = "\n\t"
                                                                    _ = v148
                                                                    _ = v149
                                                                    _ = v150
                                                                    _ = v151
                                                                }
                                                            }
                                                            var v152 uint32 = 65535
                                                            _ = v152
                                                        }
                                                    }
                                                    if true {
                                                        _ = len(v5)
                                                    }
                                                    fmt.Println(fi25)
                                                    _ = fi145
                                                }
                                                var v153 bool = true
                                                if (1 > 0) {
                                                    _ = len(v1)
                                                }
                                                _ = v153
                                            }
                                            _ = v144
                                        }
                                        fmt.Println(v2, v15, v24)
                                        v6 = v6
                                    } else {
                                        if false {
                                            v4 = (v4 * v4)
                                        } else {
                                            var v154 uint32 = v4
                                            _ = v154
                                        }
                                        var v155 uint16 = 65535
                                        _ = v155
                                    }
                                } else {
                                    for fi156 := 0; fi156 < 5; fi156++ {
                                        v157 := GenS19{F0: 896, F1: false, F2: -519}
                                        if true {
                                            fmt.Println(fi156, v4)
                                            var v158 uint16 = 32767
                                            _ = v158
                                        }
                                        _ = fi156
                                        _ = v157
                                    }
                                    v4 -= v4
                                    v159 := v5[(fi25 * -935)]
                                    _ = v159
                                }
                            case "\n\t":
                                v23 = fi25
                            case "foo bar":
                                for fi160 := 0; fi160 < 4; fi160++ {
                                    v161 := GenS19{F0: (-(v23)), F1: ((0 > 65535) && true), F2: -9223372036854775808}
                                    fmt.Println(v3)
                                    v15 += (-(v15))
                                    _ = fi160
                                    _ = v161
                                }
                            default:
                                var v162 uint32 = 358
                                _ = v162
                            }
                            v24 -= 176
                            v163 := GenS7{F0: 1, F1: 1}
                            _ = v163
                        }
                    case 55:
                        if false {
                            fmt.Println(v6)
                        }
                    default:
                        var v164 uint8 = (v24 / (((v24 | v24)) | 1))
                        _ = v164
                    }
                    _ = fi25
                }
            } else {
                if false {
                    var v165 uint16 = 255
                    var v166 string = ("\n\t" + "foo bar")
                    v6 = -2.25
                    _ = v165
                    _ = v166
                } else {
                    v167 := GenS20{F0: int16(v23), F1: v24}
                    v5[v2] = (!false)
                    v15 += -646
                    _ = v167
                }
                var v168 int64 = genFn0()
                var v169 int16 = -32768
                _ = v168
                _ = v169
            }
            _ = v24
        }
        _ = v22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
