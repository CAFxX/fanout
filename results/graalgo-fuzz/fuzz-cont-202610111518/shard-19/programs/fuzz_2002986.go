package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int64; F2 uint16 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int; F1 uint; F2 uint }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 int; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS11 struct { F0 uint16; F1 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int64; F1 int16; F2 int32 }
type GenS15 struct { F0 uint8; F1 uint32 }
type GenS16 struct { F0 bool; F1 int64; F2 uint32 }
type GenS17 struct { F0 string }

func main() {
    v1 := []int{32768, 345, 1, 1}
    var v2 float64 = (3.14159 * -2.25)
    var v3 bool = (!(false && (0 > 9223372036854775807)))
    v4 := make(map[string]int)
    v2 -= v2
    v2 = float64(1.5)
    v5 := GenS1{F0: 48, F1: 746, F2: 1}
    v6 := GenS1{F0: 42, F1: 105, F2: 127}
    var v7 string = "foo bar"
    var v8 uint = 42
    if (0 >= -32768) {
        v9 := GenS1{F0: 65535, F1: -1, F2: 84}
        v3 = false
        _ = v9
    } else {
        var v10 uint16 = 65535
        fmt.Println(v8)
        var v11 string = v7
        _ = v10
        _ = v11
    }
    fmt.Println(v7)
    if (!(-1 < -1)) {
        if v3 {
            fmt.Println(v7)
            v12 := GenS1{F0: 42, F1: 151, F2: 128}
            var v13 uint64 = 18446744073709551615
            _ = v12
            _ = v13
        }
    }
    var v14 string = (v7 + ("hello" + v7))
    if (100 <= 100) {
        v3 = ((1.5 != 1e10) && ((!(((v8 <= (v8 % ((77) | 1))) && v3) && (576 < -32768))) || (!(true && v3))))
        if (0 == -32768) {
            if v3 {
                switch v14 {
                case "foo bar":
                    fmt.Println(v2, v14)
                case "a":
                    v7 = v14
                case "\n\t":
                    switch ((v14 + ("\n\t" + ((("hello" + "foo bar") + v14) + "hello"))) + v14) {
                    case "\n\t":
                        var v15 float32 = 0.0
                        _ = v15
                    case "hello":
                        v3 = (223 >= 2)
                    case "a":
                        var v16 uint64 = 64
                        _ = v16
                    default:
                        v17 := GenS2{F0: 127, F1: v8, F2: 255}
                        _ = v17
                    }
                default:
                    if (((!((false && v3) || ((!(v3 || v3)) && (-657 <= -9223372036854775808)))) || v3) || v3) {
                        v18 := v4[v7]
                        v18 = -857
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v18
                    } else {
                        v2 = ((1e10 + ((v2 / 3.14159) + ((0.1 / v2) / -2.25))) + (v2 / (((v2 * v2) - v2) + (v2 / (v2 - ((v2 + 1e10) / -2.25))))))
                        for fi19 := 0; fi19 < 17; fi19++ {
                            for fi20 := 0; fi20 < 5; fi20++ {
                                fmt.Println(fi20, v14, v8)
                                fmt.Println(v7)
                                _ = fi20
                            }
                            _ = fi19
                        }
                        for fi21 := 0; fi21 < 4; fi21++ {
                            switch v14 {
                            case "foo bar":
                                var v22 float32 = -2.25
                                _ = v22
                            case "\n\t":
                                _ = len(v4)
                            case "":
                                for fi23 := 0; fi23 < 14; fi23++ {
                                    if v3 {
                                        switch v14 {
                                        case "foo bar":
                                            v3 = v3
                                            fmt.Println(v8, v2)
                                        case "hello":
                                            v8 -= 1
                                        case "a":
                                            v3 = ((!((!false) || true)) && (float64(((0.0 + 0.1) * ((1e10 + float32(0.0)) + -2.25))) != v2))
                                        }
                                        fmt.Println(v8)
                                        v24 := GenS3{F0: 145, F1: -180}
                                        v14 = ""
                                        v2 = (v2 * (v2 / (v2 * ((v2 / (v2 + v2)) / (float64(0.1) / (v2 - float64(0.1)))))))
                                        if (2 <= 255) {
                                            v25 := GenS4{F0: 65535}
                                            var v26 uint = 18446744073709551615
                                            _ = v25
                                            _ = v26
                                        }
                                        if v3 {
                                            v4[(v7 + v7)] = fi23
                                            v14 = "\n\t"
                                            delete(v4, "hello")
                                        }
                                        switch (fi23 / ((fi21) | 1)) {
                                        case -9223372036854775808:
                                            _ = len(v1)
                                        case 2147483647:
                                            switch 42 {
                                            case -910:
                                                fmt.Println(v8, fi21)
                                            default:
                                                v27 := GenS5{F0: fi21, F1: -1, F2: 681}
                                                _ = v27
                                            }
                                        }
                                        _ = v24
                                    }
                                    v3 = true
                                    v2 -= (v2 * ((v2 - v2) / v2))
                                    _ = fi23
                                }
                            }
                            for fi28 := 0; fi28 < 14; fi28++ {
                                delete(v4, "a")
                                switch -463 {
                                case 4294967295:
                                    if (fi21 >= fi28) {
                                        v29 := GenS5{F0: (fi21 | fi28), F1: 1, F2: 0}
                                        _ = v29
                                    }
                                case -1:
                                    switch fi21 {
                                    case 887:
                                        if (!false) {
                                            var v30 uint64 = 4294967295
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v30
                                        }
                                    case 2147483647:
                                        var v31 uint64 = 0
                                        _ = v31
                                    case 0:
                                        var v32 bool = v3
                                        _ = v32
                                    default:
                                        var v33 uint64 = 0
                                        _ = v33
                                    }
                                case 127:
                                    var v34 uint64 = 4294967295
                                    _ = v34
                                default:
                                    var v35 float32 = 1.5
                                    _ = v35
                                }
                                if v3 {
                                    switch "" {
                                    case "foo bar":
                                        var v36 uint64 = 0
                                        _ = v36
                                    case "":
                                        _ = len(v4)
                                    case "hello":
                                        if v3 {
                                            var v37 float32 = 3.14159
                                            v38 := GenS5{F0: (fi28 / (((fi21 << ((v8) & 63))) | 1)), F1: 0, F2: 32767}
                                            v7 = "a"
                                            _ = v37
                                            _ = v38
                                        } else {
                                            if false {
                                                var v39 uint64 = 1
                                                _ = v39
                                            }
                                            var v40 uint16 = 255
                                            v41 := GenS3{F0: 197, F1: 42}
                                            _ = v40
                                            _ = v41
                                        }
                                    default:
                                        v42 := GenS6{F0: 1000, F1: 2147483647}
                                        _ = v42
                                    }
                                    v43 := GenS5{F0: (fi21 >> ((255) & 63)), F1: 590, F2: 494}
                                    v3 = ((42 >= 4294967295) && v3)
                                    _ = v43
                                } else {
                                    v7 = (v7 + (((v14 + (v7 + "hello")) + v14) + "\n\t"))
                                }
                                _ = fi28
                            }
                            _ = fi21
                        }
                    }
                }
            } else {
                for fi44 := 0; fi44 < 1; fi44++ {
                    var v45 int16 = 1
                    if v3 {
                        fmt.Println(fi44)
                        if (false || (-828 <= v45)) {
                            if v3 {
                                v2 = (v2 * (v2 / (0.1 / v2)))
                                v46 := GenS7{F0: fi44, F1: v8, F2: v8}
                                fmt.Println(v14)
                                _ = v46
                            } else {
                                _ = len(v1)
                                v47 := GenS5{F0: fi44, F1: -752, F2: 104}
                                for fi48 := 0; fi48 < 11; fi48++ {
                                    fmt.Println(v3, fi48, v2)
                                    v7 = v7
                                    _ = fi48
                                }
                                _ = v47
                            }
                            if (!v3) {
                                fmt.Println(fi44)
                                switch v14 {
                                case "a":
                                    if false {
                                        v49 := GenS8{F0: v8, F1: (v45 ^ v45)}
                                        _ = v49
                                    } else {
                                        var v50 int16 = v45
                                        _ = v50
                                    }
                                default:
                                    var v51 uint64 = 1
                                    _ = v51
                                }
                                var v52 int64 = 423
                                _ = v52
                            }
                        } else {
                            v45 -= -32768
                            var v53 uint64 = 18446744073709551615
                            _ = v53
                        }
                        if len(v1) > 0 {
                            v1[0] = (fi44 - fi44)
                        }
                    } else {
                        fmt.Println(v14)
                        v54 := v4[("\n\t" + "hello")]
                        if (2147483647 != 2147483647) {
                            if v3 {
                                fmt.Println(v8, fi44)
                                var v55 uint64 = 0
                                for fi56 := 0; fi56 < 10; fi56++ {
                                    v57 := GenS3{F0: uint8(v8), F1: -32768}
                                    v54 -= fi44
                                    var v58 int8 = 43
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                }
                                _ = v55
                            }
                        }
                        _ = v54
                    }
                    if v3 {
                        for fi59 := 0; fi59 < 1; fi59++ {
                            v14 = "hello"
                            _ = fi59
                        }
                        _ = len(v4)
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                    }
                    _ = fi44
                    _ = v45
                }
                v60 := GenS7{F0: 502, F1: v8, F2: v8}
                _ = v60
            }
        } else {
            v3 = (!v3)
            v8 += (v8 << (((^(v8))) & 63))
        }
        if (-2.25 > v2) {
            v61 := GenS9{F0: 32768, F1: 42}
            if (117 < 2) {
                fmt.Println(v14, v3)
                _ = len(v4)
            }
            _ = v61
        } else {
            v62 := GenS7{F0: -484, F1: 523, F2: (v8 | v8)}
            if true {
                for fi63 := 0; fi63 < 1; fi63++ {
                    v3 = (32767 == (fi63 | fi63))
                    _ = fi63
                }
            } else {
                var v64 uint32 = uint32(v8)
                _ = v64
            }
            _ = v62
        }
    } else {
        fmt.Println(v7, v2, v14)
    }
    switch -259 {
    case -2147483648:
        if v3 {
            var v65 uint8 = 3
            if v3 {
                v3 = ((0 != (^(v8))) && v3)
            }
            _ = v65
        } else {
            var v66 uint8 = 2
            v14 = (v7 + v14)
            _ = v66
        }
    case 127:
        v14 = v7
    case 32767:
        v67 := GenS2{F0: 9223372036854775807, F1: v8, F2: 255}
        _ = v67
    default:
        v68 := v4[v7]
        _ = v68
    }
    if v3 {
        var v69 int16 = -695
        for fi70 := 0; fi70 < 17; fi70++ {
            v71 := v4[v7]
            v72 := GenS7{F0: fi70, F1: v8, F2: v8}
            v73 := GenS7{F0: (fi70 + int(v8)), F1: v8, F2: 18446744073709551615}
            _ = fi70
            _ = v71
            _ = v72
            _ = v73
        }
        delete(v4, v7)
        _ = v69
    } else {
        switch 0 {
        case -2147483648:
            if true {
                fmt.Println(v3, v14)
                for fi74 := 0; fi74 < 14; fi74++ {
                    if false {
                        fmt.Println(v7)
                    }
                    if v3 {
                        fmt.Println(v14)
                        var v75 uint32 = 4294967295
                        _ = v75
                    } else {
                        v14 = ("foo bar" + v14)
                    }
                    switch v14 {
                    case "\n\t":
                        var v76 uint8 = 161
                        _ = v76
                    case "hello":
                        v14 = ("\n\t" + "foo bar")
                    default:
                        v77 := GenS5{F0: 32767, F1: -1, F2: 377}
                        _ = v77
                    }
                    _ = fi74
                }
                fmt.Println(v14, v14)
            } else {
                var v78 int16 = 23
                if v3 {
                    if v3 {
                        if v3 {
                            v2 += v2
                            for fi79 := 0; fi79 < 19; fi79++ {
                                v80 := GenS10{F0: (1 == 42), F1: 298, F2: 42}
                                if (false || v3) {
                                    for fi81 := 0; fi81 < 16; fi81++ {
                                        v82 := GenS8{F0: (v8 >> ((v8) & 63)), F1: 698}
                                        fmt.Println(v78, v14, v14)
                                        switch (v14 + v14) {
                                        case "hello":
                                            v2 = ((float64(float32((v2 - v2))) - v2) * 1e10)
                                        case "foo bar":
                                            if (v78 != v78) {
                                                v14 = ((v7 + v7) + v14)
                                                if ((v8 % ((2) | 1)) < v8) {
                                                    fmt.Println(v78)
                                                }
                                                switch v7 {
                                                case "":
                                                    v1 = append(v1, fi81)
                                                default:
                                                    v1 = append(v1, (fi81 << ((42) & 63)))
                                                }
                                                _ = len(v1)
                                                switch v14 {
                                                case "foo bar":
                                                    v14 = v14
                                                case "":
                                                    if true {
                                                        v8 += (v8 / ((v8) | 1))
                                                        var v83 uint16 = 418
                                                        _ = v83
                                                    }
                                                case "a":
                                                    var v84 uint16 = 2
                                                    var v85 int64 = 9223372036854775807
                                                    v4[(v7 + v14)] = 497
                                                    if v3 {
                                                        v8 = (^(v8))
                                                        var v86 int64 = v85
                                                        switch (fi81 | fi81) {
                                                        case 2147483647:
                                                            v87 := GenS9{F0: fi79, F1: 255}
                                                            _ = v87
                                                        case 42:
                                                            v7 = v7
                                                            if len(v1) > 0 {
                                                                v1[0] = int(v84)
                                                            }
                                                            v88 := GenS11{F0: 42, F1: (v8 & 0)}
                                                            _ = v88
                                                        case 4294967295:
                                                            if len(v1) > 0 {
                                                                v1[0] = -1
                                                            }
                                                        default:
                                                            var v89 uint64 = 1
                                                            _ = v89
                                                        }
                                                        var v90 float32 = (0.0 * (3.14159 + -2.25))
                                                        _ = v86
                                                        _ = v90
                                                    }
                                                    if (v78 > 1000) {
                                                        v91 := GenS12{F0: 100}
                                                        switch v14 {
                                                        case "\n\t":
                                                            switch v14 {
                                                            case "hello":
                                                                var v92 uint = 2
                                                                _ = v92
                                                            case "\n\t":
                                                                v93 := GenS1{F0: 1, F1: v78, F2: 128}
                                                                _ = v93
                                                            }
                                                        case "foo bar":
                                                            v14 = ("hello" + (v7 + v14))
                                                        case "a":
                                                            var v94 int8 = -128
                                                            _ = v94
                                                        }
                                                        var v95 uint = v8
                                                        _ = v91
                                                        _ = v95
                                                    }
                                                    switch ((v7 + ((v14 + "hello") + ("" + ("foo bar" + v14)))) + "\n\t") {
                                                    case "\n\t":
                                                        if (!true) {
                                                            var v96 int = fi79
                                                            v97 := GenS13{F0: 0}
                                                            v1 = append(v1, (fi79 / (((v96 & (v96 / ((v96) | 1)))) | 1)))
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    case "":
                                                        if ((fi81 - (fi81 << ((v8) & 63))) < 32767) {
                                                            v98 := GenS12{F0: 57}
                                                            v84 -= 32767
                                                            if false {
                                                                fmt.Println(v2)
                                                                v84 = 405
                                                                v3 = (-1 >= -1)
                                                            } else {
                                                                var v99 uint8 = 186
                                                                var v100 uint = (v8 | (v8 / ((v8) | 1)))
                                                                v78 = (v78 * (v78 * 1000))
                                                                _ = v100
                                                                _ = v99
                                                            }
                                                            switch fi79 {
                                                            case 127:
                                                                fmt.Println(v7, v14)
                                                            case 9223372036854775807:
                                                                v14 = ((v7 + ("" + (v7 + ("hello" + (v14 + "a"))))) + v7)
                                                            default:
                                                                fmt.Println(v2)
                                                            }
                                                            fmt.Println(v2, v78, v8)
                                                            _ = v98
                                                        } else {
                                                            var v101 uint32 = 717
                                                            delete(v4, v14)
                                                            v85 -= (v85 & int64(v101))
                                                            _ = len(v4)
                                                            _ = v101
                                                        }
                                                    case "hello":
                                                        fmt.Println(v78)
                                                    }
                                                    _ = v84
                                                    _ = v85
                                                default:
                                                    switch 4294967295 {
                                                    case 32767:
                                                        v7 = v14
                                                    case 42:
                                                        v102 := GenS14{F0: -9223372036854775808, F1: -1, F2: 395}
                                                        _ = v102
                                                    case -801:
                                                        v2 += 0.1
                                                    default:
                                                        fmt.Println(fi79, v78)
                                                    }
                                                }
                                            } else {
                                                switch v7 {
                                                case "\n\t":
                                                    var v103 int16 = v78
                                                    _ = v103
                                                case "":
                                                    var v104 float64 = ((-2.25 - v2) + 1e10)
                                                    _ = v104
                                                default:
                                                    if v3 {
                                                        fmt.Println(v8, fi81, fi79)
                                                    } else {
                                                        var v105 int = 0
                                                        switch 9223372036854775807 {
                                                        case -1:
                                                            if (!v3) {
                                                                v106 := GenS7{F0: fi81, F1: v8, F2: (v8 ^ 18446744073709551615)}
                                                                v107 := GenS4{F0: 4294967295}
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                        case -2147483648:
                                                            switch (v14 + v14) {
                                                            case "foo bar":
                                                                var v108 int32 = -1
                                                                _ = v108
                                                            case "hello":
                                                                var v109 int32 = 42
                                                                _ = v109
                                                            case "\n\t":
                                                                v110 := GenS15{F0: 128, F1: 1}
                                                                _ = v110
                                                            }
                                                        default:
                                                            var v111 uint64 = 18446744073709551615
                                                            _ = v111
                                                        }
                                                        var v112 float64 = (float64(0.0) - ((((3.14159 / 3.14159) / v2) - 1e10) / 3.14159))
                                                        _ = v105
                                                        _ = v112
                                                    }
                                                }
                                                if v3 {
                                                    if (127 != 0) {
                                                        var v113 int32 = 278
                                                        _ = v113
                                                    } else {
                                                        v2 -= v2
                                                        delete(v4, ((v14 + "") + v14))
                                                        v114 := GenS3{F0: 0, F1: 1000}
                                                        _ = v114
                                                    }
                                                    fmt.Println(v78)
                                                    v115 := GenS2{F0: 0, F1: (v8 * 1), F2: 42}
                                                    var v116 float64 = (v2 / v2)
                                                    fmt.Println(v78, fi79)
                                                    _ = v115
                                                    _ = v116
                                                }
                                                fmt.Println(v8, v3, v3)
                                                _ = len(v1)
                                            }
                                        case "a":
                                            v117 := GenS16{F0: v3, F1: 1, F2: 2}
                                            v4[v14] = fi81
                                            v2 = (((((v2 / (((0.0 - v2) / v2) / v2)) * (1.5 - 0.0)) - float64((1e10 - float32(3.14159)))) * -2.25) + (float64(float32(-2.25)) - ((v2 + v2) * -2.25)))
                                            _ = v117
                                        default:
                                            var v118 uint8 = 2
                                            _ = v118
                                        }
                                        _ = fi81
                                        _ = v82
                                    }
                                    v14 = v14
                                    v4[""] = fi79
                                } else {
                                    if true {
                                        var v119 uint32 = 2
                                        for fi120 := 0; fi120 < 15; fi120++ {
                                            if (!false) {
                                                v3 = v3
                                                if (v3 && (-47 < -128)) {
                                                    fmt.Println(v119)
                                                }
                                            } else {
                                                _ = len(v4)
                                                var v121 uint16 = 42
                                                _ = v121
                                            }
                                            _ = fi120
                                        }
                                        _ = v119
                                    }
                                    v122 := GenS11{F0: 255, F1: (v8 * v8)}
                                    switch "\n\t" {
                                    case "a":
                                        switch v7 {
                                        case "foo bar":
                                            v78 -= (v78 / ((v78) | 1))
                                        case "\n\t":
                                            var v123 uint16 = 668
                                            _ = v123
                                        }
                                    case "foo bar":
                                        fmt.Println(v8, v3, v2)
                                    default:
                                        v78 += (^(v78))
                                    }
                                    _ = v122
                                }
                                _ = fi79
                                _ = v80
                            }
                        } else {
                            v8 = 810
                            v8 = (v8 * (v8 << (((v8 >> (((v8 / ((4294967295) | 1))) & 63))) & 63)))
                            var v124 uint64 = 12
                            _ = v124
                        }
                    } else {
                        var v125 string = ""
                        _ = v125
                    }
                    switch v14 {
                    case "":
                        v3 = (68 < 0)
                    case "a":
                        v2 = (v2 * (v2 - v2))
                    }
                    for fi126 := 0; fi126 < 18; fi126++ {
                        switch (fi126 ^ fi126) {
                        case 4294967295:
                            for fi127 := 0; fi127 < 1; fi127++ {
                                if v3 {
                                    switch (-(fi127)) {
                                    case 149:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v128 int8 = 108
                                        _ = v128
                                    default:
                                        v8 -= v8
                                    }
                                    var v129 uint8 = 115
                                    _ = v129
                                }
                                if v3 {
                                    delete(v4, v14)
                                } else {
                                    v130 := GenS17{F0: v14}
                                    _ = v130
                                }
                                _ = fi127
                            }
                        case 127:
                            v8 -= (v8 / (((^(v8))) | 1))
                        case -241:
                            for fi131 := 0; fi131 < 3; fi131++ {
                                fmt.Println(v2, v7, fi131)
                                _ = fi131
                            }
                        default:
                            v132 := GenS5{F0: (fi126 | (fi126 - fi126)), F1: 42, F2: 32767}
                            _ = v132
                        }
                        _ = fi126
                    }
                } else {
                    var v133 float32 = (1e10 + (0.0 * -2.25))
                    _ = v133
                }
                _ = v78
            }
        case 32768:
            if false {
                fmt.Println(v2)
                if len(v1) > 0 {
                    v1[0] = -139
                }
            }
        }
    }
    v134 := GenS3{F0: 161, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v134
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
