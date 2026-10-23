package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint64; F2 string }
type GenS6 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 bool; F1 string; F2 string }
func genFn0(p0 uint64) uint16 {
    var v4 int32 = -2147483648
    v4 = v4
    _ = p0
    _ = v4
    return 425
}
func genFn1(p0 uint8, p1 uint) uint8 {
    v5 := GenS1{F0: -183}
    p1 = p1
    v6 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v7 int16 = 447
    v7 = v7
    v8 := GenS1{F0: -315}
    v9 := GenS1{F0: -28}
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := []int{}
        if false {
            if (!(65535 == 842)) {
                v12 := GenS1{F0: -185}
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 bool = false
                    if (4294967295 == 883) {
                        var v15 float64 = float64((3.14159 - float32(3.14159)))
                        var v16 uint = 869
                        var v17 int = fi13
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        switch fi13 {
                        case 127:
                            v7 += (v7 ^ v7)
                        case 128:
                            var v18 float64 = float64(1.5)
                            _ = v18
                        default:
                            v11 = append(v11, (fi13 ^ (fi13 >> ((18446744073709551615) & 63))))
                        }
                    }
                    v14 = ((!(("foo bar" + "") >= (("hello" + ("foo bar" + "hello")) + ""))) || true)
                    _ = fi13
                    _ = v14
                }
                switch "\n\t" {
                case "foo bar":
                    v7 -= 0
                }
                _ = v12
            }
            for fi19 := 0; fi19 < 8; fi19++ {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(fi10, v7, v7)
                default:
                    if (128 > 127) {
                        v7 -= v7
                        v7 = v7
                        _ = len(v11)
                    } else {
                        var v20 uint32 = 65535
                        if ((((fi10 <= fi19) || ((317 < 65535) || true)) || true) || (v20 != (v20 * 42))) {
                            fmt.Println(v20, fi19)
                        }
                        _ = v20
                    }
                }
                v21 := GenS1{F0: 42}
                _ = len(v11)
                _ = fi19
                _ = v21
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v7, fi22, fi10)
                _ = fi22
            }
        } else {
            fmt.Println(fi10, fi10)
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 int64 = 559
                v25 := GenS1{F0: (v24 + v24)}
                var v26 uint32 = 771
                switch "\n\t" {
                case "\n\t":
                    v27 := GenS1{F0: (v24 ^ 42)}
                    _ = v27
                case "foo bar":
                    var v28 uint32 = v26
                    _ = v28
                }
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        _ = fi10
        _ = v11
    }
    switch (("foo bar" + ("" + "hello")) + ("" + "")) {
    case "\n\t":
        switch 32767 {
        case -1:
            var v29 uint = 65535
            _ = v29
        }
    case "a":
        for fi30 := 0; fi30 < 12; fi30++ {
            v7 -= (v7 - v7)
            _ = fi30
        }
    case "foo bar":
        fmt.Println(v7)
    default:
        if true {
            if true {
                var v31 float32 = 0.0
                for fi32 := 0; fi32 < 15; fi32++ {
                    if true {
                        var v33 uint32 = 731
                        if false {
                            v34 := []int{(fi32 / ((fi32) | 1)), fi32}
                            var v35 uint64 = 4294967295
                            if true {
                                var v36 bool = (127 < 89)
                                v36 = v36
                                _ = v36
                            } else {
                                v31 = (float32(-2.25) + 1e10)
                                var v37 int16 = v7
                                _ = v37
                            }
                            _ = v34
                            _ = v35
                        } else {
                            v7 -= v7
                            v2["foo bar"] = true
                        }
                        _ = v33
                    } else {
                        fmt.Println(fi32, v7, v31)
                    }
                    _ = fi32
                }
                _ = v31
            }
            if (!(1e10 == 3.14159)) {
                v7 -= v7
            } else {
                if (v7 >= 1000) {
                    delete(v3, 4294967295)
                    if false {
                        v38 := GenS1{F0: -683}
                        v7 += 1000
                        v7 -= 1
                        _ = v38
                    } else {
                        v7 -= v7
                        var v39 int32 = 172
                        v40 := GenS2{F0: 127, F1: genFn0(4294967295)}
                        _ = v39
                        _ = v40
                    }
                    var v41 uint8 = 127
                    _ = v41
                } else {
                    v7 = 1000
                    for fi42 := 0; fi42 < 7; fi42++ {
                        v7 = -102
                        fmt.Println(fi42, fi42)
                        if ((!(-53 < 0)) && true) {
                            v43 := GenS3{F0: (true || (false && false)), F1: 128}
                            _ = v43
                        } else {
                            switch ("" + "a") {
                            case "hello":
                                fmt.Println(v7, fi42)
                            case "a":
                                var v44 string = ""
                                _ = v44
                            case "foo bar":
                                if false {
                                    var v45 int16 = (-(v7))
                                    v45 -= v7
                                    _ = v45
                                }
                            }
                            v46 := []int{}
                            _ = v46
                        }
                        _ = fi42
                    }
                    if (false || true) {
                        var v47 uint64 = 0
                        if false {
                            v47 += (v47 << ((18446744073709551615) & 63))
                        } else {
                            var v48 float32 = 3.14159
                            var v49 int64 = 256
                            v2["foo bar"] = false
                            _ = v48
                            _ = v49
                        }
                        _ = v47
                    }
                }
                var v50 int8 = 1
                _ = len(v1)
                _ = v50
            }
        }
    }
    var v51 int8 = -128
    if false {
        v51 = (v51 >> ((664) & 63))
        var v52 int = 492
        v53 := []int{v52}
        _ = v52
        _ = v53
    }
    fmt.Println(v7, v51)
    if (!(2 <= 42)) {
        v7 += v7
        if false {
            for fi54 := 0; fi54 < 16; fi54++ {
                switch ("" + "foo bar") {
                case "":
                    var v55 bool = false
                    _ = v55
                case "\n\t":
                    v51 = (v51 & (v51 + (v51 - (v51 >> ((134) & 63)))))
                case "foo bar":
                    fmt.Println(v51, fi54)
                default:
                    v56 := GenS4{F0: 65535}
                    _ = v56
                }
                _ = fi54
            }
        }
        switch (("" + ((("a" + "\n\t") + "") + "\n\t")) + "hello") {
        case "\n\t":
            if true {
                var v57 uint32 = 634
                v58 := []int{-689}
                _ = v57
                _ = v58
            }
        case "":
            switch "foo bar" {
            case "a":
                v59 := GenS2{F0: 127, F1: 795}
                _ = v59
            case "foo bar":
                v51 = v51
            case "\n\t":
                v51 = (v51 & v51)
            }
        case "hello":
            if (-2147483648 < 408) {
                v60 := v3[2147483648]
                v61 := GenS5{F0: 42, F1: 35, F2: "foo bar"}
                _ = v60
                _ = v61
            }
        }
    }
    v51 = v51
    if true {
        v62 := GenS2{F0: 0, F1: 2}
        _ = v62
    } else {
        var v63 int64 = 9223372036854775807
        fmt.Println(v7, v7, v7)
        for fi64 := 0; fi64 < 14; fi64++ {
            for fi65 := 0; fi65 < 13; fi65++ {
                var v66 float32 = (-2.25 - (0.0 * 1e10))
                switch "\n\t" {
                case "foo bar":
                    if ("a" == "") {
                        var v67 int = fi64
                        var v68 int = 32767
                        _ = v67
                        _ = v68
                    }
                case "a":
                    v69 := v3[2147483647]
                    _ = v69
                default:
                    v7 += (v7 ^ (v7 - 42))
                    if (!false) {
                        if true {
                            fmt.Println(v63)
                            if (v7 < (^(v7))) {
                                v51 = (v51 % (((v51 & (v51 / (((v51 ^ v51)) | 1)))) | 1))
                                if false {
                                    v70 := GenS2{F0: (fi64 & (^(fi65))), F1: 1}
                                    var v71 uint16 = 663
                                    var v72 uint16 = (v71 - v71)
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                }
                                if (1 >= 225) {
                                    v73 := GenS6{F0: 471, F1: v7, F2: 4294967295}
                                    _ = v73
                                }
                            } else {
                                switch ("" + "a") {
                                case "a":
                                    v66 = float32((3.14159 + (((-2.25 * 0.1) * (3.14159 + 3.14159)) + (1e10 * 0.0))))
                                case "foo bar":
                                    v63 += v63
                                case "":
                                    if (4294967295 == 4294967295) {
                                        var v74 uint8 = genFn1(128, 18446744073709551615)
                                        var v75 uint8 = v74
                                        _ = v74
                                        _ = v75
                                    } else {
                                        v76 := []int{}
                                        if len(v76) >= 2 {
                                            v76 = v76[1:len(v76)]
                                        }
                                        var v77 uint16 = genFn0(42)
                                        v7 = (v7 * (v7 ^ 42))
                                        _ = v76
                                        _ = v77
                                    }
                                }
                                switch fi65 {
                                case 127:
                                    v78 := []int{(^(fi65))}
                                    _ = v78
                                }
                                if ((((1e10 * v66) / 1e10) * ((float32((float64(1.5) / float64(float32(float64((float32((((3.14159 * 0.1) + 0.0) + 1.5)) + v66)))))) / 1.5) - (v66 + v66))) == ((0.1 / (v66 * -2.25)) * 1.5)) {
                                    v7 = (v7 ^ v7)
                                    fmt.Println(fi65, v51)
                                    _ = len(v2)
                                }
                            }
                        }
                    } else {
                        fmt.Println(fi65, fi65, fi65)
                    }
                }
                v51 = v51
                _ = fi65
                _ = v66
            }
            v79 := GenS5{F0: (fi64 & fi64), F1: 323, F2: "\n\t"}
            _ = fi64
            _ = v79
        }
        _ = v63
    }
    for fi80 := 0; fi80 < 8; fi80++ {
        var v81 bool = false
        _ = fi80
        _ = v81
    }
    for fi82 := 0; fi82 < 3; fi82++ {
        var v83 bool = (4294967295 < 748)
        switch ("" + ("foo bar" + "")) {
        case "":
            if v83 {
                fmt.Println(v83)
                for fi84 := 0; fi84 < 6; fi84++ {
                    switch (("foo bar" + "") + "hello") {
                    case "foo bar":
                        fmt.Println(fi82, v83)
                    }
                    var v85 uint64 = 495
                    _ = fi84
                    _ = v85
                }
                v86 := GenS5{F0: fi82, F1: 0, F2: ((("\n\t" + "\n\t") + "foo bar") + "a")}
                _ = v86
            }
        case "\n\t":
            v87 := GenS2{F0: fi82, F1: 634}
            _ = v87
        case "foo bar":
            if true {
                switch "a" {
                case "":
                    var v88 uint = 255
                    _ = v88
                case "a":
                    for fi89 := 0; fi89 < 15; fi89++ {
                        v83 = (("a" + "") <= "\n\t")
                        v90 := GenS1{F0: -811}
                        if v83 {
                            _ = len(v1)
                            v91 := []int{(fi89 & (-(fi89))), fi82, (fi82 ^ (fi82 << ((4294967295) & 63))), fi82}
                            _ = v91
                        }
                        _ = fi89
                        _ = v90
                    }
                case "hello":
                    var v92 uint = 946
                    _ = v92
                }
                fmt.Println(v51, v7, v83)
                v93 := v3[0]
                _ = v93
            } else {
                if false {
                    var v94 uint64 = 229
                    var v95 float64 = 0.1
                    v96 := []int{}
                    _ = v94
                    _ = v95
                    _ = v96
                } else {
                    var v97 uint8 = 143
                    switch "hello" {
                    case "hello":
                        v98 := v3[2147483647]
                        _ = v98
                    case "":
                        switch fi82 {
                        case 4294967295:
                            delete(v1, ("a" + ""))
                        case 0:
                            v99 := []int{(fi82 - fi82), 2147483647}
                            _ = v99
                        }
                    case "foo bar":
                        var v100 int16 = (-(v7))
                        _ = v100
                    default:
                        if true {
                            switch (fi82 / (((fi82 / ((fi82) | 1))) | 1)) {
                            case 128:
                                v83 = true
                            default:
                                v83 = v83
                            }
                        } else {
                            v101 := v2[((("hello" + "foo bar") + ("" + "")) + ("foo bar" + ("" + "hello")))]
                            _ = v101
                        }
                    }
                    _ = v97
                }
                v7 = (v7 ^ v7)
                if (-1 < (v7 << ((18446744073709551615) & 63))) {
                    switch 127 {
                    case -1:
                        v102 := GenS7{F0: 255, F1: 42}
                        _ = v102
                    case 126:
                        var v103 string = "\n\t"
                        _ = v103
                    }
                    if v83 {
                        for fi104 := 0; fi104 < 17; fi104++ {
                            v105 := GenS8{F0: 2}
                            v51 += 0
                            _ = fi104
                            _ = v105
                        }
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            if (!false) {
                                switch -487 {
                                case 42:
                                    v106 := []int{-2147483648}
                                    _ = v106
                                default:
                                    var v107 uint32 = 0
                                    _ = v107
                                }
                            }
                        }
                    }
                    if (!(-2.25 >= 0.1)) {
                        var v108 uint = 0
                        for fi109 := 0; fi109 < 13; fi109++ {
                            var v110 int16 = (v7 >> (((v108 & (v108 * v108))) & 63))
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                if (genFn1(genFn1(2, v108), 42) <= 1) {
                                    v51 = -1
                                    switch 988 {
                                    case 1:
                                        var v111 uint8 = 0
                                        _ = v111
                                    case 128:
                                        v112 := GenS9{F0: (v51 % ((v51) | 1)), F1: (v110 / ((v7) | 1))}
                                        _ = v112
                                    default:
                                        if v83 {
                                            if (v83 || v83) {
                                                var v113 uint8 = 224
                                                _ = v113
                                            } else {
                                                var v114 uint64 = 2
                                                v115 := GenS3{F0: ((-(v51)) != (v51 / ((127) | 1))), F1: genFn1(127, 2)}
                                                if v83 {
                                                    v116 := GenS10{F0: uint16(v114)}
                                                    fmt.Println(v7, v51, v7)
                                                    v51 = (v51 | v51)
                                                    _ = v116
                                                } else {
                                                    v117 := v1[""]
                                                    v118 := []int{(fi82 >> (((v108 << ((v108) & 63))) & 63)), (v117 / (((^(fi109))) | 1)), (fi82 * 9223372036854775807), v117}
                                                    fmt.Println(v110, v117, v110)
                                                    _ = v117
                                                    _ = v118
                                                }
                                                _ = v114
                                                _ = v115
                                            }
                                            switch 419 {
                                            case 2147483647:
                                                if ("hello" <= (("" + "a") + "\n\t")) {
                                                    v119 := []int{(-(fi82))}
                                                    _ = v119
                                                }
                                            case 9223372036854775807:
                                                v120 := v3[2147483647]
                                                _ = v120
                                            }
                                            fmt.Println(v108, v83)
                                            v121 := GenS11{F0: ((v7 ^ v110) == v7), F1: "\n\t", F2: "hello"}
                                            _ = v121
                                        }
                                    }
                                    var v122 uint32 = 4294967295
                                    _ = v122
                                }
                                var v123 uint16 = 394
                                fmt.Println(fi82, v51)
                                var v124 float64 = 1.5
                                var v125 int = (fi82 * (-(fi82)))
                                var v126 int32 = 42
                                _ = v123
                                _ = v124
                                _ = v125
                                _ = v126
                            default:
                                v51 += v51
                            }
                            var v127 uint8 = 93
                            _ = fi109
                            _ = v110
                            _ = v127
                        }
                        _ = v108
                    }
                }
            }
        }
        for fi128 := 0; fi128 < 13; fi128++ {
            fmt.Println(v83, v83, v83)
            v129 := []int{(fi82 * fi82), fi128}
            v130 := GenS7{F0: 0, F1: 627}
            _ = fi128
            _ = v129
            _ = v130
        }
        _ = fi82
        _ = v83
    }
    for fi131 := 0; fi131 < 16; fi131++ {
        var v132 float64 = -2.25
        v132 = v132
        fmt.Println(v7)
        _ = fi131
        _ = v132
    }
    v51 -= (^(v51))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v7
    _ = v8
    _ = v9
}
