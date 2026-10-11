package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 int8; F1 int64; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint16; F1 int16 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 uint8; F2 int64 }
type GenS13 struct { F0 uint8; F1 int16 }
type GenS14 struct { F0 int8; F1 int64 }
type GenS15 struct { F0 int8; F1 int16 }
type GenS16 struct { F0 int8; F1 uint16; F2 bool }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int64; F1 uint32 }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 int32; F1 uint16 }
type GenS23 struct { F0 uint32; F1 int8 }
type GenS24 struct { F0 int; F1 string }
type GenS25 struct { F0 uint; F1 string; F2 int64 }
type GenS26 struct { F0 int16; F1 int64 }
type GenS27 struct { F0 uint64; F1 uint8 }
type GenS28 struct { F0 uint32; F1 uint }
type GenS29 struct { F0 uint64; F1 bool }
type GenS30 struct { F0 bool; F1 bool }
type GenS31 struct { F0 int64; F1 int64 }
func genFn0(p0 bool) bool {
    v7 := []int{659, 9223372036854775807, 1, -1}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int64 = -990
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        p0 = (958 != 0)
        p0 = p0
        _ = fi10
    }
    if true {
        switch 9223372036854775807 {
        case -937:
            p0 = p0
        case 2147483647:
            p0 = (!((1.5 - -2.25) < 1.5))
        case 42:
            v11 := GenS1{F0: 602}
            _ = v11
        default:
            if (("hello" + "foo bar") <= ("a" + (("a" + "") + "foo bar"))) {
                for fi12 := 0; fi12 < 5; fi12++ {
                    p0 = false
                    if (0 >= 1) {
                        for fi13 := 0; fi13 < 7; fi13++ {
                            var v14 float32 = (float32(float64(((0.0 / 3.14159) * (-2.25 / 0.1)))) / 1.5)
                            _ = fi13
                            _ = v14
                        }
                    }
                    v15 := GenS1{F0: 151}
                    _ = fi12
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 12; fi16++ {
                    p0 = p0
                    _ = fi16
                }
                var v17 uint8 = 1
                _ = v17
            }
        }
        if true {
            if (-128 <= 100) {
                var v18 uint16 = 0
                v19 := GenS1{F0: 409}
                _ = v18
                _ = v19
            }
            var v20 int16 = -67
            var v21 string = ("a" + "\n\t")
            _ = v20
            _ = v21
        }
        if p0 {
            v22 := GenS2{F0: 1, F1: 2147483647}
            _ = v22
        } else {
            v23 := GenS3{F0: 1, F1: ""}
            for fi24 := 0; fi24 < 4; fi24++ {
                for fi25 := 0; fi25 < 15; fi25++ {
                    switch (fi25 ^ fi25) {
                    case 0:
                        if len(v7) > 0 {
                            v7[0] = (fi24 >> ((1) & 63))
                        }
                    default:
                        fmt.Println(p0)
                    }
                    var v26 int16 = -32768
                    _ = fi25
                    _ = v26
                }
                if (-128 <= -1) {
                    fmt.Println(fi24, p0)
                    fmt.Println(p0, fi24)
                }
                _ = fi24
            }
            _ = v23
        }
    } else {
        var v27 uint = 18446744073709551615
        switch 42 {
        case 1:
            switch -1 {
            case 4294967295:
                var v28 float64 = 3.14159
                _ = v28
            case 128:
                if (255 == 224) {
                    var v29 int = -578
                    p0 = p0
                    _ = v29
                } else {
                    v30 := GenS3{F0: 1, F1: "hello"}
                    fmt.Println(v27, p0)
                    _ = v30
                }
            default:
                p0 = (-250 == 42)
            }
        case 128:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case 32767:
            for fi31 := 0; fi31 < 6; fi31++ {
                if p0 {
                    v7 = append(v7, fi31)
                } else {
                    v7 = append(v7, fi31)
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v32 := GenS2{F0: (fi31 & (fi31 & fi31)), F1: 2147483648}
                    _ = v32
                }
                _ = fi31
            }
        }
        _ = v27
    }
    _ = p0
    _ = v7
    return (65535 < 65535)
}

