package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int; F2 int32 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int; F2 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 string }
type GenS11 struct { F0 int8; F1 uint32 }
type GenS12 struct { F0 uint; F1 int16 }
type GenS13 struct { F0 int8; F1 int64 }
type GenS14 struct { F0 uint16; F1 uint16 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int16; F1 int }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 bool; F1 int32; F2 int32 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 bool; F1 int64; F2 int }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 bool; F1 int; F2 int64 }
type GenS23 struct { F0 int64; F1 bool; F2 string }
type GenS24 struct { F0 int16; F1 bool }
type GenS25 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS26 struct { F0 int32; F1 int32; F2 int }
type GenS27 struct { F0 uint64 }
type GenS28 struct { F0 int32; F1 int32; F2 uint32 }
type GenS29 struct { F0 bool }
type GenS30 struct { F0 int8; F1 int8 }
type GenS31 struct { F0 int32; F1 uint; F2 int16 }
type GenS32 struct { F0 bool }
type GenS33 struct { F0 int8; F1 uint64; F2 bool }
type GenS34 struct { F0 bool }
type GenS35 struct { F0 int32 }
type GenS36 struct { F0 uint32 }
type GenS37 struct { F0 int8; F1 string; F2 uint8 }
type GenS38 struct { F0 int32 }
type GenS39 struct { F0 int8; F1 int32; F2 uint32 }
type GenS40 struct { F0 uint64 }
type GenS41 struct { F0 bool; F1 bool }
type GenS42 struct { F0 bool }
func genFn0(p0 uint64, p1 uint32, p2 uint32) uint64 {
    v6 := GenS1{F0: 28, F1: 127}
    fmt.Println(p0, p0, p1)
    p2 -= (^(p1))
    p2 = (p1 - 0)
    v7 := GenS2{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 / (((p0 - (p0 >> ((234) & 63)))) | 1))
}
func genFn1() int8 {
    if true {
        var v8 int32 = 42
        var v9 int = -9223372036854775808
        _ = v8
        _ = v9
    }
    var v10 int32 = 1
    _ = v10
    return -73
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 128
    var v3 float64 = 1.5
    var v4 uint64 = 0
    v5 := make(map[string]int)
    for fi11 := 0; fi11 < 8; fi11++ {
        if ((v2 + (v2 << ((42) & 63))) > v2) {
            v4 -= 2
            if (!(359 <= 478)) {
                fmt.Println(v3)
            }
        } else {
            var v12 uint = 65535
            switch "hello" {
            case "":
                switch fi11 {
                case 32767:
                    v13 := v5[""]
                    _ = v13
                case 116:
                    switch (fi11 | (fi11 % (((fi11 ^ fi11)) | 1))) {
                    case 416:
                        fmt.Println(v4, v12, fi11)
                    }
                }
            case "hello":
                for fi14 := 0; fi14 < 4; fi14++ {
                    if true {
                        _ = len(v1)
                    }
                    fmt.Println(v3)
                    v15 := GenS3{F0: "foo bar"}
                    if (42 > 2147483647) {
                        var v16 uint = 0
                        v17 := v5[""]
                        _ = v16
                        _ = v17
                    } else {
                        v3 -= (float64(3.14159) * v3)
                        if false {
                            var v18 int64 = -722
                            v1[390] = (fi14 & 681)
                            if true {
                                var v19 string = "a"
                                if true {
                                    var v20 uint32 = uint32(v2)
                                    _ = v20
                                }
                                _ = v19
                            } else {
                                switch "" {
                                case "hello":
                                    v12 -= uint(v4)
                                }
                            }
                            _ = v18
                        }
                        var v21 float32 = 1.5
                        _ = v21
                    }
                    if true {
                        if ((v12 * 2) > (v12 + v12)) {
                            var v22 int16 = -746
                            _ = v22
                        }
                        if (true || (!((!(-2147483648 > int32(v12))) && false))) {
                            var v23 int64 = int64(v2)
                            var v24 int64 = v23
                            if true {
                                if (127 >= -1) {
                                    if ((1.5 <= -2.25) && true) {
                                        v24 -= (v23 ^ v23)
                                        fmt.Println(v2, fi11, v2)
                                    } else {
                                        v3 = v3
                                        switch (fi14 * 535) {
                                        case 79:
                                            var v25 int64 = (v24 | v23)
                                            _ = v25
                                        default:
                                            v3 = v3
                                        }
                                        var v26 int16 = 42
                                        _ = v26
                                    }
                                    if false {
                                        fmt.Println(fi11)
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            fmt.Println(fi14, v12)
                                        case "foo bar":
                                            fmt.Println(v12)
                                        case "hello":
                                            v27 := GenS3{F0: "foo bar"}
                                            _ = v27
                                        }
                                        if false {
                                            v28 := []int{fi11, (fi11 + (fi14 | 9223372036854775807))}
                                            _ = v28
                                        } else {
                                            switch fi11 {
                                            case -1:
                                                var v29 float32 = (-2.25 / -2.25)
                                                _ = v29
                                            case 4294967295:
                                                if false {
                                                    if (!((fi11 | fi11) != (fi11 + -739))) {
                                                        var v30 int = fi11
                                                        delete(v1, 65535)
                                                        _ = len(v5)
                                                        switch ((("" + "foo bar") + ("foo bar" + ("" + "hello"))) + ((("\n\t" + "a") + "") + ((("hello" + "\n\t") + ("" + "\n\t")) + ("foo bar" + ("hello" + "hello"))))) {
                                                        case "foo bar":
                                                            if true {
                                                                v1[0] = v30
                                                                var v31 bool = (!(true && false))
                                                                v2 += v2
                                                                _ = v31
                                                            } else {
                                                                if (false && (1 != v4)) {
                                                                    v32 := GenS4{F0: (v2 & uint8(v4)), F1: int(v12), F2: -843}
                                                                    v12 -= (^(v12))
                                                                    if false {
                                                                        v33 := GenS5{F0: 2, F1: 32767}
                                                                        var v34 bool = ("a" <= "foo bar")
                                                                        _ = v33
                                                                        _ = v34
                                                                    }
                                                                    _ = v32
                                                                } else {
                                                                    v3 = float64(3.14159)
                                                                    v3 -= 0.1
                                                                }
                                                            }
                                                        case "\n\t":
                                                            if ((false && (((1e10 - float64((float32((1e10 * v3)) + (1e10 * 3.14159)))) / 0.1) >= 3.14159)) && (v12 <= v12)) {
                                                                v4 += v4
                                                                v2 -= v2
                                                            } else {
                                                                v35 := GenS4{F0: 128, F1: fi11, F2: 364}
                                                                _ = v35
                                                            }
                                                        }
                                                        v30 = (fi11 >> (((v12 * (v12 >> ((18446744073709551615) & 63)))) & 63))
                                                        v36 := GenS5{F0: 0, F1: -1}
                                                        fmt.Println(v12, v3)
                                                        v37 := v5[("" + "hello")]
                                                        _ = v30
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                } else {
                                                    v38 := []int{4294967295, fi14, -1}
                                                    v39 := GenS1{F0: 0, F1: 127}
                                                    if false {
                                                        v40 := GenS4{F0: 193, F1: fi14, F2: 705}
                                                        _ = v40
                                                    } else {
                                                        switch (fi14 | (-(fi14))) {
                                                        case 149:
                                                            v41 := v5[("\n\t" + "a")]
                                                            _ = v41
                                                        case 42:
                                                            v12 = 793
                                                        case 128:
                                                            v42 := GenS6{F0: (^(v4)), F1: v24, F2: v12}
                                                            _ = v42
                                                        default:
                                                            var v43 float64 = v3
                                                            _ = v43
                                                        }
                                                    }
                                                    _ = v38
                                                    _ = v39
                                                }
                                            case 2147483647:
                                                v4 -= 774
                                            default:
                                                if true {
                                                    v3 = (0.0 * (v3 / v3))
                                                    v44 := GenS7{F0: v4, F1: 0, F2: v2}
                                                    switch (-(fi11)) {
                                                    case -405:
                                                        v4 -= v4
                                                    case -9223372036854775808:
                                                        v45 := GenS7{F0: 997, F1: 138, F2: v2}
                                                        _ = v45
                                                    default:
                                                        v24 += v23
                                                    }
                                                    v46 := GenS4{F0: 120, F1: fi11, F2: 2147483647}
                                                    _ = v44
                                                    _ = v46
                                                }
                                            }
                                            v47 := v1[1]
                                            switch ("foo bar" + "foo bar") {
                                            case "foo bar":
                                                v48 := GenS1{F0: -103, F1: -128}
                                                _ = v48
                                            case "a":
                                                fmt.Println(v24, v47)
                                            case "hello":
                                                if (41 <= 2147483648) {
                                                    if ((("a" + ("hello" + "")) + "foo bar") == "hello") {
                                                        v49 := GenS5{F0: 302, F1: 42}
                                                        _ = v49
                                                    } else {
                                                        v50 := GenS1{F0: genFn1(), F1: 42}
                                                        var v51 uint16 = 26
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    if true {
                                                        if true {
                                                            v52 := GenS7{F0: v4, F1: 65535, F2: (v2 / (((v2 | (^(v2)))) | 1))}
                                                            _ = v52
                                                        } else {
                                                            v53 := []int{(fi11 * 1)}
                                                            v2 = (v2 << ((v12) & 63))
                                                            _ = v53
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        var v54 uint32 = 42
                                                        var v55 uint16 = 96
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                }
                                                v3 = (1e10 / v3)
                                            }
                                            _ = v47
                                        }
                                        v3 += v3
                                    }
                                    _ = len(v5)
                                }
                                _ = len(v5)
                            }
                            _ = v23
                            _ = v24
                        } else {
                            var v56 int16 = 1
                            _ = v56
                        }
                    }
                    _ = fi14
                    _ = v15
                }
            case "foo bar":
                v57 := GenS6{F0: 4294967295, F1: 1, F2: v12}
                _ = v57
            }
            _ = v12
        }
        for fi58 := 0; fi58 < 17; fi58++ {
            v2 = (v2 & 2)
            _ = fi58
        }
        v59 := GenS8{F0: 65535, F1: -9223372036854775808, F2: true}
        _ = fi11
        _ = v59
    }
    if true {
        switch 32768 {
        case 127:
            if true {
                if (v2 <= v2) {
                    var v60 float32 = (-2.25 / 1.0)
                    _ = v60
                } else {
                    v61 := []int{-104, 591}
                    _ = v61
                }
                if ((151 <= (v2 - (v2 * (^(v2))))) || true) {
                    fmt.Println(v3)
                    var v62 int8 = -1
                    for fi63 := 0; fi63 < 5; fi63++ {
                        for fi64 := 0; fi64 < 7; fi64++ {
                            var v65 uint = 255
                            v3 = 1e10
                            v66 := GenS4{F0: v2, F1: fi63, F2: -2147483648}
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                        var v67 string = "a"
                        _ = fi63
                        _ = v67
                    }
                    _ = v62
                } else {
                    v4 = (v4 % (((v4 * (v4 / ((v4) | 1)))) | 1))
                }
                if (18446744073709551615 < 2) {
                    var v68 uint8 = uint8(v4)
                    v2 += v68
                    v69 := GenS1{F0: 1, F1: -1}
                    _ = v68
                    _ = v69
                }
            }
        }
        var v70 float64 = (float64((1e10 / 1.5)) / v3)
        switch "hello" {
        case "foo bar":
            for fi71 := 0; fi71 < 20; fi71++ {
                switch 774 {
                case 128:
                    v72 := GenS9{F0: "\n\t"}
                    _ = v72
                case -640:
                    var v73 uint16 = 65535
                    _ = v73
                default:
                    var v74 uint = 42
                    _ = v74
                }
                for fi75 := 0; fi75 < 20; fi75++ {
                    if false {
                        fmt.Println(v4)
                        v76 := []int{(fi71 * (fi75 & 389))}
                        _ = v76
                    } else {
                        var v77 int64 = 9223372036854775807
                        switch "foo bar" {
                        case "":
                            if false {
                                if (uint32(v4) > 54) {
                                    if false {
                                        fmt.Println(fi75, v3, v70)
                                    } else {
                                        switch 32767 {
                                        case 32768:
                                            v3 = (((((1e10 / (1.5 - (float64(((((1e10 * (0.0 + ((0.0 - (float32(-2.25) - 0.1)) - 0.0))) - (3.14159 + -2.25)) * (0.1 * (3.14159 * ((0.0 * (-2.25 + 3.14159)) + 1.5)))) / 3.14159)) * -2.25))) * (v70 * ((v70 - (1.5 / v70)) * 0.0))) - 3.14159) + v3) / v70)
                                            if false {
                                                var v78 string = "foo bar"
                                                var v79 uint8 = (^(v2))
                                                _ = v78
                                                _ = v79
                                            } else {
                                                v4 = v4
                                            }
                                            var v80 int64 = (-(v77))
                                            var v81 uint64 = (v4 * (v4 << ((42) & 63)))
                                            var v82 float32 = 1.5
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                        case 127:
                                            v2 = 1
                                        case 0:
                                            _ = len(v1)
                                        default:
                                            v70 += v3
                                            v83 := GenS4{F0: 127, F1: fi75, F2: 0}
                                            v70 += (v70 - v3)
                                            _ = v83
                                        }
                                        var v84 uint = 42
                                        _ = v84
                                    }
                                    v70 -= (v70 - (v70 / v70))
                                } else {
                                    switch (fi75 & (fi71 >> ((65535) & 63))) {
                                    case 32768:
                                        if false {
                                            _ = len(v5)
                                            v85 := GenS10{F0: -32768, F1: ("a" + ((("\n\t" + "\n\t") + "hello") + "foo bar"))}
                                            _ = v85
                                        }
                                    }
                                    v3 -= ((v70 - v3) / 3.14159)
                                }
                                var v86 float32 = 1.5
                                var v87 string = ("" + ("foo bar" + "\n\t"))
                                _ = v86
                                _ = v87
                            } else {
                                if false {
                                    var v88 int32 = 1
                                    switch "foo bar" {
                                    case "hello":
                                        var v89 int = (fi75 << ((255) & 63))
                                        _ = v89
                                    case "":
                                        v70 += (0.1 - -2.25)
                                    case "a":
                                        v90 := GenS7{F0: (v4 ^ (^(v4))), F1: 42, F2: v2}
                                        _ = v90
                                    default:
                                        switch (fi75 / ((fi75) | 1)) {
                                        case 42:
                                            switch fi75 {
                                            case -317:
                                                v70 = (((((1.5 * v3) + (v70 * v3)) - 3.14159) * -2.25) - (1e10 * (0.1 * 0.0)))
                                            case 127:
                                                v91 := GenS8{F0: 551, F1: (fi71 | (-(fi71))), F2: true}
                                                _ = v91
                                            }
                                        default:
                                            if (!(0 <= 2147483648)) {
                                                switch ("hello" + "a") {
                                                case "":
                                                    v70 = v3
                                                case "hello":
                                                    if true {
                                                        v70 -= 0.1
                                                    } else {
                                                        fmt.Println(v70, v88, v70)
                                                        var v92 uint32 = 2
                                                        v93 := GenS9{F0: ""}
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                default:
                                                    v77 = int64(v2)
                                                }
                                                v70 = float64(-2.25)
                                                fmt.Println(v2, v88)
                                                v94 := []int{fi75, (^(fi71)), (fi71 / (((-(fi71))) | 1))}
                                                var v95 uint8 = (v2 - 42)
                                                _ = v94
                                                _ = v95
                                            }
                                        }
                                    }
                                    if true {
                                        v3 += (v70 / v3)
                                        v88 = (v88 * v88)
                                    }
                                    _ = v88
                                }
                                v77 -= (v77 / (((-(v77))) | 1))
                            }
                        case "hello":
                            switch "a" {
                            case "foo bar":
                                if ("" == (("" + ("a" + "hello")) + "\n\t")) {
                                    switch "foo bar" {
                                    case "a":
                                        switch "foo bar" {
                                        case "foo bar":
                                            fmt.Println(fi75, v3)
                                        case "\n\t":
                                            v96 := GenS11{F0: 44, F1: 2}
                                            _ = v96
                                        case "hello":
                                            var v97 uint = 420
                                            _ = v97
                                        }
                                    case "hello":
                                        switch "\n\t" {
                                        case "\n\t":
                                            if true {
                                                v77 = v77
                                                if false {
                                                    v98 := GenS6{F0: 710, F1: 9223372036854775807, F2: 4294967295}
                                                    _ = v98
                                                } else {
                                                    var v99 uint16 = 255
                                                    v70 = v3
                                                    _ = v99
                                                }
                                            } else {
                                                var v100 bool = (2147483647 == 42)
                                                v4 = (v4 - 4294967295)
                                                v100 = v100
                                                var v101 int8 = 1
                                                v102 := GenS12{F0: 446, F1: 627}
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                            }
                                        case "":
                                            switch (^(fi75)) {
                                            case 32768:
                                                v103 := GenS13{F0: -1, F1: -293}
                                                _ = v103
                                            }
                                        }
                                    }
                                    v4 -= (v4 % (((v4 >> ((65535) & 63))) | 1))
                                } else {
                                    v70 += float64(3.14159)
                                    var v104 uint32 = 437
                                    v70 += 1.5
                                    _ = v104
                                }
                            default:
                                if false {
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        delete(v1, 42)
                                    case "\n\t":
                                        v77 = (v77 & (v77 ^ v77))
                                        if true {
                                            if false {
                                                var v105 int16 = -32768
                                                v2 = v2
                                                v77 -= v77
                                                var v106 uint16 = 255
                                                var v107 int = int(v106)
                                                _ = v105
                                                _ = v106
                                                _ = v107
                                            }
                                        }
                                        v108 := GenS11{F0: genFn1(), F1: 4294967295}
                                        _ = v108
                                    default:
                                        var v109 int = fi75
                                        _ = v109
                                    }
                                    v4 = (v4 + v4)
                                } else {
                                    v4 = 517
                                    switch "" {
                                    case "\n\t":
                                        v70 -= v70
                                    default:
                                        v2 = (v2 % (((v2 >> ((606) & 63))) | 1))
                                    }
                                    delete(v5, "foo bar")
                                }
                            }
                        }
                        var v110 uint16 = 1000
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            if (2147483647 < 513) {
                                _ = len(v5)
                                v111 := GenS14{F0: v110, F1: v110}
                                _ = v111
                            }
                        case "foo bar":
                            switch (fi75 & 719) {
                            case 9223372036854775807:
                                v2 -= v2
                            case 32767:
                                v70 = (1e10 - (v70 * (float64((float32((float64((((0.0 * 0.0) + (1e10 + 3.14159)) + 0.1)) * v3)) * (0.0 - (((1.5 * (1e10 - (0.1 + ((0.0 * ((((1.5 - (-2.25 + (((float32((1.5 + v70)) / 1e10) / (0.0 + -2.25)) - float32((((-2.25 - v3) * v3) - v70))))) / 3.14159) + (1e10 / float32(v3))) + 1e10)) - -2.25)))) + float32(((((v3 + float64(0.0)) * v70) + (1e10 / -2.25)) / (((((((((v3 - v3) - ((v3 - 3.14159) * v70)) - (v70 * -2.25)) - v70) - (v3 * 1e10)) * v3) + 1e10) - (v3 - 0.0)) + v70)))) / -2.25)))) - v3)))
                            default:
                                v112 := GenS15{F0: 255}
                                fmt.Println(fi75, v4, v4)
                                _ = v112
                            }
                        case "":
                            var v113 uint = 255
                            _ = v113
                        }
                        _ = v110
                        _ = v77
                    }
                    if (true && (449 < uint16(v2))) {
                        fmt.Println(fi71, fi75, v4)
                        var v114 uint32 = 1
                        _ = v114
                    } else {
                        v4 = (v4 * 2)
                    }
                    v115 := GenS13{F0: -128, F1: -51}
                    _ = fi75
                    _ = v115
                }
                delete(v5, "\n\t")
                _ = fi71
            }
        case "\n\t":
            var v116 float32 = 1e10
            _ = v116
        }
        _ = v70
    }
    switch 855 {
    case 128:
        var v117 float64 = float64(3.14159)
        _ = v117
    case -9223372036854775808:
        var v118 int16 = 1
        _ = v118
    case 615:
        var v119 int32 = 2147483647
        _ = v119
    default:
        v4 += 0
    }
    if true {
        var v120 int = -47
        if false {
            for fi121 := 0; fi121 < 13; fi121++ {
                delete(v1, uint32(v2))
                var v122 float32 = (1e10 * ((float32(3.14159) / (((1e10 + 1.5) * -2.25) - 0.1)) / (0.0 - 1e10)))
                _ = fi121
                _ = v122
            }
            delete(v5, ("" + ""))
            if true {
                fmt.Println(v3, v3, v3)
            } else {
                v4 += v4
            }
        }
        _ = v120
    } else {
        var v123 uint8 = (^(v2))
        var v124 uint32 = 381
        for fi125 := 0; fi125 < 8; fi125++ {
            var v126 uint32 = (v124 / ((v124) | 1))
            fmt.Println(v123)
            _ = fi125
            _ = v126
        }
        _ = v123
        _ = v124
    }
    var v127 int32 = -2147483648
    var v128 int64 = 9223372036854775807
    v129 := GenS12{F0: 0, F1: 42}
    v2 += (v2 % (((^(v2))) | 1))
    switch ("\n\t" + ("foo bar" + ("a" + "a"))) {
    case "\n\t":
        v5["foo bar"] = 824
    case "":
        v130 := GenS16{F0: 1000, F1: -2147483648}
        _ = v130
    case "a":
        for fi131 := 0; fi131 < 1; fi131++ {
            if (0 >= 2) {
                v132 := GenS5{F0: 32767, F1: 32767}
                switch fi131 {
                case -2147483648:
                    switch fi131 {
                    case -9223372036854775808:
                        fmt.Println(fi131, v3)
                    case -2147483648:
                        if false {
                            v133 := GenS8{F0: 0, F1: fi131, F2: true}
                            v128 -= v128
                            fmt.Println(fi131)
                            _ = v133
                        } else {
                            for fi134 := 0; fi134 < 1; fi134++ {
                                var v135 int16 = -1
                                _ = fi134
                                _ = v135
                            }
                            var v136 int8 = 30
                            _ = v136
                        }
                    default:
                        v137 := GenS9{F0: "a"}
                        _ = v137
                    }
                case 42:
                    if true {
                        for fi138 := 0; fi138 < 1; fi138++ {
                            switch "hello" {
                            case "":
                                v139 := []int{128, int(v127)}
                                _ = v139
                            default:
                                v4 = v4
                            }
                            v2 += (v2 / ((v2) | 1))
                            if false {
                                fmt.Println(v4, v127)
                                switch fi138 {
                                case 127:
                                    switch ((("hello" + "foo bar") + "\n\t") + "foo bar") {
                                    case "a":
                                        switch (fi131 | (^(fi138))) {
                                        case -863:
                                            v127 = v127
                                        case 518:
                                            var v140 uint16 = 42
                                            _ = v140
                                        case 0:
                                            switch (^(fi131)) {
                                            case 0:
                                                switch "hello" {
                                                case "a":
                                                    if true {
                                                        var v141 int32 = v127
                                                        _ = v141
                                                    }
                                                    var v142 uint16 = 373
                                                    _ = v142
                                                case "foo bar":
                                                    var v143 int16 = 1000
                                                    _ = v143
                                                case "\n\t":
                                                    if false {
                                                        if (42 == 1000) {
                                                            var v144 uint32 = 362
                                                            if (int(v4) < fi131) {
                                                                fmt.Println(v3, fi131, fi138)
                                                                v4 = (v4 ^ (v4 % (((v4 / ((v4) | 1))) | 1)))
                                                                if ((v128 > 369) || (1 < 1)) {
                                                                    fmt.Println(v128, v144, fi131)
                                                                    v145 := GenS12{F0: 4294967295, F1: 1}
                                                                    _ = v145
                                                                } else {
                                                                    v146 := []int{128, 149}
                                                                    _ = v146
                                                                }
                                                                fmt.Println(v4, v2, v3)
                                                            }
                                                            _ = v144
                                                        } else {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v147 uint8 = (^(v2))
                                                                _ = v147
                                                            case "foo bar":
                                                                v127 += v127
                                                            case "":
                                                                v2 += v2
                                                            default:
                                                                v148 := GenS4{F0: v2, F1: fi131, F2: 71}
                                                                _ = v148
                                                            }
                                                            fmt.Println(v2)
                                                            fmt.Println(v2, v128)
                                                        }
                                                    }
                                                default:
                                                    v2 = uint8(v4)
                                                }
                                            case 4294967295:
                                                if true {
                                                    delete(v5, "foo bar")
                                                    v149 := GenS17{F0: (v127 + v127)}
                                                    _ = v149
                                                } else {
                                                    switch "\n\t" {
                                                    case "a":
                                                        if ((true || (!true)) || false) {
                                                            v3 = v3
                                                            if false {
                                                                v150 := GenS18{F0: (1 != -79), F1: -2147483648, F2: v127}
                                                                var v151 string = ""
                                                                switch v151 {
                                                                case "foo bar":
                                                                    var v152 int32 = (v127 | (v127 + v127))
                                                                    _ = v152
                                                                case "\n\t":
                                                                    if false {
                                                                        v153 := []int{32767, fi131}
                                                                        if (!(false && (1 == 11))) {
                                                                            var v154 bool = false
                                                                            var v155 int8 = 42
                                                                            _ = v154
                                                                            _ = v155
                                                                        }
                                                                        switch "a" {
                                                                        case "":
                                                                            fmt.Println(v128, v151)
                                                                        case "hello":
                                                                            if (!(0.0 < 3.14159)) {
                                                                                var v156 int64 = 1
                                                                                _ = v156
                                                                            } else {
                                                                                fmt.Println(v3, v4)
                                                                            }
                                                                        default:
                                                                            fmt.Println(v4)
                                                                        }
                                                                        _ = v153
                                                                    } else {
                                                                        var v157 uint8 = uint8(v4)
                                                                        var v158 int16 = 0
                                                                        v2 -= (v157 + v157)
                                                                        if (!(18446744073709551615 != 4294967295)) {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                var v159 int32 = 462
                                                                                _ = v159
                                                                            }
                                                                            var v160 int16 = (v158 >> ((18446744073709551615) & 63))
                                                                            v151 = ""
                                                                            v127 -= (^(v127))
                                                                            v161 := GenS19{F0: false}
                                                                            if true {
                                                                                v151 = v151
                                                                            } else {
                                                                                v5[(("foo bar" + v151) + v151)] = (fi138 * int(v128))
                                                                                v162 := GenS9{F0: v151}
                                                                                _ = v162
                                                                            }
                                                                            _ = v160
                                                                            _ = v161
                                                                        } else {
                                                                            var v163 int64 = v128
                                                                            if false {
                                                                                var v164 int = fi131
                                                                                switch -662 {
                                                                                case 4294967295:
                                                                                    var v165 uint16 = 667
                                                                                    if false {
                                                                                        if (((v163 - v128) >= 42) || false) {
                                                                                            v127 += v127
                                                                                            v158 = 42
                                                                                        }
                                                                                        delete(v1, 1)
                                                                                    }
                                                                                    _ = v165
                                                                                case 377:
                                                                                    switch (("foo bar" + v151) + (v151 + v151)) {
                                                                                    case "a":
                                                                                        var v166 float64 = 3.14159
                                                                                        _ = v166
                                                                                    case "\n\t":
                                                                                        v158 = (v158 | (v158 & 60))
                                                                                    case "":
                                                                                        if false {
                                                                                            switch v151 {
                                                                                            case "":
                                                                                                if false {
                                                                                                    if ((fi138 | (fi138 | -1)) <= fi138) {
                                                                                                        v167 := GenS4{F0: v157, F1: (fi131 % (((fi131 << ((42) & 63))) | 1)), F2: int32(fi131)}
                                                                                                        _ = v167
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v168 int = fi131
                                                                                                _ = v168
                                                                                            case "hello":
                                                                                                if false {
                                                                                                    if (1e10 != (1.5 + -2.25)) {
                                                                                                        if (true || (2 <= 42)) {
                                                                                                            switch "\n\t" {
                                                                                                            case "a":
                                                                                                                v2 = (v157 * v157)
                                                                                                            case "hello":
                                                                                                                v163 = v163
                                                                                                            case "\n\t":
                                                                                                                v128 += v128
                                                                                                                var v169 uint64 = (v4 << ((1) & 63))
                                                                                                                v163 += (-(v128))
                                                                                                                var v170 int = fi138
                                                                                                                v171 := v1[2147483647]
                                                                                                                var v172 int32 = (-(v127))
                                                                                                                _ = v169
                                                                                                                _ = v170
                                                                                                                _ = v171
                                                                                                                _ = v172
                                                                                                            default:
                                                                                                                fmt.Println(v163, v128, fi138)
                                                                                                            }
                                                                                                            v173 := []int{fi131, (v164 / ((fi138) | 1)), (-(fi138)), (fi131 % ((1) | 1))}
                                                                                                            if (true || (((!true) || true) || (genFn1() > 42))) {
                                                                                                                v174 := GenS20{F0: false, F1: v128, F2: (^(fi138))}
                                                                                                                _ = len(v173)
                                                                                                                _ = v174
                                                                                                            }
                                                                                                            v3 = v3
                                                                                                            _ = v173
                                                                                                        } else {
                                                                                                            delete(v1, 2147483648)
                                                                                                        }
                                                                                                        var v175 int8 = int8(v128)
                                                                                                        if (!false) {
                                                                                                            var v176 int16 = 697
                                                                                                            fmt.Println(fi131, v163, v127)
                                                                                                            fmt.Println(v163, v157)
                                                                                                            _ = v176
                                                                                                        } else {
                                                                                                            var v177 int8 = v175
                                                                                                            _ = v177
                                                                                                        }
                                                                                                        v1[2147483647] = 32767
                                                                                                        _ = v175
                                                                                                    }
                                                                                                } else {
                                                                                                    var v178 uint64 = 326
                                                                                                    _ = v178
                                                                                                }
                                                                                            default:
                                                                                                var v179 uint64 = 195
                                                                                                _ = v179
                                                                                            }
                                                                                        } else {
                                                                                            v158 += (v158 - (v158 * (v158 | (v158 << ((255) & 63)))))
                                                                                        }
                                                                                    default:
                                                                                        v4 -= genFn0((^(v4)), 2147483648, 42)
                                                                                    }
                                                                                default:
                                                                                    _ = len(v1)
                                                                                }
                                                                                switch v151 {
                                                                                case "a":
                                                                                    if true {
                                                                                        v180 := GenS12{F0: 220, F1: v158}
                                                                                        _ = v180
                                                                                    }
                                                                                default:
                                                                                    v2 = 126
                                                                                }
                                                                                var v181 uint32 = 782
                                                                                var v182 string = v151
                                                                                _ = v164
                                                                                _ = v181
                                                                                _ = v182
                                                                            } else {
                                                                                if (v127 <= v127) {
                                                                                    v4 = (v4 + 513)
                                                                                    if true {
                                                                                        v127 = (^(v127))
                                                                                        fmt.Println(v3, fi131, v128)
                                                                                        v183 := GenS15{F0: 42}
                                                                                        _ = v183
                                                                                    }
                                                                                    v184 := GenS4{F0: (v157 | v2), F1: -736, F2: (v127 * 584)}
                                                                                    var v185 string = v151
                                                                                    _ = v184
                                                                                    _ = v185
                                                                                } else {
                                                                                    v186 := GenS21{F0: int32(v163)}
                                                                                    var v187 int8 = 0
                                                                                    _ = v186
                                                                                    _ = v187
                                                                                }
                                                                                switch v151 {
                                                                                case "\n\t":
                                                                                    v4 -= (v4 - genFn0(4294967295, 99, 176))
                                                                                case "":
                                                                                    var v188 float64 = 0.0
                                                                                    _ = v188
                                                                                default:
                                                                                    v1[2147483648] = (fi131 + (fi138 % (((fi131 << ((65535) & 63))) | 1)))
                                                                                }
                                                                            }
                                                                            _ = v163
                                                                        }
                                                                        v4 = 42
                                                                        fmt.Println(v158, v2, v3)
                                                                        if (!((v158 > (-(v158))) && (1 < uint16(v157)))) {
                                                                            if (!false) {
                                                                                var v189 string = ""
                                                                                var v190 int8 = 18
                                                                                v2 += 149
                                                                                _ = v189
                                                                                _ = v190
                                                                            } else {
                                                                                v2 -= (v157 % ((v157) | 1))
                                                                                v2 = (^(v2))
                                                                            }
                                                                            var v191 int32 = -2147483648
                                                                            v192 := []int{(fi131 & fi138), (^(fi131)), fi138}
                                                                            _ = v191
                                                                            _ = v192
                                                                        } else {
                                                                            switch fi131 {
                                                                            case 42:
                                                                                fmt.Println(v128, fi131, v4)
                                                                            default:
                                                                                v193 := []int{(-(fi138)), fi131, (fi138 - fi138)}
                                                                                _ = v193
                                                                            }
                                                                            switch fi138 {
                                                                            case -2147483648:
                                                                                if (!true) {
                                                                                    switch (fi138 ^ 705) {
                                                                                    case -2147483648:
                                                                                        if true {
                                                                                            v194 := GenS22{F0: ((0.0 / 1e10) != (1e10 * (((0.0 + (0.1 * float32((0.0 - float64(-2.25))))) + -2.25) + 1e10))), F1: 128, F2: v128}
                                                                                            var v195 uint64 = uint64(v157)
                                                                                            fmt.Println(fi131, fi131, v3)
                                                                                            if true {
                                                                                                v196 := GenS23{F0: (v128 | v128), F1: (!false), F2: v151}
                                                                                                v197 := GenS12{F0: 42, F1: (v158 >> ((425) & 63))}
                                                                                                _ = v196
                                                                                                _ = v197
                                                                                            } else {
                                                                                                var v198 uint32 = 2147483647
                                                                                                v199 := v1[2]
                                                                                                if true {
                                                                                                    v200 := GenS24{F0: (v158 + 42), F1: true}
                                                                                                    _ = v200
                                                                                                } else {
                                                                                                    if ((^(v127)) == v127) {
                                                                                                        v4 += v195
                                                                                                        v201 := GenS13{F0: -1, F1: (v128 - (v128 * (v128 ^ (^(v128)))))}
                                                                                                        _ = v201
                                                                                                    }
                                                                                                    v202 := []int{(fi131 + (fi138 << ((18446744073709551615) & 63))), (v199 | v199)}
                                                                                                    _ = v202
                                                                                                }
                                                                                                v158 -= (v158 >> ((1) & 63))
                                                                                                var v203 uint32 = (v198 * (v198 >> ((1) & 63)))
                                                                                                switch (fi138 * 409) {
                                                                                                case 32768:
                                                                                                    v204 := v5["\n\t"]
                                                                                                    _ = v204
                                                                                                default:
                                                                                                    var v205 uint64 = 2
                                                                                                    _ = v205
                                                                                                }
                                                                                                _ = v198
                                                                                                _ = v199
                                                                                                _ = v203
                                                                                            }
                                                                                            var v206 int64 = v128
                                                                                            _ = v194
                                                                                            _ = v195
                                                                                            _ = v206
                                                                                        }
                                                                                    case -1:
                                                                                        v207 := GenS2{F0: (v151 + v151)}
                                                                                        _ = v207
                                                                                    default:
                                                                                        fmt.Println(v128, v2, fi138)
                                                                                    }
                                                                                    var v208 uint64 = (v4 >> ((1) & 63))
                                                                                    _ = v208
                                                                                } else {
                                                                                    var v209 bool = (!true)
                                                                                    var v210 uint32 = 2
                                                                                    var v211 int16 = (-(v158))
                                                                                    switch ("\n\t" + "a") {
                                                                                    case "a":
                                                                                        if v209 {
                                                                                            v212 := GenS14{F0: 0, F1: 681}
                                                                                            v213 := []int{(fi131 & (^(fi138))), (fi138 & fi131), (fi138 & -936)}
                                                                                            _ = v212
                                                                                            _ = v213
                                                                                        }
                                                                                        var v214 uint = 735
                                                                                        var v215 string = v151
                                                                                        _ = v214
                                                                                        _ = v215
                                                                                    case "foo bar":
                                                                                        fmt.Println(v128, v128)
                                                                                    case "":
                                                                                        v216 := GenS25{F0: v210, F1: 255, F2: 0}
                                                                                        _ = v216
                                                                                    default:
                                                                                        v217 := GenS19{F0: v209}
                                                                                        _ = v217
                                                                                    }
                                                                                    fmt.Println(v158)
                                                                                    v3 -= 1.5
                                                                                    switch fi131 {
                                                                                    case 4294967295:
                                                                                        var v218 int = (-(fi131))
                                                                                        _ = v218
                                                                                    default:
                                                                                        v219 := GenS3{F0: "a"}
                                                                                        _ = v219
                                                                                    }
                                                                                    _ = v209
                                                                                    _ = v210
                                                                                    _ = v211
                                                                                }
                                                                                var v220 uint64 = 2
                                                                                if true {
                                                                                    v128 = (v128 / ((0) | 1))
                                                                                    _ = len(v1)
                                                                                    if true {
                                                                                        if (!(!((!(!((v220 | genFn0(v4, 65535, 962)) != v4))) || (true || true)))) {
                                                                                            var v221 bool = (("foo bar" + "") > (v151 + ("hello" + (v151 + v151))))
                                                                                            var v222 bool = (32768 <= -2147483648)
                                                                                            _ = v221
                                                                                            _ = v222
                                                                                        } else {
                                                                                            v128 = v128
                                                                                            var v223 float32 = (0.0 * 1.5)
                                                                                            var v224 uint8 = v2
                                                                                            v225 := GenS26{F0: v127, F1: v127, F2: -660}
                                                                                            _ = v223
                                                                                            _ = v224
                                                                                            _ = v225
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v158)
                                                                                    }
                                                                                    var v226 int16 = -1
                                                                                    v226 = 1
                                                                                    if ((v2 < (v2 / (((v157 * (v2 + (v157 ^ v157)))) | 1))) || true) {
                                                                                        if (1 <= 2147483647) {
                                                                                            v227 := GenS25{F0: 249, F1: 368, F2: 311}
                                                                                            var v228 uint = 2
                                                                                            v229 := v1[1]
                                                                                            _ = v227
                                                                                            _ = v228
                                                                                            _ = v229
                                                                                        }
                                                                                        v226 -= 1000
                                                                                    }
                                                                                    _ = v226
                                                                                }
                                                                                _ = v220
                                                                            case 4294967295:
                                                                                v151 = v151
                                                                            }
                                                                            if false {
                                                                                v2 -= 128
                                                                            }
                                                                            var v230 int = fi131
                                                                            _ = v230
                                                                        }
                                                                        var v231 bool = (!(fi138 < (^(fi138))))
                                                                        if v231 {
                                                                            v231 = (42 < v127)
                                                                            fmt.Println(v127)
                                                                            v232 := GenS1{F0: 127, F1: genFn1()}
                                                                            fmt.Println(v127, v231)
                                                                            var v233 bool = v231
                                                                            _ = v232
                                                                            _ = v233
                                                                        } else {
                                                                            v234 := v1[0]
                                                                            if v231 {
                                                                                v235 := GenS20{F0: false, F1: v128, F2: fi138}
                                                                                _ = v235
                                                                            } else {
                                                                                fmt.Println(v234)
                                                                                v127 = (v127 * (v127 % (((v127 ^ v127)) | 1)))
                                                                            }
                                                                            fmt.Println(v2, v4)
                                                                            _ = v234
                                                                        }
                                                                        var v236 int64 = 9223372036854775807
                                                                        if v231 {
                                                                            if v231 {
                                                                                if v231 {
                                                                                    switch v151 {
                                                                                    case "hello":
                                                                                        var v237 uint32 = 505
                                                                                        _ = v237
                                                                                    case "a":
                                                                                        if (994 >= 65535) {
                                                                                            if (uint(v4) < 699) {
                                                                                                fmt.Println(v151, v157)
                                                                                                var v238 float64 = (v3 * (v3 * v3))
                                                                                                if v231 {
                                                                                                    v239 := GenS7{F0: (v4 >> ((255) & 63)), F1: 761, F2: v2}
                                                                                                    var v240 int32 = (-(v127))
                                                                                                    if v231 {
                                                                                                        if ("foo bar" < "\n\t") {
                                                                                                            v240 -= 2147483647
                                                                                                            v128 += int64(v127)
                                                                                                            if v231 {
                                                                                                                if (!v231) {
                                                                                                                    var v241 uint = 1
                                                                                                                    v242 := GenS27{F0: (v4 % ((42) | 1))}
                                                                                                                    v4 = (v4 & (v4 * genFn0((v4 + 4294967295), 42, 2147483648)))
                                                                                                                    if (v127 > v240) {
                                                                                                                        v151 = "\n\t"
                                                                                                                        var v243 uint16 = 852
                                                                                                                        if ((0.1 - 1.5) <= 0.1) {
                                                                                                                            var v244 int64 = -1
                                                                                                                            var v245 uint32 = 159
                                                                                                                            fmt.Println(v236)
                                                                                                                            _ = v244
                                                                                                                            _ = v245
                                                                                                                        } else {
                                                                                                                            var v246 uint16 = uint16(v4)
                                                                                                                            _ = v246
                                                                                                                        }
                                                                                                                        _ = v243
                                                                                                                    } else {
                                                                                                                        fmt.Println(v128)
                                                                                                                        v158 -= (v158 + 32767)
                                                                                                                        fmt.Println(v4, fi138)
                                                                                                                    }
                                                                                                                    _ = v241
                                                                                                                    _ = v242
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v238 -= -2.25
                                                                                                                        var v247 uint8 = (v157 + v2)
                                                                                                                        if v231 {
                                                                                                                            v248 := []int{fi138, fi138, 42, fi138}
                                                                                                                            v157 += (v2 / ((v2) | 1))
                                                                                                                            _ = v248
                                                                                                                        } else {
                                                                                                                            if v231 {
                                                                                                                                v4 -= (v4 >> ((1) & 63))
                                                                                                                                v249 := GenS26{F0: (v240 / ((v127) | 1)), F1: int32(v247), F2: fi131}
                                                                                                                                _ = v249
                                                                                                                            } else {
                                                                                                                                v127 = -2147483648
                                                                                                                                v250 := GenS28{F0: (^(v240)), F1: v127, F2: 1}
                                                                                                                                var v251 int32 = -2147483648
                                                                                                                                _ = v250
                                                                                                                                _ = v251
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v252 := []int{(-(fi138)), (fi131 | -367), int(v158), fi138}
                                                                                                                        _ = v247
                                                                                                                        _ = v252
                                                                                                                    }
                                                                                                                    if ((1 <= uint16(v2)) || ((!v231) && true)) {
                                                                                                                        if (v231 && (v231 && ((!((v151 + v151) <= v151)) || v231))) {
                                                                                                                            v253 := GenS21{F0: 774}
                                                                                                                            var v254 int16 = (v158 << ((4294967295) & 63))
                                                                                                                            switch fi131 {
                                                                                                                            case -638:
                                                                                                                                v238 -= v3
                                                                                                                            case -48:
                                                                                                                                v158 = v158
                                                                                                                            }
                                                                                                                            _ = v253
                                                                                                                            _ = v254
                                                                                                                        }
                                                                                                                        if (273 > 2) {
                                                                                                                            switch v151 {
                                                                                                                            case "hello":
                                                                                                                                fmt.Println(v236, v3, v240)
                                                                                                                            case "a":
                                                                                                                                _ = len(v1)
                                                                                                                            case "\n\t":
                                                                                                                                switch fi138 {
                                                                                                                                case 9223372036854775807:
                                                                                                                                    v127 = -1
                                                                                                                                case 128:
                                                                                                                                    v255 := []int{}
                                                                                                                                    _ = v255
                                                                                                                                case 2147483647:
                                                                                                                                    fmt.Println(fi138, v238)
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v256 float64 = 1.5
                                                                                                                            v5[v151] = fi138
                                                                                                                            _ = v256
                                                                                                                        }
                                                                                                                        fmt.Println(v236)
                                                                                                                        var v257 float32 = 0.0
                                                                                                                        if v231 {
                                                                                                                            if v231 {
                                                                                                                                var v258 string = (v151 + (("hello" + v151) + v151))
                                                                                                                                _ = v258
                                                                                                                            } else {
                                                                                                                                if (567 <= 32767) {
                                                                                                                                    var v259 uint32 = 901
                                                                                                                                    fmt.Println(v158)
                                                                                                                                    var v260 string = v151
                                                                                                                                    _ = v259
                                                                                                                                    _ = v260
                                                                                                                                } else {
                                                                                                                                    v261 := GenS29{F0: (!true)}
                                                                                                                                    var v262 uint64 = v4
                                                                                                                                    fmt.Println(v236, v236, v240)
                                                                                                                                    _ = v261
                                                                                                                                    _ = v262
                                                                                                                                }
                                                                                                                            }
                                                                                                                            if (!((((v257 - v257) * v257) - (v257 + (v257 / v257))) < 0.0)) {
                                                                                                                                v158 -= v158
                                                                                                                                v128 -= -649
                                                                                                                            } else {
                                                                                                                                v2 -= 139
                                                                                                                                var v263 uint32 = 42
                                                                                                                                _ = v263
                                                                                                                            }
                                                                                                                            v238 = (v238 * (3.14159 * 1e10))
                                                                                                                        }
                                                                                                                        v264 := GenS30{F0: 0, F1: genFn1()}
                                                                                                                        fmt.Println(fi138, v157, v257)
                                                                                                                        _ = v257
                                                                                                                        _ = v264
                                                                                                                    } else {
                                                                                                                        fmt.Println(fi131, fi131, v240)
                                                                                                                        var v265 uint16 = uint16(v4)
                                                                                                                        _ = v265
                                                                                                                    }
                                                                                                                    var v266 int = (-(fi131))
                                                                                                                    _ = v266
                                                                                                                }
                                                                                                            } else {
                                                                                                                v267 := GenS31{F0: 42, F1: 4294967295, F2: v158}
                                                                                                                v2 -= v157
                                                                                                                var v268 uint16 = 42
                                                                                                                _ = v267
                                                                                                                _ = v268
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v239
                                                                                                    _ = v240
                                                                                                } else {
                                                                                                    v238 -= 1.5
                                                                                                }
                                                                                                var v269 uint32 = 4294967295
                                                                                                fmt.Println(v236)
                                                                                                _ = v238
                                                                                                _ = v269
                                                                                            } else {
                                                                                                fmt.Println(v236, v3, v151)
                                                                                                v158 = (v158 - 32767)
                                                                                            }
                                                                                            v270 := []int{(fi131 & (fi138 * 9223372036854775807)), 2147483647, (-(fi131)), fi131}
                                                                                            _ = v270
                                                                                        } else {
                                                                                            var v271 float64 = (v3 / 3.14159)
                                                                                            v272 := GenS31{F0: (v127 * -120), F1: 65535, F2: v158}
                                                                                            var v273 uint8 = v157
                                                                                            delete(v1, 2147483647)
                                                                                            v128 = -942
                                                                                            v274 := GenS19{F0: v231}
                                                                                            _ = v271
                                                                                            _ = v272
                                                                                            _ = v273
                                                                                            _ = v274
                                                                                        }
                                                                                    }
                                                                                    v236 -= (v236 * -321)
                                                                                    v275 := []int{(fi138 % ((fi131) | 1)), int(v157), fi131, (fi131 ^ 9223372036854775807)}
                                                                                    var v276 string = "\n\t"
                                                                                    _ = v275
                                                                                    _ = v276
                                                                                }
                                                                                v277 := []int{785}
                                                                                switch 32767 {
                                                                                case 9223372036854775807:
                                                                                    v127 = (^(v127))
                                                                                case 32767:
                                                                                    var v278 float32 = 3.14159
                                                                                    v151 = v151
                                                                                    v279 := v5[v151]
                                                                                    _ = v278
                                                                                    _ = v279
                                                                                case 127:
                                                                                    v158 = v158
                                                                                }
                                                                                switch fi131 {
                                                                                case 127:
                                                                                    if (1 > 65535) {
                                                                                        var v280 uint16 = 0
                                                                                        v236 = (v236 | v128)
                                                                                        var v281 int64 = (v236 | int64(v157))
                                                                                        _ = v280
                                                                                        _ = v281
                                                                                    }
                                                                                case 1:
                                                                                    var v282 bool = (!((!v231) && (!v231)))
                                                                                    _ = v282
                                                                                }
                                                                                _ = v277
                                                                            } else {
                                                                                v283 := GenS32{F0: (!false)}
                                                                                if v231 {
                                                                                    switch int(v128) {
                                                                                    case 32768:
                                                                                        v4 = (^(v4))
                                                                                    case -2147483648:
                                                                                        switch (v151 + v151) {
                                                                                        case "":
                                                                                            v284 := GenS33{F0: 98, F1: 26, F2: false}
                                                                                            _ = v284
                                                                                        default:
                                                                                            v236 = int64(v2)
                                                                                            switch v151 {
                                                                                            case "\n\t":
                                                                                                v285 := GenS34{F0: (1e10 >= -2.25)}
                                                                                                _ = v285
                                                                                            case "hello":
                                                                                                if v231 {
                                                                                                    var v286 uint32 = 2
                                                                                                    if v231 {
                                                                                                        v151 = (((("" + "hello") + v151) + v151) + ("" + (v151 + (v151 + "foo bar"))))
                                                                                                        var v287 string = (("hello" + v151) + "a")
                                                                                                        var v288 int32 = (v127 % ((-583) | 1))
                                                                                                        _ = v287
                                                                                                        _ = v288
                                                                                                    } else {
                                                                                                        v3 -= float64((0.1 * ((3.14159 * (1e10 / ((-2.25 / float32((v3 * (float64((-2.25 - 0.0)) + v3)))) / 1.5))) + 1e10)))
                                                                                                    }
                                                                                                    _ = v286
                                                                                                } else {
                                                                                                    v2 += v157
                                                                                                }
                                                                                            case "foo bar":
                                                                                                if (!v231) {
                                                                                                    v289 := GenS4{F0: v157, F1: (fi131 % ((192) | 1)), F2: (v127 << ((4294967295) & 63))}
                                                                                                    var v290 uint8 = v157
                                                                                                    if (v231 && (844 != 0)) {
                                                                                                        if (!v231) {
                                                                                                            v4 += (v4 ^ v4)
                                                                                                        }
                                                                                                    } else {
                                                                                                        if v231 {
                                                                                                            switch v151 {
                                                                                                            case "foo bar":
                                                                                                                fmt.Println(fi138, v236, fi138)
                                                                                                            case "\n\t":
                                                                                                                v231 = (!v231)
                                                                                                            case "":
                                                                                                                v291 := GenS5{F0: 0, F1: v158}
                                                                                                                _ = v291
                                                                                                            default:
                                                                                                                fmt.Println(v151)
                                                                                                            }
                                                                                                            fmt.Println(v128, v231)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v289
                                                                                                    _ = v290
                                                                                                } else {
                                                                                                    v236 = (v128 - int64(v158))
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v127, v157, fi131)
                                                                                            }
                                                                                            if v231 {
                                                                                                if (fi131 > fi131) {
                                                                                                    v292 := GenS35{F0: 2147483647}
                                                                                                    var v293 uint8 = 1
                                                                                                    v294 := v5[v151]
                                                                                                    v295 := GenS36{F0: 192}
                                                                                                    _ = v292
                                                                                                    _ = v293
                                                                                                    _ = v294
                                                                                                    _ = v295
                                                                                                } else {
                                                                                                    var v296 int32 = v127
                                                                                                    _ = v296
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        var v297 float64 = 0.1
                                                                                        _ = v297
                                                                                    }
                                                                                    var v298 float64 = (v3 * (v3 + ((v3 * 0.1) - 1e10)))
                                                                                    _ = v298
                                                                                } else {
                                                                                    v299 := GenS36{F0: 467}
                                                                                    _ = v299
                                                                                }
                                                                                v2 = (v157 / ((v2) | 1))
                                                                                _ = v283
                                                                            }
                                                                            v300 := GenS8{F0: 854, F1: -728, F2: (v3 < v3)}
                                                                            if v231 {
                                                                                var v301 int64 = v236
                                                                                _ = v301
                                                                            } else {
                                                                                _ = len(v5)
                                                                            }
                                                                            _ = v300
                                                                        }
                                                                        if (v231 || (65535 < 42)) {
                                                                            fmt.Println(fi131, v127, v236)
                                                                            fmt.Println(v158)
                                                                        }
                                                                        _ = v157
                                                                        _ = v158
                                                                        _ = v231
                                                                        _ = v236
                                                                    }
                                                                case "a":
                                                                    var v302 int16 = -235
                                                                    _ = v302
                                                                default:
                                                                    if false {
                                                                        if false {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v303 := v5[v151]
                                                                                _ = v303
                                                                            case "\n\t":
                                                                                if (!(false && (!false))) {
                                                                                    fmt.Println(v151, v3, fi131)
                                                                                    var v304 bool = (!true)
                                                                                    v2 = v2
                                                                                    v305 := GenS37{F0: -79, F1: v151, F2: 153}
                                                                                    _ = v304
                                                                                    _ = v305
                                                                                } else {
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        v4 = 4294967295
                                                                                    case "\n\t":
                                                                                        delete(v5, "a")
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v306 string = (v151 + v151)
                                                                                _ = v306
                                                                            }
                                                                        } else {
                                                                            var v307 float32 = float32((v3 * 0.1))
                                                                            _ = v307
                                                                        }
                                                                        v4 += (^(v4))
                                                                    }
                                                                }
                                                                _ = v150
                                                                _ = v151
                                                            }
                                                        } else {
                                                            v4 -= 807
                                                        }
                                                    }
                                                    v308 := GenS38{F0: (v127 | (v127 * (v127 + -2147483648)))}
                                                    v309 := []int{(^(fi131)), (fi138 / (((^(fi138))) | 1))}
                                                    _ = v308
                                                    _ = v309
                                                }
                                            case 32768:
                                                if false {
                                                    v310 := []int{42, fi138, fi138, (fi131 & (fi138 ^ int(v127)))}
                                                    _ = v310
                                                }
                                            }
                                        default:
                                            v311 := GenS24{F0: -1, F1: true}
                                            _ = v311
                                        }
                                        switch ("" + (("\n\t" + "\n\t") + ("hello" + ("\n\t" + "hello")))) {
                                        case "foo bar":
                                            _ = len(v5)
                                        case "hello":
                                            var v312 int = (^(fi131))
                                            _ = v312
                                        case "a":
                                            v3 = 0.1
                                        default:
                                            v2 += 127
                                        }
                                        v127 = (v127 % ((-974) | 1))
                                        fmt.Println(v127, fi131)
                                    default:
                                        v2 -= 128
                                    }
                                    v313 := GenS6{F0: 42, F1: (v128 & (v128 + 1)), F2: 0}
                                    v314 := GenS20{F0: (!false), F1: (-(v128)), F2: -925}
                                    var v315 uint32 = 0
                                    var v316 int = (fi138 | -1)
                                    _ = v313
                                    _ = v314
                                    _ = v315
                                    _ = v316
                                case 1:
                                    if (((1e10 * 1.5) / 0.1) > -2.25) {
                                        if false {
                                            switch fi131 {
                                            case 32768:
                                                v2 = v2
                                            case 1:
                                                var v317 int64 = (-(v128))
                                                _ = v317
                                            case -9223372036854775808:
                                                v3 += v3
                                            default:
                                                var v318 float64 = (-2.25 / (v3 / v3))
                                                _ = v318
                                            }
                                        } else {
                                            var v319 bool = true
                                            _ = v319
                                        }
                                        var v320 bool = (4294967295 != fi138)
                                        var v321 bool = (!(!(32767 == (fi138 + -9223372036854775808))))
                                        _ = v320
                                        _ = v321
                                    } else {
                                        v128 = (v128 | 1)
                                        if (974 >= 32767) {
                                            v3 = v3
                                            if (((1 >= 1) || ("hello" > "hello")) || (false || (!(v2 != (v2 & (v2 ^ uint8(v4))))))) {
                                                switch fi138 {
                                                case 261:
                                                    v322 := GenS39{F0: genFn1(), F1: v127, F2: 675}
                                                    _ = v322
                                                default:
                                                    v323 := []int{-2147483648, (fi131 << ((854) & 63))}
                                                    _ = v323
                                                }
                                                fmt.Println(v4)
                                                v324 := GenS3{F0: ""}
                                                _ = v324
                                            } else {
                                                if true {
                                                    var v325 int8 = -1
                                                    if true {
                                                        if (!true) {
                                                            var v326 int32 = (v127 + (v127 + 347))
                                                            _ = v326
                                                        } else {
                                                            fmt.Println(fi138)
                                                        }
                                                    }
                                                    switch fi131 {
                                                    case 42:
                                                        var v327 uint32 = 2147483647
                                                        switch "" {
                                                        case "":
                                                            fmt.Println(fi138, v327)
                                                        case "\n\t":
                                                            switch "foo bar" {
                                                            case "hello":
                                                                fmt.Println(v127, v2)
                                                            case "a":
                                                                switch fi131 {
                                                                case 9223372036854775807:
                                                                    var v328 uint16 = 103
                                                                    _ = v328
                                                                case 1:
                                                                    v329 := GenS40{F0: genFn0((v4 % (((v4 * (v4 + 1))) | 1)), (v327 / ((4294967295) | 1)), (^(v327)))}
                                                                    _ = v329
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(v3, v2)
                                                            default:
                                                                v2 += (^(v2))
                                                            }
                                                        }
                                                        fmt.Println(fi131, fi138)
                                                        v127 = (v127 << ((uint(v4)) & 63))
                                                        v330 := v5[("hello" + ((("foo bar" + "") + ("hello" + "foo bar")) + ("foo bar" + "a")))]
                                                        v331 := GenS37{F0: genFn1(), F1: "\n\t", F2: v2}
                                                        fmt.Println(v327)
                                                        _ = v327
                                                        _ = v330
                                                        _ = v331
                                                    case 2147483647:
                                                        fmt.Println(fi138)
                                                    default:
                                                        var v332 int8 = -128
                                                        _ = v332
                                                    }
                                                    _ = v325
                                                } else {
                                                    v333 := GenS41{F0: true, F1: (-1 > -32768)}
                                                    var v334 uint16 = 923
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v335 int8 = -1
                                                        _ = v335
                                                    default:
                                                        delete(v5, "")
                                                    }
                                                    _ = v333
                                                    _ = v334
                                                }
                                                v336 := GenS31{F0: v127, F1: 4294967295, F2: 519}
                                                if (1.5 <= ((0.1 * (0.0 - 0.0)) * (0.1 * float32((float64(1.5) * 3.14159))))) {
                                                    if true {
                                                        v337 := GenS37{F0: 127, F1: "foo bar", F2: 42}
                                                        fmt.Println(v2, v127, fi138)
                                                        _ = v337
                                                    } else {
                                                        var v338 uint64 = (^(v4))
                                                        if (v127 > (v127 - (v127 * v127))) {
                                                            var v339 uint8 = v2
                                                            var v340 float64 = v3
                                                            _ = v339
                                                            _ = v340
                                                        } else {
                                                            if false {
                                                                v127 += -1
                                                                fmt.Println(v128)
                                                                if (((1.5 / v3) / ((v3 - (-2.25 * -2.25)) + (v3 + (1e10 - 1.5)))) != v3) {
                                                                    switch fi138 {
                                                                    case -190:
                                                                        if ((!true) || ((-(v128)) < -1)) {
                                                                            var v341 string = "a"
                                                                            var v342 bool = ((true && true) && (((-(v127)) == (v127 | (-(v127)))) || ((v2 | 127) >= v2)))
                                                                            if v342 {
                                                                                delete(v5, "")
                                                                                v343 := GenS6{F0: v338, F1: (v128 << ((131) & 63)), F2: 4294967295}
                                                                                if (!false) {
                                                                                    fmt.Println(v4, v127)
                                                                                    switch ("hello" + (v341 + ((v341 + v341) + v341))) {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v4)
                                                                                    case "a":
                                                                                        v341 = ""
                                                                                    }
                                                                                }
                                                                                _ = v343
                                                                            } else {
                                                                                v3 = v3
                                                                            }
                                                                            _ = v341
                                                                            _ = v342
                                                                        }
                                                                    default:
                                                                        v2 = (v2 >> ((1) & 63))
                                                                        v128 -= (v128 % ((v128) | 1))
                                                                    }
                                                                    v344 := GenS30{F0: 127, F1: -1}
                                                                    v345 := GenS41{F0: (1 != 4294967295), F1: true}
                                                                    if true {
                                                                        v338 -= genFn0(v338, 96, 4294967295)
                                                                    } else {
                                                                        switch ("\n\t" + ("foo bar" + "a")) {
                                                                        case "foo bar":
                                                                            v346 := []int{2147483647, (fi138 >> ((18446744073709551615) & 63)), -2147483648}
                                                                            _ = v346
                                                                        case "a":
                                                                            fmt.Println(v3, v338)
                                                                        case "\n\t":
                                                                            var v347 float64 = -2.25
                                                                            _ = v347
                                                                        default:
                                                                            if true {
                                                                                delete(v1, 2147483648)
                                                                            }
                                                                        }
                                                                        var v348 bool = ((("hello" + "") + "hello") >= ("a" + "\n\t"))
                                                                        v5["a"] = -1
                                                                        _ = v348
                                                                    }
                                                                    _ = v344
                                                                    _ = v345
                                                                } else {
                                                                    v349 := GenS30{F0: 1, F1: 42}
                                                                    var v350 uint16 = 2
                                                                    _ = v349
                                                                    _ = v350
                                                                }
                                                            }
                                                            v2 = (v2 << ((509) & 63))
                                                        }
                                                        _ = v338
                                                    }
                                                    fmt.Println(fi131)
                                                    v5[("\n\t" + "\n\t")] = (fi131 | fi131)
                                                }
                                                _ = v336
                                            }
                                        } else {
                                            v351 := GenS42{F0: false}
                                            _ = v351
                                        }
                                    }
                                case -2147483648:
                                    var v352 uint = 1
                                    _ = v352
                                }
                            } else {
                                v353 := GenS7{F0: (^(v4)), F1: 2, F2: v2}
                                v127 = -493
                                _ = v353
                            }
                            _ = fi138
                        }
                    } else {
                        v354 := GenS1{F0: -1, F1: -1}
                        var v355 uint16 = 42
                        fmt.Println(v355)
                        _ = v354
                        _ = v355
                    }
                case 892:
                    var v356 uint64 = v4
                    _ = v356
                }
                for fi357 := 0; fi357 < 6; fi357++ {
                    delete(v5, ("\n\t" + "foo bar"))
                    delete(v1, 65535)
                    _ = fi357
                }
                _ = v132
            }
            _ = fi131
        }
    default:
        v358 := v1[65535]
        _ = v358
    }
    if (false && (v128 > (v128 ^ (v128 ^ (v128 << ((17) & 63)))))) {
        var v359 int16 = 28
        _ = v359
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v127
    _ = v128
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
