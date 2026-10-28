package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 uint16; F1 int64; F2 int16 }
type GenS4 struct { F0 uint8; F1 int32; F2 string }
type GenS5 struct { F0 string; F1 int64; F2 string }
type GenS6 struct { F0 string; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 bool; F1 uint; F2 uint64 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 int64; F1 string; F2 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS13 struct { F0 int; F1 int; F2 uint64 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int64; F1 uint8 }
type GenS16 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 uint32, p1 int64) uint8 {
    p0 = (p0 % ((p0) | 1))
    var v6 int64 = 42
    if false {
        if (p0 <= p0) {
            v7 := GenS1{F0: 1}
            for fi8 := 0; fi8 < 1; fi8++ {
                if false {
                    var v9 uint32 = p0
                    var v10 int64 = (v6 % ((v6) | 1))
                    v11 := []int{fi8}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                _ = fi8
            }
            fmt.Println(p1)
            _ = v7
        } else {
            for fi12 := 0; fi12 < 5; fi12++ {
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 bool = false
                    _ = fi13
                    _ = v14
                }
                p1 += v6
                var v15 int16 = 42
                _ = fi12
                _ = v15
            }
            if (0 == -128) {
                v16 := GenS2{F0: 526, F1: 2147483647, F2: 4294967295}
                if ((725 > 255) && ((1 < 1) && (false && (!false)))) {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        var v18 int = (fi17 + (fi17 | 127))
                        _ = fi17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if (0.1 <= 0.1) {
                            fmt.Println(v6, fi19)
                        } else {
                            var v20 uint = 255
                            _ = v20
                        }
                        v21 := []int{(fi19 / ((fi19) | 1)), fi19, 42, -2147483648}
                        _ = fi19
                        _ = v21
                    }
                }
                switch (("" + "") + ("foo bar" + "a")) {
                case "":
                    var v22 uint = 18446744073709551615
                    _ = v22
                case "a":
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v24 := GenS3{F0: 1, F1: p1, F2: 42}
                        var v25 string = "a"
                        fmt.Println(p0, p0)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case "\n\t":
                    p0 += (p0 << ((1) & 63))
                default:
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 int64 = (p1 ^ (p1 - v6))
                        if true {
                            var v28 int8 = 127
                            fmt.Println(v6)
                            switch fi26 {
                            case 32767:
                                if true {
                                    var v29 int64 = 0
                                    v6 -= -9223372036854775808
                                    var v30 string = (("a" + ("" + "hello")) + ("foo bar" + "hello"))
                                    _ = v29
                                    _ = v30
                                } else {
                                    v6 -= (v27 >> ((397) & 63))
                                }
                            case 127:
                                v31 := []int{(fi26 % ((fi26) | 1)), fi26}
                                _ = v31
                            case -2147483648:
                                switch int(p0) {
                                case 42:
                                    var v32 int16 = 1000
                                    _ = v32
                                case 849:
                                    var v33 uint = 995
                                    _ = v33
                                case 9223372036854775807:
                                    for fi34 := 0; fi34 < 17; fi34++ {
                                        fmt.Println(v27, p0)
                                        switch (fi34 / ((fi34) | 1)) {
                                        case 128:
                                            switch (("\n\t" + "\n\t") + ("" + "hello")) {
                                            case "a":
                                                v35 := []int{}
                                                _ = v35
                                            case "foo bar":
                                                if false {
                                                    if (0 < v28) {
                                                        var v36 float64 = 0.0
                                                        _ = v36
                                                    }
                                                } else {
                                                    v28 = (v28 ^ 127)
                                                    v37 := []int{}
                                                    _ = v37
                                                }
                                            case "hello":
                                                if (((44 >= 255) && (v27 == (v6 >> ((162) & 63)))) && false) {
                                                    var v38 int32 = 0
                                                    v39 := []int{(fi34 % ((fi26) | 1)), -192, 32768}
                                                    v39 = append(v39, int(p0))
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    fmt.Println(v6)
                                                    var v40 int32 = 42
                                                    if (!(!(699 <= (v40 + v40)))) {
                                                        if true {
                                                            p1 += (v27 + (p1 % ((p1) | 1)))
                                                            v41 := []int{(fi26 ^ fi26)}
                                                            _ = v41
                                                        }
                                                        v42 := []int{}
                                                        _ = v42
                                                    } else {
                                                        switch fi34 {
                                                        case 2147483647:
                                                            v43 := GenS4{F0: 1, F1: v40, F2: "foo bar"}
                                                            _ = v43
                                                        }
                                                    }
                                                    _ = v40
                                                }
                                            }
                                        case 32768:
                                            v44 := []int{42, fi34}
                                            if true {
                                                var v45 uint32 = (p0 * (p0 << ((1) & 63)))
                                                _ = v45
                                            }
                                            v27 = p1
                                            if false {
                                                var v46 uint16 = 667
                                                switch (fi34 / ((-513) | 1)) {
                                                case 32767:
                                                    var v47 float32 = 0.1
                                                    v46 += v46
                                                    v47 += 0.0
                                                    _ = v47
                                                default:
                                                    switch 1 {
                                                    case 32767:
                                                        var v48 uint64 = 234
                                                        _ = v48
                                                    case 128:
                                                        if len(v44) > 0 {
                                                            v44[0] = fi34
                                                        }
                                                        v46 += v46
                                                        if (false && (21 >= (v28 | (^(v28))))) {
                                                            if len(v44) > 0 {
                                                                v44[0] = (fi34 % ((fi34) | 1))
                                                            }
                                                            v6 = int64(fi26)
                                                            var v49 uint = 1
                                                            _ = v49
                                                        }
                                                    default:
                                                        p0 += (p0 | (p0 & (p0 ^ uint32(v46))))
                                                    }
                                                }
                                                var v50 int64 = 580
                                                _ = v46
                                                _ = v50
                                            }
                                            v51 := GenS5{F0: "hello", F1: v6, F2: "foo bar"}
                                            fmt.Println(p1, v28)
                                            fmt.Println(v27, v28)
                                            _ = v44
                                            _ = v51
                                        }
                                        _ = fi34
                                    }
                                }
                            default:
                                var v52 uint = 2
                                _ = v52
                            }
                            _ = v28
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                _ = v16
            }
            if ((-1 < 42) || ((42 == p0) || true)) {
                var v53 uint8 = 2
                var v54 int64 = p1
                _ = v53
                _ = v54
            }
        }
        if true {
            for fi55 := 0; fi55 < 4; fi55++ {
                p1 = v6
                if (!(42 < 590)) {
                    fmt.Println(v6, v6)
                    p0 = (p0 & (^(p0)))
                }
                _ = fi55
            }
            var v56 uint16 = 1
            _ = v56
        } else {
            v57 := []int{951, 494, -9223372036854775808, 32767}
            _ = len(v57)
            _ = v57
        }
        for fi58 := 0; fi58 < 10; fi58++ {
            fmt.Println(v6)
            _ = fi58
        }
    } else {
        if (-64 != 42) {
            v6 = 0
        }
        if false {
            if (!true) {
                p1 = (p1 - (v6 & v6))
                v59 := GenS5{F0: "foo bar", F1: v6, F2: "a"}
                v6 += (p1 / (((^(v6))) | 1))
                _ = v59
            }
            var v60 int8 = 69
            fmt.Println(v6, p0, p1)
            _ = v60
        }
        v61 := GenS6{F0: "", F1: 42, F2: 1}
        _ = v61
    }
    switch ("" + "\n\t") {
    case "":
        if false {
            var v62 uint64 = 18446744073709551615
            var v63 int64 = p1
            _ = v62
            _ = v63
        }
    case "foo bar":
        v64 := []int{32767, 785, -2147483648}
        _ = v64
    case "a":
        if (1e10 != 3.14159) {
            v6 += p1
        }
    default:
        p1 += (p1 ^ p1)
    }
    _ = p0
    _ = p1
    _ = v6
    return 87
}
func genFn1(p0 int) string {
    v65 := GenS7{F0: 1, F1: 42}
    for fi66 := 0; fi66 < 1; fi66++ {
        if true {
            if false {
                p0 = 32767
                if (255 > 1) {
                    fmt.Println(fi66, p0)
                }
                p0 = (p0 >> ((603) & 63))
            } else {
                v67 := GenS8{F0: true, F1: 516, F2: 861}
                if false {
                    v68 := []int{32768, (fi66 >> ((121) & 63)), (-(p0)), fi66}
                    if len(v68) >= 2 {
                        v68 = v68[1:len(v68)]
                    }
                    _ = v68
                } else {
                    fmt.Println(fi66, fi66)
                }
                _ = v67
            }
            if true {
                fmt.Println(p0, fi66)
            }
        }
        _ = fi66
    }
    v69 := GenS9{F0: 0, F1: 42}
    if (42 > -128) {
        if true {
            v70 := []int{(-(p0)), 110, p0}
            _ = v70
        }
        var v71 bool = true
        var v72 string = (("foo bar" + "\n\t") + ("foo bar" + "foo bar"))
        _ = v71
        _ = v72
    } else {
        var v73 uint32 = 318
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            fmt.Println(v73, p0)
        default:
            if (2 < 4294967295) {
                var v74 string = "foo bar"
                v74 = v74
                var v75 float32 = ((3.14159 * 1.5) / -2.25)
                _ = v74
                _ = v75
            }
        }
        if (2 > (v73 >> ((42) & 63))) {
            v76 := GenS9{F0: 2147483648, F1: 42}
            _ = v76
        } else {
            switch -2147483648 {
            case -2147483648:
                v73 = v73
            case 32768:
                var v77 uint8 = 14
                _ = v77
            case -1:
                for fi78 := 0; fi78 < 20; fi78++ {
                    fmt.Println(p0, v73)
                    v79 := []int{(fi78 - fi78), -9223372036854775808}
                    v80 := GenS5{F0: ((("\n\t" + "a") + "foo bar") + (("hello" + ("foo bar" + "foo bar")) + "foo bar")), F1: 1, F2: "foo bar"}
                    _ = fi78
                    _ = v79
                    _ = v80
                }
            }
            switch -9223372036854775808 {
            case 2147483647:
                p0 += p0
            case 947:
                v73 = (v73 + (v73 | v73))
            case 247:
                if (false && false) {
                    fmt.Println(p0, v73)
                    var v81 string = "a"
                    var v82 int8 = 42
                    _ = v81
                    _ = v82
                } else {
                    for fi83 := 0; fi83 < 17; fi83++ {
                        var v84 string = "hello"
                        _ = fi83
                        _ = v84
                    }
                    var v85 float32 = -2.25
                    _ = v85
                }
            }
            for fi86 := 0; fi86 < 17; fi86++ {
                switch ("foo bar" + "") {
                case "a":
                    v73 = (^(v73))
                }
                for fi87 := 0; fi87 < 18; fi87++ {
                    v88 := []int{(fi87 >> ((1) & 63))}
                    var v89 int = fi86
                    _ = fi87
                    _ = v88
                    _ = v89
                }
                if true {
                    switch (^(p0)) {
                    case -1:
                        v90 := GenS10{F0: 42, F1: "hello", F2: v73}
                        _ = v90
                    case 32767:
                        fmt.Println(p0, v73, p0)
                    default:
                        switch "hello" {
                        case "\n\t":
                            var v91 int = fi86
                            _ = v91
                        case "a":
                            v92 := GenS11{F0: v73}
                            _ = v92
                        case "hello":
                            v93 := []int{}
                            _ = v93
                        default:
                            if ((!(-643 > 42)) && true) {
                                if (!false) {
                                    p0 = (fi86 * 32767)
                                    for fi94 := 0; fi94 < 15; fi94++ {
                                        if (uint(v73) < 255) {
                                            switch 291 {
                                            case 32767:
                                                v95 := GenS11{F0: 2}
                                                _ = v95
                                            case -550:
                                                if false {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v96 := GenS1{F0: 65535}
                                                        _ = v96
                                                    }
                                                } else {
                                                    v97 := []int{(fi86 / (((p0 * (p0 * 9223372036854775807))) | 1)), fi94}
                                                    var v98 int8 = -1
                                                    if true {
                                                        var v99 uint16 = 1
                                                        _ = v99
                                                    } else {
                                                        p0 = 4294967295
                                                        var v100 uint16 = uint16(v73)
                                                        var v101 int32 = -1
                                                        v98 += v98
                                                        if (!true) {
                                                            var v102 int16 = 70
                                                            _ = v102
                                                        }
                                                        fmt.Println(fi94, p0)
                                                        if len(v97) >= 2 {
                                                            v97 = v97[1:len(v97)]
                                                        }
                                                        if (false || ((!true) || true)) {
                                                            v73 = (v73 / ((0) | 1))
                                                        }
                                                        v101 = (v101 & (v101 / ((-676) | 1)))
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                    _ = v97
                                                    _ = v98
                                                }
                                            case -291:
                                                if (588 != 255) {
                                                    var v103 uint16 = 1
                                                    if false {
                                                        if false {
                                                            switch "a" {
                                                            case "":
                                                                switch 0 {
                                                                case 106:
                                                                    var v104 string = "a"
                                                                    _ = v104
                                                                case 689:
                                                                    v105 := []int{}
                                                                    _ = v105
                                                                }
                                                            default:
                                                                var v106 uint64 = 4294967295
                                                                _ = v106
                                                            }
                                                            v107 := GenS4{F0: 246, F1: 650, F2: "hello"}
                                                            _ = v107
                                                        }
                                                        p0 = (fi86 >> ((255) & 63))
                                                        var v108 float64 = float64((float32(0.0) - float32((1.5 / 1e10))))
                                                        _ = v108
                                                    }
                                                    _ = v103
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            p0 = (p0 / ((fi94) | 1))
                                                        case "":
                                                            if (235 != 42) {
                                                                switch 1 {
                                                                case 128:
                                                                    p0 = (p0 >> ((4294967295) & 63))
                                                                case 4294967295:
                                                                    v73 -= (v73 >> ((1) & 63))
                                                                default:
                                                                    switch ("hello" + "foo bar") {
                                                                    case "hello":
                                                                        if false {
                                                                            var v109 uint16 = 2
                                                                            fmt.Println(p0, fi94)
                                                                            _ = v109
                                                                        } else {
                                                                            var v110 int16 = 400
                                                                            v73 = v73
                                                                            _ = v110
                                                                        }
                                                                    case "\n\t":
                                                                        if (true && false) {
                                                                            var v111 int32 = -1
                                                                            v111 = v111
                                                                            _ = v111
                                                                        }
                                                                    default:
                                                                        switch (^(fi86)) {
                                                                        case 32767:
                                                                            var v112 float64 = 3.14159
                                                                            _ = v112
                                                                        case 9223372036854775807:
                                                                            v73 += (v73 >> ((255) & 63))
                                                                        }
                                                                        v113 := []int{}
                                                                        if (1 != 452) {
                                                                            var v114 int = 128
                                                                            if len(v113) > 0 {
                                                                                v113[0] = fi94
                                                                            }
                                                                            _ = v114
                                                                        } else {
                                                                            if (4294967295 == 2) {
                                                                                p0 = -666
                                                                            }
                                                                            var v115 float32 = (0.1 - float32(((float64(float32((((0.0 * 3.14159) * 1e10) + (1e10 / (((float64((1e10 - 3.14159)) / 3.14159) / 1.0) - (1.5 * (0.0 - 1e10))))))) * 0.1) + 3.14159)))
                                                                            switch (^(p0)) {
                                                                            case 32767:
                                                                                p0 = -9223372036854775808
                                                                            default:
                                                                                p0 = fi94
                                                                            }
                                                                            _ = v115
                                                                        }
                                                                        _ = v113
                                                                    }
                                                                    v116 := GenS5{F0: ("" + ("hello" + (("" + "") + ("foo bar" + ("foo bar" + ""))))), F1: 9223372036854775807, F2: "foo bar"}
                                                                    _ = v116
                                                                }
                                                            }
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(p0)
                                                        var v117 int32 = 42
                                                        switch (^(fi86)) {
                                                        case -9223372036854775808:
                                                            if (!false) {
                                                                fmt.Println(v117, fi94)
                                                            }
                                                        }
                                                        _ = v117
                                                    case "hello":
                                                        v118 := GenS5{F0: "", F1: 42, F2: ("foo bar" + "a")}
                                                        _ = v118
                                                    default:
                                                        var v119 uint64 = 2
                                                        _ = v119
                                                    }
                                                    if (0.0 <= -2.25) {
                                                        var v120 string = ("\n\t" + "")
                                                        v121 := []int{2147483647, (fi86 >> ((2) & 63)), p0, fi94}
                                                        _ = v120
                                                        _ = v121
                                                    } else {
                                                        v122 := GenS10{F0: -9223372036854775808, F1: "foo bar", F2: 1}
                                                        fmt.Println(v73, fi86)
                                                        var v123 float32 = (1.5 + 1e10)
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                    fmt.Println(p0, fi86)
                                                }
                                            default:
                                                var v124 int8 = 100
                                                _ = v124
                                            }
                                        }
                                        switch -945 {
                                        case 594:
                                            v125 := GenS1{F0: 2}
                                            _ = v125
                                        case 0:
                                            if (-788 == -1) {
                                                switch fi94 {
                                                case 32767:
                                                    v73 = 300
                                                default:
                                                    if false {
                                                        fmt.Println(p0, p0)
                                                        v126 := GenS12{F0: 32767, F1: 141, F2: 904}
                                                        fmt.Println(fi94)
                                                        _ = v126
                                                    }
                                                }
                                                v127 := GenS12{F0: 1, F1: genFn0(v73, -1), F2: 1}
                                                v73 -= v73
                                                _ = v127
                                            }
                                        case 9223372036854775807:
                                            var v128 int32 = -837
                                            _ = v128
                                        }
                                        _ = fi94
                                    }
                                    v73 = (^(v73))
                                }
                                if (0.0 <= 0.0) {
                                    switch (^(fi86)) {
                                    case 127:
                                        v73 = (v73 ^ v73)
                                    case -1:
                                        var v129 uint8 = 42
                                        _ = v129
                                    default:
                                        fmt.Println(v73, fi86, v73)
                                    }
                                } else {
                                    var v130 uint64 = 876
                                    for fi131 := 0; fi131 < 7; fi131++ {
                                        p0 = (^(fi131))
                                        _ = fi131
                                    }
                                    _ = v130
                                }
                            }
                        }
                    }
                }
                _ = fi86
            }
        }
        _ = v73
    }
    _ = p0
    _ = v65
    _ = v69
    return ""
}

