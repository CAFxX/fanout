package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int32; F2 int32 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 bool; F2 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int16; F1 uint8; F2 string }
type GenS14 struct { F0 uint; F1 uint32 }
type GenS15 struct { F0 uint; F1 uint16 }
type GenS16 struct { F0 int; F1 int64 }
func genFn0(p0 uint64, p1 uint) uint {
    v4 := GenS1{F0: 1}
    fmt.Println(p1)
    p1 += (p1 >> ((p1) & 63))
    if true {
        if true {
            var v5 int = -2147483648
            _ = v5
        }
    } else {
        var v6 bool = ((42 <= -266) || false)
        if (1000 == 1) {
            fmt.Println(p0)
        }
        if (!(("\n\t" + (("foo bar" + ("\n\t" + ("a" + ("" + "foo bar")))) + (((("foo bar" + "\n\t") + "foo bar") + "a") + ""))) != "\n\t")) {
            v6 = (-9223372036854775808 <= -796)
            var v7 string = ("" + "\n\t")
            _ = v7
        }
        _ = v6
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{(fi8 % ((fi8) | 1)), fi8, (fi8 * (fi8 % ((fi8) | 1))), (fi8 | fi8)}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return (p1 * p1)
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 814
    var v3 float64 = -2.25
    fmt.Println(v1, v1, v2)
    var v10 string = v1
    v10 = ("\n\t" + ("\n\t" + v1))
    if false {
        if true {
            var v11 uint32 = 2
            v12 := GenS1{F0: 494}
            switch v1 {
            case "":
                var v13 float64 = 1e10
                _ = v13
            case "hello":
                v14 := GenS1{F0: 42}
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch fi15 {
                    case 650:
                        if true {
                            v11 = (v11 & v11)
                            v2 = (v2 >> ((250) & 63))
                            switch (fi15 | (fi15 << ((18446744073709551615) & 63))) {
                            case -721:
                                if false {
                                    var v16 bool = (1 == 89)
                                    if (v16 || false) {
                                        for fi17 := 0; fi17 < 1; fi17++ {
                                            if (2147483647 >= -2147483648) {
                                                var v18 float32 = 1.5
                                                v19 := []int{(fi17 | (-(fi17))), (fi17 & (-(fi15))), -2147483648, fi17}
                                                v2 = (v2 ^ (v2 * (^(v2))))
                                                v10 = (v1 + v10)
                                                switch v1 {
                                                case "hello":
                                                    switch (fi15 - fi17) {
                                                    case 0:
                                                        v20 := GenS1{F0: -460}
                                                        _ = v20
                                                    default:
                                                        v21 := GenS1{F0: 32767}
                                                        _ = v21
                                                    }
                                                case "":
                                                    var v22 bool = (!(4294967295 > 4294967295))
                                                    _ = v22
                                                case "a":
                                                    if v16 {
                                                        if false {
                                                            switch (v10 + "") {
                                                            case "\n\t":
                                                                v23 := GenS1{F0: 32767}
                                                                _ = v23
                                                            default:
                                                                v16 = ((255 > 127) && (2147483647 == 1))
                                                                v2 = v2
                                                                v18 = (1.5 / v18)
                                                                if v16 {
                                                                    if len(v19) > 0 {
                                                                        v19[0] = fi15
                                                                    }
                                                                    v24 := GenS1{F0: 32767}
                                                                    if ((1.5 < float32(v3)) || true) {
                                                                        if (!(!(-128 <= -101))) {
                                                                            fmt.Println(v18, fi17, fi15)
                                                                            v25 := GenS1{F0: -944}
                                                                            var v26 uint = genFn0(4294967295, 2)
                                                                            _ = v25
                                                                            _ = v26
                                                                        }
                                                                        v2 = (v2 * v2)
                                                                        if (((^(v2)) <= 2) && (2 != (v11 | v11))) {
                                                                            var v27 float64 = (v3 * v3)
                                                                            var v28 int64 = -9223372036854775808
                                                                            v19 = append(v19, fi15)
                                                                            _ = v27
                                                                            _ = v28
                                                                        } else {
                                                                            if true {
                                                                                v19 = append(v19, fi15)
                                                                            } else {
                                                                                if len(v19) > 0 {
                                                                                    v19[0] = (fi17 + fi15)
                                                                                }
                                                                                var v29 int = (fi17 >> ((4294967295) & 63))
                                                                                v30 := GenS2{F0: 943, F1: 128}
                                                                                _ = v29
                                                                                _ = v30
                                                                            }
                                                                        }
                                                                    }
                                                                    switch (fi15 + (fi15 << ((uint(v2)) & 63))) {
                                                                    case 128:
                                                                        if (((v16 || false) || v16) && false) {
                                                                            switch 2147483647 {
                                                                            case -2147483648:
                                                                                fmt.Println(v16, v16)
                                                                            case 2147483647:
                                                                                if true {
                                                                                    var v31 float64 = v3
                                                                                    switch ("hello" + (v10 + v10)) {
                                                                                    case "hello":
                                                                                        var v32 uint8 = 215
                                                                                        _ = v32
                                                                                    }
                                                                                    _ = v31
                                                                                } else {
                                                                                    v33 := GenS3{F0: (v11 ^ uint32(v2)), F1: 116, F2: v1}
                                                                                    if (1 > 100) {
                                                                                        v18 += ((v18 + v18) * -2.25)
                                                                                        if len(v19) > 0 {
                                                                                            v19[0] = (fi15 * 42)
                                                                                        }
                                                                                        v34 := GenS4{F0: (v11 * 42)}
                                                                                        v35 := GenS4{F0: v11}
                                                                                        v16 = (-1 == -2147483648)
                                                                                        _ = v34
                                                                                        _ = v35
                                                                                    } else {
                                                                                        var v36 float32 = ((float32(0.1) + (v18 + v18)) + ((((-2.25 / v18) / 3.14159) - (float32(v3) - (0.0 / v18))) * (v18 / 0.1)))
                                                                                        _ = v36
                                                                                    }
                                                                                    _ = v33
                                                                                }
                                                                            }
                                                                            switch v1 {
                                                                            case "hello":
                                                                                v1 = ((((v10 + "") + "foo bar") + v1) + (v1 + v10))
                                                                                var v37 uint64 = 623
                                                                                v38 := GenS3{F0: (^(v11)), F1: 585, F2: v10}
                                                                                _ = v37
                                                                                _ = v38
                                                                            case "foo bar":
                                                                                fmt.Println(v2, fi15)
                                                                            case "a":
                                                                                switch (fi15 / ((42) | 1)) {
                                                                                case 4294967295:
                                                                                    if len(v19) >= 2 {
                                                                                        v19 = v19[1:len(v19)]
                                                                                    }
                                                                                case -2147483648:
                                                                                    v39 := GenS3{F0: v11, F1: uint(v2), F2: (v1 + v1)}
                                                                                    _ = v39
                                                                                default:
                                                                                    v19 = append(v19, fi15)
                                                                                }
                                                                                v16 = (!false)
                                                                            default:
                                                                                if len(v19) > 0 {
                                                                                    v19[0] = (fi17 << ((1) & 63))
                                                                                }
                                                                            }
                                                                            if v16 {
                                                                                _ = len(v19)
                                                                            } else {
                                                                                switch (fi17 + 127) {
                                                                                case -754:
                                                                                    v18 = 1e10
                                                                                case 42:
                                                                                    var v40 float64 = v3
                                                                                    _ = v40
                                                                                case 486:
                                                                                    v11 = v11
                                                                                }
                                                                            }
                                                                        }
                                                                    case 1:
                                                                        var v41 uint32 = (v11 / ((350) | 1))
                                                                        _ = v41
                                                                    case 127:
                                                                        switch v10 {
                                                                        case "foo bar":
                                                                            v11 += 1
                                                                        case "hello":
                                                                            v1 = "a"
                                                                        }
                                                                    }
                                                                    _ = v24
                                                                }
                                                            }
                                                        } else {
                                                            if (-32768 == 874) {
                                                                var v42 uint8 = 2
                                                                _ = v42
                                                            }
                                                            switch v10 {
                                                            case "\n\t":
                                                                if (0 < -32768) {
                                                                    v43 := GenS1{F0: 381}
                                                                    _ = v43
                                                                } else {
                                                                    v44 := GenS2{F0: 0, F1: 54}
                                                                    var v45 bool = v16
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                                var v46 float32 = v18
                                                                _ = v46
                                                            case "a":
                                                                var v47 bool = (1.5 < ((v3 + v3) / 1.5))
                                                                _ = v47
                                                            case "foo bar":
                                                                v48 := GenS4{F0: 2}
                                                                _ = v48
                                                            }
                                                            if (2147483648 == (v11 / (((v11 | (v11 - v11))) | 1))) {
                                                                v10 = v1
                                                                var v49 bool = (2 == uint64(v11))
                                                                _ = v49
                                                            }
                                                        }
                                                        fmt.Println(v2)
                                                        v50 := GenS5{F0: -636, F1: -604, F2: int32(v11)}
                                                        _ = v50
                                                    }
                                                    v51 := GenS4{F0: v11}
                                                    _ = v51
                                                default:
                                                    fmt.Println(v2, v18)
                                                    fmt.Println(v18)
                                                    v19 = append(v19, 4294967295)
                                                }
                                                fmt.Println(v1)
                                                v1 = (v10 + "")
                                                _ = v18
                                                _ = v19
                                            }
                                            _ = fi17
                                        }
                                        var v52 bool = v16
                                        _ = v52
                                    } else {
                                        switch (((((v1 + ("hello" + "a")) + ("a" + v1)) + v10) + v10) + (v1 + ((v1 + v1) + "hello"))) {
                                        case "\n\t":
                                            for fi53 := 0; fi53 < 20; fi53++ {
                                                v11 -= (v11 | v11)
                                                if true {
                                                    switch "foo bar" {
                                                    case "":
                                                        switch "" {
                                                        case "a":
                                                            switch v10 {
                                                            case "\n\t":
                                                                switch (fi53 * (fi15 * fi53)) {
                                                                case 32768:
                                                                    v54 := GenS2{F0: genFn0(1, 509), F1: 2}
                                                                    var v55 uint32 = v11
                                                                    v56 := []int{(fi53 >> ((1) & 63))}
                                                                    switch v10 {
                                                                    case "a":
                                                                        v57 := GenS6{F0: 109, F1: (v55 & 2147483648), F2: (v2 >> ((genFn0(18446744073709551615, 2)) & 63))}
                                                                        _ = v57
                                                                    case "hello":
                                                                        v2 += v2
                                                                    case "\n\t":
                                                                        switch int(v2) {
                                                                        case -810:
                                                                            if v16 {
                                                                                fmt.Println(fi53)
                                                                            }
                                                                        case 127:
                                                                            v2 += v2
                                                                        default:
                                                                            v58 := GenS1{F0: 1000}
                                                                            _ = v58
                                                                        }
                                                                    default:
                                                                        if len(v56) > 0 {
                                                                            v56[0] = fi15
                                                                        }
                                                                    }
                                                                    var v59 int = fi15
                                                                    v3 -= v3
                                                                    _ = v54
                                                                    _ = v55
                                                                    _ = v56
                                                                    _ = v59
                                                                case 408:
                                                                    switch v1 {
                                                                    case "a":
                                                                        if true {
                                                                            var v60 float64 = v3
                                                                            _ = v60
                                                                        }
                                                                    case "hello":
                                                                        var v61 uint8 = 42
                                                                        _ = v61
                                                                    default:
                                                                        v11 += (v11 - 4294967295)
                                                                    }
                                                                }
                                                            }
                                                        case "hello":
                                                            v62 := GenS3{F0: v11, F1: genFn0(0, 4294967295), F2: "\n\t"}
                                                            v63 := GenS1{F0: 0}
                                                            if v16 {
                                                                var v64 string = "foo bar"
                                                                fmt.Println(v16)
                                                                if true {
                                                                    v65 := []int{fi53, fi15}
                                                                    _ = v65
                                                                } else {
                                                                    v10 = "foo bar"
                                                                    v16 = ((v11 & (v11 % (((v11 << ((1) & 63))) | 1))) < (^(v11)))
                                                                }
                                                                _ = v64
                                                            }
                                                            v10 = (v10 + "")
                                                            var v66 bool = false
                                                            _ = v62
                                                            _ = v63
                                                            _ = v66
                                                        default:
                                                            if (v16 || ((v1 + v1) != v10)) {
                                                                if v16 {
                                                                    switch (("hello" + ((v1 + ((v1 + ("hello" + v10)) + v10)) + (((((v10 + ((((v1 + v10) + v1) + v1) + "foo bar")) + "\n\t") + (v10 + "hello")) + v1) + v1))) + v1) {
                                                                    case "foo bar":
                                                                        var v67 uint64 = 1
                                                                        _ = v67
                                                                    case "hello":
                                                                        var v68 uint8 = 128
                                                                        _ = v68
                                                                    case "":
                                                                        v2 -= 2
                                                                    default:
                                                                        var v69 bool = true
                                                                        _ = v69
                                                                    }
                                                                    var v70 string = (v10 + v10)
                                                                    _ = v70
                                                                }
                                                                v71 := GenS7{F0: v1}
                                                                v72 := []int{}
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                            v73 := []int{fi15, fi15, (fi53 - (fi15 & fi15))}
                                                            if (!v16) {
                                                                if len(v73) > 0 {
                                                                    v73[0] = fi15
                                                                }
                                                            }
                                                            _ = v73
                                                        }
                                                    }
                                                    v1 = ("" + v10)
                                                    v74 := []int{(fi53 % (((fi15 + 128)) | 1)), (fi15 ^ (fi15 | fi53)), fi15}
                                                    _ = v74
                                                }
                                                var v75 uint8 = uint8(v2)
                                                _ = fi53
                                                _ = v75
                                            }
                                        case "foo bar":
                                            v76 := GenS8{F0: 1, F1: (0 > 42), F2: v11}
                                            _ = v76
                                        case "":
                                            v11 = 2147483648
                                        default:
                                            for fi77 := 0; fi77 < 6; fi77++ {
                                                switch fi15 {
                                                case -1:
                                                    fmt.Println(v2)
                                                default:
                                                    var v78 bool = ((-32768 >= 0) && (4294967295 > 0))
                                                    _ = v78
                                                }
                                                _ = fi77
                                            }
                                        }
                                        if false {
                                            switch v1 {
                                            case "foo bar":
                                                v79 := GenS7{F0: (v10 + "\n\t")}
                                                _ = v79
                                            }
                                            switch (fi15 | fi15) {
                                            case -2147483648:
                                                if ((true && ((fi15 + fi15) <= fi15)) && (genFn0(286, 673) <= 331)) {
                                                    var v80 float32 = 1e10
                                                    var v81 float64 = float64(v80)
                                                    _ = v80
                                                    _ = v81
                                                } else {
                                                    v82 := GenS9{F0: (0.1 != (((v3 * (v3 / v3)) - (v3 / float64(0.1))) - v3))}
                                                    switch "" {
                                                    case "":
                                                        v83 := []int{(fi15 / ((int(v2)) | 1)), fi15, 930}
                                                        _ = v83
                                                    }
                                                    _ = v82
                                                }
                                            default:
                                                for fi84 := 0; fi84 < 16; fi84++ {
                                                    switch v1 {
                                                    case "a":
                                                        v3 += 1.5
                                                    case "foo bar":
                                                        if ((fi15 | -1) < (fi84 + -19)) {
                                                            switch 57 {
                                                            case 0:
                                                                v85 := []int{fi15, fi15, -759, 4294967295}
                                                                _ = v85
                                                            case 127:
                                                                var v86 uint8 = 227
                                                                _ = v86
                                                            }
                                                            v87 := []int{(fi84 << ((42) & 63)), (fi84 * fi15), (fi15 + fi84)}
                                                            v88 := GenS10{F0: (^(v11)), F1: 4294967295}
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    default:
                                                        if (-128 == 100) {
                                                            fmt.Println(fi15, fi84, v11)
                                                            var v89 uint = 717
                                                            _ = v89
                                                        }
                                                    }
                                                    var v90 float32 = (0.1 - 0.1)
                                                    _ = fi84
                                                    _ = v90
                                                }
                                            }
                                        } else {
                                            for fi91 := 0; fi91 < 14; fi91++ {
                                                if v16 {
                                                    v92 := []int{}
                                                    if (((v16 || (141 <= 657)) && v16) && (true || v16)) {
                                                        v93 := GenS11{F0: 9223372036854775807, F1: -16}
                                                        var v94 float32 = (((float32((((float64(1.5) * ((0.1 / v3) - v3)) - v3) * float64(float32((-2.25 + (v3 * 3.14159)))))) * float32(v3)) / ((-2.25 / -2.25) + 0.0)) / 1.5)
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        v16 = false
                                                        fmt.Println(v1)
                                                        v11 += (^(v11))
                                                    }
                                                    _ = v92
                                                } else {
                                                    v1 = "\n\t"
                                                    var v95 string = ((v10 + v1) + (((("hello" + v1) + "a") + v1) + "hello"))
                                                    _ = v95
                                                }
                                                v96 := []int{fi91, (fi91 - (fi91 * 4294967295)), (-(fi15)), fi91}
                                                var v97 int64 = 418
                                                _ = fi91
                                                _ = v96
                                                _ = v97
                                            }
                                            v98 := []int{-354, fi15, fi15, fi15}
                                            v99 := GenS9{F0: (1 != 42)}
                                            _ = v98
                                            _ = v99
                                        }
                                    }
                                    fmt.Println(v10)
                                    _ = v16
                                } else {
                                    switch 393 {
                                    case 62:
                                        switch ("foo bar" + v10) {
                                        case "\n\t":
                                            if (v1 == ("\n\t" + (("hello" + v1) + v1))) {
                                                switch (("\n\t" + (v1 + "hello")) + ("a" + "\n\t")) {
                                                case "":
                                                    var v100 int = fi15
                                                    _ = v100
                                                default:
                                                    v101 := GenS12{F0: v11}
                                                    _ = v101
                                                }
                                            } else {
                                                fmt.Println(v2)
                                                v102 := GenS2{F0: 4294967295, F1: 128}
                                                _ = v102
                                            }
                                        }
                                    default:
                                        v10 = ""
                                    }
                                    fmt.Println(v11, v3)
                                    v1 = "a"
                                }
                            default:
                                v103 := GenS4{F0: v11}
                                _ = v103
                            }
                        }
                    case -179:
                        v104 := GenS11{F0: 9223372036854775807, F1: 127}
                        _ = v104
                    default:
                        if (false || (((!true) || (!(-44 < -128))) || (904 < 273))) {
                            if (true && false) {
                                v10 = v10
                                var v105 string = (v10 + ("foo bar" + v1))
                                _ = v105
                            } else {
                                v1 = (("hello" + v10) + "hello")
                                v2 += 32767
                            }
                            fmt.Println(v2, v11, fi15)
                            var v106 int16 = 0
                            _ = v106
                        }
                    }
                    var v107 float64 = float64(-2.25)
                    for fi108 := 0; fi108 < 17; fi108++ {
                        v109 := []int{fi108, (fi108 & fi108)}
                        v3 -= v3
                        v2 += (v2 | (v2 % (((^(v2))) | 1)))
                        _ = fi108
                        _ = v109
                    }
                    _ = fi15
                    _ = v107
                }
            }
            _ = v11
            _ = v12
        } else {
            var v110 bool = false
            switch v10 {
            case "":
                if true {
                    v111 := GenS10{F0: 1, F1: 4294967295}
                    v1 = ""
                    for fi112 := 0; fi112 < 17; fi112++ {
                        v1 = v10
                        _ = fi112
                    }
                    _ = v111
                } else {
                    var v113 float64 = 0.0
                    if true {
                        switch 0 {
                        case 0:
                            if v110 {
                                if v110 {
                                    v10 = v1
                                    fmt.Println(v110, v1, v2)
                                }
                                switch 32767 {
                                case -1:
                                    if v110 {
                                        v114 := []int{-397}
                                        v115 := GenS13{F0: -32768, F1: 255, F2: (v10 + "foo bar")}
                                        _ = v114
                                        _ = v115
                                    }
                                default:
                                    v113 += float64((3.14159 + ((3.14159 / 0.1) - (1e10 - 0.0))))
                                }
                            } else {
                                v2 = 523
                                var v116 float64 = v113
                                if true {
                                    var v117 bool = v110
                                    _ = v117
                                } else {
                                    fmt.Println(v10, v10)
                                    v118 := []int{844, 1, 9223372036854775807}
                                    _ = v118
                                }
                                _ = v116
                            }
                        case 2147483647:
                            var v119 int16 = -32768
                            _ = v119
                        case 339:
                            for fi120 := 0; fi120 < 5; fi120++ {
                                if ((127 > -1) || false) {
                                    v113 = v113
                                } else {
                                    fmt.Println(v2, v3)
                                }
                                v121 := []int{(-(fi120)), (fi120 << ((genFn0(18446744073709551615, 625)) & 63)), fi120, (fi120 | (fi120 + fi120))}
                                if (((!((v3 - v3) != float64((((0.0 - -2.25) - 0.0) / 0.1)))) || (42 != 42)) || v110) {
                                    _ = len(v121)
                                }
                                _ = fi120
                                _ = v121
                            }
                        }
                        var v122 int32 = 1
                        _ = v122
                    }
                    _ = v113
                }
            case "foo bar":
                for fi123 := 0; fi123 < 1; fi123++ {
                    if v110 {
                        var v124 bool = v110
                        for fi125 := 0; fi125 < 14; fi125++ {
                            switch 32768 {
                            case -329:
                                v126 := GenS10{F0: 42, F1: 2147483647}
                                _ = v126
                            case 2147483647:
                                if (false && (42 > 18446744073709551615)) {
                                    switch (-(fi123)) {
                                    case 257:
                                        v3 -= v3
                                    default:
                                        v110 = false
                                    }
                                    v127 := []int{fi125, (-(fi125)), (fi125 * fi125), -9223372036854775808}
                                    var v128 uint32 = 2
                                    _ = v127
                                    _ = v128
                                } else {
                                    v129 := GenS14{F0: 4294967295, F1: 4294967295}
                                    if (((-2.25 * -2.25) / ((float32(v3) / 0.1) * 1.5)) != 1e10) {
                                        switch fi125 {
                                        case -545:
                                            var v130 uint32 = 664
                                            _ = v130
                                        default:
                                            if true {
                                                var v131 uint16 = (^(v2))
                                                v124 = false
                                                v1 = v1
                                                v132 := []int{fi123, 32767, fi125}
                                                v133 := GenS9{F0: ((v1 + ("" + v10)) >= "foo bar")}
                                                v132 = append(v132, fi123)
                                                v2 -= (^(v2))
                                                _ = v131
                                                _ = v132
                                                _ = v133
                                            } else {
                                                v3 += v3
                                                v124 = (!v110)
                                            }
                                        }
                                        fmt.Println(v124)
                                        v134 := GenS8{F0: -310, F1: (0 <= 450), F2: 42}
                                        _ = v134
                                    } else {
                                        v135 := GenS3{F0: 2147483647, F1: 710, F2: ((v1 + (((("\n\t" + (v10 + ((((("foo bar" + v1) + (v1 + v10)) + "hello") + v10) + v1))) + v1) + ((v10 + (("" + v1) + v1)) + v10)) + "")) + "a")}
                                        _ = v135
                                    }
                                    _ = v129
                                }
                            default:
                                var v136 bool = (42 > 32767)
                                _ = v136
                            }
                            fmt.Println(v124, v10)
                            v137 := []int{(fi125 >> ((286) & 63)), 128, 4294967295, fi125}
                            _ = fi125
                            _ = v137
                        }
                        _ = v124
                    }
                    v138 := GenS11{F0: 1, F1: -128}
                    var v139 bool = false
                    _ = fi123
                    _ = v138
                    _ = v139
                }
            default:
                v140 := []int{}
                _ = v140
            }
            _ = v110
        }
        fmt.Println(v10)
    }
    v141 := []int{32768}
    v10 = ((v10 + (v1 + v1)) + v10)
    fmt.Println(v1, v3, v1)
    fmt.Println(v1, v10, v1)
    for fi142 := 0; fi142 < 11; fi142++ {
        v143 := GenS15{F0: 4294967295, F1: 1}
        v1 = (v1 + ("\n\t" + v1))
        _ = fi142
        _ = v143
    }
    switch v1 {
    case "a":
        v144 := GenS16{F0: 32767, F1: 334}
        _ = v144
    case "foo bar":
        var v145 float64 = v3
        _ = v145
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v141:", len(v141), v141)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v141
    _ = v2
    _ = v3
}
