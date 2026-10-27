package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 int16; F1 int64; F2 string }
type GenS9 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint8 = 0
    var v2 int8 = -128
    var v3 float32 = (1e10 + float32(3.14159))
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]uint32)
    v6 := GenS1{F0: v1, F1: v1, F2: "a"}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := v4[640]
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 bool = (v3 != 0.1)
        _ = fi9
        _ = v10
    }
    var v11 int8 = v2
    v12 := GenS1{F0: v1, F1: v1, F2: "hello"}
    fmt.Println(v1, v11)
    v13 := []int{4294967295, 87, 0}
    v2 = 127
    switch 32768 {
    case 0:
        if true {
            for fi14 := 0; fi14 < 11; fi14++ {
                v4[969] = (!true)
                var v15 int = (-(fi14))
                fmt.Println(v3)
                _ = fi14
                _ = v15
            }
            var v16 uint32 = 2147483648
            v3 -= (((v3 - float32(1.5)) / (0.0 - 1e10)) / float32((((1.5 / 1.5) * -2.25) - (1.5 / (0.0 + 3.14159)))))
            _ = v16
        }
    }
    if (1 != -204) {
        switch "a" {
        case "\n\t":
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int8 = v2
                var v19 uint64 = 42
                _ = fi17
                _ = v18
                _ = v19
            }
        case "a":
            if false {
                v1 += 2
                _ = len(v13)
                _ = len(v13)
            } else {
                v20 := v5[4294967295]
                delete(v4, 189)
                _ = v20
            }
        default:
            var v21 int64 = 1
            _ = v21
        }
        var v22 int8 = (v11 & v11)
        var v23 uint64 = 4294967295
        _ = v22
        _ = v23
    } else {
        fmt.Println(v3, v1)
        v3 = (0.0 + v3)
        v4[0] = (float32((1.5 - float64((-2.25 * ((3.14159 * v3) * 1e10))))) == (float32((((3.14159 - 3.14159) * 0.1) - 3.14159)) + 1e10))
    }
    var v24 uint = 547
    if false {
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 string = ""
            var v27 string = ""
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        fmt.Println(v3)
        for fi28 := 0; fi28 < 12; fi28++ {
            _ = len(v13)
            _ = fi28
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            switch fi29 {
            case 0:
                delete(v4, 4294967295)
            case -342:
                _ = len(v4)
            case 2147483647:
                if (-705 <= 1000) {
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v5[uint32(v1)] = 4294967295
                        var v31 float32 = (((v3 - v3) + v3) / ((1e10 + 1.5) + float32((((((0.1 - float64(v3)) + -2.25) * 1e10) + ((0.1 - 1e10) + ((1.5 * ((1.5 / 1.0) + ((1e10 * (1e10 / float64(0.1))) * 1.5))) + 0.0))) + (-2.25 - float64(1.5))))))
                        _ = fi30
                        _ = v31
                    }
                    if (!true) {
                        v5[1] = 842
                    }
                } else {
                    if false {
                        for fi32 := 0; fi32 < 5; fi32++ {
                            v33 := GenS2{F0: -653, F1: 2}
                            var v34 string = ""
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS3{F0: -1}
                        _ = v35
                    } else {
                        var v36 uint32 = 710
                        var v37 int32 = -175
                        var v38 int8 = v2
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    v2 = 42
                    if ((-2.25 != 0.0) && (788 == 971)) {
                        var v39 uint64 = 2
                        var v40 uint16 = 667
                        _ = v39
                        _ = v40
                    }
                }
            }
            fmt.Println(fi29, v3, v24)
            delete(v4, 0)
            _ = fi29
        }
    }
    switch ("" + "foo bar") {
    case "foo bar":
        if false {
            if true {
                if (true && false) {
                    for fi41 := 0; fi41 < 4; fi41++ {
                        fmt.Println(v24, v3)
                        _ = fi41
                    }
                } else {
                    if false {
                        if false {
                            switch "\n\t" {
                            case "hello":
                                v42 := GenS2{F0: 127, F1: 2147483647}
                                _ = v42
                            case "a":
                                v2 += v11
                            case "foo bar":
                                if false {
                                    if ((!false) || true) {
                                        v43 := GenS4{F0: -1}
                                        var v44 uint64 = 2
                                        v45 := GenS5{F0: 127}
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v46 uint8 = (v1 - v1)
                                        var v47 uint32 = 2147483647
                                        _ = v46
                                        _ = v47
                                    }
                                    for fi48 := 0; fi48 < 2; fi48++ {
                                        var v49 uint32 = 42
                                        v13 = append(v13, 32767)
                                        v11 += (v11 - (v11 >> ((2) & 63)))
                                        _ = fi48
                                        _ = v49
                                    }
                                    var v50 bool = false
                                    _ = v50
                                }
                            default:
                                v2 = v11
                            }
                        } else {
                            var v51 int = 0
                            _ = v51
                        }
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        if (!(v3 < (v3 * v3))) {
                            switch -627 {
                            case 128:
                                for fi52 := 0; fi52 < 4; fi52++ {
                                    fmt.Println(v24)
                                    _ = fi52
                                }
                            case 32768:
                                v2 = v2
                            }
                        } else {
                            fmt.Println(v24, v3, v2)
                            switch 0 {
                            case 9223372036854775807:
                                if (805 <= 4294967295) {
                                    v3 = float32(3.14159)
                                    fmt.Println(v2, v3, v2)
                                    for fi53 := 0; fi53 < 17; fi53++ {
                                        v2 = (v11 - (v11 & v2))
                                        for fi54 := 0; fi54 < 16; fi54++ {
                                            _ = len(v13)
                                            var v55 int8 = 74
                                            var v56 float64 = 1.5
                                            _ = fi54
                                            _ = v55
                                            _ = v56
                                        }
                                        v24 -= v24
                                        _ = fi53
                                    }
                                } else {
                                    for fi57 := 0; fi57 < 13; fi57++ {
                                        var v58 int64 = 9223372036854775807
                                        _ = fi57
                                        _ = v58
                                    }
                                    var v59 int8 = 100
                                    _ = v59
                                }
                            default:
                                v60 := GenS6{F0: 1000, F1: (v24 + (v24 % ((v24) | 1)))}
                                _ = v60
                            }
                        }
                    } else {
                        var v61 int32 = -2147483648
                        var v62 int = -151
                        _ = v61
                        _ = v62
                    }
                    if (18446744073709551615 > 400) {
                        fmt.Println(v2, v2)
                        var v63 float32 = (v3 - float32(0.1))
                        var v64 uint16 = 42
                        _ = v63
                        _ = v64
                    }
                }
                v1 -= (^(v1))
                v11 = v11
            }
            if (!true) {
                var v65 int32 = 144
                switch -464 {
                case 4294967295:
                    switch 0 {
                    case 9223372036854775807:
                        v1 = v1
                    case 4294967295:
                        switch 32767 {
                        case 4294967295:
                            if false {
                                fmt.Println(v65)
                                v13 = append(v13, -9223372036854775808)
                                var v66 uint16 = 2
                                _ = v66
                            } else {
                                if (0 > 9223372036854775807) {
                                    switch ("a" + "foo bar") {
                                    case "hello":
                                        fmt.Println(v1, v11, v11)
                                    case "\n\t":
                                        v24 = v24
                                    default:
                                        var v67 bool = (42 != 905)
                                        _ = v67
                                    }
                                    for fi68 := 0; fi68 < 15; fi68++ {
                                        var v69 uint16 = 0
                                        fmt.Println(v24, v2)
                                        _ = fi68
                                        _ = v69
                                    }
                                    var v70 float64 = 3.14159
                                    _ = v70
                                }
                                var v71 int64 = 202
                                v72 := GenS4{F0: v71}
                                _ = v71
                                _ = v72
                            }
                        case 1:
                            switch 32768 {
                            case 0:
                                for fi73 := 0; fi73 < 3; fi73++ {
                                    for fi74 := 0; fi74 < 7; fi74++ {
                                        v75 := GenS7{F0: 9223372036854775807, F1: 759}
                                        var v76 uint8 = (v1 << ((v24) & 63))
                                        fmt.Println(v2, v24)
                                        fmt.Println(v1)
                                        _ = fi74
                                        _ = v75
                                        _ = v76
                                    }
                                    if false {
                                        for fi77 := 0; fi77 < 3; fi77++ {
                                            var v78 uint = 2
                                            var v79 uint8 = v1
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                        var v80 uint16 = 42
                                        _ = v80
                                    } else {
                                        var v81 float64 = -2.25
                                        var v82 int32 = v65
                                        var v83 uint16 = 2
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                    }
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    _ = fi73
                                }
                            case 32767:
                                var v84 int8 = 86
                                _ = v84
                            case 1:
                                if false {
                                    var v85 float64 = ((-2.25 * (-2.25 * 1.5)) / 0.1)
                                    var v86 int32 = v65
                                    switch -201 {
                                    case -9223372036854775808:
                                        if (3.14159 <= ((v3 - 1.5) + -2.25)) {
                                            var v87 uint32 = 1
                                            for fi88 := 0; fi88 < 17; fi88++ {
                                                var v89 uint64 = 960
                                                _ = len(v5)
                                                _ = fi88
                                                _ = v89
                                            }
                                            if (32767 >= -1) {
                                                v87 = v87
                                                _ = len(v5)
                                            }
                                            _ = v87
                                        } else {
                                            v90 := GenS5{F0: 0}
                                            switch 32768 {
                                            case 32768:
                                                v91 := GenS8{F0: -247, F1: 736, F2: "foo bar"}
                                                _ = v91
                                            case -664:
                                                _ = len(v13)
                                            case 42:
                                                var v92 string = "hello"
                                                _ = v92
                                            }
                                            _ = v90
                                        }
                                    case 1:
                                        if false {
                                            var v93 uint8 = uint8(v24)
                                            for fi94 := 0; fi94 < 9; fi94++ {
                                                v5[2147483648] = 4294967295
                                                v2 = v11
                                                v95 := GenS6{F0: 1, F1: 0}
                                                _ = fi94
                                                _ = v95
                                            }
                                            _ = v93
                                        }
                                    default:
                                        v11 = (v2 & 42)
                                    }
                                    _ = v85
                                    _ = v86
                                }
                            default:
                                switch ("hello" + "foo bar") {
                                case "":
                                    if false {
                                        if false {
                                            for fi96 := 0; fi96 < 20; fi96++ {
                                                fmt.Println(v1)
                                                if true {
                                                    var v97 bool = true
                                                    _ = v97
                                                }
                                                v98 := GenS6{F0: 42, F1: (^(v24))}
                                                _ = fi96
                                                _ = v98
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                for fi99 := 0; fi99 < 7; fi99++ {
                                                    if true {
                                                        var v100 int16 = -727
                                                        v11 += (^(v11))
                                                        fmt.Println(v24, v3, v24)
                                                        _ = v100
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v1 += 42
                                                        case "foo bar":
                                                            var v101 int = (fi99 * fi99)
                                                            _ = v101
                                                        }
                                                    }
                                                    fmt.Println(fi99)
                                                    _ = fi99
                                                }
                                            }
                                        } else {
                                            fmt.Println(v3, v2)
                                            v102 := GenS4{F0: 42}
                                            for fi103 := 0; fi103 < 8; fi103++ {
                                                var v104 float64 = -2.25
                                                _ = fi103
                                                _ = v104
                                            }
                                            _ = v102
                                        }
                                        var v105 uint8 = (^(v1))
                                        _ = v105
                                    }
                                case "a":
                                    v106 := GenS9{F0: 65535, F1: 2}
                                    _ = v106
                                }
                            }
                        default:
                            for fi107 := 0; fi107 < 3; fi107++ {
                                switch ("" + ("foo bar" + "a")) {
                                case "\n\t":
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                case "a":
                                    if ((!((fi107 / (((fi107 | 9223372036854775807)) | 1)) == (fi107 | (^(fi107))))) || false) {
                                        for fi108 := 0; fi108 < 4; fi108++ {
                                            v13 = append(v13, 56)
                                            _ = fi108
                                        }
                                    } else {
                                        v3 = v3
                                        var v109 uint64 = 0
                                        var v110 uint = v24
                                        _ = v109
                                        _ = v110
                                    }
                                case "hello":
                                    v2 = v11
                                default:
                                    v11 = (v2 % (((v11 >> (((^(v24))) & 63))) | 1))
                                }
                                var v111 float64 = 0.1
                                if false {
                                    var v112 int8 = -128
                                    if true {
                                        for fi113 := 0; fi113 < 1; fi113++ {
                                            var v114 uint = v24
                                            _ = fi113
                                            _ = v114
                                        }
                                        var v115 float32 = v3
                                        for fi116 := 0; fi116 < 14; fi116++ {
                                            if false {
                                                switch (fi116 / ((32768) | 1)) {
                                                case 128:
                                                    var v117 int64 = 9223372036854775807
                                                    _ = v117
                                                default:
                                                    switch ("\n\t" + "foo bar") {
                                                    case "":
                                                        fmt.Println(v3, fi107)
                                                    }
                                                }
                                                var v118 string = "foo bar"
                                                _ = v118
                                            }
                                            _ = fi116
                                        }
                                        _ = v115
                                    }
                                    _ = v112
                                }
                                _ = fi107
                                _ = v111
                            }
                        }
                    }
                case 32768:
                    delete(v5, 541)
                }
                _ = v65
            } else {
                for fi119 := 0; fi119 < 18; fi119++ {
                    if (0 < 1) {
                        v3 = (v3 / v3)
                    }
                    _ = fi119
                }
                v11 = (v11 - v2)
            }
        } else {
            v24 = 42
            switch "\n\t" {
            case "a":
                v5[2147483648] = 2
            }
        }
    }
    fmt.Println(v1, v24, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
