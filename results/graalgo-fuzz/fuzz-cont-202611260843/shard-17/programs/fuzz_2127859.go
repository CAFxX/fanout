package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int8; F1 bool; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int8 }
type GenS4 struct { F0 int32; F1 uint8; F2 bool }
type GenS5 struct { F0 uint64; F1 int64; F2 int16 }
type GenS6 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint; F1 int16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64; F1 uint32 }
type GenS14 struct { F0 string; F1 uint64; F2 string }
type GenS15 struct { F0 string; F1 uint16; F2 int32 }
type GenS16 struct { F0 bool; F1 uint; F2 int64 }
type GenS17 struct { F0 uint8; F1 int32; F2 uint }
type GenS18 struct { F0 uint16; F1 bool; F2 int64 }
type GenS19 struct { F0 uint; F1 uint32; F2 int16 }
type GenS20 struct { F0 uint16; F1 bool }
type GenS21 struct { F0 uint64; F1 int; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 218
    var v3 int16 = 42
    var v4 uint32 = 2147483648
    var v5 bool = true
    if (32768 <= 79) {
        var v6 string = ((("hello" + "hello") + "") + "foo bar")
        _ = v6
    } else {
        if ("\n\t" >= "") {
            switch 32768 {
            case 127:
                for fi7 := 0; fi7 < 2; fi7++ {
                    v8 := GenS1{F0: (v3 >> ((1) & 63)), F1: true}
                    var v9 int16 = v3
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case 32767:
                var v10 uint = 1
                _ = v10
            case -2147483648:
                if false {
                    if v5 {
                        if ((v4 & (v4 * (v4 << ((93) & 63)))) != uint32(v2)) {
                            switch "hello" {
                            case "":
                                switch -563 {
                                case 32767:
                                    v5 = v5
                                case 128:
                                    _ = len(v1)
                                case 217:
                                    for fi11 := 0; fi11 < 18; fi11++ {
                                        v3 -= v3
                                        switch (fi11 * (fi11 % ((62) | 1))) {
                                        case -1:
                                            if true {
                                                v5 = v5
                                                _ = len(v1)
                                                delete(v1, ("foo bar" + "a"))
                                            } else {
                                                v12 := GenS1{F0: (v3 & (v3 | (v3 % (((v3 - v3)) | 1)))), F1: v5}
                                                _ = v12
                                            }
                                        case 128:
                                            var v13 int16 = (-(v3))
                                            _ = v13
                                        default:
                                            for fi14 := 0; fi14 < 20; fi14++ {
                                                if (-32768 == v3) {
                                                    v15 := GenS1{F0: v3, F1: v5}
                                                    v16 := GenS2{F0: 1, F1: false, F2: -926}
                                                    v17 := GenS3{F0: 643, F1: -9223372036854775808, F2: -128}
                                                    _ = v15
                                                    _ = v16
                                                    _ = v17
                                                } else {
                                                    v18 := GenS3{F0: 32767, F1: 1, F2: 127}
                                                    if (18446744073709551615 == 1) {
                                                        v19 := []int{(fi11 & 32767)}
                                                        _ = v19
                                                    }
                                                    _ = v18
                                                }
                                                var v20 bool = v5
                                                _ = fi14
                                                _ = v20
                                            }
                                        }
                                        _ = fi11
                                    }
                                }
                            case "hello":
                                var v21 int32 = 170
                                _ = v21
                            default:
                                fmt.Println(v5, v3, v3)
                            }
                            fmt.Println(v2)
                        } else {
                            if v5 {
                                v22 := []int{1, 907, -418, -1}
                                _ = v22
                            } else {
                                v23 := GenS4{F0: -2147483648, F1: 1, F2: (631 >= 42)}
                                if v5 {
                                    if v5 {
                                        for fi24 := 0; fi24 < 10; fi24++ {
                                            var v25 float64 = -2.25
                                            switch int(v2) {
                                            case 4294967295:
                                                for fi26 := 0; fi26 < 7; fi26++ {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v27 uint64 = 925
                                                        _ = v27
                                                    }
                                                    _ = fi26
                                                }
                                            case -992:
                                                v28 := GenS1{F0: (v3 ^ (v3 - 100)), F1: v5}
                                                _ = v28
                                            default:
                                                v29 := GenS4{F0: -1, F1: v2, F2: (true || (fi24 >= 654))}
                                                _ = v29
                                            }
                                            for fi30 := 0; fi30 < 19; fi30++ {
                                                v31 := GenS5{F0: 292, F1: -630, F2: (-(v3))}
                                                _ = fi30
                                                _ = v31
                                            }
                                            _ = fi24
                                            _ = v25
                                        }
                                        var v32 int32 = -1
                                        for fi33 := 0; fi33 < 11; fi33++ {
                                            var v34 uint32 = 42
                                            var v35 float64 = 0.1
                                            _ = fi33
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = v32
                                    } else {
                                        fmt.Println(v3, v2)
                                        fmt.Println(v3, v5)
                                    }
                                }
                                _ = v23
                            }
                            switch "a" {
                            case "foo bar":
                                fmt.Println(v4, v2)
                            case "\n\t":
                                v36 := GenS6{F0: (^(v3)), F1: 42, F2: 293}
                                _ = v36
                            case "a":
                                v37 := []int{1, 0, -2147483648}
                                _ = v37
                            }
                        }
                        fmt.Println(v4, v4, v3)
                    }
                    for fi38 := 0; fi38 < 1; fi38++ {
                        var v39 uint16 = 42
                        if (-1 <= 42) {
                            var v40 string = "\n\t"
                            _ = v40
                        } else {
                            v2 += 127
                            v41 := GenS5{F0: 166, F1: -643, F2: v3}
                            if (42 != -1) {
                                if ((v3 << ((65535) & 63)) < (v3 + (v3 | (v3 + (v3 % (((v3 & (v3 << ((25) & 63)))) | 1)))))) {
                                    v42 := GenS2{F0: -128, F1: v5, F2: 42}
                                    _ = v42
                                }
                                v4 += v4
                            } else {
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    var v44 int = (^(fi43))
                                    fmt.Println(v2, v39)
                                    _ = fi43
                                    _ = v44
                                }
                                v45 := GenS7{F0: 127}
                                _ = v45
                            }
                            _ = v41
                        }
                        if true {
                            var v46 bool = (false || false)
                            fmt.Println(v3, v4, v5)
                            var v47 int16 = 665
                            _ = v46
                            _ = v47
                        }
                        _ = fi38
                        _ = v39
                    }
                    for fi48 := 0; fi48 < 11; fi48++ {
                        switch 375 {
                        case -2147483648:
                            var v49 uint64 = 727
                            _ = v49
                        case -851:
                            v5 = ((((v2 >> ((429) & 63)) != v2) || v5) && false)
                        }
                        _ = fi48
                    }
                } else {
                    _ = len(v1)
                    var v50 bool = false
                    _ = v50
                }
            default:
                v5 = v5
            }
            fmt.Println(v4, v4, v3)
        } else {
            if v5 {
                v5 = false
            } else {
                if (false || (-2.25 <= (((float32(0.1) - float32((1e10 * 1.5))) + 3.14159) * 1.5))) {
                    if (true || true) {
                        _ = len(v1)
                        for fi51 := 0; fi51 < 19; fi51++ {
                            if v5 {
                                v2 = v2
                            } else {
                                switch (("\n\t" + ((("hello" + "hello") + "") + "foo bar")) + "a") {
                                case "hello":
                                    var v52 string = "\n\t"
                                    _ = v52
                                default:
                                    for fi53 := 0; fi53 < 11; fi53++ {
                                        delete(v1, "foo bar")
                                        _ = fi53
                                    }
                                }
                                v54 := []int{(^(fi51))}
                                v2 += v2
                                _ = v54
                            }
                            if v5 {
                                v4 = (^(v4))
                            }
                            if v5 {
                                v3 -= v3
                                v2 -= v2
                            } else {
                                v2 -= v2
                                for fi55 := 0; fi55 < 2; fi55++ {
                                    switch ("a" + "hello") {
                                    case "":
                                        v56 := GenS8{F0: (fi55 & int(v2)), F1: v2}
                                        _ = v56
                                    case "hello":
                                        var v57 uint32 = (v4 >> ((2) & 63))
                                        _ = v57
                                    case "foo bar":
                                        if (118 <= (v4 << ((18446744073709551615) & 63))) {
                                            fmt.Println(v4, fi55)
                                            v2 = v2
                                            switch (("foo bar" + ("\n\t" + (("foo bar" + "foo bar") + "\n\t"))) + (("a" + "hello") + (("\n\t" + "foo bar") + "hello"))) {
                                            case "hello":
                                                fmt.Println(v3)
                                            default:
                                                var v58 int32 = 2147483647
                                                _ = v58
                                            }
                                        } else {
                                            fmt.Println(fi55, v2, v5)
                                            var v59 bool = (v5 && true)
                                            _ = v59
                                        }
                                    }
                                    fmt.Println(v2, v2, v5)
                                    _ = fi55
                                }
                            }
                            _ = fi51
                        }
                    }
                    if (814 <= uint64(v2)) {
                        var v60 int16 = int16(v2)
                        if v5 {
                            var v61 string = "hello"
                            var v62 uint64 = 2
                            v63 := GenS3{F0: 1, F1: 0, F2: -1}
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        for fi64 := 0; fi64 < 8; fi64++ {
                            if true {
                                var v65 uint16 = 255
                                _ = v65
                            } else {
                                v4 = v4
                                if (v5 && (!v5)) {
                                    switch ("hello" + "foo bar") {
                                    case "foo bar":
                                        v66 := []int{32768, fi64, (fi64 / ((fi64) | 1)), -539}
                                        _ = v66
                                    case "a":
                                        for fi67 := 0; fi67 < 1; fi67++ {
                                            var v68 int = 1
                                            v1["hello"] = "foo bar"
                                            fmt.Println(v3, v68)
                                            _ = fi67
                                            _ = v68
                                        }
                                    }
                                } else {
                                    var v69 uint16 = 1
                                    _ = v69
                                }
                                for fi70 := 0; fi70 < 3; fi70++ {
                                    v71 := v1["foo bar"]
                                    _ = fi70
                                    _ = v71
                                }
                            }
                            v4 -= 4294967295
                            _ = fi64
                        }
                        _ = v60
                    } else {
                        var v72 int8 = 0
                        _ = v72
                    }
                } else {
                    v4 -= 921
                    if v5 {
                        v73 := GenS8{F0: -1, F1: 128}
                        _ = v73
                    } else {
                        switch ("a" + ("" + ("" + "a"))) {
                        case "foo bar":
                            switch "hello" {
                            case "\n\t":
                                switch 41 {
                                case 606:
                                    v2 = (v2 / (((v2 >> ((0) & 63))) | 1))
                                case 127:
                                    if (((((((-2.25 * float64(1e10)) * float64(1.5)) - (1e10 - 0.0)) * float64(1.5)) - (((1e10 * (3.14159 / 1e10)) * 1.5) - 0.0)) * 0.0) >= 0.1) {
                                        v5 = true
                                    }
                                }
                            default:
                                for fi74 := 0; fi74 < 6; fi74++ {
                                    var v75 int16 = (-(v3))
                                    v76 := v1["foo bar"]
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                }
                            }
                        case "a":
                            var v77 string = "foo bar"
                            _ = v77
                        }
                        fmt.Println(v2)
                        v78 := GenS2{F0: 86, F1: true, F2: -2147483648}
                        _ = v78
                    }
                    v5 = v5
                }
            }
        }
        v5 = (true || v5)
    }
    v79 := GenS1{F0: 42, F1: v5}
    for fi80 := 0; fi80 < 3; fi80++ {
        if (!((!false) || false)) {
            var v81 int64 = -1
            _ = v81
        }
        fmt.Println(v4, v4)
        _ = fi80
    }
    v82 := GenS7{F0: v2}
    v2 += (v2 - v2)
    for fi83 := 0; fi83 < 12; fi83++ {
        for fi84 := 0; fi84 < 20; fi84++ {
            if (-89 != 100) {
                v85 := GenS6{F0: -337, F1: 2, F2: (^(v4))}
                var v86 uint8 = (^(v2))
                v87 := v1["foo bar"]
                v88 := GenS9{F0: 1}
                _ = v85
                _ = v86
                _ = v87
                _ = v88
            }
            v2 = (v2 / ((127) | 1))
            if (-9223372036854775808 < 42) {
                var v89 string = "\n\t"
                _ = v89
            }
            fmt.Println(v5)
            v90 := GenS10{F0: 493, F1: 200}
            _ = fi84
            _ = v90
        }
        var v91 float64 = 3.14159
        if (v2 <= 255) {
            v92 := []int{(fi83 | (fi83 % (((-(fi83))) | 1))), (^(fi83)), (fi83 >> ((18446744073709551615) & 63))}
            _ = v92
        }
        _ = fi83
        _ = v91
    }
    if v5 {
        fmt.Println(v3, v4, v2)
    } else {
        v1["\n\t"] = "hello"
        switch "\n\t" {
        case "a":
            v93 := []int{231, 401}
            _ = v93
        case "\n\t":
            var v94 int = 127
            _ = v94
        case "hello":
            if false {
                v95 := []int{-2147483648, 0, -9223372036854775808}
                var v96 int16 = (v3 >> ((18446744073709551615) & 63))
                for fi97 := 0; fi97 < 12; fi97++ {
                    v1[("" + "")] = "hello"
                    for fi98 := 0; fi98 < 10; fi98++ {
                        switch fi97 {
                        case -2147483648:
                            if len(v95) > 0 {
                                v95[0] = int(v4)
                            }
                            if true {
                                v99 := GenS1{F0: 32767, F1: v5}
                                switch fi98 {
                                case -645:
                                    switch -1 {
                                    case 32767:
                                        v4 -= 1
                                    case 971:
                                        v1[("\n\t" + ("\n\t" + "a"))] = "\n\t"
                                    case -9223372036854775808:
                                        v100 := GenS11{F0: 1}
                                        if v5 {
                                            v96 += (-(v96))
                                            var v101 string = ("\n\t" + "foo bar")
                                            switch fi97 {
                                            case 1:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if v5 {
                                                        v95 = append(v95, fi97)
                                                        v96 += -365
                                                        if (true && (!v5)) {
                                                            if v5 {
                                                                if v5 {
                                                                    var v102 uint = 343
                                                                    _ = v102
                                                                } else {
                                                                    delete(v1, "foo bar")
                                                                    v103 := GenS5{F0: 620, F1: 0, F2: (^(v96))}
                                                                    _ = v103
                                                                }
                                                                if true {
                                                                    v101 = v101
                                                                    if (-68 < 67) {
                                                                        v101 = v101
                                                                        var v104 string = v101
                                                                        v1["hello"] = v101
                                                                        switch fi97 {
                                                                        case 485:
                                                                            fmt.Println(v101, v2)
                                                                        case 128:
                                                                            var v105 float32 = (0.1 - (float32(0.1) + 0.1))
                                                                            _ = v105
                                                                        case 42:
                                                                            if len(v95) >= 2 {
                                                                                v95 = v95[1:len(v95)]
                                                                            }
                                                                        default:
                                                                            if v5 {
                                                                                if (0 != 42) {
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    if false {
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    var v106 int32 = 1
                                                                                    if (42 == 934) {
                                                                                        fmt.Println(v2)
                                                                                        if v5 {
                                                                                            v107 := GenS6{F0: int16(v4), F1: 18446744073709551615, F2: (v4 % ((v4) | 1))}
                                                                                            if ((v2 > v2) && (0.1 != 3.14159)) {
                                                                                                v108 := GenS12{F0: 85}
                                                                                                var v109 float32 = -2.25
                                                                                                v110 := GenS4{F0: int32(v4), F1: v2, F2: (-902 > (fi98 - 4294967295))}
                                                                                                _ = v108
                                                                                                _ = v109
                                                                                                _ = v110
                                                                                            }
                                                                                            _ = v107
                                                                                        } else {
                                                                                            var v111 uint = 0
                                                                                            _ = v111
                                                                                        }
                                                                                    } else {
                                                                                        v112 := GenS9{F0: -128}
                                                                                        var v113 int16 = -82
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                    }
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v104
                                                                    } else {
                                                                        var v114 int32 = 42
                                                                        var v115 int32 = v114
                                                                        var v116 uint64 = 2
                                                                        var v117 int32 = v114
                                                                        _ = v114
                                                                        _ = v115
                                                                        _ = v116
                                                                        _ = v117
                                                                    }
                                                                    v118 := GenS1{F0: v96, F1: (!true)}
                                                                    _ = v118
                                                                }
                                                            }
                                                            if true {
                                                                switch 9223372036854775807 {
                                                                case -23:
                                                                    v101 = v101
                                                                    switch fi98 {
                                                                    case -9223372036854775808:
                                                                        v119 := GenS6{F0: (v96 & (-(v3))), F1: 617, F2: 2}
                                                                        _ = v119
                                                                    default:
                                                                        delete(v1, v101)
                                                                    }
                                                                case -1:
                                                                    switch v101 {
                                                                    case "foo bar":
                                                                        fmt.Println(v4, v5)
                                                                    case "a":
                                                                        delete(v1, "")
                                                                    case "\n\t":
                                                                        var v120 int = fi97
                                                                        _ = v120
                                                                    default:
                                                                        v3 += v96
                                                                    }
                                                                    v5 = (v5 || v5)
                                                                    v121 := GenS13{F0: 0, F1: (v4 / ((v4) | 1))}
                                                                    _ = v121
                                                                default:
                                                                    var v122 uint = 307
                                                                    _ = v122
                                                                }
                                                            }
                                                            switch (fi98 | fi97) {
                                                            case -9223372036854775808:
                                                                var v123 float64 = 1e10
                                                                _ = v123
                                                            case 32767:
                                                                switch int(v4) {
                                                                case 643:
                                                                    v2 = v2
                                                                case -1:
                                                                    if v5 {
                                                                        v124 := GenS14{F0: "", F1: 420, F2: (v101 + "foo bar")}
                                                                        var v125 uint32 = v4
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                default:
                                                                    fmt.Println(v4, fi97, v4)
                                                                }
                                                            case 128:
                                                                v126 := GenS5{F0: 74, F1: -9223372036854775808, F2: (v3 & (v96 / (((v96 << ((18446744073709551615) & 63))) | 1)))}
                                                                _ = v126
                                                            default:
                                                                var v127 int = (fi97 >> ((65535) & 63))
                                                                _ = v127
                                                            }
                                                            fmt.Println(v5, v2)
                                                        }
                                                    }
                                                case "\n\t":
                                                    v96 = v3
                                                case "":
                                                    fmt.Println(v3, v2)
                                                }
                                            case 32767:
                                                var v128 int16 = v3
                                                _ = v128
                                            case 2147483647:
                                                fmt.Println(v2)
                                            default:
                                                if ((v96 / (((v96 >> ((4294967295) & 63))) | 1)) <= v3) {
                                                    _ = len(v1)
                                                    var v129 uint64 = uint64(v2)
                                                    _ = v129
                                                } else {
                                                    switch fi98 {
                                                    case 32768:
                                                        var v130 uint8 = v2
                                                        _ = v130
                                                    default:
                                                        var v131 int64 = -9223372036854775808
                                                        _ = v131
                                                    }
                                                }
                                            }
                                            v132 := GenS10{F0: 65535, F1: 805}
                                            if len(v95) >= 2 {
                                                v95 = v95[1:len(v95)]
                                            }
                                            fmt.Println(v3)
                                            switch "a" {
                                            case "\n\t":
                                                v2 = v2
                                                var v133 float32 = 0.0
                                                v101 = v101
                                                _ = v133
                                            default:
                                                var v134 uint = 1
                                                _ = v134
                                            }
                                            _ = v101
                                            _ = v132
                                        } else {
                                            if true {
                                                var v135 int = -799
                                                _ = v135
                                            } else {
                                                switch (fi97 ^ 127) {
                                                case 742:
                                                    var v136 int = (fi98 | fi98)
                                                    if v5 {
                                                        _ = len(v95)
                                                        var v137 int = (v136 & v136)
                                                        _ = v137
                                                    }
                                                    _ = v136
                                                case 440:
                                                    var v138 string = ("foo bar" + "")
                                                    _ = v138
                                                case 1:
                                                    if len(v95) > 0 {
                                                        v95[0] = fi98
                                                    }
                                                    if ((!v5) && (((1 < 1) || v5) || ((v96 % ((-32768) | 1)) < (^(v3))))) {
                                                        v139 := GenS7{F0: (^(v2))}
                                                        var v140 int16 = 1000
                                                        fmt.Println(v5, v4)
                                                        var v141 uint64 = 0
                                                        var v142 string = ("" + "a")
                                                        _ = v139
                                                        _ = v140
                                                        _ = v141
                                                        _ = v142
                                                    } else {
                                                        var v143 int32 = 1
                                                        v143 -= 42
                                                        _ = v143
                                                    }
                                                default:
                                                    v3 = (v96 | (v96 * (v3 / ((1000) | 1))))
                                                }
                                                v144 := GenS3{F0: 2, F1: 836, F2: 127}
                                                if (0.0 < 3.14159) {
                                                    if (!(v5 || (!((0 < -1) && v5)))) {
                                                        switch fi98 {
                                                        case 9223372036854775807:
                                                            if len(v95) >= 2 {
                                                                v95 = v95[1:len(v95)]
                                                            }
                                                        default:
                                                            fmt.Println(v2)
                                                        }
                                                        if (v5 && v5) {
                                                            var v145 uint32 = 42
                                                            v2 += (v2 % ((1) | 1))
                                                            _ = v145
                                                        }
                                                    } else {
                                                        fmt.Println(fi97, v2, v5)
                                                    }
                                                    v4 = v4
                                                    v96 += (v3 + v3)
                                                } else {
                                                    v146 := GenS13{F0: -9223372036854775808, F1: v4}
                                                    v5 = (!(((1.5 + 1.5) - 1e10) != float32((1.5 / 1.0))))
                                                    v2 += (v2 >> ((442) & 63))
                                                    _ = v146
                                                }
                                                switch (((("foo bar" + "foo bar") + ("\n\t" + "hello")) + "\n\t") + "\n\t") {
                                                case "hello":
                                                    fmt.Println(v96, v3, fi98)
                                                case "a":
                                                    var v147 int64 = 61
                                                    _ = v147
                                                case "foo bar":
                                                    switch fi98 {
                                                    case 32768:
                                                        v148 := GenS15{F0: ((((("foo bar" + "") + "hello") + (("foo bar" + "") + "foo bar")) + ("\n\t" + "")) + "a"), F1: 613, F2: -409}
                                                        _ = v148
                                                    case 0:
                                                        fmt.Println(v4, fi97)
                                                    default:
                                                        if (42 > 2147483647) {
                                                            switch ("hello" + "foo bar") {
                                                            case "a":
                                                                switch ("a" + (("\n\t" + "") + "\n\t")) {
                                                                case "":
                                                                    _ = len(v95)
                                                                case "\n\t":
                                                                    if v5 {
                                                                        v5 = false
                                                                        var v149 bool = v5
                                                                        _ = v149
                                                                    } else {
                                                                        v150 := GenS5{F0: uint64(v4), F1: 1, F2: (v3 | v96)}
                                                                        var v151 int8 = 127
                                                                        _ = v150
                                                                        _ = v151
                                                                    }
                                                                case "foo bar":
                                                                    switch (("\n\t" + "foo bar") + "a") {
                                                                    case "\n\t":
                                                                        var v152 uint = 367
                                                                        _ = v152
                                                                    }
                                                                }
                                                            default:
                                                                switch (fi97 % (((fi98 * (fi98 - (fi97 % ((128) | 1))))) | 1)) {
                                                                case 152:
                                                                    var v153 string = "foo bar"
                                                                    _ = v153
                                                                case -533:
                                                                    switch fi98 {
                                                                    case -638:
                                                                        if v5 {
                                                                            v2 += v2
                                                                            v2 = (v2 * uint8(v4))
                                                                            switch (fi97 * 1) {
                                                                            case 32767:
                                                                                v3 = 32767
                                                                            case -779:
                                                                                v3 = (v96 % ((-900) | 1))
                                                                            case 9223372036854775807:
                                                                                var v154 int64 = 1
                                                                                _ = v154
                                                                            default:
                                                                                switch int(v4) {
                                                                                case 128:
                                                                                    _ = len(v95)
                                                                                case 2147483647:
                                                                                    v155 := GenS7{F0: v2}
                                                                                    _ = v155
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if v5 {
                                                                                v4 -= v4
                                                                            }
                                                                            delete(v1, "")
                                                                            if v5 {
                                                                                fmt.Println(fi97, v4)
                                                                                v5 = ((fi97 < (fi98 - (fi97 * 42))) || (v4 < (v4 << ((255) & 63))))
                                                                                if (v5 && (127 < 1)) {
                                                                                    v96 = (v96 >> ((42) & 63))
                                                                                    v3 += v96
                                                                                    v96 -= 1
                                                                                    v2 = 87
                                                                                    v3 -= (v3 | v96)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v4 -= v4
                                                        } else {
                                                            if (42 >= -399) {
                                                                v156 := GenS2{F0: -1, F1: v5, F2: -784}
                                                                switch fi97 {
                                                                case 127:
                                                                    var v157 uint16 = 283
                                                                    _ = v157
                                                                case -2147483648:
                                                                    v2 = 1
                                                                case 1:
                                                                    v158 := GenS16{F0: v5, F1: 1, F2: 504}
                                                                    _ = v158
                                                                default:
                                                                    fmt.Println(v96)
                                                                }
                                                                var v159 int8 = 42
                                                                _ = v156
                                                                _ = v159
                                                            } else {
                                                                var v160 string = "foo bar"
                                                                v161 := GenS17{F0: v2, F1: 1, F2: 2}
                                                                if (232 != 9223372036854775807) {
                                                                    v96 = (v96 >> ((1) & 63))
                                                                    fmt.Println(fi98)
                                                                }
                                                                _ = v160
                                                                _ = v161
                                                            }
                                                        }
                                                    }
                                                    var v162 int32 = 294
                                                    v1[("hello" + ("hello" + "foo bar"))] = "a"
                                                    _ = v162
                                                default:
                                                    v96 = v3
                                                }
                                                _ = v144
                                            }
                                            if v5 {
                                                v163 := GenS7{F0: v2}
                                                _ = v163
                                            } else {
                                                v164 := GenS15{F0: ("" + ""), F1: 952, F2: 1}
                                                fmt.Println(v2, fi97, v96)
                                                v5 = (!(!((float32(-2.25) > (1.5 * ((float32(0.1) / (((1.5 * ((((-2.25 - (0.0 - float32((1.5 - (1.5 * 3.14159))))) + float32(1e10)) / 1.5) + (((3.14159 * -2.25) - 3.14159) - (1.5 + 0.1)))) * (-2.25 - float32((0.0 - -2.25)))) / 1e10)) + 1.5))) || (!false))))
                                                fmt.Println(v4)
                                                v96 = v96
                                                _ = v164
                                            }
                                            v95 = append(v95, fi98)
                                        }
                                        _ = v100
                                    }
                                case 32768:
                                    if v5 {
                                        v4 += (^(v4))
                                    }
                                case 1:
                                    var v165 float64 = 3.14159
                                    _ = v165
                                }
                                if ("foo bar" == "a") {
                                    v96 = (v96 ^ 42)
                                }
                                v166 := GenS12{F0: 18446744073709551615}
                                var v167 uint32 = 1
                                _ = v166
                                _ = v167
                                _ = v99
                            } else {
                                if len(v95) > 0 {
                                    v95[0] = 128
                                }
                                v1["hello"] = "hello"
                            }
                        case 893:
                            v5 = (((350 < 18446744073709551615) && (128 == v2)) || true)
                        default:
                            if v5 {
                                v168 := GenS18{F0: 32767, F1: v5, F2: -9223372036854775808}
                                v1[""] = ("a" + "hello")
                                _ = v168
                            } else {
                                var v169 uint = 618
                                v3 -= (v3 | int16(fi98))
                                switch ("hello" + "a") {
                                case "foo bar":
                                    if (!v5) {
                                        v170 := GenS19{F0: 572, F1: (v4 - v4), F2: 1000}
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch "" {
                                            case "hello":
                                                switch "" {
                                                case "foo bar":
                                                    v169 -= (v169 * (v169 % (((v169 & (v169 & v169))) | 1)))
                                                default:
                                                    var v171 uint = (v169 & (^(v169)))
                                                    _ = v171
                                                }
                                            case "foo bar":
                                                switch "foo bar" {
                                                case "a":
                                                    var v172 bool = v5
                                                    _ = v172
                                                case "hello":
                                                    v2 = (v2 - v2)
                                                }
                                                var v173 int32 = 42
                                                _ = len(v1)
                                                _ = v173
                                            }
                                        case "":
                                            v174 := GenS20{F0: 42, F1: ((v96 ^ (v3 >> (((v169 + 1)) & 63))) >= (v96 + (v3 - v96)))}
                                            _ = v174
                                        case "a":
                                            if false {
                                                if false {
                                                    if (-9223372036854775808 < 1) {
                                                        v175 := GenS21{F0: 733, F1: int(v2), F2: 9223372036854775807}
                                                        v176 := GenS13{F0: 42, F1: 2147483648}
                                                        _ = v175
                                                        _ = v176
                                                    }
                                                    var v177 int16 = (v3 + -244)
                                                    _ = v177
                                                } else {
                                                    var v178 float64 = 0.1
                                                    v169 = (v169 - (v169 | (v169 + 65535)))
                                                    v178 = ((v178 - ((v178 + v178) + 3.14159)) / v178)
                                                    delete(v1, "foo bar")
                                                    switch fi97 {
                                                    case 0:
                                                        v179 := GenS15{F0: "", F1: 232, F2: -102}
                                                        _ = v179
                                                    case -463:
                                                        v4 += (v4 << ((202) & 63))
                                                    }
                                                    _ = v178
                                                }
                                                fmt.Println(v4, v3)
                                            } else {
                                                var v180 int8 = 101
                                                fmt.Println(v180, v169, v3)
                                                _ = v180
                                            }
                                        }
                                        _ = v170
                                    } else {
                                        if false {
                                            var v181 uint32 = v4
                                            v182 := v1["a"]
                                            if (v5 && v5) {
                                                var v183 uint8 = v2
                                                _ = v183
                                            }
                                            _ = v181
                                            _ = v182
                                        } else {
                                            v95 = append(v95, (fi98 * (^(fi98))))
                                        }
                                        var v184 uint32 = v4
                                        v185 := GenS18{F0: 255, F1: false, F2: 452}
                                        _ = v184
                                        _ = v185
                                    }
                                }
                                _ = v169
                            }
                        }
                        _ = fi98
                    }
                    v3 += v96
                    _ = fi97
                }
                _ = v95
                _ = v96
            } else {
                var v186 int64 = 42
                v4 += 65535
                var v187 uint = 390
                _ = v186
                _ = v187
            }
        default:
            v188 := GenS10{F0: 4294967295, F1: (v3 & -32768)}
            _ = v188
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v82
}
