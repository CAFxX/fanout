package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int32; F2 int64 }
type GenS9 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS10 struct { F0 uint16; F1 int8; F2 bool }
type GenS11 struct { F0 uint8; F1 int16 }
type GenS12 struct { F0 uint64; F1 uint8; F2 string }
type GenS13 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS14 struct { F0 int64; F1 uint8 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int16; F1 uint16 }
type GenS17 struct { F0 uint; F1 uint; F2 string }
type GenS18 struct { F0 int; F1 int64 }
type GenS19 struct { F0 uint16; F1 string }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0() uint {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int32 = 42
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(fi7, fi7, v6)
            switch (fi7 % (((fi5 / (((^(fi5))) | 1))) | 1)) {
            case -2147483648:
                switch 854 {
                case 100:
                    var v8 bool = (int(v6) != 4294967295)
                    _ = v8
                case 9223372036854775807:
                    v9 := []int{fi5, 9223372036854775807}
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    fmt.Println(fi5, fi7, fi5)
                    _ = v9
                case 338:
                    v10 := GenS1{F0: 694, F1: (807 == 1)}
                    _ = v10
                }
            }
            fmt.Println(v6, v6, fi5)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    v11 := GenS2{F0: 215}
    _ = v11
    return 788
}

func main() {
    var v1 int = 127
    v2 := []int{-68, 635, v1, 42}
    var v3 uint16 = 495
    var v4 string = ("" + ("foo bar" + ("a" + "foo bar")))
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(fi13, v3)
            v14 := GenS1{F0: 1, F1: (!(0.0 != (0.0 * ((1.5 * 0.0) * (0.0 - (1.5 * float32((-2.25 * 3.14159))))))))}
            var v15 string = ""
            switch 0 {
            case -881:
                if len(v2) > 0 {
                    v2[0] = -765
                }
            case -178:
                var v16 uint8 = 1
                _ = v16
            default:
                switch "" {
                case "hello":
                    if false {
                        var v17 int8 = 92
                        v18 := GenS3{F0: 276}
                        _ = v17
                        _ = v18
                    }
                case "\n\t":
                    switch "" {
                    case "":
                        var v19 float32 = 0.1
                        _ = v19
                    case "a":
                        fmt.Println(fi12, v4, fi12)
                    default:
                        v20 := GenS2{F0: 18446744073709551615}
                        _ = v20
                    }
                case "foo bar":
                    if true {
                        if true {
                            fmt.Println(v4, v1, v15)
                            if true {
                                v21 := GenS1{F0: 65535, F1: (1e10 > 0.1)}
                                _ = v21
                            }
                            switch ((v4 + "a") + "a") {
                            case "hello":
                                if (!false) {
                                    if false {
                                        fmt.Println(v3, v3)
                                        if (fi12 == v1) {
                                            if true {
                                                v15 = v4
                                                var v22 uint8 = 103
                                                _ = v22
                                            } else {
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v23 float64 = 1.5
                                                    v23 -= ((v23 + v23) + v23)
                                                    _ = v23
                                                }
                                            }
                                            switch (fi13 & (fi12 ^ v1)) {
                                            case 32767:
                                                if (((!((^(fi12)) >= -496)) || (true || ((v3 + v3) >= (v3 >> ((genFn0()) & 63))))) && true) {
                                                    switch (v4 + "") {
                                                    case "":
                                                        if false {
                                                            if len(v2) > 0 {
                                                                v2[0] = fi13
                                                            }
                                                            v1 = v1
                                                        }
                                                    }
                                                    v24 := GenS1{F0: 65535, F1: (false || true)}
                                                    if (0 != -674) {
                                                        v1 = fi12
                                                        var v25 uint16 = 475
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v25
                                                    }
                                                    _ = v24
                                                }
                                            case 42:
                                                var v26 uint16 = 255
                                                _ = v26
                                            case -98:
                                                v1 -= 9223372036854775807
                                            default:
                                                fmt.Println(v1)
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                    v1 += (v1 / ((fi13) | 1))
                                } else {
                                    var v27 uint8 = 0
                                    var v28 uint32 = 1
                                    v29 := GenS2{F0: 18446744073709551615}
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                fmt.Println(fi13)
                            }
                            if (!false) {
                                var v30 uint64 = 0
                                _ = v30
                            } else {
                                var v31 uint16 = v3
                                v31 = (v3 >> ((genFn0()) & 63))
                                if false {
                                    v31 = 0
                                    if len(v2) > 0 {
                                        v2[0] = -2147483648
                                    }
                                    fmt.Println(v4, v3, v4)
                                } else {
                                    var v32 int8 = 1
                                    _ = v32
                                }
                                _ = v31
                            }
                            v15 = v4
                        } else {
                            if true {
                                var v33 int = fi12
                                if (false || (!(42 > 2))) {
                                    switch 2147483647 {
                                    case 727:
                                        if false {
                                            if (!(!(1000 < 42))) {
                                                _ = len(v2)
                                                v1 -= (fi13 / (((v1 & (v33 & (v1 | fi12)))) | 1))
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            } else {
                                                fmt.Println(fi13, v3, v3)
                                                var v34 int64 = 9223372036854775807
                                                switch (^(fi12)) {
                                                case -591:
                                                    v1 = (fi13 | fi13)
                                                }
                                                _ = v34
                                            }
                                            _ = len(v2)
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = fi13
                                            }
                                        }
                                    case 128:
                                        var v35 float64 = -2.25
                                        _ = v35
                                    default:
                                        var v36 int64 = 9223372036854775807
                                        _ = v36
                                    }
                                    v15 = ""
                                    v4 = (v15 + v15)
                                }
                                v37 := GenS1{F0: 4294967295, F1: ((v15 > v15) || ((!(!false)) && true))}
                                _ = v33
                                _ = v37
                            }
                            switch "a" {
                            case "hello":
                                switch (v15 + v15) {
                                case "hello":
                                    v4 = ""
                                case "\n\t":
                                    if (9223372036854775807 == -9223372036854775808) {
                                        var v38 int = (fi13 * (^(fi12)))
                                        _ = v38
                                    }
                                }
                            case "\n\t":
                                fmt.Println(fi12, v15, v3)
                            default:
                                v3 = 2
                            }
                            v39 := GenS2{F0: 1}
                            _ = v39
                        }
                    }
                default:
                    v40 := GenS3{F0: 0}
                    _ = v40
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        var v41 int = fi12
        _ = fi12
        _ = v41
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = "\n\t"
    if (true && ((^(v1)) != v1)) {
        if true {
            if true {
                v42 := GenS3{F0: 0}
                if (!((-532 < 42) && (18446744073709551615 <= 4294967295))) {
                    if false {
                        switch v4 {
                        case "hello":
                            for fi43 := 0; fi43 < 6; fi43++ {
                                switch -837 {
                                case 0:
                                    v4 = ("\n\t" + "")
                                case 32768:
                                    for fi44 := 0; fi44 < 10; fi44++ {
                                        v45 := GenS2{F0: 599}
                                        fmt.Println(v1, fi44, v4)
                                        _ = fi44
                                        _ = v45
                                    }
                                default:
                                    if (!true) {
                                        for fi46 := 0; fi46 < 1; fi46++ {
                                            v1 -= fi43
                                            _ = fi46
                                        }
                                    } else {
                                        v2 = append(v2, (v1 - int(v3)))
                                        if false {
                                            for fi47 := 0; fi47 < 14; fi47++ {
                                                v48 := GenS4{F0: (fi43 != (v1 - v1))}
                                                fmt.Println(fi43)
                                                _ = fi47
                                                _ = v48
                                            }
                                            for fi49 := 0; fi49 < 4; fi49++ {
                                                v4 = v4
                                                v3 = v3
                                                var v50 uint8 = 138
                                                _ = fi49
                                                _ = v50
                                            }
                                            switch 32768 {
                                            case 9223372036854775807:
                                                for fi51 := 0; fi51 < 5; fi51++ {
                                                    v52 := GenS4{F0: false}
                                                    if true {
                                                        if false {
                                                            if true {
                                                                v53 := GenS2{F0: 4294967295}
                                                                v54 := GenS5{F0: 1, F1: 737}
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            v1 += -1
                                                        }
                                                        var v55 uint64 = 256
                                                        _ = v55
                                                    }
                                                    _ = fi51
                                                    _ = v52
                                                }
                                            case 127:
                                                var v56 uint8 = 1
                                                _ = v56
                                            default:
                                                v4 = v4
                                            }
                                        } else {
                                            var v57 uint32 = 1
                                            for fi58 := 0; fi58 < 8; fi58++ {
                                                var v59 uint16 = 2
                                                var v60 uint16 = 255
                                                _ = fi58
                                                _ = v59
                                                _ = v60
                                            }
                                            v61 := GenS1{F0: v57, F1: (true && (!true))}
                                            _ = v57
                                            _ = v61
                                        }
                                        fmt.Println(v1)
                                    }
                                }
                                for fi62 := 0; fi62 < 10; fi62++ {
                                    if (32767 > 32767) {
                                        v1 += (fi43 - 1)
                                        v63 := GenS6{F0: 169, F1: 255}
                                        v64 := GenS2{F0: 565}
                                        var v65 string = ("hello" + (v4 + v4))
                                        if (-2.25 > ((float32(1e10) - 0.0) - 0.1)) {
                                            var v66 uint16 = 1
                                            switch "\n\t" {
                                            case "hello":
                                                v1 = fi43
                                            case "foo bar":
                                                var v67 float32 = 1.5
                                                _ = v67
                                            case "a":
                                                switch (v1 >> ((0) & 63)) {
                                                case 128:
                                                    if len(v2) > 0 {
                                                        v2[0] = -1
                                                    }
                                                case -222:
                                                    var v68 uint16 = 32767
                                                    _ = v68
                                                default:
                                                    fmt.Println(v3)
                                                }
                                            }
                                            _ = v66
                                        }
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    } else {
                                        var v69 int16 = 462
                                        v2 = append(v2, (fi62 * v1))
                                        var v70 float64 = 3.14159
                                        _ = v69
                                        _ = v70
                                    }
                                    var v71 uint16 = (v3 / ((v3) | 1))
                                    var v72 uint = 1
                                    _ = fi62
                                    _ = v71
                                    _ = v72
                                }
                                _ = fi43
                            }
                        }
                        var v73 int16 = -1
                        for fi74 := 0; fi74 < 8; fi74++ {
                            for fi75 := 0; fi75 < 17; fi75++ {
                                fmt.Println(fi74, fi75)
                                v76 := GenS5{F0: 1, F1: 0}
                                switch v4 {
                                case "foo bar":
                                    switch fi75 {
                                    case 426:
                                        var v77 uint32 = 1
                                        _ = v77
                                    case -770:
                                        switch (fi74 | fi75) {
                                        case 128:
                                            v78 := GenS1{F0: 788, F1: false}
                                            _ = v78
                                        case 32768:
                                            _ = len(v2)
                                        default:
                                            if false {
                                                var v79 uint = 1
                                                _ = v79
                                            } else {
                                                switch v1 {
                                                case 4294967295:
                                                    if (577 != genFn0()) {
                                                        if true {
                                                            var v80 float32 = 1.5
                                                            fmt.Println(v1, v73, fi74)
                                                            _ = v80
                                                        } else {
                                                            v81 := GenS3{F0: (v73 * v73)}
                                                            var v82 float64 = ((3.14159 - 0.1) + 0.1)
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                        fmt.Println(v73, fi74)
                                                        v83 := GenS2{F0: 1}
                                                        if true {
                                                            var v84 bool = true
                                                            var v85 uint64 = 18446744073709551615
                                                            v1 += 2147483647
                                                            v73 += v73
                                                            v85 = (^(v85))
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = (v1 + (fi75 & 32768))
                                                        }
                                                        _ = v83
                                                    } else {
                                                        v73 = (v73 | (v73 + 1))
                                                        v1 = fi75
                                                        var v86 uint64 = 299
                                                        _ = v86
                                                    }
                                                case -260:
                                                    if true {
                                                        var v87 uint64 = 4294967295
                                                        var v88 int32 = 200
                                                        switch v1 {
                                                        case -9223372036854775808:
                                                            if false {
                                                                v89 := GenS3{F0: v73}
                                                                var v90 string = "\n\t"
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                if false {
                                                                    var v91 int64 = 0
                                                                    fmt.Println(v88, v1, v87)
                                                                    _ = v91
                                                                } else {
                                                                    var v92 uint16 = (^(v3))
                                                                    fmt.Println(v87, fi74, v73)
                                                                    _ = v92
                                                                }
                                                                fmt.Println(v87)
                                                                var v93 uint8 = 199
                                                                _ = v93
                                                            }
                                                        case 32768:
                                                            if ((true && (!false)) || false) {
                                                                if (4294967295 != 42) {
                                                                    if (2 >= 853) {
                                                                        switch (v4 + "\n\t") {
                                                                        case "foo bar":
                                                                            v88 += (v88 << ((218) & 63))
                                                                        default:
                                                                            fmt.Println(v73)
                                                                        }
                                                                    } else {
                                                                        var v94 int32 = v88
                                                                        var v95 uint16 = (v3 / ((v3) | 1))
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                    v96 := GenS7{F0: "hello"}
                                                                    _ = v96
                                                                }
                                                                v4 = ("\n\t" + (v4 + "foo bar"))
                                                                v97 := GenS8{F0: ("\n\t" <= ((v4 + (v4 + ("foo bar" + v4))) + (((((v4 + v4) + v4) + v4) + ("a" + v4)) + "\n\t"))), F1: (v88 & (v88 + v88)), F2: int64(v88)}
                                                                _ = v97
                                                            } else {
                                                                v98 := GenS9{F0: 270, F1: (^(v87)), F2: 832}
                                                                _ = v98
                                                            }
                                                        default:
                                                            v3 = v3
                                                        }
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        if (-1 < -128) {
                                                            var v99 string = v4
                                                            switch v99 {
                                                            case "":
                                                                v4 = v99
                                                            case "hello":
                                                                var v100 uint = 125
                                                                _ = v100
                                                            default:
                                                                var v101 string = v4
                                                                _ = v101
                                                            }
                                                            v102 := GenS2{F0: 2}
                                                            _ = v102
                                                            _ = v99
                                                        }
                                                        var v103 uint = 583
                                                        if true {
                                                            var v104 uint64 = 18446744073709551615
                                                            if true {
                                                                v105 := GenS10{F0: v3, F1: 1, F2: (float32((0.0 * (0.0 + (0.0 + 0.1)))) != 3.14159)}
                                                                _ = v105
                                                            } else {
                                                                var v106 uint32 = 2
                                                                _ = v106
                                                            }
                                                            _ = v104
                                                        } else {
                                                            switch v4 {
                                                            case "foo bar":
                                                                v1 = -2147483648
                                                            case "":
                                                                var v107 bool = (true || (168 >= -1))
                                                                _ = v107
                                                            case "\n\t":
                                                                fmt.Println(v3)
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                        }
                                                        if true {
                                                            if false {
                                                                var v108 bool = (2 > 324)
                                                                switch "a" {
                                                                case "":
                                                                    v109 := GenS7{F0: v4}
                                                                    _ = v109
                                                                case "foo bar":
                                                                    var v110 uint64 = 2
                                                                    _ = v110
                                                                }
                                                                fmt.Println(v108, v1)
                                                                _ = v108
                                                            }
                                                            switch v4 {
                                                            case "\n\t":
                                                                v73 -= v73
                                                            case "foo bar":
                                                                var v111 string = v4
                                                                _ = v111
                                                            default:
                                                                v112 := GenS11{F0: 2, F1: v73}
                                                                _ = v112
                                                            }
                                                        } else {
                                                            var v113 int64 = int64(v3)
                                                            _ = v113
                                                        }
                                                        v73 += v73
                                                        var v114 string = (v4 + v4)
                                                        _ = v103
                                                        _ = v114
                                                    }
                                                    if true {
                                                        v3 = v3
                                                        v115 := GenS1{F0: 2, F1: (((65535 >= 65535) || ((3.14159 > (0.1 / -2.25)) || (-4 == 0))) || (-629 < (v73 & 42)))}
                                                        _ = v115
                                                    }
                                                    v3 -= 863
                                                    v3 += (^(v3))
                                                    var v116 uint = genFn0()
                                                    var v117 int8 = 0
                                                    _ = v116
                                                    _ = v117
                                                case -421:
                                                    v118 := GenS8{F0: true, F1: -342, F2: -484}
                                                    _ = v118
                                                default:
                                                    var v119 uint32 = 628
                                                    _ = v119
                                                }
                                            }
                                        }
                                    default:
                                        v120 := GenS7{F0: "hello"}
                                        _ = v120
                                    }
                                default:
                                    if false {
                                        if (v73 <= (v73 * v73)) {
                                            var v121 bool = ((!((((0.1 * (0.0 / (-2.25 - ((-2.25 / (-2.25 * 1.5)) / 1.0)))) != -2.25) || false) && (!true))) && true)
                                            _ = v121
                                        } else {
                                            v122 := GenS12{F0: 0, F1: 2, F2: v4}
                                            var v123 int = (fi74 >> ((0) & 63))
                                            fmt.Println(v1, v3)
                                            _ = v122
                                            _ = v123
                                        }
                                    } else {
                                        switch 491 {
                                        case 196:
                                            if (3.14159 == float32(-2.25)) {
                                                fmt.Println(v4, fi75)
                                            }
                                        default:
                                            var v124 int8 = 0
                                            _ = v124
                                        }
                                        v125 := GenS13{F0: 0, F1: 4294967295, F2: -9223372036854775808}
                                        if (((0.0 - 3.14159) * (3.14159 + ((((1.5 - 1e10) * -2.25) / -2.25) - 0.1))) <= (0.1 - 1e10)) {
                                            var v126 int64 = -1
                                            _ = v126
                                        } else {
                                            v2 = append(v2, v1)
                                        }
                                        _ = v125
                                    }
                                }
                                _ = fi75
                                _ = v76
                            }
                            _ = fi74
                        }
                        _ = v73
                    } else {
                        var v127 int32 = 968
                        switch "a" {
                        case "":
                            var v128 uint = 537
                            _ = v128
                        case "hello":
                            var v129 bool = (!false)
                            _ = v129
                        case "\n\t":
                            if (!(746 != 1)) {
                                if false {
                                    for fi130 := 0; fi130 < 3; fi130++ {
                                        switch (v1 + (fi130 * (v1 | -640))) {
                                        case 42:
                                            switch v4 {
                                            case "hello":
                                                if true {
                                                    switch "a" {
                                                    case "":
                                                        for fi131 := 0; fi131 < 15; fi131++ {
                                                            v132 := GenS7{F0: "foo bar"}
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v133 := GenS7{F0: ""}
                                                            _ = fi131
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                    case "hello":
                                                        v134 := GenS14{F0: -605, F1: 128}
                                                        _ = v134
                                                    }
                                                    var v135 uint16 = v3
                                                    _ = v135
                                                } else {
                                                    _ = len(v2)
                                                    var v136 float64 = float64(1.5)
                                                    switch v4 {
                                                    case "foo bar":
                                                        var v137 uint64 = 2
                                                        _ = v137
                                                    case "hello":
                                                        v138 := GenS15{F0: (("hello" + "a") + v4)}
                                                        _ = v138
                                                    case "":
                                                        var v139 int = (v1 - fi130)
                                                        _ = v139
                                                    }
                                                    _ = v136
                                                }
                                            case "a":
                                                v127 += (v127 | v127)
                                            case "":
                                                for fi140 := 0; fi140 < 3; fi140++ {
                                                    var v141 float64 = 0.0
                                                    var v142 int16 = 925
                                                    _ = fi140
                                                    _ = v141
                                                    _ = v142
                                                }
                                            }
                                        case 0:
                                            v1 = (fi130 * (v1 >> ((genFn0()) & 63)))
                                        case 128:
                                            if (((v4 + v4) + v4) >= v4) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v4 = v4
                                                case "\n\t":
                                                    v143 := GenS16{F0: 42, F1: 889}
                                                    _ = v143
                                                case "":
                                                    v144 := GenS17{F0: 823, F1: 42, F2: ((v4 + "") + v4)}
                                                    _ = v144
                                                default:
                                                    var v145 int16 = -32768
                                                    _ = v145
                                                }
                                                var v146 uint8 = 2
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v146
                                            } else {
                                                for fi147 := 0; fi147 < 16; fi147++ {
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v1, v1, fi147)
                                                    }
                                                    switch v4 {
                                                    case "foo bar":
                                                        if false {
                                                            v3 = (v3 + v3)
                                                            switch -2147483648 {
                                                            case 4294967295:
                                                                if false {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (fi147 + (fi147 * 128))
                                                                    }
                                                                    v2 = append(v2, (^(fi130)))
                                                                } else {
                                                                    v4 = v4
                                                                    var v148 float32 = ((0.0 + (0.0 / ((3.14159 + (3.14159 - 1.5)) - 1.5))) + 1e10)
                                                                    fmt.Println(v148)
                                                                    _ = v148
                                                                }
                                                            case 1:
                                                                fmt.Println(fi147, fi130)
                                                            case 32767:
                                                                v3 = v3
                                                            default:
                                                                switch (fi130 + (v1 / ((fi147) | 1))) {
                                                                case 127:
                                                                    var v149 uint16 = (v3 >> ((2) & 63))
                                                                    _ = v149
                                                                case 520:
                                                                    if false {
                                                                        v150 := GenS18{F0: (^(v1)), F1: -1}
                                                                        if ((!false) || (42 == 255)) {
                                                                            v127 += (-(v127))
                                                                            v2 = append(v2, fi147)
                                                                            v151 := GenS19{F0: (v3 << ((255) & 63)), F1: (v4 + ("hello" + ("hello" + (v4 + "foo bar"))))}
                                                                            _ = v151
                                                                        }
                                                                        v152 := GenS8{F0: false, F1: v127, F2: -1}
                                                                        var v153 uint32 = 467
                                                                        var v154 uint8 = 127
                                                                        var v155 uint = 4294967295
                                                                        var v156 float32 = 1.5
                                                                        _ = v150
                                                                        _ = v152
                                                                        _ = v153
                                                                        _ = v154
                                                                        _ = v155
                                                                        _ = v156
                                                                    }
                                                                case 32768:
                                                                    _ = len(v2)
                                                                default:
                                                                    v127 += -324
                                                                }
                                                            }
                                                            if false {
                                                                if ((0 >= 543) && false) {
                                                                    if false {
                                                                        v157 := GenS20{F0: 65535}
                                                                        var v158 int64 = 289
                                                                        _ = v157
                                                                        _ = v158
                                                                    } else {
                                                                        v127 = v127
                                                                        v4 = v4
                                                                        var v159 int64 = -9223372036854775808
                                                                        _ = v159
                                                                    }
                                                                    v2 = append(v2, (v1 % ((fi147) | 1)))
                                                                    v3 += 32767
                                                                    var v160 uint32 = 1
                                                                    _ = v160
                                                                } else {
                                                                    if (!(true || (!false))) {
                                                                        v161 := GenS14{F0: -9223372036854775808, F1: 127}
                                                                        var v162 string = v4
                                                                        if (941 > -1) {
                                                                            var v163 int32 = 0
                                                                            _ = v163
                                                                        }
                                                                        _ = v161
                                                                        _ = v162
                                                                    } else {
                                                                        v4 = v4
                                                                        var v164 uint = genFn0()
                                                                        var v165 uint32 = 0
                                                                        _ = v164
                                                                        _ = v165
                                                                    }
                                                                    var v166 int = fi147
                                                                    _ = v166
                                                                }
                                                                var v167 float64 = (0.1 * (1.5 - (1e10 - 0.0)))
                                                                _ = v167
                                                            } else {
                                                                v168 := GenS9{F0: 0, F1: 29, F2: 42}
                                                                var v169 uint32 = 43
                                                                v4 = ""
                                                                _ = v168
                                                                _ = v169
                                                            }
                                                        }
                                                    case "":
                                                        v4 = (v4 + ((v4 + "\n\t") + v4))
                                                    case "a":
                                                        if false {
                                                            _ = len(v2)
                                                            v170 := GenS7{F0: "foo bar"}
                                                            if false {
                                                                fmt.Println(v3, v3, fi130)
                                                            }
                                                            _ = v170
                                                        } else {
                                                            switch (v1 * v1) {
                                                            case -526:
                                                                v127 = (v127 | (v127 / (((v127 * (v127 / (((v127 & v127)) | 1)))) | 1)))
                                                            case -550:
                                                                v3 = (^(v3))
                                                            default:
                                                                var v171 float32 = 0.1
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi130
                                                                }
                                                                _ = v171
                                                            }
                                                            _ = len(v2)
                                                        }
                                                    default:
                                                        switch v4 {
                                                        case "foo bar":
                                                            switch v4 {
                                                            case "foo bar":
                                                                v127 -= v127
                                                            case "":
                                                                var v172 uint8 = 128
                                                                _ = v172
                                                            }
                                                        }
                                                    }
                                                    var v173 int = (-(v1))
                                                    _ = fi147
                                                    _ = v173
                                                }
                                                for fi174 := 0; fi174 < 11; fi174++ {
                                                    if (true && true) {
                                                        if ((!true) || false) {
                                                            fmt.Println(v127)
                                                        }
                                                        var v175 string = (("hello" + "") + v4)
                                                        v176 := GenS12{F0: 4294967295, F1: 0, F2: v4}
                                                        _ = v175
                                                        _ = v176
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v3 = (v3 << ((1) & 63))
                                                    v127 = 127
                                                    v4 = v4
                                                    _ = fi174
                                                }
                                            }
                                        }
                                        v3 -= v3
                                        _ = fi130
                                    }
                                    var v177 float32 = 3.14159
                                    _ = v177
                                } else {
                                    fmt.Println(v3, v127, v3)
                                }
                            } else {
                                var v178 int = 159
                                _ = v178
                            }
                        }
                        v3 += v3
                        _ = v127
                    }
                    switch v1 {
                    case 0:
                        var v179 uint16 = 32767
                        _ = v179
                    case 32767:
                        var v180 uint = 0
                        _ = v180
                    case 2147483647:
                        v2 = append(v2, 9223372036854775807)
                    }
                }
                v2 = append(v2, v1)
                _ = v42
            } else {
                fmt.Println(v1, v3)
                for fi181 := 0; fi181 < 3; fi181++ {
                    var v182 int8 = 42
                    switch ((v4 + ((v4 + v4) + (v4 + "\n\t"))) + (("foo bar" + ("a" + (v4 + v4))) + "a")) {
                    case "a":
                        var v183 uint32 = 4294967295
                        _ = v183
                    case "":
                        var v184 float64 = 1.5
                        _ = v184
                    case "\n\t":
                        if false {
                            v185 := GenS4{F0: (918 >= 245)}
                            if len(v2) > 0 {
                                v2[0] = fi181
                            }
                            v1 = v1
                            _ = v185
                        }
                    default:
                        switch "" {
                        case "hello":
                            v4 = v4
                        default:
                            v1 += 1
                        }
                    }
                    _ = fi181
                    _ = v182
                }
            }
        }
        if (true || (true && (false || true))) {
            v3 -= v3
            var v186 float64 = 1e10
            for fi187 := 0; fi187 < 20; fi187++ {
                _ = len(v2)
                _ = fi187
            }
            _ = v186
        }
    }
    v1 += 32768
    v188 := GenS21{F0: 1000}
    var v189 uint16 = v3
    var v190 int64 = -1
    for fi191 := 0; fi191 < 4; fi191++ {
        for fi192 := 0; fi192 < 13; fi192++ {
            var v193 int64 = v190
            _ = fi192
            _ = v193
        }
        if true {
            v194 := GenS22{F0: 2, F1: (v4 + "\n\t"), F2: v3}
            _ = v194
        }
        _ = fi191
    }
    fmt.Println(v3, v189)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v195 := GenS7{F0: (v4 + (v4 + v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v189:", v189)
    fmt.Println("v190:", v190)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v188
    _ = v189
    _ = v190
    _ = v195
    _ = v2
    _ = v3
    _ = v4
}
