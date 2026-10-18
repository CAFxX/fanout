package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 uint16; F1 string }
type GenS11 struct { F0 uint64; F1 uint32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 int }
func genFn0(p0 int16, p1 uint) uint16 {
    v4 := GenS1{F0: 1, F1: 42}
    if (false && true) {
        var v5 float64 = 1.5
        _ = v5
    } else {
        p1 = p1
    }
    var v6 uint8 = 2
    v6 += 255
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 590
}

func main() {
    var v1 uint64 = 2
    var v2 float64 = float64(float32(((0.0 + float64((1.5 + (0.1 * (1e10 - 0.1))))) / (3.14159 / 3.14159))))
    v3 := make(map[uint32]int)
    var v7 uint = 1
    v2 = v2
    var v8 int8 = 65
    if true {
        v2 += -2.25
        if (uint32(v7) >= 65535) {
            switch "\n\t" {
            case "":
                v1 = 4294967295
            }
            var v9 int16 = -202
            v10 := []int{32767, -708}
            _ = v10
            _ = v9
        }
        v3[42] = -1
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v7 -= (v7 | v7)
        _ = fi11
    }
    fmt.Println(v7, v7, v7)
    fmt.Println(v7, v8)
    if false {
        switch 4294967295 {
        case -462:
            v12 := []int{42, -9223372036854775808, 9223372036854775807, -1}
            _ = v12
        }
        switch "\n\t" {
        case "hello":
            v13 := GenS1{F0: -1, F1: genFn0(-314, 756)}
            _ = v13
        default:
            for fi14 := 0; fi14 < 17; fi14++ {
                v7 = 65535
                fmt.Println(fi14)
                _ = fi14
            }
        }
        v15 := []int{}
        _ = v15
    } else {
        switch int(v1) {
        case 32768:
            if ((float64((-2.25 * 1.5)) - float64(1.5)) == v2) {
                var v16 uint16 = genFn0(1000, v7)
                fmt.Println(v16)
                _ = v16
            } else {
                switch "" {
                case "":
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v18 := GenS1{F0: -2147483648, F1: 42}
                        _ = fi17
                        _ = v18
                    }
                case "a":
                    v1 -= (^(v1))
                case "hello":
                    if true {
                        if (2 > 128) {
                            fmt.Println(v2, v7, v7)
                            switch "foo bar" {
                            case "foo bar":
                                v19 := GenS2{F0: -137, F1: 32767, F2: v1}
                                _ = v19
                            case "\n\t":
                                v20 := []int{}
                                _ = v20
                            }
                        }
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v1 = v1
                            switch (fi21 / ((fi21) | 1)) {
                            case -788:
                                fmt.Println(v2, v2)
                            case -2147483648:
                                var v22 int = (fi21 / ((9223372036854775807) | 1))
                                _ = v22
                            case 42:
                                if false {
                                    v23 := []int{fi21}
                                    _ = v23
                                } else {
                                    var v24 uint32 = 548
                                    _ = v24
                                }
                            default:
                                v8 = -128
                            }
                            switch "a" {
                            case "foo bar":
                                v3[4294967295] = (fi21 % ((fi21) | 1))
                            case "\n\t":
                                if false {
                                    switch fi21 {
                                    case 42:
                                        var v25 int8 = v8
                                        _ = v25
                                    default:
                                        fmt.Println(v7)
                                    }
                                    fmt.Println(v1)
                                } else {
                                    if false {
                                        v2 = 1e10
                                    } else {
                                        v26 := GenS3{F0: (true && false), F1: (v8 == 100)}
                                        _ = v26
                                    }
                                    v8 = v8
                                }
                            case "a":
                                v2 -= 1e10
                            default:
                                var v27 int8 = 100
                                _ = v27
                            }
                            _ = fi21
                        }
                        var v28 uint8 = 7
                        _ = v28
                    } else {
                        var v29 int = 32767
                        switch (-(v29)) {
                        case 244:
                            if true {
                                var v30 uint16 = 1
                                _ = v30
                            } else {
                                var v31 float64 = 1e10
                                v7 = (^(v7))
                                _ = v31
                            }
                        case 1:
                            for fi32 := 0; fi32 < 19; fi32++ {
                                var v33 bool = (v1 >= 4294967295)
                                for fi34 := 0; fi34 < 11; fi34++ {
                                    v33 = (!(int64(v7) < 9223372036854775807))
                                    _ = fi34
                                }
                                if v33 {
                                    v2 += 3.14159
                                    v7 -= (v7 / (((^(v7))) | 1))
                                }
                                _ = fi32
                                _ = v33
                            }
                        }
                        _ = v29
                    }
                default:
                    v7 = v7
                }
                v35 := GenS3{F0: true, F1: (true || (true || true))}
                _ = v35
            }
        case -9223372036854775808:
            var v36 int = -2147483648
            _ = v36
        case 4294967295:
            if true {
                v7 -= (v7 & v7)
                delete(v3, 811)
                var v37 float32 = 1e10
                _ = v37
            }
        }
    }
    var v38 float32 = float32(v2)
    if (((v7 % ((4294967295) | 1)) != 42) || (198 != 2)) {
        v8 = v8
        for fi39 := 0; fi39 < 14; fi39++ {
            v1 += 0
            var v40 float64 = v2
            _ = fi39
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 9; fi41++ {
            for fi42 := 0; fi42 < 8; fi42++ {
                v43 := v3[42]
                fmt.Println(v43)
                var v44 int64 = -1
                v43 = (fi41 / ((v43) | 1))
                v45 := GenS4{F0: false, F1: (v1 << ((18446744073709551615) & 63))}
                v46 := []int{}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
                _ = v46
            }
            v7 = 0
            _ = fi41
        }
        fmt.Println(v38, v8)
    }
    v47 := GenS2{F0: 645, F1: 32767, F2: (^(v1))}
    var v48 uint64 = (v1 >> ((v7) & 63))
    v49 := GenS2{F0: 172, F1: 1, F2: (v48 | (v48 + 1))}
    for fi50 := 0; fi50 < 5; fi50++ {
        for fi51 := 0; fi51 < 19; fi51++ {
            var v52 bool = false
            if false {
                v53 := GenS4{F0: (((v2 / -2.25) + v2) <= float64(v38)), F1: (v48 % ((v48) | 1))}
                _ = v53
            } else {
                v54 := GenS3{F0: (true || false), F1: (-1 < 42)}
                fmt.Println(v7)
                switch 32767 {
                case -2147483648:
                    if (false && ((v2 != (3.14159 / float64(0.1))) || (float32(((v2 + v2) * (v2 / v2))) > (0.0 / (1e10 * v38))))) {
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            switch "" {
                            case "foo bar":
                                if v52 {
                                    if ("" > ("foo bar" + "\n\t")) {
                                        if v52 {
                                            if true {
                                                v55 := []int{}
                                                switch (fi51 / (((fi50 | fi50)) | 1)) {
                                                case -493:
                                                    if (1 > 1) {
                                                        var v56 uint16 = 2
                                                        v7 = (v7 >> (((^(v7))) & 63))
                                                        var v57 uint32 = 2147483647
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                case 32767:
                                                    var v58 float32 = v38
                                                    _ = v58
                                                case -1:
                                                    v55 = append(v55, (-(fi51)))
                                                    v2 += v2
                                                default:
                                                    v2 += 0.0
                                                }
                                                v59 := GenS4{F0: (139 > 0), F1: v1}
                                                _ = v55
                                                _ = v59
                                            }
                                            if (true || (false && (((fi51 - (fi50 / ((fi50) | 1))) == (-(fi50))) || v52))) {
                                                v60 := GenS4{F0: true, F1: (^(v48))}
                                                v8 += v8
                                                var v61 uint = v7
                                                _ = v60
                                                _ = v61
                                            }
                                            var v62 uint32 = 345
                                            _ = v62
                                        } else {
                                            var v63 uint16 = 255
                                            _ = v63
                                        }
                                        if false {
                                            v3[2147483648] = int(v8)
                                        }
                                        if v52 {
                                            if (false && (((v8 << ((v7) & 63)) > 100) || ((!(128 < 42)) || false))) {
                                                if (v52 && v52) {
                                                    v64 := GenS5{F0: v8, F1: true}
                                                    _ = v64
                                                } else {
                                                    var v65 int16 = -206
                                                    v66 := GenS4{F0: ((((float32((v2 - (1.5 / 1e10))) - v38) * v38) + (v38 + v38)) <= 0.0), F1: v48}
                                                    _ = v65
                                                    _ = v66
                                                }
                                            } else {
                                                switch (^(fi51)) {
                                                case 42:
                                                    v67 := v3[1]
                                                    _ = v67
                                                case 4294967295:
                                                    if (!true) {
                                                        v38 += (v38 - 0.0)
                                                        if (((v1 >> ((18446744073709551615) & 63)) <= 0) || ((!v52) && true)) {
                                                            v2 = v2
                                                            v8 -= (v8 * -1)
                                                        }
                                                        v68 := []int{}
                                                        v38 = 3.14159
                                                        v68 = append(v68, fi51)
                                                        _ = v68
                                                    }
                                                default:
                                                    v69 := []int{fi50, fi50, fi50}
                                                    _ = v69
                                                }
                                                var v70 float32 = 3.14159
                                                _ = v70
                                            }
                                            v3[4294967295] = 4294967295
                                            if ((-1 != 107) || (1 < 2147483647)) {
                                                v8 = v8
                                                switch (fi50 >> ((v7) & 63)) {
                                                case 42:
                                                    switch (fi50 * (fi51 << ((v7) & 63))) {
                                                    case 32767:
                                                        var v71 uint16 = 658
                                                        _ = v71
                                                    }
                                                case -1:
                                                    if (781 < 1000) {
                                                        var v72 float32 = v38
                                                        var v73 uint32 = 65535
                                                        v74 := []int{(^(fi51)), (fi51 >> ((uint(v1)) & 63)), fi50, (^(fi50))}
                                                        _ = v72
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        if (!v52) {
                                                            var v75 int8 = 100
                                                            if (v52 && v52) {
                                                                var v76 uint32 = 542
                                                                var v77 int32 = 745
                                                                v78 := GenS5{F0: v8, F1: (!(v52 || (float32(v2) != float32((1e10 / v2)))))}
                                                                _ = v76
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            _ = v75
                                                        }
                                                        v79 := GenS2{F0: 0, F1: 67, F2: (v48 & v48)}
                                                        v7 -= (v7 * 261)
                                                        _ = v79
                                                    }
                                                case -9223372036854775808:
                                                    fmt.Println(v52)
                                                }
                                                var v80 int8 = v8
                                                v3[945] = fi50
                                                _ = v80
                                            } else {
                                                var v81 int32 = -350
                                                _ = v81
                                            }
                                            v8 = (-(v8))
                                            switch "hello" {
                                            case "foo bar":
                                                switch 1 {
                                                case -725:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v1 -= v1
                                                        v2 += float64((v38 / ((((v38 * v38) / (float32(((v2 * ((v2 * v2) + (1e10 + v2))) / v2)) / v38)) / v38) + 0.0)))
                                                        switch ("\n\t" + "\n\t") {
                                                        case "foo bar":
                                                            if v52 {
                                                                v1 = (v1 - (^(v48)))
                                                                v1 += v48
                                                            }
                                                        case "a":
                                                            var v82 uint8 = uint8(v7)
                                                            _ = v82
                                                        case "\n\t":
                                                            v38 -= v38
                                                        }
                                                        if true {
                                                            v8 += v8
                                                            v8 = 86
                                                        } else {
                                                            switch 2147483647 {
                                                            case 32767:
                                                                var v83 uint64 = 1
                                                                _ = v83
                                                            case -2147483648:
                                                                if v52 {
                                                                    if (32767 != -345) {
                                                                        fmt.Println(v1, v52)
                                                                        if (!true) {
                                                                            var v84 int = (fi51 / ((-2147483648) | 1))
                                                                            v85 := GenS2{F0: 1, F1: 32767, F2: (v48 / ((12) | 1))}
                                                                            if (123 == 127) {
                                                                                v7 = 2
                                                                            }
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                    }
                                                                    v8 = 0
                                                                    var v86 uint64 = 805
                                                                    _ = v86
                                                                }
                                                            default:
                                                                switch ("\n\t" + "foo bar") {
                                                                case "a":
                                                                    if v52 {
                                                                        fmt.Println(v48)
                                                                    } else {
                                                                        if (42 > 0) {
                                                                            if v52 {
                                                                                if (true || (true && false)) {
                                                                                    var v87 int16 = 770
                                                                                    _ = v87
                                                                                }
                                                                                v2 = (0.0 / v2)
                                                                            } else {
                                                                                v8 -= v8
                                                                                v88 := GenS5{F0: (v8 - 41), F1: (42 > -601)}
                                                                                _ = v88
                                                                            }
                                                                            if v52 {
                                                                                var v89 string = ""
                                                                                v90 := GenS6{F0: -2147483648, F1: v89}
                                                                                v89 = ""
                                                                                _ = v89
                                                                                _ = v90
                                                                            }
                                                                        } else {
                                                                            v91 := GenS6{F0: 116, F1: "hello"}
                                                                            v48 -= (^(v1))
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                switch ("" + (("" + ("a" + "a")) + ("hello" + "a"))) {
                                                                                case "a":
                                                                                    if ((!(1 >= int16(v7))) || false) {
                                                                                        if (1 != 9223372036854775807) {
                                                                                            switch fi50 {
                                                                                            case -379:
                                                                                                v1 = 1
                                                                                            case -2147483648:
                                                                                                v92 := GenS7{F0: (v2 == (((v2 * v2) * ((v2 * v2) * (v2 / (float64((v38 - (v38 - v38))) - v2)))) + float64(v38)))}
                                                                                                _ = v92
                                                                                            case 32767:
                                                                                                var v93 int = (fi51 | 0)
                                                                                                _ = v93
                                                                                            default:
                                                                                                if v52 {
                                                                                                    v94 := GenS4{F0: v52, F1: (^(v48))}
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    var v95 float32 = 1e10
                                                                                                    _ = v95
                                                                                                }
                                                                                            }
                                                                                            v96 := GenS1{F0: -351, F1: 2}
                                                                                            v97 := GenS8{F0: (v1 & v48), F1: -1}
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        } else {
                                                                                            v98 := []int{}
                                                                                            if len(v98) >= 2 {
                                                                                                v98 = v98[1:len(v98)]
                                                                                            }
                                                                                            _ = v98
                                                                                        }
                                                                                        v8 = v8
                                                                                        var v99 uint16 = 511
                                                                                        _ = v99
                                                                                    } else {
                                                                                        v100 := GenS9{F0: v7, F1: 65535}
                                                                                        if true {
                                                                                            var v101 float64 = v2
                                                                                            fmt.Println(v101)
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                var v102 uint8 = 0
                                                                                                _ = v102
                                                                                            case "hello":
                                                                                                var v103 float32 = ((float32(0.0) - (v38 / v38)) * (((v38 / v38) * v38) / float32((float64((3.14159 - v38)) - (v2 + v101)))))
                                                                                                _ = v103
                                                                                            }
                                                                                            _ = v101
                                                                                        }
                                                                                        v104 := []int{}
                                                                                        _ = v100
                                                                                        _ = v104
                                                                                    }
                                                                                case "\n\t":
                                                                                    v105 := GenS10{F0: 42, F1: ""}
                                                                                    _ = v105
                                                                                default:
                                                                                    var v106 int16 = -32768
                                                                                    _ = v106
                                                                                }
                                                                            case "hello":
                                                                                if v52 {
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        var v107 uint8 = 255
                                                                                        _ = v107
                                                                                    case "\n\t":
                                                                                        v108 := v3[2147483648]
                                                                                        v109 := GenS1{F0: int32(v1), F1: 1}
                                                                                        if (v52 && v52) {
                                                                                            var v110 bool = (!true)
                                                                                            _ = v110
                                                                                        }
                                                                                        _ = v108
                                                                                        _ = v109
                                                                                    case "":
                                                                                        var v111 uint8 = 128
                                                                                        _ = v111
                                                                                    }
                                                                                    if (4294967295 > 969) {
                                                                                        var v112 int32 = -413
                                                                                        _ = v112
                                                                                    } else {
                                                                                        delete(v3, 2147483648)
                                                                                        var v113 int64 = 0
                                                                                        _ = v113
                                                                                    }
                                                                                }
                                                                            default:
                                                                                fmt.Println(fi50, v1, v38)
                                                                            }
                                                                            _ = v91
                                                                        }
                                                                        v1 += 2
                                                                        if (v38 == 1.5) {
                                                                            v114 := []int{fi51, fi51, 273}
                                                                            _ = len(v114)
                                                                            v115 := GenS3{F0: v52, F1: false}
                                                                            _ = v114
                                                                            _ = v115
                                                                        } else {
                                                                            var v116 bool = ("" != (("\n\t" + "foo bar") + ""))
                                                                            v117 := []int{(fi51 - 4294967295), fi50}
                                                                            v118 := GenS11{F0: 4294967295, F1: 1}
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                        }
                                                                        v48 = (^(v48))
                                                                        v48 = (v48 & (v48 & (^(v1))))
                                                                    }
                                                                case "\n\t":
                                                                    if (((!(false || (1000 >= 1))) && false) || false) {
                                                                        v1 += (^(v1))
                                                                        v119 := []int{fi50, -2147483648}
                                                                        _ = v119
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if v52 {
                                                            var v120 uint = (v7 << (((v7 ^ v7)) & 63))
                                                            v8 -= (v8 | v8)
                                                            v121 := GenS5{F0: (-(v8)), F1: (32767 < 32767)}
                                                            _ = v120
                                                            _ = v121
                                                        } else {
                                                            var v122 int32 = 42
                                                            _ = v122
                                                        }
                                                    case "":
                                                        v123 := GenS12{F0: ((!v52) || v52)}
                                                        _ = v123
                                                    case "hello":
                                                        v124 := []int{fi51, fi51, -491}
                                                        _ = v124
                                                    default:
                                                        var v125 uint64 = v1
                                                        _ = v125
                                                    }
                                                case -9223372036854775808:
                                                    v1 = (^(v1))
                                                case -1:
                                                    v126 := GenS2{F0: 9223372036854775807, F1: 65535, F2: v48}
                                                    _ = v126
                                                default:
                                                    v127 := GenS5{F0: (-(v8)), F1: (42 > genFn0(-558, v7))}
                                                    _ = v127
                                                }
                                            case "\n\t":
                                                switch fi51 {
                                                case 32768:
                                                    v48 += 430
                                                case 627:
                                                    switch "foo bar" {
                                                    case "":
                                                        v2 -= 1.5
                                                    case "a":
                                                        v128 := []int{}
                                                        _ = v128
                                                    case "foo bar":
                                                        v129 := []int{fi51}
                                                        _ = v129
                                                    default:
                                                        var v130 uint16 = 665
                                                        _ = v130
                                                    }
                                                case 72:
                                                    var v131 uint8 = 1
                                                    _ = len(v3)
                                                    if (780 >= 2) {
                                                        v132 := []int{4294967295}
                                                        if len(v132) > 0 {
                                                            v132[0] = (fi51 >> (((v7 / ((v7) | 1))) & 63))
                                                        }
                                                        if v52 {
                                                            var v133 int64 = -1
                                                            _ = v133
                                                        }
                                                        _ = v132
                                                    } else {
                                                        var v134 uint64 = (v48 / ((v48) | 1))
                                                        v135 := []int{(^(fi50)), fi50, fi50}
                                                        _ = v134
                                                        _ = v135
                                                    }
                                                    _ = v131
                                                }
                                            }
                                        }
                                        switch (("foo bar" + "\n\t") + ("hello" + ("\n\t" + (((("hello" + "") + "hello") + ("" + "")) + ("\n\t" + "\n\t"))))) {
                                        case "foo bar":
                                            var v136 string = "foo bar"
                                            _ = v136
                                        case "a":
                                            var v137 float32 = float32(float64(v38))
                                            _ = v137
                                        default:
                                            v138 := []int{fi50, 128}
                                            _ = v138
                                        }
                                    }
                                }
                            }
                        case "":
                            var v139 bool = v52
                            _ = v139
                        case "foo bar":
                            v140 := []int{-307, fi51, (fi50 / (((fi51 / (((^(fi51))) | 1))) | 1))}
                            _ = v140
                        }
                    } else {
                        if v52 {
                            v141 := GenS12{F0: (false || v52)}
                            _ = v141
                        }
                        _ = len(v3)
                        fmt.Println(v7, v48, fi50)
                    }
                default:
                    v2 += float64(3.14159)
                }
                v142 := []int{fi51, 2147483647}
                _ = v142
                _ = v54
            }
            v2 = float64(v38)
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    if (!(-1 != (v8 ^ (v8 / ((v8) | 1))))) {
        v143 := []int{127, 128, 883}
        v1 = uint64(v7)
        _ = v143
    }
    v48 = uint64(v7)
    for fi144 := 0; fi144 < 14; fi144++ {
        fmt.Println(fi144, v7)
        if false {
            for fi145 := 0; fi145 < 10; fi145++ {
                v38 = 3.14159
                _ = fi145
            }
        } else {
            if true {
                v48 = (v1 % ((v1) | 1))
            }
            for fi146 := 0; fi146 < 19; fi146++ {
                var v147 uint16 = 0
                if false {
                    var v148 bool = false
                    var v149 uint64 = uint64(v147)
                    switch "foo bar" {
                    case "hello":
                        v150 := v3[42]
                        _ = v150
                    case "\n\t":
                        var v151 uint32 = 4294967295
                        _ = v151
                    }
                    if v148 {
                        var v152 string = "hello"
                        _ = v152
                    } else {
                        v8 += (v8 ^ (v8 >> ((uint(v147)) & 63)))
                        v1 = v48
                        v153 := GenS13{F0: 0, F1: -431}
                        _ = v153
                    }
                    var v154 int32 = -669
                    _ = v148
                    _ = v149
                    _ = v154
                } else {
                    v7 = uint(v48)
                }
                if (v147 != 255) {
                    if false {
                        var v155 float64 = float64((1.5 + v38))
                        v156 := GenS4{F0: (((v155 + v2) + 0.1) < (0.0 / v2)), F1: 734}
                        v7 = v7
                        _ = v155
                        _ = v156
                    }
                } else {
                    if true {
                        v157 := []int{(^(fi144))}
                        _ = v157
                    }
                }
                _ = fi146
                _ = v147
            }
        }
        _ = fi144
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v48:", v48)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v47
    _ = v48
    _ = v49
    _ = v7
    _ = v8
}
