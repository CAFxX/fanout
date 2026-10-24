package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 int32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int32; F1 uint; F2 int16 }
type GenS6 struct { F0 uint64; F1 int; F2 string }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64; F1 uint16 }
type GenS11 struct { F0 string; F1 uint8; F2 uint64 }
type GenS12 struct { F0 uint16; F1 int8; F2 int8 }
type GenS13 struct { F0 int8; F1 uint32 }

func main() {
    var v1 bool = (281 < 255)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483647
    var v5 string = "\n\t"
    if v1 {
        v5 = (v5 + v5)
    }
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 42, F1: v4, F2: 42}
            var v8 uint16 = 255
            _ = fi6
            _ = v7
            _ = v8
        }
        v4 = (v4 | (v4 ^ 65535))
        if ((v4 << ((867) & 63)) != 65535) {
            var v9 int8 = 42
            _ = v9
        }
    } else {
        if (32767 >= 674) {
            if (1 == 230) {
                if true {
                    var v10 int = -966
                    _ = v10
                }
                v11 := v3[v4]
                _ = v11
            }
            v4 += v4
        }
        var v12 int = 127
        _ = v12
    }
    if v1 {
        v1 = false
    }
    v13 := GenS2{F0: (true && v1), F1: v4, F2: -530}
    fmt.Println(v5, v5)
    switch 9223372036854775807 {
    case 9223372036854775807:
        if (v1 && (!v1)) {
            v14 := GenS3{F0: 127, F1: -2147483648, F2: -702}
            _ = v14
        } else {
            if true {
                switch 0 {
                case 128:
                    var v15 int64 = 42
                    _ = v15
                default:
                    fmt.Println(v5, v5)
                }
                v16 := GenS4{F0: v1, F1: 9223372036854775807}
                _ = v16
            } else {
                v17 := GenS5{F0: -2147483648, F1: 65535, F2: 1000}
                fmt.Println(v1, v1)
                _ = v17
            }
            switch v5 {
            case "hello":
                v18 := GenS4{F0: false, F1: -9223372036854775808}
                _ = v18
            case "":
                switch -2147483648 {
                case 128:
                    v19 := GenS5{F0: -1, F1: 2, F2: 42}
                    _ = v19
                case 2147483647:
                    if (0 == 255) {
                        v4 -= 2
                        var v20 uint = 323
                        for fi21 := 0; fi21 < 2; fi21++ {
                            for fi22 := 0; fi22 < 20; fi22++ {
                                v20 = (v20 * (v20 - v20))
                                _ = fi22
                            }
                            fmt.Println(v5, v5, v5)
                            if (((622 == 0) || v1) && (1 >= 0)) {
                                var v23 int8 = 0
                                _ = v23
                            }
                            _ = fi21
                        }
                        _ = v20
                    } else {
                        var v24 int64 = 42
                        var v25 uint64 = 42
                        _ = v24
                        _ = v25
                    }
                case 1:
                    delete(v3, 916)
                default:
                    var v26 uint = 4294967295
                    _ = v26
                }
            case "\n\t":
                v4 += v4
            }
        }
    case -355:
        switch v5 {
        case "foo bar":
            var v27 int = -871
            _ = v27
        case "hello":
            for fi28 := 0; fi28 < 17; fi28++ {
                v5 = "\n\t"
                if ((^(v4)) <= (v4 + v4)) {
                    v29 := GenS2{F0: v1, F1: v4, F2: -798}
                    _ = v29
                } else {
                    var v30 uint64 = 1
                    for fi31 := 0; fi31 < 20; fi31++ {
                        v30 = v30
                        var v32 float64 = 3.14159
                        if (fi28 > (fi31 * fi28)) {
                            fmt.Println(fi28)
                        } else {
                            var v33 int16 = -32768
                            _ = v33
                        }
                        var v34 uint8 = 92
                        _ = fi31
                        _ = v32
                        _ = v34
                    }
                    _ = v30
                }
                v1 = (v5 == "hello")
                _ = fi28
            }
        case "a":
            if v1 {
                v35 := GenS6{F0: 180, F1: -2147483648, F2: "hello"}
                _ = v35
            }
        }
    }
    switch (v5 + v5) {
    case "foo bar":
        var v36 uint64 = uint64(v4)
        _ = v36
    case "a":
        if v1 {
            _ = len(v3)
            v37 := GenS7{F0: 255, F1: 2147483647}
            v1 = (v1 || (false && false))
            _ = v37
        } else {
            v1 = v1
        }
    default:
        v4 = 2
    }
    switch "hello" {
    case "foo bar":
        for fi38 := 0; fi38 < 9; fi38++ {
            if v1 {
                if false {
                    fmt.Println(v5, fi38, v1)
                    for fi39 := 0; fi39 < 13; fi39++ {
                        var v40 uint32 = (v4 / (((v4 + v4)) | 1))
                        v1 = v1
                        _ = fi39
                        _ = v40
                    }
                } else {
                    v41 := GenS4{F0: (true || v1), F1: (^(fi38))}
                    fmt.Println(v1, v1)
                    _ = v41
                }
                switch v5 {
                case "foo bar":
                    v3[(v4 - v4)] = false
                case "":
                    v1 = v1
                case "hello":
                    switch 2147483647 {
                    case -1:
                        var v42 uint = 0
                        _ = v42
                    default:
                        var v43 string = (((v5 + (v5 + v5)) + "hello") + v5)
                        _ = v43
                    }
                }
                v44 := GenS8{F0: 149}
                _ = v44
            } else {
                switch "a" {
                case "foo bar":
                    v45 := []int{(^(fi38))}
                    _ = v45
                case "\n\t":
                    v46 := GenS4{F0: v1, F1: 9223372036854775807}
                    _ = v46
                case "":
                    var v47 uint64 = 0
                    _ = v47
                }
                var v48 float32 = (-2.25 - 3.14159)
                _ = v48
            }
            v49 := []int{(fi38 * fi38), fi38}
            v50 := GenS8{F0: 42}
            _ = fi38
            _ = v49
            _ = v50
        }
    case "\n\t":
        if v1 {
            fmt.Println(v1, v4)
        } else {
            v51 := []int{32768, 128, -47}
            if v1 {
                if len(v51) > 0 {
                    v51[0] = 1
                }
                v1 = (!(461 == 4294967295))
                v5 = (v5 + v5)
            } else {
                if true {
                    if v1 {
                        for fi52 := 0; fi52 < 12; fi52++ {
                            switch (fi52 - (fi52 & fi52)) {
                            case 2147483647:
                                if ((v1 || ((!(-1 < 127)) || false)) || v1) {
                                    if (-1 > 32767) {
                                        if v1 {
                                            switch (fi52 ^ fi52) {
                                            case -2147483648:
                                                var v53 uint8 = 127
                                                _ = v53
                                            case 1:
                                                var v54 int16 = 1
                                                _ = v54
                                            case -9223372036854775808:
                                                v5 = "\n\t"
                                            }
                                            v4 += v4
                                            v55 := GenS2{F0: false, F1: (v4 / (((^(v4))) | 1)), F2: -9}
                                            _ = v55
                                        } else {
                                            var v56 int16 = -1
                                            var v57 float32 = (1.5 / 0.1)
                                            fmt.Println(v5)
                                            _ = v56
                                            _ = v57
                                        }
                                        v58 := GenS7{F0: 127, F1: -1}
                                        var v59 int8 = 1
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v60 uint16 = 0
                                        if len(v51) > 0 {
                                            v51[0] = (fi52 + fi52)
                                        }
                                        _ = v60
                                    }
                                }
                            case -1:
                                switch 0 {
                                case 9223372036854775807:
                                    for fi61 := 0; fi61 < 1; fi61++ {
                                        if (100 < 58) {
                                            v4 += v4
                                        } else {
                                            var v62 float32 = -2.25
                                            _ = v62
                                        }
                                        var v63 int8 = 42
                                        _ = fi61
                                        _ = v63
                                    }
                                case 0:
                                    switch (fi52 + (fi52 + fi52)) {
                                    case -9223372036854775808:
                                        var v64 int16 = 1
                                        _ = v64
                                    case 42:
                                        if (!false) {
                                            for fi65 := 0; fi65 < 11; fi65++ {
                                                v4 -= 4294967295
                                                fmt.Println(v1, v4)
                                                _ = fi65
                                            }
                                            switch (fi52 / ((42) | 1)) {
                                            case -21:
                                                for fi66 := 0; fi66 < 7; fi66++ {
                                                    v67 := GenS8{F0: 46}
                                                    _ = len(v2)
                                                    _ = fi66
                                                    _ = v67
                                                }
                                            default:
                                                if len(v51) >= 2 {
                                                    v51 = v51[1:len(v51)]
                                                }
                                            }
                                            if ((-9223372036854775808 == 632) && (uint8(v4) != 0)) {
                                                var v68 uint32 = v4
                                                _ = v68
                                            } else {
                                                for fi69 := 0; fi69 < 17; fi69++ {
                                                    v1 = v1
                                                    _ = fi69
                                                }
                                                v4 += (v4 / (((v4 >> ((0) & 63))) | 1))
                                                v5 = ("foo bar" + v5)
                                            }
                                        }
                                    case -521:
                                        var v70 int16 = -812
                                        _ = v70
                                    default:
                                        v71 := GenS7{F0: 119, F1: 0}
                                        _ = v71
                                    }
                                }
                            case 9223372036854775807:
                                var v72 int8 = -128
                                _ = v72
                            default:
                                for fi73 := 0; fi73 < 8; fi73++ {
                                    v5 = "a"
                                    v4 += 0
                                    switch (fi73 + -2147483648) {
                                    case -1:
                                        var v74 bool = ((0.1 + float64(3.14159)) == 0.0)
                                        _ = v74
                                    case 2147483647:
                                        var v75 uint64 = 1
                                        _ = v75
                                    }
                                    if false {
                                        v76 := v3[v4]
                                        var v77 int8 = 100
                                        _ = v76
                                        _ = v77
                                    } else {
                                        var v78 float32 = 3.14159
                                        _ = v78
                                    }
                                    if len(v51) >= 2 {
                                        v51 = v51[1:len(v51)]
                                    }
                                    var v79 int64 = -1
                                    var v80 string = ""
                                    if len(v51) >= 2 {
                                        v51 = v51[1:len(v51)]
                                    }
                                    fmt.Println(fi73)
                                    if v1 {
                                        v79 = (v79 * (-(v79)))
                                    } else {
                                        v81 := GenS3{F0: 127, F1: -2147483648, F2: 2147483647}
                                        _ = v81
                                    }
                                    _ = fi73
                                    _ = v79
                                    _ = v80
                                }
                            }
                            _ = fi52
                        }
                        switch -363 {
                        case -435:
                            var v82 int32 = 42
                            _ = v82
                        case -166:
                            if len(v51) > 0 {
                                v51[0] = -209
                            }
                        }
                        switch "hello" {
                        case "":
                            if false {
                                v4 = v4
                                fmt.Println(v1, v1)
                            } else {
                                switch 4294967295 {
                                case 9223372036854775807:
                                    if v1 {
                                        var v83 float32 = 1e10
                                        fmt.Println(v5)
                                        _ = v83
                                    } else {
                                        v1 = v1
                                    }
                                case -673:
                                    switch -634 {
                                    case -2147483648:
                                        v1 = true
                                    case 756:
                                        v84 := GenS1{F0: 18446744073709551615, F1: v4, F2: 42}
                                        _ = v84
                                    }
                                }
                            }
                        case "hello":
                            if false {
                                fmt.Println(v1)
                            } else {
                                if false {
                                    fmt.Println(v5, v5)
                                    fmt.Println(v5, v4)
                                    if len(v51) >= 2 {
                                        v51 = v51[1:len(v51)]
                                    }
                                } else {
                                    if (!(!true)) {
                                        for fi85 := 0; fi85 < 17; fi85++ {
                                            var v86 uint = 231
                                            v4 -= v4
                                            _ = fi85
                                            _ = v86
                                        }
                                        var v87 uint64 = 0
                                        var v88 int64 = -9223372036854775808
                                        _ = v87
                                        _ = v88
                                    } else {
                                        if (1 >= -853) {
                                            var v89 uint64 = uint64(v4)
                                            delete(v2, 42)
                                            if (4294967295 != -637) {
                                                v5 = ""
                                                for fi90 := 0; fi90 < 2; fi90++ {
                                                    v5 = v5
                                                    _ = fi90
                                                }
                                            }
                                            _ = v89
                                        } else {
                                            for fi91 := 0; fi91 < 5; fi91++ {
                                                var v92 int16 = -1
                                                if v1 {
                                                    if (v1 && (127 == 42)) {
                                                        var v93 float32 = 1.5
                                                        _ = v93
                                                    } else {
                                                        v4 += v4
                                                        var v94 uint8 = 2
                                                        if false {
                                                            var v95 float64 = 0.1
                                                            _ = v95
                                                        }
                                                        _ = v94
                                                    }
                                                    v96 := GenS6{F0: 667, F1: 9223372036854775807, F2: v5}
                                                    _ = v96
                                                } else {
                                                    var v97 uint16 = 2
                                                    var v98 float32 = (3.14159 * (3.14159 + float32((1e10 * (-2.25 / ((float64(0.1) + (1.5 * (0.1 - (0.0 - (float64((0.1 + (3.14159 / (3.14159 * (0.1 - (-2.25 * -2.25)))))) + -2.25))))) / (((1e10 + (0.1 / 1.5)) - 0.1) + (0.0 + -2.25))))))))
                                                    _ = v97
                                                    _ = v98
                                                }
                                                _ = fi91
                                                _ = v92
                                            }
                                        }
                                        v99 := GenS6{F0: 18446744073709551615, F1: -614, F2: v5}
                                        fmt.Println(v1, v5)
                                        _ = v99
                                    }
                                    if v1 {
                                        v1 = false
                                        switch v5 {
                                        case "a":
                                            v4 -= 507
                                        case "foo bar":
                                            v4 = 65535
                                        default:
                                            switch ((v5 + "a") + (v5 + v5)) {
                                            case "":
                                                switch v5 {
                                                case "a":
                                                    v1 = v1
                                                }
                                            case "\n\t":
                                                var v100 float32 = float32(3.14159)
                                                _ = v100
                                            case "a":
                                                for fi101 := 0; fi101 < 18; fi101++ {
                                                    var v102 uint64 = 700
                                                    _ = fi101
                                                    _ = v102
                                                }
                                            }
                                        }
                                        for fi103 := 0; fi103 < 11; fi103++ {
                                            if false {
                                                fmt.Println(v4)
                                                if (!v1) {
                                                    switch (fi103 & (fi103 ^ fi103)) {
                                                    case 4294967295:
                                                        var v104 bool = (!(0 < 463))
                                                        _ = v104
                                                    }
                                                }
                                                var v105 uint32 = v4
                                                _ = v105
                                            } else {
                                                for fi106 := 0; fi106 < 3; fi106++ {
                                                    v107 := GenS9{F0: fi103}
                                                    _ = fi106
                                                    _ = v107
                                                }
                                            }
                                            switch fi103 {
                                            case 89:
                                                var v108 float32 = 1.5
                                                _ = v108
                                            case -2147483648:
                                                var v109 int32 = -93
                                                _ = v109
                                            default:
                                                v110 := GenS3{F0: 127, F1: 2147483647, F2: -2147483648}
                                                _ = v110
                                            }
                                            _ = fi103
                                        }
                                    }
                                    v5 = "hello"
                                }
                                if (128 < 127) {
                                    var v111 int64 = 9223372036854775807
                                    switch -919 {
                                    case -9223372036854775808:
                                        var v112 float64 = ((((0.1 + 1e10) * float64((3.14159 * (-2.25 * float32(((((1e10 + -2.25) * 0.0) + 0.1) + ((float64((-2.25 - 0.1)) - 3.14159) + 1.5))))))) / 1.0) * (float64(3.14159) + 0.0))
                                        _ = v112
                                    case 32768:
                                        fmt.Println(v1, v111, v111)
                                    default:
                                        if len(v51) >= 2 {
                                            v51 = v51[1:len(v51)]
                                        }
                                    }
                                    var v113 bool = (v1 || true)
                                    _ = v111
                                    _ = v113
                                } else {
                                    var v114 int8 = 0
                                    for fi115 := 0; fi115 < 11; fi115++ {
                                        v116 := GenS9{F0: 625}
                                        v4 = v4
                                        _ = fi115
                                        _ = v116
                                    }
                                    if v1 {
                                        switch -588 {
                                        case 32767:
                                            v5 = (v5 + "\n\t")
                                        default:
                                            delete(v2, v4)
                                        }
                                    } else {
                                        if v1 {
                                            switch -201 {
                                            case 127:
                                                var v117 float32 = 0.0
                                                _ = v117
                                            case -747:
                                                if ((float64(3.14159) / 1.5) == (0.1 * 1.5)) {
                                                    for fi118 := 0; fi118 < 11; fi118++ {
                                                        for fi119 := 0; fi119 < 3; fi119++ {
                                                            v1 = v1
                                                            v5 = ""
                                                            _ = fi119
                                                        }
                                                        _ = fi118
                                                    }
                                                    for fi120 := 0; fi120 < 13; fi120++ {
                                                        _ = len(v51)
                                                        _ = len(v51)
                                                        _ = len(v51)
                                                        _ = fi120
                                                    }
                                                }
                                            case 709:
                                                var v121 uint8 = 165
                                                _ = v121
                                            }
                                            var v122 uint64 = 427
                                            v4 += 2147483648
                                            _ = v122
                                        } else {
                                            var v123 int32 = -453
                                            _ = v123
                                        }
                                        var v124 bool = (-2.25 != 1.5)
                                        _ = v124
                                    }
                                    _ = v114
                                }
                                var v125 uint32 = 4294967295
                                _ = v125
                            }
                        case "\n\t":
                            if false {
                                _ = len(v2)
                            }
                        default:
                            v126 := GenS2{F0: (32767 >= 65535), F1: v4, F2: 125}
                            _ = v126
                        }
                    } else {
                        var v127 int = -9223372036854775808
                        _ = v127
                    }
                    for fi128 := 0; fi128 < 6; fi128++ {
                        v4 += 435
                        _ = fi128
                    }
                    if (-772 == -1) {
                        v2[(v4 >> ((1) & 63))] = 4294967295
                    } else {
                        for fi129 := 0; fi129 < 1; fi129++ {
                            v130 := GenS4{F0: v1, F1: (fi129 ^ 32767)}
                            delete(v2, v4)
                            _ = fi129
                            _ = v130
                        }
                        _ = len(v2)
                        switch v5 {
                        case "a":
                            var v131 float64 = (-2.25 / 0.1)
                            _ = v131
                        case "\n\t":
                            var v132 uint16 = 2
                            _ = v132
                        case "hello":
                            for fi133 := 0; fi133 < 18; fi133++ {
                                if v1 {
                                    v4 = (v4 << ((2) & 63))
                                    v5 = v5
                                    if (9223372036854775807 < int64(fi133)) {
                                        v5 = v5
                                    }
                                } else {
                                    v134 := GenS10{F0: 9223372036854775807, F1: 255}
                                    _ = v134
                                }
                                v4 -= 4294967295
                                _ = fi133
                            }
                        }
                    }
                } else {
                    v1 = v1
                }
                v4 = v4
            }
            _ = v51
        }
    default:
        for fi135 := 0; fi135 < 16; fi135++ {
            _ = len(v3)
            switch (fi135 & (fi135 >> ((18446744073709551615) & 63))) {
            case 9223372036854775807:
                fmt.Println(v5)
            case -2147483648:
                fmt.Println(v4, v5)
            case 32767:
                v136 := GenS9{F0: (^(fi135))}
                _ = v136
            default:
                var v137 uint64 = 2
                _ = v137
            }
            _ = fi135
        }
    }
    v138 := GenS11{F0: "\n\t", F1: 127, F2: 42}
    var v139 uint32 = 4294967295
    for fi140 := 0; fi140 < 17; fi140++ {
        for fi141 := 0; fi141 < 6; fi141++ {
            if (-1 >= -34) {
                fmt.Println(fi140, v5, fi140)
            }
            _ = fi141
        }
        _ = fi140
    }
    switch v5 {
    case "":
        switch -1 {
        case 32768:
            for fi142 := 0; fi142 < 1; fi142++ {
                switch v5 {
                case "hello":
                    var v143 int = fi142
                    _ = v143
                case "a":
                    if v1 {
                        fmt.Println(v4, v139)
                    } else {
                        var v144 int32 = -1
                        _ = v144
                    }
                case "foo bar":
                    switch (((v5 + ((v5 + v5) + v5)) + "") + v5) {
                    case "hello":
                        switch fi142 {
                        case -2147483648:
                            var v145 uint8 = 212
                            _ = v145
                        case -9223372036854775808:
                            if (!false) {
                                fmt.Println(v139, v139, v139)
                                var v146 int32 = 385
                                for fi147 := 0; fi147 < 5; fi147++ {
                                    v1 = true
                                    delete(v2, 0)
                                    _ = fi147
                                }
                                _ = v146
                            } else {
                                v4 = v4
                                var v148 float32 = -2.25
                                _ = v148
                            }
                        case -21:
                            if v1 {
                                switch v5 {
                                case "foo bar":
                                    for fi149 := 0; fi149 < 5; fi149++ {
                                        v150 := GenS12{F0: 255, F1: 100, F2: -1}
                                        var v151 string = v5
                                        if false {
                                            var v152 uint8 = 255
                                            v153 := []int{}
                                            _ = v152
                                            _ = v153
                                        } else {
                                            var v154 uint8 = 255
                                            _ = v154
                                        }
                                        _ = fi149
                                        _ = v150
                                        _ = v151
                                    }
                                case "":
                                    if ((v4 == (^(v139))) || false) {
                                        for fi155 := 0; fi155 < 14; fi155++ {
                                            v156 := []int{(^(fi155)), 347}
                                            _ = fi155
                                            _ = v156
                                        }
                                        for fi157 := 0; fi157 < 14; fi157++ {
                                            var v158 int16 = 1
                                            if v1 {
                                                if false {
                                                    var v159 int64 = 9223372036854775807
                                                    if ((42 != 128) && (0 == 1)) {
                                                        v158 += 333
                                                        fmt.Println(v1, v1, v159)
                                                    } else {
                                                        switch -9223372036854775808 {
                                                        case 781:
                                                            var v160 uint16 = 0
                                                            _ = v160
                                                        case -674:
                                                            var v161 int16 = v158
                                                            _ = v161
                                                        case -997:
                                                            v162 := GenS3{F0: 120, F1: 364, F2: 1}
                                                            _ = v162
                                                        }
                                                    }
                                                    _ = v159
                                                }
                                                v163 := GenS1{F0: 347, F1: v139, F2: 18446744073709551615}
                                                fmt.Println(fi157)
                                                _ = v163
                                            }
                                            v164 := []int{}
                                            _ = fi157
                                            _ = v158
                                            _ = v164
                                        }
                                    } else {
                                        v5 = ("hello" + (("\n\t" + v5) + v5))
                                    }
                                case "hello":
                                    v1 = (!(!true))
                                }
                            }
                        }
                    case "":
                        fmt.Println(fi142, v1, v5)
                    default:
                        var v165 uint = 4294967295
                        _ = v165
                    }
                }
                _ = fi142
            }
        case 128:
            if v1 {
                for fi166 := 0; fi166 < 14; fi166++ {
                    v5 = "a"
                    _ = fi166
                }
            }
        case 42:
            var v167 int = -87
            _ = v167
        default:
            for fi168 := 0; fi168 < 8; fi168++ {
                v4 -= (v4 % (((v4 / ((v4) | 1))) | 1))
                v139 -= v139
                _ = fi168
            }
        }
    case "foo bar":
        var v169 float64 = 0.1
        _ = v169
    default:
        switch "foo bar" {
        case "":
            v170 := GenS13{F0: -128, F1: (^(v4))}
            _ = v170
        }
    }
    for fi171 := 0; fi171 < 17; fi171++ {
        var v172 bool = true
        v5 = ((v5 + ("hello" + "\n\t")) + ("foo bar" + (v5 + "\n\t")))
        for fi173 := 0; fi173 < 14; fi173++ {
            v1 = false
            if false {
                if true {
                    var v174 int8 = 0
                    _ = v174
                }
            }
            _ = fi173
        }
        _ = fi171
        _ = v172
    }
    v1 = ((true && (v1 || (v1 && v1))) || v1)
    for fi175 := 0; fi175 < 4; fi175++ {
        for fi176 := 0; fi176 < 13; fi176++ {
            var v177 int16 = 42
            switch v5 {
            case "hello":
                var v178 uint = 18446744073709551615
                _ = v178
            case "\n\t":
                var v179 int64 = 9223372036854775807
                _ = v179
            case "foo bar":
                switch (("hello" + v5) + "a") {
                case "a":
                    var v180 uint32 = v139
                    _ = v180
                }
            default:
                delete(v2, v139)
            }
            v181 := []int{(fi176 ^ fi176), -806}
            _ = fi176
            _ = v177
            _ = v181
        }
        var v182 uint32 = 42
        v5 = "hello"
        _ = fi175
        _ = v182
    }
    var v183 uint16 = 65535
    var v184 string = ("foo bar" + (("foo bar" + v5) + v5))
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("v183:", v183)
    fmt.Println("v184:", v184)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v138
    _ = v139
    _ = v183
    _ = v184
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
