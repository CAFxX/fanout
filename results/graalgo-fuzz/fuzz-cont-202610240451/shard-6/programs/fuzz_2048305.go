package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32; F1 int64; F2 int32 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint8; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 bool; F1 int16; F2 int }
type GenS10 struct { F0 int64; F1 uint16; F2 int32 }
type GenS11 struct { F0 int; F1 int16; F2 uint64 }

func main() {
    var v1 uint32 = 4294967295
    var v2 string = ((("" + ("\n\t" + "a")) + "foo bar") + ("hello" + "\n\t"))
    var v3 int16 = 1
    var v4 int32 = -553
    v5 := make(map[int]bool)
    switch "foo bar" {
    case "foo bar":
        v6 := []int{-1, 2147483647, -512}
        _ = v6
    }
    if true {
        v7 := []int{32768, 9223372036854775807}
        _ = v7
    } else {
        v8 := GenS1{F0: v2}
        if false {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := v5[fi9]
                v3 = 42
                v10 = (!(!v10))
                _ = fi9
                _ = v10
            }
            v11 := []int{int(v4), 32768}
            _ = v11
        }
        _ = v8
    }
    switch 365 {
    case -979:
        for fi12 := 0; fi12 < 4; fi12++ {
            if false {
                switch fi12 {
                case 0:
                    if false {
                        var v13 uint16 = 65535
                        for fi14 := 0; fi14 < 13; fi14++ {
                            fmt.Println(v1, fi14)
                            v1 = v1
                            var v15 int32 = (v4 << ((0) & 63))
                            fmt.Println(v3)
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    } else {
                        if ((v4 >> ((0) & 63)) > int32(v3)) {
                            v1 -= v1
                        }
                        if (!(43 != 117)) {
                            if true {
                                if (!(!false)) {
                                    v16 := GenS2{F0: fi12, F1: 0, F2: 9223372036854775807}
                                    v17 := GenS2{F0: (fi12 / (((fi12 * fi12)) | 1)), F1: 511, F2: 309}
                                    fmt.Println(fi12)
                                    _ = v16
                                    _ = v17
                                }
                                var v18 int32 = 577
                                _ = v18
                            }
                            v1 = v1
                            _ = len(v5)
                        } else {
                            v19 := GenS2{F0: (fi12 | 0), F1: 255, F2: 42}
                            v4 = -1
                            _ = v19
                        }
                    }
                }
                switch 2147483647 {
                case 243:
                    fmt.Println(v1, v3, v3)
                case 986:
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 int64 = 1
                        v22 := []int{}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
                var v23 float64 = -2.25
                _ = v23
            } else {
                if true {
                    v2 = (("" + v2) + (("hello" + ((v2 + "") + v2)) + ((v2 + v2) + v2)))
                    v3 -= 32767
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v3 = (v3 << ((4294967295) & 63))
                        v2 = ("hello" + (((v2 + ((((v2 + v2) + (v2 + ((("\n\t" + (v2 + v2)) + "a") + v2))) + "a") + (v2 + ("hello" + v2)))) + ("foo bar" + v2)) + v2))
                        _ = fi24
                    }
                } else {
                    v25 := v5[(fi12 >> ((683) & 63))]
                    var v26 string = (v2 + v2)
                    _ = v25
                    _ = v26
                }
                var v27 uint8 = 109
                _ = v27
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                v29 := []int{fi12, 1}
                switch (fi28 + 718) {
                case 0:
                    var v30 uint8 = 2
                    delete(v5, (fi28 + (^(fi12))))
                    v1 = v1
                    if len(v29) > 0 {
                        v29[0] = (fi28 << ((924) & 63))
                    }
                    _ = v30
                case 4294967295:
                    var v31 int8 = -104
                    _ = v31
                default:
                    var v32 uint16 = 42
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
            fmt.Println(v4, v1, v4)
            _ = fi12
        }
    default:
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 float64 = -2.25
            _ = fi33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 uint64 = 18446744073709551615
            v5[fi35] = (((fi35 << ((65535) & 63)) > fi36) && (1 == 255))
            v3 = v3
            v38 := GenS3{F0: 493, F1: 1, F2: 457}
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    var v39 float32 = -2.25
    if true {
        var v40 int8 = -128
        _ = v40
    }
    v41 := GenS1{F0: (v2 + "hello")}
    v42 := GenS4{F0: 65535, F1: -9223372036854775808, F2: (v4 ^ (v4 | 577))}
    v43 := GenS5{F0: 42, F1: (v3 ^ (-(v3))), F2: 1}
    var v44 uint8 = 255
    if true {
        v45 := GenS5{F0: 1, F1: (^(v3)), F2: 18446744073709551615}
        if (!true) {
            for fi46 := 0; fi46 < 16; fi46++ {
                if (true || false) {
                    v47 := []int{617}
                    v2 = ("a" + v2)
                    _ = v47
                }
                if true {
                    if (!(v39 > (v39 * float32(0.0)))) {
                        fmt.Println(fi46)
                        v39 += 3.14159
                        for fi48 := 0; fi48 < 3; fi48++ {
                            v1 -= 311
                            var v49 uint16 = 42
                            _ = fi48
                            _ = v49
                        }
                    } else {
                        v2 = v2
                        v5[fi46] = true
                    }
                    fmt.Println(v39)
                } else {
                    if (!true) {
                        switch (-(fi46)) {
                        case -162:
                            switch v2 {
                            case "a":
                                v50 := []int{fi46, fi46}
                                _ = v50
                            case "":
                                fmt.Println(fi46, v1)
                            default:
                                for fi51 := 0; fi51 < 1; fi51++ {
                                    v52 := GenS2{F0: (fi51 / ((-2147483648) | 1)), F1: 65535, F2: -9223372036854775808}
                                    fmt.Println(fi46)
                                    _ = fi51
                                    _ = v52
                                }
                            }
                        case 479:
                            for fi53 := 0; fi53 < 4; fi53++ {
                                var v54 uint16 = 65535
                                _ = fi53
                                _ = v54
                            }
                        }
                        for fi55 := 0; fi55 < 18; fi55++ {
                            switch fi55 {
                            case 0:
                                var v56 int = -396
                                _ = v56
                            case 42:
                                v57 := GenS6{F0: (947 >= 18446744073709551615), F1: 2, F2: v2}
                                _ = v57
                            }
                            if true {
                                var v58 uint16 = 42
                                v3 = int16(v58)
                                v1 -= 1
                                _ = v58
                            } else {
                                v59 := GenS4{F0: (v1 % ((v1) | 1)), F1: 189, F2: (v4 | int32(v3))}
                                if false {
                                    if (false && (((v44 / ((42) | 1)) >= (v44 / ((2) | 1))) || (v44 < 2))) {
                                        switch v2 {
                                        case "":
                                            v60 := GenS1{F0: ("\n\t" + "foo bar")}
                                            _ = v60
                                        case "a":
                                            if ((fi46 <= (fi46 * fi55)) || (3.14159 <= 0.0)) {
                                                var v61 bool = false
                                                fmt.Println(fi55, v1, v61)
                                                _ = v61
                                            }
                                        case "foo bar":
                                            v62 := GenS7{F0: -1}
                                            _ = v62
                                        default:
                                            _ = len(v5)
                                        }
                                        if true {
                                            v63 := []int{0, (fi46 >> ((105) & 63)), (fi55 - fi55), (fi55 - (fi55 & (fi55 - fi46)))}
                                            v44 -= v44
                                            fmt.Println(fi46)
                                            _ = v63
                                        } else {
                                            var v64 string = ((("" + v2) + v2) + ((v2 + v2) + ((v2 + (((v2 + "") + "foo bar") + ((("\n\t" + "foo bar") + v2) + (v2 + ((v2 + v2) + v2))))) + v2)))
                                            _ = v64
                                        }
                                        delete(v5, fi55)
                                        if (!((^(v1)) < v1)) {
                                            switch fi55 {
                                            case -471:
                                                v39 = (1e10 - float32(-2.25))
                                            case 1:
                                                if ((v1 ^ 2147483647) > (v1 & 42)) {
                                                    if (v1 > uint32(v44)) {
                                                        var v65 uint16 = 970
                                                        switch fi55 {
                                                        case 127:
                                                            switch v2 {
                                                            case "foo bar":
                                                                fmt.Println(fi55, fi46, v3)
                                                            default:
                                                                v66 := []int{fi46, 128, (fi46 & (fi46 << ((65535) & 63)))}
                                                                _ = v66
                                                            }
                                                        default:
                                                            var v67 uint8 = (v44 >> ((1) & 63))
                                                            _ = v67
                                                        }
                                                        var v68 bool = false
                                                        _ = v65
                                                        _ = v68
                                                    } else {
                                                        fmt.Println(v39, v39)
                                                        if (v1 < 65535) {
                                                            var v69 bool = true
                                                            var v70 int64 = -1
                                                            var v71 int8 = int8(v1)
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        switch fi55 {
                                                        case -738:
                                                            var v72 uint16 = 1
                                                            _ = v72
                                                        case 128:
                                                            var v73 uint32 = (^(v1))
                                                            _ = v73
                                                        }
                                                    }
                                                    var v74 uint8 = v44
                                                    var v75 int32 = (^(v4))
                                                    _ = v74
                                                    _ = v75
                                                } else {
                                                    fmt.Println(v39, v44, v3)
                                                    if false {
                                                        if true {
                                                            switch "hello" {
                                                            case "a":
                                                                var v76 int16 = 1000
                                                                _ = v76
                                                            case "foo bar":
                                                                v39 -= 0.1
                                                            }
                                                            v77 := GenS2{F0: fi46, F1: 337, F2: -9223372036854775808}
                                                            _ = v77
                                                        }
                                                        v3 = int16(fi46)
                                                    } else {
                                                        v1 = 2147483647
                                                        v44 = (v44 << ((840) & 63))
                                                    }
                                                }
                                            case 127:
                                                v1 -= v1
                                            default:
                                                v44 = 180
                                            }
                                            v4 += v4
                                        } else {
                                            var v78 string = v2
                                            v79 := []int{}
                                            _ = v78
                                            _ = v79
                                        }
                                        v44 = (v44 % (((^(v44))) | 1))
                                    }
                                    var v80 uint64 = 339
                                    _ = v80
                                } else {
                                    switch (("a" + (("" + "hello") + ((("foo bar" + ("\n\t" + v2)) + v2) + "\n\t"))) + v2) {
                                    case "foo bar":
                                        var v81 bool = true
                                        if v81 {
                                            var v82 int = 32768
                                            if v81 {
                                                fmt.Println(v1, v3, v82)
                                            } else {
                                                switch v82 {
                                                case 9223372036854775807:
                                                    v4 = v4
                                                case -1:
                                                    if v81 {
                                                        var v83 uint32 = 11
                                                        _ = v83
                                                    } else {
                                                        v5[fi46] = v81
                                                    }
                                                }
                                                fmt.Println(v1, v44)
                                            }
                                            v84 := []int{fi46}
                                            _ = v82
                                            _ = v84
                                        }
                                        v5[2147483647] = (false && ((v39 < ((-2.25 / (-2.25 - 1.5)) + 1.5)) && (42 == v44)))
                                        _ = v81
                                    }
                                }
                                _ = v59
                            }
                            if true {
                                if false {
                                    v85 := v5[(fi55 ^ (fi46 & 1))]
                                    v85 = v85
                                    _ = v85
                                } else {
                                    delete(v5, fi55)
                                    v39 += 3.14159
                                }
                                if true {
                                    if false {
                                        v86 := []int{(fi46 * (^(fi55)))}
                                        v86 = append(v86, fi46)
                                        _ = v86
                                    } else {
                                        var v87 string = ""
                                        switch (v87 + "\n\t") {
                                        case "\n\t":
                                            var v88 int64 = 9223372036854775807
                                            _ = v88
                                        case "":
                                            v89 := GenS8{F0: (!false), F1: 0}
                                            _ = v89
                                        }
                                        var v90 int32 = 836
                                        _ = v87
                                        _ = v90
                                    }
                                    v44 += (v44 * v44)
                                    v2 = v2
                                }
                            }
                            _ = fi55
                        }
                        if false {
                            if ((float32(3.14159) * (v39 / float32(((float64(v39) - float64((1e10 + v39))) * 0.1)))) >= (0.0 + v39)) {
                                if ((true || false) || true) {
                                    if (1 == 127) {
                                        v91 := GenS8{F0: (!(-88 >= 42)), F1: 1}
                                        var v92 int = fi46
                                        _ = v91
                                        _ = v92
                                    } else {
                                        fmt.Println(fi46, v1)
                                        var v93 uint = 888
                                        var v94 uint16 = 32767
                                        _ = v93
                                        _ = v94
                                    }
                                }
                            } else {
                                for fi95 := 0; fi95 < 18; fi95++ {
                                    var v96 float64 = -2.25
                                    var v97 int8 = 100
                                    v2 = "foo bar"
                                    _ = fi95
                                    _ = v96
                                    _ = v97
                                }
                                var v98 int8 = 16
                                for fi99 := 0; fi99 < 3; fi99++ {
                                    fmt.Println(v3, v44, v39)
                                    if true {
                                        var v100 uint = 30
                                        _ = v100
                                    }
                                    _ = fi99
                                }
                                _ = v98
                            }
                            for fi101 := 0; fi101 < 15; fi101++ {
                                if (v4 != (v4 | 2147483647)) {
                                    var v102 int32 = (v4 << ((1) & 63))
                                    v103 := v5[fi101]
                                    _ = v102
                                    _ = v103
                                } else {
                                    if false {
                                        if ((18446744073709551615 <= 42) || ((((v2 + "") + "a") != "a") || true)) {
                                            var v104 uint8 = (v44 % ((128) | 1))
                                            var v105 float32 = v39
                                            _ = v104
                                            _ = v105
                                        } else {
                                            fmt.Println(v39, v44)
                                            var v106 int = (^(fi46))
                                            var v107 bool = (!(((v2 + "a") + "") != v2))
                                            _ = v106
                                            _ = v107
                                        }
                                    } else {
                                        if ((v4 - (v4 | 0)) == (v4 * (v4 | -1))) {
                                            if false {
                                                var v108 float32 = 1e10
                                                v39 = v39
                                                switch v2 {
                                                case "foo bar":
                                                    v2 = v2
                                                }
                                                if (!(42 >= 42)) {
                                                    v109 := GenS9{F0: true, F1: (v3 * v3), F2: fi46}
                                                    v110 := GenS9{F0: false, F1: v3, F2: fi46}
                                                    v4 = 2147483647
                                                    _ = v109
                                                    _ = v110
                                                } else {
                                                    v1 = (v1 & 2147483648)
                                                    v39 = v39
                                                    var v111 float32 = float32(((((((1.5 * 0.1) + 0.0) / 1.5) * (0.0 * float64(v39))) - 1.5) - (1.5 + (-2.25 / -2.25))))
                                                    _ = v111
                                                }
                                                if (-126 >= 0) {
                                                    if (v44 >= 127) {
                                                        v112 := []int{(fi101 - int(v44)), (^(fi101)), 0}
                                                        _ = v112
                                                    }
                                                }
                                                _ = v108
                                            }
                                        } else {
                                            var v113 uint16 = 2
                                            _ = v113
                                        }
                                    }
                                }
                                var v114 float64 = 1.5
                                fmt.Println(fi101, v44, v4)
                                _ = fi101
                                _ = v114
                            }
                            if false {
                                var v115 uint16 = 255
                                var v116 int64 = -1
                                _ = v115
                                _ = v116
                            }
                        }
                    }
                    v1 += (v1 >> ((202) & 63))
                }
                v117 := v5[fi46]
                _ = fi46
                _ = v117
            }
        } else {
            if true {
                v118 := GenS10{F0: 42, F1: 255, F2: (-(v4))}
                var v119 int32 = (v4 % ((v4) | 1))
                var v120 uint16 = 65535
                _ = v118
                _ = v119
                _ = v120
            } else {
                var v121 uint32 = 0
                v44 = 0
                _ = v121
            }
            for fi122 := 0; fi122 < 1; fi122++ {
                v39 += v39
                _ = fi122
            }
        }
        var v123 int64 = 0
        _ = v123
        _ = v45
    } else {
        v124 := []int{32768, 796, -1, -653}
        v44 += (^(v44))
        if (v1 <= v1) {
            v125 := GenS3{F0: 65535, F1: 255, F2: 433}
            if len(v124) > 0 {
                v124[0] = 9223372036854775807
            }
            _ = v125
        } else {
            if len(v124) > 0 {
                v124[0] = 128
            }
        }
        _ = v124
    }
    v126 := GenS11{F0: 97, F1: v3, F2: 1}
    var v127 int16 = (v3 | (v3 >> ((4294967295) & 63)))
    fmt.Println(v39, v1, v1)
    if ((v3 << ((2) & 63)) > (v3 - (-(v3)))) {
        v127 -= v127
        for fi128 := 0; fi128 < 4; fi128++ {
            v129 := GenS11{F0: 391, F1: (v127 + int16(fi128)), F2: 293}
            v130 := GenS7{F0: 9223372036854775807}
            _ = fi128
            _ = v129
            _ = v130
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v126
    _ = v127
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
