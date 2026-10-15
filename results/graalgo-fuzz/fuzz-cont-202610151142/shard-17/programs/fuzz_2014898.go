package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 uint16; F1 string; F2 int }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 int8; F2 int }
type GenS10 struct { F0 int64; F1 int16; F2 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 int16; F2 uint64 }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 uint64; F1 int16 }
type GenS19 struct { F0 int64; F1 int8; F2 uint32 }
func genFn0(p0 bool, p1 int64, p2 int64) int32 {
    var v5 uint16 = 37
    p1 += -9223372036854775808
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: "a", F1: v5}
        switch (fi6 >> ((18446744073709551615) & 63)) {
        case 42:
            v8 := GenS2{F0: 192}
            _ = v8
        case 4294967295:
            var v9 uint8 = 2
            _ = v9
        case 1:
            v10 := []int{(fi6 * fi6), (fi6 & fi6), fi6}
            _ = v10
        default:
            v11 := []int{-2147483648, fi6, (fi6 + (fi6 / ((fi6) | 1)))}
            _ = v11
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := []int{fi6, 4294967295, fi6}
            _ = len(v13)
            _ = fi12
            _ = v13
        }
        _ = fi6
        _ = v7
    }
    if p0 {
        p1 = 1
        if p0 {
            v14 := GenS2{F0: 255}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS2{F0: 2}
                var v17 float32 = 1e10
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    } else {
        var v18 bool = (false || (-677 >= -2147483648))
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return int32(p1)
}
func genFn1(p0 uint16, p1 int16) uint8 {
    v19 := GenS3{F0: 255, F1: 1000, F2: -1}
    p0 = (^(p0))
    var v20 uint16 = p0
    fmt.Println(p1)
    var v21 string = "foo bar"
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return 255
}

func main() {
    var v1 int = 127
    var v2 uint = 1
    var v3 uint16 = 65535
    var v4 uint32 = 1
    v22 := GenS1{F0: "foo bar", F1: (v3 | (v3 >> ((v2) & 63)))}
    if true {
        fmt.Println(v1, v4)
        v23 := []int{v1, 127, (v1 >> ((v2) & 63))}
        switch ("" + "hello") {
        case "foo bar":
            var v24 uint32 = uint32(v3)
            _ = v24
        }
        _ = v23
    } else {
        v2 = 1
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        fmt.Println(v1, v2, v3)
        var v26 string = "a"
        v4 = v4
        _ = fi25
        _ = v26
    }
    switch "" {
    case "foo bar":
        switch "" {
        case "hello":
            switch "a" {
            case "\n\t":
                v2 = v2
            case "foo bar":
                v27 := []int{v1, v1, -9223372036854775808, (v1 & 0)}
                _ = v27
            case "hello":
                for fi28 := 0; fi28 < 13; fi28++ {
                    switch "a" {
                    case "hello":
                        v29 := []int{2147483647, fi28}
                        _ = v29
                    default:
                        if (true || false) {
                            switch "hello" {
                            case "":
                                v30 := GenS2{F0: 210}
                                _ = v30
                            }
                        } else {
                            if true {
                                v4 += (v4 - (v4 - v4))
                                v2 -= 771
                                for fi31 := 0; fi31 < 1; fi31++ {
                                    v2 = (^(v2))
                                    _ = fi31
                                }
                            }
                            v32 := []int{v1, 127, 4294967295, (-(fi28))}
                            _ = v32
                        }
                    }
                    v3 = v3
                    if true {
                        var v33 float32 = 1e10
                        for fi34 := 0; fi34 < 14; fi34++ {
                            var v35 uint16 = (^(v3))
                            var v36 uint32 = v4
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        _ = v33
                    } else {
                        for fi37 := 0; fi37 < 10; fi37++ {
                            var v38 uint = (^(v2))
                            v1 = int(v4)
                            v39 := []int{int(v3), fi28, fi28, 32768}
                            var v40 float32 = (3.14159 / (((float32(0.0) * 3.14159) - 1.5) + 3.14159))
                            _ = fi37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 17; fi41++ {
                            if (!false) {
                                switch ("" + "hello") {
                                case "\n\t":
                                    fmt.Println(v2, v2)
                                case "a":
                                    v42 := []int{(fi41 % ((fi41) | 1)), fi41, 32767}
                                    _ = v42
                                }
                                v43 := GenS2{F0: genFn1((v3 >> ((v2) & 63)), 32767)}
                                var v44 uint16 = 1
                                _ = v43
                                _ = v44
                            } else {
                                fmt.Println(v1, fi41, v3)
                            }
                            _ = fi41
                        }
                    }
                    _ = fi28
                }
            }
        default:
            if true {
                for fi45 := 0; fi45 < 19; fi45++ {
                    v46 := []int{}
                    for fi47 := 0; fi47 < 10; fi47++ {
                        switch fi47 {
                        case 1:
                            v46 = append(v46, v1)
                        case 2147483647:
                            if len(v46) >= 2 {
                                v46 = v46[1:len(v46)]
                            }
                        case -566:
                            v2 = (v2 | v2)
                        default:
                            v46 = append(v46, 1)
                        }
                        _ = fi47
                    }
                    _ = fi45
                    _ = v46
                }
                v48 := []int{}
                var v49 string = ""
                _ = v48
                _ = v49
            } else {
                v50 := []int{(v1 % (((v1 * (v1 / ((515) | 1)))) | 1)), (v1 - (v1 + v1))}
                fmt.Println(v3)
                switch ("a" + "") {
                case "\n\t":
                    v1 = 9223372036854775807
                case "hello":
                    for fi51 := 0; fi51 < 15; fi51++ {
                        if len(v50) >= 2 {
                            v50 = v50[1:len(v50)]
                        }
                        if false {
                            v52 := GenS2{F0: 253}
                            _ = v52
                        } else {
                            v53 := GenS4{F0: (v3 % ((v3) | 1)), F1: ("foo bar" + "a"), F2: fi51}
                            _ = v53
                        }
                        _ = fi51
                    }
                case "a":
                    v54 := GenS3{F0: 39, F1: 32767, F2: 127}
                    _ = v54
                default:
                    if false {
                        v3 -= (v3 + v3)
                        if false {
                            if false {
                                if len(v50) > 0 {
                                    v50[0] = 0
                                }
                                for fi55 := 0; fi55 < 7; fi55++ {
                                    switch "hello" {
                                    case "":
                                        var v56 bool = false
                                        _ = v56
                                    case "\n\t":
                                        v50 = append(v50, 589)
                                    case "a":
                                        v57 := GenS3{F0: 127, F1: -32768, F2: fi55}
                                        _ = v57
                                    default:
                                        var v58 uint16 = (v3 - v3)
                                        _ = v58
                                    }
                                    var v59 uint64 = 935
                                    v59 += v59
                                    _ = fi55
                                    _ = v59
                                }
                                v60 := GenS5{F0: v2, F1: (v2 | (v2 ^ (v2 * (v2 + (v2 ^ (^(v2)))))))}
                                _ = v60
                            }
                            fmt.Println(v2)
                            for fi61 := 0; fi61 < 1; fi61++ {
                                fmt.Println(v3, fi61, fi61)
                                v50 = append(v50, 42)
                                _ = fi61
                            }
                        } else {
                            fmt.Println(v1, v3)
                            fmt.Println(v2)
                            var v62 uint16 = (^(v3))
                            _ = v62
                        }
                    }
                }
                _ = v50
            }
        }
    }
    var v63 uint16 = (v3 / (((v3 ^ 89)) | 1))
    for fi64 := 0; fi64 < 13; fi64++ {
        v4 += v4
        var v65 int16 = 0
        _ = fi64
        _ = v65
    }
    fmt.Println(v2, v3, v63)
    var v66 float32 = float32(3.14159)
    v67 := GenS5{F0: 2, F1: 2}
    if ((v2 <= 365) || ("hello" > "foo bar")) {
        for fi68 := 0; fi68 < 20; fi68++ {
            v4 -= v4
            if true {
                switch -387 {
                case 4294967295:
                    switch int(v4) {
                    case 1:
                        if false {
                            for fi69 := 0; fi69 < 20; fi69++ {
                                v70 := GenS2{F0: 48}
                                v71 := GenS6{F0: 42}
                                _ = fi69
                                _ = v70
                                _ = v71
                            }
                            switch -9223372036854775808 {
                            case -2147483648:
                                for fi72 := 0; fi72 < 17; fi72++ {
                                    var v73 int32 = genFn0((false && ((203 > uint32(v2)) && (true && true))), -313, 1)
                                    v63 -= (v3 * v3)
                                    switch "\n\t" {
                                    case "a":
                                        v73 = (-(v73))
                                    case "\n\t":
                                        var v74 int16 = 32767
                                        v75 := []int{}
                                        _ = v74
                                        _ = v75
                                    case "foo bar":
                                        if true {
                                            var v76 int64 = 1
                                            _ = v76
                                        } else {
                                            if (!(("hello" <= "a") || (genFn1(uint16(v4), 1) <= genFn1((^(v63)), 82)))) {
                                                v77 := GenS7{F0: 916, F1: (v1 << ((v2) & 63))}
                                                var v78 float64 = (0.0 * (1e10 + 1e10))
                                                _ = v77
                                                _ = v78
                                            }
                                            v1 = fi72
                                            switch fi68 {
                                            case -9223372036854775808:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v79 := GenS8{F0: "\n\t"}
                                                    _ = v79
                                                case "a":
                                                    if ((false || (v63 < uint16(v2))) || (!(v66 <= v66))) {
                                                        v80 := GenS1{F0: (("" + "a") + "a"), F1: (v3 ^ 638)}
                                                        switch 9223372036854775807 {
                                                        case -586:
                                                            fmt.Println(v2)
                                                        case 128:
                                                            v81 := []int{int(v73), (-(fi68)), v1}
                                                            _ = v81
                                                        }
                                                        v82 := []int{42, -1, (v1 / ((fi68) | 1)), (-(fi68))}
                                                        var v83 int = (fi68 / ((-747) | 1))
                                                        var v84 uint = v2
                                                        _ = v80
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        v85 := GenS6{F0: 1}
                                                        if true {
                                                            var v86 int16 = -1
                                                            var v87 float64 = (3.14159 - 0.0)
                                                            v66 = 1.5
                                                            if false {
                                                                if false {
                                                                    v88 := []int{(fi68 & fi68), (fi72 * (^(v1))), -2147483648, fi68}
                                                                    if false {
                                                                        var v89 string = ""
                                                                        if true {
                                                                            v90 := GenS5{F0: (v2 & (^(v2))), F1: (v2 % (((^(v2))) | 1))}
                                                                            if false {
                                                                                var v91 uint32 = (v4 | (v4 - 4294967295))
                                                                                var v92 float32 = v66
                                                                                _ = len(v88)
                                                                                _ = len(v88)
                                                                                _ = v91
                                                                                _ = v92
                                                                            } else {
                                                                                switch fi68 {
                                                                                case 128:
                                                                                    v66 = (float32(((v87 - v87) + 1.5)) - v66)
                                                                                case 0:
                                                                                    v2 = (v2 ^ v2)
                                                                                default:
                                                                                    fmt.Println(v66)
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v1, v3)
                                                                                } else {
                                                                                    if (2 <= (v3 >> ((uint(v4)) & 63))) {
                                                                                        v93 := GenS9{F0: v73, F1: -128, F2: fi72}
                                                                                        switch v1 {
                                                                                        case 0:
                                                                                            var v94 float64 = (1e10 - v87)
                                                                                            var v95 int16 = (v86 & v86)
                                                                                            if (-128 < -47) {
                                                                                                fmt.Println(v73)
                                                                                                var v96 uint16 = v63
                                                                                                v63 += (v63 << (((v2 / ((v2) | 1))) & 63))
                                                                                                v97 := GenS10{F0: 9223372036854775807, F1: v86, F2: 810}
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                            }
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        case 42:
                                                                                            fmt.Println(v4)
                                                                                        case -660:
                                                                                            var v98 int16 = int16(fi68)
                                                                                            _ = v98
                                                                                        }
                                                                                        v1 += v1
                                                                                        _ = v93
                                                                                    } else {
                                                                                        if false {
                                                                                            var v99 uint32 = (v4 % (((v4 ^ v4)) | 1))
                                                                                            if false {
                                                                                                switch 127 {
                                                                                                case 32768:
                                                                                                    v86 += v86
                                                                                                case 127:
                                                                                                    fmt.Println(fi72)
                                                                                                case -44:
                                                                                                    if (1 >= 127) {
                                                                                                        v86 -= (v86 | v86)
                                                                                                        fmt.Println(v99)
                                                                                                        if (2 < v2) {
                                                                                                            v63 -= (v63 ^ 2)
                                                                                                        } else {
                                                                                                            v88 = append(v88, (fi68 >> ((231) & 63)))
                                                                                                            if ((v86 / ((v86) | 1)) != (v86 >> ((399) & 63))) {
                                                                                                                v100 := GenS11{F0: (v1 & fi72)}
                                                                                                                v101 := GenS9{F0: v73, F1: 100, F2: v1}
                                                                                                                _ = v100
                                                                                                                _ = v101
                                                                                                            }
                                                                                                            v99 = (v4 % ((uint32(v3)) | 1))
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                switch int(v2) {
                                                                                                case 103:
                                                                                                    if true {
                                                                                                        _ = len(v88)
                                                                                                        var v102 uint = v2
                                                                                                        v89 = ("a" + v89)
                                                                                                        v103 := GenS1{F0: (v89 + v89), F1: v3}
                                                                                                        v87 = 1.5
                                                                                                        _ = v102
                                                                                                        _ = v103
                                                                                                    }
                                                                                                }
                                                                                                switch ((v89 + v89) + (v89 + v89)) {
                                                                                                case "a":
                                                                                                    if true {
                                                                                                        if len(v88) >= 2 {
                                                                                                            v88 = v88[1:len(v88)]
                                                                                                        }
                                                                                                        var v104 int16 = v86
                                                                                                        if true {
                                                                                                            var v105 string = ("\n\t" + (v89 + "\n\t"))
                                                                                                            _ = v105
                                                                                                        } else {
                                                                                                            var v106 float32 = v66
                                                                                                            var v107 int64 = 516
                                                                                                            var v108 uint64 = 0
                                                                                                            _ = v106
                                                                                                            _ = v107
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        fmt.Println(v66, v2, v99)
                                                                                                        _ = v104
                                                                                                    } else {
                                                                                                        var v109 float64 = v87
                                                                                                        fmt.Println(v4)
                                                                                                        _ = v109
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v66, v3)
                                                                                                }
                                                                                            } else {
                                                                                                if len(v88) >= 2 {
                                                                                                    v88 = v88[1:len(v88)]
                                                                                                }
                                                                                                v99 = v99
                                                                                                if len(v88) > 0 {
                                                                                                    v88[0] = v1
                                                                                                }
                                                                                            }
                                                                                            _ = v99
                                                                                        } else {
                                                                                            v110 := GenS12{F0: (-(fi72)), F1: (v86 * (v86 * -32768)), F2: uint64(v2)}
                                                                                            switch v89 {
                                                                                            case "hello":
                                                                                                var v111 float32 = (v66 / ((float32(float64((0.0 + -2.25))) - v66) / 1.5))
                                                                                                _ = v111
                                                                                            default:
                                                                                                _ = len(v88)
                                                                                            }
                                                                                            fmt.Println(v2)
                                                                                            _ = v110
                                                                                        }
                                                                                    }
                                                                                    if len(v88) > 0 {
                                                                                        v88[0] = (fi68 * int(v4))
                                                                                    }
                                                                                    v3 -= 32767
                                                                                }
                                                                                switch 128 {
                                                                                case 223:
                                                                                    v3 -= v3
                                                                                default:
                                                                                    var v112 uint8 = 0
                                                                                    _ = v112
                                                                                }
                                                                            }
                                                                            v88 = append(v88, 0)
                                                                            _ = v90
                                                                        }
                                                                        var v113 uint64 = 42
                                                                        _ = v113
                                                                        _ = v89
                                                                    }
                                                                    var v114 uint = v2
                                                                    var v115 string = "\n\t"
                                                                    v3 += (v3 + v3)
                                                                    _ = v114
                                                                    _ = v115
                                                                    _ = v88
                                                                } else {
                                                                    v116 := GenS13{F0: 194, F1: false}
                                                                    v117 := GenS14{F0: v3}
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                                switch "a" {
                                                                case "\n\t":
                                                                    if true {
                                                                        var v118 int8 = -5
                                                                        var v119 float64 = 1e10
                                                                        _ = v118
                                                                        _ = v119
                                                                    }
                                                                    if (float64(v66) > 3.14159) {
                                                                        v120 := GenS15{F0: -86}
                                                                        _ = v120
                                                                    }
                                                                    var v121 float64 = v87
                                                                    var v122 uint16 = v3
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v3 += v3
                                                                    case "":
                                                                        v121 = (((-2.25 - v87) + (v121 + 0.0)) - 0.0)
                                                                    }
                                                                    _ = v121
                                                                    _ = v122
                                                                case "a":
                                                                    var v123 uint16 = (^(v63))
                                                                    _ = v123
                                                                case "":
                                                                    if ((v73 ^ (v73 - (v73 - v73))) == 1) {
                                                                        var v124 int8 = 0
                                                                        _ = v124
                                                                    }
                                                                }
                                                            }
                                                            _ = v86
                                                            _ = v87
                                                        } else {
                                                            fmt.Println(v66)
                                                        }
                                                        _ = v85
                                                    }
                                                case "":
                                                    v125 := []int{(fi68 + fi68), 751, fi68}
                                                    _ = v125
                                                }
                                            }
                                        }
                                    }
                                    var v126 float32 = (v66 * 0.0)
                                    _ = fi72
                                    _ = v126
                                    _ = v73
                                }
                            case 32768:
                                v127 := []int{128}
                                _ = v127
                            }
                        }
                    case 42:
                        for fi128 := 0; fi128 < 7; fi128++ {
                            var v129 uint = v2
                            if true {
                                switch (("" + "a") + "a") {
                                case "\n\t":
                                    v130 := GenS16{F0: 0}
                                    _ = v130
                                case "":
                                    v131 := GenS1{F0: ("\n\t" + "hello"), F1: v3}
                                    _ = v131
                                case "foo bar":
                                    v132 := GenS17{F0: true}
                                    _ = v132
                                default:
                                    v133 := []int{}
                                    _ = v133
                                }
                                v66 = (-2.25 * v66)
                                v66 = v66
                            }
                            _ = fi128
                            _ = v129
                        }
                    }
                case 911:
                    if (!(!true)) {
                        for fi134 := 0; fi134 < 15; fi134++ {
                            var v135 float32 = v66
                            _ = fi134
                            _ = v135
                        }
                    }
                default:
                    var v136 int32 = 42
                    _ = v136
                }
            } else {
                v137 := GenS17{F0: (genFn1(0, 658) <= 2)}
                for fi138 := 0; fi138 < 11; fi138++ {
                    v139 := []int{(v1 | fi68)}
                    fmt.Println(v63, v66)
                    var v140 uint16 = v63
                    _ = fi138
                    _ = v139
                    _ = v140
                }
                _ = v137
            }
            var v141 float64 = 1.5
            _ = fi68
            _ = v141
        }
        var v142 string = ("hello" + "\n\t")
        _ = v142
    }
    var v143 uint = 65535
    switch "\n\t" {
    case "foo bar":
        if (true || false) {
            fmt.Println(v2, v66)
        }
    case "":
        v144 := []int{248, 1, 933}
        _ = v144
    case "hello":
        for fi145 := 0; fi145 < 13; fi145++ {
            var v146 uint32 = 411
            switch (("\n\t" + "foo bar") + (("\n\t" + "\n\t") + "")) {
            case "hello":
                switch "hello" {
                case "hello":
                    var v147 uint16 = v3
                    _ = v147
                case "a":
                    var v148 int8 = 100
                    _ = v148
                default:
                    if ("foo bar" <= ("foo bar" + "hello")) {
                        v3 = 405
                    }
                }
            }
            var v149 string = "foo bar"
            _ = fi145
            _ = v146
            _ = v149
        }
    }
    v150 := []int{v1, (v1 % (((v1 * v1)) | 1)), v1}
    var v151 int = (v1 / ((v1) | 1))
    v152 := GenS18{F0: 1, F1: 32767}
    v153 := GenS19{F0: 42, F1: 1, F2: v4}
    if true {
        if (((!(!((v3 + v63) == 42))) && true) && false) {
            if len(v150) > 0 {
                v150[0] = (v1 - v1)
            }
        } else {
            switch ("a" + ("foo bar" + "")) {
            case "":
                fmt.Println(v2)
            case "hello":
                if len(v150) > 0 {
                    v150[0] = (v151 >> ((2) & 63))
                }
            case "foo bar":
                switch ("\n\t" + "hello") {
                case "":
                    var v154 bool = false
                    _ = v154
                case "a":
                    fmt.Println(v143)
                case "\n\t":
                    if (42 <= 1) {
                        var v155 int64 = 1
                        var v156 int = v151
                        _ = v155
                        _ = v156
                    }
                }
            }
            if len(v150) >= 2 {
                v150 = v150[1:len(v150)]
            }
        }
        v3 += 32767
        var v157 string = "foo bar"
        _ = v157
    }
    for fi158 := 0; fi158 < 16; fi158++ {
        v1 = 2147483647
        v3 -= 42
        v143 = v143
        _ = fi158
    }
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("slicelen v150:", len(v150), v150)
    fmt.Println("v151:", v151)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v143
    _ = v150
    _ = v151
    _ = v152
    _ = v153
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v63
    _ = v66
    _ = v67
}
