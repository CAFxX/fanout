package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 bool }
type GenS2 struct { F0 int; F1 uint64; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint32; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int8; F2 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool; F1 uint; F2 int8 }
type GenS14 struct { F0 int16; F1 int8 }
type GenS15 struct { F0 int8; F1 uint32; F2 int8 }
type GenS16 struct { F0 int16; F1 int64 }
type GenS17 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint16 = 255
    var v2 float32 = 1.5
    v3 := make(map[int]string)
    v4 := []string{"a", ("a" + "a")}
    v5 := make(map[uint32]uint32)
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 313, F1: ((true || false) && (v2 > v2)), F2: ((float64(v2) * -2.25) < 1.5)}
        _ = fi7
        _ = v8
    }
    v1 -= (v1 + (v1 | (v1 ^ v1)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch "hello" {
            case "\n\t":
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := v5[42]
                    v12 := GenS2{F0: (fi10 | fi9), F1: 18446744073709551615, F2: 637}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "hello":
                v1 = 65535
            default:
                if ((-1 >= 111) && false) {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        if true {
                            switch (fi13 | (fi9 + 2147483647)) {
                            case 839:
                                fmt.Println(fi9)
                            case 0:
                                switch fi9 {
                                case -725:
                                    if len(v4) > 0 {
                                        v4[0] = (v6 + v6)
                                    }
                                case 9223372036854775807:
                                    switch v6 {
                                    case "a":
                                        v14 := GenS3{F0: -379}
                                        _ = v14
                                    case "foo bar":
                                        var v15 uint32 = 755
                                        _ = v15
                                    }
                                case 128:
                                    v16 := GenS3{F0: 366}
                                    _ = v16
                                default:
                                    if (0 < 377) {
                                        if false {
                                            v17 := v3[fi9]
                                            if (!false) {
                                                v17 = v6
                                                v18 := GenS4{F0: 667}
                                                v1 = (^(v1))
                                                _ = v18
                                            }
                                            v19 := GenS5{F0: -2147483648, F1: 42, F2: -1}
                                            _ = v17
                                            _ = v19
                                        }
                                        if ((false || (false && ((((-473 >= 42) && false) && true) && (310 == 590)))) && (128 < 93)) {
                                            var v20 int64 = -1
                                            if (v2 == v2) {
                                                _ = len(v4)
                                                switch v6 {
                                                case "\n\t":
                                                    if (!false) {
                                                        v21 := GenS2{F0: 2147483647, F1: 18446744073709551615, F2: 783}
                                                        v22 := GenS2{F0: fi13, F1: 0, F2: 42}
                                                        fmt.Println(v2, v20, v1)
                                                        _ = v21
                                                        _ = v22
                                                    } else {
                                                        fmt.Println(v20)
                                                    }
                                                case "foo bar":
                                                    if false {
                                                        v2 -= -2.25
                                                        var v23 float32 = (1.5 / 0.1)
                                                        if (-32768 != 173) {
                                                            var v24 float32 = float32((float64((v23 + ((v23 - 0.1) * ((((v2 * v23) - v23) - (3.14159 - (v23 * (v23 + v23)))) / v2)))) * 3.14159))
                                                            fmt.Println(v23, v6)
                                                            var v25 int8 = 100
                                                            _ = v24
                                                            _ = v25
                                                        } else {
                                                            v26 := GenS6{F0: -467}
                                                            _ = v26
                                                        }
                                                        _ = v23
                                                    } else {
                                                        var v27 uint16 = v1
                                                        v6 = v6
                                                        if false {
                                                            fmt.Println(v6, v20)
                                                            switch v6 {
                                                            case "a":
                                                                var v28 string = (v6 + v6)
                                                                _ = v28
                                                            default:
                                                                var v29 uint16 = 78
                                                                _ = v29
                                                            }
                                                            v30 := v5[65535]
                                                            _ = v30
                                                        }
                                                        _ = v27
                                                    }
                                                default:
                                                    fmt.Println(v6, fi9)
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    var v31 float32 = (((1.5 - (float32(3.14159) + ((v2 / ((0.1 * (1.5 + ((1e10 / (float32(0.1) / 1.5)) * (v2 - (((((0.0 * (v2 * v2)) / v2) - (0.0 / 1.0)) + v2) - (float32((-2.25 - -2.25)) * v2)))))) + float32((1e10 - -2.25)))) - v2))) * (1.5 + v2)) / ((float32(0.1) * 1.5) - v2))
                                                    _ = v31
                                                case "":
                                                    var v32 uint8 = 1
                                                    _ = v32
                                                }
                                            }
                                            v6 = v6
                                            _ = v20
                                        } else {
                                            v6 = v6
                                        }
                                        var v33 string = (v6 + "foo bar")
                                        fmt.Println(fi13)
                                        _ = v33
                                    } else {
                                        v34 := GenS7{F0: 513, F1: 0, F2: (2147483647 < 4294967295)}
                                        switch (fi9 | 859) {
                                        case 661:
                                            v6 = v6
                                        case -320:
                                            v6 = "foo bar"
                                        case 42:
                                            var v35 float64 = ((3.14159 / (-2.25 * 0.1)) - 3.14159)
                                            _ = v35
                                        }
                                        _ = v34
                                    }
                                }
                            case 128:
                                var v36 int64 = -1
                                var v37 int32 = 2147483647
                                _ = v36
                                _ = v37
                            default:
                                if false {
                                    v6 = "a"
                                } else {
                                    v38 := GenS1{F0: 361, F1: true, F2: false}
                                    var v39 uint16 = (^(v1))
                                    if true {
                                        switch (fi13 + 4294967295) {
                                        case 1:
                                            fmt.Println(v1, v39, v1)
                                        default:
                                            v40 := GenS8{F0: -104}
                                            _ = v40
                                        }
                                    } else {
                                        v39 -= (v1 & (v1 - (v1 % ((v39) | 1))))
                                        var v41 int8 = -128
                                        v42 := GenS1{F0: 279, F1: (!(32767 <= 1)), F2: (42 > (^(v1)))}
                                        if (295 > 364) {
                                            fmt.Println(v41)
                                            var v43 int8 = (v41 / (((v41 * v41)) | 1))
                                            _ = v43
                                        } else {
                                            fmt.Println(v1, v39)
                                            v6 = v6
                                        }
                                        v44 := GenS9{F0: -918}
                                        _ = v41
                                        _ = v42
                                        _ = v44
                                    }
                                    _ = v38
                                    _ = v39
                                }
                            }
                            _ = len(v5)
                        }
                        _ = fi13
                    }
                } else {
                    if false {
                        var v45 uint = 0
                        v46 := v5[42]
                        _ = v45
                        _ = v46
                    }
                }
            }
            _ = fi9
        }
        fmt.Println(v6)
        if false {
            v2 -= (1e10 / 0.1)
        }
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        if (!false) {
            v4 = append(v4, "")
        } else {
            var v48 float32 = ((float32(-2.25) + (v2 / v2)) - (float32(-2.25) - (0.1 + 1e10)))
            _ = v48
        }
        _ = fi47
    }
    if (1 < -128) {
        if (129 <= 0) {
            switch 32768 {
            case 431:
                var v49 uint64 = 0
                _ = v49
            case -363:
                _ = len(v4)
            }
            v1 = v1
        }
        v50 := v5[0]
        _ = v50
    }
    v51 := GenS2{F0: 127, F1: 4294967295, F2: 42}
    var v52 uint64 = 0
    v2 += 0.1
    if true {
        v53 := v3[32767]
        if (99 != 127) {
            v4 = append(v4, ("\n\t" + v53))
        }
        _ = v53
    } else {
        var v54 uint8 = 228
        var v55 uint64 = 1
        var v56 int64 = -1
        _ = v54
        _ = v55
        _ = v56
    }
    v1 = v1
    switch "" {
    case "foo bar":
        if (42 == 42) {
            v52 = 2
            delete(v5, 1)
            var v57 uint64 = (v52 & 4294967295)
            _ = v57
        } else {
            var v58 uint16 = 2
            switch 0 {
            case 4294967295:
                v59 := GenS5{F0: -706, F1: 2, F2: -1}
                _ = v59
            case -2147483648:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 42:
                v6 = v6
            }
            switch (v6 + (v6 + (v6 + v6))) {
            case "\n\t":
                switch v6 {
                case "foo bar":
                    switch v6 {
                    case "a":
                        for fi60 := 0; fi60 < 2; fi60++ {
                            var v61 uint = 2
                            delete(v3, (fi60 >> ((255) & 63)))
                            v2 -= (v2 / (v2 - (1.5 / v2)))
                            _ = fi60
                            _ = v61
                        }
                    case "":
                        var v62 uint8 = 2
                        _ = v62
                    case "foo bar":
                        for fi63 := 0; fi63 < 16; fi63++ {
                            var v64 uint64 = 18446744073709551615
                            v2 = float32((0.1 * float64((0.1 * v2))))
                            var v65 int = fi63
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    default:
                        delete(v3, -2147483648)
                    }
                case "a":
                    fmt.Println(v2)
                case "":
                    v3[2147483647] = (v6 + v6)
                default:
                    for fi66 := 0; fi66 < 16; fi66++ {
                        v58 -= (v58 / ((v1) | 1))
                        switch -10 {
                        case 32767:
                            if (535 < 967) {
                                for fi67 := 0; fi67 < 3; fi67++ {
                                    var v68 int64 = 5
                                    var v69 string = v6
                                    if ((v52 % (((v52 + (^(v52)))) | 1)) > v52) {
                                        fmt.Println(v69, v2, v52)
                                        _ = len(v4)
                                    }
                                    _ = fi67
                                    _ = v68
                                    _ = v69
                                }
                                switch -299 {
                                case 32767:
                                    var v70 uint16 = (v58 << ((uint(v52)) & 63))
                                    _ = v70
                                case -1000:
                                    fmt.Println(v6, v58, v6)
                                }
                                fmt.Println(v58, v1, v52)
                            }
                        case 4294967295:
                            var v71 uint64 = v52
                            _ = v71
                        case 1:
                            v72 := GenS9{F0: int(v1)}
                            _ = v72
                        }
                        if (88 == 218) {
                            if false {
                                var v73 uint = 18446744073709551615
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (65535 != 42) {
                                    for fi74 := 0; fi74 < 7; fi74++ {
                                        switch fi66 {
                                        case 911:
                                            var v75 int = fi74
                                            _ = v75
                                        case 127:
                                            v58 -= v58
                                        }
                                        var v76 int64 = 42
                                        _ = fi74
                                        _ = v76
                                    }
                                    switch 1 {
                                    case -2147483648:
                                        v2 = 0.1
                                    case 1:
                                        var v77 uint32 = 0
                                        _ = v77
                                    }
                                    v4 = append(v4, ("foo bar" + (v6 + "")))
                                } else {
                                    for fi78 := 0; fi78 < 11; fi78++ {
                                        v2 -= ((v2 * v2) - v2)
                                        _ = len(v4)
                                        var v79 uint16 = (v58 - 2)
                                        v80 := GenS7{F0: 2, F1: -128, F2: true}
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                    }
                                    for fi81 := 0; fi81 < 9; fi81++ {
                                        fmt.Println(v1, fi81, v2)
                                        fmt.Println(v6, fi66)
                                        v58 = v58
                                        if false {
                                            v82 := GenS7{F0: 596, F1: -10, F2: true}
                                            v83 := GenS10{F0: v6, F1: 0}
                                            if (!(true && false)) {
                                                if ((1 <= -69) && (-220 == 1)) {
                                                    delete(v3, (fi66 * fi81))
                                                } else {
                                                    if (((!(2 > 42)) && false) || false) {
                                                        var v84 bool = true
                                                        if (0 <= -800) {
                                                            if v84 {
                                                                v85 := GenS3{F0: 1}
                                                                var v86 uint16 = (^(v1))
                                                                fmt.Println(fi81, fi66, v86)
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                        } else {
                                                            var v87 uint16 = v1
                                                            if true {
                                                                if (248 <= 2) {
                                                                    switch ("a" + ((("" + v6) + v6) + (("" + v6) + "hello"))) {
                                                                    case "foo bar":
                                                                        if v84 {
                                                                            v88 := GenS11{F0: 42, F1: (v73 & (v73 | (^(v73)))), F2: 360}
                                                                            v1 += v58
                                                                            fmt.Println(v87)
                                                                            _ = v88
                                                                        } else {
                                                                            v73 -= (v73 / ((65535) | 1))
                                                                            switch (^(fi66)) {
                                                                            case 32767:
                                                                                if true {
                                                                                    v58 -= v1
                                                                                } else {
                                                                                    var v89 float64 = 1e10
                                                                                    var v90 uint64 = uint64(v87)
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                }
                                                                                if v84 {
                                                                                    fmt.Println(v6, v6)
                                                                                } else {
                                                                                    var v91 uint = (v73 << (((^(v73))) & 63))
                                                                                    _ = v91
                                                                                }
                                                                                fmt.Println(fi66, fi66, v6)
                                                                            case 2147483647:
                                                                                var v92 int8 = int8(v87)
                                                                                _ = v92
                                                                            case 4294967295:
                                                                                v93 := GenS4{F0: 361}
                                                                                _ = v93
                                                                            default:
                                                                                v58 -= (v1 / ((v58) | 1))
                                                                            }
                                                                        }
                                                                    case "":
                                                                        v3[1] = ""
                                                                    default:
                                                                        v94 := GenS1{F0: 2, F1: (!v84), F2: true}
                                                                        _ = v94
                                                                    }
                                                                    if (0 != (fi66 * 127)) {
                                                                        if (!false) {
                                                                            var v95 uint16 = (v1 % ((v87) | 1))
                                                                            if ((v73 / (((v73 * (v73 % (((v73 << (((v73 << (((v73 % ((v73) | 1))) & 63))) & 63))) | 1)))) | 1)) <= (v73 ^ (v73 + v73))) {
                                                                                if (-128 >= 0) {
                                                                                    v95 = (v87 | v58)
                                                                                }
                                                                                switch v6 {
                                                                                case "foo bar":
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v6
                                                                                    }
                                                                                }
                                                                                var v96 int32 = -252
                                                                                _ = v96
                                                                            } else {
                                                                                var v97 bool = false
                                                                                var v98 uint16 = uint16(v73)
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                            _ = v95
                                                                        }
                                                                        v99 := GenS9{F0: (fi81 + fi81)}
                                                                        if (!v84) {
                                                                            v52 += v52
                                                                            if true {
                                                                                v52 = v52
                                                                                fmt.Println(v73, fi81, fi66)
                                                                                fmt.Println(v52, fi66, v2)
                                                                            } else {
                                                                                v100 := GenS12{F0: "\n\t"}
                                                                                _ = v100
                                                                            }
                                                                        } else {
                                                                            if v84 {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                v58 = v58
                                                                                v58 = (v58 ^ 32767)
                                                                                v101 := GenS1{F0: 4294967295, F1: ((2 > 2) && false), F2: v84}
                                                                                v84 = v84
                                                                                _ = v101
                                                                            }
                                                                            v1 = (v87 / ((v87) | 1))
                                                                            if v84 {
                                                                                _ = len(v4)
                                                                            } else {
                                                                                v102 := GenS12{F0: v6}
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                        var v103 bool = false
                                                                        _ = v103
                                                                        _ = v99
                                                                    } else {
                                                                        if v84 {
                                                                            fmt.Println(v6)
                                                                            if len(v4) > 0 {
                                                                                v4[0] = "hello"
                                                                            }
                                                                            switch (v6 + v6) {
                                                                            case "foo bar":
                                                                                v104 := GenS6{F0: -1}
                                                                                _ = v104
                                                                            }
                                                                            var v105 int32 = 0
                                                                            v52 = 268
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    v106 := GenS2{F0: (^(fi81)), F1: 18446744073709551615, F2: -32768}
                                                                    _ = v106
                                                                } else {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v52 += (v52 & (v52 / (((^(v52))) | 1)))
                                                                    }
                                                                    fmt.Println(v87, v84)
                                                                }
                                                                if v84 {
                                                                    v73 = v73
                                                                }
                                                                var v107 float64 = (0.0 + -2.25)
                                                                _ = v107
                                                            } else {
                                                                var v108 uint = v73
                                                                _ = v108
                                                            }
                                                            var v109 uint8 = uint8(v73)
                                                            _ = v109
                                                            _ = v87
                                                        }
                                                        v110 := v3[(fi81 & fi66)]
                                                        var v111 int32 = -1
                                                        var v112 bool = ((!((false && (v84 && v84)) && false)) || (0 != (v52 | v52)))
                                                        _ = v110
                                                        _ = v111
                                                        _ = v112
                                                        _ = v84
                                                    } else {
                                                        v2 = v2
                                                    }
                                                    v73 = v73
                                                    var v113 float64 = 0.0
                                                    _ = v113
                                                }
                                            } else {
                                                v114 := GenS5{F0: -403, F1: uint32(v58), F2: 501}
                                                var v115 float32 = (1.5 - v2)
                                                _ = v114
                                                _ = v115
                                            }
                                            v116 := GenS13{F0: ((-2.25 < 3.14159) && (!(v52 <= 2))), F1: v73, F2: 0}
                                            var v117 float32 = v2
                                            fmt.Println(fi66)
                                            v2 -= 0.1
                                            var v118 int16 = -1
                                            _ = v116
                                            _ = v117
                                            _ = v118
                                            _ = v82
                                            _ = v83
                                        } else {
                                            v58 = 1
                                            var v119 int = int(v58)
                                            _ = v119
                                        }
                                        if true {
                                            if (fi66 <= fi66) {
                                                v1 = 556
                                                if true {
                                                    switch 9223372036854775807 {
                                                    case -9223372036854775808:
                                                        var v120 int8 = 122
                                                        _ = v120
                                                    default:
                                                        var v121 uint8 = 104
                                                        _ = v121
                                                    }
                                                    v52 = 42
                                                }
                                                if (!(!(42 != -1))) {
                                                    v6 = ("\n\t" + (v6 + ("\n\t" + ("a" + ("foo bar" + (v6 + "a"))))))
                                                    v122 := GenS4{F0: v52}
                                                    _ = v122
                                                }
                                                if (v6 > ("foo bar" + v6)) {
                                                    if len(v4) > 0 {
                                                        v4[0] = ("a" + ((v6 + ("foo bar" + v6)) + "\n\t"))
                                                    }
                                                } else {
                                                    v6 = v6
                                                    var v123 uint = 65535
                                                    v1 = v1
                                                    var v124 uint16 = 32767
                                                    var v125 int16 = 1
                                                    _ = v123
                                                    _ = v124
                                                    _ = v125
                                                }
                                                var v126 int16 = int16(v52)
                                                v58 = v1
                                                v73 = (v73 | (^(v73)))
                                                _ = v126
                                            } else {
                                                switch (^(fi66)) {
                                                case 32768:
                                                    v6 = v6
                                                }
                                                if (!false) {
                                                    fmt.Println(v58, v2, v73)
                                                    var v127 string = v6
                                                    v52 = 42
                                                    _ = v127
                                                } else {
                                                    _ = len(v5)
                                                    fmt.Println(v52)
                                                    v5[2] = 4294967295
                                                }
                                                if (-32768 < -967) {
                                                    if (false || true) {
                                                        v128 := GenS14{F0: -923, F1: 42}
                                                        _ = v128
                                                    }
                                                    var v129 float64 = 0.0
                                                    _ = v129
                                                } else {
                                                    var v130 int64 = 946
                                                    _ = len(v4)
                                                    _ = v130
                                                }
                                            }
                                            v73 -= uint(v58)
                                        }
                                        _ = fi81
                                    }
                                    v52 -= v52
                                }
                                _ = v73
                            } else {
                                fmt.Println(v58, v1, v1)
                                switch v6 {
                                case "\n\t":
                                    v2 += 0.1
                                case "hello":
                                    v6 = "foo bar"
                                }
                                switch int(v58) {
                                case 0:
                                    v131 := GenS15{F0: -1, F1: 963, F2: 105}
                                    _ = v131
                                }
                            }
                        } else {
                            v132 := GenS16{F0: -215, F1: 0}
                            var v133 uint8 = 42
                            v52 += (v52 << ((2) & 63))
                            _ = v132
                            _ = v133
                        }
                        _ = fi66
                    }
                }
            }
            _ = v58
        }
    case "":
        if (false && false) {
            switch v6 {
            case "a":
                for fi134 := 0; fi134 < 3; fi134++ {
                    fmt.Println(v2, v52, v52)
                    var v135 int64 = 1
                    switch v6 {
                    case "":
                        if (true && false) {
                            fmt.Println(v52, v1, v2)
                        }
                    default:
                        delete(v5, uint32(v1))
                    }
                    _ = fi134
                    _ = v135
                }
            case "foo bar":
                var v136 int64 = 694
                _ = v136
            case "\n\t":
                for fi137 := 0; fi137 < 6; fi137++ {
                    var v138 bool = (!true)
                    _ = fi137
                    _ = v138
                }
            }
            v139 := GenS12{F0: "hello"}
            var v140 uint16 = (v1 - v1)
            _ = v139
            _ = v140
        } else {
            for fi141 := 0; fi141 < 13; fi141++ {
                var v142 uint8 = 1
                if (((0.1 / v2) + 0.1) >= v2) {
                    v143 := GenS17{F0: 42, F1: 703}
                    _ = v143
                }
                _ = fi141
                _ = v142
            }
        }
    case "hello":
        fmt.Println(v2, v6)
    default:
        fmt.Println(v1, v52, v2)
    }
    v52 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
