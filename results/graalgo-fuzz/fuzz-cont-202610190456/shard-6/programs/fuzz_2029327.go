package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint; F1 int; F2 int }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string; F1 string; F2 uint64 }
func genFn0(p0 uint64, p1 bool) uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(p0)
        p1 = (!p1)
        _ = fi5
    }
    switch 32767 {
    case 730:
        v6 := GenS1{F0: 42, F1: p1, F2: -446}
        _ = v6
    case 127:
        var v7 uint = 0
        _ = v7
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 int8, p1 uint8, p2 uint64) uint8 {
    var v8 float64 = float64(((3.14159 / (-2.25 / -2.25)) * ((1e10 / (1.5 * -2.25)) - 3.14159)))
    for fi9 := 0; fi9 < 6; fi9++ {
        v8 -= (-2.25 * float64((((0.0 * 0.1) / 1.5) + 0.0)))
        v8 = v8
        _ = fi9
    }
    var v10 float32 = 1.5
    v10 -= (v10 + ((v10 / (float32(v8) / ((float32(v8) / v10) / v10))) / ((1.5 * v10) - -2.25)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    return p1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 0
    v3 := []uint32{2, 2147483647, 9, 1}
    v4 := make(map[uint32]uint32)
    var v11 int16 = 296
    fmt.Println(v2, v11)
    if true {
        v11 = v11
    } else {
        if false {
            v12 := GenS2{F0: 1}
            if true {
                switch ("hello" + ((("a" + (("hello" + "foo bar") + "")) + "") + "a")) {
                case "a":
                    var v13 uint64 = 1
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 16; fi14++ {
                        fmt.Println(v2, v2)
                        v2 += 1
                        for fi15 := 0; fi15 < 4; fi15++ {
                            v4[0] = 2
                            _ = fi15
                        }
                        _ = fi14
                    }
                }
            }
            _ = v12
        } else {
            var v16 string = ("foo bar" + "a")
            switch (((v16 + ((v16 + "\n\t") + v16)) + "") + v16) {
            case "":
                for fi17 := 0; fi17 < 13; fi17++ {
                    if true {
                        fmt.Println(fi17, fi17)
                        var v18 uint8 = 0
                        _ = v18
                    } else {
                        v16 = "foo bar"
                        if (255 > genFn0(0, true)) {
                            if (!true) {
                                var v19 uint16 = v2
                                _ = v19
                            } else {
                                var v20 uint8 = 42
                                for fi21 := 0; fi21 < 8; fi21++ {
                                    if ((-128 == 1) && false) {
                                        v22 := v4[809]
                                        _ = v22
                                    } else {
                                        v2 = v2
                                        if false {
                                            if (1e10 != (1e10 / (0.0 + 1.5))) {
                                                v23 := GenS2{F0: 128}
                                                v11 -= (v11 | (v11 << ((42) & 63)))
                                                if true {
                                                    switch (v16 + v16) {
                                                    case "a":
                                                        fmt.Println(v16, fi21, fi17)
                                                    default:
                                                        switch v16 {
                                                        case "":
                                                            v24 := GenS3{F0: (^(v20))}
                                                            _ = v24
                                                        case "hello":
                                                            var v25 float64 = (float64(3.14159) - -2.25)
                                                            _ = v25
                                                        case "a":
                                                            var v26 uint32 = 329
                                                            _ = v26
                                                        }
                                                    }
                                                    v20 += genFn1(-128, (v20 & v20), 451)
                                                }
                                                var v27 int64 = 984
                                                var v28 uint64 = 0
                                                _ = v23
                                                _ = v27
                                                _ = v28
                                            }
                                            v16 = ("\n\t" + ((v16 + (v16 + v16)) + ("a" + ((((v16 + "\n\t") + v16) + v16) + "foo bar"))))
                                            switch v16 {
                                            case "a":
                                                switch v16 {
                                                case "foo bar":
                                                    var v29 uint64 = 42
                                                    _ = v29
                                                case "a":
                                                    if false {
                                                        v2 = (v2 ^ (^(v2)))
                                                    } else {
                                                        v20 = (^(v20))
                                                        v30 := GenS4{F0: (((v16 + ((v16 + ((v16 + "a") + "hello")) + v16)) + ((v16 + "hello") + (v16 + ("hello" + ("" + v16))))) + v16), F1: uint32(v20)}
                                                        v31 := GenS1{F0: -51, F1: true, F2: (^(fi17))}
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                }
                                            default:
                                                v16 = v16
                                            }
                                            var v32 int32 = 1
                                            _ = v32
                                        } else {
                                            var v33 float64 = (1.5 - float64((((0.1 / 3.14159) + (1e10 / -2.25)) + (1.5 + (((3.14159 / (1e10 / 0.1)) / 3.14159) - -2.25)))))
                                            _ = v33
                                        }
                                        v34 := GenS4{F0: ("foo bar" + v16), F1: 782}
                                        _ = v34
                                    }
                                    v35 := GenS4{F0: v16, F1: 42}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi21
                                    _ = v35
                                }
                                v11 -= (-(v11))
                                _ = v20
                            }
                        }
                        var v36 uint32 = 2
                        _ = v36
                    }
                    if true {
                        v37 := GenS2{F0: 1}
                        v38 := GenS3{F0: 123}
                        for fi39 := 0; fi39 < 13; fi39++ {
                            var v40 int = 1
                            var v41 int16 = 810
                            fmt.Println(v2, v41, v41)
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        _ = v37
                        _ = v38
                    }
                    if true {
                        v2 = (v2 + (v2 ^ v2))
                    }
                    _ = fi17
                }
            default:
                var v42 uint32 = 2147483647
                _ = v42
            }
            v43 := GenS5{F0: 18446744073709551615, F1: 32768, F2: 128}
            _ = v16
            _ = v43
        }
        for fi44 := 0; fi44 < 6; fi44++ {
            var v45 int64 = 0
            switch fi44 {
            case -326:
                var v46 bool = true
                _ = v46
            case 32767:
                fmt.Println(v11, v11)
            case 867:
                v47 := GenS5{F0: 1, F1: -1, F2: fi44}
                _ = v47
            default:
                if false {
                    v3 = append(v3, 2147483648)
                    v48 := GenS6{F0: -127, F1: "foo bar"}
                    _ = v48
                }
            }
            if false {
                if (float64((1e10 / 1e10)) != float64((((((0.1 - (0.0 * ((float32(3.14159) + 3.14159) / 1.0))) + ((3.14159 - 1.5) * 3.14159)) * ((1e10 * 3.14159) / 0.1)) * 0.1) * 1.5))) {
                    var v49 bool = false
                    var v50 float64 = 1.5
                    fmt.Println(v49, v45)
                    _ = v49
                    _ = v50
                } else {
                    v11 = v11
                    v51 := GenS3{F0: 147}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v51
                }
                for fi52 := 0; fi52 < 8; fi52++ {
                    v53 := GenS7{F0: v2, F1: 1}
                    if (!true) {
                        v54 := GenS8{F0: 42, F1: (!(!(!(0 == 2147483647)))), F2: 42}
                        switch (fi52 | (fi52 % ((-40) | 1))) {
                        case -2147483648:
                            v45 = v45
                        default:
                            v45 = v45
                            v55 := GenS7{F0: v2, F1: 968}
                            v45 += 9223372036854775807
                            switch "\n\t" {
                            case "\n\t":
                                v11 = (v11 - v11)
                                switch "foo bar" {
                                case "a":
                                    v56 := GenS2{F0: 1}
                                    _ = v56
                                }
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v57 string = "hello"
                                    var v58 float32 = 0.1
                                    v59 := GenS4{F0: v57, F1: 281}
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                } else {
                                    v60 := GenS5{F0: 617, F1: (fi44 ^ fi44), F2: fi52}
                                    fmt.Println(fi44)
                                    var v61 float32 = -2.25
                                    v11 = (v11 / (((v11 % (((v11 << ((1) & 63))) | 1))) | 1))
                                    _ = v60
                                    _ = v61
                                }
                            case "hello":
                                v62 := v1[(fi52 ^ fi44)]
                                _ = v62
                            case "foo bar":
                                var v63 uint32 = 2147483648
                                _ = v63
                            default:
                                switch "hello" {
                                case "hello":
                                    switch (fi44 & -249) {
                                    case 0:
                                        v64 := v1[fi52]
                                        var v65 uint32 = 466
                                        if (!false) {
                                            v11 -= (v11 >> ((173) & 63))
                                            if false {
                                                v66 := GenS7{F0: (v2 % ((v2) | 1)), F1: 796}
                                                var v67 int16 = (v11 / ((v11) | 1))
                                                v68 := GenS5{F0: 4294967295, F1: fi44, F2: fi52}
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            } else {
                                                var v69 int32 = -786
                                                _ = v69
                                            }
                                        }
                                        var v70 uint64 = 893
                                        v11 = v11
                                        _ = v64
                                        _ = v65
                                        _ = v70
                                    }
                                case "foo bar":
                                    fmt.Println(v11, v2, fi52)
                                }
                            }
                            switch ("\n\t" + (("foo bar" + "hello") + "\n\t")) {
                            case "":
                                fmt.Println(v2, v11, v2)
                            }
                            _ = v55
                        }
                        v71 := GenS6{F0: 42, F1: ("foo bar" + "hello")}
                        _ = v54
                        _ = v71
                    } else {
                        fmt.Println(fi44)
                        fmt.Println(v45, fi44, fi44)
                        if (!false) {
                            v2 = 1
                        }
                    }
                    _ = fi52
                    _ = v53
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi44
            _ = v45
        }
        v72 := GenS9{F0: "a"}
        _ = v72
    }
    for fi73 := 0; fi73 < 9; fi73++ {
        var v74 int64 = 0
        _ = fi73
        _ = v74
    }
    var v75 int = 9223372036854775807
    for fi76 := 0; fi76 < 20; fi76++ {
        v77 := v1[32767]
        var v78 uint = 0
        _ = fi76
        _ = v77
        _ = v78
    }
    v75 -= (v75 | v75)
    for fi79 := 0; fi79 < 10; fi79++ {
        for fi80 := 0; fi80 < 17; fi80++ {
            fmt.Println(v2)
            var v81 string = "\n\t"
            _ = fi80
            _ = v81
        }
        var v82 int = 0
        _ = fi79
        _ = v82
    }
    fmt.Println(v11)
    if ((v2 & (v2 - (v2 >> ((2) & 63)))) < (v2 + (v2 + v2))) {
        for fi83 := 0; fi83 < 1; fi83++ {
            for fi84 := 0; fi84 < 10; fi84++ {
                switch fi83 {
                case 2147483647:
                    v75 += fi83
                }
                v1[976] = "foo bar"
                if true {
                    v85 := GenS6{F0: -71, F1: "a"}
                    _ = len(v1)
                    fmt.Println(fi83, fi83, fi84)
                    var v86 bool = true
                    _ = v85
                    _ = v86
                }
                _ = fi84
            }
            v75 = v75
            _ = fi83
        }
        v2 = v2
        for fi87 := 0; fi87 < 13; fi87++ {
            switch "" {
            case "a":
                var v88 uint32 = 949
                _ = v88
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                v11 += v11
            }
            _ = fi87
        }
    }
    var v89 string = ("" + "hello")
    v3 = append(v3, 940)
    for fi90 := 0; fi90 < 19; fi90++ {
        switch v89 {
        case "\n\t":
            switch "hello" {
            case "\n\t":
                switch (v89 + "hello") {
                case "a":
                    var v91 float64 = float64(1e10)
                    _ = v91
                case "hello":
                    var v92 float64 = 0.1
                    _ = v92
                }
            default:
                v2 -= 65535
            }
        case "hello":
            if (-32 == -1) {
                switch "" {
                case "hello":
                    var v93 int8 = 1
                    _ = v93
                case "":
                    if (false || (v75 != (^(fi90)))) {
                        var v94 float32 = 1e10
                        _ = v94
                    }
                default:
                    var v95 int32 = 1
                    _ = v95
                }
            }
        case "":
            if true {
                v89 = v89
                var v96 float32 = 0.0
                var v97 uint16 = 2
                _ = v96
                _ = v97
            }
        default:
            v98 := GenS2{F0: 255}
            _ = v98
        }
        var v99 uint64 = 201
        for fi100 := 0; fi100 < 20; fi100++ {
            var v101 int = v75
            v102 := v1[(-(fi100))]
            fmt.Println(v11, v99, fi90)
            _ = fi100
            _ = v101
            _ = v102
        }
        _ = fi90
        _ = v99
    }
    if (!(true || (!(127 >= 0)))) {
        switch "\n\t" {
        case "a":
            v103 := GenS10{F0: 191}
            _ = v103
        case "foo bar":
            for fi104 := 0; fi104 < 16; fi104++ {
                switch (^(fi104)) {
                case -454:
                    if ((255 >= 965) || false) {
                        v2 = v2
                        v105 := GenS3{F0: 128}
                        v106 := v1[v75]
                        _ = v105
                        _ = v106
                    } else {
                        switch 1 {
                        case 115:
                            switch v89 {
                            case "hello":
                                for fi107 := 0; fi107 < 13; fi107++ {
                                    var v108 float32 = -2.25
                                    v109 := GenS7{F0: v2, F1: 2147483647}
                                    _ = fi107
                                    _ = v108
                                    _ = v109
                                }
                            case "foo bar":
                                v89 = ((v89 + v89) + ("" + "foo bar"))
                            }
                        }
                        var v110 bool = (!((42 != 18446744073709551615) && (false || false)))
                        _ = v110
                    }
                case 32768:
                    var v111 int16 = v11
                    _ = v111
                case -420:
                    v75 += (fi104 % ((127) | 1))
                }
                _ = fi104
            }
        case "\n\t":
            if (-1 <= 42) {
                if (!(!(18446744073709551615 <= 1))) {
                    var v112 string = v89
                    _ = v112
                } else {
                    v3 = append(v3, 65535)
                    switch ("a" + v89) {
                    case "a":
                        v113 := GenS10{F0: int64(v75)}
                        _ = v113
                    default:
                        for fi114 := 0; fi114 < 2; fi114++ {
                            v115 := GenS4{F0: v89, F1: 4294967295}
                            _ = fi114
                            _ = v115
                        }
                    }
                    v11 = (v11 - (v11 / ((v11) | 1)))
                }
                if (!(!(!true))) {
                    var v116 uint = 42
                    if false {
                        for fi117 := 0; fi117 < 17; fi117++ {
                            if (-809 < -980) {
                                var v118 float32 = 0.1
                                _ = v118
                            } else {
                                var v119 uint16 = (v2 / ((v2) | 1))
                                _ = v119
                            }
                            var v120 int16 = 1
                            var v121 int32 = 0
                            _ = fi117
                            _ = v120
                            _ = v121
                        }
                        v122 := GenS11{F0: v75}
                        _ = v122
                    }
                    _ = v116
                }
            } else {
                fmt.Println(v75)
                var v123 int16 = 32767
                _ = len(v4)
                _ = v123
            }
        default:
            switch (("hello" + v89) + ((((v89 + v89) + "\n\t") + v89) + (v89 + "foo bar"))) {
            case "":
                var v124 uint32 = 2147483647
                _ = v124
            default:
                v125 := GenS8{F0: 222, F1: (!(((true && false) || (255 == 255)) || (v2 <= v2))), F2: 18446744073709551615}
                _ = v125
            }
        }
        for fi126 := 0; fi126 < 2; fi126++ {
            for fi127 := 0; fi127 < 7; fi127++ {
                var v128 uint32 = 57
                switch (v75 >> ((2) & 63)) {
                case 32768:
                    var v129 int8 = 100
                    _ = v129
                case 129:
                    var v130 int64 = 0
                    _ = v130
                }
                v131 := GenS9{F0: v89}
                switch (fi126 + (fi126 / ((88) | 1))) {
                case 352:
                    var v132 float32 = (0.0 / 1.5)
                    _ = v132
                case -9223372036854775808:
                    if false {
                        v75 = (v75 / ((42) | 1))
                        switch 127 {
                        case -240:
                            var v133 string = v89
                            _ = v133
                        default:
                            var v134 int8 = -38
                            _ = v134
                        }
                    } else {
                        fmt.Println(fi127)
                    }
                case -1:
                    if (!((!true) && (true && (v89 >= "hello")))) {
                        v75 += fi127
                        v128 = 1
                        if (uint(v128) == genFn0(18446744073709551615, false)) {
                            v3 = append(v3, 4294967295)
                        }
                    } else {
                        delete(v1, 2147483647)
                        var v135 int32 = 2147483647
                        _ = v135
                    }
                }
                _ = fi127
                _ = v128
                _ = v131
            }
            fmt.Println(v89, v2)
            _ = fi126
        }
        if true {
            v75 = v75
            if (true || false) {
                fmt.Println(v75)
                v2 += 2
            } else {
                fmt.Println(v11, v89, v11)
                switch (v75 << ((255) & 63)) {
                case 127:
                    v11 -= v11
                case 1:
                    var v136 int16 = 20
                    _ = v136
                case 9223372036854775807:
                    var v137 int32 = 951
                    _ = v137
                }
                v75 = (v75 / ((v75) | 1))
            }
            var v138 float64 = -2.25
            _ = v138
        } else {
            v2 = (v2 % ((v2) | 1))
            v4[2147483648] = 421
        }
    }
    switch ((v89 + v89) + "\n\t") {
    case "\n\t":
        var v139 int32 = -1
        _ = v139
    case "foo bar":
        switch (v89 + v89) {
        case "foo bar":
            switch "" {
            case "foo bar":
                switch ("a" + v89) {
                case "hello":
                    var v140 uint64 = 151
                    _ = v140
                }
            case "hello":
                v141 := GenS12{F0: 42}
                _ = v141
            case "a":
                v89 = "a"
            }
        default:
            for fi142 := 0; fi142 < 16; fi142++ {
                if true {
                    var v143 uint64 = 284
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v89 = (v89 + "a")
                    _ = v143
                } else {
                    v144 := GenS13{F0: (("a" + v89) + v89), F1: ((v89 + (v89 + v89)) + v89), F2: 1}
                    v2 += (v2 + v2)
                    switch v89 {
                    case "foo bar":
                        var v145 int64 = 61
                        _ = v145
                    case "\n\t":
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                    }
                    _ = v144
                }
                for fi146 := 0; fi146 < 17; fi146++ {
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    _ = fi146
                }
                var v147 float32 = 0.1
                _ = fi142
                _ = v147
            }
        }
    case "":
        v4[223] = 507
    default:
        var v148 int = (^(v75))
        _ = v148
    }
    if (!false) {
        v1[(v75 >> ((18446744073709551615) & 63))] = v89
    }
    if (!true) {
        if true {
            v4[126] = 4294967295
        } else {
            if (42 >= 32767) {
                fmt.Println(v75, v89, v2)
                var v149 int = v75
                v150 := GenS10{F0: 1}
                _ = v149
                _ = v150
            } else {
                v89 = ((("hello" + (("\n\t" + v89) + v89)) + ((v89 + ("a" + "")) + v89)) + "")
                var v151 uint = 86
                if len(v3) > 0 {
                    v3[0] = 42
                }
                _ = v151
            }
        }
        v152 := GenS8{F0: 128, F1: true, F2: 13}
        fmt.Println(v75, v11)
        _ = v152
    }
    v11 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v75:", v75)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v75
    _ = v89
}
