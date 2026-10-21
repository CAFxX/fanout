package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint; F1 uint32; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 int32; F2 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32; F1 int; F2 int }
type GenS12 struct { F0 string; F1 uint32; F2 string }
type GenS13 struct { F0 int32; F1 uint; F2 int8 }
func genFn0(p0 string, p1 uint16) uint8 {
    p1 += 862
    if ((-454 > 150) && (42 < -128)) {
        v4 := []int{159, 619}
        if true {
            for fi5 := 0; fi5 < 18; fi5++ {
                p1 -= 2
                _ = len(v4)
                _ = fi5
            }
            var v6 uint8 = 128
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: 65535}
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v4
    } else {
        var v9 float32 = -2.25
        v10 := []int{0}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        switch fi11 {
        case -9223372036854775808:
            switch (fi11 ^ (fi11 & fi11)) {
            case 32767:
                var v12 int32 = 1
                _ = v12
            case -310:
                if false {
                    v13 := []int{fi11}
                    v13 = append(v13, 32768)
                    if false {
                        if false {
                            p1 = 65535
                            var v14 uint = 493
                            if (!true) {
                                for fi15 := 0; fi15 < 3; fi15++ {
                                    var v16 uint8 = 127
                                    v17 := GenS1{F0: 0}
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                            }
                            _ = v14
                        } else {
                            for fi18 := 0; fi18 < 7; fi18++ {
                                fmt.Println(fi18, fi18)
                                p0 = p0
                                p1 = 255
                                fmt.Println(p1)
                                p0 = "foo bar"
                                var v19 int8 = 0
                                switch "\n\t" {
                                case "":
                                    if len(v13) > 0 {
                                        v13[0] = (fi18 + (^(fi11)))
                                    }
                                case "hello":
                                    p0 = "a"
                                default:
                                    if (false && false) {
                                        _ = len(v13)
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                        }
                    } else {
                        if (9223372036854775807 <= 9223372036854775807) {
                            v20 := GenS1{F0: 407}
                            _ = v20
                        }
                    }
                    _ = v13
                } else {
                    v21 := GenS2{F0: true, F1: 65535, F2: 127}
                    _ = v21
                }
            default:
                if false {
                    p1 += (p1 - (p1 & p1))
                    if (true || (!false)) {
                        var v22 int32 = 668
                        p1 = p1
                        fmt.Println(p1, p1, v22)
                        _ = v22
                    } else {
                        var v23 int32 = 1
                        _ = v23
                    }
                } else {
                    v24 := []int{-2147483648}
                    _ = v24
                }
            }
        case -1:
            var v25 int32 = 2147483647
            _ = v25
        case -304:
            switch ((p0 + (("a" + "hello") + p0)) + p0) {
            case "\n\t":
                v26 := GenS3{F0: 42}
                _ = v26
            case "foo bar":
                v27 := []int{40, (fi11 - fi11), fi11, int(p1)}
                _ = v27
            case "a":
                for fi28 := 0; fi28 < 18; fi28++ {
                    if true {
                        p1 += (p1 / (((^(p1))) | 1))
                    }
                    _ = fi28
                }
            default:
                for fi29 := 0; fi29 < 6; fi29++ {
                    if false {
                        var v30 uint64 = 2
                        var v31 uint32 = 423
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS1{F0: 97}
                    var v33 uint8 = 128
                    v34 := GenS2{F0: (false || true), F1: p1, F2: 100}
                    _ = fi29
                    _ = v32
                    _ = v33
                    _ = v34
                }
            }
        default:
            p1 += (^(p1))
        }
        switch "foo bar" {
        case "":
            p1 = (^(p1))
        default:
            fmt.Println(p0, p1)
        }
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS1{F0: 619}
            p1 += (^(p1))
            var v37 int = fi11
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi11
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        p1 = p1
        _ = fi38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 int64 = 42
        var v41 int32 = -1
        _ = fi39
        _ = v40
        _ = v41
    }
    _ = p0
    _ = p1
    return 230
}
func genFn1(p0 int8, p1 uint, p2 uint32) uint8 {
    var v42 float64 = -2.25
    v43 := GenS4{F0: "a", F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    _ = v43
    return 24
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]uint32)
    var v3 uint32 = 599
    v44 := GenS5{F0: 493, F1: v3, F2: 32767}
    v45 := GenS1{F0: v3}
    switch (v1 * 42) {
    case 181:
        if (!(3.14159 == 1e10)) {
            v46 := GenS3{F0: 255}
            _ = v46
        }
    case -9223372036854775808:
        v47 := []int{(v1 - -2147483648)}
        _ = v47
    default:
        v1 = v1
    }
    if false {
        var v48 int32 = -625
        var v49 float64 = 1.5
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 2; fi50++ {
        var v51 float32 = 1.5
        _ = fi50
        _ = v51
    }
    if (!true) {
        for fi52 := 0; fi52 < 1; fi52++ {
            if ((-502 < 0) && false) {
                for fi53 := 0; fi53 < 12; fi53++ {
                    if (0.1 >= (1.5 * (0.1 + 0.0))) {
                        v54 := GenS6{F0: (284 != 18446744073709551615)}
                        _ = v54
                    } else {
                        v55 := GenS3{F0: 1}
                        _ = v55
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v56 int8 = 127
                        _ = v56
                    case "a":
                        var v57 int8 = 74
                        _ = v57
                    case "hello":
                        v58 := GenS7{F0: -558, F1: 42, F2: -14}
                        _ = v58
                    default:
                        fmt.Println(fi52)
                    }
                    _ = fi53
                }
                v59 := GenS7{F0: -32768, F1: 42, F2: 0}
                if true {
                    v60 := GenS7{F0: 1, F1: int32(v3), F2: -128}
                    _ = v60
                }
                _ = v59
            } else {
                var v61 int16 = 0
                v61 -= -1
                _ = v61
            }
            _ = fi52
        }
        v1 = v1
    } else {
        _ = len(v2)
        v62 := GenS8{F0: "a"}
        _ = v62
    }
    if true {
        if false {
            var v63 uint32 = (v3 ^ 1)
            var v64 uint8 = 255
            _ = v63
            _ = v64
        }
        fmt.Println(v3)
    } else {
        var v65 uint = 0
        if (float64(((1.5 + float32(0.0)) + -2.25)) == 0.1) {
            if false {
                for fi66 := 0; fi66 < 8; fi66++ {
                    v3 = (v3 / ((v3) | 1))
                    v65 -= v65
                    fmt.Println(v65)
                    _ = fi66
                }
                v67 := []int{v1, v1, 754, v1}
                _ = v67
            } else {
                v3 = (v3 ^ 1)
                switch v1 {
                case 32767:
                    for fi68 := 0; fi68 < 17; fi68++ {
                        for fi69 := 0; fi69 < 2; fi69++ {
                            v65 = v65
                            _ = fi69
                        }
                        v1 -= v1
                        v65 = v65
                        _ = fi68
                    }
                case 947:
                    v70 := []int{32767}
                    _ = v70
                default:
                    if (true && (v3 != v3)) {
                        if (2147483648 != (v3 >> ((v65) & 63))) {
                            var v71 float32 = -2.25
                            v72 := []int{-1, v1, v1}
                            for fi73 := 0; fi73 < 15; fi73++ {
                                v71 = 3.14159
                                var v74 bool = ((42 < -32768) || (127 < 29))
                                _ = fi73
                                _ = v74
                            }
                            _ = v71
                            _ = v72
                        }
                    }
                }
                if false {
                    v2[((((("foo bar" + "hello") + "foo bar") + ("hello" + "hello")) + "a") + "\n\t")] = (^(v3))
                    var v75 float64 = -2.25
                    if ((1 <= 499) && false) {
                        v76 := GenS9{F0: 1}
                        v77 := GenS10{F0: 32768}
                        _ = v76
                        _ = v77
                    } else {
                        v1 -= v1
                        v78 := GenS11{F0: (v3 >> ((42) & 63)), F1: v1, F2: (v1 << ((v65) & 63))}
                        v79 := GenS1{F0: v3}
                        _ = v78
                        _ = v79
                    }
                    _ = v75
                }
            }
        } else {
            fmt.Println(v1)
            v3 -= (v3 * (v3 % ((2147483648) | 1)))
        }
        for fi80 := 0; fi80 < 9; fi80++ {
            v81 := []int{241}
            var v82 string = ("\n\t" + "\n\t")
            _ = fi80
            _ = v81
            _ = v82
        }
        _ = v65
    }
    var v83 int16 = 976
    var v84 int8 = -110
    delete(v2, "hello")
    var v85 string = "hello"
    var v86 uint64 = 4294967295
    if true {
        for fi87 := 0; fi87 < 18; fi87++ {
            var v88 uint32 = (v3 >> ((1) & 63))
            _ = fi87
            _ = v88
        }
    } else {
        var v89 string = (v85 + "hello")
        v90 := GenS6{F0: ((v3 - (v3 * v3)) <= v3)}
        for fi91 := 0; fi91 < 3; fi91++ {
            v92 := GenS12{F0: "\n\t", F1: (v3 << ((819) & 63)), F2: ((v85 + v85) + v89)}
            var v93 uint64 = (^(v86))
            v94 := v2["hello"]
            _ = fi91
            _ = v92
            _ = v93
            _ = v94
        }
        _ = v89
        _ = v90
    }
    for fi95 := 0; fi95 < 3; fi95++ {
        for fi96 := 0; fi96 < 9; fi96++ {
            fmt.Println(fi95)
            v86 = 1
            switch fi96 {
            case 128:
                v83 = 11
            default:
                v97 := []int{(fi95 / ((fi95) | 1)), -1, fi96}
                _ = v97
            }
            fmt.Println(v1)
            _ = fi96
        }
        if false {
            switch (v85 + v85) {
            case "":
                v84 = v84
            }
            for fi98 := 0; fi98 < 18; fi98++ {
                fmt.Println(fi95, v1)
                v99 := GenS11{F0: v3, F1: (fi95 & fi98), F2: v1}
                v100 := []int{v1}
                _ = fi98
                _ = v100
                _ = v99
            }
        } else {
            var v101 float64 = -2.25
            v2[(v85 + "foo bar")] = v3
            switch v85 {
            case "\n\t":
                v102 := GenS3{F0: 13}
                _ = v102
            case "a":
                for fi103 := 0; fi103 < 18; fi103++ {
                    v104 := []int{v1, fi95, (fi95 << ((734) & 63))}
                    if (!((1 == -1) || true)) {
                        v85 = ("hello" + ("" + (((v85 + v85) + v85) + ((v85 + (("hello" + v85) + v85)) + "\n\t"))))
                        v105 := GenS13{F0: -561, F1: 255, F2: 1}
                        v85 = ("" + "foo bar")
                        _ = v105
                    } else {
                        fmt.Println(v101, fi95, v86)
                        v106 := GenS1{F0: (^(v3))}
                        v107 := GenS1{F0: 2147483648}
                        _ = v106
                        _ = v107
                    }
                    var v108 int16 = -1
                    _ = fi103
                    _ = v104
                    _ = v108
                }
            default:
                v109 := v2[(("foo bar" + "foo bar") + v85)]
                _ = v109
            }
            _ = v101
        }
        var v110 int32 = 453
        _ = fi95
        _ = v110
    }
    switch (v1 ^ (^(v1))) {
    case -1:
        for fi111 := 0; fi111 < 17; fi111++ {
            if true {
                switch v1 {
                case 1:
                    v112 := GenS1{F0: v3}
                    _ = v112
                case -9223372036854775808:
                    v113 := GenS11{F0: (v3 - v3), F1: fi111, F2: (v1 % (((fi111 / ((v1) | 1))) | 1))}
                    _ = v113
                case -2147483648:
                    for fi114 := 0; fi114 < 3; fi114++ {
                        _ = len(v2)
                        switch v85 {
                        case "":
                            var v115 string = ""
                            v84 -= 126
                            _ = v115
                        default:
                            switch (v85 + ("\n\t" + ((v85 + "foo bar") + "a"))) {
                            case "hello":
                                v116 := GenS4{F0: v85, F1: 1}
                                _ = v116
                            default:
                                var v117 int8 = int8(v3)
                                _ = v117
                            }
                        }
                        if true {
                            v118 := GenS1{F0: 152}
                            if false {
                                var v119 uint16 = 0
                                v120 := GenS13{F0: 2147483647, F1: 65535, F2: -128}
                                v121 := GenS2{F0: (127 <= (fi111 + v1)), F1: v119, F2: (v84 % ((127) | 1))}
                                _ = v119
                                _ = v120
                                _ = v121
                            }
                            fmt.Println(v84, v85, v1)
                            var v122 string = ("hello" + (v85 + v85))
                            _ = v118
                            _ = v122
                        } else {
                            v3 = v3
                        }
                        _ = fi114
                    }
                default:
                    if (((-2.25 * -2.25) <= 0.1) || (!true)) {
                        for fi123 := 0; fi123 < 13; fi123++ {
                            if (-9223372036854775808 > -9223372036854775808) {
                                v85 = v85
                            }
                            v85 = ("\n\t" + (v85 + v85))
                            switch fi111 {
                            case -979:
                                v85 = ("foo bar" + v85)
                            case 128:
                                v1 = fi123
                            default:
                                v124 := GenS4{F0: v85, F1: 813}
                                _ = v124
                            }
                            _ = fi123
                        }
                        switch (^(fi111)) {
                        case -1:
                            v125 := GenS11{F0: (^(v3)), F1: fi111, F2: (v1 ^ -48)}
                            _ = v125
                        case 2147483647:
                            v126 := []int{}
                            _ = v126
                        }
                    } else {
                        v2[v85] = v3
                    }
                }
                fmt.Println(v84)
            }
            var v127 float64 = 1.5
            v128 := GenS4{F0: v85, F1: 392}
            _ = fi111
            _ = v127
            _ = v128
        }
    case -297:
        var v129 int32 = int32(v86)
        _ = v129
    case -62:
        for fi130 := 0; fi130 < 18; fi130++ {
            var v131 int32 = -596
            _ = fi130
            _ = v131
        }
    default:
        var v132 uint = 683
        _ = v132
    }
    if false {
        if (!true) {
            v84 -= (v84 & (v84 * -1))
            v133 := v2[(v85 + v85)]
            _ = v133
        }
    } else {
        delete(v2, "\n\t")
        var v134 float32 = float32(0.1)
        if true {
            v135 := []int{(v1 & v1), v1, int(v3)}
            fmt.Println(v85, v85, v83)
            if false {
                var v136 int8 = (v84 >> ((18446744073709551615) & 63))
                v137 := v2[v85]
                var v138 uint = 591
                _ = v136
                _ = v137
                _ = v138
            } else {
                v86 = v86
                for fi139 := 0; fi139 < 10; fi139++ {
                    switch v85 {
                    case "foo bar":
                        for fi140 := 0; fi140 < 6; fi140++ {
                            var v141 int = fi140
                            _ = fi140
                            _ = v141
                        }
                    case "a":
                        var v142 int64 = 621
                        _ = v142
                    }
                    _ = fi139
                }
                v1 += 2147483647
            }
            _ = v135
        } else {
            v143 := GenS2{F0: false, F1: 32767, F2: (v84 + (v84 >> ((18446744073709551615) & 63)))}
            _ = v143
        }
        _ = v134
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
