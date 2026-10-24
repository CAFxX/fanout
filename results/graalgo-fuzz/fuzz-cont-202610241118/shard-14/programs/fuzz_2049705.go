package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint64; F2 string }
type GenS6 struct { F0 int32; F1 string; F2 uint }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 uint32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int8; F1 uint32; F2 int32 }

func main() {
    v1 := []uint32{440, 2147483647}
    var v2 uint64 = 468
    var v3 bool = (v2 < 985)
    v4 := make(map[int]string)
    if (0 != -215) {
        for fi5 := 0; fi5 < 18; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                v3 = (1 == 127)
                if ((fi5 >> ((97) & 63)) > fi6) {
                    fmt.Println(v3, v2)
                }
                _ = fi6
            }
            v3 = (-185 != 0)
            _ = fi5
        }
    } else {
        var v7 int = -2147483648
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v2 = 393
        v2 = 18446744073709551615
        _ = fi8
    }
    v9 := GenS1{F0: 1000}
    if (v3 && (true || (42 < 2147483647))) {
        var v10 uint8 = 127
        if v3 {
            switch "a" {
            case "":
                v3 = v3
            case "hello":
                var v11 uint = 65535
                _ = v11
            case "foo bar":
                var v12 uint32 = 1
                _ = v12
            default:
                if (("a" >= "foo bar") || v3) {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        v13 := GenS1{F0: -124}
                        _ = v13
                    case 42:
                        for fi14 := 0; fi14 < 19; fi14++ {
                            switch "a" {
                            case "foo bar":
                                switch "a" {
                                case "hello":
                                    v1 = append(v1, 42)
                                case "foo bar":
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                case "":
                                    switch "" {
                                    case "foo bar":
                                        switch 127 {
                                        case 9223372036854775807:
                                            for fi15 := 0; fi15 < 5; fi15++ {
                                                if v3 {
                                                    v16 := GenS1{F0: 32767}
                                                    switch 128 {
                                                    case 127:
                                                        if (((0.1 + -2.25) + 0.1) != (0.0 / 1.0)) {
                                                            v17 := GenS1{F0: 1}
                                                            _ = v17
                                                        } else {
                                                            var v18 uint8 = (v10 % ((v10) | 1))
                                                            var v19 float32 = ((-2.25 + 1e10) - (-2.25 - 3.14159))
                                                            _ = v18
                                                            _ = v19
                                                        }
                                                    }
                                                    _ = v16
                                                }
                                                var v20 uint32 = 32
                                                if (v3 && (0.0 >= (((0.0 - float32((1e10 + (3.14159 + 3.14159)))) * 1.5) + 3.14159))) {
                                                    switch (fi14 - (-(fi14))) {
                                                    case 32767:
                                                        var v21 uint = 708
                                                        _ = v21
                                                    }
                                                    _ = len(v1)
                                                    v20 = v20
                                                    var v22 int8 = -49
                                                    switch 9223372036854775807 {
                                                    case 42:
                                                        v23 := GenS1{F0: 32767}
                                                        _ = v23
                                                    }
                                                    _ = v22
                                                }
                                                _ = fi15
                                                _ = v20
                                            }
                                        }
                                    default:
                                        v24 := GenS1{F0: -387}
                                        _ = v24
                                    }
                                default:
                                    v25 := GenS1{F0: 42}
                                    _ = v25
                                }
                            case "\n\t":
                                v3 = v3
                            default:
                                v3 = (((v3 || (true && (true || (("hello" + ("" + "")) == ("a" + "hello"))))) || (v3 && true)) && v3)
                            }
                            for fi26 := 0; fi26 < 19; fi26++ {
                                v27 := GenS1{F0: 844}
                                switch ("\n\t" + ("foo bar" + ((("" + "hello") + (("" + ("" + "foo bar")) + "hello")) + "\n\t"))) {
                                case "\n\t":
                                    v3 = (!(!(103 >= 127)))
                                    switch "\n\t" {
                                    case "hello":
                                        if true {
                                            v28 := GenS1{F0: 32767}
                                            if v3 {
                                                var v29 bool = ((((true && v3) || (587 > -585)) || v3) && v3)
                                                _ = len(v4)
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                                v30 := GenS1{F0: 0}
                                                v10 = v10
                                                _ = v29
                                                _ = v30
                                            }
                                            _ = v28
                                        } else {
                                            var v31 bool = true
                                            if (!((v31 || false) && (((!(!v3)) || (469 >= 839)) && ((0 <= 42) && v3)))) {
                                                if (127 <= fi26) {
                                                    v31 = v3
                                                    switch 42 {
                                                    case 32768:
                                                        fmt.Println(v10, v31, v2)
                                                    case 514:
                                                        _ = len(v4)
                                                    case -2147483648:
                                                        var v32 uint32 = 1
                                                        _ = v32
                                                    }
                                                } else {
                                                    v33 := GenS2{F0: (v2 & v2), F1: v2}
                                                    var v34 uint8 = v10
                                                    fmt.Println(fi14, v34, v34)
                                                    _ = v33
                                                    _ = v34
                                                }
                                                var v35 string = "foo bar"
                                                _ = v35
                                            }
                                            var v36 bool = (408 > 42)
                                            _ = v31
                                            _ = v36
                                        }
                                    case "\n\t":
                                        fmt.Println(v3)
                                    case "a":
                                        fmt.Println(fi14, fi26)
                                    }
                                default:
                                    v37 := GenS3{F0: -32768, F1: 1}
                                    _ = v37
                                }
                                _ = fi26
                                _ = v27
                            }
                            for fi38 := 0; fi38 < 18; fi38++ {
                                var v39 float64 = (3.14159 * 3.14159)
                                _ = fi38
                                _ = v39
                            }
                            _ = fi14
                        }
                    case 127:
                        for fi40 := 0; fi40 < 19; fi40++ {
                            var v41 int32 = -2147483648
                            v42 := GenS1{F0: 874}
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    }
                    v2 += 42
                    if (0 < 37) {
                        if false {
                            if v3 {
                                for fi43 := 0; fi43 < 16; fi43++ {
                                    if (3.14159 < 1.5) {
                                        var v44 uint32 = 65535
                                        _ = v44
                                    } else {
                                        v2 -= v2
                                        v3 = false
                                        v45 := v4[1]
                                        _ = v45
                                    }
                                    _ = fi43
                                }
                                v46 := GenS4{F0: true, F1: v10}
                                if v3 {
                                    v47 := GenS3{F0: -217, F1: 1}
                                    v2 = 2
                                    for fi48 := 0; fi48 < 10; fi48++ {
                                        for fi49 := 0; fi49 < 12; fi49++ {
                                            fmt.Println(fi49, v3)
                                            if (579 <= 1) {
                                                var v50 int16 = 0
                                                _ = v50
                                            } else {
                                                if (0.0 == 0.1) {
                                                    delete(v4, (^(fi48)))
                                                    v51 := GenS5{F0: 42, F1: v2, F2: "a"}
                                                    _ = v51
                                                }
                                                v52 := GenS6{F0: -1, F1: "foo bar", F2: 726}
                                                if (858 != -961) {
                                                    v2 = v2
                                                    v3 = true
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    v53 := GenS7{F0: -2147483648, F1: 65535}
                                                    v10 += (v10 - (v10 | v10))
                                                    _ = v53
                                                }
                                                _ = v52
                                            }
                                            v2 = uint64(v10)
                                            _ = fi49
                                        }
                                        for fi54 := 0; fi54 < 5; fi54++ {
                                            if true {
                                                if (255 == 892) {
                                                    _ = len(v1)
                                                }
                                            } else {
                                                v55 := GenS2{F0: v2, F1: v2}
                                                var v56 bool = (int64(fi54) >= 1)
                                                var v57 uint8 = v10
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            }
                                            var v58 int8 = -59
                                            _ = fi54
                                            _ = v58
                                        }
                                        _ = fi48
                                    }
                                    _ = v47
                                }
                                _ = v46
                            }
                            v2 += v2
                        } else {
                            var v59 int32 = -1
                            if v3 {
                                for fi60 := 0; fi60 < 19; fi60++ {
                                    var v61 int64 = 1
                                    var v62 uint32 = 42
                                    v2 = (v2 - v2)
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                                v63 := v4[-1]
                                v10 -= (^(v10))
                                _ = v63
                            } else {
                                v64 := GenS6{F0: (v59 << ((255) & 63)), F1: "a", F2: 2}
                                _ = v64
                            }
                            v65 := GenS4{F0: false, F1: v10}
                            _ = v59
                            _ = v65
                        }
                        if v3 {
                            v2 = v2
                        }
                        v3 = (!v3)
                    } else {
                        v66 := GenS8{F0: v2, F1: 331, F2: (v2 | (v2 % (((v2 + v2)) | 1)))}
                        _ = v66
                    }
                }
            }
        }
        var v67 uint = 18446744073709551615
        _ = v10
        _ = v67
    }
    v3 = (v3 && true)
    if (!false) {
        v68 := GenS6{F0: 2147483647, F1: "", F2: 0}
        _ = v68
    } else {
        v4[42] = ""
        for fi69 := 0; fi69 < 5; fi69++ {
            for fi70 := 0; fi70 < 18; fi70++ {
                var v71 uint8 = uint8(v2)
                var v72 uint64 = v2
                _ = fi70
                _ = v71
                _ = v72
            }
            _ = fi69
        }
    }
    v73 := GenS6{F0: 42, F1: "hello", F2: 65535}
    for fi74 := 0; fi74 < 3; fi74++ {
        v2 += (v2 << ((65535) & 63))
        if v3 {
            switch fi74 {
            case 172:
                var v75 float64 = (float64(-2.25) * 0.1)
                _ = v75
            case 127:
                v3 = (v3 || true)
            }
            if false {
                for fi76 := 0; fi76 < 20; fi76++ {
                    v2 = 732
                    _ = fi76
                }
                var v77 bool = true
                v78 := GenS7{F0: 2147483647, F1: 949}
                _ = v77
                _ = v78
            }
            v1 = append(v1, 65535)
        } else {
            if v3 {
                var v79 int32 = 0
                v79 = (v79 * v79)
                for fi80 := 0; fi80 < 2; fi80++ {
                    if true {
                        v81 := GenS4{F0: v3, F1: 128}
                        v2 = 4294967295
                        v79 += 42
                        _ = v81
                    }
                    v82 := GenS9{F0: 2}
                    var v83 uint8 = 246
                    _ = fi80
                    _ = v82
                    _ = v83
                }
                _ = v79
            }
            var v84 float32 = 0.1
            if true {
                var v85 float32 = v84
                _ = v85
            } else {
                v3 = ((v2 & v2) != (v2 >> ((18446744073709551615) & 63)))
            }
            _ = v84
        }
        v3 = v3
        _ = fi74
    }
    if v3 {
        var v86 uint = 122
        _ = v86
    } else {
        v87 := GenS10{F0: 0}
        var v88 bool = (0 >= 255)
        _ = v87
        _ = v88
    }
    if v3 {
        var v89 int32 = 0
        var v90 uint64 = (^(v2))
        fmt.Println(v2)
        _ = v89
        _ = v90
    } else {
        switch -471 {
        case -530:
            v2 -= 18446744073709551615
        case -2147483648:
            v91 := GenS11{F0: v2, F1: 845}
            _ = v91
        default:
            for fi92 := 0; fi92 < 1; fi92++ {
                switch "\n\t" {
                case "":
                    for fi93 := 0; fi93 < 4; fi93++ {
                        v94 := GenS3{F0: 42, F1: -1}
                        var v95 int64 = 0
                        _ = fi93
                        _ = v94
                        _ = v95
                    }
                case "foo bar":
                    fmt.Println(fi92, v2)
                default:
                    var v96 uint32 = 65535
                    _ = v96
                }
                switch (fi92 ^ 43) {
                case 928:
                    var v97 uint = 578
                    _ = v97
                case 1:
                    _ = len(v1)
                }
                for fi98 := 0; fi98 < 10; fi98++ {
                    switch ("foo bar" + "") {
                    case "foo bar":
                        v4[(fi98 << ((292) & 63))] = "\n\t"
                    case "":
                        fmt.Println(v2, v3)
                    case "\n\t":
                        _ = len(v1)
                    }
                    v99 := GenS10{F0: 127}
                    v100 := GenS8{F0: (v2 / (((^(v2))) | 1)), F1: -9223372036854775808, F2: (^(v2))}
                    _ = fi98
                    _ = v100
                    _ = v99
                }
                _ = fi92
            }
        }
    }
    fmt.Println(v3, v3)
    v2 -= 4294967295
    var v101 int = 0
    fmt.Println(v2, v3, v3)
    if (v2 == 18446744073709551615) {
        v102 := GenS5{F0: 616, F1: v2, F2: ""}
        var v103 uint64 = v2
        if v3 {
            v103 = v103
        }
        _ = v102
        _ = v103
    }
    for fi104 := 0; fi104 < 4; fi104++ {
        if (32767 < 874) {
            for fi105 := 0; fi105 < 8; fi105++ {
                v106 := GenS10{F0: 160}
                v4[v101] = "\n\t"
                _ = fi105
                _ = v106
            }
            v107 := GenS12{F0: 0}
            for fi108 := 0; fi108 < 2; fi108++ {
                var v109 bool = v3
                _ = fi108
                _ = v109
            }
            _ = v107
        } else {
            for fi110 := 0; fi110 < 11; fi110++ {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                _ = fi110
            }
        }
        _ = fi104
    }
    if v3 {
        var v111 uint64 = 1
        if true {
            fmt.Println(v3, v101)
            var v112 uint8 = 128
            _ = v112
        }
        switch (("hello" + "a") + (("\n\t" + "") + "\n\t")) {
        case "\n\t":
            if ((-2.25 - (0.1 - ((-2.25 + float32(0.0)) + -2.25))) < (1e10 + 3.14159)) {
                for fi113 := 0; fi113 < 5; fi113++ {
                    for fi114 := 0; fi114 < 13; fi114++ {
                        v3 = false
                        if v3 {
                            var v115 int32 = 165
                            _ = v115
                        }
                        v116 := v4[v101]
                        _ = fi114
                        _ = v116
                    }
                    if (0 <= -662) {
                        v3 = v3
                        for fi117 := 0; fi117 < 14; fi117++ {
                            fmt.Println(v2)
                            if v3 {
                                _ = len(v4)
                                var v118 uint8 = 190
                                _ = v118
                            } else {
                                var v119 int64 = -561
                                if (("" < "hello") || (!(!false))) {
                                    v2 = v111
                                } else {
                                    if true {
                                        switch (fi117 ^ fi117) {
                                        case 42:
                                            var v120 float32 = 0.1
                                            _ = v120
                                        case 32768:
                                            v119 += (v119 / (((v119 / ((1) | 1))) | 1))
                                        }
                                        v3 = v3
                                        switch ("foo bar" + "foo bar") {
                                        case "":
                                            var v121 float32 = 3.14159
                                            v3 = true
                                            switch (fi117 - 42) {
                                            case -1:
                                                v121 = (float32(((1e10 - -2.25) * -2.25)) * (((v121 * v121) / (float32(-2.25) - 0.0)) / v121))
                                            case 361:
                                                fmt.Println(v119, v111, v121)
                                            }
                                            _ = v121
                                        }
                                    }
                                    v2 = (v111 >> ((477) & 63))
                                    v111 = v2
                                }
                                _ = v119
                            }
                            v122 := GenS12{F0: 1}
                            _ = fi117
                            _ = v122
                        }
                        if v3 {
                            var v123 uint64 = v111
                            _ = v123
                        } else {
                            v4[127] = "\n\t"
                        }
                    } else {
                        if (588 > 73) {
                            if (-1 > 210) {
                                v2 += 289
                            } else {
                                var v124 int = v101
                                v125 := GenS12{F0: 2}
                                _ = v124
                                _ = v125
                            }
                            v126 := GenS3{F0: 1, F1: 1}
                            _ = v126
                        }
                        v101 = fi113
                    }
                    _ = fi113
                }
                fmt.Println(v111)
                var v127 int8 = 0
                _ = v127
            }
        default:
            if ((!(165 == 19)) && false) {
                fmt.Println(v2, v2)
                if false {
                    switch "\n\t" {
                    case "a":
                        v128 := GenS2{F0: v2, F1: (v2 + 1)}
                        _ = v128
                    case "hello":
                        var v129 float32 = 0.0
                        _ = v129
                    case "foo bar":
                        var v130 uint = 0
                        _ = v130
                    }
                } else {
                    v3 = ((4294967295 > (v2 * (v111 & (v2 << ((65535) & 63))))) || true)
                    var v131 int32 = 42
                    v132 := GenS12{F0: 0}
                    _ = v131
                    _ = v132
                }
            } else {
                v133 := GenS13{F0: -1, F1: 245, F2: int32(v2)}
                v101 = (v101 | v101)
                fmt.Println(v111)
                _ = v133
            }
        }
        _ = v111
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v9
}
