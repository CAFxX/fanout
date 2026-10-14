package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 string; F2 string }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 string; F1 int32; F2 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 string; F1 int64; F2 int }
type GenS11 struct { F0 int32 }
func genFn0(p0 string) int {
    if true {
        p0 = (p0 + p0)
    }
    if (true || true) {
        for fi7 := 0; fi7 < 9; fi7++ {
            if (549 > 42) {
                switch fi7 {
                case 4294967295:
                    var v8 uint64 = 0
                    _ = v8
                case 32768:
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 uint = 2
                        v11 := []int{}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
                var v12 float64 = 1.5
                _ = v12
            }
            var v13 int8 = 0
            _ = fi7
            _ = v13
        }
        switch -303 {
        case 0:
            for fi14 := 0; fi14 < 6; fi14++ {
                p0 = "hello"
                _ = fi14
            }
        default:
            var v15 float64 = 3.14159
            _ = v15
        }
    } else {
        var v16 uint8 = 2
        v17 := []int{}
        _ = v16
        _ = v17
    }
    p0 = (p0 + (((p0 + p0) + "a") + "a"))
    v18 := []int{1, 461, 42, -290}
    v19 := GenS1{F0: 176, F1: 9}
    _ = p0
    _ = v18
    _ = v19
    return 1
}

func main() {
    var v1 int8 = -81
    v2 := []int{966}
    v3 := make(map[string]string)
    var v4 uint64 = 0
    v5 := make(map[int]uint32)
    var v6 uint16 = uint16(v4)
    if true {
        delete(v5, -716)
        v1 = (v1 / (((v1 * 127)) | 1))
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int = (^(fi20))
            _ = fi20
            _ = v21
        }
    } else {
        fmt.Println(v4, v6)
    }
    if (1 >= -933) {
        v1 = v1
    } else {
        if len(v2) > 0 {
            v2[0] = 464
        }
        var v22 uint32 = 65535
        fmt.Println(v6, v22, v4)
        _ = v22
    }
    delete(v3, ("foo bar" + ("foo bar" + "")))
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            v4 -= v4
            _ = fi24
        }
        var v25 int8 = 100
        v25 += 100
        _ = fi23
        _ = v25
    }
    v26 := GenS2{F0: 56}
    for fi27 := 0; fi27 < 13; fi27++ {
        switch 98 {
        case 32767:
            v1 = v1
        case 2147483647:
            v1 += v1
        }
        _ = fi27
    }
    var v28 uint32 = 4294967295
    switch 4294967295 {
    case 145:
        switch -1 {
        case 1:
            for fi29 := 0; fi29 < 14; fi29++ {
                v2 = append(v2, (fi29 + 83))
                if true {
                    v30 := GenS3{F0: true}
                    fmt.Println(fi29, fi29, v1)
                    var v31 float32 = (1.5 + 1.5)
                    _ = v30
                    _ = v31
                } else {
                    v28 += 359
                    fmt.Println(v6, v1, v1)
                }
                _ = fi29
            }
        case -900:
            var v32 uint16 = v6
            _ = v32
        case 32768:
            _ = len(v2)
        default:
            v4 += 0
        }
    case -9223372036854775808:
        _ = len(v3)
    }
    v3[""] = ""
    var v33 float64 = 1e10
    var v34 uint = 2
    var v35 float32 = 3.14159
    v36 := GenS1{F0: uint8(v4), F1: v1}
    switch "" {
    case "hello":
        if false {
            if (-9 < 9223372036854775807) {
                v35 = v35
                var v37 uint = 266
                _ = v37
            } else {
                for fi38 := 0; fi38 < 5; fi38++ {
                    var v39 uint8 = 127
                    v33 = (float64(float32(1e10)) + v33)
                    _ = fi38
                    _ = v39
                }
                var v40 int16 = -32768
                v41 := GenS4{F0: v40, F1: ("a" + ("\n\t" + ("hello" + (("\n\t" + ((("\n\t" + (("foo bar" + "hello") + "\n\t")) + "hello") + "hello")) + "foo bar")))), F2: "foo bar"}
                _ = v40
                _ = v41
            }
        } else {
            var v42 uint = (v34 / ((v34) | 1))
            _ = v42
        }
    case "a":
        v1 += (v1 >> (((v34 - (v34 & v34))) & 63))
    case "foo bar":
        if false {
            v28 += (v28 * v28)
            var v43 uint32 = 0
            v44 := GenS5{F0: 1, F1: 127}
            _ = v43
            _ = v44
        } else {
            v45 := GenS4{F0: 42, F1: "hello", F2: ("foo bar" + ("hello" + "a"))}
            v46 := v3[((("" + "hello") + ("hello" + "hello")) + "foo bar")]
            switch -693 {
            case 128:
                if (v1 != (v1 | (v1 % ((v1) | 1)))) {
                    v35 = -2.25
                } else {
                    v5[-2147483648] = v28
                    v47 := GenS4{F0: 32767, F1: v46, F2: v46}
                    _ = v47
                }
            default:
                if (!(1000 > 42)) {
                    v6 = (v6 * v6)
                    fmt.Println(v33, v35)
                    v48 := GenS1{F0: 127, F1: (v1 / ((v1) | 1))}
                    _ = v48
                } else {
                    var v49 int = 1
                    _ = v49
                }
            }
            _ = v45
            _ = v46
        }
    default:
        switch -352 {
        case -1:
            if false {
                for fi50 := 0; fi50 < 4; fi50++ {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        v28 += (^(v28))
                        v52 := GenS6{F0: (("foo bar" + "\n\t") + "foo bar"), F1: (true || (!true))}
                        v53 := v3["foo bar"]
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    var v54 uint = v34
                    _ = fi50
                    _ = v54
                }
            } else {
                var v55 uint32 = (^(v28))
                for fi56 := 0; fi56 < 15; fi56++ {
                    if ((("a" != (("" + "hello") + ("\n\t" + ((("\n\t" + "\n\t") + "a") + "foo bar")))) && (9223372036854775807 <= 0)) || true) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v4 = (v4 + (^(v4)))
                        switch "a" {
                        case "":
                            switch 4294967295 {
                            case 0:
                                v33 -= (0.0 * (float64(v35) - float64(v35)))
                            default:
                                for fi57 := 0; fi57 < 12; fi57++ {
                                    var v58 int8 = 100
                                    var v59 uint16 = v6
                                    v60 := GenS6{F0: "\n\t", F1: true}
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                            }
                        case "foo bar":
                            if (((float64((v35 * v35)) / float64(v35)) + v33) >= 1.5) {
                                for fi61 := 0; fi61 < 8; fi61++ {
                                    if true {
                                        v62 := GenS1{F0: 143, F1: v1}
                                        _ = v62
                                    }
                                    v63 := GenS7{F0: "", F1: 0, F2: v4}
                                    delete(v5, fi56)
                                    _ = fi61
                                    _ = v63
                                }
                            } else {
                                v35 = v35
                                fmt.Println(v4, v34)
                                v28 = (v55 | v55)
                            }
                        case "\n\t":
                            if (59 <= (^(v1))) {
                                v64 := GenS8{F0: v1}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v64
                            }
                        default:
                            for fi65 := 0; fi65 < 17; fi65++ {
                                var v66 uint32 = 4294967295
                                if (!false) {
                                    v67 := GenS6{F0: "a", F1: true}
                                    v28 = (v28 << ((v34) & 63))
                                    if (!(!(!false))) {
                                        if true {
                                            var v68 int32 = int32(v66)
                                            var v69 bool = true
                                            _ = len(v2)
                                            _ = v68
                                            _ = v69
                                        } else {
                                            switch "\n\t" {
                                            case "a":
                                                var v70 int8 = int8(v66)
                                                _ = v70
                                            case "hello":
                                                switch (fi56 - (fi56 - 4294967295)) {
                                                case -214:
                                                    if true {
                                                        switch "a" {
                                                        case "a":
                                                            switch ("\n\t" + "") {
                                                            case "\n\t":
                                                                v28 = (v66 - v66)
                                                            case "a":
                                                                v3["foo bar"] = "a"
                                                            case "foo bar":
                                                                var v71 uint64 = v4
                                                                _ = v71
                                                            }
                                                        case "\n\t":
                                                            var v72 float32 = v35
                                                            _ = v72
                                                        default:
                                                            v66 = 2147483647
                                                        }
                                                    }
                                                default:
                                                    switch ("" + "foo bar") {
                                                    case "a":
                                                        var v73 float32 = v35
                                                        _ = v73
                                                    case "hello":
                                                        var v74 string = ("a" + "")
                                                        _ = v74
                                                    case "":
                                                        var v75 float32 = (((v35 / -2.25) / ((v35 - (-2.25 - 0.1)) * 1.5)) - (v35 / ((((3.14159 * v35) * ((v35 * 0.0) / 0.1)) - 0.0) + v35)))
                                                        _ = v75
                                                    default:
                                                        v55 = 844
                                                    }
                                                }
                                                var v76 int64 = 9223372036854775807
                                                var v77 float32 = v35
                                                _ = v76
                                                _ = v77
                                            default:
                                                switch fi56 {
                                                case 0:
                                                    switch (fi56 & fi56) {
                                                    case 2147483647:
                                                        fmt.Println(v35)
                                                    case 9223372036854775807:
                                                        var v78 float64 = (v33 - v33)
                                                        _ = v78
                                                    default:
                                                        var v79 float32 = ((3.14159 - v35) * v35)
                                                        _ = v79
                                                    }
                                                case -511:
                                                    if true {
                                                        v2 = append(v2, fi56)
                                                    }
                                                default:
                                                    var v80 int8 = (v1 * (v1 % (((-(v1))) | 1)))
                                                    _ = v80
                                                }
                                                if false {
                                                    if false {
                                                        v81 := GenS1{F0: 0, F1: (-(v1))}
                                                        delete(v5, genFn0(""))
                                                        if true {
                                                            v3[("hello" + ("foo bar" + "a"))] = ("hello" + ("\n\t" + "hello"))
                                                        }
                                                        var v82 bool = (1e10 >= ((v33 - 0.0) - 0.1))
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                    v5[fi56] = v66
                                                }
                                            }
                                        }
                                        v1 += 1
                                        v6 -= (v6 * 65535)
                                    }
                                    v35 -= 1e10
                                    switch "hello" {
                                    case "hello":
                                        switch int(v66) {
                                        case 4294967295:
                                            var v83 uint64 = (v4 & v4)
                                            _ = v83
                                        case -718:
                                            v1 = 95
                                        default:
                                            var v84 uint16 = v6
                                            _ = v84
                                        }
                                    default:
                                        switch int(v6) {
                                        case 127:
                                            v66 = (v66 | 65535)
                                            v6 -= (v6 % (((v6 - (v6 ^ (^(v6))))) | 1))
                                            if true {
                                                var v85 int = (fi65 & 127)
                                                v3["\n\t"] = (((("foo bar" + "") + "hello") + "hello") + "\n\t")
                                                switch int(v6) {
                                                case -9223372036854775808:
                                                    fmt.Println(v4)
                                                case 1:
                                                    v28 = (^(v28))
                                                }
                                                var v86 float32 = 0.1
                                                v86 = (3.14159 + (float32((((1.5 + float64((float32(v33) - 1e10))) / float64(v86)) - v33)) + (v86 / 1.5)))
                                                var v87 uint8 = 2
                                                _ = v85
                                                _ = v86
                                                _ = v87
                                            } else {
                                                _ = len(v5)
                                            }
                                            v88 := GenS9{F0: 1}
                                            if false {
                                                v66 = 65535
                                                v6 = 255
                                                fmt.Println(v35, v1)
                                            } else {
                                                v89 := GenS10{F0: ((((("foo bar" + ("\n\t" + "\n\t")) + ("" + ((("foo bar" + ("a" + "hello")) + ("\n\t" + "hello")) + ("foo bar" + "hello")))) + ("a" + "hello")) + ("foo bar" + ("" + "a"))) + (("a" + ("\n\t" + "")) + "hello")), F1: 1, F2: -1}
                                                var v90 float32 = (v35 / (v35 / -2.25))
                                                if (!(-32768 <= 1)) {
                                                    v90 = ((float32(v33) / v90) / ((v35 + (v90 - (v35 + v90))) * float32(-2.25)))
                                                }
                                                _ = v89
                                                _ = v90
                                            }
                                            _ = v88
                                        case 2147483647:
                                            var v91 int8 = v1
                                            _ = v91
                                        }
                                        if false {
                                            if (!(-1 != 42)) {
                                                var v92 uint16 = (v6 - v6)
                                                switch (("" + (("\n\t" + "hello") + "\n\t")) + "") {
                                                case "":
                                                    fmt.Println(v34, v28)
                                                case "a":
                                                    delete(v3, "hello")
                                                    v93 := GenS6{F0: ((("a" + (("a" + "foo bar") + ("a" + ("hello" + "hello")))) + "") + ("" + "hello")), F1: true}
                                                    _ = v93
                                                case "\n\t":
                                                    if true {
                                                        v66 += (v55 - (v66 & v55))
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi65 % ((565) | 1))
                                                        }
                                                        var v94 int16 = 1000
                                                        _ = v94
                                                    }
                                                default:
                                                    if true {
                                                        var v95 float64 = -2.25
                                                        if false {
                                                            var v96 uint64 = 594
                                                            _ = v96
                                                        } else {
                                                            var v97 uint32 = (v66 << (((v34 & (v34 + (v34 ^ (v34 + v34))))) & 63))
                                                            var v98 uint8 = 249
                                                            _ = v97
                                                            _ = v98
                                                        }
                                                        fmt.Println(v1, v95)
                                                        _ = v95
                                                    }
                                                }
                                                if true {
                                                    fmt.Println(v34, v92)
                                                } else {
                                                    switch (fi65 % (((fi65 >> ((uint(v4)) & 63))) | 1)) {
                                                    case 128:
                                                        var v99 float64 = (0.1 / v33)
                                                        _ = v99
                                                    case 1:
                                                        v100 := GenS10{F0: "\n\t", F1: 69, F2: (fi56 | -796)}
                                                        _ = v100
                                                    }
                                                    var v101 int32 = -492
                                                    var v102 int64 = 42
                                                    _ = v101
                                                    _ = v102
                                                }
                                                _ = v92
                                            }
                                        }
                                    }
                                    var v103 float64 = 1.5
                                    switch fi65 {
                                    case 2147483647:
                                        var v104 bool = false
                                        _ = v104
                                    case -9223372036854775808:
                                        var v105 uint8 = 2
                                        _ = v105
                                    default:
                                        v106 := GenS11{F0: 157}
                                        _ = v106
                                    }
                                    _ = v103
                                    _ = v67
                                } else {
                                    _ = len(v2)
                                    switch "hello" {
                                    case "foo bar":
                                        var v107 bool = (true && (192 <= (v66 & v55)))
                                        _ = v107
                                    case "hello":
                                        var v108 uint32 = v55
                                        _ = v108
                                    default:
                                        if ((uint64(v66) > v4) || true) {
                                            v3["a"] = ""
                                            v33 += ((v33 / v33) - float64((v35 + v35)))
                                            fmt.Println(fi56, v6)
                                        }
                                    }
                                    v33 -= 3.14159
                                }
                                _ = fi65
                                _ = v66
                            }
                        }
                    }
                    switch (fi56 * genFn0((((("foo bar" + ("" + "hello")) + "hello") + (("a" + "\n\t") + "\n\t")) + (("a" + "hello") + ("hello" + "foo bar"))))) {
                    case 593:
                        v55 -= 4294967295
                    case 127:
                        v33 -= v33
                    }
                    for fi109 := 0; fi109 < 6; fi109++ {
                        var v110 bool = true
                        fmt.Println(v110, v28, v33)
                        _ = fi109
                        _ = v110
                    }
                    _ = fi56
                }
                _ = v55
            }
        case -339:
            fmt.Println(v34, v28, v35)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
