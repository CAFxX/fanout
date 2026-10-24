package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 string; F2 uint }
type GenS4 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 string; F1 int64; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 int; F2 uint }
type GenS9 struct { F0 int32; F1 uint64; F2 int64 }
type GenS10 struct { F0 int; F1 int32; F2 uint64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string; F1 uint16; F2 int }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 string; F1 int; F2 uint32 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int8; F1 int; F2 int }
type GenS19 struct { F0 uint32; F1 string }
type GenS20 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS21 struct { F0 int }

func main() {
    var v1 uint = 1
    var v2 uint16 = 42
    var v3 string = "hello"
    v1 -= 4294967295
    if true {
        switch "" {
        case "foo bar":
            switch (v3 + (v3 + "foo bar")) {
            case "a":
                switch 127 {
                case -9223372036854775808:
                    if false {
                        var v4 uint = (v1 & (v1 - 4294967295))
                        if (!((!true) && true)) {
                            v5 := []int{-2147483648, 9223372036854775807, 2147483647}
                            _ = v5
                        }
                        _ = v4
                    } else {
                        var v6 bool = (!(!(!true)))
                        for fi7 := 0; fi7 < 6; fi7++ {
                            var v8 int8 = 1
                            switch (v3 + (v3 + v3)) {
                            case "a":
                                v9 := GenS1{F0: (fi7 / ((-2147483648) | 1)), F1: 42, F2: -1}
                                _ = v9
                            case "\n\t":
                                switch v3 {
                                case "\n\t":
                                    v6 = false
                                case "hello":
                                    var v10 float32 = -2.25
                                    _ = v10
                                case "":
                                    switch fi7 {
                                    case 42:
                                        v2 -= v2
                                    case 4294967295:
                                        if v6 {
                                            for fi11 := 0; fi11 < 9; fi11++ {
                                                if (42 < 1) {
                                                    var v12 uint = 42
                                                    var v13 uint8 = 0
                                                    var v14 uint16 = v2
                                                    _ = v12
                                                    _ = v13
                                                    _ = v14
                                                } else {
                                                    v15 := []int{(fi11 ^ (^(fi7))), (-(fi11))}
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v16 uint32 = 2147483648
                                                        _ = v16
                                                    }
                                                    v3 = v3
                                                    var v17 uint8 = 127
                                                    switch fi11 {
                                                    case 32767:
                                                        var v18 uint16 = (v2 & v2)
                                                        v1 = 2
                                                        _ = v18
                                                    case 907:
                                                        v17 -= (v17 / ((104) | 1))
                                                    case -2147483648:
                                                        var v19 uint32 = 2
                                                        _ = v19
                                                    default:
                                                        v20 := GenS2{F0: -1}
                                                        _ = v20
                                                    }
                                                    _ = v15
                                                    _ = v17
                                                }
                                                if v6 {
                                                    v1 = (v1 | v1)
                                                } else {
                                                    v21 := GenS1{F0: (fi7 / (((fi11 ^ 1)) | 1)), F1: 695, F2: 0}
                                                    v22 := GenS2{F0: -2147483648}
                                                    v6 = v6
                                                    v23 := []int{fi7, fi11, (fi11 + (fi7 % ((fi11) | 1)))}
                                                    fmt.Println(v8, v2, fi7)
                                                    var v24 int16 = -1
                                                    _ = v21
                                                    _ = v22
                                                    _ = v23
                                                    _ = v24
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    v25 := GenS3{F0: 58, F1: v3, F2: v1}
                                                    _ = v25
                                                case "a":
                                                    v6 = false
                                                }
                                                _ = fi11
                                            }
                                            v26 := []int{fi7, (fi7 + 4294967295), (^(fi7))}
                                            v27 := GenS4{F0: 711, F1: (^(v2)), F2: v2}
                                            _ = v26
                                            _ = v27
                                        }
                                    case -2147483648:
                                        fmt.Println(v2, v3, v1)
                                    default:
                                        switch fi7 {
                                        case -91:
                                            switch fi7 {
                                            case 1:
                                                v28 := GenS5{F0: -32768, F1: 2}
                                                _ = v28
                                            default:
                                                v29 := GenS6{F0: (("" + (("\n\t" + v3) + v3)) + ((("a" + ("" + "a")) + v3) + v3)), F1: 9223372036854775807, F2: -712}
                                                _ = v29
                                            }
                                        case 803:
                                            v8 = (v8 & v8)
                                        case 9223372036854775807:
                                            if v6 {
                                                var v30 float32 = (3.14159 * 0.1)
                                                fmt.Println(v6)
                                                v6 = true
                                                _ = v30
                                            }
                                        default:
                                            if (1000 >= int16(v2)) {
                                                v31 := []int{0, fi7}
                                                if len(v31) >= 2 {
                                                    v31 = v31[1:len(v31)]
                                                }
                                                if ((!false) && (false && v6)) {
                                                    switch "a" {
                                                    case "foo bar":
                                                        if (false || v6) {
                                                            v6 = (127 > 0)
                                                            for fi32 := 0; fi32 < 10; fi32++ {
                                                                _ = len(v31)
                                                                fmt.Println(v2, v8, v3)
                                                                _ = fi32
                                                            }
                                                            v33 := GenS7{F0: v6}
                                                            _ = v33
                                                        }
                                                    case "\n\t":
                                                        if false {
                                                            for fi34 := 0; fi34 < 9; fi34++ {
                                                                fmt.Println(v1, v8, v1)
                                                                _ = fi34
                                                            }
                                                            fmt.Println(v8, v3)
                                                            v8 = (v8 + (v8 ^ (v8 >> (((v1 ^ (v1 << ((uint(v2)) & 63)))) & 63))))
                                                        } else {
                                                            if v6 {
                                                                if (v6 && v6) {
                                                                    for fi35 := 0; fi35 < 12; fi35++ {
                                                                        if (0 >= 128) {
                                                                            v1 -= 255
                                                                        }
                                                                        v36 := GenS3{F0: 128, F1: v3, F2: v1}
                                                                        var v37 uint = v1
                                                                        v38 := GenS8{F0: 1, F1: (^(fi35)), F2: 89}
                                                                        v6 = (!(!true))
                                                                        if v6 {
                                                                            v39 := GenS6{F0: "foo bar", F1: 1, F2: 1000}
                                                                            _ = v39
                                                                        } else {
                                                                            if v6 {
                                                                                if len(v31) > 0 {
                                                                                    v31[0] = fi35
                                                                                }
                                                                                switch int(v8) {
                                                                                case 601:
                                                                                    if (869 >= -926) {
                                                                                        fmt.Println(v6)
                                                                                        if (9223372036854775807 >= 42) {
                                                                                            if v6 {
                                                                                                v40 := GenS1{F0: fi35, F1: 2147483648, F2: -2147483648}
                                                                                                _ = v40
                                                                                            } else {
                                                                                                switch v3 {
                                                                                                case "a":
                                                                                                    var v41 uint32 = 4294967295
                                                                                                    _ = v41
                                                                                                case "foo bar":
                                                                                                    var v42 uint32 = 2147483647
                                                                                                    _ = v42
                                                                                                case "":
                                                                                                    if false {
                                                                                                        if len(v31) > 0 {
                                                                                                            v31[0] = (fi7 - fi7)
                                                                                                        }
                                                                                                        v8 -= (v8 & v8)
                                                                                                    }
                                                                                                default:
                                                                                                    v31 = append(v31, -554)
                                                                                                }
                                                                                                var v43 string = ((v3 + v3) + v3)
                                                                                                _ = v43
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v44 uint16 = v2
                                                                                        _ = v44
                                                                                    }
                                                                                default:
                                                                                    var v45 uint = v37
                                                                                    _ = v45
                                                                                }
                                                                                v46 := GenS5{F0: 1000, F1: 462}
                                                                                _ = v46
                                                                            } else {
                                                                                var v47 uint8 = 0
                                                                                v3 = ((v3 + "a") + (v3 + v3))
                                                                                switch (fi35 % ((fi35) | 1)) {
                                                                                case -70:
                                                                                    v48 := GenS4{F0: 502, F1: v2, F2: 733}
                                                                                    _ = v48
                                                                                }
                                                                                _ = v47
                                                                            }
                                                                        }
                                                                        _ = fi35
                                                                        _ = v36
                                                                        _ = v37
                                                                        _ = v38
                                                                    }
                                                                    switch (v3 + ("a" + v3)) {
                                                                    case "\n\t":
                                                                        v1 = (v1 % ((v1) | 1))
                                                                    case "a":
                                                                        var v49 int16 = -1
                                                                        _ = v49
                                                                    default:
                                                                        v1 += (v1 * (v1 - 1))
                                                                    }
                                                                    v8 -= (v8 * -128)
                                                                }
                                                            }
                                                            var v50 int8 = v8
                                                            var v51 bool = (42 != 42)
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    case "a":
                                                        if len(v31) > 0 {
                                                            v31[0] = fi7
                                                        }
                                                    }
                                                    fmt.Println(v3, v3)
                                                } else {
                                                    v3 = v3
                                                    if v6 {
                                                        for fi52 := 0; fi52 < 15; fi52++ {
                                                            var v53 float32 = float32((1e10 + (0.1 / ((-2.25 * (1e10 * (0.0 / 3.14159))) + 1e10))))
                                                            var v54 int = fi7
                                                            _ = fi52
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        switch v3 {
                                                        case "":
                                                            var v55 float32 = 0.1
                                                            _ = v55
                                                        case "hello":
                                                            fmt.Println(v3)
                                                        case "foo bar":
                                                            for fi56 := 0; fi56 < 10; fi56++ {
                                                                switch 205 {
                                                                case -118:
                                                                    _ = len(v31)
                                                                case 4294967295:
                                                                    v57 := GenS9{F0: -1, F1: 0, F2: -1}
                                                                    if v6 {
                                                                        if true {
                                                                            if len(v31) > 0 {
                                                                                v31[0] = fi56
                                                                            }
                                                                            switch (fi56 + (-(fi7))) {
                                                                            case -9223372036854775808:
                                                                                v6 = (v1 >= 2)
                                                                            case 9223372036854775807:
                                                                                var v58 string = v3
                                                                                _ = v58
                                                                            default:
                                                                                var v59 int8 = v8
                                                                                _ = v59
                                                                            }
                                                                            v60 := GenS10{F0: fi7, F1: -991, F2: 483}
                                                                            _ = v60
                                                                        }
                                                                        if v6 {
                                                                            if v6 {
                                                                                switch 1 {
                                                                                case -264:
                                                                                    var v61 int64 = 42
                                                                                    _ = v61
                                                                                case 128:
                                                                                    switch (fi7 + 817) {
                                                                                    case 128:
                                                                                        _ = len(v31)
                                                                                    case 9223372036854775807:
                                                                                        v6 = v6
                                                                                        v62 := GenS11{F0: 127}
                                                                                        _ = v62
                                                                                    default:
                                                                                        fmt.Println(v3)
                                                                                    }
                                                                                }
                                                                                v63 := GenS5{F0: 1, F1: 4294967295}
                                                                                var v64 int = 128
                                                                                v8 = 1
                                                                                fmt.Println(fi7, v64)
                                                                                _ = v63
                                                                                _ = v64
                                                                            }
                                                                            fmt.Println(v1)
                                                                            var v65 int8 = (-(v8))
                                                                            _ = v65
                                                                        }
                                                                        v1 = v1
                                                                    }
                                                                    var v66 float32 = 0.0
                                                                    _ = v57
                                                                    _ = v66
                                                                }
                                                                var v67 float32 = -2.25
                                                                var v68 bool = (v6 || true)
                                                                _ = fi56
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        }
                                                    } else {
                                                        v69 := GenS12{F0: 127}
                                                        _ = v69
                                                    }
                                                }
                                                _ = v31
                                            } else {
                                                v70 := []int{}
                                                _ = v70
                                            }
                                        }
                                    }
                                }
                            case "hello":
                                var v71 uint64 = 299
                                _ = v71
                            }
                            _ = fi7
                            _ = v8
                        }
                        for fi72 := 0; fi72 < 3; fi72++ {
                            v73 := []int{(fi72 >> (((^(v1))) & 63)), -1, (fi72 >> ((255) & 63))}
                            v1 -= 0
                            if (42 == 83) {
                                v6 = v6
                            } else {
                                v74 := GenS13{F0: "\n\t", F1: (v2 & (^(v2))), F2: (fi72 * 127)}
                                _ = v74
                            }
                            _ = fi72
                            _ = v73
                        }
                        _ = v6
                    }
                case 156:
                    v1 -= 4294967295
                case 305:
                    switch -979 {
                    case 127:
                        v75 := []int{823, -1}
                        _ = v75
                    case -160:
                        v76 := []int{-272}
                        _ = v76
                    }
                default:
                    for fi77 := 0; fi77 < 12; fi77++ {
                        for fi78 := 0; fi78 < 18; fi78++ {
                            v79 := []int{fi77, 1}
                            switch (v3 + (v3 + "a")) {
                            case "\n\t":
                                if len(v79) > 0 {
                                    v79[0] = fi78
                                }
                            }
                            v79 = append(v79, fi78)
                            v1 = (v1 * v1)
                            v79 = append(v79, fi77)
                            var v80 int64 = 9223372036854775807
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                        _ = fi77
                    }
                }
            default:
                v81 := GenS4{F0: -1, F1: v2, F2: (v2 - (v2 - v2))}
                _ = v81
            }
        default:
            v82 := []int{-726, 154, -1, -2147483648}
            _ = v82
        }
    }
    v83 := []int{4294967295}
    fmt.Println(v3)
    v84 := GenS14{F0: 0}
    fmt.Println(v2, v2, v3)
    for fi85 := 0; fi85 < 13; fi85++ {
        switch v3 {
        case "foo bar":
            switch (fi85 * fi85) {
            case -9223372036854775808:
                var v86 uint64 = 578
                _ = v86
            default:
                if len(v83) >= 2 {
                    v83 = v83[1:len(v83)]
                }
            }
        case "\n\t":
            fmt.Println(fi85)
        case "hello":
            v3 = "\n\t"
        default:
            v3 = "hello"
        }
        for fi87 := 0; fi87 < 6; fi87++ {
            v88 := GenS11{F0: 215}
            _ = len(v83)
            _ = fi87
            _ = v88
        }
        _ = fi85
    }
    var v89 int8 = 1
    _ = len(v83)
    var v90 uint8 = 233
    var v91 float32 = 0.0
    switch "a" {
    case "a":
        if true {
            var v92 int16 = 32767
            v2 = 65535
            if (false && true) {
                v2 = (v2 >> ((18446744073709551615) & 63))
            }
            _ = v92
        } else {
            fmt.Println(v91, v91)
            v93 := GenS15{F0: -1}
            _ = v93
        }
    case "hello":
        v94 := GenS8{F0: v89, F1: 2147483647, F2: (^(v1))}
        _ = v94
    case "":
        if len(v83) > 0 {
            v83[0] = 474
        }
    default:
        v1 = 769
    }
    if (!(3.14159 < float32(1e10))) {
        v3 = v3
        v95 := GenS3{F0: v90, F1: ((v3 + v3) + "foo bar"), F2: 515}
        if true {
            v96 := GenS16{F0: "foo bar", F1: -18, F2: 2147483647}
            v91 += (v91 - float32(0.0))
            if false {
                v90 = v90
                var v97 float64 = 0.0
                var v98 float64 = (float64(v91) * 0.1)
                _ = v97
                _ = v98
            }
            _ = v96
        }
        _ = v95
    }
    for fi99 := 0; fi99 < 18; fi99++ {
        var v100 uint16 = (^(v2))
        v101 := GenS17{F0: 18446744073709551615}
        switch ((v3 + "\n\t") + v3) {
        case "foo bar":
            v102 := GenS1{F0: fi99, F1: 2147483648, F2: 42}
            _ = v102
        default:
            for fi103 := 0; fi103 < 18; fi103++ {
                var v104 bool = (false || false)
                v105 := GenS18{F0: v89, F1: -2147483648, F2: (fi99 << (((v1 + (v1 * uint(v100)))) & 63))}
                fmt.Println(v2, v2)
                v106 := GenS18{F0: (v89 | v89), F1: (fi103 >> ((1) & 63)), F2: (fi103 << ((v1) & 63))}
                switch v3 {
                case "\n\t":
                    v91 = v91
                case "a":
                    var v107 uint64 = 591
                    _ = v107
                case "foo bar":
                    v89 += v89
                }
                _ = fi103
                _ = v104
                _ = v105
                _ = v106
            }
        }
        _ = fi99
        _ = v100
        _ = v101
    }
    var v108 int64 = -681
    for fi109 := 0; fi109 < 11; fi109++ {
        if true {
            for fi110 := 0; fi110 < 11; fi110++ {
                if (1 == -32768) {
                    v111 := GenS12{F0: v90}
                    if (false || (0 <= 65535)) {
                        if true {
                            var v112 float64 = float64(((3.14159 + ((v91 + v91) - ((1e10 / (v91 / v91)) + 3.14159))) + v91))
                            _ = v112
                        } else {
                            if false {
                                v89 += (v89 | v89)
                                fmt.Println(v108)
                                var v113 float64 = 1.5
                                _ = v113
                            } else {
                                v114 := GenS19{F0: 1, F1: (((v3 + "hello") + ("hello" + "a")) + (v3 + ((v3 + v3) + "foo bar")))}
                                v90 += 127
                                if ((false && (uint(v90) < v1)) || true) {
                                    v115 := GenS20{F0: 42, F1: uint64(v1), F2: 2}
                                    fmt.Println(fi109, v89)
                                    switch (fi110 ^ fi109) {
                                    case -700:
                                        var v116 string = v3
                                        _ = v116
                                    case 42:
                                        var v117 float32 = 1.5
                                        _ = v117
                                    }
                                    _ = v115
                                }
                                switch (v3 + v3) {
                                case "":
                                    v90 += 127
                                case "a":
                                    var v118 string = (v3 + ("foo bar" + ("a" + ((v3 + ("hello" + "a")) + v3))))
                                    _ = v118
                                }
                                v119 := GenS6{F0: v3, F1: v108, F2: -1}
                                _ = v114
                                _ = v119
                            }
                            fmt.Println(fi109, v91, v2)
                            v83 = append(v83, fi109)
                        }
                    } else {
                        var v120 int16 = -793
                        _ = v120
                    }
                    var v121 int32 = 127
                    _ = v111
                    _ = v121
                }
                fmt.Println(v108, v89)
                v122 := GenS3{F0: uint8(v2), F1: "", F2: (v1 * (^(v1)))}
                v123 := GenS8{F0: 100, F1: fi109, F2: (^(v1))}
                fmt.Println(v90, fi109)
                v2 += (v2 ^ (v2 ^ v2))
                fmt.Println(v90)
                _ = fi110
                _ = v122
                _ = v123
            }
        }
        for fi124 := 0; fi124 < 14; fi124++ {
            v83 = append(v83, fi124)
            _ = fi124
        }
        _ = fi109
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v3)
    case "hello":
        v125 := GenS21{F0: 127}
        _ = v125
    default:
        var v126 int16 = -32768
        _ = v126
    }
    var v127 string = ("a" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v108
    _ = v127
    _ = v2
    _ = v3
    _ = v83
    _ = v84
    _ = v89
    _ = v90
    _ = v91
}
