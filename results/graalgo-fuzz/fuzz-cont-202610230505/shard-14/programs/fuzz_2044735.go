package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 int64; F2 int8 }
type GenS7 struct { F0 int64; F1 int32; F2 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 uint16; F2 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int; F1 int16; F2 uint }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint64; F1 string }
type GenS14 struct { F0 int }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint64; F1 uint64 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS19 struct { F0 uint8; F1 int16; F2 uint }
type GenS20 struct { F0 string; F1 string }
type GenS21 struct { F0 int16; F1 int64 }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 int }
func genFn0() uint32 {
    var v5 int16 = 0
    fmt.Println(v5)
    _ = v5
    return 485
}
func genFn1(p0 uint64) int8 {
    p0 = 18446744073709551615
    switch "" {
    case "foo bar":
        switch ("\n\t" + ("" + (("\n\t" + ("a" + "foo bar")) + "hello"))) {
        case "\n\t":
            v6 := GenS1{F0: 42, F1: 787, F2: (p0 + (p0 / ((p0) | 1)))}
            _ = v6
        }
    case "\n\t":
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint64 = 1
            _ = fi7
            _ = v8
        }
    default:
        if (42 > 0) {
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 uint64 = (^(p0))
                v11 := []int{fi9, (fi9 << ((55) & 63))}
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    if true {
                        v10 = 42
                        fmt.Println(v10)
                    } else {
                        fmt.Println(fi9, v10, v10)
                        var v12 uint64 = (p0 | (v10 >> ((1000) & 63)))
                        var v13 uint64 = (^(v10))
                        _ = v12
                        _ = v13
                    }
                case "a":
                    var v14 float64 = (-2.25 - 0.0)
                    _ = v14
                case "hello":
                    p0 = (p0 % ((1) | 1))
                default:
                    var v15 int64 = -185
                    _ = v15
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            p0 -= 620
        }
    }
    if false {
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 uint16 = 856
                var v19 uint64 = 673
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 float32 = (0.0 + (((1e10 * (3.14159 - float32((1.5 - -2.25)))) + 1e10) / 3.14159))
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(v20, fi21, fi21)
                v22 := []int{(fi21 << ((2) & 63)), -2147483648, (fi21 - 1), (fi16 << ((684) & 63))}
                if true {
                    var v23 uint64 = p0
                    var v24 uint16 = 2
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v20, v20, fi21)
                    var v25 uint = 255
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
            _ = fi16
            _ = v20
        }
        fmt.Println(p0)
    } else {
        v26 := GenS2{F0: "\n\t", F1: -128}
        var v27 uint64 = 2
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: 456, F1: 503}
    _ = p0
    _ = v28
    return 100
}

