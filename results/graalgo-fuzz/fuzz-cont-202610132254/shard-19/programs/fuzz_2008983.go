package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int8; F1 int; F2 string }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS9 struct { F0 string; F1 int8; F2 int16 }
type GenS10 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0(p0 uint16, p1 uint64, p2 string) string {
    if (1 >= 2147483647) {
        fmt.Println(p0, p0, p2)
    }
    p2 = p2
    p0 -= (^(p0))
    v6 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "\n\t"
}

func main() {
    var v1 int16 = 0
    var v2 float64 = -2.25
    v3 := []string{"a", "a", "", "foo bar"}
    var v4 string = "foo bar"
    var v5 int16 = (v1 - -32768)
    switch (v4 + ("" + ("foo bar" + v4))) {
    case "\n\t":
        fmt.Println(v4, v4)
    case "foo bar":
        var v7 int64 = 9223372036854775807
        _ = v7
    case "a":
        if len(v3) > 0 {
            v3[0] = (v4 + v4)
        }
    }
    v2 = float64(-2.25)
    if true {
        v8 := GenS2{F0: -824, F1: 37}
        v1 -= (v1 + 1)
        _ = v8
    } else {
        v4 = ("a" + v4)
    }
    v4 = v4
    if true {
        var v9 float32 = (1e10 * -2.25)
        var v10 bool = (true || (!false))
        for fi11 := 0; fi11 < 19; fi11++ {
            _ = len(v3)
            var v12 float64 = 0.1
            var v13 float64 = float64((v9 + 0.1))
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    } else {
        var v14 float32 = 1e10
        var v15 uint = 0
        _ = v14
        _ = v15
    }
    var v16 int = 9223372036854775807
    fmt.Println(v16)
    if (((float32(-2.25) * ((0.1 / 1.0) / float32(((((1.5 - v2) * 1e10) - v2) * (v2 + (float64(float32((v2 / (v2 / v2)))) + 1e10)))))) / -2.25) <= (3.14159 - float32(float64(1e10)))) {
        v4 = ("\n\t" + v4)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v17 int16 = -565
        _ = v17
    } else {
        v18 := GenS2{F0: (v16 | -937), F1: 65535}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (0.0 > 1.5) {
            var v19 bool = (65535 != 1)
            var v20 int = 128
            v21 := GenS2{F0: v20, F1: 32767}
            _ = v19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    if false {
        v22 := GenS2{F0: (v16 % ((v16) | 1)), F1: 263}
        _ = v22
    }
    if false {
        fmt.Println(v2, v1)
        _ = len(v3)
        var v23 float32 = ((3.14159 + 1e10) * 1.5)
        _ = v23
    } else {
        var v24 uint64 = 18446744073709551615
        if true {
            if (42 <= 65535) {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v1 -= (v5 / (((v1 % (((v1 - (v5 + (v1 & v1)))) | 1))) | 1))
                    for fi26 := 0; fi26 < 16; fi26++ {
                        if ((!true) && (18446744073709551615 >= 260)) {
                            fmt.Println(v5)
                        }
                        _ = fi26
                    }
                    v27 := GenS3{F0: int16(v16), F1: (v4 + v4), F2: 964}
                    _ = fi25
                    _ = v27
                }
                if len(v3) > 0 {
                    v3[0] = v4
                }
                if (true && (!false)) {
                    fmt.Println(v1)
                } else {
                    if false {
                        v4 = ("foo bar" + (("a" + v4) + v4))
                        switch v4 {
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "hello":
                            if true {
                                v28 := GenS4{F0: 2, F1: (^(v16))}
                                v1 -= (v1 ^ v5)
                                _ = v28
                            }
                        default:
                            v16 += 159
                        }
                    }
                    if (!true) {
                        switch v16 {
                        case 0:
                            if (1e10 > (v2 / float64(((3.14159 - 0.0) * 3.14159)))) {
                                v1 += v1
                                var v29 bool = false
                                _ = v29
                            } else {
                                var v30 int32 = 80
                                v30 = v30
                                v16 -= (v16 / (((v16 | v16)) | 1))
                                _ = v30
                            }
                        }
                        var v31 float32 = -2.25
                        _ = v31
                    }
                }
            } else {
                fmt.Println(v4, v4, v24)
                fmt.Println(v24)
                for fi32 := 0; fi32 < 1; fi32++ {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    v33 := GenS2{F0: (v16 * (v16 / ((fi32) | 1))), F1: 0}
                    var v34 int16 = (-(v5))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            v5 -= (v1 | (v1 / (((v1 + v5)) | 1)))
            var v35 uint64 = (v24 | v24)
            _ = v35
        } else {
            if false {
                var v36 string = (v4 + (("a" + "") + v4))
                v4 = v4
                if false {
                    if true {
                        switch (v16 | v16) {
                        case 2147483647:
                            v5 -= (v1 << ((423) & 63))
                        default:
                            var v37 uint8 = 127
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 15; fi38++ {
                            for fi39 := 0; fi39 < 6; fi39++ {
                                var v40 int16 = 42
                                var v41 string = ("foo bar" + v36)
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            _ = fi38
                        }
                        for fi42 := 0; fi42 < 2; fi42++ {
                            v16 = v16
                            _ = fi42
                        }
                    } else {
                        var v43 uint16 = 801
                        v44 := GenS1{F0: -128}
                        v45 := GenS5{F0: 100, F1: v16, F2: v4}
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                } else {
                    v46 := GenS6{F0: 9223372036854775807, F1: 42}
                    var v47 uint16 = 535
                    var v48 bool = (1 != 72)
                    _ = v46
                    _ = v47
                    _ = v48
                }
                _ = v36
            } else {
                var v49 int32 = 463
                if false {
                    if false {
                        var v50 uint16 = 1
                        if true {
                            v5 += v5
                            v16 += (v16 * (v16 % (((v16 << ((255) & 63))) | 1)))
                        }
                        v51 := GenS1{F0: 1}
                        _ = v50
                        _ = v51
                    } else {
                        if false {
                            var v52 uint16 = 286
                            for fi53 := 0; fi53 < 2; fi53++ {
                                if (-2.25 < -2.25) {
                                    fmt.Println(v49)
                                    v1 -= v5
                                } else {
                                    v49 += 275
                                    fmt.Println(v2, v2, v2)
                                    v54 := GenS1{F0: 100}
                                    _ = v54
                                }
                                _ = fi53
                            }
                            _ = v52
                        }
                        for fi55 := 0; fi55 < 7; fi55++ {
                            if (!(!false)) {
                                if (42 <= 363) {
                                    if false {
                                        v56 := GenS3{F0: (v1 - (-(v5))), F1: "a", F2: 340}
                                        v2 = v2
                                        _ = v56
                                    } else {
                                        var v57 float32 = (0.0 - (1e10 - 0.0))
                                        var v58 uint32 = 42
                                        v59 := GenS4{F0: v24, F1: -2147483648}
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    }
                                } else {
                                    for fi60 := 0; fi60 < 14; fi60++ {
                                        if true {
                                            var v61 int8 = 127
                                            v62 := GenS5{F0: 127, F1: v16, F2: v4}
                                            v63 := GenS6{F0: 1, F1: (v61 * (v61 & v61))}
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        }
                                        switch ((v4 + "a") + "") {
                                        case "\n\t":
                                            v5 = 75
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        _ = fi60
                                    }
                                    for fi64 := 0; fi64 < 8; fi64++ {
                                        if true {
                                            v65 := GenS4{F0: (v24 | (v24 / ((v24) | 1))), F1: (fi64 / ((int(v49)) | 1))}
                                            v24 = (v24 % ((2) | 1))
                                            v66 := GenS1{F0: -115}
                                            v67 := GenS2{F0: fi55, F1: 1}
                                            if (-2.25 == 0.1) {
                                                switch v4 {
                                                case "hello":
                                                    _ = len(v3)
                                                case "\n\t":
                                                    var v68 uint16 = 65535
                                                    var v69 uint16 = (v68 ^ v68)
                                                    if ((-1 > 42) && (((v2 - ((v2 / v2) * v2)) != v2) || false)) {
                                                        v70 := GenS7{F0: 1, F1: -128}
                                                        if len(v3) > 0 {
                                                            v3[0] = v4
                                                        }
                                                        switch v4 {
                                                        case "":
                                                            v69 = (v68 + 42)
                                                        case "hello":
                                                            var v71 int32 = v49
                                                            _ = v71
                                                        case "foo bar":
                                                            switch (v4 + "a") {
                                                            case "\n\t":
                                                                fmt.Println(v69, v69, v24)
                                                            default:
                                                                var v72 float32 = (1.5 - float32((float64(0.0) * float64(0.0))))
                                                                _ = v72
                                                            }
                                                        }
                                                        _ = v70
                                                    } else {
                                                        v4 = v4
                                                        v24 -= v24
                                                        var v73 uint64 = 199
                                                        _ = v73
                                                    }
                                                    _ = v68
                                                    _ = v69
                                                case "a":
                                                    if true {
                                                        var v74 int16 = v1
                                                        fmt.Println(v74, fi55)
                                                        v3 = append(v3, v4)
                                                        _ = v74
                                                    } else {
                                                        var v75 int = (v16 << ((2) & 63))
                                                        _ = v75
                                                    }
                                                default:
                                                    var v76 int8 = 1
                                                    _ = v76
                                                }
                                                var v77 string = ("a" + ("\n\t" + (v4 + (v4 + v4))))
                                                _ = v77
                                            } else {
                                                v3 = append(v3, v4)
                                                v78 := GenS4{F0: v24, F1: 654}
                                                _ = v78
                                            }
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = fi64
                                    }
                                    switch "" {
                                    case "foo bar":
                                        var v79 float32 = (float32(v2) / 1.0)
                                        _ = v79
                                    case "\n\t":
                                        var v80 int64 = -9223372036854775808
                                        _ = v80
                                    case "":
                                        var v81 float32 = (1e10 + -2.25)
                                        _ = v81
                                    }
                                }
                                var v82 float64 = 0.0
                                v83 := GenS1{F0: 74}
                                _ = v82
                                _ = v83
                            } else {
                                v84 := GenS6{F0: 0, F1: -128}
                                for fi85 := 0; fi85 < 10; fi85++ {
                                    switch v16 {
                                    case 2147483647:
                                        var v86 float64 = v2
                                        _ = v86
                                    }
                                    if true {
                                        v87 := GenS8{F0: 42, F1: 1, F2: 4294967295}
                                        v3 = append(v3, v4)
                                        fmt.Println(v24, v16, v5)
                                        _ = v87
                                    } else {
                                        v3 = append(v3, "")
                                        if false {
                                            fmt.Println(v2)
                                            fmt.Println(v4, v24, v16)
                                            var v88 int16 = 564
                                            _ = v88
                                        }
                                    }
                                    _ = fi85
                                }
                                _ = len(v3)
                                _ = v84
                            }
                            v3 = append(v3, "hello")
                            _ = fi55
                        }
                    }
                    for fi89 := 0; fi89 < 20; fi89++ {
                        if (42 != 42) {
                            v90 := GenS6{F0: 0, F1: 42}
                            _ = v90
                        } else {
                            if ((!true) || false) {
                                if (true || (255 >= 65535)) {
                                    fmt.Println(v2)
                                    for fi91 := 0; fi91 < 8; fi91++ {
                                        var v92 int8 = 1
                                        _ = fi91
                                        _ = v92
                                    }
                                    if false {
                                        var v93 float64 = float64(1.5)
                                        var v94 uint8 = 42
                                        _ = v93
                                        _ = v94
                                    }
                                } else {
                                    if (v2 > v2) {
                                        fmt.Println(v16, v49, v16)
                                    } else {
                                        if (4294967295 != (v24 << ((0) & 63))) {
                                            var v95 uint64 = 1
                                            _ = v95
                                        } else {
                                            v4 = v4
                                        }
                                        v4 = "\n\t"
                                        for fi96 := 0; fi96 < 12; fi96++ {
                                            var v97 int8 = 100
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if true {
                                                fmt.Println(v16, v2)
                                            }
                                            var v98 int32 = v49
                                            _ = fi96
                                            _ = v97
                                            _ = v98
                                        }
                                    }
                                    v49 -= 42
                                    v99 := GenS5{F0: -30, F1: (-(fi89)), F2: v4}
                                    _ = v99
                                }
                                if true {
                                    if (v4 > v4) {
                                        if len(v3) > 0 {
                                            v3[0] = ("foo bar" + v4)
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                    } else {
                                        if false {
                                            if (-2.25 <= (v2 - float64(((1e10 + 0.1) - (1.5 * (float32(v2) / -2.25)))))) {
                                                var v100 uint32 = 65535
                                                _ = v100
                                            }
                                            _ = len(v3)
                                        }
                                    }
                                }
                            }
                        }
                        var v101 int8 = 127
                        _ = fi89
                        _ = v101
                    }
                    fmt.Println(v2, v16, v24)
                }
                var v102 int = v16
                _ = v102
                _ = v49
            }
            var v103 uint = 4294967295
            _ = v103
        }
        var v104 uint32 = 826
        _ = v104
        _ = v24
    }
    if (((float64(1e10) * v2) - (v2 / ((v2 * 1e10) / 1.0))) >= float64(0.0)) {
        var v105 int = (v16 % ((v16) | 1))
        for fi106 := 0; fi106 < 4; fi106++ {
            var v107 int = (-(v105))
            _ = fi106
            _ = v107
        }
        var v108 float32 = (((-2.25 + 3.14159) * 1.5) + -2.25)
        _ = v105
        _ = v108
    }
    v109 := GenS5{F0: 127, F1: -2, F2: "a"}
    var v110 string = v4
    v110 = v110
    fmt.Println(v16)
    v110 = v110
    switch v4 {
    case "\n\t":
        v5 -= v1
    }
    if true {
        v16 -= v16
        fmt.Println(v16, v110)
    } else {
        var v111 uint8 = 80
        if true {
            var v112 int32 = -1
            var v113 uint = 18446744073709551615
            if false {
                switch v16 {
                case 2147483647:
                    var v114 int32 = (v112 << ((1) & 63))
                    _ = v114
                case 9223372036854775807:
                    var v115 uint64 = 0
                    _ = v115
                case -2147483648:
                    if (679 != uint32(v111)) {
                        for fi116 := 0; fi116 < 8; fi116++ {
                            v117 := GenS1{F0: 0}
                            for fi118 := 0; fi118 < 13; fi118++ {
                                v110 = v4
                                var v119 float32 = 3.14159
                                var v120 int = (fi116 & (fi116 & (fi116 & v16)))
                                _ = fi118
                                _ = v119
                                _ = v120
                            }
                            _ = fi116
                            _ = v117
                        }
                    }
                }
                var v121 uint = (v113 >> ((2) & 63))
                _ = v121
            }
            _ = v112
            _ = v113
        } else {
            if (255 <= 2) {
                _ = len(v3)
                v122 := GenS9{F0: v110, F1: 0, F2: (v1 & v1)}
                var v123 bool = true
                _ = v122
                _ = v123
            } else {
                for fi124 := 0; fi124 < 1; fi124++ {
                    fmt.Println(v1, v4, fi124)
                    if true {
                        v125 := GenS10{F0: 2, F1: true, F2: 2}
                        var v126 uint8 = (v111 % ((v111) | 1))
                        var v127 uint64 = 2
                        _ = v125
                        _ = v126
                        _ = v127
                    }
                    _ = fi124
                }
                var v128 float64 = v2
                _ = v128
            }
            switch -112 {
            case 127:
                var v129 uint8 = (v111 - (v111 | (v111 + v111)))
                _ = v129
            case 338:
                var v130 uint16 = 0
                _ = v130
            default:
                fmt.Println(v110, v16)
            }
        }
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v109
    _ = v110
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
