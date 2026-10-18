package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 uint32; F1 int; F2 int32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int16; F1 int; F2 string }
type GenS12 struct { F0 uint8; F1 int8; F2 uint }
type GenS13 struct { F0 uint64 }
func genFn0(p0 string, p1 int32) int {
    switch p0 {
    case "foo bar":
        p0 = ((("" + (p0 + "foo bar")) + ("foo bar" + (p0 + "\n\t"))) + ("hello" + p0))
    case "a":
        v4 := []int{42, 32768, 127, 32768}
        _ = v4
    case "hello":
        fmt.Println(p1, p0)
    }
    var v5 float32 = 1.5
    _ = p0
    _ = p1
    _ = v5
    return -104
}
func genFn1(p0 uint16) uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        if true {
            for fi7 := 0; fi7 < 4; fi7++ {
                p0 = p0
                _ = fi7
            }
        } else {
            var v8 uint64 = 2
            p0 += 0
            _ = v8
        }
        if (((!true) || false) && (("" + "foo bar") < "")) {
            var v9 uint16 = (p0 + p0)
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := []int{(fi6 | -860), (fi6 ^ fi6)}
                v12 := GenS1{F0: 255, F1: 0, F2: 42}
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int64 = 1
                switch ("foo bar" + "\n\t") {
                case "a":
                    if true {
                        fmt.Println(fi13, v14, v14)
                    } else {
                        if (true || false) {
                            var v15 uint8 = 42
                            fmt.Println(p0, v9)
                            _ = v15
                        }
                    }
                case "foo bar":
                    if (false && false) {
                        fmt.Println(v14, fi13, v14)
                    } else {
                        var v16 int8 = 1
                        if (!(false || (!(255 <= 42)))) {
                            if false {
                                var v17 uint32 = 65535
                                v18 := GenS1{F0: 128, F1: (^(v14)), F2: 1000}
                                v19 := []int{(fi6 % ((fi6) | 1)), fi13, 127, fi6}
                                _ = v17
                                _ = v18
                                _ = v19
                            } else {
                                if false {
                                    switch (fi13 >> ((4294967295) & 63)) {
                                    case 282:
                                        var v20 float32 = 1.5
                                        _ = v20
                                    case 42:
                                        fmt.Println(p0, p0, v16)
                                    }
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v14)
                                    case "\n\t":
                                        v9 += (v9 << ((954) & 63))
                                        var v21 uint32 = 4294967295
                                        v22 := GenS1{F0: uint8(v21), F1: 1, F2: 1000}
                                        _ = v21
                                        _ = v22
                                    default:
                                        v9 = (v9 / ((v9) | 1))
                                        var v23 uint8 = 0
                                        v24 := []int{(fi13 >> ((255) & 63))}
                                        _ = v23
                                        _ = v24
                                    }
                                    v14 = -1
                                    switch ("foo bar" + "foo bar") {
                                    case "a":
                                        v14 -= v14
                                    case "":
                                        var v25 int64 = (v14 & v14)
                                        _ = v25
                                    }
                                    v9 -= 2
                                }
                                fmt.Println(v16, fi13, fi6)
                                v14 -= (v14 << ((255) & 63))
                            }
                        }
                        v16 = -128
                        _ = v16
                    }
                }
                var v26 float32 = (1e10 - 0.1)
                switch 221 {
                case 128:
                    if (4294967295 != 840) {
                        var v27 string = "hello"
                        var v28 uint32 = 4294967295
                        if (v27 <= "") {
                            v29 := GenS2{F0: v14, F1: v14, F2: 449}
                            v30 := GenS3{F0: 28}
                            _ = v29
                            _ = v30
                        } else {
                            fmt.Println(fi6, fi13)
                            v14 -= -259
                            v31 := []int{}
                            _ = v31
                        }
                        _ = v27
                        _ = v28
                    } else {
                        switch "hello" {
                        case "foo bar":
                            v26 = (v26 + (float32(1.5) - 0.0))
                            v14 -= v14
                            var v32 int8 = 64
                            switch ("hello" + "hello") {
                            case "a":
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v14, v9, p0)
                                }
                            default:
                                var v33 int = (fi13 % ((-1) | 1))
                                var v34 bool = false
                                if (1e10 < 0.0) {
                                    v35 := GenS4{F0: 1, F1: v34}
                                    v36 := GenS2{F0: v14, F1: 723, F2: 1}
                                    v37 := GenS4{F0: 2, F1: (!(!(255 != 18446744073709551615)))}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                _ = v33
                                _ = v34
                            }
                            var v38 uint8 = 128
                            _ = v32
                            _ = v38
                        }
                    }
                case -62:
                    switch "a" {
                    case "foo bar":
                        var v39 int16 = 0
                        _ = v39
                    case "hello":
                        var v40 int = (fi13 | 127)
                        _ = v40
                    }
                case 9223372036854775807:
                    var v41 uint = 647
                    _ = v41
                }
                v42 := []int{fi13, fi6}
                p0 = (p0 + p0)
                _ = fi13
                _ = v14
                _ = v26
                _ = v42
            }
            _ = v9
        }
        _ = fi6
    }
    if (true || true) {
        for fi43 := 0; fi43 < 18; fi43++ {
            p0 = 0
            for fi44 := 0; fi44 < 1; fi44++ {
                v45 := []int{(fi44 << ((65535) & 63)), (-(fi43)), (-(fi44))}
                var v46 uint32 = 578
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = fi43
        }
        fmt.Println(p0)
    }
    var v47 int16 = 42
    fmt.Println(p0, p0)
    _ = p0
    _ = v47
    return 1
}

