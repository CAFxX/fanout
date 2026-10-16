package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 string; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS6 struct { F0 int8; F1 string; F2 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32; F1 int64; F2 uint32 }
type GenS12 struct { F0 uint16; F1 string; F2 uint16 }
type GenS13 struct { F0 bool; F1 uint16 }
type GenS14 struct { F0 int; F1 uint32; F2 bool }
type GenS15 struct { F0 int32; F1 uint8 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 string; F1 int16 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint8; F1 uint32; F2 int }
type GenS20 struct { F0 int32; F1 int16; F2 int8 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 string; F1 uint; F2 uint8 }
type GenS23 struct { F0 bool; F1 int64; F2 bool }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 string; F1 int32 }
type GenS27 struct { F0 uint16; F1 int; F2 int16 }
type GenS28 struct { F0 int64 }
type GenS29 struct { F0 int16; F1 uint64 }
type GenS30 struct { F0 uint64; F1 int16 }
type GenS31 struct { F0 bool; F1 uint8 }
type GenS32 struct { F0 int; F1 int8 }
type GenS33 struct { F0 uint8 }
type GenS34 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS35 struct { F0 uint32; F1 uint8 }
type GenS36 struct { F0 bool }
type GenS37 struct { F0 int; F1 uint }
type GenS38 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    var v3 uint = 42
    var v4 int16 = 597
    var v5 uint16 = 11
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{-2147483648, 32767, fi6}
        if false {
            for fi8 := 0; fi8 < 18; fi8++ {
                if (!false) {
                    if true {
                        v7 = append(v7, fi6)
                        var v9 string = ""
                        v7 = append(v7, fi6)
                        v2[761] = (("\n\t" + v9) + (v9 + v9))
                        _ = v9
                    } else {
                        if (2 >= 127) {
                            var v10 string = "a"
                            _ = v10
                        } else {
                            v4 = -1
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        v11 := GenS1{F0: ("\n\t" >= ""), F1: 504}
                        _ = v11
                    case "":
                        v12 := GenS2{F0: ("a" + "foo bar"), F1: -25}
                        _ = v12
                    case "a":
                        v13 := GenS1{F0: false, F1: -9223372036854775808}
                        var v14 int8 = 1
                        _ = v13
                        _ = v14
                    }
                    var v15 uint64 = 42
                    _ = v15
                }
                v3 += 0
                v16 := GenS2{F0: "hello", F1: int8(v5)}
                _ = fi8
                _ = v16
            }
            var v17 int32 = 42
            v17 -= (v17 ^ v17)
            _ = v17
        }
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    switch -1 {
    case -703:
        fmt.Println(v3)
    case 42:
        for fi18 := 0; fi18 < 15; fi18++ {
            v4 = (v4 >> ((18446744073709551615) & 63))
            var v19 uint = v3
            v20 := GenS3{F0: 65535, F1: (v5 >> (((v3 + 255)) & 63)), F2: 236}
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        var v21 uint8 = 0
        _ = v21
    }
    switch ("hello" + ("foo bar" + "hello")) {
    case "a":
        for fi22 := 0; fi22 < 10; fi22++ {
            v4 = v4
            _ = fi22
        }
    case "hello":
        v23 := GenS3{F0: (v5 + v5), F1: (v5 / (((^(v5))) | 1)), F2: 4294967295}
        _ = v23
    case "\n\t":
        v24 := GenS2{F0: (("foo bar" + "\n\t") + ("" + "\n\t")), F1: 92}
        _ = v24
    }
    fmt.Println(v4)
    if (65535 > 65535) {
        for fi25 := 0; fi25 < 9; fi25++ {
            fmt.Println(v4, v5)
            v26 := GenS2{F0: "\n\t", F1: 42}
            var v27 uint16 = v5
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        if (42 == 948) {
            var v28 uint16 = (v5 << ((64) & 63))
            for fi29 := 0; fi29 < 19; fi29++ {
                v3 = (v3 / (((v3 << ((42) & 63))) | 1))
                var v30 int8 = -77
                _ = fi29
                _ = v30
            }
            fmt.Println(v5, v28, v4)
            _ = v28
        }
        fmt.Println(v5, v3, v3)
    }
    v3 += 4294967295
    v31 := []int{-651, 127, 1, 42}
    switch (("foo bar" + ("" + "\n\t")) + "foo bar") {
    case "\n\t":
        v3 -= v3
    case "a":
        switch 0 {
        case 0:
            var v32 int = -290
            _ = v32
        }
    case "hello":
        v3 -= (v3 / (((^(v3))) | 1))
    default:
        switch "a" {
        case "\n\t":
            if (1.5 <= ((float64(float32((1e10 - float64(-2.25)))) * (-2.25 / 1e10)) / 1.5)) {
                fmt.Println(v3, v3)
                var v33 int16 = v4
                v34 := GenS4{F0: "\n\t", F1: 127, F2: v5}
                _ = v33
                _ = v34
            } else {
                delete(v2, 2147483648)
            }
        case "hello":
            switch "" {
            case "foo bar":
                var v35 uint16 = 2
                _ = v35
            case "a":
                if (-686 >= 9223372036854775807) {
                    var v36 int32 = -558
                    var v37 string = ("a" + ("foo bar" + ""))
                    if ("" > ((v37 + v37) + (((v37 + v37) + "a") + "foo bar"))) {
                        switch v37 {
                        case "\n\t":
                            v37 = (v37 + v37)
                        }
                    } else {
                        for fi38 := 0; fi38 < 11; fi38++ {
                            v39 := GenS3{F0: v5, F1: (v5 ^ v5), F2: 686}
                            if false {
                                v40 := v1[""]
                                v41 := GenS2{F0: "a", F1: -27}
                                var v42 uint = 2
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            switch (v37 + v37) {
                            case "foo bar":
                                switch (v37 + "") {
                                case "\n\t":
                                    v43 := GenS5{F0: 1, F1: -327, F2: v5}
                                    _ = v43
                                case "foo bar":
                                    if false {
                                        v3 = 30
                                        switch ("a" + v37) {
                                        case "":
                                            var v44 int64 = 669
                                            _ = v44
                                        default:
                                            if false {
                                                for fi45 := 0; fi45 < 6; fi45++ {
                                                    v46 := GenS2{F0: v37, F1: -128}
                                                    _ = fi45
                                                    _ = v46
                                                }
                                                fmt.Println(v5, v5)
                                            }
                                        }
                                    } else {
                                        v47 := GenS6{F0: -1, F1: v37, F2: 370}
                                        if ("hello" >= v37) {
                                            if true {
                                                var v48 bool = (v37 > "\n\t")
                                                switch fi38 {
                                                case -1:
                                                    v49 := v2[4294967295]
                                                    _ = v49
                                                default:
                                                    fmt.Println(v37, fi38)
                                                }
                                                _ = v48
                                            } else {
                                                switch (v37 + "") {
                                                case "\n\t":
                                                    for fi50 := 0; fi50 < 11; fi50++ {
                                                        var v51 uint8 = 0
                                                        switch v37 {
                                                        case "hello":
                                                            var v52 uint8 = v51
                                                            v3 = v3
                                                            _ = v52
                                                        }
                                                        if (false || false) {
                                                            v5 = 179
                                                        } else {
                                                            v31 = append(v31, -367)
                                                        }
                                                        _ = fi50
                                                        _ = v51
                                                    }
                                                case "foo bar":
                                                    v53 := GenS3{F0: v5, F1: 0, F2: 154}
                                                    _ = v53
                                                default:
                                                    var v54 uint32 = uint32(v5)
                                                    _ = v54
                                                }
                                                if true {
                                                    for fi55 := 0; fi55 < 4; fi55++ {
                                                        if false {
                                                            delete(v1, "foo bar")
                                                            v56 := GenS1{F0: (((v36 / ((-2147483648) | 1)) != (v36 ^ v36)) || true), F1: -9223372036854775808}
                                                            v57 := GenS7{F0: 42}
                                                            _ = v56
                                                            _ = v57
                                                        } else {
                                                            var v58 int8 = 0
                                                            _ = v58
                                                        }
                                                        var v59 float64 = 1.5
                                                        _ = fi55
                                                        _ = v59
                                                    }
                                                } else {
                                                    if true {
                                                        v60 := GenS3{F0: (v5 - (v5 / ((v5) | 1))), F1: (^(v5)), F2: 2147483647}
                                                        for fi61 := 0; fi61 < 8; fi61++ {
                                                            var v62 int32 = 1
                                                            v5 += 932
                                                            if false {
                                                                if len(v31) >= 2 {
                                                                    v31 = v31[1:len(v31)]
                                                                }
                                                                v63 := v1[v37]
                                                                _ = v63
                                                            }
                                                            _ = fi61
                                                            _ = v62
                                                        }
                                                        _ = v60
                                                    }
                                                    switch "a" {
                                                    case "\n\t":
                                                        v2[1] = ("" + v37)
                                                    }
                                                    v64 := GenS1{F0: true, F1: 0}
                                                    _ = v64
                                                }
                                            }
                                            for fi65 := 0; fi65 < 16; fi65++ {
                                                v36 += (v36 + v36)
                                                if true {
                                                    v3 = (^(v3))
                                                    v2[0] = v37
                                                    var v66 int8 = 0
                                                    if true {
                                                        var v67 int16 = v4
                                                        _ = v67
                                                    } else {
                                                        var v68 float32 = float32(((1.5 - 0.0) / (float64(3.14159) * float64((0.1 + (-2.25 - (0.1 / 0.1)))))))
                                                        switch fi65 {
                                                        case -519:
                                                            if (((((("" + ((((v37 + ((((v37 + "foo bar") + v37) + ((v37 + (v37 + ("" + "hello"))) + "foo bar")) + v37)) + "foo bar") + ("foo bar" + "")) + v37)) + ((v37 + "hello") + v37)) + (v37 + (v37 + "foo bar"))) + "hello") + "hello") == v37) {
                                                                fmt.Println(fi38, v36)
                                                                switch v37 {
                                                                case "":
                                                                    var v69 float64 = 1e10
                                                                    fmt.Println(v3, v3, v37)
                                                                    _ = v69
                                                                case "foo bar":
                                                                    v70 := GenS1{F0: false, F1: -9223372036854775808}
                                                                    _ = v70
                                                                case "a":
                                                                    if true {
                                                                        v3 += (v3 - v3)
                                                                        fmt.Println(fi38, v3, v5)
                                                                        var v71 uint16 = (^(v5))
                                                                        _ = v71
                                                                    }
                                                                default:
                                                                    v72 := GenS5{F0: 2, F1: 90, F2: (v5 + v5)}
                                                                    fmt.Println(v36)
                                                                    v3 = 4294967295
                                                                    v4 = (v4 / (((v4 % (((v4 - v4)) | 1))) | 1))
                                                                    _ = v72
                                                                }
                                                                switch fi65 {
                                                                case 194:
                                                                    v73 := GenS3{F0: v5, F1: v5, F2: 65535}
                                                                    _ = v73
                                                                case 42:
                                                                    if true {
                                                                        switch (fi38 - 0) {
                                                                        case -1:
                                                                            fmt.Println(fi38, v37)
                                                                            v37 = ""
                                                                        case 128:
                                                                            v74 := GenS8{F0: (v66 % ((1) | 1)), F1: 0}
                                                                            _ = v74
                                                                        default:
                                                                            var v75 int32 = int32(v5)
                                                                            _ = v75
                                                                        }
                                                                        v76 := GenS7{F0: 18446744073709551615}
                                                                        switch ("a" + ((("hello" + ((v37 + (((("foo bar" + v37) + "") + v37) + (("foo bar" + v37) + v37))) + ((v37 + ((v37 + ((v37 + "foo bar") + v37)) + ((v37 + "hello") + (((("foo bar" + "\n\t") + (v37 + v37)) + v37) + v37)))) + v37))) + "") + ("\n\t" + (v37 + ("a" + "hello"))))) {
                                                                        case "hello":
                                                                            _ = len(v1)
                                                                        case "":
                                                                            v31 = append(v31, int(v66))
                                                                        }
                                                                        _ = v76
                                                                    }
                                                                case 32767:
                                                                    var v77 int16 = 1
                                                                    _ = v77
                                                                default:
                                                                    v78 := GenS9{F0: 1, F1: 42}
                                                                    _ = v78
                                                                }
                                                            } else {
                                                                v1["\n\t"] = ("\n\t" + "a")
                                                                switch v37 {
                                                                case "foo bar":
                                                                    switch (fi65 >> ((v3) & 63)) {
                                                                    case 128:
                                                                        v66 = (-(v66))
                                                                    default:
                                                                        var v79 bool = (1.5 < 1.5)
                                                                        _ = v79
                                                                    }
                                                                case "hello":
                                                                    v5 = v5
                                                                case "":
                                                                    v80 := GenS10{F0: 674}
                                                                    _ = v80
                                                                default:
                                                                    var v81 int = (fi38 % (((fi38 << ((v3) & 63))) | 1))
                                                                    _ = v81
                                                                }
                                                            }
                                                        case 0:
                                                            fmt.Println(v4)
                                                        }
                                                        v5 -= v5
                                                        if false {
                                                            var v82 string = ""
                                                            var v83 int = 32768
                                                            if len(v31) >= 2 {
                                                                v31 = v31[1:len(v31)]
                                                            }
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        fmt.Println(fi38, v66, v5)
                                                        _ = v68
                                                    }
                                                    if len(v31) > 0 {
                                                        v31[0] = (fi38 % ((fi65) | 1))
                                                    }
                                                    _ = v66
                                                } else {
                                                    var v84 int = (fi65 + fi65)
                                                    if false {
                                                        var v85 int8 = 42
                                                        fmt.Println(v85, v5, fi65)
                                                        if true {
                                                            if true {
                                                                v85 += (v85 / ((v85) | 1))
                                                            } else {
                                                                var v86 uint64 = 456
                                                                _ = v86
                                                            }
                                                            if len(v31) > 0 {
                                                                v31[0] = fi38
                                                            }
                                                            v1[v37] = v37
                                                        }
                                                        v87 := GenS3{F0: v5, F1: v5, F2: 420}
                                                        var v88 uint16 = (v5 & (v5 >> (((^(v3))) & 63)))
                                                        switch (v37 + v37) {
                                                        case "\n\t":
                                                            var v89 bool = ((v3 & v3) > (v3 << ((uint(v5)) & 63)))
                                                            _ = v89
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                        v85 -= v85
                                                        var v90 int32 = -1
                                                        v37 = (v37 + v37)
                                                        v5 += 42
                                                        _ = v85
                                                        _ = v87
                                                        _ = v88
                                                        _ = v90
                                                    }
                                                    _ = v84
                                                }
                                                _ = fi65
                                            }
                                            var v91 uint64 = 18446744073709551615
                                            _ = v91
                                        } else {
                                            v5 -= (v5 | v5)
                                            v37 = v37
                                            if false {
                                                for fi92 := 0; fi92 < 16; fi92++ {
                                                    var v93 int8 = -5
                                                    if ((v93 ^ v93) <= (v93 | -40)) {
                                                        var v94 string = ("" + v37)
                                                        fmt.Println(v37, v94, fi92)
                                                        v95 := GenS3{F0: v5, F1: (v5 | (^(v5))), F2: 0}
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    _ = fi92
                                                    _ = v93
                                                }
                                            }
                                        }
                                        _ = v47
                                    }
                                case "":
                                    if true {
                                        v96 := GenS11{F0: v36, F1: int64(v4), F2: 65535}
                                        _ = v96
                                    }
                                }
                            }
                            _ = fi38
                            _ = v39
                        }
                        v97 := GenS8{F0: 0, F1: v3}
                        for fi98 := 0; fi98 < 1; fi98++ {
                            if true {
                                v37 = ("foo bar" + v37)
                                if false {
                                    v36 -= -1
                                    if (529 != (v4 >> (((v3 | v3)) & 63))) {
                                        v2[25] = (v37 + "a")
                                        fmt.Println(v37, v36, v36)
                                        var v99 uint16 = (v5 - (^(v5)))
                                        _ = v99
                                    }
                                    if (!(!(1 <= 42))) {
                                        var v100 float32 = (3.14159 * 0.1)
                                        _ = v100
                                    }
                                }
                            }
                            v36 = (v36 / ((2147483647) | 1))
                            v101 := GenS12{F0: (^(v5)), F1: "\n\t", F2: 0}
                            _ = fi98
                            _ = v101
                        }
                        _ = v97
                    }
                    _ = v36
                    _ = v37
                }
            }
        default:
            var v102 int32 = -2147483648
            _ = v102
        }
    }
    if true {
        for fi103 := 0; fi103 < 14; fi103++ {
            if true {
                if len(v31) > 0 {
                    v31[0] = (fi103 << (((v3 * v3)) & 63))
                }
                v3 = (v3 | (v3 & (v3 * v3)))
                v1[""] = ("\n\t" + "\n\t")
            }
            if (((!false) && (v3 >= 1)) || true) {
                var v104 int64 = 9223372036854775807
                var v105 uint = (v3 / ((255) | 1))
                v2[577] = "\n\t"
                _ = v104
                _ = v105
            } else {
                v106 := GenS2{F0: ("\n\t" + "\n\t"), F1: -128}
                delete(v1, "hello")
                var v107 string = ""
                _ = v106
                _ = v107
            }
            if len(v31) > 0 {
                v31[0] = 42
            }
            _ = fi103
        }
        var v108 int8 = -88
        _ = v108
    } else {
        v5 = v5
        _ = len(v2)
        if (2147483648 <= 42) {
            v109 := GenS13{F0: true, F1: 784}
            if (!false) {
                switch int(v4) {
                case 1:
                    var v110 uint8 = 255
                    _ = v110
                default:
                    if ("foo bar" <= ("a" + "hello")) {
                        fmt.Println(v3, v3, v4)
                        v3 = (^(v3))
                        if false {
                            if false {
                                v4 = 1000
                                v111 := GenS14{F0: 884, F1: 96, F2: (!true)}
                                var v112 int8 = -1
                                _ = v111
                                _ = v112
                            } else {
                                _ = len(v31)
                                fmt.Println(v3, v4, v3)
                            }
                            var v113 int64 = 1
                            if (!true) {
                                if true {
                                    var v114 int32 = -1
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    _ = v114
                                } else {
                                    for fi115 := 0; fi115 < 20; fi115++ {
                                        v5 = (v5 | 0)
                                        var v116 int = int(v3)
                                        switch v116 {
                                        case 4294967295:
                                            v117 := GenS15{F0: 42, F1: 218}
                                            _ = v117
                                        }
                                        _ = fi115
                                        _ = v116
                                    }
                                }
                                if false {
                                    if false {
                                        v113 += (v113 + v113)
                                        v118 := GenS16{F0: v5}
                                        _ = v118
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            for fi119 := 0; fi119 < 13; fi119++ {
                                                switch "hello" {
                                                case "a":
                                                    v3 -= v3
                                                case "hello":
                                                    fmt.Println(fi119, v4, v113)
                                                case "foo bar":
                                                    v1["a"] = ("\n\t" + "\n\t")
                                                default:
                                                    if true {
                                                        var v120 int = (fi119 >> ((65535) & 63))
                                                        _ = len(v2)
                                                        for fi121 := 0; fi121 < 17; fi121++ {
                                                            delete(v1, ("hello" + "foo bar"))
                                                            var v122 int8 = -1
                                                            _ = fi121
                                                            _ = v122
                                                        }
                                                        _ = v120
                                                    } else {
                                                        switch "hello" {
                                                        case "hello":
                                                            v123 := v1["hello"]
                                                            _ = v123
                                                        }
                                                    }
                                                }
                                                _ = fi119
                                            }
                                        default:
                                            v124 := GenS1{F0: true, F1: v113}
                                            _ = v124
                                        }
                                        v4 = v4
                                        v125 := GenS15{F0: -341, F1: 255}
                                        _ = v125
                                    }
                                    var v126 uint8 = 42
                                    v126 = uint8(v5)
                                    _ = v126
                                } else {
                                    if (true && (!(((^(v5)) < v5) || false))) {
                                        var v127 uint = v3
                                        v5 -= (^(v5))
                                        _ = v127
                                    } else {
                                        switch ("\n\t" + "a") {
                                        case "a":
                                            v128 := GenS17{F0: "\n\t", F1: 32767}
                                            _ = v128
                                        case "\n\t":
                                            if len(v31) > 0 {
                                                v31[0] = 32768
                                            }
                                        case "hello":
                                            for fi129 := 0; fi129 < 1; fi129++ {
                                                for fi130 := 0; fi130 < 16; fi130++ {
                                                    switch -2147483648 {
                                                    case 449:
                                                        v5 -= v5
                                                        var v131 int64 = v113
                                                        _ = v131
                                                    }
                                                    if (754 != 855) {
                                                        v132 := GenS7{F0: 4294967295}
                                                        if (fi130 <= fi130) {
                                                            v3 -= v3
                                                            if (((!(false || true)) || true) && (733 != 42)) {
                                                                if (0 >= 42) {
                                                                    fmt.Println(fi129)
                                                                }
                                                            } else {
                                                                var v133 bool = false
                                                                _ = v133
                                                            }
                                                            if len(v31) >= 2 {
                                                                v31 = v31[1:len(v31)]
                                                            }
                                                        } else {
                                                            v134 := GenS5{F0: 1, F1: v113, F2: 2}
                                                            var v135 uint8 = 128
                                                            var v136 bool = (!(false || ((-(v113)) == v113)))
                                                            _ = v134
                                                            _ = v135
                                                            _ = v136
                                                        }
                                                        delete(v2, 4294967295)
                                                        _ = v132
                                                    } else {
                                                        switch (fi130 % ((fi130) | 1)) {
                                                        case -1:
                                                            var v137 float32 = (((1.5 / 1.0) - float32(1.5)) / 1e10)
                                                            _ = v137
                                                        case -298:
                                                            fmt.Println(v113, v5, fi130)
                                                        case 42:
                                                            v3 -= (v3 / (((v3 & (v3 & (v3 << ((v3) & 63))))) | 1))
                                                        default:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v138 int8 = 42
                                                                _ = v138
                                                            case "":
                                                                if false {
                                                                    fmt.Println(v113, fi130, fi130)
                                                                    if true {
                                                                        var v139 uint64 = 18446744073709551615
                                                                        var v140 uint32 = 2147483647
                                                                        if (("hello" + "foo bar") != "a") {
                                                                            var v141 uint64 = v139
                                                                            fmt.Println(v4)
                                                                            _ = v141
                                                                        } else {
                                                                            v140 += v140
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v142 int8 = int8(v140)
                                                                                _ = v142
                                                                            default:
                                                                                fmt.Println(v113, fi130)
                                                                            }
                                                                        }
                                                                        _ = v139
                                                                        _ = v140
                                                                    }
                                                                } else {
                                                                    v31 = append(v31, (fi130 + (fi129 * 32768)))
                                                                    v143 := GenS5{F0: 127, F1: (v113 ^ v113), F2: (v5 - v5)}
                                                                    _ = v143
                                                                }
                                                                v4 = -1
                                                            case "hello":
                                                                v5 += uint16(v3)
                                                                fmt.Println(v3, v3, v4)
                                                                v5 = (v5 | 65535)
                                                            default:
                                                                var v144 uint = v3
                                                                _ = v144
                                                            }
                                                        }
                                                    }
                                                    var v145 bool = ((fi129 - 2147483647) >= (fi129 ^ (-(fi130))))
                                                    _ = fi130
                                                    _ = v145
                                                }
                                                for fi146 := 0; fi146 < 12; fi146++ {
                                                    v147 := GenS7{F0: 2}
                                                    v5 = (^(v5))
                                                    switch fi129 {
                                                    case 127:
                                                        v148 := GenS15{F0: -722, F1: 0}
                                                        _ = v148
                                                    case 0:
                                                        delete(v1, "foo bar")
                                                        v4 = (-(v4))
                                                    case 32767:
                                                        v5 = (^(v5))
                                                    }
                                                    v4 += v4
                                                    _ = fi146
                                                    _ = v147
                                                }
                                                v149 := GenS3{F0: v5, F1: v5, F2: 2147483648}
                                                _ = fi129
                                                _ = v149
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                        if len(v31) >= 2 {
                                            v31 = v31[1:len(v31)]
                                        }
                                    }
                                    fmt.Println(v113)
                                    v150 := v2[851]
                                    _ = v150
                                }
                            }
                            _ = v113
                        } else {
                            for fi151 := 0; fi151 < 14; fi151++ {
                                fmt.Println(v3, v5, fi151)
                                fmt.Println(v3, v4, fi151)
                                for fi152 := 0; fi152 < 5; fi152++ {
                                    if false {
                                        switch fi152 {
                                        case 32768:
                                            switch ("hello" + ("foo bar" + "foo bar")) {
                                            case "\n\t":
                                                v153 := GenS4{F0: "a", F1: 10, F2: v5}
                                                _ = v153
                                            }
                                        default:
                                            fmt.Println(v4, v5, fi152)
                                            fmt.Println(v4, fi151)
                                        }
                                        switch (fi151 / (((fi152 >> (((^(v3))) & 63))) | 1)) {
                                        case 233:
                                            if false {
                                                if (!true) {
                                                    v3 = v3
                                                    fmt.Println(v4, fi152)
                                                    switch 2147483647 {
                                                    case 4294967295:
                                                        switch "" {
                                                        case "":
                                                            v4 -= v4
                                                        case "a":
                                                            var v154 uint = 963
                                                            _ = v154
                                                        case "foo bar":
                                                            switch ("a" + "foo bar") {
                                                            case "hello":
                                                                if true {
                                                                    if false {
                                                                        var v155 int = (fi152 ^ (^(fi152)))
                                                                        v3 = 255
                                                                        var v156 int8 = -56
                                                                        _ = v155
                                                                        _ = v156
                                                                    } else {
                                                                        var v157 float32 = 1e10
                                                                        var v158 uint8 = 99
                                                                        var v159 float32 = (v157 + (((1e10 * v157) + 0.0) - float32(0.0)))
                                                                        _ = v157
                                                                        _ = v158
                                                                        _ = v159
                                                                    }
                                                                    v160 := GenS18{F0: (-854 != 2147483647)}
                                                                    _ = v160
                                                                } else {
                                                                    if (255 >= 255) {
                                                                        v161 := GenS12{F0: v5, F1: "a", F2: (v5 * 32767)}
                                                                        _ = v161
                                                                    }
                                                                    if false {
                                                                        v3 -= (v3 ^ uint(v5))
                                                                        v5 = v5
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v162 int = fi152
                                                                _ = v162
                                                            case "":
                                                                switch 1 {
                                                                case -1:
                                                                    switch (fi151 & 2147483647) {
                                                                    case 940:
                                                                        fmt.Println(v5, v5, fi151)
                                                                    case 128:
                                                                        if true {
                                                                            v163 := GenS16{F0: v5}
                                                                            _ = v163
                                                                        }
                                                                    case 32768:
                                                                        v164 := GenS18{F0: (255 <= (v5 - (v5 << ((255) & 63))))}
                                                                        if (42 >= 0) {
                                                                            if false {
                                                                                v165 := GenS9{F0: 155, F1: int64(v4)}
                                                                                _ = v165
                                                                            } else {
                                                                                var v166 int32 = 0
                                                                                if false {
                                                                                    v166 = (v166 | (^(v166)))
                                                                                } else {
                                                                                    if false {
                                                                                        if (141 == (^(v5))) {
                                                                                            v167 := GenS19{F0: 235, F1: 66, F2: 4294967295}
                                                                                            if true {
                                                                                                if len(v31) >= 2 {
                                                                                                    v31 = v31[1:len(v31)]
                                                                                                }
                                                                                                v166 = 42
                                                                                                fmt.Println(v4, fi151, fi151)
                                                                                            }
                                                                                            if (v166 >= (v166 << ((42) & 63))) {
                                                                                                if (-899 != 35) {
                                                                                                    var v168 string = "foo bar"
                                                                                                    v3 += v3
                                                                                                    _ = v168
                                                                                                } else {
                                                                                                    var v169 int32 = v166
                                                                                                    var v170 int64 = -684
                                                                                                    _ = v169
                                                                                                    _ = v170
                                                                                                }
                                                                                                if true {
                                                                                                    v4 = v4
                                                                                                    v166 = v166
                                                                                                } else {
                                                                                                    v4 = (v4 >> ((v3) & 63))
                                                                                                    if (42 <= 4294967295) {
                                                                                                        if false {
                                                                                                            if (0.0 < (1e10 - 0.0)) {
                                                                                                                if false {
                                                                                                                    switch fi152 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        switch "foo bar" {
                                                                                                                        case "foo bar":
                                                                                                                            v31 = append(v31, -906)
                                                                                                                        case "a":
                                                                                                                            switch -170 {
                                                                                                                            case -64:
                                                                                                                                v166 = (-(v166))
                                                                                                                            }
                                                                                                                        case "\n\t":
                                                                                                                            if false {
                                                                                                                                var v171 uint8 = 218
                                                                                                                                switch "\n\t" {
                                                                                                                                case "hello":
                                                                                                                                    _ = len(v31)
                                                                                                                                }
                                                                                                                                var v172 uint8 = v171
                                                                                                                                v171 = (^(v172))
                                                                                                                                v173 := GenS5{F0: 0, F1: -1, F2: (v5 / ((0) | 1))}
                                                                                                                                _ = v171
                                                                                                                                _ = v172
                                                                                                                                _ = v173
                                                                                                                            } else {
                                                                                                                                var v174 int = -9223372036854775808
                                                                                                                                _ = v174
                                                                                                                            }
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        fmt.Println(v4, v4)
                                                                                                                    }
                                                                                                                    switch "" {
                                                                                                                    case "foo bar":
                                                                                                                        v175 := GenS16{F0: 255}
                                                                                                                        _ = v175
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            if (!(2 != 2147483648)) {
                                                                                                                v166 += (^(v166))
                                                                                                                var v176 float32 = float32(float64(3.14159))
                                                                                                                v31 = append(v31, fi152)
                                                                                                                if len(v31) >= 2 {
                                                                                                                    v31 = v31[1:len(v31)]
                                                                                                                }
                                                                                                                _ = v176
                                                                                                            }
                                                                                                        }
                                                                                                        var v177 int64 = -978
                                                                                                        _ = v177
                                                                                                    } else {
                                                                                                        var v178 int64 = 42
                                                                                                        if false {
                                                                                                            v3 = v3
                                                                                                            v5 += (v5 << ((v3) & 63))
                                                                                                            if (0 < -128) {
                                                                                                                var v179 float32 = (0.0 - float32((0.0 / 1e10)))
                                                                                                                if len(v31) >= 2 {
                                                                                                                    v31 = v31[1:len(v31)]
                                                                                                                }
                                                                                                                v166 -= v166
                                                                                                                _ = v179
                                                                                                            } else {
                                                                                                                v180 := GenS20{F0: 906, F1: -242, F2: -1}
                                                                                                                var v181 uint16 = 786
                                                                                                                v4 += 1
                                                                                                                if (true || (v4 != v4)) {
                                                                                                                    _ = len(v31)
                                                                                                                    fmt.Println(v5, v166, v178)
                                                                                                                } else {
                                                                                                                    v182 := GenS21{F0: 0}
                                                                                                                    var v183 uint16 = 922
                                                                                                                    var v184 float64 = ((1.5 - 1.5) / ((0.1 * (3.14159 / (-2.25 - 1.5))) * float64(1e10)))
                                                                                                                    var v185 uint8 = 255
                                                                                                                    _ = v182
                                                                                                                    _ = v183
                                                                                                                    _ = v184
                                                                                                                    _ = v185
                                                                                                                }
                                                                                                                v3 += (^(v3))
                                                                                                                switch "" {
                                                                                                                case "a":
                                                                                                                    switch -179 {
                                                                                                                    case 117:
                                                                                                                        if true {
                                                                                                                            delete(v1, "")
                                                                                                                            v31 = append(v31, fi152)
                                                                                                                            v186 := GenS1{F0: (!false), F1: (v178 % (((v178 * (v178 / ((v178) | 1)))) | 1))}
                                                                                                                            _ = v186
                                                                                                                        } else {
                                                                                                                            _ = len(v31)
                                                                                                                            _ = len(v31)
                                                                                                                            v187 := GenS14{F0: fi151, F1: 2, F2: (false && true)}
                                                                                                                            var v188 float32 = 3.14159
                                                                                                                            _ = v187
                                                                                                                            _ = v188
                                                                                                                        }
                                                                                                                    case 4294967295:
                                                                                                                        v178 -= 9223372036854775807
                                                                                                                    case 603:
                                                                                                                        var v189 uint32 = 1
                                                                                                                        _ = v189
                                                                                                                    default:
                                                                                                                        v31 = append(v31, (fi151 << ((v3) & 63)))
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    v190 := GenS22{F0: ("\n\t" + "\n\t"), F1: (v3 & 2), F2: 42}
                                                                                                                    _ = v190
                                                                                                                default:
                                                                                                                    v191 := GenS16{F0: v181}
                                                                                                                    _ = v191
                                                                                                                }
                                                                                                                _ = v180
                                                                                                                _ = v181
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch "hello" {
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(fi152)
                                                                                                            case "a":
                                                                                                                v166 = v166
                                                                                                            case "hello":
                                                                                                                var v192 float64 = (3.14159 + -2.25)
                                                                                                                _ = v192
                                                                                                            default:
                                                                                                                var v193 bool = (((1.5 / 3.14159) * 1.5) == (float32(0.0) * 3.14159))
                                                                                                                _ = v193
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v178
                                                                                                    }
                                                                                                    var v194 uint64 = 473
                                                                                                    var v195 uint64 = (v194 - (^(v194)))
                                                                                                    switch fi152 {
                                                                                                    case 128:
                                                                                                        fmt.Println(v5)
                                                                                                    default:
                                                                                                        v196 := GenS23{F0: false, F1: 850, F2: false}
                                                                                                        _ = v196
                                                                                                    }
                                                                                                    var v197 int8 = 100
                                                                                                    fmt.Println(v3, v5)
                                                                                                    _ = v194
                                                                                                    _ = v195
                                                                                                    _ = v197
                                                                                                }
                                                                                                v31 = append(v31, fi152)
                                                                                            } else {
                                                                                                fmt.Println(v4)
                                                                                                var v198 int64 = 42
                                                                                                _ = v198
                                                                                            }
                                                                                            _ = v167
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(fi152, v4, v5)
                                                                                    if true {
                                                                                        if true {
                                                                                            v166 = 706
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                v3 += 18446744073709551615
                                                                                            case "hello":
                                                                                                var v199 uint16 = v5
                                                                                                _ = v199
                                                                                            case "foo bar":
                                                                                                v166 -= 208
                                                                                            default:
                                                                                                var v200 uint8 = 2
                                                                                                _ = v200
                                                                                            }
                                                                                            var v201 int64 = -9223372036854775808
                                                                                            _ = v201
                                                                                        } else {
                                                                                            if (uint32(v5) > 0) {
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    var v202 int16 = (v4 % ((v4) | 1))
                                                                                                    _ = v202
                                                                                                }
                                                                                                if true {
                                                                                                    var v203 float64 = (0.0 + float64(float32((0.0 + 1.5))))
                                                                                                    fmt.Println(v5, v203, v3)
                                                                                                    _ = v203
                                                                                                }
                                                                                            } else {
                                                                                                v204 := GenS24{F0: false}
                                                                                                v205 := GenS6{F0: -128, F1: "foo bar", F2: 2147483647}
                                                                                                v206 := GenS2{F0: "a", F1: 100}
                                                                                                _ = v204
                                                                                                _ = v205
                                                                                                _ = v206
                                                                                            }
                                                                                            v5 = v5
                                                                                        }
                                                                                        if false {
                                                                                            if len(v31) >= 2 {
                                                                                                v31 = v31[1:len(v31)]
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v4, fi152)
                                                                                    }
                                                                                }
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    var v207 uint = v3
                                                                                    if true {
                                                                                        v4 = (v4 / ((v4) | 1))
                                                                                        if (((1.5 - ((1e10 - 3.14159) * (0.1 / float32(((0.1 / (((-2.25 * 0.0) - ((((float64(0.0) - 0.0) - (((((3.14159 + 0.0) * 3.14159) - 0.1) + 0.0) + (((3.14159 * ((((3.14159 - 1.5) / 0.1) * 0.0) * 1.5)) / (3.14159 / 3.14159)) - ((0.1 / (1.5 + 3.14159)) + 1e10)))) / 1.0) + 0.0)) + 0.0)) - 0.1))))) + (0.1 * 0.0)) <= 0.1) {
                                                                                            _ = len(v31)
                                                                                        } else {
                                                                                            v3 -= 65535
                                                                                            v3 = 852
                                                                                            var v208 uint = (v207 / ((v207) | 1))
                                                                                            v207 = uint(v5)
                                                                                            _ = v208
                                                                                        }
                                                                                    }
                                                                                    _ = v207
                                                                                case "\n\t":
                                                                                    fmt.Println(v4, v166)
                                                                                case "":
                                                                                    fmt.Println(fi151, v5, v166)
                                                                                default:
                                                                                    var v209 bool = true
                                                                                    _ = v209
                                                                                }
                                                                                _ = v166
                                                                            }
                                                                        } else {
                                                                            if (fi151 >= (fi152 | (fi151 ^ (fi152 - (-(fi151)))))) {
                                                                                fmt.Println(fi152, v4)
                                                                                if false {
                                                                                    if ("" > "\n\t") {
                                                                                        _ = len(v31)
                                                                                    } else {
                                                                                        if (v5 >= v5) {
                                                                                            var v210 uint8 = 255
                                                                                            v3 = (v3 * (v3 >> ((v3) & 63)))
                                                                                            fmt.Println(fi152, v4)
                                                                                            _ = v210
                                                                                        } else {
                                                                                            v211 := GenS25{F0: 185}
                                                                                            v212 := GenS19{F0: 196, F1: uint32(v3), F2: (fi152 * (fi151 / ((fi152) | 1)))}
                                                                                            _ = v211
                                                                                            _ = v212
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v5)
                                                                                    _ = len(v2)
                                                                                    v213 := GenS14{F0: fi152, F1: 42, F2: true}
                                                                                    _ = v213
                                                                                }
                                                                            }
                                                                            _ = len(v31)
                                                                            v3 = v3
                                                                        }
                                                                        v4 += (v4 >> ((0) & 63))
                                                                        _ = v164
                                                                    }
                                                                case 272:
                                                                    v214 := GenS26{F0: "hello", F1: 0}
                                                                    _ = v214
                                                                default:
                                                                    v215 := GenS27{F0: 83, F1: 127, F2: v4}
                                                                    _ = v215
                                                                }
                                                                v216 := GenS28{F0: 448}
                                                                _ = v216
                                                            default:
                                                                v217 := GenS29{F0: -24, F1: 378}
                                                                _ = v217
                                                            }
                                                        }
                                                        if true {
                                                            if (11 >= -25) {
                                                                v218 := GenS13{F0: true, F1: v5}
                                                                if true {
                                                                    if (128 < (-(fi152))) {
                                                                        v5 -= (v5 | (^(v5)))
                                                                    } else {
                                                                        v1[("" + "foo bar")] = "a"
                                                                    }
                                                                } else {
                                                                    switch ("" + "") {
                                                                    case "":
                                                                        var v219 int64 = -1
                                                                        _ = v219
                                                                    }
                                                                    v220 := GenS25{F0: v4}
                                                                    _ = len(v31)
                                                                    _ = v220
                                                                }
                                                                var v221 int = fi151
                                                                _ = v218
                                                                _ = v221
                                                            }
                                                            fmt.Println(v4)
                                                            var v222 uint = v3
                                                            _ = v222
                                                        } else {
                                                            v223 := GenS30{F0: 42, F1: v4}
                                                            var v224 int = fi152
                                                            _ = v223
                                                            _ = v224
                                                        }
                                                        v5 = (v5 >> ((v3) & 63))
                                                    case 32767:
                                                        var v225 uint = v3
                                                        _ = v225
                                                    default:
                                                        var v226 string = "hello"
                                                        _ = v226
                                                    }
                                                }
                                                if len(v31) > 0 {
                                                    v31[0] = fi152
                                                }
                                            } else {
                                                fmt.Println(fi152)
                                                var v227 uint64 = uint64(v3)
                                                _ = v227
                                            }
                                        case 994:
                                            if len(v31) >= 2 {
                                                v31 = v31[1:len(v31)]
                                            }
                                        default:
                                            fmt.Println(v3, fi151, v4)
                                        }
                                        if true {
                                            if false {
                                                var v228 uint = (v3 - 65535)
                                                var v229 uint32 = 4294967295
                                                var v230 uint32 = (v229 & 2)
                                                _ = v228
                                                _ = v229
                                                _ = v230
                                            } else {
                                                v231 := GenS31{F0: false, F1: 127}
                                                v5 -= (v5 % ((868) | 1))
                                                fmt.Println(fi152, fi151)
                                                fmt.Println(fi152)
                                                switch "a" {
                                                case "":
                                                    fmt.Println(v3, fi151, v5)
                                                    if false {
                                                        v232 := GenS32{F0: (^(fi151)), F1: -18}
                                                        fmt.Println(fi152)
                                                        var v233 int64 = -1
                                                        _ = v232
                                                        _ = v233
                                                    }
                                                    var v234 string = "foo bar"
                                                    _ = v234
                                                case "hello":
                                                    var v235 uint16 = 255
                                                    _ = v235
                                                case "foo bar":
                                                    if len(v31) > 0 {
                                                        v31[0] = (fi151 % (((fi152 * fi152)) | 1))
                                                    }
                                                }
                                                _ = v231
                                            }
                                            var v236 string = "\n\t"
                                            _ = v236
                                        }
                                    } else {
                                        v4 += 1000
                                        v5 = v5
                                    }
                                    var v237 int64 = int64(fi151)
                                    if (v3 == v3) {
                                        if (fi152 >= fi151) {
                                            v3 += v3
                                        } else {
                                            v1[""] = "hello"
                                            var v238 int = (fi152 >> (((^(v3))) & 63))
                                            _ = len(v1)
                                            switch 42 {
                                            case 4294967295:
                                                v238 = fi151
                                            case 127:
                                                fmt.Println(v237, v5, v4)
                                            }
                                            v237 = (v237 | (v237 ^ 42))
                                            v3 = 504
                                            v239 := GenS19{F0: 127, F1: 2147483648, F2: (-(fi151))}
                                            if false {
                                                v237 += (v237 & v237)
                                                v240 := GenS15{F0: 42, F1: uint8(v5)}
                                                _ = v240
                                            } else {
                                                if (v237 <= (^(v237))) {
                                                    v241 := GenS33{F0: 42}
                                                    _ = v241
                                                } else {
                                                    switch ("" + "a") {
                                                    case "hello":
                                                        if (-39 == -1) {
                                                            switch -418 {
                                                            case 42:
                                                                fmt.Println(fi152)
                                                            default:
                                                                if true {
                                                                    var v242 int32 = -2147483648
                                                                    if (((float64(-2.25) < (float64(-2.25) + ((float64(3.14159) / 1.0) + (-2.25 + (1e10 / ((1.5 - 0.1) - float64(0.0))))))) && false) || false) {
                                                                        v4 = (v4 / ((v4) | 1))
                                                                        if true {
                                                                            fmt.Println(v237, v3)
                                                                            var v243 int32 = v242
                                                                            _ = v243
                                                                        } else {
                                                                            var v244 uint = (v3 >> ((65535) & 63))
                                                                            v244 += v3
                                                                            v245 := GenS34{F0: 42, F1: 1, F2: 42}
                                                                            _ = v244
                                                                            _ = v245
                                                                        }
                                                                        v242 = (^(v242))
                                                                        v1["foo bar"] = (("a" + "\n\t") + "")
                                                                        v3 -= (v3 / (((v3 >> (((v3 + v3)) & 63))) | 1))
                                                                    }
                                                                    delete(v2, 4294967295)
                                                                    _ = v242
                                                                } else {
                                                                    var v246 uint32 = 670
                                                                    _ = v246
                                                                }
                                                            }
                                                            v1[("a" + ("a" + "a"))] = "a"
                                                            delete(v2, 42)
                                                            var v247 int64 = v237
                                                            var v248 int8 = 122
                                                            _ = v247
                                                            _ = v248
                                                        } else {
                                                            v249 := v1[""]
                                                            _ = v249
                                                        }
                                                    case "a":
                                                        _ = len(v1)
                                                        _ = len(v31)
                                                    }
                                                    if false {
                                                        v4 = v4
                                                        v250 := GenS35{F0: 2147483647, F1: 127}
                                                        var v251 float64 = (0.1 + float64(3.14159))
                                                        v252 := GenS6{F0: 49, F1: "a", F2: 861}
                                                        _ = v250
                                                        _ = v251
                                                        _ = v252
                                                    } else {
                                                        var v253 float32 = 1.5
                                                        switch ((("hello" + "a") + "") + "\n\t") {
                                                        case "hello":
                                                            var v254 bool = (!true)
                                                            v5 += 578
                                                            var v255 bool = (182 <= 42)
                                                            _ = v254
                                                            _ = v255
                                                        case "":
                                                            fmt.Println(v5, fi151, v238)
                                                        case "a":
                                                            fmt.Println(fi152, v253, v253)
                                                        default:
                                                            if (!false) {
                                                                if (v4 == (v4 & (^(v4)))) {
                                                                    switch 0 {
                                                                    case 581:
                                                                        switch ("\n\t" + "") {
                                                                        case "\n\t":
                                                                            var v256 int32 = 2147483647
                                                                            _ = v256
                                                                        }
                                                                    case 42:
                                                                        var v257 uint32 = 65535
                                                                        fmt.Println(v5, v237)
                                                                        switch (fi152 * (fi152 ^ 32767)) {
                                                                        case -260:
                                                                            v258 := GenS35{F0: 189, F1: 255}
                                                                            switch ("hello" + "") {
                                                                            case "a":
                                                                                v259 := GenS23{F0: true, F1: (v237 ^ v237), F2: false}
                                                                                _ = v259
                                                                            }
                                                                            _ = v258
                                                                        case -164:
                                                                            var v260 float32 = ((3.14159 - v253) - v253)
                                                                            _ = v260
                                                                        case 9223372036854775807:
                                                                            if (v4 < v4) {
                                                                                v261 := GenS36{F0: true}
                                                                                _ = v261
                                                                            }
                                                                        }
                                                                        _ = v257
                                                                    case 2147483647:
                                                                        var v262 bool = (int64(v4) != 1)
                                                                        _ = v262
                                                                    default:
                                                                        v237 = (v237 % ((v237) | 1))
                                                                    }
                                                                    fmt.Println(v5, v5)
                                                                }
                                                            } else {
                                                                v263 := GenS37{F0: v238, F1: v3}
                                                                switch (("a" + "") + "\n\t") {
                                                                case "foo bar":
                                                                    if len(v31) >= 2 {
                                                                        v31 = v31[1:len(v31)]
                                                                    }
                                                                case "\n\t":
                                                                    var v264 uint32 = 452
                                                                    _ = v264
                                                                default:
                                                                    v253 = v253
                                                                }
                                                                var v265 float64 = float64((3.14159 * 3.14159))
                                                                var v266 float64 = 1e10
                                                                if true {
                                                                    fmt.Println(v253)
                                                                } else {
                                                                    if true {
                                                                        v4 += (v4 * v4)
                                                                    }
                                                                }
                                                                _ = v263
                                                                _ = v265
                                                                _ = v266
                                                            }
                                                        }
                                                        _ = v253
                                                    }
                                                    v3 -= v3
                                                    v4 = v4
                                                }
                                                if (((true || false) || (-1 >= -1)) || true) {
                                                    var v267 int16 = (-(v4))
                                                    _ = v267
                                                }
                                                v237 = v237
                                            }
                                            _ = v238
                                            _ = v239
                                        }
                                        var v268 uint64 = 181
                                        v237 -= 173
                                        _ = v268
                                    }
                                    _ = fi152
                                    _ = v237
                                }
                                _ = fi151
                            }
                        }
                    }
                }
                v4 = (v4 - v4)
                v31 = append(v31, 9223372036854775807)
            }
            fmt.Println(v3, v3)
            _ = v109
        }
    }
    v4 -= (v4 * (v4 % ((-1) | 1)))
    for fi269 := 0; fi269 < 13; fi269++ {
        var v270 float64 = 1e10
        _ = fi269
        _ = v270
    }
    v1["\n\t"] = "foo bar"
    v271 := GenS38{F0: "hello", F1: v4}
    switch -179 {
    case 643:
        var v272 int16 = -32768
        _ = v272
    default:
        for fi273 := 0; fi273 < 11; fi273++ {
            var v274 string = "foo bar"
            _ = fi273
            _ = v274
        }
    }
    var v275 int = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v275:", v275)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v271
    _ = v275
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
