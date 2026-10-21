package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint32; F2 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int32; F1 uint8 }
type GenS14 struct { F0 bool; F1 uint }
type GenS15 struct { F0 int64; F1 int64; F2 int8 }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int }
type GenS19 struct { F0 int32; F1 uint }
type GenS20 struct { F0 uint32; F1 int16 }
type GenS21 struct { F0 int8; F1 string }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 int; F1 uint16 }
func genFn0(p0 bool) int16 {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 1}
        p0 = p0
        p0 = (904 != 1)
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 1}
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(p0, p0)
        p0 = (0 > 969)
        _ = fi8
    }
    if p0 {
        v9 := []int{}
        for fi10 := 0; fi10 < 12; fi10++ {
            p0 = p0
            _ = fi10
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v9
    } else {
        var v11 uint64 = 2
        var v12 float32 = 3.14159
        _ = v11
        _ = v12
    }
    switch "foo bar" {
    case "":
        p0 = p0
    case "hello":
        for fi13 := 0; fi13 < 20; fi13++ {
            if ((0 <= 0) || (false || (-2.25 <= float32(-2.25)))) {
                fmt.Println(fi13)
            } else {
                var v14 string = ""
                _ = v14
            }
            var v15 int8 = -128
            _ = fi13
            _ = v15
        }
    case "foo bar":
        fmt.Println(p0)
    default:
        v16 := GenS3{F0: 42, F1: 74}
        _ = v16
    }
    _ = p0
    _ = v7
    return 366
}

