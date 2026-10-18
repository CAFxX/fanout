package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int8; F1 int64; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int32; F1 bool; F2 bool }
type GenS7 struct { F0 int64; F1 bool; F2 int }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 uint32; F1 int32; F2 int16 }
type GenS10 struct { F0 int }
func genFn0() int {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{fi5}
        var v7 float32 = (((float32(1e10) / (3.14159 + 1.5)) * -2.25) - 3.14159)
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{5, -676, 736, 32767}
    _ = v8
    return 32767
}
func genFn1(p0 string, p1 uint64) uint32 {
    v9 := []int{int(p1), 128, -683}
    if false {
        if len(v9) > 0 {
            v9[0] = 684
        }
        p0 = p0
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(p0, fi10)
        var v11 uint32 = 42
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]uint32)
    var v3 float32 = (0.0 * (((float32(-2.25) / 1.0) * (((3.14159 + 1.5) - (0.0 / -2.25)) * 1.5)) + ((1e10 + (1.5 - 1.5)) - 0.1)))
    var v4 int = 857
    fmt.Println(v4, v4, v4)
    v1 = (v1 % ((46) | 1))
    var v12 uint8 = 127
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    case "":
        v4 -= (^(v4))
    default:
        v3 += v3
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int8 = -103
        if true {
            v4 += fi13
            if true {
                fmt.Println(v14, fi13)
                v15 := GenS1{F0: (v12 & 127)}
                fmt.Println(v4, v1, v4)
                _ = v15
            } else {
                if ((v12 % ((128) | 1)) <= v12) {
                    var v16 uint16 = 32767
                    if true {
                        v17 := GenS2{F0: (v12 % ((v12) | 1)), F1: (v16 | 648), F2: "hello"}
                        _ = v17
                    }
                    _ = v16
                }
            }
        } else {
            v18 := GenS3{F0: (1000 != 0), F1: fi13}
            var v19 uint = 2
            v20 := GenS2{F0: (v12 - v12), F1: 0, F2: "foo bar"}
            _ = v18
            _ = v19
            _ = v20
        }
        switch ("foo bar" + "a") {
        case "hello":
            for fi21 := 0; fi21 < 10; fi21++ {
                v2[uint32(v12)] = 1
                v14 += (-(v1))
                if false {
                    var v22 int64 = 1
                    v23 := []int{fi21}
                    _ = v22
                    _ = v23
                }
                _ = fi21
            }
        case "a":
            v12 += 42
        }
        _ = fi13
        _ = v14
    }
    switch "a" {
    case "\n\t":
        for fi24 := 0; fi24 < 20; fi24++ {
            switch "\n\t" {
            case "hello":
                switch (v4 & v4) {
                case -924:
                    v25 := []int{genFn0()}
                    _ = v25
                }
            case "":
                if false {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := GenS2{F0: v12, F1: 255, F2: ""}
                        if false {
                            var v28 uint = 4294967295
                            _ = v28
                        } else {
                            var v29 int32 = 1
                            var v30 bool = (!false)
                            _ = v29
                            _ = v30
                        }
                        _ = fi26
                        _ = v27
                    }
                    var v31 uint = 255
                    v4 = 42
                    _ = v31
                }
            case "a":
                v12 = (v12 ^ 1)
            }
            var v32 float32 = float32(-2.25)
            v32 += v32
            _ = fi24
            _ = v32
        }
    case "":
        v33 := GenS2{F0: v12, F1: 42, F2: "hello"}
        _ = v33
    default:
        var v34 int16 = -344
        _ = v34
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 int64 = 42
        for fi37 := 0; fi37 < 14; fi37++ {
            switch "a" {
            case "hello":
                switch (fi35 ^ fi37) {
                case 0:
                    v38 := GenS3{F0: (!(!(!(v36 > v36)))), F1: v4}
                    _ = v38
                case -746:
                    switch (fi35 ^ (fi37 % ((-1) | 1))) {
                    case 1:
                        v1 = v1
                    case -387:
                        var v39 uint8 = 127
                        _ = v39
                    default:
                        v40 := []int{(v4 ^ genFn0()), (fi37 ^ (^(v4))), genFn0()}
                        _ = v40
                    }
                case 127:
                    fmt.Println(v36, v3)
                }
                fmt.Println(v4, fi37)
            case "\n\t":
                var v41 uint8 = v12
                _ = v41
            default:
                var v42 int = 552
                _ = v42
            }
            fmt.Println(v1)
            _ = fi37
        }
        for fi43 := 0; fi43 < 13; fi43++ {
            var v44 uint = 18446744073709551615
            fmt.Println(fi43, v36, fi35)
            if (1 > 2147483647) {
                v2[2] = 65535
                var v45 uint = v44
                v45 += (v44 | uint(v12))
                v45 = (^(v44))
                if (!(132 <= 18446744073709551615)) {
                    var v46 uint8 = 0
                    v47 := GenS1{F0: (v46 >> (((v45 % (((v44 / ((v45) | 1))) | 1))) & 63))}
                    v36 -= (v36 ^ 9223372036854775807)
                    _ = v46
                    _ = v47
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v48 float32 = (1.5 / 0.1)
                        _ = v48
                    default:
                        var v49 int16 = 32767
                        _ = v49
                    }
                    if false {
                        v50 := GenS3{F0: false, F1: (v4 & (fi43 * 1))}
                        v51 := []int{genFn0(), v4}
                        switch "hello" {
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                v52 := GenS1{F0: v12}
                                _ = v52
                            case "hello":
                                if ((^(v12)) < v12) {
                                    switch "foo bar" {
                                    case "":
                                        v45 = 255
                                    case "a":
                                        fmt.Println(fi43, v36, fi35)
                                    }
                                    delete(v2, genFn1((("" + ("foo bar" + "a")) + ""), 433))
                                    fmt.Println(fi35, v3, v1)
                                } else {
                                    fmt.Println(v12)
                                    v3 += v3
                                    v1 += (-(v1))
                                }
                            }
                        case "hello":
                            var v53 int8 = (-(v1))
                            _ = v53
                        case "\n\t":
                            var v54 int8 = v1
                            _ = v54
                        default:
                            var v55 int64 = v36
                            _ = v55
                        }
                        _ = v50
                        _ = v51
                    }
                    var v56 uint64 = 0
                    _ = v56
                }
                if true {
                    var v57 int16 = -1
                    if ((^(fi35)) == -2147483648) {
                        var v58 int8 = 42
                        _ = v58
                    } else {
                        fmt.Println(v3, v4)
                    }
                    _ = len(v2)
                    _ = v57
                } else {
                    switch "foo bar" {
                    case "a":
                        if (-32768 != 0) {
                            v59 := []int{}
                            _ = v59
                        } else {
                            var v60 int64 = v36
                            if true {
                                switch ("" + "a") {
                                case "hello":
                                    v61 := []int{127, genFn0()}
                                    _ = v61
                                case "a":
                                    v44 = v45
                                case "":
                                    switch (v4 + (v4 << ((18446744073709551615) & 63))) {
                                    case -2147483648:
                                        if true {
                                            switch fi35 {
                                            case -690:
                                                if ("hello" == ("hello" + "a")) {
                                                    if (621 > -1) {
                                                        v36 = (^(v36))
                                                        var v62 int = v4
                                                        var v63 bool = (v3 >= (v3 / 1e10))
                                                        var v64 string = "a"
                                                        delete(v2, uint32(v12))
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    if (1000 < 1000) {
                                                        var v65 uint32 = 65535
                                                        _ = v65
                                                    } else {
                                                        v66 := GenS4{F0: 21, F1: v36, F2: int64(v1)}
                                                        _ = v66
                                                    }
                                                } else {
                                                    if (!(659 <= 712)) {
                                                        v67 := GenS5{F0: genFn1("", 749), F1: 18446744073709551615}
                                                        var v68 uint8 = (v12 >> ((v45) & 63))
                                                        _ = v67
                                                        _ = v68
                                                    } else {
                                                        v69 := v2[2147483647]
                                                        var v70 int16 = 0
                                                        v69 = v69
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    v36 -= -2
                                                }
                                            case 1:
                                                var v71 int32 = 42
                                                switch "" {
                                                case "foo bar":
                                                    v72 := GenS6{F0: v71, F1: true, F2: (2 <= 540)}
                                                    _ = v72
                                                case "a":
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v44)
                                                    }
                                                case "":
                                                    if (!false) {
                                                        if false {
                                                            v45 += (^(v44))
                                                            var v73 uint64 = 42
                                                            _ = v73
                                                        } else {
                                                            v74 := []int{fi35, int(v60), v4, -9223372036854775808}
                                                            switch "a" {
                                                            case "a":
                                                                v4 -= (-(v4))
                                                            case "":
                                                                v12 = v12
                                                            }
                                                            if true {
                                                                if (v12 <= 128) {
                                                                    var v75 int32 = v71
                                                                    _ = v75
                                                                }
                                                                v74 = append(v74, (v4 % (((fi35 ^ 473)) | 1)))
                                                                v71 -= v71
                                                            } else {
                                                                v76 := GenS6{F0: int32(v45), F1: (!(v1 < (v1 % ((int8(v4)) | 1)))), F2: ((v71 << (((v45 >> ((v45) & 63))) & 63)) >= v71)}
                                                                if true {
                                                                    v12 = v12
                                                                    if (!false) {
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            var v77 int32 = (v71 % ((v71) | 1))
                                                                            delete(v2, 1)
                                                                            if (v44 <= (v45 - (v45 | (v45 % ((uint(v12)) | 1))))) {
                                                                                delete(v2, 42)
                                                                            }
                                                                            var v78 uint32 = 0
                                                                            var v79 float64 = (0.0 - 1.5)
                                                                            _ = v77
                                                                            _ = v78
                                                                            _ = v79
                                                                        default:
                                                                            v80 := GenS3{F0: false, F1: (fi35 | 182)}
                                                                            _ = v80
                                                                        }
                                                                        v45 -= (v45 - v45)
                                                                        v81 := GenS6{F0: 0, F1: false, F2: (668 > uint64(v12))}
                                                                        _ = v81
                                                                    } else {
                                                                        v60 = v60
                                                                        var v82 bool = (4294967295 != 1)
                                                                        _ = v82
                                                                    }
                                                                    switch ("foo bar" + (("foo bar" + "") + "foo bar")) {
                                                                    case "\n\t":
                                                                        var v83 int32 = int32(v36)
                                                                        _ = v83
                                                                    default:
                                                                        fmt.Println(v71, fi43, v44)
                                                                    }
                                                                } else {
                                                                    if ((^(v71)) != (v71 | (v71 - (^(v71))))) {
                                                                        v12 += v12
                                                                    }
                                                                    v84 := GenS4{F0: (v1 / ((v1) | 1)), F1: (v36 * v36), F2: (^(v60))}
                                                                    if true {
                                                                        v1 = v1
                                                                        var v85 int32 = 42
                                                                        _ = v85
                                                                    }
                                                                    var v86 uint32 = 4294967295
                                                                    v4 = genFn0()
                                                                    switch (("hello" + ((("" + ((("foo bar" + "hello") + ("foo bar" + "\n\t")) + "hello")) + "foo bar") + "hello")) + "") {
                                                                    case "":
                                                                        v1 = (v1 % (((-(v1))) | 1))
                                                                    }
                                                                    fmt.Println(v12)
                                                                    fmt.Println(v60, v36)
                                                                    _ = v84
                                                                    _ = v86
                                                                }
                                                                v4 = 119
                                                                _ = v76
                                                            }
                                                            _ = v74
                                                        }
                                                        v3 += v3
                                                    } else {
                                                        var v87 uint8 = (v12 >> ((v45) & 63))
                                                        _ = v87
                                                    }
                                                }
                                                _ = v71
                                            default:
                                                if true {
                                                    var v88 uint64 = 42
                                                    _ = v88
                                                } else {
                                                    v3 += 0.1
                                                    v89 := GenS2{F0: (^(v12)), F1: 255, F2: "\n\t"}
                                                    v60 = (v60 >> (((v44 | (v44 | 18446744073709551615))) & 63))
                                                    _ = v89
                                                }
                                            }
                                        } else {
                                            var v90 int16 = -649
                                            var v91 float64 = 1.5
                                            v60 += (v60 + int64(fi43))
                                            _ = v90
                                            _ = v91
                                        }
                                    default:
                                        v92 := GenS4{F0: v1, F1: (v36 + (v36 ^ (v60 ^ (v60 | (v60 % ((v36) | 1)))))), F2: v60}
                                        _ = v92
                                    }
                                }
                                var v93 int32 = -667
                                var v94 string = (("" + "a") + (("foo bar" + "a") + ("" + "foo bar")))
                                _ = v93
                                _ = v94
                            } else {
                                if ((v36 / ((1) | 1)) != 42) {
                                    var v95 uint32 = genFn1("foo bar", 62)
                                    var v96 int64 = (^(v36))
                                    v12 = 127
                                    v97 := GenS4{F0: (v1 & v1), F1: v96, F2: (-(v60))}
                                    switch (fi35 >> ((v44) & 63)) {
                                    case -828:
                                        v98 := []int{(fi35 + (v4 | (fi35 << ((v44) & 63)))), (v4 * genFn0())}
                                        _ = v98
                                    }
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                } else {
                                    if true {
                                        if (false || (-2.25 == (float64(3.14159) * 0.0))) {
                                            v36 -= (v36 << ((4294967295) & 63))
                                        } else {
                                            var v99 uint16 = 1
                                            fmt.Println(v44, v36, fi43)
                                            var v100 int8 = -1
                                            _ = v100
                                            _ = v99
                                        }
                                        var v101 bool = true
                                        v102 := GenS7{F0: -192, F1: (!(!v101)), F2: 42}
                                        _ = v101
                                        _ = v102
                                    }
                                }
                                if false {
                                    if true {
                                        v103 := GenS4{F0: v1, F1: v36, F2: (v36 / (((v36 % (((-(v36))) | 1))) | 1))}
                                        _ = v103
                                    } else {
                                        if true {
                                            fmt.Println(v4, fi35, v4)
                                        } else {
                                            v45 = (^(v44))
                                            v104 := GenS7{F0: (v60 - 383), F1: true, F2: fi43}
                                            v105 := []int{}
                                            _ = v104
                                            _ = v105
                                        }
                                    }
                                } else {
                                    if true {
                                        var v106 int = -538
                                        _ = v106
                                    } else {
                                        var v107 float64 = -2.25
                                        _ = v107
                                    }
                                    v108 := []int{genFn0(), v4, genFn0(), (fi43 >> ((4294967295) & 63))}
                                    v2[uint32(v44)] = 1
                                    v109 := GenS4{F0: v1, F1: v60, F2: v36}
                                    v108 = append(v108, -9223372036854775808)
                                    var v110 bool = true
                                    var v111 int8 = v1
                                    _ = v108
                                    _ = v109
                                    _ = v110
                                    _ = v111
                                }
                            }
                            _ = v60
                        }
                    case "hello":
                        v1 = (v1 & v1)
                    default:
                        v112 := []int{(fi43 - -2147483648), (fi35 / (((v4 + (fi43 + v4))) | 1)), fi35}
                        _ = v112
                    }
                    switch (v4 << ((v45) & 63)) {
                    case -9223372036854775808:
                        v113 := GenS8{F0: (-(v36)), F1: 793}
                        _ = v113
                    case -2147483648:
                        var v114 uint8 = 132
                        var v115 float32 = (1e10 - v3)
                        var v116 uint8 = (v12 & 42)
                        switch "\n\t" {
                        case "foo bar":
                            v117 := GenS6{F0: 42, F1: true, F2: true}
                            _ = v117
                        }
                        _ = v114
                        _ = v115
                        _ = v116
                    case 133:
                        var v118 uint16 = 42
                        _ = v118
                    default:
                        var v119 bool = true
                        var v120 uint8 = (^(v12))
                        var v121 string = ("" + "\n\t")
                        _ = v119
                        _ = v120
                        _ = v121
                    }
                }
                v2[786] = 555
                switch ((("" + "a") + "foo bar") + "hello") {
                case "hello":
                    if true {
                        v4 = (fi35 / ((int(v44)) | 1))
                        v45 = (v44 % ((uint(v12)) | 1))
                    }
                case "foo bar":
                    v36 = v36
                case "a":
                    delete(v2, uint32(v44))
                }
                if false {
                    v122 := []int{42, (fi35 << ((v45) & 63)), fi43, (-(v4))}
                    v123 := v2[65535]
                    _ = v122
                    _ = v123
                } else {
                    var v124 float32 = v3
                    v12 = 1
                    _ = v124
                }
                _ = v45
            } else {
                var v125 bool = true
                v1 += (v1 >> (((v44 - (v44 * 1))) & 63))
                _ = v125
            }
            _ = fi43
            _ = v44
        }
        _ = fi35
        _ = v36
    }
    var v126 uint64 = 1
    v127 := GenS2{F0: (v12 - v12), F1: 2, F2: ("" + ("" + "hello"))}
    v128 := GenS1{F0: (v12 + (v12 << ((1) & 63)))}
    if false {
        switch v4 {
        case 0:
            switch "hello" {
            case "foo bar":
                var v129 uint8 = (^(v12))
                _ = v129
            case "hello":
                if true {
                    var v130 uint32 = 151
                    fmt.Println(v4)
                    var v131 int32 = 0
                    _ = v130
                    _ = v131
                } else {
                    v126 += v126
                    v12 = (v12 & (v12 ^ v12))
                    var v132 float64 = -2.25
                    _ = v132
                }
            case "\n\t":
                switch "foo bar" {
                case "foo bar":
                    v12 -= (v12 ^ v12)
                }
            }
        default:
            switch "" {
            case "\n\t":
                switch "a" {
                case "hello":
                    switch (v4 << ((2) & 63)) {
                    case -110:
                        var v133 int16 = 797
                        _ = v133
                    default:
                        for fi134 := 0; fi134 < 9; fi134++ {
                            var v135 float64 = ((0.0 + 0.1) + 1e10)
                            v2[103] = 354
                            _ = fi134
                            _ = v135
                        }
                    }
                }
            case "a":
                if (((3.14159 * (0.1 + 0.0)) + 0.1) == (1e10 * 0.0)) {
                    if (v4 < v4) {
                        v136 := GenS9{F0: 65535, F1: 0, F2: 1}
                        _ = v136
                    } else {
                        for fi137 := 0; fi137 < 19; fi137++ {
                            if (!((true || false) || (((float32(0.1) - 1.5) / v3) < v3))) {
                                var v138 int8 = v1
                                _ = v138
                            }
                            switch ("" + ("a" + "hello")) {
                            case "a":
                                var v139 float32 = v3
                                _ = v139
                            case "\n\t":
                                v140 := GenS10{F0: genFn0()}
                                _ = v140
                            default:
                                v126 = 1
                            }
                            var v141 uint64 = v126
                            _ = fi137
                            _ = v141
                        }
                        if (v12 < (^(v12))) {
                            v12 = (v12 & v12)
                            var v142 uint = 137
                            _ = v142
                        } else {
                            var v143 int8 = (v1 / (((v1 & -11)) | 1))
                            v3 -= (v3 / float32(1e10))
                            if true {
                                for fi144 := 0; fi144 < 7; fi144++ {
                                    v4 = fi144
                                    for fi145 := 0; fi145 < 16; fi145++ {
                                        v12 = (v12 - 127)
                                        v146 := GenS6{F0: -598, F1: false, F2: ((^(v12)) == 41)}
                                        var v147 bool = true
                                        _ = fi145
                                        _ = v146
                                        _ = v147
                                    }
                                    var v148 int8 = (v143 * v1)
                                    _ = fi144
                                    _ = v148
                                }
                                v149 := []int{(v4 >> ((401) & 63)), v4, (v4 / (((v4 >> ((579) & 63))) | 1)), (v4 >> ((255) & 63))}
                                var v150 uint = 0
                                _ = v149
                                _ = v150
                            } else {
                                var v151 float32 = (1e10 + v3)
                                for fi152 := 0; fi152 < 18; fi152++ {
                                    fmt.Println(v12)
                                    var v153 int = 996
                                    _ = fi152
                                    _ = v153
                                }
                                v12 = v12
                                _ = v151
                            }
                            _ = v143
                        }
                    }
                }
            case "foo bar":
                v1 += v1
            default:
                if false {
                    var v154 int32 = int32(v126)
                    _ = v154
                } else {
                    _ = len(v2)
                }
            }
        }
        switch -512 {
        case 389:
            if true {
                _ = len(v2)
            }
        case -1:
            if (genFn1("", uint64(v12)) != 0) {
                v155 := GenS3{F0: true, F1: (^(v4))}
                v3 += 0.1
                v3 = v3
                _ = v155
            } else {
                var v156 float32 = (float32(3.14159) + v3)
                _ = v156
            }
        case 2147483647:
            v4 += (v4 & v4)
        }
    } else {
        if (2 > 42) {
            v12 -= 1
            var v157 uint8 = (v12 * v12)
            if true {
                delete(v2, 1)
            } else {
                v4 = (v4 - v4)
                v158 := []int{(v4 & (v4 ^ genFn0()))}
                v3 -= 1e10
                _ = v158
            }
            _ = v157
        }
    }
    if true {
        for fi159 := 0; fi159 < 3; fi159++ {
            var v160 int64 = -9223372036854775808
            _ = fi159
            _ = v160
        }
        v161 := GenS6{F0: -2147483648, F1: (0 < (v1 >> ((4294967295) & 63))), F2: false}
        v162 := []int{genFn0(), v4}
        _ = v161
        _ = v162
    } else {
        fmt.Println(v1)
        var v163 float64 = 1e10
        _ = v163
    }
    v164 := v2[286]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v126:", v126)
    fmt.Println("v164:", v164)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v126
    _ = v127
    _ = v128
    _ = v164
    _ = v2
    _ = v3
    _ = v4
}
