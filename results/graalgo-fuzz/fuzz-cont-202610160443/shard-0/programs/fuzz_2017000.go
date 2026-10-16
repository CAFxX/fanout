package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint; F1 bool; F2 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS12 struct { F0 int32; F1 uint64; F2 string }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int32; F1 int32 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int32; F1 uint; F2 int64 }
type GenS17 struct { F0 uint32; F1 bool }
type GenS18 struct { F0 int32; F1 string }
type GenS19 struct { F0 int8; F1 int16; F2 int16 }
type GenS20 struct { F0 bool; F1 uint; F2 uint32 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 uint16; F1 uint8 }
type GenS24 struct { F0 int16 }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 int }
func genFn0(p0 uint8, p1 int64, p2 uint32) uint64 {
    p2 = (p2 | 0)
    v6 := []int{}
    p2 += (p2 >> ((18446744073709551615) & 63))
    var v7 int32 = -324
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -1
    var v3 uint16 = 0
    var v4 uint8 = 0
    var v5 float32 = 0.1
    v2 += (v2 - (v2 << ((2) & 63)))
    if false {
        if false {
            if ((^(v4)) != 42) {
                if (0 >= 2) {
                    v8 := GenS1{F0: 799, F1: v4, F2: 46}
                    switch -9223372036854775808 {
                    case 0:
                        v5 = 1.5
                    case 32767:
                        for fi9 := 0; fi9 < 19; fi9++ {
                            switch ("" + ("" + "\n\t")) {
                            case "a":
                                v10 := GenS1{F0: 0, F1: v4, F2: 100}
                                _ = v10
                            case "hello":
                                var v11 uint16 = v3
                                _ = v11
                            }
                            for fi12 := 0; fi12 < 1; fi12++ {
                                switch 32768 {
                                case 42:
                                    var v13 float64 = float64((float32((1.5 * 0.0)) * (v5 / ((v5 + v5) - v5))))
                                    _ = v13
                                case 32767:
                                    if ((-1 > (v2 << ((255) & 63))) && (v2 <= v2)) {
                                        v14 := GenS2{F0: 470, F1: 215, F2: fi9}
                                        _ = v14
                                    } else {
                                        v15 := GenS2{F0: 2, F1: 2, F2: (fi9 << ((1) & 63))}
                                        v2 -= (-(v2))
                                        _ = v15
                                    }
                                case 0:
                                    v1[4294967295] = fi9
                                default:
                                    v16 := GenS1{F0: v3, F1: (v4 * 42), F2: -30}
                                    _ = v16
                                }
                                if false {
                                    v17 := GenS1{F0: v3, F1: 23, F2: -1}
                                    if (false && ((0 < 362) && false)) {
                                        switch "hello" {
                                        case "":
                                            var v18 uint64 = 911
                                            var v19 int = (fi12 | fi12)
                                            fmt.Println(v5, v3)
                                            v2 += -609
                                            var v20 int8 = -128
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        case "hello":
                                            v3 += v3
                                        case "a":
                                            var v21 string = ""
                                            _ = v21
                                        default:
                                            v22 := []int{}
                                            _ = v22
                                        }
                                        v23 := GenS3{F0: 42}
                                        if (v5 < (((0.0 * v5) / v5) * (v5 + float32(((1e10 + 0.1) / 1e10))))) {
                                            v3 = 2
                                            v24 := GenS4{F0: genFn0(v4, 9223372036854775807, 350)}
                                            if (false || true) {
                                                v25 := GenS5{F0: -1, F1: 42, F2: -32768}
                                                v26 := GenS4{F0: genFn0((v4 - (v4 % ((0) | 1))), -9223372036854775808, 2)}
                                                v27 := GenS3{F0: 42}
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            } else {
                                                if false {
                                                    v2 = v2
                                                }
                                                _ = len(v1)
                                                switch ("" + "hello") {
                                                case "a":
                                                    if (((!(genFn0(v4, -111, 793) < 0)) || (4294967295 >= 4294967295)) || (("" + "hello") >= ("" + ("foo bar" + "a")))) {
                                                        fmt.Println(fi12)
                                                    }
                                                case "hello":
                                                    v28 := GenS6{F0: v4}
                                                    _ = v28
                                                default:
                                                    if (("foo bar" + ("a" + "hello")) == "foo bar") {
                                                        if ((^(v3)) != v3) {
                                                            v29 := GenS7{F0: 0, F1: 42}
                                                            v30 := GenS8{F0: 65535, F1: ((!true) && false), F2: 42}
                                                            v31 := GenS9{F0: 2}
                                                            _ = v29
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        v2 -= v2
                                                        if false {
                                                            if (0.0 != 0.1) {
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    fmt.Println(v3, v4)
                                                                case "foo bar":
                                                                    v32 := GenS2{F0: 42, F1: 618, F2: 127}
                                                                    _ = v32
                                                                case "hello":
                                                                    v4 = v4
                                                                default:
                                                                    v2 -= (v2 / (((^(v2))) | 1))
                                                                }
                                                                v2 = -32768
                                                                if false {
                                                                    var v33 int = (fi9 >> ((255) & 63))
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        var v34 int = (-(fi9))
                                                                        _ = v34
                                                                    case "hello":
                                                                        v35 := v1[133]
                                                                        _ = v35
                                                                    case "a":
                                                                        v36 := []int{v33, -858, fi12}
                                                                        _ = v36
                                                                    default:
                                                                        var v37 int16 = v2
                                                                        _ = v37
                                                                    }
                                                                    _ = len(v1)
                                                                    _ = v33
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v24
                                        }
                                        _ = v23
                                    } else {
                                        fmt.Println(v2)
                                    }
                                    v38 := GenS10{F0: 255, F1: 558}
                                    _ = v17
                                    _ = v38
                                }
                                v2 -= v2
                                _ = fi12
                            }
                            _ = fi9
                        }
                    case 111:
                        v1[118] = 1
                    default:
                        if false {
                            for fi39 := 0; fi39 < 14; fi39++ {
                                v4 = v4
                                v3 -= v3
                                if false {
                                    v40 := GenS11{F0: 18446744073709551615, F1: (v4 << ((42) & 63)), F2: 116}
                                    var v41 int32 = -1
                                    v2 = -1
                                    _ = v40
                                    _ = v41
                                } else {
                                    fmt.Println(v2, fi39, v2)
                                    if (0.1 != ((float64(((v5 / v5) * 0.1)) / 1.5) - -2.25)) {
                                        var v42 int = fi39
                                        fmt.Println(v5, v42, v4)
                                        v3 += v3
                                        _ = v42
                                    } else {
                                        for fi43 := 0; fi43 < 7; fi43++ {
                                            var v44 uint16 = (v3 - (v3 >> ((65535) & 63)))
                                            var v45 uint64 = 759
                                            _ = fi43
                                            _ = v44
                                            _ = v45
                                        }
                                        v4 = 0
                                    }
                                }
                                _ = fi39
                            }
                        } else {
                            switch "a" {
                            case "":
                                for fi46 := 0; fi46 < 20; fi46++ {
                                    switch fi46 {
                                    case 2147483647:
                                        v4 = 2
                                    }
                                    var v47 int32 = -2147483648
                                    _ = fi46
                                    _ = v47
                                }
                            case "a":
                                v48 := []int{4294967295, 2147483647, 32767}
                                _ = v48
                            case "hello":
                                v49 := GenS5{F0: 1, F1: 921, F2: (v2 - v2)}
                                _ = v49
                            }
                            fmt.Println(v2, v5, v2)
                        }
                    }
                    v50 := GenS8{F0: 0, F1: false, F2: 42}
                    _ = v50
                    _ = v8
                }
                v51 := GenS12{F0: 2147483647, F1: 4294967295, F2: "hello"}
                _ = v51
            } else {
                fmt.Println(v2, v5, v5)
                fmt.Println(v3, v4, v5)
            }
            v2 += (v2 & 286)
        } else {
            fmt.Println(v5)
            v52 := GenS9{F0: 42}
            _ = v52
        }
        v4 = (v4 & (v4 ^ v4))
        var v53 uint32 = 155
        _ = v53
    }
    var v54 uint64 = genFn0((v4 ^ (v4 + 255)), -1, 2147483647)
    v55 := GenS13{F0: false}
    if (1 < 0) {
        if (3.14159 <= float64(v5)) {
            for fi56 := 0; fi56 < 16; fi56++ {
                v54 -= v54
                v5 = v5
                _ = fi56
            }
        }
    } else {
        for fi57 := 0; fi57 < 13; fi57++ {
            switch 0 {
            case 128:
                for fi58 := 0; fi58 < 4; fi58++ {
                    var v59 uint64 = genFn0(42, 341, 0)
                    delete(v1, 268)
                    _ = fi58
                    _ = v59
                }
            case 1:
                v54 = v54
            case 136:
                var v60 int32 = 284
                _ = v60
            default:
                v3 = 2
            }
            var v61 int16 = 0
            v54 = genFn0(v4, 42, 934)
            _ = fi57
            _ = v61
        }
        if (false && (true && true)) {
            if false {
                var v62 uint32 = 2147483647
                if false {
                    var v63 int = 127
                    var v64 int = (v63 - v63)
                    _ = v63
                    _ = v64
                } else {
                    if false {
                        var v65 int32 = -762
                        _ = v65
                    } else {
                        var v66 int32 = 532
                        v67 := GenS14{F0: int32(v62), F1: v66}
                        if (!true) {
                            for fi68 := 0; fi68 < 12; fi68++ {
                                var v69 float64 = (0.0 - 1.5)
                                _ = fi68
                                _ = v69
                            }
                            if false {
                                delete(v1, v62)
                                if true {
                                    if false {
                                        v70 := GenS6{F0: (v4 >> ((2) & 63))}
                                        _ = v70
                                    }
                                }
                                var v71 uint16 = v3
                                _ = v71
                            } else {
                                for fi72 := 0; fi72 < 19; fi72++ {
                                    v2 += (v2 / ((v2) | 1))
                                    switch "a" {
                                    case "a":
                                        switch 32768 {
                                        case -680:
                                            if (((("\n\t" + ("hello" + "foo bar")) + "foo bar") <= "") || false) {
                                                if true {
                                                    var v73 uint64 = 0
                                                    var v74 int8 = 127
                                                    _ = v73
                                                    _ = v74
                                                }
                                                for fi75 := 0; fi75 < 8; fi75++ {
                                                    v76 := []int{fi75, fi75}
                                                    _ = fi75
                                                    _ = v76
                                                }
                                                var v77 uint = 4294967295
                                                _ = v77
                                            }
                                        case 42:
                                            fmt.Println(v62, v54)
                                        case -344:
                                            v2 = (v2 & (v2 | -32768))
                                        }
                                    case "hello":
                                        if false {
                                            v4 = (v4 & 42)
                                            v78 := []int{}
                                            _ = v78
                                        } else {
                                            if true {
                                                var v79 uint32 = (v62 * 602)
                                                var v80 string = ("hello" + "foo bar")
                                                _ = v79
                                                _ = v80
                                            }
                                        }
                                    default:
                                        var v81 int16 = 42
                                        _ = v81
                                    }
                                    _ = fi72
                                }
                                switch -2147483648 {
                                case 0:
                                    var v82 float32 = v5
                                    _ = v82
                                case 575:
                                    if (false && true) {
                                        v83 := GenS14{F0: (v66 << ((0) & 63)), F1: (-(v66))}
                                        fmt.Println(v4, v3)
                                        _ = v83
                                    } else {
                                        v4 = (v4 & v4)
                                        v84 := []int{128}
                                        if (!((v3 | (v3 | v3)) >= 131)) {
                                            var v85 int64 = 9223372036854775807
                                            var v86 uint8 = (^(v4))
                                            var v87 string = "hello"
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        } else {
                                            var v88 int = 856
                                            for fi89 := 0; fi89 < 8; fi89++ {
                                                v66 += (-(v66))
                                                var v90 bool = false
                                                if (true && (!v90)) {
                                                    v88 = -220
                                                }
                                                _ = fi89
                                                _ = v90
                                            }
                                            fmt.Println(v62, v62, v66)
                                            _ = v88
                                        }
                                        _ = v84
                                    }
                                case 128:
                                    for fi91 := 0; fi91 < 10; fi91++ {
                                        switch (fi91 << ((4294967295) & 63)) {
                                        case -9223372036854775808:
                                            v1[(v62 | (^(v62)))] = (fi91 & (fi91 - -604))
                                        case 32767:
                                            for fi92 := 0; fi92 < 15; fi92++ {
                                                v2 = 0
                                                if ((v2 & v2) == (v2 + v2)) {
                                                    v5 += v5
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v4 = 120
                                                    default:
                                                        v93 := []int{}
                                                        _ = v93
                                                    }
                                                    switch fi92 {
                                                    case 127:
                                                        v94 := GenS5{F0: int32(v2), F1: 42, F2: v2}
                                                        _ = v94
                                                    case 32768:
                                                        if (65535 == 855) {
                                                            var v95 int8 = 100
                                                            switch fi91 {
                                                            case 1:
                                                                v96 := GenS15{F0: v4}
                                                                _ = v96
                                                            }
                                                            v3 -= v3
                                                            var v97 float64 = (((0.0 * 0.1) * 1.5) - (((-2.25 + 3.14159) + 1.5) * (1e10 - 1.5)))
                                                            v62 = 943
                                                            v66 += (^(v66))
                                                            if false {
                                                                if false {
                                                                    v98 := v1[v62]
                                                                    var v99 uint8 = v4
                                                                    v100 := []int{fi92, fi91}
                                                                    if (false || (int16(v3) >= v2)) {
                                                                        v100 = append(v100, -630)
                                                                    }
                                                                    if (42 == 511) {
                                                                        v5 = v5
                                                                    }
                                                                    if (v62 > (v62 | (^(v62)))) {
                                                                        v101 := GenS16{F0: (v66 % ((v66) | 1)), F1: 42, F2: int64(v54)}
                                                                        _ = v101
                                                                    }
                                                                    _ = v100
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "a":
                                                                v102 := GenS14{F0: (-(v66)), F1: v66}
                                                                _ = v102
                                                            }
                                                            switch "" {
                                                            case "foo bar":
                                                                v103 := GenS8{F0: 42, F1: true, F2: -1}
                                                                _ = v103
                                                            case "hello":
                                                                if (32767 > v2) {
                                                                    var v104 string = ""
                                                                    v97 = 1.5
                                                                    _ = v104
                                                                }
                                                            }
                                                            if ((v62 << ((958) & 63)) < v62) {
                                                                if false {
                                                                    var v105 uint8 = v4
                                                                    v106 := []int{(fi91 ^ fi91)}
                                                                    _ = v105
                                                                    _ = v106
                                                                } else {
                                                                    v4 -= 5
                                                                    if true {
                                                                        var v107 uint16 = 65535
                                                                        fmt.Println(v4)
                                                                        v5 = 0.0
                                                                        _ = v107
                                                                    }
                                                                    v54 += 18446744073709551615
                                                                }
                                                                v66 -= v66
                                                            } else {
                                                                v108 := GenS13{F0: (358 == 162)}
                                                                delete(v1, v62)
                                                                var v109 uint16 = v3
                                                                if (v3 < 924) {
                                                                    v54 += v54
                                                                    if ((v66 >> ((4294967295) & 63)) < int32(v2)) {
                                                                        v110 := v1[(^(v62))]
                                                                        if true {
                                                                            v5 -= (3.14159 / 3.14159)
                                                                            delete(v1, 2)
                                                                        } else {
                                                                            var v111 string = ""
                                                                            _ = v111
                                                                        }
                                                                        v97 -= 1e10
                                                                        var v112 int32 = v66
                                                                        _ = v110
                                                                        _ = v112
                                                                    }
                                                                    v113 := GenS17{F0: (^(v62)), F1: false}
                                                                    _ = v113
                                                                } else {
                                                                    v5 -= float32(3.14159)
                                                                    if true {
                                                                        if false {
                                                                            v114 := GenS5{F0: v66, F1: 683, F2: (v2 - (v2 / (((v2 ^ v2)) | 1)))}
                                                                            v115 := []int{fi92, fi91, fi91, (fi91 | fi92)}
                                                                            _ = v114
                                                                            _ = v115
                                                                        }
                                                                        var v116 float32 = v5
                                                                        _ = v116
                                                                    } else {
                                                                        v117 := GenS5{F0: v66, F1: 4294967295, F2: v2}
                                                                        v54 += (v54 % ((v54) | 1))
                                                                        _ = v117
                                                                    }
                                                                    v54 += genFn0(uint8(v109), -1, (v62 + 4294967295))
                                                                }
                                                                var v118 uint8 = 255
                                                                _ = v108
                                                                _ = v109
                                                                _ = v118
                                                            }
                                                            _ = v95
                                                            _ = v97
                                                        } else {
                                                            v54 -= (v54 ^ (v54 + (v54 - (v54 + 385))))
                                                            fmt.Println(v54, fi91)
                                                        }
                                                    case 128:
                                                        switch "\n\t" {
                                                        case "":
                                                            switch "a" {
                                                            case "hello":
                                                                v62 += v62
                                                            default:
                                                                var v119 int8 = -128
                                                                _ = v119
                                                            }
                                                        case "hello":
                                                            v120 := GenS18{F0: (v66 * (v66 + v66)), F1: ""}
                                                            _ = v120
                                                        case "foo bar":
                                                            v62 = (^(v62))
                                                            v4 = v4
                                                            v62 = 593
                                                        }
                                                    }
                                                    v62 += v62
                                                } else {
                                                    var v121 string = (("\n\t" + "") + "foo bar")
                                                    _ = len(v1)
                                                    _ = v121
                                                }
                                                if (false && false) {
                                                    v62 = (^(v62))
                                                }
                                                _ = fi92
                                            }
                                        case -184:
                                            switch "a" {
                                            case "a":
                                                var v122 int16 = 0
                                                _ = v122
                                            case "foo bar":
                                                v1[(^(v62))] = fi91
                                            }
                                        }
                                        if false {
                                            if false {
                                                v66 -= int32(v3)
                                            }
                                            fmt.Println(v4, v2, v66)
                                            if (v62 == (v62 / ((v62) | 1))) {
                                                for fi123 := 0; fi123 < 14; fi123++ {
                                                    fmt.Println(v3, fi123, v5)
                                                    var v124 int = 32767
                                                    _ = fi123
                                                    _ = v124
                                                }
                                                for fi125 := 0; fi125 < 14; fi125++ {
                                                    v62 -= 0
                                                    var v126 string = "foo bar"
                                                    _ = fi125
                                                    _ = v126
                                                }
                                                v5 = v5
                                            }
                                        } else {
                                            v127 := GenS18{F0: v66, F1: (("hello" + ("hello" + "")) + "\n\t")}
                                            var v128 float64 = 1e10
                                            v66 += v66
                                            _ = v127
                                            _ = v128
                                        }
                                        _ = fi91
                                    }
                                }
                                var v129 float64 = 1e10
                                _ = v129
                            }
                        }
                        _ = v66
                        _ = v67
                    }
                    v1[(^(v62))] = -30
                    for fi130 := 0; fi130 < 18; fi130++ {
                        switch ("hello" + "hello") {
                        case "a":
                            v131 := []int{}
                            _ = v131
                        case "foo bar":
                            v132 := GenS1{F0: (v3 << ((1) & 63)), F1: (v4 << ((206) & 63)), F2: int8(v3)}
                            _ = v132
                        case "\n\t":
                            fmt.Println(fi130, v2, v62)
                        }
                        _ = fi130
                    }
                }
                _ = v62
            } else {
                var v133 uint32 = 920
                if (false && true) {
                    var v134 int64 = 1
                    for fi135 := 0; fi135 < 16; fi135++ {
                        var v136 uint64 = uint64(v133)
                        v2 = v2
                        switch 0 {
                        case 1:
                            var v137 int32 = -910
                            _ = v137
                        case 9223372036854775807:
                            switch ((("\n\t" + "hello") + "foo bar") + "a") {
                            case "a":
                                for fi138 := 0; fi138 < 3; fi138++ {
                                    var v139 int64 = -1
                                    _ = fi138
                                    _ = v139
                                }
                            case "hello":
                                for fi140 := 0; fi140 < 18; fi140++ {
                                    if false {
                                        if true {
                                            if true {
                                                v136 += genFn0(v4, v134, (v133 >> ((4294967295) & 63)))
                                                switch (^(fi135)) {
                                                case 128:
                                                    var v141 int8 = -69
                                                    _ = v141
                                                case 9223372036854775807:
                                                    v142 := GenS19{F0: 1, F1: v2, F2: -32768}
                                                    _ = v142
                                                default:
                                                    var v143 int16 = int16(v4)
                                                    _ = v143
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    v4 += v4
                                                    v144 := GenS14{F0: 42, F1: 217}
                                                    var v145 int64 = (v134 << ((589) & 63))
                                                    _ = v144
                                                    _ = v145
                                                case "a":
                                                    v146 := []int{359}
                                                    _ = v146
                                                default:
                                                    v3 = (v3 * (v3 << ((4294967295) & 63)))
                                                }
                                                v147 := GenS11{F0: (v54 >> ((42) & 63)), F1: v4, F2: -524}
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v148 := v1[(v133 * (^(v133)))]
                                                    _ = v148
                                                case "":
                                                    v149 := GenS20{F0: (-2147483648 < 0), F1: uint(v3), F2: v133}
                                                    _ = v149
                                                default:
                                                    var v150 int = fi140
                                                    _ = v150
                                                }
                                                _ = v147
                                            } else {
                                                var v151 bool = true
                                                v152 := GenS21{F0: v2}
                                                _ = v151
                                                _ = v152
                                            }
                                            var v153 string = "foo bar"
                                            if false {
                                                v134 -= v134
                                            }
                                            _ = v153
                                        } else {
                                            _ = len(v1)
                                            v134 = 42
                                            v4 -= (^(v4))
                                        }
                                        fmt.Println(v134, fi140, fi140)
                                    }
                                    _ = fi140
                                }
                            case "\n\t":
                                v154 := []int{(fi135 >> ((42) & 63)), fi135, -822}
                                _ = v154
                            }
                        case 172:
                            for fi155 := 0; fi155 < 9; fi155++ {
                                var v156 uint8 = 42
                                var v157 int = fi155
                                v158 := GenS15{F0: v4}
                                switch "\n\t" {
                                case "a":
                                    var v159 float64 = 1e10
                                    _ = v159
                                case "":
                                    var v160 bool = false
                                    _ = v160
                                default:
                                    v161 := v1[v133]
                                    _ = v161
                                }
                                _ = len(v1)
                                _ = fi155
                                _ = v156
                                _ = v157
                                _ = v158
                            }
                        default:
                            v162 := GenS14{F0: -1, F1: -635}
                            _ = v162
                        }
                        _ = fi135
                        _ = v136
                    }
                    _ = v134
                }
                _ = v133
            }
        } else {
            v163 := GenS7{F0: 1, F1: 72}
            if (42 == -1) {
                for fi164 := 0; fi164 < 10; fi164++ {
                    switch (fi164 | fi164) {
                    case -642:
                        if false {
                            for fi165 := 0; fi165 < 9; fi165++ {
                                v166 := GenS21{F0: 1000}
                                _ = fi165
                                _ = v166
                            }
                            var v167 bool = (-128 > 109)
                            _ = v167
                        } else {
                            v3 -= (^(v3))
                        }
                    case 33:
                        if (false && (false && false)) {
                            if ((^(v54)) <= v54) {
                                v1[880] = 42
                                fmt.Println(v54)
                                v3 = v3
                            }
                            v2 += 42
                        }
                    default:
                        switch fi164 {
                        case 32767:
                            var v168 string = "\n\t"
                            _ = v168
                        case 4294967295:
                            _ = len(v1)
                        case -421:
                            v169 := GenS22{F0: "foo bar"}
                            _ = v169
                        }
                    }
                    var v170 int = (fi164 - fi164)
                    v2 = 52
                    _ = fi164
                    _ = v170
                }
                fmt.Println(v3)
                v5 -= v5
            } else {
                v4 = v4
                if false {
                    var v171 int16 = (v2 / ((v2) | 1))
                    v172 := GenS23{F0: v3, F1: (v4 - (^(v4)))}
                    if false {
                        var v173 uint64 = (v54 / (((v54 & 18446744073709551615)) | 1))
                        v174 := []int{}
                        switch -10 {
                        case -126:
                            var v175 int16 = (v171 | (v2 - v171))
                            _ = v175
                        case -1:
                            var v176 float64 = -2.25
                            _ = v176
                        }
                        _ = v173
                        _ = v174
                    }
                    _ = v171
                    _ = v172
                } else {
                    fmt.Println(v2, v3)
                    v5 -= (0.0 * (float32((1.5 - 1e10)) - 0.1))
                    v4 = (v4 % ((255) | 1))
                }
                var v177 uint32 = 2147483647
                _ = v177
            }
            switch -105 {
            case 1:
                for fi178 := 0; fi178 < 4; fi178++ {
                    v179 := GenS20{F0: false, F1: 829, F2: 892}
                    var v180 int = (fi178 | (fi178 - fi178))
                    _ = fi178
                    _ = v179
                    _ = v180
                }
            case -942:
                fmt.Println(v5, v3, v4)
            case -9223372036854775808:
                v3 = 19
            default:
                v181 := GenS13{F0: false}
                _ = v181
            }
            _ = v163
        }
        var v182 uint64 = v54
        _ = v182
    }
    for fi183 := 0; fi183 < 9; fi183++ {
        fmt.Println(v4, v54)
        switch 32767 {
        case 722:
            v3 = (v3 / ((v3) | 1))
        case 813:
            if (false || true) {
                v184 := []int{(fi183 * fi183), fi183, -192, (fi183 % ((1) | 1))}
                for fi185 := 0; fi185 < 3; fi185++ {
                    var v186 int8 = 1
                    v187 := GenS12{F0: 525, F1: (v54 % ((v54) | 1)), F2: ("hello" + "a")}
                    _ = fi185
                    _ = v186
                    _ = v187
                }
                _ = v184
            } else {
                if false {
                    v4 = (^(v4))
                }
            }
        case 127:
            var v188 uint8 = 51
            _ = v188
        }
        if false {
            for fi189 := 0; fi189 < 7; fi189++ {
                v5 += v5
                v2 = int16(v4)
                switch (fi183 - 2147483647) {
                case 32768:
                    v190 := GenS24{F0: 611}
                    _ = v190
                }
                _ = fi189
            }
        } else {
            v4 = v4
            for fi191 := 0; fi191 < 16; fi191++ {
                v192 := GenS25{F0: v3}
                var v193 float64 = (0.0 * 0.1)
                v194 := []int{fi191, fi191}
                _ = fi191
                _ = v192
                _ = v193
                _ = v194
            }
        }
        _ = fi183
    }
    v195 := GenS24{F0: 1}
    v54 = genFn0(84, 535, 2)
    v1[42] = -9223372036854775808
    var v196 uint32 = uint32(v4)
    var v197 float64 = -2.25
    v198 := []int{}
    var v199 uint = 255
    v200 := GenS26{F0: 0}
    v197 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v196:", v196)
    fmt.Println("v197:", v197)
    fmt.Println("slicelen v198:", len(v198), v198)
    fmt.Println("v199:", v199)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v195
    _ = v196
    _ = v197
    _ = v198
    _ = v199
    _ = v2
    _ = v200
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
}
