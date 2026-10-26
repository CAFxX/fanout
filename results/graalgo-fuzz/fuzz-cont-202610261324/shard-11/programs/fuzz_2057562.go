package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS12 struct { F0 int8; F1 bool; F2 uint32 }
type GenS13 struct { F0 uint; F1 int64; F2 uint64 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int8 }
func genFn0(p0 int) bool {
    v7 := GenS1{F0: 2, F1: p0}
    p0 -= (p0 >> ((0) & 63))
    v8 := GenS1{F0: 1, F1: p0}
    _ = p0
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{51, 65535, 2147483648, 0}
    var v3 int = 32768
    var v4 int32 = -1
    v5 := []int{v3, v3}
    var v6 uint8 = 128
    var v9 uint = 4294967295
    v10 := GenS1{F0: (^(v6)), F1: 96}
    if false {
        if (!true) {
            var v11 string = ((("" + (("\n\t" + ("\n\t" + ("" + "hello"))) + "foo bar")) + ("" + (("hello" + ((((("foo bar" + "\n\t") + "a") + "\n\t") + "foo bar") + (((("" + ("hello" + ("" + "\n\t"))) + "\n\t") + ("a" + "hello")) + "\n\t"))) + ""))) + (("a" + (("foo bar" + ("\n\t" + "hello")) + ("a" + ("\n\t" + ("foo bar" + "hello"))))) + "hello"))
            var v12 int8 = 42
            _ = v11
            _ = v12
        } else {
            v4 = v4
        }
    }
    switch v3 {
    case -709:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v9, v9)
    v13 := GenS2{F0: 275, F1: -1, F2: v6}
    var v14 uint64 = 981
    fmt.Println(v1)
    switch 9223372036854775807 {
    case 32767:
        v15 := GenS1{F0: v6, F1: 98}
        _ = v15
    case 1:
        v16 := GenS2{F0: (v9 << ((v9) & 63)), F1: 100, F2: 127}
        _ = v16
    case -2147483648:
        if (78 != -673) {
            var v17 string = ("foo bar" + "foo bar")
            var v18 float64 = (1e10 - (((1e10 + (1.5 * 0.0)) * 1e10) / (-2.25 / (0.1 + 0.0))))
            _ = v17
            _ = v18
        } else {
            var v19 int64 = 15
            var v20 string = "\n\t"
            _ = v19
            _ = v20
        }
    default:
        var v21 int16 = v1
        _ = v21
    }
    if (!false) {
        v6 = v6
        v22 := GenS1{F0: 198, F1: 726}
        _ = v22
    }
    if ((42 == -1) || (false || ((-2.25 + float32(float64(-2.25))) <= (-2.25 * 0.1)))) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if false {
        v6 += 1
    } else {
        v23 := GenS3{F0: v14, F1: v4}
        fmt.Println(v6, v9, v14)
        fmt.Println(v1, v3)
        _ = v23
    }
    if true {
        switch v3 {
        case -489:
            var v24 uint = (^(v9))
            _ = v24
        }
        var v25 bool = true
        fmt.Println(v3)
        _ = v25
    } else {
        if true {
            for fi26 := 0; fi26 < 2; fi26++ {
                v3 = (fi26 * (-(v3)))
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 uint16 = 1
                    fmt.Println(v6, v4)
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 int16 = 1
            v3 += v3
            _ = fi29
            _ = v30
        }
    }
    var v31 uint64 = (^(v14))
    v32 := GenS2{F0: (v9 << ((65535) & 63)), F1: -1, F2: (v6 % (((v6 - (v6 << (((v9 & (v9 | 712))) & 63)))) | 1))}
    if (((^(v1)) >= (v1 ^ int16(v9))) || true) {
        if ((-2.25 >= (-2.25 + (float32(0.1) - float32(1e10)))) && false) {
            var v33 bool = (!(-1 == (v4 >> (((^(v9))) & 63))))
            if v33 {
                var v34 uint = 42
                var v35 float64 = float64(((3.14159 - -2.25) * 3.14159))
                _ = v34
                _ = v35
            } else {
                v36 := GenS4{F0: 132, F1: 9223372036854775807}
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 int64 = 899
                    v4 -= v4
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
            v39 := GenS5{F0: 1, F1: 2147483648}
            _ = v33
            _ = v39
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            v41 := GenS3{F0: v31, F1: v4}
            v6 = v6
            _ = fi40
            _ = v41
        }
        v6 += (v6 % ((v6) | 1))
    } else {
        if (((uint64(v9) != 4294967295) && (false || (true && (!true)))) && false) {
            v31 = v31
            for fi42 := 0; fi42 < 11; fi42++ {
                if true {
                    v43 := GenS1{F0: v6, F1: (fi42 ^ int(v4))}
                    _ = v43
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v9, v1)
                    case "a":
                        fmt.Println(v6)
                    case "foo bar":
                        for fi44 := 0; fi44 < 13; fi44++ {
                            v45 := GenS4{F0: 885, F1: 606}
                            v46 := GenS6{F0: 166}
                            var v47 uint32 = 525
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                    }
                    fmt.Println(v3, v1, v1)
                    switch (v3 * 287) {
                    case 2147483647:
                        v48 := GenS7{F0: "hello", F1: 42}
                        _ = v48
                    case -9223372036854775808:
                        v49 := GenS8{F0: v6, F1: 836, F2: 42}
                        _ = v49
                    case -603:
                        var v50 uint64 = v14
                        _ = v50
                    }
                }
                v31 = 573
                if true {
                    for fi51 := 0; fi51 < 15; fi51++ {
                        switch "foo bar" {
                        case "a":
                            switch "" {
                            case "foo bar":
                                v52 := GenS5{F0: uint16(v9), F1: 2}
                                _ = v52
                            default:
                                if (!((!false) && false)) {
                                    fmt.Println(v3, fi51, v9)
                                    switch (fi42 & -929) {
                                    case 42:
                                        var v53 int8 = 100
                                        _ = v53
                                    case 886:
                                        fmt.Println(v9)
                                    case -9223372036854775808:
                                        switch "a" {
                                        case "":
                                            var v54 float32 = (-2.25 * 1.5)
                                            _ = v54
                                        default:
                                            switch fi42 {
                                            case -111:
                                                if ((657 <= 0) && false) {
                                                    switch fi42 {
                                                    case 4294967295:
                                                        if false {
                                                            var v55 uint = (v9 >> ((v9) & 63))
                                                            var v56 uint = (v55 * uint(v31))
                                                            if true {
                                                                if (-126 < 94) {
                                                                    v4 -= int32(v14)
                                                                }
                                                            }
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v56 = 0
                                                            case "a":
                                                                var v57 float64 = 0.1
                                                                _ = v57
                                                            case "foo bar":
                                                                v3 += 128
                                                            default:
                                                                fmt.Println(v1, v1)
                                                            }
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    case 43:
                                                        v58 := GenS8{F0: v6, F1: 152, F2: 32767}
                                                        _ = v58
                                                    default:
                                                        v59 := GenS1{F0: (v6 & (v6 * (v6 & uint8(v9)))), F1: (fi42 << ((515) & 63))}
                                                        _ = v59
                                                    }
                                                    var v60 float64 = (((-2.25 / 1e10) / (-2.25 * 1e10)) + 1e10)
                                                    var v61 uint8 = (^(v6))
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    var v62 string = "hello"
                                                    if (43 != 42) {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v63 uint16 = 197
                                                        _ = v63
                                                    } else {
                                                        v2 = append(v2, 65535)
                                                    }
                                                    _ = v62
                                                }
                                            case 817:
                                                if len(v2) > 0 {
                                                    v2[0] = 2
                                                }
                                                v4 = v4
                                            case -2147483648:
                                                var v64 string = ("foo bar" + "a")
                                                _ = v64
                                            }
                                        }
                                    }
                                    var v65 bool = false
                                    _ = v65
                                } else {
                                    fmt.Println(v4, v31)
                                    fmt.Println(v31, v31)
                                    switch "hello" {
                                    case "a":
                                        v66 := GenS4{F0: 0, F1: 1}
                                        _ = v66
                                    case "":
                                        switch 644 {
                                        case 9223372036854775807:
                                            fmt.Println(v3)
                                        case 1:
                                            var v67 int64 = 389
                                            _ = v67
                                        default:
                                            if (v4 <= int32(v14)) {
                                                switch (fi51 | fi51) {
                                                case 42:
                                                    v4 += v4
                                                }
                                                if true {
                                                    fmt.Println(v1, v1, fi42)
                                                    v31 = (v31 - (^(v14)))
                                                }
                                                _ = len(v5)
                                            }
                                        }
                                        if (("hello" + "") >= "a") {
                                            if (!true) {
                                                v9 = (v9 >> ((v9) & 63))
                                                var v68 int = -9223372036854775808
                                                v14 = (v14 * v31)
                                                _ = v68
                                            } else {
                                                v3 = v3
                                                var v69 string = "hello"
                                                _ = v69
                                            }
                                        }
                                    case "foo bar":
                                        switch int(v9) {
                                        case -2147483648:
                                            if false {
                                                if true {
                                                    v14 -= 334
                                                    var v70 string = ""
                                                    _ = v70
                                                } else {
                                                    switch (fi42 << ((v9) & 63)) {
                                                    case 441:
                                                        if false {
                                                            var v71 int8 = 71
                                                            var v72 uint64 = v31
                                                            if false {
                                                                var v73 uint16 = 581
                                                                v74 := GenS9{F0: "hello", F1: -434}
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                var v75 uint32 = 664
                                                                _ = v75
                                                            }
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    default:
                                        v9 -= uint(v14)
                                    }
                                }
                            }
                        default:
                            v6 = 42
                        }
                        var v76 int = v3
                        switch fi51 {
                        case -9223372036854775808:
                            var v77 float64 = 0.1
                            _ = v77
                        case -631:
                            switch fi51 {
                            case -254:
                                v9 += (v9 + 1)
                            case 4294967295:
                                switch (("hello" + "hello") + (("hello" + "foo bar") + "a")) {
                                case "\n\t":
                                    if false {
                                        var v78 uint8 = (v6 | v6)
                                        if false {
                                            if len(v5) > 0 {
                                                v5[0] = v3
                                            }
                                            v3 = -200
                                        } else {
                                            var v79 int = fi42
                                            switch "" {
                                            case "":
                                                if (v76 != (^(v79))) {
                                                    if ((-(v4)) > v4) {
                                                        v9 = (v9 / (((v9 ^ 255)) | 1))
                                                        v6 += v6
                                                    }
                                                } else {
                                                    var v80 float64 = 1.5
                                                    _ = v80
                                                }
                                            case "a":
                                                var v81 int = (fi42 - (^(fi42)))
                                                _ = v81
                                            case "foo bar":
                                                v3 = fi42
                                            }
                                            fmt.Println(v3)
                                            _ = v79
                                        }
                                        v82 := GenS1{F0: v78, F1: (v3 % (((^(fi42))) | 1))}
                                        var v83 uint16 = 255
                                        _ = v78
                                        _ = v82
                                        _ = v83
                                    }
                                default:
                                    var v84 uint64 = (v14 | (v14 * (v14 % (((v31 * (v14 + v31))) | 1))))
                                    _ = v84
                                }
                            }
                        case 127:
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v14 = v31
                                v85 := GenS10{F0: v1, F1: 2}
                                _ = v85
                            }
                        default:
                            v86 := GenS11{F0: uint32(v31), F1: 42, F2: 171}
                            _ = v86
                        }
                        _ = fi51
                        _ = v76
                    }
                }
                _ = fi42
            }
            switch v3 {
            case 482:
                if len(v2) > 0 {
                    v2[0] = 987
                }
            case 0:
                v6 = (v6 | v6)
            }
        } else {
            v4 = (v4 & (v4 * (v4 - (v4 | -2147483648))))
            if ("hello" == "") {
                var v87 float64 = ((0.0 - float64(3.14159)) / 1e10)
                _ = v87
            }
        }
        switch "" {
        case "":
            for fi88 := 0; fi88 < 14; fi88++ {
                _ = len(v2)
                if (!false) {
                    var v89 uint32 = 4294967295
                    v90 := GenS1{F0: v6, F1: (fi88 - v3)}
                    if false {
                        fmt.Println(v3, v4, fi88)
                        switch (-(fi88)) {
                        case -2147483648:
                            switch ("" + "foo bar") {
                            case "a":
                                for fi91 := 0; fi91 < 11; fi91++ {
                                    if (18446744073709551615 != (v9 ^ v9)) {
                                        v89 = (^(v89))
                                    } else {
                                        if ((!(!true)) && false) {
                                            switch (v3 + (fi91 / ((fi88) | 1))) {
                                            case 32768:
                                                v92 := GenS10{F0: (v1 * (-(v1))), F1: 65535}
                                                _ = v92
                                            case 128:
                                                var v93 int = (fi88 >> (((v9 | v9)) & 63))
                                                _ = v93
                                            case 32767:
                                                fmt.Println(v3, v6, fi91)
                                            }
                                            v94 := GenS12{F0: 100, F1: false, F2: v89}
                                            _ = v94
                                        }
                                        switch (^(fi91)) {
                                        case 32767:
                                            if true {
                                                fmt.Println(v4, v89, v6)
                                            } else {
                                                var v95 int16 = v1
                                                _ = v95
                                            }
                                        case -9223372036854775808:
                                            var v96 int64 = 9223372036854775807
                                            switch "a" {
                                            case "a":
                                                v31 += 70
                                            }
                                            _ = v96
                                        default:
                                            v2 = append(v2, uint32(v14))
                                            switch (fi91 & (fi88 * -9223372036854775808)) {
                                            case -209:
                                                if false {
                                                    if len(v2) > 0 {
                                                        v2[0] = (v89 & 571)
                                                    }
                                                    _ = len(v2)
                                                    var v97 float32 = float32((3.14159 - (1e10 / (0.1 / -2.25))))
                                                    var v98 uint32 = (v89 ^ (^(v89)))
                                                    switch 4294967295 {
                                                    case 9223372036854775807:
                                                        v99 := GenS10{F0: 32767, F1: 255}
                                                        _ = v99
                                                    }
                                                    v100 := GenS13{F0: (^(v9)), F1: 662, F2: 857}
                                                    if (false && false) {
                                                        var v101 bool = true
                                                        if len(v5) > 0 {
                                                            v5[0] = 9223372036854775807
                                                        }
                                                        v102 := GenS14{F0: v101}
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                    _ = v100
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    var v103 int8 = -90
                                                    if (!false) {
                                                        var v104 string = "\n\t"
                                                        _ = len(v5)
                                                        _ = v104
                                                    }
                                                    _ = v103
                                                }
                                            default:
                                                v14 = (v14 ^ 0)
                                            }
                                            v105 := GenS15{F0: -123}
                                            _ = v105
                                        }
                                        var v106 uint8 = (v6 * (v6 >> (((v9 << ((v9) & 63))) & 63)))
                                        _ = v106
                                    }
                                    v14 -= v14
                                    var v107 int32 = -2147483648
                                    _ = fi91
                                    _ = v107
                                }
                            case "hello":
                                var v108 float32 = 0.0
                                _ = v108
                            case "":
                                v3 = fi88
                            }
                        case -1:
                            var v109 float64 = (0.0 - (1e10 * (1e10 / (((1e10 / 0.1) / 1.0) - (0.0 + (1e10 - 1e10))))))
                            _ = v109
                        case 32767:
                            v110 := GenS7{F0: "\n\t", F1: v89}
                            _ = v110
                        default:
                            fmt.Println(v3, fi88, fi88)
                        }
                    }
                    _ = v89
                    _ = v90
                }
                var v111 int64 = int64(v14)
                _ = fi88
                _ = v111
            }
        default:
            v31 += v31
        }
    }
    var v112 int16 = 1
    for fi113 := 0; fi113 < 14; fi113++ {
        if true {
            v3 = v3
            var v114 string = "a"
            v31 += uint64(v9)
            _ = v114
        }
        fmt.Println(v31, v4)
        v115 := GenS6{F0: 0}
        _ = fi113
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v112
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