func main() {
    v1 := []uint32{0}
    v2 := []string{"\n\t", "\n\t", "\n\t", ""}
    v3 := make(map[uint32]int)
    v4 := []int{690}
    var v5 int32 = 0
    var v6 uint64 = 889
    fmt.Println(v5, v5)
    if (0 <= 645) {
        var v33 bool = (-32768 >= -32768)
        _ = v33
    } else {
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 string = "\n\t"
            _ = fi34
            _ = v35
        }
        var v36 int64 = -1
        if false {
            v37 := GenS4{F0: "foo bar"}
            v36 -= v36
            _ = v37
        }
        _ = v36
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 bool = true
        _ = fi38
        _ = v39
    }
    v6 = 2
    if true {
        switch ("" + "hello") {
        case "foo bar":
            v5 += -2147483648
        case "a":
            v5 = v5
        case "\n\t":
            v3[65535] = -388
        }
    } else {
        fmt.Println(v5)
        v6 += v6
        v40 := GenS5{F0: 195, F1: 2147483647, F2: 2}
        _ = v40
    }
    var v41 string = "a"
    v5 -= 42
    for fi42 := 0; fi42 < 19; fi42++ {
        v6 -= 1
        for fi43 := 0; fi43 < 11; fi43++ {
            switch (v41 + "") {
            case "a":
                if (1 != 9) {
                    switch "hello" {
                    case "foo bar":
                        v6 += 1
                    default:
                        var v44 float32 = ((1.5 * float32((0.1 / 1.0))) * ((-2.25 + 1.5) - 0.0))
                        _ = v44
                    }
                    v5 -= (v5 * (v5 | v5))
                    v6 += (v6 * v6)
                }
            case "\n\t":
                var v45 uint16 = 145
                _ = v45
            default:
                var v46 string = v41
                _ = v46
            }
            if true {
                v41 = v41
                switch (v41 + ("hello" + ((((v41 + v41) + "a") + "a") + (v41 + "")))) {
                case "hello":
                    v1 = append(v1, 271)
                    v47 := GenS3{F0: 0, F1: ("" + "foo bar")}
                    var v48 int32 = v5
                    _ = v47
                    _ = v48
                case "":
                    v6 = (^(v6))
                case "a":
                    _ = len(v3)
                default:
                    v6 += 0
                }
                fmt.Println(v6, v5)
            }
            if (fi42 <= fi42) {
                fmt.Println(fi42, v5, fi42)
            } else {
                switch v41 {
                case "a":
                    switch v41 {
                    case "foo bar":
                        switch fi42 {
                        case -856:
                            if ((((3.14159 + float64((1.5 - 1e10))) + (1e10 * (((0.0 + 0.0) / 1.5) * (0.0 * 3.14159)))) > (1e10 - (1e10 + float64(1e10)))) || false) {
                                if true {
                                    v49 := GenS2{F0: (fi42 & (fi42 | fi43)), F1: 221}
                                    _ = v49
                                } else {
                                    if ((true || false) || (309 == 1)) {
                                        v5 -= v5
                                        v4 = append(v4, (fi42 * (fi43 + (fi43 / ((fi42) | 1)))))
                                    }
                                    v5 += (v5 << ((4294967295) & 63))
                                    var v50 uint = uint(v6)
                                    _ = v50
                                }
                            } else {
                                var v51 float32 = 0.0
                                _ = v51
                            }
                        case 42:
                            _ = len(v3)
                        }
                    case "a":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        v52 := GenS5{F0: -9223372036854775808, F1: 2147483648, F2: 4294967295}
                        _ = v52
                    }
                    var v53 int = (fi43 << ((76) & 63))
                    if false {
                        if len(v4) > 0 {
                            v4[0] = v53
                        }
                        if (-6 != -1) {
                            fmt.Println(v53, fi43)
                            var v54 uint = 2
                            _ = v54
                        }
                        if (!((("hello" + (v41 + v41)) > "hello") && (-1 < 0))) {
                            switch 42 {
                            case 127:
                                fmt.Println(v5, v6, fi43)
                            default:
                                v55 := v3[2147483648]
                                _ = v55
                            }
                            switch v41 {
                            case "":
                                var v56 uint32 = 517
                                _ = v56
                            case "\n\t":
                                var v57 uint32 = 4294967295
                                _ = v57
                            case "hello":
                                v58 := GenS3{F0: 1, F1: v41}
                                _ = v58
                            }
                            v5 += -1
                        }
                    }
                    _ = v53
                case "":
                    fmt.Println(v5, fi43, v41)
                case "hello":
                    if true {
                        var v59 int64 = 1
                        _ = v59
                    } else {
                        switch fi42 {
                        case 0:
                            if (255 <= 128) {
                                v60 := v3[42]
                                v6 = v6
                                _ = v60
                            } else {
                                v41 = ("" + "")
                                v41 = "foo bar"
                            }
                        case 42:
                            v41 = "\n\t"
                        }
                        v41 = v41
                    }
                }
                switch v41 {
                case "":
                    v41 = v41
                default:
                    switch (fi42 | -358) {
                    case -2147483648:
                        var v61 uint64 = (v6 * 1)
                        _ = v61
                    case 9223372036854775807:
                        fmt.Println(fi43, v5, v6)
                        v62 := v3[58]
                        _ = v62
                    default:
                        v41 = (((v41 + (((("" + (v41 + ((v41 + v41) + "hello"))) + "a") + (v41 + "foo bar")) + v41)) + v41) + v41)
                    }
                }
                switch v41 {
                case "\n\t":
                    if false {
                        v5 += v5
                        v63 := GenS6{F0: 100, F1: 42, F2: 665}
                        _ = v63
                    }
                }
            }
            _ = fi43
        }
        _ = fi42
    }
    if (400 <= 255) {
        switch v41 {
        case "":
            var v64 bool = (4294967295 <= 2)
            _ = v64
        case "hello":
            for fi65 := 0; fi65 < 3; fi65++ {
                for fi66 := 0; fi66 < 3; fi66++ {
                    v67 := GenS7{F0: "a"}
                    _ = fi66
                    _ = v67
                }
                if len(v2) > 0 {
                    v2[0] = (v41 + v41)
                }
                _ = fi65
            }
        case "\n\t":
            v41 = v41
        default:
            if true {
                for fi68 := 0; fi68 < 11; fi68++ {
                    v5 -= v5
                    v69 := GenS3{F0: 112, F1: v41}
                    _ = fi68
                    _ = v69
                }
                switch 678 {
                case 2147483647:
                    fmt.Println(v6, v6, v6)
                default:
                    v70 := GenS8{F0: 0, F1: 2}
                    _ = v70
                }
                fmt.Println(v5, v41)
            } else {
                fmt.Println(v6, v5, v41)
                v71 := GenS5{F0: 42, F1: 2, F2: 255}
                _ = v71
            }
        }
        if true {
            var v72 bool = false
            _ = v72
        }
        var v73 float64 = 0.1
        _ = v73
    } else {
        v5 = (v5 & int32(v6))
    }
    v6 = v6
    for fi74 := 0; fi74 < 8; fi74++ {
        var v75 int64 = 42
        _ = fi74
        _ = v75
    }
    if (127 <= 15) {
        if (0.1 > 1e10) {
            var v76 int64 = 361
            for fi77 := 0; fi77 < 7; fi77++ {
                switch (fi77 % ((128) | 1)) {
                case 348:
                    v76 = (v76 << ((255) & 63))
                case 2147483647:
                    if true {
                        var v78 bool = true
                        for fi79 := 0; fi79 < 18; fi79++ {
                            switch "foo bar" {
                            case "":
                                switch v41 {
                                case "hello":
                                    var v80 int16 = 1
                                    _ = v80
                                default:
                                    if (-18 <= 1) {
                                        var v81 float32 = ((3.14159 - 1e10) - 3.14159)
                                        _ = v81
                                    } else {
                                        switch (("hello" + "foo bar") + (v41 + (v41 + v41))) {
                                        case "":
                                            switch fi79 {
                                            case 680:
                                                v82 := GenS9{F0: 0, F1: -640}
                                                _ = v82
                                            case 32767:
                                                var v83 int64 = v76
                                                if (32767 > 255) {
                                                    v5 = v5
                                                    fmt.Println(fi79)
                                                    if ((v5 + v5) >= (v5 - (v5 << ((4294967295) & 63)))) {
                                                        if true {
                                                            var v84 float64 = -2.25
                                                            v85 := GenS8{F0: -9223372036854775808, F1: 132}
                                                            v86 := GenS5{F0: (v76 - (v83 >> ((2) & 63))), F1: 2, F2: 6}
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            v87 := GenS2{F0: 4294967295, F1: 42}
                                                            v41 = v41
                                                            var v88 float64 = (3.14159 - 3.14159)
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    } else {
                                                        if (fi79 == (fi77 & (fi79 / ((9223372036854775807) | 1)))) {
                                                            var v89 uint = 18446744073709551615
                                                            _ = len(v1)
                                                            _ = v89
                                                        } else {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v83 = -146
                                                            case "":
                                                                switch (^(fi79)) {
                                                                case 127:
                                                                    var v90 int8 = 40
                                                                    _ = v90
                                                                default:
                                                                    switch fi77 {
                                                                    case 4294967295:
                                                                        var v91 float64 = 1e10
                                                                        _ = v91
                                                                    case 42:
                                                                        v41 = v41
                                                                    default:
                                                                        var v92 bool = v78
                                                                        _ = v92
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                if (377 < 42) {
                                                                    if false {
                                                                        v83 += (v76 ^ (v76 * int64(v6)))
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            var v93 float64 = 0.0
                                                                            _ = v93
                                                                        case "a":
                                                                            if false {
                                                                                var v94 float32 = 0.1
                                                                                v5 += 2147483647
                                                                                v78 = v78
                                                                                _ = v94
                                                                            }
                                                                            v41 = v41
                                                                        default:
                                                                            if v78 {
                                                                                v95 := GenS10{F0: (v6 << ((42) & 63)), F1: 39}
                                                                                if v78 {
                                                                                    v96 := GenS4{F0: v41}
                                                                                    v6 = (^(v6))
                                                                                    _ = v96
                                                                                }
                                                                                if (true && (v78 || v78)) {
                                                                                    if ((1.5 - 0.0) > (1e10 + 1e10)) {
                                                                                        fmt.Println(v5)
                                                                                        if ((0.0 / 3.14159) != -2.25) {
                                                                                            var v97 int8 = 1
                                                                                            v83 -= v76
                                                                                            fmt.Println(fi79, v6)
                                                                                            _ = v97
                                                                                        } else {
                                                                                            if (v83 <= v83) {
                                                                                                switch fi77 {
                                                                                                case -2147483648:
                                                                                                    v98 := GenS7{F0: ""}
                                                                                                    _ = v98
                                                                                                default:
                                                                                                    v83 -= v76
                                                                                                }
                                                                                                var v99 uint32 = 2147483648
                                                                                                var v100 bool = (true && v78)
                                                                                                _ = v100
                                                                                                _ = v99
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        _ = len(v2)
                                                                                    }
                                                                                    var v101 string = (v41 + v41)
                                                                                    var v102 float32 = float32(-2.25)
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                } else {
                                                                                    if (643 <= 18446744073709551615) {
                                                                                        var v103 uint32 = 2147483647
                                                                                        _ = v103
                                                                                    }
                                                                                }
                                                                                _ = v95
                                                                            } else {
                                                                                if (874 < 150) {
                                                                                    var v104 uint64 = v6
                                                                                    var v105 uint16 = 0
                                                                                    fmt.Println(v83, v76)
                                                                                    _ = v104
                                                                                    _ = v105
                                                                                } else {
                                                                                    if true {
                                                                                        fmt.Println(v5)
                                                                                        var v106 uint32 = 65535
                                                                                        _ = v106
                                                                                    } else {
                                                                                        if v78 {
                                                                                            fmt.Println(v76)
                                                                                            if v78 {
                                                                                                v107 := GenS11{F0: (fi79 << ((1) & 63))}
                                                                                                if ((-(v83)) >= (v76 - (v76 / ((v76) | 1)))) {
                                                                                                    fmt.Println(fi79, v78, fi79)
                                                                                                    v108 := v3[606]
                                                                                                    var v109 float64 = 3.14159
                                                                                                    v78 = true
                                                                                                    _ = v108
                                                                                                    _ = v109
                                                                                                } else {
                                                                                                    var v110 uint8 = 213
                                                                                                    _ = len(v3)
                                                                                                    v78 = false
                                                                                                    _ = v110
                                                                                                }
                                                                                                fmt.Println(v76, fi79, v78)
                                                                                                var v111 uint64 = v6
                                                                                                v5 -= v5
                                                                                                v112 := v3[1]
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                _ = v107
                                                                                                _ = v111
                                                                                                _ = v112
                                                                                            }
                                                                                            switch v41 {
                                                                                            case "hello":
                                                                                                switch fi79 {
                                                                                                case 820:
                                                                                                    var v113 string = (((v41 + (v41 + "a")) + "\n\t") + ("foo bar" + v41))
                                                                                                    _ = v113
                                                                                                }
                                                                                            case "":
                                                                                                if (1 <= -53) {
                                                                                                    v78 = v78
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = 4294967295
                                                                                                    }
                                                                                                    delete(v3, 977)
                                                                                                } else {
                                                                                                    var v114 uint16 = 65535
                                                                                                    _ = v114
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            var v115 int32 = (^(v5))
                                                                                            var v116 float32 = (1e10 / 1.5)
                                                                                            var v117 bool = (v5 >= (v5 - v115))
                                                                                            _ = v115
                                                                                            _ = v116
                                                                                            _ = v117
                                                                                        }
                                                                                    }
                                                                                    if ((1.5 >= float64(1.5)) && false) {
                                                                                        v3[0] = (^(fi79))
                                                                                    }
                                                                                    var v118 int32 = v5
                                                                                    _ = v118
                                                                                }
                                                                            }
                                                                        }
                                                                        v83 -= v83
                                                                        v76 += v76
                                                                    }
                                                                    var v119 string = v41
                                                                    if v78 {
                                                                        v120 := GenS6{F0: -128, F1: (-(v76)), F2: 2}
                                                                        v4 = append(v4, fi77)
                                                                        if true {
                                                                            fmt.Println(fi79, v119)
                                                                            if true {
                                                                                if (!false) {
                                                                                    v41 = v41
                                                                                    var v121 uint32 = 878
                                                                                    _ = v121
                                                                                } else {
                                                                                    var v122 int = (fi77 >> ((0) & 63))
                                                                                    switch 4294967295 {
                                                                                    case 777:
                                                                                        switch (fi79 & 639) {
                                                                                        case 538:
                                                                                            var v123 uint = 523
                                                                                            _ = v123
                                                                                        default:
                                                                                            fmt.Println(v83, v122)
                                                                                        }
                                                                                    case 42:
                                                                                        v6 = (v6 + (v6 & v6))
                                                                                    case -2147483648:
                                                                                        if (!false) {
                                                                                            v5 = 2147483647
                                                                                            v78 = v78
                                                                                            var v124 float64 = (1e10 + 3.14159)
                                                                                            v1 = append(v1, 572)
                                                                                            v41 = v41
                                                                                            _ = v124
                                                                                        }
                                                                                    default:
                                                                                        v6 += (^(v6))
                                                                                    }
                                                                                    var v125 float32 = -2.25
                                                                                    _ = v122
                                                                                    _ = v125
                                                                                }
                                                                                if ((!v78) && v78) {
                                                                                    if v78 {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 42
                                                                                        }
                                                                                        if ((v6 - (^(v6))) != 18446744073709551615) {
                                                                                            v126 := GenS12{F0: fi79, F1: 127, F2: 660}
                                                                                            switch (v119 + v41) {
                                                                                            case "a":
                                                                                                if v78 {
                                                                                                    var v127 float64 = ((float64(0.1) * 1.5) * float64((0.1 / ((float32((0.0 / 3.14159)) - (3.14159 / 3.14159)) + 1e10))))
                                                                                                    _ = v127
                                                                                                }
                                                                                            case "hello":
                                                                                                v76 -= v76
                                                                                            case "\n\t":
                                                                                                v78 = v78
                                                                                                v6 += (v6 >> ((255) & 63))
                                                                                            }
                                                                                            v6 -= (v6 % (((^(v6))) | 1))
                                                                                            _ = v126
                                                                                        } else {
                                                                                            var v128 float64 = 1.5
                                                                                            _ = v128
                                                                                        }
                                                                                    } else {
                                                                                        var v129 float64 = float64(1.5)
                                                                                        switch (("\n\t" + (v41 + v41)) + "hello") {
                                                                                        case "":
                                                                                            var v130 uint16 = 2
                                                                                            _ = v130
                                                                                        case "foo bar":
                                                                                            v131 := GenS6{F0: -128, F1: (v83 % ((0) | 1)), F2: 754}
                                                                                            _ = v131
                                                                                        case "a":
                                                                                            var v132 uint = 814
                                                                                            v119 = "\n\t"
                                                                                            _ = v132
                                                                                        }
                                                                                        v3[1000] = fi77
                                                                                        _ = v129
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if v78 {
                                                                                    if (false || (278 == 2147483647)) {
                                                                                        v133 := GenS2{F0: fi79, F1: 2147483648}
                                                                                        _ = v133
                                                                                    } else {
                                                                                        v134 := GenS9{F0: 2, F1: -1}
                                                                                        v135 := GenS12{F0: (-(fi77)), F1: 218, F2: v83}
                                                                                        var v136 uint64 = v6
                                                                                        switch ("foo bar" + (((v119 + "") + (v41 + "")) + "a")) {
                                                                                        case "hello":
                                                                                            var v137 int16 = -119
                                                                                            _ = v137
                                                                                        default:
                                                                                            var v138 bool = ((!(!("hello" < "foo bar"))) || v78)
                                                                                            _ = v138
                                                                                        }
                                                                                        var v139 float64 = -2.25
                                                                                        v41 = v41
                                                                                        v5 -= (-(v5))
                                                                                        _ = v134
                                                                                        _ = v135
                                                                                        _ = v136
                                                                                        _ = v139
                                                                                    }
                                                                                    if (-88 > -1) {
                                                                                        v76 = v76
                                                                                        var v140 bool = v78
                                                                                        switch fi77 {
                                                                                        case 2147483647:
                                                                                            switch -1 {
                                                                                            case -871:
                                                                                                var v141 string = v41
                                                                                                _ = v141
                                                                                            case -833:
                                                                                                var v142 int8 = 112
                                                                                                _ = v142
                                                                                            case 284:
                                                                                                if true {
                                                                                                    fmt.Println(v119, v78)
                                                                                                    v119 = v41
                                                                                                }
                                                                                            default:
                                                                                                delete(v3, 841)
                                                                                            }
                                                                                        }
                                                                                        _ = v140
                                                                                    }
                                                                                    if false {
                                                                                        var v143 uint32 = 2147483648
                                                                                        var v144 bool = (!(0.1 >= 0.1))
                                                                                        if (42 == 32767) {
                                                                                            if (!v144) {
                                                                                                delete(v3, v143)
                                                                                                var v145 int = fi79
                                                                                                fmt.Println(v145)
                                                                                                _ = v145
                                                                                            }
                                                                                            var v146 uint8 = 0
                                                                                            v146 = (v146 ^ (v146 - (v146 | v146)))
                                                                                            _ = v146
                                                                                        } else {
                                                                                            _ = len(v2)
                                                                                        }
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                    }
                                                                                    fmt.Println(fi79)
                                                                                    fmt.Println(v76, v41, v83)
                                                                                    v6 -= (^(v6))
                                                                                    var v147 int32 = v5
                                                                                    _ = v147
                                                                                }
                                                                                fmt.Println(fi79, v5)
                                                                                v148 := GenS3{F0: 44, F1: v41}
                                                                                _ = v148
                                                                            }
                                                                            fmt.Println(v78)
                                                                        } else {
                                                                            var v149 int = 0
                                                                            fmt.Println(v78)
                                                                            _ = v149
                                                                        }
                                                                        _ = v120
                                                                    } else {
                                                                        var v150 uint16 = 995
                                                                        if len(v2) > 0 {
                                                                            v2[0] = ""
                                                                        }
                                                                        _ = v150
                                                                    }
                                                                    _ = v119
                                                                } else {
                                                                    fmt.Println(v76)
                                                                }
                                                            default:
                                                                fmt.Println(v83)
                                                            }
                                                            v83 += (v83 & v83)
                                                            fmt.Println(v6, v5)
                                                            if false {
                                                                if (v78 && false) {
                                                                    if true {
                                                                        var v151 float32 = 1.5
                                                                        v6 -= v6
                                                                        var v152 string = ""
                                                                        _ = v151
                                                                        _ = v152
                                                                    } else {
                                                                        fmt.Println(v41)
                                                                        v153 := GenS12{F0: -9223372036854775808, F1: 255, F2: 1}
                                                                        v2 = append(v2, "")
                                                                        _ = v153
                                                                    }
                                                                } else {
                                                                    delete(v3, 2)
                                                                    v5 = (v5 - (v5 * v5))
                                                                    v76 += v83
                                                                }
                                                                fmt.Println(v76, fi79)
                                                                v154 := v3[2147483648]
                                                                var v155 uint8 = 213
                                                                var v156 int64 = v76
                                                                var v157 string = v41
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v6 = 2
                                                                _ = v154
                                                                _ = v155
                                                                _ = v156
                                                                _ = v157
                                                            } else {
                                                                fmt.Println(v83, v5, fi79)
                                                            }
                                                            v5 = v5
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi77)
                                                _ = v83
                                            case -9223372036854775808:
                                                v76 += (v76 - 561)
                                            default:
                                                switch fi79 {
                                                case 821:
                                                    if v78 {
                                                        var v158 uint = 163
                                                        fmt.Println(fi79)
                                                        _ = v158
                                                    } else {
                                                        var v159 uint32 = 2147483647
                                                        switch fi77 {
                                                        case -9223372036854775808:
                                                            var v160 uint32 = (v159 * (^(v159)))
                                                            _ = v160
                                                        }
                                                        switch v41 {
                                                        case "foo bar":
                                                            var v161 uint16 = 42
                                                            _ = v161
                                                        case "":
                                                            _ = len(v4)
                                                        default:
                                                            v76 -= (v76 >> ((65535) & 63))
                                                        }
                                                        _ = v159
                                                    }
                                                case -1:
                                                    switch v41 {
                                                    case "a":
                                                        switch v41 {
                                                        case "hello":
                                                            var v162 bool = true
                                                            _ = v162
                                                        case "":
                                                            var v163 uint64 = (v6 ^ v6)
                                                            fmt.Println(v41, v76)
                                                            _ = len(v1)
                                                            _ = v163
                                                        case "\n\t":
                                                            if v78 {
                                                                switch (("\n\t" + v41) + "a") {
                                                                case "hello":
                                                                    v41 = (v41 + ("hello" + v41))
                                                                default:
                                                                    v78 = v78
                                                                }
                                                                v164 := GenS13{F0: 255, F1: -14}
                                                                v165 := GenS11{F0: (-(fi77))}
                                                                _ = v164
                                                                _ = v165
                                                            } else {
                                                                v76 += 9223372036854775807
                                                            }
                                                        }
                                                    case "hello":
                                                        if (127 > 0) {
                                                            v78 = v78
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v166 float32 = (3.14159 * float32(((1.5 / (((-2.25 * (float64(0.1) * 3.14159)) - (1.5 + (0.1 - 1e10))) / 1e10)) / -2.25)))
                                                            _ = v166
                                                        }
                                                    case "":
                                                        if (v78 && (-2.25 <= 0.0)) {
                                                            var v167 bool = v78
                                                            v168 := GenS9{F0: 219, F1: -32768}
                                                            _ = v167
                                                            _ = v168
                                                        } else {
                                                            v169 := GenS14{F0: -118, F1: (v76 & v76)}
                                                            _ = v169
                                                        }
                                                    default:
                                                        v6 += v6
                                                    }
                                                    v78 = (!((^(v6)) < v6))
                                                    if false {
                                                        v5 -= v5
                                                    } else {
                                                        v170 := GenS15{F0: 100, F1: -879}
                                                        switch v41 {
                                                        case "foo bar":
                                                            v5 -= (v5 * v5)
                                                        case "a":
                                                            if (v78 && (3.14159 <= -2.25)) {
                                                                var v171 float64 = ((0.1 - (1e10 + (0.1 - (1.5 - -2.25)))) - 1.5)
                                                                if (703 > 2147483647) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v41
                                                                    }
                                                                    v172 := GenS13{F0: 137, F1: -32768}
                                                                    switch (v41 + ("foo bar" + "")) {
                                                                    case "\n\t":
                                                                        v6 += (^(v6))
                                                                    default:
                                                                        var v173 int8 = 1
                                                                        var v174 uint16 = 32767
                                                                        var v175 float64 = v171
                                                                        switch 32767 {
                                                                        case -981:
                                                                            var v176 uint = 2
                                                                            _ = v176
                                                                        case 521:
                                                                            if ((!true) && (false || v78)) {
                                                                                if (-32768 >= 639) {
                                                                                    v173 -= v173
                                                                                    if (552 <= 1000) {
                                                                                        var v177 float64 = ((v171 * 0.1) / (v175 - 3.14159))
                                                                                        v3[380] = (fi77 ^ (fi77 - 0))
                                                                                        var v178 uint32 = 0
                                                                                        _ = v177
                                                                                        _ = v178
                                                                                    } else {
                                                                                        if v78 {
                                                                                            if v78 {
                                                                                                var v179 uint16 = (v174 * 0)
                                                                                                v6 -= (v6 ^ (v6 | v6))
                                                                                                v180 := GenS16{F0: v173, F1: 567, F2: ((^(v174)) >= 65535)}
                                                                                                _ = v179
                                                                                                _ = v180
                                                                                            } else {
                                                                                                v181 := GenS17{F0: 42}
                                                                                                if true {
                                                                                                    var v182 uint8 = 127
                                                                                                    v183 := GenS18{F0: 948}
                                                                                                    delete(v3, 42)
                                                                                                    var v184 float32 = 3.14159
                                                                                                    v173 += 127
                                                                                                    v173 += int8(v182)
                                                                                                    _ = v182
                                                                                                    _ = v183
                                                                                                    _ = v184
                                                                                                } else {
                                                                                                    switch fi77 {
                                                                                                    case 127:
                                                                                                        var v185 uint64 = v6
                                                                                                        _ = v185
                                                                                                    case -9223372036854775808:
                                                                                                        fmt.Println(fi77, v76)
                                                                                                    case 4294967295:
                                                                                                        var v186 int = fi77
                                                                                                        _ = v186
                                                                                                    }
                                                                                                }
                                                                                                _ = v181
                                                                                            }
                                                                                            var v187 float32 = 0.1
                                                                                            _ = v187
                                                                                        } else {
                                                                                            var v188 int = (^(fi77))
                                                                                            switch (-(fi79)) {
                                                                                            case 128:
                                                                                                v188 = (fi79 + (fi79 * v188))
                                                                                            case 2147483647:
                                                                                                v41 = v41
                                                                                            case -1:
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            _ = v188
                                                                                        }
                                                                                    }
                                                                                    v78 = (true || v78)
                                                                                } else {
                                                                                    v76 -= (v76 + (v76 + v76))
                                                                                    var v189 int = fi79
                                                                                    var v190 int32 = (v5 / (((v5 - 2147483647)) | 1))
                                                                                    _ = v189
                                                                                    _ = v190
                                                                                }
                                                                            } else {
                                                                                var v191 uint32 = 334
                                                                                switch fi79 {
                                                                                case 32768:
                                                                                    var v192 int8 = v173
                                                                                    _ = v192
                                                                                case -9223372036854775808:
                                                                                    switch 42 {
                                                                                    case 128:
                                                                                        switch v41 {
                                                                                        case "a":
                                                                                            fmt.Println(v171, v175, v175)
                                                                                        case "foo bar":
                                                                                            v171 += v171
                                                                                        default:
                                                                                            var v193 int64 = v76
                                                                                            _ = v193
                                                                                        }
                                                                                    case 880:
                                                                                        fmt.Println(v175)
                                                                                        v194 := GenS8{F0: (v76 << ((137) & 63)), F1: 46}
                                                                                        var v195 int = 128
                                                                                        _ = v194
                                                                                        _ = v195
                                                                                    }
                                                                                default:
                                                                                    _ = len(v1)
                                                                                    var v196 string = "foo bar"
                                                                                    _ = v196
                                                                                }
                                                                                _ = v191
                                                                            }
                                                                        case -2147483648:
                                                                            switch fi79 {
                                                                            case -9223372036854775808:
                                                                                v197 := GenS19{F0: (v76 - -1), F1: 935}
                                                                                _ = v197
                                                                            case 818:
                                                                                v173 = (^(v173))
                                                                            default:
                                                                                v76 -= 0
                                                                            }
                                                                        }
                                                                        switch fi77 {
                                                                        case 4294967295:
                                                                            if false {
                                                                                v198 := GenS7{F0: (v41 + v41)}
                                                                                v199 := GenS2{F0: 783, F1: 756}
                                                                                v78 = (((v175 - v175) != (3.14159 * 1e10)) || (!v78))
                                                                                _ = v198
                                                                                _ = v199
                                                                            }
                                                                        case 584:
                                                                            v200 := v3[925]
                                                                            _ = v200
                                                                        case 1:
                                                                            var v201 uint16 = uint16(v6)
                                                                            _ = v201
                                                                        }
                                                                        v41 = (v41 + (v41 + "foo bar"))
                                                                        _ = v173
                                                                        _ = v174
                                                                        _ = v175
                                                                    }
                                                                    _ = v172
                                                                }
                                                                _ = len(v1)
                                                                _ = v171
                                                            } else {
                                                                v202 := GenS2{F0: (fi77 % ((fi79) | 1)), F1: 2147483647}
                                                                switch (v41 + (v41 + (v41 + (((v41 + "\n\t") + v41) + v41)))) {
                                                                case "foo bar":
                                                                    v203 := GenS20{F0: (v76 ^ v76)}
                                                                    _ = v203
                                                                case "hello":
                                                                    delete(v3, 4294967295)
                                                                case "a":
                                                                    var v204 int32 = (v5 ^ v5)
                                                                    _ = v204
                                                                }
                                                                _ = v202
                                                            }
                                                        default:
                                                            if (!false) {
                                                                switch (fi77 / ((fi79) | 1)) {
                                                                case -285:
                                                                    if (2 < 57) {
                                                                        if v78 {
                                                                            var v205 uint16 = 1
                                                                            _ = v205
                                                                        } else {
                                                                            v41 = v41
                                                                            _ = len(v1)
                                                                            if true {
                                                                                if ((^(fi77)) < fi77) {
                                                                                    if (((v76 * v76) <= 795) && (1000 != 336)) {
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = v41
                                                                                        }
                                                                                    }
                                                                                    v206 := GenS4{F0: ("foo bar" + v41)}
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = v41
                                                                                    }
                                                                                    v78 = (4294967295 < 2)
                                                                                    var v207 int32 = v5
                                                                                    switch v41 {
                                                                                    case "":
                                                                                        var v208 uint64 = 42
                                                                                        if v78 {
                                                                                            switch -789 {
                                                                                            case 2147483647:
                                                                                                if true {
                                                                                                    v209 := GenS17{F0: 465}
                                                                                                    var v210 uint64 = 590
                                                                                                    _ = v209
                                                                                                    _ = v210
                                                                                                }
                                                                                            case 32767:
                                                                                                var v211 int8 = 42
                                                                                                _ = v211
                                                                                            }
                                                                                            v76 += (^(v76))
                                                                                        } else {
                                                                                            v212 := GenS21{F0: v76}
                                                                                            v213 := v3[42]
                                                                                            v214 := GenS22{F0: v207, F1: 32767}
                                                                                            switch v41 {
                                                                                            case "foo bar":
                                                                                                if (1.5 == (((0.1 * 0.0) * 1.5) - 0.0)) {
                                                                                                    var v215 uint32 = 257
                                                                                                    var v216 uint = 4294967295
                                                                                                    if (v78 || ((^(v215)) == v215)) {
                                                                                                        v207 += v207
                                                                                                        if false {
                                                                                                            v217 := GenS6{F0: 42, F1: 425, F2: 0}
                                                                                                            _ = v217
                                                                                                        }
                                                                                                        v218 := GenS13{F0: 42, F1: -1}
                                                                                                        _ = v218
                                                                                                    } else {
                                                                                                        fmt.Println(v6, v5)
                                                                                                        fmt.Println(v78, v5, v216)
                                                                                                    }
                                                                                                    var v219 string = v41
                                                                                                    if (v208 == v208) {
                                                                                                        var v220 uint16 = 42
                                                                                                        var v221 uint64 = v6
                                                                                                        if true {
                                                                                                            v222 := GenS1{F0: v6}
                                                                                                            _ = v222
                                                                                                        }
                                                                                                        _ = v220
                                                                                                        _ = v221
                                                                                                    }
                                                                                                    _ = v215
                                                                                                    _ = v216
                                                                                                    _ = v219
                                                                                                } else {
                                                                                                    v76 -= (^(v76))
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v76 = v76
                                                                                            case "":
                                                                                                fmt.Println(v5, fi79)
                                                                                            }
                                                                                            _ = v212
                                                                                            _ = v213
                                                                                            _ = v214
                                                                                        }
                                                                                        var v223 string = v41
                                                                                        _ = v208
                                                                                        _ = v223
                                                                                    case "hello":
                                                                                        var v224 int32 = v207
                                                                                        _ = v224
                                                                                    case "a":
                                                                                        v225 := GenS6{F0: 42, F1: (v76 * 417), F2: 547}
                                                                                        _ = v225
                                                                                    }
                                                                                    _ = v206
                                                                                    _ = v207
                                                                                } else {
                                                                                    var v226 uint16 = 255
                                                                                    if v78 {
                                                                                        var v227 float64 = -2.25
                                                                                        v3[97] = (fi79 - (-(fi79)))
                                                                                        _ = v227
                                                                                    } else {
                                                                                        if (((!(42 != 1)) || v78) && (v78 && false)) {
                                                                                            fmt.Println(v76, v41, fi77)
                                                                                            v6 -= 18446744073709551615
                                                                                        }
                                                                                        if ((1.5 / (1.5 + (float32(0.1) * 0.1))) == 3.14159) {
                                                                                            v78 = v78
                                                                                            var v228 int16 = 149
                                                                                            _ = v228
                                                                                        } else {
                                                                                            _ = len(v3)
                                                                                            v41 = "hello"
                                                                                        }
                                                                                        var v229 uint64 = v6
                                                                                        _ = v229
                                                                                    }
                                                                                    var v230 int16 = 369
                                                                                    _ = v226
                                                                                    _ = v230
                                                                                }
                                                                            } else {
                                                                                var v231 uint8 = 128
                                                                                v76 += (^(v76))
                                                                                v78 = false
                                                                                _ = v231
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                v232 := GenS23{F0: 493, F1: 127}
                                                                var v233 uint16 = 495
                                                                v78 = v78
                                                                _ = v232
                                                                _ = v233
                                                            }
                                                        }
                                                        _ = v170
                                                    }
                                                    var v234 uint8 = 255
                                                    _ = v234
                                                default:
                                                    v235 := GenS14{F0: 100, F1: (^(v76))}
                                                    _ = v235
                                                }
                                            }
                                        default:
                                            fmt.Println(v41)
                                        }
                                    }
                                }
                            case "hello":
                                _ = len(v3)
                            default:
                                var v236 string = v41
                                _ = v236
                            }
                            var v237 int = (fi77 >> ((909) & 63))
                            if v78 {
                                var v238 bool = true
                                v239 := GenS24{F0: 537, F1: "hello"}
                                var v240 uint16 = 32767
                                _ = v238
                                _ = v239
                                _ = v240
                            } else {
                                var v241 bool = v78
                                _ = len(v3)
                                var v242 int64 = (v76 + (-(v76)))
                                _ = v241
                                _ = v242
                            }
                            _ = fi79
                            _ = v237
                        }
                        _ = v78
                    }
                case -1:
                    v3[431] = (fi77 * (fi77 ^ 42))
                }
                _ = fi77
            }
            _ = v76
        }
        v5 -= v5
    }
    for fi243 := 0; fi243 < 8; fi243++ {
        v41 = v41
        _ = fi243
    }
    switch 1 {
    case -9223372036854775808:
        if false {
            if (1 <= 4294967295) {
                var v244 uint16 = 255
                _ = v244
            } else {
                v41 = ((v41 + "hello") + v41)
                if (2 > 988) {
                    v5 = v5
                    var v245 int32 = int32(v6)
                    v41 = ("foo bar" + v41)
                    _ = v245
                }
            }
            var v246 int16 = 32767
            if (!(!(((("foo bar" + v41) + v41) + "") > "\n\t"))) {
                var v247 int64 = 9223372036854775807
                var v248 uint32 = 2
                if (395 <= (v5 / ((v5) | 1))) {
                    switch -573 {
                    case 0:
                        v249 := GenS13{F0: 128, F1: (-(v246))}
                        _ = v249
                    }
                } else {
                    var v250 float32 = 3.14159
                    v250 += 0.1
                    v6 = (v6 % ((v6) | 1))
                    _ = v250
                }
                _ = v247
                _ = v248
            } else {
                for fi251 := 0; fi251 < 17; fi251++ {
                    if (((-2.25 * float64(3.14159)) <= 1.5) || (true || false)) {
                        v41 = v41
                        for fi252 := 0; fi252 < 19; fi252++ {
                            v41 = ((v41 + v41) + ("" + "hello"))
                            var v253 int8 = -128
                            _ = fi252
                            _ = v253
                        }
                        v254 := GenS25{F0: 335, F1: "", F2: 9223372036854775807}
                        _ = v254
                    } else {
                        if (false || ((v6 - (v6 % (((v6 | (^(v6)))) | 1))) <= 0)) {
                            v255 := GenS26{F0: -52, F1: -9223372036854775808}
                            if (true && (!false)) {
                                _ = len(v1)
                                var v256 uint64 = v6
                                _ = v256
                            } else {
                                switch (-(fi251)) {
                                case 42:
                                    v246 -= v246
                                default:
                                    fmt.Println(fi251)
                                }
                            }
                            switch fi251 {
                            case 128:
                                fmt.Println(fi251, v41, fi251)
                            case 316:
                                v257 := GenS9{F0: 536, F1: v246}
                                _ = v257
                            case 0:
                                fmt.Println(v6)
                            default:
                                switch 568 {
                                case -9223372036854775808:
                                    var v258 bool = false
                                    _ = v258
                                case 0:
                                    var v259 string = ("foo bar" + "")
                                    _ = v259
                                }
                            }
                            _ = v255
                        }
                        if (v6 < (v6 - v6)) {
                            var v260 uint8 = 189
                            var v261 int16 = -676
                            var v262 uint64 = v6
                            _ = v260
                            _ = v261
                            _ = v262
                        } else {
                            v263 := GenS27{F0: v6, F1: 128}
                            _ = v263
                        }
                        var v264 uint64 = (^(v6))
                        _ = v264
                    }
                    var v265 uint64 = 18446744073709551615
                    _ = fi251
                    _ = v265
                }
            }
            _ = v246
        } else {
            if (((v6 % ((v6) | 1)) < v6) || true) {
                _ = len(v1)
                var v266 int8 = 96
                _ = v266
            } else {
                var v267 int64 = 363
                fmt.Println(v6)
                for fi268 := 0; fi268 < 20; fi268++ {
                    switch (fi268 % (((-(fi268))) | 1)) {
                    case 127:
                        v6 = 2
                    case 223:
                        fmt.Println(v5)
                    case 42:
                        v269 := GenS26{F0: 0, F1: 1}
                        _ = v269
                    default:
                        if true {
                            var v270 bool = false
                            delete(v3, 821)
                            _ = len(v4)
                            _ = v270
                        } else {
                            var v271 uint64 = (v6 ^ v6)
                            v272 := GenS2{F0: fi268, F1: 1}
                            _ = v271
                            _ = v272
                        }
                    }
                    var v273 uint8 = 128
                    _ = fi268
                    _ = v273
                }
                _ = v267
            }
            switch "" {
            case "\n\t":
                if false {
                    switch 4294967295 {
                    case -718:
                        v41 = ""
                    case 2147483647:
                        switch "hello" {
                        case "":
                            for fi274 := 0; fi274 < 18; fi274++ {
                                var v275 float64 = (-2.25 / (1.5 / ((1e10 - (0.0 + 3.14159)) - (-2.25 / ((0.0 - 0.1) * 3.14159)))))
                                switch (fi274 % (((fi274 << ((0) & 63))) | 1)) {
                                case -11:
                                    var v276 int64 = -753
                                    _ = v276
                                case 32768:
                                    switch fi274 {
                                    case 32768:
                                        if (!(false || true)) {
                                            _ = len(v2)
                                            v41 = (((v41 + "hello") + v41) + "a")
                                        } else {
                                            var v277 int32 = v5
                                            for fi278 := 0; fi278 < 12; fi278++ {
                                                switch (-(fi278)) {
                                                case 32767:
                                                    var v279 float32 = (1.5 - 1.5)
                                                    _ = v279
                                                case 32768:
                                                    fmt.Println(v5, v275)
                                                }
                                                _ = fi278
                                            }
                                            _ = v277
                                        }
                                    default:
                                        v280 := GenS5{F0: 42, F1: 65535, F2: 652}
                                        _ = v280
                                    }
                                case 18:
                                    if (!true) {
                                        if true {
                                            fmt.Println(v6, v275, v5)
                                        }
                                        v5 = int32(fi274)
                                        for fi281 := 0; fi281 < 1; fi281++ {
                                            fmt.Println(fi274, v6)
                                            var v282 int32 = v5
                                            if (1 == fi281) {
                                                var v283 bool = false
                                                _ = v283
                                            } else {
                                                v6 += 2
                                                v5 = (v5 | (v282 - v282))
                                                v282 = v5
                                            }
                                            _ = fi281
                                            _ = v282
                                        }
                                    } else {
                                        switch ("hello" + v41) {
                                        case "":
                                            var v284 uint64 = 0
                                            _ = v284
                                        case "a":
                                            v5 -= (-(v5))
                                        case "foo bar":
                                            for fi285 := 0; fi285 < 11; fi285++ {
                                                if len(v2) > 0 {
                                                    v2[0] = "foo bar"
                                                }
                                                _ = fi285
                                            }
                                        }
                                        v286 := GenS2{F0: (fi274 + (fi274 * (-(fi274)))), F1: 4294967295}
                                        v5 = 190
                                        _ = v286
                                    }
                                }
                                _ = fi274
                                _ = v275
                            }
                        case "foo bar":
                            for fi287 := 0; fi287 < 5; fi287++ {
                                for fi288 := 0; fi288 < 16; fi288++ {
                                    var v289 int64 = 0
                                    switch (fi288 % ((fi288) | 1)) {
                                    case 177:
                                        var v290 bool = false
                                        _ = v290
                                    case 4294967295:
                                        var v291 float32 = 0.0
                                        _ = v291
                                    case 0:
                                        v4 = append(v4, (fi287 * (fi288 ^ (fi288 * (fi287 ^ (fi287 - fi288))))))
                                    }
                                    v41 = v41
                                    switch (v41 + "hello") {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "\n\t":
                                            v292 := GenS15{F0: 37, F1: int16(v6)}
                                            _ = v292
                                        case "":
                                            var v293 int32 = -2147483648
                                            _ = v293
                                        default:
                                            fmt.Println(v6)
                                        }
                                    case "":
                                        fmt.Println(v41, fi287)
                                    case "hello":
                                        var v294 bool = (127 == 255)
                                        _ = v294
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = fi287
                                        }
                                    }
                                    switch fi287 {
                                    case -2147483648:
                                        var v295 uint64 = v6
                                        _ = v295
                                    }
                                    v289 += v289
                                    v5 = v5
                                    _ = fi288
                                    _ = v289
                                }
                                _ = len(v3)
                                if false {
                                    var v296 int64 = -964
                                    for fi297 := 0; fi297 < 18; fi297++ {
                                        fmt.Println(fi287)
                                        var v298 bool = true
                                        var v299 uint16 = 616
                                        _ = fi297
                                        _ = v298
                                        _ = v299
                                    }
                                    _ = v296
                                }
                                _ = fi287
                            }
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                        }
                    }
                    v5 = -2147483648
                }
            default:
                v300 := GenS28{F0: 46, F1: 255}
                _ = v300
            }
            if true {
                var v301 bool = (true || true)
                if (131 < 1) {
                    switch 137 {
                    case 42:
                        v302 := GenS29{F0: (v6 >> ((65535) & 63)), F1: v301}
                        _ = v302
                    case 9223372036854775807:
                        switch v41 {
                        case "\n\t":
                            v41 = ((v41 + "\n\t") + "hello")
                        case "":
                            if (v301 || (v301 && (-125 < 0))) {
                                for fi303 := 0; fi303 < 3; fi303++ {
                                    switch (fi303 + (^(fi303))) {
                                    case 32768:
                                        for fi304 := 0; fi304 < 14; fi304++ {
                                            fmt.Println(fi304, v5)
                                            v41 = (v41 + v41)
                                            _ = fi304
                                        }
                                    case -635:
                                        var v305 uint = 882
                                        _ = v305
                                    }
                                    switch fi303 {
                                    case 781:
                                        var v306 float32 = float32(-2.25)
                                        _ = v306
                                    case 42:
                                        var v307 float64 = 1.5
                                        _ = v307
                                    case 899:
                                        var v308 bool = (2 >= 157)
                                        _ = v308
                                    }
                                    _ = fi303
                                }
                            } else {
                                var v309 int16 = -846
                                var v310 string = ("hello" + ((v41 + v41) + v41))
                                v311 := GenS18{F0: 42}
                                _ = v309
                                _ = v310
                                _ = v311
                            }
                        }
                    }
                    v5 += (v5 + (^(v5)))
                    var v312 bool = (!false)
                    _ = v312
                }
                var v313 float64 = -2.25
                _ = v301
                _ = v313
            }
        }
    case 128:
        for fi314 := 0; fi314 < 2; fi314++ {
            var v315 int8 = 42
            v5 = v5
            _ = fi314
            _ = v315
        }
    default:
        v41 = "\n\t"
    }
    v41 = v41
    for fi316 := 0; fi316 < 11; fi316++ {
        if (1 == v5) {
            v317 := v3[2]
            _ = v317
        }
        switch v41 {
        case "\n\t":
            var v318 uint8 = 124
            _ = v318
        default:
            var v319 uint32 = 2147483648
            _ = v319
        }
        if (42 >= 526) {
            v5 -= 42
        } else {
            v41 = v41
            if true {
                if (-32768 != -164) {
                    v320 := GenS30{F0: (("\n\t" != ("\n\t" + "hello")) && false), F1: (18446744073709551615 >= v6)}
                    _ = v320
                } else {
                    var v321 uint64 = v6
                    var v322 uint16 = 633
                    for fi323 := 0; fi323 < 10; fi323++ {
                        if false {
                            delete(v3, 0)
                        }
                        var v324 string = ("hello" + v41)
                        switch "a" {
                        case "\n\t":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "a":
                            if ((v6 & 762) == v321) {
                                var v325 uint64 = uint64(v322)
                                _ = v325
                            } else {
                                v41 = (v324 + ((v41 + ("" + v41)) + (v41 + v324)))
                            }
                        }
                        _ = fi323
                        _ = v324
                    }
                    _ = v321
                    _ = v322
                }
                var v326 int16 = -845
                _ = v326
            } else {
                for fi327 := 0; fi327 < 14; fi327++ {
                    switch fi327 {
                    case 0:
                        v328 := GenS7{F0: ("" + v41)}
                        _ = v328
                    case -9223372036854775808:
                        var v329 float64 = (1.5 - 0.1)
                        _ = v329
                    default:
                        var v330 float64 = ((1e10 - (1e10 - 1.5)) * 1.5)
                        _ = v330
                    }
                    _ = len(v1)
                    if len(v4) > 0 {
                        v4[0] = fi316
                    }
                    _ = fi327
                }
                _ = len(v4)
                for fi331 := 0; fi331 < 10; fi331++ {
                    var v332 uint32 = 65535
                    var v333 int16 = 307
                    _ = fi331
                    _ = v332
                    _ = v333
                }
            }
            for fi334 := 0; fi334 < 13; fi334++ {
                if (127 != 25) {
                    var v335 int32 = 2147483647
                    _ = v335
                }
                v5 = 146
                if (((1e10 / (3.14159 * (0.1 + (-2.25 * 1.5)))) / 1e10) >= 3.14159) {
                    var v336 uint16 = 65535
                    _ = v336
                } else {
                    var v337 int8 = 100
                    v338 := GenS13{F0: 24, F1: 42}
                    _ = v337
                    _ = v338
                }
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                var v339 int8 = 100
                if false {
                    if ((false || true) || true) {
                        v41 = v41
                    } else {
                        fmt.Println(v41, v41)
                        if true {
                            var v340 int = fi334
                            var v341 int32 = v5
                            if ((v41 + v41) == ((v41 + (v41 + ("hello" + ((v41 + v41) + v41)))) + "")) {
                                v41 = v41
                                fmt.Println(fi334)
                            } else {
                                v340 = v340
                                if false {
                                    var v342 float64 = float64(0.1)
                                    fmt.Println(v5, v5, v341)
                                    _ = v342
                                }
                            }
                            _ = v340
                            _ = v341
                        } else {
                            v6 = 42
                            if true {
                                _ = len(v3)
                            } else {
                                var v343 int32 = (^(v5))
                                if (-1 != 542) {
                                    var v344 bool = (131 == 966)
                                    if (!(65535 != 42)) {
                                        _ = len(v3)
                                        var v345 uint32 = 450
                                        v5 -= v5
                                        _ = v345
                                    } else {
                                        if true {
                                            fmt.Println(v5, v343, fi334)
                                            v6 += 4294967295
                                            v346 := v3[2]
                                            _ = v346
                                        } else {
                                            v347 := GenS10{F0: v6, F1: 0}
                                            var v348 int8 = 42
                                            v339 -= (-(v348))
                                            v339 = v348
                                            var v349 int = 2147483647
                                            _ = v347
                                            _ = v348
                                            _ = v349
                                        }
                                        v41 = v41
                                        switch (v41 + v41) {
                                        case "foo bar":
                                            switch (fi334 * fi334) {
                                            case -1:
                                                var v350 float32 = 0.0
                                                _ = v350
                                            case 1:
                                                var v351 int = (fi316 ^ (fi334 * fi316))
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                                _ = v351
                                            case 42:
                                                fmt.Println(v5, v41)
                                            }
                                        case "\n\t":
                                            v352 := v3[42]
                                            _ = v352
                                        case "a":
                                            v353 := GenS25{F0: 18446744073709551615, F1: ((v41 + ("hello" + ("" + ((v41 + "hello") + "a")))) + "hello"), F2: 9223372036854775807}
                                            _ = v353
                                        default:
                                            switch 634 {
                                            case 149:
                                                var v354 int = 42
                                                switch -9223372036854775808 {
                                                case -540:
                                                    if v344 {
                                                        v339 -= -6
                                                    }
                                                }
                                                v354 = v354
                                                fmt.Println(v41, v354)
                                                _ = v354
                                            default:
                                                fmt.Println(v6, v41)
                                                v41 = v41
                                            }
                                        }
                                    }
                                    var v355 uint = 65535
                                    _ = v344
                                    _ = v355
                                } else {
                                    v5 += (v343 | v343)
                                    switch (fi316 << ((971) & 63)) {
                                    case 127:
                                        v356 := GenS31{F0: 1, F1: 9223372036854775807}
                                        _ = v356
                                    case 32767:
                                        var v357 int32 = (v343 ^ (-(v5)))
                                        if ((-2147483648 >= v343) || true) {
                                            var v358 uint8 = 42
                                            var v359 uint8 = v358
                                            var v360 int16 = 1
                                            _ = v358
                                            _ = v359
                                            _ = v360
                                        } else {
                                            v3[510] = (fi316 & fi316)
                                            var v361 float64 = (1e10 + 1.5)
                                            var v362 uint = 42
                                            _ = v361
                                            _ = v362
                                        }
                                        v41 = v41
                                        _ = v357
                                    }
                                    v6 -= v6
                                }
                                fmt.Println(v339, v6, v343)
                                _ = v343
                            }
                        }
                        switch fi316 {
                        case -2147483648:
                            if ((668 != 42) || (true || ((fi334 - fi334) != (fi316 | fi334)))) {
                                delete(v3, 1)
                                v41 = "\n\t"
                                var v363 int32 = v5
                                _ = v363
                            } else {
                                if ((float64(3.14159) / (-2.25 + (3.14159 - float64(float32(-2.25))))) >= 0.1) {
                                    v1 = append(v1, 289)
                                    if (594 != 339) {
                                        if (true || false) {
                                            v339 -= 100
                                        } else {
                                            var v364 int8 = (-(v339))
                                            v339 = v339
                                            _ = v364
                                        }
                                        v365 := GenS29{F0: (v6 - (^(v6))), F1: (!((4294967295 < 145) && false))}
                                        fmt.Println(v6, v339, v41)
                                        _ = v365
                                    }
                                    var v366 int = (fi334 / ((634) | 1))
                                    v41 = v41
                                    fmt.Println(v5, v5, v6)
                                    _ = v366
                                }
                            }
                        case 128:
                            v367 := GenS5{F0: 1, F1: 65535, F2: 65535}
                            _ = v367
                        case 42:
                            v368 := GenS13{F0: 0, F1: 32767}
                            _ = v368
                        }
                    }
                    v41 = v41
                }
                _ = fi334
                _ = v339
            }
        }
        _ = fi316
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
