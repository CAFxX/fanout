package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 string; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint16; F1 int; F2 int8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 uint64; F2 int16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 bool; F1 bool; F2 int16 }
type GenS12 struct { F0 int64; F1 uint }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 string; F1 string }
type GenS15 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS16 struct { F0 uint32; F1 uint32 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint; F1 int64; F2 uint32 }
type GenS20 struct { F0 int8; F1 int32; F2 uint32 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 int; F1 int8 }
func genFn0(p0 int64, p1 int32, p2 int8) int8 {
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            p0 += 1
            v7 := GenS1{F0: -99, F1: 1, F2: -319}
            p0 = p0
            _ = v7
        } else {
            v8 := []int{}
            v9 := GenS1{F0: (p2 * 100), F1: 0, F2: int16(p2)}
            _ = v8
            _ = v9
        }
        p1 += p1
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 bool = true
            var v12 uint16 = 255
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi6
    }
    var v13 int32 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (p2 & -67)
}

func main() {
    var v1 string = ""
    var v2 int = 2147483647
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    var v5 float32 = 1e10
    v5 -= (v5 / ((1e10 / ((float32(1e10) * v5) / v5)) + (v5 - float32(float64((v5 - v5))))))
    for fi14 := 0; fi14 < 18; fi14++ {
        if (-128 >= genFn0(0, -1, -1)) {
            if (false || (v2 >= (fi14 - (^(fi14))))) {
                var v15 uint = 65
                v16 := GenS1{F0: 100, F1: 1000, F2: 498}
                _ = v15
                _ = v16
            }
        }
        v17 := GenS2{F0: 1}
        v4[v1] = false
        _ = fi14
        _ = v17
    }
    if ((784 < 1) || ("" != "hello")) {
        if (!false) {
            for fi18 := 0; fi18 < 5; fi18++ {
                _ = len(v3)
                v19 := []int{127, (v2 - 45), (v2 / ((fi18) | 1))}
                v20 := GenS3{F0: true, F1: (v1 + "hello"), F2: 32767}
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 int64 = 721
            for fi22 := 0; fi22 < 14; fi22++ {
                _ = len(v4)
                if false {
                    var v23 int32 = 972
                    var v24 uint64 = 18446744073709551615
                    _ = v23
                    _ = v24
                } else {
                    var v25 float64 = 3.14159
                    _ = v25
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 uint = 1
                    switch v1 {
                    case "a":
                        v28 := []int{}
                        _ = v28
                    case "\n\t":
                        v29 := GenS3{F0: true, F1: v1, F2: 32767}
                        _ = v29
                    case "hello":
                        var v30 bool = true
                        _ = v30
                    }
                    fmt.Println(v5)
                    if false {
                        v31 := GenS2{F0: 525}
                        if (((genFn0((v21 ^ v21), -4, 0) <= 42) && ((^(v27)) <= (v27 - (v27 * (v27 / ((2) | 1)))))) && (uint16(v27) < 32767)) {
                            if ((42 >= -83) || (v1 != "hello")) {
                                if true {
                                    switch (v1 + v1) {
                                    case "\n\t":
                                        v2 -= v2
                                    case "hello":
                                        fmt.Println(fi22, v2)
                                    case "a":
                                        if true {
                                            var v32 float64 = 1.5
                                            _ = v32
                                        }
                                    }
                                } else {
                                    var v33 float64 = (3.14159 / 1.5)
                                    fmt.Println(v33, v21, v5)
                                    _ = v33
                                }
                            }
                            v3["\n\t"] = fi22
                            v34 := GenS4{F0: 557, F1: (v21 - v21), F2: 65535}
                            _ = v34
                        } else {
                            fmt.Println(v2, fi26, v2)
                            delete(v4, v1)
                            switch v2 {
                            case -596:
                                fmt.Println(v2)
                                if ((1 < 14) || (-799 >= 0)) {
                                    var v35 int = fi22
                                    var v36 int16 = int16(v27)
                                    if ((!((v21 + v21) <= v21)) || (-2.25 != -2.25)) {
                                        if (v36 > (v36 * 1000)) {
                                            var v37 uint8 = 2
                                            _ = v37
                                        } else {
                                            v38 := GenS2{F0: 1}
                                            v39 := GenS1{F0: int8(v2), F1: (v36 >> ((1) & 63)), F2: int16(v35)}
                                            v5 = float32((0.1 - float64(1e10)))
                                            _ = v38
                                            _ = v39
                                        }
                                        v1 = v1
                                        var v40 uint = (v27 | 18446744073709551615)
                                        _ = v40
                                    }
                                    _ = v35
                                    _ = v36
                                }
                            default:
                                fmt.Println(fi22, v5)
                            }
                        }
                        _ = v31
                    }
                    if true {
                        v41 := GenS5{F0: 498, F1: 32767, F2: -1}
                        switch (^(v2)) {
                        case -1:
                            switch (v1 + ("foo bar" + ("foo bar" + "hello"))) {
                            case "":
                                var v42 int = fi22
                                _ = v42
                            }
                        }
                        delete(v4, (v1 + "\n\t"))
                        v43 := GenS6{F0: 0, F1: -128}
                        v44 := GenS7{F0: 32767}
                        v45 := GenS7{F0: 255}
                        fmt.Println(v27, v5)
                        _ = v41
                        _ = v43
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint32 = 4294967295
                        _ = v46
                    }
                    v21 = v21
                    _ = fi26
                    _ = v27
                }
                _ = fi22
            }
            _ = v21
        }
        if (42 < 255) {
            for fi47 := 0; fi47 < 5; fi47++ {
                if (0 <= 128) {
                    var v48 uint = 42
                    if true {
                        var v49 int16 = 1000
                        var v50 int32 = -1
                        fmt.Println(v50)
                        _ = v49
                        _ = v50
                    }
                    _ = v48
                } else {
                    v51 := GenS1{F0: 42, F1: 305, F2: 32767}
                    switch fi47 {
                    case 4294967295:
                        switch (v1 + (("" + v1) + "foo bar")) {
                        case "foo bar":
                            for fi52 := 0; fi52 < 14; fi52++ {
                                switch (fi52 | (fi47 & (-(fi47)))) {
                                case 32768:
                                    v53 := GenS8{F0: 0}
                                    _ = v53
                                default:
                                    v2 = fi52
                                }
                                if false {
                                    fmt.Println(fi52, v5)
                                    var v54 uint16 = 42
                                    if true {
                                        v55 := v4[v1]
                                        v2 = fi52
                                        var v56 int = (fi47 % (((fi52 * (fi47 >> ((725) & 63)))) | 1))
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v1 = v1
                                    }
                                    _ = v54
                                } else {
                                    v57 := GenS9{F0: 42, F1: 18446744073709551615, F2: -148}
                                    _ = v57
                                }
                                if (1.5 <= v5) {
                                    v58 := GenS5{F0: 416, F1: 893, F2: 42}
                                    if true {
                                        v2 = (fi52 << ((926) & 63))
                                        if true {
                                            v2 = -461
                                            v2 -= fi47
                                        }
                                    } else {
                                        var v59 bool = true
                                        v60 := GenS10{F0: 777}
                                        v61 := GenS2{F0: 0}
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v58
                                }
                                fmt.Println(v1, v2, v1)
                                var v62 int64 = 42
                                _ = fi52
                                _ = v62
                            }
                        case "":
                            v63 := []int{fi47}
                            _ = v63
                        case "hello":
                            for fi64 := 0; fi64 < 5; fi64++ {
                                v5 -= v5
                                if false {
                                    if true {
                                        v2 = fi64
                                        delete(v4, "\n\t")
                                        v65 := v3[(((v1 + ("foo bar" + v1)) + v1) + "\n\t")]
                                        _ = v65
                                    } else {
                                        if true {
                                            v5 -= (0.1 * 3.14159)
                                            if false {
                                                var v66 uint = 4294967295
                                                v1 = "\n\t"
                                                fmt.Println(v5)
                                                _ = v66
                                            } else {
                                                if true {
                                                    v1 = (v1 + "")
                                                } else {
                                                    v2 -= (fi47 * (^(v2)))
                                                }
                                                var v67 float32 = v5
                                                _ = v67
                                            }
                                            var v68 float32 = ((v5 * v5) - ((3.14159 - 1e10) + 1.5))
                                            _ = v68
                                        }
                                        fmt.Println(v2, fi47, v5)
                                        var v69 int64 = -260
                                        switch 127 {
                                        case 32768:
                                            var v70 string = "a"
                                            _ = v70
                                        }
                                        v69 += v69
                                        _ = v69
                                    }
                                } else {
                                    var v71 float32 = 3.14159
                                    var v72 uint = 740
                                    _ = v71
                                    _ = v72
                                }
                                _ = fi64
                            }
                        default:
                            if (genFn0(-1, 1, 1) > 0) {
                                fmt.Println(v1, v5)
                                delete(v3, v1)
                                v73 := GenS7{F0: 1}
                                _ = v73
                            }
                        }
                    case -325:
                        var v74 float32 = float32(((1e10 / 0.1) + 0.0))
                        _ = v74
                    case 42:
                        if (!true) {
                            var v75 uint8 = 127
                            v76 := GenS10{F0: 255}
                            v77 := GenS8{F0: 42}
                            _ = v75
                            _ = v76
                            _ = v77
                        } else {
                            if false {
                                v5 -= v5
                                v78 := GenS11{F0: (264 >= 519), F1: (false && false), F2: 32767}
                                for fi79 := 0; fi79 < 10; fi79++ {
                                    v80 := GenS9{F0: int16(fi79), F1: 4294967295, F2: 275}
                                    _ = fi79
                                    _ = v80
                                }
                                _ = v78
                            }
                            for fi81 := 0; fi81 < 16; fi81++ {
                                fmt.Println(fi81, fi81, v2)
                                v3[v1] = 0
                                if true {
                                    if ((v1 + (("foo bar" + v1) + v1)) > v1) {
                                        v2 = 2147483647
                                        v82 := GenS5{F0: 2, F1: (fi81 - (v2 / (((fi47 & v2)) | 1))), F2: 116}
                                        v5 -= float32(1e10)
                                        _ = v82
                                    }
                                    if (1 >= -9223372036854775808) {
                                        switch "" {
                                        case "\n\t":
                                            v5 -= v5
                                        }
                                    }
                                    var v83 int = fi47
                                    _ = v83
                                } else {
                                    v1 = "foo bar"
                                    if false {
                                        var v84 int32 = 2147483647
                                        _ = v84
                                    } else {
                                        v85 := []int{(^(v2))}
                                        if (-142 < 0) {
                                            fmt.Println(v2, fi47)
                                            v86 := GenS8{F0: 974}
                                            _ = v86
                                        }
                                        _ = v85
                                    }
                                    if ((true && false) || false) {
                                        v87 := []int{1, v2, v2, fi81}
                                        _ = v87
                                    } else {
                                        var v88 float64 = -2.25
                                        v4[(((v1 + (v1 + v1)) + "\n\t") + v1)] = true
                                        _ = v88
                                    }
                                }
                                _ = fi81
                            }
                        }
                    default:
                        var v89 int16 = -728
                        _ = v89
                    }
                    _ = v51
                }
                _ = fi47
            }
        } else {
            v90 := GenS1{F0: 127, F1: -1, F2: 1}
            var v91 bool = false
            _ = v90
            _ = v91
        }
        fmt.Println(v1, v2, v5)
    }
    v92 := []int{312, -376, (-(v2)), 32768}
    if len(v92) > 0 {
        v92[0] = (v2 * v2)
    }
    v93 := GenS1{F0: 23, F1: -708, F2: 42}
    _ = len(v3)
    var v94 int16 = 32767
    switch ("foo bar" + "") {
    case "hello":
        if len(v92) >= 2 {
            v92 = v92[1:len(v92)]
        }
    case "":
        switch 384 {
        case 0:
            _ = len(v92)
        default:
            if (!(v2 != -549)) {
                v2 += v2
            } else {
                if (float32(((1.5 * 3.14159) / 0.1)) <= 1.5) {
                    var v95 uint32 = 4294967295
                    for fi96 := 0; fi96 < 20; fi96++ {
                        for fi97 := 0; fi97 < 5; fi97++ {
                            switch (v2 >> ((4294967295) & 63)) {
                            case -1:
                                switch v1 {
                                case "a":
                                    v1 = v1
                                default:
                                    v5 = 1e10
                                }
                            case 4294967295:
                                switch v2 {
                                case -388:
                                    fmt.Println(v5, fi96, v95)
                                case -2147483648:
                                    var v98 string = (v1 + "hello")
                                    _ = v98
                                case 132:
                                    var v99 int64 = -759
                                    _ = v99
                                default:
                                    v5 -= (0.1 * v5)
                                }
                            case 32767:
                                var v100 int32 = 1
                                _ = v100
                            default:
                                if (1.5 < float64(v5)) {
                                    switch v1 {
                                    case "hello":
                                        var v101 float32 = v5
                                        _ = v101
                                    case "foo bar":
                                        var v102 uint8 = 2
                                        _ = v102
                                    }
                                    var v103 uint32 = 1
                                    v104 := v4[v1]
                                    _ = v103
                                    _ = v104
                                }
                            }
                            _ = fi97
                        }
                        for fi105 := 0; fi105 < 18; fi105++ {
                            var v106 uint64 = 1
                            _ = fi105
                            _ = v106
                        }
                        v107 := GenS12{F0: 9223372036854775807, F1: 108}
                        _ = fi96
                        _ = v107
                    }
                    _ = v95
                } else {
                    v92 = append(v92, (v2 - (^(v2))))
                }
                v5 -= (v5 * ((3.14159 - float32(0.1)) / v5))
            }
        }
    }
    var v108 float32 = 1e10
    var v109 uint64 = 0
    switch 2147483647 {
    case 32768:
        v110 := GenS13{F0: 42}
        _ = v110
    case 7:
        var v111 uint8 = 2
        _ = v111
    default:
        v1 = ""
    }
    if false {
        fmt.Println(v1, v94)
        var v112 int64 = 42
        _ = v112
    } else {
        if true {
            v108 += v108
            if true {
                v113 := GenS6{F0: 2147483647, F1: 0}
                _ = v113
            }
            var v114 float64 = (float64(0.1) + float64(v5))
            _ = v114
        } else {
            for fi115 := 0; fi115 < 10; fi115++ {
                if (true || false) {
                    v5 -= -2.25
                    fmt.Println(v108, v2, v2)
                } else {
                    switch v1 {
                    case "foo bar":
                        for fi116 := 0; fi116 < 10; fi116++ {
                            var v117 uint32 = 2147483647
                            _ = fi116
                            _ = v117
                        }
                    }
                    var v118 float32 = (float32(1e10) * ((0.1 / 1.0) - v108))
                    _ = v118
                }
                switch v1 {
                case "\n\t":
                    v1 = ((v1 + "a") + v1)
                case "a":
                    if ((false && true) && false) {
                        v109 = v109
                        _ = len(v92)
                        v119 := GenS10{F0: 32767}
                        _ = v119
                    } else {
                        for fi120 := 0; fi120 < 14; fi120++ {
                            v121 := GenS14{F0: "foo bar", F1: v1}
                            if len(v92) > 0 {
                                v92[0] = 32768
                            }
                            if true {
                                v94 -= 0
                                v122 := GenS12{F0: -9223372036854775808, F1: 584}
                                var v123 float64 = float64(float32(3.14159))
                                _ = v122
                                _ = v123
                            } else {
                                v109 = 279
                                v124 := GenS7{F0: 0}
                                if true {
                                    if true {
                                        if false {
                                            var v125 float64 = (3.14159 - (1.5 / -2.25))
                                            if (1 <= 255) {
                                                v4["a"] = (((((v5 * (-2.25 - v5)) + (v5 + (v108 * v5))) - (float32(1.5) * (v108 - float32(v125)))) == (v5 + (0.1 + 1.5))) || true)
                                                fmt.Println(v1, v94)
                                            } else {
                                                if (!(192 >= 1)) {
                                                    var v126 int32 = -1
                                                    if (42 <= 1) {
                                                        v109 += (v109 % ((v109) | 1))
                                                        v127 := GenS15{F0: 32767, F1: v94, F2: 2147483647}
                                                        v3[v1] = 2147483647
                                                        _ = v127
                                                    }
                                                    _ = v126
                                                } else {
                                                    if true {
                                                        v128 := GenS16{F0: 670, F1: 0}
                                                        v129 := GenS13{F0: 65535}
                                                        v125 = (v125 * v125)
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                }
                                                var v130 uint16 = 42
                                                switch v1 {
                                                case "hello":
                                                    v131 := GenS17{F0: 2}
                                                    _ = v131
                                                default:
                                                    v1 = v1
                                                }
                                                _ = v130
                                            }
                                            _ = v125
                                        }
                                    }
                                    if false {
                                        v5 = v5
                                    }
                                }
                                _ = v124
                            }
                            _ = fi120
                            _ = v121
                        }
                        v109 = v109
                        if false {
                            fmt.Println(v1, fi115, v5)
                        } else {
                            if (true && false) {
                                switch (("" + (v1 + "\n\t")) + v1) {
                                case "hello":
                                    var v132 int = (-(fi115))
                                    _ = v132
                                case "a":
                                    if len(v92) >= 2 {
                                        v92 = v92[1:len(v92)]
                                    }
                                default:
                                    switch v1 {
                                    case "":
                                        switch v1 {
                                        case "hello":
                                            if true {
                                                v1 = v1
                                                if false {
                                                    if false {
                                                        v109 -= 4294967295
                                                    } else {
                                                        if (true && false) {
                                                            v1 = v1
                                                            switch v1 {
                                                            case "a":
                                                                switch v1 {
                                                                case "hello":
                                                                    v108 = (((v5 - ((float32((1.5 + 0.0)) - v5) - ((v5 - v108) - 1e10))) / v108) / (((v5 / v108) - v108) + (v108 + v5)))
                                                                case "foo bar":
                                                                    v133 := GenS15{F0: 32767, F1: 417, F2: 2147483648}
                                                                    _ = v133
                                                                case "\n\t":
                                                                    var v134 uint64 = (v109 % ((v109) | 1))
                                                                    _ = v134
                                                                }
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v135 uint16 = 167
                                                                _ = v135
                                                            default:
                                                                v136 := GenS18{F0: ""}
                                                                _ = v136
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v109 = 18446744073709551615
                                                    var v137 uint32 = 4294967295
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if true {
                                                            v138 := GenS19{F0: 2, F1: 1, F2: v137}
                                                            v108 -= (((1e10 * float32(0.0)) + 0.1) * (3.14159 - v5))
                                                            var v139 uint = 1
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                    case "a":
                                                        if (!true) {
                                                            for fi140 := 0; fi140 < 14; fi140++ {
                                                                var v141 uint64 = (v109 & 42)
                                                                var v142 uint8 = uint8(v137)
                                                                if false {
                                                                    if ((v94 * v94) >= v94) {
                                                                        v143 := GenS15{F0: uint16(v141), F1: v94, F2: (^(v137))}
                                                                        v108 = (v108 * 0.1)
                                                                        v94 = 1
                                                                        v3[((((v1 + v1) + v1) + v1) + (v1 + ((v1 + v1) + v1)))] = 2147483647
                                                                        _ = v143
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        var v144 uint32 = v137
                                                                        var v145 int64 = -281
                                                                        _ = len(v4)
                                                                        _ = v144
                                                                        _ = v145
                                                                    }
                                                                }
                                                                _ = fi140
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                            if len(v92) >= 2 {
                                                                v92 = v92[1:len(v92)]
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v146 := v4[""]
                                                        _ = v146
                                                    }
                                                    _ = v137
                                                }
                                                v3[(("hello" + v1) + v1)] = fi115
                                            } else {
                                                v147 := GenS10{F0: 95}
                                                _ = v147
                                            }
                                        case "a":
                                            _ = len(v3)
                                        case "\n\t":
                                            var v148 float32 = (v5 + v108)
                                            _ = v148
                                        default:
                                            v5 = v108
                                        }
                                    }
                                }
                                v149 := GenS10{F0: uint16(v109)}
                                _ = v149
                            }
                            v94 -= (^(v94))
                            if false {
                                if (127 == -73) {
                                    v150 := GenS16{F0: 2147483647, F1: 65535}
                                    _ = v150
                                }
                            }
                        }
                    }
                }
                _ = len(v92)
                _ = fi115
            }
        }
        switch (v2 | v2) {
        case 941:
            fmt.Println(v2, v108, v108)
        default:
            var v151 int32 = -2147483648
            _ = v151
        }
    }
    v92 = append(v92, (v2 / (((-(v2))) | 1)))
    if len(v92) >= 2 {
        v92 = v92[1:len(v92)]
    }
    fmt.Println(v109)
    if (32767 <= 65535) {
        var v152 uint8 = 128
        if (v109 < 800) {
            switch "foo bar" {
            case "foo bar":
                _ = len(v92)
            case "\n\t":
                v108 -= (v108 - v108)
            }
            if len(v92) > 0 {
                v92[0] = v2
            }
            var v153 string = ((v1 + (v1 + v1)) + (v1 + ""))
            _ = v153
        }
        _ = v152
    } else {
        switch (v2 >> ((42) & 63)) {
        case 32768:
            v108 = v108
        case -534:
            for fi154 := 0; fi154 < 10; fi154++ {
                switch v2 {
                case 9223372036854775807:
                    var v155 int16 = -32768
                    _ = v155
                case -303:
                    for fi156 := 0; fi156 < 19; fi156++ {
                        var v157 int64 = 536
                        var v158 int16 = v94
                        v108 = float32(float64((v5 / v108)))
                        if true {
                            fmt.Println(fi156, v157, v109)
                            var v159 uint16 = 352
                            v2 += fi156
                            _ = v159
                        }
                        v160 := GenS20{F0: -128, F1: -373, F2: 65535}
                        if true {
                            fmt.Println(v108)
                            if ((3.14159 + 1e10) >= (3.14159 + (((float64(float32(0.1)) * 3.14159) - 0.0) / 1.0))) {
                                fmt.Println(v5, fi154)
                                var v161 bool = (true && false)
                                _ = v161
                            } else {
                                v157 = (-(v157))
                            }
                        }
                        v162 := GenS21{F0: -32768}
                        var v163 int64 = 9223372036854775807
                        _ = fi156
                        _ = v157
                        _ = v158
                        _ = v160
                        _ = v162
                        _ = v163
                    }
                default:
                    fmt.Println(v5)
                }
                _ = fi154
            }
        case 4294967295:
            _ = len(v3)
        }
        switch -1 {
        case 32767:
            v2 -= (v2 << ((65535) & 63))
        case -395:
            if len(v92) >= 2 {
                v92 = v92[1:len(v92)]
            }
        }
        v164 := GenS22{F0: 424, F1: 127}
        _ = v164
    }
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v92:", len(v92), v92)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v92
    _ = v93
    _ = v94
}
