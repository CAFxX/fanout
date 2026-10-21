package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 uint; F1 int8; F2 string }
type GenS11 struct { F0 string; F1 uint32; F2 bool }
type GenS12 struct { F0 int32; F1 int; F2 uint16 }
func genFn0(p0 int64) uint8 {
    var v6 int64 = p0
    if false {
        v7 := []int{-48, -1}
        var v8 string = "foo bar"
        _ = v7
        _ = v8
    } else {
        v9 := []int{286, 32768, -519, 32768}
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 string = ("hello" + "a")
            if ((-(p0)) > v6) {
                var v12 uint8 = 42
                v12 = (v12 << ((255) & 63))
                _ = v12
            }
            v13 := GenS1{F0: 247}
            _ = fi10
            _ = v11
            _ = v13
        }
        v14 := GenS1{F0: 0}
        _ = v14
        _ = v9
    }
    _ = p0
    _ = v6
    return 127
}
func genFn1(p0 uint) uint {
    var v15 int16 = 0
    if (p0 <= 415) {
        for fi16 := 0; fi16 < 14; fi16++ {
            v15 -= v15
            fmt.Println(p0)
            if false {
                var v17 uint32 = 420
                var v18 float32 = float32(1.5)
                v19 := []int{}
                _ = v17
                _ = v18
                _ = v19
            } else {
                if true {
                    switch (fi16 / (((fi16 * fi16)) | 1)) {
                    case 128:
                        fmt.Println(p0, v15)
                    case -2147483648:
                        p0 = 255
                    case 32767:
                        if true {
                            var v20 int32 = 2147483647
                            for fi21 := 0; fi21 < 16; fi21++ {
                                v22 := GenS2{F0: (-(v20)), F1: fi21}
                                p0 = (p0 >> ((1) & 63))
                                _ = fi21
                                _ = v22
                            }
                            v20 -= v20
                            _ = v20
                        } else {
                            v23 := []int{}
                            v15 = (v15 + (v15 * v15))
                            _ = v23
                        }
                    default:
                        v24 := GenS3{F0: 107}
                        _ = v24
                    }
                }
            }
            _ = fi16
        }
        if ((!false) && ((0.1 * 1.5) == ((1e10 * 1.5) / (1e10 / (1.5 / 1.0))))) {
            for fi25 := 0; fi25 < 20; fi25++ {
                fmt.Println(fi25, v15, v15)
                _ = fi25
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 float32 = 3.14159
                var v28 string = ""
                _ = fi26
                _ = v27
                _ = v28
            }
        } else {
            var v29 uint64 = 1
            v29 -= 769
            _ = v29
        }
        if (true && (!true)) {
            var v30 int8 = -1
            _ = v30
        }
    } else {
        switch "foo bar" {
        case "a":
            for fi31 := 0; fi31 < 19; fi31++ {
                v32 := GenS3{F0: 100}
                fmt.Println(v15, p0, v15)
                _ = fi31
                _ = v32
            }
        case "foo bar":
            p0 -= (p0 & (p0 >> (((p0 / (((^(p0))) | 1))) & 63)))
        case "":
            fmt.Println(v15, v15)
        }
        var v33 int64 = -9223372036854775808
        switch -15 {
        case 1:
            switch -941 {
            case -9223372036854775808:
                v34 := []int{549, -9223372036854775808}
                _ = v34
            case 9223372036854775807:
                if (94 != -1) {
                    v35 := GenS4{F0: v33}
                    v36 := []int{1, -546}
                    var v37 int = 42
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    if ("" > ("" + "a")) {
                        v38 := GenS2{F0: 2147483647, F1: 71}
                        switch "" {
                        case "a":
                            v15 = (v15 * -32768)
                        case "":
                            if true {
                                switch 1 {
                                case -9223372036854775808:
                                    switch 0 {
                                    case -310:
                                        v33 = (v33 | v33)
                                    case 9223372036854775807:
                                        v33 = 42
                                    default:
                                        for fi39 := 0; fi39 < 13; fi39++ {
                                            v40 := GenS3{F0: 79}
                                            if false {
                                                fmt.Println(p0)
                                                for fi41 := 0; fi41 < 15; fi41++ {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v42 int64 = (v33 << ((65535) & 63))
                                                        _ = v42
                                                    case "":
                                                        fmt.Println(v15, fi41)
                                                    case "foo bar":
                                                        v43 := GenS3{F0: 42}
                                                        _ = v43
                                                    default:
                                                        p0 = (p0 << (((p0 / (((p0 * p0)) | 1))) & 63))
                                                    }
                                                    v44 := GenS5{F0: p0, F1: 128, F2: v33}
                                                    _ = fi41
                                                    _ = v44
                                                }
                                            }
                                            for fi45 := 0; fi45 < 1; fi45++ {
                                                var v46 uint8 = 3
                                                var v47 string = "\n\t"
                                                _ = fi45
                                                _ = v46
                                                _ = v47
                                            }
                                            _ = fi39
                                            _ = v40
                                        }
                                    }
                                case 1:
                                    for fi48 := 0; fi48 < 14; fi48++ {
                                        if false {
                                            v49 := []int{fi48}
                                            if (1.5 != 3.14159) {
                                                _ = len(v49)
                                                for fi50 := 0; fi50 < 17; fi50++ {
                                                    switch (^(fi48)) {
                                                    case 32768:
                                                        v51 := GenS4{F0: (v33 >> ((18446744073709551615) & 63))}
                                                        _ = v51
                                                    case 127:
                                                        var v52 uint = p0
                                                        _ = v52
                                                    case 4294967295:
                                                        v15 -= (v15 << (((^(p0))) & 63))
                                                    }
                                                    if (0 == 24) {
                                                        var v53 string = "a"
                                                        v15 = v15
                                                        if true {
                                                            if (2147483647 < 1) {
                                                                if ((float64(0.1) <= 0.1) && true) {
                                                                    p0 = p0
                                                                    switch fi48 {
                                                                    case 42:
                                                                        v33 += v33
                                                                    case 9223372036854775807:
                                                                        var v54 uint64 = 2
                                                                        _ = v54
                                                                    case -2147483648:
                                                                        v55 := GenS1{F0: 127}
                                                                        _ = v55
                                                                    default:
                                                                        p0 -= p0
                                                                    }
                                                                    if len(v49) >= 2 {
                                                                        v49 = v49[1:len(v49)]
                                                                    }
                                                                } else {
                                                                    v33 = v33
                                                                }
                                                            } else {
                                                                v53 = "a"
                                                                var v56 string = v53
                                                                if (3.14159 == 3.14159) {
                                                                    var v57 int16 = (^(v15))
                                                                    _ = v57
                                                                } else {
                                                                    if ((!(v15 == v15)) || true) {
                                                                        _ = len(v49)
                                                                    }
                                                                    v58 := GenS6{F0: 0, F1: (v33 << (((p0 % ((p0) | 1))) & 63))}
                                                                    _ = v58
                                                                }
                                                                _ = v56
                                                            }
                                                            v59 := GenS5{F0: p0, F1: 2, F2: 1}
                                                            _ = v59
                                                        }
                                                        _ = v53
                                                    }
                                                    v60 := GenS4{F0: v33}
                                                    fmt.Println(v15, p0, v15)
                                                    v61 := GenS7{F0: v33}
                                                    _ = fi50
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                            fmt.Println(v15, v33)
                                            _ = v49
                                        } else {
                                            p0 += p0
                                            switch "foo bar" {
                                            case "":
                                                v15 = v15
                                            default:
                                                v62 := []int{(fi48 + fi48), 128, 32768}
                                                _ = v62
                                            }
                                            if (0.1 > 1e10) {
                                                var v63 uint8 = 128
                                                _ = v63
                                            } else {
                                                v64 := []int{-623, int(p0)}
                                                _ = v64
                                            }
                                        }
                                        v65 := GenS7{F0: (v33 | (-(v33)))}
                                        var v66 uint = (p0 / (((p0 & p0)) | 1))
                                        _ = fi48
                                        _ = v65
                                        _ = v66
                                    }
                                case 32768:
                                    var v67 uint64 = 197
                                    _ = v67
                                default:
                                    var v68 string = "a"
                                    _ = v68
                                }
                                p0 = (^(p0))
                                var v69 uint8 = genFn0((v33 & v33))
                                _ = v69
                            }
                        case "\n\t":
                            var v70 float64 = float64(1.5)
                            _ = v70
                        }
                        var v71 int16 = (v15 / (((v15 * v15)) | 1))
                        _ = v38
                        _ = v71
                    } else {
                        v72 := GenS6{F0: 1, F1: (v33 * (v33 | (v33 - v33)))}
                        v73 := GenS8{F0: 18446744073709551615}
                        if (v33 < 9223372036854775807) {
                            fmt.Println(v33)
                            v74 := []int{int(v15), -30, int(p0), 1}
                            var v75 uint64 = 42
                            _ = v74
                            _ = v75
                        } else {
                            switch "a" {
                            case "\n\t":
                                v76 := []int{128, 2147483647, -838, -370}
                                _ = v76
                            case "hello":
                                var v77 float64 = 1e10
                                _ = v77
                            case "foo bar":
                                if (!false) {
                                    if true {
                                        for fi78 := 0; fi78 < 17; fi78++ {
                                            for fi79 := 0; fi79 < 9; fi79++ {
                                                v80 := []int{(fi78 >> ((259) & 63)), fi79}
                                                var v81 uint8 = 128
                                                _ = fi79
                                                _ = v80
                                                _ = v81
                                            }
                                            _ = fi78
                                        }
                                        v82 := GenS6{F0: 42, F1: (v33 | -9223372036854775808)}
                                        var v83 bool = true
                                        _ = v82
                                        _ = v83
                                    }
                                    switch -703 {
                                    case 1:
                                        var v84 uint8 = 1
                                        _ = v84
                                    case 0:
                                        v33 -= v33
                                    default:
                                        v33 += v33
                                    }
                                    for fi85 := 0; fi85 < 19; fi85++ {
                                        var v86 int64 = v33
                                        _ = fi85
                                        _ = v86
                                    }
                                }
                            }
                            v87 := []int{-136, 63, 32767}
                            _ = v87
                        }
                        _ = v72
                        _ = v73
                    }
                }
            case 32767:
                if (false && false) {
                    fmt.Println(v33)
                    p0 = (^(p0))
                }
            }
        case 42:
            v88 := []int{32768, 32767, 2147483647}
            _ = v88
        default:
            for fi89 := 0; fi89 < 15; fi89++ {
                v15 = 1000
                _ = fi89
            }
        }
        _ = v33
    }
    switch 955 {
    case 9223372036854775807:
        if false {
            p0 = (p0 ^ 1)
        }
    case 98:
        fmt.Println(p0)
    case 32767:
        v90 := []int{945, 0, 42}
        _ = v90
    }
    v91 := GenS6{F0: 42, F1: 9223372036854775807}
    _ = p0
    _ = v15
    _ = v91
    return (p0 * p0)
}

func main() {
    var v1 uint16 = 314
    v2 := []int{}
    var v3 uint32 = 1
    var v4 float64 = 0.0
    var v5 bool = false
    if false {
        v92 := GenS9{F0: v3, F1: v1}
        var v93 uint = 18446744073709551615
        if (false && true) {
            if (!((!v5) || v5)) {
                v93 = genFn1((v93 | v93))
                v94 := GenS10{F0: 42, F1: 55, F2: "a"}
                _ = v94
            } else {
                v4 = float64(float32(((1e10 / v4) + (v4 - 1e10))))
            }
        }
        _ = v92
        _ = v93
    } else {
        v95 := GenS5{F0: 2, F1: 30, F2: 1}
        _ = v95
    }
    var v96 int32 = -1
    if ((-32768 == 32767) && v5) {
        v3 = v3
    }
    switch "a" {
    case "foo bar":
        _ = len(v2)
    case "\n\t":
        fmt.Println(v5, v5)
    case "hello":
        v2 = append(v2, -1)
    default:
        if (824 == 1) {
            if v5 {
                for fi97 := 0; fi97 < 15; fi97++ {
                    v96 = v96
                    for fi98 := 0; fi98 < 3; fi98++ {
                        switch "" {
                        case "foo bar":
                            if (!v5) {
                                if v5 {
                                    v5 = false
                                    v99 := GenS2{F0: v96, F1: 1}
                                    _ = v99
                                }
                            } else {
                                v100 := GenS4{F0: 0}
                                _ = v100
                            }
                        case "a":
                            var v101 float64 = v4
                            _ = v101
                        case "hello":
                            var v102 int64 = 9223372036854775807
                            if (100 > 0) {
                                var v103 uint32 = uint32(v1)
                                v104 := GenS9{F0: (v103 << ((2) & 63)), F1: 325}
                                _ = v103
                                _ = v104
                            } else {
                                switch "hello" {
                                case "\n\t":
                                    if (42 >= 915) {
                                        v105 := GenS6{F0: v1, F1: (v102 * (v102 & 0))}
                                        v96 = 0
                                        _ = v105
                                    } else {
                                        var v106 int = 32768
                                        _ = v106
                                    }
                                default:
                                    fmt.Println(v96, v5)
                                }
                            }
                            var v107 uint16 = 245
                            _ = v102
                            _ = v107
                        default:
                            var v108 int16 = 0
                            _ = v108
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(v4, v1)
                        case "hello":
                            var v109 uint16 = 255
                            _ = v109
                        }
                        _ = fi98
                    }
                    var v110 uint64 = 42
                    _ = fi97
                    _ = v110
                }
                var v111 int = 4294967295
                v5 = (211 == 507)
                _ = v111
            } else {
                v4 = (((v4 / 3.14159) + v4) + float64(((((1e10 + 1e10) * 3.14159) / 0.1) * (3.14159 + -2.25))))
            }
        }
    }
    v4 = 0.0
    var v112 uint = 65535
    var v113 int = -978
    fmt.Println(v112)
    v113 = v113
    switch v113 {
    case 88:
        var v114 int = (v113 / ((v113) | 1))
        _ = v114
    case -2147483648:
        fmt.Println(v5, v5, v113)
    default:
        if true {
            _ = len(v2)
        } else {
            fmt.Println(v5, v96, v96)
            v115 := GenS11{F0: "", F1: v3, F2: ((^(v112)) == (v112 >> ((v112) & 63)))}
            _ = v115
        }
    }
    switch (v113 >> (((v112 >> (((^(v112))) & 63))) & 63)) {
    case 2147483647:
        v96 -= v96
    case 4294967295:
        v96 = v96
    }
    for fi116 := 0; fi116 < 10; fi116++ {
        for fi117 := 0; fi117 < 11; fi117++ {
            v118 := GenS12{F0: (v96 & v96), F1: int(v96), F2: v1}
            fmt.Println(v1, v4, v3)
            v112 += v112
            _ = fi117
            _ = v118
        }
        var v119 uint64 = 2
        fmt.Println(v112)
        _ = fi116
        _ = v119
    }
    v96 = v96
    switch ("" + "") {
    case "":
        v113 += v113
    default:
        var v120 uint8 = 229
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v112
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v96
}
