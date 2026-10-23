package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int64; F1 int16; F2 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 uint8; F1 uint16 }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64; F1 bool; F2 string }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint8; F1 int8 }
type GenS15 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS16 struct { F0 bool; F1 uint }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 bool }
func genFn0(p0 uint8, p1 string, p2 int16) int8 {
    v5 := GenS1{F0: 1, F1: (p1 + (p1 + p1)), F2: 0}
    v6 := GenS2{F0: 682, F1: 1, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint64) bool {
    var v7 uint8 = 2
    p0 = 18446744073709551615
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(p0, fi8)
        p0 -= (^(p0))
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS3{F0: 0, F1: 1}
            if true {
                var v11 string = "a"
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v12 int = 1
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            v7 = v7
            _ = fi14
        }
        if true {
            var v15 bool = false
            v16 := []int{9223372036854775807, fi13, v12, (v12 >> ((471) & 63))}
            for fi17 := 0; fi17 < 14; fi17++ {
                v16 = append(v16, v12)
                v18 := GenS1{F0: -1, F1: ("hello" + "foo bar"), F2: 42}
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
        p0 = 4294967295
        _ = fi13
    }
    _ = p0
    _ = v12
    _ = v7
    return false
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1000
    var v4 uint32 = 562
    fmt.Println(v3)
    for fi19 := 0; fi19 < 11; fi19++ {
        v3 += (v3 / (((v3 & v3)) | 1))
        var v20 float32 = float32(((((0.1 / (1e10 + 3.14159)) * float64(3.14159)) + -2.25) + (1.5 - -2.25)))
        var v21 uint64 = 1
        _ = fi19
        _ = v20
        _ = v21
    }
    switch 4294967295 {
    case 32768:
        var v22 uint32 = v4
        _ = v22
    case -343:
        v4 = (v4 >> ((0) & 63))
    default:
        var v23 bool = false
        _ = v23
    }
    if (v3 == v3) {
        v2[(v4 ^ 514)] = (v4 << ((18446744073709551615) & 63))
        switch -1 {
        case 287:
            if (((v4 + 65535) < v4) || false) {
                for fi24 := 0; fi24 < 5; fi24++ {
                    v25 := GenS3{F0: 61, F1: 704}
                    _ = fi24
                    _ = v25
                }
            } else {
                fmt.Println(v3)
            }
        case 9223372036854775807:
            for fi26 := 0; fi26 < 7; fi26++ {
                if len(v1) > 0 {
                    v1[0] = v4
                }
                v3 -= 325
                v2[v4] = 2147483647
                _ = fi26
            }
        default:
            v27 := GenS4{F0: -500, F1: 258}
            _ = v27
        }
        fmt.Println(v3, v4)
    } else {
        if false {
            _ = len(v2)
            v2[(v4 ^ (v4 >> ((65535) & 63)))] = v4
            delete(v2, (v4 >> ((65535) & 63)))
        } else {
            for fi28 := 0; fi28 < 4; fi28++ {
                v1 = append(v1, v4)
                fmt.Println(fi28)
                _ = fi28
            }
        }
        v4 += (v4 & (v4 + v4))
        switch (("" + (("" + "foo bar") + "foo bar")) + "hello") {
        case "\n\t":
            if (false && false) {
                switch 4294967295 {
                case -628:
                    if (((v4 >> ((697) & 63)) <= v4) || true) {
                        if (v3 > v3) {
                            v29 := GenS1{F0: 1, F1: ("\n\t" + ""), F2: -1}
                            for fi30 := 0; fi30 < 9; fi30++ {
                                for fi31 := 0; fi31 < 20; fi31++ {
                                    v4 = v4
                                    _ = fi31
                                }
                                v3 -= (v3 >> ((42) & 63))
                                if (!((v3 | v3) == v3)) {
                                    if true {
                                        for fi32 := 0; fi32 < 20; fi32++ {
                                            v3 -= 42
                                            if true {
                                                delete(v2, v4)
                                            }
                                            fmt.Println(v4, v3)
                                            _ = fi32
                                        }
                                    }
                                } else {
                                    for fi33 := 0; fi33 < 15; fi33++ {
                                        if true {
                                            v34 := GenS4{F0: 32768, F1: 4294967295}
                                            var v35 uint64 = 4294967295
                                            switch fi33 {
                                            case 613:
                                                fmt.Println(v3, v4, fi30)
                                            case 32768:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if true {
                                                v3 += (v3 ^ v3)
                                                var v36 uint = 18446744073709551615
                                                _ = v36
                                            }
                                            v37 := GenS4{F0: 128, F1: 2}
                                            v38 := GenS5{F0: (v4 | (v4 * (v4 & v4))), F1: (42 > 60)}
                                            switch (("" + "") + ("\n\t" + ("\n\t" + "\n\t"))) {
                                            case "":
                                                v39 := GenS6{F0: 0, F1: v3, F2: 1}
                                                _ = v39
                                            }
                                            v40 := GenS7{F0: -1}
                                            if false {
                                                if (!(!false)) {
                                                    switch -9223372036854775808 {
                                                    case 128:
                                                        v2[(v4 | v4)] = (v4 / ((452) | 1))
                                                    case -408:
                                                        var v41 bool = (-2147483648 <= int32(fi30))
                                                        _ = v41
                                                    case -129:
                                                        switch "a" {
                                                        case "":
                                                            if ((v3 + -1) != v3) {
                                                                switch 0 {
                                                                case 2147483647:
                                                                    var v42 float64 = 3.14159
                                                                    _ = v42
                                                                case 986:
                                                                    fmt.Println(fi33)
                                                                case 1:
                                                                    switch (-(fi30)) {
                                                                    case -912:
                                                                        var v43 uint64 = 272
                                                                        _ = v43
                                                                    case -1:
                                                                        var v44 int8 = 100
                                                                        _ = v44
                                                                    case 0:
                                                                        var v45 uint16 = 65535
                                                                        _ = v45
                                                                    }
                                                                }
                                                            } else {
                                                                var v46 uint8 = 128
                                                                _ = v46
                                                            }
                                                        }
                                                    }
                                                }
                                                var v47 uint8 = 128
                                                _ = v47
                                            }
                                            if false {
                                                var v48 uint64 = 1
                                                if true {
                                                    if true {
                                                        var v49 uint8 = 78
                                                        v50 := GenS3{F0: genFn0(v49, ("foo bar" + ""), v3), F1: int64(fi30)}
                                                        if ((fi30 + (fi30 >> ((42) & 63))) <= (fi30 % (((-(fi30))) | 1))) {
                                                            v3 = (v3 % ((v3) | 1))
                                                            switch fi30 {
                                                            case 32768:
                                                                var v51 int8 = -82
                                                                _ = v51
                                                            case 440:
                                                                fmt.Println(v4)
                                                            }
                                                            var v52 int8 = 0
                                                            _ = v52
                                                        } else {
                                                            v53 := GenS2{F0: -262, F1: 245, F2: uint16(v48)}
                                                            delete(v2, (v4 / ((v4) | 1)))
                                                            _ = v53
                                                        }
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        v3 -= 42
                                                        v54 := GenS1{F0: 111, F1: "\n\t", F2: genFn0(128, "", v3)}
                                                        if (9223372036854775807 != 1) {
                                                            v48 = v48
                                                        } else {
                                                            v55 := GenS4{F0: 649, F1: (v48 >> ((4294967295) & 63))}
                                                            _ = v55
                                                        }
                                                        _ = v54
                                                    }
                                                } else {
                                                    var v56 int16 = 32767
                                                    _ = v56
                                                }
                                                v48 += v48
                                                _ = v48
                                            } else {
                                                v4 -= v4
                                                v3 = (-(v3))
                                                v57 := GenS8{F0: (v3 ^ -32768), F1: genFn0(201, "foo bar", (v3 * v3))}
                                                if false {
                                                    v58 := GenS9{F0: 9, F1: 524}
                                                    var v59 uint64 = 42
                                                    var v60 int = fi33
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    var v61 bool = (0 <= 32767)
                                                    switch fi33 {
                                                    case 127:
                                                        if ((float32(-2.25) > 0.1) && v61) {
                                                            if len(v1) > 0 {
                                                                v1[0] = (v4 / (((v4 << ((1) & 63))) | 1))
                                                            }
                                                            _ = len(v1)
                                                            v62 := GenS3{F0: 67, F1: 832}
                                                            _ = v62
                                                        } else {
                                                            var v63 int = fi30
                                                            v64 := GenS10{F0: 9223372036854775807, F1: 1}
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                    }
                                                    v65 := v2[4294967295]
                                                    _ = v61
                                                    _ = v65
                                                }
                                                _ = v57
                                            }
                                            switch fi30 {
                                            case -25:
                                                v3 = v3
                                            default:
                                                fmt.Println(v4, v3, v3)
                                            }
                                            if false {
                                                if false {
                                                    if (true && true) {
                                                        var v66 float32 = float32((0.0 / 3.14159))
                                                        fmt.Println(v4, fi30)
                                                        v66 = v66
                                                        _ = v66
                                                    } else {
                                                        fmt.Println(fi30)
                                                        fmt.Println(fi30)
                                                        v4 = v4
                                                    }
                                                    v4 = (v4 * v4)
                                                    v4 -= 1
                                                    if ((-21 >= 1) || true) {
                                                        if (1 <= 4294967295) {
                                                            if (3.14159 < ((1.5 / 1e10) - (float64(-2.25) + 1.5))) {
                                                                var v67 int32 = 42
                                                                delete(v2, (v4 - v4))
                                                                delete(v2, v4)
                                                                v4 -= v4
                                                                if true {
                                                                    fmt.Println(v67, v67, fi30)
                                                                    switch (^(fi30)) {
                                                                    case 57:
                                                                        if (925 == 65535) {
                                                                            var v68 int = 644
                                                                            _ = v68
                                                                        }
                                                                    case 46:
                                                                        fmt.Println(fi30)
                                                                    }
                                                                    fmt.Println(v67)
                                                                } else {
                                                                    v2[(v4 / (((v4 - v4)) | 1))] = (v4 | 65535)
                                                                    var v69 int16 = (v3 * 1)
                                                                    if (!true) {
                                                                        _ = len(v1)
                                                                    } else {
                                                                        var v70 uint8 = 42
                                                                        fmt.Println(v70, v70, v67)
                                                                        switch fi33 {
                                                                        case 496:
                                                                            switch ("hello" + "a") {
                                                                            case "foo bar":
                                                                                if (false && ((("a" + ("" + "\n\t")) + "foo bar") == ("hello" + "\n\t"))) {
                                                                                    v70 += (v70 & v70)
                                                                                }
                                                                            case "a":
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    var v71 int16 = -270
                                                                                    _ = v71
                                                                                }
                                                                            default:
                                                                                var v72 uint64 = 18446744073709551615
                                                                                _ = v72
                                                                            }
                                                                        case 127:
                                                                            v73 := GenS11{F0: false}
                                                                            _ = v73
                                                                        case 128:
                                                                            switch (fi30 << ((0) & 63)) {
                                                                            case 131:
                                                                                if false {
                                                                                    v3 -= v3
                                                                                }
                                                                            default:
                                                                                var v74 bool = true
                                                                                _ = v74
                                                                            }
                                                                        }
                                                                        _ = v70
                                                                    }
                                                                    _ = v69
                                                                }
                                                                _ = v67
                                                            } else {
                                                                v4 += v4
                                                                fmt.Println(v4, v4)
                                                            }
                                                        }
                                                        v4 -= 451
                                                    }
                                                    if true {
                                                        if (false || (!(!false))) {
                                                            var v75 int64 = -1
                                                            if len(v1) > 0 {
                                                                v1[0] = (v4 + (v4 ^ (v4 ^ v4)))
                                                            }
                                                            _ = v75
                                                        }
                                                        v3 = (v3 + v3)
                                                        var v76 string = "a"
                                                        _ = v76
                                                    }
                                                    var v77 int16 = v3
                                                    if len(v1) > 0 {
                                                        v1[0] = v4
                                                    }
                                                    v4 -= (v4 + v4)
                                                    _ = v77
                                                }
                                                if true {
                                                    v4 += 42
                                                    v2[(v4 ^ 2)] = v4
                                                }
                                            } else {
                                                v3 = (v3 | 283)
                                            }
                                            _ = v37
                                            _ = v38
                                            _ = v40
                                        }
                                        if false {
                                            var v78 uint8 = 43
                                            switch ("" + (("foo bar" + ("foo bar" + "foo bar")) + "hello")) {
                                            case "":
                                                var v79 int8 = 0
                                                _ = v79
                                            }
                                            _ = v78
                                        } else {
                                            var v80 int64 = -9223372036854775808
                                            _ = v80
                                        }
                                        v81 := GenS12{F0: 373, F1: true, F2: (("" + "a") + ("foo bar" + "a"))}
                                        var v82 int64 = 1
                                        _ = fi33
                                        _ = v81
                                        _ = v82
                                    }
                                    var v83 uint64 = 18446744073709551615
                                    v3 = v3
                                    _ = v83
                                }
                                _ = fi30
                            }
                            _ = v29
                        }
                        var v84 int64 = 9223372036854775807
                        _ = v84
                    } else {
                        for fi85 := 0; fi85 < 16; fi85++ {
                            var v86 uint = 255
                            switch 2147483647 {
                            case -1:
                                var v87 uint64 = 4294967295
                                _ = v87
                            case 1:
                                if false {
                                    switch "\n\t" {
                                    case "":
                                        var v88 uint64 = 0
                                        _ = v88
                                    case "\n\t":
                                        v89 := GenS13{F0: -1}
                                        _ = v89
                                    case "hello":
                                        for fi90 := 0; fi90 < 12; fi90++ {
                                            if (true && false) {
                                                switch (^(fi85)) {
                                                case 1:
                                                    var v91 string = ("a" + "hello")
                                                    _ = v91
                                                case 32768:
                                                    fmt.Println(v4)
                                                case 201:
                                                    var v92 uint8 = 2
                                                    _ = v92
                                                default:
                                                    v93 := GenS14{F0: 128, F1: 127}
                                                    _ = v93
                                                }
                                                var v94 string = "a"
                                                var v95 bool = (v86 >= (^(v86)))
                                                fmt.Println(fi90, v94, fi90)
                                                delete(v2, v4)
                                                _ = v94
                                                _ = v95
                                            }
                                            _ = fi90
                                        }
                                    default:
                                        fmt.Println(fi85)
                                    }
                                } else {
                                    var v96 uint8 = 42
                                    _ = v96
                                }
                            }
                            var v97 uint64 = 959
                            _ = fi85
                            _ = v86
                            _ = v97
                        }
                        v98 := GenS15{F0: true, F1: 2, F2: 0}
                        _ = v98
                    }
                case -703:
                    var v99 int64 = 0
                    _ = v99
                default:
                    v3 = v3
                }
                switch 42 {
                case 32767:
                    var v100 string = "a"
                    _ = v100
                default:
                    for fi101 := 0; fi101 < 20; fi101++ {
                        fmt.Println(v3, v3)
                        v4 = (v4 - v4)
                        _ = fi101
                    }
                }
                if ((((^(v4)) <= 1) || ("a" != "\n\t")) || true) {
                    v102 := GenS7{F0: 9223372036854775807}
                    for fi103 := 0; fi103 < 20; fi103++ {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v3, v3, v3)
                        default:
                            for fi104 := 0; fi104 < 9; fi104++ {
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                var v105 int8 = 42
                                _ = fi104
                                _ = v105
                            }
                        }
                        for fi106 := 0; fi106 < 4; fi106++ {
                            if true {
                                var v107 uint = 963
                                var v108 uint64 = 1
                                _ = v107
                                _ = v108
                            }
                            v3 -= (v3 - v3)
                            if (false || ((("" + "") + "foo bar") == (((("" + ("\n\t" + ((("\n\t" + "a") + "\n\t") + (("a" + (("" + "hello") + "")) + ((("\n\t" + ("\n\t" + "a")) + "hello") + ""))))) + ("foo bar" + "")) + "") + "hello"))) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v3 -= v3
                                var v109 float32 = 1e10
                                _ = v109
                            }
                            _ = fi106
                        }
                        _ = fi103
                    }
                    v3 += (v3 ^ v3)
                    _ = v102
                }
            }
        case "a":
            for fi110 := 0; fi110 < 6; fi110++ {
                switch "hello" {
                case "a":
                    if false {
                        fmt.Println(v4, v4)
                    } else {
                        if (127 >= 103) {
                            if true {
                                if true {
                                    var v111 int8 = 19
                                    v111 -= v111
                                    _ = v111
                                }
                                fmt.Println(v4)
                                if (!true) {
                                    if true {
                                        switch "hello" {
                                        case "hello":
                                            if true {
                                                switch "foo bar" {
                                                case "hello":
                                                    v4 -= 4294967295
                                                case "foo bar":
                                                    v4 += 2
                                                default:
                                                    v3 += int16(v4)
                                                }
                                                for fi112 := 0; fi112 < 14; fi112++ {
                                                    var v113 uint = 711
                                                    v4 += 42
                                                    switch "a" {
                                                    case "hello":
                                                        _ = len(v2)
                                                    case "a":
                                                        v114 := GenS2{F0: 42, F1: (v113 % ((v113) | 1)), F2: 32767}
                                                        _ = v114
                                                    default:
                                                        v4 = (v4 ^ v4)
                                                    }
                                                    _ = fi112
                                                    _ = v113
                                                }
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                    }
                                    if false {
                                        switch 42 {
                                        case -236:
                                            v4 += v4
                                        case 0:
                                            if ((278 >= -352) && false) {
                                                if false {
                                                    v115 := GenS16{F0: (((59 >= 128) && (0.0 <= ((0.0 / -2.25) - (float32((-2.25 + (float64((3.14159 - 1e10)) - -2.25))) * ((1e10 / 3.14159) / 0.1))))) && (0 != 17)), F1: 836}
                                                    _ = v115
                                                }
                                                if true {
                                                    if false {
                                                        var v116 uint64 = 303
                                                        if (255 <= 128) {
                                                            var v117 uint32 = 65535
                                                            var v118 int32 = 42
                                                            _ = v117
                                                            _ = v118
                                                        } else {
                                                            v2[(v4 | v4)] = v4
                                                            var v119 uint32 = (v4 >> ((4294967295) & 63))
                                                            _ = v119
                                                        }
                                                        if (!false) {
                                                            v120 := GenS9{F0: uint8(v4), F1: 65535}
                                                            v121 := GenS4{F0: fi110, F1: (v116 << ((0) & 63))}
                                                            delete(v2, v4)
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        _ = v116
                                                    } else {
                                                        v3 = v3
                                                        v3 -= v3
                                                    }
                                                    v122 := GenS13{F0: -2147483648}
                                                    var v123 int8 = 100
                                                    _ = v122
                                                    _ = v123
                                                } else {
                                                    switch 4294967295 {
                                                    case -74:
                                                        v124 := v2[43]
                                                        _ = v124
                                                    case -632:
                                                        v125 := GenS11{F0: (true || false)}
                                                        _ = v125
                                                    }
                                                    v3 += (v3 | 42)
                                                    _ = len(v2)
                                                }
                                                switch ((("" + "a") + "") + "hello") {
                                                case "\n\t":
                                                    v126 := GenS3{F0: -4, F1: -1}
                                                    _ = v126
                                                case "foo bar":
                                                    v3 = -1
                                                default:
                                                    if (-312 <= -9223372036854775808) {
                                                        v127 := GenS17{F0: 2}
                                                        if false {
                                                            var v128 bool = false
                                                            var v129 uint32 = v4
                                                            _ = v128
                                                            _ = v129
                                                        }
                                                        v130 := GenS7{F0: 42}
                                                        _ = v127
                                                        _ = v130
                                                    }
                                                }
                                            } else {
                                                var v131 int64 = -162
                                                var v132 int64 = (v131 * v131)
                                                _ = v131
                                                _ = v132
                                            }
                                        default:
                                            switch (-(fi110)) {
                                            case -44:
                                                v133 := GenS18{F0: v3}
                                                _ = v133
                                            case 128:
                                                if true {
                                                    if true {
                                                        var v134 uint32 = v4
                                                        switch (-(fi110)) {
                                                        case 0:
                                                            var v135 float64 = (-2.25 * (float64(3.14159) * 1e10))
                                                            _ = v135
                                                        case 32768:
                                                            for fi136 := 0; fi136 < 16; fi136++ {
                                                                switch ("foo bar" + "a") {
                                                                case "foo bar":
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                case "hello":
                                                                    if (true || false) {
                                                                        v137 := GenS12{F0: uint64(v4), F1: true, F2: ("a" + "hello")}
                                                                        v3 = (v3 / ((v3) | 1))
                                                                        _ = v137
                                                                    }
                                                                default:
                                                                    var v138 int8 = -1
                                                                    _ = v138
                                                                }
                                                                _ = fi136
                                                            }
                                                        default:
                                                            switch 46 {
                                                            case -752:
                                                                var v139 uint8 = 74
                                                                _ = v139
                                                            default:
                                                                v140 := GenS17{F0: 306}
                                                                _ = v140
                                                            }
                                                        }
                                                        _ = v134
                                                    } else {
                                                        var v141 int16 = (v3 + (v3 & -252))
                                                        var v142 float32 = 1.5
                                                        v4 += (v4 << ((0) & 63))
                                                        _ = v141
                                                        _ = v142
                                                    }
                                                    v4 = v4
                                                    for fi143 := 0; fi143 < 5; fi143++ {
                                                        var v144 int8 = 63
                                                        _ = fi143
                                                        _ = v144
                                                    }
                                                } else {
                                                    v4 = (v4 % ((v4) | 1))
                                                    v3 += (-(v3))
                                                }
                                            }
                                        }
                                    } else {
                                        var v145 uint = 0
                                        _ = v145
                                    }
                                    v3 -= (v3 / ((v3) | 1))
                                }
                            }
                        }
                        _ = len(v1)
                    }
                case "":
                    for fi146 := 0; fi146 < 9; fi146++ {
                        var v147 uint = 23
                        _ = fi146
                        _ = v147
                    }
                case "hello":
                    var v148 int64 = -1
                    _ = v148
                default:
                    for fi149 := 0; fi149 < 20; fi149++ {
                        switch fi149 {
                        case 32768:
                            v3 = (v3 >> ((18446744073709551615) & 63))
                        case -346:
                            if false {
                                var v150 int16 = v3
                                _ = v150
                            } else {
                                v4 += 2147483648
                            }
                        }
                        _ = fi149
                    }
                }
                _ = fi110
            }
        }
    }
    var v151 float64 = 0.1
    var v152 uint8 = 182
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("\n\t" + "") {
    case "hello":
        if (0 >= 255) {
            for fi153 := 0; fi153 < 5; fi153++ {
                if len(v1) > 0 {
                    v1[0] = v4
                }
                v152 = (v152 & (v152 >> ((65535) & 63)))
                v152 += v152
                _ = fi153
            }
            if len(v1) > 0 {
                v1[0] = 65535
            }
            for fi154 := 0; fi154 < 9; fi154++ {
                switch (fi154 % (((fi154 * (fi154 - -1))) | 1)) {
                case -975:
                    v155 := GenS16{F0: true, F1: 525}
                    _ = v155
                case -2147483648:
                    v4 = v4
                case 127:
                    fmt.Println(v151)
                default:
                    if false {
                        var v156 uint64 = 133
                        _ = v156
                    } else {
                        v157 := GenS2{F0: 1, F1: 939, F2: 37}
                        var v158 uint = 705
                        v159 := GenS14{F0: v152, F1: int8(v3)}
                        _ = v157
                        _ = v158
                        _ = v159
                    }
                }
                _ = len(v1)
                _ = fi154
            }
        }
    case "":
        v160 := GenS10{F0: 42, F1: -620}
        _ = v160
    default:
        for fi161 := 0; fi161 < 4; fi161++ {
            if false {
                var v162 int32 = -2147483648
                _ = v162
            }
            for fi163 := 0; fi163 < 10; fi163++ {
                v164 := GenS12{F0: 2, F1: false, F2: "foo bar"}
                switch (fi163 ^ 9223372036854775807) {
                case 2147483647:
                    v165 := v2[(v4 * (v4 % (((^(v4))) | 1)))]
                    _ = v165
                case 178:
                    v166 := GenS10{F0: -9223372036854775808, F1: 42}
                    var v167 float32 = (-2.25 - 0.1)
                    v152 -= (v152 | 255)
                    _ = v166
                    _ = v167
                default:
                    v168 := GenS19{F0: 1}
                    _ = v168
                }
                _ = fi163
                _ = v164
            }
            v169 := GenS11{F0: false}
            _ = fi161
            _ = v169
        }
    }
    for fi170 := 0; fi170 < 9; fi170++ {
        v151 += v151
        var v171 int64 = 9223372036854775807
        var v172 int16 = v3
        _ = fi170
        _ = v171
        _ = v172
    }
    if true {
        v152 -= v152
        if false {
            var v173 string = ("hello" + ("" + ((((("" + "a") + (("hello" + "foo bar") + (("" + "hello") + "\n\t"))) + (("hello" + "foo bar") + "a")) + "a") + "a")))
            _ = v173
        } else {
            _ = len(v1)
        }
        _ = len(v2)
    } else {
        fmt.Println(v151, v151)
        v2[(^(v4))] = v4
        fmt.Println(v152, v151, v3)
    }
    for fi174 := 0; fi174 < 3; fi174++ {
        for fi175 := 0; fi175 < 11; fi175++ {
            fmt.Println(v4, v4, v151)
            _ = fi175
        }
        _ = fi174
    }
    v4 = v4
    if (int64(v152) > -467) {
        v151 += float64(((((1e10 / 0.1) / -2.25) - 1.5) + 1.5))
    }
    if ((v4 / ((v4) | 1)) <= v4) {
        v176 := GenS20{F0: false}
        for fi177 := 0; fi177 < 13; fi177++ {
            var v178 int16 = v3
            _ = fi177
            _ = v178
        }
        v4 = v4
        _ = v176
    } else {
        v179 := GenS4{F0: 0, F1: 42}
        for fi180 := 0; fi180 < 12; fi180++ {
            for fi181 := 0; fi181 < 10; fi181++ {
                fmt.Println(v152)
                _ = fi181
            }
            v151 = (v151 - 3.14159)
            _ = fi180
        }
        for fi182 := 0; fi182 < 8; fi182++ {
            v152 = v152
            _ = fi182
        }
        _ = v179
    }
    for fi183 := 0; fi183 < 7; fi183++ {
        var v184 uint16 = 146
        _ = fi183
        _ = v184
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v4
}
