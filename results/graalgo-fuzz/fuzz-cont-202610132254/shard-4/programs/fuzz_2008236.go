package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int32; F2 uint64 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32; F1 int32; F2 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8; F1 int16; F2 int8 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint64; F1 uint8 }
type GenS16 struct { F0 int16; F1 uint8; F2 uint }
type GenS17 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 uint32; F1 uint16 }
type GenS20 struct { F0 uint64; F1 int; F2 uint8 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 uint64; F1 bool; F2 int8 }
type GenS25 struct { F0 int16; F1 uint32; F2 bool }
type GenS26 struct { F0 uint8 }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 int }
type GenS29 struct { F0 bool; F1 int64; F2 int32 }
type GenS30 struct { F0 uint8; F1 uint }
func genFn0(p0 uint8) bool {
    if false {
        v6 := GenS1{F0: 128, F1: (p0 >> ((4294967295) & 63)), F2: (-1 == 1)}
        _ = v6
    } else {
        switch -2147483648 {
        case 1:
            if (true || ((!(!false)) || false)) {
                switch "hello" {
                case "":
                    switch -9223372036854775808 {
                    case 42:
                        if (!(-144 == -177)) {
                            var v7 float32 = 0.1
                            _ = v7
                        } else {
                            v8 := []int{0, 0}
                            _ = v8
                        }
                    default:
                        v9 := GenS1{F0: p0, F1: 0, F2: false}
                        _ = v9
                    }
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        var v10 int8 = 1
                        _ = v10
                    default:
                        v11 := []int{}
                        _ = v11
                    }
                case "hello":
                    switch 0 {
                    case -715:
                        for fi12 := 0; fi12 < 13; fi12++ {
                            var v13 uint32 = 65535
                            _ = fi12
                            _ = v13
                        }
                    case 128:
                        var v14 uint16 = 65535
                        _ = v14
                    default:
                        for fi15 := 0; fi15 < 4; fi15++ {
                            var v16 bool = false
                            for fi17 := 0; fi17 < 7; fi17++ {
                                var v18 bool = (3.14159 >= 1e10)
                                if (false && v18) {
                                    fmt.Println(p0, v16, v16)
                                }
                                _ = fi17
                                _ = v18
                            }
                            _ = fi15
                            _ = v16
                        }
                    }
                default:
                    var v19 uint32 = 2147483647
                    _ = v19
                }
            }
        case -9223372036854775808:
            p0 = p0
        case -2147483648:
            var v20 uint32 = 0
            _ = v20
        }
    }
    v21 := GenS1{F0: 0, F1: 127, F2: true}
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS1{F0: 3, F1: (p0 + (p0 * p0)), F2: true}
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = v21
    return (0 > 32768)
}
func genFn1(p0 uint32, p1 int16) uint8 {
    for fi24 := 0; fi24 < 1; fi24++ {
        v25 := GenS1{F0: 0, F1: 127, F2: ((fi24 & -2147483648) <= -361)}
        var v26 uint32 = (p0 + p0)
        v27 := []int{}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 int8 = -128
    var v29 uint16 = 993
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    return 255
}

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    var v3 string = ""
    var v4 bool = false
    v5 := []string{"hello", ("" + v3), "", ("a" + "hello")}
    if len(v5) > 0 {
        v5[0] = (v3 + v3)
    }
    if len(v5) > 0 {
        v5[0] = ""
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        fmt.Println(v1)
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 uint64 = 42
            v32 -= v32
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        switch ("\n\t" + "") {
        case "hello":
            v34 := GenS1{F0: 127, F1: 255, F2: (853 >= 1)}
            _ = v34
        default:
            var v35 int32 = -341
            _ = v35
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            v1 = v1
            _ = fi36
        }
        _ = fi33
    }
    if len(v5) > 0 {
        v5[0] = ("\n\t" + v3)
    }
    v37 := GenS1{F0: 0, F1: 128, F2: (!((1.5 * ((0.0 * 1.5) / 0.1)) != 3.14159))}
    var v38 uint = 383
    var v39 uint32 = 4294967295
    var v40 int16 = -32768
    switch (v3 + "hello") {
    case "hello":
        if false {
            if v4 {
                if true {
                    v41 := GenS1{F0: uint8(v38), F1: 2, F2: (true && false)}
                    for fi42 := 0; fi42 < 15; fi42++ {
                        switch ("\n\t" + ((v3 + v3) + (v3 + "\n\t"))) {
                        case "":
                            if (!false) {
                                v43 := GenS1{F0: 42, F1: 126, F2: (!false)}
                                _ = v43
                            } else {
                                switch (v3 + v3) {
                                case "foo bar":
                                    v44 := GenS2{F0: -1, F1: -701}
                                    _ = v44
                                }
                                for fi45 := 0; fi45 < 11; fi45++ {
                                    var v46 uint8 = 2
                                    v1 = -128
                                    fmt.Println(v1, v40, v3)
                                    fmt.Println(fi42, v46)
                                    _ = fi45
                                    _ = v46
                                }
                                if (!v4) {
                                    v4 = ((v40 | v40) < -1)
                                } else {
                                    v47 := GenS2{F0: -628, F1: 10}
                                    _ = v47
                                }
                            }
                        case "a":
                            switch (fi42 - fi42) {
                            case -1:
                                if false {
                                    if (v4 && false) {
                                        for fi48 := 0; fi48 < 20; fi48++ {
                                            v38 += (^(v38))
                                            _ = fi48
                                        }
                                        if v4 {
                                            if false {
                                                v49 := v2[((v3 + "a") + v3)]
                                                _ = v49
                                            }
                                            v3 = ("a" + v3)
                                        } else {
                                            var v50 string = "\n\t"
                                            switch fi42 {
                                            case 42:
                                                v38 = uint(v39)
                                            }
                                            for fi51 := 0; fi51 < 15; fi51++ {
                                                v40 -= -1
                                                v4 = true
                                                var v52 uint8 = 212
                                                _ = fi51
                                                _ = v52
                                            }
                                            _ = v50
                                        }
                                    } else {
                                        var v53 int8 = v1
                                        switch v3 {
                                        case "":
                                            delete(v2, v3)
                                        case "foo bar":
                                            for fi54 := 0; fi54 < 1; fi54++ {
                                                v4 = (uint(v39) != (v38 / ((v38) | 1)))
                                                var v55 float32 = -2.25
                                                _ = fi54
                                                _ = v55
                                            }
                                        case "hello":
                                            var v56 int16 = 809
                                            _ = v56
                                        }
                                        _ = v53
                                    }
                                    if (((0 != 879) || true) || (2147483647 < -409)) {
                                        var v57 int8 = int8(v40)
                                        var v58 int8 = v57
                                        var v59 int32 = 42
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v60 float64 = float64(-2.25)
                                        if v4 {
                                            if v4 {
                                                var v61 int8 = v1
                                                for fi62 := 0; fi62 < 5; fi62++ {
                                                    v63 := GenS1{F0: 89, F1: genFn1(v39, -32768), F2: (uint64(v39) == 0)}
                                                    _ = len(v5)
                                                    _ = fi62
                                                    _ = v63
                                                }
                                                _ = v61
                                            }
                                            for fi64 := 0; fi64 < 4; fi64++ {
                                                if v4 {
                                                    switch v3 {
                                                    case "":
                                                        fmt.Println(v38)
                                                    case "foo bar":
                                                        var v65 uint64 = 372
                                                        _ = v65
                                                    }
                                                } else {
                                                    var v66 float32 = float32(v60)
                                                    _ = v66
                                                }
                                                v1 -= 127
                                                var v67 uint = 111
                                                _ = fi64
                                                _ = v67
                                            }
                                        } else {
                                            if ((v38 % (((^(v38))) | 1)) >= v38) {
                                                v5 = append(v5, v3)
                                                if len(v5) > 0 {
                                                    v5[0] = "hello"
                                                }
                                            } else {
                                                var v68 int8 = (^(v1))
                                                v4 = true
                                                v40 += (v40 & 1)
                                                _ = v68
                                            }
                                        }
                                        _ = v60
                                    }
                                }
                            case 128:
                                switch fi42 {
                                case -251:
                                    if v4 {
                                        v39 = 206
                                        var v69 uint32 = 2147483647
                                        var v70 int = 4294967295
                                        _ = v69
                                        _ = v70
                                    } else {
                                        v1 -= 14
                                        v39 -= v39
                                    }
                                default:
                                    var v71 int16 = (v40 & v40)
                                    _ = v71
                                }
                            default:
                                var v72 float64 = (1.5 / (0.1 / (1.5 + -2.25)))
                                _ = v72
                            }
                        default:
                            for fi73 := 0; fi73 < 14; fi73++ {
                                v74 := GenS1{F0: genFn1((^(v39)), (^(v40))), F1: 138, F2: true}
                                if false {
                                    fmt.Println(fi73, fi73)
                                    v40 = v40
                                    if (42 > (fi42 ^ fi73)) {
                                        if true {
                                            var v75 int32 = -1
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v75
                                        }
                                        switch v3 {
                                        case "foo bar":
                                            var v76 float32 = (0.0 - float32(3.14159))
                                            _ = v76
                                        case "\n\t":
                                            v77 := v2[("" + "foo bar")]
                                            fmt.Println(v3)
                                            switch (("foo bar" + (v3 + v3)) + v3) {
                                            case "\n\t":
                                                v40 = (v40 >> (((^(v38))) & 63))
                                            case "hello":
                                                fmt.Println(fi73)
                                            }
                                            fmt.Println(v1, v39)
                                            if false {
                                                if v4 {
                                                    switch v3 {
                                                    case "foo bar":
                                                        fmt.Println(fi42, v1)
                                                    case "\n\t":
                                                        _ = len(v2)
                                                        v4 = (v39 != 1)
                                                        v78 := GenS3{F0: v38, F1: v40}
                                                        _ = v78
                                                    case "":
                                                        switch (v77 / (((fi42 + fi73)) | 1)) {
                                                        case -9223372036854775808:
                                                            var v79 uint16 = 407
                                                            _ = v79
                                                        default:
                                                            if len(v5) > 0 {
                                                                v5[0] = v3
                                                            }
                                                        }
                                                    }
                                                    v1 -= 100
                                                    v80 := GenS1{F0: 135, F1: 123, F2: v4}
                                                    switch "" {
                                                    case "a":
                                                        v81 := GenS3{F0: (v38 - (v38 * 65535)), F1: (v40 / ((-606) | 1))}
                                                        _ = v81
                                                    case "\n\t":
                                                        switch v3 {
                                                        case "\n\t":
                                                            v3 = ("" + v3)
                                                        default:
                                                            v4 = v4
                                                        }
                                                    case "foo bar":
                                                        var v82 uint64 = 2
                                                        _ = v82
                                                    default:
                                                        if (-2.25 == (((0.0 - 0.1) * float32(-2.25)) + 0.0)) {
                                                            v83 := GenS1{F0: 128, F1: 128, F2: (uint8(v38) < 127)}
                                                            _ = v83
                                                        }
                                                        switch int(v39) {
                                                        case 42:
                                                            v1 = (v1 << ((v38) & 63))
                                                        case 32767:
                                                            v2["hello"] = fi73
                                                        default:
                                                            var v84 int32 = 2147483647
                                                            _ = v84
                                                        }
                                                    }
                                                    v39 = v39
                                                    var v85 float64 = 3.14159
                                                    var v86 uint64 = 835
                                                    _ = v80
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    var v87 uint64 = 749
                                                    fmt.Println(v40, v3, v39)
                                                    v39 = 0
                                                    _ = v87
                                                }
                                            }
                                            _ = v77
                                        default:
                                            switch v3 {
                                            case "hello":
                                                var v88 string = v3
                                                _ = v88
                                            case "foo bar":
                                                if v4 {
                                                    if ((v38 ^ 18446744073709551615) < (v38 + 65535)) {
                                                        v39 += (^(v39))
                                                        v3 = (("" + v3) + "foo bar")
                                                        if (127 == 0) {
                                                            v39 = (v39 / (((v39 ^ v39)) | 1))
                                                        } else {
                                                            v89 := GenS4{F0: (^(v1))}
                                                            if (6 != 2) {
                                                                var v90 uint8 = 1
                                                                _ = v90
                                                            }
                                                            var v91 uint16 = 0
                                                            _ = v89
                                                            _ = v91
                                                        }
                                                        var v92 uint8 = 255
                                                        if ((1.5 / 0.1) == (0.0 * (1e10 / 1e10))) {
                                                            fmt.Println(fi73)
                                                        }
                                                        v93 := GenS2{F0: -9223372036854775808, F1: -2147483648}
                                                        var v94 int16 = (v40 >> ((65535) & 63))
                                                        _ = v92
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        v95 := GenS5{F0: (v38 >> ((2) & 63))}
                                                        v39 = (v39 * v39)
                                                        _ = v95
                                                    }
                                                }
                                                v3 = (v3 + ("foo bar" + (v3 + (((v3 + v3) + v3) + ""))))
                                                v39 = v39
                                            }
                                        }
                                        var v96 uint8 = 2
                                        _ = v96
                                    }
                                } else {
                                    if v4 {
                                        v97 := GenS6{F0: v3, F1: -147, F2: 42}
                                        v2[((((((((((v3 + ((("hello" + v3) + (v3 + "foo bar")) + v3)) + v3) + v3) + "foo bar") + v3) + v3) + ("\n\t" + v3)) + v3) + "hello") + v3)] = (fi42 / (((-(fi73))) | 1))
                                        delete(v2, "hello")
                                        if len(v5) > 0 {
                                            v5[0] = v3
                                        }
                                        if v4 {
                                            switch fi73 {
                                            case 1:
                                                v39 = v39
                                            case -9223372036854775808:
                                                v98 := GenS7{F0: v40, F1: -729}
                                                _ = v98
                                            case -694:
                                                switch -839 {
                                                case -1:
                                                    var v99 float32 = (1e10 * float32(3.14159))
                                                    var v100 float64 = 0.1
                                                    v101 := GenS7{F0: (v40 * v40), F1: -32768}
                                                    if v4 {
                                                        fmt.Println(v100)
                                                        if v4 {
                                                            v2["foo bar"] = 128
                                                            v40 += v40
                                                        }
                                                        var v102 float32 = v99
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v102
                                                    }
                                                    delete(v2, "a")
                                                    if ((v4 && (v4 || (42 <= v38))) || false) {
                                                        if v4 {
                                                            if len(v5) > 0 {
                                                                v5[0] = v3
                                                            }
                                                        } else {
                                                            v103 := GenS8{F0: (0 != -169), F1: 2147483647}
                                                            if (v99 >= (float32(((v100 * v100) + 3.14159)) * (v99 - (((v99 * v99) * float32(((v100 / (v100 / 1.5)) - float64(0.0)))) / (v99 * (1.5 - ((float32((v100 + (v100 * (1e10 - ((-2.25 / (v100 - (3.14159 - 1e10))) / (v100 / (((1e10 + v100) / v100) / (((float64(v99) + 0.0) / -2.25) / v100)))))))) - v99) / v99))))))) {
                                                                v104 := v2[v3]
                                                                _ = len(v2)
                                                                _ = v104
                                                            } else {
                                                                if ((fi73 ^ (fi42 >> (((^(v38))) & 63))) <= (fi42 * (fi73 - 127))) {
                                                                    var v105 int32 = 1
                                                                    _ = v105
                                                                }
                                                                var v106 float64 = v100
                                                                var v107 int = int(v39)
                                                                switch (v107 / (((fi42 | 4294967295)) | 1)) {
                                                                case -140:
                                                                    if (true && v4) {
                                                                        v1 += (v1 & v1)
                                                                        v106 -= float64(0.0)
                                                                        v108 := GenS9{F0: (^(v107))}
                                                                        _ = v108
                                                                    }
                                                                default:
                                                                    v109 := GenS10{F0: int32(fi73), F1: 42, F2: 2}
                                                                    _ = v109
                                                                }
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                            _ = v103
                                                        }
                                                        switch v3 {
                                                        case "":
                                                            switch fi42 {
                                                            case 9223372036854775807:
                                                                v1 += 1
                                                            case 466:
                                                                v110 := v2[(v3 + "foo bar")]
                                                                _ = v110
                                                            case 99:
                                                                var v111 uint = (v38 >> (((v38 / ((v38) | 1))) & 63))
                                                                _ = v111
                                                            default:
                                                                fmt.Println(v99, v100, v38)
                                                            }
                                                        case "\n\t":
                                                            v38 = (v38 * v38)
                                                        }
                                                        switch (^(fi42)) {
                                                        case 32768:
                                                            var v112 int16 = v40
                                                            _ = v112
                                                        }
                                                    }
                                                    _ = v100
                                                    _ = v101
                                                    _ = v99
                                                case -225:
                                                    v113 := GenS11{F0: (v38 * v38)}
                                                    _ = v113
                                                default:
                                                    switch -376 {
                                                    case 681:
                                                        switch v3 {
                                                        case "":
                                                            switch fi42 {
                                                            case 27:
                                                                fmt.Println(v1, v39, v4)
                                                            case 0:
                                                                v38 = v38
                                                            default:
                                                                v38 = v38
                                                            }
                                                        case "hello":
                                                            switch (fi73 << ((65535) & 63)) {
                                                            case 127:
                                                                fmt.Println(v1, v38, v4)
                                                            case 4294967295:
                                                                delete(v2, v3)
                                                            default:
                                                                switch v3 {
                                                                case "\n\t":
                                                                    v114 := v2[v3]
                                                                    _ = v114
                                                                case "foo bar":
                                                                    if v4 {
                                                                        fmt.Println(v1, fi42)
                                                                    }
                                                                case "a":
                                                                    var v115 uint32 = (v39 << ((42) & 63))
                                                                    _ = v115
                                                                default:
                                                                    if v4 {
                                                                        v2["\n\t"] = fi73
                                                                        var v116 uint32 = v39
                                                                        v40 -= v40
                                                                        _ = v116
                                                                    } else {
                                                                        fmt.Println(fi42)
                                                                        var v117 uint16 = 1
                                                                        fmt.Println(v40)
                                                                        _ = v117
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            switch fi42 {
                                                            case -853:
                                                                v118 := GenS12{F0: (v1 % ((v1) | 1)), F1: (v40 | (v40 << ((466) & 63))), F2: v1}
                                                                _ = v118
                                                            case 4294967295:
                                                                if v4 {
                                                                    v38 = (v38 / (((v38 << ((v38) & 63))) | 1))
                                                                    switch -1 {
                                                                    case 1:
                                                                        v4 = (255 < 17)
                                                                    case 4294967295:
                                                                        if v4 {
                                                                            v38 -= 2
                                                                            fmt.Println(v38, v3)
                                                                            switch 626 {
                                                                            case 42:
                                                                                var v119 uint16 = 42
                                                                                _ = v119
                                                                            case -1:
                                                                                var v120 uint16 = 497
                                                                                _ = v120
                                                                            default:
                                                                                v121 := GenS13{F0: 2}
                                                                                _ = v121
                                                                            }
                                                                        }
                                                                    case 127:
                                                                        fmt.Println(v4, v4)
                                                                    default:
                                                                        v122 := GenS14{F0: 0}
                                                                        _ = v122
                                                                    }
                                                                    var v123 uint = v38
                                                                    _ = v123
                                                                }
                                                            default:
                                                                v39 -= (v39 - 2)
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        if true {
                                                            v1 = (v1 / ((v1) | 1))
                                                            if v4 {
                                                                v40 = v40
                                                                var v124 uint = uint(v39)
                                                                if (0 == 1) {
                                                                    fmt.Println(fi73, v3)
                                                                    v3 = (v3 + (v3 + v3))
                                                                }
                                                                _ = v124
                                                            } else {
                                                                v125 := GenS12{F0: 127, F1: (v40 + -32768), F2: v1}
                                                                switch v3 {
                                                                case "foo bar":
                                                                    if v4 {
                                                                        if (4294967295 < 18446744073709551615) {
                                                                            if v4 {
                                                                                var v126 int16 = int16(v1)
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    if v4 {
                                                                                        v127 := GenS6{F0: v3, F1: -1, F2: 0}
                                                                                        v126 -= (v126 * 1000)
                                                                                        v128 := GenS12{F0: (^(v1)), F1: v40, F2: (v1 << (((v38 % (((v38 % ((v38) | 1))) | 1))) & 63))}
                                                                                        var v129 uint16 = 0
                                                                                        v130 := GenS15{F0: 4294967295, F1: 255}
                                                                                        var v131 uint = v38
                                                                                        fmt.Println(v4, v39, v131)
                                                                                        _ = v127
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                        _ = v130
                                                                                        _ = v131
                                                                                    }
                                                                                case "foo bar":
                                                                                    v132 := GenS11{F0: v38}
                                                                                    _ = v132
                                                                                }
                                                                                switch 9223372036854775807 {
                                                                                case -430:
                                                                                    fmt.Println(v1, v126, fi42)
                                                                                case 127:
                                                                                    v1 += v1
                                                                                case 452:
                                                                                    v3 = "foo bar"
                                                                                }
                                                                                v3 = v3
                                                                                _ = v126
                                                                            }
                                                                            v3 = v3
                                                                            if true {
                                                                                var v133 uint64 = 1
                                                                                v38 = 2
                                                                                fmt.Println(v133)
                                                                                _ = v133
                                                                            } else {
                                                                                v4 = (!false)
                                                                                switch (v3 + v3) {
                                                                                case "hello":
                                                                                    fmt.Println(v1)
                                                                                case "":
                                                                                    if (!v4) {
                                                                                        _ = len(v2)
                                                                                        v4 = (v3 < v3)
                                                                                    }
                                                                                }
                                                                                v134 := GenS15{F0: 394, F1: 2}
                                                                                var v135 uint32 = v39
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (v3 + "a")
                                                                                }
                                                                                _ = v134
                                                                                _ = v135
                                                                            }
                                                                        } else {
                                                                            switch v3 {
                                                                            case "":
                                                                                var v136 uint = v38
                                                                                _ = v136
                                                                            case "hello":
                                                                                var v137 bool = true
                                                                                _ = v137
                                                                            }
                                                                            v3 = v3
                                                                            var v138 int16 = v40
                                                                            _ = v138
                                                                        }
                                                                        if (-1 != 42) {
                                                                            if ((!true) || ((0.1 != 3.14159) && (((false && false) || ((v40 * v40) < 1)) || v4))) {
                                                                                v2[v3] = fi73
                                                                            }
                                                                            v139 := GenS10{F0: -521, F1: 2147483647, F2: 2}
                                                                            v140 := GenS7{F0: 0, F1: (v40 * int16(v38))}
                                                                            _ = v139
                                                                            _ = v140
                                                                        } else {
                                                                            var v141 uint8 = 42
                                                                            v142 := GenS16{F0: (v40 ^ v40), F1: v141, F2: 18446744073709551615}
                                                                            _ = v141
                                                                            _ = v142
                                                                        }
                                                                        v40 = (v40 >> ((v38) & 63))
                                                                    }
                                                                }
                                                                _ = v125
                                                            }
                                                            v3 = "\n\t"
                                                        }
                                                    case 128:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                }
                                            default:
                                                v1 -= v1
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if (!false) {
                                                var v143 uint16 = 255
                                                v40 += v40
                                                var v144 string = v3
                                                switch (^(fi42)) {
                                                case 128:
                                                    switch 4294967295 {
                                                    case -1:
                                                        switch v3 {
                                                        case "\n\t":
                                                            fmt.Println(v38, v38)
                                                        default:
                                                            fmt.Println(v144)
                                                        }
                                                    case 9223372036854775807:
                                                        var v145 uint = 255
                                                        _ = v145
                                                    case 816:
                                                        if v4 {
                                                            delete(v2, v144)
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                if v4 {
                                                                    var v146 uint = 255
                                                                    var v147 int = fi42
                                                                    _ = v146
                                                                    _ = v147
                                                                }
                                                            case "\n\t":
                                                                v148 := GenS17{F0: -851, F1: 1, F2: 230}
                                                                _ = v148
                                                            case "a":
                                                                v4 = true
                                                            default:
                                                                v144 = (((v144 + "\n\t") + (("foo bar" + ((((v144 + v3) + ("foo bar" + "hello")) + "foo bar") + (v3 + v3))) + v144)) + "hello")
                                                                var v149 uint = (^(v38))
                                                                _ = v149
                                                            }
                                                            if v4 {
                                                                if v4 {
                                                                    v150 := GenS2{F0: 1, F1: 1}
                                                                    var v151 float32 = (0.0 * 1.5)
                                                                    if v4 {
                                                                        var v152 string = "\n\t"
                                                                        var v153 int8 = 42
                                                                        var v154 uint64 = 2
                                                                        _ = v152
                                                                        _ = v153
                                                                        _ = v154
                                                                    }
                                                                    var v155 bool = v4
                                                                    _ = v150
                                                                    _ = v151
                                                                    _ = v155
                                                                } else {
                                                                    var v156 int = (fi42 - (fi42 - (fi73 >> (((v38 % ((uint(v39)) | 1))) & 63))))
                                                                    v3 = v3
                                                                    if true {
                                                                        v144 = ("\n\t" + v144)
                                                                        v144 = ""
                                                                    } else {
                                                                        v38 = (v38 | v38)
                                                                        var v157 uint64 = 18446744073709551615
                                                                        switch (v3 + "") {
                                                                        case "foo bar":
                                                                            if (!(((v40 * v40) != int16(v156)) || v4)) {
                                                                                v38 = v38
                                                                                switch 128 {
                                                                                case -726:
                                                                                    fmt.Println(v157, v39)
                                                                                case 1:
                                                                                    v143 += (v143 & (v143 & (v143 * (v143 / ((2) | 1)))))
                                                                                case -288:
                                                                                    v1 = 50
                                                                                }
                                                                                v156 = v156
                                                                            }
                                                                        case "hello":
                                                                            switch "" {
                                                                            case "":
                                                                                var v158 uint32 = 4294967295
                                                                                _ = v158
                                                                            case "\n\t":
                                                                                v3 = (("hello" + "foo bar") + "foo bar")
                                                                            case "foo bar":
                                                                                delete(v2, (v144 + "a"))
                                                                            default:
                                                                                v159 := GenS4{F0: 31}
                                                                                _ = v159
                                                                            }
                                                                        }
                                                                        _ = v157
                                                                    }
                                                                    _ = v156
                                                                }
                                                            } else {
                                                                if len(v5) > 0 {
                                                                    v5[0] = "\n\t"
                                                                }
                                                                var v160 uint8 = 128
                                                                var v161 uint8 = (v160 >> ((250) & 63))
                                                                if (!v4) {
                                                                    switch v144 {
                                                                    case "foo bar":
                                                                        var v162 int8 = (v1 & v1)
                                                                        _ = v162
                                                                    case "hello":
                                                                        v4 = (float32(float64(-2.25)) <= 0.0)
                                                                    case "":
                                                                        fmt.Println(v160)
                                                                        var v163 bool = false
                                                                        _ = v163
                                                                    default:
                                                                        fmt.Println(v4, v160)
                                                                    }
                                                                    fmt.Println(v4, v4)
                                                                } else {
                                                                    if ((^(fi73)) <= -753) {
                                                                        _ = len(v5)
                                                                        v3 = "hello"
                                                                        var v164 int = int(v161)
                                                                        _ = v164
                                                                    } else {
                                                                        v160 += genFn1(v39, v40)
                                                                        fmt.Println(v4, fi73)
                                                                        var v165 int16 = (v40 >> ((v38) & 63))
                                                                        _ = v165
                                                                    }
                                                                    v3 = ((v3 + "\n\t") + ((v3 + "\n\t") + ((("a" + "") + (("a" + "a") + v3)) + (v144 + v3))))
                                                                    v160 += uint8(v143)
                                                                    if v4 {
                                                                        v5 = append(v5, v144)
                                                                        v1 += (v1 | (v1 << ((2) & 63)))
                                                                    }
                                                                    var v166 int32 = 1
                                                                    _ = v166
                                                                }
                                                                fmt.Println(v143, fi73)
                                                                var v167 uint32 = (v39 >> (((v38 ^ v38)) & 63))
                                                                _ = v160
                                                                _ = v161
                                                                _ = v167
                                                            }
                                                            if false {
                                                                if v4 {
                                                                    v38 -= (v38 & v38)
                                                                    v168 := GenS17{F0: 9223372036854775807, F1: (v143 & v143), F2: 42}
                                                                    var v169 float32 = -2.25
                                                                    _ = v168
                                                                    _ = v169
                                                                } else {
                                                                    var v170 string = "\n\t"
                                                                    if v4 {
                                                                        v144 = v3
                                                                    }
                                                                    var v171 int64 = 42
                                                                    _ = v170
                                                                    _ = v171
                                                                }
                                                                fmt.Println(v4)
                                                            } else {
                                                                var v172 int64 = 730
                                                                _ = v172
                                                            }
                                                            var v173 string = ""
                                                            _ = v173
                                                        }
                                                    default:
                                                        fmt.Println(v143, v1, v144)
                                                        v174 := GenS1{F0: genFn1(65535, v40), F1: 42, F2: (42 <= -9223372036854775808)}
                                                        switch "" {
                                                        case "a":
                                                            var v175 int = (fi42 % (((fi42 * 407)) | 1))
                                                            _ = v175
                                                        case "\n\t":
                                                            var v176 float64 = (1.5 + 1.5)
                                                            _ = v176
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                        _ = v174
                                                    }
                                                default:
                                                    var v177 string = ""
                                                    _ = v177
                                                }
                                                var v178 int16 = v40
                                                _ = v143
                                                _ = v144
                                                _ = v178
                                            } else {
                                                var v179 uint32 = 701
                                                _ = v179
                                            }
                                            switch 4294967295 {
                                            case 32768:
                                                fmt.Println(fi42, v1, v3)
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    var v180 float64 = 3.14159
                                                    _ = v180
                                                case "foo bar":
                                                    if v4 {
                                                        if len(v5) > 0 {
                                                            v5[0] = v3
                                                        }
                                                    } else {
                                                        var v181 uint32 = (v39 - (v39 ^ 1))
                                                        v182 := GenS18{F0: fi73}
                                                        v39 += v181
                                                        _ = v181
                                                        _ = v182
                                                    }
                                                default:
                                                    fmt.Println(v1, fi73)
                                                }
                                            }
                                            switch "" {
                                            case "a":
                                                if (true && false) {
                                                    var v183 int32 = 0
                                                    _ = v183
                                                }
                                            case "":
                                                var v184 float64 = (3.14159 - (0.0 * float64(1e10)))
                                                _ = v184
                                            case "hello":
                                                var v185 int = fi73
                                                _ = v185
                                            default:
                                                v5 = append(v5, v3)
                                            }
                                        } else {
                                            v4 = v4
                                            switch v3 {
                                            case "":
                                                v186 := GenS12{F0: v1, F1: v40, F2: v1}
                                                _ = v186
                                            case "\n\t":
                                                v187 := GenS8{F0: (((("\n\t" + "hello") <= (v3 + "\n\t")) || v4) && v4), F1: -2147483648}
                                                _ = v187
                                            default:
                                                if (!false) {
                                                    fmt.Println(v40, v3)
                                                    var v188 int = fi73
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v188
                                                } else {
                                                    if v4 {
                                                        delete(v2, (v3 + ((("" + ((v3 + "\n\t") + (v3 + v3))) + (v3 + ("\n\t" + "foo bar"))) + "foo bar")))
                                                        var v189 int32 = 320
                                                        if v4 {
                                                            fmt.Println(v1, v1)
                                                            if (0 > (v189 * (^(v189)))) {
                                                                v190 := GenS13{F0: 255}
                                                                v191 := GenS19{F0: v39, F1: 32767}
                                                                _ = len(v5)
                                                                v192 := GenS20{F0: 4294967295, F1: fi73, F2: genFn1(0, v40)}
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v193 uint8 = 128
                                                                    _ = v193
                                                                case "foo bar":
                                                                    v2[(v3 + "hello")] = (fi42 - (fi42 * fi73))
                                                                case "":
                                                                    v194 := GenS21{F0: (v1 + (v1 & (-(v1))))}
                                                                    _ = v194
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                                _ = v190
                                                                _ = v191
                                                                _ = v192
                                                            }
                                                            if false {
                                                                v1 = -128
                                                                v38 = (v38 + v38)
                                                                v195 := GenS22{F0: 65535}
                                                                _ = v195
                                                            }
                                                        }
                                                        _ = v189
                                                    } else {
                                                        v196 := GenS4{F0: v1}
                                                        var v197 int8 = -125
                                                        _ = v196
                                                        _ = v197
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if v4 {
                                                            fmt.Println(fi73, v3, fi42)
                                                            switch v3 {
                                                            case "a":
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v198 uint32 = v39
                                                                    _ = v198
                                                                case "":
                                                                    fmt.Println(fi73, fi73, v40)
                                                                case "a":
                                                                    switch (fi42 + fi42) {
                                                                    case -2147483648:
                                                                        var v199 int8 = -1
                                                                        _ = v199
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                var v200 uint8 = 2
                                                                _ = v200
                                                            }
                                                            fmt.Println(v40)
                                                        } else {
                                                            if (2147483647 <= -2147483648) {
                                                                v201 := GenS13{F0: 0}
                                                                if v4 {
                                                                    var v202 uint = (v38 >> ((v38) & 63))
                                                                    fmt.Println(v38, v1, v39)
                                                                    v40 = v40
                                                                    _ = v202
                                                                }
                                                                _ = v201
                                                            } else {
                                                                var v203 uint = v38
                                                                v1 = (v1 * v1)
                                                                _ = v203
                                                            }
                                                            v204 := v2["foo bar"]
                                                            var v205 int32 = -278
                                                            _ = v204
                                                            _ = v205
                                                        }
                                                    case "\n\t":
                                                        v40 = -1
                                                    }
                                                    var v206 uint32 = 0
                                                    var v207 float32 = 3.14159
                                                    var v208 bool = (v40 <= (v40 * -366))
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v206
                                                    _ = v207
                                                    _ = v208
                                                }
                                            }
                                        }
                                        var v209 uint16 = 1
                                        _ = v209
                                        _ = v97
                                    } else {
                                        v4 = v4
                                    }
                                    v1 += v1
                                }
                                _ = fi73
                                _ = v74
                            }
                        }
                        fmt.Println(v4)
                        _ = fi42
                    }
                    _ = v41
                } else {
                    for fi210 := 0; fi210 < 14; fi210++ {
                        fmt.Println(v40, v39)
                        v40 = (v40 | (v40 << ((v38) & 63)))
                        var v211 float64 = 3.14159
                        _ = fi210
                        _ = v211
                    }
                    if v4 {
                        var v212 string = v3
                        if true {
                            v213 := GenS14{F0: 286}
                            v2[(v212 + (v212 + (v212 + v212)))] = -1
                            _ = v213
                        } else {
                            var v214 float32 = float32((1.5 / 1.0))
                            _ = v214
                        }
                        if (((1e10 <= (1e10 - -2.25)) && true) || false) {
                            if v4 {
                                var v215 float64 = ((3.14159 / -2.25) * float64(((0.0 / float32(1.5)) + (((float32(3.14159) - -2.25) - (-2.25 / float32(0.1))) + 0.0))))
                                _ = v215
                            } else {
                                v39 = v39
                                v39 = 2
                                switch "" {
                                case "a":
                                    if v4 {
                                        for fi216 := 0; fi216 < 14; fi216++ {
                                            fmt.Println(v1)
                                            if v4 {
                                                v217 := GenS13{F0: 254}
                                                v212 = (("foo bar" + v212) + v212)
                                                _ = v217
                                            }
                                            _ = fi216
                                        }
                                        fmt.Println(v1, v3, v1)
                                        v4 = ((v4 || v4) && v4)
                                    } else {
                                        if (0.0 <= -2.25) {
                                            if (0 == int32(v39)) {
                                                v218 := GenS23{F0: (v38 << (((^(v38))) & 63))}
                                                _ = v218
                                            }
                                            fmt.Println(v1, v40, v1)
                                            v219 := GenS8{F0: (false || (1.5 == float32((1.5 * 0.1)))), F1: -2147483648}
                                            _ = v219
                                        }
                                    }
                                default:
                                    v5 = append(v5, v3)
                                }
                            }
                            fmt.Println(v38, v40, v212)
                            v220 := GenS24{F0: 1, F1: v4, F2: v1}
                            _ = v220
                        }
                        _ = v212
                    }
                    switch (v3 + v3) {
                    case "\n\t":
                        for fi221 := 0; fi221 < 16; fi221++ {
                            if false {
                                var v222 uint8 = 0
                                var v223 uint16 = 65535
                                var v224 int16 = v40
                                _ = v222
                                _ = v223
                                _ = v224
                            }
                            _ = fi221
                        }
                    case "foo bar":
                        fmt.Println(v4, v40, v40)
                    case "":
                        var v225 uint8 = 1
                        _ = v225
                    default:
                        if (616 > 129) {
                            v38 -= v38
                            if len(v5) > 0 {
                                v5[0] = v3
                            }
                        } else {
                            var v226 int = int(v38)
                            for fi227 := 0; fi227 < 8; fi227++ {
                                switch fi227 {
                                case 127:
                                    v228 := GenS25{F0: (v40 - 397), F1: 541, F2: v4}
                                    _ = v228
                                }
                                for fi229 := 0; fi229 < 5; fi229++ {
                                    var v230 float64 = 1.5
                                    if (-1 >= 42) {
                                        var v231 int = fi229
                                        v4 = true
                                        if v4 {
                                            if (4294967295 < v38) {
                                                v232 := v2[(v3 + ("hello" + v3))]
                                                _ = v232
                                            } else {
                                                fmt.Println(v40)
                                                v40 -= v40
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = v3
                                            }
                                            var v233 float64 = v230
                                            _ = v233
                                        } else {
                                            var v234 uint16 = 898
                                            var v235 int8 = (v1 << ((97) & 63))
                                            fmt.Println(v3)
                                            _ = v234
                                            _ = v235
                                        }
                                        if v4 {
                                            if v4 {
                                                var v236 float32 = float32(v230)
                                                _ = v236
                                            }
                                            var v237 uint32 = (v39 % ((v39) | 1))
                                            if (!true) {
                                                if v4 {
                                                    v238 := GenS10{F0: 0, F1: -1, F2: 255}
                                                    var v239 float64 = 0.0
                                                    fmt.Println(v231, v4, v38)
                                                    fmt.Println(v38, v237, v1)
                                                    v239 = ((v230 + (0.0 / v239)) / ((v239 / ((v230 / v230) * (float64(3.14159) - -2.25))) + ((((v230 / v230) - (0.1 - ((v230 / ((3.14159 * v239) * 0.0)) / v230))) + v239) + v239)))
                                                    var v240 int8 = v1
                                                    var v241 uint8 = 1
                                                    v242 := GenS26{F0: (v241 / ((v241) | 1))}
                                                    if (float64(0.1) < v239) {
                                                        if (v4 && (!(9223372036854775807 >= 42))) {
                                                            if (true && (v4 && v4)) {
                                                                v230 -= (v230 + ((float64(0.0) / float64(float32((v230 + v230)))) + (float64(0.0) / (3.14159 / (-2.25 * v239)))))
                                                                v2[v3] = 676
                                                            }
                                                            switch (v3 + ((v3 + (v3 + ("\n\t" + v3))) + v3)) {
                                                            case "":
                                                                v241 = 0
                                                            case "a":
                                                                var v243 float32 = (1.5 + float32(v239))
                                                                v2["foo bar"] = v231
                                                                _ = v243
                                                            case "hello":
                                                                var v244 float32 = 0.1
                                                                _ = v244
                                                            default:
                                                                v231 = fi229
                                                            }
                                                            var v245 uint64 = 18446744073709551615
                                                            _ = v245
                                                        } else {
                                                            v246 := GenS27{F0: 2}
                                                            if (v239 <= 1e10) {
                                                                v247 := GenS28{F0: (v231 / (((fi229 ^ 4294967295)) | 1))}
                                                                _ = v247
                                                            }
                                                            if (!(0.1 < 0.0)) {
                                                                var v248 float64 = v230
                                                                _ = v248
                                                            }
                                                            _ = len(v2)
                                                            v226 += v231
                                                            _ = v246
                                                        }
                                                        switch (v226 | (v231 << (((v38 + v38)) & 63))) {
                                                        case 439:
                                                            var v249 uint64 = 42
                                                            _ = v249
                                                        case 4294967295:
                                                            if (-2147483648 >= -830) {
                                                                var v250 uint32 = uint32(v241)
                                                                if ((v240 * (v240 * (v240 % ((-128) | 1)))) == 1) {
                                                                    if (!(656 == 18446744073709551615)) {
                                                                        fmt.Println(v240)
                                                                    } else {
                                                                        var v251 uint16 = 42
                                                                        fmt.Println(v4, v231, v38)
                                                                        if (false || (597 != (v251 & v251))) {
                                                                            if true {
                                                                                var v252 uint8 = 139
                                                                                v241 = v241
                                                                                _ = v252
                                                                            }
                                                                            fmt.Println(v230)
                                                                        } else {
                                                                            var v253 bool = v4
                                                                            _ = len(v5)
                                                                            _ = v253
                                                                        }
                                                                        _ = v251
                                                                    }
                                                                    v254 := v2[v3]
                                                                    _ = v254
                                                                }
                                                                _ = v250
                                                            }
                                                        }
                                                    }
                                                    if v4 {
                                                        v1 = v240
                                                        v230 = float64((0.1 / 0.1))
                                                    } else {
                                                        fmt.Println(fi229)
                                                        if (float32(1e10) >= 1.5) {
                                                            if v4 {
                                                                var v255 uint8 = v241
                                                                var v256 uint = (v38 % ((251) | 1))
                                                                _ = v255
                                                                _ = v256
                                                            } else {
                                                                v226 = 128
                                                                var v257 uint16 = 65535
                                                                _ = v257
                                                            }
                                                            if v4 {
                                                                var v258 uint8 = (v241 - uint8(v237))
                                                                fmt.Println(v4, v230)
                                                                _ = v258
                                                            }
                                                            if ((v4 && v4) && v4) {
                                                                switch (v3 + v3) {
                                                                case "a":
                                                                    var v259 int64 = 42
                                                                    _ = v259
                                                                case "foo bar":
                                                                    var v260 uint8 = (v241 & uint8(v38))
                                                                    _ = v260
                                                                default:
                                                                    v261 := GenS12{F0: (v1 + (^(v240))), F1: (v40 + v40), F2: v1}
                                                                    _ = v261
                                                                }
                                                                v3 = "foo bar"
                                                            } else {
                                                                if ((false && v4) && v4) {
                                                                    v3 = "hello"
                                                                    v262 := GenS29{F0: true, F1: 9223372036854775807, F2: 1}
                                                                    v38 = v38
                                                                    _ = v262
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v238
                                                    _ = v239
                                                    _ = v240
                                                    _ = v241
                                                    _ = v242
                                                } else {
                                                    v38 -= v38
                                                    var v263 uint64 = 1
                                                    v231 -= 1
                                                    _ = v263
                                                }
                                                v1 += (v1 % ((v1) | 1))
                                            }
                                            _ = v237
                                        }
                                        _ = v231
                                    }
                                    v264 := GenS27{F0: 2}
                                    _ = fi229
                                    _ = v230
                                    _ = v264
                                }
                                v265 := GenS27{F0: 956}
                                _ = fi227
                                _ = v265
                            }
                            var v266 uint16 = 32767
                            _ = v226
                            _ = v266
                        }
                    }
                }
                fmt.Println(v38, v40)
                for fi267 := 0; fi267 < 3; fi267++ {
                    if false {
                        v39 -= v39
                    }
                    _ = fi267
                }
            } else {
                var v268 uint32 = 4294967295
                if (-306 == v40) {
                    v3 = ("foo bar" + v3)
                    if (v4 && v4) {
                        v269 := GenS12{F0: int8(v40), F1: (^(v40)), F2: int8(v39)}
                        v4 = (float32(1.5) != 0.0)
                        _ = v269
                    }
                }
                v270 := GenS11{F0: (v38 % ((v38) | 1))}
                _ = v268
                _ = v270
            }
            if (!((!((255 <= 32767) && (2 != 1))) || v4)) {
                var v271 int8 = v1
                if v4 {
                    if len(v5) > 0 {
                        v5[0] = v3
                    }
                }
                _ = v271
            }
            var v272 int = 127
            _ = v272
        }
    case "a":
        if true {
            var v273 float32 = 0.0
            v38 = 0
            _ = v273
        } else {
            var v274 int16 = v40
            for fi275 := 0; fi275 < 6; fi275++ {
                v40 += v40
                if true {
                    var v276 uint64 = uint64(v39)
                    v277 := GenS30{F0: 251, F1: v38}
                    _ = v276
                    _ = v277
                } else {
                    for fi278 := 0; fi278 < 8; fi278++ {
                        switch fi278 {
                        case 665:
                            if v4 {
                                _ = len(v2)
                                var v279 int64 = 24
                                _ = v279
                            } else {
                                if v4 {
                                    var v280 string = v3
                                    fmt.Println(v3, v3)
                                    v39 += 449
                                    _ = v280
                                }
                            }
                        }
                        _ = fi278
                    }
                    v40 -= 1000
                    for fi281 := 0; fi281 < 14; fi281++ {
                        if ((v3 + (((v3 + (v3 + "")) + v3) + v3)) == "hello") {
                            switch fi275 {
                            case -2147483648:
                                var v282 uint8 = genFn1((v39 | (v39 & v39)), v274)
                                _ = v282
                            case 766:
                                if v4 {
                                    var v283 int8 = (v1 & v1)
                                    _ = v283
                                } else {
                                    fmt.Println(v39, v4)
                                    v274 += v40
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v1, fi281)
                                    case "hello":
                                        if (1 == int64(v39)) {
                                            v274 = (v274 - -32768)
                                            v40 += (-(v40))
                                        }
                                    default:
                                        v1 += v1
                                    }
                                    fmt.Println(v38, v40)
                                }
                            }
                            v284 := GenS19{F0: (v39 - v39), F1: 210}
                            v38 = (v38 + 2)
                            _ = v284
                        } else {
                            var v285 uint32 = v39
                            v39 = v39
                            _ = v285
                        }
                        var v286 int16 = v274
                        _ = fi281
                        _ = v286
                    }
                }
                var v287 int64 = -1
                _ = fi275
                _ = v287
            }
            _ = v274
        }
    case "\n\t":
        fmt.Println(v39)
    default:
        var v288 uint8 = 1
        _ = v288
    }
    fmt.Println(v3, v39, v4)
    switch 32767 {
    case -2147483648:
        var v289 int32 = 0
        _ = v289
    default:
        for fi290 := 0; fi290 < 13; fi290++ {
            if (113 != 255) {
                var v291 uint8 = 86
                v38 -= (v38 & v38)
                delete(v2, v3)
                _ = v291
            } else {
                v1 = v1
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi290
        }
    }
    for fi292 := 0; fi292 < 16; fi292++ {
        var v293 uint64 = 4294967295
        v4 = (v4 || false)
        v2[v3] = int(v40)
        _ = fi292
        _ = v293
    }
    v3 = ("\n\t" + v3)
    if ((57 <= 42) || true) {
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v294 := GenS12{F0: v1, F1: (v40 & (^(v40))), F2: v1}
        _ = v294
    }
    v4 = (2 == 219)
    var v295 float64 = (0.0 + (-2.25 / 1e10))
    var v296 int32 = -566
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v295:", v295)
    fmt.Println("v296:", v296)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v295
    _ = v296
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
