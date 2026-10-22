package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int64 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS11 struct { F0 uint16; F1 int16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS14 struct { F0 int64 }

func main() {
    var v1 uint32 = 83
    var v2 uint64 = 214
    var v3 int8 = 1
    var v4 float64 = -2.25
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: (v3 | v3), F1: (v3 | v3), F2: 343}
    for fi8 := 0; fi8 < 8; fi8++ {
        if (2 != 32767) {
            v9 := GenS2{F0: 1, F1: true}
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := []int{fi8, (fi10 % ((-2147483648) | 1)), (fi10 | fi8), 9223372036854775807}
                var v12 uint8 = 255
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint8 = uint8(v2)
            _ = v13
            _ = v9
        } else {
            var v14 int32 = 42
            _ = len(v6)
            _ = v14
        }
        switch "" {
        case "\n\t":
            v4 -= 3.14159
        default:
            if false {
                fmt.Println(v2, v4, v2)
            }
        }
        v15 := GenS2{F0: -1, F1: true}
        _ = fi8
        _ = v15
    }
    if (3.14159 != 0.0) {
        if false {
            v16 := GenS3{F0: 32767, F1: 538, F2: 32767}
            _ = v16
        }
    }
    fmt.Println(v4)
    v4 = v4
    if false {
        if true {
            if true {
                v2 -= 0
            }
        } else {
            v17 := []int{985}
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 5; fi19++ {
                    fmt.Println(fi18)
                    _ = fi19
                }
                v1 = v1
                if (v3 != v3) {
                    if true {
                        v2 += (v2 + (v2 + 2))
                    }
                    v20 := GenS3{F0: 2, F1: -272, F2: fi18}
                    var v21 uint = 65535
                    _ = v20
                    _ = v21
                } else {
                    if true {
                        fmt.Println(v3)
                        v22 := GenS4{F0: true, F1: -128}
                        fmt.Println(v4, v1)
                        _ = v22
                    }
                }
                _ = fi18
            }
            _ = v17
        }
        switch (((("\n\t" + ("\n\t" + ("a" + ("" + "")))) + "hello") + "") + "hello") {
        case "":
            v23 := GenS4{F0: true, F1: (v3 - (v3 & v3))}
            _ = v23
        case "foo bar":
            v24 := GenS2{F0: -128, F1: (!false)}
            _ = v24
        }
        switch 128 {
        case -433:
            if (2147483647 != (^(v1))) {
                if false {
                    if (!false) {
                        var v25 string = ("foo bar" + "\n\t")
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 3; fi26++ {
                        if (-2147483648 == 0) {
                            switch fi26 {
                            case 2147483647:
                                v6[fi26] = v1
                            case -2147483648:
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    var v28 uint16 = 32767
                                    v6[(fi26 & (fi27 << ((2) & 63)))] = (v1 >> ((1) & 63))
                                    _ = fi27
                                    _ = v28
                                }
                            }
                        } else {
                            var v29 string = ""
                            _ = v29
                        }
                        v4 = v4
                        _ = fi26
                    }
                } else {
                    switch 0 {
                    case -811:
                        for fi30 := 0; fi30 < 12; fi30++ {
                            var v31 int16 = -716
                            v32 := GenS5{F0: 9223372036854775807}
                            var v33 float64 = (float64(0.1) + v4)
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    case 9223372036854775807:
                        var v34 uint64 = (v2 / ((v2) | 1))
                        _ = v34
                    case 42:
                        v4 -= float64(((1e10 + (3.14159 + 3.14159)) / 1e10))
                    default:
                        if (false || (true || true)) {
                            var v35 float32 = (1e10 - float32(float64(-2.25)))
                            v36 := GenS6{F0: 103}
                            v37 := GenS6{F0: 1}
                            _ = v35
                            _ = v36
                            _ = v37
                        } else {
                            delete(v6, 2147483647)
                        }
                    }
                }
                var v38 uint32 = 360
                var v39 string = "foo bar"
                _ = v38
                _ = v39
            } else {
                v4 += v4
            }
        default:
            for fi40 := 0; fi40 < 8; fi40++ {
                var v41 int = (fi40 - int(v1))
                if (v1 > v1) {
                    fmt.Println(v4)
                    v2 = (^(v2))
                    if true {
                        v42 := []int{fi40, 32768, fi40, (v41 & int(v1))}
                        _ = len(v42)
                        _ = v42
                    }
                } else {
                    v43 := []int{}
                    for fi44 := 0; fi44 < 19; fi44++ {
                        v45 := GenS6{F0: 497}
                        _ = fi44
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 16; fi46++ {
                        switch "" {
                        case "\n\t":
                            var v47 int8 = v3
                            _ = v47
                        default:
                            var v48 int8 = v3
                            _ = v48
                        }
                        if (!true) {
                            v2 = v2
                        }
                        _ = fi46
                    }
                    _ = v43
                }
                _ = fi40
                _ = v41
            }
        }
    } else {
        switch 128 {
        case 42:
            var v49 string = "a"
            _ = v49
        default:
            if (!true) {
                v50 := GenS3{F0: 42, F1: -9223372036854775808, F2: -1}
                _ = v50
            } else {
                switch ("\n\t" + (("foo bar" + "hello") + ("foo bar" + (("hello" + "a") + "foo bar")))) {
                case "hello":
                    v51 := GenS3{F0: 2, F1: -103, F2: 32767}
                    _ = v51
                default:
                    switch ("hello" + "hello") {
                    case "foo bar":
                        if (4294967295 >= v2) {
                            for fi52 := 0; fi52 < 2; fi52++ {
                                v53 := []int{(fi52 << ((290) & 63)), (fi52 % ((fi52) | 1)), 9223372036854775807, fi52}
                                var v54 int16 = 42
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                            fmt.Println(v2, v1)
                        }
                    default:
                        v1 += (v1 | (v1 % (((v1 ^ 42)) | 1)))
                    }
                }
                fmt.Println(v2, v1)
                var v55 int = 127
                _ = v55
            }
        }
    }
    fmt.Println(v3, v3)
    for fi56 := 0; fi56 < 13; fi56++ {
        if false {
            var v57 uint32 = v1
            delete(v5, -238)
            _ = v57
        } else {
            var v58 int16 = 0
            v1 = v1
            if ((!("a" >= "a")) && true) {
                switch fi56 {
                case 4294967295:
                    v2 = (v2 / ((v2) | 1))
                case -638:
                    var v59 int8 = 100
                    _ = v59
                default:
                    v60 := GenS6{F0: 0}
                    _ = v60
                }
                v2 = (v2 ^ 872)
                if true {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v3, v3)
                    case "a":
                        if (340 < 92) {
                            switch "a" {
                            case "hello":
                                for fi61 := 0; fi61 < 12; fi61++ {
                                    var v62 float64 = (v4 * (float64((1e10 - (1e10 / -2.25))) / v4))
                                    if (false && false) {
                                        v3 = (v3 >> ((0) & 63))
                                    }
                                    var v63 int16 = (v58 ^ (v58 * (v58 & (v58 | (^(v58))))))
                                    _ = fi61
                                    _ = v62
                                    _ = v63
                                }
                            case "\n\t":
                                if false {
                                    v64 := []int{(^(fi56))}
                                    if (v4 == (v4 + v4)) {
                                        v65 := GenS5{F0: (fi56 | fi56)}
                                        _ = v65
                                    }
                                    _ = v64
                                }
                            case "foo bar":
                                v4 = ((v4 - (v4 * 1.5)) / ((v4 + v4) * 0.1))
                            }
                        } else {
                            v2 += 581
                            var v66 int = int(v2)
                            v2 = 4294967295
                            _ = v66
                        }
                    case "foo bar":
                        v67 := []int{}
                        _ = v67
                    default:
                        var v68 uint = 1
                        _ = v68
                    }
                    var v69 int16 = (v58 * v58)
                    _ = v69
                }
            }
            _ = v58
        }
        _ = fi56
    }
    if false {
        for fi70 := 0; fi70 < 17; fi70++ {
            for fi71 := 0; fi71 < 6; fi71++ {
                v72 := GenS7{F0: "a", F1: 185}
                if ("" <= "hello") {
                    var v73 int32 = 42
                    v74 := GenS3{F0: 65535, F1: 0, F2: (fi70 - (fi70 ^ -1))}
                    v75 := GenS2{F0: (v3 ^ (v3 + v3)), F1: (v2 != 18446744073709551615)}
                    _ = v73
                    _ = v74
                    _ = v75
                } else {
                    var v76 int = fi70
                    var v77 int8 = (v3 & 1)
                    _ = v76
                    _ = v77
                }
                if (!false) {
                    if false {
                        switch (fi71 & (fi70 | fi71)) {
                        case 42:
                            v78 := []int{(fi71 - (fi71 + fi71)), (^(fi70))}
                            _ = v78
                        case -2147483648:
                            v3 += 127
                        default:
                            var v79 float64 = float64((0.0 - 3.14159))
                            _ = v79
                        }
                        switch (fi71 * fi70) {
                        case 32767:
                            if (fi71 == 9223372036854775807) {
                                switch (("" + "foo bar") + "foo bar") {
                                case "hello":
                                    fmt.Println(fi70, v2, v2)
                                case "foo bar":
                                    if ((-1 >= -270) && (int16(v2) < 42)) {
                                        v80 := GenS8{F0: (!false)}
                                        v2 -= v2
                                        v5[-697] = fi70
                                        v81 := []int{(fi70 << ((65535) & 63)), -407, (-(fi71))}
                                        fmt.Println(v1, v1, fi71)
                                        var v82 uint8 = 59
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    } else {
                                        var v83 bool = (-674 == 737)
                                        switch ((("hello" + "") + "\n\t") + ("" + "\n\t")) {
                                        case "":
                                            if v83 {
                                                v84 := GenS6{F0: 42}
                                                v3 = (v3 ^ v3)
                                                var v85 uint = 4294967295
                                                _ = v84
                                                _ = v85
                                            } else {
                                                var v86 bool = (v4 != 1.5)
                                                fmt.Println(v3, fi70)
                                                _ = v86
                                            }
                                        default:
                                            v87 := GenS9{F0: (v3 << ((209) & 63)), F1: 36}
                                            _ = v87
                                        }
                                        fmt.Println(v2)
                                        switch fi70 {
                                        case -504:
                                            var v88 uint8 = uint8(v2)
                                            _ = v88
                                        default:
                                            v4 = (v4 - 1e10)
                                        }
                                        _ = v83
                                    }
                                }
                            } else {
                                if true {
                                    v4 = float64(3.14159)
                                    if (fi70 >= 1) {
                                        if ((("\n\t" + "foo bar") >= "foo bar") && false) {
                                            switch "hello" {
                                            case "a":
                                                switch "\n\t" {
                                                case "hello":
                                                    if false {
                                                        v89 := []int{9223372036854775807, -564, -1}
                                                        v90 := GenS10{F0: false, F1: 1, F2: 138}
                                                        v91 := GenS11{F0: 267, F1: 1000}
                                                        _ = v89
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        v3 -= (^(v3))
                                                        switch ("hello" + "foo bar") {
                                                        case "foo bar":
                                                            v92 := GenS12{F0: 914}
                                                            _ = v92
                                                        }
                                                        v93 := []int{}
                                                        _ = v93
                                                    }
                                                case "":
                                                    v94 := []int{(fi70 + (-(fi70)))}
                                                    _ = v94
                                                case "\n\t":
                                                    fmt.Println(v1)
                                                }
                                            }
                                            v95 := GenS10{F0: true, F1: (v1 * uint32(v2)), F2: 42}
                                            v4 = v4
                                            _ = v95
                                        }
                                        var v96 bool = ((v1 / (((v1 + v1)) | 1)) > v1)
                                        v97 := GenS9{F0: v3, F1: 2}
                                        _ = v96
                                        _ = v97
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v2)
                                    case "\n\t":
                                        v1 = (v1 >> ((18446744073709551615) & 63))
                                    case "foo bar":
                                        if false {
                                            fmt.Println(v4)
                                        } else {
                                            var v98 float64 = float64(((3.14159 / 1e10) - float32(1e10)))
                                            if (fi70 > fi71) {
                                                var v99 string = "a"
                                                v100 := []int{}
                                                _ = v100
                                                _ = v99
                                            } else {
                                                fmt.Println(fi70, v2)
                                                v1 = v1
                                                var v101 uint64 = (^(v2))
                                                _ = v101
                                            }
                                            v102 := GenS13{F0: 808, F1: (v1 >> ((42) & 63)), F2: (v2 | 2)}
                                            _ = v102
                                            _ = v98
                                        }
                                    default:
                                        _ = len(v5)
                                    }
                                    v103 := GenS10{F0: (!true), F1: (v1 & (v1 | v1)), F2: 1}
                                    if (1.5 <= 3.14159) {
                                        var v104 uint64 = v2
                                        _ = v104
                                    }
                                    switch -789 {
                                    case 0:
                                        if false {
                                            switch -9223372036854775808 {
                                            case 228:
                                                v2 = (v2 >> ((1) & 63))
                                            case 963:
                                                delete(v6, (fi71 + fi70))
                                            }
                                            v5[(fi71 << ((1) & 63))] = fi70
                                            var v105 uint64 = 18446744073709551615
                                            _ = v105
                                        }
                                    case 4294967295:
                                        v1 = (v1 | (v1 & 4294967295))
                                    default:
                                        var v106 int16 = 32767
                                        v2 = v2
                                        _ = v106
                                    }
                                    _ = v103
                                } else {
                                    var v107 int32 = -1
                                    switch (fi71 % ((4294967295) | 1)) {
                                    case -2147483648:
                                        var v108 float32 = 3.14159
                                        _ = v108
                                    case -9223372036854775808:
                                        v2 = v2
                                    case 42:
                                        var v109 uint = 255
                                        _ = v109
                                    }
                                    _ = v107
                                }
                                var v110 int64 = 42
                                v111 := []int{}
                                fmt.Println(fi71, v3)
                                _ = v110
                                _ = v111
                            }
                        case -1:
                            v112 := GenS13{F0: 42, F1: v1, F2: v2}
                            _ = v112
                        }
                    }
                    v4 = v4
                    var v113 uint16 = 1
                    v114 := []int{-759}
                    _ = v113
                    _ = v114
                }
                _ = fi71
                _ = v72
            }
            fmt.Println(v4)
            _ = fi70
        }
    } else {
        fmt.Println(v3, v1)
        v115 := []int{2147483647, 1, int(v3), -1}
        _ = v115
    }
    v116 := GenS14{F0: int64(v3)}
    v117 := v5[0]
    v118 := GenS11{F0: uint16(v2), F1: 521}
    if true {
        fmt.Println(v1, v1)
        for fi119 := 0; fi119 < 8; fi119++ {
            fmt.Println(v2, v2, v3)
            var v120 float64 = (v4 * v4)
            _ = fi119
            _ = v120
        }
        var v121 int64 = -9223372036854775808
        _ = v121
    } else {
        fmt.Println(v117, v2, v117)
        v2 -= (v2 << ((1) & 63))
    }
    delete(v5, v117)
    v1 += (v1 | v1)
    var v122 uint = 0
    var v123 uint8 = 108
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v116
    _ = v117
    _ = v118
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
