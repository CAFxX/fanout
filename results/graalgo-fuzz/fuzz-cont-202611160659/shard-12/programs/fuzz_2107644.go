package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 uint; F1 int32; F2 uint16 }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 string; F1 uint64; F2 uint32 }
type GenS12 struct { F0 uint8; F1 string; F2 uint64 }
type GenS13 struct { F0 string; F1 uint16; F2 int32 }
type GenS14 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS15 struct { F0 int8; F1 uint64; F2 int32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int32; F1 bool; F2 uint64 }
type GenS19 struct { F0 bool; F1 int64; F2 int8 }
type GenS20 struct { F0 string; F1 int; F2 uint8 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 uint16 }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        switch "\n\t" {
        case "hello":
            var v5 uint = 0
            _ = v5
        }
        v6 := GenS1{F0: 65535}
        var v7 uint32 = 2147483647
        _ = fi4
        _ = v6
        _ = v7
    }
    var v8 uint64 = 2
    switch 32767 {
    case 4294967295:
        v8 -= v8
    case 127:
        for fi9 := 0; fi9 < 14; fi9++ {
            if ((-(fi9)) > 127) {
                v8 -= 18446744073709551615
                var v10 float32 = float32(((1e10 * 1.5) - 1.5))
                v11 := []int{fi9, 127, (fi9 << ((706) & 63)), (fi9 & (fi9 * (-(fi9))))}
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                v8 -= v8
                if (("" > "a") && true) {
                    switch "" {
                    case "hello":
                        v13 := GenS2{F0: 255, F1: -1}
                        _ = v13
                    case "\n\t":
                        v14 := GenS1{F0: 844}
                        _ = v14
                    }
                    v15 := GenS3{F0: 1, F1: 65535}
                    if (!(3 != -9223372036854775808)) {
                        if false {
                            v8 += v8
                            var v16 uint32 = 763
                            _ = v16
                        } else {
                            v17 := []int{fi9, (fi9 + 9223372036854775807)}
                            v8 = v8
                            v18 := GenS1{F0: 525}
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = v15
                } else {
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        var v19 int8 = 127
                        _ = v19
                    case "hello":
                        fmt.Println(v8, fi12)
                    case "":
                        var v20 uint = 1
                        _ = v20
                    }
                }
                _ = fi12
            }
            v21 := []int{(fi9 % (((fi9 / ((18) | 1))) | 1)), (fi9 - fi9)}
            _ = fi9
            _ = v21
        }
    }
    v22 := []int{-9223372036854775808, 0, -9223372036854775808, 1}
    _ = v22
    _ = v8
    return 645
}

