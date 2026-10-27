package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 int64; F1 uint; F2 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16 }
func genFn0(p0 int16, p1 int, p2 int) int16 {
    var v6 uint32 = 2
    for fi7 := 0; fi7 < 16; fi7++ {
        switch ("\n\t" + "a") {
        case "\n\t":
            v8 := GenS1{F0: 84, F1: (p0 * p0)}
            _ = v8
        case "hello":
            var v9 int32 = int32(v6)
            _ = v9
        }
        var v10 bool = false
        if (!true) {
            if (32767 <= 2) {
                var v11 bool = false
                if v11 {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 int = (^(fi7))
                        if v11 {
                            v14 := GenS1{F0: 608, F1: 32767}
                            v15 := []int{(fi7 << ((0) & 63))}
                            _ = v14
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                    switch ("hello" + "") {
                    case "\n\t":
                        switch 42 {
                        case 0:
                            v16 := []int{p2}
                            _ = v16
                        case 32768:
                            if v10 {
                                v17 := GenS1{F0: 354, F1: p0}
                                _ = v17
                            } else {
                                v11 = (v10 || v10)
                                v6 += (^(v6))
                            }
                        }
                    }
                } else {
                    if true {
                        if true {
                            fmt.Println(v6)
                        } else {
                            p2 = p1
                            p2 += -2147483648
                            v18 := []int{}
                            _ = v18
                        }
                        if (false || (0 > 65535)) {
                            if (!((^(p0)) < (p0 + p0))) {
                                var v19 uint = 42
                                var v20 string = "foo bar"
                                _ = v19
                                _ = v20
                            }
                            v10 = (-2.25 == float32(float64((((float32(0.0) / 1e10) * 3.14159) - 1e10))))
                            for fi21 := 0; fi21 < 16; fi21++ {
                                v22 := []int{p1, 856, (-(fi7))}
                                v22 = append(v22, fi21)
                                v23 := GenS2{F0: false}
                                var v24 int16 = 32767
                                var v25 float64 = -2.25
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        }
                        p2 = fi7
                    } else {
                        fmt.Println(p1, v11)
                        var v26 uint16 = 65535
                        var v27 bool = (42 <= 0)
                        _ = v26
                        _ = v27
                    }
                    var v28 int16 = 166
                    _ = v28
                }
                v29 := []int{(p1 << ((255) & 63)), (p2 | (-(fi7))), -9223372036854775808, p1}
                _ = v11
                _ = v29
            } else {
                var v30 uint32 = v6
                _ = v30
            }
            v31 := GenS1{F0: 18446744073709551615, F1: p0}
            _ = v31
        }
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 1000
    v3 := []int{116, 30, 42, 4294967295}
    v4 := []string{}
    var v5 string = ""
    if true {
        fmt.Println(v5)
    } else {
        v1 += v1
        switch -1 {
        case -41:
            if len(v4) > 0 {
                v4[0] = (((("foo bar" + v5) + "foo bar") + "foo bar") + v5)
            }
        case 687:
            v4 = append(v4, ("a" + "a"))
        default:
            v32 := GenS2{F0: false}
            _ = v32
        }
    }
    fmt.Println(v5)
    var v33 int64 = 1
    switch ("a" + v5) {
    case "foo bar":
        v2 += 1000
    }
    v2 -= 195
    fmt.Println(v2)
    var v34 float64 = float64(0.0)
    var v35 float64 = 1e10
    var v36 int32 = 42
    fmt.Println(v1, v34, v33)
    var v37 bool = false
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := GenS3{F0: 2, F1: v37}
        if v37 {
            if (v37 && (v37 || v37)) {
                v1 = (v1 & (^(v1)))
            }
        }
        v40 := GenS4{F0: v5, F1: v36}
        _ = fi38
        _ = v39
        _ = v40
    }
    switch "\n\t" {
    case "":
        switch v5 {
        case "hello":
            for fi41 := 0; fi41 < 5; fi41++ {
                v42 := GenS5{F0: v5, F1: false}
                if true {
                    v43 := GenS3{F0: 2147483647, F1: (255 == 229)}
                    v34 = (((-2.25 - v35) / v35) * (v34 - v34))
                    _ = v43
                }
                _ = fi41
                _ = v42
            }
        case "\n\t":
            v4 = append(v4, (v5 + "a"))
        }
    case "a":
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 bool = true
            v46 := GenS2{F0: v37}
            _ = fi44
            _ = v45
            _ = v46
        }
    case "\n\t":
        if v37 {
            if (905 > 421) {
                if (v37 && (!(!v37))) {
                    switch 4294967295 {
                    case -9223372036854775808:
                        for fi47 := 0; fi47 < 10; fi47++ {
                            for fi48 := 0; fi48 < 8; fi48++ {
                                if v37 {
                                    v3 = append(v3, fi48)
                                    v3 = append(v3, 1)
                                }
                                if (0 >= 92) {
                                    fmt.Println(v34)
                                }
                                if v37 {
                                    var v49 uint16 = 6
                                    if false {
                                        if v37 {
                                            var v50 int64 = v33
                                            switch int(v1) {
                                            case 9223372036854775807:
                                                if (1e10 <= (-2.25 - 0.1)) {
                                                    if len(v4) > 0 {
                                                        v4[0] = v5
                                                    }
                                                    fmt.Println(fi48, v2)
                                                    v4 = append(v4, (((v5 + "hello") + "a") + v5))
                                                } else {
                                                    v50 = (v33 * (-(v33)))
                                                    v35 += 1.5
                                                }
                                            case -224:
                                                if v37 {
                                                    var v51 uint16 = v49
                                                    _ = v51
                                                } else {
                                                    var v52 int32 = v36
                                                    v2 = (^(v2))
                                                    _ = v52
                                                }
                                            case -637:
                                                fmt.Println(fi48)
                                            }
                                            if v37 {
                                                v36 += v36
                                            }
                                            _ = v50
                                        } else {
                                            var v53 int = fi47
                                            switch (fi47 & (v53 >> ((18446744073709551615) & 63))) {
                                            case -2147483648:
                                                if (18446744073709551615 >= 18446744073709551615) {
                                                    v3 = append(v3, v53)
                                                    v36 = -2147483648
                                                    fmt.Println(v53, v34, v33)
                                                }
                                            case 1:
                                                fmt.Println(fi47)
                                            case 32768:
                                                if false {
                                                    fmt.Println(fi48, v5)
                                                    v1 += v1
                                                    v4 = append(v4, (v5 + "\n\t"))
                                                    v54 := GenS6{F0: 18446744073709551615, F1: (-(v33))}
                                                    v4 = append(v4, "foo bar")
                                                    _ = v54
                                                }
                                            }
                                            fmt.Println(v2)
                                            var v55 uint8 = uint8(v49)
                                            _ = v53
                                            _ = v55
                                        }
                                        v36 = v36
                                        if v37 {
                                            var v56 float64 = v34
                                            var v57 int8 = 0
                                            var v58 uint64 = 18446744073709551615
                                            _ = v56
                                            _ = v57
                                            _ = v58
                                        }
                                    } else {
                                        if ((v1 ^ v1) == (v1 ^ (v1 % ((v1) | 1)))) {
                                            _ = len(v3)
                                            switch "hello" {
                                            case "hello":
                                                v1 += (v1 * (v1 / (((v1 << ((1) & 63))) | 1)))
                                                v4 = append(v4, ((("" + ((((("foo bar" + "hello") + (v5 + ((v5 + v5) + (v5 + v5)))) + ("" + v5)) + "a") + v5)) + ((v5 + "") + (v5 + "a"))) + v5))
                                                v59 := GenS1{F0: v1, F1: -1}
                                                _ = v59
                                            }
                                            if v37 {
                                                if (false || (!v37)) {
                                                    v33 -= v33
                                                } else {
                                                    switch v5 {
                                                    case "a":
                                                        if v37 {
                                                            switch "a" {
                                                            case "foo bar":
                                                                fmt.Println(v35, v37)
                                                            case "a":
                                                                var v60 uint16 = (^(v49))
                                                                _ = v60
                                                            case "hello":
                                                                v35 = ((v34 + v34) + (v35 - float64(1.5)))
                                                            default:
                                                                v61 := GenS6{F0: (v1 % ((2) | 1)), F1: 1}
                                                                _ = v61
                                                            }
                                                        }
                                                    case "foo bar":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "\n\t":
                                                        switch fi47 {
                                                        case 2147483647:
                                                            fmt.Println(fi47)
                                                        }
                                                    default:
                                                        v62 := GenS3{F0: uint32(v1), F1: (!(v49 != v49))}
                                                        _ = v62
                                                    }
                                                    v2 = 32767
                                                    v4 = append(v4, v5)
                                                }
                                            }
                                        } else {
                                            switch v5 {
                                            case "":
                                                var v63 uint16 = (^(v49))
                                                _ = v63
                                            case "\n\t":
                                                var v64 int64 = (^(v33))
                                                _ = v64
                                            }
                                            v65 := GenS6{F0: v1, F1: (v33 ^ -9223372036854775808)}
                                            var v66 uint16 = 695
                                            _ = v65
                                            _ = v66
                                        }
                                        v49 = v49
                                    }
                                    var v67 uint64 = 0
                                    _ = v49
                                    _ = v67
                                } else {
                                    var v68 uint8 = 2
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v69 := GenS7{F0: (v36 ^ (v36 * 388)), F1: (v1 * (v1 << (((v1 % ((uint(v68)) | 1))) & 63)))}
                                        _ = v69
                                    } else {
                                        var v70 uint32 = 1
                                        _ = len(v4)
                                        if v37 {
                                            switch int(v33) {
                                            case 128:
                                                fmt.Println(v1, v70, v36)
                                            case -735:
                                                if len(v4) > 0 {
                                                    v4[0] = ""
                                                }
                                            default:
                                                if v37 {
                                                    if true {
                                                        if v37 {
                                                            var v71 int = (-(fi47))
                                                            v68 += v68
                                                            v72 := GenS8{F0: v33, F1: v1, F2: v70}
                                                            _ = v71
                                                            _ = v72
                                                        } else {
                                                            v3 = append(v3, (fi47 - (fi47 ^ (fi48 | fi47))))
                                                            v70 = (v70 % ((2147483647) | 1))
                                                        }
                                                        var v73 float64 = 1.5
                                                        _ = v73
                                                    }
                                                    fmt.Println(v5, v1)
                                                    var v74 uint16 = 65535
                                                    _ = v74
                                                } else {
                                                    if (-1 < 100) {
                                                        v1 += (v1 << ((v1) & 63))
                                                    } else {
                                                        if (99 == (v68 / (((v68 / (((v68 + v68)) | 1))) | 1))) {
                                                            var v75 uint64 = 2
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v37 = v37
                                                            _ = v75
                                                        } else {
                                                            switch ("\n\t" + "") {
                                                            case "a":
                                                                v76 := GenS9{F0: 380}
                                                                _ = v76
                                                            }
                                                        }
                                                        var v77 uint8 = (v68 | 120)
                                                        switch (v5 + (v5 + v5)) {
                                                        case "\n\t":
                                                            v78 := GenS1{F0: 488, F1: v2}
                                                            _ = v78
                                                        }
                                                        v70 -= (v70 ^ v70)
                                                        v3 = append(v3, fi47)
                                                        var v79 string = v5
                                                        _ = v77
                                                        _ = v79
                                                    }
                                                    v35 = float64((3.14159 + 0.1))
                                                    v36 = (v36 * v36)
                                                    switch v5 {
                                                    case "a":
                                                        fmt.Println(fi47)
                                                    case "foo bar":
                                                        v70 = v70
                                                    }
                                                    fmt.Println(v36, v33)
                                                }
                                            }
                                        } else {
                                            v80 := GenS6{F0: 255, F1: v33}
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(v34, v35, fi48)
                                            default:
                                                v34 += float64((1e10 - 1.5))
                                            }
                                            _ = v80
                                        }
                                        _ = v70
                                    }
                                    _ = v68
                                }
                                _ = fi48
                            }
                            _ = fi47
                        }
                    default:
                        v81 := GenS9{F0: 617}
                        _ = v81
                    }
                    for fi82 := 0; fi82 < 8; fi82++ {
                        v83 := GenS2{F0: (127 <= 1)}
                        switch ((v5 + ((v5 + v5) + "\n\t")) + v5) {
                        case "":
                            var v84 int = fi82
                            _ = v84
                        case "hello":
                            v35 += (float64(3.14159) * (1.5 + 1e10))
                        case "foo bar":
                            if v37 {
                                var v85 bool = v37
                                if false {
                                    v37 = (1.5 <= 0.0)
                                    for fi86 := 0; fi86 < 6; fi86++ {
                                        v87 := GenS8{F0: v33, F1: v1, F2: 2147483648}
                                        v4 = append(v4, v5)
                                        v88 := GenS2{F0: false}
                                        _ = fi86
                                        _ = v87
                                        _ = v88
                                    }
                                    v85 = (958 == v1)
                                }
                                _ = v85
                            } else {
                                v4 = append(v4, v5)
                                var v89 uint16 = 42
                                var v90 bool = (v37 || false)
                                _ = v89
                                _ = v90
                            }
                        }
                        _ = fi82
                        _ = v83
                    }
                } else {
                    var v91 float32 = (-2.25 - 1.5)
                    _ = v91
                }
                v35 -= (v34 * (((1.5 + 3.14159) - 0.0) + (1e10 + v35)))
            } else {
                if ((float32(v34) * (1e10 - (3.14159 - 3.14159))) != 3.14159) {
                    var v92 int64 = (v33 % (((v33 << (((v1 * v1)) & 63))) | 1))
                    for fi93 := 0; fi93 < 7; fi93++ {
                        var v94 float32 = ((float32((((((0.0 / v35) * (float64((1e10 * 1.5)) - 1e10)) * v35) * (3.14159 + 1e10)) / (0.1 - 0.0))) - 0.1) / 1.5)
                        if v37 {
                            switch fi93 {
                            case 127:
                                if len(v4) > 0 {
                                    v4[0] = v5
                                }
                            case 32768:
                                v36 = (v36 / ((v36) | 1))
                            case 32767:
                                v33 = -1
                            }
                            for fi95 := 0; fi95 < 7; fi95++ {
                                if ((v1 / (((v1 * (v1 | v1))) | 1)) != (v1 | (v1 * (^(v1))))) {
                                    var v96 int8 = 42
                                    v94 = (v94 - (float32(-2.25) / 0.1))
                                    v37 = ((!false) || true)
                                    _ = v96
                                }
                                _ = fi95
                            }
                        }
                        _ = fi93
                        _ = v94
                    }
                    _ = v92
                } else {
                    if (!(true && v37)) {
                        switch ((("hello" + ("foo bar" + v5)) + (("" + ((v5 + ("a" + v5)) + "\n\t")) + v5)) + "a") {
                        case "hello":
                            var v97 int32 = (^(v36))
                            _ = v97
                        case "a":
                            v37 = (v36 >= (v36 << ((v1) & 63)))
                        case "foo bar":
                            v5 = (("\n\t" + "a") + v5)
                        default:
                            var v98 float32 = 1e10
                            _ = v98
                        }
                        v1 = v1
                        v3 = append(v3, 29)
                    } else {
                        var v99 string = "hello"
                        var v100 int = 1
                        switch (v100 ^ v100) {
                        case 0:
                            for fi101 := 0; fi101 < 13; fi101++ {
                                var v102 float32 = (1.5 + 1e10)
                                var v103 float64 = 1.5
                                switch v5 {
                                case "\n\t":
                                    switch v99 {
                                    case "foo bar":
                                        var v104 uint32 = 1
                                        _ = v104
                                    default:
                                        v105 := GenS4{F0: "hello", F1: v36}
                                        _ = v105
                                    }
                                case "a":
                                    v5 = v5
                                }
                                _ = fi101
                                _ = v102
                                _ = v103
                            }
                        case 128:
                            for fi106 := 0; fi106 < 10; fi106++ {
                                for fi107 := 0; fi107 < 9; fi107++ {
                                    v33 -= (v33 / (((v33 >> ((42) & 63))) | 1))
                                    v5 = (("\n\t" + ((((v5 + "\n\t") + v5) + "") + v5)) + (v99 + ""))
                                    _ = fi107
                                }
                                var v108 float64 = v35
                                for fi109 := 0; fi109 < 17; fi109++ {
                                    v100 = (fi109 % (((fi106 + fi106)) | 1))
                                    _ = fi109
                                }
                                _ = fi106
                                _ = v108
                            }
                        case -9223372036854775808:
                            switch v99 {
                            case "foo bar":
                                var v110 float64 = (0.0 * v34)
                                _ = v110
                            case "a":
                                v36 = (-(v36))
                            default:
                                var v111 uint8 = 42
                                _ = v111
                            }
                        default:
                            for fi112 := 0; fi112 < 4; fi112++ {
                                v113 := GenS7{F0: (v36 >> (((^(v1))) & 63)), F1: (v1 & 255)}
                                v1 = (v1 << (((v1 >> ((16) & 63))) & 63))
                                var v114 uint64 = 1
                                _ = fi112
                                _ = v113
                                _ = v114
                            }
                        }
                        _ = v100
                        _ = v99
                    }
                }
            }
        }
    default:
        if (((((3.14159 < 1e10) && (v2 != genFn0(v2, -986, 128))) && (0.0 >= v35)) && false) || ((v5 + (v5 + v5)) != ("" + ((v5 + v5) + ("hello" + "hello"))))) {
            for fi115 := 0; fi115 < 16; fi115++ {
                _ = len(v4)
                if true {
                    if (42 <= 138) {
                        var v116 uint64 = 290
                        if (true || (v1 < 784)) {
                            var v117 uint = (v1 - v1)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v117
                        }
                        _ = v116
                    }
                    v118 := GenS10{F0: v1}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v118
                }
                v4 = append(v4, (v5 + v5))
                _ = fi115
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    for fi119 := 0; fi119 < 3; fi119++ {
        v35 -= 1e10
        _ = fi119
    }
    switch 127 {
    case -9223372036854775808:
        if (!(v1 > 0)) {
            switch 98 {
            case 403:
                switch 42 {
                case -264:
                    fmt.Println(v37, v33)
                case -330:
                    var v120 string = ("\n\t" + "\n\t")
                    _ = v120
                default:
                    for fi121 := 0; fi121 < 8; fi121++ {
                        var v122 string = "a"
                        var v123 float64 = 1.5
                        var v124 uint32 = 65535
                        _ = fi121
                        _ = v122
                        _ = v123
                        _ = v124
                    }
                }
            case 32768:
                switch -9223372036854775808 {
                case 42:
                    v37 = true
                case 438:
                    v35 = v34
                case -56:
                    fmt.Println(v34, v34)
                default:
                    v36 += 1
                }
            case 2147483647:
                for fi125 := 0; fi125 < 9; fi125++ {
                    fmt.Println(v36, v1, v1)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi125
                }
            }
            v4 = append(v4, "a")
            v2 -= genFn0(genFn0(genFn0((v2 % (((v2 * v2)) | 1)), 436, -9223372036854775808), 4294967295, 32767), 505, 0)
        } else {
            v3 = append(v3, 32768)
            var v126 int64 = v33
            _ = v126
        }
    default:
        for fi127 := 0; fi127 < 1; fi127++ {
            var v128 uint = v1
            v36 = int32(v33)
            for fi129 := 0; fi129 < 5; fi129++ {
                v130 := GenS11{F0: 65535}
                v131 := GenS1{F0: v128, F1: 1000}
                switch fi127 {
                case 128:
                    v132 := GenS9{F0: 1}
                    _ = v132
                case 262:
                    var v133 float32 = (1.5 + (-2.25 / (float32(((1.5 - ((v34 / v34) - (1.5 - v34))) / (v35 - (0.0 + v35)))) + 1e10)))
                    v134 := GenS7{F0: 2147483647, F1: (v1 - (v1 & v1))}
                    _ = v133
                    _ = v134
                }
                if len(v4) > 0 {
                    v4[0] = (v5 + v5)
                }
                v3 = append(v3, fi129)
                _ = fi129
                _ = v130
                _ = v131
            }
            _ = fi127
            _ = v128
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
