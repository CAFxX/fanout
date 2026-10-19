package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int32; F1 string; F2 uint }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 int64; F1 int8; F2 int32 }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string; F1 int8; F2 uint64 }
type GenS14 struct { F0 uint16 }

func main() {
    var v1 bool = true
    var v2 int = 9223372036854775807
    var v3 int = v2
    var v4 string = (("a" + (("" + ("\n\t" + "foo bar")) + ("hello" + (("a" + "hello") + ("a" + "foo bar"))))) + ("foo bar" + "hello"))
    v5 := []uint32{4294967295, 2147483647, 331}
    var v6 string = "\n\t"
    switch (("\n\t" + "\n\t") + ((("a" + "hello") + (("a" + "\n\t") + "foo bar")) + "hello")) {
    case "\n\t":
        v3 -= -104
    case "a":
        if (0 == -783) {
            fmt.Println(v3, v1, v4)
            if (!(32767 >= 0)) {
                var v7 int8 = 0
                _ = v7
            } else {
                v8 := GenS1{F0: int32(v2), F1: 747, F2: v6}
                v1 = (!v1)
                _ = v8
            }
        }
    }
    if v1 {
        v5 = append(v5, 42)
        if (2 != 0) {
            fmt.Println(v4)
            v3 += (v2 * v2)
            var v9 uint32 = 1
            _ = v9
        } else {
            v3 = v3
            for fi10 := 0; fi10 < 14; fi10++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch "hello" {
                case "hello":
                    var v11 uint8 = 108
                    _ = v11
                case "\n\t":
                    v12 := GenS2{F0: 0, F1: 42}
                    _ = v12
                default:
                    var v13 uint8 = 1
                    _ = v13
                }
                _ = fi10
            }
        }
        v14 := GenS2{F0: 880, F1: -722}
        _ = v14
    } else {
        var v15 int32 = 1
        if v1 {
            if false {
                _ = len(v5)
                var v16 uint32 = 4294967295
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v16
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint = 42
                _ = fi17
                _ = v18
            }
            _ = len(v5)
        } else {
            if len(v5) > 0 {
                v5[0] = 65535
            }
            v19 := GenS3{F0: v15, F1: ((v4 + "a") + "a"), F2: 828}
            switch v4 {
            case "":
                var v20 string = v6
                _ = v20
            case "foo bar":
                var v21 bool = (-9223372036854775808 <= 0)
                _ = v21
            default:
                for fi22 := 0; fi22 < 3; fi22++ {
                    fmt.Println(fi22)
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 uint64 = 2
                        switch ("a" + (((((v4 + (v6 + ((("a" + "") + ((v6 + "hello") + "a")) + (("" + ((v6 + v4) + (v4 + (v6 + "foo bar")))) + (("\n\t" + v6) + ((("foo bar" + v6) + "") + (v4 + ""))))))) + v4) + v4) + "\n\t") + v4)) {
                        case "foo bar":
                            v24 = v24
                        }
                        var v25 uint32 = 963
                        var v26 int = (fi23 & fi22)
                        v5 = append(v5, 2)
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = fi22
                }
            }
            _ = v19
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 uint = 18446744073709551615
            _ = fi27
            _ = v28
        }
        _ = v15
    }
    var v29 int16 = -691
    for fi30 := 0; fi30 < 15; fi30++ {
        v3 = fi30
        if len(v5) > 0 {
            v5[0] = 419
        }
        var v31 uint = 65535
        _ = fi30
        _ = v31
    }
    var v32 uint32 = 926
    var v33 int32 = -2147483648
    for fi34 := 0; fi34 < 20; fi34++ {
        v33 -= v33
        if (!(true || (uint8(v32) > 127))) {
            if (!true) {
                var v35 float64 = 1e10
                var v36 uint = 0
                for fi37 := 0; fi37 < 6; fi37++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v38 string = "\n\t"
                    _ = fi37
                    _ = v38
                }
                _ = v35
                _ = v36
            }
        }
        _ = fi34
    }
    if false {
        if ((((^(v2)) != (v3 + v2)) && (-1 <= 174)) || false) {
            v39 := GenS1{F0: (v33 / (((v33 << ((1) & 63))) | 1)), F1: -9223372036854775808, F2: v6}
            switch (v3 >> ((65535) & 63)) {
            case -838:
                var v40 int16 = v29
                _ = v40
            case 127:
                for fi41 := 0; fi41 < 5; fi41++ {
                    if (!v1) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v42 := GenS2{F0: 4294967295, F1: (v33 + 455)}
                        v29 -= 42
                        _ = v42
                    }
                    _ = fi41
                }
            case 9223372036854775807:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v3 -= v3
            _ = v39
        } else {
            var v43 bool = false
            var v44 float32 = 1.5
            for fi45 := 0; fi45 < 1; fi45++ {
                if ("" == v4) {
                    fmt.Println(v44, v43)
                    v46 := GenS4{F0: (!(("" + "hello") != (v6 + "hello"))), F1: 675}
                    if (true && v1) {
                        if false {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v47 int32 = (^(v33))
                            if v43 {
                                for fi48 := 0; fi48 < 4; fi48++ {
                                    v4 = v4
                                    var v49 uint16 = 255
                                    _ = fi48
                                    _ = v49
                                }
                            } else {
                                v5 = append(v5, v32)
                                v29 = (v29 + (v29 >> ((42) & 63)))
                            }
                            _ = v47
                        } else {
                            v5 = append(v5, v32)
                            v43 = v43
                        }
                        v43 = v1
                        var v50 uint = 774
                        _ = v50
                    }
                    _ = v46
                }
                _ = fi45
            }
            _ = v43
            _ = v44
        }
        fmt.Println(v1, v32)
        for fi51 := 0; fi51 < 20; fi51++ {
            for fi52 := 0; fi52 < 11; fi52++ {
                var v53 int = fi52
                if (!false) {
                    if ("foo bar" > (v6 + ((((((("" + "a") + v6) + v6) + ("" + v6)) + "\n\t") + "foo bar") + (v6 + v6)))) {
                        v54 := GenS5{F0: 9223372036854775807, F1: 102}
                        v55 := GenS2{F0: 42, F1: -1}
                        fmt.Println(v1)
                        _ = v54
                        _ = v55
                    } else {
                        var v56 uint8 = 127
                        v53 = (^(fi51))
                        if v1 {
                            fmt.Println(v1)
                            if v1 {
                                fmt.Println(v3, v29, v32)
                                v33 -= -2147483648
                                v2 = (v53 % ((v2) | 1))
                            } else {
                                switch fi52 {
                                case 4294967295:
                                    switch (v2 * (-(v53))) {
                                    case 32767:
                                        fmt.Println(fi52)
                                    default:
                                        v56 -= v56
                                    }
                                }
                                var v57 int64 = -9223372036854775808
                                fmt.Println(v1, v4)
                                _ = v57
                            }
                            v58 := GenS4{F0: v1, F1: 4294967295}
                            _ = v58
                        } else {
                            v59 := GenS3{F0: (v33 - (v33 + v33)), F1: (v6 + (v4 + (v4 + v4))), F2: 837}
                            switch 32768 {
                            case -2147483648:
                                v60 := GenS6{F0: (v32 & (v32 >> ((835) & 63)))}
                                v61 := GenS2{F0: uint(v32), F1: -2147483648}
                                _ = v60
                                _ = v61
                            case -1:
                                v62 := GenS7{F0: (!v1)}
                                _ = v62
                            default:
                                v63 := GenS8{F0: -9223372036854775808, F1: 969}
                                _ = v63
                            }
                            _ = v59
                        }
                        switch (-(fi52)) {
                        case 9223372036854775807:
                            v29 = -32768
                        default:
                            if v1 {
                                _ = len(v5)
                            } else {
                                v5 = append(v5, (v32 / ((42) | 1)))
                            }
                        }
                        _ = v56
                    }
                    v2 = (v2 * v3)
                    v33 += v33
                } else {
                    fmt.Println(v4, v3, v6)
                    if false {
                        v64 := GenS9{F0: v29, F1: 501}
                        var v65 uint32 = (v32 - v32)
                        _ = v64
                        _ = v65
                    } else {
                        v32 += v32
                    }
                }
                _ = fi52
                _ = v53
            }
            fmt.Println(v32, v1, v6)
            _ = fi51
        }
    }
    var v66 float64 = 3.14159
    fmt.Println(v6)
    if (((3.14159 - ((((1e10 - 1e10) * float32(v66)) + 0.1) / 1.5)) < 3.14159) && false) {
        if ("hello" == (v6 + "")) {
            v33 -= 2147483647
            var v67 int64 = -1
            _ = v67
        } else {
            if true {
                var v68 int64 = 1
                v33 += v33
                _ = v68
            } else {
                switch v4 {
                case "a":
                    if false {
                        if v1 {
                            v29 += (-(v29))
                            switch v2 {
                            case 32768:
                                fmt.Println(v4)
                            case -9223372036854775808:
                                var v69 bool = (!(32767 >= 5))
                                _ = v69
                            }
                        }
                    } else {
                        switch (v6 + (v6 + v4)) {
                        case "\n\t":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "a":
                            var v70 int16 = (v29 / (((v29 & v29)) | 1))
                            _ = v70
                        default:
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                for fi71 := 0; fi71 < 13; fi71++ {
                                    var v72 uint32 = v32
                                    fmt.Println(v3)
                                    v6 = "\n\t"
                                    _ = fi71
                                    _ = v72
                                }
                                for fi73 := 0; fi73 < 19; fi73++ {
                                    fmt.Println(v4)
                                    _ = fi73
                                }
                            }
                        }
                        v74 := GenS3{F0: (v33 * (v33 >> ((0) & 63))), F1: v6, F2: 2}
                        _ = v74
                    }
                case "hello":
                    v2 = v3
                case "foo bar":
                    if v1 {
                        switch "hello" {
                        case "hello":
                            v6 = v6
                        }
                        var v75 float32 = -2.25
                        if (1 == 2) {
                            var v76 uint8 = 128
                            for fi77 := 0; fi77 < 19; fi77++ {
                                v6 = ""
                                if (false || (187 != 9223372036854775807)) {
                                    var v78 uint16 = 255
                                    if ((53 <= 100) && (((v75 - v75) > 0.0) || (true && false))) {
                                        for fi79 := 0; fi79 < 4; fi79++ {
                                            v33 += 1
                                            _ = fi79
                                        }
                                        v80 := GenS7{F0: (true && v1)}
                                        _ = v80
                                    } else {
                                        if (4294967295 == (^(v32))) {
                                            switch v6 {
                                            case "\n\t":
                                                var v81 uint8 = v76
                                                _ = v81
                                            case "hello":
                                                for fi82 := 0; fi82 < 5; fi82++ {
                                                    if v1 {
                                                        var v83 int8 = 114
                                                        var v84 uint8 = uint8(v32)
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    if (v1 && v1) {
                                                        v32 = (v32 & v32)
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v78 = (v78 / ((255) | 1))
                                                        var v85 int16 = (v29 + v29)
                                                        if false {
                                                            switch (((((v6 + v6) + "a") + (v6 + "hello")) + v4) + v4) {
                                                            case "hello":
                                                                v75 += 0.0
                                                            case "foo bar":
                                                                var v86 uint = 255
                                                                _ = v86
                                                            case "":
                                                                var v87 uint64 = 0
                                                                _ = v87
                                                            }
                                                            v88 := GenS10{F0: -827, F1: 0, F2: v33}
                                                            v66 = float64(0.1)
                                                            v89 := GenS11{F0: 421, F1: 0}
                                                            v85 -= 1
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        _ = v85
                                                    }
                                                    if (88 < 1) {
                                                        v90 := GenS11{F0: 0, F1: 42}
                                                        _ = v90
                                                    }
                                                    var v91 int = int(v32)
                                                    v32 = (v32 ^ 0)
                                                    fmt.Println(v1, v75, v76)
                                                    _ = fi82
                                                    _ = v91
                                                }
                                            default:
                                                v4 = ((v6 + (v6 + (v4 + "a"))) + (v6 + "hello"))
                                            }
                                            switch (fi77 % ((-808) | 1)) {
                                            case -637:
                                                v76 -= (^(v76))
                                            case 4294967295:
                                                v5 = append(v5, (v32 / (((^(v32))) | 1)))
                                            default:
                                                if v1 {
                                                    fmt.Println(v1, v78, fi77)
                                                    var v92 uint16 = (v78 >> ((65535) & 63))
                                                    _ = v92
                                                }
                                            }
                                        } else {
                                            fmt.Println(v3)
                                        }
                                    }
                                    var v93 int8 = -67
                                    _ = v78
                                    _ = v93
                                }
                                v5 = append(v5, (^(v32)))
                                _ = fi77
                            }
                            _ = v76
                        } else {
                            v4 = v6
                            var v94 float32 = v75
                            for fi95 := 0; fi95 < 8; fi95++ {
                                if ((v33 <= v33) || v1) {
                                    var v96 uint64 = 2
                                    _ = v96
                                }
                                _ = fi95
                            }
                            _ = v94
                        }
                        _ = v75
                    }
                default:
                    for fi97 := 0; fi97 < 13; fi97++ {
                        fmt.Println(v2, v3)
                        var v98 uint = 18446744073709551615
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi97
                        _ = v98
                    }
                }
                switch 9223372036854775807 {
                case -9223372036854775808:
                    if v1 {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        for fi99 := 0; fi99 < 12; fi99++ {
                            if true {
                                for fi100 := 0; fi100 < 13; fi100++ {
                                    fmt.Println(v6, v29)
                                    _ = fi100
                                }
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v2)
                                case "a":
                                    v101 := GenS12{F0: ((v4 + (v6 + "foo bar")) + "hello")}
                                    _ = v101
                                default:
                                    switch fi99 {
                                    case 32768:
                                        var v102 uint32 = v32
                                        _ = v102
                                    }
                                }
                            } else {
                                v2 += fi99
                                for fi103 := 0; fi103 < 11; fi103++ {
                                    v66 = v66
                                    if v1 {
                                        v4 = v4
                                        v5 = append(v5, v32)
                                        fmt.Println(v33)
                                    } else {
                                        switch v3 {
                                        case 2147483647:
                                            if (!false) {
                                                if v1 {
                                                    v4 = ""
                                                }
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = (^(v32))
                                                }
                                            }
                                        case -9223372036854775808:
                                            v104 := GenS10{F0: 9223372036854775807, F1: 42, F2: v33}
                                            _ = v104
                                        case 981:
                                            fmt.Println(v1, v6)
                                        }
                                        var v105 uint32 = v32
                                        _ = v105
                                    }
                                    if ((v29 / ((v29) | 1)) >= (-(v29))) {
                                        switch (fi99 << ((1) & 63)) {
                                        case -480:
                                            switch (fi99 * 0) {
                                            case 4294967295:
                                                var v106 uint8 = 42
                                                _ = v106
                                            case 9223372036854775807:
                                                if false {
                                                    var v107 int = 128
                                                    v108 := GenS3{F0: 1, F1: (v4 + "\n\t"), F2: 255}
                                                    _ = v107
                                                    _ = v108
                                                } else {
                                                    v109 := GenS1{F0: (v33 | (v33 ^ (v33 / ((v33) | 1)))), F1: 42, F2: "foo bar"}
                                                    _ = v109
                                                }
                                            default:
                                                if len(v5) > 0 {
                                                    v5[0] = 2147483648
                                                }
                                            }
                                        case 4294967295:
                                            fmt.Println(v1, v32, v33)
                                        case -744:
                                            var v110 bool = v1
                                            _ = v110
                                        }
                                        fmt.Println(fi99, v6, fi103)
                                    } else {
                                        switch v4 {
                                        case "a":
                                            if len(v5) > 0 {
                                                v5[0] = v32
                                            }
                                        case "":
                                            var v111 int = v3
                                            _ = v111
                                        case "\n\t":
                                            var v112 uint64 = 1
                                            _ = v112
                                        }
                                        v5 = append(v5, v32)
                                    }
                                    _ = fi103
                                }
                                v2 = v3
                            }
                            v113 := GenS13{F0: ("a" + ""), F1: -13, F2: 2}
                            _ = fi99
                            _ = v113
                        }
                        v114 := GenS14{F0: 42}
                        _ = v114
                    } else {
                        var v115 uint8 = 38
                        _ = v115
                    }
                default:
                    for fi116 := 0; fi116 < 1; fi116++ {
                        switch v4 {
                        case "":
                            var v117 float64 = float64(float32(v66))
                            _ = v117
                        case "hello":
                            for fi118 := 0; fi118 < 9; fi118++ {
                                if v1 {
                                    v119 := GenS4{F0: v1, F1: 0}
                                    var v120 uint16 = 255
                                    v5 = append(v5, (v32 >> ((42) & 63)))
                                    _ = v119
                                    _ = v120
                                }
                                _ = len(v5)
                                var v121 int8 = -128
                                _ = fi118
                                _ = v121
                            }
                        }
                        v122 := GenS9{F0: v29, F1: -883}
                        _ = fi116
                        _ = v122
                    }
                }
            }
        }
        if (!v1) {
            var v123 int = (v3 * (v3 - 395))
            if (float64((1e10 / 1.0)) > v66) {
                var v124 int64 = 1
                _ = v124
            } else {
                fmt.Println(v32, v4)
                v125 := GenS5{F0: -1, F1: 127}
                _ = v125
            }
            v126 := GenS8{F0: 0, F1: 65535}
            _ = v123
            _ = v126
        }
        fmt.Println(v3)
    }
    v32 += (^(v32))
    v32 = 979
    var v127 uint = 278
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v127
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
