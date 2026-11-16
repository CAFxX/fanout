package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 int }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint8; F1 uint16; F2 string }
type GenS6 struct { F0 uint64; F1 int16; F2 string }
type GenS7 struct { F0 uint64; F1 bool; F2 int16 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 uint32; F1 int; F2 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 bool }
type GenS12 struct { F0 int64; F1 int32; F2 string }
type GenS13 struct { F0 int16; F1 int; F2 uint8 }
type GenS14 struct { F0 int; F1 int8; F2 uint64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint32; F1 string; F2 uint64 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS19 struct { F0 string }
func genFn0() int8 {
    var v5 uint16 = 65535
    v5 += (v5 ^ (v5 + (^(v5))))
    v6 := GenS1{F0: 957, F1: 2}
    v5 -= v5
    _ = v5
    _ = v6
    return 6
}

func main() {
    var v1 int8 = -63
    v2 := []uint32{1, 65535, 4294967295}
    var v3 int16 = 1
    var v4 int8 = 4
    if (false && (214 > 42)) {
        if len(v2) > 0 {
            v2[0] = 132
        }
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(fi7, v4)
        for fi8 := 0; fi8 < 10; fi8++ {
            v2 = append(v2, 1)
            _ = fi8
        }
        var v9 uint16 = 337
        _ = fi7
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(fi10, fi10)
        _ = fi10
    }
    v4 -= -1
    v11 := GenS1{F0: 2147483648, F1: 236}
    switch "a" {
    case "a":
        if false {
            v2 = append(v2, 1)
            var v12 uint64 = 230
            v3 -= v3
            _ = v12
        } else {
            v13 := GenS2{F0: 42, F1: int32(v1), F2: 42}
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint64 = 0
                for fi16 := 0; fi16 < 10; fi16++ {
                    v1 += v4
                    _ = fi16
                }
                _ = fi14
                _ = v15
            }
            var v17 int64 = int64(v3)
            _ = v13
            _ = v17
        }
    case "hello":
        fmt.Println(v3, v3, v4)
    case "\n\t":
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 int32 = 0
            v20 := GenS3{F0: 99, F1: v3}
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        if ((0.1 * 1e10) <= 1e10) {
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 uint32 = 42
                fmt.Println(v4, v3)
                _ = fi21
                _ = v22
            }
            if false {
                var v23 float64 = 0.0
                _ = v23
            } else {
                v1 += (-(v4))
                v4 += 115
                var v24 float64 = (3.14159 * -2.25)
                _ = v24
            }
            fmt.Println(v4, v3)
        }
    }
    var v25 float32 = 1e10
    v26 := GenS2{F0: 229, F1: -1, F2: 2147483647}
    if true {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 float64 = ((-2.25 / ((0.1 * ((0.0 / ((3.14159 + 0.1) * (3.14159 / (0.0 + -2.25)))) + float64((v25 * float32(1.5))))) + ((1.5 * float64(v25)) * (float64(v25) * 3.14159)))) - float64(v25))
            _ = fi27
            _ = v28
        }
        v25 = 0.1
        v1 = 127
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v29 := GenS2{F0: 197, F1: 1, F2: -9223372036854775808}
    if (false && false) {
        if false {
            v30 := GenS3{F0: 280, F1: (v3 | v3)}
            if len(v2) > 0 {
                v2[0] = 831
            }
            var v31 uint32 = 42
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            switch "\n\t" {
            case "a":
                var v33 uint32 = 437
                _ = v33
            case "foo bar":
                v3 = (v3 & 1)
            }
            if false {
                if false {
                    switch fi32 {
                    case 2147483647:
                        var v34 int8 = v1
                        _ = v34
                    case -940:
                        fmt.Println(v1, fi32, fi32)
                    }
                }
            }
            _ = fi32
        }
        if (((v25 * (v25 / v25)) <= (1e10 / 3.14159)) || ((v25 * v25) == (v25 - v25))) {
            var v35 int8 = (v1 * v4)
            _ = v35
        } else {
            var v36 int64 = -1
            fmt.Println(v25, v1, v36)
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        v3 = -200
        fmt.Println(fi37, v1)
        if len(v2) > 0 {
            v2[0] = 65535
        }
        _ = fi37
    }
    var v38 bool = ((2 <= 255) || true)
    var v39 uint64 = 1
    if v38 {
        fmt.Println(v3, v4, v3)
        _ = len(v2)
    }
    v38 = v38
    for fi40 := 0; fi40 < 6; fi40++ {
        v41 := GenS3{F0: -9223372036854775808, F1: v3}
        switch (^(fi40)) {
        case 32767:
            v3 = (v3 * (^(v3)))
        case 0:
            _ = len(v2)
        default:
            if v38 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi42 := 0; fi42 < 5; fi42++ {
                    var v43 uint16 = 2
                    v44 := GenS2{F0: uint8(v39), F1: 522, F2: fi40}
                    var v45 int8 = v4
                    switch 32767 {
                    case 812:
                        var v46 int64 = -1
                        var v47 uint8 = 70
                        v2 = append(v2, 715)
                        var v48 float64 = (float64((((float32(0.0) / v25) * ((v25 / (v25 / v25)) * (float32(0.0) - (((v25 * float32(((float64(v25) / 1.0) / -2.25))) * v25) / ((float32(1e10) - 0.0) * v25))))) + ((v25 + 0.1) + 0.0))) / (1.5 / 3.14159))
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                    v45 = genFn0()
                    var v49 int8 = (v4 + (v1 * (v4 + (v1 * (-(v1))))))
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v49
                }
            }
        }
        if v38 {
            switch "foo bar" {
            case "":
                v38 = v38
            case "a":
                var v50 uint = 1
                _ = v50
            default:
                if (-705 < 1) {
                    var v51 uint16 = 664
                    _ = v51
                }
            }
            if (v38 || true) {
                if ((((117 != -398) || v38) && (9223372036854775807 <= -572)) || ((!(!(36 <= 1))) || (2147483648 < 960))) {
                    for fi52 := 0; fi52 < 7; fi52++ {
                        switch fi52 {
                        case 0:
                            _ = len(v2)
                        case 128:
                            v53 := GenS4{F0: 42, F1: v38}
                            _ = v53
                        case -788:
                            v4 += (v1 % ((v4) | 1))
                        default:
                            switch -596 {
                            case 4294967295:
                                v54 := GenS5{F0: 127, F1: 2, F2: ("" + "")}
                                _ = v54
                            case 965:
                                var v55 int32 = -2147483648
                                _ = v55
                            case 33:
                                v2 = append(v2, 1)
                            }
                            if v38 {
                                _ = len(v2)
                            } else {
                                v56 := GenS6{F0: 4294967295, F1: (v3 | -343), F2: "a"}
                                _ = v56
                            }
                        }
                        var v57 uint64 = (v39 << ((1) & 63))
                        _ = fi52
                        _ = v57
                    }
                }
                v3 = int16(fi40)
                if (0 > v3) {
                    var v58 uint = 251
                    v25 += v25
                    v38 = (v38 || (!((60 > 159) || v38)))
                    _ = v58
                }
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch ("" + "") {
            case "\n\t":
                if (float32(((((0.1 * (3.14159 * 3.14159)) - -2.25) - 0.1) * 1.5)) != 3.14159) {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    if true {
                        switch "a" {
                        case "foo bar":
                            v59 := GenS6{F0: 42, F1: (v3 >> ((42) & 63)), F2: ""}
                            _ = v59
                        case "":
                            if v38 {
                                for fi60 := 0; fi60 < 1; fi60++ {
                                    var v61 uint8 = 42
                                    if v38 {
                                        var v62 int8 = (v1 >> ((1) & 63))
                                        switch (fi60 | -1) {
                                        case 32767:
                                            if v38 {
                                                v39 = v39
                                                if v38 {
                                                    if (373 <= 2147483647) {
                                                        fmt.Println(v39, v3)
                                                    }
                                                } else {
                                                    if true {
                                                        var v63 int32 = 42
                                                        _ = v63
                                                    }
                                                }
                                            }
                                        case 0:
                                            if ((!(v3 <= (v3 % ((v3) | 1)))) || (!((^(v3)) <= (v3 | v3)))) {
                                                v64 := GenS2{F0: (v61 % (((v61 + (v61 % ((99) | 1)))) | 1)), F1: 1, F2: fi40}
                                                _ = v64
                                            } else {
                                                if (v38 && true) {
                                                    var v65 int8 = v4
                                                    _ = v65
                                                }
                                                v25 = (v25 - 0.1)
                                            }
                                        }
                                        v3 += (-(v3))
                                        _ = v62
                                    }
                                    v1 += v4
                                    _ = fi60
                                    _ = v61
                                }
                                if true {
                                    var v66 uint64 = (v39 % (((v39 * 42)) | 1))
                                    var v67 int8 = genFn0()
                                    _ = v66
                                    _ = v67
                                } else {
                                    for fi68 := 0; fi68 < 10; fi68++ {
                                        var v69 float64 = 1e10
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v70 uint16 = 255
                                        v71 := GenS1{F0: 4294967295, F1: 130}
                                        v1 -= 0
                                        _ = fi68
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = 65535
                                    }
                                    v39 = (v39 + v39)
                                }
                                v2 = append(v2, 1)
                            }
                        default:
                            v3 -= (v3 * 702)
                        }
                        fmt.Println(v38, v3, v1)
                    } else {
                        switch fi40 {
                        case 32767:
                            if (!false) {
                                for fi72 := 0; fi72 < 2; fi72++ {
                                    if true {
                                        v38 = v38
                                        fmt.Println(fi40, fi40, v3)
                                        if v38 {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v38, fi40)
                                            _ = len(v2)
                                        } else {
                                            v1 = (v1 - (^(v1)))
                                        }
                                    }
                                    v73 := GenS7{F0: 18446744073709551615, F1: ((2 != 0) || true), F2: v3}
                                    if (!v38) {
                                        if v38 {
                                            if (true && ("hello" > "a")) {
                                                var v74 float64 = (1.5 * 1.5)
                                                _ = v74
                                            } else {
                                                fmt.Println(fi72, fi72)
                                                v75 := GenS7{F0: v39, F1: v38, F2: 42}
                                                _ = v75
                                            }
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v25 -= float32((3.14159 / -2.25))
                                        }
                                        var v76 uint16 = 42
                                        _ = v76
                                    } else {
                                        v4 -= v1
                                    }
                                    _ = fi72
                                    _ = v73
                                }
                                fmt.Println(v39)
                            }
                        case 42:
                            fmt.Println(v1, v38, fi40)
                        case 1:
                            for fi77 := 0; fi77 < 1; fi77++ {
                                var v78 int64 = 9223372036854775807
                                _ = fi77
                                _ = v78
                            }
                        }
                    }
                    fmt.Println(v4, fi40)
                }
            case "foo bar":
                if ((391 != -9223372036854775808) || v38) {
                    for fi79 := 0; fi79 < 17; fi79++ {
                        v2 = append(v2, 497)
                        v38 = false
                        switch fi40 {
                        case -248:
                            v80 := GenS1{F0: 222, F1: 1}
                            _ = v80
                        case 32768:
                            var v81 float32 = 0.0
                            _ = v81
                        default:
                            fmt.Println(v38, v38, v3)
                            var v82 bool = (4294967295 > 2147483648)
                            _ = v82
                        }
                        _ = fi79
                    }
                    var v83 int8 = (v1 * (v1 % ((-112) | 1)))
                    if v38 {
                        if (!(0.0 > ((v25 * (((float32(1.5) + 3.14159) * float32(3.14159)) * 3.14159)) / v25))) {
                            if v38 {
                                if v38 {
                                    var v84 int64 = 0
                                    v85 := GenS8{F0: 0, F1: 841}
                                    _ = v84
                                    _ = v85
                                } else {
                                    if (((("foo bar" + "") + "\n\t") + "hello") == "a") {
                                        var v86 uint8 = 128
                                        if v38 {
                                            if ((v83 * genFn0()) <= (v4 | int8(fi40))) {
                                                switch "hello" {
                                                case "":
                                                    var v87 bool = (v38 && v38)
                                                    _ = v87
                                                case "\n\t":
                                                    var v88 float32 = (((v25 * v25) / ((-2.25 * 0.1) + v25)) + 0.0)
                                                    _ = v88
                                                case "foo bar":
                                                    switch int(v86) {
                                                    case -1:
                                                        var v89 uint = 18446744073709551615
                                                        _ = v89
                                                    case 4294967295:
                                                        if ((!v38) || false) {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v90 uint32 = 0
                                                                _ = v90
                                                            default:
                                                                var v91 string = ("hello" + "\n\t")
                                                                _ = v91
                                                            }
                                                            for fi92 := 0; fi92 < 19; fi92++ {
                                                                v1 = (v1 - -123)
                                                                _ = fi92
                                                            }
                                                        }
                                                    case 128:
                                                        if true {
                                                            for fi93 := 0; fi93 < 15; fi93++ {
                                                                if len(v2) > 0 {
                                                                    v2[0] = uint32(v86)
                                                                }
                                                                if ((!v38) || (85 == (v86 - v86))) {
                                                                    switch "a" {
                                                                    case "a":
                                                                        v83 += (^(v4))
                                                                    case "\n\t":
                                                                        v39 = 408
                                                                    }
                                                                    var v94 int64 = 784
                                                                    v95 := GenS5{F0: (^(v86)), F1: 32767, F2: "\n\t"}
                                                                    _ = v94
                                                                    _ = v95
                                                                } else {
                                                                    v96 := GenS9{F0: 2147483647, F1: -140, F2: 940}
                                                                    _ = v96
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    if (!false) {
                                                                        if (2147483647 < int32(v3)) {
                                                                            var v97 bool = v38
                                                                            _ = v97
                                                                        }
                                                                        v98 := GenS10{F0: 0}
                                                                        v83 = genFn0()
                                                                        _ = v98
                                                                    } else {
                                                                        var v99 string = (("\n\t" + "a") + "\n\t")
                                                                        _ = v99
                                                                    }
                                                                }
                                                                var v100 int64 = 64
                                                                _ = fi93
                                                                _ = v100
                                                            }
                                                            fmt.Println(v1, v4, v25)
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = 635
                                                            }
                                                            if (float32(0.1) == 0.0) {
                                                                var v101 uint8 = v86
                                                                for fi102 := 0; fi102 < 17; fi102++ {
                                                                    var v103 uint = 421
                                                                    _ = fi102
                                                                    _ = v103
                                                                }
                                                                _ = v101
                                                            } else {
                                                                fmt.Println(v4)
                                                            }
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v2 = append(v2, 42)
                                                            case "":
                                                                v3 = v3
                                                            case "hello":
                                                                fmt.Println(fi40, v3, fi40)
                                                            default:
                                                                var v104 int = 849
                                                                _ = v104
                                                            }
                                                        }
                                                    default:
                                                        var v105 int32 = -436
                                                        _ = v105
                                                    }
                                                default:
                                                    var v106 int8 = v1
                                                    _ = v106
                                                }
                                                var v107 uint64 = 2
                                                v108 := GenS2{F0: v86, F1: 1, F2: 524}
                                                _ = v107
                                                _ = v108
                                            } else {
                                                var v109 bool = v38
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v109
                                            }
                                        } else {
                                            v110 := GenS8{F0: 0, F1: 4294967295}
                                            _ = v110
                                        }
                                        _ = v86
                                    } else {
                                        fmt.Println(fi40, v83)
                                        var v111 int32 = 42
                                        _ = v111
                                    }
                                    v2 = append(v2, 2)
                                    var v112 uint64 = 42
                                    _ = v112
                                }
                                switch fi40 {
                                case 9223372036854775807:
                                    fmt.Println(v4, v38, v4)
                                case 1:
                                    if v38 {
                                        for fi113 := 0; fi113 < 15; fi113++ {
                                            var v114 int64 = int64(v39)
                                            _ = fi113
                                            _ = v114
                                        }
                                    }
                                case 885:
                                    for fi115 := 0; fi115 < 3; fi115++ {
                                        if len(v2) > 0 {
                                            v2[0] = 4294967295
                                        }
                                        v4 -= 0
                                        switch -1 {
                                        case 440:
                                            v116 := GenS8{F0: -198, F1: 551}
                                            _ = v116
                                        case 9223372036854775807:
                                            v1 -= 100
                                        case -1:
                                            fmt.Println(v3, v38)
                                        default:
                                            fmt.Println(fi115, v83, v3)
                                        }
                                        var v117 uint64 = (v39 | (v39 - (v39 & (v39 * (v39 & (v39 | 587))))))
                                        _ = fi115
                                        _ = v117
                                    }
                                }
                            } else {
                                v118 := GenS11{F0: 551, F1: v38}
                                _ = v118
                            }
                            v119 := GenS10{F0: 42}
                            if v38 {
                                var v120 int = fi40
                                for fi121 := 0; fi121 < 10; fi121++ {
                                    var v122 int8 = genFn0()
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi121
                                    _ = v122
                                }
                                _ = v120
                            } else {
                                if true {
                                    fmt.Println(v1, v83)
                                } else {
                                    v123 := GenS12{F0: 519, F1: 2147483647, F2: "hello"}
                                    _ = v123
                                }
                                v38 = true
                                for fi124 := 0; fi124 < 2; fi124++ {
                                    v125 := GenS13{F0: -312, F1: fi124, F2: 1}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v1)
                                    _ = fi124
                                    _ = v125
                                }
                            }
                            _ = v119
                        } else {
                            v39 -= v39
                            switch ((("" + "") + "") + "hello") {
                            case "\n\t":
                                switch "\n\t" {
                                case "":
                                    if v38 {
                                        var v126 uint32 = 1
                                        fmt.Println(fi40, v1, v25)
                                        switch fi40 {
                                        case 0:
                                            switch ("hello" + "\n\t") {
                                            case "hello":
                                                if (false && true) {
                                                    v1 += (v1 * 1)
                                                    if (2147483647 > -250) {
                                                        switch fi40 {
                                                        case 32767:
                                                            if len(v2) > 0 {
                                                                v2[0] = uint32(v39)
                                                            }
                                                        case 127:
                                                            fmt.Println(v39, v39, fi40)
                                                        default:
                                                            if v38 {
                                                                for fi127 := 0; fi127 < 15; fi127++ {
                                                                    v38 = (1 >= 42)
                                                                    if false {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            if ((v1 / (((v83 >> ((784) & 63))) | 1)) <= 33) {
                                                                                v128 := GenS14{F0: (fi40 >> ((uint(v39)) & 63)), F1: v1, F2: (v39 / (((v39 >> ((18446744073709551615) & 63))) | 1))}
                                                                                var v129 uint32 = v126
                                                                                v130 := GenS15{F0: 1}
                                                                                _ = v128
                                                                                _ = v129
                                                                                _ = v130
                                                                            } else {
                                                                                if (684 > v39) {
                                                                                    var v131 uint = 18446744073709551615
                                                                                    v83 -= (v4 / (((-(v1))) | 1))
                                                                                    v1 += v1
                                                                                    _ = v131
                                                                                } else {
                                                                                    v132 := GenS10{F0: 0}
                                                                                    var v133 int16 = int16(fi127)
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                }
                                                                                v83 = (v4 ^ v1)
                                                                                v134 := GenS16{F0: uint32(v39), F1: ("hello" + "a"), F2: (v39 + v39)}
                                                                                _ = v134
                                                                            }
                                                                        case "a":
                                                                            var v135 uint8 = 73
                                                                            _ = v135
                                                                        case "foo bar":
                                                                            var v136 uint64 = v39
                                                                            _ = v136
                                                                        default:
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v39 = v39
                                                                                switch fi40 {
                                                                                case 4294967295:
                                                                                    fmt.Println(v38, v25)
                                                                                    v25 -= v25
                                                                                case 0:
                                                                                    v2 = append(v2, (v126 % ((2) | 1)))
                                                                                default:
                                                                                    var v137 bool = (v25 >= v25)
                                                                                    _ = v137
                                                                                }
                                                                            case "\n\t":
                                                                                fmt.Println(v25)
                                                                            }
                                                                        }
                                                                        v38 = v38
                                                                    }
                                                                    _ = fi127
                                                                }
                                                                if (4294967295 != 898) {
                                                                    for fi138 := 0; fi138 < 9; fi138++ {
                                                                        v25 += float32(0.1)
                                                                        if v38 {
                                                                            var v139 bool = (v25 <= ((1e10 + v25) / (((float32(0.1) / ((-2.25 - (-2.25 * 0.0)) + (v25 * ((((v25 / float32((((3.14159 / 0.1) + 1.5) * 1.5))) + float32(0.0)) / v25) * v25)))) * (v25 / 3.14159)) - v25)))
                                                                            _ = v139
                                                                        } else {
                                                                            v38 = ((0 < 454) || (v3 >= (^(v3))))
                                                                            fmt.Println(fi40, fi40, fi40)
                                                                        }
                                                                        _ = fi138
                                                                    }
                                                                    v140 := GenS2{F0: 6, F1: 0, F2: fi40}
                                                                    fmt.Println(v39, v1, v83)
                                                                    _ = v140
                                                                }
                                                            }
                                                        }
                                                        if (v38 || ((fi40 - fi40) < 127)) {
                                                            v4 += (v1 + v83)
                                                            v39 = v39
                                                        }
                                                    } else {
                                                        if v38 {
                                                            v3 -= v3
                                                            switch fi40 {
                                                            case -637:
                                                                v39 = v39
                                                            case 607:
                                                                v83 += (v83 + (v4 + -128))
                                                            case 32768:
                                                                if true {
                                                                    v25 += (1.5 - 1.5)
                                                                    for fi141 := 0; fi141 < 12; fi141++ {
                                                                        var v142 uint8 = 128
                                                                        v143 := GenS17{F0: 42}
                                                                        v144 := GenS17{F0: -1}
                                                                        _ = fi141
                                                                        _ = v142
                                                                        _ = v143
                                                                        _ = v144
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 1
                                                                    }
                                                                }
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        for fi145 := 0; fi145 < 2; fi145++ {
                                                            v1 += (v1 | v1)
                                                            switch -17 {
                                                            case 128:
                                                                if (1 <= 0) {
                                                                    var v146 float64 = (0.0 / (((1.5 * float64((v25 * v25))) / (3.14159 - 0.0)) - 0.0))
                                                                    _ = v146
                                                                }
                                                            case 2147483647:
                                                                v147 := GenS7{F0: v39, F1: true, F2: (^(v3))}
                                                                _ = v147
                                                            default:
                                                                fmt.Println(v83, v3, v1)
                                                            }
                                                            _ = fi145
                                                        }
                                                        v148 := GenS11{F0: 0, F1: v38}
                                                        _ = v148
                                                    }
                                                }
                                            }
                                        case 128:
                                            for fi149 := 0; fi149 < 20; fi149++ {
                                                var v150 int = fi40
                                                switch (v150 * int(v4)) {
                                                case 9223372036854775807:
                                                    fmt.Println(v126, v3, v39)
                                                    if v38 {
                                                        v2 = append(v2, v126)
                                                        v3 = (v3 - v3)
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            _ = len(v2)
                                                        case "a":
                                                            v151 := GenS18{F0: (v126 * 2), F1: 1, F2: 123}
                                                            _ = v151
                                                        }
                                                        var v152 float64 = (((3.14159 / float64(v25)) - float64(v25)) + 1e10)
                                                        _ = v152
                                                    }
                                                    v126 -= 377
                                                case -9223372036854775808:
                                                    fmt.Println(v38, v38)
                                                default:
                                                    var v153 uint16 = 310
                                                    _ = v153
                                                }
                                                if (!v38) {
                                                    _ = len(v2)
                                                    v154 := GenS12{F0: -579, F1: 458, F2: ""}
                                                    _ = v154
                                                }
                                                _ = fi149
                                                _ = v150
                                            }
                                        case 4294967295:
                                            for fi155 := 0; fi155 < 6; fi155++ {
                                                v1 = 1
                                                v83 = (^(v4))
                                                _ = fi155
                                            }
                                        }
                                        _ = v126
                                    }
                                case "foo bar":
                                    var v156 uint64 = 908
                                    _ = v156
                                case "\n\t":
                                    v2 = append(v2, 227)
                                default:
                                    if v38 {
                                        for fi157 := 0; fi157 < 8; fi157++ {
                                            v4 += 127
                                            _ = fi157
                                        }
                                    }
                                }
                            case "hello":
                                v38 = (255 > 65535)
                            case "":
                                fmt.Println(v3, v3, v1)
                            default:
                                v1 = v4
                            }
                            for fi158 := 0; fi158 < 18; fi158++ {
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                                var v159 uint32 = 2147483648
                                _ = fi158
                                _ = v159
                            }
                        }
                        v25 = (v25 / (3.14159 / v25))
                        v160 := GenS13{F0: (v3 << ((650) & 63)), F1: (fi40 % (((fi40 ^ fi40)) | 1)), F2: 101}
                        _ = v160
                    } else {
                        v161 := GenS6{F0: v39, F1: (^(v3)), F2: ("" + ("a" + "\n\t"))}
                        if v38 {
                            v4 += genFn0()
                        }
                        var v162 uint8 = 0
                        _ = v161
                        _ = v162
                    }
                    _ = v83
                }
            case "hello":
                for fi163 := 0; fi163 < 12; fi163++ {
                    var v164 int64 = -420
                    var v165 int8 = -18
                    switch "hello" {
                    case "\n\t":
                        var v166 uint64 = 4294967295
                        _ = v166
                    default:
                        if false {
                            var v167 float32 = 0.0
                            switch int(v165) {
                            case 4294967295:
                                var v168 uint32 = 2147483648
                                _ = v168
                            case 32768:
                                if ("a" <= "") {
                                    v3 = -32768
                                    v169 := GenS5{F0: 128, F1: 101, F2: "\n\t"}
                                    v2 = append(v2, 1)
                                    _ = v169
                                } else {
                                    v170 := GenS7{F0: (^(v39)), F1: (true || (412 < 797)), F2: (v3 & (v3 & (v3 * v3)))}
                                    if true {
                                        v171 := GenS13{F0: v3, F1: (-(fi40)), F2: 127}
                                        v172 := GenS15{F0: 2}
                                        _ = len(v2)
                                        _ = v171
                                        _ = v172
                                    } else {
                                        v173 := GenS19{F0: "\n\t"}
                                        if (v38 || v38) {
                                            _ = len(v2)
                                            fmt.Println(v165)
                                            v38 = (590 < 255)
                                        }
                                        _ = v173
                                    }
                                    v164 -= 137
                                    _ = v170
                                }
                                var v174 int = fi163
                                _ = len(v2)
                                fmt.Println(v38, v4, v167)
                                _ = v174
                            }
                            if v38 {
                                fmt.Println(v38, fi40)
                            } else {
                                v175 := GenS12{F0: (v164 ^ 453), F1: 2147483647, F2: "hello"}
                                v4 -= (v4 >> ((0) & 63))
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                                _ = v175
                            }
                            _ = v167
                        }
                    }
                    _ = fi163
                    _ = v164
                    _ = v165
                }
            default:
                var v176 bool = false
                _ = v176
            }
            v38 = v38
        }
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
