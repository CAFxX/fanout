package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS6 struct { F0 int32; F1 string; F2 int32 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS16 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 string; F1 uint8 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 bool; F1 string; F2 uint16 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 string }
func genFn0(p0 int8, p1 bool, p2 string) uint {
    switch 0 {
    case 128:
        p1 = (!p1)
    case 32768:
        v5 := []int{128, 0}
        _ = v5
    case 1:
        switch "" {
        case "foo bar":
            p2 = ("\n\t" + (("hello" + p2) + p2))
        case "a":
            switch p2 {
            case "hello":
                var v6 uint16 = 510
                _ = v6
            default:
                switch "foo bar" {
                case "hello":
                    var v7 string = p2
                    _ = v7
                }
            }
        }
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 string = "a"
            p0 -= 0
            switch (fi8 + (fi8 << ((454) & 63))) {
            case -9223372036854775808:
                v10 := []int{}
                _ = v10
            case 2147483647:
                var v11 uint16 = 42
                _ = v11
            default:
                var v12 int32 = 42
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 338
}

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 2
    var v3 int8 = 1
    v4 := []uint32{4294967295, 224, 832, 4294967295}
    v3 = v3
    fmt.Println(v1, v1)
    _ = len(v4)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 string = "foo bar"
        var v15 string = v14
        var v16 uint = 4294967295
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: 686, F1: (v1 - 42), F2: ""}
        var v19 int8 = v3
        var v20 bool = true
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: -1, F1: genFn0(-1, false, "\n\t")}
    switch ("\n\t" + "a") {
    case "\n\t":
        v4 = append(v4, 2)
    case "":
        switch "\n\t" {
        case "\n\t":
            v4 = append(v4, 2147483648)
        case "hello":
            fmt.Println(v3, v1)
        }
    case "a":
        v3 = v3
    default:
        for fi22 := 0; fi22 < 7; fi22++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi22
        }
    }
    if false {
        if (false || (!true)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v23 string = ""
            _ = v23
        }
        var v24 uint32 = uint32(v2)
        switch "" {
        case "foo bar":
            for fi25 := 0; fi25 < 5; fi25++ {
                if true {
                    if (!true) {
                        if (((!false) && (true && false)) || ((^(v24)) == (^(v24)))) {
                            fmt.Println(fi25)
                        } else {
                            for fi26 := 0; fi26 < 2; fi26++ {
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch fi25 {
                                    case 813:
                                        v27 := GenS2{F0: 86, F1: genFn0((v3 & 42), false, "hello")}
                                        v28 := GenS3{F0: v3}
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                var v29 uint = 18446744073709551615
                                _ = fi26
                                _ = v29
                            }
                            if false {
                                var v30 bool = false
                                _ = v30
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (!(!(v3 <= 78))) {
                                    switch "hello" {
                                    case "foo bar":
                                        var v31 uint32 = v24
                                        _ = v31
                                    case "a":
                                        var v32 int16 = 32767
                                        _ = v32
                                    }
                                }
                                var v33 uint32 = v24
                                _ = v33
                            }
                            switch fi25 {
                            case -948:
                                if true {
                                    for fi34 := 0; fi34 < 8; fi34++ {
                                        if true {
                                            v24 -= (v24 * 42)
                                        } else {
                                            v35 := GenS1{F0: (^(v1)), F1: 2, F2: (("a" + ("\n\t" + "")) + "a")}
                                            v36 := GenS1{F0: (^(v1)), F1: 1, F2: ""}
                                            _ = v35
                                            _ = v36
                                        }
                                        var v37 bool = true
                                        _ = fi34
                                        _ = v37
                                    }
                                }
                            case 725:
                                switch fi25 {
                                case -287:
                                    v38 := GenS3{F0: 0}
                                    _ = v38
                                case 1:
                                    for fi39 := 0; fi39 < 3; fi39++ {
                                        var v40 uint64 = v2
                                        _ = fi39
                                        _ = v40
                                    }
                                case -2147483648:
                                    v4 = append(v4, v24)
                                }
                            case -72:
                                switch -307 {
                                case 844:
                                    var v41 int = (fi25 - 9223372036854775807)
                                    _ = v41
                                case 42:
                                    v24 = v24
                                default:
                                    fmt.Println(v3, v2)
                                }
                            }
                        }
                        v2 -= (v2 & v2)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        var v42 int32 = 577
                        if true {
                            if (255 > 4294967295) {
                                for fi43 := 0; fi43 < 5; fi43++ {
                                    if (62 != v3) {
                                        if (3.14159 == (3.14159 * 3.14159)) {
                                            fmt.Println(fi25, v1, fi25)
                                            var v44 bool = ((fi43 & fi25) > (fi43 >> ((genFn0((v3 % ((-1) | 1)), false, "hello")) & 63)))
                                            _ = v44
                                        }
                                        fmt.Println(fi43, fi43)
                                    }
                                    _ = fi43
                                }
                                switch fi25 {
                                case 128:
                                    v45 := GenS3{F0: 56}
                                    _ = v45
                                }
                            } else {
                                switch fi25 {
                                case -455:
                                    v46 := GenS3{F0: (v3 >> ((0) & 63))}
                                    _ = v46
                                default:
                                    var v47 uint64 = (v2 / (((v2 & (v2 | (^(v2))))) | 1))
                                    _ = v47
                                }
                                if false {
                                    var v48 string = "foo bar"
                                    _ = v48
                                }
                                if (0 >= 140) {
                                    for fi49 := 0; fi49 < 11; fi49++ {
                                        if (-2.25 > 0.0) {
                                            var v50 int32 = (-(v42))
                                            if true {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v2 -= v2
                                                case "a":
                                                    _ = len(v4)
                                                case "":
                                                    if true {
                                                        if (float64(3.14159) >= float64(1.5)) {
                                                            var v51 uint8 = 255
                                                            _ = v51
                                                        }
                                                        switch "" {
                                                        case "a":
                                                            if true {
                                                                if (((v3 | (v3 % ((v3) | 1))) >= 56) || (!true)) {
                                                                    var v52 uint16 = 835
                                                                    v53 := GenS2{F0: (v50 - (^(v42))), F1: 450}
                                                                    v4 = append(v4, v24)
                                                                    _ = v52
                                                                    _ = v53
                                                                } else {
                                                                    var v54 uint8 = 128
                                                                    _ = v54
                                                                }
                                                                _ = len(v4)
                                                                if (-129 >= 42) {
                                                                    if (1.5 <= 1e10) {
                                                                        fmt.Println(v50)
                                                                    } else {
                                                                        v50 -= v50
                                                                        fmt.Println(v3, v2, fi25)
                                                                        var v55 float64 = (0.0 * 0.0)
                                                                        _ = v55
                                                                    }
                                                                    fmt.Println(v3, v42)
                                                                } else {
                                                                    var v56 uint8 = 128
                                                                    _ = v56
                                                                }
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v57 int16 = 32767
                                                            _ = v57
                                                        case "foo bar":
                                                            if false {
                                                                if false {
                                                                    v42 += (^(v50))
                                                                    v58 := GenS2{F0: v42, F1: 2}
                                                                    _ = v58
                                                                }
                                                                v59 := GenS1{F0: (^(v1)), F1: (v1 & v1), F2: "foo bar"}
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    if ((2 < 2) && false) {
                                                                        v42 = v50
                                                                        _ = len(v4)
                                                                        if true {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (^(v24))
                                                                            }
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v60 int8 = -128
                                                                    _ = v60
                                                                case "hello":
                                                                    v61 := GenS2{F0: (-(v50)), F1: 293}
                                                                    _ = v61
                                                                }
                                                                _ = v59
                                                            }
                                                        case "hello":
                                                            v3 = -128
                                                        }
                                                        var v62 int16 = -41
                                                        _ = v62
                                                    } else {
                                                        v63 := GenS2{F0: 0, F1: 42}
                                                        if false {
                                                            switch 32768 {
                                                            case 128:
                                                                v24 = v24
                                                            default:
                                                                v64 := GenS2{F0: v42, F1: 848}
                                                                v65 := GenS3{F0: v3}
                                                                var v66 bool = true
                                                                v67 := GenS3{F0: v3}
                                                                v68 := GenS4{F0: v1}
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                            v1 -= v1
                                                            v69 := GenS4{F0: v1}
                                                            _ = v69
                                                        } else {
                                                            v50 += v50
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        _ = v63
                                                    }
                                                }
                                            }
                                            fmt.Println(v1)
                                            v70 := GenS5{F0: 27, F1: v3, F2: v24}
                                            v71 := GenS1{F0: v1, F1: (v1 / ((v1) | 1)), F2: "hello"}
                                            _ = v50
                                            _ = v70
                                            _ = v71
                                        }
                                        var v72 int32 = v42
                                        v72 -= v42
                                        if (float64(3.14159) == ((1e10 + (float64((-2.25 + 3.14159)) * -2.25)) / 1.0)) {
                                            v73 := GenS6{F0: (v42 ^ (v72 | 42)), F1: ("foo bar" + "foo bar"), F2: v42}
                                            var v74 uint = 294
                                            _ = v73
                                            _ = v74
                                        } else {
                                            fmt.Println(v72, v24, fi49)
                                        }
                                        _ = fi49
                                        _ = v72
                                    }
                                    for fi75 := 0; fi75 < 9; fi75++ {
                                        switch (fi75 >> ((genFn0((v3 | (v3 ^ v3)), true, ("foo bar" + "\n\t"))) & 63)) {
                                        case 9223372036854775807:
                                            if true {
                                                v76 := GenS7{F0: (v24 - v24), F1: (v1 ^ 641), F2: v3}
                                                _ = v76
                                            }
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = (v24 * v24)
                                        }
                                        var v77 float32 = (1e10 - (((-2.25 / float32((3.14159 - 1.5))) * 0.0) * 1e10))
                                        _ = fi75
                                        _ = v77
                                    }
                                    v4 = append(v4, v24)
                                }
                            }
                            for fi78 := 0; fi78 < 16; fi78++ {
                                v4 = append(v4, (v24 ^ (v24 << ((42) & 63))))
                                v79 := GenS8{F0: 569, F1: int8(v42)}
                                v4 = append(v4, (v24 - v24))
                                _ = fi78
                                _ = v79
                            }
                        } else {
                            var v80 uint64 = (v2 / (((v2 - (v2 << ((743) & 63)))) | 1))
                            v1 = v1
                            _ = v80
                        }
                        v2 += 32
                        _ = v42
                    }
                } else {
                    var v81 int32 = 0
                    _ = v81
                }
                _ = fi25
            }
        case "a":
            if (!true) {
                switch 434 {
                case 616:
                    for fi82 := 0; fi82 < 11; fi82++ {
                        for fi83 := 0; fi83 < 10; fi83++ {
                            v2 = (v2 ^ (v2 % ((uint64(v1)) | 1)))
                            var v84 bool = true
                            if false {
                                var v85 string = "a"
                                var v86 int32 = -1
                                var v87 int64 = 1
                                _ = v85
                                _ = v86
                                _ = v87
                            } else {
                                if ((^(v24)) < 1) {
                                    if v84 {
                                        var v88 int16 = 42
                                        v1 = (^(v1))
                                        _ = v88
                                    } else {
                                        var v89 float32 = ((1e10 * 1e10) + 0.1)
                                        if (v84 && ((fi83 >= (fi83 << ((18446744073709551615) & 63))) || v84)) {
                                            if (1 > 214) {
                                                if len(v4) > 0 {
                                                    v4[0] = 65535
                                                }
                                            } else {
                                                if false {
                                                    v3 -= (v3 ^ v3)
                                                    v90 := GenS9{F0: v1}
                                                    v4 = append(v4, 4294967295)
                                                    switch ("" + "a") {
                                                    case "\n\t":
                                                        fmt.Println(fi82)
                                                    case "":
                                                        fmt.Println(fi83, v2, v24)
                                                    default:
                                                        var v91 int32 = -1
                                                        _ = v91
                                                    }
                                                    _ = v90
                                                } else {
                                                    fmt.Println(v84)
                                                    if len(v4) > 0 {
                                                        v4[0] = 957
                                                    }
                                                }
                                                v24 += (^(v24))
                                            }
                                            v4 = append(v4, 2147483647)
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = 2147483647
                                            }
                                            var v92 float32 = float32(float64(0.1))
                                            if (v84 && ((v1 % ((42) | 1)) != v1)) {
                                                switch ("\n\t" + "hello") {
                                                case "hello":
                                                    if (42 < -914) {
                                                        if ((!v84) || (("\n\t" + (("hello" + (("hello" + "foo bar") + "a")) + "foo bar")) >= "hello")) {
                                                            v3 = 1
                                                        }
                                                        v93 := GenS10{F0: uint8(v24)}
                                                        var v94 int = (fi82 / ((fi83) | 1))
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = v24
                                                        }
                                                        v89 -= (v89 - ((0.0 / 1.5) / (float32(3.14159) - v89)))
                                                        var v95 bool = ("\n\t" == "\n\t")
                                                        _ = v95
                                                    }
                                                case "":
                                                    v2 += (v2 - (v2 * v2))
                                                }
                                                v96 := GenS11{F0: v1, F1: -381, F2: 222}
                                                _ = v96
                                            }
                                            _ = v92
                                        }
                                        var v97 float32 = v89
                                        _ = v89
                                        _ = v97
                                    }
                                    var v98 float64 = 0.1
                                    if v84 {
                                        v84 = ((v84 || v84) && (!true))
                                        var v99 float64 = (v98 * -2.25)
                                        var v100 int8 = 0
                                        var v101 uint32 = v24
                                        if (!true) {
                                            v4 = append(v4, 2147483648)
                                        } else {
                                            switch fi82 {
                                            case 0:
                                                var v102 uint = 18446744073709551615
                                                _ = v102
                                            case -9223372036854775808:
                                                switch "hello" {
                                                case "hello":
                                                    var v103 int8 = 0
                                                    _ = v103
                                                case "\n\t":
                                                    fmt.Println(v99, fi82, v2)
                                                case "foo bar":
                                                    switch fi83 {
                                                    case 4294967295:
                                                        var v104 int = 32767
                                                        _ = v104
                                                    }
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            }
                                        }
                                        _ = v100
                                        _ = v101
                                        _ = v99
                                    } else {
                                        v4 = append(v4, v24)
                                        if false {
                                            var v105 uint32 = v24
                                            _ = v105
                                        }
                                        var v106 float64 = (1.5 * (v98 + v98))
                                        _ = v106
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v98
                                } else {
                                    v2 += v2
                                    var v107 int16 = 481
                                    if len(v4) > 0 {
                                        v4[0] = v24
                                    }
                                    _ = v107
                                }
                                if (-32768 != 659) {
                                    if (!v84) {
                                        var v108 int16 = 1
                                        if v84 {
                                            var v109 uint16 = 0
                                            _ = v109
                                        }
                                        _ = v108
                                    } else {
                                        var v110 int8 = v3
                                        _ = v110
                                    }
                                } else {
                                    var v111 uint32 = (^(v24))
                                    v112 := GenS2{F0: 2147483647, F1: 255}
                                    _ = v111
                                    _ = v112
                                }
                                if v84 {
                                    var v113 float32 = 3.14159
                                    var v114 int32 = 1
                                    fmt.Println(v3)
                                    _ = v113
                                    _ = v114
                                }
                            }
                            _ = fi83
                            _ = v84
                        }
                        v115 := GenS12{F0: "a"}
                        _ = fi82
                        _ = v115
                    }
                default:
                    v116 := GenS9{F0: v1}
                    _ = v116
                }
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v117 uint64 = v2
                _ = v117
            } else {
                v24 = (v24 ^ v24)
                for fi118 := 0; fi118 < 14; fi118++ {
                    if (!true) {
                        for fi119 := 0; fi119 < 11; fi119++ {
                            v120 := GenS13{F0: fi119}
                            var v121 int8 = -61
                            var v122 int8 = int8(v1)
                            _ = fi119
                            _ = v120
                            _ = v121
                            _ = v122
                        }
                        switch "hello" {
                        case "hello":
                            v3 = (v3 | 85)
                        }
                    }
                    if (((v1 - uint16(v2)) == 42) || true) {
                        switch "" {
                        case "a":
                            switch (("hello" + "a") + "a") {
                            case "\n\t":
                                var v123 float64 = ((1e10 + -2.25) - 1e10)
                                _ = v123
                            case "a":
                                if false {
                                    for fi124 := 0; fi124 < 5; fi124++ {
                                        v24 = uint32(v1)
                                        _ = fi124
                                    }
                                    v1 -= (^(v1))
                                    v2 += (v2 << ((18446744073709551615) & 63))
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        v24 = v24
                                    case "foo bar":
                                        v2 = (v2 + v2)
                                    default:
                                        var v125 bool = true
                                        _ = v125
                                    }
                                    v126 := GenS9{F0: 42}
                                    v127 := GenS10{F0: 114}
                                    _ = v126
                                    _ = v127
                                }
                            case "":
                                if (4294967295 == 42) {
                                    v128 := GenS14{F0: 42}
                                    var v129 string = "foo bar"
                                    _ = v128
                                    _ = v129
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = (v24 + v24)
                                    }
                                    for fi130 := 0; fi130 < 14; fi130++ {
                                        if false {
                                            if true {
                                                var v131 float32 = float32(1e10)
                                                v24 -= (^(v24))
                                                _ = v131
                                            } else {
                                                fmt.Println(fi118, v3)
                                                _ = len(v4)
                                                if true {
                                                    if true {
                                                        switch (fi118 + int(v2)) {
                                                        case -1:
                                                            var v132 int8 = 1
                                                            _ = v132
                                                        case 209:
                                                            fmt.Println(fi130, v24)
                                                        case 32768:
                                                            switch fi118 {
                                                            case 32767:
                                                                if ("hello" < "foo bar") {
                                                                    var v133 string = (((("hello" + ("hello" + "")) + ("hello" + "\n\t")) + "") + "a")
                                                                    _ = v133
                                                                } else {
                                                                    var v134 float64 = 1e10
                                                                    v135 := GenS9{F0: v1}
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                            case 9223372036854775807:
                                                                if (!false) {
                                                                    var v136 uint64 = v2
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        if false {
                                                                            v3 -= (^(v3))
                                                                        }
                                                                    case "hello":
                                                                        switch (-(fi118)) {
                                                                        case 9223372036854775807:
                                                                            if (!false) {
                                                                                if true {
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        if (18446744073709551615 > (^(v2))) {
                                                                                            if (4294967295 < 4294967295) {
                                                                                                var v137 bool = (true && true)
                                                                                                _ = v137
                                                                                            }
                                                                                            if false {
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = 1
                                                                                                }
                                                                                                v2 += 1
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = 2
                                                                                                }
                                                                                            } else {
                                                                                                v136 -= (v2 + (v2 << ((1) & 63)))
                                                                                                var v138 string = "a"
                                                                                                _ = v138
                                                                                            }
                                                                                            fmt.Println(v136, v2)
                                                                                            _ = len(v4)
                                                                                            if true {
                                                                                                v4 = append(v4, 2147483648)
                                                                                            } else {
                                                                                                var v139 uint8 = 255
                                                                                                v2 = v136
                                                                                                v139 = (v139 + v139)
                                                                                                _ = v139
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v140 string = ("foo bar" + "a")
                                                                                        _ = v140
                                                                                    default:
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            v1 = (v1 % ((v1) | 1))
                                                                                        case "a":
                                                                                            var v141 uint64 = v136
                                                                                            _ = v141
                                                                                        }
                                                                                    }
                                                                                    v4 = append(v4, 670)
                                                                                    v142 := GenS13{F0: (^(fi130))}
                                                                                    fmt.Println(v24, v136)
                                                                                    _ = v142
                                                                                }
                                                                                if true {
                                                                                    v3 += (v3 - v3)
                                                                                }
                                                                            } else {
                                                                                var v143 int64 = 61
                                                                                _ = v143
                                                                            }
                                                                        default:
                                                                            v144 := GenS15{F0: -2147483648, F1: 0, F2: v2}
                                                                            _ = v144
                                                                        }
                                                                    default:
                                                                        var v145 uint = 2
                                                                        _ = v145
                                                                    }
                                                                    _ = v136
                                                                } else {
                                                                    if ((v1 ^ (v1 - (v1 + 232))) >= (v1 - (v1 >> ((genFn0(v3, (1 <= -1), ("foo bar" + "foo bar"))) & 63)))) {
                                                                        var v146 int = fi118
                                                                        if false {
                                                                            var v147 int32 = -1
                                                                            _ = v147
                                                                        }
                                                                        var v148 int32 = 2147483647
                                                                        _ = v146
                                                                        _ = v148
                                                                    } else {
                                                                        var v149 float64 = 1.5
                                                                        if (float64(1.5) > v149) {
                                                                            v150 := GenS12{F0: "\n\t"}
                                                                            var v151 int64 = 1
                                                                            _ = v150
                                                                            _ = v151
                                                                        } else {
                                                                            var v152 bool = false
                                                                            var v153 float64 = 1e10
                                                                            var v154 bool = v152
                                                                            _ = v152
                                                                            _ = v153
                                                                            _ = v154
                                                                        }
                                                                        _ = v149
                                                                    }
                                                                    var v155 int32 = -1
                                                                    _ = v155
                                                                }
                                                            }
                                                        default:
                                                            var v156 uint16 = v1
                                                            _ = v156
                                                        }
                                                        fmt.Println(fi130, v1, fi130)
                                                        switch fi118 {
                                                        case -1:
                                                            if (!(false && false)) {
                                                                fmt.Println(v3, v2, v3)
                                                            }
                                                        default:
                                                            v2 += (v2 - (v2 << ((genFn0(v3, ("foo bar" != ""), "foo bar")) & 63)))
                                                        }
                                                    }
                                                    fmt.Println(v2, v2, fi130)
                                                }
                                            }
                                            if (!((^(v1)) > uint16(v2))) {
                                                var v157 float64 = (0.1 * (1e10 - ((((3.14159 * (float64((1.5 * 0.0)) / 0.1)) - 0.0) + 1e10) - 1e10)))
                                                _ = v157
                                            } else {
                                                if (128 >= 127) {
                                                    if true {
                                                        var v158 uint32 = 1
                                                        v3 = v3
                                                        var v159 string = "foo bar"
                                                        _ = v158
                                                        _ = v159
                                                    }
                                                    v2 -= (v2 & (v2 * v2))
                                                    var v160 int32 = -2147483648
                                                    _ = v160
                                                }
                                            }
                                            _ = len(v4)
                                        }
                                        _ = len(v4)
                                        _ = fi130
                                    }
                                }
                            default:
                                v161 := GenS16{F0: -1, F1: (^(v24)), F2: v2}
                                _ = v161
                            }
                        default:
                            v162 := GenS17{F0: 2}
                            _ = v162
                        }
                        v1 += (v1 * v1)
                    }
                    if (42 > genFn0(42, false, ("a" + "a"))) {
                        if (42 < 1) {
                            fmt.Println(fi118, v3, v2)
                        } else {
                            if false {
                                v163 := GenS18{F0: "hello", F1: 1}
                                if (-1 < -904) {
                                    for fi164 := 0; fi164 < 4; fi164++ {
                                        var v165 uint32 = (v24 / (((v24 / ((v24) | 1))) | 1))
                                        fmt.Println(fi118, fi164, v2)
                                        v165 -= (v165 ^ v165)
                                        _ = fi164
                                        _ = v165
                                    }
                                    if (!((v1 % ((v1) | 1)) != v1)) {
                                        for fi166 := 0; fi166 < 2; fi166++ {
                                            v3 -= v3
                                            fmt.Println(v24)
                                            _ = fi166
                                        }
                                        v24 += 499
                                    }
                                    v4 = append(v4, 2147483648)
                                }
                                v4 = append(v4, 4294967295)
                                _ = v163
                            }
                            if (0 >= 9223372036854775807) {
                                for fi167 := 0; fi167 < 18; fi167++ {
                                    if len(v4) > 0 {
                                        v4[0] = 2
                                    }
                                    if true {
                                        var v168 uint = 42
                                        v169 := GenS19{F0: 0}
                                        v4 = append(v4, v24)
                                        _ = v168
                                        _ = v169
                                    } else {
                                        v3 -= v3
                                        switch 32767 {
                                        case 32767:
                                            var v170 int8 = (v3 << ((321) & 63))
                                            _ = v170
                                        default:
                                            var v171 int64 = 285
                                            _ = v171
                                        }
                                    }
                                    fmt.Println(v3)
                                    _ = fi167
                                }
                                v3 = -1
                            } else {
                                v4 = append(v4, 1)
                            }
                            v172 := GenS20{F0: true, F1: (("foo bar" + "\n\t") + ""), F2: 65535}
                            _ = v172
                        }
                        v1 = (v1 - (v1 % ((v1) | 1)))
                        fmt.Println(v2, v1, v2)
                    }
                    _ = fi118
                }
            }
        }
        _ = v24
    }
    v173 := GenS21{F0: (^(v2))}
    v174 := GenS3{F0: v3}
    v2 = (v2 >> ((42) & 63))
    v2 = v2
    v175 := GenS22{F0: "foo bar"}
    _ = len(v4)
    v4 = append(v4, 1)
    if (true && ((!false) && (false && true))) {
        v1 += uint16(v2)
        v3 += 110
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v173
    _ = v174
    _ = v175
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