func main() {
    var v1 uint = 4294967295
    var v2 float64 = 0.1
    v3 := make(map[string]bool)
    var v4 int32 = 1
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int16 = 1
        _ = fi17
        _ = v18
    }
    v4 = 307
    _ = len(v3)
    v4 = v4
    v4 -= -1
    if (("" + ("hello" + ("" + "\n\t"))) != "foo bar") {
        if (false || true) {
            v2 += 0.1
        } else {
            var v19 int64 = 1
            v19 -= (v19 - v19)
            v20 := GenS4{F0: -128}
            _ = v19
            _ = v20
        }
        v21 := GenS3{F0: -813, F1: 100}
        _ = v21
    } else {
        var v22 float32 = 0.0
        for fi23 := 0; fi23 < 8; fi23++ {
            if (!false) {
                v1 -= v1
                v24 := v3[("\n\t" + "hello")]
                _ = v24
            } else {
                var v25 uint32 = 645
                _ = v25
            }
            _ = fi23
        }
        v26 := GenS2{F0: (v4 % ((int32(v1)) | 1))}
        _ = v22
        _ = v26
    }
    switch 1 {
    case 9223372036854775807:
        v1 = v1
    case 601:
        switch "\n\t" {
        case "":
            v27 := []int{-383, 769, -9223372036854775808}
            _ = v27
        default:
            v2 += float64(((1e10 / 1.5) - -2.25))
        }
    case -406:
        if true {
            fmt.Println(v1)
            for fi28 := 0; fi28 < 2; fi28++ {
                v2 += v2
                _ = fi28
            }
        } else {
            switch "a" {
            case "":
                if false {
                    v29 := GenS5{F0: -9223372036854775808}
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        fmt.Println(fi30)
                        v2 += ((0.0 / (v2 + v2)) * 1e10)
                        _ = fi30
                    }
                    v2 = (v2 / float64((1e10 / float32(float64(0.1)))))
                    if false {
                        v1 -= v1
                        switch "hello" {
                        case "hello":
                            v31 := GenS2{F0: 42}
                            _ = v31
                        case "\n\t":
                            v32 := GenS6{F0: (("" + "a") + "\n\t"), F1: 4294967295, F2: -1}
                            _ = v32
                        case "":
                            switch "hello" {
                            case "hello":
                                var v33 uint64 = 42
                                _ = v33
                            case "a":
                                v1 = (v1 >> (((v1 | 255)) & 63))
                            }
                        default:
                            fmt.Println(v1, v2)
                        }
                        var v34 int8 = -1
                        _ = v34
                    } else {
                        fmt.Println(v1, v1)
                    }
                }
            case "\n\t":
                if false {
                    for fi35 := 0; fi35 < 11; fi35++ {
                        switch "a" {
                        case "":
                            for fi36 := 0; fi36 < 1; fi36++ {
                                fmt.Println(v2, fi36, v1)
                                _ = fi36
                            }
                        case "\n\t":
                            v4 = (^(v4))
                        default:
                            v37 := []int{881, 42}
                            _ = v37
                        }
                        v4 = (v4 * 42)
                        _ = fi35
                    }
                    fmt.Println(v2, v1)
                    v38 := []int{2147483647, 2147483647}
                    _ = v38
                } else {
                    var v39 uint = (v1 & v1)
                    _ = v39
                }
            default:
                for fi40 := 0; fi40 < 16; fi40++ {
                    if true {
                        if ((890 != 865) && false) {
                            var v41 int8 = -42
                            _ = v41
                        } else {
                            for fi42 := 0; fi42 < 1; fi42++ {
                                fmt.Println(v4, v2)
                                _ = fi42
                            }
                            v43 := GenS4{F0: -107}
                            v44 := GenS4{F0: 9}
                            _ = v43
                            _ = v44
                        }
                        v4 += v4
                    }
                    _ = fi40
                }
            }
            for fi45 := 0; fi45 < 10; fi45++ {
                for fi46 := 0; fi46 < 14; fi46++ {
                    var v47 float32 = (float32((-2.25 + (v2 - float64(0.0)))) - float32(1.5))
                    _ = fi46
                    _ = v47
                }
                v48 := GenS6{F0: "\n\t", F1: 2, F2: genFn0((1 >= 65535))}
                fmt.Println(v1, v4)
                _ = fi45
                _ = v48
            }
            if (106 > 255) {
                for fi49 := 0; fi49 < 20; fi49++ {
                    var v50 bool = ((0.1 > 0.0) && (false && false))
                    _ = fi49
                    _ = v50
                }
            } else {
                v2 += ((float64(1.5) / (0.1 - v2)) * 3.14159)
                if false {
                    v51 := GenS6{F0: ("hello" + "foo bar"), F1: 0, F2: 32767}
                    _ = v51
                }
            }
        }
    }
    fmt.Println(v2, v1, v4)
    if true {
        var v52 float64 = (1e10 * 3.14159)
        _ = v52
    }
    if true {
        v53 := GenS7{F0: 2147483647}
        for fi54 := 0; fi54 < 5; fi54++ {
            if (!true) {
                if ("foo bar" == ("a" + "")) {
                    switch 42 {
                    case -262:
                        switch "foo bar" {
                        case "a":
                            switch fi54 {
                            case -606:
                                delete(v3, ("" + "\n\t"))
                            case 153:
                                v55 := GenS8{F0: 2147483647}
                                _ = v55
                            }
                        }
                    }
                    var v56 uint32 = 327
                    _ = v56
                } else {
                    if (-2147483648 <= v4) {
                        v57 := []int{fi54, (fi54 / ((9223372036854775807) | 1))}
                        for fi58 := 0; fi58 < 2; fi58++ {
                            v57 = append(v57, 42)
                            var v59 uint8 = uint8(v1)
                            delete(v3, "foo bar")
                            _ = fi58
                            _ = v59
                        }
                        if (-1 <= 42) {
                            v1 = v1
                        } else {
                            v2 = v2
                            if (false && true) {
                                for fi60 := 0; fi60 < 20; fi60++ {
                                    v4 = 1
                                    if false {
                                        switch fi54 {
                                        case 32768:
                                            v61 := GenS6{F0: ("hello" + "hello"), F1: uint32(v1), F2: 32767}
                                            _ = v61
                                        case -9223372036854775808:
                                            v2 = (0.1 * v2)
                                            fmt.Println(v2, v1)
                                        case 4294967295:
                                            var v62 float64 = 0.0
                                            _ = v62
                                        }
                                        v1 = 2
                                        var v63 uint32 = 4294967295
                                        var v64 uint64 = 1
                                        v65 := GenS3{F0: 32767, F1: 1}
                                        if ((v1 ^ 42) < v1) {
                                            if false {
                                                if (!false) {
                                                    v66 := GenS6{F0: "foo bar", F1: (v63 * v63), F2: genFn0(false)}
                                                    if (!(false || true)) {
                                                        if (true && false) {
                                                            v2 = 1.5
                                                            switch "\n\t" {
                                                            case "hello":
                                                                fmt.Println(v63, v4, v2)
                                                            case "a":
                                                                v64 -= (v64 * v64)
                                                            case "":
                                                                if true {
                                                                    if len(v57) > 0 {
                                                                        v57[0] = int(v63)
                                                                    }
                                                                    v67 := GenS3{F0: -1, F1: 0}
                                                                    _ = v67
                                                                }
                                                            default:
                                                                var v68 int8 = -1
                                                                v69 := GenS9{F0: fi60}
                                                                var v70 uint = (v1 - (v1 + 4294967295))
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                        } else {
                                                            fmt.Println(v64, v1, v63)
                                                            v1 -= (v1 & (v1 + 255))
                                                        }
                                                        var v71 float64 = 1e10
                                                        v72 := v3["foo bar"]
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v73 uint = v1
                                                            _ = v73
                                                        }
                                                        v63 -= v63
                                                    }
                                                    var v74 float32 = 0.1
                                                    if true {
                                                        var v75 int16 = 1
                                                        if (!true) {
                                                            var v76 int = fi54
                                                            if ((-(v4)) == -79) {
                                                                v77 := GenS2{F0: 401}
                                                                if ((-365 >= 42) || true) {
                                                                    var v78 float32 = v74
                                                                    v74 -= (v74 + v78)
                                                                    _ = v78
                                                                } else {
                                                                    var v79 int8 = -1
                                                                    if true {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v80 := GenS3{F0: 0, F1: (v79 >> (((v1 & 4294967295)) & 63))}
                                                                            _ = v80
                                                                        case "\n\t":
                                                                            switch fi54 {
                                                                            case 2147483647:
                                                                                v63 = (v63 | v63)
                                                                            case 32768:
                                                                                v76 = v76
                                                                            }
                                                                        case "a":
                                                                            _ = len(v3)
                                                                        default:
                                                                            v4 = (v4 / (((-(v4))) | 1))
                                                                        }
                                                                    } else {
                                                                        var v81 int = v76
                                                                        v1 = v1
                                                                        if ((v81 * 649) <= v81) {
                                                                            var v82 int8 = (v79 << (((v1 / ((v1) | 1))) & 63))
                                                                            _ = v82
                                                                        } else {
                                                                            if false {
                                                                                v83 := GenS10{F0: 0}
                                                                                v2 += float64((((v74 / 1e10) - 0.1) / 1.0))
                                                                                v2 = 1.5
                                                                                _ = v83
                                                                            } else {
                                                                                if len(v57) > 0 {
                                                                                    v57[0] = 9223372036854775807
                                                                                }
                                                                                switch v76 {
                                                                                case 32768:
                                                                                    v3["a"] = ((v1 >= v1) && false)
                                                                                }
                                                                            }
                                                                            switch (fi60 | (fi54 << (((v1 + (^(v1)))) & 63))) {
                                                                            case 42:
                                                                                v75 -= (v75 >> ((255) & 63))
                                                                            default:
                                                                                switch "a" {
                                                                                case "a":
                                                                                    if false {
                                                                                        if false {
                                                                                            _ = len(v3)
                                                                                            if true {
                                                                                                v84 := GenS11{F0: 65535, F1: 2147483647, F2: v75}
                                                                                                if true {
                                                                                                    switch (fi54 * 0) {
                                                                                                    case 610:
                                                                                                        switch "foo bar" {
                                                                                                        case "":
                                                                                                            switch ("foo bar" + ((("" + "a") + "foo bar") + "\n\t")) {
                                                                                                            case "":
                                                                                                                var v85 float32 = v74
                                                                                                                _ = v85
                                                                                                            case "foo bar":
                                                                                                                if (((v4 ^ (-(v4))) != v4) || false) {
                                                                                                                    v86 := GenS12{F0: 620}
                                                                                                                    _ = v86
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            var v87 int16 = genFn0((255 != 0))
                                                                                                            _ = v87
                                                                                                        }
                                                                                                    case 128:
                                                                                                        v63 -= (v63 - 2)
                                                                                                    case 9223372036854775807:
                                                                                                        if false {
                                                                                                            v88 := GenS13{F0: (v4 - -487), F1: uint8(v64)}
                                                                                                            var v89 uint = 194
                                                                                                            _ = v88
                                                                                                            _ = v89
                                                                                                        } else {
                                                                                                            v76 = (fi54 / (((fi60 & 2147483647)) | 1))
                                                                                                        }
                                                                                                    }
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        v2 -= float64(v74)
                                                                                                    }
                                                                                                    v81 = v76
                                                                                                } else {
                                                                                                    v90 := GenS8{F0: (-(v4))}
                                                                                                    var v91 string = ((("\n\t" + "\n\t") + "a") + "foo bar")
                                                                                                    if len(v57) > 0 {
                                                                                                        v57[0] = fi60
                                                                                                    }
                                                                                                    v4 = v4
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                }
                                                                                                v4 = (v4 & (v4 * (-(v4))))
                                                                                                if len(v57) >= 2 {
                                                                                                    v57 = v57[1:len(v57)]
                                                                                                }
                                                                                                var v92 int16 = -32768
                                                                                                if len(v57) >= 2 {
                                                                                                    v57 = v57[1:len(v57)]
                                                                                                }
                                                                                                _ = v84
                                                                                                _ = v92
                                                                                            } else {
                                                                                                v57 = append(v57, (fi60 + v81))
                                                                                                switch (v81 * v76) {
                                                                                                case 32767:
                                                                                                    var v93 uint64 = v64
                                                                                                    _ = v93
                                                                                                case 9223372036854775807:
                                                                                                    var v94 uint8 = 2
                                                                                                    _ = v94
                                                                                                case 40:
                                                                                                    v95 := GenS7{F0: 42}
                                                                                                    _ = v95
                                                                                                }
                                                                                                fmt.Println(v74, v1, v63)
                                                                                            }
                                                                                            delete(v3, ("foo bar" + "foo bar"))
                                                                                        }
                                                                                        switch fi60 {
                                                                                        case 32767:
                                                                                            _ = len(v57)
                                                                                        default:
                                                                                            v75 = (v75 | v75)
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                v63 -= 42
                                                                                            case "\n\t":
                                                                                                if false {
                                                                                                    if (true && false) {
                                                                                                        if (false && false) {
                                                                                                            fmt.Println(v81, v1, v1)
                                                                                                            v3["\n\t"] = (2 != v1)
                                                                                                        }
                                                                                                        switch (v81 ^ (-(fi60))) {
                                                                                                        case -614:
                                                                                                            var v96 uint16 = 1
                                                                                                            v2 = -2.25
                                                                                                            _ = v96
                                                                                                        case 127:
                                                                                                            v97 := GenS5{F0: fi54}
                                                                                                            _ = v97
                                                                                                        case 941:
                                                                                                            v1 -= (v1 / ((4294967295) | 1))
                                                                                                        }
                                                                                                        var v98 int8 = 0
                                                                                                        _ = v98
                                                                                                    }
                                                                                                } else {
                                                                                                    var v99 int32 = v4
                                                                                                    v100 := GenS14{F0: (!false), F1: (v1 & (v1 / (((v1 - v1)) | 1)))}
                                                                                                    _ = v100
                                                                                                    _ = v99
                                                                                                }
                                                                                            default:
                                                                                                if (v74 >= 0.0) {
                                                                                                    v2 = (v2 - 1e10)
                                                                                                    switch "a" {
                                                                                                    case "":
                                                                                                        if len(v57) > 0 {
                                                                                                            v57[0] = 65
                                                                                                        }
                                                                                                    default:
                                                                                                        if (!(v75 >= (v75 & int16(v76)))) {
                                                                                                            var v101 float64 = v2
                                                                                                            var v102 int8 = (v79 ^ 127)
                                                                                                            _ = v101
                                                                                                            _ = v102
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                var v103 uint64 = v64
                                                                                                                var v104 uint64 = (v64 & 1)
                                                                                                                if ((v64 / ((v104) | 1)) <= v64) {
                                                                                                                    var v105 int32 = int32(v75)
                                                                                                                    var v106 float32 = v74
                                                                                                                    if ("" <= "\n\t") {
                                                                                                                        fmt.Println(v2, fi54)
                                                                                                                    }
                                                                                                                    _ = v105
                                                                                                                    _ = v106
                                                                                                                } else {
                                                                                                                    var v107 float32 = float32(3.14159)
                                                                                                                    v108 := GenS6{F0: "a", F1: v63, F2: v75}
                                                                                                                    _ = v107
                                                                                                                    _ = v108
                                                                                                                }
                                                                                                                v109 := GenS15{F0: 9223372036854775807, F1: 42, F2: v79}
                                                                                                                v110 := GenS15{F0: 42, F1: 656, F2: (v79 * v79)}
                                                                                                                var v111 float32 = ((float32((v2 * v2)) / 1.5) * 1.5)
                                                                                                                v112 := v3[("a" + "")]
                                                                                                                var v113 int16 = (v75 + (^(v75)))
                                                                                                                _ = v103
                                                                                                                _ = v104
                                                                                                                _ = v109
                                                                                                                _ = v110
                                                                                                                _ = v111
                                                                                                                _ = v112
                                                                                                                _ = v113
                                                                                                            } else {
                                                                                                                if (((v79 ^ v79) >= (v79 & 100)) || true) {
                                                                                                                    if ((v64 / (((^(v64))) | 1)) != (v64 << (((v1 ^ 2)) & 63))) {
                                                                                                                        v76 -= 4294967295
                                                                                                                        v114 := GenS16{F0: "", F1: "\n\t"}
                                                                                                                        if ((-(fi54)) <= 42) {
                                                                                                                            var v115 int = (fi60 & (v76 & v76))
                                                                                                                            v116 := GenS3{F0: genFn0((!true)), F1: v79}
                                                                                                                            _ = v115
                                                                                                                            _ = v116
                                                                                                                        }
                                                                                                                        _ = v114
                                                                                                                    }
                                                                                                                    fmt.Println(v74, v81, v75)
                                                                                                                    _ = len(v57)
                                                                                                                }
                                                                                                                if true {
                                                                                                                    v76 = 396
                                                                                                                    v117 := GenS6{F0: "", F1: 1, F2: (v75 % (((v75 | (v75 + int16(v64)))) | 1))}
                                                                                                                    v79 -= 100
                                                                                                                    v118 := GenS15{F0: -1, F1: int64(fi60), F2: (v79 & v79)}
                                                                                                                    v75 = 0
                                                                                                                    _ = v117
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                            }
                                                                                                            v79 += (v79 % ((v79) | 1))
                                                                                                            v63 += v63
                                                                                                        }
                                                                                                    }
                                                                                                    var v119 uint = (v1 * (v1 - v1))
                                                                                                    _ = v119
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = len(v3)
                                                                                    }
                                                                                case "hello":
                                                                                    v57 = append(v57, fi54)
                                                                                }
                                                                            }
                                                                        }
                                                                        var v120 int8 = 58
                                                                        var v121 int16 = v75
                                                                        _ = v120
                                                                        _ = v121
                                                                        _ = v81
                                                                    }
                                                                    v57 = append(v57, (fi60 & v76))
                                                                    _ = v79
                                                                }
                                                                switch ("" + ("" + ("foo bar" + "foo bar"))) {
                                                                case "a":
                                                                    if (42 >= 9223372036854775807) {
                                                                        if (v64 != v64) {
                                                                            v122 := GenS4{F0: 127}
                                                                            v74 = 3.14159
                                                                            _ = v122
                                                                        } else {
                                                                            fmt.Println(v64, v1)
                                                                            if true {
                                                                                v74 -= float32((3.14159 * 0.0))
                                                                                if false {
                                                                                    var v123 int8 = 1
                                                                                    v124 := GenS17{F0: (!true)}
                                                                                    _ = v123
                                                                                    _ = v124
                                                                                } else {
                                                                                    v125 := GenS18{F0: (fi54 ^ fi60)}
                                                                                    _ = v125
                                                                                }
                                                                            } else {
                                                                                v1 = 1
                                                                                v126 := GenS5{F0: v76}
                                                                                _ = v126
                                                                            }
                                                                            v76 = fi60
                                                                        }
                                                                        v127 := GenS18{F0: fi54}
                                                                        fmt.Println(v2)
                                                                        _ = v127
                                                                    }
                                                                case "foo bar":
                                                                    v64 += v64
                                                                }
                                                                _ = v77
                                                            } else {
                                                                var v128 uint8 = 1
                                                                v129 := v3[(("a" + "foo bar") + "hello")]
                                                                v2 = 0.0
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                            fmt.Println(v74)
                                                            _ = v76
                                                        }
                                                        _ = v75
                                                    } else {
                                                        v74 -= (v74 + -2.25)
                                                        if true {
                                                            switch "a" {
                                                            case "":
                                                                var v130 int64 = 9223372036854775807
                                                                _ = v130
                                                            case "a":
                                                                if (!(v4 == -2147483648)) {
                                                                    var v131 int8 = 16
                                                                    _ = v131
                                                                } else {
                                                                    var v132 int32 = 1
                                                                    _ = v132
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v4, v4)
                                                            default:
                                                                v2 += ((0.1 / v2) / v2)
                                                            }
                                                            _ = len(v57)
                                                        } else {
                                                            fmt.Println(v64, fi60)
                                                        }
                                                        v133 := GenS19{F0: -1, F1: (v1 + (v1 >> ((565) & 63)))}
                                                        _ = v133
                                                    }
                                                    var v134 float32 = 3.14159
                                                    _ = v134
                                                    _ = v66
                                                    _ = v74
                                                } else {
                                                    v2 = ((v2 / ((((1e10 * (float64(0.1) + 3.14159)) - (v2 - ((((-2.25 + float64((-2.25 / -2.25))) / v2) * v2) * float64((float32((v2 - 1e10)) + 0.1))))) - float64((-2.25 / 1e10))) * (((v2 - float64(3.14159)) - v2) * v2))) + -2.25)
                                                }
                                            }
                                            if false {
                                                v135 := GenS20{F0: 703, F1: -610}
                                                switch "hello" {
                                                case "\n\t":
                                                    switch fi54 {
                                                    case -890:
                                                        var v136 uint64 = v64
                                                        _ = v136
                                                    }
                                                }
                                                var v137 uint64 = (v64 - v64)
                                                _ = v135
                                                _ = v137
                                            } else {
                                                if (!(v4 <= (v4 >> ((255) & 63)))) {
                                                    v138 := GenS21{F0: 56, F1: "a"}
                                                    if false {
                                                        var v139 int8 = 100
                                                        _ = v139
                                                    }
                                                    _ = v138
                                                } else {
                                                    var v140 int8 = -16
                                                    v141 := GenS22{F0: (v64 >> ((65) & 63))}
                                                    _ = v140
                                                    _ = v141
                                                }
                                            }
                                        }
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    } else {
                                        _ = len(v57)
                                    }
                                    v142 := v3[("a" + ("\n\t" + "hello"))]
                                    _ = fi60
                                    _ = v142
                                }
                                v1 += (v1 * 0)
                                switch 9223372036854775807 {
                                case -9223372036854775808:
                                    var v143 int64 = -152
                                    _ = v143
                                case 9:
                                    for fi144 := 0; fi144 < 15; fi144++ {
                                        v4 = 0
                                        if (!true) {
                                            var v145 uint = 2
                                            var v146 int16 = -1
                                            if ((fi54 / ((fi144) | 1)) == (^(fi144))) {
                                                _ = len(v3)
                                                var v147 bool = (!(!(3.14159 != v2)))
                                                _ = v147
                                            }
                                            _ = v145
                                            _ = v146
                                        } else {
                                            v148 := v3["\n\t"]
                                            var v149 bool = v148
                                            var v150 bool = v148
                                            switch fi54 {
                                            case 128:
                                                switch "foo bar" {
                                                case "hello":
                                                    v151 := GenS11{F0: 255, F1: 2, F2: 0}
                                                    _ = v151
                                                case "foo bar":
                                                    var v152 int8 = 42
                                                    _ = v152
                                                }
                                                var v153 float64 = v2
                                                var v154 uint = v1
                                                _ = v153
                                                _ = v154
                                            case 32768:
                                                var v155 string = ("a" + "hello")
                                                _ = v155
                                            default:
                                                fmt.Println(fi54, v148, v2)
                                            }
                                            _ = v148
                                            _ = v149
                                            _ = v150
                                        }
                                        _ = fi144
                                    }
                                }
                            }
                            if len(v57) > 0 {
                                v57[0] = (fi54 * fi54)
                            }
                        }
                        _ = v57
                    }
                    v1 += v1
                }
                var v156 string = "a"
                var v157 uint8 = 127
                _ = v156
                _ = v157
            } else {
                if (1 >= 468) {
                    v4 = v4
                    v3[((("foo bar" + "foo bar") + "\n\t") + ("foo bar" + "\n\t"))] = ((v1 >= (v1 + (v1 % ((v1) | 1)))) && false)
                    v158 := GenS6{F0: "\n\t", F1: 981, F2: 32767}
                    _ = v158
                } else {
                    if ((!true) || true) {
                        var v159 string = "hello"
                        switch (v159 + ("" + "\n\t")) {
                        case "\n\t":
                            v160 := []int{(fi54 ^ 32767)}
                            _ = v160
                        case "a":
                            for fi161 := 0; fi161 < 8; fi161++ {
                                var v162 bool = true
                                if (!v162) {
                                    v1 = 4294967295
                                    var v163 float64 = v2
                                    var v164 uint64 = 0
                                    _ = v163
                                    _ = v164
                                } else {
                                    v165 := GenS5{F0: (-(fi54))}
                                    _ = v165
                                }
                                v166 := GenS23{F0: (fi54 % ((9223372036854775807) | 1)), F1: 526}
                                _ = fi161
                                _ = v162
                                _ = v166
                            }
                        case "hello":
                            v159 = v159
                        }
                        var v167 int64 = -9223372036854775808
                        _ = v159
                        _ = v167
                    }
                    var v168 int64 = -9223372036854775808
                    var v169 int8 = -1
                    _ = v168
                    _ = v169
                }
            }
            var v170 bool = (0.0 != v2)
            _ = fi54
            _ = v170
        }
        v171 := []int{4294967295, -1, 4294967295, 0}
        _ = v171
        _ = v53
    }
    for fi172 := 0; fi172 < 15; fi172++ {
        switch "" {
        case "a":
            v173 := []int{fi172, fi172, fi172}
            _ = v173
        case "foo bar":
            var v174 uint = v1
            _ = v174
        case "\n\t":
            v4 = 42
        default:
            v175 := GenS23{F0: fi172, F1: 42}
            _ = v175
        }
        _ = fi172
    }
    v176 := []int{}
    var v177 int16 = genFn0((1e10 == 3.14159))
    switch "" {
    case "\n\t":
        fmt.Println(v177)
    default:
        v3[("hello" + ("" + ("a" + "a")))] = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v176:", len(v176), v176)
    fmt.Println("v177:", v177)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v176
    _ = v177
    _ = v2
    _ = v3
    _ = v4
}
