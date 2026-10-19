package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint16; F1 int16; F2 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint; F2 int8 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int64; F1 uint16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 uint16; F1 int8; F2 int32 }
type GenS15 struct { F0 uint16; F1 string }

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    var v3 bool = v1
    var v4 int32 = 42
    var v5 int16 = 0
    var v6 int8 = -11
    v3 = v1
    v4 += v4
    switch "\n\t" {
    case "\n\t":
        v7 := []int{-1, -9223372036854775808, 1}
        _ = v7
    case "hello":
        v8 := []int{128, 128, 584, 128}
        _ = v8
    default:
        var v9 int32 = 190
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        if v1 {
            switch (-(fi10)) {
            case 1:
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 uint64 = 42
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v4, fi10)
                    case "\n\t":
                        var v13 int = (fi11 << ((529) & 63))
                        _ = v13
                    }
                    var v14 uint = 1
                    _ = fi11
                    _ = v12
                    _ = v14
                }
            case 2147483647:
                switch "hello" {
                case "\n\t":
                    v1 = false
                case "a":
                    if v3 {
                        v1 = false
                    } else {
                        v15 := GenS1{F0: "hello", F1: (fi10 * fi10), F2: 255}
                        _ = v15
                    }
                default:
                    var v16 int = (^(fi10))
                    _ = v16
                }
            default:
                v17 := GenS1{F0: "hello", F1: fi10, F2: 2}
                _ = v17
            }
            for fi18 := 0; fi18 < 10; fi18++ {
                v1 = v1
                _ = fi18
            }
            v19 := []int{-2147483648, fi10, (fi10 & 0), fi10}
            _ = v19
        }
        var v20 uint = 42
        _ = fi10
        _ = v20
    }
    v6 += 1
    if (((1.5 * 0.1) != 1.5) && (1e10 >= 3.14159)) {
        if v1 {
            switch "a" {
            case "":
                if (!v1) {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            if (!(!v3)) {
                                var v23 float64 = float64(-2.25)
                                v3 = false
                                fmt.Println(v23, v5, fi22)
                                if (false && ("" > "\n\t")) {
                                    if (65535 < 2147483648) {
                                        v1 = v1
                                        if v3 {
                                            if v1 {
                                                fmt.Println(v5, v4, v4)
                                            }
                                            v24 := GenS2{F0: 32767, F1: v3}
                                            var v25 uint64 = 18446744073709551615
                                            _ = len(v2)
                                            _ = v24
                                            _ = v25
                                        } else {
                                            v26 := GenS3{F0: 706, F1: 385, F2: -32768}
                                            _ = v26
                                        }
                                        var v27 uint8 = 255
                                        _ = v27
                                    }
                                }
                                switch (((("" + "foo bar") + ("\n\t" + "foo bar")) + "") + "") {
                                case "a":
                                    var v28 uint16 = 32767
                                    _ = len(v2)
                                    v6 -= v6
                                    _ = v28
                                case "hello":
                                    v29 := GenS3{F0: 0, F1: (v5 << ((1) & 63)), F2: (^(v5))}
                                    _ = v29
                                default:
                                    v30 := GenS1{F0: "", F1: fi21, F2: 361}
                                    _ = v30
                                }
                                _ = v23
                            } else {
                                _ = len(v2)
                                var v31 bool = ((((-1 < (v6 >> ((2) & 63))) && v3) || (540 >= fi21)) && (128 <= 216))
                                v32 := []int{}
                                _ = v31
                                _ = v32
                            }
                            v3 = v3
                            _ = fi22
                        }
                        if (false || (!(42 == 2))) {
                            if v1 {
                                for fi33 := 0; fi33 < 12; fi33++ {
                                    v34 := GenS1{F0: "\n\t", F1: int(v5), F2: 871}
                                    v5 = int16(v6)
                                    _ = fi33
                                    _ = v34
                                }
                                v35 := GenS4{F0: v6}
                                _ = v35
                            }
                            _ = len(v2)
                            v36 := GenS4{F0: (v6 / (((^(v6))) | 1))}
                            _ = v36
                        }
                        fmt.Println(v1)
                        _ = fi21
                    }
                    v3 = (v1 && (!(414 == 996)))
                }
            case "foo bar":
                v4 = v4
            }
        } else {
            delete(v2, 127)
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 uint64 = 42
                _ = fi37
                _ = v38
            }
        }
        var v39 float64 = -2.25
        _ = v39
    } else {
        v4 = 370
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 int64 = -1
            v41 = (v41 | 983)
            var v42 int8 = (v6 | 65)
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    v3 = false
    fmt.Println(v1)
    for fi43 := 0; fi43 < 8; fi43++ {
        switch "\n\t" {
        case "":
            v44 := []int{fi43, (fi43 - int(v6)), 0}
            _ = v44
        case "foo bar":
            v45 := GenS5{F0: fi43, F1: 42, F2: -128}
            _ = v45
        case "a":
            for fi46 := 0; fi46 < 20; fi46++ {
                switch (fi43 ^ (-(fi46))) {
                case -447:
                    var v47 uint32 = 1
                    _ = v47
                case 4294967295:
                    if ((^(fi46)) > -2147483648) {
                        v4 -= v4
                        var v48 float32 = 0.0
                        _ = v48
                    } else {
                        switch "a" {
                        case "foo bar":
                            v3 = (!(!(1 == 2)))
                        case "\n\t":
                            switch 0 {
                            case 9223372036854775807:
                                switch 791 {
                                case 32767:
                                    if true {
                                        var v49 uint32 = 42
                                        var v50 bool = false
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v51 := GenS2{F0: 32767, F1: v1}
                                        _ = v51
                                    }
                                }
                            case 1:
                                var v52 float32 = (1e10 * -2.25)
                                _ = v52
                            default:
                                var v53 int = fi46
                                _ = v53
                            }
                        case "hello":
                            var v54 float32 = (0.0 / 1.5)
                            _ = v54
                        default:
                            switch fi46 {
                            case 1:
                                var v55 int8 = (v6 / (((v6 / ((int8(fi43)) | 1))) | 1))
                                var v56 uint64 = 42
                                var v57 bool = v1
                                _ = v55
                                _ = v56
                                _ = v57
                            case 42:
                                v58 := []int{(^(fi43)), (fi43 >> ((4294967295) & 63))}
                                if (v4 <= (v4 - (-(v4)))) {
                                    fmt.Println(v6, fi46)
                                    fmt.Println(fi46, v4, v4)
                                } else {
                                    if ((^(v5)) >= v5) {
                                        var v59 float64 = (0.1 - (-2.25 / 1.0))
                                        delete(v2, 127)
                                        _ = v59
                                    } else {
                                        var v60 float64 = 0.1
                                        if (4294967295 > 32768) {
                                            v61 := GenS4{F0: v6}
                                            v2[fi43] = ((fi43 == (fi46 >> ((2) & 63))) || v1)
                                            switch fi43 {
                                            case -2147483648:
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                            case -860:
                                                fmt.Println(v60)
                                            case 179:
                                                switch "hello" {
                                                case "":
                                                    if v3 {
                                                        var v62 int8 = v6
                                                        v62 -= 127
                                                        _ = v62
                                                    } else {
                                                        var v63 uint8 = 99
                                                        var v64 uint16 = 768
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                case "hello":
                                                    var v65 uint = 18446744073709551615
                                                    _ = v65
                                                case "foo bar":
                                                    v4 += (v4 ^ v4)
                                                    var v66 int = 930
                                                    _ = v66
                                                }
                                            default:
                                                switch fi46 {
                                                case 729:
                                                    switch (("" + "a") + "foo bar") {
                                                    case "hello":
                                                        var v67 string = ("hello" + (((("" + "") + "\n\t") + "hello") + ("\n\t" + ("\n\t" + "foo bar"))))
                                                        _ = v67
                                                    case "":
                                                        if len(v58) > 0 {
                                                            v58[0] = (fi46 + fi43)
                                                        }
                                                    default:
                                                        v4 -= 1
                                                        v68 := GenS6{F0: 1, F1: 948}
                                                        if v1 {
                                                            if v3 {
                                                                var v69 int64 = 803
                                                                _ = v69
                                                            }
                                                            v70 := GenS7{F0: 2, F1: 0}
                                                            _ = v70
                                                        } else {
                                                            v4 -= (v4 % (((v4 >> ((152) & 63))) | 1))
                                                            if v1 {
                                                                var v71 int16 = (v5 + -804)
                                                                if (("" >= ("a" + "a")) || (!false)) {
                                                                    if v3 {
                                                                        v6 -= int8(v5)
                                                                        v2[(^(fi46))] = false
                                                                        fmt.Println(v4, v6, v71)
                                                                        var v72 float64 = ((-2.25 / v60) - v60)
                                                                        v73 := GenS7{F0: 4294967295, F1: 479}
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    v74 := GenS8{F0: 663, F1: 65535}
                                                                    if v3 {
                                                                        if len(v58) > 0 {
                                                                            v58[0] = fi46
                                                                        }
                                                                        v5 = (v5 << ((0) & 63))
                                                                        if (196 < 18446744073709551615) {
                                                                            v75 := GenS2{F0: 32767, F1: (1 > 1)}
                                                                            _ = v75
                                                                        }
                                                                    } else {
                                                                        var v76 int = -9223372036854775808
                                                                        _ = len(v2)
                                                                        _ = v76
                                                                    }
                                                                    _ = v74
                                                                }
                                                                _ = v71
                                                            }
                                                            fmt.Println(v5, v3, fi46)
                                                        }
                                                        _ = v68
                                                    }
                                                default:
                                                    var v77 float64 = -2.25
                                                    _ = v77
                                                }
                                            }
                                            if (!true) {
                                                v2[fi46] = (1.5 != 3.14159)
                                            } else {
                                                var v78 float32 = (1e10 - 0.1)
                                                switch "hello" {
                                                case "a":
                                                    fmt.Println(v6)
                                                case "\n\t":
                                                    if v3 {
                                                        switch "a" {
                                                        case "":
                                                            if len(v58) > 0 {
                                                                v58[0] = (fi46 - (fi43 * 4294967295))
                                                            }
                                                        case "foo bar":
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                        default:
                                                            v4 = int32(v6)
                                                            _ = len(v58)
                                                            if (v1 && ((-(fi46)) <= (fi43 / (((fi46 % ((fi46) | 1))) | 1)))) {
                                                                fmt.Println(v60, fi46)
                                                                if (127 > 128) {
                                                                    if v3 {
                                                                        if true {
                                                                            switch fi43 {
                                                                            case 9223372036854775807:
                                                                                v4 = (-(v4))
                                                                            }
                                                                            fmt.Println(v60)
                                                                        }
                                                                        v79 := GenS1{F0: "hello", F1: fi43, F2: 825}
                                                                        _ = v79
                                                                    } else {
                                                                        v80 := GenS9{F0: 42}
                                                                        _ = v80
                                                                    }
                                                                    v3 = ((v78 * 0.1) == 3.14159)
                                                                    if (v3 && (1 >= 397)) {
                                                                        v81 := GenS2{F0: 786, F1: (((fi43 + 127) <= (fi43 * -985)) && (620 < 2))}
                                                                        _ = v81
                                                                    } else {
                                                                        fmt.Println(v5, v78, fi43)
                                                                    }
                                                                } else {
                                                                    v4 = v4
                                                                }
                                                                v3 = v1
                                                            } else {
                                                                var v82 string = "foo bar"
                                                                var v83 int64 = 571
                                                                var v84 int16 = 1000
                                                                v85 := GenS10{F0: v83, F1: 556}
                                                                if (42 != 964) {
                                                                    switch 9223372036854775807 {
                                                                    case 323:
                                                                        if v1 {
                                                                            v84 = -32768
                                                                        }
                                                                        if (((1.5 * 1e10) / v78) <= float32((3.14159 / v60))) {
                                                                            var v86 float32 = 3.14159
                                                                            if (true || ((0.1 + 1e10) <= v78)) {
                                                                                if len(v58) > 0 {
                                                                                    v58[0] = (fi46 >> ((4294967295) & 63))
                                                                                }
                                                                            }
                                                                            if (!(1 < 1)) {
                                                                                var v87 uint = 18446744073709551615
                                                                                _ = v87
                                                                            }
                                                                            _ = v86
                                                                        } else {
                                                                            v88 := GenS11{F0: true}
                                                                            var v89 uint = 255
                                                                            v84 += v5
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                    case 32767:
                                                                        v78 = (float32(v60) * v78)
                                                                    case 1:
                                                                        var v90 int16 = v84
                                                                        _ = v90
                                                                    }
                                                                } else {
                                                                    v91 := GenS12{F0: 79}
                                                                    v1 = v1
                                                                    switch v82 {
                                                                    case "foo bar":
                                                                        v92 := GenS10{F0: v83, F1: 32767}
                                                                        _ = v92
                                                                    case "a":
                                                                        var v93 float64 = (((v60 + (v60 - v60)) / float64((-2.25 + v78))) - float64((v78 * (v78 * v78))))
                                                                        _ = v93
                                                                    }
                                                                    _ = v91
                                                                }
                                                                _ = v82
                                                                _ = v83
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                        }
                                                    }
                                                case "foo bar":
                                                    v94 := GenS5{F0: (-(fi46)), F1: 840, F2: (v6 * 42)}
                                                    fmt.Println(v4, v5)
                                                    _ = v94
                                                }
                                                v60 -= (-2.25 * float64((float32(0.1) / -2.25)))
                                                _ = v78
                                            }
                                            var v95 uint16 = 32767
                                            if v1 {
                                                var v96 int32 = v4
                                                fmt.Println(v60)
                                                v60 += v60
                                                _ = v96
                                            } else {
                                                var v97 int16 = (^(v5))
                                                if false {
                                                    v98 := GenS11{F0: true}
                                                    v58 = append(v58, (fi43 % ((int(v95)) | 1)))
                                                    var v99 uint8 = 1
                                                    v99 = 128
                                                    var v100 int8 = v6
                                                    _ = v100
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    v3 = false
                                                    fmt.Println(v97, v97)
                                                    v101 := GenS9{F0: 1}
                                                    switch "\n\t" {
                                                    case "a":
                                                        if (65535 >= 193) {
                                                            fmt.Println(v3, fi46, v97)
                                                            fmt.Println(fi46)
                                                        } else {
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                            fmt.Println(fi43, v1, fi43)
                                                            if ((v60 / (v60 - 3.14159)) <= float64((0.0 + (-2.25 + (1.5 * 1.5))))) {
                                                                v60 += 0.1
                                                                v97 -= v97
                                                            }
                                                        }
                                                    default:
                                                        if (((((((v60 - v60) - v60) / float64(3.14159)) + v60) > (0.0 / v60)) && (v1 && (!true))) || (255 > 0)) {
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                            var v102 int32 = v4
                                                            if v3 {
                                                                _ = len(v2)
                                                                v103 := GenS12{F0: 2}
                                                                _ = len(v2)
                                                                _ = v103
                                                            }
                                                            if v1 {
                                                                if len(v58) > 0 {
                                                                    v58[0] = fi43
                                                                }
                                                            } else {
                                                                v60 = (v60 - v60)
                                                                v4 = v4
                                                            }
                                                            v97 = (v97 - (v5 & (v5 % ((3) | 1))))
                                                            if (v3 && ((-(fi46)) == (fi46 >> ((68) & 63)))) {
                                                                fmt.Println(v3, v6, v102)
                                                                var v104 uint = 65535
                                                                fmt.Println(v60, v6)
                                                                v105 := v2[(fi43 << ((376) & 63))]
                                                                v97 = (^(v5))
                                                                _ = v104
                                                                _ = v105
                                                            }
                                                            _ = v102
                                                        } else {
                                                            v1 = v1
                                                            switch fi46 {
                                                            case 2147483647:
                                                                switch (fi46 & fi46) {
                                                                case 0:
                                                                    fmt.Println(v1)
                                                                case -2147483648:
                                                                    var v106 float64 = 1e10
                                                                    _ = v106
                                                                }
                                                            case 128:
                                                                var v107 uint32 = 0
                                                                _ = v107
                                                            default:
                                                                v95 = (v95 >> ((278) & 63))
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v97)
                                                    _ = v101
                                                }
                                                v108 := GenS4{F0: int8(fi43)}
                                                _ = v108
                                                _ = v97
                                            }
                                            v109 := GenS13{F0: 2147483647, F1: v1}
                                            v2[fi46] = (!true)
                                            _ = v109
                                            _ = v61
                                            _ = v95
                                        }
                                        _ = v60
                                    }
                                    var v110 float64 = 1.5
                                    v111 := GenS14{F0: 2, F1: (v6 % (((-(v6))) | 1)), F2: (v4 % (((v4 + -2147483648)) | 1))}
                                    _ = v110
                                    _ = v111
                                }
                                fmt.Println(fi43, v5, v4)
                                _ = v58
                            }
                        }
                        v112 := GenS15{F0: 2, F1: (("foo bar" + ("\n\t" + "")) + "foo bar")}
                        switch 127 {
                        case 128:
                            fmt.Println(v3, v3)
                        case 0:
                            var v113 uint = 1
                            _ = v113
                        case -131:
                            var v114 int64 = 871
                            _ = v114
                        }
                        _ = v112
                    }
                default:
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v6)
                    default:
                        v4 -= (v4 << ((65535) & 63))
                        if false {
                            fmt.Println(v6, fi43, v1)
                        }
                    }
                }
                v115 := v2[32768]
                v115 = (32767 == v5)
                _ = fi46
                _ = v115
            }
        }
        delete(v2, fi43)
        var v116 int32 = v4
        _ = fi43
        _ = v116
    }
    v6 -= -57
    var v117 uint32 = 126
    _ = len(v2)
    var v118 uint16 = uint16(v117)
    if v1 {
        switch -9223372036854775808 {
        case -750:
            switch 4294967295 {
            case 9223372036854775807:
                v1 = (1 >= uint(v117))
            case -1:
                v5 -= v5
            }
        }
    }
    fmt.Println(v1)
    v119 := v2[32768]
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
