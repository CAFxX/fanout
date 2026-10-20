package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int8; F2 int32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint32; F1 int16; F2 bool }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int64; F1 int; F2 uint32 }
type GenS9 struct { F0 uint64; F1 int16; F2 int8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 uint16 }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint16; F1 bool }
type GenS15 struct { F0 string }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint }
func genFn0(p0 uint16, p1 uint16, p2 bool) int8 {
    switch 0 {
    case 1:
        p1 = (^(p1))
    default:
        var v4 float64 = -2.25
        _ = v4
    }
    fmt.Println(p0, p0, p1)
    if ((p2 && p2) && false) {
        if (!p2) {
            v5 := []int{32767, 845}
            v6 := GenS1{F0: p0, F1: 32768, F2: 42}
            switch "a" {
            case "a":
                p2 = p2
            default:
                var v7 uint32 = 0
                _ = v7
            }
            _ = v5
            _ = v6
        }
        if (1 > 1) {
            var v8 string = (("" + "") + ("\n\t" + "hello"))
            _ = v8
        } else {
            p1 += p1
            var v9 uint32 = 42
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int16 = -1
                v9 -= (v9 / ((4294967295) | 1))
                v12 := GenS1{F0: p0, F1: (-(fi10)), F2: 0}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        }
    } else {
        var v13 uint = 65535
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 string, p1 string, p2 int64) int {
    var v14 uint32 = 760
    v14 = 67
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 32767
}

func main() {
    var v1 uint8 = 42
    var v2 string = (((("foo bar" + ("a" + "hello")) + "hello") + "\n\t") + "a")
    v3 := make(map[int]uint32)
    var v15 int = -382
    v1 = v1
    if true {
        v16 := GenS2{F0: 128}
        v2 = (v2 + v2)
        _ = v16
    }
    switch genFn1(v2, v2, -9223372036854775808) {
    case 127:
        for fi17 := 0; fi17 < 10; fi17++ {
            v1 = 223
            _ = fi17
        }
    case 9223372036854775807:
        v3[-249] = 2
    case 128:
        if (true && false) {
            if false {
                var v18 string = (v2 + v2)
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 uint64 = 976
                    fmt.Println(fi19, v15, fi19)
                    for fi21 := 0; fi21 < 6; fi21++ {
                        switch "a" {
                        case "a":
                            switch (^(v15)) {
                            case 32767:
                                fmt.Println(v15, v18, fi21)
                            case 32768:
                                if false {
                                    var v22 uint8 = 160
                                    var v23 float64 = (1e10 * 0.1)
                                    _ = v22
                                    _ = v23
                                } else {
                                    fmt.Println(v18)
                                    var v24 float32 = (((float32(float64((1e10 * 0.1))) / 1.0) / 1.5) / 1.5)
                                    v24 += float32(1e10)
                                    _ = v24
                                }
                            }
                        case "\n\t":
                            v25 := []int{}
                            _ = v25
                        case "foo bar":
                            v26 := GenS1{F0: 1, F1: (^(fi19)), F2: (^(v1))}
                            _ = v26
                        }
                        _ = fi21
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        }
    }
    var v27 int64 = 9223372036854775807
    if true {
        switch v2 {
        case "":
            _ = len(v3)
        case "hello":
            var v28 int8 = int8(v15)
            _ = v28
        }
    } else {
        if true {
            var v29 int8 = 42
            _ = v29
        } else {
            var v30 string = ((v2 + ("foo bar" + "")) + "a")
            _ = v30
        }
        delete(v3, 4294967295)
        var v31 bool = (128 <= v1)
        _ = v31
    }
    switch v2 {
    case "\n\t":
        var v32 uint = 0
        _ = v32
    case "a":
        if (0 >= 18446744073709551615) {
            v33 := []int{v15, v15, (v15 / (((-(v15))) | 1))}
            if ((0.1 != 3.14159) || (166 != 255)) {
                v2 = v2
            } else {
                var v34 uint64 = 1
                v35 := GenS3{F0: v2, F1: 127, F2: -2147483648}
                v15 += 4294967295
                _ = v34
                _ = v35
            }
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = v33
        }
    case "":
        v15 -= (v15 >> ((18446744073709551615) & 63))
    default:
        switch (v2 + v2) {
        case "a":
            v1 += (v1 % (((v1 % ((v1) | 1))) | 1))
        case "\n\t":
            var v36 float64 = 0.1
            _ = v36
        case "hello":
            switch 379 {
            case 2147483647:
                if true {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        fmt.Println(v2, v2)
                        v38 := GenS2{F0: (fi37 / ((32768) | 1))}
                        var v39 int32 = 42
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    var v40 uint8 = 127
                    switch (v2 + v2) {
                    case "foo bar":
                        fmt.Println(v27, v27)
                    default:
                        v27 = (^(v27))
                    }
                    _ = v40
                } else {
                    switch v15 {
                    case 4294967295:
                        if false {
                            switch v15 {
                            case -1:
                                switch (v2 + v2) {
                                case "\n\t":
                                    var v41 int = v15
                                    _ = v41
                                case "hello":
                                    if false {
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            fmt.Println(v1, v2, v1)
                                        case "hello":
                                            v42 := GenS1{F0: 360, F1: (-(v15)), F2: v1}
                                            _ = v42
                                        case "":
                                            if false {
                                                v43 := GenS2{F0: (v15 + v15)}
                                                _ = v43
                                            } else {
                                                var v44 uint64 = 42
                                                _ = v44
                                            }
                                        }
                                    } else {
                                        v45 := GenS3{F0: (("\n\t" + ("foo bar" + "")) + v2), F1: 42, F2: 42}
                                        if (false && (2 > 42)) {
                                            if ((((1e10 + 1e10) < 3.14159) && (4294967295 == 0)) && (("foo bar" + v2) < v2)) {
                                                switch v15 {
                                                case -2147483648:
                                                    var v46 bool = ((v27 == -826) && (!true))
                                                    _ = v46
                                                }
                                                v1 = (v1 & v1)
                                                v15 = (v15 ^ v15)
                                            } else {
                                                for fi47 := 0; fi47 < 7; fi47++ {
                                                    v48 := []int{428, fi47, (v15 + fi47), v15}
                                                    _ = fi47
                                                    _ = v48
                                                }
                                            }
                                            var v49 uint16 = 982
                                            _ = len(v3)
                                            _ = v49
                                        } else {
                                            v50 := GenS4{F0: 926, F1: 2}
                                            v27 = (v27 - (v27 | -1))
                                            switch (((v2 + "a") + ((v2 + v2) + v2)) + "hello") {
                                            case "":
                                                delete(v3, (v15 / (((v15 - genFn1("\n\t", "a", (-(v27))))) | 1)))
                                            case "\n\t":
                                                if (!false) {
                                                    v27 -= (-(v27))
                                                    var v51 uint64 = 42
                                                    _ = v51
                                                } else {
                                                    v52 := GenS1{F0: 0, F1: -927, F2: (v1 & (v1 * 128))}
                                                    _ = v52
                                                }
                                            default:
                                                if true {
                                                    fmt.Println(v27)
                                                    if true {
                                                        v53 := GenS2{F0: (-(v15))}
                                                        for fi54 := 0; fi54 < 17; fi54++ {
                                                            v55 := GenS2{F0: v15}
                                                            v56 := GenS5{F0: 2, F1: genFn0(382, 42, ((-118 >= 230) || true))}
                                                            _ = fi54
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        _ = v53
                                                    } else {
                                                        if (9223372036854775807 > v27) {
                                                            v27 -= (v27 * v27)
                                                            switch (v15 % (((v15 % ((v15) | 1))) | 1)) {
                                                            case 127:
                                                                v27 += (v27 >> ((0) & 63))
                                                            case 477:
                                                                v57 := []int{}
                                                                _ = v57
                                                            default:
                                                                v27 -= (v27 - (v27 << ((18446744073709551615) & 63)))
                                                            }
                                                            fmt.Println(v15)
                                                        } else {
                                                            v2 = "a"
                                                        }
                                                        var v58 int32 = 2147483647
                                                        _ = v58
                                                    }
                                                } else {
                                                    fmt.Println(v1, v27)
                                                    v59 := GenS6{F0: 2147483647, F1: -439, F2: (65535 >= 65535)}
                                                    _ = v59
                                                }
                                            }
                                            _ = v50
                                        }
                                        _ = v45
                                    }
                                default:
                                    v3[v15] = 1
                                }
                            case -9223372036854775808:
                                var v60 float64 = 0.1
                                _ = v60
                            }
                            switch v2 {
                            case "":
                                switch "a" {
                                case "":
                                    var v61 float64 = 0.0
                                    _ = v61
                                default:
                                    v62 := GenS7{F0: (v15 << ((1) & 63)), F1: 341}
                                    _ = v62
                                }
                            case "\n\t":
                                switch (v15 % (((^(v15))) | 1)) {
                                case -2147483648:
                                    fmt.Println(v27, v27)
                                case 32768:
                                    switch genFn1("\n\t", ("foo bar" + (v2 + ("foo bar" + v2))), 1) {
                                    case 127:
                                        var v63 uint16 = 163
                                        _ = v63
                                    case 2147483647:
                                        for fi64 := 0; fi64 < 19; fi64++ {
                                            for fi65 := 0; fi65 < 18; fi65++ {
                                                v1 += (v1 & v1)
                                                fmt.Println(v15, v15)
                                                v15 = int(v27)
                                                var v66 int8 = 127
                                                _ = fi65
                                                _ = v66
                                            }
                                            v1 += v1
                                            var v67 uint8 = (v1 & v1)
                                            _ = fi64
                                            _ = v67
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                default:
                                    var v68 bool = (((589 <= 385) || true) && false)
                                    _ = v68
                                }
                            case "hello":
                                if (!false) {
                                    var v69 uint = 727
                                    for fi70 := 0; fi70 < 18; fi70++ {
                                        if ((42 <= 42) && false) {
                                            v69 = (v69 - (v69 + uint(v1)))
                                            delete(v3, (fi70 + -290))
                                        }
                                        _ = fi70
                                    }
                                    v1 += v1
                                    _ = v69
                                } else {
                                    fmt.Println(v15)
                                }
                            default:
                                var v71 int16 = -32768
                                _ = v71
                            }
                            v72 := []int{}
                            _ = v72
                        }
                    }
                }
            }
        default:
            v1 = 197
        }
    }
    v73 := []int{genFn1("", v2, v27), (^(v15)), int(v1), genFn1("", "hello", v27)}
    for fi74 := 0; fi74 < 16; fi74++ {
        v75 := GenS7{F0: (v15 ^ 32768), F1: 581}
        if false {
            switch genFn1("\n\t", v2, (v27 & (v27 >> ((65535) & 63)))) {
            case 0:
                switch "hello" {
                case "foo bar":
                    for fi76 := 0; fi76 < 5; fi76++ {
                        v15 += genFn1(v2, v2, v27)
                        var v77 uint64 = 0
                        var v78 uint = 255
                        var v79 int16 = 478
                        if false {
                            if (!true) {
                                var v80 uint = v78
                                var v81 float64 = ((-2.25 / 1.5) + float64(0.1))
                                v82 := GenS6{F0: 1, F1: v79, F2: false}
                                if (false || (!true)) {
                                    v83 := GenS8{F0: (v27 & v27), F1: (v15 + fi74), F2: 0}
                                    if (!true) {
                                        switch v2 {
                                        case "foo bar":
                                            if (false || false) {
                                                v84 := GenS9{F0: v77, F1: (-(v79)), F2: -1}
                                                var v85 int8 = 1
                                                v86 := v3[v15]
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                        case "a":
                                            switch (fi76 / (((v15 / (((fi74 & fi76)) | 1))) | 1)) {
                                            case 0:
                                                var v87 uint = (v78 & v78)
                                                _ = v87
                                            default:
                                                v2 = v2
                                            }
                                        case "":
                                            fmt.Println(v2, v77)
                                        }
                                        if false {
                                            v1 = 2
                                        } else {
                                            v3[(-(fi76))] = 842
                                            v15 = -1
                                            var v88 bool = false
                                            _ = v88
                                        }
                                        v3[fi74] = 0
                                    }
                                    delete(v3, (^(fi76)))
                                    _ = v83
                                }
                                _ = v80
                                _ = v81
                                _ = v82
                            } else {
                                if ((v78 & v78) < (v78 * (v78 % ((255) | 1)))) {
                                    v1 = v1
                                }
                                if true {
                                    var v89 float32 = 3.14159
                                    var v90 int8 = genFn0(0, 279, (true || (2147483647 >= -699)))
                                    _ = v89
                                    _ = v90
                                } else {
                                    fmt.Println(fi74)
                                }
                            }
                        } else {
                            var v91 int8 = 42
                            _ = v91
                        }
                        _ = fi76
                        _ = v77
                        _ = v78
                        _ = v79
                    }
                case "hello":
                    v15 += (v15 - (v15 - fi74))
                case "a":
                    for fi92 := 0; fi92 < 13; fi92++ {
                        if true {
                            fmt.Println(v15)
                            v93 := GenS10{F0: (^(v27))}
                            if ((v27 - (v27 % ((v27) | 1))) >= 1) {
                                v15 += genFn1((("\n\t" + (v2 + ("foo bar" + ("\n\t" + "hello")))) + ((v2 + "") + "")), v2, v27)
                                fmt.Println(fi92, v15)
                            } else {
                                var v94 int8 = -128
                                var v95 uint8 = v1
                                if true {
                                    v3[(fi74 * -2147483648)] = 262
                                }
                                _ = v94
                                _ = v95
                            }
                            _ = v93
                        }
                        v96 := GenS5{F0: 65535, F1: 42}
                        if true {
                            v73 = append(v73, (fi74 >> ((4294967295) & 63)))
                            v97 := v3[fi92]
                            if true {
                                v98 := GenS11{F0: 42, F1: 65535}
                                if (!true) {
                                    switch v2 {
                                    case "":
                                        v99 := GenS12{F0: 47, F1: 0}
                                        _ = v99
                                    }
                                    switch v2 {
                                    case "\n\t":
                                        _ = len(v73)
                                    default:
                                        if false {
                                            switch v2 {
                                            case "":
                                                v97 = v97
                                            default:
                                                v97 -= (^(v97))
                                                if (true || false) {
                                                    v73 = append(v73, fi74)
                                                    var v100 float32 = 1.5
                                                    var v101 uint64 = 221
                                                    _ = v100
                                                    _ = v101
                                                } else {
                                                    fmt.Println(fi92, v97, v2)
                                                    v102 := GenS3{F0: ("hello" + (v2 + v2)), F1: -128, F2: 862}
                                                    v2 = v2
                                                    _ = v102
                                                }
                                                switch -462 {
                                                case -2147483648:
                                                    if len(v73) > 0 {
                                                        v73[0] = (^(v15))
                                                    }
                                                case 4294967295:
                                                    var v103 uint32 = (^(v97))
                                                    _ = v103
                                                }
                                            }
                                            v73 = append(v73, genFn1("", v2, (v27 ^ (^(v27)))))
                                            v2 = (v2 + "a")
                                        }
                                    }
                                    fmt.Println(fi74, fi74, fi74)
                                    var v104 int64 = (^(v27))
                                    _ = v104
                                } else {
                                    var v105 int16 = -791
                                    var v106 float32 = 1e10
                                    v107 := GenS13{F0: (v15 & genFn1("hello", v2, (v27 + (^(v27)))))}
                                    _ = v105
                                    _ = v106
                                    _ = v107
                                }
                                _ = v98
                            }
                            _ = v97
                        } else {
                            if (3.14159 > 1e10) {
                                v1 = (v1 - (v1 - v1))
                                if ((-316 > 776) || false) {
                                    fmt.Println(v1)
                                    if false {
                                        v2 = ((v2 + "foo bar") + "a")
                                    } else {
                                        var v108 uint64 = 4294967295
                                        var v109 uint32 = 65535
                                        if (!(("hello" + v2) >= "foo bar")) {
                                            if (-696 <= (-(v15))) {
                                                v2 = ""
                                            } else {
                                                var v110 int = 4294967295
                                                var v111 int32 = 1
                                                v112 := GenS14{F0: 703, F1: false}
                                                _ = v110
                                                _ = v111
                                                _ = v112
                                            }
                                            var v113 uint64 = uint64(v1)
                                            var v114 int32 = 42
                                            _ = v113
                                            _ = v114
                                        } else {
                                            if len(v73) > 0 {
                                                v73[0] = genFn1((v2 + v2), v2, (v27 / ((v27) | 1)))
                                            }
                                        }
                                        _ = v108
                                        _ = v109
                                    }
                                }
                            }
                            if false {
                                if false {
                                    var v115 int64 = v27
                                    var v116 float64 = 0.0
                                    _ = v115
                                    _ = v116
                                }
                                v117 := GenS3{F0: v2, F1: 100, F2: -174}
                                _ = v117
                            } else {
                                switch v2 {
                                case "":
                                    if (!(39 > 18446744073709551615)) {
                                        if false {
                                            var v118 int8 = genFn0(176, 2, false)
                                            var v119 bool = (!(42 > v27))
                                            _ = v118
                                            _ = v119
                                        } else {
                                            var v120 float64 = 1e10
                                            if false {
                                                v121 := GenS2{F0: (-(v15))}
                                                switch v2 {
                                                case "":
                                                    delete(v3, (v15 >> ((18446744073709551615) & 63)))
                                                case "a":
                                                    switch (v2 + (v2 + (v2 + v2))) {
                                                    case "a":
                                                        v120 = v120
                                                    }
                                                case "\n\t":
                                                    switch "a" {
                                                    case "a":
                                                        v1 -= 128
                                                        v122 := GenS15{F0: "foo bar"}
                                                        fmt.Println(v1, v1)
                                                        _ = v122
                                                    case "foo bar":
                                                        v73 = append(v73, (v15 >> ((138) & 63)))
                                                    case "\n\t":
                                                        v120 = ((float64(((0.0 * 3.14159) - float32(v120))) + float64(1e10)) / float64(3.14159))
                                                    }
                                                }
                                                _ = v121
                                            } else {
                                                var v123 float32 = (float32(1e10) / (3.14159 - 0.1))
                                                fmt.Println(fi92, v123)
                                                _ = v123
                                            }
                                            var v124 float64 = v120
                                            _ = v120
                                            _ = v124
                                        }
                                        if (372 > 4294967295) {
                                            if false {
                                                if true {
                                                    if false {
                                                        var v125 int64 = (v27 - (v27 | v27))
                                                        _ = v125
                                                    }
                                                    v126 := GenS16{F0: (v2 + "foo bar")}
                                                    _ = v126
                                                }
                                                v1 = (v1 ^ (v1 + v1))
                                                fmt.Println(v2)
                                            }
                                            switch -68 {
                                            case 127:
                                                v27 -= 9223372036854775807
                                            case -9223372036854775808:
                                                var v127 uint = 65535
                                                _ = v127
                                            default:
                                                var v128 int8 = -56
                                                _ = v128
                                            }
                                            v73 = append(v73, (-(fi92)))
                                        }
                                    }
                                }
                                v27 -= (-(v27))
                            }
                        }
                        _ = fi92
                        _ = v96
                    }
                }
            case -754:
                if false {
                    v129 := GenS2{F0: (-(fi74))}
                    v1 += v1
                    v130 := GenS17{F0: 1}
                    _ = v129
                    _ = v130
                }
            default:
                for fi131 := 0; fi131 < 6; fi131++ {
                    v27 = (^(v27))
                    _ = fi131
                }
            }
        }
        var v132 uint64 = 0
        _ = fi74
        _ = v132
        _ = v75
    }
    v3[genFn1(v2, ("a" + v2), (v27 - v27))] = 1
    for fi133 := 0; fi133 < 13; fi133++ {
        if (!(((float32(1e10) - ((0.1 + (((3.14159 + 1.5) + 0.0) / 1.0)) / (0.1 / ((0.0 + 1e10) + (0.1 / 0.1))))) + -2.25) < ((3.14159 - (3.14159 * 0.1)) / float32((-2.25 - 1.5))))) {
            v1 += (v1 * (v1 + v1))
        } else {
            for fi134 := 0; fi134 < 10; fi134++ {
                var v135 uint64 = 0
                _ = fi134
                _ = v135
            }
            fmt.Println(v1, v27)
        }
        if false {
            if (!((!(!true)) || false)) {
                for fi136 := 0; fi136 < 12; fi136++ {
                    switch "" {
                    case "a":
                        var v137 float64 = 1.5
                        var v138 uint16 = 496
                        var v139 int8 = 4
                        _ = v137
                        _ = v138
                        _ = v139
                    default:
                        var v140 uint8 = v1
                        _ = v140
                    }
                    _ = fi136
                }
                var v141 int32 = 42
                _ = v141
            } else {
                for fi142 := 0; fi142 < 19; fi142++ {
                    v27 = v27
                    _ = len(v3)
                    _ = fi142
                }
            }
            for fi143 := 0; fi143 < 15; fi143++ {
                v15 += 42
                switch (fi143 ^ v15) {
                case -415:
                    var v144 uint16 = 255
                    _ = v144
                default:
                    v1 += v1
                }
                _ = fi143
            }
        }
        v27 = (v27 * -604)
        _ = fi133
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v73:", len(v73), v73)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v73
}
