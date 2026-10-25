package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint64; F1 bool; F2 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int32; F1 int; F2 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int; F1 uint16; F2 string }
type GenS13 struct { F0 uint8; F1 uint64 }
type GenS14 struct { F0 bool; F1 string; F2 uint8 }
type GenS15 struct { F0 int8; F1 int32; F2 bool }
type GenS16 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{42, 65535, 538}
    var v3 float32 = (0.1 * (((0.1 / ((-2.25 + 1e10) * (1.5 / (1.5 / float32((0.1 - 3.14159)))))) + 0.1) * -2.25))
    var v4 uint8 = 216
    v5 := []string{"hello", "foo bar", ("foo bar" + "a"), "foo bar"}
    v4 += 163
    v6 := v1[36]
    v3 += float32(1e10)
    v3 -= ((v3 / v3) - v3)
    if v6 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v7 := GenS1{F0: 182, F1: 0}
        v4 -= 1
        _ = v7
    }
    if false {
        v6 = true
    } else {
        if false {
            switch 560 {
            case 0:
                if (2 > 248) {
                    if (!(!(!((("a" + "hello") < ("foo bar" + ("hello" + ""))) || (32768 != -806))))) {
                        switch 42 {
                        case -2147483648:
                            for fi8 := 0; fi8 < 4; fi8++ {
                                var v9 uint32 = 916
                                _ = fi8
                                _ = v9
                            }
                        case 32768:
                            switch "" {
                            case "\n\t":
                                if v6 {
                                    if v6 {
                                        switch 42 {
                                        case 127:
                                            v10 := GenS1{F0: 2147483647, F1: -2147483648}
                                            _ = v10
                                        case 0:
                                            switch ("" + "") {
                                            case "":
                                                for fi11 := 0; fi11 < 20; fi11++ {
                                                    for fi12 := 0; fi12 < 14; fi12++ {
                                                        v6 = true
                                                        switch fi11 {
                                                        case 0:
                                                            switch (fi11 & 32768) {
                                                            case 128:
                                                                if (1 <= 723) {
                                                                    v6 = (("foo bar" + "") > "foo bar")
                                                                    if (65535 > 158) {
                                                                        fmt.Println(fi12)
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v6 = true
                                                                    case "foo bar":
                                                                        v13 := GenS2{F0: "\n\t"}
                                                                        _ = v13
                                                                    case "\n\t":
                                                                        v14 := GenS3{F0: (!(v6 || (65535 < 32767))), F1: 32767, F2: 459}
                                                                        _ = v14
                                                                    }
                                                                    switch fi12 {
                                                                    case -9223372036854775808:
                                                                        v6 = false
                                                                    }
                                                                    v6 = ((^(v4)) > (v4 >> ((18446744073709551615) & 63)))
                                                                    v6 = v6
                                                                    var v15 uint16 = 65535
                                                                    v5 = append(v5, "a")
                                                                    var v16 float64 = 0.0
                                                                    if (2 <= v4) {
                                                                        fmt.Println(v15)
                                                                    }
                                                                    if false {
                                                                        v17 := GenS3{F0: ("hello" >= (("" + (((("\n\t" + ("\n\t" + "a")) + "") + ("\n\t" + "foo bar")) + ("\n\t" + "hello"))) + "\n\t")), F1: 977, F2: 724}
                                                                        switch fi12 {
                                                                        case -1:
                                                                            v18 := GenS4{F0: 596}
                                                                            _ = v18
                                                                        case 1:
                                                                            switch (("hello" + "a") + "\n\t") {
                                                                            case "hello":
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            case "":
                                                                                fmt.Println(v4, fi11, v4)
                                                                            case "\n\t":
                                                                                if ((!true) || (!v6)) {
                                                                                    var v19 int16 = 42
                                                                                    v1[1] = ((1.5 == 1e10) || v6)
                                                                                    v20 := GenS5{F0: v4, F1: -71, F2: 2}
                                                                                    _ = v19
                                                                                    _ = v20
                                                                                }
                                                                            default:
                                                                                var v21 bool = true
                                                                                _ = v21
                                                                            }
                                                                            if (0 < -9223372036854775808) {
                                                                                if false {
                                                                                    var v22 uint64 = 891
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    _ = v22
                                                                                } else {
                                                                                    var v23 uint8 = 70
                                                                                    switch (-(fi11)) {
                                                                                    case 893:
                                                                                        switch "" {
                                                                                        case "":
                                                                                            delete(v1, 2)
                                                                                        case "\n\t":
                                                                                            v24 := v1[0]
                                                                                            _ = v24
                                                                                        }
                                                                                        if true {
                                                                                            var v25 int = (fi11 >> ((42) & 63))
                                                                                            var v26 int64 = -1
                                                                                            var v27 uint64 = 4294967295
                                                                                            var v28 uint32 = 582
                                                                                            var v29 float32 = float32((v16 + -2.25))
                                                                                            fmt.Println(v23, v25, v16)
                                                                                            switch (-(v25)) {
                                                                                            case 9223372036854775807:
                                                                                                var v30 int32 = -1
                                                                                                _ = v30
                                                                                            default:
                                                                                                if true {
                                                                                                    if (42 < 1) {
                                                                                                        var v31 float32 = float32(v16)
                                                                                                        if (!v6) {
                                                                                                            v3 = float32(v16)
                                                                                                            if (((float64(v3) / v16) < 0.1) && v6) {
                                                                                                                fmt.Println(v29, v28)
                                                                                                            } else {
                                                                                                                switch "\n\t" {
                                                                                                                case "a":
                                                                                                                    if (((1e10 - 0.1) * ((0.1 + -2.25) * v29)) == v29) {
                                                                                                                        switch (-(fi11)) {
                                                                                                                        case 432:
                                                                                                                            fmt.Println(fi11, v28)
                                                                                                                        case -1:
                                                                                                                            v32 := GenS4{F0: (v28 | v28)}
                                                                                                                            _ = v32
                                                                                                                        case 2:
                                                                                                                            var v33 int8 = 114
                                                                                                                            _ = v33
                                                                                                                        default:
                                                                                                                            v25 -= -2147483648
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v5 = append(v5, "")
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    var v34 int16 = 504
                                                                                                                    _ = v34
                                                                                                                default:
                                                                                                                    var v35 uint16 = 2
                                                                                                                    _ = v35
                                                                                                                }
                                                                                                            }
                                                                                                            if len(v5) > 0 {
                                                                                                                v5[0] = ("" + "foo bar")
                                                                                                            }
                                                                                                        } else {
                                                                                                            v29 = 1e10
                                                                                                            if false {
                                                                                                                var v36 int = 2147483647
                                                                                                                var v37 int64 = 42
                                                                                                                if (255 != uint(v27)) {
                                                                                                                    v27 = v27
                                                                                                                    var v38 float32 = ((((0.0 - (0.1 * ((float32((v16 - (((3.14159 - v16) * (0.0 * ((3.14159 + (1e10 * -2.25)) + 0.0))) / (float64((0.0 + v31)) - (v16 - float64(((-2.25 / v31) + (v31 - v3)))))))) * float32(float64(((float32(v16) - v3) * (v31 * 0.1))))) - 0.0))) / v3) - v29) + v29)
                                                                                                                    var v39 int64 = (v26 + v37)
                                                                                                                    _ = v38
                                                                                                                    _ = v39
                                                                                                                } else {
                                                                                                                    var v40 uint64 = (^(v27))
                                                                                                                    if (v6 || v6) {
                                                                                                                        if v6 {
                                                                                                                            if v6 {
                                                                                                                                v41 := v1[v28]
                                                                                                                                v16 = v16
                                                                                                                                var v42 float64 = v16
                                                                                                                                _ = v41
                                                                                                                                _ = v42
                                                                                                                            }
                                                                                                                            v43 := GenS6{F0: v27, F1: v6, F2: v37}
                                                                                                                            var v44 uint8 = v23
                                                                                                                            _ = v43
                                                                                                                            _ = v44
                                                                                                                        } else {
                                                                                                                            if (float32(float64(-2.25)) == 1e10) {
                                                                                                                                if false {
                                                                                                                                    var v45 float64 = ((v16 - (0.0 / 1e10)) + ((v16 * (v16 + float64((v31 * v29)))) + (v16 - 1e10)))
                                                                                                                                    v46 := GenS5{F0: uint8(v28), F1: -56, F2: 20}
                                                                                                                                    v47 := GenS3{F0: (654 != -1), F1: 1, F2: v15}
                                                                                                                                    var v48 int32 = 42
                                                                                                                                    var v49 float32 = (1e10 * (float32(0.1) / v31))
                                                                                                                                    _ = v45
                                                                                                                                    _ = v46
                                                                                                                                    _ = v47
                                                                                                                                    _ = v48
                                                                                                                                    _ = v49
                                                                                                                                } else {
                                                                                                                                    _ = len(v1)
                                                                                                                                    fmt.Println(v23)
                                                                                                                                    v40 += (v40 & 42)
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v50 int64 = 9223372036854775807
                                                                                                                                _ = v50
                                                                                                                            }
                                                                                                                            switch v36 {
                                                                                                                            case 9223372036854775807:
                                                                                                                                fmt.Println(v37, v6)
                                                                                                                            case 128:
                                                                                                                                fmt.Println(v40)
                                                                                                                            }
                                                                                                                            var v51 float32 = 1e10
                                                                                                                            _ = v51
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v25 += 178
                                                                                                                    fmt.Println(v36, v26)
                                                                                                                    _ = v40
                                                                                                                }
                                                                                                                _ = v36
                                                                                                                _ = v37
                                                                                                            } else {
                                                                                                                var v52 bool = v6
                                                                                                                _ = v52
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v31
                                                                                                    }
                                                                                                    var v53 int8 = -1
                                                                                                    _ = v53
                                                                                                }
                                                                                            }
                                                                                            if v6 {
                                                                                                if false {
                                                                                                    v54 := GenS6{F0: (^(v27)), F1: false, F2: (v26 / ((-9223372036854775808) | 1))}
                                                                                                    if v6 {
                                                                                                        if (v29 != v3) {
                                                                                                            if ((((v28 << ((65535) & 63)) >= v28) || ((v4 % (((^(v4))) | 1)) >= v4)) || (470 < v15)) {
                                                                                                                var v55 uint32 = (v28 + v28)
                                                                                                                if v6 {
                                                                                                                    v56 := GenS4{F0: 970}
                                                                                                                    if (uint32(v15) == v28) {
                                                                                                                        var v57 int32 = 42
                                                                                                                        v58 := GenS5{F0: 127, F1: 42, F2: 2147483648}
                                                                                                                        var v59 int8 = 42
                                                                                                                        fmt.Println(v4)
                                                                                                                        v59 = (-(v59))
                                                                                                                        _ = v57
                                                                                                                        _ = v58
                                                                                                                        _ = v59
                                                                                                                    } else {
                                                                                                                        var v60 float32 = 3.14159
                                                                                                                        if v6 {
                                                                                                                            delete(v1, v55)
                                                                                                                            fmt.Println(v26)
                                                                                                                            if ((v6 && v6) || v6) {
                                                                                                                                v29 -= (-2.25 / v3)
                                                                                                                                var v61 uint16 = uint16(v4)
                                                                                                                                switch fi11 {
                                                                                                                                case 9223372036854775807:
                                                                                                                                    v62 := GenS5{F0: (^(v4)), F1: 0, F2: (v28 >> ((uint(v4)) & 63))}
                                                                                                                                    _ = v62
                                                                                                                                case 1:
                                                                                                                                    var v63 float32 = ((3.14159 / v3) + v29)
                                                                                                                                    _ = v63
                                                                                                                                case 0:
                                                                                                                                    if ((1e10 > v60) && (true && v6)) {
                                                                                                                                        v4 -= uint8(v55)
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    v64 := GenS7{F0: 655}
                                                                                                                                    _ = v64
                                                                                                                                }
                                                                                                                                _ = v61
                                                                                                                            } else {
                                                                                                                                if v6 {
                                                                                                                                    v65 := GenS5{F0: v4, F1: -18, F2: 1}
                                                                                                                                    v28 -= 629
                                                                                                                                    v66 := GenS8{F0: (v25 / ((4294967295) | 1))}
                                                                                                                                    _ = v65
                                                                                                                                    _ = v66
                                                                                                                                } else {
                                                                                                                                    fmt.Println(fi11, fi11)
                                                                                                                                    v60 += (((v60 * ((v3 * v29) + v29)) + v60) * v29)
                                                                                                                                    fmt.Println(v60)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            var v67 int = 4294967295
                                                                                                                            fmt.Println(v28, v28)
                                                                                                                            var v68 int16 = -32768
                                                                                                                            if ((v16 < (1.5 + (float64((v29 * v60)) / (float64(v60) - 1e10)))) || v6) {
                                                                                                                                if (v16 >= (v16 / 0.1)) {
                                                                                                                                    var v69 int8 = 42
                                                                                                                                    var v70 string = "\n\t"
                                                                                                                                    _ = v69
                                                                                                                                    _ = v70
                                                                                                                                } else {
                                                                                                                                    if (!((-128 > -35) || (fi11 == (-(v67))))) {
                                                                                                                                        v15 = (v15 | v15)
                                                                                                                                    }
                                                                                                                                    v15 = v15
                                                                                                                                    v71 := v1[v28]
                                                                                                                                    _ = v71
                                                                                                                                }
                                                                                                                                v72 := GenS7{F0: uint(v15)}
                                                                                                                                v29 += (v60 + (v60 + 1.5))
                                                                                                                                _ = v72
                                                                                                                            } else {
                                                                                                                                delete(v1, (v28 | 2147483647))
                                                                                                                                v73 := GenS4{F0: (v55 ^ 854)}
                                                                                                                                if (1 < -2147483648) {
                                                                                                                                    if true {
                                                                                                                                        switch (v67 - fi11) {
                                                                                                                                        case 32767:
                                                                                                                                            var v74 int32 = 2147483647
                                                                                                                                            _ = v74
                                                                                                                                        case 128:
                                                                                                                                            var v75 int8 = 0
                                                                                                                                            _ = v75
                                                                                                                                        }
                                                                                                                                        var v76 bool = (v26 < 9223372036854775807)
                                                                                                                                        _ = v76
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if v6 {
                                                                                                                                        v16 += float64(v29)
                                                                                                                                        v26 = v26
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v15, v29)
                                                                                                                                    }
                                                                                                                                    fmt.Println(v25, v23, v25)
                                                                                                                                }
                                                                                                                                var v77 uint = 0
                                                                                                                                fmt.Println(v28, v3, v16)
                                                                                                                                _ = v73
                                                                                                                                _ = v77
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                v78 := GenS9{F0: 53, F1: v15}
                                                                                                                                _ = v78
                                                                                                                            } else {
                                                                                                                                if (2 != 600) {
                                                                                                                                    v55 = v55
                                                                                                                                    if v6 {
                                                                                                                                        switch "" {
                                                                                                                                        case "":
                                                                                                                                            switch "hello" {
                                                                                                                                            case "hello":
                                                                                                                                                v79 := GenS9{F0: 716, F1: v15}
                                                                                                                                                _ = v79
                                                                                                                                            case "":
                                                                                                                                                if false {
                                                                                                                                                    delete(v1, v28)
                                                                                                                                                } else {
                                                                                                                                                    fmt.Println(v25)
                                                                                                                                                    fmt.Println(v23)
                                                                                                                                                    fmt.Println(v29)
                                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        v16 = v16
                                                                                                                                        if false {
                                                                                                                                            var v80 bool = v6
                                                                                                                                            var v81 int32 = 2147483647
                                                                                                                                            if (!v6) {
                                                                                                                                                var v82 bool = (!(uint(v4) > 18446744073709551615))
                                                                                                                                                var v83 int64 = (v26 / (((v26 - (v26 / (((v26 >> ((65535) & 63))) | 1)))) | 1))
                                                                                                                                                _ = v82
                                                                                                                                                _ = v83
                                                                                                                                            } else {
                                                                                                                                                v84 := GenS9{F0: v26, F1: 2}
                                                                                                                                                switch v67 {
                                                                                                                                                case 2147483647:
                                                                                                                                                    v6 = v6
                                                                                                                                                case -430:
                                                                                                                                                    v80 = true
                                                                                                                                                case -903:
                                                                                                                                                    var v85 int16 = int16(v55)
                                                                                                                                                    _ = v85
                                                                                                                                                default:
                                                                                                                                                    var v86 int8 = -128
                                                                                                                                                    _ = v86
                                                                                                                                                }
                                                                                                                                                _ = v84
                                                                                                                                            }
                                                                                                                                            _ = v80
                                                                                                                                            _ = v81
                                                                                                                                        } else {
                                                                                                                                            v87 := GenS1{F0: -1, F1: 1}
                                                                                                                                            switch ((("hello" + "") + (((("" + "") + ("" + "foo bar")) + ("a" + "foo bar")) + "")) + "foo bar") {
                                                                                                                                            case "foo bar":
                                                                                                                                                _ = len(v1)
                                                                                                                                            default:
                                                                                                                                                v16 = ((v16 * 1.5) + ((1.5 + -2.25) * 0.1))
                                                                                                                                            }
                                                                                                                                            _ = v87
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    var v88 uint = 255
                                                                                                                                    _ = v88
                                                                                                                                }
                                                                                                                                var v89 int = int(v55)
                                                                                                                                v90 := GenS5{F0: 0, F1: -13, F2: 863}
                                                                                                                                _ = v89
                                                                                                                                _ = v90
                                                                                                                            }
                                                                                                                            _ = v67
                                                                                                                            _ = v68
                                                                                                                        }
                                                                                                                        _ = len(v1)
                                                                                                                        _ = v60
                                                                                                                    }
                                                                                                                    fmt.Println(v25, v3, v29)
                                                                                                                    v91 := GenS10{F0: 1, F1: fi11, F2: (v55 / ((v55) | 1))}
                                                                                                                    _ = v56
                                                                                                                    _ = v91
                                                                                                                } else {
                                                                                                                    var v92 int8 = 42
                                                                                                                    fmt.Println(v92, v28, v23)
                                                                                                                    fmt.Println(fi11, v4, v55)
                                                                                                                    _ = v92
                                                                                                                }
                                                                                                                _ = v55
                                                                                                            } else {
                                                                                                                v29 = 1e10
                                                                                                                switch "" {
                                                                                                                case "hello":
                                                                                                                    v93 := GenS9{F0: int64(v15), F1: v15}
                                                                                                                    _ = v93
                                                                                                                case "foo bar":
                                                                                                                    if false {
                                                                                                                        var v94 int = 304
                                                                                                                        switch "foo bar" {
                                                                                                                        case "":
                                                                                                                            switch -190 {
                                                                                                                            case 9223372036854775807:
                                                                                                                                var v95 float32 = ((0.1 + v29) * (-2.25 - v29))
                                                                                                                                _ = v95
                                                                                                                            case 128:
                                                                                                                                v27 -= 42
                                                                                                                            }
                                                                                                                        case "\n\t":
                                                                                                                            var v96 float64 = v16
                                                                                                                            _ = v96
                                                                                                                        default:
                                                                                                                            var v97 float32 = (v29 + v3)
                                                                                                                            _ = v97
                                                                                                                        }
                                                                                                                        _ = v94
                                                                                                                    } else {
                                                                                                                        v98 := GenS3{F0: false, F1: (v15 << ((125) & 63)), F2: v15}
                                                                                                                        v28 -= v28
                                                                                                                        var v99 uint8 = 255
                                                                                                                        _ = v98
                                                                                                                        _ = v99
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if (!(-128 < 35)) {
                                                                                                                        switch "" {
                                                                                                                        case "hello":
                                                                                                                            v100 := v1[v28]
                                                                                                                            _ = v100
                                                                                                                        case "a":
                                                                                                                            v1[(v28 * (v28 / ((v28) | 1)))] = false
                                                                                                                            v25 += 0
                                                                                                                            if ((v6 || v6) || v6) {
                                                                                                                                var v101 uint = 18446744073709551615
                                                                                                                                _ = v101
                                                                                                                            }
                                                                                                                        case "foo bar":
                                                                                                                            v16 -= v16
                                                                                                                        default:
                                                                                                                            v1[uint32(v23)] = true
                                                                                                                        }
                                                                                                                        v16 = (float64(v3) / v16)
                                                                                                                        if (uint32(v15) <= 402) {
                                                                                                                            if (v27 > v27) {
                                                                                                                                v4 += v23
                                                                                                                                var v102 int16 = -32768
                                                                                                                                var v103 uint8 = 33
                                                                                                                                _ = v102
                                                                                                                                _ = v103
                                                                                                                            } else {
                                                                                                                                if (((!(v27 < (v27 & v27))) || (-1 != (fi12 << ((0) & 63)))) || v6) {
                                                                                                                                    v104 := v1[v28]
                                                                                                                                    switch ("" + "\n\t") {
                                                                                                                                    case "":
                                                                                                                                        v27 += 265
                                                                                                                                    case "hello":
                                                                                                                                        v26 = -337
                                                                                                                                    case "foo bar":
                                                                                                                                        var v105 float32 = v3
                                                                                                                                        _ = v105
                                                                                                                                    }
                                                                                                                                    _ = v104
                                                                                                                                }
                                                                                                                                fmt.Println(fi11)
                                                                                                                                v1[(v28 + (v28 | (^(v28))))] = v6
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                v106 := GenS6{F0: v27, F1: (v6 || (v16 <= (float64(v29) + -2.25))), F2: (-(v26))}
                                                                                                                _ = v106
                                                                                                            }
                                                                                                            if v6 {
                                                                                                                if ((^(v27)) != (v27 % ((18446744073709551615) | 1))) {
                                                                                                                    var v107 string = "a"
                                                                                                                    if v6 {
                                                                                                                        if v6 {
                                                                                                                            if true {
                                                                                                                                if v6 {
                                                                                                                                    v16 -= v16
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                fmt.Println(v27, v107)
                                                                                                                            }
                                                                                                                            if ((v15 % ((32767) | 1)) != 32767) {
                                                                                                                                v26 -= 0
                                                                                                                                switch v25 {
                                                                                                                                case 9223372036854775807:
                                                                                                                                    switch v107 {
                                                                                                                                    case "":
                                                                                                                                        v29 = v29
                                                                                                                                    }
                                                                                                                                case -1:
                                                                                                                                    v108 := GenS11{F0: 2147483647}
                                                                                                                                    _ = v108
                                                                                                                                case 127:
                                                                                                                                    fmt.Println(v23, fi12, v6)
                                                                                                                                default:
                                                                                                                                    v5 = append(v5, "")
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v109 int16 = -1
                                                                                                                                switch v107 {
                                                                                                                                case "a":
                                                                                                                                    v110 := GenS2{F0: v107}
                                                                                                                                    _ = v110
                                                                                                                                case "foo bar":
                                                                                                                                    v15 = v15
                                                                                                                                case "\n\t":
                                                                                                                                    var v111 int8 = -128
                                                                                                                                    _ = v111
                                                                                                                                }
                                                                                                                                _ = v109
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v112 := v1[uint32(v4)]
                                                                                                                        v1[v28] = false
                                                                                                                        _ = v112
                                                                                                                    }
                                                                                                                    fmt.Println(v15, v28)
                                                                                                                    v113 := GenS2{F0: v107}
                                                                                                                    var v114 int16 = -203
                                                                                                                    v115 := GenS6{F0: v27, F1: v6, F2: v26}
                                                                                                                    _ = v107
                                                                                                                    _ = v113
                                                                                                                    _ = v114
                                                                                                                    _ = v115
                                                                                                                }
                                                                                                                var v116 int = 32767
                                                                                                                v117 := GenS4{F0: v28}
                                                                                                                _ = v116
                                                                                                                _ = v117
                                                                                                            } else {
                                                                                                                var v118 uint8 = v23
                                                                                                                if (-2147483648 < 2147483647) {
                                                                                                                    fmt.Println(fi11, v27, v28)
                                                                                                                }
                                                                                                                _ = v118
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch "a" {
                                                                                                            case "":
                                                                                                                var v119 uint = 255
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            delete(v1, uint32(v15))
                                                                                                        }
                                                                                                        v120 := v1[v28]
                                                                                                        _ = len(v1)
                                                                                                        _ = v120
                                                                                                    } else {
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            var v121 float32 = v29
                                                                                                            _ = v121
                                                                                                        case "foo bar":
                                                                                                            v25 = (v25 % ((fi11) | 1))
                                                                                                        default:
                                                                                                            switch (("foo bar" + "\n\t") + "\n\t") {
                                                                                                            case "foo bar":
                                                                                                                if true {
                                                                                                                    if v6 {
                                                                                                                        var v122 float64 = (0.0 - v16)
                                                                                                                        var v123 string = (((("hello" + ("\n\t" + (("a" + ("a" + ("a" + "\n\t"))) + "foo bar"))) + "foo bar") + (("" + "a") + "a")) + "foo bar")
                                                                                                                        _ = v122
                                                                                                                        _ = v123
                                                                                                                    } else {
                                                                                                                        if v6 {
                                                                                                                            fmt.Println(fi12, v23, v6)
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case "":
                                                                                                                v4 -= (v4 * uint8(v27))
                                                                                                                delete(v1, v28)
                                                                                                                if (!false) {
                                                                                                                    v124 := GenS8{F0: (fi11 / (((fi11 ^ (-(v25)))) | 1))}
                                                                                                                    v125 := v1[(^(v28))]
                                                                                                                    switch (fi11 & 127) {
                                                                                                                    case -1:
                                                                                                                        v126 := GenS6{F0: uint64(v23), F1: v125, F2: (-(v26))}
                                                                                                                        var v127 uint16 = v15
                                                                                                                        fmt.Println(fi11)
                                                                                                                        var v128 bool = v125
                                                                                                                        v3 = 0.0
                                                                                                                        _ = v126
                                                                                                                        _ = v127
                                                                                                                        _ = v128
                                                                                                                    case -292:
                                                                                                                        _ = len(v2)
                                                                                                                        var v129 float64 = v16
                                                                                                                        _ = v129
                                                                                                                    default:
                                                                                                                        delete(v1, (v28 % (((v28 ^ (v28 * v28))) | 1)))
                                                                                                                    }
                                                                                                                    _ = v124
                                                                                                                    _ = v125
                                                                                                                }
                                                                                                                v3 += (((((v3 / 1.0) * v29) / (((1e10 - ((v29 - v29) + 3.14159)) + float32(((v16 - 3.14159) + 1e10))) * v3)) - (3.14159 - v29)) - (v3 - 0.1))
                                                                                                                var v130 bool = (((v16 / v16) - (v16 + float64(v29))) != 0.0)
                                                                                                                _ = v130
                                                                                                            case "\n\t":
                                                                                                                v3 = ((v29 / (v3 * ((v29 / 1e10) + 3.14159))) + 0.1)
                                                                                                            }
                                                                                                        }
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            v25 = 1
                                                                                                        }
                                                                                                    }
                                                                                                    v131 := GenS1{F0: -598, F1: 0}
                                                                                                    _ = v131
                                                                                                    _ = v54
                                                                                                } else {
                                                                                                    switch "hello" {
                                                                                                    case "foo bar":
                                                                                                        var v132 uint16 = (v15 - v15)
                                                                                                        _ = v132
                                                                                                    case "":
                                                                                                        fmt.Println(fi12)
                                                                                                    case "a":
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            var v133 uint = 1
                                                                                                            _ = v133
                                                                                                        case "":
                                                                                                            if (false && v6) {
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = (v28 << ((886) & 63))
                                                                                                                }
                                                                                                                v27 = v27
                                                                                                            } else {
                                                                                                                var v134 uint8 = 165
                                                                                                                _ = v134
                                                                                                            }
                                                                                                        default:
                                                                                                            if ("foo bar" < (("\n\t" + "foo bar") + "\n\t")) {
                                                                                                                var v135 string = "a"
                                                                                                                fmt.Println(v15, v135, v27)
                                                                                                                fmt.Println(v29)
                                                                                                                _ = v135
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v136 := GenS3{F0: v6, F1: (^(v15)), F2: (v15 & v15)}
                                                                                                v5 = append(v5, "hello")
                                                                                                _ = v136
                                                                                            }
                                                                                            _ = v25
                                                                                            _ = v26
                                                                                            _ = v27
                                                                                            _ = v28
                                                                                            _ = v29
                                                                                        }
                                                                                    default:
                                                                                        var v137 int16 = 32767
                                                                                        _ = v137
                                                                                    }
                                                                                    v16 -= float64(((0.0 - ((1.5 + v3) - (0.0 / v3))) / 3.14159))
                                                                                    _ = v23
                                                                                }
                                                                                fmt.Println(fi12)
                                                                                fmt.Println(v16)
                                                                            }
                                                                        }
                                                                        v6 = (2 > v15)
                                                                        _ = v17
                                                                    } else {
                                                                        if (!(6 != 4294967295)) {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = ("" + "")
                                                                            }
                                                                            if true {
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    delete(v1, 0)
                                                                                default:
                                                                                    var v138 uint16 = (v15 >> ((255) & 63))
                                                                                    _ = v138
                                                                                }
                                                                                v3 = v3
                                                                                v3 = v3
                                                                            } else {
                                                                                var v139 int64 = -1
                                                                                v140 := GenS1{F0: -354, F1: 328}
                                                                                v16 += (v16 + (float64(v3) * v16))
                                                                                _ = v139
                                                                                _ = v140
                                                                            }
                                                                        }
                                                                        var v141 int64 = -9223372036854775808
                                                                        var v142 uint16 = v15
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                    _ = v15
                                                                    _ = v16
                                                                } else {
                                                                    v3 += v3
                                                                    fmt.Println(v6, v3)
                                                                    var v143 uint16 = 65535
                                                                    _ = v143
                                                                }
                                                            case -9223372036854775808:
                                                                if v6 {
                                                                    v6 = v6
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v144 := GenS12{F0: (-(fi12)), F1: 0, F2: "\n\t"}
                                                                    _ = v144
                                                                } else {
                                                                    v4 += 128
                                                                    v145 := v1[1]
                                                                    v4 = 255
                                                                    _ = v145
                                                                }
                                                            }
                                                        case -306:
                                                            delete(v1, 2)
                                                        case 127:
                                                            if false {
                                                                v6 = true
                                                                var v146 uint32 = 842
                                                                var v147 uint16 = 32767
                                                                _ = v146
                                                                _ = v147
                                                            }
                                                        default:
                                                            v6 = (0 == 1)
                                                        }
                                                        _ = fi12
                                                    }
                                                    _ = fi11
                                                }
                                            case "foo bar":
                                                fmt.Println(v3, v3, v3)
                                            }
                                        default:
                                            v3 = v3
                                        }
                                        var v148 uint64 = 2
                                        _ = v148
                                    }
                                    v4 = v4
                                    var v149 int16 = 1000
                                    _ = v149
                                } else {
                                    v4 = (v4 << ((958) & 63))
                                    var v150 int16 = -635
                                    _ = v150
                                }
                            }
                        }
                    }
                }
            case 695:
                v4 += (v4 + (v4 & v4))
            case -330:
                if (-1 < -1) {
                    v151 := GenS13{F0: v4, F1: 18446744073709551615}
                    for fi152 := 0; fi152 < 13; fi152++ {
                        if (-2.25 != 0.0) {
                            delete(v1, 98)
                            v153 := GenS14{F0: (!true), F1: "a", F2: 118}
                            _ = v153
                        }
                        _ = fi152
                    }
                    _ = v151
                }
            }
        } else {
            for fi154 := 0; fi154 < 1; fi154++ {
                v155 := GenS15{F0: 100, F1: 1, F2: v6}
                if (42 != 2) {
                    if len(v5) > 0 {
                        v5[0] = ((("\n\t" + (("a" + "hello") + "\n\t")) + "\n\t") + "\n\t")
                    }
                    switch 9223372036854775807 {
                    case -997:
                        delete(v1, 65535)
                    }
                    var v156 int64 = -1
                    _ = v156
                }
                _ = fi154
                _ = v155
            }
            v157 := GenS7{F0: 1}
            _ = v157
        }
        v6 = false
        var v158 string = ((("hello" + "") + (((("hello" + "hello") + "\n\t") + "\n\t") + ("a" + ""))) + "a")
        _ = v158
    }
    var v159 uint64 = 0
    for fi160 := 0; fi160 < 20; fi160++ {
        v4 = 1
        switch 116 {
        case 224:
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        default:
            if (false || v6) {
                var v161 float32 = 1.5
                if false {
                    if (1 >= 1) {
                        var v162 int16 = 1000
                        _ = v162
                    }
                    for fi163 := 0; fi163 < 16; fi163++ {
                        v3 = ((v161 + v3) - (-2.25 + v161))
                        _ = fi163
                    }
                    v3 += v3
                } else {
                    for fi164 := 0; fi164 < 6; fi164++ {
                        v4 = (v4 & 42)
                        v161 = v3
                        v165 := GenS16{F0: ((fi160 | fi164) >= -954), F1: -18}
                        _ = fi164
                        _ = v165
                    }
                    v6 = v6
                }
                _ = v161
            } else {
                var v166 string = ("\n\t" + ((("foo bar" + "foo bar") + "hello") + ""))
                v1[4294967295] = v6
                _ = v166
            }
        }
        _ = fi160
    }
    fmt.Println(v159)
    if (1 == 42) {
        v3 -= (v3 + 0.0)
    } else {
        for fi167 := 0; fi167 < 15; fi167++ {
            v168 := GenS7{F0: 871}
            switch fi167 {
            case -872:
                fmt.Println(v4)
            }
            _ = fi167
            _ = v168
        }
    }
    _ = len(v5)
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v159:", v159)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v159
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
