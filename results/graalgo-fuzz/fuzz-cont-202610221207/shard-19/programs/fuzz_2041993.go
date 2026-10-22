package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 string; F1 int32; F2 bool }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8; F1 uint64; F2 int8 }
type GenS10 struct { F0 uint32; F1 int8; F2 int16 }
type GenS11 struct { F0 uint32; F1 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string; F1 bool; F2 uint16 }
type GenS14 struct { F0 int64; F1 uint }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 string; F1 uint8 }
type GenS17 struct { F0 int8; F1 int; F2 int8 }
type GenS18 struct { F0 uint32; F1 string }
type GenS19 struct { F0 uint; F1 int8 }
func genFn0() string {
    var v4 int64 = 9223372036854775807
    v5 := []int{527, 945}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (-128 >= 98) {
        v6 := GenS1{F0: 1000, F1: ("foo bar" + (("a" + "a") + "foo bar")), F2: 183}
        switch "a" {
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        case "foo bar":
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int64 = (^(v4))
                v4 = 1
                _ = fi7
                _ = v8
            }
        default:
            if true {
                switch 42 {
                case 9223372036854775807:
                    v9 := GenS2{F0: 255, F1: 217}
                    _ = v9
                case 42:
                    v4 -= v4
                case -1:
                    v4 = (v4 / ((-801) | 1))
                }
                v4 = v4
                for fi10 := 0; fi10 < 17; fi10++ {
                    _ = len(v5)
                    _ = fi10
                }
            }
        }
        _ = v6
    }
    v5 = append(v5, 128)
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1() int32 {
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS3{F0: (915 == 32767), F1: 42}
                var v14 uint8 = 127
                if ("" != ("hello" + "hello")) {
                    var v15 int16 = -32768
                    var v16 float64 = 3.14159
                    v15 += v15
                    _ = v15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            if false {
                var v17 float64 = 1.5
                v18 := []int{fi11}
                _ = v17
                _ = v18
            } else {
                var v19 uint16 = 255
                var v20 float32 = float32(((-2.25 - (-2.25 * 3.14159)) / (0.1 + (1e10 + 0.0))))
                fmt.Println(fi11)
                _ = v19
                _ = v20
            }
            _ = fi11
        }
        if true {
            for fi21 := 0; fi21 < 2; fi21++ {
                if false {
                    var v22 uint64 = 2
                    _ = v22
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        var v23 uint = 4294967295
                        _ = v23
                    case "foo bar":
                        var v24 int64 = 1
                        _ = v24
                    default:
                        if ((!false) || false) {
                            v25 := GenS4{F0: "\n\t", F1: -767, F2: true}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 19; fi26++ {
                                switch fi26 {
                                case 9223372036854775807:
                                    if true {
                                        if ((fi21 != fi21) || true) {
                                            v27 := []int{(fi26 ^ (fi21 ^ (fi21 ^ (fi26 % (((^(fi21))) | 1))))), fi26}
                                            _ = v27
                                        }
                                        if true {
                                            v28 := GenS4{F0: "", F1: 1, F2: true}
                                            v29 := []int{(fi21 / ((fi21) | 1)), (fi21 + fi21), fi21, (fi21 * (fi26 + 713))}
                                            _ = v28
                                            _ = v29
                                        }
                                    }
                                case -9223372036854775808:
                                    var v30 uint32 = 42
                                    _ = v30
                                case 297:
                                    var v31 int8 = 127
                                    _ = v31
                                }
                                if ((fi26 ^ fi26) == fi21) {
                                    switch fi21 {
                                    case 4294967295:
                                        var v32 uint8 = 128
                                        _ = v32
                                    case -1:
                                        var v33 int64 = 0
                                        _ = v33
                                    }
                                } else {
                                    var v34 int = (fi21 - fi21)
                                    _ = v34
                                }
                                var v35 uint16 = 2
                                _ = fi26
                                _ = v35
                            }
                        }
                    }
                }
                fmt.Println(fi21)
                _ = fi21
            }
            var v36 float32 = 0.1
            v37 := []int{-9223372036854775808, 128, 4294967295}
            _ = v36
            _ = v37
        }
        v38 := GenS5{F0: 914, F1: 4294967295}
        _ = v38
    }
    if false {
        v39 := GenS1{F0: 42, F1: "\n\t", F2: 1}
        _ = v39
    } else {
        var v40 string = "a"
        _ = v40
    }
    var v41 float32 = 0.1
    v41 = -2.25
    _ = v41
    return 2147483647
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    var v42 uint16 = 632
    switch -9223372036854775808 {
    case 128:
        for fi43 := 0; fi43 < 13; fi43++ {
            v44 := v2[2147483647]
            var v45 uint32 = v44
            for fi46 := 0; fi46 < 13; fi46++ {
                var v47 int32 = 2147483647
                _ = fi46
                _ = v47
            }
            _ = fi43
            _ = v44
            _ = v45
        }
    case 1:
        v3 = (v3 + v3)
    }
    fmt.Println(v3, v42)
    v1["\n\t"] = 460
    _ = len(v2)
    var v48 uint = 18446744073709551615
    v49 := GenS1{F0: 42, F1: "\n\t", F2: (v42 ^ 153)}
    var v50 float64 = (1.5 * float64(0.0))
    var v51 uint64 = 42
    if (!(!false)) {
        if (true && false) {
            var v52 float64 = 1e10
            _ = v52
        }
        v53 := []int{854, -1, 32767, 4294967295}
        _ = v53
    } else {
        if false {
            v50 = (float64(float32((3.14159 / 3.14159))) / 3.14159)
            switch 42 {
            case 128:
                v42 = 585
            case 921:
                v3 += (v3 & int64(v48))
            case -2147483648:
                for fi54 := 0; fi54 < 10; fi54++ {
                    v55 := GenS6{F0: (uint16(v51) == (v42 % (((v42 * uint16(v51))) | 1))), F1: (v51 ^ v51)}
                    if ((v48 / ((v48) | 1)) <= v48) {
                        fmt.Println(fi54, v48)
                        switch "foo bar" {
                        case "":
                            v56 := []int{}
                            _ = v56
                        case "hello":
                            var v57 int16 = 42
                            _ = v57
                        }
                        var v58 int8 = 68
                        _ = v58
                    } else {
                        v50 += 1.5
                        if true {
                            v51 = v51
                            if ((65535 <= v42) && (42 < v42)) {
                                var v59 uint8 = 116
                                fmt.Println(v51)
                                _ = v59
                            } else {
                                v60 := GenS7{F0: ""}
                                v61 := GenS8{F0: 2147483647}
                                v51 = (^(v51))
                                _ = v60
                                _ = v61
                            }
                            switch fi54 {
                            case 28:
                                fmt.Println(fi54)
                            case 9223372036854775807:
                                switch fi54 {
                                case 32767:
                                    for fi62 := 0; fi62 < 16; fi62++ {
                                        var v63 uint = (v48 + v48)
                                        if false {
                                            v64 := []int{}
                                            var v65 uint64 = 42
                                            var v66 int32 = -2147483648
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                        var v67 string = ""
                                        _ = fi62
                                        _ = v63
                                        _ = v67
                                    }
                                case 2147483647:
                                    switch (fi54 << ((v48) & 63)) {
                                    case 32767:
                                        if false {
                                            if true {
                                                for fi68 := 0; fi68 < 11; fi68++ {
                                                    switch 774 {
                                                    case 32767:
                                                        v3 = (v3 & (v3 / ((v3) | 1)))
                                                    case -640:
                                                        fmt.Println(fi54, v42, fi68)
                                                    case 336:
                                                        var v69 uint64 = v51
                                                        _ = v69
                                                    }
                                                    if true {
                                                        v70 := GenS5{F0: 429, F1: 2}
                                                        delete(v2, 2147483647)
                                                        v50 -= float64((3.14159 + float32(v50)))
                                                        _ = v70
                                                    } else {
                                                        var v71 uint64 = v51
                                                        v1["a"] = 805
                                                        switch ("a" + "") {
                                                        case "foo bar":
                                                            if (-1 == 102) {
                                                                switch "a" {
                                                                case "a":
                                                                    v50 -= (1e10 + ((0.1 + 0.0) * v50))
                                                                case "hello":
                                                                    v51 += (v71 >> ((v48) & 63))
                                                                }
                                                                v3 -= int64(fi54)
                                                                v50 = 0.1
                                                            } else {
                                                                v72 := v1["a"]
                                                                fmt.Println(v50)
                                                                v72 = 2
                                                                _ = v72
                                                            }
                                                        }
                                                        _ = v71
                                                    }
                                                    _ = fi68
                                                }
                                            }
                                            v73 := GenS9{F0: -1, F1: (^(v51)), F2: 127}
                                            var v74 bool = false
                                            _ = v73
                                            _ = v74
                                        } else {
                                            v50 = float64(0.0)
                                            var v75 float64 = 0.0
                                            _ = v75
                                        }
                                    case -1:
                                        if false {
                                            switch -193 {
                                            case 32768:
                                                v76 := []int{4294967295, (fi54 / ((fi54) | 1))}
                                                _ = v76
                                            case -242:
                                                if true {
                                                    var v77 string = "hello"
                                                    _ = v77
                                                }
                                            }
                                            v50 += ((float64((((0.1 * 0.1) - 3.14159) + 0.0)) / v50) - v50)
                                            switch fi54 {
                                            case 954:
                                                if false {
                                                    if true {
                                                        var v78 uint8 = 128
                                                        _ = v78
                                                    } else {
                                                        var v79 int32 = 1
                                                        v80 := []int{-959, fi54, fi54, fi54}
                                                        for fi81 := 0; fi81 < 2; fi81++ {
                                                            fmt.Println(fi54, v3, v3)
                                                            var v82 int32 = v79
                                                            if (float32(v50) <= ((0.0 / -2.25) / -2.25)) {
                                                                var v83 int16 = -32768
                                                                var v84 float32 = (-2.25 * (-2.25 - 1e10))
                                                                v85 := GenS10{F0: 39, F1: 0, F2: (v83 | (v83 - 0))}
                                                                _ = v83
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                if (1e10 >= (-2.25 + float64(0.0))) {
                                                                    v82 += int32(v42)
                                                                    v86 := GenS11{F0: 4294967295, F1: 127}
                                                                    _ = v86
                                                                }
                                                            }
                                                            _ = fi81
                                                            _ = v82
                                                        }
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                } else {
                                                    for fi87 := 0; fi87 < 3; fi87++ {
                                                        v88 := GenS8{F0: 156}
                                                        var v89 int32 = -2147483648
                                                        var v90 uint16 = v42
                                                        if (729 == 0) {
                                                            v91 := GenS7{F0: (("hello" + "") + ("\n\t" + ((("\n\t" + "\n\t") + "foo bar") + "a")))}
                                                            _ = v91
                                                        } else {
                                                            _ = len(v1)
                                                        }
                                                        fmt.Println(v90)
                                                        _ = fi87
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                    var v92 uint8 = 2
                                                    for fi93 := 0; fi93 < 6; fi93++ {
                                                        v50 -= 3.14159
                                                        fmt.Println(fi54)
                                                        if (false && (uint8(v51) == (v92 | (^(v92))))) {
                                                            v2[685] = 2
                                                            v51 -= v51
                                                        } else {
                                                            v94 := []int{fi93}
                                                            switch "a" {
                                                            case "hello":
                                                                v3 = -948
                                                            default:
                                                                v92 -= 0
                                                            }
                                                            var v95 uint16 = v42
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                        fmt.Println(v50, fi54)
                                                        _ = fi93
                                                    }
                                                    _ = v92
                                                }
                                            case 175:
                                                for fi96 := 0; fi96 < 4; fi96++ {
                                                    v97 := []int{(fi54 * (fi96 - fi96)), (^(fi54)), fi96, 127}
                                                    switch "\n\t" {
                                                    case "":
                                                        if ((v42 / ((2) | 1)) >= 2) {
                                                            v98 := GenS12{F0: 255}
                                                            var v99 string = "hello"
                                                            var v100 int8 = 127
                                                            switch fi54 {
                                                            case 32767:
                                                                var v101 float32 = float32(0.0)
                                                                v50 = v50
                                                                v102 := GenS13{F0: (v99 + ("\n\t" + "\n\t")), F1: (!(false && true)), F2: v42}
                                                                _ = v101
                                                                _ = v102
                                                            case 127:
                                                                v42 += (v42 * v42)
                                                            case -1:
                                                                v50 = (1.5 * v50)
                                                            default:
                                                                switch "a" {
                                                                case "":
                                                                    var v103 uint64 = (v51 % ((1) | 1))
                                                                    _ = v103
                                                                case "\n\t":
                                                                    v104 := GenS14{F0: 42, F1: (v48 | v48)}
                                                                    v105 := GenS9{F0: v100, F1: v51, F2: (^(v100))}
                                                                    var v106 uint16 = v42
                                                                    _ = v104
                                                                    _ = v105
                                                                    _ = v106
                                                                }
                                                            }
                                                            _ = v100
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                    }
                                                    v3 = 750
                                                    _ = fi96
                                                    _ = v97
                                                }
                                            default:
                                                var v107 string = ""
                                                _ = v107
                                            }
                                        }
                                    default:
                                        v51 = 796
                                    }
                                }
                            default:
                                if true {
                                    switch 42 {
                                    case 28:
                                        if true {
                                            v108 := GenS15{F0: (v48 << (((v48 ^ v48)) & 63))}
                                            if false {
                                                for fi109 := 0; fi109 < 11; fi109++ {
                                                    v110 := GenS16{F0: ("" + ("" + "hello")), F1: 127}
                                                    _ = fi109
                                                    _ = v110
                                                }
                                                var v111 uint32 = 2
                                                var v112 int32 = -1
                                                _ = v111
                                                _ = v112
                                            }
                                            for fi113 := 0; fi113 < 4; fi113++ {
                                                if ((v48 % ((674) | 1)) != (v48 | v48)) {
                                                    if true {
                                                        switch fi54 {
                                                        case 40:
                                                            v48 = 18446744073709551615
                                                        case 90:
                                                            var v114 int32 = -1
                                                            v51 -= (^(v51))
                                                            _ = v114
                                                        case -1:
                                                            var v115 string = "foo bar"
                                                            _ = v115
                                                        }
                                                        if true {
                                                            v48 = v48
                                                        }
                                                        v116 := GenS1{F0: -232, F1: "foo bar", F2: 0}
                                                        if (!false) {
                                                            fmt.Println(v51, v3)
                                                        }
                                                        v3 = (v3 >> (((v48 & 80)) & 63))
                                                        v3 -= -732
                                                        _ = v116
                                                    } else {
                                                        if ((0.1 > -2.25) || false) {
                                                            fmt.Println(fi54, fi113, v48)
                                                        }
                                                    }
                                                    switch (("foo bar" + "\n\t") + ((("foo bar" + "foo bar") + ("foo bar" + "foo bar")) + "foo bar")) {
                                                    case "\n\t":
                                                        switch fi113 {
                                                        case -2147483648:
                                                            fmt.Println(fi54)
                                                            v1[("\n\t" + "hello")] = 65535
                                                        case 127:
                                                            v2[1] = uint32(v42)
                                                        default:
                                                            var v117 int64 = (v3 % (((v3 % ((v3) | 1))) | 1))
                                                            _ = v117
                                                        }
                                                    case "foo bar":
                                                        v118 := []int{604, -1, fi113, (fi113 / (((fi54 << (((v48 ^ (v48 >> (((^(v48))) & 63)))) & 63))) | 1))}
                                                        _ = v118
                                                    case "hello":
                                                        if (-92 > -128) {
                                                            var v119 int32 = 2147483647
                                                            v48 -= v48
                                                            var v120 uint = (v48 - v48)
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                    }
                                                } else {
                                                    var v121 int32 = -2147483648
                                                    _ = v121
                                                }
                                                switch ("a" + ("" + "hello")) {
                                                case "":
                                                    v122 := GenS8{F0: 42}
                                                    _ = v122
                                                }
                                                if false {
                                                    if (!(("\n\t" + (("\n\t" + "\n\t") + ("foo bar" + ""))) < ("a" + "\n\t"))) {
                                                        var v123 float64 = ((((v50 / 3.14159) / v50) + v50) - v50)
                                                        var v124 uint32 = 79
                                                        if false {
                                                            v1["a"] = v124
                                                            v48 = (v48 % ((v48) | 1))
                                                        } else {
                                                            v125 := GenS6{F0: false, F1: (v51 % (((v51 - v51)) | 1))}
                                                            if (0 == -1) {
                                                                v124 += (v124 >> ((209) & 63))
                                                            } else {
                                                                var v126 int8 = -67
                                                                var v127 uint16 = (v42 | (v42 & (v42 >> ((v48) & 63))))
                                                                var v128 uint16 = v42
                                                                _ = v126
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                            _ = v125
                                                        }
                                                        _ = v123
                                                        _ = v124
                                                    } else {
                                                        switch (fi54 + 32767) {
                                                        case 32767:
                                                            v129 := []int{0}
                                                            _ = v129
                                                        case -617:
                                                            var v130 int32 = genFn1()
                                                            _ = v130
                                                        case -1:
                                                            v48 = v48
                                                        default:
                                                            switch ("a" + ("hello" + "a")) {
                                                            case "\n\t":
                                                                if ((false || false) && ((v42 | 0) == 1)) {
                                                                    v48 = (v48 | 83)
                                                                    v131 := []int{}
                                                                    var v132 uint32 = uint32(v42)
                                                                    _ = v131
                                                                    _ = v132
                                                                }
                                                            case "hello":
                                                                var v133 bool = false
                                                                _ = v133
                                                            }
                                                        }
                                                        var v134 int8 = 42
                                                        if (126 >= 255) {
                                                            switch 42 {
                                                            case 0:
                                                                v135 := []int{fi54, fi54, 675, (fi113 << (((v48 >> ((212) & 63))) & 63))}
                                                                _ = v135
                                                            case 4294967295:
                                                                var v136 string = ("\n\t" + "foo bar")
                                                                _ = v136
                                                            case -617:
                                                                v50 = -2.25
                                                            }
                                                            switch "hello" {
                                                            case "a":
                                                                var v137 string = "a"
                                                                _ = v137
                                                            }
                                                            var v138 float32 = 1e10
                                                            _ = v138
                                                        }
                                                        _ = v134
                                                    }
                                                }
                                                _ = fi113
                                            }
                                            _ = v108
                                        } else {
                                            for fi139 := 0; fi139 < 4; fi139++ {
                                                if false {
                                                    var v140 int16 = 32767
                                                    _ = v140
                                                } else {
                                                    v141 := []int{}
                                                    if true {
                                                        v142 := GenS17{F0: 0, F1: -704, F2: 42}
                                                        _ = v142
                                                    }
                                                    fmt.Println(fi139, fi139)
                                                    _ = v141
                                                }
                                                _ = fi139
                                            }
                                            for fi143 := 0; fi143 < 7; fi143++ {
                                                var v144 int8 = 100
                                                fmt.Println(v42, fi54)
                                                v42 += v42
                                                _ = fi143
                                                _ = v144
                                            }
                                        }
                                    default:
                                        if (!(genFn1() < 1)) {
                                            var v145 float64 = v50
                                            fmt.Println(v48, v145, v145)
                                            _ = v145
                                        } else {
                                            v146 := []int{(fi54 ^ fi54), (fi54 * (^(fi54)))}
                                            v48 = (^(v48))
                                            for fi147 := 0; fi147 < 18; fi147++ {
                                                v3 -= (v3 - (v3 * -542))
                                                switch 32767 {
                                                case 4294967295:
                                                    v146 = append(v146, 9223372036854775807)
                                                case 1:
                                                    _ = len(v146)
                                                case 868:
                                                    if ((!(!((fi147 & 4294967295) == (fi54 ^ 1)))) && false) {
                                                        var v148 int32 = 151
                                                        _ = len(v1)
                                                        _ = v148
                                                    } else {
                                                        v149 := GenS3{F0: (true || (fi147 != fi54)), F1: 422}
                                                        _ = v149
                                                    }
                                                default:
                                                    var v150 float64 = ((v50 + float64((0.0 * -2.25))) / v50)
                                                    _ = v150
                                                }
                                                _ = fi147
                                            }
                                            _ = v146
                                        }
                                    }
                                } else {
                                    if (!true) {
                                        var v151 uint8 = 0
                                        switch (-(fi54)) {
                                        case 127:
                                            v152 := []int{-9223372036854775808, (fi54 << ((v48) & 63)), fi54}
                                            _ = v152
                                        case 32768:
                                            if (v50 <= v50) {
                                                if true {
                                                    v153 := []int{255, 42, -780}
                                                    _ = v153
                                                }
                                                for fi154 := 0; fi154 < 5; fi154++ {
                                                    if false {
                                                        switch fi54 {
                                                        case 683:
                                                            _ = len(v1)
                                                        }
                                                    }
                                                    switch "foo bar" {
                                                    case "":
                                                        switch int(v151) {
                                                        case 42:
                                                            v155 := []int{-274, fi54}
                                                            _ = v155
                                                        }
                                                    case "hello":
                                                        v42 += 42
                                                    case "foo bar":
                                                        v156 := []int{fi54}
                                                        _ = v156
                                                    default:
                                                        v42 = (^(v42))
                                                    }
                                                    var v157 uint = (v48 + v48)
                                                    var v158 int16 = 0
                                                    v159 := GenS18{F0: 1, F1: "a"}
                                                    var v160 uint8 = 50
                                                    _ = fi154
                                                    _ = v157
                                                    _ = v158
                                                    _ = v159
                                                    _ = v160
                                                }
                                            } else {
                                                v50 = v50
                                            }
                                        case -279:
                                            v161 := GenS19{F0: (v48 ^ (^(v48))), F1: 127}
                                            _ = v161
                                        }
                                        _ = v151
                                    } else {
                                        fmt.Println(v51)
                                        v162 := []int{(fi54 + (fi54 - (-(fi54)))), fi54, 0}
                                        _ = v162
                                    }
                                    var v163 int32 = 2147483647
                                    _ = v163
                                }
                            }
                        }
                        var v164 uint = (^(v48))
                        _ = v164
                    }
                    _ = fi54
                    _ = v55
                }
            }
        } else {
            var v165 int16 = -32768
            if true {
                if false {
                    v166 := []int{-1, 2147483647}
                    v3 += -949
                    for fi167 := 0; fi167 < 2; fi167++ {
                        switch (fi167 & (fi167 - 186)) {
                        case 32768:
                            switch (fi167 / ((fi167) | 1)) {
                            case -719:
                                v168 := GenS3{F0: true, F1: (v48 * 65535)}
                                _ = v168
                            case 2147483647:
                                var v169 int32 = 42
                                _ = v169
                            }
                        case 42:
                            v165 = v165
                        case 2147483647:
                            switch 32767 {
                            case 9223372036854775807:
                                if false {
                                    if len(v166) > 0 {
                                        v166[0] = (fi167 >> ((v48) & 63))
                                    }
                                    v170 := GenS16{F0: "", F1: 2}
                                    _ = v170
                                } else {
                                    switch (fi167 * (fi167 | (-(fi167)))) {
                                    case 12:
                                        for fi171 := 0; fi171 < 20; fi171++ {
                                            _ = len(v166)
                                            _ = fi171
                                        }
                                    }
                                }
                            case -475:
                                v48 = (v48 >> (((^(v48))) & 63))
                            }
                        }
                        _ = fi167
                    }
                    _ = v166
                }
                fmt.Println(v50, v165)
            }
            if (!(float32(float64((-2.25 + (0.1 * (((1.5 * 1e10) * (-2.25 + 1e10)) - 3.14159))))) > ((-2.25 / 3.14159) + 0.0))) {
                v172 := []int{4294967295, 858, 569, -9223372036854775808}
                v2[1] = 65535
                _ = v172
            } else {
                fmt.Println(v42, v42)
                if (16 < 137) {
                    for fi173 := 0; fi173 < 10; fi173++ {
                        fmt.Println(v165)
                        for fi174 := 0; fi174 < 10; fi174++ {
                            v50 = v50
                            _ = fi174
                        }
                        _ = fi173
                    }
                    switch 42 {
                    case 32767:
                        v42 = 1
                    case 127:
                        if false {
                            var v175 int16 = v165
                            _ = v175
                        } else {
                            v165 -= v165
                        }
                    default:
                        v176 := v2[2147483648]
                        _ = v176
                    }
                }
                var v177 int64 = 1
                _ = v177
            }
            _ = v165
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v48
    _ = v49
    _ = v50
    _ = v51
}
