package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int32; F1 int64; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int8; F2 int32 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 bool; F2 uint32 }
type GenS12 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS13 struct { F0 int16; F1 uint16; F2 int16 }
type GenS14 struct { F0 uint32; F1 uint8 }
type GenS15 struct { F0 int; F1 int64; F2 int64 }
type GenS16 struct { F0 uint16; F1 int64; F2 uint }
type GenS17 struct { F0 bool; F1 int8 }
type GenS18 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint; F1 int32; F2 uint8 }
type GenS21 struct { F0 int8; F1 uint32; F2 uint }
type GenS22 struct { F0 int; F1 int64 }
type GenS23 struct { F0 int }
type GenS24 struct { F0 int; F1 int16; F2 uint64 }
type GenS25 struct { F0 bool; F1 int8; F2 int8 }
type GenS26 struct { F0 uint32; F1 int64 }
type GenS27 struct { F0 int32; F1 uint }
type GenS28 struct { F0 int32 }
type GenS29 struct { F0 uint8; F1 string }
type GenS30 struct { F0 uint64 }
type GenS31 struct { F0 bool; F1 uint64 }
type GenS32 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS33 struct { F0 bool; F1 int16 }
type GenS34 struct { F0 bool }
type GenS35 struct { F0 int64 }
type GenS36 struct { F0 int32; F1 uint8; F2 int16 }
type GenS37 struct { F0 string; F1 uint32; F2 string }
type GenS38 struct { F0 uint; F1 uint }
type GenS39 struct { F0 uint16 }
type GenS40 struct { F0 uint; F1 uint8 }
type GenS41 struct { F0 int64; F1 uint64; F2 string }
type GenS42 struct { F0 int16 }
type GenS43 struct { F0 int8; F1 int32; F2 int }
type GenS44 struct { F0 int32 }
func genFn0(p0 uint8) uint {
    switch "\n\t" {
    case "hello":
        v7 := GenS1{F0: 1, F1: "\n\t"}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int16 = 42
                switch fi9 {
                case -143:
                    var v11 int64 = 636
                    _ = v11
                case -2147483648:
                    fmt.Println(fi9)
                case 1:
                    var v12 int8 = -128
                    fmt.Println(v10, p0)
                    _ = v12
                default:
                    if (32767 < 1) {
                        var v13 int64 = 9223372036854775807
                        if (-558 <= v10) {
                            var v14 uint64 = 152
                            _ = v14
                        } else {
                            v13 += v13
                        }
                        _ = v13
                    }
                }
                fmt.Println(fi8)
                _ = fi9
                _ = v10
            }
            switch fi8 {
            case 42:
                if (true && true) {
                    if (true && (4294967295 <= 65535)) {
                        switch "a" {
                        case "\n\t":
                            switch "" {
                            case "hello":
                                fmt.Println(p0, p0)
                            default:
                                if true {
                                    var v15 int64 = 1
                                    p0 = 42
                                    _ = v15
                                }
                            }
                        case "foo bar":
                            var v16 float64 = 1.5
                            _ = v16
                        case "hello":
                            v17 := GenS1{F0: -79, F1: "a"}
                            _ = v17
                        default:
                            v18 := []int{(fi8 / (((fi8 % ((-56) | 1))) | 1)), fi8}
                            _ = v18
                        }
                        switch fi8 {
                        case 9223372036854775807:
                            if (!(32767 < -475)) {
                                if (!("foo bar" == (("a" + "\n\t") + "foo bar"))) {
                                    v19 := GenS2{F0: 255}
                                    if true {
                                        p0 -= p0
                                        var v20 float64 = -2.25
                                        for fi21 := 0; fi21 < 4; fi21++ {
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(p0, fi8)
                                            }
                                            _ = fi21
                                        }
                                        _ = v20
                                    } else {
                                        var v22 uint16 = 826
                                        v23 := GenS1{F0: -128, F1: ("a" + "\n\t")}
                                        _ = v22
                                        _ = v23
                                    }
                                    _ = v19
                                } else {
                                    var v24 uint32 = 1
                                    _ = v24
                                }
                            } else {
                                fmt.Println(fi8)
                            }
                        }
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v26 := GenS3{F0: -128, F1: (^(p0))}
                            fmt.Println(p0, fi25)
                            _ = fi25
                            _ = v26
                        }
                    }
                }
            case -9223372036854775808:
                var v27 uint8 = 255
                _ = v27
            case 885:
                v28 := []int{42, (fi8 / ((fi8) | 1)), (fi8 - (fi8 + fi8))}
                _ = v28
            default:
                for fi29 := 0; fi29 < 7; fi29++ {
                    v30 := GenS4{F0: 42, F1: (^(p0))}
                    fmt.Println(fi8, p0, p0)
                    _ = fi29
                    _ = v30
                }
            }
            _ = fi8
        }
    case "foo bar":
        p0 = (p0 & p0)
    default:
        if false {
            switch "hello" {
            case "\n\t":
                var v31 float64 = ((1e10 - ((-2.25 + (0.0 * 3.14159)) + 1.5)) + (float64(((float32(0.1) + 1.5) + 0.1)) * 0.1))
                _ = v31
            case "a":
                var v32 int64 = -1
                _ = v32
            case "foo bar":
                var v33 string = "hello"
                _ = v33
            default:
                var v34 uint8 = 191
                _ = v34
            }
        } else {
            var v35 float32 = (1.5 + (1e10 / -2.25))
            v36 := GenS1{F0: -14, F1: ((("foo bar" + "hello") + ("foo bar" + "a")) + "\n\t")}
            _ = v35
            _ = v36
        }
    }
    switch "\n\t" {
    case "":
        v37 := []int{9223372036854775807, -853, 9223372036854775807, 32768}
        _ = v37
    }
    fmt.Println(p0)
    p0 = 80
    p0 = 42
    _ = p0
    return 1
}
func genFn1(p0 uint16, p1 uint64, p2 string) uint32 {
    switch 32767 {
    case 128:
        if true {
            for fi38 := 0; fi38 < 1; fi38++ {
                v39 := GenS5{F0: 1, F1: 0, F2: 42}
                p0 = (p0 & (p0 >> ((748) & 63)))
                _ = fi38
                _ = v39
            }
            p0 = p0
            p1 = (p1 << ((2) & 63))
        }
    case 2147483647:
        var v40 int16 = -43
        _ = v40
    case 945:
        switch 2147483647 {
        case 128:
            var v41 int16 = -1
            _ = v41
        case -852:
            var v42 uint = 18446744073709551615
            _ = v42
        default:
            v43 := GenS3{F0: 42, F1: 255}
            _ = v43
        }
    }
    if (1 >= 9223372036854775807) {
        fmt.Println(p1)
        if (2147483647 > 0) {
            p0 += p0
        } else {
            switch -9223372036854775808 {
            case 0:
                switch p2 {
                case "":
                    switch -1 {
                    case 0:
                        var v44 uint32 = 607
                        _ = v44
                    default:
                        if ((127 == 202) && (!(9223372036854775807 >= 1))) {
                            if false {
                                if (true && false) {
                                    if (!(true || ((230 == 109) || (966 >= 32768)))) {
                                        var v45 float32 = -2.25
                                        v46 := []int{127}
                                        var v47 int64 = 80
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    } else {
                                        p2 = p2
                                        fmt.Println(p1, p0)
                                    }
                                    if false {
                                        var v48 int = 32767
                                        p1 += 18446744073709551615
                                        _ = v48
                                    }
                                    v49 := []int{32767, 180, -769, 32767}
                                    _ = v49
                                }
                                var v50 uint16 = p0
                                _ = v50
                            }
                        }
                    }
                case "hello":
                    var v51 int8 = -1
                    _ = v51
                case "\n\t":
                    v52 := []int{42, 0, -9223372036854775808}
                    _ = v52
                }
            case -757:
                p0 = 0
            case 4294967295:
                p2 = p2
            default:
                p2 = ("hello" + ("" + "\n\t"))
            }
            v53 := GenS4{F0: 133, F1: 127}
            for fi54 := 0; fi54 < 7; fi54++ {
                v55 := GenS6{F0: 0}
                _ = fi54
                _ = v55
            }
            _ = v53
        }
        v56 := GenS7{F0: -1, F1: -1, F2: 206}
        _ = v56
    } else {
        p0 = p0
        var v57 string = "a"
        if (32767 >= (^(p0))) {
            v58 := []int{611, 42}
            var v59 string = "a"
            _ = v58
            _ = v59
        } else {
            v57 = v57
        }
        _ = v57
    }
    v60 := []int{1, 128}
    var v61 uint64 = p1
    var v62 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v60
    _ = v61
    _ = v62
    return (v62 >> ((65535) & 63))
}

