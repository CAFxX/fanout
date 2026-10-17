package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 int; F1 int32; F2 int }
type GenS4 struct { F0 string; F1 int16; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS7 struct { F0 string; F1 int; F2 uint16 }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 int8; F2 int32 }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 string; F1 uint16; F2 uint32 }
type GenS14 struct { F0 uint8; F1 int16 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint8; F1 uint64 }
type GenS18 struct { F0 uint8; F1 bool; F2 bool }
type GenS19 struct { F0 int32; F1 int64; F2 int64 }
type GenS20 struct { F0 int; F1 uint; F2 uint16 }
type GenS21 struct { F0 uint8; F1 string }
type GenS22 struct { F0 int64; F1 int8 }
type GenS23 struct { F0 int16; F1 int8 }
func genFn0() int16 {
    var v7 uint16 = 32767
    if (0 >= 0) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: -528, F1: 1}
            switch (fi8 << ((18446744073709551615) & 63)) {
            case 655:
                fmt.Println(v7)
            case 2147483647:
                if true {
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v7 -= v7
                        _ = fi10
                    }
                    v7 = (^(v7))
                    fmt.Println(v7, fi8)
                } else {
                    fmt.Println(fi8, v7)
                }
            default:
                v11 := GenS2{F0: -1, F1: "hello", F2: 42}
                _ = v11
            }
            var v12 string = "\n\t"
            _ = fi8
            _ = v12
            _ = v9
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            switch "" {
            case "hello":
                v14 := GenS1{F0: fi13, F1: -1}
                _ = v14
            case "":
                if (false && (-9223372036854775808 == 430)) {
                    var v15 int64 = 0
                    v16 := GenS2{F0: -1, F1: ((((("\n\t" + "hello") + ("foo bar" + "foo bar")) + "hello") + ("a" + ("" + "foo bar"))) + (((("a" + (("" + "") + "")) + "") + "hello") + (("hello" + ("a" + "hello")) + ""))), F2: v15}
                    if (!(4294967295 <= 18446744073709551615)) {
                        fmt.Println(v15, v15, v7)
                        var v17 int64 = -161
                        for fi18 := 0; fi18 < 19; fi18++ {
                            if false {
                                v7 = v7
                                v19 := []int{}
                                v20 := GenS3{F0: (fi13 | fi13), F1: -2147483648, F2: fi18}
                                _ = v19
                                _ = v20
                            } else {
                                v15 += (v17 % ((v15) | 1))
                                v7 += v7
                                fmt.Println(v17, fi13, v17)
                                if true {
                                    switch (("hello" + "\n\t") + ("\n\t" + "foo bar")) {
                                    case "hello":
                                        if true {
                                            if (false && (true || true)) {
                                                fmt.Println(v17)
                                                var v21 int16 = -274
                                                v15 = v15
                                                _ = v21
                                            } else {
                                                var v22 uint64 = 18446744073709551615
                                                v7 += (v7 << ((42) & 63))
                                                if false {
                                                    v23 := GenS4{F0: "a", F1: 644, F2: 257}
                                                    var v24 uint8 = 0
                                                    v15 = v17
                                                    _ = v23
                                                    _ = v24
                                                } else {
                                                    var v25 string = ""
                                                    var v26 uint = 1
                                                    var v27 uint64 = (v22 % (((^(v22))) | 1))
                                                    fmt.Println(fi13)
                                                    var v28 int8 = -3
                                                    _ = v25
                                                    _ = v26
                                                    _ = v27
                                                    _ = v28
                                                }
                                                _ = v22
                                            }
                                            var v29 string = "hello"
                                            _ = v29
                                        } else {
                                            fmt.Println(v7, fi13, v15)
                                        }
                                    case "foo bar":
                                        switch (fi13 - fi18) {
                                        case 32767:
                                            v30 := []int{(fi18 ^ 42), -1, (fi18 >> ((18446744073709551615) & 63))}
                                            _ = v30
                                        case -285:
                                            fmt.Println(fi13)
                                            v31 := GenS2{F0: -128, F1: ((("a" + "") + "foo bar") + ((("hello" + "hello") + "hello") + "a")), F2: 1}
                                            if false {
                                                v32 := GenS3{F0: (fi18 & (fi18 | (fi18 + fi18))), F1: 930, F2: fi13}
                                                _ = v32
                                            } else {
                                                if (int64(fi18) <= 0) {
                                                    var v33 int32 = 2147483647
                                                    var v34 uint = 18446744073709551615
                                                    _ = v33
                                                    _ = v34
                                                }
                                                v35 := GenS3{F0: fi18, F1: 0, F2: 42}
                                                _ = v35
                                            }
                                            v36 := GenS2{F0: 42, F1: (("hello" + "") + "a"), F2: int64(fi18)}
                                            _ = v31
                                            _ = v36
                                        }
                                    case "a":
                                        v15 = v17
                                    default:
                                        fmt.Println(v15)
                                    }
                                    v37 := []int{(^(fi13)), 1, 127, fi18}
                                    var v38 uint32 = 42
                                    _ = v37
                                    _ = v38
                                }
                                if (!(false && true)) {
                                    if false {
                                        v15 -= (-(v15))
                                        v39 := []int{fi13}
                                        _ = len(v39)
                                        v7 = (v7 / (((v7 ^ v7)) | 1))
                                        switch fi13 {
                                        case -794:
                                            fmt.Println(v7)
                                        case 4294967295:
                                            v7 -= (v7 << ((255) & 63))
                                        case -1:
                                            v7 += v7
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                var v40 int = 895
                                                _ = v40
                                            case "":
                                                if (1 < 2) {
                                                    v7 = v7
                                                    fmt.Println(v15, v17, fi18)
                                                    if len(v39) >= 2 {
                                                        v39 = v39[1:len(v39)]
                                                    }
                                                } else {
                                                    var v41 int8 = 0
                                                    switch ("" + ("hello" + (((("hello" + "foo bar") + ("foo bar" + ("a" + "\n\t"))) + "foo bar") + (("\n\t" + ("foo bar" + "a")) + "\n\t")))) {
                                                    case "foo bar":
                                                        v42 := GenS2{F0: (^(v41)), F1: "hello", F2: v15}
                                                        _ = v42
                                                    default:
                                                        switch ("hello" + "foo bar") {
                                                        case "a":
                                                            v43 := GenS2{F0: v41, F1: ("" + ("\n\t" + ("a" + "hello"))), F2: 42}
                                                            _ = v43
                                                        case "\n\t":
                                                            if (true || false) {
                                                                v17 = v17
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(fi18, fi13)
                                                    switch "a" {
                                                    case "a":
                                                        v44 := GenS5{F0: fi18}
                                                        _ = v44
                                                    }
                                                    if len(v39) >= 2 {
                                                        v39 = v39[1:len(v39)]
                                                    }
                                                    _ = v41
                                                }
                                                v17 += -1
                                                var v45 string = "a"
                                                _ = v45
                                            case "\n\t":
                                                fmt.Println(fi18, fi18, v17)
                                            }
                                        }
                                        v15 = (v15 / ((-9223372036854775808) | 1))
                                        _ = v39
                                    } else {
                                        v17 = v17
                                    }
                                    if true {
                                        var v46 float64 = 0.0
                                        _ = v46
                                    }
                                    var v47 uint64 = 876
                                    _ = v47
                                }
                            }
                            switch "a" {
                            case "foo bar":
                                var v48 string = ""
                                _ = v48
                            case "":
                                if false {
                                    if true {
                                        var v49 uint32 = 2147483648
                                        _ = v49
                                    }
                                    v50 := []int{fi18}
                                    switch "" {
                                    case "hello":
                                        if len(v50) >= 2 {
                                            v50 = v50[1:len(v50)]
                                        }
                                    case "\n\t":
                                        switch fi13 {
                                        case 42:
                                            if len(v50) >= 2 {
                                                v50 = v50[1:len(v50)]
                                            }
                                        case 32768:
                                            if false {
                                                if true {
                                                    if len(v50) >= 2 {
                                                        v50 = v50[1:len(v50)]
                                                    }
                                                    if len(v50) >= 2 {
                                                        v50 = v50[1:len(v50)]
                                                    }
                                                } else {
                                                    if true {
                                                        var v51 int = (fi13 ^ -2147483648)
                                                        v17 += -410
                                                        v17 = 0
                                                        _ = v51
                                                    } else {
                                                        if true {
                                                            v17 = int64(v7)
                                                            v52 := GenS2{F0: 100, F1: ("" + "foo bar"), F2: v17}
                                                            v15 = (v15 / (((v15 << ((608) & 63))) | 1))
                                                            _ = len(v50)
                                                            _ = v52
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "a":
                                                        v17 -= v17
                                                        if true {
                                                            v53 := GenS4{F0: "hello", F1: 819, F2: (v15 - (v17 << ((0) & 63)))}
                                                            v7 -= v7
                                                            _ = v53
                                                        } else {
                                                            v15 = -9223372036854775808
                                                            switch 32767 {
                                                            case 32768:
                                                                if (!true) {
                                                                    if (!true) {
                                                                        v50 = append(v50, (fi18 >> ((65535) & 63)))
                                                                        var v54 float64 = 1e10
                                                                        _ = v54
                                                                    } else {
                                                                        v55 := GenS5{F0: fi13}
                                                                        if ((-(fi13)) == -757) {
                                                                            if false {
                                                                                v56 := GenS2{F0: int8(fi18), F1: "\n\t", F2: -193}
                                                                                if (!true) {
                                                                                    var v57 uint8 = 2
                                                                                    _ = v57
                                                                                }
                                                                                _ = v56
                                                                            }
                                                                            var v58 float32 = 3.14159
                                                                            switch (^(fi18)) {
                                                                            case 42:
                                                                                var v59 uint8 = 143
                                                                                _ = v59
                                                                            case 798:
                                                                                if true {
                                                                                    if len(v50) > 0 {
                                                                                        v50[0] = (fi18 * -429)
                                                                                    }
                                                                                } else {
                                                                                    var v60 int = 127
                                                                                    fmt.Println(v58, v15)
                                                                                    _ = v60
                                                                                }
                                                                            case -364:
                                                                                _ = len(v50)
                                                                            default:
                                                                                var v61 uint32 = 0
                                                                                _ = v61
                                                                            }
                                                                            if true {
                                                                                v50 = append(v50, fi18)
                                                                                var v62 int16 = 1
                                                                                switch (fi13 * 42) {
                                                                                case 127:
                                                                                    v63 := GenS5{F0: fi13}
                                                                                    _ = v63
                                                                                case 2147483647:
                                                                                    fmt.Println(fi18, fi13)
                                                                                case 32768:
                                                                                    switch 2147483647 {
                                                                                    case 1:
                                                                                        fmt.Println(fi13, fi18, v62)
                                                                                    case 2147483647:
                                                                                        var v64 float64 = (1e10 / (-2.25 - float64(v58)))
                                                                                        _ = v64
                                                                                    default:
                                                                                        fmt.Println(fi18, v62, fi18)
                                                                                    }
                                                                                }
                                                                                _ = v62
                                                                            }
                                                                            v65 := GenS6{F0: 54, F1: (v7 * (^(v7))), F2: 18446744073709551615}
                                                                            v15 += v17
                                                                            v58 -= ((float32(((0.1 - (-2.25 + (1.5 - (0.0 * -2.25)))) * 1e10)) - 3.14159) * (((v58 / v58) * v58) * v58))
                                                                            _ = v58
                                                                            _ = v65
                                                                        } else {
                                                                            if len(v50) >= 2 {
                                                                                v50 = v50[1:len(v50)]
                                                                            }
                                                                        }
                                                                        v66 := GenS7{F0: "hello", F1: 42, F2: (v7 << ((2) & 63))}
                                                                        if false {
                                                                            v67 := GenS3{F0: 128, F1: 42, F2: fi13}
                                                                            fmt.Println(v15, v15)
                                                                            _ = v67
                                                                        }
                                                                        _ = v55
                                                                        _ = v66
                                                                    }
                                                                    var v68 uint16 = (v7 >> ((18446744073709551615) & 63))
                                                                    var v69 uint = 1
                                                                    _ = v68
                                                                    _ = v69
                                                                } else {
                                                                    v70 := GenS5{F0: 127}
                                                                    var v71 float32 = 1e10
                                                                    fmt.Println(v71, v71)
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            default:
                                                                v72 := GenS6{F0: 2, F1: (v7 + v7), F2: 1}
                                                                _ = v72
                                                            }
                                                        }
                                                        v15 += 42
                                                    case "hello":
                                                        if false {
                                                            v73 := GenS6{F0: 2, F1: v7, F2: 298}
                                                            _ = len(v50)
                                                            v17 = v15
                                                            _ = v73
                                                        } else {
                                                            if len(v50) >= 2 {
                                                                v50 = v50[1:len(v50)]
                                                            }
                                                            var v74 int64 = (-(v15))
                                                            _ = v74
                                                        }
                                                    default:
                                                        v7 = 408
                                                    }
                                                    v15 = (v15 / (((v17 >> ((uint(v7)) & 63))) | 1))
                                                }
                                            } else {
                                                var v75 int32 = -2147483648
                                                var v76 string = ("a" + "a")
                                                switch (fi13 - fi13) {
                                                case 1:
                                                    fmt.Println(fi18, v15, fi18)
                                                case 9223372036854775807:
                                                    fmt.Println(v75, v76, fi18)
                                                }
                                                var v77 int32 = v75
                                                switch v76 {
                                                case "a":
                                                    v76 = "foo bar"
                                                case "\n\t":
                                                    var v78 uint32 = 2
                                                    _ = v78
                                                case "foo bar":
                                                    v79 := GenS7{F0: ("\n\t" + v76), F1: fi18, F2: v7}
                                                    _ = v79
                                                }
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                            }
                                        default:
                                            var v80 int16 = 42
                                            _ = v80
                                        }
                                    case "":
                                        var v81 int = (fi13 + (fi13 & fi18))
                                        _ = v81
                                    }
                                    _ = v50
                                }
                            case "\n\t":
                                if (-1 < -1) {
                                    v17 = v17
                                    var v82 uint8 = 1
                                    _ = v82
                                }
                            }
                            v83 := GenS6{F0: 1, F1: v7, F2: 611}
                            if true {
                                fmt.Println(fi13)
                            }
                            v84 := GenS8{F0: 4294967295, F1: uint(v7)}
                            _ = fi18
                            _ = v83
                            _ = v84
                        }
                        _ = v17
                    }
                    _ = v15
                    _ = v16
                } else {
                    v85 := GenS9{F0: 42, F1: 2147483648}
                    _ = v85
                }
            case "a":
                v7 -= 1
            default:
                if (32767 == (v7 << ((42) & 63))) {
                    v86 := GenS10{F0: 2}
                    switch "foo bar" {
                    case "foo bar":
                        v7 = 32767
                    default:
                        v7 = v7
                    }
                    _ = v86
                }
            }
            v7 -= (^(v7))
            var v87 int64 = 9223372036854775807
            _ = fi13
            _ = v87
        }
        switch 1 {
        case 127:
            if false {
                if (78 >= 100) {
                    if false {
                        v88 := GenS1{F0: 2147483647, F1: -109}
                        for fi89 := 0; fi89 < 19; fi89++ {
                            fmt.Println(fi89)
                            for fi90 := 0; fi90 < 1; fi90++ {
                                var v91 uint8 = 250
                                v91 = (v91 ^ v91)
                                var v92 uint32 = 42
                                _ = fi90
                                _ = v91
                                _ = v92
                            }
                            for fi93 := 0; fi93 < 15; fi93++ {
                                v94 := GenS2{F0: 1, F1: "a", F2: -964}
                                v95 := GenS4{F0: ("a" + "hello"), F1: 0, F2: 1}
                                _ = fi93
                                _ = v94
                                _ = v95
                            }
                            _ = fi89
                        }
                        fmt.Println(v7)
                        _ = v88
                    } else {
                        for fi96 := 0; fi96 < 18; fi96++ {
                            if (true || true) {
                                v97 := []int{9223372036854775807, fi96, -1}
                                _ = v97
                            } else {
                                var v98 int64 = 389
                                v99 := GenS11{F0: -388, F1: 1, F2: -2147483648}
                                v98 -= (v98 | v98)
                                _ = v98
                                _ = v99
                            }
                            if ((false && (1.5 <= 0.1)) || ((1 == 2147483648) && ((!(!true)) && (3.14159 <= 1e10)))) {
                                var v100 float32 = float32(0.1)
                                v101 := GenS6{F0: 0, F1: (v7 | v7), F2: 408}
                                v102 := []int{(fi96 - (fi96 * (^(fi96)))), 1, 127, fi96}
                                _ = v100
                                _ = v101
                                _ = v102
                            } else {
                                for fi103 := 0; fi103 < 17; fi103++ {
                                    switch (fi103 - fi103) {
                                    case 1:
                                        if (743 != 272) {
                                            if true {
                                                if (fi96 >= 931) {
                                                    v104 := []int{fi103, fi103, 0, 2147483647}
                                                    _ = v104
                                                }
                                            }
                                            v7 = (^(v7))
                                            v7 += (v7 << ((202) & 63))
                                            v7 = v7
                                            var v105 int64 = -9223372036854775808
                                            _ = v105
                                        }
                                    case 32767:
                                        switch fi96 {
                                        case 32768:
                                            v106 := []int{}
                                            _ = v106
                                        case 2147483647:
                                            var v107 uint32 = 42
                                            _ = v107
                                        case 127:
                                            v108 := GenS1{F0: fi96, F1: -1}
                                            _ = v108
                                        }
                                    }
                                    var v109 uint64 = 2
                                    _ = fi103
                                    _ = v109
                                }
                                switch -2147483648 {
                                case 9223372036854775807:
                                    for fi110 := 0; fi110 < 16; fi110++ {
                                        fmt.Println(fi110)
                                        v7 = (^(v7))
                                        _ = fi110
                                    }
                                case 637:
                                    for fi111 := 0; fi111 < 19; fi111++ {
                                        v112 := GenS3{F0: fi111, F1: 0, F2: fi96}
                                        switch "\n\t" {
                                        case "\n\t":
                                            v113 := []int{(^(fi96)), -2147483648, fi111, fi96}
                                            _ = v113
                                        }
                                        v114 := GenS12{F0: true, F1: 2147483647}
                                        _ = fi111
                                        _ = v112
                                        _ = v114
                                    }
                                case 4294967295:
                                    var v115 uint64 = 1
                                    _ = v115
                                }
                                if (fi96 >= (^(fi96))) {
                                    switch fi96 {
                                    case 4294967295:
                                        fmt.Println(v7, fi96)
                                    }
                                    v7 += 0
                                } else {
                                    var v116 int64 = -476
                                    var v117 int32 = -1
                                    v117 -= 332
                                    _ = v116
                                    _ = v117
                                }
                            }
                            v118 := GenS1{F0: (fi96 * fi96), F1: 1}
                            _ = fi96
                            _ = v118
                        }
                        fmt.Println(v7)
                        var v119 uint32 = 65535
                        _ = v119
                    }
                    v120 := []int{}
                    v7 = v7
                    _ = v120
                }
            } else {
                var v121 bool = false
                for fi122 := 0; fi122 < 7; fi122++ {
                    var v123 int32 = 42
                    v123 -= 42
                    if true {
                        v124 := GenS10{F0: 680}
                        v125 := GenS2{F0: 120, F1: "hello", F2: -1}
                        _ = v124
                        _ = v125
                    }
                    _ = fi122
                    _ = v123
                }
                var v126 int16 = 32767
                _ = v121
                _ = v126
            }
        }
    }
    v127 := GenS9{F0: 42, F1: 523}
    v128 := []int{}
    _ = v127
    _ = v128
    _ = v7
    return 1
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{42, 2147483648, 940, 923}
    v3 := []uint32{4294967295, 230, 855, 2147483648}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 int64 = 361
    if false {
        for fi129 := 0; fi129 < 5; fi129++ {
            v1 += v1
            _ = fi129
        }
    } else {
        v1 += v1
        v130 := GenS13{F0: (("\n\t" + "hello") + ("" + "")), F1: 255, F2: 1}
        if false {
            v131 := GenS14{F0: 127, F1: 32767}
            v132 := GenS11{F0: -1, F1: -102, F2: 1}
            v133 := v4[2147483647]
            _ = v131
            _ = v132
            _ = v133
        } else {
            switch ("a" + "hello") {
            case "hello":
                if (0 > 1) {
                    var v134 float32 = float32(3.14159)
                    _ = v134
                } else {
                    switch "\n\t" {
                    case "hello":
                        v6 -= (v6 ^ v6)
                    case "\n\t":
                        v135 := GenS8{F0: 2147483648, F1: 255}
                        _ = v135
                    case "a":
                        v6 += v6
                    }
                    switch 32768 {
                    case 2147483647:
                        v136 := GenS3{F0: -1, F1: 479, F2: 0}
                        _ = v136
                    case -949:
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                    case -2147483648:
                        var v137 uint8 = 128
                        _ = v137
                    }
                    for fi138 := 0; fi138 < 17; fi138++ {
                        var v139 uint8 = 83
                        _ = fi138
                        _ = v139
                    }
                }
            case "foo bar":
                for fi140 := 0; fi140 < 7; fi140++ {
                    if (!(1 != 2)) {
                        v141 := GenS11{F0: 1, F1: -1, F2: 1}
                        switch -317 {
                        case -9223372036854775808:
                            var v142 uint = 2
                            _ = v142
                        case 877:
                            if false {
                                var v143 uint16 = 190
                                v144 := v4[fi140]
                                _ = v143
                                _ = v144
                            }
                        case 32767:
                            v1 += (-2.25 - v1)
                        default:
                            v1 -= (1e10 / v1)
                        }
                        _ = v141
                    } else {
                        for fi145 := 0; fi145 < 11; fi145++ {
                            var v146 float32 = float32(-2.25)
                            v1 = (v1 * ((0.0 / v1) + (v146 - (v1 - float32(-2.25)))))
                            _ = fi145
                            _ = v146
                        }
                        for fi147 := 0; fi147 < 1; fi147++ {
                            switch "\n\t" {
                            case "foo bar":
                                switch ("foo bar" + ("\n\t" + "\n\t")) {
                                case "":
                                    v6 = (v6 % ((42) | 1))
                                case "a":
                                    fmt.Println(fi140, v6, fi140)
                                default:
                                    var v148 int16 = 1
                                    _ = v148
                                }
                            case "hello":
                                var v149 float64 = 1e10
                                _ = v149
                            }
                            v1 = 1e10
                            if ((!false) || true) {
                                v150 := GenS15{F0: (((1e10 - (0.0 - float64(float32(float64((v1 / 3.14159)))))) - 1e10) < -2.25)}
                                if true {
                                    v151 := GenS9{F0: 1, F1: 0}
                                    var v152 bool = (false || ((1 <= 0) && (((float32(-2.25) - -2.25) + ((v1 + (v1 - v1)) / v1)) == 0.1)))
                                    var v153 uint64 = 18446744073709551615
                                    _ = v151
                                    _ = v152
                                    _ = v153
                                }
                                if len(v3) > 0 {
                                    v3[0] = 2
                                }
                                _ = v150
                            }
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v6, fi140)
                            case "\n\t":
                                v1 += v1
                            default:
                                v154 := GenS6{F0: 255, F1: 948, F2: 596}
                                _ = v154
                            }
                            var v155 float32 = v1
                            _ = len(v2)
                            _ = fi147
                            _ = v155
                        }
                    }
                    _ = fi140
                }
            }
            var v156 uint16 = 32767
            for fi157 := 0; fi157 < 4; fi157++ {
                for fi158 := 0; fi158 < 6; fi158++ {
                    fmt.Println(v1, fi158)
                    if (42 < 0) {
                        var v159 string = "foo bar"
                        v156 = (v156 & 945)
                        _ = v159
                    } else {
                        v1 = v1
                        switch fi157 {
                        case -217:
                            v156 = (^(v156))
                        }
                    }
                    var v160 uint8 = 1
                    _ = fi158
                    _ = v160
                }
                _ = fi157
            }
            _ = v156
        }
        _ = v130
    }
    var v161 int = 128
    v2 = append(v2, 1)
    var v162 float64 = (((1.5 * 1.5) * 3.14159) + ((((-2.25 + -2.25) + (0.1 - 0.1)) * 1.5) * 3.14159))
    v163 := GenS13{F0: ("foo bar" + ("" + "foo bar")), F1: 477, F2: 2}
    for fi164 := 0; fi164 < 18; fi164++ {
        v165 := GenS4{F0: ("a" + "\n\t"), F1: -1, F2: 143}
        if ((((2 == 623) || true) && (!true)) && ((true || false) && true)) {
            if false {
                fmt.Println(v1, v162, fi164)
                for fi166 := 0; fi166 < 1; fi166++ {
                    var v167 float32 = (((float32(v162) * -2.25) / ((v1 + v1) - float32(float64(1e10)))) * (-2.25 - (0.1 / -2.25)))
                    switch "" {
                    case "a":
                        v167 -= v167
                    case "":
                        v162 += ((((1e10 / -2.25) / v162) / float64(v1)) / float64((float32(v162) + 1.5)))
                    }
                    _ = fi166
                    _ = v167
                }
            }
        } else {
            var v168 int64 = v6
            _ = v168
        }
        for fi169 := 0; fi169 < 3; fi169++ {
            switch ("a" + "\n\t") {
            case "foo bar":
                if true {
                    v170 := GenS16{F0: 18446744073709551615}
                    _ = v170
                } else {
                    var v171 string = ("\n\t" + ("foo bar" + ("\n\t" + "foo bar")))
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    _ = v171
                }
            case "":
                fmt.Println(v161, v162)
            case "a":
                v1 += float32(float64(3.14159))
                var v172 int32 = 0
                v6 = (v6 / ((v6) | 1))
                _ = v172
            default:
                var v173 int64 = 42
                _ = v173
            }
            _ = fi169
        }
        _ = fi164
        _ = v165
    }
    v174 := v4[v161]
    for fi175 := 0; fi175 < 8; fi175++ {
        var v176 int16 = 1000
        if (!false) {
            if (!(42 < 1)) {
                if false {
                    var v177 uint16 = 963
                    v178 := GenS17{F0: 127, F1: 239}
                    _ = v177
                    _ = v178
                } else {
                    var v179 uint32 = v174
                    _ = v179
                }
            }
            v162 = v162
        }
        fmt.Println(v1, v174, v6)
        _ = fi175
        _ = v176
    }
    for fi180 := 0; fi180 < 14; fi180++ {
        _ = len(v2)
        fmt.Println(v162, v6, v161)
        switch "" {
        case "foo bar":
            for fi181 := 0; fi181 < 4; fi181++ {
                v1 -= v1
                if (32767 <= 2) {
                    v182 := GenS18{F0: 65, F1: (!((v162 != ((float64(v1) / 1.0) / v162)) || (!((-49 <= -882) || false)))), F2: false}
                    if true {
                        v6 = v6
                    }
                    v1 += float32(v162)
                    _ = v182
                }
                _ = fi181
            }
        }
        _ = fi180
    }
    if true {
        var v183 int = 32768
        fmt.Println(v161)
        _ = v183
    } else {
        if false {
            for fi184 := 0; fi184 < 14; fi184++ {
                var v185 int64 = v6
                if true {
                    v3 = append(v3, 314)
                    var v186 int8 = 100
                    _ = v186
                } else {
                    v187 := GenS5{F0: (-(v161))}
                    switch "a" {
                    case "":
                        var v188 int8 = 0
                        _ = v188
                    case "foo bar":
                        _ = len(v5)
                    case "a":
                        fmt.Println(fi184, v162)
                    default:
                        switch "\n\t" {
                        case "hello":
                            if (v174 < v174) {
                                if (false && false) {
                                    for fi189 := 0; fi189 < 7; fi189++ {
                                        v185 = (v6 / ((-365) | 1))
                                        _ = len(v2)
                                        _ = fi189
                                    }
                                    if false {
                                        if (("hello" + "") == ((("" + ((("\n\t" + ("" + "a")) + "hello") + "a")) + (("foo bar" + "\n\t") + "a")) + "foo bar")) {
                                            v6 -= v6
                                            var v190 int8 = 1
                                            switch ("foo bar" + "foo bar") {
                                            case "hello":
                                                if (4294967295 > 213) {
                                                    for fi191 := 0; fi191 < 18; fi191++ {
                                                        var v192 uint32 = (v174 << ((33) & 63))
                                                        switch "\n\t" {
                                                        case "a":
                                                            v174 = v192
                                                        case "hello":
                                                            if (-666 == -2147483648) {
                                                                v185 = v185
                                                                fmt.Println(v1, v174, v174)
                                                            } else {
                                                                v162 = 1.5
                                                                switch ("\n\t" + ("a" + "a")) {
                                                                case "foo bar":
                                                                    var v193 uint16 = 42
                                                                    _ = v193
                                                                case "":
                                                                    var v194 int32 = 2147483647
                                                                    _ = v194
                                                                case "hello":
                                                                    v195 := GenS7{F0: "\n\t", F1: v161, F2: 569}
                                                                    _ = v195
                                                                default:
                                                                    var v196 float32 = v1
                                                                    var v197 float64 = v162
                                                                    _ = v196
                                                                    _ = v197
                                                                }
                                                            }
                                                        }
                                                        _ = fi191
                                                        _ = v192
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        v5[(v174 - v174)] = (65535 > 1)
                                                    case "foo bar":
                                                        var v198 int8 = 95
                                                        _ = v198
                                                    case "\n\t":
                                                        v199 := GenS10{F0: 4294967295}
                                                        _ = v199
                                                    }
                                                    if (!true) {
                                                        for fi200 := 0; fi200 < 5; fi200++ {
                                                            var v201 uint16 = 0
                                                            if true {
                                                                var v202 string = ("hello" + "hello")
                                                                var v203 int64 = 0
                                                                _ = v202
                                                                _ = v203
                                                            } else {
                                                                v204 := GenS1{F0: (-(fi200)), F1: 134}
                                                                delete(v5, (v174 + (v174 / ((v174) | 1))))
                                                                _ = v204
                                                            }
                                                            _ = fi200
                                                            _ = v201
                                                        }
                                                        v3 = append(v3, 966)
                                                        if false {
                                                            var v205 float64 = ((0.0 * v162) + 0.1)
                                                            delete(v5, v174)
                                                            var v206 uint = 0
                                                            _ = v205
                                                            _ = v206
                                                        } else {
                                                            for fi207 := 0; fi207 < 1; fi207++ {
                                                                v1 = float32(v162)
                                                                switch v161 {
                                                                case 0:
                                                                    var v208 int = -369
                                                                    _ = v208
                                                                }
                                                                switch (("hello" + "\n\t") + ("hello" + "foo bar")) {
                                                                case "foo bar":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v174 - v174)
                                                                    }
                                                                default:
                                                                    v1 -= v1
                                                                }
                                                                _ = fi207
                                                            }
                                                            v209 := GenS15{F0: false}
                                                            if (!(738 != 0)) {
                                                                switch "a" {
                                                                case "hello":
                                                                    v210 := GenS12{F0: (2147483647 != -154), F1: (v161 - fi184)}
                                                                    _ = v210
                                                                case "":
                                                                    v5[v174] = (65535 != 255)
                                                                default:
                                                                    v162 -= ((0.0 * (3.14159 / v162)) * (v162 - 1.5))
                                                                }
                                                                var v211 float32 = 1.5
                                                                v190 -= (v190 ^ v190)
                                                                _ = v211
                                                            } else {
                                                                if true {
                                                                    if (((v162 + v162) - v162) < ((v162 / 1e10) - ((-2.25 + -2.25) + (v162 / 1.0)))) {
                                                                        for fi212 := 0; fi212 < 17; fi212++ {
                                                                            switch "" {
                                                                            case "hello":
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (^(v174))
                                                                                }
                                                                            case "\n\t":
                                                                                switch fi184 {
                                                                                case 9223372036854775807:
                                                                                    if false {
                                                                                        var v213 uint64 = 18446744073709551615
                                                                                        _ = v213
                                                                                    }
                                                                                case 32767:
                                                                                    var v214 string = (("foo bar" + ("foo bar" + "\n\t")) + "hello")
                                                                                    _ = v214
                                                                                default:
                                                                                    v215 := GenS19{F0: -1, F1: v185, F2: v185}
                                                                                    _ = v215
                                                                                }
                                                                            case "":
                                                                                switch "hello" {
                                                                                case "":
                                                                                    switch (v161 | (-(fi212))) {
                                                                                    case 4294967295:
                                                                                        v162 -= v162
                                                                                    default:
                                                                                        var v216 uint64 = 1
                                                                                        _ = v216
                                                                                    }
                                                                                case "\n\t":
                                                                                    switch -1 {
                                                                                    case -2147483648:
                                                                                        fmt.Println(v185)
                                                                                    }
                                                                                case "hello":
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                default:
                                                                                    var v217 string = "a"
                                                                                    _ = v217
                                                                                }
                                                                            default:
                                                                                fmt.Println(v6, v161, v162)
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            case "hello":
                                                                                v218 := GenS20{F0: 9223372036854775807, F1: 1, F2: 42}
                                                                                _ = v218
                                                                            default:
                                                                                fmt.Println(v161, v1)
                                                                            }
                                                                            _ = fi212
                                                                        }
                                                                        v4[fi184] = 42
                                                                    }
                                                                    v219 := v5[v174]
                                                                    _ = v219
                                                                }
                                                                fmt.Println(v1, v6, v174)
                                                            }
                                                            _ = v209
                                                        }
                                                    }
                                                }
                                            case "":
                                                var v220 float64 = v162
                                                _ = v220
                                            default:
                                                v185 = 9223372036854775807
                                            }
                                            _ = v190
                                        } else {
                                            v161 += v161
                                        }
                                        v221 := GenS18{F0: 87, F1: false, F2: (false || true)}
                                        _ = v221
                                    } else {
                                        var v222 uint8 = 1
                                        var v223 int32 = 42
                                        v224 := GenS21{F0: v222, F1: "foo bar"}
                                        _ = v222
                                        _ = v223
                                        _ = v224
                                    }
                                } else {
                                    fmt.Println(v6, v161)
                                    v162 -= v162
                                }
                            }
                        default:
                            v162 += 1e10
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = v174
                    }
                    _ = v187
                }
                v6 -= v6
                _ = fi184
                _ = v185
            }
        }
        v225 := GenS22{F0: 9223372036854775807, F1: 1}
        _ = v225
    }
    v6 = (v6 + v6)
    if true {
        v161 += v161
    } else {
        v4[(v161 & v161)] = 2147483647
        if false {
            var v226 int = v161
            v227 := GenS23{F0: 0, F1: 100}
            _ = v226
            _ = v227
        }
    }
    var v228 int8 = 96
    if true {
        if (v1 > -2.25) {
            for fi229 := 0; fi229 < 6; fi229++ {
                v4[(-(fi229))] = v174
                _ = fi229
            }
        }
        var v230 uint32 = 2147483647
        _ = v230
    }
    var v231 uint64 = 42
    if len(v3) > 0 {
        v3[0] = (^(v174))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v161:", v161)
    fmt.Println("v162:", v162)
    fmt.Println("v174:", v174)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v228:", v228)
    fmt.Println("v231:", v231)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v161
    _ = v162
    _ = v163
    _ = v174
    _ = v2
    _ = v228
    _ = v231
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
