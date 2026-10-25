package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 int16; F1 uint32; F2 bool }
type GenS8 struct { F0 int16; F1 int16; F2 int64 }
type GenS9 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS10 struct { F0 bool; F1 int; F2 uint32 }
type GenS11 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0(p0 uint64) uint32 {
    if false {
        if true {
            v7 := GenS1{F0: 42, F1: 131, F2: 2147483647}
            if true {
                v8 := []int{32768, 127, 32767}
                if len(v8) > 0 {
                    v8[0] = 32767
                }
                _ = v8
            } else {
                var v9 uint64 = 1
                v10 := []int{-706}
                var v11 bool = true
                _ = v10
                _ = v11
                _ = v9
            }
            _ = v7
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 string = ("" + "")
                if (false || false) {
                    switch (fi12 % (((fi12 & fi12)) | 1)) {
                    case 2147483647:
                        fmt.Println(fi12, v13)
                    case 32767:
                        for fi14 := 0; fi14 < 8; fi14++ {
                            fmt.Println(fi12)
                            fmt.Println(v13, fi14, v13)
                            var v15 uint8 = 1
                            _ = fi14
                            _ = v15
                        }
                    }
                }
                _ = fi12
                _ = v13
            }
            switch "foo bar" {
            case "":
                switch 4294967295 {
                case 42:
                    for fi16 := 0; fi16 < 11; fi16++ {
                        p0 = p0
                        v17 := []int{fi16, (fi16 & (fi16 * (fi16 | (fi16 - (^(fi16)))))), fi16, (-(fi16))}
                        _ = fi16
                        _ = v17
                    }
                case -104:
                    v18 := []int{-2147483648}
                    _ = v18
                }
            case "hello":
                var v19 int64 = 0
                _ = v19
            case "a":
                p0 = p0
            default:
                v20 := GenS1{F0: -972, F1: 42, F2: int32(p0)}
                _ = v20
            }
            fmt.Println(p0)
        }
    }
    p0 = 499
    var v21 uint32 = 2
    _ = p0
    _ = v21
    return (^(v21))
}
func genFn1(p0 bool, p1 uint64) bool {
    var v22 bool = ((p0 && ("hello" != ((("hello" + "\n\t") + "") + ("\n\t" + ("a" + ("a" + "\n\t")))))) || (p0 || true))
    v23 := GenS2{F0: (0 <= 9223372036854775807)}
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return ((-9223372036854775808 > 32768) && v22)
}

