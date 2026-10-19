package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int; F2 bool }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 string; F2 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int16; F1 uint; F2 int }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int64; F1 int32; F2 int64 }
type GenS17 struct { F0 int32; F1 int }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int; F1 int; F2 uint8 }
type GenS21 struct { F0 int8; F1 int16 }
type GenS22 struct { F0 string; F1 int8; F2 bool }
type GenS23 struct { F0 uint32; F1 int32; F2 uint }
type GenS24 struct { F0 bool; F1 uint8; F2 uint }
type GenS25 struct { F0 string }
func genFn0(p0 uint32, p1 int32, p2 uint16) int32 {
    v7 := GenS1{F0: 9223372036854775807, F1: 1, F2: -32768}
    fmt.Println(p0)
    v8 := GenS1{F0: 0, F1: 9223372036854775807, F2: 42}
    var v9 float32 = (1.5 / (((1.5 + float32(-2.25)) - 1.5) - 1e10))
    v10 := []int{-503, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p1 - p1)
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (3.14159 * 3.14159)
    var v3 bool = true
    var v4 bool = true
    var v5 int8 = 100
    var v6 float64 = 1.5
    for fi11 := 0; fi11 < 20; fi11++ {
        v5 -= -128
        _ = len(v1)
        _ = fi11
    }
    var v12 uint64 = 756
    switch "a" {
    case "\n\t":
        for fi13 := 0; fi13 < 15; fi13++ {
            switch ("\n\t" + ("a" + ("a" + "a"))) {
            case "":
                var v14 int = fi13
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS2{F0: 42}
                    switch (fi15 ^ fi15) {
                    case 2147483647:
                        var v17 uint8 = 42
                        v4 = (v4 && false)
                        fmt.Println(v3)
                        if (false || (!v3)) {
                            v3 = (true || (!v4))
                            v6 = ((1.5 / (0.1 - (v6 / 3.14159))) - float64(((float32(float64(0.0)) - 0.1) - (-2.25 - ((-2.25 - 1e10) - ((0.0 - 0.1) + float32(v6)))))))
                            v3 = v3
                        }
                        _ = v17
                    case 32768:
                        v2 -= (float64((1.5 - float32((v2 + float64(0.0))))) * (v6 + float64(float32(float64((0.1 * ((float32(1.5) / -2.25) + 1.5)))))))
                    default:
                        if v3 {
                            fmt.Println(fi15, fi13)
                        } else {
                            v18 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807, F2: 0}
                            _ = v18
                        }
                    }
                    v19 := []int{(fi13 >> ((65535) & 63)), 32767, fi15}
                    if (((-1 == 803) && v4) && v3) {
                        if (-1 >= 42) {
                            v20 := GenS3{F0: "", F1: fi15, F2: (285 <= 42)}
                            _ = v20
                        } else {
                            v12 = (v12 & v12)
                            v19 = append(v19, (fi15 * (fi13 & -2147483648)))
                        }
                        var v21 int = (fi13 / ((4294967295) | 1))
                        _ = v21
                    }
                    v22 := GenS4{F0: 2, F1: 1}
                    var v23 uint16 = 2
                    _ = fi15
                    _ = v16
                    _ = v19
                    _ = v22
                    _ = v23
                }
            case "foo bar":
                if (("foo bar" + "foo bar") >= "hello") {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        if (18446744073709551615 < 2) {
                            if false {
                                var v25 float32 = (0.0 / 0.1)
                                _ = v25
                            }
                        }
                        var v26 int16 = 42
                        _ = fi24
                        _ = v26
                    }
                } else {
                    v12 = (v12 * 1)
                    if (((("\n\t" + "") + "hello") + ("\n\t" + (("\n\t" + "hello") + "foo bar"))) > "foo bar") {
                        fmt.Println(fi13)
                    } else {
                        if v3 {
                            if false {
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    var v28 int = (fi27 | (fi27 - fi13))
                                    if true {
                                        v12 = v12
                                    }
                                    fmt.Println(v5)
                                    _ = fi27
                                    _ = v28
                                }
                            }
                        }
                        switch "hello" {
                        case "a":
                            fmt.Println(v12, v12, v4)
                        case "foo bar":
                            if false {
                                if (967 < 4294967295) {
                                    for fi29 := 0; fi29 < 10; fi29++ {
                                        if v3 {
                                            switch ("" + "hello") {
                                            case "a":
                                                v30 := GenS5{F0: 32767}
                                                _ = v30
                                            default:
                                                if (!true) {
                                                    fmt.Println(v6, v12, v2)
                                                    v31 := GenS6{F0: fi29}
                                                    if (705 == -9223372036854775808) {
                                                        v32 := GenS4{F0: 413, F1: 42}
                                                        var v33 int8 = 123
                                                        v6 += 3.14159
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        v4 = (v6 < 0.1)
                                                        v34 := GenS7{F0: int64(fi13)}
                                                        v35 := v1[(fi29 & (fi13 % ((42) | 1)))]
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    delete(v1, (fi13 * 128))
                                                    _ = v31
                                                }
                                            }
                                        } else {
                                            var v36 uint32 = 2147483648
                                            var v37 uint32 = (^(v36))
                                            _ = v36
                                            _ = v37
                                        }
                                        v38 := []int{32768, fi29}
                                        _ = fi29
                                        _ = v38
                                    }
                                    var v39 uint64 = (^(v12))
                                    switch fi13 {
                                    case 4294967295:
                                        v6 = v6
                                    case 748:
                                        fmt.Println(v5, fi13, v2)
                                    case 9223372036854775807:
                                        v40 := GenS4{F0: 32767, F1: 9223372036854775807}
                                        _ = v40
                                    default:
                                        fmt.Println(v12, v5)
                                    }
                                    _ = v39
                                } else {
                                    for fi41 := 0; fi41 < 6; fi41++ {
                                        var v42 int = fi41
                                        v43 := []int{1, fi13}
                                        var v44 uint16 = 32767
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    v3 = true
                                }
                                var v45 uint32 = 339
                                _ = v45
                            } else {
                                var v46 bool = true
                                _ = v46
                            }
                        case "\n\t":
                            v12 -= (v12 - v12)
                        default:
                            v47 := GenS7{F0: 779}
                            _ = v47
                        }
                        var v48 uint32 = 753
                        _ = v48
                    }
                }
            }
            for fi49 := 0; fi49 < 1; fi49++ {
                if (!false) {
                    v4 = true
                }
                var v50 int = (^(fi13))
                var v51 int8 = (v5 | (^(v5)))
                if (2 <= 255) {
                    fmt.Println(fi49, v2, v12)
                    fmt.Println(fi49)
                }
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = fi13
        }
    }
    fmt.Println(v2, v3, v4)
    switch 9223372036854775807 {
    case 2147483647:
        if (0 <= 901) {
            switch ("\n\t" + "a") {
            case "":
                if false {
                    v52 := GenS8{F0: -1, F1: ("" + "\n\t"), F2: 1}
                    fmt.Println(v6, v6, v2)
                    _ = v52
                }
            case "foo bar":
                var v53 float64 = (v6 * v2)
                _ = v53
            case "a":
                var v54 int8 = v5
                _ = v54
            default:
                delete(v1, 522)
            }
            v55 := GenS9{F0: 4294967295}
            _ = v55
        } else {
            switch "hello" {
            case "foo bar":
                v56 := v1[4294967295]
                _ = v56
            default:
                var v57 float64 = v6
                _ = v57
            }
        }
    case 526:
        switch 32767 {
        case 2147483647:
            v58 := v1[-9223372036854775808]
            _ = v58
        case 9223372036854775807:
            var v59 bool = false
            _ = v59
        case 32767:
            v60 := GenS10{F0: "foo bar"}
            _ = v60
        }
    }
    v12 = (v12 | 42)
    if v4 {
        v2 = ((((v2 - v6) * v6) + 1.5) / float64((3.14159 / 1.0)))
        var v61 int64 = 1
        var v62 uint32 = 733
        _ = v61
        _ = v62
    }
    v12 = v12
    if false {
        if false {
            var v63 int8 = (v5 | v5)
            v64 := GenS11{F0: 214}
            var v65 string = ""
            _ = v63
            _ = v64
            _ = v65
        } else {
            switch "hello" {
            case "":
                v66 := GenS12{F0: (v12 >> ((255) & 63))}
                _ = v66
            default:
                for fi67 := 0; fi67 < 9; fi67++ {
                    v68 := GenS7{F0: 1}
                    v2 = ((3.14159 - 1.5) * (v2 / 1.5))
                    _ = fi67
                    _ = v68
                }
            }
            if v4 {
                var v69 int = 0
                _ = v69
            } else {
                v70 := []int{}
                _ = v70
            }
            v71 := GenS10{F0: "a"}
            _ = v71
        }
    } else {
        v72 := GenS13{F0: (-(v5))}
        if v4 {
            var v73 int32 = 1
            for fi74 := 0; fi74 < 7; fi74++ {
                if (v6 <= ((-2.25 - float64(1.5)) / 1e10)) {
                    for fi75 := 0; fi75 < 19; fi75++ {
                        switch (-(fi75)) {
                        case -2147483648:
                            v6 -= float64(1.5)
                        case 42:
                            delete(v1, fi74)
                        }
                        v5 = v5
                        v76 := []int{fi75, (fi74 ^ (fi74 % ((128) | 1))), fi75, (fi74 << ((255) & 63))}
                        _ = fi75
                        _ = v76
                    }
                    for fi77 := 0; fi77 < 19; fi77++ {
                        if (!((v5 / ((v5) | 1)) == v5)) {
                            v2 -= (((0.0 / 1.5) - float64(((0.1 + (3.14159 * (3.14159 * 0.1))) + (1.5 + (0.0 / 1.0))))) * v6)
                            var v78 int16 = -32768
                            var v79 int = (fi74 % ((232) | 1))
                            _ = v78
                            _ = v79
                        } else {
                            v80 := []int{}
                            _ = v80
                        }
                        _ = fi77
                    }
                }
                var v81 int8 = int8(v12)
                switch (fi74 | (fi74 << ((255) & 63))) {
                case -1:
                    var v82 uint64 = 1
                    _ = v82
                }
                _ = fi74
                _ = v81
            }
            fmt.Println(v73)
            _ = v73
        }
        v12 -= (v12 ^ v12)
        _ = v72
    }
    var v83 uint32 = 848
    var v84 float64 = (v6 - 3.14159)
    var v85 bool = (true || v3)
    v86 := GenS5{F0: 1}
    _ = len(v1)
    for fi87 := 0; fi87 < 20; fi87++ {
        v2 += 0.0
        v6 = 3.14159
        switch "hello" {
        case "foo bar":
            fmt.Println(fi87)
        case "a":
            v88 := GenS14{F0: -578, F1: 345, F2: 2147483647}
            _ = v88
        }
        _ = fi87
    }
    for fi89 := 0; fi89 < 1; fi89++ {
        v4 = false
        _ = fi89
    }
    if (4294967295 <= 941) {
        for fi90 := 0; fi90 < 19; fi90++ {
            v83 = 42
            _ = fi90
        }
        switch 2147483647 {
        case -1:
            var v91 float32 = 0.1
            _ = v91
        case 411:
            v84 += (((3.14159 * 3.14159) * v6) + 0.1)
        case 2147483647:
            if ((9223372036854775807 != 0) && (288 < -543)) {
                var v92 uint8 = 128
                v93 := GenS1{F0: 1, F1: 422, F2: -32768}
                var v94 uint8 = v92
                _ = v92
                _ = v93
                _ = v94
            } else {
                v95 := GenS15{F0: 32768}
                for fi96 := 0; fi96 < 12; fi96++ {
                    var v97 int64 = -236
                    switch "hello" {
                    case "hello":
                        if v4 {
                            v12 = v12
                        } else {
                            for fi98 := 0; fi98 < 20; fi98++ {
                                v5 -= (-(v5))
                                _ = fi98
                            }
                            for fi99 := 0; fi99 < 16; fi99++ {
                                var v100 int32 = -825
                                switch "\n\t" {
                                case "":
                                    v101 := GenS16{F0: 9223372036854775807, F1: -1, F2: v97}
                                    _ = v101
                                case "\n\t":
                                    v84 = v6
                                    if false {
                                        var v102 bool = (42 != 4294967295)
                                        v103 := GenS14{F0: 1000, F1: 748, F2: fi96}
                                        if (0 == int16(v83)) {
                                            switch "hello" {
                                            case "\n\t":
                                                switch "a" {
                                                case "foo bar":
                                                    var v104 uint = 65535
                                                    _ = v104
                                                case "":
                                                    if v4 {
                                                        var v105 float32 = (3.14159 / (((1e10 + 1e10) / 1.5) / (-2.25 * -2.25)))
                                                        _ = v105
                                                    }
                                                default:
                                                    v106 := GenS16{F0: (^(v97)), F1: v100, F2: 9223372036854775807}
                                                    var v107 int32 = genFn0(v83, 0, 883)
                                                    _ = v106
                                                    _ = v107
                                                }
                                            default:
                                                fmt.Println(fi99)
                                                var v108 uint64 = v12
                                                v109 := []int{fi96}
                                                _ = v108
                                                _ = v109
                                            }
                                            switch (("foo bar" + ("a" + ("" + ("\n\t" + "")))) + ("hello" + "foo bar")) {
                                            case "\n\t":
                                                var v110 uint64 = (v12 >> ((0) & 63))
                                                _ = v110
                                            case "a":
                                                v97 = v97
                                            }
                                            if (0 > 171) {
                                                var v111 uint16 = 214
                                                _ = v111
                                            }
                                        }
                                        _ = v102
                                        _ = v103
                                    } else {
                                        var v112 bool = ("\n\t" != "")
                                        if (!(128 < 1)) {
                                            var v113 uint = 2
                                            switch (fi96 & (fi99 % ((fi99) | 1))) {
                                            case 2147483647:
                                                fmt.Println(v5, v83, fi96)
                                            case 1:
                                                switch (-(fi96)) {
                                                case 42:
                                                    v112 = v4
                                                    if true {
                                                        var v114 float64 = v84
                                                        if v85 {
                                                            var v115 bool = (!v3)
                                                            var v116 float64 = v114
                                                            _ = v115
                                                            _ = v116
                                                        } else {
                                                            var v117 float64 = v2
                                                            _ = v117
                                                        }
                                                        _ = v114
                                                    } else {
                                                        switch (("foo bar" + (("a" + "foo bar") + "a")) + "\n\t") {
                                                        case "hello":
                                                            v3 = (!(v3 && (1 <= 101)))
                                                        case "":
                                                            fmt.Println(v6)
                                                            v4 = (v97 < v97)
                                                        case "\n\t":
                                                            if v85 {
                                                                var v118 int = 272
                                                                if (v85 || v85) {
                                                                    switch "" {
                                                                    case "a":
                                                                        if (!v85) {
                                                                            var v119 uint16 = 32767
                                                                            _ = v119
                                                                        }
                                                                    case "hello":
                                                                        var v120 int16 = -32768
                                                                        _ = v120
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v112 = (v4 && true)
                                                                        var v121 uint64 = 892
                                                                        var v122 int64 = v97
                                                                        _ = v121
                                                                        _ = v122
                                                                    default:
                                                                        v123 := GenS17{F0: v100, F1: 128}
                                                                        _ = v123
                                                                    }
                                                                    _ = len(v1)
                                                                } else {
                                                                    v12 -= (v12 * (v12 + 4294967295))
                                                                    if false {
                                                                        var v124 int8 = (v5 + (^(v5)))
                                                                        switch 605 {
                                                                        case -16:
                                                                            if ((4294967295 != (v12 & v12)) || v4) {
                                                                                v118 = v118
                                                                                fmt.Println(v3)
                                                                            } else {
                                                                                var v125 uint32 = v83
                                                                                switch 0 {
                                                                                case 32767:
                                                                                    fmt.Println(v85, fi99, v124)
                                                                                case -2147483648:
                                                                                    var v126 uint8 = 127
                                                                                    _ = v126
                                                                                case 127:
                                                                                    var v127 bool = (v6 > ((float64(3.14159) + float64((((0.1 / (1.5 + 0.1)) * ((1e10 / 1.5) - (0.0 * -2.25))) - (float32((v2 + 3.14159)) * (1.5 + 0.0))))) / v6))
                                                                                    _ = v127
                                                                                }
                                                                                _ = v125
                                                                            }
                                                                        }
                                                                        fmt.Println(v3)
                                                                        v128 := GenS5{F0: 2}
                                                                        v129 := GenS18{F0: 315}
                                                                        _ = v124
                                                                        _ = v128
                                                                        _ = v129
                                                                    }
                                                                    v130 := []int{}
                                                                    _ = v130
                                                                }
                                                                _ = v118
                                                            } else {
                                                                var v131 bool = true
                                                                if (v4 || (!v3)) {
                                                                    v132 := GenS14{F0: 0, F1: (v113 ^ 1), F2: fi99}
                                                                    v133 := GenS11{F0: 1}
                                                                    v97 += (v97 | (v97 >> (((^(v113))) & 63)))
                                                                    _ = v132
                                                                    _ = v133
                                                                } else {
                                                                    var v134 bool = (!v112)
                                                                    _ = v134
                                                                }
                                                                if v112 {
                                                                    v135 := GenS8{F0: 1000, F1: "\n\t", F2: (v97 & v97)}
                                                                    _ = v135
                                                                } else {
                                                                    fmt.Println(v12, fi96, v84)
                                                                    if v4 {
                                                                        if v131 {
                                                                            var v136 uint64 = v12
                                                                            _ = v136
                                                                        }
                                                                        v131 = v4
                                                                    } else {
                                                                        v137 := []int{}
                                                                        var v138 int16 = int16(v97)
                                                                        fmt.Println(v6)
                                                                        _ = v137
                                                                        _ = v138
                                                                    }
                                                                    var v139 int64 = (v97 & 166)
                                                                    _ = v139
                                                                }
                                                                _ = v131
                                                            }
                                                        }
                                                    }
                                                case 32767:
                                                    var v140 uint32 = v83
                                                    _ = v140
                                                default:
                                                    v1[fi99] = ("" + "foo bar")
                                                }
                                            default:
                                                if (-1 <= (v100 << ((v113) & 63))) {
                                                    v141 := GenS19{F0: v12}
                                                    var v142 float32 = float32(((-2.25 + v84) / v6))
                                                    switch ((("hello" + "hello") + "foo bar") + "foo bar") {
                                                    case "hello":
                                                        v143 := []int{fi99, fi96, -881}
                                                        _ = v143
                                                    }
                                                    var v144 string = "foo bar"
                                                    _ = v141
                                                    _ = v142
                                                    _ = v144
                                                }
                                            }
                                            v112 = v85
                                            _ = v113
                                        }
                                        v4 = (2 > 863)
                                        if true {
                                            fmt.Println(v3)
                                            var v145 int64 = 9223372036854775807
                                            v146 := []int{}
                                            _ = v145
                                            _ = v146
                                        } else {
                                            if true {
                                                if v85 {
                                                    fmt.Println(fi99, v4)
                                                    v147 := []int{(fi99 << ((795) & 63)), 0, fi99}
                                                    fmt.Println(v12)
                                                    _ = v147
                                                }
                                                v83 -= 65535
                                            } else {
                                                switch (fi99 ^ 876) {
                                                case -164:
                                                    var v148 string = ("a" + ("a" + ""))
                                                    _ = v148
                                                case 9223372036854775807:
                                                    fmt.Println(fi96, v6, v97)
                                                case 0:
                                                    var v149 float64 = float64((0.0 * 0.0))
                                                    _ = v149
                                                default:
                                                    switch (fi99 / (((fi96 / ((705) | 1))) | 1)) {
                                                    case -9223372036854775808:
                                                        fmt.Println(v12, v6, v83)
                                                    case -232:
                                                        if (!(!(!v4))) {
                                                            _ = len(v1)
                                                        } else {
                                                            switch ("\n\t" + ("foo bar" + "a")) {
                                                            case "foo bar":
                                                                v150 := []int{-9223372036854775808, (-(fi96)), -9223372036854775808}
                                                                _ = v150
                                                            case "":
                                                                v84 = ((float64(0.0) + v84) / (((v84 + -2.25) / (float64(-2.25) + float64((0.1 + float32(0.1))))) / v6))
                                                            case "hello":
                                                                var v151 uint16 = 0
                                                                _ = v151
                                                            }
                                                            if false {
                                                                var v152 int16 = int16(v100)
                                                                _ = v152
                                                            } else {
                                                                v6 += float64(1e10)
                                                                v5 = 42
                                                                v85 = (!(v3 || (42 >= 99)))
                                                            }
                                                            v97 -= (v97 + (v97 ^ v97))
                                                        }
                                                    case 9223372036854775807:
                                                        var v153 int64 = (v97 >> ((0) & 63))
                                                        _ = v153
                                                    default:
                                                        fmt.Println(v83, v3, v12)
                                                    }
                                                }
                                                v1[(^(fi96))] = "\n\t"
                                            }
                                            v154 := GenS20{F0: fi99, F1: (^(fi96)), F2: 2}
                                            v155 := []int{fi99}
                                            _ = v154
                                            _ = v155
                                        }
                                        if v112 {
                                            var v156 uint8 = 42
                                            var v157 uint16 = 42
                                            _ = v156
                                            _ = v157
                                        } else {
                                            switch 1 {
                                            case 296:
                                                v158 := GenS21{F0: -84, F1: 42}
                                                _ = v158
                                            case 42:
                                                var v159 int8 = 100
                                                _ = v159
                                            default:
                                                v97 = 42
                                            }
                                        }
                                        _ = v112
                                    }
                                    v85 = (!("\n\t" <= "hello"))
                                }
                                var v160 float32 = (0.0 / float32((v84 + (1.5 * float64(0.0)))))
                                _ = fi99
                                _ = v100
                                _ = v160
                            }
                        }
                    default:
                        v161 := []int{-44, fi96}
                        _ = v161
                    }
                    for fi162 := 0; fi162 < 8; fi162++ {
                        v163 := GenS22{F0: "", F1: (v5 << ((4294967295) & 63)), F2: v3}
                        if v3 {
                            fmt.Println(v84, v83, v84)
                        } else {
                            v164 := GenS23{F0: v83, F1: 0, F2: 588}
                            _ = v164
                        }
                        var v165 int32 = 251
                        _ = fi162
                        _ = v163
                        _ = v165
                    }
                    _ = fi96
                    _ = v97
                }
                _ = v95
            }
        }
        switch 42 {
        case -9223372036854775808:
            v85 = false
        case 9223372036854775807:
            var v166 string = ("\n\t" + "a")
            _ = v166
        case 1:
            fmt.Println(v6)
        default:
            if (v83 == v83) {
                var v167 int16 = 42
                v168 := GenS24{F0: (!(!v3)), F1: 247, F2: 4294967295}
                var v169 float64 = 0.0
                _ = v167
                _ = v168
                _ = v169
            } else {
                v170 := GenS25{F0: "hello"}
                _ = v170
            }
        }
    }
    v84 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
