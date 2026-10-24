package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int32; F2 int16 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int8; F1 int16; F2 int8 }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool; F1 int64; F2 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64; F1 int; F2 bool }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 uint32; F1 uint32; F2 bool }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int32; F1 int }
type GenS18 struct { F0 uint; F1 bool; F2 int }
func genFn0(p0 uint8, p1 uint, p2 int64) int16 {
    p1 = p1
    v7 := []int{9223372036854775807, 0, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 uint8 = 255
    var v4 uint = 4294967295
    var v5 uint64 = 2
    var v6 uint32 = 626
    var v8 uint = 0
    var v9 int64 = -1
    _ = len(v1)
    var v10 int16 = 32767
    if false {
        v9 += int64(v5)
        if true {
            var v11 int32 = 42
            fmt.Println(v6)
            if ((3.14159 + 0.0) >= float64(float32(0.0))) {
                switch "hello" {
                case "a":
                    for fi12 := 0; fi12 < 4; fi12++ {
                        var v13 int64 = (v9 - (^(v9)))
                        v9 += v13
                        v14 := GenS1{F0: (-(v10))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case "\n\t":
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v16 := GenS1{F0: (^(v10))}
                        switch 42 {
                        case -911:
                            v17 := []int{fi15, fi15}
                            _ = v17
                        case 128:
                            fmt.Println(fi15)
                        default:
                            v4 += 0
                        }
                        _ = fi15
                        _ = v16
                    }
                case "hello":
                    if ((!(!((!((^(v6)) < (v6 / ((v6) | 1)))) && true))) || true) {
                        var v18 int8 = 100
                        _ = v18
                    } else {
                        if (!false) {
                            var v19 int16 = 183
                            switch "a" {
                            case "\n\t":
                                v20 := GenS1{F0: -32768}
                                _ = v20
                            case "foo bar":
                                var v21 uint16 = 255
                                _ = v21
                            default:
                                v11 -= v11
                            }
                            _ = v19
                        } else {
                            if true {
                                v11 = v11
                                v22 := []int{32768, 305, 32768, -863}
                                _ = v22
                            } else {
                                v23 := GenS1{F0: int16(v4)}
                                _ = v23
                            }
                            v24 := v2[v6]
                            var v25 float32 = 1e10
                            _ = v24
                            _ = v25
                        }
                    }
                }
                for fi26 := 0; fi26 < 3; fi26++ {
                    if (((v9 % (((v9 & int64(v4))) | 1)) >= (v9 | v9)) && true) {
                        v27 := GenS1{F0: -260}
                        v3 += v3
                        v28 := []int{(fi26 | (fi26 << ((2) & 63))), (fi26 >> ((18446744073709551615) & 63)), (fi26 + fi26), fi26}
                        _ = v27
                        _ = v28
                    }
                    var v29 int = -9223372036854775808
                    v9 = (v9 & v9)
                    _ = fi26
                    _ = v29
                }
                if (!false) {
                    v30 := []int{0, -1, 32768, 770}
                    v31 := GenS2{F0: (v10 + int16(v3)), F1: (v5 / ((v5) | 1))}
                    if false {
                        if true {
                            var v32 int = -2147483648
                            v33 := GenS3{F0: (v3 & uint8(v6)), F1: v11, F2: (v10 | (v10 % ((0) | 1)))}
                            var v34 string = (("\n\t" + ("" + (("\n\t" + (("a" + ((("hello" + "") + "\n\t") + "a")) + "foo bar")) + ("a" + "foo bar")))) + "\n\t")
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS1{F0: v10}
                        fmt.Println(v10, v5)
                        _ = v35
                    }
                    _ = v30
                    _ = v31
                }
            } else {
                var v36 uint8 = 2
                var v37 uint32 = uint32(v8)
                _ = v36
                _ = v37
            }
            _ = v11
        } else {
            v4 += (v8 * v4)
            var v38 uint16 = 2
            _ = v38
        }
    } else {
        v1[42] = (v6 - v6)
    }
    fmt.Println(v10, v4, v6)
    switch "" {
    case "foo bar":
        switch -2147483648 {
        case 42:
            var v39 int32 = 42
            _ = v39
        case -976:
            v40 := GenS4{F0: v4, F1: 145}
            _ = v40
        case 238:
            if ((1 != -1) && true) {
                v4 += v4
                var v41 int8 = -49
                _ = v41
            } else {
                v5 = v5
            }
        default:
            if (v4 == v8) {
                for fi42 := 0; fi42 < 12; fi42++ {
                    v43 := GenS4{F0: 2, F1: (v5 ^ v5)}
                    _ = fi42
                    _ = v43
                }
                var v44 uint32 = 2147483647
                _ = v44
            }
        }
    case "":
        for fi45 := 0; fi45 < 3; fi45++ {
            fmt.Println(v5, v9, v5)
            _ = fi45
        }
    }
    v8 = v8
    for fi46 := 0; fi46 < 15; fi46++ {
        v9 -= (v9 % ((v9) | 1))
        var v47 uint16 = 42
        v6 = v6
        _ = fi46
        _ = v47
    }
    delete(v1, (v6 + 2))
    if ((v10 | 1000) != 0) {
        v10 = (v10 - 32767)
        var v48 int = -2147483648
        var v49 int16 = v10
        _ = v48
        _ = v49
    } else {
        if (0 <= -20) {
            v8 = (v8 >> ((v8) & 63))
            for fi50 := 0; fi50 < 2; fi50++ {
                var v51 float64 = 1e10
                v3 -= v3
                fmt.Println(v8)
                _ = fi50
                _ = v51
            }
        } else {
            for fi52 := 0; fi52 < 11; fi52++ {
                switch fi52 {
                case 2147483647:
                    if (-9223372036854775808 < v9) {
                        for fi53 := 0; fi53 < 9; fi53++ {
                            v54 := GenS2{F0: genFn0(v3, (v4 % (((v4 / ((4294967295) | 1))) | 1)), (v9 / ((v9) | 1))), F1: (v5 - 42)}
                            v4 = v4
                            _ = fi53
                            _ = v54
                        }
                    }
                case -2147483648:
                    if false {
                        if (false && true) {
                            v55 := GenS1{F0: v10}
                            _ = v55
                        }
                    } else {
                        v56 := GenS3{F0: v3, F1: -2147483648, F2: (v10 ^ 1000)}
                        if false {
                            for fi57 := 0; fi57 < 20; fi57++ {
                                var v58 float32 = 3.14159
                                _ = fi57
                                _ = v58
                            }
                        }
                        fmt.Println(v9, v10)
                        _ = v56
                    }
                case 4294967295:
                    v10 = 0
                }
                for fi59 := 0; fi59 < 19; fi59++ {
                    v3 += v3
                    v60 := GenS5{F0: 4294967295, F1: (v10 >> ((v8) & 63))}
                    v9 -= (v9 + (-(v9)))
                    _ = fi59
                    _ = v60
                }
                v61 := GenS1{F0: (v10 / ((v10) | 1))}
                _ = fi52
                _ = v61
            }
            v8 += 368
        }
        if (!true) {
            if false {
                v4 = uint(v3)
                switch "" {
                case "":
                    if (!true) {
                        for fi62 := 0; fi62 < 4; fi62++ {
                            switch "foo bar" {
                            case "hello":
                                if (v5 > 0) {
                                    fmt.Println(fi62, v4)
                                    var v63 int32 = 42
                                    if true {
                                        switch "" {
                                        case "":
                                            v4 += 18446744073709551615
                                        default:
                                            v9 = (v9 ^ (v9 % (((v9 % ((v9) | 1))) | 1)))
                                        }
                                        v9 -= 0
                                    } else {
                                        v64 := []int{0, fi62, int(v5)}
                                        v65 := GenS4{F0: (v4 * v8), F1: v5}
                                        v66 := GenS6{F0: 127, F1: (v10 + (v10 << (((^(v4))) & 63))), F2: 127}
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    _ = v63
                                } else {
                                    var v67 uint64 = (^(v5))
                                    v2[(v6 + v6)] = int(v67)
                                    v9 = (v9 + v9)
                                    _ = v67
                                }
                            default:
                                switch int(v9) {
                                case 784:
                                    if (-2.25 >= float32((-2.25 * 3.14159))) {
                                        if (!true) {
                                            v5 -= (v5 * v5)
                                        }
                                    }
                                case 4294967295:
                                    for fi68 := 0; fi68 < 17; fi68++ {
                                        if ((v5 < 59) && (-2147483648 < -2147483648)) {
                                            if true {
                                                delete(v2, v6)
                                            } else {
                                                v69 := GenS5{F0: v8, F1: genFn0(v3, (v8 - (v4 - (v8 | 18446744073709551615))), (v9 ^ 0))}
                                                v70 := []int{-2147483648}
                                                v71 := GenS7{F0: 0, F1: 715}
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            }
                                            var v72 int16 = v10
                                            _ = v72
                                        } else {
                                            if false {
                                                v73 := []int{fi62, fi62, 32767, fi68}
                                                v74 := GenS8{F0: v9}
                                                _ = v73
                                                _ = v74
                                            }
                                            if true {
                                                if (v6 <= v6) {
                                                    var v75 uint8 = 42
                                                    v76 := GenS3{F0: 128, F1: -1, F2: (v10 % ((genFn0(uint8(v5), v8, v9)) | 1))}
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    v77 := GenS7{F0: 42, F1: 841}
                                                    var v78 bool = (false && (-1 <= -1))
                                                    v79 := GenS9{F0: ((-(v10)) == -32768), F1: -716, F2: false}
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                }
                                                var v80 int32 = -990
                                                v9 = (v9 >> ((v4) & 63))
                                                _ = v80
                                            } else {
                                                v8 += 42
                                                if true {
                                                    if (v6 <= v6) {
                                                        v81 := []int{4294967295, 9223372036854775807, fi68, int(v9)}
                                                        fmt.Println(v9, v5)
                                                        switch (-(fi68)) {
                                                        case 2147483647:
                                                            v8 += (v4 ^ 42)
                                                        default:
                                                            v82 := GenS10{F0: "a"}
                                                            _ = v82
                                                        }
                                                        _ = v81
                                                    } else {
                                                        v83 := GenS4{F0: (v8 + v8), F1: v5}
                                                        var v84 float64 = (1.5 - float64(((0.0 * (1.5 * 1.5)) + 1.5)))
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    v5 += v5
                                                } else {
                                                    var v85 bool = ((4294967295 <= 2147483647) || false)
                                                    _ = v85
                                                }
                                                var v86 uint8 = (v3 ^ v3)
                                                v87 := v1[(v6 % (((^(v6))) | 1))]
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                        switch fi62 {
                                        case 0:
                                            v88 := []int{(fi62 >> ((2) & 63)), fi68, (fi62 / ((fi62) | 1))}
                                            _ = v88
                                        case -2147483648:
                                            var v89 float32 = (float32(3.14159) * ((float32((-2.25 - ((-2.25 - (float64((float32((float64(3.14159) * (((1.5 - (1.5 - (0.1 + (-2.25 - 1e10)))) - 0.1) / (1.5 - 0.0)))) * float32(1e10))) * ((1e10 + 3.14159) * 0.1))) + ((((0.0 * 1.5) * -2.25) / float64((((0.0 / float32(3.14159)) / 3.14159) + 1.5))) / ((3.14159 / (0.1 * 1.5)) + 3.14159))))) + ((3.14159 - 1e10) / ((-2.25 - -2.25) + (((-2.25 - float32((3.14159 * 3.14159))) / 0.1) + -2.25)))) * ((0.1 * (((3.14159 * 1e10) * 0.0) / (1.5 - 3.14159))) * 0.1)))
                                            _ = v89
                                        case 32767:
                                            var v90 uint8 = (v3 - 127)
                                            var v91 int8 = -110
                                            if (!(!(0 > 42))) {
                                                var v92 uint8 = (v90 * v3)
                                                _ = v92
                                            } else {
                                                v93 := GenS10{F0: "foo bar"}
                                                if (!true) {
                                                    _ = len(v1)
                                                    v91 = v91
                                                } else {
                                                    switch (("\n\t" + "") + ("hello" + ("hello" + ("a" + (("hello" + "foo bar") + ("" + ("hello" + ""))))))) {
                                                    case "foo bar":
                                                        if true {
                                                            v94 := GenS11{F0: (v9 + v9), F1: -1, F2: true}
                                                            _ = v94
                                                        } else {
                                                            if false {
                                                                switch fi62 {
                                                                case -1:
                                                                    switch ("\n\t" + (("" + "foo bar") + "foo bar")) {
                                                                    case "":
                                                                        fmt.Println(v91, fi68, v3)
                                                                    case "foo bar":
                                                                        v95 := GenS8{F0: (^(v9))}
                                                                        _ = v95
                                                                    }
                                                                }
                                                                var v96 int = fi62
                                                                var v97 uint16 = 32767
                                                                _ = v96
                                                                _ = v97
                                                            } else {
                                                                v10 += (v10 ^ int16(v3))
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v4 = 718
                                                                case "foo bar":
                                                                    v98 := GenS12{F0: v5, F1: uint16(v5)}
                                                                    _ = v98
                                                                case "a":
                                                                    fmt.Println(v90, v8, v91)
                                                                }
                                                                var v99 int64 = (^(v9))
                                                                v90 -= v90
                                                                _ = v99
                                                            }
                                                        }
                                                    case "hello":
                                                        delete(v1, (v6 >> ((v4) & 63)))
                                                    case "a":
                                                        var v100 int32 = 556
                                                        _ = v100
                                                    }
                                                }
                                                if true {
                                                    v101 := []int{}
                                                    _ = v101
                                                }
                                                v8 = (v8 & (v4 - v4))
                                                var v102 int64 = 9223372036854775807
                                                var v103 uint = (v4 * 0)
                                                if ((v6 >= 2) || (false && true)) {
                                                    var v104 uint32 = (^(v6))
                                                    fmt.Println(v8, v102)
                                                    if (true && ((^(v4)) >= v8)) {
                                                        switch (^(fi68)) {
                                                        case 9223372036854775807:
                                                            v105 := v1[430]
                                                            _ = v105
                                                        case 32768:
                                                            var v106 int32 = 1
                                                            _ = v106
                                                        default:
                                                            v3 = (^(v3))
                                                        }
                                                        v10 += (v10 * (v10 * (v10 % ((v10) | 1))))
                                                    } else {
                                                        v91 = (v91 << (((v103 >> ((v103) & 63))) & 63))
                                                        var v107 bool = false
                                                        fmt.Println(v4)
                                                        var v108 float32 = 1e10
                                                        _ = v107
                                                        _ = v108
                                                    }
                                                    _ = v104
                                                } else {
                                                    var v109 int8 = 1
                                                    v3 += 127
                                                    _ = v109
                                                }
                                                _ = v102
                                                _ = v103
                                                _ = v93
                                            }
                                            var v110 int32 = -1
                                            _ = v110
                                            _ = v90
                                            _ = v91
                                        default:
                                            var v111 int = (fi68 * fi68)
                                            v9 = v9
                                            if true {
                                                if false {
                                                    delete(v1, (v6 % ((uint32(v5)) | 1)))
                                                    if ((fi62 >> (((v8 | (v8 << (((v4 - v4)) & 63)))) & 63)) >= (-(fi68))) {
                                                        v9 = (v9 / ((int64(v6)) | 1))
                                                    } else {
                                                        v112 := v1[(v6 % ((v6) | 1))]
                                                        v10 = (v10 % (((^(v10))) | 1))
                                                        switch v111 {
                                                        case 9223372036854775807:
                                                            var v113 bool = ("foo bar" >= ("hello" + (("" + "a") + "\n\t")))
                                                            v2[uint32(v4)] = (^(v111))
                                                            var v114 uint = (v4 >> ((699) & 63))
                                                            v115 := GenS10{F0: ""}
                                                            v1[42] = v6
                                                            _ = v113
                                                            _ = v114
                                                            _ = v115
                                                        case 127:
                                                            fmt.Println(v5, v5, v10)
                                                        default:
                                                            v6 = 2147483647
                                                        }
                                                        if true {
                                                            v6 += 980
                                                            fmt.Println(v9, v4)
                                                            switch ("" + "hello") {
                                                            case "foo bar":
                                                                var v116 bool = ("\n\t" < ((("hello" + "foo bar") + "a") + (("foo bar" + "") + ("foo bar" + "hello"))))
                                                                _ = v116
                                                            case "\n\t":
                                                                v117 := GenS6{F0: 42, F1: v10, F2: -48}
                                                                _ = v117
                                                            case "a":
                                                                switch (fi62 >> ((820) & 63)) {
                                                                case 1:
                                                                    var v118 string = "a"
                                                                    _ = v118
                                                                default:
                                                                    var v119 int = fi68
                                                                    _ = v119
                                                                }
                                                            default:
                                                                var v120 int8 = -128
                                                                _ = v120
                                                            }
                                                        }
                                                        _ = v112
                                                    }
                                                    var v121 string = "\n\t"
                                                    _ = v121
                                                } else {
                                                    fmt.Println(v8)
                                                    if false {
                                                        switch (fi62 ^ (fi68 ^ (-(fi68)))) {
                                                        case -124:
                                                            v111 += fi62
                                                        case 128:
                                                            if false {
                                                                v122 := []int{(-(fi62)), fi68}
                                                                _ = v122
                                                            }
                                                        }
                                                    } else {
                                                        switch "" {
                                                        case "a":
                                                            switch "hello" {
                                                            case "a":
                                                                v123 := GenS7{F0: 127, F1: 574}
                                                                _ = v123
                                                            case "":
                                                                switch v111 {
                                                                case 0:
                                                                    fmt.Println(v4)
                                                                case 42:
                                                                    v124 := GenS1{F0: (v10 ^ (v10 & v10))}
                                                                    _ = v124
                                                                default:
                                                                    v125 := v2[(v6 ^ (^(v6)))]
                                                                    _ = v125
                                                                }
                                                            default:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v126 := v2[v6]
                                                                    _ = v126
                                                                case "\n\t":
                                                                    var v127 string = ("a" + ("a" + "foo bar"))
                                                                    _ = v127
                                                                default:
                                                                    switch v111 {
                                                                    case 1:
                                                                        if true {
                                                                            if false {
                                                                                var v128 int64 = (-(v9))
                                                                                switch -433 {
                                                                                case 0:
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        var v129 uint16 = 65535
                                                                                        _ = v129
                                                                                    case "":
                                                                                        fmt.Println(v4)
                                                                                    case "hello":
                                                                                        v128 = v9
                                                                                    default:
                                                                                        v3 -= (^(v3))
                                                                                    }
                                                                                }
                                                                                var v130 int16 = 32767
                                                                                _ = v128
                                                                                _ = v130
                                                                            } else {
                                                                                v5 = 4294967295
                                                                                fmt.Println(v111)
                                                                                fmt.Println(v8, v10, v111)
                                                                                var v131 int = fi62
                                                                                v5 += (v5 ^ 220)
                                                                                if (0 <= 2) {
                                                                                    var v132 float32 = (0.1 + 1e10)
                                                                                    v133 := []int{}
                                                                                    v133 = append(v133, (fi62 << (((v4 | (v4 % ((v8) | 1)))) & 63)))
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                }
                                                                                _ = v131
                                                                            }
                                                                            if true {
                                                                                var v134 uint64 = v5
                                                                                var v135 uint = (v4 << ((v8) & 63))
                                                                                v8 += v8
                                                                                _ = v134
                                                                                _ = v135
                                                                            } else {
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v136 uint8 = (^(v3))
                                                                                    _ = v136
                                                                                default:
                                                                                    v4 = (v4 & (v4 >> (((v8 & v4)) & 63)))
                                                                                }
                                                                                v5 += uint64(v6)
                                                                                var v137 float64 = (0.0 * ((float64(-2.25) + 0.1) / 0.1))
                                                                                _ = v137
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (!false) {
                                                                            if true {
                                                                                if true {
                                                                                    var v138 uint64 = (v5 & v5)
                                                                                    switch -1 {
                                                                                    case 127:
                                                                                        var v139 float32 = 1.5
                                                                                        _ = v139
                                                                                    }
                                                                                    var v140 int8 = 127
                                                                                    _ = v138
                                                                                    _ = v140
                                                                                } else {
                                                                                    fmt.Println(v4)
                                                                                }
                                                                                v3 = (v3 / ((v3) | 1))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case "hello":
                                                            v111 = -379
                                                        }
                                                        v8 += 4294967295
                                                    }
                                                    switch "foo bar" {
                                                    case "":
                                                        var v141 int32 = 0
                                                        _ = v141
                                                    }
                                                    fmt.Println(fi62, v8, v8)
                                                }
                                                if false {
                                                    switch (v111 >> ((v8) & 63)) {
                                                    case 128:
                                                        var v142 int = 32768
                                                        _ = v142
                                                    case 1:
                                                        var v143 int = fi62
                                                        fmt.Println(v6, fi62)
                                                        v144 := []int{}
                                                        _ = v143
                                                        _ = v144
                                                    case 32767:
                                                        var v145 int = 42
                                                        _ = v145
                                                    default:
                                                        var v146 float64 = float64((1e10 * (3.14159 - 1e10)))
                                                        _ = v146
                                                    }
                                                } else {
                                                    var v147 uint32 = 2
                                                    v148 := GenS13{F0: 264, F1: (^(v6)), F2: false}
                                                    delete(v1, (v147 - v6))
                                                    v10 = genFn0((v3 - (v3 - v3)), v8, (v9 | (v9 - v9)))
                                                    _ = v147
                                                    _ = v148
                                                }
                                            }
                                            _ = v111
                                        }
                                        var v149 int = (fi62 << ((v8) & 63))
                                        _ = fi68
                                        _ = v149
                                    }
                                case 347:
                                    var v150 uint8 = (v3 | v3)
                                    _ = v150
                                }
                            }
                            _ = fi62
                        }
                    }
                case "a":
                    for fi151 := 0; fi151 < 13; fi151++ {
                        v3 += v3
                        _ = fi151
                    }
                case "\n\t":
                    if ((false || true) || false) {
                        v5 = v5
                        if true {
                            var v152 uint32 = v6
                            if true {
                                fmt.Println(v10, v152, v5)
                            }
                            delete(v1, v152)
                            _ = v152
                        }
                    }
                default:
                    v153 := []int{32768, 2147483647}
                    _ = v153
                }
            } else {
                v154 := GenS14{F0: 42}
                for fi155 := 0; fi155 < 10; fi155++ {
                    v156 := []int{(fi155 + fi155)}
                    _ = fi155
                    _ = v156
                }
                if (!(v4 >= 65535)) {
                    for fi157 := 0; fi157 < 18; fi157++ {
                        if false {
                            switch "" {
                            case "hello":
                                if true {
                                    var v158 int64 = v9
                                    v6 += (v6 / ((v6) | 1))
                                    if false {
                                        var v159 uint = (v8 * 836)
                                        for fi160 := 0; fi160 < 20; fi160++ {
                                            v158 += v9
                                            switch ("a" + ("hello" + ("\n\t" + "foo bar"))) {
                                            case "hello":
                                                var v161 uint32 = (v6 & (v6 % (((v6 * (v6 + (v6 % ((uint32(v5)) | 1))))) | 1)))
                                                _ = v161
                                            case "\n\t":
                                                var v162 uint64 = (^(v5))
                                                _ = v162
                                            case "foo bar":
                                                var v163 bool = true
                                                _ = v163
                                            }
                                            v164 := GenS15{F0: (v5 >> ((1) & 63))}
                                            _ = fi160
                                            _ = v164
                                        }
                                        _ = v159
                                    }
                                    _ = v158
                                }
                            case "foo bar":
                                for fi165 := 0; fi165 < 14; fi165++ {
                                    if false {
                                        v8 = (v8 + (v4 ^ (v4 % ((18446744073709551615) | 1))))
                                        var v166 uint32 = 4294967295
                                        v167 := []int{(fi157 + 127)}
                                        _ = v166
                                        _ = v167
                                    }
                                    v168 := GenS16{F0: v6}
                                    v169 := GenS10{F0: "\n\t"}
                                    var v170 uint64 = 844
                                    if false {
                                        switch "a" {
                                        case "a":
                                            fmt.Println(v8, fi165, v4)
                                        }
                                        switch ("" + "a") {
                                        case "a":
                                            v3 += 127
                                        case "foo bar":
                                            if false {
                                                var v171 int64 = v9
                                                if false {
                                                    v8 += 255
                                                    v6 += 42
                                                } else {
                                                    v172 := v1[v6]
                                                    _ = v172
                                                }
                                                fmt.Println(v4)
                                                _ = v171
                                            } else {
                                                fmt.Println(v4, v4)
                                                fmt.Println(v10, v3, v170)
                                                fmt.Println(fi157)
                                            }
                                        case "\n\t":
                                            v173 := []int{fi157, (fi157 * fi165)}
                                            _ = v173
                                        default:
                                            var v174 int8 = 127
                                            _ = v174
                                        }
                                    }
                                    v175 := GenS4{F0: (v8 % ((4294967295) | 1)), F1: 530}
                                    _ = fi165
                                    _ = v168
                                    _ = v169
                                    _ = v170
                                    _ = v175
                                }
                            case "a":
                                switch (fi157 / ((42) | 1)) {
                                case 835:
                                    v1[(^(v6))] = v6
                                }
                            }
                            fmt.Println(v3, v9, v6)
                        }
                        var v176 bool = true
                        var v177 string = "a"
                        _ = fi157
                        _ = v176
                        _ = v177
                    }
                    if false {
                        if (((-2.25 * ((0.1 * (3.14159 / 3.14159)) - 0.1)) * 0.1) != (3.14159 + 1e10)) {
                            if (1 == 1) {
                                v178 := []int{177, -198}
                                _ = v178
                            } else {
                                if ((v5 > v5) && (!(((float64(-2.25) * ((0.1 / -2.25) - 1.5)) + -2.25) == 1.5))) {
                                    if (v6 <= (v6 ^ (v6 | (v6 & 589)))) {
                                        v179 := GenS17{F0: 1, F1: 0}
                                        _ = v179
                                    }
                                    v3 = v3
                                } else {
                                    delete(v2, 42)
                                    v180 := []int{}
                                    _ = v180
                                }
                                v3 -= (v3 * v3)
                                var v181 uint64 = (v5 / (((v5 * 38)) | 1))
                                _ = v181
                            }
                            var v182 float32 = (3.14159 * 0.0)
                            _ = v182
                        }
                        v183 := GenS18{F0: 65535, F1: false, F2: -9223372036854775808}
                        v184 := []int{127, -1, -60}
                        _ = v183
                        _ = v184
                    } else {
                        var v185 int = 9223372036854775807
                        v9 = 9223372036854775807
                        var v186 int8 = 86
                        _ = v185
                        _ = v186
                    }
                    v187 := GenS3{F0: v3, F1: -2147483648, F2: 0}
                    _ = v187
                }
                _ = v154
            }
            var v188 uint8 = v3
            if true {
                var v189 float64 = 3.14159
                switch 623 {
                case 2147483647:
                    v5 -= (v5 >> ((v8) & 63))
                case 321:
                    v6 -= v6
                case 1:
                    switch "foo bar" {
                    case "":
                        var v190 bool = (1 < -367)
                        _ = v190
                    default:
                        fmt.Println(v6)
                    }
                default:
                    var v191 int32 = -1
                    _ = v191
                }
                _ = v189
            }
            _ = v188
        } else {
            for fi192 := 0; fi192 < 15; fi192++ {
                v8 += (v8 * 972)
                _ = fi192
            }
            switch (("hello" + "hello") + ((("hello" + "") + "hello") + "")) {
            case "a":
                fmt.Println(v5, v6)
            case "hello":
                fmt.Println(v8, v6)
            }
            var v193 uint64 = v5
            _ = v193
        }
        var v194 bool = (true && (1 != (^(v5))))
        _ = v194
    }
    var v195 int64 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v195:", v195)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v195
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
