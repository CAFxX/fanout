package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS5 struct { F0 bool; F1 int32; F2 int64 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 uint16; F1 uint32 }
type GenS12 struct { F0 bool; F1 uint8 }
type GenS13 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS14 struct { F0 uint32; F1 int8 }
type GenS15 struct { F0 uint64; F1 int; F2 int16 }
type GenS16 struct { F0 uint; F1 uint16; F2 int32 }
type GenS17 struct { F0 string; F1 uint64 }
type GenS18 struct { F0 uint64; F1 string; F2 bool }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int8; F1 int64; F2 uint32 }
type GenS21 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 int16, p1 uint8, p2 int32) int32 {
    var v6 int32 = (-(p2))
    fmt.Println(p1, v6)
    var v7 uint8 = p1
    var v8 int16 = p0
    v9 := GenS1{F0: 43, F1: ((!true) && false), F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 42
    v2 := []int{}
    var v3 int16 = 42
    var v4 int16 = (v3 & 42)
    var v5 uint32 = 2147483647
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            switch (("\n\t" + "\n\t") + "") {
            case "a":
                switch fi11 {
                case -553:
                    v4 = (v3 ^ -1)
                default:
                    switch 32768 {
                    case 4294967295:
                        fmt.Println(v4, v3, v5)
                    case 42:
                        v12 := GenS2{F0: (v5 & 4294967295), F1: 2, F2: 42}
                        _ = v12
                    }
                }
            case "":
                if (((fi11 / ((fi11) | 1)) < (fi10 - 4294967295)) || (0 <= v1)) {
                    v13 := GenS1{F0: 255, F1: false, F2: int16(fi10)}
                    v5 = (^(v5))
                    if (-1 > v4) {
                        fmt.Println(v5, v5)
                        var v14 string = ("a" + "hello")
                        v15 := GenS2{F0: 42, F1: 8, F2: (v5 << ((2) & 63))}
                        _ = v14
                        _ = v15
                    } else {
                        if (788 <= 898) {
                            var v16 uint = 65535
                            var v17 int64 = 791
                            fmt.Println(v17)
                            _ = v16
                            _ = v17
                        } else {
                            switch (fi10 << ((2) & 63)) {
                            case 32768:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 32767:
                                v4 -= v4
                            case 128:
                                var v18 int16 = (v3 % (((v3 - (-(v4)))) | 1))
                                _ = v18
                            default:
                                v3 = 0
                            }
                            v1 -= v1
                        }
                        v19 := GenS2{F0: v5, F1: 42, F2: (v5 * v5)}
                        _ = v19
                    }
                    _ = v13
                } else {
                    fmt.Println(v1, fi10)
                    switch "a" {
                    case "foo bar":
                        v5 = 0
                    case "a":
                        if true {
                            v1 -= 0
                            fmt.Println(v4, v5)
                            switch (fi10 / ((fi11) | 1)) {
                            case 183:
                                _ = len(v2)
                            case -694:
                                v2 = append(v2, (fi11 | -42))
                            case 4294967295:
                                v2 = append(v2, (fi10 ^ fi11))
                            }
                        }
                    }
                }
            }
            fmt.Println(v1)
            v5 += v5
            _ = fi11
        }
        _ = fi10
    }
    var v20 uint = 2
    if (!false) {
        switch -788 {
        case -1:
            if false {
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 uint64 = v1
                    v2 = append(v2, 972)
                    for fi23 := 0; fi23 < 17; fi23++ {
                        fmt.Println(v1, v4, v5)
                        var v24 float64 = 3.14159
                        switch fi21 {
                        case 299:
                            v22 += uint64(v20)
                        case -1:
                            var v25 int16 = v3
                            _ = v25
                        }
                        if false {
                            v20 = 1
                            var v26 int8 = 127
                            v2 = append(v2, (fi21 >> (((v20 / ((4294967295) | 1))) & 63)))
                            _ = v26
                        } else {
                            if (0 != 113) {
                                var v27 uint8 = uint8(v20)
                                _ = v27
                            }
                            v28 := GenS2{F0: v5, F1: 1, F2: v5}
                            if len(v2) > 0 {
                                v2[0] = (-(fi23))
                            }
                            _ = v28
                        }
                        _ = fi23
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
                var v29 uint32 = (v5 ^ 2147483647)
                fmt.Println(v3, v20, v1)
                _ = v29
            } else {
                if len(v2) > 0 {
                    v2[0] = 32767
                }
            }
        }
        fmt.Println(v4)
    } else {
        if true {
            v30 := GenS3{F0: v20}
            _ = v30
        } else {
            switch "" {
            case "a":
                _ = len(v2)
            }
        }
        if true {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        } else {
            v31 := GenS4{F0: v5, F1: 42, F2: -432}
            var v32 uint16 = 720
            _ = v31
            _ = v32
        }
    }
    switch ("" + "hello") {
    case "hello":
        var v33 bool = (0 >= 255)
        _ = v33
    case "\n\t":
        var v34 int8 = 1
        _ = v34
    default:
        switch (("hello" + ((("a" + ("\n\t" + "foo bar")) + "a") + ((("hello" + "") + "") + (("\n\t" + "foo bar") + "a")))) + "") {
        case "":
            if true {
                if true {
                    switch 1 {
                    case 32767:
                        v35 := GenS5{F0: (32768 == -87), F1: 0, F2: -9223372036854775808}
                        _ = v35
                    case 32768:
                        v4 += (v4 >> ((v20) & 63))
                    case 42:
                        v20 = 0
                    }
                } else {
                    v20 -= (v20 / ((4294967295) | 1))
                    var v36 uint64 = (v1 % (((v1 | 0)) | 1))
                    _ = v36
                }
            } else {
                for fi37 := 0; fi37 < 18; fi37++ {
                    if (true || (!(2147483647 > 83))) {
                        switch -703 {
                        case 32768:
                            if false {
                                switch "foo bar" {
                                case "hello":
                                    v38 := GenS6{F0: (v5 % ((87) | 1)), F1: (fi37 & (fi37 & (fi37 >> (((v20 / (((v20 - v20)) | 1))) & 63))))}
                                    _ = v38
                                case "a":
                                    v39 := GenS6{F0: uint32(v1), F1: 42}
                                    _ = v39
                                default:
                                    var v40 uint64 = v1
                                    _ = v40
                                }
                            } else {
                                v41 := GenS7{F0: 105, F1: int32(v20)}
                                if (!(false || (true || (((("hello" + "\n\t") + "\n\t") + "") >= ((("a" + "") + "\n\t") + "\n\t"))))) {
                                    for fi42 := 0; fi42 < 8; fi42++ {
                                        if false {
                                            var v43 uint16 = 42
                                            v44 := GenS8{F0: 269, F1: -1}
                                            if len(v2) > 0 {
                                                v2[0] = (fi37 >> ((v20) & 63))
                                            }
                                            _ = v43
                                            _ = v44
                                        } else {
                                            v5 -= 620
                                            v20 = 813
                                            var v45 uint8 = 255
                                            _ = v45
                                        }
                                        _ = fi42
                                    }
                                    _ = len(v2)
                                    for fi46 := 0; fi46 < 7; fi46++ {
                                        fmt.Println(v20, fi46)
                                        _ = fi46
                                    }
                                } else {
                                    var v47 bool = (1000 > 361)
                                    fmt.Println(v4)
                                    _ = v47
                                }
                                _ = v41
                            }
                        case 127:
                            switch fi37 {
                            case 1:
                                switch fi37 {
                                case -260:
                                    if false {
                                        fmt.Println(v1, v4)
                                        var v48 int32 = 0
                                        _ = v48
                                    } else {
                                        var v49 int8 = 100
                                        var v50 float32 = 3.14159
                                        _ = v49
                                        _ = v50
                                    }
                                case 4294967295:
                                    var v51 int16 = 1
                                    _ = v51
                                case -2147483648:
                                    fmt.Println(v4, v5, v4)
                                }
                            case 328:
                                if (false && true) {
                                    var v52 float32 = -2.25
                                    for fi53 := 0; fi53 < 3; fi53++ {
                                        v20 -= (v20 >> ((v20) & 63))
                                        v54 := GenS3{F0: (v20 | (v20 << (((v20 / (((v20 / ((v20) | 1))) | 1))) & 63)))}
                                        _ = fi53
                                        _ = v54
                                    }
                                    for fi55 := 0; fi55 < 8; fi55++ {
                                        var v56 int = (fi37 << ((v20) & 63))
                                        _ = fi55
                                        _ = v56
                                    }
                                    _ = v52
                                } else {
                                    v4 += -1
                                }
                            case -231:
                                var v57 int8 = -1
                                _ = v57
                            }
                        case -432:
                            var v58 bool = (((v1 % ((v1) | 1)) < v1) && (42 != 32767))
                            _ = v58
                        default:
                            var v59 float64 = 1e10
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 7; fi60++ {
                            if (!true) {
                                var v61 uint16 = 65535
                                if (-1 != 42) {
                                    v20 += v20
                                    var v62 uint64 = (v1 | (v1 | (v1 * (v1 * v1))))
                                    _ = v62
                                }
                                switch "" {
                                case "a":
                                    _ = len(v2)
                                case "foo bar":
                                    var v63 float64 = 1.5
                                    _ = v63
                                }
                                _ = v61
                            } else {
                                var v64 int8 = -128
                                var v65 uint32 = 0
                                v1 += (v1 * (v1 - 2))
                                _ = v64
                                _ = v65
                            }
                            if (v20 != v20) {
                                v1 = (v1 >> (((v20 % ((v20) | 1))) & 63))
                                v2 = append(v2, 660)
                                v20 = (v20 * v20)
                            }
                            fmt.Println(fi60, v1)
                            v2 = append(v2, fi37)
                            _ = fi60
                        }
                    }
                    _ = fi37
                }
                v3 += v4
                v2 = append(v2, 0)
            }
        }
    }
    var v66 uint32 = (v5 >> ((v20) & 63))
    fmt.Println(v4, v66, v4)
    if false {
        if (0 > 1) {
            var v67 uint64 = (v1 ^ v1)
            var v68 int64 = 9223372036854775807
            v69 := GenS3{F0: (v20 & v20)}
            _ = v67
            _ = v68
            _ = v69
        }
    } else {
        if (uint16(v5) != 65535) {
            v3 -= v4
            var v70 bool = false
            _ = v70
        } else {
            var v71 uint16 = 42
            switch -615 {
            case -1:
                var v72 float64 = 1e10
                _ = v72
            default:
                var v73 uint16 = v71
                _ = v73
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v71
        }
    }
    if false {
        switch "" {
        case "foo bar":
            v74 := GenS5{F0: (!true), F1: 488, F2: -1}
            _ = v74
        case "a":
            fmt.Println(v1, v66)
        }
        for fi75 := 0; fi75 < 7; fi75++ {
            var v76 uint16 = 609
            for fi77 := 0; fi77 < 9; fi77++ {
                v78 := GenS9{F0: -1, F1: 40}
                switch "hello" {
                case "":
                    var v79 float64 = 0.1
                    _ = v79
                case "hello":
                    if (("hello" + "") != "hello") {
                        var v80 int = (fi75 & int(v4))
                        switch ("hello" + (("foo bar" + "foo bar") + "\n\t")) {
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "\n\t":
                            var v81 int = (fi77 + fi77)
                            _ = v81
                        case "hello":
                            switch fi75 {
                            case 889:
                                fmt.Println(v4)
                            case -399:
                                if (v5 != (v5 % ((v66) | 1))) {
                                    v3 = (v3 * 42)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    switch -203 {
                                    case 32768:
                                        if len(v2) > 0 {
                                            v2[0] = (fi75 % ((-1) | 1))
                                        }
                                    case 78:
                                        var v82 float32 = 3.14159
                                        _ = v82
                                    case 42:
                                        v2 = append(v2, (fi75 * 42))
                                    default:
                                        fmt.Println(v76)
                                    }
                                } else {
                                    if false {
                                        if (1 <= 904) {
                                            v83 := GenS10{F0: (fi77 >> ((245) & 63)), F1: uint32(v20)}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v76 += v76
                                            v76 = 2
                                            _ = v83
                                        }
                                    } else {
                                        var v84 int8 = -1
                                        fmt.Println(v1)
                                        _ = v84
                                    }
                                    v85 := GenS1{F0: v20, F1: (!true), F2: v3}
                                    v20 = v20
                                    _ = v85
                                }
                            case 128:
                                switch "" {
                                case "a":
                                    switch "foo bar" {
                                    case "":
                                        var v86 int16 = 42
                                        _ = v86
                                    case "a":
                                        var v87 bool = true
                                        _ = v87
                                    default:
                                        var v88 int16 = int16(v80)
                                        _ = v88
                                    }
                                default:
                                    var v89 bool = (false && (0.0 == (-2.25 - -2.25)))
                                    _ = v89
                                }
                            }
                        default:
                            if false {
                                v80 = v80
                                if false {
                                    switch ("\n\t" + "hello") {
                                    case "":
                                        switch "a" {
                                        case "hello":
                                            var v90 int16 = (v4 ^ v3)
                                            _ = v90
                                        }
                                    }
                                    if ((v76 + (^(v76))) != v76) {
                                        if (42 > 425) {
                                            switch fi77 {
                                            case 127:
                                                var v91 uint = (v20 - 159)
                                                _ = v91
                                            case 4294967295:
                                                if false {
                                                    var v92 uint = (v20 << ((879) & 63))
                                                    _ = v92
                                                }
                                            }
                                        }
                                        var v93 float32 = 0.1
                                        _ = v93
                                    } else {
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            fmt.Println(v4)
                                        default:
                                            var v94 int16 = v4
                                            _ = v94
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if (genFn0(-1, 127, 379) >= -2147483648) {
                                            _ = len(v2)
                                            fmt.Println(v4)
                                            if (!false) {
                                                v5 = 547
                                                var v95 int8 = 127
                                                v2 = append(v2, -904)
                                                _ = v95
                                            } else {
                                                v96 := GenS5{F0: false, F1: -2147483648, F2: 9223372036854775807}
                                                if true {
                                                    v20 -= v20
                                                } else {
                                                    v97 := GenS1{F0: 0, F1: (1e10 <= 1e10), F2: -228}
                                                    var v98 int16 = v3
                                                    _ = v97
                                                    _ = v98
                                                }
                                                fmt.Println(v4, fi77)
                                                _ = v96
                                            }
                                            if (!false) {
                                                fmt.Println(fi75, v4, v80)
                                                v1 += 1
                                            } else {
                                                switch (fi75 & fi75) {
                                                case 127:
                                                    if (!((!(2147483647 != -112)) || false)) {
                                                        v2 = append(v2, fi77)
                                                        var v99 float32 = 0.1
                                                        var v100 uint16 = (v76 / ((0) | 1))
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        v20 += v20
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v101 uint64 = (v1 << ((v20) & 63))
                                                            _ = v101
                                                        case "":
                                                            if (127 >= 9) {
                                                                switch (fi77 % (((v80 << ((v20) & 63))) | 1)) {
                                                                case -890:
                                                                    v66 += 42
                                                                case 4294967295:
                                                                    v20 += v20
                                                                case 128:
                                                                    if ((((3.14159 - 3.14159) - (3.14159 / 1.5)) == (1e10 + (0.0 - (-2.25 * -2.25)))) && false) {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(v66, v4, v3)
                                                            } else {
                                                                v2 = append(v2, v80)
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v80 = 0
                                                                case "":
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v102 := GenS11{F0: v76, F1: v66}
                                                                        _ = v102
                                                                    case "\n\t":
                                                                        fmt.Println(fi77)
                                                                    }
                                                                    var v103 float32 = 3.14159
                                                                    if false {
                                                                        fmt.Println(v4, v80)
                                                                    }
                                                                    _ = v103
                                                                }
                                                                if false {
                                                                    switch -961 {
                                                                    case 32768:
                                                                        v4 -= v4
                                                                    default:
                                                                        v104 := GenS12{F0: ("hello" != ""), F1: 128}
                                                                        if (false && true) {
                                                                            v105 := GenS13{F0: (v1 << ((v20) & 63)), F1: 42, F2: v3}
                                                                            v20 -= v20
                                                                            if false {
                                                                                v106 := GenS14{F0: (v5 / ((v5) | 1)), F1: 1}
                                                                                v76 = 230
                                                                                v5 = 1
                                                                                var v107 uint = v20
                                                                                if true {
                                                                                    if true {
                                                                                        v3 = (-(v3))
                                                                                        v5 = 247
                                                                                    } else {
                                                                                        v20 -= v20
                                                                                        v108 := GenS15{F0: (v1 + v1), F1: v80, F2: (v3 | v4)}
                                                                                        _ = v108
                                                                                    }
                                                                                } else {
                                                                                    if ((^(v76)) >= v76) {
                                                                                        fmt.Println(v107, v76)
                                                                                        var v109 float64 = (3.14159 / 1.0)
                                                                                        _ = v109
                                                                                    } else {
                                                                                        var v110 float64 = 0.1
                                                                                        var v111 int32 = -2147483648
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    }
                                                                                    v112 := GenS7{F0: 246, F1: -2147483648}
                                                                                    v1 += (v1 * v1)
                                                                                    v80 += (fi75 ^ (^(v80)))
                                                                                    _ = v112
                                                                                }
                                                                                _ = v106
                                                                                _ = v107
                                                                            } else {
                                                                                var v113 int32 = -652
                                                                                if false {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = v80
                                                                                    }
                                                                                    if (true && false) {
                                                                                        v2 = append(v2, 127)
                                                                                        fmt.Println(fi75)
                                                                                    } else {
                                                                                        v4 -= -32768
                                                                                        var v114 int64 = 9223372036854775807
                                                                                        var v115 uint = 457
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                    }
                                                                                    if true {
                                                                                        v2 = append(v2, v80)
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = (v80 << (((v20 * (v20 / ((255) | 1)))) & 63))
                                                                                        }
                                                                                        var v116 float64 = 0.0
                                                                                        _ = v116
                                                                                    } else {
                                                                                        var v117 float32 = 0.1
                                                                                        var v118 string = "a"
                                                                                        if true {
                                                                                            var v119 float64 = 0.0
                                                                                            _ = v119
                                                                                        }
                                                                                        _ = v117
                                                                                        _ = v118
                                                                                    }
                                                                                } else {
                                                                                    var v120 float32 = 0.0
                                                                                    v2 = append(v2, (-(v80)))
                                                                                    fmt.Println(v20, v1, v66)
                                                                                    if true {
                                                                                        fmt.Println(v80, v66, fi75)
                                                                                        v113 -= v113
                                                                                    } else {
                                                                                        switch -7 {
                                                                                        case 0:
                                                                                            var v121 uint = uint(v1)
                                                                                            _ = v121
                                                                                        default:
                                                                                            v5 += v5
                                                                                        }
                                                                                        v80 = fi75
                                                                                    }
                                                                                    _ = v120
                                                                                }
                                                                                var v122 float32 = 3.14159
                                                                                v1 -= uint64(v20)
                                                                                _ = v113
                                                                                _ = v122
                                                                            }
                                                                            v2 = append(v2, (-(fi77)))
                                                                            _ = v105
                                                                        } else {
                                                                            var v123 int16 = v4
                                                                            var v124 string = "a"
                                                                            _ = v123
                                                                            _ = v124
                                                                        }
                                                                        _ = v104
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            var v125 uint64 = 2
                                                            _ = v125
                                                        }
                                                        v76 = 0
                                                        if false {
                                                            if true {
                                                                var v126 int = 83
                                                                if (v76 == v76) {
                                                                    v4 += 42
                                                                    switch ("\n\t" + "a") {
                                                                    case "foo bar":
                                                                        _ = len(v2)
                                                                    case "\n\t":
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            fmt.Println(fi75)
                                                                        case "a":
                                                                            if true {
                                                                                if false {
                                                                                    v127 := GenS14{F0: 0, F1: 1}
                                                                                    fmt.Println(fi75, v20)
                                                                                    _ = v127
                                                                                }
                                                                            } else {
                                                                                switch (fi75 << ((v20) & 63)) {
                                                                                case 127:
                                                                                    fmt.Println(v4, v66, v3)
                                                                                case -1:
                                                                                    if (!(457 >= (v1 >> (((v20 | (v20 + v20))) & 63)))) {
                                                                                        v126 -= 32767
                                                                                    } else {
                                                                                        v66 -= v66
                                                                                    }
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    if true {
                                                                                        v80 = fi75
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = fi77
                                                                                            }
                                                                                            v126 += (^(v126))
                                                                                            fmt.Println(v20, v5)
                                                                                        case "hello":
                                                                                            v126 = (v80 + v80)
                                                                                        }
                                                                                        v2 = append(v2, 32768)
                                                                                        v128 := GenS16{F0: v20, F1: 32767, F2: -2147483648}
                                                                                        var v129 uint = (v20 * uint(v76))
                                                                                        v130 := GenS3{F0: v129}
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                        _ = v130
                                                                                    }
                                                                                default:
                                                                                    var v131 uint = v20
                                                                                    _ = v131
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v5, v80)
                                                                                    v126 = v80
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            var v132 int8 = int8(v3)
                                                                            _ = v132
                                                                        default:
                                                                            if (!(-1 < int64(v76))) {
                                                                                var v133 float32 = 0.0
                                                                                v80 += (-(v80))
                                                                                fmt.Println(v1)
                                                                                _ = v133
                                                                            } else {
                                                                                v134 := GenS2{F0: v5, F1: 0, F2: (v5 / (((v66 ^ (^(v66)))) | 1))}
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    var v135 uint8 = 20
                                                                                    _ = v135
                                                                                case "a":
                                                                                    var v136 uint32 = v5
                                                                                    v80 += -889
                                                                                    var v137 uint64 = (v1 + (v1 | v1))
                                                                                    _ = v136
                                                                                    _ = v137
                                                                                default:
                                                                                    if (!false) {
                                                                                        var v138 int32 = 2147483647
                                                                                        _ = len(v2)
                                                                                        _ = v138
                                                                                    }
                                                                                }
                                                                                switch "hello" {
                                                                                case "":
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = v126
                                                                                    }
                                                                                case "a":
                                                                                    v20 = (^(v20))
                                                                                default:
                                                                                    v3 -= (v4 << (((v20 % (((^(v20))) | 1))) & 63))
                                                                                }
                                                                                fmt.Println(v126)
                                                                                v1 = (v1 - 4294967295)
                                                                                _ = v134
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        if (0 > 127) {
                                                                            switch fi75 {
                                                                            case 0:
                                                                                var v139 uint8 = 0
                                                                                _ = v139
                                                                            }
                                                                        } else {
                                                                            v5 = v66
                                                                        }
                                                                    default:
                                                                        var v140 float64 = 1.5
                                                                        _ = v140
                                                                    }
                                                                } else {
                                                                    switch "a" {
                                                                    case "a":
                                                                        if (0 != 2147483647) {
                                                                            if true {
                                                                                if ((104 >= 2) && true) {
                                                                                    if false {
                                                                                        v141 := GenS2{F0: (^(v66)), F1: 127, F2: v66}
                                                                                        _ = len(v2)
                                                                                        _ = v141
                                                                                    } else {
                                                                                        v1 = 0
                                                                                    }
                                                                                    v5 += (v5 / (((v66 + (v66 >> (((v20 + 255)) & 63)))) | 1))
                                                                                    switch (-(v80)) {
                                                                                    case 32768:
                                                                                        v142 := GenS17{F0: ("hello" + "hello"), F1: 18446744073709551615}
                                                                                        _ = v142
                                                                                    case 42:
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = 32768
                                                                                        }
                                                                                    case 365:
                                                                                        switch ("a" + (("\n\t" + "foo bar") + ("" + "hello"))) {
                                                                                        case "hello":
                                                                                            var v143 uint16 = (v76 % (((^(v76))) | 1))
                                                                                            _ = v143
                                                                                        case "":
                                                                                            if false {
                                                                                                v144 := GenS15{F0: v1, F1: 0, F2: (v3 | int16(v20))}
                                                                                                _ = v144
                                                                                            }
                                                                                            v4 += (v4 ^ (v3 >> ((255) & 63)))
                                                                                            _ = len(v2)
                                                                                        case "a":
                                                                                            v3 -= v3
                                                                                        default:
                                                                                            if (197 < uint8(v20)) {
                                                                                                v145 := GenS2{F0: v66, F1: 0, F2: v5}
                                                                                                v2 = append(v2, (v126 * -847))
                                                                                                _ = v145
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v5, fi75, v76)
                                                                                    }
                                                                                }
                                                                                fmt.Println(v4)
                                                                                if (532 == -9223372036854775808) {
                                                                                    v76 -= (v76 | v76)
                                                                                } else {
                                                                                    if (true && (((("hello" + "a") + ("\n\t" + ("\n\t" + "foo bar"))) + "\n\t") >= "\n\t")) {
                                                                                        v146 := GenS11{F0: v76, F1: 28}
                                                                                        v147 := GenS18{F0: (^(v1)), F1: "\n\t", F2: (!(v20 != v20))}
                                                                                        v80 = (v126 + 4294967295)
                                                                                        _ = v146
                                                                                        _ = v147
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            v66 = 325
                                                                                            switch fi75 {
                                                                                            case 42:
                                                                                                v5 = v66
                                                                                            }
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = int(v76)
                                                                                            }
                                                                                        case "foo bar":
                                                                                            _ = len(v2)
                                                                                        default:
                                                                                            switch 32768 {
                                                                                            case 273:
                                                                                                v66 += 2147483647
                                                                                            }
                                                                                            v126 += fi77
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    v76 = 1
                                                                                }
                                                                                if (v126 > -565) {
                                                                                    var v148 uint8 = 143
                                                                                    v2 = append(v2, fi75)
                                                                                    v76 += v76
                                                                                    _ = v148
                                                                                } else {
                                                                                    if ((3.14159 * 0.1) > 1e10) {
                                                                                        v149 := GenS11{F0: v76, F1: (v66 % (((^(v5))) | 1))}
                                                                                        _ = v149
                                                                                    } else {
                                                                                        var v150 int32 = -1
                                                                                        _ = v150
                                                                                    }
                                                                                }
                                                                                if (v5 >= (v66 << (((v20 << (((v20 / (((v20 / (((^(v20))) | 1))) | 1))) & 63))) & 63))) {
                                                                                    var v151 uint32 = v5
                                                                                    v151 += 2147483647
                                                                                    _ = v151
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v152 uint = (^(v20))
                                                                            _ = v152
                                                                        }
                                                                    default:
                                                                        var v153 int = (fi77 - fi75)
                                                                        _ = v153
                                                                    }
                                                                }
                                                                _ = v126
                                                            }
                                                            if ((3.14159 == float64((3.14159 - 0.0))) || true) {
                                                                v154 := GenS5{F0: false, F1: -902, F2: 969}
                                                                fmt.Println(v4, v3, fi77)
                                                                switch (fi75 & (fi77 - fi75)) {
                                                                case 32767:
                                                                    v3 -= (-(v3))
                                                                default:
                                                                    var v155 uint64 = (v1 << ((780) & 63))
                                                                    _ = v155
                                                                }
                                                                var v156 int16 = 1
                                                                _ = v154
                                                                _ = v156
                                                            } else {
                                                                if (!false) {
                                                                    var v157 int8 = 42
                                                                    v1 = (v1 ^ v1)
                                                                    _ = v157
                                                                }
                                                                v20 += 297
                                                                _ = len(v2)
                                                            }
                                                            v1 -= v1
                                                        } else {
                                                            if true {
                                                                if len(v2) > 0 {
                                                                    v2[0] = (^(v80))
                                                                }
                                                                v158 := GenS19{F0: 42}
                                                                _ = v158
                                                            }
                                                        }
                                                    }
                                                case 32767:
                                                    var v159 uint64 = (v1 & 204)
                                                    _ = v159
                                                }
                                                var v160 uint = v20
                                                _ = v160
                                            }
                                        } else {
                                            var v161 uint32 = v5
                                            v80 = v80
                                            _ = v161
                                        }
                                        if false {
                                            var v162 int16 = v4
                                            _ = v162
                                        }
                                        switch (^(fi77)) {
                                        case 128:
                                            v2 = append(v2, fi77)
                                        case -9223372036854775808:
                                            if len(v2) > 0 {
                                                v2[0] = (fi77 | v80)
                                            }
                                        case 1:
                                            v5 += uint32(v76)
                                        }
                                    }
                                    v20 -= uint(v5)
                                } else {
                                    switch "" {
                                    case "a":
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if (48 <= 98) {
                                                var v163 string = "hello"
                                                fmt.Println(v163)
                                                v164 := GenS14{F0: (v66 >> ((131) & 63)), F1: 122}
                                                _ = v163
                                                _ = v164
                                            }
                                            var v165 int32 = genFn0((v4 - 0), 0, 1)
                                            _ = v165
                                        }
                                    case "hello":
                                        var v166 int64 = -743
                                        _ = v166
                                    }
                                }
                                if true {
                                    if false {
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = (fi75 - fi75)
                                            }
                                        } else {
                                            var v167 uint32 = (v5 << ((18446744073709551615) & 63))
                                            v20 = v20
                                            fmt.Println(v76)
                                            _ = v167
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        switch 0 {
                                        case 9223372036854775807:
                                            v168 := GenS19{F0: 2147483647}
                                            _ = v168
                                        case 1:
                                            _ = len(v2)
                                        case 32768:
                                            _ = len(v2)
                                        default:
                                            var v169 string = ("a" + "a")
                                            _ = v169
                                        }
                                        var v170 uint16 = v76
                                        var v171 int16 = v4
                                        _ = v170
                                        _ = v171
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                var v172 int8 = 39
                                _ = len(v2)
                                if (-997 == genFn0(int16(v172), 122, 42)) {
                                    if false {
                                        var v173 bool = ((v76 | (v76 / ((v76) | 1))) >= v76)
                                        _ = v173
                                    }
                                    v2 = append(v2, v80)
                                } else {
                                    if (v66 <= v5) {
                                        v66 += (v66 % ((v66) | 1))
                                        v3 -= (v3 | v4)
                                    }
                                    var v174 uint8 = 1
                                    if true {
                                        if (true || (!(0 <= -9223372036854775808))) {
                                            v76 = uint16(v20)
                                            v2 = append(v2, (v80 ^ (fi75 ^ v80)))
                                            switch "hello" {
                                            case "hello":
                                                v80 = (fi77 & fi77)
                                                v175 := GenS9{F0: (v172 | (v172 % ((68) | 1))), F1: (v66 ^ (v5 & (v66 + (^(v66)))))}
                                                v176 := GenS5{F0: (true && false), F1: -860, F2: 129}
                                                _ = v175
                                                _ = v176
                                            case "":
                                                v4 = 982
                                            }
                                        } else {
                                            switch (("a" + "a") + "") {
                                            case "hello":
                                                if (!((2147483648 != uint32(v1)) || (!((0.0 - -2.25) >= (1.5 * float64(1.5)))))) {
                                                    var v177 uint16 = v76
                                                    _ = v177
                                                } else {
                                                    var v178 uint = 2
                                                    v179 := GenS9{F0: v172, F1: (^(v66))}
                                                    var v180 uint64 = v1
                                                    _ = v178
                                                    _ = v179
                                                    _ = v180
                                                }
                                            case "a":
                                                v2 = append(v2, (-(v80)))
                                                switch fi77 {
                                                case -9223372036854775808:
                                                    var v181 int32 = 2147483647
                                                    _ = v181
                                                case 0:
                                                    v174 = 2
                                                case 474:
                                                    v182 := GenS15{F0: uint64(v20), F1: (fi75 ^ (fi77 << (((v20 / ((18446744073709551615) | 1))) & 63))), F2: (v4 - -840)}
                                                    _ = v182
                                                default:
                                                    var v183 uint64 = (^(v1))
                                                    _ = v183
                                                }
                                            }
                                            switch "\n\t" {
                                            case "":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "hello":
                                                if len(v2) > 0 {
                                                    v2[0] = int(v20)
                                                }
                                            case "\n\t":
                                                switch v80 {
                                                case 9223372036854775807:
                                                    var v184 int32 = -788
                                                    _ = v184
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = int(v1)
                                                }
                                            }
                                            v66 = 2147483647
                                        }
                                        fmt.Println(v172)
                                        if ((true && false) || (!false)) {
                                            var v185 float32 = (-2.25 - -2.25)
                                            _ = v185
                                        } else {
                                            var v186 string = "a"
                                            if (65535 == v5) {
                                                var v187 uint16 = (v76 * (v76 - v76))
                                                _ = v187
                                            } else {
                                                v188 := GenS8{F0: v4, F1: -171}
                                                switch "hello" {
                                                case "":
                                                    if ((840 > -1) && (1 == 1)) {
                                                        v2 = append(v2, (-(fi77)))
                                                    }
                                                    if true {
                                                        v189 := GenS16{F0: (v20 / ((527) | 1)), F1: 719, F2: -2147483648}
                                                        var v190 int16 = 32767
                                                        _ = v189
                                                        _ = v190
                                                    } else {
                                                        switch ((v186 + v186) + v186) {
                                                        case "hello":
                                                            v191 := GenS20{F0: (v172 * v172), F1: 399, F2: 621}
                                                            _ = v191
                                                        case "a":
                                                            switch v186 {
                                                            case "hello":
                                                                v1 -= v1
                                                            default:
                                                                v20 += v20
                                                            }
                                                        case "":
                                                            v192 := GenS4{F0: (v66 ^ (^(v66))), F1: (v174 + (^(v174))), F2: int32(v1)}
                                                            _ = v192
                                                        }
                                                        switch (v186 + "a") {
                                                        case "\n\t":
                                                            var v193 float64 = -2.25
                                                            _ = v193
                                                        default:
                                                            if (!false) {
                                                                var v194 int64 = -9223372036854775808
                                                                var v195 float32 = (1.5 / float32((0.1 - float64(float32(((-2.25 / 1.5) * ((1.5 * -2.25) - 3.14159)))))))
                                                                switch (-(fi77)) {
                                                                case 9223372036854775807:
                                                                    switch 0 {
                                                                    case 4294967295:
                                                                        if false {
                                                                            var v196 int32 = 1
                                                                            _ = v196
                                                                        } else {
                                                                            var v197 int64 = v194
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v197
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            var v198 float32 = 3.14159
                                                                            if true {
                                                                                v76 += 0
                                                                                v1 = (v1 & (v1 * (v1 & v1)))
                                                                            }
                                                                            _ = v198
                                                                        }
                                                                    }
                                                                case 1:
                                                                    var v199 float32 = v195
                                                                    _ = v199
                                                                case -792:
                                                                    v172 += -90
                                                                }
                                                                if false {
                                                                    var v200 int16 = int16(v194)
                                                                    _ = v200
                                                                }
                                                                var v201 uint64 = 42
                                                                _ = v194
                                                                _ = v195
                                                                _ = v201
                                                            } else {
                                                                if (true || (true || ((v174 * 42) != v174))) {
                                                                    fmt.Println(v3)
                                                                    fmt.Println(v186)
                                                                }
                                                                var v202 uint = (v20 + (v20 - (v20 | (v20 / (((v20 & v20)) | 1)))))
                                                                var v203 bool = (!true)
                                                                _ = v202
                                                                _ = v203
                                                            }
                                                        }
                                                        if false {
                                                            var v204 uint64 = (v1 >> ((0) & 63))
                                                            v5 = 0
                                                            _ = v204
                                                        } else {
                                                            v2 = append(v2, fi77)
                                                            v172 -= v172
                                                        }
                                                    }
                                                }
                                                v172 -= v172
                                                _ = v188
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v1 = (v1 & 994)
                                            if len(v2) > 0 {
                                                v2[0] = 32768
                                            }
                                            _ = v186
                                        }
                                    } else {
                                        if (1 != 163) {
                                            v1 = v1
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v205 uint64 = v1
                                            if (v3 == 1000) {
                                                v172 = v172
                                                switch fi77 {
                                                case -793:
                                                    fmt.Println(v76, v205, v20)
                                                case 128:
                                                    var v206 float32 = (-2.25 + ((1e10 * 0.0) * (0.1 - 3.14159)))
                                                    _ = v206
                                                }
                                            }
                                            v4 += (v4 - 1000)
                                            var v207 float64 = 0.0
                                            _ = v205
                                            _ = v207
                                        } else {
                                            v174 += (v174 | (v174 * v174))
                                        }
                                        var v208 int32 = 42
                                        v2 = append(v2, (fi77 - fi75))
                                        _ = v208
                                    }
                                    _ = v174
                                }
                                if ((v1 > v1) || (0.0 != 0.1)) {
                                    v76 = (v76 + (^(v76)))
                                    v209 := GenS4{F0: (^(v66)), F1: 218, F2: int32(v20)}
                                    _ = v209
                                } else {
                                    var v210 int64 = -1
                                    if (!true) {
                                        v211 := GenS21{F0: 1000, F1: 42, F2: uint16(v66)}
                                        switch (-(fi75)) {
                                        case -1:
                                            v212 := GenS20{F0: v172, F1: (v210 ^ v210), F2: v5}
                                            _ = v212
                                        default:
                                            var v213 int16 = v4
                                            _ = v213
                                        }
                                        _ = v211
                                    }
                                    fmt.Println(v20)
                                    _ = v210
                                }
                                _ = v172
                            }
                        }
                        _ = v80
                    } else {
                        v66 += (^(v5))
                        _ = len(v2)
                        v1 -= (v1 % (((^(v1))) | 1))
                    }
                }
                fmt.Println(v1, fi75, v4)
                _ = fi77
                _ = v78
            }
            var v214 uint16 = uint16(v1)
            _ = fi75
            _ = v214
            _ = v76
        }
        _ = len(v2)
    } else {
        _ = len(v2)
        switch -907 {
        case -1:
            var v215 string = (("a" + "foo bar") + ("" + "\n\t"))
            _ = v215
        }
        v1 += 607
    }
    var v216 int32 = 2147483647
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 127
    }
    switch ("hello" + ("hello" + (("foo bar" + "a") + ((("a" + "\n\t") + (("a" + ("foo bar" + "")) + (("hello" + ("hello" + "a")) + "a"))) + "foo bar")))) {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 127
        }
    case "":
        if false {
            switch "a" {
            case "":
                var v217 uint = v20
                _ = v217
            default:
                if ("foo bar" > "") {
                    v218 := GenS3{F0: (v20 % ((v20) | 1))}
                    for fi219 := 0; fi219 < 7; fi219++ {
                        for fi220 := 0; fi220 < 6; fi220++ {
                            if false {
                                v5 = (v66 | 0)
                                var v221 uint32 = uint32(v1)
                                _ = v221
                            }
                            _ = fi220
                        }
                        _ = fi219
                    }
                    _ = v218
                } else {
                    var v222 int = 128
                    _ = v222
                }
            }
            for fi223 := 0; fi223 < 4; fi223++ {
                var v224 int32 = 42
                _ = fi223
                _ = v224
            }
        } else {
            v4 += v4
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
        }
    case "hello":
        if false {
            var v225 int8 = -128
            var v226 float64 = -2.25
            _ = v225
            _ = v226
        } else {
            v3 = 1
        }
    }
    if len(v2) > 0 {
        v2[0] = int(v216)
    }
    var v227 uint = (v20 * v20)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v216:", v216)
    fmt.Println("v227:", v227)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v20
    _ = v216
    _ = v227
    _ = v3
    _ = v4
    _ = v5
    _ = v66
}
