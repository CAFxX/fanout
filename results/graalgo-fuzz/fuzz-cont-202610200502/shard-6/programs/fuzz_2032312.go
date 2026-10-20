package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
type GenS5 struct { F0 int32; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint32; F1 bool; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 string; F2 bool }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint; F1 int }
type GenS13 struct { F0 int64; F1 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int32; F1 int32 }
type GenS18 struct { F0 uint8; F1 uint16 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint16; F1 uint64; F2 bool }
type GenS21 struct { F0 int8; F1 string }
type GenS22 struct { F0 int64; F1 uint16 }
type GenS23 struct { F0 bool; F1 string }
type GenS24 struct { F0 int; F1 int16; F2 int16 }
type GenS25 struct { F0 int8; F1 uint64 }
type GenS26 struct { F0 uint32; F1 uint16 }
func genFn0() uint {
    v4 := []int{-9223372036854775808, -9223372036854775808, 0}
    if (!false) {
        for fi5 := 0; fi5 < 9; fi5++ {
            if (!(!false)) {
                var v6 int = 32767
                if true {
                    v6 = (-(v6))
                    v7 := GenS1{F0: 0, F1: 1, F2: 74}
                    _ = v7
                } else {
                    v6 += 32768
                    if (42 != 18446744073709551615) {
                        for fi8 := 0; fi8 < 19; fi8++ {
                            if false {
                                if true {
                                    v6 -= (fi5 % ((128) | 1))
                                    switch ((("a" + "a") + ("\n\t" + "")) + ("hello" + "a")) {
                                    case "foo bar":
                                        var v9 uint32 = 0
                                        _ = v9
                                    case "a":
                                        v6 += (v6 & v6)
                                    case "hello":
                                        var v10 int16 = 408
                                        _ = v10
                                    default:
                                        switch (("foo bar" + "\n\t") + "foo bar") {
                                        case "\n\t":
                                            v11 := GenS1{F0: 1000, F1: 1, F2: 4294967295}
                                            _ = v11
                                        case "a":
                                            if true {
                                                fmt.Println(v6, fi8, fi8)
                                            }
                                        case "":
                                            v6 -= v6
                                        }
                                    }
                                } else {
                                    v12 := GenS1{F0: 32767, F1: 11, F2: 65535}
                                    _ = v12
                                }
                                fmt.Println(fi8, fi8)
                            }
                            v6 = (fi5 / (((-(v6))) | 1))
                            v4 = append(v4, (fi5 ^ fi8))
                            _ = fi8
                        }
                        if true {
                            v13 := GenS2{F0: 100}
                            var v14 int64 = -1
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v13
                            _ = v14
                        } else {
                            v6 = (fi5 | -1)
                        }
                    }
                }
                _ = v6
            }
            var v15 int64 = 95
            _ = fi5
            _ = v15
        }
        v16 := GenS3{F0: "\n\t", F1: 1}
        _ = v16
    }
    var v17 int32 = 1
    _ = v17
    _ = v4
    return 0
}
func genFn1(p0 uint64, p1 int8, p2 int8) uint {
    if false {
        var v18 uint8 = 42
        if (0 >= -1) {
            p2 = (p1 >> ((65535) & 63))
            v19 := []int{127, 42}
            _ = v19
        } else {
            v20 := GenS4{F0: -1, F1: (^(p1)), F2: v18}
            _ = v20
        }
        _ = v18
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        v21 := []int{4294967295}
        _ = v21
    case "a":
        v22 := GenS3{F0: "\n\t", F1: int32(p1)}
        _ = v22
    default:
        v23 := []int{-9223372036854775808, 128, 1, 1}
        _ = v23
    }
    var v24 uint64 = p0
    p2 = p2
    var v25 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 42
    var v2 float32 = ((-2.25 * 1.5) - float32(0.0))
    v3 := make(map[int]string)
    if (210 <= 128) {
        v1 -= (v1 % ((0) | 1))
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(v2, fi26)
            _ = fi26
        }
        if false {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint32 = (v1 + (v1 / ((v1) | 1)))
                v29 := GenS5{F0: 0, F1: 1, F2: 2}
                v2 += v2
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 uint64 = 18446744073709551615
            v2 += 3.14159
            _ = v30
        } else {
            v2 += (v2 + (3.14159 / (0.1 * v2)))
        }
    } else {
        if (877 >= (v1 - (v1 >> ((205) & 63)))) {
            switch 1 {
            case 127:
                v1 -= v1
            case -2147483648:
                v2 += (0.1 + float32((float64(v2) / 3.14159)))
            default:
                v2 = ((v2 * v2) * -2.25)
            }
        }
        fmt.Println(v2, v2)
        v31 := GenS6{F0: (v1 - (v1 & v1)), F1: false, F2: "\n\t"}
        _ = v31
    }
    v32 := []int{1, 128}
    if (false || false) {
        v2 = (float32(1e10) * 1e10)
        if len(v32) > 0 {
            v32[0] = -2147483648
        }
        if (!(0 < 2147483647)) {
            for fi33 := 0; fi33 < 2; fi33++ {
                v1 -= 2147483648
                switch "foo bar" {
                case "":
                    for fi34 := 0; fi34 < 4; fi34++ {
                        if (279 >= 32767) {
                            if false {
                                v2 = ((-2.25 + 0.1) * ((-2.25 - 0.0) / -2.25))
                                _ = len(v32)
                            }
                            var v35 uint16 = 65535
                            var v36 int64 = 1
                            var v37 int64 = -1
                            v38 := v3[fi33]
                            v2 += (float32((-2.25 - -2.25)) + (v2 - v2))
                            switch v38 {
                            case "a":
                                var v39 string = v38
                                _ = v39
                            }
                            fmt.Println(fi34, v37, fi34)
                            var v40 float64 = 3.14159
                            v35 -= v35
                            if false {
                                switch fi33 {
                                case 2147483647:
                                    var v41 string = ("" + ("\n\t" + "\n\t"))
                                    _ = v41
                                case 4294967295:
                                    var v42 int16 = 1000
                                    _ = v42
                                default:
                                    if (664 < 42) {
                                        var v43 bool = false
                                        v1 = (v1 & v1)
                                        _ = v43
                                    } else {
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                    }
                                }
                                var v44 uint16 = 0
                                var v45 int = fi33
                                _ = v44
                                _ = v45
                            } else {
                                var v46 uint16 = (v35 * (v35 % ((v35) | 1)))
                                var v47 float32 = (((((float32(v40) - 0.1) - (v2 / -2.25)) - ((float32(3.14159) / v2) - (v2 / v2))) - -2.25) - (v2 + ((((v2 + v2) / (0.0 * (v2 / v2))) * v2) / ((v2 * v2) * v2))))
                                _ = v46
                                _ = v47
                            }
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v40
                        }
                        v2 -= 1.5
                        _ = fi34
                    }
                case "foo bar":
                    var v48 uint32 = (v1 % ((v1) | 1))
                    _ = v48
                case "hello":
                    v2 = (-2.25 - v2)
                }
                delete(v3, -94)
                _ = fi33
            }
            if true {
                for fi49 := 0; fi49 < 10; fi49++ {
                    if false {
                        if false {
                            fmt.Println(v2, v1)
                            v3[fi49] = ""
                        }
                        var v50 uint16 = 0
                        _ = v50
                    }
                    v1 -= (^(v1))
                    for fi51 := 0; fi51 < 5; fi51++ {
                        v1 = v1
                        v52 := GenS2{F0: 127}
                        fmt.Println(v2)
                        _ = fi51
                        _ = v52
                    }
                    _ = fi49
                }
                v2 = (v2 * -2.25)
                var v53 uint64 = 0
                _ = v53
            }
            switch ((("" + "hello") + ("\n\t" + "\n\t")) + "\n\t") {
            case "hello":
                var v54 uint32 = v1
                _ = v54
            }
        }
    } else {
        for fi55 := 0; fi55 < 20; fi55++ {
            v1 -= (v1 ^ (^(v1)))
            if true {
                var v56 int8 = 49
                _ = v56
            } else {
                switch (fi55 ^ fi55) {
                case 42:
                    v2 = v2
                case 127:
                    v1 += (v1 - v1)
                case 630:
                    if (float32((float64(1.5) - 1e10)) >= ((v2 / float32(((0.1 + (1.5 * (0.1 * (3.14159 + ((((0.1 + 1e10) * ((float64(v2) - 0.1) / ((float64(1.5) + 0.0) * 3.14159))) + (1.5 * 3.14159)) + 1.5))))) / 1e10))) - v2)) {
                        v1 += 2147483648
                    }
                default:
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                }
            }
            _ = fi55
        }
        if (-69 > -128) {
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            if (1 <= 1) {
                var v57 uint16 = 784
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                for fi58 := 0; fi58 < 13; fi58++ {
                    var v59 int16 = 42
                    v60 := v3[fi58]
                    v2 = 0.0
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                _ = v57
            } else {
                if (0 <= 2) {
                    v61 := GenS2{F0: 100}
                    fmt.Println(v1, v2)
                    _ = v61
                }
                for fi62 := 0; fi62 < 20; fi62++ {
                    fmt.Println(v1)
                    _ = fi62
                }
            }
            if false {
                if (!(((1e10 + 3.14159) == -2.25) || true)) {
                    var v63 uint8 = 1
                    v64 := GenS7{F0: 100}
                    if (!true) {
                        v65 := GenS3{F0: ("foo bar" + ("a" + ("a" + "\n\t"))), F1: 1}
                        v66 := GenS4{F0: 1000, F1: 0, F2: v63}
                        var v67 string = "hello"
                        _ = v65
                        _ = v66
                        _ = v67
                    } else {
                        var v68 float64 = 1.5
                        _ = v68
                    }
                    _ = v63
                    _ = v64
                }
                switch 4294967295 {
                case 0:
                    v69 := GenS2{F0: -128}
                    _ = v69
                case 32767:
                    var v70 uint32 = v1
                    _ = v70
                default:
                    switch "\n\t" {
                    case "\n\t":
                        var v71 int16 = -1
                        _ = v71
                    case "foo bar":
                        v72 := GenS3{F0: "\n\t", F1: 0}
                        _ = v72
                    default:
                        v2 += v2
                    }
                }
                for fi73 := 0; fi73 < 7; fi73++ {
                    _ = len(v3)
                    var v74 string = (("hello" + ("hello" + "hello")) + "\n\t")
                    switch v74 {
                    case "":
                        v2 = (v2 + float32(float64(v2)))
                    default:
                        switch 4294967295 {
                        case 1:
                            v75 := GenS8{F0: ((((v74 + v74) + (v74 + v74)) + v74) + "hello")}
                            _ = v75
                        case 9223372036854775807:
                            v76 := GenS3{F0: v74, F1: int32(v1)}
                            _ = v76
                        case 0:
                            if ((-2.25 / 1.5) != 1e10) {
                                for fi77 := 0; fi77 < 17; fi77++ {
                                    v2 = v2
                                    var v78 int32 = 532
                                    _ = fi77
                                    _ = v78
                                }
                                for fi79 := 0; fi79 < 10; fi79++ {
                                    if (false && (96 <= 127)) {
                                        var v80 uint = 42
                                        switch "a" {
                                        case "foo bar":
                                            var v81 int = fi73
                                            _ = v81
                                        case "":
                                            if true {
                                                fmt.Println(fi73, fi73)
                                                v1 += 0
                                            } else {
                                                fmt.Println(fi79, v2)
                                                fmt.Println(v2, v2)
                                                if true {
                                                    fmt.Println(v1, v1)
                                                    v80 += (v80 | v80)
                                                    var v82 int8 = -7
                                                    _ = v82
                                                } else {
                                                    switch "" {
                                                    case "hello":
                                                        if (179 >= 138) {
                                                            if false {
                                                                var v83 bool = false
                                                                _ = v83
                                                            }
                                                            switch -2147483648 {
                                                            case 32768:
                                                                if (2147483647 != 42) {
                                                                    var v84 string = ("\n\t" + ("" + ("a" + "foo bar")))
                                                                    if true {
                                                                        fmt.Println(fi73, fi73, v84)
                                                                        v1 = v1
                                                                        if (fi73 < fi79) {
                                                                            v84 = v84
                                                                            _ = len(v32)
                                                                        }
                                                                    }
                                                                    if true {
                                                                        v85 := GenS9{F0: -32768, F1: v84, F2: (32767 == 520)}
                                                                        v86 := GenS10{F0: 281, F1: 1}
                                                                        switch 449 {
                                                                        case 32767:
                                                                            if (1 <= 0) {
                                                                                if (!false) {
                                                                                    if false {
                                                                                        v87 := GenS11{F0: -32768}
                                                                                        var v88 string = v84
                                                                                        v89 := GenS12{F0: genFn1(18446744073709551615, 0, 114), F1: fi79}
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                    }
                                                                                    switch (v74 + v84) {
                                                                                    case "foo bar":
                                                                                        var v90 uint8 = 74
                                                                                        _ = v90
                                                                                    default:
                                                                                        var v91 bool = (false && (((v2 * v2) < (v2 / 1.5)) && ((32767 > fi79) && (false || true))))
                                                                                        v80 -= v80
                                                                                        _ = v91
                                                                                    }
                                                                                } else {
                                                                                    switch ("hello" + (v84 + v84)) {
                                                                                    case "":
                                                                                        if (1.5 == 3.14159) {
                                                                                            var v92 float32 = v2
                                                                                            _ = v92
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            switch v74 {
                                                                                            case "":
                                                                                                switch v74 {
                                                                                                case "hello":
                                                                                                    fmt.Println(v84, fi73, fi73)
                                                                                                case "":
                                                                                                    var v93 uint64 = uint64(v80)
                                                                                                    var v94 uint64 = 2
                                                                                                    _ = v93
                                                                                                    _ = v94
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v2, fi73, fi73)
                                                                                                }
                                                                                            case "\n\t":
                                                                                                if true {
                                                                                                    v95 := GenS13{F0: -1, F1: 127}
                                                                                                    _ = v95
                                                                                                }
                                                                                            case "foo bar":
                                                                                                switch ((v84 + "\n\t") + (v84 + v84)) {
                                                                                                case "hello":
                                                                                                    v96 := GenS4{F0: 1, F1: 42, F2: 2}
                                                                                                    _ = v96
                                                                                                default:
                                                                                                    if (42 > (^(v80))) {
                                                                                                        if ((v80 + (v80 & v80)) == uint(v1)) {
                                                                                                            v97 := GenS14{F0: v84}
                                                                                                            delete(v3, (fi79 % ((fi73) | 1)))
                                                                                                            fmt.Println(v74)
                                                                                                            _ = v97
                                                                                                        }
                                                                                                        _ = len(v32)
                                                                                                    } else {
                                                                                                        var v98 int64 = -1
                                                                                                        _ = v98
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    fmt.Println(fi73, v74, v80)
                                                                                                    var v99 int32 = 0
                                                                                                    _ = v99
                                                                                                } else {
                                                                                                    var v100 float32 = 1.5
                                                                                                    v74 = ""
                                                                                                    _ = v100
                                                                                                }
                                                                                            }
                                                                                            if true {
                                                                                                switch "" {
                                                                                                case "":
                                                                                                    v3[-2147483648] = ""
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v2, v1, v74)
                                                                                                }
                                                                                                var v101 uint64 = 0
                                                                                                _ = v101
                                                                                            } else {
                                                                                                var v102 uint16 = 561
                                                                                                if false {
                                                                                                    v1 += v1
                                                                                                }
                                                                                                var v103 int = fi73
                                                                                                _ = v102
                                                                                                _ = v103
                                                                                            }
                                                                                            switch (v84 + (((v74 + ("hello" + v74)) + v74) + v84)) {
                                                                                            case "\n\t":
                                                                                                v3[fi73] = ""
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch fi73 {
                                                                                        case 128:
                                                                                            v104 := GenS9{F0: 504, F1: v84, F2: (219 >= 127)}
                                                                                            _ = v104
                                                                                        case 123:
                                                                                            if (v2 <= v2) {
                                                                                                fmt.Println(v74, v84, v80)
                                                                                                v105 := GenS8{F0: "\n\t"}
                                                                                                var v106 float32 = float32((((((1.5 / ((3.14159 + ((((((0.0 * 0.0) + 0.0) * 1.5) * 0.1) - -2.25) * (1.5 / 1e10))) * 1e10)) - -2.25) + 1.5) - (-2.25 * (1e10 - 3.14159))) - ((1e10 - float64(float32(1.5))) - (((float64(0.1) - (1e10 / 1.0)) + 3.14159) + (float64(v2) / (float64((v2 + (1e10 * v2))) / 3.14159))))))
                                                                                                if len(v32) >= 2 {
                                                                                                    v32 = v32[1:len(v32)]
                                                                                                }
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            } else {
                                                                                                fmt.Println(v1, v1, v80)
                                                                                            }
                                                                                        case 32768:
                                                                                            if true {
                                                                                                var v107 uint32 = v1
                                                                                                switch (fi79 - (fi79 | (fi73 | (fi73 + fi79)))) {
                                                                                                case -2147483648:
                                                                                                    var v108 uint = genFn1(42, 0, -128)
                                                                                                    _ = v108
                                                                                                case -223:
                                                                                                    _ = len(v32)
                                                                                                default:
                                                                                                    v80 -= (v80 & v80)
                                                                                                }
                                                                                                var v109 int16 = -32768
                                                                                                var v110 uint32 = (v107 & (v107 % ((v107) | 1)))
                                                                                                _ = v107
                                                                                                _ = v109
                                                                                                _ = v110
                                                                                            } else {
                                                                                                if (!(2 < (v1 << ((genFn0()) & 63)))) {
                                                                                                    var v111 int32 = 2147483647
                                                                                                    _ = v111
                                                                                                } else {
                                                                                                    fmt.Println(v84, fi73)
                                                                                                    var v112 bool = true
                                                                                                    var v113 uint32 = v1
                                                                                                    _ = v112
                                                                                                    _ = v113
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v114 int8 = 0
                                                                                    _ = v114
                                                                                }
                                                                                v115 := GenS3{F0: ("foo bar" + "hello"), F1: 2147483647}
                                                                                v80 -= v80
                                                                                var v116 int = fi79
                                                                                fmt.Println(fi73, fi73, fi73)
                                                                                v117 := GenS15{F0: 1}
                                                                                _ = v115
                                                                                _ = v116
                                                                                _ = v117
                                                                            } else {
                                                                                fmt.Println(v1)
                                                                                _ = len(v3)
                                                                                if false {
                                                                                    if true {
                                                                                        if true {
                                                                                            fmt.Println(v80, v2, v2)
                                                                                            v118 := v3[30]
                                                                                            _ = v118
                                                                                        } else {
                                                                                            _ = len(v32)
                                                                                            switch 127 {
                                                                                            case -9223372036854775808:
                                                                                                var v119 uint32 = (v1 % ((1) | 1))
                                                                                                _ = v119
                                                                                            case 42:
                                                                                                v120 := GenS16{F0: 117}
                                                                                                _ = v120
                                                                                            case -681:
                                                                                                fmt.Println(v74)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v121 int16 = -533
                                                                                    v122 := GenS11{F0: -1}
                                                                                    _ = v121
                                                                                    _ = v122
                                                                                }
                                                                            }
                                                                        case 32768:
                                                                            if true {
                                                                                var v123 float32 = -2.25
                                                                                _ = v123
                                                                            }
                                                                        case -728:
                                                                            v74 = ""
                                                                        }
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    _ = v84
                                                                } else {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        var v124 uint64 = 524
                                                                        _ = v124
                                                                    case "foo bar":
                                                                        if (0 == 704) {
                                                                            if false {
                                                                                v1 = v1
                                                                            }
                                                                            var v125 string = v74
                                                                            v126 := GenS12{F0: v80, F1: fi73}
                                                                            _ = v125
                                                                            _ = v126
                                                                        } else {
                                                                            _ = len(v32)
                                                                            v74 = v74
                                                                            v127 := GenS8{F0: v74}
                                                                            _ = v127
                                                                        }
                                                                    case "\n\t":
                                                                        switch (fi79 - fi73) {
                                                                        case 2147483647:
                                                                            v128 := GenS10{F0: v1, F1: 32767}
                                                                            _ = v128
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                if len(v32) > 0 {
                                                                    v32[0] = fi73
                                                                }
                                                            }
                                                        } else {
                                                            v80 = 18446744073709551615
                                                        }
                                                    case "a":
                                                        v129 := GenS17{F0: int32(fi79), F1: 42}
                                                        _ = v129
                                                    }
                                                    v1 = (^(v1))
                                                    v1 = 2147483648
                                                }
                                            }
                                        case "a":
                                            var v130 float64 = 0.1
                                            _ = v130
                                        }
                                        var v131 uint64 = 42
                                        _ = v131
                                        _ = v80
                                    }
                                    _ = fi79
                                }
                                if ((0 < -1) || true) {
                                    switch v74 {
                                    case "foo bar":
                                        v2 = 0.1
                                    case "a":
                                        var v132 uint64 = 1
                                        _ = v132
                                    default:
                                        v74 = "a"
                                    }
                                    var v133 float32 = float32(float64(v2))
                                    switch "a" {
                                    case "\n\t":
                                        for fi134 := 0; fi134 < 16; fi134++ {
                                            if (!true) {
                                                v133 -= v133
                                                v3[(fi134 << ((9) & 63))] = ((((v74 + ((v74 + "foo bar") + v74)) + v74) + "foo bar") + (((v74 + "foo bar") + "hello") + v74))
                                                var v135 uint8 = 1
                                                _ = v135
                                            } else {
                                                _ = len(v32)
                                                v74 = v74
                                                var v136 bool = false
                                                _ = v136
                                            }
                                            v137 := GenS3{F0: "foo bar", F1: -760}
                                            fmt.Println(v133, fi73)
                                            _ = fi134
                                            _ = v137
                                        }
                                    case "foo bar":
                                        v138 := GenS12{F0: 980, F1: 4294967295}
                                        _ = v138
                                    case "a":
                                        if true {
                                            v3[1] = v74
                                            for fi139 := 0; fi139 < 6; fi139++ {
                                                v74 = v74
                                                _ = fi139
                                            }
                                            for fi140 := 0; fi140 < 3; fi140++ {
                                                if false {
                                                    var v141 float32 = 0.1
                                                    _ = v141
                                                }
                                                var v142 uint = 1
                                                _ = fi140
                                                _ = v142
                                            }
                                        } else {
                                            v32 = append(v32, fi73)
                                        }
                                    default:
                                        fmt.Println(fi73, v74)
                                    }
                                    _ = v133
                                } else {
                                    for fi143 := 0; fi143 < 7; fi143++ {
                                        _ = len(v32)
                                        v144 := GenS18{F0: 128, F1: 255}
                                        v74 = v74
                                        var v145 float32 = 0.0
                                        var v146 int32 = -649
                                        if ((1.5 + ((0.1 + 1e10) / (3.14159 / ((-2.25 * -2.25) + 0.1)))) > (0.0 * 1e10)) {
                                            v74 = v74
                                        }
                                        v3[(fi143 & 378)] = ((((("\n\t" + v74) + v74) + (v74 + v74)) + v74) + v74)
                                        fmt.Println(v74)
                                        var v147 uint16 = 65535
                                        v148 := GenS19{F0: -44}
                                        v149 := GenS20{F0: 65535, F1: uint64(v147), F2: true}
                                        if false {
                                            switch 0 {
                                            case -1:
                                                v2 = float32(float64(0.1))
                                            case 4294967295:
                                                v150 := GenS21{F0: 42, F1: v74}
                                                _ = v150
                                            default:
                                                fmt.Println(v2, fi143)
                                            }
                                            if ((fi73 < 0) && (!(66 <= 1))) {
                                                if (-737 >= 597) {
                                                    var v151 int16 = -771
                                                    v152 := GenS13{F0: 224, F1: 193}
                                                    var v153 string = ""
                                                    _ = v151
                                                    _ = v152
                                                    _ = v153
                                                } else {
                                                    v146 = (v146 | (v146 ^ v146))
                                                }
                                                v154 := GenS22{F0: -9223372036854775808, F1: (v147 & (^(v147)))}
                                                _ = v154
                                            } else {
                                                if (0.1 == 1.5) {
                                                    v155 := GenS8{F0: v74}
                                                    switch (fi143 >> ((803) & 63)) {
                                                    case -1:
                                                        v145 += v2
                                                    }
                                                    var v156 int16 = -447
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v157 int = fi73
                                                        var v158 int8 = 42
                                                        _ = v157
                                                        _ = v158
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                    _ = v155
                                                    _ = v156
                                                }
                                                v159 := GenS8{F0: v74}
                                                v1 = 2147483648
                                                v160 := GenS17{F0: -714, F1: v146}
                                                _ = v159
                                                _ = v160
                                            }
                                        }
                                        _ = fi143
                                        _ = v144
                                        _ = v145
                                        _ = v146
                                        _ = v147
                                        _ = v148
                                        _ = v149
                                    }
                                    if true {
                                        switch "" {
                                        case "a":
                                            if false {
                                                v2 = v2
                                                v161 := GenS23{F0: true, F1: ""}
                                                _ = v161
                                            } else {
                                                v74 = v74
                                            }
                                        case "foo bar":
                                            _ = len(v32)
                                        default:
                                            var v162 uint64 = 2
                                            _ = v162
                                        }
                                    } else {
                                        fmt.Println(fi73, v74)
                                    }
                                }
                            } else {
                                v163 := GenS9{F0: 1000, F1: "foo bar", F2: false}
                                _ = v163
                            }
                        }
                    }
                    _ = fi73
                    _ = v74
                }
            } else {
                for fi164 := 0; fi164 < 11; fi164++ {
                    var v165 int = fi164
                    var v166 float64 = 3.14159
                    _ = fi164
                    _ = v165
                    _ = v166
                }
                for fi167 := 0; fi167 < 9; fi167++ {
                    var v168 uint16 = 65535
                    v169 := GenS19{F0: 100}
                    switch "\n\t" {
                    case "a":
                        var v170 uint16 = (^(v168))
                        _ = v170
                    case "\n\t":
                        _ = len(v32)
                    case "hello":
                        var v171 int64 = -9223372036854775808
                        _ = v171
                    default:
                        var v172 uint8 = 255
                        _ = v172
                    }
                    _ = fi167
                    _ = v168
                    _ = v169
                }
            }
        }
        var v173 float32 = (1e10 * v2)
        _ = v173
    }
    v1 -= (v1 & 0)
    var v174 float64 = -2.25
    for fi175 := 0; fi175 < 20; fi175++ {
        if (true && true) {
            v176 := GenS24{F0: (fi175 - (fi175 >> ((569) & 63))), F1: 511, F2: -32768}
            v1 = (v1 - v1)
            var v177 int = fi175
            _ = v176
            _ = v177
        }
        v2 += v2
        _ = fi175
    }
    if true {
        for fi178 := 0; fi178 < 11; fi178++ {
            fmt.Println(v1, fi178, v1)
            var v179 int32 = 1
            for fi180 := 0; fi180 < 17; fi180++ {
                var v181 float64 = (0.1 * v174)
                _ = fi180
                _ = v181
            }
            _ = fi178
            _ = v179
        }
        if false {
            for fi182 := 0; fi182 < 15; fi182++ {
                v2 = v2
                _ = len(v32)
                fmt.Println(fi182, fi182)
                _ = fi182
            }
            var v183 int16 = -333
            _ = v183
        }
        switch 1 {
        case -2147483648:
            v184 := GenS25{F0: 84, F1: 1}
            _ = v184
        }
    } else {
        if (108 == 127) {
            if len(v32) > 0 {
                v32[0] = 824
            }
        }
    }
    fmt.Println(v174)
    fmt.Println(v1, v2)
    switch -625 {
    case 4294967295:
        var v185 float64 = ((v174 / ((1e10 - v174) - v174)) / 0.1)
        _ = v185
    }
    for fi186 := 0; fi186 < 15; fi186++ {
        v1 = (^(v1))
        var v187 int = fi186
        if (true || (65535 <= 2)) {
            switch v187 {
            case 1:
                if true {
                    switch (v187 - (v187 / ((fi186) | 1))) {
                    case 497:
                        for fi188 := 0; fi188 < 5; fi188++ {
                            _ = len(v3)
                            if true {
                                switch (-(fi186)) {
                                case 32767:
                                    v189 := v3[fi188]
                                    _ = v189
                                case 127:
                                    v1 += 42
                                default:
                                    v190 := GenS6{F0: 0, F1: (!false), F2: ""}
                                    _ = v190
                                }
                            } else {
                                v191 := GenS22{F0: -1, F1: 637}
                                v187 += fi188
                                _ = v191
                            }
                            var v192 uint = 1
                            _ = fi188
                            _ = v192
                        }
                    default:
                        v193 := GenS26{F0: (v1 - v1), F1: 655}
                        _ = v193
                    }
                    v2 += v2
                }
            }
        }
        _ = fi186
        _ = v187
    }
    v194 := GenS11{F0: 1000}
    v1 += (^(v1))
    v2 = float32(0.1)
    var v195 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v174:", v174)
    fmt.Println("v195:", v195)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v174
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v32
}
