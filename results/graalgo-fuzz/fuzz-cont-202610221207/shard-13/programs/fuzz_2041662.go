package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int16; F2 uint32 }
type GenS8 struct { F0 int64; F1 int64; F2 uint16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 bool }
type GenS12 struct { F0 uint16; F1 int16; F2 int32 }
type GenS13 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS14 struct { F0 uint; F1 string; F2 int8 }
type GenS15 struct { F0 bool; F1 uint64; F2 string }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int16; F1 int32 }
func genFn0(p0 int, p1 int8, p2 int8) uint64 {
    v6 := []int{(p0 << ((2) & 63))}
    var v7 uint32 = 0
    var v8 string = "\n\t"
    if (true && (42 <= -2147483648)) {
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(v7, fi9)
            for fi10 := 0; fi10 < 8; fi10++ {
                switch 1 {
                case -350:
                    if (p1 < 127) {
                        v8 = "a"
                        v6 = append(v6, (fi9 ^ 42))
                        fmt.Println(p1, v7, v7)
                    } else {
                        if true {
                            p2 = p1
                        }
                        var v11 uint8 = 87
                        _ = v11
                    }
                case 1:
                    v12 := GenS1{F0: p0, F1: 9223372036854775807}
                    _ = v12
                case 762:
                    switch "hello" {
                    case "":
                        if (!(0 >= p1)) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    default:
                        var v13 int = p0
                        _ = v13
                    }
                default:
                    if false {
                        v14 := GenS2{F0: 0, F1: fi9}
                        var v15 int32 = 42
                        _ = v14
                        _ = v15
                    }
                }
                v16 := GenS2{F0: 1, F1: 4294967295}
                var v17 string = "hello"
                if (-1 <= -1) {
                    if false {
                        v18 := GenS3{F0: 1, F1: ((p0 / ((fi10) | 1)) < p0)}
                        var v19 uint64 = 18446744073709551615
                        if true {
                            if false {
                                if (float64(3.14159) >= float64(float32(1.5))) {
                                    v20 := GenS2{F0: uint8(v19), F1: (p0 << ((42) & 63))}
                                    p1 = p1
                                    v19 = (v19 + 311)
                                    _ = v20
                                }
                                var v21 int32 = 0
                                v22 := GenS3{F0: 0, F1: false}
                                _ = v21
                                _ = v22
                            } else {
                                p2 = 100
                            }
                            if ((true || true) || (v19 >= (v19 - 42))) {
                                v23 := GenS3{F0: (v19 >> ((42) & 63)), F1: true}
                                fmt.Println(fi10)
                                p2 = (p2 % ((-1) | 1))
                                var v24 uint16 = 803
                                _ = v23
                                _ = v24
                            }
                        }
                        v25 := GenS3{F0: v19, F1: false}
                        v8 = "foo bar"
                        _ = v18
                        _ = v19
                        _ = v25
                    } else {
                        var v26 float64 = 3.14159
                        var v27 uint16 = 65535
                        if false {
                            fmt.Println(v26)
                        } else {
                            var v28 uint = 255
                            if true {
                                switch 0 {
                                case -149:
                                    v29 := GenS2{F0: 128, F1: fi9}
                                    _ = v29
                                case 127:
                                    var v30 float32 = -2.25
                                    _ = v30
                                case 1:
                                    var v31 uint8 = 128
                                    _ = v31
                                }
                            }
                            _ = v28
                        }
                        _ = v26
                        _ = v27
                    }
                    p0 += (fi9 >> ((uint(v7)) & 63))
                } else {
                    var v32 float32 = 3.14159
                    _ = v32
                }
                _ = fi10
                _ = v16
                _ = v17
            }
            if (!false) {
                var v33 bool = (1 != 896)
                p0 -= fi9
                _ = len(v6)
                _ = v33
            }
            _ = fi9
        }
        var v34 uint = 42
        fmt.Println(v7)
        _ = v34
    } else {
        var v35 uint8 = 42
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 65535
    var v2 string = ("hello" + "foo bar")
    v3 := make(map[uint32]string)
    var v4 string = v2
    var v5 uint = 190
    var v36 uint8 = 1
    if true {
        var v37 float64 = 1.5
        _ = v37
    } else {
        if (uint16(v1) >= 330) {
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 uint64 = 1
                for fi40 := 0; fi40 < 5; fi40++ {
                    switch v4 {
                    case "foo bar":
                        v41 := GenS3{F0: v39, F1: true}
                        _ = v41
                    case "hello":
                        v2 = v4
                        if true {
                            v42 := []int{fi38, (-(fi38))}
                            _ = v42
                        } else {
                            var v43 uint64 = 482
                            switch (fi40 / ((fi40) | 1)) {
                            case 32768:
                                v43 = 1
                            case 127:
                                v44 := GenS1{F0: 9223372036854775807, F1: 42}
                                _ = v44
                            }
                            _ = v43
                        }
                        var v45 bool = (v39 != v39)
                        _ = v45
                    case "\n\t":
                        var v46 uint64 = 464
                        _ = v46
                    }
                    v47 := []int{}
                    v2 = ((v2 + "hello") + v2)
                    _ = fi40
                    _ = v47
                }
                _ = fi38
                _ = v39
            }
            var v48 int = 128
            for fi49 := 0; fi49 < 14; fi49++ {
                var v50 bool = false
                _ = fi49
                _ = v50
            }
            _ = v48
        } else {
            if false {
                v51 := GenS1{F0: 32768, F1: 42}
                _ = v51
            } else {
                if false {
                    for fi52 := 0; fi52 < 1; fi52++ {
                        if (!true) {
                            v53 := GenS4{F0: v36, F1: (^(v5)), F2: 32767}
                            _ = v53
                        } else {
                            switch fi52 {
                            case 0:
                                var v54 uint = 18446744073709551615
                                _ = v54
                            case 128:
                                v55 := GenS2{F0: 156, F1: (fi52 & (-(fi52)))}
                                _ = v55
                            }
                            fmt.Println(v2, fi52, v36)
                        }
                        _ = fi52
                    }
                    fmt.Println(v36)
                    fmt.Println(v5, v5, v2)
                }
                if false {
                    var v56 int64 = -180
                    v2 = v2
                    v57 := v3[(v1 - (v1 | 2147483647))]
                    _ = v56
                    _ = v57
                } else {
                    var v58 float32 = 1.5
                    for fi59 := 0; fi59 < 17; fi59++ {
                        v4 = (v4 + ("a" + v2))
                        for fi60 := 0; fi60 < 10; fi60++ {
                            if (((float64(v58) / 1.5) >= 1.5) && true) {
                                v5 = (v5 | 682)
                                var v61 float64 = -2.25
                                _ = v61
                            } else {
                                fmt.Println(v5, v58, v58)
                            }
                            _ = fi60
                        }
                        _ = fi59
                    }
                    for fi62 := 0; fi62 < 3; fi62++ {
                        if true {
                            v63 := GenS3{F0: 227, F1: true}
                            v1 -= (v1 ^ (v1 ^ v1))
                            _ = v63
                        } else {
                            var v64 float32 = 0.0
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 20; fi65++ {
                            v66 := []int{fi65}
                            v67 := GenS5{F0: 603, F1: uint32(v36), F2: fi65}
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                        if false {
                            var v68 int16 = int16(v36)
                            v69 := GenS4{F0: v36, F1: v5, F2: 32767}
                            _ = v68
                            _ = v69
                        }
                        _ = fi62
                    }
                    _ = v58
                }
            }
        }
    }
    v5 -= (v5 | v5)
    v70 := GenS6{F0: (v1 / (((v1 >> (((v5 | v5)) & 63))) | 1))}
    switch v4 {
    case "a":
        if false {
            var v71 int = 1
            v2 = v2
            fmt.Println(v2, v1, v4)
            _ = v71
        } else {
            if false {
                for fi72 := 0; fi72 < 20; fi72++ {
                    if false {
                        if (!false) {
                            switch v2 {
                            case "":
                                fmt.Println(v2, v2, fi72)
                            case "foo bar":
                                v2 = v4
                            default:
                                if (!(2 != 18446744073709551615)) {
                                    v36 -= (v36 + v36)
                                    v73 := GenS3{F0: uint64(v36), F1: false}
                                    _ = v73
                                }
                            }
                            v2 = v4
                        }
                    } else {
                        fmt.Println(v4, v4)
                        if ((^(v1)) != (v1 - v1)) {
                            var v74 int8 = 42
                            _ = v74
                        } else {
                            var v75 int8 = -26
                            _ = v75
                        }
                    }
                    var v76 float64 = ((float64((((0.0 / 1.0) / 1.5) + (3.14159 + 1e10))) / -2.25) - 3.14159)
                    _ = fi72
                    _ = v76
                }
                v77 := GenS3{F0: genFn0(-770, 42, 0), F1: false}
                _ = v77
            } else {
                var v78 uint64 = uint64(v36)
                switch v4 {
                case "hello":
                    _ = len(v3)
                case "\n\t":
                    var v79 uint64 = genFn0(2147483647, 1, 2)
                    _ = v79
                case "foo bar":
                    if false {
                        v36 = (v36 | (^(v36)))
                    }
                }
                v5 = uint(v78)
                _ = v78
            }
        }
    case "foo bar":
        v5 = (v5 + (v5 | v5))
    }
    var v80 bool = (42 >= -9223372036854775808)
    for fi81 := 0; fi81 < 17; fi81++ {
        var v82 string = v2
        v80 = true
        _ = fi81
        _ = v82
    }
    var v83 uint8 = (v36 & 42)
    v84 := []int{2147483647, 0, 9223372036854775807}
    v85 := GenS6{F0: v1}
    for fi86 := 0; fi86 < 10; fi86++ {
        switch ("\n\t" + v2) {
        case "\n\t":
            switch v4 {
            case "hello":
                v87 := GenS7{F0: 100, F1: 32767, F2: (^(v1))}
                _ = v87
            case "":
                var v88 float32 = 3.14159
                _ = v88
            }
        case "hello":
            switch ((v4 + v4) + ((v4 + "") + (v4 + ("hello" + (v2 + "\n\t"))))) {
            case "foo bar":
                switch (-(fi86)) {
                case 32768:
                    v83 = v83
                }
            case "a":
                for fi89 := 0; fi89 < 18; fi89++ {
                    var v90 int = 32767
                    if len(v84) > 0 {
                        v84[0] = int(v83)
                    }
                    v91 := GenS2{F0: v36, F1: 32767}
                    v92 := GenS5{F0: 434, F1: 2, F2: v90}
                    _ = fi89
                    _ = v90
                    _ = v91
                    _ = v92
                }
            default:
                v93 := GenS1{F0: fi86, F1: 1}
                _ = v93
            }
        }
        _ = len(v84)
        _ = fi86
    }
    switch (v4 + v4) {
    case "a":
        if len(v84) >= 2 {
            v84 = v84[1:len(v84)]
        }
    case "":
        v2 = v4
    }
    v3[v1] = v4
    switch "foo bar" {
    case "":
        v94 := GenS8{F0: 42, F1: 1, F2: 65535}
        _ = v94
    case "\n\t":
        v2 = v2
    case "a":
        for fi95 := 0; fi95 < 20; fi95++ {
            if len(v84) >= 2 {
                v84 = v84[1:len(v84)]
            }
            if v80 {
                fmt.Println(v5, v83, v1)
            } else {
                if false {
                    v96 := v3[(^(v1))]
                    switch fi95 {
                    case 127:
                        for fi97 := 0; fi97 < 12; fi97++ {
                            v98 := GenS9{F0: 42}
                            v99 := GenS10{F0: 4294967295}
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                    }
                    _ = v96
                }
                var v100 int16 = 594
                _ = v100
            }
            _ = fi95
        }
    }
    var v101 bool = true
    if (v4 > v4) {
        for fi102 := 0; fi102 < 14; fi102++ {
            var v103 int64 = int64(v36)
            for fi104 := 0; fi104 < 4; fi104++ {
                if v101 {
                    if v80 {
                        fmt.Println(fi102, v101, fi104)
                    } else {
                        fmt.Println(v1, v101, v80)
                        v105 := v3[v1]
                        _ = v105
                    }
                    var v106 bool = false
                    _ = v106
                } else {
                    switch ("" + v4) {
                    case "foo bar":
                        v107 := GenS11{F0: -1, F1: (2147483647 >= 978)}
                        _ = v107
                    case "hello":
                        if (!true) {
                            var v108 uint8 = 1
                            if true {
                                var v109 uint16 = uint16(v108)
                                v83 = (v108 | v108)
                                _ = v109
                            } else {
                                var v110 float64 = -2.25
                                var v111 float32 = 1.5
                                _ = v110
                                _ = v111
                            }
                            _ = v108
                        } else {
                            if (366 >= (v5 + 18446744073709551615)) {
                                var v112 int32 = 42
                                _ = len(v3)
                                fmt.Println(fi102, v36)
                                _ = v112
                            } else {
                                switch fi102 {
                                case 32767:
                                    var v113 float32 = (1.5 - float32(0.1))
                                    _ = v113
                                case 2147483647:
                                    var v114 uint16 = 255
                                    _ = v114
                                default:
                                    if ((v5 ^ 162) >= v5) {
                                        switch (fi104 % (((^(fi104))) | 1)) {
                                        case -1:
                                            switch 813 {
                                            case 217:
                                                fmt.Println(v5, v2, v4)
                                            case 42:
                                                v3[2147483647] = (("foo bar" + ((v2 + v2) + (v2 + v2))) + v4)
                                            case 32767:
                                                v80 = v101
                                            default:
                                                var v115 bool = v101
                                                var v116 int = (fi102 >> ((v5) & 63))
                                                _ = v115
                                                _ = v116
                                            }
                                        case -508:
                                            v103 += v103
                                        }
                                        v4 = ""
                                        v117 := GenS12{F0: 868, F1: -993, F2: 867}
                                        switch "\n\t" {
                                        case "":
                                            var v118 uint8 = (v36 - uint8(v5))
                                            _ = v118
                                        case "\n\t":
                                            var v119 uint8 = v36
                                            v120 := GenS13{F0: 32767, F1: 42, F2: 4294967295}
                                            v80 = (v36 == v119)
                                            v80 = (!false)
                                            v5 -= v5
                                            v103 -= 1
                                            if (700 > int16(v1)) {
                                                v101 = v101
                                                fmt.Println(fi102, v36, v1)
                                                if ((296 == 0) && (genFn0(fi104, 127, 42) != genFn0(fi104, 100, -1))) {
                                                    v83 -= v119
                                                    v5 -= v5
                                                    v5 = 42
                                                } else {
                                                    switch (v2 + v2) {
                                                    case "\n\t":
                                                        fmt.Println(v80, v103)
                                                    default:
                                                        switch (^(fi102)) {
                                                        case 1:
                                                            delete(v3, v1)
                                                        }
                                                    }
                                                }
                                                v121 := GenS14{F0: v5, F1: (v4 + (v2 + "hello")), F2: 42}
                                                _ = v121
                                            } else {
                                                if v80 {
                                                    var v122 int = 42
                                                    v36 = v36
                                                    _ = v122
                                                } else {
                                                    v123 := v3[(v1 ^ (v1 & 2))]
                                                    var v124 int8 = 84
                                                    _ = v123
                                                    _ = v124
                                                }
                                            }
                                            _ = v119
                                            _ = v120
                                        }
                                        var v125 int64 = (v103 & (^(v103)))
                                        _ = v117
                                        _ = v125
                                    }
                                }
                                var v126 int8 = 1
                                v103 -= (v103 ^ (-(v103)))
                                _ = v126
                            }
                            switch v4 {
                            case "foo bar":
                                var v127 uint = v5
                                var v128 float32 = 1e10
                                switch fi104 {
                                case 128:
                                    fmt.Println(v36, v36)
                                case 783:
                                    _ = len(v3)
                                default:
                                    fmt.Println(fi102)
                                }
                                _ = v127
                                _ = v128
                            case "":
                                var v129 int64 = (v103 - 1)
                                _ = v129
                            default:
                                switch fi104 {
                                case 32767:
                                    fmt.Println(v36, v2)
                                case 32768:
                                    v103 -= -581
                                case 4294967295:
                                    if (v101 && true) {
                                        var v130 string = v4
                                        _ = v130
                                    }
                                }
                            }
                            if v80 {
                                v131 := GenS2{F0: (v83 - v36), F1: (fi104 * (fi102 << (((v5 + v5)) & 63)))}
                                fmt.Println(v1)
                                switch "foo bar" {
                                case "":
                                    if ((0.1 / (float64(3.14159) / ((3.14159 / 1.5) + (-2.25 * ((3.14159 + (float64((3.14159 * ((-2.25 - -2.25) / ((-2.25 - 1e10) - 1e10)))) * (1e10 + 1e10))) - (0.1 + 1.5)))))) <= float64(3.14159)) {
                                        switch v2 {
                                        case "\n\t":
                                            var v132 float64 = ((0.1 + ((((3.14159 / -2.25) / 1.5) - 0.0) - 3.14159)) + 0.0)
                                            _ = v132
                                        case "":
                                            v5 = 840
                                        }
                                        v83 += v83
                                    }
                                case "foo bar":
                                    switch "hello" {
                                    case "\n\t":
                                        var v133 float32 = 3.14159
                                        _ = v133
                                    }
                                case "\n\t":
                                    if v101 {
                                        v84 = append(v84, (fi102 & (fi102 ^ fi104)))
                                        v36 = v83
                                        var v134 float64 = (3.14159 - 0.0)
                                        _ = v134
                                    }
                                default:
                                    if v80 {
                                        if v80 {
                                            v36 -= (v36 % ((v83) | 1))
                                            var v135 bool = v101
                                            _ = v135
                                        }
                                        var v136 uint = v5
                                        switch 42 {
                                        case 32767:
                                            var v137 int16 = 0
                                            _ = v137
                                        default:
                                            if (v80 && (!v101)) {
                                                var v138 int32 = 20
                                                fmt.Println(v1, fi104, v101)
                                                _ = v138
                                            }
                                        }
                                        _ = v136
                                    }
                                }
                                if v101 {
                                    fmt.Println(v4, v1)
                                    if (false || v80) {
                                        if true {
                                            var v139 int = (fi104 % ((fi104) | 1))
                                            if (65535 < v1) {
                                                fmt.Println(v80)
                                                var v140 int64 = (v103 * 1)
                                                v1 = v1
                                                _ = v140
                                            } else {
                                                switch int(v83) {
                                                case -873:
                                                    v36 = 2
                                                    switch v2 {
                                                    case "a":
                                                        if (((1e10 + 0.0) != (3.14159 + (1e10 + (((0.0 - (-2.25 + 3.14159)) / 1.0) - 0.0)))) || (-2.25 > ((((-2.25 - (3.14159 + ((3.14159 * float64(1e10)) / 1.5))) + (((1.5 * 1.5) - float64(0.1)) * 3.14159)) * 1.5) * 0.1))) {
                                                            _ = len(v84)
                                                            v103 += v103
                                                        } else {
                                                            v141 := v3[(^(v1))]
                                                            _ = v141
                                                        }
                                                    default:
                                                        if v101 {
                                                            v142 := GenS7{F0: 0, F1: 1, F2: v1}
                                                            _ = v142
                                                        } else {
                                                            if (v103 != 1) {
                                                                var v143 int8 = -1
                                                                v143 += (v143 << ((v5) & 63))
                                                                _ = v143
                                                            }
                                                        }
                                                    }
                                                    var v144 int16 = -1
                                                    var v145 uint16 = 32767
                                                    _ = v144
                                                    _ = v145
                                                case 0:
                                                    if v101 {
                                                        var v146 int32 = 2147483647
                                                        if (!v101) {
                                                            fmt.Println(fi102, v5, v139)
                                                            fmt.Println(v83, v36)
                                                            var v147 int32 = (v146 ^ v146)
                                                            v83 += v83
                                                            _ = v147
                                                        }
                                                        if true {
                                                            switch v139 {
                                                            case 127:
                                                                if (((fi102 * (v139 + (fi104 * 25))) == fi102) || false) {
                                                                    v148 := GenS11{F0: v146, F1: false}
                                                                    var v149 uint = 255
                                                                    _ = v148
                                                                    _ = v149
                                                                }
                                                            case 1:
                                                                var v150 int64 = 0
                                                                _ = v150
                                                            default:
                                                                if len(v84) >= 2 {
                                                                    v84 = v84[1:len(v84)]
                                                                }
                                                            }
                                                        } else {
                                                            if v80 {
                                                                var v151 uint64 = uint64(v1)
                                                                if len(v84) > 0 {
                                                                    v84[0] = -1
                                                                }
                                                                v151 = 4294967295
                                                                v139 = int(v151)
                                                                _ = v151
                                                            } else {
                                                                v152 := GenS15{F0: v80, F1: 18446744073709551615, F2: v2}
                                                                if len(v84) > 0 {
                                                                    v84[0] = -9223372036854775808
                                                                }
                                                                if v80 {
                                                                    if len(v84) > 0 {
                                                                        v84[0] = 42
                                                                    }
                                                                    v153 := GenS15{F0: v80, F1: 913, F2: v4}
                                                                    v154 := GenS3{F0: 993, F1: ((v36 * v83) != (v36 << ((v5) & 63)))}
                                                                    switch (fi102 & v139) {
                                                                    case 262:
                                                                        var v155 int8 = 42
                                                                        _ = v155
                                                                    }
                                                                    v156 := GenS16{F0: (^(v83))}
                                                                    _ = v153
                                                                    _ = v154
                                                                    _ = v156
                                                                }
                                                                _ = v152
                                                            }
                                                        }
                                                        switch v4 {
                                                        case "hello":
                                                            v157 := GenS17{F0: -420, F1: (-(v146))}
                                                            _ = v157
                                                        case "foo bar":
                                                            var v158 uint8 = v36
                                                            _ = v158
                                                        case "":
                                                            var v159 uint8 = v83
                                                            _ = v159
                                                        default:
                                                            v83 -= v36
                                                        }
                                                        _ = v146
                                                    }
                                                case 673:
                                                    if v80 {
                                                        v160 := GenS11{F0: 2147483647, F1: v80}
                                                        v4 = (((v4 + v2) + ((("foo bar" + (("a" + "\n\t") + "")) + v2) + v2)) + "foo bar")
                                                        _ = v160
                                                    }
                                                default:
                                                    var v161 bool = true
                                                    _ = v161
                                                }
                                            }
                                            _ = v139
                                        } else {
                                            v162 := GenS11{F0: 0, F1: (v80 || (v101 && v80))}
                                            var v163 int32 = 993
                                            _ = len(v84)
                                            if ((42 >= -1) || (2 >= 42)) {
                                                fmt.Println(v36, v103)
                                                if len(v84) >= 2 {
                                                    v84 = v84[1:len(v84)]
                                                }
                                                v103 = (v103 - int64(v36))
                                            } else {
                                                v164 := GenS4{F0: (v83 >> ((v5) & 63)), F1: (^(v5)), F2: 65535}
                                                if true {
                                                    switch v2 {
                                                    case "a":
                                                        v165 := GenS10{F0: genFn0(32767, -128, -96)}
                                                        _ = v165
                                                    case "hello":
                                                        fmt.Println(v1, v80, v83)
                                                    default:
                                                        if (v1 >= 2) {
                                                            v166 := GenS9{F0: 100}
                                                            delete(v3, v1)
                                                            var v167 int8 = 87
                                                            _ = v166
                                                            _ = v167
                                                        } else {
                                                            var v168 uint = 18446744073709551615
                                                            var v169 int8 = 38
                                                            _ = v168
                                                            _ = v169
                                                        }
                                                    }
                                                    v103 -= v103
                                                    switch (fi102 >> ((878) & 63)) {
                                                    case -2147483648:
                                                        v103 += (v103 & 42)
                                                        if v80 {
                                                            if v101 {
                                                                var v170 float32 = 1e10
                                                                fmt.Println(v170)
                                                                var v171 int64 = (v103 | (^(v103)))
                                                                _ = v170
                                                                _ = v171
                                                            }
                                                        }
                                                        var v172 int64 = (-(v103))
                                                        _ = v172
                                                    case 685:
                                                        var v173 uint32 = v1
                                                        _ = v173
                                                    case 128:
                                                        if (false || (-119 <= 1)) {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(v101, v103)
                                                            case "foo bar":
                                                                var v174 int = -1
                                                                _ = v174
                                                            case "":
                                                                var v175 float32 = 0.0
                                                                _ = v175
                                                            }
                                                        } else {
                                                            if (v163 <= (v163 ^ (v163 | -1))) {
                                                                switch v4 {
                                                                case "foo bar":
                                                                    v176 := GenS4{F0: v36, F1: v5, F2: 255}
                                                                    _ = v176
                                                                default:
                                                                    if (1.5 != (0.0 * (1e10 + 3.14159))) {
                                                                        v103 -= 9223372036854775807
                                                                    }
                                                                }
                                                                var v177 bool = v101
                                                                _ = v177
                                                            } else {
                                                                var v178 int8 = -1
                                                                _ = v178
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    switch fi104 {
                                                    case -90:
                                                        v101 = (42 >= 255)
                                                    case 127:
                                                        v36 = (^(v83))
                                                    case 128:
                                                        if ((v163 ^ (v163 >> (((v5 / ((65535) | 1))) & 63))) < (v163 * v163)) {
                                                            var v179 int32 = v163
                                                            _ = v179
                                                        } else {
                                                            _ = len(v84)
                                                            v180 := GenS17{F0: 386, F1: (v163 ^ (v163 >> ((626) & 63)))}
                                                            var v181 uint = v5
                                                            _ = v180
                                                            _ = v181
                                                        }
                                                    default:
                                                        var v182 float32 = 0.1
                                                        _ = v182
                                                    }
                                                    _ = len(v84)
                                                    v183 := GenS2{F0: v83, F1: fi102}
                                                    if (v4 < (v2 + v4)) {
                                                        v36 = uint8(v5)
                                                    } else {
                                                        v184 := GenS1{F0: fi102, F1: 0}
                                                        var v185 uint64 = genFn0(128, 1, 42)
                                                        _ = v184
                                                        _ = v185
                                                    }
                                                    _ = v183
                                                }
                                                var v186 int16 = 42
                                                _ = v164
                                                _ = v186
                                            }
                                            v187 := GenS11{F0: (v163 & -1), F1: v101}
                                            _ = v162
                                            _ = v163
                                            _ = v187
                                        }
                                        fmt.Println(fi102)
                                    }
                                    var v188 int = (fi102 ^ fi102)
                                    v5 = v5
                                    _ = v188
                                }
                                _ = v131
                            }
                        }
                    }
                }
                _ = fi104
            }
            var v189 uint16 = 99
            _ = fi102
            _ = v103
            _ = v189
        }
    } else {
        fmt.Println(v101)
        var v190 uint64 = 18446744073709551615
        switch "\n\t" {
        case "hello":
            for fi191 := 0; fi191 < 3; fi191++ {
                fmt.Println(v1)
                _ = fi191
            }
        default:
            v3[4294967295] = v4
        }
        _ = v190
    }
    v192 := v3[uint32(v5)]
    v2 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v192:", v192)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("v83:", v83)
    fmt.Println("slicelen v84:", len(v84), v84)
    _ = v1
    _ = v101
    _ = v192
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v70
    _ = v80
    _ = v83
    _ = v84
    _ = v85
}
