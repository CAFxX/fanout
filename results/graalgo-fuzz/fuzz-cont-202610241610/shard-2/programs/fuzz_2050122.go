package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int }

func main() {
    v1 := []int{1, 32768, 4294967295, -784}
    var v2 int32 = -2147483648
    v3 := make(map[int]bool)
    var v4 int64 = 0
    switch "\n\t" {
    case "":
        var v5 float64 = 1e10
        _ = v5
    case "foo bar":
        if ((3.14159 - 3.14159) < 3.14159) {
            switch 764 {
            case -1:
                var v6 uint64 = 152
                _ = v6
            case -115:
                var v7 string = "foo bar"
                _ = v7
            case 567:
                switch 9223372036854775807 {
                case 42:
                    for fi8 := 0; fi8 < 19; fi8++ {
                        var v9 uint32 = 0
                        if true {
                            _ = len(v1)
                            var v10 int = 2147483647
                            if ((!(1 < 32768)) || false) {
                                v11 := GenS1{F0: 133, F1: (v4 ^ v4), F2: 568}
                                _ = v11
                            }
                            _ = v10
                        } else {
                            switch "a" {
                            case "hello":
                                v12 := v3[(fi8 ^ 4294967295)]
                                _ = v12
                            case "foo bar":
                                var v13 string = ("\n\t" + (("hello" + "\n\t") + "hello"))
                                _ = v13
                            default:
                                v1 = append(v1, 48)
                            }
                            fmt.Println(fi8)
                        }
                        _ = fi8
                        _ = v9
                    }
                case -532:
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 int64 = int64(v2)
                        v4 = 1
                        if false {
                            var v16 string = ("a" + "foo bar")
                            switch fi14 {
                            case 4294967295:
                                var v17 uint32 = 930
                                _ = v17
                            case 128:
                                delete(v3, (fi14 / ((32768) | 1)))
                            case 0:
                                switch fi14 {
                                case 229:
                                    v16 = ""
                                }
                            default:
                                fmt.Println(v4)
                            }
                            var v18 int8 = 89
                            _ = v16
                            _ = v18
                        } else {
                            fmt.Println(v2)
                            if (1 != (-(fi14))) {
                                if (-1 >= -32768) {
                                    var v19 float32 = 0.1
                                    _ = v19
                                } else {
                                    if true {
                                        fmt.Println(fi14)
                                    }
                                    switch ("foo bar" + ((("\n\t" + "a") + "") + ("foo bar" + "foo bar"))) {
                                    case "\n\t":
                                        for fi20 := 0; fi20 < 14; fi20++ {
                                            _ = len(v1)
                                            var v21 uint16 = 42
                                            v2 -= (v2 | (v2 ^ v2))
                                            _ = fi20
                                            _ = v21
                                        }
                                    case "a":
                                        v2 = v2
                                    case "hello":
                                        switch (fi14 / (((fi14 ^ (fi14 | fi14))) | 1)) {
                                        case 2147483647:
                                            fmt.Println(fi14)
                                        case -706:
                                            v22 := GenS2{F0: v15, F1: fi14}
                                            _ = v22
                                        default:
                                            var v23 int64 = v15
                                            _ = v23
                                        }
                                    default:
                                        var v24 int = -9223372036854775808
                                        _ = v24
                                    }
                                    if false {
                                        v25 := GenS2{F0: v15, F1: int(v15)}
                                        v26 := v3[32767]
                                        if true {
                                            if (!true) {
                                                _ = len(v3)
                                                for fi27 := 0; fi27 < 7; fi27++ {
                                                    if (!(((2 > 0) && (!(!((float64(((0.0 * (1.5 - float32(1e10))) / 1.0)) / (-2.25 * 1.5)) > ((0.0 / (1e10 / ((0.0 - -2.25) + float64((3.14159 - (-2.25 / 0.1)))))) - (1e10 - (float64(3.14159) + ((((3.14159 / (3.14159 * 0.1)) + 0.1) + (-2.25 / 1e10)) + ((0.0 * ((0.1 + ((1e10 + (1.5 * float64(((0.0 - 1.5) + (3.14159 + 0.0))))) / (1e10 + 1.5))) / 1.0)) + 0.1))))))))) || (!((3.14159 / 1.5) > (0.0 + 0.0))))) {
                                                        var v28 int8 = int8(v2)
                                                        if false {
                                                            if false {
                                                                fmt.Println(v15, v4)
                                                            } else {
                                                                v1 = append(v1, (fi27 >> ((65535) & 63)))
                                                                var v29 float64 = 1.5
                                                                _ = v29
                                                            }
                                                            v1 = append(v1, (-(fi27)))
                                                            if (1 > 42) {
                                                                v26 = v26
                                                            } else {
                                                                v4 -= v4
                                                                switch 0 {
                                                                case -2147483648:
                                                                    var v30 string = ""
                                                                    v26 = false
                                                                    var v31 string = ((v30 + v30) + ((v30 + ((v30 + (("" + "foo bar") + (v30 + v30))) + (v30 + "a"))) + (((("\n\t" + ("" + v30)) + v30) + v30) + (v30 + v30))))
                                                                    _ = v30
                                                                    _ = v31
                                                                case 32768:
                                                                    if false {
                                                                        v4 += (v15 | v15)
                                                                        v32 := GenS3{F0: 42, F1: 793}
                                                                        v4 += v15
                                                                        _ = v32
                                                                    }
                                                                case 1:
                                                                    var v33 int32 = (v2 + (-(v2)))
                                                                    _ = v33
                                                                }
                                                            }
                                                        }
                                                        _ = v28
                                                    } else {
                                                        if true {
                                                            var v34 int = fi27
                                                            v35 := GenS1{F0: -167, F1: (v15 << ((42) & 63)), F2: v2}
                                                            var v36 string = "foo bar"
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            v2 -= 2147483647
                                                            v2 = (v2 + 2147483647)
                                                            var v37 uint8 = 127
                                                            _ = v37
                                                        }
                                                        delete(v3, fi14)
                                                        var v38 uint64 = 418
                                                        _ = v38
                                                    }
                                                    var v39 int = fi27
                                                    _ = fi27
                                                    _ = v39
                                                }
                                            }
                                            switch (fi14 ^ -2147483648) {
                                            case 32768:
                                                switch (^(fi14)) {
                                                case 32768:
                                                    fmt.Println(v2)
                                                case 2147483647:
                                                    if false {
                                                        for fi40 := 0; fi40 < 1; fi40++ {
                                                            v41 := GenS1{F0: int16(fi14), F1: v15, F2: v2}
                                                            v2 = v2
                                                            v42 := GenS4{F0: fi14, F1: 2, F2: (v15 << ((0) & 63))}
                                                            if v26 {
                                                                switch 287 {
                                                                case -458:
                                                                    switch ("" + ("\n\t" + ("foo bar" + "hello"))) {
                                                                    case "\n\t":
                                                                        fmt.Println(fi40, fi40, fi14)
                                                                    case "hello":
                                                                        var v43 string = "a"
                                                                        _ = v43
                                                                    case "a":
                                                                        var v44 uint64 = 0
                                                                        _ = v44
                                                                    default:
                                                                        v45 := GenS5{F0: 0}
                                                                        _ = v45
                                                                    }
                                                                case 405:
                                                                    v26 = v26
                                                                default:
                                                                    switch (("\n\t" + "a") + "") {
                                                                    case "hello":
                                                                        if (2 <= 0) {
                                                                            switch ("\n\t" + "\n\t") {
                                                                            case "hello":
                                                                                switch ("foo bar" + "a") {
                                                                                case "":
                                                                                    v46 := v3[(fi14 << ((230) & 63))]
                                                                                    _ = v46
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v15, v4)
                                                                            default:
                                                                                v2 = (v2 | (-(v2)))
                                                                            }
                                                                            fmt.Println(fi40, v26, fi40)
                                                                            fmt.Println(v2)
                                                                        }
                                                                    case "\n\t":
                                                                        v26 = ((680 != 0) || (!v26))
                                                                    }
                                                                }
                                                                v26 = ((741 >= 2147483648) || v26)
                                                                v4 -= (v4 + (v4 & (v15 * v15)))
                                                            }
                                                            _ = fi40
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    } else {
                                                        v47 := GenS5{F0: 186}
                                                        var v48 uint = 255
                                                        var v49 uint8 = 2
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                case 54:
                                                    var v50 uint16 = 42
                                                    _ = v50
                                                }
                                            case 32767:
                                                switch ("" + "a") {
                                                case "hello":
                                                    if len(v1) > 0 {
                                                        v1[0] = fi14
                                                    }
                                                case "foo bar":
                                                    v51 := GenS2{F0: 0, F1: fi14}
                                                    _ = v51
                                                case "\n\t":
                                                    var v52 int16 = 32767
                                                    _ = v52
                                                default:
                                                    var v53 int16 = 850
                                                    _ = v53
                                                }
                                            }
                                        } else {
                                            for fi54 := 0; fi54 < 4; fi54++ {
                                                v26 = v26
                                                var v55 int = 698
                                                v1 = append(v1, fi54)
                                                if len(v1) > 0 {
                                                    v1[0] = v55
                                                }
                                                v56 := v3[(fi14 + fi14)]
                                                v57 := v3[127]
                                                _ = len(v3)
                                                _ = fi54
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            }
                                        }
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v58 int16 = 32767
                                        _ = v58
                                    }
                                }
                                delete(v3, 0)
                            } else {
                                fmt.Println(v2)
                                fmt.Println(v2, fi14)
                                v2 = -962
                            }
                            var v59 uint32 = 282
                            _ = v59
                        }
                        _ = fi14
                        _ = v15
                    }
                case -9223372036854775808:
                    for fi60 := 0; fi60 < 7; fi60++ {
                        v4 -= (-(v4))
                        _ = fi60
                    }
                default:
                    var v61 bool = (int8(v4) > 1)
                    _ = v61
                }
            default:
                var v62 int8 = 100
                _ = v62
            }
            v2 = v2
            v63 := GenS6{F0: 42}
            _ = v63
        } else {
            v4 -= (v4 + v4)
        }
    case "\n\t":
        switch "\n\t" {
        case "":
            fmt.Println(v4, v2)
        case "a":
            fmt.Println(v2, v2)
        case "hello":
            if ((!(0.1 != float64(((float32(1.5) - float32(3.14159)) + ((((-2.25 + (float32(((0.0 / 1.0) - 0.1)) / 0.1)) / 3.14159) - (0.1 * 0.1)) * -2.25))))) && (!true)) {
                _ = len(v3)
                var v64 uint16 = 500
                _ = v64
            } else {
                switch int(v2) {
                case 215:
                    var v65 bool = (!true)
                    _ = v65
                case 0:
                    if (false || (32767 >= 255)) {
                        if (2 != 2) {
                            if false {
                                v4 = (v4 % ((v4) | 1))
                                fmt.Println(v2, v2)
                            }
                        } else {
                            for fi66 := 0; fi66 < 5; fi66++ {
                                switch ("hello" + ("foo bar" + "")) {
                                case "a":
                                    if (65535 < 32767) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch ("" + "") {
                                        case "hello":
                                            _ = len(v3)
                                        case "a":
                                            v2 = v2
                                        case "":
                                            if (!(v4 != -530)) {
                                                var v67 int8 = 100
                                                var v68 int16 = 42
                                                _ = v67
                                                _ = v68
                                            }
                                        default:
                                            var v69 int = fi66
                                            _ = v69
                                        }
                                        v70 := GenS2{F0: v4, F1: (fi66 | fi66)}
                                        _ = v70
                                    }
                                case "":
                                    var v71 uint16 = 295
                                    _ = v71
                                default:
                                    var v72 uint32 = 2
                                    _ = v72
                                }
                                _ = len(v1)
                                switch fi66 {
                                case 0:
                                    v73 := GenS4{F0: (fi66 & fi66), F1: 1, F2: (^(v4))}
                                    _ = v73
                                default:
                                    v3[(fi66 ^ fi66)] = (919 == 18446744073709551615)
                                }
                                _ = fi66
                            }
                        }
                        var v74 int = -1
                        _ = v74
                    } else {
                        v4 = (v4 | 0)
                        var v75 float32 = 0.0
                        _ = v75
                    }
                case 127:
                    var v76 int64 = v4
                    _ = v76
                default:
                    if false {
                        var v77 uint16 = 32767
                        var v78 int32 = -2147483648
                        _ = v77
                        _ = v78
                    }
                }
                v2 = 42
            }
        }
    default:
        if (((((("\n\t" + "foo bar") + "foo bar") + "") + "hello") >= "a") || (0 > 42)) {
            var v79 float64 = 1.5
            if false {
                for fi80 := 0; fi80 < 5; fi80++ {
                    v79 += v79
                    var v81 int64 = -1
                    var v82 int = fi80
                    _ = fi80
                    _ = v81
                    _ = v82
                }
            }
            for fi83 := 0; fi83 < 5; fi83++ {
                v2 += 0
                fmt.Println(v2, v2, v79)
                for fi84 := 0; fi84 < 15; fi84++ {
                    fmt.Println(fi83, v2)
                    v2 += v2
                    _ = fi84
                }
                _ = fi83
            }
            _ = v79
        }
    }
    for fi85 := 0; fi85 < 11; fi85++ {
        if (32767 != 187) {
            if true {
                v86 := GenS2{F0: (v4 ^ (v4 & v4)), F1: 127}
                v2 -= 1
                fmt.Println(v2)
                _ = v86
            } else {
                v4 += (v4 % (((v4 >> ((193) & 63))) | 1))
                var v87 int32 = -1
                v4 -= 206
                _ = v87
            }
        }
        if true {
            _ = len(v3)
            for fi88 := 0; fi88 < 3; fi88++ {
                if ("" >= "") {
                    if (((v2 | v2) > v2) && true) {
                        if (!(false || false)) {
                            _ = len(v3)
                            var v89 float64 = (1e10 * 3.14159)
                            _ = v89
                        }
                        if (true || true) {
                            v90 := v3[(fi88 ^ int(v2))]
                            _ = v90
                        }
                        var v91 int32 = (-(v2))
                        _ = v91
                    } else {
                        var v92 int16 = 0
                        _ = v92
                    }
                    if true {
                        var v93 uint32 = 2147483648
                        var v94 int32 = -1
                        _ = v93
                        _ = v94
                    } else {
                        v4 = (v4 + v4)
                        v2 = 2147483647
                    }
                } else {
                    v2 += int32(fi88)
                    fmt.Println(v4, fi85)
                    v3[fi88] = false
                }
                if (true || (((746 <= 558) || (946 != 380)) && ((65535 != 2) && (!true)))) {
                    _ = len(v3)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi88
            }
            if (0 == (-(fi85))) {
                if false {
                    if len(v1) > 0 {
                        v1[0] = -81
                    }
                }
            }
        }
        _ = fi85
    }
    if ((1e10 > (-2.25 + (0.0 * float64((0.0 - 0.0))))) && (false && (false || false))) {
        for fi95 := 0; fi95 < 1; fi95++ {
            v4 = v4
            var v96 int16 = -1
            var v97 float32 = 1e10
            _ = fi95
            _ = v96
            _ = v97
        }
    } else {
        v4 -= v4
        v2 -= (v2 % (((v2 % (((-(v2))) | 1))) | 1))
        v2 -= (v2 + v2)
    }
    for fi98 := 0; fi98 < 19; fi98++ {
        fmt.Println(v2)
        _ = fi98
    }
    v2 += (v2 * (v2 - v2))
    v99 := GenS6{F0: 32767}
    var v100 uint8 = 2
    switch -9223372036854775808 {
    case -2147483648:
        v3[128] = (((-1 <= 32767) && (-52 == 0)) && (44 >= 1))
    }
    for fi101 := 0; fi101 < 12; fi101++ {
        v2 = (v2 - (v2 + v2))
        _ = fi101
    }
    fmt.Println(v4, v4)
    for fi102 := 0; fi102 < 15; fi102++ {
        fmt.Println(v4, v100)
        _ = fi102
    }
    _ = len(v3)
    var v103 bool = (!false)
    v100 = v100
    fmt.Println(v100, v103)
    v104 := GenS7{F0: -958}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v99
}
