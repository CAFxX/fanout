package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int; F2 string }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 int32; F1 int64; F2 uint64 }
type GenS12 struct { F0 int16; F1 int64; F2 uint }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 int; F1 int8; F2 uint8 }
type GenS17 struct { F0 uint; F1 int64 }
type GenS18 struct { F0 string; F1 uint16 }
type GenS19 struct { F0 int8; F1 int16 }
type GenS20 struct { F0 int16; F1 int16; F2 uint16 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 int8 }
func genFn0(p0 uint, p1 bool) int8 {
    var v4 int8 = -1
    v5 := GenS1{F0: 42, F1: 0}
    switch 1 {
    case 0:
        v6 := GenS1{F0: 559, F1: (p0 << ((255) & 63))}
        _ = v6
    }
    p1 = (-164 <= 787)
    v7 := GenS1{F0: 4294967295, F1: p0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    return -84
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    var v3 uint32 = 105
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            v3 = v3
            fmt.Println(v1)
        } else {
            v2[fi8] = true
            for fi9 := 0; fi9 < 6; fi9++ {
                fmt.Println(fi8)
                switch (fi9 * 128) {
                case 1:
                    if true {
                        var v10 bool = true
                        v11 := []int{822, 4294967295, (fi9 >> ((1) & 63)), fi9}
                        fmt.Println(v3, v3)
                        _ = v10
                        _ = v11
                    } else {
                        if (int32(fi9) <= (^(v1))) {
                            switch fi8 {
                            case 605:
                                fmt.Println(fi9)
                            }
                            switch (fi9 & 128) {
                            case 789:
                                var v12 uint = 2
                                var v13 int32 = -1
                                _ = v12
                                _ = v13
                            case -9223372036854775808:
                                var v14 uint32 = (v3 * (v3 % ((v3) | 1)))
                                _ = v14
                            case 385:
                                var v15 int32 = v1
                                _ = v15
                            }
                            if (!((!(!true)) && ((4294967295 > 255) && (4294967295 == 65535)))) {
                                var v16 int16 = 538
                                v17 := GenS1{F0: 361, F1: 65535}
                                _ = v16
                                _ = v17
                            } else {
                                if (true || false) {
                                    var v18 uint16 = 1
                                    var v19 uint32 = v3
                                    _ = v18
                                    _ = v19
                                } else {
                                    fmt.Println(fi8, fi8, fi8)
                                }
                            }
                            var v20 float64 = 0.0
                            _ = v20
                        } else {
                            if ("\n\t" >= ("" + ("foo bar" + (("a" + "hello") + "foo bar")))) {
                                fmt.Println(v3, v1)
                                v3 += v3
                                if (569 >= 255) {
                                    delete(v2, 0)
                                    v21 := GenS1{F0: 406, F1: 1}
                                    var v22 uint = 18446744073709551615
                                    _ = v21
                                    _ = v22
                                }
                            }
                        }
                        v23 := GenS1{F0: 891, F1: 65535}
                        v3 -= (v3 >> ((1) & 63))
                        _ = v23
                    }
                }
                _ = fi9
            }
            v24 := GenS2{F0: (^(v3))}
            _ = v24
        }
        v25 := GenS3{F0: 622, F1: fi8}
        _ = fi8
        _ = v25
    }
    v26 := GenS4{F0: "hello"}
    if (0.1 == 0.0) {
        v27 := GenS1{F0: 4294967295, F1: 882}
        var v28 bool = false
        v29 := GenS1{F0: 4294967295, F1: 1}
        _ = v27
        _ = v28
        _ = v29
    } else {
        fmt.Println(v3, v1)
        switch "foo bar" {
        case "":
            delete(v2, 42)
        case "\n\t":
            fmt.Println(v1)
        default:
            v30 := GenS1{F0: 851, F1: 1}
            _ = v30
        }
    }
    v31 := []int{469, 42, -1, 1}
    var v32 int = 9223372036854775807
    for fi33 := 0; fi33 < 4; fi33++ {
        if true {
            _ = len(v2)
        }
        switch (v32 % (((fi33 % ((v32) | 1))) | 1)) {
        case 433:
            var v34 string = ("foo bar" + "")
            _ = v34
        case -1:
            if false {
                v35 := GenS5{F0: v32, F1: 2147483647, F2: ("a" + "hello")}
                v36 := v2[(v32 / (((fi33 | (v32 + v32))) | 1))]
                _ = v35
                _ = v36
            }
        }
        _ = fi33
    }
    fmt.Println(v1)
    for fi37 := 0; fi37 < 19; fi37++ {
        fmt.Println(fi37, v3)
        _ = fi37
    }
    v3 += v3
    if (2 != 127) {
        var v38 uint64 = 42
        fmt.Println(v38, v32, v3)
        _ = v38
    } else {
        for fi39 := 0; fi39 < 6; fi39++ {
            for fi40 := 0; fi40 < 20; fi40++ {
                var v41 int64 = -342
                var v42 uint32 = 0
                v43 := GenS6{F0: 1000, F1: true}
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
            var v44 uint = 4294967295
            for fi45 := 0; fi45 < 15; fi45++ {
                var v46 bool = true
                if (!v46) {
                    var v47 int32 = (v1 | -874)
                    _ = v47
                } else {
                    v44 += 18446744073709551615
                    if v46 {
                        fmt.Println(fi45, v46, v44)
                        if len(v31) > 0 {
                            v31[0] = 32768
                        }
                        if ((("a" + ("a" + ("hello" + "\n\t"))) != "\n\t") && (42 >= -1)) {
                            switch "a" {
                            case "hello":
                                if v46 {
                                    v1 = v1
                                    fmt.Println(v1, v1, v1)
                                }
                            case "\n\t":
                                switch "\n\t" {
                                case "":
                                    var v48 int16 = 42
                                    _ = v48
                                case "hello":
                                    if (0 == uint64(v3)) {
                                        v49 := GenS7{F0: 2}
                                        var v50 uint8 = 127
                                        v51 := GenS6{F0: int16(v3), F1: (uint32(v50) > 891)}
                                        fmt.Println(v3, v46)
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    } else {
                                        fmt.Println(v46, v44)
                                    }
                                case "a":
                                    fmt.Println(fi45, v44)
                                default:
                                    v1 = v1
                                }
                            case "foo bar":
                                if v46 {
                                    if v46 {
                                        if ((^(fi39)) >= v32) {
                                            v52 := GenS6{F0: 0, F1: false}
                                            v32 -= fi45
                                            _ = v52
                                        } else {
                                            fmt.Println(v46, v32, fi45)
                                        }
                                        v44 = v44
                                        v1 = (v1 | (v1 & (v1 / ((v1) | 1))))
                                    }
                                    v53 := GenS8{F0: 1, F1: v1}
                                    _ = v53
                                } else {
                                    switch ("\n\t" + "\n\t") {
                                    case "hello":
                                        switch "hello" {
                                        case "a":
                                            v54 := GenS5{F0: fi39, F1: fi45, F2: (("foo bar" + "\n\t") + "")}
                                            _ = v54
                                        }
                                    default:
                                        v32 += 42
                                    }
                                    switch fi45 {
                                    case 9223372036854775807:
                                        v44 = (^(v44))
                                    }
                                    var v55 int64 = 42
                                    _ = v55
                                }
                            default:
                                if false {
                                    if false {
                                        v32 = int(v3)
                                        delete(v2, (v32 & v32))
                                    } else {
                                        var v56 int32 = v1
                                        _ = v56
                                    }
                                    v32 -= -976
                                    v57 := v2[(v32 % ((fi39) | 1))]
                                    var v58 float32 = float32(0.1)
                                    _ = v57
                                    _ = v58
                                }
                            }
                            if false {
                                v44 = (v44 >> ((1) & 63))
                                switch (fi39 | (v32 << ((v44) & 63))) {
                                case 9223372036854775807:
                                    if (!(-987 <= 42)) {
                                        var v59 float32 = (3.14159 * (1.5 + (-2.25 - (float32(3.14159) - ((-2.25 + 0.0) + 1e10)))))
                                        _ = len(v2)
                                        var v60 string = ("a" + "hello")
                                        _ = v59
                                        _ = v60
                                    } else {
                                        if false {
                                            var v61 float64 = (1e10 / float64(1.5))
                                            _ = v61
                                        }
                                        if (0 <= 0) {
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(fi45, fi45, fi39)
                                            case "\n\t":
                                                v46 = ((((1.5 * 0.1) + 1.5) + 0.1) != 1.5)
                                            case "foo bar":
                                                var v62 bool = false
                                                _ = v62
                                            }
                                            v63 := GenS1{F0: 504, F1: 0}
                                            v44 = 255
                                            _ = v63
                                        }
                                        switch (("a" + "hello") + ("a" + ("foo bar" + "a"))) {
                                        case "foo bar":
                                            var v64 int16 = -1
                                            _ = v64
                                        case "hello":
                                            v46 = false
                                        case "a":
                                            v65 := GenS4{F0: ("\n\t" + "hello")}
                                            _ = v65
                                        default:
                                            fmt.Println(v46)
                                        }
                                        if v46 {
                                            v66 := GenS9{F0: v1}
                                            v67 := v2[-1]
                                            var v68 float64 = 0.0
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        } else {
                                            if false {
                                                var v69 uint64 = 18446744073709551615
                                                if v46 {
                                                    if v46 {
                                                        v70 := GenS5{F0: (^(fi39)), F1: (^(fi45)), F2: "foo bar"}
                                                        v71 := GenS9{F0: int32(fi45)}
                                                        switch v32 {
                                                        case -9223372036854775808:
                                                            v32 = fi39
                                                            fmt.Println(v32, v46)
                                                        case 2147483647:
                                                            if (!(18446744073709551615 > (v69 ^ v69))) {
                                                                if ((v44 ^ (v44 + v44)) > 0) {
                                                                    v72 := GenS10{F0: 2, F1: 65535}
                                                                    _ = v72
                                                                } else {
                                                                    v46 = v46
                                                                    v3 = v3
                                                                }
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v73 int32 = v1
                                                                    _ = v73
                                                                case "a":
                                                                    var v74 uint16 = 32767
                                                                    _ = v74
                                                                }
                                                                fmt.Println(v3, fi39, fi39)
                                                            }
                                                        }
                                                        var v75 bool = (558 >= 856)
                                                        v76 := GenS11{F0: (^(v1)), F1: -1, F2: 0}
                                                        var v77 int64 = -1
                                                        v78 := GenS7{F0: 255}
                                                        v32 -= (fi45 >> ((0) & 63))
                                                        _ = v70
                                                        _ = v71
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    if true {
                                                        fmt.Println(fi39, v69)
                                                    } else {
                                                        switch v32 {
                                                        case 32768:
                                                            var v79 int32 = (v1 >> (((v44 % ((1) | 1))) & 63))
                                                            _ = v79
                                                        default:
                                                            var v80 float64 = 0.1
                                                            switch "hello" {
                                                            case "a":
                                                                if v46 {
                                                                    if v46 {
                                                                        v81 := GenS4{F0: ("hello" + "hello")}
                                                                        _ = v81
                                                                    } else {
                                                                        var v82 uint32 = 65535
                                                                        v83 := GenS12{F0: -187, F1: -1, F2: (v44 & v44)}
                                                                        if (!v46) {
                                                                            var v84 uint = v44
                                                                            v44 -= (v84 - v84)
                                                                            v85 := v2[fi39]
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                    if false {
                                                                        v86 := GenS3{F0: uint64(v44), F1: (fi45 - (fi39 - (^(fi45))))}
                                                                        var v87 uint64 = uint64(v3)
                                                                        var v88 string = ""
                                                                        v32 = 0
                                                                        _ = v86
                                                                        _ = v87
                                                                        _ = v88
                                                                    }
                                                                    if (v69 != v69) {
                                                                        var v89 int32 = (-(v1))
                                                                        var v90 uint = v44
                                                                        _ = v89
                                                                        _ = v90
                                                                    } else {
                                                                        v32 -= fi39
                                                                        v80 += 0.1
                                                                    }
                                                                } else {
                                                                    var v91 uint16 = 1
                                                                    _ = v91
                                                                }
                                                            default:
                                                                if false {
                                                                    var v92 int8 = 0
                                                                    _ = v92
                                                                }
                                                            }
                                                            v93 := GenS11{F0: (v1 % ((v1) | 1)), F1: 0, F2: v69}
                                                            _ = v80
                                                            _ = v93
                                                        }
                                                        fmt.Println(fi45, v46)
                                                        if v46 {
                                                            v44 = 1
                                                            switch (-(fi39)) {
                                                            case -108:
                                                                v94 := GenS8{F0: 1000, F1: v1}
                                                                _ = v94
                                                            case -724:
                                                                fmt.Println(v44, fi39, fi45)
                                                            }
                                                            var v95 int32 = v1
                                                            v96 := GenS2{F0: (^(v3))}
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            var v97 uint8 = 44
                                                            if (!false) {
                                                                v3 -= 42
                                                                switch (fi39 / ((v32) | 1)) {
                                                                case 32768:
                                                                    _ = len(v31)
                                                                }
                                                            }
                                                            v32 = (fi45 << (((v44 | uint(v69))) & 63))
                                                            _ = v97
                                                        }
                                                        if (1 == (v3 + 4294967295)) {
                                                            v98 := GenS11{F0: (v1 | v1), F1: 1, F2: 0}
                                                            var v99 uint64 = v69
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            var v100 int = (fi45 + (v32 | (v32 / (((fi39 / (((-(fi45))) | 1))) | 1))))
                                                            if false {
                                                                var v101 bool = true
                                                                if (1e10 <= -2.25) {
                                                                    var v102 int64 = 0
                                                                    _ = v102
                                                                } else {
                                                                    var v103 uint32 = uint32(v69)
                                                                    if (((^(v32)) <= int(v1)) && (42 <= -1)) {
                                                                        var v104 int16 = 32767
                                                                        v2[fi45] = (!(!v46))
                                                                        switch "a" {
                                                                        case "":
                                                                            v105 := GenS6{F0: (^(v104)), F1: (v69 > (v69 | v69))}
                                                                            _ = v105
                                                                        case "a":
                                                                            var v106 uint32 = (v3 - v103)
                                                                            _ = v106
                                                                        case "hello":
                                                                            if v46 {
                                                                                if ((1 < 2) && v46) {
                                                                                    var v107 int16 = (^(v104))
                                                                                    _ = v107
                                                                                } else {
                                                                                    if (-64 < -72) {
                                                                                        v108 := GenS8{F0: (v104 / (((^(v104))) | 1)), F1: (-(v1))}
                                                                                        _ = v108
                                                                                    } else {
                                                                                        delete(v2, v100)
                                                                                        fmt.Println(fi39, v46)
                                                                                    }
                                                                                }
                                                                                v109 := GenS10{F0: 167, F1: 32767}
                                                                                _ = v109
                                                                            } else {
                                                                                if v46 {
                                                                                    if v46 {
                                                                                        var v110 int8 = -128
                                                                                        var v111 bool = false
                                                                                        var v112 uint = v44
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    }
                                                                                    fmt.Println(v44, v101)
                                                                                    v44 = 396
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v113 bool = (674 < uint32(v69))
                                                                            _ = v113
                                                                        }
                                                                        if false {
                                                                            v3 -= (v3 & 541)
                                                                            v103 = (v3 + v103)
                                                                            if v46 {
                                                                                if (((1e10 * -2.25) / float64(float32(-2.25))) > -2.25) {
                                                                                    var v114 int64 = 50
                                                                                    var v115 int64 = int64(v44)
                                                                                    v46 = v101
                                                                                    _ = v114
                                                                                    _ = v115
                                                                                }
                                                                                v104 -= 32767
                                                                                if v46 {
                                                                                    v32 = -1
                                                                                    switch 128 {
                                                                                    case -638:
                                                                                        v100 += (v32 - fi39)
                                                                                    case 32767:
                                                                                        v116 := GenS13{F0: (^(v3)), F1: ""}
                                                                                        _ = v116
                                                                                    default:
                                                                                        v117 := GenS12{F0: (v104 << ((v44) & 63)), F1: 0, F2: (v44 - uint(v69))}
                                                                                        _ = v117
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        var v118 string = "foo bar"
                                                                                        _ = v118
                                                                                    default:
                                                                                        if v101 {
                                                                                            if ((!false) || ((!v46) && v46)) {
                                                                                                v119 := GenS4{F0: ("" + ("\n\t" + (("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))) + ("\n\t" + "foo bar"))))}
                                                                                                var v120 uint8 = 1
                                                                                                if len(v31) >= 2 {
                                                                                                    v31 = v31[1:len(v31)]
                                                                                                }
                                                                                                if v46 {
                                                                                                    v121 := GenS7{F0: 419}
                                                                                                    switch (v32 ^ fi45) {
                                                                                                    case 128:
                                                                                                        if v101 {
                                                                                                            v3 += v103
                                                                                                            v69 -= v69
                                                                                                            var v122 string = "foo bar"
                                                                                                            var v123 uint64 = v69
                                                                                                            v32 += (v100 % ((-85) | 1))
                                                                                                            _ = v122
                                                                                                            _ = v123
                                                                                                        } else {
                                                                                                            v124 := GenS14{F0: (("\n\t" + "a") + (("foo bar" + ("foo bar" + "a")) + ("a" + "\n\t")))}
                                                                                                            if (v101 && (0.1 < 0.1)) {
                                                                                                                var v125 uint64 = v69
                                                                                                                v44 = (v44 - 1)
                                                                                                                _ = v125
                                                                                                            }
                                                                                                            v2[int(v3)] = true
                                                                                                            switch "" {
                                                                                                            case "foo bar":
                                                                                                                var v126 int = fi39
                                                                                                                _ = v126
                                                                                                            case "hello":
                                                                                                                v127 := GenS15{F0: v104, F1: v1}
                                                                                                                _ = v127
                                                                                                            default:
                                                                                                                v31 = append(v31, v100)
                                                                                                            }
                                                                                                            _ = v124
                                                                                                        }
                                                                                                    case -398:
                                                                                                        var v128 uint = (^(v44))
                                                                                                        var v129 uint16 = 2
                                                                                                        v130 := GenS16{F0: 260, F1: 32, F2: uint8(v44)}
                                                                                                        _ = v128
                                                                                                        _ = v129
                                                                                                        _ = v130
                                                                                                    }
                                                                                                    v131 := GenS15{F0: v104, F1: v1}
                                                                                                    _ = v121
                                                                                                    _ = v131
                                                                                                } else {
                                                                                                    switch "foo bar" {
                                                                                                    case "foo bar":
                                                                                                        switch ("" + "") {
                                                                                                        case "hello":
                                                                                                            var v132 int32 = (v1 >> (((v44 + v44)) & 63))
                                                                                                            _ = v132
                                                                                                        case "":
                                                                                                            v44 = v44
                                                                                                        }
                                                                                                    case "a":
                                                                                                        v133 := GenS13{F0: 4294967295, F1: "hello"}
                                                                                                        _ = v133
                                                                                                    }
                                                                                                    if v46 {
                                                                                                        fmt.Println(v69, v3, v101)
                                                                                                    }
                                                                                                    fmt.Println(fi39, fi39)
                                                                                                }
                                                                                                _ = v119
                                                                                                _ = v120
                                                                                            } else {
                                                                                                v134 := GenS7{F0: 42}
                                                                                                if v46 {
                                                                                                    if v46 {
                                                                                                        v135 := GenS16{F0: (-(v32)), F1: int8(v104), F2: 255}
                                                                                                        _ = v135
                                                                                                    }
                                                                                                } else {
                                                                                                    if (v46 && v101) {
                                                                                                        switch fi45 {
                                                                                                        case 32768:
                                                                                                            v69 += (v69 / ((1) | 1))
                                                                                                        }
                                                                                                    }
                                                                                                    v1 = (v1 >> ((v44) & 63))
                                                                                                }
                                                                                                v136 := GenS15{F0: v104, F1: (v1 >> ((v44) & 63))}
                                                                                                _ = v134
                                                                                                _ = v136
                                                                                            }
                                                                                        } else {
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                v137 := GenS10{F0: 32767, F1: 2}
                                                                                                _ = v137
                                                                                            case "":
                                                                                                v138 := GenS10{F0: 65535, F1: 255}
                                                                                                _ = v138
                                                                                            case "a":
                                                                                                if (2 >= 255) {
                                                                                                    switch "\n\t" {
                                                                                                    case "a":
                                                                                                        if ((v46 && (!(!(v46 || (116 >= (v3 + v3)))))) || false) {
                                                                                                            v101 = (v44 == 18446744073709551615)
                                                                                                        }
                                                                                                    default:
                                                                                                        switch (^(v32)) {
                                                                                                        case 782:
                                                                                                            if len(v31) >= 2 {
                                                                                                                v31 = v31[1:len(v31)]
                                                                                                            }
                                                                                                        case 2147483647:
                                                                                                            if (65 < 255) {
                                                                                                                v139 := GenS2{F0: (v3 | (v103 & 328))}
                                                                                                                fmt.Println(fi39, fi45, v104)
                                                                                                                var v140 uint32 = v103
                                                                                                                _ = v139
                                                                                                                _ = v140
                                                                                                            }
                                                                                                        default:
                                                                                                            var v141 int8 = -107
                                                                                                            _ = v141
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    switch (fi45 - v32) {
                                                                                                    case 42:
                                                                                                        if (42 >= 1) {
                                                                                                            var v142 uint32 = 2
                                                                                                            _ = v142
                                                                                                        }
                                                                                                    case 408:
                                                                                                        v103 += v3
                                                                                                    case -1:
                                                                                                        if len(v31) > 0 {
                                                                                                            v31[0] = 42
                                                                                                        }
                                                                                                        fmt.Println(v103, fi45)
                                                                                                        var v143 float32 = (float32(1e10) - float32(float64(0.0)))
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    v144 := GenS10{F0: uint16(v69), F1: 2}
                                                                                                    _ = v144
                                                                                                }
                                                                                            }
                                                                                            var v145 string = "a"
                                                                                            _ = v145
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v146 := v2[v32]
                                                                                    if false {
                                                                                        _ = len(v31)
                                                                                        if v146 {
                                                                                            var v147 int = 2147483647
                                                                                            if false {
                                                                                                var v148 int = (v147 ^ (fi45 << ((634) & 63)))
                                                                                                if false {
                                                                                                    var v149 float32 = (1.5 / float32(3.14159))
                                                                                                    if true {
                                                                                                        switch -1 {
                                                                                                        case -321:
                                                                                                            switch v32 {
                                                                                                            case -490:
                                                                                                                if len(v31) > 0 {
                                                                                                                    v31[0] = v148
                                                                                                                }
                                                                                                            case -911:
                                                                                                                v104 += int16(v32)
                                                                                                            }
                                                                                                        case 9223372036854775807:
                                                                                                            delete(v2, fi45)
                                                                                                        default:
                                                                                                            v147 = v32
                                                                                                        }
                                                                                                        fmt.Println(v32, v147, fi39)
                                                                                                        v3 -= (v3 - v103)
                                                                                                    }
                                                                                                    _ = v149
                                                                                                }
                                                                                                _ = v148
                                                                                            } else {
                                                                                                switch fi45 {
                                                                                                case -1:
                                                                                                    v150 := v2[fi45]
                                                                                                    _ = v150
                                                                                                case -39:
                                                                                                    var v151 uint = (v44 >> (((v44 % (((v44 << (((^(v44))) & 63))) | 1))) & 63))
                                                                                                    _ = v151
                                                                                                }
                                                                                                if v101 {
                                                                                                    v152 := GenS15{F0: (v104 & v104), F1: (v1 + (v1 << ((538) & 63)))}
                                                                                                    _ = v152
                                                                                                } else {
                                                                                                    v44 = v44
                                                                                                    var v153 uint = v44
                                                                                                    _ = v153
                                                                                                }
                                                                                            }
                                                                                            _ = v147
                                                                                        } else {
                                                                                            var v154 int64 = 0
                                                                                            var v155 uint16 = 2
                                                                                            if (32767 >= 42) {
                                                                                                var v156 uint64 = v69
                                                                                                _ = v156
                                                                                            } else {
                                                                                                delete(v2, (v32 << ((0) & 63)))
                                                                                            }
                                                                                            _ = v154
                                                                                            _ = v155
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v69, v100, v103)
                                                                                        var v157 uint16 = 142
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            var v158 int32 = (-(v1))
                                                                                            _ = v158
                                                                                        case "":
                                                                                            v157 += uint16(v3)
                                                                                            var v159 uint8 = 42
                                                                                            fmt.Println(v46, fi39)
                                                                                            _ = v159
                                                                                        case "foo bar":
                                                                                            if v101 {
                                                                                                fmt.Println(v104)
                                                                                            } else {
                                                                                                v69 = 599
                                                                                                var v160 string = ("\n\t" + "foo bar")
                                                                                                _ = v160
                                                                                            }
                                                                                            var v161 string = "hello"
                                                                                            v157 = 412
                                                                                            if ((((float64((1e10 / 3.14159)) * float64(0.0)) - 1.5) - float64(1.5)) > float64(float32(1.5))) {
                                                                                                v162 := GenS16{F0: fi45, F1: 0, F2: 118}
                                                                                                fmt.Println(v103, v101)
                                                                                                _ = len(v2)
                                                                                                _ = v162
                                                                                            }
                                                                                            _ = v161
                                                                                        default:
                                                                                            fmt.Println(v103, v146, v104)
                                                                                        }
                                                                                        if v46 {
                                                                                            fmt.Println(fi45)
                                                                                        }
                                                                                        _ = v157
                                                                                    }
                                                                                    if false {
                                                                                        var v163 int32 = (v1 - (v1 >> (((v44 ^ v44)) & 63)))
                                                                                        var v164 string = "foo bar"
                                                                                        delete(v2, v32)
                                                                                        _ = v163
                                                                                        _ = v164
                                                                                    }
                                                                                    _ = v146
                                                                                }
                                                                            }
                                                                        }
                                                                        v104 -= (v104 - v104)
                                                                        _ = v104
                                                                    } else {
                                                                        v46 = (!v46)
                                                                        v100 += 0
                                                                    }
                                                                    v165 := GenS4{F0: ("\n\t" + "foo bar")}
                                                                    _ = v103
                                                                    _ = v165
                                                                }
                                                                _ = v101
                                                            } else {
                                                                var v166 uint8 = 127
                                                                _ = v166
                                                            }
                                                            _ = v100
                                                        }
                                                        if false {
                                                            switch "" {
                                                            case "foo bar":
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v167 uint32 = 2147483648
                                                                    _ = v167
                                                                case "a":
                                                                    var v168 uint8 = 128
                                                                    _ = v168
                                                                default:
                                                                    switch ("hello" + ("" + ("a" + ("hello" + ("foo bar" + "hello"))))) {
                                                                    case "\n\t":
                                                                        var v169 int16 = int16(v1)
                                                                        _ = v169
                                                                    case "foo bar":
                                                                        v170 := GenS8{F0: -356, F1: (v1 + v1)}
                                                                        _ = v170
                                                                    case "hello":
                                                                        v46 = v46
                                                                    default:
                                                                        var v171 uint64 = (v69 ^ 1)
                                                                        if (v44 < (v44 % ((v44) | 1))) {
                                                                            switch fi45 {
                                                                            case 400:
                                                                                var v172 bool = false
                                                                                _ = v172
                                                                            case 128:
                                                                                if (0.0 != 1.5) {
                                                                                    v2[fi45] = ((v44 - v44) == (v44 << ((0) & 63)))
                                                                                    v3 = (v3 - v3)
                                                                                }
                                                                            default:
                                                                                var v173 uint16 = 42
                                                                                _ = v173
                                                                            }
                                                                            v2[(^(v32))] = ((v46 && true) || ((!(v46 || (!((0 < 1) && (!(!((1.5 >= 3.14159) && (251 >= 42)))))))) || v46))
                                                                            if v46 {
                                                                                v174 := GenS5{F0: v32, F1: -548, F2: "foo bar"}
                                                                                var v175 int16 = -416
                                                                                var v176 int32 = v1
                                                                                v171 -= v69
                                                                                _ = v174
                                                                                _ = v175
                                                                                _ = v176
                                                                            }
                                                                        }
                                                                        _ = v171
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                if v46 {
                                                                    v2[fi39] = false
                                                                    var v177 int32 = (v1 + v1)
                                                                    _ = v177
                                                                } else {
                                                                    switch (("" + "foo bar") + "hello") {
                                                                    case "":
                                                                        v44 = v44
                                                                    case "foo bar":
                                                                        v32 = v32
                                                                    case "a":
                                                                        if (65535 <= 1) {
                                                                            var v178 float64 = (1.5 * -2.25)
                                                                            switch ("" + "foo bar") {
                                                                            case "foo bar":
                                                                                var v179 int64 = 42
                                                                                _ = v179
                                                                            default:
                                                                                if len(v31) > 0 {
                                                                                    v31[0] = fi39
                                                                                }
                                                                            }
                                                                            v3 -= v3
                                                                            var v180 bool = ((!(!((0.1 <= (v178 - 3.14159)) || v46))) || (!true))
                                                                            _ = v178
                                                                            _ = v180
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            fmt.Println(fi39)
                                                                            delete(v2, fi39)
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v181 uint32 = (v3 * (v3 >> ((109) & 63)))
                                                                                _ = v181
                                                                            case "\n\t":
                                                                                fmt.Println(v46, v32)
                                                                            case "a":
                                                                                if v46 {
                                                                                    v69 = (v69 ^ (v69 & (v69 ^ v69)))
                                                                                } else {
                                                                                    fmt.Println(v32, v69)
                                                                                    v32 -= (v32 * (-(fi39)))
                                                                                    v182 := GenS1{F0: 18446744073709551615, F1: 65535}
                                                                                    var v183 uint32 = 2147483647
                                                                                    _ = v182
                                                                                    _ = v183
                                                                                }
                                                                            default:
                                                                                v46 = (0.1 <= float64(-2.25))
                                                                            }
                                                                            switch "" {
                                                                            case "a":
                                                                                v184 := v2[fi39]
                                                                                _ = v184
                                                                            case "\n\t":
                                                                                v44 = v44
                                                                            default:
                                                                                var v185 int16 = int16(v3)
                                                                                _ = v185
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case "hello":
                                                                if (((v32 << ((v44) & 63)) != fi45) || true) {
                                                                    var v186 uint32 = v3
                                                                    _ = v186
                                                                } else {
                                                                    if false {
                                                                        var v187 float64 = (float64(1.5) - -2.25)
                                                                        _ = v187
                                                                    }
                                                                    fmt.Println(v46, v44, v46)
                                                                }
                                                            }
                                                        } else {
                                                            var v188 uint = (v44 + (v44 % (((v44 % (((v44 | (v44 % (((v44 >> (((v44 - v44)) & 63))) | 1)))) | 1))) | 1)))
                                                            _ = v188
                                                        }
                                                    }
                                                    if v46 {
                                                        switch (v32 & fi45) {
                                                        case -1:
                                                            fmt.Println(v44, v1)
                                                        case -366:
                                                            if true {
                                                                v32 += 877
                                                                var v189 uint32 = (v3 | v3)
                                                                _ = v189
                                                            }
                                                        default:
                                                            if (0.1 <= (((3.14159 / (1e10 / float32(-2.25))) / 0.1) * 1.5)) {
                                                                _ = len(v2)
                                                            }
                                                        }
                                                        var v190 uint64 = (v69 << (((v44 + v44)) & 63))
                                                        _ = v190
                                                    } else {
                                                        var v191 int = 823
                                                        fmt.Println(v3, v191)
                                                        _ = v191
                                                    }
                                                }
                                                _ = v69
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v192 string = "\n\t"
                                                    _ = v192
                                                case "hello":
                                                    fmt.Println(fi45, fi45, fi39)
                                                    v193 := GenS17{F0: v44, F1: 470}
                                                    if (220 < 42) {
                                                        if v46 {
                                                            v194 := GenS18{F0: "foo bar", F1: 32767}
                                                            _ = v194
                                                        } else {
                                                            v46 = (!((v32 >> ((0) & 63)) >= (-(v32))))
                                                            fmt.Println(v3, v44, fi45)
                                                            fmt.Println(v46, v46)
                                                        }
                                                    }
                                                    _ = v193
                                                }
                                                v46 = v46
                                                v2[(fi39 | (fi45 | (^(fi45))))] = true
                                            }
                                        }
                                    }
                                case 127:
                                    v32 += (v32 | 127)
                                case 128:
                                    v44 -= (v44 + v44)
                                }
                                var v195 uint = (^(v44))
                                v1 = (v1 | (v1 >> (((v44 - (v44 >> ((v44) & 63)))) & 63)))
                                if v46 {
                                    fmt.Println(v3)
                                    var v196 int16 = 885
                                    _ = v196
                                }
                                _ = v195
                            } else {
                                var v197 uint32 = (v3 % ((2147483648) | 1))
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                                v198 := GenS19{F0: 0, F1: 816}
                                _ = v197
                                _ = v198
                            }
                        } else {
                            if (32767 > 32767) {
                                v199 := GenS5{F0: v32, F1: (fi45 << (((v44 | v44)) & 63)), F2: (("\n\t" + "foo bar") + "foo bar")}
                                switch "" {
                                case "a":
                                    v31 = append(v31, (v32 + -463))
                                    delete(v2, -2147483648)
                                }
                                v200 := GenS2{F0: 73}
                                _ = v199
                                _ = v200
                            }
                            var v201 uint16 = 0
                            fmt.Println(v46, v32, v46)
                            _ = v201
                        }
                    } else {
                        v44 = 18446744073709551615
                    }
                    var v202 uint16 = 443
                    _ = v202
                }
                _ = fi45
                _ = v46
            }
            _ = fi39
            _ = v44
        }
        _ = len(v2)
        v3 = v3
    }
    var v203 int64 = -9223372036854775808
    var v204 uint8 = 75
    v32 += 9223372036854775807
    if true {
        v1 += v1
        if (true || true) {
            if false {
                var v205 float64 = float64(1e10)
                if true {
                    var v206 int64 = v203
                    v207 := GenS20{F0: int16(v204), F1: int16(v1), F2: 0}
                    _ = v206
                    _ = v207
                }
                for fi208 := 0; fi208 < 1; fi208++ {
                    fmt.Println(fi208)
                    v204 = (v204 >> ((255) & 63))
                    _ = fi208
                }
                _ = v205
            }
            for fi209 := 0; fi209 < 9; fi209++ {
                var v210 int32 = -394
                switch fi209 {
                case 42:
                    fmt.Println(v1)
                case 115:
                    var v211 int64 = v203
                    _ = v211
                default:
                    _ = len(v2)
                }
                _ = fi209
                _ = v210
            }
        } else {
            _ = len(v31)
            var v212 uint16 = 0
            var v213 float32 = (float32(float64(0.0)) - (1.5 / float32(((1.5 + -2.25) * 1e10))))
            _ = v212
            _ = v213
        }
    } else {
        switch v32 {
        case -9223372036854775808:
            v214 := GenS21{F0: (^(v204))}
            _ = v214
        default:
            var v215 int8 = -5
            _ = v215
        }
        _ = len(v2)
    }
    var v216 int = (v32 << ((18446744073709551615) & 63))
    if true {
        for fi217 := 0; fi217 < 13; fi217++ {
            switch int(v3) {
            case 2147483647:
                if (!false) {
                    for fi218 := 0; fi218 < 18; fi218++ {
                        var v219 float64 = 1e10
                        v3 += v3
                        var v220 int8 = genFn0(1, false)
                        _ = fi218
                        _ = v219
                        _ = v220
                    }
                    var v221 bool = true
                    if false {
                        _ = len(v31)
                    } else {
                        if (v221 && ((-(fi217)) >= 9223372036854775807)) {
                            v1 = 2147483647
                        } else {
                            v222 := GenS4{F0: "foo bar"}
                            if len(v31) > 0 {
                                v31[0] = (v32 | 136)
                            }
                            v223 := GenS22{F0: genFn0(1, (!((false || v221) && (v221 || (4294967295 > 571)))))}
                            _ = v222
                            _ = v223
                        }
                        v224 := GenS20{F0: -683, F1: -35, F2: 1}
                        if (42 >= v1) {
                            v2[9223372036854775807] = v221
                        }
                        _ = v224
                    }
                    _ = v221
                } else {
                    for fi225 := 0; fi225 < 8; fi225++ {
                        if true {
                            var v226 string = "\n\t"
                            v2[(v32 << ((2) & 63))] = true
                            if (v216 != (^(fi225))) {
                                var v227 uint32 = v3
                                _ = v227
                            }
                            _ = v226
                        }
                        var v228 int32 = 2147483647
                        v229 := v2[42]
                        _ = fi225
                        _ = v228
                        _ = v229
                    }
                    fmt.Println(v204)
                }
            case -9223372036854775808:
                for fi230 := 0; fi230 < 9; fi230++ {
                    if (1e10 >= (1.5 / (1.5 / float32(-2.25)))) {
                        var v231 float64 = -2.25
                        if len(v31) > 0 {
                            v31[0] = (-(fi230))
                        }
                        var v232 uint8 = (v204 | 42)
                        v233 := GenS7{F0: 40}
                        v203 = (v203 + (v203 | (v203 * (v203 / ((v203) | 1)))))
                        v2[(fi230 | fi230)] = (true || true)
                        _ = v231
                        _ = v232
                        _ = v233
                    }
                    var v234 uint8 = uint8(v3)
                    var v235 uint32 = v3
                    _ = fi230
                    _ = v234
                    _ = v235
                }
            }
            _ = fi217
        }
        switch "" {
        case "\n\t":
            var v236 int64 = v203
            _ = v236
        case "hello":
            switch 32767 {
            case -181:
                for fi237 := 0; fi237 < 5; fi237++ {
                    if false {
                        var v238 int64 = (v203 << ((255) & 63))
                        var v239 float32 = (float32(0.1) + 1.5)
                        if (0.0 >= -2.25) {
                            v240 := GenS2{F0: (v3 * (v3 / ((900) | 1)))}
                            _ = v240
                        }
                        _ = v238
                        _ = v239
                    } else {
                        for fi241 := 0; fi241 < 2; fi241++ {
                            fmt.Println(v3, v204, v216)
                            switch "hello" {
                            case "":
                                var v242 int16 = -1
                                _ = v242
                            default:
                                var v243 int16 = -452
                                _ = v243
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v216 = int(v203)
                            case "hello":
                                v3 += (v3 >> ((42) & 63))
                            }
                            _ = fi241
                        }
                        fmt.Println(v216, v1, v3)
                        v204 -= (v204 + 2)
                    }
                    v244 := GenS2{F0: (v3 >> ((0) & 63))}
                    _ = fi237
                    _ = v244
                }
            case 1:
                v2[4294967295] = (float64(1.5) == (1e10 + (3.14159 + 0.0)))
            case -9223372036854775808:
                v216 = (v32 | (v32 << ((42) & 63)))
            }
        case "":
            var v245 int64 = (-(v203))
            _ = v245
        default:
            if true {
                var v246 float64 = 0.0
                _ = v246
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v203:", v203)
    fmt.Println("v204:", v204)
    fmt.Println("v216:", v216)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v203
    _ = v204
    _ = v216
    _ = v26
    _ = v3
    _ = v31
    _ = v32
}
