package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 string }
type GenS3 struct { F0 uint; F1 string; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool; F2 string }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 string; F2 uint32 }
type GenS11 struct { F0 int }
func genFn0() uint64 {
    v6 := []int{0, -2147483648}
    var v7 int = 32767
    _ = len(v6)
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 0
    v3 := []string{"a", ("foo bar" + "foo bar"), "foo bar"}
    var v4 bool = true
    var v5 string = ("\n\t" + "")
    v1 += v1
    var v8 float32 = 1e10
    var v9 int64 = v2
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    v5 = v5
    if v4 {
        v1 += (v1 * (-(v1)))
        switch 4294967295 {
        case 1:
            for fi10 := 0; fi10 < 10; fi10++ {
                switch fi10 {
                case 32767:
                    switch (^(fi10)) {
                    case -9223372036854775808:
                        var v11 uint32 = 902
                        _ = v11
                    }
                case -295:
                    var v12 uint8 = 167
                    _ = v12
                case 338:
                    var v13 uint16 = 829
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v1 += v1
                        if true {
                            v1 += v1
                            _ = len(v3)
                        }
                        _ = fi14
                    }
                }
                v1 = (^(v1))
                var v15 uint16 = 927
                _ = fi10
                _ = v15
            }
        case 854:
            v3 = append(v3, ("foo bar" + (v5 + v5)))
        case 32767:
            v16 := GenS1{F0: 1, F1: -700}
            _ = v16
        }
    }
    if true {
        _ = len(v3)
        for fi17 := 0; fi17 < 10; fi17++ {
            switch (^(fi17)) {
            case -9223372036854775808:
                v3 = append(v3, v5)
            }
            _ = fi17
        }
        var v18 int8 = 2
        _ = v18
    } else {
        var v19 string = v5
        if v4 {
            if (false || v4) {
                switch v19 {
                case "\n\t":
                    _ = len(v3)
                default:
                    switch (v19 + ("" + v5)) {
                    case "\n\t":
                        v3 = append(v3, "\n\t")
                    case "":
                        if v4 {
                            for fi20 := 0; fi20 < 4; fi20++ {
                                fmt.Println(v4)
                                v4 = v4
                                _ = len(v3)
                                _ = fi20
                            }
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v5 = v19
                                _ = fi21
                            }
                            fmt.Println(v19, v8)
                        } else {
                            if (18446744073709551615 != genFn0()) {
                                v8 = (((v8 * float32(((float64(v8) / ((0.1 / 1.5) * 0.1)) / 1.0))) + 0.0) * v8)
                                if (-2.25 > (v8 * (float32(((1e10 / 1.0) - (-2.25 / -2.25))) + 0.1))) {
                                    switch -2147483648 {
                                    case -894:
                                        var v22 int = 4294967295
                                        _ = v22
                                    case 127:
                                        for fi23 := 0; fi23 < 6; fi23++ {
                                            v3 = append(v3, (v19 + (v19 + v5)))
                                            switch v5 {
                                            case "a":
                                                v19 = "a"
                                            case "hello":
                                                switch v5 {
                                                case "":
                                                    var v24 uint64 = 0
                                                    _ = v24
                                                }
                                            case "":
                                                v25 := GenS2{F0: v19, F1: -52, F2: ("hello" + v19)}
                                                _ = v25
                                            default:
                                                fmt.Println(fi23, v9)
                                            }
                                            _ = fi23
                                        }
                                    default:
                                        if v4 {
                                            for fi26 := 0; fi26 < 17; fi26++ {
                                                for fi27 := 0; fi27 < 5; fi27++ {
                                                    var v28 string = v19
                                                    v29 := GenS1{F0: 65535, F1: -345}
                                                    _ = len(v3)
                                                    var v30 int32 = 2147483647
                                                    _ = fi27
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = "foo bar"
                                                }
                                                for fi31 := 0; fi31 < 9; fi31++ {
                                                    v8 = v8
                                                    if v4 {
                                                        v1 -= int32(fi26)
                                                        v8 = (v8 - v8)
                                                        var v32 uint64 = 1
                                                        _ = v32
                                                    } else {
                                                        v19 = "hello"
                                                        var v33 uint32 = 4294967295
                                                        _ = v33
                                                    }
                                                    switch fi31 {
                                                    case 9223372036854775807:
                                                        if (929 <= 4294967295) {
                                                            v34 := GenS2{F0: v19, F1: 0, F2: ((("foo bar" + (v19 + ("a" + v5))) + v5) + "foo bar")}
                                                            fmt.Println(v9, v9)
                                                            v4 = v4
                                                            _ = v34
                                                        } else {
                                                            if (!v4) {
                                                                switch -2147483648 {
                                                                case -2147483648:
                                                                    if false {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v5
                                                                        }
                                                                        v4 = v4
                                                                        v35 := GenS2{F0: v19, F1: 1, F2: v5}
                                                                        _ = v35
                                                                    }
                                                                    _ = len(v3)
                                                                case 32768:
                                                                    v1 += (v1 << ((0) & 63))
                                                                case 0:
                                                                    v2 -= (v2 << ((654) & 63))
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        fmt.Println(v4, fi31, fi31)
                                                                    case "a":
                                                                        var v36 uint16 = 255
                                                                        _ = v36
                                                                    case "hello":
                                                                        var v37 int8 = 127
                                                                        _ = v37
                                                                    default:
                                                                        switch (v19 + "a") {
                                                                        case "hello":
                                                                            v8 -= 3.14159
                                                                        case "foo bar":
                                                                            if (!v4) {
                                                                                if (!(!false)) {
                                                                                    v1 += 0
                                                                                    switch (-(fi31)) {
                                                                                    case 4294967295:
                                                                                        if v4 {
                                                                                            fmt.Println(v1)
                                                                                        }
                                                                                    case 42:
                                                                                        if v4 {
                                                                                            v19 = (v19 + v19)
                                                                                            var v38 float32 = (v8 - v8)
                                                                                            var v39 float64 = 1e10
                                                                                            _ = v38
                                                                                            _ = v39
                                                                                        } else {
                                                                                            switch (fi31 % (((fi31 & 42)) | 1)) {
                                                                                            case -9223372036854775808:
                                                                                                var v40 int16 = -341
                                                                                                _ = v40
                                                                                            default:
                                                                                                var v41 float64 = (((((0.1 - 0.0) + 0.1) + 3.14159) * (0.1 * (1e10 - 0.1))) * ((1e10 * (0.1 - 0.1)) / 0.1))
                                                                                                _ = v41
                                                                                            }
                                                                                            var v42 float32 = v8
                                                                                            var v43 float64 = 3.14159
                                                                                            _ = v42
                                                                                            _ = v43
                                                                                        }
                                                                                    }
                                                                                    switch fi26 {
                                                                                    case 32767:
                                                                                        v2 = 1
                                                                                    default:
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = v5
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v44 int32 = (^(v1))
                                                                                if true {
                                                                                    fmt.Println(v2)
                                                                                    switch ("" + v19) {
                                                                                    case "foo bar":
                                                                                        if (18446744073709551615 > 2) {
                                                                                            if (!true) {
                                                                                                if (v4 && (!(18446744073709551615 < 18446744073709551615))) {
                                                                                                    var v45 float64 = 0.0
                                                                                                    if true {
                                                                                                        fmt.Println(v44, fi26)
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        v8 = ((v8 + v8) + float32(v45))
                                                                                                        var v46 bool = v4
                                                                                                        if (!v46) {
                                                                                                            if v46 {
                                                                                                                fmt.Println(v1, v45)
                                                                                                            }
                                                                                                            v47 := GenS1{F0: 312, F1: -670}
                                                                                                            _ = v47
                                                                                                        }
                                                                                                        _ = v46
                                                                                                    }
                                                                                                    var v48 int64 = v9
                                                                                                    _ = v45
                                                                                                    _ = v48
                                                                                                }
                                                                                                var v49 uint = 1
                                                                                                var v50 int = -1
                                                                                                _ = v49
                                                                                                _ = v50
                                                                                            } else {
                                                                                                var v51 int16 = 0
                                                                                                if v4 {
                                                                                                    var v52 uint64 = 716
                                                                                                    _ = v52
                                                                                                }
                                                                                                v1 -= (-(v44))
                                                                                                _ = v51
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if (!(!(v4 && v4))) {
                                                                                            var v53 float32 = ((v8 / v8) / 0.1)
                                                                                            _ = v53
                                                                                        }
                                                                                    case "hello":
                                                                                        v44 += -2147483648
                                                                                    default:
                                                                                        v54 := GenS1{F0: 65535, F1: 1000}
                                                                                        v5 = ("foo bar" + (v5 + v19))
                                                                                        if v4 {
                                                                                            var v55 int16 = 1
                                                                                            _ = v55
                                                                                        } else {
                                                                                            if (-22 > (-(v9))) {
                                                                                                switch v5 {
                                                                                                case "":
                                                                                                    var v56 int16 = 0
                                                                                                    _ = v56
                                                                                                case "a":
                                                                                                    if v4 {
                                                                                                        v2 = (v9 / ((v9) | 1))
                                                                                                        var v57 bool = (v4 || ((((-9223372036854775808 < 211) && (!false)) || (65535 <= 669)) || (!v4)))
                                                                                                        _ = v57
                                                                                                    } else {
                                                                                                        if v4 {
                                                                                                            fmt.Println(fi26)
                                                                                                            var v58 bool = ((!(397 < 0)) && ((v2 ^ (v9 << ((0) & 63))) >= 42))
                                                                                                            _ = v58
                                                                                                        } else {
                                                                                                            switch (fi31 | (-(fi31))) {
                                                                                                            case 2147483647:
                                                                                                                v59 := GenS3{F0: 42, F1: v19, F2: 125}
                                                                                                                _ = v59
                                                                                                            case 32768:
                                                                                                                v2 -= int64(v1)
                                                                                                            default:
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = (((((v19 + ("hello" + (v19 + ((v5 + ((v5 + v19) + v19)) + "foo bar")))) + v5) + "") + v5) + (v19 + "hello"))
                                                                                                                }
                                                                                                            }
                                                                                                            switch "a" {
                                                                                                            case "":
                                                                                                                if len(v3) >= 2 {
                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                v60 := GenS3{F0: 65535, F1: v5, F2: 127}
                                                                                                                _ = v60
                                                                                                            default:
                                                                                                                var v61 int64 = v9
                                                                                                                _ = v61
                                                                                                            }
                                                                                                            v44 -= int32(v2)
                                                                                                        }
                                                                                                        v8 += float32(1.5)
                                                                                                        switch (fi26 ^ 987) {
                                                                                                        case -205:
                                                                                                            switch (^(fi31)) {
                                                                                                            case 127:
                                                                                                                var v62 bool = (1 > 250)
                                                                                                                _ = v62
                                                                                                            case 2147483647:
                                                                                                                var v63 uint8 = 2
                                                                                                                _ = v63
                                                                                                            }
                                                                                                        case 2147483647:
                                                                                                            var v64 uint16 = 42
                                                                                                            _ = v64
                                                                                                        case 1:
                                                                                                            if (1 < 4294967295) {
                                                                                                                _ = len(v3)
                                                                                                            } else {
                                                                                                                switch v5 {
                                                                                                                case "\n\t":
                                                                                                                    v19 = ("hello" + v5)
                                                                                                                }
                                                                                                                var v65 uint32 = 0
                                                                                                                _ = v65
                                                                                                            }
                                                                                                        default:
                                                                                                            v9 -= (-(v9))
                                                                                                        }
                                                                                                        var v66 int8 = 25
                                                                                                        switch fi26 {
                                                                                                        case 0:
                                                                                                            var v67 int = (fi31 * (fi31 - fi31))
                                                                                                            _ = v67
                                                                                                        default:
                                                                                                            v1 = v44
                                                                                                        }
                                                                                                        _ = v66
                                                                                                    }
                                                                                                default:
                                                                                                    var v68 bool = v4
                                                                                                    _ = v68
                                                                                                }
                                                                                                var v69 float32 = (float32(0.0) / 0.1)
                                                                                                _ = v69
                                                                                            } else {
                                                                                                if (float64(1.5) < 1e10) {
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    if false {
                                                                                                        var v70 int = (fi31 / (((^(fi31))) | 1))
                                                                                                        var v71 uint32 = 9
                                                                                                        v1 += (v44 / ((v1) | 1))
                                                                                                        _ = v70
                                                                                                        _ = v71
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v54
                                                                                    }
                                                                                    v9 = -930
                                                                                } else {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    v1 += v1
                                                                                    v2 = 0
                                                                                }
                                                                                if v4 {
                                                                                    var v72 int8 = 127
                                                                                    v73 := GenS3{F0: 1, F1: v5, F2: (v72 >> ((1) & 63))}
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                } else {
                                                                                    if false {
                                                                                        if (v1 < (v1 & (v1 & 2147483647))) {
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (((("foo bar" + "foo bar") + ((v5 + v5) + ("foo bar" + v19))) + ((("a" + v19) + v5) + (v19 + v5))) + v5)
                                                                                            }
                                                                                            v9 -= (v2 << ((856) & 63))
                                                                                        } else {
                                                                                            v74 := GenS4{F0: v4}
                                                                                            if v4 {
                                                                                                v9 = v9
                                                                                            } else {
                                                                                                v9 -= -1
                                                                                                var v75 uint8 = 2
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = ((v5 + v19) + v19)
                                                                                                }
                                                                                                _ = v75
                                                                                            }
                                                                                            if v4 {
                                                                                                v76 := GenS5{F0: 2, F1: v4, F2: "foo bar"}
                                                                                                v2 = (v2 & v2)
                                                                                                _ = v76
                                                                                            } else {
                                                                                                var v77 float32 = (v8 / ((1e10 / ((v8 - 0.0) - v8)) - v8))
                                                                                                _ = v77
                                                                                            }
                                                                                            _ = v74
                                                                                        }
                                                                                        v4 = (((v5 >= ("\n\t" + "\n\t")) && true) || ((v8 * (v8 + 0.0)) > (3.14159 + float32(0.0))))
                                                                                    } else {
                                                                                        if ((v44 < (v1 ^ (v1 | 0))) && (float32(0.1) == v8)) {
                                                                                            if v4 {
                                                                                                if true {
                                                                                                    var v78 uint64 = 2
                                                                                                    _ = len(v3)
                                                                                                    _ = v78
                                                                                                } else {
                                                                                                    v5 = v19
                                                                                                    var v79 bool = (2147483648 == 311)
                                                                                                    var v80 uint32 = 872
                                                                                                    _ = v79
                                                                                                    _ = v80
                                                                                                }
                                                                                                v1 += 2147483647
                                                                                            }
                                                                                            v1 = v1
                                                                                            v44 = (v44 / (((v1 | v44)) | 1))
                                                                                        } else {
                                                                                            v44 -= (v44 & 0)
                                                                                        }
                                                                                        var v81 uint32 = 65535
                                                                                        v5 = "a"
                                                                                        _ = v81
                                                                                    }
                                                                                }
                                                                                var v82 uint8 = 167
                                                                                if ((v82 << ((2) & 63)) <= v82) {
                                                                                    v8 = 1e10
                                                                                } else {
                                                                                    v83 := GenS2{F0: v5, F1: -28, F2: "hello"}
                                                                                    _ = v83
                                                                                }
                                                                                _ = v44
                                                                                _ = v82
                                                                            }
                                                                        case "":
                                                                            if false {
                                                                                v2 += (v2 * 0)
                                                                                v9 = -1
                                                                            } else {
                                                                                fmt.Println(fi26)
                                                                            }
                                                                        default:
                                                                            v9 = 0
                                                                            if v4 {
                                                                                var v84 float32 = float32(((float64(3.14159) + (1e10 - 0.1)) + 3.14159))
                                                                                _ = v84
                                                                            } else {
                                                                                v3 = append(v3, (("" + (((v5 + ("a" + "a")) + (("" + "a") + ((((((v19 + (v5 + "foo bar")) + v5) + v5) + ("\n\t" + ("foo bar" + (v19 + v5)))) + "foo bar") + v5))) + (("hello" + "hello") + ""))) + ("foo bar" + v5)))
                                                                                var v85 int64 = 42
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        switch fi31 {
                                                                        case -9223372036854775808:
                                                                            var v86 int8 = int8(fi26)
                                                                            _ = v86
                                                                        }
                                                                    case "hello":
                                                                        if (2 > 2) {
                                                                            v87 := GenS6{F0: (18446744073709551615 < 1), F1: 42}
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            _ = v87
                                                                        }
                                                                    case "\n\t":
                                                                        v5 = "foo bar"
                                                                    }
                                                                    v88 := GenS4{F0: (4294967295 <= 4294967295)}
                                                                    v1 -= (v1 - 42)
                                                                    _ = v88
                                                                }
                                                            }
                                                            v19 = v5
                                                            if v4 {
                                                                if v4 {
                                                                    var v89 int32 = (v1 % ((2147483647) | 1))
                                                                    _ = v89
                                                                } else {
                                                                    var v90 bool = v4
                                                                    _ = v90
                                                                }
                                                            }
                                                        }
                                                    case 270:
                                                        v91 := GenS1{F0: 273, F1: 1000}
                                                        _ = v91
                                                    }
                                                    _ = fi31
                                                }
                                                _ = fi26
                                            }
                                        }
                                    }
                                    v3 = append(v3, "")
                                } else {
                                    var v92 int64 = 979
                                    v1 = v1
                                    var v93 string = v19
                                    _ = v92
                                    _ = v93
                                }
                            }
                            for fi94 := 0; fi94 < 1; fi94++ {
                                v3 = append(v3, v19)
                                _ = fi94
                            }
                            for fi95 := 0; fi95 < 19; fi95++ {
                                switch (fi95 | (fi95 >> ((856) & 63))) {
                                case 2147483647:
                                    var v96 string = v5
                                    _ = v96
                                }
                                v5 = v19
                                var v97 uint16 = 42
                                _ = fi95
                                _ = v97
                            }
                        }
                    }
                }
                v98 := GenS7{F0: 42}
                _ = v98
            } else {
                var v99 uint8 = 2
                v9 -= 0
                if true {
                    v1 += 42
                    var v100 int32 = v1
                    _ = v100
                }
                _ = v99
            }
        } else {
            v101 := GenS8{F0: "a"}
            _ = v101
        }
        var v102 int64 = 548
        _ = v102
        _ = v19
    }
    fmt.Println(v1, v1, v4)
    v103 := GenS9{F0: 32768}
    v8 += 1e10
    v104 := GenS6{F0: ((4294967295 > 496) && v4), F1: -990}
    for fi105 := 0; fi105 < 3; fi105++ {
        v4 = (!v4)
        switch v5 {
        case "a":
            _ = len(v3)
        case "":
            var v106 int32 = v1
            _ = v106
        case "\n\t":
            for fi107 := 0; fi107 < 15; fi107++ {
                switch (fi107 | (fi105 | fi105)) {
                case -471:
                    v108 := GenS6{F0: true, F1: 0}
                    _ = v108
                case 32767:
                    if false {
                        if v4 {
                            v109 := GenS1{F0: 321, F1: 32767}
                            var v110 int8 = 0
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v109
                            _ = v110
                        }
                        v111 := GenS5{F0: 42, F1: (v4 || ((944 < 926) && v4)), F2: (v5 + v5)}
                        v112 := GenS10{F0: 4294967295, F1: v5, F2: 42}
                        _ = v111
                        _ = v112
                    }
                case -491:
                    var v113 int8 = 100
                    v4 = (v4 || v4)
                    _ = v113
                default:
                    v114 := GenS11{F0: (^(fi107))}
                    _ = v114
                }
                var v115 uint8 = 244
                _ = fi107
                _ = v115
            }
        default:
            if ((genFn0() == 0) && v4) {
                for fi116 := 0; fi116 < 9; fi116++ {
                    v9 -= (v2 << ((255) & 63))
                    _ = len(v3)
                    _ = fi116
                }
                if (v8 != v8) {
                    v5 = v5
                    v5 = v5
                    for fi117 := 0; fi117 < 17; fi117++ {
                        v9 = (v9 * v2)
                        if (18446744073709551615 > 2) {
                            var v118 int8 = 100
                            var v119 int64 = (v9 * (v2 & (v2 << ((420) & 63))))
                            _ = v118
                            _ = v119
                        } else {
                            var v120 int16 = -32768
                            var v121 uint64 = 2
                            v122 := GenS8{F0: ("" + v5)}
                            _ = v120
                            _ = v121
                            _ = v122
                        }
                        _ = fi117
                    }
                }
            } else {
                var v123 int = (fi105 >> ((554) & 63))
                for fi124 := 0; fi124 < 6; fi124++ {
                    var v125 int16 = -1
                    _ = fi124
                    _ = v125
                }
                _ = v123
            }
        }
        v5 = v5
        _ = fi105
    }
    var v126 int64 = (v2 * (v2 >> ((266) & 63)))
    fmt.Println(v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v103
    _ = v104
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
