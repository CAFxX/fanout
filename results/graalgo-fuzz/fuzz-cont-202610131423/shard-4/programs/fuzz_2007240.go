package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint; F1 bool; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint32; F2 string }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint64; F1 int }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint16; F1 int16; F2 int32 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int16; F1 int8; F2 uint32 }
type GenS17 struct { F0 uint64; F1 uint }
type GenS18 struct { F0 int32; F1 int }
func genFn0(p0 uint32, p1 uint, p2 uint8) uint {
    for fi5 := 0; fi5 < 9; fi5++ {
        p0 = p0
        v6 := GenS1{F0: 508}
        _ = fi5
        _ = v6
    }
    p1 = (p1 + (p1 << ((p1) & 63)))
    var v7 uint64 = 4294967295
    switch "" {
    case "":
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int8 = 1
            var v10 uint64 = (v7 >> (((p1 + 65535)) & 63))
            _ = fi8
            _ = v10
            _ = v9
        }
    case "\n\t":
        p2 = 239
    case "foo bar":
        p1 -= p1
    default:
        var v11 uint = (p1 ^ p1)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 | (p1 + p1))
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 uint64 = 436
    if (0 < 65535) {
        v4 += v4
    }
    if ((("" + "foo bar") + "") == ((("" + "") + "a") + "foo bar")) {
        v12 := []int{-9223372036854775808, 27}
        v3[42] = ("" + "\n\t")
        switch -711 {
        case 127:
            if false {
                fmt.Println(v4)
                switch -669 {
                case -621:
                    _ = len(v3)
                case -1:
                    v1 = -1
                case 1:
                    v4 -= v4
                default:
                    var v13 uint32 = 535
                    _ = v13
                }
                fmt.Println(v4, v1)
            }
        }
        _ = v12
    } else {
        v1 += v1
        if (true || true) {
            v4 -= (v4 | (v4 | (^(v4))))
            v14 := GenS1{F0: 18446744073709551615}
            _ = v14
        }
        v2[0] = 2
    }
    v4 = (v4 / ((v4) | 1))
    for fi15 := 0; fi15 < 7; fi15++ {
        v1 = 33
        for fi16 := 0; fi16 < 16; fi16++ {
            if true {
                fmt.Println(fi16, v1, v1)
                v1 = -63
                fmt.Println(v1)
                if false {
                    v4 += v4
                    v17 := GenS1{F0: v4}
                    v18 := GenS1{F0: v4}
                    _ = v17
                    _ = v18
                } else {
                    v19 := []int{(fi16 - fi16), fi15, 32767, (fi15 / ((-2147483648) | 1))}
                    v20 := GenS1{F0: (v4 - (^(v4)))}
                    _ = v19
                    _ = v20
                }
                var v21 int8 = (^(v1))
                v4 -= (^(v4))
                _ = v21
            } else {
                v22 := []int{fi15, (fi15 ^ fi15)}
                var v23 uint32 = 42
                if ("\n\t" <= (((("a" + ("foo bar" + ("a" + "\n\t"))) + "a") + "hello") + ("hello" + ("foo bar" + "\n\t")))) {
                    fmt.Println(fi15)
                } else {
                    var v24 float32 = (((float32(1.5) - 0.1) - float32(1e10)) / 1.0)
                    _ = v24
                }
                _ = v22
                _ = v23
            }
            v25 := GenS2{F0: 286}
            _ = fi16
            _ = v25
        }
        v3[42] = "foo bar"
        _ = fi15
    }
    v26 := []int{}
    for fi27 := 0; fi27 < 2; fi27++ {
        if false {
            for fi28 := 0; fi28 < 10; fi28++ {
                v1 = v1
                v29 := GenS3{F0: 271}
                switch fi28 {
                case 219:
                    v1 += v1
                case -9223372036854775808:
                    v30 := GenS4{F0: 42, F1: 116}
                    _ = v30
                case 740:
                    var v31 uint32 = 2147483647
                    _ = v31
                }
                var v32 int8 = 0
                if true {
                    var v33 int16 = 1000
                    v34 := GenS1{F0: v4}
                    delete(v3, 42)
                    _ = v33
                    _ = v34
                } else {
                    if ((v4 >> ((826) & 63)) == v4) {
                        var v35 float32 = 0.1
                        var v36 uint8 = 136
                        _ = v35
                        _ = v36
                    }
                }
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = fi28
                _ = v29
                _ = v32
            }
        } else {
            if ((0.1 / 1e10) <= -2.25) {
                v4 -= v4
                fmt.Println(fi27)
            } else {
                _ = len(v26)
                switch ("hello" + "a") {
                case "hello":
                    v37 := v3[2]
                    _ = v37
                case "\n\t":
                    var v38 int64 = 9223372036854775807
                    _ = v38
                case "a":
                    var v39 int32 = 758
                    _ = v39
                default:
                    var v40 uint64 = 211
                    _ = v40
                }
                if (-486 != -2147483648) {
                    if (false && (!(262 < 2))) {
                        var v41 uint8 = 127
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            var v43 int = -2147483648
                            v43 += (-(fi42))
                            _ = fi42
                            _ = v43
                        }
                        v1 = (v1 / (((v1 & (v1 << ((255) & 63)))) | 1))
                    }
                    var v44 bool = (9223372036854775807 > 1)
                    var v45 uint16 = 32767
                    _ = v44
                    _ = v45
                }
            }
        }
        if (false || (2147483647 > 42)) {
            v46 := GenS2{F0: 42}
            _ = v46
        } else {
            var v47 uint16 = 2
            switch fi27 {
            case -869:
                if true {
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    for fi48 := 0; fi48 < 5; fi48++ {
                        var v49 uint = 255
                        if true {
                            fmt.Println(fi48, fi48, v1)
                            v50 := GenS4{F0: -960, F1: 128}
                            if (!false) {
                                v51 := GenS5{F0: (^(v49)), F1: true, F2: 987}
                                _ = v51
                            }
                            _ = v50
                        } else {
                            if true {
                                if (("hello" <= "a") && true) {
                                    switch (fi27 >> (((v49 | v49)) & 63)) {
                                    case 32768:
                                        v49 += (v49 | v49)
                                    case -221:
                                        v4 += 1
                                    }
                                    v4 = v4
                                    var v52 bool = (v47 > 255)
                                    _ = v52
                                } else {
                                    var v53 uint16 = v47
                                    var v54 uint8 = 2
                                    _ = v53
                                    _ = v54
                                }
                                _ = len(v3)
                                v26 = append(v26, (^(fi27)))
                            }
                            fmt.Println(v47, v1)
                            if true {
                                v55 := v3[175]
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                _ = v55
                            }
                        }
                        _ = fi48
                        _ = v49
                    }
                }
            case 9223372036854775807:
                var v56 float64 = float64(((((3.14159 * (-2.25 - 3.14159)) - 1.5) * -2.25) / 1.0))
                _ = v56
            case -9223372036854775808:
                if true {
                    fmt.Println(v1, v47, v1)
                    var v57 string = "a"
                    for fi58 := 0; fi58 < 8; fi58++ {
                        if true {
                            v59 := GenS5{F0: 994, F1: false, F2: -1}
                            if false {
                                if true {
                                    v60 := GenS6{F0: 607, F1: 42}
                                    var v61 int16 = 472
                                    v61 = -1
                                    delete(v2, 65535)
                                    delete(v2, 2147483648)
                                    _ = v60
                                    _ = v61
                                }
                                var v62 uint8 = 42
                                _ = v62
                            }
                            _ = v59
                        } else {
                            if true {
                                v1 = -1
                                v63 := GenS4{F0: -2147483648, F1: 1}
                                switch (fi58 * fi58) {
                                case -2147483648:
                                    if ((!true) && true) {
                                        var v64 uint16 = 968
                                        _ = v64
                                    }
                                case 0:
                                    v47 += 0
                                }
                                switch 127 {
                                case 114:
                                    fmt.Println(v57)
                                case -711:
                                    v65 := GenS3{F0: -565}
                                    _ = v65
                                case 244:
                                    v47 = (v47 % (((v47 / ((v47) | 1))) | 1))
                                default:
                                    v57 = ""
                                }
                                _ = v63
                            }
                        }
                        v47 = v47
                        _ = fi58
                    }
                    _ = v57
                }
            }
            _ = v47
        }
        v1 = (v1 * (v1 ^ v1))
        _ = fi27
    }
    for fi66 := 0; fi66 < 10; fi66++ {
        for fi67 := 0; fi67 < 16; fi67++ {
            switch fi67 {
            case 128:
                v4 = (v4 % (((v4 ^ v4)) | 1))
                v68 := GenS7{F0: 501, F1: 4294967295, F2: ("" + ("foo bar" + (("\n\t" + "") + "")))}
                _ = v68
            case -938:
                v69 := GenS5{F0: 255, F1: (!false), F2: 42}
                _ = v69
            }
            v26 = append(v26, fi67)
            switch "\n\t" {
            case "foo bar":
                v70 := GenS3{F0: -1}
                fmt.Println(fi67, v1)
                _ = v70
            }
            _ = fi67
        }
        v4 += (^(v4))
        _ = len(v3)
        _ = fi66
    }
    var v71 string = "hello"
    if ((v1 / (((v1 % ((1) | 1))) | 1)) != (v1 - int8(v4))) {
        v72 := GenS8{F0: true, F1: (v4 & (v4 - 42))}
        switch v71 {
        case "":
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        default:
            if false {
                if false {
                    _ = len(v26)
                    var v73 int16 = -153
                    _ = v73
                } else {
                    v71 = "hello"
                    v1 -= v1
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                }
                switch 4294967295 {
                case 42:
                    if (0 > 1) {
                        v71 = (v71 + (v71 + v71))
                        var v74 uint32 = 2147483648
                        _ = v74
                    }
                default:
                    _ = len(v26)
                }
            } else {
                var v75 uint32 = 1
                v76 := GenS9{F0: 2}
                _ = v75
                _ = v76
            }
        }
        v1 -= (v1 / ((v1) | 1))
        _ = v72
    } else {
        var v77 int64 = 9223372036854775807
        for fi78 := 0; fi78 < 14; fi78++ {
            if ((2 > 867) || false) {
                v1 = (v1 >> ((2) & 63))
                if true {
                    var v79 int8 = (v1 ^ (v1 & 100))
                    _ = v79
                } else {
                    if (65535 <= 0) {
                        if false {
                            v80 := GenS10{F0: v4, F1: (fi78 % ((int(v1)) | 1))}
                            if true {
                                for fi81 := 0; fi81 < 10; fi81++ {
                                    var v82 int32 = -1
                                    if true {
                                        var v83 uint8 = 255
                                        v84 := GenS11{F0: 255}
                                        _ = v83
                                        _ = v84
                                    } else {
                                        if (3.14159 <= ((1.5 + float64(-2.25)) - 0.1)) {
                                            var v85 string = ((v71 + v71) + (v71 + "\n\t"))
                                            var v86 uint64 = 2
                                            v77 -= v77
                                            _ = v85
                                            _ = v86
                                        } else {
                                            var v87 float64 = (1e10 * 0.1)
                                            _ = v87
                                        }
                                        if true {
                                            v71 = (v71 + ("hello" + ("a" + (((v71 + "foo bar") + v71) + ("hello" + v71)))))
                                            v88 := GenS12{F0: 2}
                                            if len(v26) >= 2 {
                                                v26 = v26[1:len(v26)]
                                            }
                                            _ = v88
                                        } else {
                                            var v89 int64 = -879
                                            v71 = "foo bar"
                                            _ = v89
                                        }
                                        if ((v77 | v77) < 9223372036854775807) {
                                            if len(v26) >= 2 {
                                                v26 = v26[1:len(v26)]
                                            }
                                        } else {
                                            v82 = (^(v82))
                                        }
                                    }
                                    var v90 int8 = (v1 | (v1 ^ v1))
                                    _ = fi81
                                    _ = v82
                                    _ = v90
                                }
                                v91 := GenS10{F0: v4, F1: (fi78 << ((4294967295) & 63))}
                                _ = v91
                            } else {
                                fmt.Println(fi78, v77)
                                switch 0 {
                                case -946:
                                    var v92 int64 = v77
                                    _ = v92
                                case -2147483648:
                                    var v93 string = "foo bar"
                                    _ = v93
                                }
                            }
                            _ = v80
                        } else {
                            switch v71 {
                            case "foo bar":
                                v4 = v4
                            case "hello":
                                v4 += v4
                            case "":
                                v4 = v4
                            default:
                                if false {
                                    switch "hello" {
                                    case "a":
                                        v77 = (v77 + v77)
                                    case "":
                                        var v94 int8 = v1
                                        _ = v94
                                    default:
                                        v95 := GenS13{F0: 32767, F1: -1, F2: -252}
                                        _ = v95
                                    }
                                    if (540 > 1) {
                                        for fi96 := 0; fi96 < 18; fi96++ {
                                            v97 := GenS9{F0: 32767}
                                            var v98 float64 = float64(0.0)
                                            var v99 int64 = v77
                                            if false {
                                                v100 := GenS11{F0: 106}
                                                _ = v100
                                            }
                                            var v101 bool = ((0 >= v77) && false)
                                            var v102 int16 = 1
                                            _ = fi96
                                            _ = v101
                                            _ = v102
                                            _ = v97
                                            _ = v98
                                            _ = v99
                                        }
                                        v103 := GenS11{F0: 32767}
                                        _ = v103
                                    } else {
                                        switch v71 {
                                        case "hello":
                                            for fi104 := 0; fi104 < 20; fi104++ {
                                                fmt.Println(fi104)
                                                v1 = 127
                                                _ = fi104
                                            }
                                        case "\n\t":
                                            var v105 uint64 = 4294967295
                                            _ = v105
                                        default:
                                            v71 = ((("" + v71) + "foo bar") + "\n\t")
                                        }
                                    }
                                    for fi106 := 0; fi106 < 10; fi106++ {
                                        v77 = 396
                                        switch -2147483648 {
                                        case -822:
                                            v77 -= v77
                                        case 1:
                                            switch ("\n\t" + ("" + "\n\t")) {
                                            case "hello":
                                                v107 := GenS14{F0: (-862 <= -1)}
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if (4294967295 < 1) {
                                                        fmt.Println(v4, fi78, fi106)
                                                    } else {
                                                        fmt.Println(fi106, v1)
                                                    }
                                                case "hello":
                                                    v108 := GenS15{F0: 210}
                                                    _ = v108
                                                case "":
                                                    if true {
                                                        var v109 string = (v71 + v71)
                                                        _ = v109
                                                    } else {
                                                        switch fi78 {
                                                        case 4294967295:
                                                            var v110 int8 = v1
                                                            _ = v110
                                                        case -338:
                                                            if false {
                                                                if (497 > v4) {
                                                                    if (false || true) {
                                                                        _ = len(v26)
                                                                        var v111 uint32 = 2147483647
                                                                        _ = v111
                                                                    } else {
                                                                        v112 := GenS15{F0: 42}
                                                                        if ("a" > v71) {
                                                                            if (-32768 != 1) {
                                                                                if false {
                                                                                    _ = len(v26)
                                                                                    delete(v2, 133)
                                                                                    if (0 >= -504) {
                                                                                        if len(v26) > 0 {
                                                                                            v26[0] = (fi106 / (((fi106 % (((fi106 / ((fi78) | 1))) | 1))) | 1))
                                                                                        }
                                                                                    } else {
                                                                                        if (1000 <= 42) {
                                                                                            fmt.Println(v4, v77)
                                                                                            fmt.Println(v4)
                                                                                            var v113 float64 = 0.1
                                                                                            _ = v113
                                                                                        } else {
                                                                                            v71 = ((v71 + v71) + (("a" + "") + (v71 + v71)))
                                                                                            v1 = -128
                                                                                            v114 := GenS16{F0: -1, F1: (v1 & (v1 % ((-1) | 1))), F2: 4294967295}
                                                                                            _ = v114
                                                                                        }
                                                                                        if false {
                                                                                            if len(v26) >= 2 {
                                                                                                v26 = v26[1:len(v26)]
                                                                                            }
                                                                                            var v115 int64 = 965
                                                                                            fmt.Println(v71, fi78)
                                                                                            _ = v115
                                                                                        } else {
                                                                                            if (((fi78 + (fi78 + fi78)) < fi106) || ((fi78 >> ((65535) & 63)) >= (fi106 | fi106))) {
                                                                                                v116 := GenS1{F0: (v4 + v4)}
                                                                                                if false {
                                                                                                    var v117 uint8 = 42
                                                                                                    var v118 string = v71
                                                                                                    var v119 int16 = 0
                                                                                                    _ = v117
                                                                                                    _ = v118
                                                                                                    _ = v119
                                                                                                } else {
                                                                                                    v1 = v1
                                                                                                    switch "a" {
                                                                                                    case "":
                                                                                                        var v120 int64 = (v77 & -9223372036854775808)
                                                                                                        _ = v120
                                                                                                    case "a":
                                                                                                        v77 = v77
                                                                                                    case "foo bar":
                                                                                                        v121 := GenS9{F0: 65535}
                                                                                                        _ = v121
                                                                                                    }
                                                                                                }
                                                                                                _ = v116
                                                                                            } else {
                                                                                                var v122 float32 = (0.1 + 3.14159)
                                                                                                v71 = (v71 + "hello")
                                                                                                fmt.Println(fi106, v1)
                                                                                                if len(v26) >= 2 {
                                                                                                    v26 = v26[1:len(v26)]
                                                                                                }
                                                                                                v123 := GenS17{F0: v4, F1: 4294967295}
                                                                                                switch fi78 {
                                                                                                case 128:
                                                                                                    v124 := GenS5{F0: 255, F1: (502 <= 0), F2: v77}
                                                                                                    _ = v124
                                                                                                case 127:
                                                                                                    v125 := GenS12{F0: 144}
                                                                                                    _ = v125
                                                                                                default:
                                                                                                    v71 = v71
                                                                                                }
                                                                                                _ = v122
                                                                                                _ = v123
                                                                                            }
                                                                                            var v126 uint64 = (v4 - v4)
                                                                                            var v127 uint16 = 1
                                                                                            _ = v126
                                                                                            _ = v127
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    v71 = ("" + ("\n\t" + (("\n\t" + v71) + v71)))
                                                                                    if true {
                                                                                        v1 = -128
                                                                                        fmt.Println(v4, fi106)
                                                                                        if false {
                                                                                            fmt.Println(v1)
                                                                                            v77 = v77
                                                                                            v128 := GenS8{F0: true, F1: 4294967295}
                                                                                            _ = v128
                                                                                        }
                                                                                    } else {
                                                                                        if true {
                                                                                            var v129 uint64 = v4
                                                                                            _ = v129
                                                                                        }
                                                                                        v2[4294967295] = 2147483648
                                                                                        v1 = 42
                                                                                        v130 := GenS17{F0: (v4 ^ (v4 + (v4 | v4))), F1: 255}
                                                                                        v71 = v71
                                                                                        if len(v26) > 0 {
                                                                                            v26[0] = 128
                                                                                        }
                                                                                        _ = v130
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi106)
                                                                                }
                                                                                v131 := v3[820]
                                                                                _ = v131
                                                                            }
                                                                            v132 := GenS8{F0: ((v1 + v1) != (v1 + 100)), F1: (v4 / ((v4) | 1))}
                                                                            _ = v132
                                                                        }
                                                                        v4 -= 42
                                                                        _ = v112
                                                                    }
                                                                    v71 = v71
                                                                    fmt.Println(v4, v4, v71)
                                                                    v133 := GenS15{F0: 42}
                                                                    var v134 int64 = v77
                                                                    if len(v26) >= 2 {
                                                                        v26 = v26[1:len(v26)]
                                                                    }
                                                                    switch "a" {
                                                                    case "a":
                                                                        v135 := GenS11{F0: 42}
                                                                        _ = v135
                                                                    case "\n\t":
                                                                        v134 += (v134 % ((int64(v1)) | 1))
                                                                    }
                                                                    fmt.Println(fi106, fi106, v77)
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                                var v136 float64 = 1.5
                                                                _ = v136
                                                            }
                                                        default:
                                                            v4 = 0
                                                        }
                                                        fmt.Println(v77)
                                                        if (false && (!((!(((^(v77)) > v77) && (float32(1e10) <= 0.1))) || ((v4 % ((v4) | 1)) == v4)))) {
                                                            v71 = v71
                                                            var v137 uint32 = 2
                                                            _ = v137
                                                        } else {
                                                            v71 = (("foo bar" + v71) + "\n\t")
                                                            v77 -= (v77 / ((v77) | 1))
                                                        }
                                                    }
                                                default:
                                                    v71 = "foo bar"
                                                }
                                                _ = v107
                                            case "a":
                                                v77 += 0
                                            case "\n\t":
                                                fmt.Println(fi106, fi106, v1)
                                            }
                                            v138 := v2[81]
                                            _ = v138
                                        case -47:
                                            v77 = v77
                                        }
                                        v139 := GenS18{F0: 758, F1: -838}
                                        switch (fi78 * 32767) {
                                        case 394:
                                            var v140 uint16 = 871
                                            _ = v140
                                        case 42:
                                            var v141 uint16 = 846
                                            _ = v141
                                        }
                                        v4 += v4
                                        _ = fi106
                                        _ = v139
                                    }
                                } else {
                                    _ = len(v26)
                                }
                            }
                            delete(v3, 872)
                            var v142 uint16 = 73
                            _ = v142
                        }
                    }
                    var v143 uint64 = (^(v4))
                    _ = v143
                }
            } else {
                var v144 uint16 = 2
                _ = v144
            }
            v2[684] = 2147483648
            _ = fi78
        }
        var v145 uint16 = 42
        _ = v145
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v71
}
