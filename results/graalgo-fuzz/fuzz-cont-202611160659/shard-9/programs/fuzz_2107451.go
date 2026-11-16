package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 bool; F1 bool; F2 int }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS9 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 int16, p1 uint32, p2 bool) int8 {
    var v4 int32 = 1
    fmt.Println(v4, v4)
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 124
}
func genFn1(p0 uint8, p1 int32, p2 uint16) uint16 {
    var v5 int16 = 42
    var v6 uint8 = (p0 / ((127) | 1))
    v7 := GenS1{F0: p0}
    switch "foo bar" {
    case "foo bar":
        switch ("foo bar" + "a") {
        case "foo bar":
            v8 := GenS1{F0: v6}
            _ = v8
        case "\n\t":
            v9 := GenS2{F0: 71, F1: 128}
            _ = v9
        default:
            var v10 int16 = -32768
            _ = v10
        }
    case "hello":
        var v11 float64 = 3.14159
        _ = v11
    case "a":
        switch -9223372036854775808 {
        case 42:
            v12 := []int{4294967295, 154}
            _ = v12
        case 896:
            var v13 uint = 141
            _ = v13
        case -9223372036854775808:
            var v14 uint8 = (p0 / ((p0) | 1))
            _ = v14
        }
    default:
        if (-91 >= -90) {
            v15 := []int{2147483647}
            _ = v15
        } else {
            v16 := GenS1{F0: v6}
            p1 = int32(v5)
            _ = v16
        }
    }
    fmt.Println(p2, v6, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (p2 >> ((2) & 63))
}

func main() {
    var v1 float32 = float32(((3.14159 - 0.1) * ((((-2.25 / float64(float32(3.14159))) + 0.1) / 0.1) - (((0.1 + 1e10) - 1e10) + -2.25))))
    v2 := []string{}
    var v3 int16 = 0
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 bool = true
        _ = fi17
        _ = v18
    }
    switch 299 {
    case 1:
        fmt.Println(v3, v3)
    default:
        var v19 uint32 = 42
        _ = v19
    }
    v3 -= v3
    var v20 int64 = -158
    if true {
        v1 = v1
        for fi21 := 0; fi21 < 20; fi21++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi21
        }
    } else {
        if ((!false) || false) {
            if false {
                v20 = v20
                if true {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    v3 += 32767
                } else {
                    switch 25 {
                    case -713:
                        var v22 uint8 = 127
                        _ = v22
                    case 127:
                        if (true && ((v3 << ((18446744073709551615) & 63)) == v3)) {
                            v3 = (v3 + (v3 % ((v3) | 1)))
                            for fi23 := 0; fi23 < 10; fi23++ {
                                for fi24 := 0; fi24 < 1; fi24++ {
                                    if (!false) {
                                        if (70 > 128) {
                                            var v25 uint = 0
                                            if false {
                                                var v26 uint64 = 42
                                                _ = v26
                                            } else {
                                                v27 := GenS3{F0: 128}
                                                v3 = v3
                                                v20 = 1
                                                fmt.Println(fi23)
                                                v2 = append(v2, "hello")
                                                v3 += (v3 - (v3 ^ v3))
                                                if len(v2) > 0 {
                                                    v2[0] = (("hello" + "hello") + ("" + ((((("a" + "") + "") + "foo bar") + "foo bar") + "hello")))
                                                }
                                                fmt.Println(fi23)
                                                _ = v27
                                            }
                                            var v28 int16 = v3
                                            _ = v25
                                            _ = v28
                                        }
                                        var v29 int8 = 127
                                        v30 := GenS3{F0: 42}
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v31 := GenS1{F0: 255}
                                        if true {
                                            switch (fi24 / ((0) | 1)) {
                                            case -62:
                                                switch int(v3) {
                                                case 2147483647:
                                                    var v32 int16 = (v3 - (v3 % ((v3) | 1)))
                                                    _ = v32
                                                default:
                                                    v3 = (v3 >> ((2) & 63))
                                                }
                                            case 9223372036854775807:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v33 := GenS4{F0: 1, F1: 2147483648}
                                                _ = v33
                                            case 32768:
                                                v1 += ((v1 - 1e10) / v1)
                                            }
                                        } else {
                                            var v34 int8 = 127
                                            if false {
                                                fmt.Println(v34)
                                                v35 := GenS5{F0: (genFn1(13, 603, 65535) == 42), F1: false, F2: (-(fi23))}
                                                _ = v35
                                            }
                                            _ = v34
                                        }
                                        var v36 uint16 = 943
                                        if false {
                                            if (v20 <= (v20 ^ v20)) {
                                                fmt.Println(v36, fi23)
                                            } else {
                                                switch "" {
                                                case "a":
                                                    if (4294967295 >= 0) {
                                                        v3 += v3
                                                        var v37 string = ""
                                                        _ = v37
                                                    }
                                                case "":
                                                    v38 := GenS3{F0: 107}
                                                    _ = v38
                                                case "\n\t":
                                                    v20 = (v20 / ((v20) | 1))
                                                default:
                                                    var v39 uint64 = 42
                                                    _ = v39
                                                }
                                                if true {
                                                    var v40 bool = false
                                                    fmt.Println(v3)
                                                    switch 32768 {
                                                    case 32768:
                                                        fmt.Println(v40)
                                                    default:
                                                        var v41 int = fi23
                                                        switch ("hello" + (((("hello" + "a") + "\n\t") + "") + "a")) {
                                                        case "hello":
                                                            switch 74 {
                                                            case 1:
                                                                v42 := GenS6{F0: 524, F1: -9223372036854775808}
                                                                _ = v42
                                                            case 128:
                                                                v40 = ((^(v36)) > genFn1(131, 0, (v36 ^ 255)))
                                                            case -918:
                                                                var v43 uint16 = (v36 + genFn1(0, 1, v36))
                                                                _ = v43
                                                            default:
                                                                var v44 float64 = (-2.25 / 0.1)
                                                                _ = v44
                                                            }
                                                        case "a":
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                        }
                                                        _ = v41
                                                    }
                                                    _ = v40
                                                } else {
                                                    if (-1 > 859) {
                                                        v20 -= v20
                                                        v2 = append(v2, "hello")
                                                        v1 = float32((1e10 / -2.25))
                                                    }
                                                    var v45 uint = 0
                                                    var v46 int = (fi24 + int(v36))
                                                    var v47 uint = (^(v45))
                                                    v2 = append(v2, (((("" + "foo bar") + "") + "hello") + "foo bar"))
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                }
                                            }
                                            var v48 int32 = 1
                                            _ = v48
                                        } else {
                                            fmt.Println(v1, v1)
                                            var v49 uint8 = 7
                                            v2 = append(v2, "")
                                            _ = v49
                                        }
                                        _ = v31
                                        _ = v36
                                    }
                                    var v50 uint32 = 0
                                    _ = fi24
                                    _ = v50
                                }
                                if false {
                                    _ = len(v2)
                                    var v51 int32 = -2147483648
                                    _ = v51
                                }
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                _ = fi23
                            }
                        }
                    case 128:
                        if true {
                            var v52 uint = 614
                            v3 = v3
                            fmt.Println(v20)
                            _ = v52
                        }
                    }
                }
                v53 := GenS3{F0: 255}
                _ = v53
            } else {
                _ = len(v2)
                if ((255 == 21) || true) {
                    v54 := GenS3{F0: 1}
                    var v55 uint = 600
                    _ = v54
                    _ = v55
                }
                v56 := GenS2{F0: -44, F1: 35}
                _ = v56
            }
            var v57 int32 = 815
            var v58 int32 = -437
            _ = v57
            _ = v58
        }
        v59 := GenS1{F0: 122}
        _ = v59
    }
    switch 32767 {
    case 32767:
        for fi60 := 0; fi60 < 5; fi60++ {
            v3 += v3
            _ = fi60
        }
    case -342:
        var v61 bool = true
        _ = v61
    default:
        var v62 int8 = genFn0((v3 | (-(v3))), 42, false)
        _ = v62
    }
    var v63 int64 = (v20 ^ v20)
    switch "hello" {
    case "":
        var v64 int8 = -122
        _ = v64
    default:
        v63 = 9223372036854775807
    }
    if true {
        switch 1 {
        case 889:
            v20 = v20
        case -2147483648:
            switch "" {
            case "a":
                if true {
                    for fi65 := 0; fi65 < 6; fi65++ {
                        v3 = v3
                        _ = fi65
                    }
                } else {
                    v1 = float32(((0.0 * -2.25) / ((1e10 * (0.0 + 0.1)) / 3.14159)))
                    v66 := GenS3{F0: 1}
                    for fi67 := 0; fi67 < 12; fi67++ {
                        if (int8(fi67) < genFn0(1000, 1, ((v63 * v20) != v20))) {
                            var v68 int16 = v3
                            var v69 uint16 = 65535
                            _ = v68
                            _ = v69
                        } else {
                            switch (fi67 >> ((18446744073709551615) & 63)) {
                            case -9223372036854775808:
                                v1 += v1
                            case 1:
                                var v70 uint64 = 18446744073709551615
                                _ = v70
                            case 447:
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        for fi71 := 0; fi71 < 18; fi71++ {
                                            if ((fi67 / (((-(fi71))) | 1)) == fi67) {
                                                v3 = (v3 % ((v3) | 1))
                                                fmt.Println(v3, v20)
                                                if (42 > 2147483647) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v72 uint8 = 0
                                                        v20 = (v63 * (v20 / (((v20 % ((0) | 1))) | 1)))
                                                        var v73 uint32 = 175
                                                        _ = v72
                                                        _ = v73
                                                    case "":
                                                        fmt.Println(v3, v20)
                                                    case "foo bar":
                                                        if (v3 < (v3 * v3)) {
                                                            var v74 uint = 353
                                                            _ = v74
                                                        } else {
                                                            var v75 uint32 = 65535
                                                            v76 := GenS5{F0: false, F1: (((0.0 - 1e10) >= ((0.1 + float64(v1)) - float64(((float32(-2.25) - (-2.25 * 0.0)) * v1)))) || true), F2: fi71}
                                                            v63 = 1
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v1 += float32(0.1)
                                                        case "\n\t":
                                                            if (!(!false)) {
                                                                var v77 uint32 = 4294967295
                                                                v78 := GenS4{F0: genFn1(1, -862, 375), F1: v77}
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                        default:
                                                            switch fi67 {
                                                            case 4294967295:
                                                                v2 = append(v2, "\n\t")
                                                            }
                                                        }
                                                        v79 := GenS4{F0: 0, F1: 0}
                                                        _ = v79
                                                    default:
                                                        var v80 int16 = v3
                                                        _ = v80
                                                    }
                                                }
                                            } else {
                                                v81 := GenS7{F0: true}
                                                var v82 uint16 = 32767
                                                _ = v81
                                                _ = v82
                                            }
                                            switch fi71 {
                                            case 128:
                                                var v83 string = "\n\t"
                                                _ = v83
                                            case 2147483647:
                                                if (genFn1(131, 0, 255) <= 65535) {
                                                    var v84 uint8 = 50
                                                    _ = v84
                                                }
                                            case 32767:
                                                var v85 uint16 = genFn1(1, -1, 65535)
                                                _ = v85
                                            }
                                            _ = fi71
                                        }
                                    case "a":
                                        var v86 float64 = 3.14159
                                        _ = v86
                                    }
                                } else {
                                    v20 = 905
                                    v1 -= 3.14159
                                }
                            default:
                                if false {
                                    v3 = (v3 & 1000)
                                    var v87 uint = 2
                                    _ = v87
                                } else {
                                    fmt.Println(v20, v1)
                                }
                            }
                            switch "" {
                            case "foo bar":
                                v88 := GenS4{F0: 990, F1: 0}
                                _ = v88
                            }
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                        }
                        if true {
                            var v89 float64 = 1.5
                            _ = v89
                        } else {
                            if true {
                                var v90 int32 = 2147483647
                                v90 = v90
                                v91 := GenS6{F0: 65535, F1: int64(v90)}
                                _ = v90
                                _ = v91
                            }
                            switch "" {
                            case "a":
                                v92 := GenS8{F0: 2, F1: 952, F2: 42}
                                _ = v92
                            default:
                                v93 := GenS5{F0: false, F1: (0 == 2), F2: 0}
                                _ = v93
                            }
                        }
                        _ = fi67
                    }
                    _ = v66
                }
            default:
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi94 := 0; fi94 < 9; fi94++ {
                        if (false || (32767 <= 32767)) {
                            switch "\n\t" {
                            case "a":
                                switch fi94 {
                                case -227:
                                    v2 = append(v2, "\n\t")
                                case -826:
                                    var v95 int64 = (-(v63))
                                    _ = v95
                                default:
                                    var v96 uint = 4294967295
                                    _ = v96
                                }
                            case "":
                                v63 = (v20 | v20)
                            default:
                                switch fi94 {
                                case -1:
                                    var v97 uint16 = 65535
                                    _ = v97
                                case 32767:
                                    if true {
                                        if true {
                                            var v98 float64 = -2.25
                                            v3 += (v3 & 42)
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v99 int = (fi94 << ((2) & 63))
                                                _ = v99
                                            case "foo bar":
                                                var v100 int8 = 42
                                                _ = v100
                                            case "hello":
                                                if (255 < 65535) {
                                                    var v101 uint64 = 42
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v102 int64 = 9223372036854775807
                                                    _ = v101
                                                    _ = v102
                                                }
                                            default:
                                                v63 += (^(v63))
                                            }
                                            _ = v98
                                        }
                                        for fi103 := 0; fi103 < 11; fi103++ {
                                            var v104 uint32 = 2147483648
                                            switch ("\n\t" + ((("a" + "") + "a") + "\n\t")) {
                                            case "a":
                                                var v105 float32 = 3.14159
                                                _ = v105
                                            default:
                                                v106 := GenS6{F0: v104, F1: int64(v3)}
                                                _ = v106
                                            }
                                            var v107 int16 = (v3 >> ((255) & 63))
                                            _ = fi103
                                            _ = v104
                                            _ = v107
                                        }
                                    } else {
                                        v3 = (v3 % ((v3) | 1))
                                    }
                                case 32768:
                                    v63 += (v20 - v20)
                                default:
                                    v108 := GenS1{F0: 169}
                                    _ = v108
                                }
                            }
                            var v109 int = (fi94 % (((fi94 & (fi94 & (^(fi94))))) | 1))
                            _ = v109
                        }
                        _ = fi94
                    }
                    if ("hello" <= "hello") {
                        var v110 int8 = 0
                        v20 += 42
                        _ = v110
                    } else {
                        for fi111 := 0; fi111 < 4; fi111++ {
                            v112 := GenS9{F0: 100, F1: genFn0((v3 - v3), 65535, (v3 >= v3)), F2: 1}
                            v20 -= (v63 ^ v20)
                            _ = fi111
                            _ = v112
                        }
                        if ("\n\t" > (("\n\t" + "") + ((("" + "a") + "hello") + "foo bar"))) {
                            v113 := GenS5{F0: (false || false), F1: true, F2: -639}
                            for fi114 := 0; fi114 < 8; fi114++ {
                                var v115 int64 = int64(fi114)
                                _ = fi114
                                _ = v115
                            }
                            if (2 <= 4294967295) {
                                v63 -= 42
                                v2 = append(v2, ("" + "foo bar"))
                                var v116 string = "hello"
                                _ = v116
                            } else {
                                var v117 int = 128
                                switch v117 {
                                case 128:
                                    _ = len(v2)
                                case -807:
                                    var v118 int64 = 0
                                    _ = v118
                                case 127:
                                    for fi119 := 0; fi119 < 18; fi119++ {
                                        v63 = (v63 / (((v20 >> ((4294967295) & 63))) | 1))
                                        _ = fi119
                                    }
                                default:
                                    v1 -= (-2.25 + v1)
                                }
                                _ = v117
                            }
                            _ = v113
                        } else {
                            var v120 uint16 = 255
                            v120 -= (v120 >> ((65535) & 63))
                            _ = v120
                        }
                        fmt.Println(v63)
                    }
                } else {
                    v121 := GenS8{F0: 677, F1: 4294967295, F2: 128}
                    v122 := GenS2{F0: 0, F1: 128}
                    v63 = (-(v20))
                    _ = v121
                    _ = v122
                }
            }
        case 4294967295:
            if false {
                var v123 int32 = 0
                var v124 string = "a"
                _ = v123
                _ = v124
            }
        }
        var v125 uint64 = 728
        v20 -= v20
        _ = v125
    }
    v2 = append(v2, "")
    var v126 int64 = v63
    var v127 uint64 = 18446744073709551615
    v1 = 0.0
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v3, v20, v20)
    for fi128 := 0; fi128 < 14; fi128++ {
        var v129 bool = true
        _ = fi128
        _ = v129
    }
    var v130 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v130:", v130)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v126
    _ = v127
    _ = v130
    _ = v2
    _ = v20
    _ = v3
    _ = v63
}
