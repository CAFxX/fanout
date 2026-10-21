package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int8; F2 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 uint8; F1 uint; F2 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32; F1 int64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint64; F1 uint16 }
type GenS16 struct { F0 string; F1 int }
type GenS17 struct { F0 int16; F1 uint64; F2 bool }
type GenS18 struct { F0 uint }
func genFn0(p0 uint32, p1 uint8, p2 bool) int32 {
    var v5 uint16 = 240
    v6 := GenS1{F0: p0, F1: 632}
    v7 := GenS1{F0: (p0 / (((p0 ^ 2147483648)) | 1)), F1: 2147483647}
    var v8 float64 = (0.1 + ((1e10 / 1.5) * 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint) bool {
    var v9 uint32 = 0
    var v10 uint = (^(p0))
    _ = p0
    _ = v10
    _ = v9
    return ("foo bar" <= ("foo bar" + ""))
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    var v4 uint64 = 162
    if (0.1 <= (1.5 / 1e10)) {
        v11 := GenS1{F0: 299, F1: 127}
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                switch fi13 {
                case 128:
                    var v14 bool = true
                    _ = v14
                }
                switch ("hello" + ("" + "a")) {
                case "\n\t":
                    var v15 int8 = 102
                    _ = v15
                case "foo bar":
                    v16 := GenS1{F0: 652, F1: fi13}
                    if true {
                        if (v1 > v1) {
                            switch "a" {
                            case "foo bar":
                                if false {
                                    v1 -= 1
                                    v17 := GenS2{F0: 2147483648}
                                    _ = v17
                                } else {
                                    v18 := GenS3{F0: "a"}
                                    _ = v18
                                }
                                v4 = 538
                                fmt.Println(v1, fi13)
                            case "":
                                v2[1] = 65535
                            case "a":
                                switch (fi12 >> ((874) & 63)) {
                                case -832:
                                    v19 := GenS3{F0: ("\n\t" + "")}
                                    _ = v19
                                case 32768:
                                    fmt.Println(v4, v4)
                                case 2147483647:
                                    if false {
                                        if (!(true && (42 < 1))) {
                                            v4 = v4
                                        }
                                    }
                                default:
                                    var v20 int16 = -32768
                                    _ = v20
                                }
                            default:
                                v4 = 0
                            }
                            v1 = v1
                        }
                        if len(v3) > 0 {
                            v3[0] = fi12
                        }
                        if len(v3) > 0 {
                            v3[0] = (fi12 | fi12)
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = fi13
                        }
                    }
                    var v21 string = "foo bar"
                    _ = v16
                    _ = v21
                }
                _ = fi13
            }
            switch "\n\t" {
            case "a":
                var v22 uint64 = 2
                _ = v22
            case "foo bar":
                if false {
                    v4 = 1
                } else {
                    var v23 uint64 = (^(v4))
                    _ = v23
                }
            case "\n\t":
                fmt.Println(v1)
            }
            switch -2147483648 {
            case 4294967295:
                fmt.Println(fi12, fi12)
            case 730:
                switch "a" {
                case "a":
                    _ = len(v2)
                }
            case 127:
                switch fi12 {
                case -9223372036854775808:
                    if (-9223372036854775808 <= 9223372036854775807) {
                        var v24 float32 = 1.5
                        var v25 float32 = 3.14159
                        _ = v24
                        _ = v25
                    }
                case 128:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        var v27 int8 = 0
                        _ = fi26
                        _ = v27
                    }
                case 0:
                    v3 = append(v3, (fi12 - (fi12 + 1)))
                }
            default:
                if true {
                    var v28 int = fi12
                    _ = v28
                } else {
                    v4 = 4294967295
                    v4 = v4
                    var v29 uint16 = 961
                    _ = v29
                }
            }
            _ = fi12
        }
        if ((749 != 174) && (0.0 > 1e10)) {
            var v30 int = -9223372036854775808
            fmt.Println(v1)
            _ = v30
        }
        _ = v11
    }
    var v31 uint32 = 4294967295
    if ((3.14159 < (0.0 * -2.25)) && (genFn0(v31, 127, ("" >= "foo bar")) >= -2147483648)) {
        v3 = append(v3, -395)
        if false {
            var v32 float64 = 3.14159
            v33 := v2[537]
            if (false || true) {
                switch "\n\t" {
                case "a":
                    _ = len(v3)
                default:
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v31 += 1
                        if false {
                            var v35 int8 = 5
                            _ = v35
                        }
                        v36 := GenS1{F0: (v31 * (^(v33))), F1: 0}
                        _ = fi34
                        _ = v36
                    }
                }
                if true {
                    if true {
                        _ = len(v2)
                        v37 := GenS4{F0: 42}
                        _ = v37
                    }
                    v33 += (v33 >> ((0) & 63))
                    if (-1 >= 810) {
                        var v38 uint16 = 32767
                        if ((v31 << ((65535) & 63)) != (^(v33))) {
                            v32 += float64((((((3.14159 * 0.1) / -2.25) + 1.5) + ((3.14159 - float32((v32 + v32))) / 1e10)) * float32(v32)))
                            if (((("hello" + "\n\t") + "foo bar") + "a") >= "") {
                                switch 0 {
                                case -399:
                                    for fi39 := 0; fi39 < 7; fi39++ {
                                        if false {
                                            v38 += (^(v38))
                                            if false {
                                                var v40 int8 = 71
                                                for fi41 := 0; fi41 < 3; fi41++ {
                                                    if ((((1.5 / (1.5 / (1e10 + 1.5))) / 1e10) > -2.25) || (0.0 <= float32((v32 / 3.14159)))) {
                                                        if true {
                                                            if (42 <= uint8(v31)) {
                                                                v32 = ((v32 / v32) * -2.25)
                                                                v38 = 42
                                                            }
                                                        } else {
                                                            var v42 uint16 = v38
                                                            if false {
                                                                if (!false) {
                                                                    var v43 uint8 = 158
                                                                    var v44 uint = 4294967295
                                                                    if (0 < v1) {
                                                                        var v45 int8 = v40
                                                                        _ = v45
                                                                    } else {
                                                                        fmt.Println(v42, v33, v43)
                                                                        v46 := GenS5{F0: (v43 - (v43 / (((^(v43))) | 1))), F1: v40, F2: "a"}
                                                                        _ = v46
                                                                    }
                                                                    v47 := GenS1{F0: v31, F1: (fi41 | (-(fi41)))}
                                                                    _ = v43
                                                                    _ = v44
                                                                    _ = v47
                                                                } else {
                                                                    if true {
                                                                        v38 = (v42 + (^(v42)))
                                                                        v40 = (v40 << ((227) & 63))
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                    }
                                                                    v48 := GenS6{F0: 42}
                                                                    _ = v48
                                                                }
                                                            } else {
                                                                var v49 int16 = (-(v1))
                                                                _ = v49
                                                            }
                                                            _ = v42
                                                        }
                                                    }
                                                    delete(v2, uint32(v4))
                                                    v2[(v31 | v33)] = (v33 + (v33 & 110))
                                                    _ = fi41
                                                }
                                                _ = v40
                                            } else {
                                                fmt.Println(v4, fi39)
                                                v32 = v32
                                                fmt.Println(v38, v4)
                                            }
                                        }
                                        _ = fi39
                                    }
                                default:
                                    v50 := GenS7{F0: -2147483648}
                                    _ = v50
                                }
                            } else {
                                if false {
                                    var v51 uint16 = (v38 >> ((2) & 63))
                                    v32 -= v32
                                    v52 := GenS8{F0: 127, F1: 140, F2: v51}
                                    _ = v51
                                    _ = v52
                                }
                                if false {
                                    v53 := GenS9{F0: (v4 % ((v4) | 1))}
                                    if (1 >= 42) {
                                        var v54 int32 = -1
                                        _ = len(v3)
                                        _ = len(v2)
                                        _ = v54
                                    }
                                    _ = v53
                                } else {
                                    if false {
                                        for fi55 := 0; fi55 < 10; fi55++ {
                                            v31 -= v31
                                            var v56 uint8 = 42
                                            _ = fi55
                                            _ = v56
                                        }
                                    } else {
                                        var v57 int = -187
                                        _ = v57
                                    }
                                    v32 = v32
                                    var v58 int32 = -127
                                    _ = v58
                                }
                                v2[(^(v33))] = (v31 & 42)
                            }
                            var v59 int64 = -1
                            _ = v59
                        } else {
                            if (v1 <= (v1 + v1)) {
                                var v60 int16 = (v1 + (^(v1)))
                                _ = v60
                            } else {
                                for fi61 := 0; fi61 < 16; fi61++ {
                                    v31 = 2147483648
                                    fmt.Println(v38)
                                    _ = fi61
                                }
                                for fi62 := 0; fi62 < 18; fi62++ {
                                    v63 := GenS3{F0: "foo bar"}
                                    delete(v2, (^(v31)))
                                    var v64 float32 = 0.1
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                            }
                            switch -150 {
                            case 4294967295:
                                v38 = 155
                            default:
                                for fi65 := 0; fi65 < 16; fi65++ {
                                    fmt.Println(v4, v1, v32)
                                    v33 -= (^(v31))
                                    for fi66 := 0; fi66 < 10; fi66++ {
                                        var v67 bool = false
                                        var v68 int8 = 87
                                        v1 -= (v1 >> ((18446744073709551615) & 63))
                                        v3 = append(v3, fi66)
                                        _ = len(v2)
                                        _ = fi66
                                        _ = v67
                                        _ = v68
                                    }
                                    _ = fi65
                                }
                            }
                        }
                        _ = v38
                    }
                } else {
                    switch 42 {
                    case 1:
                        if (false || (681 > uint16(v4))) {
                            if true {
                                var v69 uint32 = v31
                                _ = v69
                            } else {
                                fmt.Println(v32)
                                fmt.Println(v32)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            switch (("\n\t" + "") + ("foo bar" + "foo bar")) {
                            case "a":
                                fmt.Println(v33, v4)
                            case "foo bar":
                                v33 = 4294967295
                            case "":
                                v1 += 32767
                            default:
                                v4 = (v4 * (v4 % (((v4 + v4)) | 1)))
                            }
                        }
                    case 32767:
                        if (42 >= int32(v4)) {
                            if (true || false) {
                                if (-1 >= -111) {
                                    for fi70 := 0; fi70 < 1; fi70++ {
                                        for fi71 := 0; fi71 < 12; fi71++ {
                                            if false {
                                                var v72 bool = true
                                                v73 := GenS10{F0: 0, F1: fi71}
                                                v74 := GenS9{F0: v4}
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            } else {
                                                v75 := GenS1{F0: v31, F1: 32768}
                                                _ = len(v2)
                                                switch 591 {
                                                case 2147483647:
                                                    var v76 float32 = (0.0 * 1.5)
                                                    _ = v76
                                                case 911:
                                                    v31 = 2147483648
                                                case 521:
                                                    var v77 int64 = -9223372036854775808
                                                    _ = v77
                                                default:
                                                    v78 := GenS11{F0: 119, F1: 18446744073709551615, F2: fi71}
                                                    _ = v78
                                                }
                                                _ = v75
                                            }
                                            v32 = v32
                                            var v79 int64 = -9223372036854775808
                                            v33 = (v31 - 65535)
                                            delete(v2, 122)
                                            _ = fi71
                                            _ = v79
                                        }
                                        if false {
                                            v4 -= (v4 / (((v4 * v4)) | 1))
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v80 uint16 = 2
                                            _ = v80
                                        }
                                        _ = fi70
                                    }
                                } else {
                                    var v81 int64 = -1
                                    if (42 < 134) {
                                        if ((v31 >> ((0) & 63)) <= v33) {
                                            var v82 int = -1
                                            for fi83 := 0; fi83 < 20; fi83++ {
                                                if true {
                                                    for fi84 := 0; fi84 < 16; fi84++ {
                                                        if true {
                                                            switch "a" {
                                                            case "":
                                                                if (!false) {
                                                                    delete(v2, 806)
                                                                    switch -1 {
                                                                    case 128:
                                                                        v2[911] = v33
                                                                    case 32767:
                                                                        v31 -= (^(v31))
                                                                    case 42:
                                                                        if true {
                                                                            v82 = 128
                                                                            v33 -= (v31 - (v31 - v33))
                                                                        } else {
                                                                            if true {
                                                                                v32 += v32
                                                                            }
                                                                            var v85 float64 = v32
                                                                            var v86 int = -981
                                                                            _ = v85
                                                                            _ = v86
                                                                        }
                                                                    default:
                                                                        fmt.Println(v31)
                                                                    }
                                                                } else {
                                                                    if ((v4 & (v4 - 4294967295)) >= (v4 + (v4 ^ 974))) {
                                                                        v31 = v33
                                                                        var v87 int16 = v1
                                                                        _ = v87
                                                                    }
                                                                    switch -104 {
                                                                    case 42:
                                                                        v1 = v1
                                                                    case 587:
                                                                        var v88 int8 = 1
                                                                        _ = v88
                                                                    default:
                                                                        v32 -= v32
                                                                    }
                                                                }
                                                            case "hello":
                                                                v89 := GenS12{F0: 1}
                                                                _ = v89
                                                            default:
                                                                fmt.Println(v81, v1, v32)
                                                            }
                                                        }
                                                        var v90 bool = (1.5 < 0.1)
                                                        v1 = v1
                                                        switch int(v81) {
                                                        case -625:
                                                            v91 := GenS13{F0: (v33 ^ v33), F1: (v81 & (v81 - v81))}
                                                            if ((200 == 1) || false) {
                                                                var v92 float32 = 3.14159
                                                                _ = v92
                                                            } else {
                                                                v82 = (v82 * fi84)
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v91
                                                        case -267:
                                                            v82 = -9223372036854775808
                                                        }
                                                        if v90 {
                                                            v4 += (v4 % (((v4 + 0)) | 1))
                                                            var v93 float64 = float64(0.1)
                                                            _ = v93
                                                        }
                                                        _ = fi84
                                                        _ = v90
                                                    }
                                                    v94 := GenS2{F0: (v31 << ((732) & 63))}
                                                    _ = v94
                                                }
                                                _ = fi83
                                            }
                                            for fi95 := 0; fi95 < 11; fi95++ {
                                                _ = len(v3)
                                                for fi96 := 0; fi96 < 13; fi96++ {
                                                    v1 -= v1
                                                    _ = fi96
                                                }
                                                _ = fi95
                                            }
                                            _ = v82
                                        }
                                        if true {
                                            v97 := GenS11{F0: 127, F1: 0, F2: 42}
                                            _ = v97
                                        }
                                    }
                                    fmt.Println(v4, v31, v31)
                                    _ = v81
                                }
                            }
                            v98 := GenS14{F0: 128}
                            _ = v98
                        } else {
                            v1 = v1
                        }
                    }
                    for fi99 := 0; fi99 < 15; fi99++ {
                        if false {
                            fmt.Println(v1)
                            for fi100 := 0; fi100 < 9; fi100++ {
                                v101 := GenS15{F0: v4, F1: 2}
                                var v102 uint16 = 65535
                                var v103 int16 = int16(v33)
                                _ = fi100
                                _ = v101
                                _ = v102
                                _ = v103
                            }
                            for fi104 := 0; fi104 < 1; fi104++ {
                                switch ("" + "") {
                                case "a":
                                    fmt.Println(v1, v33)
                                }
                                v31 = v31
                                var v105 uint8 = 190
                                _ = fi104
                                _ = v105
                            }
                        } else {
                            v106 := GenS16{F0: ("hello" + "a"), F1: 300}
                            v1 = (v1 * (v1 & -738))
                            var v107 bool = false
                            _ = v106
                            _ = v107
                        }
                        _ = fi99
                    }
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = -2147483648
                }
                v33 += 2147483647
                if (true || ((v4 | (v4 * (v4 - v4))) < (^(v4)))) {
                    var v108 int16 = (v1 * 0)
                    var v109 uint8 = 200
                    v110 := GenS17{F0: (-(v108)), F1: v4, F2: true}
                    _ = v108
                    _ = v109
                    _ = v110
                } else {
                    delete(v2, v33)
                }
            }
            _ = v32
            _ = v33
        } else {
            v3 = append(v3, 9223372036854775807)
            if (115 != 28) {
                for fi111 := 0; fi111 < 5; fi111++ {
                    var v112 int16 = int16(v31)
                    _ = fi111
                    _ = v112
                }
            }
            v4 -= v4
        }
    } else {
        var v113 int16 = v1
        var v114 int16 = 32767
        v114 = (v1 + v1)
        _ = v113
        _ = v114
    }
    v115 := GenS18{F0: 18446744073709551615}
    for fi116 := 0; fi116 < 9; fi116++ {
        v117 := v2[(v31 & v31)]
        _ = fi116
        _ = v117
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v118 int32 = 2147483647
        v3 = append(v3, -1)
        _ = v118
    } else {
        v31 -= v31
        v4 -= 0
        switch ("a" + "foo bar") {
        case "foo bar":
            v119 := GenS8{F0: -1, F1: 127, F2: 1}
            _ = v119
        }
    }
    var v120 int32 = 228
    switch 128 {
    case 32767:
        switch 42 {
        case 127:
            var v121 uint = 2
            _ = v121
        case 9223372036854775807:
            v122 := GenS15{F0: (^(v4)), F1: 32767}
            _ = v122
        default:
            v4 = (^(v4))
        }
    case 32768:
        var v123 uint8 = 2
        _ = v123
    case 42:
        v1 += (v1 / ((v1) | 1))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi124 := 0; fi124 < 15; fi124++ {
        fmt.Println(fi124)
        _ = fi124
    }
    v1 += (v1 ^ v1)
    v120 = v120
    if (!(222 >= 128)) {
        v125 := GenS9{F0: (^(v4))}
        v2[v31] = v31
        _ = v125
    }
    v31 += (v31 / (((v31 + v31)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v115
    _ = v120
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
