package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 int8; F1 int16; F2 int8 }
type GenS10 struct { F0 int32; F1 string; F2 int64 }
type GenS11 struct { F0 uint32; F1 uint }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS14 struct { F0 int16; F1 string }
type GenS15 struct { F0 int32; F1 uint64 }
type GenS16 struct { F0 int; F1 uint32 }
type GenS17 struct { F0 uint; F1 int; F2 int32 }
type GenS18 struct { F0 uint32; F1 string; F2 bool }
type GenS19 struct { F0 uint64; F1 int16 }
type GenS20 struct { F0 uint64; F1 int16; F2 int32 }
type GenS21 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int8) int32 {
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 uint = 255
            switch fi6 {
            case 128:
                var v9 uint = v8
                _ = v9
            case -2147483648:
                v10 := []int{(fi6 << (((v8 % ((797) | 1))) & 63)), (fi6 % ((fi6) | 1))}
                _ = v10
            default:
                var v11 int32 = 703
                _ = v11
            }
            if false {
                v12 := GenS1{F0: 1, F1: (p1 - (p1 >> (((v8 * (v8 ^ 290))) & 63)))}
                var v13 uint8 = 193
                _ = v12
                _ = v13
            } else {
                if false {
                    v14 := []int{(fi7 & 128)}
                    var v15 uint32 = 65535
                    if false {
                        if len(v14) > 0 {
                            v14[0] = 4294967295
                        }
                        if true {
                            var v16 int16 = 1
                            _ = v16
                        } else {
                            fmt.Println(fi7, fi7, v8)
                        }
                        v17 := GenS1{F0: -1, F1: p1}
                        if len(v14) > 0 {
                            v14[0] = fi6
                        }
                        var v18 float64 = -2.25
                        _ = v17
                        _ = v18
                    }
                    _ = v14
                    _ = v15
                } else {
                    var v19 string = ("a" + "foo bar")
                    var v20 uint32 = 2
                    switch (("\n\t" + (("hello" + v19) + ("" + "foo bar"))) + v19) {
                    case "foo bar":
                        var v21 int64 = -9223372036854775808
                        _ = v21
                    default:
                        v22 := []int{(fi7 - 32767), -708}
                        _ = v22
                    }
                    _ = v19
                    _ = v20
                }
                var v23 uint64 = 4294967295
                _ = v23
            }
            _ = fi7
            _ = v8
        }
        if ((!(-1 <= 991)) && true) {
            var v24 float32 = 1e10
            var v25 int64 = 623
            _ = v24
            _ = v25
        } else {
            if true {
                var v26 int = (fi6 >> ((4294967295) & 63))
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 int32 = 2147483647
                    if ((!false) && true) {
                        v29 := GenS1{F0: (v28 & 42), F1: 127}
                        _ = v29
                    } else {
                        v30 := GenS2{F0: -32768, F1: (^(p0))}
                        _ = v30
                    }
                    fmt.Println(p0, fi27)
                    switch "a" {
                    case "\n\t":
                        switch (fi27 - 4294967295) {
                        case -1:
                            fmt.Println(fi6, p0)
                        case -233:
                            v31 := GenS2{F0: 1000, F1: (p0 % ((2) | 1))}
                            _ = v31
                        default:
                            var v32 int = (fi6 & 4294967295)
                            _ = v32
                        }
                    default:
                        fmt.Println(fi27)
                    }
                    p0 = (p0 & 320)
                    _ = fi27
                    _ = v28
                }
                _ = v26
            } else {
                var v33 bool = (!false)
                _ = v33
            }
            var v34 float64 = 1.5
            _ = v34
        }
        _ = fi6
    }
    var v35 uint32 = 156
    fmt.Println(p1, p1, p1)
    v36 := []int{}
    if false {
        switch 128 {
        case 42:
            for fi37 := 0; fi37 < 12; fi37++ {
                if true {
                    v38 := GenS3{F0: (fi37 - 42), F1: uint8(v35), F2: 865}
                    for fi39 := 0; fi39 < 2; fi39++ {
                        p0 -= (p0 & 305)
                        switch (fi39 >> ((4294967295) & 63)) {
                        case 32768:
                            if false {
                                if true {
                                    switch fi37 {
                                    case 9223372036854775807:
                                        var v40 int16 = 1
                                        _ = v40
                                    case 42:
                                        var v41 uint16 = p0
                                        _ = v41
                                    case 4294967295:
                                        v42 := GenS4{F0: 42}
                                        _ = v42
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        var v43 int32 = 258
                                        _ = v43
                                    case "foo bar":
                                        if false {
                                            p0 += 535
                                            if len(v36) >= 2 {
                                                v36 = v36[1:len(v36)]
                                            }
                                        }
                                    case "a":
                                        switch "\n\t" {
                                        case "a":
                                            if false {
                                                var v44 uint8 = 2
                                                _ = v44
                                            } else {
                                                if true {
                                                    var v45 int16 = -1
                                                    _ = v45
                                                }
                                                var v46 bool = (p0 != (p0 / ((p0) | 1)))
                                                p0 -= p0
                                                v35 = 65535
                                                _ = v46
                                            }
                                            _ = len(v36)
                                            var v47 int8 = p1
                                            var v48 int = (fi39 - 401)
                                            v36 = append(v36, (fi37 + 0))
                                            if false {
                                                var v49 float32 = (3.14159 * -2.25)
                                                _ = v49
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                    var v50 string = "\n\t"
                                    _ = v50
                                }
                            } else {
                                _ = len(v36)
                                var v51 uint8 = uint8(p0)
                                _ = v51
                            }
                        }
                        p1 = (p1 << ((557) & 63))
                        var v52 uint16 = p0
                        _ = fi39
                        _ = v52
                    }
                    v53 := GenS5{F0: 42, F1: (fi37 * fi37)}
                    _ = v38
                    _ = v53
                }
                if false {
                    p0 += (p0 & p0)
                    for fi54 := 0; fi54 < 12; fi54++ {
                        if ((2 <= 65535) || true) {
                            switch "a" {
                            case "":
                                _ = len(v36)
                                var v55 uint16 = (p0 ^ p0)
                                var v56 int64 = 0
                                _ = v55
                                _ = v56
                            default:
                                var v57 uint = 921
                                _ = v57
                            }
                        }
                        v35 -= v35
                        _ = fi54
                    }
                }
                p0 = (p0 | p0)
                _ = fi37
            }
        case 0:
            fmt.Println(p1, p0)
        }
        var v58 int8 = p1
        _ = v58
    }
    _ = p0
    _ = p1
    _ = v35
    _ = v36
    return -183
}

func main() {
    var v1 uint64 = 401
    v2 := []string{("\n\t" + "foo bar")}
    var v3 int = 128
    v4 := []string{"", ((("" + ("" + "")) + "foo bar") + "\n\t"), "a"}
    v5 := make(map[uint32]int)
    var v59 int32 = 2147483647
    if false {
        var v60 float32 = 1e10
        _ = v60
    } else {
        v1 = v1
        for fi61 := 0; fi61 < 11; fi61++ {
            var v62 bool = (2147483648 != 942)
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = fi61
            _ = v62
        }
    }
    v59 = 0
    var v63 float64 = 0.0
    for fi64 := 0; fi64 < 4; fi64++ {
        fmt.Println(v1, fi64)
        _ = fi64
    }
    var v65 uint8 = 42
    var v66 float32 = (0.1 - (-2.25 / float32((v63 + v63))))
    v67 := GenS1{F0: 1, F1: 100}
    v68 := GenS4{F0: -9223372036854775808}
    for fi69 := 0; fi69 < 10; fi69++ {
        for fi70 := 0; fi70 < 18; fi70++ {
            if true {
                v66 += -2.25
                var v71 float64 = float64((v66 / v66))
                switch (fi69 % ((v3) | 1)) {
                case -160:
                    fmt.Println(fi69, v1)
                }
                v72 := GenS1{F0: v59, F1: -1}
                v3 = v3
                _ = v71
                _ = v72
            }
            v3 = (v3 | int(v1))
            switch 128 {
            case 127:
                v73 := GenS5{F0: v59, F1: (fi70 ^ 127)}
                _ = v73
            case 1:
                if false {
                    var v74 uint16 = 65535
                    switch 32768 {
                    case 9223372036854775807:
                        v75 := GenS4{F0: 9223372036854775807}
                        _ = v75
                    case -631:
                        fmt.Println(fi69)
                    case 128:
                        var v76 uint16 = 32767
                        _ = v76
                    default:
                        v5[487] = (fi69 % (((-(v3))) | 1))
                    }
                    _ = v74
                } else {
                    if (!false) {
                        v77 := GenS6{F0: 32767}
                        switch fi70 {
                        case 4294967295:
                            if (v3 == (fi69 >> ((42) & 63))) {
                                fmt.Println(v65, fi69)
                                if true {
                                    switch fi70 {
                                    case 60:
                                        v2 = append(v2, "\n\t")
                                    case -1:
                                        var v78 bool = (v1 <= (v1 & v1))
                                        _ = v78
                                    default:
                                        switch "hello" {
                                        case "":
                                            v1 += (v1 - (v1 * v1))
                                        case "\n\t":
                                            v79 := GenS7{F0: (v65 | (v65 + 42)), F1: 2}
                                            _ = v79
                                        }
                                    }
                                    v4 = append(v4, "")
                                    fmt.Println(fi69)
                                    v63 = v63
                                    v63 = v63
                                } else {
                                    if (-2147483648 <= (fi69 & -210)) {
                                        v80 := GenS8{F0: -908, F1: 0}
                                        v59 = genFn0(255, 100)
                                        var v81 float64 = (((((1.5 * (1e10 - (v63 * (1.5 - v63)))) * ((v63 / float64(3.14159)) / 3.14159)) - (v63 - 1.5)) - ((v63 + v63) / v63)) - (1.5 * v63))
                                        var v82 uint = 408
                                        fmt.Println(v65, v59)
                                        v66 += ((v66 * float32(v63)) / ((v66 - v66) + (v66 - float32(((v81 / 1e10) - (v63 * v81))))))
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    } else {
                                        if false {
                                            if (42 != 127) {
                                                if (127 < 127) {
                                                    _ = len(v5)
                                                } else {
                                                    v66 = 0.1
                                                    v5[42] = (fi69 * v3)
                                                    if (v63 <= ((((v63 / (v63 - v63)) - v63) / (v63 / v63)) + (1.5 / v63))) {
                                                        v83 := GenS9{F0: 87, F1: 42, F2: 0}
                                                        fmt.Println(v65)
                                                        _ = v83
                                                    }
                                                }
                                                if ((v59 * genFn0(1, 42)) >= -86) {
                                                    if false {
                                                        var v84 uint8 = v65
                                                        var v85 uint16 = 32767
                                                        var v86 string = "\n\t"
                                                        _ = v84
                                                        _ = v85
                                                        _ = v86
                                                    } else {
                                                        v59 = v59
                                                        if true {
                                                            switch "a" {
                                                            case "foo bar":
                                                                v87 := GenS10{F0: v59, F1: ((("foo bar" + ("hello" + ("foo bar" + "foo bar"))) + "\n\t") + "a"), F2: -9223372036854775808}
                                                                _ = v87
                                                            case "\n\t":
                                                                v65 = (v65 / ((v65) | 1))
                                                            case "":
                                                                fmt.Println(v63, fi69, v65)
                                                            default:
                                                                var v88 float64 = 1e10
                                                                var v89 int = -2147483648
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            var v90 uint16 = 891
                                                            v91 := GenS8{F0: 42, F1: 888}
                                                            _ = v90
                                                            _ = v91
                                                        } else {
                                                            var v92 uint = 0
                                                            switch 842 {
                                                            case 2147483647:
                                                                v66 += (float32((1e10 / v63)) + v66)
                                                            default:
                                                                if (35 != 1) {
                                                                    var v93 int = fi70
                                                                    if false {
                                                                        var v94 string = "\n\t"
                                                                        _ = v94
                                                                    }
                                                                    var v95 int8 = -128
                                                                    _ = v93
                                                                    _ = v95
                                                                }
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = ("hello" + "hello")
                                                            }
                                                            _ = v92
                                                        }
                                                        switch 42 {
                                                        case 32767:
                                                            fmt.Println(v65)
                                                        case 32768:
                                                            v96 := GenS8{F0: -833, F1: 2147483648}
                                                            _ = v96
                                                        case 1:
                                                            switch v3 {
                                                            case 428:
                                                                v65 = 255
                                                            case 32768:
                                                                v97 := GenS7{F0: 127, F1: (^(v65))}
                                                                _ = v97
                                                            case 9223372036854775807:
                                                                v98 := GenS4{F0: 0}
                                                                _ = v98
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        if false {
                                                            fmt.Println(v59, v3, v59)
                                                            var v99 int = 127
                                                            _ = v99
                                                        }
                                                    } else {
                                                        var v100 int = 0
                                                        _ = v100
                                                    }
                                                    v101 := GenS10{F0: v59, F1: "foo bar", F2: 0}
                                                    _ = v101
                                                } else {
                                                    fmt.Println(v63, fi69)
                                                    if ((-2.25 / 1e10) > v63) {
                                                        v63 = 1.5
                                                        fmt.Println(v59)
                                                        if (uint16(v1) >= 32767) {
                                                            fmt.Println(fi69, fi70)
                                                            v63 = (v63 * 1e10)
                                                        }
                                                    } else {
                                                        v2 = append(v2, (("hello" + ("\n\t" + ("\n\t" + (("foo bar" + "hello") + "a")))) + "hello"))
                                                        v102 := GenS11{F0: 0, F1: 611}
                                                        switch v3 {
                                                        case 317:
                                                            if ((^(v1)) > uint64(v65)) {
                                                                v65 = 128
                                                                v3 = int(v59)
                                                                if true {
                                                                    switch (fi69 | v3) {
                                                                    case -9223372036854775808:
                                                                        var v103 int16 = 734
                                                                        _ = v103
                                                                    }
                                                                    v63 = v63
                                                                }
                                                                if len(v2) > 0 {
                                                                    v2[0] = ""
                                                                }
                                                                var v104 uint16 = 2
                                                                v105 := GenS7{F0: (v65 << ((2) & 63)), F1: (v65 ^ 127)}
                                                                v65 -= 2
                                                                if (!(!(-32768 < -90))) {
                                                                    v3 = (v3 & fi70)
                                                                } else {
                                                                    if true {
                                                                        var v106 uint = 0
                                                                        if (!true) {
                                                                            v107 := v5[216]
                                                                            var v108 int32 = int32(v1)
                                                                            _ = v107
                                                                            _ = v108
                                                                        }
                                                                        _ = v106
                                                                    }
                                                                    var v109 uint = uint(v1)
                                                                    var v110 string = "hello"
                                                                    _ = v109
                                                                    _ = v110
                                                                }
                                                                var v111 uint32 = 396
                                                                if (!(true && true)) {
                                                                    v112 := GenS9{F0: 100, F1: -1, F2: -1}
                                                                    var v113 uint64 = (v1 % ((v1) | 1))
                                                                    v111 += (^(v111))
                                                                    _ = v112
                                                                    _ = v113
                                                                }
                                                                _ = v104
                                                                _ = v105
                                                                _ = v111
                                                            }
                                                        default:
                                                            fmt.Println(v66, v65)
                                                        }
                                                        if (!true) {
                                                            var v114 int64 = 1
                                                            _ = v114
                                                        }
                                                        _ = len(v2)
                                                        _ = v102
                                                    }
                                                    if false {
                                                        fmt.Println(v66, v65, fi70)
                                                    }
                                                    fmt.Println(fi69, v66)
                                                    fmt.Println(fi70, v59, v3)
                                                    v115 := GenS7{F0: (v65 % ((v65) | 1)), F1: 128}
                                                    _ = v115
                                                }
                                                fmt.Println(v63)
                                                v4 = append(v4, "\n\t")
                                                v63 = (((1e10 - v63) + 0.1) + 0.0)
                                            } else {
                                                switch -1 {
                                                case 128:
                                                    if false {
                                                        if false {
                                                            if false {
                                                                if (!true) {
                                                                    v63 = v63
                                                                    var v116 uint8 = (v65 + 1)
                                                                    var v117 int64 = -714
                                                                    _ = len(v5)
                                                                    _ = v116
                                                                    _ = v117
                                                                } else {
                                                                    v118 := v5[646]
                                                                    if ((100 <= 0) || (!true)) {
                                                                        var v119 uint16 = 255
                                                                        _ = v119
                                                                    } else {
                                                                        v2 = append(v2, "a")
                                                                        if false {
                                                                            v118 -= (v3 | (^(fi69)))
                                                                        } else {
                                                                            v120 := GenS12{F0: 1}
                                                                            v121 := GenS3{F0: (fi69 / ((fi70) | 1)), F1: v65, F2: v1}
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ("a" + "")
                                                                            }
                                                                            _ = v120
                                                                            _ = v121
                                                                        }
                                                                        var v122 float64 = v63
                                                                        fmt.Println(v3, v118, v1)
                                                                        switch "" {
                                                                        case "":
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            default:
                                                                                var v123 float64 = v122
                                                                                _ = v123
                                                                            }
                                                                        case "\n\t":
                                                                            v118 -= (fi69 ^ -630)
                                                                        default:
                                                                            var v124 uint32 = 65535
                                                                            _ = v124
                                                                        }
                                                                        _ = v122
                                                                    }
                                                                    v125 := GenS11{F0: 2, F1: 65535}
                                                                    _ = v118
                                                                    _ = v125
                                                                }
                                                                fmt.Println(fi69, v65, v66)
                                                                var v126 int = fi70
                                                                _ = v126
                                                            }
                                                            v59 -= v59
                                                            v65 += 2
                                                        } else {
                                                            switch v3 {
                                                            case 1:
                                                                var v127 int16 = 32767
                                                                _ = v127
                                                            default:
                                                                v65 += uint8(v1)
                                                                var v128 uint16 = 65535
                                                                v3 = (fi69 * v3)
                                                                _ = v128
                                                            }
                                                            if (true || false) {
                                                                _ = len(v2)
                                                                if ((1.5 + ((0.0 / (-2.25 * (-2.25 / (v63 - v63)))) - v63)) >= v63) {
                                                                    var v129 bool = true
                                                                    var v130 int8 = 0
                                                                    var v131 int = fi69
                                                                    _ = v129
                                                                    _ = v130
                                                                    _ = v131
                                                                } else {
                                                                    var v132 float32 = v66
                                                                    _ = v132
                                                                }
                                                            } else {
                                                                var v133 uint64 = (v1 / (((v1 & v1)) | 1))
                                                                _ = v133
                                                            }
                                                            fmt.Println(v66, v59, v59)
                                                        }
                                                    }
                                                    v65 -= (v65 * (^(v65)))
                                                case -438:
                                                    if true {
                                                        switch fi70 {
                                                        case -9223372036854775808:
                                                            _ = len(v4)
                                                            fmt.Println(v3, v63, v65)
                                                            if (!(!("a" == ("\n\t" + "foo bar")))) {
                                                                v134 := GenS11{F0: 568, F1: 65535}
                                                                _ = v134
                                                            } else {
                                                                if false {
                                                                    v3 -= v3
                                                                    fmt.Println(v63, v66, v63)
                                                                    v135 := GenS10{F0: 1, F1: ("" + "\n\t"), F2: 0}
                                                                    _ = v135
                                                                } else {
                                                                    if ((-43 > 166) || false) {
                                                                        var v136 uint64 = 4294967295
                                                                        v137 := GenS13{F0: uint32(v65), F1: 685, F2: 334}
                                                                        var v138 uint32 = 42
                                                                        _ = v136
                                                                        _ = v137
                                                                        _ = v138
                                                                    } else {
                                                                        v66 -= (v66 + -2.25)
                                                                        fmt.Println(v59, fi69, v1)
                                                                    }
                                                                    if false {
                                                                        var v139 uint64 = v1
                                                                        _ = v139
                                                                    } else {
                                                                        v5[1] = fi69
                                                                    }
                                                                    fmt.Println(v63)
                                                                }
                                                                v59 += v59
                                                            }
                                                        case 32767:
                                                            fmt.Println(v66, v65, fi69)
                                                        case 77:
                                                            if false {
                                                                v1 -= (v1 ^ (^(v1)))
                                                                v140 := GenS4{F0: -146}
                                                                _ = v140
                                                            }
                                                        }
                                                        v141 := v5[4294967295]
                                                        v1 += 1
                                                        _ = v141
                                                    }
                                                default:
                                                    var v142 uint8 = 2
                                                    _ = v142
                                                }
                                                if ((v3 & (fi69 + (fi69 - v3))) > (-(fi70))) {
                                                    var v143 float64 = (float64(float32(v63)) * float64(v66))
                                                    v59 = 2147483647
                                                    var v144 int32 = (v59 & (v59 % ((v59) | 1)))
                                                    _ = v143
                                                    _ = v144
                                                } else {
                                                    _ = len(v4)
                                                    var v145 float32 = (((v66 / v66) / v66) * 3.14159)
                                                    _ = v145
                                                }
                                            }
                                            switch (fi70 - -2147483648) {
                                            case -1:
                                                switch ("foo bar" + "foo bar") {
                                                case "hello":
                                                    v146 := GenS14{F0: -1, F1: "\n\t"}
                                                    _ = v146
                                                }
                                            case 32768:
                                                _ = len(v5)
                                            }
                                        } else {
                                            var v147 uint = 18446744073709551615
                                            v63 = v63
                                            var v148 int64 = 184
                                            _ = v147
                                            _ = v148
                                        }
                                        _ = len(v4)
                                    }
                                    v66 = (v66 - v66)
                                }
                                var v149 uint32 = 2
                                _ = v149
                            }
                        }
                        v150 := GenS3{F0: 9223372036854775807, F1: (v65 << ((632) & 63)), F2: 640}
                        _ = v150
                        _ = v77
                    }
                    fmt.Println(v59, v66, v66)
                }
            case 32767:
                if (true && ((65535 >= 65535) && true)) {
                    if (v65 > 42) {
                        var v151 int32 = (v59 * v59)
                        _ = v151
                    } else {
                        var v152 int32 = (v59 - (v59 & (v59 / (((v59 << ((65535) & 63))) | 1))))
                        v153 := GenS4{F0: 42}
                        _ = v152
                        _ = v153
                    }
                    var v154 int32 = (v59 % ((genFn0(721, -1)) | 1))
                    var v155 bool = (1 <= 0)
                    _ = v154
                    _ = v155
                } else {
                    v63 = ((0.1 * 1.5) + (v63 - v63))
                    var v156 int16 = int16(v59)
                    v157 := GenS14{F0: v156, F1: "\n\t"}
                    _ = v156
                    _ = v157
                }
            default:
                v1 = (v1 ^ v1)
            }
            _ = fi70
        }
        _ = len(v4)
        _ = fi69
    }
    fmt.Println(v3, v63)
    if true {
        for fi158 := 0; fi158 < 11; fi158++ {
            if (false || false) {
                for fi159 := 0; fi159 < 17; fi159++ {
                    _ = len(v5)
                    _ = fi159
                }
                v160 := GenS15{F0: v59, F1: v1}
                _ = v160
            } else {
                var v161 uint8 = v65
                v161 -= 2
                _ = v161
            }
            switch fi158 {
            case 32768:
                switch ((("a" + (("foo bar" + "hello") + "hello")) + ("foo bar" + ("foo bar" + ""))) + "\n\t") {
                case "a":
                    for fi162 := 0; fi162 < 19; fi162++ {
                        if (((0.0 * v63) * (3.14159 - v63)) <= v63) {
                            var v163 uint8 = 127
                            if (1000 < 483) {
                                _ = len(v5)
                                v164 := GenS16{F0: int(v59), F1: 42}
                                if true {
                                    var v165 int16 = 286
                                    if len(v4) > 0 {
                                        v4[0] = (("\n\t" + "") + "foo bar")
                                    }
                                    if (1 <= -128) {
                                        if true {
                                            switch "hello" {
                                            case "\n\t":
                                                if (v65 > v65) {
                                                    if true {
                                                        var v166 int = 32767
                                                        _ = v166
                                                    } else {
                                                        v165 -= v165
                                                    }
                                                    if (false || true) {
                                                        fmt.Println(v66, v65, v59)
                                                        fmt.Println(v165, v1, v66)
                                                    }
                                                    if ((("a" + "\n\t") == "\n\t") && (!false)) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    } else {
                                                        _ = len(v2)
                                                        v66 += 0.0
                                                    }
                                                }
                                            default:
                                                v167 := GenS17{F0: 993, F1: (^(fi162)), F2: (v59 << ((2) & 63))}
                                                _ = v167
                                            }
                                        } else {
                                            switch 127 {
                                            case 0:
                                                var v168 uint64 = (v1 | v1)
                                                _ = v168
                                            }
                                            if false {
                                                var v169 int = -951
                                                _ = len(v2)
                                                var v170 int = v169
                                                _ = v169
                                                _ = v170
                                            }
                                            var v171 int16 = v165
                                            _ = v171
                                        }
                                    } else {
                                        v172 := GenS16{F0: 4294967295, F1: 2147483648}
                                        if false {
                                            var v173 uint8 = (v163 % (((v163 & (^(v163)))) | 1))
                                            v65 -= (v173 % (((v65 & v173)) | 1))
                                            _ = v173
                                        }
                                        _ = v172
                                    }
                                    _ = v165
                                }
                                v163 = (v65 - (v163 * v65))
                                v66 = v66
                                v5[2147483648] = 127
                                switch (v3 << ((567) & 63)) {
                                case 32768:
                                    var v174 int64 = 42
                                    _ = v174
                                case 0:
                                    v175 := GenS18{F0: 0, F1: "a", F2: (0 != int64(v59))}
                                    _ = v175
                                case 435:
                                    v1 -= 1
                                default:
                                    var v176 int64 = 9223372036854775807
                                    _ = v176
                                }
                                _ = v164
                            } else {
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v63)
                                default:
                                    var v177 int = (v3 ^ 0)
                                    _ = v177
                                }
                                v178 := GenS19{F0: v1, F1: 0}
                                _ = v178
                            }
                            var v179 int8 = -35
                            if (true || (true && (!((18446744073709551615 < v1) || true)))) {
                                v65 += v163
                            }
                            v180 := GenS5{F0: 0, F1: (v3 << ((42) & 63))}
                            _ = v163
                            _ = v179
                            _ = v180
                        }
                        var v181 uint8 = 226
                        v65 = (v181 * (v181 - v181))
                        fmt.Println(fi162, v63, v66)
                        _ = fi162
                        _ = v181
                    }
                case "hello":
                    v66 += ((-2.25 - v66) * (((v66 + v66) - 1.5) + v66))
                default:
                    switch fi158 {
                    case 42:
                        v65 = v65
                    case 127:
                        if true {
                            v182 := GenS20{F0: v1, F1: 0, F2: (v59 | genFn0(45, -128))}
                            _ = v182
                        }
                    }
                }
            case 2147483647:
                for fi183 := 0; fi183 < 20; fi183++ {
                    var v184 uint64 = (v1 - v1)
                    var v185 uint16 = 0
                    var v186 float32 = 0.1
                    _ = fi183
                    _ = v184
                    _ = v185
                    _ = v186
                }
            case -44:
                for fi187 := 0; fi187 < 3; fi187++ {
                    v188 := GenS13{F0: 55, F1: 32767, F2: 1}
                    _ = fi187
                    _ = v188
                }
            }
            if false {
                v189 := GenS18{F0: 65535, F1: "hello", F2: (true && (2 == v1))}
                _ = v189
            }
            _ = fi158
        }
    } else {
        _ = len(v4)
        delete(v5, 2147483648)
        v190 := GenS21{F0: 128}
        _ = v190
    }
    var v191 bool = false
    for fi192 := 0; fi192 < 5; fi192++ {
        if v191 {
            v193 := GenS13{F0: 4294967295, F1: int16(v1), F2: 42}
            _ = v193
        } else {
            _ = len(v4)
            v1 = (v1 >> ((820) & 63))
            var v194 string = ("" + ("\n\t" + ("" + (("\n\t" + "\n\t") + ("foo bar" + "a")))))
            _ = v194
        }
        _ = fi192
    }
    fmt.Println(v66)
    v66 += v66
    var v195 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v191:", v191)
    fmt.Println("v195:", v195)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v191
    _ = v195
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v63
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
