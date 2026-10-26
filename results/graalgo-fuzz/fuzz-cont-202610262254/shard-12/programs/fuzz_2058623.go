package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64; F1 int; F2 int }
type GenS3 struct { F0 uint16; F1 int64; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int; F1 uint64 }
type GenS14 struct { F0 uint16; F1 bool }
type GenS15 struct { F0 bool; F1 uint; F2 int32 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int32; F1 int }
type GenS18 struct { F0 string; F1 int8 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 int8; F1 int8 }
type GenS21 struct { F0 uint8; F1 int }
type GenS22 struct { F0 int64 }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]string)
    var v3 uint16 = 255
    v4 := make(map[int]string)
    var v5 float64 = float64((float32(float64(3.14159)) / 1.5))
    var v6 float32 = (3.14159 + float32((float64(1e10) * v5)))
    var v7 uint = 4294967295
    v1 = (v1 & (v1 * v1))
    v7 += v7
    var v8 uint8 = 0
    v9 := []int{1, 128}
    if (!("\n\t" >= ("hello" + "foo bar"))) {
        switch "" {
        case "foo bar":
            v8 += (v8 - (v8 & (v8 / (((v8 & (v8 & 2))) | 1))))
        case "\n\t":
            for fi10 := 0; fi10 < 3; fi10++ {
                switch 460 {
                case 127:
                    v8 = v8
                case 1:
                    fmt.Println(v3)
                default:
                    var v11 bool = (-128 == 0)
                    _ = v11
                }
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := GenS1{F0: true, F1: (fi12 >> ((v7) & 63))}
                    v7 -= 336
                    var v14 int8 = 102
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                switch int(v8) {
                case -404:
                    switch "" {
                    case "":
                        v7 = 255
                    default:
                        var v15 int8 = 127
                        _ = v15
                    }
                case -2147483648:
                    var v16 int = (fi10 | fi10)
                    _ = v16
                case -435:
                    if false {
                        if ("a" != "\n\t") {
                            var v17 bool = false
                            switch "foo bar" {
                            case "\n\t":
                                for fi18 := 0; fi18 < 7; fi18++ {
                                    var v19 uint8 = (^(v8))
                                    v7 -= 679
                                    _ = fi18
                                    _ = v19
                                }
                            case "hello":
                                switch "foo bar" {
                                case "":
                                    if false {
                                        var v20 string = ("\n\t" + "\n\t")
                                        _ = v20
                                    }
                                case "a":
                                    v21 := v4[int(v1)]
                                    _ = v21
                                default:
                                    if (int(v7) >= (-(fi10))) {
                                        if false {
                                            if v17 {
                                                var v22 int32 = -475
                                                _ = len(v9)
                                                var v23 float32 = v6
                                                _ = v22
                                                _ = v23
                                            }
                                            switch (-(fi10)) {
                                            case 127:
                                                v7 += (v7 * 0)
                                            }
                                            v5 = (3.14159 + (1.5 + v5))
                                        } else {
                                            if false {
                                                var v24 int64 = v1
                                                if false {
                                                    v5 = float64(v6)
                                                    _ = len(v9)
                                                    var v25 float64 = v5
                                                    _ = v25
                                                }
                                                v26 := v4[fi10]
                                                _ = v24
                                                _ = v26
                                            }
                                            var v27 uint16 = v3
                                            v28 := GenS2{F0: 18446744073709551615, F1: (fi10 % ((-520) | 1)), F2: (fi10 - fi10)}
                                            _ = v27
                                            _ = v28
                                        }
                                        switch fi10 {
                                        case -1:
                                            v29 := GenS3{F0: v3, F1: (v1 + (v1 / ((930) | 1))), F2: -128}
                                            _ = v29
                                        default:
                                            v30 := GenS2{F0: 2, F1: fi10, F2: (-(fi10))}
                                            _ = v30
                                        }
                                        var v31 int32 = 42
                                        _ = v31
                                    } else {
                                        var v32 float32 = v6
                                        var v33 string = ("" + "")
                                        var v34 float64 = float64(-2.25)
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            default:
                                switch fi10 {
                                case 810:
                                    switch ("" + (("foo bar" + (("\n\t" + "") + "a")) + "\n\t")) {
                                    case "hello":
                                        v35 := GenS4{F0: 485, F1: 128, F2: int64(v3)}
                                        _ = v35
                                    case "\n\t":
                                        var v36 int16 = 1
                                        _ = v36
                                    default:
                                        if ((("a" + "a") > ("hello" + ("a" + ("\n\t" + "foo bar")))) && v17) {
                                            if (1 != 134) {
                                                var v37 bool = true
                                                _ = v37
                                            }
                                            fmt.Println(v1, fi10, v7)
                                        }
                                    }
                                }
                            }
                            if ((v3 << ((v7) & 63)) != v3) {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                            _ = v17
                        } else {
                            for fi38 := 0; fi38 < 5; fi38++ {
                                var v39 uint32 = 2147483648
                                v40 := v2[(v39 ^ (v39 >> ((v7) & 63)))]
                                v41 := GenS5{F0: (^(v7))}
                                var v42 uint = (v7 - 325)
                                if ((1 > int16(v8)) || (!false)) {
                                    _ = len(v2)
                                    var v43 uint16 = 65535
                                    _ = v43
                                } else {
                                    var v44 int16 = 42
                                    _ = v44
                                }
                                _ = fi38
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            switch "foo bar" {
                            case "":
                                var v45 int8 = 0
                                _ = v45
                            case "hello":
                                switch "a" {
                                case "hello":
                                    if true {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v46 := GenS6{F0: 9223372036854775807}
                                            _ = v46
                                        case "hello":
                                            v7 = v7
                                        }
                                        for fi47 := 0; fi47 < 1; fi47++ {
                                            v8 = 246
                                            if false {
                                                fmt.Println(v8, v3)
                                                v48 := GenS7{F0: v8}
                                                if ((0 > int16(fi10)) || (294 < 0)) {
                                                    v6 += 3.14159
                                                }
                                                _ = v48
                                            } else {
                                                var v49 uint64 = 998
                                                _ = v49
                                            }
                                            var v50 int32 = 65
                                            _ = fi47
                                            _ = v50
                                        }
                                    } else {
                                        for fi51 := 0; fi51 < 15; fi51++ {
                                            v52 := GenS1{F0: (32768 > (fi51 & (fi51 << ((v7) & 63)))), F1: fi10}
                                            fmt.Println(v7, v5)
                                            _ = fi51
                                            _ = v52
                                        }
                                        if (!false) {
                                            for fi53 := 0; fi53 < 4; fi53++ {
                                                v8 += (v8 / ((42) | 1))
                                                var v54 int16 = int16(v3)
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                _ = fi53
                                                _ = v54
                                            }
                                            v6 = (-2.25 + v6)
                                        }
                                        v5 += float64(v6)
                                    }
                                case "":
                                    for fi55 := 0; fi55 < 3; fi55++ {
                                        switch "hello" {
                                        case "foo bar":
                                            _ = len(v9)
                                        case "":
                                            if false {
                                                v56 := GenS5{F0: v7}
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "":
                                                        if false {
                                                            switch (fi10 & fi10) {
                                                            case 354:
                                                                if (!false) {
                                                                    v9 = append(v9, (^(fi10)))
                                                                    var v57 bool = true
                                                                    v5 -= (v5 + v5)
                                                                    var v58 uint8 = (v8 * (v8 >> ((v7) & 63)))
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    v59 := GenS8{F0: 622}
                                                                    v4[fi55] = "foo bar"
                                                                    _ = v59
                                                                }
                                                            case 575:
                                                                v7 = 413
                                                            }
                                                            var v60 int8 = 1
                                                            var v61 uint64 = 0
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            if false {
                                                                v62 := GenS9{F0: 0}
                                                                v63 := GenS2{F0: 0, F1: (-(fi10)), F2: fi10}
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                            fmt.Println(v8, v1, v3)
                                                        }
                                                    case "hello":
                                                        if false {
                                                            if (0 < 1) {
                                                                v64 := GenS9{F0: 803}
                                                                v6 -= (((float32(0.0) * v6) / (-2.25 * v6)) + v6)
                                                                _ = v64
                                                            }
                                                            if ((2 > 414) && (((122 >= -1) || true) || (1 == 127))) {
                                                                fmt.Println(v3, v1)
                                                                _ = len(v9)
                                                                v9 = append(v9, fi10)
                                                            }
                                                            if (205 >= (^(v3))) {
                                                                v4[fi10] = "a"
                                                                _ = len(v4)
                                                                v6 = 3.14159
                                                                v3 = 65535
                                                            } else {
                                                                var v65 uint64 = 764
                                                                _ = v65
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v5)
                                                    }
                                                case "":
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                case "hello":
                                                    v3 = 2
                                                default:
                                                    var v66 float64 = (3.14159 * v5)
                                                    _ = v66
                                                }
                                                _ = v56
                                            } else {
                                                if ((fi10 ^ fi10) > (fi55 % ((42) | 1))) {
                                                    v67 := GenS7{F0: v8}
                                                    v5 = 0.1
                                                    if (v7 == (v7 & v7)) {
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        if (v3 > (v3 - (^(v3)))) {
                                                            v8 -= (v8 | v8)
                                                            v9 = append(v9, (fi10 + (fi10 | 1)))
                                                        }
                                                        v7 += v7
                                                        if (127 != -35) {
                                                            var v68 int64 = v1
                                                            if false {
                                                                v69 := GenS7{F0: v8}
                                                                _ = v69
                                                            } else {
                                                                v7 += 18446744073709551615
                                                                if false {
                                                                    v70 := GenS10{F0: 32767, F1: 1}
                                                                    _ = len(v2)
                                                                    if (true || true) {
                                                                        var v71 int = fi55
                                                                        var v72 uint16 = (v3 >> ((0) & 63))
                                                                        var v73 bool = true
                                                                        _ = v71
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    _ = v70
                                                                }
                                                                v9 = append(v9, fi55)
                                                            }
                                                            _ = v68
                                                        } else {
                                                            var v74 uint32 = 2147483647
                                                            _ = v74
                                                        }
                                                        var v75 int16 = -32768
                                                        _ = v75
                                                    } else {
                                                        v76 := GenS3{F0: v3, F1: (^(v1)), F2: 57}
                                                        if false {
                                                            v77 := GenS7{F0: (v8 + 128)}
                                                            _ = len(v2)
                                                            _ = v77
                                                        } else {
                                                            v7 = (v7 % ((v7) | 1))
                                                            if len(v9) > 0 {
                                                                v9[0] = fi10
                                                            }
                                                        }
                                                        _ = v76
                                                    }
                                                    _ = v67
                                                }
                                            }
                                        default:
                                            v3 -= (v3 | 255)
                                        }
                                        _ = fi55
                                    }
                                case "a":
                                    if false {
                                        var v78 int8 = -23
                                        _ = v78
                                    }
                                }
                            default:
                                fmt.Println(v5)
                            }
                        }
                        if true {
                            if true {
                                switch "a" {
                                case "\n\t":
                                    var v79 int32 = 0
                                    _ = v79
                                default:
                                    fmt.Println(v3)
                                }
                                var v80 int32 = 42
                                for fi81 := 0; fi81 < 15; fi81++ {
                                    fmt.Println(v8, v8)
                                    _ = fi81
                                }
                                _ = v80
                            }
                            if ("hello" <= ("a" + (("a" + "hello") + ("" + "foo bar")))) {
                                for fi82 := 0; fi82 < 15; fi82++ {
                                    if ("" != "\n\t") {
                                        if true {
                                            v1 = 42
                                            var v83 float64 = (v5 + 1e10)
                                            v7 = (v7 | v7)
                                            fmt.Println(v3, v5)
                                            v9 = append(v9, (fi82 ^ (fi10 % ((fi10) | 1))))
                                            _ = v83
                                        }
                                    }
                                    var v84 uint = v7
                                    var v85 int64 = 947
                                    var v86 float32 = (((0.1 + v6) + (v6 * (float32(-2.25) / (v6 / v6)))) / ((0.0 + (v6 * v6)) + v6))
                                    var v87 uint8 = uint8(v7)
                                    _ = fi82
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                }
                            } else {
                                for fi88 := 0; fi88 < 10; fi88++ {
                                    switch 950 {
                                    case 127:
                                        var v89 uint = (^(v7))
                                        _ = v89
                                    default:
                                        v90 := GenS10{F0: -1, F1: 1}
                                        _ = v90
                                    }
                                    v7 = v7
                                    switch "hello" {
                                    case "a":
                                        switch "" {
                                        case "":
                                            var v91 uint8 = 127
                                            _ = v91
                                        }
                                    case "":
                                        v7 = v7
                                    }
                                    _ = fi88
                                }
                            }
                        } else {
                            switch -2147483648 {
                            case 0:
                                if false {
                                    v3 = v3
                                    var v92 uint = (^(v7))
                                    if ((false || true) && (720 <= 0)) {
                                        v93 := GenS5{F0: (v92 << ((v7) & 63))}
                                        _ = v93
                                    } else {
                                        v94 := GenS11{F0: v1, F1: (v1 ^ v1)}
                                        _ = v94
                                    }
                                    _ = v92
                                } else {
                                    delete(v4, -658)
                                    v1 = (v1 << ((v7) & 63))
                                    fmt.Println(v8)
                                }
                            case 127:
                                var v95 int16 = 32767
                                _ = v95
                            case -77:
                                for fi96 := 0; fi96 < 10; fi96++ {
                                    switch (("foo bar" + "foo bar") + "hello") {
                                    case "a":
                                        v97 := GenS8{F0: 2}
                                        _ = v97
                                    case "\n\t":
                                        switch fi96 {
                                        case 0:
                                            switch -948 {
                                            case 42:
                                                if len(v9) > 0 {
                                                    v9[0] = (fi96 ^ fi96)
                                                }
                                            case 2147483647:
                                                _ = len(v9)
                                            case 32768:
                                                var v98 string = ""
                                                _ = v98
                                            default:
                                                _ = len(v9)
                                            }
                                        case -1:
                                            v1 += -9223372036854775808
                                        case 9223372036854775807:
                                            if true {
                                                var v99 float64 = float64(v6)
                                                switch ("a" + "hello") {
                                                case "hello":
                                                    var v100 int64 = (v1 + -1)
                                                    var v101 uint16 = v3
                                                    if (42 <= 993) {
                                                        v102 := GenS12{F0: 652}
                                                        _ = v102
                                                    } else {
                                                        if false {
                                                            var v103 float64 = v5
                                                            switch fi10 {
                                                            case 460:
                                                                fmt.Println(v3, fi96, v100)
                                                            case 32768:
                                                                v5 += ((v103 / v103) * v5)
                                                            default:
                                                                var v104 int64 = (v1 | (v100 | (v100 / ((-384) | 1))))
                                                                fmt.Println(v8)
                                                                if false {
                                                                    v103 -= float64(v6)
                                                                    if (-21 < 0) {
                                                                        var v105 int32 = 605
                                                                        v106 := GenS13{F0: (^(fi96)), F1: 128}
                                                                        _ = v105
                                                                        _ = v106
                                                                    }
                                                                    v104 = (v104 % ((v100) | 1))
                                                                }
                                                                _ = v104
                                                            }
                                                            _ = v103
                                                        } else {
                                                            _ = len(v2)
                                                            var v107 bool = true
                                                            switch 9223372036854775807 {
                                                            case 127:
                                                                delete(v4, fi10)
                                                            case -9223372036854775808:
                                                                switch -1 {
                                                                case 635:
                                                                    if v107 {
                                                                        var v108 float64 = (((v99 - (v99 / 1.5)) * v5) / (((v99 - v99) + float64(float32(0.1))) * float64(v6)))
                                                                        fmt.Println(v99, fi10, v6)
                                                                        var v109 int32 = 42
                                                                        _ = v108
                                                                        _ = v109
                                                                    }
                                                                    if len(v9) > 0 {
                                                                        v9[0] = fi10
                                                                    }
                                                                case 32767:
                                                                    if v107 {
                                                                        if v107 {
                                                                            if len(v9) >= 2 {
                                                                                v9 = v9[1:len(v9)]
                                                                            }
                                                                            v110 := GenS14{F0: (v101 & 65535), F1: v107}
                                                                            switch ("\n\t" + "hello") {
                                                                            case "\n\t":
                                                                                v111 := GenS14{F0: v101, F1: (!(false && true))}
                                                                                _ = v111
                                                                            case "hello":
                                                                                switch fi10 {
                                                                                case 400:
                                                                                    v5 += (-2.25 - v5)
                                                                                }
                                                                            }
                                                                            _ = v110
                                                                        } else {
                                                                            if v107 {
                                                                                v112 := v2[uint32(v7)]
                                                                                switch v112 {
                                                                                case "\n\t":
                                                                                    var v113 uint64 = 4294967295
                                                                                    _ = v113
                                                                                case "a":
                                                                                    var v114 string = ((v112 + v112) + v112)
                                                                                    _ = v114
                                                                                case "hello":
                                                                                    v2[414] = v112
                                                                                    v115 := GenS1{F0: v107, F1: (fi10 / (((^(fi10))) | 1))}
                                                                                    fmt.Println(v99)
                                                                                    v116 := GenS5{F0: 65535}
                                                                                    fmt.Println(v7, v1)
                                                                                    if v107 {
                                                                                        v7 = v7
                                                                                        v117 := GenS15{F0: v107, F1: v7, F2: 2147483647}
                                                                                        v3 = v3
                                                                                        switch ("\n\t" + ("" + ("foo bar" + v112))) {
                                                                                        case "hello":
                                                                                            var v118 float64 = v5
                                                                                            _ = v118
                                                                                        case "":
                                                                                            var v119 int8 = -23
                                                                                            _ = v119
                                                                                        case "\n\t":
                                                                                            if true {
                                                                                                var v120 uint32 = 42
                                                                                                var v121 float32 = (v6 + v6)
                                                                                                v7 = (v7 * 0)
                                                                                                if len(v9) >= 2 {
                                                                                                    v9 = v9[1:len(v9)]
                                                                                                }
                                                                                                _ = len(v2)
                                                                                                if v107 {
                                                                                                    var v122 float64 = float64(0.1)
                                                                                                    v9 = append(v9, (fi10 / ((fi96) | 1)))
                                                                                                    if (fi10 < fi10) {
                                                                                                        v112 = ("foo bar" + v112)
                                                                                                        fmt.Println(v112)
                                                                                                    } else {
                                                                                                        v123 := GenS7{F0: 2}
                                                                                                        v124 := GenS3{F0: v3, F1: v1, F2: 39}
                                                                                                        v122 = v99
                                                                                                        v2[(v120 ^ v120)] = (v112 + v112)
                                                                                                        _ = v123
                                                                                                        _ = v124
                                                                                                    }
                                                                                                    _ = len(v9)
                                                                                                    _ = v122
                                                                                                } else {
                                                                                                    var v125 uint64 = 4294967295
                                                                                                    v9 = append(v9, 42)
                                                                                                    _ = v125
                                                                                                }
                                                                                                var v126 int16 = 1000
                                                                                                if (v8 >= 37) {
                                                                                                    fmt.Println(fi96, v121)
                                                                                                }
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                                _ = v126
                                                                                            }
                                                                                            fmt.Println(v6, fi10)
                                                                                        }
                                                                                        if (2 >= 0) {
                                                                                            if len(v9) > 0 {
                                                                                                v9[0] = (fi96 & (fi10 - (fi96 % (((fi96 % ((fi96) | 1))) | 1))))
                                                                                            }
                                                                                            v100 = v1
                                                                                            fmt.Println(v112, v107, v8)
                                                                                        } else {
                                                                                            switch fi10 {
                                                                                            case 32767:
                                                                                                fmt.Println(fi10, v1)
                                                                                            }
                                                                                            v100 = -9223372036854775808
                                                                                        }
                                                                                        _ = v117
                                                                                    } else {
                                                                                        fmt.Println(fi10, v7, v8)
                                                                                    }
                                                                                    _ = v115
                                                                                    _ = v116
                                                                                default:
                                                                                    var v127 string = ((v112 + ((("foo bar" + v112) + v112) + "hello")) + "a")
                                                                                    _ = v127
                                                                                }
                                                                                _ = v112
                                                                            }
                                                                        }
                                                                        var v128 int8 = -128
                                                                        if (v7 >= (v7 >> (((v7 << ((v7) & 63))) & 63))) {
                                                                            v101 -= v101
                                                                        } else {
                                                                            if (!(-122 < 100)) {
                                                                                switch ("foo bar" + ("a" + ("\n\t" + ("hello" + ((("\n\t" + "a") + ("foo bar" + "a")) + ("a" + "")))))) {
                                                                                case "foo bar":
                                                                                    v129 := GenS16{F0: ((("a" + "a") + ("hello" + "hello")) + "hello")}
                                                                                    _ = v129
                                                                                }
                                                                            }
                                                                            _ = len(v2)
                                                                        }
                                                                        _ = v128
                                                                    }
                                                                case 1:
                                                                    var v130 uint32 = 375
                                                                    _ = v130
                                                                }
                                                            case 4294967295:
                                                                var v131 string = ("" + ("" + ""))
                                                                if (42 == -128) {
                                                                    var v132 uint32 = 4294967295
                                                                    _ = v132
                                                                } else {
                                                                    v133 := GenS17{F0: 2147483647, F1: fi10}
                                                                    _ = v133
                                                                }
                                                                fmt.Println(v6, v101, v1)
                                                                _ = v131
                                                            default:
                                                                var v134 int32 = 42
                                                                _ = v134
                                                            }
                                                            _ = v107
                                                        }
                                                        v101 = v3
                                                    }
                                                    if (v1 >= -9223372036854775808) {
                                                        fmt.Println(v3, v99, fi10)
                                                        v135 := v4[int(v8)]
                                                        v7 = v7
                                                        if (v7 < v7) {
                                                            v136 := GenS18{F0: "a", F1: 1}
                                                            v3 = v3
                                                            switch v135 {
                                                            case "":
                                                                if len(v9) > 0 {
                                                                    v9[0] = (fi10 % ((fi96) | 1))
                                                                }
                                                            }
                                                            var v137 int16 = -1
                                                            fmt.Println(v6, v135, fi10)
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                        v3 = (v3 / ((v3) | 1))
                                                        _ = v135
                                                    } else {
                                                        if (true || true) {
                                                            v7 += (^(v7))
                                                            v99 = 3.14159
                                                            var v138 int64 = v100
                                                            _ = v138
                                                        } else {
                                                            v139 := v4[fi96]
                                                            var v140 int16 = -32768
                                                            v101 = v3
                                                            if len(v9) >= 2 {
                                                                v9 = v9[1:len(v9)]
                                                            }
                                                            _ = v139
                                                            _ = v140
                                                        }
                                                        var v141 int32 = -1
                                                        var v142 int8 = 127
                                                        _ = v141
                                                        _ = v142
                                                    }
                                                    var v143 float32 = v6
                                                    v144 := v2[4294967295]
                                                    delete(v4, (fi10 ^ -2147483648))
                                                    _ = v100
                                                    _ = v101
                                                    _ = v143
                                                    _ = v144
                                                default:
                                                    if true {
                                                        if false {
                                                            if ((v7 - v7) >= (v7 >> (((v7 - (v7 >> (((v7 + 18446744073709551615)) & 63)))) & 63))) {
                                                                _ = len(v9)
                                                                fmt.Println(v6, v6)
                                                            } else {
                                                                v99 += 1.5
                                                                switch fi96 {
                                                                case 389:
                                                                    v145 := GenS13{F0: int(v7), F1: 18446744073709551615}
                                                                    _ = v145
                                                                }
                                                                var v146 float32 = 3.14159
                                                                switch "foo bar" {
                                                                case "":
                                                                    var v147 float32 = (v146 / v146)
                                                                    _ = v147
                                                                case "\n\t":
                                                                    v148 := GenS12{F0: 2147483647}
                                                                    _ = v148
                                                                }
                                                                _ = v146
                                                            }
                                                            v7 -= 42
                                                            v149 := GenS9{F0: 0}
                                                            _ = v149
                                                        }
                                                    } else {
                                                        switch 0 {
                                                        case 42:
                                                            v7 = (v7 >> ((uint(v8)) & 63))
                                                        default:
                                                            var v150 uint = (^(v7))
                                                            _ = v150
                                                        }
                                                        var v151 int64 = (v1 | (v1 & v1))
                                                        _ = v151
                                                    }
                                                }
                                                var v152 uint16 = (v3 << (((v7 + (v7 % (((v7 / ((v7) | 1))) | 1)))) & 63))
                                                var v153 bool = (v5 > v5)
                                                var v154 string = (("" + "hello") + "")
                                                v5 += (v99 / v99)
                                                v152 = (v3 ^ 17)
                                                v2[4294967295] = v154
                                                _ = v152
                                                _ = v153
                                                _ = v154
                                                _ = v99
                                            }
                                        }
                                    case "":
                                        switch ("" + ("\n\t" + (("foo bar" + "foo bar") + "foo bar"))) {
                                        case "foo bar":
                                            v6 += v6
                                        default:
                                            if ((v3 ^ v3) >= (v3 & v3)) {
                                                v6 = 1.5
                                                var v155 uint8 = (v8 & 42)
                                                _ = v155
                                            }
                                        }
                                    }
                                    v156 := GenS15{F0: false, F1: 4294967295, F2: -2147483648}
                                    var v157 float32 = ((float32(1.5) - 0.1) / v6)
                                    v4[(fi10 << ((679) & 63))] = ("" + "hello")
                                    _ = fi96
                                    _ = v156
                                    _ = v157
                                }
                            default:
                                v9 = append(v9, fi10)
                            }
                            delete(v2, 600)
                            var v158 int = -771
                            _ = v158
                        }
                        for fi159 := 0; fi159 < 12; fi159++ {
                            v3 += v3
                            _ = fi159
                        }
                    }
                default:
                    if len(v9) > 0 {
                        v9[0] = (fi10 / (((fi10 << (((v7 % ((4294967295) | 1))) & 63))) | 1))
                    }
                }
                _ = fi10
            }
        default:
            if (42 != -263) {
                switch -1 {
                case 134:
                    fmt.Println(v6)
                }
                var v160 uint64 = 104
                var v161 int = 32767
                _ = v160
                _ = v161
            } else {
                if (v1 <= v1) {
                    v7 += v7
                    v7 = v7
                } else {
                    if ((0.0 * v6) < (1.5 + 0.1)) {
                        if (!false) {
                            v162 := GenS18{F0: "a", F1: -1}
                            switch 128 {
                            case 493:
                                v163 := GenS19{F0: -1}
                                _ = v163
                            default:
                                var v164 uint16 = v3
                                _ = v164
                            }
                            fmt.Println(v3, v5)
                            _ = v162
                        }
                        var v165 int32 = 0
                        v5 = v5
                        _ = v165
                    } else {
                        v8 = (^(v8))
                    }
                    var v166 uint = v7
                    for fi167 := 0; fi167 < 16; fi167++ {
                        v168 := GenS4{F0: 2, F1: v8, F2: v1}
                        if true {
                            if true {
                                v169 := GenS2{F0: uint64(v3), F1: (fi167 & fi167), F2: -178}
                                _ = v169
                            }
                            if true {
                                for fi170 := 0; fi170 < 9; fi170++ {
                                    v1 = 0
                                    var v171 bool = (42 < -1)
                                    _ = fi170
                                    _ = v171
                                }
                            }
                        } else {
                            if (v166 != (v166 << ((v7) & 63))) {
                                v172 := GenS20{F0: -102, F1: -128}
                                _ = v172
                            }
                            fmt.Println(v3)
                        }
                        if (!(42 < int8(v7))) {
                            for fi173 := 0; fi173 < 6; fi173++ {
                                v6 = (0.0 + v6)
                                v3 = (^(v3))
                                var v174 float32 = (((1.5 / ((v6 + 1.5) / ((v6 - v6) * v6))) - (((v6 / (1e10 + 0.0)) / v6) + v6)) / 1e10)
                                _ = fi173
                                _ = v174
                            }
                        }
                        _ = fi167
                        _ = v168
                    }
                    _ = v166
                }
                var v175 int64 = v1
                v1 = (v175 / (((v1 & v1)) | 1))
                _ = v175
            }
        }
        var v176 float32 = v6
        v176 += (0.1 + 0.1)
        _ = v176
    } else {
        for fi177 := 0; fi177 < 14; fi177++ {
            if (false || (v8 > v8)) {
                v178 := GenS8{F0: uint64(v3)}
                _ = v178
            } else {
                var v179 uint8 = (v8 ^ (v8 + v8))
                if (int16(fi177) >= -15) {
                    v3 -= v3
                    for fi180 := 0; fi180 < 7; fi180++ {
                        if (false && ((fi180 ^ fi177) <= fi180)) {
                            v4[(fi177 ^ 1)] = ""
                            if false {
                                var v181 string = ""
                                v3 = v3
                                var v182 int = 128
                                _ = v181
                                _ = v182
                            }
                            if (2 == (v7 & 648)) {
                                v7 = (v7 | (v7 * (v7 * v7)))
                                if (-2147483648 >= 2147483647) {
                                    var v183 int8 = 83
                                    var v184 string = "\n\t"
                                    _ = v183
                                    _ = v184
                                }
                            } else {
                                fmt.Println(v1, v6, v8)
                            }
                        }
                        var v185 uint64 = 2
                        if false {
                            v186 := GenS19{F0: 100}
                            v187 := GenS17{F0: -1, F1: (fi180 * (fi177 & fi177))}
                            _ = v186
                            _ = v187
                        }
                        _ = fi180
                        _ = v185
                    }
                } else {
                    if (!true) {
                        if false {
                            if (4294967295 > 63) {
                                v3 = (^(v3))
                                for fi188 := 0; fi188 < 8; fi188++ {
                                    switch 32767 {
                                    case 4294967295:
                                        var v189 int64 = (v1 - -9223372036854775808)
                                        _ = v189
                                    case -2147483648:
                                        v4[fi177] = ""
                                    default:
                                        v190 := GenS8{F0: 18446744073709551615}
                                        _ = v190
                                    }
                                    fmt.Println(v3)
                                    if (-347 != 245) {
                                        var v191 int64 = 9223372036854775807
                                        _ = v191
                                    }
                                    v6 = v6
                                    v7 = 912
                                    var v192 bool = true
                                    _ = fi188
                                    _ = v192
                                }
                                v8 -= v8
                            }
                            if ((v179 | v8) <= v179) {
                                if len(v9) > 0 {
                                    v9[0] = fi177
                                }
                            }
                        }
                        var v193 float32 = (0.1 - v6)
                        _ = v193
                    }
                    v7 += v7
                }
                v194 := GenS21{F0: 0, F1: (fi177 ^ fi177)}
                _ = v179
                _ = v194
            }
            _ = fi177
        }
        v195 := GenS22{F0: (^(v1))}
        var v196 uint64 = uint64(v3)
        _ = v195
        _ = v196
    }
    if true {
        if (-109 != 1) {
            for fi197 := 0; fi197 < 4; fi197++ {
                v3 -= v3
                _ = fi197
            }
            fmt.Println(v5)
            v7 = uint(v8)
        } else {
            v2[30] = "\n\t"
            if ("hello" <= "\n\t") {
                v2[273] = "a"
                var v198 uint16 = (v3 - v3)
                _ = v198
            } else {
                _ = len(v9)
            }
            v6 = v6
        }
        switch ("" + "a") {
        case "hello":
            if false {
                v199 := GenS1{F0: (v3 <= v3), F1: 4294967295}
                _ = v199
            }
        case "a":
            fmt.Println(v3, v7, v5)
        case "foo bar":
            v7 = (v7 - v7)
        }
        var v200 int64 = (v1 - 9223372036854775807)
        _ = v200
    }
    for fi201 := 0; fi201 < 2; fi201++ {
        v5 += (v5 + v5)
        _ = fi201
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