func main() {
    var v1 uint8 = 0
    var v2 uint8 = v1
    v3 := []int{127, -545, 128}
    fmt.Println(v1, v1)
    var v48 float32 = (1e10 / 0.1)
    if false {
        _ = len(v3)
        var v49 bool = ((1 != 1) || false)
        _ = v49
    } else {
        var v50 int16 = 1000
        switch -537 {
        case -410:
            v3 = append(v3, 9223372036854775807)
        }
        _ = v50
    }
    v51 := GenS4{F0: 42, F1: false}
    for fi52 := 0; fi52 < 18; fi52++ {
        for fi53 := 0; fi53 < 12; fi53++ {
            var v54 uint32 = 0
            fmt.Println(v54)
            _ = fi53
            _ = v54
        }
        v2 -= v1
        if true {
            v1 = v2
            v3 = append(v3, fi52)
            for fi55 := 0; fi55 < 12; fi55++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (32767 == 2) {
                    v3 = append(v3, fi52)
                }
                _ = fi55
            }
        } else {
            var v56 float64 = 0.0
            _ = v56
        }
        _ = fi52
    }
    v48 += ((3.14159 + v48) + 1.5)
    var v57 int = 32767
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v57 -= (-(v57))
    if ((false || (918 != -546)) && (!(!(!(18446744073709551615 > 0))))) {
        v58 := GenS5{F0: ("a" + ("hello" + ("foo bar" + (("hello" + "a") + "hello")))), F1: -1}
        _ = len(v3)
        switch "a" {
        case "":
            if (true && false) {
                v3 = append(v3, genFn0("\n\t", -1))
                var v59 int8 = 127
                if true {
                    if false {
                        v59 += 0
                    } else {
                        _ = len(v3)
                        v60 := GenS5{F0: "hello", F1: 0}
                        if false {
                            v61 := GenS3{F0: 2}
                            if (!true) {
                                if (!true) {
                                    v3 = append(v3, (v57 ^ (^(v57))))
                                    for fi62 := 0; fi62 < 12; fi62++ {
                                        if (4294967295 == 65535) {
                                            v63 := GenS6{F0: 137, F1: 42}
                                            v64 := GenS4{F0: 4294967295, F1: false}
                                            _ = v63
                                            _ = v64
                                        } else {
                                            for fi65 := 0; fi65 < 6; fi65++ {
                                                switch "" {
                                                case "":
                                                    v3 = append(v3, genFn0("a", int32(fi62)))
                                                }
                                                _ = fi65
                                            }
                                            switch ("\n\t" + "\n\t") {
                                            case "\n\t":
                                                var v66 uint = genFn1(70)
                                                _ = v66
                                            }
                                        }
                                        _ = fi62
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            } else {
                                for fi67 := 0; fi67 < 19; fi67++ {
                                    var v68 bool = false
                                    var v69 string = ("hello" + "hello")
                                    v1 -= 128
                                    _ = fi67
                                    _ = v68
                                    _ = v69
                                }
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v70 uint8 = v2
                                    _ = v70
                                }
                                var v71 int64 = 40
                                _ = v71
                            }
                            _ = v61
                        }
                        _ = v60
                    }
                }
                _ = v59
            } else {
                for fi72 := 0; fi72 < 12; fi72++ {
                    var v73 uint8 = (^(v1))
                    fmt.Println(v73, v48, v2)
                    if ((^(v73)) == (v1 << ((42) & 63))) {
                        var v74 bool = true
                        var v75 uint32 = 2147483648
                        for fi76 := 0; fi76 < 11; fi76++ {
                            if (965 == genFn1(1)) {
                                v75 = (v75 ^ (^(v75)))
                                v77 := GenS3{F0: 2}
                                var v78 float32 = 1e10
                                _ = v77
                                _ = v78
                            }
                            switch "a" {
                            case "\n\t":
                                if v74 {
                                    if (!true) {
                                        v79 := GenS7{F0: 1, F1: v75}
                                        v80 := GenS8{F0: 42, F1: 100}
                                        v48 -= float32(0.1)
                                        v75 = 2147483648
                                        _ = v79
                                        _ = v80
                                    } else {
                                        if (false && (18446744073709551615 > 4294967295)) {
                                            if len(v3) > 0 {
                                                v3[0] = (-(fi72))
                                            }
                                            v57 = genFn0(("foo bar" + ""), -2147483648)
                                            v74 = v74
                                        } else {
                                            var v81 int64 = -959
                                            if v74 {
                                                if (123 <= 42) {
                                                    var v82 int8 = 100
                                                    _ = len(v3)
                                                    v74 = (genFn1(954) <= 18446744073709551615)
                                                    var v83 uint8 = (v2 << ((42) & 63))
                                                    _ = v82
                                                    _ = v83
                                                } else {
                                                    if (227 < (^(v81))) {
                                                        v73 += (v1 - v2)
                                                    }
                                                    switch v57 {
                                                    case 9223372036854775807:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            fmt.Println(v75)
                                                        case "a":
                                                            if v74 {
                                                                if ((v81 * (v81 - v81)) < (v81 + (v81 / (((v81 * v81)) | 1)))) {
                                                                    var v84 bool = ((128 > (v57 << ((898) & 63))) || true)
                                                                    fmt.Println(v2, fi72, v73)
                                                                    v57 = (^(v57))
                                                                    _ = v84
                                                                }
                                                                v57 = -141
                                                            }
                                                        default:
                                                            var v85 float64 = ((1.5 + (0.0 / (((0.1 * 0.0) / (3.14159 * -2.25)) / float64((v48 + v48))))) / 1.0)
                                                            _ = v85
                                                        }
                                                    default:
                                                        v3 = append(v3, (v57 | fi72))
                                                    }
                                                }
                                                var v86 int64 = 695
                                                if (false && v74) {
                                                    var v87 bool = v74
                                                    var v88 int8 = -128
                                                    _ = v87
                                                    _ = v88
                                                } else {
                                                    switch "" {
                                                    case "":
                                                        v89 := GenS1{F0: 127, F1: (v81 + (v81 >> ((466) & 63))), F2: 42}
                                                        _ = v89
                                                    default:
                                                        if (((-2.25 * ((0.1 - -2.25) * 3.14159)) * float64(((-2.25 / v48) + float32(-2.25)))) == 0.1) {
                                                            var v90 int32 = 2147483647
                                                            _ = v90
                                                        } else {
                                                            v91 := GenS9{F0: v75, F1: 349, F2: -962}
                                                            _ = v91
                                                        }
                                                    }
                                                }
                                                _ = v86
                                            }
                                            _ = len(v3)
                                            _ = v81
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = (^(v57))
                                        }
                                        var v92 uint8 = v1
                                        _ = v92
                                    }
                                } else {
                                    if ((220 != 1) && (-601 < 32767)) {
                                        fmt.Println(v2, v1)
                                        if (18446744073709551615 == 4294967295) {
                                            _ = len(v3)
                                            v57 += -751
                                            switch "hello" {
                                            case "a":
                                                var v93 float32 = 1e10
                                                _ = v93
                                            case "foo bar":
                                                var v94 float64 = -2.25
                                                _ = v94
                                            case "":
                                                v48 = ((1e10 / 0.1) - v48)
                                            default:
                                                v48 = ((v48 * float32(0.1)) + 1e10)
                                            }
                                        } else {
                                            v95 := GenS5{F0: "\n\t", F1: 987}
                                            if false {
                                                if v74 {
                                                    var v96 int32 = 42
                                                    var v97 string = ((("foo bar" + ("hello" + ("\n\t" + "a"))) + "hello") + "a")
                                                    _ = v96
                                                    _ = v97
                                                }
                                                fmt.Println(v57)
                                                v75 = 65535
                                            } else {
                                                v2 = 1
                                                if false {
                                                    var v98 uint32 = v75
                                                    _ = v98
                                                } else {
                                                    var v99 uint64 = 2
                                                    v100 := GenS10{F0: (v74 && false)}
                                                    v2 += v2
                                                    fmt.Println(v74, v75, fi72)
                                                    if (0 == (v99 % ((4294967295) | 1))) {
                                                        switch (fi72 + (-(v57))) {
                                                        case 113:
                                                            var v101 bool = v74
                                                            _ = v101
                                                        default:
                                                            if true {
                                                                v102 := GenS11{F0: 32767, F1: 611, F2: "foo bar"}
                                                                fmt.Println(v75, fi72)
                                                                _ = v102
                                                            }
                                                        }
                                                        v57 += genFn0(((("foo bar" + "\n\t") + ("" + "hello")) + "foo bar"), -1)
                                                    }
                                                    _ = v100
                                                    _ = v99
                                                }
                                                v75 -= v75
                                            }
                                            var v103 uint16 = uint16(v75)
                                            _ = v103
                                            _ = v95
                                        }
                                        if ((v2 >> ((650) & 63)) != (v2 * v73)) {
                                            if v74 {
                                                var v104 int = 127
                                                _ = v104
                                            }
                                            if true {
                                                v57 -= 2147483647
                                                v75 = (^(v75))
                                                v57 -= (fi72 * genFn0((("a" + "\n\t") + "a"), -1))
                                            } else {
                                                v105 := GenS5{F0: ("\n\t" + ("a" + "a")), F1: 2147483647}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if v74 {
                                                    v48 += ((0.1 * (float32(3.14159) - 1.5)) - -2.25)
                                                }
                                                _ = v105
                                            }
                                            if (-1 >= 679) {
                                                v106 := GenS12{F0: (v73 / ((v2) | 1)), F1: 42, F2: 0}
                                                switch -9223372036854775808 {
                                                case 32767:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case 584:
                                                    v1 = 214
                                                case 9223372036854775807:
                                                    v2 += (v2 + 42)
                                                default:
                                                    if (0 > -994) {
                                                        if len(v3) > 0 {
                                                            v3[0] = genFn0("", 0)
                                                        }
                                                        var v107 int16 = -32768
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v107
                                                    }
                                                }
                                                v108 := GenS13{F0: 18446744073709551615}
                                                v109 := GenS5{F0: "\n\t", F1: -391}
                                                _ = v106
                                                _ = v108
                                                _ = v109
                                            } else {
                                                var v110 int16 = 134
                                                fmt.Println(v2, v57, fi72)
                                                var v111 uint8 = 255
                                                _ = v110
                                                _ = v111
                                            }
                                        }
                                    }
                                }
                            case "hello":
                                var v112 int64 = -9223372036854775808
                                _ = v112
                            case "":
                                var v113 int = 128
                                _ = v113
                            default:
                                v74 = (v2 <= v2)
                            }
                            _ = fi76
                        }
                        _ = v74
                        _ = v75
                    } else {
                        fmt.Println(fi72, v48)
                    }
                    _ = fi72
                    _ = v73
                }
                switch v57 {
                case -2147483648:
                    for fi114 := 0; fi114 < 7; fi114++ {
                        var v115 uint8 = v2
                        fmt.Println(v48, v115, v48)
                        _ = fi114
                        _ = v115
                    }
                default:
                    v116 := GenS7{F0: -32, F1: 65535}
                    _ = v116
                }
                v57 += (-(v57))
            }
        case "a":
            var v117 uint8 = 2
            _ = v117
        default:
            v1 = 212
        }
        _ = v58
    }
    switch "foo bar" {
    case "foo bar":
        v57 += 1
    case "\n\t":
        if (1.5 > -2.25) {
            var v118 uint8 = v1
            _ = v118
        } else {
            for fi119 := 0; fi119 < 5; fi119++ {
                var v120 int32 = 42
                _ = fi119
                _ = v120
            }
            v48 = 0.1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    default:
        fmt.Println(v48)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v51
    _ = v57
}
