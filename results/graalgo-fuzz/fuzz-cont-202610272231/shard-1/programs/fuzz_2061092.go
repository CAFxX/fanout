package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int16; F2 int16 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 uint8; F1 string; F2 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int64; F1 int8; F2 int64 }
type GenS14 struct { F0 int8; F1 uint32; F2 bool }
func genFn0(p0 int32) bool {
    if (2 != 1) {
        v6 := GenS1{F0: 162, F1: 100, F2: -88}
        _ = v6
    } else {
        var v7 bool = true
        if ((("\n\t" + ("foo bar" + "\n\t")) + "a") != "a") {
            p0 -= p0
        } else {
            switch ("a" + ("foo bar" + "\n\t")) {
            case "hello":
                for fi8 := 0; fi8 < 1; fi8++ {
                    var v9 uint16 = 2
                    v9 = (^(v9))
                    _ = fi8
                    _ = v9
                }
            case "":
                var v10 int = 933
                _ = v10
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                if (0.0 <= -2.25) {
                    fmt.Println(v7)
                    switch ("" + "hello") {
                    case "hello":
                        v12 := GenS1{F0: 127, F1: 100, F2: 62}
                        _ = v12
                    case "a":
                        fmt.Println(fi11, p0, fi11)
                    case "\n\t":
                        v13 := GenS1{F0: 42, F1: -1, F2: -1}
                        _ = v13
                    default:
                        p0 = p0
                    }
                } else {
                    if (-312 != fi11) {
                        var v14 int64 = 9223372036854775807
                        if v7 {
                            v15 := []int{}
                            for fi16 := 0; fi16 < 1; fi16++ {
                                fmt.Println(v7)
                                var v17 bool = v7
                                _ = fi16
                                _ = v17
                            }
                            _ = v15
                        } else {
                            var v18 uint = 42
                            _ = v18
                        }
                        _ = v14
                    }
                    if (128 != 128) {
                        v19 := []int{}
                        _ = v19
                    }
                }
                fmt.Println(v7, fi11, fi11)
                _ = fi11
            }
        }
        _ = v7
    }
    var v20 bool = (16 != 987)
    p0 += p0
    _ = p0
    _ = v20
    return ((605 >= 1) || false)
}
func genFn1(p0 uint8) uint8 {
    v21 := []int{1, 231}
    _ = len(v21)
    if len(v21) > 0 {
        v21[0] = -810
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        if ((!(2 != 96)) || (32767 > 32767)) {
            if (1e10 == 0.0) {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v24 := GenS1{F0: p0, F1: 42, F2: -39}
                    p0 = p0
                    _ = fi23
                    _ = v24
                }
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            } else {
                v25 := GenS1{F0: p0, F1: 0, F2: -114}
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 int64 = 491
                    var v28 uint16 = 65535
                    var v29 float32 = 1.5
                    p0 += (p0 & 64)
                    v21 = append(v21, (fi26 / (((-(fi26))) | 1)))
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = v25
            }
        }
        _ = fi22
    }
    p0 -= p0
    _ = p0
    _ = v21
    return 2
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]int)
    var v3 uint8 = 0
    v4 := make(map[int]string)
    var v5 string = ""
    var v30 int8 = 90
    v3 += v3
    var v31 int = 32768
    fmt.Println(v30, v5, v3)
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 int = (v31 / ((v31) | 1))
        if true {
            var v34 int32 = -1
            _ = v34
        } else {
            var v35 uint = 18446744073709551615
            if (!(("a" + "") == (v5 + "hello"))) {
                for fi36 := 0; fi36 < 12; fi36++ {
                    v2[4294967295] = 2147483647
                    var v37 uint32 = 2
                    switch "hello" {
                    case "a":
                        v38 := GenS1{F0: (v3 - 255), F1: (v30 ^ int8(v1)), F2: v30}
                        var v39 float64 = (-2.25 * 3.14159)
                        switch -2147483648 {
                        case 127:
                            v40 := GenS1{F0: v3, F1: 127, F2: v30}
                            _ = v40
                        case 4294967295:
                            v31 -= (fi36 ^ 507)
                        }
                        _ = v38
                        _ = v39
                    }
                    _ = fi36
                    _ = v37
                }
                v41 := GenS2{F0: (^(v3)), F1: 127}
                _ = v41
            }
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        if ((-(v30)) < (v30 ^ 0)) {
            v43 := GenS2{F0: genFn1((v3 % ((172) | 1))), F1: v30}
            if true {
                for fi44 := 0; fi44 < 14; fi44++ {
                    var v45 int8 = v30
                    fmt.Println(fi44)
                    var v46 uint16 = v1
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            var v47 int8 = (v30 << ((65535) & 63))
            _ = v43
            _ = v47
        } else {
            v30 = -128
            if false {
                var v48 int64 = 783
                _ = v48
            } else {
                for fi49 := 0; fi49 < 3; fi49++ {
                    v50 := GenS2{F0: 127, F1: -79}
                    _ = fi49
                    _ = v50
                }
                if true {
                    if (!(1 < (v1 | v1))) {
                        v1 = 408
                        switch v5 {
                        case "":
                            var v51 uint32 = 2147483648
                            _ = v51
                        case "hello":
                            var v52 float32 = (1e10 / (0.1 * (3.14159 / (3.14159 - 1e10))))
                            _ = v52
                        case "\n\t":
                            var v53 uint64 = 2
                            _ = v53
                        default:
                            switch (-(v31)) {
                            case 2147483647:
                                delete(v2, 65535)
                            case -2147483648:
                                var v54 int8 = (v30 << ((255) & 63))
                                _ = v54
                            default:
                                v55 := v4[v31]
                                _ = v55
                            }
                        }
                    } else {
                        for fi56 := 0; fi56 < 12; fi56++ {
                            v57 := []int{2147483647}
                            _ = fi56
                            _ = v57
                        }
                        v31 = -2147483648
                        if (985 >= 65535) {
                            for fi58 := 0; fi58 < 16; fi58++ {
                                var v59 uint32 = 65535
                                switch ("hello" + (v5 + ((("\n\t" + (("" + v5) + "foo bar")) + "a") + "\n\t"))) {
                                case "a":
                                    fmt.Println(v1, v5)
                                case "foo bar":
                                    if true {
                                        v60 := GenS1{F0: (v3 * v3), F1: v30, F2: v30}
                                        _ = v60
                                    } else {
                                        v1 = (v1 | v1)
                                        if (!(((v5 + v5) >= "a") && false)) {
                                            var v61 int32 = -213
                                            var v62 uint32 = v59
                                            if true {
                                                v63 := GenS1{F0: (v3 % ((v3) | 1)), F1: v30, F2: v30}
                                                fmt.Println(fi58, fi42, v59)
                                                if (v5 == (((((v5 + v5) + (v5 + v5)) + ("\n\t" + "")) + v5) + (v5 + v5))) {
                                                    v1 += (v1 - v1)
                                                }
                                                _ = v63
                                            }
                                            _ = v61
                                            _ = v62
                                        }
                                        if (!true) {
                                            v2[(^(v59))] = v31
                                        }
                                    }
                                }
                                if true {
                                    if true {
                                        if false {
                                            if false {
                                                if ((441 < -684) || (4294967295 < 4294967295)) {
                                                    switch (fi58 - fi42) {
                                                    case 32768:
                                                        v64 := GenS1{F0: v3, F1: (v30 % ((-85) | 1)), F2: (v30 / ((v30) | 1))}
                                                        _ = v64
                                                    default:
                                                        switch v5 {
                                                        case "a":
                                                            switch "a" {
                                                            case "foo bar":
                                                                switch (v31 + (fi58 - (fi42 >> ((0) & 63)))) {
                                                                case 45:
                                                                    v65 := GenS1{F0: 127, F1: v30, F2: v30}
                                                                    _ = v65
                                                                case 9223372036854775807:
                                                                    var v66 float32 = 3.14159
                                                                    _ = v66
                                                                case 32767:
                                                                    v1 = (v1 << ((255) & 63))
                                                                }
                                                            case "":
                                                                v67 := GenS3{F0: false, F1: (v59 * 4294967295), F2: (797 > 642)}
                                                                _ = v67
                                                            case "a":
                                                                var v68 string = ((("foo bar" + v5) + v5) + v5)
                                                                _ = v68
                                                            default:
                                                                v69 := GenS2{F0: (v3 << ((65535) & 63)), F1: 127}
                                                                v70 := GenS4{F0: ("a" + v5)}
                                                                if false {
                                                                    v71 := GenS5{F0: "a", F1: -326, F2: -661}
                                                                    _ = v71
                                                                }
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                        case "foo bar":
                                                            if false {
                                                                v1 += v1
                                                                v30 = (^(v30))
                                                                switch v31 {
                                                                case 9223372036854775807:
                                                                    v5 = ("hello" + v5)
                                                                case 128:
                                                                    v30 += v30
                                                                default:
                                                                    var v72 int8 = v30
                                                                    _ = v72
                                                                }
                                                                if (-9223372036854775808 > -1) {
                                                                    v73 := []int{fi58, 1, fi42}
                                                                    _ = v73
                                                                }
                                                            } else {
                                                                v74 := GenS6{F0: fi42, F1: genFn1(v3)}
                                                                v75 := GenS5{F0: v5, F1: -1, F2: -32768}
                                                                fmt.Println(v1)
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                        default:
                                                            v76 := []int{fi42}
                                                            _ = v76
                                                        }
                                                    }
                                                    if true {
                                                        v77 := GenS7{F0: (v3 - v3), F1: (v5 + v5), F2: genFn1((^(v3)))}
                                                        var v78 bool = (float64(1.5) > (3.14159 + float64(1e10)))
                                                        if v78 {
                                                            if (v78 || ((v31 - (fi42 - 6)) == fi58)) {
                                                                var v79 int32 = 796
                                                                var v80 int32 = (v79 - v79)
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                            if (!v78) {
                                                                v81 := GenS8{F0: v78}
                                                                _ = v81
                                                            }
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    var v82 uint16 = uint16(v59)
                                                    v59 += (v59 | (v59 | v59))
                                                    v82 -= (v1 + (^(v82)))
                                                    _ = v82
                                                }
                                                v5 = "foo bar"
                                                var v83 int32 = 1
                                                _ = v83
                                            }
                                            var v84 float64 = (float64(1e10) / -2.25)
                                            if (!(true || false)) {
                                                delete(v4, (v31 / (((^(v31))) | 1)))
                                                v59 -= (v59 / ((v59) | 1))
                                                var v85 float64 = (v84 * v84)
                                                _ = v85
                                            } else {
                                                v86 := GenS7{F0: genFn1((v3 % (((v3 << ((255) & 63))) | 1))), F1: v5, F2: (v3 * 2)}
                                                switch "a" {
                                                case "hello":
                                                    v3 -= (v3 % (((v3 ^ (v3 - 42))) | 1))
                                                case "foo bar":
                                                    v87 := GenS9{F0: false, F1: (fi58 % ((fi58) | 1))}
                                                    _ = v87
                                                case "":
                                                    v88 := GenS5{F0: v5, F1: -32768, F2: 0}
                                                    _ = v88
                                                }
                                                switch "hello" {
                                                case "a":
                                                    switch (fi58 % (((fi58 + fi58)) | 1)) {
                                                    case -973:
                                                        _ = len(v2)
                                                    case -34:
                                                        var v89 uint = 4294967295
                                                        _ = v89
                                                    case 287:
                                                        if (((!(1000 < 0)) || (18446744073709551615 >= 2)) && ((v59 + v59) <= (v59 / (((v59 ^ 768)) | 1)))) {
                                                            v59 -= v59
                                                            v90 := []int{(fi42 ^ -462)}
                                                            _ = v90
                                                        }
                                                    }
                                                }
                                                _ = v86
                                            }
                                            if false {
                                                v91 := GenS4{F0: v5}
                                                fmt.Println(v84)
                                                switch "" {
                                                case "\n\t":
                                                    _ = len(v2)
                                                case "foo bar":
                                                    switch (fi58 - fi42) {
                                                    case 195:
                                                        v59 = v59
                                                    default:
                                                        var v92 bool = (!true)
                                                        _ = v92
                                                    }
                                                case "hello":
                                                    var v93 uint64 = 18446744073709551615
                                                    _ = v93
                                                }
                                                _ = v91
                                            }
                                            _ = v84
                                        } else {
                                            v94 := GenS10{F0: v1}
                                            switch "" {
                                            case "hello":
                                                switch v5 {
                                                case "a":
                                                    var v95 bool = true
                                                    _ = v95
                                                case "hello":
                                                    v96 := GenS9{F0: true, F1: fi42}
                                                    if ("\n\t" < v5) {
                                                        fmt.Println(v1, fi42, v1)
                                                        v1 = 0
                                                        _ = len(v4)
                                                        var v97 int8 = v30
                                                        v98 := []int{2147483647, fi42, fi42, (fi58 + (fi58 * v31))}
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        var v99 int8 = (v30 ^ int8(v1))
                                                        if (-224 == 471) {
                                                            v100 := GenS6{F0: int(v3), F1: 1}
                                                            fmt.Println(v30, v5, v30)
                                                            fmt.Println(v31, v3)
                                                            var v101 bool = true
                                                            fmt.Println(v59, v3)
                                                            var v102 uint = 1
                                                            v103 := v2[v59]
                                                            if v101 {
                                                                if true {
                                                                    if (v5 != "foo bar") {
                                                                        var v104 int16 = -958
                                                                        var v105 int = (v103 << ((969) & 63))
                                                                        v106 := GenS2{F0: uint8(v59), F1: (v99 << ((2) & 63))}
                                                                        _ = v104
                                                                        _ = v105
                                                                        _ = v106
                                                                    } else {
                                                                        v103 += (v31 % ((fi42) | 1))
                                                                    }
                                                                    if ((!false) && false) {
                                                                        if (1000 < -196) {
                                                                            var v107 uint8 = v3
                                                                            var v108 uint = v102
                                                                            var v109 string = v5
                                                                            var v110 int8 = -1
                                                                            _ = v107
                                                                            _ = v108
                                                                            _ = v109
                                                                            _ = v110
                                                                        }
                                                                        fmt.Println(fi42, v1)
                                                                        v111 := []int{(-(v31))}
                                                                        _ = v111
                                                                    }
                                                                    if v101 {
                                                                        switch 32767 {
                                                                        case -408:
                                                                            var v112 int32 = 0
                                                                            v112 = v112
                                                                            _ = v112
                                                                        default:
                                                                            v1 += v1
                                                                        }
                                                                        _ = len(v4)
                                                                        v103 = (fi58 | (fi58 - v103))
                                                                        fmt.Println(v5, v30, v5)
                                                                        switch v31 {
                                                                        case -675:
                                                                            v113 := GenS11{F0: v101, F1: v102}
                                                                            _ = v113
                                                                        case 700:
                                                                            v102 = uint(v1)
                                                                        case 127:
                                                                            var v114 uint64 = 516
                                                                            _ = v114
                                                                        }
                                                                    } else {
                                                                        switch ("\n\t" + v5) {
                                                                        case "hello":
                                                                            v31 = 293
                                                                        case "a":
                                                                            var v115 uint16 = v1
                                                                            _ = v115
                                                                        case "\n\t":
                                                                            if (4294967295 == 74) {
                                                                                v116 := []int{}
                                                                                _ = v116
                                                                            }
                                                                        }
                                                                        var v117 int8 = -23
                                                                        _ = v117
                                                                    }
                                                                } else {
                                                                    var v118 uint16 = (v1 & 1)
                                                                    if (((1.5 * 1e10) - 0.0) <= (float32(-2.25) + 1.5)) {
                                                                        v119 := GenS11{F0: (!(v118 > v118)), F1: 65535}
                                                                        if (((1e10 * (3.14159 + float32(3.14159))) / ((3.14159 - ((3.14159 + (float32(0.0) * (float32(-2.25) + ((3.14159 - (0.0 + (0.1 - float32((float64((float32(3.14159) + 0.1)) * ((1.5 * 1e10) - 1.5)))))) * (1.5 * (1.5 + -2.25)))))) / ((0.1 + 3.14159) * 0.1))) - float32(0.1))) == (((1e10 - 0.1) - (((0.1 + (1e10 + 3.14159)) - 3.14159) / float32((0.1 * (1e10 / 1.0))))) / 0.1)) {
                                                                            var v120 int64 = 0
                                                                            v99 -= (v30 + (v30 - int8(v102)))
                                                                            _ = v120
                                                                        }
                                                                        _ = v119
                                                                    }
                                                                    fmt.Println(v118, v103)
                                                                    _ = v118
                                                                }
                                                                v121 := GenS12{F0: v3}
                                                                v122 := GenS13{F0: -1, F1: v30, F2: 0}
                                                                _ = v121
                                                                _ = v122
                                                            } else {
                                                                v123 := []int{v31}
                                                                _ = v123
                                                            }
                                                            var v124 int16 = -1
                                                            _ = v100
                                                            _ = v101
                                                            _ = v102
                                                            _ = v103
                                                            _ = v124
                                                        }
                                                        v3 = (v3 | 1)
                                                        _ = v99
                                                    }
                                                    _ = v96
                                                default:
                                                    fmt.Println(v30)
                                                }
                                            case "\n\t":
                                                v31 = v31
                                            case "":
                                                var v125 float64 = -2.25
                                                _ = v125
                                            default:
                                                if false {
                                                    var v126 uint16 = v1
                                                    _ = v126
                                                }
                                            }
                                            switch "" {
                                            case "":
                                                v127 := GenS4{F0: v5}
                                                _ = v127
                                            }
                                            _ = v94
                                        }
                                        if (42 < 497) {
                                            if true {
                                                var v128 uint = uint(v59)
                                                var v129 string = "hello"
                                                _ = v128
                                                _ = v129
                                            }
                                            v130 := []int{fi58, (v31 + fi58)}
                                            v31 += fi58
                                            _ = v130
                                        }
                                        var v131 int16 = -32768
                                        _ = v131
                                    } else {
                                        fmt.Println(v5, v5)
                                    }
                                    if true {
                                        switch ((((("hello" + "foo bar") + v5) + v5) + ((v5 + v5) + "foo bar")) + (v5 + "foo bar")) {
                                        case "a":
                                            if (4294967295 >= 255) {
                                                v132 := GenS14{F0: v30, F1: (v59 & 65535), F2: false}
                                                var v133 uint16 = 255
                                                _ = v132
                                                _ = v133
                                            }
                                        case "hello":
                                            v31 += (v31 << ((4294967295) & 63))
                                        case "foo bar":
                                            if (!(!false)) {
                                                v59 = (v59 * (v59 / ((42) | 1)))
                                            }
                                        default:
                                            var v134 bool = (!(363 > 65535))
                                            _ = v134
                                        }
                                        if true {
                                            fmt.Println(v1, v3)
                                            delete(v2, uint32(v1))
                                        }
                                    }
                                    v59 = v59
                                }
                                _ = fi58
                                _ = v59
                            }
                        } else {
                            v135 := []int{934, (fi42 ^ 461)}
                            var v136 float64 = float64(float32(0.1))
                            v137 := GenS7{F0: v3, F1: "", F2: v3}
                            _ = v135
                            _ = v136
                            _ = v137
                        }
                    }
                    fmt.Println(v5, fi42)
                    for fi138 := 0; fi138 < 6; fi138++ {
                        fmt.Println(v5)
                        if (-9223372036854775808 != -360) {
                            v1 = 1
                            v31 = 127
                            switch (fi138 + v31) {
                            case -637:
                                var v139 int16 = -32768
                                _ = v139
                            case -935:
                                fmt.Println(v1, fi138, v30)
                            }
                        } else {
                            var v140 int64 = -151
                            var v141 int32 = 0
                            v5 = v5
                            _ = v140
                            _ = v141
                        }
                        var v142 int = -14
                        _ = fi138
                        _ = v142
                    }
                }
                if (!(4294967295 < 387)) {
                    var v143 float64 = (-2.25 * (float64(1e10) * 3.14159))
                    v5 = (v5 + v5)
                    v144 := []int{32767, (v31 * -776)}
                    _ = v143
                    _ = v144
                }
            }
        }
        _ = fi42
    }
    switch 0 {
    case 32768:
        if false {
            v145 := v2[65535]
            var v146 uint16 = 1
            _ = v145
            _ = v146
        } else {
            var v147 float32 = (0.1 + 1e10)
            v148 := GenS12{F0: (v3 * 127)}
            _ = v147
            _ = v148
        }
    case 127:
        v30 = (-(v30))
    }
    v31 -= 2147483647
    var v149 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v149
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