func main() {
    v1 := []int{}
    v2 := []uint32{2147483647, 580, 557, 2}
    v3 := []string{}
    var v4 uint64 = 297
    var v5 bool = ("\n\t" > "")
    v6 := []uint32{2147483648}
    switch 0 {
    case -794:
        fmt.Println(v4, v5)
    }
    switch "hello" {
    case "foo bar":
        var v63 uint = genFn0(20)
        _ = v63
    }
    var v64 uint8 = 32
    if v5 {
        v64 = v64
        if true {
            var v65 uint16 = 2
            if (("foo bar" < "\n\t") || v5) {
                for fi66 := 0; fi66 < 4; fi66++ {
                    v4 += 4294967295
                    switch "hello" {
                    case "":
                        var v67 int = 9223372036854775807
                        _ = v67
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        v65 += v65
                    }
                    switch (("foo bar" + "\n\t") + "foo bar") {
                    case "a":
                        v65 = 42
                    case "":
                        var v68 string = ("hello" + "")
                        _ = v68
                    default:
                        fmt.Println(v65, v5)
                    }
                    _ = fi66
                }
                v65 = (^(v65))
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v64 -= v64
            _ = v65
        } else {
            v5 = (482 < -1)
            v6 = append(v6, 2147483647)
        }
    }
    v4 -= (v4 & v4)
    var v69 uint8 = (v64 + v64)
    v70 := GenS4{F0: 2147483647, F1: (v64 ^ v64)}
    switch 1 {
    case 763:
        for fi71 := 0; fi71 < 3; fi71++ {
            var v72 uint16 = 255
            _ = fi71
            _ = v72
        }
    case 32767:
        switch "" {
        case "":
            fmt.Println(v4)
        case "hello":
            if (v5 || true) {
                for fi73 := 0; fi73 < 16; fi73++ {
                    switch (fi73 | fi73) {
                    case 42:
                        if false {
                            fmt.Println(fi73, fi73)
                            var v74 uint64 = v4
                            _ = v74
                        }
                    case 32768:
                        var v75 float32 = ((((0.1 - -2.25) * 1.5) * 0.0) - (1.5 + float32(1.5)))
                        _ = v75
                    case -415:
                        fmt.Println(v4, v4, v69)
                    }
                    for fi76 := 0; fi76 < 6; fi76++ {
                        _ = len(v6)
                        v5 = v5
                        if false {
                            v69 = 0
                        }
                        if (v5 && v5) {
                            var v77 int16 = 42
                            var v78 uint64 = 18446744073709551615
                            _ = v77
                            _ = v78
                        }
                        _ = fi76
                    }
                    _ = fi73
                }
                v5 = v5
            } else {
                var v79 bool = true
                _ = v79
            }
        case "a":
            var v80 int32 = -1
            _ = v80
        }
    }
    var v81 int = 32767
    switch (v81 & 9223372036854775807) {
    case -150:
        if ((!(32767 < 295)) && v5) {
            var v82 int32 = 2147483647
            if v5 {
                fmt.Println(v5, v82, v82)
                v69 = (^(v64))
                var v83 uint8 = (^(v69))
                _ = v83
            }
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = v82
        } else {
            fmt.Println(v5, v64, v4)
            fmt.Println(v64, v4, v69)
            if false {
                if false {
                    if true {
                        for fi84 := 0; fi84 < 13; fi84++ {
                            var v85 string = ("hello" + "")
                            for fi86 := 0; fi86 < 18; fi86++ {
                                fmt.Println(v69)
                                switch "foo bar" {
                                case "\n\t":
                                    switch fi84 {
                                    case 822:
                                        switch "hello" {
                                        case "":
                                            var v87 float64 = float64(((1.5 - 3.14159) - (0.0 + ((1.5 + 0.1) / float32(0.1)))))
                                            _ = v87
                                        case "a":
                                            v5 = (!(genFn0(v69) > 18446744073709551615))
                                            v64 += v64
                                            fmt.Println(v5, v64)
                                            if v5 {
                                                switch v81 {
                                                case 2147483647:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                default:
                                                    if v5 {
                                                        v81 -= (fi84 - (-(fi86)))
                                                        v4 += v4
                                                    }
                                                }
                                                var v88 int8 = -81
                                                var v89 bool = v5
                                                _ = v88
                                                _ = v89
                                            } else {
                                                v4 = v4
                                                v4 -= v4
                                            }
                                        case "\n\t":
                                            v90 := GenS8{F0: (float32((0.0 + 3.14159)) < 1e10), F1: 4294967295}
                                            v81 += (fi84 - 726)
                                            var v91 uint32 = 65535
                                            if len(v6) > 0 {
                                                v6[0] = (v91 * (^(v91)))
                                            }
                                            _ = len(v2)
                                            _ = v90
                                            _ = v91
                                        default:
                                            v81 = (fi84 * (v81 & (fi86 & (fi86 + (fi84 - fi84)))))
                                        }
                                    default:
                                        v3 = append(v3, (v85 + ""))
                                        var v92 float32 = 0.0
                                        var v93 float32 = 1e10
                                        _ = v92
                                        _ = v93
                                    }
                                default:
                                    var v94 int32 = 2147483647
                                    _ = v94
                                }
                                var v95 int32 = -2147483648
                                _ = fi86
                                _ = v95
                            }
                            _ = fi84
                            _ = v85
                        }
                        v64 = 138
                    } else {
                        var v96 float32 = float32((1.5 + 0.0))
                        for fi97 := 0; fi97 < 18; fi97++ {
                            v1 = append(v1, (^(fi97)))
                            v98 := GenS9{F0: 4294967295}
                            _ = fi97
                            _ = v98
                        }
                        _ = v96
                    }
                }
                v64 = (^(v64))
            } else {
                if len(v1) > 0 {
                    v1[0] = v81
                }
                switch (v81 << ((255) & 63)) {
                case 4294967295:
                    switch v81 {
                    case -60:
                        switch "hello" {
                        case "foo bar":
                            v4 = v4
                        case "":
                            var v99 int8 = -1
                            _ = v99
                        }
                    default:
                        var v100 int16 = -629
                        _ = v100
                    }
                }
            }
        }
    case 32767:
        fmt.Println(v81, v64)
    default:
        for fi101 := 0; fi101 < 17; fi101++ {
            if (-1 == 42) {
                var v102 int = fi101
                for fi103 := 0; fi103 < 20; fi103++ {
                    var v104 uint64 = uint64(v64)
                    var v105 bool = (!(-101 <= 913))
                    switch "" {
                    case "a":
                        var v106 float32 = 0.1
                        _ = v106
                    case "\n\t":
                        switch "foo bar" {
                        case "a":
                            v107 := GenS10{F0: "hello"}
                            _ = v107
                        default:
                            v108 := GenS8{F0: (2147483647 > 869), F1: 42}
                            _ = v108
                        }
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    switch 349 {
                    case 42:
                        var v109 string = "a"
                        _ = v109
                    }
                    _ = len(v6)
                    _ = fi103
                    _ = v104
                    _ = v105
                }
                if v5 {
                    var v110 float64 = -2.25
                    if v5 {
                        switch ("foo bar" + ("a" + (("hello" + "") + ("foo bar" + "a")))) {
                        case "hello":
                            switch -541 {
                            case 4294967295:
                                fmt.Println(v4, v64, v4)
                            case 2147483647:
                                for fi111 := 0; fi111 < 3; fi111++ {
                                    var v112 uint8 = v64
                                    if (true || v5) {
                                        var v113 bool = (v64 <= (v64 % (((v64 - v112)) | 1)))
                                        var v114 float32 = 0.0
                                        v110 -= (1e10 - 1.5)
                                        _ = v113
                                        _ = v114
                                    }
                                    v5 = (0.0 <= 3.14159)
                                    _ = fi111
                                    _ = v112
                                }
                            }
                        case "a":
                            fmt.Println(v69)
                        case "\n\t":
                            if v5 {
                                switch "foo bar" {
                                case "a":
                                    var v115 uint8 = (v64 * (v64 - v64))
                                    _ = v115
                                case "":
                                    v81 = (v81 % ((fi101) | 1))
                                case "foo bar":
                                    v116 := GenS8{F0: false, F1: 480}
                                    _ = v116
                                default:
                                    for fi117 := 0; fi117 < 1; fi117++ {
                                        var v118 uint64 = 1
                                        _ = fi117
                                        _ = v118
                                    }
                                }
                                var v119 int64 = 1
                                v120 := GenS11{F0: 997, F1: (-460 > 1000), F2: 0}
                                _ = v119
                                _ = v120
                            }
                        }
                        for fi121 := 0; fi121 < 20; fi121++ {
                            if (((^(v64)) == v64) || v5) {
                                switch "hello" {
                                case "\n\t":
                                    if v5 {
                                        var v122 bool = v5
                                        v123 := GenS12{F0: (^(v64)), F1: (-32768 != 0), F2: (v69 - 255)}
                                        var v124 uint = 0
                                        _ = v122
                                        _ = v123
                                        _ = v124
                                    } else {
                                        if v5 {
                                            var v125 float64 = v110
                                            switch ("foo bar" + "foo bar") {
                                            case "hello":
                                                v126 := GenS12{F0: 100, F1: (v110 == (v110 / 1e10)), F2: (v64 % ((0) | 1))}
                                                _ = v126
                                            case "a":
                                                v102 += fi101
                                            }
                                            v102 -= -1
                                            _ = v125
                                        } else {
                                            v81 = (v81 >> ((18446744073709551615) & 63))
                                        }
                                        if (!false) {
                                            v3 = append(v3, "\n\t")
                                        } else {
                                            v127 := GenS13{F0: 42, F1: 65535, F2: -238}
                                            switch v102 {
                                            case 4294967295:
                                                v64 = uint8(v4)
                                            }
                                            _ = v127
                                        }
                                        v4 = 42
                                        switch 32768 {
                                        case 42:
                                            v81 = (^(v81))
                                        case 9223372036854775807:
                                            if (1 == (^(v69))) {
                                                v69 -= v69
                                            } else {
                                                switch fi101 {
                                                case 827:
                                                    v64 = (v69 / (((v69 * v69)) | 1))
                                                    if (!(!(-1 > 1))) {
                                                        var v128 int16 = 1000
                                                        _ = v128
                                                    }
                                                    if true {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if false {
                                                            var v129 uint = 2
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v130 := GenS11{F0: (v102 / (((v81 - fi121)) | 1)), F1: true, F2: 1}
                                                            _ = v129
                                                            _ = v130
                                                        }
                                                        switch "\n\t" {
                                                        case "a":
                                                            v81 = v102
                                                        case "":
                                                            fmt.Println(v69, v69, v69)
                                                        case "hello":
                                                            v131 := GenS14{F0: 2147483648, F1: 2}
                                                            _ = v131
                                                        default:
                                                            var v132 int = 0
                                                            _ = v132
                                                        }
                                                    } else {
                                                        v133 := GenS15{F0: v102, F1: -1, F2: -1}
                                                        if (v5 && true) {
                                                            v3 = append(v3, "")
                                                        }
                                                        switch fi121 {
                                                        case 32767:
                                                            _ = len(v6)
                                                        case 0:
                                                            fmt.Println(v4, v64, fi121)
                                                        case -9223372036854775808:
                                                            if false {
                                                                var v134 float32 = 0.0
                                                                fmt.Println(v110, v4, fi101)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v134
                                                            } else {
                                                                var v135 uint64 = v4
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v136 := GenS3{F0: 0, F1: (v64 + 127)}
                                                                    _ = v136
                                                                case "":
                                                                    if (65535 >= 786) {
                                                                        v64 += 127
                                                                    }
                                                                    v81 -= (v81 + v102)
                                                                    v137 := GenS16{F0: uint16(v69), F1: -677, F2: genFn0(255)}
                                                                    v135 = (v4 + (v4 | 246))
                                                                    _ = v137
                                                                default:
                                                                    var v138 uint16 = 65535
                                                                    _ = v138
                                                                }
                                                                var v139 int16 = -458
                                                                var v140 int16 = v139
                                                                _ = v135
                                                                _ = v139
                                                                _ = v140
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v141 int16 = 430
                                                                v81 = 32767
                                                                _ = v141
                                                            case "":
                                                                if v5 {
                                                                    v142 := GenS8{F0: (((3.14159 - float64(1e10)) + (-2.25 * 1e10)) != v110), F1: 0}
                                                                    _ = v142
                                                                } else {
                                                                    var v143 int32 = 42
                                                                    _ = v143
                                                                }
                                                            case "a":
                                                                if (533 < 716) {
                                                                    switch ("a" + "a") {
                                                                    case "foo bar":
                                                                        v144 := GenS3{F0: 0, F1: (v64 + v69)}
                                                                        _ = v144
                                                                    default:
                                                                        var v145 uint = 420
                                                                        _ = v145
                                                                    }
                                                                    v146 := GenS17{F0: (!v5), F1: -1}
                                                                    switch (v102 << ((18446744073709551615) & 63)) {
                                                                    case -997:
                                                                        if (94 < -128) {
                                                                            var v147 string = ""
                                                                            v148 := GenS1{F0: -1, F1: ("a" + "foo bar")}
                                                                            var v149 float64 = 1e10
                                                                            _ = v147
                                                                            _ = v148
                                                                            _ = v149
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                if false {
                                                                                    if v5 {
                                                                                        switch ("hello" + "a") {
                                                                                        case "\n\t":
                                                                                            var v150 int64 = -9223372036854775808
                                                                                            _ = v150
                                                                                        case "hello":
                                                                                            switch (fi101 ^ (^(fi101))) {
                                                                                            case 20:
                                                                                                if (-1 < 127) {
                                                                                                    _ = len(v2)
                                                                                                    v81 -= (fi121 ^ (fi121 + v81))
                                                                                                } else {
                                                                                                    if v5 {
                                                                                                        v64 = v64
                                                                                                        fmt.Println(v4, v5)
                                                                                                        switch "\n\t" {
                                                                                                        case "\n\t":
                                                                                                            var v151 int = v102
                                                                                                            _ = v151
                                                                                                        case "a":
                                                                                                            if v5 {
                                                                                                                var v152 float32 = float32(((float64((1.5 + -2.25)) + v110) / v110))
                                                                                                                v153 := GenS18{F0: v64, F1: genFn1(497, v4, "\n\t"), F2: 803}
                                                                                                                _ = v152
                                                                                                                _ = v153
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            if (((1 <= -32768) && v5) || true) {
                                                                                                                var v154 uint64 = v4
                                                                                                                v5 = (v5 || v5)
                                                                                                                if true {
                                                                                                                    if (-131 >= 473) {
                                                                                                                        if ((^(v4)) > (v4 ^ v4)) {
                                                                                                                            var v155 int32 = 0
                                                                                                                            if len(v6) > 0 {
                                                                                                                                v6[0] = 2147483647
                                                                                                                            }
                                                                                                                            fmt.Println(v4)
                                                                                                                            _ = v155
                                                                                                                        }
                                                                                                                        if len(v3) > 0 {
                                                                                                                            v3[0] = "hello"
                                                                                                                        }
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(fi101, v110)
                                                                                                                    v4 = v154
                                                                                                                }
                                                                                                                _ = v154
                                                                                                            } else {
                                                                                                                switch "foo bar" {
                                                                                                                case "":
                                                                                                                    switch v102 {
                                                                                                                    case 9223372036854775807:
                                                                                                                        v110 -= v110
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v5 = (-64 == 0)
                                                                                                                }
                                                                                                                fmt.Println(v4, fi121, v69)
                                                                                                            }
                                                                                                        }
                                                                                                        v156 := GenS1{F0: 1, F1: ("" + (("hello" + ("a" + "foo bar")) + ("\n\t" + "")))}
                                                                                                        _ = v156
                                                                                                    } else {
                                                                                                        var v157 string = "a"
                                                                                                        _ = v157
                                                                                                    }
                                                                                                    if (!true) {
                                                                                                        var v158 int64 = -1
                                                                                                        fmt.Println(v64)
                                                                                                        fmt.Println(fi121)
                                                                                                        if (!v5) {
                                                                                                            v159 := GenS18{F0: v64, F1: 65535, F2: 2147483647}
                                                                                                            var v160 float64 = ((0.1 + (float64((1e10 - 3.14159)) * float64((0.1 - (float32(-2.25) * 0.0))))) - (v110 / 1.5))
                                                                                                            if ((!true) && false) {
                                                                                                                v161 := GenS7{F0: -282, F1: -1, F2: 1}
                                                                                                                v3 = append(v3, ("" + "foo bar"))
                                                                                                                fmt.Println(v69, v69)
                                                                                                                _ = v161
                                                                                                            } else {
                                                                                                                if (0 == 1) {
                                                                                                                    if (v81 == (fi121 & fi121)) {
                                                                                                                        var v162 int32 = -629
                                                                                                                        if len(v6) > 0 {
                                                                                                                            v6[0] = 2147483647
                                                                                                                        }
                                                                                                                        var v163 uint64 = 42
                                                                                                                        if (int(v64) > (v102 << ((0) & 63))) {
                                                                                                                            switch "\n\t" {
                                                                                                                            case "foo bar":
                                                                                                                                v164 := GenS19{F0: 0}
                                                                                                                                _ = v164
                                                                                                                            }
                                                                                                                            v158 -= v158
                                                                                                                            var v165 uint16 = 2
                                                                                                                            _ = v165
                                                                                                                        }
                                                                                                                        _ = v162
                                                                                                                        _ = v163
                                                                                                                    }
                                                                                                                    v4 = 0
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v159
                                                                                                            _ = v160
                                                                                                        } else {
                                                                                                            v3 = append(v3, "")
                                                                                                            if true {
                                                                                                                if (!((-275 != 1000) && (!v5))) {
                                                                                                                    v69 -= 48
                                                                                                                } else {
                                                                                                                    v64 -= 1
                                                                                                                    if len(v3) >= 2 {
                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v64, v5, v158)
                                                                                                            }
                                                                                                            if (!(2147483647 != 1)) {
                                                                                                                var v166 uint8 = 252
                                                                                                                var v167 int16 = 858
                                                                                                                v1 = append(v1, (^(fi101)))
                                                                                                                var v168 int32 = -2147483648
                                                                                                                _ = v166
                                                                                                                _ = v167
                                                                                                                _ = v168
                                                                                                            } else {
                                                                                                                v169 := GenS20{F0: 18, F1: 878, F2: 0}
                                                                                                                v2 = append(v2, 0)
                                                                                                                _ = v169
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v158
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v69 -= v69
                                                                                            }
                                                                                        case "":
                                                                                            fmt.Println(fi101)
                                                                                        }
                                                                                        if (973 == -1) {
                                                                                            v5 = v5
                                                                                        } else {
                                                                                            v81 += (fi121 - 243)
                                                                                        }
                                                                                    } else {
                                                                                        var v170 uint32 = 2147483647
                                                                                        fmt.Println(v170, v81, v69)
                                                                                        _ = v170
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(fi101, v4)
                                                                            case "foo bar":
                                                                                fmt.Println(fi121)
                                                                            default:
                                                                                v171 := GenS21{F0: 1, F1: genFn1(770, (v4 % ((v4) | 1)), ""), F2: genFn0(v69)}
                                                                                var v172 int32 = int32(v4)
                                                                                _ = v171
                                                                                _ = v172
                                                                            }
                                                                            fmt.Println(v102, fi101)
                                                                            var v173 uint32 = genFn1(255, (^(v4)), (((("" + (("foo bar" + "hello") + "hello")) + ("" + "hello")) + "a") + "hello"))
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            _ = v173
                                                                        }
                                                                    case 4294967295:
                                                                        var v174 uint64 = (v4 ^ (v4 | (v4 % (((v4 * (^(v4)))) | 1))))
                                                                        _ = v174
                                                                    }
                                                                    _ = v146
                                                                } else {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                }
                                                            }
                                                            v4 += 42
                                                        default:
                                                            v81 += (fi101 >> ((4294967295) & 63))
                                                        }
                                                        _ = v133
                                                    }
                                                case 0:
                                                    v102 = (fi121 >> ((1) & 63))
                                                case 32768:
                                                    var v175 float64 = (float64(float32(float64((((0.0 + ((0.1 / ((((1e10 * 1e10) - 1e10) + 0.1) * 1e10)) - 3.14159)) * (float32(v110) - -2.25)) * 3.14159)))) - 1e10)
                                                    _ = v175
                                                default:
                                                    v102 = (v81 / ((42) | 1))
                                                }
                                            }
                                        case 1:
                                            if (!(v5 || ((!(true || v5)) && false))) {
                                                switch fi121 {
                                                case 1:
                                                    if v5 {
                                                        _ = len(v6)
                                                        var v176 uint32 = 2147483648
                                                        _ = v176
                                                    }
                                                }
                                            }
                                        }
                                        var v177 int8 = 1
                                        _ = v177
                                    }
                                case "hello":
                                    var v178 uint = 971
                                    _ = v178
                                case "":
                                    fmt.Println(v5)
                                default:
                                    fmt.Println(v81, v110)
                                }
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                            }
                            if (!(((float32(float64(0.1)) / 1e10) - float32(3.14159)) <= -2.25)) {
                                var v179 string = ("" + "")
                                var v180 int = -1
                                switch v102 {
                                case 4294967295:
                                    if (v179 >= v179) {
                                        var v181 int = 298
                                        _ = v181
                                    } else {
                                        switch int(v64) {
                                        case -9223372036854775808:
                                            if v5 {
                                                if false {
                                                    if true {
                                                        v64 = v64
                                                        fmt.Println(fi101, v180, v102)
                                                        var v182 int = (-(fi121))
                                                        v183 := GenS14{F0: 151, F1: uint8(v4)}
                                                        if v5 {
                                                            v102 += 32768
                                                            v110 += ((float64((1e10 * ((0.1 - (-2.25 + -2.25)) + ((1.5 + -2.25) - 3.14159)))) * v110) + (v110 + (float64((float32(v110) - float32(v110))) / (v110 - (v110 + (0.0 + float64(1e10)))))))
                                                        }
                                                        _ = v182
                                                        _ = v183
                                                    }
                                                    switch v102 {
                                                    case -415:
                                                        v184 := GenS20{F0: 65535, F1: 811, F2: 128}
                                                        _ = v184
                                                    }
                                                    if v5 {
                                                        fmt.Println(v69)
                                                    } else {
                                                        fmt.Println(v81)
                                                        var v185 float64 = (v110 / ((v110 / -2.25) / v110))
                                                        _ = v185
                                                    }
                                                    switch 857 {
                                                    case 678:
                                                        v186 := GenS22{F0: (fi121 | (v102 << ((uint(v69)) & 63))), F1: int64(fi101)}
                                                        _ = v186
                                                    case 0:
                                                        v187 := GenS23{F0: (v102 << ((42) & 63))}
                                                        v4 = v4
                                                        if (((v69 >> ((777) & 63)) >= v64) && true) {
                                                            if v5 {
                                                                if false {
                                                                    v188 := GenS5{F0: 2147483647, F1: 1, F2: -9223372036854775808}
                                                                    _ = len(v2)
                                                                    switch (fi121 & (fi101 << ((42) & 63))) {
                                                                    case 127:
                                                                        v110 += (v110 / v110)
                                                                    case -9223372036854775808:
                                                                        v110 = 3.14159
                                                                    default:
                                                                        fmt.Println(fi121)
                                                                    }
                                                                    _ = v188
                                                                } else {
                                                                    var v189 int16 = 1000
                                                                    _ = v189
                                                                }
                                                                v81 = 0
                                                                v3 = append(v3, (v179 + v179))
                                                                fmt.Println(v81)
                                                            }
                                                            if true {
                                                                if (0 == genFn0(2)) {
                                                                    v190 := GenS20{F0: 255, F1: int32(v64), F2: 0}
                                                                    _ = v190
                                                                } else {
                                                                    var v191 int64 = -365
                                                                    fmt.Println(fi121, v110, v102)
                                                                    _ = v191
                                                                }
                                                                fmt.Println(v64, fi101, v4)
                                                                v69 = 127
                                                            } else {
                                                                v64 = (v64 ^ (v64 << ((4294967295) & 63)))
                                                                if (!(!false)) {
                                                                    var v192 uint32 = 562
                                                                    v193 := GenS24{F0: fi121, F1: -659, F2: 18446744073709551615}
                                                                    v194 := GenS2{F0: 45}
                                                                    fmt.Println(v180, v4, v180)
                                                                    var v195 uint16 = 255
                                                                    _ = v192
                                                                    _ = v193
                                                                    _ = v194
                                                                    _ = v195
                                                                }
                                                                v196 := GenS7{F0: 0, F1: 0, F2: 1}
                                                                _ = v196
                                                            }
                                                            var v197 int8 = -48
                                                            _ = v197
                                                        }
                                                        _ = v187
                                                    case 166:
                                                        v64 = (v64 | 1)
                                                    }
                                                    var v198 int16 = 1
                                                    _ = v198
                                                } else {
                                                    v1 = append(v1, fi101)
                                                    switch (v179 + "\n\t") {
                                                    case "foo bar":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    default:
                                                        if true {
                                                            if v5 {
                                                                var v199 float64 = (float64((1e10 * 3.14159)) / v110)
                                                                fmt.Println(v179)
                                                                _ = v199
                                                            } else {
                                                                v200 := GenS25{F0: v5, F1: -72, F2: 100}
                                                                _ = v200
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(fi101)
                                                            case "foo bar":
                                                                v102 -= v81
                                                            default:
                                                                v4 = (v4 >> ((580) & 63))
                                                            }
                                                            var v201 string = "foo bar"
                                                            switch fi101 {
                                                            case -2147483648:
                                                                var v202 string = v179
                                                                _ = v202
                                                            case 0:
                                                                v5 = true
                                                            case -1:
                                                                var v203 uint8 = 0
                                                                _ = v203
                                                            default:
                                                                switch v180 {
                                                                case 694:
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        switch (v179 + v201) {
                                                                        case "hello":
                                                                            v204 := GenS18{F0: (v69 >> ((18446744073709551615) & 63)), F1: 0, F2: 2147483647}
                                                                            _ = v204
                                                                        case "foo bar":
                                                                            switch fi101 {
                                                                            case 0:
                                                                                fmt.Println(fi121)
                                                                            case 270:
                                                                                var v205 float64 = (v110 / (v110 + ((v110 / v110) + -2.25)))
                                                                                _ = v205
                                                                            case -2147483648:
                                                                                fmt.Println(v179)
                                                                            }
                                                                        case "":
                                                                            if v5 {
                                                                                switch (-(fi101)) {
                                                                                case 1:
                                                                                    if v5 {
                                                                                        var v206 int16 = -1
                                                                                        v201 = v179
                                                                                        v201 = "hello"
                                                                                        _ = v206
                                                                                    } else {
                                                                                        v207 := GenS17{F0: (!(!(0 >= 1))), F1: 55}
                                                                                        var v208 float32 = 0.1
                                                                                        _ = v207
                                                                                        _ = v208
                                                                                    }
                                                                                }
                                                                                v81 = v180
                                                                                var v209 int = (v102 | (v102 - 32768))
                                                                                var v210 uint8 = (v64 - v64)
                                                                                _ = v209
                                                                                _ = v210
                                                                            } else {
                                                                                fmt.Println(v69, v180, v4)
                                                                            }
                                                                        default:
                                                                            if v5 {
                                                                                var v211 uint64 = 42
                                                                                _ = v211
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        _ = len(v1)
                                                                    case "\n\t":
                                                                        v81 += (-(v102))
                                                                    }
                                                                case 155:
                                                                    v1 = append(v1, fi101)
                                                                default:
                                                                    var v212 int = v102
                                                                    _ = v212
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "a":
                                                                if v5 {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    fmt.Println(v201, v110, fi121)
                                                                    if (!v5) {
                                                                        if v5 {
                                                                            if v5 {
                                                                                var v213 int8 = 127
                                                                                var v214 int8 = v213
                                                                                _ = v213
                                                                                _ = v214
                                                                            } else {
                                                                                if ((-103 != 0) && true) {
                                                                                    var v215 uint64 = (v4 / ((v4) | 1))
                                                                                    v216 := GenS21{F0: 0, F1: 4294967295, F2: genFn0(173)}
                                                                                    var v217 int16 = int16(v64)
                                                                                    _ = v215
                                                                                    _ = v216
                                                                                    _ = v217
                                                                                }
                                                                                v218 := GenS11{F0: fi121, F1: ((0.1 / -2.25) != 1.5), F2: 2147483647}
                                                                                v219 := GenS26{F0: 2147483647, F1: -818}
                                                                                _ = v218
                                                                                _ = v219
                                                                            }
                                                                            fmt.Println(fi101, v69, v5)
                                                                            var v220 int32 = -2147483648
                                                                            _ = v220
                                                                        }
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v221 := GenS27{F0: 42, F1: genFn0((v69 & (v69 ^ v64)))}
                                                                _ = v221
                                                            case "\n\t":
                                                                var v222 int16 = 870
                                                                _ = v222
                                                            default:
                                                                var v223 bool = (-608 == (v102 - (v81 << ((65535) & 63))))
                                                                if v5 {
                                                                    if (!(-2.25 > ((float32(float64((0.0 + 1.5))) - 0.1) - (1e10 / (-2.25 / (1.5 * (1.5 - (0.1 - (-2.25 - 1e10))))))))) {
                                                                        v4 = (v4 ^ v4)
                                                                    }
                                                                    var v224 int32 = -2147483648
                                                                    _ = v224
                                                                } else {
                                                                    switch (v102 << ((495) & 63)) {
                                                                    case -333:
                                                                        var v225 uint64 = (v4 * (v4 ^ 0))
                                                                        _ = v225
                                                                    case 566:
                                                                        _ = len(v3)
                                                                    case 775:
                                                                        if (v69 != (v69 << ((1) & 63))) {
                                                                            var v226 uint32 = 2
                                                                            fmt.Println(v64)
                                                                            v180 -= (v180 & (v102 >> ((genFn0(v69)) & 63)))
                                                                            switch v179 {
                                                                            case "hello":
                                                                                if v5 {
                                                                                    var v227 float64 = v110
                                                                                    _ = v227
                                                                                }
                                                                            case "\n\t":
                                                                                var v228 string = "foo bar"
                                                                                _ = v228
                                                                            }
                                                                            _ = v226
                                                                        } else {
                                                                            var v229 int8 = 31
                                                                            switch v179 {
                                                                            case "foo bar":
                                                                                var v230 int16 = 32767
                                                                                _ = v230
                                                                            case "a":
                                                                                switch 2147483647 {
                                                                                case 0:
                                                                                    var v231 int16 = 42
                                                                                    _ = v231
                                                                                case 128:
                                                                                    _ = len(v2)
                                                                                    fmt.Println(v5, v180)
                                                                                case 4294967295:
                                                                                    if (162 == 462) {
                                                                                        v81 -= (v102 & 4294967295)
                                                                                    } else {
                                                                                        switch v81 {
                                                                                        case -673:
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                v6 = append(v6, genFn1(709, v4, v179))
                                                                                            default:
                                                                                                v223 = ((v229 / (((v229 - v229)) | 1)) < (v229 ^ 100))
                                                                                            }
                                                                                        case 649:
                                                                                            fmt.Println(v201, v4)
                                                                                        default:
                                                                                            v232 := GenS25{F0: true, F1: v229, F2: (v229 % ((36) | 1))}
                                                                                            _ = v232
                                                                                        }
                                                                                        if true {
                                                                                            v1 = append(v1, (v102 % (((v102 ^ fi121)) | 1)))
                                                                                        } else {
                                                                                            v5 = (-889 >= (^(v102)))
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v233 float32 = 3.14159
                                                                            _ = v229
                                                                            _ = v233
                                                                        }
                                                                    }
                                                                    switch v179 {
                                                                    case "a":
                                                                        if v223 {
                                                                            v234 := GenS28{F0: 42}
                                                                            var v235 uint = 136
                                                                            var v236 uint64 = 42
                                                                            if false {
                                                                                var v237 string = (v179 + v179)
                                                                                var v238 string = ("" + "a")
                                                                                v179 = v201
                                                                                _ = v237
                                                                                _ = v238
                                                                            }
                                                                            switch v201 {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    v6 = append(v6, 2147483647)
                                                                                    if v223 {
                                                                                        fmt.Println(v180)
                                                                                        var v239 float32 = (-2.25 / 0.1)
                                                                                        v236 = (^(v236))
                                                                                        _ = v239
                                                                                    }
                                                                                    _ = len(v2)
                                                                                }
                                                                            case "hello":
                                                                                v240 := GenS10{F0: (v179 + "foo bar")}
                                                                                _ = v240
                                                                            default:
                                                                                switch v81 {
                                                                                case 671:
                                                                                    v110 += (float64(((0.0 * 0.0) + 0.0)) * v110)
                                                                                case 127:
                                                                                    switch v102 {
                                                                                    case 87:
                                                                                        v241 := GenS29{F0: 52, F1: v179}
                                                                                        _ = v241
                                                                                    case 32768:
                                                                                        var v242 float32 = 0.1
                                                                                        _ = v242
                                                                                    }
                                                                                default:
                                                                                    v102 += 953
                                                                                }
                                                                            }
                                                                            switch v102 {
                                                                            case 236:
                                                                                v236 -= 1
                                                                            }
                                                                            v69 += (v64 * (v69 % (((^(v64))) | 1)))
                                                                            _ = v234
                                                                            _ = v235
                                                                            _ = v236
                                                                        }
                                                                    default:
                                                                        v243 := GenS28{F0: 2147483647}
                                                                        _ = v243
                                                                    }
                                                                    v4 -= 4294967295
                                                                }
                                                                _ = v223
                                                            }
                                                            _ = v201
                                                        } else {
                                                            if v5 {
                                                                v5 = false
                                                                v102 = v102
                                                            } else {
                                                                var v244 uint32 = 65535
                                                                var v245 int8 = -1
                                                                var v246 float32 = 1e10
                                                                var v247 uint64 = v4
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v179 + v179)
                                                                }
                                                                _ = v244
                                                                _ = v245
                                                                _ = v246
                                                                _ = v247
                                                            }
                                                            v64 = (v69 >> ((65) & 63))
                                                            if v5 {
                                                                switch v179 {
                                                                case "":
                                                                    switch v179 {
                                                                    case "\n\t":
                                                                        switch (fi121 % ((42) | 1)) {
                                                                        case 1:
                                                                            v179 = (("" + (v179 + ((v179 + v179) + v179))) + "a")
                                                                        case 32767:
                                                                            if true {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (fi101 | (v102 ^ (v180 % ((fi101) | 1))))
                                                                                }
                                                                            } else {
                                                                                if (32767 > 65535) {
                                                                                    v180 -= 2147483647
                                                                                    fmt.Println(v4, v180, v81)
                                                                                }
                                                                            }
                                                                        case 2147483647:
                                                                            if v5 {
                                                                                if (32767 < -54) {
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        if v5 {
                                                                                            var v248 uint64 = (v4 % ((v4) | 1))
                                                                                            fmt.Println(fi101)
                                                                                            switch v179 {
                                                                                            case "hello":
                                                                                                v179 = "foo bar"
                                                                                            default:
                                                                                                v110 += 1e10
                                                                                            }
                                                                                            _ = v248
                                                                                        } else {
                                                                                            v249 := GenS30{F0: v4}
                                                                                            var v250 uint32 = 2147483648
                                                                                            _ = v249
                                                                                            _ = v250
                                                                                        }
                                                                                        if (-1 == 42) {
                                                                                            if (32767 == 32767) {
                                                                                                switch fi101 {
                                                                                                case 0:
                                                                                                    switch v81 {
                                                                                                    case 4294967295:
                                                                                                        var v251 int8 = 100
                                                                                                        _ = v251
                                                                                                    case 9223372036854775807:
                                                                                                        var v252 float32 = ((0.0 - 0.1) * 3.14159)
                                                                                                        _ = v252
                                                                                                    case 2147483647:
                                                                                                        v69 += (v69 / ((134) | 1))
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    v2 = append(v2, 636)
                                                                                                } else {
                                                                                                    v69 = (^(v69))
                                                                                                    v253 := GenS31{F0: v5, F1: v4}
                                                                                                    _ = v253
                                                                                                }
                                                                                                v179 = v179
                                                                                            }
                                                                                            v254 := GenS1{F0: 100, F1: "a"}
                                                                                            _ = v254
                                                                                        }
                                                                                        fmt.Println(v69)
                                                                                    case "a":
                                                                                        if (1 >= 0) {
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                var v255 uint8 = (v64 + v64)
                                                                                                _ = v255
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v256 int16 = -406
                                                                                        _ = v256
                                                                                    default:
                                                                                        fmt.Println(v102, v180)
                                                                                        var v257 uint32 = 623
                                                                                        _ = v257
                                                                                    }
                                                                                    fmt.Println(fi121)
                                                                                    v180 = (v102 % (((fi121 >> ((uint(v4)) & 63))) | 1))
                                                                                }
                                                                                fmt.Println(v179, fi101, v69)
                                                                                var v258 int = (fi121 << ((731) & 63))
                                                                                _ = v258
                                                                            }
                                                                        default:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v179
                                                                            }
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v102, v179)
                                                                default:
                                                                    var v259 int16 = -583
                                                                    _ = v259
                                                                }
                                                                if len(v6) > 0 {
                                                                    v6[0] = genFn1(65535, v4, (v179 + v179))
                                                                }
                                                                v102 -= fi101
                                                            } else {
                                                                v180 += fi121
                                                                v81 = v180
                                                                var v260 int16 = -723
                                                                v110 = ((v110 / v110) - v110)
                                                                if false {
                                                                    var v261 bool = false
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v261
                                                                } else {
                                                                    v262 := GenS13{F0: v260, F1: 2, F2: v260}
                                                                    v64 = v64
                                                                    fmt.Println(fi121, v5)
                                                                    _ = v262
                                                                }
                                                                _ = v260
                                                            }
                                                            switch 813 {
                                                            case 1:
                                                                v110 = 0.1
                                                            case -659:
                                                                fmt.Println(v69, v4)
                                                            default:
                                                                v5 = false
                                                            }
                                                            v263 := GenS27{F0: 843, F1: 789}
                                                            _ = v263
                                                        }
                                                    }
                                                }
                                                if (!v5) {
                                                    v69 -= (v64 % (((v69 + 255)) | 1))
                                                    v180 -= fi101
                                                }
                                                var v264 int = -860
                                                fmt.Println(v5, v81)
                                                _ = len(v1)
                                                _ = v264
                                            }
                                        }
                                    }
                                case 672:
                                    if (2 != 983) {
                                        if (127 < 127) {
                                            switch (v180 / (((v180 - v102)) | 1)) {
                                            case 32767:
                                                fmt.Println(v4, v5, v102)
                                            }
                                            fmt.Println(v5, v102, v4)
                                            v3 = append(v3, "foo bar")
                                        } else {
                                            v265 := GenS26{F0: 2, F1: 42}
                                            if len(v6) > 0 {
                                                v6[0] = 2147483648
                                            }
                                            _ = v265
                                        }
                                        var v266 int = (v81 - (v81 ^ v180))
                                        if v5 {
                                            v267 := GenS23{F0: v81}
                                            v179 = "hello"
                                            _ = v267
                                        }
                                        _ = v266
                                    } else {
                                        fmt.Println(v110, v179, v110)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                case 32768:
                                    var v268 uint16 = 255
                                    _ = v268
                                }
                                _ = v179
                                _ = v180
                            } else {
                                v81 -= (v81 * v81)
                                var v269 int64 = -324
                                _ = v269
                            }
                            if v5 {
                                v270 := GenS27{F0: 2147483647, F1: 255}
                                if (271 >= 926) {
                                    fmt.Println(fi101, fi101, v102)
                                    fmt.Println(fi101)
                                    fmt.Println(fi121, fi101, v81)
                                } else {
                                    v110 -= v110
                                    var v271 uint = 944
                                    v5 = (v5 && (v110 != v110))
                                    _ = v271
                                }
                                _ = v270
                            } else {
                                if (!v5) {
                                    fmt.Println(fi101)
                                    switch "" {
                                    case "foo bar":
                                        var v272 string = ("" + ("\n\t" + "a"))
                                        _ = v272
                                    case "":
                                        fmt.Println(v102, v81)
                                    case "hello":
                                        switch ((("\n\t" + ((("foo bar" + ("foo bar" + (("a" + (("" + "hello") + "hello")) + "foo bar"))) + "\n\t") + ("a" + "foo bar"))) + "hello") + "foo bar") {
                                        case "foo bar":
                                            if v5 {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                fmt.Println(v81)
                                                if len(v6) > 0 {
                                                    v6[0] = uint32(v69)
                                                }
                                            }
                                        default:
                                            v69 += v64
                                        }
                                    default:
                                        var v273 uint64 = v4
                                        _ = v273
                                    }
                                    v110 = v110
                                    v81 -= v102
                                    v110 -= (v110 / v110)
                                } else {
                                    v274 := GenS25{F0: v5, F1: -128, F2: -1}
                                    if v5 {
                                        var v275 uint64 = v4
                                        v64 = 255
                                        if ((0.1 / float64((3.14159 / 0.1))) < ((v110 / 3.14159) - 0.0)) {
                                            switch 4294967295 {
                                            case 0:
                                                switch v102 {
                                                case 128:
                                                    fmt.Println(v5, v64)
                                                case 0:
                                                    _ = len(v3)
                                                default:
                                                    if v5 {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if len(v6) > 0 {
                                                            v6[0] = 2147483647
                                                        }
                                                        if ((-2.25 / v110) == v110) {
                                                            v3 = append(v3, "foo bar")
                                                            v275 += 1
                                                            _ = len(v2)
                                                        }
                                                        var v276 uint8 = v69
                                                        _ = v276
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = 0
                                                        }
                                                        var v277 uint8 = (v69 / ((182) | 1))
                                                        _ = v277
                                                    }
                                                }
                                            }
                                            v102 -= v81
                                        } else {
                                            v5 = (0 > uint16(v275))
                                            if len(v2) > 0 {
                                                v2[0] = genFn1(0, 4294967295, ("hello" + "\n\t"))
                                            }
                                            v81 = fi101
                                            switch (-(fi121)) {
                                            case -791:
                                                v102 = (fi121 ^ (v102 * (fi101 * (v81 >> ((4294967295) & 63)))))
                                            default:
                                                v102 -= fi101
                                            }
                                        }
                                        _ = v275
                                    }
                                    _ = len(v1)
                                    _ = v274
                                }
                                v5 = v5
                                switch "hello" {
                                case "a":
                                    var v278 uint64 = 1
                                    _ = v278
                                default:
                                    if v5 {
                                        v4 += (v4 * 591)
                                        v81 += v81
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            if ((v4 | v4) > (^(v4))) {
                                                fmt.Println(v5, v81, v69)
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if v5 {
                                                    switch "" {
                                                    case "\n\t":
                                                        if v5 {
                                                            if v5 {
                                                                v4 = (v4 << ((255) & 63))
                                                                var v279 uint32 = genFn1(65535, v4, ("foo bar" + ("\n\t" + "foo bar")))
                                                                _ = v279
                                                            } else {
                                                                v280 := GenS32{F0: 231, F1: 17, F2: uint8(v4)}
                                                                switch "a" {
                                                                case "foo bar":
                                                                    if v5 {
                                                                        v2 = append(v2, genFn1(42, (^(v4)), (("foo bar" + "") + "a")))
                                                                        if true {
                                                                            if false {
                                                                                v102 = (v81 >> ((2) & 63))
                                                                            } else {
                                                                                if true {
                                                                                    _ = len(v3)
                                                                                } else {
                                                                                    if v5 {
                                                                                        if (!false) {
                                                                                            v5 = (3.14159 <= v110)
                                                                                            v69 = (v64 * v64)
                                                                                            v110 = (v110 * v110)
                                                                                            fmt.Println(v102)
                                                                                        }
                                                                                    }
                                                                                    v281 := GenS33{F0: (128 <= v102), F1: 1}
                                                                                    if false {
                                                                                        v102 = fi101
                                                                                    }
                                                                                    _ = v281
                                                                                }
                                                                                v102 -= v102
                                                                                v4 = (^(v4))
                                                                            }
                                                                            v282 := GenS4{F0: 65535, F1: 2}
                                                                            var v283 int32 = 42
                                                                            _ = v282
                                                                            _ = v283
                                                                        } else {
                                                                            var v284 uint = 42
                                                                            _ = v284
                                                                        }
                                                                    }
                                                                case "a":
                                                                    if ((4294967295 == 2147483648) && v5) {
                                                                        var v285 int64 = 584
                                                                        if true {
                                                                            v4 = 225
                                                                            fmt.Println(v81, v81)
                                                                        } else {
                                                                            switch 0 {
                                                                            case 1:
                                                                                var v286 int32 = 42
                                                                                _ = v286
                                                                            case 677:
                                                                                var v287 uint8 = (^(v64))
                                                                                _ = v287
                                                                            case 969:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                if (true && v5) {
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        var v288 int8 = 109
                                                                                        _ = v288
                                                                                    }
                                                                                    v289 := GenS34{F0: v5}
                                                                                    _ = v289
                                                                                } else {
                                                                                    if v5 {
                                                                                        if (((float32(((0.1 / v110) - v110)) - (0.1 - 3.14159)) * (1e10 - 0.0)) >= -2.25) {
                                                                                            v102 = fi101
                                                                                        } else {
                                                                                            v102 -= v102
                                                                                            var v290 int32 = -652
                                                                                            v6 = append(v6, 4294967295)
                                                                                            _ = v290
                                                                                        }
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = (fi121 % (((v102 * v81)) | 1))
                                                                                        }
                                                                                    }
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 2
                                                                                    }
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "\n\t"
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (238 > -192) {
                                                                                v291 := GenS4{F0: 253, F1: v69}
                                                                                _ = v291
                                                                            } else {
                                                                                v292 := GenS35{F0: v285}
                                                                                switch (^(v102)) {
                                                                                case 0:
                                                                                    fmt.Println(v81)
                                                                                case -9223372036854775808:
                                                                                    var v293 int = (v81 % ((-274) | 1))
                                                                                    _ = v293
                                                                                case -2147483648:
                                                                                    v294 := GenS36{F0: -2147483648, F1: 2, F2: 0}
                                                                                    _ = v294
                                                                                }
                                                                                switch (((("hello" + ("a" + "\n\t")) + ("a" + "foo bar")) + ("a" + ("hello" + "a"))) + ("foo bar" + (((("a" + "foo bar") + "") + "foo bar") + ""))) {
                                                                                case "a":
                                                                                    v5 = ((!(false || (!v5))) && v5)
                                                                                case "foo bar":
                                                                                    v110 = 1e10
                                                                                }
                                                                                v285 = int64(v102)
                                                                                v102 += fi121
                                                                                _ = v292
                                                                            }
                                                                            v295 := GenS10{F0: "\n\t"}
                                                                            v296 := GenS2{F0: (v69 ^ v69)}
                                                                            var v297 string = ""
                                                                            _ = v295
                                                                            _ = v296
                                                                            _ = v297
                                                                        }
                                                                        var v298 int8 = 127
                                                                        _ = v285
                                                                        _ = v298
                                                                    }
                                                                case "hello":
                                                                    v2 = append(v2, 42)
                                                                default:
                                                                    v299 := GenS37{F0: "", F1: 0, F2: "\n\t"}
                                                                    _ = v299
                                                                }
                                                                _ = v280
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = genFn1(1, uint64(v64), "foo bar")
                                                            }
                                                            var v300 uint = genFn0(v64)
                                                            _ = v300
                                                        }
                                                    case "hello":
                                                        var v301 uint = uint(v69)
                                                        _ = v301
                                                    }
                                                    if v5 {
                                                        fmt.Println(v81, v64)
                                                    }
                                                    v302 := GenS13{F0: 32767, F1: uint16(v4), F2: 1000}
                                                    _ = v302
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v303 uint8 = 153
                                                        _ = v303
                                                    case "\n\t":
                                                        var v304 uint64 = 0
                                                        _ = v304
                                                    case "a":
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    default:
                                                        v69 = (v69 + v64)
                                                    }
                                                    if v5 {
                                                        v102 = (v81 + (-(fi121)))
                                                        var v305 int32 = 1
                                                        _ = v305
                                                    }
                                                    var v306 int8 = 0
                                                    if true {
                                                        switch (fi121 >> ((358) & 63)) {
                                                        case 2147483647:
                                                            v307 := GenS22{F0: v102, F1: 1}
                                                            _ = v307
                                                        case -632:
                                                            if ((-2147483648 == 889) || v5) {
                                                                v110 = v110
                                                                if (1 >= 26) {
                                                                    v308 := GenS28{F0: -2147483648}
                                                                    if v5 {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (fi121 * v81)
                                                                        }
                                                                    } else {
                                                                        var v309 bool = (!(699 >= 4294967295))
                                                                        _ = v309
                                                                    }
                                                                    _ = v308
                                                                } else {
                                                                    if v5 {
                                                                        v110 = v110
                                                                        v310 := GenS1{F0: (^(v306)), F1: ((("foo bar" + "a") + "foo bar") + "a")}
                                                                        v110 -= (v110 + (v110 / v110))
                                                                        _ = v310
                                                                    } else {
                                                                        fmt.Println(v81, v5)
                                                                        fmt.Println(v69, fi121, v102)
                                                                        v64 -= (^(v69))
                                                                    }
                                                                }
                                                                v102 += (-(fi101))
                                                                if false {
                                                                    if v5 {
                                                                        _ = len(v6)
                                                                    } else {
                                                                        if (false && v5) {
                                                                            v311 := GenS30{F0: (v4 + 14)}
                                                                            v312 := GenS38{F0: 384, F1: 18446744073709551615}
                                                                            _ = v311
                                                                            _ = v312
                                                                        } else {
                                                                            v5 = false
                                                                            var v313 int32 = int32(v69)
                                                                            var v314 bool = true
                                                                            _ = v313
                                                                            _ = v314
                                                                        }
                                                                        v110 += v110
                                                                        var v315 uint = 0
                                                                        _ = v315
                                                                    }
                                                                } else {
                                                                    v5 = ((v4 + 160) < 2)
                                                                    if ((v110 + float64(float32(((v110 * (v110 + (v110 + 3.14159))) / ((0.1 - 1.5) - 3.14159))))) == v110) {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        var v316 uint32 = 189
                                                                        var v317 uint16 = 1
                                                                        v318 := GenS20{F0: 835, F1: -2147483648, F2: v69}
                                                                        _ = v316
                                                                        _ = v317
                                                                        _ = v318
                                                                    }
                                                                }
                                                            } else {
                                                                if false {
                                                                    var v319 float32 = 0.0
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v4 = v4
                                                                    case "a":
                                                                        if v5 {
                                                                            v320 := GenS39{F0: 1}
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                fmt.Println(v102)
                                                                            }
                                                                            if v5 {
                                                                                var v321 int = (fi101 >> ((2) & 63))
                                                                                _ = v321
                                                                            } else {
                                                                                v5 = v5
                                                                                v102 -= (-(fi101))
                                                                                var v322 int8 = (v306 % ((v306) | 1))
                                                                                _ = v322
                                                                            }
                                                                            _ = v320
                                                                        }
                                                                    }
                                                                    _ = v319
                                                                } else {
                                                                    v323 := GenS19{F0: 32767}
                                                                    switch "a" {
                                                                    case "":
                                                                        v4 += (v4 ^ v4)
                                                                    default:
                                                                        if v5 {
                                                                            var v324 int16 = 1
                                                                            _ = v324
                                                                        }
                                                                    }
                                                                    _ = v323
                                                                }
                                                                var v325 uint8 = v64
                                                                switch ("\n\t" + "") {
                                                                case "":
                                                                    var v326 int = 365
                                                                    _ = v326
                                                                case "hello":
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v327 := GenS40{F0: 1, F1: (v325 - v69)}
                                                                        _ = v327
                                                                    case "hello":
                                                                        if true {
                                                                            switch fi101 {
                                                                            case 127:
                                                                                v328 := GenS41{F0: -704, F1: (^(v4)), F2: "\n\t"}
                                                                                _ = v328
                                                                            default:
                                                                                var v329 uint = 42
                                                                                _ = v329
                                                                            }
                                                                        } else {
                                                                            v69 = v325
                                                                            v4 = 462
                                                                            v6 = append(v6, 42)
                                                                            v330 := GenS42{F0: -32768}
                                                                            _ = v330
                                                                        }
                                                                    default:
                                                                        v331 := GenS43{F0: -128, F1: -1, F2: fi121}
                                                                        _ = v331
                                                                    }
                                                                }
                                                                _ = v325
                                                            }
                                                        default:
                                                            var v332 uint = 1
                                                            _ = v332
                                                        }
                                                        fmt.Println(v69, v5)
                                                    }
                                                    var v333 int = 273
                                                    switch ("foo bar" + "") {
                                                    case "foo bar":
                                                        var v334 int32 = 42
                                                        _ = v334
                                                    default:
                                                        var v335 uint = 0
                                                        _ = v335
                                                    }
                                                    _ = v306
                                                    _ = v333
                                                }
                                                v2 = append(v2, 2147483647)
                                            }
                                        case "":
                                            v64 += (v64 >> ((632) & 63))
                                        case "\n\t":
                                            v102 -= int(v69)
                                        }
                                        v110 = v110
                                        var v336 uint = 42
                                        _ = v336
                                    }
                                    v1 = append(v1, (fi121 >> ((1) & 63)))
                                }
                            }
                            _ = len(v6)
                            switch (-(v81)) {
                            case -842:
                                var v337 uint16 = 42
                                var v338 float32 = (1.5 - 0.0)
                                v5 = (9223372036854775807 > 643)
                                _ = v337
                                _ = v338
                            }
                            v339 := GenS23{F0: v102}
                            v69 = (v64 / (((v64 % ((uint8(v4)) | 1))) | 1))
                            var v340 uint64 = 42
                            _ = fi121
                            _ = v339
                            _ = v340
                        }
                    }
                    switch "" {
                    case "\n\t":
                        var v341 uint16 = 255
                        _ = v341
                    default:
                        v4 += (v4 & v4)
                    }
                    _ = v110
                }
                _ = v102
            } else {
                for fi342 := 0; fi342 < 11; fi342++ {
                    switch (v81 + (^(fi101))) {
                    case 4294967295:
                        if v5 {
                            fmt.Println(v5, v64)
                        } else {
                            fmt.Println(fi101, fi342, v4)
                        }
                    case 128:
                        if true {
                            var v343 uint64 = 390
                            fmt.Println(v343)
                            if len(v1) > 0 {
                                v1[0] = 32768
                            }
                            fmt.Println(v4, v81, v69)
                            switch ((("" + ("\n\t" + "foo bar")) + "") + "") {
                            case "hello":
                                v344 := GenS21{F0: 100, F1: 115, F2: genFn0((v69 - 127))}
                                _ = v344
                            case "foo bar":
                                v81 += (fi101 & 494)
                            }
                            _ = v343
                        } else {
                            switch "" {
                            case "hello":
                                fmt.Println(v4, v69, v5)
                            default:
                                v345 := GenS7{F0: 2147483647, F1: -128, F2: 1}
                                _ = v345
                            }
                            if (2 == 820) {
                                var v346 uint32 = 2
                                v4 += 1
                                fmt.Println(fi101)
                                switch 680 {
                                case 289:
                                    switch int(v4) {
                                    case 0:
                                        var v347 bool = (uint8(v346) >= v64)
                                        _ = v347
                                    case 1:
                                        if v5 {
                                            v81 = fi342
                                            if false {
                                                _ = len(v6)
                                            }
                                        } else {
                                            v348 := GenS26{F0: 1, F1: 947}
                                            _ = v348
                                        }
                                        var v349 uint = genFn0((v69 << ((1) & 63)))
                                        _ = v349
                                    default:
                                        var v350 uint32 = v346
                                        _ = v350
                                    }
                                }
                                _ = v346
                            }
                            fmt.Println(v81, v69, v81)
                        }
                    case -9223372036854775808:
                        if false {
                            var v351 string = ""
                            v352 := GenS10{F0: v351}
                            _ = v351
                            _ = v352
                        } else {
                            switch (v81 | fi101) {
                            case 9223372036854775807:
                                _ = len(v2)
                            case -206:
                                fmt.Println(v5, v5, v64)
                            case 32767:
                                if true {
                                    v69 = (v69 * v69)
                                    if (1.5 < 3.14159) {
                                        if false {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            fmt.Println(v5)
                                            var v353 int16 = 1000
                                            var v354 uint64 = 2
                                            _ = v353
                                            _ = v354
                                        }
                                        var v355 int64 = 0
                                        _ = v355
                                    }
                                    v356 := GenS21{F0: 42, F1: 0, F2: 732}
                                    _ = v356
                                }
                            default:
                                var v357 float64 = -2.25
                                _ = v357
                            }
                            if v5 {
                                if v5 {
                                    if ("a" != "hello") {
                                        v6 = append(v6, genFn1(496, (v4 + 0), "foo bar"))
                                    } else {
                                        var v358 bool = (42 == -1)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v4, fi101)
                                        _ = v358
                                    }
                                    var v359 string = (((("\n\t" + ("a" + "foo bar")) + "hello") + "") + "hello")
                                    switch fi342 {
                                    case -9223372036854775808:
                                        switch "hello" {
                                        case "\n\t":
                                            if (true && v5) {
                                                if (2 < 42) {
                                                    var v360 int64 = 9223372036854775807
                                                    _ = v360
                                                }
                                            } else {
                                                var v361 int32 = -960
                                                _ = v361
                                            }
                                        case "foo bar":
                                            if (!(712 < 1)) {
                                                _ = len(v6)
                                                v359 = ""
                                                var v362 uint32 = 2147483648
                                                _ = v362
                                            }
                                        case "hello":
                                            fmt.Println(v5, fi342, v4)
                                        default:
                                            v363 := GenS44{F0: 1}
                                            _ = v363
                                        }
                                    }
                                    _ = v359
                                } else {
                                    var v364 uint16 = 255
                                    _ = v364
                                }
                                if ((1000 >= 983) || v5) {
                                    var v365 uint64 = v4
                                    fmt.Println(fi101)
                                    v366 := GenS14{F0: 830, F1: (^(v64))}
                                    v81 = fi342
                                    v69 = (v64 & v64)
                                    _ = v365
                                    _ = v366
                                }
                            }
                            var v367 uint16 = 65535
                            v69 = v64
                            var v368 int16 = 1000
                            v69 += uint8(v367)
                            _ = v367
                            _ = v368
                        }
                    default:
                        if (!false) {
                            v5 = (((-2.25 * 0.0) - 0.0) >= (1e10 / -2.25))
                            v81 = 128
                            if v5 {
                                v2 = append(v2, 1)
                            }
                        } else {
                            var v369 uint = 2
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            _ = v369
                        }
                    }
                    fmt.Println(v64, v64, fi101)
                    switch "a" {
                    case "hello":
                        v81 = fi101
                    case "foo bar":
                        v3 = append(v3, "")
                    case "\n\t":
                        v69 = v64
                        v4 -= 329
                    default:
                        if v5 {
                            v69 -= uint8(v4)
                            _ = len(v2)
                            var v370 float64 = (0.0 - (3.14159 - 0.1))
                            v371 := GenS30{F0: (v4 ^ 4294967295)}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v370
                            _ = v371
                        }
                    }
                    _ = fi342
                }
                var v372 bool = (!("foo bar" > ""))
                var v373 int32 = 42
                _ = v372
                _ = v373
            }
            _ = fi101
        }
    }
    v374 := GenS29{F0: v69, F1: "hello"}
    var v375 int16 = 32767
    var v376 int16 = v375
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v64 -= 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v375:", v375)
    fmt.Println("v376:", v376)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    fmt.Println("v69:", v69)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v374
    _ = v375
    _ = v376
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v69
    _ = v70
    _ = v81
}
