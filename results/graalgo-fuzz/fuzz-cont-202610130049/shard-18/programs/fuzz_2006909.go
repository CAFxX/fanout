package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64; F1 uint }
type GenS12 struct { F0 int }
func genFn0(p0 uint32, p1 string) bool {
    var v4 int = 706
    switch "foo bar" {
    case "":
        fmt.Println(p0)
    case "a":
        p1 = p1
    case "\n\t":
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := []int{}
            fmt.Println(v4)
            _ = fi5
            _ = v6
        }
    default:
        v7 := []int{v4, -101}
        _ = v7
    }
    if (!true) {
        v8 := GenS1{F0: 18446744073709551615}
        fmt.Println(p0)
        var v9 string = "\n\t"
        _ = v8
        _ = v9
    } else {
        switch v4 {
        case -9223372036854775808:
            var v10 int8 = 1
            _ = v10
        case 4294967295:
            v11 := []int{v4, (^(v4))}
            _ = v11
        case 32768:
            var v12 uint64 = 2
            _ = v12
        }
        switch v4 {
        case -934:
            v13 := []int{v4}
            _ = v13
        case -691:
            p0 = p0
        case 2147483647:
            v14 := []int{v4}
            _ = v14
        }
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    return (1 > 671)
}
func genFn1(p0 uint16, p1 int) uint32 {
    v15 := GenS1{F0: 1}
    v16 := GenS1{F0: 2}
    p1 = (^(p1))
    p1 += p1
    p1 += 32768
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"", "hello"}
    var v3 uint16 = 547
    if true {
        switch "foo bar" {
        case "a":
            v3 = v3
        case "\n\t":
            if false {
                v3 = v3
                v17 := GenS1{F0: 523}
                _ = v17
            } else {
                var v18 uint = 4294967295
                var v19 uint8 = 155
                if (-516 < 42) {
                    if (!false) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi20 := 0; fi20 < 16; fi20++ {
                            var v21 float32 = (-2.25 / (-2.25 + 3.14159))
                            _ = fi20
                            _ = v21
                        }
                    } else {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v23 := v1[(^(fi22))]
                            if v23 {
                                v24 := GenS1{F0: 1}
                                _ = v24
                            } else {
                                v3 += v3
                            }
                            _ = fi22
                            _ = v23
                        }
                        if (!(!(true || true))) {
                            v25 := GenS1{F0: v18}
                            switch ("a" + "a") {
                            case "foo bar":
                                v26 := GenS1{F0: 2}
                                _ = v26
                            case "":
                                var v27 string = "\n\t"
                                _ = v27
                            case "\n\t":
                                switch ("foo bar" + "a") {
                                case "\n\t":
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        var v29 uint64 = 1
                                        if (42 >= 127) {
                                            v30 := GenS1{F0: (v18 ^ 393)}
                                            var v31 string = "\n\t"
                                            var v32 uint32 = 2147483647
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                        if false {
                                            fmt.Println(v18, fi28)
                                            var v33 uint16 = v3
                                            switch (fi28 * int(v3)) {
                                            case 0:
                                                var v34 uint64 = (^(v29))
                                                _ = v34
                                            case 128:
                                                v3 -= (v33 >> (((v18 % ((42) | 1))) & 63))
                                            }
                                            _ = v33
                                        }
                                        _ = fi28
                                        _ = v29
                                    }
                                case "hello":
                                    v3 -= 935
                                default:
                                    fmt.Println(v18, v18)
                                }
                            default:
                                v35 := GenS1{F0: 2}
                                _ = v35
                            }
                            v36 := GenS1{F0: (v18 * 0)}
                            _ = v25
                            _ = v36
                        }
                        if false {
                            for fi37 := 0; fi37 < 5; fi37++ {
                                fmt.Println(v3)
                                var v38 int64 = 42
                                _ = fi37
                                _ = v38
                            }
                        } else {
                            v39 := GenS2{F0: 32768, F1: 2}
                            v40 := GenS3{F0: 4294967295, F1: 0}
                            _ = v39
                            _ = v40
                        }
                    }
                    for fi41 := 0; fi41 < 6; fi41++ {
                        v18 = 4294967295
                        switch "\n\t" {
                        case "a":
                            if (false && true) {
                                v42 := GenS4{F0: (fi41 / ((fi41) | 1))}
                                v18 -= 65535
                                _ = v42
                            }
                        case "hello":
                            var v43 int32 = 478
                            _ = v43
                        }
                        if (true || ((fi41 != fi41) || true)) {
                            if false {
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v19 = v19
                                fmt.Println(fi41, v18)
                            }
                            if true {
                                v44 := GenS4{F0: (fi41 % (((fi41 & fi41)) | 1))}
                                _ = v44
                            }
                        } else {
                            if false {
                                for fi45 := 0; fi45 < 16; fi45++ {
                                    v18 = 4294967295
                                    _ = fi45
                                }
                                if false {
                                    if false {
                                        fmt.Println(fi41)
                                        v3 = v3
                                    } else {
                                        switch -742 {
                                        case 319:
                                            fmt.Println(fi41, v19, v18)
                                        case 253:
                                            for fi46 := 0; fi46 < 3; fi46++ {
                                                v47 := GenS1{F0: (v18 * (v18 ^ v18))}
                                                var v48 bool = (true && (!(("" + "a") == "hello")))
                                                switch "" {
                                                case "a":
                                                    if len(v2) > 0 {
                                                        v2[0] = "hello"
                                                    }
                                                case "hello":
                                                    v49 := GenS1{F0: (^(v18))}
                                                    _ = v49
                                                default:
                                                    _ = len(v1)
                                                }
                                                _ = fi46
                                                _ = v47
                                                _ = v48
                                            }
                                        case 42:
                                            fmt.Println(v3)
                                        default:
                                            var v50 uint64 = 42
                                            _ = v50
                                        }
                                        v51 := GenS4{F0: (fi41 - (fi41 / ((128) | 1)))}
                                        _ = v51
                                    }
                                    if (!false) {
                                        v18 = v18
                                        for fi52 := 0; fi52 < 6; fi52++ {
                                            if true {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v18 = (v18 & (^(v18)))
                                                case "a":
                                                    v19 += (v19 % (((v19 * v19)) | 1))
                                                    var v53 uint8 = 248
                                                    fmt.Println(fi52)
                                                    _ = v53
                                                case "hello":
                                                    v2 = append(v2, "")
                                                }
                                                v3 = 2
                                                switch ("" + "\n\t") {
                                                case "hello":
                                                    v18 = (v18 * v18)
                                                case "\n\t":
                                                    v19 -= v19
                                                case "a":
                                                    switch 42 {
                                                    case 32768:
                                                        if true {
                                                            var v54 uint = (v18 & (v18 / ((uint(v19)) | 1)))
                                                            _ = v54
                                                        } else {
                                                            v55 := GenS4{F0: fi41}
                                                            _ = v55
                                                        }
                                                    case -2147483648:
                                                        var v56 int64 = 9223372036854775807
                                                        _ = v56
                                                    case -460:
                                                        var v57 uint8 = (v19 - v19)
                                                        _ = v57
                                                    }
                                                }
                                            }
                                            if false {
                                                if (-340 >= 2147483647) {
                                                    switch ("a" + "") {
                                                    case "\n\t":
                                                        var v58 uint64 = 392
                                                        _ = v58
                                                    }
                                                    v59 := GenS5{F0: (v18 >> ((uint(v3)) & 63)), F1: 42}
                                                    var v60 string = "a"
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    _ = len(v1)
                                                }
                                                v3 += 2
                                            }
                                            _ = fi52
                                        }
                                        for fi61 := 0; fi61 < 10; fi61++ {
                                            var v62 float64 = 1e10
                                            v62 += (v62 + v62)
                                            v63 := GenS6{F0: 42}
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                        }
                                    }
                                    if (!(1e10 != float32(0.0))) {
                                        if true {
                                            var v64 float32 = 1e10
                                            v18 = (v18 ^ v18)
                                            _ = v64
                                        }
                                        if false {
                                            v18 -= (^(v18))
                                            fmt.Println(v18, v3, v19)
                                        }
                                    } else {
                                        switch "" {
                                        case "a":
                                            for fi65 := 0; fi65 < 4; fi65++ {
                                                _ = len(v1)
                                                if false {
                                                    fmt.Println(fi41, v3, fi41)
                                                } else {
                                                    v66 := GenS7{F0: -1}
                                                    var v67 int = (fi41 ^ (fi41 / (((fi41 >> ((42) & 63))) | 1)))
                                                    _ = v66
                                                    _ = v67
                                                }
                                                v68 := GenS2{F0: fi41, F1: 227}
                                                _ = fi65
                                                _ = v68
                                            }
                                        case "foo bar":
                                            for fi69 := 0; fi69 < 3; fi69++ {
                                                var v70 float64 = (1e10 + 1e10)
                                                fmt.Println(v70)
                                                _ = fi69
                                                _ = v70
                                            }
                                        default:
                                            v18 = 290
                                        }
                                        v3 = v3
                                    }
                                } else {
                                    _ = len(v2)
                                    v18 += 42
                                }
                            } else {
                                if (-61 > -1) {
                                    v18 += v18
                                    if false {
                                        v18 = v18
                                    }
                                }
                            }
                            v71 := GenS8{F0: "a", F1: "foo bar"}
                            for fi72 := 0; fi72 < 6; fi72++ {
                                v18 = (v18 & v18)
                                delete(v1, fi41)
                                var v73 uint64 = 120
                                _ = fi72
                                _ = v73
                            }
                            _ = v71
                        }
                        _ = fi41
                    }
                    v74 := GenS3{F0: 315, F1: uint64(v19)}
                    _ = v74
                } else {
                    var v75 int64 = -9223372036854775808
                    for fi76 := 0; fi76 < 13; fi76++ {
                        if true {
                            delete(v1, fi76)
                            v77 := GenS9{F0: false}
                            for fi78 := 0; fi78 < 16; fi78++ {
                                var v79 uint16 = (v3 >> ((v18) & 63))
                                v18 = (v18 + uint(v19))
                                var v80 int64 = (v75 * (v75 + v75))
                                switch (fi76 ^ 4294967295) {
                                case 1:
                                    v81 := GenS10{F0: -907}
                                    _ = v81
                                case 523:
                                    var v82 int = (fi78 | (fi78 ^ (fi76 | fi76)))
                                    _ = v82
                                }
                                _ = fi78
                                _ = v79
                                _ = v80
                            }
                            _ = v77
                        }
                        for fi83 := 0; fi83 < 10; fi83++ {
                            v2 = append(v2, "\n\t")
                            v3 -= (^(v3))
                            _ = fi83
                        }
                        for fi84 := 0; fi84 < 7; fi84++ {
                            var v85 int = (fi84 + 9223372036854775807)
                            switch v85 {
                            case -1:
                                var v86 bool = true
                                _ = v86
                            case 32768:
                                v87 := v1[fi84]
                                _ = v87
                            case 1:
                                fmt.Println(v19)
                            }
                            _ = fi84
                            _ = v85
                        }
                        _ = fi76
                    }
                    _ = v75
                }
                _ = v18
                _ = v19
            }
        }
        for fi88 := 0; fi88 < 2; fi88++ {
            var v89 uint64 = 1
            v2 = append(v2, "hello")
            for fi90 := 0; fi90 < 18; fi90++ {
                var v91 int64 = -9223372036854775808
                v92 := GenS8{F0: "foo bar", F1: "hello"}
                switch -9223372036854775808 {
                case 2147483647:
                    v93 := GenS2{F0: fi88, F1: (^(v89))}
                    _ = v93
                case -794:
                    v3 += (^(v3))
                case -2147483648:
                    switch "foo bar" {
                    case "":
                        if true {
                            v89 = (v89 % ((v89) | 1))
                            if len(v2) > 0 {
                                v2[0] = ("hello" + "foo bar")
                            }
                            if true {
                                var v94 float64 = 3.14159
                                v3 = v3
                                var v95 float32 = (float32((v94 - 0.1)) - float32(v94))
                                _ = v94
                                _ = v95
                            }
                            var v96 uint = 65535
                            var v97 int32 = 42
                            _ = v96
                            _ = v97
                        } else {
                            if false {
                                var v98 float64 = 0.0
                                _ = len(v1)
                                _ = v98
                            } else {
                                var v99 int64 = -364
                                if (0.1 != 1e10) {
                                    if false {
                                        v3 -= (v3 + 255)
                                        v100 := GenS3{F0: 2147483648, F1: v89}
                                        _ = v100
                                    }
                                    fmt.Println(v91)
                                    var v101 uint8 = 2
                                    _ = v101
                                } else {
                                    fmt.Println(v3)
                                    if (false && true) {
                                        switch "" {
                                        case "foo bar":
                                            if (uint8(v3) <= 194) {
                                                v99 -= v99
                                            }
                                        case "hello":
                                            switch -548 {
                                            case 127:
                                                var v102 int64 = int64(v3)
                                                _ = v102
                                            case 375:
                                                var v103 uint8 = 85
                                                _ = v103
                                            case 0:
                                                var v104 uint64 = (^(v89))
                                                v105 := v1[fi88]
                                                v99 -= v91
                                                _ = v104
                                                _ = v105
                                            default:
                                                fmt.Println(v3)
                                            }
                                        }
                                    } else {
                                        if ((true && (false && true)) || (1e10 < 3.14159)) {
                                            v91 = -158
                                            v106 := GenS6{F0: -426}
                                            var v107 int16 = 0
                                            delete(v1, (fi88 * (fi90 % ((int(v3)) | 1))))
                                            _ = v106
                                            _ = v107
                                        }
                                        v108 := GenS5{F0: 463, F1: v91}
                                        _ = v108
                                    }
                                    if false {
                                        fmt.Println(v91, v89, v99)
                                    } else {
                                        if true {
                                            _ = len(v2)
                                        }
                                        v109 := GenS1{F0: 42}
                                        v110 := GenS3{F0: 2147483648, F1: (v89 - (v89 + v89))}
                                        v111 := GenS4{F0: 0}
                                        _ = v109
                                        _ = v110
                                        _ = v111
                                    }
                                }
                                _ = v99
                            }
                        }
                    case "a":
                        if (4294967295 > 2) {
                            fmt.Println(v89, fi90)
                        } else {
                            if (-1 < -128) {
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = ("foo bar" + ("hello" + "hello"))
                                    }
                                } else {
                                    v112 := GenS6{F0: v91}
                                    _ = v112
                                }
                                switch -291 {
                                case 33:
                                    var v113 float32 = -2.25
                                    _ = v113
                                }
                                switch (fi90 << ((408) & 63)) {
                                case 321:
                                    if true {
                                        var v114 int64 = 42
                                        fmt.Println(v3, v3)
                                        switch "foo bar" {
                                        case "":
                                            var v115 int64 = (^(v91))
                                            _ = v115
                                        case "foo bar":
                                            v116 := GenS2{F0: -9223372036854775808, F1: (v89 / (((v89 * (v89 >> ((uint(v3)) & 63)))) | 1))}
                                            _ = v116
                                        case "hello":
                                            _ = len(v1)
                                        default:
                                            var v117 uint16 = v3
                                            _ = v117
                                        }
                                        _ = v114
                                    }
                                case -9223372036854775808:
                                    v118 := v1[(fi90 << ((577) & 63))]
                                    _ = v118
                                case -1:
                                    switch fi88 {
                                    case -9223372036854775808:
                                        v91 = (v91 >> ((576) & 63))
                                    case 32768:
                                        if (!false) {
                                            var v119 int8 = -111
                                            _ = v119
                                        }
                                    }
                                    var v120 uint64 = v89
                                    switch "hello" {
                                    case "\n\t":
                                        switch (((("\n\t" + "foo bar") + ("\n\t" + "\n\t")) + "a") + "\n\t") {
                                        case "\n\t":
                                            fmt.Println(fi88)
                                        }
                                    }
                                    _ = v120
                                default:
                                    fmt.Println(v91, v89, fi88)
                                }
                            }
                            switch "foo bar" {
                            case "a":
                                var v121 int8 = -1
                                _ = v121
                            default:
                                if (v89 > v89) {
                                    switch "a" {
                                    case "":
                                        v91 = v91
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        if false {
                                            _ = len(v1)
                                            if false {
                                                var v122 bool = ((v89 ^ v89) == (^(v89)))
                                                _ = v122
                                            } else {
                                                v3 = 558
                                            }
                                            fmt.Println(v91, fi88, fi90)
                                        }
                                    case "a":
                                        var v123 bool = false
                                        _ = v123
                                    case "":
                                        v89 += (v89 % ((v89) | 1))
                                    default:
                                        v124 := GenS11{F0: (v91 & v91), F1: 1}
                                        _ = v124
                                    }
                                    v89 += (^(v89))
                                }
                            }
                        }
                    case "foo bar":
                        v2 = append(v2, ("a" + "a"))
                    }
                }
                _ = fi90
                _ = v91
                _ = v92
            }
            _ = fi88
            _ = v89
        }
    }
    var v125 int16 = 1
    var v126 uint32 = 2147483648
    for fi127 := 0; fi127 < 1; fi127++ {
        v125 += v125
        v126 = genFn1((v3 >> ((4294967295) & 63)), fi127)
        for fi128 := 0; fi128 < 20; fi128++ {
            switch ("hello" + "hello") {
            case "\n\t":
                _ = len(v1)
            case "":
                if ((18446744073709551615 != 465) && false) {
                    if len(v2) > 0 {
                        v2[0] = ("foo bar" + "foo bar")
                    }
                    fmt.Println(fi127)
                    var v129 int = int(v126)
                    _ = v129
                } else {
                    v130 := v1[fi127]
                    _ = v130
                }
            case "a":
                var v131 int64 = 9223372036854775807
                _ = v131
            }
            var v132 uint32 = genFn1((^(v3)), fi128)
            var v133 float64 = 1e10
            _ = fi128
            _ = v132
            _ = v133
        }
        _ = fi127
    }
    for fi134 := 0; fi134 < 7; fi134++ {
        for fi135 := 0; fi135 < 11; fi135++ {
            var v136 uint32 = genFn1((v3 | v3), (-(fi135)))
            _ = fi135
            _ = v136
        }
        _ = fi134
    }
    _ = len(v2)
    var v137 int16 = (v125 & int16(v3))
    if false {
        if ((4294967295 >= 1) || (("a" + ("hello" + "\n\t")) < "\n\t")) {
            var v138 int = 858
            v3 = (v3 - 42)
            v139 := GenS12{F0: (v138 + (v138 / ((v138) | 1)))}
            _ = v138
            _ = v139
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v137:", v137)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v125
    _ = v126
    _ = v137
    _ = v2
    _ = v3
}
