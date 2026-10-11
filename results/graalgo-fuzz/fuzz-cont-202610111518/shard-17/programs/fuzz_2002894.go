package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint8; F1 string; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint8; F2 int64 }
type GenS7 struct { F0 int }

func main() {
    var v1 int32 = -1
    var v2 uint = 2
    v3 := []int{}
    var v4 int16 = -464
    if true {
        fmt.Println(v1, v4)
        v5 := GenS1{F0: ((-128 <= 127) && true), F1: 303, F2: -9223372036854775808}
        _ = v5
    }
    var v6 uint8 = 1
    v3 = append(v3, -366)
    var v7 int64 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 4294967295 {
    case 32767:
        v7 -= v7
    case 127:
        if (65535 >= 1) {
            if ((v7 + (^(v7))) != (v7 % (((v7 / ((-9223372036854775808) | 1))) | 1))) {
                switch "a" {
                case "\n\t":
                    var v8 bool = true
                    _ = v8
                case "":
                    if (2 >= 0) {
                        var v9 uint8 = 62
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v9
                    }
                case "foo bar":
                    for fi10 := 0; fi10 < 6; fi10++ {
                        if ((true && (((((("\n\t" + ("foo bar" + "")) + (((("a" + "\n\t") + "") + "hello") + ("a" + "foo bar"))) + "hello") + "a") == "a") || false)) && true) {
                            if (1 < 2) {
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(v2)
                                case "a":
                                    if len(v3) > 0 {
                                        v3[0] = (fi10 / ((-984) | 1))
                                    }
                                case "\n\t":
                                    var v11 int32 = (v1 * 1)
                                    _ = v11
                                default:
                                    v12 := GenS1{F0: (fi10 != 0), F1: 18446744073709551615, F2: (-(v7))}
                                    _ = v12
                                }
                            } else {
                                for fi13 := 0; fi13 < 12; fi13++ {
                                    if (0 < 180) {
                                        if (0.0 > float64(0.0)) {
                                            if (v4 < -518) {
                                                v14 := GenS2{F0: 2}
                                                var v15 int32 = v1
                                                _ = v14
                                                _ = v15
                                            }
                                            v16 := GenS2{F0: 2}
                                            var v17 int8 = 1
                                            _ = v16
                                            _ = v17
                                        }
                                        if (1 == 0) {
                                            v3 = append(v3, 9223372036854775807)
                                            var v18 uint16 = 747
                                            _ = v18
                                        } else {
                                            var v19 uint8 = v6
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v19)
                                            case "foo bar":
                                                v20 := GenS2{F0: 911}
                                                _ = v20
                                            }
                                            _ = v19
                                        }
                                        fmt.Println(v4)
                                        var v21 float64 = ((1e10 * -2.25) / float64(-2.25))
                                        switch "" {
                                        case "a":
                                            v3 = append(v3, 32768)
                                        case "hello":
                                            var v22 uint = (v2 + v2)
                                            _ = v22
                                        case "foo bar":
                                            _ = len(v3)
                                        }
                                        v4 += int16(fi10)
                                        var v23 uint64 = 1
                                        if true {
                                            if false {
                                                var v24 int16 = (v4 | (^(v4)))
                                                v25 := GenS2{F0: 42}
                                                _ = v24
                                                _ = v25
                                            }
                                        } else {
                                            v26 := GenS2{F0: 993}
                                            var v27 int = (-(fi13))
                                            _ = v26
                                            _ = v27
                                        }
                                        if true {
                                            if true {
                                                if (((v23 | (v23 | (v23 & (v23 - (v23 << ((v2) & 63)))))) > 180) && ((fi13 % ((4294967295) | 1)) == (fi13 + (fi13 >> ((20) & 63))))) {
                                                    var v28 string = ""
                                                    v1 = v1
                                                    _ = v28
                                                } else {
                                                    v1 -= 1
                                                    fmt.Println(v1, v1, fi10)
                                                    v1 = (v1 * v1)
                                                    switch ("\n\t" + (("" + "a") + "foo bar")) {
                                                    case "":
                                                        fmt.Println(v23)
                                                    case "foo bar":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "\n\t":
                                                        var v29 string = ("hello" + ((("foo bar" + "a") + "\n\t") + "\n\t"))
                                                        _ = v29
                                                    default:
                                                        fmt.Println(v4)
                                                    }
                                                }
                                                v21 = 3.14159
                                                var v30 uint = (^(v2))
                                                var v31 uint8 = (v6 << ((v30) & 63))
                                                var v32 int = (fi13 & fi10)
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                            var v33 int64 = v7
                                            _ = v33
                                        } else {
                                            v34 := GenS1{F0: true, F1: (^(v23)), F2: (^(v7))}
                                            fmt.Println(fi10, fi13, v1)
                                            var v35 uint8 = v6
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 uint32 = 1
                                        fmt.Println(v2, v21)
                                        _ = v21
                                        _ = v23
                                        _ = v36
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v37 float64 = 0.1
                                        fmt.Println(v4)
                                        _ = v37
                                    }
                                    var v38 bool = (127 < -103)
                                    v2 = v2
                                    _ = fi13
                                    _ = v38
                                }
                            }
                        } else {
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = fi10
                                }
                            }
                        }
                        _ = fi10
                    }
                default:
                    switch "a" {
                    case "foo bar":
                        v3 = append(v3, int(v1))
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 32768
                        }
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        fmt.Println(v4, v7, v4)
                    }
                }
                fmt.Println(v7, v4, v1)
            } else {
                if false {
                    for fi39 := 0; fi39 < 17; fi39++ {
                        var v40 uint8 = v6
                        if true {
                            fmt.Println(v6, v7)
                        }
                        _ = fi39
                        _ = v40
                    }
                    var v41 float32 = -2.25
                    _ = v41
                } else {
                    if (4294967295 == 18446744073709551615) {
                        v2 += (^(v2))
                        for fi42 := 0; fi42 < 13; fi42++ {
                            var v43 float64 = float64(1.5)
                            if true {
                                v43 += float64((1e10 / (1.5 / 1e10)))
                                v44 := GenS3{F0: 935, F1: v6}
                                _ = v44
                            }
                            _ = fi42
                            _ = v43
                        }
                        v4 = v4
                    } else {
                        switch -19 {
                        case -9223372036854775808:
                            v45 := GenS1{F0: (31 > 32767), F1: 0, F2: -1}
                            _ = v45
                        case 42:
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                        default:
                            for fi46 := 0; fi46 < 14; fi46++ {
                                for fi47 := 0; fi47 < 4; fi47++ {
                                    v7 = -1
                                    var v48 uint8 = 249
                                    _ = fi47
                                    _ = v48
                                }
                                v49 := GenS1{F0: true, F1: 12, F2: (^(v7))}
                                _ = fi46
                                _ = v49
                            }
                        }
                        for fi50 := 0; fi50 < 2; fi50++ {
                            v1 = v1
                            fmt.Println(fi50, v4, v7)
                            _ = fi50
                        }
                    }
                    switch 128 {
                    case -2147483648:
                        switch "" {
                        case "foo bar":
                            v4 += -764
                        case "":
                            var v51 int = 127
                            _ = v51
                        }
                    case 127:
                        var v52 float32 = 0.1
                        _ = v52
                    default:
                        if true {
                            if ((v4 >> ((2) & 63)) < (v4 - (v4 ^ v4))) {
                                var v53 float32 = float32((float64(1.5) / 0.1))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v54 int64 = v7
                                    _ = v54
                                case "a":
                                    var v55 uint16 = 42
                                    _ = v55
                                case "\n\t":
                                    var v56 int32 = -1
                                    _ = v56
                                }
                                _ = v53
                            } else {
                                if false {
                                    switch 128 {
                                    case 32767:
                                        for fi57 := 0; fi57 < 7; fi57++ {
                                            for fi58 := 0; fi58 < 4; fi58++ {
                                                v6 -= 2
                                                _ = fi58
                                            }
                                            _ = fi57
                                        }
                                    case -1:
                                        for fi59 := 0; fi59 < 20; fi59++ {
                                            v1 += (v1 << ((709) & 63))
                                            var v60 string = ("hello" + "")
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi59
                                            _ = v60
                                        }
                                    default:
                                        v61 := GenS3{F0: 255, F1: (v6 & (v6 >> ((3) & 63)))}
                                        _ = v61
                                    }
                                }
                            }
                            switch 42 {
                            case 530:
                                fmt.Println(v7, v6)
                            case 1:
                                v1 += 1
                            default:
                                v6 = (v6 | 138)
                            }
                            if (!((v2 * v2) != v2)) {
                                var v62 int = 0
                                v63 := GenS1{F0: (float32(0.0) < (3.14159 / float32(3.14159))), F1: 18446744073709551615, F2: v7}
                                for fi64 := 0; fi64 < 7; fi64++ {
                                    for fi65 := 0; fi65 < 6; fi65++ {
                                        v4 += v4
                                        v3 = append(v3, v62)
                                        switch 32768 {
                                        case 4294967295:
                                            v1 = v1
                                        case 713:
                                            v1 += (v1 << (((v2 ^ v2)) & 63))
                                            if len(v3) > 0 {
                                                v3[0] = (v62 | fi65)
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case 22:
                                            if ((^(v1)) == (-(v1))) {
                                                v3 = append(v3, (v62 * (v62 & fi64)))
                                                if ((!false) || ((((v1 + (v1 + v1)) == (-(v1))) && (true || ((42 >= int8(v2)) && ((-9223372036854775808 != (fi65 ^ v62)) || (100 <= 1))))) && ((0.1 * (3.14159 + (1.5 * float32(3.14159)))) == 1e10))) {
                                                    if len(v3) > 0 {
                                                        v3[0] = fi65
                                                    }
                                                    if false {
                                                        switch 285 {
                                                        case -651:
                                                            if (false || true) {
                                                                var v66 uint16 = 618
                                                                var v67 uint = (v2 + v2)
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                v6 += (v6 >> ((v2) & 63))
                                                                var v68 int64 = (v7 ^ v7)
                                                                v68 = v68
                                                                v4 -= v4
                                                                _ = v68
                                                            }
                                                        default:
                                                            var v69 uint32 = 1
                                                            _ = v69
                                                        }
                                                        switch 304 {
                                                        case 42:
                                                            v7 -= int64(v1)
                                                        }
                                                    }
                                                    v7 = (v7 >> ((v2) & 63))
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = fi65
                                                    }
                                                }
                                                var v70 int8 = -26
                                                var v71 int8 = (v70 ^ (v70 / ((v70) | 1)))
                                                v70 += (v70 + (v71 % ((v71) | 1)))
                                                _ = v70
                                                _ = v71
                                            } else {
                                                var v72 uint8 = 128
                                                switch fi64 {
                                                case 32768:
                                                    var v73 float32 = (((1.5 / 3.14159) * (1.5 - (-2.25 / 1.5))) * (float32(0.1) + float32((((1e10 - float64(((-2.25 / -2.25) + 0.1))) / 1e10) + (-2.25 * 0.0)))))
                                                    v74 := GenS4{F0: v72, F1: ("" + ("foo bar" + ("" + ""))), F2: 11}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if true {
                                                        var v75 uint16 = 2
                                                        _ = v75
                                                    } else {
                                                        v4 += v4
                                                        switch (fi64 << (((^(v2))) & 63)) {
                                                        case 32767:
                                                            _ = len(v3)
                                                        }
                                                        v4 = 42
                                                    }
                                                    var v76 bool = false
                                                    if (v76 || v76) {
                                                        v77 := GenS1{F0: ((v4 | v4) >= -359), F1: 607, F2: -877}
                                                        switch (v62 ^ -233) {
                                                        case -498:
                                                            var v78 int64 = -396
                                                            var v79 uint = (v2 << ((v2) & 63))
                                                            _ = v78
                                                            _ = v79
                                                        case -9223372036854775808:
                                                            _ = len(v3)
                                                        case 1:
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                if false {
                                                                    var v80 float64 = (1.5 / float64((float32((1e10 / float64(3.14159))) / v73)))
                                                                    if ((v4 >> ((625) & 63)) <= v4) {
                                                                        v2 += (v2 - 748)
                                                                        var v81 float32 = (1e10 / float32(1.5))
                                                                        var v82 uint8 = (v6 % (((v72 % ((v6) | 1))) | 1))
                                                                        v83 := GenS1{F0: (v80 >= v80), F1: 18446744073709551615, F2: (v7 ^ (v7 ^ (v7 >> (((v2 << ((65535) & 63))) & 63))))}
                                                                        if true {
                                                                            v73 = ((((v73 - 1.5) * (1e10 - v73)) / (((0.1 / (v73 - (1e10 / v81))) - float32(-2.25)) + v73)) / (v73 + float32(v80)))
                                                                        } else {
                                                                            v6 += 161
                                                                            switch -2147483648 {
                                                                            case 508:
                                                                                var v84 uint64 = 2
                                                                                _ = v84
                                                                            case -23:
                                                                                v85 := GenS5{F0: 32767}
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                        _ = v81
                                                                        _ = v82
                                                                        _ = v83
                                                                    } else {
                                                                        if (v76 || false) {
                                                                            var v86 float64 = (((((v80 + (v80 + ((v80 / v80) + v80))) / -2.25) / float64(float32(v80))) * v80) / float64(float32((0.0 / (v80 / ((float64(v73) - float64(float32((v80 + ((v80 - v80) - v80))))) / (((v80 - (v80 * float64(v73))) - (v80 * 0.1)) / v80)))))))
                                                                            v76 = false
                                                                            _ = v86
                                                                        }
                                                                        v6 = v72
                                                                    }
                                                                    _ = v80
                                                                }
                                                            case "\n\t":
                                                                v3 = append(v3, fi65)
                                                            case "hello":
                                                                var v87 uint = v2
                                                                _ = v87
                                                            }
                                                            fmt.Println(fi64, v6)
                                                        }
                                                        _ = v77
                                                    }
                                                    _ = v73
                                                    _ = v74
                                                    _ = v76
                                                case 1:
                                                    var v88 uint16 = 65535
                                                    _ = v88
                                                case 0:
                                                    v1 -= int32(v6)
                                                default:
                                                    var v89 float64 = 0.1
                                                    _ = v89
                                                }
                                                _ = v72
                                            }
                                        }
                                        _ = fi65
                                    }
                                    v7 = (v7 << ((v2) & 63))
                                    var v90 uint8 = 42
                                    _ = fi64
                                    _ = v90
                                }
                                _ = v62
                                _ = v63
                            }
                        } else {
                            v91 := GenS6{F0: false, F1: v6, F2: (v7 - 9223372036854775807)}
                            var v92 uint64 = 4294967295
                            _ = v91
                            _ = v92
                        }
                    }
                    var v93 uint8 = v6
                    _ = v93
                }
                var v94 uint32 = 4294967295
                switch "\n\t" {
                case "hello":
                    v2 = 2
                }
                _ = v94
            }
            v95 := GenS2{F0: 65535}
            var v96 int16 = -730
            _ = v95
            _ = v96
        }
    case 42:
        _ = len(v3)
    default:
        v7 += v7
    }
    for fi97 := 0; fi97 < 8; fi97++ {
        switch (fi97 & 32768) {
        case 1:
            for fi98 := 0; fi98 < 8; fi98++ {
                var v99 int16 = v4
                v7 = -1
                _ = fi98
                _ = v99
            }
        case 9223372036854775807:
            var v100 int32 = -761
            _ = v100
        case 4294967295:
            fmt.Println(v7)
        default:
            for fi101 := 0; fi101 < 18; fi101++ {
                v1 = -2147483648
                v6 -= (^(v6))
                switch "a" {
                case "foo bar":
                    if (!(42 > (^(v6)))) {
                        fmt.Println(v6)
                        v3 = append(v3, (fi101 << (((v2 % (((^(v2))) | 1))) & 63)))
                        var v102 bool = true
                        _ = v102
                    } else {
                        var v103 float32 = (-2.25 * 0.1)
                        _ = v103
                    }
                case "a":
                    v2 = v2
                default:
                    v2 += (v2 | (v2 ^ (v2 % ((65535) | 1))))
                }
                _ = fi101
            }
        }
        if (647 <= int32(v2)) {
            var v104 float32 = -2.25
            var v105 uint = (v2 & v2)
            _ = v104
            _ = v105
        }
        _ = fi97
    }
    for fi106 := 0; fi106 < 12; fi106++ {
        for fi107 := 0; fi107 < 9; fi107++ {
            var v108 int16 = v4
            v4 -= (^(v4))
            _ = fi107
            _ = v108
        }
        for fi109 := 0; fi109 < 13; fi109++ {
            v1 = 609
            _ = fi109
        }
        if (((v1 & 0) <= v1) && false) {
            if (v4 < v4) {
                v4 = v4
            }
            v110 := GenS7{F0: 32767}
            _ = v110
        } else {
            v111 := GenS6{F0: (0.1 < float64((0.0 + 3.14159))), F1: (v6 ^ v6), F2: v7}
            var v112 uint = (v2 ^ (v2 + v2))
            _ = v111
            _ = v112
        }
        _ = fi106
    }
    v113 := GenS2{F0: 598}
    _ = len(v3)
    _ = len(v3)
    v114 := GenS6{F0: (2147483647 >= 42), F1: (v6 / ((128) | 1)), F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
