package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint32; F2 int }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 uint64, p1 bool) uint {
    var v5 int32 = 2147483647
    v6 := []int{-50, 42, 1}
    if ((2 == 255) && (!p1)) {
        p0 = p0
        var v7 uint32 = 42
        var v8 int32 = v5
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            p0 = (p0 << ((42) & 63))
            for fi10 := 0; fi10 < 2; fi10++ {
                p0 += p0
                var v11 int = 128
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        p1 = false
    }
    if (!p1) {
        var v12 int32 = (v5 / ((v5) | 1))
        _ = v12
    } else {
        var v13 float64 = 1e10
        v14 := GenS1{F0: 32767}
        _ = v13
        _ = v14
    }
    var v15 float32 = (0.1 / 1.0)
    _ = p0
    _ = p1
    _ = v15
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int) string {
    switch ("hello" + "a") {
    case "\n\t":
        fmt.Println(p0)
    case "foo bar":
        p0 = 392
    case "":
        var v16 bool = ((2 > 42) || (true || ((((((3.14159 / ((-2.25 - 3.14159) / (float32(3.14159) * float32((1e10 * ((float64((0.0 + (((0.0 / 1.0) * -2.25) / (1e10 - -2.25)))) / (-2.25 + -2.25)) + (-2.25 / 3.14159))))))) - float32((3.14159 + (float64(float32(0.1)) * (0.1 * ((((-2.25 * float64(1e10)) / (-2.25 * 0.1)) * 1.5) + 0.0)))))) - float32((0.0 * 1.5))) - 1e10) + 3.14159) >= (0.0 * 1.5))))
        _ = v16
    }
    switch "" {
    case "hello":
        p0 = 2147483647
    case "":
        p0 += (p0 & p0)
    default:
        switch 4294967295 {
        case -470:
            p0 = p0
        case 32767:
            p0 -= -340
        case 1:
            var v17 bool = (false || ((!((p0 ^ 32767) > -9223372036854775808)) && false))
            _ = v17
        default:
            if true {
                fmt.Println(p0)
            } else {
                v18 := GenS1{F0: 32767}
                p0 -= (p0 & 128)
                var v19 float32 = 3.14159
                _ = v18
                _ = v19
            }
        }
    }
    _ = p0
    return "a"
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 string = "hello"
    var v4 uint32 = 2147483648
    fmt.Println(v4, v4)
    v4 = (^(v4))
    var v20 uint64 = 66
    v21 := GenS1{F0: 0}
    switch -598 {
    case 32768:
        fmt.Println(v20, v4)
    case 0:
        v20 += (v20 * v20)
    case 127:
        v22 := GenS2{F0: 42}
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        if true {
            v4 += v4
            for fi24 := 0; fi24 < 2; fi24++ {
                fmt.Println(v20)
                v4 = (v4 & v4)
                _ = fi24
            }
            v20 -= v20
        }
        var v25 int8 = -128
        switch (fi23 / ((-2147483648) | 1)) {
        case 32767:
            v26 := GenS3{F0: "a"}
            _ = v26
        case -9223372036854775808:
            if false {
                if (false && ((((float32(1e10) * 0.1) + 0.1) < 0.1) && (!true))) {
                    if true {
                        v27 := []int{(^(fi23)), 295, 32767}
                        if (v4 == (v4 | (v4 * (v4 << ((65535) & 63))))) {
                            v28 := GenS4{F0: -2147483648, F1: (v4 << ((18446744073709551615) & 63)), F2: 303}
                            _ = v28
                        }
                        var v29 float32 = 3.14159
                        _ = v27
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 16; fi30++ {
                        var v31 int16 = -32768
                        _ = fi30
                        _ = v31
                    }
                    v3 = ("a" + "a")
                }
            } else {
                v25 -= (-(v25))
            }
        }
        _ = fi23
        _ = v25
    }
    switch v3 {
    case "foo bar":
        v32 := v2[32767]
        _ = v32
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        for fi34 := 0; fi34 < 2; fi34++ {
            if (255 > 0) {
                switch fi33 {
                case 40:
                    switch fi34 {
                    case -1:
                        var v35 bool = false
                        _ = v35
                    case 2147483647:
                        var v36 uint = 0
                        v20 = (v20 >> (((v36 << ((v36) & 63))) & 63))
                        _ = v36
                    case 42:
                        switch v3 {
                        case "foo bar":
                            var v37 int = (fi33 % ((fi33) | 1))
                            _ = v37
                        case "":
                            var v38 uint16 = 553
                            _ = v38
                        case "a":
                            fmt.Println(fi34, v4)
                        }
                    }
                case -826:
                    if (-641 <= -2147483648) {
                        v20 -= v20
                    } else {
                        v4 -= 882
                    }
                case 9223372036854775807:
                    switch fi34 {
                    case 128:
                        var v39 uint64 = 0
                        _ = v39
                    case -123:
                        v40 := []int{fi33}
                        _ = v40
                    case 127:
                        var v41 uint8 = 7
                        _ = v41
                    default:
                        var v42 float64 = ((1e10 / 0.1) / 1.0)
                        _ = v42
                    }
                    v43 := v2[fi33]
                    var v44 float32 = (0.1 - 1e10)
                    _ = v43
                    _ = v44
                }
            } else {
                v45 := GenS3{F0: v3}
                v4 = 42
                _ = v45
            }
            v46 := GenS5{F0: 2147483647, F1: (v20 ^ (^(v20)))}
            var v47 float32 = 1e10
            switch (fi34 << ((2) & 63)) {
            case 4294967295:
                var v48 string = "\n\t"
                _ = v48
            case 128:
                v49 := GenS6{F0: 100, F1: v20}
                _ = v49
            default:
                v50 := GenS7{F0: v4, F1: 2, F2: int8(v4)}
                _ = v50
            }
            fmt.Println(v47, fi34, v3)
            v3 = "a"
            v4 = v4
            v51 := []int{(-(fi33)), (fi33 + (fi34 | (^(fi34)))), fi33, (fi33 | fi34)}
            v51 = append(v51, 4294967295)
            _ = fi34
            _ = v46
            _ = v47
            _ = v51
        }
        v52 := []int{128, 154}
        _ = fi33
        _ = v52
    }
    for fi53 := 0; fi53 < 15; fi53++ {
        for fi54 := 0; fi54 < 4; fi54++ {
            v20 = 18446744073709551615
            v4 = v4
            _ = fi54
        }
        _ = fi53
    }
    v20 = (v20 >> ((0) & 63))
    for fi55 := 0; fi55 < 14; fi55++ {
        switch fi55 {
        case -2147483648:
            v4 = v4
        case 128:
            v3 = v3
        case 963:
            switch (fi55 >> ((660) & 63)) {
            case -1000:
                switch (fi55 % (((fi55 / ((fi55) | 1))) | 1)) {
                case 42:
                    switch "hello" {
                    case "":
                        var v56 uint8 = 1
                        _ = v56
                    case "a":
                        var v57 uint16 = 32767
                        _ = v57
                    case "hello":
                        var v58 float32 = float32((float64(0.1) * ((((((float64((((3.14159 + (0.0 / 1e10)) - 3.14159) / float32(-2.25))) + (1.5 / 1e10)) + (-2.25 / 1e10)) / (float64((1e10 / ((((float32(-2.25) / (float32(-2.25) / 1e10)) - 0.0) * (((0.1 + ((float32(1.5) / -2.25) - ((1e10 * (0.1 - (1e10 - 3.14159))) + 0.1))) + 3.14159) / 1e10)) + 0.1))) - 3.14159)) - ((float64((float32(0.1) + 3.14159)) - (3.14159 + -2.25)) / -2.25)) + float64(1e10)) * (1.5 - -2.25))))
                        _ = v58
                    }
                case 32768:
                    v3 = ("foo bar" + v3)
                }
            case -249:
                var v59 uint32 = (v4 * (v4 | (v4 / ((v4) | 1))))
                _ = v59
            }
        default:
            v3 = v3
        }
        delete(v1, (-(fi55)))
        _ = fi55
    }
    var v60 string = v3
    switch -47 {
    case -1:
        if false {
            for fi61 := 0; fi61 < 3; fi61++ {
                v3 = v3
                if (false || false) {
                    if (0.0 != float32(0.0)) {
                        switch (fi61 + (fi61 ^ (fi61 | (fi61 / ((fi61) | 1))))) {
                        case 496:
                            v4 += (v4 - v4)
                        case 32768:
                            var v62 uint32 = 663
                            _ = v62
                        case 138:
                            v20 -= 329
                        default:
                            for fi63 := 0; fi63 < 12; fi63++ {
                                delete(v1, (fi61 % ((-1) | 1)))
                                if false {
                                    v64 := GenS3{F0: (v60 + v60)}
                                    _ = v64
                                }
                                v4 = (v4 >> ((genFn0((v20 % ((0) | 1)), false)) & 63))
                                _ = fi63
                            }
                        }
                        var v65 float32 = 3.14159
                        fmt.Println(v3)
                        _ = v65
                    } else {
                        fmt.Println(v4, v20)
                    }
                }
                v4 = 42
                _ = fi61
            }
            v60 = (v3 + v3)
        } else {
            switch v60 {
            case "":
                if (!(false && (-1 <= 42))) {
                    if true {
                        v66 := GenS6{F0: -1, F1: (v20 * 18446744073709551615)}
                        v4 = 399
                        var v67 uint16 = 255
                        _ = v66
                        _ = v67
                    } else {
                        switch 1 {
                        case -416:
                            v68 := GenS8{F0: -1}
                            _ = v68
                        }
                        v60 = ((v60 + v60) + ("hello" + "\n\t"))
                    }
                } else {
                    switch -300 {
                    case 9223372036854775807:
                        for fi69 := 0; fi69 < 4; fi69++ {
                            var v70 float32 = 0.0
                            v60 = v3
                            var v71 uint8 = 2
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                    case -1:
                        v60 = v3
                    }
                }
            case "foo bar":
                fmt.Println(v3)
            default:
                v4 = v4
            }
        }
    case 1:
        var v72 float32 = 0.1
        _ = v72
    case 42:
        v60 = ("hello" + v60)
    default:
        var v73 int8 = int8(v4)
        _ = v73
    }
    _ = len(v2)
    v74 := GenS2{F0: 42}
    var v75 uint = 119
    var v76 uint64 = (v20 / ((341) | 1))
    if false {
        for fi77 := 0; fi77 < 12; fi77++ {
            var v78 int8 = 0
            var v79 int = (fi77 << (((v75 << (((v75 << (((v75 * (^(v75)))) & 63))) & 63))) & 63))
            v75 = genFn0(4294967295, (uint8(v4) >= 1))
            _ = fi77
            _ = v78
            _ = v79
        }
        if ((3.14159 - (0.0 * 3.14159)) != (-2.25 / 1e10)) {
            if (!true) {
                v3 = v60
                if (!(-128 < 3)) {
                    switch 0 {
                    case 910:
                        fmt.Println(v3)
                    case -77:
                        switch v60 {
                        case "hello":
                            var v80 int64 = -1
                            _ = v80
                        case "foo bar":
                            if true {
                                for fi81 := 0; fi81 < 16; fi81++ {
                                    if false {
                                        switch -2147483648 {
                                        case 9223372036854775807:
                                            switch (v60 + v60) {
                                            case "\n\t":
                                                var v82 int = -2147483648
                                                _ = v82
                                            case "":
                                                v83 := GenS9{F0: fi81, F1: 896, F2: 1000}
                                                _ = v83
                                            case "hello":
                                                var v84 uint8 = 0
                                                _ = v84
                                            default:
                                                var v85 uint8 = 255
                                                _ = v85
                                            }
                                        case 127:
                                            fmt.Println(v60, v60, v75)
                                        default:
                                            v76 = v20
                                        }
                                        for fi86 := 0; fi86 < 8; fi86++ {
                                            if true {
                                                switch (fi81 - fi81) {
                                                case 1:
                                                    v3 = v3
                                                case 127:
                                                    switch v60 {
                                                    case "hello":
                                                        v4 = 156
                                                    case "foo bar":
                                                        var v87 float32 = 0.0
                                                        _ = v87
                                                    case "a":
                                                        v88 := v2[fi86]
                                                        _ = v88
                                                    default:
                                                        var v89 string = "foo bar"
                                                        var v90 float32 = (1.5 * 0.0)
                                                        v20 += (v76 & 2)
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                default:
                                                    var v91 float64 = 0.1
                                                    _ = v91
                                                }
                                                v92 := []int{(^(fi86))}
                                                _ = v92
                                            }
                                            v2[fi81] = true
                                            _ = fi86
                                        }
                                        if true {
                                            var v93 bool = ((42 > -1) && false)
                                            v94 := []int{(fi81 >> ((genFn0(v20, (((fi81 <= fi81) || ((!(-2.25 == -2.25)) || (127 == 5))) || v93))) & 63)), int(v76), fi81, (fi81 * -916)}
                                            if false {
                                                fmt.Println(v76, v76)
                                                v20 = v20
                                            } else {
                                                for fi95 := 0; fi95 < 8; fi95++ {
                                                    v3 = ("foo bar" + "")
                                                    _ = fi95
                                                }
                                                if (111 <= 127) {
                                                    v96 := GenS2{F0: 42}
                                                    for fi97 := 0; fi97 < 5; fi97++ {
                                                        var v98 uint8 = 20
                                                        var v99 uint8 = v98
                                                        switch (^(fi81)) {
                                                        case 1:
                                                            fmt.Println(v76, v60)
                                                        }
                                                        v76 = (v76 | v20)
                                                        _ = fi97
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                    _ = v96
                                                } else {
                                                    var v100 int32 = 0
                                                    _ = v100
                                                }
                                            }
                                            _ = v93
                                            _ = v94
                                        }
                                    }
                                    v101 := GenS9{F0: (fi81 | fi81), F1: (v4 / (((v4 - v4)) | 1)), F2: 44}
                                    for fi102 := 0; fi102 < 2; fi102++ {
                                        var v103 uint8 = 42
                                        v20 -= 217
                                        _ = fi102
                                        _ = v103
                                    }
                                    _ = fi81
                                    _ = v101
                                }
                            }
                        case "":
                            if (false && (uint16(v75) == 42)) {
                                var v104 uint16 = 32767
                                _ = len(v2)
                                fmt.Println(v60)
                                _ = v104
                            } else {
                                for fi105 := 0; fi105 < 19; fi105++ {
                                    for fi106 := 0; fi106 < 19; fi106++ {
                                        if (-128 < 42) {
                                            v107 := GenS8{F0: 655}
                                            _ = v107
                                        } else {
                                            var v108 int8 = -128
                                            v109 := GenS7{F0: 2147483647, F1: 0, F2: 100}
                                            _ = v108
                                            _ = v109
                                        }
                                        v75 = v75
                                        var v110 uint = (v75 >> ((genFn0(v76, false)) & 63))
                                        var v111 bool = false
                                        v4 -= v4
                                        _ = fi106
                                        _ = v110
                                        _ = v111
                                    }
                                    _ = fi105
                                }
                                switch int(v75) {
                                case 2147483647:
                                    for fi112 := 0; fi112 < 3; fi112++ {
                                        v113 := []int{fi112}
                                        _ = fi112
                                        _ = v113
                                    }
                                case 1:
                                    v60 = "foo bar"
                                default:
                                    v60 = ""
                                }
                                var v114 int32 = -528
                                _ = v114
                            }
                        }
                    default:
                        for fi115 := 0; fi115 < 8; fi115++ {
                            v75 += v75
                            _ = fi115
                        }
                    }
                    fmt.Println(v20)
                }
            } else {
                var v116 int32 = -2147483648
                _ = v116
            }
            var v117 uint32 = v4
            if (!(-1 <= -9223372036854775808)) {
                var v118 int = -108
                for fi119 := 0; fi119 < 3; fi119++ {
                    v2[1] = false
                    _ = fi119
                }
                _ = v118
            }
            _ = v117
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v60
    _ = v74
    _ = v75
    _ = v76
}
