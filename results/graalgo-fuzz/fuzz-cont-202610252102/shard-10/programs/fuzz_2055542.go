package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 int; F1 string; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int; F1 uint32; F2 int }
type GenS11 struct { F0 int8; F1 int8 }
type GenS12 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS13 struct { F0 uint; F1 int8; F2 uint8 }
type GenS14 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int; F1 uint16 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint16; F1 int16 }
func genFn0(p0 int, p1 uint16) uint {
    var v4 int8 = -128
    var v5 int8 = (v4 - v4)
    var v6 int = (p0 & 9223372036854775807)
    if (v5 != v4) {
        switch ("hello" + ("a" + ("" + ""))) {
        case "":
            if ((1e10 - (1.5 - -2.25)) == -2.25) {
                p1 += 42
            } else {
                for fi7 := 0; fi7 < 13; fi7++ {
                    v4 = (v5 - 98)
                    v5 = (v4 / ((-85) | 1))
                    _ = fi7
                }
                var v8 uint64 = 18446744073709551615
                p0 += p0
                _ = v8
            }
        case "foo bar":
            fmt.Println(v6, v6, v4)
        case "hello":
            v5 = (-(v4))
        default:
            switch ("foo bar" + "a") {
            case "foo bar":
                fmt.Println(v4)
            case "a":
                for fi9 := 0; fi9 < 20; fi9++ {
                    v6 += (^(fi9))
                    switch (("" + ("\n\t" + "hello")) + "foo bar") {
                    case "hello":
                        v10 := GenS1{F0: 1, F1: 9223372036854775807, F2: -1}
                        _ = v10
                    case "":
                        switch "\n\t" {
                        case "":
                            var v11 float64 = 0.1
                            _ = v11
                        case "a":
                            if false {
                                p0 += v6
                            } else {
                                var v12 int64 = 916
                                v13 := GenS1{F0: 2, F1: 42, F2: -1}
                                _ = v12
                                _ = v13
                            }
                        case "hello":
                            if (18446744073709551615 >= 2) {
                                if (((0.0 != 0.1) || true) && false) {
                                    v14 := GenS2{F0: (p0 % (((v6 * fi9)) | 1)), F1: 783, F2: -1}
                                    var v15 uint64 = 0
                                    _ = v14
                                    _ = v15
                                }
                                v16 := []int{}
                                v17 := GenS3{F0: (p1 >> ((401) & 63)), F1: "\n\t", F2: 1}
                                _ = v16
                                _ = v17
                            }
                        }
                    case "\n\t":
                        var v18 uint16 = (^(p1))
                        _ = v18
                    }
                    _ = fi9
                }
            case "hello":
                for fi19 := 0; fi19 < 4; fi19++ {
                    switch fi19 {
                    case -363:
                        fmt.Println(v6, fi19, v4)
                    case 127:
                        v20 := GenS1{F0: 4294967295, F1: 42, F2: 785}
                        _ = v20
                    case 0:
                        var v21 float32 = float32(((3.14159 + (float64((1e10 * 1.5)) / ((1.5 + 3.14159) / 1e10))) * 1e10))
                        _ = v21
                    }
                    _ = fi19
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 string = "foo bar"
    var v3 uint16 = 65535
    fmt.Println(v1, v3, v3)
    if false {
        fmt.Println(v2, v1, v2)
        fmt.Println(v3, v1, v2)
        if true {
            if (!(0.1 >= 1e10)) {
                var v22 uint32 = 1
                var v23 int = 9223372036854775807
                _ = v22
                _ = v23
            }
        } else {
            v24 := []int{775, 2147483647}
            _ = v24
        }
    } else {
        v3 -= v3
    }
    if (!false) {
        v2 = v2
    } else {
        v25 := GenS1{F0: 42, F1: 42, F2: (v1 + -2147483648)}
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                fmt.Println(v1)
                v2 = (v2 + v2)
                v1 = (v1 & (v1 / (((v1 % ((v1) | 1))) | 1)))
                _ = fi27
            }
            if (true && (1 >= 119)) {
                v28 := GenS4{F0: int(v3), F1: (("a" + ("hello" + "")) + (v2 + "")), F2: 42}
                var v29 uint64 = 4294967295
                _ = v28
                _ = v29
            }
            _ = fi26
        }
        if true {
            v3 -= v3
            v30 := []int{984, 4294967295}
            v31 := GenS3{F0: v3, F1: ((v2 + v2) + "hello"), F2: 555}
            _ = v30
            _ = v31
        }
        _ = v25
    }
    v2 = (v2 + (v2 + (v2 + ("foo bar" + "a"))))
    fmt.Println(v3, v1, v2)
    if false {
        v32 := []int{}
        for fi33 := 0; fi33 < 20; fi33++ {
            v34 := GenS5{F0: fi33}
            var v35 bool = true
            v36 := GenS1{F0: uint(v3), F1: 476, F2: v1}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        var v37 int8 = 67
        _ = v32
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        for fi39 := 0; fi39 < 14; fi39++ {
            if (!true) {
                v1 += (v1 % ((-446) | 1))
                if false {
                    switch v2 {
                    case "foo bar":
                        var v40 int = (fi38 << ((1) & 63))
                        _ = v40
                    case "hello":
                        v41 := []int{4294967295}
                        _ = v41
                    case "":
                        var v42 int = -1
                        _ = v42
                    }
                    v43 := GenS3{F0: v3, F1: (v2 + v2), F2: 0}
                    var v44 int = (fi38 ^ (fi38 - fi38))
                    _ = v43
                    _ = v44
                } else {
                    v45 := []int{}
                    if ((false && true) && true) {
                        if true {
                            v45 = append(v45, -873)
                            if true {
                                if false {
                                    switch ("\n\t" + (v2 + "\n\t")) {
                                    case "":
                                        if true {
                                            if len(v45) >= 2 {
                                                v45 = v45[1:len(v45)]
                                            }
                                            fmt.Println(fi39, v2)
                                            var v46 int32 = 512
                                            _ = v46
                                        }
                                    case "hello":
                                        switch "a" {
                                        case "foo bar":
                                            var v47 uint16 = (v3 >> ((1) & 63))
                                            switch "hello" {
                                            case "\n\t":
                                                if (true || true) {
                                                    if len(v45) > 0 {
                                                        v45[0] = (-(fi39))
                                                    }
                                                    if len(v45) >= 2 {
                                                        v45 = v45[1:len(v45)]
                                                    }
                                                    if len(v45) >= 2 {
                                                        v45 = v45[1:len(v45)]
                                                    }
                                                } else {
                                                    var v48 int32 = (v1 - (v1 << ((genFn0(32768, 255)) & 63)))
                                                    if false {
                                                        v1 += v1
                                                    }
                                                    var v49 uint8 = 128
                                                    _ = len(v45)
                                                    fmt.Println(v3, v1)
                                                    _ = v48
                                                    _ = v49
                                                }
                                                switch 42 {
                                                case 32768:
                                                    v50 := GenS6{F0: 4294967295}
                                                    _ = v50
                                                default:
                                                    var v51 string = ((v2 + v2) + v2)
                                                    _ = v51
                                                }
                                            }
                                            var v52 uint8 = 0
                                            _ = v47
                                            _ = v52
                                        default:
                                            var v53 int = (fi38 + fi39)
                                            if (434 == 831) {
                                                if len(v45) >= 2 {
                                                    v45 = v45[1:len(v45)]
                                                }
                                            } else {
                                                var v54 bool = false
                                                fmt.Println(v2)
                                                _ = v54
                                            }
                                            var v55 int16 = -1
                                            if ((v1 % (((v1 << ((65535) & 63))) | 1)) > (v1 % (((v1 % ((v1) | 1))) | 1))) {
                                                v56 := GenS5{F0: v53}
                                                var v57 float64 = 0.1
                                                _ = v56
                                                _ = v57
                                            }
                                            v1 += 471
                                            _ = v53
                                            _ = v55
                                        }
                                    case "a":
                                        switch "" {
                                        case "foo bar":
                                            if ("\n\t" != "\n\t") {
                                                if false {
                                                    if ("foo bar" > (((v2 + "a") + v2) + ("\n\t" + ""))) {
                                                        v2 = ("" + ("hello" + "hello"))
                                                    } else {
                                                        v2 = v2
                                                        if (2 == 65535) {
                                                            fmt.Println(v1)
                                                            var v58 int = 2147483647
                                                            var v59 uint32 = 65535
                                                            if (v1 >= (v1 - v1)) {
                                                                _ = len(v45)
                                                                var v60 uint8 = 44
                                                                switch -2147483648 {
                                                                case -9223372036854775808:
                                                                    if (-2.25 >= (-2.25 + ((float64(3.14159) + 0.1) / (3.14159 / 1.5)))) {
                                                                        var v61 uint8 = v60
                                                                        v59 -= (v59 ^ v59)
                                                                        switch fi39 {
                                                                        case -596:
                                                                            if len(v45) > 0 {
                                                                                v45[0] = fi39
                                                                            }
                                                                        case 128:
                                                                            fmt.Println(v58)
                                                                        default:
                                                                            fmt.Println(v3, v2)
                                                                        }
                                                                        _ = v61
                                                                    } else {
                                                                        var v62 float32 = (3.14159 / 1e10)
                                                                        _ = v62
                                                                    }
                                                                case -152:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        switch fi38 {
                                                                        case 32767:
                                                                            v2 = v2
                                                                        case 4294967295:
                                                                            switch (v2 + v2) {
                                                                            case "foo bar":
                                                                                v63 := GenS5{F0: (^(fi39))}
                                                                                _ = v63
                                                                            case "\n\t":
                                                                                var v64 bool = (v2 > v2)
                                                                                _ = v64
                                                                            default:
                                                                                v65 := GenS7{F0: 2}
                                                                                _ = v65
                                                                            }
                                                                        default:
                                                                            var v66 int64 = -1
                                                                            _ = v66
                                                                        }
                                                                    case "hello":
                                                                        var v67 float64 = 3.14159
                                                                        _ = len(v45)
                                                                        _ = v67
                                                                    default:
                                                                        var v68 int = (fi39 << ((42) & 63))
                                                                        _ = v68
                                                                    }
                                                                default:
                                                                    var v69 float32 = 0.1
                                                                    _ = v69
                                                                }
                                                                _ = v60
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v59)
                                                                    if true {
                                                                        v70 := GenS8{F0: true}
                                                                        if ((!(true || true)) || (uint64(v3) == 933)) {
                                                                            switch ("foo bar" + v2) {
                                                                            case "":
                                                                                var v71 int16 = 0
                                                                                _ = v71
                                                                            case "\n\t":
                                                                                var v72 int8 = 42
                                                                                _ = v72
                                                                            default:
                                                                                v73 := GenS6{F0: v59}
                                                                                _ = v73
                                                                            }
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                fmt.Println(fi39, v59)
                                                                                switch v2 {
                                                                                case "":
                                                                                    var v74 uint16 = 65535
                                                                                    _ = v74
                                                                                case "\n\t":
                                                                                    if (false || true) {
                                                                                        if len(v45) >= 2 {
                                                                                            v45 = v45[1:len(v45)]
                                                                                        }
                                                                                        var v75 uint8 = 0
                                                                                        v76 := GenS7{F0: 18446744073709551615}
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    }
                                                                                }
                                                                            default:
                                                                                switch v2 {
                                                                                case "":
                                                                                    v45 = append(v45, -1)
                                                                                    var v77 uint = 42
                                                                                    if len(v45) >= 2 {
                                                                                        v45 = v45[1:len(v45)]
                                                                                    }
                                                                                    _ = v77
                                                                                }
                                                                            }
                                                                            if len(v45) >= 2 {
                                                                                v45 = v45[1:len(v45)]
                                                                            }
                                                                        }
                                                                        v78 := GenS9{F0: (v3 - 56)}
                                                                        _ = v70
                                                                        _ = v78
                                                                    } else {
                                                                        v3 += (v3 | (v3 - 65535))
                                                                        fmt.Println(v3)
                                                                        var v79 uint16 = (v3 - (v3 << ((572) & 63)))
                                                                        var v80 uint16 = 167
                                                                        _ = len(v45)
                                                                        _ = v79
                                                                        _ = v80
                                                                    }
                                                                    var v81 string = (v2 + v2)
                                                                    _ = v81
                                                                } else {
                                                                    fmt.Println(v1, v58)
                                                                }
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            v1 = (v1 * (v1 ^ v1))
                                                            _ = len(v45)
                                                        }
                                                    }
                                                    var v82 int8 = -128
                                                    _ = v82
                                                }
                                                v83 := GenS10{F0: (fi38 / ((fi38) | 1)), F1: 65535, F2: fi38}
                                                _ = v83
                                            } else {
                                                var v84 float64 = (-2.25 - (float64(float32(float64(-2.25))) - 0.1))
                                                if true {
                                                    v84 = 1.5
                                                }
                                                _ = v84
                                            }
                                        }
                                    }
                                    var v85 float32 = 1e10
                                    _ = v85
                                }
                                if (!(!((v2 + "\n\t") != ""))) {
                                    v2 = v2
                                    var v86 int32 = -1
                                    switch "a" {
                                    case "foo bar":
                                        v3 -= (v3 + 42)
                                    case "hello":
                                        if false {
                                            _ = len(v45)
                                        } else {
                                            var v87 float32 = -2.25
                                            v2 = "\n\t"
                                            _ = v87
                                        }
                                    default:
                                        if false {
                                            v88 := GenS2{F0: fi39, F1: 1, F2: 32767}
                                            switch (^(fi39)) {
                                            case 127:
                                                if false {
                                                    if true {
                                                        var v89 float64 = 0.1
                                                        var v90 uint8 = 127
                                                        _ = v89
                                                        _ = v90
                                                    } else {
                                                        v86 += (v86 & 2147483647)
                                                        if len(v45) >= 2 {
                                                            v45 = v45[1:len(v45)]
                                                        }
                                                        if false {
                                                            v3 += 42
                                                        }
                                                    }
                                                } else {
                                                    if (v86 < (v1 / ((0) | 1))) {
                                                        if true {
                                                            v45 = append(v45, 599)
                                                            fmt.Println(v2, fi38, v2)
                                                        }
                                                        v2 = "hello"
                                                        fmt.Println(v86, fi38, v2)
                                                    } else {
                                                        var v91 uint64 = 4294967295
                                                        fmt.Println(v1)
                                                        fmt.Println(v91)
                                                        _ = v91
                                                    }
                                                    _ = len(v45)
                                                }
                                            default:
                                                var v92 uint16 = 255
                                                _ = v92
                                            }
                                            _ = v88
                                        }
                                    }
                                    _ = v86
                                } else {
                                    var v93 int16 = -26
                                    v2 = v2
                                    v2 = ""
                                    _ = v93
                                }
                                if false {
                                    v94 := GenS3{F0: 854, F1: (v2 + ((("a" + "foo bar") + (("foo bar" + "") + v2)) + v2)), F2: -9223372036854775808}
                                    _ = v94
                                } else {
                                    fmt.Println(v1)
                                    var v95 uint32 = 65535
                                    _ = v95
                                }
                            } else {
                                v96 := GenS11{F0: 42, F1: 122}
                                var v97 int64 = -9223372036854775808
                                _ = v96
                                _ = v97
                            }
                            v2 = v2
                            switch 2147483647 {
                            case -2147483648:
                                v98 := GenS12{F0: (^(v3)), F1: 59, F2: 1}
                                _ = v98
                            case 42:
                                if true {
                                    v3 += (v3 / ((255) | 1))
                                    fmt.Println(fi38, v2)
                                } else {
                                    v99 := GenS6{F0: 2147483648}
                                    v1 += v1
                                    v3 += (v3 >> ((2) & 63))
                                    switch "hello" {
                                    case "hello":
                                        v2 = ("hello" + v2)
                                    case "a":
                                        switch fi39 {
                                        case -694:
                                            var v100 float32 = 1.5
                                            _ = v100
                                        case 32767:
                                            if (false && true) {
                                                fmt.Println(v1, fi38)
                                            } else {
                                                if (1.5 == (((float32(0.0) * 1.5) * float32(0.0)) / (((0.1 / 3.14159) - 3.14159) / (-2.25 + 1e10)))) {
                                                    v1 = (v1 ^ v1)
                                                }
                                                v1 = (-(v1))
                                                v101 := GenS11{F0: -36, F1: -1}
                                                _ = v101
                                            }
                                        case -980:
                                            v102 := GenS10{F0: (fi39 + fi38), F1: 42, F2: fi38}
                                            _ = v102
                                        }
                                    default:
                                        v1 += v1
                                        var v103 uint = 255
                                        v3 += (v3 >> (((^(v103))) & 63))
                                        _ = v103
                                    }
                                    _ = v99
                                }
                            case 950:
                                switch v2 {
                                case "hello":
                                    if len(v45) > 0 {
                                        v45[0] = fi39
                                    }
                                case "a":
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                }
                                switch (v2 + "foo bar") {
                                case "\n\t":
                                    var v104 string = v2
                                    var v105 uint16 = (v3 >> ((909) & 63))
                                    _ = v104
                                    _ = v105
                                case "a":
                                    if false {
                                        var v106 int32 = -1
                                        if ((("foo bar" + (v2 + v2)) + ("a" + v2)) <= "a") {
                                            switch -9223372036854775808 {
                                            case 0:
                                                var v107 int8 = 74
                                                _ = v107
                                            }
                                        } else {
                                            v3 = v3
                                        }
                                        _ = v106
                                    }
                                case "":
                                    v108 := GenS13{F0: 42, F1: 0, F2: 0}
                                    _ = v108
                                }
                                v3 += (v3 % (((v3 + (v3 / (((v3 % ((v3) | 1))) | 1)))) | 1))
                                v109 := GenS14{F0: -32768, F1: 192, F2: 0}
                                _ = v109
                            }
                            var v110 int32 = (-(v1))
                            _ = v110
                        }
                        var v111 bool = false
                        _ = v111
                    } else {
                        var v112 int8 = 1
                        if (!false) {
                            if true {
                                switch (fi39 & fi38) {
                                case 694:
                                    switch "\n\t" {
                                    case "hello":
                                        var v113 uint32 = 42
                                        _ = v113
                                    default:
                                        v45 = append(v45, int(v3))
                                    }
                                case -9223372036854775808:
                                    var v114 int32 = (-(v1))
                                    _ = v114
                                case -772:
                                    _ = len(v45)
                                default:
                                    var v115 uint = 2
                                    _ = v115
                                }
                                switch fi38 {
                                case 650:
                                    v45 = append(v45, (fi38 % ((fi38) | 1)))
                                case -2147483648:
                                    v116 := GenS7{F0: 893}
                                    _ = v116
                                }
                                switch v2 {
                                case "\n\t":
                                    if (true || true) {
                                        var v117 string = v2
                                        _ = v117
                                    }
                                case "a":
                                    var v118 uint16 = (v3 / ((0) | 1))
                                    _ = v118
                                case "foo bar":
                                    if false {
                                        if false {
                                            if len(v45) > 0 {
                                                v45[0] = (-(fi39))
                                            }
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            var v119 uint = 210
                                            _ = v119
                                        case "\n\t":
                                            var v120 uint32 = 2147483648
                                            _ = v120
                                        case "":
                                            v121 := GenS6{F0: 592}
                                            _ = v121
                                        }
                                        v122 := GenS15{F0: 142}
                                        _ = v122
                                    } else {
                                        switch v2 {
                                        case "foo bar":
                                            fmt.Println(v2, v2, v3)
                                        case "\n\t":
                                            var v123 int = int(v112)
                                            _ = v123
                                        case "a":
                                            switch (fi39 - fi39) {
                                            case 296:
                                                if true {
                                                    v124 := GenS7{F0: 0}
                                                    v45 = append(v45, fi38)
                                                    if false {
                                                        if len(v45) >= 2 {
                                                            v45 = v45[1:len(v45)]
                                                        }
                                                        v112 += (v112 << ((genFn0((fi38 & fi39), v3)) & 63))
                                                        if len(v45) >= 2 {
                                                            v45 = v45[1:len(v45)]
                                                        }
                                                    } else {
                                                        if len(v45) > 0 {
                                                            v45[0] = fi38
                                                        }
                                                        v45 = append(v45, (fi39 % (((fi39 + (fi39 >> ((213) & 63)))) | 1)))
                                                    }
                                                    var v125 float32 = 0.0
                                                    _ = v124
                                                    _ = v125
                                                }
                                            case 32767:
                                                v2 = v2
                                            case -2147483648:
                                                var v126 bool = (18446744073709551615 <= 805)
                                                fmt.Println(fi38)
                                                _ = v126
                                            }
                                        }
                                        _ = len(v45)
                                    }
                                default:
                                    if (true || false) {
                                        var v127 uint8 = 42
                                        _ = v127
                                    }
                                }
                            }
                            switch fi38 {
                            case 32767:
                                var v128 float64 = ((float64(0.0) / 1.0) - (1.5 * float64((1e10 / 1e10))))
                                _ = v128
                            case -1:
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                            case -504:
                                var v129 float64 = (1.5 * 0.0)
                                _ = v129
                            default:
                                switch (v2 + "") {
                                case "\n\t":
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                case "a":
                                    if (!(fi39 <= -1)) {
                                        v1 = (v1 % ((v1) | 1))
                                        v1 = (^(v1))
                                        fmt.Println(v3)
                                        if (825 != -1) {
                                            var v130 int64 = 0
                                            _ = v130
                                        }
                                        v131 := GenS5{F0: (fi39 | 580)}
                                        v3 = (v3 << ((0) & 63))
                                        var v132 uint32 = 2147483648
                                        if ((0.0 * 0.1) > 0.0) {
                                            if (v2 == v2) {
                                                v1 = (v1 + v1)
                                                fmt.Println(fi39)
                                            }
                                            if (true || (v132 > (v132 % ((2147483647) | 1)))) {
                                                var v133 uint64 = 18446744073709551615
                                                v134 := GenS14{F0: int16(v132), F1: v132, F2: (v132 % (((v132 ^ v132)) | 1))}
                                                _ = v133
                                                _ = v134
                                            }
                                            var v135 float32 = (0.1 / (float32(0.0) + float32(float64((1e10 / 1.0)))))
                                            _ = v135
                                        }
                                        v1 += v1
                                        _ = v131
                                        _ = v132
                                    }
                                case "":
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                    if true {
                                        v112 = 100
                                        v45 = append(v45, (-(fi38)))
                                        var v136 int32 = -1
                                        _ = v136
                                    } else {
                                        v137 := GenS3{F0: v3, F1: v2, F2: -9223372036854775808}
                                        _ = v137
                                    }
                                default:
                                    v1 = (v1 | v1)
                                }
                            }
                        } else {
                            _ = len(v45)
                        }
                        v1 += v1
                        _ = v112
                    }
                    if true {
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                        var v138 int32 = v1
                        _ = v138
                    }
                    _ = v45
                }
                if true {
                    fmt.Println(v3)
                }
                var v139 uint8 = 2
                switch fi39 {
                case -2147483648:
                    var v140 uint32 = 2
                    _ = v140
                default:
                    v139 = 169
                    v3 += v3
                    fmt.Println(v139, v139, v1)
                }
                _ = v139
            } else {
                switch (v2 + "\n\t") {
                case "a":
                    fmt.Println(fi39)
                case "\n\t":
                    switch v2 {
                    case "foo bar":
                        v141 := []int{}
                        _ = v141
                    case "":
                        switch fi38 {
                        case 4294967295:
                            fmt.Println(fi38, fi39)
                        case -294:
                            v142 := GenS6{F0: 1}
                            _ = v142
                        }
                    case "a":
                        v3 = v3
                        if false {
                            var v143 int = fi38
                            v144 := GenS16{F0: -912, F1: (v3 * 42)}
                            switch (fi38 & 676) {
                            case -9223372036854775808:
                                v145 := GenS17{F0: 70}
                                _ = v145
                            case 868:
                                v2 = v2
                            case -877:
                                v146 := GenS8{F0: true}
                                _ = v146
                            default:
                                if (!(true && true)) {
                                    v147 := GenS10{F0: (v143 >> ((18446744073709551615) & 63)), F1: 2147483647, F2: fi39}
                                    switch (^(fi38)) {
                                    case 2147483647:
                                        var v148 float64 = (0.1 / 1.5)
                                        _ = v148
                                    case -477:
                                        v143 += (fi38 % ((fi38) | 1))
                                    case 32767:
                                        var v149 int8 = 42
                                        _ = v149
                                    }
                                    _ = v147
                                } else {
                                    var v150 float32 = 3.14159
                                    var v151 string = "\n\t"
                                    var v152 uint = 65535
                                    _ = v150
                                    _ = v151
                                    _ = v152
                                }
                            }
                            if false {
                                v153 := GenS18{F0: (v3 >> ((genFn0((fi38 / ((v143) | 1)), (v3 * v3))) & 63)), F1: 0}
                                _ = v153
                            } else {
                                var v154 uint32 = 4294967295
                                _ = v154
                            }
                            var v155 uint = 2
                            var v156 uint32 = 826
                            switch (fi39 + -1) {
                            case 473:
                                var v157 uint16 = v3
                                _ = v157
                            default:
                                v158 := []int{}
                                _ = v158
                            }
                            _ = v143
                            _ = v144
                            _ = v155
                            _ = v156
                        } else {
                            v159 := GenS16{F0: (fi38 + (fi38 >> ((255) & 63))), F1: (v3 + v3)}
                            var v160 int64 = -9223372036854775808
                            switch v2 {
                            case "a":
                                switch (-(fi39)) {
                                case 0:
                                    var v161 bool = true
                                    _ = v161
                                }
                            case "":
                                switch (fi39 * 190) {
                                case 9223372036854775807:
                                    if (12 <= 127) {
                                        v162 := GenS16{F0: fi39, F1: (v3 + 2)}
                                        v163 := []int{}
                                        _ = v162
                                        _ = v163
                                    }
                                case 32767:
                                    switch fi38 {
                                    case 1:
                                        if (3.14159 != 1.5) {
                                            v164 := []int{fi38, fi38, (fi39 | -1)}
                                            _ = v164
                                        } else {
                                            var v165 uint8 = 149
                                            var v166 float64 = 0.1
                                            _ = v165
                                            _ = v166
                                        }
                                    case -9223372036854775808:
                                        v167 := []int{fi38}
                                        _ = v167
                                    }
                                }
                            }
                            _ = v159
                            _ = v160
                        }
                        var v168 uint32 = 42
                        _ = v168
                    default:
                        fmt.Println(fi39, v1)
                    }
                case "":
                    fmt.Println(fi39)
                default:
                    var v169 uint8 = 42
                    _ = v169
                }
                var v170 string = v2
                var v171 int16 = -499
                _ = v170
                _ = v171
            }
            v172 := []int{1, (fi38 % (((-(fi39))) | 1)), fi39, fi39}
            _ = fi39
            _ = v172
        }
        var v173 uint8 = 1
        _ = fi38
        _ = v173
    }
    v1 += -2147483648
    for fi174 := 0; fi174 < 2; fi174++ {
        if (genFn0(fi174, (v3 * (v3 >> ((928) & 63)))) < 65535) {
            if false {
                v175 := GenS14{F0: -516, F1: 2147483647, F2: 731}
                _ = v175
            }
        } else {
            for fi176 := 0; fi176 < 1; fi176++ {
                fmt.Println(fi176, fi174, v1)
                fmt.Println(v2, v1, v3)
                _ = fi176
            }
            if false {
                var v177 uint32 = 2147483648
                var v178 float32 = float32(0.0)
                var v179 uint64 = 0
                _ = v177
                _ = v178
                _ = v179
            }
            var v180 int = (fi174 & 735)
            _ = v180
        }
        _ = fi174
    }
    v3 = v3
    var v181 uint16 = v3
    fmt.Println(v3)
    fmt.Println(v1)
    if false {
        v1 = (v1 - (v1 * v1))
        fmt.Println(v2, v2, v2)
        v182 := []int{4294967295}
        _ = v182
    } else {
        var v183 float64 = (-2.25 * 3.14159)
        _ = v183
    }
    var v184 int64 = 751
    fmt.Println("v1:", v1)
    fmt.Println("v181:", v181)
    fmt.Println("v184:", v184)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v181
    _ = v184
    _ = v2
    _ = v3
}
