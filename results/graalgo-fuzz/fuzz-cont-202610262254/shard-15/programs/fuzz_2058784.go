package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16; F1 int; F2 uint8 }
type GenS8 struct { F0 uint8; F1 int; F2 int64 }
type GenS9 struct { F0 int32; F1 uint8; F2 int16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint64; F1 int16; F2 int16 }
type GenS13 struct { F0 uint8; F1 uint64 }
type GenS14 struct { F0 int32; F1 int64; F2 uint64 }
func genFn0(p0 uint16, p1 int32, p2 int) string {
    v6 := GenS1{F0: (p2 - (p2 % (((p2 | p2)) | 1)))}
    switch (("" + "a") + (("" + "") + ("" + "hello"))) {
    case "hello":
        v7 := GenS1{F0: p2}
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            p0 -= (p0 + p0)
            _ = fi8
        }
        var v9 uint32 = 1
        for fi10 := 0; fi10 < 17; fi10++ {
            if false {
                v11 := []int{fi10}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 bool = ((true && false) || ((v9 / (((^(v9))) | 1)) == 0))
                    var v14 uint8 = 0
                    if v13 {
                        fmt.Println(v9, fi10)
                        var v15 bool = (v13 || (p0 != p0))
                        fmt.Println(v14, p1)
                        _ = v15
                    } else {
                        fmt.Println(p1, v14, p2)
                    }
                    var v16 uint32 = v9
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v16
                }
            }
            _ = fi10
        }
        _ = v9
    } else {
        p1 -= p1
        v17 := GenS1{F0: (-(p2))}
        v18 := GenS1{F0: p2}
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"foo bar"}
    var v3 int64 = -1
    var v4 uint = 1
    var v5 string = ""
    switch -531 {
    case 32767:
        switch 0 {
        case -1:
            v5 = "\n\t"
        case 127:
            fmt.Println(v3, v4, v3)
        }
    case 1:
        var v19 bool = (-2147483648 != 2147483647)
        _ = v19
    case -375:
        v20 := v1["foo bar"]
        _ = v20
    }
    var v21 int = 32767
    v5 = ""
    if len(v2) > 0 {
        v2[0] = ((v5 + "") + v5)
    }
    if (float32(1.5) <= 3.14159) {
        v22 := GenS1{F0: v21}
        for fi23 := 0; fi23 < 9; fi23++ {
            v3 = -1
            v21 = 2147483647
            _ = fi23
        }
        switch v5 {
        case "foo bar":
            fmt.Println(v3, v5, v5)
        }
        _ = v22
    } else {
        var v24 int = v21
        v25 := GenS1{F0: v24}
        v26 := GenS2{F0: 191, F1: int32(v4)}
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 16; fi28++ {
            if (861 > 0) {
                _ = len(v2)
            }
            var v29 uint = v4
            if (int(v4) > v21) {
                switch "hello" {
                case "foo bar":
                    v5 = (v5 + v5)
                case "hello":
                    var v30 float32 = (1e10 * 3.14159)
                    _ = v30
                case "\n\t":
                    fmt.Println(v29, v21)
                default:
                    var v31 float64 = 3.14159
                    _ = v31
                }
            } else {
                switch (v21 / (((fi28 / (((fi28 | (^(fi27)))) | 1))) | 1)) {
                case 2147483647:
                    if false {
                        switch (fi28 << ((v4) & 63)) {
                        case 315:
                            var v32 bool = false
                            _ = v32
                        case -2147483648:
                            var v33 int8 = -30
                            _ = len(v2)
                            _ = v33
                        case 4294967295:
                            var v34 bool = (false && true)
                            _ = v34
                        }
                        if len(v2) > 0 {
                            v2[0] = v5
                        }
                        v35 := GenS3{F0: true, F1: true, F2: 42}
                        fmt.Println(fi28, v21, v5)
                        if len(v2) > 0 {
                            v2[0] = v5
                        }
                        var v36 int32 = 0
                        var v37 float64 = 3.14159
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                default:
                    if false {
                        var v38 int32 = -525
                        var v39 int32 = v38
                        fmt.Println(v5, v21, v3)
                        v38 = v39
                        _ = v38
                        _ = v39
                    } else {
                        var v40 int32 = 740
                        fmt.Println(v21)
                        fmt.Println(v4, v40, fi28)
                        _ = v40
                    }
                }
                v5 = v5
                v41 := GenS2{F0: 42, F1: 762}
                v42 := GenS1{F0: 2147483647}
                var v43 int8 = -128
                _ = v41
                _ = v42
                _ = v43
            }
            _ = fi28
            _ = v29
        }
        var v44 uint16 = 746
        var v45 uint32 = 2147483647
        _ = fi27
        _ = v44
        _ = v45
    }
    v5 = (v5 + v5)
    if (!true) {
        var v46 int8 = 90
        _ = v46
    }
    if true {
        if false {
            v5 = (((v5 + "foo bar") + (v5 + (((v5 + "") + ((v5 + v5) + (("foo bar" + "hello") + ""))) + (v5 + v5)))) + (((("hello" + "\n\t") + "foo bar") + v5) + v5))
            v4 = 18446744073709551615
        } else {
            for fi47 := 0; fi47 < 15; fi47++ {
                if true {
                    v48 := GenS3{F0: true, F1: false, F2: 0}
                    _ = v48
                } else {
                    _ = len(v1)
                }
                _ = fi47
            }
        }
        fmt.Println(v4, v4, v3)
        var v49 uint16 = 950
        _ = v49
    } else {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        if true {
            fmt.Println(v5, v4, v3)
        } else {
            var v50 float64 = (1e10 * 1.5)
            _ = v50
        }
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v3)
    for fi51 := 0; fi51 < 14; fi51++ {
        v1["hello"] = true
        for fi52 := 0; fi52 < 15; fi52++ {
            if (true || (false || (1.5 < 3.14159))) {
                if ((false || false) && false) {
                    fmt.Println(v21)
                } else {
                    var v53 string = (("a" + (("" + v5) + (v5 + v5))) + ("foo bar" + v5))
                    switch v53 {
                    case "\n\t":
                        v4 -= v4
                    }
                    var v54 uint8 = 1
                    _ = v53
                    _ = v54
                }
                v55 := GenS4{F0: 1, F1: -1, F2: 566}
                _ = v55
            } else {
                v56 := GenS3{F0: false, F1: (!false), F2: 1000}
                if true {
                    var v57 float64 = (0.1 + 3.14159)
                    if len(v2) > 0 {
                        v2[0] = v5
                    }
                    _ = v57
                } else {
                    switch fi52 {
                    case 1:
                        var v58 uint32 = 2
                        _ = v58
                    case 128:
                        var v59 uint8 = 223
                        _ = v59
                    }
                    v21 = (fi52 - (-(fi51)))
                    v3 += v3
                }
                v60 := GenS3{F0: true, F1: false, F2: 0}
                _ = v56
                _ = v60
            }
            if (false && (fi51 != (-(fi52)))) {
                v21 = -30
            }
            v3 -= v3
            _ = fi52
        }
        _ = fi51
    }
    var v61 float64 = 0.1
    switch v5 {
    case "\n\t":
        var v62 string = "a"
        _ = v62
    case "foo bar":
        if ((!true) || (2 < 1)) {
            v5 = "hello"
            if (-2.25 > v61) {
                switch 32767 {
                case 1:
                    for fi63 := 0; fi63 < 15; fi63++ {
                        var v64 int8 = -128
                        var v65 bool = true
                        for fi66 := 0; fi66 < 6; fi66++ {
                            v3 -= (v3 & (v3 + 527))
                            v5 = (("hello" + ("\n\t" + v5)) + "foo bar")
                            v3 = v3
                            _ = fi66
                        }
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                case 199:
                    _ = len(v2)
                default:
                    if (3.14159 > 3.14159) {
                        if false {
                            if false {
                                fmt.Println(v21, v4, v3)
                                switch "a" {
                                case "a":
                                    v3 += (v3 >> (((v4 * v4)) & 63))
                                }
                                for fi67 := 0; fi67 < 2; fi67++ {
                                    switch 2147483647 {
                                    case 0:
                                        v4 = (v4 * v4)
                                    case 507:
                                        v68 := GenS5{F0: uint16(v4), F1: 82}
                                        _ = v68
                                    case 460:
                                        if false {
                                            if false {
                                                v3 += (v3 / (((v3 * (v3 & -9223372036854775808))) | 1))
                                                for fi69 := 0; fi69 < 3; fi69++ {
                                                    if ((42 <= 175) || true) {
                                                        fmt.Println(v21, fi67)
                                                        _ = len(v2)
                                                        if false {
                                                            var v70 float32 = -2.25
                                                            v70 = -2.25
                                                            v70 -= v70
                                                            _ = v70
                                                        } else {
                                                            v71 := GenS5{F0: 24, F1: 1}
                                                            if false {
                                                                v21 -= 32767
                                                                fmt.Println(v61)
                                                            } else {
                                                                var v72 uint = (v4 - 2)
                                                                v3 += v3
                                                                fmt.Println(v3, v3)
                                                                _ = v72
                                                            }
                                                            var v73 string = "\n\t"
                                                            fmt.Println(fi67)
                                                            v3 += v3
                                                            _ = v71
                                                            _ = v73
                                                        }
                                                    }
                                                    var v74 uint64 = 18446744073709551615
                                                    _ = fi69
                                                    _ = v74
                                                }
                                                var v75 bool = true
                                                _ = v75
                                            } else {
                                                for fi76 := 0; fi76 < 9; fi76++ {
                                                    var v77 int = fi76
                                                    v5 = v5
                                                    _ = fi76
                                                    _ = v77
                                                }
                                                v1[v5] = false
                                                v78 := GenS6{F0: false}
                                                _ = v78
                                            }
                                        } else {
                                            fmt.Println(fi67, v4, v3)
                                            v5 = v5
                                            if (-2 == v3) {
                                                fmt.Println(v61, fi67, v21)
                                            }
                                        }
                                    }
                                    _ = fi67
                                }
                            } else {
                                var v79 uint64 = 42
                                _ = v79
                            }
                            if true {
                                for fi80 := 0; fi80 < 10; fi80++ {
                                    v81 := GenS7{F0: 42, F1: 610, F2: 128}
                                    v82 := GenS1{F0: (fi80 << ((v4) & 63))}
                                    if true {
                                        for fi83 := 0; fi83 < 1; fi83++ {
                                            _ = len(v2)
                                            _ = fi83
                                        }
                                        v84 := GenS8{F0: 255, F1: v21, F2: -9223372036854775808}
                                        if (!false) {
                                            v21 = 16
                                            var v85 int32 = 273
                                            fmt.Println(v5)
                                            _ = v85
                                        } else {
                                            if false {
                                                fmt.Println(v3, v61, v21)
                                                v5 = (v5 + ((v5 + "a") + v5))
                                                fmt.Println(v3, fi80, v61)
                                            } else {
                                                fmt.Println(fi80, v61, v4)
                                                switch v21 {
                                                case 1:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                default:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v86 := GenS9{F0: 1, F1: 25, F2: -32768}
                                                        _ = v86
                                                    }
                                                }
                                                _ = len(v2)
                                            }
                                            v87 := GenS9{F0: 2147483647, F1: 27, F2: -32768}
                                            _ = v87
                                        }
                                        _ = v84
                                    } else {
                                        for fi88 := 0; fi88 < 9; fi88++ {
                                            if true {
                                                var v89 int64 = v3
                                                v90 := GenS2{F0: 128, F1: -1}
                                                if true {
                                                    v61 += v61
                                                }
                                                _ = v89
                                                _ = v90
                                            } else {
                                                if (!((v5 + v5) > "\n\t")) {
                                                    var v91 uint64 = 18446744073709551615
                                                    _ = v91
                                                } else {
                                                    fmt.Println(fi80)
                                                }
                                                v92 := GenS10{F0: 304}
                                                v61 = v61
                                                if (127 > 75) {
                                                    v3 -= (-(v3))
                                                    v4 = v4
                                                }
                                                v61 = 1e10
                                                v3 = 9223372036854775807
                                                _ = v92
                                            }
                                            var v93 uint8 = 105
                                            if (-9223372036854775808 != (fi80 / ((int(v93)) | 1))) {
                                                v94 := GenS11{F0: -32768}
                                                var v95 int8 = -1
                                                _ = v94
                                                _ = v95
                                            } else {
                                                switch (v5 + v5) {
                                                case "hello":
                                                    if ((1 == 2147483647) && false) {
                                                        switch ("\n\t" + ((v5 + v5) + v5)) {
                                                        case "hello":
                                                            v2 = append(v2, v5)
                                                        case "a":
                                                            fmt.Println(v21, v3)
                                                        default:
                                                            var v96 uint8 = v93
                                                            _ = v96
                                                        }
                                                        switch v5 {
                                                        case "hello":
                                                            fmt.Println(v5)
                                                        }
                                                    }
                                                    fmt.Println(fi88, v4)
                                                case "":
                                                    var v97 uint = v4
                                                    _ = v97
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = fi88
                                            _ = v93
                                        }
                                    }
                                    _ = fi80
                                    _ = v81
                                    _ = v82
                                }
                                if true {
                                    switch v5 {
                                    case "":
                                        if (18446744073709551615 != 419) {
                                            switch v21 {
                                            case 4294967295:
                                                var v98 uint64 = 1
                                                _ = v98
                                            default:
                                                if true {
                                                    v99 := GenS7{F0: 993, F1: v21, F2: 0}
                                                    var v100 string = v5
                                                    _ = v100
                                                    _ = v99
                                                } else {
                                                    for fi101 := 0; fi101 < 2; fi101++ {
                                                        fmt.Println(v3, fi101)
                                                        if true {
                                                            v102 := GenS12{F0: 2, F1: -1, F2: 42}
                                                            for fi103 := 0; fi103 < 15; fi103++ {
                                                                var v104 int16 = 0
                                                                _ = fi103
                                                                _ = v104
                                                            }
                                                            _ = v102
                                                        }
                                                        if (((v4 >> ((v4) & 63)) < (v4 ^ 65535)) && false) {
                                                            var v105 uint8 = 2
                                                            fmt.Println(fi101, v4)
                                                            var v106 bool = ((0.1 / 1.5) >= 1e10)
                                                            _ = v105
                                                            _ = v106
                                                        } else {
                                                            if (float64(float32((v61 + (v61 * v61)))) == (v61 - (((v61 / v61) + (v61 * v61)) / (((1.5 / v61) / (v61 - 0.0)) + v61)))) {
                                                                v5 = v5
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                fmt.Println(fi101, v4, v61)
                                                            }
                                                            var v107 int = fi101
                                                            for fi108 := 0; fi108 < 11; fi108++ {
                                                                v3 -= 42
                                                                switch v5 {
                                                                case "hello":
                                                                    if (true && (307 == 324)) {
                                                                        v5 = ""
                                                                        v2 = append(v2, v5)
                                                                        v107 = 4294967295
                                                                    } else {
                                                                        switch v21 {
                                                                        case 32767:
                                                                            _ = len(v2)
                                                                        case -9223372036854775808:
                                                                            v21 -= (v107 & (fi108 + (fi108 - (v21 - (v21 | v107)))))
                                                                        }
                                                                        _ = len(v1)
                                                                        var v109 float64 = v61
                                                                        _ = v109
                                                                    }
                                                                default:
                                                                    v1[(v5 + (("foo bar" + (v5 + v5)) + v5))] = true
                                                                }
                                                                v110 := GenS5{F0: 255, F1: -25}
                                                                _ = fi108
                                                                _ = v110
                                                            }
                                                            _ = v107
                                                        }
                                                        _ = fi101
                                                    }
                                                    if (2147483647 > -658) {
                                                        _ = len(v2)
                                                        fmt.Println(v21, v4, v61)
                                                    } else {
                                                        switch v21 {
                                                        case -2147483648:
                                                            var v111 int8 = -128
                                                            _ = v111
                                                        }
                                                    }
                                                }
                                            }
                                            var v112 int32 = 1
                                            var v113 bool = (((v5 == "") || true) && (v4 != (^(v4))))
                                            _ = v112
                                            _ = v113
                                        }
                                    case "foo bar":
                                        v114 := GenS13{F0: 203, F1: 1}
                                        _ = v114
                                    }
                                    var v115 string = ((("hello" + v5) + v5) + v5)
                                    _ = v115
                                }
                                var v116 bool = (!((v3 % (((v3 - (v3 >> (((v4 >> ((18446744073709551615) & 63))) & 63)))) | 1)) >= v3))
                                _ = v116
                            } else {
                                for fi117 := 0; fi117 < 2; fi117++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if true {
                                        v118 := GenS12{F0: 0, F1: 668, F2: 0}
                                        _ = v118
                                    }
                                    _ = fi117
                                }
                                v4 -= 4294967295
                            }
                            for fi119 := 0; fi119 < 3; fi119++ {
                                var v120 int = (^(fi119))
                                v2 = append(v2, v5)
                                _ = fi119
                                _ = v120
                            }
                        } else {
                            if false {
                                v21 = (v21 % ((int(v4)) | 1))
                                if (!(128 >= 197)) {
                                    v1[v5] = (false && (0.0 < -2.25))
                                    switch v21 {
                                    case 275:
                                        var v121 uint8 = 255
                                        _ = v121
                                    case 32768:
                                        if true {
                                            var v122 uint32 = 4294967295
                                            _ = v122
                                        }
                                    }
                                    var v123 uint32 = 65535
                                    _ = v123
                                } else {
                                    switch v5 {
                                    case "":
                                        _ = len(v2)
                                    case "foo bar":
                                        v124 := GenS14{F0: 2147483647, F1: -1, F2: 4294967295}
                                        _ = v124
                                    }
                                    if (1 == 32767) {
                                        var v125 float32 = (1.5 - 1.5)
                                        _ = v125
                                    } else {
                                        v4 = 42
                                    }
                                }
                            } else {
                                if false {
                                    var v126 int = v21
                                    _ = v126
                                } else {
                                    fmt.Println(v21, v5)
                                }
                                for fi127 := 0; fi127 < 11; fi127++ {
                                    var v128 uint64 = 1
                                    _ = fi127
                                    _ = v128
                                }
                            }
                        }
                        v129 := v1[v5]
                        _ = v129
                    }
                }
            }
            var v130 uint8 = 211
            _ = v130
        }
    case "hello":
        v3 = (v3 * 417)
    }
    if (!false) {
        v4 = (v4 | 389)
    }
    v4 -= (v4 * 914)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v61
}
