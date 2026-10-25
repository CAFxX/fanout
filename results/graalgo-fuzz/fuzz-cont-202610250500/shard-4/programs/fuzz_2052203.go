package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 uint32; F1 string; F2 bool }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint32; F1 uint; F2 bool }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 bool; F1 uint8; F2 int64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 string; F1 uint32 }
type GenS19 struct { F0 int; F1 uint64 }
type GenS20 struct { F0 int16; F1 uint64; F2 uint8 }
func genFn0(p0 string) uint32 {
    v6 := []int{32768}
    if (!false) {
        v7 := GenS1{F0: (42 >= 127)}
        v8 := GenS1{F0: (false && (true || false))}
        if (!(false || false)) {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS2{F0: (-2.25 == 0.0), F1: -53}
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: (42 <= 128)}
            _ = v11
        } else {
            p0 = ("\n\t" + "\n\t")
            v12 := GenS1{F0: true}
            if false {
                for fi13 := 0; fi13 < 2; fi13++ {
                    if (true && false) {
                        if true {
                            fmt.Println(p0)
                            if true {
                                v14 := GenS3{F0: 1000}
                                _ = v14
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = (fi13 - (-(fi13)))
                                }
                            }
                        }
                        switch p0 {
                        case "a":
                            v15 := GenS4{F0: -1, F1: 65535}
                            _ = v15
                        case "hello":
                            fmt.Println(p0, p0)
                        case "":
                            var v16 uint16 = 283
                            _ = v16
                        }
                        v17 := GenS1{F0: (((255 < 136) || false) && (!true))}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 14; fi18++ {
                        var v19 uint16 = 1
                        if ((true && false) || (true || false)) {
                            var v20 uint = 42
                            _ = v20
                        } else {
                            switch "foo bar" {
                            case "hello":
                                var v21 uint16 = (v19 | v19)
                                _ = v21
                            case "":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            case "a":
                                if (!(true && false)) {
                                    var v22 float32 = 1.5
                                    if true {
                                        v22 += (0.0 + v22)
                                        if true {
                                            var v23 uint64 = 4294967295
                                            switch ("foo bar" + (p0 + "hello")) {
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "a":
                                                    var v24 float64 = (1e10 - ((0.0 - 0.0) + (0.0 * 0.0)))
                                                    _ = v24
                                                case "hello":
                                                    _ = len(v6)
                                                default:
                                                    if ((!(((!true) || true) && (218 > 852))) && false) {
                                                        var v25 uint8 = 255
                                                        if false {
                                                            var v26 int = fi18
                                                            var v27 float32 = v22
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        _ = v25
                                                    } else {
                                                        if (-128 != 1) {
                                                            v28 := GenS1{F0: (!(-2147483648 != 0))}
                                                            switch p0 {
                                                            case "a":
                                                                switch p0 {
                                                                case "a":
                                                                    if false {
                                                                        fmt.Println(v19, v19, fi13)
                                                                        switch -366 {
                                                                        case -9223372036854775808:
                                                                            v29 := GenS5{F0: 127}
                                                                            _ = v29
                                                                        case 42:
                                                                            var v30 int = (fi18 / ((fi13) | 1))
                                                                            _ = v30
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v31 uint32 = 595
                                                                    _ = v31
                                                                case "":
                                                                    var v32 int8 = 0
                                                                    var v33 uint64 = 4294967295
                                                                    _ = v32
                                                                    _ = v33
                                                                }
                                                            case "foo bar":
                                                                v34 := GenS1{F0: true}
                                                                _ = v34
                                                            default:
                                                                var v35 string = ("" + ("hello" + ((((p0 + (p0 + p0)) + "foo bar") + "\n\t") + p0)))
                                                                _ = v35
                                                            }
                                                            _ = v28
                                                        } else {
                                                            if ((true || (!((!false) && false))) && (true && false)) {
                                                                if len(v6) > 0 {
                                                                    v6[0] = (fi18 - (fi18 << ((1) & 63)))
                                                                }
                                                            }
                                                        }
                                                        if true {
                                                            var v36 float32 = v22
                                                            switch 2147483647 {
                                                            case 32767:
                                                                v19 += (v19 / ((65535) | 1))
                                                            case 128:
                                                                v37 := GenS4{F0: -9223372036854775808, F1: 65535}
                                                                _ = v37
                                                            case -249:
                                                                v36 -= v22
                                                                var v38 uint32 = 1
                                                                if false {
                                                                    v39 := GenS6{F0: v23, F1: 4294967295}
                                                                    if (18446744073709551615 < uint64(v38)) {
                                                                        switch (fi13 & 0) {
                                                                        case 128:
                                                                            var v40 bool = ((false || (v19 <= v19)) || true)
                                                                            _ = v40
                                                                        case -2147483648:
                                                                            var v41 bool = (true && false)
                                                                            _ = v41
                                                                        }
                                                                        fmt.Println(fi18, p0, v38)
                                                                        var v42 string = (((p0 + "a") + "hello") + p0)
                                                                        _ = v42
                                                                    } else {
                                                                        fmt.Println(v23, v36)
                                                                        switch "a" {
                                                                        case "a":
                                                                            v38 = v38
                                                                        case "foo bar":
                                                                            if (42 <= 435) {
                                                                                switch (fi13 % ((fi18) | 1)) {
                                                                                case -497:
                                                                                    v23 -= 4294967295
                                                                                case 223:
                                                                                    var v43 int32 = 2147483647
                                                                                    _ = v43
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v36)
                                                                                default:
                                                                                    v44 := GenS7{F0: v23, F1: false}
                                                                                    _ = v44
                                                                                }
                                                                            }
                                                                        default:
                                                                            _ = len(v6)
                                                                        }
                                                                        var v45 uint32 = v38
                                                                        v46 := GenS8{F0: 127, F1: -93}
                                                                        if (442 != 0) {
                                                                            switch (("\n\t" + p0) + "foo bar") {
                                                                            case "\n\t":
                                                                                switch fi13 {
                                                                                case 32768:
                                                                                    var v47 bool = false
                                                                                    _ = v47
                                                                                default:
                                                                                    var v48 int64 = -9223372036854775808
                                                                                    _ = v48
                                                                                }
                                                                            case "a":
                                                                                v49 := GenS9{F0: 42, F1: (v19 << ((255) & 63)), F2: v23}
                                                                                _ = v49
                                                                            case "":
                                                                                if true {
                                                                                    fmt.Println(v22)
                                                                                    v50 := GenS10{F0: false, F1: 1}
                                                                                    _ = v50
                                                                                } else {
                                                                                    _ = len(v6)
                                                                                }
                                                                            }
                                                                            fmt.Println(v45, v45, v19)
                                                                        }
                                                                        _ = v45
                                                                        _ = v46
                                                                    }
                                                                    v6 = append(v6, fi13)
                                                                    _ = v39
                                                                }
                                                                _ = v38
                                                            }
                                                            var v51 int8 = 100
                                                            _ = v36
                                                            _ = v51
                                                        }
                                                        v19 = v19
                                                    }
                                                    var v52 uint8 = 155
                                                    _ = v52
                                                }
                                                var v53 int64 = 727
                                                _ = v53
                                            }
                                            _ = v23
                                        } else {
                                            v19 = (v19 | v19)
                                        }
                                    } else {
                                        switch (p0 + p0) {
                                        case "":
                                            if len(v6) > 0 {
                                                v6[0] = fi13
                                            }
                                        case "a":
                                            v54 := GenS11{F0: 2, F1: ("foo bar" + p0), F2: (-1 == -1)}
                                            _ = v54
                                        case "\n\t":
                                            if true {
                                                p0 = p0
                                                fmt.Println(p0, p0, v22)
                                                v55 := GenS11{F0: 151, F1: "a", F2: false}
                                                var v56 uint32 = 558
                                                switch fi13 {
                                                case -2147483648:
                                                    var v57 uint8 = uint8(v19)
                                                    _ = v57
                                                }
                                                _ = v55
                                                _ = v56
                                            } else {
                                                p0 = (p0 + p0)
                                                v19 -= (v19 % ((v19) | 1))
                                                var v58 uint16 = 968
                                                _ = v58
                                            }
                                        }
                                    }
                                    if (-2.25 > 1e10) {
                                        switch ((p0 + "hello") + (("hello" + p0) + p0)) {
                                        case "hello":
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v22, fi13, fi13)
                                            case "foo bar":
                                                v59 := GenS4{F0: -9223372036854775808, F1: 65535}
                                                _ = v59
                                            }
                                        case "a":
                                            var v60 int16 = 42
                                            _ = v60
                                        case "\n\t":
                                            switch fi13 {
                                            case 128:
                                                _ = len(v6)
                                            default:
                                                var v61 uint16 = v19
                                                _ = v61
                                            }
                                        default:
                                            var v62 int32 = 1
                                            _ = v62
                                        }
                                        var v63 string = p0
                                        v6 = append(v6, 2147483647)
                                        v22 -= float32((float64(float32(-2.25)) + ((float64((3.14159 + ((v22 / float32(-2.25)) / 1.5))) / 1.5) - -2.25)))
                                        fmt.Println(v19)
                                        v6 = append(v6, 2147483647)
                                        var v64 uint32 = 320
                                        _ = v63
                                        _ = v64
                                    } else {
                                        v6 = append(v6, fi18)
                                    }
                                    _ = v22
                                } else {
                                    v65 := GenS12{F0: 42}
                                    _ = v65
                                }
                                switch p0 {
                                case "a":
                                    v66 := GenS13{F0: 408, F1: 1, F2: (2 < v19)}
                                    _ = v66
                                }
                            }
                        }
                        v67 := GenS13{F0: 2147483647, F1: 608, F2: true}
                        _ = fi18
                        _ = v19
                        _ = v67
                    }
                    _ = fi13
                }
                var v68 int8 = 0
                _ = v68
            } else {
                if false {
                    for fi69 := 0; fi69 < 9; fi69++ {
                        var v70 int64 = 622
                        fmt.Println(v70, v70)
                        if true {
                            var v71 int64 = v70
                            _ = v71
                        } else {
                            switch ("a" + "\n\t") {
                            case "hello":
                                v72 := GenS14{F0: 1}
                                _ = v72
                            case "foo bar":
                                var v73 uint16 = 249
                                _ = v73
                            case "\n\t":
                                v74 := GenS15{F0: 1}
                                _ = v74
                            default:
                                v70 += v70
                            }
                            var v75 int16 = 1000
                            switch p0 {
                            case "":
                                fmt.Println(p0, p0, v70)
                            default:
                                v6 = append(v6, fi69)
                            }
                            _ = v75
                        }
                        _ = fi69
                        _ = v70
                    }
                    fmt.Println(p0)
                } else {
                    v76 := GenS7{F0: 1, F1: ((557 != -9223372036854775808) && ((((3.14159 + 1.5) * 0.1) >= 0.1) && (true || (!false))))}
                    switch 2147483647 {
                    case -9223372036854775808:
                        if (1 <= 409) {
                            if true {
                                fmt.Println(p0)
                                var v77 int32 = 42
                                _ = v77
                            }
                            fmt.Println(p0)
                            fmt.Println(p0)
                        } else {
                            for fi78 := 0; fi78 < 7; fi78++ {
                                var v79 int = (fi78 - fi78)
                                _ = fi78
                                _ = v79
                            }
                            switch ((p0 + p0) + p0) {
                            case "hello":
                                if (128 != 63) {
                                    fmt.Println(p0)
                                    v80 := GenS10{F0: true, F1: 1}
                                    _ = v80
                                } else {
                                    var v81 int16 = -558
                                    for fi82 := 0; fi82 < 12; fi82++ {
                                        for fi83 := 0; fi83 < 1; fi83++ {
                                            if (true && false) {
                                                v81 += (^(v81))
                                                switch 0 {
                                                case -2147483648:
                                                    if true {
                                                        v84 := GenS16{F0: (!(fi83 <= (fi82 | (-(fi83))))), F1: 116, F2: 42}
                                                        if (true && true) {
                                                            if len(v6) > 0 {
                                                                v6[0] = 32767
                                                            }
                                                            v85 := GenS5{F0: -1}
                                                            _ = v85
                                                        }
                                                        _ = v84
                                                    } else {
                                                        if ((v81 << ((0) & 63)) > (v81 ^ (v81 % ((-32768) | 1)))) {
                                                            var v86 uint = 0
                                                            v81 += (-(v81))
                                                            _ = v86
                                                        } else {
                                                            if true {
                                                                p0 = p0
                                                                fmt.Println(fi83, fi83, fi83)
                                                                var v87 uint32 = 4294967295
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v87 = v87
                                                                case "":
                                                                    v6 = append(v6, (fi82 ^ fi83))
                                                                case "hello":
                                                                    p0 = ""
                                                                default:
                                                                    switch ("\n\t" + p0) {
                                                                    case "hello":
                                                                        v81 = -320
                                                                    default:
                                                                        if len(v6) > 0 {
                                                                            v6[0] = 1
                                                                        }
                                                                    }
                                                                }
                                                                _ = v87
                                                            }
                                                        }
                                                        fmt.Println(p0)
                                                        if (!((!((v81 % (((v81 & v81)) | 1)) <= v81)) && true)) {
                                                            if len(v6) > 0 {
                                                                v6[0] = -2147483648
                                                            }
                                                        } else {
                                                            v81 = 1
                                                            if len(v6) > 0 {
                                                                v6[0] = (^(fi83))
                                                            }
                                                            switch (("foo bar" + p0) + p0) {
                                                            case "a":
                                                                fmt.Println(p0)
                                                            case "":
                                                                if true {
                                                                    _ = len(v6)
                                                                } else {
                                                                    var v88 uint = 1
                                                                    _ = v88
                                                                }
                                                            default:
                                                                var v89 int16 = (v81 >> ((4294967295) & 63))
                                                                _ = v89
                                                            }
                                                        }
                                                    }
                                                    var v90 string = "foo bar"
                                                    p0 = "a"
                                                    _ = v90
                                                }
                                                switch fi82 {
                                                case -9223372036854775808:
                                                    _ = len(v6)
                                                case 32768:
                                                    v91 := GenS17{F0: -124}
                                                    _ = v91
                                                case -2147483648:
                                                    p0 = p0
                                                }
                                                var v92 uint32 = 4294967295
                                                _ = v92
                                            } else {
                                                var v93 float64 = 1.5
                                                var v94 int64 = 409
                                                switch p0 {
                                                case "":
                                                    _ = len(v6)
                                                case "foo bar":
                                                    switch (fi82 ^ fi83) {
                                                    case 32767:
                                                        var v95 uint8 = 178
                                                        _ = v95
                                                    case 43:
                                                        if false {
                                                            v94 = (v94 / (((v94 ^ v94)) | 1))
                                                            p0 = (p0 + ((p0 + "foo bar") + p0))
                                                            if true {
                                                                _ = len(v6)
                                                                _ = len(v6)
                                                                v96 := GenS18{F0: ((p0 + (p0 + p0)) + (((((p0 + p0) + ("" + "hello")) + p0) + "\n\t") + ((p0 + ("hello" + p0)) + ((p0 + p0) + "")))), F1: 578}
                                                                _ = v96
                                                            }
                                                            var v97 int = 0
                                                            v94 = -9223372036854775808
                                                            _ = v97
                                                        } else {
                                                            v93 += (v93 / (v93 - 1.5))
                                                        }
                                                    case -194:
                                                        v81 -= int16(fi82)
                                                        v94 = v94
                                                        var v98 uint32 = 101
                                                        _ = v98
                                                    default:
                                                        var v99 float32 = 0.0
                                                        _ = v99
                                                    }
                                                }
                                                var v100 float32 = (3.14159 / float32(v93))
                                                v101 := GenS19{F0: (fi83 | fi82), F1: 2}
                                                var v102 float32 = 1.5
                                                v103 := GenS2{F0: (!false), F1: 1}
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                                _ = v93
                                                _ = v94
                                            }
                                            _ = fi83
                                        }
                                        var v104 int = (fi82 - (fi82 + (fi82 & fi82)))
                                        _ = fi82
                                        _ = v104
                                    }
                                    v81 += 319
                                    _ = v81
                                }
                            case "a":
                                var v105 int8 = 127
                                _ = v105
                            case "":
                                if false {
                                    var v106 string = p0
                                    for fi107 := 0; fi107 < 6; fi107++ {
                                        for fi108 := 0; fi108 < 10; fi108++ {
                                            var v109 uint64 = 1
                                            fmt.Println(v109, fi107, fi108)
                                            if true {
                                                var v110 float32 = -2.25
                                                _ = v110
                                            }
                                            _ = fi108
                                            _ = v109
                                        }
                                        if (!(((42 <= -176) || ((65535 != 1) || true)) || (0 > 1))) {
                                            if false {
                                                var v111 float32 = 3.14159
                                                var v112 uint8 = 79
                                                _ = v111
                                                _ = v112
                                            }
                                            p0 = v106
                                        }
                                        _ = fi107
                                    }
                                    for fi113 := 0; fi113 < 5; fi113++ {
                                        if true {
                                            for fi114 := 0; fi114 < 13; fi114++ {
                                                v6 = append(v6, (fi113 << ((65535) & 63)))
                                                _ = fi114
                                            }
                                            fmt.Println(p0, v106, fi113)
                                        }
                                        if true {
                                            v6 = append(v6, -1)
                                        } else {
                                            p0 = ""
                                            _ = len(v6)
                                            var v115 int32 = -946
                                            _ = v115
                                        }
                                        switch fi113 {
                                        case 4294967295:
                                            var v116 uint16 = 462
                                            _ = v116
                                        case -176:
                                            v117 := GenS16{F0: false, F1: 127, F2: -515}
                                            _ = v117
                                        case 127:
                                            for fi118 := 0; fi118 < 3; fi118++ {
                                                p0 = "a"
                                                _ = len(v6)
                                                v106 = (p0 + p0)
                                                _ = fi118
                                            }
                                        default:
                                            fmt.Println(fi113)
                                        }
                                        _ = fi113
                                    }
                                    _ = v106
                                }
                            default:
                                v119 := GenS5{F0: 100}
                                _ = v119
                            }
                        }
                    case -523:
                        p0 = p0
                    }
                    _ = v76
                }
            }
            _ = v12
        }
        _ = v7
        _ = v8
    }
    v120 := GenS20{F0: 1000, F1: 18446744073709551615, F2: 127}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v120
    _ = v6
    return 65535
}
func genFn1(p0 int, p1 string, p2 bool) uint16 {
    fmt.Println(p0, p1, p0)
    for fi121 := 0; fi121 < 5; fi121++ {
        v122 := []int{}
        p1 = ((p1 + ("" + "foo bar")) + "\n\t")
        p2 = p2
        _ = fi121
        _ = v122
    }
    v123 := GenS2{F0: (-2.25 <= float32((0.1 * 0.0))), F1: int8(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v123
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 243
    v3 := []string{("" + (("\n\t" + "hello") + "\n\t")), "\n\t"}
    var v4 float32 = (1e10 + float32(3.14159))
    v5 := make(map[uint32]uint32)
    if true {
        v2 = v2
    }
    v4 = v4
    var v124 float32 = 3.14159
    var v125 string = "\n\t"
    v124 -= v124
    v2 -= (v2 - v2)
    if false {
        var v126 float32 = float32(-2.25)
        _ = v126
    }
    switch 32767 {
    case 2147483647:
        fmt.Println(v4, v125, v4)
    case 154:
        for fi127 := 0; fi127 < 7; fi127++ {
            var v128 string = v125
            fmt.Println(v125)
            _ = fi127
            _ = v128
        }
    default:
        switch v125 {
        case "":
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        default:
            var v129 uint32 = 2147483648
            _ = v129
        }
    }
    v2 += (v2 / ((0) | 1))
    v125 = v125
    var v130 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v124
    _ = v125
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