func main() {
    var v1 int64 = 0
    var v2 uint32 = 0
    var v3 uint = 667
    fmt.Println(v3, v2, v3)
    fmt.Println(v2, v3)
    switch ("a" + ("foo bar" + "a")) {
    case "a":
        if ("\n\t" == "a") {
            v23 := GenS4{F0: (^(v3)), F1: -2147483648}
            _ = v23
        } else {
            v3 -= (^(v3))
            v24 := []int{0, int(v1), -9223372036854775808}
            _ = v24
        }
    case "foo bar":
        fmt.Println(v2, v1, v3)
    case "hello":
        if false {
            if (false || (((v2 - uint32(v3)) < v2) || (!(0 < 128)))) {
                var v25 string = ("\n\t" + ("hello" + (((((("a" + "hello") + ("a" + "hello")) + "\n\t") + ("hello" + "a")) + "\n\t") + ("\n\t" + ""))))
                v25 = "foo bar"
                _ = v25
            }
        }
    default:
        if true {
            v26 := []int{127, 1}
            var v27 uint32 = v2
            _ = v26
            _ = v27
        }
    }
    if ((!(!(!true))) && (("hello" + ("a" + "hello")) < "\n\t")) {
        v28 := []int{4294967295, -9223372036854775808}
        fmt.Println(v1, v3, v1)
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 16; fi29++ {
            v2 += (^(v2))
            v30 := GenS5{F0: int(v3), F1: (fi29 >> ((65535) & 63))}
            var v31 float32 = 3.14159
            _ = fi29
            _ = v30
            _ = v31
        }
        switch "hello" {
        case "\n\t":
            var v32 int = 1
            _ = v32
        case "a":
            for fi33 := 0; fi33 < 4; fi33++ {
                switch fi33 {
                case -746:
                    if true {
                        for fi34 := 0; fi34 < 7; fi34++ {
                            v3 = 71
                            _ = fi34
                        }
                        v35 := GenS4{F0: (v3 & 65535), F1: (^(fi33))}
                        _ = v35
                    } else {
                        v1 -= (-(v1))
                    }
                }
                v1 = v1
                fmt.Println(v3, v1)
                _ = fi33
            }
        default:
            if (-49 <= 127) {
                if (0.1 < ((float32(float64((float32(3.14159) * 1.5))) / ((0.1 + 1e10) * 1e10)) / 1.0)) {
                    if (156 == 255) {
                        switch "hello" {
                        case "hello":
                            switch "\n\t" {
                            case "":
                                v36 := []int{int(v2), int(v2), -1}
                                _ = v36
                            case "hello":
                                v37 := GenS6{F0: uint64(v3), F1: 4294967295}
                                _ = v37
                            case "a":
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    v3 = 65535
                                    _ = fi38
                                }
                            }
                        case "":
                            v39 := GenS6{F0: 2, F1: 18446744073709551615}
                            _ = v39
                        case "\n\t":
                            v3 = (^(v3))
                        default:
                            if (true && (!(42 > int8(v2)))) {
                                var v40 float64 = 1e10
                                fmt.Println(v2, v3, v1)
                                _ = v40
                            } else {
                                switch 0 {
                                case 522:
                                    v41 := []int{2147483647, -2147483648, -9223372036854775808, 2147483647}
                                    _ = v41
                                case 0:
                                    fmt.Println(v1)
                                case -920:
                                    var v42 int = 0
                                    _ = v42
                                }
                                var v43 int32 = -497
                                for fi44 := 0; fi44 < 16; fi44++ {
                                    switch (("\n\t" + "hello") + (("hello" + "hello") + "\n\t")) {
                                    case "\n\t":
                                        v45 := []int{(-(fi44)), 42, 32767, fi44}
                                        _ = v45
                                    case "":
                                        if false {
                                            var v46 int64 = (v1 * v1)
                                            _ = v46
                                        }
                                    default:
                                        v43 = v43
                                    }
                                    if (v1 <= (^(v1))) {
                                        switch fi44 {
                                        case -9223372036854775808:
                                            var v47 string = "\n\t"
                                            _ = v47
                                        default:
                                            fmt.Println(fi44, v3)
                                        }
                                        var v48 int16 = 32767
                                        if (((826 <= 879) || (1e10 != 0.1)) && false) {
                                            v49 := GenS6{F0: 2, F1: 579}
                                            fmt.Println(fi44, v1)
                                            _ = v49
                                        } else {
                                            if true {
                                                var v50 uint16 = 32767
                                                var v51 uint = 1
                                                _ = v50
                                                _ = v51
                                            } else {
                                                var v52 int64 = (v1 >> ((4294967295) & 63))
                                                _ = v52
                                            }
                                            var v53 string = "a"
                                            v54 := []int{int(v1)}
                                            _ = v53
                                            _ = v54
                                        }
                                        _ = v48
                                    }
                                    _ = fi44
                                }
                                _ = v43
                            }
                        }
                        if (65535 > 454) {
                            var v55 uint = (v3 + v3)
                            v3 -= (v3 << (((^(v3))) & 63))
                            _ = v55
                        } else {
                            var v56 uint = 65535
                            _ = v56
                        }
                    }
                } else {
                    fmt.Println(v1)
                    var v57 uint = v3
                    _ = v57
                }
                var v58 int8 = 65
                for fi59 := 0; fi59 < 9; fi59++ {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        fmt.Println(v1, v2)
                    case "":
                        if (!true) {
                            var v60 int16 = 226
                            switch "a" {
                            case "":
                                switch fi59 {
                                case -2147483648:
                                    var v61 float32 = 1e10
                                    _ = v61
                                default:
                                    v62 := []int{(fi59 % (((fi59 % ((int(v1)) | 1))) | 1)), (-(fi59)), fi59}
                                    _ = v62
                                }
                            case "a":
                                v1 = 9223372036854775807
                            case "foo bar":
                                switch (fi59 & (fi59 / ((127) | 1))) {
                                case 4294967295:
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v63 uint8 = 42
                                        _ = v63
                                    case "":
                                        v3 -= (v3 % ((v3) | 1))
                                    case "hello":
                                        switch fi59 {
                                        case 125:
                                            if (2 == 163) {
                                                v64 := []int{fi59, fi59, fi59, (-(fi59))}
                                                v64 = append(v64, 4294967295)
                                                var v65 uint8 = 127
                                                _ = v64
                                                _ = v65
                                            } else {
                                                switch fi59 {
                                                case 127:
                                                    v60 += 1000
                                                case -126:
                                                    v66 := GenS6{F0: 4294967295, F1: 42}
                                                    _ = v66
                                                case 128:
                                                    v1 = 851
                                                default:
                                                    var v67 uint64 = 1
                                                    _ = v67
                                                }
                                                fmt.Println(fi59)
                                                var v68 string = "a"
                                                _ = v68
                                            }
                                        case -609:
                                            fmt.Println(v1)
                                        default:
                                            for fi69 := 0; fi69 < 12; fi69++ {
                                                var v70 int32 = -162
                                                v71 := GenS3{F0: 4294967295, F1: 163}
                                                v72 := GenS3{F0: 344, F1: 2}
                                                switch ("foo bar" + ("hello" + "a")) {
                                                case "\n\t":
                                                    var v73 uint16 = genFn0()
                                                    _ = v73
                                                case "a":
                                                    var v74 int32 = (v70 / ((v70) | 1))
                                                    _ = v74
                                                case "":
                                                    if false {
                                                        v75 := []int{}
                                                        _ = v75
                                                    } else {
                                                        v76 := []int{fi59}
                                                        _ = v76
                                                    }
                                                default:
                                                    switch (("a" + ("" + "a")) + "foo bar") {
                                                    case "\n\t":
                                                        var v77 uint16 = 42
                                                        _ = v77
                                                    case "a":
                                                        var v78 float32 = 1e10
                                                        _ = v78
                                                    default:
                                                        v79 := []int{fi69}
                                                        _ = v79
                                                    }
                                                }
                                                _ = fi69
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                        }
                                    }
                                case -601:
                                    if (true || true) {
                                        var v80 float64 = 1e10
                                        if (true && (v1 <= v1)) {
                                            v2 = v2
                                            switch (fi59 / ((127) | 1)) {
                                            case 517:
                                                v81 := GenS4{F0: v3, F1: fi59}
                                                _ = v81
                                            case -9223372036854775808:
                                                v82 := GenS6{F0: 42, F1: 0}
                                                _ = v82
                                            default:
                                                v83 := GenS5{F0: (fi59 * 1), F1: 4294967295}
                                                _ = v83
                                            }
                                            v3 = (v3 / ((18446744073709551615) | 1))
                                        }
                                        _ = v80
                                    } else {
                                        switch fi59 {
                                        case 9223372036854775807:
                                            switch (^(fi59)) {
                                            case 954:
                                                var v84 uint16 = 161
                                                _ = v84
                                            case 9223372036854775807:
                                                v85 := []int{}
                                                _ = v85
                                            }
                                        case 312:
                                            if true {
                                                v86 := GenS7{F0: v3, F1: 32767}
                                                for fi87 := 0; fi87 < 4; fi87++ {
                                                    switch fi87 {
                                                    case 0:
                                                        v2 += v2
                                                    default:
                                                        switch (fi87 & (fi59 / ((127) | 1))) {
                                                        case 32768:
                                                            var v88 int = fi87
                                                            _ = v88
                                                        case 32767:
                                                            if true {
                                                                v58 = (v58 * v58)
                                                            } else {
                                                                var v89 float64 = 1.5
                                                                var v90 float64 = v89
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        default:
                                                            v91 := GenS5{F0: fi59, F1: fi87}
                                                            _ = v91
                                                        }
                                                    }
                                                    _ = fi87
                                                }
                                                _ = v86
                                            }
                                        case -777:
                                            switch "hello" {
                                            case "foo bar":
                                                v92 := GenS8{F0: v3, F1: 42, F2: genFn0()}
                                                _ = v92
                                            case "":
                                                switch fi59 {
                                                case -811:
                                                    v3 += v3
                                                case 4294967295:
                                                    for fi93 := 0; fi93 < 8; fi93++ {
                                                        if ((255 >= 32767) && ((-(v60)) != v60)) {
                                                            v94 := GenS4{F0: uint(v2), F1: (fi93 << ((v3) & 63))}
                                                            v60 += 51
                                                            v95 := GenS9{F0: 42, F1: 255}
                                                            v96 := []int{(fi93 % ((127) | 1)), 4294967295, fi59}
                                                            if len(v96) >= 2 {
                                                                v96 = v96[1:len(v96)]
                                                            }
                                                            _ = v94
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            if false {
                                                                if ((472 >= uint64(v2)) || false) {
                                                                    v97 := GenS9{F0: 31, F1: 65535}
                                                                    var v98 string = ("hello" + ("hello" + "\n\t"))
                                                                    var v99 string = v98
                                                                    v100 := GenS8{F0: v3, F1: -1, F2: 0}
                                                                    _ = v100
                                                                    _ = v97
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                            }
                                                            var v101 string = ("foo bar" + "")
                                                            fmt.Println(v1, v3, v2)
                                                            _ = v101
                                                        }
                                                        if (2 != 18446744073709551615) {
                                                            v102 := GenS10{F0: (v3 / ((4294967295) | 1))}
                                                            _ = v102
                                                        } else {
                                                            var v103 int16 = v60
                                                            _ = v103
                                                        }
                                                        v1 -= v1
                                                        _ = fi93
                                                    }
                                                default:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v2 -= v2
                                                    default:
                                                        var v104 uint8 = 2
                                                        _ = v104
                                                    }
                                                }
                                            }
                                        default:
                                            var v105 float64 = 1e10
                                            _ = v105
                                        }
                                    }
                                }
                            }
                            _ = v60
                        } else {
                            var v106 int32 = 0
                            _ = v106
                        }
                    case "a":
                        var v107 uint = v3
                        _ = v107
                    default:
                        for fi108 := 0; fi108 < 7; fi108++ {
                            switch "foo bar" {
                            case "":
                                var v109 int8 = (v58 + int8(fi108))
                                _ = v109
                            case "\n\t":
                                var v110 uint64 = 42
                                _ = v110
                            default:
                                v1 -= (v1 % (((v1 << ((42) & 63))) | 1))
                            }
                            switch fi108 {
                            case -1:
                                var v111 int32 = -231
                                _ = v111
                            case 127:
                                var v112 string = "\n\t"
                                _ = v112
                            default:
                                var v113 uint8 = 0
                                _ = v113
                            }
                            _ = fi108
                        }
                    }
                    _ = fi59
                }
                _ = v58
            } else {
                for fi114 := 0; fi114 < 19; fi114++ {
                    switch -645 {
                    case -709:
                        v115 := GenS4{F0: (^(v3)), F1: fi114}
                        _ = v115
                    case -844:
                        if false {
                            var v116 float32 = 1.5
                            switch "hello" {
                            case "hello":
                                v3 += (^(v3))
                            case "a":
                                switch "\n\t" {
                                case "foo bar":
                                    var v117 uint16 = 804
                                    _ = v117
                                }
                            case "\n\t":
                                if (float32((3.14159 / 3.14159)) <= ((float32(((3.14159 / (1e10 + 3.14159)) * 3.14159)) - (v116 - ((v116 + 3.14159) - ((v116 / 1.5) / (3.14159 / v116))))) * v116)) {
                                    if true {
                                        var v118 int8 = 42
                                        for fi119 := 0; fi119 < 17; fi119++ {
                                            v120 := GenS11{F0: ("a" + "foo bar"), F1: 1, F2: uint32(v3)}
                                            v116 = 1.5
                                            if ((v116 + v116) < v116) {
                                                v1 += v1
                                                var v121 uint64 = 741
                                                _ = v121
                                            } else {
                                                v122 := GenS3{F0: 0, F1: 255}
                                                var v123 int16 = -201
                                                v124 := []int{4294967295}
                                                _ = v122
                                                _ = v123
                                                _ = v124
                                            }
                                            _ = fi119
                                            _ = v120
                                        }
                                        var v125 uint16 = genFn0()
                                        _ = v118
                                        _ = v125
                                    } else {
                                        var v126 uint = v3
                                        _ = v126
                                    }
                                    v127 := GenS12{F0: 105, F1: ("" + "a"), F2: 42}
                                    _ = v127
                                }
                            }
                            for fi128 := 0; fi128 < 15; fi128++ {
                                v129 := []int{}
                                if (true || false) {
                                    v116 = ((float32(-2.25) + v116) * float32(1e10))
                                    var v130 uint64 = 1
                                    v3 -= (v3 * v3)
                                    _ = v130
                                } else {
                                    switch fi114 {
                                    case 410:
                                        if false {
                                            fmt.Println(v1)
                                            fmt.Println(v2, fi114, v1)
                                            var v131 float32 = v116
                                            _ = v131
                                        }
                                    case 1:
                                        v132 := GenS11{F0: ("" + "hello"), F1: 0, F2: (v2 | 187)}
                                        _ = v132
                                    }
                                    var v133 int = fi114
                                    _ = v133
                                }
                                v134 := GenS9{F0: -2147483648, F1: v3}
                                _ = fi128
                                _ = v129
                                _ = v134
                            }
                            _ = v116
                        }
                    default:
                        for fi135 := 0; fi135 < 3; fi135++ {
                            switch -2147483648 {
                            case 0:
                                v136 := GenS8{F0: v3, F1: 5, F2: 32767}
                                _ = v136
                            case -9223372036854775808:
                                switch "a" {
                                case "foo bar":
                                    v3 = (^(v3))
                                    switch "foo bar" {
                                    case "\n\t":
                                        v137 := GenS3{F0: 534, F1: 326}
                                        _ = v137
                                    case "a":
                                        v1 += v1
                                    }
                                    v138 := []int{fi135}
                                    _ = v138
                                default:
                                    v3 -= v3
                                    var v139 int64 = v1
                                    v2 = v2
                                    if false {
                                        if false {
                                            v140 := []int{(fi135 * fi114), (fi114 / (((fi114 / (((^(fi114))) | 1))) | 1)), (fi135 * (fi135 >> ((65535) & 63)))}
                                            _ = v140
                                        } else {
                                            if (84 > v3) {
                                                v2 = (v2 - 394)
                                            }
                                            var v141 int16 = 0
                                            switch fi114 {
                                            case 60:
                                                v139 = (v1 | v139)
                                            case 128:
                                                v141 += v141
                                            case 32767:
                                                v1 = v139
                                            default:
                                                fmt.Println(v1, v3, v139)
                                            }
                                            _ = v141
                                        }
                                    } else {
                                        v142 := GenS5{F0: (fi135 >> ((1) & 63)), F1: fi114}
                                        var v143 string = "a"
                                        _ = v142
                                        _ = v143
                                    }
                                    var v144 bool = (!(4294967295 < (^(v3))))
                                    _ = v139
                                    _ = v144
                                }
                            default:
                                var v145 int64 = 42
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v2, v145)
                                case "\n\t":
                                    v1 = (v1 / ((0) | 1))
                                }
                                _ = v145
                            }
                            var v146 int64 = v1
                            var v147 bool = true
                            _ = fi135
                            _ = v146
                            _ = v147
                        }
                    }
                    v2 -= (^(v2))
                    v148 := GenS13{F0: "hello", F1: 108, F2: 2147483647}
                    _ = fi114
                    _ = v148
                }
                var v149 float64 = (0.1 * 0.0)
                switch "a" {
                case "":
                    if true {
                        switch "hello" {
                        case "\n\t":
                            v150 := GenS14{F0: 42, F1: 80, F2: v1}
                            _ = v150
                        default:
                            if true {
                                v149 -= v149
                                v151 := []int{-15, 127}
                                _ = v151
                            } else {
                                v149 += (v149 - 3.14159)
                            }
                        }
                        var v152 bool = (false || false)
                        v152 = (true || v152)
                        _ = v152
                    }
                case "a":
                    var v153 uint8 = 128
                    _ = v153
                case "hello":
                    v154 := GenS2{F0: 0, F1: 1}
                    _ = v154
                }
                _ = v149
            }
        }
    }
    v155 := []int{128, -2147483648, -9223372036854775808, -822}
    v3 += 13
    v2 -= (^(v2))
    var v156 uint = v3
    for fi157 := 0; fi157 < 18; fi157++ {
        if len(v155) > 0 {
            v155[0] = (fi157 * (fi157 + 32768))
        }
        switch 812 {
        case 32767:
            for fi158 := 0; fi158 < 6; fi158++ {
                var v159 uint64 = 4294967295
                _ = len(v155)
                var v160 float32 = 0.1
                _ = fi158
                _ = v159
                _ = v160
            }
        default:
            for fi161 := 0; fi161 < 16; fi161++ {
                v156 += v3
                v1 = 0
                v1 = (v1 | (v1 | (v1 % (((v1 % ((-1) | 1))) | 1))))
                _ = fi161
            }
        }
        _ = fi157
    }
    switch "" {
    case "a":
        for fi162 := 0; fi162 < 13; fi162++ {
            v3 += v156
            var v163 string = ("foo bar" + "\n\t")
            var v164 int8 = -128
            _ = fi162
            _ = v163
            _ = v164
        }
    case "foo bar":
        if len(v155) > 0 {
            v155[0] = 32767
        }
    case "":
        var v165 float64 = -2.25
        _ = v165
    }
    if (!(3.14159 != float64(3.14159))) {
        v1 += v1
    } else {
        if len(v155) > 0 {
            v155[0] = 0
        }
        for fi166 := 0; fi166 < 19; fi166++ {
            v1 += v1
            v3 += (v3 % (((v156 << ((42) & 63))) | 1))
            _ = fi166
        }
        v167 := GenS7{F0: (v156 % ((18446744073709551615) | 1)), F1: -920}
        _ = v167
    }
    if (-215 > 42) {
        if (127 != -63) {
            v168 := GenS10{F0: (v156 + v156)}
            _ = v168
        }
        for fi169 := 0; fi169 < 1; fi169++ {
            v1 = (v1 | (v1 ^ v1))
            if len(v155) > 0 {
                v155[0] = (fi169 << ((v156) & 63))
            }
            switch fi169 {
            case 128:
                v156 -= (v3 >> ((v156) & 63))
            case -432:
                fmt.Println(fi169, fi169)
            default:
                fmt.Println(fi169, v156, v3)
            }
            _ = fi169
        }
    } else {
        switch 127 {
        case -1:
            v170 := GenS15{F0: 100, F1: 434, F2: -1}
            _ = v170
        default:
            for fi171 := 0; fi171 < 8; fi171++ {
                if true {
                    if true {
                        v1 = (v1 << ((4294967295) & 63))
                        v172 := GenS16{F0: int32(v2)}
                        _ = v172
                    } else {
                        v1 += v1
                    }
                    for fi173 := 0; fi173 < 15; fi173++ {
                        switch fi171 {
                        case 0:
                            v174 := GenS17{F0: 4294967295}
                            _ = v174
                        case -1:
                            var v175 int8 = 104
                            _ = v175
                        default:
                            fmt.Println(v156, v156, fi171)
                        }
                        var v176 int16 = 42
                        _ = fi173
                        _ = v176
                    }
                    if false {
                        v3 = (v156 / ((v156) | 1))
                        if len(v155) >= 2 {
                            v155 = v155[1:len(v155)]
                        }
                        v2 = (v2 ^ 437)
                    }
                }
                var v177 float64 = 0.0
                _ = len(v155)
                _ = fi171
                _ = v177
            }
        }
        v3 = (^(v156))
    }
    if ((!false) || (!true)) {
        for fi178 := 0; fi178 < 6; fi178++ {
            for fi179 := 0; fi179 < 10; fi179++ {
                v156 = uint(v2)
                switch ("\n\t" + ("foo bar" + ("foo bar" + ""))) {
                case "foo bar":
                    if true {
                        if len(v155) > 0 {
                            v155[0] = (fi178 % (((fi178 / ((fi179) | 1))) | 1))
                        }
                        v180 := GenS16{F0: -2147483648}
                        _ = v180
                    } else {
                        fmt.Println(fi179, v156, fi179)
                        if ((v2 << ((v3) & 63)) == 2147483648) {
                            _ = len(v155)
                            v156 -= (v3 - (v3 & v156))
                        }
                    }
                case "":
                    var v181 int8 = 100
                    _ = v181
                }
                _ = fi179
            }
            v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
            _ = fi178
        }
    }
    if len(v155) > 0 {
        v155[0] = 2147483647
    }
    for fi182 := 0; fi182 < 12; fi182++ {
        for fi183 := 0; fi183 < 20; fi183++ {
            _ = len(v155)
            _ = fi183
        }
        for fi184 := 0; fi184 < 9; fi184++ {
            switch "foo bar" {
            case "hello":
                switch fi182 {
                case 0:
                    v156 = 427
                case 2147483647:
                    if false {
                        if (1e10 == 3.14159) {
                            v185 := GenS4{F0: (v3 << ((992) & 63)), F1: 0}
                            _ = v185
                        } else {
                            v2 += 42
                            switch ((("" + "") + (("hello" + "foo bar") + "a")) + "hello") {
                            case "foo bar":
                                if false {
                                    var v186 int = fi182
                                    if (691 < v1) {
                                        if true {
                                            v186 += fi182
                                            switch (fi182 ^ -2147483648) {
                                            case -230:
                                                var v187 float32 = (1.5 + (-2.25 + float32(((((1e10 / -2.25) - 3.14159) + (-2.25 - -2.25)) / ((-2.25 * (0.1 * (1.5 / ((1e10 / 1e10) * -2.25)))) + 0.1)))))
                                                switch "a" {
                                                case "foo bar":
                                                    v2 -= (^(v2))
                                                case "":
                                                    if (!("foo bar" != "\n\t")) {
                                                        if len(v155) >= 2 {
                                                            v155 = v155[1:len(v155)]
                                                        }
                                                    }
                                                }
                                                _ = v187
                                            case 42:
                                                var v188 uint16 = 255
                                                _ = v188
                                            }
                                            switch (^(fi184)) {
                                            case -9223372036854775808:
                                                if true {
                                                    switch (fi184 << (((v156 ^ (v3 * v3))) & 63)) {
                                                    case -2147483648:
                                                        if len(v155) >= 2 {
                                                            v155 = v155[1:len(v155)]
                                                        }
                                                    default:
                                                        if ((((v156 << (((v156 | v3)) & 63)) >= v156) && true) && (!(32767 != 255))) {
                                                            var v189 bool = false
                                                            _ = v189
                                                        }
                                                    }
                                                    if true {
                                                        v190 := GenS9{F0: 0, F1: v3}
                                                        _ = v190
                                                    }
                                                    v186 -= (fi184 + -1)
                                                }
                                            case 0:
                                                var v191 int32 = 842
                                                var v192 int16 = -1
                                                _ = v191
                                                _ = v192
                                            }
                                        }
                                        var v193 uint32 = v2
                                        fmt.Println(v156, v193, v1)
                                        v193 -= v2
                                        if len(v155) >= 2 {
                                            v155 = v155[1:len(v155)]
                                        }
                                        v1 -= v1
                                        v3 = 1
                                        var v194 bool = false
                                        switch (fi182 + 4294967295) {
                                        case 32768:
                                            fmt.Println(v2, fi184)
                                        default:
                                            v186 = (^(v186))
                                        }
                                        v3 = (v156 % ((v156) | 1))
                                        _ = v193
                                        _ = v194
                                    }
                                    v195 := GenS10{F0: v3}
                                    var v196 int64 = v1
                                    _ = v186
                                    _ = v195
                                    _ = v196
                                }
                            case "a":
                                switch "foo bar" {
                                case "\n\t":
                                    switch "a" {
                                    case "hello":
                                        var v197 bool = false
                                        if len(v155) > 0 {
                                            v155[0] = (fi182 * -707)
                                        }
                                        if false {
                                            v197 = (fi184 != 248)
                                            if v197 {
                                                v2 = (^(v2))
                                                v198 := GenS3{F0: 18446744073709551615, F1: 886}
                                                _ = v198
                                            }
                                        }
                                        var v199 int64 = (v1 / (((v1 | 9223372036854775807)) | 1))
                                        var v200 float32 = (1e10 * ((0.0 + -2.25) * 0.1))
                                        _ = v197
                                        _ = v199
                                        _ = v200
                                    case "":
                                        if (!(("foo bar" + "") < ((("foo bar" + (("foo bar" + "foo bar") + ("" + "foo bar"))) + "a") + ("a" + ("hello" + ("hello" + "a")))))) {
                                            v156 += (^(v3))
                                            v155 = append(v155, fi184)
                                            switch "hello" {
                                            case "":
                                                v3 = 255
                                            case "\n\t":
                                                var v201 int16 = 1
                                                _ = v201
                                            default:
                                                var v202 uint64 = 0
                                                _ = v202
                                            }
                                        }
                                    }
                                case "":
                                    v203 := GenS7{F0: v3, F1: fi182}
                                    _ = v203
                                default:
                                    if len(v155) >= 2 {
                                        v155 = v155[1:len(v155)]
                                    }
                                }
                            }
                            var v204 int8 = 42
                            v1 += v1
                            fmt.Println(v3)
                            _ = v204
                        }
                        switch ("\n\t" + "hello") {
                        case "a":
                            v1 -= 9223372036854775807
                        case "foo bar":
                            if (genFn0() <= 814) {
                                v205 := GenS11{F0: "foo bar", F1: 0, F2: 710}
                                v3 -= 42
                                _ = v205
                            }
                        default:
                            switch ("\n\t" + "\n\t") {
                            case "foo bar":
                                var v206 uint32 = (v2 << ((v156) & 63))
                                _ = v206
                            case "hello":
                                v207 := GenS5{F0: fi182, F1: fi182}
                                _ = v207
                            case "a":
                                v208 := GenS11{F0: ("foo bar" + "hello"), F1: 18446744073709551615, F2: v2}
                                _ = v208
                            }
                            if (false || (!(v1 > -604))) {
                                var v209 int = (-(fi184))
                                v210 := GenS12{F0: 255, F1: ("foo bar" + "\n\t"), F2: 18446744073709551615}
                                var v211 int16 = -32768
                                _ = v209
                                _ = v210
                                _ = v211
                            } else {
                                var v212 uint = v3
                                if false {
                                    var v213 uint = v156
                                    var v214 bool = true
                                    v215 := GenS2{F0: 6, F1: 1}
                                    _ = v213
                                    _ = v214
                                    _ = v215
                                }
                                v3 = (v212 / (((v156 % (((v212 & v156)) | 1))) | 1))
                                fmt.Println(v3, v212)
                                _ = v212
                            }
                            v1 = (v1 / (((v1 * 42)) | 1))
                        }
                        var v216 int8 = int8(v2)
                        _ = v216
                    }
                case 9223372036854775807:
                    v2 += 533
                default:
                    if (446 > -251) {
                        fmt.Println(fi182, fi182)
                        if (687 != 0) {
                            if len(v155) >= 2 {
                                v155 = v155[1:len(v155)]
                            }
                            v3 = (^(v156))
                            var v217 int32 = 635
                            _ = v217
                        } else {
                            v156 += (v156 >> ((392) & 63))
                        }
                        if len(v155) > 0 {
                            v155[0] = fi182
                        }
                        var v218 int64 = (v1 & -766)
                        _ = v218
                    }
                }
            case "\n\t":
                if false {
                    fmt.Println(fi184, fi182, v156)
                    v3 += 80
                    if true {
                        switch 670 {
                        case 2147483647:
                            v219 := GenS18{F0: 79, F1: ((!((1 >= 65535) && true)) && (((1.5 + 0.0) * (float32(-2.25) * float32((0.0 - 1e10)))) != -2.25)), F2: 4294967295}
                            _ = v219
                        case 176:
                            v2 = (v2 / (((v2 ^ v2)) | 1))
                        case 343:
                            fmt.Println(v156, v156)
                        }
                    }
                } else {
                    v220 := GenS10{F0: (v3 ^ 18446744073709551615)}
                    if (true || true) {
                        v221 := GenS19{F0: (v1 < v1), F1: (v1 * 9223372036854775807), F2: 100}
                        fmt.Println(v1, v3, v3)
                        if (true && true) {
                            var v222 uint16 = 1
                            if (42 < 65) {
                                var v223 int64 = 42
                                v3 = uint(v222)
                                _ = v223
                            }
                            _ = v222
                        } else {
                            _ = len(v155)
                        }
                        _ = v221
                    }
                    if true {
                        switch 4294967295 {
                        case 9223372036854775807:
                            if len(v155) > 0 {
                                v155[0] = fi184
                            }
                        }
                        switch (fi182 & fi182) {
                        case 4294967295:
                            switch "\n\t" {
                            case "":
                                if false {
                                    if true {
                                        if true {
                                            v1 -= v1
                                            v224 := GenS19{F0: (float32((-2.25 + 0.0)) > 1.5), F1: v1, F2: -105}
                                            _ = v224
                                        } else {
                                            v155 = append(v155, fi184)
                                        }
                                        var v225 float64 = 1e10
                                        var v226 int32 = 76
                                        fmt.Println(fi182)
                                        if true {
                                            switch ("" + "hello") {
                                            case "hello":
                                                var v227 uint = (^(v156))
                                                _ = v227
                                            case "\n\t":
                                                _ = len(v155)
                                            case "a":
                                                switch (("\n\t" + "hello") + ("foo bar" + "hello")) {
                                                case "a":
                                                    var v228 uint = (^(v156))
                                                    switch ("" + (("hello" + "a") + "")) {
                                                    case "hello":
                                                        var v229 int8 = -128
                                                        _ = v229
                                                    case "foo bar":
                                                        fmt.Println(v156)
                                                    case "a":
                                                        fmt.Println(fi184)
                                                    }
                                                    if (("hello" + "") <= "") {
                                                        switch (("foo bar" + ("hello" + "a")) + "foo bar") {
                                                        case "\n\t":
                                                            v228 = v3
                                                            if true {
                                                                v1 = 0
                                                            }
                                                            v230 := GenS1{F0: 4294967295}
                                                            fmt.Println(v3, v228)
                                                            var v231 uint32 = v2
                                                            _ = v230
                                                            _ = v231
                                                        }
                                                    }
                                                    _ = v228
                                                case "":
                                                    v1 += int64(v3)
                                                case "\n\t":
                                                    if true {
                                                        if len(v155) >= 2 {
                                                            v155 = v155[1:len(v155)]
                                                        }
                                                        v232 := GenS20{F0: "a", F1: fi184, F2: 255}
                                                        _ = v232
                                                    } else {
                                                        v156 += 18446744073709551615
                                                        v155 = append(v155, 42)
                                                    }
                                                    v225 -= v225
                                                    if len(v155) > 0 {
                                                        v155[0] = (fi184 >> ((659) & 63))
                                                    }
                                                default:
                                                    if true {
                                                        if true {
                                                            switch ("hello" + "a") {
                                                            case "hello":
                                                                if false {
                                                                    v2 = v2
                                                                    v1 = 42
                                                                    if false {
                                                                        var v233 uint = v156
                                                                        v155 = append(v155, int(v1))
                                                                        _ = v233
                                                                    } else {
                                                                        v155 = append(v155, (^(fi184)))
                                                                        v234 := GenS13{F0: ("foo bar" + "\n\t"), F1: 0, F2: v226}
                                                                        if false {
                                                                            var v235 int16 = 42
                                                                            v225 += v225
                                                                            var v236 int16 = v235
                                                                            _ = v235
                                                                            _ = v236
                                                                        } else {
                                                                            v3 = v156
                                                                            if false {
                                                                                v237 := GenS8{F0: uint(v2), F1: v226, F2: 255}
                                                                                v156 += (v3 % ((uint(v2)) | 1))
                                                                                _ = v237
                                                                            } else {
                                                                                var v238 uint16 = 42
                                                                                fmt.Println(v226, v156)
                                                                                _ = v238
                                                                            }
                                                                            _ = len(v155)
                                                                        }
                                                                        _ = v234
                                                                    }
                                                                    v1 -= v1
                                                                }
                                                            case "\n\t":
                                                                _ = len(v155)
                                                            case "a":
                                                                if false {
                                                                    switch 960 {
                                                                    case 723:
                                                                        v2 = (^(v2))
                                                                        v225 = v225
                                                                    }
                                                                    v239 := GenS21{F0: 874}
                                                                    _ = v239
                                                                } else {
                                                                    var v240 int64 = v1
                                                                    _ = v240
                                                                }
                                                            }
                                                            fmt.Println(fi184)
                                                        }
                                                    } else {
                                                        switch "foo bar" {
                                                        case "":
                                                            v156 -= (v3 | uint(v2))
                                                            v2 = v2
                                                            if len(v155) >= 2 {
                                                                v155 = v155[1:len(v155)]
                                                            }
                                                        case "foo bar":
                                                            var v241 int32 = (v226 - v226)
                                                            v242 := GenS14{F0: 1, F1: 938, F2: -1}
                                                            _ = v241
                                                            _ = v242
                                                        }
                                                    }
                                                }
                                            default:
                                                var v243 int8 = 44
                                                if (int64(v156) == 42) {
                                                    v2 = 623
                                                    v243 = v243
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        v226 = v226
                                                    case "":
                                                        v3 -= 532
                                                    default:
                                                        var v244 int8 = (v243 >> ((162) & 63))
                                                        _ = v244
                                                    }
                                                }
                                                _ = v243
                                            }
                                            v245 := GenS6{F0: 4294967295, F1: 1}
                                            _ = v245
                                        }
                                        _ = v225
                                        _ = v226
                                    }
                                    if len(v155) > 0 {
                                        v155[0] = fi184
                                    }
                                    fmt.Println(v156, fi184, fi182)
                                }
                            default:
                                v3 -= v156
                            }
                            if ("\n\t" == "a") {
                                var v246 float64 = 1.5
                                _ = v246
                            }
                        }
                    } else {
                        var v247 float32 = 0.0
                        v156 = 18446744073709551615
                        _ = v247
                    }
                    if false {
                        if true {
                            v155 = append(v155, fi182)
                            switch (fi182 * -940) {
                            case 961:
                                _ = len(v155)
                            }
                        } else {
                            var v248 float64 = 1e10
                            _ = v248
                        }
                    } else {
                        if len(v155) > 0 {
                            v155[0] = (fi182 << ((v156) & 63))
                        }
                    }
                    _ = v220
                }
                var v249 float32 = 0.0
                _ = v249
            case "a":
                if true {
                    var v250 int8 = 100
                    v2 -= (v2 >> ((0) & 63))
                    v2 += v2
                    _ = v250
                } else {
                    v2 = (v2 | v2)
                    switch ("\n\t" + ("foo bar" + ("foo bar" + ""))) {
                    case "a":
                        v1 = (v1 % ((736) | 1))
                    case "\n\t":
                        v251 := GenS22{F0: 0}
                        _ = v251
                    case "foo bar":
                        v3 = 326
                    }
                }
            }
            v1 += (v1 * (^(v1)))
            fmt.Println(v1, v2)
            _ = fi184
        }
        _ = fi182
    }
    v3 += 255
    v155 = append(v155, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v155:", len(v155), v155)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v155
    _ = v156
    _ = v2
    _ = v3
}
