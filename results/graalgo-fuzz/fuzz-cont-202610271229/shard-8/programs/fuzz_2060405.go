package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint16; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint32; F2 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32; F1 bool; F2 bool }
type GenS10 struct { F0 uint; F1 bool; F2 uint8 }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int32; F1 uint32; F2 int64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint16; F1 int }
func genFn0() int64 {
    switch "hello" {
    case "a":
        if true {
            for fi5 := 0; fi5 < 14; fi5++ {
                var v6 bool = false
                v7 := []int{}
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 float32 = -2.25
            v9 := []int{-716, 2147483647}
            _ = v8
            _ = v9
        }
    case "hello":
        var v10 uint32 = 87
        _ = v10
    default:
        var v11 uint8 = 12
        _ = v11
    }
    var v12 float32 = 3.14159
    var v13 uint32 = 87
    if (v13 < (v13 - 2147483648)) {
        v13 -= 0
    } else {
        v14 := []int{4294967295, 128, -407}
        if (false || false) {
            var v15 bool = (127 != -128)
            v16 := GenS1{F0: v13}
            switch 4294967295 {
            case 9223372036854775807:
                switch 976 {
                case -9223372036854775808:
                    var v17 int = 127
                    _ = v17
                case 2147483647:
                    v12 += v12
                case 1:
                    if len(v14) > 0 {
                        v14[0] = 2147483647
                    }
                }
            }
            _ = v15
            _ = v16
        } else {
            for fi18 := 0; fi18 < 20; fi18++ {
                v12 -= v12
                switch "\n\t" {
                case "a":
                    v19 := GenS1{F0: 65535}
                    _ = v19
                case "foo bar":
                    if (fi18 == fi18) {
                        for fi20 := 0; fi20 < 12; fi20++ {
                            v21 := GenS2{F0: 32767, F1: uint64(v13)}
                            if (604 <= 1) {
                                v22 := GenS2{F0: 1, F1: 1}
                                fmt.Println(fi20, fi18)
                                if (42 == 2) {
                                    var v23 uint8 = 1
                                    var v24 uint8 = 0
                                    fmt.Println(v24, v12, fi20)
                                    _ = v23
                                    _ = v24
                                }
                                v12 -= v12
                                _ = v22
                            } else {
                                v14 = append(v14, (fi20 + 128))
                            }
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            v25 := GenS1{F0: 42}
                            switch ((("\n\t" + "foo bar") + ("a" + ((("hello" + "a") + "hello") + "hello"))) + (("hello" + "hello") + (("foo bar" + "hello") + "\n\t"))) {
                            case "":
                                if (false && (true || false)) {
                                    switch (fi20 - (fi20 & (-(fi20)))) {
                                    case -468:
                                        var v26 uint8 = 33
                                        _ = v26
                                    case -1:
                                        var v27 uint16 = 65535
                                        v28 := GenS3{F0: 0, F1: (v27 / (((v27 - (^(v27)))) | 1)), F2: 726}
                                        if false {
                                            v12 = v12
                                            if false {
                                                var v29 uint32 = 2147483648
                                                v30 := GenS4{F0: fi20}
                                                v31 := GenS2{F0: -1, F1: 0}
                                                fmt.Println(fi18, v27)
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                        } else {
                                            var v32 int64 = 0
                                            if false {
                                                v32 = -428
                                                if ((v13 / (((^(v13))) | 1)) > v13) {
                                                    v33 := GenS4{F0: fi20}
                                                    v27 -= (v27 & 1)
                                                    _ = v33
                                                } else {
                                                    if false {
                                                        v32 += (v32 << ((0) & 63))
                                                        fmt.Println(v27, v27)
                                                    } else {
                                                        fmt.Println(v32)
                                                        var v34 float32 = (v12 / float32(1.5))
                                                        if false {
                                                            var v35 uint8 = 153
                                                            v34 -= float32(((3.14159 + 3.14159) - ((3.14159 * 0.0) * -2.25)))
                                                            if true {
                                                                switch int(v35) {
                                                                case 42:
                                                                    v13 = (v13 % ((v13) | 1))
                                                                case 2147483647:
                                                                    fmt.Println(v27, fi18)
                                                                    if true {
                                                                        var v36 float32 = (1.5 + ((v34 + v12) + v34))
                                                                        fmt.Println(v27, v13, v35)
                                                                        _ = v36
                                                                    }
                                                                default:
                                                                    v35 = (v35 << ((uint(v13)) & 63))
                                                                }
                                                                fmt.Println(v35, v34)
                                                                if (861 >= 958) {
                                                                    v37 := GenS3{F0: -273, F1: (v27 + v27), F2: -32768}
                                                                    fmt.Println(v35)
                                                                    switch (fi20 << ((631) & 63)) {
                                                                    case -2147483648:
                                                                        var v38 uint8 = v35
                                                                        _ = v38
                                                                    case 32768:
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v14 = append(v14, fi20)
                                                                        }
                                                                    }
                                                                    _ = v37
                                                                } else {
                                                                    v13 = v13
                                                                }
                                                            } else {
                                                                var v39 uint8 = v35
                                                                if ((v35 ^ 141) <= (v39 | (^(v39)))) {
                                                                    v34 -= v34
                                                                    var v40 int8 = 99
                                                                    _ = v40
                                                                }
                                                                _ = v39
                                                            }
                                                            _ = v35
                                                        }
                                                        _ = v34
                                                    }
                                                    if (2147483647 != 2147483647) {
                                                        var v41 int8 = -128
                                                        v14 = append(v14, 700)
                                                        if true {
                                                            v13 = (v13 % (((v13 / (((v13 | v13)) | 1))) | 1))
                                                        } else {
                                                            if (720 >= 97) {
                                                                if (v13 == v13) {
                                                                    if false {
                                                                        _ = len(v14)
                                                                    }
                                                                } else {
                                                                    v41 = v41
                                                                    v42 := GenS4{F0: (fi18 & fi18)}
                                                                    v41 -= v41
                                                                    v41 = (v41 ^ 127)
                                                                    v12 += (v12 + v12)
                                                                    if true {
                                                                        var v43 int16 = 0
                                                                        var v44 uint64 = 0
                                                                        _ = v43
                                                                        _ = v44
                                                                    } else {
                                                                        if (!false) {
                                                                            switch fi20 {
                                                                            case 32767:
                                                                                fmt.Println(v32, v41)
                                                                            case -2147483648:
                                                                                v27 -= 398
                                                                            }
                                                                            switch fi18 {
                                                                            case 547:
                                                                                v45 := GenS3{F0: v32, F1: (^(v27)), F2: 193}
                                                                                _ = v45
                                                                            case 79:
                                                                                var v46 float64 = (1e10 - 0.0)
                                                                                _ = v46
                                                                            }
                                                                            if len(v14) >= 2 {
                                                                                v14 = v14[1:len(v14)]
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v42
                                                                }
                                                            } else {
                                                                switch (^(fi20)) {
                                                                case -9223372036854775808:
                                                                    v32 = v32
                                                                case 4294967295:
                                                                    var v47 uint8 = 1
                                                                    _ = v47
                                                                case -435:
                                                                    var v48 int16 = 372
                                                                    _ = v48
                                                                default:
                                                                    fmt.Println(v32)
                                                                }
                                                                v49 := GenS3{F0: (v32 * v32), F1: 2, F2: -32768}
                                                                if (290 != 18446744073709551615) {
                                                                    switch (fi18 | -435) {
                                                                    case 128:
                                                                        fmt.Println(v27)
                                                                    case 765:
                                                                        switch "" {
                                                                        case "hello":
                                                                            fmt.Println(v27, v32, v32)
                                                                        case "":
                                                                            fmt.Println(v12)
                                                                        default:
                                                                            v14 = append(v14, (fi18 | fi20))
                                                                        }
                                                                    case 4294967295:
                                                                        var v50 uint64 = 517
                                                                        _ = v50
                                                                    default:
                                                                        v12 = v12
                                                                        _ = len(v14)
                                                                        v27 = 32767
                                                                    }
                                                                    v12 -= (float32((1e10 / 3.14159)) / 0.1)
                                                                } else {
                                                                    fmt.Println(fi20, v41, fi20)
                                                                }
                                                                v51 := GenS3{F0: v32, F1: (v27 & 2), F2: -903}
                                                                v14 = append(v14, (fi20 << ((0) & 63)))
                                                                _ = v49
                                                                _ = v51
                                                            }
                                                            _ = len(v14)
                                                        }
                                                        _ = v41
                                                    }
                                                    if true {
                                                        v52 := GenS1{F0: v13}
                                                        _ = v52
                                                    }
                                                }
                                                fmt.Println(fi18, fi18, fi20)
                                            } else {
                                                if true {
                                                    if true {
                                                        var v53 string = (("" + "foo bar") + "\n\t")
                                                        fmt.Println(v27, v53, fi18)
                                                        fmt.Println(fi18)
                                                        switch ("\n\t" + ("foo bar" + v53)) {
                                                        case "foo bar":
                                                            var v54 uint8 = 128
                                                            _ = v54
                                                        default:
                                                            switch fi18 {
                                                            case 128:
                                                                fmt.Println(v13, v13, v12)
                                                            }
                                                        }
                                                        _ = v53
                                                    } else {
                                                        var v55 int32 = -2147483648
                                                        v55 -= (v55 | (v55 / (((v55 / ((v55) | 1))) | 1)))
                                                        if (true || true) {
                                                            var v56 uint = 1
                                                            var v57 int64 = 0
                                                            v12 = float32(-2.25)
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        _ = v55
                                                    }
                                                    fmt.Println(v12, v27)
                                                }
                                                if false {
                                                    if len(v14) > 0 {
                                                        v14[0] = (fi20 + (fi20 * fi20))
                                                    }
                                                } else {
                                                    fmt.Println(v12, v13, v13)
                                                    if (641 <= 889) {
                                                        v13 += v13
                                                        if true {
                                                            v58 := GenS3{F0: v32, F1: v27, F2: 533}
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v59 := GenS1{F0: 2147483648}
                                                                _ = v59
                                                            case "a":
                                                                fmt.Println(fi18)
                                                            case "\n\t":
                                                                fmt.Println(v27)
                                                                if false {
                                                                    var v60 int8 = -20
                                                                    if (-1 == -1) {
                                                                        fmt.Println(v27, v13, fi20)
                                                                        v13 += (v13 * 68)
                                                                        v12 = float32(-2.25)
                                                                        v61 := GenS5{F0: 18446744073709551615}
                                                                        _ = v61
                                                                    } else {
                                                                        v60 -= v60
                                                                    }
                                                                    _ = v60
                                                                }
                                                                var v62 uint8 = 239
                                                                _ = v62
                                                            }
                                                            v32 = v32
                                                            _ = v58
                                                        } else {
                                                            if len(v14) > 0 {
                                                                v14[0] = (fi20 - 276)
                                                            }
                                                            var v63 string = ""
                                                            fmt.Println(v63, v32, v63)
                                                            _ = v63
                                                        }
                                                        var v64 uint32 = (^(v13))
                                                        _ = v64
                                                    } else {
                                                        if false {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v65 uint64 = 0
                                                                _ = v65
                                                            case "foo bar":
                                                                var v66 float32 = (1.5 / (float32(float64(0.1)) / (v12 - v12)))
                                                                _ = v66
                                                            default:
                                                                switch 456 {
                                                                case -438:
                                                                    var v67 uint32 = (^(v13))
                                                                    _ = v67
                                                                case 32767:
                                                                    fmt.Println(fi18)
                                                                    var v68 int32 = 687
                                                                    if len(v14) > 0 {
                                                                        v14[0] = (fi20 << ((4294967295) & 63))
                                                                    }
                                                                    _ = v68
                                                                }
                                                                v12 = ((v12 / v12) * v12)
                                                            }
                                                            v12 += float32(1.5)
                                                        }
                                                        if (-633 < 528) {
                                                            fmt.Println(v32)
                                                            _ = len(v14)
                                                            v14 = append(v14, 9223372036854775807)
                                                        }
                                                    }
                                                    var v69 int8 = 1
                                                    _ = v69
                                                }
                                                if false {
                                                    v32 += 42
                                                }
                                                var v70 int16 = 1000
                                                v12 = (float32(1e10) * ((v12 * ((v12 - float32(float64(v12))) - v12)) * (float32(1e10) * 0.1)))
                                                var v71 uint16 = (^(v27))
                                                if true {
                                                    if false {
                                                        v72 := GenS2{F0: 474, F1: 380}
                                                        v14 = append(v14, fi18)
                                                        v73 := GenS6{F0: v70}
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        v74 := GenS6{F0: v70}
                                                        _ = v74
                                                    }
                                                    v32 = v32
                                                    v75 := GenS7{F0: "foo bar", F1: (v13 & 523), F2: uint32(v27)}
                                                    _ = v75
                                                }
                                                _ = v70
                                                _ = v71
                                            }
                                            v12 -= v12
                                            _ = v32
                                        }
                                        _ = v27
                                        _ = v28
                                    default:
                                        if (false || false) {
                                            var v76 float32 = v12
                                            v13 = v13
                                            _ = v76
                                        } else {
                                            if len(v14) > 0 {
                                                v14[0] = 681
                                            }
                                            if (2 < 747) {
                                                var v77 uint = 42
                                                switch fi20 {
                                                case 127:
                                                    v13 = uint32(v77)
                                                default:
                                                    switch "a" {
                                                    case "a":
                                                        var v78 uint64 = 1
                                                        switch (fi18 - fi20) {
                                                        case -9223372036854775808:
                                                            var v79 int16 = 0
                                                            _ = v79
                                                        default:
                                                            var v80 float32 = (v12 * v12)
                                                            _ = v80
                                                        }
                                                        if false {
                                                            if false {
                                                                var v81 int = (fi20 / ((fi20) | 1))
                                                                if len(v14) > 0 {
                                                                    v14[0] = v81
                                                                }
                                                                var v82 int8 = 1
                                                                v78 -= v78
                                                                if false {
                                                                    var v83 int = (^(fi18))
                                                                    if (true && (((^(v82)) == v82) || (!false))) {
                                                                        v13 = (^(v13))
                                                                        v13 = (^(v13))
                                                                    } else {
                                                                        v84 := GenS7{F0: "a", F1: v13, F2: (v13 >> (((^(v77))) & 63))}
                                                                        _ = v84
                                                                    }
                                                                    v82 -= (v82 % ((1) | 1))
                                                                    _ = v83
                                                                }
                                                                _ = v81
                                                                _ = v82
                                                            } else {
                                                                v77 += (v77 / (((v77 + 255)) | 1))
                                                                var v85 float64 = (1e10 / -2.25)
                                                                switch fi18 {
                                                                case -2147483648:
                                                                    var v86 int32 = int32(v78)
                                                                    _ = v86
                                                                case 32768:
                                                                    switch int(v78) {
                                                                    case 9223372036854775807:
                                                                        v87 := GenS7{F0: "", F1: (v13 / ((v13) | 1)), F2: 42}
                                                                        _ = v87
                                                                    case 32768:
                                                                        var v88 int = 963
                                                                        _ = v88
                                                                    case 4294967295:
                                                                        var v89 int32 = 42
                                                                        _ = v89
                                                                    default:
                                                                        v13 += (v13 >> (((v77 << ((v77) & 63))) & 63))
                                                                    }
                                                                case 32767:
                                                                    v14 = append(v14, int(v78))
                                                                default:
                                                                    if true {
                                                                        var v90 string = ("\n\t" + ((("a" + "hello") + "hello") + "foo bar"))
                                                                        v85 -= v85
                                                                        if (!false) {
                                                                            v14 = append(v14, (^(fi18)))
                                                                        } else {
                                                                            _ = len(v14)
                                                                            var v91 int = (fi20 / ((fi20) | 1))
                                                                            var v92 float64 = v85
                                                                            _ = v91
                                                                            _ = v92
                                                                        }
                                                                        _ = v90
                                                                    } else {
                                                                        var v93 uint = 615
                                                                        _ = v93
                                                                    }
                                                                }
                                                                _ = v85
                                                            }
                                                            switch 2147483647 {
                                                            case 4294967295:
                                                                v94 := GenS2{F0: 1, F1: v78}
                                                                _ = v94
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v78 -= (v78 | (v78 & (^(v78))))
                                                            case "a":
                                                                v12 = v12
                                                            }
                                                        } else {
                                                            if true {
                                                                if len(v14) > 0 {
                                                                    v14[0] = fi18
                                                                }
                                                                if true {
                                                                    if false {
                                                                        v95 := GenS1{F0: uint32(v77)}
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if false {
                                                                                var v96 bool = true
                                                                                if false {
                                                                                    var v97 int8 = 0
                                                                                    var v98 int64 = int64(v97)
                                                                                    _ = v97
                                                                                    _ = v98
                                                                                } else {
                                                                                    v14 = append(v14, fi18)
                                                                                    fmt.Println(v96, v78, v77)
                                                                                }
                                                                                _ = v96
                                                                            } else {
                                                                                v13 -= v13
                                                                                var v99 int8 = -1
                                                                                _ = v99
                                                                            }
                                                                            v12 += v12
                                                                            if (-677 == 1) {
                                                                                switch 1 {
                                                                                case -1:
                                                                                    switch fi20 {
                                                                                    case -2147483648:
                                                                                        if ((v77 - 42) != (v77 | (v77 % ((v77) | 1)))) {
                                                                                            fmt.Println(v13, v78, fi18)
                                                                                            var v100 int = fi18
                                                                                            _ = v100
                                                                                        } else {
                                                                                            var v101 bool = true
                                                                                            if len(v14) > 0 {
                                                                                                v14[0] = fi18
                                                                                            }
                                                                                            _ = v101
                                                                                        }
                                                                                    }
                                                                                case 769:
                                                                                    v13 -= (v13 + v13)
                                                                                case 295:
                                                                                    fmt.Println(fi18, v78)
                                                                                default:
                                                                                    if false {
                                                                                        v78 -= v78
                                                                                        var v102 uint = v77
                                                                                        _ = len(v14)
                                                                                        _ = v102
                                                                                    }
                                                                                }
                                                                                v103 := GenS2{F0: 32767, F1: 0}
                                                                                v13 -= (v13 % ((v13) | 1))
                                                                                _ = v103
                                                                            }
                                                                        case "\n\t":
                                                                            if (!(213 == 0)) {
                                                                                switch (fi20 + 32767) {
                                                                                case 1:
                                                                                    var v104 float64 = 0.1
                                                                                    _ = v104
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(fi18, v13, fi18)
                                                                                case -282:
                                                                                    if true {
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            v105 := GenS4{F0: fi20}
                                                                                            _ = v105
                                                                                        }
                                                                                    } else {
                                                                                        v14 = append(v14, 127)
                                                                                        var v106 string = "\n\t"
                                                                                        _ = v106
                                                                                    }
                                                                                default:
                                                                                    if (!(-2.25 != ((((0.0 * (((1e10 + 0.0) + 1e10) / 1.5)) - 3.14159) - 3.14159) + 3.14159))) {
                                                                                        fmt.Println(fi20)
                                                                                        if true {
                                                                                            v12 += v12
                                                                                        } else {
                                                                                            v13 = 1
                                                                                            v13 = (v13 >> (((^(v77))) & 63))
                                                                                            var v107 int = fi20
                                                                                            _ = v107
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v12, v12, v13)
                                                                                    }
                                                                                }
                                                                                switch ("" + (("a" + "") + ("a" + "foo bar"))) {
                                                                                case "\n\t":
                                                                                    fmt.Println(v77, v12, fi18)
                                                                                case "":
                                                                                    var v108 int8 = 1
                                                                                    _ = v108
                                                                                default:
                                                                                    var v109 uint8 = 2
                                                                                    _ = v109
                                                                                }
                                                                            }
                                                                        default:
                                                                            v78 -= v78
                                                                        }
                                                                        v110 := GenS1{F0: v13}
                                                                        if len(v14) >= 2 {
                                                                            v14 = v14[1:len(v14)]
                                                                        }
                                                                        _ = v110
                                                                        _ = v95
                                                                    } else {
                                                                        if len(v14) > 0 {
                                                                            v14[0] = -2147483648
                                                                        }
                                                                        v12 -= (0.1 - v12)
                                                                        v111 := GenS1{F0: v13}
                                                                        fmt.Println(v78)
                                                                        _ = v111
                                                                    }
                                                                    v13 -= 4294967295
                                                                    v13 = v13
                                                                } else {
                                                                    _ = len(v14)
                                                                }
                                                            } else {
                                                                switch ("\n\t" + (("foo bar" + (("foo bar" + ("\n\t" + ("" + ""))) + "\n\t")) + "\n\t")) {
                                                                case "a":
                                                                    if len(v14) > 0 {
                                                                        v14[0] = (fi20 + fi18)
                                                                    }
                                                                case "\n\t":
                                                                    _ = len(v14)
                                                                case "hello":
                                                                    if false {
                                                                        v112 := GenS6{F0: 519}
                                                                        var v113 string = "foo bar"
                                                                        var v114 uint = v77
                                                                        _ = v112
                                                                        _ = v113
                                                                        _ = v114
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(v78, v78, fi18)
                                                                            var v115 int64 = 720
                                                                            if false {
                                                                                _ = len(v14)
                                                                            }
                                                                            _ = v115
                                                                        } else {
                                                                            if false {
                                                                                var v116 uint8 = 147
                                                                                if ((false && false) && ((!false) || false)) {
                                                                                    v116 = 2
                                                                                }
                                                                                _ = v116
                                                                            } else {
                                                                                if (false && (116 != 42)) {
                                                                                    v117 := GenS8{F0: 0}
                                                                                    v118 := GenS3{F0: 0, F1: 0, F2: -1}
                                                                                    v78 += v78
                                                                                    _ = v117
                                                                                    _ = v118
                                                                                }
                                                                                if len(v14) > 0 {
                                                                                    v14[0] = fi18
                                                                                }
                                                                                if ((!false) || true) {
                                                                                    var v119 string = ("" + "a")
                                                                                    if len(v14) > 0 {
                                                                                        v14[0] = fi18
                                                                                    }
                                                                                    fmt.Println(v12, v119)
                                                                                    switch ((v119 + "\n\t") + "") {
                                                                                    case "":
                                                                                        v120 := GenS9{F0: 0, F1: false, F2: true}
                                                                                        _ = v120
                                                                                    default:
                                                                                        if (true && (0.0 >= 0.0)) {
                                                                                            switch 0 {
                                                                                            case -1:
                                                                                                v121 := GenS10{F0: v77, F1: (-1 == 100), F2: 128}
                                                                                                v119 = v119
                                                                                                _ = v121
                                                                                            default:
                                                                                                v14 = append(v14, (fi20 / ((9223372036854775807) | 1)))
                                                                                            }
                                                                                            _ = len(v14)
                                                                                        }
                                                                                    }
                                                                                    _ = len(v14)
                                                                                    _ = v119
                                                                                } else {
                                                                                    if (42 <= -1) {
                                                                                        if len(v14) >= 2 {
                                                                                            v14 = v14[1:len(v14)]
                                                                                        }
                                                                                        v13 = 877
                                                                                    } else {
                                                                                        v77 += (v77 + v77)
                                                                                        if len(v14) > 0 {
                                                                                            v14[0] = -1
                                                                                        }
                                                                                        var v122 float64 = 0.1
                                                                                        _ = v122
                                                                                    }
                                                                                    fmt.Println(fi20, v78)
                                                                                }
                                                                            }
                                                                        }
                                                                        v12 += 1e10
                                                                    }
                                                                }
                                                            }
                                                            v77 += v77
                                                        }
                                                        switch fi20 {
                                                        case 127:
                                                            var v123 uint32 = uint32(v77)
                                                            _ = v123
                                                        case 2147483647:
                                                            v124 := GenS3{F0: 9223372036854775807, F1: 848, F2: 32767}
                                                            _ = v124
                                                        case -195:
                                                            fmt.Println(fi20, fi18, fi20)
                                                            if len(v14) >= 2 {
                                                                v14 = v14[1:len(v14)]
                                                            }
                                                            if true {
                                                                if false {
                                                                    fmt.Println(v13, v77, fi20)
                                                                    var v125 uint32 = (v13 << (((v77 << (((v77 % ((v77) | 1))) & 63))) & 63))
                                                                    _ = v125
                                                                } else {
                                                                    if (!(42 > 286)) {
                                                                        var v126 uint8 = 2
                                                                        if len(v14) > 0 {
                                                                            v14[0] = (^(fi20))
                                                                        }
                                                                        if len(v14) >= 2 {
                                                                            v14 = v14[1:len(v14)]
                                                                        }
                                                                        _ = v126
                                                                    } else {
                                                                        var v127 int32 = 42
                                                                        var v128 uint16 = 318
                                                                        v129 := GenS1{F0: 330}
                                                                        _ = v127
                                                                        _ = v128
                                                                        _ = v129
                                                                    }
                                                                    var v130 float64 = 1.5
                                                                    _ = v130
                                                                }
                                                                if len(v14) >= 2 {
                                                                    v14 = v14[1:len(v14)]
                                                                }
                                                                v12 -= float32(0.1)
                                                                fmt.Println(fi18)
                                                            }
                                                        }
                                                        _ = v78
                                                    default:
                                                        var v131 float32 = (v12 + v12)
                                                        _ = v131
                                                    }
                                                    var v132 uint16 = 635
                                                    fmt.Println(fi18)
                                                    _ = v132
                                                }
                                                _ = v77
                                            } else {
                                                v133 := GenS11{F0: fi20, F1: -128}
                                                _ = len(v14)
                                                fmt.Println(fi20, v12)
                                                _ = v133
                                            }
                                            v13 = 852
                                        }
                                    }
                                    v134 := GenS5{F0: 583}
                                    _ = v134
                                }
                            case "\n\t":
                                var v135 int32 = -1
                                _ = v135
                            }
                            v136 := GenS12{F0: ("" + ("a" + "\n\t"))}
                            _ = fi20
                            _ = v136
                            _ = v21
                            _ = v25
                        }
                        var v137 float32 = v12
                        v138 := GenS2{F0: 1, F1: 18446744073709551615}
                        _ = v137
                        _ = v138
                    }
                }
                _ = fi18
            }
            var v139 int = -2147483648
            _ = v139
        }
        _ = v14
    }
    var v140 bool = false
    _ = v12
    _ = v13
    _ = v140
    return -1
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]bool)
    var v3 bool = ((65535 >= 32767) && true)
    var v4 uint8 = 127
    fmt.Println(v1, v1)
    v3 = (v4 > v4)
    v141 := GenS13{F0: -27, F1: 1, F2: (v1 + genFn0())}
    var v142 int16 = 0
    var v143 uint16 = 32767
    v144 := []int{-1, -1, 127}
    v4 = 255
    _ = len(v2)
    switch "hello" {
    case "\n\t":
        for fi145 := 0; fi145 < 3; fi145++ {
            if (!(!true)) {
                var v146 uint64 = 18446744073709551615
                _ = v146
            }
            v4 = (v4 << ((1) & 63))
            _ = fi145
        }
    case "foo bar":
        v3 = false
    case "a":
        switch 9223372036854775807 {
        case 1:
            fmt.Println(v142)
        case 127:
            if (v3 && v3) {
                if (true && false) {
                    v1 = genFn0()
                } else {
                    v144 = append(v144, 756)
                    switch 852 {
                    case 203:
                        v3 = v3
                    default:
                        v147 := GenS14{F0: 18446744073709551615}
                        _ = v147
                    }
                    v1 = 9223372036854775807
                }
            }
        case -9223372036854775808:
            for fi148 := 0; fi148 < 20; fi148++ {
                switch ("\n\t" + (("\n\t" + "a") + "")) {
                case "a":
                    var v149 uint = 0
                    _ = v149
                case "hello":
                    var v150 float64 = 0.1
                    _ = v150
                default:
                    if v3 {
                        if v3 {
                            delete(v2, -2147483648)
                            v1 = v1
                        }
                        v2[(fi148 - 42)] = (244 > 41)
                        for fi151 := 0; fi151 < 9; fi151++ {
                            v143 = (v143 + v143)
                            fmt.Println(fi151, v143)
                            var v152 string = "\n\t"
                            var v153 int = (fi151 - fi148)
                            _ = fi151
                            _ = v152
                            _ = v153
                        }
                    }
                }
                _ = fi148
            }
        }
    default:
        var v154 string = ("\n\t" + ("hello" + ("foo bar" + "a")))
        _ = v154
    }
    v143 -= v143
    var v155 float32 = 1.5
    switch "" {
    case "\n\t":
        var v156 int64 = (^(v1))
        _ = v156
    default:
        switch -1 {
        case 0:
            for fi157 := 0; fi157 < 2; fi157++ {
                fmt.Println(v155, v3)
                if false {
                    v3 = ((v142 * -1) >= 1)
                    v3 = true
                    for fi158 := 0; fi158 < 16; fi158++ {
                        var v159 float64 = -2.25
                        v4 += v4
                        _ = fi158
                        _ = v159
                    }
                } else {
                    fmt.Println(fi157, fi157, v3)
                    v143 = v143
                }
                _ = fi157
            }
        case -9223372036854775808:
            if (((v155 + -2.25) - (v155 - 0.1)) <= 1.5) {
                var v160 uint = uint(v4)
                fmt.Println(v155, v142)
                _ = v160
            } else {
                fmt.Println(v4)
                for fi161 := 0; fi161 < 2; fi161++ {
                    v162 := GenS6{F0: v142}
                    _ = fi161
                    _ = v162
                }
                v163 := GenS12{F0: (((("hello" + "") + (("foo bar" + ("foo bar" + "")) + "hello")) + "hello") + "\n\t")}
                _ = v163
            }
        }
    }
    fmt.Println(v143)
    switch 0 {
    case 2147483647:
        var v164 uint = 18446744073709551615
        _ = v164
    }
    if v3 {
        v1 = v1
        if (v4 != (v4 ^ 42)) {
            if v3 {
                for fi165 := 0; fi165 < 11; fi165++ {
                    var v166 string = "foo bar"
                    if (4294967295 == 12) {
                        if (v3 || v3) {
                            v167 := v2[(fi165 & fi165)]
                            _ = v167
                        }
                    }
                    for fi168 := 0; fi168 < 12; fi168++ {
                        v169 := GenS9{F0: int32(fi168), F1: (51 != 42), F2: ((((3.14159 + 0.1) * 3.14159) * 1.5) != (-2.25 * (1e10 - (0.0 + 3.14159))))}
                        if (!(!false)) {
                            v170 := GenS15{F0: ((v166 + "hello") != v166)}
                            _ = v170
                        } else {
                            var v171 uint = 408
                            var v172 int8 = 127
                            _ = v171
                            _ = v172
                        }
                        var v173 int16 = v142
                        _ = fi168
                        _ = v169
                        _ = v173
                    }
                    _ = fi165
                    _ = v166
                }
                var v174 int8 = -128
                if len(v144) >= 2 {
                    v144 = v144[1:len(v144)]
                }
                _ = v174
            } else {
                v1 = (v1 ^ v1)
                var v175 int16 = v142
                var v176 int16 = v142
                _ = v175
                _ = v176
            }
        }
        var v177 int64 = v1
        _ = v177
    } else {
        switch (("hello" + "\n\t") + "hello") {
        case "a":
            if v3 {
                v178 := GenS16{F0: v1}
                v2[2147483647] = false
                _ = v178
            } else {
                for fi179 := 0; fi179 < 1; fi179++ {
                    v1 = 0
                    var v180 int32 = -2147483648
                    fmt.Println(v143, v3)
                    _ = fi179
                    _ = v180
                }
            }
        case "foo bar":
            v181 := GenS17{F0: v143, F1: -747}
            _ = v181
        }
        v142 = -196
        fmt.Println(v142)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("slicelen v144:", len(v144), v144)
    fmt.Println("v155:", v155)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v141
    _ = v142
    _ = v143
    _ = v144
    _ = v155
    _ = v2
    _ = v3
    _ = v4
}
