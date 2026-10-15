package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS14 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS15 struct { F0 int8; F1 uint32 }
type GenS16 struct { F0 bool; F1 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint64) int64 {
    for fi4 := 0; fi4 < 15; fi4++ {
        switch fi4 {
        case 1:
            v5 := []int{}
            _ = v5
        case 220:
            p1 += 450
        case 0:
            switch ("\n\t" + "") {
            case "foo bar":
                fmt.Println(p2, p1)
            case "hello":
                if true {
                    var v6 int16 = 827
                    if (127 < 128) {
                        v6 += v6
                        var v7 bool = true
                        _ = v7
                    }
                    _ = v6
                }
            }
        default:
            p2 -= p0
        }
        var v8 int16 = 42
        _ = fi4
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint64 = 215
        var v11 uint = 42
        v10 -= (p0 / (((v10 % (((^(v10))) | 1))) | 1))
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 bool, p1 int64) int16 {
    v12 := GenS1{F0: -79, F1: 4294967295, F2: 2147483648}
    v13 := GenS2{F0: 250, F1: p0}
    v14 := []int{-2147483648, -379}
    var v15 int16 = 32767
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return v15
}

func main() {
    var v1 uint64 = 42
    var v2 string = ("a" + (("hello" + "\n\t") + ("" + "foo bar")))
    var v3 int32 = 2147483647
    v16 := []int{}
    switch v2 {
    case "hello":
        switch "a" {
        case "foo bar":
            _ = len(v16)
        case "hello":
            if (0 > 42) {
                v17 := GenS2{F0: 2, F1: (!(1 != 42))}
                _ = v17
            } else {
                fmt.Println(v2)
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 float64 = 3.14159
                    fmt.Println(v3, v2, v19)
                    v16 = append(v16, -158)
                    _ = fi18
                    _ = v19
                }
                if (v1 == (v1 % ((698) | 1))) {
                    var v20 float32 = float32(-2.25)
                    for fi21 := 0; fi21 < 15; fi21++ {
                        var v22 string = v2
                        var v23 bool = (v1 > (v1 + v1))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v3 += (v3 << ((65535) & 63))
                    _ = v20
                }
            }
        }
    case "a":
        var v24 uint8 = 2
        _ = v24
    case "foo bar":
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    v25 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: genFn0((v1 >> ((568) & 63)), 255, (v1 | 730))}
    fmt.Println(v3, v2, v1)
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS2{F0: 128, F1: true}
        if true {
            var v28 int8 = 127
            v29 := GenS3{F0: v1, F1: 2, F2: -9223372036854775808}
            v3 += 42
            _ = v28
            _ = v29
        }
        fmt.Println(v1, v2, fi26)
        _ = fi26
        _ = v27
    }
    if (true || ((108 != v1) && true)) {
        if true {
            v30 := GenS4{F0: v2, F1: v3}
            _ = v30
        } else {
            if (!(false || (true || (987 == 4294967295)))) {
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 float32 = ((1e10 / (3.14159 - (3.14159 - -2.25))) - (0.1 * 3.14159))
                    v32 = 0.0
                    _ = fi31
                    _ = v32
                }
                if (314 > 437) {
                    switch 32768 {
                    case 250:
                        if ((true && false) && (3.14159 >= float64((0.0 / 1.0)))) {
                            var v33 int32 = v3
                            _ = v33
                        } else {
                            var v34 int8 = -14
                            v3 -= v3
                            var v35 bool = ((false || true) && (-2.25 == -2.25))
                            _ = v34
                            _ = v35
                        }
                    case -2147483648:
                        switch -51 {
                        case -9223372036854775808:
                            if (948 <= 42) {
                                switch ((v2 + "") + (v2 + (v2 + ((v2 + v2) + ("foo bar" + (("hello" + (v2 + v2)) + (v2 + v2))))))) {
                                case "\n\t":
                                    _ = len(v16)
                                case "foo bar":
                                    if (-9223372036854775808 >= 42) {
                                        v2 = (v2 + "hello")
                                        v36 := GenS4{F0: "hello", F1: v3}
                                        switch 32767 {
                                        case 1:
                                            for fi37 := 0; fi37 < 7; fi37++ {
                                                if (!(((((float64(float32(float64(1.5))) + 1.5) + 1.5) * -2.25) + 3.14159) != (float64(float32((0.0 / 0.1))) * -2.25))) {
                                                    fmt.Println(v2)
                                                    var v38 bool = false
                                                    _ = v38
                                                } else {
                                                    switch fi37 {
                                                    case 32768:
                                                        v2 = "foo bar"
                                                    case 32767:
                                                        for fi39 := 0; fi39 < 19; fi39++ {
                                                            switch fi39 {
                                                            case 107:
                                                                fmt.Println(v1)
                                                            case -196:
                                                                var v40 bool = (-1 < -1)
                                                                _ = v40
                                                            case 566:
                                                                if len(v16) > 0 {
                                                                    v16[0] = 0
                                                                }
                                                            default:
                                                                v3 = v3
                                                            }
                                                            var v41 uint16 = 2
                                                            _ = fi39
                                                            _ = v41
                                                        }
                                                    case 42:
                                                        for fi42 := 0; fi42 < 17; fi42++ {
                                                            if (1 != (v1 - v1)) {
                                                                if false {
                                                                    fmt.Println(fi37)
                                                                    v43 := GenS2{F0: 42, F1: (568 <= 665)}
                                                                    _ = v43
                                                                } else {
                                                                    v1 -= v1
                                                                    v16 = append(v16, fi42)
                                                                }
                                                                v44 := GenS3{F0: 42, F1: 65535, F2: 56}
                                                                _ = v44
                                                            }
                                                            var v45 int32 = 1
                                                            var v46 int = fi42
                                                            var v47 int16 = 1000
                                                            v45 += 1
                                                            _ = fi42
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    default:
                                                        switch v2 {
                                                        case "a":
                                                            if false {
                                                                for fi48 := 0; fi48 < 20; fi48++ {
                                                                    var v49 uint16 = 303
                                                                    if (-305 < -604) {
                                                                        switch v2 {
                                                                        case "\n\t":
                                                                            switch v2 {
                                                                            case "a":
                                                                                fmt.Println(v1, v2)
                                                                            case "hello":
                                                                                var v50 uint64 = (^(v1))
                                                                                _ = v50
                                                                            default:
                                                                                v49 -= (v49 - (v49 + (v49 % ((v49) | 1))))
                                                                            }
                                                                        }
                                                                        fmt.Println(v3, fi37)
                                                                    } else {
                                                                        var v51 uint8 = 1
                                                                        v51 += 1
                                                                        fmt.Println(v51)
                                                                        _ = v51
                                                                    }
                                                                    v52 := GenS2{F0: 1, F1: (genFn1(((1e10 > 3.14159) && (((1.5 - -2.25) == (1e10 * float32((0.0 * (0.1 * (((float64(0.0) * 3.14159) + -2.25) / 1.0)))))) || (!(1 >= 57)))), genFn0((v1 | (^(v1))), v49, (v1 + (v1 % (((v1 / (((v1 - 2)) | 1))) | 1))))) != -771)}
                                                                    v16 = append(v16, 1)
                                                                    v16 = append(v16, (fi48 * fi37))
                                                                    _ = fi48
                                                                    _ = v49
                                                                    _ = v52
                                                                }
                                                                v16 = append(v16, fi37)
                                                                var v53 int8 = 62
                                                                _ = v53
                                                            } else {
                                                                for fi54 := 0; fi54 < 14; fi54++ {
                                                                    v3 -= v3
                                                                    _ = fi54
                                                                }
                                                                v55 := GenS2{F0: 213, F1: true}
                                                                for fi56 := 0; fi56 < 10; fi56++ {
                                                                    if true {
                                                                        v1 = v1
                                                                    } else {
                                                                        var v57 int = fi37
                                                                        v58 := GenS5{F0: v3, F1: 0}
                                                                        v57 -= (-(fi37))
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                    v59 := GenS3{F0: (v1 / (((^(v1))) | 1)), F1: 255, F2: 1}
                                                                    _ = fi56
                                                                    _ = v59
                                                                }
                                                                _ = v55
                                                            }
                                                        default:
                                                            v60 := GenS5{F0: (v3 & (^(v3))), F1: 0}
                                                            _ = v60
                                                        }
                                                    }
                                                    fmt.Println(v1, v1)
                                                    v61 := GenS6{F0: -9223372036854775808}
                                                    _ = v61
                                                }
                                                _ = len(v16)
                                                for fi62 := 0; fi62 < 11; fi62++ {
                                                    if len(v16) > 0 {
                                                        v16[0] = -9223372036854775808
                                                    }
                                                    _ = len(v16)
                                                    if false {
                                                        if (362 < 1) {
                                                            fmt.Println(v1, v1, v2)
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                fmt.Println(fi62)
                                                            default:
                                                                if (!(!true)) {
                                                                    var v63 uint = 65535
                                                                    v64 := GenS5{F0: 1, F1: 678}
                                                                    var v65 int32 = v3
                                                                    _ = v63
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            }
                                                            if true {
                                                                if false {
                                                                    v2 = v2
                                                                    v66 := GenS6{F0: -817}
                                                                    v2 = v2
                                                                    _ = v66
                                                                }
                                                                v16 = append(v16, (^(fi37)))
                                                                if (!false) {
                                                                    var v67 uint = 18446744073709551615
                                                                    _ = v67
                                                                }
                                                            }
                                                        } else {
                                                            if false {
                                                                if true {
                                                                    var v68 float64 = float64((((-2.25 - (3.14159 - 3.14159)) - 0.1) + float32(float64((1.5 * (0.0 * float32(float64(float32(float64(3.14159))))))))))
                                                                    _ = v68
                                                                } else {
                                                                    if true {
                                                                        if len(v16) >= 2 {
                                                                            v16 = v16[1:len(v16)]
                                                                        }
                                                                        switch 1 {
                                                                        case 2147483647:
                                                                            if len(v16) >= 2 {
                                                                                v16 = v16[1:len(v16)]
                                                                            }
                                                                        case 128:
                                                                            var v69 uint16 = 331
                                                                            _ = v69
                                                                        case 127:
                                                                            v70 := GenS7{F0: 0, F1: 1, F2: v3}
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                }
                                                                v1 += v1
                                                                v1 = v1
                                                            }
                                                            if len(v16) >= 2 {
                                                                v16 = v16[1:len(v16)]
                                                            }
                                                            v71 := GenS7{F0: 498, F1: v1, F2: (v3 + 1)}
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = fi62
                                                }
                                                _ = fi37
                                            }
                                        }
                                        _ = v36
                                    }
                                case "a":
                                    v16 = append(v16, 4294967295)
                                }
                            } else {
                                for fi72 := 0; fi72 < 1; fi72++ {
                                    if true {
                                        for fi73 := 0; fi73 < 7; fi73++ {
                                            v2 = v2
                                            v74 := GenS8{F0: 42}
                                            v1 += v1
                                            if ((1 == 255) || (v3 == (v3 ^ (-(v3))))) {
                                                var v75 int8 = -58
                                                var v76 int16 = 32767
                                                _ = len(v16)
                                                _ = v75
                                                _ = v76
                                            }
                                            v3 = -2147483648
                                            var v77 int = (fi73 ^ fi73)
                                            _ = fi73
                                            _ = v74
                                            _ = v77
                                        }
                                    }
                                    if (v1 != v1) {
                                        _ = len(v16)
                                        if false {
                                            v2 = ((v2 + (v2 + v2)) + v2)
                                        }
                                    } else {
                                        for fi78 := 0; fi78 < 14; fi78++ {
                                            v2 = "\n\t"
                                            if true {
                                                switch int(v3) {
                                                case 42:
                                                    var v79 int16 = int16(fi78)
                                                    v80 := GenS5{F0: v3, F1: 229}
                                                    v81 := GenS9{F0: v2, F1: 100}
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                case 32768:
                                                    v1 = v1
                                                default:
                                                    v2 = v2
                                                }
                                                fmt.Println(fi72, v1, fi72)
                                                v2 = ""
                                            } else {
                                                var v82 float32 = ((float32(1.5) * (float32(1.5) * 0.0)) + (1e10 + 0.0))
                                                _ = v82
                                            }
                                            v1 += 797
                                            _ = fi78
                                        }
                                    }
                                    if false {
                                        v83 := GenS10{F0: 199, F1: (^(v3))}
                                        fmt.Println(v2, fi72, v1)
                                        _ = v83
                                    } else {
                                        v84 := GenS11{F0: 429, F1: 26, F2: 131}
                                        if len(v16) > 0 {
                                            v16[0] = fi72
                                        }
                                        v85 := GenS12{F0: v3}
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = fi72
                                }
                                for fi86 := 0; fi86 < 1; fi86++ {
                                    v16 = append(v16, 32768)
                                    if ((false && (2147483648 >= 926)) && true) {
                                        v16 = append(v16, (-(fi86)))
                                        var v87 int8 = 59
                                        _ = v87
                                    } else {
                                        var v88 uint = 4294967295
                                        for fi89 := 0; fi89 < 6; fi89++ {
                                            v3 -= v3
                                            v1 = (v1 - (v1 * (v1 & v1)))
                                            if true {
                                                v90 := GenS7{F0: 262, F1: (v1 - v1), F2: -313}
                                                switch ((("foo bar" + ((v2 + (v2 + v2)) + (v2 + (v2 + (v2 + v2))))) + ("foo bar" + v2)) + ("" + ("" + "hello"))) {
                                                case "hello":
                                                    if ((!true) && ((v1 % (((^(v1))) | 1)) == (v1 - v1))) {
                                                        var v91 uint16 = 42
                                                        _ = v91
                                                    } else {
                                                        if false {
                                                            switch ("foo bar" + (((v2 + ((v2 + "") + (v2 + "hello"))) + (((("foo bar" + (v2 + v2)) + ("" + ("a" + (v2 + v2)))) + (v2 + "hello")) + ("hello" + v2))) + "hello")) {
                                                            case "a":
                                                                _ = len(v16)
                                                            case "hello":
                                                                if (true && true) {
                                                                    v2 = "foo bar"
                                                                    if ((false && false) || true) {
                                                                        var v92 float64 = -2.25
                                                                        var v93 uint = 1
                                                                        v1 = (v1 - (v1 & 0))
                                                                        fmt.Println(v93)
                                                                        var v94 uint = v93
                                                                        _ = v92
                                                                        _ = v93
                                                                        _ = v94
                                                                    } else {
                                                                        var v95 uint32 = uint32(v1)
                                                                        fmt.Println(v1, v95)
                                                                        _ = v95
                                                                    }
                                                                }
                                                            }
                                                            var v96 int16 = 42
                                                            v2 = v2
                                                            var v97 bool = false
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        if len(v16) > 0 {
                                                            v16[0] = (^(fi89))
                                                        }
                                                    }
                                                default:
                                                    var v98 int32 = v3
                                                    _ = v98
                                                }
                                                v88 -= 1
                                                _ = v90
                                            }
                                            _ = fi89
                                        }
                                        _ = v88
                                    }
                                    for fi99 := 0; fi99 < 9; fi99++ {
                                        var v100 uint64 = 425
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        fmt.Println(v100, v100)
                                        v1 = 1
                                        _ = fi99
                                        _ = v100
                                    }
                                    _ = fi86
                                }
                                _ = len(v16)
                            }
                        case 127:
                            v16 = append(v16, int(v1))
                        }
                    case 4294967295:
                        var v101 uint16 = 1
                        _ = v101
                    }
                    fmt.Println(v2, v3)
                    v16 = append(v16, -385)
                }
            }
            var v102 bool = true
            _ = v102
        }
    } else {
        var v103 int16 = 32767
        v104 := GenS13{F0: 9223372036854775807, F1: v1, F2: 0}
        for fi105 := 0; fi105 < 19; fi105++ {
            v2 = (v2 + (v2 + v2))
            if (42 > 129) {
                var v106 int64 = 1
                if true {
                    if true {
                        v3 = v3
                        var v107 int8 = -1
                        _ = v107
                    }
                }
                _ = v106
            }
            _ = fi105
        }
        _ = v103
        _ = v104
    }
    fmt.Println(v3, v2)
    switch -265 {
    case -1:
        fmt.Println(v3, v2)
    case -2147483648:
        if (127 <= 255) {
            var v108 int64 = -1
            fmt.Println(v1, v3)
            _ = v108
        }
    }
    if false {
        v109 := GenS4{F0: v2, F1: (v3 + v3)}
        v110 := GenS14{F0: 42, F1: int16(v3), F2: 235}
        _ = v109
        _ = v110
    }
    if (218 > 0) {
        v111 := GenS13{F0: -574, F1: (v1 / ((42) | 1)), F2: 65535}
        _ = v111
    } else {
        v1 -= (v1 ^ (v1 | v1))
    }
    for fi112 := 0; fi112 < 14; fi112++ {
        v16 = append(v16, 128)
        v113 := GenS2{F0: 2, F1: true}
        v114 := GenS6{F0: genFn0(v1, 46, v1)}
        _ = fi112
        _ = v113
        _ = v114
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v115 uint16 = 2
    v115 = v115
    fmt.Println(v3, v115, v3)
    v3 = (v3 ^ v3)
    for fi116 := 0; fi116 < 12; fi116++ {
        if true {
            fmt.Println(v1, fi116, v2)
        }
        for fi117 := 0; fi117 < 13; fi117++ {
            _ = len(v16)
            v118 := GenS15{F0: 1, F1: 2147483647}
            _ = fi117
            _ = v118
        }
        for fi119 := 0; fi119 < 17; fi119++ {
            v120 := GenS16{F0: true, F1: 1000}
            v16 = append(v16, (fi116 / (((fi116 * 4294967295)) | 1)))
            v115 += v115
            _ = fi119
            _ = v120
        }
        _ = fi116
    }
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v115
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
