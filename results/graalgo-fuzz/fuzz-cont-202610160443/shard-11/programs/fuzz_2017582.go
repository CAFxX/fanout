package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 string; F1 string; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 bool; F2 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int; F1 int8; F2 int8 }

func main() {
    v1 := []uint32{2147483647, 854, 1, 65535}
    v2 := make(map[string]uint32)
    var v3 uint16 = 42
    v4 := make(map[int]string)
    var v5 uint8 = 255
    var v6 uint8 = uint8(v3)
    if true {
        if true {
            var v7 int8 = 46
            var v8 int64 = -9223372036854775808
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 42, F1: 65535}
            v1 = append(v1, 775)
            for fi10 := 0; fi10 < 16; fi10++ {
                if true {
                    v11 := GenS2{F0: "foo bar", F1: (("" + "hello") + "hello"), F2: 0}
                    var v12 int16 = -1
                    if (true || false) {
                        if (false || true) {
                            if true {
                                var v13 uint64 = 433
                                v13 -= (v13 % (((v13 & v13)) | 1))
                                if false {
                                    if false {
                                        v13 -= (^(v13))
                                    }
                                    if (v5 == v6) {
                                        for fi14 := 0; fi14 < 14; fi14++ {
                                            if (false && (((!false) || (false && false)) && true)) {
                                                v13 = v13
                                            } else {
                                                v15 := GenS3{F0: -788}
                                                _ = v15
                                            }
                                            if (false || (522 == 42)) {
                                                fmt.Println(fi14)
                                            } else {
                                                v13 -= v13
                                                fmt.Println(fi10)
                                            }
                                            v16 := GenS2{F0: (("a" + "") + "\n\t"), F1: "\n\t", F2: 42}
                                            _ = fi14
                                            _ = v16
                                        }
                                        v12 = (^(v12))
                                        v17 := GenS3{F0: -618}
                                        _ = v17
                                    }
                                    switch "\n\t" {
                                    case "":
                                        var v18 uint32 = 42
                                        _ = v18
                                    case "hello":
                                        switch "hello" {
                                        case "":
                                            var v19 int8 = 74
                                            _ = v19
                                        }
                                    case "foo bar":
                                        switch fi10 {
                                        case 644:
                                            v20 := GenS2{F0: ("\n\t" + "hello"), F1: (((("" + "foo bar") + "a") + ("foo bar" + "foo bar")) + "foo bar"), F2: 11}
                                            _ = v20
                                        }
                                    }
                                }
                                _ = v13
                            } else {
                                _ = len(v4)
                                var v21 uint16 = 2
                                switch (fi10 ^ fi10) {
                                case 42:
                                    v3 = (v3 / ((v21) | 1))
                                case 32767:
                                    fmt.Println(fi10)
                                case -292:
                                    v2[""] = 1
                                default:
                                    v12 += v12
                                }
                                _ = v21
                            }
                            var v22 int64 = -1
                            switch "a" {
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v23 := GenS4{F0: -1}
                                _ = v23
                            }
                            _ = v22
                        } else {
                            var v24 int8 = 0
                            v25 := GenS5{F0: -1}
                            fmt.Println(v12)
                            _ = v24
                            _ = v25
                        }
                        v12 += v12
                    } else {
                        var v26 bool = (4294967295 != 511)
                        for fi27 := 0; fi27 < 12; fi27++ {
                            if len(v1) > 0 {
                                v1[0] = 393
                            }
                            fmt.Println(v12, v26, v26)
                            _ = fi27
                        }
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v6 += (v5 - v6)
                            var v29 uint32 = 398
                            _ = fi28
                            _ = v29
                        }
                        _ = v26
                    }
                    _ = v11
                    _ = v12
                }
                var v30 int = fi10
                var v31 int8 = 1
                _ = fi10
                _ = v30
                _ = v31
            }
            _ = v9
        }
        fmt.Println(v5, v5, v6)
        fmt.Println(v5, v3, v5)
    }
    v5 -= v6
    fmt.Println(v6)
    if (100 <= 127) {
        v1 = append(v1, 65535)
        fmt.Println(v6, v5)
        v32 := GenS6{F0: "foo bar"}
        _ = v32
    }
    v3 = 0
    for fi33 := 0; fi33 < 4; fi33++ {
        switch 128 {
        case -9223372036854775808:
            for fi34 := 0; fi34 < 8; fi34++ {
                v3 = 1
                v1 = append(v1, uint32(v6))
                _ = fi34
            }
        case 9223372036854775807:
            if false {
                if true {
                    if (!(!true)) {
                        if true {
                            if (!false) {
                                if true {
                                    var v35 int16 = 32767
                                    _ = v35
                                } else {
                                    var v36 int32 = 42
                                    _ = v36
                                }
                            } else {
                                v5 += v6
                                v1 = append(v1, 2147483648)
                            }
                            if ((v6 + 1) > (v6 + v5)) {
                                if true {
                                    switch -9223372036854775808 {
                                    case -9223372036854775808:
                                        delete(v4, (fi33 * (fi33 >> ((2) & 63))))
                                    case 2147483647:
                                        v3 = v3
                                    case -104:
                                        v5 += (v5 << ((18446744073709551615) & 63))
                                    default:
                                        v6 = 244
                                    }
                                    v37 := GenS7{F0: 42}
                                    _ = v37
                                } else {
                                    for fi38 := 0; fi38 < 4; fi38++ {
                                        v39 := GenS3{F0: 9223372036854775807}
                                        if false {
                                            if true {
                                                _ = len(v2)
                                                var v40 string = "\n\t"
                                                _ = v40
                                            }
                                            var v41 float64 = 0.0
                                            _ = v41
                                        } else {
                                            _ = len(v1)
                                            if ((true || true) && true) {
                                                v6 += v5
                                                if false {
                                                    fmt.Println(v5, v6)
                                                    var v42 bool = true
                                                    v43 := GenS4{F0: 664}
                                                    var v44 uint64 = 123
                                                    _ = v42
                                                    _ = v43
                                                    _ = v44
                                                }
                                                v1 = append(v1, 1)
                                            } else {
                                                v5 -= (v5 * 129)
                                                switch "hello" {
                                                case "a":
                                                    v45 := GenS1{F0: 670, F1: 559}
                                                    _ = v45
                                                case "\n\t":
                                                    v46 := GenS1{F0: 42, F1: 4294967295}
                                                    _ = v46
                                                case "foo bar":
                                                    v6 += 223
                                                default:
                                                    _ = len(v1)
                                                }
                                                v47 := GenS1{F0: 0, F1: 2147483648}
                                                _ = len(v2)
                                                v48 := GenS7{F0: 65535}
                                                if len(v1) > 0 {
                                                    v1[0] = 2
                                                }
                                                v49 := GenS8{F0: (false || (1 != 42)), F1: true, F2: 2}
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                        _ = len(v4)
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                                v50 := GenS6{F0: "hello"}
                                v51 := GenS9{F0: -35}
                                _ = v50
                                _ = v51
                            } else {
                                if false {
                                    var v52 uint8 = 42
                                    var v53 uint32 = 1
                                    _ = v52
                                    _ = v53
                                } else {
                                    fmt.Println(v3, v6, v3)
                                    v54 := GenS3{F0: int64(v5)}
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 13; fi55++ {
                                    v3 = (v3 / ((334) | 1))
                                    v56 := GenS2{F0: "foo bar", F1: "hello", F2: -128}
                                    _ = fi55
                                    _ = v56
                                }
                            }
                        } else {
                            v6 = 42
                            var v57 int16 = 1000
                            _ = v57
                        }
                        var v58 int32 = 2147483647
                        _ = len(v2)
                        _ = v58
                    } else {
                        for fi59 := 0; fi59 < 5; fi59++ {
                            var v60 int16 = -743
                            switch (fi33 / ((fi59) | 1)) {
                            case 32767:
                                v4[(fi59 - (fi59 - -428))] = "foo bar"
                            case 607:
                                v6 -= v6
                            }
                            fmt.Println(v5)
                            _ = fi59
                            _ = v60
                        }
                    }
                } else {
                    var v61 int64 = 0
                    v62 := GenS2{F0: ("" + "foo bar"), F1: "a", F2: -125}
                    _ = v61
                    _ = v62
                }
                for fi63 := 0; fi63 < 20; fi63++ {
                    switch "" {
                    case "\n\t":
                        switch (^(fi63)) {
                        case 42:
                            v6 -= (v6 & v5)
                        case 32767:
                            var v64 string = ("foo bar" + "foo bar")
                            _ = v64
                        default:
                            var v65 uint16 = 2
                            _ = v65
                        }
                    case "":
                        v6 = v5
                    case "foo bar":
                        v66 := GenS7{F0: 910}
                        _ = v66
                    }
                    _ = fi63
                }
            }
        case 541:
            var v67 bool = (!(!false))
            _ = v67
        }
        switch "hello" {
        case "":
            v68 := GenS9{F0: 127}
            _ = v68
        case "foo bar":
            for fi69 := 0; fi69 < 9; fi69++ {
                fmt.Println(v3)
                if (42 != (v3 << ((613) & 63))) {
                    v6 = 2
                    v6 = (v5 * (v6 | (v5 ^ 255)))
                    if false {
                        v70 := GenS10{F0: uint8(v3)}
                        switch -9223372036854775808 {
                        case 0:
                            v5 -= (^(v5))
                        case 222:
                            fmt.Println(fi33, v5, v5)
                            v71 := GenS7{F0: 42}
                            _ = v71
                        case 9223372036854775807:
                            if (false || false) {
                                v72 := GenS11{F0: true}
                                switch fi69 {
                                case -9223372036854775808:
                                    var v73 string = ""
                                    _ = v73
                                case -206:
                                    v6 = 127
                                default:
                                    v4[fi69] = "hello"
                                }
                                v74 := GenS7{F0: uint32(v6)}
                                _ = len(v2)
                                _ = v72
                                _ = v74
                            }
                        default:
                            var v75 int32 = 42
                            _ = v75
                        }
                        _ = v70
                    } else {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(fi33)
                        case "hello":
                            if true {
                                var v76 float32 = -2.25
                                _ = v76
                            }
                        case "":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v77 int16 = -32768
                        var v78 uint64 = 4294967295
                        _ = v77
                        _ = v78
                    }
                }
                switch (fi69 >> ((255) & 63)) {
                case 1:
                    var v79 uint8 = v5
                    _ = v79
                case 32768:
                    var v80 int32 = 816
                    _ = v80
                }
                _ = fi69
            }
        default:
            var v81 float64 = -2.25
            _ = v81
        }
        v5 = (v6 << ((2) & 63))
        _ = fi33
    }
    v5 += v5
    v82 := GenS12{F0: (^(v5))}
    for fi83 := 0; fi83 < 7; fi83++ {
        switch fi83 {
        case -9223372036854775808:
            var v84 int = (fi83 ^ (fi83 + 127))
            _ = v84
        case 0:
            var v85 float32 = ((float32(0.0) / 3.14159) - -2.25)
            _ = v85
        default:
            for fi86 := 0; fi86 < 6; fi86++ {
                if false {
                    v87 := GenS1{F0: -437, F1: 2147483647}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (!true) {
                        var v88 uint = 1
                        v89 := GenS13{F0: 1}
                        _ = v88
                        _ = v89
                    }
                    _ = v87
                }
                if true {
                    if true {
                        fmt.Println(v3)
                    }
                    var v90 int64 = 825
                    if (!true) {
                        switch (-(fi86)) {
                        case 791:
                            var v91 uint16 = 2
                            _ = v91
                        case -2147483648:
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                        default:
                            var v92 uint8 = 127
                            _ = v92
                        }
                        var v93 string = (("a" + ("hello" + "")) + "foo bar")
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v93
                    }
                    _ = v90
                } else {
                    fmt.Println(fi83, fi86)
                }
                _ = fi86
            }
        }
        if false {
            v3 = (v3 % ((v3) | 1))
            for fi94 := 0; fi94 < 3; fi94++ {
                if (0.1 == ((float64(1e10) * ((1.5 * (-2.25 / (1.5 + 0.1))) - 1.5)) / ((float64(1e10) - -2.25) - 0.1))) {
                    v95 := v2[("hello" + "foo bar")]
                    _ = v95
                }
                var v96 float32 = ((1e10 + 3.14159) - 1.5)
                if false {
                    var v97 int = 973
                    v6 -= (v5 >> ((1) & 63))
                    var v98 uint32 = 1
                    if true {
                        v97 = 172
                        if true {
                            v99 := GenS1{F0: 2147483647, F1: (v98 << ((42) & 63))}
                            _ = v99
                        }
                    }
                    switch "hello" {
                    case "a":
                        var v100 string = ("foo bar" + "")
                        if false {
                            v101 := GenS14{F0: -269, F1: 127, F2: -97}
                            _ = v101
                        } else {
                            fmt.Println(v5)
                            fmt.Println(fi94, fi94, v3)
                        }
                        _ = v100
                    }
                    v102 := GenS2{F0: "\n\t", F1: ("foo bar" + "hello"), F2: 42}
                    _ = v102
                    _ = v97
                    _ = v98
                } else {
                    if true {
                        fmt.Println(v96)
                        var v103 uint = 59
                        _ = len(v2)
                        _ = v103
                    } else {
                        var v104 uint = 1
                        var v105 uint64 = 1
                        _ = v104
                        _ = v105
                    }
                }
                _ = fi94
                _ = v96
            }
        } else {
            v3 = 32767
            fmt.Println(v5, v6)
            switch fi83 {
            case 9223372036854775807:
                var v106 bool = true
                _ = v106
            default:
                switch (^(fi83)) {
                case 9223372036854775807:
                    v5 = v5
                case 32767:
                    var v107 uint16 = (v3 * v3)
                    _ = v107
                }
            }
        }
        _ = fi83
    }
    var v108 int = 32767
    switch "foo bar" {
    case "\n\t":
        if false {
            v108 = int(v3)
            if false {
                if true {
                    fmt.Println(v108, v108, v5)
                    fmt.Println(v6, v6)
                    fmt.Println(v108, v6, v3)
                }
            }
        } else {
            delete(v4, -638)
            var v109 uint32 = 1
            _ = v109
        }
    case "a":
        var v110 uint8 = v5
        _ = v110
    case "foo bar":
        switch (("foo bar" + "hello") + (("foo bar" + ("\n\t" + "\n\t")) + "a")) {
        case "foo bar":
            if (false && (false && ((v3 + (v3 << ((0) & 63))) < v3))) {
                switch v108 {
                case 127:
                    v5 = (v5 + v5)
                case 2147483647:
                    var v111 int = 42
                    _ = v111
                case -2147483648:
                    var v112 int32 = -2147483648
                    _ = v112
                default:
                    var v113 int16 = 955
                    _ = v113
                }
                switch (v108 + (^(v108))) {
                case 138:
                    if false {
                        var v114 uint8 = v6
                        v6 = (^(v6))
                        _ = v114
                    }
                case -61:
                    fmt.Println(v108)
                case 128:
                    var v115 int64 = -1
                    _ = v115
                }
                var v116 uint32 = 866
                _ = v116
            } else {
                fmt.Println(v5, v108)
            }
        default:
            v117 := v2["foo bar"]
            _ = v117
        }
    }
    v5 += (^(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v108:", v108)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
}
