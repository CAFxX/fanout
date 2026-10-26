package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int64; F1 int16; F2 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8; F1 int64; F2 string }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint8; F1 uint }
type GenS13 struct { F0 bool; F1 uint }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int8; F1 uint16 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int32; F1 bool; F2 string }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 string; F1 int64 }
type GenS20 struct { F0 uint32; F1 int16 }
type GenS21 struct { F0 int32; F1 int16; F2 int8 }
type GenS22 struct { F0 bool; F1 int16 }
func genFn0(p0 uint32, p1 int) int64 {
    v6 := []int{}
    switch "hello" {
    case "\n\t":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "foo bar":
        fmt.Println(p1)
    case "a":
        if (42 <= 275) {
            if false {
                p0 = 2147483648
                fmt.Println(p1, p1)
                var v7 int8 = -80
                _ = v7
            } else {
                var v8 uint8 = 128
                v9 := GenS1{F0: 65535, F1: -32768, F2: 30}
                _ = v8
                _ = v9
            }
            var v10 uint64 = 0
            if (p1 != p1) {
                p0 = (p0 ^ (p0 << ((214) & 63)))
            } else {
                switch "\n\t" {
                case "foo bar":
                    var v11 uint16 = 0
                    _ = v11
                case "a":
                    if len(v6) > 0 {
                        v6[0] = (p1 >> ((276) & 63))
                    }
                case "\n\t":
                    if (false || ((!(128 == (p1 >> ((18446744073709551615) & 63)))) && (false && false))) {
                        _ = len(v6)
                        switch p1 {
                        case -2147483648:
                            if (18446744073709551615 == (v10 >> ((4294967295) & 63))) {
                                var v12 string = "hello"
                                _ = v12
                            } else {
                                switch p1 {
                                case 32768:
                                    var v13 uint16 = 2
                                    _ = v13
                                case 0:
                                    for fi14 := 0; fi14 < 1; fi14++ {
                                        switch ("hello" + ("foo bar" + "a")) {
                                        case "":
                                            for fi15 := 0; fi15 < 9; fi15++ {
                                                fmt.Println(v10)
                                                if (!(p0 != p0)) {
                                                    var v16 uint = 42
                                                    var v17 uint16 = 568
                                                    v6 = append(v6, 127)
                                                    _ = v16
                                                    _ = v17
                                                } else {
                                                    p0 += (p0 & (^(p0)))
                                                    var v18 int8 = 1
                                                    _ = v18
                                                }
                                                v6 = append(v6, 4294967295)
                                                if false {
                                                    var v19 uint64 = 4294967295
                                                    var v20 int8 = 42
                                                    p1 -= p1
                                                    _ = v19
                                                    _ = v20
                                                } else {
                                                    switch fi15 {
                                                    case 42:
                                                        if (9223372036854775807 > 1) {
                                                            if len(v6) > 0 {
                                                                v6[0] = (p1 & fi15)
                                                            }
                                                            v10 = v10
                                                        } else {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if (1 != 71) {
                                                                    fmt.Println(fi15)
                                                                    if (uint8(p0) < 0) {
                                                                        v21 := GenS1{F0: 255, F1: 1, F2: 0}
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi15
                                                                        }
                                                                        p1 = int(p0)
                                                                        if len(v6) > 0 {
                                                                            v6[0] = (fi14 & (fi14 & 0))
                                                                        }
                                                                        _ = v21
                                                                    }
                                                                    if ((true && ((2147483647 >= 406) && ((!(!(!(true && (!(1.5 <= (3.14159 / float32((3.14159 - (1.5 / ((float64(0.0) * float64(1.5)) + 3.14159))))))))))) || (0 <= 42)))) && true) {
                                                                        var v22 int64 = 0
                                                                        v10 += v10
                                                                        _ = v22
                                                                    }
                                                                    if false {
                                                                        v23 := GenS2{F0: v10}
                                                                        var v24 int8 = -128
                                                                        if true {
                                                                            p0 -= 2147483647
                                                                        } else {
                                                                            var v25 uint32 = p0
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                var v26 int32 = 1
                                                                                _ = v26
                                                                            default:
                                                                                p0 = 370
                                                                            }
                                                                            _ = v25
                                                                        }
                                                                        v24 = (v24 - 0)
                                                                        _ = v23
                                                                        _ = v24
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            fmt.Println(v10)
                                                                            var v27 uint = 42
                                                                            _ = v27
                                                                        }
                                                                        p1 = fi15
                                                                    }
                                                                    var v28 int8 = 0
                                                                    _ = v28
                                                                } else {
                                                                    v29 := GenS1{F0: 42, F1: 1000, F2: 0}
                                                                    _ = v29
                                                                }
                                                            case "foo bar":
                                                                p0 -= (p0 - uint32(v10))
                                                            default:
                                                                v10 = 34
                                                            }
                                                            p0 -= 2147483647
                                                            fmt.Println(p0, fi15, fi15)
                                                        }
                                                    case 128:
                                                        switch -2147483648 {
                                                        case -1:
                                                            var v30 int = fi14
                                                            _ = v30
                                                        default:
                                                            if len(v6) > 0 {
                                                                v6[0] = (-(fi14))
                                                            }
                                                        }
                                                    case -9223372036854775808:
                                                        v31 := GenS1{F0: 255, F1: -32768, F2: -1}
                                                        _ = v31
                                                    }
                                                    if (9223372036854775807 > 0) {
                                                        if (("\n\t" + ("\n\t" + "")) <= "a") {
                                                            p1 += 581
                                                        } else {
                                                            switch (-(fi15)) {
                                                            case 2147483647:
                                                                var v32 uint64 = (v10 | 2)
                                                                _ = v32
                                                            case 127:
                                                                if (!true) {
                                                                    var v33 bool = false
                                                                    _ = v33
                                                                } else {
                                                                    var v34 float32 = 0.1
                                                                    v35 := GenS2{F0: 302}
                                                                    v36 := GenS3{F0: 255, F1: false}
                                                                    _ = v34
                                                                    _ = v35
                                                                    _ = v36
                                                                }
                                                                var v37 uint = 718
                                                                _ = v37
                                                            default:
                                                                v10 = (v10 ^ v10)
                                                            }
                                                            fmt.Println(fi14, fi15)
                                                            if (-2.25 >= (float64(1.5) / (1e10 * 0.1))) {
                                                                fmt.Println(fi14, fi14, fi14)
                                                                p0 -= p0
                                                                v38 := GenS4{F0: 0, F1: (false || false), F2: (p0 & p0)}
                                                                _ = v38
                                                            } else {
                                                                p0 -= (p0 % ((2147483647) | 1))
                                                            }
                                                            p0 -= (p0 >> ((65535) & 63))
                                                        }
                                                        v39 := GenS5{F0: (^(v10)), F1: (883 <= (v10 / ((v10) | 1)))}
                                                        _ = v39
                                                    }
                                                }
                                                _ = fi15
                                            }
                                        default:
                                            p0 += 65535
                                        }
                                        p1 = (fi14 + fi14)
                                        fmt.Println(fi14)
                                        _ = fi14
                                    }
                                case 9223372036854775807:
                                    var v40 uint32 = 2
                                    _ = v40
                                default:
                                    p1 += (p1 + (p1 & (p1 / ((p1) | 1))))
                                }
                                var v41 uint16 = 32767
                                var v42 uint8 = 4
                                _ = v41
                                _ = v42
                            }
                        case -637:
                            var v43 uint16 = 918
                            _ = v43
                        case -331:
                            for fi44 := 0; fi44 < 3; fi44++ {
                                switch p1 {
                                case 127:
                                    switch ("" + "") {
                                    case "hello":
                                        if len(v6) > 0 {
                                            v6[0] = -2147483648
                                        }
                                    case "a":
                                        var v45 int64 = 702
                                        _ = v45
                                    case "foo bar":
                                        v6 = append(v6, (fi44 + fi44))
                                    }
                                case 0:
                                    fmt.Println(p1)
                                case -785:
                                    for fi46 := 0; fi46 < 4; fi46++ {
                                        v47 := GenS1{F0: 4294967295, F1: -32768, F2: -99}
                                        if len(v6) > 0 {
                                            v6[0] = 1
                                        }
                                        var v48 int8 = 0
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                _ = fi44
                            }
                        }
                        switch (^(p1)) {
                        case 42:
                            for fi49 := 0; fi49 < 20; fi49++ {
                                fmt.Println(v10, p1, p0)
                                if true {
                                    p1 -= (p1 % (((fi49 & 9223372036854775807)) | 1))
                                } else {
                                    fmt.Println(p1)
                                    if (fi49 <= fi49) {
                                        v50 := GenS5{F0: (v10 | 382), F1: (((false && true) && false) && false)}
                                        if (!true) {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = fi49
                                        }
                                        _ = v50
                                    }
                                }
                                var v51 uint8 = 70
                                _ = fi49
                                _ = v51
                            }
                        case -624:
                            v52 := GenS6{F0: -38, F1: 2}
                            _ = v52
                        case -9223372036854775808:
                            switch "a" {
                            case "hello":
                                if true {
                                    var v53 string = "hello"
                                    _ = v53
                                }
                            case "\n\t":
                                var v54 int8 = 43
                                _ = v54
                            case "a":
                                p0 += (p0 - p0)
                            default:
                                for fi55 := 0; fi55 < 1; fi55++ {
                                    var v56 int = (fi55 - fi55)
                                    if (1.5 != 0.1) {
                                        if false {
                                            p1 -= fi55
                                        } else {
                                            if false {
                                                var v57 float64 = 3.14159
                                                _ = v57
                                            } else {
                                                fmt.Println(v10)
                                                var v58 uint8 = 42
                                                if true {
                                                    var v59 bool = false
                                                    _ = v59
                                                }
                                                _ = v58
                                            }
                                        }
                                        v10 += v10
                                        switch v56 {
                                        case -2147483648:
                                            var v60 int64 = -634
                                            _ = v60
                                        case 512:
                                            var v61 uint = 18446744073709551615
                                            _ = v61
                                        case 782:
                                            if (1.5 != -2.25) {
                                                var v62 int32 = 2147483647
                                                switch (p1 + (p1 + fi55)) {
                                                case 2147483647:
                                                    var v63 uint8 = 57
                                                    _ = v63
                                                case -322:
                                                    if false {
                                                        fmt.Println(p0, p1)
                                                        var v64 int32 = v62
                                                        var v65 bool = (211 > 210)
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        v10 += 1
                                                        p1 = (fi55 << ((255) & 63))
                                                        for fi66 := 0; fi66 < 18; fi66++ {
                                                            v67 := GenS7{F0: 0, F1: 691, F2: 2}
                                                            if true {
                                                                var v68 uint8 = 113
                                                                v56 -= -528
                                                                var v69 uint8 = (v68 % (((v68 << ((18446744073709551615) & 63))) | 1))
                                                                _ = v68
                                                                _ = v69
                                                            } else {
                                                                var v70 int = int(v62)
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v71 uint8 = 0
                                                                    _ = v71
                                                                case "hello":
                                                                    var v72 int = (fi55 % ((v56) | 1))
                                                                    _ = v72
                                                                default:
                                                                    p0 += p0
                                                                }
                                                                _ = v70
                                                            }
                                                            if (!false) {
                                                                switch ("foo bar" + "\n\t") {
                                                                case "\n\t":
                                                                    switch ("\n\t" + "\n\t") {
                                                                    case "foo bar":
                                                                        if (p0 == (p0 & p0)) {
                                                                            fmt.Println(v56)
                                                                            v62 = (v62 | v62)
                                                                            if false {
                                                                                p1 = fi66
                                                                            } else {
                                                                                if false {
                                                                                    var v73 uint64 = v10
                                                                                    var v74 int32 = (v62 ^ (v62 ^ (^(v62))))
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                } else {
                                                                                    v10 -= v10
                                                                                }
                                                                                p0 -= p0
                                                                            }
                                                                        } else {
                                                                            var v75 uint16 = 32767
                                                                            _ = v75
                                                                        }
                                                                    case "hello":
                                                                        if true {
                                                                            v6 = append(v6, p1)
                                                                        }
                                                                    case "a":
                                                                        v76 := GenS6{F0: -1, F1: 128}
                                                                        _ = v76
                                                                    }
                                                                }
                                                            }
                                                            var v77 uint64 = (v10 & (^(v10)))
                                                            v10 -= (v77 * (v10 & (v10 + v10)))
                                                            _ = fi66
                                                            _ = v67
                                                            _ = v77
                                                        }
                                                    }
                                                case 128:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                default:
                                                    switch (p1 / ((fi55) | 1)) {
                                                    case 32767:
                                                        fmt.Println(v62)
                                                    case 9223372036854775807:
                                                        fmt.Println(v62)
                                                    case -981:
                                                        fmt.Println(fi55)
                                                    default:
                                                        v56 += 32768
                                                    }
                                                }
                                                _ = v62
                                            }
                                        }
                                    } else {
                                        p0 -= 2
                                    }
                                    p1 = 128
                                    _ = fi55
                                    _ = v56
                                }
                            }
                        default:
                            if (2 == (p0 | (^(p0)))) {
                                var v78 int16 = -490
                                _ = v78
                            } else {
                                for fi79 := 0; fi79 < 15; fi79++ {
                                    if false {
                                        var v80 uint64 = (v10 >> ((2) & 63))
                                        v81 := GenS4{F0: (v80 >> ((0) & 63)), F1: true, F2: (p0 / ((p0) | 1))}
                                        var v82 float64 = 0.0
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    } else {
                                        v83 := GenS8{F0: 970}
                                        _ = v83
                                    }
                                    for fi84 := 0; fi84 < 13; fi84++ {
                                        if false {
                                            switch fi84 {
                                            case 42:
                                                if true {
                                                    var v85 int64 = 1
                                                    _ = v85
                                                } else {
                                                    var v86 int8 = 11
                                                    _ = len(v6)
                                                    _ = v86
                                                }
                                            }
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = -965
                                            }
                                            if true {
                                                var v87 int8 = 100
                                                v10 = v10
                                                var v88 uint8 = 128
                                                _ = v87
                                                _ = v88
                                            } else {
                                                var v89 bool = (-9223372036854775808 < -9223372036854775808)
                                                fmt.Println(fi84, v89, p1)
                                                v90 := GenS2{F0: (v10 >> ((1) & 63))}
                                                _ = v89
                                                _ = v90
                                            }
                                        }
                                        v91 := GenS9{F0: -624}
                                        _ = fi84
                                        _ = v91
                                    }
                                    fmt.Println(fi79, v10, p1)
                                    _ = fi79
                                }
                            }
                        }
                    } else {
                        switch -246 {
                        case 32767:
                            if len(v6) > 0 {
                                v6[0] = p1
                            }
                        default:
                            p1 += (p1 + (p1 >> ((18446744073709551615) & 63)))
                        }
                        var v92 float32 = float32(1.5)
                        _ = v92
                    }
                default:
                    v6 = append(v6, p1)
                }
            }
            _ = v10
        }
    default:
        p1 = p1
    }
    _ = len(v6)
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 uint32, p1 int16, p2 uint32) string {
    v93 := GenS1{F0: 18446744073709551615, F1: (p1 + p1), F2: 1}
    v94 := GenS10{F0: 54, F1: 0, F2: "foo bar"}
    v95 := GenS11{F0: -27}
    for fi96 := 0; fi96 < 17; fi96++ {
        fmt.Println(p2)
        _ = fi96
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v93
    _ = v94
    _ = v95
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 bool = (!false)
    var v2 float64 = float64((0.0 * (3.14159 * (0.0 * (float32(-2.25) * 1e10)))))
    var v3 uint8 = 42
    v4 := []string{"foo bar"}
    var v5 uint64 = 772
    var v97 int32 = int32(v5)
    var v98 float64 = (v2 / ((v2 * 0.0) / (((((v2 * float64(0.1)) * (float64(0.0) + 3.14159)) / (v2 / 3.14159)) / (1e10 * (0.1 - v2))) + v2)))
    v97 -= (v97 - (v97 / (((-(v97))) | 1)))
    if ((42 != 2) && ("foo bar" != (("" + "foo bar") + "a"))) {
        switch "hello" {
        case "a":
            v98 += (1e10 + (v98 / (v98 / (((float64(3.14159) + float64(0.0)) * v98) / (0.0 * ((((v98 - 1.5) * 0.1) - ((v98 / 1.0) / 1.5)) + (float64(-2.25) / -2.25)))))))
        case "\n\t":
            v5 = (v5 | v5)
        }
    }
    for fi99 := 0; fi99 < 15; fi99++ {
        v100 := GenS12{F0: (^(v3)), F1: 42}
        v101 := GenS11{F0: 0}
        if (v1 || (!v1)) {
            fmt.Println(v3)
            for fi102 := 0; fi102 < 6; fi102++ {
                fmt.Println(v2)
                _ = fi102
            }
        } else {
            if v1 {
                if (1 < 801) {
                    fmt.Println(v1, v97)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v103 int16 = 550
                    _ = v103
                }
                for fi104 := 0; fi104 < 18; fi104++ {
                    switch ((("hello" + "\n\t") + ("foo bar" + "a")) + "foo bar") {
                    case "":
                        v105 := GenS11{F0: 42}
                        _ = v105
                    case "foo bar":
                        v106 := GenS13{F0: (-200 > v97), F1: 4294967295}
                        _ = v106
                    default:
                        fmt.Println(v1, fi104, v1)
                        if (!(!(v1 && (!false)))) {
                            v4 = append(v4, "")
                            var v107 uint16 = 255
                            var v108 float32 = 3.14159
                            _ = v107
                            _ = v108
                        } else {
                            if v1 {
                                v109 := GenS9{F0: 2147483647}
                                _ = v109
                            }
                            if (4294967295 != 764) {
                                v110 := GenS9{F0: int32(v3)}
                                v4 = append(v4, "a")
                                _ = v110
                            }
                        }
                        var v111 uint32 = 385
                        _ = v111
                    }
                    v4 = append(v4, (("hello" + "") + "a"))
                    _ = fi104
                }
                for fi112 := 0; fi112 < 16; fi112++ {
                    v97 -= v97
                    _ = fi112
                }
            }
            fmt.Println(v3, v98)
            for fi113 := 0; fi113 < 19; fi113++ {
                var v114 uint64 = (v5 >> ((4294967295) & 63))
                v115 := GenS3{F0: 931, F1: v1}
                if ("" != ("hello" + "a")) {
                    v3 += 0
                    var v116 uint64 = 834
                    _ = v116
                } else {
                    if (!v1) {
                        v117 := GenS14{F0: 1}
                        var v118 int = fi113
                        v114 -= 1
                        if true {
                            v118 -= (fi113 | 2147483647)
                            var v119 float32 = 3.14159
                            var v120 uint8 = v3
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "\n\t":
                                v121 := GenS15{F0: 1, F1: 42}
                                _ = v121
                            case "a":
                                v122 := GenS7{F0: 42, F1: -32768, F2: 801}
                                _ = v122
                            }
                            _ = v119
                            _ = v120
                        } else {
                            switch (fi99 * fi113) {
                            case 0:
                                switch (fi99 / (((v118 << ((18446744073709551615) & 63))) | 1)) {
                                case 501:
                                    v123 := GenS10{F0: v3, F1: 42, F2: ("a" + ((("foo bar" + ("foo bar" + "")) + (((((("\n\t" + "foo bar") + "\n\t") + "hello") + "a") + "") + "foo bar")) + "\n\t"))}
                                    _ = v123
                                case 2147483647:
                                    var v124 int64 = -9223372036854775808
                                    if true {
                                        var v125 int64 = (-(v124))
                                        var v126 int = (v118 ^ fi113)
                                        v127 := GenS6{F0: 127, F1: (^(v3))}
                                        _ = v125
                                        _ = v126
                                        _ = v127
                                    } else {
                                        fmt.Println(v2, v98)
                                        var v128 uint = 731
                                        switch (fi113 << ((v128) & 63)) {
                                        case 128:
                                            switch (fi99 / ((701) | 1)) {
                                            case 0:
                                                var v129 float32 = 1e10
                                                _ = v129
                                            case 2147483647:
                                                if len(v4) > 0 {
                                                    v4[0] = "hello"
                                                }
                                            }
                                        case 32767:
                                            var v130 uint8 = (v3 * 201)
                                            _ = v130
                                        default:
                                            v131 := GenS14{F0: 2147483647}
                                            _ = v131
                                        }
                                        _ = v128
                                    }
                                    _ = v124
                                case -407:
                                    var v132 uint8 = (^(v3))
                                    _ = v132
                                }
                            case 128:
                                v133 := GenS14{F0: (v97 ^ -624)}
                                _ = v133
                            case 32768:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v98 = 0.0
                            }
                            _ = len(v4)
                        }
                        v3 = (v3 ^ 1)
                        _ = v117
                        _ = v118
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v97 -= (^(v97))
                    v134 := GenS12{F0: (^(v3)), F1: 18446744073709551615}
                    _ = v134
                }
                v98 += v2
                fmt.Println(v1, v114, v2)
                _ = fi113
                _ = v114
                _ = v115
            }
        }
        _ = fi99
        _ = v100
        _ = v101
    }
    v98 -= (float64((0.1 / 1.0)) * ((float64(1e10) - 1.5) - v98))
    if true {
        switch 931 {
        case -2147483648:
            if ((v2 / v2) == v2) {
                fmt.Println(v3, v2, v1)
                switch "" {
                case "a":
                    v2 -= v2
                case "hello":
                    for fi135 := 0; fi135 < 5; fi135++ {
                        var v136 uint = 86
                        var v137 uint32 = 483
                        _ = fi135
                        _ = v136
                        _ = v137
                    }
                case "foo bar":
                    if (genFn0(2, 2147483647) >= 161) {
                        if ((^(v3)) == (v3 >> ((2) & 63))) {
                            v98 = (v98 / 3.14159)
                            v138 := GenS16{F0: -74}
                            if (!v1) {
                                v139 := GenS12{F0: v3, F1: 18446744073709551615}
                                fmt.Println(v3, v98, v2)
                                _ = v139
                            }
                            _ = v138
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                default:
                    switch "hello" {
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "foo bar":
                        v1 = (-1 >= -1)
                    default:
                        var v140 uint8 = 185
                        _ = v140
                    }
                }
                var v141 uint = 1
                _ = v141
            } else {
                v5 = (v5 - 2)
            }
        case 32767:
            v4 = append(v4, "foo bar")
        case 2147483647:
            v142 := GenS14{F0: 560}
            _ = v142
        default:
            v98 = float64(float32(v98))
        }
        switch "hello" {
        case "foo bar":
            v143 := GenS15{F0: 1, F1: 1}
            _ = v143
        default:
            for fi144 := 0; fi144 < 4; fi144++ {
                if true {
                    var v145 uint64 = (^(v5))
                    switch "a" {
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                    case "hello":
                        if v1 {
                            if ((v97 & -412) < (v97 - (v97 + v97))) {
                                var v146 uint32 = 265
                                _ = v146
                            }
                        } else {
                            for fi147 := 0; fi147 < 4; fi147++ {
                                fmt.Println(fi147, v5)
                                _ = fi147
                            }
                            for fi148 := 0; fi148 < 7; fi148++ {
                                var v149 uint16 = 255
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 = (float64((1.5 - (0.0 + 3.14159))) / (3.14159 - float64((float32(v98) + -2.25))))
                                _ = fi148
                                _ = v149
                            }
                            for fi150 := 0; fi150 < 6; fi150++ {
                                if ((!v1) && ("hello" < ("a" + (("foo bar" + "\n\t") + "foo bar")))) {
                                    var v151 int32 = (v97 + v97)
                                    if (2 != 0) {
                                        switch "a" {
                                        case "":
                                            v4 = append(v4, "hello")
                                        case "hello":
                                            var v152 uint = 713
                                            _ = v152
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v153 := GenS17{F0: (v151 * (-(v151))), F1: v1, F2: "hello"}
                                        _ = v153
                                    }
                                    _ = v151
                                } else {
                                    var v154 uint = 255
                                    v155 := GenS18{F0: (722 > 650)}
                                    _ = v154
                                    _ = v155
                                }
                                _ = fi150
                            }
                        }
                    case "foo bar":
                        var v156 bool = (68 >= -30)
                        _ = v156
                    default:
                        v3 = (v3 | (v3 / ((v3) | 1)))
                    }
                    if (65535 == 65535) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (false || true) {
                            if (97 >= 255) {
                                if false {
                                    v157 := GenS17{F0: -2147483648, F1: v1, F2: "hello"}
                                    if (v1 || (2147483647 > 65535)) {
                                        var v158 uint = 255
                                        var v159 string = ("a" + "hello")
                                        _ = v158
                                        _ = v159
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v157
                                } else {
                                    v97 += (v97 & 0)
                                }
                            }
                            v5 -= (v5 * v145)
                        }
                        fmt.Println(v2)
                    } else {
                        if v1 {
                            for fi160 := 0; fi160 < 6; fi160++ {
                                switch (("" + "\n\t") + "\n\t") {
                                case "\n\t":
                                    fmt.Println(v98, fi160)
                                    if v1 {
                                        fmt.Println(v5, v98)
                                        var v161 uint16 = 409
                                        _ = v161
                                    }
                                default:
                                    v162 := GenS7{F0: 575, F1: 385, F2: 642}
                                    _ = v162
                                }
                                _ = fi160
                            }
                            v145 = uint64(v3)
                            if (int32(v3) > v97) {
                                for fi163 := 0; fi163 < 15; fi163++ {
                                    if (v2 == 3.14159) {
                                        v4 = append(v4, "\n\t")
                                        fmt.Println(v2)
                                        v5 -= (^(v145))
                                    }
                                    _ = fi163
                                }
                                for fi164 := 0; fi164 < 2; fi164++ {
                                    var v165 int = (fi144 + fi164)
                                    fmt.Println(fi164, v98)
                                    v166 := GenS19{F0: "foo bar", F1: 315}
                                    _ = fi164
                                    _ = v165
                                    _ = v166
                                }
                            } else {
                                var v167 uint = 1
                                _ = len(v4)
                                _ = v167
                            }
                        } else {
                            var v168 float32 = 1e10
                            if (1 <= 9223372036854775807) {
                                v2 = 1e10
                                switch fi144 {
                                case 893:
                                    fmt.Println(v3)
                                case 0:
                                    var v169 float32 = -2.25
                                    _ = v169
                                default:
                                    fmt.Println(v5, v97, v2)
                                }
                            }
                            if (v1 && (138 == v3)) {
                                var v170 int8 = -118
                                v1 = v1
                                v171 := GenS18{F0: ((v3 >> ((874) & 63)) > (v3 | v3))}
                                _ = v170
                                _ = v171
                            }
                            _ = v168
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = v145
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if v1 {
                        _ = len(v4)
                        v3 = 24
                    } else {
                        v172 := GenS17{F0: v97, F1: (v1 && v1), F2: ("\n\t" + (((("foo bar" + ("a" + "hello")) + (("hello" + "hello") + ("" + "hello"))) + "foo bar") + ("\n\t" + ("a" + ("" + "\n\t")))))}
                        v5 = v5
                        _ = v172
                    }
                    for fi173 := 0; fi173 < 18; fi173++ {
                        v174 := GenS20{F0: 42, F1: 1}
                        _ = fi173
                        _ = v174
                    }
                }
                _ = fi144
            }
        }
    } else {
        var v175 bool = (!(186 > 2147483648))
        for fi176 := 0; fi176 < 20; fi176++ {
            v175 = v175
            _ = fi176
        }
        _ = v175
    }
    v2 += v98
    switch "" {
    case "\n\t":
        for fi177 := 0; fi177 < 12; fi177++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi177
        }
    case "a":
        var v178 uint32 = 2
        _ = v178
    default:
        for fi179 := 0; fi179 < 2; fi179++ {
            fmt.Println(v98, v5)
            _ = fi179
        }
    }
    v180 := GenS21{F0: (v97 * (v97 & v97)), F1: 2, F2: 42}
    fmt.Println(v97)
    v2 = (1e10 * v98)
    v2 = v98
    if v1 {
        v181 := GenS22{F0: (4294967295 == 4294967295), F1: -1}
        switch 195 {
        case 4294967295:
            if true {
                var v182 int64 = 9223372036854775807
                _ = v182
            } else {
                var v183 int = 1
                _ = v183
            }
        case 32767:
            _ = len(v4)
        }
        _ = v181
    }
    for fi184 := 0; fi184 < 3; fi184++ {
        v185 := GenS11{F0: -1}
        v2 -= 1.5
        _ = fi184
        _ = v185
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v180
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v97
    _ = v98
}
