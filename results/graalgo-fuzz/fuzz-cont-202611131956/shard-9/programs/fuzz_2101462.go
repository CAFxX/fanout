package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS9 struct { F0 int8; F1 uint8; F2 bool }
type GenS10 struct { F0 string; F1 int8 }
func genFn0(p0 int16, p1 int64, p2 int) int {
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(p2)
        v6 := []int{(-(fi5)), fi5}
        var v7 string = ""
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 205
}
func genFn1(p0 uint, p1 string) uint16 {
    p0 = (p0 << ((p0) & 63))
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint16 = 42
        v11 := []int{fi9, (^(fi9))}
        var v12 int16 = 407
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{2147483647}
    switch -2147483648 {
    case 4294967295:
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint32 = 42
            switch "hello" {
            case "\n\t":
                var v16 uint64 = 2
                _ = v16
            case "a":
                switch "\n\t" {
                case "hello":
                    var v17 int16 = 32767
                    _ = v17
                case "a":
                    p0 -= p0
                }
            }
            v18 := GenS1{F0: 65535, F1: -1}
            _ = fi14
            _ = v15
            _ = v18
        }
    default:
        p0 = 255
    }
    v19 := GenS2{F0: -9223372036854775808, F1: 226}
    _ = p0
    _ = p1
    _ = v13
    _ = v19
    return 32767
}

