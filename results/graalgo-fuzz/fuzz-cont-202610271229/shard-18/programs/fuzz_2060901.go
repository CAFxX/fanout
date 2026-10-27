package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64; F1 bool; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 int16; F1 uint16; F2 int16 }
type GenS10 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS11 struct { F0 uint8 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 uint64 = 1
    v5 := GenS1{F0: -106, F1: -686}
    if ((!false) || false) {
        switch 127 {
        case 32767:
            v6 := v2[4294967295]
            _ = v6
        case 0:
            v1 = v1
        default:
            v7 := []int{-629, -9223372036854775808, -1, 2147483647}
            _ = v7
        }
        switch 127 {
        case -228:
            switch 32768 {
            case -942:
                switch 577 {
                case 128:
                    if (4294967295 < (v4 + 471)) {
                        v8 := GenS2{F0: 2}
                        switch v1 {
                        case "\n\t":
                            for fi9 := 0; fi9 < 4; fi9++ {
                                fmt.Println(fi9)
                                _ = fi9
                            }
                        case "":
                            v1 = ("\n\t" + v1)
                        }
                        _ = v8
                    }
                case -1:
                    var v10 int32 = -2147483648
                    _ = v10
                }
            case 1:
                for fi11 := 0; fi11 < 18; fi11++ {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 int64 = 42
                        v4 = v4
                        var v14 float32 = 0.0
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                }
            default:
                switch 9223372036854775807 {
                case 127:
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 bool = false
                        v17 := []int{(-(fi15)), -571, -2147483648}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case -801:
                    v18 := []int{127, 32767}
                    _ = v18
                case 728:
                    var v19 uint64 = 4294967295
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v1 = v1
                        var v21 uint8 = 2
                        _ = fi20
                        _ = v21
                    }
                }
            }
        case 779:
            switch 128 {
            case -735:
                if (true && false) {
                    v22 := GenS3{F0: 18446744073709551615, F1: 42}
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if (1 != -1) {
                            fmt.Println(v4, v1, fi23)
                        }
                        _ = fi23
                    }
                    _ = v22
                }
            case 127:
                switch "foo bar" {
                case "":
                    v24 := GenS4{F0: -2147483648, F1: v1}
                    _ = v24
                }
            }
        case 4294967295:
            v3[-9223372036854775808] = 2147483647
        default:
            v25 := []int{-1}
            _ = v25
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            switch (v1 + ("hello" + (v1 + v1))) {
            case "\n\t":
                for fi27 := 0; fi27 < 7; fi27++ {
                    if true {
                        if true {
                            v28 := GenS2{F0: 2147483647}
                            fmt.Println(fi27, v1)
                            _ = v28
                        } else {
                            var v29 int32 = -1
                            var v30 int64 = 42
                            _ = v29
                            _ = v30
                        }
                        if (uint16(v4) <= 65535) {
                            v31 := GenS4{F0: 288, F1: ("\n\t" + v1)}
                            var v32 uint64 = 2
                            if ((1 >= 127) || false) {
                                var v33 int = (fi26 & fi27)
                                if false {
                                    v34 := GenS3{F0: 1, F1: 52}
                                    v35 := GenS4{F0: 211, F1: ""}
                                    _ = v34
                                    _ = v35
                                }
                                _ = v33
                            }
                            _ = v31
                            _ = v32
                        } else {
                            v1 = v1
                            v36 := []int{(fi26 / ((-9223372036854775808) | 1)), (fi27 ^ 799)}
                            v37 := GenS4{F0: -725, F1: v1}
                            _ = v36
                            _ = v37
                        }
                        v2[0] = 2147483648
                    } else {
                        var v38 bool = (((((1e10 + 0.0) * (-2.25 * 1.5)) - 3.14159) * -2.25) >= ((3.14159 - 0.0) / (1e10 / ((0.1 / (0.1 + 3.14159)) * (1.5 - 0.1)))))
                        _ = v38
                    }
                    _ = fi27
                }
            case "a":
                switch "a" {
                case "foo bar":
                    v1 = "\n\t"
                case "hello":
                    v39 := []int{}
                    _ = v39
                default:
                    var v40 int8 = 127
                    _ = v40
                }
            default:
                if (false || false) {
                    if false {
                        for fi41 := 0; fi41 < 20; fi41++ {
                            var v42 int8 = 122
                            if true {
                                var v43 uint8 = 255
                                var v44 uint64 = (v4 * (v4 << ((42) & 63)))
                                v45 := GenS2{F0: 2147483647}
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                var v46 uint32 = 2147483648
                                switch (fi41 - (fi26 << ((2) & 63))) {
                                case 32768:
                                    switch (fi26 / (((fi26 * fi41)) | 1)) {
                                    case 1:
                                        v47 := GenS5{F0: 0, F1: (0.1 != 1e10), F2: 9223372036854775807}
                                        _ = v47
                                    case -1:
                                        var v48 uint8 = 1
                                        _ = v48
                                    case 2147483647:
                                        switch fi26 {
                                        case 953:
                                            var v49 int = fi41
                                            _ = v49
                                        case -905:
                                            switch (-(fi26)) {
                                            case 42:
                                                v46 += v46
                                            case 0:
                                                var v50 int = (fi26 / ((-2147483648) | 1))
                                                _ = v50
                                            case -9223372036854775808:
                                                if false {
                                                    switch fi26 {
                                                    case -607:
                                                        var v51 uint32 = (^(v46))
                                                        _ = v51
                                                    case -260:
                                                        v52 := GenS4{F0: 1, F1: (("hello" + (("a" + "") + (v1 + ""))) + v1)}
                                                        _ = v52
                                                    case 166:
                                                        var v53 float64 = 0.0
                                                        _ = v53
                                                    }
                                                    var v54 int64 = 0
                                                    _ = v54
                                                }
                                            default:
                                                v55 := []int{}
                                                _ = v55
                                            }
                                        case 42:
                                            var v56 float64 = float64(3.14159)
                                            _ = v56
                                        }
                                    default:
                                        v3[2147483647] = fi41
                                    }
                                    if true {
                                        fmt.Println(v46, v4)
                                        v1 = ("hello" + v1)
                                        var v57 bool = true
                                        _ = v57
                                    }
                                case 128:
                                    if (1 > -779) {
                                        if ((v1 + (((v1 + v1) + v1) + ("" + v1))) >= ((((v1 + v1) + v1) + v1) + "a")) {
                                            var v58 int16 = 0
                                            _ = v58
                                        } else {
                                            v59 := GenS4{F0: 42, F1: "foo bar"}
                                            switch (^(fi26)) {
                                            case 42:
                                                switch v1 {
                                                case "":
                                                    v60 := []int{fi41, fi26, (fi41 ^ fi26)}
                                                    _ = v60
                                                case "a":
                                                    v61 := v2[42]
                                                    _ = v61
                                                case "hello":
                                                    v46 += v46
                                                }
                                            case -210:
                                                v62 := GenS2{F0: v46}
                                                _ = v62
                                            case 32767:
                                                switch 0 {
                                                case 1:
                                                    v46 = v46
                                                case 9223372036854775807:
                                                    var v63 float32 = ((0.1 * 3.14159) / -2.25)
                                                    _ = v63
                                                default:
                                                    _ = len(v3)
                                                }
                                            }
                                            var v64 string = ((v1 + (v1 + (v1 + v1))) + "a")
                                            _ = v59
                                            _ = v64
                                        }
                                        switch v1 {
                                        case "a":
                                            v65 := v3[(fi41 ^ 127)]
                                            _ = v65
                                        }
                                        v3[(^(fi41))] = (fi26 - 2147483647)
                                    }
                                }
                                switch ("" + v1) {
                                case "\n\t":
                                    v4 = 116
                                case "a":
                                    var v66 uint16 = 1
                                    switch 128 {
                                    case 128:
                                        switch (fi41 << ((18446744073709551615) & 63)) {
                                        case 0:
                                            v1 = v1
                                        case 9223372036854775807:
                                            v66 = (v66 - (v66 - v66))
                                        }
                                        var v67 uint64 = 1
                                        _ = v67
                                    case 9223372036854775807:
                                        fmt.Println(fi26)
                                    case -2147483648:
                                        var v68 int = (fi41 & (fi26 * fi26))
                                        _ = v68
                                    }
                                    fmt.Println(fi41, v46, v42)
                                    _ = v66
                                case "hello":
                                    switch 813 {
                                    case -120:
                                        switch (fi41 + fi26) {
                                        case 157:
                                            switch "hello" {
                                            case "":
                                                v3[(fi26 / ((fi26) | 1))] = (fi26 ^ fi26)
                                            }
                                            if (false || (!true)) {
                                                fmt.Println(fi41, fi41, v1)
                                                v69 := v3[(fi26 & (fi41 & fi41))]
                                                _ = v69
                                            } else {
                                                v70 := []int{32767, -790, (fi26 ^ (fi26 + fi26))}
                                                _ = v70
                                            }
                                            fmt.Println(v4, v46)
                                        }
                                        var v71 int16 = 1
                                        switch ((("\n\t" + v1) + v1) + (v1 + "\n\t")) {
                                        case "foo bar":
                                            switch fi26 {
                                            case -2147483648:
                                                v4 += 1
                                            default:
                                                switch v1 {
                                                case "\n\t":
                                                    _ = len(v3)
                                                    var v72 bool = true
                                                    _ = v72
                                                case "a":
                                                    v73 := v2[(v46 % ((uint32(v4)) | 1))]
                                                    _ = v73
                                                }
                                            }
                                        case "\n\t":
                                            if (true && (int8(fi41) != v42)) {
                                                v71 += (^(v71))
                                                v74 := GenS1{F0: (v42 % (((v42 - (v42 << ((uint(v4)) & 63)))) | 1)), F1: -32768}
                                                var v75 bool = (!true)
                                                _ = v74
                                                _ = v75
                                            }
                                        default:
                                            var v76 uint16 = 609
                                            _ = v76
                                        }
                                        _ = v71
                                    case 128:
                                        v77 := []int{(fi26 % (((^(fi41))) | 1))}
                                        _ = v77
                                    case 1:
                                        var v78 int8 = v42
                                        _ = v78
                                    }
                                default:
                                    v79 := GenS3{F0: 0, F1: 1}
                                    _ = v79
                                }
                                v80 := v2[2147483647]
                                var v81 int8 = (v42 >> ((18446744073709551615) & 63))
                                v80 -= (v80 >> ((0) & 63))
                                v1 = ((v1 + "\n\t") + ("\n\t" + v1))
                                _ = v46
                                _ = v80
                                _ = v81
                            }
                            v42 += (^(v42))
                            _ = fi41
                            _ = v42
                        }
                    }
                    if true {
                        var v82 string = ""
                        _ = v82
                    } else {
                        v4 = (v4 / ((v4) | 1))
                    }
                }
            }
            switch fi26 {
            case 32768:
                if (((v1 + ("" + (v1 + ((("hello" + v1) + "") + v1)))) + v1) < (v1 + (v1 + (v1 + v1)))) {
                    for fi83 := 0; fi83 < 6; fi83++ {
                        if (true || true) {
                            switch v1 {
                            case "\n\t":
                                v84 := GenS6{F0: 2147483648}
                                _ = v84
                            default:
                                if true {
                                    v1 = "hello"
                                    var v85 bool = (!(!(!true)))
                                    _ = v85
                                }
                            }
                            if false {
                                v86 := GenS7{F0: fi83, F1: 0}
                                _ = v86
                            } else {
                                var v87 int32 = int32(fi83)
                                switch v1 {
                                case "\n\t":
                                    v88 := GenS8{F0: 255, F1: -272}
                                    _ = v88
                                default:
                                    fmt.Println(fi26, v4, v4)
                                }
                                _ = v87
                            }
                            var v89 float32 = -2.25
                            _ = v89
                        } else {
                            var v90 int32 = 42
                            _ = v90
                        }
                        fmt.Println(fi26, v1)
                        _ = fi83
                    }
                    switch "" {
                    case "":
                        v1 = (v1 + ((((("" + "") + "") + "hello") + v1) + ""))
                    case "foo bar":
                        var v91 bool = (331 < 2)
                        _ = v91
                    }
                } else {
                    for fi92 := 0; fi92 < 8; fi92++ {
                        if true {
                            if true {
                                var v93 int8 = 48
                                _ = v93
                            } else {
                                v1 = (("\n\t" + (v1 + "")) + ((((v1 + v1) + "hello") + v1) + (v1 + "hello")))
                                v94 := []int{fi26, fi26}
                                _ = v94
                            }
                            v4 += v4
                            var v95 bool = true
                            v96 := GenS3{F0: 1, F1: 2}
                            _ = v95
                            _ = v96
                        } else {
                            fmt.Println(v1)
                            var v97 float64 = float64((1.5 * float32(0.0)))
                            v98 := v2[42]
                            _ = v97
                            _ = v98
                        }
                        fmt.Println(v4, v4, fi26)
                        v1 = (v1 + v1)
                        _ = fi92
                    }
                    v1 = (v1 + "\n\t")
                    v1 = ("\n\t" + ("hello" + v1))
                }
            case 518:
                v99 := GenS9{F0: 32767, F1: 313, F2: -32768}
                _ = v99
            }
            _ = fi26
        }
    } else {
        if true {
            v1 = v1
        }
    }
    delete(v3, int(v4))
    v100 := GenS10{F0: 0, F1: 32767, F2: 812}
    if false {
        switch 32767 {
        case -973:
            v101 := GenS8{F0: 1, F1: -1}
            _ = v101
        case 32768:
            fmt.Println(v4, v4)
        case -2147483648:
            for fi102 := 0; fi102 < 8; fi102++ {
                v103 := v3[int(v4)]
                _ = fi102
                _ = v103
            }
        }
        for fi104 := 0; fi104 < 5; fi104++ {
            var v105 int16 = 32767
            _ = fi104
            _ = v105
        }
    } else {
        v106 := GenS11{F0: 154}
        var v107 uint = 442
        var v108 uint64 = v4
        _ = v106
        _ = v107
        _ = v108
    }
    if true {
        var v109 uint32 = 2147483647
        _ = v109
    }
    switch (v1 + v1) {
    case "foo bar":
        v1 = v1
    default:
        v4 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
