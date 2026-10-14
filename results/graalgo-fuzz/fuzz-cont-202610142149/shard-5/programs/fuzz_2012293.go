package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int; F2 uint16 }
type GenS6 struct { F0 int16; F1 uint64; F2 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 uint16; F1 int; F2 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint; F1 uint; F2 uint }
type GenS12 struct { F0 int32; F1 int64; F2 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint; F1 int64 }
type GenS15 struct { F0 int; F1 uint; F2 uint8 }
func genFn0(p0 uint64, p1 uint64, p2 int32) bool {
    var v5 int64 = -9223372036854775808
    if (65535 < 282) {
        if (269 < 433) {
            for fi6 := 0; fi6 < 16; fi6++ {
                p0 = (^(p0))
                v7 := GenS1{F0: fi6, F1: (fi6 << ((0) & 63))}
                if true {
                    for fi8 := 0; fi8 < 11; fi8++ {
                        fmt.Println(fi8)
                        v9 := GenS1{F0: -2147483648, F1: fi6}
                        if false {
                            var v10 float32 = 1.5
                            v11 := []int{853}
                            _ = v10
                            _ = v11
                        }
                        _ = fi8
                        _ = v9
                    }
                    var v12 int8 = 68
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if true {
                            v14 := []int{(fi13 ^ fi13), -248}
                            _ = v14
                        } else {
                            v15 := GenS1{F0: (fi6 ^ fi6), F1: fi6}
                            v16 := []int{-2147483648, (fi6 | fi13), fi13}
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                    }
                    _ = v12
                }
                _ = fi6
                _ = v7
            }
        } else {
            for fi17 := 0; fi17 < 19; fi17++ {
                fmt.Println(v5, p0)
                if ((1e10 + 1.5) > (-2.25 - (0.1 + (3.14159 - float32(1e10))))) {
                    p2 += 42
                    v18 := GenS1{F0: 32768, F1: fi17}
                    _ = v18
                } else {
                    var v19 bool = (true || ((((1e10 - 0.0) <= 3.14159) && true) || false))
                    var v20 int8 = -1
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    var v22 float32 = ((3.14159 / 1.5) - 0.1)
                    _ = fi21
                    _ = v22
                }
                _ = fi17
            }
            var v23 float64 = 0.0
            _ = v23
        }
        v24 := GenS2{F0: 385}
        var v25 uint16 = 65535
        _ = v24
        _ = v25
    } else {
        if (3.14159 >= (1e10 - (1.5 + (-2.25 * (0.1 - 0.0))))) {
            fmt.Println(v5)
            v5 = -1
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (18446744073709551615 != 818)
}
func genFn1() int64 {
    var v26 int8 = -128
    switch 9223372036854775807 {
    case 32768:
        switch "foo bar" {
        case "foo bar":
            for fi27 := 0; fi27 < 5; fi27++ {
                fmt.Println(v26, fi27)
                for fi28 := 0; fi28 < 13; fi28++ {
                    if false {
                        v29 := GenS2{F0: 1}
                        _ = v29
                    } else {
                        v26 += v26
                        if (true && ((65535 > 2) && (!(255 < 32767)))) {
                            fmt.Println(fi27, fi27, fi27)
                            var v30 uint64 = 18446744073709551615
                            _ = v30
                        }
                        switch (("" + "\n\t") + "foo bar") {
                        case "hello":
                            var v31 uint16 = 1
                            _ = v31
                        default:
                            switch "hello" {
                            case "foo bar":
                                v32 := GenS1{F0: 42, F1: (fi28 * fi28)}
                                _ = v32
                            default:
                                switch fi28 {
                                case 4294967295:
                                    v33 := GenS2{F0: 32767}
                                    _ = v33
                                default:
                                    v34 := []int{1, (fi27 % ((fi28) | 1)), fi28}
                                    _ = v34
                                }
                            }
                        }
                    }
                    v26 -= v26
                    if true {
                        var v35 uint32 = 42
                        switch (fi27 << ((1) & 63)) {
                        case -9223372036854775808:
                            var v36 int16 = 0
                            _ = v36
                        case 127:
                            var v37 float64 = ((1.5 * 1.5) / (1e10 - 0.1))
                            _ = v37
                        case 32767:
                            if ((false && false) && (!(false || (0 <= 42)))) {
                                fmt.Println(v26, fi28, v26)
                                v38 := []int{}
                                var v39 int = (fi27 & (fi27 << ((1) & 63)))
                                _ = v38
                                _ = v39
                            } else {
                                v40 := []int{fi27, (^(fi28)), (fi27 / ((9223372036854775807) | 1))}
                                _ = v40
                            }
                        default:
                            if true {
                                fmt.Println(fi28, v35, v35)
                                var v41 string = "foo bar"
                                _ = v41
                            }
                        }
                        switch "hello" {
                        case "a":
                            if false {
                                v42 := GenS1{F0: (fi28 / (((fi28 % (((-(fi27))) | 1))) | 1)), F1: fi27}
                                v26 = (v26 ^ -113)
                                if (!true) {
                                    if false {
                                        var v43 float32 = 1.5
                                        _ = v43
                                    } else {
                                        if true {
                                            switch "foo bar" {
                                            case "":
                                                switch "hello" {
                                                case "\n\t":
                                                    switch (fi27 & (fi27 << ((0) & 63))) {
                                                    case 0:
                                                        v44 := GenS3{F0: (fi28 ^ fi28)}
                                                        _ = v44
                                                    }
                                                    fmt.Println(fi27, fi28)
                                                default:
                                                    var v45 uint = 698
                                                    _ = v45
                                                }
                                            case "a":
                                                v46 := GenS1{F0: 128, F1: (fi27 * (fi27 / ((32768) | 1)))}
                                                _ = v46
                                            case "foo bar":
                                                v47 := GenS2{F0: 42}
                                                _ = v47
                                            }
                                            if false {
                                                if true {
                                                    if false {
                                                        var v48 bool = (255 >= uint16(v35))
                                                        _ = v48
                                                    } else {
                                                        v26 -= (^(v26))
                                                        if (32767 != 2) {
                                                            v35 = 2147483647
                                                            switch (fi28 << ((18446744073709551615) & 63)) {
                                                            case -1:
                                                                var v49 uint32 = (v35 * (v35 & v35))
                                                                _ = v49
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        if false {
                                                            v35 += 42
                                                            var v50 uint16 = 42
                                                            _ = v50
                                                        }
                                                        v51 := GenS4{F0: fi28, F1: -442, F2: (v35 >> ((845) & 63))}
                                                        if true {
                                                            fmt.Println(v35, fi28, fi28)
                                                            switch (fi28 - 32767) {
                                                            case 858:
                                                                var v52 float64 = (-2.25 + ((3.14159 / 1.0) / ((-2.25 / 3.14159) + ((0.1 * (-2.25 + 0.1)) + 0.0))))
                                                                _ = v52
                                                            case -104:
                                                                if true {
                                                                    if (4294967295 != 4294967295) {
                                                                        v35 = v35
                                                                        v26 -= 6
                                                                        v53 := []int{fi28, 2147483647, fi27, fi27}
                                                                        _ = v53
                                                                    } else {
                                                                        v54 := GenS5{F0: 2147483647, F1: fi28, F2: uint16(v35)}
                                                                        _ = v54
                                                                    }
                                                                    switch fi27 {
                                                                    case 1:
                                                                        var v55 int16 = 32767
                                                                        _ = v55
                                                                    case -1:
                                                                        v26 += (v26 | v26)
                                                                    case 32768:
                                                                        var v56 int64 = 0
                                                                        _ = v56
                                                                    default:
                                                                        var v57 uint8 = 206
                                                                        _ = v57
                                                                    }
                                                                    v58 := []int{fi27, 2147483647, 328}
                                                                    _ = v58
                                                                } else {
                                                                    v35 = 65535
                                                                    switch "" {
                                                                    case "":
                                                                        switch 449 {
                                                                        case -1:
                                                                            var v59 float64 = (((-2.25 * -2.25) - 3.14159) * (3.14159 / 0.1))
                                                                            v26 += 0
                                                                            v26 = 100
                                                                            _ = v59
                                                                        case -633:
                                                                            if true {
                                                                                v60 := []int{}
                                                                                var v61 bool = (0 != 1)
                                                                                var v62 int = fi28
                                                                                _ = v60
                                                                                _ = v61
                                                                                _ = v62
                                                                            } else {
                                                                                var v63 uint8 = 0
                                                                                fmt.Println(v26, fi27, fi28)
                                                                                if (true || (false || false)) {
                                                                                    fmt.Println(v26)
                                                                                    var v64 float32 = 1e10
                                                                                    fmt.Println(v35, v64, v26)
                                                                                    _ = v64
                                                                                }
                                                                                _ = v63
                                                                            }
                                                                        default:
                                                                            var v65 float64 = 1.5
                                                                            var v66 uint8 = 128
                                                                            _ = v65
                                                                            _ = v66
                                                                        }
                                                                    default:
                                                                        v26 -= -1
                                                                    }
                                                                    if (9223372036854775807 == 42) {
                                                                        if false {
                                                                            switch 112 {
                                                                            case 0:
                                                                                v35 += 2
                                                                            default:
                                                                                var v67 string = ""
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                        fmt.Println(fi28, v26, fi27)
                                                                    }
                                                                }
                                                            default:
                                                                v26 = 127
                                                            }
                                                            var v68 float64 = (float64(float32(0.0)) + 1e10)
                                                            _ = v68
                                                        }
                                                        _ = v51
                                                    } else {
                                                        v26 -= v26
                                                    }
                                                    var v69 int64 = -192
                                                    _ = v69
                                                }
                                                v26 = v26
                                            } else {
                                                v26 = v26
                                                v70 := GenS2{F0: 32767}
                                                _ = v70
                                            }
                                            fmt.Println(fi27, v26, v26)
                                            v26 = 0
                                            if (!true) {
                                                fmt.Println(v35)
                                            }
                                            var v71 int64 = -166
                                            var v72 uint = 1
                                            _ = v71
                                            _ = v72
                                        }
                                        if true {
                                            v26 -= 127
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(fi28, v35, v26)
                                            default:
                                                var v73 int16 = 32767
                                                _ = v73
                                            }
                                        }
                                    }
                                    v26 = v26
                                    var v74 int8 = (^(v26))
                                    _ = v74
                                }
                                v26 = (v26 ^ 1)
                                v75 := GenS4{F0: fi28, F1: 602, F2: 42}
                                _ = v42
                                _ = v75
                            }
                        default:
                            v26 -= (v26 << ((1) & 63))
                        }
                        v35 -= 42
                        fmt.Println(fi27, v35)
                        v76 := []int{573, fi28, fi27, (fi28 | fi27)}
                        _ = v35
                        _ = v76
                    }
                    _ = fi28
                }
                v26 = v26
                _ = fi27
            }
        case "hello":
            if (!false) {
                for fi77 := 0; fi77 < 10; fi77++ {
                    v78 := []int{}
                    var v79 int = (fi77 - (fi77 % ((324) | 1)))
                    var v80 int32 = -587
                    _ = fi77
                    _ = v78
                    _ = v79
                    _ = v80
                }
                var v81 uint64 = 1
                if false {
                    v81 = 4294967295
                } else {
                    if true {
                        for fi82 := 0; fi82 < 20; fi82++ {
                            v81 = 2
                            _ = fi82
                        }
                        fmt.Println(v26, v81)
                        v81 -= v81
                    } else {
                        v81 += (v81 + v81)
                        var v83 int8 = -128
                        v84 := GenS5{F0: 42, F1: 150, F2: 382}
                        _ = v83
                        _ = v84
                    }
                    v85 := []int{535}
                    if (float32((1e10 + (float64(float32((float64(0.0) / -2.25))) / 1.0))) >= float32(1e10)) {
                        fmt.Println(v26)
                    }
                    _ = v85
                }
                _ = v81
            }
        }
    case 0:
        switch 42 {
        case -323:
            v26 += v26
        case 0:
            switch -2147483648 {
            case -2147483648:
                v26 -= v26
            case 42:
                for fi86 := 0; fi86 < 1; fi86++ {
                    v87 := GenS4{F0: fi86, F1: -450, F2: 4294967295}
                    _ = fi86
                    _ = v87
                }
            case -9223372036854775808:
                fmt.Println(v26)
            }
        default:
            v26 -= v26
        }
    default:
        fmt.Println(v26)
    }
    var v88 float32 = 0.0
    switch 127 {
    case 9223372036854775807:
        v89 := []int{42, 9223372036854775807, 32767}
        _ = v89
    case 535:
        switch -483 {
        case -1:
            if true {
                v90 := GenS4{F0: 9223372036854775807, F1: 756, F2: 2147483648}
                var v91 bool = true
                _ = v90
                _ = v91
            }
        case 2147483647:
            switch 128 {
            case 42:
                if (v26 <= (v26 - (v26 ^ v26))) {
                    for fi92 := 0; fi92 < 19; fi92++ {
                        switch 128 {
                        case 9223372036854775807:
                            if true {
                                for fi93 := 0; fi93 < 2; fi93++ {
                                    fmt.Println(fi92, fi92)
                                    v94 := GenS4{F0: 9223372036854775807, F1: 1000, F2: 2147483648}
                                    _ = fi93
                                    _ = v94
                                }
                            } else {
                                var v95 uint = 18446744073709551615
                                _ = v95
                            }
                        case -1:
                            var v96 uint = 0
                            _ = v96
                        default:
                            var v97 string = "hello"
                            _ = v97
                        }
                        if false {
                            if true {
                                var v98 int32 = 42
                                v99 := []int{(-(fi92)), fi92, (fi92 | fi92)}
                                _ = v98
                                _ = v99
                            }
                            v26 = (v26 / (((-(v26))) | 1))
                            var v100 uint64 = 18446744073709551615
                            _ = v100
                        }
                        _ = fi92
                    }
                    if false {
                        v26 -= -1
                        for fi101 := 0; fi101 < 12; fi101++ {
                            v26 += v26
                            if (918 != 2147483647) {
                                fmt.Println(v26, v26)
                            } else {
                                v26 = v26
                                v102 := GenS4{F0: fi101, F1: -827, F2: 754}
                                v103 := GenS6{F0: 1000, F1: 274, F2: fi101}
                                _ = v102
                                _ = v103
                            }
                            _ = fi101
                        }
                    } else {
                        v88 -= (-2.25 + (((v88 * 1.5) - -2.25) * 0.0))
                    }
                    switch -302 {
                    case 2147483647:
                        fmt.Println(v26)
                    case 127:
                        for fi104 := 0; fi104 < 5; fi104++ {
                            v105 := []int{fi104, int(v26)}
                            v26 = v26
                            _ = fi104
                            _ = v105
                        }
                    case -1:
                        switch -2147483648 {
                        case 128:
                            v106 := []int{32768, -442}
                            _ = v106
                        case 9223372036854775807:
                            v107 := []int{60, 103, -9223372036854775808, -658}
                            _ = v107
                        }
                    default:
                        if true {
                            v108 := GenS7{F0: -2147483648}
                            _ = v108
                        } else {
                            var v109 int = -9223372036854775808
                            var v110 int8 = v26
                            _ = v109
                            _ = v110
                        }
                    }
                } else {
                    v111 := []int{1, -2147483648}
                    if false {
                        for fi112 := 0; fi112 < 5; fi112++ {
                            for fi113 := 0; fi113 < 3; fi113++ {
                                var v114 int64 = 83
                                v26 += v26
                                _ = fi113
                                _ = v114
                            }
                            _ = fi112
                        }
                        var v115 int8 = v26
                        var v116 uint64 = 42
                        _ = v115
                        _ = v116
                    } else {
                        fmt.Println(v26, v26)
                        for fi117 := 0; fi117 < 8; fi117++ {
                            v118 := GenS8{F0: 9223372036854775807, F1: 65535}
                            if true {
                                fmt.Println(v88, v88, v88)
                                var v119 int32 = 1
                                var v120 uint16 = 0
                                _ = v119
                                _ = v120
                            } else {
                                v26 = (v26 & (v26 % ((v26) | 1)))
                                if true {
                                    for fi121 := 0; fi121 < 5; fi121++ {
                                        v26 = 75
                                        _ = fi121
                                    }
                                    for fi122 := 0; fi122 < 19; fi122++ {
                                        v26 = 100
                                        switch "\n\t" {
                                        case "\n\t":
                                            if len(v111) >= 2 {
                                                v111 = v111[1:len(v111)]
                                            }
                                        case "hello":
                                            if (("hello" + ((("hello" + "foo bar") + "\n\t") + "\n\t")) != ("" + ("foo bar" + ("hello" + ("hello" + "hello"))))) {
                                                var v123 int8 = (v26 ^ v26)
                                                _ = v123
                                            }
                                        case "a":
                                            if (false && (-343 > -311)) {
                                                v124 := GenS5{F0: 482, F1: fi117, F2: 640}
                                                _ = v124
                                            } else {
                                                var v125 string = ("a" + "hello")
                                                var v126 uint32 = 65535
                                                _ = v125
                                                _ = v126
                                            }
                                        default:
                                            var v127 float64 = 3.14159
                                            _ = v127
                                        }
                                        v128 := GenS3{F0: 127}
                                        _ = fi122
                                        _ = v128
                                    }
                                    if (42 < 0) {
                                        for fi129 := 0; fi129 < 17; fi129++ {
                                            if ("foo bar" != "") {
                                                v111 = append(v111, 32767)
                                            } else {
                                                fmt.Println(fi129, v26, v26)
                                                if false {
                                                    var v130 uint = 2
                                                    _ = v130
                                                }
                                                var v131 float32 = ((v88 - (v88 / v88)) + 1e10)
                                                v131 = ((v88 / (((v88 / ((float32(-2.25) - (1e10 * (v88 * v131))) * v88)) / float32(-2.25)) + v131)) + 0.0)
                                                v132 := GenS6{F0: 0, F1: 916, F2: (-(fi117))}
                                                _ = len(v111)
                                                _ = v131
                                                _ = v132
                                            }
                                            _ = fi129
                                        }
                                        for fi133 := 0; fi133 < 9; fi133++ {
                                            if (true || (!false)) {
                                                if false {
                                                    var v134 uint8 = 42
                                                    if true {
                                                        var v135 uint16 = 1
                                                        _ = v135
                                                    } else {
                                                        var v136 bool = true
                                                        switch (fi133 << ((42) & 63)) {
                                                        case -2147483648:
                                                            var v137 int8 = v26
                                                            _ = v137
                                                        default:
                                                            v138 := GenS9{F0: 884, F1: (fi133 >> ((881) & 63)), F2: v134}
                                                            _ = v138
                                                        }
                                                        _ = v136
                                                    }
                                                    var v139 uint = uint(v134)
                                                    _ = v134
                                                    _ = v139
                                                } else {
                                                    if (32767 < 65535) {
                                                        v26 = 45
                                                    } else {
                                                        switch "" {
                                                        case "":
                                                            switch ("a" + "foo bar") {
                                                            case "\n\t":
                                                                if len(v111) >= 2 {
                                                                    v111 = v111[1:len(v111)]
                                                                }
                                                            case "":
                                                                _ = len(v111)
                                                            }
                                                        case "foo bar":
                                                            var v140 int8 = (v26 + (v26 >> ((962) & 63)))
                                                            _ = v140
                                                        case "a":
                                                            if true {
                                                                var v141 int32 = 42
                                                                _ = v141
                                                            }
                                                        }
                                                    }
                                                    v88 -= 1.5
                                                }
                                                if (!((-176 == 0) && (!(!(1.5 <= ((v88 - ((3.14159 * (v88 - v88)) * 0.0)) / 1e10)))))) {
                                                    if len(v111) >= 2 {
                                                        v111 = v111[1:len(v111)]
                                                    }
                                                    var v142 int = (fi133 ^ (fi133 * -2147483648))
                                                    v142 = fi133
                                                    _ = v142
                                                } else {
                                                    switch ("\n\t" + ("" + ((("\n\t" + "a") + "") + ((("foo bar" + "a") + "a") + "hello")))) {
                                                    case "hello":
                                                        var v143 uint64 = 1
                                                        _ = v143
                                                    case "foo bar":
                                                        v88 -= 1.5
                                                    default:
                                                        var v144 float64 = 0.1
                                                        _ = v144
                                                    }
                                                    fmt.Println(v88, v26)
                                                }
                                            } else {
                                                var v145 string = "a"
                                                var v146 uint32 = 0
                                                fmt.Println(fi133, v26)
                                                _ = v145
                                                _ = v146
                                            }
                                            v26 += (v26 / ((v26) | 1))
                                            fmt.Println(v88, fi133, v26)
                                            _ = fi133
                                        }
                                        var v147 float64 = -2.25
                                        _ = v147
                                    } else {
                                        v148 := GenS1{F0: fi117, F1: -9223372036854775808}
                                        var v149 uint32 = 2147483648
                                        _ = v148
                                        _ = v149
                                    }
                                } else {
                                    v150 := GenS10{F0: 988}
                                    _ = v150
                                }
                            }
                            switch 32767 {
                            case 1:
                                var v151 uint8 = 0
                                _ = v151
                            case -1:
                                var v152 uint64 = 4294967295
                                _ = v152
                            default:
                                switch fi117 {
                                case 32767:
                                    var v153 int8 = v26
                                    _ = v153
                                default:
                                    v88 = (v88 * -2.25)
                                }
                            }
                            _ = fi117
                            _ = v118
                        }
                        var v154 bool = true
                        _ = v154
                    }
                    fmt.Println(v88, v88)
                    _ = v111
                }
            }
        default:
            if (761 <= 0) {
                v155 := []int{496, 128, 2147483647}
                v26 = v26
                _ = v155
            } else {
                for fi156 := 0; fi156 < 13; fi156++ {
                    switch fi156 {
                    case 786:
                        switch (fi156 / ((32768) | 1)) {
                        case -510:
                            v157 := GenS9{F0: 32767, F1: (fi156 << ((879) & 63)), F2: 0}
                            _ = v157
                        default:
                            var v158 int16 = int16(fi156)
                            _ = v158
                        }
                    }
                    v26 = (v26 / (((v26 / (((v26 ^ v26)) | 1))) | 1))
                    var v159 int8 = 0
                    _ = fi156
                    _ = v159
                }
            }
        }
    case 0:
        fmt.Println(v26, v26)
    default:
        v26 = -12
    }
    _ = v26
    _ = v88
    return -1
}

func main() {
    var v1 bool = (4294967295 > 687)
    var v2 int16 = 42
    var v3 uint = 255
    var v4 uint64 = 42
    var v160 float32 = (-2.25 + -2.25)
    v4 = (v4 * v4)
    var v161 int = 32768
    v162 := GenS11{F0: (v3 / ((976) | 1)), F1: (v3 / ((v3) | 1)), F2: (v3 / (((v3 / ((v3) | 1))) | 1))}
    v163 := []int{9223372036854775807}
    var v164 int8 = -42
    var v165 uint32 = 37
    if ((v165 & 1) < 290) {
        for fi166 := 0; fi166 < 12; fi166++ {
            if (v165 < v165) {
                if (v164 == 127) {
                    fmt.Println(v2, v4)
                } else {
                    var v167 bool = ((("" + ("hello" + "\n\t")) >= ("\n\t" + "\n\t")) && v1)
                    _ = v167
                }
                if true {
                    for fi168 := 0; fi168 < 20; fi168++ {
                        v169 := GenS7{F0: -1}
                        _ = fi168
                        _ = v169
                    }
                    var v170 float64 = -2.25
                    _ = v170
                } else {
                    for fi171 := 0; fi171 < 6; fi171++ {
                        var v172 int32 = 1
                        fmt.Println(fi166, v172)
                        _ = fi171
                        _ = v172
                    }
                    v173 := GenS6{F0: (v2 - v2), F1: v4, F2: -286}
                    _ = v173
                }
            }
            _ = len(v163)
            _ = fi166
        }
    }
    fmt.Println(v165, v2)
    var v174 int32 = int32(v3)
    if (!v1) {
        fmt.Println(v1, v160)
        switch v161 {
        case 42:
            var v175 uint32 = v165
            _ = v175
        case -37:
            if len(v163) > 0 {
                v163[0] = (v161 & int(v3))
            }
        case 128:
            if (true && false) {
                if ((-(v164)) != v164) {
                    for fi176 := 0; fi176 < 1; fi176++ {
                        var v177 uint8 = 42
                        if true {
                            v178 := GenS12{F0: (^(v174)), F1: -9223372036854775808, F2: "hello"}
                            v179 := GenS9{F0: uint16(v165), F1: v161, F2: v177}
                            _ = v178
                            _ = v179
                        } else {
                            if len(v163) >= 2 {
                                v163 = v163[1:len(v163)]
                            }
                            v163 = append(v163, (fi176 / ((-9223372036854775808) | 1)))
                        }
                        if false {
                            if len(v163) >= 2 {
                                v163 = v163[1:len(v163)]
                            }
                            var v180 int16 = (v2 - (v2 * -688))
                            if len(v163) > 0 {
                                v163[0] = v161
                            }
                            _ = v180
                        } else {
                            v177 += v177
                        }
                        _ = fi176
                        _ = v177
                    }
                    switch v161 {
                    case 42:
                        for fi181 := 0; fi181 < 1; fi181++ {
                            fmt.Println(v165)
                            _ = fi181
                        }
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    v182 := GenS13{F0: 136}
                    _ = v182
                default:
                    fmt.Println(v165, v3, v1)
                }
            } else {
                v174 -= 0
                if len(v163) >= 2 {
                    v163 = v163[1:len(v163)]
                }
                fmt.Println(v164)
            }
        }
    } else {
        for fi183 := 0; fi183 < 9; fi183++ {
            var v184 int16 = 204
            v160 -= v160
            _ = fi183
            _ = v184
        }
        if len(v163) > 0 {
            v163[0] = 0
        }
    }
    if v1 {
        var v185 int16 = int16(v165)
        _ = v185
    } else {
        v186 := GenS14{F0: (^(v3)), F1: 1}
        var v187 uint16 = 32767
        v188 := GenS15{F0: (-(v161)), F1: (^(v3)), F2: 128}
        _ = v186
        _ = v187
        _ = v188
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v160:", v160)
    fmt.Println("v161:", v161)
    fmt.Println("slicelen v163:", len(v163), v163)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v160
    _ = v161
    _ = v162
    _ = v163
    _ = v164
    _ = v165
    _ = v174
    _ = v2
    _ = v3
    _ = v4
}