func main() {
    v1 := []uint32{828, 2147483648, 4294967295}
    var v2 int8 = -94
    var v3 uint16 = 310
    v4 := []int{}
    v3 -= (v3 / (((v3 & v3)) | 1))
    var v29 int8 = v2
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 7; fi31++ {
            v1 = append(v1, 0)
            if true {
                v32 := GenS4{F0: 0, F1: 0, F2: true}
                var v33 float32 = 1.5
                fmt.Println(v29, v33, v33)
                _ = v32
                _ = v33
            }
            _ = fi31
        }
        _ = fi30
    }
    v29 += v2
    v2 = (v29 * genFn1(18446744073709551615))
    v1 = append(v1, 2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 += (v3 >> ((4294967295) & 63))
    var v34 bool = true
    if len(v4) > 0 {
        v4[0] = 105
    }
    var v35 uint32 = 65535
    var v36 float32 = 0.0
    var v37 int8 = genFn1(0)
    for fi38 := 0; fi38 < 8; fi38++ {
        fmt.Println(v35, v2)
        var v39 uint32 = 4294967295
        v36 += v36
        _ = fi38
        _ = v39
    }
    if (376 >= 42) {
        if true {
            v40 := GenS4{F0: v35, F1: 663, F2: true}
            _ = v40
        } else {
            if (v34 && (v34 || v34)) {
                fmt.Println(v36, v2, v34)
                if (v34 && v34) {
                    v29 = 127
                } else {
                    fmt.Println(v35)
                    v34 = v34
                }
                var v41 float32 = v36
                _ = v41
            } else {
                v42 := GenS3{F0: -1, F1: -914}
                if v34 {
                    var v43 uint64 = 1
                    fmt.Println(v36, v2, v35)
                    _ = v43
                } else {
                    var v44 float32 = 0.1
                    switch "hello" {
                    case "":
                        v35 -= 2
                    case "hello":
                        var v45 uint8 = 179
                        _ = v45
                    case "\n\t":
                        var v46 float64 = 1e10
                        _ = v46
                    }
                    v47 := GenS4{F0: genFn0(), F1: 46, F2: (!(1.5 <= 0.1))}
                    _ = v44
                    _ = v47
                }
                for fi48 := 0; fi48 < 4; fi48++ {
                    if (("foo bar" + "hello") >= "") {
                        var v49 bool = v34
                        _ = v49
                    } else {
                        v50 := GenS5{F0: v3}
                        _ = v50
                    }
                    _ = fi48
                }
                _ = v42
            }
            switch "\n\t" {
            case "\n\t":
                v51 := GenS1{F0: v3, F1: 1, F2: 0}
                _ = v51
            case "hello":
                if (1e10 == 0.1) {
                    v52 := GenS6{F0: (("foo bar" + "\n\t") + "\n\t"), F1: 50, F2: (v29 | (v37 * v37))}
                    _ = v52
                } else {
                    v53 := GenS4{F0: (^(v35)), F1: 0, F2: v34}
                    for fi54 := 0; fi54 < 10; fi54++ {
                        switch "foo bar" {
                        case "hello":
                            if v34 {
                                for fi55 := 0; fi55 < 4; fi55++ {
                                    if (1e10 < float64(v36)) {
                                        v2 -= (v29 & v2)
                                        v4 = append(v4, 127)
                                        switch ("" + "foo bar") {
                                        case "":
                                            v35 = v35
                                        }
                                    }
                                    v56 := GenS7{F0: -461, F1: 774, F2: 624}
                                    if (628 > 385) {
                                        var v57 float32 = float32((((-2.25 - -2.25) * 0.0) - 3.14159))
                                        _ = v57
                                    } else {
                                        switch -616 {
                                        case 32767:
                                            switch (fi54 / ((-835) | 1)) {
                                            case 4294967295:
                                                switch "" {
                                                case "":
                                                    v35 -= (^(v35))
                                                case "hello":
                                                    var v58 float32 = v36
                                                    _ = v58
                                                default:
                                                    v59 := GenS1{F0: 32767, F1: genFn0(), F2: 0}
                                                    v60 := GenS8{F0: 0}
                                                    fmt.Println(fi54, fi55)
                                                    if v34 {
                                                        if (1 > -1) {
                                                            v61 := GenS2{F0: "\n\t", F1: genFn1(0)}
                                                            var v62 int = -658
                                                            v2 = -128
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            switch (fi54 >> ((520) & 63)) {
                                                            case -872:
                                                                if v34 {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v35 += v35
                                                                    case "foo bar":
                                                                        if (v3 > uint16(v35)) {
                                                                            fmt.Println(v34)
                                                                            switch (fi55 | fi55) {
                                                                            case 42:
                                                                                _ = len(v4)
                                                                            case 1:
                                                                                v37 += v2
                                                                            default:
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    fmt.Println(v3, fi55, v34)
                                                                                case "foo bar":
                                                                                    fmt.Println(v29, v37)
                                                                                }
                                                                            }
                                                                            v63 := GenS9{F0: 128, F1: (v3 - 1), F2: ("hello" + ("foo bar" + "\n\t"))}
                                                                            _ = v63
                                                                        }
                                                                    default:
                                                                        if (!(!false)) {
                                                                            switch fi55 {
                                                                            case 32768:
                                                                                if (!v34) {
                                                                                    if (!((1.5 <= (1.5 - 3.14159)) && ((v34 && (!false)) && v34))) {
                                                                                        var v64 uint16 = (v3 * 42)
                                                                                        v65 := GenS10{F0: 2}
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    } else {
                                                                                        var v66 uint32 = v35
                                                                                        if v34 {
                                                                                            _ = len(v4)
                                                                                            v3 -= 2
                                                                                        }
                                                                                        v67 := GenS10{F0: 2}
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                    v36 -= v36
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v2 -= genFn1(18)
                                                                    v34 = ((v35 % ((v35) | 1)) < genFn0())
                                                                } else {
                                                                    if (("" + "a") >= "a") {
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            v68 := GenS9{F0: 130, F1: v3, F2: "a"}
                                                                            _ = v68
                                                                        case "":
                                                                            v36 += (float32(float64(v36)) * (0.0 / v36))
                                                                        case "hello":
                                                                            v29 = (v29 & genFn1(2))
                                                                        default:
                                                                            fmt.Println(v2)
                                                                        }
                                                                        var v69 float32 = v36
                                                                        fmt.Println(v36, v37, v35)
                                                                        fmt.Println(v29)
                                                                        _ = v69
                                                                    } else {
                                                                        switch ("foo bar" + ("" + "hello")) {
                                                                        case "\n\t":
                                                                            v37 = v29
                                                                        case "foo bar":
                                                                            v70 := GenS11{F0: 694, F1: -32768, F2: 65535}
                                                                            _ = v70
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 4294967295
                                                                        }
                                                                        var v71 string = ("foo bar" + "foo bar")
                                                                        _ = v71
                                                                    }
                                                                }
                                                            case 660:
                                                                v37 = (v29 >> ((255) & 63))
                                                            default:
                                                                v35 = genFn0()
                                                            }
                                                            v36 -= 0.0
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                if (v34 && v34) {
                                                                    v72 := GenS12{F0: 255}
                                                                    if (!(((1e10 * 1e10) / (float64(v36) + (1.5 + (0.1 * 3.14159)))) < 0.0)) {
                                                                        v73 := GenS4{F0: v35, F1: 74, F2: (("a" + ("\n\t" + "")) == "hello")}
                                                                        _ = v73
                                                                    } else {
                                                                        v74 := GenS13{F0: 864, F1: ("hello" + "\n\t")}
                                                                        if false {
                                                                            var v75 float32 = (1e10 + 1.5)
                                                                            _ = v75
                                                                        } else {
                                                                            v76 := GenS14{F0: 42}
                                                                            v2 -= (v2 ^ genFn1(42))
                                                                            _ = v76
                                                                        }
                                                                        _ = v74
                                                                    }
                                                                    _ = v72
                                                                }
                                                            case "a":
                                                                if (v36 <= 1e10) {
                                                                    v77 := GenS15{F0: false}
                                                                    fmt.Println(v29, v29, v34)
                                                                    var v78 float64 = 0.0
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                            }
                                                            var v79 bool = true
                                                            _ = v79
                                                        }
                                                        var v80 int32 = 1
                                                        v1 = append(v1, (v35 / ((548) | 1)))
                                                        if true {
                                                            if (0 > 632) {
                                                                var v81 int8 = (v37 ^ -93)
                                                                v82 := GenS16{F0: 4294967295, F1: 42}
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                            v83 := GenS2{F0: ("\n\t" + ""), F1: 42}
                                                            _ = v83
                                                        }
                                                        _ = v80
                                                    } else {
                                                        var v84 uint64 = 489
                                                        if false {
                                                            v84 += (v84 & v84)
                                                            fmt.Println(v35, v3)
                                                            var v85 uint = 919
                                                            if (int(v85) > fi54) {
                                                                if ((((v84 << ((948) & 63)) == (v84 + v84)) && (v34 && ("hello" != "hello"))) || (42 < -352)) {
                                                                    v34 = (v34 && v34)
                                                                } else {
                                                                    v86 := GenS17{F0: uint8(v85)}
                                                                    v87 := GenS18{F0: (^(v2)), F1: (^(v3)), F2: (v35 >> ((v85) & 63))}
                                                                    switch (fi54 / ((int(v35)) | 1)) {
                                                                    case 32767:
                                                                        var v88 float64 = 1.5
                                                                        _ = v88
                                                                    default:
                                                                        var v89 bool = v34
                                                                        _ = v89
                                                                    }
                                                                    _ = v86
                                                                    _ = v87
                                                                }
                                                            }
                                                            _ = v85
                                                        } else {
                                                            if false {
                                                                fmt.Println(v35, fi55)
                                                            } else {
                                                                v90 := GenS19{F0: 0, F1: -32768, F2: 255}
                                                                _ = v90
                                                            }
                                                        }
                                                        _ = v84
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = v35
                                                    }
                                                    _ = v59
                                                    _ = v60
                                                }
                                            }
                                        case 42:
                                            var v91 uint64 = 18446744073709551615
                                            v91 = 0
                                            if (53 == 793) {
                                                fmt.Println(v29, v37)
                                                var v92 int32 = -302
                                                _ = v92
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = v35
                                                }
                                            }
                                            var v93 uint32 = uint32(v91)
                                            _ = v91
                                            _ = v93
                                        case 4294967295:
                                            v1 = append(v1, (v35 >> ((4294967295) & 63)))
                                        }
                                        fmt.Println(fi55)
                                    }
                                    if v34 {
                                        v94 := GenS4{F0: (v35 | (v35 << ((2) & 63))), F1: 42, F2: true}
                                        fmt.Println(fi54)
                                        _ = v94
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                            } else {
                                var v95 uint16 = v3
                                var v96 uint64 = 4294967295
                                fmt.Println(v2)
                                _ = v95
                                _ = v96
                            }
                        case "\n\t":
                            if (!(((v36 * float32(-2.25)) + v36) <= v36)) {
                                v37 += (v29 << ((2) & 63))
                            } else {
                                if true {
                                    var v97 float64 = 1e10
                                    _ = v97
                                }
                            }
                        case "a":
                            var v98 int = 1
                            _ = v98
                        }
                        _ = fi54
                    }
                    if len(v4) > 0 {
                        v4[0] = 128
                    }
                    _ = v53
                }
            case "a":
                if v34 {
                    v29 -= v2
                    fmt.Println(v37, v2)
                }
            default:
                v3 = (v3 * (v3 * 65535))
            }
        }
        fmt.Println(v29)
        switch 318 {
        case 128:
            if (-32768 <= -32768) {
                fmt.Println(v35, v35, v37)
                if (1 == -2147483648) {
                    v29 += v37
                }
                if (v36 == (v36 * float32(0.1))) {
                    v99 := GenS20{F0: ("hello" + ("" + "foo bar")), F1: ("\n\t" + "a")}
                    if true {
                        v35 -= v35
                        v100 := GenS21{F0: 1000, F1: 0}
                        v101 := GenS16{F0: 42, F1: 18446744073709551615}
                        _ = v100
                        _ = v101
                    }
                    for fi102 := 0; fi102 < 20; fi102++ {
                        for fi103 := 0; fi103 < 5; fi103++ {
                            var v104 uint64 = 427
                            switch fi102 {
                            case 9223372036854775807:
                                if v34 {
                                    v2 -= v37
                                    if (0 >= 34) {
                                        fmt.Println(fi102)
                                        switch (fi102 & (^(fi103))) {
                                        case -2147483648:
                                            var v105 int64 = 42
                                            _ = v105
                                        case -457:
                                            if true {
                                                v3 = (v3 / ((2) | 1))
                                                var v106 uint64 = v104
                                                _ = v106
                                            }
                                        }
                                        var v107 uint64 = v104
                                        v36 -= float32((float64(v36) - (1e10 * ((1e10 + (-2.25 + 0.1)) / 0.1))))
                                        _ = v107
                                    }
                                } else {
                                    var v108 int = (fi103 << ((504) & 63))
                                    _ = v108
                                }
                            case -72:
                                v29 -= genFn1((v104 & uint64(v35)))
                            case -9223372036854775808:
                                v104 -= v104
                            default:
                                v3 += 1
                            }
                            switch fi102 {
                            case 114:
                                var v109 float32 = v36
                                v36 = v109
                                v110 := GenS22{F0: (v3 + (v3 % (((^(v3))) | 1)))}
                                _ = v109
                                _ = v110
                            default:
                                if v34 {
                                    var v111 uint = 630
                                    fmt.Println(v104)
                                    v111 += (v111 / ((v111) | 1))
                                    v37 = 42
                                    _ = v111
                                }
                                if v34 {
                                    switch -9223372036854775808 {
                                    case 2147483647:
                                        if v34 {
                                            switch -63 {
                                            case -100:
                                                if v34 {
                                                    var v112 int16 = -328
                                                    _ = v112
                                                } else {
                                                    if (v34 && false) {
                                                        if true {
                                                            v29 = (v37 + genFn1(2))
                                                            v3 -= 568
                                                        }
                                                        v2 = (v2 * genFn1(18446744073709551615))
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = 2147483647
                                                        }
                                                        if (v34 && (!(int8(fi102) == (v2 + 100)))) {
                                                            fmt.Println(v37, v104)
                                                            v104 += 1
                                                            var v113 uint32 = v35
                                                            _ = v113
                                                        } else {
                                                            v104 += (v104 / ((v104) | 1))
                                                            switch ("\n\t" + "") {
                                                            case "\n\t":
                                                                v36 = (v36 - v36)
                                                            default:
                                                                _ = len(v4)
                                                            }
                                                        }
                                                        var v114 int16 = -32768
                                                        _ = v114
                                                    }
                                                    var v115 int = (fi102 << ((42) & 63))
                                                    var v116 string = ("hello" + "hello")
                                                    _ = v115
                                                    _ = v116
                                                }
                                            case 834:
                                                fmt.Println(v37)
                                            default:
                                                var v117 float64 = 0.0
                                                _ = v117
                                            }
                                            v36 = ((1.5 / v36) - 3.14159)
                                        } else {
                                            v2 = (v37 ^ 42)
                                            if true {
                                                v2 = (v2 ^ (-(v2)))
                                                v118 := GenS23{F0: (fi103 | fi102)}
                                                _ = v118
                                            } else {
                                                v36 = v36
                                            }
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v36)
                                    case -9223372036854775808:
                                        v104 -= (v104 + (^(v104)))
                                    default:
                                        fmt.Println(fi103)
                                    }
                                    var v119 uint = 724
                                    _ = v119
                                }
                            }
                            _ = fi103
                            _ = v104
                        }
                        _ = fi102
                    }
                    _ = v99
                } else {
                    v34 = v34
                    v37 = v2
                }
            } else {
                var v120 int8 = -14
                if ((!(v36 > v36)) || (v35 >= 946)) {
                    if ((122 == 393) && ("foo bar" > "\n\t")) {
                        for fi121 := 0; fi121 < 2; fi121++ {
                            for fi122 := 0; fi122 < 18; fi122++ {
                                switch "" {
                                case "a":
                                    fmt.Println(v34, v35, v34)
                                }
                                fmt.Println(v35, v29)
                                v36 = v36
                                _ = fi122
                            }
                            v2 += v2
                            _ = fi121
                        }
                        fmt.Println(v3)
                    }
                } else {
                    if false {
                        var v123 int16 = -641
                        if (("a" + "") <= "a") {
                            var v124 int16 = (v123 % (((v123 >> ((18446744073709551615) & 63))) | 1))
                            _ = v124
                        }
                        v29 += v29
                        _ = v123
                    }
                }
                v125 := GenS17{F0: 255}
                _ = v120
                _ = v125
            }
        case 142:
            if false {
                fmt.Println(v2)
                var v126 int = -95
                _ = v126
            } else {
                for fi127 := 0; fi127 < 10; fi127++ {
                    for fi128 := 0; fi128 < 5; fi128++ {
                        var v129 bool = false
                        _ = fi128
                        _ = v129
                    }
                    var v130 string = ((("\n\t" + (("a" + "\n\t") + "")) + "\n\t") + "hello")
                    var v131 int16 = -27
                    _ = fi127
                    _ = v130
                    _ = v131
                }
                v132 := GenS8{F0: 0}
                _ = v132
            }
        }
    }
    var v133 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v133
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
