package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint64, p1 string) uint {
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            fmt.Println(fi5, fi6)
            _ = fi6
        }
        v7 := []int{}
        _ = fi5
        _ = v7
    }
    if true {
        v8 := GenS1{F0: -1, F1: 671}
        v9 := []int{879}
        v10 := GenS1{F0: -2147483648, F1: -39}
        _ = v10
        _ = v8
        _ = v9
    } else {
        var v11 int8 = -127
        v11 += v11
        _ = v11
    }
    if (false && (1 > -32768)) {
        p1 = "\n\t"
    } else {
        var v12 float32 = (((0.1 * (float32(float64(1e10)) * 3.14159)) + 0.1) + (0.0 / -2.25))
        v13 := []int{32767}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 793
}
func genFn1(p0 int64, p1 int, p2 uint8) string {
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint16 = (v14 << ((genFn0(4294967295, "foo bar")) & 63))
        var v17 int8 = 127
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        switch p1 {
        case 0:
            var v19 float32 = 1.5
            _ = v19
        case 42:
            if false {
                v20 := GenS1{F0: -2147483648, F1: (p0 - 394)}
                var v21 int64 = p0
                _ = v20
                _ = v21
            } else {
                var v22 bool = ("" < "\n\t")
                for fi23 := 0; fi23 < 18; fi23++ {
                    v14 -= (v14 << ((18446744073709551615) & 63))
                    _ = fi23
                }
                _ = v22
            }
        default:
            v24 := GenS2{F0: 0, F1: p2}
            _ = v24
        }
        v14 -= 32767
        for fi25 := 0; fi25 < 5; fi25++ {
            v26 := GenS1{F0: 1, F1: 1}
            var v27 int8 = 41
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi18
    }
    var v28 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v28
    return ""
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 int16 = -770
    var v4 string = (("\n\t" + "foo bar") + "hello")
    var v29 int16 = v3
    fmt.Println(v29)
    var v30 uint32 = 807
    for fi31 := 0; fi31 < 7; fi31++ {
        for fi32 := 0; fi32 < 6; fi32++ {
            switch fi32 {
            case 128:
                v29 += (^(v3))
            case 765:
                var v33 int16 = v3
                _ = v33
            }
            var v34 uint64 = 2
            v35 := v2[""]
            _ = fi32
            _ = v34
            _ = v35
        }
        v4 = v4
        v30 = v30
        _ = fi31
    }
    var v36 int8 = 0
    v2[((("\n\t" + v4) + v4) + v4)] = "a"
    if (!true) {
        switch "" {
        case "\n\t":
            v29 += v3
        case "":
            var v37 bool = ((!(975 != 2)) || true)
            _ = v37
        }
        var v38 uint8 = 127
        _ = v38
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            var v41 uint64 = 995
            v42 := GenS3{F0: v41, F1: (fi40 % ((fi40) | 1)), F2: 227}
            if false {
                var v43 int16 = (v29 | 0)
                v44 := GenS4{F0: fi39}
                switch (fi39 << ((0) & 63)) {
                case -294:
                    var v45 string = "hello"
                    fmt.Println(v41, v43, fi39)
                    if true {
                        v3 = (v29 & v43)
                        var v46 uint16 = 2
                        _ = v46
                    } else {
                        v47 := GenS5{F0: 42}
                        fmt.Println(v29, v43, v43)
                        delete(v1, (fi39 / ((fi40) | 1)))
                        _ = v47
                    }
                    fmt.Println(v43, v3, v36)
                    fmt.Println(v3, v4, v43)
                    fmt.Println(v43)
                    var v48 float64 = 3.14159
                    _ = v45
                    _ = v48
                case -227:
                    fmt.Println(v30, fi39, v41)
                case -453:
                    var v49 uint = 536
                    _ = v49
                default:
                    var v50 float64 = (0.0 / 1.0)
                    _ = v50
                }
                _ = v43
                _ = v44
            }
            v29 = v3
            switch (fi40 ^ 128) {
            case -9223372036854775808:
                switch (fi40 << ((2) & 63)) {
                case 128:
                    v30 = v30
                case -967:
                    var v51 int64 = 9223372036854775807
                    _ = v51
                default:
                    var v52 uint = 597
                    _ = v52
                }
            case -1:
                var v53 uint = 909
                _ = v53
            case 2147483647:
                fmt.Println(v4)
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    if true {
        for fi54 := 0; fi54 < 3; fi54++ {
            if true {
                var v55 int64 = -1
                _ = v55
            } else {
                if (128 <= (fi54 + 32768)) {
                    var v56 float64 = 1.5
                    var v57 int64 = -1
                    _ = v56
                    _ = v57
                } else {
                    for fi58 := 0; fi58 < 18; fi58++ {
                        delete(v2, ((v4 + ((("a" + (v4 + ("\n\t" + "a"))) + (v4 + v4)) + "a")) + v4))
                        v59 := GenS6{F0: -1}
                        fmt.Println(fi58, v29)
                        _ = fi58
                        _ = v59
                    }
                    v4 = v4
                    v60 := []int{int(v30), fi54, (fi54 << ((18446744073709551615) & 63)), fi54}
                    _ = v60
                }
            }
            switch v4 {
            case "foo bar":
                for fi61 := 0; fi61 < 13; fi61++ {
                    v62 := []int{}
                    v4 = ("foo bar" + ((v4 + v4) + "a"))
                    if false {
                        v62 = append(v62, fi61)
                    }
                    _ = fi61
                    _ = v62
                }
            default:
                var v63 uint8 = 4
                _ = v63
            }
            fmt.Println(v4, v36, fi54)
            _ = fi54
        }
        v64 := GenS6{F0: 0}
        var v65 int16 = -555
        _ = v64
        _ = v65
    }
    var v66 float64 = (0.1 - (float64((0.0 + float32(1.5))) * ((-2.25 - 3.14159) * (1.5 * (((0.0 * -2.25) / -2.25) / (1.5 * ((1.5 - (3.14159 - (-2.25 * 0.1))) - 3.14159)))))))
    if (v29 != v29) {
        v67 := []int{}
        if ((1.5 * float32(v66)) <= 0.1) {
            if false {
                var v68 float64 = 0.1
                if ("hello" == v4) {
                    switch ("a" + ("a" + v4)) {
                    case "foo bar":
                        switch v4 {
                        case "a":
                            switch v4 {
                            case "foo bar":
                                if false {
                                    v66 += v68
                                    var v69 uint16 = 0
                                    if len(v67) >= 2 {
                                        v67 = v67[1:len(v67)]
                                    }
                                    _ = v69
                                } else {
                                    v29 += (v29 << ((18446744073709551615) & 63))
                                    switch (((((v4 + ("foo bar" + ((v4 + "foo bar") + v4))) + v4) + ((v4 + "a") + "\n\t")) + v4) + ("a" + ((v4 + v4) + "\n\t"))) {
                                    case "hello":
                                        var v70 int32 = 498
                                        _ = v70
                                    case "\n\t":
                                        for fi71 := 0; fi71 < 18; fi71++ {
                                            if (-1 < (fi71 + (fi71 ^ (fi71 / (((fi71 & (fi71 >> ((255) & 63)))) | 1))))) {
                                                if (599 <= 4294967295) {
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v72 float32 = 0.1
                                                        _ = v72
                                                    case "":
                                                        var v73 string = v4
                                                        _ = v73
                                                    default:
                                                        if true {
                                                            for fi74 := 0; fi74 < 4; fi74++ {
                                                                v75 := GenS6{F0: 9223372036854775807}
                                                                if false {
                                                                    var v76 bool = true
                                                                    if ((3.14159 < (float32(1.5) * (0.0 + (((1.5 + (-2.25 / float32(v68))) / (3.14159 * 3.14159)) * 3.14159)))) || v76) {
                                                                        switch (fi74 ^ -487) {
                                                                        case 1:
                                                                            v77 := GenS7{F0: -748, F1: (^(v3))}
                                                                            _ = v77
                                                                        case -7:
                                                                            _ = len(v67)
                                                                        case -258:
                                                                            switch fi71 {
                                                                            case -9223372036854775808:
                                                                                var v78 int64 = 0
                                                                                _ = v78
                                                                            case 32768:
                                                                                v67 = append(v67, 856)
                                                                            case 4294967295:
                                                                                v36 -= -14
                                                                            }
                                                                        default:
                                                                            var v79 float32 = (((3.14159 - (float32(((v68 + 0.1) - v68)) * (1.5 * 0.1))) - (((3.14159 - 1.5) * 1e10) + float32(v68))) * ((1e10 + 1.5) / float32(v66)))
                                                                            _ = v79
                                                                        }
                                                                    } else {
                                                                        var v80 uint64 = 534
                                                                        v30 -= v30
                                                                        _ = len(v1)
                                                                        _ = v80
                                                                    }
                                                                    v81 := GenS1{F0: 933, F1: 9223372036854775807}
                                                                    _ = v76
                                                                    _ = v81
                                                                } else {
                                                                    v29 = v3
                                                                    var v82 int16 = (v3 + (^(v3)))
                                                                    _ = len(v1)
                                                                    _ = v82
                                                                }
                                                                _ = fi74
                                                                _ = v75
                                                            }
                                                        }
                                                    }
                                                    v3 += v3
                                                }
                                            } else {
                                                v3 = v29
                                                var v83 int32 = 2147483647
                                                var v84 bool = false
                                                _ = v83
                                                _ = v84
                                            }
                                            _ = fi71
                                        }
                                    default:
                                        v85 := GenS2{F0: 18446744073709551615, F1: 127}
                                        _ = v85
                                    }
                                    v86 := GenS8{F0: 2}
                                    _ = v86
                                }
                            case "a":
                                var v87 bool = ((4294967295 <= 0) || false)
                                _ = v87
                            }
                        }
                    case "a":
                        v4 = ""
                    case "\n\t":
                        for fi88 := 0; fi88 < 3; fi88++ {
                            if (fi88 >= 944) {
                                var v89 uint64 = 799
                                _ = v89
                            } else {
                                for fi90 := 0; fi90 < 7; fi90++ {
                                    var v91 float32 = -2.25
                                    v4 = v4
                                    _ = fi90
                                    _ = v91
                                }
                            }
                            _ = fi88
                        }
                    default:
                        if true {
                            if false {
                                if len(v67) >= 2 {
                                    v67 = v67[1:len(v67)]
                                }
                                v68 += -2.25
                                for fi92 := 0; fi92 < 15; fi92++ {
                                    if len(v67) >= 2 {
                                        v67 = v67[1:len(v67)]
                                    }
                                    _ = fi92
                                }
                            }
                            if len(v67) >= 2 {
                                v67 = v67[1:len(v67)]
                            }
                            v93 := GenS9{F0: 65535}
                            _ = v93
                        } else {
                            for fi94 := 0; fi94 < 14; fi94++ {
                                v30 = 4294967295
                                _ = fi94
                            }
                        }
                    }
                    fmt.Println(v30, v66, v68)
                }
                _ = v68
            } else {
                var v95 uint32 = v30
                _ = v95
            }
            switch v4 {
            case "":
                for fi96 := 0; fi96 < 17; fi96++ {
                    v36 += v36
                    if false {
                        var v97 int8 = v36
                        switch int(v29) {
                        case 128:
                            var v98 float64 = 1.5
                            _ = v98
                        case 2147483647:
                            var v99 uint = 18446744073709551615
                            _ = v99
                        default:
                            switch v4 {
                            case "hello":
                                switch (fi96 | fi96) {
                                case 32768:
                                    if true {
                                        switch 127 {
                                        case 1:
                                            v100 := GenS10{F0: 2147483647}
                                            _ = v100
                                        case -2147483648:
                                            var v101 uint8 = 255
                                            _ = v101
                                        }
                                    } else {
                                        fmt.Println(v66, v30)
                                    }
                                case 1:
                                    var v102 int32 = 42
                                    _ = v102
                                default:
                                    v97 = (v97 >> ((0) & 63))
                                }
                            default:
                                var v103 int8 = (v36 + (v36 ^ (v97 | v97)))
                                _ = v103
                            }
                        }
                        _ = v97
                    }
                    if false {
                        for fi104 := 0; fi104 < 6; fi104++ {
                            var v105 uint = 255
                            switch fi104 {
                            case -9223372036854775808:
                                v66 += v66
                            case 127:
                                v30 += v30
                            case 9223372036854775807:
                                var v106 int = fi104
                                _ = v106
                            }
                            switch (v4 + (("" + v4) + v4)) {
                            case "a":
                                switch ("\n\t" + v4) {
                                case "\n\t":
                                    if (!(1 == 2)) {
                                        if (!(((v105 % (((v105 % ((v105) | 1))) | 1)) != 255) || (!false))) {
                                            if (!(((3.14159 - v66) + v66) == (v66 - v66))) {
                                                var v107 int8 = v36
                                                _ = v107
                                            } else {
                                                if false {
                                                    var v108 uint = v105
                                                    var v109 int8 = v36
                                                    v110 := v1[fi104]
                                                    _ = v108
                                                    _ = v109
                                                    _ = v110
                                                }
                                            }
                                            v30 = v30
                                            v111 := GenS11{F0: v36}
                                            _ = v111
                                        }
                                        if false {
                                            v112 := GenS6{F0: 42}
                                            var v113 string = v4
                                            _ = v112
                                            _ = v113
                                        }
                                    } else {
                                        var v114 float32 = 1.5
                                        var v115 uint64 = 42
                                        fmt.Println(v36, v4, v105)
                                        _ = v114
                                        _ = v115
                                    }
                                default:
                                    var v116 int16 = (v3 / ((v29) | 1))
                                    _ = v116
                                }
                                var v117 int16 = v3
                                _ = v117
                            case "\n\t":
                                fmt.Println(v66)
                            case "hello":
                                fmt.Println(v36)
                            default:
                                var v118 uint64 = 4294967295
                                _ = v118
                            }
                            _ = fi104
                            _ = v105
                        }
                    }
                    _ = fi96
                }
            case "a":
                fmt.Println(v29)
            case "foo bar":
                switch 670 {
                case 13:
                    v119 := GenS12{F0: (((v36 << ((383) & 63)) > v36) || true), F1: 128}
                    _ = v119
                default:
                    if true {
                        v120 := GenS9{F0: (^(v30))}
                        v121 := GenS12{F0: true, F1: 2}
                        _ = v120
                        _ = v121
                    } else {
                        switch "a" {
                        case "foo bar":
                            if true {
                                fmt.Println(v36, v30)
                                switch -1 {
                                case 361:
                                    v122 := GenS3{F0: 771, F1: -9223372036854775808, F2: 1}
                                    _ = v122
                                case 9223372036854775807:
                                    var v123 uint = 255
                                    _ = v123
                                default:
                                    if (true && (v4 <= (("\n\t" + v4) + (v4 + "")))) {
                                        v36 = (v36 >> ((977) & 63))
                                        for fi124 := 0; fi124 < 8; fi124++ {
                                            v125 := GenS6{F0: -558}
                                            _ = fi124
                                            _ = v125
                                        }
                                        var v126 uint = 65535
                                        _ = v126
                                    }
                                }
                                var v127 int16 = (^(v29))
                                _ = v127
                            }
                        case "":
                            var v128 int = 42
                            _ = v128
                        default:
                            v129 := GenS2{F0: 42, F1: 255}
                            _ = v129
                        }
                    }
                }
            default:
                var v130 string = v4
                _ = v130
            }
            v131 := GenS7{F0: (v3 / ((v29) | 1)), F1: (v3 % (((v3 / ((v29) | 1))) | 1))}
            _ = v131
        }
        if len(v67) > 0 {
            v67[0] = -139
        }
        _ = v67
    } else {
        var v132 bool = true
        var v133 int32 = -2147483648
        delete(v2, v4)
        _ = v132
        _ = v133
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v66
}
