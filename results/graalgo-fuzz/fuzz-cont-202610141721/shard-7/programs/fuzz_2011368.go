package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint32; F2 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 bool; F2 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 int32; F2 uint }
type GenS12 struct { F0 int64; F1 int32 }
type GenS13 struct { F0 int16; F1 uint64 }
type GenS14 struct { F0 uint32; F1 uint32; F2 bool }
type GenS15 struct { F0 uint32; F1 int; F2 int }
type GenS16 struct { F0 int8; F1 uint16; F2 bool }
type GenS17 struct { F0 string; F1 uint64; F2 uint32 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 string }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 uint16; F1 uint64 }
type GenS22 struct { F0 string; F1 int; F2 uint8 }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]int)
    v3 := []int{32768, -2147483648, -2147483648, -2147483648}
    var v4 string = "hello"
    var v5 int = 32767
    var v6 uint32 = 2
    if (42 == 472) {
        if (0 == int(v1)) {
            if len(v3) > 0 {
                v3[0] = (v5 >> ((612) & 63))
            }
        } else {
            var v7 uint16 = 1
            _ = v7
        }
    } else {
        _ = len(v3)
        switch "a" {
        case "foo bar":
            if false {
                v8 := GenS1{F0: (68 < 42)}
                var v9 int = (v5 | v5)
                _ = v8
                _ = v9
            } else {
                var v10 uint = 851
                _ = v10
            }
        case "\n\t":
            v2[v5] = 8
        }
        v1 = (v1 | 1)
    }
    v4 = (v4 + v4)
    if true {
        v5 -= 128
        var v11 uint = 255
        v5 = (v5 >> (((v11 - v11)) & 63))
        _ = v11
    } else {
        delete(v2, (^(v5)))
        v1 += (v1 >> ((4294967295) & 63))
    }
    if true {
        fmt.Println(v6)
        v12 := v2[v5]
        for fi13 := 0; fi13 < 6; fi13++ {
            v6 -= 2147483647
            fmt.Println(v12)
            _ = fi13
        }
        _ = v12
    } else {
        v6 = v6
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS1{F0: true}
            for fi16 := 0; fi16 < 18; fi16++ {
                fmt.Println(v5, v6, v1)
                if ((!false) && (!((3.14159 - float32(1.5)) >= (3.14159 * float32(0.1))))) {
                    delete(v2, v5)
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
    }
    v5 -= v5
    var v17 float64 = 0.0
    for fi18 := 0; fi18 < 14; fi18++ {
        v1 -= -17
        if true {
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS2{F0: fi18}
                _ = fi19
                _ = v20
            }
        } else {
            v5 -= (fi18 * v5)
            v6 = v6
        }
        _ = fi18
    }
    v1 = v1
    fmt.Println(v5, v5)
    v21 := GenS3{F0: 18446744073709551615, F1: int(v1)}
    var v22 bool = (true && (v1 > (v1 * v1)))
    if (v22 || (false || (1 != 255))) {
        var v23 uint64 = uint64(v6)
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                if v22 {
                    var v26 uint16 = 122
                    if v22 {
                        fmt.Println(v22, v5)
                    }
                    _ = v26
                } else {
                    v27 := GenS4{F0: (fi25 / ((fi24) | 1)), F1: 255, F2: ((0 <= 0) && (v22 && v22))}
                    if ((1.5 <= 1e10) || (3.14159 >= v17)) {
                        v1 += (v1 * 1)
                        if v22 {
                            if (784 > 335) {
                                switch (v5 << ((18446744073709551615) & 63)) {
                                case 599:
                                    v28 := GenS4{F0: (v5 | int(v1)), F1: 639, F2: v22}
                                    v23 -= (v23 * 42)
                                    _ = v28
                                case 0:
                                    var v29 uint64 = (v23 + (^(v23)))
                                    _ = v29
                                case 32767:
                                    if ((0 > 42) || (!true)) {
                                        v17 += 3.14159
                                        var v30 int8 = v1
                                        _ = v30
                                    }
                                }
                                var v31 int8 = -63
                                _ = v31
                            } else {
                                var v32 bool = (v22 || (((3.14159 + 1e10) / (-2.25 / (0.0 - (1e10 / float32(0.1))))) >= 1e10))
                                v33 := GenS5{F0: 127}
                                if (fi25 < int(v6)) {
                                    v23 += v23
                                } else {
                                    v6 = (v6 >> ((1) & 63))
                                    if len(v3) > 0 {
                                        v3[0] = 32768
                                    }
                                }
                                _ = v32
                                _ = v33
                            }
                            var v34 uint64 = (v23 / ((18446744073709551615) | 1))
                            _ = v34
                        }
                    } else {
                        if (!(fi24 > -1)) {
                            var v35 float64 = (v17 / 1.0)
                            v23 -= v23
                            v36 := GenS2{F0: int(v6)}
                            _ = v35
                            _ = v36
                        }
                        if (v22 || (1.5 >= 3.14159)) {
                            var v37 int32 = 467
                            _ = v37
                        } else {
                            var v38 int64 = int64(v1)
                            v17 = 3.14159
                            if len(v3) > 0 {
                                v3[0] = (fi25 >> ((4294967295) & 63))
                            }
                            _ = v38
                        }
                        var v39 int = fi24
                        _ = v39
                    }
                    _ = v27
                }
                fmt.Println(v22)
                var v40 string = ("\n\t" + v4)
                if v22 {
                    switch "foo bar" {
                    case "":
                        fmt.Println(v23, v6, v5)
                    case "foo bar":
                        switch (fi25 & int(v6)) {
                        case -1:
                            v41 := GenS3{F0: (v23 - (v23 << ((65535) & 63))), F1: int(v6)}
                            _ = v41
                        case -562:
                            var v42 int8 = (v1 | -128)
                            _ = v42
                        }
                    default:
                        v43 := GenS4{F0: fi25, F1: 2, F2: v22}
                        _ = v43
                    }
                    v44 := GenS6{F0: (v1 >> ((18446744073709551615) & 63)), F1: (v6 & v6), F2: v5}
                    v22 = (!(9223372036854775807 == 1))
                    _ = v44
                } else {
                    if false {
                        if false {
                            v22 = v22
                        }
                        var v45 string = (v40 + "foo bar")
                        if v22 {
                            v4 = "\n\t"
                        }
                        var v46 int = (fi24 * fi24)
                        var v47 int16 = -32768
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    delete(v2, (fi25 | (fi24 - (v5 + 2147483647))))
                }
                _ = fi25
                _ = v40
            }
            for fi48 := 0; fi48 < 18; fi48++ {
                v22 = (-1 <= 0)
                if (((int32(v6) <= 42) && v22) && (0.1 < v17)) {
                    if true {
                        var v49 uint64 = (^(v23))
                        if true {
                            v23 = v49
                            if (!v22) {
                                _ = len(v2)
                                var v50 uint16 = 0
                                _ = v50
                            } else {
                                if (v22 || (((fi24 >> ((4294967295) & 63)) > fi24) || (2147483647 < 471))) {
                                    delete(v2, v5)
                                    v2[-9223372036854775808] = -798
                                    var v51 bool = (((fi24 << ((194) & 63)) < (fi24 + (-(fi48)))) && true)
                                    _ = v51
                                }
                                switch v4 {
                                case "":
                                    if true {
                                        v22 = (4294967295 < v5)
                                        v22 = false
                                        v52 := GenS1{F0: (!(fi48 < fi48))}
                                        _ = v52
                                    }
                                case "foo bar":
                                    if (!false) {
                                        if len(v3) > 0 {
                                            v3[0] = v5
                                        }
                                        _ = len(v3)
                                        _ = len(v3)
                                        var v53 string = v4
                                        v1 -= (v1 / (((v1 - (v1 << ((uint(v23)) & 63)))) | 1))
                                        _ = v53
                                    } else {
                                        if v22 {
                                            v4 = v4
                                        } else {
                                            switch (^(v5)) {
                                            case -2147483648:
                                                var v54 bool = true
                                                _ = v54
                                            case 972:
                                                v55 := v2[(fi48 ^ (v5 ^ 32768))]
                                                _ = v55
                                            }
                                            var v56 bool = false
                                            _ = v56
                                        }
                                    }
                                case "\n\t":
                                    if false {
                                        var v57 float32 = 0.0
                                        v1 += v1
                                        var v58 int64 = 102
                                        _ = v57
                                        _ = v58
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        } else {
                            delete(v2, 1)
                            v59 := GenS5{F0: 127}
                            _ = v59
                        }
                        if true {
                            v6 += 165
                            if (!v22) {
                                var v60 uint8 = 1
                                _ = v60
                            }
                            v23 = (v49 % ((306) | 1))
                        }
                        var v61 string = v4
                        if v22 {
                            if false {
                                switch ((("a" + v4) + v4) + v4) {
                                case "hello":
                                    var v62 int16 = 0
                                    _ = v62
                                case "":
                                    v1 = (-(v1))
                                }
                                _ = len(v3)
                                if true {
                                    v6 = 2147483647
                                    fmt.Println(v61, v23, fi24)
                                } else {
                                    fmt.Println(fi48, v6, v61)
                                    v5 = 32767
                                    if false {
                                        v1 += (v1 ^ (v1 << ((160) & 63)))
                                    }
                                }
                            } else {
                                v5 += fi48
                                v63 := GenS3{F0: 512, F1: (fi48 + (fi24 % (((fi48 * fi24)) | 1)))}
                                var v64 string = ("hello" + ((("hello" + v4) + v61) + (((v61 + v4) + "hello") + (v61 + "hello"))))
                                _ = v63
                                _ = v64
                            }
                            fmt.Println(v5, fi24)
                            v65 := GenS7{F0: 42}
                            _ = v65
                        } else {
                            if false {
                                v17 += (v17 / 1.0)
                            }
                            if ((v61 + v61) < ("a" + (v4 + v61))) {
                                _ = len(v3)
                                _ = len(v2)
                                var v66 int16 = 665
                                _ = v66
                            } else {
                                v49 = (^(v23))
                            }
                            var v67 int16 = 523
                            v5 = 4294967295
                            var v68 float32 = ((float32(v17) + float32(v17)) / (1.5 + -2.25))
                            _ = v67
                            _ = v68
                        }
                        _ = v49
                        _ = v61
                    } else {
                        var v69 bool = v22
                        fmt.Println(v1, v6)
                        _ = v69
                    }
                    if ((v6 + v6) == (v6 - (v6 / ((v6) | 1)))) {
                        if len(v3) > 0 {
                            v3[0] = (-(v5))
                        }
                        if (2 != (v6 | (v6 * (v6 >> ((255) & 63))))) {
                            v1 = (^(v1))
                        }
                        v1 += (v1 + (v1 | (v1 + (v1 >> ((847) & 63)))))
                        v70 := GenS8{F0: 890, F1: v22, F2: (v23 - (v23 / ((v23) | 1)))}
                        _ = v70
                    } else {
                        if (0 < 630) {
                            v17 = v17
                        } else {
                            switch (v5 / ((510) | 1)) {
                            case -9223372036854775808:
                                v71 := v2[32768]
                                _ = v71
                            case 32767:
                                _ = len(v2)
                            default:
                                var v72 uint8 = 205
                                _ = v72
                            }
                            switch "hello" {
                            case "foo bar":
                                if (v23 != v23) {
                                    switch v4 {
                                    case "a":
                                        switch ("\n\t" + v4) {
                                        case "\n\t":
                                            var v73 int32 = 42
                                            v74 := GenS9{F0: -468}
                                            var v75 uint64 = v23
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        }
                                        if (v4 > "") {
                                            switch "hello" {
                                            case "\n\t":
                                                var v76 float64 = -2.25
                                                _ = v76
                                            case "":
                                                switch v5 {
                                                case 0:
                                                    v22 = (143 >= 42)
                                                case 128:
                                                    v3 = append(v3, v5)
                                                    var v77 int8 = -36
                                                    _ = v77
                                                case 981:
                                                    v23 = (v23 + v23)
                                                }
                                            case "a":
                                                delete(v2, 1)
                                            default:
                                                fmt.Println(v4, v5)
                                            }
                                            v78 := GenS10{F0: 255}
                                            _ = v78
                                        } else {
                                            var v79 uint8 = 128
                                            if (v17 >= (v17 * (v17 * -2.25))) {
                                                if ((!true) && (!(!(9223372036854775807 < -9223372036854775808)))) {
                                                    if v22 {
                                                        v80 := GenS8{F0: -254, F1: (0 > 0), F2: (v23 << ((147) & 63))}
                                                        if (-316 == -1) {
                                                            v2[42] = v5
                                                            fmt.Println(v4, fi48)
                                                            fmt.Println(v4)
                                                        }
                                                        switch (fi48 << ((4294967295) & 63)) {
                                                        case 1:
                                                            var v81 int64 = 42
                                                            _ = v81
                                                        case 2147483647:
                                                            fmt.Println(v5, v79)
                                                        }
                                                        _ = v80
                                                    } else {
                                                        if v22 {
                                                            var v82 uint64 = (^(v23))
                                                            v5 = -126
                                                            switch fi24 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v22)
                                                            case 826:
                                                                var v83 uint64 = (v82 | (v23 % ((v23) | 1)))
                                                                _ = len(v3)
                                                                fmt.Println(v6, v17)
                                                                fmt.Println(fi24, v79, v23)
                                                                _ = v83
                                                            case -879:
                                                                delete(v2, 4294967295)
                                                                v84 := GenS9{F0: int64(v79)}
                                                                if v22 {
                                                                    v4 = ((v4 + "a") + v4)
                                                                } else {
                                                                    switch ((v4 + (v4 + "foo bar")) + (v4 + v4)) {
                                                                    case "foo bar":
                                                                        switch v5 {
                                                                        case -9223372036854775808:
                                                                            var v85 int64 = 42
                                                                            _ = v85
                                                                        case 42:
                                                                            if (65535 > 32767) {
                                                                                var v86 uint32 = (v6 ^ v6)
                                                                                _ = v86
                                                                            } else {
                                                                                if v22 {
                                                                                    _ = len(v2)
                                                                                    if ((((128 < v5) || (!(!v22))) || (100 >= 118)) && true) {
                                                                                        v82 += v82
                                                                                        var v87 int64 = -565
                                                                                        v88 := GenS8{F0: (^(v87)), F1: v22, F2: v23}
                                                                                        _ = len(v3)
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                    } else {
                                                                                        if v22 {
                                                                                            v89 := GenS11{F0: uint32(v23), F1: int32(v1), F2: 2}
                                                                                            v17 += -2.25
                                                                                            if (780 == 807) {
                                                                                                fmt.Println(v6)
                                                                                                _ = len(v3)
                                                                                                v23 = (v23 | (v82 - v82))
                                                                                                if (-961 < 433) {
                                                                                                    v90 := GenS12{F0: -9223372036854775808, F1: -24}
                                                                                                    var v91 int8 = 1
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                } else {
                                                                                                    switch (fi24 & (fi24 & v5)) {
                                                                                                    case 481:
                                                                                                        if (42 < 1) {
                                                                                                            fmt.Println(v17, v5)
                                                                                                            v92 := v2[fi48]
                                                                                                            var v93 int32 = 0
                                                                                                            _ = v92
                                                                                                            _ = v93
                                                                                                        } else {
                                                                                                            var v94 uint = 65535
                                                                                                            _ = v94
                                                                                                        }
                                                                                                    case -640:
                                                                                                        var v95 float64 = (v17 + v17)
                                                                                                        _ = v95
                                                                                                    default:
                                                                                                        fmt.Println(v6)
                                                                                                    }
                                                                                                    var v96 float64 = 0.1
                                                                                                    v97 := GenS13{F0: 1000, F1: (^(v23))}
                                                                                                    _ = v96
                                                                                                    _ = v97
                                                                                                }
                                                                                            } else {
                                                                                                v98 := GenS14{F0: v6, F1: 2147483648, F2: false}
                                                                                                _ = v98
                                                                                            }
                                                                                            if (100 >= v1) {
                                                                                                fmt.Println(v22, v79)
                                                                                                var v99 int16 = 1
                                                                                                _ = v99
                                                                                            } else {
                                                                                                var v100 float32 = 1.5
                                                                                                if (v22 || (v22 && v22)) {
                                                                                                    delete(v2, (v5 ^ fi48))
                                                                                                    v22 = v22
                                                                                                    if true {
                                                                                                        v101 := v2[fi48]
                                                                                                        v102 := GenS15{F0: 251, F1: fi24, F2: 220}
                                                                                                        var v103 int16 = -1
                                                                                                        if false {
                                                                                                            v79 -= 128
                                                                                                        } else {
                                                                                                            v104 := GenS8{F0: -1, F1: v22, F2: uint64(v79)}
                                                                                                            _ = v104
                                                                                                        }
                                                                                                        _ = v101
                                                                                                        _ = v102
                                                                                                        _ = v103
                                                                                                    }
                                                                                                } else {
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    v1 = (v1 ^ (v1 % (((v1 / (((v1 - -1)) | 1))) | 1)))
                                                                                                    v2[-698] = fi24
                                                                                                }
                                                                                                _ = v100
                                                                                            }
                                                                                            v23 = v23
                                                                                            _ = v89
                                                                                        } else {
                                                                                            v105 := v2[fi24]
                                                                                            var v106 int = 32767
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    if false {
                                                                                        var v107 uint8 = (v79 ^ v79)
                                                                                        v3 = append(v3, (-(v5)))
                                                                                        v5 = fi48
                                                                                        _ = v107
                                                                                    } else {
                                                                                        switch fi48 {
                                                                                        case 818:
                                                                                            switch v4 {
                                                                                            case "foo bar":
                                                                                                v82 += (^(v23))
                                                                                            case "a":
                                                                                                var v108 int64 = -1
                                                                                                _ = v108
                                                                                            case "":
                                                                                                fmt.Println(v22)
                                                                                            default:
                                                                                                var v109 int16 = 118
                                                                                                _ = v109
                                                                                            }
                                                                                        case 0:
                                                                                            switch int(v23) {
                                                                                            case -62:
                                                                                                if v22 {
                                                                                                    var v110 uint16 = 2
                                                                                                    v111 := GenS2{F0: v5}
                                                                                                    v5 += v5
                                                                                                    _ = v110
                                                                                                    _ = v111
                                                                                                } else {
                                                                                                    v17 = v17
                                                                                                    var v112 float32 = 0.1
                                                                                                    _ = v112
                                                                                                }
                                                                                            default:
                                                                                                var v113 uint16 = 32767
                                                                                                _ = v113
                                                                                            }
                                                                                        default:
                                                                                            var v114 uint16 = 140
                                                                                            _ = v114
                                                                                        }
                                                                                    }
                                                                                    var v115 uint32 = v6
                                                                                    _ = v115
                                                                                }
                                                                                v22 = v22
                                                                            }
                                                                        default:
                                                                            var v116 int64 = 42
                                                                            _ = v116
                                                                        }
                                                                    case "hello":
                                                                        var v117 float32 = 3.14159
                                                                        _ = v117
                                                                    }
                                                                    v118 := GenS12{F0: 1, F1: 61}
                                                                    var v119 uint = uint(v82)
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                                fmt.Println(v23)
                                                                _ = v84
                                                            default:
                                                                switch "" {
                                                                case "\n\t":
                                                                    v120 := GenS16{F0: int8(v5), F1: 32767, F2: v22}
                                                                    v121 := GenS17{F0: "\n\t", F1: 744, F2: 2147483648}
                                                                    _ = v120
                                                                    _ = v121
                                                                case "foo bar":
                                                                    v122 := v2[(fi48 / (((fi24 * (fi24 >> ((65535) & 63)))) | 1))]
                                                                    _ = v122
                                                                default:
                                                                    if v22 {
                                                                        var v123 float32 = 0.0
                                                                        _ = v123
                                                                    }
                                                                }
                                                            }
                                                            _ = v82
                                                        } else {
                                                            v79 -= (v79 - (^(v79)))
                                                        }
                                                        switch (^(fi48)) {
                                                        case 4294967295:
                                                            var v124 int = (fi24 % ((fi24) | 1))
                                                            _ = v124
                                                        case 32768:
                                                            if false {
                                                                switch (v5 << ((42) & 63)) {
                                                                case 1:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v125 := GenS18{F0: (fi24 + 4294967295)}
                                                                        _ = v125
                                                                    case "a":
                                                                        if (v1 > v1) {
                                                                            v23 = uint64(v79)
                                                                            var v126 float32 = -2.25
                                                                            var v127 int32 = int32(v1)
                                                                            var v128 int16 = 550
                                                                            v129 := GenS4{F0: fi48, F1: 32767, F2: v22}
                                                                            if true {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = int(v23)
                                                                                }
                                                                                v5 -= (fi48 ^ 603)
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = int(v1)
                                                                                }
                                                                            }
                                                                            _ = v126
                                                                            _ = v127
                                                                            _ = v128
                                                                            _ = v129
                                                                        }
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v130 uint32 = v6
                                                                    _ = v130
                                                                default:
                                                                    var v131 uint32 = 0
                                                                    _ = v131
                                                                }
                                                                v6 -= 484
                                                            }
                                                        case -757:
                                                            var v132 string = "a"
                                                            _ = v132
                                                        }
                                                        if (v22 && true) {
                                                            var v133 int8 = 1
                                                            if (!((^(v23)) <= v23)) {
                                                                var v134 string = ""
                                                                var v135 uint64 = 18446744073709551615
                                                                _ = v134
                                                                _ = v135
                                                            } else {
                                                                v6 += (v6 ^ (v6 + v6))
                                                                v6 = 65535
                                                            }
                                                            fmt.Println(v5, v4, v1)
                                                            _ = v133
                                                        } else {
                                                            if (!v22) {
                                                                var v136 bool = true
                                                                var v137 int8 = 0
                                                                switch v4 {
                                                                case "foo bar":
                                                                    if (!v136) {
                                                                        if v22 {
                                                                            fmt.Println(fi48, v1)
                                                                            var v138 uint = 65535
                                                                            _ = v138
                                                                        } else {
                                                                            if (v23 <= 1) {
                                                                                if false {
                                                                                    fmt.Println(v136, v137)
                                                                                    v139 := GenS14{F0: 0, F1: 0, F2: (((v6 - (v6 * 99)) > v6) && (-730 == -502))}
                                                                                    v140 := GenS19{F0: v4}
                                                                                    fmt.Println(v5, v5)
                                                                                    v141 := GenS10{F0: uint16(v23)}
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                } else {
                                                                                    switch v4 {
                                                                                    case "":
                                                                                        var v142 uint64 = v23
                                                                                        _ = len(v3)
                                                                                        switch (^(fi24)) {
                                                                                        case -2147483648:
                                                                                            v5 -= fi24
                                                                                        case 42:
                                                                                            switch (-(fi24)) {
                                                                                            case 638:
                                                                                                if (!v22) {
                                                                                                    var v143 string = "hello"
                                                                                                    var v144 uint16 = 255
                                                                                                    _ = v143
                                                                                                    _ = v144
                                                                                                } else {
                                                                                                    var v145 float64 = -2.25
                                                                                                    if false {
                                                                                                        var v146 string = (((v4 + v4) + (v4 + "a")) + "foo bar")
                                                                                                        _ = v146
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v147 := GenS1{F0: ((v6 / ((v6) | 1)) == 582)}
                                                                                                            v79 = v79
                                                                                                            _ = len(v2)
                                                                                                            _ = v147
                                                                                                        }
                                                                                                        if (!(false || v136)) {
                                                                                                            v148 := GenS9{F0: 9223372036854775807}
                                                                                                            var v149 uint64 = (v23 / ((v142) | 1))
                                                                                                            var v150 string = v4
                                                                                                            v5 = fi48
                                                                                                            _ = v148
                                                                                                            _ = v149
                                                                                                            _ = v150
                                                                                                        }
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        var v151 int32 = -2147483648
                                                                                                        _ = v151
                                                                                                    }
                                                                                                    _ = v145
                                                                                                }
                                                                                                var v152 uint16 = 383
                                                                                                _ = v152
                                                                                            case 1:
                                                                                                var v153 int = 352
                                                                                                _ = v153
                                                                                            }
                                                                                        default:
                                                                                            switch ("\n\t" + v4) {
                                                                                            case "":
                                                                                                switch (v5 >> ((682) & 63)) {
                                                                                                case 127:
                                                                                                    v17 = (v17 - (v17 * v17))
                                                                                                case 32768:
                                                                                                    v5 = v5
                                                                                                    var v154 bool = true
                                                                                                    v155 := GenS14{F0: (v6 % (((v6 ^ (v6 + v6))) | 1)), F1: (v6 * (^(v6))), F2: false}
                                                                                                    _ = v154
                                                                                                    _ = v155
                                                                                                }
                                                                                            default:
                                                                                                v156 := GenS20{F0: -2147483648}
                                                                                                _ = v156
                                                                                            }
                                                                                        }
                                                                                        _ = v142
                                                                                    case "foo bar":
                                                                                        if v136 {
                                                                                            v22 = ("a" > "")
                                                                                            var v157 int16 = 989
                                                                                            if ((!(1 == -1)) || (v157 > v157)) {
                                                                                                var v158 uint32 = (v6 ^ v6)
                                                                                                if v22 {
                                                                                                    v157 += v157
                                                                                                }
                                                                                                fmt.Println(v22, v5)
                                                                                                _ = v158
                                                                                            } else {
                                                                                                v159 := v2[(fi48 / ((-160) | 1))]
                                                                                                _ = v159
                                                                                            }
                                                                                            _ = v157
                                                                                        }
                                                                                    case "a":
                                                                                        if (v136 && true) {
                                                                                            v160 := v2[v5]
                                                                                            fmt.Println(v136, v1, v17)
                                                                                            v79 = v79
                                                                                            var v161 uint = 4294967295
                                                                                            var v162 uint8 = (v79 | (v79 >> (((v161 << ((v161) & 63))) & 63)))
                                                                                            _ = v160
                                                                                            _ = v161
                                                                                            _ = v162
                                                                                        } else {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v163 int64 = 42
                                                                                        _ = v163
                                                                                    }
                                                                                    delete(v2, fi24)
                                                                                    var v164 uint8 = (v79 / ((v79) | 1))
                                                                                    _ = v164
                                                                                }
                                                                                v79 += 170
                                                                                switch v5 {
                                                                                case 4294967295:
                                                                                    fmt.Println(v23, v1, v136)
                                                                                case -2147483648:
                                                                                    var v165 bool = (1.5 <= 1e10)
                                                                                    _ = v165
                                                                                case 2147483647:
                                                                                    var v166 string = (((((((v4 + "foo bar") + "hello") + (v4 + (v4 + ("foo bar" + (v4 + v4))))) + v4) + "") + "foo bar") + v4)
                                                                                    _ = v166
                                                                                default:
                                                                                    var v167 uint32 = (v6 & v6)
                                                                                    _ = v167
                                                                                }
                                                                                var v168 float64 = float64(0.0)
                                                                                _ = v168
                                                                            }
                                                                        }
                                                                        var v169 int = 1
                                                                        _ = v169
                                                                    } else {
                                                                        fmt.Println(v79, v137, v5)
                                                                    }
                                                                case "":
                                                                    var v170 float64 = 0.1
                                                                    _ = v170
                                                                case "\n\t":
                                                                    v171 := GenS15{F0: 212, F1: fi24, F2: v5}
                                                                    _ = v171
                                                                }
                                                                var v172 uint64 = uint64(v6)
                                                                var v173 int = 982
                                                                _ = v136
                                                                _ = v137
                                                                _ = v172
                                                                _ = v173
                                                            }
                                                            var v174 uint64 = (v23 % (((v23 & 18446744073709551615)) | 1))
                                                            _ = v174
                                                        }
                                                        switch 32767 {
                                                        case 2147483647:
                                                            if v22 {
                                                                var v175 float32 = 0.1
                                                                _ = v175
                                                            } else {
                                                                v17 = v17
                                                                var v176 int8 = 1
                                                                _ = v176
                                                            }
                                                        case 1:
                                                            v177 := GenS21{F0: 164, F1: v23}
                                                            _ = v177
                                                        case -9223372036854775808:
                                                            v23 = v23
                                                        }
                                                        v17 -= v17
                                                    }
                                                    fmt.Println(fi24)
                                                    var v178 int8 = v1
                                                    _ = v178
                                                }
                                                var v179 bool = (-9223372036854775808 > 42)
                                                _ = v179
                                            }
                                            _ = v79
                                        }
                                        var v180 int8 = 42
                                        _ = v180
                                    case "":
                                        var v181 uint16 = 42
                                        _ = v181
                                    case "\n\t":
                                        var v182 int = int(v23)
                                        _ = v182
                                    }
                                    var v183 float32 = ((1.5 - (0.1 * float32((v17 + 0.0)))) / 1.5)
                                    v23 = v23
                                    _ = v183
                                } else {
                                    v184 := GenS6{F0: (v1 % (((v1 ^ (v1 * int8(v6)))) | 1)), F1: v6, F2: (-(fi48))}
                                    v185 := GenS11{F0: 133, F1: -2147483648, F2: 42}
                                    _ = v184
                                    _ = v185
                                }
                            case "\n\t":
                                v6 -= v6
                            default:
                                v4 = (v4 + v4)
                            }
                            var v186 int16 = 1
                            _ = v186
                        }
                        var v187 int64 = 1
                        _ = v187
                    }
                    v22 = v22
                } else {
                    if v22 {
                        var v188 uint64 = v23
                        var v189 int32 = 32
                        if v22 {
                            fmt.Println(v4, v23)
                            var v190 uint8 = 47
                            _ = v190
                        }
                        _ = v188
                        _ = v189
                    } else {
                        v191 := v2[fi48]
                        _ = v191
                    }
                }
                _ = fi48
            }
            _ = fi24
        }
        _ = v23
    } else {
        for fi192 := 0; fi192 < 11; fi192++ {
            var v193 int8 = (v1 / (((^(v1))) | 1))
            _ = fi192
            _ = v193
        }
        v194 := GenS22{F0: v4, F1: v5, F2: 87}
        _ = v194
    }
    v17 -= v17
    delete(v2, v5)
    v195 := v2[(v5 << ((0) & 63))]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v195:", v195)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v195
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
