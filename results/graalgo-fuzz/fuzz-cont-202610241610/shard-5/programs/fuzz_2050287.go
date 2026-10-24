package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 bool }
type GenS2 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
type GenS4 struct { F0 int32; F1 bool; F2 uint }
type GenS5 struct { F0 bool; F1 int32; F2 int }
type GenS6 struct { F0 int64; F1 bool; F2 int64 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 uint32; F1 string; F2 uint32 }
type GenS10 struct { F0 int32; F1 int64; F2 uint8 }
type GenS11 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 9223372036854775807
    var v3 uint16 = 32767
    v4 := make(map[uint32]string)
    var v5 uint16 = (v3 % (((v3 % (((v3 ^ (v3 >> ((448) & 63)))) | 1))) | 1))
    v1[("\n\t" + "hello")] = 0
    v2 = (^(v2))
    if true {
        v6 := []int{(v2 >> ((42) & 63))}
        for fi7 := 0; fi7 < 5; fi7++ {
            fmt.Println(v2, v3)
            _ = fi7
        }
        _ = v6
    } else {
        if false {
            for fi8 := 0; fi8 < 15; fi8++ {
                for fi9 := 0; fi9 < 17; fi9++ {
                    v1[""] = fi8
                    fmt.Println(fi9, v5, v3)
                    _ = fi9
                }
                for fi10 := 0; fi10 < 3; fi10++ {
                    if false {
                        if true {
                            fmt.Println(fi10, v3)
                            var v11 bool = false
                            _ = v11
                        } else {
                            if (true || false) {
                                v5 = 255
                            }
                        }
                        fmt.Println(v5, v3, v5)
                        switch fi10 {
                        case 42:
                            switch (fi8 + (fi10 >> ((193) & 63))) {
                            case 719:
                                var v12 int64 = 523
                                if (104 < 1) {
                                    var v13 string = "\n\t"
                                    fmt.Println(v3, fi10, fi8)
                                    var v14 uint = 818
                                    _ = v13
                                    _ = v14
                                } else {
                                    var v15 int64 = (-(v12))
                                    var v16 int16 = 32767
                                    _ = v15
                                    _ = v16
                                }
                                _ = v12
                            case 32767:
                                var v17 uint = 158
                                _ = v17
                            }
                        }
                        switch v2 {
                        case 32767:
                            var v18 float64 = ((-2.25 * 1.5) + 1.5)
                            _ = v18
                        case 9223372036854775807:
                            v19 := GenS1{F0: 0, F1: -1, F2: false}
                            var v20 uint = 544
                            _ = v19
                            _ = v20
                        case -2147483648:
                            v5 -= 65535
                        }
                        var v21 uint64 = 971
                        _ = v21
                    } else {
                        switch 4294967295 {
                        case 290:
                            v22 := []int{v2, v2, (fi10 | 134), fi10}
                            _ = v22
                        default:
                            var v23 float32 = ((0.1 / ((3.14159 * 1.5) + 3.14159)) * 0.1)
                            _ = v23
                        }
                    }
                    var v24 int32 = -1
                    _ = fi10
                    _ = v24
                }
                _ = fi8
            }
        } else {
            v3 = v3
            var v25 int16 = -10
            v25 = (-(v25))
            _ = v25
        }
        var v26 string = ""
        _ = v26
    }
    if (65535 != 0) {
        v3 += 0
        if false {
            if false {
                if false {
                    switch "a" {
                    case "a":
                        fmt.Println(v3)
                    case "\n\t":
                        fmt.Println(v3, v3, v5)
                    case "":
                        v5 = v5
                    default:
                        fmt.Println(v3, v3)
                    }
                    v3 += (v3 << ((642) & 63))
                } else {
                    var v27 int16 = 474
                    fmt.Println(v2, v2, v3)
                    v28 := GenS2{F0: 362, F1: (v5 / ((v5) | 1)), F2: (v3 * (^(v3)))}
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 14; fi29++ {
                    if true {
                        v30 := GenS2{F0: 0, F1: v5, F2: v3}
                        if (false && false) {
                            if false {
                                v31 := []int{fi29, v2, -2147483648, (fi29 % (((fi29 - (^(fi29)))) | 1))}
                                if true {
                                    var v32 int32 = 42
                                    var v33 int16 = -362
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 uint = 258
                                    _ = v34
                                }
                                switch "foo bar" {
                                case "a":
                                    for fi35 := 0; fi35 < 10; fi35++ {
                                        v5 = (^(v3))
                                        var v36 uint16 = v5
                                        _ = fi35
                                        _ = v36
                                    }
                                }
                                _ = v31
                            } else {
                                v5 = v3
                                if (v2 != (v2 % (((fi29 + fi29)) | 1))) {
                                    if true {
                                        v37 := v1[((("a" + "foo bar") + "") + "\n\t")]
                                        var v38 int64 = 42
                                        fmt.Println(v2, v2)
                                        _ = v37
                                        _ = v38
                                    }
                                    v2 = v2
                                    v39 := GenS2{F0: 42, F1: v3, F2: v3}
                                    _ = v39
                                } else {
                                    var v40 int = v2
                                    for fi41 := 0; fi41 < 17; fi41++ {
                                        delete(v4, 351)
                                        _ = fi41
                                    }
                                    _ = v40
                                }
                                v42 := v1["foo bar"]
                                _ = v42
                            }
                            switch v2 {
                            case 32768:
                                if true {
                                    var v43 uint = 255
                                    _ = v43
                                } else {
                                    if true {
                                        v44 := GenS1{F0: 42, F1: -2147483648, F2: ((float64(3.14159) * 0.0) < (0.1 - (1e10 - 1.5)))}
                                        v45 := GenS3{F0: (v2 * (v2 * fi29)), F1: -32768, F2: 4294967295}
                                        _ = v44
                                        _ = v45
                                    } else {
                                        for fi46 := 0; fi46 < 8; fi46++ {
                                            v5 -= (v3 ^ v5)
                                            var v47 uint16 = v3
                                            if (!("a" == "")) {
                                                v47 = 42
                                            } else {
                                                _ = len(v1)
                                                v3 = v5
                                                if false {
                                                    if ((fi29 >= fi46) || (false || (0 <= -866))) {
                                                        switch (fi29 * fi46) {
                                                        case 240:
                                                            v48 := GenS3{F0: v2, F1: -437, F2: 198}
                                                            _ = v48
                                                        }
                                                        fmt.Println(v2, v2, v2)
                                                    }
                                                }
                                                v49 := []int{(v2 - -9223372036854775808), v2, (^(fi29))}
                                                _ = v49
                                            }
                                            v5 = 147
                                            _ = fi46
                                            _ = v47
                                        }
                                        var v50 uint16 = (^(v3))
                                        _ = v50
                                    }
                                }
                            case 999:
                                if true {
                                    switch (v2 * (v2 % (((^(fi29))) | 1))) {
                                    case -60:
                                        var v51 int32 = 2147483647
                                        _ = v51
                                    default:
                                        switch (fi29 ^ v2) {
                                        case 128:
                                            if (2147483647 > 460) {
                                                for fi52 := 0; fi52 < 16; fi52++ {
                                                    v53 := GenS4{F0: -2147483648, F1: true, F2: 2}
                                                    v54 := GenS1{F0: (v5 * (v3 & 919)), F1: 690, F2: true}
                                                    _ = fi52
                                                    _ = v53
                                                    _ = v54
                                                }
                                            }
                                        case 479:
                                            var v55 int64 = 917
                                            _ = v55
                                        case 0:
                                            v3 -= 255
                                        }
                                    }
                                    if false {
                                        _ = len(v1)
                                        var v56 int8 = 1
                                        if (-1 != 32767) {
                                            if (!false) {
                                                v3 += 495
                                                switch "hello" {
                                                case "a":
                                                    v56 -= v56
                                                default:
                                                    v3 += v3
                                                }
                                                switch "" {
                                                case "a":
                                                    v57 := GenS5{F0: (float32((0.1 / 0.1)) == ((float32(0.1) / 1.0) * (1e10 - 0.0))), F1: -80, F2: (fi29 / ((222) | 1))}
                                                    _ = v57
                                                }
                                            } else {
                                                v58 := GenS2{F0: -1, F1: (v5 + 42), F2: 32767}
                                                _ = v58
                                            }
                                            v59 := GenS5{F0: false, F1: -1, F2: (^(v2))}
                                            _ = v59
                                        }
                                        _ = v56
                                    }
                                } else {
                                    var v60 bool = ((-1 >= -9223372036854775808) && false)
                                    _ = v60
                                }
                            default:
                                var v61 int16 = 0
                                _ = v61
                            }
                        }
                        _ = v30
                    }
                    _ = fi29
                }
                var v62 int16 = 84
                _ = v62
            }
        } else {
            v63 := GenS1{F0: v3, F1: 156, F2: false}
            fmt.Println(v2, v5, v2)
            delete(v4, 2)
            _ = v63
        }
    } else {
        fmt.Println(v3, v3, v2)
        v64 := GenS3{F0: v2, F1: 943, F2: 42}
        _ = v64
    }
    switch (^(v2)) {
    case 4294967295:
        v65 := GenS6{F0: -469, F1: ((!true) && (v2 <= v2)), F2: 0}
        _ = v65
    }
    var v66 float32 = (0.0 * float32(3.14159))
    if ("foo bar" <= ("\n\t" + ("" + ("a" + "foo bar")))) {
        v67 := GenS7{F0: 554, F1: 243}
        v68 := GenS3{F0: v2, F1: -883, F2: 4294967295}
        _ = v67
        _ = v68
    }
    fmt.Println(v66, v5, v66)
    v69 := []int{v2, 127}
    switch (v2 % ((0) | 1)) {
    case -67:
        if (!true) {
            if true {
                var v70 bool = ((!true) && false)
                _ = v70
            } else {
                fmt.Println(v2)
            }
            if false {
                v71 := GenS2{F0: 298, F1: v3, F2: (v5 - 366)}
                for fi72 := 0; fi72 < 9; fi72++ {
                    for fi73 := 0; fi73 < 18; fi73++ {
                        var v74 float64 = float64(0.0)
                        _ = fi73
                        _ = v74
                    }
                    if true {
                        fmt.Println(fi72, v66, v2)
                        if (141 <= uint8(v3)) {
                            var v75 uint16 = v5
                            _ = v75
                        }
                    } else {
                        if false {
                            if ((v3 / (((^(v5))) | 1)) <= (v5 - 454)) {
                                for fi76 := 0; fi76 < 8; fi76++ {
                                    v77 := v4[899]
                                    _ = fi76
                                    _ = v77
                                }
                            } else {
                                var v78 uint32 = 741
                                _ = len(v69)
                                v79 := GenS5{F0: true, F1: 625, F2: (v2 & (fi72 << ((18446744073709551615) & 63)))}
                                _ = v78
                                _ = v79
                            }
                        }
                    }
                    var v80 int = fi72
                    _ = fi72
                    _ = v80
                }
                _ = v71
            } else {
                var v81 string = ("a" + "a")
                var v82 int8 = 53
                _ = v81
                _ = v82
            }
            for fi83 := 0; fi83 < 18; fi83++ {
                v84 := GenS2{F0: -1, F1: 255, F2: v3}
                if false {
                    v3 -= v3
                    v3 += (v3 / ((977) | 1))
                    switch 32768 {
                    case -246:
                        var v85 string = ("" + "")
                        _ = v85
                    case 32768:
                        if true {
                            v69 = append(v69, v2)
                        }
                    case 127:
                        if (65535 <= v3) {
                            if (true && ((!true) && (127 != -94))) {
                                if true {
                                    var v86 int32 = 42
                                    fmt.Println(v86)
                                    fmt.Println(v3, v3)
                                    _ = v86
                                } else {
                                    v87 := GenS8{F0: -1, F1: v5}
                                    var v88 float64 = (1e10 / 1.5)
                                    v88 = v88
                                    _ = v87
                                    _ = v88
                                }
                            } else {
                                for fi89 := 0; fi89 < 12; fi89++ {
                                    if (true || true) {
                                        if (!true) {
                                            v5 += (v3 + v3)
                                            var v90 uint16 = (v5 - v5)
                                            _ = v90
                                        }
                                        v5 = (v3 ^ v3)
                                        v66 -= ((v66 - (1e10 - (0.1 / (v66 + ((float32((0.0 + 0.1)) - v66) * 0.1))))) * (0.0 * float32((0.1 + -2.25))))
                                        switch ((("\n\t" + "\n\t") + (("foo bar" + (("hello" + "a") + ("" + ("foo bar" + "foo bar")))) + "foo bar")) + ("" + "foo bar")) {
                                        case "hello":
                                            switch (fi83 << ((2) & 63)) {
                                            case -678:
                                                if len(v69) >= 2 {
                                                    v69 = v69[1:len(v69)]
                                                }
                                            case 127:
                                                var v91 uint8 = 127
                                                _ = v91
                                            case 42:
                                                v3 = (^(v3))
                                            }
                                        case "":
                                            v92 := GenS3{F0: (fi83 % (((-(v2))) | 1)), F1: 42, F2: 42}
                                            _ = v92
                                        default:
                                            fmt.Println(v5, fi89)
                                        }
                                        v5 = (v3 * (v5 + 255))
                                    } else {
                                        var v93 string = ("\n\t" + "a")
                                        var v94 float32 = v66
                                        v66 += v66
                                        _ = v93
                                        _ = v94
                                    }
                                    var v95 int8 = -26
                                    _ = fi89
                                    _ = v95
                                }
                                v96 := GenS7{F0: 255, F1: 42}
                                _ = v96
                            }
                            v66 -= v66
                            if true {
                                var v97 float32 = 3.14159
                                v66 = v97
                                v98 := GenS9{F0: 0, F1: ("foo bar" + (("foo bar" + "hello") + "")), F2: 2147483647}
                                _ = v97
                                _ = v98
                            }
                        }
                    }
                }
                switch "\n\t" {
                case "\n\t":
                    v1[""] = -9223372036854775808
                case "foo bar":
                    var v99 uint = 255
                    _ = v99
                default:
                    v100 := GenS10{F0: -730, F1: -461, F2: 0}
                    _ = v100
                }
                _ = fi83
                _ = v84
            }
        }
    case 4294967295:
        for fi101 := 0; fi101 < 15; fi101++ {
            v102 := GenS11{F0: 561, F1: 42}
            v2 += v2
            _ = fi101
            _ = v102
        }
    }
    v69 = append(v69, (v2 | v2))
    if len(v69) >= 2 {
        v69 = v69[1:len(v69)]
    }
    var v103 int32 = 0
    var v104 int32 = v103
    if true {
        fmt.Println(v104)
        for fi105 := 0; fi105 < 6; fi105++ {
            var v106 int = 32767
            if len(v69) > 0 {
                v69[0] = v2
            }
            _ = fi105
            _ = v106
        }
        if true {
            for fi107 := 0; fi107 < 13; fi107++ {
                fmt.Println(fi107)
                _ = fi107
            }
            if false {
                if false {
                    if true {
                        v104 += 642
                        var v108 float32 = v66
                        _ = v108
                    }
                }
                v109 := GenS6{F0: 0, F1: true, F2: 42}
                _ = v109
            } else {
                v69 = append(v69, v2)
                _ = len(v69)
                v66 -= v66
            }
            for fi110 := 0; fi110 < 14; fi110++ {
                v2 -= v2
                var v111 int32 = v104
                var v112 uint = 606
                _ = fi110
                _ = v111
                _ = v112
            }
        } else {
            var v113 float32 = v66
            v66 -= float32(0.0)
            _ = v113
        }
    } else {
        fmt.Println(v2, v3)
        var v114 uint64 = 4294967295
        switch v2 {
        case 67:
            var v115 float32 = (v66 + (float32((0.0 + ((3.14159 * (0.0 * float64((((1e10 + float32(0.0)) - v66) + v66)))) / -2.25))) * 3.14159))
            _ = v115
        }
        _ = v114
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v69:", len(v69), v69)
    _ = v1
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v69
}
