package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 uint }
type GenS12 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS13 struct { F0 uint; F1 uint }
type GenS14 struct { F0 uint8; F1 uint }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int8; F1 uint }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 bool; F1 bool }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int; F1 int16 }
type GenS21 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 bool; F1 int16 }
type GenS24 struct { F0 uint8; F1 int8 }
type GenS25 struct { F0 uint64 }
type GenS26 struct { F0 uint64; F1 int; F2 int8 }

func main() {
    v1 := []uint32{440, 584, 405}
    var v2 uint16 = 65535
    v3 := make(map[string]bool)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 -= (v2 >> ((1) & 63))
    var v4 string = ""
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 bool = true
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        if len(v1) > 0 {
            v1[0] = 431
        }
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 2, F1: (v2 & 65535)}
    for fi8 := 0; fi8 < 18; fi8++ {
        v2 += 255
        var v9 uint = 2
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        switch (v4 + "foo bar") {
        case "hello":
            var v11 int64 = -409
            _ = v11
        default:
            switch fi10 {
            case 9223372036854775807:
                var v12 float64 = float64((float32(0.1) * 0.0))
                _ = v12
            default:
                if true {
                    var v13 uint64 = 4294967295
                    v4 = "hello"
                    fmt.Println(fi10, v13, v4)
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 int = fi10
                        v16 := GenS2{F0: 1}
                        if false {
                            if false {
                                if true {
                                    v3[v4] = (true && true)
                                    v15 += (fi14 % ((32768) | 1))
                                    v15 = v15
                                }
                                if ((v4 + (v4 + v4)) != "hello") {
                                    delete(v3, (v4 + v4))
                                    var v17 uint = 42
                                    _ = v17
                                } else {
                                    fmt.Println(v15, fi14)
                                    if false {
                                        var v18 int16 = 32767
                                        switch (-(v15)) {
                                        case 42:
                                            switch fi14 {
                                            case -2147483648:
                                                if (636 == 128) {
                                                    var v19 float64 = 3.14159
                                                    var v20 int32 = -450
                                                    _ = v19
                                                    _ = v20
                                                }
                                                v18 = -32768
                                            }
                                        case -464:
                                            fmt.Println(v18, fi14, v4)
                                        }
                                        _ = v18
                                    }
                                }
                                if false {
                                    v4 = ("\n\t" + "")
                                    _ = len(v3)
                                    if ((fi10 * (v15 % ((fi14) | 1))) == (^(fi10))) {
                                        switch fi14 {
                                        case 905:
                                            var v21 string = ""
                                            _ = v21
                                        case -501:
                                            if true {
                                                v4 = (v4 + ("a" + v4))
                                                switch fi10 {
                                                case -379:
                                                    var v22 int8 = 127
                                                    _ = v22
                                                case 128:
                                                    var v23 float64 = 0.0
                                                    _ = v23
                                                default:
                                                    v4 = ""
                                                }
                                                v15 += fi10
                                            } else {
                                                if (65535 == 2) {
                                                    var v24 uint16 = 32767
                                                    var v25 uint64 = 42
                                                    v24 -= (v2 >> ((42) & 63))
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    switch (v15 ^ fi14) {
                                                    case 32768:
                                                        v4 = ("a" + "")
                                                    }
                                                    v15 -= v15
                                                }
                                                v2 -= v2
                                            }
                                        case 1:
                                            v15 += 32768
                                        }
                                        v26 := GenS3{F0: 0, F1: 2147483647, F2: -455}
                                        if (219 >= -9223372036854775808) {
                                            var v27 uint64 = 42
                                            var v28 float64 = (3.14159 * 3.14159)
                                            fmt.Println(v27, v27, v28)
                                            _ = v27
                                            _ = v28
                                        } else {
                                            var v29 int16 = 42
                                            if len(v1) > 0 {
                                                v1[0] = 4294967295
                                            }
                                            if (v15 > (fi14 | (fi14 >> ((255) & 63)))) {
                                                v30 := GenS3{F0: -310, F1: 65535, F2: 42}
                                                _ = v30
                                            }
                                            _ = v29
                                        }
                                        _ = v26
                                    }
                                }
                                v4 = v4
                                fmt.Println(fi10, v2, fi14)
                            } else {
                                v15 = v15
                            }
                            fmt.Println(v2, v15)
                            v31 := GenS4{F0: 42}
                            _ = v31
                        } else {
                            v1 = append(v1, 0)
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    var v32 int8 = 0
                    _ = v32
                }
            }
        }
        fmt.Println(v4, v4, fi10)
        v4 = ""
        _ = fi10
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch fi33 {
            case -2147483648:
                v2 += v2
            default:
                var v35 int = (fi33 - -1)
                _ = v35
            }
            if true {
                var v36 int32 = 244
                if false {
                    v36 -= -1
                }
                _ = v36
            }
            _ = fi34
        }
        v37 := GenS5{F0: 997, F1: v4}
        _ = fi33
        _ = v37
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        if (!((float32(3.14159) < (1e10 * -2.25)) || false)) {
            _ = len(v3)
            for fi39 := 0; fi39 < 17; fi39++ {
                v2 += v2
                var v40 int16 = 521
                if (42 < 898) {
                    fmt.Println(v4, v40, fi38)
                    var v41 uint64 = 998
                    _ = v41
                } else {
                    var v42 uint16 = v2
                    fmt.Println(v40, v4, v42)
                    if (2147483647 == 2147483647) {
                        v2 = (v42 * 255)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        var v43 int = fi38
                        if (0 < -128) {
                            if true {
                                if ((v43 <= (-(fi39))) && false) {
                                    var v44 uint8 = 42
                                    if (833 == 4294967295) {
                                        fmt.Println(v44, v42, fi38)
                                        if ((v40 / ((32767) | 1)) != v40) {
                                            v42 += (v42 >> ((18446744073709551615) & 63))
                                            switch v4 {
                                            case "hello":
                                                switch ((v4 + v4) + (v4 + ("hello" + v4))) {
                                                case "":
                                                    var v45 int64 = 42
                                                    _ = v45
                                                }
                                            case "a":
                                                var v46 uint32 = 4294967295
                                                _ = v46
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            fmt.Println(v44, fi39, v2)
                                        }
                                    }
                                    _ = v44
                                }
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                            } else {
                                v42 += (v42 | (^(v2)))
                                if true {
                                    v2 += v42
                                } else {
                                    var v47 int16 = v40
                                    v3[((v4 + (v4 + ("a" + "\n\t"))) + (v4 + v4))] = (v2 <= v2)
                                    _ = v47
                                }
                            }
                            switch v4 {
                            case "\n\t":
                                switch (v4 + v4) {
                                case "foo bar":
                                    v2 -= v2
                                case "hello":
                                    v4 = (((v4 + (v4 + v4)) + v4) + "a")
                                case "":
                                    var v48 bool = (true && (v2 >= v42))
                                    _ = v48
                                default:
                                    if true {
                                        v43 -= -428
                                        v2 -= (v42 * 2)
                                    }
                                }
                            case "hello":
                                var v49 int16 = v40
                                _ = v49
                            default:
                                fmt.Println(v4, fi39, v2)
                            }
                        } else {
                            switch v4 {
                            case "foo bar":
                                if false {
                                    fmt.Println(v4, v4)
                                    if (v2 <= v42) {
                                        v43 += (v43 * 128)
                                        v4 = v4
                                        if (-74 != -4) {
                                            v50 := GenS3{F0: 9223372036854775807, F1: 2, F2: -9223372036854775808}
                                            if false {
                                                switch "foo bar" {
                                                case "hello":
                                                    switch (((v4 + "") + v4) + ("hello" + ((v4 + v4) + v4))) {
                                                    case "":
                                                        switch v4 {
                                                        case "hello":
                                                            v43 = (v43 | (fi38 % (((^(fi39))) | 1)))
                                                        case "foo bar":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    v51 := GenS6{F0: 65535, F1: v2, F2: 313}
                                                    _ = v51
                                                case "foo bar":
                                                    if true {
                                                        v52 := GenS7{F0: 421}
                                                        fmt.Println(v43, v43, fi38)
                                                        if (uint(v2) == 42) {
                                                            v43 = fi38
                                                            if false {
                                                                switch (v4 + (v4 + v4)) {
                                                                case "a":
                                                                    v53 := GenS7{F0: 158}
                                                                    _ = v53
                                                                case "\n\t":
                                                                    var v54 uint16 = v2
                                                                    _ = v54
                                                                }
                                                                var v55 bool = true
                                                                v2 -= v42
                                                                fmt.Println(v2, v55, v40)
                                                                _ = v55
                                                            }
                                                        }
                                                        _ = v52
                                                    } else {
                                                        v1 = append(v1, 1)
                                                        v3[v4] = (fi39 <= (v43 << ((14) & 63)))
                                                        v56 := GenS8{F0: 357, F1: v42}
                                                        _ = v56
                                                    }
                                                }
                                                v42 = (v42 * 2)
                                                if true {
                                                    var v57 int64 = -1
                                                    _ = v57
                                                } else {
                                                    var v58 float32 = 1e10
                                                    _ = v58
                                                }
                                                v59 := v3[v4]
                                                _ = v59
                                            } else {
                                                var v60 float64 = float64(1e10)
                                                var v61 uint16 = v42
                                                _ = v60
                                                _ = v61
                                            }
                                            v40 -= -760
                                            _ = v50
                                        }
                                        v62 := GenS9{F0: 1, F1: false}
                                        v40 = v40
                                        _ = v62
                                    }
                                } else {
                                    v63 := GenS8{F0: 2147483648, F1: v42}
                                    fmt.Println(fi39, v43)
                                    _ = v63
                                }
                            }
                            v2 -= 946
                            var v64 int = (fi39 - 127)
                            v43 = fi38
                            _ = v64
                        }
                        v1 = append(v1, 535)
                        _ = v43
                    }
                    _ = v42
                }
                _ = fi39
                _ = v40
            }
        } else {
            v65 := GenS2{F0: -32768}
            if false {
                for fi66 := 0; fi66 < 9; fi66++ {
                    if false {
                        v4 = v4
                        v4 = ("a" + v4)
                        v2 = (v2 >> ((1000) & 63))
                    } else {
                        v4 = v4
                        var v67 uint = 18446744073709551615
                        _ = v67
                    }
                    v68 := GenS10{F0: v2}
                    v69 := GenS1{F0: 962, F1: (v2 >> ((18446744073709551615) & 63))}
                    if len(v1) > 0 {
                        v1[0] = 923
                    }
                    _ = fi66
                    _ = v68
                    _ = v69
                }
                v2 -= (v2 % (((v2 ^ 2)) | 1))
            } else {
                var v70 float32 = 3.14159
                var v71 uint16 = (v2 | 771)
                _ = v70
                _ = v71
            }
            _ = len(v3)
            _ = v65
        }
        for fi72 := 0; fi72 < 5; fi72++ {
            var v73 float64 = 1.5
            if (!(0 == -117)) {
                fmt.Println(v73, v73, v4)
                var v74 int = fi38
                _ = v74
            }
            var v75 int8 = -77
            if true {
                var v76 uint = 59
                _ = v76
            } else {
                var v77 uint = 895
                _ = v77
            }
            _ = fi72
            _ = v73
            _ = v75
        }
        var v78 int16 = 32767
        _ = fi38
        _ = v78
    }
    v4 = v4
    fmt.Println(v2, v2)
    if false {
        for fi79 := 0; fi79 < 17; fi79++ {
            switch v4 {
            case "hello":
                var v80 float64 = (1e10 * ((-2.25 - float64(1e10)) - (1e10 / (1.5 / float64((1e10 + (((float32(((1.5 + -2.25) - float64(-2.25))) + 1.5) - (1e10 + 3.14159)) / ((-2.25 * ((((((float32(0.0) / (3.14159 + (0.0 + 0.1))) + ((1.5 * (1e10 - 1e10)) / 1.5)) / 1.0) + (((3.14159 / ((0.1 - 1.5) / 0.1)) * (0.0 - ((float32(0.0) * 0.1) - 1e10))) - 1.5)) - (1.5 / 1.0)) / 3.14159)) / (((1.5 - 0.1) - (((1e10 + 0.0) - 0.1) * 3.14159)) * (3.14159 - (0.1 / float32(1e10))))))))))))
                _ = v80
            case "\n\t":
                if (true && false) {
                    if true {
                        switch "" {
                        case "hello":
                            for fi81 := 0; fi81 < 12; fi81++ {
                                v2 = v2
                                if (2 == 1) {
                                    v4 = ((v4 + "") + (v4 + ((v4 + v4) + ("a" + "a"))))
                                } else {
                                    if false {
                                        var v82 float32 = -2.25
                                        if true {
                                            v83 := GenS4{F0: 4294967295}
                                            var v84 int64 = 9223372036854775807
                                            _ = v83
                                            _ = v84
                                        } else {
                                            v4 = "foo bar"
                                            var v85 uint8 = 133
                                            _ = v85
                                        }
                                        v86 := GenS11{F0: 199, F1: 255}
                                        _ = v82
                                        _ = v86
                                    } else {
                                        var v87 uint = 42
                                        if ("\n\t" > v4) {
                                            v88 := GenS4{F0: v87}
                                            _ = v88
                                        } else {
                                            fmt.Println(v2, v2)
                                            v2 -= (^(v2))
                                            var v89 int32 = 1
                                            _ = v89
                                        }
                                        delete(v3, v4)
                                        _ = v87
                                    }
                                    if false {
                                        delete(v3, (v4 + (v4 + v4)))
                                        var v90 int64 = 42
                                        var v91 int64 = v90
                                        _ = v90
                                        _ = v91
                                    } else {
                                        v92 := GenS10{F0: (^(v2))}
                                        _ = v92
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = len(v1)
                                }
                                switch fi79 {
                                case -685:
                                    var v93 int16 = -794
                                    _ = v93
                                case 42:
                                    v2 = v2
                                }
                                _ = fi81
                            }
                        }
                        v94 := GenS12{F0: 114, F1: 42, F2: uint32(v2)}
                        _ = v94
                    } else {
                        if ((-1 == -1) && (!false)) {
                            v95 := GenS13{F0: 1, F1: 4294967295}
                            switch v4 {
                            case "a":
                                for fi96 := 0; fi96 < 10; fi96++ {
                                    v4 = v4
                                    var v97 int32 = 559
                                    var v98 int = int(v97)
                                    var v99 uint8 = 0
                                    var v100 int32 = (v97 >> ((42) & 63))
                                    v101 := GenS5{F0: 127, F1: (v4 + ((((v4 + "a") + (((v4 + (v4 + (v4 + ""))) + v4) + ("\n\t" + "hello"))) + v4) + "a"))}
                                    var v102 int64 = 1
                                    var v103 int = fi79
                                    v4 = (v4 + v4)
                                    switch v4 {
                                    case "foo bar":
                                        v2 += (v2 - (v2 << ((42) & 63)))
                                    case "a":
                                        if true {
                                            v104 := GenS12{F0: 253, F1: 42, F2: 239}
                                            if false {
                                                v105 := GenS5{F0: 4294967295, F1: v4}
                                                _ = v105
                                            }
                                            _ = v104
                                        } else {
                                            fmt.Println(v103)
                                            v106 := GenS14{F0: 11, F1: 858}
                                            v107 := v3[(v4 + ("a" + "a"))]
                                            fmt.Println(v107)
                                            delete(v3, ((v4 + v4) + "hello"))
                                            _ = v106
                                            _ = v107
                                        }
                                    case "":
                                        fmt.Println(v2, v2, v103)
                                    default:
                                        fmt.Println(fi96)
                                    }
                                    _ = fi96
                                    _ = v100
                                    _ = v101
                                    _ = v102
                                    _ = v103
                                    _ = v97
                                    _ = v98
                                    _ = v99
                                }
                            case "\n\t":
                                for fi108 := 0; fi108 < 12; fi108++ {
                                    if true {
                                        v1 = append(v1, 955)
                                        v1 = append(v1, 550)
                                        v109 := GenS1{F0: 4294967295, F1: v2}
                                        _ = v109
                                    }
                                    _ = fi108
                                }
                            case "foo bar":
                                switch fi79 {
                                case 128:
                                    for fi110 := 0; fi110 < 6; fi110++ {
                                        fmt.Println(v2)
                                        switch v4 {
                                        case "a":
                                            _ = len(v3)
                                        case "foo bar":
                                            if ((fi110 / (((fi79 % ((fi79) | 1))) | 1)) <= fi79) {
                                                var v111 float32 = ((1e10 + (3.14159 * (0.1 * 0.0))) - (0.1 - (-2.25 * 1.5)))
                                                fmt.Println(v4)
                                                v2 = 32767
                                                _ = v111
                                            }
                                        default:
                                            fmt.Println(v2, v4)
                                        }
                                        _ = fi110
                                    }
                                case -2147483648:
                                    v4 = ""
                                case 127:
                                    _ = len(v3)
                                }
                            default:
                                var v112 int8 = 127
                                _ = v112
                            }
                            _ = v95
                        }
                        v113 := GenS15{F0: (v2 | 52)}
                        v4 = v4
                        _ = v113
                    }
                } else {
                    if (-1 >= 290) {
                        for fi114 := 0; fi114 < 12; fi114++ {
                            v4 = ((v4 + v4) + (v4 + (("" + (("\n\t" + ("\n\t" + "")) + v4)) + v4)))
                            fmt.Println(v2, v2, v2)
                            v4 = v4
                            fmt.Println(v4, fi79)
                            v2 -= 0
                            _ = fi114
                        }
                    } else {
                        v115 := GenS16{F0: 42, F1: 335}
                        fmt.Println(v4, v4)
                        _ = v115
                    }
                }
            default:
                var v116 int8 = 8
                _ = v116
            }
            if ((-32768 <= 32767) && true) {
                if false {
                    for fi117 := 0; fi117 < 18; fi117++ {
                        if true {
                            v2 += 42
                            v118 := GenS2{F0: 218}
                            _ = v118
                        } else {
                            fmt.Println(fi79, v2)
                        }
                        delete(v3, (v4 + (v4 + v4)))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v119 uint8 = 255
                        switch v4 {
                        case "foo bar":
                            var v120 uint8 = 255
                            _ = v120
                        }
                        if false {
                            var v121 uint64 = 18446744073709551615
                            switch v4 {
                            case "":
                                if (!((42 < 0) && (false && true))) {
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                    var v122 int32 = -2147483648
                                    v123 := GenS14{F0: v119, F1: 18446744073709551615}
                                    v124 := GenS11{F0: 378, F1: 0}
                                    _ = v122
                                    _ = v123
                                    _ = v124
                                }
                            case "foo bar":
                                v125 := GenS17{F0: 505}
                                _ = v125
                            }
                            _ = v121
                        } else {
                            switch v4 {
                            case "":
                                if ((1.5 - float32(0.0)) > 3.14159) {
                                    v3[((v4 + (v4 + v4)) + v4)] = true
                                }
                            case "\n\t":
                                if (18446744073709551615 < 18446744073709551615) {
                                    var v126 int = (-(fi117))
                                    if false {
                                        var v127 uint32 = 0
                                        if false {
                                            switch v4 {
                                            case "hello":
                                                var v128 float64 = 0.0
                                                _ = v128
                                            case "":
                                                v119 = 1
                                                fmt.Println(fi117, v2)
                                                if false {
                                                    var v129 string = v4
                                                    if (v129 >= v129) {
                                                        v3[v4] = (false && (!(42 >= 1)))
                                                        var v130 int = 128
                                                        _ = v130
                                                    } else {
                                                        var v131 int16 = -284
                                                        _ = len(v1)
                                                        _ = v131
                                                    }
                                                    _ = v129
                                                }
                                            case "\n\t":
                                                if false {
                                                    v132 := GenS7{F0: 1}
                                                    switch ((v4 + v4) + v4) {
                                                    case "\n\t":
                                                        var v133 uint16 = v2
                                                        _ = v133
                                                    }
                                                    v4 = (("a" + v4) + v4)
                                                    _ = v132
                                                }
                                            }
                                            v4 = v4
                                            switch "\n\t" {
                                            case "":
                                                switch "hello" {
                                                case "":
                                                    var v134 uint = 255
                                                    _ = v134
                                                case "\n\t":
                                                    if (v119 < (v119 * (v119 + (v119 + (v119 * (v119 | (v119 % ((2) | 1)))))))) {
                                                        v119 = v119
                                                    }
                                                }
                                            case "a":
                                                var v135 int16 = 1
                                                _ = v135
                                            case "foo bar":
                                                switch "" {
                                                case "hello":
                                                    var v136 uint16 = (v2 - v2)
                                                    v1 = append(v1, 42)
                                                    v137 := GenS18{F0: (!(true && true)), F1: true}
                                                    _ = v136
                                                    _ = v137
                                                case "":
                                                    if (v2 != (v2 ^ 989)) {
                                                        if false {
                                                            if true {
                                                                delete(v3, v4)
                                                                var v138 bool = ((^(v2)) < (v2 % (((v2 >> ((1) & 63))) | 1)))
                                                                if false {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (v127 * v127)
                                                                    }
                                                                }
                                                                _ = v138
                                                            }
                                                            v127 -= 42
                                                            switch 4294967295 {
                                                            case 127:
                                                                var v139 bool = true
                                                                _ = v139
                                                            default:
                                                                _ = len(v3)
                                                            }
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            delete(v3, v4)
                                                        }
                                                        if (1 != 2) {
                                                            fmt.Println(fi117, fi79)
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    fmt.Println(v119, v127)
                                                                    v3[v4] = (v127 > v127)
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v127 -= v127
                                                                        switch (fi79 + 576) {
                                                                        case 1:
                                                                            if true {
                                                                                fmt.Println(v127)
                                                                                v140 := GenS19{F0: 1}
                                                                                _ = v140
                                                                            }
                                                                        case 42:
                                                                            v2 = (v2 << ((uint(v119)) & 63))
                                                                        case -290:
                                                                            if (1e10 > 1.5) {
                                                                                v119 += (v119 - v119)
                                                                                v119 += 128
                                                                            } else {
                                                                                switch ("\n\t" + v4) {
                                                                                case "hello":
                                                                                    var v141 uint64 = 18446744073709551615
                                                                                    _ = v141
                                                                                }
                                                                                var v142 int32 = 0
                                                                                _ = v142
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch "a" {
                                                                    case "":
                                                                        v126 += (fi79 | fi117)
                                                                    case "\n\t":
                                                                        var v143 int64 = -9223372036854775808
                                                                        _ = v143
                                                                    case "hello":
                                                                        if (587 > 42) {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            if false {
                                                                                var v144 bool = (true || (v126 >= 1))
                                                                                _ = v144
                                                                            } else {
                                                                                if (!false) {
                                                                                    var v145 int16 = -657
                                                                                    var v146 uint32 = (^(v127))
                                                                                    _ = v145
                                                                                    _ = v146
                                                                                } else {
                                                                                    v127 -= v127
                                                                                }
                                                                                v147 := GenS8{F0: (v127 ^ 1), F1: v2}
                                                                                _ = v147
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch v4 {
                                                                        case "\n\t":
                                                                            switch v4 {
                                                                            case "":
                                                                                delete(v3, v4)
                                                                            case "\n\t":
                                                                                var v148 float64 = ((-2.25 - -2.25) / float64(float32(float64(-2.25))))
                                                                                _ = v148
                                                                            case "hello":
                                                                                _ = len(v3)
                                                                            }
                                                                        case "hello":
                                                                            var v149 int16 = -1
                                                                            _ = v149
                                                                        }
                                                                    }
                                                                    v119 = 111
                                                                    v2 = (v2 * (v2 ^ 42))
                                                                }
                                                                var v150 uint = 0
                                                                switch (fi117 & (v126 - fi117)) {
                                                                case 42:
                                                                    switch fi79 {
                                                                    case 127:
                                                                        v151 := GenS19{F0: 0}
                                                                        _ = v151
                                                                    case -9223372036854775808:
                                                                        if (fi79 != (v126 / (((v126 << ((18446744073709551615) & 63))) | 1))) {
                                                                            v127 += 65535
                                                                        }
                                                                    case 4294967295:
                                                                        fmt.Println(fi117, v127)
                                                                    default:
                                                                        v152 := GenS20{F0: fi79, F1: 373}
                                                                        _ = v152
                                                                    }
                                                                case 2147483647:
                                                                    var v153 int64 = int64(v126)
                                                                    _ = v153
                                                                default:
                                                                    v150 = (v150 | (^(v150)))
                                                                }
                                                                _ = v150
                                                            } else {
                                                                if (v4 == ((("a" + v4) + (v4 + (v4 + (v4 + "\n\t")))) + v4)) {
                                                                    var v154 float64 = 1.5
                                                                    _ = v154
                                                                }
                                                                if false {
                                                                    var v155 uint = 65535
                                                                    v156 := GenS17{F0: v127}
                                                                    _ = v155
                                                                    _ = v156
                                                                } else {
                                                                    if (!(false || (!(((!(127 > 75)) || (0.1 < 0.0)) || (130 <= 2))))) {
                                                                        v2 += v2
                                                                        var v157 float32 = 0.1
                                                                        var v158 int32 = 1
                                                                        _ = v157
                                                                        _ = v158
                                                                    }
                                                                }
                                                                var v159 int32 = 203
                                                                if false {
                                                                    var v160 int16 = 1000
                                                                    v161 := GenS21{F0: uint16(v119), F1: v127, F2: -1}
                                                                    v4 = ("a" + "hello")
                                                                    v162 := GenS3{F0: 42, F1: (v127 * v127), F2: 1}
                                                                    _ = v160
                                                                    _ = v161
                                                                    _ = v162
                                                                } else {
                                                                    v163 := GenS22{F0: v4}
                                                                    v164 := GenS21{F0: 255, F1: (^(v127)), F2: 9223372036854775807}
                                                                    _ = v163
                                                                    _ = v164
                                                                }
                                                                if ((1e10 + -2.25) < 0.1) {
                                                                    fmt.Println(v119)
                                                                    v4 = ""
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                                fmt.Println(v126, v159, v119)
                                                                _ = v159
                                                            }
                                                        }
                                                    } else {
                                                        if (!false) {
                                                            v165 := GenS21{F0: (v2 >> ((386) & 63)), F1: uint32(v119), F2: 9223372036854775807}
                                                            v4 = ("" + ("" + v4))
                                                            fmt.Println(fi79, fi117)
                                                            _ = v165
                                                        }
                                                        if (v119 >= 2) {
                                                            var v166 uint64 = 42
                                                            _ = v166
                                                        } else {
                                                            if (true && true) {
                                                                if (513 != 209) {
                                                                    _ = len(v1)
                                                                    var v167 int32 = 701
                                                                    v168 := GenS23{F0: false, F1: 644}
                                                                    _ = v167
                                                                    _ = v168
                                                                }
                                                                v4 = v4
                                                            }
                                                            switch v4 {
                                                            case "hello":
                                                                var v169 uint8 = (v119 ^ (v119 + uint8(v2)))
                                                                _ = v169
                                                            default:
                                                                fmt.Println(fi79, v4, v127)
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = (v127 % ((v127) | 1))
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v3[v4] = (1e10 >= (-2.25 * 0.1))
                                                }
                                            }
                                        } else {
                                            if false {
                                                var v170 float32 = float32(float64((-2.25 / 1.0)))
                                                fmt.Println(fi117, v170, v2)
                                                v171 := GenS20{F0: (v126 % ((127) | 1)), F1: 0}
                                                _ = v170
                                                _ = v171
                                            } else {
                                                switch 1 {
                                                case 128:
                                                    var v172 int8 = 42
                                                    _ = v172
                                                }
                                                if ((0.0 / 1.0) <= (1e10 - 1.5)) {
                                                    fmt.Println(v127)
                                                } else {
                                                    v173 := GenS23{F0: false, F1: 329}
                                                    v2 = 32767
                                                    _ = v173
                                                }
                                            }
                                            fmt.Println(v4, v2)
                                            v126 += v126
                                            v2 = v2
                                            fmt.Println(v126)
                                            if ((!false) && false) {
                                                if ((v119 - v119) < (v119 << ((52) & 63))) {
                                                    v126 += fi79
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = len(v1)
                                            } else {
                                                fmt.Println(v127, v126)
                                                var v174 int16 = 0
                                                _ = v174
                                            }
                                        }
                                        var v175 int64 = -853
                                        _ = v127
                                        _ = v175
                                    } else {
                                        if (false && false) {
                                            v2 = (v2 << ((1) & 63))
                                            v176 := GenS21{F0: 65535, F1: 1, F2: 855}
                                            _ = v176
                                        } else {
                                            var v177 int64 = 0
                                            v2 = 32767
                                            _ = len(v1)
                                            _ = v177
                                        }
                                        fmt.Println(v126)
                                    }
                                    _ = v126
                                }
                            default:
                                var v178 int8 = 41
                                v179 := GenS24{F0: v119, F1: (v178 % ((v178) | 1))}
                                _ = len(v1)
                                _ = v178
                                _ = v179
                            }
                        }
                        v180 := GenS25{F0: 0}
                        _ = fi117
                        _ = v119
                        _ = v180
                    }
                    switch (v4 + (v4 + ((v4 + ("foo bar" + v4)) + "\n\t"))) {
                    case "foo bar":
                        v2 = v2
                    default:
                        switch 9223372036854775807 {
                        case 1:
                            v181 := GenS3{F0: 42, F1: 0, F2: 1}
                            _ = v181
                        case 4294967295:
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                        default:
                            v4 = (v4 + (("\n\t" + (v4 + ((v4 + ("" + (("hello" + v4) + ""))) + ((v4 + v4) + (v4 + ("\n\t" + (v4 + v4))))))) + "foo bar"))
                        }
                    }
                } else {
                    if (645 < 150) {
                        for fi182 := 0; fi182 < 2; fi182++ {
                            v4 = v4
                            if (!true) {
                                if (false || false) {
                                    var v183 uint32 = 970
                                    v184 := GenS26{F0: 884, F1: fi79, F2: 32}
                                    v185 := GenS6{F0: v183, F1: (v2 << ((870) & 63)), F2: 1}
                                    _ = v183
                                    _ = v184
                                    _ = v185
                                }
                            } else {
                                v2 -= (^(v2))
                                var v186 float32 = 1e10
                                v2 -= (v2 >> ((1) & 63))
                                _ = v186
                            }
                            _ = fi182
                        }
                        if (!(-1 >= -9223372036854775808)) {
                            for fi187 := 0; fi187 < 19; fi187++ {
                                v4 = ((v4 + (v4 + ("hello" + v4))) + "a")
                                _ = fi187
                            }
                            if (((v2 / ((v2) | 1)) != 340) && true) {
                                v1 = append(v1, 297)
                                var v188 string = ""
                                _ = v188
                            }
                            fmt.Println(v4)
                        } else {
                            switch (fi79 ^ (fi79 * 9223372036854775807)) {
                            case -2147483648:
                                var v189 uint64 = 4294967295
                                _ = v189
                            }
                            for fi190 := 0; fi190 < 6; fi190++ {
                                v4 = v4
                                _ = len(v1)
                                v4 = v4
                                _ = fi190
                            }
                            v2 -= v2
                        }
                    } else {
                        v2 += (^(v2))
                    }
                    v3[(v4 + "a")] = (65535 < v2)
                    v4 = v4
                }
                switch (v4 + "") {
                case "a":
                    var v191 bool = (396 > v2)
                    _ = v191
                }
            }
            var v192 float64 = 1.5
            _ = fi79
            _ = v192
        }
        fmt.Println(v2, v2)
        fmt.Println(v4, v4)
    }
    var v193 int64 = 42
    v2 += (v2 - (v2 | 65535))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v193:", v193)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v193
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
