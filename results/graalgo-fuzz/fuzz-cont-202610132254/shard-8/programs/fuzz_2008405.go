package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int16; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 int64; F2 int8 }
type GenS8 struct { F0 int64; F1 uint; F2 int }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS11 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint16; F1 int16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int8; F1 int16 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int32; F1 string }
func genFn0(p0 uint8, p1 uint8, p2 int8) uint32 {
    if false {
        var v7 int8 = 1
        var v8 uint16 = 0
        var v9 int16 = 0
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: -205, F1: 648, F2: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 4294967295
}

func main() {
    var v1 int16 = 291
    var v2 float64 = (float64(0.1) / -2.25)
    var v3 uint32 = 42
    var v4 uint8 = 212
    var v5 uint16 = 1
    v6 := make(map[uint32]int)
    v1 += v1
    switch 32768 {
    case 888:
        v11 := []int{}
        _ = v11
    case 9223372036854775807:
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(v5)
                if (true && (v2 < 0.1)) {
                    v14 := []int{}
                    var v15 uint64 = 784
                    var v16 int16 = 32767
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    if ((-1 > 42) && false) {
                        v17 := GenS1{F0: int(v3), F1: 65535, F2: (fi13 + 0)}
                        v18 := []int{414, (fi13 << ((255) & 63)), fi12, (fi12 / ((42) | 1))}
                        v19 := GenS1{F0: (fi12 / ((-9223372036854775808) | 1)), F1: (v5 & v5), F2: fi12}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    v20 := []int{4294967295, fi12, -1, fi12}
                    if false {
                        v21 := GenS2{F0: fi12, F1: 0, F2: 1}
                        switch "a" {
                        case "a":
                            v1 = 32767
                        }
                        _ = v21
                    } else {
                        var v22 uint64 = 553
                        v23 := GenS1{F0: fi13, F1: uint16(v22), F2: (fi12 * fi13)}
                        var v24 int = 127
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                }
                _ = fi13
            }
            v25 := GenS3{F0: (^(v1)), F1: 211}
            _ = fi12
            _ = v25
        }
    case 630:
        if ((v3 | v3) != genFn0(v4, (v4 / ((v4) | 1)), 84)) {
            var v26 float32 = 3.14159
            if false {
                v1 -= v1
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint32 = v3
                    _ = fi27
                    _ = v28
                }
                v5 -= v5
            } else {
                _ = len(v6)
                if ("hello" != ("" + "")) {
                    if false {
                        var v29 int8 = -1
                        v30 := []int{0, -2147483648}
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS4{F0: (!(0 > -128))}
                    for fi32 := 0; fi32 < 16; fi32++ {
                        var v33 int16 = 964
                        _ = fi32
                        _ = v33
                    }
                    _ = v31
                }
            }
            _ = v26
        }
    }
    if false {
        v6[genFn0((v4 / ((2) | 1)), uint8(v3), -63)] = -2147483648
        v3 = (v3 - (v3 << ((18446744073709551615) & 63)))
    }
    switch 128 {
    case 4294967295:
        delete(v6, 2)
    case 127:
        switch "" {
        case "":
            if true {
                v2 += float64(0.1)
                v3 += v3
                v34 := []int{2147483647, 128}
                _ = v34
            } else {
                for fi35 := 0; fi35 < 2; fi35++ {
                    v1 = 1000
                    var v36 int64 = 0
                    fmt.Println(v4, v1)
                    _ = fi35
                    _ = v36
                }
            }
        default:
            v37 := GenS4{F0: (float32(0.1) >= (-2.25 * (-2.25 - (0.0 * 0.0))))}
            _ = v37
        }
    default:
        v38 := []int{int(v5), -250, 0, int(v1)}
        _ = v38
    }
    if (v1 == int16(v5)) {
        for fi39 := 0; fi39 < 4; fi39++ {
            if (255 < 42) {
                var v40 float32 = ((0.1 * 0.0) / (0.1 + 0.1))
                var v41 int64 = 1
                _ = v40
                _ = v41
            }
            _ = fi39
        }
    } else {
        if (genFn0((v4 + (^(v4))), v4, -1) > genFn0((^(v4)), (v4 * (v4 ^ 0)), 54)) {
            v1 = (v1 % (((v1 << ((65535) & 63))) | 1))
        } else {
            var v42 uint32 = v3
            v5 += (v5 | (v5 >> ((474) & 63)))
            _ = v42
        }
    }
    if true {
        for fi43 := 0; fi43 < 4; fi43++ {
            for fi44 := 0; fi44 < 15; fi44++ {
                v45 := GenS5{F0: 540, F1: (v1 - (v1 * v1)), F2: v1}
                _ = fi44
                _ = v45
            }
            v3 += (v3 - (v3 * genFn0(42, (^(v4)), -128)))
            v4 = (v4 + (^(v4)))
            _ = fi43
        }
        for fi46 := 0; fi46 < 8; fi46++ {
            fmt.Println(v4, fi46)
            v3 += (v3 & v3)
            if (!(0 < -973)) {
                if (-1 == 42) {
                    switch (^(fi46)) {
                    case 9223372036854775807:
                        switch fi46 {
                        case -1:
                            for fi47 := 0; fi47 < 16; fi47++ {
                                switch "\n\t" {
                                case "foo bar":
                                    v48 := GenS6{F0: 794}
                                    _ = v48
                                case "a":
                                    v2 += 1.5
                                }
                                v5 -= 32767
                                var v49 bool = (!(("" + (((("" + ("foo bar" + "a")) + "") + ("" + (("hello" + "\n\t") + ("hello" + "hello")))) + "a")) != "foo bar"))
                                var v50 int64 = 0
                                fmt.Println(v4, v5)
                                _ = fi47
                                _ = v49
                                _ = v50
                            }
                        case 9223372036854775807:
                            for fi51 := 0; fi51 < 8; fi51++ {
                                v1 = (v1 % (((v1 | v1)) | 1))
                                if ("a" == "a") {
                                    var v52 bool = true
                                    v53 := []int{fi51, fi51, fi46}
                                    _ = v52
                                    _ = v53
                                } else {
                                    v2 = v2
                                }
                                v54 := v6[65535]
                                _ = fi51
                                _ = v54
                            }
                        default:
                            switch fi46 {
                            case 9223372036854775807:
                                for fi55 := 0; fi55 < 20; fi55++ {
                                    var v56 uint32 = v3
                                    if true {
                                        v56 += (^(v56))
                                        fmt.Println(v3, v2)
                                    } else {
                                        var v57 uint64 = 42
                                        fmt.Println(fi55, v57)
                                        _ = v57
                                    }
                                    var v58 uint16 = v5
                                    fmt.Println(fi55, v5, fi55)
                                    switch "foo bar" {
                                    case "\n\t":
                                        if true {
                                            if (false || true) {
                                                var v59 bool = (!((!(!false)) && (0 <= 65535)))
                                                _ = v59
                                            }
                                        }
                                        switch (fi46 + -266) {
                                        case -560:
                                            delete(v6, v56)
                                        }
                                        var v60 int8 = 42
                                        _ = v60
                                    default:
                                        v1 = (v1 % (((v1 ^ v1)) | 1))
                                    }
                                    v61 := GenS7{F0: -9223372036854775808, F1: 1, F2: 0}
                                    var v62 int16 = -908
                                    _ = fi55
                                    _ = v56
                                    _ = v58
                                    _ = v61
                                    _ = v62
                                }
                            case 128:
                                switch -312 {
                                case -2147483648:
                                    for fi63 := 0; fi63 < 12; fi63++ {
                                        if (-2147483648 < 42) {
                                            v1 -= 1000
                                        } else {
                                            switch -129 {
                                            case 32767:
                                                v3 = (v3 - (v3 << ((1) & 63)))
                                            case 42:
                                                v4 = 140
                                            }
                                            v2 = (((float64(1e10) * 0.0) * float64((1e10 / 3.14159))) / 0.1)
                                        }
                                        var v64 uint16 = (^(v5))
                                        var v65 uint = 2
                                        _ = fi63
                                        _ = v64
                                        _ = v65
                                    }
                                case -9223372036854775808:
                                    for fi66 := 0; fi66 < 19; fi66++ {
                                        v4 -= v4
                                        v3 -= 841
                                        v67 := []int{}
                                        _ = fi66
                                        _ = v67
                                    }
                                }
                            }
                        }
                    }
                    v68 := []int{fi46, fi46, 32768, (fi46 / (((fi46 % ((9223372036854775807) | 1))) | 1))}
                    var v69 int8 = -128
                    _ = v68
                    _ = v69
                } else {
                    v70 := GenS6{F0: 18446744073709551615}
                    _ = v70
                }
            } else {
                v71 := GenS7{F0: -387, F1: 0, F2: 100}
                _ = v71
            }
            _ = fi46
        }
        v72 := GenS2{F0: 42, F1: -1, F2: (v5 % (((v5 << ((2) & 63))) | 1))}
        _ = v72
    } else {
        v1 = 1000
        switch 1 {
        case 102:
            delete(v6, genFn0((^(v4)), (v4 ^ v4), 0))
        case 1:
            for fi73 := 0; fi73 < 19; fi73++ {
                v3 += genFn0((^(v4)), v4, 42)
                if false {
                    for fi74 := 0; fi74 < 15; fi74++ {
                        v2 = v2
                        switch 127 {
                        case -2147483648:
                            v75 := GenS7{F0: 649, F1: -783, F2: -47}
                            _ = v75
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v2 = (((-2.25 + (((float64(((1e10 - 0.0) + (1e10 / 1.0))) * -2.25) + (0.0 + v2)) + float64(((float32(v2) - float32(v2)) - ((0.1 - (float32(-2.25) - -2.25)) + 0.1))))) * ((((float64((-2.25 - float32((float64(0.0) - ((v2 / 0.1) - (v2 - float64((-2.25 * (0.1 / (-2.25 * float32((float64(0.1) - v2)))))))))))) + v2) / 1.0) / (v2 - v2)) * (v2 / float64((3.14159 / 3.14159))))) * (v2 / float64((float32(float64(1.5)) + 3.14159))))
                        default:
                            v76 := GenS8{F0: -9223372036854775808, F1: 4294967295, F2: (fi73 >> ((4294967295) & 63))}
                            _ = v76
                        }
                        v77 := GenS6{F0: 2}
                        _ = fi74
                        _ = v77
                    }
                    var v78 bool = false
                    _ = v78
                } else {
                    v79 := []int{fi73, (fi73 & (fi73 >> ((151) & 63))), fi73}
                    _ = v79
                }
                for fi80 := 0; fi80 < 13; fi80++ {
                    var v81 int8 = -4
                    if (v5 > 42) {
                        if (((v2 * ((0.1 / 0.1) - v2)) == 1.5) && (true && true)) {
                            var v82 int16 = v1
                            v83 := GenS6{F0: 18446744073709551615}
                            fmt.Println(v1, v3, v81)
                            fmt.Println(v5)
                            v84 := []int{fi73, fi73, -2147483648, (fi80 & fi80)}
                            if (!false) {
                                v2 = 0.1
                                v85 := GenS3{F0: (v82 * (v82 >> ((870) & 63))), F1: 1}
                                v6[(v3 * (v3 / ((v3) | 1)))] = (fi80 >> ((255) & 63))
                                _ = v85
                            } else {
                                delete(v6, v3)
                                v86 := GenS1{F0: fi80, F1: (v5 * 0), F2: fi73}
                                var v87 int = fi80
                                v1 = v1
                                if len(v84) >= 2 {
                                    v84 = v84[1:len(v84)]
                                }
                                _ = v86
                                _ = v87
                            }
                            v5 += (v5 / ((0) | 1))
                            _ = v82
                            _ = v83
                            _ = v84
                        } else {
                            var v88 uint16 = uint16(v3)
                            var v89 uint8 = (^(v4))
                            v90 := []int{}
                            _ = v88
                            _ = v89
                            _ = v90
                        }
                        var v91 int64 = 42
                        _ = v91
                    } else {
                        switch int(v1) {
                        case -1:
                            var v92 uint32 = (v3 * v3)
                            _ = v92
                        case 601:
                            v93 := GenS9{F0: -1, F1: (false && (0.1 > 0.0))}
                            v94 := GenS4{F0: (!(1.5 <= 1e10))}
                            if false {
                                var v95 uint16 = (v5 % ((895) | 1))
                                fmt.Println(v1, v1)
                                _ = v95
                            } else {
                                fmt.Println(v2, v3)
                            }
                            _ = v93
                            _ = v94
                        default:
                            fmt.Println(v2, fi73)
                        }
                        fmt.Println(v3)
                    }
                    _ = fi80
                    _ = v81
                }
                _ = fi73
            }
        case 9223372036854775807:
            for fi96 := 0; fi96 < 10; fi96++ {
                switch "\n\t" {
                case "foo bar":
                    switch (fi96 * -762) {
                    case 42:
                        var v97 int64 = 0
                        _ = v97
                    case 127:
                        var v98 uint8 = 36
                        _ = v98
                    case 0:
                        delete(v6, 672)
                    }
                }
                for fi99 := 0; fi99 < 5; fi99++ {
                    var v100 int64 = 0
                    var v101 int64 = (v100 * v100)
                    _ = fi99
                    _ = v100
                    _ = v101
                }
                _ = fi96
            }
        }
    }
    for fi102 := 0; fi102 < 20; fi102++ {
        if (714 == 638) {
            v6[v3] = (fi102 | (fi102 % ((fi102) | 1)))
        }
        _ = fi102
    }
    v103 := []int{42, 0}
    for fi104 := 0; fi104 < 18; fi104++ {
        if (!((v1 << ((2) & 63)) != (-(v1)))) {
            var v105 uint16 = v5
            switch fi104 {
            case 2147483647:
                v3 += (v3 & v3)
            case 0:
                switch (fi104 + (fi104 - 4294967295)) {
                case 257:
                    delete(v6, (v3 * v3))
                case 195:
                    if len(v103) > 0 {
                        v103[0] = 128
                    }
                case 4294967295:
                    switch -730 {
                    case -282:
                        if true {
                            var v106 int32 = int32(v5)
                            var v107 uint32 = (v3 & (v3 / ((v3) | 1)))
                            v107 = genFn0((v4 + v4), v4, 1)
                            _ = v106
                            _ = v107
                        } else {
                            v108 := GenS10{F0: 42, F1: v4, F2: int64(v4)}
                            var v109 int32 = 0
                            _ = v108
                            _ = v109
                        }
                    }
                }
            case 42:
                fmt.Println(v4, v1)
            }
            v105 += (v105 ^ v105)
            _ = v105
        } else {
            delete(v6, v3)
        }
        if len(v103) >= 2 {
            v103 = v103[1:len(v103)]
        }
        if false {
            if len(v103) >= 2 {
                v103 = v103[1:len(v103)]
            }
        } else {
            for fi110 := 0; fi110 < 3; fi110++ {
                var v111 int8 = 0
                _ = fi110
                _ = v111
            }
            var v112 int64 = 314
            _ = v112
        }
        _ = fi104
    }
    for fi113 := 0; fi113 < 6; fi113++ {
        for fi114 := 0; fi114 < 2; fi114++ {
            fmt.Println(v5, v2, v3)
            if (false || (-111 == int8(v5))) {
                var v115 int8 = 0
                v5 += v5
                switch fi114 {
                case 4294967295:
                    fmt.Println(v3, fi114, v1)
                case -1:
                    var v116 int8 = (v115 ^ v115)
                    _ = v116
                case 1:
                    if false {
                        fmt.Println(v115, v2)
                    } else {
                        if (0 != 0) {
                            if len(v103) > 0 {
                                v103[0] = (fi113 << ((65535) & 63))
                            }
                            switch (fi113 / ((fi114) | 1)) {
                            case 32767:
                                if ((true || true) && true) {
                                    v117 := v6[uint32(v5)]
                                    switch "hello" {
                                    case "hello":
                                        var v118 bool = (false && true)
                                        _ = v118
                                    case "":
                                        var v119 uint = 216
                                        if (1e10 <= ((v2 + (v2 + v2)) * 3.14159)) {
                                            if false {
                                                fmt.Println(v1, v119, v117)
                                                if true {
                                                    var v120 int8 = (v115 % (((v115 % (((^(v115))) | 1))) | 1))
                                                    if len(v103) >= 2 {
                                                        v103 = v103[1:len(v103)]
                                                    }
                                                    _ = len(v6)
                                                    _ = v120
                                                } else {
                                                    var v121 int8 = (v115 % ((1) | 1))
                                                    if false {
                                                        fmt.Println(v3, v2, v117)
                                                        var v122 float64 = 1e10
                                                        _ = v122
                                                    } else {
                                                        v117 += (^(fi113))
                                                    }
                                                    _ = v121
                                                }
                                            } else {
                                                if (true || (false && (v1 > (^(v1))))) {
                                                    if (2147483647 == 1) {
                                                        var v123 uint64 = uint64(v5)
                                                        var v124 uint = (v119 + v119)
                                                        _ = v123
                                                        _ = v124
                                                    }
                                                } else {
                                                    var v125 float32 = 0.1
                                                    if (-128 != (-(v115))) {
                                                        v126 := GenS4{F0: true}
                                                        var v127 string = (("foo bar" + (((("" + "") + ((("foo bar" + "hello") + "\n\t") + "a")) + "foo bar") + "\n\t")) + "hello")
                                                        _ = v126
                                                        _ = v127
                                                    }
                                                    if true {
                                                        v3 = v3
                                                        var v128 float32 = 0.0
                                                        _ = v128
                                                    } else {
                                                        var v129 float32 = v125
                                                        if (1.5 == v2) {
                                                            var v130 uint64 = 2
                                                            v129 += (v125 - ((float32(0.0) + 1e10) / ((float32(0.0) - v129) + float32(v2))))
                                                            _ = v130
                                                        }
                                                        _ = v129
                                                    }
                                                    _ = v125
                                                }
                                                v5 = (^(v5))
                                            }
                                            var v131 uint16 = (v5 << (((v119 & v119)) & 63))
                                            _ = v131
                                        } else {
                                            var v132 int16 = int16(v115)
                                            switch ("foo bar" + "a") {
                                            case "a":
                                                if true {
                                                    var v133 string = "foo bar"
                                                    v119 = (^(v119))
                                                    if (!(true && false)) {
                                                        v134 := v6[v3]
                                                        v135 := GenS7{F0: -1, F1: -860, F2: (v115 ^ -128)}
                                                        if false {
                                                            v3 = (v3 % ((v3) | 1))
                                                            v136 := GenS5{F0: -1, F1: 91, F2: int16(v119)}
                                                            _ = len(v103)
                                                            _ = v136
                                                        } else {
                                                            if false {
                                                                v117 += 255
                                                                v137 := GenS4{F0: true}
                                                                var v138 uint8 = (v4 & (v4 << ((v119) & 63)))
                                                                if false {
                                                                    v139 := GenS11{F0: v5, F1: 4294967295, F2: (v115 & (v115 * (v115 ^ (v115 / ((v115) | 1)))))}
                                                                    _ = v139
                                                                }
                                                                v103 = append(v103, fi113)
                                                                v138 += v138
                                                                _ = v137
                                                                _ = v138
                                                            }
                                                            delete(v6, (^(v3)))
                                                            v3 -= 375
                                                        }
                                                        _ = v134
                                                        _ = v135
                                                    }
                                                    _ = v133
                                                } else {
                                                    v5 -= (v5 ^ (v5 * (v5 ^ (v5 ^ (^(v5))))))
                                                }
                                            }
                                            if len(v103) >= 2 {
                                                v103 = v103[1:len(v103)]
                                            }
                                            _ = v132
                                        }
                                        v115 -= v115
                                        _ = v119
                                    }
                                    if (!(true || true)) {
                                        var v140 int = fi114
                                        _ = v140
                                    } else {
                                        if true {
                                            switch "foo bar" {
                                            case "":
                                                fmt.Println(v5, v4)
                                            case "foo bar":
                                                switch "a" {
                                                case "\n\t":
                                                    if (!(!false)) {
                                                        switch "foo bar" {
                                                        case "a":
                                                            _ = len(v103)
                                                            var v141 uint8 = (v4 - (v4 - 193))
                                                            _ = v141
                                                        case "\n\t":
                                                            v142 := GenS12{F0: 4294967295}
                                                            _ = v142
                                                        case "":
                                                            var v143 uint32 = uint32(v5)
                                                            _ = v143
                                                        default:
                                                            switch ("\n\t" + "foo bar") {
                                                            case "\n\t":
                                                                var v144 int8 = v115
                                                                _ = v144
                                                            case "":
                                                                v145 := GenS9{F0: 32768, F1: false}
                                                                _ = v145
                                                            case "a":
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v146 int16 = (v1 / (((v1 ^ v1)) | 1))
                                                                    _ = v146
                                                                default:
                                                                    v147 := GenS13{F0: 42}
                                                                    _ = v147
                                                                }
                                                            }
                                                        }
                                                    }
                                                case "foo bar":
                                                    var v148 uint8 = (^(v4))
                                                    _ = v148
                                                case "hello":
                                                    v5 = (v5 - v5)
                                                default:
                                                    v149 := v6[v3]
                                                    _ = v149
                                                }
                                            case "\n\t":
                                                switch fi113 {
                                                case 0:
                                                    var v150 bool = (fi113 <= -2147483648)
                                                    _ = v150
                                                case -109:
                                                    v2 = 1.5
                                                default:
                                                    var v151 int32 = 0
                                                    _ = v151
                                                }
                                            }
                                            v1 = v1
                                        } else {
                                            if true {
                                                v115 -= v115
                                                _ = len(v6)
                                            }
                                        }
                                        v152 := GenS3{F0: (v1 % ((-32768) | 1)), F1: 473}
                                        _ = v152
                                    }
                                    _ = v117
                                }
                            case 4294967295:
                                if false {
                                    v115 += v115
                                }
                            }
                            v3 -= 0
                        }
                    }
                    switch 128 {
                    case 52:
                        var v153 uint = 1
                        _ = v153
                    case 32768:
                        v154 := GenS14{F0: (v5 / ((v5) | 1)), F1: v1}
                        _ = v154
                    case -488:
                        switch ("a" + "") {
                        case "a":
                            if len(v103) > 0 {
                                v103[0] = fi114
                            }
                        case "":
                            v3 += (^(v3))
                        default:
                            var v155 uint8 = v4
                            _ = v155
                        }
                    default:
                        var v156 int = fi114
                        switch (fi114 & (fi114 >> ((65535) & 63))) {
                        case -9223372036854775808:
                            if ((v5 == 65535) || (-2.25 != 1e10)) {
                                v157 := v6[genFn0((v4 >> ((18446744073709551615) & 63)), (^(v4)), v115)]
                                v158 := GenS11{F0: (v5 + v5), F1: 218, F2: (v115 + (v115 - (-(v115))))}
                                v1 = -32768
                                switch ("\n\t" + ("hello" + "hello")) {
                                case "foo bar":
                                    v1 = 1000
                                case "\n\t":
                                    v3 += uint32(v5)
                                }
                                v156 = fi113
                                switch -820 {
                                case 543:
                                    var v159 uint32 = (^(v3))
                                    _ = v159
                                case 2147483647:
                                    switch int(v5) {
                                    case -513:
                                        switch "a" {
                                        case "\n\t":
                                            v5 -= (v5 << ((734) & 63))
                                        case "a":
                                            v160 := GenS15{F0: -105}
                                            _ = v160
                                        case "hello":
                                            v103 = append(v103, fi113)
                                        default:
                                            if true {
                                                var v161 uint32 = genFn0(0, v4, (v115 / ((v115) | 1)))
                                                v162 := GenS11{F0: uint16(v3), F1: 4294967295, F2: v115}
                                                fmt.Println(v3)
                                                _ = v161
                                                _ = v162
                                            }
                                        }
                                    case 1:
                                        v157 = fi113
                                    case -1:
                                        v5 = 408
                                    default:
                                        v103 = append(v103, 9223372036854775807)
                                    }
                                case 9223372036854775807:
                                    if (v2 <= v2) {
                                        v2 += (-2.25 / 1.0)
                                        v3 -= genFn0(v4, (v4 << ((0) & 63)), 1)
                                    } else {
                                        v163 := GenS10{F0: 42, F1: 2, F2: 42}
                                        var v164 uint8 = (v4 & (v4 + 42))
                                        if true {
                                            var v165 uint = 65535
                                            _ = v165
                                        }
                                        v166 := GenS10{F0: 0, F1: v164, F2: 42}
                                        if (false && ((v5 == v5) && false)) {
                                            v103 = append(v103, v156)
                                        } else {
                                            switch fi114 {
                                            case 2147483647:
                                                var v167 uint16 = 1
                                                _ = v167
                                            default:
                                                var v168 uint = 4294967295
                                                _ = v168
                                            }
                                            if false {
                                                _ = len(v103)
                                            }
                                            switch ("hello" + "") {
                                            case "":
                                                if false {
                                                    v169 := GenS1{F0: (fi113 + v156), F1: 794, F2: 330}
                                                    _ = v169
                                                } else {
                                                    var v170 int8 = (v115 << ((255) & 63))
                                                    v171 := GenS14{F0: (v5 << ((2) & 63)), F1: v1}
                                                    var v172 int32 = 42
                                                    _ = v170
                                                    _ = v171
                                                    _ = v172
                                                }
                                            }
                                        }
                                        _ = v163
                                        _ = v164
                                        _ = v166
                                    }
                                default:
                                    fmt.Println(v1, v5, fi114)
                                }
                                _ = v157
                                _ = v158
                            } else {
                                var v173 int32 = -10
                                _ = v173
                            }
                        }
                        if false {
                            switch "hello" {
                            case "":
                                fmt.Println(v115)
                            case "a":
                                var v174 int8 = -1
                                _ = v174
                            case "\n\t":
                                if false {
                                    fmt.Println(v1, v2, v5)
                                    if (!(!(int(v3) != fi114))) {
                                        _ = len(v103)
                                        fmt.Println(fi114, v115)
                                        switch v156 {
                                        case 32768:
                                            v175 := GenS3{F0: (v1 ^ v1), F1: -967}
                                            _ = v175
                                        case 128:
                                            v176 := GenS16{F0: (v115 ^ (v115 >> ((200) & 63))), F1: v1}
                                            _ = v176
                                        default:
                                            var v177 string = ("a" + "")
                                            _ = v177
                                        }
                                        v5 = v5
                                        v1 = 194
                                    }
                                    fmt.Println(v156, v1)
                                } else {
                                    if true {
                                        _ = len(v103)
                                        var v178 int8 = 127
                                        _ = v178
                                    } else {
                                        switch "" {
                                        case "hello":
                                            var v179 string = ("hello" + ("\n\t" + "hello"))
                                            _ = v179
                                        case "a":
                                            var v180 float32 = 1e10
                                            _ = v180
                                        case "\n\t":
                                            v181 := GenS17{F0: (v156 >> ((361) & 63))}
                                            _ = v181
                                        }
                                        var v182 string = "hello"
                                        fmt.Println(v3)
                                        _ = v182
                                    }
                                    if (!false) {
                                        var v183 uint8 = (^(v4))
                                        _ = v183
                                    } else {
                                        if (true && true) {
                                            var v184 int32 = 768
                                            _ = v184
                                        }
                                        if false {
                                            switch -2147483648 {
                                            case 892:
                                                switch ("a" + (("a" + (((((("" + ("a" + "hello")) + "") + "a") + "a") + "a") + "hello")) + "a")) {
                                                case "\n\t":
                                                    var v185 float32 = 0.1
                                                    _ = v185
                                                case "":
                                                    var v186 float64 = v2
                                                    _ = v186
                                                default:
                                                    fmt.Println(v3)
                                                }
                                            case -9223372036854775808:
                                                var v187 uint8 = 133
                                                _ = v187
                                            default:
                                                var v188 uint16 = v5
                                                _ = v188
                                            }
                                            v2 = float64((1.5 / 0.1))
                                            switch v156 {
                                            case -81:
                                                var v189 uint = 42
                                                _ = v189
                                            }
                                        }
                                        v190 := GenS3{F0: 1000, F1: 42}
                                        fmt.Println(v115)
                                        var v191 uint16 = 1
                                        _ = v190
                                        _ = v191
                                    }
                                }
                            }
                            v192 := v6[(v3 >> ((682) & 63))]
                            _ = v192
                        } else {
                            v4 = (v4 / (((^(v4))) | 1))
                            v3 += (^(v3))
                        }
                        _ = v156
                    }
                    fmt.Println(v115, fi114)
                    fmt.Println(v1, v5, v1)
                }
                _ = v115
            }
            v6[2] = fi113
            _ = fi114
        }
        var v193 float64 = ((v2 - v2) * (((1e10 - v2) / -2.25) + -2.25))
        if false {
            v194 := GenS15{F0: int8(v1)}
            _ = v194
        } else {
            v195 := GenS3{F0: (^(v1)), F1: -1}
            _ = v195
        }
        _ = fi113
        _ = v193
    }
    if false {
        v5 -= (v5 >> ((2) & 63))
        var v196 int = 42
        v5 = v5
        _ = v196
    }
    var v197 uint32 = (v3 & 2147483648)
    v198 := GenS18{F0: 1, F1: ""}
    v3 -= (v3 / ((v197) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v103:", len(v103), v103)
    fmt.Println("v197:", v197)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v103
    _ = v197
    _ = v198
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
