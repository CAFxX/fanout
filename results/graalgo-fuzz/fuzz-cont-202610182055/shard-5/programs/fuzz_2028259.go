package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int8; F1 int; F2 int }
type GenS6 struct { F0 int64; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 int }
func genFn0(p0 uint, p1 uint8) uint64 {
    var v4 float32 = float32(3.14159)
    var v5 bool = (3.14159 == 0.1)
    v6 := []int{744}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint8 = 136
    var v2 string = "foo bar"
    var v3 uint = 65535
    v7 := GenS1{F0: 648, F1: 32767, F2: (0.0 > -2.25)}
    fmt.Println(v1, v2)
    v2 = v2
    fmt.Println(v3)
    var v8 float64 = 0.1
    v9 := []int{128}
    if true {
        var v10 uint64 = 4294967295
        _ = len(v9)
        _ = v10
    } else {
        switch ("" + v2) {
        case "foo bar":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        case "hello":
            var v11 int = 32767
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            if false {
                v13 := GenS1{F0: 0, F1: 42, F2: (!false)}
                var v14 uint16 = 42
                _ = v13
                _ = v14
            }
            switch 9223372036854775807 {
            case -343:
                for fi15 := 0; fi15 < 1; fi15++ {
                    fmt.Println(v1, v8)
                    fmt.Println(v1, v8)
                    _ = fi15
                }
            case 663:
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 int16 = -32768
                    var v18 int8 = -122
                    v9 = append(v9, 42)
                    fmt.Println(v2, v17)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                v3 = (v3 / (((v3 | v3)) | 1))
            }
            _ = fi12
        }
    }
    v19 := GenS1{F0: 588, F1: 32767, F2: true}
    var v20 string = v2
    v9 = append(v9, 32768)
    fmt.Println(v2)
    if ((42 >= 2147483647) || true) {
        if ("a" <= v20) {
            var v21 int = -914
            _ = v21
        }
    }
    if true {
        for fi22 := 0; fi22 < 8; fi22++ {
            if ((v20 + v20) > (v2 + ("\n\t" + ("hello" + v20)))) {
                var v23 uint8 = (v1 | (^(v1)))
                var v24 int8 = 0
                fmt.Println(v2)
                _ = v23
                _ = v24
            } else {
                if true {
                    var v25 int = fi22
                    v25 -= v25
                    if false {
                        v9 = append(v9, (fi22 << ((v3) & 63)))
                    } else {
                        var v26 float32 = float32((v8 - v8))
                        if (!false) {
                            switch 42 {
                            case 9223372036854775807:
                                for fi27 := 0; fi27 < 1; fi27++ {
                                    var v28 int8 = 127
                                    if (-993 >= int32(v1)) {
                                        switch fi22 {
                                        case 4294967295:
                                            v29 := GenS1{F0: 2147483647, F1: 0, F2: true}
                                            _ = v29
                                        case 127:
                                            var v30 int64 = -319
                                            _ = v30
                                        case -143:
                                            switch 32768 {
                                            case 1:
                                                v3 = v3
                                                fmt.Println(v26, v20, v20)
                                                v2 = (((v2 + v20) + v20) + v2)
                                            case 127:
                                                if (358 == fi27) {
                                                    _ = len(v9)
                                                    if false {
                                                        fmt.Println(v20)
                                                        var v31 int = int(v28)
                                                        var v32 float64 = v8
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        v2 = v20
                                                        _ = len(v9)
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        v33 := GenS2{F0: 0}
                                                        var v34 int8 = (-(v28))
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    var v35 bool = (!(("hello" + v2) <= (v2 + "foo bar")))
                                                    _ = v35
                                                }
                                                var v36 uint8 = (v1 / ((v1) | 1))
                                                v28 = v28
                                                _ = v36
                                            }
                                        default:
                                            if (true && true) {
                                                var v37 int16 = -32768
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                if (!false) {
                                                    v2 = (v2 + "hello")
                                                    if (((v28 / (((v28 >> (((v3 + 18446744073709551615)) & 63))) | 1)) >= (^(v28))) && (!(true && true))) {
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                    } else {
                                                        if (true && false) {
                                                            v3 -= (v3 << (((^(v3))) & 63))
                                                            _ = len(v9)
                                                        } else {
                                                            v20 = (v2 + v2)
                                                            v25 += (v25 | fi22)
                                                        }
                                                    }
                                                    v9 = append(v9, (fi22 | fi27))
                                                } else {
                                                    v9 = append(v9, (v25 << (((v3 - (v3 - 853))) & 63)))
                                                    var v38 float32 = (v26 - ((v26 / v26) / 3.14159))
                                                    switch v20 {
                                                    case "\n\t":
                                                        var v39 uint = 2
                                                        _ = v39
                                                    case "hello":
                                                        v20 = "foo bar"
                                                    case "foo bar":
                                                        v1 += (^(v1))
                                                    }
                                                    _ = v38
                                                }
                                                _ = v37
                                            } else {
                                                _ = len(v9)
                                            }
                                        }
                                        v40 := GenS1{F0: 4294967295, F1: 467, F2: false}
                                        var v41 uint64 = 656
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 uint32 = 2
                                    _ = fi27
                                    _ = v28
                                    _ = v42
                                }
                            case -1:
                                switch (v2 + (v20 + "\n\t")) {
                                case "":
                                    v9 = append(v9, (fi22 * (fi22 + (v25 & v25))))
                                }
                            default:
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    switch ("a" + "") {
                                    case "a":
                                        v3 = (v3 / (((v3 + v3)) | 1))
                                    case "hello":
                                        v25 += v25
                                    case "foo bar":
                                        switch (fi22 % (((fi22 << ((v3) & 63))) | 1)) {
                                        case -2147483648:
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        default:
                                            v2 = (v20 + v2)
                                        }
                                    }
                                    _ = fi43
                                }
                            }
                            if ((v8 / (((v8 - 1.5) + 0.1) + v8)) < v8) {
                                switch (^(v25)) {
                                case -9223372036854775808:
                                    for fi44 := 0; fi44 < 15; fi44++ {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v45 bool = (((v26 - v26) / (float32(v8) * (1e10 + v26))) == (1.5 * v26))
                                            _ = v45
                                        case "a":
                                            v46 := GenS2{F0: 2147483647}
                                            var v47 uint8 = v1
                                            _ = v46
                                            _ = v47
                                        case "\n\t":
                                            v8 = float64(v26)
                                        }
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        _ = fi44
                                    }
                                case -2147483648:
                                    v9 = append(v9, v25)
                                case 1:
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                default:
                                    if false {
                                        v25 += (fi22 % (((v25 >> ((v3) & 63))) | 1))
                                    } else {
                                        _ = len(v9)
                                        switch v20 {
                                        case "a":
                                            if true {
                                                if (65535 < 1) {
                                                    v20 = v2
                                                    v26 = v26
                                                }
                                                var v48 int32 = 42
                                                _ = len(v9)
                                                _ = v48
                                            } else {
                                                var v49 uint = v3
                                                _ = v49
                                            }
                                        default:
                                            v50 := GenS3{F0: 2147483647}
                                            _ = v50
                                        }
                                        for fi51 := 0; fi51 < 7; fi51++ {
                                            fmt.Println(v1, v1)
                                            if false {
                                                v52 := GenS2{F0: -1}
                                                var v53 string = ""
                                                var v54 uint = (^(v3))
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                            if len(v9) > 0 {
                                                v9[0] = 431
                                            }
                                            _ = fi51
                                        }
                                    }
                                }
                                var v55 int = -9223372036854775808
                                for fi56 := 0; fi56 < 1; fi56++ {
                                    if false {
                                        if false {
                                            v1 += v1
                                            if true {
                                                var v57 float32 = v26
                                                var v58 int16 = -734
                                                if false {
                                                    v59 := GenS2{F0: -2147483648}
                                                    v60 := GenS4{F0: (0 <= 0), F1: v3}
                                                    var v61 int16 = -32768
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                    fmt.Println(v25, v58)
                                                    if true {
                                                        v62 := GenS3{F0: -2147483648}
                                                        var v63 string = v2
                                                        if len(v9) > 0 {
                                                            v9[0] = -206
                                                        }
                                                        fmt.Println(v26)
                                                        var v64 int64 = 42
                                                        v57 += v57
                                                        fmt.Println(fi56, v3)
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                }
                                                _ = v57
                                                _ = v58
                                            } else {
                                                if true {
                                                    var v65 int8 = -1
                                                    _ = v65
                                                } else {
                                                    var v66 int = 2147483647
                                                    _ = v66
                                                }
                                            }
                                        } else {
                                            var v67 uint = 65535
                                            _ = v67
                                        }
                                        var v68 int64 = -1
                                        _ = v68
                                    } else {
                                        switch v2 {
                                        case "hello":
                                            v69 := GenS5{F0: 100, F1: 252, F2: fi56}
                                            _ = v69
                                        case "\n\t":
                                            switch v2 {
                                            case "a":
                                                fmt.Println(v2)
                                            case "hello":
                                                var v70 int = 127
                                                _ = v70
                                            default:
                                                v26 -= v26
                                            }
                                        default:
                                            if len(v9) > 0 {
                                                v9[0] = v55
                                            }
                                        }
                                        if len(v9) > 0 {
                                            v9[0] = v25
                                        }
                                    }
                                    var v71 uint16 = 2
                                    v72 := GenS4{F0: true, F1: v3}
                                    _ = fi56
                                    _ = v71
                                    _ = v72
                                }
                                _ = v55
                            }
                            if (((1.5 * 3.14159) >= 3.14159) || false) {
                                var v73 uint8 = (v1 * v1)
                                switch (fi22 | v25) {
                                case 0:
                                    switch int(v3) {
                                    case 9223372036854775807:
                                        if (!(-2.25 > v26)) {
                                            v20 = ""
                                        }
                                    case 741:
                                        var v74 string = "foo bar"
                                        _ = v74
                                    case 545:
                                        if (true && false) {
                                            var v75 int16 = -865
                                            if false {
                                                var v76 uint16 = 2
                                                for fi77 := 0; fi77 < 14; fi77++ {
                                                    v8 = float64(v26)
                                                    v9 = append(v9, fi77)
                                                    _ = fi77
                                                }
                                                _ = v76
                                            } else {
                                                v26 += v26
                                                v78 := GenS6{F0: 42, F1: v3, F2: 255}
                                                fmt.Println(v75, v75, v25)
                                                _ = v78
                                            }
                                            _ = v75
                                        }
                                    }
                                case 758:
                                    v9 = append(v9, (v25 % (((v25 & fi22)) | 1)))
                                }
                                _ = v73
                            } else {
                                switch v20 {
                                case "a":
                                    var v79 string = "hello"
                                    _ = v79
                                }
                                v80 := GenS6{F0: 42, F1: (v3 | v3), F2: 255}
                                _ = v80
                            }
                        }
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = v26
                    }
                    _ = v25
                }
            }
            v81 := GenS2{F0: 42}
            v20 = ((v20 + v20) + "foo bar")
            _ = fi22
            _ = v81
        }
    } else {
        v3 = v3
        v1 += 113
        if false {
            if (false || ((65535 < 42) || (2147483647 != 534))) {
                if (!false) {
                    if (!true) {
                        _ = len(v9)
                        switch v2 {
                        case "":
                            for fi82 := 0; fi82 < 12; fi82++ {
                                v83 := GenS6{F0: 9223372036854775807, F1: (v3 & v3), F2: 32767}
                                for fi84 := 0; fi84 < 9; fi84++ {
                                    if (true && (0 > 640)) {
                                        v8 -= v8
                                    } else {
                                        _ = len(v9)
                                        v85 := GenS6{F0: -683, F1: (v3 | (v3 << ((2) & 63))), F2: 11}
                                        v86 := GenS3{F0: int32(v3)}
                                        _ = v85
                                        _ = v86
                                    }
                                    _ = fi84
                                }
                                if ((v3 >= (v3 + v3)) && true) {
                                    fmt.Println(v2, fi82)
                                    for fi87 := 0; fi87 < 4; fi87++ {
                                        var v88 uint = v3
                                        fmt.Println(v88, fi87, v2)
                                        _ = fi87
                                        _ = v88
                                    }
                                    if (((0.1 - (1.5 / (float64(1e10) + v8))) + v8) >= ((v8 + v8) - v8)) {
                                        if false {
                                            switch fi82 {
                                            case 162:
                                                var v89 int64 = 9223372036854775807
                                                _ = v89
                                            case 32768:
                                                fmt.Println(v20, fi82)
                                            case 1:
                                                fmt.Println(v20)
                                            default:
                                                var v90 int8 = 1
                                                _ = v90
                                            }
                                        } else {
                                            switch fi82 {
                                            case 42:
                                                for fi91 := 0; fi91 < 5; fi91++ {
                                                    v3 = v3
                                                    _ = fi91
                                                }
                                            }
                                        }
                                        switch fi82 {
                                        case 4294967295:
                                            v2 = ("a" + ("a" + v20))
                                        case 1:
                                            _ = len(v9)
                                        case 32767:
                                            v8 = v8
                                        }
                                    } else {
                                        v2 = ("a" + "hello")
                                        v8 -= float64(3.14159)
                                        v92 := GenS4{F0: (false || false), F1: 2}
                                        _ = v92
                                    }
                                }
                                _ = fi82
                                _ = v83
                            }
                        default:
                            if false {
                                var v93 int32 = 42
                                switch "hello" {
                                case "foo bar":
                                    _ = len(v9)
                                case "hello":
                                    if false {
                                        v8 = v8
                                        v94 := GenS5{F0: -31, F1: 1, F2: -775}
                                        for fi95 := 0; fi95 < 7; fi95++ {
                                            if false {
                                                if (-1 > 397) {
                                                    var v96 float64 = (1.5 + v8)
                                                    v2 = "foo bar"
                                                    if true {
                                                        for fi97 := 0; fi97 < 18; fi97++ {
                                                            v2 = ("hello" + "")
                                                            _ = fi97
                                                        }
                                                    }
                                                    _ = v96
                                                } else {
                                                    v20 = "hello"
                                                    v93 = -196
                                                    if false {
                                                        v98 := GenS4{F0: false, F1: v3}
                                                        if len(v9) > 0 {
                                                            v9[0] = fi95
                                                        }
                                                        v3 += 951
                                                        _ = v98
                                                    }
                                                }
                                                fmt.Println(v2, v1, fi95)
                                                var v99 int32 = int32(v1)
                                                _ = v99
                                            } else {
                                                switch fi95 {
                                                case 2147483647:
                                                    if len(v9) > 0 {
                                                        v9[0] = fi95
                                                    }
                                                }
                                                switch (("foo bar" + (v2 + (("a" + "\n\t") + v2))) + "a") {
                                                case "hello":
                                                    v100 := GenS3{F0: v93}
                                                    _ = v100
                                                case "\n\t":
                                                    for fi101 := 0; fi101 < 1; fi101++ {
                                                        if (int32(v1) != 2147483647) {
                                                            v93 = v93
                                                            if true {
                                                                var v102 float32 = (0.1 - 1e10)
                                                                _ = v102
                                                            }
                                                            if true {
                                                                fmt.Println(v20)
                                                                if ((-128 == -128) && (275 < 42)) {
                                                                    if false {
                                                                        if false {
                                                                            switch 0 {
                                                                            case -191:
                                                                                v8 = v8
                                                                            case 9223372036854775807:
                                                                                var v103 uint8 = (v1 / (((v1 ^ 42)) | 1))
                                                                                _ = v103
                                                                            case -1:
                                                                                v3 = v3
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (-2.25 <= 1.5) {
                                                                            fmt.Println(v1)
                                                                            var v104 int64 = 0
                                                                            _ = v104
                                                                        }
                                                                        if (!(0.1 <= 0.1)) {
                                                                            fmt.Println(v1)
                                                                            var v105 uint = (v3 << ((0) & 63))
                                                                            if (true || true) {
                                                                                v9 = append(v9, (fi101 << ((v105) & 63)))
                                                                                fmt.Println(fi101)
                                                                                switch (v2 + "hello") {
                                                                                case "":
                                                                                    v9 = append(v9, (fi101 << (((v105 << ((v105) & 63))) & 63)))
                                                                                case "\n\t":
                                                                                    v20 = ("" + "")
                                                                                }
                                                                                v9 = append(v9, (fi101 & -76))
                                                                            } else {
                                                                                if len(v9) >= 2 {
                                                                                    v9 = v9[1:len(v9)]
                                                                                }
                                                                                if (v105 > (v105 << (((v3 ^ v3)) & 63))) {
                                                                                    var v106 uint32 = uint32(v1)
                                                                                    var v107 string = ((v2 + "hello") + (v2 + "hello"))
                                                                                    fmt.Println(v105)
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        v9 = append(v9, fi101)
                                                                                        if false {
                                                                                            var v108 int32 = 0
                                                                                            v8 = v8
                                                                                            _ = v108
                                                                                        }
                                                                                        if false {
                                                                                            v20 = v2
                                                                                        }
                                                                                    case "":
                                                                                        if len(v9) > 0 {
                                                                                            v9[0] = fi101
                                                                                        }
                                                                                    }
                                                                                    v20 = (v2 + "")
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                }
                                                                                if (v8 > v8) {
                                                                                    v105 -= (v105 | v105)
                                                                                    if (false || ((0.1 > (1.5 / -2.25)) || ((!false) || (!(v1 >= (v1 - (v1 + (v1 ^ (v1 + 104))))))))) {
                                                                                        var v109 uint = v3
                                                                                        _ = len(v9)
                                                                                        _ = v109
                                                                                    }
                                                                                }
                                                                                v20 = (v2 + (v20 + v20))
                                                                            }
                                                                            var v110 int16 = int16(fi95)
                                                                            switch (-(fi101)) {
                                                                            case 9223372036854775807:
                                                                                switch (fi95 * fi101) {
                                                                                case 1:
                                                                                    v20 = v20
                                                                                }
                                                                            }
                                                                            _ = v105
                                                                            _ = v110
                                                                        } else {
                                                                            v9 = append(v9, fi101)
                                                                            fmt.Println(fi101, v1, v93)
                                                                            var v111 int16 = 32767
                                                                            _ = v111
                                                                        }
                                                                    }
                                                                    if true {
                                                                        v112 := GenS6{F0: 932, F1: (v3 / ((137) | 1)), F2: 1}
                                                                        var v113 float32 = 0.1
                                                                        v8 += (1.5 + (v8 * v8))
                                                                        _ = v112
                                                                        _ = v113
                                                                    } else {
                                                                        switch v20 {
                                                                        case "\n\t":
                                                                            var v114 uint16 = 589
                                                                            _ = v114
                                                                        default:
                                                                            v1 -= v1
                                                                        }
                                                                        switch v2 {
                                                                        case "":
                                                                            var v115 string = ("\n\t" + v2)
                                                                            _ = v115
                                                                        }
                                                                        if true {
                                                                            v2 = "foo bar"
                                                                            v116 := GenS2{F0: v93}
                                                                            var v117 int16 = int16(v3)
                                                                            _ = v116
                                                                            _ = v117
                                                                        }
                                                                    }
                                                                    fmt.Println(v93, v1)
                                                                }
                                                            } else {
                                                                v9 = append(v9, fi95)
                                                                if true {
                                                                    fmt.Println(v3, fi95)
                                                                }
                                                            }
                                                        } else {
                                                            v9 = append(v9, -2147483648)
                                                            var v118 int64 = int64(v1)
                                                            _ = v118
                                                        }
                                                        _ = fi101
                                                    }
                                                case "foo bar":
                                                    var v119 int64 = 0
                                                    _ = v119
                                                default:
                                                    var v120 uint16 = 65535
                                                    _ = v120
                                                }
                                            }
                                            _ = fi95
                                        }
                                        _ = v94
                                    } else {
                                        if (v2 <= v20) {
                                            for fi121 := 0; fi121 < 12; fi121++ {
                                                for fi122 := 0; fi122 < 11; fi122++ {
                                                    switch v20 {
                                                    case "\n\t":
                                                        v1 = (v1 / (((v1 + 128)) | 1))
                                                    case "":
                                                        if ((-(fi121)) <= fi121) {
                                                            v93 += -2147483648
                                                        }
                                                    }
                                                    if len(v9) > 0 {
                                                        v9[0] = 0
                                                    }
                                                    var v123 float32 = ((1e10 + -2.25) + (float32(v8) * ((-2.25 * 0.0) - -2.25)))
                                                    _ = fi122
                                                    _ = v123
                                                }
                                                v124 := GenS7{F0: v1}
                                                v8 = 1e10
                                                _ = fi121
                                                _ = v124
                                            }
                                            var v125 bool = (1000 >= int16(v1))
                                            v125 = v125
                                            _ = v125
                                        }
                                    }
                                case "":
                                    var v126 int64 = 1
                                    _ = v126
                                default:
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                }
                                _ = v93
                            }
                        }
                        var v127 uint = (v3 & v3)
                        _ = v127
                    }
                    v1 = v1
                } else {
                    v128 := GenS8{F0: 941, F1: 255}
                    for fi129 := 0; fi129 < 13; fi129++ {
                        v8 = v8
                        var v130 int = (^(fi129))
                        if true {
                            v3 -= (v3 >> (((v3 | 18446744073709551615)) & 63))
                            for fi131 := 0; fi131 < 9; fi131++ {
                                var v132 int = v130
                                _ = fi131
                                _ = v132
                            }
                            if len(v9) > 0 {
                                v9[0] = (v130 | v130)
                            }
                        } else {
                            v8 += (v8 / 1.5)
                            if (true && false) {
                                if (!(1 == -1)) {
                                    var v133 bool = (genFn0(v3, 255) != 4294967295)
                                    v9 = append(v9, fi129)
                                    _ = v133
                                } else {
                                    v8 -= 1e10
                                    v3 += v3
                                    v20 = v2
                                }
                                if (true || true) {
                                    var v134 uint8 = (^(v1))
                                    switch v20 {
                                    case "":
                                        var v135 uint32 = 1
                                        _ = v135
                                    default:
                                        if len(v9) > 0 {
                                            v9[0] = fi129
                                        }
                                    }
                                    _ = v134
                                }
                            }
                            v20 = v20
                        }
                        _ = fi129
                        _ = v130
                    }
                    _ = v128
                }
                v136 := GenS9{F0: 927}
                _ = v136
            }
            fmt.Println(v2)
            for fi137 := 0; fi137 < 19; fi137++ {
                v2 = v2
                var v138 int32 = 0
                var v139 int32 = -2147483648
                _ = fi137
                _ = v138
                _ = v139
            }
        } else {
            switch ("" + (v20 + ("" + "foo bar"))) {
            case "a":
                var v140 int = -85
                _ = v140
            }
            if false {
                if false {
                    if false {
                        var v141 uint = 984
                        v8 += ((((((1e10 * 1e10) + ((0.0 - v8) * float64((0.1 + 0.1)))) * (v8 / v8)) / v8) - v8) + float64((((3.14159 + (0.1 * -2.25)) * 0.1) - 0.1)))
                        _ = v141
                    }
                    v142 := GenS3{F0: -2147483648}
                    _ = v142
                }
                var v143 float32 = 0.1
                _ = v143
            }
            v8 = v8
        }
    }
    switch ("hello" + v20) {
    case "foo bar":
        fmt.Println(v20)
    case "\n\t":
        if (!(-15 != -1)) {
            for fi144 := 0; fi144 < 17; fi144++ {
                v3 += v3
                v1 = 42
                _ = fi144
            }
            var v145 uint16 = 65535
            _ = v145
        }
    }
    fmt.Println(v8, v8, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
