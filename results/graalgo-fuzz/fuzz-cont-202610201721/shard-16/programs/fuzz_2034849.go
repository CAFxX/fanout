package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int8; F1 uint32; F2 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16 }

func main() {
    v1 := []uint32{2, 42, 42, 2147483648}
    v2 := []int{4294967295, 2147483647, 0, -1}
    var v3 int64 = 0
    var v4 uint32 = 65535
    var v5 int16 = 1000
    for fi6 := 0; fi6 < 15; fi6++ {
        v5 += (v5 % ((v5) | 1))
        _ = fi6
    }
    if (float32(-2.25) > float32(1e10)) {
        v7 := GenS1{F0: (84 > 0), F1: ("\n\t" + "\n\t")}
        _ = v7
    }
    v5 = v5
    v3 -= (-(v3))
    v4 -= v4
    if false {
        var v8 uint8 = 127
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 int = -2147483648
                var v12 uint8 = 65
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    }
    var v13 int = 2147483647
    var v14 float32 = ((1e10 + 0.1) * (1.5 + ((0.1 - float32(1e10)) + 1e10)))
    fmt.Println(v3)
    switch (v13 << ((0) & 63)) {
    case 718:
        if ((v3 >= v3) && (false && true)) {
            v14 = v14
        } else {
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int = 1
                _ = fi15
                _ = v16
            }
            v3 += (v3 << ((2) & 63))
        }
    case 0:
        for fi17 := 0; fi17 < 2; fi17++ {
            v13 = (v13 + fi17)
            v5 = (v5 - v5)
            if false {
                if true {
                    var v18 float64 = (0.0 - ((0.1 + 3.14159) / 3.14159))
                    fmt.Println(v14, v3, v18)
                    _ = v18
                } else {
                    var v19 string = "hello"
                    _ = v19
                }
                switch 0 {
                case -9223372036854775808:
                    v4 -= v4
                case -2147483648:
                    var v20 int8 = 57
                    _ = v20
                case 1:
                    fmt.Println(v13, v14, v4)
                default:
                    switch (v13 | (^(v13))) {
                    case -375:
                        for fi21 := 0; fi21 < 3; fi21++ {
                            var v22 float64 = 1e10
                            fmt.Println(v4, v4, v4)
                            v4 -= 212
                            _ = fi21
                            _ = v22
                        }
                    case 0:
                        if (583 == 2) {
                            var v23 bool = false
                            _ = v23
                        } else {
                            v1 = append(v1, v4)
                            if (uint64(v4) == 1) {
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    var v25 float64 = 1.5
                                    if true {
                                        switch ("hello" + "") {
                                        case "a":
                                            v13 = (fi17 - -9223372036854775808)
                                        case "foo bar":
                                            fmt.Println(fi24)
                                        case "":
                                            if len(v1) > 0 {
                                                v1[0] = v4
                                            }
                                        }
                                        switch "" {
                                        case "\n\t":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "a":
                                            fmt.Println(v25, v25)
                                        default:
                                            v13 += v13
                                            v4 += 65535
                                            var v26 bool = true
                                            _ = v26
                                        }
                                    } else {
                                        var v27 float64 = 0.0
                                        fmt.Println(v13)
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v28 uint64 = 4294967295
                                            _ = v28
                                        case "":
                                            switch fi17 {
                                            case -9223372036854775808:
                                                if false {
                                                    if true {
                                                        if true {
                                                            fmt.Println(fi24, v3, v3)
                                                        } else {
                                                            if false {
                                                                fmt.Println(v3, v3)
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(v13)
                                                            case "foo bar":
                                                                var v29 uint16 = 388
                                                                _ = v29
                                                            default:
                                                                var v30 int16 = v5
                                                                _ = v30
                                                            }
                                                        }
                                                    }
                                                    var v31 uint64 = 4294967295
                                                    v32 := GenS1{F0: ("\n\t" > "\n\t"), F1: ((((("\n\t" + "hello") + "a") + ("foo bar" + "a")) + "foo bar") + "\n\t")}
                                                    _ = v31
                                                    _ = v32
                                                }
                                            case -2147483648:
                                                fmt.Println(v14, v27, v13)
                                            case 330:
                                                var v33 float64 = 1e10
                                                _ = v33
                                            }
                                        case "a":
                                            var v34 int64 = v3
                                            var v35 int = fi24
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v5 = v5
                                            _ = v34
                                            _ = v35
                                        default:
                                            v1 = append(v1, 0)
                                        }
                                        v36 := GenS1{F0: (127 == 100), F1: "a"}
                                        _ = v27
                                        _ = v36
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                                var v37 int16 = v5
                                _ = v37
                            } else {
                                v5 = v5
                            }
                        }
                    }
                }
            } else {
                if (1 != 1) {
                    if false {
                        if (!(!false)) {
                            var v38 int64 = 379
                            if true {
                                for fi39 := 0; fi39 < 10; fi39++ {
                                    var v40 int8 = 100
                                    if (42 > 42) {
                                        v38 -= 1
                                        v14 = (float32(1e10) * 1.5)
                                        v41 := GenS1{F0: false, F1: "a"}
                                        _ = v41
                                    } else {
                                        var v42 int32 = -1
                                        v3 = v38
                                        _ = v42
                                    }
                                    if false {
                                        var v43 uint8 = 2
                                        fmt.Println(v13, v3, v5)
                                        _ = v43
                                    }
                                    _ = fi39
                                    _ = v40
                                }
                                if (!false) {
                                    v5 = (v5 / (((v5 + v5)) | 1))
                                } else {
                                    v44 := GenS1{F0: (("foo bar" < ("foo bar" + ("\n\t" + ""))) || ((uint64(v4) != 920) || (v4 != v4))), F1: "\n\t"}
                                    fmt.Println(v5)
                                    fmt.Println(v5)
                                    _ = v44
                                }
                                v13 = fi17
                            }
                            for fi45 := 0; fi45 < 13; fi45++ {
                                if (!true) {
                                    if (!true) {
                                        switch "a" {
                                        case "":
                                            var v46 uint16 = 32767
                                            _ = v46
                                        case "a":
                                            var v47 float32 = float32(-2.25)
                                            _ = v47
                                        case "\n\t":
                                            v5 += (v5 ^ v5)
                                        }
                                    } else {
                                        v13 = -569
                                        switch ("hello" + ((("\n\t" + "hello") + (("\n\t" + "foo bar") + "a")) + "")) {
                                        case "\n\t":
                                            _ = len(v2)
                                        case "":
                                            if true {
                                                switch fi45 {
                                                case -9223372036854775808:
                                                    var v48 float32 = (-2.25 * (v14 / (v14 * (((float32(((float64(1e10) * 3.14159) - 1.5)) + 0.0) * (float32(-2.25) - v14)) / (v14 * ((((((v14 * (1e10 - (v14 - -2.25))) - 1.5) / 1.0) / v14) - 1.5) / float32((3.14159 - float64(v14)))))))))
                                                    _ = v48
                                                case 911:
                                                    if false {
                                                        var v49 uint64 = 324
                                                        v3 -= v38
                                                        var v50 int16 = -446
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                case 32768:
                                                    var v51 uint = 2
                                                    v52 := GenS2{F0: 100, F1: 80}
                                                    _ = v51
                                                    _ = v52
                                                }
                                                if ((3.14159 - ((3.14159 * v14) - v14)) < -2.25) {
                                                    v53 := GenS3{F0: "hello", F1: true}
                                                    _ = v53
                                                }
                                                v54 := GenS2{F0: int8(v4), F1: uint8(v4)}
                                                _ = v54
                                            }
                                        default:
                                            v14 -= 1e10
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = v4
                                        }
                                    }
                                    fmt.Println(v38, v3, v4)
                                    switch "\n\t" {
                                    case "a":
                                        if true {
                                            switch "foo bar" {
                                            case "":
                                                v13 -= (fi45 >> ((2) & 63))
                                                v38 += (^(v3))
                                            case "\n\t":
                                                if ((v5 >> ((18446744073709551615) & 63)) == v5) {
                                                    switch "" {
                                                    case "hello":
                                                        var v55 uint16 = 871
                                                        _ = v55
                                                    case "a":
                                                        switch int(v5) {
                                                        case 32767:
                                                            switch 373 {
                                                            case 127:
                                                                fmt.Println(v4, v4, fi45)
                                                            default:
                                                                v56 := GenS3{F0: "a", F1: false}
                                                                _ = v56
                                                            }
                                                            v38 = (v3 << ((255) & 63))
                                                            switch (v13 & v13) {
                                                            case 32768:
                                                                var v57 float32 = v14
                                                                _ = v57
                                                            case -9223372036854775808:
                                                                fmt.Println(v14, v3)
                                                            case -208:
                                                                v58 := GenS1{F0: true, F1: (("a" + "") + "")}
                                                                _ = v58
                                                            }
                                                            if (!false) {
                                                                v59 := GenS2{F0: -37, F1: 255}
                                                                _ = v59
                                                            } else {
                                                                v14 += v14
                                                                fmt.Println(v13)
                                                                v60 := GenS2{F0: 64, F1: 255}
                                                                switch "" {
                                                                case "hello":
                                                                    if (0 < 4294967295) {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        v61 := GenS2{F0: -1, F1: 127}
                                                                        v3 = -9223372036854775808
                                                                        _ = v61
                                                                    } else {
                                                                        var v62 uint32 = 2147483648
                                                                        _ = v62
                                                                    }
                                                                default:
                                                                    var v63 uint16 = 530
                                                                    _ = v63
                                                                }
                                                                fmt.Println(v4)
                                                                fmt.Println(fi17)
                                                                v64 := GenS3{F0: (((("\n\t" + "\n\t") + "") + "a") + ""), F1: false}
                                                                _ = v60
                                                                _ = v64
                                                            }
                                                            v13 = -631
                                                            v13 -= (v13 >> ((639) & 63))
                                                        default:
                                                            v5 -= (^(v5))
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                if false {
                                                    fmt.Println(v14, v4)
                                                }
                                            default:
                                                v3 += 9223372036854775807
                                            }
                                            v13 = v13
                                            fmt.Println(v3, v13, v38)
                                        } else {
                                            var v65 int16 = v5
                                            v66 := GenS4{F0: 1, F1: v4, F2: ""}
                                            switch ("foo bar" + "hello") {
                                            case "hello":
                                                fmt.Println(fi45, fi45)
                                                v14 = v14
                                                if (true || true) {
                                                    var v67 int64 = -1
                                                    if ((883 >= v4) || false) {
                                                        fmt.Println(fi45)
                                                        v67 = v67
                                                    } else {
                                                        v2 = append(v2, int(v5))
                                                        var v68 int64 = (-(v3))
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case "a":
                                                            v69 := GenS5{F0: "foo bar"}
                                                            _ = v69
                                                        case "hello":
                                                            v68 = v3
                                                        default:
                                                            var v70 uint16 = 877
                                                            _ = v70
                                                        }
                                                        var v71 uint64 = 2
                                                        var v72 float32 = v14
                                                        _ = v68
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    _ = v67
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = (v4 | 2147483648)
                                                    }
                                                    v73 := GenS6{F0: 487}
                                                    _ = v73
                                                }
                                            case "a":
                                                if false {
                                                    var v74 int8 = -128
                                                    if false {
                                                        var v75 float32 = v14
                                                        var v76 int64 = v3
                                                        var v77 uint = 252
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        v78 := GenS6{F0: 80}
                                                        switch fi17 {
                                                        case 2147483647:
                                                            v74 = 100
                                                        }
                                                        v3 = v3
                                                        _ = v78
                                                    }
                                                    v79 := GenS4{F0: -128, F1: v4, F2: "hello"}
                                                    _ = v74
                                                    _ = v79
                                                }
                                                fmt.Println(v13, fi45)
                                            case "":
                                                v65 += v5
                                            }
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                }
                                v80 := GenS6{F0: 0}
                                var v81 uint64 = 1
                                _ = fi45
                                _ = v80
                                _ = v81
                            }
                            _ = v38
                        } else {
                            var v82 uint8 = 236
                            if false {
                                for fi83 := 0; fi83 < 11; fi83++ {
                                    v84 := GenS7{F0: 0}
                                    fmt.Println(v13, v14, v13)
                                    _ = fi83
                                    _ = v84
                                }
                                if false {
                                    if (3.14159 > 1.5) {
                                        switch (fi17 ^ (v13 & (v13 & fi17))) {
                                        case -619:
                                            fmt.Println(v4, v3)
                                        case -1:
                                            var v85 float64 = (0.1 / 0.1)
                                            _ = v85
                                        }
                                    }
                                } else {
                                    v86 := GenS8{F0: 586, F1: (v82 | v82)}
                                    if true {
                                        var v87 int32 = -1
                                        _ = v87
                                    } else {
                                        fmt.Println(v3, v3, v13)
                                        var v88 bool = false
                                        _ = v88
                                    }
                                    _ = v86
                                }
                            }
                            if false {
                                for fi89 := 0; fi89 < 13; fi89++ {
                                    var v90 uint = 18446744073709551615
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = len(v2)
                                    _ = fi89
                                    _ = v90
                                }
                                for fi91 := 0; fi91 < 16; fi91++ {
                                    if len(v2) > 0 {
                                        v2[0] = fi91
                                    }
                                    if false {
                                        var v92 uint32 = (v4 * v4)
                                        if false {
                                            v82 = (v82 + v82)
                                            fmt.Println(fi17, v14)
                                            if (!(!(210 >= v82))) {
                                                var v93 float64 = 0.1
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v93
                                            } else {
                                                var v94 int16 = 42
                                                v94 = (-(v5))
                                                _ = v94
                                            }
                                            _ = len(v1)
                                            v14 = v14
                                            fmt.Println(fi91, fi91)
                                        }
                                        _ = v92
                                    } else {
                                        var v95 float32 = v14
                                        v13 += (^(v13))
                                        _ = v95
                                    }
                                    _ = fi91
                                }
                            } else {
                                switch ("\n\t" + "") {
                                case "foo bar":
                                    fmt.Println(v14)
                                case "a":
                                    if ((v3 >> ((4294967295) & 63)) < 42) {
                                        if true {
                                            fmt.Println(v5, v13)
                                        } else {
                                            for fi96 := 0; fi96 < 14; fi96++ {
                                                var v97 string = "foo bar"
                                                _ = fi96
                                                _ = v97
                                            }
                                            for fi98 := 0; fi98 < 4; fi98++ {
                                                v14 = (v14 + ((v14 * v14) / v14))
                                                v1 = append(v1, (^(v4)))
                                                _ = fi98
                                            }
                                        }
                                        for fi99 := 0; fi99 < 5; fi99++ {
                                            if (v5 == 1) {
                                                var v100 int = (v13 & (v13 - -1))
                                                _ = v100
                                            } else {
                                                v101 := GenS9{F0: v4}
                                                var v102 string = ""
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v101
                                                _ = v102
                                            }
                                            v2 = append(v2, fi17)
                                            _ = fi99
                                        }
                                        v103 := GenS7{F0: 924}
                                        _ = v103
                                    }
                                case "hello":
                                    fmt.Println(v3)
                                }
                                for fi104 := 0; fi104 < 2; fi104++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v105 float32 = ((((v14 + v14) + v14) * v14) * v14)
                                    var v106 string = "hello"
                                    _ = fi104
                                    _ = v105
                                    _ = v106
                                }
                            }
                            _ = v82
                        }
                    }
                    v107 := GenS10{F0: 867}
                    _ = v107
                }
                var v108 uint8 = 0
                _ = v108
            }
            _ = fi17
        }
    default:
        v3 += -9223372036854775808
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    v109 := GenS3{F0: "hello", F1: ((v13 <= v13) || ((1e10 * 3.14159) <= 0.1))}
    for fi110 := 0; fi110 < 12; fi110++ {
        v111 := GenS6{F0: 456}
        if (v13 < (v13 % ((v13) | 1))) {
            fmt.Println(v5)
            var v112 int64 = v3
            _ = v112
        }
        var v113 uint16 = 1
        _ = fi110
        _ = v111
        _ = v113
    }
    v3 = v3
    var v114 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v109
    _ = v114
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
