package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint64; F2 string }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint16; F1 int64; F2 int }
type GenS8 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS9 struct { F0 string; F1 int16 }
func genFn0() bool {
    v5 := GenS1{F0: -1, F1: 32767, F2: 0}
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint = 246
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 uint16 = 32767
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]uint32)
    v4 := []int{32767, 32768}
    if (v2 <= 2147483647) {
        v1 += v1
    } else {
        switch 32768 {
        case -2147483648:
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: -2147483648, F1: (fi9 + 926), F2: 100}
                var v11 uint8 = 255
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := GenS2{F0: fi12, F1: -431, F2: 4294967295}
                    v14 := GenS2{F0: int(v2), F1: -1, F2: 2}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        case 474:
            var v15 int64 = 9223372036854775807
            _ = v15
        case 4294967295:
            _ = len(v4)
        default:
            if false {
                v1 -= 1
                for fi16 := 0; fi16 < 17; fi16++ {
                    v17 := v3[(^(v2))]
                    var v18 uint32 = 2
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                fmt.Println(v1)
            } else {
                for fi19 := 0; fi19 < 8; fi19++ {
                    if false {
                        v1 = 2
                    } else {
                        v20 := GenS1{F0: -2147483648, F1: 128, F2: 127}
                        if false {
                            var v21 int32 = 2147483647
                            if (false && (!false)) {
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    switch (-(fi22)) {
                                    case 4294967295:
                                        _ = len(v4)
                                    case -9223372036854775808:
                                        if (!(true && false)) {
                                            v21 = (v21 >> ((65535) & 63))
                                            v2 += (v2 * 2147483647)
                                            var v23 int = (-(fi19))
                                            _ = v23
                                        } else {
                                            v24 := GenS1{F0: v21, F1: fi19, F2: 65}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v24
                                        }
                                        v4 = append(v4, fi19)
                                        var v25 uint = 255
                                        var v26 string = "a"
                                        _ = v25
                                        _ = v26
                                    default:
                                        if false {
                                            v27 := GenS3{F0: (v2 & v2), F1: (v21 % ((0) | 1))}
                                            _ = v27
                                        }
                                    }
                                    _ = fi22
                                }
                                delete(v3, (v2 >> ((513) & 63)))
                            } else {
                                if (((1.5 * 1e10) * 0.1) != 1.5) {
                                    fmt.Println(v21, v1)
                                } else {
                                    _ = len(v3)
                                    fmt.Println(fi19, fi19)
                                }
                            }
                            _ = v21
                        } else {
                            var v28 float32 = ((3.14159 / 1e10) + 1e10)
                            if ((!(3.14159 != (0.0 - 1.5))) || (106 >= 42)) {
                                fmt.Println(v28)
                            } else {
                                var v29 float64 = 1e10
                                _ = v29
                            }
                            _ = v28
                        }
                        _ = v20
                    }
                    v30 := v3[(^(v2))]
                    _ = fi19
                    _ = v30
                }
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    switch ("" + "\n\t") {
    case "foo bar":
        for fi31 := 0; fi31 < 16; fi31++ {
            v1 = v1
            var v32 uint16 = (v1 << ((65535) & 63))
            _ = fi31
            _ = v32
        }
    case "\n\t":
        switch ("\n\t" + ("a" + "")) {
        case "foo bar":
            v2 = 1
        case "hello":
            v2 = (v2 | 2147483647)
        case "\n\t":
            _ = len(v3)
        default:
            if (!(((float32((3.14159 * (0.1 - (1e10 * (1e10 + (((1e10 / 1.5) + float64(0.0)) - ((-2.25 * 0.0) + ((-2.25 * float64(3.14159)) + 1e10)))))))) - 1e10) * 1e10) != -2.25)) {
                switch 32767 {
                case 42:
                    if false {
                        var v33 int = -237
                        _ = v33
                    } else {
                        if (!(18446744073709551615 != 735)) {
                            for fi34 := 0; fi34 < 15; fi34++ {
                                var v35 uint16 = (v1 ^ 65535)
                                v36 := GenS2{F0: fi34, F1: -456, F2: 18446744073709551615}
                                if len(v4) > 0 {
                                    v4[0] = fi34
                                }
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 8; fi37++ {
                                for fi38 := 0; fi38 < 16; fi38++ {
                                    v39 := GenS4{F0: 15}
                                    if (!(true && (!(127 > -1)))) {
                                        var v40 uint64 = 129
                                        _ = v40
                                    } else {
                                        if true {
                                            v2 = (^(v2))
                                            if false {
                                                switch "" {
                                                case "foo bar":
                                                    if true {
                                                        v41 := GenS1{F0: int32(v2), F1: fi37, F2: 127}
                                                        if (!(!true)) {
                                                            if true {
                                                                v2 += (^(v2))
                                                                var v42 string = "a"
                                                                _ = v42
                                                            }
                                                            v2 = v2
                                                            v1 += 2
                                                        } else {
                                                            var v43 int32 = 679
                                                            v44 := GenS1{F0: v43, F1: fi38, F2: 127}
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        switch fi38 {
                                                        case 2147483647:
                                                            v45 := GenS3{F0: (v2 * v2), F1: int32(v1)}
                                                            _ = v45
                                                        case 9223372036854775807:
                                                            if (true || ((128 > 2) && false)) {
                                                                if (2147483647 > -713) {
                                                                    var v46 int = 995
                                                                    _ = v46
                                                                }
                                                            } else {
                                                                v47 := GenS1{F0: -2147483648, F1: -1, F2: 127}
                                                                v1 = v1
                                                                _ = v47
                                                            }
                                                        }
                                                        _ = v41
                                                    } else {
                                                        var v48 float64 = (1e10 + (0.1 + (((1.5 - (0.0 * (1e10 - float64(1.5)))) * 1.5) + (1e10 + (1e10 - 0.1)))))
                                                        v1 -= (v1 / (((v1 << ((65535) & 63))) | 1))
                                                        v1 = 0
                                                        _ = v48
                                                    }
                                                case "\n\t":
                                                    v49 := v3[uint32(v1)]
                                                    _ = v49
                                                case "a":
                                                    v1 = (^(v1))
                                                }
                                                v50 := GenS3{F0: v2, F1: 650}
                                                _ = v50
                                            }
                                        } else {
                                            if ((fi38 % ((fi37) | 1)) < 128) {
                                                v51 := GenS5{F0: "a", F1: 42, F2: "foo bar"}
                                                _ = v51
                                            }
                                        }
                                        if false {
                                            var v52 uint32 = (v2 & 315)
                                            v4 = append(v4, 894)
                                            fmt.Println(v52, fi37)
                                            _ = v52
                                        } else {
                                            v53 := GenS1{F0: 375, F1: (fi38 ^ fi37), F2: -1}
                                            _ = v53
                                        }
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                                v1 = v1
                                fmt.Println(v2, v1, v2)
                                _ = fi37
                            }
                        } else {
                            v1 -= (v1 | v1)
                            switch 9223372036854775807 {
                            case 818:
                                v2 = v2
                            case 0:
                                v2 = (v2 / ((4294967295) | 1))
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                }
                var v54 float64 = 3.14159
                var v55 uint64 = uint64(v1)
                _ = v54
                _ = v55
            }
        }
    case "hello":
        v2 += (v2 * (^(v2)))
    }
    for fi56 := 0; fi56 < 19; fi56++ {
        var v57 bool = (!true)
        _ = len(v4)
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 17; fi58++ {
        switch (fi58 - fi58) {
        case 2147483647:
            if (-618 == 127) {
                var v59 string = "hello"
                for fi60 := 0; fi60 < 20; fi60++ {
                    v2 = v2
                    fmt.Println(v1, v1)
                    _ = fi60
                }
                _ = v59
            } else {
                var v61 uint16 = v1
                for fi62 := 0; fi62 < 15; fi62++ {
                    var v63 uint16 = (v1 >> ((295) & 63))
                    switch (-(fi62)) {
                    case 32768:
                        var v64 bool = false
                        _ = v64
                    case -2147483648:
                        var v65 bool = true
                        _ = v65
                    case 249:
                        var v66 uint32 = v2
                        _ = v66
                    default:
                        fmt.Println(v2, v61, v61)
                    }
                    _ = fi62
                    _ = v63
                }
                _ = v61
            }
        case 42:
            if false {
                var v67 uint32 = v2
                v68 := GenS6{F0: 42, F1: ((((("a" + "\n\t") + "") + "foo bar") + "\n\t") + "\n\t")}
                _ = v67
                _ = v68
            }
        case -292:
            fmt.Println(v2, fi58)
        }
        var v69 uint8 = 1
        switch "foo bar" {
        case "foo bar":
            var v70 uint = 7
            _ = v70
        default:
            v69 += 42
        }
        _ = fi58
        _ = v69
    }
    fmt.Println(v1, v2)
    delete(v3, 42)
    v71 := GenS7{F0: 32767, F1: 1, F2: 32767}
    if true {
        if (2147483647 < 0) {
            if (-9223372036854775808 >= -1) {
                fmt.Println(v2, v1)
                v1 += (v1 + 42)
                switch -1 {
                case 9223372036854775807:
                    var v72 int = -499
                    _ = v72
                default:
                    v1 = v1
                }
            } else {
                switch -279 {
                case -174:
                    var v73 int32 = 2147483647
                    _ = v73
                case 128:
                    v74 := GenS5{F0: "", F1: 734, F2: (("hello" + ("" + ("" + "foo bar"))) + ("a" + ("foo bar" + "foo bar")))}
                    _ = v74
                }
                switch -293 {
                case 32768:
                    v2 = 639
                case 0:
                    fmt.Println(v1, v2)
                default:
                    switch "hello" {
                    case "":
                        v75 := GenS8{F0: 0, F1: 42, F2: v2}
                        _ = v75
                    }
                }
            }
            fmt.Println(v2, v2)
        } else {
            var v76 float64 = 3.14159
            for fi77 := 0; fi77 < 8; fi77++ {
                var v78 bool = (269 == fi77)
                _ = fi77
                _ = v78
            }
            _ = v76
        }
        if (!("a" > (("hello" + "hello") + "foo bar"))) {
            if (-32768 < -511) {
                var v79 bool = true
                var v80 int16 = 1
                _ = v79
                _ = v80
            } else {
                var v81 int64 = 9223372036854775807
                switch ("" + "foo bar") {
                case "foo bar":
                    fmt.Println(v81, v81, v81)
                case "\n\t":
                    fmt.Println(v2, v81, v81)
                default:
                    var v82 bool = (!(!false))
                    _ = v82
                }
                for fi83 := 0; fi83 < 15; fi83++ {
                    delete(v3, (v2 << ((336) & 63)))
                    if false {
                        if (true && true) {
                            var v84 uint64 = 1
                            _ = v84
                        } else {
                            if (52 > 144) {
                                if false {
                                    var v85 uint32 = (v2 * v2)
                                    v85 = v85
                                    _ = v85
                                }
                                if false {
                                    switch 9223372036854775807 {
                                    case 9223372036854775807:
                                        v3[v2] = (^(v2))
                                    }
                                } else {
                                    var v86 int32 = 0
                                    if false {
                                        v81 = (v81 * v81)
                                        v87 := GenS1{F0: 42, F1: (^(fi83)), F2: 0}
                                        _ = v87
                                    } else {
                                        var v88 uint8 = 1
                                        fmt.Println(v2, v86)
                                        v88 = v88
                                        _ = v88
                                    }
                                    _ = v86
                                }
                            }
                            var v89 uint8 = 127
                            _ = v89
                        }
                        v1 += (^(v1))
                        if (255 < 253) {
                            fmt.Println(v1, fi83, fi83)
                        } else {
                            v90 := GenS5{F0: "", F1: 18446744073709551615, F2: ""}
                            switch "" {
                            case "":
                                if false {
                                    v2 = 0
                                    if len(v4) > 0 {
                                        v4[0] = fi83
                                    }
                                    var v91 uint = 42
                                    _ = v91
                                } else {
                                    v2 = (v2 - (v2 * v2))
                                    for fi92 := 0; fi92 < 10; fi92++ {
                                        v1 -= v1
                                        if (632 != (v2 * 2147483647)) {
                                            v93 := GenS4{F0: -128}
                                            _ = v93
                                        } else {
                                            _ = len(v3)
                                        }
                                        _ = fi92
                                    }
                                }
                            default:
                                v2 += v2
                            }
                            switch "a" {
                            case "foo bar":
                                if true {
                                    v2 += v2
                                }
                            }
                            _ = v90
                        }
                    } else {
                        if true {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(v1, fi83)
                        } else {
                            for fi94 := 0; fi94 < 18; fi94++ {
                                v95 := GenS6{F0: (v81 - v81), F1: "a"}
                                var v96 int = int(v1)
                                fmt.Println(fi94)
                                _ = fi94
                                _ = v95
                                _ = v96
                            }
                            _ = len(v4)
                            fmt.Println(v81, v1, v1)
                        }
                        var v97 float64 = 0.0
                        var v98 float32 = 0.0
                        _ = v97
                        _ = v98
                    }
                    _ = fi83
                }
                _ = v81
            }
        } else {
            if false {
                if (-1 > 42) {
                    if (127 <= 1) {
                        switch 558 {
                        case 127:
                            if true {
                                var v99 float32 = -2.25
                                _ = v99
                            } else {
                                if true {
                                    v100 := GenS4{F0: 42}
                                    if ((1e10 * 3.14159) > (float64(1e10) - -2.25)) {
                                        var v101 uint64 = 748
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v102 bool = false
                                        _ = v101
                                        _ = v102
                                    } else {
                                        switch -9223372036854775808 {
                                        case -749:
                                            v103 := GenS6{F0: 42, F1: "\n\t"}
                                            _ = v103
                                        default:
                                            v104 := GenS9{F0: "a", F1: 32767}
                                            _ = v104
                                        }
                                        v2 += (^(v2))
                                        var v105 int32 = -2147483648
                                        _ = v105
                                    }
                                    v106 := v3[(v2 ^ (v2 >> ((4294967295) & 63)))]
                                    _ = v100
                                    _ = v106
                                }
                                switch -934 {
                                case 1:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case 715:
                                    for fi107 := 0; fi107 < 12; fi107++ {
                                        var v108 int8 = 127
                                        v3[402] = 4294967295
                                        _ = fi107
                                        _ = v108
                                    }
                                case 9223372036854775807:
                                    var v109 float64 = 1e10
                                    _ = v109
                                }
                            }
                        case -1:
                            v110 := GenS2{F0: -249, F1: 0, F2: 486}
                            _ = v110
                        case 42:
                            v2 = (v2 % ((v2) | 1))
                        }
                        v2 -= (v2 ^ (v2 ^ v2))
                        for fi111 := 0; fi111 < 1; fi111++ {
                            var v112 uint64 = 4294967295
                            var v113 uint64 = (v112 / ((18446744073709551615) | 1))
                            _ = fi111
                            _ = v112
                            _ = v113
                        }
                    } else {
                        v1 = v1
                        v114 := GenS4{F0: 100}
                        var v115 float64 = ((-2.25 / -2.25) * 1e10)
                        _ = v114
                        _ = v115
                    }
                    for fi116 := 0; fi116 < 17; fi116++ {
                        if true {
                            v1 = (v1 ^ v1)
                            if (v1 < v1) {
                                v2 += (v2 ^ 708)
                                var v117 string = ""
                                var v118 uint8 = 19
                                _ = v117
                                _ = v118
                            } else {
                                var v119 uint32 = (v2 % (((^(v2))) | 1))
                                _ = v119
                            }
                        }
                        _ = fi116
                    }
                } else {
                    for fi120 := 0; fi120 < 9; fi120++ {
                        fmt.Println(v1, v1)
                        _ = fi120
                    }
                }
                v121 := GenS3{F0: 113, F1: 42}
                var v122 float64 = 3.14159
                _ = v121
                _ = v122
            }
        }
    }
    if true {
        var v123 uint16 = v1
        var v124 float32 = 1e10
        _ = v123
        _ = v124
    } else {
        if true {
            v2 -= v2
            v2 = v2
            v1 = 42
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi125 := 0; fi125 < 8; fi125++ {
                for fi126 := 0; fi126 < 9; fi126++ {
                    delete(v3, 1)
                    _ = fi126
                }
                _ = fi125
            }
            var v127 float64 = ((1e10 / 0.1) - 3.14159)
            _ = v127
        }
        v2 += (v2 * v2)
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
}
