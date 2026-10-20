package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int32; F1 uint; F2 string }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 int32; F1 int32; F2 uint }
type GenS12 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int64; F1 int16; F2 uint }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int; F1 uint16 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 uint; F1 uint64 }
type GenS22 struct { F0 uint8 }
type GenS23 struct { F0 int16; F1 int32; F2 uint16 }
type GenS24 struct { F0 uint }
type GenS25 struct { F0 string; F1 uint16 }
func genFn0(p0 uint8, p1 bool, p2 uint16) string {
    var v5 float32 = float32(((float64((((-2.25 - (3.14159 * (float32(0.1) * (0.0 * (-2.25 - (((3.14159 - 0.0) - 3.14159) * 0.1)))))) / 3.14159) * 0.0)) + 0.0) / -2.25))
    v5 += 3.14159
    switch 320 {
    case 32768:
        v6 := []int{-2147483648, 925, 4294967295, -1}
        _ = v6
    case 725:
        v5 -= float32(1e10)
    case -1:
        v7 := GenS1{F0: 761, F1: -1}
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int64 = -540
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("\n\t" + ("hello" + "a"))
}
func genFn1(p0 int16, p1 int32, p2 uint8) int {
    p0 = p0
    v10 := []int{-9223372036854775808, 127, 0}
    v11 := GenS1{F0: 566, F1: 42}
    for fi12 := 0; fi12 < 16; fi12++ {
        v10 = append(v10, 2147483647)
        var v13 string = "\n\t"
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v10 = append(v10, fi14)
        v15 := GenS1{F0: 18446744073709551615, F1: 626}
        p1 -= p1
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 bool = (537 == 1)
    var v2 bool = (!false)
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    for fi16 := 0; fi16 < 1; fi16++ {
        _ = len(v3)
        if false {
            fmt.Println(v1, v1)
            v2 = (9223372036854775807 > -1)
            v1 = (66 <= 127)
        } else {
            if (!(!(23 != 1))) {
                v1 = true
                v17 := GenS2{F0: 60, F1: 42}
                for fi18 := 0; fi18 < 3; fi18++ {
                    if (false || (32767 == 2)) {
                        var v19 uint32 = 1
                        _ = v19
                    } else {
                        if v1 {
                            if v1 {
                                if ((-1 <= -835) && (1.5 > (1e10 / float32(-2.25)))) {
                                    v2 = (-9223372036854775808 >= 441)
                                }
                                if (false && (fi18 < 266)) {
                                    var v20 float64 = float64(((0.0 + ((-2.25 / -2.25) - float32(float64(1.5)))) - (float32(float64(3.14159)) * (float32(1e10) + (-2.25 / 1.5)))))
                                    if true {
                                        fmt.Println(v1, v20, v1)
                                        if (1.5 > 1.5) {
                                            var v21 uint8 = 0
                                            _ = v21
                                        } else {
                                            var v22 bool = v2
                                            if v22 {
                                                v20 = (float64((3.14159 - (-2.25 * 1e10))) - v20)
                                                v1 = v2
                                            }
                                            _ = v22
                                        }
                                    }
                                    v2 = v1
                                    _ = v20
                                }
                                v23 := []int{-843}
                                _ = v23
                            } else {
                                v24 := []int{fi16, (fi18 * -9223372036854775808), fi16, (fi18 << ((4294967295) & 63))}
                                v1 = (65535 > 529)
                                var v25 uint16 = 65535
                                _ = v24
                                _ = v25
                            }
                            fmt.Println(fi16, fi18)
                            v26 := GenS2{F0: 42, F1: 278}
                            v27 := []int{fi16}
                            _ = v26
                            _ = v27
                        }
                        switch (fi16 & genFn1(1000, -2147483648, 127)) {
                        case 1:
                            var v28 string = "\n\t"
                            _ = v28
                        }
                    }
                    if v2 {
                        _ = len(v3)
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            v1 = (fi16 != 716)
                        case "hello":
                            v2 = (0 < 0)
                        case "":
                            v29 := GenS3{F0: 833, F1: 65535}
                            _ = v29
                        }
                        delete(v4, 961)
                        switch ("\n\t" + "foo bar") {
                        case "\n\t":
                            var v30 uint8 = 128
                            _ = v30
                        case "foo bar":
                            delete(v4, 2147483648)
                        default:
                            switch fi16 {
                            case -9223372036854775808:
                                if v2 {
                                    var v31 uint64 = 4294967295
                                    v32 := GenS4{F0: -9223372036854775808, F1: 2}
                                    v33 := v3[9223372036854775807]
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                } else {
                                    if (0 >= 42) {
                                        v34 := GenS5{F0: ("foo bar" + ("hello" + ("a" + "")))}
                                        switch "" {
                                        case "foo bar":
                                            v35 := GenS6{F0: 65535, F1: 65535}
                                            _ = v35
                                        case "hello":
                                            v1 = (true && (!(42 <= -128)))
                                            var v36 bool = v1
                                            v36 = v1
                                            _ = v36
                                        case "":
                                            v1 = false
                                        default:
                                            if (-1 >= -1) {
                                                fmt.Println(fi18)
                                                v3[(fi18 & fi16)] = "\n\t"
                                                if (fi18 != (fi18 >> ((1) & 63))) {
                                                    v37 := []int{genFn1(42, -1, 157)}
                                                    var v38 uint8 = 2
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                        }
                                        _ = v34
                                    } else {
                                        var v39 uint8 = 255
                                        _ = v39
                                    }
                                    var v40 string = ("" + "")
                                    _ = v40
                                }
                            case 2147483647:
                                if false {
                                    v41 := GenS7{F0: 1, F1: 4294967295, F2: "a"}
                                    switch "a" {
                                    case "foo bar":
                                        v2 = (!v1)
                                    }
                                    v1 = v2
                                    _ = v41
                                } else {
                                    var v42 float32 = 0.1
                                    switch (fi18 ^ genFn1(int16(fi16), 1, 0)) {
                                    case -467:
                                        var v43 uint8 = 128
                                        _ = v43
                                    case 481:
                                        v44 := []int{(^(fi18)), -2147483648}
                                        _ = v44
                                    default:
                                        v1 = (v42 <= 1.5)
                                    }
                                    _ = v42
                                }
                            default:
                                if (0 >= -9223372036854775808) {
                                    if v2 {
                                        fmt.Println(v2)
                                    } else {
                                        var v45 uint16 = 32767
                                        v1 = (4294967295 == 348)
                                        _ = v45
                                    }
                                    v1 = (-92 < 1)
                                    var v46 int16 = 366
                                    _ = v46
                                } else {
                                    v2 = (1.5 <= 0.1)
                                }
                            }
                        }
                    }
                    _ = fi18
                }
                _ = v17
            } else {
                if v2 {
                    v47 := []int{(fi16 / ((127) | 1)), genFn1(32767, 0, 1), (fi16 - 128), fi16}
                    var v48 int64 = -956
                    if v2 {
                        for fi49 := 0; fi49 < 14; fi49++ {
                            if v1 {
                                v50 := GenS8{F0: -1, F1: int16(v48)}
                                _ = v50
                            } else {
                                v51 := GenS9{F0: 42}
                                v52 := GenS10{F0: (fi16 | (fi49 * 32768)), F1: 898}
                                v48 -= (v48 % ((v48) | 1))
                                _ = len(v47)
                                _ = v51
                                _ = v52
                            }
                            fmt.Println(fi16)
                            _ = fi49
                        }
                        v1 = v1
                        fmt.Println(v48, v1, v48)
                    } else {
                        for fi53 := 0; fi53 < 5; fi53++ {
                            if (!(!((0.0 - 0.1) != float32((0.1 - 1e10))))) {
                                if (!v1) {
                                    if len(v47) > 0 {
                                        v47[0] = genFn1(-32768, -608, 0)
                                    }
                                } else {
                                    fmt.Println(v1, v1, fi16)
                                    v54 := GenS1{F0: 756, F1: v48}
                                    _ = v54
                                }
                                var v55 bool = true
                                switch "\n\t" {
                                case "a":
                                    v56 := GenS9{F0: 420}
                                    _ = v56
                                case "\n\t":
                                    v57 := GenS2{F0: 32767, F1: 42}
                                    _ = v57
                                case "hello":
                                    if (127 >= -38) {
                                        switch genFn1(519, 2147483647, 97) {
                                        case -386:
                                            var v58 uint16 = 42
                                            _ = v58
                                        case 2147483647:
                                            switch ("hello" + "") {
                                            case "a":
                                                v55 = (0.0 <= -2.25)
                                            case "":
                                                switch "" {
                                                case "foo bar":
                                                    switch "" {
                                                    case "\n\t":
                                                        v48 = v48
                                                    case "":
                                                        v2 = v1
                                                    case "foo bar":
                                                        fmt.Println(v2, fi16)
                                                    default:
                                                        var v59 bool = (2 <= 42)
                                                        _ = v59
                                                    }
                                                    var v60 string = ("" + "\n\t")
                                                    v61 := GenS4{F0: 0, F1: 27}
                                                    _ = v60
                                                    _ = v61
                                                case "a":
                                                    v2 = ((fi53 << ((4294967295) & 63)) >= fi16)
                                                case "":
                                                    v62 := GenS3{F0: 65535, F1: 255}
                                                    _ = v62
                                                }
                                            case "hello":
                                                v63 := v3[1]
                                                _ = v63
                                            default:
                                                fmt.Println(v2)
                                            }
                                        case -1:
                                            fmt.Println(v2, fi16, v48)
                                        default:
                                            switch "" {
                                            case "hello":
                                                var v64 uint = 1
                                                _ = v64
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi16, fi53)
                                        if (65535 > 706) {
                                            fmt.Println(fi53, fi53, v48)
                                            switch (^(fi53)) {
                                            case 42:
                                                var v65 string = "\n\t"
                                                _ = v65
                                            }
                                        } else {
                                            if v2 {
                                                v55 = v2
                                                if (350 > 255) {
                                                    var v66 float64 = (1e10 + 0.1)
                                                    if len(v47) > 0 {
                                                        v47[0] = (fi53 << ((65535) & 63))
                                                    }
                                                    if false {
                                                        v2 = v55
                                                        fmt.Println(v48, v66)
                                                    }
                                                    var v67 uint = 541
                                                    if (2147483647 >= 0) {
                                                        if true {
                                                            v47 = append(v47, fi53)
                                                        }
                                                    } else {
                                                        fmt.Println(fi53, v67)
                                                        var v68 uint64 = 2
                                                        var v69 uint16 = 493
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    if v1 {
                                                        if len(v47) > 0 {
                                                            v47[0] = fi16
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v70 uint32 = 0
                                                            if (128 == 127) {
                                                                var v71 int = (fi53 % ((fi16) | 1))
                                                                var v72 string = "hello"
                                                                v73 := GenS1{F0: (v67 & v67), F1: (v48 * (v48 * (v48 + (v48 << ((65535) & 63)))))}
                                                                switch -358 {
                                                                case -1:
                                                                    v48 = v48
                                                                    _ = len(v47)
                                                                case -204:
                                                                    if (v2 || v1) {
                                                                        var v74 int64 = (v48 / ((v48) | 1))
                                                                        var v75 int = (v71 - genFn1(-1, 0, 0))
                                                                        var v76 int32 = 996
                                                                        _ = v74
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                case -139:
                                                                    var v77 int64 = -393
                                                                    _ = v77
                                                                default:
                                                                    switch v72 {
                                                                    case "a":
                                                                        var v78 int = fi16
                                                                        _ = v78
                                                                    case "\n\t":
                                                                        if (!(217 >= 0)) {
                                                                            switch ("hello" + v72) {
                                                                            case "a":
                                                                                v2 = v1
                                                                            case "foo bar":
                                                                                if (0.0 < (0.0 - 3.14159)) {
                                                                                    if v55 {
                                                                                        v47 = append(v47, fi16)
                                                                                        var v79 bool = (!("hello" == "foo bar"))
                                                                                        _ = v79
                                                                                    }
                                                                                    var v80 bool = true
                                                                                    _ = v80
                                                                                } else {
                                                                                    v81 := GenS3{F0: 1, F1: (^(v67))}
                                                                                    var v82 int32 = -1
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                            var v83 uint16 = 1
                                                                            var v84 string = "a"
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            if false {
                                                                                var v85 int8 = 88
                                                                                _ = v85
                                                                            } else {
                                                                                switch v72 {
                                                                                case "":
                                                                                    var v86 float64 = -2.25
                                                                                    _ = v86
                                                                                case "a":
                                                                                    v87 := GenS11{F0: 42, F1: -2147483648, F2: (v67 ^ (v67 ^ (v67 - v67)))}
                                                                                    _ = v87
                                                                                default:
                                                                                    v66 += v66
                                                                                }
                                                                                var v88 int8 = 1
                                                                                delete(v4, 2)
                                                                                _ = v88
                                                                            }
                                                                            var v89 uint8 = 42
                                                                            var v90 bool = (62 > int8(v67))
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                    case "hello":
                                                                        v48 = 1
                                                                    }
                                                                }
                                                                v91 := GenS12{F0: 59, F1: 42, F2: 474}
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                                _ = v91
                                                            } else {
                                                                v55 = v2
                                                                fmt.Println(v1, v66, v66)
                                                            }
                                                            v92 := GenS13{F0: 73}
                                                            v2 = (2147483647 <= 1)
                                                            v66 += 1.5
                                                            _ = v70
                                                            _ = v92
                                                        case "":
                                                            var v93 int64 = v48
                                                            _ = v93
                                                        case "hello":
                                                            v67 = v67
                                                        default:
                                                            var v94 int = (fi53 << (((v67 << ((4294967295) & 63))) & 63))
                                                            _ = v94
                                                        }
                                                        v95 := GenS8{F0: fi16, F1: -1}
                                                        var v96 float32 = 1.5
                                                        var v97 int = (fi16 ^ 32767)
                                                        _ = v95
                                                        _ = v96
                                                        _ = v97
                                                    } else {
                                                        v98 := GenS14{F0: (42 != -1)}
                                                        var v99 int = 42
                                                        v100 := v4[6]
                                                        _ = v100
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                    if (65535 != 1) {
                                                        v1 = (564 < 686)
                                                        if true {
                                                            var v101 int16 = -32768
                                                            _ = v101
                                                        }
                                                        switch ("\n\t" + ("foo bar" + "foo bar")) {
                                                        case "a":
                                                            delete(v3, (fi53 >> ((4294967295) & 63)))
                                                        default:
                                                            v3[genFn1(32767, 2147483647, 2)] = "foo bar"
                                                        }
                                                    } else {
                                                        v102 := GenS4{F0: v48, F1: 343}
                                                        var v103 int16 = 0
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    if ("a" <= ("foo bar" + ("\n\t" + "\n\t"))) {
                                                        if false {
                                                            switch fi53 {
                                                            case 54:
                                                                fmt.Println(v48, v1, v48)
                                                            case 199:
                                                                v2 = v55
                                                            default:
                                                                switch "" {
                                                                case "foo bar":
                                                                    fmt.Println(fi53, v55, v48)
                                                                case "a":
                                                                    var v104 string = ("foo bar" + "")
                                                                    switch (fi16 & fi53) {
                                                                    case 128:
                                                                        v2 = v1
                                                                    case 127:
                                                                        fmt.Println(v48, fi53, fi53)
                                                                    case 4294967295:
                                                                        var v105 int64 = (v48 | v48)
                                                                        _ = v105
                                                                    default:
                                                                        v2 = false
                                                                        switch genFn1(1, 0, 170) {
                                                                        case -749:
                                                                            var v106 int16 = -485
                                                                            _ = v106
                                                                        case 128:
                                                                            v107 := GenS15{F0: 542}
                                                                            _ = v107
                                                                        default:
                                                                            var v108 bool = v2
                                                                            _ = v108
                                                                        }
                                                                        v109 := GenS15{F0: 22}
                                                                        if v1 {
                                                                            var v110 uint8 = 247
                                                                            _ = v110
                                                                        }
                                                                        _ = v109
                                                                    }
                                                                    _ = v104
                                                                default:
                                                                    v48 += (-(v48))
                                                                }
                                                            }
                                                            v1 = v2
                                                            if v2 {
                                                                if (false && true) {
                                                                    v111 := GenS6{F0: 649, F1: 2}
                                                                    _ = v111
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v48 -= v48
                                                        var v112 int64 = v48
                                                        v113 := GenS14{F0: (1000 == 42)}
                                                        fmt.Println(fi16)
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                }
                                                switch fi16 {
                                                case -592:
                                                    fmt.Println(v1)
                                                    if (42 != 0) {
                                                        if (42 <= 132) {
                                                            if (!v1) {
                                                                v114 := GenS12{F0: -54, F1: 1, F2: 42}
                                                                fmt.Println(v1, fi53, fi16)
                                                                _ = v114
                                                            }
                                                            fmt.Println(v1, fi53)
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if v55 {
                                                                    fmt.Println(v1, v1)
                                                                    v115 := GenS16{F0: (v48 >> ((270) & 63)), F1: -513, F2: 0}
                                                                    var v116 bool = ((!v2) || (656 <= 65535))
                                                                    _ = v115
                                                                    _ = v116
                                                                }
                                                            case "":
                                                                v1 = v2
                                                            }
                                                        }
                                                        v2 = true
                                                        v2 = (34 > 95)
                                                    } else {
                                                        fmt.Println(fi16, v48)
                                                        var v117 uint16 = 32767
                                                        _ = v117
                                                    }
                                                    if v1 {
                                                        fmt.Println(v48)
                                                        v48 -= v48
                                                    } else {
                                                        if false {
                                                            v118 := GenS17{F0: 0}
                                                            switch ("hello" + "\n\t") {
                                                            case "foo bar":
                                                                var v119 float32 = (((float32(float64(-2.25)) / 1e10) - 1.5) * 1.5)
                                                                _ = v119
                                                            }
                                                            v48 -= 9223372036854775807
                                                            _ = v118
                                                        } else {
                                                            v120 := GenS15{F0: 395}
                                                            if true {
                                                                v2 = v2
                                                            }
                                                            v121 := GenS4{F0: v48, F1: 2}
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        _ = len(v3)
                                                        v55 = (1.5 > 1e10)
                                                    }
                                                    if len(v47) > 0 {
                                                        v47[0] = fi53
                                                    }
                                                case 1:
                                                    v122 := GenS5{F0: (("hello" + "") + "hello")}
                                                    _ = v122
                                                case 676:
                                                    if len(v47) >= 2 {
                                                        v47 = v47[1:len(v47)]
                                                    }
                                                    if false {
                                                        if v2 {
                                                            fmt.Println(fi16, fi53, fi16)
                                                            var v123 float32 = 0.0
                                                            v55 = false
                                                            switch (fi53 ^ genFn1(1, 0, 255)) {
                                                            case -752:
                                                                v124 := GenS14{F0: v2}
                                                                _ = v124
                                                            case -1:
                                                                var v125 uint64 = 1
                                                                _ = v125
                                                            case -9223372036854775808:
                                                                if ((-128 < int8(v48)) && false) {
                                                                    fmt.Println(v123)
                                                                    if v2 {
                                                                        v48 = 0
                                                                        fmt.Println(v123, v123, v123)
                                                                    } else {
                                                                        v1 = (154 < -32768)
                                                                        var v126 int = genFn1(1, -112, 1)
                                                                        _ = v126
                                                                    }
                                                                }
                                                            }
                                                            _ = v123
                                                        } else {
                                                            v127 := GenS10{F0: (fi53 * fi16), F1: 0}
                                                            v1 = v1
                                                            _ = v127
                                                        }
                                                        v55 = false
                                                        fmt.Println(fi16, v55)
                                                        v128 := GenS18{F0: 32768, F1: 1}
                                                        _ = v128
                                                    } else {
                                                        v129 := GenS15{F0: 18446744073709551615}
                                                        _ = v129
                                                    }
                                                    var v130 int64 = v48
                                                    _ = v130
                                                }
                                            }
                                        }
                                        v131 := GenS15{F0: 0}
                                        _ = v131
                                    }
                                default:
                                    switch ((("\n\t" + "foo bar") + "foo bar") + (("\n\t" + (("hello" + "hello") + "a")) + "hello")) {
                                    case "foo bar":
                                        var v132 uint = 0
                                        _ = v132
                                    case "\n\t":
                                        var v133 string = "a"
                                        _ = v133
                                    }
                                }
                                _ = v55
                            }
                            _ = fi53
                        }
                        if v1 {
                            fmt.Println(v48)
                            switch "" {
                            case "\n\t":
                                if (!(!(v2 && v2))) {
                                    v4[4294967295] = 4294967295
                                    for fi134 := 0; fi134 < 2; fi134++ {
                                        v135 := GenS18{F0: 455, F1: 2}
                                        _ = fi134
                                        _ = v135
                                    }
                                    v2 = v1
                                }
                            case "foo bar":
                                for fi136 := 0; fi136 < 12; fi136++ {
                                    var v137 uint64 = 0
                                    _ = fi136
                                    _ = v137
                                }
                            case "hello":
                                for fi138 := 0; fi138 < 2; fi138++ {
                                    fmt.Println(fi16, v48, v2)
                                    fmt.Println(fi138)
                                    delete(v4, 1)
                                    _ = fi138
                                }
                            default:
                                v139 := v4[2]
                                _ = v139
                            }
                            fmt.Println(v1, v48)
                        } else {
                            if len(v47) >= 2 {
                                v47 = v47[1:len(v47)]
                            }
                        }
                        var v140 int8 = int8(fi16)
                        _ = v140
                    }
                    _ = v47
                    _ = v48
                }
            }
            v2 = ((-454 < 562) && ((!true) || v1))
        }
        fmt.Println(v1, v1, v2)
        _ = fi16
    }
    if ((3.14159 == (float32((((1.5 - (3.14159 - 0.1)) * 0.1) / (((-2.25 - (0.0 * 1e10)) * ((((0.0 - 0.0) / 3.14159) - (0.0 - (((0.0 / 1.0) - (float64((1.5 + (0.0 + 0.1))) / (float64(3.14159) / 1.5))) - 1e10))) / 1.0)) * 1e10))) - 1e10)) || false) {
        var v141 uint64 = 1
        _ = v141
    } else {
        for fi142 := 0; fi142 < 13; fi142++ {
            var v143 uint8 = 42
            _ = fi142
            _ = v143
        }
        v144 := v3[genFn1(0, -2147483648, 92)]
        var v145 uint8 = 127
        _ = v144
        _ = v145
    }
    for fi146 := 0; fi146 < 9; fi146++ {
        v1 = v2
        v1 = v2
        if (1 < 448) {
            for fi147 := 0; fi147 < 11; fi147++ {
                var v148 float32 = 1.5
                _ = fi147
                _ = v148
            }
            v4[784] = 0
            for fi149 := 0; fi149 < 10; fi149++ {
                v2 = v2
                if v2 {
                    var v150 int16 = -492
                    switch 0 {
                    case -2:
                        switch 1 {
                        case -9223372036854775808:
                            if (true || v1) {
                                v1 = v2
                                v151 := []int{fi149, fi146}
                                var v152 int32 = 42
                                fmt.Println(v1, v152, v1)
                                switch "a" {
                                case "foo bar":
                                    var v153 float32 = ((1e10 + (0.1 - ((0.1 * 0.0) + 0.1))) - (-2.25 + 0.1))
                                    _ = v153
                                case "hello":
                                    fmt.Println(v150)
                                }
                                var v154 uint = 1
                                _ = v151
                                _ = v152
                                _ = v154
                            }
                        case 42:
                            var v155 int8 = 100
                            switch "" {
                            case "foo bar":
                                if (v150 < v150) {
                                    v155 = v155
                                    fmt.Println(fi149)
                                } else {
                                    v155 = (v155 << ((0) & 63))
                                    switch -9223372036854775808 {
                                    case -1:
                                        v150 -= (v150 & 42)
                                    case 32767:
                                        fmt.Println(fi146)
                                    default:
                                        var v156 uint64 = 568
                                        _ = v156
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        delete(v3, -9223372036854775808)
                                    }
                                }
                            case "a":
                                v157 := GenS11{F0: -212, F1: -2147483648, F2: 4294967295}
                                if v1 {
                                    switch ("a" + "foo bar") {
                                    case "foo bar":
                                        v1 = false
                                    case "a":
                                        if v2 {
                                            v158 := []int{genFn1((v150 + (v150 | 32767)), -1, 128)}
                                            _ = v158
                                        } else {
                                            v155 -= (v155 + (v155 | (v155 >> ((643) & 63))))
                                            var v159 bool = v2
                                            _ = v159
                                        }
                                        v155 = v155
                                    case "hello":
                                        var v160 int32 = 42
                                        _ = v160
                                    }
                                    v161 := GenS11{F0: 1, F1: 0, F2: 4294967295}
                                    if (!(-1 < 345)) {
                                        var v162 bool = ((v1 || false) && (42 == 42))
                                        v163 := GenS5{F0: "hello"}
                                        _ = v162
                                        _ = v163
                                    }
                                    _ = v161
                                }
                                _ = len(v3)
                                v2 = (!false)
                                _ = v157
                            case "\n\t":
                                if v2 {
                                    var v164 int = 2147483647
                                    v155 += v155
                                    var v165 uint8 = 60
                                    v165 -= 255
                                    _ = v164
                                    _ = v165
                                } else {
                                    switch (fi146 / (((fi149 | (fi149 * 9223372036854775807))) | 1)) {
                                    case 42:
                                        v166 := []int{32767, genFn1(7, 2147483647, 21), genFn1(v150, -20, 42)}
                                        _ = v166
                                    case -851:
                                        v2 = true
                                    case 128:
                                        fmt.Println(v1, v1)
                                        var v167 int16 = 42
                                        var v168 uint8 = 42
                                        v2 = ((0 <= -522) || ((!v1) && (!true)))
                                        v169 := []int{(fi149 & (fi146 % ((fi149) | 1))), (fi146 - -404)}
                                        switch fi149 {
                                        case 723:
                                            if v2 {
                                                v168 = v168
                                            } else {
                                                v168 += 2
                                                v4[2147483647] = 91
                                            }
                                        case 2147483647:
                                            v170 := GenS19{F0: 710}
                                            _ = v170
                                        }
                                        _ = v167
                                        _ = v168
                                        _ = v169
                                    }
                                    _ = len(v4)
                                    v171 := []int{fi149, (fi149 & fi146), (fi146 - (fi146 - (fi146 - 4294967295)))}
                                    switch ("" + "hello") {
                                    case "a":
                                        v2 = (2 <= 4294967295)
                                    }
                                    _ = v171
                                }
                            }
                            v172 := GenS20{F0: 2}
                            if ((v150 * v150) != 0) {
                                v2 = (-128 <= 1)
                                var v173 uint64 = 42
                                _ = v173
                            } else {
                                v155 -= v155
                                if ((!v2) || v1) {
                                    var v174 int64 = 0
                                    v175 := v3[genFn1((v150 << ((477) & 63)), 0, 128)]
                                    _ = v174
                                    _ = v175
                                } else {
                                    v176 := GenS21{F0: 551, F1: 2}
                                    _ = v176
                                }
                            }
                            var v177 int8 = v155
                            _ = v155
                            _ = v172
                            _ = v177
                        }
                    case 42:
                        var v178 float32 = -2.25
                        _ = v178
                    default:
                        switch "a" {
                        case "a":
                            v150 = 112
                        case "hello":
                            switch (fi146 % ((fi149) | 1)) {
                            case 516:
                                switch genFn1((v150 >> ((0) & 63)), 1, 42) {
                                case -825:
                                    var v179 bool = ((true || v2) || (1 < 4294967295))
                                    _ = v179
                                default:
                                    if (!(v150 >= (v150 << ((0) & 63)))) {
                                        if v2 {
                                            delete(v3, (-(fi149)))
                                            if v2 {
                                                v180 := []int{-2147483648, (fi149 + (^(fi149))), (fi149 ^ 4294967295), fi146}
                                                v2 = v1
                                                var v181 int32 = -396
                                                v182 := v3[127]
                                                _ = v180
                                                _ = v181
                                                _ = v182
                                            }
                                            fmt.Println(fi149, v1)
                                        } else {
                                            v2 = (!false)
                                        }
                                        delete(v3, genFn1((v150 & v150), 1, 128))
                                        v1 = false
                                    } else {
                                        if (2147483647 == -2147483648) {
                                            if v1 {
                                                var v183 int64 = 417
                                                v184 := GenS21{F0: 255, F1: 191}
                                                v3[(fi149 + fi149)] = ("hello" + "foo bar")
                                                _ = v183
                                                _ = v184
                                            }
                                            v150 += v150
                                            var v185 uint = 65535
                                            _ = v185
                                        } else {
                                            var v186 uint16 = 32767
                                            _ = v186
                                        }
                                        delete(v4, 65535)
                                        switch "hello" {
                                        case "foo bar":
                                            v150 -= v150
                                        case "hello":
                                            if v2 {
                                                v187 := GenS22{F0: 1}
                                                v188 := []int{fi146}
                                                _ = v187
                                                _ = v188
                                            } else {
                                                v150 = (v150 + v150)
                                            }
                                        case "":
                                            if v1 {
                                                if (255 != 1) {
                                                    v1 = v1
                                                } else {
                                                    v189 := GenS13{F0: 124}
                                                    switch fi149 {
                                                    case -474:
                                                        if (((1e10 + (((((float64(1e10) / 1e10) - ((0.1 * 1.5) * 0.1)) * 0.1) / 3.14159) * (0.1 * 1.5))) - 1e10) > (0.1 + 1.5)) {
                                                            var v190 uint16 = 10
                                                            v191 := v3[(fi146 - (fi146 - 112))]
                                                            fmt.Println(v190)
                                                            v192 := []int{fi149, -455}
                                                            _ = v190
                                                            _ = v191
                                                            _ = v192
                                                        }
                                                    case -1:
                                                        v193 := GenS9{F0: 2}
                                                        _ = v193
                                                    }
                                                    var v194 string = "hello"
                                                    _ = v189
                                                    _ = v194
                                                }
                                            }
                                        }
                                    }
                                }
                            case -902:
                                v195 := GenS23{F0: v150, F1: 42, F2: 1}
                                _ = v195
                            }
                        default:
                            fmt.Println(v150, v2, v2)
                        }
                    }
                    var v196 int64 = 1
                    _ = v150
                    _ = v196
                } else {
                    v1 = (false && false)
                    v1 = true
                    v2 = v1
                    v2 = v1
                }
                delete(v4, 42)
                switch ("a" + ("" + "\n\t")) {
                case "hello":
                    if (v2 || (4 == 2)) {
                        var v197 int = 42
                        _ = v197
                    } else {
                        fmt.Println(v2, fi149, fi146)
                        v1 = true
                    }
                case "":
                    var v198 uint16 = 32767
                    _ = v198
                }
                v2 = (!true)
                var v199 uint8 = 127
                _ = fi149
                _ = v199
            }
        } else {
            v200 := GenS24{F0: 415}
            if false {
                var v201 uint16 = 0
                fmt.Println(v2)
                _ = v201
            }
            _ = v200
        }
        _ = fi146
    }
    switch "a" {
    case "foo bar":
        switch (("a" + "\n\t") + "") {
        case "foo bar":
            v1 = v2
        default:
            var v202 float32 = ((0.1 - 1.5) / 0.1)
            _ = v202
        }
    case "hello":
        var v203 string = ("" + "a")
        _ = v203
    case "a":
        v2 = (1.5 <= (1.5 / 1.0))
    }
    var v204 uint = 42
    if (0 <= 1) {
        for fi205 := 0; fi205 < 5; fi205++ {
            for fi206 := 0; fi206 < 14; fi206++ {
                v207 := []int{(fi205 * 654), (fi206 / (((fi206 >> ((v204) & 63))) | 1)), fi206}
                delete(v4, 673)
                _ = fi206
                _ = v207
            }
            fmt.Println(v1, v1, v2)
            v208 := GenS25{F0: ("a" + ("" + "foo bar")), F1: 255}
            _ = fi205
            _ = v208
        }
        v209 := []int{825, 0}
        _ = v209
    }
    var v210 float64 = 0.1
    v2 = v2
    fmt.Println(v204, v204, v2)
    fmt.Println(v2, v210, v204)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v204:", v204)
    fmt.Println("v210:", v210)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v204
    _ = v210
    _ = v3
    _ = v4
}
