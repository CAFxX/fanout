package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint; F1 int; F2 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64 }
func genFn0(p0 bool, p1 uint64, p2 uint) bool {
    v7 := GenS1{F0: p1, F1: 138}
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint16 = 1
            _ = fi8
            _ = v9
        }
        switch 626 {
        case -2147483648:
            p0 = true
        case 0:
            for fi10 := 0; fi10 < 14; fi10++ {
                v11 := []int{}
                if (("\n\t" + ((("foo bar" + ("" + "")) + "hello") + (("hello" + ("a" + "hello")) + "a"))) != "") {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        var v13 uint8 = 186
                        _ = fi12
                        _ = v13
                    }
                    if len(v11) > 0 {
                        v11[0] = 32767
                    }
                    if (("foo bar" + "foo bar") != "foo bar") {
                        fmt.Println(p0, p1)
                        var v14 int32 = -1
                        _ = v14
                    }
                }
                v15 := GenS1{F0: p1, F1: 255}
                _ = fi10
                _ = v11
                _ = v15
            }
        default:
            v16 := []int{}
            _ = v16
        }
        v17 := []int{-277, 140, 4294967295}
        _ = v17
    } else {
        var v18 int32 = 2147483647
        var v19 uint64 = (p1 << ((p2) & 63))
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}
func genFn1(p0 int32, p1 bool) uint {
    var v20 int16 = 1
    if (!(157 < 128)) {
        v21 := GenS1{F0: 42, F1: 398}
        if (!(!(p1 || p1))) {
            v22 := []int{-892, -2147483648, 793}
            for fi23 := 0; fi23 < 18; fi23++ {
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 uint8 = 1
                    switch int(v25) {
                    case 128:
                        v26 := GenS1{F0: 0, F1: 65535}
                        _ = v26
                    case 4294967295:
                        var v27 uint32 = 42
                        _ = v27
                    }
                    switch fi23 {
                    case 32767:
                        switch "\n\t" {
                        case "a":
                            if p1 {
                                v25 += (v25 << ((0) & 63))
                                switch -33 {
                                case 42:
                                    p0 = p0
                                case -9223372036854775808:
                                    v28 := GenS2{F0: false, F1: 1, F2: 1}
                                    _ = v28
                                case 128:
                                    if (42 == (p0 / (((p0 * p0)) | 1))) {
                                        _ = len(v22)
                                    } else {
                                        var v29 int = fi23
                                        var v30 uint64 = 18446744073709551615
                                        fmt.Println(v30)
                                        var v31 float32 = (float32((float64((1.5 * -2.25)) / (1e10 * (0.0 - -2.25)))) * 0.1)
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            } else {
                                var v32 uint32 = 1
                                var v33 int16 = (v20 * -879)
                                _ = v32
                                _ = v33
                            }
                        default:
                            var v34 uint32 = 633
                            _ = v34
                        }
                    case -2147483648:
                        switch "hello" {
                        case "hello":
                            switch (fi23 << ((506) & 63)) {
                            case 127:
                                var v35 uint32 = 65535
                                _ = v35
                            case 4294967295:
                                p0 = (p0 * (^(p0)))
                                fmt.Println(fi24)
                                fmt.Println(fi23, fi23, v25)
                            }
                        }
                    default:
                        if (-91 != 0) {
                            if len(v22) > 0 {
                                v22[0] = fi23
                            }
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                switch "\n\t" {
                case "foo bar":
                    for fi36 := 0; fi36 < 1; fi36++ {
                        fmt.Println(p1, p0)
                        p0 = (p0 - 2147483647)
                        _ = fi36
                    }
                case "a":
                    if p1 {
                        if p1 {
                            var v37 string = "a"
                            _ = v37
                        }
                        if p1 {
                            v38 := GenS3{F0: 2147483647, F1: (!(255 != 1))}
                            p0 = p0
                            var v39 uint16 = 32767
                            _ = v38
                            _ = v39
                        }
                        _ = len(v22)
                    } else {
                        v40 := GenS1{F0: 803, F1: 1}
                        _ = v40
                    }
                case "\n\t":
                    var v41 float32 = 1e10
                    _ = v41
                default:
                    var v42 int = fi23
                    _ = v42
                }
                _ = fi23
            }
            p0 -= p0
            _ = v22
        }
        fmt.Println(p0, p0)
        _ = v21
    }
    fmt.Println(v20)
    var v43 uint32 = 2
    _ = p0
    _ = p1
    _ = v20
    _ = v43
    return 0
}

func main() {
    var v1 uint8 = 95
    var v2 int = 42
    var v3 int16 = -1
    v4 := []string{"a", "a", "foo bar", ("\n\t" + ("\n\t" + "hello"))}
    var v5 float64 = 1.5
    var v6 uint = 42
    var v44 int = (-(v2))
    v45 := GenS1{F0: 859, F1: 110}
    var v46 bool = true
    switch 2147483647 {
    case 1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 9223372036854775807:
        var v47 float64 = 0.0
        _ = v47
    default:
        fmt.Println(v44, v3, v5)
    }
    v48 := GenS3{F0: v44, F1: v46}
    switch "" {
    case "a":
        if v46 {
            switch (v44 | int(v1)) {
            case 2147483647:
                v49 := GenS2{F0: (2147483647 <= 1), F1: 9223372036854775807, F2: 255}
                _ = v49
            case 9223372036854775807:
                var v50 int64 = -1
                _ = v50
            default:
                if v46 {
                    if false {
                        var v51 int32 = -184
                        for fi52 := 0; fi52 < 5; fi52++ {
                            switch (v44 - v2) {
                            case 4294967295:
                                for fi53 := 0; fi53 < 7; fi53++ {
                                    fmt.Println(v6, v5, fi53)
                                    _ = len(v4)
                                    if v46 {
                                        v54 := GenS4{F0: (^(v6))}
                                        v55 := GenS1{F0: 1, F1: 373}
                                        var v56 uint64 = 1
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    } else {
                                        if true {
                                            v57 := GenS1{F0: 2, F1: 0}
                                            _ = v57
                                        } else {
                                            v58 := GenS5{F0: 1}
                                            v59 := GenS3{F0: fi52, F1: (!v46)}
                                            _ = v58
                                            _ = v59
                                        }
                                        var v60 float32 = 0.1
                                        _ = v60
                                    }
                                    v4 = append(v4, (("" + "hello") + "\n\t"))
                                    v61 := GenS4{F0: 893}
                                    var v62 uint64 = 2
                                    v3 -= (v3 + int16(v51))
                                    v63 := GenS2{F0: v46, F1: 9223372036854775807, F2: (^(v6))}
                                    var v64 uint64 = v62
                                    _ = fi53
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                            case -2147483648:
                                fmt.Println(v1, v1, v5)
                            case 32768:
                                v65 := GenS3{F0: (v2 << ((v6) & 63)), F1: true}
                                _ = v65
                            default:
                                var v66 uint = (v6 | (v6 << (((^(v6))) & 63)))
                                _ = v66
                            }
                            var v67 int16 = (v3 << ((1) & 63))
                            for fi68 := 0; fi68 < 17; fi68++ {
                                if (4294967295 >= 1) {
                                    _ = len(v4)
                                } else {
                                    if (!v46) {
                                        var v69 float64 = v5
                                        fmt.Println(v51)
                                        v67 = -1
                                        _ = v69
                                    }
                                    _ = len(v4)
                                }
                                _ = fi68
                            }
                            _ = fi52
                            _ = v67
                        }
                        _ = v51
                    } else {
                        switch "" {
                        case "a":
                            var v70 float64 = (v5 + (v5 - (v5 + 3.14159)))
                            _ = v70
                        case "\n\t":
                            v5 += (float64(1e10) / v5)
                        }
                    }
                    if false {
                        if (-9223372036854775808 != 370) {
                            for fi71 := 0; fi71 < 8; fi71++ {
                                for fi72 := 0; fi72 < 19; fi72++ {
                                    switch "hello" {
                                    case "\n\t":
                                        var v73 uint8 = v1
                                        _ = v73
                                    case "":
                                        var v74 int64 = -997
                                        _ = v74
                                    }
                                    _ = fi72
                                }
                                _ = fi71
                            }
                        }
                        if v46 {
                            v75 := GenS6{F0: false, F1: 0}
                            _ = len(v4)
                            var v76 uint = (v6 | (^(v6)))
                            _ = v75
                            _ = v76
                        }
                    }
                    var v77 bool = v46
                    _ = v77
                } else {
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + ((("\n\t" + "a") + "\n\t") + ("foo bar" + "hello")))
                    }
                    var v78 string = (("hello" + ("hello" + (("\n\t" + "hello") + ("foo bar" + "")))) + "\n\t")
                    if (v46 || (148 > 255)) {
                        v3 -= int16(v1)
                        for fi79 := 0; fi79 < 6; fi79++ {
                            if v46 {
                                var v80 uint32 = 2
                                switch "a" {
                                case "\n\t":
                                    var v81 float64 = 1e10
                                    _ = v81
                                case "foo bar":
                                    v5 -= (-2.25 - v5)
                                }
                                switch 32767 {
                                case -9223372036854775808:
                                    v5 = (v5 + float64(-2.25))
                                case 42:
                                    if (v6 != (v6 % (((^(v6))) | 1))) {
                                        v82 := GenS3{F0: (fi79 & -1), F1: (!(v80 > 0))}
                                        var v83 string = ("foo bar" + (v78 + "a"))
                                        _ = v82
                                        _ = v83
                                    }
                                case -736:
                                    var v84 int32 = 2147483647
                                    _ = v84
                                }
                                _ = v80
                            }
                            v6 -= (v6 / (((v6 << ((v6) & 63))) | 1))
                            _ = fi79
                        }
                        var v85 uint16 = 42
                        _ = v85
                    }
                    _ = v78
                }
            }
            v86 := GenS6{F0: v46, F1: -761}
            _ = v86
        }
    case "foo bar":
        v5 += (1.5 / v5)
    }
    var v87 int = (v44 ^ (-(v2)))
    _ = len(v4)
    switch "foo bar" {
    case "a":
        var v88 uint64 = 1
        _ = v88
    default:
        var v89 uint16 = 32767
        _ = v89
    }
    for fi90 := 0; fi90 < 4; fi90++ {
        switch (v87 % ((v44) | 1)) {
        case 127:
            for fi91 := 0; fi91 < 17; fi91++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch (v2 + (^(v44))) {
                case 0:
                    switch -2147483648 {
                    case -911:
                        v6 = (v6 << ((120) & 63))
                    case 128:
                        if (v5 >= v5) {
                            if false {
                                switch "\n\t" {
                                case "a":
                                    var v92 float32 = ((-2.25 * (1.5 / ((1e10 * -2.25) - 0.0))) - 3.14159)
                                    _ = v92
                                case "":
                                    if (-2.25 != float64((1e10 / 1.0))) {
                                        var v93 uint64 = 650
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v94 uint16 = 32767
                                        _ = v93
                                        _ = v94
                                    } else {
                                        if v46 {
                                            var v95 uint64 = 1
                                            _ = v95
                                        }
                                        if v46 {
                                            if (-68 == 0) {
                                                if false {
                                                    var v96 int8 = 12
                                                    var v97 int64 = -954
                                                    var v98 uint8 = v1
                                                    _ = v96
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    fmt.Println(v5, v87, fi91)
                                                    if (v3 < v3) {
                                                        switch "a" {
                                                        case "":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        case "hello":
                                                            _ = len(v4)
                                                        default:
                                                            fmt.Println(fi90, v2)
                                                        }
                                                        v44 -= (fi91 ^ (v2 % ((int(v1)) | 1)))
                                                    }
                                                }
                                                switch "a" {
                                                case "a":
                                                    if true {
                                                        v99 := GenS3{F0: int(v3), F1: ((0 > -43) || (!(!v46)))}
                                                        v100 := GenS7{F0: v6, F1: int(v3), F2: (v6 / ((255) | 1))}
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                case "":
                                                    v4 = append(v4, "hello")
                                                case "foo bar":
                                                    if (v46 && (657 != 2)) {
                                                        if v46 {
                                                            fmt.Println(v5, v2, v2)
                                                            var v101 uint32 = 2147483648
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v87 -= (v2 + (fi90 * (v2 * fi91)))
                                                                fmt.Println(v2)
                                                                switch (("hello" + ("foo bar" + "")) + "") {
                                                                case "":
                                                                    var v102 uint = v6
                                                                    _ = v102
                                                                default:
                                                                    fmt.Println(v101, v44, v5)
                                                                }
                                                            case "":
                                                                var v103 int16 = (v3 + (v3 & (v3 - (v3 << ((genFn1(1, v46)) & 63)))))
                                                                _ = v103
                                                            }
                                                            _ = v101
                                                        } else {
                                                            _ = len(v4)
                                                        }
                                                    } else {
                                                        v44 += (v87 ^ (v44 * (fi91 | v87)))
                                                    }
                                                }
                                                v5 = v5
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v1 -= 255
                                            } else {
                                                if v46 {
                                                    var v104 uint16 = 255
                                                    var v105 int32 = -1
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    if (!(-1 > -300)) {
                                                        var v106 int64 = -1
                                                        if len(v4) > 0 {
                                                            v4[0] = "foo bar"
                                                        }
                                                        var v107 int32 = 0
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                    v3 -= (-(v3))
                                                    v2 -= fi91
                                                }
                                                v108 := GenS8{F0: 42}
                                                fmt.Println(v46)
                                                v4 = append(v4, "a")
                                                _ = v108
                                            }
                                        } else {
                                            var v109 float32 = 0.1
                                            _ = len(v4)
                                            var v110 int16 = v3
                                            _ = v109
                                            _ = v110
                                        }
                                        switch ("" + "\n\t") {
                                        case "":
                                            v111 := GenS7{F0: (v6 | v6), F1: 1, F2: (v6 * (^(v6)))}
                                            _ = v111
                                        case "a":
                                            v112 := GenS1{F0: 18446744073709551615, F1: 255}
                                            _ = v112
                                        case "\n\t":
                                            v5 += v5
                                        }
                                        var v113 uint = 1
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                        _ = v113
                                    }
                                case "\n\t":
                                    v114 := GenS3{F0: fi90, F1: v46}
                                    _ = v114
                                }
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v115 := GenS9{F0: 0}
                            v116 := GenS7{F0: (v6 ^ v6), F1: -464, F2: 65535}
                            _ = v115
                            _ = v116
                        } else {
                            var v117 uint8 = v1
                            var v118 uint16 = 1
                            v117 = (v117 ^ v1)
                            _ = v117
                            _ = v118
                        }
                    case -876:
                        if v46 {
                            var v119 uint16 = 32767
                            _ = v119
                        }
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                case 32767:
                    if ((-(v3)) <= 42) {
                        var v120 int8 = 127
                        _ = v120
                    } else {
                        var v121 int16 = -1
                        v6 = genFn1(2147483647, v46)
                        v122 := GenS4{F0: (^(v6))}
                        v123 := GenS4{F0: (^(v6))}
                        var v124 string = "hello"
                        _ = v121
                        _ = v122
                        _ = v123
                        _ = v124
                    }
                case 886:
                    fmt.Println(v87, v6)
                default:
                    var v125 bool = (!(int32(v6) <= 0))
                    _ = v125
                }
                _ = fi91
            }
        case 32768:
            for fi126 := 0; fi126 < 3; fi126++ {
                v127 := GenS4{F0: (v6 | genFn1(42, (!false)))}
                v3 = -621
                _ = fi126
                _ = v127
            }
        default:
            v3 = v3
        }
        var v128 bool = (127 <= v1)
        fmt.Println(v1, v46, v2)
        _ = fi90
        _ = v128
    }
    v129 := GenS4{F0: (v6 << ((v6) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v6
    _ = v87
}
