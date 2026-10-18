package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int8; F1 int64; F2 int64 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int64; F1 int64; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int; F1 int64; F2 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int; F1 uint16; F2 uint64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint64; F1 int32; F2 int8 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int16; F1 int64 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int16; F1 bool; F2 string }

func main() {
    var v1 bool = (true && true)
    var v2 int8 = -128
    var v3 string = ("foo bar" + ("a" + "foo bar"))
    var v4 uint32 = 2
    v5 := []string{(("" + v3) + "foo bar"), v3, (v3 + ((v3 + ("hello" + v3)) + (((v3 + "foo bar") + "") + "a")))}
    v6 := []int{32768, 2147483647}
    for fi7 := 0; fi7 < 2; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if v1 {
            if false {
                for fi8 := 0; fi8 < 17; fi8++ {
                    v9 := GenS1{F0: 42, F1: 53}
                    fmt.Println(fi7, fi8)
                    _ = fi8
                    _ = v9
                }
            } else {
                if (!v1) {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        var v11 uint16 = 872
                        _ = len(v5)
                        _ = fi10
                        _ = v11
                    }
                }
            }
            v12 := GenS2{F0: 1, F1: 42}
            v1 = (679 >= 32767)
            _ = v12
        } else {
            v13 := GenS2{F0: 42, F1: -640}
            var v14 float32 = (0.1 * 3.14159)
            _ = v13
            _ = v14
        }
        _ = fi7
    }
    switch -743 {
    case -553:
        for fi15 := 0; fi15 < 18; fi15++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi16 := 0; fi16 < 4; fi16++ {
                switch v3 {
                case "\n\t":
                    switch (fi16 >> ((826) & 63)) {
                    case 32768:
                        v17 := GenS1{F0: v4, F1: 127}
                        _ = v17
                    default:
                        v3 = v3
                        var v18 int8 = v2
                        v19 := GenS3{F0: 2147483647, F1: 403}
                        _ = v18
                        _ = v19
                    }
                case "foo bar":
                    if true {
                        v1 = false
                        var v20 uint64 = 387
                        v21 := GenS1{F0: (v4 >> ((255) & 63)), F1: 2}
                        v3 = v3
                        v22 := GenS4{F0: (v2 >> ((484) & 63)), F1: 2}
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        switch (-(fi15)) {
                        case 1:
                            v23 := GenS5{F0: v2, F1: 0, F2: 0}
                            _ = v23
                        }
                        var v24 int64 = 644
                        if v1 {
                            var v25 int32 = 42
                            v3 = v3
                            var v26 int8 = (v2 / (((v2 >> ((42) & 63))) | 1))
                            _ = v25
                            _ = v26
                        }
                        _ = v24
                    }
                default:
                    var v27 uint = 65535
                    _ = v27
                }
                _ = fi16
            }
            fmt.Println(fi15, v3)
            _ = fi15
        }
    case -2147483648:
        var v28 uint32 = v4
        _ = v28
    case -1:
        _ = len(v6)
    }
    fmt.Println(v3)
    v2 -= 0
    var v29 bool = (0 != -12)
    for fi30 := 0; fi30 < 18; fi30++ {
        for fi31 := 0; fi31 < 7; fi31++ {
            if len(v6) > 0 {
                v6[0] = (fi30 ^ (fi31 / ((fi30) | 1)))
            }
            _ = fi31
        }
        fmt.Println(v3, v4)
        switch (fi30 + fi30) {
        case -796:
            var v32 int = (fi30 - fi30)
            _ = v32
        case 42:
            if len(v5) > 0 {
                v5[0] = ""
            }
        default:
            v33 := GenS3{F0: 0, F1: 1000}
            _ = v33
        }
        _ = fi30
    }
    if (!(!false)) {
        var v34 uint8 = 128
        v1 = v29
        _ = v34
    } else {
        for fi35 := 0; fi35 < 2; fi35++ {
            if true {
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 int16 = 1
                    if false {
                        v38 := GenS4{F0: v2, F1: 328}
                        var v39 uint8 = 255
                        _ = v38
                        _ = v39
                    }
                    _ = fi36
                    _ = v37
                }
                if v29 {
                    v5 = append(v5, ("hello" + v3))
                    var v40 float64 = (-2.25 + 0.1)
                    _ = v40
                } else {
                    v2 += (v2 / (((v2 - v2)) | 1))
                    if true {
                        var v41 int32 = -549
                        if ((v4 % ((v4) | 1)) > v4) {
                            switch (fi35 ^ fi35) {
                            case 592:
                                v41 += (v41 * (v41 / (((v41 << ((634) & 63))) | 1)))
                            }
                            if (!v29) {
                                for fi42 := 0; fi42 < 10; fi42++ {
                                    v4 += (v4 % ((4294967295) | 1))
                                    _ = len(v6)
                                    var v43 float64 = (((3.14159 + 0.1) * 0.0) / 1e10)
                                    var v44 int8 = v2
                                    var v45 uint8 = 255
                                    v46 := GenS6{F0: -1, F1: 0}
                                    fmt.Println(v2)
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            } else {
                                var v47 int = (fi35 << ((300) & 63))
                                _ = v47
                            }
                        }
                        v29 = v29
                        _ = v41
                    } else {
                        if true {
                            for fi48 := 0; fi48 < 11; fi48++ {
                                if v1 {
                                    v2 -= v2
                                    if (((((0.0 + 0.0) - 3.14159) / 1.0) > (1.5 / float64(0.1))) && (!((fi48 - 32768) < (fi35 >> ((42) & 63))))) {
                                        v4 = (^(v4))
                                        v49 := GenS1{F0: 65535, F1: 42}
                                        switch "a" {
                                        case "hello":
                                            if v1 {
                                                switch fi35 {
                                                case -486:
                                                    v3 = v3
                                                default:
                                                    v5 = append(v5, (((v3 + v3) + (((v3 + (("hello" + "a") + "a")) + v3) + v3)) + v3))
                                                }
                                            } else {
                                                v50 := GenS3{F0: 2147483647, F1: -973}
                                                _ = v50
                                            }
                                        case "foo bar":
                                            v29 = v29
                                        default:
                                            v2 = v2
                                        }
                                        v2 = v2
                                        v3 = (((v3 + "foo bar") + v3) + "foo bar")
                                        _ = v49
                                    }
                                    if v1 {
                                        switch (fi35 - (-(fi35))) {
                                        case 329:
                                            if len(v5) > 0 {
                                                v5[0] = "hello"
                                            }
                                        case 0:
                                            if (!(!false)) {
                                                v4 = v4
                                            } else {
                                                if v1 {
                                                    v3 = ("" + v3)
                                                }
                                                var v51 uint32 = (v4 + v4)
                                                v52 := GenS7{F0: 9223372036854775807, F1: -9223372036854775808, F2: 18446744073709551615}
                                                _ = v51
                                                _ = v52
                                            }
                                        }
                                    }
                                } else {
                                    var v53 int = 128
                                    if (!(128 == 150)) {
                                        v53 = 128
                                        v3 = v3
                                        fmt.Println(fi48, v1, v4)
                                        v29 = (!false)
                                        var v54 string = ("" + "a")
                                        _ = v54
                                    }
                                    var v55 uint8 = 255
                                    _ = v53
                                    _ = v55
                                }
                                v1 = ((3.14159 > ((0.1 / 3.14159) + float64(-2.25))) || ((1000 == -852) || (18446744073709551615 > 376)))
                                if (v1 && (v29 && (!v1))) {
                                    v1 = v29
                                }
                                _ = fi48
                            }
                            for fi56 := 0; fi56 < 15; fi56++ {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if (-1 > 809) {
                                    if true {
                                        _ = len(v5)
                                        fmt.Println(fi56, fi56, v1)
                                        if (!v1) {
                                            if v29 {
                                                switch fi56 {
                                                case 861:
                                                    v57 := GenS8{F0: (fi56 * 2147483647)}
                                                    fmt.Println(fi35, fi35, fi56)
                                                    if true {
                                                        v2 += v2
                                                    }
                                                    _ = v57
                                                case 42:
                                                    var v58 uint32 = (v4 | v4)
                                                    _ = v58
                                                default:
                                                    var v59 string = (v3 + "\n\t")
                                                    if v29 {
                                                        switch 2147483647 {
                                                        case 392:
                                                            v3 = v3
                                                            var v60 int16 = -1
                                                            _ = v60
                                                        case -316:
                                                            var v61 string = ("hello" + v59)
                                                            _ = v61
                                                        case 1:
                                                            var v62 int = (fi56 + -266)
                                                            _ = v62
                                                        default:
                                                            fmt.Println(v1, v1)
                                                        }
                                                        var v63 uint16 = 30
                                                        v59 = "foo bar"
                                                        switch ((v59 + ((v3 + "a") + v59)) + (v59 + (v59 + ((v59 + v59) + v3)))) {
                                                        case "a":
                                                            v1 = (v2 > v2)
                                                        case "\n\t":
                                                            if (!(((fi35 ^ fi56) < 42) && (347 <= 0))) {
                                                                var v64 float64 = (float64(float32(((0.1 + ((float64(1.5) + 1e10) * 0.0)) + -2.25))) * 1.5)
                                                                var v65 uint32 = 42
                                                                var v66 int32 = -2147483648
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                switch (("\n\t" + "\n\t") + v3) {
                                                                case "":
                                                                    v4 -= v4
                                                                case "hello":
                                                                    v6 = append(v6, (fi35 + 147))
                                                                default:
                                                                    v4 = v4
                                                                }
                                                                var v67 float32 = 3.14159
                                                                v63 = (v63 | v63)
                                                                _ = v67
                                                            }
                                                        }
                                                        if (!false) {
                                                            v68 := GenS3{F0: 42, F1: 1000}
                                                            v1 = (1000 != 1)
                                                            _ = v68
                                                        }
                                                        _ = v63
                                                    } else {
                                                        if true {
                                                            switch -1 {
                                                            case 32767:
                                                                if false {
                                                                    if v1 {
                                                                        v69 := GenS9{F0: fi35, F1: -1, F2: 2147483647}
                                                                        if v29 {
                                                                            fmt.Println(v59)
                                                                        } else {
                                                                            v70 := GenS10{F0: 35}
                                                                            _ = v70
                                                                        }
                                                                        v71 := GenS8{F0: (-(fi56))}
                                                                        _ = v69
                                                                        _ = v71
                                                                    }
                                                                    if ((0 == -1) && v29) {
                                                                        v3 = v3
                                                                    } else {
                                                                        v6 = append(v6, (^(fi35)))
                                                                        var v72 uint16 = 2
                                                                        if v29 {
                                                                            var v73 uint64 = 240
                                                                            v3 = ""
                                                                            v74 := GenS11{F0: 54}
                                                                            _ = v73
                                                                            _ = v74
                                                                        } else {
                                                                            switch fi56 {
                                                                            case 179:
                                                                                v75 := GenS1{F0: (^(v4)), F1: 42}
                                                                                _ = v75
                                                                            case 44:
                                                                                if v29 {
                                                                                    _ = len(v5)
                                                                                    v76 := GenS2{F0: v72, F1: 32767}
                                                                                    _ = v76
                                                                                }
                                                                            case -1:
                                                                                var v77 int8 = v2
                                                                                _ = v77
                                                                            }
                                                                            v78 := GenS9{F0: (^(fi35)), F1: int64(v4), F2: 1}
                                                                            if ((v2 / (((v2 << ((uint(v72)) & 63))) | 1)) > (-(v2))) {
                                                                                _ = len(v5)
                                                                                var v79 int8 = v2
                                                                                switch 128 {
                                                                                case -1:
                                                                                    v80 := GenS12{F0: 32767}
                                                                                    _ = v80
                                                                                default:
                                                                                    v81 := GenS13{F0: (fi56 / ((fi35) | 1)), F1: (v72 | (v72 | v72)), F2: 0}
                                                                                    _ = v81
                                                                                }
                                                                                var v82 int = fi35
                                                                                v79 = (-(v2))
                                                                                fmt.Println(v79, v82)
                                                                                v72 = (v72 ^ v72)
                                                                                _ = v79
                                                                                _ = v82
                                                                            }
                                                                            _ = v78
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                    fmt.Println(v29)
                                                                    v1 = v1
                                                                } else {
                                                                    var v83 uint16 = 65535
                                                                    v84 := GenS14{F0: (v83 >> ((4294967295) & 63))}
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                            default:
                                                                var v85 uint8 = 175
                                                                _ = v85
                                                            }
                                                            switch (fi35 / ((4294967295) | 1)) {
                                                            case 794:
                                                                fmt.Println(v29)
                                                            case 2147483647:
                                                                var v86 bool = v29
                                                                _ = v86
                                                            case 0:
                                                                if (!(1 != 0)) {
                                                                    fmt.Println(v1, v29, v1)
                                                                } else {
                                                                    if (v29 || ("\n\t" > "foo bar")) {
                                                                        var v87 string = v59
                                                                        _ = v87
                                                                    } else {
                                                                        fmt.Println(fi35)
                                                                        fmt.Println(fi56, v29, v4)
                                                                        var v88 int16 = -1
                                                                        _ = v88
                                                                    }
                                                                    if true {
                                                                        if v1 {
                                                                            v2 += v2
                                                                            v89 := GenS11{F0: v4}
                                                                            _ = v89
                                                                        } else {
                                                                            fmt.Println(v29, fi56)
                                                                            var v90 uint = 18446744073709551615
                                                                            var v91 float32 = (-2.25 + 1e10)
                                                                            if v29 {
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            } else {
                                                                                switch fi35 {
                                                                                case -851:
                                                                                    v92 := GenS2{F0: 65535, F1: -1}
                                                                                    _ = v92
                                                                                case -1:
                                                                                    if v29 {
                                                                                        switch int(v90) {
                                                                                        case 543:
                                                                                            v3 = v3
                                                                                        case 128:
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = fi35
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            if len(v5) > 0 {
                                                                                v5[0] = v59
                                                                            }
                                                                            v91 = (v91 / 1e10)
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                        fmt.Println(v59, fi56)
                                                                    } else {
                                                                        v93 := GenS3{F0: 2147483647, F1: 306}
                                                                        if (549 <= 63) {
                                                                            fmt.Println(v29, v4)
                                                                            v94 := GenS12{F0: (fi35 ^ (fi56 << ((255) & 63)))}
                                                                            _ = v94
                                                                        } else {
                                                                            if true {
                                                                                v2 = 1
                                                                                if (42 <= v4) {
                                                                                    switch fi56 {
                                                                                    case -1:
                                                                                        fmt.Println(v29, v4, v2)
                                                                                    case -53:
                                                                                        var v95 bool = (0.1 <= 3.14159)
                                                                                        _ = v95
                                                                                    }
                                                                                    var v96 uint32 = v4
                                                                                    _ = v96
                                                                                }
                                                                            } else {
                                                                                var v97 int = 32768
                                                                                v98 := GenS15{F0: 255}
                                                                                fmt.Println(v3, v1)
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                            fmt.Println(v3, fi35)
                                                                        }
                                                                        _ = v93
                                                                    }
                                                                    v1 = ((v2 + v2) < -13)
                                                                }
                                                            default:
                                                                v29 = true
                                                            }
                                                            v2 -= (v2 | (v2 << ((1) & 63)))
                                                        } else {
                                                            var v99 bool = (1e10 >= 3.14159)
                                                            switch v59 {
                                                            case "\n\t":
                                                                var v100 float32 = -2.25
                                                                _ = v100
                                                            default:
                                                                fmt.Println(fi35, v4, fi35)
                                                            }
                                                            _ = v99
                                                        }
                                                    }
                                                    _ = v59
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = ""
                                                }
                                                v1 = (-805 < -1)
                                            } else {
                                                fmt.Println(fi35, v4, v3)
                                            }
                                            v2 = -107
                                        } else {
                                            var v101 int = (-(fi35))
                                            v102 := GenS16{F0: 32767}
                                            _ = v101
                                            _ = v102
                                        }
                                    }
                                    fmt.Println(v3)
                                    _ = len(v6)
                                }
                                _ = fi56
                            }
                            var v103 int32 = -884
                            _ = v103
                        }
                        for fi104 := 0; fi104 < 12; fi104++ {
                            v1 = (fi104 == -461)
                            v4 -= v4
                            v105 := GenS10{F0: 42}
                            _ = fi104
                            _ = v105
                        }
                    }
                }
            } else {
                for fi106 := 0; fi106 < 8; fi106++ {
                    switch fi35 {
                    case 788:
                        v1 = true
                    case 0:
                        if false {
                            switch v3 {
                            case "a":
                                v107 := GenS11{F0: 65535}
                                _ = v107
                            case "":
                                v29 = (v29 || (((0 < 42) && true) && v29))
                            case "foo bar":
                                v4 -= 1
                            default:
                                if ((v3 + "") >= v3) {
                                    if ((-2147483648 <= 1) && v29) {
                                        if false {
                                            v108 := GenS8{F0: fi106}
                                            _ = v108
                                        } else {
                                            var v109 uint64 = 716
                                            _ = v109
                                        }
                                        v3 = "a"
                                        var v110 uint8 = 0
                                        _ = v110
                                    }
                                } else {
                                    v1 = v1
                                    v3 = (((("" + "") + (v3 + ("foo bar" + v3))) + (v3 + ((("hello" + v3) + (v3 + v3)) + v3))) + "foo bar")
                                }
                            }
                            if (false || ((18446744073709551615 > 4294967295) && v1)) {
                                v3 = "hello"
                                var v111 int16 = -1
                                if v29 {
                                    var v112 uint = 65535
                                    if ((^(fi35)) == fi35) {
                                        if len(v6) > 0 {
                                            v6[0] = (fi35 | fi35)
                                        }
                                    } else {
                                        if (2 < 1) {
                                            _ = len(v5)
                                        } else {
                                            v3 = (((v3 + v3) + v3) + v3)
                                            if len(v5) > 0 {
                                                v5[0] = ((v3 + "foo bar") + v3)
                                            }
                                            v113 := GenS17{F0: 1, F1: -1, F2: 38}
                                            _ = v113
                                        }
                                        v4 += (v4 >> ((v112) & 63))
                                        _ = len(v5)
                                    }
                                    fmt.Println(fi35, v111, v3)
                                    _ = v112
                                }
                                _ = v111
                            }
                            var v114 uint16 = 1
                            switch 1 {
                            case -661:
                                if (((("\n\t" + v3) + "\n\t") <= v3) || false) {
                                    var v115 int32 = 0
                                    _ = len(v6)
                                    var v116 float32 = 3.14159
                                    _ = v115
                                    _ = v116
                                } else {
                                    if true {
                                        var v117 uint64 = 42
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v118 bool = v29
                                        _ = v117
                                        _ = v118
                                    }
                                    _ = len(v5)
                                    v119 := GenS18{F0: v1}
                                    _ = v119
                                }
                            case 1:
                                v5 = append(v5, "foo bar")
                            case 9223372036854775807:
                                if (-1 < -2147483648) {
                                    v120 := GenS9{F0: fi106, F1: -655, F2: 42}
                                    var v121 int16 = -1
                                    fmt.Println(v2, v29)
                                    var v122 uint64 = 0
                                    _ = v120
                                    _ = v121
                                    _ = v122
                                } else {
                                    v123 := GenS5{F0: 127, F1: 1, F2: 9223372036854775807}
                                    _ = v123
                                }
                            }
                            _ = v114
                        } else {
                            v2 = (v2 * v2)
                            if (((0.0 / 1.0) + 1e10) > 1e10) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            if len(v5) > 0 {
                                v5[0] = v3
                            }
                            var v124 uint = 0
                            v5 = append(v5, (v3 + v3))
                            _ = v124
                        }
                    case 32767:
                        var v125 string = ("hello" + (v3 + ""))
                        _ = v125
                    }
                    if (float32(float64(((1e10 * 0.0) / 3.14159))) != (0.1 / float32(3.14159))) {
                        switch (fi35 << ((2) & 63)) {
                        case 992:
                            fmt.Println(v4, v2, v29)
                        case -2147483648:
                            v1 = v29
                        case 4294967295:
                            if v1 {
                                fmt.Println(fi106, fi35)
                                v6 = append(v6, -1)
                                switch 1 {
                                case 32767:
                                    switch (fi106 - fi35) {
                                    case 128:
                                        if v29 {
                                            fmt.Println(v2, v1, v1)
                                            var v126 uint32 = v4
                                            _ = v126
                                        } else {
                                            switch (-(fi35)) {
                                            case 4294967295:
                                                var v127 uint8 = 159
                                                _ = v127
                                            case -927:
                                                v128 := GenS19{F0: 1000, F1: -9223372036854775808}
                                                _ = v128
                                            }
                                            switch ((v3 + v3) + v3) {
                                            case "hello":
                                                var v129 int16 = 601
                                                _ = v129
                                            case "\n\t":
                                                switch v3 {
                                                case "\n\t":
                                                    if v1 {
                                                        switch (v3 + "\n\t") {
                                                        case "hello":
                                                            _ = len(v5)
                                                        case "a":
                                                            v2 = 100
                                                        case "foo bar":
                                                            v130 := GenS20{F0: (2 <= 18446744073709551615)}
                                                            _ = v130
                                                        }
                                                    }
                                                case "foo bar":
                                                    v131 := GenS2{F0: 32767, F1: 32767}
                                                    _ = v131
                                                default:
                                                    v1 = (42 != -1)
                                                }
                                            case "a":
                                                var v132 int16 = 42
                                                _ = v132
                                            }
                                            if len(v6) > 0 {
                                                v6[0] = (fi106 ^ fi35)
                                            }
                                        }
                                    default:
                                        v5 = append(v5, (("hello" + v3) + v3))
                                        fmt.Println(v3, fi35, v4)
                                    }
                                }
                            } else {
                                var v133 bool = v29
                                if false {
                                    var v134 int8 = -56
                                    v29 = v1
                                    v135 := GenS15{F0: 68}
                                    _ = v134
                                    _ = v135
                                }
                                fmt.Println(v1, v133)
                                _ = v133
                            }
                        }
                    } else {
                        if v29 {
                            v29 = true
                            v1 = true
                            var v136 bool = false
                            var v137 uint16 = 0
                            _ = v136
                            _ = v137
                        } else {
                            if (910 == 506) {
                                if (v2 < v2) {
                                    if ((v1 || (true || v29)) || v1) {
                                        fmt.Println(fi35, fi106, v3)
                                    }
                                } else {
                                    if (718 < -303) {
                                        switch fi35 {
                                        case 1:
                                            v138 := GenS6{F0: 33, F1: v2}
                                            _ = v138
                                        case -628:
                                            var v139 float64 = float64(3.14159)
                                            _ = v139
                                        case -522:
                                            fmt.Println(fi35, fi106)
                                        default:
                                            _ = len(v5)
                                        }
                                        v4 -= v4
                                        fmt.Println(v1, v29, v29)
                                        v2 = (v2 | v2)
                                    }
                                    var v140 float64 = 3.14159
                                    if v29 {
                                        v5 = append(v5, ((v3 + "hello") + v3))
                                    } else {
                                        var v141 int32 = 42
                                        v140 = v140
                                        _ = v141
                                    }
                                    _ = v140
                                }
                                v142 := GenS21{F0: 32767, F1: (!(32767 >= 244)), F2: v3}
                                _ = v142
                            }
                        }
                    }
                    if (!((v4 % ((456) | 1)) == 1)) {
                        var v143 int16 = 32767
                        switch fi35 {
                        case -233:
                            var v144 float32 = (-2.25 * 3.14159)
                            _ = v144
                        case 127:
                            v29 = (936 != 203)
                        case -427:
                            v4 -= (v4 % ((v4) | 1))
                        }
                        _ = len(v5)
                        _ = v143
                    } else {
                        var v145 int64 = 528
                        _ = v145
                    }
                    _ = fi106
                }
                v3 = v3
                v3 = (v3 + v3)
            }
            v5 = append(v5, (v3 + v3))
            fmt.Println(v1)
            _ = fi35
        }
        var v146 uint8 = 0
        v2 = (v2 / (((v2 + -128)) | 1))
        _ = v146
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi147 := 0; fi147 < 11; fi147++ {
        v5 = append(v5, "\n\t")
        v29 = ((3.14159 + 1e10) < 0.1)
        var v148 int64 = -1
        _ = fi147
        _ = v148
    }
    var v149 uint32 = (v4 & (v4 | v4))
    if (((v4 / ((4294967295) | 1)) != 65535) || v1) {
        for fi150 := 0; fi150 < 2; fi150++ {
            v151 := GenS17{F0: 18446744073709551615, F1: 42, F2: v2}
            if (-106 >= 32767) {
                v4 += v4
            }
            _ = fi150
            _ = v151
        }
        fmt.Println(v2, v4, v29)
        _ = len(v6)
    } else {
        for fi152 := 0; fi152 < 10; fi152++ {
            var v153 uint = 2
            for fi154 := 0; fi154 < 20; fi154++ {
                v3 = "\n\t"
                v1 = v29
                fmt.Println(v153)
                _ = fi154
            }
            v3 = "foo bar"
            _ = fi152
            _ = v153
        }
    }
    var v155 bool = v1
    for fi156 := 0; fi156 < 1; fi156++ {
        var v157 uint16 = 0
        _ = fi156
        _ = v157
    }
    if true {
        var v158 string = ""
        if v155 {
            v1 = (9223372036854775807 == 9223372036854775807)
        }
        _ = v158
    }
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("v155:", v155)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v149
    _ = v155
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
