package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint32; F1 int32; F2 uint }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint16; F1 int64; F2 string }
type GenS7 struct { F0 uint64; F1 int16; F2 bool }
type GenS8 struct { F0 string; F1 string; F2 uint64 }
type GenS9 struct { F0 int16; F1 uint; F2 uint32 }
type GenS10 struct { F0 int8; F1 uint64; F2 int8 }
type GenS11 struct { F0 int32; F1 uint }
type GenS12 struct { F0 uint8; F1 int16; F2 bool }
type GenS13 struct { F0 int64; F1 string; F2 int32 }
func genFn0(p0 uint8) bool {
    v6 := []int{2147483647, 9223372036854775807, 4294967295, -618}
    for fi7 := 0; fi7 < 9; fi7++ {
        if ((fi7 % ((478) | 1)) >= 457) {
            v8 := GenS1{F0: (fi7 & (fi7 / (((fi7 << ((4294967295) & 63))) | 1))), F1: true}
            _ = v8
        } else {
            var v9 int32 = -1
            _ = v9
        }
        switch (fi7 << ((811) & 63)) {
        case 1:
            p0 -= (p0 ^ p0)
        case 4294967295:
            var v10 int16 = -1
            _ = v10
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return false
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{32768, 127, 1}
    var v3 string = "a"
    var v4 uint8 = 255
    v5 := make(map[uint32]uint32)
    v4 -= (v4 | 0)
    for fi11 := 0; fi11 < 1; fi11++ {
        if false {
            fmt.Println(v4, v4)
            switch (fi11 * (fi11 << ((4294967295) & 63))) {
            case 1:
                fmt.Println(v1, fi11, fi11)
            case 32768:
                var v12 uint16 = 82
                _ = v12
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            fmt.Println(v4, v4, v3)
        } else {
            if (!false) {
                var v13 int64 = (v1 + (v1 << ((231) & 63)))
                switch (fi11 - 42) {
                case -1:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        _ = len(v5)
                        switch "" {
                        case "a":
                            _ = len(v2)
                        case "":
                            v4 = v4
                        default:
                            var v15 float64 = (((0.0 - 3.14159) + 1.5) / (1.5 - ((-2.25 + 0.1) / 3.14159)))
                            _ = v15
                        }
                        _ = fi14
                    }
                default:
                    switch fi11 {
                    case 4294967295:
                        for fi16 := 0; fi16 < 8; fi16++ {
                            v17 := GenS2{F0: 2147483647, F1: -2147483648, F2: 101}
                            v4 += v4
                            if (!(false || (((1 <= -685) || (42 != 1)) || false))) {
                                fmt.Println(v3, v1, v13)
                            }
                            _ = fi16
                            _ = v17
                        }
                    case 2147483647:
                        for fi18 := 0; fi18 < 14; fi18++ {
                            switch v3 {
                            case "foo bar":
                                if false {
                                    v19 := GenS3{F0: v4, F1: 42, F2: 4294967295}
                                    _ = v19
                                }
                            case "a":
                                v3 = v3
                            default:
                                v20 := GenS1{F0: (fi11 ^ -1), F1: (false || ("foo bar" != "\n\t"))}
                                _ = v20
                            }
                            var v21 string = v3
                            _ = fi18
                            _ = v21
                        }
                    case 384:
                        for fi22 := 0; fi22 < 10; fi22++ {
                            var v23 uint16 = 255
                            if true {
                                if (42 > 32767) {
                                    var v24 float32 = 0.0
                                    v23 += 325
                                    _ = v24
                                } else {
                                    if false {
                                        v25 := GenS4{F0: 398, F1: (v3 + "\n\t")}
                                        _ = v25
                                    } else {
                                        var v26 float64 = -2.25
                                        _ = v26
                                    }
                                    v5[42] = 42
                                    var v27 bool = (!true)
                                    _ = v27
                                }
                                var v28 int = fi11
                                v29 := GenS3{F0: (v4 >> ((668) & 63)), F1: 1, F2: 1}
                                var v30 int8 = -128
                                if false {
                                    var v31 int8 = 127
                                    var v32 string = v3
                                    _ = v31
                                    _ = v32
                                }
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                _ = len(v2)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            v1 = int64(fi22)
                            _ = fi22
                            _ = v23
                        }
                    }
                }
                v33 := GenS5{F0: (-(v13)), F1: -127}
                _ = v13
                _ = v33
            }
        }
        _ = fi11
    }
    var v34 bool = (false && true)
    var v35 uint = 131
    v35 = v35
    v36 := GenS5{F0: v1, F1: -128}
    v37 := GenS1{F0: 1, F1: v34}
    var v38 bool = (!((-408 == -858) && v34))
    if false {
        v39 := GenS5{F0: v1, F1: 72}
        fmt.Println(v4, v35)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v39
    } else {
        switch v3 {
        case "":
            v3 = v3
        default:
            var v40 uint16 = uint16(v35)
            _ = v40
        }
    }
    fmt.Println(v38, v1)
    v4 = v4
    if false {
        if len(v2) > 0 {
            v2[0] = -770
        }
        switch (v3 + (("foo bar" + v3) + ("hello" + v3))) {
        case "\n\t":
            var v41 uint64 = 0
            _ = v41
        case "foo bar":
            fmt.Println(v38)
        }
    }
    if (255 <= 65535) {
        v38 = (v34 || (100 < 100))
        var v42 string = (v3 + ((v3 + v3) + v3))
        var v43 uint16 = 2
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        for fi45 := 0; fi45 < 17; fi45++ {
            switch fi44 {
            case 42:
                if len(v2) > 0 {
                    v2[0] = fi45
                }
            case -1:
                v5[65535] = 662
            case -970:
                fmt.Println(fi44, fi45, fi45)
            default:
                var v46 float64 = 0.1
                fmt.Println(v35, v46)
                var v47 bool = ((1 < 790) && v38)
                _ = v46
                _ = v47
            }
            fmt.Println(v3, v38)
            _ = fi45
        }
        _ = fi44
    }
    var v48 int8 = -128
    for fi49 := 0; fi49 < 4; fi49++ {
        switch (fi49 * fi49) {
        case -9223372036854775808:
            v3 = v3
        case 1:
            var v50 int8 = (^(v48))
            _ = v50
        case 127:
            var v51 string = (v3 + v3)
            _ = v51
        }
        if (false && false) {
            if (((0.0 - 0.1) + ((float64(0.0) - 1.5) / (((1e10 - (1.5 - 0.0)) * -2.25) + 3.14159))) < (float64(1.5) * 0.1)) {
                var v52 int16 = 32767
                _ = v52
            }
        }
        if (1 != -32768) {
            if (v38 && v34) {
                v48 += (v48 >> ((v35) & 63))
                var v53 uint32 = 0
                var v54 bool = (-877 < 511)
                _ = v53
                _ = v54
            }
            v35 = (v35 & 42)
        } else {
            var v55 uint16 = uint16(v35)
            _ = v55
        }
        _ = fi49
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        for fi57 := 0; fi57 < 13; fi57++ {
            if v34 {
                v1 -= v1
                var v58 uint = (^(v35))
                v59 := GenS4{F0: 4294967295, F1: "foo bar"}
                _ = v58
                _ = v59
            } else {
                var v60 int8 = v48
                fmt.Println(fi56, fi57, v48)
                _ = v60
            }
            if (fi57 != fi56) {
                v2 = append(v2, (fi57 >> (((v35 & (v35 << (((v35 % (((v35 % (((v35 % ((4294967295) | 1))) | 1))) | 1))) & 63)))) & 63)))
                if v34 {
                    var v61 int64 = (v1 / ((9223372036854775807) | 1))
                    _ = v61
                }
            } else {
                v62 := GenS6{F0: 2, F1: v1, F2: (v3 + ((((v3 + v3) + v3) + ((v3 + (v3 + "")) + ((v3 + "") + "foo bar"))) + "\n\t"))}
                if (v38 && (!true)) {
                    v63 := GenS4{F0: 2, F1: v3}
                    _ = v63
                } else {
                    fmt.Println(fi56, v48)
                    switch (^(fi57)) {
                    case -2147483648:
                        if false {
                            switch ("foo bar" + "a") {
                            case "hello":
                                v2 = append(v2, (-(fi57)))
                            case "\n\t":
                                if (42 == -2) {
                                    v38 = (((1 <= 255) || (float64(1e10) < 0.0)) || (0 < -998))
                                    v1 = -9223372036854775808
                                } else {
                                    switch fi57 {
                                    case 127:
                                        v48 = v48
                                    }
                                }
                            case "a":
                                var v64 int16 = 148
                                _ = v64
                            }
                            v35 = v35
                            v35 -= v35
                        } else {
                            switch -2147483648 {
                            case 1:
                                v5[2] = uint32(v35)
                            case 279:
                                if (862 == 2147483647) {
                                    if (42 == -1) {
                                        if true {
                                            if v38 {
                                                var v65 int32 = -2147483648
                                                fmt.Println(v65, v35)
                                                v48 = 42
                                                _ = v65
                                            }
                                            var v66 int16 = -824
                                            if true {
                                                v48 = 127
                                                v35 += 0
                                                var v67 string = "hello"
                                                _ = v67
                                            } else {
                                                v68 := GenS7{F0: 736, F1: int16(v1), F2: true}
                                                _ = v68
                                            }
                                            v38 = v34
                                            var v69 uint32 = 368
                                            _ = v66
                                            _ = v69
                                        }
                                        var v70 uint8 = 108
                                        _ = v70
                                    } else {
                                        switch (fi56 >> (((^(v35))) & 63)) {
                                        case 1:
                                            var v71 float64 = 0.1
                                            _ = v71
                                        default:
                                            v72 := v5[363]
                                            _ = v72
                                        }
                                        v73 := GenS2{F0: 854, F1: 1, F2: uint(v4)}
                                        var v74 uint64 = 4294967295
                                        switch fi57 {
                                        case -909:
                                            _ = len(v2)
                                        case 9223372036854775807:
                                            var v75 int8 = 1
                                            _ = v75
                                        }
                                        _ = v73
                                        _ = v74
                                    }
                                }
                            case 30:
                                v35 -= v35
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v48 = 0
                                delete(v5, 502)
                                if ((v4 != v4) || (!(!v38))) {
                                    v76 := GenS8{F0: "hello", F1: "\n\t", F2: 4294967295}
                                    _ = v76
                                } else {
                                    if v34 {
                                        v34 = true
                                        if v34 {
                                            switch (fi57 - fi57) {
                                            case -384:
                                                var v77 uint32 = 65535
                                                _ = v77
                                            default:
                                                v78 := GenS7{F0: 830, F1: -32768, F2: v34}
                                                _ = v78
                                            }
                                            if (-2.25 == 0.0) {
                                                switch v3 {
                                                case "foo bar":
                                                    v3 = v3
                                                }
                                            } else {
                                                v38 = false
                                                _ = len(v5)
                                                var v79 uint = (v35 - (^(v35)))
                                                v35 -= v35
                                                v48 += (v48 + int8(v4))
                                                if (!v38) {
                                                    v2 = append(v2, fi57)
                                                    v1 = (v1 & (v1 / ((v1) | 1)))
                                                    if (v4 != 127) {
                                                        if false {
                                                            var v80 uint = v79
                                                            v81 := GenS5{F0: (v1 - v1), F1: v48}
                                                            v79 += (^(v79))
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        v82 := GenS9{F0: -32768, F1: v79, F2: 2}
                                                        _ = v82
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v1 += (^(v1))
                                                    }
                                                    var v83 int64 = (v1 | (v1 & v1))
                                                    var v84 int8 = (v48 & -21)
                                                    v85 := GenS2{F0: 517, F1: 1, F2: v79}
                                                    v86 := GenS3{F0: (v4 + v4), F1: 42, F2: 2}
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    var v87 uint8 = (v4 + (^(v4)))
                                                    v3 = "hello"
                                                    var v88 uint16 = 331
                                                    v89 := GenS10{F0: int8(v88), F1: 329, F2: v48}
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                }
                                                _ = len(v2)
                                                switch fi57 {
                                                case 42:
                                                    var v90 bool = ((v1 >> ((v35) & 63)) <= (v1 << (((v35 << ((v35) & 63))) & 63)))
                                                    _ = v90
                                                case -265:
                                                    var v91 float32 = 3.14159
                                                    _ = v91
                                                case 2147483647:
                                                    v92 := GenS11{F0: 2147483647, F1: (^(v79))}
                                                    _ = v92
                                                }
                                                var v93 int = (fi57 / ((fi57) | 1))
                                                _ = v79
                                                _ = v93
                                            }
                                            v34 = v34
                                        } else {
                                            v94 := v5[1]
                                            v95 := v5[(^(v94))]
                                            _ = v94
                                            _ = v95
                                        }
                                    } else {
                                        _ = len(v5)
                                        if (false || (false || false)) {
                                            fmt.Println(v35, v1)
                                            v4 = (v4 << (((^(v35))) & 63))
                                            fmt.Println(v35, fi57)
                                        }
                                    }
                                    v96 := GenS2{F0: 2147483648, F1: -2147483648, F2: v35}
                                    _ = v96
                                }
                            }
                            switch v3 {
                            case "foo bar":
                                v2 = append(v2, fi56)
                            case "\n\t":
                                fmt.Println(v1, v35)
                            }
                            var v97 uint32 = 970
                            _ = v97
                        }
                    case 4294967295:
                        v38 = true
                        if ((v4 << ((v35) & 63)) < 136) {
                            fmt.Println(v1)
                            fmt.Println(v38, v48)
                        }
                    case -9223372036854775808:
                        switch v3 {
                        case "":
                            fmt.Println(fi56, v34, v4)
                        case "foo bar":
                            v98 := GenS5{F0: v1, F1: int8(v4)}
                            var v99 string = "a"
                            switch fi57 {
                            case 133:
                                fmt.Println(v1, v35)
                            case 42:
                                delete(v5, 2147483647)
                            }
                            _ = v98
                            _ = v99
                        case "a":
                            v100 := GenS7{F0: 1, F1: 32767, F2: (!true)}
                            _ = v100
                        }
                    }
                    var v101 int = 157
                    _ = v101
                }
                v2 = append(v2, (fi57 + (fi56 - 127)))
                _ = v62
            }
            _ = fi57
        }
        for fi102 := 0; fi102 < 15; fi102++ {
            if false {
                if (!(!v34)) {
                    var v103 uint8 = v4
                    if ((v48 / ((v48) | 1)) > v48) {
                        var v104 int16 = 285
                        fmt.Println(v35)
                        v1 = (v1 % ((-1) | 1))
                        if false {
                            v38 = (!((v103 / ((v103) | 1)) <= (^(v103))))
                            var v105 uint16 = 2
                            var v106 uint16 = v105
                            _ = v105
                            _ = v106
                        } else {
                            v107 := GenS9{F0: v104, F1: 255, F2: 42}
                            _ = len(v2)
                            _ = v107
                        }
                        var v108 uint8 = (v103 * (v103 ^ 1))
                        _ = v104
                        _ = v108
                    } else {
                        if false {
                            _ = len(v2)
                        } else {
                            var v109 int8 = v48
                            switch v3 {
                            case "hello":
                                switch v3 {
                                case "foo bar":
                                    var v110 int16 = 1000
                                    _ = v110
                                default:
                                    v111 := GenS12{F0: (v103 | v4), F1: -32768, F2: v34}
                                    _ = v111
                                }
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = v109
                        }
                        var v112 int16 = 1000
                        _ = v112
                    }
                    _ = v103
                }
                switch (((v3 + "hello") + v3) + (v3 + (("foo bar" + v3) + (v3 + v3)))) {
                case "hello":
                    var v113 uint8 = (v4 - (v4 * 250))
                    _ = v113
                case "a":
                    v3 = v3
                }
            }
            _ = fi102
        }
        for fi114 := 0; fi114 < 7; fi114++ {
            if len(v2) > 0 {
                v2[0] = fi56
            }
            if ((-2.25 - 1e10) != 1.5) {
                fmt.Println(v35, fi56, fi56)
            } else {
                if (1 < 42) {
                    var v115 uint16 = 255
                    v116 := GenS13{F0: int64(v4), F1: "", F2: -353}
                    if len(v2) > 0 {
                        v2[0] = fi114
                    }
                    v117 := GenS5{F0: (v1 & (v1 % (((v1 | 1)) | 1))), F1: v48}
                    if (0 == -580) {
                        var v118 float32 = float32(0.0)
                        _ = v118
                    }
                    v1 += v1
                    v5[1] = 894
                    _ = v115
                    _ = v116
                    _ = v117
                }
                v2 = append(v2, (-(fi114)))
            }
            _ = fi114
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v48
    _ = v5
}
