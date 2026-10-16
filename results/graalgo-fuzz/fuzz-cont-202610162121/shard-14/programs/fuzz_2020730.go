package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 string }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 string; F1 int64; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 uint32; F2 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 bool; F2 int32 }
type GenS11 struct { F0 int32; F1 uint; F2 uint32 }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 uint; F1 int64; F2 uint }
type GenS14 struct { F0 int8; F1 int8; F2 int8 }
type GenS15 struct { F0 int32; F1 uint16; F2 bool }
type GenS16 struct { F0 uint; F1 string; F2 uint }
type GenS17 struct { F0 uint; F1 int64; F2 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 int16 = 42
    v6 := make(map[int]uint32)
    v2 = ("\n\t" + (v2 + v2))
    var v7 int8 = 100
    v1 = v1
    v1 += v1
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            if true {
                v5 -= v5
                v10 := GenS1{F0: ("\n\t" + ("" + "hello")), F1: 457, F2: ((v2 + ("\n\t" + "\n\t")) + "")}
                v7 = 0
                _ = v10
            }
            _ = fi9
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            v2 = "hello"
            _ = fi11
        }
        _ = fi8
    }
    switch 32767 {
    case 128:
        v12 := GenS1{F0: "hello", F1: 1, F2: (v2 + v2)}
        _ = v12
    case 32768:
        v7 = -1
    default:
        v13 := v6[-1]
        _ = v13
    }
    switch 2147483647 {
    case -1:
        for fi14 := 0; fi14 < 9; fi14++ {
            if (628 == 42) {
                if true {
                    if (!true) {
                        fmt.Println(v1, v5)
                    } else {
                        switch ("a" + "\n\t") {
                        case "a":
                            switch fi14 {
                            case 127:
                                delete(v6, (fi14 << ((4294967295) & 63)))
                            case 2147483647:
                                v15 := GenS1{F0: (v2 + v2), F1: 1, F2: "a"}
                                _ = v15
                            case 4294967295:
                                v6[(^(fi14))] = 71
                            }
                        case "hello":
                            v1 -= v1
                        case "foo bar":
                            var v16 int64 = -183
                            _ = v16
                        default:
                            for fi17 := 0; fi17 < 15; fi17++ {
                                if (599 != 42) {
                                    var v18 float32 = (3.14159 * 0.0)
                                    _ = v18
                                } else {
                                    if true {
                                        switch 2147483647 {
                                        case -1:
                                            v7 -= (v7 >> ((4294967295) & 63))
                                        case 0:
                                            switch (fi14 ^ -565) {
                                            case -2147483648:
                                                v5 -= v5
                                            }
                                        case -9223372036854775808:
                                            var v19 uint16 = v1
                                            _ = v19
                                        default:
                                            v20 := []int{}
                                            var v21 int64 = -1
                                            v22 := GenS1{F0: v2, F1: 169, F2: v2}
                                            if ((fi14 > fi17) || false) {
                                                var v23 uint64 = 0
                                                _ = len(v20)
                                                _ = v23
                                            }
                                            v5 = int16(v1)
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                        switch fi14 {
                                        case 477:
                                            var v24 uint = 2
                                            _ = v24
                                        case 2147483647:
                                            v25 := GenS2{F0: 4294967295, F1: (v2 + v2)}
                                            _ = v25
                                        case 0:
                                            if false {
                                                var v26 uint = 0
                                                _ = v26
                                            } else {
                                                switch fi17 {
                                                case 32768:
                                                    if (!true) {
                                                        var v27 uint8 = 1
                                                        if (-1 < (v5 * v5)) {
                                                            v28 := []int{fi17}
                                                            if ((^(v5)) > 32767) {
                                                                var v29 uint = 844
                                                                switch ((v2 + (("\n\t" + v2) + v2)) + (v2 + ("\n\t" + v2))) {
                                                                case "":
                                                                    v1 += (v1 >> ((v29) & 63))
                                                                case "hello":
                                                                    if ((^(v5)) > 42) {
                                                                        var v30 uint16 = (v1 * (v1 * (v1 | v1)))
                                                                        delete(v3, ((((("\n\t" + (v2 + ("foo bar" + (v2 + v2)))) + v2) + v2) + "") + v2))
                                                                        var v31 float64 = 0.1
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                default:
                                                                    fmt.Println(fi17)
                                                                    _ = len(v6)
                                                                    if false {
                                                                        var v32 int8 = v7
                                                                        switch -376 {
                                                                        case 32768:
                                                                            var v33 uint64 = 42
                                                                            _ = v33
                                                                        }
                                                                        var v34 float64 = 1.5
                                                                        _ = v32
                                                                        _ = v34
                                                                    } else {
                                                                        if false {
                                                                            if (v7 >= v7) {
                                                                                if ((("a" + "foo bar") + ((v2 + ("" + "\n\t")) + ("hello" + v2))) <= v2) {
                                                                                    if len(v28) >= 2 {
                                                                                        v28 = v28[1:len(v28)]
                                                                                    }
                                                                                    fmt.Println(fi17, fi17, v2)
                                                                                    v6[(fi14 - fi14)] = 2147483647
                                                                                } else {
                                                                                    v3[((("\n\t" + v2) + ("foo bar" + v2)) + "foo bar")] = "\n\t"
                                                                                }
                                                                                delete(v3, v2)
                                                                            } else {
                                                                                if (true || false) {
                                                                                    if false {
                                                                                        var v35 uint16 = v1
                                                                                        fmt.Println(fi17, v7)
                                                                                        if false {
                                                                                            switch 4294967295 {
                                                                                            case 0:
                                                                                                v2 = (("" + v2) + v2)
                                                                                            default:
                                                                                                fmt.Println(v7)
                                                                                            }
                                                                                        }
                                                                                        _ = v35
                                                                                    }
                                                                                    var v36 int64 = 9223372036854775807
                                                                                    fmt.Println(v7)
                                                                                    _ = v36
                                                                                }
                                                                                v2 = ""
                                                                            }
                                                                        }
                                                                        v2 = (v2 + "\n\t")
                                                                    }
                                                                }
                                                                v1 += v1
                                                                var v37 int64 = 9223372036854775807
                                                                _ = v29
                                                                _ = v37
                                                            }
                                                            _ = v28
                                                        } else {
                                                            switch 2147483647 {
                                                            case 4294967295:
                                                                var v38 uint32 = 2147483648
                                                                _ = v38
                                                            case 32768:
                                                                if ((^(v27)) <= v27) {
                                                                    v2 = "a"
                                                                } else {
                                                                    v39 := v6[fi14]
                                                                    v40 := GenS1{F0: "foo bar", F1: 2, F2: "hello"}
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                            }
                                                            var v41 int32 = -2147483648
                                                            var v42 int32 = 1
                                                            var v43 int16 = (v5 ^ -64)
                                                            v27 -= 1
                                                            v7 += v7
                                                            var v44 int32 = v41
                                                            _ = v41
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        if true {
                                                            var v45 uint16 = v1
                                                            _ = v45
                                                        }
                                                        if ((699 > 2147483647) || false) {
                                                            switch fi14 {
                                                            case 32767:
                                                                v46 := GenS1{F0: v2, F1: 2, F2: v2}
                                                                _ = v46
                                                            case -717:
                                                                v5 += (v5 & (v5 ^ -534))
                                                                v27 += (v27 ^ v27)
                                                            default:
                                                                var v47 bool = (255 >= 42)
                                                                _ = v47
                                                            }
                                                            if (true || true) {
                                                                var v48 int8 = v7
                                                                v5 -= (v5 >> ((2) & 63))
                                                                v1 += v1
                                                                _ = v48
                                                            }
                                                            v1 += 876
                                                        } else {
                                                            var v49 uint64 = 2
                                                            v7 = (v7 >> ((255) & 63))
                                                            switch "" {
                                                            case "":
                                                                v50 := []int{1, fi17, (fi14 ^ (fi17 ^ 205)), fi17}
                                                                _ = v50
                                                            case "\n\t":
                                                                fmt.Println(fi17, fi17)
                                                            default:
                                                                v51 := GenS3{F0: (^(v49))}
                                                                _ = v51
                                                            }
                                                            _ = v49
                                                        }
                                                        _ = v27
                                                    } else {
                                                        if (4294967295 < 1) {
                                                            switch 495 {
                                                            case -781:
                                                                var v52 int = (fi17 / (((^(fi17))) | 1))
                                                                _ = v52
                                                            }
                                                            var v53 uint32 = 2147483648
                                                            if true {
                                                                v54 := GenS2{F0: 42, F1: (((v2 + v2) + v2) + v2)}
                                                                var v55 string = (v2 + v2)
                                                                _ = v54
                                                                _ = v55
                                                            } else {
                                                                v1 = (v1 & 1)
                                                            }
                                                            if (3.14159 < 3.14159) {
                                                                if false {
                                                                    fmt.Println(v2)
                                                                    if true {
                                                                        var v56 bool = (!false)
                                                                        v53 = v53
                                                                        _ = v56
                                                                    }
                                                                    var v57 uint = 4294967295
                                                                    _ = v57
                                                                } else {
                                                                    var v58 float32 = 1e10
                                                                    v59 := GenS4{F0: true, F1: -2147483648}
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                                if true {
                                                                    v53 += v53
                                                                    v53 += (v53 * v53)
                                                                    var v60 bool = (((v5 >> ((4294967295) & 63)) == 1) || (!((false && false) && true)))
                                                                    _ = v60
                                                                }
                                                                var v61 uint64 = 4294967295
                                                                var v62 int32 = 2147483647
                                                                delete(v4, fi14)
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                v63 := GenS4{F0: false, F1: fi14}
                                                                _ = v63
                                                            }
                                                            _ = v53
                                                        } else {
                                                            if false {
                                                                v64 := []int{}
                                                                var v65 int32 = -402
                                                                v66 := GenS4{F0: (fi14 >= fi14), F1: fi17}
                                                                switch "a" {
                                                                case "foo bar":
                                                                    switch v2 {
                                                                    case "a":
                                                                        var v67 float64 = 1.5
                                                                        _ = v67
                                                                    case "hello":
                                                                        var v68 int8 = v7
                                                                        _ = v68
                                                                    case "foo bar":
                                                                        _ = len(v64)
                                                                    }
                                                                default:
                                                                    v5 += (v5 | -943)
                                                                }
                                                                if len(v64) >= 2 {
                                                                    v64 = v64[1:len(v64)]
                                                                }
                                                                var v69 int64 = -381
                                                                if true {
                                                                    switch (-(fi14)) {
                                                                    case 4294967295:
                                                                        v3["\n\t"] = v2
                                                                        var v70 bool = (1 != 4294967295)
                                                                        _ = v70
                                                                    case 307:
                                                                        var v71 uint64 = 0
                                                                        _ = v71
                                                                    case 688:
                                                                        switch fi17 {
                                                                        case 9223372036854775807:
                                                                            fmt.Println(v69, v7)
                                                                        default:
                                                                            switch (fi14 - fi17) {
                                                                            case 42:
                                                                                switch fi14 {
                                                                                case 9223372036854775807:
                                                                                    switch 32768 {
                                                                                    case 2147483647:
                                                                                        if false {
                                                                                            fmt.Println(fi14, fi14)
                                                                                        }
                                                                                    case -2147483648:
                                                                                        var v72 bool = (false || false)
                                                                                        _ = v72
                                                                                    case 1:
                                                                                        var v73 int8 = 100
                                                                                        _ = v73
                                                                                    default:
                                                                                        v5 += (v5 ^ 954)
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    switch (("a" + v2) + "a") {
                                                                                    case "foo bar":
                                                                                        v74 := GenS4{F0: false, F1: (fi14 + fi14)}
                                                                                        _ = v74
                                                                                    case "hello":
                                                                                        fmt.Println(v1, v2)
                                                                                        switch fi17 {
                                                                                        case -2147483648:
                                                                                            v75 := GenS5{F0: v2, F1: (v69 * v69), F2: fi14}
                                                                                            _ = v75
                                                                                        case 604:
                                                                                            _ = len(v64)
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v76 int8 = (-(v7))
                                                                                        _ = v76
                                                                                    default:
                                                                                        v77 := v4[2147483647]
                                                                                        _ = v77
                                                                                    }
                                                                                }
                                                                            case 9223372036854775807:
                                                                                v5 -= (-(v5))
                                                                                if false {
                                                                                    var v78 uint32 = 2
                                                                                    if (1 != (v1 + (^(v1)))) {
                                                                                        var v79 bool = false
                                                                                        var v80 int32 = (v65 * (v65 / (((v65 >> ((255) & 63))) | 1)))
                                                                                        v81 := v6[(fi17 + fi17)]
                                                                                        switch fi14 {
                                                                                        case 4294967295:
                                                                                            var v82 uint32 = v78
                                                                                            _ = v82
                                                                                        case 1:
                                                                                            fmt.Println(v5, v81)
                                                                                        case 471:
                                                                                            fmt.Println(v79, v79, v80)
                                                                                        default:
                                                                                            v5 -= v5
                                                                                        }
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    }
                                                                                    _ = v78
                                                                                } else {
                                                                                    v7 += v7
                                                                                    if (-2.25 == 3.14159) {
                                                                                        if len(v64) > 0 {
                                                                                            v64[0] = fi14
                                                                                        }
                                                                                    }
                                                                                    var v83 float64 = ((0.1 + float64(1.5)) * 0.0)
                                                                                    _ = v83
                                                                                }
                                                                            default:
                                                                                if (true && false) {
                                                                                    v7 -= (v7 - (v7 >> ((42) & 63)))
                                                                                } else {
                                                                                    v69 = v69
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v84 := GenS4{F0: true, F1: int(v65)}
                                                                    _ = v84
                                                                }
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                                _ = v69
                                                            }
                                                            v85 := GenS4{F0: true, F1: 0}
                                                            if true {
                                                                if false {
                                                                    if (-477 == v5) {
                                                                        if false {
                                                                            fmt.Println(v2)
                                                                            if (-1 >= -764) {
                                                                                v2 = v2
                                                                            }
                                                                        } else {
                                                                            var v86 float32 = (((((1.5 - float32((3.14159 * -2.25))) - 0.0) * 1e10) + (3.14159 - 3.14159)) - 1.5)
                                                                            switch 663 {
                                                                            case 713:
                                                                                v87 := GenS2{F0: 256, F1: v2}
                                                                                _ = v87
                                                                            case 1:
                                                                                v88 := GenS1{F0: (("foo bar" + v2) + "\n\t"), F1: 317, F2: v2}
                                                                                _ = v88
                                                                            case -1:
                                                                                switch v2 {
                                                                                case "a":
                                                                                    v89 := GenS4{F0: true, F1: (fi14 + (fi14 % (((fi17 | fi17)) | 1)))}
                                                                                    _ = v89
                                                                                case "":
                                                                                    v2 = v2
                                                                                case "foo bar":
                                                                                    var v90 int8 = (-(v7))
                                                                                    _ = v90
                                                                                default:
                                                                                    var v91 float32 = (float32(-2.25) - v86)
                                                                                    _ = v91
                                                                                }
                                                                            default:
                                                                                if (9223372036854775807 > -1) {
                                                                                    v4[852] = false
                                                                                    v2 = v2
                                                                                    fmt.Println(fi17)
                                                                                }
                                                                            }
                                                                            var v92 int32 = 2147483647
                                                                            _ = v86
                                                                            _ = v92
                                                                        }
                                                                    } else {
                                                                        if (228 < 4294967295) {
                                                                            if false {
                                                                                v93 := GenS6{F0: true}
                                                                                v94 := GenS5{F0: "", F1: -9223372036854775808, F2: fi17}
                                                                                v7 -= v7
                                                                                v95 := []int{(-(fi17)), (fi14 ^ fi14), -1}
                                                                                switch ((v2 + ((v2 + ((v2 + (v2 + v2)) + ("foo bar" + v2))) + v2)) + (v2 + v2)) {
                                                                                case "\n\t":
                                                                                    var v96 float64 = 0.1
                                                                                    _ = v96
                                                                                case "hello":
                                                                                    v1 = (v1 * 2)
                                                                                case "foo bar":
                                                                                    var v97 float32 = (0.1 / ((0.0 + 1.5) - 1e10))
                                                                                    _ = v97
                                                                                }
                                                                                if (false && true) {
                                                                                    var v98 uint = 255
                                                                                    if true {
                                                                                        v99 := GenS2{F0: 343, F1: ""}
                                                                                        _ = v99
                                                                                    } else {
                                                                                        v100 := GenS7{F0: (v7 - int8(v98))}
                                                                                        if ((v1 / ((1) | 1)) <= (v1 & (v1 / ((v1) | 1)))) {
                                                                                            switch ("" + v2) {
                                                                                            case "foo bar":
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    fmt.Println(v7, v5)
                                                                                                default:
                                                                                                    var v101 uint64 = 42
                                                                                                    _ = v101
                                                                                                }
                                                                                            case "a":
                                                                                                delete(v4, fi14)
                                                                                            default:
                                                                                                v102 := GenS1{F0: v2, F1: 1, F2: v2}
                                                                                                _ = v102
                                                                                            }
                                                                                            v7 -= (v7 << (((v98 - (v98 & (v98 << ((v98) & 63))))) & 63))
                                                                                            v103 := GenS8{F0: v98, F1: 2147483647, F2: false}
                                                                                            _ = v103
                                                                                        }
                                                                                        _ = v100
                                                                                    }
                                                                                    _ = v98
                                                                                } else {
                                                                                    v104 := GenS1{F0: "\n\t", F1: 1, F2: "hello"}
                                                                                    _ = v104
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v2, v1)
                                                                                    switch v2 {
                                                                                    case "\n\t":
                                                                                        v5 -= (v5 + (v5 >> ((65535) & 63)))
                                                                                    case "hello":
                                                                                        var v105 uint16 = v1
                                                                                        _ = v105
                                                                                    case "a":
                                                                                        v5 = (v5 / ((v5) | 1))
                                                                                    }
                                                                                    v1 += 0
                                                                                    if len(v95) >= 2 {
                                                                                        v95 = v95[1:len(v95)]
                                                                                    }
                                                                                    var v106 int8 = (v7 + 42)
                                                                                    _ = v106
                                                                                }
                                                                                if (165 == -1) {
                                                                                    var v107 int8 = (v7 << ((2) & 63))
                                                                                    v108 := GenS2{F0: 18446744073709551615, F1: (("" + v2) + "")}
                                                                                    _ = v107
                                                                                    _ = v108
                                                                                } else {
                                                                                    switch fi14 {
                                                                                    case 32768:
                                                                                        var v109 uint64 = 2
                                                                                        _ = v109
                                                                                    case -1:
                                                                                        if (fi17 > (fi14 | fi14)) {
                                                                                            v5 = v5
                                                                                            if true {
                                                                                                v7 -= 48
                                                                                                v7 = v7
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if ((1 > 562) || false) {
                                                                                        if true {
                                                                                            v5 += 0
                                                                                        } else {
                                                                                            v110 := GenS9{F0: (v7 - -128)}
                                                                                            v1 -= v1
                                                                                            _ = v110
                                                                                        }
                                                                                        v2 = (v2 + v2)
                                                                                        fmt.Println(fi14, v7)
                                                                                    }
                                                                                }
                                                                                fmt.Println(v1)
                                                                                _ = v93
                                                                                _ = v94
                                                                                _ = v95
                                                                            } else {
                                                                                var v111 int16 = (v5 & v5)
                                                                                var v112 uint = 65535
                                                                                v113 := GenS2{F0: 2, F1: "a"}
                                                                                _ = v111
                                                                                _ = v112
                                                                                _ = v113
                                                                            }
                                                                            v114 := []int{fi14, fi17}
                                                                            _ = v114
                                                                        } else {
                                                                            var v115 int32 = -434
                                                                            v116 := GenS10{F0: 1, F1: (183 != 0), F2: v115}
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                    v2 = "hello"
                                                                }
                                                            }
                                                            if (!true) {
                                                                fmt.Println(v7, v5, fi17)
                                                            } else {
                                                                var v117 int = (^(fi17))
                                                                v118 := GenS3{F0: 4294967295}
                                                                _ = v117
                                                                _ = v118
                                                            }
                                                            var v119 bool = false
                                                            fmt.Println(v7)
                                                            _ = v119
                                                            _ = v85
                                                        }
                                                        v120 := v6[(fi14 / ((fi14) | 1))]
                                                        v120 = v120
                                                        _ = v120
                                                    }
                                                case 32767:
                                                    switch v2 {
                                                    case "hello":
                                                        var v121 uint64 = 0
                                                        _ = v121
                                                    case "\n\t":
                                                        fmt.Println(v5, fi17, v2)
                                                    }
                                                }
                                                if true {
                                                    var v122 int32 = -2147483648
                                                    v122 = (v122 - (v122 ^ (v122 >> ((1) & 63))))
                                                    _ = v122
                                                }
                                            }
                                        }
                                    } else {
                                        delete(v6, 188)
                                    }
                                    if true {
                                        v7 -= (v7 + (v7 >> ((0) & 63)))
                                        v123 := GenS11{F0: -1, F1: 2, F2: 919}
                                        if (938 > 926) {
                                            if true {
                                                switch v2 {
                                                case "hello":
                                                    var v124 int16 = -730
                                                    _ = v124
                                                case "\n\t":
                                                    v125 := GenS12{F0: (v1 % ((728) | 1)), F1: -900}
                                                    _ = v125
                                                case "foo bar":
                                                    var v126 bool = (!false)
                                                    _ = v126
                                                }
                                                var v127 uint32 = 2
                                                var v128 int = fi17
                                                _ = v127
                                                _ = v128
                                            } else {
                                                if ((-(v5)) > (v5 >> ((255) & 63))) {
                                                    delete(v6, fi17)
                                                    switch "a" {
                                                    case "a":
                                                        v2 = v2
                                                    case "\n\t":
                                                        v2 = (v2 + "")
                                                    case "":
                                                        v1 -= (^(v1))
                                                    default:
                                                        v2 = (("hello" + v2) + v2)
                                                    }
                                                    var v129 uint64 = 220
                                                    _ = v129
                                                }
                                                var v130 int32 = 42
                                                _ = v130
                                            }
                                            if (!true) {
                                                if (!(!(128 == 207))) {
                                                    v7 = (v7 + 42)
                                                    fmt.Println(v5, fi17)
                                                } else {
                                                    var v131 int8 = v7
                                                    var v132 uint8 = 128
                                                    if false {
                                                        v1 -= 42
                                                        if false {
                                                            v7 -= v131
                                                            v133 := GenS13{F0: 486, F1: 0, F2: 2}
                                                            _ = v133
                                                        }
                                                    } else {
                                                        if false {
                                                            v134 := []int{fi17, fi17, (fi17 | fi17)}
                                                            v1 = v1
                                                            v131 = v7
                                                            var v135 uint64 = 1
                                                            _ = v134
                                                            _ = v135
                                                        } else {
                                                            if (true || false) {
                                                                v131 = 100
                                                                v5 = (v5 << ((42) & 63))
                                                            }
                                                            if false {
                                                                if false {
                                                                    if (18446744073709551615 <= 294) {
                                                                        var v136 float32 = 0.1
                                                                        _ = v136
                                                                    }
                                                                    v137 := GenS2{F0: 1, F1: "\n\t"}
                                                                    v1 = 32767
                                                                    if (42 <= 697) {
                                                                        if (0 >= v132) {
                                                                            if ((v131 + (v7 % ((-1) | 1))) == 0) {
                                                                                fmt.Println(v1, fi14)
                                                                                var v138 int32 = -1
                                                                                v5 += (v5 + v5)
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    fmt.Println(v131, v2, v1)
                                                                                case "":
                                                                                    if (v132 > 2) {
                                                                                        delete(v6, fi17)
                                                                                        v139 := []int{(fi17 << ((972) & 63)), fi17}
                                                                                        var v140 uint16 = v1
                                                                                        _ = v139
                                                                                        _ = v140
                                                                                    } else {
                                                                                        if (true && false) {
                                                                                            var v141 int32 = v138
                                                                                            v131 += 127
                                                                                            _ = v141
                                                                                        }
                                                                                        fmt.Println(fi14, fi14, v7)
                                                                                    }
                                                                                }
                                                                                _ = v138
                                                                            }
                                                                        }
                                                                        v142 := []int{}
                                                                        _ = v142
                                                                    } else {
                                                                        delete(v4, (^(fi14)))
                                                                        v132 += (^(v132))
                                                                    }
                                                                    _ = v137
                                                                }
                                                                v143 := GenS10{F0: 338, F1: (0 >= 4294967295), F2: 42}
                                                                _ = len(v3)
                                                                _ = v143
                                                            }
                                                            var v144 uint64 = 18446744073709551615
                                                            _ = v144
                                                        }
                                                        var v145 float64 = (-2.25 * ((1.5 + float64((-2.25 - 0.1))) - (0.0 / (1.5 * 1e10))))
                                                        _ = v145
                                                    }
                                                    _ = v131
                                                    _ = v132
                                                }
                                                switch fi17 {
                                                case -2147483648:
                                                    fmt.Println(fi14, fi14, v7)
                                                }
                                            } else {
                                                v146 := []int{-9223372036854775808}
                                                _ = v146
                                            }
                                            var v147 string = ""
                                            _ = v147
                                        }
                                        _ = v123
                                    }
                                    fmt.Println(v1, v1, fi14)
                                }
                                v7 = (v7 << ((255) & 63))
                                var v148 int16 = (v5 << ((65535) & 63))
                                _ = fi17
                                _ = v148
                            }
                        }
                        v1 += (v1 * v1)
                    }
                }
                var v149 int8 = v7
                if (!true) {
                    fmt.Println(fi14, fi14)
                    var v150 uint32 = 899
                    _ = v150
                }
                _ = v149
            }
            for fi151 := 0; fi151 < 6; fi151++ {
                var v152 int16 = (v5 | (v5 ^ -991))
                v2 = (v2 + ("a" + "hello"))
                _ = fi151
                _ = v152
            }
            _ = fi14
        }
    case 720:
        for fi153 := 0; fi153 < 16; fi153++ {
            for fi154 := 0; fi154 < 6; fi154++ {
                if ((!(false || ("\n\t" > (v2 + ("hello" + (((v2 + (v2 + ((v2 + v2) + "foo bar"))) + ((v2 + v2) + "\n\t")) + "foo bar")))))) || (!true)) {
                    var v155 int32 = 2147483647
                    var v156 string = v2
                    v157 := GenS14{F0: 42, F1: (v7 << ((4294967295) & 63)), F2: v7}
                    _ = v155
                    _ = v156
                    _ = v157
                }
                v2 = ("foo bar" + "hello")
                _ = fi154
            }
            var v158 uint64 = 42
            _ = fi153
            _ = v158
        }
    }
    v159 := []int{128, 42}
    if true {
        _ = len(v159)
        if false {
            v3[v2] = ((("foo bar" + v2) + "\n\t") + (v2 + (v2 + v2)))
            v160 := GenS15{F0: 2147483647, F1: v1, F2: (v2 >= "\n\t")}
            fmt.Println(v1, v2)
            _ = v160
        }
    }
    if false {
        v161 := GenS16{F0: 18446744073709551615, F1: "\n\t", F2: 0}
        v7 += (v7 << ((1) & 63))
        var v162 uint8 = 255
        _ = v161
        _ = v162
    }
    var v163 string = ((((v2 + v2) + (v2 + v2)) + v2) + (v2 + ""))
    v7 = 127
    if false {
        v163 = "\n\t"
        v164 := GenS17{F0: 17, F1: 285, F2: 192}
        v2 = v2
        _ = v164
    } else {
        var v165 uint32 = 2147483647
        if true {
            for fi166 := 0; fi166 < 4; fi166++ {
                v1 = v1
                _ = fi166
            }
            v167 := GenS16{F0: 922, F1: (("\n\t" + ("foo bar" + "hello")) + v2), F2: 255}
            _ = v167
        }
        _ = v165
    }
    v168 := GenS9{F0: v7}
    v2 = ((("" + ((("a" + v2) + v163) + v2)) + "hello") + (("\n\t" + v2) + (v2 + (v2 + (v2 + "")))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v159:", len(v159), v159)
    fmt.Println("v163:", v163)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v159
    _ = v163
    _ = v168
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
