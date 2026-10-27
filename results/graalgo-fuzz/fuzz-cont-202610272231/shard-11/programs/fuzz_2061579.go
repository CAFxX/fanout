package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 bool; F1 int16; F2 int8 }
type GenS4 struct { F0 uint16; F1 int64; F2 int32 }
type GenS5 struct { F0 string; F1 int16; F2 int32 }
type GenS6 struct { F0 int32; F1 int64; F2 int64 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8; F1 uint64; F2 bool }
type GenS10 struct { F0 bool; F1 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32; F1 int; F2 string }
type GenS13 struct { F0 int; F1 int8; F2 int8 }
type GenS14 struct { F0 int8; F1 uint32 }
type GenS15 struct { F0 int; F1 int64 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int8; F1 uint8 }
type GenS18 struct { F0 int16; F1 int32; F2 int64 }
type GenS19 struct { F0 bool; F1 int32 }
type GenS20 struct { F0 uint; F1 int }
func genFn0() int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int32 = -1
        _ = fi5
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = ("a" + "")
            for fi9 := 0; fi9 < 13; fi9++ {
                if false {
                    var v10 int64 = -9223372036854775808
                    v11 := GenS1{F0: 295}
                    switch v8 {
                    case "hello":
                        v8 = "a"
                    case "":
                        var v12 uint16 = 42
                        _ = v12
                    case "a":
                        fmt.Println(fi7, v8, fi9)
                    default:
                        v13 := GenS1{F0: 42}
                        _ = v13
                    }
                    _ = v10
                    _ = v11
                }
                var v14 string = "\n\t"
                var v15 float64 = float64((1e10 / (((0.0 / -2.25) / float32(1.5)) + -2.25)))
                _ = fi9
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(fi16, fi7, fi7)
                fmt.Println(v8, fi7, v8)
                var v17 bool = true
                var v18 uint8 = 235
                if (v18 >= 0) {
                    if (16 >= -259) {
                        v19 := []int{(fi16 % (((-(fi7))) | 1)), fi16, fi16}
                        if v17 {
                            fmt.Println(fi7, fi16, fi7)
                            switch fi7 {
                            case 9223372036854775807:
                                v20 := GenS2{F0: 65535, F1: 255}
                                if true {
                                    v21 := GenS3{F0: (-2.25 < (float64(0.1) - 0.1)), F1: 32767, F2: -122}
                                    _ = v21
                                }
                                switch fi16 {
                                case 32768:
                                    var v22 uint = 4294967295
                                    _ = v22
                                }
                                _ = v20
                            case -1:
                                var v23 uint = 18446744073709551615
                                _ = v23
                            case 128:
                                var v24 float32 = (1.5 + (1.5 + 0.1))
                                _ = v24
                            }
                        } else {
                            if (!v17) {
                                switch fi7 {
                                case -9223372036854775808:
                                    v18 -= (^(v18))
                                case 1:
                                    if true {
                                        if true {
                                            if v17 {
                                                v8 = (v8 + v8)
                                            } else {
                                                switch fi7 {
                                                case -903:
                                                    fmt.Println(v17, fi7)
                                                case -9223372036854775808:
                                                    fmt.Println(v18)
                                                case 1:
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                default:
                                                    v25 := GenS3{F0: v17, F1: 0, F2: 1}
                                                    _ = v25
                                                }
                                                v17 = v17
                                                switch fi7 {
                                                case 128:
                                                    var v26 float64 = 1.5
                                                    _ = v26
                                                case 4294967295:
                                                    v18 = v18
                                                case -878:
                                                    if v17 {
                                                        v27 := GenS1{F0: 4294967295}
                                                        var v28 uint16 = 65535
                                                        fmt.Println(v8)
                                                        _ = v27
                                                        _ = v28
                                                    } else {
                                                        fmt.Println(v18, fi16, v8)
                                                        var v29 uint16 = 375
                                                        switch fi16 {
                                                        case 9223372036854775807:
                                                            _ = len(v19)
                                                            v29 -= (v29 % (((v29 | 69)) | 1))
                                                        default:
                                                            switch v8 {
                                                            case "hello":
                                                                if (1 >= 42) {
                                                                    v30 := GenS2{F0: (v29 ^ v29), F1: (v29 & (v29 << ((42) & 63)))}
                                                                    v31 := GenS4{F0: (^(v29)), F1: -9223372036854775808, F2: 42}
                                                                    v32 := GenS5{F0: v8, F1: -1, F2: 0}
                                                                    v29 += v29
                                                                    _ = v30
                                                                    _ = v31
                                                                    _ = v32
                                                                } else {
                                                                    switch (fi7 >> ((2) & 63)) {
                                                                    case 32767:
                                                                        if v17 {
                                                                            _ = len(v19)
                                                                            var v33 int = (fi7 | fi7)
                                                                            _ = v33
                                                                        }
                                                                    case -968:
                                                                        v18 = v18
                                                                    }
                                                                    var v34 int8 = 42
                                                                    if (v18 < (v18 << ((4294967295) & 63))) {
                                                                        v35 := GenS4{F0: v29, F1: -1, F2: 724}
                                                                        if (9223372036854775807 < 0) {
                                                                            v36 := GenS6{F0: 1, F1: -667, F2: 0}
                                                                            v37 := GenS6{F0: -207, F1: -9223372036854775808, F2: 0}
                                                                            _ = v36
                                                                            _ = v37
                                                                        }
                                                                        if (2 == 45) {
                                                                            switch (^(fi16)) {
                                                                            case 0:
                                                                                if true {
                                                                                    fmt.Println(v17)
                                                                                } else {
                                                                                    if v17 {
                                                                                        fmt.Println(v8, fi16)
                                                                                        v38 := GenS5{F0: (((v8 + ("\n\t" + (("foo bar" + v8) + "hello"))) + (v8 + (v8 + v8))) + v8), F1: 42, F2: -1}
                                                                                        _ = v38
                                                                                    } else {
                                                                                        if len(v19) > 0 {
                                                                                            v19[0] = 32768
                                                                                        }
                                                                                        var v39 uint64 = 0
                                                                                        _ = v39
                                                                                    }
                                                                                    v40 := GenS7{F0: v34, F1: -887}
                                                                                    if (-9223372036854775808 > -1) {
                                                                                        var v41 int16 = -32768
                                                                                        if len(v19) > 0 {
                                                                                            v19[0] = 407
                                                                                        }
                                                                                        v42 := GenS8{F0: fi16}
                                                                                        _ = v41
                                                                                        _ = v42
                                                                                    } else {
                                                                                        v29 = (v29 << ((360) & 63))
                                                                                    }
                                                                                    _ = v40
                                                                                }
                                                                            default:
                                                                                v34 -= (v34 + 42)
                                                                                switch fi16 {
                                                                                case -528:
                                                                                    var v43 int32 = 0
                                                                                    _ = v43
                                                                                case 32768:
                                                                                    var v44 int64 = 1
                                                                                    _ = v44
                                                                                case 184:
                                                                                    if ((((fi16 | (fi16 - fi16)) >= (fi16 | 32768)) && true) || (-35 > 9223372036854775807)) {
                                                                                        switch ((v8 + (v8 + v8)) + v8) {
                                                                                        case "":
                                                                                            v45 := GenS6{F0: -680, F1: -1, F2: 803}
                                                                                            _ = v45
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            switch v8 {
                                                                                            case "hello":
                                                                                                fmt.Println(v17, fi16, fi7)
                                                                                                if (-715 != 42) {
                                                                                                    fmt.Println(fi16, v8, v17)
                                                                                                    v34 = -1
                                                                                                    v18 -= v18
                                                                                                } else {
                                                                                                    fmt.Println(v17, v17, fi16)
                                                                                                }
                                                                                            case "":
                                                                                                v19 = append(v19, -9223372036854775808)
                                                                                                var v46 int16 = int16(fi16)
                                                                                                if v17 {
                                                                                                    var v47 float64 = 1e10
                                                                                                    _ = v47
                                                                                                }
                                                                                                _ = v46
                                                                                            case "foo bar":
                                                                                                v29 -= 32767
                                                                                            default:
                                                                                                var v48 int16 = -902
                                                                                                if false {
                                                                                                    v49 := GenS1{F0: 255}
                                                                                                    var v50 float64 = 0.1
                                                                                                    _ = v49
                                                                                                    _ = v50
                                                                                                }
                                                                                                v18 -= v18
                                                                                                _ = v48
                                                                                            }
                                                                                        }
                                                                                        v34 -= (v34 | (v34 >> ((373) & 63)))
                                                                                        v8 = ((v8 + (v8 + (v8 + (((v8 + v8) + (v8 + "hello")) + "foo bar")))) + v8)
                                                                                        v18 += (v18 << ((42) & 63))
                                                                                        if len(v19) >= 2 {
                                                                                            v19 = v19[1:len(v19)]
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v29 = v29
                                                                                }
                                                                            }
                                                                            v18 = (v18 | (v18 + (v18 << ((4294967295) & 63))))
                                                                        }
                                                                        v34 -= 1
                                                                        if ((v18 % (((v18 * (^(v18)))) | 1)) == v18) {
                                                                            switch (fi16 & 9223372036854775807) {
                                                                            case 583:
                                                                                v19 = append(v19, (fi16 >> ((0) & 63)))
                                                                            case -2147483648:
                                                                                fmt.Println(v8)
                                                                            default:
                                                                                if len(v19) >= 2 {
                                                                                    v19 = v19[1:len(v19)]
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v35
                                                                    }
                                                                    var v51 uint16 = v29
                                                                    _ = v34
                                                                    _ = v51
                                                                }
                                                            case "":
                                                                fmt.Println(v8, fi16, v29)
                                                            case "\n\t":
                                                                fmt.Println(v29, fi7)
                                                            default:
                                                                v18 += uint8(v29)
                                                            }
                                                        }
                                                        _ = v29
                                                    }
                                                default:
                                                    switch v8 {
                                                    case "hello":
                                                        fmt.Println(fi16)
                                                    default:
                                                        if true {
                                                            var v52 int16 = 42
                                                            var v53 uint16 = 924
                                                            if len(v19) >= 2 {
                                                                v19 = v19[1:len(v19)]
                                                            }
                                                            _ = v52
                                                            _ = v53
                                                        } else {
                                                            v19 = append(v19, (fi7 >> ((65535) & 63)))
                                                            var v54 int16 = 1
                                                            _ = v54
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            switch fi16 {
                                            case 2147483647:
                                                v18 -= 11
                                            case 128:
                                                if v17 {
                                                    v55 := GenS9{F0: (^(v18)), F1: 18446744073709551615, F2: (!(((fi16 - int(v18)) <= (-(fi7))) || v17))}
                                                    v56 := GenS4{F0: 589, F1: -1, F2: 2147483647}
                                                    if true {
                                                        _ = len(v19)
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                }
                                            default:
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                            }
                                            var v57 int8 = 0
                                            _ = v57
                                        }
                                        var v58 int16 = -32768
                                        _ = v58
                                    } else {
                                        v59 := GenS3{F0: v17, F1: 1000, F2: 42}
                                        var v60 float32 = (3.14159 + 1e10)
                                        v8 = v8
                                        v61 := GenS10{F0: v17, F1: 1}
                                        switch fi7 {
                                        case -789:
                                            if (!v17) {
                                                fmt.Println(v8)
                                                v18 = 127
                                                if len(v19) > 0 {
                                                    v19[0] = fi16
                                                }
                                            }
                                        case 42:
                                            fmt.Println(v60, fi7, v17)
                                            v62 := GenS3{F0: (((v60 + float32(3.14159)) / float32(3.14159)) <= v60), F1: -1, F2: 1}
                                            _ = v62
                                        default:
                                            v63 := GenS5{F0: (v8 + v8), F1: 0, F2: 42}
                                            _ = v63
                                        }
                                        v60 += float32((float64((float32(-2.25) * (((3.14159 / 3.14159) * v60) - float32(1e10)))) / 1.0))
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                }
                                v18 += (v18 << ((2) & 63))
                            } else {
                                var v64 uint8 = 1
                                if len(v19) > 0 {
                                    v19[0] = (fi16 / ((fi16) | 1))
                                }
                                _ = v64
                            }
                            _ = len(v19)
                            v17 = ((fi7 - (fi16 | (fi16 << ((241) & 63)))) <= 128)
                        }
                        var v65 float32 = (0.0 / 1.5)
                        _ = v19
                        _ = v65
                    } else {
                        var v66 uint16 = 0
                        var v67 uint = 0
                        fmt.Println(v67, v18, v67)
                        _ = v66
                        _ = v67
                    }
                    v8 = v8
                    switch v8 {
                    case "hello":
                        v17 = (0.1 <= (0.0 / 1e10))
                    }
                }
                switch "" {
                case "hello":
                    var v68 string = "\n\t"
                    _ = v68
                case "a":
                    v18 -= (v18 ^ (v18 << ((4294967295) & 63)))
                default:
                    if false {
                        var v69 uint = 886
                        _ = v69
                    } else {
                        v70 := GenS1{F0: 4294967295}
                        var v71 int = (fi16 / ((fi16) | 1))
                        fmt.Println(v8, v18, fi7)
                        _ = v70
                        _ = v71
                    }
                }
                var v72 int = (fi7 - (fi7 & 1))
                _ = fi16
                _ = v17
                _ = v18
                _ = v72
            }
            _ = fi7
            _ = v8
        }
        var v73 uint16 = 346
        v74 := GenS1{F0: 398}
        _ = v73
        _ = v74
    } else {
        var v75 uint64 = 0
        v76 := GenS3{F0: (32768 > -572), F1: 42, F2: -128}
        _ = v75
        _ = v76
    }
    var v77 uint32 = 78
    for fi78 := 0; fi78 < 2; fi78++ {
        var v79 string = ""
        if false {
            var v80 uint32 = v77
            v81 := []int{(fi78 | -504), -125, fi78}
            v77 = (v77 >> ((65535) & 63))
            _ = v80
            _ = v81
        } else {
            if true {
                for fi82 := 0; fi82 < 13; fi82++ {
                    if ((!(!false)) || false) {
                        v77 -= (v77 + v77)
                        if true {
                            if ((("foo bar" + v79) + "\n\t") > (v79 + v79)) {
                                v83 := GenS1{F0: 1}
                                v79 = (((v79 + v79) + v79) + ("\n\t" + "hello"))
                                v84 := GenS11{F0: fi82}
                                _ = v83
                                _ = v84
                            }
                            v85 := []int{fi78, fi82, 9223372036854775807, (^(fi82))}
                            var v86 uint16 = 534
                            _ = v85
                            _ = v86
                        } else {
                            var v87 uint16 = 1
                            v88 := []int{fi78, (fi78 % ((fi82) | 1))}
                            v79 = ((((v79 + "a") + (v79 + (("a" + v79) + v79))) + "\n\t") + v79)
                            var v89 int32 = 2147483647
                            _ = v87
                            _ = v88
                            _ = v89
                        }
                        v79 = v79
                        v90 := GenS12{F0: 657, F1: -557, F2: (v79 + (v79 + (v79 + "\n\t")))}
                        switch (v79 + "foo bar") {
                        case "foo bar":
                            if (false && false) {
                                var v91 uint16 = 58
                                if false {
                                    v91 = (v91 ^ v91)
                                    switch (fi78 | (fi82 / ((fi78) | 1))) {
                                    case 32767:
                                        v92 := []int{(fi78 ^ (-(fi78))), (fi82 & (fi82 << ((991) & 63))), (fi78 | fi82)}
                                        _ = v92
                                    case 42:
                                        switch (v79 + ("hello" + v79)) {
                                        case "\n\t":
                                            if (v77 >= (v77 << ((0) & 63))) {
                                                v91 = (v91 ^ v91)
                                                var v93 int = (fi82 + (^(fi78)))
                                                fmt.Println(fi78, v79, v77)
                                                var v94 uint64 = 18446744073709551615
                                                _ = v93
                                                _ = v94
                                            } else {
                                                fmt.Println(v91, v79)
                                                var v95 uint16 = (v91 + (^(v91)))
                                                _ = v95
                                            }
                                        case "foo bar":
                                            fmt.Println(v77)
                                        default:
                                            var v96 uint16 = 508
                                            _ = v96
                                        }
                                    case -9223372036854775808:
                                        if (v77 >= 1) {
                                            fmt.Println(fi78, v91, fi82)
                                            v91 += 0
                                            var v97 int = (^(fi78))
                                            switch ("hello" + v79) {
                                            case "a":
                                                fmt.Println(v77)
                                            case "foo bar":
                                                switch (fi78 % ((v97) | 1)) {
                                                case 42:
                                                    if false {
                                                        if true {
                                                            v97 = fi78
                                                        } else {
                                                            v97 -= fi82
                                                            v98 := GenS3{F0: false, F1: -1, F2: -70}
                                                            v91 += v91
                                                            _ = v98
                                                        }
                                                        v99 := []int{}
                                                        if len(v99) >= 2 {
                                                            v99 = v99[1:len(v99)]
                                                        }
                                                        _ = v99
                                                    }
                                                case -842:
                                                    var v100 float64 = 3.14159
                                                    _ = v100
                                                case -9223372036854775808:
                                                    v91 = 32767
                                                default:
                                                    v101 := []int{-1, fi78, fi78, 624}
                                                    _ = v101
                                                }
                                            default:
                                                v97 += -39
                                            }
                                            _ = v97
                                        }
                                    default:
                                        v77 -= (v77 / ((675) | 1))
                                    }
                                    if (!((!(false || true)) || (false && true))) {
                                        if (fi78 > int(v91)) {
                                            v102 := GenS12{F0: 511, F1: fi78, F2: v79}
                                            v103 := GenS11{F0: fi78}
                                            _ = v102
                                            _ = v103
                                        }
                                    }
                                    var v104 string = v79
                                    v105 := GenS9{F0: 1, F1: 1, F2: true}
                                    _ = v104
                                    _ = v105
                                }
                                if true {
                                    var v106 uint = 490
                                    switch ("\n\t" + v79) {
                                    case "a":
                                        var v107 uint32 = v77
                                        _ = v107
                                    case "foo bar":
                                        switch (fi78 + (-(fi78))) {
                                        case 0:
                                            var v108 int32 = -294
                                            _ = v108
                                        default:
                                            if ((fi78 << ((v106) & 63)) >= -2147483648) {
                                                var v109 float32 = 1e10
                                                var v110 uint64 = 1
                                                var v111 int64 = 507
                                                _ = v109
                                                _ = v110
                                                _ = v111
                                            } else {
                                                var v112 uint8 = 83
                                                switch "" {
                                                case "\n\t":
                                                    v77 += (v77 & (v77 << ((v106) & 63)))
                                                case "foo bar":
                                                    fmt.Println(v77, fi78)
                                                case "":
                                                    switch (fi82 * -9223372036854775808) {
                                                    case 42:
                                                        v77 = (v77 % (((v77 + (v77 | (v77 % ((4294967295) | 1))))) | 1))
                                                    case -1:
                                                        switch (v79 + "foo bar") {
                                                        case "foo bar":
                                                            switch (fi82 / ((-293) | 1)) {
                                                            case 127:
                                                                var v113 int16 = 912
                                                                _ = v113
                                                            case -499:
                                                                fmt.Println(v106, fi78)
                                                            }
                                                        case "\n\t":
                                                            var v114 uint8 = v112
                                                            _ = v114
                                                        }
                                                    case -917:
                                                        v115 := []int{fi78, (fi82 & fi82), fi82, fi78}
                                                        _ = v115
                                                    default:
                                                        v79 = v79
                                                    }
                                                }
                                                switch fi82 {
                                                case -538:
                                                    v91 = 65535
                                                case -825:
                                                    v106 += 981
                                                case 143:
                                                    v106 = 718
                                                    if true {
                                                        v116 := []int{(fi82 / ((fi82) | 1)), (fi78 % ((fi82) | 1))}
                                                        var v117 int16 = 348
                                                        _ = v116
                                                        _ = v117
                                                    } else {
                                                        if (false && true) {
                                                            switch 32767 {
                                                            case 904:
                                                                var v118 float32 = -2.25
                                                                _ = v118
                                                            case -736:
                                                                v119 := []int{(fi82 | int(v106)), (fi78 + fi82)}
                                                                _ = v119
                                                            case 32768:
                                                                if (792 > -1) {
                                                                    if (-1 >= 42) {
                                                                        var v120 uint8 = v112
                                                                        v121 := GenS13{F0: -9223372036854775808, F1: 100, F2: -128}
                                                                        var v122 int16 = 0
                                                                        _ = v120
                                                                        _ = v121
                                                                        _ = v122
                                                                    } else {
                                                                        v123 := GenS2{F0: (^(v91)), F1: v91}
                                                                        if (!(true || (int8(v106) > 1))) {
                                                                            v124 := []int{int(v112), (fi82 * 42), (fi82 & int(v91)), (fi82 / (((fi78 & (fi78 | (^(fi78))))) | 1))}
                                                                            _ = v124
                                                                        } else {
                                                                            v125 := []int{int(v106), fi78, fi78}
                                                                            var v126 int16 = -140
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                        _ = v123
                                                                    }
                                                                    switch fi78 {
                                                                    case -9223372036854775808:
                                                                        var v127 uint32 = 1
                                                                        _ = v127
                                                                    case 32768:
                                                                        switch v79 {
                                                                        case "a":
                                                                            v106 += (^(v106))
                                                                        case "":
                                                                            v128 := GenS2{F0: v91, F1: v91}
                                                                            _ = v128
                                                                        }
                                                                        v129 := GenS14{F0: -119, F1: 0}
                                                                        _ = v129
                                                                    case 1:
                                                                        v106 += (v106 % ((v106) | 1))
                                                                    }
                                                                    fmt.Println(v112)
                                                                    v77 = 926
                                                                    v77 += v77
                                                                } else {
                                                                    if true {
                                                                        if ((v79 + (("hello" + (v79 + "foo bar")) + v79)) <= "\n\t") {
                                                                            fmt.Println(fi78)
                                                                            if (656 < 2) {
                                                                                var v130 uint = 65535
                                                                                v131 := []int{}
                                                                                var v132 int32 = int32(v130)
                                                                                _ = v130
                                                                                _ = v131
                                                                                _ = v132
                                                                            }
                                                                            v133 := GenS7{F0: 115, F1: 2147483647}
                                                                            if true {
                                                                                v134 := GenS15{F0: 880, F1: 0}
                                                                                switch (fi78 << (((v106 % (((v106 % (((v106 | 4294967295)) | 1))) | 1))) & 63)) {
                                                                                case -1:
                                                                                    v77 -= v77
                                                                                default:
                                                                                    fmt.Println(fi78, v91)
                                                                                }
                                                                                v91 = (v91 - (v91 >> ((v106) & 63)))
                                                                                _ = v134
                                                                            }
                                                                            _ = v133
                                                                        } else {
                                                                            v135 := GenS8{F0: (-(fi82))}
                                                                            var v136 int16 = -831
                                                                            switch (fi78 - fi78) {
                                                                            case -1:
                                                                                switch v79 {
                                                                                case "a":
                                                                                    v77 = v77
                                                                                case "hello":
                                                                                    fmt.Println(fi78)
                                                                                case "":
                                                                                    v137 := []int{fi82, -354}
                                                                                    _ = v137
                                                                                }
                                                                            default:
                                                                                var v138 uint = v106
                                                                                _ = v138
                                                                            }
                                                                            _ = v135
                                                                            _ = v136
                                                                        }
                                                                        var v139 string = ""
                                                                        _ = v139
                                                                    }
                                                                }
                                                            }
                                                            v79 = "foo bar"
                                                        } else {
                                                            v140 := GenS9{F0: v112, F1: 2, F2: true}
                                                            var v141 int = (fi82 ^ (fi78 ^ int(v77)))
                                                            _ = v140
                                                            _ = v141
                                                        }
                                                    }
                                                default:
                                                    var v142 int16 = -1
                                                    _ = v142
                                                }
                                                _ = v112
                                            }
                                        }
                                    case "\n\t":
                                        switch 2147483647 {
                                        case -9223372036854775808:
                                            var v143 int32 = 2147483647
                                            _ = v143
                                        case 42:
                                            if false {
                                                v77 = (^(v77))
                                                v144 := []int{fi82, (fi78 % ((42) | 1)), fi78}
                                                v77 = (v77 - v77)
                                                v144 = append(v144, 9223372036854775807)
                                                _ = v144
                                            }
                                        case -868:
                                            var v145 uint64 = 846
                                            _ = v145
                                        default:
                                            v106 = v106
                                        }
                                    default:
                                        switch 813 {
                                        case 66:
                                            var v146 int32 = 0
                                            _ = v146
                                        }
                                    }
                                    _ = v106
                                } else {
                                    v147 := GenS16{F0: 0}
                                    v91 -= (^(v91))
                                    _ = v147
                                }
                                _ = v91
                            }
                        default:
                            v148 := []int{fi78, (^(fi78))}
                            _ = v148
                        }
                        _ = v90
                    }
                    _ = fi82
                }
                var v149 int = (-(fi78))
                switch ("hello" + (v79 + v79)) {
                case "foo bar":
                    if ((v77 | (v77 ^ v77)) <= v77) {
                        fmt.Println(v77, v77, v77)
                        v149 += fi78
                        fmt.Println(v79, v77)
                    }
                }
                _ = v149
            } else {
                switch -681 {
                case 1:
                    var v150 float64 = float64((float32((0.1 - 1e10)) - (0.0 / float32((0.1 * 1.5)))))
                    _ = v150
                case -9223372036854775808:
                    fmt.Println(v79, fi78)
                case 32768:
                    v151 := []int{971}
                    _ = v151
                default:
                    v152 := GenS16{F0: 2}
                    _ = v152
                }
                var v153 float32 = (0.1 / 3.14159)
                _ = v153
            }
            switch "a" {
            case "hello":
                v154 := []int{}
                _ = v154
            default:
                for fi155 := 0; fi155 < 18; fi155++ {
                    v156 := GenS15{F0: fi78, F1: 42}
                    v157 := []int{}
                    _ = fi155
                    _ = v156
                    _ = v157
                }
            }
        }
        _ = fi78
        _ = v79
    }
    v77 = v77
    _ = v77
    return 0
}
func genFn1(p0 int8, p1 uint64, p2 int8) string {
    p1 = 2
    v158 := GenS17{F0: p2, F1: 255}
    v159 := []int{1}
    for fi160 := 0; fi160 < 1; fi160++ {
        var v161 int8 = (p0 - (p0 >> ((189) & 63)))
        var v162 int64 = 0
        var v163 int16 = 299
        _ = fi160
        _ = v161
        _ = v162
        _ = v163
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v158
    _ = v159
    return ""
}

func main() {
    v1 := []string{""}
    v2 := []string{("a" + "\n\t"), ("foo bar" + ("hello" + "\n\t")), "a", (("\n\t" + ("" + (("hello" + "\n\t") + "foo bar"))) + ("foo bar" + "\n\t"))}
    var v3 string = (((("hello" + "a") + "\n\t") + "hello") + ((("hello" + "") + "hello") + ("a" + "")))
    var v4 uint64 = 728
    var v164 uint = 18446744073709551615
    _ = len(v1)
    switch "hello" {
    case "":
        v164 += 0
    default:
        v165 := GenS18{F0: -70, F1: -382, F2: -9223372036854775808}
        _ = v165
    }
    switch 127 {
    case 978:
        var v166 int8 = 127
        _ = v166
    case -310:
        switch -703 {
        case 2147483647:
            if true {
                var v167 uint16 = 42
                _ = v167
            } else {
                _ = len(v2)
                if false {
                    for fi168 := 0; fi168 < 15; fi168++ {
                        fmt.Println(v164, fi168)
                        var v169 uint8 = 127
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi168
                        _ = v169
                    }
                } else {
                    var v170 uint64 = v4
                    var v171 float32 = 1.5
                    switch "" {
                    case "a":
                        for fi172 := 0; fi172 < 19; fi172++ {
                            v171 += v171
                            _ = fi172
                        }
                    case "foo bar":
                        v1 = append(v1, v3)
                    case "hello":
                        var v173 int = -5
                        _ = v173
                    }
                    _ = v170
                    _ = v171
                }
                if ((-684 >= 127) || ((33 <= 97) && false)) {
                    v174 := GenS15{F0: 4294967295, F1: -997}
                    _ = len(v1)
                    switch 32767 {
                    case 32767:
                        v175 := GenS14{F0: 100, F1: 65535}
                        _ = v175
                    default:
                        var v176 int16 = -1
                        _ = v176
                    }
                    _ = v174
                }
            }
        case 42:
            v164 = (v164 / (((v164 & (^(v164)))) | 1))
        default:
            v164 = 554
        }
    case 127:
        var v177 float32 = -2.25
        _ = v177
    default:
        switch -2147483648 {
        case 4294967295:
            fmt.Println(v4)
        case -2147483648:
            if ((!false) || true) {
                if (1 >= 280) {
                    if true {
                        var v178 int = -9223372036854775808
                        _ = v178
                    } else {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v164 += (v164 / (((v164 * (^(v164)))) | 1))
                        }
                        if false {
                            for fi179 := 0; fi179 < 18; fi179++ {
                                for fi180 := 0; fi180 < 4; fi180++ {
                                    v4 = v4
                                    v164 += 255
                                    _ = fi180
                                }
                                var v181 float64 = 0.0
                                v181 -= (v181 / (-2.25 * v181))
                                _ = fi179
                                _ = v181
                            }
                            v4 += v4
                            fmt.Println(v164, v4, v3)
                        }
                    }
                    fmt.Println(v3)
                    switch "" {
                    case "a":
                        for fi182 := 0; fi182 < 5; fi182++ {
                            v183 := GenS15{F0: (fi182 - fi182), F1: 1}
                            _ = fi182
                            _ = v183
                        }
                    case "\n\t":
                        var v184 string = "foo bar"
                        _ = v184
                    default:
                        if true {
                            v4 += (v4 * v4)
                            var v185 bool = false
                            _ = v185
                        } else {
                            v3 = (v3 + "foo bar")
                        }
                    }
                } else {
                    for fi186 := 0; fi186 < 20; fi186++ {
                        v3 = (((("a" + v3) + v3) + ("" + v3)) + "\n\t")
                        var v187 uint8 = 106
                        v164 = (v164 / ((2) | 1))
                        _ = fi186
                        _ = v187
                    }
                    var v188 int32 = 0
                    _ = v188
                }
                switch -697 {
                case 42:
                    fmt.Println(v4, v164)
                case -1:
                    switch 1 {
                    case 1:
                        if false {
                            var v189 uint32 = 65535
                            var v190 float32 = 1e10
                            fmt.Println(v189)
                            _ = v189
                            _ = v190
                        } else {
                            var v191 float64 = 0.1
                            switch (((((((("a" + "hello") + "a") + (v3 + ("a" + v3))) + "") + (v3 + "\n\t")) + ("" + v3)) + ("\n\t" + "\n\t")) + v3) {
                            case "\n\t":
                                fmt.Println(v4, v3)
                            case "foo bar":
                                _ = len(v2)
                            }
                            fmt.Println(v3, v191, v4)
                            _ = v191
                        }
                    case 9223372036854775807:
                        var v192 uint16 = 42
                        _ = v192
                    case -323:
                        var v193 int = -745
                        _ = v193
                    }
                case -2147483648:
                    var v194 int8 = 42
                    _ = v194
                }
                for fi195 := 0; fi195 < 20; fi195++ {
                    var v196 int8 = -65
                    var v197 bool = true
                    _ = fi195
                    _ = v196
                    _ = v197
                }
            } else {
                switch 4294967295 {
                case 639:
                    if true {
                        var v198 uint16 = 65535
                        var v199 float64 = 1.5
                        _ = v198
                        _ = v199
                    } else {
                        for fi200 := 0; fi200 < 14; fi200++ {
                            var v201 uint64 = (v4 >> ((65535) & 63))
                            _ = fi200
                            _ = v201
                        }
                        for fi202 := 0; fi202 < 13; fi202++ {
                            if (true && (!true)) {
                                v203 := GenS14{F0: 0, F1: 2147483648}
                                _ = v203
                            }
                            if false {
                                switch (^(fi202)) {
                                case 663:
                                    v204 := GenS13{F0: 42, F1: 0, F2: -8}
                                    _ = v204
                                case -211:
                                    v4 -= v4
                                case 2147483647:
                                    if (!false) {
                                        if false {
                                            _ = len(v2)
                                            var v205 uint = v164
                                            _ = v205
                                        }
                                    } else {
                                        var v206 uint32 = 2147483648
                                        if true {
                                            v164 += 255
                                        } else {
                                            for fi207 := 0; fi207 < 7; fi207++ {
                                                v164 = v164
                                                _ = fi207
                                            }
                                            if true {
                                                var v208 string = v3
                                                v164 = v164
                                                _ = v208
                                            } else {
                                                if false {
                                                    v209 := GenS11{F0: (fi202 >> ((18446744073709551615) & 63))}
                                                    _ = v209
                                                } else {
                                                    if false {
                                                        for fi210 := 0; fi210 < 10; fi210++ {
                                                            fmt.Println(fi210, v164)
                                                            if (fi210 != fi210) {
                                                                v3 = "foo bar"
                                                            } else {
                                                                _ = len(v1)
                                                                v164 = 18446744073709551615
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            _ = fi210
                                                        }
                                                        v211 := GenS19{F0: (true || false), F1: 0}
                                                        fmt.Println(fi202)
                                                        _ = v211
                                                    }
                                                    if (!((fi202 * fi202) == 1)) {
                                                        switch (fi202 >> ((18446744073709551615) & 63)) {
                                                        case 127:
                                                            v206 = 481
                                                        }
                                                    }
                                                    fmt.Println(v3, v206)
                                                }
                                                fmt.Println(v3, v164, fi202)
                                                var v212 int32 = 0
                                                _ = v212
                                            }
                                        }
                                        var v213 uint32 = 332
                                        _ = v206
                                        _ = v213
                                    }
                                }
                                if (fi202 < int(v4)) {
                                    if len(v2) > 0 {
                                        v2[0] = (v3 + v3)
                                    }
                                }
                                fmt.Println(fi202, v3, v4)
                            } else {
                                fmt.Println(fi202)
                            }
                            for fi214 := 0; fi214 < 13; fi214++ {
                                var v215 uint32 = 2147483647
                                v216 := GenS20{F0: 4294967295, F1: (fi202 | (fi202 | fi202))}
                                _ = fi214
                                _ = v215
                                _ = v216
                            }
                            _ = fi202
                        }
                        v3 = "\n\t"
                    }
                }
            }
        default:
            var v217 float32 = (3.14159 - (1.5 - 1e10))
            _ = v217
        }
    }
    v218 := GenS16{F0: v4}
    fmt.Println(v4, v4, v4)
    var v219 uint64 = (v4 + v4)
    var v220 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v164:", v164)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v219:", v219)
    fmt.Println("v220:", v220)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v164
    _ = v2
    _ = v218
    _ = v219
    _ = v220
    _ = v3
    _ = v4
}
