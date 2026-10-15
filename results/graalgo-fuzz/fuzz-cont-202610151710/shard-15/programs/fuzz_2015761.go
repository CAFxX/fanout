package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32; F1 int16; F2 int8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 int) int16 {
    var v4 int = p0
    v4 = (-(v4))
    if false {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 uint32 = 2
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 14; fi7++ {
            switch 42 {
            case -596:
                switch (fi7 & 0) {
                case 0:
                    v8 := GenS1{F0: 9223372036854775807, F1: 42, F2: 18446744073709551615}
                    _ = v8
                case -9223372036854775808:
                    var v9 string = (("foo bar" + ("hello" + "foo bar")) + "foo bar")
                    _ = v9
                }
            case 162:
                v10 := GenS2{F0: 612, F1: -602, F2: -1}
                _ = v10
            case 32767:
                var v11 uint = 287
                _ = v11
            default:
                var v12 int8 = 114
                _ = v12
            }
            _ = fi7
        }
    } else {
        var v13 string = ("foo bar" + "a")
        var v14 int64 = 9223372036854775807
        _ = v13
        _ = v14
    }
    var v15 float32 = -2.25
    fmt.Println(v15, p0, v15)
    _ = p0
    _ = v15
    _ = v4
    return -32768
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    v3 := make(map[int]string)
    v16 := []int{-9223372036854775808, 4294967295, -2147483648, 128}
    if ((false || (42 == 255)) || (!(true || true))) {
        var v17 uint32 = 4294967295
        v17 += v17
        v18 := v3[-9223372036854775808]
        _ = v17
        _ = v18
    } else {
        if false {
            var v19 uint8 = 159
            if false {
                if (!((^(v19)) >= (v19 >> ((255) & 63)))) {
                    switch -9223372036854775808 {
                    case 2147483647:
                        fmt.Println(v19, v2)
                    case -9223372036854775808:
                        fmt.Println(v19, v19)
                    default:
                        for fi20 := 0; fi20 < 18; fi20++ {
                            fmt.Println(fi20)
                            v21 := GenS1{F0: 836, F1: 0, F2: 42}
                            _ = fi20
                            _ = v21
                        }
                    }
                    delete(v3, 32767)
                    fmt.Println(v19, v2)
                }
                var v22 bool = (!false)
                v23 := GenS3{F0: 32767, F1: (^(v19))}
                _ = v22
                _ = v23
            }
            _ = v19
        }
        var v24 float32 = 0.1
        if (!(("hello" + ("\n\t" + "")) > "foo bar")) {
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := GenS4{F0: 1, F1: 995}
                fmt.Println(v24)
                _ = fi25
                _ = v26
            }
            if len(v16) > 0 {
                v16[0] = 127
            }
        } else {
            if false {
                if (0 != uint16(v2)) {
                    v2 += 2
                    fmt.Println(v24)
                }
                v27 := GenS1{F0: -689, F1: 255, F2: 1}
                _ = v27
            }
            v28 := GenS5{F0: 102, F1: (("" + "\n\t") + ("hello" + "a"))}
            v2 = v2
            _ = v28
        }
        _ = v24
    }
    fmt.Println(v2)
    switch "foo bar" {
    case "hello":
        for fi29 := 0; fi29 < 6; fi29++ {
            v30 := GenS3{F0: -32768, F1: 255}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            v2 += 397
            _ = fi29
            _ = v30
        }
    case "a":
        v31 := GenS4{F0: 2, F1: 2147483648}
        _ = v31
    default:
        var v32 uint32 = 224
        _ = v32
    }
    var v33 int = 1
    if (42 <= 32767) {
        v34 := GenS6{F0: (("\n\t" + "hello") + ("a" + "foo bar"))}
        _ = v34
    } else {
        var v35 int32 = 0
        _ = v35
    }
    if (("" + "\n\t") >= ("\n\t" + ("foo bar" + ("foo bar" + "")))) {
        if true {
            for fi36 := 0; fi36 < 8; fi36++ {
                fmt.Println(fi36, v2)
                _ = fi36
            }
            v37 := GenS4{F0: 18446744073709551615, F1: 4294967295}
            if (0 >= 255) {
                v38 := GenS7{F0: 0}
                if (1 != 482) {
                    if false {
                        var v39 uint16 = 2
                        _ = v39
                    } else {
                        switch "" {
                        case "\n\t":
                            _ = len(v16)
                        case "":
                            var v40 int8 = 42
                            _ = v40
                        case "hello":
                            v33 += 13
                        }
                    }
                    v41 := GenS6{F0: "foo bar"}
                    _ = v41
                } else {
                    v42 := GenS3{F0: 42, F1: 1}
                    switch ("hello" + "") {
                    case "\n\t":
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    case "hello":
                        if false {
                            if false {
                                var v43 uint = 298
                                _ = v43
                            } else {
                                if false {
                                    delete(v1, 42)
                                    v33 = -642
                                    v2 -= 1
                                } else {
                                    v2 += 1
                                    v44 := GenS3{F0: -637, F1: 0}
                                    if true {
                                        v45 := GenS1{F0: 9223372036854775807, F1: 2, F2: 18446744073709551615}
                                        switch (v33 >> ((2) & 63)) {
                                        case -2147483648:
                                            for fi46 := 0; fi46 < 9; fi46++ {
                                                v47 := GenS8{F0: (v2 - 614), F1: 1}
                                                v48 := GenS2{F0: uint32(v2), F1: genFn0((v33 >> ((2) & 63))), F2: -128}
                                                v33 -= v33
                                                _ = fi46
                                                _ = v47
                                                _ = v48
                                            }
                                        case 759:
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                        default:
                                            if (!(!(32767 <= v33))) {
                                                fmt.Println(v33, v33)
                                            }
                                        }
                                        if (0 == 150) {
                                            var v49 uint16 = 2
                                            v2 = 1
                                            _ = v49
                                        }
                                        _ = v45
                                    } else {
                                        for fi50 := 0; fi50 < 16; fi50++ {
                                            if (0.1 != (((float32(float64(3.14159)) / 0.1) / 0.1) - 3.14159)) {
                                                _ = len(v3)
                                                var v51 uint32 = 2147483647
                                                v52 := GenS9{F0: (^(v51))}
                                                _ = v51
                                                _ = v52
                                            } else {
                                                var v53 int64 = -91
                                                if true {
                                                    for fi54 := 0; fi54 < 16; fi54++ {
                                                        var v55 int8 = -1
                                                        v16 = append(v16, 69)
                                                        v56 := GenS1{F0: (v53 | (v53 * int64(v2))), F1: 2, F2: 42}
                                                        v16 = append(v16, fi50)
                                                        _ = fi54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    for fi57 := 0; fi57 < 15; fi57++ {
                                                        var v58 int64 = (v53 / (((v53 | v53)) | 1))
                                                        if true {
                                                            delete(v1, 1)
                                                            v33 = 2147483647
                                                            v33 = (v33 - fi57)
                                                        }
                                                        _ = fi57
                                                        _ = v58
                                                    }
                                                }
                                                _ = v53
                                            }
                                            var v59 int64 = -985
                                            _ = fi50
                                            _ = v59
                                        }
                                    }
                                    _ = v44
                                }
                                for fi60 := 0; fi60 < 11; fi60++ {
                                    v2 -= (v2 / ((v2) | 1))
                                    _ = fi60
                                }
                            }
                        }
                    case "a":
                        for fi61 := 0; fi61 < 12; fi61++ {
                            var v62 uint16 = 32767
                            _ = fi61
                            _ = v62
                        }
                    default:
                        v33 = -866
                    }
                    _ = v42
                }
                v2 -= 0
                _ = v38
            } else {
                if len(v16) > 0 {
                    v16[0] = (-(v33))
                }
                var v63 float64 = 1e10
                _ = v63
            }
            _ = v37
        }
        v64 := GenS8{F0: v2, F1: 1}
        _ = v64
    } else {
        var v65 uint64 = 4294967295
        if false {
            v2 += v65
        }
        switch v33 {
        case 835:
            v66 := GenS10{F0: v33, F1: 1}
            _ = v66
        case -730:
            var v67 string = ("" + "foo bar")
            _ = v67
        default:
            var v68 string = ("foo bar" + ("a" + (("\n\t" + (("" + "a") + "\n\t")) + (("\n\t" + "hello") + ("hello" + (("foo bar" + "a") + "a"))))))
            _ = v68
        }
        _ = v65
    }
    switch (("" + "foo bar") + "hello") {
    case "hello":
        v33 = -1
    case "\n\t":
        for fi69 := 0; fi69 < 13; fi69++ {
            if ((^(fi69)) > (fi69 / ((fi69) | 1))) {
                switch 42 {
                case 146:
                    if false {
                        v33 -= -1
                        if len(v16) > 0 {
                            v16[0] = (fi69 - (fi69 << ((255) & 63)))
                        }
                    }
                case 1:
                    v33 -= (^(fi69))
                case -2147483648:
                    for fi70 := 0; fi70 < 9; fi70++ {
                        v33 += (fi69 / ((747) | 1))
                        v2 = (v2 / (((v2 >> ((1) & 63))) | 1))
                        _ = fi70
                    }
                default:
                    for fi71 := 0; fi71 < 8; fi71++ {
                        v33 = fi71
                        v72 := GenS10{F0: (fi69 << ((287) & 63)), F1: 65535}
                        var v73 int16 = 1
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                }
                switch v33 {
                case 32767:
                    if len(v16) > 0 {
                        v16[0] = 1
                    }
                default:
                    v33 = v33
                }
            } else {
                v16 = append(v16, v33)
                for fi74 := 0; fi74 < 9; fi74++ {
                    if (v2 == (v2 * v2)) {
                        v2 = 2
                        var v75 uint64 = (v2 * v2)
                        var v76 int = 127
                        _ = v75
                        _ = v76
                    }
                    if true {
                        var v77 uint = 97
                        fmt.Println(fi74, v2)
                        var v78 int32 = 62
                        _ = v77
                        _ = v78
                    } else {
                        switch "" {
                        case "foo bar":
                            _ = len(v16)
                        case "":
                            fmt.Println(v33, v2)
                        case "a":
                            var v79 float32 = 1e10
                            _ = v79
                        }
                        var v80 uint16 = 65535
                        _ = v80
                    }
                    _ = len(v16)
                    _ = fi74
                }
            }
            switch "a" {
            case "hello":
                delete(v3, (v33 / (((v33 + v33)) | 1)))
            case "a":
                v2 += (v2 + v2)
            case "\n\t":
                v2 -= (v2 << ((18446744073709551615) & 63))
            }
            _ = fi69
        }
    case "foo bar":
        switch v33 {
        case 1:
            switch v33 {
            case 1:
                if (2 >= uint8(v2)) {
                    v33 -= v33
                } else {
                    v81 := GenS11{F0: -118, F1: 2}
                    v82 := GenS12{F0: -587}
                    _ = v81
                    _ = v82
                }
            case -959:
                fmt.Println(v2)
            case 127:
                var v83 float32 = ((float32(1e10) - ((-2.25 - 1e10) * -2.25)) / 1e10)
                _ = v83
            default:
                var v84 int64 = 42
                _ = v84
            }
        case -882:
            fmt.Println(v2, v33)
        }
    default:
        v33 -= v33
    }
    for fi85 := 0; fi85 < 11; fi85++ {
        v33 += (fi85 ^ v33)
        _ = fi85
    }
    v86 := v1[4294967295]
    var v87 string = "hello"
    for fi88 := 0; fi88 < 15; fi88++ {
        var v89 int8 = 55
        v90 := GenS9{F0: 990}
        if true {
            var v91 uint32 = 2147483647
            var v92 int = -883
            v93 := GenS13{F0: 9223372036854775807}
            _ = v91
            _ = v92
            _ = v93
        } else {
            var v94 uint64 = v2
            var v95 int8 = (-(v89))
            _ = v94
            _ = v95
        }
        _ = fi88
        _ = v89
        _ = v90
    }
    var v96 string = ((v87 + "hello") + "hello")
    for fi97 := 0; fi97 < 11; fi97++ {
        var v98 uint32 = (v86 % ((964) | 1))
        if (65535 < 441) {
            var v99 uint = 355
            _ = v99
        } else {
            var v100 uint32 = (^(v86))
            _ = v100
        }
        _ = fi97
        _ = v98
    }
    if len(v16) > 0 {
        v16[0] = v33
    }
    var v101 int64 = 9223372036854775807
    if false {
        if false {
            if ((2 != 1) || false) {
                if false {
                    for fi102 := 0; fi102 < 1; fi102++ {
                        if false {
                            v16 = append(v16, (^(fi102)))
                            v103 := GenS5{F0: (v86 & (v86 >> ((65535) & 63))), F1: "a"}
                            v104 := v1[uint32(v2)]
                            _ = v103
                            _ = v104
                        }
                        v33 = fi102
                        for fi105 := 0; fi105 < 20; fi105++ {
                            fmt.Println(v2, v96, v2)
                            _ = fi105
                        }
                        _ = fi102
                    }
                    if (0.1 > (1.5 / (float64((-2.25 * -2.25)) - ((0.0 / -2.25) / (1.5 - ((3.14159 / 1.0) - -2.25)))))) {
                        var v106 uint32 = uint32(v2)
                        for fi107 := 0; fi107 < 16; fi107++ {
                            if true {
                                if false {
                                    var v108 float64 = 0.1
                                    _ = v108
                                } else {
                                    var v109 uint64 = 2
                                    _ = v109
                                }
                            }
                            _ = fi107
                        }
                        _ = v106
                    }
                } else {
                    if ((1.5 * 0.1) > 1e10) {
                        v110 := GenS5{F0: v86, F1: v87}
                        _ = v110
                    }
                    v111 := GenS14{F0: v2}
                    _ = len(v16)
                    _ = v111
                }
                var v112 float64 = 0.0
                _ = v112
            }
            v87 = v96
        }
        v101 -= v101
        switch (v87 + v96) {
        case "\n\t":
            v113 := GenS14{F0: (v2 / (((v2 >> ((uint(v2)) & 63))) | 1))}
            _ = v113
        case "foo bar":
            var v114 bool = true
            _ = v114
        case "":
            switch "a" {
            case "a":
                fmt.Println(v2, v2)
            case "hello":
                var v115 float32 = 0.0
                _ = v115
            default:
                v116 := GenS8{F0: (^(v2)), F1: (v86 | 1)}
                _ = v116
            }
        default:
            switch (v33 / ((v33) | 1)) {
            case 286:
                var v117 uint = 2
                _ = v117
            case 127:
                fmt.Println(v87)
            case 128:
                v96 = ((v96 + v96) + v96)
            default:
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            }
        }
    } else {
        if ((v101 >> ((1) & 63)) != 9223372036854775807) {
            if (v86 == 1) {
                v118 := GenS15{F0: -32768, F1: v33, F2: "foo bar"}
                fmt.Println(v86)
                _ = v118
            } else {
                fmt.Println(v101, v101)
            }
            switch (((v87 + v96) + "") + "hello") {
            case "\n\t":
                fmt.Println(v96, v87, v33)
            case "foo bar":
                for fi119 := 0; fi119 < 12; fi119++ {
                    v2 -= (v2 & v2)
                    v101 = -1
                    _ = fi119
                }
            case "a":
                var v120 float32 = 0.1
                _ = v120
            default:
                switch v96 {
                case "foo bar":
                    if (v96 >= "\n\t") {
                        v2 += v2
                        _ = len(v16)
                    }
                case "hello":
                    fmt.Println(v96, v87)
                case "a":
                    for fi121 := 0; fi121 < 10; fi121++ {
                        v101 -= (v101 / ((94) | 1))
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        _ = fi121
                    }
                default:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                }
            }
            v122 := GenS10{F0: 0, F1: (^(v86))}
            _ = v122
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v101
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v86
    _ = v87
    _ = v96
}
