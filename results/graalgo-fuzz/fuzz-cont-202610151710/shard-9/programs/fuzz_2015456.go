package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int8; F2 int8 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 uint32; F2 int64 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64; F1 int32 }
type GenS13 struct { F0 int32; F1 int32 }
type GenS14 struct { F0 uint8; F1 string }

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    var v3 uint32 = 458
    var v4 uint = 18446744073709551615
    var v5 uint = v4
    v6 := make(map[uint32]string)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 uint = (v4 * (v4 & (v4 | (^(v5)))))
    v8 := v6[v3]
    if false {
        fmt.Println(v4, v8, v3)
    } else {
        var v9 int64 = -344
        for fi10 := 0; fi10 < 14; fi10++ {
            v8 = v8
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
        }
        _ = v9
    }
    if ((1e10 * 1.5) <= 3.14159) {
        if (true || true) {
            v4 += 4294967295
            if ((-63 == 896) && (1 < 2)) {
                var v11 string = (("foo bar" + v8) + v8)
                fmt.Println(v8, v7)
                v3 += v3
                _ = v11
            } else {
                v2[v8] = v3
            }
            var v12 uint32 = v3
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 bool = false
            _ = fi13
            _ = v14
        }
    }
    v4 = (v7 * (^(v4)))
    fmt.Println(v7, v8)
    if true {
        if false {
            v6[2147483648] = (v8 + "foo bar")
        }
        fmt.Println(v4, v8)
        v8 = v8
    }
    switch -9223372036854775808 {
    case 128:
        var v15 int32 = 1
        _ = v15
    case -284:
        v16 := GenS1{F0: -897, F1: -1}
        _ = v16
    }
    if ((1.5 - float64(((((0.1 + -2.25) - 3.14159) * 0.1) * 1e10))) != 1e10) {
        switch (("\n\t" + v8) + "foo bar") {
        case "foo bar":
            var v17 bool = (true || (-490 > -9223372036854775808))
            _ = v17
        case "":
            v7 = (^(v4))
        default:
            if (-2.25 != 1.5) {
                fmt.Println(v8, v3)
                var v18 int = 127
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v18
            }
        }
    }
    if true {
        if ((0.1 + 3.14159) <= 0.0) {
            v3 = (v3 % (((v3 >> (((v5 % ((v5) | 1))) & 63))) | 1))
        } else {
            if (true && true) {
                fmt.Println(v4, v3)
                delete(v2, v8)
                var v19 bool = (!false)
                _ = v19
            }
            if false {
                v20 := GenS1{F0: 32768, F1: -53}
                _ = v20
            }
        }
        if false {
            v5 += (v5 / ((65535) | 1))
            var v21 string = "hello"
            if true {
                v6[v3] = "foo bar"
                _ = len(v2)
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := GenS2{F0: 4294967295, F1: 1}
                    for fi24 := 0; fi24 < 17; fi24++ {
                        delete(v2, "hello")
                        _ = fi24
                    }
                    _ = fi22
                    _ = v23
                }
            } else {
                var v25 uint64 = 18446744073709551615
                if (0.0 >= 1.5) {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v2["foo bar"] = (v3 - 1)
                        _ = fi26
                    }
                    fmt.Println(v21, v4, v8)
                    var v27 uint = 999
                    _ = v27
                } else {
                    if (65535 <= 255) {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            v4 -= v5
                            v29 := GenS1{F0: (fi28 | (fi28 % ((fi28) | 1))), F1: 127}
                            for fi30 := 0; fi30 < 13; fi30++ {
                                switch -451 {
                                case 42:
                                    delete(v6, v3)
                                case 4294967295:
                                    var v31 int64 = 42
                                    _ = v31
                                case -9223372036854775808:
                                    var v32 uint8 = 203
                                    _ = v32
                                }
                                delete(v2, v21)
                                _ = fi30
                            }
                            _ = fi28
                            _ = v29
                        }
                        var v33 uint64 = uint64(v4)
                        for fi34 := 0; fi34 < 16; fi34++ {
                            if true {
                                for fi35 := 0; fi35 < 12; fi35++ {
                                    v5 = uint(v25)
                                    var v36 uint8 = 127
                                    var v37 int8 = 100
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                            }
                            for fi38 := 0; fi38 < 6; fi38++ {
                                v5 -= (v4 * v4)
                                fmt.Println(v5)
                                v39 := GenS3{F0: 42}
                                v8 = v8
                                _ = fi38
                                _ = v39
                            }
                            _ = fi34
                        }
                        _ = v33
                    } else {
                        if (!(true && true)) {
                            v40 := GenS4{F0: 2, F1: 100, F2: 1}
                            var v41 bool = (!((-2.25 + (3.14159 + ((-2.25 * ((1.5 * 0.1) / 1e10)) * ((-2.25 * -2.25) / (1e10 / 1.0))))) == -2.25))
                            for fi42 := 0; fi42 < 1; fi42++ {
                                v43 := GenS2{F0: v25, F1: 9223372036854775807}
                                _ = fi42
                                _ = v43
                            }
                            _ = v40
                            _ = v41
                        } else {
                            if (v5 == 234) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v8, v3, v21)
                            }
                            v44 := GenS1{F0: 128, F1: -7}
                            _ = v44
                        }
                        var v45 bool = false
                        _ = v45
                    }
                    var v46 int64 = -1
                    if (!true) {
                        v21 = (v21 + v21)
                    }
                    _ = v46
                }
                if false {
                    var v47 bool = false
                    v21 = "foo bar"
                    var v48 uint64 = 42
                    _ = v47
                    _ = v48
                }
                _ = v25
            }
            _ = v21
        } else {
            fmt.Println(v8, v4, v8)
            if true {
                for fi49 := 0; fi49 < 13; fi49++ {
                    v3 = v3
                    switch ("hello" + (v8 + (((v8 + "hello") + "\n\t") + "foo bar"))) {
                    case "":
                        v50 := GenS5{F0: 1, F1: (v7 >> ((v7) & 63))}
                        _ = v50
                    case "hello":
                        var v51 float32 = 3.14159
                        _ = v51
                    case "\n\t":
                        var v52 int = fi49
                        _ = v52
                    }
                    v5 = 2
                    _ = fi49
                }
                if false {
                    if (42 >= 1) {
                        v7 = (^(v4))
                        v53 := v6[42]
                        _ = v53
                    }
                    v8 = v8
                    v54 := GenS3{F0: 0}
                    _ = v54
                }
                switch v8 {
                case "\n\t":
                    if (1.5 < 1.5) {
                        v8 = v8
                        switch "hello" {
                        case "\n\t":
                            _ = len(v1)
                        case "hello":
                            v8 = v8
                        case "a":
                            var v55 int32 = -1
                            _ = v55
                        default:
                            if (1 == 4294967295) {
                                for fi56 := 0; fi56 < 7; fi56++ {
                                    v57 := GenS3{F0: 255}
                                    _ = fi56
                                    _ = v57
                                }
                                v58 := GenS2{F0: 4294967295, F1: -1}
                                v59 := GenS2{F0: 0, F1: -1}
                                _ = v58
                                _ = v59
                            } else {
                                switch -2147483648 {
                                case -2147483648:
                                    v60 := GenS6{F0: v4, F1: 873}
                                    _ = v60
                                }
                                v61 := v6[uint32(v7)]
                                _ = v61
                            }
                        }
                    } else {
                        v8 = (v8 + v8)
                    }
                default:
                    for fi62 := 0; fi62 < 1; fi62++ {
                        var v63 uint8 = 255
                        var v64 int8 = 20
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                }
            } else {
                var v65 uint8 = 75
                _ = v65
            }
            v6[(v3 + (v3 << ((18446744073709551615) & 63)))] = v8
        }
        if (-9223372036854775808 == -1) {
            for fi66 := 0; fi66 < 7; fi66++ {
                v5 = (v7 ^ (v4 + v5))
                for fi67 := 0; fi67 < 14; fi67++ {
                    var v68 uint64 = 949
                    if ((v7 ^ v4) == v7) {
                        fmt.Println(v7, v8, v5)
                    }
                    _ = fi67
                    _ = v68
                }
                v4 = (v7 >> (((^(v4))) & 63))
                _ = fi66
            }
        }
    } else {
        if true {
            for fi69 := 0; fi69 < 12; fi69++ {
                v70 := v2[("foo bar" + v8)]
                _ = fi69
                _ = v70
            }
            var v71 int32 = 0
            _ = v71
        } else {
            var v72 int32 = 2147483647
            if false {
                v73 := GenS6{F0: (^(v4)), F1: 660}
                fmt.Println(v3, v7)
                _ = v73
            }
            for fi74 := 0; fi74 < 1; fi74++ {
                v75 := GenS4{F0: 161, F1: 0, F2: 42}
                fmt.Println(v4)
                _ = fi74
                _ = v75
            }
            _ = v72
        }
        var v76 uint16 = 0
        v77 := GenS6{F0: v7, F1: v76}
        _ = v76
        _ = v77
    }
    switch -828 {
    case -578:
        v5 = 18446744073709551615
    case 84:
        for fi78 := 0; fi78 < 11; fi78++ {
            v79 := GenS7{F0: 42}
            for fi80 := 0; fi80 < 5; fi80++ {
                if ((v5 - 42) <= (v5 | 1)) {
                    if (!true) {
                        fmt.Println(fi80, v4)
                    }
                }
                if true {
                    v4 -= v4
                } else {
                    if false {
                        if (-90 != 16) {
                            if (true && (2147483647 > 2147483647)) {
                                if ((0.1 > (0.0 + (1e10 + 1.5))) || (42 == -108)) {
                                    if false {
                                        v81 := GenS8{F0: ("a" + v8), F1: v3, F2: 9223372036854775807}
                                        v3 = v3
                                        fmt.Println(fi78, v4)
                                        _ = v81
                                    }
                                } else {
                                    var v82 float64 = ((3.14159 - 1.5) * float64(0.1))
                                    _ = v82
                                }
                                fmt.Println(v3, v3, v8)
                                v7 = 348
                            }
                        }
                        var v83 uint64 = 2
                        var v84 bool = true
                        _ = v83
                        _ = v84
                    }
                    v4 -= v5
                    if (42 != 0) {
                        v5 += v5
                    } else {
                        switch (fi78 * fi78) {
                        case 9223372036854775807:
                            if (!false) {
                                if (-1 > int64(v3)) {
                                    _ = len(v2)
                                    v3 = (v3 | v3)
                                    v7 += v4
                                    v85 := GenS8{F0: "foo bar", F1: (^(v3)), F2: 42}
                                    _ = v85
                                } else {
                                    var v86 uint64 = 42
                                    _ = v86
                                }
                                switch 415 {
                                case 9223372036854775807:
                                    v87 := GenS9{F0: 127, F1: v3}
                                    fmt.Println(v8)
                                    switch v8 {
                                    case "hello":
                                        if (-76 == 127) {
                                            v88 := v2[v8]
                                            v89 := GenS3{F0: 32767}
                                            fmt.Println(v3, v5, fi80)
                                            _ = v88
                                            _ = v89
                                        }
                                    case "foo bar":
                                        var v90 uint8 = 0
                                        _ = v90
                                    default:
                                        v91 := v6[1]
                                        _ = v91
                                    }
                                    var v92 int8 = 100
                                    _ = v87
                                    _ = v92
                                }
                                var v93 uint64 = 184
                                _ = v93
                            }
                        case -9:
                            if false {
                                var v94 int64 = 0
                                v95 := v2[v8]
                                _ = v94
                                _ = v95
                            } else {
                                v4 += v5
                            }
                            if (!(true || false)) {
                                v3 -= v3
                                v1 = append(v1, fi78)
                                fmt.Println(v3, v5, v7)
                            }
                        default:
                            v6[v3] = ("foo bar" + ("" + ""))
                        }
                        var v96 float32 = 1.5
                        _ = v96
                    }
                }
                v97 := GenS1{F0: fi80, F1: 1}
                fmt.Println(fi78, v3, fi80)
                var v98 float64 = -2.25
                _ = fi80
                _ = v97
                _ = v98
            }
            _ = fi78
            _ = v79
        }
    case -2147483648:
        v5 = 4294967295
    }
    v8 = "hello"
    if (32768 <= -551) {
        v3 -= (^(v3))
        v8 = (("" + v8) + "")
    }
    var v99 int16 = 492
    if true {
        switch 9223372036854775807 {
        case -1:
            for fi100 := 0; fi100 < 11; fi100++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi101 := 0; fi101 < 18; fi101++ {
                    fmt.Println(v5)
                    if len(v1) > 0 {
                        v1[0] = fi100
                    }
                    _ = fi101
                }
                _ = fi100
            }
        default:
            v102 := GenS10{F0: -128}
            _ = v102
        }
        switch 9223372036854775807 {
        case -147:
            v103 := GenS11{F0: (3.14159 > 0.1)}
            _ = v103
        case 778:
            if (4294967295 <= 128) {
                for fi104 := 0; fi104 < 3; fi104++ {
                    for fi105 := 0; fi105 < 16; fi105++ {
                        v106 := GenS10{F0: 42}
                        fmt.Println(v7, v4, v5)
                        var v107 float32 = 3.14159
                        _ = fi105
                        _ = v106
                        _ = v107
                    }
                    v2[("a" + "a")] = v3
                    _ = fi104
                }
                var v108 uint = (v7 % (((v4 % ((v7) | 1))) | 1))
                _ = v108
            } else {
                var v109 int = 32767
                for fi110 := 0; fi110 < 13; fi110++ {
                    v8 = v8
                    v111 := GenS7{F0: -1}
                    _ = fi110
                    _ = v111
                }
                _ = v109
            }
        case 32767:
            _ = len(v1)
        }
        for fi112 := 0; fi112 < 15; fi112++ {
            for fi113 := 0; fi113 < 13; fi113++ {
                switch "hello" {
                case "foo bar":
                    var v114 uint8 = 71
                    _ = v114
                case "a":
                    v3 -= 2147483648
                default:
                    switch fi113 {
                    case 127:
                        if (((!true) || (2147483647 < -2147483648)) && ((!((fi112 - fi112) > fi113)) && (!false))) {
                            v115 := GenS12{F0: 4294967295, F1: 2147483647}
                            _ = len(v1)
                            v3 += 177
                            _ = v115
                        } else {
                            switch (v8 + (("a" + "\n\t") + v8)) {
                            case "":
                                fmt.Println(v99, fi112, v8)
                            case "\n\t":
                                switch (-(fi113)) {
                                case 1:
                                    var v116 int = fi112
                                    _ = v116
                                case 0:
                                    fmt.Println(fi112, v3)
                                case -9223372036854775808:
                                    var v117 uint64 = 1
                                    _ = v117
                                default:
                                    v118 := GenS13{F0: -2147483648, F1: 0}
                                    _ = v118
                                }
                            case "a":
                                switch (fi113 | 1) {
                                case 1:
                                    v119 := GenS14{F0: 255, F1: "hello"}
                                    _ = v119
                                }
                            }
                            v7 -= v5
                            if true {
                                fmt.Println(v3)
                            }
                        }
                    default:
                        v8 = (("a" + v8) + ((v8 + (v8 + v8)) + ("" + "foo bar")))
                    }
                }
                _ = fi113
            }
            _ = fi112
        }
    }
    if (42 > 323) {
        for fi120 := 0; fi120 < 2; fi120++ {
            for fi121 := 0; fi121 < 18; fi121++ {
                var v122 uint = (v4 | 255)
                _ = fi121
                _ = v122
            }
            v123 := GenS2{F0: 4294967295, F1: 9223372036854775807}
            if len(v1) > 0 {
                v1[0] = 688
            }
            _ = fi120
            _ = v123
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v99
}
