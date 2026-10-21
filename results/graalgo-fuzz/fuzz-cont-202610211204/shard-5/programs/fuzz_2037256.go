package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 bool; F2 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 bool; F1 string; F2 int32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8; F1 uint; F2 int64 }
type GenS16 struct { F0 uint32; F1 uint32; F2 uint }
type GenS17 struct { F0 bool; F1 int16 }
type GenS18 struct { F0 uint64; F1 uint; F2 string }
type GenS19 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS20 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS21 struct { F0 int64; F1 bool }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 uint64 }
type GenS24 struct { F0 string; F1 string; F2 uint64 }
type GenS25 struct { F0 int; F1 uint64; F2 string }
type GenS26 struct { F0 uint16; F1 int }
type GenS27 struct { F0 bool }
type GenS28 struct { F0 bool }
func genFn0() int8 {
    var v7 int16 = -258
    v7 = -32768
    _ = v7
    return -48
}

func main() {
    var v1 uint8 = 128
    v2 := []string{}
    var v3 uint8 = (^(v1))
    var v4 int = 42
    var v5 int = -9223372036854775808
    var v6 float64 = (((float64(3.14159) - (-2.25 - 0.1)) / 1.5) + 3.14159)
    if true {
        v5 = 703
    } else {
        var v8 int8 = -15
        if true {
            v9 := GenS1{F0: 2147483648, F1: -820, F2: false}
            if false {
                switch -2147483648 {
                case -102:
                    if len(v2) > 0 {
                        v2[0] = ("hello" + ("a" + ""))
                    }
                case 42:
                    var v10 bool = ("hello" <= "")
                    _ = v10
                default:
                    switch (v4 >> ((2) & 63)) {
                    case -9223372036854775808:
                        var v11 int8 = v8
                        _ = v11
                    case 4294967295:
                        v6 -= float64(3.14159)
                    default:
                        var v12 uint8 = 1
                        _ = v12
                    }
                }
                var v13 int = v4
                _ = v13
            } else {
                if false {
                    v4 = (^(v5))
                    switch v5 {
                    case 2147483647:
                        fmt.Println(v3, v3)
                    case 1:
                        switch ("hello" + "hello") {
                        case "a":
                            var v14 bool = false
                            _ = v14
                        case "hello":
                            for fi15 := 0; fi15 < 5; fi15++ {
                                var v16 int16 = 0
                                _ = len(v2)
                                var v17 uint64 = 236
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                    }
                }
                v5 = (v5 % ((127) | 1))
                switch (^(v4)) {
                case 4294967295:
                    v6 = (v6 / (float64((3.14159 * 1.5)) + 0.1))
                }
            }
            var v18 float64 = 0.0
            _ = v18
            _ = v9
        }
        _ = v8
    }
    switch "foo bar" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    case "\n\t":
        v19 := GenS1{F0: 65535, F1: -32768, F2: false}
        _ = v19
    default:
        _ = len(v2)
    }
    if true {
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(v6, v4, fi20)
            v3 = (v3 + (^(v3)))
            if (42 >= 127) {
                if (-87 != -1) {
                    switch (("\n\t" + ("\n\t" + (("hello" + "a") + ("\n\t" + "\n\t")))) + "") {
                    case "":
                        if true {
                            for fi21 := 0; fi21 < 14; fi21++ {
                                switch (fi21 ^ (fi20 & (v4 | v4))) {
                                case 133:
                                    v1 -= v3
                                case 0:
                                    var v22 int16 = 1
                                    _ = v22
                                default:
                                    fmt.Println(v5, fi21, v1)
                                }
                                _ = fi21
                            }
                            if (("a" + ("\n\t" + "foo bar")) > "hello") {
                                var v23 uint16 = 472
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    v1 = v1
                                    var v25 uint = 18446744073709551615
                                    fmt.Println(v6)
                                    if (v23 < (v23 % ((v23) | 1))) {
                                        if len(v2) > 0 {
                                            v2[0] = ("hello" + "")
                                        }
                                    } else {
                                        v6 = 1e10
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    var v26 float64 = float64(0.0)
                                    _ = v26
                                case "hello":
                                    v2 = append(v2, ("hello" + ("\n\t" + "a")))
                                case "a":
                                    fmt.Println(fi20)
                                }
                                _ = v23
                            }
                            if true {
                                fmt.Println(v4)
                                switch 899 {
                                case -53:
                                    var v27 uint = 255
                                    _ = v27
                                case 863:
                                    var v28 uint32 = 2
                                    _ = v28
                                }
                                var v29 string = ("" + (("\n\t" + "\n\t") + ""))
                                _ = v29
                            }
                        } else {
                            if false {
                                fmt.Println(fi20)
                            } else {
                                for fi30 := 0; fi30 < 11; fi30++ {
                                    switch fi20 {
                                    case 0:
                                        var v31 uint64 = 11
                                        _ = v31
                                    case 643:
                                        var v32 int64 = 1
                                        _ = v32
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(fi20, v5, v4)
                                    case "a":
                                        var v33 int64 = 0
                                        _ = v33
                                    case "":
                                        var v34 uint64 = 871
                                        _ = v34
                                    default:
                                        var v35 uint16 = 2
                                        if false {
                                            v6 -= 1e10
                                            _ = len(v2)
                                        } else {
                                            if true {
                                                v4 = 1
                                                if (42 != 65535) {
                                                    fmt.Println(v1, v5, v5)
                                                    v36 := GenS1{F0: 2147483647, F1: -1, F2: true}
                                                    fmt.Println(v4, v1)
                                                    _ = v36
                                                }
                                                switch "hello" {
                                                case "a":
                                                    if (!true) {
                                                        var v37 string = ""
                                                        v1 = 128
                                                        var v38 int16 = 42
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                case "":
                                                    if len(v2) > 0 {
                                                        v2[0] = ("\n\t" + "a")
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        switch "a" {
                                                        case "\n\t":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            v35 += 32767
                                                        }
                                                    }
                                                    var v39 int16 = 920
                                                    switch ("foo bar" + "foo bar") {
                                                    case "":
                                                        var v40 float64 = 0.1
                                                        _ = v40
                                                    case "foo bar":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case "a":
                                                        if (-684 <= 9223372036854775807) {
                                                            switch -1 {
                                                            case 4294967295:
                                                                var v41 float64 = v6
                                                                _ = v41
                                                            default:
                                                                if len(v2) > 0 {
                                                                    v2[0] = "foo bar"
                                                                }
                                                            }
                                                            if false {
                                                                v5 += (fi30 - -9223372036854775808)
                                                                var v42 int32 = -95
                                                                v43 := GenS1{F0: 0, F1: v39, F2: true}
                                                                if (2147483648 < 4294967295) {
                                                                    v5 -= -696
                                                                    var v44 uint16 = (v35 * (v35 * (v35 << ((65535) & 63))))
                                                                    v35 -= 334
                                                                    _ = v44
                                                                }
                                                                _ = len(v2)
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                v45 := GenS2{F0: (^(v5)), F1: 0}
                                                                var v46 bool = true
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                            fmt.Println(v5, v39)
                                                            switch (v5 - (-(v4))) {
                                                            case -948:
                                                                _ = len(v2)
                                                            case 9223372036854775807:
                                                                v2 = append(v2, "hello")
                                                            }
                                                        } else {
                                                            if false {
                                                                if (352 != int32(v3)) {
                                                                    v47 := GenS2{F0: (v4 & fi20), F1: (v35 * v35)}
                                                                    var v48 uint32 = 733
                                                                    v1 = 127
                                                                    v2 = append(v2, ("\n\t" + "\n\t"))
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        fmt.Println(v3, v39)
                                                                    default:
                                                                        if false {
                                                                            if false {
                                                                                if true {
                                                                                    var v49 float64 = v6
                                                                                    switch fi20 {
                                                                                    case 32768:
                                                                                        v50 := GenS2{F0: -1, F1: 42}
                                                                                        _ = v50
                                                                                    default:
                                                                                        var v51 uint8 = (v1 - (v1 & (^(v1))))
                                                                                        _ = v51
                                                                                    }
                                                                                    _ = v49
                                                                                } else {
                                                                                    v4 = fi20
                                                                                }
                                                                                _ = len(v2)
                                                                                if (0 < fi30) {
                                                                                    fmt.Println(fi20, fi30, fi20)
                                                                                    var v52 int16 = v39
                                                                                    v35 = (v35 & (^(v35)))
                                                                                    _ = v52
                                                                                } else {
                                                                                    var v53 int64 = 42
                                                                                    _ = v53
                                                                                }
                                                                            }
                                                                            v54 := GenS3{F0: v48, F1: -648}
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    var v55 uint8 = v3
                                                                    v56 := GenS2{F0: (^(v4)), F1: (v35 & (v35 & v35))}
                                                                    var v57 uint32 = 543
                                                                    _ = v47
                                                                    _ = v48
                                                                    _ = v55
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                            }
                                                            var v58 string = "hello"
                                                            var v59 uint16 = v35
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    }
                                                    _ = v39
                                                default:
                                                    fmt.Println(v1)
                                                }
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch (("foo bar" + "a") + "") {
                                            case "":
                                                fmt.Println(v4, v6, v35)
                                            case "\n\t":
                                                var v60 uint16 = v35
                                                _ = v60
                                            case "hello":
                                                v61 := GenS4{F0: (v35 * v35), F1: (-(fi30))}
                                                _ = len(v2)
                                                v3 -= (v3 - 128)
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                                _ = v61
                                            }
                                        }
                                        var v62 uint = 609
                                        _ = v35
                                        _ = v62
                                    }
                                    v5 += (fi20 | fi30)
                                    var v63 int8 = 0
                                    v64 := GenS5{F0: (v1 ^ v3), F1: fi20}
                                    if false {
                                        var v65 int = v4
                                        _ = v65
                                    } else {
                                        if true {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        switch 9223372036854775807 {
                                        case -758:
                                            v6 = 3.14159
                                            v1 = v3
                                            v3 -= (v3 / ((v1) | 1))
                                        }
                                        var v66 uint32 = 789
                                        _ = v66
                                    }
                                    v1 = v3
                                    _ = fi30
                                    _ = v63
                                    _ = v64
                                }
                                v67 := GenS3{F0: 598, F1: 42}
                                if len(v2) > 0 {
                                    v2[0] = ("hello" + "foo bar")
                                }
                                _ = v67
                            }
                            switch "a" {
                            case "":
                                var v68 uint16 = 43
                                _ = v68
                            }
                        }
                    case "hello":
                        v1 = 2
                    case "\n\t":
                        for fi69 := 0; fi69 < 16; fi69++ {
                            if true {
                                var v70 float32 = (-2.25 * (((1e10 - 0.1) - (3.14159 + 0.0)) * (float32((-2.25 * float64(0.1))) * (1e10 * -2.25))))
                                if true {
                                    if (!(833 == 255)) {
                                        fmt.Println(v70, fi20, v70)
                                        v1 -= (v3 * (v1 & 255))
                                        var v71 string = ("a" + "\n\t")
                                        switch v71 {
                                        case "":
                                            v72 := GenS1{F0: 4294967295, F1: 32767, F2: false}
                                            _ = v72
                                        case "a":
                                            var v73 int8 = 100
                                            _ = v73
                                        default:
                                            switch -127 {
                                            case 128:
                                                if false {
                                                    v1 += (v1 & v3)
                                                } else {
                                                    var v74 uint64 = 971
                                                    var v75 bool = true
                                                    if (1.5 < (v70 / ((v70 + float32(v6)) + float32((0.1 * v6))))) {
                                                        v70 += 0.0
                                                        if (0 != v74) {
                                                            v74 += (v74 + 1)
                                                            v3 += v3
                                                            var v76 uint16 = 518
                                                            _ = v76
                                                        } else {
                                                            _ = len(v2)
                                                            if false {
                                                                switch "hello" {
                                                                case "":
                                                                    if v75 {
                                                                        _ = len(v2)
                                                                        var v77 uint = 1
                                                                        _ = v77
                                                                    }
                                                                    var v78 uint16 = 235
                                                                    var v79 int = (fi20 + v4)
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                            } else {
                                                                fmt.Println(v70)
                                                                switch v4 {
                                                                case 853:
                                                                    v74 += (v74 | v74)
                                                                }
                                                                switch -199 {
                                                                case 783:
                                                                    v80 := GenS2{F0: v4, F1: 255}
                                                                    _ = v80
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v74
                                                    _ = v75
                                                }
                                                _ = len(v2)
                                                v81 := GenS6{F0: 0}
                                                _ = v81
                                            case -2147483648:
                                                if true {
                                                    var v82 uint16 = 2
                                                    var v83 uint32 = uint32(v1)
                                                    v82 += (v82 % (((v82 ^ v82)) | 1))
                                                    v84 := GenS7{F0: v83, F1: (-1 == -414), F2: 42}
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    switch v71 {
                                                    case "foo bar":
                                                        switch (v71 + ("hello" + (v71 + v71))) {
                                                        case "foo bar":
                                                            v70 -= (v70 + (v70 + (v70 - 3.14159)))
                                                        case "hello":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case "a":
                                                            var v85 uint = 42
                                                            _ = v85
                                                        }
                                                    case "":
                                                        if false {
                                                            var v86 uint8 = 139
                                                            switch v4 {
                                                            case 128:
                                                                var v87 uint32 = 631
                                                                _ = v87
                                                            default:
                                                                var v88 uint64 = 42
                                                                _ = v88
                                                            }
                                                            var v89 uint = 4294967295
                                                            v1 = uint8(v89)
                                                            _ = v86
                                                            _ = v89
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    case "hello":
                                                        if (-1 == -1) {
                                                            switch fi20 {
                                                            case -617:
                                                                fmt.Println(v3, v70, fi20)
                                                            case 42:
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v90 uint16 = uint16(v1)
                                                                    _ = v90
                                                                case "":
                                                                    var v91 int8 = 39
                                                                    _ = v91
                                                                case "\n\t":
                                                                    v5 = (fi69 / ((32768) | 1))
                                                                }
                                                            }
                                                            fmt.Println(v1, v3, fi69)
                                                            var v92 uint = 990
                                                            _ = v92
                                                        } else {
                                                            v70 -= (0.0 - (float32((1.5 / 1.5)) * float32(0.0)))
                                                            var v93 uint64 = 2
                                                            v94 := GenS2{F0: fi69, F1: 781}
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                        if true {
                                                            if true {
                                                                if (0.1 == float64(v70)) {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if true {
                                                                            if (false && true) {
                                                                                var v95 int64 = 0
                                                                                _ = v95
                                                                            } else {
                                                                                var v96 bool = (-32768 > 42)
                                                                                v5 += fi20
                                                                                v97 := GenS6{F0: 0}
                                                                                _ = v96
                                                                                _ = v97
                                                                            }
                                                                        } else {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            v2 = append(v2, ((v71 + "foo bar") + "hello"))
                                                                            v98 := GenS5{F0: 42, F1: (v4 + v5)}
                                                                            var v99 uint64 = 509
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                    case "\n\t":
                                                                        if (((255 != 53) && (!(1 > 838))) || (v71 == ("a" + v71))) {
                                                                            var v100 bool = false
                                                                            if false {
                                                                                v5 -= fi20
                                                                            } else {
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    if v100 {
                                                                                        var v101 uint = 1
                                                                                        _ = v101
                                                                                    } else {
                                                                                        v100 = (!true)
                                                                                        var v102 uint16 = 65535
                                                                                        _ = v102
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v103 uint64 = 0
                                                                                    _ = v103
                                                                                default:
                                                                                    var v104 int8 = -19
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                            v100 = v100
                                                                            var v105 int16 = -339
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (v71 + (v71 + v71))
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v71
                                                                            }
                                                                            _ = v100
                                                                            _ = v105
                                                                        }
                                                                    case "hello":
                                                                        if (-1 >= -295) {
                                                                            var v106 int8 = -1
                                                                            _ = v106
                                                                        }
                                                                    default:
                                                                        v5 = (fi20 << ((42) & 63))
                                                                    }
                                                                } else {
                                                                    v4 = v5
                                                                    v107 := GenS3{F0: 26, F1: 1}
                                                                    var v108 uint8 = (v3 + (^(v1)))
                                                                    _ = v107
                                                                    _ = v108
                                                                }
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v109 float32 = v70
                                                                _ = v109
                                                            case "\n\t":
                                                                v6 = (v6 - (v6 - (1.5 + v6)))
                                                            }
                                                        }
                                                        if (true && ("hello" <= v71)) {
                                                            if false {
                                                                if (255 != 42) {
                                                                    if false {
                                                                        if true {
                                                                            v110 := GenS7{F0: 42, F1: true, F2: 39}
                                                                            var v111 int = v5
                                                                            _ = v110
                                                                            _ = v111
                                                                        }
                                                                        v2 = append(v2, ("\n\t" + "hello"))
                                                                    } else {
                                                                        if (!true) {
                                                                            fmt.Println(fi20)
                                                                            v112 := GenS7{F0: 0, F1: (126 < 2147483648), F2: 114}
                                                                            _ = v112
                                                                        }
                                                                        fmt.Println(fi69, v70, v3)
                                                                        if (1 <= 935) {
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                var v113 int16 = 1000
                                                                                _ = v113
                                                                            case "":
                                                                                var v114 int32 = 0
                                                                                _ = v114
                                                                            case "\n\t":
                                                                                var v115 int64 = -9223372036854775808
                                                                                _ = v115
                                                                            default:
                                                                                if true {
                                                                                    if (-2.25 == float64((1e10 * 3.14159))) {
                                                                                        v3 = v3
                                                                                        v5 = 9223372036854775807
                                                                                        var v116 float32 = v70
                                                                                        _ = v116
                                                                                    }
                                                                                    var v117 uint32 = 2
                                                                                    _ = v117
                                                                                } else {
                                                                                    v118 := GenS6{F0: 42}
                                                                                    v70 -= v70
                                                                                    _ = v118
                                                                                }
                                                                            }
                                                                            if false {
                                                                                if true {
                                                                                    switch (fi69 | (fi20 * fi20)) {
                                                                                    case 128:
                                                                                        _ = len(v2)
                                                                                    case -2147483648:
                                                                                        _ = len(v2)
                                                                                    }
                                                                                    var v119 int32 = -312
                                                                                    _ = v119
                                                                                } else {
                                                                                    v1 += v1
                                                                                }
                                                                            } else {
                                                                                var v120 int8 = 127
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    if (false && false) {
                                                                                        switch v71 {
                                                                                        case "foo bar":
                                                                                            v5 += fi69
                                                                                        case "a":
                                                                                            v121 := GenS8{F0: 32767}
                                                                                            _ = v121
                                                                                        default:
                                                                                            switch (v5 ^ v5) {
                                                                                            case 2147483647:
                                                                                                v1 = 255
                                                                                            case 42:
                                                                                                v3 = v1
                                                                                            case -2147483648:
                                                                                                var v122 uint8 = v1
                                                                                                _ = v122
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v123 uint32 = 65535
                                                                                        _ = v123
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v124 bool = true
                                                                                    _ = v124
                                                                                }
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = v71
                                                                                }
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                    }
                                                                    var v125 float32 = ((v70 * float32(float64((v70 + (v70 - 3.14159))))) * (1e10 * v70))
                                                                    _ = len(v2)
                                                                    _ = v125
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (v71 + "hello")
                                                                    }
                                                                }
                                                            }
                                                            v5 += 4294967295
                                                            v71 = v71
                                                        }
                                                    default:
                                                        if false {
                                                            fmt.Println(v70, v70, v3)
                                                            v126 := GenS9{F0: 399}
                                                            fmt.Println(v6, v4, v5)
                                                            _ = v126
                                                        } else {
                                                            if false {
                                                                var v127 uint = 42
                                                                _ = v127
                                                            } else {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                switch v5 {
                                                                case 861:
                                                                    v4 = v4
                                                                case 4294967295:
                                                                    if (true || ((-2.25 + v70) >= 1.5)) {
                                                                        fmt.Println(v3, v71, v3)
                                                                        fmt.Println(fi69, v5)
                                                                        v5 = 42
                                                                        if true {
                                                                            v128 := GenS7{F0: 444, F1: false, F2: 127}
                                                                            v4 -= v4
                                                                            v129 := GenS2{F0: fi20, F1: 621}
                                                                            _ = v128
                                                                            _ = v129
                                                                        }
                                                                        fmt.Println(v71, v4)
                                                                    }
                                                                default:
                                                                    var v130 uint8 = v3
                                                                    _ = v130
                                                                }
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    }
                                                    if false {
                                                        var v131 float32 = 1.5
                                                        if len(v2) > 0 {
                                                            v2[0] = (v71 + "a")
                                                        }
                                                        switch v71 {
                                                        case "foo bar":
                                                            _ = len(v2)
                                                        default:
                                                            var v132 uint64 = 2
                                                            _ = v132
                                                        }
                                                        _ = v131
                                                    } else {
                                                        v133 := GenS10{F0: 692, F1: -128}
                                                        if (!false) {
                                                            v134 := GenS11{F0: true, F1: "a", F2: 134}
                                                            if true {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                fmt.Println(v6, v6, fi20)
                                                            } else {
                                                                fmt.Println(v3, fi69)
                                                                if len(v2) > 0 {
                                                                    v2[0] = ""
                                                                }
                                                            }
                                                            _ = v134
                                                        } else {
                                                            var v135 int64 = int64(v1)
                                                            v1 -= 127
                                                            fmt.Println(v4)
                                                            _ = v135
                                                        }
                                                        _ = v133
                                                    }
                                                }
                                            case -141:
                                                if false {
                                                    if (!false) {
                                                        fmt.Println(v1, v5)
                                                        switch "foo bar" {
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = ("foo bar" + v71)
                                                            }
                                                        case "hello":
                                                            var v136 int = (v4 / ((4294967295) | 1))
                                                            _ = v136
                                                        default:
                                                            if (genFn0() < 63) {
                                                                v137 := GenS9{F0: 1}
                                                                if (((-2.25 + (1.5 + 0.0)) * float32((v6 + (v6 / v6)))) != v70) {
                                                                    switch 32768 {
                                                                    case 32767:
                                                                        v6 = (((v6 + (v6 * (v6 * (v6 + (v6 * v6))))) - v6) / (0.1 + v6))
                                                                    case -739:
                                                                        switch fi69 {
                                                                        case -518:
                                                                            var v138 uint8 = 217
                                                                            _ = v138
                                                                        case -994:
                                                                            var v139 uint64 = 42
                                                                            _ = v139
                                                                        case -987:
                                                                            var v140 uint64 = 0
                                                                            _ = v140
                                                                        }
                                                                    }
                                                                    v141 := GenS12{F0: (v3 / ((v1) | 1))}
                                                                    var v142 float64 = (v6 - ((float64(float32(0.0)) * float64(0.1)) + v6))
                                                                    var v143 uint64 = 1
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v144 bool = true
                                                                        _ = v144
                                                                    default:
                                                                        if (!(-1 < 948)) {
                                                                            v145 := GenS8{F0: 32767}
                                                                            v143 += 691
                                                                            _ = v145
                                                                        } else {
                                                                            if (!(765 == 32767)) {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            } else {
                                                                                v146 := GenS2{F0: 651, F1: 937}
                                                                                var v147 int16 = 1000
                                                                                fmt.Println(v1, v142, v5)
                                                                                _ = v146
                                                                                _ = v147
                                                                            }
                                                                            if ((391 != 1) || true) {
                                                                                v143 = v143
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = "a"
                                                                                }
                                                                                var v148 uint64 = 42
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = ((v71 + v71) + v71)
                                                                                }
                                                                                var v149 int8 = 0
                                                                                _ = v148
                                                                                _ = v149
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v141
                                                                    _ = v142
                                                                    _ = v143
                                                                } else {
                                                                    switch (v5 + 2147483647) {
                                                                    case 42:
                                                                        if (!(237 > (v5 - (v5 & fi20)))) {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        } else {
                                                                            switch fi69 {
                                                                            case 4294967295:
                                                                                v150 := GenS1{F0: uint32(v1), F1: 42, F2: false}
                                                                                _ = v150
                                                                            case 105:
                                                                                v151 := GenS12{F0: (v1 - (v3 >> ((uint(v1)) & 63)))}
                                                                                _ = v151
                                                                            default:
                                                                                switch fi20 {
                                                                                case 127:
                                                                                    v152 := GenS12{F0: v1}
                                                                                    _ = v152
                                                                                case 9223372036854775807:
                                                                                    if (100 < 42) {
                                                                                        switch 127 {
                                                                                        case 1:
                                                                                            var v153 string = v71
                                                                                            _ = v153
                                                                                        case 4294967295:
                                                                                            if true {
                                                                                                if (0.0 != 3.14159) {
                                                                                                    if (!true) {
                                                                                                        v5 = v5
                                                                                                        if ((v3 * v1) == (v3 >> ((255) & 63))) {
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = "hello"
                                                                                                            }
                                                                                                            v6 -= (1e10 / v6)
                                                                                                        }
                                                                                                        fmt.Println(fi69, v6, fi69)
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v6 = v6
                                                                                                            var v154 uint64 = 280
                                                                                                            if (!false) {
                                                                                                                v71 = v71
                                                                                                            } else {
                                                                                                                v3 = 1
                                                                                                                v6 += 0.0
                                                                                                            }
                                                                                                            _ = v154
                                                                                                        }
                                                                                                    }
                                                                                                    var v155 uint64 = 617
                                                                                                    _ = v155
                                                                                                } else {
                                                                                                    v156 := GenS11{F0: true, F1: "\n\t", F2: 1}
                                                                                                    var v157 int = (v4 + (v4 >> ((255) & 63)))
                                                                                                    switch ((v71 + v71) + ("" + "")) {
                                                                                                    case "a":
                                                                                                        v157 -= (v157 & (v157 / ((fi20) | 1)))
                                                                                                        v158 := GenS13{F0: 0}
                                                                                                        v159 := GenS8{F0: 190}
                                                                                                        _ = v158
                                                                                                        _ = v159
                                                                                                    case "":
                                                                                                        v160 := GenS14{F0: 432}
                                                                                                        _ = v160
                                                                                                    case "foo bar":
                                                                                                        var v161 float64 = 3.14159
                                                                                                        _ = v161
                                                                                                    default:
                                                                                                        fmt.Println(v70, v157)
                                                                                                    }
                                                                                                    _ = v156
                                                                                                    _ = v157
                                                                                                }
                                                                                                if false {
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                    v162 := GenS8{F0: -1}
                                                                                                    _ = v162
                                                                                                }
                                                                                            }
                                                                                        case 128:
                                                                                            var v163 int16 = 1
                                                                                            _ = v163
                                                                                        default:
                                                                                            v6 += -2.25
                                                                                        }
                                                                                        v164 := GenS1{F0: 0, F1: 0, F2: ((false || (4294967295 >= 0)) || (0 > uint(v1)))}
                                                                                        var v165 uint32 = 4294967295
                                                                                        _ = v164
                                                                                        _ = v165
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v137
                                                            }
                                                        }
                                                    }
                                                    if ((v4 >> ((42) & 63)) <= v4) {
                                                        v1 = 2
                                                    }
                                                } else {
                                                    v5 = fi20
                                                }
                                            }
                                        }
                                        _ = v71
                                    } else {
                                        if (0.1 != v6) {
                                            v5 = v4
                                        } else {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v166 := GenS3{F0: 2, F1: -1}
                                                _ = v166
                                            case "a":
                                                var v167 uint16 = 0
                                                _ = v167
                                            case "":
                                                if len(v2) > 0 {
                                                    v2[0] = ("" + "")
                                                }
                                            }
                                            if (false || true) {
                                                v4 = (v5 - v5)
                                                switch v4 {
                                                case -671:
                                                    v4 += 128
                                                    v3 += v3
                                                case 42:
                                                    v168 := GenS15{F0: 43, F1: 42, F2: -1}
                                                    _ = v168
                                                default:
                                                    v3 = v3
                                                }
                                                if true {
                                                    var v169 float32 = (v70 / 0.1)
                                                    if false {
                                                        switch (v4 + v5) {
                                                        case 1:
                                                            v1 -= (^(v3))
                                                        case -45:
                                                            fmt.Println(v6, v3)
                                                            v70 += 1.5
                                                        }
                                                    } else {
                                                        if true {
                                                            if (2 != 131) {
                                                                v70 += ((float32(v6) / ((v70 - 1.5) + v70)) * v70)
                                                            } else {
                                                                fmt.Println(v3)
                                                                switch "foo bar" {
                                                                case "a":
                                                                    switch (fi20 & 32768) {
                                                                    case 32767:
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            fmt.Println(fi69, v5, v5)
                                                                        default:
                                                                            if ((964 != 65535) || true) {
                                                                                v2 = append(v2, ("hello" + "a"))
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    v170 := GenS13{F0: 42}
                                                                                    _ = v170
                                                                                case "":
                                                                                    var v171 int16 = 32767
                                                                                    v172 := GenS1{F0: 42, F1: v171, F2: (887 == 42)}
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    _ = v171
                                                                                    _ = v172
                                                                                case "foo bar":
                                                                                    v173 := GenS16{F0: 0, F1: 42, F2: 650}
                                                                                    _ = v173
                                                                                }
                                                                            } else {
                                                                                var v174 uint64 = 1
                                                                                _ = v174
                                                                            }
                                                                        }
                                                                    case -1:
                                                                        var v175 int = fi20
                                                                        _ = v175
                                                                    default:
                                                                        var v176 int16 = 0
                                                                        _ = v176
                                                                    }
                                                                case "\n\t":
                                                                    if false {
                                                                        var v177 int32 = 1
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "\n\t"
                                                                        }
                                                                        _ = v177
                                                                    } else {
                                                                        v169 -= v169
                                                                        var v178 float32 = 1e10
                                                                        _ = v178
                                                                    }
                                                                case "foo bar":
                                                                    if false {
                                                                        switch (("a" + ("hello" + ("hello" + ("a" + ((("\n\t" + "\n\t") + "foo bar") + "\n\t"))))) + ("hello" + (("foo bar" + "hello") + ("\n\t" + "a")))) {
                                                                        case "\n\t":
                                                                            v179 := GenS4{F0: 255, F1: (v5 + (v5 % (((fi20 + v4)) | 1)))}
                                                                            _ = v179
                                                                        default:
                                                                            v2 = append(v2, ("" + "hello"))
                                                                        }
                                                                        v5 += 5
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = ""
                                                                        }
                                                                        v4 -= v5
                                                                    }
                                                                default:
                                                                    var v180 float32 = v70
                                                                    var v181 int = (v5 - (v4 % ((1) | 1)))
                                                                    v6 -= 0.0
                                                                    _ = v180
                                                                    _ = v181
                                                                }
                                                            }
                                                            if false {
                                                                v5 = v5
                                                                switch ("\n\t" + "hello") {
                                                                case "":
                                                                    var v182 uint32 = 0
                                                                    _ = v182
                                                                case "\n\t":
                                                                    var v183 bool = ((-32768 <= -1) || false)
                                                                    _ = v183
                                                                case "a":
                                                                    if false {
                                                                        v4 += (fi69 + v4)
                                                                    }
                                                                }
                                                                v6 += (1e10 / v6)
                                                            }
                                                            v1 += 128
                                                        }
                                                        var v184 int8 = 1
                                                        _ = v184
                                                    }
                                                    _ = v169
                                                } else {
                                                    fmt.Println(v6, fi69)
                                                }
                                            }
                                            v1 = (v3 & v3)
                                        }
                                        fmt.Println(v4)
                                    }
                                } else {
                                    var v185 float64 = v6
                                    v186 := GenS6{F0: 42}
                                    v6 += (-2.25 - ((float64(v70) + 3.14159) - -2.25))
                                    _ = v185
                                    _ = v186
                                }
                                if false {
                                    if (4294967295 <= 65535) {
                                        if (0 != 65) {
                                            fmt.Println(fi20, v1)
                                            fmt.Println(fi20, fi69)
                                        }
                                        var v187 bool = ((v3 << ((498) & 63)) < (v3 ^ 62))
                                        switch "a" {
                                        case "\n\t":
                                            var v188 uint = 447
                                            _ = v188
                                        case "":
                                            v2 = append(v2, "")
                                        }
                                        _ = v187
                                    }
                                    fmt.Println(v70)
                                }
                                _ = v70
                            }
                            var v189 string = ("a" + "hello")
                            _ = fi69
                            _ = v189
                        }
                    }
                    v5 -= 1
                    for fi190 := 0; fi190 < 3; fi190++ {
                        var v191 float64 = v6
                        var v192 uint8 = 42
                        _ = fi190
                        _ = v191
                        _ = v192
                    }
                } else {
                    v193 := GenS5{F0: (v1 | 1), F1: 9223372036854775807}
                    for fi194 := 0; fi194 < 5; fi194++ {
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        _ = fi194
                    }
                    _ = v193
                }
                switch fi20 {
                case -820:
                    fmt.Println(v3, fi20)
                case -1:
                    fmt.Println(v4)
                }
                var v195 int8 = 127
                _ = v195
            } else {
                for fi196 := 0; fi196 < 4; fi196++ {
                    v6 = (((((float64(((1.5 * -2.25) + float32(v6))) - (v6 + v6)) + 1.5) * ((-2.25 - (v6 + -2.25)) - v6)) - float64((1e10 / 0.1))) - 1.5)
                    fmt.Println(v6, v6)
                    _ = fi196
                }
            }
            _ = fi20
        }
        var v197 uint32 = 1
        v4 += v4
        _ = v197
    }
    var v198 int = (v4 & v5)
    v6 += v6
    for fi199 := 0; fi199 < 14; fi199++ {
        if (0 != 2147483647) {
            v5 = fi199
        } else {
            if (0 < 42) {
                fmt.Println(v1)
            }
        }
        switch (v5 ^ v4) {
        case 2147483647:
            for fi200 := 0; fi200 < 18; fi200++ {
                if true {
                    v5 = v198
                    v4 = -62
                    switch "" {
                    case "hello":
                        switch -515 {
                        case 1:
                            var v201 uint32 = 1
                            _ = v201
                        case 9223372036854775807:
                            if (39 >= 601) {
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                case "a":
                                    v6 = v6
                                case "hello":
                                    v4 = (v4 | v4)
                                }
                            } else {
                                var v202 uint = 276
                                switch "hello" {
                                case "\n\t":
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                }
                                switch "" {
                                case "hello":
                                    var v203 uint = v202
                                    _ = v203
                                case "":
                                    var v204 uint64 = 1
                                    _ = v204
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v1 -= v3
                                _ = v202
                            }
                        case 4294967295:
                            var v205 uint32 = 2
                            _ = v205
                        default:
                            if (127 == -128) {
                                v198 += v198
                                var v206 string = "hello"
                                v207 := GenS17{F0: false, F1: -1}
                                _ = v206
                                _ = v207
                            } else {
                                v2 = append(v2, "\n\t")
                                v3 += (v3 << ((42) & 63))
                            }
                        }
                    case "\n\t":
                        switch (-(v198)) {
                        case 9223372036854775807:
                            v2 = append(v2, "a")
                        case 452:
                            var v208 int = (v4 << ((18446744073709551615) & 63))
                            switch (-(v5)) {
                            case 829:
                                v209 := GenS8{F0: 1000}
                                _ = v209
                            case 0:
                                if (!true) {
                                    var v210 int64 = 9223372036854775807
                                    var v211 string = ("foo bar" + "foo bar")
                                    _ = v210
                                    _ = v211
                                }
                            default:
                                v2 = append(v2, ("hello" + ""))
                            }
                            _ = v208
                        case -424:
                            fmt.Println(v6, v4)
                        default:
                            v6 -= v6
                        }
                    case "a":
                        fmt.Println(v198, v5)
                    }
                    v5 += fi200
                    v198 = 9223372036854775807
                    var v212 uint = 495
                    var v213 int32 = 1
                    v214 := GenS9{F0: v213}
                    _ = v212
                    _ = v213
                    _ = v214
                } else {
                    switch ("" + "a") {
                    case "a":
                        switch ("hello" + ((("a" + "a") + ("foo bar" + "hello")) + (("\n\t" + "\n\t") + "a"))) {
                        case "hello":
                            v1 = 0
                        case "foo bar":
                            var v215 int = v5
                            _ = v215
                        case "":
                            if false {
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        v4 += fi199
                                    }
                                    var v216 int16 = 225
                                    _ = v216
                                } else {
                                    if true {
                                        v198 = fi199
                                        var v217 int32 = 2147483647
                                        v3 = 1
                                        _ = v217
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    v198 = (v4 >> ((470) & 63))
                                }
                                fmt.Println(v198, fi199, fi199)
                                fmt.Println(v6, fi199, v1)
                            }
                        default:
                            switch ("foo bar" + ((("a" + ((("a" + "foo bar") + "") + "a")) + ((("a" + "a") + ((("" + "\n\t") + "hello") + "a")) + "foo bar")) + "a")) {
                            case "hello":
                                if (9223372036854775807 <= 1) {
                                    var v218 int64 = 9223372036854775807
                                    var v219 int = v4
                                    v218 = v218
                                    _ = v218
                                    _ = v219
                                } else {
                                    v220 := GenS18{F0: 4294967295, F1: 18446744073709551615, F2: ("foo bar" + (("hello" + "") + ("a" + "\n\t")))}
                                    _ = v220
                                }
                                if ("" < (("a" + "foo bar") + ("" + ""))) {
                                    switch 1 {
                                    case 41:
                                        v221 := GenS19{F0: 2, F1: 2147483648, F2: -255}
                                        _ = v221
                                    default:
                                        v222 := GenS17{F0: (int(v3) <= fi200), F1: -568}
                                        _ = v222
                                    }
                                    v223 := GenS20{F0: 212, F1: 4294967295, F2: 0}
                                    _ = v223
                                }
                            }
                        }
                    default:
                        var v224 uint = 65535
                        _ = v224
                    }
                    switch "a" {
                    case "a":
                        v198 = (^(fi199))
                    case "\n\t":
                        v225 := GenS6{F0: 42}
                        _ = v225
                    case "":
                        v226 := GenS17{F0: true, F1: 529}
                        _ = v226
                    default:
                        switch (v5 & (^(v198))) {
                        case -2147483648:
                            var v227 int32 = -2147483648
                            _ = v227
                        case -370:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            fmt.Println(v6)
                        }
                    }
                }
                _ = fi200
            }
        default:
            if true {
                var v228 uint16 = 255
                _ = v228
            }
        }
        _ = fi199
    }
    switch v4 {
    case 127:
        for fi229 := 0; fi229 < 12; fi229++ {
            var v230 uint = 18446744073709551615
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi229
            _ = v230
        }
    case -2147483648:
        if false {
            var v231 int64 = 42
            _ = v231
        }
    case 42:
        v2 = append(v2, "foo bar")
    default:
        var v232 uint8 = 127
        _ = v232
    }
    for fi233 := 0; fi233 < 20; fi233++ {
        var v234 uint = uint(v1)
        _ = len(v2)
        if false {
            fmt.Println(v4, v6, v4)
            if false {
                var v235 int32 = 0
                switch "hello" {
                case "a":
                    v2 = append(v2, "")
                }
                if false {
                    v234 -= v234
                    for fi236 := 0; fi236 < 12; fi236++ {
                        if false {
                            if ((float64(float32(v6)) + (float64(3.14159) * ((v6 + float64(-2.25)) - v6))) == float64(1.5)) {
                                switch "a" {
                                case "a":
                                    v234 += 255
                                default:
                                    var v237 string = "a"
                                    _ = v237
                                }
                                if true {
                                    v2 = append(v2, ("a" + "\n\t"))
                                    v6 += (v6 * ((0.1 + 0.1) / (v6 + 3.14159)))
                                } else {
                                    fmt.Println(v4, fi236)
                                    v238 := GenS16{F0: 65535, F1: 1, F2: v234}
                                    var v239 uint8 = v1
                                    _ = v238
                                    _ = v239
                                }
                            }
                        }
                        if (((fi236 >> (((v234 >> (((v234 + (v234 << ((255) & 63)))) & 63))) & 63)) == v198) || true) {
                            if len(v2) > 0 {
                                v2[0] = (("a" + ("\n\t" + ("" + ("hello" + ((("" + ("hello" + "\n\t")) + (((("a" + "\n\t") + "a") + "\n\t") + "")) + ((((("\n\t" + "a") + "a") + "hello") + "") + "hello")))))) + ((("\n\t" + "hello") + "foo bar") + ("" + ("hello" + ("foo bar" + ((((("foo bar" + (("" + ("" + "a")) + "hello")) + "\n\t") + "") + "a") + "hello"))))))
                            }
                            v2 = append(v2, "hello")
                            switch -2147483648 {
                            case 4294967295:
                                fmt.Println(v5, v3, v234)
                            default:
                                if (!true) {
                                    switch (v4 | (^(v198))) {
                                    case 32768:
                                        var v240 uint16 = 0
                                        _ = v240
                                    case -61:
                                        var v241 float64 = 0.0
                                        _ = v241
                                    default:
                                        v198 = (fi233 - 4294967295)
                                    }
                                } else {
                                    if false {
                                        v242 := GenS21{F0: -742, F1: true}
                                        var v243 uint = v234
                                        _ = v242
                                        _ = v243
                                    } else {
                                        var v244 uint32 = 2147483648
                                        fmt.Println(v6, v1)
                                        v245 := GenS22{F0: 410}
                                        _ = v244
                                        _ = v245
                                    }
                                    v4 += (fi233 >> ((v234) & 63))
                                    switch (("foo bar" + "hello") + "\n\t") {
                                    case "\n\t":
                                        var v246 uint16 = 1
                                        _ = v246
                                    case "hello":
                                        v247 := GenS14{F0: 1}
                                        _ = v247
                                    default:
                                        v3 -= 255
                                    }
                                    v248 := GenS23{F0: 2}
                                    var v249 string = (("" + "hello") + "")
                                    _ = v248
                                    _ = v249
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            switch fi233 {
                            case -205:
                                if (0 < 2) {
                                    v4 = (^(fi233))
                                } else {
                                    v198 -= 4294967295
                                }
                            case 42:
                                fmt.Println(v3)
                            default:
                                v198 = v5
                            }
                            _ = len(v2)
                            var v250 int = (v198 % ((v4) | 1))
                            _ = v250
                        }
                        v251 := GenS24{F0: "hello", F1: ("" + ("" + "\n\t")), F2: 42}
                        v252 := GenS25{F0: v4, F1: 42, F2: "foo bar"}
                        v2 = append(v2, "foo bar")
                        _ = fi236
                        _ = v251
                        _ = v252
                    }
                }
                _ = v235
            }
            v253 := GenS26{F0: 163, F1: 42}
            _ = v253
        }
        _ = fi233
        _ = v234
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v6)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v1 = 1
    switch 0 {
    case -1:
        for fi254 := 0; fi254 < 18; fi254++ {
            for fi255 := 0; fi255 < 19; fi255++ {
                switch "\n\t" {
                case "\n\t":
                    var v256 uint = 4294967295
                    _ = v256
                default:
                    v257 := GenS2{F0: v5, F1: 65535}
                    _ = v257
                }
                var v258 uint = 18446744073709551615
                _ = fi255
                _ = v258
            }
            v259 := GenS25{F0: -9223372036854775808, F1: 18446744073709551615, F2: "hello"}
            fmt.Println(v3)
            _ = fi254
            _ = v259
        }
    }
    var v260 int32 = 0
    var v261 string = "foo bar"
    v262 := GenS12{F0: 42}
    v263 := GenS27{F0: ((v198 - (v4 | v5)) != (v4 % (((v198 | v5)) | 1)))}
    v264 := GenS28{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v198:", v198)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v260:", v260)
    fmt.Println("v261:", v261)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v198
    _ = v2
    _ = v260
    _ = v261
    _ = v262
    _ = v263
    _ = v264
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
