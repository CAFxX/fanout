package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 string; F1 uint32; F2 bool }
type GenS11 struct { F0 uint; F1 uint64 }
type GenS12 struct { F0 int; F1 bool; F2 uint8 }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 int8; F1 uint32 }
type GenS15 struct { F0 uint16; F1 string }

func main() {
    v1 := []uint32{47, 2, 1}
    v2 := []int{0, -2147483648, 2147483647, 4294967295}
    var v3 float32 = (0.0 - 1e10)
    v4 := make(map[string]bool)
    v5 := []string{}
    if false {
        if false {
            if false {
                v6 := GenS1{F0: 255}
                if (false && ((float32(0.0) / float32(1e10)) < v3)) {
                    _ = len(v4)
                } else {
                    var v7 float64 = 0.0
                    _ = v7
                }
                _ = v6
            }
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint16 = 1
                if false {
                    fmt.Println(v3)
                    v3 = (v3 + (1e10 / (v3 - float32(-2.25))))
                } else {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        var v11 int8 = 8
                        var v12 uint32 = 1
                        if ("" <= "") {
                            if (!((2 > 255) || false)) {
                                fmt.Println(fi10)
                                fmt.Println(v11, v3)
                            } else {
                                var v13 uint16 = (^(v9))
                                v9 = v13
                                v13 = v9
                                _ = v13
                            }
                            var v14 float64 = (float64((v3 + ((0.0 / v3) * (float32(1.5) + v3)))) - 1e10)
                            v14 = (-2.25 / 1e10)
                            _ = v14
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                var v15 uint8 = 24
                                v16 := GenS2{F0: 1}
                                _ = v15
                                _ = v16
                            case "a":
                                var v17 int64 = -1
                                _ = v17
                            case "hello":
                                v18 := GenS2{F0: 1}
                                _ = v18
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    v19 := GenS3{F0: (^(v11)), F1: 1}
                                    _ = v19
                                default:
                                    if false {
                                        v9 -= 0
                                        _ = len(v4)
                                        var v20 uint16 = v9
                                        var v21 uint64 = 907
                                        v11 = -60
                                        if (int32(v9) <= -1) {
                                            var v22 uint8 = 1
                                            var v23 uint8 = (v22 % ((uint8(v21)) | 1))
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if (!true) {
                                            v12 = (^(v12))
                                        } else {
                                            if (true && (2 == (v12 | 576))) {
                                                var v24 int32 = -229
                                                var v25 int64 = -114
                                                v24 = v24
                                                _ = v24
                                                _ = v25
                                            } else {
                                                switch ("foo bar" + "foo bar") {
                                                case "foo bar":
                                                    if true {
                                                        fmt.Println(v9, v11, v9)
                                                        var v26 int8 = v11
                                                        v12 -= 551
                                                        if false {
                                                            v27 := GenS3{F0: (v26 * (^(v26))), F1: 1}
                                                            switch -586 {
                                                            case 0:
                                                                v28 := GenS4{F0: true}
                                                                var v29 int32 = int32(v11)
                                                                v30 := GenS5{F0: 212}
                                                                _ = v28
                                                                _ = v29
                                                                _ = v30
                                                            case 544:
                                                                v4["a"] = (!true)
                                                            case 128:
                                                                var v31 int32 = -2147483648
                                                                _ = v31
                                                            }
                                                            _ = v27
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v26
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        var v32 int32 = -2147483648
                                                        v12 -= (^(v12))
                                                        var v33 uint32 = (v12 % ((v12) | 1))
                                                        switch "hello" {
                                                        case "hello":
                                                            v33 = v33
                                                            v34 := GenS6{F0: -1}
                                                            v35 := GenS1{F0: 204}
                                                            _ = v34
                                                            _ = v35
                                                        default:
                                                            var v36 int = (fi8 + (fi10 + (^(fi10))))
                                                            _ = v36
                                                        }
                                                        if false {
                                                            var v37 uint = 0
                                                            v3 -= float32(0.1)
                                                            _ = v37
                                                        }
                                                        v38 := GenS5{F0: 255}
                                                        var v39 bool = false
                                                        var v40 uint = 4294967295
                                                        var v41 bool = false
                                                        _ = v32
                                                        _ = v33
                                                        _ = v38
                                                        _ = v39
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                case "hello":
                                                    switch ("" + ("a" + "a")) {
                                                    case "\n\t":
                                                        switch int(v9) {
                                                        case 2147483647:
                                                            v11 -= v11
                                                        case 42:
                                                            v11 = (^(v11))
                                                        default:
                                                            if (true && (true && true)) {
                                                                if true {
                                                                    if false {
                                                                        _ = len(v4)
                                                                    } else {
                                                                        v12 += (v12 % ((uint32(v9)) | 1))
                                                                        var v42 int64 = 0
                                                                        v4["foo bar"] = (0 >= 2147483647)
                                                                        _ = v42
                                                                    }
                                                                    switch fi10 {
                                                                    case 691:
                                                                        if true {
                                                                            v43 := GenS3{F0: v11, F1: uint64(v9)}
                                                                            var v44 uint64 = 18446744073709551615
                                                                            v45 := GenS1{F0: 0}
                                                                            _ = len(v5)
                                                                            _ = v43
                                                                            _ = v44
                                                                            _ = v45
                                                                        }
                                                                    case 32767:
                                                                        var v46 int16 = 1
                                                                        _ = v46
                                                                    }
                                                                    var v47 uint64 = 18446744073709551615
                                                                    _ = v47
                                                                } else {
                                                                    fmt.Println(v9, fi8)
                                                                }
                                                                v12 = v12
                                                                if ((42 >= 4294967295) || (!true)) {
                                                                    var v48 int = (fi8 % ((417) | 1))
                                                                    switch fi8 {
                                                                    case -816:
                                                                        v12 += v12
                                                                    case -2147483648:
                                                                        if true {
                                                                            var v49 int16 = 1
                                                                            v9 -= v9
                                                                            if false {
                                                                                v50 := GenS5{F0: 128}
                                                                                v51 := v4[(("\n\t" + "foo bar") + "hello")]
                                                                                _ = v50
                                                                                _ = v51
                                                                            } else {
                                                                                if false {
                                                                                    delete(v4, "")
                                                                                    v11 = (^(v11))
                                                                                }
                                                                                var v52 int32 = -2147483648
                                                                                fmt.Println(v9, v9, v9)
                                                                                v12 = (v12 ^ (v12 / (((v12 | 1)) | 1)))
                                                                                if true {
                                                                                    v49 = v49
                                                                                } else {
                                                                                    var v53 uint32 = uint32(v9)
                                                                                    v12 = (v12 | 4294967295)
                                                                                    _ = v53
                                                                                }
                                                                                _ = v52
                                                                            }
                                                                            _ = v49
                                                                        }
                                                                    case -122:
                                                                        if false {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v54 float64 = 3.14159
                                                                            var v55 float64 = v54
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                    }
                                                                    var v56 int64 = -1
                                                                    switch -437 {
                                                                    case -9223372036854775808:
                                                                        if (!(((!false) || false) && false)) {
                                                                            var v57 int32 = 42
                                                                            _ = v57
                                                                        } else {
                                                                            v48 = int(v56)
                                                                            if (!(!true)) {
                                                                                delete(v4, "\n\t")
                                                                            } else {
                                                                                _ = len(v1)
                                                                                v3 += v3
                                                                                var v58 uint32 = (v12 * (^(v12)))
                                                                                _ = v58
                                                                            }
                                                                            fmt.Println(v3)
                                                                            var v59 float64 = 0.1
                                                                            fmt.Println(v3)
                                                                            _ = v59
                                                                        }
                                                                    case -234:
                                                                        var v60 int16 = 398
                                                                        _ = v60
                                                                    default:
                                                                        if true {
                                                                            if (((3.14159 - (0.0 * float64(((v3 + (v3 + v3)) * (v3 * v3))))) * 0.1) > 1e10) {
                                                                                if (0.0 == v3) {
                                                                                    if false {
                                                                                        var v61 string = ""
                                                                                        switch (v48 % ((fi8) | 1)) {
                                                                                        case -1:
                                                                                            fmt.Println(v56, v12, v56)
                                                                                        case 499:
                                                                                            v62 := v4[v61]
                                                                                            _ = v62
                                                                                        }
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            v5 = append(v5, v61)
                                                                                        case "a":
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = "\n\t"
                                                                                            }
                                                                                        default:
                                                                                            v4[(("\n\t" + "foo bar") + "hello")] = true
                                                                                        }
                                                                                        _ = v61
                                                                                    } else {
                                                                                        v63 := GenS5{F0: 255}
                                                                                        var v64 int8 = (^(v11))
                                                                                        v11 += (v11 % ((127) | 1))
                                                                                        v2 = append(v2, 426)
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                    }
                                                                                    v65 := GenS6{F0: 1000}
                                                                                    if true {
                                                                                        if true {
                                                                                            v11 -= (v11 << ((42) & 63))
                                                                                        } else {
                                                                                            fmt.Println(v9, v3, v56)
                                                                                            v66 := GenS7{F0: v56, F1: int16(v48)}
                                                                                            _ = v66
                                                                                        }
                                                                                        var v67 uint8 = 1
                                                                                        switch (("" + ("hello" + "")) + "foo bar") {
                                                                                        case "hello":
                                                                                            if (315 >= 356) {
                                                                                                _ = len(v4)
                                                                                                v68 := GenS1{F0: v67}
                                                                                                _ = v68
                                                                                            }
                                                                                        case "foo bar":
                                                                                            var v69 uint8 = v67
                                                                                            _ = v69
                                                                                        case "":
                                                                                            var v70 uint64 = 448
                                                                                            _ = v70
                                                                                        default:
                                                                                            var v71 int64 = 1
                                                                                            _ = v71
                                                                                        }
                                                                                        _ = v67
                                                                                    }
                                                                                    v56 = (v56 - 983)
                                                                                    fmt.Println(v48, v3)
                                                                                    _ = v65
                                                                                }
                                                                            } else {
                                                                                v72 := GenS5{F0: 128}
                                                                                v56 -= (v56 << ((42) & 63))
                                                                                v3 += v3
                                                                                _ = v72
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                switch ("\n\t" + "foo bar") {
                                                                                case "foo bar":
                                                                                    fmt.Println(fi10, v11, fi10)
                                                                                    var v73 int8 = (v11 << ((uint(v9)) & 63))
                                                                                    var v74 int8 = (-(v11))
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                default:
                                                                                    fmt.Println(v48, fi8)
                                                                                }
                                                                                fmt.Println(v9, v56)
                                                                                var v75 bool = false
                                                                                _ = v75
                                                                            } else {
                                                                                if true {
                                                                                    switch (-(fi8)) {
                                                                                    case 9223372036854775807:
                                                                                        v12 = (^(v12))
                                                                                    case 32767:
                                                                                        if true {
                                                                                            fmt.Println(v11)
                                                                                            fmt.Println(v48, v12, v11)
                                                                                            switch ("hello" + "\n\t") {
                                                                                            case "foo bar":
                                                                                                var v76 uint8 = 1
                                                                                                _ = v76
                                                                                            case "\n\t":
                                                                                                v77 := v4[((((("" + "a") + "") + "a") + "\n\t") + "a")]
                                                                                                v48 += 0
                                                                                                if (v77 || (v77 && (!false))) {
                                                                                                    fmt.Println(v3, v12)
                                                                                                }
                                                                                                fmt.Println(v9)
                                                                                                _ = v77
                                                                                            default:
                                                                                                switch v48 {
                                                                                                case 4294967295:
                                                                                                    v78 := GenS1{F0: 128}
                                                                                                    _ = v78
                                                                                                case 529:
                                                                                                    v79 := v4[("\n\t" + (((("" + "a") + "") + ("foo bar" + "hello")) + "foo bar"))]
                                                                                                    _ = v79
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v3 = float32(1e10)
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                v80 := v4[("foo bar" + ("foo bar" + (("a" + (("foo bar" + (("" + "hello") + (("hello" + "a") + ""))) + ("a" + "a"))) + "\n\t")))]
                                                                                                _ = v80
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v81 int32 = 122
                                                                                        _ = v81
                                                                                    }
                                                                                }
                                                                                var v82 int8 = (v11 & -85)
                                                                                _ = v82
                                                                            }
                                                                            fmt.Println(v12, fi8)
                                                                            v9 += v9
                                                                        }
                                                                    }
                                                                    _ = v48
                                                                    _ = v56
                                                                } else {
                                                                    if true {
                                                                        var v83 float64 = 1e10
                                                                        if false {
                                                                            fmt.Println(v11)
                                                                            switch (fi8 >> ((4294967295) & 63)) {
                                                                            case 1:
                                                                                if false {
                                                                                    var v84 uint = 18446744073709551615
                                                                                    _ = v84
                                                                                }
                                                                            case -314:
                                                                                v3 = (v3 + (float32(1.5) * float32(float64(v3))))
                                                                            case 9223372036854775807:
                                                                                var v85 float32 = (v3 - (v3 - (((((v3 - ((((1e10 * (v3 - v3)) - ((0.0 / v3) + (v3 - v3))) * 0.0) + 0.1)) - float32(v83)) - v3) - (1e10 + (0.0 + v3))) / (v3 * v3))))
                                                                                _ = v85
                                                                            default:
                                                                                fmt.Println(v3)
                                                                            }
                                                                        } else {
                                                                            var v86 uint16 = v9
                                                                            v87 := GenS1{F0: 2}
                                                                            v88 := GenS4{F0: ((fi10 / ((-285) | 1)) == fi10)}
                                                                            _ = v86
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        var v89 string = ("hello" + "\n\t")
                                                                        _ = v83
                                                                        _ = v89
                                                                    } else {
                                                                        var v90 int64 = 936
                                                                        _ = v90
                                                                    }
                                                                }
                                                            } else {
                                                                v12 = v12
                                                            }
                                                        }
                                                    default:
                                                        var v91 int = 0
                                                        _ = v91
                                                    }
                                                default:
                                                    if true {
                                                        switch 0 {
                                                        case -2147483648:
                                                            if (!(2 >= 509)) {
                                                                var v92 int16 = 42
                                                                _ = v92
                                                            } else {
                                                                var v93 string = "hello"
                                                                var v94 uint16 = 0
                                                                v12 = v12
                                                                var v95 float64 = ((3.14159 * -2.25) * (3.14159 - (1e10 - 0.0)))
                                                                _ = v93
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                        }
                                                        v12 = v12
                                                        v3 -= 1e10
                                                        _ = len(v2)
                                                    } else {
                                                        v11 = 100
                                                        if (true && true) {
                                                            v11 = (v11 / ((v11) | 1))
                                                            _ = len(v4)
                                                        } else {
                                                            v9 -= v9
                                                            v11 = 100
                                                        }
                                                    }
                                                }
                                                v9 += (^(v9))
                                                v12 = v12
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = ("foo bar" + "hello")
                                            }
                                            if ((-1 <= -2147483648) || false) {
                                                v11 -= (^(v11))
                                            } else {
                                                v96 := v4["hello"]
                                                _ = v96
                                            }
                                        }
                                        var v97 int8 = (v11 - v11)
                                        v4[""] = (42 > v9)
                                        _ = v97
                                    }
                                }
                            }
                            v3 -= (1.5 * (0.0 / v3))
                        }
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    if true {
                        if true {
                            if false {
                                v3 -= (float32((3.14159 + ((1.5 - (float64(1.5) * 3.14159)) + 3.14159))) + v3)
                                for fi98 := 0; fi98 < 20; fi98++ {
                                    switch (fi8 >> ((65535) & 63)) {
                                    case 1:
                                        v9 -= v9
                                    case 127:
                                        var v99 int = (fi8 + fi8)
                                        _ = v99
                                    default:
                                        var v100 string = (("foo bar" + "hello") + "\n\t")
                                        _ = v100
                                    }
                                    if ((0.0 * -2.25) == ((1.5 * 1e10) * 0.0)) {
                                        v9 = (v9 + (v9 | v9))
                                        _ = len(v4)
                                    }
                                    _ = fi98
                                }
                                var v101 float64 = (3.14159 * ((0.1 / 1.5) + 3.14159))
                                _ = v101
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            v3 -= float32(0.1)
                            fmt.Println(v9, v9, v3)
                        }
                    }
                }
                v9 = 65535
                _ = fi8
                _ = v9
            }
            var v102 bool = (true || true)
            v3 -= v3
            _ = v102
        }
    } else {
        v3 = ((v3 * 3.14159) + ((v3 + 1e10) + 3.14159))
        var v103 string = "a"
        var v104 int16 = -32768
        _ = v103
        _ = v104
    }
    fmt.Println(v3)
    var v105 int = 845
    var v106 int = 0
    if (2 < 18446744073709551615) {
        v107 := GenS3{F0: 0, F1: 0}
        v108 := GenS1{F0: 128}
        for fi109 := 0; fi109 < 11; fi109++ {
            v105 -= 128
            v3 += v3
            _ = fi109
        }
        _ = v107
        _ = v108
    } else {
        for fi110 := 0; fi110 < 12; fi110++ {
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
            if ((2 <= 255) && ((!false) && true)) {
                if true {
                    v5 = append(v5, "a")
                    if len(v2) > 0 {
                        v2[0] = v105
                    }
                } else {
                    var v111 int64 = 42
                    if false {
                        switch fi110 {
                        case 127:
                            v105 -= (fi110 / (((v106 << ((65535) & 63))) | 1))
                        case -290:
                            fmt.Println(v3, v111)
                        case 100:
                            var v112 uint = 0
                            _ = v112
                        }
                        v105 = -464
                        fmt.Println(v105, v106, v111)
                    } else {
                        switch (v106 >> ((2) & 63)) {
                        case 9223372036854775807:
                            fmt.Println(v3, v105, v3)
                        case 804:
                            var v113 string = "hello"
                            _ = v113
                        case 42:
                            v105 += v106
                        default:
                            var v114 uint16 = 65535
                            _ = v114
                        }
                        if (127 > 1) {
                            switch ("\n\t" + ("hello" + (((("foo bar" + "\n\t") + ("a" + "\n\t")) + "a") + "foo bar"))) {
                            case "foo bar":
                                var v115 int = v105
                                _ = v115
                            default:
                                for fi116 := 0; fi116 < 15; fi116++ {
                                    var v117 float32 = v3
                                    _ = fi116
                                    _ = v117
                                }
                            }
                            var v118 float32 = 3.14159
                            _ = v118
                        } else {
                            v105 = (fi110 - 128)
                            for fi119 := 0; fi119 < 18; fi119++ {
                                delete(v4, "hello")
                                _ = fi119
                            }
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v3)
                            }
                        }
                    }
                    fmt.Println(v106, v105, fi110)
                    _ = v111
                }
                v1 = append(v1, 2147483648)
                if true {
                    v106 = (v105 + fi110)
                    if ((false || false) || false) {
                        if (!true) {
                            v120 := GenS2{F0: 1000}
                            for fi121 := 0; fi121 < 10; fi121++ {
                                switch "a" {
                                case "":
                                    var v122 int8 = 1
                                    _ = v122
                                case "a":
                                    v123 := GenS3{F0: 12, F1: 4294967295}
                                    _ = v123
                                case "hello":
                                    var v124 float32 = 3.14159
                                    _ = v124
                                default:
                                    switch "" {
                                    case "a":
                                        _ = len(v1)
                                    case "\n\t":
                                        v125 := GenS2{F0: 1000}
                                        _ = v125
                                    default:
                                        v105 -= 127
                                    }
                                }
                                v126 := GenS8{F0: v106, F1: 32767}
                                v3 -= 0.0
                                v127 := GenS1{F0: 46}
                                _ = fi121
                                _ = v126
                                _ = v127
                            }
                            var v128 string = "\n\t"
                            _ = v120
                            _ = v128
                        }
                        var v129 int16 = 31
                        for fi130 := 0; fi130 < 8; fi130++ {
                            switch v105 {
                            case 4294967295:
                                switch "foo bar" {
                                case "":
                                    v105 += fi130
                                case "a":
                                    var v131 uint = 65535
                                    _ = v131
                                case "foo bar":
                                    var v132 int32 = 961
                                    _ = v132
                                }
                                if false {
                                    fmt.Println(fi130, v106, v105)
                                    fmt.Println(fi110, v129)
                                }
                                v129 = (^(v129))
                            case -654:
                                var v133 uint32 = 1
                                _ = v133
                            }
                            v134 := GenS9{F0: 42}
                            v105 = fi110
                            var v135 int8 = -128
                            _ = fi130
                            _ = v134
                            _ = v135
                        }
                        _ = v129
                    } else {
                        fmt.Println(fi110)
                        if (!(669 >= -1)) {
                            var v136 float64 = 1.5
                            switch v106 {
                            case 42:
                                for fi137 := 0; fi137 < 13; fi137++ {
                                    var v138 float64 = (float64(float32(v136)) + v136)
                                    _ = fi137
                                    _ = v138
                                }
                            case -9223372036854775808:
                                v139 := GenS9{F0: 16}
                                _ = v139
                            default:
                                if (667 <= 251) {
                                    v106 += v105
                                }
                            }
                            v105 = v106
                            _ = v136
                        } else {
                            v140 := GenS10{F0: "a", F1: 1, F2: true}
                            _ = v140
                        }
                    }
                } else {
                    var v141 uint64 = 4294967295
                    switch 534 {
                    case -450:
                        if (255 >= 127) {
                            v1 = append(v1, 673)
                        } else {
                            var v142 uint32 = 2147483648
                            if true {
                                v3 = ((3.14159 * ((v3 / (float32((float64((0.0 / float32(3.14159))) + float64(v3))) + (v3 * v3))) - 1.5)) / v3)
                            }
                            _ = v142
                        }
                    default:
                        v106 = v106
                    }
                    _ = v141
                }
            } else {
                var v143 bool = true
                v143 = (-2147483648 < v105)
                var v144 uint32 = 0
                _ = v143
                _ = v144
            }
            _ = fi110
        }
        if false {
            var v145 string = "hello"
            for fi146 := 0; fi146 < 3; fi146++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = len(v5)
                v147 := GenS11{F0: 4294967295, F1: 862}
                _ = fi146
                _ = v147
            }
            var v148 int = v105
            _ = v145
            _ = v148
        } else {
            fmt.Println(v105, v106, v105)
        }
    }
    if len(v2) > 0 {
        v2[0] = v105
    }
    switch -544 {
    case 42:
        v3 += 0.1
    }
    if (((-2147483648 != -529) && true) && (3.14159 == (v3 - 0.1))) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v149 uint16 = 1
        _ = v149
    } else {
        switch "hello" {
        case "":
            for fi150 := 0; fi150 < 15; fi150++ {
                v105 += (fi150 >> ((42) & 63))
                _ = fi150
            }
        case "hello":
            if (0 != -2147483648) {
                var v151 string = "\n\t"
                var v152 uint64 = 42
                _ = v151
                _ = v152
            } else {
                switch "hello" {
                case "foo bar":
                    var v153 uint16 = 588
                    _ = v153
                case "\n\t":
                    var v154 uint32 = 1
                    _ = v154
                }
                fmt.Println(v105, v106)
                for fi155 := 0; fi155 < 13; fi155++ {
                    for fi156 := 0; fi156 < 17; fi156++ {
                        fmt.Println(fi156, fi155)
                        switch 42 {
                        case 128:
                            _ = len(v4)
                        case 0:
                            v106 += v105
                        }
                        switch "a" {
                        case "hello":
                            var v157 float32 = (v3 / v3)
                            _ = v157
                        default:
                            if false {
                                fmt.Println(fi155, v105)
                                if (0.1 >= v3) {
                                    fmt.Println(v3)
                                } else {
                                    fmt.Println(fi155)
                                }
                            } else {
                                v3 -= v3
                                var v158 int16 = 0
                                v3 += (1.5 * 0.1)
                                _ = v158
                            }
                        }
                        _ = fi156
                    }
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        switch (((("hello" + "") + ("" + ((("a" + "\n\t") + "foo bar") + "\n\t"))) + ("a" + "a")) + "foo bar") {
                        case "hello":
                            fmt.Println(v105)
                        case "a":
                            var v159 uint = 497
                            _ = v159
                        case "":
                            for fi160 := 0; fi160 < 8; fi160++ {
                                switch fi155 {
                                case 2147483647:
                                    switch "hello" {
                                    case "foo bar":
                                        v161 := GenS7{F0: 253, F1: 1000}
                                        _ = v161
                                    case "a":
                                        if (127 < 53) {
                                            if true {
                                                var v162 int16 = 1
                                                _ = v162
                                            } else {
                                                v163 := GenS3{F0: -34, F1: 187}
                                                v164 := GenS8{F0: fi155, F1: 1}
                                                _ = v163
                                                _ = v164
                                            }
                                            v3 = (v3 - v3)
                                            var v165 int32 = -1
                                            v165 = 2147483647
                                            _ = v165
                                        } else {
                                            var v166 int64 = 1
                                            v106 = (v106 | (fi160 ^ (fi155 + (fi155 + (fi155 ^ v105)))))
                                            _ = v166
                                        }
                                    }
                                default:
                                    v167 := GenS6{F0: 0}
                                    _ = v167
                                }
                                var v168 float32 = 1.5
                                var v169 uint8 = 42
                                var v170 int = v106
                                if (42 != 4294967295) {
                                    _ = len(v2)
                                    fmt.Println(v105, v170)
                                    if true {
                                        if true {
                                            switch v170 {
                                            case 0:
                                                var v171 bool = false
                                                _ = v171
                                            case 2147483647:
                                                v172 := GenS12{F0: v105, F1: true, F2: v169}
                                                _ = v172
                                            case -160:
                                                var v173 uint16 = 2
                                                var v174 uint64 = 970
                                                if false {
                                                    if ("foo bar" <= "a") {
                                                        if (false || false) {
                                                            v175 := GenS10{F0: "hello", F1: 2147483647, F2: ((v169 ^ v169) >= v169)}
                                                            var v176 float32 = v3
                                                            v177 := GenS13{F0: 1, F1: "hello"}
                                                            _ = v175
                                                            _ = v176
                                                            _ = v177
                                                        }
                                                        var v178 int64 = 42
                                                        v179 := GenS7{F0: -284, F1: 1}
                                                        _ = v178
                                                        _ = v179
                                                    } else {
                                                        switch ("hello" + "") {
                                                        case "":
                                                            v180 := GenS14{F0: 41, F1: 2147483648}
                                                            _ = v180
                                                        case "foo bar":
                                                            var v181 uint32 = 0
                                                            _ = v181
                                                        case "a":
                                                            v182 := GenS5{F0: uint8(v173)}
                                                            _ = v182
                                                        default:
                                                            v106 -= (fi160 + (v106 ^ 4294967295))
                                                            v183 := GenS15{F0: 683, F1: "foo bar"}
                                                            v184 := GenS7{F0: 0, F1: 0}
                                                            _ = v183
                                                            _ = v184
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v170, v168)
                                                    var v185 uint = uint(v174)
                                                    _ = v185
                                                }
                                                _ = len(v4)
                                                if true {
                                                    delete(v4, "a")
                                                } else {
                                                    var v186 int16 = 1000
                                                    _ = v186
                                                }
                                                v3 = -2.25
                                                fmt.Println(v3)
                                                _ = v173
                                                _ = v174
                                            default:
                                                v187 := GenS2{F0: -150}
                                                _ = v187
                                            }
                                            var v188 int16 = 0
                                            var v189 uint = 382
                                            _ = v188
                                            _ = v189
                                        }
                                        var v190 string = "\n\t"
                                        var v191 uint8 = (^(v169))
                                        _ = v190
                                        _ = v191
                                    }
                                    v168 += -2.25
                                } else {
                                    var v192 int = v106
                                    _ = v192
                                }
                                _ = fi160
                                _ = v168
                                _ = v169
                                _ = v170
                            }
                        default:
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v193 uint = 4294967295
                                v194 := GenS4{F0: (v3 != 0.0)}
                                _ = v193
                                _ = v194
                            } else {
                                v105 = fi155
                            }
                        }
                        v3 = -2.25
                    }
                    v105 = v105
                    _ = fi155
                }
            }
        case "foo bar":
            fmt.Println(v106)
        }
    }
    v105 = (v106 & (-(v106)))
    for fi195 := 0; fi195 < 6; fi195++ {
        for fi196 := 0; fi196 < 10; fi196++ {
            v105 = 2147483647
            _ = fi196
        }
        v3 = (1e10 + v3)
        _ = fi195
    }
    switch "a" {
    case "\n\t":
        v3 += float32(0.1)
    default:
        v3 = v3
    }
    v3 += v3
    switch (v106 / ((32768) | 1)) {
    case 4294967295:
        var v197 int32 = -360
        _ = v197
    case -25:
        var v198 int = 2147483647
        _ = v198
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
