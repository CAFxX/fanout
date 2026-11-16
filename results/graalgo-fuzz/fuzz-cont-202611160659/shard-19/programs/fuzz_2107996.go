package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint64; F1 bool; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint16; F1 int16; F2 int16 }
type GenS10 struct { F0 uint8; F1 string; F2 int8 }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 int16; F1 uint }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int16; F1 int32; F2 uint }

func main() {
    var v1 uint16 = 205
    var v2 int32 = 932
    var v3 int32 = (v2 ^ 42)
    var v4 float64 = -2.25
    if (!(!(!(32767 != 542)))) {
        switch "hello" {
        case "":
            if true {
                if false {
                    if (2147483647 > 905) {
                        if (false && ((!(true || true)) && true)) {
                            v5 := []int{0, -678, -2147483648, 0}
                            v4 += v4
                            _ = v5
                        }
                        v6 := GenS1{F0: 696, F1: 127}
                        _ = v6
                    }
                    if (true || false) {
                        v7 := GenS1{F0: 1, F1: 254}
                        for fi8 := 0; fi8 < 8; fi8++ {
                            v2 += -2147483648
                            _ = fi8
                        }
                        _ = v7
                    }
                    switch ("a" + "a") {
                    case "hello":
                        v9 := []int{42, 2147483647}
                        _ = v9
                    case "":
                        if (!true) {
                            switch ("hello" + ("hello" + "foo bar")) {
                            case "":
                                for fi10 := 0; fi10 < 17; fi10++ {
                                    v2 += 1
                                    if (1.5 <= (float32(float64(3.14159)) * (((3.14159 + 3.14159) + 1.5) / 1.5))) {
                                        v3 = (-(v3))
                                        for fi11 := 0; fi11 < 16; fi11++ {
                                            v12 := []int{fi11}
                                            v4 -= v4
                                            _ = fi11
                                            _ = v12
                                        }
                                        v2 += (v3 % (((v2 * v3)) | 1))
                                    }
                                    switch 42 {
                                    case 128:
                                        v2 = -2147483648
                                    }
                                    _ = fi10
                                }
                            case "a":
                                var v13 int32 = (v2 & (v2 & (v2 & -2147483648)))
                                _ = v13
                            default:
                                var v14 uint16 = (v1 % (((v1 / (((v1 & v1)) | 1))) | 1))
                                _ = v14
                            }
                            v1 -= 760
                            v15 := []int{0, -776, 1, 9223372036854775807}
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 5; fi16++ {
                                fmt.Println(fi16, v4)
                                if false {
                                    var v17 int32 = (v2 >> ((uint(v1)) & 63))
                                    _ = v17
                                }
                                _ = fi16
                            }
                        }
                    default:
                        var v18 uint64 = 2
                        _ = v18
                    }
                }
                v19 := []int{}
                _ = v19
            }
        case "\n\t":
            fmt.Println(v1)
        }
    }
    v20 := []int{int(v1), -2147483648, int(v1), 211}
    fmt.Println(v1, v4)
    var v21 float32 = 0.1
    var v22 uint = 2
    var v23 int8 = -23
    var v24 uint = 0
    switch -1 {
    case -504:
        for fi25 := 0; fi25 < 17; fi25++ {
            if false {
                if ("a" <= ((("" + "") + ("a" + (("foo bar" + "\n\t") + "a"))) + "foo bar")) {
                    if (!("foo bar" >= "foo bar")) {
                        v26 := GenS2{F0: (fi25 | 9223372036854775807), F1: int8(v22), F2: 65535}
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v28 := GenS3{F0: 42, F1: -25}
                            var v29 uint8 = 44
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        if true {
                            switch "foo bar" {
                            case "\n\t":
                                for fi30 := 0; fi30 < 10; fi30++ {
                                    v21 -= 0.1
                                    v3 = v2
                                    _ = fi30
                                }
                            case "a":
                                var v31 bool = false
                                _ = v31
                            default:
                                v20 = append(v20, (fi25 ^ fi25))
                            }
                            if false {
                                _ = len(v20)
                                _ = len(v20)
                                var v32 uint32 = 4294967295
                                _ = v32
                            } else {
                                var v33 int = (fi25 / (((fi25 * (fi25 / (((-(fi25))) | 1)))) | 1))
                                _ = v33
                            }
                            var v34 uint32 = 0
                            _ = v34
                        } else {
                            fmt.Println(v1)
                            if ((fi25 - (fi25 + fi25)) <= (fi25 >> (((v22 ^ v24)) & 63))) {
                                if (false && true) {
                                    fmt.Println(v21, fi25, v22)
                                    v24 += (v22 | (^(v22)))
                                    var v35 int64 = -354
                                    _ = v35
                                }
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            } else {
                                switch (-(fi25)) {
                                case 252:
                                    fmt.Println(v3, v4)
                                case -1:
                                    for fi36 := 0; fi36 < 4; fi36++ {
                                        var v37 string = ((("foo bar" + "hello") + "a") + ("" + "a"))
                                        _ = fi36
                                        _ = v37
                                    }
                                default:
                                    for fi38 := 0; fi38 < 8; fi38++ {
                                        v24 -= 85
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                        switch 32767 {
                                        case -2147483648:
                                            var v39 bool = false
                                            if (!(6 < 128)) {
                                                v20 = append(v20, 127)
                                                var v40 uint8 = 37
                                                _ = v40
                                            } else {
                                                v41 := GenS4{F0: 4294967295, F1: v39, F2: fi25}
                                                v21 = (((3.14159 - -2.25) / 1e10) * v21)
                                                if (1 < 541) {
                                                    if (65535 < (^(v1))) {
                                                        v24 = (^(v22))
                                                        v39 = (16 <= 42)
                                                        fmt.Println(v23, v22)
                                                    } else {
                                                        v20 = append(v20, (^(fi25)))
                                                        if (3.14159 == v4) {
                                                            v3 = (v3 - 2147483647)
                                                        } else {
                                                            var v42 uint16 = 32767
                                                            v39 = (v39 && (false || (9223372036854775807 <= 469)))
                                                            v4 = (v4 - (float64(1e10) + 0.0))
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    if (!v39) {
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                    } else {
                                                                        var v43 uint = v22
                                                                        _ = v43
                                                                    }
                                                                case "":
                                                                    if v39 {
                                                                        var v44 uint64 = 182
                                                                        if (9 >= 1) {
                                                                            var v45 int64 = -205
                                                                            switch ("foo bar" + "foo bar") {
                                                                            case "a":
                                                                                v1 = v1
                                                                            case "hello":
                                                                                if v39 {
                                                                                    v20 = append(v20, fi38)
                                                                                    v46 := GenS1{F0: 1, F1: uint8(v24)}
                                                                                    _ = v46
                                                                                }
                                                                            case "foo bar":
                                                                                if (float32(float64((v21 - 1e10))) > 0.1) {
                                                                                    fmt.Println(v3, v22)
                                                                                    fmt.Println(v21)
                                                                                    if (!(((v44 >> ((v24) & 63)) <= (v44 >> (((v24 - (v24 | (v22 % ((581) | 1))))) & 63))) || (92 == v23))) {
                                                                                        v47 := GenS1{F0: (v24 + (v24 % (((^(v24))) | 1))), F1: 127}
                                                                                        var v48 uint16 = v42
                                                                                        if len(v20) >= 2 {
                                                                                            v20 = v20[1:len(v20)]
                                                                                        }
                                                                                        if true {
                                                                                            var v49 uint16 = 722
                                                                                            if v39 {
                                                                                                if (false && (-2.25 > float64(((v21 / v21) + 0.1)))) {
                                                                                                    fmt.Println(v4, v24, v45)
                                                                                                    if true {
                                                                                                        _ = len(v20)
                                                                                                        var v50 uint8 = 255
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    fmt.Println(v4, v24)
                                                                                                }
                                                                                                switch fi38 {
                                                                                                case -2147483648:
                                                                                                    v23 = (-(v23))
                                                                                                }
                                                                                            } else {
                                                                                                if (-32768 > 752) {
                                                                                                    var v51 float64 = (v4 + 0.0)
                                                                                                    var v52 int = -1
                                                                                                    v23 += -82
                                                                                                    _ = v51
                                                                                                    _ = v52
                                                                                                }
                                                                                            }
                                                                                            _ = v49
                                                                                        }
                                                                                        _ = v47
                                                                                        _ = v48
                                                                                    }
                                                                                } else {
                                                                                    switch (fi25 % ((int(v44)) | 1)) {
                                                                                    case 704:
                                                                                        var v53 uint16 = v42
                                                                                        _ = v53
                                                                                    case 2147483647:
                                                                                        v22 = v24
                                                                                    case 1:
                                                                                        fmt.Println(v24, v1)
                                                                                    default:
                                                                                        if false {
                                                                                            var v54 int = (fi25 >> (((v24 | 66)) & 63))
                                                                                            v1 = (v42 | (v1 | v42))
                                                                                            _ = v54
                                                                                        } else {
                                                                                            v3 += (-(v2))
                                                                                            var v55 uint = v22
                                                                                            var v56 int32 = (^(v2))
                                                                                            _ = v55
                                                                                            _ = v56
                                                                                        }
                                                                                    }
                                                                                    v57 := GenS5{F0: v39}
                                                                                    _ = v57
                                                                                }
                                                                            }
                                                                            v23 += -112
                                                                            _ = v45
                                                                        }
                                                                        fmt.Println(v1)
                                                                        _ = v44
                                                                    } else {
                                                                        if (!v39) {
                                                                            switch 4294967295 {
                                                                            case 128:
                                                                                v58 := GenS6{F0: -1, F1: v42, F2: 128}
                                                                                _ = v58
                                                                            case -50:
                                                                                v21 = 0.1
                                                                                v59 := GenS7{F0: v39}
                                                                                v3 -= (v3 & 0)
                                                                                _ = v59
                                                                            case -418:
                                                                                v22 = v22
                                                                                v60 := GenS1{F0: v22, F1: 10}
                                                                                v3 -= v2
                                                                                _ = v60
                                                                            default:
                                                                                v61 := GenS6{F0: int16(v3), F1: v1, F2: 255}
                                                                                _ = v61
                                                                            }
                                                                            var v62 int64 = 0
                                                                            if (false && (!(677 == (v2 / ((v2) | 1))))) {
                                                                                v63 := GenS8{F0: 0, F1: v62}
                                                                                _ = v63
                                                                            } else {
                                                                                v20 = append(v20, fi38)
                                                                                v20 = append(v20, int(v22))
                                                                            }
                                                                            if v39 {
                                                                                v64 := GenS9{F0: v1, F1: 0, F2: int16(fi25)}
                                                                                _ = v64
                                                                            }
                                                                            _ = v62
                                                                        }
                                                                        fmt.Println(v39, fi25, v23)
                                                                    }
                                                                }
                                                            }
                                                            v21 = v21
                                                            _ = v42
                                                        }
                                                        v65 := GenS8{F0: 1, F1: -1}
                                                        _ = v65
                                                    }
                                                }
                                                _ = v41
                                            }
                                            _ = v39
                                        case -1:
                                            v66 := GenS10{F0: 127, F1: "", F2: (v23 + 42)}
                                            _ = v66
                                        case 32767:
                                            if false {
                                                var v67 string = ""
                                                v68 := GenS11{F0: v24, F1: 1}
                                                v4 = 1.5
                                                var v69 uint16 = v1
                                                _ = v67
                                                _ = v68
                                                _ = v69
                                            }
                                        }
                                        if true {
                                            fmt.Println(v24)
                                            if ((!true) || (!((v23 & v23) > v23))) {
                                                if true {
                                                    if false {
                                                        var v70 int64 = 932
                                                        _ = v70
                                                    } else {
                                                        var v71 int8 = (v23 >> (((^(v22))) & 63))
                                                        _ = v71
                                                    }
                                                } else {
                                                    var v72 int = -732
                                                    v20 = append(v20, v72)
                                                    var v73 uint64 = 18446744073709551615
                                                    var v74 uint8 = uint8(v1)
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if false {
                                                            var v75 int = fi38
                                                            v75 += (fi38 ^ (v72 & fi38))
                                                            var v76 int8 = (^(v23))
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            switch fi25 {
                                                            case 32768:
                                                                v77 := GenS11{F0: (v22 / ((uint(v73)) | 1)), F1: (v22 / ((v22) | 1))}
                                                                _ = v77
                                                            case 4294967295:
                                                                _ = len(v20)
                                                            case 0:
                                                                if len(v20) > 0 {
                                                                    v20[0] = (v72 << ((v24) & 63))
                                                                }
                                                            default:
                                                                v73 += 42
                                                            }
                                                            var v78 float32 = v21
                                                            var v79 uint16 = (v1 - (^(v1)))
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                    default:
                                                        v80 := GenS1{F0: (v24 & v24), F1: 2}
                                                        _ = v80
                                                    }
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                }
                                                v81 := GenS4{F0: 4294967295, F1: true, F2: (fi38 - fi25)}
                                                v82 := GenS12{F0: 1000, F1: 65535}
                                                _ = v81
                                                _ = v82
                                            }
                                        }
                                        _ = fi38
                                    }
                                }
                                v22 += (v24 << (((v22 * (^(v24)))) & 63))
                            }
                            v21 = (float32(((((3.14159 * float64(v21)) + (v4 - v4)) * ((1e10 + (v4 + v4)) * v4)) * (0.1 * 1.5))) * 1e10)
                        }
                        _ = v26
                    } else {
                        switch (fi25 + fi25) {
                        case -2147483648:
                            v3 += -2147483648
                        case 127:
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        case 0:
                            if (true && (((9223372036854775807 == 0) && ((-9223372036854775808 != 862) || true)) && (false && false))) {
                                switch "foo bar" {
                                case "hello":
                                    var v83 float32 = v21
                                    _ = v83
                                default:
                                    v2 = int32(fi25)
                                }
                                v22 += v22
                            }
                        default:
                            for fi84 := 0; fi84 < 14; fi84++ {
                                fmt.Println(fi84)
                                fmt.Println(v2, v21)
                                if true {
                                    v85 := GenS1{F0: (v22 & (v24 >> ((v22) & 63))), F1: 1}
                                    switch 32768 {
                                    case 320:
                                        var v86 uint8 = 127
                                        _ = v86
                                    case 32768:
                                        fmt.Println(fi25)
                                    }
                                    switch fi84 {
                                    case -970:
                                        v87 := GenS13{F0: 42}
                                        _ = v87
                                    case 919:
                                        _ = len(v20)
                                        var v88 float64 = (v4 / float64(float32(v4)))
                                        _ = v88
                                    case 680:
                                        v3 = (v3 * 724)
                                    }
                                    _ = v85
                                } else {
                                    v2 = (v2 * 2147483647)
                                    v3 -= int32(v1)
                                    v24 = 255
                                    if (true || (!(-9223372036854775808 <= -9223372036854775808))) {
                                        var v89 uint64 = 4294967295
                                        switch (fi25 >> (((v22 & v22)) & 63)) {
                                        case 127:
                                            var v90 int64 = 9223372036854775807
                                            _ = v90
                                        case 9223372036854775807:
                                            fmt.Println(v89)
                                        default:
                                            fmt.Println(v4, v23, fi25)
                                        }
                                        v91 := GenS1{F0: 431, F1: 2}
                                        if (false && false) {
                                            v92 := GenS11{F0: (v24 + (v24 + v24)), F1: (v24 + 389)}
                                            fmt.Println(fi25, v1, v21)
                                            _ = v92
                                        }
                                        _ = v89
                                        _ = v91
                                    } else {
                                        v20 = append(v20, (fi25 >> (((v24 * (^(v22)))) & 63)))
                                    }
                                }
                                v4 = (1e10 - v4)
                                var v93 float64 = (float64(v21) * (0.0 + (1e10 / float64((v21 * (v21 / (0.0 + (3.14159 + ((v21 + v21) / (v21 - 1.5))))))))))
                                _ = fi84
                                _ = v93
                            }
                        }
                        v1 -= 937
                    }
                }
            }
            fmt.Println(fi25)
            _ = fi25
        }
    default:
        v94 := GenS8{F0: 0, F1: -9223372036854775808}
        _ = v94
    }
    var v95 int8 = 1
    var v96 float64 = ((((v4 - -2.25) + float64((float32(-2.25) * 3.14159))) * v4) * (float64(1.5) - v4))
    switch "hello" {
    case "a":
        switch 0 {
        case -625:
            v2 = 869
        default:
            fmt.Println(v21)
        }
    case "hello":
        if false {
            v97 := GenS14{F0: -32768, F1: (v2 - -373), F2: v24}
            _ = v97
        }
    default:
        if ((^(v22)) >= 1) {
            var v98 uint8 = 255
            switch int(v98) {
            case 9223372036854775807:
                v23 -= v23
            }
            _ = v98
        }
    }
    switch -757 {
    case -732:
        var v99 uint8 = 0
        _ = v99
    case -319:
        v24 += (^(v22))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v95
    _ = v96
}
