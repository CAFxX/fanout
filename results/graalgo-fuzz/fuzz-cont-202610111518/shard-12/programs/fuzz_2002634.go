package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS3 struct { F0 int; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int16; F2 int64 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS12 struct { F0 int8; F1 uint; F2 int }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int32; F1 int32; F2 uint32 }
type GenS16 struct { F0 int64; F1 int8 }
type GenS17 struct { F0 uint64 }
func genFn0(p0 int16) int16 {
    v7 := []int{1}
    switch 2147483647 {
    case 42:
        if false {
            if (2 <= 255) {
                for fi8 := 0; fi8 < 16; fi8++ {
                    if true {
                        switch "hello" {
                        case "":
                            if (((!true) || false) && false) {
                                p0 -= (p0 * p0)
                            }
                        case "foo bar":
                            v9 := GenS1{F0: p0, F1: -1, F2: 255}
                            _ = v9
                        }
                        for fi10 := 0; fi10 < 1; fi10++ {
                            v11 := GenS1{F0: (p0 & 32767), F1: 1, F2: 65535}
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        var v12 int8 = 127
                        var v13 int8 = v12
                        var v14 uint16 = 866
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi8
                }
                p0 -= -448
            }
            p0 += (p0 | -1)
            var v15 uint8 = 43
            _ = v15
        } else {
            p0 = p0
        }
    case 127:
        p0 = 1
    }
    v16 := GenS1{F0: p0, F1: 0, F2: 463}
    for fi17 := 0; fi17 < 17; fi17++ {
        switch ("" + "foo bar") {
        case "\n\t":
            var v18 uint32 = 220
            _ = v18
        case "a":
            p0 += p0
        default:
            _ = len(v7)
        }
        _ = fi17
    }
    v19 := GenS1{F0: (p0 ^ (p0 << ((4294967295) & 63))), F1: 1, F2: 0}
    _ = p0
    _ = v16
    _ = v19
    _ = v7
    return (p0 >> ((0) & 63))
}

func main() {
    var v1 float32 = (0.0 - ((-2.25 / 1.5) - 1.5))
    v2 := []int{}
    v3 := []int{1, -9223372036854775808, -2147483648, -5}
    var v4 uint64 = 477
    var v5 int8 = -1
    v6 := make(map[string]bool)
    switch 32768 {
    case 1:
        v20 := v6[(("\n\t" + ("\n\t" + (("a" + "foo bar") + "hello"))) + ("a" + "foo bar"))]
        _ = v20
    case 128:
        v4 += 42
    default:
        var v21 uint64 = v4
        _ = v21
    }
    fmt.Println(v4)
    var v22 int64 = 0
    v3 = append(v3, 1)
    switch "" {
    case "hello":
        v23 := GenS1{F0: 42, F1: -2147483648, F2: 421}
        _ = v23
    case "foo bar":
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v24 int32 = 42
            _ = v24
        case -642:
            for fi25 := 0; fi25 < 16; fi25++ {
                if (false && (-2.25 >= (v1 + (3.14159 * (v1 / v1))))) {
                    v3 = append(v3, (fi25 + (fi25 >> ((1) & 63))))
                    if len(v3) > 0 {
                        v3[0] = fi25
                    }
                    v26 := GenS1{F0: -683, F1: -2147483648, F2: 1}
                    _ = v26
                }
                fmt.Println(v22, v4, v22)
                _ = fi25
            }
        case 128:
            v27 := GenS1{F0: 42, F1: -4, F2: 0}
            _ = v27
        default:
            switch "" {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                v4 = v4
            default:
                v28 := GenS2{F0: 32767, F1: 128, F2: -32768}
                _ = v28
            }
        }
    }
    v1 += (v1 / v1)
    switch ("hello" + (("foo bar" + "a") + ("" + ""))) {
    case "\n\t":
        switch "a" {
        case "":
            for fi29 := 0; fi29 < 3; fi29++ {
                v4 = 706
                for fi30 := 0; fi30 < 6; fi30++ {
                    v5 -= 78
                    _ = fi30
                }
                if true {
                    v31 := GenS1{F0: 32767, F1: 0, F2: 0}
                    if (403 != 579) {
                        if true {
                            _ = len(v2)
                            if true {
                                for fi32 := 0; fi32 < 10; fi32++ {
                                    v1 += (1.5 / v1)
                                    v33 := GenS3{F0: fi29, F1: (v22 / ((v22) | 1)), F2: 128}
                                    v5 = (v5 - -12)
                                    fmt.Println(v5, v5, fi32)
                                    fmt.Println(v4)
                                    switch ((((("foo bar" + "\n\t") + (("" + ("a" + ("hello" + ""))) + ("" + "a"))) + "hello") + "a") + "foo bar") {
                                    case "a":
                                        v5 += (v5 | v5)
                                    case "hello":
                                        if ((148 == 42) && (v4 <= (^(v4)))) {
                                            if true {
                                                var v34 string = ("" + "")
                                                switch v34 {
                                                case "\n\t":
                                                    switch (fi29 & 42) {
                                                    case -1:
                                                        v1 += 3.14159
                                                    }
                                                case "a":
                                                    switch 32768 {
                                                    case 127:
                                                        if (true || true) {
                                                            var v35 int16 = 0
                                                            var v36 uint64 = (v4 ^ 247)
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            fmt.Println(v5, fi29)
                                                            _ = len(v2)
                                                            var v37 int16 = -808
                                                            _ = v37
                                                        }
                                                        var v38 string = v34
                                                        if false {
                                                            switch 612 {
                                                            case 9223372036854775807:
                                                                v5 -= (v5 ^ (v5 >> ((700) & 63)))
                                                            case 444:
                                                                if (30 != 1) {
                                                                    v4 += v4
                                                                    if false {
                                                                        if (!(!true)) {
                                                                            fmt.Println(v4, v38)
                                                                            switch "hello" {
                                                                            case "a":
                                                                                delete(v6, v34)
                                                                            default:
                                                                                var v39 bool = (0.1 > v1)
                                                                                _ = v39
                                                                            }
                                                                        } else {
                                                                            var v40 uint32 = 0
                                                                            v22 -= (^(v22))
                                                                            v41 := GenS4{F0: 127, F1: 0, F2: v22}
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                        if (((fi29 % ((int(v22)) | 1)) != (-(fi32))) || true) {
                                                                            v34 = ("" + "hello")
                                                                            v42 := GenS4{F0: 0, F1: -32768, F2: v22}
                                                                            _ = v42
                                                                        }
                                                                        switch (fi29 + fi29) {
                                                                        case 0:
                                                                            switch 1 {
                                                                            case 4294967295:
                                                                                v43 := GenS5{F0: -607, F1: v5}
                                                                                _ = v43
                                                                            case -768:
                                                                                var v44 uint64 = 2
                                                                                _ = v44
                                                                            case 937:
                                                                                v5 = (v5 & (^(v5)))
                                                                            default:
                                                                                if ((v38 >= v34) || (-1 <= 1)) {
                                                                                    switch v38 {
                                                                                    case "":
                                                                                        var v45 int8 = (v5 % (((v5 % ((-1) | 1))) | 1))
                                                                                        _ = v45
                                                                                    default:
                                                                                        var v46 string = v38
                                                                                        var v47 int32 = 0
                                                                                        if (!true) {
                                                                                            var v48 uint = 42
                                                                                            _ = v48
                                                                                        } else {
                                                                                            v34 = v34
                                                                                            var v49 string = v34
                                                                                            _ = v49
                                                                                        }
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    }
                                                                                    v4 -= (v4 - v4)
                                                                                    v3 = append(v3, (fi32 << ((0) & 63)))
                                                                                } else {
                                                                                    v1 += float32(0.0)
                                                                                }
                                                                            }
                                                                        case 32767:
                                                                            var v50 uint32 = 65535
                                                                            _ = v50
                                                                        case 9:
                                                                            v51 := v6[((v38 + v38) + (v34 + ("foo bar" + "hello")))]
                                                                            _ = v51
                                                                        default:
                                                                            v5 = v5
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            var v52 int16 = -416
                                                                            var v53 uint64 = (v4 >> ((1000) & 63))
                                                                            if (42 == 46) {
                                                                                v38 = ("hello" + "foo bar")
                                                                                var v54 uint = 2
                                                                                if false {
                                                                                    v55 := GenS6{F0: false, F1: v54}
                                                                                    _ = v55
                                                                                }
                                                                                if false {
                                                                                    v56 := GenS3{F0: (-(fi32)), F1: (^(v22)), F2: 107}
                                                                                    fmt.Println(v38)
                                                                                    v6[(v38 + "\n\t")] = true
                                                                                    v4 = (v53 | (v53 / (((^(v53))) | 1)))
                                                                                    var v57 float32 = (v1 * (v1 + v1))
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                } else {
                                                                                    v6["hello"] = (!true)
                                                                                    v2 = append(v2, (fi29 / ((42) | 1)))
                                                                                    var v58 string = ("" + v38)
                                                                                    _ = v58
                                                                                }
                                                                                _ = v54
                                                                            } else {
                                                                                if false {
                                                                                    if false {
                                                                                        if true {
                                                                                            v38 = (v38 + v34)
                                                                                            v59 := GenS7{F0: 42, F1: (v22 & v22)}
                                                                                            _ = v59
                                                                                        }
                                                                                        if false {
                                                                                            v4 = (v4 - (v4 % ((v4) | 1)))
                                                                                        } else {
                                                                                            if false {
                                                                                                fmt.Println(v4, v52)
                                                                                                v1 += (v1 / ((0.1 * float32(((0.0 * 1.5) + float64((v1 / (0.1 - v1)))))) / (0.1 / v1)))
                                                                                            }
                                                                                            v53 = (v53 >> ((2) & 63))
                                                                                            var v60 int8 = (^(v5))
                                                                                            _ = v60
                                                                                        }
                                                                                        if ((3.14159 + 0.0) == -2.25) {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = (fi29 >> ((39) & 63))
                                                                                            }
                                                                                        }
                                                                                        var v61 float64 = -2.25
                                                                                        _ = v61
                                                                                    } else {
                                                                                        var v62 uint = 0
                                                                                        v5 += 42
                                                                                        delete(v6, "")
                                                                                        var v63 int64 = 0
                                                                                        v53 -= (^(v4))
                                                                                        _ = v62
                                                                                        _ = v63
                                                                                    }
                                                                                    v64 := GenS3{F0: (fi32 + (fi29 << ((42) & 63))), F1: (v22 >> ((4294967295) & 63)), F2: 233}
                                                                                    v65 := GenS2{F0: 513, F1: 127, F2: 1}
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                } else {
                                                                                    v22 = v22
                                                                                    var v66 float32 = -2.25
                                                                                    _ = v66
                                                                                }
                                                                                v67 := GenS6{F0: (!(((3.14159 + 0.0) >= 1.5) && (-106 > v5))), F1: 4294967295}
                                                                                _ = v67
                                                                            }
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        v38 = "\n\t"
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            v6[v34] = false
                                                                            var v68 int64 = 0
                                                                            var v69 int8 = 103
                                                                            _ = v68
                                                                            _ = v69
                                                                        default:
                                                                            var v70 int8 = v5
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                    v34 = ""
                                                                }
                                                            case -2147483648:
                                                                v4 -= v4
                                                            default:
                                                                fmt.Println(fi29, fi32, v38)
                                                            }
                                                            var v71 uint8 = 255
                                                            if true {
                                                                if (float64(1e10) >= 1.5) {
                                                                    if true {
                                                                        fmt.Println(v4, v5, v22)
                                                                    } else {
                                                                        _ = len(v6)
                                                                        var v72 int16 = 0
                                                                        v4 -= 0
                                                                        _ = v72
                                                                    }
                                                                } else {
                                                                    var v73 int32 = 1
                                                                    var v74 uint8 = 127
                                                                    v75 := v6["foo bar"]
                                                                    _ = v73
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                v22 += v22
                                                                delete(v6, v34)
                                                            }
                                                            _ = v71
                                                        } else {
                                                            var v76 int8 = v5
                                                            _ = v76
                                                        }
                                                        _ = v38
                                                    case 2147483647:
                                                        if true {
                                                            if (true || (975 < 905)) {
                                                                v5 -= v5
                                                                var v77 float32 = v1
                                                                _ = v77
                                                            }
                                                        }
                                                    case 1:
                                                        if (1 >= 410) {
                                                            if false {
                                                                v6[((v34 + (v34 + "hello")) + (((v34 + (v34 + "\n\t")) + "\n\t") + v34))] = (!(1.5 <= float64(float32(float64(v1)))))
                                                                var v78 int64 = v22
                                                                switch fi32 {
                                                                case 1:
                                                                    var v79 uint8 = 2
                                                                    _ = v79
                                                                case 32767:
                                                                    var v80 uint64 = (v4 << ((612) & 63))
                                                                    _ = v80
                                                                case -1:
                                                                    switch v34 {
                                                                    case "\n\t":
                                                                        if false {
                                                                            v81 := v6[(v34 + v34)]
                                                                            _ = v81
                                                                        }
                                                                    case "":
                                                                        v82 := GenS3{F0: fi32, F1: (v78 % ((-608) | 1)), F2: 1}
                                                                        _ = v82
                                                                    case "foo bar":
                                                                        v83 := GenS5{F0: -9223372036854775808, F1: (v5 + (v5 % ((100) | 1)))}
                                                                        var v84 float32 = v1
                                                                        _ = len(v6)
                                                                        if true {
                                                                            if (-2.25 == ((((((1.5 / (1.5 + 0.0)) / v1) - v1) + 0.0) * v84) * v84)) {
                                                                                var v85 uint32 = 42
                                                                                var v86 uint16 = 32767
                                                                                v87 := GenS8{F0: 65535, F1: (v5 & -128)}
                                                                                _ = v85
                                                                                _ = v86
                                                                                _ = v87
                                                                            }
                                                                            v88 := GenS2{F0: 42, F1: 255, F2: -32768}
                                                                            _ = v88
                                                                        }
                                                                        v89 := GenS7{F0: 260, F1: v78}
                                                                        _ = v83
                                                                        _ = v84
                                                                        _ = v89
                                                                    }
                                                                default:
                                                                    if (!(18446744073709551615 != 47)) {
                                                                        v90 := v6[""]
                                                                        _ = v90
                                                                    }
                                                                }
                                                                if false {
                                                                    v78 = v78
                                                                } else {
                                                                    v91 := GenS9{F0: 1, F1: ("\n\t" + "hello")}
                                                                    v6[(v34 + v34)] = (!false)
                                                                    if false {
                                                                        var v92 int32 = 922
                                                                        fmt.Println(v78, v22, v5)
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v92
                                                                    }
                                                                    _ = v91
                                                                }
                                                                _ = v78
                                                            }
                                                            v22 = (v22 ^ (-(v22)))
                                                        }
                                                    }
                                                case "hello":
                                                    if (false && (float32(3.14159) >= (float32(((float64(1e10) - 0.0) + (((3.14159 + (0.0 * 1.5)) / ((((((1e10 / 1.0) * 1e10) / 0.1) - 0.0) / 1.0) * 0.1)) * (1e10 / float64(((v1 - (1.5 + v1)) - (v1 - (((v1 - (v1 * ((v1 / v1) + -2.25))) + 0.0) / 3.14159)))))))) * -2.25))) {
                                                        v93 := GenS8{F0: 511, F1: (^(v5))}
                                                        _ = v93
                                                    }
                                                }
                                                if false {
                                                    v5 = 100
                                                } else {
                                                    if (1.5 <= 3.14159) {
                                                        v5 -= int8(v4)
                                                        v22 = v22
                                                        if false {
                                                            v94 := GenS4{F0: 127, F1: -32768, F2: 1}
                                                            v95 := GenS10{F0: 388, F1: 42}
                                                            _ = v94
                                                            _ = v95
                                                        } else {
                                                            var v96 bool = (!true)
                                                            if v96 {
                                                                v96 = v96
                                                                var v97 float64 = 1.5
                                                                _ = v97
                                                            } else {
                                                                v98 := GenS1{F0: 483, F1: 1, F2: 65535}
                                                                _ = v98
                                                            }
                                                            _ = v96
                                                        }
                                                        v99 := GenS9{F0: int32(fi32), F1: v34}
                                                        _ = v99
                                                    } else {
                                                        var v100 uint16 = 0
                                                        _ = v100
                                                    }
                                                    if (!true) {
                                                        if ((v5 - (v5 + v5)) != (v5 % (((v5 << ((233) & 63))) | 1))) {
                                                            fmt.Println(v1)
                                                        }
                                                        var v101 float64 = 0.0
                                                        v34 = "hello"
                                                        v4 = (v4 % (((v4 << ((922) & 63))) | 1))
                                                        _ = v101
                                                    }
                                                    if false {
                                                        delete(v6, v34)
                                                        v102 := GenS11{F0: 164, F1: 4294967295, F2: 2147483647}
                                                        v1 -= (v1 / (-2.25 * ((3.14159 / (((v1 - v1) - ((v1 / (v1 / float32(((-2.25 + 1.5) - 1e10)))) - float32(1e10))) + ((-2.25 * 1.5) / 0.1))) + 3.14159)))
                                                        _ = v102
                                                    } else {
                                                        v103 := v6[("a" + (v34 + v34))]
                                                        switch ("a" + (v34 + (v34 + v34))) {
                                                        case "":
                                                            var v104 uint = 4294967295
                                                            _ = v104
                                                        }
                                                        _ = v103
                                                    }
                                                }
                                                _ = v34
                                            } else {
                                                v105 := GenS6{F0: false, F1: 756}
                                                if false {
                                                    v22 += v22
                                                    if false {
                                                        if len(v2) > 0 {
                                                            v2[0] = fi32
                                                        }
                                                        var v106 int8 = 34
                                                        v1 += (float32((0.1 * 0.1)) - v1)
                                                        _ = v106
                                                    }
                                                    if (84 >= 0) {
                                                        var v107 int8 = (-(v5))
                                                        _ = v107
                                                    } else {
                                                        var v108 int8 = v5
                                                        v22 = (v22 >> ((231) & 63))
                                                        if true {
                                                            _ = len(v6)
                                                            if true {
                                                                if true {
                                                                    fmt.Println(fi29, v22, v1)
                                                                    if true {
                                                                        fmt.Println(v22, v108, fi29)
                                                                        var v109 string = "hello"
                                                                        v109 = v109
                                                                        _ = v109
                                                                    } else {
                                                                        if (127 < 71) {
                                                                            switch (fi32 & (-(fi29))) {
                                                                            case 32768:
                                                                                if true {
                                                                                    if (255 >= 367) {
                                                                                        v110 := GenS12{F0: 0, F1: 1, F2: (fi29 % (((fi29 / ((fi32) | 1))) | 1))}
                                                                                        v4 -= v4
                                                                                        v111 := GenS10{F0: -1, F1: 32767}
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    }
                                                                                    v108 -= -81
                                                                                    v6["hello"] = false
                                                                                }
                                                                            }
                                                                            var v112 uint = 0
                                                                            _ = len(v3)
                                                                            _ = v112
                                                                        }
                                                                    }
                                                                    _ = len(v3)
                                                                } else {
                                                                    _ = len(v6)
                                                                    fmt.Println(v1, fi29)
                                                                    fmt.Println(v22, fi29, v22)
                                                                    v108 = v108
                                                                }
                                                            } else {
                                                                if (!true) {
                                                                    var v113 float64 = 0.0
                                                                    _ = v113
                                                                } else {
                                                                    if false {
                                                                        var v114 uint = 1
                                                                        v1 -= 1e10
                                                                        _ = v114
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            v6["\n\t"] = (!true)
                                                                        default:
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                fmt.Println(v1)
                                                                                if false {
                                                                                    v22 = v22
                                                                                    _ = len(v6)
                                                                                }
                                                                                v108 -= 63
                                                                            default:
                                                                                v115 := GenS13{F0: 42}
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                        v3 = append(v3, (fi32 % ((fi29) | 1)))
                                                                        if (-331 > 838) {
                                                                            v4 = (v4 >> ((1) & 63))
                                                                            var v116 int16 = genFn0(32767)
                                                                            if ("a" != "a") {
                                                                                var v117 int8 = v108
                                                                                v118 := GenS14{F0: 1}
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                            _ = v116
                                                                        } else {
                                                                            fmt.Println(v108, v4)
                                                                            var v119 int64 = (^(v22))
                                                                            v22 = (^(v119))
                                                                            _ = v119
                                                                        }
                                                                    }
                                                                }
                                                                v120 := GenS11{F0: uint8(v4), F1: 4294967295, F2: 1}
                                                                _ = v120
                                                            }
                                                        } else {
                                                            switch "\n\t" {
                                                            case "":
                                                                switch 976 {
                                                                case 565:
                                                                    v121 := GenS15{F0: 106, F1: -1, F2: 4294967295}
                                                                    _ = v121
                                                                case 32768:
                                                                    _ = len(v3)
                                                                case 127:
                                                                    switch fi29 {
                                                                    case 9223372036854775807:
                                                                        v1 = v1
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (-(fi32))
                                                                        }
                                                                        v22 = -1
                                                                    case -51:
                                                                        v122 := GenS8{F0: 65535, F1: 85}
                                                                        _ = v122
                                                                    default:
                                                                        v108 = -6
                                                                    }
                                                                }
                                                            case "a":
                                                                v22 += -829
                                                            case "hello":
                                                                v123 := GenS16{F0: 0, F1: (v5 ^ v108)}
                                                                _ = v123
                                                            default:
                                                                fmt.Println(v108, v4, v1)
                                                            }
                                                        }
                                                        _ = len(v6)
                                                        _ = v108
                                                    }
                                                } else {
                                                    var v124 uint16 = 955
                                                    v125 := GenS6{F0: ((463 == 65535) && (!true)), F1: uint(v124)}
                                                    if ((1 > 2) && ((v4 % ((v4) | 1)) != v4)) {
                                                        if (false && (652 == 1)) {
                                                            v126 := GenS9{F0: 0, F1: "\n\t"}
                                                            v6[("foo bar" + (("" + "hello") + (("\n\t" + ("" + "foo bar")) + "foo bar")))] = false
                                                            _ = v126
                                                        } else {
                                                            switch ("\n\t" + "a") {
                                                            case "hello":
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v127 := GenS7{F0: v124, F1: v22}
                                                                    _ = v127
                                                                case "a":
                                                                    v1 -= (((v1 + v1) - v1) + float32(0.0))
                                                                case "\n\t":
                                                                    if false {
                                                                        fmt.Println(v22)
                                                                    }
                                                                default:
                                                                    v128 := GenS17{F0: v4}
                                                                    _ = v128
                                                                }
                                                            case "\n\t":
                                                                v5 = (-(v5))
                                                            case "a":
                                                                v5 += v5
                                                            }
                                                            if (("hello" < (("" + "hello") + "\n\t")) && (1 != 0)) {
                                                                if len(v2) > 0 {
                                                                    v2[0] = (fi29 % ((fi29) | 1))
                                                                }
                                                                v1 += v1
                                                                v129 := GenS10{F0: -230, F1: 42}
                                                                _ = v129
                                                            } else {
                                                                var v130 uint8 = 0
                                                                _ = v130
                                                            }
                                                        }
                                                        switch ("\n\t" + "\n\t") {
                                                        case "hello":
                                                            var v131 uint32 = 2147483648
                                                            _ = v131
                                                        case "foo bar":
                                                            var v132 uint = 2
                                                            _ = v132
                                                        case "a":
                                                            var v133 string = ((("a" + ("" + "foo bar")) + ("foo bar" + "foo bar")) + ("" + "foo bar"))
                                                            _ = v133
                                                        }
                                                        _ = len(v2)
                                                        v22 = v22
                                                        delete(v6, "foo bar")
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v134 int = (fi32 / (((fi29 >> ((18446744073709551615) & 63))) | 1))
                                                    _ = v124
                                                    _ = v125
                                                    _ = v134
                                                }
                                                var v135 int = fi32
                                                _ = v105
                                                _ = v135
                                            }
                                        }
                                    default:
                                        v2 = append(v2, 9223372036854775807)
                                    }
                                    v4 = v4
                                    _ = fi32
                                    _ = v33
                                }
                                v4 -= 314
                                if ((^(v5)) != (-(v5))) {
                                    v1 += ((v1 * 0.1) * float32((1e10 - ((0.1 - 1e10) * (3.14159 - ((1.5 + 1e10) * ((1.5 - 0.0) + (0.1 * ((1.5 / 0.1) / (3.14159 * (1e10 + 0.0)))))))))))
                                }
                            }
                        }
                    }
                    v136 := GenS12{F0: (v5 % ((127) | 1)), F1: 1, F2: int(v22)}
                    _ = v136
                    _ = v31
                }
                _ = fi29
            }
        case "a":
            switch -2147483648 {
            case -1:
                var v137 uint64 = 1
                _ = v137
            default:
                v1 = ((v1 + 3.14159) - float32(1e10))
            }
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    default:
        v4 = v4
    }
    v22 = v22
    fmt.Println(v1, v5, v5)
    switch 32768 {
    case 1:
        for fi138 := 0; fi138 < 6; fi138++ {
            var v139 uint32 = 2147483648
            _ = fi138
            _ = v139
        }
    case 127:
        v22 = (v22 % (((v22 & v22)) | 1))
    case 32767:
        v5 = (v5 - (v5 + 0))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