func main() {
    v1 := []uint32{859, 2}
    v2 := []int{-9223372036854775808}
    var v3 uint32 = 2147483647
    var v4 int64 = 42
    v3 = v3
    var v20 int32 = 1
    fmt.Println(v4, v3, v3)
    if (42 > genFn0(-389, int64(v20), int(v3))) {
        v20 = (v20 << ((18446744073709551615) & 63))
        v2 = append(v2, -1)
    }
    switch "\n\t" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = 868
        }
    case "":
        v21 := GenS1{F0: 189, F1: 0}
        _ = v21
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = v3
        }
    default:
        for fi22 := 0; fi22 < 16; fi22++ {
            switch "a" {
            case "hello":
                var v23 string = ("hello" + "foo bar")
                _ = v23
            case "a":
                fmt.Println(fi22, v4)
            default:
                v20 = v20
            }
            var v24 bool = false
            var v25 int16 = int16(v4)
            _ = fi22
            _ = v24
            _ = v25
        }
    }
    v20 = (v20 - v20)
    switch 0 {
    case -9223372036854775808:
        if ((18446744073709551615 <= 1) && false) {
            if true {
                v4 = v4
            } else {
                v26 := GenS2{F0: v4, F1: 128}
                _ = v26
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                _ = len(v1)
                if (!(fi27 == genFn0(int16(v20), (v4 | 1), fi27))) {
                    var v28 string = "foo bar"
                    _ = v28
                } else {
                    var v29 uint16 = 2
                    var v30 int32 = -1
                    for fi31 := 0; fi31 < 11; fi31++ {
                        switch -1 {
                        case 1:
                            var v32 float32 = 0.1
                            var v33 uint64 = 0
                            var v34 float32 = 1.5
                            _ = v32
                            _ = v33
                            _ = v34
                        case 127:
                            if true {
                                switch ("hello" + "hello") {
                                case "":
                                    var v35 uint64 = 42
                                    _ = v35
                                case "a":
                                    var v36 bool = true
                                    _ = v36
                                default:
                                    _ = len(v1)
                                }
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    v30 += (v20 % (((v30 | (v30 + v20))) | 1))
                                case "foo bar":
                                    var v37 bool = false
                                    _ = v37
                                }
                                if len(v2) > 0 {
                                    v2[0] = fi31
                                }
                            }
                        case 1000:
                            var v38 uint16 = v29
                            _ = v38
                        }
                        _ = fi31
                    }
                    _ = v29
                    _ = v30
                }
                var v39 float32 = 3.14159
                _ = fi27
                _ = v39
            }
            if false {
                var v40 int8 = -62
                if (v4 < (v4 & v4)) {
                    if false {
                        v41 := GenS3{F0: 32767, F1: ("hello" + ""), F2: 1}
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 15; fi42++ {
                        switch "" {
                        case "hello":
                            var v43 uint = 1
                            _ = v43
                        case "foo bar":
                            var v44 int64 = (v4 - (v4 / ((v4) | 1)))
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 4; fi45++ {
                            if ((!(!(v20 > 2147483647))) && ((v4 + (v4 / ((0) | 1))) == v4)) {
                                if false {
                                    if (42 <= 1) {
                                        v46 := GenS3{F0: 1000, F1: "a", F2: 42}
                                        _ = v46
                                    } else {
                                        if false {
                                            v2 = append(v2, 1)
                                        }
                                        fmt.Println(fi45)
                                        fmt.Println(v40, v20, v4)
                                    }
                                }
                                var v47 int64 = v4
                                fmt.Println(v4)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch "" {
                                case "\n\t":
                                    _ = len(v2)
                                    var v48 float32 = 0.1
                                    switch (("foo bar" + ("" + "a")) + "hello") {
                                    case "foo bar":
                                        var v49 int64 = (v47 >> ((65535) & 63))
                                        _ = v49
                                    }
                                    _ = v48
                                case "":
                                    if false {
                                        var v50 uint32 = (v3 % ((295) | 1))
                                        fmt.Println(v40, v3)
                                        if (("a" + "foo bar") < "a") {
                                            var v51 int32 = v20
                                            _ = len(v2)
                                            var v52 int64 = (^(v47))
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v50
                                    } else {
                                        v47 = (v4 >> ((255) & 63))
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v53 := GenS2{F0: v47, F1: 255}
                                        _ = v53
                                    }
                                default:
                                    v54 := GenS1{F0: 811, F1: 0}
                                    _ = v54
                                }
                                _ = v47
                            }
                            v3 -= (^(v3))
                            _ = fi45
                        }
                        _ = fi42
                    }
                }
                if (((float64(0.1) / ((float64(float32(3.14159)) - -2.25) + float64(float32((0.0 / float64(1.5)))))) + (1e10 + float64(float32(0.0)))) == 0.0) {
                    var v55 int32 = (v20 >> ((634) & 63))
                    for fi56 := 0; fi56 < 17; fi56++ {
                        v55 = v55
                        _ = fi56
                    }
                    _ = v55
                } else {
                    v40 -= 0
                    v1 = append(v1, (v3 / ((42) | 1)))
                    for fi57 := 0; fi57 < 14; fi57++ {
                        if len(v1) > 0 {
                            v1[0] = (v3 << ((9) & 63))
                        }
                        v58 := GenS4{F0: 113, F1: 0}
                        if true {
                            v59 := GenS2{F0: v4, F1: 42}
                            v20 = v20
                            var v60 string = "foo bar"
                            _ = v59
                            _ = v60
                        } else {
                            for fi61 := 0; fi61 < 5; fi61++ {
                                v62 := GenS4{F0: 900, F1: 42}
                                if (true || false) {
                                    v4 -= -117
                                    v4 -= (v4 >> ((482) & 63))
                                } else {
                                    if false {
                                        if (2 < 0) {
                                            v20 += v20
                                            var v63 bool = (v3 > 2147483648)
                                            _ = v63
                                        }
                                        v4 = v4
                                    } else {
                                        var v64 int16 = -1
                                        var v65 int32 = v20
                                        _ = v64
                                        _ = v65
                                    }
                                    fmt.Println(v3, v40, v4)
                                    switch "\n\t" {
                                    case "a":
                                        v20 += (v20 ^ (v20 - v20))
                                    }
                                }
                                var v66 int8 = (v40 / ((v40) | 1))
                                _ = fi61
                                _ = v62
                                _ = v66
                            }
                            v67 := GenS2{F0: (v4 / ((int64(v20)) | 1)), F1: 128}
                            var v68 int = (^(fi57))
                            _ = v67
                            _ = v68
                        }
                        _ = fi57
                        _ = v58
                    }
                }
                _ = v40
            } else {
                v69 := GenS2{F0: (v4 << ((2) & 63)), F1: 102}
                _ = v69
            }
        } else {
            fmt.Println(v20, v20)
            v70 := GenS1{F0: 2, F1: -1}
            _ = v70
        }
    case -1:
        var v71 uint8 = 10
        _ = v71
    }
    if (true || true) {
        v4 -= 1
        v1 = append(v1, (v3 / ((v3) | 1)))
    } else {
        v72 := GenS4{F0: 473, F1: int32(v4)}
        _ = v72
    }
    var v73 string = (("" + ("a" + "")) + "hello")
    switch v73 {
    case "":
        var v74 int32 = int32(v4)
        _ = v74
    case "hello":
        var v75 int = 42
        _ = v75
    default:
        switch -9223372036854775808 {
        case 42:
            v76 := GenS4{F0: 887, F1: (v20 % ((-2147483648) | 1))}
            _ = v76
        case 128:
            var v77 float64 = 0.1
            _ = v77
        case -1:
            if ((-(v20)) >= (v20 * 766)) {
                fmt.Println(v73, v3, v4)
                if ((((-2.25 * 3.14159) / 1.0) / -2.25) >= float64((3.14159 + 3.14159))) {
                    var v78 bool = true
                    for fi79 := 0; fi79 < 13; fi79++ {
                        v1 = append(v1, v3)
                        if ((v20 & (^(v20))) < v20) {
                            v1 = append(v1, 215)
                            if v78 {
                                var v80 int64 = -1
                                _ = v80
                            } else {
                                switch "" {
                                case "hello":
                                    if true {
                                        var v81 string = "foo bar"
                                        v82 := GenS5{F0: 4294967295, F1: v78, F2: 397}
                                        if (v20 <= (-(v20))) {
                                            switch fi79 {
                                            case 0:
                                                for fi83 := 0; fi83 < 11; fi83++ {
                                                    var v84 uint8 = 1
                                                    _ = fi83
                                                    _ = v84
                                                }
                                            case 2147483647:
                                                v85 := GenS6{F0: (v3 + (v3 & (v3 % (((v3 & v3)) | 1)))), F1: true}
                                                _ = v85
                                            case 42:
                                                v73 = v73
                                            default:
                                                v73 = v73
                                            }
                                            if (((1 < 196) && true) && v78) {
                                                for fi86 := 0; fi86 < 16; fi86++ {
                                                    if v78 {
                                                        var v87 int8 = -67
                                                        v4 += v4
                                                        _ = len(v2)
                                                        _ = v87
                                                    }
                                                    if v78 {
                                                        var v88 uint = 4294967295
                                                        switch -317 {
                                                        case 239:
                                                            var v89 bool = false
                                                            _ = v89
                                                        case 1:
                                                            v90 := GenS6{F0: v3, F1: v78}
                                                            v73 = "a"
                                                            _ = v90
                                                        case 32767:
                                                            v91 := GenS6{F0: 0, F1: (255 != genFn1(v88, (v81 + ((v81 + (v81 + ("\n\t" + v81))) + "a"))))}
                                                            _ = v91
                                                        }
                                                        switch -67 {
                                                        case 128:
                                                            var v92 uint64 = 4294967295
                                                            _ = v92
                                                        case -1:
                                                            v93 := GenS1{F0: 32767, F1: 42}
                                                            _ = v93
                                                        case 9223372036854775807:
                                                            switch fi79 {
                                                            case 4294967295:
                                                                if ((fi86 == (fi79 % (((fi86 % ((fi86) | 1))) | 1))) && (1000 > 1000)) {
                                                                    v94 := GenS7{F0: -1, F1: 0}
                                                                    switch ("a" + v73) {
                                                                    case "\n\t":
                                                                        v78 = (65535 < (v3 % ((v3) | 1)))
                                                                    case "a":
                                                                        v20 = 2147483647
                                                                    case "":
                                                                        v78 = false
                                                                        v4 = v4
                                                                    }
                                                                    _ = v94
                                                                }
                                                            }
                                                        }
                                                        _ = v88
                                                    }
                                                    if (602 >= 1) {
                                                        if ((0.0 + ((3.14159 + (1e10 - (3.14159 + -2.25))) / 1.5)) == (-2.25 / 1.5)) {
                                                            if len(v2) > 0 {
                                                                v2[0] = fi79
                                                            }
                                                            v3 -= v3
                                                        } else {
                                                            v78 = ((fi86 | (fi86 << ((18446744073709551615) & 63))) != (fi86 % (((fi86 / (((^(fi86))) | 1))) | 1)))
                                                            if (2 >= 18446744073709551615) {
                                                                v1 = append(v1, 2147483648)
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi86
                                                                }
                                                                if (!(0 == 42)) {
                                                                    var v95 int16 = -641
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    var v96 uint64 = 2
                                                                    _ = v95
                                                                    _ = v96
                                                                } else {
                                                                    var v97 float32 = 1e10
                                                                    var v98 int64 = (v4 << ((18446744073709551615) & 63))
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = (^(v3))
                                                                }
                                                            }
                                                        }
                                                        var v99 int64 = (v4 % ((v4) | 1))
                                                        _ = v99
                                                    }
                                                    _ = fi86
                                                }
                                                v2 = append(v2, genFn0(-865, -1, 4294967295))
                                                switch v81 {
                                                case "":
                                                    var v100 int32 = 2147483647
                                                    _ = v100
                                                }
                                            }
                                            if (42 != 1) {
                                                if (int16(v4) == 44) {
                                                    var v101 uint8 = 1
                                                    _ = v101
                                                }
                                                switch (fi79 >> ((255) & 63)) {
                                                case 32768:
                                                    var v102 uint32 = (^(v3))
                                                    _ = v102
                                                default:
                                                    var v103 bool = v78
                                                    _ = v103
                                                }
                                                fmt.Println(v4, v73)
                                            } else {
                                                fmt.Println(v73, v3, v20)
                                                v104 := GenS8{F0: v78, F1: 414, F2: 569}
                                                _ = v104
                                            }
                                        } else {
                                            if (!v78) {
                                                var v105 uint8 = 127
                                                _ = v105
                                            }
                                        }
                                        _ = v81
                                        _ = v82
                                    } else {
                                        fmt.Println(v73)
                                    }
                                case "":
                                    var v106 int32 = (-(v20))
                                    _ = v106
                                case "a":
                                    if v78 {
                                        switch fi79 {
                                        case 1:
                                            var v107 int16 = int16(v20)
                                            _ = v107
                                        case 597:
                                            var v108 uint = 0
                                            _ = v108
                                        case 587:
                                            for fi109 := 0; fi109 < 5; fi109++ {
                                                switch v73 {
                                                case "\n\t":
                                                    var v110 int64 = (v4 << ((42) & 63))
                                                    _ = v110
                                                case "a":
                                                    v78 = (!("a" >= (v73 + v73)))
                                                    var v111 int64 = v4
                                                    _ = v111
                                                case "":
                                                    switch genFn0(-1, (v4 / ((v4) | 1)), 274) {
                                                    case -2147483648:
                                                        fmt.Println(v73, v73)
                                                    case -9223372036854775808:
                                                        v112 := GenS3{F0: 1, F1: (((v73 + (v73 + v73)) + "foo bar") + v73), F2: 65535}
                                                        _ = v112
                                                    }
                                                default:
                                                    var v113 uint16 = 65535
                                                    _ = v113
                                                }
                                                var v114 float32 = (0.1 - 3.14159)
                                                _ = fi109
                                                _ = v114
                                            }
                                        default:
                                            fmt.Println(v78, fi79)
                                        }
                                        for fi115 := 0; fi115 < 20; fi115++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if (2 > 2) {
                                                var v116 string = ((v73 + "\n\t") + v73)
                                                var v117 float64 = float64(0.0)
                                                _ = v116
                                                _ = v117
                                            }
                                            _ = fi115
                                        }
                                        var v118 bool = v78
                                        _ = v118
                                    }
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            switch v73 {
                            case "foo bar":
                                v119 := GenS5{F0: v3, F1: (65535 == 18446744073709551615), F2: 65535}
                                _ = v119
                            case "hello":
                                fmt.Println(v78, v78, v78)
                            case "\n\t":
                                switch v73 {
                                case "":
                                    if v78 {
                                        var v120 float32 = 1.5
                                        v73 = v73
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v120
                                    } else {
                                        switch (v73 + v73) {
                                        case "hello":
                                            var v121 bool = (int16(v20) <= 1000)
                                            _ = v121
                                        case "":
                                            v122 := GenS9{F0: -1, F1: 255, F2: false}
                                            _ = v122
                                        }
                                        fmt.Println(v20)
                                        var v123 int16 = -365
                                        _ = v123
                                    }
                                case "\n\t":
                                    fmt.Println(v3)
                                }
                            default:
                                if (v4 >= v4) {
                                    for fi124 := 0; fi124 < 18; fi124++ {
                                        if true {
                                            v20 += -1
                                            v73 = v73
                                        }
                                        var v125 int16 = -779
                                        var v126 uint8 = 127
                                        _ = fi124
                                        _ = v125
                                        _ = v126
                                    }
                                    fmt.Println(v4)
                                }
                            }
                        }
                        _ = fi79
                    }
                    v78 = (0.0 < -2.25)
                    _ = v78
                }
                v3 = v3
            } else {
                _ = len(v2)
                var v127 uint = 336
                fmt.Println(v73)
                _ = v127
            }
        }
    }
    if true {
        if len(v1) > 0 {
            v1[0] = 42
        }
    } else {
        var v128 float32 = 0.1
        v73 = (v73 + v73)
        _ = v128
    }
    var v129 uint = 363
    for fi130 := 0; fi130 < 7; fi130++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi130
    }
    switch genFn0(-1, v4, -2147483648) {
    case 128:
        v131 := GenS7{F0: 190, F1: 42}
        _ = v131
    case -1:
        var v132 int8 = 0
        _ = v132
    case 4294967295:
        _ = len(v1)
    default:
        _ = len(v1)
    }
    fmt.Println(v73, v129, v4)
    switch v73 {
    case "a":
        fmt.Println(v20, v3)
    case "":
        var v133 float64 = 1e10
        _ = v133
    case "\n\t":
        v134 := GenS10{F0: "", F1: -105}
        _ = v134
    }
    var v135 int = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v129:", v129)
    fmt.Println("v135:", v135)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v129
    _ = v135
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v73
}