func main() {
    var v1 int32 = 2147483647
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    v132 := GenS10{F0: int64(v1), F1: "\n\t", F2: 42}
    for fi133 := 0; fi133 < 9; fi133++ {
        var v134 bool = true
        _ = fi133
        _ = v134
    }
    v1 = v1
    v135 := GenS13{F0: 993, F1: 32768, F2: 159}
    v1 -= (^(v1))
    v2 += -9223372036854775808
    if (810 == 42) {
        v1 -= v1
    } else {
        v136 := GenS14{F0: 127}
        v137 := GenS14{F0: -1}
        if true {
            v2 = (v2 - 880)
            var v138 uint8 = 128
            v139 := GenS15{F0: (v2 ^ v2), F1: v138}
            _ = v138
            _ = v139
        }
        _ = v136
        _ = v137
    }
    v2 += v2
    for fi140 := 0; fi140 < 16; fi140++ {
        fmt.Println(fi140)
        _ = fi140
    }
    var v141 int8 = 1
    if true {
        for fi142 := 0; fi142 < 15; fi142++ {
            if false {
                var v143 bool = (0.1 != 0.0)
                _ = v143
            } else {
                if true {
                    v2 -= (v2 - 42)
                }
                var v144 string = ("" + "hello")
                for fi145 := 0; fi145 < 8; fi145++ {
                    var v146 uint32 = 65535
                    var v147 uint32 = 2147483648
                    _ = fi145
                    _ = v146
                    _ = v147
                }
                _ = v144
            }
            var v148 uint16 = 32767
            var v149 uint32 = 2147483648
            _ = fi142
            _ = v148
            _ = v149
        }
    }
    v150 := GenS16{F0: 18446744073709551615, F1: 645, F2: (v2 / ((v2) | 1))}
    var v151 uint = 2
    var v152 bool = (127 == 0)
    fmt.Println("v1:", v1)
    fmt.Println("v141:", v141)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v132
    _ = v135
    _ = v141
    _ = v150
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
