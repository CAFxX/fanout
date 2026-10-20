package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int8; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint }
type GenS7 struct { F0 uint32; F1 int; F2 int64 }
type GenS8 struct { F0 int; F1 uint; F2 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 int64; F1 int8; F2 uint32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int32; F1 uint32 }
type GenS17 struct { F0 int32; F1 uint8 }
type GenS18 struct { F0 uint8; F1 int16 }
type GenS19 struct { F0 int; F1 int }
type GenS20 struct { F0 int; F1 uint64; F2 int8 }
type GenS21 struct { F0 uint; F1 int8 }
func genFn0(p0 uint32) uint {
    fmt.Println(p0)
    var v6 int16 = 1
    v7 := GenS1{F0: 1}
    var v8 bool = ((p0 % ((0) | 1)) < (p0 % ((p0) | 1)))
    fmt.Println(v8, v6)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 942
}
func genFn1(p0 uint, p1 int32, p2 uint16) int8 {
    switch ("" + "foo bar") {
    case "foo bar":
        if true {
            switch ("\n\t" + "hello") {
            case "a":
                v9 := []int{-218}
                _ = v9
            case "foo bar":
                switch -9223372036854775808 {
                case 0:
                    if ((p0 * p0) >= (p0 + (p0 >> ((p0) & 63)))) {
                        v10 := GenS2{F0: uint8(p2)}
                        fmt.Println(p1, p2, p1)
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            var v12 uint8 = 95
                            if (-2147483648 < (-(p1))) {
                                v13 := []int{fi11}
                                _ = v13
                            } else {
                                v14 := []int{2147483647}
                                _ = v14
                            }
                            switch "a" {
                            case "":
                                var v15 uint32 = 228
                                _ = v15
                            case "a":
                                p2 -= (p2 + 226)
                            default:
                                fmt.Println(fi11, p2, v12)
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                case 127:
                    p0 = (p0 << (((p0 % ((p0) | 1))) & 63))
                case 2147483647:
                    fmt.Println(p1)
                }
            default:
                p2 = 0
            }
        } else {
            var v16 bool = (0 != 885)
            _ = v16
        }
    case "":
        if true {
            var v17 int8 = int8(p1)
            var v18 float64 = 0.0
            _ = v17
            _ = v18
        }
    default:
        v19 := GenS2{F0: 6}
        _ = v19
    }
    var v20 int64 = -908
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 75
}

func main() {
    var v1 int16 = -233
    v2 := make(map[uint32]int)
    var v3 uint64 = 1
    var v4 uint32 = 42
    v5 := []uint32{v4, 4294967295, v4}
    if len(v5) > 0 {
        v5[0] = (^(v4))
    }
    switch "\n\t" {
    case "\n\t":
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS2{F0: 64}
            switch (fi21 + fi21) {
            case -1:
                fmt.Println(fi21, v3, v1)
            case 0:
                v23 := GenS1{F0: 887}
                _ = v23
            case 9223372036854775807:
                if true {
                    v1 += -5
                    if false {
                        v24 := v2[(v4 >> ((2) & 63))]
                        _ = v24
                    }
                }
            default:
                v3 = (v3 / ((v3) | 1))
            }
            _ = fi21
            _ = v22
        }
    case "":
        v3 = v3
    case "hello":
        v25 := GenS1{F0: 2147483647}
        _ = v25
    default:
        var v26 uint32 = v4
        _ = v26
    }
    var v27 string = ("hello" + ("\n\t" + "foo bar"))
    var v28 int = -698
    switch -1 {
    case 32768:
        var v29 float64 = -2.25
        _ = v29
    case 4294967295:
        for fi30 := 0; fi30 < 9; fi30++ {
            switch v27 {
            case "foo bar":
                v28 += (fi30 << ((1) & 63))
            case "hello":
                switch v27 {
                case "hello":
                    var v31 int = 661
                    _ = v31
                default:
                    switch "\n\t" {
                    case "":
                        var v32 uint64 = (^(v3))
                        _ = v32
                    default:
                        for fi33 := 0; fi33 < 12; fi33++ {
                            v5 = append(v5, (v4 << ((2) & 63)))
                            fmt.Println(v4)
                            _ = fi33
                        }
                    }
                }
            case "":
                v34 := GenS1{F0: -951}
                _ = v34
            }
            var v35 uint64 = (^(v3))
            switch (v28 + (fi30 - 4294967295)) {
            case 1:
                fmt.Println(v28, v27, v28)
            case 0:
                v4 = 65535
            case 32768:
                v3 -= 0
            default:
                v27 = v27
            }
            _ = fi30
            _ = v35
        }
    case 1:
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 uint8 = 1
            if false {
                v38 := GenS1{F0: 42}
                if len(v5) > 0 {
                    v5[0] = (v4 | (v4 % (((v4 << ((255) & 63))) | 1)))
                }
                v39 := GenS1{F0: -1}
                _ = v38
                _ = v39
            }
            fmt.Println(v37)
            _ = fi36
            _ = v37
        }
    }
    fmt.Println(v4)
    switch (v28 / ((v28) | 1)) {
    case -497:
        if len(v5) > 0 {
            v5[0] = v4
        }
    case 9223372036854775807:
        v1 = v1
    case -2147483648:
        if (v3 < v3) {
            var v40 string = "\n\t"
            var v41 uint8 = 186
            _ = v40
            _ = v41
        }
    }
    if true {
        v27 = ("hello" + ((("" + v27) + "\n\t") + "foo bar"))
        v1 = (v1 ^ (-(v1)))
    } else {
        if false {
            v42 := GenS3{F0: 255}
            if (!true) {
                for fi43 := 0; fi43 < 11; fi43++ {
                    var v44 int8 = 1
                    v44 += (v44 * v44)
                    for fi45 := 0; fi45 < 15; fi45++ {
                        var v46 uint8 = 216
                        v47 := GenS2{F0: v46}
                        fmt.Println(v3)
                        var v48 int32 = 2147483647
                        _ = fi45
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                    _ = fi43
                    _ = v44
                }
                v1 = (v1 % ((-66) | 1))
            }
            if len(v5) > 0 {
                v5[0] = uint32(v3)
            }
            _ = v42
        }
    }
    v49 := GenS3{F0: 255}
    switch "" {
    case "foo bar":
        fmt.Println(v27)
    case "a":
        v50 := GenS4{F0: 188, F1: -20, F2: (^(v3))}
        _ = v50
    case "":
        switch v28 {
        case 329:
            switch (v27 + (v27 + "foo bar")) {
            case "foo bar":
                var v51 float64 = float64((float32(1e10) + 3.14159))
                _ = v51
            default:
                v52 := v2[(v4 + 42)]
                _ = v52
            }
        case 32768:
            for fi53 := 0; fi53 < 10; fi53++ {
                for fi54 := 0; fi54 < 13; fi54++ {
                    v27 = v27
                    _ = fi54
                }
                switch v27 {
                case "hello":
                    if false {
                        _ = len(v5)
                        if ((3.14159 == 0.1) && false) {
                            switch -9223372036854775808 {
                            case 32768:
                                var v55 int = v28
                                _ = v55
                            case 2147483647:
                                if (((0.1 + ((1e10 * 1.5) * float64(float32(float64(float32((0.1 - float64(0.0)))))))) * (float64((((1.5 - 0.0) - -2.25) / (3.14159 + 3.14159))) + ((float64(0.0) / ((1.5 * 0.1) / 1.5)) * (0.1 - (0.1 + -2.25))))) > 0.0) {
                                    v56 := GenS3{F0: 128}
                                    v27 = ((v27 + "a") + v27)
                                    _ = v56
                                } else {
                                    v57 := GenS5{F0: (v1 + (^(v1))), F1: 18446744073709551615, F2: 538}
                                    _ = v57
                                }
                            default:
                                var v58 uint16 = 65535
                                _ = v58
                            }
                            v27 = ("" + ("foo bar" + "hello"))
                        }
                        var v59 float64 = -2.25
                        _ = v59
                    } else {
                        var v60 float32 = ((float32((((3.14159 - ((3.14159 + 3.14159) * (1.5 + 3.14159))) / 3.14159) / float64(1e10))) * 0.1) / 3.14159)
                        if (255 != 0) {
                            v61 := GenS1{F0: 1}
                            var v62 uint32 = 2
                            _ = v61
                            _ = v62
                        }
                        _ = v60
                    }
                case "foo bar":
                    var v63 string = (v27 + "hello")
                    _ = v63
                }
                _ = fi53
            }
        case 916:
            for fi64 := 0; fi64 < 11; fi64++ {
                var v65 int32 = 0
                var v66 string = (v27 + "\n\t")
                _ = fi64
                _ = v65
                _ = v66
            }
        }
    }
    var v67 bool = false
    v68 := GenS6{F0: (v3 % (((v3 >> ((2) & 63))) | 1)), F1: 42, F2: 0}
    fmt.Println(v28, v4)
    var v69 uint64 = 18446744073709551615
    var v70 uint32 = 0
    if (0.0 > 1.5) {
        fmt.Println(v1, v4, v3)
        fmt.Println(v69)
        _ = len(v2)
    }
    for fi71 := 0; fi71 < 3; fi71++ {
        fmt.Println(v4, v70)
        if v67 {
            switch v28 {
            case -432:
                if (!(598 != 1)) {
                    for fi72 := 0; fi72 < 12; fi72++ {
                        var v73 int = 2147483647
                        v74 := v2[4294967295]
                        v70 -= 0
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                    if false {
                        delete(v2, v70)
                        for fi75 := 0; fi75 < 1; fi75++ {
                            v76 := v2[(v70 ^ (v4 + 65535))]
                            _ = fi75
                            _ = v76
                        }
                        if true {
                            switch v27 {
                            case "hello":
                                var v77 int = v28
                                _ = v77
                            }
                        } else {
                            for fi78 := 0; fi78 < 6; fi78++ {
                                var v79 float32 = 0.1
                                _ = fi78
                                _ = v79
                            }
                            for fi80 := 0; fi80 < 11; fi80++ {
                                v70 -= v4
                                var v81 int64 = 984
                                _ = fi80
                                _ = v81
                            }
                        }
                    }
                } else {
                    v82 := GenS6{F0: (v3 % (((v69 % (((v3 | (v3 + v69))) | 1))) | 1)), F1: -381, F2: 38}
                    var v83 uint64 = 2
                    _ = v82
                    _ = v83
                }
            case 12:
                if v67 {
                    switch v27 {
                    case "a":
                        var v84 float64 = (0.0 * 3.14159)
                        _ = v84
                    case "\n\t":
                        switch (fi71 / (((^(fi71))) | 1)) {
                        case 32768:
                            fmt.Println(v27, v70, v67)
                        default:
                            var v85 string = v27
                            _ = v85
                        }
                    case "":
                        fmt.Println(v3, v1)
                    }
                }
            case -458:
                if (!true) {
                    v67 = v67
                    var v86 string = v27
                    _ = v86
                } else {
                    for fi87 := 0; fi87 < 4; fi87++ {
                        v88 := GenS7{F0: v70, F1: (v28 + (fi71 * v28)), F2: 42}
                        fmt.Println(v27, fi87)
                        fmt.Println(v70)
                        _ = len(v2)
                        v28 = fi87
                        v2[(^(v70))] = fi87
                        if v67 {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            if (genFn1(42, 0, 1) == 100) {
                                v27 = v27
                                if false {
                                    var v89 uint16 = 197
                                    v90 := GenS3{F0: 96}
                                    var v91 uint32 = (v4 >> ((234) & 63))
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                            } else {
                                if (2 <= 337) {
                                    var v92 int32 = 42
                                    v70 += v70
                                    _ = v92
                                } else {
                                    var v93 int16 = -1
                                    var v94 int8 = -95
                                    _ = v93
                                    _ = v94
                                }
                                v95 := GenS8{F0: 0, F1: 0, F2: 220}
                                _ = v95
                            }
                        }
                        delete(v2, uint32(v3))
                        v96 := GenS8{F0: fi71, F1: 924, F2: 0}
                        _ = fi87
                        _ = v88
                        _ = v96
                    }
                    v69 = v69
                }
            default:
                v97 := GenS8{F0: 2147483647, F1: 569, F2: 2}
                _ = v97
            }
            v27 = "\n\t"
            var v98 int64 = 9223372036854775807
            _ = v98
        } else {
            var v99 float32 = -2.25
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            switch (fi71 | v28) {
            case -433:
                switch v27 {
                case "a":
                    var v100 uint16 = 32767
                    _ = v100
                case "foo bar":
                    v1 = v1
                default:
                    if (v99 >= ((v99 / 0.1) + -2.25)) {
                        for fi101 := 0; fi101 < 8; fi101++ {
                            if (485 != 255) {
                                switch v27 {
                                case "":
                                    var v102 int32 = 2147483647
                                    _ = v102
                                case "hello":
                                    if (954 < -2147483648) {
                                        v103 := GenS6{F0: (v3 >> ((0) & 63)), F1: -130, F2: 18}
                                        switch v27 {
                                        case "hello":
                                            var v104 string = v27
                                            _ = v104
                                        case "foo bar":
                                            v99 = v99
                                        case "\n\t":
                                            var v105 float64 = (1e10 / float64(1e10))
                                            _ = v105
                                        }
                                        _ = v103
                                    }
                                case "a":
                                    var v106 uint = 2
                                    _ = v106
                                default:
                                    fmt.Println(fi101, fi101, v67)
                                }
                                v107 := GenS9{F0: (v28 | v28)}
                                fmt.Println(v4, v3, v70)
                                _ = v107
                            } else {
                                if (!((v4 | v4) <= 0)) {
                                    var v108 int16 = v1
                                    v109 := GenS2{F0: 128}
                                    _ = v108
                                    _ = v109
                                }
                                v1 -= v1
                                v110 := GenS1{F0: -1}
                                _ = v110
                            }
                            _ = fi101
                        }
                        v111 := GenS8{F0: v28, F1: 145, F2: 42}
                        if v67 {
                            if true {
                                _ = len(v5)
                                switch "hello" {
                                case "\n\t":
                                    v2[0] = fi71
                                case "":
                                    v112 := GenS4{F0: 42, F1: -36, F2: v3}
                                    _ = v112
                                }
                            }
                            var v113 uint = 42
                            _ = v113
                        }
                        _ = v111
                    } else {
                        if (v67 || (!false)) {
                            fmt.Println(v4)
                            for fi114 := 0; fi114 < 18; fi114++ {
                                var v115 int8 = 100
                                _ = fi114
                                _ = v115
                            }
                            v28 = (fi71 + fi71)
                        } else {
                            switch 586 {
                            case 127:
                                switch fi71 {
                                case -307:
                                    for fi116 := 0; fi116 < 4; fi116++ {
                                        switch v27 {
                                        case "":
                                            var v117 bool = true
                                            switch v27 {
                                            case "hello":
                                                v118 := GenS2{F0: uint8(v69)}
                                                _ = v118
                                            case "a":
                                                switch v27 {
                                                case "foo bar":
                                                    v119 := GenS10{F0: 0, F1: 42}
                                                    _ = v119
                                                case "hello":
                                                    v117 = true
                                                }
                                            case "foo bar":
                                                var v120 int16 = int16(v28)
                                                _ = v120
                                            default:
                                                fmt.Println(v27, v99)
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                v121 := GenS9{F0: (-(v28))}
                                                v117 = ((false || v117) && false)
                                                _ = v121
                                            case "a":
                                                if false {
                                                    v70 = v70
                                                } else {
                                                    v122 := GenS10{F0: 42, F1: -128}
                                                    switch (fi71 + fi116) {
                                                    case -39:
                                                        v123 := v2[(v4 % ((v70) | 1))]
                                                        _ = v123
                                                    }
                                                    _ = v122
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v124 bool = (!(v67 || v67))
                                                    _ = v124
                                                case "":
                                                    var v125 float32 = 1.5
                                                    _ = v125
                                                default:
                                                    v126 := GenS2{F0: 255}
                                                    _ = v126
                                                }
                                                if (!v117) {
                                                    switch 1 {
                                                    case -2147483648:
                                                        var v127 int32 = 486
                                                        _ = v127
                                                    default:
                                                        v5 = append(v5, 4294967295)
                                                    }
                                                }
                                                v128 := GenS11{F0: v28, F1: (v3 << ((255) & 63))}
                                                v28 = (fi116 >> ((uint(v69)) & 63))
                                                _ = v128
                                            case "":
                                                v117 = (4294967295 <= (^(v28)))
                                            default:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            _ = v117
                                        case "hello":
                                            v69 += (v69 * 584)
                                        case "\n\t":
                                            v129 := GenS7{F0: v4, F1: (fi116 ^ 42), F2: 42}
                                            _ = v129
                                        }
                                        if (false && (65535 != 2)) {
                                            fmt.Println(fi116, v70)
                                            if v67 {
                                                if v67 {
                                                    v130 := GenS7{F0: v70, F1: (fi71 ^ (^(fi116))), F2: -9223372036854775808}
                                                    _ = v130
                                                } else {
                                                    var v131 uint = genFn0(v4)
                                                    if (!v67) {
                                                        var v132 float64 = float64(0.0)
                                                        var v133 int8 = genFn1(uint(v70), -2147483648, 1)
                                                        _ = v132
                                                        _ = v133
                                                    } else {
                                                        v131 += (v131 & 1)
                                                        v134 := v2[(v70 >> ((v131) & 63))]
                                                        delete(v2, (v4 + uint32(v131)))
                                                        if v67 {
                                                            if ((0 != int64(v3)) || v67) {
                                                                switch ("a" + "") {
                                                                case "\n\t":
                                                                    v4 -= (v70 * v70)
                                                                default:
                                                                    v67 = v67
                                                                }
                                                                fmt.Println(v67, fi71)
                                                                if v67 {
                                                                    fmt.Println(v27, v134)
                                                                    v135 := GenS8{F0: -2147483648, F1: (v131 / ((genFn0(v4)) | 1)), F2: 224}
                                                                    v2[v70] = v134
                                                                    _ = v135
                                                                }
                                                            } else {
                                                                delete(v2, (^(v70)))
                                                                var v136 uint16 = 65535
                                                                fmt.Println(v1, v70)
                                                                if ((v1 > (v1 >> ((v131) & 63))) || true) {
                                                                    fmt.Println(v134, v3, v27)
                                                                    if false {
                                                                        if (((42 <= 103) || (v3 >= v69)) || (128 >= 2)) {
                                                                            var v137 bool = (!true)
                                                                            _ = v137
                                                                        }
                                                                        var v138 int64 = 143
                                                                        _ = v138
                                                                    } else {
                                                                        var v139 uint16 = 42
                                                                        var v140 int32 = 2147483647
                                                                        v141 := v2[(v70 * uint32(v3))]
                                                                        _ = v139
                                                                        _ = v140
                                                                        _ = v141
                                                                    }
                                                                    v142 := GenS12{F0: 981, F1: 42, F2: 2}
                                                                    _ = v142
                                                                } else {
                                                                    fmt.Println(v69, v28)
                                                                    v143 := GenS13{F0: 42}
                                                                    if v67 {
                                                                        v3 += v69
                                                                    }
                                                                    _ = v143
                                                                }
                                                                _ = v136
                                                            }
                                                            _ = len(v2)
                                                        } else {
                                                            v144 := GenS13{F0: 1}
                                                            if v67 {
                                                                fmt.Println(v1, v67, v67)
                                                            }
                                                            _ = v144
                                                        }
                                                        if v67 {
                                                            fmt.Println(fi116, v67)
                                                        }
                                                        _ = len(v5)
                                                        switch -1 {
                                                        case 32767:
                                                            v28 += 32767
                                                        default:
                                                            switch "" {
                                                            case "a":
                                                                var v145 string = v27
                                                                _ = v145
                                                            case "":
                                                                fmt.Println(v27, v70)
                                                            case "foo bar":
                                                                v146 := GenS14{F0: (v27 + v27)}
                                                                _ = v146
                                                            }
                                                        }
                                                        v147 := GenS15{F0: v131}
                                                        v5 = append(v5, uint32(v3))
                                                        _ = v134
                                                        _ = v147
                                                    }
                                                    _ = v131
                                                }
                                                var v148 int64 = int64(v70)
                                                v149 := GenS12{F0: v148, F1: 50, F2: (v70 * 1)}
                                                switch 893 {
                                                case 32767:
                                                    v67 = v67
                                                case 1:
                                                    if v67 {
                                                        var v150 int16 = (v1 / (((v1 * (v1 % (((^(v1))) | 1)))) | 1))
                                                        var v151 int16 = v150
                                                        var v152 uint64 = v3
                                                        var v153 float32 = (float32((float64((v99 - v99)) / 1e10)) / v99)
                                                        _ = v150
                                                        _ = v151
                                                        _ = v152
                                                        _ = v153
                                                    }
                                                default:
                                                    v67 = true
                                                }
                                                _ = v148
                                                _ = v149
                                            }
                                            v154 := GenS16{F0: 1, F1: (v70 + v70)}
                                            _ = v154
                                        }
                                        _ = fi116
                                    }
                                default:
                                    var v155 uint32 = v70
                                    _ = v155
                                }
                            case 2147483647:
                                switch v28 {
                                case 42:
                                    if v67 {
                                        _ = len(v5)
                                        if v67 {
                                            v99 -= ((v99 + 0.1) * 1e10)
                                            switch v27 {
                                            case "\n\t":
                                                var v156 int64 = 42
                                                _ = v156
                                            case "hello":
                                                fmt.Println(v99)
                                            case "":
                                                v157 := GenS17{F0: 42, F1: 2}
                                                _ = v157
                                            }
                                            _ = len(v2)
                                        } else {
                                            var v158 int16 = 120
                                            fmt.Println(v28)
                                            fmt.Println(v158, v158)
                                            _ = v158
                                        }
                                        var v159 bool = true
                                        _ = v159
                                    } else {
                                        for fi160 := 0; fi160 < 4; fi160++ {
                                            var v161 int = int(v70)
                                            v99 = v99
                                            var v162 uint16 = uint16(v69)
                                            _ = fi160
                                            _ = v161
                                            _ = v162
                                        }
                                        if (255 >= 42) {
                                            v5 = append(v5, uint32(v3))
                                            fmt.Println(v3, v3, v69)
                                        } else {
                                            if v67 {
                                                for fi163 := 0; fi163 < 18; fi163++ {
                                                    var v164 string = ""
                                                    _ = fi163
                                                    _ = v164
                                                }
                                                for fi165 := 0; fi165 < 12; fi165++ {
                                                    v166 := GenS3{F0: 155}
                                                    switch (v28 * (fi71 >> ((2) & 63))) {
                                                    case 128:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            switch v27 {
                                                            case "hello":
                                                                v4 += 2147483648
                                                            case "\n\t":
                                                                v3 -= 18446744073709551615
                                                            case "foo bar":
                                                                var v167 uint32 = (v70 % ((v70) | 1))
                                                                _ = v167
                                                            default:
                                                                var v168 uint16 = 1
                                                                if v67 {
                                                                    v69 += (v3 + (^(v69)))
                                                                    v99 = float32(((3.14159 / ((3.14159 * ((((float64((((v99 - 1.5) * 1e10) - v99)) + 1.5) * 0.1) + 1e10) / (-2.25 + -2.25))) + 3.14159)) - 0.0))
                                                                    v168 -= (v168 << ((2) & 63))
                                                                    fmt.Println(v69, v3, v70)
                                                                } else {
                                                                    v169 := GenS18{F0: 162, F1: v1}
                                                                    _ = v169
                                                                }
                                                                _ = v168
                                                            }
                                                        }
                                                    case -908:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    var v170 int8 = 67
                                                    _ = fi165
                                                    _ = v166
                                                    _ = v170
                                                }
                                            } else {
                                                var v171 int32 = -198
                                                _ = v171
                                            }
                                            v70 = v4
                                        }
                                    }
                                case 84:
                                    v3 -= 42
                                case -128:
                                    if (uint(v69) >= 30) {
                                        switch ((v27 + "hello") + (v27 + ("hello" + v27))) {
                                        case "hello":
                                            if true {
                                                if v67 {
                                                    if len(v5) > 0 {
                                                        v5[0] = v4
                                                    }
                                                    for fi172 := 0; fi172 < 8; fi172++ {
                                                        switch (v28 % (((fi71 ^ 32768)) | 1)) {
                                                        case 0:
                                                            var v173 uint32 = v4
                                                            _ = v173
                                                        case -9223372036854775808:
                                                            v174 := v2[(v4 << ((496) & 63))]
                                                            if len(v5) > 0 {
                                                                v5[0] = v70
                                                            }
                                                            fmt.Println(v67, v4, v99)
                                                            _ = v174
                                                        default:
                                                            v1 = v1
                                                            if ((209 != 1) && (v1 < 125)) {
                                                                var v175 string = v27
                                                                _ = v175
                                                            }
                                                        }
                                                        v176 := GenS13{F0: uint16(v4)}
                                                        var v177 uint = 0
                                                        _ = fi172
                                                        _ = v176
                                                        _ = v177
                                                    }
                                                    v4 -= 0
                                                } else {
                                                    v3 = v3
                                                    v69 -= v3
                                                }
                                                var v178 int16 = (v1 - 42)
                                                _ = v178
                                            }
                                        default:
                                            if ((!(uint(v4) == 1)) || true) {
                                                if (0.1 == 3.14159) {
                                                    for fi179 := 0; fi179 < 15; fi179++ {
                                                        switch (v27 + ((v27 + "\n\t") + "foo bar")) {
                                                        case "foo bar":
                                                            var v180 int = v28
                                                            _ = v180
                                                        case "a":
                                                            var v181 uint32 = v70
                                                            _ = v181
                                                        case "":
                                                            v182 := GenS9{F0: fi179}
                                                            _ = v182
                                                        }
                                                        _ = fi179
                                                    }
                                                    var v183 uint16 = 42
                                                    var v184 int16 = -419
                                                    _ = v183
                                                    _ = v184
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v185 := GenS1{F0: -404}
                                                _ = v185
                                            } else {
                                                var v186 int32 = 656
                                                if ((((v27 + (v27 + v27)) + ("a" + "")) + (v27 + "a")) < (v27 + (v27 + v27))) {
                                                    for fi187 := 0; fi187 < 14; fi187++ {
                                                        _ = len(v2)
                                                        if v67 {
                                                            v27 = v27
                                                            v188 := GenS19{F0: -752, F1: -2147483648}
                                                            var v189 uint = 0
                                                            _ = v188
                                                            _ = v189
                                                        } else {
                                                            if (4294967295 >= 42) {
                                                                v27 = ("\n\t" + ("hello" + v27))
                                                                var v190 int32 = (v186 * 2147483647)
                                                                var v191 float32 = (float32((float64(3.14159) / 1e10)) / -2.25)
                                                                _ = v190
                                                                _ = v191
                                                            } else {
                                                                v192 := GenS18{F0: 2, F1: -260}
                                                                _ = v192
                                                            }
                                                            v70 = (v70 * 42)
                                                            if v67 {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            } else {
                                                                var v193 int16 = v1
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v28 = (v28 ^ v28)
                                                                _ = v193
                                                            }
                                                        }
                                                        var v194 int8 = 42
                                                        if v67 {
                                                            var v195 string = "a"
                                                            v5 = append(v5, 115)
                                                            _ = v195
                                                        }
                                                        _ = fi187
                                                        _ = v194
                                                    }
                                                    v1 = (-(v1))
                                                } else {
                                                    switch (fi71 ^ fi71) {
                                                    case 1:
                                                        v196 := GenS20{F0: (v28 / ((42) | 1)), F1: (v3 | v69), F2: -74}
                                                        _ = v196
                                                    case 128:
                                                        if v67 {
                                                            var v197 int64 = -615
                                                            v198 := GenS18{F0: 1, F1: 1000}
                                                            if v67 {
                                                                var v199 uint64 = v3
                                                                _ = v199
                                                            } else {
                                                                if len(v5) > 0 {
                                                                    v5[0] = (v4 * v4)
                                                                }
                                                            }
                                                            _ = v197
                                                            _ = v198
                                                        }
                                                    }
                                                    fmt.Println(v99, v67, v69)
                                                    switch (v27 + ("" + "hello")) {
                                                    case "hello":
                                                        if v67 {
                                                            v4 = 4294967295
                                                        }
                                                    }
                                                }
                                                _ = v186
                                            }
                                        }
                                    }
                                }
                            case 128:
                                var v200 int64 = -420
                                _ = v200
                            default:
                                v201 := GenS13{F0: 395}
                                _ = v201
                            }
                            v3 = (v69 | v3)
                            v3 += v3
                        }
                        switch v27 {
                        case "a":
                            fmt.Println(v99, v70)
                        case "\n\t":
                            v67 = v67
                        case "foo bar":
                            for fi202 := 0; fi202 < 8; fi202++ {
                                v203 := v2[(^(v4))]
                                _ = fi202
                                _ = v203
                            }
                        }
                        v204 := GenS10{F0: uint16(v4), F1: 79}
                        _ = v204
                    }
                }
            case -791:
                if v67 {
                    var v205 uint32 = v70
                    v206 := GenS21{F0: 0, F1: 127}
                    var v207 uint16 = 1
                    _ = v205
                    _ = v206
                    _ = v207
                }
            case -2147483648:
                fmt.Println(fi71, v70, v27)
            default:
                switch (^(v28)) {
                case -391:
                    var v208 int32 = -2147483648
                    _ = v208
                case 2147483647:
                    for fi209 := 0; fi209 < 14; fi209++ {
                        var v210 int64 = -9223372036854775808
                        v1 = 1000
                        _ = fi209
                        _ = v210
                    }
                case 32767:
                    for fi211 := 0; fi211 < 6; fi211++ {
                        v5 = append(v5, v70)
                        _ = fi211
                    }
                default:
                    var v212 int = (fi71 << ((4294967295) & 63))
                    _ = v212
                }
            }
            _ = v99
        }
        _ = fi71
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
