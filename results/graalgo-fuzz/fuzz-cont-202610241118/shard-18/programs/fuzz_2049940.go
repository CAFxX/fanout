package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 int; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int64; F1 int64; F2 bool }
type GenS7 struct { F0 int; F1 uint16; F2 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 uint64; F1 uint64 }
type GenS12 struct { F0 int; F1 int64; F2 uint64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int16; F1 uint64 }
func genFn0() int {
    if false {
        var v5 uint8 = 225
        var v6 int8 = 127
        v7 := []int{}
        _ = v5
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            if ((fi8 ^ 127) == fi8) {
                var v9 int64 = -252
                v9 += v9
                _ = v9
            }
            _ = fi8
        }
    }
    switch "a" {
    case "a":
        var v10 uint32 = 2147483648
        _ = v10
    default:
        var v11 int16 = -1
        _ = v11
    }
    return 42
}
func genFn1(p0 uint32, p1 uint16, p2 uint) int16 {
    p2 -= p2
    p2 += 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 float32 = ((-2.25 + float32(-2.25)) / 1e10)
    v2 := []int{-2147483648, -9223372036854775808, -9223372036854775808}
    var v3 int32 = 1
    var v4 int = 127
    if (1 != 1) {
        for fi12 := 0; fi12 < 13; fi12++ {
            v1 -= 3.14159
            _ = fi12
        }
    } else {
        if (!("\n\t" > "hello")) {
            var v13 int32 = -948
            if (!(42 == 2)) {
                if true {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v13 -= v13
                        _ = fi14
                    }
                    v1 = (v1 - 0.1)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    v4 += (-(fi15))
                    v16 := GenS1{F0: 42, F1: -592, F2: 0}
                    if (1e10 > -2.25) {
                        v2 = append(v2, (v4 / (((-(fi15))) | 1)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (!(true && ((v3 < (^(v13))) && (!false)))) {
                            switch "a" {
                            case "\n\t":
                                v3 -= -1
                            case "":
                                v17 := GenS1{F0: 762, F1: 0, F2: -1}
                                _ = v17
                            case "foo bar":
                                switch (v4 << ((0) & 63)) {
                                case 2147483647:
                                    v4 = (-(fi15))
                                case -9223372036854775808:
                                    for fi18 := 0; fi18 < 2; fi18++ {
                                        var v19 int32 = v13
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v20 int32 = v19
                                            _ = v20
                                        }
                                        var v21 float32 = 3.14159
                                        _ = fi18
                                        _ = v19
                                        _ = v21
                                    }
                                }
                            }
                        }
                    } else {
                        var v22 int16 = genFn1(2147483648, 65535, 18446744073709551615)
                        _ = v22
                    }
                    _ = fi15
                    _ = v16
                }
            }
            if (255 == 255) {
                var v23 uint64 = 370
                _ = v23
            } else {
                var v24 int64 = -9223372036854775808
                _ = v24
            }
            _ = v13
        } else {
            var v25 int = -23
            _ = v25
        }
        var v26 int32 = v3
        _ = v26
    }
    switch ("foo bar" + ("" + "a")) {
    case "hello":
        if false {
            for fi27 := 0; fi27 < 16; fi27++ {
                if false {
                    v28 := GenS2{F0: 0, F1: (fi27 + (-(v4))), F2: -32768}
                    for fi29 := 0; fi29 < 9; fi29++ {
                        if (true && (65535 > 255)) {
                            switch "\n\t" {
                            case "a":
                                var v30 bool = (false && (((174 == 127) || true) || true))
                                _ = v30
                            case "":
                                v31 := GenS3{F0: 178}
                                _ = v31
                            }
                            switch 2147483647 {
                            case -9223372036854775808:
                                fmt.Println(v4, v1, v1)
                            case -2147483648:
                                switch "hello" {
                                case "":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "a":
                                    var v32 int8 = -77
                                    _ = v32
                                case "\n\t":
                                    v3 -= v3
                                    var v33 string = "a"
                                    _ = v33
                                }
                            case -264:
                                fmt.Println(fi29, v3)
                            default:
                                var v34 float32 = (3.14159 * 1e10)
                                _ = v34
                            }
                            v4 = genFn0()
                            v3 = 1
                            switch (("hello" + ("\n\t" + "foo bar")) + "foo bar") {
                            case "\n\t":
                                if (!false) {
                                    v35 := GenS1{F0: 32767, F1: 1, F2: 0}
                                    _ = v35
                                } else {
                                    v36 := GenS4{F0: 1, F1: 1}
                                    switch "hello" {
                                    case "":
                                        v37 := GenS2{F0: 496, F1: fi27, F2: -1}
                                        _ = v37
                                    case "hello":
                                        v38 := GenS5{F0: ("" + ("foo bar" + "a")), F1: 467}
                                        _ = v38
                                    default:
                                        v2 = append(v2, 4294967295)
                                    }
                                    v39 := GenS1{F0: 368, F1: int16(v4), F2: 42}
                                    _ = v36
                                    _ = v39
                                }
                            case "foo bar":
                                v1 -= -2.25
                            }
                            v4 = (fi27 - fi27)
                            var v40 int64 = -9223372036854775808
                            switch ("" + "") {
                            case "hello":
                                fmt.Println(fi27, v40)
                            default:
                                v41 := GenS6{F0: 515, F1: 1, F2: false}
                                _ = v41
                            }
                            _ = v40
                        } else {
                            v4 = v4
                            if false {
                                var v42 string = "a"
                                _ = v42
                            }
                            var v43 uint32 = 2147483647
                            v3 += -553
                            var v44 int16 = 990
                            fmt.Println(v44, v44)
                            _ = v43
                            _ = v44
                        }
                        fmt.Println(v3, v1)
                        _ = fi29
                    }
                    v4 = genFn0()
                    _ = v28
                }
                var v45 int32 = v3
                _ = fi27
                _ = v45
            }
            if (v4 > genFn0()) {
                fmt.Println(v3)
                var v46 float64 = 1e10
                for fi47 := 0; fi47 < 2; fi47++ {
                    fmt.Println(v1)
                    _ = fi47
                }
                _ = v46
            } else {
                fmt.Println(v4, v4)
                v3 -= (v3 | (v3 ^ (v3 << ((4294967295) & 63))))
                v2 = append(v2, v4)
            }
            v3 += v3
        } else {
            _ = len(v2)
        }
    }
    var v48 uint = 255
    var v49 uint = 65535
    switch "a" {
    case "a":
        switch (^(v4)) {
        case 32767:
            v4 = genFn0()
        case -9223372036854775808:
            fmt.Println(v1)
        case 338:
            var v50 string = "\n\t"
            _ = v50
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v51 := GenS6{F0: 1, F1: -9223372036854775808, F2: true}
    for fi52 := 0; fi52 < 14; fi52++ {
        var v53 int16 = 0
        v1 = v1
        _ = fi52
        _ = v53
    }
    var v54 float32 = v1
    if true {
        v4 = genFn0()
        switch genFn0() {
        case 419:
            fmt.Println(v4, v54)
        default:
            v49 -= (v49 & v49)
        }
        v1 += v1
    } else {
        v2 = append(v2, 9223372036854775807)
        if (false && true) {
            var v55 int16 = 1
            if false {
                var v56 bool = true
                if (!false) {
                    if v56 {
                        var v57 uint16 = 1
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if v56 {
                            v58 := GenS7{F0: int(v55), F1: 65535, F2: 533}
                            switch -2147483648 {
                            case 32767:
                                if false {
                                    switch "foo bar" {
                                    case "hello":
                                        if v56 {
                                            _ = len(v2)
                                        } else {
                                            for fi59 := 0; fi59 < 3; fi59++ {
                                                if v56 {
                                                    var v60 uint32 = 584
                                                    var v61 int32 = v3
                                                    _ = v60
                                                    _ = v61
                                                }
                                                for fi62 := 0; fi62 < 11; fi62++ {
                                                    v63 := GenS5{F0: "foo bar", F1: 2}
                                                    var v64 bool = v56
                                                    _ = fi62
                                                    _ = v63
                                                    _ = v64
                                                }
                                                for fi65 := 0; fi65 < 20; fi65++ {
                                                    fmt.Println(v55)
                                                    fmt.Println(v4, v54, v54)
                                                    var v66 uint64 = 2
                                                    fmt.Println(v54, v57, fi65)
                                                    _ = fi65
                                                    _ = v66
                                                }
                                                _ = fi59
                                            }
                                            var v67 int8 = -128
                                            _ = v67
                                        }
                                    default:
                                        if v56 {
                                            v68 := GenS3{F0: (v57 >> ((v49) & 63))}
                                            _ = v68
                                        }
                                    }
                                    v69 := GenS8{F0: "hello"}
                                    _ = v69
                                }
                            }
                            var v70 int64 = -9223372036854775808
                            _ = v58
                            _ = v70
                        } else {
                            if (42 == 18446744073709551615) {
                                v57 -= (^(v57))
                                v71 := GenS9{F0: (v49 - (v49 ^ (v49 - (v49 / ((2) | 1))))), F1: (("foo bar" + ("a" + "")) + "")}
                                for fi72 := 0; fi72 < 4; fi72++ {
                                    v73 := GenS10{F0: v57, F1: -1}
                                    v55 += (v55 | 389)
                                    if v56 {
                                        var v74 uint64 = 969
                                        v75 := GenS11{F0: v74, F1: 885}
                                        fmt.Println(v3, v55, v74)
                                        _ = v74
                                        _ = v75
                                    }
                                    _ = fi72
                                    _ = v73
                                }
                                _ = v71
                            } else {
                                v56 = true
                                switch genFn0() {
                                case 0:
                                    v76 := GenS2{F0: 230, F1: 0, F2: v55}
                                    _ = v76
                                }
                                for fi77 := 0; fi77 < 8; fi77++ {
                                    var v78 int8 = 89
                                    for fi79 := 0; fi79 < 15; fi79++ {
                                        var v80 int64 = -9223372036854775808
                                        switch ("foo bar" + "") {
                                        case "":
                                            v78 += v78
                                        }
                                        _ = fi79
                                        _ = v80
                                    }
                                    _ = fi77
                                    _ = v78
                                }
                            }
                            if v56 {
                                var v81 uint8 = 0
                                _ = v81
                            } else {
                                fmt.Println(v1, v57, v4)
                                fmt.Println(v3, v1, v57)
                            }
                            if v56 {
                                for fi82 := 0; fi82 < 14; fi82++ {
                                    v2 = append(v2, 42)
                                    if (-1 >= -1) {
                                        fmt.Println(v57)
                                        v4 = 919
                                        v49 -= 179
                                    } else {
                                        v2 = append(v2, fi82)
                                        if len(v2) > 0 {
                                            v2[0] = fi82
                                        }
                                    }
                                    _ = len(v2)
                                    _ = fi82
                                }
                            } else {
                                v3 = (v3 % ((v3) | 1))
                            }
                        }
                        _ = v57
                    }
                } else {
                    fmt.Println(v56, v4)
                    v83 := GenS9{F0: v49, F1: (((("foo bar" + "\n\t") + "hello") + "hello") + "foo bar")}
                    if (float64(1e10) == 0.0) {
                        v55 -= (v55 * 32767)
                    } else {
                        fmt.Println(v56, v4, v4)
                        switch "\n\t" {
                        case "foo bar":
                            v2 = append(v2, 9223372036854775807)
                        case "":
                            for fi84 := 0; fi84 < 2; fi84++ {
                                v2 = append(v2, int(v3))
                                if (v48 == (v49 * 65535)) {
                                    if (!(v1 >= 1.5)) {
                                        for fi85 := 0; fi85 < 16; fi85++ {
                                            switch -9223372036854775808 {
                                            case 0:
                                                _ = len(v2)
                                            case 128:
                                                v56 = false
                                            case 32768:
                                                _ = len(v2)
                                            default:
                                                var v86 int8 = -1
                                                _ = v86
                                            }
                                            v87 := GenS5{F0: "hello", F1: 0}
                                            v56 = v56
                                            var v88 int64 = 42
                                            if v56 {
                                                var v89 uint8 = 0
                                                _ = v89
                                            } else {
                                                var v90 uint32 = 2
                                                _ = v90
                                            }
                                            _ = fi85
                                            _ = v87
                                            _ = v88
                                        }
                                        for fi91 := 0; fi91 < 18; fi91++ {
                                            v56 = (v56 && false)
                                            var v92 bool = (!false)
                                            switch 4294967295 {
                                            case -972:
                                                fmt.Println(v4)
                                                var v93 uint8 = 255
                                                switch fi91 {
                                                case 4294967295:
                                                    switch (fi91 / (((v4 / ((int(v55)) | 1))) | 1)) {
                                                    case 128:
                                                        var v94 string = (("" + "a") + "foo bar")
                                                        _ = v94
                                                    case 2147483647:
                                                        v48 = v49
                                                    default:
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi91 / (((v4 % (((^(v4))) | 1))) | 1))
                                                        }
                                                    }
                                                default:
                                                    v95 := GenS12{F0: -1, F1: -9223372036854775808, F2: 485}
                                                    _ = v95
                                                }
                                                _ = v93
                                            }
                                            _ = fi91
                                            _ = v92
                                        }
                                    }
                                    v54 = v54
                                    var v96 int32 = v3
                                    _ = v96
                                } else {
                                    switch (("foo bar" + "hello") + ("foo bar" + "")) {
                                    case "hello":
                                        if false {
                                            v97 := GenS13{F0: -54}
                                            _ = v97
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    default:
                                        for fi98 := 0; fi98 < 16; fi98++ {
                                            switch (v4 ^ (^(fi84))) {
                                            case -734:
                                                if (genFn0() >= v4) {
                                                    fmt.Println(v1, v4, v3)
                                                    var v99 uint32 = 700
                                                    _ = v99
                                                } else {
                                                    var v100 uint32 = 42
                                                    switch "hello" {
                                                    case "":
                                                        v54 = -2.25
                                                    case "\n\t":
                                                        var v101 string = "hello"
                                                        _ = v101
                                                    case "a":
                                                        v102 := GenS2{F0: 42, F1: (fi98 + 1), F2: v55}
                                                        _ = v102
                                                    }
                                                    var v103 int8 = 42
                                                    _ = v100
                                                    _ = v103
                                                }
                                            default:
                                                v104 := GenS14{F0: 1}
                                                _ = v104
                                            }
                                            v56 = (((18446744073709551615 > 418) && true) && v56)
                                            _ = fi98
                                        }
                                    }
                                    if (0 <= 65535) {
                                        switch fi84 {
                                        case 9223372036854775807:
                                            v105 := GenS7{F0: v4, F1: 255, F2: 982}
                                            _ = v105
                                        case 2147483647:
                                            v106 := GenS7{F0: (v4 >> ((v48) & 63)), F1: 65535, F2: 42}
                                            _ = v106
                                        case -797:
                                            v48 -= v48
                                        default:
                                            v3 = 1
                                        }
                                    }
                                }
                                _ = fi84
                            }
                        }
                    }
                    _ = v83
                }
                _ = v56
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v107 uint8 = 0
                _ = v107
            }
            _ = v55
        } else {
            _ = len(v2)
            for fi108 := 0; fi108 < 18; fi108++ {
                var v109 uint16 = 1
                _ = fi108
                _ = v109
            }
            var v110 string = "\n\t"
            _ = v110
        }
        v49 = (v49 ^ v48)
    }
    if false {
        v3 = (v3 * v3)
    } else {
        v1 = (v1 - ((0.0 + ((-2.25 / v1) - (v1 * (-2.25 - ((v1 / 1.5) * v1))))) * v1))
    }
    for fi111 := 0; fi111 < 4; fi111++ {
        var v112 int64 = -9223372036854775808
        _ = fi111
        _ = v112
    }
    var v113 float32 = (v54 + v54)
    v114 := GenS7{F0: v4, F1: 1, F2: 1}
    v54 += 3.14159
    fmt.Println(v3, v3, v1)
    if true {
        for fi115 := 0; fi115 < 9; fi115++ {
            fmt.Println(v48)
            _ = fi115
        }
        for fi116 := 0; fi116 < 1; fi116++ {
            v117 := GenS15{F0: -1, F1: 0}
            for fi118 := 0; fi118 < 11; fi118++ {
                var v119 uint16 = 0
                switch fi118 {
                case -9223372036854775808:
                    var v120 float64 = (0.0 + 1e10)
                    _ = v120
                }
                _ = fi118
                _ = v119
            }
            _ = fi116
            _ = v117
        }
    }
    for fi121 := 0; fi121 < 6; fi121++ {
        for fi122 := 0; fi122 < 16; fi122++ {
            v54 = (v1 * 1e10)
            var v123 int8 = 42
            if true {
                v3 = v3
                switch 184 {
                case 0:
                    _ = len(v2)
                case -324:
                    v1 = (float32(0.0) / float32(3.14159))
                }
            } else {
                var v124 float64 = float64(v54)
                _ = v124
            }
            _ = fi122
            _ = v123
        }
        _ = fi121
    }
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v51
    _ = v54
}
