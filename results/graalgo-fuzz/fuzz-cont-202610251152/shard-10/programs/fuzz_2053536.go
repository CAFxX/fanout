package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 int8; F1 int16; F2 uint8 }
type GenS9 struct { F0 int8; F1 int64; F2 uint32 }
func genFn0(p0 int8) uint64 {
    if (0.0 >= (1.5 * (0.1 * (1e10 * ((float32(0.0) * 3.14159) - -2.25))))) {
        var v6 float64 = ((float64((0.1 - 1.5)) + ((((-2.25 / ((1e10 * 1e10) / 0.1)) + 1e10) * (0.1 - 1e10)) * 1e10)) / 3.14159)
        _ = v6
    } else {
        p0 = p0
        v7 := []int{4294967295, 127, 473, -887}
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{(fi8 | 2147483647), fi8, (fi8 / (((fi8 % (((fi8 & (fi8 - fi8))) | 1))) | 1))}
            _ = fi8
            _ = v9
        }
        p0 -= 42
    }
    p0 -= -1
    _ = p0
    return 1
}
func genFn1(p0 uint32, p1 bool) uint8 {
    v10 := []int{2147483647}
    for fi11 := 0; fi11 < 12; fi11++ {
        if p1 {
            fmt.Println(p0, fi11)
            if true {
                var v12 int32 = 2147483647
                v13 := GenS1{F0: (p0 >> ((42) & 63)), F1: 626, F2: 1}
                v14 := GenS1{F0: (p0 % ((961) | 1)), F1: 0, F2: 2}
                _ = v12
                _ = v13
                _ = v14
            } else {
                if true {
                    var v15 uint32 = p0
                    _ = v15
                }
                if false {
                    var v16 float32 = 0.1
                    p0 -= (p0 | 609)
                    for fi17 := 0; fi17 < 6; fi17++ {
                        switch -1 {
                        case 42:
                            p1 = p1
                        case -9223372036854775808:
                            if len(v10) > 0 {
                                v10[0] = fi11
                            }
                        default:
                            p0 = 2147483648
                        }
                        p0 -= 2147483648
                        var v18 string = "hello"
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
            }
        } else {
            switch fi11 {
            case -413:
                var v19 int32 = -2147483648
                _ = v19
            }
            for fi20 := 0; fi20 < 11; fi20++ {
                if (true || (456 != 1)) {
                    var v21 int64 = 9223372036854775807
                    _ = v21
                }
                _ = fi20
            }
            switch fi11 {
            case 9223372036854775807:
                v10 = append(v10, fi11)
            case 2147483647:
                v22 := GenS2{F0: 2}
                _ = v22
            case -309:
                for fi23 := 0; fi23 < 4; fi23++ {
                    if (!(true || ((2 != 127) || p1))) {
                        fmt.Println(fi23)
                        var v24 float32 = 3.14159
                        _ = v24
                    }
                    if p1 {
                        var v25 float64 = -2.25
                        var v26 uint = 0
                        var v27 bool = false
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = fi23
                }
            default:
                if (false || p1) {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        fmt.Println(fi28)
                        _ = fi28
                    }
                    v29 := GenS3{F0: (^(p0))}
                    _ = v29
                }
            }
        }
        if true {
            p0 += p0
            for fi30 := 0; fi30 < 16; fi30++ {
                switch fi11 {
                case 824:
                    _ = len(v10)
                case -2147483648:
                    p0 = p0
                default:
                    switch 0 {
                    case 32767:
                        if p1 {
                            p0 -= 0
                            p0 = (^(p0))
                            fmt.Println(fi11)
                        }
                    case 1:
                        if (42 >= -138) {
                            p0 -= p0
                            if len(v10) > 0 {
                                v10[0] = (fi30 >> ((324) & 63))
                            }
                        } else {
                            var v31 string = "\n\t"
                            _ = v31
                        }
                    default:
                        v10 = append(v10, 128)
                    }
                }
                _ = fi30
            }
        } else {
            if (127 < 2) {
                if p1 {
                    v10 = append(v10, fi11)
                    p1 = (!p1)
                } else {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        v33 := GenS2{F0: 32767}
                        _ = fi32
                        _ = v33
                    }
                    _ = len(v10)
                }
                for fi34 := 0; fi34 < 7; fi34++ {
                    var v35 int = (fi34 << ((2) & 63))
                    if (2 >= uint(p0)) {
                        var v36 uint16 = 1
                        _ = v36
                    }
                    fmt.Println(fi34)
                    _ = fi34
                    _ = v35
                }
            }
        }
        for fi37 := 0; fi37 < 18; fi37++ {
            var v38 bool = true
            _ = fi37
            _ = v38
        }
        _ = fi11
    }
    v39 := GenS2{F0: 1}
    for fi40 := 0; fi40 < 11; fi40++ {
        v41 := GenS2{F0: 1}
        _ = fi40
        _ = v41
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v39
    return 1
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    v5 := []int{-9223372036854775808, -9223372036854775808}
    v42 := GenS3{F0: 65535}
    var v43 bool = true
    switch "a" {
    case "\n\t":
        v43 = false
    case "a":
        if (!(!true)) {
            var v44 bool = (65535 == 2147483647)
            _ = v44
        } else {
            switch ((("hello" + "\n\t") + "hello") + "hello") {
            case "foo bar":
                switch "hello" {
                case "hello":
                    for fi45 := 0; fi45 < 2; fi45++ {
                        if false {
                            fmt.Println(v43)
                            v43 = v43
                            v43 = v43
                        }
                        for fi46 := 0; fi46 < 19; fi46++ {
                            fmt.Println(v43, v43, fi46)
                            var v47 bool = v43
                            var v48 int = (-(fi46))
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                        _ = fi45
                    }
                }
            case "\n\t":
                var v49 bool = v43
                _ = v49
            }
        }
    default:
        for fi50 := 0; fi50 < 14; fi50++ {
            _ = len(v5)
            _ = fi50
        }
    }
    switch ("foo bar" + "") {
    case "foo bar":
        for fi51 := 0; fi51 < 18; fi51++ {
            var v52 int = fi51
            if (v43 && true) {
                if v43 {
                    for fi53 := 0; fi53 < 14; fi53++ {
                        v54 := GenS4{F0: 2}
                        _ = fi53
                        _ = v54
                    }
                    for fi55 := 0; fi55 < 18; fi55++ {
                        v43 = (722 < -413)
                        _ = fi55
                    }
                    v56 := GenS5{F0: 420}
                    _ = v56
                } else {
                    if ((false || (3.14159 < 1.5)) && v43) {
                        for fi57 := 0; fi57 < 11; fi57++ {
                            if v43 {
                                delete(v2, "foo bar")
                                fmt.Println(v43)
                                var v58 int32 = -899
                                _ = v58
                            } else {
                                v59 := v2["a"]
                                var v60 int64 = 9223372036854775807
                                v61 := GenS4{F0: 2}
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                            v62 := GenS6{F0: 0}
                            _ = fi57
                            _ = v62
                        }
                        fmt.Println(fi51, v43, fi51)
                        v1 = v1
                    }
                    v52 = fi51
                }
                v63 := GenS5{F0: 223}
                v64 := GenS4{F0: 440}
                _ = v63
                _ = v64
            } else {
                fmt.Println(v43, fi51)
                fmt.Println(v43, v43, v1)
                if (42 < 0) {
                    v52 = -557
                    var v65 int64 = 0
                    _ = v65
                } else {
                    v43 = false
                    if v43 {
                        v52 = (v52 / ((32767) | 1))
                        v66 := GenS3{F0: 2147483648}
                        _ = v66
                    } else {
                        v1 -= v1
                        for fi67 := 0; fi67 < 4; fi67++ {
                            v1 = (v1 | 1)
                            var v68 bool = false
                            var v69 int = (v52 >> ((280) & 63))
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                        for fi70 := 0; fi70 < 19; fi70++ {
                            v52 = 1
                            _ = fi70
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            for fi71 := 0; fi71 < 18; fi71++ {
                v43 = false
                var v72 int = (fi51 >> ((733) & 63))
                _ = len(v5)
                _ = fi71
                _ = v72
            }
            _ = fi51
            _ = v52
        }
    }
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    _ = len(v5)
    var v73 int32 = 2147483647
    if false {
        v73 += v73
    }
    switch "foo bar" {
    case "hello":
        if v43 {
            var v74 uint16 = 1
            if (false && (!(4294967295 <= 995))) {
                if true {
                    v1 += (v1 | v1)
                    switch ("a" + (("hello" + "\n\t") + "hello")) {
                    case "a":
                        v5 = append(v5, 1)
                    }
                    if ((v43 || true) && (((9223372036854775807 < -317) || (!(32767 == (v74 + (v74 ^ v74))))) && (!(!true)))) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        for fi75 := 0; fi75 < 15; fi75++ {
                            var v76 float32 = -2.25
                            switch ("hello" + "a") {
                            case "foo bar":
                                if v43 {
                                    v43 = (false || v43)
                                    v5 = append(v5, fi75)
                                }
                            }
                            _ = fi75
                            _ = v76
                        }
                        if (("a" + (("hello" + "foo bar") + "\n\t")) == "foo bar") {
                            _ = len(v5)
                            for fi77 := 0; fi77 < 13; fi77++ {
                                v73 = v73
                                v73 += (^(v73))
                                _ = len(v5)
                                _ = fi77
                            }
                            for fi78 := 0; fi78 < 8; fi78++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi78
                            }
                        }
                        v74 -= v74
                    }
                }
                switch "hello" {
                case "hello":
                    v43 = (2147483647 != 354)
                case "a":
                    for fi79 := 0; fi79 < 7; fi79++ {
                        var v80 uint16 = (v74 & (v74 / (((v74 >> ((255) & 63))) | 1)))
                        _ = fi79
                        _ = v80
                    }
                case "\n\t":
                    v81 := GenS7{F0: 113, F1: (v73 & v73)}
                    _ = v81
                default:
                    for fi82 := 0; fi82 < 13; fi82++ {
                        var v83 int32 = -2147483648
                        var v84 int32 = (v83 & v73)
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                }
                v73 += (v73 % ((v73) | 1))
            }
            if (881 != 374) {
                fmt.Println(v73)
                v85 := GenS8{F0: -16, F1: (v1 >> ((1) & 63)), F2: 42}
                _ = v85
            } else {
                v3[236] = 587
                switch 32767 {
                case -9223372036854775808:
                    v86 := v4[-9223372036854775808]
                    _ = v86
                case 127:
                    switch 32768 {
                    case 42:
                        if v43 {
                            v87 := GenS6{F0: 457}
                            if v43 {
                                var v88 uint = 65535
                                fmt.Println(v74, v1)
                                var v89 float32 = float32(3.14159)
                                _ = v88
                                _ = v89
                            }
                            var v90 int8 = 1
                            _ = v87
                            _ = v90
                        } else {
                            fmt.Println(v74, v1, v1)
                            var v91 bool = (false || true)
                            _ = v91
                        }
                    case 4294967295:
                        switch "foo bar" {
                        case "":
                            for fi92 := 0; fi92 < 5; fi92++ {
                                v93 := GenS5{F0: 255}
                                _ = len(v3)
                                _ = fi92
                                _ = v93
                            }
                        case "\n\t":
                            switch "\n\t" {
                            case "foo bar":
                                if v43 {
                                    var v94 float32 = (1.5 - 3.14159)
                                    _ = v94
                                } else {
                                    v43 = (true && true)
                                }
                            case "":
                                if v43 {
                                    var v95 uint32 = 958
                                    v95 -= v95
                                    _ = v95
                                } else {
                                    delete(v3, 65535)
                                    v96 := GenS3{F0: 1}
                                    v1 = v1
                                    _ = v96
                                }
                            }
                        }
                    case -2147483648:
                        v5 = append(v5, int(v74))
                    default:
                        var v97 int16 = v1
                        _ = v97
                    }
                case 42:
                    v98 := GenS7{F0: (v74 | (v74 % ((v74) | 1))), F1: (v73 + int32(v74))}
                    _ = v98
                default:
                    var v99 int64 = 42
                    _ = v99
                }
            }
            _ = v74
        } else {
            fmt.Println(v1)
        }
    case "a":
        switch "" {
        case "hello":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "a":
            if true {
                switch "" {
                case "hello":
                    for fi100 := 0; fi100 < 12; fi100++ {
                        v101 := GenS6{F0: 716}
                        var v102 uint64 = 1
                        _ = fi100
                        _ = v101
                        _ = v102
                    }
                }
            }
        case "foo bar":
            v103 := v2[("\n\t" + "\n\t")]
            _ = v103
        default:
            for fi104 := 0; fi104 < 19; fi104++ {
                for fi105 := 0; fi105 < 20; fi105++ {
                    var v106 uint16 = 0
                    if v43 {
                        var v107 uint8 = genFn1(65535, false)
                        var v108 int = (fi105 % (((fi105 & 2147483647)) | 1))
                        v109 := GenS5{F0: uint(v106)}
                        _ = v107
                        _ = v108
                        _ = v109
                    } else {
                        switch (fi104 >> ((0) & 63)) {
                        case -1:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 9223372036854775807:
                            v110 := GenS1{F0: 1, F1: 2, F2: v106}
                            v3[720] = 2147483647
                            _ = v110
                        case 42:
                            v111 := GenS2{F0: 65535}
                            _ = v111
                        }
                        v73 += (v73 / ((0) | 1))
                    }
                    v5 = append(v5, -723)
                    _ = fi105
                    _ = v106
                }
                _ = fi104
            }
        }
    case "foo bar":
        v73 = v73
    default:
        _ = len(v5)
    }
    var v112 int32 = 42
    fmt.Println(v1, v43, v1)
    v113 := GenS9{F0: 25, F1: 9223372036854775807, F2: 65535}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v114 uint16 = 602
    v43 = (!v43)
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v73
}
