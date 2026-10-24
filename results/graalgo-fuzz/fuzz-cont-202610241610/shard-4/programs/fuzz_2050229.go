package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 bool; F1 int32; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 uint8; F1 string; F2 string }
type GenS12 struct { F0 uint16; F1 int16; F2 int8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int; F1 int8; F2 int32 }
type GenS17 struct { F0 string; F1 uint64; F2 uint32 }
type GenS18 struct { F0 uint; F1 int }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 bool; F1 bool; F2 int32 }
type GenS23 struct { F0 int32 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]bool)
    var v3 int32 = 91
    var v4 uint = 42
    v5 := make(map[uint32]bool)
    var v6 uint16 = 524
    v4 = (v4 << ((v4) & 63))
    v6 = (^(v6))
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(v3)
        fmt.Println(v3, v3, v6)
        _ = fi7
    }
    var v8 bool = ("foo bar" == "foo bar")
    var v9 int8 = 0
    if true {
        if v8 {
            v10 := GenS1{F0: "", F1: 1, F2: 255}
            v3 += -2147483648
            _ = v10
        }
        v9 = v9
        v11 := GenS1{F0: "a", F1: uint32(v4), F2: (v4 & 18446744073709551615)}
        _ = v11
    } else {
        if v8 {
            switch 9223372036854775807 {
            case 1:
                if (((!true) || v8) || false) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch ("" + ("foo bar" + ("foo bar" + ("a" + ("foo bar" + "foo bar"))))) {
                        case "foo bar":
                            v13 := GenS1{F0: ("foo bar" + "a"), F1: 1, F2: (^(v4))}
                            _ = v13
                        case "a":
                            fmt.Println(v4, v1, v6)
                        default:
                            if (1e10 <= 0.1) {
                                var v14 uint32 = 870
                                _ = v14
                            }
                        }
                        if v8 {
                            v15 := v5[0]
                            v6 -= v6
                            _ = v15
                        }
                        _ = fi12
                    }
                    if v8 {
                        v1 = v1
                        if true {
                            if v8 {
                                switch 2147483647 {
                                case 371:
                                    var v16 uint = (v4 / (((v4 ^ 2)) | 1))
                                    _ = v16
                                case 42:
                                    if (1 == 1) {
                                        if false {
                                            v17 := GenS1{F0: "a", F1: 560, F2: v4}
                                            v3 = (v3 / ((v3) | 1))
                                            _ = v17
                                        }
                                        v18 := GenS1{F0: "hello", F1: 2147483648, F2: v4}
                                        switch (("hello" + "hello") + "hello") {
                                        case "a":
                                            fmt.Println(v1, v6, v8)
                                        case "hello":
                                            for fi19 := 0; fi19 < 14; fi19++ {
                                                if (v9 < v9) {
                                                    if v8 {
                                                        var v20 string = (("hello" + "hello") + "a")
                                                        _ = v20
                                                    } else {
                                                        for fi21 := 0; fi21 < 2; fi21++ {
                                                            v22 := GenS2{F0: v3}
                                                            v6 += uint16(v4)
                                                            v9 += v9
                                                            _ = fi21
                                                            _ = v22
                                                        }
                                                        for fi23 := 0; fi23 < 15; fi23++ {
                                                            v8 = (!(1 > 42))
                                                            v24 := []int{(fi19 % ((fi19) | 1)), fi23}
                                                            _ = len(v24)
                                                            _ = fi23
                                                            _ = v24
                                                        }
                                                    }
                                                    v8 = true
                                                    v25 := []int{}
                                                    _ = v25
                                                }
                                                for fi26 := 0; fi26 < 14; fi26++ {
                                                    if v8 {
                                                        fmt.Println(v9, v8)
                                                        switch ("\n\t" + ("foo bar" + "\n\t")) {
                                                        case "a":
                                                            if v8 {
                                                                var v27 uint = v4
                                                                switch (fi19 >> ((720) & 63)) {
                                                                case 416:
                                                                    if v8 {
                                                                        if ((v6 ^ v6) == uint16(v1)) {
                                                                            switch (("a" + "a") + "hello") {
                                                                            case "a":
                                                                                if (true && v8) {
                                                                                    var v28 int32 = 42
                                                                                    v1 = v1
                                                                                    v1 += v1
                                                                                    _ = v28
                                                                                } else {
                                                                                    v5[500] = v8
                                                                                    v29 := GenS1{F0: ("foo bar" + ""), F1: 453, F2: v4}
                                                                                    _ = v29
                                                                                }
                                                                            case "foo bar":
                                                                                if (v8 || v8) {
                                                                                    v30 := GenS2{F0: (v3 ^ v3)}
                                                                                    var v31 uint = 0
                                                                                    v32 := GenS1{F0: ("hello" + (("\n\t" + "a") + ("a" + "\n\t"))), F1: 172, F2: v27}
                                                                                    _ = v30
                                                                                    _ = v31
                                                                                    _ = v32
                                                                                }
                                                                            case "hello":
                                                                                var v33 float64 = 1.5
                                                                                _ = v33
                                                                            }
                                                                            fmt.Println(v6)
                                                                        } else {
                                                                            if (uint8(v4) > v1) {
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        v27 -= (v27 - v27)
                                                                                        v3 -= v3
                                                                                    case "\n\t":
                                                                                        fmt.Println(v8, v1, v8)
                                                                                    default:
                                                                                        v27 -= (v4 | v4)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    var v34 string = "hello"
                                                                                    _ = v34
                                                                                default:
                                                                                    v35 := GenS1{F0: ("a" + (("a" + "a") + "a")), F1: uint32(v27), F2: v27}
                                                                                    if true {
                                                                                        v9 += v9
                                                                                    } else {
                                                                                        var v36 int64 = 714
                                                                                        v37 := GenS2{F0: (v3 | v3)}
                                                                                        _ = v36
                                                                                        _ = v37
                                                                                    }
                                                                                    _ = v35
                                                                                }
                                                                            }
                                                                            fmt.Println(v1)
                                                                        }
                                                                    } else {
                                                                        if v8 {
                                                                            switch (fi26 % (((fi26 * (fi26 & -953))) | 1)) {
                                                                            case 9223372036854775807:
                                                                                var v38 int = (fi26 % ((-1) | 1))
                                                                                v39 := []int{v38}
                                                                                _ = v38
                                                                                _ = v39
                                                                            case 587:
                                                                                var v40 uint32 = uint32(v6)
                                                                                _ = v40
                                                                            case 1:
                                                                                fmt.Println(v8, v4, v3)
                                                                            }
                                                                        }
                                                                        v1 = (v1 + v1)
                                                                    }
                                                                case 831:
                                                                    fmt.Println(fi26)
                                                                }
                                                                if v8 {
                                                                    v8 = (!((v6 <= 2) || (!v8)))
                                                                    v41 := GenS3{F0: 0, F1: "hello"}
                                                                    _ = v41
                                                                }
                                                                _ = v27
                                                            } else {
                                                                fmt.Println(fi26, v4)
                                                                v42 := GenS3{F0: 324, F1: "foo bar"}
                                                                v1 -= (v1 ^ 128)
                                                                _ = v42
                                                            }
                                                            v43 := GenS4{F0: v8, F1: -2147483648, F2: 1}
                                                            v9 = v9
                                                            v4 = (v4 | v4)
                                                            _ = v43
                                                        default:
                                                            v6 = v6
                                                        }
                                                    }
                                                    v1 += (v1 << ((v4) & 63))
                                                    _ = fi26
                                                }
                                                _ = fi19
                                            }
                                        default:
                                            var v44 int32 = (v3 - v3)
                                            _ = v44
                                        }
                                        _ = v18
                                    } else {
                                        fmt.Println(v4, v8, v1)
                                        for fi45 := 0; fi45 < 19; fi45++ {
                                            for fi46 := 0; fi46 < 16; fi46++ {
                                                if (!false) {
                                                    var v47 int = fi46
                                                    if v8 {
                                                        v8 = ((77 == 2) && true)
                                                    }
                                                    v1 = v1
                                                    _ = v47
                                                } else {
                                                    v9 = v9
                                                }
                                                _ = fi46
                                            }
                                            v9 += (^(v9))
                                            for fi48 := 0; fi48 < 2; fi48++ {
                                                v49 := GenS5{F0: v3}
                                                v50 := v5[2147483648]
                                                if (v4 == 4294967295) {
                                                    var v51 uint8 = (v1 ^ 171)
                                                    if true {
                                                        var v52 int = -9223372036854775808
                                                        v6 = (^(v6))
                                                        switch fi45 {
                                                        case 43:
                                                            var v53 uint64 = 1
                                                            _ = v53
                                                        case 128:
                                                            if false {
                                                                switch (fi48 >> ((65535) & 63)) {
                                                                case -9223372036854775808:
                                                                    var v54 int64 = 1
                                                                    _ = v54
                                                                case 400:
                                                                    if ("foo bar" < "foo bar") {
                                                                        v50 = (v3 < (v3 >> (((v4 ^ (^(v4)))) & 63)))
                                                                        v55 := GenS5{F0: v3}
                                                                        var v56 uint64 = 388
                                                                        _ = v55
                                                                        _ = v56
                                                                    } else {
                                                                        if (true && true) {
                                                                            if ((^(v4)) <= 0) {
                                                                                switch "" {
                                                                                case "a":
                                                                                    if v50 {
                                                                                        v50 = v50
                                                                                    } else {
                                                                                        v57 := []int{(fi48 % ((int(v6)) | 1)), (fi48 & (-(fi48))), (fi48 / ((fi45) | 1))}
                                                                                        v1 = 1
                                                                                        _ = v57
                                                                                    }
                                                                                    fmt.Println(v8)
                                                                                    var v58 int8 = int8(v3)
                                                                                    _ = v58
                                                                                case "foo bar":
                                                                                    v51 = v1
                                                                                }
                                                                            }
                                                                            v3 = (-(v3))
                                                                            v52 -= (fi45 + (v52 | 42))
                                                                        } else {
                                                                            if true {
                                                                                if ((fi48 * v52) < (fi48 - (fi45 / ((855) | 1)))) {
                                                                                    v1 += 42
                                                                                    switch (-(fi48)) {
                                                                                    case 363:
                                                                                        if (!(2 < 18446744073709551615)) {
                                                                                            switch (^(fi48)) {
                                                                                            case 32768:
                                                                                                switch -9223372036854775808 {
                                                                                                case 534:
                                                                                                    fmt.Println(v6)
                                                                                                case 2147483647:
                                                                                                    v3 = v3
                                                                                                case 0:
                                                                                                    var v59 uint = v4
                                                                                                    _ = v59
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    v60 := GenS5{F0: v3}
                                                                                                    var v61 float32 = 3.14159
                                                                                                    v51 -= 2
                                                                                                    _ = v60
                                                                                                    _ = v61
                                                                                                } else {
                                                                                                    if false {
                                                                                                        var v62 uint16 = (v6 & v6)
                                                                                                        fmt.Println(v62, v3, v3)
                                                                                                        var v63 uint64 = 0
                                                                                                        _ = v62
                                                                                                        _ = v63
                                                                                                    } else {
                                                                                                        v64 := GenS6{F0: ((v4 >> ((v4) & 63)) != (v4 - (v4 ^ (^(v4)))))}
                                                                                                        _ = v64
                                                                                                    }
                                                                                                    v65 := GenS6{F0: v8}
                                                                                                    var v66 int16 = -923
                                                                                                    _ = v65
                                                                                                    _ = v66
                                                                                                }
                                                                                            }
                                                                                            switch fi48 {
                                                                                            case 331:
                                                                                                if v50 {
                                                                                                    if false {
                                                                                                        var v67 int = (-(fi48))
                                                                                                        _ = v67
                                                                                                    }
                                                                                                    var v68 int8 = int8(v52)
                                                                                                    v9 += v9
                                                                                                    _ = v68
                                                                                                }
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    if (1 >= 2) {
                                                                                                        if v50 {
                                                                                                            v69 := []int{fi45, 726, 2147483647}
                                                                                                            var v70 uint8 = 0
                                                                                                            v71 := GenS7{F0: v1, F1: (v1 / ((v1) | 1))}
                                                                                                            _ = v69
                                                                                                            _ = v70
                                                                                                            _ = v71
                                                                                                        }
                                                                                                        var v72 uint32 = 142
                                                                                                        v73 := []int{}
                                                                                                        var v74 bool = (167 != (v72 / ((v72) | 1)))
                                                                                                        _ = v72
                                                                                                        _ = v73
                                                                                                        _ = v74
                                                                                                    }
                                                                                                case "hello":
                                                                                                    if ((0.1 + 1.5) == 1e10) {
                                                                                                        v51 = v51
                                                                                                        v75 := GenS1{F0: "", F1: 2, F2: (v4 >> ((65535) & 63))}
                                                                                                        var v76 float64 = -2.25
                                                                                                        _ = v75
                                                                                                        _ = v76
                                                                                                    } else {
                                                                                                        var v77 int64 = -372
                                                                                                        v4 += v4
                                                                                                        _ = v77
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v8 = true
                                                                                                }
                                                                                            case 32767:
                                                                                                var v78 int8 = v9
                                                                                                _ = v78
                                                                                            }
                                                                                        } else {
                                                                                            var v79 uint64 = 1
                                                                                            _ = v79
                                                                                        }
                                                                                    default:
                                                                                        if (!(-128 < v9)) {
                                                                                            var v80 int8 = 42
                                                                                            var v81 int8 = v80
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        }
                                                                                    }
                                                                                    v1 = (v51 / ((v51) | 1))
                                                                                } else {
                                                                                    var v82 bool = v50
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        delete(v5, 0)
                                                                                    }
                                                                                    v83 := GenS4{F0: true, F1: (v3 * 681), F2: 226}
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                }
                                                                            } else {
                                                                                if ((("a" + "a") > "") && false) {
                                                                                    v84 := GenS1{F0: ("hello" + "hello"), F1: 700, F2: 910}
                                                                                    _ = v84
                                                                                } else {
                                                                                    v9 -= v9
                                                                                }
                                                                            }
                                                                            fmt.Println(v8, v4)
                                                                            switch (fi48 * (fi45 << ((v4) & 63))) {
                                                                            case -156:
                                                                                v85 := []int{fi48, 42}
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                        if ((v9 << ((v4) & 63)) != (v9 << ((255) & 63))) {
                                                                            var v86 int16 = -17
                                                                            _ = v86
                                                                        } else {
                                                                            var v87 uint = 0
                                                                            _ = v87
                                                                        }
                                                                    }
                                                                case 32767:
                                                                    fmt.Println(v3)
                                                                default:
                                                                    switch (((("\n\t" + "\n\t") + "\n\t") + "\n\t") + "hello") {
                                                                    case "a":
                                                                        var v88 int16 = 1
                                                                        _ = v88
                                                                    default:
                                                                        var v89 float64 = ((float64((float32((-2.25 / 1.0)) * (-2.25 + (0.0 - 1.5)))) / float64(-2.25)) / 1.5)
                                                                        _ = v89
                                                                    }
                                                                }
                                                                var v90 uint64 = 879
                                                                var v91 int8 = v9
                                                                _ = v90
                                                                _ = v91
                                                            } else {
                                                                v92 := GenS8{F0: (v6 ^ v6)}
                                                                v93 := []int{(fi48 / ((606) | 1))}
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                        case 127:
                                                            if true {
                                                                v94 := GenS7{F0: v51, F1: 128}
                                                                _ = v94
                                                            } else {
                                                                if false {
                                                                    var v95 uint64 = 648
                                                                    _ = v95
                                                                } else {
                                                                    _ = len(v5)
                                                                    var v96 float64 = 1e10
                                                                    _ = v96
                                                                }
                                                                fmt.Println(v50, v3, fi48)
                                                            }
                                                        }
                                                        _ = v52
                                                    } else {
                                                        v9 += 0
                                                        if v50 {
                                                            v2[1] = v50
                                                            v3 -= (v3 ^ (v3 ^ 42))
                                                        }
                                                    }
                                                    _ = v51
                                                } else {
                                                    v1 = 0
                                                    if ((int(v3) <= 0) && v50) {
                                                        v97 := GenS2{F0: v3}
                                                        _ = v97
                                                    }
                                                    var v98 int32 = v3
                                                    _ = v98
                                                }
                                                _ = fi48
                                                _ = v49
                                                _ = v50
                                            }
                                            _ = fi45
                                        }
                                        if (!true) {
                                            var v99 uint64 = 0
                                            var v100 int64 = -872
                                            switch 4294967295 {
                                            case -1:
                                                switch -9223372036854775808 {
                                                case -2147483648:
                                                    v101 := GenS7{F0: 200, F1: uint8(v4)}
                                                    _ = v101
                                                case 0:
                                                    v3 = v3
                                                }
                                            case -619:
                                                if true {
                                                    var v102 string = ""
                                                    _ = v102
                                                }
                                            default:
                                                v8 = (0.1 >= -2.25)
                                            }
                                            _ = v100
                                            _ = v99
                                        }
                                    }
                                case 128:
                                    var v103 string = ("a" + ((("" + ("foo bar" + "foo bar")) + "hello") + ("foo bar" + "a")))
                                    _ = v103
                                default:
                                    v6 -= 65535
                                }
                                v4 = (v4 * v4)
                            }
                        }
                    }
                }
            case -2147483648:
                var v104 int = 127
                _ = v104
            case 127:
                v3 = (v3 << ((4294967295) & 63))
            default:
                if v8 {
                    for fi105 := 0; fi105 < 7; fi105++ {
                        v106 := GenS9{F0: 42}
                        var v107 string = "hello"
                        v108 := v2[0]
                        _ = fi105
                        _ = v106
                        _ = v107
                        _ = v108
                    }
                    delete(v5, 539)
                    var v109 uint64 = 18446744073709551615
                    _ = v109
                } else {
                    if (((939 <= uint64(v4)) || (!(!false))) && v8) {
                        v4 -= v4
                        v110 := GenS10{F0: v6, F1: v1}
                        _ = v110
                    }
                    var v111 uint = v4
                    _ = v111
                }
            }
            v112 := GenS11{F0: (v1 - (v1 + v1)), F1: "hello", F2: "foo bar"}
            v113 := []int{}
            _ = v112
            _ = v113
        }
    }
    if v8 {
        switch "foo bar" {
        case "hello":
            v114 := GenS12{F0: (v6 * (^(v6))), F1: -413, F2: 127}
            _ = v114
        case "a":
            switch 351 {
            case 1:
                var v115 int64 = -1
                _ = v115
            case 4294967295:
                var v116 uint8 = v1
                _ = v116
            }
        case "\n\t":
            switch 128 {
            case -629:
                for fi117 := 0; fi117 < 18; fi117++ {
                    v118 := v2[4294967295]
                    if ((1e10 - 3.14159) < (((((float64((0.0 - (1.5 - (((0.1 + 3.14159) / 1.5) + 0.0)))) + (1.5 * 1e10)) - (float64(1.5) + 1.5)) / 1.5) / 1.5) * (float64(1e10) / 1.5))) {
                        v119 := GenS7{F0: v1, F1: uint8(v6)}
                        var v120 uint16 = 65535
                        _ = v119
                        _ = v120
                    }
                    if ("a" < "\n\t") {
                        if (42 > 0) {
                            var v121 string = "a"
                            v122 := GenS13{F0: 534}
                            _ = v121
                            _ = v122
                        }
                    } else {
                        delete(v2, 853)
                        v9 = v9
                        _ = len(v5)
                    }
                    _ = fi117
                    _ = v118
                }
            case 127:
                v8 = v8
            case 9223372036854775807:
                v123 := GenS14{F0: (v4 * 2), F1: 716, F2: (v6 + (v6 ^ v6))}
                _ = v123
            }
        default:
            if v8 {
                v6 -= (^(v6))
                for fi124 := 0; fi124 < 10; fi124++ {
                    if v8 {
                        switch -690 {
                        case -686:
                            v2[794] = true
                        default:
                            if (2 <= 1) {
                                var v125 uint = v4
                                var v126 uint16 = v6
                                v127 := GenS14{F0: (^(v4)), F1: (v126 + 32767), F2: v126}
                                _ = v125
                                _ = v126
                                _ = v127
                            }
                        }
                    } else {
                        fmt.Println(v6)
                        var v128 float64 = 0.0
                        _ = v128
                    }
                    v3 = 42
                    v129 := []int{985, (^(fi124)), (fi124 % ((fi124) | 1)), fi124}
                    _ = fi124
                    _ = v129
                }
                fmt.Println(v9)
            } else {
                switch 177 {
                case 32768:
                    if true {
                        var v130 bool = (1 <= -42)
                        _ = v130
                    } else {
                        var v131 bool = true
                        v4 += v4
                        _ = v131
                    }
                case 127:
                    v132 := []int{32767}
                    _ = v132
                case 0:
                    v1 += (v1 >> (((^(v4))) & 63))
                }
                v1 -= 96
                v133 := GenS15{F0: 9223372036854775807}
                _ = v133
            }
        }
    } else {
        v134 := []int{9223372036854775807, int(v4), -640, -2147483648}
        _ = v134
    }
    var v135 int64 = -1
    switch 42 {
    case -451:
        v136 := GenS16{F0: 127, F1: v9, F2: v3}
        _ = v136
    case -2147483648:
        v3 -= -1
    case 1:
        for fi137 := 0; fi137 < 7; fi137++ {
            var v138 uint64 = 298
            var v139 uint = (v4 | (v4 ^ (v4 + v4)))
            fmt.Println(v135, v135, v3)
            _ = fi137
            _ = v138
            _ = v139
        }
    }
    v4 -= v4
    v140 := []int{9223372036854775807, -569}
    if ((-2.25 + float32((0.0 / 1.5))) >= (0.0 + 0.0)) {
        v141 := GenS4{F0: false, F1: v3, F2: 906}
        v140 = append(v140, -1)
        fmt.Println(v9, v6)
        _ = v141
    } else {
        switch 1 {
        case 786:
            for fi142 := 0; fi142 < 1; fi142++ {
                if false {
                    v135 -= (v135 * v135)
                    v1 = 0
                    if v8 {
                        switch "hello" {
                        case "\n\t":
                            var v143 uint = 1
                            _ = v143
                        case "":
                            if v8 {
                                v144 := GenS17{F0: "hello", F1: 975, F2: 1}
                                fmt.Println(v6, v6)
                                _ = v144
                            } else {
                                delete(v2, 4294967295)
                                v145 := GenS18{F0: 75, F1: (-(fi142))}
                                _ = v145
                            }
                        }
                        var v146 int8 = v9
                        var v147 uint = (v4 / ((1) | 1))
                        _ = v146
                        _ = v147
                    }
                }
                _ = fi142
            }
        default:
            switch 32768 {
            case 0:
                fmt.Println(v3)
            case 127:
                if ((v1 >= (v1 & v1)) || v8) {
                    if (-2.25 <= 0.0) {
                        v148 := v5[817]
                        v5[2147483648] = ((((0.1 - 3.14159) / 1e10) * 1e10) >= 3.14159)
                        v149 := GenS5{F0: v3}
                        _ = v148
                        _ = v149
                    } else {
                        if len(v140) > 0 {
                            v140[0] = -260
                        }
                        v150 := GenS19{F0: 42}
                        v1 += (v1 - 11)
                        _ = v150
                    }
                    for fi151 := 0; fi151 < 17; fi151++ {
                        v3 = -1
                        var v152 int32 = 42
                        _ = fi151
                        _ = v152
                    }
                }
            case -135:
                for fi153 := 0; fi153 < 13; fi153++ {
                    v3 -= int32(v9)
                    var v154 string = "a"
                    var v155 int64 = v135
                    _ = fi153
                    _ = v154
                    _ = v155
                }
            }
        }
        switch "a" {
        case "hello":
            for fi156 := 0; fi156 < 18; fi156++ {
                for fi157 := 0; fi157 < 19; fi157++ {
                    v140 = append(v140, (fi156 - (fi156 | 0)))
                    if v8 {
                        v158 := v2[2147483648]
                        v140 = append(v140, fi157)
                        var v159 string = ("hello" + "a")
                        v1 += 128
                        _ = v158
                        _ = v159
                    }
                    var v160 bool = (v135 >= (v135 ^ (^(v135))))
                    v161 := GenS16{F0: fi157, F1: (^(v9)), F2: v3}
                    _ = fi157
                    _ = v160
                    _ = v161
                }
                if v8 {
                    var v162 int8 = (^(v9))
                    v162 = -128
                    switch "foo bar" {
                    case "hello":
                        v163 := v2[822]
                        _ = v163
                    case "":
                        v8 = false
                    default:
                        switch "a" {
                        case "hello":
                            switch ("a" + "") {
                            case "\n\t":
                                var v164 int16 = 1
                                _ = v164
                            case "hello":
                                if len(v140) > 0 {
                                    v140[0] = fi156
                                }
                            case "foo bar":
                                for fi165 := 0; fi165 < 16; fi165++ {
                                    var v166 float32 = (float32(float64(float32((1.5 * 1e10)))) - 3.14159)
                                    if false {
                                        _ = len(v140)
                                        v167 := v5[2147483647]
                                        _ = v167
                                    }
                                    _ = fi165
                                    _ = v166
                                }
                            }
                        case "\n\t":
                            v5[4294967295] = (!v8)
                        case "":
                            v140 = append(v140, (fi156 << ((v4) & 63)))
                        }
                    }
                    _ = v162
                } else {
                    v135 = -72
                }
                _ = fi156
            }
        case "\n\t":
            v4 = 0
        case "":
            if true {
                for fi168 := 0; fi168 < 12; fi168++ {
                    if false {
                        var v169 int = (-(fi168))
                        for fi170 := 0; fi170 < 17; fi170++ {
                            fmt.Println(v9, v8, fi168)
                            fmt.Println(v1, fi168)
                            fmt.Println(v6, v8)
                            _ = fi170
                        }
                        if ((65535 >= 403) && v8) {
                            fmt.Println(v3, fi168)
                        }
                        _ = v169
                    }
                    v171 := GenS20{F0: 2}
                    v172 := GenS21{F0: "hello"}
                    _ = fi168
                    _ = v171
                    _ = v172
                }
            } else {
                v173 := GenS22{F0: v8, F1: v8, F2: v3}
                for fi174 := 0; fi174 < 8; fi174++ {
                    var v175 int16 = 42
                    v140 = append(v140, 9223372036854775807)
                    _ = fi174
                    _ = v175
                }
                for fi176 := 0; fi176 < 7; fi176++ {
                    var v177 uint64 = 69
                    v4 += uint(v177)
                    _ = fi176
                    _ = v177
                }
                _ = v173
            }
        default:
            var v178 uint32 = 365
            _ = v178
        }
    }
    if (v8 && false) {
        fmt.Println(v9, v3, v8)
        if (v8 || ((v4 % (((v4 >> ((v4) & 63))) | 1)) < v4)) {
            var v179 float32 = 1e10
            v1 = v1
            switch "foo bar" {
            case "a":
                v2[4294967295] = v8
            default:
                v4 -= v4
            }
            _ = v179
        } else {
            var v180 float32 = ((1.5 * float32((1.5 + (3.14159 + 1.5)))) * 1e10)
            _ = v180
        }
        if v8 {
            v9 = (v9 - (v9 & (v9 - (v9 % (((v9 & v9)) | 1)))))
            v140 = append(v140, 0)
        } else {
            for fi181 := 0; fi181 < 17; fi181++ {
                v9 -= v9
                v182 := GenS23{F0: (-(v3))}
                var v183 uint = (v4 + (^(v4)))
                _ = fi181
                _ = v182
                _ = v183
            }
            if v8 {
                var v184 int8 = v9
                v1 = (^(v1))
                _ = v184
            } else {
                var v185 float32 = float32(-2.25)
                var v186 uint16 = (v6 << (((^(v4))) & 63))
                _ = v185
                _ = v186
            }
            fmt.Println(v3, v8)
        }
    } else {
        var v187 float32 = -2.25
        v2[uint32(v4)] = (!(v4 >= v4))
        if false {
            var v188 uint64 = uint64(v1)
            _ = v188
        } else {
            for fi189 := 0; fi189 < 18; fi189++ {
                var v190 uint = v4
                switch fi189 {
                case 0:
                    v3 = v3
                case -9223372036854775808:
                    fmt.Println(v187, v4, v6)
                case 4294967295:
                    v191 := GenS17{F0: ("" + ("foo bar" + "\n\t")), F1: 18446744073709551615, F2: 0}
                    _ = v191
                }
                _ = fi189
                _ = v190
            }
            delete(v5, 4294967295)
        }
        _ = v187
    }
    switch 32767 {
    case 42:
        switch "foo bar" {
        case "foo bar":
            switch (("foo bar" + "\n\t") + "a") {
            case "\n\t":
                if false {
                    switch 32768 {
                    case -2147483648:
                        _ = len(v140)
                    default:
                        for fi192 := 0; fi192 < 14; fi192++ {
                            fmt.Println(v8)
                            var v193 uint16 = (^(v6))
                            _ = fi192
                            _ = v193
                        }
                    }
                } else {
                    var v194 uint8 = (^(v1))
                    _ = len(v140)
                    v194 -= (v194 >> ((uint(v6)) & 63))
                    _ = v194
                }
            case "hello":
                var v195 int32 = (v3 ^ 1)
                _ = v195
            case "a":
                v4 -= 255
            default:
                var v196 int64 = int64(v6)
                _ = v196
            }
        case "hello":
            if ((v8 || (!(!v8))) || v8) {
                fmt.Println(v6)
                v197 := GenS19{F0: 0}
                for fi198 := 0; fi198 < 10; fi198++ {
                    for fi199 := 0; fi199 < 3; fi199++ {
                        var v200 int32 = (v3 / (((^(v3))) | 1))
                        _ = fi199
                        _ = v200
                    }
                    _ = fi198
                }
                _ = v197
            }
        }
    case 2147483647:
        var v201 uint = 602
        _ = v201
    }
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("slicelen v140:", len(v140), v140)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v135
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