func main() {
    var v1 uint16 = 255
    var v2 string = ""
    var v3 int8 = -81
    v4 := make(map[int]int)
    var v5 uint = 4294967295
    var v6 int64 = 394
    v5 -= 4294967295
    switch ((v2 + (v2 + ("foo bar" + v2))) + (v2 + ("hello" + v2))) {
    case "\n\t":
        v24 := []int{1, 4294967295, 4294967295, -2147483648}
        _ = v24
    default:
        v5 = v5
    }
    v25 := GenS3{F0: uint32(v5), F1: 42, F2: 42}
    if false {
        if ((true || (2147483647 == -2147483648)) && (!false)) {
            fmt.Println(v3, v3, v3)
            v2 = v2
        } else {
            switch v2 {
            case "a":
                v26 := GenS3{F0: 4294967295, F1: genFn0(1), F2: 42}
                _ = v26
            case "foo bar":
                v5 += (v5 | 42)
            case "\n\t":
                if true {
                    fmt.Println(v1)
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v3 -= (v3 * -81)
                        v3 += (v3 + v3)
                        _ = fi27
                    }
                }
            }
            var v28 string = (v2 + v2)
            _ = v28
        }
        v29 := v4[435]
        var v30 int = (v29 - 933)
        _ = v29
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 8; fi31++ {
            if false {
                v32 := GenS4{F0: "\n\t"}
                var v33 uint32 = genFn0(1)
                v34 := GenS3{F0: v33, F1: genFn0(4294967295), F2: -2147483648}
                _ = v32
                _ = v33
                _ = v34
            } else {
                fmt.Println(v1, v1, v5)
                v6 = 442
                delete(v4, fi31)
            }
            switch fi31 {
            case -71:
                v4[(^(fi31))] = 32768
            case 2147483647:
                v35 := GenS3{F0: 65535, F1: genFn0(979), F2: -730}
                _ = v35
            }
            _ = fi31
        }
    }
    if (1e10 >= (((1.5 - (1.5 - 3.14159)) * 0.1) + 1.5)) {
        if (true || false) {
            var v36 uint = 4294967295
            switch 127 {
            case 32768:
                v5 += v36
            }
            _ = v36
        }
        var v37 string = (v2 + v2)
        _ = v37
    }
    v2 = "\n\t"
    v38 := GenS5{F0: false, F1: false}
    switch -1 {
    case 4294967295:
        var v39 int32 = 0
        _ = v39
    case -9223372036854775808:
        switch v2 {
        case "hello":
            fmt.Println(v3, v1)
        case "\n\t":
            v3 -= (v3 / (((v3 / (((v3 & 127)) | 1))) | 1))
        case "":
            if false {
                switch "hello" {
                case "hello":
                    v6 = v6
                case "a":
                    switch 127 {
                    case -1:
                        v40 := GenS6{F0: 9223372036854775807, F1: -533}
                        _ = v40
                    case 42:
                        var v41 int8 = -1
                        _ = v41
                    case 128:
                        v42 := GenS2{F0: ((42 > 253) || true)}
                        _ = v42
                    }
                }
                v43 := GenS4{F0: (("foo bar" + (((v2 + ("" + v2)) + "a") + (v2 + "hello"))) + v2)}
                v5 = v5
                _ = v43
            }
        default:
            var v44 bool = false
            _ = v44
        }
    }
    if false {
        if true {
            var v45 uint = (^(v5))
            _ = v45
        }
    } else {
        if true {
            v3 += v3
            v3 = v3
            var v46 int32 = 271
            _ = v46
        }
        if ((170 < 0) && true) {
            switch -30 {
            case -2147483648:
                switch v2 {
                case "":
                    v47 := v4[-1]
                    _ = v47
                case "\n\t":
                    v5 += 0
                case "a":
                    switch -967 {
                    case 848:
                        if true {
                            switch "hello" {
                            case "":
                                v5 -= (v5 - (v5 | v5))
                            case "hello":
                                switch "a" {
                                case "foo bar":
                                    v48 := v4[516]
                                    _ = v48
                                case "a":
                                    v49 := []int{42, 9223372036854775807, 42}
                                    _ = v49
                                case "":
                                    var v50 string = "foo bar"
                                    _ = v50
                                }
                            default:
                                v5 = (v5 / (((v5 << ((v5) & 63))) | 1))
                            }
                            var v51 int32 = -1
                            _ = v51
                        }
                    case -1:
                        if false {
                            var v52 uint16 = v1
                            if (!true) {
                                var v53 uint16 = v52
                                if false {
                                    if ((((3.14159 / 3.14159) - (3.14159 + 1.5)) * float64(((0.1 + 0.1) * 1e10))) >= (0.0 * ((1.5 + -2.25) / 3.14159))) {
                                        fmt.Println(v6)
                                        for fi54 := 0; fi54 < 12; fi54++ {
                                            if ((^(v1)) < v1) {
                                                v3 = v3
                                                var v55 int8 = v3
                                                _ = v55
                                            }
                                            v53 = (v53 % (((v52 ^ v53)) | 1))
                                            _ = fi54
                                        }
                                        for fi56 := 0; fi56 < 2; fi56++ {
                                            for fi57 := 0; fi57 < 18; fi57++ {
                                                if false {
                                                    if false {
                                                        var v58 float32 = (float32(float64(1e10)) - ((1.5 * float32(0.0)) * ((0.1 / float32(-2.25)) + 3.14159)))
                                                        v59 := GenS7{F0: 1000, F1: 815, F2: true}
                                                        var v60 uint16 = 318
                                                        switch "foo bar" {
                                                        case "a":
                                                            switch v2 {
                                                            case "hello":
                                                                fmt.Println(v1, v53, v53)
                                                            case "":
                                                                v61 := []int{}
                                                                _ = v61
                                                            default:
                                                                fmt.Println(v2, v1, v1)
                                                            }
                                                        case "\n\t":
                                                            if (!(!false)) {
                                                                switch fi57 {
                                                                case 791:
                                                                    fmt.Println(v3, fi57)
                                                                case 4294967295:
                                                                    v62 := v4[fi56]
                                                                    _ = v62
                                                                }
                                                            } else {
                                                                switch 32768 {
                                                                case 32767:
                                                                    var v63 uint16 = v52
                                                                    _ = v63
                                                                case -2147483648:
                                                                    v64 := []int{fi57, fi56}
                                                                    _ = v64
                                                                }
                                                                if true {
                                                                    fmt.Println(v60, v1)
                                                                }
                                                            }
                                                        default:
                                                            v65 := GenS8{F0: 0, F1: 770, F2: (v6 - (v6 / ((int64(v5)) | 1)))}
                                                            _ = v65
                                                        }
                                                        v3 += v3
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    v6 = 9223372036854775807
                                                }
                                                _ = fi57
                                            }
                                            v1 = v52
                                            _ = fi56
                                        }
                                    }
                                    if (!true) {
                                        fmt.Println(v5, v53)
                                        if false {
                                            if true {
                                                switch 2147483647 {
                                                case 2147483647:
                                                    v5 = (v5 & 437)
                                                case -173:
                                                    switch (v2 + v2) {
                                                    case "foo bar":
                                                        v66 := GenS9{F0: v1, F1: 793, F2: 42}
                                                        _ = v66
                                                    case "":
                                                        v67 := GenS5{F0: true, F1: false}
                                                        _ = v67
                                                    }
                                                }
                                            } else {
                                                v68 := GenS6{F0: -9223372036854775808, F1: 32768}
                                                v5 = (v5 % (((^(v5))) | 1))
                                                _ = v68
                                            }
                                            if ((0 == v6) || false) {
                                                v3 += v3
                                                v52 = 65535
                                                switch 32767 {
                                                case 2147483647:
                                                    var v69 int16 = 1
                                                    _ = v69
                                                case -481:
                                                    v53 -= v53
                                                case 4294967295:
                                                    var v70 uint32 = 796
                                                    _ = v70
                                                default:
                                                    var v71 int64 = (v6 << (((v5 % ((v5) | 1))) & 63))
                                                    _ = v71
                                                }
                                            } else {
                                                var v72 float64 = 3.14159
                                                _ = v72
                                            }
                                            var v73 int = 32768
                                            _ = v73
                                        } else {
                                            v74 := GenS1{F0: (-(v6)), F1: -32768, F2: 1}
                                            if (float64((-2.25 - 0.1)) < -2.25) {
                                                fmt.Println(v3, v2)
                                            }
                                            var v75 int16 = 42
                                            _ = v74
                                            _ = v75
                                        }
                                    } else {
                                        var v76 string = v2
                                        switch 353 {
                                        case -9223372036854775808:
                                            if true {
                                                v53 = v1
                                            } else {
                                                fmt.Println(v53, v5)
                                                var v77 int16 = 1
                                                _ = v77
                                            }
                                        case -461:
                                            v6 = v6
                                        case 0:
                                            fmt.Println(v1)
                                        }
                                        _ = v76
                                    }
                                    v53 = (v53 % ((v1) | 1))
                                }
                                _ = v53
                            }
                            v78 := GenS10{F0: (255 <= 42), F1: 127, F2: 42}
                            _ = v52
                            _ = v78
                        } else {
                            switch "hello" {
                            case "\n\t":
                                if false {
                                    v1 = 65535
                                } else {
                                    var v79 uint64 = 787
                                    for fi80 := 0; fi80 < 14; fi80++ {
                                        var v81 uint64 = (v79 & (v79 * 4294967295))
                                        switch (fi80 - (-(fi80))) {
                                        case 59:
                                            v2 = v2
                                        case -691:
                                            v6 -= 0
                                        default:
                                            v5 = (v5 >> (((v5 & (^(v5)))) & 63))
                                        }
                                        _ = fi80
                                        _ = v81
                                    }
                                    _ = v79
                                }
                            case "hello":
                                switch v2 {
                                case "foo bar":
                                    v82 := GenS11{F0: (v6 & v6), F1: v1, F2: 0}
                                    _ = v82
                                case "\n\t":
                                    v6 -= v6
                                case "":
                                    for fi83 := 0; fi83 < 7; fi83++ {
                                        var v84 bool = (2147483647 < 1)
                                        v85 := []int{int(v5), (fi83 ^ (-(fi83))), (fi83 * fi83), (fi83 % ((fi83) | 1))}
                                        v2 = ((v2 + v2) + "")
                                        _ = fi83
                                        _ = v84
                                        _ = v85
                                    }
                                }
                            }
                        }
                    case 42:
                        v86 := v4[-687]
                        _ = v86
                    default:
                        for fi87 := 0; fi87 < 17; fi87++ {
                            fmt.Println(v6, fi87, v3)
                            _ = fi87
                        }
                    }
                }
            case 9223372036854775807:
                var v88 string = "hello"
                _ = v88
            case 2147483647:
                for fi89 := 0; fi89 < 1; fi89++ {
                    for fi90 := 0; fi90 < 18; fi90++ {
                        fmt.Println(v5)
                        v3 = v3
                        if false {
                            if true {
                                fmt.Println(fi89)
                            }
                            v91 := []int{4294967295, (fi89 << ((v5) & 63)), (fi90 / ((fi89) | 1))}
                            _ = v91
                        }
                        if true {
                            v2 = "hello"
                        } else {
                            delete(v4, 2147483647)
                            if false {
                                var v92 int16 = 0
                                switch (^(fi90)) {
                                case 0:
                                    if false {
                                        if true {
                                            var v93 uint8 = 0
                                            switch fi90 {
                                            case -1:
                                                v6 += v6
                                            case 9223372036854775807:
                                                fmt.Println(v93, v93)
                                            case 1:
                                                var v94 float64 = (0.1 - -2.25)
                                                _ = v94
                                            }
                                            _ = v93
                                        }
                                        if false {
                                            var v95 int16 = 42
                                            _ = len(v4)
                                            _ = v95
                                        }
                                    } else {
                                        var v96 int16 = (-(v92))
                                        _ = v96
                                    }
                                case 32767:
                                    var v97 bool = false
                                    _ = v97
                                default:
                                    v2 = v2
                                }
                                _ = v92
                            }
                        }
                        _ = fi90
                    }
                    for fi98 := 0; fi98 < 19; fi98++ {
                        var v99 float32 = float32((1e10 + float64(1e10)))
                        fmt.Println(v2, v6)
                        _ = fi98
                        _ = v99
                    }
                    if (-2.25 < ((1e10 - ((float64(0.0) / (0.1 + (1e10 * 1e10))) * 0.0)) - 1.5)) {
                        for fi100 := 0; fi100 < 17; fi100++ {
                            v5 = (v5 * 18446744073709551615)
                            var v101 int8 = (-(v3))
                            _ = fi100
                            _ = v101
                        }
                        for fi102 := 0; fi102 < 20; fi102++ {
                            switch fi89 {
                            case 9223372036854775807:
                                fmt.Println(fi102)
                            case 32767:
                                var v103 string = v2
                                _ = v103
                            default:
                                var v104 int32 = 2147483647
                                _ = v104
                            }
                            switch "" {
                            case "\n\t":
                                var v105 float64 = 0.0
                                _ = v105
                            case "foo bar":
                                var v106 uint64 = 42
                                _ = v106
                            case "":
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v2)
                                case "\n\t":
                                    var v107 uint16 = (v1 & (v1 * (^(v1))))
                                    _ = v107
                                case "hello":
                                    if true {
                                        fmt.Println(fi89, v5)
                                        v1 += (^(v1))
                                    }
                                }
                                var v108 uint16 = v1
                                _ = v108
                            default:
                                v109 := GenS6{F0: int(v1), F1: fi102}
                                _ = v109
                            }
                            fmt.Println(v2, v3)
                            _ = fi102
                        }
                    }
                    _ = fi89
                }
            default:
                v1 = (v1 | (^(v1)))
            }
            v1 = v1
        }
        v4[int(v5)] = -2147483648
    }
    var v110 uint16 = v1
    v110 -= 1
    v2 = ("hello" + (("hello" + (("hello" + v2) + v2)) + ""))
    switch int(v110) {
    case -625:
        v111 := v4[32767]
        _ = v111
    case 32767:
        for fi112 := 0; fi112 < 1; fi112++ {
            if true {
                fmt.Println(fi112, v1)
            } else {
                v3 = (v3 % ((v3) | 1))
            }
            v3 -= (v3 % (((v3 - v3)) | 1))
            _ = fi112
        }
    case 0:
        var v113 int32 = 838
        _ = v113
    }
    var v114 int8 = (v3 / (((v3 | -128)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v110
    _ = v114
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
